#!/usr/bin/env python3
"""Bundle a completed build comparison and its search evidence."""

import argparse
import json
import math
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--build", required=True)
    parser.add_argument("--reference", required=True, type=Path)
    parser.add_argument("--candidate", required=True, type=Path)
    parser.add_argument("--rotation-search", action="append", type=Path, default=[])
    parser.add_argument("--talent-search", action="append", type=Path, default=[])
    parser.add_argument("--experiment", action="append", type=Path, default=[])
    parser.add_argument("--note", action="append", default=[])
    parser.add_argument("--output", type=Path)
    args = parser.parse_args()
    reference = json.loads(args.reference.read_text())
    candidate = json.loads(args.candidate.read_text())
    old = {r["Race"]: r for r in reference["Results"]}
    new = {r["Race"]: r for r in candidate["Results"]}
    if old.keys() != new.keys():
        raise ValueError("Race coverage differs")
    if {k: v for k, v in reference.items() if k != "Results"} != {
        k: v for k, v in candidate.items() if k != "Results"
    }:
        raise ValueError("Benchmark settings differ")
    comparisons = []
    for race, r in new.items():
        b = old[race]
        if r["Key"] != args.build or b["Key"] != args.build:
            raise ValueError("Wrong build")
        if r["Warnings"] or b["Warnings"] or min(r["Iterations"], b["Iterations"]) < 5000:
            raise ValueError("Incomplete validation")
        p, q = b["BaselinePlayer"], r["BaselinePlayer"]
        if {k: v for k, v in p.items() if k not in ("talentsString", "rotation")} != {
            k: v for k, v in q.items() if k not in ("talentsString", "rotation")
        }:
            raise ValueError(f"{race}: a fixed player field changed")
        margin = 2.5 * math.hypot(r["StandardError"], b["StandardError"])
        gain = r["DPS"] - b["DPS"]
        if p != q and gain <= margin:
            raise ValueError(f"{race}: changed build failed validation")
        comparisons.append({
            "race": race, "baselineDPS": b["DPS"], "retainedDPS": r["DPS"],
            "gainDPS": gain, "gainPercent": 100*gain/b["DPS"],
            "acceptanceMarginDPS": margin, "changed": p != q,
        })
    rotations = []
    for folder in args.rotation_search:
        ledger = json.loads((folder / "ledger.json").read_text())
        if ledger["retained"] is None:
            raise ValueError(f"Incomplete search: {folder}")
        for trial in ledger["trials"]:
            data = json.loads((folder / trial.pop("result")).read_text())
            trial["request"] = data["Results"][0]["Request"]
            trial["mechanics"] = data.get("Mechanics")
        rotations.append(ledger)
    bundle = {
        "schemaVersion": 1, "build": args.build,
        "comparison": comparisons, "notes": args.note,
        "rotationSearches": rotations,
        "talentSearches": [json.loads(p.read_text()) for p in args.talent_search],
        "additionalExperiments": [json.loads(p.read_text()) for p in args.experiment],
        "reference": reference, "retained": candidate,
    }
    output = args.output or Path("artifacts/optimization") / f"{args.build}.json"
    output.parent.mkdir(parents=True, exist_ok=True)
    output.write_text(json.dumps(bundle, indent=2) + "\n")
    print(output)


if __name__ == "__main__":
    main()
