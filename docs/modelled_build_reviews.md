# Build reviews

These are hypothetical modeled-only gear results, not obtainable item rankings. Sources linked for modeled items are allocation references only and have different stats. Gear is compared under the [model assumptions](modelled_gear.md).

The tables and [matrix](../artifacts/modelled_gear/forever_dps_5min.png) use the same 201 common-seed replays. The modeled search and original real-item benchmark are separate; neither proves available launch gear. The [October 16 follow-up](october-16-follow-up.md) records the imbue slots, real relics, Rake scaling, rotation review and confirmed re-search of every DPS profile; the [October 15 update](october-15-update.md) records the client 1.60.1.70425 changes and the earlier [October 11 audit](october-11-audit.md) and [October 11 log audit](october-11-log-audit.md) record the corrections before it. Earlier search gains are historical, not current proof of an optimum.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and hit from selected gear, enchants, talents and racials only. [Scenario and exchange model](../tools/forever_bench/README.md) · [Questions and coverage](uncertainties.md)

Tank rows include frontal incoming attacks and modeled healing, with their recorded shared raid support except assigned self-maintained duties. They do not share the non-attacking DPS encounter. [Support and Judgement correction](tank_support.md). [Tank selection and guardrails](tank_selection.md).

## Paladin — Retribution

