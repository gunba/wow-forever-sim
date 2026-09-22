#!/usr/bin/env python3
"""Rebuild race-weighted comparison summaries from the archived native runs."""

import argparse
from copy import deepcopy
import gzip
import json
import math
from pathlib import Path


def summarize(data):
    output = {key: data[key] for key in ("EngineBaseRevision", "Correction", "Seeds", "IterationsPerArmPerSeed")}
    output["Choices"] = []
    for choice in data["Choices"]:
        item = deepcopy(choice)
        comparisons = [row for row in data["Comparisons"] if row["build"] == choice["Key"]]
        item["Races"] = []
        for race in sorted({row["race"] for row in comparisons}):
            arms = {}
            for candidate in ("baseline", choice["Candidate"]):
                entries = [row for row in comparisons if row["race"] == race and row["candidate"] == candidate]
                if sorted(row["seed"] for row in entries) != sorted(data["Seeds"]):
                    raise ValueError(f"{choice['Key']}/{race}/{candidate}: unmatched seeds")
                runs = [data["Runs"][row["path"]] for row in entries]
                if any(row["Warnings"] or abs(row["Hit"]["Balance"]) > 1e-6 for row in runs):
                    raise ValueError(f"{choice['Key']}/{race}/{candidate}: invalid run")
                arms[candidate] = {
                    "DPS": sum(row["DPS"] for row in runs) / len(runs),
                    "SE": math.sqrt(sum(row["StandardError"] ** 2 for row in runs)) / len(runs),
                    "OOM": sum(row["OOMSeconds"] for row in runs) / len(runs),
                }
            base, new = arms["baseline"], arms[choice["Candidate"]]
            item["Races"].append({
                "Race": race, "BaselineDPS": base["DPS"], "CandidateDPS": new["DPS"],
                "GainDPS": new["DPS"] - base["DPS"], "GainPercent": 100 * (new["DPS"] / base["DPS"] - 1),
                "Conservative95BoundDPS": 1.96 * (new["SE"] + base["SE"]),
                "BaselineOOM": base["OOM"], "CandidateOOM": new["OOM"],
            })
        item["MeanGainPercent"] = sum(row["GainPercent"] for row in item["Races"]) / len(item["Races"])
        output["Choices"].append(item)
    return output


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--archive", type=Path, default=Path("artifacts/research_builds/validation.json.gz"))
    parser.add_argument("--output", type=Path, default=Path("artifacts/research_builds/summary.json"))
    args = parser.parse_args()
    with gzip.open(args.archive, "rt") as stream:
        result = summarize(json.load(stream))
    args.output.write_text(json.dumps(result, indent=2) + "\n")


if __name__ == "__main__":
    main()
