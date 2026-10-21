#!/usr/bin/env python3
"""Training-time estimates per build from measured environment data.

Input: one JSON line per build from `rlbench -json` (teacher replay, fixed talents): fights per second, decisions per
fight, fight-to-fight DPS spread. Constants below come from the Fire pilots on this machine (RTX 4070 Ti, 12 of 24
threads); replace them when a pilot is rerun.

    python3 tools/forever_rl/estimate.py ~/.cache/forever-sim-audit/rl/est/builds.jsonl
"""
import argparse
import json
import math

# Measured (Fire pilots, October 2026).
LEARNER_DECISIONS_PER_S = 350_000   # PPO loop, 2048 envs, 12 threads: 317k-410k across Fire, Fury, Enhancement, Balance, BM, Frostfire
EVAL_DECISIONS_PER_S = 500_000      # inference-only rollouts (no training), Fire: (t_env + t_inf) per 2048x64 decisions
BC_DECISIONS = 4e6                  # behaviour cloning to 99.8% agreement and the published DPS (Fire: 30 iterations)
ROTATION_DECISIONS = 86e6           # fixed-talent PPO, Fire: 947 (cloned) -> 954 DPS, most of it by ~30M decisions
FIRE_DECISIONS_PER_FIGHT = 195
# Assumptions for the two-stage talent search (not yet measured).
TALENT_CONDITIONED_FACTOR = 3       # a rotation network that plays many builds needs ~3x the single-build budget
CANDIDATE_BUILDS = 1000             # builds the search evaluates
SE_TARGET_DPS = 1.0                 # standard error of each build's mean DPS; shared seeds roughly halve the fights needed


def minutes(decisions, rate):
    return decisions / rate / 60


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("builds", help="rlbench -json output, one line per build")
    ap.add_argument("--markdown", action="store_true")
    a = ap.parse_args()
    rows = [json.loads(l) for l in open(a.builds) if l.strip()]
    out = []
    for r in rows:
        if "error" in r:
            continue
        dpf = r["decisionsPerFight"]
        scale = dpf / FIRE_DECISIONS_PER_FIGHT
        bc = minutes(BC_DECISIONS * max(1, scale), LEARNER_DECISIONS_PER_S)
        rot_lo = minutes(ROTATION_DECISIONS, LEARNER_DECISIONS_PER_S)
        rot_hi = minutes(ROTATION_DECISIONS * max(1, scale), LEARNER_DECISIONS_PER_S)
        fights_per_build = math.ceil((r["stdDPS"] / SE_TARGET_DPS) ** 2 / 2)
        search = CANDIDATE_BUILDS * fights_per_build * dpf / EVAL_DECISIONS_PER_S / 60
        tal_lo = TALENT_CONDITIONED_FACTOR * rot_lo + search
        tal_hi = TALENT_CONDITIONED_FACTOR * rot_hi + search
        out.append(dict(build=r["build"], dpf=dpf, sd=r["stdDPS"], bc=bc, rot_lo=rot_lo, rot_hi=rot_hi,
                        fights=fights_per_build, search=search, tal_lo=tal_lo, tal_hi=tal_hi))
    out.sort(key=lambda x: x["dpf"])
    if a.markdown:
        print("| Build | Decisions/fight | Warm start | Rotation RL | Talents + rotation |")
        print("|---|---|---|---|---|")
        for x in out:
            print(f"| {x['build']} | {x['dpf']:.0f} | {x['bc']:.1f} min | {x['rot_lo']:.0f}-{x['rot_hi']:.0f} min | "
                  f"{x['tal_lo'] / 60:.1f}-{x['tal_hi'] / 60:.1f} h |")
    else:
        for x in out:
            print(f"{x['build']:22s} dpf {x['dpf']:6.0f} sd {x['sd']:5.1f}  warm {x['bc']:5.1f} min  rotation {x['rot_lo']:4.0f}-{x['rot_hi']:5.0f} min"
                  f"  search {x['search']:6.0f} min ({x['fights']} fights/build)  talents+rotation {x['tal_lo'] / 60:5.1f}-{x['tal_hi'] / 60:5.1f} h")
    tot = lambda k: sum(x[k] for x in out) / 60
    print(f"\nall {len(out)} builds: warm start {tot('bc'):.1f} h, rotation {tot('rot_lo'):.1f}-{tot('rot_hi'):.1f} h, "
          f"talents+rotation {tot('tal_lo'):.0f}-{tot('tal_hi'):.0f} h (sequential, one machine)")


if __name__ == "__main__":
    main()
