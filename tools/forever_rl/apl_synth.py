#!/usr/bin/env python3
"""Turn a learned policy into an APL: edits to the published APL, each chosen in the simulator.

Each iteration starts from the current APL (first the published one), which the environment plays as its teacher:
1. The network plays fights and the APL plays fights; in every state the network's logits and the APL's choice are
   logged, so the network labels both its own states and the ones the APL reaches.
2. Overrides are induced on top of the APL: each an action and up to two thresholds on observation features, placed
   where the network prefers that action to the APL's choice (logit difference). Each threshold compiles to the APL
   value that reads the same quantity (the environment exports the map). Actions the network takes far more often
   than the APL are also offered unconditionally.
3. The simulator selects, on paired fights (one random stream per roll label): removing each APL entry, then each
   override and promotion, every edit kept only if it raises DPS by --accept-se standard errors; thresholds are then
   tuned the same way. The result is the next iteration's APL; iterations stop when nothing is kept.
The final APL is checked on fresh paired fights against the published one and the network.

    python3 tools/forever_rl/apl_synth.py --build fury --race Orc --model runs/fury/model.pt --out runs/fury_apl
"""
import argparse
import json
import math
import time
from pathlib import Path

import numpy as np
import torch

from forever_env import VecEnv
from train import Net


class DecisionList:
    """Entries (action index, [(feature, op, threshold)]); op is ">=", "<=", "is" (bool true) or "not" (bool false).
    The first entry whose action is legal and whose condition holds fires; otherwise the policy waits."""

    def __init__(self, entries, wait):
        self.entries, self.wait = entries, wait

    @staticmethod
    def holds(X, conds):
        ok = np.ones(len(X), bool)
        for f, op, t in conds:
            x = X[:, f]
            if op == ">=":
                ok &= x >= t
            elif op == "<=":
                ok &= x <= t
            elif op == "is":
                ok &= x > 0.5
            else:
                ok &= x <= 0.5
        return ok

    def act(self, X, mask, base=None):
        """Actions for a batch; with base (the published APL's choices), entries override it and the rest follow it."""
        out = np.full(len(X), self.wait, np.int32) if base is None else base.astype(np.int32).copy()
        open_ = np.ones(len(X), bool)
        for a, conds in self.entries:
            fire = open_ & (mask[:, a] == 1) & self.holds(X, conds)
            out[fire] = a
            open_ &= ~fire
        return out


