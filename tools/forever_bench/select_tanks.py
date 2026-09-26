#!/usr/bin/env python3
"""Select independently checked tank recipes against their frozen controls."""

import argparse
import json
import math
from pathlib import Path
import statistics

from research_tanks import failures


def mean(rows, field):
    return statistics.mean(row["Tank"][field] for row in rows)


def pooled_se(rows, field):
    return math.sqrt(sum(row["Tank"][field] ** 2 for row in rows)) / len(rows)


def death_interval(rows):
    # Wilson interval on the independent iteration count. No covariance
    # between control/candidate death indicators is assumed.
    n = sum(row["Iterations"] for row in rows)
    p = sum(row["Tank"]["ChanceOfDeath"] * row["Iterations"] for row in rows) / n
    # Two marginal 97.5% intervals give at least 95% joint coverage
    # (Bonferroni), including zero-observed-death arms.
    z = 2.241402727604947
    center = (p + z*z/(2*n)) / (1 + z*z/n)
    half = z * math.sqrt(p*(1-p)/n + z*z/(4*n*n)) / (1 + z*z/n)
    return p, max(0, center-half), min(1, center+half)


def compare(pairs, controls):
    checks, rejected = {}, []
    for scenario in ("Single", "Multi"):
        candidates = [pair[scenario] for pair in pairs]
        references = [pair[scenario] for pair in controls]
        measures = {}
        for field, error, factor, direction in (
                ("DTPS", "DTPSStandardError", 1.01, "upper"),
                ("TMI", "TMIStandardError", 1.02, "upper"),
                ("TPS", "TPSStandardError", .99, "lower")):
            delta = mean(candidates, field) - factor*mean(references, field)
            bound = 1.96 * (pooled_se(candidates, error) + factor*pooled_se(references, error))
            passed = delta+bound <= 0 if direction == "upper" else delta-bound >= 0
            measures[field] = {"candidate": mean(candidates, field),
                               "reference": mean(references, field),
                               "boundaryDifference": delta, "conservative95Bound": bound,
                               "passed": passed}
            if not passed:
                rejected.append(scenario + ":" + field + " noninferiority unresolved/failed")
        c, r = death_interval(candidates), death_interval(references)
        death_pass = c[2] - r[1] <= .0025
        measures["death"] = {"candidate": c[0], "reference": r[0],
                             "candidateWilson97_5": c[1:], "referenceWilson97_5": r[1:],
                             "conservativeUpperDifference": c[2]-r[1], "passed": death_pass}
        if not death_pass:
            rejected.append(scenario + ":death noninferiority unresolved/failed")
        target_pass = all(c["Tank"]["LeastTargetTPS"] >= .98*r["Tank"]["LeastTargetTPS"]
                          for c, r in zip(candidates, references))
        measures["leastTargetTPS"] = {
            "candidate": mean(candidates, "LeastTargetTPS"),
            "reference": mean(references, "LeastTargetTPS"), "passed": target_pass,
            "uncertainty": "Per-target threat variance is not exported; the threshold must pass each independent seed.",
        }
        if not target_pass:
            rejected.append(scenario + ":least-target TPS")
        if any(row["Warnings"] for row in candidates):
            rejected.append(scenario + ":APL warnings")
        checks[scenario] = measures
    candidates = [pair["Single"] for pair in pairs]
    references = [pair["Single"] for pair in controls]
    delta = statistics.mean(row["DPS"] for row in candidates) - statistics.mean(row["DPS"] for row in references)
    bound = 1.96 * sum(math.sqrt(sum(row["StandardError"]**2 for row in group))/len(group)
                       for group in (candidates, references))
    if delta <= bound:
        rejected.append("DPS gain not established")
    return {"checks": checks, "DPSGain": delta, "DPSConservative95Bound": bound,
            "RejectedFor": rejected,
            "MeanDPS": statistics.mean(row["DPS"] for row in candidates),
            "perSeedMeanGuardFailures": [failures(c, r) for c, r in zip(pairs, controls)]}


