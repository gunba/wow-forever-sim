#!/usr/bin/env python3
"""Compare class-associated client damage records with a native spell inventory."""

import argparse
import collections
import csv
import hashlib
import json
from pathlib import Path

from spell_client import CLASS_MASK, CLASS_SKILLS, FOREVER


def audit(client_data, registered_path, build):
    sources = {}

    def table(name):
        path = client_data / f"{name}_{build}.csv"
        sources[name] = hashlib.sha256(path.read_bytes()).hexdigest()
        with path.open() as stream:
            return list(csv.DictReader(stream))

    names = {int(r["ID"]): r["Name_lang"] for r in table("SpellName")}
    descriptions = {int(r["ID"]): r.get("Description_lang", "") for r in table("Spell")}
    levels = {int(r["SpellID"]): int(r["SpellLevel"]) for r in table("SpellLevels") if r["DifficultyID"] == "0"}
    effects = collections.defaultdict(list)
    for row in table("SpellEffect"):
        if row["DifficultyID"] == "0":
            effects[int(row["SpellID"])].append(row)
    registered = collections.defaultdict(set)
    for row in json.loads(registered_path.read_text()):
        registered[row["Class"].removeprefix("Class").lower()].update(row["SpellIDs"])
        registered[row["Class"].removeprefix("Class").lower()].update(row.get("PetSpellIDs", []))
    aliases = collections.defaultdict(set)
    root = Path(__file__).resolve().parents[2]
    for path in (root / "ui/core/spells").glob("*.json"):
        for native, info in json.loads(path.read_text()).items():
            if info.get("foreverId"):
                aliases[int(info["foreverId"])].add(int(native))

    def has_damage(spell, seen=None):
        seen = set() if seen is None else seen
        if spell in seen:
            return False
        seen.add(spell)
        for row in effects[spell]:
            if int(row["Effect"]) in (2, 9, 17, 31, 58, 62, 121) or int(row["EffectAura"]) in (3, 53, 89):
                return True
            child = int(row["EffectTriggerSpell"])
            if child and has_damage(child, seen):
                return True
        return False

    def descendants(spell, seen=None):
        seen = set() if seen is None else seen
        if spell in seen:
            return seen
        seen.add(spell)
        for row in effects[spell]:
            child = int(row["EffectTriggerSpell"])
            if child:
                descendants(child, seen)
        return seen

    families = {}
    excluded = []
    for row in table("SkillLineAbility"):
        spell = int(row["Spell"])
        if not has_damage(spell):
            continue
        for cls, mask in CLASS_MASK.items():
            if not (int(row["ClassMask"] or 0) & mask or int(row["SkillLine"]) in CLASS_SKILLS[cls]):
                continue
            entry = {
                "SpellID": spell, "Level": levels.get(spell, 0), "SkillLine": int(row["SkillLine"]),
                "ClassMask": int(row["ClassMask"] or 0), "RaceMask": int(row["RaceMasks_0"]),
                "AcquireMethod": int(row["AcquireMethod"]), "SkillRowID": int(row["ID"]),
                "Description": descriptions.get(spell, ""),
                "RegisteredIDs": sorted((descendants(spell) | aliases[spell]) & registered[cls]),
            }
            if entry["Level"] > 60 or entry["AcquireMethod"] == 3:
                excluded.append({"Class": cls, "Name": names.get(spell, str(spell)), **entry})
                continue
            key = (cls, names.get(spell, str(spell)))
            family = families.setdefault(key, {"Class": cls, "Name": key[1], "Ranks": []})
            if not any(rank["SpellID"] == spell for rank in family["Ranks"]):
                family["Ranks"].append(entry)
    rows = []
    dispositions = json.loads((root / "assets/spell_coverage_dispositions.json").read_text())
    for key in sorted(families):
        family = families[key]
        family["Ranks"].sort(key=lambda r: (r["Level"], r["SpellID"]))
        family["AnyRegistered"] = any(r["RegisteredIDs"] for r in family["Ranks"])
        if family["AnyRegistered"]:
            family["Disposition"] = "registered"
            family["Detail"] = "At least one direct, aliased or triggered action is registered; inspect individual ranks."
        else:
            try:
                family["Disposition"], family["Detail"] = dispositions[key[0]][key[1]]
            except KeyError:
                raise ValueError(f"No disposition for {key[0]}/{key[1]}") from None
        rows.append(family)
    return {
        "ClientBuild": build, "SourceSHA256": sources,
        "InventorySHA256": hashlib.sha256(registered_path.read_bytes()).hexdigest(),
        "Scope": "Class-mask and class-skill-line damage candidates at level 60 or below. Association is not proof of current availability. Registration is not proof that every effect is verified.",
        "Families": rows, "ExcludedAcquisitionOrLevel": excluded,
    }


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--client-data", type=Path, required=True)
    parser.add_argument("--registered", type=Path, required=True)
    parser.add_argument("--build", default=FOREVER)
    parser.add_argument("--output", type=Path, required=True)
    args = parser.parse_args()
    result = audit(args.client_data, args.registered, args.build)
    args.output.parent.mkdir(parents=True, exist_ok=True)
    args.output.write_text(json.dumps(result, indent=2) + "\n")
    print(f"{len(result['Families'])} candidate damage families; "
          f"{sum(not row['AnyRegistered'] for row in result['Families'])} absent from the sampled spellbooks")


if __name__ == "__main__":
    main()
