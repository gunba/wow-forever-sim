# Forever rule changes modelled in this fork

One line per rule the Forever ruleset changes, with where it lives here and what it was read
from. Written so an implementation on a different engine (the official wowsims fork) can work
from the rule rather than from this fork's diff. Numbers marked *demo* were read off BlizzCon
2026 tooltips, mostly rank 1; the per-rank assumptions behind them are listed in
`forever_beta_checklist.md`.

This file includes the initial demo-era interpretation. The class pages in
`beta-pass/` and [shared integration audit](beta-pass/core.md) record later
client checks and corrections; they supersede conflicting entries here.
The [auto-attack audit](auto_attack_audit.md) and
[Energy audit](energy_audit.md) describe the current casting-continuity rules and
the explicit general-haste regeneration assumption.

The ruleset is a `SimOptions.ruleset` enum (`RulesetClassic`, `RulesetForever`); every rule
below is gated on `IsForever()` unless it says the class code is Forever-only.

## Core

| Rule | Source | Here |
|---|---|---|
| Periodic damage can crit: dots and bleeds roll for critical strikes using the snapshot crit chance. Spells that must not (Ignite) carry `SpellFlagNoPeriodicCrit`. | Tooltip wording ("non-periodic" qualifiers on Nature's Grace, Primal Fury; Pandemic exists) | `sim/core/ruleset.go`, `sim/core/dot.go` |
| Hit and crit from gear apply to every kind of attack: an item's melee/spell hit and crit are summed and paid into both pools. Attribute conversions unchanged. | Panel | `sim/core/ruleset.go` `unifyEquipHitAndCrit` |
| Bonus healing on gear carries a damage component: `SpellDamage += HealingPower / 3`. | Panel | `sim/core/ruleset.go` `addHealingSpellDamage` |
| Improved Shadow Bolt and Stormstrike are personal: they raise only their caster's damage and are no longer raid debuffs. | Panel, confirmed by search | `sim/core/debuffs.go`, `sim/shaman/stormstrike.go` |
| Improved Shadow Bolt lasts a flat 12 s (not consumed per charge). | Published talent text | `sim/warlock/talents.go` |
| Stormstrike's Nature vulnerability lasts its full duration instead of being consumed by Nature hits. | Tooltip | `sim/core/debuffs.go` |
| Skinning: +5% damage to Beasts and Dragonkin. Mining: +5% health. | Panel (only professions with figures) | `sim/core/professions.go` |
| World buffs (Rallying Cry, Songflower, Darkmoon Faire, Warchief's Blessing, Dire Maul tribute, Spirit of Zandalar) do not work inside raids; the engine ignores them and the picker hides them. | Demo report, 13 September | `sim/core/buffs.go`, `settings_tab.ts` |
| Paladin blessings and auras are available to Horde raids because Undead Paladins are playable. | Forever race/class matrix | `sim/core/buffs.go` |

## Racials

| Rule | Source | Here |
|---|---|---|
| Every +10 resistance racial is removed. | Racials guide | `sim/core/racials.go` |
| Weapon skill racials become +1% crit (both pools) while the matching weapon is held; Mace Specialization moves to Dwarves. | Racials guide, *demo* | `sim/core/specializations.go` |
| Dwarf gains Beast Slaying-style +5% vs Beasts ("Big Game Hunter"). Troll keeps Beast Slaying; both ranged specializations removed. | Racials guide | `sim/core/racials.go` |
| Orc: Command removed. Blood Fury dynamically increases melee AP, ranged AP and spell power by 10%, for 15 seconds; free and off the GCD, two-minute cooldown. | Beta client 20572 | `sim/core/racials.go` |
| Gnome Expansive Mind raises Mana/Rage/Energy pools by 5%, not Intellect. Eureka has three charges for 15 s, a 2 min cooldown and class-specific cost reductions. | Client spell effects and class masks; [race audit](forever_races.md) | `sim/core/racials.go`, `sim/core/forever_gnome.go` |
| Night Elf Elune's Light (1259799): +10% crit for 15 s, 3 min cooldown. | Client spell effects, duration and cooldown | `sim/core/racials.go` |
| Skyborne: 1% haste and 5% damage versus Elementals. Windshaper's Skysight grants movement speed, not attack or spell power. | Beta client 1259710, 1259707, 1259686 | `sim/core/racials.go` |
| Troll Berserking: exactly 10% attack and casting speed for ten seconds, free and off the GCD, three-minute cooldown. | Beta client 20554 | `sim/core/racials.go` |
| Undead Touch of the Grave: 5% proc chance for Warrior/Paladin/Rogue, 10% for Priest/Mage/Warlock; one-second ICD; Shadow drain equal to 5% of max health before mitigation, cannot crit. | Beta client 1260189, 1260201, 1260198 and SpellAuraOptions | `sim/core/racials.go` |

## Warrior

| Rule | Source | Here |
|---|---|---|
| Slam no longer resets the swing timer. | Panel | `sim/warrior/slam.go` |
| Thunder Clap usable in Defensive Stance. | Panel | `sim/warrior/thunder_clap.go` |
| Improved Shield Wall shortens the cooldown instead of extending the duration. | Tooltip | `sim/warrior/shield_wall.go` |
| Tactical Mastery retains 10 Rage baseline; Improved Tactical Mastery adds 3 per point. | Demo spellbook, beta ability description | `sim/warrior/stances.go` |
| Enrage: any damage taken has a chance to grant a flat +2% Physical damage (was crit-only, scaled per point). | *demo* | `sim/warrior/talents.go` |
| Improved Cleave discounts Rage instead of adding damage. | *demo* | `sim/warrior/heroic_strike_cleave.go` |
| Improved Battle Shout and Improved Demoralizing Shout are gone from the tree; assumed baseline. Booming Voice only widens the radius. | Tree | `sim/warrior/shouts.go`, `demoralizing_shout.go` |
| Victory Rush baseline: not modelled (needs a killing blow). | Panel | - |

## Druid

| Rule | Source | Here |
|---|---|---|
| Energy regenerates smoothly at a base 10/sec; 100-ms integration approximates beta-observed behavior. Haste scaling remains unverified. | Client PowerType plus player beta reports; [energy audit](energy_audit.md) | `sim/core/energy.go` |
| Furor: shifting into Cat carries over a share of the energy left, plus a little per second out of form (was a chance at a flat 40). Per-rank scaling assumed linear. | *demo* | `sim/druid/forms.go` |
| Tiger's Fury: no Energy cost and a 30 sec cooldown (Wrath's shape), so that King of the Jungle's 60 Energy is a cooldown and not an engine. Assumed from the talent's wording. | Tree | `sim/druid/tigers_fury.go` |
| Nature's Grace: a short haste buff (also shortens the GCD) instead of a cast time cut on the next cast. | *demo* | `sim/druid/talents.go` |
| Thick Hide: flat base Armor from level and defense skill instead of an armor multiplier. | *demo* | `sim/druid/talents.go` |
| Primal Fury absorbs the old Blood Frenzy combo point proc. | Tree | `sim/druid/talents.go` |
| Savage Fury includes Shred. | Tooltip | `sim/druid/shred.go` |
| Improved Mark of the Wild and the feral Faerie Fire talent are gone; assumed baseline. | Tree | `sim/druid/druid.go`, `faerie_fire.go` |
| Feral Aggression is gone; Demoralizing Roar's attack power reduction is assumed baseline at full strength. | Tree | `sim/druid/demoralizing_roar.go` |
| Ferocity also cuts the Rage cost of Mangle. | Tooltip | `sim/druid/mangle.go` |
| Moonkin Form grants 360% more armor from items and the party crit aura, and nothing else: no spell damage and no Moonfire bonus. | Tooltip | `sim/druid/forms.go` |
| Feral Instinct is Swipe damage, not Bear Form threat; Bear Form's threat is the flat 1.3x. | Tooltip | `sim/druid/swipe.go`, `forms.go` |
| Natural Reaction also gives a 20% chance at 5 Rage on every dodge. | *demo* | `sim/druid/talents.go` |
| Mangle (Bear): 100% weapon damage plus 26; Berserk lifts its cooldown and widens it to 3 targets. | *demo* | `sim/druid/mangle.go`, `berserk.go` |
| Lacerate exists (Shredding Attacks cuts its Rage cost); modelled on the Season of Discovery Lacerate until a tooltip is seen. | Tree | `sim/druid/lacerate.go` |

## Paladin

| Rule | Source | Here |
|---|---|---|
| Precision: flat hit to all spells and attacks (Classic: melee only). | Tooltip | `sim/paladin/talents.go` |
| Redoubt triggers on any landed melee hit taken (was on being crit); Reckoning also has a smaller chance to fire on a block. | *demo* | `sim/paladin/talents.go` |
| Holy Shield block damage 110/161/220 (rank 1 ratio). | *demo* | `sim/paladin/holy_shield.go` |
| New abilities: Holy Strike (and its three dependent talents), Swift Judgement, Templar's Bulwark, Twist of Light; Consecrated Ground and Holy Conduit build on Consecration. | Tree, *demo* | `sim/paladin/holy_strike.go`, `swift_judgement.go`, `templars_bulwark.go`, `twist_of_light.go`, `consecration.go` |
| Holy Strike: instant, 20 mana, 12 sec cooldown, 40% weapon damage plus 36 to 46 Holy damage. | Published tooltip | `sim/paladin/holy_strike.go` |

## Priest

| Rule | Source | Here |
|---|---|---|
| Shadow Weaving buffs the priest, not the target. | Panel | `sim/priest/talents.go` |
| Divine Spirit and Improved Power Word: Fortitude are gone from the trees; assumed baseline raid buffs. | Tree | `sim/priest/priest.go` |
| Mind Flay base damage per tick from the rank 1 tooltip (119 vs 75), other ranks by ratio. | *demo* | `sim/priest/mind_flay.go` |
| Devouring Plague is castable by every race, not just the Undead: Devouring Contagion sits in the Shadow tree and does nothing otherwise. | Tree | `sim/priest/priest.go` |

## Mage

| Rule | Source | Here |
|---|---|---|
| Improved Scorch's fire vulnerability is personal to the mage who stacked it. Winter's Chill is a single personal stack for Frostbolt and Ice Lance. | Panel | `sim/mage/talents.go` |
| Ignite is excluded from periodic crits (it is already a share of a crit). | Design | `sim/mage/ignite.go` |
| Ignite pays out exactly 40% of the crit that lit it: the ticks skip the mage's and the target's damage multipliers, which the crit already carried, and a second crit rolls the damage still owed into the new dot instead of restarting it. | Design | `sim/mage/ignite.go` |
| Pyroblast dot damage from the rank 1 tooltip (76 vs 56), other ranks by ratio. Improved Fireball added to the tree. | *demo* | `sim/mage/pyroblast.go`, tree |
| Arcane rotation rebuilt around the Forever arcane talents (Arcane Impact, Arcane Shielding, ...). | Tree | `sim/mage/talents.go` |

## Rogue

| Rule | Source | Here |
|---|---|---|
| Malice and Precision also add spell crit/hit, since poisons roll against spell stats. | Tooltip | `sim/rogue/talents.go` |
| Venom rescales a Deadly Poison that is already ticking, not just the stacks applied while it is up. | Tooltip | `sim/rogue/poisons.go` |

## Shaman

| Rule | Source | Here |
|---|---|---|
| Dual wield is available. | Panel | `ui/core/proto_utils/utils.ts` `canDualWield` |
| Enhancing Totems is gone; Strength of Earth and Grace of Air always at the improved value (assumed baseline). | Tree | `sim/shaman/totems.go` |
| Call of Flame includes Lava Burst and Flame Shock. | Tooltip | `sim/shaman/lava_burst.go`, `flame_shock.go` |
| New talents/spells: Lava Burst, Lightning Overload, Maelstrom Weapon. | Tree, *demo* | `sim/shaman/` |

## Warlock

| Rule | Source | Here |
|---|---|---|
| Bane split into separate Shadow Bolt / Immolate reductions; Nightfall, Demonic Sacrifice and Pandemic per the Forever tree. | Tree, *demo* | `sim/warlock/talents.go` |
| Demonic Knowledge only pays its spell damage out while a demon is active. | Tooltip | `sim/warlock/talents.go` |
| Decimation's damage bonus belongs to Shadow Bolt and Searing Pain, the spells that trigger it; only the Soul Fire cast time carries the ten second window. | Tooltip | `sim/warlock/talents.go` |

## Hunter

| Rule | Source | Here |
|---|---|---|
| Barrage applies to Aimed Shot as well as Multi-Shot and Volley. | Tooltip | `sim/hunter/aimed_shot.go` |

## Encounters

| Rule | Source | Here |
|---|---|---|
| Tier 1 opens 9 December 2026: Barrow Deeps (10), Hyjal Summit (20), Onyxia's Lair (40). Only Onyxia has a known encounter. Molten Core kept as a target. | BlizzCon, Wowhead overview | `sim/encounters/register_all.go` |

## Spellbook coverage

The initial demo inventory was incomplete. Later client checks found omitted
damage abilities, including Frostfire Bolt, and resource abilities such as Dark
Sacrifice. Both are now implemented; see [shared integration](beta-pass/core.md)
for current limitations and remaining work.

- **Aspect of the Beast** is now implemented at level 60: 110 melee AP, mutually
  exclusive with Hawk, with Deadly Aspects' melee haste proc.
- **Shadow Word: Death** is not modelled, and neither is the Early Demise talent that buffs
  it, which is marked `notSimulated` in the tree so the picker says so.
- **Seal of Fury**, **Victory Rush**, **Fear Ward**, **Totemic Projection**, **Call of the
  Elements**, **Comprehend Scroll**, **Subjugate Demon**, **Incubus**, **Call Owl** and the
  movement, profession and dispel racials: none of them move a damage number, or the sim
  has no model for what they do (absorbs, self-healing, totem placement).

Several abilities the search does not find by name are implemented under another one:
Quickness is an `AddStat(stats.Dodge, 1)`, Axe Specialization is
`AddWeaponSpecializationCrit`, Bane of Agony lives in `sim/warlock/curses.go`. Check for the
effect before concluding an ability is missing.

## The Legacy system and the priest race abilities (checked 17 September)

Two more published datasets, neither of which the sim implements, both deliberately.

**Legacy** is three account-wide trees of 20 perks bought with a point per challenge
completed. None of the 20 touches a combat stat - they are rested experience, mount speed,
profession skill-ups, reputation and vendor discounts. The one that comes closest,
Adventure's "Talented", starts talent points at level 5 instead of 10, which changes
nothing for a level 60 character who has all 51 either way.

**Priest race abilities** are a per-race spell each, listed in the data as `class_racials`.
Only Starshards is modelled, and `sim/priest/starshards.go` gates it to Night Elves, which
is right. The other eleven are heals, crowd control, dispels, or retaliation that fires
only when the priest is hit - Touch of Weakness and Shadowguard both damage an attacker, so
they never trigger for a priest the sim never has attacked.

If either dataset grows a combat number, these are the two places to re-check.
