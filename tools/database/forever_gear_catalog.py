#!/usr/bin/env python3
"""Assemble current Forever crafting, dungeon and exported PvP item data.

ItemSparse allocations multiplied by RandPropPoints recover stats that Wowhead
does not render. These allocation multipliers are NOT stat exchange prices.
"""

from __future__ import annotations

import argparse
import collections
import hashlib
import json
import math
import os
from pathlib import Path
import sys

from import_forever_vendor import parse_planner, read_vendor_exports, shield_block_value

sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "data_watch"))
import spell_client


# Area IDs, not map IDs. Source phase/item level cannot distinguish raid tokens
# from current dungeon items, or new crafting from old crafting.
DUNGEONS = {
    209, 491, 717, 718, 719, 721, 722, 796, 1176, 1337, 1477,
    1581, 1583, 1584, 2017, 2057, 2100, 2437, 2557,
}
RAIDS = {1977, 2159, 2677, 2717, 3428, 3429, 3456}
CRAFT_SKILLS = {164, 165, 171, 197, 202, 333}
# Battleground quartermasters, Timbermaw, dungeon currency and open-world PvP.
# Raid-reputation vendors such as Anachronos are deliberately not in this list.
PVP_VENDORS = {15127, 15126, 14754, 14753, 263570}
NONRAID_VENDORS = PVP_VENDORS | {11557, 227853}
VENDOR_FACTIONS = {15127: "Alliance", 14753: "Alliance", 15126: "Horde", 14754: "Horde"}
DARKMOON_TICKET_REWARDS = {7940, 7981}
SLOT_POINTS = {
    **dict.fromkeys((1, 5, 7, 17, 20), 0),
    **dict.fromkeys((3, 6, 8, 10, 12), 1),
    **dict.fromkeys((2, 9, 11, 14, 16, 23), 2),
    **dict.fromkeys((13, 21, 22), 3),
    **dict.fromkeys((15, 25, 26, 28), 4),
}
QUALITY_POINTS = {2: "Good", 3: "Superior", 4: "Epic", 5: "Epic"}

# Values remain raw client units. In particular HitRating/CritRating are not %.
STAT_NAMES = {
    0: "Mana", 1: "Health", 3: "Agility", 4: "Strength", 5: "Intellect",
    6: "Spirit", 7: "Stamina", 12: "DefenseRating", 13: "DodgeRating",
    14: "ParryRating", 15: "BlockRating", 16: "MeleeHitRating",
    17: "RangedHitRating", 18: "SpellHitRating", 19: "MeleeCritRating",
    20: "RangedCritRating", 21: "SpellCritRating", 31: "HitRating",
    32: "CritRating", 35: "ResilienceRating", 36: "HasteRating",
    37: "ExpertiseRating", 38: "AttackPower", 39: "RangedAttackPower",
    40: "FeralAttackPower", 41: "HealingPower", 42: "SpellDamage",
    43: "MP5", 44: "ArmorPenetrationRating", 45: "SpellPower",
    46: "HealthRegen", 47: "SpellPenetration", 48: "BlockValue",
    50: "BonusArmor", 51: "FireResistance", 52: "FrostResistance",
    53: "HolyResistance", 54: "ShadowResistance", 55: "NatureResistance",
    56: "ArcaneResistance", 83: "BonusPhysicalDamage",
    84: "HolyPower", 85: "FirePower", 86: "NaturePower", 87: "FrostPower",
    88: "ShadowPower", 89: "ArcanePower",
}
# Beta GlobalStrings 58882..58912 orders these new skill modifiers; examples
# include Huge Thorium Battleaxe (90) and Servomechanic Sledgehammer (91).
STAT_NAMES.update(dict(enumerate((
    "TwoHandedAxesSkill", "TwoHandedMacesSkill", "TwoHandedSwordsSkill",
    "AxesSkill", "BowsSkill", "CrossbowsSkill", "DaggersSkill", "DualWieldSkill",
    "FistWeaponsSkill", "GunsSkill", "MacesSkill", "PolearmsSkill", "StavesSkill",
    "SwordsSkill", "ThrownSkill", "WandsSkill",
), 90)))
STAT_NAMES.update(dict(enumerate((
    "AlchemySkill", "BlacksmithingSkill", "EnchantingSkill", "EngineeringSkill",
    "JewelcraftingSkill", "LeatherworkingSkill", "HerbalismSkill", "MiningSkill",
    "SkinningSkill", "CookingSkill", "FirstAidSkill", "FishingSkill", "TailoringSkill",
), 106)))

