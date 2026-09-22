#!/usr/bin/env python3
"""Capture client records supporting the mechanics and history review."""
import argparse
import hashlib
import json
from pathlib import Path
import sys

ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "tools/data_watch"))
import spell_client as client

ENCHANTS = (2505, 2566, 2584, 2590, 2591, 2604, 2605, 2617, 2715, 7603)
MANA_ITEMS = (3827, 6149, 4381, 12662, 13443, 13444)
SHAMAN_SPELLS = (16164, 16246, 16268, 1238931, 408498, 408505, 425336,
                 8232, 8233, 16362, 16361, 16342, 408341, 408345)
PRIEST_SPELLS = (15473, 15237, 27801, 10901, 10916, 10929, 10060, 14751, 1309957,
                 15286, 19280, 10797, 19296, 19299, 19302, 19303, 19304, 19305, 20591,
                 15270, 15335, 15336, 15337, 15338)
PROJECTILE_SPELLS = (75, 3044, 14287, 19434, 20904, 2643, 14280, 1978, 25295,
                    1310687, 1310785, 1310786, 686, 25307, 412758, 1293812, 1293813)
PALADIN_COOLDOWNS = (1311606, 20473, 20929, 20930, 879, 5614, 5615, 10312,
                     10313, 10314, 2812, 10318, 20925, 20927, 20928)
BUFF_ACCESS = (20218, 14752, 14818, 14819, 27841)
WARLOCK_COOLDOWNS = (6353, 17924, 6789, 17925, 17926, 17877, 18867, 18868,
                     18869, 18870, 18871, 1293817, 1293818, 17962, 18930, 18931, 18932)
FIRESTONE = (6366, 17951, 17952, 17953, 758, 17945, 17947, 17949,
             23480, 23481, 23482, 23483)
