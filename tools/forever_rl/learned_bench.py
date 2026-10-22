#!/usr/bin/env python3
"""Benchmark the learned profiles (synth_all.py) exactly as the matrix replays published ones: the native bench, the
matrix seed and iteration count, enchants refreshed, natural hit. Adds to each artifacts/ml/<key>.json:

  native   the network and the published setup on the same fights, the published setup played by the engine's own APL
           path (apl_synth.py's reference); training's APL emulation chooses cooldowns at decision points and can differ
  profile  the converted profile: bench DPS beside the published cell, its paired gain and its APL edits

and writes artifacts/ml/learned_results.json (results format, for export_ui_profiles.py --learned) with the profiles
that beat the published setup.

    python3 tools/forever_rl/learned_bench.py --binary /tmp/forever-bench --runs ~/.cache/forever-sim-audit/rl/all
"""
import argparse
import json
import re
import subprocess
from concurrent.futures import ThreadPoolExecutor
from pathlib import Path

HERE = Path(__file__).resolve().parent
REPO = HERE.parents[1]


def edits(key, race, talents, published, final):
    """The profile's APL against the published one, in words: (added, removed)."""
    from apl_synth import apl_diff, describe_item
    from forever_env import VecEnv
    probe = VecEnv(dict(build=key, race=race, fullChoices=True, pollSeconds=0, threads=1, labeledRands=True, seed=17,
                        talents=talents), 1)
    spec = probe.spec
    probe.close()
    added, removed = apl_diff(published["priorityList"], final["priorityList"])
    return [describe_item(x, spec) for x in added], [describe_item(x, spec) for x in removed]


def race_file(race):
    return re.sub(r"[^a-z0-9]+", "_", race.lower()).strip("_")


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--binary", type=Path, required=True)
    ap.add_argument("--runs", type=Path, required=True)
    ap.add_argument("--ml", type=Path, default=REPO / "artifacts/ml")
    ap.add_argument("--matrix", type=Path, default=REPO / "artifacts/modelled_gear/forever_dps_5min.json")
    ap.add_argument("--inputs", type=Path, default=REPO / "artifacts/modelled_gear/forever_input_profiles.json")
    ap.add_argument("--iterations", type=int, default=5000)
    ap.add_argument("--seed", type=int, default=1793600001)
    ap.add_argument("--workers", type=int, default=12)
    a = ap.parse_args()
    profiles = a.runs / "learned_profiles.json"
    data = json.load(open(profiles))
    out = a.runs / "learned_bench"
    out.mkdir(exist_ok=True)

    def bench(row, source=profiles, suffix="", extra=("-any-talents",)):
        prefix = out / f"{row['Key']}__{race_file(row['Race'])}{suffix}"
        cmd = [str(a.binary), "-build", row["Key"], "-race", row["Race"], "-baseline-results", str(source),
               "-iterations", str(a.iterations), "-seed", str(a.seed), "-output", str(prefix), "-refresh-enchants", "-natural-hit", *extra]
        subprocess.run(cmd, check=True, cwd=REPO, stdout=subprocess.DEVNULL)
        return json.load(open(prefix.with_suffix(".json")))["Results"][0]

    with ThreadPoolExecutor(a.workers) as pool:
        results = list(pool.map(bench, data["Results"]))
        # The published cells again with this binary: a learned profile is only compared with the same code.
        published = list(pool.map(lambda row: bench(row, a.inputs, "__published", ()), data["Results"]))
    matrix = {(r["Key"], r["Race"]): r for r in json.load(open(a.matrix))["Results"]}
    for row in published:
        cell = matrix[(row["Key"], row["Race"])]["DPS"]
        if abs(row["DPS"] - cell) > 0.01:
            print(f"warning: {row['Key']} {row['Race']} published {row['DPS']:.3f} with this binary, {cell:.3f} in the matrix")
    matrix = {(r["Key"], r["Race"]): r for r in published}
    inputs = {(r["Key"], r["Race"]): r["BaselinePlayer"] for r in json.load(open(a.inputs))["Results"]}
    learned_players = {(r["Key"], r["Race"]): r["BaselinePlayer"] for r in data["Results"]}
    gaining = []
    for row in results:
        key, race = row["Key"], row["Race"]
        report_path = a.ml / f"{key}.json"
        report = json.load(open(report_path))
        pub = matrix[(key, race)]
        synth = a.runs / key / "synth"
        chosen = json.load(open(synth / "report.json"))
        net = json.load(open(synth / "learned" / "report.json"))
        saved = json.load(open(synth / "profile.json"))
        tag = saved.get("TalentSource", saved.get("Talents"))
        added, removed = edits(key, race, row["Talents"], inputs[(key, race)]["rotation"], learned_players[(key, race)]["rotation"])
        report["native"] = {
            "fights": net["fights"], "published": net["reference"], "network": net["network"],
            "network_vs_published": net["network_vs_reference"][0], "network_vs_published_se": net["network_vs_reference"][1],
        }
        gain, se = chosen["synthesised_vs_reference"]
        report["profile"] = {
            "id": f"{key}__{race_file(race)}__learned", "dps": row["DPS"], "se": row["StandardError"],
            "published_dps": pub["DPS"], "published_se": pub["StandardError"], "talents": row["Talents"], "talent_source": tag,
            "overrides": added, "removed": removed,
            "paired_vs_published": gain, "paired_vs_published_se": se, "paired_fights": chosen["fights"],
            "gain": bool(gain > 2 * se and row["DPS"] > pub["DPS"]),
        }
        report_path.write_text(json.dumps(report, indent=1) + "\n")
        if report["profile"]["gain"]:
            gaining.append(row)
        n = report["native"]
        print(f"{key:22s} published {pub['DPS']:7.1f}  profile {row['DPS']:7.1f} ({row['DPS'] - pub['DPS']:+6.1f}, paired {gain:+.1f} ± {se:.1f})"
              f"  network {n['network']:7.1f} ({n['network_vs_published']:+.1f} ± {n['network_vs_published_se']:.1f})", flush=True)
    (a.ml / "learned_results.json").write_text(json.dumps({"GearScenario": data.get("GearScenario"), "Results": gaining}) + "\n")
    index = []
    for path in sorted(a.ml.glob("*.json")):
        r = json.load(open(path))
        if "key" not in r or "native" not in r:
            continue
        n, q = r["native"], r["profile"]
        index.append({"key": r["key"], "race": r["race"], "published_dps": q["published_dps"], "network_dps": n["network"],
                      "network_vs_published": n["network_vs_published"], "network_vs_published_se": n["network_vs_published_se"],
                      "profile": q["id"] if q["gain"] else None, "profile_dps": q["dps"],
                      "profile_vs_published": q["dps"] - q["published_dps"], "talent_changes": len(r["talent_changes"])})
    (a.ml / "index.json").write_text(json.dumps(index, indent=1) + "\n")


if __name__ == "__main__":
    main()
