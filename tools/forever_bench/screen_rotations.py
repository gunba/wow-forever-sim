#!/usr/bin/env python3
"""Screen APL variants, independently validate, then confirm a retained change."""

import argparse
import hashlib
import json
import math
from pathlib import Path
import subprocess

from rotation_candidates import variants


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--binary", required=True, type=Path)
    parser.add_argument("--build", required=True)
    parser.add_argument("--race", required=True)
    parser.add_argument("--output", required=True, type=Path)
    parser.add_argument("--profiles", type=Path, default=Path("artifacts/forever_optimization_baselines.json"))
    parser.add_argument("--screen", type=int, default=250)
    parser.add_argument("--validate", type=int, default=2000)
    parser.add_argument("--confirm", type=int, default=5000)
    parser.add_argument("--seed", type=int, default=20261031)
    parser.add_argument("--candidate-prefix", action="append", default=[],
                        help="Screen only matching candidate labels, plus the baseline")
    args = parser.parse_args()
    args.output.mkdir(parents=True, exist_ok=False)
    baseline = json.loads(args.profiles.read_text())["profiles"][args.build]
    players = dict(variants(args.build, baseline))
    if args.candidate_prefix:
        players = {label: p for label, p in players.items()
                   if label == "baseline" or any(label.startswith(prefix) for prefix in args.candidate_prefix)}
        if len(players) == 1:
            raise ValueError("No candidates match the selected prefixes")
    for label, player in players.items():
        # Rotation is the only candidate field; the binary selects the race.
        if {k: v for k, v in player.items() if k != "rotation"} != {
            k: v for k, v in baseline.items() if k != "rotation"
        }:
            raise ValueError(f"{label} changed a non-rotation field")
        (args.output / f"{label}.player.json").write_text(json.dumps(player, indent=2) + "\n")
    ledger = {
        "schemaVersion": 1, "build": args.build, "race": args.race,
        "binarySha256": hashlib.sha256(args.binary.read_bytes()).hexdigest(),
        "profilesSha256": hashlib.sha256(args.profiles.read_bytes()).hexdigest(),
        "acceptanceStandardErrors": 2.5, "trials": [], "retained": None,
    }

    def persist():
        (args.output / "ledger.json").write_text(json.dumps(ledger, indent=2) + "\n")

    def run(label, stage, iterations, seed):
        prefix = args.output / f"{stage}-{label}"
        cmd = [str(args.binary.resolve()), "-build", args.build, "-race", args.race,
               "-player", str(args.output / f"{label}.player.json"), "-iterations", str(iterations),
               "-seed", str(seed), "-duration", "300", "-output", str(prefix)]
        with Path(str(prefix) + ".log").open("w") as log:
            subprocess.run(cmd, stdout=log, stderr=subprocess.STDOUT, check=True)
        row = json.loads(Path(str(prefix) + ".json").read_text())["Results"][0]
        trial = {
            "label": label, "stage": stage, "seed": seed, "iterations": iterations,
            "result": prefix.name + ".json", "dps": row["DPS"],
            "standardError": row["StandardError"], "oomSeconds": row["OOMSeconds"],
            "warnings": row["Warnings"],
        }
        ledger["trials"].append(trial)
        persist()
        print(stage, label, f"{row['DPS']:.2f}", f"OOM {row['OOMSeconds']:.2f}", flush=True)
        return trial

    def improves(candidate, reference):
        margin = 2.5 * math.hypot(candidate["standardError"], reference["standardError"])
        return not candidate["warnings"] and candidate["dps"] > reference["dps"] + margin

    screened = [run(label, "screen", args.screen, args.seed) for label in players]
    baseline_screen = screened[0]
    if baseline_screen["warnings"]:
        raise ValueError("Baseline APL warnings")
    shortlist = sorted(
        (r for r in screened[1:] if not r["warnings"] and r["dps"] > baseline_screen["dps"]),
        key=lambda r: r["dps"], reverse=True,
    )[:5]
    winner = run("baseline", "validate", args.validate, args.seed + 10000)
    for row in shortlist:
        candidate = run(row["label"], "validate", args.validate, args.seed + 10000)
        if improves(candidate, winner):
            winner = candidate
    selected = "baseline"
    if winner["label"] != "baseline":
        reference = run("baseline", "confirm", args.confirm, args.seed + 20000)
        candidate = run(winner["label"], "confirm", args.confirm, args.seed + 20000)
        if improves(candidate, reference):
            selected = candidate["label"]
    ledger["retained"] = selected
    persist()
    (args.output / "retained.player.json").write_text(json.dumps(players[selected], indent=2) + "\n")
    print("RETAIN", selected, flush=True)


if __name__ == "__main__":
    main()
