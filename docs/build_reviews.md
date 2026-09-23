# Build reviews

These summaries describe the published loadouts. They are simulation results, not independent confirmation of server mechanics or proof of a global optimum.

The tables and [matrix](../artifacts/forever_dps_5min.png) use the same 171 common-seed replays. Equipment selections came from an earlier mechanics revision; these results use the corrected engine. Historical search gains are not directly comparable to this release.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and paid shared-hit normalization. [Scenario and exchange model](../tools/forever_bench/README.md) · [In-game checks](in_game_checks.md)

## Paladin — Retribution

**Talents:** 13/7/31 · `253003-232-052052310012330301`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Separate seal Echoes can coexist and fire on a landed white swing; verify their simultaneous behavior in game (T53). Lower-rank seals and Consecration also need confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1141.49 | 0.62 | 0.21 |
| Human | 1118.99 | 0.61 | 0.74 |
| Dwarf | 1130.07 | 0.61 | 0.32 |

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
| Hands | [Raider Handwraps](https://www.wowhead.com/forever/item=272097) | 65 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Rapidity |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Destroyer](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Libram of Law](https://www.wowhead.com/forever/item=272435) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -1.5s: [Spell 20919](https://www.wowhead.com/forever/spell=20919).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Holy Strike](https://www.wowhead.com/forever/spell=10333).
3. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271).
4. Cast [Spell 20919](https://www.wowhead.com/forever/spell=20919) when `{"currentSealRemainingTime":{}}` ≤ 1s.
5. Cast [Spell 20293](https://www.wowhead.com/forever/spell=20293) when (Mana fraction ≥ 15% AND NOT [Echo of Command](https://www.wowhead.com/forever/spell=1311703) active AND [Spell 20919](https://www.wowhead.com/forever/spell=20919) active).
6. Cast [Spell 20919](https://www.wowhead.com/forever/spell=20919) when (Mana fraction ≥ 15% AND NOT [Echo of Righteousness](https://www.wowhead.com/forever/spell=1311704) active AND [Spell 20293](https://www.wowhead.com/forever/spell=20293) active).
7. Cast [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239).
8. Cast [Consecration](https://www.wowhead.com/forever/spell=26573) when Mana fraction ≥ 20%.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 309.41 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 197.95 |
| Auto-attack (tag 1) | 180.15 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 98.78 |
| Auto-attack (tag 3) | 89.09 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 69.90 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 69.52 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 53.37 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -12137.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11269.7 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -11011.0 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +4035.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3606.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3604.1 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3221.6 |
| Mana | OtherActionManaRegen (tag 1) | +3108.3 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -2329.2 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -2136.6 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -522.0 |
| Mana | OtherActionManaRegen (tag 2) | +16.5 |

## Hunter — Marksmanship

**Talents:** 5/35/11 · `5-0050552011523051-50005001`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 922.19 | 0.42 | 0.60 |
| Tauren | 912.12 | 0.41 | 0.28 |
| Troll | 915.72 | 0.41 | 0.76 |
| Windshaper | 916.63 | 0.41 | 0.63 |
| Human | 927.24 | 0.42 | 0.48 |
| Dwarf | 905.59 | 0.41 | 0.65 |
| Night Elf | 922.68 | 0.42 | 0.59 |
| High Order | 916.63 | 0.41 | 0.63 |

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
2. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=25295) when (NOT [Serpent Sting](https://www.wowhead.com/forever/spell=25295) DoT active AND Time remaining ≥ 12s).
3. Cast [Aimed Shot](https://www.wowhead.com/forever/spell=20904).
4. Cast [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) when Mana fraction ≥ 30%.
5. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287) when Mana fraction ≥ 15%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Shoot | 498.48 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 199.62 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 85.70 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 83.23 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 60.20 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -11436.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8194.1 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -6250.8 |
| Mana | OtherActionManaRegen (tag 1) | +6064.5 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3990.0 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -3822.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3599.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3597.5 |
| Mana | OtherActionManaRegen (tag 2) | +642.4 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

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

## Hunter — Pet/Melee

**Talents:** 16/10/25 · `53200005001-005005-5302002300502201`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Hawk damage uses an approximate guardian/pet model. Tracking talents affect its comparison with Survival on different creature types.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 862.85 | 0.43 | 0.00 |
| Windshaper | 865.33 | 0.42 | 0.00 |
| Orc | 875.95 | 0.44 | 0.00 |
| Human | 881.86 | 0.45 | 0.00 |
| Night Elf | 874.37 | 0.44 | 0.00 |
| Dwarf | 858.86 | 0.42 | 0.00 |
| Troll | 864.98 | 0.42 | 0.00 |
| High Order | 865.33 | 0.42 | 0.00 |

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
| Main hand | [Premier High Warlord's Blade](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Premier High Warlord's Quickblade](https://www.wowhead.com/forever/item=272684) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4s: [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447).
- -2s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.

### Rotation priorities

1. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
2. Use ready automatic cooldowns.
3. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
4. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
5. Cast [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) when (NOT [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) DoT active AND Time remaining ≥ 9s).
6. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when Time remaining ≥ 12s.
7. Cast [Wing Clip](https://www.wowhead.com/forever/spell=14268) when Mana fraction ≥ 40%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 163.76 |
| Auto-attack (tag 1) | 125.81 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 109.72 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 104.65 |
| Cat: Auto-attack (tag 1) | 96.14 |
| Auto-attack (tag 3) | 76.32 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 72.61 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 63.84 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8713.9 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -3494.7 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3205.3 |
| Mana | OtherActionManaRegen (tag 1) | +3053.6 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -1734.9 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -1322.3 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -977.8 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -828.4 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 853.33 | 0.46 | 0.00 |
| Tauren | 841.17 | 0.45 | 0.00 |
| Troll | 841.68 | 0.44 | 0.00 |
| Windshaper | 839.17 | 0.44 | 0.00 |
| Human | 858.14 | 0.46 | 0.00 |
| Dwarf | 836.97 | 0.45 | 0.00 |
| Night Elf | 850.48 | 0.45 | 0.00 |
| High Order | 839.17 | 0.44 | 0.00 |

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
| Main hand | [Premier High Warlord's Blade](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Premier High Warlord's Quickblade](https://www.wowhead.com/forever/item=272684) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4s: [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447).
- -2s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.

### Rotation priorities

1. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
2. Use ready automatic cooldowns.
3. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
4. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
5. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when Time remaining ≥ 12s.
6. Cast [Wing Clip](https://www.wowhead.com/forever/spell=14268) when Mana fraction ≥ 40%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 168.93 |
| Auto-attack (tag 1) | 129.95 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 110.65 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 103.11 |
| Cat: Auto-attack (tag 1) | 83.60 |
| Auto-attack (tag 3) | 79.87 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 65.43 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 53.68 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +16650.4 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -9571.5 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4350.7 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -3920.0 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3397.7 |
| Mana | OtherActionManaRegen (tag 1) | +3154.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2183.5 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -2038.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +503.4 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 844.26 | 0.38 | 0.00 |
| Troll | 839.19 | 0.38 | 0.00 |
| Undead | 848.82 | 0.39 | 0.00 |
| Human | 836.28 | 0.39 | 0.00 |
| Gnome | 847.73 | 0.38 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

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
| Ring 2 | [Advisor's Ring](https://www.wowhead.com/forever/item=19518) | 63 | — |
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
4. Cast [Life Tap](https://www.wowhead.com/forever/spell=11689) when Mana fraction ≤ 10%.
5. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
6. Cast [Searing Pain](https://www.wowhead.com/forever/spell=17923) when (NOT [Demonic Brand](https://www.wowhead.com/forever/spell=1293696) active OR [Demonic Brand](https://www.wowhead.com/forever/spell=1293696) stacks ≤ 0).
7. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND Time remaining ≥ 61s).
8. Cast [Bane of Agony](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Agony](https://www.wowhead.com/forever/spell=11713) DoT active AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND Time remaining ≥ 16s).
9. Cast [Corruption](https://www.wowhead.com/forever/spell=25311) when ([Corruption](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311}}}` AND Time remaining ≥ 12s).
10. Cast [Immolate](https://www.wowhead.com/forever/spell=25309) when ([Immolate](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309}}}` AND Time remaining ≥ 8s).
11. Cast [Soul Fire](https://www.wowhead.com/forever/spell=17924) when [Decimation](https://www.wowhead.com/forever/spell=440873) active.
12. Cast [Searing Pain](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 2s.
13. Cast [Shadow Bolt](https://www.wowhead.com/forever/spell=25307).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 309.54 |
| Succubus: Auto-attack (tag 1) | 91.76 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 90.43 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 84.19 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 82.20 |
| Succubus: [Demonic Brand](https://www.wowhead.com/forever/spell=1293697) | 62.49 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 49.46 |
| [Soul Fire](https://www.wowhead.com/forever/spell=17924) | 45.53 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -24417.3 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +13048.2 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +10307.3 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6576.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5285.2 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5150.0 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -4184.7 |
| Mana | OtherActionManaRegen (tag 1) | +3877.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3860.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3503.2 |
| Mana | [Soul Fire](https://www.wowhead.com/forever/spell=17924) | -3104.2 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |

## Priest — Shadow

**Talents:** 18/0/33 · `305030001303--504020501201302251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 801.83 | 0.23 | 0.00 |
| Undead | 811.89 | 0.23 | 0.00 |
| Human | 802.78 | 0.23 | 0.00 |
| Dwarf | 802.76 | 0.23 | 0.00 |
| Night Elf | 808.70 | 0.23 | 0.00 |
| Gnome | 809.05 | 0.23 | 0.00 |

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
4. Cast [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636).
5. `{"condition":{},"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10947}}}]}}` when `{}`.
6. Cast [Mind Blast (rank 9)](https://www.wowhead.com/forever/spell=10947).
7. Cast [Mind Flay (rank 6)](https://www.wowhead.com/forever/spell=18807).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 318.65 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 154.49 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 144.62 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 97.01 |
| [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | 84.62 |
| [Shadow Word: Death Backlash](https://www.wowhead.com/forever/spell=1309598) | 26.57 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 13.75 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +9117.0 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -6187.5 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -5938.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4100.6 |
| Mana | [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | -3200.5 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2848.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2402.3 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1597.8 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1514.0 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +50.7 |

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

## Warrior — Fury

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Level-60 rage, queued off-hand hit and Flurry charge timing still need in-game confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 786.32 | 0.47 | 0.00 |
| Tauren | 778.19 | 0.46 | 0.00 |
| Troll | 776.23 | 0.47 | 0.00 |
| Undead | 799.43 | 0.47 | 0.00 |
| Windshaper | 781.61 | 0.45 | 0.00 |
| Human | 781.33 | 0.48 | 0.00 |
| Dwarf | 779.66 | 0.46 | 0.00 |
| Night Elf | 771.25 | 0.47 | 0.00 |
| Gnome | 771.01 | 0.46 | 0.00 |
| High Order | 781.61 | 0.45 | 0.00 |

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
3. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active AND Rage ≤ 25 AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND NOT [Recklessness](https://www.wowhead.com/forever/spell=1719) active).
4. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
5. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when NOT [Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active.
6. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
7. Cast [Bloodthirst](https://www.wowhead.com/forever/spell=23894).
8. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
9. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 30.
10. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 30 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 141.94 |
| Auto-attack (tag 1) | 133.07 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 110.09 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 89.37 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 89.25 |
| Auto-attack (tag 3) | 58.03 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 48.14 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 40.56 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +1097.3 |
| Rage | Auto-attack (tag 1) | +905.7 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -876.4 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -612.9 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -495.0 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -439.1 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +179.6 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +178.6 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -147.5 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +100.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +98.8 |
| Rage | OtherActionRefund | +94.7 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 789.82 | 0.31 | 0.00 |
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

## Warrior — 2H Bloodthirst

**Talents:** 20/31/0 · `30304203032-050500320051310051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 740.72 | 0.48 | 0.00 |
| Tauren | 735.79 | 0.48 | 0.00 |
| Undead | 734.41 | 0.47 | 0.00 |
| Troll | 728.16 | 0.46 | 0.00 |
| Windshaper | 728.70 | 0.48 | 0.00 |
| Human | 736.62 | 0.47 | 0.00 |
| Dwarf | 730.95 | 0.48 | 0.00 |
| Gnome | 729.66 | 0.47 | 0.00 |
| Night Elf | 728.15 | 0.47 | 0.00 |
| High Order | 728.70 | 0.48 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

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
| Main hand | [Premier High Warlord's Battle Axe](https://www.wowhead.com/forever/item=272593) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Premier High Warlord's Recurve](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719).
3. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active AND Rage ≤ 25).
4. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
5. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when NOT [Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active.
6. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
7. Cast [Bloodthirst](https://www.wowhead.com/forever/spell=23894).
8. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
9. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 30.
10. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 50 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 228.14 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 101.02 |
| Auto-attack (tag 3) | 95.05 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 73.92 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 64.48 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 55.69 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 49.73 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 38.82 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1852.9 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -942.4 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -543.1 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -513.1 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +174.6 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -150.6 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +142.8 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -135.4 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +92.7 |
| Rage | OtherActionRefund | +77.9 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -54.8 |
| Rage | [Berserker Stance](https://www.wowhead.com/forever/spell=2458) | -26.9 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `420033152-053030031005102251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 738.45 | 0.57 | 0.05 |
| Tauren | 738.02 | 0.56 | 0.01 |
| Troll | 732.22 | 0.55 | 0.02 |
| Windshaper | 731.18 | 0.57 | 0.02 |
| Dwarf | 739.53 | 0.55 | 0.02 |

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
4. Cast [Searing Totem](https://www.wowhead.com/forever/spell=10438) when (Fire totem time remaining ≤ 0s AND Time remaining ≥ 5s).
5. Cast [Fire Nova](https://www.wowhead.com/forever/spell=408345) when [Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active.
6. Cast [Stormstrike](https://www.wowhead.com/forever/spell=17364).
7. Cast [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) when [Maelstrom Weapon](https://www.wowhead.com/forever/spell=408505) stacks ≥ 5.
8. Cast [Flame Shock](https://www.wowhead.com/forever/spell=29228) when (NOT [Flame Shock](https://www.wowhead.com/forever/spell=29228) DoT active AND Time remaining ≥ 10s).
9. Cast [Earth Shock](https://www.wowhead.com/forever/spell=10414).
10. Cast [Fire Nova](https://www.wowhead.com/forever/spell=408345) when Mana fraction ≥ 20%.

### Damage breakdown — Dwarf

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 208.88 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 139.13 |
| [Fire Nova](https://www.wowhead.com/forever/spell=408345) | 98.30 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 93.38 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 54.24 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 46.73 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 46.62 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 30.53 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Nova](https://www.wowhead.com/forever/spell=408345) | -14076.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8991.1 |
| Mana | OtherActionManaRegen (tag 1) | +6795.2 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4987.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4408.8 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -4089.4 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3617.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3593.0 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | [Grace of Air Totem](https://www.wowhead.com/forever/spell=10627) | -499.5 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -450.0 |
| Mana | OtherActionManaRegen (tag 2) | +376.9 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-550500000002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 735.56 | 0.49 | 0.00 |
| Tauren | 732.80 | 0.49 | 0.00 |
| Troll | 729.89 | 0.49 | 0.00 |
| Undead | 739.37 | 0.50 | 0.00 |
| Windshaper | 731.67 | 0.49 | 0.00 |
| Human | 738.28 | 0.49 | 0.00 |
| Dwarf | 728.26 | 0.48 | 0.00 |
| Night Elf | 733.88 | 0.49 | 0.00 |
| Gnome | 730.90 | 0.48 | 0.00 |
| High Order | 731.67 | 0.49 | 0.00 |

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
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Gloves - Minor Haste |
| Waist | [Might of the Timbermaw](https://www.wowhead.com/forever/item=19044) | 58 | — |
| Legs | [Titanic Leggings](https://www.wowhead.com/forever/item=22385) | 60 | Lesser Arcanum of Voracity |
| Feet | [Scalegut Treaders](https://www.wowhead.com/forever/item=275618) | 58 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
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

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 161.93 |
| Auto-attack (tag 3) | 114.25 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 111.00 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 110.71 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 72.95 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 43.40 |
| [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | 34.80 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 33.65 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1740.0 |
| Rage | [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | -1034.2 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -501.2 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -205.8 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -168.5 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +156.5 |
| Rage | [Rend](https://www.wowhead.com/forever/spell=11574) | -145.2 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -131.2 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +129.8 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.8 |
| Rage | OtherActionRefund | +96.5 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -95.1 |

## Mage — Arcane–Frost

**Talents:** 28/0/23 · `05020500310031053--05550002010003002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Uses an assumed Ice Lance coefficient and unresolved Fingers of Frost, Missile Barrage and Clearcasting timing.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 710.51 | 0.37 | 0.00 |
| Troll | 710.21 | 0.37 | 0.00 |
| Undead | 719.32 | 0.38 | 0.00 |
| Human | 706.52 | 0.37 | 0.00 |
| Gnome | 711.78 | 0.37 | 0.00 |
| High Order | 711.62 | 0.37 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

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
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction ≤ 15% AND Time remaining ≥ 25s).
4. Cast [Ice Lance](https://www.wowhead.com/forever/spell=30455) when [Fingers of Frost](https://www.wowhead.com/forever/spell=44543) active.
5. Cast [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) when [Missile Barrage](https://www.wowhead.com/forever/spell=44404) active.
6. Cast [Frostbolt](https://www.wowhead.com/forever/spell=25304).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 422.49 |
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 153.89 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 130.49 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 12.46 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -21627.1 |
| Mana | OtherActionManaRegen (tag 1) | +11817.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5581.5 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -2968.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2051.0 |
| Mana | OtherActionManaRegen (tag 2) | +1758.5 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.9 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +829.9 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +275.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Arcane

**Talents:** 33/3/15 · `053005023100311531-03-005500032`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 694.96 | 0.50 | 0.08 |
| Troll | 692.50 | 0.50 | 0.07 |
| Undead | 706.95 | 0.51 | 0.06 |
| Human | 691.52 | 0.50 | 0.06 |
| Gnome | 698.89 | 0.49 | 0.01 |
| High Order | 694.28 | 0.50 | 0.07 |

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
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 336.01 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 314.57 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 40.96 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 15.40 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -30563.7 |
| Mana | OtherActionManaRegen (tag 1) | +12013.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +6891.9 |
| Mana | OtherActionManaRegen (tag 2) | +4204.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4099.0 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -1628.2 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.9 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +851.1 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | -642.2 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +548.1 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Fire

**Talents:** 17/31/3 · `0501252000002-23450000130133051-003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 687.13 | 0.46 | 3.17 |
| Troll | 678.59 | 0.45 | 3.18 |
| Undead | 695.89 | 0.45 | 2.82 |
| Human | 686.04 | 0.44 | 2.25 |
| Gnome | 687.90 | 0.44 | 1.61 |
| High Order | 677.49 | 0.45 | 4.82 |

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
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 283.54 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 163.84 |
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 145.73 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 85.53 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 12.22 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 5.02 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -16729.7 |
| Mana | OtherActionManaRegen (tag 1) | +14633.4 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -14280.5 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -6604.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5571.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4111.5 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +3631.1 |
| Mana | OtherActionManaRegen (tag 2) | +1897.6 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.4 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.7 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +600.3 |
| Mana | [Fireball](https://www.wowhead.com/forever/spell=25306) | -371.0 |

## Mage — Frost

**Talents:** 11/3/37 · `050005001-03-0555003321001301251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 680.99 | 0.42 | 0.18 |
| Troll | 681.36 | 0.41 | 0.13 |
| Undead | 685.18 | 0.42 | 0.15 |
| Human | 677.07 | 0.41 | 0.11 |
| Gnome | 683.19 | 0.42 | 0.04 |
| High Order | 682.72 | 0.41 | 0.15 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 490.07 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 186.86 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 8.25 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -23541.3 |
| Mana | OtherActionManaRegen (tag 1) | +10825.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3722.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3707.9 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3582.8 |
| Mana | OtherActionManaRegen (tag 2) | +2358.9 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.6 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +848.7 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +599.7 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Druid — Feral

**Talents:** 9/34/8 · `050022-3521002023032213041-053`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 677.88 | 0.21 | 0.00 |
| Windshaper | 679.94 | 0.21 | 0.00 |
| Night Elf | 679.07 | 0.21 | 0.00 |
| High Order | 679.94 | 0.21 | 0.00 |

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

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Use ready automatic cooldowns.
3. Cast [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) when Energy ≤ 40.
4. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Time remaining ≤ 4s AND Combo points ≥ 3).
5. Cast [Rip](https://www.wowhead.com/forever/spell=9896) when (Combo points ≥ 4 AND NOT [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND Time remaining ≥ 8s).
6. Cast [Shred](https://www.wowhead.com/forever/spell=9830).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Shred](https://www.wowhead.com/forever/spell=9830) | 265.45 |
| Auto-attack (tag 1) | 226.80 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 180.35 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 4.70 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.63 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -3450.0 |
| Energy | OtherActionEnergyRegen | +3047.5 |
| Energy | [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) | +690.0 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -588.7 |
| Energy | OtherActionRefund | +179.1 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.0 |
| ComboPoints | [Rip](https://www.wowhead.com/forever/spell=9896) | -98.3 |
| ComboPoints | [Shred](https://www.wowhead.com/forever/spell=9830) | +66.3 |
| ComboPoints | [Primal Fury](https://www.wowhead.com/forever/spell=37117) | +37.4 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -35.2 |
| ComboPoints | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -4.2 |
| Mana | OtherActionManaRegen (tag 2) | +0.0 |

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
| Orc | 556.75 | 0.22 | 0.00 |
| Troll | 558.26 | 0.22 | 0.00 |
| Undead | 580.14 | 0.23 | 0.00 |
| Windshaper | 556.78 | 0.22 | 0.00 |
| Human | 562.52 | 0.22 | 0.00 |
| Dwarf | 550.79 | 0.22 | 0.00 |
| Night Elf | 559.84 | 0.22 | 0.00 |
| Gnome | 559.96 | 0.22 | 0.00 |
| High Order | 556.78 | 0.22 | 0.00 |

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

1. Cast [Slice and Dice](https://www.wowhead.com/forever/spell=6774) when (Combo points ≥ 1 AND NOT [Slice and Dice](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6s).
2. Cast [Vanish](https://www.wowhead.com/forever/spell=1856) when (NOT [Stealth](https://www.wowhead.com/forever/spell=1787) active AND Energy ≥ 60 AND Combo points ≤ 1 AND [Slice and Dice](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 10s).
3. Cast [Premeditation](https://www.wowhead.com/forever/spell=14183) when Combo points ≤ 2.
4. Cast [Ambush](https://www.wowhead.com/forever/spell=11269) when Combo points ≤ 3.
5. Cast [Slice and Dice](https://www.wowhead.com/forever/spell=6774) when (Combo points ≥ 5 AND `{"auraRemainingTime":{"auraId":{"spellId":6774}}}` ≤ 2s AND Time remaining ≥ 6s).
6. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
7. Use ready automatic cooldowns.
8. Cast [Hemorrhage](https://www.wowhead.com/forever/spell=16511) when (NOT [Hemorrhage](https://www.wowhead.com/forever/spell=16511) active AND Combo points ≥ 2 AND NOT [Rupture](https://www.wowhead.com/forever/spell=11275) DoT active AND Time remaining ≥ 10s).
9. Cast [Rupture](https://www.wowhead.com/forever/spell=11275) when (Combo points ≥ 3 AND NOT [Rupture](https://www.wowhead.com/forever/spell=11275) DoT active AND Time remaining ≥ 10s).
10. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 5.
11. Cast [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) when Combo points ≤ 3.
12. Cast [Hemorrhage](https://www.wowhead.com/forever/spell=16511).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 166.32 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 132.24 |
| Auto-attack (tag 2) | 84.00 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 39.48 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 36.48 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 28.11 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 24.93 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 18.82 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +2984.7 |
| Energy | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | -2827.6 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +538.6 |
| Energy | [Rupture](https://www.wowhead.com/forever/spell=11275) | -482.9 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -345.3 |
| Energy | OtherActionRefund | +151.7 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +98.9 |
| Energy | [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) | -84.6 |
| ComboPoints | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | +84.6 |
| Energy | [Ambush](https://www.wowhead.com/forever/spell=11269) | -82.1 |
| ComboPoints | [Rupture](https://www.wowhead.com/forever/spell=11275) | -62.9 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -39.5 |

## Priest — Smite

**Talents:** 31/17/3 · `515030031305001031-00505023002-003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** The rank-2 Smite fallback depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 562.50 | 0.15 | 0.00 |
| Undead | 566.39 | 0.15 | 0.00 |
| Human | 559.97 | 0.14 | 0.00 |
| Dwarf | 559.48 | 0.14 | 0.00 |
| Night Elf | 562.23 | 0.14 | 0.00 |
| Gnome | 564.01 | 0.15 | 0.00 |

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
6. Cast [Smite](https://www.wowhead.com/forever/spell=591).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Smite](https://www.wowhead.com/forever/spell=10934) | 201.59 |
| [Penance](https://www.wowhead.com/forever/spell=1316995) | 140.00 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 128.58 |
| [Smite](https://www.wowhead.com/forever/spell=591) | 87.55 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 8.67 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -13533.9 |
| Mana | OtherActionManaRegen (tag 1) | +9163.4 |
| Mana | [Penance](https://www.wowhead.com/forever/spell=1316995) | -8289.8 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -6276.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4264.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3612.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.4 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1551.0 |
| Mana | [Smite](https://www.wowhead.com/forever/spell=591) | -1001.6 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |

## Shaman — Stormcaller

**Talents:** 28/23/0 · `550032150010303-055030031004002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** The rank-2 Lightning Bolt filler depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 539.75 | 0.22 | 0.30 |
| Tauren | 538.73 | 0.22 | 0.29 |
| Troll | 535.56 | 0.22 | 0.26 |
| Windshaper | 535.61 | 0.22 | 0.26 |
| Dwarf | 535.32 | 0.22 | 0.28 |

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
7. Cast [Lightning Bolt](https://www.wowhead.com/forever/spell=529).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 287.73 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 86.78 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 82.96 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 33.22 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 28.93 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 14.30 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 4.38 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.44 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -17420.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5290.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5278.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3599.4 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3244.8 |
| Mana | OtherActionManaRegen (tag 1) | +3082.1 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | -1149.1 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | OtherActionManaRegen (tag 2) | +137.9 |

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
