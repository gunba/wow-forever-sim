# Build reviews

These summaries describe the published loadouts. They are simulation results, not independent confirmation of server mechanics or proof of a global optimum.

The tables and [matrix](../artifacts/forever_dps_5min.png) use the same 147 common-seed replays. Baseline comparisons use the starting gear and enchants under the same engine and seed. Talents and APLs are unchanged by the equipment search.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and paid shared-hit normalization. [Scenario and exchange model](../tools/forever_bench/README.md) · [In-game checks](in_game_checks.md)

## Paladin — Retribution

**Talents:** 12/0/39 · `250003002--052253312012331321`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Undead | 896.29 | 1090.69 | +21.69% | 0.00 |
| Human | 881.78 | 1069.51 | +21.29% | 0.00 |
| Dwarf | 880.67 | 1067.76 | +21.24% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Arcanum of Rapidity |
| Neck | [Beads of Ogre Mojo](https://www.wowhead.com/forever/item=22149) | 63 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Spitfire Bracers](https://www.wowhead.com/forever/item=20481) | 62 | Enchant Bracer - Superior Strength |
| Hands | [Raider Handwraps](https://www.wowhead.com/forever/item=272097) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Rapidity |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Blackfury](https://www.wowhead.com/forever/item=19167) | 66 | Enchant Weapon - Crusader |
| Ranged/relic | [Libram of Law](https://www.wowhead.com/forever/item=272435) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -1.5s: [Spell 20920 (rank 5)](https://www.wowhead.com/forever/spell=20920).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Holy Strike (rank 8)](https://www.wowhead.com/forever/spell=10333).
3. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271) when `{"spellCanCast":{"spellId":{"spellId":20271}}}`.
4. Cast [Spell 20920 (rank 5)](https://www.wowhead.com/forever/spell=20920) when `{"currentSealRemainingTime":{}}` ≤ 1.
5. Cast [Spell 20293 (rank 8)](https://www.wowhead.com/forever/spell=20293) when (Mana fraction ≥ 15% AND NOT [Twist of Light](https://www.wowhead.com/forever/spell=77485) active AND [Spell 20920 (rank 5)](https://www.wowhead.com/forever/spell=20920) active).
6. Cast [Spell 20920 (rank 5)](https://www.wowhead.com/forever/spell=20920) when (Mana fraction ≥ 15% AND NOT [Twist of Light](https://www.wowhead.com/forever/spell=77485) active AND [Spell 20293 (rank 8)](https://www.wowhead.com/forever/spell=20293) active).
7. Cast [Hammer of Wrath (rank 3)](https://www.wowhead.com/forever/spell=24239).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 269.23 |
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 247.66 |
| Auto-attack (tag 1) | 191.64 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 100.17 |
| Auto-attack (tag 3) | 71.45 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20966) | 66.85 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 59.92 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 36.65 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +15250.1 |
| Mana | [Spell 20920](https://www.wowhead.com/forever/spell=20920) | -14367.6 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -13596.1 |
| Mana | OtherActionManaRegen (tag 1) | +4592.8 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +4235.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3541.7 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3122.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3037.5 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1565.7 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -524.6 |
| Mana | OtherActionManaRegen (tag 2) | +1.2 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Fire

**Talents:** 17/31/3 · `0501252000002-23450000130133051-003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 759.45 | 947.95 | +24.82% | 1.62 |
| Troll | 753.96 | 940.93 | +24.80% | 1.75 |
| Undead | 766.93 | 956.68 | +24.74% | 1.34 |
| Human | 757.08 | 943.74 | +24.66% | 1.42 |
| Gnome | 762.44 | 949.87 | +24.58% | 1.04 |
| High Order | 757.87 | 944.90 | +24.68% | 1.54 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherflame Cuffs](https://www.wowhead.com/forever/item=254065) | 50 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Fire Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 80%.
2. Use ready automatic cooldowns.
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 20% AND Time remaining > 25s).
4. Cast [Pyroblast (rank 8)](https://www.wowhead.com/forever/spell=18809) when [Hot Streak](https://www.wowhead.com/forever/spell=44445) stacks = 3.
5. Cast [Fire Blast (rank 7)](https://www.wowhead.com/forever/spell=10199) when Mana ≥ Time remaining × 10.
6. Cast [Scorch (rank 7)](https://www.wowhead.com/forever/spell=10207) when ([Improved Scorch](https://www.wowhead.com/forever/spell=12873) stacks < 5 OR `{"auraRemainingTime":{"auraId":{"spellId":12873}}}` ≤ 5s).
7. Cast [Fireball](https://www.wowhead.com/forever/spell=25306) when Mana ≥ Time remaining × 150.
8. Cast [Scorch](https://www.wowhead.com/forever/spell=10207).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 561.19 |
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 143.54 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 112.60 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 102.88 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 24.17 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 12.29 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -16012.9 |
| Mana | OtherActionManaRegen (tag 2) | +5504.4 |
| Mana | OtherActionManaRegen (tag 1) | +4580.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4440.1 |
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -3928.8 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -3888.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3457.9 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +2149.2 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.9 |
| Mana | [Fireball](https://www.wowhead.com/forever/spell=25306) | -868.4 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +839.1 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +406.1 |

## Warrior — Fury

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 799.66 | 951.37 | +18.97% | 0.00 |
| Tauren | 797.41 | 933.65 | +17.09% | 0.00 |
| Troll | 796.64 | 935.26 | +17.40% | 0.00 |
| Undead | 810.49 | 953.82 | +17.69% | 0.00 |
| Windshaper | 795.46 | 932.68 | +17.25% | 0.00 |
| Human | 807.31 | 946.13 | +17.20% | 0.00 |
| Dwarf | 797.14 | 934.15 | +17.19% | 0.00 |
| Night Elf | 795.06 | 932.10 | +17.24% | 0.00 |
| Gnome | 789.50 | 923.79 | +17.01% | 0.00 |
| High Order | 795.46 | 932.68 | +17.25% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Lionheart Helm](https://www.wowhead.com/forever/item=12640) | 61 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Defiler's Plate Spaulders](https://www.wowhead.com/forever/item=20212) | 65 | Chromatic Mantle of the Dawn |
| Back | [Deathguard's Cloak](https://www.wowhead.com/forever/item=20068) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Raider Gauntlets](https://www.wowhead.com/forever/item=272095) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Might of the Timbermaw](https://www.wowhead.com/forever/item=19044) | 58 | — |
| Legs | [Titanic Leggings](https://www.wowhead.com/forever/item=22385) | 60 | Lesser Arcanum of Voracity |
| Feet | [Scalegut Treaders](https://www.wowhead.com/forever/item=275618) | 58 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Legionnaire's Band](https://www.wowhead.com/forever/item=19510) | 63 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Molten Heart of the Mountain](https://www.wowhead.com/forever/item=249470) | 55 | — |
| Main hand | [Premier High Warlord's Blade](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Premier High Warlord's Quickblade](https://www.wowhead.com/forever/item=272684) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719) when Time remaining ≤ 300s.
2. Use ready automatic cooldowns.
3. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower](https://www.wowhead.com/forever/spell=11585) active AND Rage ≤ 25 AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND NOT [Recklessness](https://www.wowhead.com/forever/spell=1719) active).
4. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
5. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when NOT [Overpower](https://www.wowhead.com/forever/spell=11585) active.
6. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
7. Cast [Bloodthirst](https://www.wowhead.com/forever/spell=23894).
8. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
9. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 30.
10. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 30 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 291.89 |
| Auto-attack (tag 2) | 142.81 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 141.25 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 114.90 |
| Auto-attack (tag 3) | 53.87 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 46.81 |
| Auto-attack (tag 1) | 45.93 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 36.89 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +2816.8 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -1289.4 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -1258.4 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -1132.8 |
| Rage | Auto-attack (tag 1) | +1018.5 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -572.6 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -204.1 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +179.5 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +160.6 |
| Rage | OtherActionRefund | +157.9 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +96.6 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +89.4 |

## Hunter — Beast Mastery

**Talents:** 31/20/0 · `5320001505101251-00531510005`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 772.00 | 880.94 | +14.11% | 0.00 |
| Tauren | 770.11 | 870.91 | +13.09% | 0.00 |
| Troll | 771.94 | 876.72 | +13.57% | 0.00 |
| Windshaper | 773.15 | 877.42 | +13.49% | 0.00 |
| Human | 769.43 | 883.41 | +14.81% | 0.00 |
| Dwarf | 767.25 | 868.58 | +13.21% | 0.00 |
| Night Elf | 777.91 | 880.06 | +13.13% | 0.00 |
| High Order | 773.15 | 877.42 | +13.49% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Shifting Cloak](https://www.wowhead.com/forever/item=18511) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Boots - Minor Agility |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Weapon - Agility |
| Waist | [Molten Belt](https://www.wowhead.com/forever/item=19163) | 70 | — |
| Legs | [Sentinel's Chain Leggings](https://www.wowhead.com/forever/item=22748) | 65 | Lesser Arcanum of Voracity |
| Feet | [Scalegut Treaders](https://www.wowhead.com/forever/item=275618) | 58 | Enchant Boots - Greater Agility |
| Ring 1 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Ring 2 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Agility |
| Off hand | [Premier High Warlord's Blade](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Core Marksman Rifle](https://www.wowhead.com/forever/item=18282) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -10s: [Aspect of the Hawk](https://www.wowhead.com/forever/spell=25296).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Intimidation](https://www.wowhead.com/forever/spell=19577) when Mana fraction > 80%.
3. Cast [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) when NOT [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) DoT active.
4. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=25295) when NOT [Serpent Sting](https://www.wowhead.com/forever/spell=25295) DoT active.
5. Cast [Aimed Shot](https://www.wowhead.com/forever/spell=20902) when `{"autoTimeToNext":{"autoType":"Ranged"}}` ≥ 0s.
6. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287) when Mana fraction ≥ 15%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Shoot | 381.80 |
| Cat: Auto-attack (tag 1) | 148.84 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | 130.12 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 82.35 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 67.74 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 49.51 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 12.55 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 10.48 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | -9216.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7630.6 |
| Mana | OtherActionManaRegen (tag 1) | +6113.4 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5535.8 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -4607.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3597.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3591.7 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3159.9 |
| Mana | OtherActionManaRegen (tag 2) | +864.2 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Intimidation](https://www.wowhead.com/forever/spell=19577) | -405.6 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Hunter — Marksmanship

**Talents:** 18/33/0 · `5023000503-0053451001503051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 732.64 | 851.27 | +16.19% | 0.02 |
| Tauren | 730.35 | 840.42 | +15.07% | 0.02 |
| Troll | 732.50 | 845.00 | +15.36% | 0.03 |
| Windshaper | 733.67 | 846.57 | +15.39% | 0.02 |
| Human | 729.69 | 853.91 | +17.02% | 0.01 |
| Dwarf | 727.56 | 837.77 | +15.15% | 0.02 |
| Night Elf | 738.73 | 850.04 | +15.07% | 0.02 |
| High Order | 733.67 | 846.57 | +15.39% | 0.02 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Shifting Cloak](https://www.wowhead.com/forever/item=18511) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Boots - Minor Agility |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Weapon - Agility |
| Waist | [Molten Belt](https://www.wowhead.com/forever/item=19163) | 70 | — |
| Legs | [Sentinel's Chain Leggings](https://www.wowhead.com/forever/item=22748) | 65 | Lesser Arcanum of Voracity |
| Feet | [Scalegut Treaders](https://www.wowhead.com/forever/item=275618) | 58 | Enchant Boots - Greater Agility |
| Ring 1 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Ring 2 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Agility |
| Off hand | [Premier High Warlord's Blade](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Core Marksman Rifle](https://www.wowhead.com/forever/item=18282) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -10s: [Aspect of the Hawk](https://www.wowhead.com/forever/spell=25296).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) when Mana fraction ≥ 40%.
3. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=25295) when NOT [Serpent Sting](https://www.wowhead.com/forever/spell=25295) DoT active.
4. Cast [Aimed Shot](https://www.wowhead.com/forever/spell=20901) when `{"autoTimeToNext":{"autoType":"Ranged"}}` ≥ 0s.
5. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287) when `{"autoTimeToNext":{"autoType":"Ranged"}}` ≥ 0s.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Shoot | 401.46 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20901) | 129.47 |
| Cat: Auto-attack (tag 1) | 103.31 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 72.84 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 70.16 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 60.50 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 12.55 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 3.62 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8216.0 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -6782.9 |
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20901) | -5783.8 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -4929.7 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3731.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3603.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3580.2 |
| Mana | OtherActionManaRegen (tag 1) | +2923.9 |
| Mana | OtherActionManaRegen (tag 2) | +700.8 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 634.92 | 805.28 | +26.83% | 0.00 |
| Troll | 631.11 | 797.56 | +26.37% | 0.00 |
| Undead | 639.27 | 807.68 | +26.35% | 0.00 |
| Human | 641.22 | 797.00 | +24.29% | 0.00 |
| Gnome | 640.87 | 810.37 | +26.45% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Shadow Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Lorekeeper's Ring](https://www.wowhead.com/forever/item=19522) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -1s: [Soul Link](https://www.wowhead.com/forever/spell=19028).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 75%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 60%).
3. Use ready automatic cooldowns.
4. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
5. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 10%.
6. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
7. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
8. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`).
9. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) when [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311,"rank":7}}}`.
10. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) when [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309,"rank":8}}}`.
11. Cast [Shadow Bolt (rank 9)](https://www.wowhead.com/forever/spell=25307).

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 421.15 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 96.66 |
| Succubus: Auto-attack (tag 1) | 91.69 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 89.22 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 85.15 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 17.80 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 5.73 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.95 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -32266.4 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +12875.2 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +11216.0 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6363.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5367.5 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5224.7 |
| Mana | OtherActionManaRegen (tag 1) | +3890.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3621.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2557.8 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -488.7 |
| Mana | [Soul Link](https://www.wowhead.com/forever/spell=19028) | -274.6 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2525000013520105--0550015103`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 623.12 | 797.34 | +27.96% | 0.01 |
| Troll | 619.45 | 790.91 | +27.68% | 0.01 |
| Undead | 626.34 | 798.33 | +27.46% | 0.01 |
| Human | 629.58 | 788.70 | +25.27% | 0.01 |
| Gnome | 628.33 | 802.03 | +27.64% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Shadow Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Lorekeeper's Ring](https://www.wowhead.com/forever/item=19522) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -5s: [Amplify Curse](https://www.wowhead.com/forever/spell=18288).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 75%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 60%).
3. Use ready automatic cooldowns.
4. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871) when Time remaining ≤ 1.5.
5. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
6. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 10%.
7. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
8. Cast [Amplify Curse](https://www.wowhead.com/forever/spell=18288).
9. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
10. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when ((NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`) AND Time remaining ≥ 12s).
11. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) when ([Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311,"rank":7}}}` AND Time remaining ≥ 12s).
12. Cast [Siphon Life (rank 4)](https://www.wowhead.com/forever/spell=18881) when ([Siphon Life (rank 4)](https://www.wowhead.com/forever/spell=18881) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":18881,"rank":4}}}` AND Time remaining ≥ 12s).
13. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) when ([Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309,"rank":8}}}` AND Time remaining ≥ 12s).
14. Cast [Shadow Bolt (rank 9)](https://www.wowhead.com/forever/spell=25307).

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 368.02 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 104.97 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 97.67 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 83.47 |
| Succubus: Auto-attack (tag 1) | 81.02 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 39.43 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 18.85 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 3.83 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -30136.2 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24740.2 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6454.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5294.7 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5107.2 |
| Mana | OtherActionManaRegen (tag 1) | +3871.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3423.9 |
| Mana | [Siphon Life](https://www.wowhead.com/forever/spell=18881) | -3370.7 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.1 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -354.7 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -342.8 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 662.24 | 791.47 | +19.51% | 0.00 |
| Tauren | 659.51 | 780.05 | +18.28% | 0.00 |
| Troll | 661.00 | 783.49 | +18.53% | 0.00 |
| Windshaper | 661.52 | 786.47 | +18.89% | 0.00 |
| Human | 671.01 | 795.70 | +18.58% | 0.00 |
| Dwarf | 656.13 | 776.14 | +18.29% | 0.00 |
| Night Elf | 666.17 | 788.82 | +18.41% | 0.00 |
| High Order | 661.52 | 786.47 | +18.89% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Cloak of the Honor Guard](https://www.wowhead.com/forever/item=20073) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Weapon - Agility |
| Waist | [Highlander's Chain Girdle](https://www.wowhead.com/forever/item=20043) | 63 | — |
| Legs | [Sentinel's Chain Leggings](https://www.wowhead.com/forever/item=22748) | 65 | Lesser Arcanum of Voracity |
| Feet | [Scalegut Treaders](https://www.wowhead.com/forever/item=275618) | 58 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Molten Heart of the Mountain](https://www.wowhead.com/forever/item=249470) | 55 | — |
| Main hand | [Premier High Warlord's Blade](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Lifestealing |
| Off hand | [Premier High Warlord's Quickblade](https://www.wowhead.com/forever/item=272684) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4s: [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447).
- -2s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
3. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
4. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
5. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 160.90 |
| Auto-attack (tag 1) | 133.74 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 103.92 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 99.60 |
| Cat: Auto-attack (tag 1) | 83.61 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 68.91 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 52.77 |
| Auto-attack (tag 3) | 51.87 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11736.1 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4284.6 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -4129.5 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3719.1 |
| Mana | OtherActionManaRegen (tag 1) | +2464.0 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -2047.0 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | OtherActionManaRegen (tag 2) | +127.8 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 645.42 | 779.92 | +20.84% | 0.00 |
| Troll | 639.19 | 778.02 | +21.72% | 0.00 |
| Undead | 650.91 | 789.67 | +21.32% | 0.00 |
| Human | 640.27 | 777.60 | +21.45% | 0.00 |
| Gnome | 646.27 | 786.65 | +21.72% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherflame Cuffs](https://www.wowhead.com/forever/item=254065) | 50 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Fire Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Advisor's Ring](https://www.wowhead.com/forever/item=19518) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 75%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 60%).
3. Use ready automatic cooldowns.
4. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871) when Time remaining ≤ 1.5.
5. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
6. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 10%.
7. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
8. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) when [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309,"rank":8}}}`.
9. Cast [Conflagrate](https://www.wowhead.com/forever/spell=18932) when [Immolate](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ 15s.
10. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
11. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`).
12. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) when [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311,"rank":7}}}`.
13. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871).
14. Cast [Incinerate](https://www.wowhead.com/forever/spell=1293813).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 256.56 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 112.57 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 89.60 |
| Succubus: Auto-attack (tag 1) | 89.11 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 82.22 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 76.84 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 49.02 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 14.29 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24223.3 |
| Mana | [Incinerate](https://www.wowhead.com/forever/spell=1293813) | -21144.0 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6696.9 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -6271.9 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -5647.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5223.3 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5138.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4519.4 |
| Mana | OtherActionManaRegen (tag 1) | +3879.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3605.7 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -463.0 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-550500000002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 654.05 | 777.38 | +18.86% | 0.00 |
| Tauren | 652.71 | 775.24 | +18.77% | 0.00 |
| Troll | 650.56 | 771.87 | +18.65% | 0.00 |
| Undead | 663.26 | 780.02 | +17.60% | 0.00 |
| Windshaper | 650.59 | 773.31 | +18.86% | 0.00 |
| Human | 646.57 | 781.06 | +20.80% | 0.00 |
| Dwarf | 646.96 | 770.62 | +19.11% | 0.00 |
| Night Elf | 651.42 | 775.44 | +19.04% | 0.00 |
| Gnome | 648.91 | 772.09 | +18.98% | 0.00 |
| High Order | 650.59 | 773.31 | +18.86% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Lionheart Helm](https://www.wowhead.com/forever/item=12640) | 61 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Highlander's Plate Spaulders](https://www.wowhead.com/forever/item=20057) | 65 | Chromatic Mantle of the Dawn |
| Back | [Cloak of the Honor Guard](https://www.wowhead.com/forever/item=20073) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Gloves - Minor Haste |
| Waist | [Might of the Timbermaw](https://www.wowhead.com/forever/item=19044) | 58 | — |
| Legs | [Titanic Leggings](https://www.wowhead.com/forever/item=22385) | 60 | Lesser Arcanum of Voracity |
| Feet | [Scalegut Treaders](https://www.wowhead.com/forever/item=275618) | 58 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Protector's Band](https://www.wowhead.com/forever/item=19514) | 63 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Molten Heart of the Mountain](https://www.wowhead.com/forever/item=249470) | 55 | — |
| Main hand | [Premier High Warlord's Greatsword](https://www.wowhead.com/forever/item=272604) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when Time remaining ≤ 15s.
2. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719) when Time remaining ≤ 15s.
3. Use ready automatic cooldowns.
4. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when (NOT Time remaining ≤ 15s AND NOT [Recklessness](https://www.wowhead.com/forever/spell=1719) active).
5. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
6. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
7. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
8. Cast [Mortal Strike](https://www.wowhead.com/forever/spell=21553).
9. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
10. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 15.
11. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 70 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
12. Cast [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) when Rage ≥ 15.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 130.96 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 124.13 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 116.17 |
| Auto-attack (tag 3) | 110.34 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 98.72 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 92.60 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 45.01 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 33.68 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +2386.8 |
| Rage | [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | -1115.8 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -794.5 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -330.3 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -206.0 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -170.6 |
| Rage | [Rend](https://www.wowhead.com/forever/spell=11574) | -147.3 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -133.5 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +130.7 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +122.5 |
| Rage | OtherActionRefund | +113.5 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +96.7 |

## Priest — Shadow

**Talents:** 20/0/31 · `005300231303--505120501201300051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Troll | 637.18 | 761.03 | +19.44% | 0.00 |
| Undead | 646.27 | 770.64 | +19.25% | 0.00 |
| Human | 637.94 | 761.86 | +19.43% | 0.00 |
| Dwarf | 643.21 | 761.71 | +18.42% | 0.00 |
| Night Elf | 643.10 | 767.35 | +19.32% | 0.00 |
| Gnome | 642.86 | 767.86 | +19.44% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherpearl Cuffs](https://www.wowhead.com/forever/item=254067) | 50 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Shadow Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Advisor's Ring](https://www.wowhead.com/forever/item=19518) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -1s: [Shadowform](https://www.wowhead.com/forever/spell=15473).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) when NOT [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) DoT active.
3. Cast [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) when (NOT [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) DoT active AND Time remaining ≥ 10).
4. `{"condition":{},"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10947}}}]}}` when `{}`.
5. Cast [Mind Blast (rank 9)](https://www.wowhead.com/forever/spell=10947).
6. Cast [Mind Flay (rank 6)](https://www.wowhead.com/forever/spell=18807).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 369.74 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 153.25 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 141.37 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 94.72 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 11.55 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +9115.3 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -5904.6 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -5774.9 |
| Mana | [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | -4786.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3807.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2408.0 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2278.4 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1583.7 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1435.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +255.1 |
| Mana | OtherActionManaRegen (tag 2) | +0.1 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `252200001351-0025003001-0550005103`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 570.13 | 757.02 | +32.78% | 0.00 |
| Troll | 566.47 | 751.21 | +32.61% | 0.00 |
| Undead | 575.48 | 760.57 | +32.16% | 0.00 |
| Human | 576.82 | 748.17 | +29.71% | 0.00 |
| Gnome | 575.21 | 762.08 | +32.49% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Shadow Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Lorekeeper's Ring](https://www.wowhead.com/forever/item=19522) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -5s: [Demonic Sacrifice](https://www.wowhead.com/forever/spell=18788).
- -5s: [Amplify Curse](https://www.wowhead.com/forever/spell=18288).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 75%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 60%).
3. Use ready automatic cooldowns.
4. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871) when Time remaining ≤ 1.5.
5. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
6. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 10%.
7. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
8. Cast [Amplify Curse](https://www.wowhead.com/forever/spell=18288).
9. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
10. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`).
11. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) when [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311,"rank":7}}}`.
12. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) when [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309,"rank":8}}}`.
13. Cast [Shadow Bolt (rank 9)](https://www.wowhead.com/forever/spell=25307).

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 433.26 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 115.76 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 104.16 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 82.03 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 21.23 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 3.44 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.20 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -32373.4 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +22760.1 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6415.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5373.7 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5229.0 |
| Mana | OtherActionManaRegen (tag 1) | +3888.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3655.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3598.8 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -522.3 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -357.4 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -204.0 |

## Mage — Arcane

**Talents:** 33/3/15 · `053005023100311531-03-005500032`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 570.76 | 700.99 | +22.82% | 0.09 |
| Troll | 569.19 | 698.05 | +22.64% | 0.10 |
| Undead | 582.41 | 711.52 | +22.17% | 0.09 |
| Human | 576.77 | 696.69 | +20.79% | 0.08 |
| Gnome | 573.53 | 703.62 | +22.68% | 0.03 |
| High Order | 570.87 | 700.36 | +22.68% | 0.14 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Nethershine Cuffs](https://www.wowhead.com/forever/item=254069) | 50 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Advisor's Ring](https://www.wowhead.com/forever/item=19518) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 15% AND Time remaining > 20s).
3. Cast [Arcane Missiles (rank 8)](https://www.wowhead.com/forever/spell=25345) when [Missile Barrage](https://www.wowhead.com/forever/spell=44404) active.
4. Cast [Frostbolt (rank 11)](https://www.wowhead.com/forever/spell=25304) when ([Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks ≥ 4 OR Mana fraction < 20%).
5. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 340.75 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 321.60 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 33.57 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 15.60 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -23550.2 |
| Mana | OtherActionManaRegen (tag 1) | +6440.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +6373.8 |
| Mana | OtherActionManaRegen (tag 2) | +4118.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3459.6 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.1 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +844.7 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -770.4 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | -469.9 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +433.2 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Frost

**Talents:** 11/3/37 · `050005001-03-0555003321001301251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 567.95 | 694.11 | +22.21% | 0.07 |
| Troll | 569.12 | 694.70 | +22.06% | 0.06 |
| Undead | 572.97 | 697.96 | +21.82% | 0.07 |
| Human | 573.66 | 690.18 | +20.31% | 0.06 |
| Gnome | 570.76 | 695.98 | +21.94% | 0.03 |
| High Order | 570.30 | 696.08 | +22.06% | 0.05 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherfroth Cuffs](https://www.wowhead.com/forever/item=254063) | 50 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Frost Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Advisor's Ring](https://www.wowhead.com/forever/item=19518) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 80%.
2. Use ready automatic cooldowns.
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 15% AND Time remaining > 25s).
4. Cast [Ice Lance](https://www.wowhead.com/forever/spell=30455) when [Fingers of Frost](https://www.wowhead.com/forever/spell=44543) active.
5. Cast [Frostbolt (rank 11)](https://www.wowhead.com/forever/spell=25304).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 498.77 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 190.75 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 8.44 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -14132.7 |
| Mana | OtherActionManaRegen (tag 2) | +4133.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3229.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2637.0 |
| Mana | OtherActionManaRegen (tag 1) | +2423.3 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -2002.7 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.0 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +827.9 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +355.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `05043305-055030031005102051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 571.05 | 692.18 | +21.21% | 36.65 |
| Tauren | 562.94 | 693.88 | +23.26% | 29.74 |
| Troll | 563.80 | 688.24 | +22.07% | 28.32 |
| Windshaper | 564.63 | 689.51 | +22.12% | 32.24 |
| Dwarf | 560.21 | 695.88 | +24.22% | 32.14 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Dwarf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Chromatic Cloak](https://www.wowhead.com/forever/item=18509) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Tranquil Wristguards](https://www.wowhead.com/forever/item=279256) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Gloves - Minor Haste |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Lesser Arcanum of Voracity |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Main hand | [Premier High Warlord's Destroyer](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Fiery Weapon |
| Ranged/relic | [Totem of Thunder](https://www.wowhead.com/forever/item=228176) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4.5s: [Strength of Earth Totem (rank 4)](https://www.wowhead.com/forever/spell=10442).
- -3s: [Grace of Air Totem (rank 2)](https://www.wowhead.com/forever/spell=10627).
- -1.5s: [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438).

### Rotation priorities

1. Cast [Strength of Earth Totem (rank 4)](https://www.wowhead.com/forever/spell=10442) when Earth totem time remaining ≤ 0s.
2. Cast [Grace of Air Totem (rank 2)](https://www.wowhead.com/forever/spell=10627) when Air totem time remaining ≤ 0s.
3. Use ready automatic cooldowns.
4. Cast [Stormstrike (rank 1)](https://www.wowhead.com/forever/spell=17364).
5. Cast [Chain Lightning](https://www.wowhead.com/forever/spell=10605) when [Maelstrom Weapon](https://www.wowhead.com/forever/spell=408505) stacks ≥ 3.
6. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when [Maelstrom Weapon](https://www.wowhead.com/forever/spell=408505) stacks ≥ 3.
7. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Fire totem time remaining ≤ 0s AND Time remaining ≥ 5s).
8. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active.
9. Cast [Earth Shock (rank 7)](https://www.wowhead.com/forever/spell=10414).

### Damage breakdown — Dwarf

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 218.83 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 141.53 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 88.79 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 65.30 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 50.46 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 50.15 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 29.78 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 28.91 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8822.2 |
| Mana | [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | -8154.1 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -5467.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5138.4 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4696.4 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -4080.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3603.4 |
| Mana | OtherActionManaRegen (tag 1) | +3073.5 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -1601.6 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1005.4 |
| Mana | [Grace of Air Totem](https://www.wowhead.com/forever/spell=10627) | -453.8 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -431.1 |

## Rogue — Mutilate

**Talents:** 31/20/0 · `0053031035140105-302303202014`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 517.02 | 608.77 | +17.75% | 0.00 |
| Troll | 517.93 | 605.42 | +16.89% | 0.00 |
| Undead | 535.90 | 629.73 | +17.51% | 0.00 |
| Windshaper | 518.92 | 610.86 | +17.72% | 0.00 |
| Human | 514.14 | 604.68 | +17.61% | 0.00 |
| Dwarf | 513.01 | 603.82 | +17.70% | 0.00 |
| Night Elf | 520.95 | 613.38 | +17.74% | 0.00 |
| Gnome | 523.82 | 616.67 | +17.73% | 0.00 |
| High Order | 518.92 | 610.86 | +17.72% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Outlaw's Collar](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Deathguard's Cloak](https://www.wowhead.com/forever/item=20068) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Timbermaw Brawlers](https://www.wowhead.com/forever/item=19049) | 61 | Enchant Gloves - Minor Haste |
| Waist | [Defiler's Leather Girdle](https://www.wowhead.com/forever/item=20190) | 63 | — |
| Legs | [Outrider's Leather Pants](https://www.wowhead.com/forever/item=22740) | 65 | Lesser Arcanum of Voracity |
| Feet | [Prowler's Leather Boots](https://www.wowhead.com/forever/item=252468) | 52 | Enchant Boots - Greater Agility |
| Ring 1 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Ring 2 | [Legionnaire's Band](https://www.wowhead.com/forever/item=19510) | 63 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Molten Heart of the Mountain](https://www.wowhead.com/forever/item=249470) | 55 | — |
| Main hand | [Premier High Warlord's Razor](https://www.wowhead.com/forever/item=272596) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Premier High Warlord's Razor](https://www.wowhead.com/forever/item=272596) | 65 | Enchant Weapon - Fiery Weapon |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND `{"auraRemainingTime":{"auraId":{"spellId":6774,"rank":2}}}` < 3 AND Time remaining > 9)).
3. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
4. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when (Combo points ≥ 4 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active).
5. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 3.
6. Cast [Mutilate](https://www.wowhead.com/forever/spell=1241584).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 160.11 |
| Auto-attack (tag 2) | 96.60 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 71.49 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 65.27 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 59.79 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 44.76 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 38.68 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 28.68 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3014.2 |
| Energy | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | -2899.4 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -767.9 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +740.3 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -410.9 |
| Energy | OtherActionRefund | +152.2 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -125.4 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.7 |
| ComboPoints | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | +90.3 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +31.4 |
| ComboPoints | [Seal Fate](https://www.wowhead.com/forever/spell=14195) | +28.0 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -22.7 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 506.52 | 606.88 | +19.81% | 0.00 |
| Troll | 508.08 | 606.36 | +19.34% | 0.00 |
| Undead | 520.08 | 625.30 | +20.23% | 0.00 |
| Windshaper | 509.04 | 609.83 | +19.80% | 0.00 |
| Human | 511.69 | 611.01 | +19.41% | 0.00 |
| Dwarf | 506.01 | 600.78 | +18.73% | 0.00 |
| Night Elf | 509.29 | 608.79 | +19.54% | 0.00 |
| Gnome | 510.00 | 609.06 | +19.42% | 0.00 |
| High Order | 509.04 | 609.83 | +19.80% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Outlaw's Collar](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Deathguard's Cloak](https://www.wowhead.com/forever/item=20068) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Timbermaw Brawlers](https://www.wowhead.com/forever/item=19049) | 61 | Enchant Gloves - Minor Haste |
| Waist | [Might of the Timbermaw](https://www.wowhead.com/forever/item=19044) | 58 | — |
| Legs | [Outrider's Leather Pants](https://www.wowhead.com/forever/item=22740) | 65 | Lesser Arcanum of Voracity |
| Feet | [Prowler's Leather Boots](https://www.wowhead.com/forever/item=252468) | 52 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Molten Heart of the Mountain](https://www.wowhead.com/forever/item=249470) | 55 | — |
| Main hand | [Premier High Warlord's Right Claw](https://www.wowhead.com/forever/item=272597) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Premier High Warlord's Quickblade](https://www.wowhead.com/forever/item=272684) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -1s: [Blade Flurry](https://www.wowhead.com/forever/spell=13877).

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND `{"auraRemainingTime":{"auraId":{"spellId":6774,"rank":2}}}` < 3 AND Time remaining > 9)).
3. Cast [Adrenaline Rush](https://www.wowhead.com/forever/spell=13750) when Energy ≤ 40.
4. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
5. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when Combo points ≥ 3.
6. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 3.
7. Cast [Sinister Strike (rank 8)](https://www.wowhead.com/forever/spell=11294) when (`{"autoTimeToNext":{"autoType":"MainHand"}}` > `{"autoSwingTime":{"autoType":"MainHand"}}` − 0.5 OR Time remaining < 6 OR Energy ≥ 79).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 178.47 |
| [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | 136.68 |
| Auto-attack (tag 2) | 110.92 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 89.08 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 33.90 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 25.42 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 16.38 |
| Auto-attack (tag 3) | 12.40 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3176.1 |
| Energy | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | -2911.8 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -469.5 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -463.2 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +462.8 |
| Energy | OtherActionRefund | +103.4 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +96.5 |
| Energy | [Blade Flurry](https://www.wowhead.com/forever/spell=13877) | -75.0 |
| ComboPoints | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | +69.6 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -66.4 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -26.0 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +24.6 |

## Druid — Feral

**Talents:** 9/34/8 · `050022-3521002023032213041-053`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Tauren | 537.80 | 602.08 | +11.95% | 0.00 |
| Windshaper | 538.68 | 602.68 | +11.88% | 0.00 |
| Night Elf | 537.84 | 603.02 | +12.12% | 0.00 |
| High Order | 538.68 | 602.68 | +11.88% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Night Elf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Outlaw's Collar](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Shifting Cloak](https://www.wowhead.com/forever/item=18511) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Timbermaw Brawlers](https://www.wowhead.com/forever/item=19049) | 61 | Enchant Weapon - Agility |
| Waist | [Might of the Timbermaw](https://www.wowhead.com/forever/item=19044) | 58 | — |
| Legs | [Warbear Woolies](https://www.wowhead.com/forever/item=15065) | 57 | Lesser Arcanum of Voracity |
| Feet | [Mongoose Boots](https://www.wowhead.com/forever/item=18506) | 62 | Enchant Boots - Greater Agility |
| Ring 1 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Ring 2 | [Protector's Band](https://www.wowhead.com/forever/item=19514) | 63 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Molten Heart of the Mountain](https://www.wowhead.com/forever/item=249470) | 55 | — |
| Main hand | [Premier High Warlord's Pulverizer](https://www.wowhead.com/forever/item=272601) | 65 | Enchant 2H Weapon - Agility |
| Ranged/relic | [Idol of the Dream](https://www.wowhead.com/forever/item=220606) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) when Energy ≤ 40.
3. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Time remaining ≤ 4s AND Combo points ≥ 3).
4. Cast [Rip](https://www.wowhead.com/forever/spell=9896) when (Combo points ≥ 4 AND NOT [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND Time remaining ≥ 8s).
5. Cast [Shred](https://www.wowhead.com/forever/spell=9830).

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 227.43 |
| [Shred](https://www.wowhead.com/forever/spell=9830) | 192.46 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 176.00 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 4.62 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.51 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -3313.3 |
| Energy | OtherActionEnergyRegen | +3029.3 |
| Energy | [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) | +719.5 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -634.1 |
| Energy | OtherActionRefund | +172.4 |
| ComboPoints | [Rip](https://www.wowhead.com/forever/spell=9896) | -92.8 |
| ComboPoints | [Shred](https://www.wowhead.com/forever/spell=9830) | +63.1 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -56.2 |
| ComboPoints | [Primal Fury](https://www.wowhead.com/forever/spell=37117) | +34.4 |
| ComboPoints | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -3.5 |
| Mana | OtherActionManaRegen (tag 2) | +0.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +0.0 |

## Priest — Smite

**Talents:** 31/17/3 · `515030031305001031-00505023002-003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Troll | 468.35 | 594.41 | +26.91% | 0.00 |
| Undead | 474.85 | 597.77 | +25.89% | 0.00 |
| Human | 467.33 | 592.97 | +26.88% | 0.00 |
| Dwarf | 468.52 | 592.22 | +26.40% | 0.00 |
| Night Elf | 469.40 | 595.57 | +26.88% | 0.00 |
| Gnome | 473.32 | 597.41 | +26.22% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Ghostweave Hood](https://www.wowhead.com/forever/item=254135) | 61 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Argent Shoulders](https://www.wowhead.com/forever/item=19059) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherlight Cuffs](https://www.wowhead.com/forever/item=254071) | 50 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Advisor's Ring](https://www.wowhead.com/forever/item=19518) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Brilliant Wand](https://www.wowhead.com/forever/item=249385) | 60 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -3s: [Holy Fire (rank 8)](https://www.wowhead.com/forever/spell=15261).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Penance](https://www.wowhead.com/forever/spell=1316995).
3. Cast [Holy Fire (rank 8)](https://www.wowhead.com/forever/spell=15261) when [Holy Fire (rank 8)](https://www.wowhead.com/forever/spell=15261) DoT time remaining < 3s.
4. `{"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10934,"rank":8}}}]}}`.
5. Cast [Smite](https://www.wowhead.com/forever/spell=10934) when Mana ≥ Time remaining × 20.
6. Cast [Smite](https://www.wowhead.com/forever/spell=598).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Smite](https://www.wowhead.com/forever/spell=10934) | 232.94 |
| [Penance](https://www.wowhead.com/forever/spell=1316995) | 131.21 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 123.11 |
| [Smite](https://www.wowhead.com/forever/spell=598) | 102.29 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 8.22 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -14503.2 |
| Mana | OtherActionManaRegen (tag 1) | +9243.1 |
| Mana | [Penance](https://www.wowhead.com/forever/spell=1316995) | -7124.0 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -5566.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4039.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3602.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.0 |
| Mana | [Smite](https://www.wowhead.com/forever/spell=598) | -1727.8 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1529.6 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |

## Druid — Balance

**Talents:** 38/0/13 · `5232220115501351--505003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Tauren | 458.37 | 588.75 | +28.44% | 0.01 |
| Windshaper | 455.98 | 586.34 | +28.59% | 0.02 |
| Night Elf | 458.26 | 588.71 | +28.47% | 0.02 |
| High Order | 455.98 | 586.34 | +28.59% | 0.02 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Tauren

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Living Crown](https://www.wowhead.com/forever/item=252561) | 61 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of the Greatfather](https://www.wowhead.com/forever/item=17721) | 38 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Advisor's Ring](https://www.wowhead.com/forever/item=19518) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Swarming Idol](https://www.wowhead.com/forever/item=272430) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -2s: [Moonkin Form](https://www.wowhead.com/forever/spell=24858).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Insect Swarm (rank 5)](https://www.wowhead.com/forever/spell=24977) when NOT [Insect Swarm (rank 5)](https://www.wowhead.com/forever/spell=24977) DoT active.
3. Cast [Moonfire (rank 10)](https://www.wowhead.com/forever/spell=9835) when NOT [Moonfire (rank 10)](https://www.wowhead.com/forever/spell=9835) DoT active.
4. Cast [Starfire (rank 7)](https://www.wowhead.com/forever/spell=25298) when [Eclipse](https://www.wowhead.com/forever/spell=48518) stacks ≥ 1.
5. Cast [Wrath (rank 8)](https://www.wowhead.com/forever/spell=9912).

### Damage breakdown — Tauren

| Action | DPS |
|---|---:|
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 333.84 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 89.50 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 85.34 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 80.07 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=25298) | -19207.4 |
| Mana | OtherActionManaRegen (tag 1) | +10465.2 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -5321.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4468.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4224.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3593.5 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1885.1 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1314.0 |
| Mana | [Moonkin Form](https://www.wowhead.com/forever/spell=24858) | -326.6 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | -46.6 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 449.12 | 548.74 | +22.18% | 0.00 |
| Troll | 444.58 | 542.07 | +21.93% | 0.00 |
| Undead | 463.09 | 572.27 | +23.58% | 0.00 |
| Windshaper | 444.87 | 540.29 | +21.45% | 0.00 |
| Human | 453.82 | 554.35 | +22.15% | 0.00 |
| Dwarf | 448.78 | 542.75 | +20.94% | 0.00 |
| Night Elf | 451.52 | 551.61 | +22.17% | 0.00 |
| Gnome | 450.71 | 546.96 | +21.36% | 0.00 |
| High Order | 444.87 | 540.29 | +21.45% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Outlaw's Collar](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Deathguard's Cloak](https://www.wowhead.com/forever/item=20068) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Timbermaw Brawlers](https://www.wowhead.com/forever/item=19049) | 61 | Enchant Weapon - Agility |
| Waist | [Might of the Timbermaw](https://www.wowhead.com/forever/item=19044) | 58 | — |
| Legs | [Outrider's Leather Pants](https://www.wowhead.com/forever/item=22740) | 65 | Lesser Arcanum of Voracity |
| Feet | [Prowler's Leather Boots](https://www.wowhead.com/forever/item=252468) | 52 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Trinket 2 | [Molten Heart of the Mountain](https://www.wowhead.com/forever/item=249470) | 55 | — |
| Main hand | [Premier High Warlord's Razor](https://www.wowhead.com/forever/item=272596) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Premier High Warlord's Quickblade](https://www.wowhead.com/forever/item=272684) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND `{"auraRemainingTime":{"auraId":{"spellId":6774,"rank":2}}}` < 3 AND Time remaining > 9)).
2. Cast [Vanish](https://www.wowhead.com/forever/spell=1856) when (Energy ≥ 60 AND Combo points ≤ 1 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining > 10).
3. Cast [Premeditation](https://www.wowhead.com/forever/spell=14183).
4. Cast [Ambush](https://www.wowhead.com/forever/spell=11269).
5. Cast [Rupture](https://www.wowhead.com/forever/spell=11275) when (Combo points ≥ 4 AND NOT [Rupture](https://www.wowhead.com/forever/spell=11275) DoT active AND Time remaining > 8).
6. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
7. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when (Combo points ≥ 4 AND ([Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active OR Energy ≥ 79 OR Time remaining < 6)).
8. Cast [Hemorrhage](https://www.wowhead.com/forever/spell=16511).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 166.89 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 133.38 |
| Auto-attack (tag 2) | 83.92 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 49.88 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 36.52 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 27.97 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 24.62 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 11.70 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +2999.2 |
| Energy | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | -2950.5 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +534.8 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -353.1 |
| Energy | [Rupture](https://www.wowhead.com/forever/spell=11275) | -336.3 |
| Energy | OtherActionRefund | +152.7 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -126.4 |
| Energy | [Ambush](https://www.wowhead.com/forever/spell=11269) | -91.1 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +88.3 |
| ComboPoints | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | +85.9 |
| ComboPoints | [Rupture](https://www.wowhead.com/forever/spell=11275) | -51.1 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -39.0 |

## Shaman — Stormcaller

**Talents:** 28/23/0 · `550032150010303-055030031004002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 413.71 | 523.09 | +26.44% | 1.39 |
| Tauren | 414.07 | 522.41 | +26.16% | 1.40 |
| Troll | 411.42 | 518.54 | +26.04% | 1.41 |
| Windshaper | 413.23 | 520.17 | +25.88% | 1.46 |
| Dwarf | 414.16 | 518.69 | +25.24% | 1.44 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Blue Dragonscale Helm](https://www.wowhead.com/forever/item=252604) | 61 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of the Greatfather](https://www.wowhead.com/forever/item=17721) | 38 | Enchant Gloves - Fire Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Ring 2 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Burning Totem](https://www.wowhead.com/forever/item=272433) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -2s: [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Rotation priorities

1. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Target count = 1 AND NOT [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) DoT active).
2. Cast [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) when (Target count ≥ 2 AND NOT [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) DoT active).
3. Use ready automatic cooldowns.
4. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active.
5. Cast [Chain Lightning (rank 4)](https://www.wowhead.com/forever/spell=10605) when ([Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active OR Target count ≥ 2).
6. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when Mana ≥ Time remaining × 10.
7. Cast [Lightning Bolt (rank 4)](https://www.wowhead.com/forever/spell=915).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 265.48 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=915) | 95.77 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 82.63 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 33.08 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 26.73 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 13.28 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=915) | 4.79 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.33 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -16366.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5279.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4905.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3597.4 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3216.1 |
| Mana | OtherActionManaRegen (tag 1) | +3074.7 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=915) | -2029.8 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1019.7 |
| Mana | OtherActionManaRegen (tag 2) | +164.9 |

## Shaman — Elemental

**Talents:** 31/6/14 · `5502301500123031-0500001-053050001`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 395.40 | 502.63 | +27.12% | 0.12 |
| Tauren | 395.80 | 502.30 | +26.91% | 0.13 |
| Troll | 393.48 | 500.46 | +27.19% | 0.12 |
| Windshaper | 393.85 | 499.82 | +26.91% | 0.11 |
| Dwarf | 395.22 | 498.00 | +26.01% | 0.13 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Blue Dragonscale Helm](https://www.wowhead.com/forever/item=252604) | 61 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Greater Intellect |
| Hands | [Gloves of the Greatfather](https://www.wowhead.com/forever/item=17721) | 38 | Enchant Gloves - Fire Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Advisor's Ring](https://www.wowhead.com/forever/item=19518) | 63 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Spellblade](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Premier High Warlord's Tome of Destruction](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Burning Totem](https://www.wowhead.com/forever/item=272433) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -2s: [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Rotation priorities

1. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Target count = 1 AND NOT [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) DoT active).
2. Cast [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) when (Target count ≥ 2 AND NOT [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) DoT active).
3. Use ready automatic cooldowns.
4. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active.
5. Cast [Lava Burst (rank 3)](https://www.wowhead.com/forever/spell=1238300).
6. Cast [Chain Lightning (rank 4)](https://www.wowhead.com/forever/spell=10605) when ([Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active OR Target count ≥ 2).
7. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when Mana ≥ Time remaining × 35.
8. Cast [Lightning Bolt (rank 6)](https://www.wowhead.com/forever/spell=6041).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 179.49 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | 96.99 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 80.29 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 78.64 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 32.56 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 20.68 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 8.94 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 4.01 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +7281.5 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | -6927.6 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -6575.7 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | -5449.7 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -5412.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4826.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3592.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3566.0 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -765.0 |
| Mana | OtherActionManaRegen (tag 2) | +162.2 |
