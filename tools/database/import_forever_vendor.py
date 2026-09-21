#!/usr/bin/env python3
"""Build the curated Forever PvP-vendor item input used by the database generator.

Vendor exports override planner stats, restrictions and weapon damage. Other
level-60 class armor remains provisional: its planner values use the same
rating conversions established by the exported armor.
"""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path
from typing import Any

from import_forever_ratings import load_level_60_coefficients


STAT_COUNT = 44
RATING_COEFFICIENTS = load_level_60_coefficients()

# Generic item hit/crit go into one pool. EquipStats shares that pool with
# spells; populating both fields here would double the item's contribution.
PLANNER_RATINGS = {
    "hitrtng": (18, "Hit - Melee"),
    "critstrkrtng": (19, "Crit - Melee"),
    "dodgertng": (31, "Dodge"),
    "parryrtng": (32, "Parry"),
    "blockrtng": (29, "Block"),
    "defrtng": (28, "Defense Skill"),
}

# UI/proto stat indexes.  These intentionally mirror proto/common.proto.
STAT = {
    "str": 0,
    "agi": 1,
    "sta": 2,
    "int": 3,
    "spi": 4,
    "splpwr": 5,
    "arcsplpwr": 6,
    "firsplpwr": 7,
    "frosplpwr": 8,
    "holysplpwr": 9,
    "natsplpwr": 10,
    "shdwsplpwr": 11,
    "manargn": 12,
    "splpen": 16,
    "atkpwr": 17,
    "rgdatkpwr": 27,
    "blockvalue": 30,
    "health": 34,
    "arcres": 35,
    "firres": 36,
    "frosres": 37,
    "natres": 38,
    "shadres": 39,
    "armorbonus": 40,
    "healpwr": 41,
    "spldmg": 42,
    "feralap": 43,
}

CLASS_MASK_TO_PROTO_CLASS = {
    1: 9,  # Warrior
    2: 4,  # Paladin
    4: 2,  # Hunter
    8: 6,  # Rogue
    16: 5,  # Priest
    64: 7,  # Shaman
    128: 3,  # Mage
    256: 8,  # Warlock
    1024: 1,  # Druid
}

INVENTORY_TO_TYPE = {
    1: 1,  # Head
    2: 2,  # Neck
    3: 3,  # Shoulder
    5: 5,  # Chest
    6: 8,  # Waist
    7: 9,  # Legs
    8: 10,  # Feet
    9: 6,  # Wrist
    10: 7,  # Hands
    11: 11,  # Finger
    12: 12,  # Trinket
    16: 4,  # Back
    20: 5,  # Robe
}

WEAPON_TYPES = {
    0: 1,  # Axe
    1: 1,  # Two-handed axe
    4: 4,  # Mace
    5: 4,  # Two-handed mace
    6: 6,  # Polearm
    7: 9,  # Sword
    8: 9,  # Two-handed sword
    10: 8,  # Staff
    13: 3,  # Fist
    15: 2,  # Dagger
}

EXPORTED_STATS = {
    "ITEM_MOD_STRENGTH_SHORT": 0,
    "ITEM_MOD_AGILITY_SHORT": 1,
    "ITEM_MOD_STAMINA_SHORT": 2,
    "ITEM_MOD_INTELLECT_SHORT": 3,
    "ITEM_MOD_SPIRIT_SHORT": 4,
    "ITEM_MOD_SPELL_POWER_SHORT": 5,
    "ITEM_MOD_MANA_REGENERATION_SHORT": 12,
    "ITEM_MOD_ATTACK_POWER_SHORT": 17,
    "RESISTANCE0_NAME": 26,
    "ITEM_MOD_SPELL_HEALING_DONE_SHORT": 41,
    "ITEM_MOD_SPELL_DAMAGE_DONE_SHORT": 42,
}

CLASS_NAMES = {
    "Druid": 1, "Hunter": 2, "Mage": 3, "Paladin": 4, "Priest": 5,
    "Rogue": 6, "Shaman": 7, "Warlock": 8, "Warrior": 9,
}

