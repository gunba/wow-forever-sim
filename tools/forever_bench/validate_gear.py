#!/usr/bin/env python3
"""Confirm coordinate-search loadouts with a second independent full-length pair."""

import argparse
import concurrent.futures
from copy import deepcopy
import hashlib
import json
import math
import os
from pathlib import Path
import re
import subprocess

from build_display import expected_roster


def fixed_player(row):
    player = deepcopy(row["BaselinePlayer"])
    player.pop("equipment")
    player.pop("profession2", None)
    return player


def compare_inputs(base, candidate):
    if fixed_player(base) != fixed_player(candidate):
        raise ValueError(f"Non-gear profile changed: {base['Key']}/{base['Race']}")
    request = deepcopy(base["Request"])
    request["raid"]["parties"][0]["players"][0] = candidate["Request"]["raid"]["parties"][0]["players"][0]
    if request != candidate["Request"]:
        raise ValueError("Encounter, buffs or simulation options changed between arms")
    for row in (base, candidate):
        if row["Warnings"] or abs(row["Hit"]["Balance"]) > 1e-6:
            raise ValueError("APL warnings or an unbalanced paid-hit budget")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--binary", required=True, type=Path)
    parser.add_argument("--search", required=True, type=Path)
    parser.add_argument("--output", required=True, type=Path)
    parser.add_argument("--seed", type=int, default=20296071)
    parser.add_argument("--iterations", type=int, default=5000)
    cpus = len(os.sched_getaffinity(0)) if hasattr(os, "sched_getaffinity") else os.cpu_count() or 1
    parser.add_argument("--workers", type=int, default=cpus)
    args = parser.parse_args()
    args.binary = args.binary.resolve()
    args.search = args.search.resolve()
    args.output.mkdir(parents=True, exist_ok=True)
    inputs = {arm: args.search / name for arm, name in
              (("baseline", "baseline.json"), ("candidate", "results.json"))}
    payloads = {arm: json.loads(path.read_text()) for arm, path in inputs.items()}
    rows = {arm: {(row["Key"], row["Race"]): row for row in data["Results"]}
            for arm, data in payloads.items()}
    if any(set(lookup) != expected_roster() or len(data["Results"]) != len(lookup)
           for arm, lookup in rows.items() for data in [payloads[arm]]):
        raise ValueError("Expected the complete race/build roster")
    changed = []
    for pair, base in rows["baseline"].items():
        candidate = rows["candidate"][pair]
        compare_inputs(base, candidate)
        for row in (base, candidate):
            options = row["Request"]["simOptions"]
            if options["iterations"] != args.iterations or int(options["randomSeed"]) == args.seed:
                raise ValueError("First confirmation pair needs equal iterations and an independent seed")
        if base["BaselinePlayer"] != candidate["BaselinePlayer"]:
            changed.append(pair)

    digest = hashlib.sha256()
    for path in [args.binary, *inputs.values()]:
        digest.update(path.read_bytes())
    manifest = {"inputsSHA256": digest.hexdigest(), "seed": args.seed, "iterations": args.iterations}
    manifest_path = args.output / "manifest.json"
    if manifest_path.exists() and json.loads(manifest_path.read_text()) != manifest:
        raise ValueError("Validation inputs changed; use a new output directory")
    manifest_path.write_text(json.dumps(manifest, indent=2) + "\n")

    def run(job):
        arm, pair = job
        key, race = pair
        slug = key + "__" + re.sub(r"[^a-z0-9]+", "_", race.lower()).strip("_")
        prefix = args.output / (arm + "__" + slug)
        path = prefix.with_suffix(".json")
        if not path.exists():
            command = [str(args.binary), "-build", key, "-race", race,
                       "-baseline-results", str(inputs[arm]), "-iterations", str(args.iterations),
                       "-seed", str(args.seed), "-output", str(prefix)]
            with prefix.with_suffix(".log").open("w") as log:
                subprocess.run(command, stdout=log, stderr=subprocess.STDOUT, check=True,
                               env={**os.environ, "GOMAXPROCS": "1"})
        row = json.loads(path.read_text())["Results"][0]
        if (row["Key"], row["Race"]) != pair or row["BaselinePlayer"] != rows[arm][pair]["BaselinePlayer"]:
            raise ValueError("Confirmation replay changed its input profile")
        options = row["Request"]["simOptions"]
        if options["iterations"] != args.iterations or int(options["randomSeed"]) != args.seed:
            raise ValueError("Confirmation replay has the wrong seed or iteration count")
        return job, row

    jobs = [(arm, pair) for pair in changed for arm in inputs]
    with concurrent.futures.ThreadPoolExecutor(max_workers=args.workers) as executor:
        second = dict(executor.map(run, jobs))
    selections, evidence = [], []
    selected = deepcopy(payloads["candidate"])
    selected["Results"] = []
    for pair, original in rows["baseline"].items():
        candidate = rows["candidate"][pair]
        if pair not in changed:
            selected["Results"].append(original)
            selections.append({"Key": pair[0], "Race": pair[1], "Decision": "unchanged"})
            continue
        base2, candidate2 = second[("baseline", pair)], second[("candidate", pair)]
        compare_inputs(base2, candidate2)
        base_mean = (original["DPS"] + base2["DPS"]) / 2
        candidate_mean = (candidate["DPS"] + candidate2["DPS"]) / 2
        base_se = math.hypot(original["StandardError"], base2["StandardError"]) / 2
        candidate_se = math.hypot(candidate["StandardError"], candidate2["StandardError"]) / 2
        delta = candidate_mean - base_mean
        bound = 1.96 * (base_se + candidate_se)
        keep = delta > bound
        selections.append({
            "Key": pair[0], "Race": pair[1], "Decision": "retain" if keep else "revert",
            "BaselineDPS": base_mean, "CandidateDPS": candidate_mean,
            "GainDPS": delta, "GainPercent": 100 * delta / base_mean,
            "Conservative95BoundDPS": bound,
        })
        evidence.append({"Key": pair[0], "Race": pair[1],
                         "Baseline": [original, base2], "Candidate": [candidate, candidate2]})
        selected["Results"].append(candidate if keep else original)
    for name, value in (("selections", selections), ("pairs", evidence), ("selected", selected)):
        (args.output / (name + ".json")).write_text(json.dumps(value, indent=2) + "\n")
    counts = {decision: sum(row["Decision"] == decision for row in selections)
              for decision in ("retain", "revert", "unchanged")}
    print(json.dumps(counts), flush=True)


if __name__ == "__main__":
    main()