def induce_residual(X, mask, logits, base, actions, usable, wait, max_entries=16, kinds=None, cap=40_000, quantiles=24,
                    min_gain_frac=0.002, seed=5, literals=2):
    """Overrides on top of a base policy (the published APL): each entry fires where its action is legal and its
    condition holds, before the base list. The value of firing action a in a state is the network's preference for a
    over the base's choice there (logit difference, clipped to +-5): positive where the network would rather do a,
    negative where the base was right. Greedy: each position takes the (action, up to two thresholds) with the
    largest total value over the states still open, until nothing gains more than min_gain_frac of the total."""
    rng = np.random.default_rng(seed)
    n = len(X)
    open_ = np.ones(n, bool)
    lb = logits[np.arange(n), base]
    total = float(np.clip(logits.max(1) - lb, 0, 5).sum())
    cands = {}
    for f in usable:
        if kinds[f] != "bool":
            v = np.unique(np.quantile(X[:, f], np.linspace(0.02, 0.98, quantiles)))
            if len(v) > 1:
                cands[f] = v
    entries = []

    def best_literal(idx, b, used):
        pos, neg = np.maximum(b, 0), np.maximum(-b, 0)
        best_g, best_lit = -np.inf, None
        for f in usable:
            if f in used:
                continue
            x = X[idx, f]
            if kinds[f] == "bool":
                on = x > 0.5
                for op, m in (("is", on), ("not", ~on)):
                    g = b[m].sum()
                    if g > best_g:
                        best_g, best_lit = float(g), (f, op, 0.5)
                continue
            ts = cands.get(f)
            if ts is None:
                continue
            L = len(ts)
            r = np.searchsorted(ts, x, side="right")
            l = np.searchsorted(ts, x, side="left")
            br = np.bincount(r, weights=b, minlength=L + 1)
            bl = np.bincount(l, weights=b, minlength=L + 1)
            gge = np.cumsum(br[::-1])[::-1][1:L + 1]
            gle = np.cumsum(bl)[:L]
            k1, k2 = int(np.argmax(gge)), int(np.argmax(gle))
            if gge[k1] > best_g:
                best_g, best_lit = float(gge[k1]), (f, ">=", float(ts[k1]))
            if gle[k2] > best_g:
                best_g, best_lit = float(gle[k2]), (f, "<=", float(ts[k2]))
        return best_g, best_lit

    for _ in range(max_entries):
        best = None
        for a in actions:
            Rall = np.nonzero(open_ & (mask[:, a] == 1) & (base != a))[0]
            if len(Rall) == 0:
                continue
            R = rng.choice(Rall, cap, replace=False) if len(Rall) > cap else Rall
            scale = len(Rall) / len(R)
            bval = np.clip(logits[R, a] - lb[R], -5, 5)
            if bval.max() <= 0:
                continue
            conds, idx, b = [], R, bval
            gain = float(b.sum())
            for _lit in range(literals):
                g, lit = best_literal(idx, b, {c[0] for c in conds})
                # The first literal is taken whenever it gains at all (an unconditional override would replace the
                # base outright); a second one only if it adds to the first.
                if lit is None or g <= 0 or (conds and g <= gain + 1e-9):
                    break
                keep = DecisionList.holds(X[idx], [lit])
                conds.append(lit)
                idx, b, gain = idx[keep], b[keep], g
            if not conds:
                continue
            gain *= scale
            if best is None or gain > best[0]:
                best = (gain, a, conds)
        if best is None or best[0] <= min_gain_frac * total:
            break
        gain, a, conds = best
        entries.append((a, conds))
        open_ &= ~(mask[:, a].astype(bool) & DecisionList.holds(X, conds))
    return DecisionList(entries, wait)


def compile_apl(dl, spec, base_rotation=None, removed=()):
    """Decision list -> APLRotation JSON, each threshold converted to the APL value's units; with base_rotation, the
    entries go above its priority list, less the entries at the indices in removed. A wait entry becomes a 50 ms APL
    wait (the engine's poll)."""
    fa, aa = spec["featureAPL"], spec["actionAPL"]
    T = spec["talents"]
    items = []
    for a, conds in dl.entries:
        if aa[a - T] is None:
            act = {"wait": {"duration": {"const": {"val": "0.05s"}}}}
        else:
            act = dict(aa[a - T])
        lits = []
        for f, op, t in conds:
            m = fa[f]
            v = m["value"]
            if op == "is":
                lits.append(v)
            elif op == "not":
                lits.append({"not": {"val": v}})
            else:
                x = t * m["scale"]
                val = f"{x:.3f}s" if m["kind"] == "time" else f"{x:.4g}"
                lits.append({"cmp": {"op": "OpGe" if op == ">=" else "OpLe", "lhs": v, "rhs": {"const": {"val": val}}}})
        if len(lits) == 1:
            act["condition"] = lits[0]
        elif lits:
            act["condition"] = {"and": {"vals": lits}}
        items.append({"action": act})
    if base_rotation is not None:
        # Everything else in the base (prepull actions above all) stays as it is.
        out = json.loads(json.dumps(base_rotation))
        out["priorityList"] = items + [it for i, it in enumerate(base_rotation["priorityList"]) if i not in removed]
        return out
    return {"type": "TypeAPL", "priorityList": items}