FAMILY_PREFIXES = (
    "Shadowhide",
    "Felweave",
    "Mortarplate",
    "Lunarhide",
    "Dragonhide",
    "Chain",
    "Magus",
    "Satin",
    "Mail",
    "Linked",
    "Chevalier",
)

FOREVER_SET_NAMES = {
    2081: "Lieutenant Commander's Wildhide",
    2086: "Lieutenant Commander's Vindication",
    2094: "Champion's Thunderfist",
}

ARMOR_SLOTS = {1, 3, 5, 6, 7, 8, 9, 10}
PLANNER_ARMOR_SLOTS = ARMOR_SLOTS | {20}


def number(value: Any) -> float:
    if isinstance(value, bool) or not isinstance(value, (int, float)):
        return 0
    return value


def parse_planner(path: Path) -> dict[str, dict[str, Any]]:
    text = path.read_text()
    marker = 'WH.setPageData("wow.gearPlanner.classicplus.item"'
    start = text.index("{", text.index(marker))
    end = text.index(");", start)
    payload = re.sub(r",\s*([}\]])", r"\1", text[start:end])
    return json.loads(payload)


def read_vendor_exports(path: Path) -> tuple[dict[int, dict[str, Any]], dict[int, str]]:
    source_items: dict[int, dict[str, Any]] = {}
    set_names: dict[int, str] = {}

    for source_path in sorted(path.glob("*.txt")):
        data = json.loads(source_path.read_text())
        for wrapper in data["items"]:
            item_id = int(wrapper["itemID"])
            source_items[item_id] = wrapper
            for line in wrapper.get("tooltip", {}).get("lines", []):
                left = str(line.get("left", "")).strip()
                match = re.match(r"(.+)\s+\(0/\d+\)$", left)
                if match:
                    set_names[item_id] = match.group(1)
                    break

    return source_items, set_names


def existing_set_names(db_path: Path) -> dict[int, str]:
    if not db_path.exists():
        return {}
    data = json.loads(db_path.read_text())
    result: dict[int, str] = {}
    for item in data.get("items", []):
        set_id = int(item.get("setId", 0))
        set_name = item.get("setName", "")
        if set_id and set_name and set_id not in result:
            result[set_id] = set_name
    return result


def selected_ids(
    planner: dict[str, dict[str, Any]],
    source_items: dict[int, dict[str, Any]],
) -> set[int]:
    selected = {
        item_id
        for item_id, wrapper in source_items.items()
        if not wrapper["item"].get("equipLocation", "").startswith("INVTYPE_NON")
    }

    for family in FAMILY_PREFIXES:
        family_items = [
            item
            for item in planner.values()
            if item.get("itemLevel") == 60
            and str(item.get("name", "")).startswith(f"Premier {family} ")
            and int(item.get("inventoryType", 0)) in PLANNER_ARMOR_SLOTS
        ]
        by_slot: dict[int, list[dict[str, Any]]] = {}
        for item in family_items:
            inventory_type = int(item["inventoryType"])
            slot = 5 if inventory_type == 20 else inventory_type
            by_slot.setdefault(slot, []).append(item)

        missing = ARMOR_SLOTS - by_slot.keys()
        if missing:
            raise ValueError(f"{family} is missing armor slots: {sorted(missing)}")
        for slot_items in by_slot.values():
            if len(slot_items) != 1:
                names = ", ".join(item["name"] for item in slot_items)
                raise ValueError(f"{family} has ambiguous slot data: {names}")
            selected.add(int(slot_items[0]["id"]))

    return selected


def tooltip_lines(source_item: dict[str, Any]) -> list[str]:
    return [str(line.get("left", "")).strip()
            for line in source_item.get("tooltip", {}).get("lines", [])]


def tooltip_percent(source_item: dict[str, Any], pattern: str) -> float:
    for line in tooltip_lines(source_item):
        match = re.search(pattern, line)
        if match:
            return float(match.group(1))
    raise ValueError(f"Missing percentage tooltip for item {source_item['itemID']}")


