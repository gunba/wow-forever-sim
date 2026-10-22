#!/usr/bin/env python3
"""Talents and rotation learned for every DPS build at its highest-simming race (codesign.py), with the action report
(explain.py) the site shows beside the published results. Resumable: finished builds are skipped.

Per build: co-design (budget by decisions per fight), then a paired 4,096-fight confirmation inside codesign.py, then
1,024 fights of the network against the published APL for the actions. Writes artifacts/ml/<key>.json and
artifacts/ml/index.json; working files go to --runs.

    python3 tools/forever_rl/run_all.py --runs ~/.cache/forever-sim-audit/rl/all
"""
import argparse
import json
import subprocess
import sys
import time
from pathlib import Path

HERE = Path(__file__).resolve().parent
REPO = HERE.parents[1]
TANKS = {"tank_warrior", "protection_paladin", "feral_tank_druid"}

# Budgets by fight length in decisions (melee and weaving builds decide far more often than casters).
BUDGETS = {
    "melee": ["--max-decisions", "500e6", "--min-decisions", "250e6", "--patience", "120e6", "--minutes", "24",
              "--eval-every", "30e6"],
    "caster": ["--max-decisions", "150e6", "--min-decisions", "60e6", "--patience", "45e6", "--minutes", "8",
               "--eval-every", "15e6", "--rollout", "64", "--lam", "0.95", "--bc-decay", "30e6"],
}


def best_races(results):
    best = {}
    for r in results["Results"]:
        if r["Key"] in TANKS:
            continue
        if r["Key"] not in best or r["DPS"] > best[r["Key"]]["DPS"]:
            best[r["Key"]] = r
    return best