PLANNER_STATS = {
    0: "mana", 1: "health", 3: "agi", 4: "str", 5: "int", 6: "spi", 7: "sta",
    12: "defrtng", 13: "dodgertng", 14: "parryrtng", 15: "blockrtng",
    31: "hitrtng", 32: "critstrkrtng", 35: "resirtng", 36: "hastertng",
    37: "exprtng", 38: "atkpwr", 39: "rgdatkpwr", 41: "healpwr",
    42: "spldmg", 43: "manargn", 45: "splpwr", 47: "splpen",
    48: "blockvalue", 50: "armorbonus", 51: "firres", 52: "frores",
    54: "shares", 55: "natres", 56: "arcres",
}


def acquisition(item: dict, vendor_export: dict | None = None) -> tuple[list[dict], str | None]:
    details = item.get("sourcemore", [])
    # Quest category c can expose raid-token rewards hidden by a vendor label.
    if any(s.get("z") in RAIDS or s.get("c") in RAIDS for s in details):
        return [], "raid-derived"
    sources = []
    for s in details:
        if s.get("t") == 6 and s.get("s") in CRAFT_SKILLS and s.get("ti"):
            sources.append({"kind": "crafted", "skill": s["s"], "spellId": s["ti"]})
        elif s.get("t") == 5 and s.get("c") in DUNGEONS:
            sources.append({
                "kind": "dungeon-quest", "zoneId": s["c"],
                "questId": s["ti"], "name": s.get("n", ""),
            })
        elif s.get("t") == 5 and s.get("ti") in DARKMOON_TICKET_REWARDS:
            sources.append({
                "kind": "ticket-exchange", "questId": s["ti"], "name": s.get("n", ""),
            })
        elif s.get("t") in (None, 1, 2) and s.get("z") in DUNGEONS:
            sources.append({
                "kind": "dungeon", "zoneId": s["z"], "entityId": s.get("ti", 0),
                "entityType": s.get("t", 0), "name": s.get("n", "Dungeon drop"),
            })
        elif 5 in item.get("source", []) and s.get("t") == 1 and s.get("ti") in NONRAID_VENDORS:
            sources.append({
                "kind": "pvp" if s["ti"] in PVP_VENDORS else "vendor",
                "entityId": s["ti"], "name": s.get("n", ""), "zoneId": s.get("z", 0),
                "faction": VENDOR_FACTIONS.get(s["ti"], ""),
            })
    if vendor_export and any(c.get("currencyName") == "Honor Points" for c in vendor_export.get("costs", [])):
        sources.append({"kind": "pvp", "exported": True})
    return sources, None if sources else "no-crafting-or-dungeon-source"


def item_stats(sparse: dict, points: dict) -> tuple[dict, list[dict], float]:
    quality = int(sparse["OverallQualityID"])
    slot = int(sparse["InventoryType"])
    prefix = QUALITY_POINTS[quality]
    scale = float(points[sparse["ItemLevel"]][f"{prefix}F_{SLOT_POINTS[slot]}"])
    values = collections.defaultdict(int)
    allocations = []
    for n in range(10):
        mod = int(sparse.get(f"StatModifier_bonusStat_{n}", -1))
        allocation = int(sparse.get(f"StatPercentEditor_{n}", 0))
        if mod < 0 or not allocation:
            continue
        value = math.floor(scale * allocation / 10000 + .5)
        name = STAT_NAMES.get(mod, f"UnmappedItemMod{mod}")
        values[name] += value
        allocations.append({"mod": mod, "allocation": allocation, "value": value})
    return dict(values), allocations, scale


