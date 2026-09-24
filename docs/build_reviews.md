# Build reviews

These summaries describe the published loadouts. They are simulation results, not independent confirmation of server mechanics or proof of a global optimum.

The tables and [matrix](../artifacts/forever_dps_5min.png) use the same 174 common-seed replays. Equipment selections came from an earlier mechanics revision; these results use the corrected engine. Historical search gains are not directly comparable to this release.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and paid shared-hit normalization. [Scenario and exchange model](../tools/forever_bench/README.md) · [In-game checks](in_game_checks.md)

## Paladin — Retribution

**Talents:** 13/7/31 · `253003-232-052052310012330301`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Separate seal Echoes can coexist and fire on a landed white swing; verify their simultaneous behavior in game (T53). Lower-rank seals and Consecration also need confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1138.41 | 0.59 | 0.06 |
| Human | 1134.19 | 0.61 | 0.21 |
| Dwarf | 1144.18 | 0.60 | 0.06 |

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
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 312.41 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 200.76 |
| Auto-attack (tag 1) | 180.42 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 100.09 |
| Auto-attack (tag 3) | 89.25 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 69.83 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 68.00 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 64.69 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -11531.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11181.9 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -10457.2 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3779.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3597.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3585.4 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3189.1 |
| Mana | OtherActionManaRegen (tag 1) | +3109.3 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -2723.5 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -2334.7 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -521.7 |
| Mana | OtherActionManaRegen (tag 2) | +12.7 |

## Hunter — Pet/Melee

**Talents:** 16/10/25 · `53200005001-005005-5302002300502201`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Hawk damage uses an approximate guardian/pet model. Tracking talents affect its comparison with Survival on different creature types.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 861.84 | 0.42 | 0.00 |
| Windshaper | 865.87 | 0.43 | 0.00 |
| Orc | 876.25 | 0.44 | 0.00 |
| Human | 882.09 | 0.44 | 0.00 |
| Night Elf | 874.77 | 0.43 | 0.00 |
| Dwarf | 857.87 | 0.42 | 0.00 |
| Troll | 864.82 | 0.43 | 0.00 |
| High Order | 865.87 | 0.43 | 0.00 |

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
| Auto-attack (tag 2) | 163.88 |
| Auto-attack (tag 1) | 125.88 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 110.06 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 104.22 |
| Cat: Auto-attack (tag 1) | 96.13 |
| Auto-attack (tag 3) | 76.45 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 72.56 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 63.75 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8717.9 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -3498.3 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3204.7 |
| Mana | OtherActionManaRegen (tag 1) | +3050.8 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -1735.6 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -1322.9 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -977.9 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -825.2 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Paladin — Physical Ret

**Talents:** 13/7/31 · `250003003-232-052253310012330001`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Separate seal Echoes still require the T53 in-game check. This is a Strength/AP-oriented Ret comparison, not a proven best build.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 869.33 | 0.55 | 0.02 |
| Human | 878.36 | 0.56 | 0.01 |
| Dwarf | 874.64 | 0.55 | 0.00 |

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
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 299.16 |
| Auto-attack (tag 1) | 230.87 |
| Auto-attack (tag 3) | 112.88 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 90.63 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 59.56 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 27.24 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 25.66 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 15.00 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -11829.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11207.3 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -10724.4 |
| Mana | OtherActionManaRegen (tag 1) | +5181.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3603.2 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3536.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3286.9 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -2993.7 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -1607.9 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1281.0 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -522.0 |
| Mana | OtherActionManaRegen (tag 2) | +7.1 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 853.33 | 0.45 | 0.00 |
| Tauren | 841.09 | 0.45 | 0.00 |
| Troll | 841.78 | 0.44 | 0.00 |
| Windshaper | 838.97 | 0.43 | 0.00 |
| Human | 857.98 | 0.45 | 0.00 |
| Dwarf | 836.91 | 0.44 | 0.00 |
| Night Elf | 850.33 | 0.45 | 0.00 |
| High Order | 838.97 | 0.43 | 0.00 |

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
| Auto-attack (tag 2) | 168.73 |
| Auto-attack (tag 1) | 129.81 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 110.51 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 102.91 |
| Cat: Auto-attack (tag 1) | 83.62 |
| Auto-attack (tag 3) | 80.03 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 65.68 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 53.84 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +16656.7 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -9572.0 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4352.0 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -3920.0 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3397.7 |
| Mana | OtherActionManaRegen (tag 1) | +3154.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2163.7 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -2037.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +521.8 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Hunter — Beast Mastery

