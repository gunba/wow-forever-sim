#!/usr/bin/env python3
"""Export benchmark requests as importable individual-simulator settings."""

import argparse
from copy import deepcopy
import hashlib
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
    parser.add_argument("--bundle", type=Path, help="also write the web default-profile bundle")
    args = parser.parse_args()
    data = json.loads(args.results.read_text())
    args.output.mkdir(parents=True, exist_ok=True)
    index = []
    profiles = []
    for row in data["Results"]:
        race = re.sub(r"[^a-z0-9]+", "_", row["Race"].lower()).strip("_")
        name = f"{row['Key']}__{race}.json"
        (args.output / name).write_text(json.dumps(ui_settings(row), indent=2) + "\n")
        index.append({"build": row["Build"], "race": row["Race"], "file": name,
                      "dps": row["DPS"], "hitAdjustment": row["Hit"]})
        profiles.append({
            "id": f"{row['Key']}__{race}", "key": row["Key"], "build": row["Build"],
            "race": row["Race"], "dps": row["DPS"], "settings": ui_settings(row),
            "hitAdjustment": row["Hit"], "unmodeledSetBonuses": row.get("UnmodeledSetBonuses") or [],
        })
    (args.output / "index.json").write_text(json.dumps(index, indent=2) + "\n")
    if args.bundle:
        if len(profiles) != 147 or len({p["id"] for p in profiles}) != 147:
            raise ValueError("The web bundle requires the complete 147-profile roster")
        for profile in profiles:
            player = profile["settings"]["player"]
            if not player.get("foreverTier1Bonuses") or player.get("equipmentScale", 1) != 1:
                raise ValueError("Web defaults must use the Tier-on, unscaled benchmark")
            if sum(bool(item.get("enchant")) for item in player["equipment"]["items"]) < 9:
                raise ValueError(f"{profile['id']} does not contain a fully enchanted ranking loadout")
        args.bundle.parent.mkdir(parents=True, exist_ok=True)
        args.bundle.write_text(json.dumps({
            "sourceSHA256": hashlib.sha256(args.results.read_bytes()).hexdigest(),
            "profiles": profiles,
        }, separators=(",", ":")) + "\n")
    print(f"{len(index)} profiles exported to {args.output}")


if __name__ == "__main__":
    main()
