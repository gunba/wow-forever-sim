# Build reviews

These are hypothetical modeled-only gear results, not obtainable item rankings. Sources linked for modeled items are allocation references only and have different stats. Gear is compared under the [model assumptions](modelled_gear.md).

The tables and [matrix](../artifacts/modelled_gear/forever_dps_5min.png) use the same 226 common-seed replays. The modeled search and original real-item benchmark are separate; neither proves available launch gear. The [October 22 learned rotations](october-22-learned-rotations.md) add the matrix's Learned column: what a reinforcement-learning agent reached for each DPS build and, where rules converted from it beat the published APL, a simmable profile; the published profiles do not change. The [October 21 Warrior rotations](october-21-warrior-rotations.md) record the Fury, Fury (Sunder) and 2H Bloodthirst rotation rules learned by reinforcement learning and the 19/31/1 Fury talents; the [October 21 follow-up](october-21-follow-up.md) records the Marksmanship and Beast Mastery weave, the Fury Hamstring weave, one Twist of Light Echo at a time, measured rage from damage taken and the original race offsets; the [October 21 update](october-21-update.md) records the wowsims/forever review, server hotfixes, Forever consumables, the SV Weave, Frostfire and Shockadin builds and the re-searched talents and consumables; the [October 16 follow-up](october-16-follow-up.md) and [October 15 update](october-15-update.md) record the client 1.60.1.70425 changes and the earlier [October 11 audit](october-11-audit.md) and [October 11 log audit](october-11-log-audit.md) record the corrections before it. Earlier search gains are historical, not current proof of an optimum.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and hit from selected gear, enchants, talents and racials only. [Scenario and exchange model](../tools/forever_bench/README.md) · [Questions and coverage](uncertainties.md)

Tank rows include frontal incoming attacks and modeled healing, with their recorded shared raid support except assigned self-maintained duties. They do not share the non-attacking DPS encounter. [Support and Judgement correction](tank_support.md). [Tank selection and guardrails](tank_selection.md).

## Hunter — SV Weave

