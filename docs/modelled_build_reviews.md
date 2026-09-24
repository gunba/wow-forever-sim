# Build reviews

These are hypothetical modeled-gear results, not obtainable item rankings. Sources linked for modeled items are allocation references only and have different stats. Gear is compared under the [model assumptions](modelled_gear.md).

The tables and [matrix](../artifacts/modelled_gear/forever_dps_5min.png) use the same 184 common-seed replays. The modeled search and original real-item benchmark are separate; neither proves available launch gear.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and paid shared-hit normalization. [Scenario and exchange model](../tools/forever_bench/README.md) · [In-game checks](in_game_checks.md)

## Paladin — Retribution

**Talents:** 13/7/31 · `253003-232-052052310012330301`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Separate seal Echoes can coexist and fire on a landed white swing; verify their simultaneous behavior in game (T53). Lower-rank seals and Consecration also need confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1272.56 | 0.63 | 0.00 |
| Human | 1268.62 | 0.65 | 0.00 |
| Dwarf | 1273.24 | 0.63 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Dwarf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / intellect / spell power / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | — |
| Shoulders | [Modeled: Strength / intellect / spell power / hit — Shoulders (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / intellect / spell power / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / intellect / spell power / hit — Chest (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / intellect / spell power / hit — Wrists (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / intellect / spell power / hit — Hands (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Lesser Arcanum of Voracity |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / intellect / spell power / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | — |
| Ring 2 | [Modeled: Strength / intellect / spell power / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / intellect / spell power / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / intellect / spell power / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | — |
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
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 349.86 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 223.96 |
| Auto-attack (tag 1) | 195.52 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 111.91 |
| Auto-attack (tag 3) | 96.54 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 83.29 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 77.75 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 70.41 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -11820.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11248.8 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -10715.0 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3773.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3608.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3596.1 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3181.6 |
| Mana | OtherActionManaRegen (tag 1) | +3112.2 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -2885.8 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -2093.4 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -521.4 |
| Mana | OtherActionManaRegen (tag 2) | +2.4 |

## Paladin — Physical Ret

**Talents:** 13/7/31 · `250003003-232-052250310012330301`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Strength/AP equipment is unchanged. Champion of the Light converts existing Intellect into spell power without equipping caster gear; the separate seal Echoes still require the T53 in-game check.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1090.03 | 0.64 | 0.00 |
| Human | 1093.77 | 0.66 | 0.00 |
| Dwarf | 1086.66 | 0.65 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Titanic Leggings](https://www.wowhead.com/forever/item=22385) | 60 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
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
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 359.64 |
| Auto-attack (tag 1) | 263.43 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 129.32 |
| Auto-attack (tag 3) | 127.80 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 79.65 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 39.19 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 37.90 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 29.45 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -11974.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11246.4 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -10853.1 |
| Mana | OtherActionManaRegen (tag 1) | +5183.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3586.3 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3546.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3263.1 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3000.8 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -1684.2 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1215.5 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -522.0 |
| Mana | OtherActionManaRegen (tag 2) | +3.2 |

## Hunter — Beast Mastery

**Talents:** 31/20/0 · `5320001505101251-00531510005`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 1069.07 | 0.31 | 0.81 |
| Tauren | 1075.97 | 0.34 | 0.42 |
| Troll | 1078.59 | 0.34 | 0.42 |
| Windshaper | 1080.31 | 0.34 | 0.33 |
| Human | 1070.79 | 0.31 | 0.61 |
| Dwarf | 1072.69 | 0.34 | 0.36 |
| Night Elf | 1085.28 | 0.34 | 0.35 |
| High Order | 1080.31 | 0.34 | 0.33 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Night Elf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Minor Agility |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Sentinel's Chain Leggings](https://www.wowhead.com/forever/item=22748) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
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
| Shoot | 414.03 |
| Cat: Auto-attack (tag 1) | 249.16 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | 145.20 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 89.83 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 73.27 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 57.39 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 54.29 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 2.11 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | -9084.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7372.0 |
| Mana | OtherActionManaRegen (tag 1) | +5898.5 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5413.1 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -4614.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3606.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3604.7 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3118.5 |
| Mana | OtherActionManaRegen (tag 2) | +741.8 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Intimidation](https://www.wowhead.com/forever/spell=19577) | -305.7 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Hunter — Pet/Melee

**Talents:** 16/10/25 · `53200005001-005005-5302002300502201`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Hawk damage uses an approximate guardian/pet model. Tracking talents affect its comparison with Survival on different creature types.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 1029.16 | 0.46 | 0.00 |
| Windshaper | 1033.56 | 0.46 | 0.00 |
| Orc | 1035.05 | 0.49 | 0.00 |
| Human | 1027.83 | 0.48 | 0.00 |
| Night Elf | 1034.98 | 0.48 | 0.00 |
| Dwarf | 1025.70 | 0.46 | 0.00 |
| Troll | 1030.42 | 0.46 | 0.00 |
| High Order | 1033.56 | 0.46 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Outrider's Chain Leggings](https://www.wowhead.com/forever/item=22673) | 65 | Lesser Arcanum of Voracity |
| Feet | [Scalegut Treaders](https://www.wowhead.com/forever/item=275618) | 58 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Premier High Warlord's Cleaver](https://www.wowhead.com/forever/item=272592) | 65 | Enchant Weapon - Crusader |
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

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 176.84 |
| Cat: Auto-attack (tag 1) | 161.74 |
| Auto-attack (tag 1) | 136.79 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 119.58 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 114.11 |
| Auto-attack (tag 3) | 84.07 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 76.65 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 69.39 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +8262.0 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -3546.6 |
| Mana | OtherActionManaRegen (tag 1) | +3334.3 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3076.0 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -1731.3 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -1322.3 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -884.7 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -828.7 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -100.0 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 996.06 | 0.50 | 0.00 |
| Tauren | 988.75 | 0.49 | 0.00 |
| Troll | 991.85 | 0.47 | 0.00 |
| Windshaper | 993.18 | 0.47 | 0.00 |
| Human | 999.94 | 0.49 | 0.00 |
| Dwarf | 987.43 | 0.48 | 0.00 |
| Night Elf | 995.72 | 0.48 | 0.00 |
| High Order | 993.18 | 0.47 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Sentinel's Chain Leggings](https://www.wowhead.com/forever/item=22748) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Blackstone Ring](https://www.wowhead.com/forever/item=17713) | 54 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
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
| Auto-attack (tag 2) | 178.66 |
| Cat: Auto-attack (tag 1) | 146.04 |
| Auto-attack (tag 1) | 137.86 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 121.30 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 113.69 |
| Auto-attack (tag 3) | 85.81 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 72.33 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 52.16 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +16618.4 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -9538.2 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4333.8 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -3666.3 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3398.1 |
| Mana | OtherActionManaRegen (tag 1) | +3156.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2179.3 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -2039.3 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +191.2 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |
| Mana | OtherActionManaRegen (tag 2) | +1.3 |

## Hunter — Marksmanship

**Talents:** 5/35/11 · `5-0050552011523051-50005001`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 917.78 | 0.43 | 2.89 |
| Tauren | 906.11 | 0.43 | 2.98 |
| Troll | 926.83 | 0.45 | 2.20 |
| Windshaper | 927.97 | 0.45 | 1.97 |
| Human | 921.88 | 0.43 | 2.40 |
| Dwarf | 902.43 | 0.43 | 2.98 |
| Night Elf | 917.90 | 0.43 | 2.90 |
| High Order | 927.97 | 0.45 | 1.97 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Minor Agility |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Outrider's Chain Leggings](https://www.wowhead.com/forever/item=22673) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
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
| Shoot | 473.99 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 221.00 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 90.68 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 82.80 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 59.50 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -11340.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7322.6 |
| Mana | OtherActionManaRegen (tag 1) | +5747.8 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5701.0 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3963.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3609.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3596.9 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -3392.7 |
| Mana | OtherActionManaRegen (tag 2) | +673.6 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Warrior — Fury

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Level-60 rage, queued off-hand hit and Flurry charge timing still need in-game confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 892.97 | 0.50 | 0.00 |
| Tauren | 889.10 | 0.50 | 0.00 |
| Troll | 886.15 | 0.49 | 0.00 |
| Undead | 910.87 | 0.50 | 0.00 |
| Windshaper | 893.44 | 0.49 | 0.00 |
| Human | 887.64 | 0.49 | 0.00 |
| Dwarf | 889.24 | 0.49 | 0.00 |
| Night Elf | 884.43 | 0.49 | 0.00 |
| Gnome | 889.16 | 0.49 | 0.00 |
| High Order | 893.44 | 0.49 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Lionheart Helm](https://www.wowhead.com/forever/item=12640) | 61 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Titanic Leggings](https://www.wowhead.com/forever/item=22385) | 60 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
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
| Auto-attack (tag 2) | 167.46 |
| Auto-attack (tag 1) | 145.34 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 126.91 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 106.43 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 97.74 |
| Auto-attack (tag 3) | 63.62 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 55.30 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 46.78 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +1116.0 |
| Rage | Auto-attack (tag 1) | +917.4 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -884.2 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -621.4 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -497.2 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -452.4 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +181.8 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +180.2 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -150.1 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +100.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +98.8 |
| Rage | OtherActionRefund | +95.8 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 898.51 | 0.40 | 0.00 |
| Troll | 895.49 | 0.40 | 0.00 |
| Undead | 899.25 | 0.41 | 0.00 |
| Human | 890.10 | 0.40 | 0.00 |
| Gnome | 904.41 | 0.41 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Shadow Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Crit / shadow power / intellect — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Crit / shadow power / intellect — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
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

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 344.37 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 96.87 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 89.61 |
| Succubus: Auto-attack (tag 1) | 86.64 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 85.19 |
| Succubus: [Demonic Brand](https://www.wowhead.com/forever/spell=1293697) | 71.46 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 53.17 |
| [Soul Fire](https://www.wowhead.com/forever/spell=17924) | 48.50 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -24706.9 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +12735.9 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +10317.6 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6411.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5396.2 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -4949.0 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -4041.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3933.0 |
| Mana | OtherActionManaRegen (tag 1) | +3116.7 |
| Mana | [Soul Fire](https://www.wowhead.com/forever/spell=17924) | -3066.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2391.9 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -909.0 |

## Priest — Shadow

**Talents:** 18/0/33 · `305030001303--504020501201302251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 860.94 | 0.25 | 0.00 |
| Undead | 873.54 | 0.26 | 0.00 |
| Human | 860.82 | 0.25 | 0.00 |
| Dwarf | 860.75 | 0.25 | 0.00 |
| Night Elf | 866.98 | 0.25 | 0.00 |
| Gnome | 862.75 | 0.25 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Crit / shadow power / intellect — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Arcanum of Focus |
| Neck | [Modeled: Crit / shadow power / intellect — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Shoulders | [Modeled: Crit / shadow power / intellect — Shoulders (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Crit / shadow power / intellect — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Shadow Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Crit / shadow power / intellect — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Ring 2 | [Modeled: Crit / shadow power / intellect — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Crit / shadow power / intellect — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
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
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 339.59 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 165.39 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 154.97 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 102.89 |
| [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | 95.37 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 15.34 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +8071.8 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -7103.5 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -5951.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4382.0 |
| Mana | [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | -3396.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2931.1 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2849.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +1771.5 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1590.8 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |

## Warrior — 2H Bloodthirst

**Talents:** 20/31/0 · `30304203032-050500320051310051`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 859.75 | 0.52 | 0.00 |
| Tauren | 848.02 | 0.51 | 0.00 |
| Undead | 846.17 | 0.53 | 0.00 |
| Troll | 848.13 | 0.52 | 0.00 |
| Windshaper | 838.35 | 0.52 | 0.00 |
| Human | 853.60 | 0.51 | 0.00 |
| Dwarf | 848.99 | 0.51 | 0.00 |
| Gnome | 847.22 | 0.51 | 0.00 |
| Night Elf | 845.32 | 0.51 | 0.00 |
| High Order | 838.35 | 0.52 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Lionheart Helm](https://www.wowhead.com/forever/item=12640) | 61 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Titanic Leggings](https://www.wowhead.com/forever/item=22385) | 60 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
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
| Auto-attack (tag 1) | 269.08 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 125.02 |
| Auto-attack (tag 3) | 110.33 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 80.37 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 72.35 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 57.84 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 57.49 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 41.94 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1869.2 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -989.5 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -542.9 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -500.2 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +174.5 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +143.1 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -140.6 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -126.5 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +92.7 |
| Rage | OtherActionRefund | +79.9 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -56.0 |
| Rage | [Berserker Stance](https://www.wowhead.com/forever/spell=2458) | -29.1 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2525000013520105--0550015103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 854.54 | 0.48 | 0.00 |
| Troll | 851.86 | 0.49 | 0.00 |
| Undead | 857.54 | 0.48 | 0.00 |
| Human | 842.59 | 0.47 | 0.00 |
| Gnome | 857.75 | 0.50 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Crit / shadow power / intellect — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Arcanum of Focus |
| Neck | [Modeled: Crit / shadow power / intellect — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Crit / shadow power / intellect — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Shadow Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 406.00 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 113.07 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 101.38 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 85.22 |
| Succubus: Auto-attack (tag 1) | 77.17 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 42.03 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 19.57 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 8.11 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -30303.4 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24800.5 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6433.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5380.6 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5085.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3598.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3437.5 |
| Mana | [Siphon Life](https://www.wowhead.com/forever/spell=18881) | -3354.9 |
| Mana | OtherActionManaRegen (tag 1) | +3104.7 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -900.0 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -351.7 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -350.2 |

## Warrior — Fury (Sunder)

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 854.49 | 0.51 | 0.00 |
| Tauren | 848.72 | 0.51 | 0.00 |
| Troll | 829.64 | 0.50 | 0.00 |
| Undead | 842.97 | 0.51 | 0.00 |
| Windshaper | 845.66 | 0.51 | 0.00 |
| Human | 846.53 | 0.51 | 0.00 |
| Dwarf | 847.35 | 0.51 | 0.00 |
| Night Elf | 831.62 | 0.50 | 0.00 |
| Gnome | 849.57 | 0.51 | 0.00 |
| High Order | 845.66 | 0.51 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Lionheart Helm](https://www.wowhead.com/forever/item=12640) | 61 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Raider Gauntlets](https://www.wowhead.com/forever/item=272095) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Titanic Leggings](https://www.wowhead.com/forever/item=22385) | 60 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Premier High Warlord's Cleaver](https://www.wowhead.com/forever/item=272592) | 65 | Enchant Weapon - Crusader |
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

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 168.70 |
| Auto-attack (tag 1) | 154.92 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 107.58 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 99.41 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 91.15 |
| Auto-attack (tag 3) | 66.93 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 54.32 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 44.14 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +1110.8 |
| Rage | Auto-attack (tag 1) | +955.6 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -798.7 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -585.3 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -458.1 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -386.5 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -273.0 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +180.8 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +180.1 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -128.2 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +100.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +100.0 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 836.59 | 0.34 | 0.00 |
| Troll | 834.85 | 0.33 | 0.00 |
| Undead | 847.69 | 0.34 | 0.00 |
| Human | 830.92 | 0.33 | 0.00 |
| Gnome | 837.98 | 0.34 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Fire Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 279.47 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 121.31 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 95.95 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 89.23 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 84.58 |
| Succubus: Auto-attack (tag 1) | 82.63 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 52.12 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 16.33 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24233.4 |
| Mana | [Incinerate](https://www.wowhead.com/forever/spell=1293813) | -21343.3 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6715.8 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -6298.3 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -5564.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5335.7 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5104.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4546.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.1 |
| Mana | OtherActionManaRegen (tag 1) | +3113.2 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -498.2 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-550500000002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 836.61 | 0.55 | 0.00 |
| Tauren | 830.83 | 0.54 | 0.00 |
| Troll | 825.33 | 0.55 | 0.00 |
| Undead | 828.94 | 0.54 | 0.00 |
| Windshaper | 826.89 | 0.54 | 0.00 |
| Human | 833.90 | 0.54 | 0.00 |
| Dwarf | 821.69 | 0.54 | 0.00 |
| Night Elf | 835.58 | 0.55 | 0.00 |
| Gnome | 832.32 | 0.54 | 0.00 |
| High Order | 826.89 | 0.54 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Lionheart Helm](https://www.wowhead.com/forever/item=12640) | 61 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Titanic Leggings](https://www.wowhead.com/forever/item=22385) | 60 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Plate) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
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

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 191.61 |
| Auto-attack (tag 3) | 135.18 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 129.10 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 127.72 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 76.66 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 51.22 |
| [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | 41.32 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 35.25 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1739.8 |
| Rage | [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | -1035.2 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -500.2 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -205.6 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -168.4 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +156.9 |
| Rage | [Rend](https://www.wowhead.com/forever/spell=11574) | -145.5 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -131.4 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +130.2 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.8 |
| Rage | OtherActionRefund | +96.3 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -95.3 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `252200001351-0025003001-0550005103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 814.94 | 0.53 | 0.00 |
| Troll | 815.91 | 0.51 | 0.00 |
| Undead | 831.65 | 0.53 | 0.00 |
| Human | 818.04 | 0.51 | 0.00 |
| Gnome | 818.31 | 0.52 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Crit / shadow power / intellect — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Fel Cape](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Crit / shadow power / intellect — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Shadow Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Crit / shadow power / intellect — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Crit / shadow power / intellect — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 474.48 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 124.97 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 108.15 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 81.81 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 22.17 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 14.08 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 3.86 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.14 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -32682.9 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24524.7 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6838.6 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5494.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5408.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3654.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3603.4 |
| Mana | OtherActionManaRegen (tag 1) | +3114.5 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -608.2 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -360.7 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -198.4 |

## Druid — Feral

**Talents:** 9/34/8 · `050022-3521002023032213041-053`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 808.58 | 0.26 | 0.00 |
| Windshaper | 811.60 | 0.26 | 0.00 |
| Night Elf | 805.45 | 0.25 | 0.00 |
| High Order | 811.60 | 0.26 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
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
| [Shred](https://www.wowhead.com/forever/spell=9830) | 299.20 |
| Auto-attack (tag 1) | 282.41 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 198.51 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 28.75 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.72 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -3315.1 |
| Energy | OtherActionEnergyRegen | +3047.4 |
| Energy | [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) | +678.6 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -570.2 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -178.3 |
| Energy | OtherActionRefund | +170.9 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.5 |
| ComboPoints | [Rip](https://www.wowhead.com/forever/spell=9896) | -94.1 |
| ComboPoints | [Shred](https://www.wowhead.com/forever/spell=9830) | +69.8 |
| ComboPoints | [Primal Fury](https://www.wowhead.com/forever/spell=37117) | +45.5 |
| ComboPoints | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -20.1 |
| Mana | OtherActionManaRegen (tag 2) | +0.0 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `420033152-053030031005102251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 793.55 | 0.61 | 0.08 |
| Tauren | 790.91 | 0.61 | 0.03 |
| Troll | 797.70 | 0.60 | 0.01 |
| Windshaper | 785.56 | 0.59 | 0.02 |
| Dwarf | 797.51 | 0.60 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Troll

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Black Dragonscale Helm](https://www.wowhead.com/forever/item=252605) | 61 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Defiler's Mail Pauldrons](https://www.wowhead.com/forever/item=20203) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Timbermaw Tunic](https://www.wowhead.com/forever/item=252484) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Chromatic Gauntlets](https://www.wowhead.com/forever/item=19157) | 70 | Enchant Gloves - Minor Haste |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Outrider's Mail Leggings](https://www.wowhead.com/forever/item=22676) | 65 | Lesser Arcanum of Voracity |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Premier High Warlord's Destroyer](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Lifestealing |
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

### Damage breakdown — Troll

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 238.89 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 154.34 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 103.84 |
| [Fire Nova](https://www.wowhead.com/forever/spell=408345) | 101.99 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 53.56 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 50.75 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 44.11 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 30.93 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Nova](https://www.wowhead.com/forever/spell=408345) | -14693.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9203.8 |
| Mana | OtherActionManaRegen (tag 1) | +7332.1 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4978.9 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -4492.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4274.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3586.2 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3213.9 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | [Grace of Air Totem](https://www.wowhead.com/forever/spell=10627) | -499.6 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -450.0 |
| Mana | OtherActionManaRegen (tag 2) | +318.0 |

## Mage — Arcane–Frost

**Talents:** 28/0/23 · `05020500310031053--05550002010003002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Uses an assumed Ice Lance coefficient and unresolved Fingers of Frost, Missile Barrage and Clearcasting timing.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 754.86 | 0.39 | 0.00 |
| Troll | 750.54 | 0.38 | 0.00 |
| Undead | 763.52 | 0.40 | 0.00 |
| Human | 750.34 | 0.38 | 0.00 |
| Gnome | 755.38 | 0.39 | 0.00 |
| High Order | 756.72 | 0.39 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Frost Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 449.02 |
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 165.87 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 134.77 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 13.86 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -21624.3 |
| Mana | OtherActionManaRegen (tag 1) | +10747.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5668.0 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -2924.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2627.8 |
| Mana | OtherActionManaRegen (tag 2) | +1724.1 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1097.9 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +808.9 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +271.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +54.2 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Fire

**Talents:** 17/31/3 · `0501252000002-23450000130133051-003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 736.22 | 0.48 | 2.94 |
| Troll | 739.45 | 0.48 | 2.69 |
| Undead | 755.22 | 0.48 | 3.01 |
| Human | 726.26 | 0.46 | 2.59 |
| Gnome | 738.14 | 0.47 | 1.99 |
| High Order | 730.24 | 0.43 | 1.46 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / fire power / crit — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / fire power / crit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / fire power / crit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / fire power / crit — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Fire Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / fire power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | — |
| Ring 2 | [Channeler's Ring](https://www.wowhead.com/forever/item=272406) | 65 | — |
| Trinket 1 | [Modeled: Passive Intellect / fire power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / fire power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | — |
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
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 286.77 |
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 173.19 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 172.80 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 101.44 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 14.11 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 6.90 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -16129.0 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -14594.3 |
| Mana | OtherActionManaRegen (tag 1) | +13338.2 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -7591.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5562.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4721.2 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +4230.6 |
| Mana | OtherActionManaRegen (tag 2) | +1967.4 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.3 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +850.1 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +599.5 |
| Mana | [Fireball](https://www.wowhead.com/forever/spell=25306) | -496.7 |

## Mage — Arcane

**Talents:** 33/3/15 · `053005023100311531-03-005500032`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 729.57 | 0.52 | 0.05 |
| Troll | 727.44 | 0.51 | 0.05 |
| Undead | 734.57 | 0.53 | 0.02 |
| Human | 719.30 | 0.51 | 0.03 |
| Gnome | 738.76 | 0.52 | 0.01 |
| High Order | 728.71 | 0.53 | 0.07 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 357.65 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 337.59 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 43.52 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -30362.4 |
| Mana | OtherActionManaRegen (tag 1) | +10805.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +6924.5 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4327.7 |
| Mana | OtherActionManaRegen (tag 2) | +3784.9 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -1548.3 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.3 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.5 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +552.7 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | -526.7 |

## Rogue — Mutilate

**Talents:** 31/20/0 · `0053031035140105-302303202014`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 701.80 | 0.31 | 0.00 |
| Troll | 698.71 | 0.32 | 0.00 |
| Undead | 727.13 | 0.33 | 0.00 |
| Windshaper | 696.51 | 0.32 | 0.00 |
| Human | 696.22 | 0.32 | 0.00 |
| Dwarf | 698.37 | 0.31 | 0.00 |
| Night Elf | 699.43 | 0.31 | 0.00 |
| Gnome | 706.10 | 0.31 | 0.00 |
| High Order | 696.51 | 0.32 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Agility / crit / hit — Waist (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
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
| Auto-attack (tag 1) | 183.05 |
| Auto-attack (tag 2) | 109.75 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 90.48 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 84.59 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 67.46 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 48.78 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 47.01 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 30.24 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3013.4 |
| Energy | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | -2872.6 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -852.8 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +798.1 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -407.7 |
| Energy | OtherActionRefund | +149.7 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -139.9 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.6 |
| ComboPoints | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | +89.5 |
| ComboPoints | [Seal Fate](https://www.wowhead.com/forever/spell=14195) | +38.9 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +33.6 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -20.4 |

## Mage — Frost

**Talents:** 11/3/37 · `050005001-03-0555003321001301251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 715.13 | 0.43 | 0.17 |
| Troll | 711.29 | 0.42 | 0.15 |
| Undead | 715.65 | 0.42 | 0.15 |
| Human | 706.24 | 0.42 | 0.13 |
| Gnome | 712.36 | 0.43 | 0.02 |
| High Order | 713.04 | 0.42 | 0.15 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Hide of the Wild](https://www.wowhead.com/forever/item=18510) | 62 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of Spell Mastery](https://www.wowhead.com/forever/item=14146) | 62 | Enchant Gloves - Frost Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 513.02 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 193.58 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 9.05 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -23520.3 |
| Mana | OtherActionManaRegen (tag 1) | +9943.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3948.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3713.9 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3593.5 |
| Mana | OtherActionManaRegen (tag 2) | +2286.7 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.1 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +850.6 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +600.3 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 688.49 | 0.31 | 0.00 |
| Troll | 685.01 | 0.31 | 0.00 |
| Undead | 707.05 | 0.32 | 0.00 |
| Windshaper | 681.11 | 0.30 | 0.00 |
| Human | 688.10 | 0.31 | 0.00 |
| Dwarf | 672.87 | 0.29 | 0.00 |
| Night Elf | 686.06 | 0.31 | 0.00 |
| Gnome | 689.94 | 0.32 | 0.00 |
| High Order | 681.11 | 0.30 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Darkspear Pauldrons](https://www.wowhead.com/forever/item=272105) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Agility / crit / hit — Waist (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Premier High Warlord's Blade](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Fiery Weapon |
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
| Auto-attack (tag 1) | 199.54 |
| [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | 150.51 |
| Auto-attack (tag 2) | 129.70 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 94.36 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 35.53 |
| Auto-attack (tag 3) | 28.82 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 27.15 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 17.74 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3176.1 |
| Energy | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | -2911.9 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -469.2 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -462.4 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +460.3 |
| Energy | OtherActionRefund | +104.8 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +96.5 |
| Energy | [Blade Flurry](https://www.wowhead.com/forever/spell=13877) | -75.0 |
| ComboPoints | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | +69.5 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -66.2 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -26.1 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +24.5 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 639.69 | 0.26 | 0.00 |
| Troll | 639.68 | 0.27 | 0.00 |
| Undead | 665.79 | 0.27 | 0.00 |
| Windshaper | 645.58 | 0.26 | 0.00 |
| Human | 639.86 | 0.27 | 0.00 |
| Dwarf | 637.18 | 0.25 | 0.00 |
| Night Elf | 639.63 | 0.25 | 0.00 |
| Gnome | 645.03 | 0.25 | 0.00 |
| High Order | 645.58 | 0.26 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Agility / crit / hit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Dawn Armor](https://www.wowhead.com/forever/item=252483) | 61 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit — Waist (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Cutthroat's Signet](https://www.wowhead.com/forever/item=272408) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / AP / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=20068) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand (2.5s) (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Modeled: Fist off-hand — Off hand weapon (1.7s) (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Agility |
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
| Auto-attack (tag 1) | 192.36 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 163.65 |
| Auto-attack (tag 2) | 96.08 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 44.69 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 39.69 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 28.99 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 23.04 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 21.24 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +2984.5 |
| Energy | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | -2829.9 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +538.5 |
| Energy | [Rupture](https://www.wowhead.com/forever/spell=11275) | -482.2 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -345.5 |
| Energy | OtherActionRefund | +153.2 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +98.9 |
| ComboPoints | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | +84.6 |
| Energy | [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) | -83.6 |
| Energy | [Ambush](https://www.wowhead.com/forever/spell=11269) | -82.8 |
| ComboPoints | [Rupture](https://www.wowhead.com/forever/spell=11275) | -62.8 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -39.6 |

## Druid — Balance

**Talents:** 38/0/13 · `5232220115501351--505003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 656.13 | 0.29 | 0.00 |
| Windshaper | 656.19 | 0.29 | 0.00 |
| Night Elf | 652.44 | 0.29 | 0.00 |
| High Order | 656.19 | 0.29 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Mantle of the Timbermaw](https://www.wowhead.com/forever/item=19050) | 61 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit — Hands (Leather) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 387.58 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 96.49 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 90.80 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 81.33 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=25298) | -15467.5 |
| Mana | OtherActionManaRegen (tag 1) | +7141.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4574.9 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -4265.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3602.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.5 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1554.2 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1231.2 |
| Mana | [Moonkin Form](https://www.wowhead.com/forever/spell=24858) | -435.4 |
| Mana | OtherActionManaRegen (tag 2) | +57.6 |

## Shaman — Stormcaller

**Talents:** 28/23/0 · `550032150010303-055030031004002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** The rank-2 Lightning Bolt filler depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 597.75 | 0.25 | 0.28 |
| Tauren | 598.27 | 0.25 | 0.28 |
| Troll | 597.62 | 0.25 | 0.28 |
| Windshaper | 598.59 | 0.25 | 0.32 |
| Dwarf | 594.05 | 0.25 | 0.27 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit — Hands (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 345.27 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 87.83 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 76.96 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 35.09 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 30.65 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 17.33 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 3.90 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.56 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -18970.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5250.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5243.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3607.7 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3144.2 |
| Mana | OtherActionManaRegen (tag 1) | +3123.8 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | -928.0 |
| Mana | OtherActionManaRegen (tag 2) | +57.8 |

## Priest — Smite

**Talents:** 31/17/3 · `515030031305001031-00505023002-003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** The rank-2 Smite fallback depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 591.86 | 0.15 | 0.00 |
| Undead | 596.10 | 0.16 | 0.00 |
| Human | 588.44 | 0.15 | 0.00 |
| Dwarf | 588.41 | 0.15 | 0.00 |
| Night Elf | 591.32 | 0.15 | 0.00 |
| Gnome | 591.86 | 0.15 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Argent Shoulders](https://www.wowhead.com/forever/item=19059) | 61 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / intellect / spell power / hit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit — Hands (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Bloodvine Leggings](https://www.wowhead.com/forever/item=19683) | 65 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Blessed Band of Light](https://www.wowhead.com/forever/item=272407) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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
| [Smite](https://www.wowhead.com/forever/spell=10934) | 221.97 |
| [Penance](https://www.wowhead.com/forever/spell=1316995) | 146.01 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 135.00 |
| [Smite](https://www.wowhead.com/forever/spell=591) | 83.31 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 9.81 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -14215.4 |
| Mana | OtherActionManaRegen (tag 1) | +8466.3 |
| Mana | [Penance](https://www.wowhead.com/forever/spell=1316995) | -8287.6 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -6282.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4244.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3616.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3599.3 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1574.2 |
| Mana | [Smite](https://www.wowhead.com/forever/spell=591) | -903.1 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |

## Shaman — Elemental

**Talents:** 31/7/13 · `5502301500123031-052-05305`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 544.64 | 0.23 | 0.17 |
| Tauren | 543.74 | 0.23 | 0.17 |
| Troll | 543.02 | 0.23 | 0.15 |
| Windshaper | 544.69 | 0.23 | 0.17 |
| Dwarf | 539.84 | 0.23 | 0.19 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Chains of the Lich](https://www.wowhead.com/forever/item=23125) | 60 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Bloodvine Vest](https://www.wowhead.com/forever/item=19682) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Mail) (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Gloves of the Greatfather](https://www.wowhead.com/forever/item=17721) | 38 | Enchant Gloves - Healing Power |
| Waist | [Belt of the Archmage](https://www.wowhead.com/forever/item=18405) | 62 | — |
| Legs | [Ironfeather Leggings](https://www.wowhead.com/forever/item=252486) | 61 | Arcanum of Focus |
| Feet | [Bloodvine Boots](https://www.wowhead.com/forever/item=19684) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Weakness Analyzer](https://www.wowhead.com/forever/item=272438) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 182.03 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 111.44 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | 99.58 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 80.09 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 33.74 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 22.06 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 9.09 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 5.57 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +6964.8 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -6886.2 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | -6429.4 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -6333.1 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | -5237.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4868.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3607.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3548.1 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -765.0 |
| Mana | OtherActionManaRegen (tag 2) | +73.4 |
