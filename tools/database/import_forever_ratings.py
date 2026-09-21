#!/usr/bin/env python3
"""Extract a level's rating coefficients from the Forever client GameTable."""

import argparse
import csv
import hashlib
import io
import json
import math
from pathlib import Path


REQUIRED = (
    "Hit - Melee", "Hit - Ranged", "Hit - Spell",
    "Crit - Melee", "Crit - Ranged", "Crit - Spell",
    "Haste - Melee", "Haste - Ranged", "Haste - Spell",
    "Dodge", "Parry", "Block", "Defense Skill",
)

def load_level_60_coefficients() -> dict[str, float]:
    path = Path(__file__).resolve().parents[2] / "assets/db_inputs/forever_combat_ratings.json"
    data = json.loads(path.read_text())
    if data["level"] != 60:
        raise ValueError("The level-60 importer needs level-60 rating coefficients")
    return data["coefficients"]


def read_coefficients(data: bytes, level: int) -> dict[str, float]:
    rows = list(csv.DictReader(io.StringIO(data.decode("utf-8-sig")), delimiter="\t"))
    matching = [r for r in rows if int(r["Level"]) == level]
    if len(matching) != 1:
        raise ValueError(f"Expected exactly one row for level {level}")
    result = {k: float(v) for k, v in matching[0].items() if k != "Level"}
    for name in REQUIRED:
        if name not in result or not math.isfinite(result[name]) or result[name] <= 0:
            raise ValueError(f"Invalid or missing rating coefficient: {name}")
    return result


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("table", type=Path)
    parser.add_argument("--build", required=True)
    parser.add_argument("--build-config", required=True)
    parser.add_argument("--cdn-config", required=True)
    parser.add_argument("--level", type=int, default=60)
    parser.add_argument("--output", type=Path, default=Path("assets/db_inputs/forever_combat_ratings.json"))
    args = parser.parse_args()
    data = args.table.read_bytes()
    output = {
        "schemaVersion": 1,
        "clientBuild": args.build,
        "level": args.level,
        "source": {
            "product": "wow_classic_beta",
            "buildConfig": args.build_config,
            "cdnConfig": args.cdn_config,
            "fileDataId": 1391669,
            "filename": "GameTables/CombatRatings.txt",
            "sha256": hashlib.sha256(data).hexdigest(),
        },
        "coefficients": read_coefficients(data, args.level),
    }
    args.output.write_text(json.dumps(output, indent=2) + "\n")


if __name__ == "__main__":
    main()