**Talents:** 13/7/31 · `55003-232-05225231001330301`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Separate seal Echoes can coexist and fire on a landed white swing; their simultaneous server behavior remains qualified (PAL-002). Active presets use maximum registered ranks, not deliberate downranking.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 998.62 | 0.56 | 0.00 |
| Human | 1005.42 | 0.56 | 0.00 |
| Dwarf | 1000.47 | 0.56 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

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
| Main hand | [Modeled: Spell two-hand — Two-hand [+25 Sta, +24 Int, +24 SP, +15 Str, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272682) | 65 | Enchant Weapon - Lifestealing |
| Ranged/relic | [Projected: Libram of Law — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001035) | 65 | — |

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

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 319.65 |
| Auto-attack (tag 1) | 205.18 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 100.95 |
| Auto-attack (tag 3) | 99.08 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 87.81 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 46.83 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 45.69 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 45.34 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +14246.6 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -9664.8 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -8761.5 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -7105.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3604.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3453.2 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3183.7 |
| Mana | OtherActionManaRegen (tag 1) | +3107.4 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +3105.4 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1291.8 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -521.5 |
| Mana | OtherActionManaRegen (tag 2) | +5.5 |

## Warrior — Fury

**Talents:** 18/33/0 · `20305213002-25053005141010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Level-60/off-hand rage and some extra-swing interactions remain provisional. The queued-Heroic-Strike off-hand hit bug is removed. Deep Wounds uses the reported intended no-AP rollover model; live delivery and off-hand payload attribution remain qualified.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 992.88 | 0.52 | 0.00 |
| Tauren | 994.98 | 0.52 | 0.00 |
| Troll | 982.21 | 0.50 | 0.00 |
| Undead | 983.97 | 0.50 | 0.00 |
| Windshaper | 999.91 | 0.51 | 0.00 |
| Human | 989.62 | 0.51 | 0.00 |
| Dwarf | 986.82 | 0.52 | 0.00 |
| Night Elf | 980.16 | 0.52 | 0.00 |
| Gnome | 982.43 | 0.51 | 0.00 |
| High Order | 999.91 | 0.51 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Strength / agility / crit / hit — Head (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 599 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Rapidity |
| Neck | [Modeled: Strength / crit / hit — Neck [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 189.91 |
| Auto-attack (tag 2) | 163.59 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 140.05 |
| Auto-attack (tag 1) | 118.87 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 112.31 |
| Auto-attack (tag 3) | 71.64 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 63.47 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 54.24 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 2) | +1382.8 |
| Rage | Auto-attack (tag 1) | +1343.1 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -1060.2 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -888.6 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -792.8 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -483.1 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +180.6 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +157.0 |
| Rage | OtherActionRefund | +118.9 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.9 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +89.5 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -88.3 |

## Hunter — Beast Mastery

**Talents:** 31/20/0 · `5320001505101251-00522510005`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Hawk lifetimes and ranged-speed behavior are source-backed; high-rank assault damage, crit and Ferocity transfer remain qualified (HUN-004). The two-Hawk policy is conditional on the central calibration.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 975.39 | 0.28 | 0.03 |
| Tauren | 967.76 | 0.28 | 0.03 |
| Troll | 968.90 | 0.28 | 0.02 |
| Windshaper | 981.72 | 0.29 | 0.02 |
| Human | 981.92 | 0.28 | 0.02 |
| Dwarf | 969.54 | 0.28 | 0.03 |
| Night Elf | 975.19 | 0.28 | 0.02 |
| High Order | 981.72 | 0.29 | 0.02 |

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
| Shoot | 387.54 |
| Cat: Auto-attack (tag 1) | 225.11 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 130.16 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 85.26 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 52.17 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 31.86 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 30.87 |
| Hawk 1: Auto-attack (tag 1) | 14.69 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -12887.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9232.1 |
| Mana | OtherActionManaRegen (tag 1) | +6388.5 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -5576.3 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -4462.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3606.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3601.1 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -2677.8 |
| Mana | OtherActionManaRegen (tag 2) | +743.8 |
| Mana | [Bestial Wrath](https://www.wowhead.com/forever/spell=19574) | -619.2 |
| Mana | [Intimidation](https://www.wowhead.com/forever/spell=19577) | -339.7 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Hunter — Marksmanship

**Talents:** 5/35/11 · `5-0050552011523051-50005001`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Lone Wolf requires no active pet. Improved Tracking assumes tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 965.82 | 0.38 | 0.00 |
| Tauren | 961.31 | 0.39 | 0.00 |
| Troll | 959.56 | 0.37 | 0.00 |
| Windshaper | 973.25 | 0.38 | 0.00 |
| Human | 971.53 | 0.39 | 0.00 |
| Dwarf | 963.22 | 0.39 | 0.00 |
| Night Elf | 966.86 | 0.38 | 0.00 |
| High Order | 973.25 | 0.38 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

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
2. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=25295) when (NOT [Serpent Sting](https://www.wowhead.com/forever/spell=25295) DoT active AND Time remaining ≥ 12s).
3. Cast [Aimed Shot](https://www.wowhead.com/forever/spell=20904).
4. Cast [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) when Mana fraction ≥ 30%.
5. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287) when Mana fraction ≥ 15%.

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Shoot | 508.98 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | 176.96 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 114.18 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 92.04 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 71.16 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 9.92 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Aimed Shot](https://www.wowhead.com/forever/spell=20904) | -11306.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +10650.0 |
| Mana | [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | -6798.7 |
| Mana | OtherActionManaRegen (tag 1) | +6202.0 |
| Mana | [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | -4941.5 |
| Mana | [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | -3935.7 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3602.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3584.2 |
| Mana | OtherActionManaRegen (tag 2) | +482.1 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |

## Warrior — Fury (Sunder)

**Talents:** 18/33/0 · `20305213002-25053005141010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** This Warrior personally builds and refreshes five Sunder stacks. External Sunder and Expose Armor are disabled only for this row; its initial armor ramp and lost globals are included. The provisional Warrior rage model still applies.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 960.57 | 0.51 | 0.00 |
| Tauren | 967.36 | 0.53 | 0.00 |
| Troll | 956.01 | 0.51 | 0.00 |
| Undead | 962.46 | 0.50 | 0.00 |
| Windshaper | 973.04 | 0.52 | 0.00 |
| Human | 968.25 | 0.51 | 0.00 |
| Dwarf | 960.39 | 0.53 | 0.00 |
| Night Elf | 953.64 | 0.52 | 0.00 |
| Gnome | 957.01 | 0.51 | 0.00 |
| High Order | 973.04 | 0.52 | 0.00 |

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
| Legs | [Modeled: Strength / crit / hit — Legs (Plate) [+18 Sta, +21 Crit, +21 Str, +23 Hit, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
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
11. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
12. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 30 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 176.66 |
| Auto-attack (tag 2) | 163.19 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 129.24 |
| Auto-attack (tag 1) | 125.31 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 106.72 |
| Auto-attack (tag 3) | 74.11 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 62.72 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 54.17 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +1415.3 |
| Rage | Auto-attack (tag 2) | +1392.6 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -980.5 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -857.2 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -749.7 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -451.6 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -275.1 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +184.7 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +171.5 |
| Rage | OtherActionRefund | +126.3 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +99.8 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +93.9 |

## Hunter — Pet/Melee

**Talents:** 16/11/24 · `53200005001-0050051-5302002300501201`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Two independent Hawks use provisional high-rank damage and crit calibration (HUN-004); no owner-crit or unverified Lone Wolf coexistence is assumed. Tracking matches the Dragonkin reference target.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 950.30 | 0.40 | 0.00 |
| Windshaper | 963.13 | 0.38 | 0.00 |
| Orc | 955.44 | 0.40 | 0.00 |
| Human | 962.93 | 0.39 | 0.00 |
| Night Elf | 953.00 | 0.39 | 0.00 |
| Dwarf | 954.24 | 0.39 | 0.00 |
| Troll | 956.26 | 0.39 | 0.00 |
| High Order | 963.13 | 0.38 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Mail) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
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
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword one-hand — One-hand [+2 Sta, +24 AP, +14 Crit, 109–165 dmg @ 2.9s] (modeled; reference only)](https://www.wowhead.com/forever/item=272452) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Crusader |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 168.07 |
| Cat: Auto-attack (tag 1) | 137.76 |
| Auto-attack (tag 1) | 136.24 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 109.04 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 101.00 |
| Auto-attack (tag 3) | 77.93 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 57.52 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 36.52 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +16114.1 |
| Mana | [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | -7642.7 |
| Mana | OtherActionManaRegen (tag 1) | +5909.0 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -4318.3 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3319.2 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -3083.9 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -1658.2 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -1544.1 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) | -180.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +0.5 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 954.63 | 0.40 | 0.00 |
| Tauren | 950.70 | 0.40 | 0.00 |
| Troll | 946.75 | 0.39 | 0.00 |
| Windshaper | 956.08 | 0.39 | 0.00 |
| Human | 958.66 | 0.40 | 0.00 |
| Dwarf | 951.70 | 0.40 | 0.00 |
| Night Elf | 953.29 | 0.40 | 0.00 |
| High Order | 956.08 | 0.39 | 0.00 |

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
| Auto-attack (tag 2) | 172.32 |
| Auto-attack (tag 1) | 135.50 |
| Cat: Auto-attack (tag 1) | 128.04 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 114.94 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 105.28 |
| Auto-attack (tag 3) | 89.22 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 65.25 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 62.32 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +17074.4 |
| Mana | [Wing Clip](https://www.wowhead.com/forever/spell=14268) | -9081.9 |
| Mana | [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | -4339.0 |
| Mana | [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | -3659.3 |
| Mana | [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | -3397.7 |
| Mana | OtherActionManaRegen (tag 1) | +3156.5 |
| Mana | [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | -2375.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +1952.6 |
| Mana | [Rapid Fire](https://www.wowhead.com/forever/spell=3045) | -200.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +193.2 |
| Mana | [Hunter's Mark (rank 4)](https://www.wowhead.com/forever/spell=14325) | -180.0 |
| Mana | [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447) | -110.0 |

## Paladin — Physical Ret

**Talents:** 13/7/31 · `52003003-232-05225031101330311`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Uses a physical equipment budget without caster gear. Champion of the Light converts existing Intellect into damage-only spell bonus, not healing power. Seal Echo ordering remains qualified (PAL-002).


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 936.46 | 0.53 | 0.00 |
| Human | 949.58 | 0.59 | 0.01 |
| Dwarf | 945.52 | 0.59 | 0.01 |

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
| Hands | [Modeled: Strength / agility / crit / hit — Hands (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 461 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Weapon - Agility |
| Waist | [Modeled: Strength / agility / crit / hit — Waist (Plate) [+21 Str, +14 Sta, +11 Crit, +8 Hit, +14 Agi, 415 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Strength / agility / crit / hit — Legs (Plate) [+28 Str, +18 Sta, +14 Crit, +10 Hit, +18 Agi, 646 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Strength / crit / hit — Feet (Plate) [+14 Sta, +16 Crit, +16 Str, +18 Hit, 507 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Strength / crit / hit — Finger [+11 Sta, +12 Crit, +12 Str, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Strength / crit / hit — Trinket [+11 Sta, +13 Crit, +13 Str, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Sword two-hand — Two-hand [+34 Sta, +23 Str, +14 Crit, 187–281 dmg @ 3.8s] (modeled; reference only)](https://www.wowhead.com/forever/item=272604) | 65 | Enchant Weapon - Crusader |
| Ranged/relic | [Projected: Libram of Invocation — Relic [+2 Sta, +1 Agi, +6.5 AP]](https://www.wowhead.com/forever/item=920001015) | 65 | — |

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

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 324.34 |
| Auto-attack (tag 1) | 230.60 |
| Auto-attack (tag 3) | 111.47 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 84.25 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 81.29 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 32.06 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 32.00 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20965) | 30.10 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +11309.2 |
| Mana | [Spell 20293](https://www.wowhead.com/forever/spell=20293) | -9090.1 |
| Mana | [Spell 20919](https://www.wowhead.com/forever/spell=20919) | -8239.6 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -6250.0 |
| Mana | OtherActionManaRegen (tag 1) | +5419.4 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3598.9 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3393.6 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -3201.5 |
| Mana | [Sanctified Judgement](https://www.wowhead.com/forever/spell=31876) | +2904.5 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -1375.8 |
| Mana | [Holy Strike](https://www.wowhead.com/forever/spell=10333) | -521.1 |
| Mana | OtherActionManaRegen (tag 2) | +5.0 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 916.52 | 0.42 | 0.00 |
| Troll | 902.66 | 0.41 | 0.00 |
| Undead | 920.30 | 0.41 | 0.00 |
| Human | 922.52 | 0.43 | 0.00 |
| Gnome | 908.99 | 0.41 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit — Neck [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit / hit (matched cost) — Back (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit — Chest (Cloth) [+18 Sta, +23 Int, +29 SP, +14 Crit, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Shadow Power |
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

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 342.80 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 98.34 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 91.16 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 90.57 |
| Succubus: [Demonic Brand](https://www.wowhead.com/forever/spell=1293697) | 83.12 |
| Succubus: Auto-attack (tag 1) | 77.17 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 55.87 |
| [Soul Fire](https://www.wowhead.com/forever/spell=17924) | 51.44 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -24327.3 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +12220.6 |
| Mana | [Fel Energy](https://www.wowhead.com/forever/spell=18792) | +11715.8 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6560.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5396.7 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5157.3 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -4265.4 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3885.7 |
| Mana | [Soul Fire](https://www.wowhead.com/forever/spell=17924) | -3176.5 |
| Mana | OtherActionManaRegen (tag 1) | +3113.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2635.8 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-55050000002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 869.00 | 0.60 | 0.00 |
| Tauren | 869.38 | 0.61 | 0.00 |
| Troll | 858.57 | 0.60 | 0.00 |
| Undead | 863.11 | 0.59 | 0.00 |
| Windshaper | 868.03 | 0.58 | 0.00 |
| Human | 869.87 | 0.61 | 0.00 |
| Dwarf | 868.20 | 0.61 | 0.00 |
| Night Elf | 865.12 | 0.60 | 0.00 |
| Gnome | 857.90 | 0.60 | 0.00 |
| High Order | 868.03 | 0.58 | 0.00 |

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
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 159.73 |
| Auto-attack (tag 1) | 145.84 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 127.57 |
| Auto-attack (tag 3) | 117.73 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 83.96 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 64.29 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 50.69 |
| [Rend](https://www.wowhead.com/forever/spell=11574) | 39.42 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +2163.0 |
| Rage | [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | -1083.5 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -664.9 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -215.0 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -176.9 |
| Rage | [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) | -170.2 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -154.6 |
| Rage | [Rend](https://www.wowhead.com/forever/spell=11574) | -145.0 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +144.5 |
| Rage | OtherActionRefund | +101.0 |
| Rage | [Anger Management](https://www.wowhead.com/forever/spell=12296) | +98.9 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +85.6 |

## Mage — Arcane–Frost

**Talents:** 28/0/23 · `050200500310031053--05550002010003002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Ice Lance uses a provisional 0.10 SP share calibrated from lower ranks. Fingers of Frost ordering and Arcane Missiles/Clearcasting proc interpretation remain unresolved; Missile Barrage now requires a landed eligible hit.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 851.26 | 0.66 | 0.72 |
| Troll | 838.72 | 0.64 | 0.80 |
| Undead | 850.21 | 0.64 | 0.63 |
| Human | 856.77 | 0.64 | 0.55 |
| Gnome | 851.75 | 0.64 | 0.46 |
| High Order | 855.72 | 0.63 | 0.71 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

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
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Passive Intellect / spell power / crit — Trinket [+11 Sta, +14 Int, +18 SP, +9 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 80%.
2. Use ready automatic cooldowns.
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction ≤ 15% AND Time remaining ≥ 25s).
4. Cast [Ice Lance](https://www.wowhead.com/forever/spell=30455) when [Fingers of Frost](https://www.wowhead.com/forever/spell=44543) active.
5. Cast [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) when [Missile Barrage](https://www.wowhead.com/forever/spell=44404) active.
6. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451) when ([Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks < 4 AND Mana fraction > 20%).
7. Cast [Frostbolt](https://www.wowhead.com/forever/spell=25304).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 487.69 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 172.49 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 154.59 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 42.01 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -27751.5 |
| Mana | OtherActionManaRegen (tag 1) | +10476.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9209.2 |
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -6959.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4565.4 |
| Mana | OtherActionManaRegen (tag 2) | +4424.0 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -1277.7 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1097.9 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +832.1 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +469.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +227.8 |

## Mage — Arcane

**Talents:** 33/3/15 · `0512005023100311531-03-005500032`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 850.69 | 0.71 | 0.80 |
| Troll | 831.46 | 0.69 | 0.94 |
| Undead | 844.44 | 0.69 | 0.81 |
| Human | 856.34 | 0.68 | 0.57 |
| Gnome | 854.00 | 0.69 | 0.50 |
| High Order | 854.20 | 0.69 | 0.78 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 81 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Neck | [Modeled: Intellect / spell power / crit / hit (matched cost) — Neck [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Intellect / spell power / crit / hit (matched cost) — Shoulders (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 75 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Chromatic Mantle of the Dawn |
| Back | [Modeled: Intellect / spell power / crit — Back (Cloth) [+11 Sta, +13 Int, +17 SP, +8 Crit, 50 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Cloak - Superior Defense |
| Chest | [Modeled: Intellect / spell power / crit / hit (matched cost) — Chest (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 100 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Chest - Greater Stats |
| Wrists | [Modeled: Intellect / spell power / crit / hit (matched cost) — Wrists (Cloth) [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int, 44 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Bracer - Healing Power |
| Hands | [Modeled: Intellect / spell power / crit / hit (matched cost) — Hands (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 63 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Enchant Gloves - Healing Power |
| Waist | [Modeled: Intellect / spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +11 Crit, +8 Hit, +24 SP, +14 Int, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Legs | [Modeled: Intellect / spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
| Feet | [Modeled: Intellect / spell power / crit — Feet (Cloth) [+14 Sta, +18 Int, +22 SP, +11 Crit, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Intellect / spell power / crit — Finger [+11 Sta, +13 Int, +17 SP, +8 Crit] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ring 2 | [Modeled: Intellect / spell power / crit / hit (matched cost) — Finger [+11 Sta, +8 Crit, +6 Hit, +18 SP, +10 Int] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
| Off hand | [Modeled: Spell-power held off-hand — Held off-hand [+15 Sta, +14 Int, +18 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | — |
| Ranged/relic | [Modeled: Spell-power wand (projected) — Ranged/wand [+9 Int, +12 SP, 49–92 dmg @ 1.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=279246) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 15% AND Time remaining > 20s).
3. Cast [Arcane Missiles (rank 8)](https://www.wowhead.com/forever/spell=25345) when [Missile Barrage](https://www.wowhead.com/forever/spell=44404) active.
4. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451) when ([Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks < 4 AND Mana fraction > 20%).
5. Cast [Arcane Wrath](https://www.wowhead.com/forever/spell=1326082).
6. Cast [Arcane Missiles (rank 8)](https://www.wowhead.com/forever/spell=25345).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 584.97 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 166.08 |
| [Arcane Wrath](https://www.wowhead.com/forever/spell=1326082) | 105.30 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | -27570.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9825.8 |
| Mana | OtherActionManaRegen (tag 1) | +8948.5 |
| Mana | OtherActionManaRegen (tag 2) | +6227.3 |
| Mana | [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | -5590.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4773.4 |
| Mana | [Arcane Wrath](https://www.wowhead.com/forever/spell=1326082) | -3876.4 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.4 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +849.6 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +589.3 |

## Druid — Feral

**Talents:** 7/34/10 · `05002-45010032021032212051-055`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Blizzard supports proc-enchant form eligibility, and Crusader's sourced health return is modeled. Its exact PPM/conditional interactions remain qualified (DRU-011); the Night Elf profile uses that enchant.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 848.45 | 0.29 | 2.34 |
| Windshaper | 851.08 | 0.28 | 0.98 |
| Night Elf | 837.61 | 0.28 | 2.08 |
| High Order | 851.08 | 0.28 | 0.98 |

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
3. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Time remaining ≤ 4s AND Combo points ≥ 3).
4. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Combo points ≥ 5 AND [Rip](https://www.wowhead.com/forever/spell=9896) DoT time remaining ≥ 5s).
5. Cast [Rip](https://www.wowhead.com/forever/spell=9896) when (Combo points ≥ 4 AND NOT [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND Time remaining ≥ 8s).
6. Cast [Rake](https://www.wowhead.com/forever/spell=9904) when (NOT [Rake](https://www.wowhead.com/forever/spell=9904) DoT active AND Time remaining ≥ 9s AND Combo points < 5).
7. Cast [Shifting Power](https://www.wowhead.com/forever/spell=1322605) when Energy ≤ 60.
8. Cast [Shred](https://www.wowhead.com/forever/spell=9830).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 230.30 |
| [Shred](https://www.wowhead.com/forever/spell=9830) | 220.85 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 182.88 |
| [Rake](https://www.wowhead.com/forever/spell=9904) | 103.44 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 57.11 |
| Auto-attack (tag 3) | 53.17 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 3.34 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shifting Power](https://www.wowhead.com/forever/spell=1322605) | -28116.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +14644.6 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3952.1 |
| Mana | OtherActionManaRegen (tag 2) | +3492.6 |
| Energy | [Shred](https://www.wowhead.com/forever/spell=9830) | -3257.7 |
| Energy | OtherActionEnergyRegen | +3107.4 |
| Energy | [Shifting Power](https://www.wowhead.com/forever/spell=1322605) | +1643.7 |
| Mana | OtherActionManaRegen (tag 1) | +1632.3 |
| Energy | [Rake](https://www.wowhead.com/forever/spell=9904) | -924.9 |
| Energy | [Rip](https://www.wowhead.com/forever/spell=9896) | -553.0 |
| Energy | [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | -444.6 |
| Energy | OtherActionRefund | +263.2 |

## Warrior — 2H Bloodthirst

**Talents:** 20/31/0 · `30304203032-55050020151010501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Level-60 rage and Flurry charge timing remain unverified; this row uses the provisional Warrior model.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 846.35 | 0.58 | 0.00 |
| Tauren | 842.26 | 0.58 | 0.00 |
| Undead | 838.74 | 0.58 | 0.00 |
| Troll | 836.87 | 0.57 | 0.00 |
| Windshaper | 843.98 | 0.57 | 0.00 |
| Human | 843.07 | 0.58 | 0.00 |
| Dwarf | 840.94 | 0.58 | 0.00 |
| Gnome | 833.62 | 0.58 | 0.00 |
| Night Elf | 835.52 | 0.57 | 0.00 |
| High Order | 843.98 | 0.57 | 0.00 |

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
| Auto-attack (tag 1) | 203.16 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 148.21 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 106.98 |
| Auto-attack (tag 3) | 91.26 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 87.43 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 70.46 |
| [Deep Wounds](https://www.wowhead.com/forever/spell=12867) | 47.85 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 46.09 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Rage | Auto-attack (tag 1) | +2361.4 |
| Rage | [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | -1028.2 |
| Rage | [Execute](https://www.wowhead.com/forever/spell=20662) | -731.8 |
| Rage | [Whirlwind](https://www.wowhead.com/forever/spell=1680) | -475.3 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -287.6 |
| Rage | [Item 13442](https://www.wowhead.com/forever/item=13442) | +143.7 |
| Rage | [Slam](https://www.wowhead.com/forever/spell=11605) | -127.5 |
| Rage | OtherActionRefund | +89.6 |
| Rage | [Bloodrage](https://www.wowhead.com/forever/spell=2687) | +87.5 |
| Rage | [Berserker Stance](https://www.wowhead.com/forever/spell=2458) | -55.9 |
| Rage | [Unbridled Wrath](https://www.wowhead.com/forever/spell=12964) | +55.5 |
| Rage | [Overpower](https://www.wowhead.com/forever/spell=11585) | -49.5 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2550020013520105--0550015103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 838.64 | 0.49 | 0.00 |
| Troll | 823.43 | 0.49 | 0.00 |
| Undead | 831.18 | 0.48 | 0.00 |
| Human | 844.36 | 0.49 | 0.00 |
| Gnome | 822.26 | 0.48 | 0.00 |

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
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 405.99 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 118.79 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 102.13 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 85.09 |
| Succubus: Auto-attack (tag 1) | 67.20 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 31.65 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 19.79 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 8.14 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -29440.7 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +24495.5 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6581.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5369.3 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5228.0 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3604.4 |
| Mana | [Siphon Life](https://www.wowhead.com/forever/spell=18881) | -3462.5 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3332.5 |
| Mana | OtherActionManaRegen (tag 1) | +3092.8 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -462.0 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -343.1 |

## Priest — Shadow

**Talents:** 18/0/33 · `30503300103--50452001201302251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Death's script value 150 and backlash interactions remain unverified; no extra execute multiplier is inferred. Self-damage is recorded without healer survival constraints.

**Rotation variants:** The priorities below are for Human. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 820.69 | 0.24 | 0.00 |
| Undead | 828.86 | 0.25 | 0.00 |
| Human | 835.32 | 0.25 | 0.00 |
| Dwarf | 833.53 | 0.26 | 0.00 |
| Night Elf | 828.60 | 0.25 | 0.00 |
| Gnome | 827.12 | 0.25 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Human

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
2. Cast [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) when NOT [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) DoT active.
3. Cast [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) when (NOT [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) DoT active AND Time remaining ≥ 10).
4. Cast [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636).
5. `{"condition":{},"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10947}}}]}}` when `{}`.
6. Cast [Mind Blast (rank 9)](https://www.wowhead.com/forever/spell=10947).
7. Cast [Mind Flay (rank 6)](https://www.wowhead.com/forever/spell=18807).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 342.41 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 165.21 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 150.41 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 102.54 |
| [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | 74.75 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +8149.9 |
| Mana | [Mind Flay](https://www.wowhead.com/forever/spell=18807) | -6358.3 |
| Mana | [Mind Blast](https://www.wowhead.com/forever/spell=10947) | -6079.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4081.3 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -2850.9 |
| Mana | [Shadow Word: Death](https://www.wowhead.com/forever/spell=1309636) | -2720.5 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +2452.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2359.1 |
| Mana | [Shadowform](https://www.wowhead.com/forever/spell=15473) | -550.4 |
| Mana | OtherActionManaRegen (tag 2) | +34.7 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2530000003-0005-0050355103101351`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 812.37 | 0.34 | 0.00 |
| Troll | 812.76 | 0.34 | 0.00 |
| Undead | 811.17 | 0.33 | 0.00 |
| Human | 819.66 | 0.34 | 0.00 |
| Gnome | 802.73 | 0.33 | 0.00 |

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
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
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
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 263.80 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 124.51 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 99.41 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 91.81 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 84.61 |
| Succubus: Auto-attack (tag 1) | 72.75 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 54.13 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 16.36 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +21748.9 |
| Mana | [Incinerate](https://www.wowhead.com/forever/spell=1293813) | -19453.1 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6686.2 |
| Mana | [Conflagrate](https://www.wowhead.com/forever/spell=18932) | -6349.7 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -5669.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5380.3 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5113.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4380.6 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3593.9 |
| Mana | OtherActionManaRegen (tag 1) | +3101.1 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -490.5 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `453203102-053030031005102251`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 816.50 | 0.58 | 0.00 |
| Tauren | 814.69 | 0.58 | 0.00 |
| Troll | 806.56 | 0.58 | 0.00 |
| Windshaper | 817.75 | 0.57 | 0.00 |
| Dwarf | 816.25 | 0.58 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit — Head (Mail) [+18 Sta, +23 Int, +29 SP, +14 Crit, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=272685) | 65 | Lesser Arcanum of Voracity |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 245.73 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 154.85 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 101.75 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 79.52 |
| [Fire Nova](https://www.wowhead.com/forever/spell=408345) | 63.37 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 48.76 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 47.46 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 28.70 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Fire Nova](https://www.wowhead.com/forever/spell=408345) | -12447.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +9343.1 |
| Mana | OtherActionManaRegen (tag 1) | +7289.9 |
| Mana | [Earth Shock](https://www.wowhead.com/forever/spell=10414) | -7017.9 |
| Mana | [Stormstrike](https://www.wowhead.com/forever/spell=17364) | -4648.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3686.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3576.5 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -2237.6 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |
| Mana | [Spell 10614](https://www.wowhead.com/forever/spell=10614) | -500.0 |
| Mana | [Strength of Earth Totem](https://www.wowhead.com/forever/spell=10442) | -450.0 |
| Mana | OtherActionManaRegen (tag 2) | +163.7 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `254000001351-0025003001-0550005103`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 794.70 | 0.51 | 0.00 |
| Troll | 778.14 | 0.50 | 0.00 |
| Undead | 795.22 | 0.50 | 0.00 |
| Human | 804.08 | 0.51 | 0.00 |
| Gnome | 789.85 | 0.51 | 0.00 |

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
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 456.53 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 124.76 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 106.78 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 88.44 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 21.41 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 3.86 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.31 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | -31348.4 |
| Mana | [Life Tap](https://www.wowhead.com/forever/spell=11689) | +22372.4 |
| Mana | [Immolate](https://www.wowhead.com/forever/spell=25309) | -6904.4 |
| Mana | [Corruption](https://www.wowhead.com/forever/spell=25311) | -5478.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5412.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3607.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3546.0 |
| Mana | OtherActionManaRegen (tag 1) | +3117.4 |
| Mana | [Bane of Doom](https://www.wowhead.com/forever/spell=603) | -1200.0 |
| Mana | [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | -575.3 |
| Mana | [Shadowburn](https://www.wowhead.com/forever/spell=18871) | -362.7 |
| Mana | [Searing Pain](https://www.wowhead.com/forever/spell=17923) | -192.1 |

## Mage — Fire

**Talents:** 18/31/2 · `05001252000003-23450000130135031-002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Talent variants:** The allocation above is for Human. Other races retain their saved allocations in the exact profiles. Historical validation ranges and source qualifications are recorded separately.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 784.13 | 0.55 | 6.46 |
| Troll | 775.56 | 0.52 | 5.07 |
| Undead | 783.70 | 0.55 | 6.07 |
| Human | 794.40 | 0.54 | 5.28 |
| Gnome | 787.54 | 0.56 | 6.03 |
| High Order | 757.70 | 0.56 | 10.38 |

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
| Waist | [Modeled: Spell power / crit / hit (matched cost) — Waist (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 56 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Legs | [Modeled: Spell power / crit / hit (matched cost) — Legs (Cloth) [+18 Sta, +21 Crit, +23 Hit, +23 SP, 88 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Focus |
| Feet | [Modeled: Spell power / crit / hit (matched cost) — Feet (Cloth) [+14 Sta, +16 Crit, +18 Hit, +18 SP, 69 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Spirit |
| Ring 1 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Spell power / crit / hit (matched cost) — Finger [+11 Sta, +12 Crit, +13 Hit, +13 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Spell power / crit / hit (matched cost) — Trinket [+11 Sta, +13 Crit, +14 Hit, +14 SP] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Caster dagger one-hand (projected) — One-hand [+11 Sta, +10 Int, +12 SP, 66–123 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Spell Power |
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

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 260.00 |
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 247.89 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 156.41 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 117.35 |
| [Item 15993](https://www.wowhead.com/forever/item=15993) | 8.64 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 4.11 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | OtherActionManaRegen (tag 1) | +14983.8 |
| Mana | [Scorch](https://www.wowhead.com/forever/spell=10207) | -14637.6 |
| Mana | [Fire Blast](https://www.wowhead.com/forever/spell=10199) | -13442.7 |
| Mana | [Pyroblast](https://www.wowhead.com/forever/spell=18809) | -12349.1 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +5445.7 |
| Mana | [Master of Elements](https://www.wowhead.com/forever/spell=29076) | +5266.2 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4843.8 |
| Mana | OtherActionManaRegen (tag 2) | +1788.1 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1099.4 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +850.2 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +600.3 |
| Mana | [Fireball](https://www.wowhead.com/forever/spell=25306) | -312.3 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 760.21 | 0.32 | 0.00 |
| Troll | 754.77 | 0.32 | 0.00 |
| Undead | 759.55 | 0.32 | 0.00 |
| Windshaper | 773.58 | 0.33 | 0.00 |
| Human | 762.37 | 0.32 | 0.00 |
| Dwarf | 761.82 | 0.33 | 0.00 |
| Night Elf | 758.40 | 0.32 | 0.00 |
| Gnome | 758.33 | 0.33 | 0.00 |
| High Order | 773.58 | 0.33 | 0.00 |

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

- -1s: [Blade Flurry](https://www.wowhead.com/forever/spell=13877).

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) time remaining < 3 AND Time remaining > 9)).
3. Cast [Adrenaline Rush](https://www.wowhead.com/forever/spell=13750) when Energy ≤ 40.
4. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
5. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when Combo points ≥ 3.
6. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 3.
7. Cast [Sinister Strike (rank 8)](https://www.wowhead.com/forever/spell=11294) when (`{"autoTimeToNext":{"autoType":"MainHand"}}` > `{"autoSwingTime":{"autoType":"MainHand"}}` − 0.5 OR Time remaining < 6 OR Energy ≥ 79).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 200.20 |
| [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | 141.80 |
| Auto-attack (tag 2) | 134.34 |
| Auto-attack (tag 3) | 108.88 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 90.50 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 39.41 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 25.35 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 13.26 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3271.4 |
| Energy | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | -3022.8 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -480.8 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +473.1 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -468.8 |
| Energy | OtherActionRefund | +125.6 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +96.7 |
| Energy | [Blade Flurry](https://www.wowhead.com/forever/spell=13877) | -75.0 |
| ComboPoints | [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | +72.2 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -68.9 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -26.5 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +24.9 |

## Rogue — Mutilate

**Talents:** 31/20/0 · `0053031035140105-302303202014`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 749.44 | 0.31 | 0.00 |
| Troll | 741.17 | 0.31 | 0.00 |
| Undead | 749.99 | 0.29 | 0.00 |
| Windshaper | 761.62 | 0.31 | 0.00 |
| Human | 752.72 | 0.31 | 0.00 |
| Dwarf | 748.74 | 0.30 | 0.00 |
| Night Elf | 747.97 | 0.30 | 0.00 |
| Gnome | 743.11 | 0.30 | 0.00 |
| High Order | 761.62 | 0.31 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

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
| Legs | [Modeled: Agility / crit / hit — Legs (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 173 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Feet | [Modeled: Agility / crit / hit — Feet (Leather) [+14 Sta, +16 Agi, +16 Crit, +18 Hit, 136 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Enchant Boots - Greater Agility |
| Ring 1 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Ring 2 | [Modeled: Agility / crit / hit — Finger [+11 Sta, +12 Agi, +12 Crit, +13 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 1 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Trinket 2 | [Modeled: Passive Agility / crit / hit — Trinket [+11 Sta, +13 Agi, +13 Crit, +14 Hit] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | — |
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 82–154 dmg @ 2.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Fiery Weapon |
| Off hand | [Modeled: Dagger off-hand (unverified) — Off hand weapon [+8 Sta, +7 Int, +15 Crit, 56–105 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Agility |
| Ranged/relic | [Modeled: Crossbow ranged — Ranged/wand [+6 Sta, +32 AP, 98–148 dmg @ 2s] (modeled; reference only)](https://www.wowhead.com/forever/item=272595) | 65 | SAF-T Ultra Precision Scope |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) time remaining < 3 AND Time remaining > 9)).
3. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
4. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when (Combo points ≥ 4 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active).
5. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 3.
6. Cast [Mutilate](https://www.wowhead.com/forever/spell=1241584).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 178.30 |
| Auto-attack (tag 2) | 113.41 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 88.13 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 78.83 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 66.99 |
| Auto-attack (tag 3) | 58.16 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 51.09 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 44.32 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3103.3 |
| Energy | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | -3019.5 |
| Energy | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -882.7 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +830.1 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -406.0 |
| Energy | OtherActionRefund | +204.2 |
| ComboPoints | [Spell 31016](https://www.wowhead.com/forever/spell=31016) | -144.3 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +99.3 |
| ComboPoints | [Mutilate](https://www.wowhead.com/forever/spell=1241584) | +94.1 |
| ComboPoints | [Seal Fate](https://www.wowhead.com/forever/spell=14195) | +39.8 |
| ComboPoints | [Ruthlessness](https://www.wowhead.com/forever/spell=14161) | +34.5 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -22.4 |

## Druid — Balance

**Talents:** 38/0/13 · `5232220115501351--505003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 715.40 | 0.34 | 0.00 |
| Windshaper | 723.39 | 0.32 | 0.00 |
| Night Elf | 713.21 | 0.32 | 0.00 |
| High Order | 723.39 | 0.32 | 0.00 |

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
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 423.17 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 106.68 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 97.18 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 96.36 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Starfire](https://www.wowhead.com/forever/spell=25298) | -17972.6 |
| Mana | OtherActionManaRegen (tag 1) | +7682.0 |
| Mana | [Moonfire](https://www.wowhead.com/forever/spell=9835) | -4833.9 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4653.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +3718.4 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3603.8 |
| Mana | [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | -1662.9 |
| Mana | [Wrath](https://www.wowhead.com/forever/spell=9912) | -1263.3 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | +947.6 |
| Mana | [Moonkin Form](https://www.wowhead.com/forever/spell=24858) | -435.4 |
| Mana | OtherActionManaRegen (tag 2) | +55.0 |
| Mana | [Innervate](https://www.wowhead.com/forever/spell=29166) | -27.4 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 705.23 | 0.27 | 0.00 |
| Troll | 700.52 | 0.27 | 0.00 |
| Undead | 703.70 | 0.26 | 0.00 |
| Windshaper | 717.64 | 0.27 | 0.00 |
| Human | 705.74 | 0.28 | 0.00 |
| Dwarf | 703.68 | 0.27 | 0.00 |
| Night Elf | 702.94 | 0.27 | 0.00 |
| Gnome | 704.63 | 0.28 | 0.00 |
| High Order | 717.64 | 0.27 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Windshaper

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit — Head (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Arcanum of Rapidity |
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
| Main hand | [Modeled: Dagger main-hand — Main hand [+8 Sta, +7 Int, +15 Crit, 82–154 dmg @ 2.5s] (modeled; reference only)](https://www.wowhead.com/forever/item=272683) | 65 | Enchant Weapon - Crusader |
| Off hand | [Modeled: Fist off-hand — Off hand weapon [+2 Sta, +24 AP, +14 Crit, 64–97 dmg @ 1.7s] (modeled; reference only)](https://www.wowhead.com/forever/item=272598) | 65 | Enchant Weapon - Crusader |
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

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 190.47 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 154.89 |
| Auto-attack (tag 2) | 102.37 |
| Auto-attack (tag 3) | 64.72 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 42.19 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 41.41 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 31.47 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 22.66 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Energy | OtherActionEnergyRegen | +3104.4 |
| Energy | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | -2933.4 |
| Energy | [Relentless Strikes](https://www.wowhead.com/forever/spell=14179) | +561.3 |
| Energy | [Rupture](https://www.wowhead.com/forever/spell=11275) | -490.8 |
| Energy | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -334.9 |
| Energy | OtherActionRefund | +188.4 |
| Energy | [Ghostly Strike](https://www.wowhead.com/forever/spell=14278) | -153.5 |
| Energy | [Item 7676](https://www.wowhead.com/forever/item=7676) | +98.4 |
| ComboPoints | [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | +87.6 |
| Energy | [Ambush](https://www.wowhead.com/forever/spell=11269) | -83.5 |
| ComboPoints | [Rupture](https://www.wowhead.com/forever/spell=11275) | -65.4 |
| ComboPoints | [Slice and Dice](https://www.wowhead.com/forever/spell=6774) | -41.1 |

## Mage — Frost

**Talents:** 12/3/36 · `0500105001-03-0555003331000301241`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 639.34 | 0.41 | 9.39 |
| Troll | 627.56 | 0.40 | 11.38 |
| Undead | 633.93 | 0.41 | 10.28 |
| Human | 643.43 | 0.40 | 7.79 |
| Gnome | 652.66 | 0.36 | 3.08 |
| High Order | 635.12 | 0.40 | 12.91 |

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
6. Cast [Frostbolt (rank 11)](https://www.wowhead.com/forever/spell=25304).

### Damage breakdown — Gnome

| Action | DPS |
|---|---:|
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 419.62 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 136.29 |
| [Cone of Cold](https://www.wowhead.com/forever/spell=10161) | 96.75 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Frostbolt](https://www.wowhead.com/forever/spell=25304) | -19069.7 |
| Mana | OtherActionManaRegen (tag 1) | +9323.2 |
| Mana | [Cone of Cold](https://www.wowhead.com/forever/spell=10161) | -7331.3 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4767.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3910.5 |
| Mana | [Ice Lance](https://www.wowhead.com/forever/spell=30455) | -3700.2 |
| Mana | OtherActionManaRegen (tag 2) | +2448.1 |
| Mana | [Item 8008](https://www.wowhead.com/forever/item=8008) | +1100.2 |
| Mana | [Item 8007](https://www.wowhead.com/forever/item=8007) | +850.2 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +775.9 |
| Mana | [Item 5513](https://www.wowhead.com/forever/item=5513) | +212.1 |

## Druid — Bear · Tank

**Talents:** 1/40/10 · `01-50022332120132012551-055`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Bear's outgoing base Rage remains the qualified damage-based model (DRU-012), with sourced double critical Rage. Incoming Rage now shares Warrior's configurable provisional expected-health scenario. The form threat multiplier is sourced at 1.5; this is not a survival ranking. Proc-enchant form eligibility is supported. Recovery has a sourced heal/ICD; Revelation requires an explicit, unverified proc model. Historical enchant-selection gains do not validate those assumptions.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Tauren | 552.02 | 0.37 | 0.00 |
| Windshaper | 551.47 | 0.37 | 0.00 |
| Night Elf | 545.43 | 0.37 | 0.00 |
| High Order | 551.47 | 0.37 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Tauren | 1899.14 | 717.56 | 73.49 | 0.28% |
| Windshaper | 1888.92 | 715.66 | 78.42 | 0.48% |
| Night Elf | 1876.29 | 700.30 | 77.88 | 0.50% |
| High Order | 1888.92 | 715.66 | 78.42 | 0.48% |

These stress-scenario results are not measured boss balance or an equal-support survival ranking.

### Equipment — Tauren

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Agility / crit / hit (leather) — Head (Leather) [+18 Sta, +21 Agi, +21 Crit, +23 Hit, 160 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=279253) | 65 | Lesser Arcanum of Voracity |
| Neck | [Modeled: Strength / agility / crit / hit — Neck [+16 Str, +11 Sta, +8 Crit, +6 Hit, +10 Agi] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | — |
| Shoulders | [Modeled: Agility / AP / crit — Shoulders (Leather) [+14 Sta, +7 Agi, +45 AP, +3 Crit, 148 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=20068) | 65 | Chromatic Mantle of the Dawn |
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

### Damage breakdown — Tauren

| Action | DPS |
|---|---:|
| [Maul](https://www.wowhead.com/forever/spell=9881) | 229.41 |
| [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073) | 115.33 |
| [Lacerate](https://www.wowhead.com/forever/spell=414644) | 91.59 |
| Auto-attack (tag 1) | 51.68 |
| [Lacerate](https://www.wowhead.com/forever/spell=414647) | 34.19 |
| Auto-attack (tag 3) | 27.62 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.20 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -216626.0 |
| Health | OtherActionHealingModel | +214222.9 |
| Rage | [Lacerate](https://www.wowhead.com/forever/spell=414644) | -1269.2 |
| Mana | [Faerie Fire](https://www.wowhead.com/forever/spell=9907) | -929.0 |
| Rage | Auto-attack (tag 1) | +892.1 |
| Rage | [Maul](https://www.wowhead.com/forever/spell=9881) | -870.1 |
| Rage | OtherActionDamageTaken | +812.9 |
| Rage | [Primal Bite (Bear)](https://www.wowhead.com/forever/spell=1238073) | -787.3 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +733.5 |
| Health | [Holy Strength](https://www.wowhead.com/forever/spell=20007) | +561.1 |
| Rage | [Blood Frenzy](https://www.wowhead.com/forever/spell=37117) | +559.1 |
| Rage | OtherActionRefund | +424.7 |

## Shaman — Elemental

**Talents:** 31/9/11 · `5503331000011351-054-05303`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 533.92 | 0.35 | 31.21 |
| Tauren | 531.37 | 0.36 | 31.76 |
| Troll | 526.11 | 0.34 | 33.02 |
| Windshaper | 525.69 | 0.36 | 37.85 |
| Dwarf | 532.06 | 0.36 | 31.67 |

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
| Ranged/relic | [Projected: Polished Driftwood Icon — Relic [+2 Sta, +1.75 Int, +2.25 SP]](https://www.wowhead.com/forever/item=920001012) | 65 | — |

Other races can use different equipment. Their complete setups are available in the simulator's Ranked builds selector.

### Before the pull

- -2s: [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Rotation priorities

1. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Target count = 1 AND NOT [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) DoT active).
2. Cast [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) when (Target count ≥ 2 AND NOT [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) DoT active).
3. Use ready automatic cooldowns.
4. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active.
5. Cast [Lava Burst (rank 4)](https://www.wowhead.com/forever/spell=1326047).
6. Cast [Chain Lightning (rank 4)](https://www.wowhead.com/forever/spell=10605) when ([Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active OR Target count ≥ 2).
7. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when Mana ≥ Time remaining × 35.
8. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Damage breakdown — Orc

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 295.11 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1326047) | 97.86 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 68.84 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 35.00 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 21.18 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 14.82 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.10 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -18637.7 |
| Mana | OtherActionManaRegen (tag 1) | +7777.8 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5206.3 |
| Mana | [Lava Burst](https://www.wowhead.com/forever/spell=1326047) | -5008.0 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +4441.1 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -4140.8 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3597.8 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -867.0 |
| Mana | OtherActionManaRegen (tag 2) | +649.5 |

## Paladin — Protection · Tank

**Talents:** 0/42/9 · `-5520513321301551-15102`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tank replay evidence](../artifacts/tanks/current/validation.json)

**Model limitations:** Tank encounter: frontal attacks, incoming boss damage and modeled healing. Seal of Fury is now selectable with sourced damage, shield requirement and talent-only depletion mana. Its shared replacement shield pool and independent lifetime are provisional; the primary seal and lifetime are recorded, not silently substituted. Threat-based taunt target selection is not modeled. Gear selections must pass both workload guard policies, which permit limited mitigation tradeoffs. This is not a survival ranking. Reflections use explicit provider power, never the receiving tank's spell power.

**Talent variants:** The allocation above is for Undead. Other races retain their saved allocations in the exact profiles. Historical validation ranges and source qualifications are recorded separately.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Undead | 531.22 | 0.26 | 0.05 |
| Human | 530.22 | 0.25 | 0.12 |
| Dwarf | 523.10 | 0.26 | 0.09 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Undead | 898.44 | 704.02 | 77.23 | 0.16% |
| Human | 916.72 | 702.40 | 78.36 | 0.18% |
| Dwarf | 899.74 | 704.85 | 77.87 | 0.36% |

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
| Auto-attack (tag 1) | 104.58 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 70.40 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 68.47 |
| [Holy Shield](https://www.wowhead.com/forever/spell=20957) | 64.99 |
| [Consecration](https://www.wowhead.com/forever/spell=20924) | 64.53 |
| Auto-attack (tag 3) | 47.09 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 35.75 |
| [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | 22.38 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -211205.4 |
| Health | OtherActionHealingModel | +208887.6 |
| Mana | [Consecration](https://www.wowhead.com/forever/spell=20924) | -17977.6 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +16708.5 |
| Mana | [Shield Specialization](https://www.wowhead.com/forever/spell=20148) | +8643.3 |
| Mana | [Holy Shield](https://www.wowhead.com/forever/spell=20928) | -6479.8 |
| Mana | [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | -3101.9 |
| Mana | [Judgement of Light](https://www.wowhead.com/forever/spell=20271) | -2682.1 |
| Mana | OtherActionManaRegen (tag 1) | +2357.0 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +2305.9 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +2302.3 |
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
| Orc | 501.82 | 0.35 | 0.00 |
| Tauren | 513.71 | 0.36 | 0.00 |
| Troll | 504.85 | 0.35 | 0.00 |
| Undead | 516.64 | 0.31 | 0.00 |
| Windshaper | 509.94 | 0.36 | 0.00 |
| Human | 508.16 | 0.37 | 0.00 |
| Dwarf | 507.38 | 0.36 | 0.00 |
| Night Elf | 497.20 | 0.33 | 0.00 |
| Gnome | 507.11 | 0.35 | 0.00 |
| High Order | 509.94 | 0.36 | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Tank metrics

| Race | TPS | DTPS | TMI | Modeled death probability |
|---|---:|---:|---:|---:|
| Orc | 1256.72 | 504.55 | 51.39 | 0.00% |
| Tauren | 1261.20 | 480.10 | 46.54 | 0.00% |
| Troll | 1242.66 | 490.24 | 49.93 | 0.00% |
| Undead | 1202.79 | 465.18 | 49.47 | 0.00% |
| Windshaper | 1197.56 | 486.02 | 49.91 | 0.00% |
| Human | 1189.36 | 472.42 | 49.48 | 0.00% |
| Dwarf | 1268.80 | 496.81 | 50.56 | 0.00% |
| Night Elf | 1164.06 | 481.44 | 49.83 | 0.00% |
| Gnome | 1194.00 | 481.96 | 49.49 | 0.00% |
| High Order | 1197.56 | 486.02 | 49.91 | 0.00% |

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
| [Revenge](https://www.wowhead.com/forever/spell=25288) | 162.48 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 117.19 |
| [Shield Slam](https://www.wowhead.com/forever/spell=23925) | 91.49 |
| Auto-attack (tag 1) | 79.22 |
| Auto-attack (tag 3) | 37.65 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 12.51 |
| [Thunder Clap](https://www.wowhead.com/forever/spell=11581) | 7.28 |
| [Fiery Weapon](https://www.wowhead.com/forever/spell=13897) | 6.80 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Health | OtherActionDamageTaken | -139552.6 |
| Health | OtherActionHealingModel | +137858.3 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1693.4 |
| Rage | Auto-attack (tag 1) | +908.8 |
| Rage | [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | -733.0 |
| Rage | [Shield Slam](https://www.wowhead.com/forever/spell=23925) | -687.4 |
| Rage | OtherActionDamageTaken | +527.0 |
| Rage | [Shield Block](https://www.wowhead.com/forever/spell=2565) | -511.1 |
| Rage | [Shield Specialization](https://www.wowhead.com/forever/spell=12727) | +428.5 |
| Rage | OtherActionRefund | +247.6 |
| Rage | [Sunder Armor](https://www.wowhead.com/forever/spell=11597) | -189.7 |
| Rage | [Shield Specialization](https://www.wowhead.com/forever/spell=12727) | +161.9 |

## Shaman — Stormcaller

**Talents:** 29/22/0 · `553233100000034-055030021004002`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Uses maximum registered Lightning Bolt. Overload uses its own client rank rows; Maelstrom/proc and server-script questions remain qualified.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Orc | 468.59 | 0.31 | 60.05 |
| Tauren | 473.35 | 0.32 | 58.78 |
| Troll | 461.76 | 0.29 | 58.57 |
| Windshaper | 461.33 | 0.30 | 64.69 |
| Dwarf | 469.18 | 0.30 | 59.72 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Equipment — Tauren

| Slot | Item | Item level | Enchant |
|---|---|---:|---|
| Head | [Modeled: Intellect / spell power / crit / hit (matched cost) — Head (Mail) [+18 Sta, +14 Crit, +10 Hit, +31 SP, +18 Int, 338 armor] (modeled; reference only)](https://www.wowhead.com/forever/item=22651) | 65 | Arcanum of Focus |
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
| Ranged/relic | [Projected: Polished Driftwood Icon — Relic [+2 Sta, +1.75 Int, +2.25 SP]](https://www.wowhead.com/forever/item=920001012) | 65 | — |

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
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 322.33 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 76.28 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 34.75 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 22.64 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 16.19 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.17 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | -19426.9 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +5250.2 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3939.1 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3598.4 |
| Mana | OtherActionManaRegen (tag 1) | +3382.4 |
| Mana | [Flame Shock](https://www.wowhead.com/forever/spell=29228) | -2360.1 |
| Mana | OtherActionManaRegen (tag 2) | +1363.0 |
| Mana | [Searing Totem](https://www.wowhead.com/forever/spell=10438) | -1020.0 |

## Priest — Smite

**Talents:** 31/17/3 · `515035031030001031-00505023002-003`

[Requests and results](../artifacts/modelled_gear/forever_dps_5min.json) · [Equipment search](../artifacts/correctness/october_15/selection-evidence.json)

**Model limitations:** Uses maximum registered ranks. Penance's static rank anomaly and actual skill-book acquisition remain source gaps; no efficient lower-rank fallback is assumed.

**Rotation variants:** The priorities below are for Undead. Other races may use a different saved APL; their exact rotations are in the Ranked builds selector and raw requests.


### Results

| Race | DPS | Standard error | Mana-limited seconds |
|---|---:|---:|---:|
| Troll | 442.22 | 0.20 | 62.17 |
| Undead | 456.78 | 0.19 | 56.88 |
| Human | 450.88 | 0.20 | 58.03 |
| Dwarf | 447.13 | 0.20 | 59.93 |
| Night Elf | 446.12 | 0.20 | 60.46 |
| Gnome | 456.10 | 0.20 | 55.17 |

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
| [Smite](https://www.wowhead.com/forever/spell=10934) | 256.93 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 122.79 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 69.40 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 7.66 |

### Resource flow

| Resource | Action | Net amount per fight |
|---|---|---:|
| Mana | [Smite](https://www.wowhead.com/forever/spell=10934) | -17946.3 |
| Mana | OtherActionManaRegen (tag 1) | +7399.5 |
| Mana | [Holy Fire](https://www.wowhead.com/forever/spell=15261) | -6185.9 |
| Mana | [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | -4930.1 |
| Mana | [Item 13444](https://www.wowhead.com/forever/item=13444) | +4135.7 |
| Mana | [Item 12662](https://www.wowhead.com/forever/item=12662) | +3604.8 |
| Mana | [Judgement of Wisdom](https://www.wowhead.com/forever/spell=20355) | +3243.7 |
| Health | OtherActionDamageTaken | -1600.0 |
| Mana | [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328) | +1600.0 |
| Health | [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | +1574.9 |
| Mana | OtherActionManaRegen (tag 2) | +1438.5 |
| Mana | [Power Infusion](https://www.wowhead.com/forever/spell=10060) | -495.4 |