def describe_item(item, spec):
    """An APL entry in words: its action as the agent knows it and its condition, in the observation features' terms
    where the condition is one compile_apl writes; other conditions are only marked."""
    act = dict(item.get("action", {}))
    cond = act.pop("condition", None)
    name = None
    for k, raw in enumerate(spec["actionAPL"]):
        if raw is not None and raw == act:
            name = spec["actionNames"][spec["talents"] + k]
            break
    if name is None:
        name = "wait" if act == {"wait": {"duration": {"const": {"val": "0.05s"}}}} else json.dumps(act, separators=(",", ":"))[:160]
    if cond is None:
        return name
    fa, fnames = spec["featureAPL"], spec["featureNames"]

    def feature(value):
        for f, m in enumerate(fa):
            if m is not None and m["value"] == value:
                return f
        return None

    def literal(c):
        if "not" in c:
            f = feature(c["not"]["val"])
            return None if f is None else "not " + fnames[f]
        if "cmp" in c and c["cmp"]["op"] in ("OpGe", "OpLe") and "const" in c["cmp"]["rhs"]:
            f = feature(c["cmp"]["lhs"])
            if f is None:
                return None
            val = c["cmp"]["rhs"]["const"]["val"]
            x = float(val.rstrip("s"))
            unit = "s" if fa[f]["kind"] == "time" else ""
            return f"{fnames[f]} {'>=' if c['cmp']['op'] == 'OpGe' else '<='} {x:.2f}{unit}"
        f = feature(c)
        return None if f is None else fnames[f]

    parts = [literal(c) for c in (cond["and"]["vals"] if "and" in cond else [cond])]
    if any(x is None for x in parts):
        return name + " (conditional)"
    return name + " if " + " and ".join(parts)


def apl_diff(published, final):
    """(added, removed): final is a block of added entries above a subsequence of published; the block is the shortest
    prefix whose rest is such a subsequence (an added entry may equal a published one)."""
    key = [json.dumps(x, sort_keys=True) for x in published]
    fin = [json.dumps(x, sort_keys=True) for x in final]

    def matched(rest):
        """Indices of published matched in order by rest, or None if rest is not a subsequence."""
        out, j = [], 0
        for x in rest:
            while j < len(key) and key[j] != x:
                j += 1
            if j == len(key):
                return None
            out.append(j)
            j += 1
        return out

    for i in range(len(fin) + 1):
        m = matched(fin[i:])
        if m is not None:
            kept = set(m)
            return final[:i], [x for j, x in enumerate(published) if j not in kept]
    return final, list(published)


