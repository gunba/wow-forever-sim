"""Build labels and existing game icons, matching the simulator's talent trees."""

import json
from pathlib import Path

BUILDS = [
    ("balance", "Druid", "Balance", "spell_nature_starfall"),
    ("feral", "Druid", "Feral", "ability_druid_catform"),
    ("beast_mastery", "Hunter", "Beast Mastery", "ability_hunter_beasttaming"),
    ("marksmanship", "Hunter", "Marksmanship", "ability_marksmanship"),
    ("survival", "Hunter", "Survival", "ability_hunter_swiftstrike"),
    ("pet_melee", "Hunter", "Pet/Melee", "ability_hunter_beasttaming"),
    ("arcane", "Mage", "Arcane", "spell_holy_magicalsentry"),
    ("fire", "Mage", "Fire", "spell_fire_firebolt02"),
    ("frost", "Mage", "Frost", "spell_frost_frostbolt02"),
    ("arcane_frost", "Mage", "Arcane–Frost", "spell_frost_frostbolt02"),
    ("retribution", "Paladin", "Retribution", "spell_holy_auraoflight"),
    ("retribution_physical", "Paladin", "Physical Ret", "spell_holy_blessingofstrength"),
    ("smite", "Priest", "Smite", "spell_holy_holysmite"),
    ("shadow", "Priest", "Shadow", "spell_shadow_shadowwordpain"),
    ("combat", "Rogue", "Combat", "ability_backstab"),
    ("mutilate", "Rogue", "Mutilate", "ability_rogue_eviscerate"),
    ("subtlety", "Rogue", "Subtlety", "ability_stealth"),
    ("elemental", "Shaman", "Elemental", "spell_nature_lightning"),
    ("stormcaller", "Shaman", "Stormcaller", "spell_nature_lightning"),
    ("enhancement", "Shaman", "Enhancement", "ability_shaman_stormstrike"),
    ("affliction", "Warlock", "Affliction", "spell_shadow_deathcoil"),
    ("demonology", "Warlock", "Demonic Pact", "spell_shadow_metamorphosis"),
    ("ds_ruin", "Warlock", "DS/Ruin", "spell_shadow_psychicscream"),
    ("destruction", "Warlock", "Destruction", "spell_shadow_rainoffire"),
    ("arms", "Warrior", "Arms", "ability_warrior_savageblow"),
    ("fury", "Warrior", "Fury", "ability_warrior_innerrage"),
    ("fury_sunder", "Warrior", "Fury (Sunder)", "ability_warrior_sunder"),
    ("fury_2h", "Warrior", "2H Bloodthirst", "ability_warrior_innerrage"),
    ("tank_warrior", "Warrior", "Protection · Tank", "ability_warrior_defensivestance"),
    ("protection_paladin", "Paladin", "Protection · Tank", "spell_holy_devotionaura"),
    ("feral_tank_druid", "Druid", "Bear · Tank", "ability_racial_bearform"),
]

BUILD_CAVEATS = {
    "feral": ["Crusader's proc eligibility and PPM in animal forms are unverified (DRU-011); the current Night Elf profile uses that enchant."],
    "pet_melee": ["Hawk damage uses an approximate guardian/pet model. Tracking talents affect its comparison with Survival on different creature types."],
    "marksmanship": ["Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target."],
    "shadow": ["Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints."],
    "arcane_frost": ["Uses an assumed Ice Lance coefficient and unresolved Fingers of Frost, Missile Barrage and Clearcasting timing."],
    "fury": ["Level-60 rage, queued off-hand hit and Flurry charge timing still need in-game confirmation."],
    "fury_sunder": ["This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies."],
    "fury_2h": ["Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model."],
    "retribution": ["Separate seal Echoes can coexist and fire on a landed white swing; verify their simultaneous behavior in game (T53). Lower-rank seals and Consecration also need confirmation."],
    "retribution_physical": ["Strength/AP equipment is unchanged. Champion of the Light converts existing Intellect into spell power without equipping caster gear; the separate seal Echoes still require the T53 in-game check."],
    "smite": ["The rank-2 Smite fallback depends on the modeled low-rank spell-power coefficient."],
    "stormcaller": ["The rank-2 Lightning Bolt filler depends on the modeled low-rank spell-power coefficient."],
    "tank_warrior": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Different external support from the non-attacking DPS rows. Rage, shield proc eligibility and scripted Tier effects remain qualified in the uncertainty register."],
    "protection_paladin": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Seal of Fury is not implemented; shield and Spiritual Attunement questions remain in the uncertainty register. This is not a survival ranking."],
    "feral_tank_druid": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Bear rage and threat coefficients remain provisional; see the uncertainty register. This is not a survival ranking.", "Crusader's proc eligibility and PPM in animal forms are unverified (DRU-011); the selected enchant gain depends on that model."],
}

HYBRID_PARENTS = {"pet_melee": "survival", "arcane_frost": "frost", "fury_2h": "arms", "fury_sunder": "fury"}


def expected_roster():
    races = json.loads((Path(__file__).resolve().parents[2] /
                        "assets/db_inputs/forever_races.json").read_text())["races"]
    return {(key, race["name"]) for key, cls, _, _ in BUILDS
            for race in races if cls in race["classes"]}


RACES = [
    "Human", "Dwarf", "Night Elf", "Gnome", "High Order",
    "Orc", "Tauren", "Troll", "Undead", "Windshaper",
]

CLASS_COLORS = {
    "Druid": "#ff7d0a", "Hunter": "#abd473", "Mage": "#69ccf0",
    "Paladin": "#f58cba", "Priest": "#c6cbd3", "Rogue": "#e6cc22",
    "Shaman": "#0070de", "Warlock": "#9482c9", "Warrior": "#c79c6e",
}
