# Build reviews

These are hypothetical modeled-only gear results, not obtainable item rankings. Sources linked for modeled items are allocation references only and have different stats. Gear is compared under the [model assumptions](modelled_gear.md).

The tables and [matrix](../artifacts/modelled_gear/forever_dps_5min.png) use the same 201 common-seed replays. The modeled search and original real-item benchmark are separate; neither proves available launch gear. The [October 11 audit](october-11-audit.md) records the engine corrections, rotation and consumable selections and the bounded gear re-search. Earlier search gains are historical, not current proof of an optimum.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and hit from selected gear, enchants, talents and racials only. [Scenario and exchange model](../tools/forever_bench/README.md) · [Questions and coverage](uncertainties.md)

Tank rows include frontal incoming attacks and modeled healing, with their recorded shared raid support except assigned self-maintained duties. They do not share the non-attacking DPS encounter. [Support and Judgement correction](tank_support.md). [Tank selection and guardrails](tank_selection.md).

## Paladin — Retribution

**Talents:** 13/7/31 · `55003-232-05225231001330301`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Separate seal Echoes can coexist and fire on a landed white swing; their simultaneous server behavior remains qualified (PAL-002). Active presets use maximum registered ranks, not deliberate downranking.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1058.13 | 0.51 | 0.00 |
| Human | 1051.09 | 0.53 | 0.00 |
| Dwarf | 1052.83 | 0.52 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Plate) [+18 Sta, +23 Int, +29 SP, +14 Crit, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Plate) [+14 Sta, +18 Int, +22 SP, +11 Crit, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Intellect / spell power / crit — Chest (Plate) [+18 Sta, +23 Int, +29 SP, +14 Crit, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / crit / hit — Hands (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / crit / hit — Legs (Plate) [+18 Sta, +21 Crit, +21 Str, +23 Hit, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Spell two-hand — Two-hand [+25 Sta, +24 Int, +24 SP, +15 Str, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Fiery Weapon |
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
8. Cast [Consecration](https://www.wowhead.com/forever/spell=20924) when Mana fraction ≥ 20%.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 305.90 |
| Auto-attack (tag 1) | 172.30 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 158.54 |
| Auto-attack (tag 3) | 84.55 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 82.84 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 67.32 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 51.26 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 46.39 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +14108.8 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -10056.5 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -9112.4 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -6126.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3588.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3296.0 |
| Mana | OtherActionManaRegen (tag 1) | +3108.8 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -2940.5 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +2866.9 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1055.4 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -520.8 |
| Mana | OtherActionManaRegen (tag 2) | +1.4 |

## Hunter — Beast Mastery

**Talents:** 31/20/0 · `5320001505101251-00522510005`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Hawk lifetimes and ranged-speed behavior are source-backed; high-rank assault damage, crit and Ferocity transfer remain qualified (HUN-004). The two-Hawk policy is conditional on the central calibration.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 987.36 | 0.29 | 0.02 |
| Tauren | 970.12 | 0.27 | 0.03 |
| Troll | 979.33 | 0.28 | 0.02 |
| Windshaper | 980.12 | 0.27 | 0.02 |
| Human | 989.38 | 0.28 | 0.02 |
| Dwarf | 971.80 | 0.27 | 0.03 |
| Night Elf | 984.54 | 0.28 | 0.02 |
| High Order | 980.12 | 0.27 | 0.02 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Agility / crit / hit — Neck [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Minor Agility |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Bow ranged — Ranged bow [+6 Sta, +32 AP, 88–133 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -10s: [Aspect of the Hawk](https://www.wowhead.com/forever/spell=25296).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Intimidation](https://www.wowhead.com/forever/spell=19577) when Mana fraction > 80%.
3. Cast [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) when (Active Hawks < 2 OR Earliest Hawk expiry ≤ [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) travel time).
4. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=25295) when NOT [Serpent Sting](https://www.wowhead.com/forever/spell=25295) DoT active.
5. Cast [Aimed Shot](https://www.wowhead.com/forever/spell=20904) when `{"autoTimeToNext":{"autoType":"Ranged"}}` ≥ 0s.
6. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287) when Mana fraction ≥ 15%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Shoot | 387.67 |
| Cat: Auto-attack (tag 1) | 230.73 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 130.12 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 85.25 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 53.89 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 31.82 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 30.89 |
| Hawk 1: Auto-attack (tag 1) | 14.70 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -12885.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9231.7 |
| Mana | OtherActionManaRegen (tag 1) | +6389.0 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -5576.5 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -4462.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3606.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3601.2 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -2678.1 |
| Mana | OtherActionManaRegen (tag 2) | +743.0 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Intimidation](https://www.wowhead.com/forever/spell=19577) | -340.2 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Paladin — Physical Ret

**Talents:** 13/7/31 · `52003003-232-05225031101330311`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Uses a physical equipment budget without caster gear. Champion of the Light converts existing Intellect into damage-only spell bonus, not healing power. Seal Echo ordering remains qualified (PAL-002).


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 968.55 | 0.53 | 0.00 |
| Human | 960.18 | 0.58 | 0.02 |
| Dwarf | 950.20 | 0.58 | 0.02 |

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
| Main hand | [Modeled: Axe two-hand — Two-hand [+34 Sta, +23 Str, +14 Crit, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272593) | 65 | Enchant Weapon - Fiery Weapon |
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
8. Cast [Consecration](https://www.wowhead.com/forever/spell=20924) when Mana fraction ≥ 20%.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 302.40 |
| Auto-attack (tag 1) | 207.97 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 101.19 |
| Auto-attack (tag 3) | 100.65 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 78.73 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 43.07 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 34.28 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 30.68 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +14135.3 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -9912.3 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -8983.9 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -6510.2 |
| Mana | OtherActionManaRegen (tag 1) | +5392.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3418.7 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3199.6 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3122.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2825.5 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1333.3 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -520.7 |
| Mana | OtherActionManaRegen (tag 2) | +1.9 |

## Hunter — Marksmanship

**Talents:** 5/35/11 · `5-0050552011523051-50005001`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 961.32 | 0.37 | 0.00 |
| Tauren | 943.45 | 0.37 | 0.00 |
| Troll | 953.36 | 0.37 | 0.00 |
| Windshaper | 953.39 | 0.37 | 0.00 |
| Human | 965.50 | 0.37 | 0.00 |
| Dwarf | 945.49 | 0.37 | 0.00 |
| Night Elf | 960.86 | 0.37 | 0.00 |
| High Order | 953.39 | 0.37 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Agility / crit / hit — Neck [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Agility / crit / hit — Shoulders (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Agility / crit / hit — Back (Cloth) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Agility / crit / hit — Chest (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Agility / crit / hit — Wrists (Mail) [+11 Sta, +12 Agi, +12 Crit, +13 Hit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Greater Intellect |
| Hands | [Modeled: Agility / crit / hit — Hands (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Agility / crit / hit — Waist (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Agility / crit / hit — Legs (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Mail) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Bow ranged — Ranged bow [+6 Sta, +32 AP, 88–133 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272594) | 65 | SAF-T Ultra Precision Scope |

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
| Shoot | 503.77 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 178.28 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 114.37 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 93.10 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 65.99 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 9.99 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -11178.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +10412.3 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -6704.0 |
| Mana | OtherActionManaRegen (tag 1) | +6248.9 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -5742.8 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3904.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3599.7 |
| Mana | OtherActionManaRegen (tag 2) | +601.5 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Warrior — Fury

**Talents:** 18/33/0 · `20305213002-25053005141010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Level-60/off-hand rage and some extra-swing interactions remain provisional. The queued-Heroic-Strike off-hand hit bug is removed. Deep Wounds uses the reported intended no-AP rollover model; live delivery and off-hand payload attribution remain qualified.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 948.08 | 0.51 | 0.00 |
| Tauren | 940.54 | 0.51 | 0.00 |
| Troll | 932.97 | 0.49 | 0.00 |
| Undead | 947.93 | 0.51 | 0.00 |
| Windshaper | 930.83 | 0.50 | 0.00 |
| Human | 945.08 | 0.50 | 0.00 |
| Dwarf | 934.58 | 0.50 | 0.00 |
| Night Elf | 934.62 | 0.49 | 0.00 |
| Gnome | 931.53 | 0.51 | 0.00 |
| High Order | 930.83 | 0.50 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / crit / hit — Hands (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Axe one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272592) | 65 | Enchant Weapon - Crusader |
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
9. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
10. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 30 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 171.08 |
| Auto-attack (tag 2) | 155.22 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 138.05 |
| Auto-attack (tag 1) | 117.05 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 103.86 |
| Auto-attack (tag 3) | 68.41 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 56.49 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 52.87 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1288.0 |
| Rage | Auto-attack (tag 2) | +1286.6 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -1043.1 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -836.4 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -712.9 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -480.3 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +174.1 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +158.4 |
| Rage | OtherActionRefund | +114.3 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.9 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +89.8 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -87.0 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 915.53 | 0.42 | 0.00 |
| Troll | 905.85 | 0.41 | 0.00 |
| Undead | 928.65 | 0.42 | 0.00 |
| Human | 922.27 | 0.42 | 0.00 |
| Gnome | 911.82 | 0.41 | 0.00 |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 337.19 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 97.36 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 89.61 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 89.54 |
| Succubus: [Demonic Brand](https://www.wowhead.com/forever/spell=1293697) | 82.63 |
| Succubus: Auto-attack (tag 1) | 78.67 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 55.32 |
| [Soul Fire](https://www.wowhead.com/forever/spell=17924) | 50.72 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -24366.5 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +12069.4 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +11859.6 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6566.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5380.0 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5155.9 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -4263.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3884.1 |
| Mana | [Soul Fire](https://www.wowhead.com/forever/spell=17924) | -3157.2 |
| Mana | OtherActionManaRegen (tag 1) | +3113.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2591.5 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |

## Warrior — Fury (Sunder)

**Talents:** 18/33/0 · `20305213002-25053005141010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 914.97 | 0.52 | 0.00 |
| Tauren | 907.19 | 0.51 | 0.00 |
| Troll | 907.41 | 0.51 | 0.00 |
| Undead | 914.93 | 0.50 | 0.00 |
| Windshaper | 904.56 | 0.49 | 0.00 |
| Human | 920.13 | 0.50 | 0.00 |
| Dwarf | 895.15 | 0.50 | 0.00 |
| Night Elf | 903.00 | 0.50 | 0.00 |
| Gnome | 902.19 | 0.50 | 0.00 |
| High Order | 904.56 | 0.49 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / crit / hit — Legs (Plate) [+18 Sta, +21 Crit, +21 Str, +23 Hit, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
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
11. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
12. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 30 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 160.50 |
| Auto-attack (tag 2) | 154.13 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 124.62 |
| Auto-attack (tag 1) | 122.33 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 100.57 |
| Auto-attack (tag 3) | 71.19 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 56.17 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 51.65 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1379.2 |
| Rage | Auto-attack (tag 2) | +1320.3 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -964.8 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -820.3 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -692.7 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -450.7 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -275.0 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +173.2 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +172.2 |
| Rage | OtherActionRefund | +122.2 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.8 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +94.2 |

## Hunter — Pet/Melee

**Talents:** 16/11/24 · `53200005001-0050051-5302002300501201`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Two independent Hawks use provisional high-rank damage and crit calibration (HUN-004); no owner-crit or unverified Lone Wolf coexistence is assumed. Tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 903.26 | 0.36 | 0.00 |
| Windshaper | 910.83 | 0.36 | 0.00 |
| Orc | 914.73 | 0.37 | 0.00 |
| Human | 916.36 | 0.37 | 0.00 |
| Night Elf | 910.67 | 0.37 | 0.00 |
| Dwarf | 903.18 | 0.36 | 0.00 |
| Troll | 905.73 | 0.36 | 0.00 |
| High Order | 910.83 | 0.36 | 0.00 |

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
- -1.5s: [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325).

### Rotation priorities

1. Cast [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) when [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) time remaining ≤ 5s.
2. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
3. Use ready automatic cooldowns.
4. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
5. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
6. Cast [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) when (Time elapsed ≥ 0s AND Time remaining ≥ 9s).
7. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when Time remaining ≥ 12s.
8. Cast [Wing Clip](https://www.wowhead.com/forever/spell=14268) when Mana fraction ≥ 40%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 157.63 |
| Cat: Auto-attack (tag 1) | 143.20 |
| Auto-attack (tag 1) | 127.21 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 105.74 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 97.89 |
| Auto-attack (tag 3) | 73.99 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 56.60 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 37.85 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +15159.5 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -7643.3 |
| Mana | OtherActionManaRegen (tag 1) | +6650.7 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -4353.7 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3320.2 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -3027.7 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -1628.9 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -1543.9 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) | -180.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +28.1 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 902.72 | 0.37 | 0.00 |
| Tauren | 888.35 | 0.36 | 0.00 |
| Troll | 892.43 | 0.37 | 0.00 |
| Windshaper | 893.63 | 0.36 | 0.00 |
| Human | 903.82 | 0.37 | 0.00 |
| Dwarf | 887.14 | 0.37 | 0.00 |
| Night Elf | 898.69 | 0.37 | 0.00 |
| High Order | 893.63 | 0.36 | 0.00 |

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
- -1.5s: [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325).

### Rotation priorities

1. Cast [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) when [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) time remaining ≤ 5s.
2. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
3. Use ready automatic cooldowns.
4. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
5. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
6. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when Time remaining ≥ 12s.
7. Cast [Wing Clip](https://www.wowhead.com/forever/spell=14268) when Mana fraction ≥ 40%.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 161.81 |
| Cat: Auto-attack (tag 1) | 128.17 |
| Auto-attack (tag 1) | 126.99 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 108.56 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 99.35 |
| Auto-attack (tag 3) | 84.00 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 58.65 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 51.97 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +17071.8 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -9081.8 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4338.9 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -3659.3 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3397.7 |
| Mana | OtherActionManaRegen (tag 1) | +3156.5 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -2375.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +1956.3 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +188.7 |
| Mana | [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) | -180.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-55050000002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 826.66 | 0.58 | 0.00 |
| Tauren | 820.11 | 0.58 | 0.00 |
| Troll | 822.25 | 0.59 | 0.00 |
| Undead | 829.14 | 0.58 | 0.00 |
| Windshaper | 825.55 | 0.59 | 0.00 |
| Human | 832.72 | 0.59 | 0.00 |
| Dwarf | 818.62 | 0.58 | 0.00 |
| Night Elf | 829.83 | 0.59 | 0.00 |
| Gnome | 820.92 | 0.59 | 0.00 |
| High Order | 825.55 | 0.59 | 0.00 |

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
10. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 50.
11. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 70 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
12. Cast [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) when Rage ≥ 15.

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 155.17 |
| Auto-attack (tag 1) | 142.77 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 122.01 |
| Auto-attack (tag 3) | 114.30 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 80.47 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 57.12 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 45.95 |
| [Rend](https://www.wowhead.com/forever/spell=11574) | 37.93 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +2122.6 |
| Rage | [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | -1079.8 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -648.5 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -199.2 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -176.9 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -170.1 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -152.1 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +145.4 |
| Rage | [Rend](https://www.wowhead.com/forever/spell=11574) | -145.0 |
| Rage | OtherActionRefund | +99.9 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +85.9 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2525000013520105--0550015103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 816.72 | 0.48 | 0.00 |
| Troll | 806.74 | 0.47 | 0.00 |
| Undead | 818.27 | 0.47 | 0.00 |
| Human | 824.01 | 0.47 | 0.00 |
| Gnome | 805.23 | 0.46 | 0.00 |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 386.48 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 112.50 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 102.20 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 87.58 |
| Succubus: Auto-attack (tag 1) | 70.60 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 31.47 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 19.67 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 8.12 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -29396.7 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24650.2 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6577.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5346.0 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5232.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3599.5 |
| Mana | [Siphon Life](https://www.wowhead.com/forever/spell=18881) | -3450.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3327.2 |
| Mana | OtherActionManaRegen (tag 1) | +3096.7 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -459.1 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -338.4 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 813.33 | 0.34 | 0.00 |
| Troll | 816.57 | 0.34 | 0.00 |
| Undead | 820.95 | 0.34 | 0.00 |
| Human | 821.02 | 0.34 | 0.00 |
| Gnome | 805.62 | 0.34 | 0.00 |

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
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Fire Power |
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Spell Power |
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

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 262.29 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 124.88 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 99.31 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 90.37 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 85.18 |
| Succubus: Auto-attack (tag 1) | 76.43 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 53.92 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 16.37 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +21878.9 |
| Mana | [Incinerate](https://www.wowhead.com/forever/spell=1293813) | -19426.8 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6692.9 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -6347.6 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -5656.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5364.4 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5119.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4373.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3607.5 |
| Mana | OtherActionManaRegen (tag 1) | +3099.5 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -487.6 |

## Priest — Shadow

**Talents:** 18/0/33 · `305030001303--504020501201302251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 805.93 | 0.24 | 0.00 |
| Undead | 812.37 | 0.25 | 0.00 |
| Human | 805.30 | 0.24 | 0.00 |
| Dwarf | 811.41 | 0.25 | 0.00 |
| Night Elf | 811.23 | 0.24 | 0.00 |
| Gnome | 808.28 | 0.25 | 0.00 |

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
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 309.39 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 159.68 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 146.87 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 99.07 |
| [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | 87.52 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 9.84 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +8072.6 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -7105.5 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -5950.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4376.7 |
| Mana | [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | -3396.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2930.6 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2848.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +1782.6 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1411.7 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |

## Warrior — 2H Bloodthirst

**Talents:** 20/31/0 · `30304203032-55050020151010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 809.32 | 0.56 | 0.00 |
| Tauren | 792.12 | 0.55 | 0.00 |
| Undead | 802.62 | 0.57 | 0.00 |
| Troll | 798.16 | 0.56 | 0.00 |
| Windshaper | 794.95 | 0.56 | 0.00 |
| Human | 805.18 | 0.56 | 0.00 |
| Dwarf | 799.63 | 0.56 | 0.00 |
| Gnome | 792.98 | 0.55 | 0.00 |
| Night Elf | 797.65 | 0.55 | 0.00 |
| High Order | 794.95 | 0.56 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Rapidity |
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
9. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
10. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 50.
11. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 50 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 198.56 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 142.29 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 98.29 |
| Auto-attack (tag 3) | 88.27 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 82.70 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 68.09 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 45.09 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 43.30 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +2303.4 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -1020.2 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -709.0 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -474.3 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -270.0 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +145.1 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -123.4 |
| Rage | OtherActionRefund | +88.0 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +87.6 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +56.1 |
| Rage | [Berserker Stance](https://www.wowhead.com/forever/spell=2458) | -52.2 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -49.7 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `252200001351-0025003001-0550005103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 791.12 | 0.51 | 0.00 |
| Troll | 778.63 | 0.51 | 0.00 |
| Undead | 801.59 | 0.51 | 0.00 |
| Human | 805.67 | 0.51 | 0.00 |
| Gnome | 791.24 | 0.51 | 0.00 |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 456.99 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 122.32 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 108.74 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 89.66 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 21.79 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 3.87 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.29 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -31346.0 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +22387.3 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6903.3 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5481.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5402.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3543.3 |
| Mana | OtherActionManaRegen (tag 1) | +3117.4 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -573.6 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -362.7 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -194.1 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `550133102-053030031005102251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 761.52 | 0.57 | 0.02 |
| Tauren | 756.29 | 0.55 | 0.00 |
| Troll | 762.04 | 0.56 | 0.00 |
| Windshaper | 759.83 | 0.55 | 0.00 |
| Dwarf | 764.60 | 0.55 | 0.00 |

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
| Auto-attack (tag 1) | 222.53 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 141.46 |
| [Fire Nova](https://www.wowhead.com/forever/spell=408345) | 92.69 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 91.07 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 62.25 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 49.85 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 44.86 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 28.72 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Nova](https://www.wowhead.com/forever/spell=408345) | -14783.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9294.5 |
| Mana | OtherActionManaRegen (tag 1) | +7236.8 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -5380.1 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4735.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4488.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3596.7 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3307.8 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | [Spell 10614](https://www.wowhead.com/forever/spell=10614) | -499.8 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -450.0 |
| Mana | OtherActionManaRegen (tag 2) | +217.5 |

## Mage — Arcane

**Talents:** 33/3/15 · `053005023100311531-03-005500032`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 724.18 | 0.44 | 0.00 |
| Troll | 724.09 | 0.43 | 0.00 |
| Undead | 743.01 | 0.47 | 0.00 |
| Human | 728.87 | 0.44 | 0.00 |
| Gnome | 734.25 | 0.43 | 0.00 |
| High Order | 729.34 | 0.44 | 0.00 |

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
4. Cast [Frostbolt](https://www.wowhead.com/forever/spell=25304) when ([Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks ≥ 3 OR Mana fraction < 20%).
5. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 339.37 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 327.41 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 55.09 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 21.14 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -25740.2 |
| Mana | OtherActionManaRegen (tag 1) | +9894.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +6843.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3611.8 |
| Mana | OtherActionManaRegen (tag 2) | +2762.8 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -2093.6 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.9 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.2 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | -622.4 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +533.5 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Mage — Fire

**Talents:** 18/31/2 · `0501252000003-23450000130133051-002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Talent variants:** The allocation above is for Undead. Other races retain their saved allocations in the exact profiles. Historical validation ranges and source qualifications are recorded separately.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 728.49 | 0.46 | 1.73 |
| Troll | 723.31 | 0.43 | 0.92 |
| Undead | 741.30 | 0.44 | 1.50 |
| Human | 739.05 | 0.45 | 1.30 |
| Gnome | 730.48 | 0.45 | 1.40 |
| High Order | 714.22 | 0.45 | 2.71 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
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
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 268.86 |
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 171.89 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 166.56 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 99.22 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 17.30 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 8.86 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 8.62 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -15728.0 |
| Mana | OtherActionManaRegen (tag 1) | +14643.9 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -14427.8 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -7934.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5491.8 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +4419.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4171.4 |
| Mana | OtherActionManaRegen (tag 2) | +1859.4 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.1 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +850.5 |
| Mana | [Fireball](https://www.wowhead.com/forever/spell=25306) | -642.6 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +599.4 |

## Mage — Arcane–Frost

**Talents:** 28/0/23 · `05020500310031053--05550002010003002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Ice Lance uses a provisional 0.10 SP share calibrated from lower ranks. Fingers of Frost ordering and Arcane Missiles/Clearcasting proc interpretation remain unresolved; Missile Barrage now requires a landed eligible hit.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 726.60 | 0.37 | 0.00 |
| Troll | 724.80 | 0.36 | 0.00 |
| Undead | 739.74 | 0.37 | 0.00 |
| Human | 736.47 | 0.37 | 0.00 |
| Gnome | 732.61 | 0.37 | 0.00 |
| High Order | 732.92 | 0.37 | 0.00 |

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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 441.52 |
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 154.99 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 126.08 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 17.15 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -21894.2 |
| Mana | OtherActionManaRegen (tag 1) | +10762.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5630.0 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3132.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2832.4 |
| Mana | OtherActionManaRegen (tag 2) | +1792.3 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1065.2 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +704.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +341.2 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +243.2 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Rogue — Mutilate

**Talents:** 31/20/0 · `0053031035140105-302303202014`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 705.64 | 0.30 | 0.00 |
| Troll | 702.35 | 0.29 | 0.00 |
| Undead | 736.31 | 0.29 | 0.00 |
| Windshaper | 707.83 | 0.29 | 0.00 |
| Human | 702.31 | 0.30 | 0.00 |
| Dwarf | 699.75 | 0.29 | 0.00 |
| Night Elf | 708.96 | 0.29 | 0.00 |
| Gnome | 705.28 | 0.29 | 0.00 |
| High Order | 707.83 | 0.29 | 0.00 |

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
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Fiery Weapon |
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
| Auto-attack (tag 1) | 159.53 |
| Auto-attack (tag 2) | 100.25 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 79.03 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 72.57 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 56.03 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 50.90 |
| Auto-attack (tag 3) | 47.73 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 40.18 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3013.0 |
| Energy | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | -2937.0 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -840.8 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +797.5 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -400.6 |
| Energy | OtherActionRefund | +195.9 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -137.5 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.7 |
| ComboPoints | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | +91.6 |
| ComboPoints | [Seal Fate](https://www.wowhead.com/forever/spell=14195) | +37.1 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +33.3 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -22.7 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 718.58 | 0.31 | 0.00 |
| Troll | 709.43 | 0.30 | 0.00 |
| Undead | 735.42 | 0.31 | 0.00 |
| Windshaper | 713.95 | 0.30 | 0.00 |
| Human | 716.70 | 0.30 | 0.00 |
| Dwarf | 709.44 | 0.31 | 0.00 |
| Night Elf | 712.51 | 0.29 | 0.00 |
| Gnome | 710.79 | 0.29 | 0.00 |
| High Order | 713.95 | 0.30 | 0.00 |

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
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Fiery Weapon |
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
| Auto-attack (tag 1) | 174.30 |
| [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | 125.84 |
| Auto-attack (tag 2) | 116.72 |
| Auto-attack (tag 3) | 95.53 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 83.91 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 38.58 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 30.97 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 24.41 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3176.0 |
| Energy | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | -2932.6 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -467.7 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -463.6 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +462.5 |
| Energy | OtherActionRefund | +122.6 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +97.3 |
| Energy | [Blade Flurry](https://www.wowhead.com/forever/spell=13877) | -75.0 |
| ComboPoints | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | +70.0 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -66.4 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -26.4 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +24.5 |

## Druid — Feral

**Talents:** 9/34/8 · `050022-35210032021032212041-053`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Blizzard supports proc-enchant form eligibility, and Crusader's sourced health return is modeled. Its exact PPM/conditional interactions remain qualified (DRU-011); the Night Elf profile uses that enchant.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 726.44 | 0.23 | 2.36 |
| Windshaper | 729.39 | 0.24 | 1.65 |
| Night Elf | 730.62 | 0.23 | 2.13 |
| High Order | 729.39 | 0.24 | 1.65 |

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
3. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Time remaining ≤ 4s AND Combo points ≥ 3).
4. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Combo points ≥ 5 AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 5s).
5. Cast [Rip](https://www.wowhead.com/forever/spell=9896) when (Combo points ≥ 4 AND NOT [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND Time remaining ≥ 8s).
6. Cast [Shifting Power](https://www.wowhead.com/forever/spell=1322605) when Energy ≤ 60.
7. Cast [Shred](https://www.wowhead.com/forever/spell=9830).

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| [Shred](https://www.wowhead.com/forever/spell=9830) | 275.35 |
| Auto-attack (tag 1) | 205.90 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 175.76 |
| Auto-attack (tag 3) | 48.05 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 22.30 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 3.26 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shifting Power](https://www.wowhead.com/forever/spell=1322605) | -28399.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +14601.1 |
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -4149.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4072.4 |
| Mana | OtherActionManaRegen (tag 2) | +3458.3 |
| Energy | OtherActionEnergyRegen | +3010.4 |
| Energy | [Shifting Power](https://www.wowhead.com/forever/spell=1322605) | +1660.3 |
| Mana | OtherActionManaRegen (tag 1) | +1641.6 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -692.6 |
| Energy | OtherActionRefund | +260.9 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -253.6 |
| ComboPoints | [Rip](https://www.wowhead.com/forever/spell=9896) | -112.2 |

## Druid — Balance

**Talents:** 38/0/13 · `5232220115501351--505003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 707.93 | 0.33 | 0.00 |
| Windshaper | 711.17 | 0.33 | 0.00 |
| Night Elf | 711.18 | 0.33 | 0.00 |
| High Order | 711.17 | 0.33 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Night Elf

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

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 411.67 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 107.16 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 98.26 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 94.09 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=25298) | -15057.0 |
| Mana | OtherActionManaRegen (tag 1) | +7684.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4560.7 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -4394.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3604.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.3 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1735.0 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1213.9 |
| Mana | [Moonkin Form](https://www.wowhead.com/forever/spell=24858) | -435.4 |
| Mana | OtherActionManaRegen (tag 2) | +50.6 |

## Mage — Frost

**Talents:** 11/3/37 · `050005001-03-0555003321001301251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 681.35 | 0.38 | 0.22 |
| Troll | 678.64 | 0.38 | 0.19 |
| Undead | 683.59 | 0.38 | 0.35 |
| Human | 685.76 | 0.38 | 0.15 |
| Gnome | 690.80 | 0.39 | 0.02 |
| High Order | 680.82 | 0.38 | 0.21 |

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
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 518.27 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 172.53 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -23391.5 |
| Mana | OtherActionManaRegen (tag 1) | +9425.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3975.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3723.1 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3580.9 |
| Mana | OtherActionManaRegen (tag 2) | +2030.5 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +852.0 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +848.4 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +175.2 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 665.40 | 0.26 | 0.00 |
| Troll | 659.27 | 0.25 | 0.00 |
| Undead | 682.96 | 0.26 | 0.00 |
| Windshaper | 661.66 | 0.26 | 0.00 |
| Human | 664.24 | 0.26 | 0.00 |
| Dwarf | 653.83 | 0.25 | 0.00 |
| Night Elf | 662.62 | 0.26 | 0.00 |
| Gnome | 661.65 | 0.26 | 0.00 |
| High Order | 661.66 | 0.26 | 0.00 |

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
| Auto-attack (tag 1) | 165.09 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 138.16 |
| Auto-attack (tag 2) | 88.88 |
| Auto-attack (tag 3) | 57.17 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 41.90 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 41.46 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 30.63 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 29.84 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +2984.2 |
| Energy | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | -2826.7 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +544.6 |
| Energy | [Rupture](https://www.wowhead.com/forever/spell=11275) | -486.5 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -344.9 |
| Energy | OtherActionRefund | +179.4 |
| Energy | [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) | -111.2 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +98.9 |
| ComboPoints | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | +84.7 |
| Energy | [Ambush](https://www.wowhead.com/forever/spell=11269) | -84.2 |
| ComboPoints | [Rupture](https://www.wowhead.com/forever/spell=11275) | -63.6 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -39.6 |

## Paladin — Protection · Tank

**Talents:** 0/42/9 · `-5520513321301551-15102`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Seal of Fury is now selectable with sourced damage, shield requirement and talent-only depletion mana. Its shared replacement shield pool and independent lifetime are provisional; the primary seal and lifetime are recorded, not silently substituted. Threat-based taunt target selection is not modeled. Gear selections must pass both workload guard policies, which permit limited mitigation tradeoffs. This is not a survival ranking. Reflections use explicit provider power, never the receiving tank's spell power.

**Talent variants:** The allocation above is for Undead. Other races retain their saved allocations in the exact profiles. Historical validation ranges and source qualifications are recorded separately.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 613.58 | 0.29 | 0.01 |
| Human | 599.49 | 0.27 | 0.01 |
| Dwarf | 586.05 | 0.27 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Undead | 914.31 | 697.99 | 75.29 | 0.16% |
| Human | 917.12 | 695.66 | 76.07 | 0.38% |
| Dwarf | 894.40 | 697.99 | 75.97 | 0.08% |

These stress-scenario results are not measured boss balance or an equal-support survival ranking.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Plate) [+18 Sta, +23 Int, +29 SP, +14 Crit, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Plate) [+14 Sta, +18 Int, +22 SP, +11 Crit, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Intellect / spell power / crit — Chest (Plate) [+18 Sta, +23 Int, +29 SP, +14 Crit, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Spell power / crit / hit (matched cost) — Wrists (Plate) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Spell power / crit / hit (matched cost) — Hands (Plate) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Plate) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Spell power / crit / hit (matched cost) — Legs (Plate) [+18 Sta, +21 Crit, +23 Hit, +23 SP, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Spell power / crit / hit (matched cost) — Feet (Plate) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Fiery Weapon |
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
8. Cast [Consecration](https://www.wowhead.com/forever/spell=20924) when Mana fraction ≥ 15%.
9. Cast [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239).
10. Cast [Holy Strike](https://www.wowhead.com/forever/spell=10333).
11. Cast [Holy Shield](https://www.wowhead.com/forever/spell=20928).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 106.21 |
| Auto-attack (tag 1) | 99.01 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 94.96 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 73.83 |
| [Holy Shield](https://www.wowhead.com/forever/spell=20957) | 63.21 |
| Auto-attack (tag 3) | 45.44 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 36.02 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 33.32 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -209396.0 |
| Health | OtherActionHealingModel | +203581.6 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -17982.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +15103.9 |
| Mana | [Shield Specialization](https://www.wowhead.com/forever/spell=20148) | +11878.0 |
| Mana | [Holy Shield](https://www.wowhead.com/forever/spell=20928) | -6480.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +5793.9 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -3097.6 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -2659.4 |
| Mana | OtherActionManaRegen (tag 1) | +2265.0 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -1980.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +1764.7 |

## Druid — Bear · Tank

**Talents:** 1/40/10 · `01-50022332120132012551-055`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Bear's outgoing base Rage remains the qualified damage-based model (DRU-012), with sourced double critical Rage. Incoming Rage now shares Warrior's configurable provisional expected-health scenario. The form threat multiplier is sourced at 1.5; this is not a survival ranking. Proc-enchant form eligibility is supported. Recovery has a sourced heal/ICD; Revelation requires an explicit, unverified proc model. Historical enchant-selection gains do not validate those assumptions.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 543.27 | 0.35 | 0.00 |
| Windshaper | 545.39 | 0.36 | 0.00 |
| Night Elf | 544.61 | 0.36 | 0.00 |
| High Order | 545.39 | 0.36 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Tauren | 1868.32 | 715.29 | 72.64 | 0.18% |
| Windshaper | 1869.07 | 714.94 | 77.61 | 0.34% |
| Night Elf | 1871.08 | 699.86 | 76.91 | 0.30% |
| High Order | 1869.07 | 714.94 | 77.61 | 0.34% |

These stress-scenario results are not measured boss balance or an equal-support survival ranking.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / AP / crit — Head (Leather) [+18 Sta, +9 Agi, +59 AP, +4 Crit, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=20068) | 65 | Lesser Arcanum of Voracity |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Maul](https://www.wowhead.com/forever/spell=9881) | 230.24 |
| [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073) | 110.48 |
| [Lacerate](https://www.wowhead.com/forever/spell=414644) | 90.91 |
| Auto-attack (tag 1) | 51.40 |
| [Lacerate](https://www.wowhead.com/forever/spell=414647) | 33.34 |
| Auto-attack (tag 3) | 26.82 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.21 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -215775.8 |
| Health | OtherActionHealingModel | +212857.9 |
| Rage | [Lacerate](https://www.wowhead.com/forever/spell=414644) | -1254.3 |
| Health | [Frenzied Regeneration](https://www.wowhead.com/forever/spell=22842) | +978.2 |
| Mana | [Faerie Fire](https://www.wowhead.com/forever/spell=9907) | -929.9 |
| Rage | Auto-attack (tag 1) | +892.9 |
| Rage | [Maul](https://www.wowhead.com/forever/spell=9881) | -844.5 |
| Rage | OtherActionDamageTaken | +810.0 |
| Rage | [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073) | -743.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +732.1 |
| Health | [Holy Strength](https://www.wowhead.com/forever/spell=20007) | +591.8 |
| Rage | [Blood Frenzy](https://www.wowhead.com/forever/spell=37117) | +541.0 |

## Warrior — Protection · Tank

**Talents:** 8/6/37 · `35-0501-251333120230021351`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Incoming Rage uses a configurable, unverified expected-creature-health/reference-Armor scenario shared with Bear. Its coefficient, level lookup and reference curve are assumptions, not measured server behavior; resolved parameters accompany each replay. Heroic Strike/Cleave replace swings and the APL reserves Rage for self-maintained Sunder. Earlier direct-cast Protection searches are superseded. Last Stand removes its temporary health on expiration; the nonlethal floor follows the qualified Classic convention.

**Talent variants:** The allocation above is for Undead. Other races retain their saved allocations in the exact profiles. Historical validation ranges and source qualifications are recorded separately.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 501.95 | 0.35 | 0.00 |
| Tauren | 508.51 | 0.35 | 0.00 |
| Troll | 504.97 | 0.35 | 0.00 |
| Undead | 523.24 | 0.33 | 0.00 |
| Windshaper | 505.32 | 0.35 | 0.00 |
| Human | 508.28 | 0.37 | 0.00 |
| Dwarf | 502.29 | 0.35 | 0.00 |
| Night Elf | 497.44 | 0.33 | 0.00 |
| Gnome | 507.23 | 0.35 | 0.00 |
| High Order | 505.32 | 0.35 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Orc | 1256.91 | 504.55 | 51.39 | 0.00% |
| Tauren | 1250.31 | 480.10 | 46.54 | 0.00% |
| Troll | 1242.85 | 490.24 | 49.93 | 0.00% |
| Undead | 1211.80 | 464.97 | 49.42 | 0.00% |
| Windshaper | 1188.65 | 485.86 | 50.03 | 0.00% |
| Human | 1189.54 | 472.42 | 49.48 | 0.00% |
| Dwarf | 1258.17 | 496.81 | 50.56 | 0.00% |
| Night Elf | 1164.41 | 481.44 | 49.83 | 0.00% |
| Gnome | 1194.17 | 481.96 | 49.49 | 0.00% |
| High Order | 1188.65 | 485.86 | 50.03 | 0.00% |

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
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Modeled: Physical shield — Shield [+24 Sta, +10 Str, 2468 armor, 44 Block] (modeled; reference only)](https://www.wowhead.com/forever/item=272591) | 65 | Enchant Shield - Greater Stamina |
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
| [Revenge](https://www.wowhead.com/forever/spell=25288) | 162.35 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 117.11 |
| [Shield Slam](https://www.wowhead.com/forever/spell=23925) | 91.20 |
| Auto-attack (tag 1) | 79.28 |
| Auto-attack (tag 3) | 37.41 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 19.57 |
| [Thunder Clap](https://www.wowhead.com/forever/spell=11581) | 7.36 |
| [Fiery Weapon](https://www.wowhead.com/forever/spell=13897) | 6.91 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -139490.8 |
| Health | OtherActionHealingModel | +136913.7 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +2575.4 |
| Rage | Auto-attack (tag 1) | +907.6 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -732.2 |
| Rage | [Shield Slam](https://www.wowhead.com/forever/spell=23925) | -686.8 |
| Rage | OtherActionDamageTaken | +526.8 |
| Rage | [Shield Block](https://www.wowhead.com/forever/spell=2565) | -511.2 |
| Rage | [Shield Specialization](https://www.wowhead.com/forever/spell=12727) | +428.6 |
| Rage | OtherActionRefund | +247.8 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -190.1 |
| Rage | [Shield Specialization](https://www.wowhead.com/forever/spell=12727) | +161.6 |

## Shaman — Elemental

**Talents:** 31/7/13 · `5502301300123051-052-05305`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 495.05 | 0.35 | 41.71 |
| Tauren | 490.00 | 0.35 | 42.17 |
| Troll | 490.86 | 0.34 | 43.31 |
| Windshaper | 490.49 | 0.35 | 44.14 |
| Dwarf | 491.08 | 0.35 | 42.08 |

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
5. Cast [Lava Burst (rank 3)](https://www.wowhead.com/forever/spell=1238300).
6. Cast [Chain Lightning (rank 4)](https://www.wowhead.com/forever/spell=10605) when ([Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active OR Target count ≥ 2).
7. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when Mana ≥ Time remaining × 35.
8. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 269.11 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | 86.24 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 73.10 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 34.82 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 17.40 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 13.52 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 0.86 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -16633.9 |
| Mana | OtherActionManaRegen (tag 1) | +6908.1 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -6250.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5254.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4167.3 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | -4120.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3604.9 |
| Mana | OtherActionManaRegen (tag 2) | +972.9 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -765.0 |

## Priest — Smite

**Talents:** 31/17/3 · `515030031305001031-00505023002-003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Uses maximum registered ranks. Penance's static rank anomaly and actual skill-book acquisition remain source gaps; no efficient lower-rank fallback is assumed.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 443.65 | 0.18 | 41.03 |
| Undead | 464.15 | 0.19 | 31.55 |
| Human | 446.32 | 0.18 | 36.88 |
| Dwarf | 443.51 | 0.18 | 38.52 |
| Night Elf | 444.75 | 0.18 | 39.70 |
| Gnome | 454.73 | 0.18 | 34.22 |

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
2. Cast [Holy Fire (rank 8)](https://www.wowhead.com/forever/spell=15261) when [Holy Fire (rank 8)](https://www.wowhead.com/forever/spell=15261) DoT time remaining < 3s.
3. `{"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10934,"rank":8}}}]}}`.
4. Cast [Smite](https://www.wowhead.com/forever/spell=10934) when Mana ≥ Time remaining × 20.
5. Cast [Smite](https://www.wowhead.com/forever/spell=10934).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Smite](https://www.wowhead.com/forever/spell=10934) | 328.75 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 124.24 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 11.16 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -23150.3 |
| Mana | OtherActionManaRegen (tag 1) | +7841.1 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -6261.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3604.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3603.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3548.9 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1574.6 |
| Mana | OtherActionManaRegen (tag 2) | +721.6 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |

## Shaman — Stormcaller

**Talents:** 29/22/0 · `550133130010304-055030021004002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_11/selection-evidence.json)

**Model limitations:** Uses maximum registered Lightning Bolt. Overload uses its own client rank rows; Maelstrom/proc and server-script questions remain qualified.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 457.51 | 0.31 | 65.33 |
| Tauren | 461.01 | 0.31 | 62.98 |
| Troll | 453.07 | 0.30 | 66.95 |
| Windshaper | 452.88 | 0.30 | 67.63 |
| Dwarf | 455.32 | 0.31 | 64.96 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Tauren

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Mail) [+18 Sta, +23 Int, +29 SP, +14 Crit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Mail) [+14 Sta, +18 Int, +22 SP, +11 Crit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Mail) [+11 Sta, +13 Int, +17 SP, +8 Crit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
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
7. Cast [Lightning Bolt](https://www.wowhead.com/forever/spell=15208).

### Damage breakdown — Tauren

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 301.87 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 85.84 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 34.42 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 22.57 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 15.17 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.13 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -17792.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5307.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3905.1 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -3795.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3599.5 |
| Mana | OtherActionManaRegen (tag 1) | +2731.7 |
| Mana | OtherActionManaRegen (tag 2) | +1478.2 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
