# Build reviews

These are hypothetical modeled-only gear results, not obtainable item rankings. Sources linked for modeled items are allocation references only and have different stats. Gear is compared under the [model assumptions](modelled_gear.md).

The tables and [matrix](../artifacts/modelled_gear/forever_dps_5min.png) use the same 201 common-seed replays. The modeled search and original real-item benchmark are separate; neither proves available launch gear. Equipment is frozen from an earlier mechanics revision; historical search controls and gains do not validate this corrected release.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and hit from selected gear, enchants, talents and racials only. [Scenario and exchange model](../tools/forever_bench/README.md) · [Questions and coverage](uncertainties.md)

Tank rows include frontal incoming attacks and modeled healing, with their recorded shared raid support except assigned self-maintained duties. They do not share the non-attacking DPS encounter. [Support and Judgement correction](tank_support.md). [Tank selection and guardrails](tank_selection.md).

## Paladin — Retribution

**Talents:** 13/7/31 · `55003-232-05225231001330301`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Separate seal Echoes can coexist and fire on a landed white swing; verify their simultaneous behavior in game (T53). Lower-rank seals and Consecration also need confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1103.06 | 0.56 | 0.00 |
| Human | 1096.50 | 0.55 | 0.00 |
| Dwarf | 1103.73 | 0.56 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Dwarf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Plate) [+18 Sta, +23 Int, +29 SP, +14 Crit, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Plate) [+14 Sta, +18 Int, +22 SP, +11 Crit, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Intellect / spell power / crit — Chest (Plate) [+18 Sta, +23 Int, +29 SP, +14 Crit, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / crit / hit — Hands (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / crit / hit — Legs (Plate) [+18 Sta, +21 Crit, +21 Str, +23 Hit, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Spell two-hand — Two-hand [+25 Sta, +24 Int, +24 SP, +15 Str, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Physical libram — Relic [+8 Sta, +4 Agi, +26 AP] (modeled; reference only)](https://www.wowhead.com/forever/item=279247) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -1.5s: [Spell 20919](https://www.wowhead.com/forever/spell=20919).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Holy Strike](https://www.wowhead.com/forever/spell=10333).
3. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271).
4. Cast [Spell 20919](https://www.wowhead.com/forever/spell=20919) when Seal time remaining ≤ 1s.
5. Cast [Spell 20293](https://www.wowhead.com/forever/spell=20293) when (Mana fraction ≥ 15% AND NOT [Echo of Command](https://www.wowhead.com/forever/spell=1311703) active AND [Spell 20919](https://www.wowhead.com/forever/spell=20919) active).
6. Cast [Spell 20919](https://www.wowhead.com/forever/spell=20919) when (Mana fraction ≥ 15% AND NOT [Echo of Righteousness](https://www.wowhead.com/forever/spell=1311704) active AND [Spell 20293](https://www.wowhead.com/forever/spell=20293) active).
7. Cast [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239).
8. Cast [Consecration](https://www.wowhead.com/forever/spell=26573) when Mana fraction ≥ 20%.

### Damage breakdown — Dwarf

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 326.01 |
| Auto-attack (tag 1) | 183.90 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 177.79 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 91.58 |
| Auto-attack (tag 3) | 89.91 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 63.71 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 62.34 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 56.21 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11247.4 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -9787.6 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -8873.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3412.9 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3179.4 |
| Mana | OtherActionManaRegen (tag 1) | +3111.2 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3104.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2634.7 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -1759.0 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1220.1 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -521.2 |
| Mana | OtherActionManaRegen (tag 2) | +1.3 |

## Paladin — Physical Ret

**Talents:** 13/7/31 · `52003003-232-05225031101330311`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Strength/AP equipment is unchanged. Champion of the Light converts existing Intellect into spell power without equipping caster gear; the separate seal Echoes still require the T53 in-game check.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 985.85 | 0.58 | 0.00 |
| Human | 983.63 | 0.57 | 0.00 |
| Dwarf | 974.79 | 0.57 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Axe two-hand — Two-hand [+34 Sta, +23 Str, +14 Crit, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272593) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Physical libram — Relic [+8 Sta, +4 Agi, +26 AP] (modeled; reference only)](https://www.wowhead.com/forever/item=279247) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -1.5s: [Spell 20919](https://www.wowhead.com/forever/spell=20919).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Holy Strike](https://www.wowhead.com/forever/spell=10333).
3. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271).
4. Cast [Spell 20919](https://www.wowhead.com/forever/spell=20919) when Seal time remaining ≤ 1s.
5. Cast [Spell 20293](https://www.wowhead.com/forever/spell=20293) when (Mana fraction ≥ 15% AND NOT [Echo of Command](https://www.wowhead.com/forever/spell=1311703) active AND [Spell 20919](https://www.wowhead.com/forever/spell=20919) active).
6. Cast [Spell 20919](https://www.wowhead.com/forever/spell=20919) when (Mana fraction ≥ 15% AND NOT [Echo of Righteousness](https://www.wowhead.com/forever/spell=1311704) active AND [Spell 20293](https://www.wowhead.com/forever/spell=20293) active).
7. Cast [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239).
8. Cast [Consecration](https://www.wowhead.com/forever/spell=26573) when Mana fraction ≥ 20%.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 325.06 |
| Auto-attack (tag 1) | 219.02 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 109.79 |
| Auto-attack (tag 3) | 106.17 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 84.01 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 37.26 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 33.87 |
| [Consecration](https://www.wowhead.com/forever/spell=26573) | 23.94 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11364.4 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -9944.0 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -9013.6 |
| Mana | OtherActionManaRegen (tag 1) | +5165.4 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3198.6 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3123.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2405.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +1957.0 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=26573) | -1577.4 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1321.4 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -520.7 |
| Mana | OtherActionManaRegen (tag 2) | +1.2 |

## Hunter — Beast Mastery

**Talents:** 31/20/0 · `5320001505101251-00531510005`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 943.51 | 0.27 | 1.08 |
| Tauren | 941.42 | 0.28 | 0.61 |
| Troll | 947.93 | 0.29 | 0.61 |
| Windshaper | 949.99 | 0.29 | 0.48 |
| Human | 945.54 | 0.27 | 0.74 |
| Dwarf | 943.03 | 0.28 | 0.51 |
| Night Elf | 954.84 | 0.29 | 0.46 |
| High Order | 949.99 | 0.29 | 0.48 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Night Elf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Agility / crit / hit — Neck [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Minor Agility |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 123–185 dmg @ 2.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

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
| Shoot | 348.70 |
| Cat: Auto-attack (tag 1) | 228.57 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | 123.04 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 75.55 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 70.36 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 53.43 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 45.08 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 8.17 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | -9021.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7329.2 |
| Mana | OtherActionManaRegen (tag 1) | +5859.4 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5373.1 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -4609.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3599.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3595.3 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -2998.8 |
| Mana | OtherActionManaRegen (tag 2) | +799.3 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Intimidation](https://www.wowhead.com/forever/spell=19577) | -310.8 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Hunter — Pet/Melee

**Talents:** 16/10/25 · `53200005001-005005-5302002300502201`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Hawk damage uses an approximate guardian/pet model. Tracking talents affect its comparison with Survival on different creature types.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 901.99 | 0.36 | 0.00 |
| Windshaper | 909.34 | 0.36 | 0.00 |
| Orc | 915.18 | 0.36 | 0.00 |
| Human | 917.26 | 0.36 | 0.00 |
| Night Elf | 911.65 | 0.36 | 0.00 |
| Dwarf | 899.55 | 0.36 | 0.00 |
| Troll | 905.50 | 0.35 | 0.00 |
| High Order | 909.34 | 0.36 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / crit / hit — Neck [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 123–185 dmg @ 2.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4s: [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447).
- -2s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.
- -1.5s: [Hunter's Mark](https://www.wowhead.com/forever/spell=14325).

### Rotation priorities

1. Cast [Hunter's Mark](https://www.wowhead.com/forever/spell=14325) when [Hunter's Mark](https://www.wowhead.com/forever/spell=14325) time remaining ≤ 5s.
2. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
3. Use ready automatic cooldowns.
4. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
5. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
6. Cast [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) when (NOT [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) DoT active AND Time remaining ≥ 9s).
7. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when Time remaining ≥ 12s.
8. Cast [Wing Clip](https://www.wowhead.com/forever/spell=14268) when Mana fraction ≥ 40%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 153.85 |
| Cat: Auto-attack (tag 1) | 143.33 |
| Auto-attack (tag 1) | 120.94 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 101.02 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 98.29 |
| Auto-attack (tag 3) | 77.71 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 65.89 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 56.61 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9659.6 |
| Mana | OtherActionManaRegen (tag 1) | +3890.6 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -3314.4 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3314.0 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -3029.7 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -1728.7 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -890.7 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -890.0 |
| Mana | [Hunter's Mark](https://www.wowhead.com/forever/spell=14325) | -180.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -100.0 |
| Health | [Holy Strength](https://www.wowhead.com/forever/spell=20007) | +0.0 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 892.32 | 0.38 | 0.00 |
| Tauren | 879.85 | 0.37 | 0.00 |
| Troll | 882.59 | 0.36 | 0.00 |
| Windshaper | 882.84 | 0.36 | 0.00 |
| Human | 893.58 | 0.38 | 0.00 |
| Dwarf | 876.92 | 0.37 | 0.00 |
| Night Elf | 888.48 | 0.38 | 0.00 |
| High Order | 882.84 | 0.36 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Agility / crit / hit — Neck [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 123–185 dmg @ 2.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4s: [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447).
- -2s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.
- -1.5s: [Hunter's Mark](https://www.wowhead.com/forever/spell=14325).

### Rotation priorities

1. Cast [Hunter's Mark](https://www.wowhead.com/forever/spell=14325) when [Hunter's Mark](https://www.wowhead.com/forever/spell=14325) time remaining ≤ 5s.
2. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
3. Use ready automatic cooldowns.
4. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
5. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
6. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when Time remaining ≥ 12s.
7. Cast [Wing Clip](https://www.wowhead.com/forever/spell=14268) when Mana fraction ≥ 40%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 160.69 |
| Cat: Auto-attack (tag 1) | 128.12 |
| Auto-attack (tag 1) | 126.16 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 101.10 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 99.26 |
| Auto-attack (tag 3) | 82.84 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 58.59 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 52.06 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +17065.2 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -9236.0 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4338.0 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -3661.0 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3397.8 |
| Mana | OtherActionManaRegen (tag 1) | +3156.3 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -2207.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +1996.1 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Hunter's Mark](https://www.wowhead.com/forever/spell=14325) | -180.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +98.8 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 877.56 | 0.41 | 0.00 |
| Troll | 872.71 | 0.40 | 0.00 |
| Undead | 892.48 | 0.41 | 0.00 |
| Human | 884.94 | 0.41 | 0.00 |
| Gnome | 875.50 | 0.40 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Spell power / crit / hit (matched cost) — Neck [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Shadow Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 322.54 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 94.78 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 88.73 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 87.38 |
| Succubus: Auto-attack (tag 1) | 78.54 |
| Succubus: [Demonic Brand](https://www.wowhead.com/forever/spell=1293697) | 70.30 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 54.08 |
| [Soul Fire](https://www.wowhead.com/forever/spell=17924) | 49.15 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -23779.0 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +11858.5 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +11538.7 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6515.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5383.1 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5144.7 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -4266.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3824.4 |
| Mana | [Soul Fire](https://www.wowhead.com/forever/spell=17924) | -3127.8 |
| Mana | OtherActionManaRegen (tag 1) | +3111.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2488.8 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |

## Warrior — Fury

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Level-60 rage, queued off-hand hit and Flurry charge timing still need in-game confirmation.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 833.09 | 0.42 | 0.00 |
| Tauren | 823.97 | 0.41 | 0.00 |
| Troll | 823.43 | 0.42 | 0.00 |
| Undead | 842.15 | 0.43 | 0.00 |
| Windshaper | 825.73 | 0.42 | 0.00 |
| Human | 832.28 | 0.42 | 0.00 |
| Dwarf | 821.28 | 0.42 | 0.00 |
| Night Elf | 822.26 | 0.41 | 0.00 |
| Gnome | 822.34 | 0.42 | 0.00 |
| High Order | 825.73 | 0.42 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / crit / hit — Neck [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / crit / hit — Hands (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

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
| Auto-attack (tag 2) | 149.38 |
| Auto-attack (tag 1) | 138.59 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 120.55 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 94.91 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 94.42 |
| Auto-attack (tag 3) | 68.46 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 48.45 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 40.52 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +1195.3 |
| Rage | Auto-attack (tag 1) | +1008.9 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -937.1 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -684.9 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -505.6 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -503.8 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +182.0 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +179.7 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -145.0 |
| Rage | OtherActionRefund | +103.0 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +100.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +98.6 |

## Priest — Shadow

**Talents:** 18/0/33 · `305030001303--504020501201302251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 826.56 | 0.25 | 0.00 |
| Undead | 832.50 | 0.25 | 0.00 |
| Human | 825.19 | 0.25 | 0.00 |
| Dwarf | 831.54 | 0.25 | 0.00 |
| Night Elf | 831.50 | 0.25 | 0.00 |
| Gnome | 826.87 | 0.25 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Crit / shadow power / intellect — Head (Cloth) [+18 Sta, +26 Shadow, +24 Crit, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Arcanum of Focus |
| Neck | [Modeled: Crit / shadow power / intellect — Neck [+11 Sta, +15 Shadow, +14 Crit, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Shoulders | [Modeled: Crit / shadow power / intellect — Shoulders (Cloth) [+14 Sta, +20 Shadow, +18 Crit, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Crit / shadow power / intellect — Back (Cloth) [+11 Sta, +15 Shadow, +14 Crit, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Crit / shadow power / intellect — Chest (Cloth) [+18 Sta, +26 Shadow, +24 Crit, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Crit / shadow power / intellect — Wrists (Cloth) [+11 Sta, +15 Shadow, +14 Crit, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Crit / shadow power / intellect — Hands (Cloth) [+14 Sta, +20 Shadow, +18 Crit, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Gloves - Shadow Power |
| Waist | [Modeled: Crit / shadow power / intellect — Waist (Cloth) [+14 Sta, +20 Shadow, +18 Crit, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Legs | [Modeled: Spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +21 Crit, +23 Hit, +23 SP, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Focus |
| Feet | [Modeled: Spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Mace one-hand — One-hand [+8 Agi, +12 Sta, +11 Hit, 76–142 dmg @ 2.3s] (modeled; reference only)](https://www.wowhead.com/forever/item=279261) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 324.81 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 159.63 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 147.27 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 99.23 |
| [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | 91.82 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 9.74 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +8072.1 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -7103.6 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -5950.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4371.1 |
| Mana | [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | -3395.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2934.1 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2849.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +1780.9 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1399.7 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2525000013520105--0550015103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 801.77 | 0.47 | 0.00 |
| Troll | 790.87 | 0.45 | 0.00 |
| Undead | 802.52 | 0.46 | 0.00 |
| Human | 808.52 | 0.46 | 0.00 |
| Gnome | 790.64 | 0.46 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Crit / shadow power / intellect — Neck [+11 Sta, +15 Shadow, +14 Crit, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Crit / shadow power / intellect — Chest (Cloth) [+18 Sta, +26 Shadow, +24 Crit, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Shadow Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit — Legs (Cloth) [+18 Sta, +23 Int, +29 SP, +14 Crit, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 366.14 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 109.97 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 99.23 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 88.31 |
| Succubus: Auto-attack (tag 1) | 70.58 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 41.95 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 19.22 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 8.02 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -28518.7 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24128.4 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6765.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5353.6 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5231.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3603.0 |
| Mana | [Siphon Life](https://www.wowhead.com/forever/spell=18881) | -3602.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3266.1 |
| Mana | OtherActionManaRegen (tag 1) | +3090.5 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -459.7 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -338.9 |

## Warrior — Fury (Sunder)

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 801.98 | 0.43 | 0.00 |
| Tauren | 794.98 | 0.42 | 0.00 |
| Troll | 795.81 | 0.42 | 0.00 |
| Undead | 807.67 | 0.41 | 0.00 |
| Windshaper | 791.50 | 0.40 | 0.00 |
| Human | 799.29 | 0.42 | 0.00 |
| Dwarf | 783.38 | 0.42 | 0.00 |
| Night Elf | 792.66 | 0.42 | 0.00 |
| Gnome | 789.66 | 0.42 | 0.00 |
| High Order | 791.50 | 0.40 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / crit / hit — Head (Plate) [+18 Sta, +21 Crit, +21 Str, +23 Hit, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / crit / hit — Neck [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Strength / crit / hit — Shoulders (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / crit / hit — Chest (Plate) [+18 Sta, +21 Crit, +21 Str, +23 Hit, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / crit / hit — Hands (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Rapidity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) when ([Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) stacks < 5 OR [Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) time remaining < 8s).
2. `{"condition":{"and":{"vals":[{"or":{"vals":[{"cmp":{"op":"OpLt","lhs":{"auraNumStacks":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597,"rank":5}}},"rhs":{"const":{"val":"5"}}}},{"cmp":{"op":"OpLt","lhs":{"auraRemainingTime":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597,"rank":5}}},"rhs":{"const":{"val":"8s"}}}}]}},{"cmp":{"op":"OpLt","lhs":{"currentRage":{}},"rhs":{"const":{"val":"15"}}}}]}},"wait":{"duration":{"const":{"val":"500ms"}}}}` when (([Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) stacks < 5 OR [Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) time remaining < 8s) AND Rage < 15).
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
| Auto-attack (tag 2) | 145.42 |
| Auto-attack (tag 1) | 144.94 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 103.53 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 90.39 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 84.92 |
| Auto-attack (tag 3) | 72.18 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 46.31 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 37.18 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +1202.9 |
| Rage | Auto-attack (tag 1) | +1100.8 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -866.8 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -668.0 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -474.1 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -457.2 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -273.3 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +188.6 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +180.2 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -135.1 |
| Rage | OtherActionRefund | +110.2 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +100.0 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 792.74 | 0.33 | 0.00 |
| Troll | 796.32 | 0.33 | 0.00 |
| Undead | 800.21 | 0.33 | 0.00 |
| Human | 799.99 | 0.33 | 0.00 |
| Gnome | 786.02 | 0.32 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Fire Power |
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit — Legs (Cloth) [+18 Sta, +23 Int, +29 SP, +14 Crit, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Intellect / spell power / crit — Trinket [+11 Sta, +14 Int, +18 SP, +9 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 247.55 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 118.66 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 94.49 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 87.27 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 81.98 |
| Succubus: Auto-attack (tag 1) | 74.94 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 51.89 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 15.90 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +21874.7 |
| Mana | [Incinerate](https://www.wowhead.com/forever/spell=1293813) | -19402.5 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6701.5 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -6347.4 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -5644.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5367.6 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5114.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4358.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.1 |
| Mana | OtherActionManaRegen (tag 1) | +3099.7 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -476.7 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-550500000002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 784.45 | 0.50 | 0.00 |
| Tauren | 777.77 | 0.49 | 0.00 |
| Troll | 777.28 | 0.50 | 0.00 |
| Undead | 785.26 | 0.51 | 0.00 |
| Windshaper | 781.70 | 0.51 | 0.00 |
| Human | 785.67 | 0.51 | 0.00 |
| Dwarf | 776.53 | 0.49 | 0.00 |
| Night Elf | 782.69 | 0.50 | 0.00 |
| Gnome | 776.33 | 0.51 | 0.00 |
| High Order | 781.70 | 0.51 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Sword two-hand — Two-hand [+34 Sta, +23 Str, +14 Crit, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272604) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

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
| Auto-attack (tag 1) | 170.71 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 157.79 |
| Auto-attack (tag 3) | 126.32 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 116.07 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 69.77 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 38.61 |
| [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | 36.84 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 28.18 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1751.5 |
| Rage | [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | -1025.1 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -495.4 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -179.8 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -171.6 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -166.2 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +157.4 |
| Rage | [Rend](https://www.wowhead.com/forever/spell=11574) | -145.5 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +122.4 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.8 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -98.1 |
| Rage | OtherActionRefund | +92.6 |

## Hunter — Marksmanship

**Talents:** 5/35/11 · `5-0050552011523051-50005001`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 774.71 | 0.37 | 5.55 |
| Tauren | 760.60 | 0.36 | 5.31 |
| Troll | 781.93 | 0.38 | 4.68 |
| Windshaper | 781.23 | 0.39 | 4.55 |
| Human | 778.40 | 0.37 | 3.06 |
| Dwarf | 761.58 | 0.36 | 3.54 |
| Night Elf | 772.88 | 0.37 | 6.27 |
| High Order | 781.23 | 0.39 | 4.55 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Troll

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Mail) [+14 Sta, +18 Int, +22 SP, +11 Crit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Greater Intellect |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 123–185 dmg @ 2.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

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
| Shoot | 394.46 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 183.28 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 79.96 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 66.91 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 47.35 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 9.97 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -11214.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7193.0 |
| Mana | OtherActionManaRegen (tag 1) | +5722.8 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -5224.3 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3917.2 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -3738.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3631.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3603.7 |
| Mana | OtherActionManaRegen (tag 2) | +741.1 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `252200001351-0025003001-0550005103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 762.40 | 0.50 | 0.00 |
| Troll | 752.78 | 0.49 | 0.00 |
| Undead | 776.15 | 0.50 | 0.00 |
| Human | 777.54 | 0.51 | 0.00 |
| Gnome | 763.77 | 0.50 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Shadow Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 435.40 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 121.12 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 105.70 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 87.33 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 21.84 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 3.76 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.38 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -30648.7 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +21709.8 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6838.1 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5492.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5400.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3604.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3503.4 |
| Mana | OtherActionManaRegen (tag 1) | +3111.6 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -625.6 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -360.9 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -206.6 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `550133102-053030031005102251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 763.53 | 0.57 | 0.04 |
| Tauren | 759.73 | 0.56 | 0.01 |
| Troll | 763.74 | 0.56 | 0.00 |
| Windshaper | 762.72 | 0.55 | 0.01 |
| Dwarf | 768.28 | 0.56 | 0.01 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Dwarf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / intellect / crit / hit (mail) — Shoulders (Mail) [+11 Str, +17 Int, +17 Sta, +11 Crit, +14 Hit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Intellect / spell power / crit — Chest (Mail) [+18 Sta, +23 Int, +29 SP, +14 Crit, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / intellect / crit / hit (mail) — Wrists (Mail) [+8 Str, +13 Int, +13 Sta, +8 Crit, +10 Hit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / intellect / crit / hit (mail) — Hands (Mail) [+11 Str, +17 Int, +17 Sta, +11 Crit, +14 Hit, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / intellect / crit / hit (mail) — Waist (Mail) [+11 Str, +17 Int, +17 Sta, +11 Crit, +14 Hit, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | — |
| Legs | [Modeled: Strength / intellect / crit / hit (mail) — Legs (Mail) [+14 Str, +22 Int, +22 Sta, +14 Crit, +18 Hit, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / intellect / crit / hit (mail) — Feet (Mail) [+11 Str, +17 Int, +17 Sta, +11 Crit, +14 Hit, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Spell two-hand — Two-hand [+25 Sta, +24 Int, +24 SP, +15 Str, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Fiery Weapon |
| Ranged/relic | [Modeled: Physical totem — Relic [+8 Sta, +4 Agi, +26 AP] (modeled; reference only)](https://www.wowhead.com/forever/item=279249) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4.5s: [Strength of Earth Totem (rank 4)](https://www.wowhead.com/forever/spell=10442).
- -3s: [Spell 10614 (rank 3)](https://www.wowhead.com/forever/spell=10614).
- -1.5s: [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438).

### Rotation priorities

1. Cast [Strength of Earth Totem (rank 4)](https://www.wowhead.com/forever/spell=10442) when Earth totem time remaining ≤ 0s.
2. Cast [Spell 10614 (rank 3)](https://www.wowhead.com/forever/spell=10614) when Air totem time remaining ≤ 0s.
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
| Auto-attack (tag 1) | 224.90 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 143.71 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 91.71 |
| [Fire Nova](https://www.wowhead.com/forever/spell=408345) | 90.97 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 62.39 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 49.82 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 44.88 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 28.72 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Nova](https://www.wowhead.com/forever/spell=408345) | -14399.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9301.3 |
| Mana | OtherActionManaRegen (tag 1) | +6774.4 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -5394.4 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4735.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4590.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3598.5 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3312.7 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | [Spell 10614](https://www.wowhead.com/forever/spell=10614) | -499.6 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -450.0 |
| Mana | OtherActionManaRegen (tag 2) | +204.7 |

## Mage — Arcane–Frost

**Talents:** 28/0/23 · `05020500310031053--05550002010003002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Uses an assumed Ice Lance coefficient and unresolved Fingers of Frost, Missile Barrage and Clearcasting timing.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 742.42 | 0.39 | 0.00 |
| Troll | 740.60 | 0.39 | 0.00 |
| Undead | 755.54 | 0.40 | 0.00 |
| Human | 751.64 | 0.39 | 0.00 |
| Gnome | 747.48 | 0.39 | 0.00 |
| High Order | 748.75 | 0.39 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Frost Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 441.63 |
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 162.29 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 134.26 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 17.35 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -21659.4 |
| Mana | OtherActionManaRegen (tag 1) | +10763.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5664.7 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -2951.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2648.5 |
| Mana | OtherActionManaRegen (tag 2) | +1719.2 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1061.1 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +692.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +336.7 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +216.8 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Warrior — 2H Bloodthirst

**Talents:** 20/31/0 · `30304203032-050500320051310051`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 748.69 | 0.46 | 0.00 |
| Tauren | 734.95 | 0.46 | 0.00 |
| Undead | 745.41 | 0.47 | 0.00 |
| Troll | 741.32 | 0.46 | 0.00 |
| Windshaper | 734.56 | 0.46 | 0.00 |
| Human | 744.64 | 0.46 | 0.00 |
| Dwarf | 739.47 | 0.46 | 0.00 |
| Gnome | 737.11 | 0.46 | 0.00 |
| Night Elf | 738.11 | 0.46 | 0.00 |
| High Order | 734.56 | 0.46 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / agility / crit / hit — Feet (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Axe two-hand — Two-hand [+34 Sta, +23 Str, +14 Crit, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272593) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

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
| Auto-attack (tag 1) | 236.08 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 109.01 |
| Auto-attack (tag 3) | 95.74 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 71.27 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 63.94 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 50.52 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 50.12 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 33.61 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1854.6 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -984.5 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -537.7 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -500.4 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +174.9 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +130.2 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -126.5 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -123.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +92.6 |
| Rage | OtherActionRefund | +77.4 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -55.3 |
| Rage | [Berserker Stance](https://www.wowhead.com/forever/spell=2458) | -28.0 |

## Mage — Arcane

**Talents:** 33/3/15 · `053005023100311531-03-005500032`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 719.62 | 0.54 | 0.08 |
| Troll | 720.22 | 0.53 | 0.04 |
| Undead | 739.21 | 0.55 | 0.04 |
| Human | 724.84 | 0.54 | 0.06 |
| Gnome | 732.17 | 0.53 | 0.02 |
| High Order | 725.55 | 0.54 | 0.06 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Intellect / spell power / crit — Waist (Cloth) [+14 Sta, +18 Int, +22 SP, +11 Crit, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit — Legs (Cloth) [+18 Sta, +23 Int, +29 SP, +14 Crit, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 348.16 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 324.86 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 44.97 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 21.24 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -30465.3 |
| Mana | OtherActionManaRegen (tag 1) | +10945.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +6956.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4436.7 |
| Mana | OtherActionManaRegen (tag 2) | +3879.7 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -1712.3 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.8 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.9 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | -640.0 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +558.3 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Fire

**Talents:** 17/31/3 · `0501252000002-23450000130133051-003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 721.24 | 0.50 | 5.89 |
| Troll | 720.33 | 0.46 | 4.33 |
| Undead | 734.51 | 0.48 | 5.48 |
| Human | 731.62 | 0.49 | 5.19 |
| Gnome | 725.90 | 0.48 | 5.06 |
| High Order | 720.93 | 0.44 | 2.79 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / fire power / crit — Head (Cloth) [+18 Sta, +24 Int, +26 Fire, +18 Crit, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / fire power / crit — Neck [+11 Sta, +14 Int, +15 Fire, +10 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Cloth) [+14 Sta, +18 Int, +22 SP, +11 Crit, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / fire power / crit — Back (Cloth) [+11 Sta, +14 Int, +15 Fire, +10 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / fire power / crit — Chest (Cloth) [+18 Sta, +24 Int, +26 Fire, +18 Crit, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / fire power / crit — Wrists (Cloth) [+11 Sta, +14 Int, +15 Fire, +10 Crit, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Fire Power |
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +21 Crit, +23 Hit, +23 SP, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Focus |
| Feet | [Modeled: Spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 80%.
2. Use ready automatic cooldowns.
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 20% AND Time remaining > 25s).
4. Cast [Pyroblast (rank 8)](https://www.wowhead.com/forever/spell=18809) when [Hot Streak](https://www.wowhead.com/forever/spell=44445) stacks = 3.
5. Cast [Fire Blast (rank 7)](https://www.wowhead.com/forever/spell=10199) when Mana ≥ Time remaining × 10.
6. Cast [Scorch (rank 7)](https://www.wowhead.com/forever/spell=10207) when ([Improved Scorch](https://www.wowhead.com/forever/spell=12873) stacks < 5 OR [Improved Scorch](https://www.wowhead.com/forever/spell=12873) time remaining ≤ 5s).
7. Cast [Fireball](https://www.wowhead.com/forever/spell=25306) when Mana ≥ Time remaining × 150.
8. Cast [Scorch](https://www.wowhead.com/forever/spell=10207).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 271.66 |
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 173.33 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 159.06 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 101.29 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 17.06 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 8.63 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 3.48 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -16009.4 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -13738.8 |
| Mana | OtherActionManaRegen (tag 1) | +13154.9 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -7888.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5435.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4673.6 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +4394.4 |
| Mana | OtherActionManaRegen (tag 2) | +1778.7 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.9 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.6 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +600.7 |
| Mana | [Fireball](https://www.wowhead.com/forever/spell=25306) | -276.8 |

## Mage — Frost

**Talents:** 11/3/37 · `050005001-03-0555003321001301251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 703.56 | 0.43 | 0.26 |
| Troll | 700.98 | 0.42 | 0.19 |
| Undead | 704.68 | 0.42 | 0.36 |
| Human | 707.89 | 0.42 | 0.15 |
| Gnome | 713.35 | 0.44 | 0.02 |
| High Order | 702.80 | 0.42 | 0.20 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Cloth) [+18 Sta, +23 Int, +29 SP, +14 Crit, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Frost Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 518.63 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 194.73 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -23398.8 |
| Mana | OtherActionManaRegen (tag 1) | +9432.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3975.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3714.1 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3581.4 |
| Mana | OtherActionManaRegen (tag 2) | +2033.1 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.2 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +850.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +846.2 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +176.8 |

## Druid — Balance

**Talents:** 38/0/13 · `5232220115501351--505003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 708.05 | 0.33 | 0.00 |
| Windshaper | 711.90 | 0.33 | 0.00 |
| Night Elf | 711.64 | 0.33 | 0.00 |
| High Order | 711.90 | 0.33 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Leather) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Leather) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 148 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Leather) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 197 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Leather) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 86 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Leather) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 123 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Leather) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 111 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Leather) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 173 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Leather) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 136 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power idol — Relic [+8 Sta, +7 Int, +9 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279250) | 65 | — |

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
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 413.33 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 106.39 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 97.47 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 94.70 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=25298) | -15277.7 |
| Mana | OtherActionManaRegen (tag 1) | +7151.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4583.6 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -4347.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3596.1 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1745.4 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1227.0 |
| Mana | [Moonkin Form](https://www.wowhead.com/forever/spell=24858) | -435.4 |
| Mana | OtherActionManaRegen (tag 2) | +42.7 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 687.08 | 0.29 | 0.00 |
| Troll | 685.16 | 0.29 | 0.00 |
| Undead | 708.42 | 0.29 | 0.00 |
| Windshaper | 688.38 | 0.29 | 0.00 |
| Human | 690.02 | 0.28 | 0.00 |
| Dwarf | 680.28 | 0.29 | 0.00 |
| Night Elf | 688.65 | 0.29 | 0.00 |
| Gnome | 686.09 | 0.29 | 0.00 |
| High Order | 688.38 | 0.29 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit (leather) — Head (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Agility / crit / hit — Neck [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit (leather) — Shoulders (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 148 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 197 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Leather) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 86 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 123 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Agility / crit / hit — Waist (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 111 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 173 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Feet | [Modeled: Agility / crit / hit — Feet (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 136 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Fist main-hand — Main hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272597) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -1s: [Blade Flurry](https://www.wowhead.com/forever/spell=13877).

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) time remaining < 3 AND Time remaining > 9)).
3. Cast [Adrenaline Rush](https://www.wowhead.com/forever/spell=13750) when Energy ≤ 40.
4. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
5. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when Combo points ≥ 3.
6. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 3.
7. Cast [Sinister Strike (rank 8)](https://www.wowhead.com/forever/spell=11294) when (`{"autoTimeToNext":{"autoType":"MainHand"}}` > `{"autoSwingTime":{"autoType":"MainHand"}}` − 0.5 OR Time remaining < 6 OR Energy ≥ 79).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 178.58 |
| [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | 127.93 |
| Auto-attack (tag 2) | 114.23 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 85.69 |
| Auto-attack (tag 3) | 77.78 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 38.54 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 29.32 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 23.13 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3176.2 |
| Energy | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | -2914.8 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -468.9 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -462.7 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +462.3 |
| Energy | OtherActionRefund | +105.2 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +97.1 |
| Energy | [Blade Flurry](https://www.wowhead.com/forever/spell=13877) | -75.0 |
| ComboPoints | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | +69.6 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -66.4 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -26.0 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +24.6 |

## Rogue — Mutilate

**Talents:** 31/20/0 · `0053031035140105-302303202014`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 672.76 | 0.27 | 0.00 |
| Troll | 669.61 | 0.27 | 0.00 |
| Undead | 703.89 | 0.29 | 0.00 |
| Windshaper | 674.37 | 0.27 | 0.00 |
| Human | 667.16 | 0.27 | 0.00 |
| Dwarf | 667.50 | 0.27 | 0.00 |
| Night Elf | 675.61 | 0.27 | 0.00 |
| Gnome | 675.38 | 0.27 | 0.00 |
| High Order | 674.37 | 0.27 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Agility / crit / hit — Neck [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 148 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 197 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Leather) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 86 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 123 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Agility / crit / hit — Waist (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 111 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 173 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Feet | [Modeled: Agility / crit / hit — Feet (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 136 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Modeled: Dagger off-hand (unverified) — Off hand weapon [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Fiery Weapon |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) time remaining < 3 AND Time remaining > 9)).
3. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
4. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when (Combo points ≥ 4 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active).
5. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 3.
6. Cast [Mutilate](https://www.wowhead.com/forever/spell=1241584).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 156.40 |
| Auto-attack (tag 2) | 97.72 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 74.25 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 71.02 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 51.50 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 50.75 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 39.93 |
| Auto-attack (tag 3) | 38.83 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3013.6 |
| Energy | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | -2877.6 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -832.7 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +782.4 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -409.0 |
| Energy | OtherActionRefund | +151.5 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -136.6 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.6 |
| ComboPoints | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | +89.6 |
| ComboPoints | [Seal Fate](https://www.wowhead.com/forever/spell=14195) | +36.5 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +33.2 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -21.0 |

## Druid — Feral

**Talents:** 9/34/8 · `050022-3521002023032213041-053`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** Blizzard supports proc-enchant form eligibility, and Crusader's sourced health return is modeled. Its exact PPM/conditional interactions remain qualified (DRU-011); the Night Elf profile uses that enchant.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 689.22 | 0.23 | 0.00 |
| Windshaper | 694.02 | 0.23 | 0.00 |
| Night Elf | 695.36 | 0.25 | 0.00 |
| High Order | 694.02 | 0.23 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Night Elf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit (leather) — Head (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit (leather) — Shoulders (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 148 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit (leather) — Chest (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 197 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit (leather) — Wrists (Leather) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 86 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit (leather) — Hands (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 123 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit (leather) — Waist (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 111 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit (leather) — Legs (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 173 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit (leather) — Feet (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 136 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Mace two-hand — Two-hand [+34 Sta, +23 Str, +14 Crit, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272601) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Physical idol — Relic [+8 Sta, +4 Agi, +26 AP] (modeled; reference only)](https://www.wowhead.com/forever/item=279250) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Use ready automatic cooldowns.
3. Cast [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) when Energy ≤ 40.
4. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Time remaining ≤ 4s AND Combo points ≥ 3).
5. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Combo points ≥ 5 AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 5s).
6. Cast [Rip](https://www.wowhead.com/forever/spell=9896) when (Combo points ≥ 4 AND NOT [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND Time remaining ≥ 8s).
7. Cast [Shred](https://www.wowhead.com/forever/spell=9830).

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| [Shred](https://www.wowhead.com/forever/spell=9830) | 239.06 |
| Auto-attack (tag 1) | 212.45 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 172.53 |
| Auto-attack (tag 3) | 47.92 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 20.21 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 3.18 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -3241.2 |
| Energy | OtherActionEnergyRegen | +3017.2 |
| Energy | [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) | +686.8 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -654.3 |
| Energy | OtherActionRefund | +167.6 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -139.4 |
| ComboPoints | [Rip](https://www.wowhead.com/forever/spell=9896) | -105.6 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.5 |
| ComboPoints | [Shred](https://www.wowhead.com/forever/spell=9830) | +76.5 |
| ComboPoints | [Blood Frenzy](https://www.wowhead.com/forever/spell=37117) | +46.5 |
| ComboPoints | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -16.3 |
| Mana | OtherActionManaRegen (tag 2) | +0.0 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 632.64 | 0.25 | 0.00 |
| Troll | 632.82 | 0.26 | 0.00 |
| Undead | 655.85 | 0.27 | 0.00 |
| Windshaper | 635.11 | 0.26 | 0.00 |
| Human | 636.22 | 0.25 | 0.00 |
| Dwarf | 627.50 | 0.25 | 0.00 |
| Night Elf | 636.62 | 0.26 | 0.00 |
| Gnome | 635.37 | 0.26 | 0.00 |
| High Order | 635.11 | 0.26 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Agility / crit / hit — Neck [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit (leather) — Shoulders (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 148 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 197 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Leather) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 86 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit — Hands (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 123 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit — Waist (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 111 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 173 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 136 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 82–154 dmg @ 2.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Slice and Dice](https://www.wowhead.com/forever/spell=6774) when (Combo points ≥ 1 AND NOT [Slice and Dice](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6s).
2. Cast [Vanish](https://www.wowhead.com/forever/spell=1856) when (NOT [Stealth](https://www.wowhead.com/forever/spell=1787) active AND Energy ≥ 60 AND Combo points ≤ 1 AND [Slice and Dice](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 10s).
3. Cast [Premeditation](https://www.wowhead.com/forever/spell=14183) when Combo points ≤ 2.
4. Cast [Ambush](https://www.wowhead.com/forever/spell=11269) when Combo points ≤ 3.
5. Cast [Slice and Dice](https://www.wowhead.com/forever/spell=6774) when (Combo points ≥ 5 AND [Slice and Dice](https://www.wowhead.com/forever/spell=6774) time remaining ≤ 2s AND Time remaining ≥ 6s).
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
| Auto-attack (tag 1) | 161.10 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 135.95 |
| Auto-attack (tag 2) | 86.73 |
| Auto-attack (tag 3) | 55.39 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 41.71 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 40.76 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 29.58 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 29.49 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +2984.6 |
| Energy | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | -2827.2 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +538.2 |
| Energy | [Rupture](https://www.wowhead.com/forever/spell=11275) | -482.0 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -345.3 |
| Energy | OtherActionRefund | +150.6 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +98.9 |
| ComboPoints | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | +84.6 |
| Energy | [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) | -83.8 |
| Energy | [Ambush](https://www.wowhead.com/forever/spell=11269) | -82.5 |
| ComboPoints | [Rupture](https://www.wowhead.com/forever/spell=11275) | -62.8 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -39.7 |

## Druid — Bear · Tank

**Talents:** 1/40/10 · `01-5002232123132210551-055`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Bear retains unverified damage-based rage, not Warrior's normalized formula (DRU-012); threat coefficients also remain provisional. This is not a survival ranking. Proc-enchant form eligibility is supported; exact PPM/conditional interactions remain qualified (DRU-011). Historical enchant-selection gains do not validate the current model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 606.59 | 0.37 | 0.00 |
| Windshaper | 610.07 | 0.38 | 0.00 |
| Night Elf | 611.92 | 0.37 | 0.00 |
| High Order | 610.07 | 0.38 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Tauren | 1616.20 | 716.64 | 72.81 | 0.14% |
| Windshaper | 1617.41 | 715.09 | 77.53 | 0.36% |
| Night Elf | 1626.85 | 702.95 | 77.19 | 0.24% |
| High Order | 1617.41 | 715.09 | 77.53 | 0.36% |

These stress-scenario results are not measured boss balance or an equal-support survival ranking.

### Equipment — Night Elf

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit (leather) — Head (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit (leather) — Shoulders (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 148 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / agility / crit / hit — Back (Cloth) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit (leather) — Chest (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 197 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit (leather) — Wrists (Leather) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 86 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Agility / crit / hit (leather) — Hands (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 123 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit (leather) — Waist (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 111 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit (leather) — Legs (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 173 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit (leather) — Feet (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 136 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / agility / crit / hit — Finger [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Strength / agility / crit / hit — Finger [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Mace two-hand — Two-hand [+34 Sta, +23 Str, +14 Crit, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272601) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Physical idol — Relic [+8 Sta, +4 Agi, +26 AP] (modeled; reference only)](https://www.wowhead.com/forever/item=279250) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Frenzied Regeneration](https://www.wowhead.com/forever/spell=22842) when `{"currentHealthPercent":{}}` < 30%.
2. Cast [Barkskin](https://www.wowhead.com/forever/spell=22812).
3. Use ready automatic cooldowns.
4. Cast [Demoralizing Roar](https://www.wowhead.com/forever/spell=9898) when ([Demoralizing Roar](https://www.wowhead.com/forever/spell=9898) time remaining ≤ 4s AND Time remaining ≥ 5s).
5. Cast [Faerie Fire](https://www.wowhead.com/forever/spell=9907) when ([Faerie Fire](https://www.wowhead.com/forever/spell=9907) time remaining ≤ 3s AND Time remaining ≥ 4s).
6. Cast [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073).
7. Cast [Swipe](https://www.wowhead.com/forever/spell=9908) when (Target count ≥ 2 AND Rage ≥ 15).
8. Cast [Lacerate](https://www.wowhead.com/forever/spell=414644) when ([Lacerate](https://www.wowhead.com/forever/spell=414647) stacks < 5 OR [Lacerate](https://www.wowhead.com/forever/spell=414647) DoT time remaining ≤ 4.5s).
9. Cast [Maul (rank 7)](https://www.wowhead.com/forever/spell=9881) when Rage ≥ 25.
10. Cast [Lacerate](https://www.wowhead.com/forever/spell=414644).

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| [Maul](https://www.wowhead.com/forever/spell=9881) | 345.41 |
| [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073) | 119.83 |
| [Lacerate](https://www.wowhead.com/forever/spell=414644) | 96.86 |
| [Lacerate](https://www.wowhead.com/forever/spell=414647) | 34.01 |
| Auto-attack (tag 1) | 7.39 |
| Auto-attack (tag 3) | 6.11 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.31 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -212180.0 |
| Health | OtherActionHealingModel | +209914.9 |
| Rage | OtherActionDamageTaken | +2075.6 |
| Rage | [Lacerate](https://www.wowhead.com/forever/spell=414644) | -1387.7 |
| Rage | [Maul](https://www.wowhead.com/forever/spell=9881) | -1265.9 |
| Rage | [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073) | -841.8 |
| Rage | [Blood Frenzy](https://www.wowhead.com/forever/spell=37117) | +586.6 |
| Health | [Holy Strength](https://www.wowhead.com/forever/spell=20007) | +524.2 |
| Rage | OtherActionRefund | +505.9 |
| Health | [Frenzied Regeneration](https://www.wowhead.com/forever/spell=22842) | +393.7 |
| Rage | [Natural Reaction](https://www.wowhead.com/forever/spell=57878) | +210.7 |
| Rage | Auto-attack (tag 1) | +154.0 |

## Priest — Smite

**Talents:** 31/17/3 · `515030031305001031-00505023002-003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** The rank-2 Smite fallback depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 546.46 | 0.14 | 0.00 |
| Undead | 552.94 | 0.15 | 0.00 |
| Human | 543.47 | 0.15 | 0.00 |
| Dwarf | 543.18 | 0.15 | 0.00 |
| Night Elf | 546.44 | 0.15 | 0.00 |
| Gnome | 547.54 | 0.15 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

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
| [Smite](https://www.wowhead.com/forever/spell=10934) | 207.58 |
| [Penance](https://www.wowhead.com/forever/spell=1316995) | 136.55 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 125.17 |
| [Smite](https://www.wowhead.com/forever/spell=591) | 72.69 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 10.94 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -14430.4 |
| Mana | [Penance](https://www.wowhead.com/forever/spell=1316995) | -8299.8 |
| Mana | OtherActionManaRegen (tag 1) | +8174.2 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -6288.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4221.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3630.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.8 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1580.6 |
| Mana | [Smite](https://www.wowhead.com/forever/spell=591) | -869.9 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |

## Shaman — Stormcaller

**Talents:** 29/22/0 · `550133130010304-054030031004002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)

**Model limitations:** The rank-2 Lightning Bolt filler depends on the modeled low-rank spell-power coefficient.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 549.36 | 0.21 | 0.27 |
| Tauren | 548.34 | 0.21 | 0.31 |
| Troll | 546.44 | 0.21 | 0.27 |
| Windshaper | 547.36 | 0.21 | 0.36 |
| Dwarf | 544.77 | 0.21 | 0.28 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Mail) [+18 Sta, +23 Int, +29 SP, +14 Crit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Mail) [+14 Sta, +18 Int, +22 SP, +11 Crit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Spell power / crit / hit (matched cost) — Wrists (Mail) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Spell power / crit / hit (matched cost) — Hands (Mail) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Mail) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Spell power / crit / hit (matched cost) — Legs (Mail) [+18 Sta, +21 Crit, +23 Hit, +23 SP, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Focus |
| Feet | [Modeled: Spell power / crit / hit (matched cost) — Feet (Mail) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power totem — Relic [+8 Sta, +7 Int, +9 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279249) | 65 | — |

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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 279.94 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 93.59 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 91.05 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 34.92 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 29.84 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 13.90 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 4.63 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.48 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -16471.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5316.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5286.3 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3907.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.3 |
| Mana | OtherActionManaRegen (tag 1) | +3078.1 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | -1211.5 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | OtherActionManaRegen (tag 2) | +145.7 |

## Paladin — Protection · Tank

**Talents:** 0/43/8 · `-5521513321301551-15002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Seal of Fury is not implemented; shield and Spiritual Attunement questions remain in the uncertainty register. This is not a survival ranking.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 548.59 | 0.29 | 0.01 |
| Human | 528.53 | 0.28 | 0.01 |
| Dwarf | 520.24 | 0.27 | 0.02 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Undead | 820.68 | 698.43 | 75.34 | 0.14% |
| Human | 803.91 | 697.21 | 75.33 | 0.14% |
| Dwarf | 792.48 | 698.66 | 75.04 | 0.18% |

These stress-scenario results are not measured boss balance or an equal-support survival ranking.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Plate) [+18 Sta, +23 Int, +29 SP, +14 Crit, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Plate) [+14 Sta, +18 Int, +22 SP, +11 Crit, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Intellect / spell power / crit — Chest (Plate) [+18 Sta, +23 Int, +29 SP, +14 Crit, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / crit / hit — Hands (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / crit / hit — Legs (Plate) [+18 Sta, +21 Crit, +21 Str, +23 Hit, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Hybrid damage shield — Shield [+21 Sta, +9 Int, +6 Dmg, +18 Heal, 2468 armor, 44 Block] (modeled; reference only)](https://www.wowhead.com/forever/item=278469) | 65 | Enchant Shield - Greater Stamina |
| Ranged/relic | [Modeled: Physical libram — Relic [+8 Sta, +4 Agi, +26 AP] (modeled; reference only)](https://www.wowhead.com/forever/item=279247) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -3s: Cast primary seal.
- -1.5s: [Holy Shield](https://www.wowhead.com/forever/spell=20928).

### Rotation priorities

1. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271).
2. Cast [Swift Judgement](https://www.wowhead.com/forever/spell=53671) when `{"spellTimeToReady":{"spellId":{"spellId":20271}}}` ≥ 6s.
3. Use ready automatic cooldowns.
4. Cast [Templar's Bulwark](https://www.wowhead.com/forever/spell=1311015) when `{"currentHealthPercent":{}}` < 30%.
5. Cast [Holy Shield](https://www.wowhead.com/forever/spell=20928) when [Holy Shield](https://www.wowhead.com/forever/spell=20928) time remaining ≤ 2s.
6. Cast primary seal when Seal time remaining ≤ 2s.
7. Cast [Consecration](https://www.wowhead.com/forever/spell=20924) when Mana fraction ≥ 30%.
8. Cast [Consecration](https://www.wowhead.com/forever/spell=26573) when Mana fraction ≥ 15%.
9. Cast [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239).
10. Cast [Holy Strike](https://www.wowhead.com/forever/spell=10333).
11. Cast [Holy Shield](https://www.wowhead.com/forever/spell=20928).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 120.16 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 81.43 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 79.76 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 57.36 |
| Auto-attack (tag 3) | 51.50 |
| [Holy Shield](https://www.wowhead.com/forever/spell=20957) | 50.64 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 36.24 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 28.85 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -209529.8 |
| Health | OtherActionHealingModel | +203972.3 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -17854.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +14229.8 |
| Mana | [Shield Specialization](https://www.wowhead.com/forever/spell=20148) | +10702.2 |
| Mana | [Holy Shield](https://www.wowhead.com/forever/spell=20928) | -6480.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +4949.3 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -3098.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2473.2 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -2449.4 |
| Mana | OtherActionManaRegen (tag 1) | +2321.8 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -1980.0 |

## Shaman — Elemental

**Talents:** 31/7/13 · `5502301300123051-052-05305`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/modelled_gear_search/current/summary.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 545.95 | 0.23 | 0.34 |
| Tauren | 540.82 | 0.23 | 0.34 |
| Troll | 543.16 | 0.23 | 0.33 |
| Windshaper | 544.59 | 0.23 | 0.35 |
| Dwarf | 542.01 | 0.23 | 0.34 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Mail) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Mail) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Mail) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Mail) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Mail) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Axe one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272592) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power totem — Relic [+8 Sta, +7 Int, +9 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279249) | 65 | — |

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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 171.76 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 118.70 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | 96.63 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 87.22 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 35.19 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 20.80 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 8.64 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 5.97 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -7652.2 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -7342.0 |
| Mana | OtherActionManaRegen (tag 1) | +6882.3 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | -6085.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4849.3 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | -4500.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3607.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3569.8 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -765.0 |
| Mana | OtherActionManaRegen (tag 2) | +206.9 |

## Warrior — Protection · Tank

**Talents:** 4/6/41 · `31-0501-532531232300210351`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Different external support from the non-attacking DPS rows. Rage, shield proc eligibility and scripted Tier effects remain qualified in the uncertainty register. Heroic Strike/Cleave now replace swings. Earlier Protection gear/talent search gains used an invalid direct-cast APL and are superseded; the existing gear and talents have been replayed, not re-optimized.

**Talent variants:** The allocation above is for Undead. Other races retain their saved allocations in the exact profiles. Historical validation ranges and source qualifications are recorded separately.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 355.51 | 0.28 | 0.00 |
| Tauren | 357.04 | 0.28 | 0.00 |
| Troll | 354.95 | 0.27 | 0.00 |
| Undead | 368.41 | 0.29 | 0.00 |
| Windshaper | 355.90 | 0.28 | 0.00 |
| Human | 360.40 | 0.28 | 0.00 |
| Dwarf | 355.50 | 0.28 | 0.00 |
| Night Elf | 354.96 | 0.28 | 0.00 |
| Gnome | 357.39 | 0.27 | 0.00 |
| High Order | 355.90 | 0.28 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Orc | 814.38 | 480.47 | 48.54 | 0.00% |
| Tauren | 814.66 | 462.02 | 44.61 | 0.00% |
| Troll | 812.84 | 475.10 | 47.94 | 0.00% |
| Undead | 797.41 | 465.36 | 47.64 | 0.00% |
| Windshaper | 780.18 | 471.93 | 48.11 | 0.00% |
| Human | 787.28 | 475.77 | 48.09 | 0.00% |
| Dwarf | 813.55 | 473.33 | 47.69 | 0.00% |
| Night Elf | 778.20 | 461.71 | 47.58 | 0.00% |
| Gnome | 781.90 | 463.14 | 47.43 | 0.00% |
| High Order | 780.18 | 471.93 | 48.11 | 0.00% |

These stress-scenario results are not measured boss balance or an equal-support survival ranking.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / crit / hit — Neck [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Hybrid damage shield — Shield [+21 Sta, +9 Int, +6 Dmg, +18 Heal, 2468 armor, 44 Block] (modeled; reference only)](https://www.wowhead.com/forever/item=278469) | 65 | Enchant Shield - Greater Stamina |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Battle Shout Rank 7](https://www.wowhead.com/forever/spell=25289) when ([Battle Shout Rank 7](https://www.wowhead.com/forever/spell=25289) time remaining ≤ 3s AND Time remaining ≥ 8s AND Rage ≥ 10).
2. Cast [Shield Wall](https://www.wowhead.com/forever/spell=871) when `{"currentHealthPercent":{}}` < 35%.
3. Cast [Last Stand](https://www.wowhead.com/forever/spell=12975) when `{"currentHealthPercent":{}}` < 30%.
4. Cast [Shield Block](https://www.wowhead.com/forever/spell=2565) when (Rage ≥ 15 AND [Shield Block](https://www.wowhead.com/forever/spell=2565) time remaining < 1s).
5. Cast [Bloodrage](https://www.wowhead.com/forever/spell=2687) when (Rage < 35 AND `{"currentHealthPercent":{}}` > 50%).
6. Cast [Sunder Armor](https://www.wowhead.com/forever/spell=11597) when ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining ≤ 8s AND Time remaining > [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining).
7. Cast [Thunder Clap](https://www.wowhead.com/forever/spell=11581) when ([Thunder Clap](https://www.wowhead.com/forever/spell=11581) time remaining ≤ 3s AND Time remaining ≥ 4s).
8. Cast [Demoralizing Shout Rank 5](https://www.wowhead.com/forever/spell=11556) when ([Demoralizing Shout Rank 5](https://www.wowhead.com/forever/spell=11556) time remaining ≤ 4s AND Time remaining ≥ 5s).
9. Cast [Sunder Armor](https://www.wowhead.com/forever/spell=11597) when [Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks < 5.
10. Cast [Thunder Clap](https://www.wowhead.com/forever/spell=11581) when (Target count ≥ 3 AND (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Thunder Clap](https://www.wowhead.com/forever/spell=11581) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost)).
11. Cast [Revenge](https://www.wowhead.com/forever/spell=25288) when (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Revenge](https://www.wowhead.com/forever/spell=25288) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost).
12. Cast [Shield Slam](https://www.wowhead.com/forever/spell=23925) when (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Shield Slam](https://www.wowhead.com/forever/spell=23925) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost).
13. Cast [Cleave](https://www.wowhead.com/forever/spell=20569) when ((Target count > 1 AND Rage ≥ 20) AND (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Cleave](https://www.wowhead.com/forever/spell=20569) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost)).
14. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when ((Target count = 1 AND Rage ≥ 20) AND (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Heroic Strike](https://www.wowhead.com/forever/spell=25286) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost)).
15. Cast [Sunder Armor](https://www.wowhead.com/forever/spell=11597) when Rage ≥ 45.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 100.57 |
| [Shield Slam](https://www.wowhead.com/forever/spell=23925) | 85.74 |
| Auto-attack (tag 1) | 79.95 |
| Auto-attack (tag 3) | 44.89 |
| [Revenge](https://www.wowhead.com/forever/spell=25288) | 39.41 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 12.66 |
| [Thunder Clap](https://www.wowhead.com/forever/spell=11581) | 3.36 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 1.83 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -139606.9 |
| Health | OtherActionHealingModel | +137322.8 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1817.2 |
| Rage | Auto-attack (tag 1) | +684.0 |
| Rage | [Shield Slam](https://www.wowhead.com/forever/spell=23925) | -645.4 |
| Rage | [Shield Block](https://www.wowhead.com/forever/spell=2565) | -494.8 |
| Rage | OtherActionDamageTaken | +483.2 |
| Health | [Holy Strength](https://www.wowhead.com/forever/spell=20007) | +455.4 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -442.7 |
| Rage | [Shield Specialization](https://www.wowhead.com/forever/spell=12727) | +430.6 |
| Rage | OtherActionRefund | +199.8 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -188.3 |
