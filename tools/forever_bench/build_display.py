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
    "feral": ["Blizzard supports proc-enchant form eligibility, and Crusader's sourced health return is modeled. Its exact PPM/conditional interactions remain qualified (DRU-011); the Night Elf profile uses that enchant."],
    "pet_melee": ["Two independent Hawks use provisional high-rank damage and crit calibration (HUN-004); no owner-crit or unverified Lone Wolf coexistence is assumed. Tracking matches the Dragonkin reference target."],
    "beast_mastery": ["Hawk lifetimes and ranged-speed behavior are source-backed; high-rank assault damage, crit and Ferocity transfer remain qualified (HUN-004). The two-Hawk policy is conditional on the central calibration."],
    "marksmanship": ["Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target."],
    "shadow": ["Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints."],
    "arcane_frost": ["Ice Lance uses a provisional 0.10 SP share calibrated from lower ranks. Fingers of Frost ordering and Arcane Missiles/Clearcasting proc interpretation remain unresolved; Missile Barrage now requires a landed eligible hit."],
    "fury": ["Level-60/off-hand rage and some extra-swing interactions remain provisional. The queued-Heroic-Strike off-hand hit bug is removed. Deep Wounds uses the reported intended no-AP rollover model; live delivery and off-hand payload attribution remain qualified."],
    "fury_sunder": ["This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies."],
    "fury_2h": ["Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model."],
    "retribution": ["Separate seal Echoes can coexist and fire on a landed white swing; their simultaneous server behavior remains qualified (PAL-002). Active presets use maximum registered ranks, not deliberate downranking."],
    "retribution_physical": ["Uses a physical equipment budget without caster gear. Champion of the Light converts existing Intellect into damage-only spell bonus, not healing power. Seal Echo ordering remains qualified (PAL-002)."],
    "smite": ["Uses maximum registered ranks. Penance's static rank anomaly and actual skill-book acquisition remain source gaps; no efficient lower-rank fallback is assumed."],
    "stormcaller": ["Uses maximum registered Lightning Bolt. Overload uses its own client rank rows; Maelstrom/proc and server-script questions remain qualified."],
    "tank_warrior": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Different external support from the non-attacking DPS rows. Rage, shield proc eligibility and scripted Tier effects remain qualified in the uncertainty register.", "Heroic Strike/Cleave now replace swings. Earlier Protection gear/talent search gains used an invalid direct-cast APL and are superseded; the existing gear and talents have been replayed, not re-optimized."],
    "protection_paladin": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Seal of Fury and finite shielding are implemented; remaining shield/Spiritual Attunement details are tracked separately. Gear selections must pass both workload guard policies, which permit limited mitigation tradeoffs. This is not a survival ranking."],
    "feral_tank_druid": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Bear retains unverified damage-based rage, not Warrior's normalized formula (DRU-012); threat coefficients also remain provisional. This is not a survival ranking.", "Proc-enchant form eligibility is supported; exact PPM/conditional interactions remain qualified (DRU-011). Historical enchant-selection gains do not validate the current model."],
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
