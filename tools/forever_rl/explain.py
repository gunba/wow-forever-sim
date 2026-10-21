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
    ap.add_argument("--talents", default="published")
    ap.add_argument("--model", type=Path, required=True)
    ap.add_argument("--rotation", type=Path, help="compare against this rotation (APLRotation JSON) instead of the published APL")
    ap.add_argument("--fights", type=int, default=512)
    ap.add_argument("--threads", type=int, default=12)
    ap.add_argument("--json", type=Path)
    a = ap.parse_args()
    cfg = dict(build=a.build, race=a.race, fullChoices=True, pollSeconds=0.25, rewardScale=1e-5, threads=a.threads,
               labeledRands=True, seed=99, rotation=a.rotation.read_text() if a.rotation else "")
    F = a.fights
    ev = VecEnv(cfg, 2 * F)
    spec = ev.spec
    B = Builds(spec)
    build = B.published if a.talents == "published" else B.canon(a.talents)
    names = spec["actionNames"]
    feat = spec["featureNames"]
    fi = {k: feat.index(k) for k in ("elapsed", "rage", "execute20")}
    dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    ck = torch.load(a.model, map_location="cpu", weights_only=False)
    net = Net(spec["obsDim"], spec["actDim"], ck["net"]["trunk.0.weight"].shape[0]).to(dev)
    net.load_state_dict(ck["net"])
    net.eval()
    for i in range(2 * F):
        ev.set_talents(i, build)
    obs, mask, tch = ev.restart(np.tile(31_337 + np.arange(F, dtype=np.int64), 2))
    is_net = np.arange(2 * F) >= F
    dps = np.full(2 * F, np.nan)
    log = [collections.defaultdict(list), collections.defaultdict(list)]  # action -> [(time, rage, execute, env)]
    agree = collections.Counter()
    dur = spec["duration"] / 1e9 if spec["duration"] > 1e6 else spec["duration"]
    while np.isnan(dps).any():
        with torch.no_grad():
            logits, _ = net(torch.from_numpy(obs).to(dev), torch.from_numpy(mask).to(dev).bool())
        greedy = logits.argmax(-1).cpu().numpy().astype(np.int32)
        act = np.where(is_net, greedy, tch.astype(np.int32))
        live = np.isnan(dps)
        for i in np.nonzero(live)[0]:
            p = int(is_net[i])
            t = float(obs[i, fi["elapsed"]]) * dur
            log[p][int(act[i])].append((t, float(obs[i, fi["rage"]]) * 100, bool(obs[i, fi["execute20"]] > 0.5), int(i)))
            if p == 1:
                agree["same" if act[i] == tch[i] else "different"] += 1
                if act[i] != tch[i]:
                    agree[(names[int(tch[i])], names[int(act[i])])] += 1
        obs, mask, tch, _, done, info = ev.step(act)
        new = (done == 1) & np.isnan(dps)
        dps[new] = info[new, 0]
    ev.close()
    d_apl, d_net = dps[:F], dps[F:]
    diff = d_net - d_apl
    print(f"{B.describe(build)}: APL {d_apl.mean():.1f}, policy {d_net.mean():.1f}, difference {diff.mean():+.2f} ± {diff.std() / np.sqrt(F):.2f}")
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
                     "apl_rage": g0, "policy_rage": g1})
        print(f"{names[k]:32s} {len(r0) / F:10.2f} {len(r1) / F:12.2f} {e0:9.2f} {e1:9.2f} {g0:9.1f} {g1:9.1f}")
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
        print(f"   {nm:32s} APL {out[0]}    policy {out[1]}")
    if a.json:
        a.json.write_text(json.dumps({"apl": float(d_apl.mean()), "policy": float(d_net.mean()), "diff": float(diff.mean()),
                                      "diff_se": float(diff.std() / np.sqrt(F)), "actions": rows}, indent=1, default=float))


if __name__ == "__main__":
    main()
