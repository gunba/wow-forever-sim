#!/usr/bin/env python3
"""Turn a learned policy into an APL: policy extraction into a priority list, verified in the simulator.

1. The network plays greedy fights; every decision is logged (observation, legal actions, its choice and how strongly
   it preferred it). Later rounds let the candidate APL play and the network label the states the APL reaches
   (DAgger), so the list learns to recover from its own mistakes.
2. A decision list is induced with APL semantics, where the first ready entry fires: each position takes the action
   and condition (up to two thresholds on observation features) that best separate "the network chose it" from "it
   was ready but the network did something else", weighted by the network's preference. Every action the network
   uses at least --min-uses times a fight gets an entry.
3. Each condition compiles to the APL value that reads the same quantity (the environment exports the map), so the
   list becomes an APL that means the same thing.
4. The APL is checked on paired fights (one random stream per roll label) against the published APL and the network;
   thresholds are then tuned and entries that do nothing are removed, on the same paired fights.

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


def induce(X, mask, y, w, actions, usable, wait, max_entries=24, alpha=1.0, min_uses=None, per_fight=None, quantiles=24,
           kinds=None, cap=30_000, seed=5):
    """Greedy sequential covering for an ordered decision list with APL semantics.

    X: features (n, D); mask: legal actions (n, A); y: the network's action; w: weight per decision; actions: candidate
    action indices; usable: feature indices with an APL counterpart; kinds[f]: "bool" or numeric. Splits are searched on
    at most cap decisions per action (sampled), then applied to all."""
    rng = np.random.default_rng(seed)
    n = len(X)
    open_ = np.ones(n, bool)
    entries = []
    cands = {}
    for f in usable:
        if kinds[f] == "bool":
            continue
        v = np.unique(np.quantile(X[:, f], np.linspace(0.02, 0.98, quantiles)))
        if len(v) > 1:
            cands[f] = v
    must = set(a for a in actions if min_uses and per_fight is not None and per_fight.get(a, 0) >= min_uses)

    def best_literal(idx, pos_w, neg_w, used):
        """The literal maximising pos - alpha * neg over the decisions idx, with its gain; pos_w/neg_w are per idx."""
        best_g, best_lit = -np.inf, None
        for f in usable:
            if f in used:
                continue
            x = X[idx, f]
            if kinds[f] == "bool":
                on = x > 0.5
                for op, m in (("is", on), ("not", ~on)):
                    g = pos_w[m].sum() - alpha * neg_w[m].sum()
                    if g > best_g:
                        best_g, best_lit = g, (f, op, 0.5)
                continue
            ts = cands.get(f)
            if ts is None:
                continue
            L = len(ts)
            r = np.searchsorted(ts, x, side="right")  # thresholds <= x: x >= ts[k] <=> r >= k + 1
            l = np.searchsorted(ts, x, side="left")   # thresholds < x:  x <= ts[k] <=> l <= k
            pr, nr = np.bincount(r, weights=pos_w, minlength=L + 1), np.bincount(r, weights=neg_w, minlength=L + 1)
            pl, nl = np.bincount(l, weights=pos_w, minlength=L + 1), np.bincount(l, weights=neg_w, minlength=L + 1)
            pge = np.cumsum(pr[::-1])[::-1][1:L + 1]
            nge = np.cumsum(nr[::-1])[::-1][1:L + 1]
            ple, nle = np.cumsum(pl)[:L], np.cumsum(nl)[:L]
            gge, gle = pge - alpha * nge, ple - alpha * nle
            k1, k2 = int(np.argmax(gge)), int(np.argmax(gle))
            if gge[k1] > best_g:
                best_g, best_lit = float(gge[k1]), (f, ">=", float(ts[k1]))
            if gle[k2] > best_g:
                best_g, best_lit = float(gle[k2]), (f, "<=", float(ts[k2]))
        return best_g, best_lit

    def entry_for(a):
        """The best condition (up to two literals) for action a among the open decisions, and its gain."""
        R = np.nonzero(open_ & (mask[:, a] == 1))[0]
        if len(R) == 0:
            return None
        if len(R) > cap:
            R = rng.choice(R, cap, replace=False)
        is_a = y[R] == a
        pos_w, neg_w = np.where(is_a, w[R], 0.0), np.where(~is_a, w[R], 0.0)
        P = pos_w.sum()
        if P <= 0:
            return None
        conds, idx, pw, nw = [], R, pos_w, neg_w
        gain = P - alpha * neg_w.sum()
        for _ in range(2):
            g, lit = best_literal(idx, pw, nw, {c[0] for c in conds})
            if lit is None or g <= gain + 0.01 * P:
                break
            keep = DecisionList.holds(X[idx], [lit])
            conds.append(lit)
            idx, pw, nw, gain = idx[keep], pw[keep], nw[keep], g
        return gain / len(R) * (open_ & (mask[:, a] == 1)).sum(), conds

    for _ in range(max_entries):
        options = []
        for a in actions:
            r = entry_for(a)
            if r is not None:
                options.append((r[0], a, r[1]))
        if not options:
            break
        options.sort(key=lambda o: -o[0])
        gain, a, conds = options[0]
        if gain <= 0:
            forced = [o for o in options if o[1] in must]
            if not forced:
                break
            gain, a, conds = forced[0]
        entries.append((a, conds))
        must.discard(a)
        open_ &= ~(mask[:, a].astype(bool) & DecisionList.holds(X, conds))
    return DecisionList(entries, wait)


def induce_residual(X, mask, logits, base, actions, usable, wait, max_entries=16, kinds=None, cap=40_000, quantiles=24,
                    min_gain_frac=0.002, seed=5):
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
            for _lit in range(2):
                g, lit = best_literal(idx, b, {c[0] for c in conds})
                if lit is None or g <= gain + 1e-9:
                    break
                keep = DecisionList.holds(X[idx], [lit])
                conds.append(lit)
                idx, b, gain = idx[keep], b[keep], g
            if not conds:
                continue  # an unconditional override would replace the base outright
            gain *= scale
            if best is None or gain > best[0]:
                best = (gain, a, conds)
        if best is None or best[0] <= min_gain_frac * total:
            break
        gain, a, conds = best
        entries.append((a, conds))
        open_ &= ~(mask[:, a].astype(bool) & DecisionList.holds(X, conds))
    return DecisionList(entries, wait)


def compile_apl(dl, spec, base_rotation=None):
    """Decision list -> APLRotation JSON, each threshold converted to the APL value's units; with base_rotation, the
    entries go above its priority list. A wait entry becomes a 50 ms APL wait (the engine's poll)."""
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
        out["priorityList"] = items + base_rotation["priorityList"]
        return out
    return {"type": "TypeAPL", "priorityList": items}


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
    ap.add_argument("--mode", choices=["residual", "scratch"], default="residual",
                    help="residual: overrides above the published APL; scratch: a whole new list")
    ap.add_argument("--fights", type=int, default=256, help="fights per data round")
    ap.add_argument("--rounds", type=int, default=3, help="data rounds (1 = the network only; more add DAgger rounds)")
    ap.add_argument("--max-entries", type=int, default=16)
    ap.add_argument("--min-gain", type=float, default=0.0005, help="residual mode: stop proposing below this share of the total preference")
    ap.add_argument("--alpha", type=float, default=1.0, help="scratch mode: cost of firing where the network did not")
    ap.add_argument("--min-uses", type=float, default=0.5, help="scratch mode: actions used this often a fight get an entry")
    ap.add_argument("--samples", type=int, default=400_000, help="decisions kept for induction")
    ap.add_argument("--eval-fights", type=int, default=2048)
    ap.add_argument("--tune-rounds", type=int, default=1)
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
    base_rotation = None
    if a.mode == "residual":
        rows = json.load(open(a.profiles))["Results"]
        base_rotation = next(r["BaselinePlayer"]["rotation"] for r in rows if r["Key"] == a.build and r["Race"] == a.race)
    ck = torch.load(a.model, map_location="cpu", weights_only=False)
    if ck["spec"]["obsDim"] != spec["obsDim"] or ck["spec"]["actDim"] != spec["actDim"]:
        raise SystemExit("model spaces differ from this environment (same build, race and settings?)")
    net = Net(spec["obsDim"], spec["actDim"], ck["net"]["trunk.0.weight"].shape[0]).to(dev)
    net.load_state_dict(ck["net"])
    net.eval()
    residual = a.mode == "residual"

    def rollout(dl, fights, seed0, keep=True):
        """Fights played by the network (dl None) or the list (over the published APL when residual); every decision
        labelled with the network's logits and the published APL's choice."""
        ev = VecEnv(cfg, fights)
        obs, mask, tch = ev.restart(seed0 + np.arange(fights, dtype=np.int64))
        dps = np.full(fights, np.nan)
        Xs, Ms, Ls, Ts = [], [], [], []
        while np.isnan(dps).any():
            live = np.isnan(dps)
            with torch.no_grad():
                logits, _ = net(torch.from_numpy(obs).to(dev), torch.from_numpy(mask).to(dev).bool())
            lg = logits.float().cpu().numpy()
            if dl is None:
                act = lg.argmax(-1).astype(np.int32)
            else:
                act = dl.act(obs, mask, tch if residual else None)
            if keep:
                Xs.append(obs[live].copy()); Ms.append(mask[live].copy()); Ls.append(np.maximum(lg[live], -1e4)); Ts.append(tch[live].astype(np.int32))
            obs, mask, tch, _, done, info = ev.step(act)
            new = (done == 1) & np.isnan(dps)
            dps[new] = info[new, 0]
        ev.close()
        if not keep:
            return dps
        return np.concatenate(Xs), np.concatenate(Ms), np.concatenate(Ls), np.concatenate(Ts), dps

    def apl_dps(rotation, fights, seed0):
        """Paired fights of an APL (the published one when rotation is None) in the plain engine path."""
        c = {**cfg, "fullChoices": False}
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

    def build(Xa, Ma, La, Ta, per_fight):
        Ya = La.argmax(1).astype(np.int32)
        Ya = np.where(np.isin(Ya, list(waits)), wait, Ya)
        if residual:
            Lf = La.astype(np.float32)
            # Every wait variant counts as the plain wait, whose value is the best of them.
            wv = [w_ for w_ in waits]
            Lf[:, wait] = Lf[:, wv].max(1)
            Tb = np.where(np.isin(Ta, list(waits)), wait, Ta)
            cand = sorted({int(v) for v in range(T, T + K)} - (waits - {wait}))
            return induce_residual(Xa, Ma, Lf, Tb, cand, usable, wait, a.max_entries, kinds=kinds, min_gain_frac=a.min_gain)
        srt = np.sort(La.astype(np.float32), -1)
        weight = 0.1 + np.clip(srt[:, -1] - srt[:, -2], 0, 5)
        cand = sorted({int(v) for v in np.unique(Ya)} - {wait})
        return induce(Xa, Ma, Ya, weight, cand, usable, wait, a.max_entries, a.alpha, a.min_uses, per_fight, kinds=kinds)

    t0 = time.time()
    data = None
    dl = None
    log = []
    per_fight = {}
    for rnd in range(a.rounds):
        X, M, L, Tch, dps = rollout(dl, a.fights, 100_000 + 10_000 * rnd)
        if dl is None:
            u, c = np.unique(L.argmax(1), return_counts=True)
            per_fight = {int(k): v / a.fights for k, v in zip(u, c)}
        player = "network" if dl is None else "list"
        log.append({"round": rnd, "player": player, "dps": float(np.mean(dps)), "decisions": int(len(X))})
        print(f"round {rnd}: {player} {np.mean(dps):.1f} DPS over {a.fights} fights, {len(X):,} decisions ({time.time() - t0:.0f}s)", flush=True)
        data = (X, M, L, Tch) if data is None else tuple(np.concatenate([d, n]) for d, n in zip(data, (X, M, L, Tch)))
        Xa, Ma, La, Ta = data
        if len(Xa) > a.samples:
            keep = rng.choice(len(Xa), a.samples, replace=False)
            Xa, Ma, La, Ta = Xa[keep], Ma[keep], La[keep], Ta[keep]
        dl = build(Xa, Ma, La, Ta, per_fight)
        print(f"   {len(dl.entries)} entries", flush=True)
        log[-1]["entries"] = len(dl.entries)

    def paired(x, y):
        d = x - y
        return float(d.mean()), float(d.std() / math.sqrt(len(d)))

    # Verify: the compiled APL in the engine, the same list evaluated here, the published APL and the network.
    F, S = a.eval_fights, 7_000_000
    apl = compile_apl(dl, spec, base_rotation)
    pub = apl_dps(None, F, S)
    syn = apl_dps(apl, F, S)
    lst = rollout(dl, F, S, keep=False)
    net_dps = rollout(None, F, S, keep=False)
    print(f"published {pub.mean():.1f}; compiled APL {syn.mean():.1f} ({paired(syn, pub)[0]:+.2f} ± {paired(syn, pub)[1]:.2f}); "
          f"same list here {lst.mean():.1f}; network {net_dps.mean():.1f} ({paired(net_dps, pub)[0]:+.2f})", flush=True)

    # Selection on paired fights: the induced entries in order, each kept only if it raises DPS; then thresholds.
    TF, TS = max(512, F // 2), 9_000_000
    if residual:
        best_dl = DecisionList([], wait)
        cur = apl_dps(base_rotation, TF, TS).mean()
        for e in dl.entries:
            trial = DecisionList(best_dl.entries + [e], wait)
            m = apl_dps(compile_apl(trial, spec, base_rotation), TF, TS).mean()
            if m > cur + 0.3:
                best_dl, cur = trial, m
    else:
        best_dl = dl
        cur = apl_dps(compile_apl(best_dl, spec, base_rotation), TF, TS).mean()
    for _ in range(a.tune_rounds):
        for e in range(len(best_dl.entries)):
            act_, conds = best_dl.entries[e]
            for j in range(len(conds)):
                f, op, t = conds[j]
                if op not in (">=", "<="):
                    continue
                for factor in (0.85, 1.15):
                    c2 = list(conds)
                    c2[j] = (f, op, t * factor)
                    trial = list(best_dl.entries)
                    trial[e] = (act_, c2)
                    cdl = DecisionList(trial, wait)
                    m = apl_dps(compile_apl(cdl, spec, base_rotation), TF, TS).mean()
                    if m > cur + 0.3:
                        best_dl, cur, conds = cdl, m, c2
    best_apl = compile_apl(best_dl, spec, base_rotation)
    tuned = apl_dps(best_apl, F, S)
    result = {"build": a.build, "race": a.race, "mode": a.mode, "fights": F, "published": float(pub.mean()),
              "compiled_untuned": float(syn.mean()), "list_untuned_in_env": float(lst.mean()),
              "synthesised": float(tuned.mean()), "synthesised_vs_published": paired(tuned, pub),
              "network": float(net_dps.mean()), "network_vs_published": paired(net_dps, pub),
              "entries": len(best_dl.entries), "rounds": log, "minutes": round((time.time() - t0) / 60, 1)}
    (a.out / "apl.json").write_text(json.dumps(best_apl, indent=2) + "\n")
    (a.out / "rules.txt").write_text(describe(best_dl, spec) + "\n")
    (a.out / "report.json").write_text(json.dumps(result, indent=1) + "\n")
    print(describe(best_dl, spec))
    print(json.dumps(result, indent=1), flush=True)


if __name__ == "__main__":
    main()
