#!/usr/bin/env python3
"""Translate the source catalog into UIItem stats and equipment metadata.

Proc and set effects are separate runtime registrations. Translation alone does
not certify that an item's effects are implemented.
"""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from import_forever_ratings import load_level_60_coefficients
from import_forever_vendor import (
    CLASS_MASK_TO_PROTO_CLASS, INVENTORY_TO_TYPE, WEAPON_TYPES,
    build_stats, class_allowlist, read_vendor_exports, weapon_fields,
)


STAT_INDEX = {
    "Strength": 0, "Agility": 1, "Stamina": 2, "Intellect": 3, "Spirit": 4,
    "SpellPower": 5, "ArcanePower": 6, "FirePower": 7, "FrostPower": 8,
    "HolyPower": 9, "NaturePower": 10, "ShadowPower": 11, "MP5": 12,
    "SpellPenetration": 16, "AttackPower": 17, "Mana": 23, "RangedAttackPower": 27,
    "BlockValue": 30, "Health": 34, "ArcaneResistance": 35, "FireResistance": 36,
    "FrostResistance": 37, "NatureResistance": 38, "ShadowResistance": 39,
    "BonusArmor": 40, "HealingPower": 41, "SpellDamage": 42,
}
RATINGS = {
    "HitRating": (18, "Hit - Melee"), "CritRating": (19, "Crit - Melee"),
    "DefenseRating": (28, "Defense Skill"), "BlockRating": (29, "Block"),
    "DodgeRating": (31, "Dodge"), "ParryRating": (32, "Parry"),
    "ExpertiseRating": (22, "Expertise"),
}
SKILLS = {
    "AxesSkill": 1, "SwordsSkill": 2, "MacesSkill": 3, "DaggersSkill": 4,
    "FistWeaponsSkill": 5, "TwoHandedAxesSkill": 6, "TwoHandedSwordsSkill": 7,
    "TwoHandedMacesSkill": 8, "PolearmsSkill": 9, "StavesSkill": 10,
    "ThrownSkill": 11, "BowsSkill": 12, "CrossbowsSkill": 13, "GunsSkill": 14,
}
NONCOMBAT_STATS = {"MiningSkill", "HerbalismSkill"}
PROFESSIONS = {171: 1, 164: 2, 333: 3, 202: 4, 182: 5, 165: 8, 186: 9, 393: 10, 197: 11}


def compile_stats(item: dict, rates: dict) -> tuple[list[float], list[float], float]:
    stats, skills = [0.0] * 44, [0.0] * 16
    physical_damage = 0
    for name, value in item["stats"].items():
        if name in STAT_INDEX:
            stats[STAT_INDEX[name]] += value
        elif name in RATINGS:
            index, coefficient = RATINGS[name]
            stats[index] += value / rates[coefficient]
        elif name == "HasteRating":
            stats[15] += value / rates["Haste - Spell"]
            stats[20] += value / rates["Haste - Melee"]
        elif name in SKILLS:
            skills[SKILLS[name]] += value
        elif name == "BonusPhysicalDamage":
            physical_damage += value
        elif name not in NONCOMBAT_STATS:
            raise ValueError(f"unsupported stat {name}={value}")
    # Forever's published armor line excludes the separate Bonus Armor line:
    # Cloak of Warding is 44 base + 170 bonus, not 170 within a total of 44.
    stats[26] = item["armor"]
    stats[27] += stats[17]
    return stats, skills, physical_damage


