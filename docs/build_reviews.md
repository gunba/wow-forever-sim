# Build reviews

These summaries describe the published loadouts. They are simulation results, not independent confirmation of server mechanics or proof of a global optimum.

The tables and [matrix](../artifacts/forever_dps_5min.png) use the same 147 common-seed replays. Equipment selections came from an earlier mechanics revision; these results use the corrected engine. Historical search gains are not directly comparable to this release.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and paid shared-hit normalization. [Scenario and exchange model](../tools/forever_bench/README.md) · [In-game checks](in_game_checks.md)

## Paladin — Retribution

**Talents:** 12/0/39 · `250003002--052253312012331321`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1025.38 | 0.46 | 0.00 |
| Human | 1004.48 | 0.46 | 0.00 |
| Dwarf | 1002.87 | 0.45 | 0.00 |

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
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 246.01 |
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 226.85 |
| Auto-attack (tag 1) | 191.64 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 91.85 |
| Auto-attack (tag 3) | 71.45 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20966) | 61.30 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 54.97 |
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

## Warrior — Fury

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 951.37 | 0.53 | 0.00 |
| Tauren | 933.65 | 0.49 | 0.00 |
| Troll | 935.26 | 0.52 | 0.00 |
| Undead | 953.82 | 0.50 | 0.00 |
| Windshaper | 932.68 | 0.50 | 0.00 |
| Human | 946.13 | 0.51 | 0.00 |
| Dwarf | 934.15 | 0.49 | 0.00 |
| Night Elf | 932.10 | 0.51 | 0.00 |
| Gnome | 923.79 | 0.51 | 0.00 |
| High Order | 932.68 | 0.50 | 0.00 |

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

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 882.91 | 0.30 | 0.00 |
| Tauren | 872.88 | 0.29 | 0.01 |
| Troll | 877.98 | 0.29 | 0.00 |
| Windshaper | 878.64 | 0.30 | 0.00 |
| Human | 885.70 | 0.30 | 0.00 |
| Dwarf | 870.38 | 0.29 | 0.01 |
| Night Elf | 881.50 | 0.30 | 0.00 |
| High Order | 878.64 | 0.30 | 0.00 |

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
| Shoot | 382.34 |
| Cat: Auto-attack (tag 1) | 149.19 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | 129.81 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 83.57 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 67.66 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 49.48 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 13.17 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 10.47 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | -9191.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7649.5 |
| Mana | OtherActionManaRegen (tag 1) | +6124.9 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5528.6 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -4664.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3600.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3597.9 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3158.5 |
| Mana | OtherActionManaRegen (tag 2) | +846.2 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Intimidation](https://www.wowhead.com/forever/spell=19577) | -407.1 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Hunter — Marksmanship

**Talents:** 18/33/0 · `5023000503-0053451001503051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 852.16 | 0.32 | 0.03 |
| Tauren | 841.39 | 0.32 | 0.03 |
| Troll | 847.00 | 0.32 | 0.03 |
| Windshaper | 847.37 | 0.32 | 0.02 |
| Human | 854.77 | 0.32 | 0.03 |
| Dwarf | 838.68 | 0.32 | 0.03 |
| Night Elf | 850.92 | 0.32 | 0.03 |
| High Order | 847.37 | 0.32 | 0.02 |

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
| Shoot | 401.67 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20901) | 129.72 |
| Cat: Auto-attack (tag 1) | 103.26 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 73.12 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 70.18 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 60.45 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 12.57 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 3.80 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8212.6 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -6784.0 |
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20901) | -5784.5 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -4920.8 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3739.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3603.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3576.0 |
| Mana | OtherActionManaRegen (tag 1) | +2926.8 |
| Mana | OtherActionManaRegen (tag 2) | +691.6 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 812.45 | 0.42 | 0.00 |
| Troll | 804.23 | 0.42 | 0.00 |
| Undead | 814.48 | 0.41 | 0.00 |
| Human | 803.89 | 0.41 | 0.00 |
| Gnome | 816.48 | 0.42 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Healing Power |
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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 425.27 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 97.06 |
| Succubus: Auto-attack (tag 1) | 91.70 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 89.70 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 85.62 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 18.02 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 5.73 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.96 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -32241.2 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +13101.7 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +11005.5 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6356.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5372.1 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5223.2 |
| Mana | OtherActionManaRegen (tag 1) | +3890.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3611.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2635.2 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -492.7 |
| Mana | [Soul Link](https://www.wowhead.com/forever/spell=19028) | -274.6 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2525000013520105--0550015103`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 804.41 | 0.46 | 0.01 |
| Troll | 798.42 | 0.45 | 0.01 |
| Undead | 805.16 | 0.45 | 0.01 |
| Human | 796.51 | 0.44 | 0.01 |
| Gnome | 809.01 | 0.46 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Healing Power |
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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 372.32 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 105.73 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 98.62 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 84.08 |
| Succubus: Auto-attack (tag 1) | 80.93 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 39.67 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 19.00 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 3.84 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -30125.4 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24916.7 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6455.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5254.4 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5109.4 |
| Mana | OtherActionManaRegen (tag 1) | +3869.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3425.9 |
| Mana | [Siphon Life](https://www.wowhead.com/forever/spell=18881) | -3386.4 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.1 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -355.0 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -340.1 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 784.98 | 0.31 | 0.00 |
| Troll | 782.16 | 0.30 | 0.00 |
| Undead | 795.20 | 0.31 | 0.00 |
| Human | 782.58 | 0.31 | 0.00 |
| Gnome | 791.12 | 0.31 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherflame Cuffs](https://www.wowhead.com/forever/item=254065) | 50 | Enchant Bracer - Healing Power |
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
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 258.00 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 113.24 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 90.34 |
| Succubus: Auto-attack (tag 1) | 89.25 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 83.30 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 77.55 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 49.35 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 14.51 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24361.3 |
| Mana | [Incinerate](https://www.wowhead.com/forever/spell=1293813) | -21116.6 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6694.8 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -6281.8 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -5662.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5218.6 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5164.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4513.4 |
| Mana | OtherActionManaRegen (tag 1) | +3878.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3602.6 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -464.7 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 779.73 | 0.45 | 0.00 |
| Tauren | 768.42 | 0.45 | 0.00 |
| Troll | 772.18 | 0.45 | 0.00 |
| Windshaper | 769.72 | 0.45 | 0.00 |
| Human | 783.81 | 0.46 | 0.00 |
| Dwarf | 764.56 | 0.44 | 0.00 |
| Night Elf | 777.05 | 0.45 | 0.00 |
| High Order | 769.72 | 0.45 | 0.00 |

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
| Auto-attack (tag 2) | 160.85 |
| Auto-attack (tag 1) | 133.97 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 103.80 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 88.54 |
| Cat: Auto-attack (tag 1) | 83.58 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 69.09 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 53.03 |
| Auto-attack (tag 3) | 50.93 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11524.9 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4283.2 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -4138.4 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3721.0 |
| Mana | OtherActionManaRegen (tag 1) | +2438.0 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -1813.3 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | OtherActionManaRegen (tag 2) | +142.6 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-550500000002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 777.38 | 0.53 | 0.00 |
| Tauren | 775.24 | 0.52 | 0.00 |
| Troll | 771.87 | 0.52 | 0.00 |
| Undead | 780.02 | 0.53 | 0.00 |
| Windshaper | 773.31 | 0.52 | 0.00 |
| Human | 781.06 | 0.52 | 0.00 |
| Dwarf | 770.62 | 0.52 | 0.00 |
| Night Elf | 775.44 | 0.52 | 0.00 |
| Gnome | 772.09 | 0.53 | 0.00 |
| High Order | 773.31 | 0.52 | 0.00 |

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

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 763.95 | 0.20 | 0.00 |
| Undead | 773.65 | 0.21 | 0.00 |
| Human | 765.06 | 0.21 | 0.00 |
| Dwarf | 764.74 | 0.20 | 0.00 |
| Night Elf | 770.47 | 0.21 | 0.00 |
| Gnome | 771.00 | 0.21 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherpearl Cuffs](https://www.wowhead.com/forever/item=254067) | 50 | Enchant Bracer - Healing Power |
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

- -1.5s: [Shadowform](https://www.wowhead.com/forever/spell=15473).

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
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 371.09 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 153.92 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 142.02 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 95.09 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 11.53 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +9107.1 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -5907.2 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -5775.1 |
| Mana | [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | -4787.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3809.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2503.5 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2279.8 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1590.5 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1506.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +990.8 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `252200001351-0025003001-0550005103`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 764.74 | 0.50 | 0.00 |
| Troll | 759.06 | 0.48 | 0.00 |
| Undead | 767.94 | 0.49 | 0.00 |
| Human | 755.90 | 0.49 | 0.00 |
| Gnome | 770.10 | 0.50 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Healing Power |
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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 439.31 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 116.38 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 104.69 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 82.60 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 21.45 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 3.46 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.22 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -32382.6 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +22906.5 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6415.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5363.0 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5224.3 |
| Mana | OtherActionManaRegen (tag 1) | +3887.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3656.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.1 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -525.5 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -352.3 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -204.2 |

## Mage — Arcane

**Talents:** 33/3/15 · `053005023100311531-03-005500032`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 681.31 | 0.59 | 1.32 |
| Troll | 678.37 | 0.59 | 1.62 |
| Undead | 692.62 | 0.60 | 1.32 |
| Human | 678.04 | 0.58 | 1.16 |
| Gnome | 687.20 | 0.55 | 0.60 |
| High Order | 680.25 | 0.58 | 1.55 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Nethershine Cuffs](https://www.wowhead.com/forever/item=254069) | 50 | Enchant Bracer - Healing Power |
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
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 322.77 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 296.08 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 58.80 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 14.96 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -28426.5 |
| Mana | OtherActionManaRegen (tag 1) | +8824.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +6831.3 |
| Mana | OtherActionManaRegen (tag 2) | +4708.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4496.1 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -2431.0 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.9 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +850.5 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | -631.6 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +572.7 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Druid — Feral

**Talents:** 9/34/8 · `050022-3521002023032213041-053`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 671.39 | 0.21 | 0.00 |
| Windshaper | 673.60 | 0.21 | 0.00 |
| Night Elf | 672.47 | 0.21 | 0.00 |
| High Order | 673.60 | 0.21 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Outlaw's Collar](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Shifting Cloak](https://www.wowhead.com/forever/item=18511) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Timbermaw Brawlers](https://www.wowhead.com/forever/item=19049) | 61 | Enchant Weapon - Agility |
| Waist | [Might of the Timbermaw](https://www.wowhead.com/forever/item=19044) | 58 | — |
| Legs | [Warbear Woolies](https://www.wowhead.com/forever/item=15065) | 57 | Lesser Arcanum of Voracity |
| Feet | [Mongoose Boots](https://www.wowhead.com/forever/item=18506) | 62 | Enchant Boots - Greater Agility |
| Ring 1 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Ring 2 | [Legionnaire's Band](https://www.wowhead.com/forever/item=19510) | 63 | — |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Shred](https://www.wowhead.com/forever/spell=9830) | 258.05 |
| Auto-attack (tag 1) | 226.79 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 181.50 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 4.64 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.62 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -3346.5 |
| Energy | OtherActionEnergyRegen | +3057.9 |
| Energy | [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) | +690.2 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -597.0 |
| Energy | OtherActionRefund | +172.3 |
| ComboPoints | [Rip](https://www.wowhead.com/forever/spell=9896) | -98.8 |
| ComboPoints | [Shred](https://www.wowhead.com/forever/spell=9830) | +66.3 |
| ComboPoints | [Primal Fury](https://www.wowhead.com/forever/spell=37117) | +37.7 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -35.2 |
| ComboPoints | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -4.2 |
| Mana | OtherActionManaRegen (tag 2) | +0.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +0.0 |

## Mage — Frost

**Talents:** 11/3/37 · `050005001-03-0555003321001301251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 627.50 | 0.52 | 24.07 |
| Troll | 627.05 | 0.50 | 24.67 |
| Undead | 633.43 | 0.52 | 23.42 |
| Human | 627.42 | 0.52 | 22.48 |
| Gnome | 654.60 | 0.52 | 12.91 |
| High Order | 628.13 | 0.51 | 24.60 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Bloodvine Goggles](https://www.wowhead.com/forever/item=19999) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherfroth Cuffs](https://www.wowhead.com/forever/item=254063) | 50 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Frost Power |
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

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 80%.
2. Use ready automatic cooldowns.
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 15% AND Time remaining > 25s).
4. Cast [Ice Lance](https://www.wowhead.com/forever/spell=30455) when [Fingers of Frost](https://www.wowhead.com/forever/spell=44543) active.
5. Cast [Frostbolt (rank 11)](https://www.wowhead.com/forever/spell=25304).

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 473.45 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 181.15 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -21537.5 |
| Mana | OtherActionManaRegen (tag 1) | +5914.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4563.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3554.9 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3347.3 |
| Mana | OtherActionManaRegen (tag 2) | +2645.6 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.5 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.9 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +600.5 |

## Rogue — Mutilate

**Talents:** 31/20/0 · `0053031035140105-302303202014`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 618.45 | 0.29 | 0.00 |
| Troll | 615.26 | 0.29 | 0.00 |
| Undead | 639.28 | 0.29 | 0.00 |
| Windshaper | 620.87 | 0.29 | 0.00 |
| Human | 614.49 | 0.28 | 0.00 |
| Dwarf | 613.40 | 0.29 | 0.00 |
| Night Elf | 623.19 | 0.29 | 0.00 |
| Gnome | 627.36 | 0.30 | 0.00 |
| High Order | 620.87 | 0.29 | 0.00 |

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
| Auto-attack (tag 1) | 159.66 |
| Auto-attack (tag 2) | 96.32 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 72.22 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 70.48 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 59.70 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 44.68 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 43.34 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 28.69 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3013.9 |
| Energy | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | -2890.3 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -795.8 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +759.9 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -409.1 |
| Energy | OtherActionRefund | +149.8 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -130.4 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.7 |
| ComboPoints | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | +90.1 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +32.0 |
| ComboPoints | [Seal Fate](https://www.wowhead.com/forever/spell=14195) | +31.9 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -22.0 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 606.88 | 0.27 | 0.00 |
| Troll | 606.36 | 0.27 | 0.00 |
| Undead | 625.30 | 0.27 | 0.00 |
| Windshaper | 609.83 | 0.27 | 0.00 |
| Human | 611.01 | 0.28 | 0.00 |
| Dwarf | 600.78 | 0.27 | 0.00 |
| Night Elf | 608.79 | 0.27 | 0.00 |
| Gnome | 609.06 | 0.28 | 0.00 |
| High Order | 609.83 | 0.27 | 0.00 |

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

## Mage — Fire

**Talents:** 17/31/3 · `0501252000002-23450000130133051-003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 605.27 | 0.50 | 32.01 |
| Troll | 598.31 | 0.49 | 32.05 |
| Undead | 620.02 | 0.52 | 29.66 |
| Human | 612.52 | 0.52 | 28.34 |
| Gnome | 622.15 | 0.51 | 24.70 |
| High Order | 595.12 | 0.49 | 34.61 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

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
| Ring 2 | [Lorekeeper's Ring](https://www.wowhead.com/forever/item=19522) | 63 | — |
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

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 273.49 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 135.60 |
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 129.79 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 78.17 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 5.09 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -15732.0 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -11518.2 |
| Mana | OtherActionManaRegen (tag 1) | +9462.2 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -5674.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5265.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5127.5 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +3279.1 |
| Mana | OtherActionManaRegen (tag 2) | +1732.0 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1098.6 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +850.2 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +600.2 |
| Mana | [Fireball](https://www.wowhead.com/forever/spell=25306) | -400.7 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `05043305-055030031005102051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 601.88 | 0.50 | 45.86 |
| Tauren | 602.01 | 0.49 | 40.22 |
| Troll | 596.65 | 0.48 | 39.18 |
| Windshaper | 599.77 | 0.48 | 42.36 |
| Dwarf | 605.13 | 0.49 | 42.35 |

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
| Auto-attack (tag 1) | 167.68 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 117.79 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 87.74 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 59.59 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 48.94 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 48.12 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 29.55 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 26.91 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7666.0 |
| Mana | [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | -7536.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5340.7 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -5179.6 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4661.6 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3992.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3602.1 |
| Mana | OtherActionManaRegen (tag 1) | +3060.7 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -1792.4 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -999.1 |
| Mana | [Grace of Air Totem](https://www.wowhead.com/forever/spell=10627) | -478.1 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -443.4 |

## Druid — Balance

**Talents:** 38/0/13 · `5232220115501351--505003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 603.87 | 0.26 | 0.00 |
| Windshaper | 603.28 | 0.26 | 0.00 |
| Night Elf | 604.08 | 0.26 | 0.00 |
| High Order | 603.28 | 0.26 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Night Elf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Living Crown](https://www.wowhead.com/forever/item=252561) | 61 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Healing Power |
| Hands | [Gloves of the Greatfather](https://www.wowhead.com/forever/item=17721) | 38 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Lorekeeper's Ring](https://www.wowhead.com/forever/item=19522) | 63 | — |
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

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 351.92 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 91.81 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 86.25 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 74.09 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=25298) | -14990.4 |
| Mana | OtherActionManaRegen (tag 1) | +7736.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4478.9 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -4231.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3596.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3595.0 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1569.8 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1199.5 |
| Mana | [Moonkin Form](https://www.wowhead.com/forever/spell=24858) | -435.4 |
| Mana | OtherActionManaRegen (tag 2) | +75.1 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | +1.0 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | -0.0 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 548.74 | 0.21 | 0.00 |
| Troll | 542.07 | 0.22 | 0.00 |
| Undead | 572.27 | 0.23 | 0.00 |
| Windshaper | 540.29 | 0.22 | 0.00 |
| Human | 554.35 | 0.21 | 0.00 |
| Dwarf | 542.75 | 0.21 | 0.00 |
| Night Elf | 551.61 | 0.21 | 0.00 |
| Gnome | 546.96 | 0.22 | 0.00 |
| High Order | 540.29 | 0.22 | 0.00 |

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

## Priest — Smite

**Talents:** 31/17/3 · `515030031305001031-00505023002-003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 551.42 | 0.15 | 0.00 |
| Undead | 555.30 | 0.15 | 0.00 |
| Human | 550.13 | 0.15 | 0.00 |
| Dwarf | 549.36 | 0.15 | 0.00 |
| Night Elf | 552.57 | 0.15 | 0.00 |
| Gnome | 554.27 | 0.15 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Ghostweave Hood](https://www.wowhead.com/forever/item=254135) | 61 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Argent Shoulders](https://www.wowhead.com/forever/item=19059) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Netherlight Cuffs](https://www.wowhead.com/forever/item=254071) | 50 | Enchant Bracer - Healing Power |
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
| [Smite](https://www.wowhead.com/forever/spell=10934) | 213.97 |
| [Penance](https://www.wowhead.com/forever/spell=1316995) | 121.36 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 114.27 |
| [Smite](https://www.wowhead.com/forever/spell=598) | 97.47 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 8.23 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -14335.8 |
| Mana | OtherActionManaRegen (tag 1) | +9243.1 |
| Mana | [Penance](https://www.wowhead.com/forever/spell=1316995) | -7123.5 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -5566.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4040.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3602.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.1 |
| Mana | [Smite](https://www.wowhead.com/forever/spell=598) | -1764.0 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1528.7 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |

## Shaman — Stormcaller

**Talents:** 28/23/0 · `550032150010303-055030031004002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 529.44 | 0.22 | 1.42 |
| Tauren | 528.58 | 0.22 | 1.43 |
| Troll | 524.93 | 0.23 | 1.44 |
| Windshaper | 526.57 | 0.23 | 1.47 |
| Dwarf | 525.02 | 0.22 | 1.41 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Blue Dragonscale Helm](https://www.wowhead.com/forever/item=252604) | 61 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Healing Power |
| Hands | [Gloves of the Greatfather](https://www.wowhead.com/forever/item=17721) | 38 | Enchant Gloves - Healing Power |
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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 267.14 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=915) | 100.05 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 82.36 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 33.02 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 27.18 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 13.32 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=915) | 5.01 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.35 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -16197.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5274.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4902.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3598.9 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3213.4 |
| Mana | OtherActionManaRegen (tag 1) | +3074.7 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=915) | -2075.9 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1019.7 |
| Mana | OtherActionManaRegen (tag 2) | +164.9 |

## Shaman — Elemental

**Talents:** 31/7/13 · `5502301500123031-052-05305`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 510.58 | 0.21 | 0.12 |
| Tauren | 510.16 | 0.21 | 0.12 |
| Troll | 508.40 | 0.21 | 0.11 |
| Windshaper | 507.86 | 0.21 | 0.11 |
| Dwarf | 505.84 | 0.21 | 0.13 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Blue Dragonscale Helm](https://www.wowhead.com/forever/item=252604) | 61 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Runecloth Cuffs](https://www.wowhead.com/forever/item=254123) | 59 | Enchant Bracer - Healing Power |
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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 182.14 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | 97.07 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 84.53 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 78.71 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 32.58 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 21.19 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 9.08 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 4.22 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +7281.5 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | -6849.2 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -6577.3 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -5567.8 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | -5450.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4826.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3593.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3566.4 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -765.0 |
| Mana | OtherActionManaRegen (tag 2) | +162.3 |