def talent_changes(spec, published, learned):
    sys.path.insert(0, str(HERE))
    from codesign import Talents
    tal = Talents(spec)
    a, b = tal.decode(published), tal.decode(learned)
    out = []
    for i in range(tal.n):
        if a[i] != b[i]:
            out.append({"talent": tal.names[i], "tree": spec["treeNames"][tal.tree[i]], "published": int(a[i]), "learned": int(b[i])})
    return out


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--runs", type=Path, required=True)
    ap.add_argument("--out", type=Path, default=REPO / "artifacts/ml")
    ap.add_argument("--only", default="", help="comma-separated build keys")
    ap.add_argument("--threads", type=int, default=24)
    ap.add_argument("--decisions", type=Path, default=Path.home() / ".cache/forever-sim-audit/rl/preflight.json",
                    help="decisions per fight per build (melee budget from 1,000)")
    a = ap.parse_args()
    a.runs.mkdir(parents=True, exist_ok=True)
    a.out.mkdir(parents=True, exist_ok=True)
    results = json.load(open(REPO / "artifacts/modelled_gear/forever_dps_5min.json"))
    best = best_races(results)
    # Decisions per fight with full choices, measured by the pre-flight (rlbench -full -poll 0).
    speeds = json.load(open(a.decisions)) if a.decisions.exists() else {}
    keys = sorted(best, key=lambda k: -best[k]["DPS"])
    if a.only:
        keys = [k for k in keys if k in a.only.split(",")]
    progress = (a.runs / "progress.jsonl").open("a")
    # A second pass retries any build whose first attempt failed (e.g. killed for memory), from scratch.
    for key in keys + [None] + keys:
        if key is None:
            for k in keys:
                run = a.runs / k
                if not (a.out / f"{k}.json").exists() and run.exists() and not (run / "summary.json").exists():
                    import shutil
                    shutil.rmtree(run)
            continue
        row = best[key]
        report = a.out / f"{key}.json"
        if report.exists():
            continue
        run = a.runs / key
        kind = "melee" if speeds.get(key, 0) >= 1000 else "caster"
        t0 = time.time()
        budgets = dict(BUDGETS)
        if (a.runs / "budgets.json").exists():
            # Overrides read before every build, so a running queue can be retuned.
            budgets.update(json.load(open(a.runs / "budgets.json")))
        cmd = [sys.executable, str(HERE / "codesign.py"), "--build", key, "--race", row["Race"], "--threads", str(a.threads),
               "--out", str(run), *budgets[kind]]
        print(f"{time.strftime('%H:%M:%S')} {key} ({row['Race']}, {kind})", flush=True)
        rc = 0
        if not (run / "summary.json").exists():
            with open(run.with_suffix(".out"), "w") as fh:
                rc = subprocess.run(cmd, stdout=fh, stderr=subprocess.STDOUT, cwd=REPO).returncode
        if rc != 0 or not (run / "summary.json").exists():
            progress.write(json.dumps({"key": key, "status": "failed", "rc": rc, "minutes": round((time.time() - t0) / 60, 1)}) + "\n")
            progress.flush()
            continue
        summary = json.load(open(run / "summary.json"))
        spec = json.load(open(run / "spec.json"))
        ex_json = run / "explain.json"
        cmd = [sys.executable, str(HERE / "explain.py"), "--build", key, "--race", row["Race"], "--model", str(run / "model.pt"),
               "--network-talents", summary["best_talents"], "--fights", "1024", "--threads", str(a.threads), "--json", str(ex_json)]
        with open(run / "explain.out", "w") as fh:
            rc = subprocess.run(cmd, stdout=fh, stderr=subprocess.STDOUT, cwd=REPO).returncode
        ex = json.load(open(ex_json)) if rc == 0 and ex_json.exists() else None
        conf = summary["confirm"]
        net_pub = next(r for r in conf["rows"] if r["talents"] == summary["published_talents"])
        net_best = next((r for r in conf["rows"] if r["talents"] == summary["best_talents"]), net_pub)
        out = {
            "key": key, "build": row.get("Build", key), "race": row["Race"], "class": spec.get("class", ""),
            "published": {"talents": summary["published_talents"], "dps": conf["apl"], "se": conf["apl_se"],
                          "benchmark_dps": row["DPS"]},
            "learned": {"talents": summary["best_talents"], "dps": net_best["dps"], "vs_published": net_best["vs_apl"],
                        "vs_published_se": net_best["vs_apl_se"],
                        "published_talents_dps": net_pub["dps"], "published_talents_vs_published": net_pub["vs_apl"],
                        "published_talents_vs_published_se": net_pub["vs_apl_se"]},
            "talent_changes": talent_changes(spec, summary["published_talents"], summary["best_talents"]),
            "fights": conf["fights"],
            "training": {k: summary[k] for k in ("stop", "train_minutes", "minutes", "decisions", "fights", "builds_seen",
                                                  "decisions_per_fight", "best_decisions")},
            "actions": ex,
        }
        report.write_text(json.dumps(out, indent=1) + "\n")
        progress.write(json.dumps({"key": key, "status": "done", "minutes": round((time.time() - t0) / 60, 1),
                                   "vs_published": round(net_best["vs_apl"], 2)}) + "\n")
        progress.flush()
        print(f"   {key}: {conf['apl']:.1f} -> {net_best['dps']:.1f} ({net_best['vs_apl']:+.2f} ± {net_best['vs_apl_se']:.2f})"
              f" in {(time.time() - t0) / 60:.1f} min", flush=True)
    index = []
    for p in sorted(a.out.glob("*.json")):
        if p.name == "index.json":
            continue
        r = json.load(open(p))
        index.append({"key": r["key"], "race": r["race"], "published_dps": r["published"]["dps"], "learned_dps": r["learned"]["dps"],
                      "vs_published": r["learned"]["vs_published"], "vs_published_se": r["learned"]["vs_published_se"],
                      "talents": r["learned"]["talents"], "talent_changes": len(r["talent_changes"])})
    (a.out / "index.json").write_text(json.dumps({"builds": index}, indent=1) + "\n")


if __name__ == "__main__":
    main()