def compile_item(item: dict, rates: dict, vendor: dict | None = None) -> dict:
    stats, skills, physical_damage = compile_stats(item, rates)
    inv, cls, sub = item["inventoryType"], item["class"], item["subclass"]
    out = {
        "id": item["id"], "name": item["name"], "icon": item["icon"],
        "type": INVENTORY_TO_TYPE.get(inv, 0), "stats": stats,
        "ilvl": item["itemLevel"], "quality": item["quality"], "phase": 1, "expansion": 1,
    }
    if cls == 4 and sub in (1, 2, 3, 4):
        out["armorType"] = sub
    if inv in (14, 23):
        out.update(type=13, weaponType=7 if inv == 14 else 5, handType=3)
    elif inv == 28:
        out.update(type=14, rangedWeaponType={7: 5, 8: 4, 9: 7}[sub])
    elif cls == 2:
        ranged = {2: 1, 18: 2, 3: 3, 16: 6, 19: 8}.get(sub)
        if ranged:
            out.update(type=14, rangedWeaponType=ranged)
        else:
            out.update(type=13, weaponType=WEAPON_TYPES[sub],
                       handType={17: 4, 21: 1, 22: 3, 13: 2}[inv])
        weapon = item["weapon"]
        out.update(weaponDamageMin=weapon["min"], weaponDamageMax=weapon["max"],
                   weaponSpeed=weapon["speed"])
    if not out["type"]:
        raise ValueError(f"unsupported inventory type {inv}")
    if any(skills):
        out["weaponSkills"] = skills
    if physical_damage:
        out["bonusPhysicalDamage"] = physical_damage
    mask = item["classMask"]
    if mask > 0:
        out["classAllowlist"] = [cls for bit, cls in CLASS_MASK_TO_PROTO_CLASS.items() if mask & bit]
        if not out["classAllowlist"]:
            raise ValueError(f"class mask {mask} excludes every supported class")
    if item["raceMasks"] not in ([-1, -1], [0, 0]):
        raise ValueError(f"race-restricted record needs a UI mapping: {item['raceMasks']}")
    source_factions = {source.get("faction", "") for source in item["sources"]}
    if source_factions in ({"Alliance"}, {"Horde"}):
        out["factionRestriction"] = 1 if "Alliance" in source_factions else 2
    if item["requiredSkill"]:
        out["requiredProfession"] = PROFESSIONS[item["requiredSkill"]]
    if item["maxCount"] == 1 or item.get("uniqueEquipped") or item.get("limitCategoryQuantity") == 1:
        out["unique"] = True
    if item["setId"]:
        out.update(setId=item["setId"], setName=item["setName"])
    sources = []
    for source in item["sources"]:
        if source["kind"] == "crafted":
            sources.append({"crafted": {
                "profession": PROFESSIONS[source["skill"]], "spellId": source["spellId"],
            }})
        elif source["kind"] == "dungeon":
            drop = {"difficulty": 1, "zoneId": source["zoneId"]}
            if source["entityType"] == 1:
                drop["npcId"] = source["entityId"]
            else:
                drop["otherName"] = source.get("name", "")
            sources.append({"drop": drop})
        elif source["kind"] in ("pvp", "vendor"):
            sources.append({"soldBy": {
                "npcName": source.get("name", "PvP quartermasters"),
                "npcId": source.get("entityId", 0),
                "zoneId": source.get("zoneId", 0),
            }})
        elif source["kind"] in ("dungeon-quest", "ticket-exchange"):
            sources.append({"quest": {"id": source["questId"], "name": source["name"]}})
        elif source["kind"] == "synthetic":
            # No fabricated vendor, crafting recipe or drop for modeled gear.
            continue
        else:
            raise ValueError(f"unsupported acquisition source {source['kind']}")
    out["sources"] = sources
    if vendor is not None:
        out["stats"] = build_stats(item, vendor)
        out["classAllowlist"] = class_allowlist(item, vendor)
        exported = vendor["item"]
        for source, dest in (("name", "name"), ("quality", "quality"), ("actualItemLevel", "ilvl")):
            if source in exported:
                out[dest] = exported[source]
        if cls == 2:
            _, _, _, _, _, low, high, speed = weapon_fields(item, vendor)
            out.update(weaponDamageMin=low, weaponDamageMax=high, weaponSpeed=speed)
    return out


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--catalog", type=Path, default=Path("assets/db_inputs/forever_gear_catalog.json"))
    parser.add_argument("--synthetic-catalog", type=Path,
                        default=Path("assets/db_inputs/forever_synthetic_gear.json"))
    parser.add_argument("--vendor-dir", type=Path, default=Path("assets/db_inputs/forever_vendor"))
    parser.add_argument("--out", type=Path, default=Path("assets/db_inputs/forever_equipment.json"))
    args = parser.parse_args()
    catalog = json.loads(args.catalog.read_text())
    vendors, _ = read_vendor_exports(args.vendor_dir)
    rates = load_level_60_coefficients()
    compiled = []
    for item in catalog["items"]:
        try:
            compiled.append(compile_item(item, rates, vendors.get(item["id"])))
        except (ValueError, KeyError) as error:
            print(f"Excluded {item['id']} {item['name']}: {error}")
    if args.synthetic_catalog.exists():
        synthetic = json.loads(args.synthetic_catalog.read_text())
        for item in synthetic["items"]:
            try:
                compiled.append(compile_item(item, rates))
            except (ValueError, KeyError) as error:
                raise ValueError(f"Modeled item {item['id']} cannot compile: {error}") from error
    args.out.parent.mkdir(parents=True, exist_ok=True)
    args.out.write_text(json.dumps({"items": compiled}, indent=2, ensure_ascii=False) + "\n")
    print(f"Compiled {len(compiled)} of {len(catalog['items'])} items into {args.out}")


if __name__ == "__main__":
    main()