**Talents:** 31/20/0 · `5320001505101251-00531510005`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 835.07 | 0.27 | 0.01 |
| Tauren | 845.98 | 0.29 | 0.00 |
| Troll | 848.10 | 0.30 | 0.00 |
| Windshaper | 846.56 | 0.29 | 0.00 |
| Human | 837.57 | 0.27 | 0.01 |
| Dwarf | 843.57 | 0.29 | 0.00 |
| Night Elf | 854.30 | 0.29 | 0.00 |
| High Order | 846.56 | 0.29 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Night Elf

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

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| Shoot | 358.14 |
| Cat: Auto-attack (tag 1) | 149.42 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | 125.94 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 82.24 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 66.34 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 48.66 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 13.07 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 10.48 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | -9083.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7391.7 |
| Mana | OtherActionManaRegen (tag 1) | +6020.1 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5474.3 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -4614.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3602.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3596.4 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3116.8 |
| Mana | OtherActionManaRegen (tag 2) | +762.1 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Intimidation](https://www.wowhead.com/forever/spell=19577) | -378.3 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 844.76 | 0.39 | 0.00 |
| Troll | 840.23 | 0.38 | 0.00 |
| Undead | 849.54 | 0.39 | 0.00 |
| Human | 836.48 | 0.39 | 0.00 |
| Gnome | 847.49 | 0.38 | 0.00 |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 309.87 |
| Succubus: Auto-attack (tag 1) | 91.88 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 90.40 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 84.31 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 82.25 |
| Succubus: [Demonic Brand](https://www.wowhead.com/forever/spell=1293697) | 62.53 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 49.46 |
| [Soul Fire](https://www.wowhead.com/forever/spell=17924) | 45.67 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -24423.5 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +13047.4 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +10306.3 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6576.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5283.7 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5146.6 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -4184.1 |
| Mana | OtherActionManaRegen (tag 1) | +3877.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3859.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3516.3 |
| Mana | [Soul Fire](https://www.wowhead.com/forever/spell=17924) | -3105.2 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |

## Hunter — Marksmanship

**Talents:** 5/35/11 · `5-0050552011523051-50005001`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 805.54 | 0.38 | 2.01 |
| Tauren | 796.33 | 0.37 | 1.23 |
| Troll | 815.21 | 0.40 | 1.45 |
| Windshaper | 814.46 | 0.40 | 1.40 |
| Human | 810.07 | 0.38 | 1.58 |
| Dwarf | 791.01 | 0.37 | 2.13 |
| Night Elf | 805.97 | 0.37 | 1.99 |
| High Order | 814.46 | 0.40 | 1.40 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Troll

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

### Damage breakdown — Troll

| Action | DPS |
|---|---:|
| Shoot | 410.81 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 192.15 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 83.46 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 75.67 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 53.13 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -11338.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7338.4 |
| Mana | OtherActionManaRegen (tag 1) | +5873.0 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5803.9 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3951.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3615.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3603.2 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -3467.4 |
| Mana | OtherActionManaRegen (tag 2) | +684.6 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Priest — Shadow

**Talents:** 18/0/33 · `305030001303--504020501201302251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 803.61 | 0.23 | 0.00 |
| Undead | 811.74 | 0.24 | 0.00 |
| Human | 802.99 | 0.23 | 0.00 |
| Dwarf | 802.88 | 0.23 | 0.00 |
| Night Elf | 809.02 | 0.23 | 0.00 |
| Gnome | 809.20 | 0.23 | 0.00 |

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
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 154.44 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 144.71 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 97.01 |
| [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | 84.55 |
| [Shadow Word: Death Backlash](https://www.wowhead.com/forever/spell=1309598) | 26.55 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 13.68 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +9117.6 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -6188.8 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -5938.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4096.0 |
| Mana | [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | -3199.2 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2847.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2398.2 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1598.1 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1514.1 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +50.1 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2525000013520105--0550015103`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 804.43 | 0.45 | 0.01 |
| Troll | 799.65 | 0.45 | 0.01 |
| Undead | 805.26 | 0.45 | 0.01 |
| Human | 796.15 | 0.44 | 0.01 |
| Gnome | 809.58 | 0.46 | 0.00 |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 372.57 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 105.84 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 98.67 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 84.14 |
| Succubus: Auto-attack (tag 1) | 81.05 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 39.69 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 18.93 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 3.84 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -30127.7 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24903.7 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6456.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5257.1 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5104.2 |
| Mana | OtherActionManaRegen (tag 1) | +3869.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3598.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3420.7 |
| Mana | [Siphon Life](https://www.wowhead.com/forever/spell=18881) | -3385.4 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -352.4 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -338.5 |

## Warrior — Fury

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Level-60 rage, queued off-hand hit and Flurry charge timing still need in-game confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 791.77 | 0.47 | 0.00 |
| Tauren | 778.48 | 0.47 | 0.00 |
| Troll | 774.80 | 0.47 | 0.00 |
| Undead | 799.70 | 0.47 | 0.00 |
| Windshaper | 780.56 | 0.45 | 0.00 |
| Human | 780.88 | 0.47 | 0.00 |
| Dwarf | 778.93 | 0.47 | 0.00 |
| Night Elf | 770.64 | 0.47 | 0.00 |
| Gnome | 778.39 | 0.45 | 0.00 |
| High Order | 780.56 | 0.45 | 0.00 |

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

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 790.01 | 0.31 | 0.00 |
| Troll | 785.09 | 0.30 | 0.00 |
| Undead | 795.20 | 0.31 | 0.00 |
| Human | 783.03 | 0.31 | 0.00 |
| Gnome | 791.03 | 0.31 | 0.00 |

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
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 258.14 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 113.43 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 90.21 |
| Succubus: Auto-attack (tag 1) | 89.26 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 83.35 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 77.37 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 49.43 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 14.37 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24365.1 |
| Mana | [Incinerate](https://www.wowhead.com/forever/spell=1293813) | -21119.7 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6698.8 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -6280.7 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -5664.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5218.3 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5159.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4513.6 |
| Mana | OtherActionManaRegen (tag 1) | +3878.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3602.1 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -461.7 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `252200001351-0025003001-0550005103`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 765.87 | 0.49 | 0.00 |
| Troll | 760.33 | 0.48 | 0.00 |
| Undead | 768.27 | 0.48 | 0.00 |
| Human | 756.87 | 0.48 | 0.00 |
| Gnome | 770.84 | 0.49 | 0.00 |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 439.92 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 116.60 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 104.48 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 82.68 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 21.45 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 3.50 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.22 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -32381.1 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +22911.3 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6414.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5354.2 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5223.4 |
| Mana | OtherActionManaRegen (tag 1) | +3887.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3650.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3604.1 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -523.6 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -353.0 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -203.6 |

## Warrior — 2H Bloodthirst

**Talents:** 20/31/0 · `30304203032-050500320051310051`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 746.80 | 0.49 | 0.00 |
| Tauren | 735.50 | 0.49 | 0.00 |
| Undead | 740.46 | 0.48 | 0.00 |
| Troll | 736.22 | 0.47 | 0.00 |
| Windshaper | 728.55 | 0.48 | 0.00 |
| Human | 741.71 | 0.48 | 0.00 |
| Dwarf | 737.47 | 0.49 | 0.00 |
| Gnome | 736.32 | 0.48 | 0.00 |
| Night Elf | 733.92 | 0.48 | 0.00 |
| High Order | 728.55 | 0.48 | 0.00 |

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
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 41.93 |

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

## Shaman — Enhancement

**Talents:** 20/31/0 · `420033152-053030031005102251`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 737.67 | 0.58 | 0.05 |
| Tauren | 736.88 | 0.56 | 0.01 |
| Troll | 731.59 | 0.56 | 0.02 |
| Windshaper | 737.13 | 0.56 | 0.02 |
| Dwarf | 743.19 | 0.56 | 0.02 |

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
| Auto-attack (tag 1) | 208.81 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 138.49 |
| [Fire Nova](https://www.wowhead.com/forever/spell=408345) | 97.76 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 93.45 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 55.04 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 51.51 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 45.85 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 30.56 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Nova](https://www.wowhead.com/forever/spell=408345) | -14067.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8969.3 |
| Mana | OtherActionManaRegen (tag 1) | +6820.6 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4979.1 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -4491.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4380.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3599.8 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3218.9 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | [Grace of Air Totem](https://www.wowhead.com/forever/spell=10627) | -499.6 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -450.0 |
| Mana | OtherActionManaRegen (tag 2) | +336.4 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-550500000002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 735.64 | 0.49 | 0.00 |
| Tauren | 732.62 | 0.49 | 0.00 |
| Troll | 730.47 | 0.47 | 0.00 |
| Undead | 739.08 | 0.50 | 0.00 |
| Windshaper | 732.18 | 0.50 | 0.00 |
| Human | 738.48 | 0.49 | 0.00 |
| Dwarf | 728.16 | 0.49 | 0.00 |
| Night Elf | 734.01 | 0.49 | 0.00 |
| Gnome | 732.42 | 0.48 | 0.00 |
| High Order | 732.18 | 0.50 | 0.00 |

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
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 33.68 |

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
| Troll | 711.93 | 0.37 | 0.00 |
| Undead | 719.11 | 0.38 | 0.00 |
| Human | 706.75 | 0.37 | 0.00 |
| Gnome | 712.02 | 0.38 | 0.00 |
| High Order | 713.45 | 0.37 | 0.00 |

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
| Troll | 695.78 | 0.50 | 0.08 |
| Undead | 706.03 | 0.51 | 0.06 |
| Human | 691.83 | 0.49 | 0.04 |
| Gnome | 699.37 | 0.49 | 0.02 |
| High Order | 696.04 | 0.51 | 0.08 |

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
| Orc | 686.67 | 0.45 | 3.11 |
| Troll | 683.18 | 0.44 | 2.68 |
| Undead | 696.05 | 0.44 | 2.83 |
| Human | 685.59 | 0.44 | 2.13 |
| Gnome | 687.78 | 0.44 | 1.63 |
| High Order | 681.83 | 0.41 | 1.05 |

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
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 146.01 |
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
| Troll | 681.90 | 0.41 | 0.15 |
| Undead | 685.35 | 0.41 | 0.14 |
| Human | 676.87 | 0.41 | 0.12 |
| Gnome | 683.31 | 0.42 | 0.04 |
| High Order | 683.75 | 0.41 | 0.17 |

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
| Tauren | 678.39 | 0.21 | 0.00 |
| Windshaper | 680.02 | 0.20 | 0.00 |
| Night Elf | 679.55 | 0.21 | 0.00 |
| High Order | 680.02 | 0.20 | 0.00 |

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
| [Shred](https://www.wowhead.com/forever/spell=9830) | 265.25 |
| Auto-attack (tag 1) | 226.89 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 180.53 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 4.71 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.63 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -3451.1 |
| Energy | OtherActionEnergyRegen | +3047.4 |
| Energy | [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) | +689.6 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -588.4 |
| Energy | OtherActionRefund | +180.0 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.0 |
| ComboPoints | [Rip](https://www.wowhead.com/forever/spell=9896) | -98.3 |
| ComboPoints | [Shred](https://www.wowhead.com/forever/spell=9830) | +66.2 |
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
| Orc | 618.11 | 0.29 | 0.00 |
| Troll | 615.32 | 0.29 | 0.00 |
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
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 28.64 |

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
| Orc | 606.96 | 0.27 | 0.00 |
| Troll | 606.48 | 0.27 | 0.00 |
| Undead | 624.96 | 0.28 | 0.00 |
| Windshaper | 609.50 | 0.27 | 0.00 |
| Human | 610.43 | 0.28 | 0.00 |
| Dwarf | 600.86 | 0.27 | 0.00 |
| Night Elf | 608.88 | 0.27 | 0.00 |
| Gnome | 608.29 | 0.28 | 0.00 |
| High Order | 609.50 | 0.27 | 0.00 |

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
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 25.44 |
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
| Tauren | 605.82 | 0.26 | 0.00 |
| Windshaper | 605.24 | 0.26 | 0.00 |
| Night Elf | 606.15 | 0.26 | 0.00 |
| High Order | 605.24 | 0.26 | 0.00 |

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
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 354.33 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 91.46 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 85.96 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 74.40 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=25298) | -15141.5 |
| Mana | OtherActionManaRegen (tag 1) | +7736.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4496.5 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -4212.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3600.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3589.2 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1556.5 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1205.2 |
| Mana | [Moonkin Form](https://www.wowhead.com/forever/spell=24858) | -435.4 |
| Mana | OtherActionManaRegen (tag 2) | +76.0 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | +3.5 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | -0.1 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 556.75 | 0.21 | 0.00 |
| Troll | 558.50 | 0.22 | 0.00 |
| Undead | 579.83 | 0.23 | 0.00 |
| Windshaper | 556.73 | 0.22 | 0.00 |
| Human | 562.87 | 0.22 | 0.00 |
| Dwarf | 551.15 | 0.21 | 0.00 |
| Night Elf | 560.25 | 0.22 | 0.00 |
| Gnome | 559.61 | 0.22 | 0.00 |
| High Order | 556.73 | 0.22 | 0.00 |

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
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 39.43 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 36.34 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 28.02 |
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
| Troll | 563.16 | 0.14 | 0.00 |
| Undead | 566.47 | 0.15 | 0.00 |
| Human | 559.69 | 0.14 | 0.00 |
| Dwarf | 559.20 | 0.14 | 0.00 |
| Night Elf | 562.05 | 0.14 | 0.00 |
| Gnome | 563.93 | 0.14 | 0.00 |

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
| [Smite](https://www.wowhead.com/forever/spell=10934) | 201.73 |
| [Penance](https://www.wowhead.com/forever/spell=1316995) | 140.01 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 128.51 |
| [Smite](https://www.wowhead.com/forever/spell=591) | 87.50 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 8.71 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -13536.0 |
| Mana | OtherActionManaRegen (tag 1) | +9166.2 |
| Mana | [Penance](https://www.wowhead.com/forever/spell=1316995) | -8293.5 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -6275.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4267.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3612.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3603.9 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1557.2 |
| Mana | [Smite](https://www.wowhead.com/forever/spell=591) | -1001.3 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |

## Shaman — Stormcaller

**Talents:** 28/23/0 · `550032150010303-055030031004002`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)

**Model limitations:** The rank-2 Lightning Bolt filler depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 539.87 | 0.22 | 0.29 |
| Tauren | 538.85 | 0.22 | 0.28 |
| Troll | 535.50 | 0.21 | 0.28 |
| Windshaper | 537.34 | 0.21 | 0.25 |
| Dwarf | 535.23 | 0.22 | 0.31 |

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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 287.72 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 86.82 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 82.97 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 33.24 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 28.83 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 14.46 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 4.38 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.46 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -17435.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5290.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5277.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.9 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3248.0 |
| Mana | OtherActionManaRegen (tag 1) | +3082.2 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | -1150.7 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | OtherActionManaRegen (tag 2) | +137.7 |

## Shaman — Elemental

**Talents:** 31/7/13 · `5502301500123031-052-05305`

[Requests and results](../artifacts/forever_dps_5min.json) · [Equipment search](../artifacts/gear_search/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 510.44 | 0.21 | 0.13 |
| Tauren | 510.01 | 0.21 | 0.13 |
| Troll | 508.19 | 0.21 | 0.12 |
| Windshaper | 508.20 | 0.21 | 0.11 |
| Dwarf | 505.91 | 0.21 | 0.13 |

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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 182.06 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | 97.13 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 84.55 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 78.69 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 32.56 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 21.12 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 9.06 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 4.21 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +7281.2 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | -6850.6 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -6585.3 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -5571.2 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | -5444.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4837.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3603.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3553.6 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -765.0 |
| Mana | OtherActionManaRegen (tag 2) | +163.3 |
