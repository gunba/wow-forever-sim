#!/usr/bin/env python3
"""Export benchmark requests as importable individual-simulator settings."""

import argparse
from copy import deepcopy
import json
from pathlib import Path
import re


def ui_settings(row):
    request = row["Request"]
    raid = request["raid"]
    party = raid["parties"][0]
    return deepcopy({
        "settings": {
            "iterations": request["simOptions"]["iterations"],
            "fixedRngSeed": request["simOptions"]["randomSeed"],
            "ruleset": request["simOptions"]["ruleset"],
            "phase": 1,
            "faction": row["Faction"],
            "showDamageMetrics": True,
        },
        "raidBuffs": raid.get("buffs", {}),
        "partyBuffs": party.get("buffs", {}),
        "debuffs": raid.get("debuffs", {}),
        "tanks": raid.get("tanks", []),
        "player": party["players"][0],
        "encounter": request["encounter"],
    })


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--results", required=True, type=Path)
    parser.add_argument("--output", required=True, type=Path)
    args = parser.parse_args()
    data = json.loads(args.results.read_text())
    args.output.mkdir(parents=True, exist_ok=True)
    index = []
    for row in data["Results"]:
        race = re.sub(r"[^a-z0-9]+", "_", row["Race"].lower()).strip("_")
        name = f"{row['Key']}__{race}.json"
        (args.output / name).write_text(json.dumps(ui_settings(row), indent=2) + "\n")
        index.append({"build": row["Build"], "race": row["Race"], "file": name,
                      "dps": row["DPS"], "hitAdjustment": row["Hit"]})
    (args.output / "index.json").write_text(json.dumps(index, indent=2) + "\n")
    print(f"{len(index)} profiles exported to {args.output}")


if __name__ == "__main__":
    main()
