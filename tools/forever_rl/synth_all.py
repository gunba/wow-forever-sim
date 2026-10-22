#!/usr/bin/env python3
"""Simulatable learned profiles: for every build with a learned result (run_all.py), turn its policy into edits to the
published APL (apl_synth.py) and assemble a profile: the published player for that race with the chosen talents and
the converted APL. Builds whose learned talents differ from the published ones are converted on both, and the better
profile against the published setup is kept (the APL may not use the learned talents as the network does).
Resumable: builds with a profile are skipped.

    python3 tools/forever_rl/synth_all.py --runs ~/.cache/forever-sim-audit/rl/all --jobs 3
"""
import argparse
import json
import shutil
import subprocess
import sys
import threading
import time
from concurrent.futures import ThreadPoolExecutor
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
    ap.add_argument("--jobs", type=int, default=3, help="conversions at once")
    ap.add_argument("--threads", type=int, default=24, help="threads shared by the jobs")
    ap.add_argument("--iterations", type=int, default=3)
    a = ap.parse_args()
    rows = {(r["Key"], r["Race"]): r for r in json.load(open(a.profiles))["Results"]}
    lock = threading.Lock()
    log = (a.runs / "synth_progress.jsonl").open("a")
    reports = {}
    for path in sorted(a.ml.glob("*.json")):
        r = json.load(open(path))
        if "key" in r and (not a.only or r["key"] in a.only.split(",")):
            reports[r["key"]] = r

    def say(text):
        with lock:
            print(f"{time.strftime('%H:%M:%S')} {text}", flush=True)

    def convert(key, tag, talents):
        r = reports[key]
        d = a.runs / key / "synth" / tag
        d.mkdir(parents=True, exist_ok=True)
        cmd = [sys.executable, str(HERE / "apl_synth.py"), "--build", key, "--race", r["race"],
               "--model", str(a.runs / key / "model.pt"), "--talents", talents, "--iterations", str(a.iterations),
               "--threads", str(max(1, a.threads // a.jobs)), "--out", str(d)]
        t0 = time.time()
        with open(d / "apl_synth.out", "w") as fh:
            rc = subprocess.run(cmd, stdout=fh, stderr=subprocess.STDOUT, cwd=REPO).returncode
        if rc != 0 or not (d / "report.json").exists():
            say(f"{key} [{tag} talents] failed ({rc})")
            return None
        rep = json.load(open(d / "report.json"))
        say(f"{key} [{tag} talents]: APL {rep['synthesised_vs_reference'][0]:+.1f} ± {rep['synthesised_vs_reference'][1]:.1f}, "
            f"network {rep['network_vs_reference'][0]:+.1f} against the published setup ({len(rep['added'])} added, "
            f"{len(rep['removed'])} removed, {(time.time() - t0) / 60:.1f} min)")
        return rep

    tasks = []
    for key, r in reports.items():
        if (a.runs / key / "synth" / "profile.json").exists():
            continue
        published = rows[(key, r["race"])]["BaselinePlayer"]
        tasks.append((key, "learned", r["learned"]["talents"]))
        if r["learned"]["talents"] != published["talentsString"]:
            tasks.append((key, "published", published["talentsString"]))
    def finish(key, got):
        """Keep the better candidate as the build's profile, once all its candidates are in."""
        r = reports[key]
        synth = a.runs / key / "synth"
        got = {t: rep for t, rep in got.items() if rep is not None}
        if not got:
            log.write(json.dumps({"key": key, "status": "failed"}) + "\n")
            log.flush()
            return
        tag = max(got, key=lambda t: got[t]["synthesised_vs_reference"][0])
        for name in ("apl.json", "rules.txt", "report.json"):
            shutil.copyfile(synth / tag / name, synth / name)
        player = deepcopy(rows[(key, r["race"])]["BaselinePlayer"])
        player["talentsString"] = got[tag]["talents"]
        player["rotation"] = json.load(open(synth / "apl.json"))
        (synth / "profile.json").write_text(json.dumps({"Key": key, "Race": r["race"], "TalentSource": tag,
                                                        "BaselinePlayer": player}, indent=1) + "\n")
        log.write(json.dumps({"key": key, "status": "done", "talents": tag,
                              "vs_published": got[tag]["synthesised_vs_reference"],
                              "network_vs_published": got[tag]["network_vs_reference"]}) + "\n")
        log.flush()
        say(f"{key}: kept the {tag} talents")

    pending = {}
    for key, tag, _ in tasks:
        pending.setdefault(key, set()).add(tag)
    results = {key: {} for key in pending}

    def run(task):
        key, tag, _ = task
        rep = convert(*task)
        with lock:
            results[key][tag] = rep
            pending[key].discard(tag)
            last = not pending[key]
        if last:
            finish(key, results[key])

    with ThreadPoolExecutor(a.jobs) as pool:
        list(pool.map(run, tasks))
    # Every learned player in the frozen-profile format, for the bench (one row per build, at its best race).
    out_rows = []
    for key in sorted(reports):
        path = a.runs / key / "synth" / "profile.json"
        if path.exists():
            row = json.load(open(path))
            out_rows.append({k: row[k] for k in ("Key", "Race", "BaselinePlayer")})
    (a.runs / "learned_profiles.json").write_text(json.dumps({"GearScenario": json.load(open(a.profiles)).get("GearScenario"),
                                                              "Results": out_rows}, indent=1) + "\n")


if __name__ == "__main__":
    main()
