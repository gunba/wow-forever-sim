#!/usr/bin/env python3
"""What a trained policy does differently from the published APL: both play the same fights (one random stream per
roll label), and every decision is recorded with the fight time, rage and phase. Prints actions per fight in and out of
the execute phase, when each cooldown is used, and the rage at which rage-spending actions are chosen.

    python3 tools/forever_rl/explain.py --build fury --race Orc --model runs/fury_codesign/model.pt --fights 512
"""
import argparse
import collections
import json
from pathlib import Path

import numpy as np
import torch

from codesign import Builds
from forever_env import VecEnv
from train import Net


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--build", default="fury")
    ap.add_argument("--race", default="Orc")
    ap.add_argument("--talents", default="published", help="the APL's talents")
    ap.add_argument("--network-talents", default="", help="the network's talents (default: the APL's)")
    ap.add_argument("--timeline-seconds", type=float, default=60, help="record fight 0's actions for this long")
    ap.add_argument("--model", type=Path, required=True)
    ap.add_argument("--rotation", type=Path, help="compare against this rotation (APLRotation JSON) instead of the published APL")
    ap.add_argument("--fights", type=int, default=512)
    ap.add_argument("--threads", type=int, default=12)
    ap.add_argument("--json", type=Path)
    a = ap.parse_args()
    cfg = dict(build=a.build, race=a.race, fullChoices=True, pollSeconds=0, rewardScale=1e-5, threads=a.threads,
               labeledRands=True, seed=99, rotation=a.rotation.read_text() if a.rotation else "")
    F = a.fights
    ev = VecEnv(cfg, 2 * F)
    spec = ev.spec
    B = Builds(spec)
    build = B.published if a.talents == "published" else B.canon(a.talents)
    net_build = B.canon(a.network_talents) if a.network_talents else build
    names = spec["actionNames"]
    feat = spec["featureNames"]
    fi = {k: feat.index(k) for k in ("elapsed", "rage", "execute20", "mana_frac", "energy", "combo")}
    dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    ck = torch.load(a.model, map_location="cpu", weights_only=False)
    net = Net(spec["obsDim"], spec["actDim"], ck["net"]["trunk.0.weight"].shape[0]).to(dev)
    net.load_state_dict(ck["net"])
    net.eval()
    for i in range(2 * F):
        ev.set_talents(i, build if i < F else net_build)
    obs, mask, tch = ev.restart(np.tile(31_337 + np.arange(F, dtype=np.int64), 2))
    is_net = np.arange(2 * F) >= F
    dps = np.full(2 * F, np.nan)
    timeline = [[], []]  # fight 0 of each player: (time, action, rage, mana %, energy, combo points)
    dur = spec["duration"] / 1e9 if spec["duration"] > 1e6 else spec["duration"]
    rec = {k: [] for k in ("act", "t", "rage", "exec", "env", "tch")}
    while np.isnan(dps).any():
        with torch.no_grad():
            logits, _ = net(torch.from_numpy(obs).to(dev), torch.from_numpy(mask).to(dev).bool())
        greedy = logits.argmax(-1).cpu().numpy().astype(np.int32)
        act = np.where(is_net, greedy, tch.astype(np.int32))
        live = np.nonzero(np.isnan(dps))[0]
        t_all = obs[live, fi["elapsed"]] * dur
        rec["act"].append(act[live]); rec["t"].append(t_all); rec["rage"].append(obs[live, fi["rage"]] * 100)
        rec["exec"].append(obs[live, fi["execute20"]] > 0.5); rec["env"].append(live); rec["tch"].append(tch[live].astype(np.int32))
        for i in (0, F):
            if np.isnan(dps[i]):
                t = float(obs[i, fi["elapsed"]]) * dur
                if t <= a.timeline_seconds and names[int(act[i])] != "wait":
                    timeline[int(i >= F)].append({"t": round(t, 2), "action": names[int(act[i])], "rage": round(float(obs[i, fi["rage"]]) * 100),
                                                  "mana_pct": round(float(obs[i, fi["mana_frac"]]) * 100),
                                                  "energy": round(float(obs[i, fi["energy"]]) * 100), "combo": round(float(obs[i, fi["combo"]]) * 5)})
        obs, mask, tch, _, done, info = ev.step(act)
        new = (done == 1) & np.isnan(dps)
        dps[new] = info[new, 0]
    ev.close()
    R_ = {k: np.concatenate(v) for k, v in rec.items()}
    who = (R_["env"] >= F).astype(np.int64)
    log = [collections.defaultdict(list), collections.defaultdict(list)]
    for p_ in (0, 1):
        sel = who == p_
        acts, ts, rg, ex, en = R_["act"][sel], R_["t"][sel], R_["rage"][sel], R_["exec"][sel], R_["env"][sel]
        order = np.argsort(acts, kind="stable")
        acts, ts, rg, ex, en = acts[order], ts[order], rg[order], ex[order], en[order]
        bounds = np.flatnonzero(np.diff(acts)) + 1
        for chunk in np.split(np.arange(len(acts)), bounds):
            if len(chunk):
                k = int(acts[chunk[0]])
                log[p_][k] = list(zip(ts[chunk].tolist(), rg[chunk].tolist(), ex[chunk].tolist(), en[chunk].tolist()))
    agree = collections.Counter()
    sel = who == 1
    same = R_["act"][sel] == R_["tch"][sel]
    agree["same"], agree["different"] = int(same.sum()), int((~same).sum())
    diff_pairs = np.stack([R_["tch"][sel][~same], R_["act"][sel][~same]], 1)
    if len(diff_pairs):
        u, c = np.unique(diff_pairs, axis=0, return_counts=True)
        for (t_, n_), v in zip(u.tolist(), c.tolist()):
            agree[(names[t_], names[n_])] = v
    d_apl, d_net = dps[:F], dps[F:]
    diff = d_net - d_apl
    print(f"APL on {build} {d_apl.mean():.1f}, policy on {net_build} {d_net.mean():.1f}, difference {diff.mean():+.2f} ± {diff.std() / np.sqrt(F):.2f}")
    total = agree["same"] + agree["different"]
    print(f"policy decisions matching the APL's choice: {agree['same'] / max(1, total):.1%}")
    print("most common disagreements (APL would -> policy did), per fight:")
    pairs = sorted(((k, v) for k, v in agree.items() if isinstance(k, tuple)), key=lambda kv: -kv[1])[:15]
    for (t_, n_), v in pairs:
        print(f"   {v / F:6.2f}  {t_:28s} -> {n_}")
    print(f"\n{'action':32s} {'APL/fight':>10s} {'policy/fight':>12s} {'APL exec':>9s} {'pol exec':>9s} {'APL rage':>9s} {'pol rage':>9s}")
    rows = []
    for k in sorted(set(log[0]) | set(log[1]), key=lambda k: -(len(log[1].get(k, [])) + len(log[0].get(k, [])))):
        r0, r1 = log[0].get(k, []), log[1].get(k, [])
        e0 = sum(1 for x in r0 if x[2]) / F
        e1 = sum(1 for x in r1 if x[2]) / F
        g0 = np.mean([x[1] for x in r0]) if r0 else float("nan")
        g1 = np.mean([x[1] for x in r1]) if r1 else float("nan")
        rows.append({"action": names[k], "apl": len(r0) / F, "policy": len(r1) / F, "apl_execute": e0, "policy_execute": e1,
                     "apl_rage": None if np.isnan(g0) else float(g0), "policy_rage": None if np.isnan(g1) else float(g1)})
        print(f"{names[k]:32s} {len(r0) / F:10.2f} {len(r1) / F:12.2f} {e0:9.2f} {e1:9.2f} {g0:9.1f} {g1:9.1f}")
    cooldowns = []
    print("\ncooldown timing (mean fight time of the 1st / 2nd / 3rd use, seconds):")
    for k in sorted(set(log[0]) | set(log[1])):
        nm = names[k]
        if not any(s in nm for s in ("Death Wish", "Recklessness", "Blood Fury", "Mighty Rage", "Sapper", "Bloodrage",
                                     "Berserker Rage", "Sweeping", "Demonic", "Battle Shout")):
            continue
        out = []
        for p in range(2):
            by_fight = collections.defaultdict(list)
            for x in log[p].get(k, []):
                by_fight[x[3]].append(x[0])
            uses = [[], [], []]
            for ts in by_fight.values():
                for n, t in enumerate(sorted(ts)[:3]):
                    uses[n].append(t)
            out.append(" / ".join(f"{np.mean(u):5.0f} ({len(u) / F:.0%})" if u else "    -      " for u in uses))
            if p == 0:
                cd_row = {"action": nm, "apl": [], "policy": []}
            cd_row["apl" if p == 0 else "policy"] = [{"mean_s": round(float(np.mean(u)), 1), "share": round(len(u) / F, 3)} for u in uses if u]
        cooldowns.append(cd_row)
        print(f"   {nm:32s} APL {out[0]}    policy {out[1]}")
    if a.json:
        dis = [{"apl_would": t_, "policy_did": n_, "per_fight": round(v / F, 3)} for (t_, n_), v in pairs]
        a.json.write_text(json.dumps({"fights": F, "apl_talents": build, "policy_talents": net_build, "apl": float(d_apl.mean()),
                                      "policy": float(d_net.mean()), "diff": float(diff.mean()), "diff_se": float(diff.std() / np.sqrt(F)),
                                      "agreement": agree["same"] / max(1, total), "disagreements": dis, "actions": rows,
                                      "cooldowns": cooldowns, "timeline": {"apl": timeline[0], "policy": timeline[1]}},
                                     indent=1, default=float))


if __name__ == "__main__":
    main()
