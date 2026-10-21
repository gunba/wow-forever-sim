#!/usr/bin/env python3
"""Paired comparison of builds and players: every (build, player) pair plays the same fights (fight f uses seed f, one
random stream per roll label), and each row reports its difference to the first row with a paired standard error.

A player is "apl" (the published rotation), "apl=FILE" (another rotation, APLRotation JSON) or the path of a model
saved by codesign.py or train.py (greedy).

    python3 tools/forever_rl/compare.py --build fury --race Orc --fights 4096 \\
        --row published:apl --row published:runs/fury_codesign/model.pt --row 30305213002-25053005122010501:apl
"""
import argparse
import json
import math
from pathlib import Path

import numpy as np
import torch

from codesign import Builds
from forever_env import VecEnv
from train import Net


def play_rows(ev, rows, idxs, dps, F, per, nets, a, dev):
    """Fill dps[j] for the rows idxs, per fights at a time, on the shared seeds."""
    for c in range(0, len(idxs), per):
        group = idxs[c:c + per]
        n = len(group) * F
        for k, j in enumerate(group):
            for f in range(F):
                ev.set_talents(k * F + f, rows[j][0])
        seeds = np.ones(ev.n, np.int64)
        seeds[:n] = np.tile(a.seed0 + np.arange(F, dtype=np.int64), len(group))
        obs, mask, tch = ev.restart(seeds)
        got = np.full(ev.n, np.nan)
        got[n:] = 0
        player_of = np.array([rows[j][1] for j in group for _ in range(F)] + ["idle"] * (ev.n - n))
        while np.isnan(got).any():
            act = tch.astype(np.int32).copy()
            for p, net in nets.items():
                sel = np.nonzero(player_of == p)[0]
                if len(sel):
                    with torch.no_grad():
                        logits, _ = net(torch.from_numpy(obs[sel]).to(dev), torch.from_numpy(mask[sel]).to(dev).bool())
                    act[sel] = logits.argmax(-1).cpu().numpy().astype(np.int32)
            obs, mask, tch, _, done, info = ev.step(act)
            new = (done == 1) & np.isnan(got)
            got[new] = info[new, 0]
        for k, j in enumerate(group):
            dps[j] = got[k * F:(k + 1) * F]


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--build", default="fury")
    ap.add_argument("--race", default="Orc")
    ap.add_argument("--row", action="append", required=True,
                    help="BUILD:PLAYER; BUILD is \"published\" or a talent string")
    ap.add_argument("--fights", type=int, default=2048)
    ap.add_argument("--seed0", type=int, default=31_337)
    ap.add_argument("--threads", type=int, default=12)
    ap.add_argument("--chunk", type=int, default=4096)
    ap.add_argument("--json", type=Path, help="also write the table here")
    a = ap.parse_args()
    cfg = dict(build=a.build, race=a.race, fullChoices=True, pollSeconds=0.25, rewardScale=1e-5, threads=a.threads,
               labeledRands=True, seed=99)
    ev = VecEnv(cfg, 1)
    spec = ev.spec
    ev.close()
    B = Builds(spec)
    dev = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    nets = {}
    rows = []
    for r in a.row:
        build, _, player = r.rpartition(":")
        build = B.published if build == "published" else B.canon(build)
        if not player.startswith("apl") and player not in nets:
            ck = torch.load(player, map_location="cpu", weights_only=False)
            if ck["spec"]["obsDim"] != spec["obsDim"] or ck["spec"]["actDim"] != spec["actDim"]:
                raise SystemExit(f"{player}: model spaces differ from this environment")
            net = Net(spec["obsDim"], spec["actDim"], ck["net"]["trunk.0.weight"].shape[0]).to(dev)
            net.load_state_dict(ck["net"])
            net.eval()
            nets[player] = net
        rows.append((build, player))

    F = a.fights
    per = max(1, a.chunk // F)
    dps = np.zeros((len(rows), F))
    # Rows played by another rotation need an environment built on it; models play on the published one.
    by_rot = {}
    for j, (b, p) in enumerate(rows):
        by_rot.setdefault(p.split("=", 1)[1] if p.startswith("apl=") else "", []).append(j)
    for rot, idxs in by_rot.items():
        ev = VecEnv({**cfg, "rotation": Path(rot).read_text()} if rot else cfg, a.chunk)
        play_rows(ev, rows, idxs, dps, F, per, nets, a, dev)
        ev.close()
    out = []

    for j, (b, p) in enumerate(rows):
        d = dps[j] - dps[0]
        out.append({"build": b, "described": B.describe(b), "player": p, "dps": round(float(dps[j].mean()), 2),
                    "se": round(float(dps[j].std() / math.sqrt(F)), 2), "vs_first": round(float(d.mean()), 2),
                    "vs_first_se": round(float(d.std() / math.sqrt(F)), 2)})
        print(f"{out[-1]['dps']:8.1f}  {out[-1]['vs_first']:+7.2f} ± {out[-1]['vs_first_se']:.2f}  "
              f"{p if p.startswith('apl') else Path(p).parent.name + '/' + Path(p).name:34s} {out[-1]['described']}", flush=True)
    if a.json:
        a.json.write_text(json.dumps({"fights": F, "rows": out}, indent=1))


if __name__ == "__main__":
    main()