def build_stats(
    planner_item: dict[str, Any], source_item: dict[str, Any] | None = None,
) -> list[float]:
    stats = [0.0] * STAT_COUNT
    raw = planner_item.get("stats", {})

    if source_item is not None:
        exported = source_item["item"].get("stats", {})
        for key, index in EXPORTED_STATS.items():
            stats[index] = number(exported.get(key))
        for line in tooltip_lines(source_item):
            bonus_armor = re.fullmatch(r"\+(\d+) Armor", line)
            if bonus_armor:
                stats[40] += int(bonus_armor[1])
        stats[26] -= stats[40]
        if exported.get("ITEM_MOD_CRIT_RATING_SHORT"):
            stats[19] = tooltip_percent(source_item, r"critical strike by ([\d.]+)%")
        if exported.get("ITEM_MOD_DODGE_RATING_SHORT"):
            stats[31] = tooltip_percent(source_item, r"Dodge an attack by ([\d.]+)%")
    else:
        for key, index in STAT.items():
            stats[index] = number(raw.get(key))
        stats[26] = number(raw.get("armor"))
        # Client GameTables establish these coefficients; the exported crit
        # and dodge percentages independently agree.
        for key, (index, coefficient) in PLANNER_RATINGS.items():
            stats[index] = number(raw.get(key)) / RATING_COEFFICIENTS[coefficient]
        # Haste does not pass through the hit/crit sharing operation.
        stats[20] = number(raw.get("hastertng")) / RATING_COEFFICIENTS["Haste - Melee"]
        stats[15] = number(raw.get("hastertng")) / RATING_COEFFICIENTS["Haste - Spell"]
        unsupported = {key for key in raw if key.endswith("rtng")
                       and raw[key] and key not in PLANNER_RATINGS and key != "hastertng"}
        if unsupported:
            raise ValueError(f"Unsupported rating conversions: {sorted(unsupported)}")

    stats[27] += stats[17]
    return stats


def class_allowlist(
    planner_item: dict[str, Any], source_item: dict[str, Any] | None,
) -> list[int]:
    if source_item is not None:
        for line in tooltip_lines(source_item):
            if line.startswith("Classes: "):
                return [CLASS_NAMES[name.strip()] for name in line[9:].split(",")]
        return []
    raw = planner_item.get("classMask")
    if raw is None:
        raw = planner_item.get("stats", {}).get("classes", 0)
    mask = int(number(raw))
    return [
        proto_class
        for bit, proto_class in CLASS_MASK_TO_PROTO_CLASS.items()
        if mask & bit
    ]


def weapon_fields(
    planner_item: dict[str, Any],
    source_item: dict[str, Any] | None,
) -> tuple[int, int, int, int, list[float], float, float, float]:
    item_class = int(planner_item.get("class", 0))
    subclass = int(planner_item.get("subclass", 0))
    inventory_type = int(planner_item.get("inventoryType", 0))

    if item_class == 4 and inventory_type == 14:
        return 13, 7, 3, 0, [0.0] * 16, 0, 0, 0
    if item_class == 4 and inventory_type == 23 and subclass == -5:
        return 13, 5, 3, 0, [0.0] * 16, 0, 0, 0

    if item_class != 2:
        return 0, 0, 0, 0, [0.0] * 16, 0, 0, 0

    ranged_type = {2: 1, 18: 2, 3: 3, 16: 6, 19: 8}.get(subclass, 0)
    item_type = 14 if ranged_type else 13
    weapon_type = WEAPON_TYPES.get(subclass, 0)

    if item_type == 14:
        hand_type = 0
    elif inventory_type == 17:
        hand_type = 4
    elif inventory_type == 21:
        hand_type = 1
    elif inventory_type == 22:
        hand_type = 3
    else:
        hand_type = 2

    if source_item is None:
        raise ValueError(f"Weapon {planner_item['id']} needs exported damage and speed")
    for line in source_item["tooltip"]["lines"]:
        damage = re.fullmatch(r"([\d.]+) - ([\d.]+) Damage", line.get("left", ""))
        speed = re.fullmatch(r"Speed ([\d.]+)", line.get("right", ""))
        if damage and speed:
            return (item_type, weapon_type, hand_type, ranged_type, [0.0] * 16,
                    float(damage[1]), float(damage[2]), float(speed[1]))
    raise ValueError(f"Missing weapon tooltip for item {planner_item['id']}")