SPELLSTONE = (2362, 17727, 17728, 1237152, 1237156, 1237159, 1237162, 1237164, 1237165)
DEMONIC_BRAND = (1293695, 1293696, 1293697, 1293698, 1293699)
DRUID_HISTORY = (16880, 16886, 417141, 33876, 407993, 407995, 1238069, 1238070,
                 1238073, 414644, 1235826, 1235827, 5176, 2912, 8921, 5570, 16914, 9907, 768, 24858)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--cache", required=True)
    parser.add_argument("--build", default=client.FOREVER)
    parser.add_argument("--output", type=Path, default=ROOT / "assets/db_inputs/forever_effect_audit.json")
    args = parser.parse_args()
    client.CACHE = args.cache
    names = ("SpellItemEnchantment", "SpellEffect", "SpellAuraOptions", "Spell",
             "SpellLevels", "SpellMisc", "SpellDuration", "SpellClassOptions",
             "SpellShapeshift", "SpellPower", "SpellCooldowns", "ItemEffect", "ItemXItemEffect",
             "SpellCategories", "SpellEquippedItems", "SkillLineAbility", "Talent", "TraitDefinition",
             "SpellXDescriptionVariables", "SpellDescriptionVariables")
    tables = {name: client.table(args.build, name) for name in names}
    mana_links = [r for r in tables["ItemXItemEffect"] if int(r["ItemID"]) in MANA_ITEMS]
    mana_effect_ids = {r["ItemEffectID"] for r in mana_links}
    mana_effects = [r for r in tables["ItemEffect"] if r["ID"] in mana_effect_ids]
    mana_spell_ids = {r["SpellID"] for r in mana_effects}
    mana_misc = [r for r in tables["SpellMisc"] if r["SpellID"] in mana_spell_ids]
    mana_forms = [r for r in tables["SpellShapeshift"] if r["SpellID"] in mana_spell_ids]
    assert len(mana_links) == len(MANA_ITEMS) and len(mana_misc) == len(MANA_ITEMS)
    assert not any(int(r["Attributes_0"]) & 0x10000 for r in mana_misc)
    assert not mana_forms
    enchants = [r for r in tables["SpellItemEnchantment"] if int(r["ID"]) in ENCHANTS]
    assert len(enchants) == len(ENCHANTS)
    brand_links = [r for r in tables["SpellXDescriptionVariables"] if int(r["SpellID"]) in DEMONIC_BRAND]
    brand_variables = {r["SpellDescriptionVariablesID"] for r in brand_links}
    records = []
    for enchant in enchants:
        spell_ids = {int(enchant[f"EffectArg_{i}"]) for i in range(3) if enchant[f"Effect_{i}"] == "3"}
        effects = [e for e in tables["SpellEffect"] if int(e["SpellID"]) in spell_ids]
        stats = {"HealingPower": 0, "SpellDamage": 0}
        for effect in effects:
            if effect["Effect"] == "6" and effect["EffectAura"] in ("13", "135"):
                assert effect["EffectMiscValue_0"] == "126", effect
                stat = "SpellDamage" if effect["EffectAura"] == "13" else "HealingPower"
                stats[stat] += float(effect["EffectBasePointsF"])
        records.append({"enchantId": int(enchant["ID"]), "record": enchant,
                        "effects": effects, "healingAndDamageStats": stats})
    output = {
        "schemaVersion": 1,
        "build": args.build,
        "sources": [{
            "table": name,
            "url": f"https://wago.tools/db2/{name}/csv?build={args.build}",
            "sha256": hashlib.sha256((Path(args.cache) / f"{name}_{args.build}.csv").read_bytes()).hexdigest(),
        } for name in names],
        "enchants": records,
        "hotStreak": {
            "talentId": 400624, "auraId": 400625,
            "auraOptions": [r for r in tables["SpellAuraOptions"] if int(r["SpellID"]) in (400624, 400625)],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in (400624, 400625)],
        },
        "mageClearcasting": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] in ("11213", "12536")],
            "auraOptions": [r for r in tables["SpellAuraOptions"] if r["SpellID"] in ("11213", "12536")],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] in ("11213", "12536")],
        },
        "projectiles": {
            "misc": [r for r in tables["SpellMisc"] if int(r["SpellID"]) in PROJECTILE_SPELLS],
        },
        "hunterPetBite": {
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] == "17261"],
        },
        "hunterCastingRegen": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] in ("1242688", "1242512")],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] in ("1242688", "1242512")],
        },
        "paladinCooldowns": {
            "cooldowns": [r for r in tables["SpellCooldowns"] if int(r["SpellID"]) in PALADIN_COOLDOWNS],
            "categories": [r for r in tables["SpellCategories"] if int(r["SpellID"]) in PALADIN_COOLDOWNS],
            "equipment": [r for r in tables["SpellEquippedItems"] if int(r["SpellID"]) in (20925, 20927, 20928)],
        },
        "warlockCooldowns": {
            "cooldowns": [r for r in tables["SpellCooldowns"] if int(r["SpellID"]) in WARLOCK_COOLDOWNS],
            "categories": [r for r in tables["SpellCategories"] if int(r["SpellID"]) in WARLOCK_COOLDOWNS],
        },
        "decimation": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] in ("440870", "440873")],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] in ("440870", "440873")],
        },
        "firestone": {
            "descriptions": [r for r in tables["Spell"] if int(r["ID"]) in FIRESTONE],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in FIRESTONE],
            "levels": [r for r in tables["SpellLevels"] if int(r["SpellID"]) in FIRESTONE],
        },
        "spellstone": {
            "descriptions": [r for r in tables["Spell"] if int(r["ID"]) in SPELLSTONE],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in SPELLSTONE],
            "levels": [r for r in tables["SpellLevels"] if int(r["SpellID"]) in SPELLSTONE],
            "limitation": "The effect school mask is 36 (Fire + Shadow), while the tooltip names only Shadow. The model follows the effect mask; the Fire portion needs gameplay confirmation.",
        },
        "demonicBrand": {
            "descriptions": [r for r in tables["Spell"] if int(r["ID"]) in DEMONIC_BRAND],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in DEMONIC_BRAND],
            "auraOptions": [r for r in tables["SpellAuraOptions"] if int(r["SpellID"]) in DEMONIC_BRAND],
            "misc": [r for r in tables["SpellMisc"] if int(r["SpellID"]) in DEMONIC_BRAND],
            "variableLinks": brand_links,
            "variables": [r for r in tables["SpellDescriptionVariables"] if r["ID"] in brand_variables],
            "model": "Uses matching-school owner spell power at the pet hit, then pet damage multipliers once; Imp uses Fire, other base pets currently use Shadow. Target-scoped charges accept direct pet spells and melee. Child spells have the ALWAYS_HIT attribute.",
            "limitations": [
                "Owner versus pet power attribution and snapshot timing need a controlled capture.",
                "Imp/Fire and Succubus/Shadow follow their spell schools. Voidwalker/Felhunter mapping remains provisional; a physical child spell also exists.",
                "The 3x threat multiplier and noncritical proc behavior remain inherited assumptions.",
            ],
        },
        "druidHistory": {
            "queriedSpellIds": DRUID_HISTORY,
            "descriptions": [r for r in tables["Spell"] if int(r["ID"]) in DRUID_HISTORY],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in DRUID_HISTORY],
            "families": [r for r in tables["SpellClassOptions"] if int(r["SpellID"]) in DRUID_HISTORY],
            "shapeshift": [r for r in tables["SpellShapeshift"] if int(r["SpellID"]) in DRUID_HISTORY],
            "traits": [r for r in tables["TraitDefinition"] if int(r["SpellID"]) in DRUID_HISTORY],
            "skills": [r for r in tables["SkillLineAbility"] if int(r["Spell"]) in DRUID_HISTORY],
            "interpretation": "No sourced Cat Mangle; the available ranks require Bear form. Berserk's extra-target mask selects Mangle, not Lacerate. Nature's Grace separately grants 10% cast haste and reduces eligible GCDs by 10%, including instant Balance spells and Faerie Fire.",
        },
        "felArmorAvailability": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] == "403619"],
            "skills": [r for r in tables["SkillLineAbility"] if r["Spell"] == "403619"],
            "traits": [r for r in tables["TraitDefinition"] if r["SpellID"] == "403619"],
            "interpretation": "Retained spell and acquisition-method-3 skill entry do not establish current learnability. The misleading UI entry selecting Demon Armor under a Fel Armor tooltip was removed; no Fel Armor bonus is granted.",
        },
        "buffAvailability": {
            "queriedSpellIds": BUFF_ACCESS,
            "skills": [r for r in tables["SkillLineAbility"] if int(r["Spell"]) in BUFF_ACCESS],
            "currentTraits": [r for r in tables["TraitDefinition"] if int(r["SpellID"]) in BUFF_ACCESS],
            "legacyTalents": [r for r in tables["Talent"]
                              if int(r["SpellID"]) in BUFF_ACCESS or any(int(r[f"SpellRank_{i}"]) in BUFF_ACCESS for i in range(9))],
            "interpretation": "Sanctity Aura retains a legacy talent/spell record but no current trait or class-skill entry; excluded from Forever. Divine Spirit retains Priest class-skill entries, but trainer access after talent removal still needs gameplay confirmation.",
        },
        "roguePuncturingWounds": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] in ("1224716", "1241584")],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] in ("1224716", "1241584", "1241586", "1241590")],
            "families": [r for r in tables["SpellClassOptions"] if r["SpellID"] in ("1241586", "1241590")],
        },
        "warriorShieldTalents": {
            "descriptions": [r for r in tables["Spell"] if int(r["ID"]) in (12792, 16538, 23602, 1310316)],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in (12792, 16538, 23602, 1310316)],
        },
        "druidOmen": {
            "descriptions": [r for r in tables["Spell"] if int(r["ID"]) in (16864, 16870)],
            "auraOptions": [r for r in tables["SpellAuraOptions"] if int(r["SpellID"]) in (16864, 16870)],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in (16864, 16870)],
            "levels": [r for r in tables["SpellLevels"] if int(r["SpellID"]) in (16864, 16870)],
            "misc": [r for r in tables["SpellMisc"] if int(r["SpellID"]) in (16864, 16870)],
            "duration": [r for r in tables["SpellDuration"] if r["ID"] == "8"],
            "model": "Provisional literal client model: 100% chance after ten-second ICD; server-side proc overrides unverified.",
        },
        "druidMoonglow": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] == "16845"],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] == "16845"],
            "representativeSpellFamilies": [r for r in tables["SpellClassOptions"]
                                           if int(r["SpellID"]) in (5176, 2912, 8921, 5570, 16914, 9907, 768, 24858, 5185)],
        },
        "druidEnrage": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] == "5229"],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] == "5229"],
        },
        "druidThickHide": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] == "16929"],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] == "16929"],
            "limitation": "Form scope is explicit. The inherited .67 Defense-to-armor coefficient per point remains provisional; current raw scalar and captured rank text need reconciliation.",
        },
        "innervate": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] == "29166"],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] == "29166"],
            "power": [r for r in tables["SpellPower"] if r["SpellID"] == "29166"],
            "cooldowns": [r for r in tables["SpellCooldowns"] if r["SpellID"] == "29166"],
            "accounting": "Bonus regeneration is credited to Innervate; ordinary regeneration receives first claim on missing mana. Casting cost remains a separate negative event.",
        },
        "druidFaerieFire": {
            "shapeshift": [r for r in tables["SpellShapeshift"] if r["SpellID"] in ("770", "9907")],
            "power": [r for r in tables["SpellPower"] if r["SpellID"] in ("770", "9907")],
            "cooldowns": [r for r in tables["SpellCooldowns"] if r["SpellID"] in ("770", "9907")],
            "misc": [r for r in tables["SpellMisc"] if r["SpellID"] in ("770", "9907")],
        },
        "manaConsumables": {
            "itemLinks": mana_links,
            "itemEffects": mana_effects,
            "misc": mana_misc,
            "shapeshift": mana_forms,
            "model": "No NOT_SHAPESHIFT attribute or form exclusion; these mana consumables preserve the current form.",
        },
        "shaman": {
            "descriptions": [r for r in tables["Spell"] if int(r["ID"]) in SHAMAN_SPELLS],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in SHAMAN_SPELLS],
            "auraOptions": [r for r in tables["SpellAuraOptions"] if int(r["SpellID"]) in SHAMAN_SPELLS],
            "cooldowns": [r for r in tables["SpellCooldowns"] if int(r["SpellID"]) in SHAMAN_SPELLS],
            "limitations": [
                "Maelstrom's effective proc rate and Windfury's ICD/extra-attack table remain unverified.",
                "Fire Nova's active-Fire-totem requirement does not establish target-position or threat magnitudes.",
            ],
        },
        "priest": {
            "descriptions": [r for r in tables["Spell"] if int(r["ID"]) in PRIEST_SPELLS],
            "effects": [r for r in tables["SpellEffect"] if int(r["SpellID"]) in PRIEST_SPELLS],
            "power": [r for r in tables["SpellPower"] if int(r["SpellID"]) in PRIEST_SPELLS],
            "cooldowns": [r for r in tables["SpellCooldowns"] if int(r["SpellID"]) in PRIEST_SPELLS],
            "shapeshift": [r for r in tables["SpellShapeshift"] if int(r["SpellID"]) in PRIEST_SPELLS],
        },
        "hunter": {
            "descriptions": [r for r in tables["Spell"] if r["ID"] in ("1310532", "1310627", "1310726")],
            "effects": [r for r in tables["SpellEffect"] if r["SpellID"] in ("1310532", "1310627", "1310726")],
            "auraOptions": [r for r in tables["SpellAuraOptions"] if r["SpellID"] in ("1310532", "1310627", "1310726")],
        },
        "limitations": [
            "Enchant effects do not establish recipe availability.",
            "Hot Streak has three stacks but one proc charge; server-side interrupted-cast timing remains untested.",
            "Omen's proc frequency is a provisional interpretation of client metadata, not an observed server rate.",
        ],
    }
    args.output.write_text(json.dumps(output, indent=2) + "\n")


if __name__ == "__main__":
    main()