def describe(dl, spec):
    names, fnames, fa = spec["actionNames"], spec["featureNames"], spec["featureAPL"]
    lines = []
    for k, (a, conds) in enumerate(dl.entries, 1):
        parts = []
        for f, op, t in conds:
            m = fa[f]
            if op in ("is", "not"):
                parts.append(("not " if op == "not" else "") + fnames[f])
            else:
                x = t * m["scale"]
                parts.append(f"{fnames[f]} {op} {x:.2f}{'s' if m['kind'] == 'time' else ''}")
        lines.append(f"{k:2d}. {names[a]}" + (" if " + " and ".join(parts) if parts else ""))
    return "\n".join(lines)


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--build", required=True)
    ap.add_argument("--race", required=True)
    ap.add_argument("--model", type=Path, required=True)
    ap.add_argument("--talents", default="", help="talents (default: the published ones)")
    ap.add_argument("--iterations", type=int, default=3, help="induce-and-select iterations (stops early when nothing is kept)")
    ap.add_argument("--decisions", type=int, default=400_000, help="decisions logged per iteration, half from each player")
    ap.add_argument("--min-fights", type=int, default=128)
    ap.add_argument("--max-fights", type=int, default=4096)
    ap.add_argument("--max-entries", type=int, default=16)
    ap.add_argument("--literals", type=int, default=2, help="thresholds per override at most")
    ap.add_argument("--min-gain", type=float, default=0.0005, help="stop proposing below this share of the total preference")
    ap.add_argument("--promote", type=float, default=0.5, help="actions the network takes this much more a fight than the APL are offered unconditionally")
    ap.add_argument("--samples", type=int, default=400_000, help="decisions kept for induction")
    ap.add_argument("--select-decisions", type=int, default=1_000_000, help="decisions per selection test (sets its paired fights)")
    ap.add_argument("--accept-se", type=float, default=2.0, help="a selection edit must gain this many standard errors")
    ap.add_argument("--eval-fights", type=int, default=2048)
    ap.add_argument("--threads", type=int, default=24)
    ap.add_argument("--profiles", type=Path, default=Path(__file__).resolve().parents[2] / "artifacts/modelled_gear/forever_input_profiles.json")
    ap.add_argument("--out", type=Path, required=True)
    a = ap.parse_args()
    a.out.mkdir(parents=True, exist_ok=True)
    dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    rng = np.random.default_rng(3)
    cfg = dict(build=a.build, race=a.race, fullChoices=True, pollSeconds=0, rewardScale=1e-5, threads=a.threads,
               labeledRands=True, seed=17)
    if a.talents:
        cfg["talents"] = a.talents
    probe = VecEnv(cfg, 1)
    spec = probe.spec
    probe.close()
    T, K = spec["talents"], spec["agentActions"]
    wait = T + K - 1
    names = spec["actionNames"]
    waits = {T + k for k in range(K) if names[T + k].startswith("wait")}
    kinds = [m["kind"] if m else None for m in spec["featureAPL"]]
    usable = [f for f, m in enumerate(spec["featureAPL"]) if m is not None]
    published = next(r["BaselinePlayer"] for r in json.load(open(a.profiles))["Results"] if r["Key"] == a.build and r["Race"] == a.race)
    ck = torch.load(a.model, map_location="cpu", weights_only=False)
    if ck["spec"]["obsDim"] != spec["obsDim"] or ck["spec"]["actDim"] != spec["actDim"]:
        raise SystemExit("model spaces differ from this environment (same build, race and settings?)")
    net = Net(spec["obsDim"], spec["actDim"], ck["net"]["trunk.0.weight"].shape[0]).to(dev)
    net.load_state_dict(ck["net"])
    net.eval()

    def rollout(teacher, fights, seed0, network=True, keep=True):
        """Fights played by the network (or by the teacher APL), every decision labelled with the network's logits
        and the teacher's choice."""
        ev = VecEnv({**cfg, "teacherRotation": json.dumps(teacher)}, fights)
        obs, mask, tch = ev.restart(seed0 + np.arange(fights, dtype=np.int64))
        dps = np.full(fights, np.nan)
        Xs, Ms, Ls, Ts = [], [], [], []
        while np.isnan(dps).any():
            live = np.isnan(dps)
            with torch.no_grad():
                logits, _ = net(torch.from_numpy(obs).to(dev), torch.from_numpy(mask).to(dev).bool())
            lg = logits.float().cpu().numpy()
            act = lg.argmax(-1).astype(np.int32) if network else tch.astype(np.int32)
            if keep:
                Xs.append(obs[live].copy()); Ms.append(mask[live].copy()); Ls.append(np.maximum(lg[live], -1e4)); Ts.append(tch[live].astype(np.int32))
            obs, mask, tch, _, done, info = ev.step(act)
            new = (done == 1) & np.isnan(dps)
            dps[new] = info[new, 0]
        ev.close()
        if not keep:
            return dps
        return np.concatenate(Xs), np.concatenate(Ms), np.concatenate(Ls), np.concatenate(Ts), dps

    def apl_dps(rotation, fights, seed0, talents=None):
        """Paired fights of an APL (the published one when rotation is None) in the engine's own APL path."""
        c = {**cfg, "fullChoices": False}
        if talents is not None:
            c["talents"] = talents
        if rotation is not None:
            c["rotation"] = json.dumps(rotation)
        ev = VecEnv(c, fights)
        obs, mask, tch = ev.restart(seed0 + np.arange(fights, dtype=np.int64))
        dps = np.full(fights, np.nan)
        while np.isnan(dps).any():
            obs, mask, tch, _, done, info = ev.step(tch.astype(np.int32))
            new = (done == 1) & np.isnan(dps)
            dps[new] = info[new, 0]
        ev.close()
        return dps

    def paired(x, y):
        d = x - y
        return float(d.mean()), float(d.std() / math.sqrt(len(d)))

    def fights_for(decisions, per_fight):
        return int(min(a.max_fights, max(a.min_fights, math.ceil(decisions / per_fight))))

    t0 = time.time()
    base = json.loads(json.dumps(published["rotation"]))
    edits, log = [], []
    per_dec = None
    TS = 9_000_000
    for it in range(a.iterations):
        # 1. Data: the network's states and the APL's own, both labelled by the network against the APL's choice.
        if per_dec is None:
            X, M, L, Tch, d = rollout(base, a.min_fights, 100_000)
            per_dec = len(X) / a.min_fights
            n = fights_for(a.decisions / 2, per_dec) - a.min_fights
            if n > 0:
                parts = [(X, M, L, Tch, d), rollout(base, n, 100_000 + a.min_fights)]
                X, M, L, Tch, d = (np.concatenate(z) for z in zip(*parts))
            TF = fights_for(a.select_decisions, per_dec)
        else:
            X, M, L, Tch, d = rollout(base, fights_for(a.decisions / 2, per_dec), 100_000 + 50_000 * it)
        net_dps, nf = float(np.mean(d)), len(d)
        X2, M2, L2, T2, d2 = rollout(base, fights_for(a.decisions / 2, per_dec), 1_000_000 + 50_000 * it, network=False)
        # Actions the network takes far more often than the APL would in the same states.
        net_n = np.bincount(L.argmax(1), minlength=T + K)[T:] / nf
        apl_n = np.bincount(Tch.astype(np.int64), minlength=T + K)[T:] / nf
        promotions = [(T + int(k), []) for k in np.argsort(-(net_n - apl_n))
                      if net_n[k] - apl_n[k] >= a.promote and T + k != wait and T + k not in waits]
        # 2. Overrides on top of the current APL, induced on the network's states, the APL's and both.
        cand = sorted({int(v) for v in range(T, T + K)} - (waits - {wait}))
        proposals = []
        views = ((X, M, L, Tch), (X2, M2, L2, T2), tuple(np.concatenate([u, v]) for u, v in zip((X, M, L, Tch), (X2, M2, L2, T2))))
        for Xa, Ma, La, Ta in views:
            if len(Xa) > a.samples:
                keep = rng.choice(len(Xa), a.samples, replace=False)
                Xa, Ma, La, Ta = Xa[keep], Ma[keep], La[keep], Ta[keep]
            Lf = La.astype(np.float32)
            Lf[:, wait] = Lf[:, sorted(waits)].max(1)  # every wait variant counts as the plain wait, worth the best of them
            Tb = np.where(np.isin(Ta, list(waits)), wait, Ta)
            for e in induce_residual(Xa, Ma, Lf, Tb, cand, usable, wait, a.max_entries, kinds=kinds, min_gain_frac=a.min_gain,
                                     literals=a.literals).entries:
                if e not in proposals:
                    proposals.append(e)
        dl = DecisionList(proposals, wait)
        print(f"iteration {it}: network {net_dps:.1f} and APL {np.mean(d2):.1f} DPS (teacher path), {len(X) + len(X2):,} decisions, "
              f"{len(dl.entries)} overrides proposed, {len(promotions)} promotions ({time.time() - t0:.0f}s)", flush=True)
        for line in describe(dl, spec).splitlines():
            print("      " + line, flush=True)

        # 3. Selection in the simulator.
        items = base["priorityList"]
        removed, chosen = set(), []

        def assemble(entries, rem):
            return compile_apl(DecisionList(entries, wait), spec, base, rem)

        cur = apl_dps(base, TF, TS)
        start = cur.mean()

        def better(rot):
            dd = apl_dps(rot, TF, TS)
            g, se = paired(dd, cur)
            return g > max(0.2, a.accept_se * se), dd, g, se

        for pass_ in range(2):
            changed = False
            for i in range(len(items)):
                if i in removed:
                    continue
                ok, dd, g, se = better(assemble(chosen, removed | {i}))
                if ok:
                    removed.add(i)
                    cur, changed = dd, True
                    edits.append({"iteration": it, "edit": "remove", "entry": describe_item(items[i], spec), "gain": g, "se": se})
                    print(f"   remove {edits[-1]['entry']}: {g:+.2f} ± {se:.2f}", flush=True)
            # Best first: every proposal on its own, then added in order of its own gain while each still helps.
            order = []
            for e in (dl.entries + promotions) if pass_ == 0 else []:
                ok, dd, g, se = better(assemble(chosen + [e], removed))
                if ok:
                    order.append((g, e))
            for _, e in sorted(order, key=lambda x: -x[0]):
                ok, dd, g, se = better(assemble(chosen + [e], removed))
                if ok:
                    chosen.append(e)
                    cur, changed = dd, True
                    edits.append({"iteration": it, "edit": "add", "entry": describe(DecisionList([e], wait), spec).split(". ", 1)[1],
                                  "gain": g, "se": se})
                    print(f"   add {edits[-1]['entry']}: {g:+.2f} ± {se:.2f}", flush=True)
            if not changed:
                break
        for e in range(len(chosen)):
            act_, conds = chosen[e]
            for j in range(len(conds)):
                for factor in (0.85, 1.15):
                    f, op, t = conds[j]
                    if op not in (">=", "<="):
                        continue
                    c2 = list(conds)
                    c2[j] = (f, op, t * factor)
                    trial = chosen[:e] + [(act_, c2)] + chosen[e + 1:]
                    ok, dd, g, se = better(assemble(trial, removed))
                    if ok:
                        chosen, cur, conds = trial, dd, c2
        log.append({"iteration": it, "network_teacher_path": net_dps, "apl_teacher_path": float(np.mean(d2)),
                    "decisions": int(len(X) + len(X2)), "proposed": len(dl.entries), "promotions": len(promotions),
                    "kept": len(chosen), "removed": len(removed), "select_fights": TF,
                    "gain": float(cur.mean() - start)})
        if not chosen and not removed:
            break
        base = assemble(chosen, removed)

    # Verify on fresh fights: the published APL, the result, the network; and the published setup (published talents).
    F, S = a.eval_fights, 7_000_000
    pub = apl_dps(None, F, S)
    tuned = apl_dps(base, F, S)
    net_dps = rollout(published["rotation"], F, S, keep=False)
    ref = pub if (a.talents or published["talentsString"]) == published["talentsString"] else \
        apl_dps(None, F, S, talents=published["talentsString"])
    # The final APL against the published one: edits always add above everything, so it is a block of added entries
    # followed by the published entries that remain, in order.
    added, dropped = apl_diff(published["rotation"]["priorityList"], base["priorityList"])
    added = [describe_item(x, spec) for x in added]
    dropped = [describe_item(x, spec) for x in dropped]
    result = {"build": a.build, "race": a.race, "fights": F, "published": float(pub.mean()),
              "synthesised": float(tuned.mean()), "synthesised_vs_published": paired(tuned, pub),
              "network": float(net_dps.mean()), "network_vs_published": paired(net_dps, pub),
              "talents": a.talents or published["talentsString"], "reference": float(ref.mean()),
              "synthesised_vs_reference": paired(tuned, ref), "network_vs_reference": paired(net_dps, ref),
              "added": added, "removed": dropped, "entries": len(added), "edits": edits, "iterations": log,
              "minutes": round((time.time() - t0) / 60, 1)}
    (a.out / "apl.json").write_text(json.dumps(base, indent=2) + "\n")
    (a.out / "rules.txt").write_text("".join(f"{k}. {x}\n" for k, x in enumerate(added, 1))
                                     + "".join(f"removed: {x}\n" for x in dropped))
    (a.out / "report.json").write_text(json.dumps(result, indent=1) + "\n")
    print(f"published {pub.mean():.1f}; APL {tuned.mean():.1f} ({result['synthesised_vs_published'][0]:+.2f} ± "
          f"{result['synthesised_vs_published'][1]:.2f}); network {net_dps.mean():.1f} ({result['network_vs_published'][0]:+.2f}); "
          f"{len(added)} added, {len(dropped)} removed ({result['minutes']} min)", flush=True)


if __name__ == "__main__":
    main()