def validate_planner(raw: dict, allocations: list[dict]) -> list[dict]:
    conflicts = []
    for a in allocations:
        key = PLANNER_STATS.get(a["mod"])
        if key and key in raw and raw[key] != a["value"]:
            conflicts.append({"stat": key, "client": a["value"], "planner": raw[key]})
    return conflicts


def planner_record(item: dict, sources: list[dict]) -> dict:
    raw = item.get("stats", {})
    names = {key: STAT_NAMES[mod] for mod, key in PLANNER_STATS.items()}
    names.update({
        "arcsplpwr": "ArcanePower", "firsplpwr": "FirePower", "frosplpwr": "FrostPower",
        "holsplpwr": "HolyPower", "natsplpwr": "NaturePower", "shasplpwr": "ShadowPower",
        "splheal": "HealingPower", "feratkpwr": "FeralAttackPower",
    })
    stats = {name: raw[key] for key, name in names.items() if raw.get(key)}
    if not stats and item["class"] != 2:
        raise ValueError("no published combat stats; requires tooltip/effect review")
    result = {
        "id": item["id"], "name": item["name"], "icon": item.get("icon", ""),
        "itemLevel": item["itemLevel"], "quality": item["quality"],
        "class": item["class"], "subclass": item["subclass"],
        "inventoryType": item["inventoryType"],
        "requiredLevel": item.get("requiredLevel", 0),
        "classMask": item.get("classMask", 0),
        "raceMasks": [item.get("raceMask", 0), 0],
        "requiredSkill": raw.get("reqskill", 0),
        "requiredSkillRank": raw.get("reqskillrank", 0),
        "maxCount": raw.get("maxcount", 0), "limitCategory": item.get("limitCategory", 0),
        "setId": raw.get("itemset", 0), "sources": sources,
        "stats": stats, "armor": raw.get("armor", 0),
        "statSource": "wowhead-forever", "clientStatRecord": False,
    }
    add_weapon(result, item)
    return result


def add_weapon(record: dict, item: dict) -> None:
    raw = item.get("stats", {})
    if item["class"] == 2:
        if not all(k in raw for k in ("dmgmin1", "dmgmax1", "speed")):
            raise ValueError("missing weapon damage/speed")
        record["weapon"] = {
            "min": raw["dmgmin1"], "max": raw["dmgmax1"], "speed": raw["speed"],
            "school": raw.get("dmgtype1", 0),
        }
    if item.get("randomEnchants"):
        record["randomSuffixes"] = item["randomEnchants"]


def make_record(item: dict, sparse: dict, base: dict, points: dict, sources: list[dict]) -> dict:
    stats, allocations, scale = item_stats(sparse, points)
    raw = item.get("stats", {})
    conflicts = validate_planner(raw, allocations)
    if conflicts:
        raise ValueError(f"client/planner stat conflict: {conflicts}")
    result = {
        "id": item["id"], "name": sparse["Display_lang"], "icon": item.get("icon", ""),
        "itemLevel": int(sparse["ItemLevel"]),
        "quality": int(sparse["OverallQualityID"]),
        "class": int(base["ClassID"]), "subclass": int(base["SubclassID"]),
        "inventoryType": int(sparse["InventoryType"]),
        "requiredLevel": int(sparse["RequiredLevel"]),
        "classMask": int(sparse["AllowableClass"]),
        "raceMasks": [int(sparse["AllowableRace_0"]), int(sparse["AllowableRace_1"])],
        "requiredSkill": int(sparse["RequiredSkill"]),
        "requiredSkillRank": int(sparse["RequiredSkillRank"]),
        "requiredAbility": int(sparse["RequiredAbility"]),
        "maxCount": int(sparse["MaxCount"]), "limitCategory": int(sparse["LimitCategory"]),
        "uniqueEquipped": bool(int(sparse["Flags_0"]) & 0x80000),
        "binding": int(sparse["Bonding"]), "setId": int(sparse["ItemSet"]),
        "sources": sources, "stats": stats,
        "armor": raw.get("armor", 0),
        "statScale": scale, "statAllocations": allocations,
        "statSource": "client", "clientStatRecord": True,
    }
    add_weapon(result, item)
    return result


