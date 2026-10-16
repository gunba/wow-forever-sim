# October 15 update

Baseline: published `0facf43a0` (results from the [October 11 log audit](october-11-log-audit.md)).

Evidence:
- **Client 1.60.1.70425**, compared with 1.60.1.70374 field by field ([wago.tools diff](https://wago.tools/builds-diff?to=1.60.1.70425&from=1.60.1.70374)):
  spell effects, durations, cooldowns, costs, aura options, class masks, item effects and talent trees.
- **The [October 15 development notes](https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-october-15/2360696/6)**,
  every item, not only the highlights.
- **Discord reports since October 11:** WoW Sims (368 messages), Fight Club Warrior channels (13,993 messages, 186 relevant) and
  Hunter Lodge (4,404 messages since the notes).
- **Upstream:** wowsims/classic has no new commits; the official Forever repository is private.

Register dispositions are in [Simulator questions and coverage](uncertainties.md); new open questions have test
procedures in [Beta test procedures](beta-test-procedures.md). Raw evidence is in
[`artifacts/correctness/october_15/`](../artifacts/correctness/october_15/). The previous results are
[archived](../artifacts/history/0facf43a0/manifest.json).

## Changes

### Races and raid buffs

| Change | Evidence | Register |
|---|---|---|
| Tauren Endurance: 2% critical damage instead of 1% hit. Dwarf Earth Fury replaces Mace Specialization with the same 2%. The whole critical multiplier is scaled (2.04x melee, 1.53x spells). | Notes; client 20550 (aura 163) | CORE-020 |
| Human Ruthless (+2% crit with spells and attacks) and Orc Savage (+1%) replace the weapon specializations. Orc Blood Fury: 3-minute cooldown. | Notes; client 20597, 20574, 20572 | CORE-015 |
| Skyborne Wind Blessed: 3% haste. | Notes | |
| Touch of the Grave: 20% (Warrior, Paladin, Rogue) or 30% (casters), 25-second cooldown, no Shadow scaling. | Notes; client aura options | CORE-021 |
| Eureka: an 8-second buff with no charges. | Notes; client 1259812-1259823 | CORE-009 |
| Leader of the Pack is 3% melee and ranged crit again, Moonkin Aura 3% spell crit; both are raid-wide. | Notes; client 24932 (aura 52), 24907 (aura 57) | DONE-005 |
| Battle Shout, Trueshot Aura and the Strength of Earth, Grace of Air, Windfury and Flametongue Totems are raid-wide: a provider in one party reaches every party. Grace of Air no longer excludes Windfury or Flametongue; those two still exclude each other. | Notes; client effect 35 to 65 | DONE-005 |

### Classes

| Class | Change | Register |
|---|---|---|
| Druid | Faerie Fire in Cat or Bear Form cannot proc Omen of Clarity. Furor keeps the saved Cat Energy through Bear Form, and Bear time earns none. Howling Idol takes 0.5 s off Shifting Power (was 1 s). | DRU-015 |
| Hunter | Multi-Shot is instant. Mortal Shots and Efficiency cover Sniper Shot. Pets inherit the owner's Expertise and crit (no crit from their own Agility). Resourcefulness procs only from melee and ranged ability crits. Fire traps gain Attack Power (provisional: 10% over the trap). | HUN-016, HUN-017 |
| Mage | Arcane Blast: 1.5 s, 40% less damage, each stack gives 15% of Intellect as Spell Damage for 15 s. New Arcane Wrath. Missile Barrage: 2 points, also procs from Arcane Missiles and Arcane Wrath. Fire Power: 1/2/3% per Fire effect on the target. Critical Mass: 5 points. Improved Fire Blast replaces Wake of Fire. Heating Up: Fire Blast always adds a stack and Presence of Mind keeps them. Winter's Chill: +40% Cone of Cold per stack, no crit. Ice Lance: x3 against frozen targets. Clearcasting rolls once per cast. | MAG-001, MAG-005, MAG-009, MAG-010 |
| Paladin | Judgement of Righteousness and Fury 0.429; Seal of Righteousness takes Spell Damage once; Consecration 0.042 a tick; Righteous Fury +90% Holy threat; Champion of the Light 10/20/30% of Intellect; Shield Specialization 4% mana; Sanctified Judgement's mana makes no threat; Reckoning stores up to 4 charges for the next swing. | PAL-001, PAL-016 |
| Priest | Devouring Plague 24 s, Shadow Word: Death 18 s, Vampiric Embrace 30 s. Mental Strength 2% Intellect a point. Silent Resolve covers all spells. Power in Light also counts Shadow Word: Pain. Twin Disciplines covers Shadow Word: Death, Mind Flay, Starshards and Penance. Improved Mind Flay escalates each tick. New Dark Inspiration. Holy Fire rank 5 DoT 45. Shadow Affinity and Improved Fade removed. | PRI-006 |
| Rogue | Cold Blood is spent when Mutilate lands. | ROG-001 |
| Shaman | Chain Lightning costs about 35% less. Fire Nova: 0.143, lower damage and cost. Lava Burst ranks at 30/40/50/60 (new rank 4). Elemental Mastery is the 31-point talent. Call of Flame adds 3/6/9 s to Flame Shock. Call of Thunder removed. Stormstrike costs 8% of base mana and its charge is only spent by a spell that lands. Water Shield triggers on any spell crit. | SHA-010 |
| Warlock | Malediction becomes Grim Reach (range). Shadow Mastery 2% a point. | |
| Warrior | Weaponmaster's extra attack no longer triggers from Shield Slam. Battle Shout is raid-wide. Slam pushback, Overpower against low-level enemies, Lingering Rage and Improved Berserker Rage change nothing simulated. | WAR-024 |

### Items

Bloodrazor (five 31-damage ticks over 15 s), Destiny (12 s), Fiery War Axe (231-293 plus 20 a tick) and Sulfuras'
Immolation (7) follow the client (DATA-017). The client's mass change of armor auras from type 22 to 674 keeps every value.

### Talent trees

Mage, Priest, Shaman and Warlock trees moved talents, changed ranks and added or removed talents. The trees, protos,
presets and every profile were re-encoded by talent name. Points from removed or shortened talents were re-spent:
Elemental and Stormcaller into Reverberation and Elemental Mastery, Fire into Critical Mass 5/5, Affliction, DS/Ruin
and Destruction into Improved Corruption and Improved Drains, Arcane and Arcane-Frost into Arcane Wrath and Missile
Barrage 2/2.

### Engine fixes found during the update

- Fire Power counted no Ignite (its DoT lives on the tick spell) and crashed without Impact.
- Reckoning's block cooldown was set after the trigger was built and never applied.
- `TestTalentTreesMatchTheirProtos` failed for Warrior and Druid, whose Forever trees are read by talent name; the test
  now checks those trees against their name lists.

## Selections

- **Talents.** Every changed build ran up to four rounds of one-point reallocation screens with independent
  confirmation. Elemental, Enhancement and Frost moved one or two points; the rest kept their migrated builds.
- **Rotations.** Arcane now stacks Arcane Blast to four before Arcane Wrath and fills with Arcane Missiles (+2.5% over
  the first draft on two seeds). Arcane-Frost stacks Arcane Blast before Frostbolts (+14%); Arcane Wrath is weaker
  than Frostbolt there. A Frost variant with Arcane Blast ran out of mana and was dropped.
- **Raid buffs.** The benchmark raid carries both Leader of the Pack and Moonkin Aura, and melee groups take Grace of Air
  alongside Windfury (CORE-022). Grace of Air alone is worth 3.6-6.1% to melee rows (Enhancement 11.7%). Tank support is unchanged.
- **Gear.** All 184 DPS profiles were re-searched on the final engine (seed 1793700101) and every change confirmed on
  an independent seed (1793900001). 44 changes were retained and none reverted: Retribution +3.3-3.7%, Arcane +0.7-3.0%,
  Arcane-Frost +0.5-2.1%, and smaller changes in 13 other builds.

## Benchmark

All 1,005 scenarios were replayed: baseline, Tier 1 off, +10% and +50% equipment, and the previous setups on the new
engine. Seed 1793400001, 5,000 iterations each. All 68 tank workload checks were replayed with unchanged tank setups.
Averaged over races:

| Build | Published | New engine, old setup | Released | Change |
|---|---:|---:|---:|---:|
| Arcane | 732 | 807 | 849 | +15.9% |
| Arcane–Frost | 743 | 740 | 850 | +14.5% |
| Protection Paladin | 603 | 528 | 528 | -12.3% |
| Frost | 690 | 602 | 608 | -11.9% |
| Enhancement | 727 | 797 | 801 | +10.2% |
| Feral | 703 | 757 | 757 | +7.6% |
| Fire | 729 | 780 | 781 | +7.0% |
| Elemental | 491 | 520 | 525 | +6.9% |
| Survival | 895 | 953 | 953 | +6.5% |
| Subtlety | 664 | 707 | 707 | +6.5% |
| Combat | 716 | 760 | 762 | +6.4% |
| Fury (Sunder) | 907 | 959 | 962 | +6.1% |
| Mutilate | 708 | 751 | 751 | +6.0% |
| Fury | 938 | 988 | 989 | +5.5% |
| 2H Bloodthirst | 799 | 841 | 841 | +5.2% |
| Pet/Melee | 909 | 956 | 956 | +5.1% |
| Arms | 825 | 866 | 866 | +4.9% |
| Retribution | 1054 | 971 | 1004 | -4.7% |
| Shadow | 809 | 827 | 827 | +2.3% |
| Affliction | 814 | 831 | 832 | +2.2% |
| Stormcaller | 456 | 463 | 464 | +1.7% |
| Physical Ret | 960 | 943 | 943 | -1.7% |
| Marksmanship | 955 | 967 | 967 | +1.3% |
| Smite | 450 | 444 | 444 | -1.2% |
| Balance | 710 | 717 | 717 | +1.1% |
| Bear | 540 | 544 | 544 | +0.6% |
| Beast Mastery | 980 | 975 | 975 | -0.5% |
| Destruction | 815 | 811 | 811 | -0.5% |
| Demonic Pact | 917 | 914 | 914 | -0.4% |
| Protection Warrior | 506 | 508 | 508 | +0.2% |
| DS/Ruin | 794 | 792 | 792 | -0.2% |

- Melee rows gain from Grace of Air; without it, Human and Orc Fury, Arms, Combat and Survival land within 0.7% of the published values (Enhancement 1.6% below).
- Frost loses Winter's Chill's crit. Arcane gains from the Arcane Blast and Arcane Wrath redesign.
- Retribution and Protection Paladin lose Seal of Righteousness' second Spell Damage share, Judgement and
  Consecration coefficient and half of Champion of the Light. Protection Paladin threat is unchanged (-1.1% single
  target, +0.2% three targets) because Righteous Fury rose to +90%.
- Tank survival stays within the earlier range: the highest modeled death probability is 0.6% (Bear, one attacker).

## Rotation planner

The Warrior rank tables, trait map, base stats and planner tables were regenerated from 70425; nothing but the build
stamp changed. The planner reads the racial critical-damage scale from the engine (a Tauren crits for 2.04x), and
Tauren characters no longer get 1% hit.

## Not simulated

Frost Blast (a level-24 utility spell), Salvation and Revoke, Cone of Cold's extra slow, Plainsrunning, Shatter Curse's
magic reduction, Troll Regeneration, Lingering Rage and the other movement, utility and out-of-combat changes.
