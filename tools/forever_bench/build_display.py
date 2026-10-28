"""Build labels and existing game icons, matching the simulator's talent trees."""

import json
from pathlib import Path

BUILDS = [
    ("balance", "Druid", "Balance", "spell_nature_starfall"),
    ("feral", "Druid", "Feral", "ability_druid_catform"),
    ("beast_mastery", "Hunter", "BM Weave", "ability_hunter_beasttaming"),
    ("marksmanship", "Hunter", "MM Weave", "ability_marksmanship"),
    ("marksmanship_ranged", "Hunter", "Marksmanship", "ability_trueshot"),
    ("survival", "Hunter", "Survival", "ability_hunter_swiftstrike"),
    ("pet_melee", "Hunter", "Pet/Melee", "ability_hunter_beasttaming"),
    ("survival_weave", "Hunter", "SV Weave", "ability_hunter_swiftstrike"),
    ("arcane", "Mage", "Arcane", "spell_holy_magicalsentry"),
    ("fire", "Mage", "Fire", "spell_fire_firebolt02"),
    ("frost", "Mage", "Frost", "spell_frost_frostbolt02"),
    ("arcane_frost", "Mage", "Arcane–Frost", "spell_frost_frostbolt02"),
    ("frostfire", "Mage", "Frostfire", "ability_mage_frostfirebolt"),
    ("retribution", "Paladin", "Retribution", "spell_holy_auraoflight"),
    ("retribution_physical", "Paladin", "Physical Ret", "spell_holy_blessingofstrength"),
    ("shockadin", "Paladin", "Shockadin", "spell_holy_searinglight"),
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
    "beast_mastery": ["Melee weave from 12 yards with the Cat and two Hawks out: steps in for Raptor Strike, Mongoose Bite and Immolation Trap while the next Auto Shot is more than 1.5 s away, then back out (wowsims/forever BM Weave). Movement takes the engine's run speed; no latency is modelled.", "Hawk lifetimes and ranged-speed behavior are source-backed; high-rank assault damage, crit and Ferocity transfer remain qualified (HUN-004). The two-Hawk policy is conditional on the central calibration."],
    "marksmanship": ["Melee weave from 12 yards with Marksmanship and Survival talents (wowsims/forever MM Weave): steps in for Raptor Strike, Mongoose Bite, Strider Kick and Immolation Trap while the next Auto Shot is more than 1.5 s away, then back out, and keeps its own Hunter's Mark for Expose Prey. Movement takes the engine's run speed; no latency is modelled.", "Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target."],
    "marksmanship_ranged": ["Ranged only: Auto Shot from 12 yards and never in melee. Aimed Shot, Arcane Shot and Sniper Shot on cooldown, Serpent Sting kept up, Multi-Shot only when Aimed Shot cannot be cast (they share a cooldown), Rapid Fire and Thorium Grenades; an Immolation Trap is laid before the pull. Cat pet. Trueshot Aura is taken, but the benchmark raid already has it, so it adds no DPS here."],
    "shadow": ["Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints."],
    "arcane_frost": ["Ice Lance uses a provisional 0.10 SP share calibrated from lower ranks. Fingers of Frost ordering and Arcane Missiles/Clearcasting proc interpretation remain unresolved; Missile Barrage now requires a landed eligible hit."],
    "fury": ["Level-60/off-hand rage and some extra-swing interactions remain provisional. The queued-Heroic-Strike off-hand hit bug is removed. Deep Wounds uses the reported intended no-AP rollover model; live delivery and off-hand payload attribution remain qualified.", "Hamstring is woven for Windfury Totem procs when Bloodthirst and Whirlwind are at least 1.5 s away, outside execute (wowsims/forever #243)."],
    "fury_sunder": ["This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies.", "Hamstring is woven for Windfury Totem procs when Bloodthirst and Whirlwind are at least 1.5 s away, outside execute (wowsims/forever #243)."],
    "fury_2h": ["Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model.", "Hamstring is woven for Windfury Totem procs when Bloodthirst and Whirlwind are at least 1.5 s away, outside execute (wowsims/forever #243)."],
    "retribution": ["Separate seal Echoes can coexist and fire on a landed white swing; their simultaneous server behavior remains qualified (PAL-002). Active presets use maximum registered ranks, not deliberate downranking."],
    "survival_weave": ["Melee weave from 12 yards: steps in for Raptor Strike, Mongoose Bite, Strider Kick and Immolation Trap while the next Auto Shot is more than 1.5 s away, then back out. No pet. Movement takes the engine's run speed; no latency is modelled."],
    "frostfire": ["Frostfire Bolt spam behind five Fire Vulnerability stacks, Pyroblast on three Heating Up stacks. Uses the hotfixed Frost tree (20 October 2026)."],
    "shockadin": ["Holy Shock Retribution with Seal of Righteousness, Judgement, Consecration, Hammer of Wrath and Holy Strike. Champion of the Light at the hotfixed 20/40/60% of Intellect. No Manual Crowd Pummeler swap."],
    "retribution_physical": ["Uses a physical equipment budget without caster gear. Champion of the Light converts existing Intellect into damage-only spell bonus, not healing power. Seal Echo ordering remains qualified (PAL-002)."],
    "smite": ["Uses maximum registered ranks. Penance's static rank anomaly and actual skill-book acquisition remain source gaps; no efficient lower-rank fallback is assumed."],
    "stormcaller": ["Uses maximum registered Lightning Bolt. Overload uses its own client rank rows; Maelstrom/proc and server-script questions remain qualified."],
    "tank_warrior": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Incoming Rage uses a configurable, unverified expected-creature-health/reference-Armor scenario shared with Bear. Its coefficient, level lookup and reference curve are assumptions, not measured server behavior; resolved parameters accompany each replay.", "Heroic Strike/Cleave replace swings and the APL reserves Rage for self-maintained Sunder. Earlier direct-cast Protection searches are superseded. Last Stand removes its temporary health on expiration; the nonlethal floor follows the qualified Classic convention."],
    "protection_paladin": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Seal of Fury is now selectable with sourced damage, shield requirement and talent-only depletion mana. Its shared replacement shield pool and independent lifetime are provisional; the primary seal and lifetime are recorded, not silently substituted. Threat-based taunt target selection is not modeled.", "Gear selections must pass both workload guard policies, which permit limited mitigation tradeoffs. This is not a survival ranking. Reflections use explicit provider power, never the receiving tank's spell power."],
    "feral_tank_druid": ["Tank encounter: frontal attacks, incoming boss damage and modeled healing. Bear's outgoing base Rage remains the qualified damage-based model (DRU-012), with sourced double critical Rage. Incoming Rage now shares Warrior's configurable provisional expected-health scenario. The form threat multiplier is sourced at 1.5; this is not a survival ranking.", "Proc-enchant form eligibility is supported. Recovery has a sourced heal/ICD; Revelation requires an explicit, unverified proc model. Historical enchant-selection gains do not validate those assumptions."],
}

HYBRID_PARENTS = {"pet_melee": "survival", "survival_weave": "survival", "arcane_frost": "frost", "frostfire": "fire",
                  "fury_2h": "arms", "fury_sunder": "fury", "shockadin": "retribution"}


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