def summarize(pairs, scenario):
    rows = [pair[scenario] for pair in pairs]
    result = {field: mean(rows, field) for field in
              ("TPS", "DTPS", "TMI", "ChanceOfDeath", "LeastTargetTPS")}
    for field in ("TPSStandardError", "DTPSStandardError", "TMIStandardError"):
        result[field] = pooled_se(rows, field)
    result["TargetTPS"] = [statistics.mean(values)
                           for values in zip(*(row["Tank"]["TargetTPS"] for row in rows))]
    result["Iterations"] = sum(row["Iterations"] for row in rows)
    return result


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--confirmation", type=Path, nargs="+", required=True)
    parser.add_argument("--output", type=Path, required=True)
    args = parser.parse_args()
    if len(args.confirmation) < 2:
        raise SystemExit("At least two independent seeds are required")
    manifests = [json.loads((path / "manifest.json").read_text()) for path in args.confirmation]
    if len({manifest["seed"] for manifest in manifests}) != len(manifests):
        raise SystemExit("Seeds are not independent")
    if any(set(m["jobs"]) != set(manifests[0]["jobs"]) for m in manifests):
        raise SystemExit("Candidate coverage differs")
    if len({m["binarySHA256"] for m in manifests}) != 1:
        raise SystemExit("Different engines")
    if any((m["iterations"], m.get("iterationsByBuild", {})) !=
           (manifests[0]["iterations"], manifests[0].get("iterationsByBuild", {})) for m in manifests):
        raise SystemExit("Iteration allocation differs between seeds")
    jobs = manifests[0]["jobs"]
    pairs = {uid: [json.loads((path / (uid+".pair.json")).read_text()) for path in args.confirmation]
             for uid in jobs}
    controls = {(job["build"], job["race"]): uid for uid, job in jobs.items()
                if "published-control" in job["labels"]}
    comparisons, selected, rows = [], [], []
    for (build, race), control in controls.items():
        eligible = []
        for uid, job in jobs.items():
            if (job["build"], job["race"]) != (build, race) or uid == control:
                continue
            result = compare(pairs[uid], pairs[control])
            result.update(ID=uid, Key=build, Race=race, Control=control)
            comparisons.append(result)
            if not result["RejectedFor"]:
                eligible.append(result)
        winner = max(eligible, key=lambda row: row["MeanDPS"]) if eligible else None
        uid = winner["ID"] if winner else control
        pair = pairs[uid][0]
        rows.append(pair["Single"])
        selected.append({
            "Key": build, "Race": race, "SelectedID": uid, "ControlID": control,
            "Disposition": "retain" if winner else "retain-published-control",
            "DPSGain": winner["DPSGain"] if winner else 0,
            "DPSConservative95Bound": winner["DPSConservative95Bound"] if winner else 0,
            "Single": summarize(pairs[uid], "Single"), "Multi": summarize(pairs[uid], "Multi"),
        })
        print(build, race, selected[-1]["Disposition"], round(selected[-1]["DPSGain"], 2), flush=True)
    args.output.mkdir(parents=True, exist_ok=True)
    (args.output / "selected.json").write_text(json.dumps({"GearScenario": "modeled-65-v2", "Results": rows}, indent=2)+"\n")
    (args.output / "validation.json").write_text(json.dumps({
        "engineSHA256": manifests[0]["binarySHA256"],
        "seeds": [m["seed"] for m in manifests], "iterationsPerSeed": [m["iterations"] for m in manifests],
        "iterationsPerSeedByBuild": {
            build: [m.get("iterationsByBuild", {}).get(build, m["iterations"]) for m in manifests]
            for build, _ in controls
        },
        "selectionPolicy": "DPS gain plus conservative guard checks against frozen published controls; otherwise retain the control.",
        "scenarios": {"Single": "300s, one 3000-base/2s attacker, 1500 HPS",
                      "Multi": "90s, three 1500-base/2s attackers, 2500 HPS"},
        "selections": selected, "comparisons": comparisons,
    }, indent=2)+"\n")


if __name__ == "__main__":
    main()