def item_to_proto(
    planner_item: dict[str, Any],
    source_item: dict[str, Any] | None,
    source_set_name: str | None,
    old_set_names: dict[int, str],
) -> dict[str, Any]:
    item_id = int(planner_item["id"])
    inventory_type = int(planner_item.get("inventoryType", 0))
    item_class = int(planner_item.get("class", 0))
    subclass = int(planner_item.get("subclass", 0))
    item_type, weapon_type, hand_type, ranged_type, skills, damage_min, damage_max, speed = weapon_fields(
        planner_item, source_item
    )

    if item_type == 0:
        item_type = INVENTORY_TO_TYPE.get(inventory_type, 0)

    armor_type = 0
    if item_class == 4 and subclass in (1, 2, 3, 4):
        armor_type = subclass

    raw_stats = planner_item.get("stats", {})
    source = source_item["item"] if source_item is not None else {}
    set_id = int(source.get("setID") or raw_stats.get("itemset") or 0)

    result: dict[str, Any] = {
        "id": item_id,
        "name": source.get("name", planner_item.get("name", "")),
        "icon": planner_item.get("icon", "inv_misc_questionmark"),
        "type": item_type,
        "stats": build_stats(planner_item, source_item),
        "ilvl": int(source.get("actualItemLevel", planner_item.get("itemLevel", 0))),
        "phase": 1,
        "quality": int(source.get("quality", planner_item.get("quality", 0))),
        "expansion": 1,
        "factionRestriction": 2,
    }

    if armor_type:
        result["armorType"] = armor_type
    if weapon_type:
        result["weaponType"] = weapon_type
    if hand_type:
        result["handType"] = hand_type
    if ranged_type:
        result["rangedWeaponType"] = ranged_type
    if any(skills):
        result["weaponSkills"] = skills
    if damage_min:
        result["weaponDamageMin"] = damage_min
    if damage_max:
        result["weaponDamageMax"] = damage_max
    if speed:
        result["weaponSpeed"] = speed

    allowlist = class_allowlist(planner_item, source_item)
    if allowlist:
        result["classAllowlist"] = allowlist

    if set_id:
        result["setId"] = set_id
        result["setName"] = source_set_name or old_set_names.get(set_id, "")
        if not result["setName"]:
            result.pop("setName")

    return result


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--planner", type=Path, default=Path("assets/db_inputs/wowhead_forever_gearplanner.txt"))
    parser.add_argument("--vendor-dir", type=Path, default=Path("assets/db_inputs/forever_vendor"))
    parser.add_argument("--db", type=Path, default=Path("assets/database/db.json"))
    parser.add_argument("--output", type=Path, default=Path("assets/db_inputs/forever_vendor_items.json"))
    args = parser.parse_args()

    planner = parse_planner(args.planner)
    source_items, source_set_names = read_vendor_exports(args.vendor_dir)
    old_set_names = existing_set_names(args.db)
    ids = selected_ids(planner, source_items)

    missing = sorted(item_id for item_id in ids if str(item_id) not in planner)
    if missing:
        raise ValueError(f"Selected vendor items missing from gearplanner snapshot: {missing}")

    items = []
    for item_id in sorted(ids):
        planner_item = planner[str(item_id)]
        source = source_items.get(item_id)
        set_id = int(number(planner_item.get("stats", {}).get("itemset")))
        set_name = source_set_names.get(item_id)
        if not set_name and set_id:
            set_name = old_set_names.get(set_id) or FOREVER_SET_NAMES.get(set_id)
        items.append(item_to_proto(planner_item, source, set_name, old_set_names))

    args.output.parent.mkdir(parents=True, exist_ok=True)
    args.output.write_text(json.dumps({"items": items}, indent=2) + "\n")
    verified = len(ids & source_items.keys())
    print(f"Wrote {len(items)} items to {args.output}: "
          f"{verified} exported, {len(items) - verified} provisional class armor")


if __name__ == "__main__":
    main()
