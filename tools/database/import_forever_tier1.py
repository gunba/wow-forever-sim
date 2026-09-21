#!/usr/bin/env python3
"""Capture the Forever Tier 1 definitions, independently of item availability."""
import argparse
import collections
import json
import pathlib
import sys

ROOT = pathlib.Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "tools/data_watch"))
import spell_client as client


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--cache", required=True)
    args = parser.parse_args()
    client.CACHE = args.cache
    cl = client.Client()
    sets = {row["ID"]: row for row in client.table(client.FOREVER, "ItemSet")}
    bonuses = collections.defaultdict(list)
    for row in client.table(client.FOREVER, "ItemSetSpell"):
        bonuses[row["ItemSetID"]].append(row)
    result = []
    for set_id, rows in sorted(bonuses.items(), key=lambda entry: int(entry[0])):
        if not any(
            "1.60.0" in cl.names.get(int(row["SpellID"]), "")
            and "Tier 1" in cl.names.get(int(row["SpellID"]), "")
            for row in rows
        ):
            continue
        result.append({
            "id": int(set_id),
            "name": sets[set_id]["Name_lang"],
            "bonuses": [
                {
                    "pieces": int(row["Threshold"]),
                    "spell": cl.spell(row["SpellID"]),
                    "rawEffects": cl.effects[int(row["SpellID"])],
                }
                for row in sorted(rows, key=lambda row: int(row["Threshold"]))
            ],
        })
    destination = ROOT / "assets/db_inputs/forever_tier1_bonuses.json"
    destination.write_text(json.dumps({
        "build": client.FOREVER,
        "sources": [
            f"https://wago.tools/db2/{table}?build={client.FOREVER}"
            for table in ("ItemSet", "ItemSetSpell", "SpellEffect", "Spell")
        ],
        "sets": result,
    }, indent=2) + "\n")
    print(f"{len(result)} Tier 1 sets -> {destination.relative_to(ROOT)}")


if __name__ == "__main__":
    main()