**Talents:** 0/20/31 · `-00530510105-530200230050122141`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Melee weave from 12 yards: steps in for Raptor Strike, Mongoose Bite, Strider Kick and Immolation Trap while the next Auto Shot is more than 1.5 s away, then back out. No pet. Movement takes the engine's run speed; no latency is modelled.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 1373.03 | 0.44 | 0.46 |
| Tauren | 1358.42 | 0.43 | 0.53 |
| Troll | 1353.56 | 0.44 | 0.48 |
| Windshaper | 1385.91 | 0.44 | 0.48 |
| Human | 1378.00 | 0.44 | 0.23 |
| Dwarf | 1362.64 | 0.43 | 0.49 |
| Night Elf | 1381.75 | 0.44 | 0.45 |
| High Order | 1385.91 | 0.44 | 0.48 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

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
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 90–136 dmg @ 2.4s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 90–136 dmg @ 2.4s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 147–222 dmg @ 3s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -30s: [Immolation Trap](https://www.wowhead.com/forever/spell=14305).
- -10s: [Aspect of the Hawk](https://www.wowhead.com/forever/spell=14322).
- -4s: [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325).
- -1.25s: [Aimed Shot](https://www.wowhead.com/forever/spell=20904).

### Rotation priorities

1. Cast [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) when [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) time remaining ≤ 5s.
2. Use ready automatic cooldowns.
3. Cast [Rapid Fire](https://www.wowhead.com/forever/spell=3045).
4. Cast [Multi-Shot](https://www.wowhead.com/forever/spell=2643).
5. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=13555) when [Serpent Sting](https://www.wowhead.com/forever/spell=13555) DoT time remaining ≤ 3s.
6. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287).
7. `{"condition":{"cmp":{"op":"OpGt","lhs":{"autoTimeToNext":{"autoType":"Ranged"}},"rhs":{"const":{"val":"1.5s"}}}},"move":{"rangeFromTarget":{"const":{"val":"5"}}}}` when `{"autoTimeToNext":{"autoType":"Ranged"}}` > 1.5s.
8. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when (Time remaining ≥ 6s AND `{"spellCanCast":{"spellId":{"spellId":14268}}}` AND [Immolation Trap](https://www.wowhead.com/forever/spell=14305) DoT time remaining ≤ 0s).
9. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
10. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
11. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
12. `{"condition":{"and":{"vals":[{"not":{"val":{"auraIsActive":{"auraId":{"spellId":14266}}}}},{"cmp":{"op":"OpGt","lhs":{"autoTimeToNext":{"autoType":"Melee"}},"rhs":{"const":{"val":"0.5s"}}}}]}},"move":{"rangeFromTarget":{"const":{"val":"12"}}}}` when (NOT [Raptor Strike](https://www.wowhead.com/forever/spell=14266) active AND `{"autoTimeToNext":{"autoType":"Melee"}}` > 0.5s).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Shoot | 439.40 |
| Auto-attack (tag 2) | 142.06 |
| [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | 134.23 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | 112.10 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 110.05 |
| Auto-attack (tag 1) | 107.74 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 92.00 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 81.25 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +15170.5 |
| Mana | [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | -9169.4 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -7084.7 |
| Mana | OtherActionManaRegen (tag 1) | +6387.9 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | -4619.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3392.0 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -2799.8 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -2555.5 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -1414.2 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -1283.4 |
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -310.0 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Hunter — MM Weave

**Talents:** 0/27/24 · `-005305031050302-5302002300501201`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Melee weave from 12 yards with Marksmanship and Survival talents (wowsims/forever MM Weave): steps in for Raptor Strike, Mongoose Bite, Strider Kick and Immolation Trap while the next Auto Shot is more than 1.5 s away, then back out, and keeps its own Hunter's Mark for Expose Prey. Movement takes the engine's run speed; no latency is modelled. Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 1313.06 | 0.41 | 0.26 |
| Tauren | 1299.28 | 0.40 | 0.32 |
| Troll | 1297.39 | 0.41 | 0.34 |
| Windshaper | 1328.04 | 0.41 | 0.19 |
| Human | 1316.68 | 0.41 | 0.14 |
| Dwarf | 1302.89 | 0.40 | 0.30 |
| Night Elf | 1320.68 | 0.41 | 0.26 |
| High Order | 1328.04 | 0.41 | 0.19 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

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
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 90–136 dmg @ 2.4s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 90–136 dmg @ 2.4s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 147–222 dmg @ 3s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -15s: [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447).
- -12s: [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325).
- 0.05s − `{"spellCastTime":{"spellId":{"spellId":20904}}}` × 5 / 7 + [Aimed Shot](https://www.wowhead.com/forever/spell=20904) travel time: [Aimed Shot](https://www.wowhead.com/forever/spell=20904).
- -7s: [Rapid Fire](https://www.wowhead.com/forever/spell=3045).
- -30s: [Explosive Trap](https://www.wowhead.com/forever/spell=14317).

### Rotation priorities

1. Cast [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) when [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) time remaining ≤ 6s.
2. Use ready automatic cooldowns.
3. Cast [Rapid Fire](https://www.wowhead.com/forever/spell=3045).
4. Cast [Explosive Trap](https://www.wowhead.com/forever/spell=14317) when (Time remaining ≤ 10s AND `{"spellCanCast":{"spellId":{"spellId":14268}}}`).
5. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when `{"spellCanCast":{"spellId":{"spellId":14268}}}`.
6. Cast [Multi-Shot](https://www.wowhead.com/forever/spell=2643).
7. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=13555) when [Serpent Sting](https://www.wowhead.com/forever/spell=13555) DoT time remaining ≤ 0.5s.
8. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287).
9. `{"condition":{"cmp":{"op":"OpGt","lhs":{"autoTimeToNext":{"autoType":"Ranged"}},"rhs":{"const":{"val":"1.5s"}}}},"move":{"rangeFromTarget":{"const":{"val":"5"}}}}` when `{"autoTimeToNext":{"autoType":"Ranged"}}` > 1.5s.
10. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
11. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
12. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
13. `{"condition":{"and":{"vals":[{"cmp":{"op":"OpGt","lhs":{"autoTimeToNext":{"autoType":"Melee"}},"rhs":{"const":{"val":"0.5s"}}}}]}},"move":{"rangeFromTarget":{"const":{"val":"12"}}}}` when (`{"autoTimeToNext":{"autoType":"Melee"}}` > 0.5s).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Shoot | 419.05 |
| [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | 144.72 |
| Auto-attack (tag 2) | 143.31 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 112.20 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | 111.27 |
| Auto-attack (tag 1) | 109.87 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 92.85 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 67.55 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +15217.6 |
| Mana | [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | -9313.8 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -7438.0 |
| Mana | OtherActionManaRegen (tag 1) | +6461.2 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | -4592.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3508.6 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -2799.8 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -1550.0 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -1531.9 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -1369.1 |
| Mana | [Explosive Trap](https://www.wowhead.com/forever/spell=14317) | -706.2 |
| Mana | OtherActionManaRegen (tag 2) | +393.0 |

## Hunter — Survival

**Talents:** 7/12/32 · `502-005205-500230230050122151`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 1280.34 | 0.42 | 0.01 |
| Tauren | 1259.35 | 0.42 | 0.02 |
| Troll | 1267.41 | 0.42 | 0.01 |
| Windshaper | 1278.04 | 0.43 | 0.01 |
| Human | 1275.03 | 0.43 | 0.00 |
| Dwarf | 1263.04 | 0.42 | 0.01 |
| Night Elf | 1279.32 | 0.42 | 0.01 |
| High Order | 1278.04 | 0.43 | 0.01 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

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
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 147–222 dmg @ 3s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -31s: [Explosive Trap](https://www.wowhead.com/forever/spell=14317).
- -4s: [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447).
- -2s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.
- -1.5s: [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325).

### Rotation priorities

1. Cast [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) when [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) time remaining ≤ 5s.
2. Use ready automatic cooldowns.
3. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
4. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when (Time remaining ≥ 5s AND `{"spellCanCast":{"spellId":{"spellId":14268}}}` AND [Immolation Trap](https://www.wowhead.com/forever/spell=14305) DoT time remaining ≤ 0.2s).
5. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
6. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=13555) when [Serpent Sting](https://www.wowhead.com/forever/spell=13555) DoT time remaining ≤ 0.3s.
7. Cast [Multi-Shot](https://www.wowhead.com/forever/spell=2643).
8. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
9. `{"condition":{"cmp":{"op":"OpGt","lhs":{"autoTimeToNext":{"autoType":"Ranged"}},"rhs":{"const":{"val":"0.25s"}}}},"move":{"rangeFromTarget":{"const":{"val":"5"}}}}` when `{"autoTimeToNext":{"autoType":"Ranged"}}` > 0.25s.
10. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287).
11. `{"condition":{"and":{"vals":[{"not":{"val":{"auraIsActive":{"auraId":{"spellId":14266}}}}},{"cmp":{"op":"OpGt","lhs":{"autoTimeToNext":{"autoType":"Melee"}},"rhs":{"const":{"val":"0.75s"}}}},{"or":{"vals":[{"not":{"val":{"auraIsActive":{"auraId":{"spellId":1299448}}}}},{"not":{"val":{"auraIsActive":{"auraId":{"spellId":1310726}}}}}]}}]}},"move":{"rangeFromTarget":{"const":{"val":"12"}}}}` when (NOT [Raptor Strike](https://www.wowhead.com/forever/spell=14266) active AND `{"autoTimeToNext":{"autoType":"Melee"}}` > 0.75s AND (NOT [Quick Strikes](https://www.wowhead.com/forever/spell=1299448) active OR NOT [Expose Prey state](https://www.wowhead.com/forever/spell=1310726) active)).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| Shoot | 274.46 |
| Auto-attack (tag 2) | 141.17 |
| Cat: Auto-attack (tag 1) | 131.91 |
| Auto-attack (tag 1) | 109.65 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 93.77 |
| [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | 89.69 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | 76.75 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 68.86 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +13912.6 |
| Mana | [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | -7807.4 |
| Mana | OtherActionManaRegen (tag 1) | +6379.7 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -4928.7 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | -4133.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2837.0 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -2717.9 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -2510.4 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -2451.7 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -1496.0 |
| Mana | [Explosive Trap](https://www.wowhead.com/forever/spell=14317) | -364.0 |
| Mana | OtherActionManaRegen (tag 2) | +328.3 |

## Hunter — BM Weave

**Talents:** 31/20/0 · `5320001505101251-00532500005`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Melee weave from 12 yards with the Cat and two Hawks out: steps in for Raptor Strike, Mongoose Bite and Immolation Trap while the next Auto Shot is more than 1.5 s away, then back out (wowsims/forever BM Weave). Movement takes the engine's run speed; no latency is modelled. Hawk lifetimes and ranged-speed behavior are source-backed; high-rank assault damage, crit and Ferocity transfer remain qualified (HUN-004). The two-Hawk policy is conditional on the central calibration.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 1201.49 | 0.33 | 3.87 |
| Tauren | 1176.57 | 0.32 | 4.21 |
| Troll | 1178.08 | 0.33 | 4.94 |
| Windshaper | 1192.48 | 0.33 | 4.81 |
| Human | 1193.21 | 0.32 | 2.57 |
| Dwarf | 1179.23 | 0.32 | 4.16 |
| Night Elf | 1195.85 | 0.33 | 3.84 |
| High Order | 1192.48 | 0.33 | 4.81 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

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
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 90–136 dmg @ 2.4s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 90–136 dmg @ 2.4s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 147–222 dmg @ 3s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -31s: [Immolation Trap](https://www.wowhead.com/forever/spell=14305).
- -10s: [Aspect of the Hawk](https://www.wowhead.com/forever/spell=14322).
- -1.7s: [Aimed Shot](https://www.wowhead.com/forever/spell=20904).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Intimidation](https://www.wowhead.com/forever/spell=19577) when Mana fraction > 100%.
3. Cast [Rapid Fire](https://www.wowhead.com/forever/spell=3045).
4. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287).
5. Cast [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) when (Active Hawks < 2 OR Earliest Hawk expiry ≤ [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) travel time).
6. Cast [Multi-Shot](https://www.wowhead.com/forever/spell=2643).
7. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=13555) when [Serpent Sting](https://www.wowhead.com/forever/spell=13555) DoT time remaining ≤ 3s.
8. `{"condition":{"cmp":{"op":"OpGt","lhs":{"autoTimeToNext":{"autoType":"Ranged"}},"rhs":{"const":{"val":"1.5s"}}}},"move":{"rangeFromTarget":{"const":{"val":"5"}}}}` when `{"autoTimeToNext":{"autoType":"Ranged"}}` > 1.5s.
9. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305) when (`{"spellCanCast":{"spellId":{"spellId":14268}}}` AND [Immolation Trap](https://www.wowhead.com/forever/spell=14305) DoT time remaining ≤ 0s).
10. Cast [Wing Clip](https://www.wowhead.com/forever/spell=14268) when Mana fraction > 10%.
11. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
12. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
13. `{"condition":{"and":{"vals":[{"not":{"val":{"auraIsActive":{"auraId":{"spellId":14266}}}}},{"cmp":{"op":"OpGt","lhs":{"autoTimeToNext":{"autoType":"Melee"}},"rhs":{"const":{"val":"0.5s"}}}}]}},"move":{"rangeFromTarget":{"const":{"val":"12"}}}}` when (NOT [Raptor Strike](https://www.wowhead.com/forever/spell=14266) active AND `{"autoTimeToNext":{"autoType":"Melee"}}` > 0.5s).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| Shoot | 343.59 |
| Cat: Auto-attack (tag 1) | 229.49 |
| [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | 98.56 |
| Auto-attack (tag 1) | 82.87 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | 80.14 |
| Auto-attack (tag 2) | 70.84 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 64.27 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 52.54 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +14157.0 |
| Mana | [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | -8057.8 |
| Mana | OtherActionManaRegen (tag 1) | +6642.8 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -4671.5 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | -3931.5 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -3710.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3593.7 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -3436.0 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -2794.0 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -2453.0 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -291.4 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003231220301351-0550005003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 1106.79 | 0.40 | 0.00 |
| Troll | 1085.12 | 0.40 | 0.00 |
| Undead | 1108.20 | 0.40 | 0.00 |
| Human | 1090.22 | 0.39 | 0.00 |
| Gnome | 1087.15 | 0.40 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Spell power / crit / hit (matched cost) — Neck [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Shadow Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -25s: [Summon Imp](https://www.wowhead.com/forever/spell=688).
- -18s: [Demonic Sacrifice](https://www.wowhead.com/forever/spell=18788).
- -16s: [Summon Succubus](https://www.wowhead.com/forever/spell=712).
- -5s: [Soul Link](https://www.wowhead.com/forever/spell=19028).
- -2.3s: [Shadow Bolt](https://www.wowhead.com/forever/spell=11661).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 80%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 70%).
3. Use ready automatic cooldowns.
4. Cast [Life Tap](https://www.wowhead.com/forever/spell=11689) when Mana fraction ≤ 1%.
5. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
6. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND Time remaining ≥ 61s).
7. Cast [Corruption](https://www.wowhead.com/forever/spell=11672) when ([Corruption](https://www.wowhead.com/forever/spell=11672) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11672}}}` AND Time remaining ≥ 12s).
8. Cast [Searing Pain](https://www.wowhead.com/forever/spell=17923) when (NOT [Demonic Brand](https://www.wowhead.com/forever/spell=1293696) active OR [Demonic Brand](https://www.wowhead.com/forever/spell=1293696) stacks ≤ 0).
9. Cast [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) when Time remaining ≤ 15s.
10. Cast [Bane of Agony](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Agony](https://www.wowhead.com/forever/spell=11713) DoT active AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND Time remaining ≥ 16s).
11. Cast [Immolate](https://www.wowhead.com/forever/spell=11668) when ([Immolate](https://www.wowhead.com/forever/spell=11668) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11668}}}` AND Time remaining ≥ 20s).
12. Cast [Searing Pain](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 2s.
13. Cast [Shadow Bolt](https://www.wowhead.com/forever/spell=695) when Mana fraction ≤ 80%.
14. Cast [Shadow Bolt](https://www.wowhead.com/forever/spell=11661).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=695) | 443.33 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 120.11 |
| [Corruption](https://www.wowhead.com/forever/spell=11672) | 119.88 |
| Succubus: [Demonic Brand](https://www.wowhead.com/forever/spell=1293697) | 91.21 |
| [Immolate](https://www.wowhead.com/forever/spell=11668) | 89.69 |
| Succubus: Auto-attack (tag 1) | 76.39 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 60.98 |
| [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) | 50.36 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Immolate](https://www.wowhead.com/forever/spell=11668) | -6245.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5358.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5258.2 |
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=695) | -4502.6 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=11672) | -4357.9 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -4196.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3597.8 |
| Mana | OtherActionManaRegen (tag 1) | +3033.6 |
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) | -2774.9 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Summon Imp](https://www.wowhead.com/forever/spell=688) | -823.8 |
| Mana | [Summon Succubus](https://www.wowhead.com/forever/spell=712) | -823.8 |

## Hunter — Marksmanship

**Talents:** 7/37/7 · `502-0053552001523051-50002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Ranged only: Auto Shot from 12 yards and never in melee. Aimed Shot, Arcane Shot and Sniper Shot on cooldown, Serpent Sting kept up, Multi-Shot only when Aimed Shot cannot be cast (they share a cooldown), Rapid Fire and Thorium Grenades; an Immolation Trap is laid before the pull. Cat pet. Trueshot Aura is taken, but the benchmark raid already has it, so it adds no DPS here.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 1039.36 | 0.39 | 3.34 |
| Tauren | 1022.21 | 0.38 | 3.77 |
| Troll | 1029.79 | 0.34 | 0.00 |
| Windshaper | 1040.54 | 0.39 | 4.08 |
| Human | 1038.55 | 0.39 | 1.91 |
| Dwarf | 1024.09 | 0.39 | 3.67 |
| Night Elf | 1041.59 | 0.39 | 3.38 |
| High Order | 1040.54 | 0.39 | 4.08 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Night Elf

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
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 90–136 dmg @ 2.4s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Off hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 90–136 dmg @ 2.4s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 147–222 dmg @ 3s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -10s: [Aspect of the Hawk](https://www.wowhead.com/forever/spell=14322).
- -31s: [Immolation Trap](https://www.wowhead.com/forever/spell=14305).
- 0.05s − `{"spellCastTime":{"spellId":{"spellId":20904}}}` + [Aimed Shot](https://www.wowhead.com/forever/spell=20904) travel time: [Aimed Shot](https://www.wowhead.com/forever/spell=20904).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Rapid Fire](https://www.wowhead.com/forever/spell=3045).
3. Cast [Aimed Shot](https://www.wowhead.com/forever/spell=20904).
4. Cast [Multi-Shot](https://www.wowhead.com/forever/spell=2643).
5. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=13555) when ([Serpent Sting](https://www.wowhead.com/forever/spell=13555) DoT time remaining ≤ 0.5s AND Time remaining ≥ 8s).
6. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287).
7. Cast [Sniper Shot](https://www.wowhead.com/forever/spell=1310786).

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| Shoot | 424.68 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 179.41 |
| Cat: Auto-attack (tag 1) | 123.92 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 110.62 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | 87.94 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 67.50 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 31.68 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 9.37 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -11323.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7553.4 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -6958.0 |
| Mana | OtherActionManaRegen (tag 1) | +5999.9 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -4728.1 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=13555) | -3741.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3616.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.5 |
| Mana | OtherActionManaRegen (tag 2) | +816.2 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -245.0 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Multi-Shot](https://www.wowhead.com/forever/spell=2643) | -123.6 |

## Warrior — Fury

**Talents:** 20/31/0 · `30305213102-25053005112010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Level-60/off-hand rage and some extra-swing interactions remain provisional. The queued-Heroic-Strike off-hand hit bug is removed. Deep Wounds uses the reported intended no-AP rollover model; live delivery and off-hand payload attribution remain qualified. Hamstring is woven for Windfury Totem procs when Bloodthirst and Whirlwind are at least 1 s away and the next main-hand swing at least 0.7 s away, outside execute (wowsims/forever #243). Spearing Strike, Overpower and Rend are cast in Battle Stance after a stance swap; the rage kept through Tactical Mastery follows the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 1007.42 | 0.49 | 0.00 |
| Tauren | 1004.32 | 0.50 | 0.00 |
| Troll | 1006.96 | 0.49 | 0.00 |
| Undead | 1015.60 | 0.49 | 0.00 |
| Windshaper | 1015.63 | 0.51 | 0.00 |
| Human | 1004.91 | 0.48 | 0.00 |
| Dwarf | 1011.03 | 0.51 | 0.00 |
| Night Elf | 1006.61 | 0.49 | 0.00 |
| Gnome | 1007.28 | 0.50 | 0.00 |
| High Order | 1015.63 | 0.51 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

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

### Before the pull

- -10s: [Bloodrage](https://www.wowhead.com/forever/spell=2687).
- -1.5s: [Death Wish](https://www.wowhead.com/forever/spell=12328).

### Rotation priorities

1. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719) when Time remaining ≤ 300s.
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when Rage ≤ 40.
3. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=11567) when (Rage ≥ 42 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
4. Cast [Bloodthirst](https://www.wowhead.com/forever/spell=23894).
5. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
6. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when NOT ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active OR (NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND Rage ≤ 75)).
7. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
8. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
9. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when (NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND Rage ≤ 30).
10. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active AND Rage ≤ 35 AND `{"spellTimeToReady":{"spellId":{"spellId":11585}}}` ≤ 1s).
11. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active AND Rage ≤ 35 AND `{"spellTimeToReady":{"spellId":{"spellId":1310222}}}` ≤ 1s).
12. Cast [Spearing Strike](https://www.wowhead.com/forever/spell=1310222).
13. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when ([Battle Stance](https://www.wowhead.com/forever/spell=2457) active AND NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
14. Use ready automatic cooldowns.
15. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
16. Cast [Hamstring](https://www.wowhead.com/forever/spell=7373) when (`{"auraIsKnown":{"auraId":{"spellId":10610}}}` AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND `{"spellTimeToReady":{"spellId":{"spellId":23894}}}` ≥ 1s AND `{"spellTimeToReady":{"spellId":{"spellId":1680}}}` ≥ 1s AND `{"autoTimeToNext":{"autoType":"MainHand"}}` ≥ 0.7s).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 157.10 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 153.13 |
| Auto-attack (tag 1) | 146.00 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=11567) | 109.33 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 87.03 |
| Auto-attack (tag 3) | 84.48 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 81.44 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 70.29 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1651.3 |
| Rage | Auto-attack (tag 2) | +1335.2 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -1244.0 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -658.1 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -444.7 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=11567) | -436.7 |
| Rage | [Hamstring](https://www.wowhead.com/forever/spell=7373) | -425.4 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +193.0 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +171.6 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -137.1 |
| Rage | OtherActionRefund | +135.3 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -134.2 |

## Paladin — Retribution

**Talents:** 16/2/33 · `55003003-2-05225331001330311`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Separate seal Echoes can coexist and fire on a landed white swing; their simultaneous server behavior remains qualified (PAL-002). Active presets use maximum registered ranks, not deliberate downranking.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 1008.36 | 0.53 | 0.93 |
| Human | 1000.68 | 0.54 | 0.64 |
| Dwarf | 991.75 | 0.54 | 0.95 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / agility / crit / hit — Wrists (Plate) [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Strength / crit / hit — Legs (Plate) [+18 Sta, +21 Crit, +21 Str, +23 Hit, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Spell two-hand — Two-hand [+25 Sta, +24 Int, +24 SP, +15 Str, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Projected: Libram of Law — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001035) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -3s: [Spell 20920](https://www.wowhead.com/forever/spell=20920).
- -1.5s: [Spell 20293](https://www.wowhead.com/forever/spell=20293).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239).
3. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271).
4. Cast [Holy Strike](https://www.wowhead.com/forever/spell=10333).
5. Cast [Spell 20293](https://www.wowhead.com/forever/spell=20293) when Seal time remaining ≤ 1s.
6. Cast [Spell 20293](https://www.wowhead.com/forever/spell=20293) when (Mana fraction ≥ 0% AND [Spell 20920](https://www.wowhead.com/forever/spell=20920) active AND `{"autoTimeToNext":{"autoType":"MainHand"}}` ≤ 3s).
7. Cast [Consecration](https://www.wowhead.com/forever/spell=20924) when Mana fraction ≥ 10%.
8. Cast [Spell 20920](https://www.wowhead.com/forever/spell=20920) when (Mana fraction ≥ 0% AND NOT [Echo of Command](https://www.wowhead.com/forever/spell=1311703) active AND [Spell 20293](https://www.wowhead.com/forever/spell=20293) active).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 272.59 |
| Auto-attack (tag 1) | 209.02 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 95.56 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 94.21 |
| Auto-attack (tag 3) | 90.88 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 77.60 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 66.24 |
| [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | 34.80 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +10078.3 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -9291.6 |
| Mana | [Spell 20920](https://www.wowhead.com/forever/spell=20920) | -9054.6 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -8556.4 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +6212.6 |
| Mana | OtherActionManaRegen (tag 1) | +5192.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3593.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3554.3 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3212.7 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -2242.4 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -516.7 |
| Mana | OtherActionManaRegen (tag 2) | +2.5 |

## Mage — Fire

**Talents:** 18/31/2 · `03022050021003-23450000130135031-002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 989.40 | 0.48 | 0.37 |
| Troll | 987.47 | 0.48 | 0.42 |
| Undead | 1000.25 | 0.48 | 0.40 |
| Human | 998.35 | 0.48 | 0.21 |
| Gnome | 991.86 | 0.50 | 0.24 |
| High Order | 975.51 | 0.48 | 0.23 |

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
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
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

### Before the pull

- -6.05s: [Pyroblast](https://www.wowhead.com/forever/spell=18809).

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 0%.
2. Use ready automatic cooldowns.
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 5% AND Time remaining > 25s).
4. Cast [Fire Blast (rank 7)](https://www.wowhead.com/forever/spell=10199) when [Hot Streak](https://www.wowhead.com/forever/spell=44445) stacks ≤ 2.
5. Cast [Pyroblast (rank 8)](https://www.wowhead.com/forever/spell=18809) when [Hot Streak](https://www.wowhead.com/forever/spell=44445) stacks = 3.
6. Cast [Fireball](https://www.wowhead.com/forever/spell=10151) when [Clearcasting](https://www.wowhead.com/forever/spell=12536) active.
7. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451) when (Mana ≥ Time remaining × 3 AND [Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks ≤ 0).
8. Cast [Arcane Missiles](https://www.wowhead.com/forever/spell=10212) when [Missile Barrage](https://www.wowhead.com/forever/spell=44404) active.
9. Cast [Scorch](https://www.wowhead.com/forever/spell=10207).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 300.60 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 196.60 |
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 187.94 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 141.46 |
| [Fireball](https://www.wowhead.com/forever/spell=10151) | 58.99 |
| [Arcane Missiles](https://www.wowhead.com/forever/spell=10212) | 57.12 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 33.22 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 14.66 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -14691.3 |
| Mana | OtherActionManaRegen (tag 1) | +13987.0 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -12210.7 |
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -10183.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +6212.2 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +5319.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4553.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3589.1 |
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -2811.9 |
| Mana | OtherActionManaRegen (tag 2) | +1826.8 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Warrior — Fury (Sunder)

**Talents:** 20/31/0 · `30305213102-25053005112010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies. Hamstring is woven for Windfury Totem procs when Bloodthirst and Whirlwind are at least 1 s away and the next main-hand swing at least 0.7 s away, outside execute (wowsims/forever #243). Spearing Strike, Overpower and Rend are cast in Battle Stance after a stance swap; the rage kept through Tactical Mastery follows the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 972.37 | 0.51 | 0.00 |
| Tauren | 966.40 | 0.51 | 0.00 |
| Troll | 959.50 | 0.50 | 0.00 |
| Undead | 983.09 | 0.50 | 0.00 |
| Windshaper | 978.53 | 0.49 | 0.00 |
| Human | 976.83 | 0.50 | 0.00 |
| Dwarf | 964.69 | 0.50 | 0.00 |
| Night Elf | 971.62 | 0.49 | 0.00 |
| Gnome | 970.65 | 0.50 | 0.00 |
| High Order | 978.53 | 0.49 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Strength / agility / crit / hit — Shoulders (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Strength / agility / crit / hit — Chest (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / crit / hit — Wrists (Plate) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Strength / crit / hit — Hands (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / crit / hit — Waist (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Rapidity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -10s: [Bloodrage](https://www.wowhead.com/forever/spell=2687).
- -1.5s: [Death Wish](https://www.wowhead.com/forever/spell=12328).

### Rotation priorities

1. Cast [Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) when ([Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) stacks < 5 OR [Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) time remaining < 8s).
2. `{"condition":{"and":{"vals":[{"or":{"vals":[{"cmp":{"op":"OpLt","lhs":{"auraNumStacks":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597,"rank":5}}},"rhs":{"const":{"val":"5"}}}},{"cmp":{"op":"OpLt","lhs":{"auraRemainingTime":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597,"rank":5}}},"rhs":{"const":{"val":"8s"}}}}]}},{"cmp":{"op":"OpLt","lhs":{"currentRage":{}},"rhs":{"const":{"val":"15"}}}}]}},"wait":{"duration":{"const":{"val":"500ms"}}}}` when (([Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) stacks < 5 OR [Sunder Armor (rank 5)](https://www.wowhead.com/forever/spell=11597) time remaining < 8s) AND Rage < 15).
3. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719) when Time remaining ≤ 300s.
4. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when Rage ≤ 40.
5. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=11567) when (Rage ≥ 42 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
6. Cast [Bloodthirst](https://www.wowhead.com/forever/spell=23894).
7. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
8. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when NOT ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active OR (NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND Rage ≤ 75)).
9. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
10. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
11. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when (NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND Rage ≤ 30).
12. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active AND Rage ≤ 35 AND `{"spellTimeToReady":{"spellId":{"spellId":11585}}}` ≤ 1s).
13. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active AND Rage ≤ 35 AND `{"spellTimeToReady":{"spellId":{"spellId":1310222}}}` ≤ 1s).
14. Cast [Spearing Strike](https://www.wowhead.com/forever/spell=1310222).
15. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when ([Battle Stance](https://www.wowhead.com/forever/spell=2457) active AND NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
16. Use ready automatic cooldowns.
17. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
18. Cast [Hamstring](https://www.wowhead.com/forever/spell=7373) when (`{"auraIsKnown":{"auraId":{"spellId":10610}}}` AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND `{"spellTimeToReady":{"spellId":{"spellId":23894}}}` ≥ 1s AND `{"spellTimeToReady":{"spellId":{"spellId":1680}}}` ≥ 1s AND `{"autoTimeToNext":{"autoType":"MainHand"}}` ≥ 0.7s).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 153.41 |
| Auto-attack (tag 1) | 146.33 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 142.34 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=11567) | 96.38 |
| Auto-attack (tag 3) | 83.74 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 78.95 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 78.32 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 65.57 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1649.1 |
| Rage | Auto-attack (tag 2) | +1307.5 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -1166.9 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -608.5 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -407.5 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=11567) | -395.3 |
| Rage | [Hamstring](https://www.wowhead.com/forever/spell=7373) | -344.0 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -273.2 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +192.5 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +176.3 |
| Rage | OtherActionRefund | +138.5 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -132.6 |

## Paladin — Physical Ret

**Talents:** 16/1/34 · `55003003-1-05225331001330321`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Uses a physical equipment budget without caster gear. Champion of the Light converts existing Intellect into damage-only spell bonus, not healing power. Seal Echo ordering remains qualified (PAL-002).


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 940.28 | 0.54 | 3.75 |
| Human | 937.93 | 0.54 | 0.28 |
| Dwarf | 927.35 | 0.53 | 0.39 |

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
| Ranged/relic | [Projected: Libram of Law — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001035) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -5s: [Spell 20293](https://www.wowhead.com/forever/spell=20293).
- -3s: [Spell 20920](https://www.wowhead.com/forever/spell=20920).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239).
3. Cast [Spell 20920](https://www.wowhead.com/forever/spell=20920) when (Mana fraction ≥ 0% AND NOT [Echo of Command](https://www.wowhead.com/forever/spell=1311703) active AND [Spell 20293](https://www.wowhead.com/forever/spell=20293) active AND `{"spellTimeToReady":{"spellId":{"spellId":20271}}}` ≤ 5.9s).
4. Cast [Holy Strike](https://www.wowhead.com/forever/spell=10333).
5. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271) when [Spell 20920](https://www.wowhead.com/forever/spell=20920) active.
6. Cast [Spell 20293](https://www.wowhead.com/forever/spell=20293) when Seal time remaining ≤ 1s.
7. Cast [Spell 20293](https://www.wowhead.com/forever/spell=20293) when (Mana fraction ≥ 0% AND [Spell 20920](https://www.wowhead.com/forever/spell=20920) active AND `{"spellTimeToReady":{"spellId":{"spellId":20271}}}` > 0.5s).
8. Cast [Consecration](https://www.wowhead.com/forever/spell=20924) when Mana fraction ≥ 10%.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 276.64 |
| Auto-attack (tag 1) | 221.66 |
| Auto-attack (tag 3) | 109.11 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 86.89 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 76.25 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20966) | 68.78 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 44.33 |
| [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | 29.01 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +10155.3 |
| Mana | [Spell 20920](https://www.wowhead.com/forever/spell=20920) | -9573.6 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -9198.5 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -7752.6 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +5765.5 |
| Mana | OtherActionManaRegen (tag 1) | +5220.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3606.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3560.7 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -2874.6 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -2191.5 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -503.8 |
| Mana | OtherActionManaRegen (tag 2) | +11.6 |

## Warlock — Destruction

**Talents:** 20/0/31 · `25500000035--055030510310035`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 938.36 | 0.44 | 1.22 |
| Troll | 928.52 | 0.44 | 0.99 |
| Undead | 932.16 | 0.43 | 1.24 |
| Human | 934.03 | 0.43 | 1.18 |
| Gnome | 923.83 | 0.42 | 0.96 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Shadow Power |
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Intellect / spell power / crit — Trinket [+11 Sta, +14 Int, +18 SP, +9 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4s: [Soul Fire](https://www.wowhead.com/forever/spell=17924).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 70%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 80%).
3. Use ready automatic cooldowns.
4. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871) when Time remaining ≤ 1.5.
5. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
6. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
7. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=11668) when [Immolate (rank 8)](https://www.wowhead.com/forever/spell=11668) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11668,"rank":8}}}`.
8. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=11672) when [Corruption (rank 7)](https://www.wowhead.com/forever/spell=11672) DoT time remaining ≤ 0s.
9. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
10. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 5%.
11. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`).
12. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871) when Time remaining ≥ 15s.
13. Cast [Conflagrate](https://www.wowhead.com/forever/spell=18932) when [Immolate](https://www.wowhead.com/forever/spell=11668) DoT time remaining ≤ 15s.
14. Cast [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) when Mana fraction ≥ 60%.
15. Cast [Shadow Bolt](https://www.wowhead.com/forever/spell=695).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=695) | 261.27 |
| [Corruption](https://www.wowhead.com/forever/spell=11672) | 125.42 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 116.51 |
| [Immolate](https://www.wowhead.com/forever/spell=11668) | 105.42 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 93.97 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 73.44 |
| Succubus: Auto-attack (tag 1) | 69.53 |
| [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) | 52.85 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +6567.0 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=11668) | -6338.5 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -6144.5 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -5519.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5398.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5103.2 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=11672) | -4650.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3592.9 |
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) | -3116.6 |
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=695) | -2924.8 |
| Mana | OtherActionManaRegen (tag 1) | +2880.0 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |

## Mage — Frostfire

**Talents:** 0/29/22 · `-0355100013013403-2035010311013002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Frostfire Bolt spam behind five Fire Vulnerability stacks, Pyroblast on three Heating Up stacks. Uses the hotfixed Frost tree (20 October 2026).


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 923.78 | 0.71 | 0.01 |
| Troll | 917.88 | 0.70 | 0.01 |
| Undead | 932.69 | 0.70 | 0.01 |
| Human | 937.84 | 0.68 | 0.00 |
| Gnome | 916.08 | 0.65 | 0.00 |
| High Order | 922.47 | 0.75 | 0.02 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / fire power / crit — Shoulders (Cloth) [+14 Sta, +18 Int, +20 Fire, +14 Crit, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / fire power / crit — Back (Cloth) [+11 Sta, +14 Int, +15 Fire, +10 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / fire power / crit — Chest (Cloth) [+18 Sta, +24 Int, +26 Fire, +18 Crit, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / fire power / crit — Wrists (Cloth) [+11 Sta, +14 Int, +15 Fire, +10 Crit, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Fire Power |
| Waist | [Modeled: Intellect / fire power / crit — Waist (Cloth) [+14 Sta, +18 Int, +20 Fire, +14 Crit, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279266) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- 0.05s − `{"spellCastTime":{"spellId":{"spellId":18809}}}` + [Pyroblast](https://www.wowhead.com/forever/spell=18809) travel time: [Pyroblast](https://www.wowhead.com/forever/spell=18809).

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 0%.
2. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 35% AND Time remaining > 25s).
3. Cast [Pyroblast (rank 8)](https://www.wowhead.com/forever/spell=18809) when ([Hot Streak](https://www.wowhead.com/forever/spell=44445) stacks = 3 AND ([Pyroblast](https://www.wowhead.com/forever/spell=18809) DoT time remaining ≤ 2.5s OR Time remaining ≤ 5s)).
4. Use ready automatic cooldowns.
5. Cast [Ice Lance](https://www.wowhead.com/forever/spell=30455) when Time remaining ≤ 3s.
6. Cast [Scorch](https://www.wowhead.com/forever/spell=2948) when ([Improved Scorch](https://www.wowhead.com/forever/spell=12873) stacks < 5 OR [Improved Scorch](https://www.wowhead.com/forever/spell=12873) time remaining ≤ 6s).
7. Cast [Fire Blast](https://www.wowhead.com/forever/spell=2136) when (NOT [Fingers of Frost](https://www.wowhead.com/forever/spell=44543) active AND [Hot Streak](https://www.wowhead.com/forever/spell=44445) stacks ≥ 1 AND [Hot Streak](https://www.wowhead.com/forever/spell=44445) stacks < 2).
8. Cast [Frostfire Bolt](https://www.wowhead.com/forever/spell=1237313).
9. Cast [Scorch](https://www.wowhead.com/forever/spell=2948).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Frostfire Bolt](https://www.wowhead.com/forever/spell=1237313) | 498.19 |
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 199.42 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 180.83 |
| [Scorch](https://www.wowhead.com/forever/spell=2948) | 31.85 |
| [Fire Blast](https://www.wowhead.com/forever/spell=2136) | 17.68 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 7.85 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 2.02 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostfire Bolt](https://www.wowhead.com/forever/spell=1237313) | -27600.8 |
| Mana | OtherActionManaRegen (tag 1) | +10565.1 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -8697.6 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +6770.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4805.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4040.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.8 |
| Mana | OtherActionManaRegen (tag 2) | +1719.7 |
| Mana | [Scorch](https://www.wowhead.com/forever/spell=2948) | -899.4 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=2136) | -416.7 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -240.1 |

## Hunter — Pet/Melee

**Talents:** 16/11/24 · `53200005001-0050051-5302002300501201`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Two independent Hawks use provisional high-rank damage and crit calibration (HUN-004); no owner-crit or unverified Lone Wolf coexistence is assumed. Tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 917.01 | 0.35 | 0.00 |
| Windshaper | 932.75 | 0.35 | 0.00 |
| Orc | 935.04 | 0.36 | 0.00 |
| Human | 930.76 | 0.36 | 0.00 |
| Night Elf | 931.88 | 0.35 | 0.00 |
| Dwarf | 921.96 | 0.35 | 0.00 |
| Troll | 926.69 | 0.35 | 0.00 |
| High Order | 932.75 | 0.35 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Orc

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
| Trinket 1 | [Modeled: Passive Strength / agility / crit / hit — Trinket [+17 Str, +11 Sta, +9 Crit, +6 Hit, +11 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Axe one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272592) | 65 | Enchant Weapon - Crusader |
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

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 165.18 |
| Cat: Auto-attack (tag 1) | 146.28 |
| Auto-attack (tag 1) | 129.46 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 102.05 |
| Auto-attack (tag 3) | 76.10 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 73.27 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 59.65 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 41.19 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +15323.5 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -7885.3 |
| Mana | OtherActionManaRegen (tag 1) | +6351.1 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -4469.0 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3306.2 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -3008.2 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -1620.3 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -1543.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +259.9 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) | -180.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Warlock — Affliction

**Talents:** 30/0/21 · `2550002003520015--0550305003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 926.27 | 0.50 | 0.08 |
| Troll | 917.56 | 0.49 | 0.07 |
| Undead | 929.51 | 0.50 | 0.08 |
| Human | 926.36 | 0.50 | 0.07 |
| Gnome | 914.29 | 0.49 | 0.04 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Crit / shadow power / intellect — Chest (Cloth) [+18 Sta, +26 Shadow, +24 Crit, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Gloves - Shadow Power |
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Crit / shadow power / intellect — Legs (Cloth) [+18 Sta, +26 Shadow, +24 Crit, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Intellect / spell power / crit — Trinket [+11 Sta, +14 Int, +18 SP, +9 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -4.3s: [Soul Fire](https://www.wowhead.com/forever/spell=17924).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 65%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 80%).
3. Use ready automatic cooldowns.
4. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
5. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 6%.
6. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
7. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=11672) when ([Corruption (rank 7)](https://www.wowhead.com/forever/spell=11672) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11672,"rank":7}}}` AND Time remaining ≥ 4s).
8. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
9. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when ((NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`) AND Time remaining ≥ 12s).
10. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=11668) when ([Immolate (rank 8)](https://www.wowhead.com/forever/spell=11668) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11668,"rank":8}}}` AND Time remaining ≥ 12s).
11. Cast [Shadow Bolt (rank 9)](https://www.wowhead.com/forever/spell=11661).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) | 487.86 |
| [Corruption](https://www.wowhead.com/forever/spell=11672) | 124.86 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 112.29 |
| [Immolate](https://www.wowhead.com/forever/spell=11668) | 83.17 |
| Succubus: Auto-attack (tag 1) | 68.32 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 21.00 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 15.64 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 8.81 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) | -28664.1 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +19174.0 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=11668) | -5915.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5398.2 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=11672) | -4640.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3659.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3597.7 |
| Mana | OtherActionManaRegen (tag 1) | +2892.3 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -430.0 |
| Mana | [Soul Fire](https://www.wowhead.com/forever/spell=17924) | -301.5 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -160.7 |

## Priest — Shadow

**Talents:** 18/0/33 · `50500500102--50542001201302251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 891.65 | 0.24 | 0.00 |
| Undead | 910.35 | 0.25 | 0.00 |
| Human | 919.51 | 0.25 | 0.00 |
| Dwarf | 913.47 | 0.25 | 0.00 |
| Night Elf | 918.57 | 0.25 | 0.00 |
| Gnome | 918.02 | 0.25 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Crit / shadow power / intellect — Neck [+11 Sta, +15 Shadow, +14 Crit, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
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
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -3s: [Shadowform](https://www.wowhead.com/forever/spell=15473).
- 0s − `{"spellCastTime":{"spellId":{"spellId":10947}}}` + [Mind Blast](https://www.wowhead.com/forever/spell=10947) travel time: [Mind Blast](https://www.wowhead.com/forever/spell=10947).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636).
3. `{"condition":{},"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10947}}}]}}` when `{}`.
4. Cast [Mind Blast (rank 9)](https://www.wowhead.com/forever/spell=10947).
5. Cast [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) when ([Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) DoT time remaining ≤ 3s AND Time remaining ≥ 9s).
6. Cast [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) when ([Devouring Plague](https://www.wowhead.com/forever/spell=19280) DoT time remaining ≤ 5s AND Time remaining ≥ 12s).
7. Cast [Mind Flay (rank 6)](https://www.wowhead.com/forever/spell=18807).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 359.67 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 185.34 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 172.67 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 112.45 |
| [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | 89.38 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +6386.0 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -6252.5 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -6125.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4155.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3596.9 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -3055.0 |
| Mana | [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | -2890.0 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |

## Warrior — Arms

**Talents:** 36/13/2 · `20305213132515201-3505-2`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 888.19 | 0.63 | 0.00 |
| Tauren | 880.24 | 0.63 | 0.00 |
| Troll | 897.05 | 0.62 | 0.00 |
| Undead | 908.63 | 0.63 | 0.00 |
| Windshaper | 887.31 | 0.61 | 0.00 |
| Human | 901.88 | 0.62 | 0.00 |
| Dwarf | 878.78 | 0.63 | 0.00 |
| Night Elf | 885.09 | 0.62 | 0.00 |
| Gnome | 899.04 | 0.62 | 0.00 |
| High Order | 887.31 | 0.61 | 0.00 |

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

### Before the pull

- -1.9s: [Recklessness](https://www.wowhead.com/forever/spell=1719).
- -0.5s: [Battle Stance](https://www.wowhead.com/forever/spell=2457).
- -5s: [Bloodrage](https://www.wowhead.com/forever/spell=2687).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Rend](https://www.wowhead.com/forever/spell=11574) DoT time remaining ≤ 1.5s OR (([Bloodthrill](https://www.wowhead.com/forever/spell=56638) active OR [Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active) AND `{"spellTimeToReady":{"spellId":{"spellId":11585}}}` ≤ 1.5s AND Rage ≤ 45)).
3. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when ([Rend](https://www.wowhead.com/forever/spell=11574) DoT time remaining > 6s AND NOT ([Bloodthrill](https://www.wowhead.com/forever/spell=56638) active OR [Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active) AND Rage ≤ 26).
4. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719) when [Berserker Stance](https://www.wowhead.com/forever/spell=2458) active.
5. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=11567) when (Rage ≥ 85 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
6. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
7. Cast [Mortal Strike](https://www.wowhead.com/forever/spell=21553).
8. Cast [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) when Rage ≥ 15.
9. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 15.
10. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when [Rend](https://www.wowhead.com/forever/spell=11574) DoT time remaining ≤ 3s.
11. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
12. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
13. Cast [Hamstring](https://www.wowhead.com/forever/spell=7373) when Rage ≥ 20.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 159.85 |
| Auto-attack (tag 3) | 154.71 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 134.50 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 128.87 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 61.32 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 58.86 |
| [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | 43.18 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 38.32 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +2599.2 |
| Rage | [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | -1152.2 |
| Rage | [Hamstring](https://www.wowhead.com/forever/spell=7373) | -433.6 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -334.7 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -297.8 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -253.2 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -189.9 |
| Rage | [Battle Stance](https://www.wowhead.com/forever/spell=2457) | -185.4 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +164.4 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +155.0 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -147.6 |
| Rage | [Rend](https://www.wowhead.com/forever/spell=11574) | -146.8 |

## Warrior — 2H Bloodthirst

**Talents:** 20/31/0 · `20305203032-45050020152010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model. Hamstring is woven for Windfury Totem procs when Bloodthirst and Whirlwind are at least 1.5 s away, outside execute (wowsims/forever #243).


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 890.99 | 0.63 | 0.00 |
| Tauren | 881.34 | 0.62 | 0.00 |
| Undead | 895.55 | 0.63 | 0.00 |
| Troll | 882.99 | 0.62 | 0.00 |
| Windshaper | 886.08 | 0.63 | 0.00 |
| Human | 886.82 | 0.63 | 0.00 |
| Dwarf | 880.27 | 0.62 | 0.00 |
| Gnome | 886.67 | 0.63 | 0.00 |
| Night Elf | 887.35 | 0.62 | 0.00 |
| High Order | 886.08 | 0.63 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

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

### Before the pull

- -10s: [Bloodrage](https://www.wowhead.com/forever/spell=2687).
- -2.5s: [Death Wish](https://www.wowhead.com/forever/spell=12328).
- 0s: [Recklessness](https://www.wowhead.com/forever/spell=1719).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when Rage ≤ 60.
3. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719).
4. Cast [Bloodthirst](https://www.wowhead.com/forever/spell=23894).
5. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when ([Battle Stance](https://www.wowhead.com/forever/spell=2457) active AND NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
6. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
7. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when NOT [Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active.
8. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
9. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
10. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
11. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower opportunity](https://www.wowhead.com/forever/spell=1282733) active AND Rage ≤ 80).
12. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when (Rage ≥ 50 AND `{"autoTimeToNext":{"autoType":"MainHand"}}` ≥ 2.25s).
13. Cast [Hamstring](https://www.wowhead.com/forever/spell=7373) when (`{"auraIsKnown":{"auraId":{"spellId":10610}}}` AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND `{"autoTimeToNext":{"autoType":"MainHand"}}` ≥ 1s AND Rage ≥ 0).
14. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=11567) when (Rage ≥ 95 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 230.51 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 162.93 |
| Auto-attack (tag 3) | 138.63 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 71.29 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 71.17 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 62.84 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 54.98 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 24.94 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +2970.7 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -1223.0 |
| Rage | [Hamstring](https://www.wowhead.com/forever/spell=7373) | -675.0 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -516.3 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -431.7 |
| Rage | [Battle Stance](https://www.wowhead.com/forever/spell=2457) | -173.4 |
| Rage | OtherActionRefund | +117.9 |
| Rage | [Berserker Stance](https://www.wowhead.com/forever/spell=2458) | -116.7 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +115.2 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +113.7 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -97.2 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -78.6 |

## Paladin — Shockadin

**Talents:** 23/0/28 · `550030032101012--052153300012303`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Holy Shock Retribution with Seal of Righteousness, Judgement, Consecration, Hammer of Wrath and Holy Strike. Champion of the Light at the hotfixed 20/40/60% of Intellect. No Manual Crowd Pummeler swap.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 879.80 | 0.33 | 0.00 |
| Human | 870.17 | 0.34 | 0.00 |
| Dwarf | 864.21 | 0.34 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Plate) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Plate) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 553 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Plate) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 738 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Plate) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 323 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Plate) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Plate) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Spell power / crit / hit (matched cost) — Legs (Plate) [+18 Sta, +21 Crit, +23 Hit, +23 SP, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
| Feet | [Modeled: Spell power / crit / hit (matched cost) — Feet (Plate) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Spell two-hand — Two-hand [+25 Sta, +24 Int, +24 SP, +15 Str, 137–207 dmg @ 2.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Fiery Weapon |
| Ranged/relic | [Projected: Libram of Law — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001035) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -3s: [Divine Favor](https://www.wowhead.com/forever/spell=20216).
- -1.5s: [Spell 20293](https://www.wowhead.com/forever/spell=20293).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Spell 20293](https://www.wowhead.com/forever/spell=20293) when Seal time remaining ≤ 1s.
3. Cast [Consecration](https://www.wowhead.com/forever/spell=20924) when Mana fraction ≥ 20%.
4. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271).
5. Cast [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239).
6. Cast [Holy Strike](https://www.wowhead.com/forever/spell=10333).
7. Cast [Holy Shock](https://www.wowhead.com/forever/spell=20930).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 151.02 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 141.69 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 138.82 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 105.26 |
| [Holy Shock](https://www.wowhead.com/forever/spell=20930) | 104.59 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 93.22 |
| Auto-attack (tag 3) | 39.86 |
| [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | 39.12 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -14644.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11901.9 |
| Mana | [Holy Shock](https://www.wowhead.com/forever/spell=20930) | -8482.5 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +5136.8 |
| Mana | OtherActionManaRegen (tag 1) | +5136.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3474.6 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3184.3 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2749.7 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -2720.0 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -1980.0 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -540.0 |
| Mana | OtherActionManaRegen (tag 2) | +102.0 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `254000001351-0025003001-0550005103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 861.90 | 0.51 | 0.00 |
| Troll | 852.59 | 0.51 | 0.00 |
| Undead | 874.86 | 0.52 | 0.00 |
| Human | 873.63 | 0.53 | 0.00 |
| Gnome | 863.74 | 0.51 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +12 Crit, +13 Hit, +13 SP, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit — Hands (Cloth) [+14 Sta, +18 Int, +22 SP, +11 Crit, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Gloves - Shadow Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Crit / shadow power / intellect — Finger [+11 Sta, +15 Shadow, +14 Crit, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=279269) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
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
11. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=11672) when [Corruption (rank 7)](https://www.wowhead.com/forever/spell=11672) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11672,"rank":7}}}`.
12. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=11668) when [Immolate (rank 8)](https://www.wowhead.com/forever/spell=11668) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11668,"rank":8}}}`.
13. Cast [Shadow Bolt (rank 9)](https://www.wowhead.com/forever/spell=11661).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) | 499.93 |
| [Corruption](https://www.wowhead.com/forever/spell=11672) | 125.99 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 117.30 |
| [Immolate](https://www.wowhead.com/forever/spell=11668) | 86.94 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 22.06 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 16.13 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 4.02 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.48 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=11661) | -30932.4 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +21033.5 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=11668) | -6666.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5385.7 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=11672) | -4632.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3601.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3592.7 |
| Mana | OtherActionManaRegen (tag 1) | +2877.6 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +629.6 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -522.8 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -362.4 |

## Mage — Arcane

**Talents:** 32/3/16 · `0502205003100311531-03-05350003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 843.13 | 0.48 | 0.70 |
| Troll | 830.98 | 0.48 | 1.05 |
| Undead | 847.66 | 0.48 | 0.74 |
| Human | 851.62 | 0.47 | 0.41 |
| Gnome | 852.13 | 0.47 | 0.42 |
| High Order | 851.76 | 0.48 | 0.90 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Gnome

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit — Wrists (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 2 | [Modeled: Passive Intellect / spell power / crit — Trinket [+11 Sta, +14 Int, +18 SP, +9 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -2.8s: [Frostbolt](https://www.wowhead.com/forever/spell=10181).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 0% AND Time remaining > 20s).
3. Cast [Arcane Wrath](https://www.wowhead.com/forever/spell=1326082) when [Presence of Mind](https://www.wowhead.com/forever/spell=12043) active.
4. Cast [Arcane Missiles (rank 8)](https://www.wowhead.com/forever/spell=10212) when [Missile Barrage](https://www.wowhead.com/forever/spell=44404) active.
5. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451) when [Clearcasting](https://www.wowhead.com/forever/spell=12536) active.
6. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451) when (Time remaining < 20s AND Mana fraction > 10%).
7. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451) when ([Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks < 3 AND Mana fraction > 40%).
8. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451) when ([Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks < 2 AND Mana fraction > 0%).
9. Cast [Arcane Wrath](https://www.wowhead.com/forever/spell=1326082).
10. Cast [Arcane Wrath](https://www.wowhead.com/forever/spell=1326079).
11. Cast [Frostbolt](https://www.wowhead.com/forever/spell=10181).

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Arcane Missiles](https://www.wowhead.com/forever/spell=10212) | 328.91 |
| [Arcane Wrath](https://www.wowhead.com/forever/spell=1326082) | 190.44 |
| [Frostbolt](https://www.wowhead.com/forever/spell=10181) | 175.24 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 154.77 |
| [Arcane Wrath](https://www.wowhead.com/forever/spell=1326079) | 2.77 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -20547.8 |
| Mana | OtherActionManaRegen (tag 1) | +11672.7 |
| Mana | [Arcane Wrath](https://www.wowhead.com/forever/spell=1326082) | -8019.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7482.8 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=10181) | -7037.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5059.7 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.5 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +848.9 |
| Mana | OtherActionManaRegen (tag 2) | +695.4 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +599.8 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=10212) | -498.1 |
| Mana | [Arcane Wrath](https://www.wowhead.com/forever/spell=1326079) | -51.5 |

## Druid — Feral

**Talents:** 7/34/10 · `05002-55000032021032212051-055`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Blizzard supports proc-enchant form eligibility, and Crusader's sourced health return is modeled. Its exact PPM/conditional interactions remain qualified (DRU-011); the Night Elf profile uses that enchant.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 830.89 | 0.27 | 1.48 |
| Windshaper | 841.20 | 0.27 | 0.94 |
| Night Elf | 834.24 | 0.28 | 1.30 |
| High Order | 841.20 | 0.27 | 0.94 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

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
| Ranged/relic | [Projected: Idol of the Dream — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001001) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Use ready automatic cooldowns.
3. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=22829) when (Time remaining ≤ 2s AND Combo points ≥ 3).
4. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=22829) when ([Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 5s AND (Combo points ≥ 4 OR (Combo points ≥ 4 AND `{"spellTimeToReady":{"spellId":{"spellId":1322605}}}` ≤ 4s)) AND Energy ≤ 45 AND [Shred](https://www.wowhead.com/forever/spell=9830) current cost > 0).
5. Cast [Rake](https://www.wowhead.com/forever/spell=9904) when ((NOT [Rake](https://www.wowhead.com/forever/spell=9904) DoT active AND Time remaining ≥ 3s AND Combo points < 5) AND (NOT (Combo points ≥ 5 AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining < 5s AND Time remaining ≥ 8s) OR ([Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining > 3s OR ([Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 1s AND ([Rake](https://www.wowhead.com/forever/spell=9904) current cost ≤ 0 OR [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 100 − Energy / 10.403)))) AND [Shred](https://www.wowhead.com/forever/spell=9830) current cost > 0).
6. Cast [Rip](https://www.wowhead.com/forever/spell=9896) when (Combo points ≥ 4 AND NOT [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND Time remaining ≥ 6s AND [Shred](https://www.wowhead.com/forever/spell=9830) current cost > 0).
7. Cast [Shifting Power](https://www.wowhead.com/forever/spell=1322605) when (Energy ≤ 50 AND NOT (Combo points ≥ 5 AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining < 0s AND Energy ≥ [Rip](https://www.wowhead.com/forever/spell=9896) current cost AND Time remaining ≥ 8s)).
8. Cast [Shred](https://www.wowhead.com/forever/spell=9830) when ((NOT (Combo points ≥ 5 AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining < 5s AND Time remaining ≥ 8s) OR ([Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining > 3s OR ([Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 0s AND ([Shred](https://www.wowhead.com/forever/spell=9830) current cost ≤ 0 OR [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 100 − Energy / 10.403))))).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 222.34 |
| [Shred](https://www.wowhead.com/forever/spell=9830) | 211.95 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 181.56 |
| [Rake](https://www.wowhead.com/forever/spell=9904) | 106.06 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=22829) | 63.77 |
| Auto-attack (tag 3) | 51.92 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 3.61 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shifting Power](https://www.wowhead.com/forever/spell=1322605) | -27972.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +14695.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3850.3 |
| Mana | OtherActionManaRegen (tag 2) | +3445.4 |
| Energy | OtherActionEnergyRegen | +3076.8 |
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -2944.2 |
| Energy | [Shifting Power](https://www.wowhead.com/forever/spell=1322605) | +1635.3 |
| Mana | OtherActionManaRegen (tag 1) | +1465.1 |
| Energy | [Rake](https://www.wowhead.com/forever/spell=9904) | -1053.7 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -624.1 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=22829) | -530.6 |
| Energy | OtherActionRefund | +265.0 |

## Shaman — Enhancement

**Talents:** 19/31/1 · `053503102-055030011005102251-01`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 826.12 | 0.56 | 0.25 |
| Tauren | 819.20 | 0.55 | 0.28 |
| Troll | 822.68 | 0.57 | 0.25 |
| Windshaper | 830.45 | 0.56 | 0.26 |
| Dwarf | 822.80 | 0.56 | 0.16 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / crit / hit — Neck [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit — Shoulders (Mail) [+14 Sta, +18 Int, +22 SP, +11 Crit, 312 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Strength / crit / hit — Back (Cloth) [+11 Sta, +12 Crit, +12 Str, +13 Hit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Lesser Agility |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 416 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Strength / intellect / crit / hit (mail) — Wrists (Mail) [+8 Str, +13 Int, +13 Sta, +8 Crit, +10 Hit, 182 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | Enchant Bracer - Superior Strength |
| Hands | [Modeled: Intellect / spell power / crit — Hands (Mail) [+14 Sta, +18 Int, +22 SP, +11 Crit, 260 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Gloves - Minor Haste |
| Waist | [Modeled: Strength / intellect / crit / hit (mail) — Waist (Mail) [+11 Str, +17 Int, +17 Sta, +11 Crit, +14 Hit, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | — |
| Legs | [Modeled: Strength / intellect / crit / hit (mail) — Legs (Mail) [+14 Str, +22 Int, +22 Sta, +14 Crit, +18 Hit, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / intellect / crit / hit (mail) — Feet (Mail) [+11 Str, +17 Int, +17 Sta, +11 Crit, +14 Hit, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22676) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Spell two-hand — Two-hand [+25 Sta, +24 Int, +24 SP, +15 Str, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Fiery Weapon |
| Ranged/relic | [Projected: Totem of the Storm — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001029) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -2.25s: [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Searing Totem](https://www.wowhead.com/forever/spell=10438) when (Fire totem time remaining ≤ 0.5s AND Time remaining ≥ 5s).
3. Cast [Flame Shock](https://www.wowhead.com/forever/spell=29228) when (NOT [Flame Shock](https://www.wowhead.com/forever/spell=29228) DoT active AND Time remaining ≥ 10s).
4. Cast [Stormstrike](https://www.wowhead.com/forever/spell=17364).
5. Cast [Earth Shock](https://www.wowhead.com/forever/spell=10414).
6. Cast [Fire Nova](https://www.wowhead.com/forever/spell=408345) when [Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active.
7. Cast [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) when [Maelstrom Weapon](https://www.wowhead.com/forever/spell=408505) stacks ≥ 5.
8. Cast [Fire Nova](https://www.wowhead.com/forever/spell=408345) when Mana fraction ≥ 10%.

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 224.28 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 140.54 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 101.80 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 100.69 |
| [Fire Nova](https://www.wowhead.com/forever/spell=408345) | 68.33 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 55.71 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 55.18 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 31.23 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Nova](https://www.wowhead.com/forever/spell=408345) | -12366.9 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -9488.7 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9406.5 |
| Mana | OtherActionManaRegen (tag 1) | +7071.0 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4737.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3780.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3591.9 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -2641.7 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | OtherActionManaRegen (tag 2) | +98.7 |

## Mage — Arcane–Frost

**Talents:** 29/0/22 · `050220500310030053--0535000301003002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Ice Lance uses a provisional 0.10 SP share calibrated from lower ranks. Fingers of Frost ordering and Arcane Missiles/Clearcasting proc interpretation remain unresolved; Missile Barrage now requires a landed eligible hit.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 787.37 | 0.48 | 1.21 |
| Troll | 777.92 | 0.48 | 1.55 |
| Undead | 794.73 | 0.49 | 1.14 |
| Human | 794.23 | 0.47 | 0.60 |
| Gnome | 793.11 | 0.48 | 0.67 |
| High Order | 792.40 | 0.49 | 1.45 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Undead

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit — Finger [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 1 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Trinket [+11 Sta, +9 Crit, +6 Hit, +19 SP, +11 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
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
5. Cast [Arcane Missiles](https://www.wowhead.com/forever/spell=10212) when [Missile Barrage](https://www.wowhead.com/forever/spell=44404) active.
6. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451) when ([Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks < 4 AND Mana fraction > 20%).
7. Cast [Frostbolt](https://www.wowhead.com/forever/spell=10181).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Arcane Missiles](https://www.wowhead.com/forever/spell=10212) | 304.38 |
| [Frostbolt](https://www.wowhead.com/forever/spell=10181) | 265.14 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 139.62 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 70.91 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 14.69 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -27085.3 |
| Mana | OtherActionManaRegen (tag 1) | +13698.1 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=10181) | -10015.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +7440.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5335.0 |
| Mana | OtherActionManaRegen (tag 2) | +2564.1 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -1849.1 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.1 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.8 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +520.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +159.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +0.0 |

## Druid — Balance

**Talents:** 38/0/13 · `533222011550135--505003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 777.68 | 0.33 | 0.00 |
| Windshaper | 794.44 | 0.33 | 0.00 |
| Night Elf | 786.72 | 0.33 | 0.00 |
| High Order | 794.44 | 0.33 | 0.00 |

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
| Ranged/relic | [Projected: Swarming Idol — Relic [+2 Sta, +1.75 Int, +2.25 SP]](https://www.wowhead.com/forever/item=920001026) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- 0.05s − `{"spellCastTime":{"spellId":{"spellId":9876}}}`: [Starfire](https://www.wowhead.com/forever/spell=9876).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Moonfire (rank 10)](https://www.wowhead.com/forever/spell=9835) when Time remaining ≤ 2.4s.
3. Cast [Insect Swarm (rank 5)](https://www.wowhead.com/forever/spell=24977) when ([Insect Swarm (rank 5)](https://www.wowhead.com/forever/spell=24977) DoT time remaining ≤ 1.0s AND Time remaining > 9s).
4. Cast [Moonfire (rank 10)](https://www.wowhead.com/forever/spell=9835) when (NOT [Moonfire (rank 10)](https://www.wowhead.com/forever/spell=9835) DoT active AND Time remaining > 6s).
5. Cast [Starfire (rank 7)](https://www.wowhead.com/forever/spell=9876) when [Eclipse](https://www.wowhead.com/forever/spell=48518) stacks ≥ 1.
6. Cast [Wrath (rank 8)](https://www.wowhead.com/forever/spell=9912).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Starfire](https://www.wowhead.com/forever/spell=9876) | 455.68 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 120.54 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 110.14 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 108.08 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=9876) | -16672.4 |
| Mana | OtherActionManaRegen (tag 1) | +7391.4 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -5053.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4704.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3600.4 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | +2726.2 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1729.3 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1245.2 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | -62.2 |
| Mana | OtherActionManaRegen (tag 2) | +56.4 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 736.77 | 0.30 | 0.00 |
| Troll | 734.90 | 0.30 | 0.00 |
| Undead | 746.71 | 0.30 | 0.00 |
| Windshaper | 746.80 | 0.31 | 0.00 |
| Human | 738.12 | 0.29 | 0.00 |
| Dwarf | 733.81 | 0.30 | 0.00 |
| Night Elf | 737.82 | 0.30 | 0.00 |
| Gnome | 738.05 | 0.30 | 0.00 |
| High Order | 746.80 | 0.31 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

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
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 68–102 dmg @ 1.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -2.5s: [Blade Flurry](https://www.wowhead.com/forever/spell=13877).

### Rotation priorities

1. Cast [Garrote](https://www.wowhead.com/forever/spell=11290).
2. Cast [Vanish](https://www.wowhead.com/forever/spell=1856) when (Energy ≥ 50 AND Combo points ≤ 3 AND Time remaining > 25).
3. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
4. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 12) OR (Combo points ≥ 5 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) time remaining < 0 AND Time remaining > 9)).
5. Cast [Adrenaline Rush](https://www.wowhead.com/forever/spell=13750) when Energy ≤ 30.
6. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
7. Cast [Spell 11300](https://www.wowhead.com/forever/spell=11300) when (Time remaining ≤ 3 AND Combo points ≥ 1).
8. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when Combo points ≥ 5.
9. Cast [Spell 11300](https://www.wowhead.com/forever/spell=11300) when Combo points ≥ 5.
10. Cast [Sinister Strike (rank 8)](https://www.wowhead.com/forever/spell=11294) when (`{"autoTimeToNext":{"autoType":"MainHand"}}` > `{"autoSwingTime":{"autoType":"MainHand"}}` − 0.15 OR Time remaining < 6 OR Energy ≥ 90).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 193.10 |
| [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | 150.92 |
| Auto-attack (tag 2) | 132.40 |
| Auto-attack (tag 3) | 104.29 |
| [Spell 11300](https://www.wowhead.com/forever/spell=11300) | 63.06 |
| [Spell 11356](https://www.wowhead.com/forever/spell=11356) | 32.92 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 26.85 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 14.22 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | -3270.8 |
| Energy | OtherActionEnergyRegen | +3257.4 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +475.7 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -372.3 |
| Energy | [Spell 11300](https://www.wowhead.com/forever/spell=11300) | -267.0 |
| Energy | OtherActionRefund | +126.7 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +94.1 |
| ComboPoints | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | +78.1 |
| Energy | [Blade Flurry](https://www.wowhead.com/forever/spell=13877) | -75.0 |
| ComboPoints | [Spell 11300](https://www.wowhead.com/forever/spell=11300) | -58.3 |
| Energy | [Garrote](https://www.wowhead.com/forever/spell=11290) | -50.0 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -36.8 |

## Rogue — Mutilate

**Talents:** 30/21/0 · `0053031033150105-302303202015`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 722.99 | 0.29 | 0.00 |
| Troll | 719.26 | 0.29 | 0.00 |
| Undead | 735.20 | 0.27 | 0.00 |
| Windshaper | 732.15 | 0.29 | 0.00 |
| Human | 724.59 | 0.29 | 0.00 |
| Dwarf | 719.14 | 0.28 | 0.00 |
| Night Elf | 725.25 | 0.28 | 0.00 |
| Gnome | 720.20 | 0.28 | 0.00 |
| High Order | 732.15 | 0.29 | 0.00 |

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

### Before the pull

- -1s: [Stealth (rank 1)](https://www.wowhead.com/forever/spell=1784).

### Rotation priorities

1. Cast [Ambush](https://www.wowhead.com/forever/spell=11269) when Time remaining ≤ 2s.
2. Cast [Vanish](https://www.wowhead.com/forever/spell=1856) when (Time remaining ≤ 1s AND Energy ≥ 60).
3. Cast [Garrote](https://www.wowhead.com/forever/spell=11290).
4. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
5. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when (Combo points ≥ 2 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) time remaining ≤ 0s AND Time remaining ≥ 12s).
6. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when (Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 12s).
7. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
8. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when (Combo points ≥ 4 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active).
9. Cast [Spell 11300](https://www.wowhead.com/forever/spell=11300) when (Combo points ≥ 1 AND Time remaining ≤ 1s).
10. Cast [Spell 11300](https://www.wowhead.com/forever/spell=11300) when (Combo points ≥ 3 AND Energy ≥ 50).
11. Cast [Mutilate](https://www.wowhead.com/forever/spell=1241584) when (Combo points < 3 AND Energy ≥ 70).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 171.71 |
| Auto-attack (tag 2) | 116.57 |
| [Spell 11300](https://www.wowhead.com/forever/spell=11300) | 75.29 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 60.71 |
| Auto-attack (tag 3) | 50.10 |
| [Spell 11300](https://www.wowhead.com/forever/spell=11300) | 49.52 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 46.09 |
| [Spell 11356](https://www.wowhead.com/forever/spell=11356) | 40.13 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | -3023.6 |
| Energy | OtherActionEnergyRegen | +3020.5 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +814.6 |
| Energy | [Spell 11300](https://www.wowhead.com/forever/spell=11300) | -749.2 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -353.8 |
| Energy | OtherActionRefund | +195.4 |
| ComboPoints | [Spell 11300](https://www.wowhead.com/forever/spell=11300) | -124.8 |
| ComboPoints | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | +94.3 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +93.4 |
| Energy | [Garrote](https://www.wowhead.com/forever/spell=11290) | -50.0 |
| ComboPoints | [Seal Fate](https://www.wowhead.com/forever/spell=14195) | +39.0 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -38.3 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `005303103005--3320203310013211501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 709.27 | 0.28 | 0.00 |
| Troll | 708.19 | 0.28 | 0.00 |
| Undead | 718.52 | 0.26 | 0.00 |
| Windshaper | 718.23 | 0.27 | 0.00 |
| Human | 708.63 | 0.28 | 0.00 |
| Dwarf | 705.85 | 0.27 | 0.00 |
| Night Elf | 710.87 | 0.27 | 0.00 |
| Gnome | 710.95 | 0.28 | 0.00 |
| High Order | 718.23 | 0.27 | 0.00 |

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

### Before the pull

- -1s: [Stealth (rank 1)](https://www.wowhead.com/forever/spell=1784).

### Rotation priorities

1. Cast [Slice and Dice](https://www.wowhead.com/forever/spell=6774) when (Combo points ≥ 1 AND NOT [Slice and Dice](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6s).
2. Cast [Vanish](https://www.wowhead.com/forever/spell=1856) when (NOT [Stealth](https://www.wowhead.com/forever/spell=1787) active AND Energy ≥ 60 AND Combo points ≤ 1 AND [Slice and Dice](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 10s).
3. Cast [Premeditation](https://www.wowhead.com/forever/spell=14183) when Combo points ≤ 2.
4. Cast [Garrote](https://www.wowhead.com/forever/spell=11290) when NOT [Garrote](https://www.wowhead.com/forever/spell=11290) DoT active.
5. Cast [Ambush](https://www.wowhead.com/forever/spell=11269) when Combo points ≤ 3.
6. Cast [Slice and Dice](https://www.wowhead.com/forever/spell=6774) when (Combo points ≥ 5 AND [Slice and Dice](https://www.wowhead.com/forever/spell=6774) time remaining ≤ 0s AND Time remaining ≥ 6s).
7. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
8. Cast [Hemorrhage](https://www.wowhead.com/forever/spell=16511) when (NOT [Hemorrhage](https://www.wowhead.com/forever/spell=16511) active AND Combo points ≥ 2 AND NOT [Rupture](https://www.wowhead.com/forever/spell=11275) DoT active AND Time remaining ≥ 10s).
9. Cast [Rupture](https://www.wowhead.com/forever/spell=11275) when (Combo points ≥ 3 AND NOT [Rupture](https://www.wowhead.com/forever/spell=11275) DoT active AND Time remaining ≥ 4s).
10. Cast [Spell 11300](https://www.wowhead.com/forever/spell=11300) when (Combo points ≥ 5 AND Energy ≥ 60).
11. Cast [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) when Combo points ≤ 3.
12. Cast [Hemorrhage](https://www.wowhead.com/forever/spell=16511) when Energy ≥ 60.
13. Use ready automatic cooldowns.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 175.45 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 115.03 |
| Auto-attack (tag 2) | 97.25 |
| Auto-attack (tag 3) | 59.14 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 49.20 |
| [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) | 38.15 |
| [Spell 11356](https://www.wowhead.com/forever/spell=11356) | 34.41 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 33.01 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +2985.8 |
| Energy | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | -2107.1 |
| Energy | [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) | -595.8 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +559.9 |
| Energy | [Rupture](https://www.wowhead.com/forever/spell=11275) | -534.1 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -407.3 |
| Energy | OtherActionRefund | +171.8 |
| Energy | [Ambush](https://www.wowhead.com/forever/spell=11269) | -99.4 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +94.8 |
| Energy | [Garrote](https://www.wowhead.com/forever/spell=11290) | -70.3 |
| ComboPoints | [Rupture](https://www.wowhead.com/forever/spell=11275) | -68.0 |
| ComboPoints | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | +66.1 |

## Shaman — Stormcaller

**Talents:** 29/22/0 · `551333100000134-055030021004002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Uses maximum registered Lightning Bolt. Overload uses its own client rank rows; Maelstrom/proc and server-script questions remain qualified.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 658.59 | 0.24 | 0.09 |
| Tauren | 644.36 | 0.23 | 0.09 |
| Troll | 654.57 | 0.23 | 0.08 |
| Windshaper | 654.21 | 0.23 | 0.11 |
| Dwarf | 655.59 | 0.23 | 0.09 |

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
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Mail) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 234 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 364 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Spell power / crit / hit (matched cost) — Feet (Mail) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 286 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Projected: Polished Driftwood Icon — Relic [+2 Sta, +1.75 Int, +2.25 SP]](https://www.wowhead.com/forever/item=920001012) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -10s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.
- -2.5s: [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438).
- -1.5s: [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Rotation priorities

1. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Target count = 1 AND NOT [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) DoT active).
2. Cast [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) when (Target count ≥ 2 AND NOT [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) DoT active).
3. Use ready automatic cooldowns.
4. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when (NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active AND Time remaining > 4s).
5. Cast [Earth Shock](https://www.wowhead.com/forever/spell=10414) when [Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active.
6. Cast [Lava Burst](https://www.wowhead.com/forever/spell=1326047).
7. Cast [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) when [Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active.
8. Cast [Earth Shock](https://www.wowhead.com/forever/spell=10414) when Mana ≥ Time remaining × 999.
9. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when Mana ≥ Time remaining × 18.
10. Cast [Lightning Bolt (rank 7)](https://www.wowhead.com/forever/spell=10391) when Mana ≥ Time remaining × 10.
11. Cast [Lightning Bolt (rank 2)](https://www.wowhead.com/forever/spell=529).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 253.80 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1326047) | 120.74 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 85.40 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=10391) | 85.14 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 38.60 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 28.38 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | 28.11 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 12.76 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -12023.0 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1326047) | -4999.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4972.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4917.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3603.9 |
| Mana | OtherActionManaRegen (tag 1) | +3479.4 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=10391) | -3070.4 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -2283.0 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | OtherActionManaRegen (tag 2) | +495.3 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=529) | -337.1 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -77.3 |

## Shaman — Elemental

**Talents:** 31/11/9 · `5503331000011351-055000001-05301`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 656.46 | 0.25 | 0.63 |
| Tauren | 650.02 | 0.25 | 0.64 |
| Troll | 652.60 | 0.25 | 0.36 |
| Windshaper | 658.02 | 0.25 | 0.50 |
| Dwarf | 653.98 | 0.25 | 0.60 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

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
| Ranged/relic | [Projected: Polished Driftwood Icon — Relic [+2 Sta, +1.75 Int, +2.25 SP]](https://www.wowhead.com/forever/item=920001012) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -8s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.
- -2s: [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Rotation priorities

1. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Target count = 1 AND NOT [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) DoT active).
2. Cast [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) when (Target count ≥ 2 AND NOT [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) DoT active).
3. Use ready automatic cooldowns.
4. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active.
5. Cast [Lava Burst (rank 4)](https://www.wowhead.com/forever/spell=1326047) when Time remaining > 2s.
6. Cast [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) when [Elemental Mastery](https://www.wowhead.com/forever/spell=16166) active.
7. Cast [Earth Shock](https://www.wowhead.com/forever/spell=10414) when (([Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active OR Target count ≥ 2) AND [Flame Shock](https://www.wowhead.com/forever/spell=29228) DoT time remaining > 2s).
8. Cast [Earth Shock](https://www.wowhead.com/forever/spell=10414) when ([Flame Shock](https://www.wowhead.com/forever/spell=29228) DoT time remaining > 2s AND Mana ≥ Time remaining × 1000).
9. Cast [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) when [Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active.
10. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when Mana ≥ Time remaining × 23.
11. Cast [Lightning Bolt (rank 7)](https://www.wowhead.com/forever/spell=10391).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 355.32 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1326047) | 114.91 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 82.04 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 37.78 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=10391) | 32.58 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 17.80 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 15.97 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=10391) | 1.62 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -18541.7 |
| Mana | OtherActionManaRegen (tag 1) | +7849.2 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1326047) | -5209.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5058.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4622.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3592.3 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -2325.7 |
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=10391) | -1270.3 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -969.0 |
| Mana | OtherActionManaRegen (tag 2) | +170.2 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -76.9 |

## Mage — Frost

**Talents:** 12/3/36 · `0500205-03-253520330100301251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 602.64 | 0.32 | 2.83 |
| Troll | 592.31 | 0.33 | 5.32 |
| Undead | 603.58 | 0.33 | 3.98 |
| Human | 604.95 | 0.31 | 1.30 |
| Gnome | 609.64 | 0.30 | 0.80 |
| High Order | 596.56 | 0.33 | 5.04 |

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
5. Cast [Cone of Cold](https://www.wowhead.com/forever/spell=10161) when [Winter's Chill](https://www.wowhead.com/forever/spell=12579) stacks ≥ 5.
6. Cast [Frostbolt (rank 11)](https://www.wowhead.com/forever/spell=10181).

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Frostbolt](https://www.wowhead.com/forever/spell=10181) | 387.24 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 142.06 |
| [Cone of Cold](https://www.wowhead.com/forever/spell=10161) | 80.34 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=10181) | -16627.0 |
| Mana | OtherActionManaRegen (tag 1) | +9025.0 |
| Mana | [Cone of Cold](https://www.wowhead.com/forever/spell=10161) | -9023.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4521.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4033.7 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3833.9 |
| Mana | OtherActionManaRegen (tag 2) | +2487.4 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.2 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +780.8 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +209.8 |

## Druid — Bear · Tank

**Talents:** 1/40/10 · `01-50022332120132012551-055`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Bear's outgoing base Rage remains the qualified damage-based model (DRU-012), with sourced double critical Rage. Incoming Rage now shares Warrior's configurable provisional expected-health scenario. The form threat multiplier is sourced at 1.5; this is not a survival ranking. Proc-enchant form eligibility is supported. Recovery has a sourced heal/ICD; Revelation requires an explicit, unverified proc model. Historical enchant-selection gains do not validate those assumptions.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 592.05 | 0.37 | 0.00 |
| Windshaper | 599.16 | 0.37 | 0.00 |
| Night Elf | 597.83 | 0.37 | 0.00 |
| High Order | 599.16 | 0.37 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Tauren | 2087.26 | 718.39 | 73.56 | 0.16% |
| Windshaper | 2107.24 | 715.48 | 78.36 | 0.42% |
| Night Elf | 2102.86 | 702.96 | 78.13 | 0.34% |
| High Order | 2107.24 | 715.48 | 78.36 | 0.42% |

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
| Ranged/relic | [Projected: Enraged Idol — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001021) | 65 | — |

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
| [Maul](https://www.wowhead.com/forever/spell=9881) | 357.45 |
| [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073) | 111.91 |
| [Lacerate](https://www.wowhead.com/forever/spell=414644) | 91.86 |
| [Lacerate](https://www.wowhead.com/forever/spell=414647) | 33.25 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.40 |
| Auto-attack (tag 1) | 1.57 |
| Auto-attack (tag 3) | 0.71 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -215936.1 |
| Health | OtherActionHealingModel | +211878.5 |
| Rage | OtherActionDamageTaken | +2398.0 |
| Health | [Frenzied Regeneration](https://www.wowhead.com/forever/spell=22842) | +2041.0 |
| Rage | [Maul](https://www.wowhead.com/forever/spell=9881) | -1414.0 |
| Rage | [Lacerate](https://www.wowhead.com/forever/spell=414644) | -1345.6 |
| Mana | [Faerie Fire](https://www.wowhead.com/forever/spell=9907) | -923.2 |
| Rage | [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073) | -826.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +738.0 |
| Health | [Holy Strength](https://www.wowhead.com/forever/spell=20007) | +669.8 |
| Rage | [Blood Frenzy](https://www.wowhead.com/forever/spell=37117) | +530.3 |
| Rage | OtherActionRefund | +517.3 |

## Paladin — Protection · Tank

**Talents:** 0/42/9 · `-5520513321301551-15102`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Seal of Fury is now selectable with sourced damage, shield requirement and talent-only depletion mana. Its shared replacement shield pool and independent lifetime are provisional; the primary seal and lifetime are recorded, not silently substituted. Threat-based taunt target selection is not modeled. Gear selections must pass both workload guard policies, which permit limited mitigation tradeoffs. This is not a survival ranking. Reflections use explicit provider power, never the receiving tank's spell power.

**Talent variants:** The allocation above is for Undead. Other races retain their saved allocations in the exact profiles. Historical validation ranges and source qualifications are recorded separately.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 568.38 | 0.27 | 0.34 |
| Human | 555.89 | 0.26 | 0.63 |
| Dwarf | 547.36 | 0.26 | 0.40 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Undead | 966.69 | 703.99 | 77.27 | 0.14% |
| Human | 973.22 | 702.84 | 78.34 | 0.34% |
| Dwarf | 954.05 | 704.53 | 77.90 | 0.28% |

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
| Ranged/relic | [Projected: Libram of Law — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001035) | 65 | — |

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
| Auto-attack (tag 1) | 98.60 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 76.15 |
| [Holy Shield](https://www.wowhead.com/forever/spell=20957) | 75.18 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 70.54 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 69.67 |
| Auto-attack (tag 3) | 44.84 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 36.89 |
| [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | 27.72 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -211197.1 |
| Health | OtherActionHealingModel | +207452.9 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -18468.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +16643.9 |
| Mana | [Shield Specialization](https://www.wowhead.com/forever/spell=20148) | +8619.5 |
| Mana | [Holy Shield](https://www.wowhead.com/forever/spell=20928) | -6473.7 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +3734.0 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -3325.2 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -2678.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2370.5 |
| Mana | OtherActionManaRegen (tag 1) | +2105.2 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -1980.0 |

## Warrior — Protection · Tank

**Talents:** 8/6/37 · `35-0501-251333120230021351`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Incoming Rage uses a configurable, unverified expected-creature-health/reference-Armor scenario shared with Bear. Its coefficient, level lookup and reference curve are assumptions, not measured server behavior; resolved parameters accompany each replay. Heroic Strike/Cleave replace swings and the APL reserves Rage for self-maintained Sunder. Earlier direct-cast Protection searches are superseded. Last Stand removes its temporary health on expiration; the nonlethal floor follows the qualified Classic convention.

**Talent variants:** The allocation above is for Undead. Other races retain their saved allocations in the exact profiles. Historical validation ranges and source qualifications are recorded separately.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 509.58 | 0.33 | 0.00 |
| Tauren | 519.15 | 0.34 | 0.00 |
| Troll | 515.48 | 0.34 | 0.00 |
| Undead | 549.22 | 0.29 | 0.00 |
| Windshaper | 518.54 | 0.34 | 0.00 |
| Human | 519.93 | 0.34 | 0.00 |
| Dwarf | 514.81 | 0.34 | 0.00 |
| Night Elf | 508.84 | 0.32 | 0.00 |
| Gnome | 518.08 | 0.34 | 0.00 |
| High Order | 518.54 | 0.34 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Orc | 1302.22 | 513.38 | 52.57 | 0.00% |
| Tauren | 1323.56 | 492.77 | 48.26 | 0.00% |
| Troll | 1316.43 | 502.11 | 51.67 | 0.00% |
| Undead | 1324.39 | 471.63 | 50.55 | 0.00% |
| Windshaper | 1263.74 | 496.35 | 51.81 | 0.00% |
| Human | 1265.33 | 481.06 | 50.73 | 0.00% |
| Dwarf | 1314.72 | 505.90 | 51.68 | 0.00% |
| Night Elf | 1233.25 | 493.24 | 51.65 | 0.00% |
| Gnome | 1264.66 | 494.17 | 51.33 | 0.00% |
| High Order | 1263.74 | 496.35 | 51.81 | 0.00% |

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

1. Cast [Battle Shout Rank 6](https://www.wowhead.com/forever/spell=11551) when ([Battle Shout Rank 6](https://www.wowhead.com/forever/spell=11551) time remaining ≤ 3s AND Time remaining ≥ 8s AND Rage ≥ 10).
2. Cast [Shield Wall](https://www.wowhead.com/forever/spell=871) when `{"currentHealthPercent":{}}` < 35%.
3. Cast [Last Stand](https://www.wowhead.com/forever/spell=12975) when `{"currentHealthPercent":{}}` < 30%.
4. Cast [Shield Block](https://www.wowhead.com/forever/spell=2565) when (Rage ≥ 15 AND [Shield Block](https://www.wowhead.com/forever/spell=2565) time remaining < 1s).
5. Cast [Bloodrage](https://www.wowhead.com/forever/spell=2687) when (Rage < 35 AND `{"currentHealthPercent":{}}` > 50%).
6. Cast [Sunder Armor](https://www.wowhead.com/forever/spell=11597) when ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining ≤ 8s AND Time remaining > [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining).
7. Cast [Thunder Clap](https://www.wowhead.com/forever/spell=11581) when ([Thunder Clap](https://www.wowhead.com/forever/spell=11581) time remaining ≤ 3s AND Time remaining ≥ 4s).
8. Cast [Demoralizing Shout Rank 5](https://www.wowhead.com/forever/spell=11556) when ([Demoralizing Shout Rank 5](https://www.wowhead.com/forever/spell=11556) time remaining ≤ 4s AND Time remaining ≥ 5s).
9. Cast [Sunder Armor](https://www.wowhead.com/forever/spell=11597) when [Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks < 5.
10. Cast [Thunder Clap](https://www.wowhead.com/forever/spell=11581) when (Target count ≥ 3 AND (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Thunder Clap](https://www.wowhead.com/forever/spell=11581) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost)).
11. Cast [Revenge](https://www.wowhead.com/forever/spell=11601) when (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Revenge](https://www.wowhead.com/forever/spell=11601) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost).
12. Cast [Shield Slam](https://www.wowhead.com/forever/spell=23925) when (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Shield Slam](https://www.wowhead.com/forever/spell=23925) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost).
13. Cast [Cleave](https://www.wowhead.com/forever/spell=20569) when ((Target count > 1 AND Rage ≥ 20) AND (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Cleave](https://www.wowhead.com/forever/spell=20569) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost)).
14. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=11567) when ((Target count = 1 AND Rage ≥ 20) AND (NOT `{"auraShouldRefresh":{"sourceUnit":{"type":"CurrentTarget"},"auraId":{"spellId":11597},"maxOverlap":{"const":{"val":"8s"}}}}` OR ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) stacks = 5 AND ([Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining > 8s OR Time remaining ≤ [Sunder Armor](https://www.wowhead.com/forever/spell=11597) time remaining)) OR Rage ≥ [Heroic Strike](https://www.wowhead.com/forever/spell=11567) current cost + [Sunder Armor](https://www.wowhead.com/forever/spell=11597) current cost + [Shield Block](https://www.wowhead.com/forever/spell=2565) current cost)).
15. Cast [Sunder Armor](https://www.wowhead.com/forever/spell=11597) when Rage ≥ 45.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Heroic Strike](https://www.wowhead.com/forever/spell=11567) | 235.57 |
| [Revenge](https://www.wowhead.com/forever/spell=11601) | 141.68 |
| [Shield Slam](https://www.wowhead.com/forever/spell=23925) | 99.05 |
| Auto-attack (tag 3) | 28.33 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 20.06 |
| [Fiery Weapon](https://www.wowhead.com/forever/spell=13897) | 8.61 |
| [Thunder Clap](https://www.wowhead.com/forever/spell=11581) | 7.21 |
| Auto-attack (tag 1) | 6.29 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -141489.6 |
| Health | OtherActionHealingModel | +138814.2 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +2672.9 |
| Rage | OtherActionDamageTaken | +2461.3 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=11567) | -1557.2 |
| Rage | [Shield Slam](https://www.wowhead.com/forever/spell=23925) | -755.3 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -662.2 |
| Rage | [Shield Block](https://www.wowhead.com/forever/spell=2565) | -524.1 |
| Rage | OtherActionRefund | +444.7 |
| Rage | [Shield Specialization](https://www.wowhead.com/forever/spell=12727) | +380.6 |
| Rage | Auto-attack (tag 1) | +259.1 |
| Rage | [Shield Specialization](https://www.wowhead.com/forever/spell=12727) | +163.9 |

## Priest — Smite

**Talents:** 31/17/3 · `515035031030001031-00505023002-003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Uses maximum registered ranks. Penance's static rank anomaly and actual skill-book acquisition remain source gaps; no efficient lower-rank fallback is assumed.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 484.74 | 0.20 | 62.38 |
| Undead | 503.61 | 0.20 | 57.06 |
| Human | 495.34 | 0.20 | 56.57 |
| Dwarf | 487.63 | 0.20 | 59.73 |
| Night Elf | 489.21 | 0.19 | 60.58 |
| Gnome | 500.47 | 0.20 | 55.06 |

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
3. Cast [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) when (NOT [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) DoT active AND Time remaining ≥ 12s).
4. `{"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10934,"rank":8}}}]}}`.
5. Cast [Smite](https://www.wowhead.com/forever/spell=10934) when Mana ≥ Time remaining × 20.
6. Cast [Smite](https://www.wowhead.com/forever/spell=10934).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Smite](https://www.wowhead.com/forever/spell=10934) | 282.44 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 133.86 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 75.90 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 11.42 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -17979.9 |
| Mana | OtherActionManaRegen (tag 1) | +7157.7 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -6140.3 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -4938.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4248.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3601.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3275.5 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1599.9 |
| Mana | OtherActionManaRegen (tag 2) | +1457.9 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |
