#!/usr/bin/env python3
"""Paired talent screen: every legal one-point move from a base build, played by the published APL (or another
rotation, or a trained network), on shared seeds with one random stream per roll label so fights of two builds stay
aligned.

Each round screens all moves on --fights fights, confirms the leaders on --confirm-fights fresh fights and moves the base
to the best confirmed improvement (more than --z standard errors). Results go to <out>/screen.jsonl.

    python3 tools/forever_rl/talent_screen.py --build fury --race Orc --out ~/.cache/forever-sim-audit/rl/runs/fury_screen
"""
import argparse
import json
import math
import time
from pathlib import Path

import numpy as np

from forever_env import VecEnv
from codesign import Builds


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--build", default="fury")
    ap.add_argument("--race", default="Orc")
    ap.add_argument("--base", default="", help="starting build (default: the published one)")
    ap.add_argument("--fights", type=int, default=512)
    ap.add_argument("--confirm-fights", type=int, default=4096)
    ap.add_argument("--confirm-top", type=int, default=24)
    ap.add_argument("--rounds", type=int, default=3)
    ap.add_argument("--z", type=float, default=2.5)
    ap.add_argument("--threads", type=int, default=12)
    ap.add_argument("--chunk", type=int, default=8192, help="environments per batch")
    ap.add_argument("--model", type=Path, help="play with this codesign/train model (greedy) instead of the published APL")
    ap.add_argument("--rotation", type=Path, help="play with this rotation (APLRotation JSON) instead of the published APL")
    ap.add_argument("--seed0", type=int, default=1_000_001)
    ap.add_argument("--out", type=Path, required=True)
    a = ap.parse_args()
    a.out.mkdir(parents=True, exist_ok=True)
    cfg = dict(build=a.build, race=a.race, fullChoices=True, pollSeconds=0.25, rewardScale=1e-5, threads=a.threads,
               labeledRands=True, seed=77, rotation=a.rotation.read_text() if a.rotation else "")
    ev = VecEnv(cfg, a.chunk)
    spec = ev.spec
    B = Builds(spec)
    tal = B.tal
    net = None
    if a.model:
        import torch
        from train import Net
        ck = torch.load(a.model, map_location="cpu", weights_only=False)
        if ck["spec"]["actDim"] != spec["actDim"] or ck["spec"]["obsDim"] != spec["obsDim"]:
            raise SystemExit("model spaces differ from this environment")
        dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")
        hidden = ck["net"]["trunk.0.weight"].shape[0]
        net = Net(spec["obsDim"], spec["actDim"], hidden).to(dev)
        net.load_state_dict(ck["net"])
        net.eval()
    log = (a.out / "screen.jsonl").open("a")

    def play(builds, fights, seed0):
        """DPS per fight for each build, fight f on seed seed0 + f for every build."""
        per = max(1, a.chunk // fights)
        out = np.zeros((len(builds), fights))
        for c in range(0, len(builds), per):
            group = builds[c:c + per]
            n = len(group) * fights
            for j, b in enumerate(group):
                for f in range(fights):
                    ev.set_talents(j * fights + f, b)
            seeds = np.zeros(a.chunk, np.int64)
            seeds[:n] = np.tile(seed0 + np.arange(fights, dtype=np.int64), len(group))
            seeds[n:] = 1  # idle environments
            obs, mask, tch = ev.restart(seeds)
            got = np.full(a.chunk, np.nan)
            got[n:] = 0
            while np.isnan(got).any():
                if net is None:
                    act = tch.astype(np.int32)
                else:
                    import torch
                    with torch.no_grad():
                        logits, _ = net(torch.from_numpy(obs).to(next(net.parameters()).device),
                                        torch.from_numpy(mask).to(next(net.parameters()).device).bool())
                    act = logits.argmax(-1).cpu().numpy().astype(np.int32)
                obs, mask, tch, _, done, info = ev.step(act)
                new = (done == 1) & np.isnan(got)
                got[new] = info[new, 0]
            out[c:c + len(group)] = got[:n].reshape(len(group), fights)
        return out

    def moves(base):
        p = tal.decode(base)
        res = []
        for i in np.nonzero(p)[0]:
            q = p.copy()
            q[i] -= 1
            for j in range(tal.n):
                if j == i:
                    continue
                r = q.copy()
                r[j] += 1
                if r[j] <= tal.max[j] and tal.valid(r):
                    res.append((f"{tal.names[i]} -> {tal.names[j]}", tal.encode(r)))
        return res

    base = B.canon(a.base) if a.base else B.published
    seed = a.seed0
    for rnd in range(1, a.rounds + 1):
        t0 = time.time()
        mv = moves(base)
        builds = [base] + [b for _, b in mv]
        d = play(builds, a.fights, seed)
        seed += a.fights
        diff = d[1:] - d[0]
        mean, se = diff.mean(1), diff.std(1) / math.sqrt(a.fights)
        rows = []
        for k, (name, b) in enumerate(mv):
            rows.append({"round": rnd, "stage": "screen", "move": name, "talents": b, "diff": round(float(mean[k]), 2),
                         "se": round(float(se[k]), 2)})
        for r in rows:
            log.write(json.dumps(r) + "\n")
        log.flush()
        print(f"round {rnd}: base {base} {d[0].mean():.1f}; {len(mv)} moves in {time.time() - t0:.0f}s", flush=True)
        top = sorted(rows, key=lambda r: -r["diff"])[:a.confirm_top]
        d2 = play([base] + [r["talents"] for r in top], a.confirm_fights, seed)
        seed += a.confirm_fights
        diff2 = d2[1:] - d2[0]
        best = None
        for k, r in enumerate(top):
            r2 = {"round": rnd, "stage": "confirm", "move": r["move"], "talents": r["talents"],
                  "diff": round(float(diff2[k].mean()), 2), "se": round(float(diff2[k].std() / math.sqrt(a.confirm_fights)), 2),
                  "screen_diff": r["diff"], "base": base, "base_dps": round(float(d2[0].mean()), 1)}
            log.write(json.dumps(r2) + "\n")
            print(f"   {r2['diff']:+6.2f} ± {r2['se']:.2f}  (screen {r['diff']:+6.2f})  {r['move']:55s} {r['talents']}", flush=True)
            if r2["diff"] > a.z * r2["se"] and (best is None or r2["diff"] > best["diff"]):
                best = r2
        log.flush()
        if best is None:
            print(f"round {rnd}: no move beats {base} by {a.z} standard errors", flush=True)
            break
        print(f"round {rnd}: take {best['move']} ({best['diff']:+.2f} ± {best['se']:.2f}) -> {best['talents']}", flush=True)
        base = best["talents"]
    (a.out / "result.json").write_text(json.dumps({"final": base, "published": B.published}, indent=1))
    ev.close()


if __name__ == "__main__":
    main()
