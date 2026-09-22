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
    ("fury_2h", "Warrior", "2H Bloodthirst", "ability_warrior_innerrage"),
]

BUILD_CAVEATS = {
    "pet_melee": ["Hawk damage uses an approximate guardian/pet model. Its advantage over the revised Survival rotation is small."],
    "marksmanship": ["The Hawk/Sniper build depends on the approximate Hawk guardian and pet-inheritance model."],
    "arcane_frost": ["Uses an assumed Ice Lance coefficient and unresolved Fingers of Frost, Missile Barrage and Clearcasting timing."],
    "fury_2h": ["Level-60 rage generation remains unverified; this row uses the current Warrior rage model."],
    "retribution": ["Includes lower-rank seals and Consecration; their server scaling and proc interactions still need confirmation."],
    "smite": ["The rank-2 Smite fallback depends on the modeled low-rank spell-power coefficient."],
    "stormcaller": ["The rank-2 Lightning Bolt filler depends on the modeled low-rank spell-power coefficient."],
}

HYBRID_PARENTS = {"pet_melee": "survival", "arcane_frost": "frost", "fury_2h": "arms"}


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
