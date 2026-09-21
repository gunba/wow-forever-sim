#!/usr/bin/env python3
"""Capture the Forever character-creation roster and racial spell evidence."""

import argparse
import json
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "data_watch"))
import spell_client as client

RACES = {
    1: ("RaceHuman", "Human"),
    2: ("RaceOrc", "Orc"),
    3: ("RaceDwarf", "Dwarf"),
    4: ("RaceNightElf", "Night Elf"),
    5: ("RaceUndead", "Undead"),
    6: ("RaceTauren", "Tauren"),
    7: ("RaceGnome", "Gnome"),
    8: ("RaceTroll", "Troll"),
    95: ("RaceSkyborneHighOrder", "High Order"),
    96: ("RaceSkyborneWindshaper", "Windshaper"),
}
RACIALS = [
    20591, 1259802, 1259803,  # Expansive Mind: mana, rage, energy
    1259812, 1259813, 1259817, 1259821, 1259823,  # Eureka class variants
    1259799, 20582, 20597, 20598, 1259719,
    1259705, 1259691, 1270842, 1259710, 1259707,
]


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--cache", required=True)
    parser.add_argument("--output", default="assets/db_inputs/forever_races.json")
    args = parser.parse_args()
    client.CACHE = args.cache
    build = client.FOREVER
    c = client.Client(build)
    races = {int(r["ID"]): r for r in client.table(build, "ChrRaces")}
    classes = {int(r["ID"]): r["Name_lang"] for r in client.table(build, "ChrClasses")}
    combinations = client.table(build, "CharBaseInfo")
    skills = client.table(build, "SkillLineAbility")
    aura_options = client.table(build, "SpellAuraOptions")
    roster = []
    for race_id, (proto_name, name) in RACES.items():
        r = races[race_id]
        roster.append({
            "clientId": race_id,
            "proto": proto_name,
            "name": name,
            "faction": "Alliance" if r["Alliance"] == "0" else "Horde",
            "playableBit": int(r["PlayableRaceBit"]),
            "classes": sorted(classes[int(p["ClassID"])] for p in combinations if int(p["RaceID"]) == race_id),
        })
    evidence = []
    for spell_id in RACIALS:
        evidence.append({
            "spell": c.spell(spell_id),
            "rawEffects": c.effects[spell_id],
            "auraOptions": [r for r in aura_options if int(r["SpellID"]) == spell_id],
            "skillLineAbilities": [r for r in skills if int(r["Spell"]) == spell_id],
        })
    data = {
        "build": build,
        "source": f"https://wago.tools/db2/CharBaseInfo?build={build}",
        "races": roster,
        "racialSpells": evidence,
    }
    Path(args.output).write_text(json.dumps(data, indent=2) + "\n")
    print(f"{len(roster)} races, {sum(len(r['classes']) for r in roster)} race/class pairs")


if __name__ == "__main__":
    main()
