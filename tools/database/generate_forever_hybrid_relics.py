#!/usr/bin/env python3
"""Project every real Forever relic with a quarter of the modeled relic stats.

Relics are the one slot where the benchmark uses real items: each real idol, libram and totem keeps its own
scripted effect (the projected copy names it in effectItemId) and gains a quarter of the modeled level-65
relic's stats, in either the physical or the spell-power allocation. The search then picks the best relic
for each profile. The full-capacity modeled relics stay in the catalog but are no longer ranking gear.
"""
from __future__ import annotations

import json
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
REAL = ROOT / "assets/db_inputs/forever_gear_catalog.json"
MODELED = ROOT / "assets/db_inputs/forever_synthetic_gear_v2.json"
OUT = ROOT / "assets/db_inputs/forever_hybrid_relics.json"
FIRST_ID = 920001001
SHARE = 0.25
# Modeled level-65 relics by relic subclass (7 libram, 8 idol, 9 totem) and allocation.
MODELED_RELICS = {(8, "physical"): 920000499, (8, "spell"): 920000500, (7, "physical"): 920000501,
                  (7, "spell"): 920000502, (9, "physical"): 920000503, (9, "spell"): 920000504}
ABBREVIATIONS = {"Stamina": "Sta", "Agility": "Agi", "AttackPower": "AP", "Intellect": "Int", "SpellPower": "SP"}


def number(value: float) -> str:
    return f"{value:g}"


def main() -> None:
    real = [item for item in json.loads(REAL.read_text())["items"] if item["inventoryType"] == 28]
    modeled = {item["id"]: item for item in json.loads(MODELED.read_text())["items"]}
    items = []
    for relic in sorted(real, key=lambda item: item["id"]):
        for allocation in ("physical", "spell"):
            source = modeled[MODELED_RELICS[relic["subclass"], allocation]]
            stats = {name: value * SHARE for name, value in source["stats"].items()}
            label = ", ".join(f"+{number(value)} {ABBREVIATIONS[name]}" for name, value in stats.items())
            items.append({
                "id": FIRST_ID + len(items),
                "name": f"Projected: {relic['name']} \u2014 Relic [{label}]",
                "icon": relic["icon"], "itemLevel": 65, "requiredLevel": 60, "quality": relic["quality"],
                "class": relic["class"], "subclass": relic["subclass"], "inventoryType": 28,
                "classMask": relic["classMask"], "raceMasks": relic["raceMasks"],
                "requiredSkill": 0, "requiredSkillRank": 0, "requiredAbility": 0,
                "maxCount": 0, "limitCategory": 0, "setId": 0,
                "sources": [{"kind": "synthetic"}], "effects": [], "stats": stats, "armor": 0,
                "clientStatRecord": False, "statSource": "real-relic-effect-plus-quarter-modeled-stats",
                "synthetic": True, "modelVersion": 2, "benchmarkEligible": True,
                "effectItemId": relic["id"], "modelReferenceItemID": relic["id"],
                "modelSourceSlot": "Relic", "modelAllocation": allocation, "modelBudgetRatio": SHARE,
                "modelPolicy": "Real relic effect plus a quarter of the modeled relic stats",
                "modelConfidence": "unverified",
                "modelNotes": f"{relic['name']}'s own effect; a quarter of {source['name'].split(' \u2014 ')[0][9:]}'s stats.",
            })
    OUT.write_text(json.dumps({
        "schemaVersion": 1,
        "method": "Real relic effects with a quarter of the modeled level-65 relic stats (physical or spell allocation).",
        "share": SHARE, "items": items,
    }, indent=2, ensure_ascii=False) + "\n")
    print(f"{len(items)} projected relics from {len(real)} real relics")


if __name__ == "__main__":
    main()