def build_catalog(planner: dict, tables: dict, vendors: dict | None = None) -> dict:
    vendors = vendors or {}
    sparse = {r["ID"]: r for r in tables["ItemSparse"]}
    base = {r["ID"]: r for r in tables["Item"]}
    points = {r["ID"]: r for r in tables["RandPropPoints"]}
    effects = {r["ID"]: r for r in tables["ItemEffect"]}
    item_effects = collections.defaultdict(list)
    for r in tables["ItemXItemEffect"]:
        item_effects[r["ItemID"]].append(r["ItemEffectID"])
    set_effects = collections.defaultdict(list)
    for r in tables["ItemSetSpell"]:
        set_effects[r["ItemSetID"]].append({
            "pieces": int(r["Threshold"]), "spellId": int(r["SpellID"]),
            "specialization": int(r["ChrSpecID"]), "traitSubTree": int(r["TraitSubTreeID"]),
        })
    sets = {r["ID"]: r for r in tables["ItemSet"]}
    limits = {r["ID"]: r for r in tables["ItemLimitCategory"]}
    admitted, unavailable, excluded = [], [], collections.Counter()
    for id, item in sorted(planner.items(), key=lambda kv: int(kv[0])):
        if (item.get("requiredLevel", 0) > 60 or item["quality"] not in QUALITY_POINTS
                or item.get("inventoryType") not in SLOT_POINTS or item["class"] not in (2, 4)):
            continue
        vendor = vendors.get(int(id))
        sources, reason = acquisition(item, vendor)
        if reason:
            excluded[reason] += 1
            continue
        row = sparse.get(id)
        if row and (int(row["RequiredLevel"]) > 60 or int(row["OverallQualityID"]) not in QUALITY_POINTS
                or int(row["InventoryType"]) not in SLOT_POINTS
                or int(base[id]["ClassID"]) not in (2, 4)):
            continue
        if row and (int(row["RequiredPVPRank"]) or int(row["RequiredPVPMedal"])) and not any(
                s["kind"] == "pvp" for s in sources):
            excluded["pvp-requirement"] += 1
            continue
        try:
            # Static DB2 is not the server's complete item database. Current
            # dungeon items such as Shadowcraft Cap are absent from ItemSparse;
            # that does not make their published Forever values unavailable.
            record = make_record(item, row, base[id], points, sources) if row else planner_record(item, sources)
            missing = [e for e in item_effects[id] if e not in effects]
            if missing:
                raise ValueError(f"missing ItemEffect records: {missing}")
            if record["limitCategory"]:
                limit = limits[str(record["limitCategory"])]
                record["limitCategoryName"] = limit["Name_lang"]
                record["limitCategoryQuantity"] = int(limit["Quantity"])
                record["limitCategoryFlags"] = int(limit["Flags"])
        except (ValueError, KeyError) as e:
            unavailable.append({"id": int(id), "name": item["name"], "reason": str(e)})
            continue
        record["effects"] = [
            {k: int(e[k]) for k in ("SpellID", "TriggerType", "Charges", "CoolDownMSec",
                                     "CategoryCoolDownMSec", "SpellCategoryID", "PlayerConditionID")}
            for e in sorted((effects[e] for e in item_effects[id]), key=lambda e: int(e["LegacySlotIndex"]))
        ]
        record["effectReviewRequired"] = bool(record["effects"]) or bool(record["setId"]) or not row
        if vendor:
            record["vendorExport"] = True
            record["vendorTooltip"] = [line["left"] for line in vendor["tooltip"]["lines"] if line.get("left")]
        set_id = str(record["setId"])
        if set_id in sets:
            record["setName"] = sets[set_id]["Name_lang"]
            record["setBonuses"] = sorted(set_effects[set_id], key=lambda e: e["pieces"])
            record["setRequiredSkill"] = int(sets[set_id]["RequiredSkill"])
            record["setRequiredSkillRank"] = int(sets[set_id]["RequiredSkillRank"])
        admitted.append(record)
    return {"items": admitted, "unavailable": unavailable, "exclusionCounts": dict(excluded)}


