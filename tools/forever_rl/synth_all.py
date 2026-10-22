#!/usr/bin/env python3
"""Simulatable learned profiles: for every build with a learned result (run_all.py), turn its policy into APL overrides
(apl_synth.py) and assemble a profile: the published player for that race with the learned talents and the
converted APL. Resumable: builds with a profile are skipped.

    python3 tools/forever_rl/synth_all.py --runs ~/.cache/forever-sim-audit/rl/all
"""
import argparse
import json
import subprocess
import sys
import time
from copy import deepcopy
from pathlib import Path

HERE = Path(__file__).resolve().parent
REPO = HERE.parents[1]


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--runs", type=Path, required=True, help="run_all.py's working folder")
    ap.add_argument("--ml", type=Path, default=REPO / "artifacts/ml")
    ap.add_argument("--profiles", type=Path, default=REPO / "artifacts/modelled_gear/forever_input_profiles.json")
    ap.add_argument("--only", default="")
    ap.add_argument("--threads", type=int, default=24)
    ap.add_argument("--fights", type=int, default=256)
    ap.add_argument("--rounds", type=int, default=3)
    a = ap.parse_args()
    rows = {(r["Key"], r["Race"]): r for r in json.load(open(a.profiles))["Results"]}
    out_rows = []
    log = (a.runs / "synth_progress.jsonl").open("a")
    for report in sorted(a.ml.glob("*.json")):
        if report.name == "index.json":
            continue
        r = json.load(open(report))
        key, race = r["key"], r["race"]
        if a.only and key not in a.only.split(","):
            continue
        run = a.runs / key
        synth = run / "synth"
        profile_path = synth / "profile.json"
        if not profile_path.exists():
            t0 = time.time()
            cmd = [sys.executable, str(HERE / "apl_synth.py"), "--build", key, "--race", race, "--model", str(run / "model.pt"),
                   "--talents", r["learned"]["talents"], "--fights", str(a.fights), "--rounds", str(a.rounds),
                   "--threads", str(a.threads), "--out", str(synth)]
            print(f"{time.strftime('%H:%M:%S')} {key} ({race})", flush=True)
            synth.mkdir(parents=True, exist_ok=True)
            with open(synth / "apl_synth.out", "w") as fh:
                rc = subprocess.run(cmd, stdout=fh, stderr=subprocess.STDOUT, cwd=REPO).returncode
            if rc != 0 or not (synth / "apl.json").exists():
                log.write(json.dumps({"key": key, "status": "failed", "rc": rc}) + "\n"); log.flush()
                continue
            res = json.load(open(synth / "report.json"))
            player = deepcopy(rows[(key, race)]["BaselinePlayer"])
            player["talentsString"] = r["learned"]["talents"]
            player["rotation"] = json.load(open(synth / "apl.json"))
            profile_path.write_text(json.dumps({"Key": key, "Race": race, "BaselinePlayer": player}, indent=1) + "\n")
            log.write(json.dumps({"key": key, "status": "done", "minutes": round((time.time() - t0) / 60, 1),
                                  "synthesised_vs_published": res["synthesised_vs_published"],
                                  "network_vs_published": res["network_vs_published"], "entries": res["entries"]}) + "\n")
            log.flush()
            print(f"   {key}: APL {res['synthesised_vs_published'][0]:+.1f} ± {res['synthesised_vs_published'][1]:.1f}, "
                  f"network {res['network_vs_published'][0]:+.1f} ({res['entries']} overrides, "
                  f"{(time.time() - t0) / 60:.1f} min)", flush=True)
        out_rows.append(json.load(open(profile_path)))
    # All learned players in the frozen-profile format, for the bench (one row per build, at its best race).
    (a.runs / "learned_profiles.json").write_text(json.dumps({"GearScenario": json.load(open(a.profiles)).get("GearScenario"),
                                                              "Results": out_rows}, indent=1) + "\n")


if __name__ == "__main__":
    main()
