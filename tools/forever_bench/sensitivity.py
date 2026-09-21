#!/usr/bin/env python3
"""Compare fixed-profile Tier 1 and proportional equipment sensitivities."""

import argparse
import copy
import hashlib
import json
from pathlib import Path
import statistics

SCENARIOS = [
    ("tier1", "Tier 1 gain", "tier1_off", 1, False),
    ("gear10", "Gear +10%", "gear_110", 1.1, True),
    ("gear20", "Gear +20%", "gear_120", 1.2, True),
]


def gain(new, reference):
    """Percent gain and a conservative delta-method 95% Monte Carlo bound."""
    a, b = new["DPS"], reference["DPS"]
    if a <= 0 or b <= 0:
        raise ValueError("DPS must be positive")
    # Common seeds correlate simulations. Without iteration-paired samples,
    # use the sum of marginal SE contributions, not an independence assumption.
    se_bound = 100 * (new["StandardError"] / b + a * reference["StandardError"] / b**2)
    return 100 * (a / b - 1), 1.96 * se_bound


def compare(base, scenario, scale, tier):
    if (base["Duration"], base["StartingArmor"], base["Mechanics"]) != (
            scenario["Duration"], scenario["StartingArmor"], scenario["Mechanics"]):
        raise ValueError("Scenario mechanics or encounter differ")
    if scenario["Tier1Bonuses"] != tier or scenario["EquipmentScale"] != scale:
        raise ValueError("Wrong sensitivity scenario")
    controls = {(r["Key"], r["Race"]): r for r in base["Results"]}
    results = {(r["Key"], r["Race"]): r for r in scenario["Results"]}
    if len(controls) != len(base["Results"]) or len(results) != len(scenario["Results"]) or controls.keys() != results.keys():
        raise ValueError("Race/build coverage differs or contains duplicates")
    paired = []
    for pair, control in controls.items():
        row = results[pair]
        expected = copy.deepcopy(control["BaselinePlayer"])
        if tier:
            expected["foreverTier1Bonuses"] = True
        else:
            expected.pop("foreverTier1Bonuses", None)
        if scale != 1:
            expected["equipmentScale"] = scale
        if expected != row["BaselinePlayer"]:
            raise ValueError(f"Unexpected profile change: {pair}")
        # Only the player and its separately budgeted hit adjustment may differ.
        expected_request = copy.deepcopy(control["Request"])
        expected_request["raid"]["parties"][0]["players"][0] = row["Request"]["raid"]["parties"][0]["players"][0]
        if expected_request != row["Request"] or row["Iterations"] != control["Iterations"]:
            raise ValueError(f"Encounter/buffs/seed/iterations changed: {pair}")
        if abs(row["Hit"]["Balance"]) > 1e-6:
            raise ValueError(f"Unbalanced hit budget: {pair}")
        new, reference = (control, row) if not tier else (row, control)
        pct, bound = gain(new, reference)
        paired.append({
            "Key": pair[0], "Race": pair[1], "Faction": row["Faction"],
            "GainPercent": pct, "MonteCarlo95Bound": bound,
            "BaselineDPS": control["DPS"], "ScenarioDPS": row["DPS"],
            "OOMSeconds": row["OOMSeconds"], "Warnings": row["Warnings"],
        })
    return paired


def summarize(pairs, faction="all"):
    selected = [p for p in pairs if faction == "all" or p["Faction"].lower() == faction]
    # Equal race weights; common seeds can correlate races too, so averaging
    # the conservative bounds avoids incorrectly treating races as independent.
    return {
        "Races": len(selected),
        "GainPercent": statistics.mean(p["GainPercent"] for p in selected),
        "MonteCarlo95Bound": statistics.mean(p["MonteCarlo95Bound"] for p in selected),
    }


def load_columns(path, results_path, faction="all"):
    if path is None:
        return {}
    data = json.loads(path.read_text())
    if data["BaselineSHA256"] != hashlib.sha256(results_path.read_bytes()).hexdigest():
        raise ValueError("Sensitivity columns belong to another baseline")
    return {b["Key"]: {key: summarize(b["Comparisons"][key], faction)
                       for key, *_ in SCENARIOS} for b in data["Builds"]}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--baseline", type=Path, default=Path("artifacts/forever_dps_5min.json"))
    parser.add_argument("--directory", type=Path, default=Path("artifacts/sensitivity"))
    parser.add_argument("--output", type=Path, default=Path("artifacts/forever_sensitivity.json"))
    args = parser.parse_args()
    base = json.loads(args.baseline.read_text())
    comparisons = {}
    sources = {}
    for key, _, filename, scale, tier in SCENARIOS:
        path = args.directory / f"{filename}.json"
        comparisons[key] = compare(base, json.loads(path.read_text()), scale, tier)
        sources[key] = {"Path": str(path), "SHA256": hashlib.sha256(path.read_bytes()).hexdigest()}
    builds = []
    for key in dict.fromkeys(r["Key"] for r in base["Results"]):
        pairs = {name: [p for p in rows if p["Key"] == key] for name, rows in comparisons.items()}
        builds.append({"Key": key, "Averages": {name: summarize(rows) for name, rows in pairs.items()},
                       "Comparisons": pairs})
    payload = {
        "Baseline": str(args.baseline),
        "BaselineSHA256": hashlib.sha256(args.baseline.read_bytes()).hexdigest(),
        "Sources": sources,
        "Model": {
            "tier1": "100*(DPS with Tier 1 / DPS without Tier 1 - 1); fixed talents and APL",
            "gear": "100*(scaled DPS / baseline DPS - 1); Tier 1 remains on",
            "scaling": "Item and suffix stats, weapon min/max and item flat bonus damage scale; enchants, weapon speed/skill, procs, sets and external effects stay fixed. Paid hit is recalculated.",
            "averaging": "Arithmetic mean of per-race percentage gains, equal weight per available race",
            "uncertainty": "Conservative delta-method 95% Monte Carlo bound from marginal SEs; allows within- and between-race correlation. Not a mechanics confidence interval.",
            "interpretation": "Hypothetical proportional upgrades, not stat weights or actual future items. Fixed rotations can encounter resource/rotation thresholds. Cat weapon-DPS scaling remains unresolved.",
        },
        "Builds": builds,
    }
    args.output.write_text(json.dumps(payload, indent=2) + "\n")
    print(args.output)


if __name__ == "__main__":
    main()
