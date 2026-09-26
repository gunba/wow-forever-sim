#!/usr/bin/env python3
"""Replay saved loadouts and paired sensitivities across available CPU cores."""

import argparse
import concurrent.futures
import csv
import hashlib
import json
import os
from pathlib import Path
import re
import subprocess

from build_display import expected_roster


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--binary", required=True, type=Path)
    parser.add_argument("--profiles", required=True, type=Path)
    parser.add_argument("--original-baselines", type=Path)
    parser.add_argument("--natural-hit", action="store_true",
                        help="replay the modeled equipment-only hit scenario without the paid exchange")
    parser.add_argument("--output", required=True, type=Path)
    cpus = len(os.sched_getaffinity(0)) if hasattr(os, "sched_getaffinity") else (os.cpu_count() or 1)
    parser.add_argument("--workers", type=int, default=cpus)
    parser.add_argument("--iterations", type=int, default=5000)
    parser.add_argument("--seed", type=int, default=20291951)
    args = parser.parse_args()
    args.binary = args.binary.resolve()
    args.profiles = args.profiles.resolve()
    args.output.mkdir(parents=True, exist_ok=True)
    profile_data = json.loads(args.profiles.read_text())
    if profile_data.get("GearScenario") == "modeled-65-v2" and not args.natural_hit:
        raise SystemExit("The v2 modeled benchmark requires --natural-hit")
    roster = [(row["Key"], row["Race"]) for row in profile_data["Results"]]
    if set(roster) != expected_roster() or len(roster) != len(set(roster)):
        raise SystemExit("Expected the complete current race/build roster.")

    scenarios = {
        "main": (args.profiles, ["-refresh-enchants"]),
        "tier1_off": (args.profiles, ["-refresh-enchants", "-tier1=false"]),
        "gear_110": (args.profiles, ["-refresh-enchants", "-equipment-scale", "1.1"]),
        "gear_150": (args.profiles, ["-refresh-enchants", "-equipment-scale", "1.5"]),
    }
    if args.original_baselines:
        args.original_baselines = args.original_baselines.resolve()
        scenarios["original"] = (args.original_baselines, [])
    paths = {args.binary, args.profiles, Path("assets/database/db.json")}
    paths.update(Path("artifacts/tanks").glob("*_1t_20261993_selected.json"))
    if args.original_baselines:
        paths.add(args.original_baselines)
    paths.update(Path("assets/db_inputs").glob("forever_*.json"))
    paths.update(p for p in Path("ui").rglob("*")
                 if (p.suffix == ".json" or p.name == "presets.ts")
                 and p.name not in {"forever_ranked_profiles.json", "forever_synthetic_item_metadata.json"})
    digest = hashlib.sha256()
    for path in sorted(paths):
        digest.update(str(path).encode())
        digest.update(path.read_bytes())
    manifest = {
        "inputsSHA256": digest.hexdigest(), "iterations": args.iterations,
        "seed": args.seed, "scenarios": list(scenarios),
        "naturalHit": args.natural_hit,
    }
    manifest_path = args.output / "manifest.json"
    if manifest_path.exists() and json.loads(manifest_path.read_text()) != manifest:
        raise SystemExit("Replay inputs changed; use a new output directory.")
    manifest_path.write_text(json.dumps(manifest, indent=2) + "\n")

    def replay(job):
        scenario, build, race = job
        slug = build + "__" + re.sub(r"[^a-z0-9]+", "_", race.lower()).strip("_")
        prefix = args.output / scenario / slug
        prefix.parent.mkdir(exist_ok=True)
        result = prefix.with_suffix(".json")
        csv_path = prefix.with_suffix(".csv")
        cached = False
        if result.exists() and csv_path.exists():
            try:
                cached = len(json.loads(result.read_text())["Results"]) == 1
                with csv_path.open() as stream:
                    cached = cached and len(list(csv.DictReader(stream))) == 1
            except (ValueError, KeyError):
                cached = False
        if not cached:
            profiles, flags = scenarios[scenario]
            command = [
                str(args.binary), "-build", build, "-race", race,
                "-baseline-results", str(profiles), "-iterations", str(args.iterations),
                "-seed", str(args.seed), "-output", str(prefix), *flags,
            ]
            if args.natural_hit and scenario != "original":
                command.append("-natural-hit")
            with prefix.with_suffix(".log").open("w") as log:
                subprocess.run(command, stdout=log, stderr=subprocess.STDOUT,
                               env={**os.environ, "GOMAXPROCS": "1"}, check=True)
        data = json.loads(result.read_text())
        row = data["Results"][0]
        options = row["Request"]["simOptions"]
        if (row["Key"], row["Race"]) != (build, race) or row.get("Warnings"):
            raise ValueError(f"Invalid replay or APL warnings: {job}")
        if args.natural_hit and scenario != "original" and (
                row["Hit"]["Model"] != "equipped-only-v2" or row["Hit"]["RawHitDelta"] != 0):
            raise ValueError(f"Paid hit adjustment in an equipment-only replay: {job}")
        if options["iterations"] != args.iterations or int(options["randomSeed"]) != args.seed:
            raise ValueError(f"Iteration/seed mismatch: {job}")
        if data["Tier1Bonuses"] != (scenario != "tier1_off") or data["EquipmentScale"] != {
            "gear_110": 1.1, "gear_150": 1.5,
        }.get(scenario, 1):
            raise ValueError(f"Scenario mismatch: {job}")
        return job, prefix, data

    jobs = [(scenario, build, race) for scenario in scenarios for build, race in roster]
    results, failures = {}, []
    executor = concurrent.futures.ThreadPoolExecutor(max_workers=args.workers)
    try:
        futures = {executor.submit(replay, job): job for job in jobs}
        for future in concurrent.futures.as_completed(futures):
            try:
                job, prefix, data = future.result()
                results[job] = (prefix, data)
            except Exception as error:
                failures.append({"job": futures[future], "error": str(error)})
            settled = len(results) + len(failures)
            if settled % 25 == 0 or settled == len(jobs):
                print(f"{settled}/{len(jobs)} replays; {len(failures)} failures", flush=True)
    finally:
        executor.shutdown(wait=True, cancel_futures=True)
    (args.output / "summary.json").write_text(json.dumps({
        "completed": len(results), "requested": len(jobs), "failures": failures,
    }, indent=2) + "\n")
    if failures:
        raise SystemExit("Some replays failed; inspect summary.json.")
    for scenario in scenarios:
        members = [results[(scenario, *key)] for key in roster]
        combined = {**members[0][1], "ReplayManifest": manifest,
                    "GearScenario": ("real-reference" if scenario == "original" else
                                     profile_data.get("GearScenario", "real-reference")),
                    "Results": [data["Results"][0] for _, data in members]}
        (args.output / (scenario + ".json")).write_text(json.dumps(combined, indent=2) + "\n")
        with (args.output / (scenario + ".csv")).open("w") as stream:
            writer = csv.writer(stream, lineterminator="\n")
            for index, (prefix, _) in enumerate(members):
                with prefix.with_suffix(".csv").open() as source:
                    rows = list(csv.reader(source))
                if index == 0:
                    writer.writerow(rows[0])
                writer.writerows(rows[1:])


if __name__ == "__main__":
    main()
