#!/usr/bin/env python3
"""Run reproducible, resumable native equipment searches across the roster."""

import argparse
import concurrent.futures
import hashlib
import json
import os
from pathlib import Path
import re
import subprocess


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--binary", required=True, type=Path)
    parser.add_argument("--roster", type=Path, default=Path("artifacts/forever_dps_5min.json"))
    parser.add_argument("--baseline-results", type=Path)
    parser.add_argument("--builds", help="comma-separated build keys")
    parser.add_argument("--output", required=True, type=Path)
    parser.add_argument("--representatives", action="store_true")
    available_cpus = len(os.sched_getaffinity(0)) if hasattr(os, "sched_getaffinity") else (os.cpu_count() or 1)
    parser.add_argument("--workers", type=int, default=available_cpus)
    parser.add_argument("--screen", type=int, default=100)
    parser.add_argument("--validate", type=int, default=1000)
    parser.add_argument("--passes", type=int, default=4)
    parser.add_argument("--iterations", type=int, default=5000)
    parser.add_argument("--seed", type=int, default=20291941)
    args = parser.parse_args()
    args.binary = args.binary.resolve()
    args.output.mkdir(parents=True, exist_ok=True)
    roster = json.loads(args.roster.read_text())["Results"]
    jobs = []
    seen = set()
    for row in roster:
        key = row["Key"] if args.representatives else (row["Key"], row["Race"])
        if key not in seen:
            jobs.append((row["Key"], row["Race"]))
            seen.add(key)
    expected = 23 if args.representatives else 147
    if len(jobs) != expected:
        raise SystemExit(f"Roster has {len(jobs)} jobs; expected {expected}.")
    if args.builds:
        wanted = set(args.builds.split(","))
        unknown = wanted - {build for build, _ in jobs}
        if unknown:
            raise SystemExit(f"Unknown builds: {sorted(unknown)}")
        jobs = [(build, race) for build, race in jobs if build in wanted]

    digest = hashlib.sha256()
    paths = {args.binary, Path("assets/database/db.json")}
    if args.baseline_results:
        args.baseline_results = args.baseline_results.resolve()
        paths.add(args.baseline_results)
    paths.update(Path("assets/db_inputs").glob("forever_*.json"))
    paths.update(p for p in Path("ui").rglob("*")
                 if (p.suffix == ".json" or p.name == "presets.ts")
                 and p.name != "forever_ranked_profiles.json")
    for path in sorted(paths):
        digest.update(str(path).encode())
        digest.update(path.read_bytes())
    settings = {key: getattr(args, key) for key in ("screen", "validate", "passes", "iterations", "seed")}
    manifest = {"inputsSHA256": digest.hexdigest(), "settings": settings}
    if args.baseline_results or args.builds:
        manifest["jobs"] = [list(job) for job in jobs]
    manifest_path = args.output / "manifest.json"
    if manifest_path.exists() and json.loads(manifest_path.read_text()) != manifest:
        raise SystemExit("Search inputs changed; choose a new output directory.")
    manifest_path.write_text(json.dumps(manifest, indent=2) + "\n")

    def search(job):
        build, race = job
        slug = build + "__" + re.sub(r"[^a-z0-9]+", "_", race.lower()).strip("_")
        prefix = args.output / slug
        result_file = prefix.with_suffix(".json")
        report_file = Path(str(prefix) + "." + build + "." + race.lower().replace(" ", "-") + ".gear-search.json")
        if not result_file.exists() or not report_file.exists():
            command = [
                str(args.binary), "-build", build, "-race", race, "-search-gear",
                "-gear-screen", str(args.screen), "-gear-validate", str(args.validate),
                "-gear-passes", str(args.passes), "-iterations", str(args.iterations),
                "-seed", str(args.seed), "-output", str(prefix),
            ]
            if args.baseline_results:
                command.extend(["-baseline-results", str(args.baseline_results)])
            with prefix.with_suffix(".log").open("w") as log:
                subprocess.run(command, stdout=log, stderr=subprocess.STDOUT,
                               env={**os.environ, "GOMAXPROCS": "1"}, check=True)
        report = json.loads(report_file.read_text())
        print(f"{build}/{race}: {report['Baseline']['DPS']:.2f} -> {report['Final']['DPS']:.2f}; "
              f"{report['Passes']} passes, converged={report['Converged']}", flush=True)
        return job, report["Baseline"], report["Final"], report["Converged"], str(report_file)

    completed, failures = [], []
    with concurrent.futures.ThreadPoolExecutor(max_workers=args.workers) as executor:
        futures = {executor.submit(search, job): job for job in jobs}
        for future in concurrent.futures.as_completed(futures):
            try:
                completed.append(future.result())
            except Exception as error:
                failures.append((futures[future], str(error)))
                print(f"FAILED {futures[future]}: {error}", flush=True)
    completed.sort(key=lambda value: jobs.index(value[0]))
    for name, index in (("baseline", 1), ("results", 2)):
        payload = {
            "Duration": 300, "StartingArmor": 3731, "Tier1Bonuses": True,
            "EquipmentScale": 1, "SearchManifest": manifest,
            "Results": [result[index] for result in completed],
        }
        (args.output / f"{name}.json").write_text(json.dumps(payload, indent=2) + "\n")
    summary = {
        "completed": len(completed), "requested": len(jobs), "failures": failures,
        "unconverged": [job for job, _, _, converged, _ in completed if not converged],
        "reports": [path for *_, path in completed],
    }
    (args.output / "summary.json").write_text(json.dumps(summary, indent=2) + "\n")
    if failures:
        raise SystemExit(1)


if __name__ == "__main__":
    main()
