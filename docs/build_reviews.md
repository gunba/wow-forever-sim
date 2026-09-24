# Build reviews

These summaries describe the published loadouts. They are simulation results, not independent confirmation of server mechanics or proof of a global optimum.

The tables and [matrix](../artifacts/forever_dps_5min.png) use the same 184 common-seed replays. Equipment selections came from an earlier mechanics revision; these results use the corrected engine. Historical search gains are not directly comparable to this release.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and paid shared-hit normalization. [Scenario and exchange model](../tools/forever_bench/README.md) · [In-game checks](in_game_checks.md)

## Paladin — Retribution

**Talents:** 13/7/31 · `253003-232-052052310012330301`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Separate seal Echoes can coexist and fire on a landed white swing; verify their simultaneous behavior in game (T53). Lower-rank seals and Consecration also need confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1137.52 | 0.61 | 0.06 |
| Human | 1133.40 | 0.61 | 0.22 |
| Dwarf | 1143.55 | 0.60 | 0.07 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Dwarf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Beads of Ogre Mojo](https://www.wowhead.com/forever/item=22149) | 63 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Spitfire Bracers](https://www.wowhead.com/forever/item=20481) | 62 | Enchant Bracer - Superior Strength |
| Hands | [Raider Handwraps](https://www.wowhead.com/forever/item=272097) | 65 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Lesser Arcanum of Voracity |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Frozen Heart of the Mountain](https://www.wowhead.com/forever/item=249469) | 55 | — |
| Main hand | [Premier High Warlord's Destroyer](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Libram of Invocation](https://www.wowhead.com/forever/item=249442) | 45 | — |

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

### Damage breakdown — Dwarf

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 311.57 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 200.53 |
| Auto-attack (tag 1) | 180.78 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 100.18 |
| Auto-attack (tag 3) | 88.49 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 70.60 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 68.83 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 64.07 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -11511.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11175.0 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -10439.2 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3782.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3592.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3592.2 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3189.3 |
| Mana | OtherActionManaRegen (tag 1) | +3109.2 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -2749.5 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -2327.0 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -521.8 |
| Mana | OtherActionManaRegen (tag 2) | +12.5 |

## Hunter — Pet/Melee

**Talents:** 16/10/25 · `53200005001-005005-5302002300502201`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Hawk damage uses an approximate guardian/pet model. Tracking talents affect its comparison with Survival on different creature types.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 927.39 | 0.43 | 0.00 |
| Windshaper | 930.37 | 0.43 | 0.00 |
| Orc | 944.57 | 0.44 | 0.00 |
| Human | 951.48 | 0.44 | 0.00 |
| Night Elf | 943.15 | 0.44 | 0.00 |
| Dwarf | 922.87 | 0.43 | 0.00 |
| Troll | 930.70 | 0.44 | 0.00 |
| High Order | 930.37 | 0.43 | 0.00 |

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
| Auto-attack (tag 2) | 163.81 |
| Cat: Auto-attack (tag 1) | 147.16 |
| Auto-attack (tag 1) | 125.77 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 109.72 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 104.66 |
| Auto-attack (tag 3) | 77.27 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 69.51 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 63.66 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8253.5 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -3548.8 |
| Mana | OtherActionManaRegen (tag 1) | +3342.4 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3076.2 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -1730.5 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -1322.3 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -884.5 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -826.9 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -100.0 |

## Hunter — Beast Mastery

**Talents:** 31/20/0 · `5320001505101251-00531510005`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 944.58 | 0.27 | 0.01 |
| Tauren | 926.79 | 0.26 | 0.02 |
| Troll | 930.07 | 0.27 | 0.03 |
| Windshaper | 950.03 | 0.30 | 0.00 |
| Human | 948.60 | 0.27 | 0.01 |
| Dwarf | 923.81 | 0.26 | 0.02 |
| Night Elf | 937.69 | 0.26 | 0.02 |
| High Order | 950.03 | 0.30 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Shifting Cloak](https://www.wowhead.com/forever/item=18511) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Boots - Minor Agility |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Gloves - Minor Haste |
| Waist | [Molten Belt](https://www.wowhead.com/forever/item=19163) | 70 | — |
| Legs | [Outrider's Chain Leggings](https://www.wowhead.com/forever/item=22673) | 65 | Lesser Arcanum of Voracity |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Shoot | 353.17 |
| Cat: Auto-attack (tag 1) | 223.20 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | 124.52 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 81.68 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 65.84 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 51.84 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 47.93 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 1.85 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | -9096.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7402.2 |
| Mana | OtherActionManaRegen (tag 1) | +6004.9 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5476.9 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -4614.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3601.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3598.5 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3120.1 |
| Mana | OtherActionManaRegen (tag 2) | +784.6 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Intimidation](https://www.wowhead.com/forever/spell=19577) | -378.9 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Paladin — Physical Ret

**Talents:** 13/7/31 · `250003003-232-052250310012330301`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Strength/AP equipment is unchanged. Champion of the Light converts existing Intellect into spell power without equipping caster gear; the separate seal Echoes still require the T53 in-game check.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 933.52 | 0.57 | 0.02 |
| Human | 943.98 | 0.56 | 0.00 |
| Dwarf | 939.69 | 0.57 | 0.00 |

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
| Ranged/relic | [Libram of Invocation](https://www.wowhead.com/forever/item=249442) | 45 | — |

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

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 308.09 |
| Auto-attack (tag 1) | 226.82 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 113.08 |
| Auto-attack (tag 3) | 110.31 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 68.39 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 34.56 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 33.31 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 24.70 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -11824.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11211.2 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -10719.6 |
| Mana | OtherActionManaRegen (tag 1) | +5182.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3589.6 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3553.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3273.5 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3006.6 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -1618.3 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1285.0 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -522.0 |
| Mana | OtherActionManaRegen (tag 2) | +6.5 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 915.16 | 0.46 | 0.00 |
| Tauren | 901.42 | 0.45 | 0.00 |
| Troll | 901.18 | 0.44 | 0.00 |
| Windshaper | 896.90 | 0.44 | 0.00 |
| Human | 920.62 | 0.46 | 0.00 |
| Dwarf | 896.77 | 0.45 | 0.00 |
| Night Elf | 912.11 | 0.45 | 0.00 |
| High Order | 896.90 | 0.44 | 0.00 |

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
| Auto-attack (tag 2) | 168.68 |
| Cat: Auto-attack (tag 1) | 130.58 |
| Auto-attack (tag 1) | 129.67 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 110.15 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 103.65 |
| Auto-attack (tag 3) | 80.75 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 65.77 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 49.99 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +16654.7 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -9640.1 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4336.7 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -3663.9 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3397.7 |
| Mana | OtherActionManaRegen (tag 1) | +3157.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2118.5 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -2050.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +382.7 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 848.09 | 0.38 | 0.00 |
| Troll | 845.49 | 0.37 | 0.00 |
| Undead | 853.57 | 0.38 | 0.00 |
| Human | 840.85 | 0.37 | 0.00 |
| Gnome | 846.60 | 0.37 | 0.00 |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 308.20 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 90.25 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 84.31 |
| Succubus: Auto-attack (tag 1) | 83.78 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 82.01 |
| Succubus: [Demonic Brand](https://www.wowhead.com/forever/spell=1293697) | 69.84 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 50.29 |
| [Soul Fire](https://www.wowhead.com/forever/spell=17924) | 46.16 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -24309.3 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +13007.5 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +10307.0 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6553.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5284.7 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5147.7 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -4251.5 |
| Mana | OtherActionManaRegen (tag 1) | +3876.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3868.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3503.8 |
| Mana | [Soul Fire](https://www.wowhead.com/forever/spell=17924) | -3146.6 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |

## Priest — Shadow

**Talents:** 18/0/33 · `305030001303--504020501201302251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 804.28 | 0.23 | 0.00 |
| Undead | 813.94 | 0.23 | 0.00 |
| Human | 803.94 | 0.23 | 0.00 |
| Dwarf | 803.96 | 0.23 | 0.00 |
| Night Elf | 810.06 | 0.23 | 0.00 |
| Gnome | 805.54 | 0.23 | 0.00 |

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
2. Cast [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636).
3. Cast [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) when NOT [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) DoT active.
4. Cast [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) when (NOT [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) DoT active AND Time remaining ≥ 10).
5. `{"condition":{},"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10947}}}]}}` when `{}`.
6. Cast [Mind Blast (rank 9)](https://www.wowhead.com/forever/spell=10947).
7. `{"channelSpell":{"spellId":{"spellId":18807,"rank":6},"interruptIf":{"and":{"vals":[{"cmp":{"op":"OpGe","lhs":{"spellChanneledTicks":{"spellId":{"spellId":18807}}},"rhs":{"const":{"val":"2"}}}},{"or":{"vals":[{"spellCanCast":{"spellId":{"spellId":10947}}},{"spellCanCast":{"spellId":{"spellId":1309636}}},{"and":{"vals":[{"not":{"val":{"dotIsActive":{"spellId":{"spellId":19280}}}}},{"spellCanCast":{"spellId":{"spellId":19280}}}]}},{"and":{"vals":[{"not":{"val":{"dotIsActive":{"spellId":{"spellId":10894}}}}},{"spellCanCast":{"spellId":{"spellId":10894}}}]}}]}}]}}}}`.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 316.55 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 154.54 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 143.87 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 96.07 |
| [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | 89.69 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 13.21 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +9116.6 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -7062.8 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -5945.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4366.5 |
| Mana | [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | -3397.3 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2848.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2435.3 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1599.6 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1565.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +802.8 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |

## Hunter — Marksmanship

**Talents:** 5/35/11 · `5-0050552011523051-50005001`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 806.13 | 0.38 | 2.03 |
| Tauren | 796.96 | 0.37 | 1.22 |
| Troll | 797.61 | 0.37 | 2.03 |
| Windshaper | 813.13 | 0.39 | 1.31 |
| Human | 810.77 | 0.38 | 1.67 |
| Dwarf | 791.53 | 0.37 | 2.16 |
| Night Elf | 806.94 | 0.38 | 2.06 |
| High Order | 813.13 | 0.39 | 1.31 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Amulet of the Darkmoon](https://www.wowhead.com/forever/item=19491) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Shifting Cloak](https://www.wowhead.com/forever/item=18511) | 62 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Primal Batskin Bracers](https://www.wowhead.com/forever/item=19687) | 65 | Enchant Boots - Minor Agility |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Gloves - Minor Haste |
| Waist | [Molten Belt](https://www.wowhead.com/forever/item=19163) | 70 | — |
| Legs | [Outrider's Chain Leggings](https://www.wowhead.com/forever/item=22673) | 65 | Lesser Arcanum of Voracity |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Shoot | 408.75 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 191.77 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 83.83 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 75.70 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 53.08 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -11340.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7308.4 |
| Mana | OtherActionManaRegen (tag 1) | +5873.5 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5808.7 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3964.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3612.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.8 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -3463.5 |
| Mana | OtherActionManaRegen (tag 2) | +677.1 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Warrior — Fury

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Level-60 rage, queued off-hand hit and Flurry charge timing still need in-game confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 794.31 | 0.47 | 0.00 |
| Tauren | 781.01 | 0.46 | 0.00 |
| Troll | 775.87 | 0.46 | 0.00 |
| Undead | 802.29 | 0.46 | 0.00 |
| Windshaper | 783.15 | 0.44 | 0.00 |
| Human | 781.94 | 0.47 | 0.00 |
| Dwarf | 781.45 | 0.46 | 0.00 |
| Night Elf | 771.72 | 0.47 | 0.00 |
| Gnome | 780.95 | 0.45 | 0.00 |
| High Order | 783.15 | 0.44 | 0.00 |

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
| Auto-attack (tag 2) | 142.05 |
| Auto-attack (tag 1) | 133.06 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 110.44 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 89.32 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 89.25 |
| Auto-attack (tag 3) | 58.05 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 47.88 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 40.61 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +1097.8 |
| Rage | Auto-attack (tag 1) | +905.8 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -875.7 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -613.3 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -495.2 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -440.9 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +179.5 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +178.9 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -147.5 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +100.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +98.9 |
| Rage | OtherActionRefund | +95.4 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2525000013520105--0550015103`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 800.53 | 0.46 | 0.01 |
| Troll | 796.69 | 0.45 | 0.01 |
| Undead | 802.23 | 0.45 | 0.01 |
| Human | 794.26 | 0.45 | 0.01 |
| Gnome | 800.22 | 0.45 | 0.00 |

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

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 367.83 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 103.59 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 93.28 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 83.87 |
| Succubus: Auto-attack (tag 1) | 74.24 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 39.32 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 18.10 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 9.45 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -30467.0 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +26652.7 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6608.0 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5230.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5158.5 |
| Mana | OtherActionManaRegen (tag 1) | +3845.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.8 |
| Mana | [Siphon Life](https://www.wowhead.com/forever/spell=18881) | -3574.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3390.2 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -459.3 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -327.7 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 785.60 | 0.31 | 0.00 |
| Troll | 780.66 | 0.30 | 0.00 |
| Undead | 790.34 | 0.31 | 0.00 |
| Human | 779.18 | 0.30 | 0.00 |
| Gnome | 783.40 | 0.31 | 0.00 |

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
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 257.97 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 113.68 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 90.05 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 83.23 |
| Succubus: Auto-attack (tag 1) | 80.45 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 77.60 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 49.30 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 14.32 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24347.7 |
| Mana | [Incinerate](https://www.wowhead.com/forever/spell=1293813) | -21125.3 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6701.1 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -6282.3 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -5656.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5228.0 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5159.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4525.6 |
| Mana | OtherActionManaRegen (tag 1) | +3878.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3606.2 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -459.1 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `252200001351-0025003001-0550005103`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 764.31 | 0.49 | 0.00 |
| Troll | 760.51 | 0.48 | 0.00 |
| Undead | 768.49 | 0.49 | 0.00 |
| Human | 757.39 | 0.48 | 0.00 |
| Gnome | 762.79 | 0.48 | 0.00 |

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

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 434.84 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 114.96 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 99.25 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 81.34 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 20.53 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 12.11 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 3.40 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.07 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -32561.5 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24680.5 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6833.8 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5497.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5334.3 |
| Mana | OtherActionManaRegen (tag 1) | +3882.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3630.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3595.5 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -622.9 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -346.4 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -189.9 |

## Warrior — Fury (Sunder)

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 757.71 | 0.47 | 0.00 |
| Tauren | 744.97 | 0.47 | 0.00 |
| Troll | 740.14 | 0.48 | 0.00 |
| Undead | 766.98 | 0.48 | 0.00 |
| Windshaper | 746.68 | 0.46 | 0.00 |
| Human | 745.68 | 0.50 | 0.00 |
| Dwarf | 746.28 | 0.46 | 0.00 |
| Night Elf | 736.34 | 0.49 | 0.00 |
| Gnome | 747.16 | 0.46 | 0.00 |
| High Order | 746.68 | 0.46 | 0.00 |

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

1. Cast [Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) when ([Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) stacks < 5 OR `{"auraRemainingTime":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597,"rank":5}}}` < 8s).
2. `{"condition":{"and":{"vals":[{"or":{"vals":[{"cmp":{"op":"OpLt","lhs":{"auraNumStacks":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597,"rank":5}}},"rhs":{"const":{"val":"5"}}}},{"cmp":{"op":"OpLt","lhs":{"auraRemainingTime":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597,"rank":5}}},"rhs":{"const":{"val":"8s"}}}}]}},{"cmp":{"op":"OpLt","lhs":{"currentRage":{}},"rhs":{"const":{"val":"15"}}}}]}},"wait":{"duration":{"const":{"val":"500ms"}}}}` when (([Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) stacks < 5 OR `{"auraRemainingTime":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597,"rank":5}}}` < 8s) AND Rage < 15).
3. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719) when Time remaining ≤ 300s.
4. Use ready automatic cooldowns.
5. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active AND Rage ≤ 25 AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND NOT [Recklessness](https://www.wowhead.com/forever/spell=1719) active).
6. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
7. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when NOT [Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active.
8. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
9. Cast [Bloodthirst](https://www.wowhead.com/forever/spell=23894).
10. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
11. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 30.
12. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 30 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 140.95 |
| Auto-attack (tag 1) | 139.33 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 90.72 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 83.13 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 81.75 |
| Auto-attack (tag 3) | 60.43 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 46.67 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 37.61 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +1093.1 |
| Rage | Auto-attack (tag 1) | +943.9 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -792.6 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -576.1 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -455.4 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -376.2 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -273.2 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +180.0 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +177.9 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -125.1 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +100.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +100.0 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `420033152-053030031005102251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 738.69 | 0.57 | 0.06 |
| Tauren | 737.10 | 0.55 | 0.01 |
| Troll | 731.06 | 0.56 | 0.02 |
| Windshaper | 736.99 | 0.56 | 0.03 |
| Dwarf | 743.94 | 0.55 | 0.02 |

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
| Ranged/relic | [Burning Totem](https://www.wowhead.com/forever/item=272433) | 65 | — |

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
| Auto-attack (tag 1) | 209.14 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 138.66 |
| [Fire Nova](https://www.wowhead.com/forever/spell=408345) | 98.13 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 93.47 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 55.07 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 51.40 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 45.77 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 30.54 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Nova](https://www.wowhead.com/forever/spell=408345) | -14092.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8990.1 |
| Mana | OtherActionManaRegen (tag 1) | +6819.5 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4977.5 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -4490.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4402.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3597.7 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3220.6 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | [Grace of Air Totem](https://www.wowhead.com/forever/spell=10627) | -499.4 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -450.0 |
| Mana | OtherActionManaRegen (tag 2) | +338.3 |

## Warrior — 2H Bloodthirst

**Talents:** 20/31/0 · `30304203032-050500320051310051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 743.14 | 0.49 | 0.00 |
| Tauren | 731.93 | 0.49 | 0.00 |
| Undead | 736.91 | 0.48 | 0.00 |
| Troll | 732.84 | 0.47 | 0.00 |
| Windshaper | 726.21 | 0.48 | 0.00 |
| Human | 738.13 | 0.48 | 0.00 |
| Dwarf | 733.90 | 0.49 | 0.00 |
| Gnome | 732.78 | 0.48 | 0.00 |
| Night Elf | 730.35 | 0.48 | 0.00 |
| High Order | 726.21 | 0.48 | 0.00 |

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
| Hands | [Raider Gauntlets](https://www.wowhead.com/forever/item=272095) | 65 | Enchant Weapon - Agility |
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
| Auto-attack (tag 1) | 227.87 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 104.93 |
| Auto-attack (tag 3) | 94.61 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 73.37 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 62.96 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 51.12 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 50.01 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 38.28 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1847.3 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -978.2 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -538.0 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -497.5 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +174.7 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +141.0 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -138.5 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -123.8 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +92.8 |
| Rage | OtherActionRefund | +77.7 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -54.9 |
| Rage | [Berserker Stance](https://www.wowhead.com/forever/spell=2458) | -28.1 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-550500000002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 734.10 | 0.49 | 0.00 |
| Tauren | 731.14 | 0.49 | 0.00 |
| Troll | 729.01 | 0.47 | 0.00 |
| Undead | 737.62 | 0.50 | 0.00 |
| Windshaper | 730.71 | 0.50 | 0.00 |
| Human | 737.00 | 0.49 | 0.00 |
| Dwarf | 726.69 | 0.49 | 0.00 |
| Night Elf | 732.54 | 0.49 | 0.00 |
| Gnome | 730.85 | 0.48 | 0.00 |
| High Order | 730.71 | 0.50 | 0.00 |

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
| Auto-attack (tag 1) | 161.14 |
| Auto-attack (tag 3) | 114.67 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 111.00 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 110.60 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 73.06 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 43.31 |
| [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | 34.78 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 32.22 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1740.3 |
| Rage | [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | -1034.6 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -500.7 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -205.8 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -168.3 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +156.8 |
| Rage | [Rend](https://www.wowhead.com/forever/spell=11574) | -145.4 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -131.0 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +130.0 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.8 |
| Rage | OtherActionRefund | +96.5 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -95.8 |

## Mage — Arcane–Frost

**Talents:** 28/0/23 · `05020500310031053--05550002010003002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Uses an assumed Ice Lance coefficient and unresolved Fingers of Frost, Missile Barrage and Clearcasting timing.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 711.32 | 0.37 | 0.00 |
| Troll | 711.28 | 0.38 | 0.00 |
| Undead | 719.11 | 0.38 | 0.00 |
| Human | 706.75 | 0.37 | 0.00 |
| Gnome | 712.02 | 0.38 | 0.00 |
| High Order | 712.98 | 0.37 | 0.00 |

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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 422.42 |
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 154.48 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 129.70 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 12.51 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -21637.7 |
| Mana | OtherActionManaRegen (tag 1) | +11816.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5602.5 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -2937.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2025.7 |
| Mana | OtherActionManaRegen (tag 2) | +1753.9 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.8 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +827.4 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +274.2 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Arcane

**Talents:** 33/3/15 · `053005023100311531-03-005500032`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 695.59 | 0.51 | 0.06 |
| Troll | 693.13 | 0.50 | 0.06 |
| Undead | 706.03 | 0.51 | 0.06 |
| Human | 691.83 | 0.49 | 0.04 |
| Gnome | 699.37 | 0.49 | 0.02 |
| High Order | 694.77 | 0.50 | 0.08 |

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
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 334.71 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 314.90 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 41.02 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 15.40 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -30631.4 |
| Mana | OtherActionManaRegen (tag 1) | +12045.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +6878.8 |
| Mana | OtherActionManaRegen (tag 2) | +4234.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4121.5 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -1631.4 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.8 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.4 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | -637.1 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +549.6 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Fire

**Talents:** 17/31/3 · `0501252000002-23450000130133051-003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 686.68 | 0.45 | 3.11 |
| Troll | 683.22 | 0.44 | 2.67 |
| Undead | 696.16 | 0.44 | 2.83 |
| Human | 685.70 | 0.44 | 2.13 |
| Gnome | 687.89 | 0.44 | 1.63 |
| High Order | 681.58 | 0.41 | 1.01 |

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
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 283.40 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 163.96 |
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 146.11 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 85.53 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 12.23 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 4.93 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -16737.2 |
| Mana | OtherActionManaRegen (tag 1) | +14653.7 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -14276.3 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -6607.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5572.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4106.8 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +3632.9 |
| Mana | OtherActionManaRegen (tag 2) | +1874.8 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.3 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +851.0 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +600.3 |
| Mana | [Fireball](https://www.wowhead.com/forever/spell=25306) | -359.2 |

## Mage — Frost

**Talents:** 11/3/37 · `050005001-03-0555003321001301251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 680.67 | 0.41 | 0.20 |
| Troll | 681.66 | 0.41 | 0.15 |
| Undead | 685.35 | 0.41 | 0.14 |
| Human | 676.87 | 0.41 | 0.12 |
| Gnome | 683.31 | 0.42 | 0.04 |
| High Order | 683.79 | 0.41 | 0.16 |

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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 489.75 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 187.31 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 8.30 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -23507.4 |
| Mana | OtherActionManaRegen (tag 1) | +10826.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3715.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3705.6 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3601.4 |
| Mana | OtherActionManaRegen (tag 2) | +2359.8 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.0 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.4 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +599.7 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Druid — Feral

**Talents:** 9/34/8 · `050022-3521002023032213041-053`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 680.07 | 0.22 | 0.00 |
| Windshaper | 681.90 | 0.22 | 0.00 |
| Night Elf | 681.66 | 0.22 | 0.00 |
| High Order | 681.90 | 0.22 | 0.00 |

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
5. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Combo points ≥ 5 AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 5s).
6. Cast [Rip](https://www.wowhead.com/forever/spell=9896) when (Combo points ≥ 4 AND NOT [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND Time remaining ≥ 8s).
7. Cast [Shred](https://www.wowhead.com/forever/spell=9830).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Shred](https://www.wowhead.com/forever/spell=9830) | 256.10 |
| Auto-attack (tag 1) | 226.77 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 172.81 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 23.61 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.62 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -3328.0 |
| Energy | OtherActionEnergyRegen | +3047.3 |
| Energy | [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) | +679.1 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -571.3 |
| Energy | OtherActionRefund | +172.4 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -165.8 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.5 |
| ComboPoints | [Rip](https://www.wowhead.com/forever/spell=9896) | -94.2 |
| ComboPoints | [Shred](https://www.wowhead.com/forever/spell=9830) | +71.9 |
| ComboPoints | [Primal Fury](https://www.wowhead.com/forever/spell=37117) | +41.9 |
| ComboPoints | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -18.5 |
| Mana | OtherActionManaRegen (tag 2) | +0.0 |

## Rogue — Mutilate

**Talents:** 31/20/0 · `0053031035140105-302303202014`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 618.11 | 0.29 | 0.00 |
| Troll | 613.17 | 0.28 | 0.00 |
| Undead | 639.72 | 0.30 | 0.00 |
| Windshaper | 620.42 | 0.29 | 0.00 |
| Human | 614.36 | 0.29 | 0.00 |
| Dwarf | 613.09 | 0.28 | 0.00 |
| Night Elf | 623.38 | 0.28 | 0.00 |
| Gnome | 627.02 | 0.29 | 0.00 |
| High Order | 620.42 | 0.29 | 0.00 |

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
| Auto-attack (tag 1) | 159.75 |
| Auto-attack (tag 2) | 96.43 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 72.23 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 70.82 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 59.61 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 44.67 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 43.40 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 28.60 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3013.8 |
| Energy | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | -2889.5 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -796.3 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +762.0 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -410.0 |
| Energy | OtherActionRefund | +148.7 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -130.6 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.6 |
| ComboPoints | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | +90.1 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +32.3 |
| ComboPoints | [Seal Fate](https://www.wowhead.com/forever/spell=14195) | +31.9 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -22.0 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 606.94 | 0.27 | 0.00 |
| Troll | 604.14 | 0.27 | 0.00 |
| Undead | 624.96 | 0.28 | 0.00 |
| Windshaper | 609.52 | 0.27 | 0.00 |
| Human | 610.43 | 0.28 | 0.00 |
| Dwarf | 600.84 | 0.27 | 0.00 |
| Night Elf | 608.86 | 0.27 | 0.00 |
| Gnome | 608.29 | 0.28 | 0.00 |
| High Order | 609.52 | 0.27 | 0.00 |

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
| Auto-attack (tag 1) | 178.40 |
| [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | 136.73 |
| Auto-attack (tag 2) | 110.86 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 88.91 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 33.85 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 25.41 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 16.34 |
| Auto-attack (tag 3) | 12.39 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3175.9 |
| Energy | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | -2912.4 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -469.4 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -462.7 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +460.4 |
| Energy | OtherActionRefund | +105.9 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +96.5 |
| Energy | [Blade Flurry](https://www.wowhead.com/forever/spell=13877) | -75.0 |
| ComboPoints | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | +69.5 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -66.3 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -26.0 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +24.5 |

## Druid — Balance

**Talents:** 38/0/13 · `5232220115501351--505003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 605.67 | 0.26 | 0.00 |
| Windshaper | 605.46 | 0.26 | 0.00 |
| Night Elf | 606.24 | 0.26 | 0.00 |
| High Order | 605.46 | 0.26 | 0.00 |

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
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 354.47 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 91.46 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 85.95 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 74.36 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=25298) | -15143.7 |
| Mana | OtherActionManaRegen (tag 1) | +7735.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4499.7 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -4218.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3599.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3592.8 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1553.4 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1205.2 |
| Mana | [Moonkin Form](https://www.wowhead.com/forever/spell=24858) | -435.4 |
| Mana | OtherActionManaRegen (tag 2) | +77.7 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | +2.0 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | -0.0 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 556.67 | 0.21 | 0.00 |
| Troll | 556.79 | 0.22 | 0.00 |
| Undead | 579.82 | 0.23 | 0.00 |
| Windshaper | 556.79 | 0.22 | 0.00 |
| Human | 562.85 | 0.22 | 0.00 |
| Dwarf | 551.13 | 0.21 | 0.00 |
| Night Elf | 560.24 | 0.22 | 0.00 |
| Gnome | 559.23 | 0.22 | 0.00 |
| High Order | 556.79 | 0.22 | 0.00 |

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
| Auto-attack (tag 1) | 166.39 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 132.19 |
| Auto-attack (tag 2) | 83.94 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 39.41 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 36.34 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 27.97 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 24.88 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 18.98 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +2984.6 |
| Energy | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | -2827.6 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +538.6 |
| Energy | [Rupture](https://www.wowhead.com/forever/spell=11275) | -483.0 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -345.6 |
| Energy | OtherActionRefund | +151.9 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +98.9 |
| ComboPoints | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | +84.6 |
| Energy | [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) | -84.1 |
| Energy | [Ambush](https://www.wowhead.com/forever/spell=11269) | -82.3 |
| ComboPoints | [Rupture](https://www.wowhead.com/forever/spell=11275) | -62.9 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -39.6 |

## Priest — Smite

**Talents:** 31/17/3 · `515030031305001031-00505023002-003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** The rank-2 Smite fallback depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 562.66 | 0.15 | 0.00 |
| Undead | 566.75 | 0.15 | 0.00 |
| Human | 560.00 | 0.15 | 0.00 |
| Dwarf | 559.52 | 0.14 | 0.00 |
| Night Elf | 562.13 | 0.15 | 0.00 |
| Gnome | 562.83 | 0.14 | 0.00 |

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
| [Smite](https://www.wowhead.com/forever/spell=10934) | 201.38 |
| [Penance](https://www.wowhead.com/forever/spell=1316995) | 140.06 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 128.80 |
| [Smite](https://www.wowhead.com/forever/spell=591) | 87.85 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 8.66 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -13516.4 |
| Mana | OtherActionManaRegen (tag 1) | +9164.0 |
| Mana | [Penance](https://www.wowhead.com/forever/spell=1316995) | -8295.4 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -6273.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4273.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3598.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3592.3 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1552.9 |
| Mana | [Smite](https://www.wowhead.com/forever/spell=591) | -1004.5 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |

## Shaman — Stormcaller

**Talents:** 28/23/0 · `550032150010303-055030031004002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** The rank-2 Lightning Bolt filler depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 539.79 | 0.22 | 0.29 |
| Tauren | 538.79 | 0.22 | 0.27 |
| Troll | 536.41 | 0.22 | 0.29 |
| Windshaper | 537.24 | 0.22 | 0.34 |
| Dwarf | 535.32 | 0.22 | 0.27 |

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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 287.79 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 86.87 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 82.95 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 33.23 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 28.81 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 14.35 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 4.38 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.40 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -17419.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5285.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5284.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.9 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3250.4 |
| Mana | OtherActionManaRegen (tag 1) | +3081.3 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | -1151.9 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | OtherActionManaRegen (tag 2) | +140.5 |

## Shaman — Elemental

**Talents:** 31/7/13 · `5502301500123031-052-05305`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 510.57 | 0.21 | 0.13 |
| Tauren | 510.23 | 0.21 | 0.13 |
| Troll | 508.67 | 0.21 | 0.11 |
| Windshaper | 508.39 | 0.21 | 0.13 |
| Dwarf | 506.14 | 0.21 | 0.13 |

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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 182.33 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | 97.10 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 84.34 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 78.73 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 32.56 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 21.05 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 9.17 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 4.24 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +7281.2 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | -6860.5 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -6583.4 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -5558.1 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | -5438.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4829.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3599.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3554.0 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -765.0 |
| Mana | OtherActionManaRegen (tag 2) | +163.1 |