TABLES = ("Item", "ItemSparse", "RandPropPoints", "ItemEffect", "ItemXItemEffect",
          "ItemSet", "ItemSetSpell", "ItemLimitCategory", "GlobalStrings")


def add_shield_block_values(items: list[dict], evidence: list[dict], vendors: dict) -> None:
    captured = {row["id"]: row for row in evidence}
    for item in items:
        if item["inventoryType"] != 14:
            continue
        source = captured.get(item["id"])
        if source:
            item["baseBlockValue"] = source["baseBlockValue"]
            item["baseBlockSource"] = source["source"]
        vendor = vendors.get(item["id"])
        value = shield_block_value(vendor) if vendor else None
        if value is not None:
            item["baseBlockValue"] = value
            item["baseBlockSource"] = "Forever vendor export: intrinsic Block tooltip line"


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--planner", type=Path, default=Path("assets/db_inputs/wowhead_forever_gearplanner.txt"))
    parser.add_argument("--vendor-dir", type=Path, default=Path("assets/db_inputs/forever_vendor"))
    parser.add_argument("--build", default=spell_client.FOREVER)
    parser.add_argument("--cache", type=Path, default=Path(os.environ.get("XDG_CACHE_HOME", Path.home() / ".cache")) / "wowsims-forever")
    parser.add_argument("--out", type=Path, default=Path("assets/db_inputs/forever_gear_catalog.json"))
    parser.add_argument("--reviews", type=Path, default=Path("assets/db_inputs/forever_gear_reviews.json"))
    parser.add_argument("--shield-block", type=Path,
                        default=Path("assets/db_inputs/forever_shield_block.json"))
    args = parser.parse_args()
    spell_client.CACHE = str(args.cache)
    tables = {name: spell_client.table(args.build, name) for name in TABLES}
    vendors, _ = read_vendor_exports(args.vendor_dir)
    planner = parse_planner(args.planner)
    pool_path = Path("assets/db_inputs/forever_ilvl65_items.json")
    if pool_path.exists():
        for item in json.loads(pool_path.read_text())["items"]:
            existing = planner.get(str(item["id"]))
            if existing is not None:
                for field in ("source", "sourcemore"):
                    if field in item:
                        existing[field] = item[field]
    catalog = build_catalog(planner, tables, vendors)
    add_shield_block_values(catalog["items"], json.loads(args.shield_block.read_text()), vendors)
    reviews = json.loads(args.reviews.read_text())["items"]
    for item in catalog["items"]:
        review = reviews.get(str(item["id"]), {})
        if review.get("statOverrides"):
            item["stats"].update(review["statOverrides"])
            item["statSupplementSource"] = review["source"]
    catalog = {
        "schemaVersion": 2, "clientBuild": args.build,
        "plannerSHA256": hashlib.sha256(args.planner.read_bytes()).hexdigest(),
        "ilvl65ListSHA256": hashlib.sha256(pool_path.read_bytes()).hexdigest() if pool_path.exists() else None,
        "vendorExports": {
            path.name: hashlib.sha256(path.read_bytes()).hexdigest()
            for path in sorted(args.vendor_dir.glob("*.txt"))
        },
        "clientTables": {
            name: {
                "url": f"https://wago.tools/db2/{name}/csv?build={args.build}",
                "sha256": hashlib.sha256((args.cache / f"{name}_{args.build}.csv").read_bytes()).hexdigest(),
            } for name in TABLES
        },
        **catalog,
    }
    args.out.parent.mkdir(parents=True, exist_ok=True)
    args.out.write_text(json.dumps(catalog, indent=2, ensure_ascii=False) + "\n")
    print(f"{len(catalog['items'])} items; {len(catalog['unavailable'])} candidates unresolved")
    print("Slots:", dict(sorted(collections.Counter(i["inventoryType"] for i in catalog["items"]).items())))


if __name__ == "__main__":
    main()
