#!/usr/bin/env python3
"""Capture client energy rates and energy-spell inputs for the mechanics audit."""
import argparse
import json
from pathlib import Path
import sys

ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "tools/data_watch"))
import spell_client as client

SPELLS = (
    9830, 9850, 9904, 9896, 31018, 5217, 1312152, 417046, 417045,
    11294, 11281, 25300, 1241584, 16511, 11300, 31016, 6774, 1310703, 11275,
    13750, 14179, 14181, 14983, 1310711,  # Adrenaline Rush, Relentless Strikes, Vigor, Flawless Execution
    14186, 14189,  # Seal Fate passive and triggered combo point
)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--cache", required=True)
    parser.add_argument("--output", type=Path, default=ROOT / "assets/db_inputs/forever_energy.json")
    args = parser.parse_args()
    client.CACHE = args.cache
    builds = (client.FOREVER, "1.60.1.69913", client.ERA)
    power = []
    for build in builds:
        row = next(r for r in client.table(build, "PowerType") if r["PowerTypeEnum"] == "3")
        power.append({
            "build": build,
            "source": f"https://wago.tools/db2/PowerType/csv?build={build}",
            "record": row,
        })
    spells = client.Client()
    records = [spells.spell(id) for id in SPELLS]
    if any(record is None for record in records):
        raise ValueError("An audited spell is absent from the selected client build")
    out = {
        "schemaVersion": 1,
        "powerTypes": power,
        "spellBuild": client.FOREVER,
        "spells": records,
        "procOptions": [r for r in client.table(client.FOREVER, "SpellAuraOptions")
                        if int(r["SpellID"]) == 14186],
        "limitations": [
            "PowerType specifies base regeneration, not server update cadence or haste behavior.",
            "Identical flags across Classic and Forever do not establish identical server mechanics.",
            "SpellPower costs are before talents, sets, temporary discounts and miss refunds.",
        ],
    }
    args.output.write_text(json.dumps(out, indent=2) + "\n")


if __name__ == "__main__":
    main()
