# Shaman beta pass (17 September 2026)

Beta client `1.60.1.69893` against Classic Era `1.15.9.69722`, read with `tools/data_watch/spell_client.py`, the
talent curves with `tools/data_watch/trait_curve.mjs`, `SpellAuraOptions` for charges and proc recovery, and
`../beta/shaman.json`.

## Weapon layouts

Shamans cannot dual wield. The engine and gear picker allow a two-handed weapon,
or a one-handed weapon with a shield or held off-hand item. Imports and weapon
swaps must obey the same restriction. Swing synchronization and off-hand weapon
EP have been removed.

Earlier Enhancement comparisons used dual wield and are invalid. Their search
evidence is retained under `artifacts/optimization/archive/`; the current
starting preset uses Forest Defender's Axe. A legal-layout rerun is separate
from the wider equipment and enchant selection.

## How the numbers were read

A rank's damage in the sim is the client's base plus `EffectRealPointsPerLevel` for every level from the rank's base
level up to its max level, capped at 60, low end rounded down and high end up (single values rounded to nearest). Against
the Era client that rule reproduces the sim's old tables for every max rank and most lower ranks; the ones it did not
(Lightning Bolt 7-9, Earth Shock 4-6, Frost Shock 3, Frostbrand 1-4) were already off in the sim and now follow the rule.
Forever stores a range as `EffectBasePointsF` with `Variance`, and the per-level points are added before rounding.

Flame Shock's dot is stored per tick in both clients (4 ticks of 3 sec); the sim's table holds the four ticks' total and
applies the client's 0.1 per tick, so no coefficient conversion was needed.

## Spells changed (old -> new, level 60 rank unless noted)

| Spell | Change |
|---|---|
| Lightning Bolt (10 ranks) | rank 10 428-477 -> 190-212, 3 -> 2.5 sec cast (ranks 4-10), 265 -> 220 mana; coefficients .123/.314/.554/.857 -> .429/.571/.714 from rank 3 |
| Chain Lightning (4 ranks) | rank 4 505-564 -> 119-134, 2.5 -> 2 sec cast, 605 -> 485 mana, .714 -> .571 (rank 3 .517, as the client stores it) |
| Earth Shock (7 ranks) | rank 7 517-545 -> 293-309; ranks 1-3 coefficients .154/.212/.299 -> .386; costs unchanged |
| Flame Shock (6 ranks) | rank 6 292 + 320 over 12 sec -> 166 + 176; ranks 1-2 coefficients .134/.063 and .198/.093 -> .214/.1 |
| Frost Shock (4 ranks) | rank 4 492-520 -> 278-295; .386 unchanged |
| Lightning Shield orbs (7 ranks) | rank 7 198 -> 178, ranks 2-6 lower; ranks 1-2 coefficients .147/.227 -> .267. Rank 7's 370 mana was missing from the sim's table (it cost 0) |
| Searing Totem attack (6 ranks) | damage unchanged, coefficient .052/.083 -> .017 |
| Magma Totem pulse (4 ranks) | 75 -> 73 (22/37/54 -> 20/35/52) |
| Fire Nova Totem -> Fire Nova | Fire Nova Totem (1535, 8498, 8499, 11314, 11315) no longer exists. The spellbook teaches Fire Nova 408341-408345: instant, 10 sec cooldown (was 15, after a 5 sec totem delay), the totem's mana costs, rank 5 413-459 at .143 -> 397-443 at .214, hitting at once. It is a spell now, not a fire totem. |
| Lava Burst | guessed single spell 51505 (158-187, 2 sec, 8 sec cd, 10% base mana, .5714) -> three ranks 408490 / 1238299 / 1238300 at 40/50/60: 105-135 / 165-211 / 192-248, 2.5 sec cast, 10 sec cooldown, 165/230/265 mana, .714; Flame Shock bonus 20% confirmed |
| Stormstrike (Forever ruleset) | 20 -> 8 sec cooldown, 21% base mana -> 125 mana; +20% to all Nature damage for 12 sec -> +20% to the next Lightning Bolt, Chain Lightning or Earth Shock within 12 sec (one charge) |
| Healing Stream (5 ranks) | 14 -> 11 a tick |
| Frostbrand Attack | rank 5 187 -> 169 (.1 unchanged) |
| Windfury Totem buff | 122/229/315 -> 95/179/246 attack power (table in `air_totems.go` only, see core changes) |
| Strength of Earth / Grace of Air (own totems) | 77 x 1.15 = 88 Strength -> 53; 88 Agility -> 89 |
| Every totem | global cooldown 1.5 -> 1 sec; Strength of Earth, Stoneskin, Tremor, Windfury, Grace of Air, Windwall, Healing Stream and Mana Spring last 5 min (were 2 or 1 min). Searing and Magma durations unchanged |
| Water Shield | id 52127 -> 408510, 6% base mana -> free, adds a 15 sec cooldown; 3 globes, 2% and 3.5 sec confirmed |
| Rage of the Farseer | id 2825 -> 425336; 30%, 25 sec, 3 min confirmed |
| Improved Stormstrike | buff id 51521 -> 1238931; rank 2 regeneration 100% -> 50% (the buff is 50% at both ranks, only the chances scale) |
| Maelstrom Weapon | buff id 51530 -> 408505; 5 stacks, 30 sec, 4% per point confirmed |
| Elemental Alacrity | 0.17/0.34/0.51 -> 0.17/0.33/0.5 sec |

Unchanged in the client: shock mana costs and the 6 sec shock cooldown, Chain Lightning's 6 sec cooldown and 30% bounce
falloff, Lightning Shield's 3 charges and 3.5 sec, Mana Spring (4-10 mana every 2 sec), Rockbiter, Flametongue and
Windfury Weapon proc spells, Elemental Devastation (3% per point, 10 sec), Elemental Focus (10%, 15 sec), Nature's
Swiftness (3 min), Flurry's talent curve 5-25% (already in the sim).

The APLs follow the new ids: the elemental default casts Lava Burst rank 3 (1238300), the enhancement default reads
Maelstrom Weapon stacks from 408505. `ui/core/spells/shaman.json` has no unreviewed entries left (106 reviewed, budget
762 -> 656): changed spells are `forever` with tooltips; Searing Totem, Flametongue Weapon, Windfury Weapon and Elemental
Devastation are `classic`. The Stormstrike entry in `ui/core/spells/core.json` now carries the client's text.

## Checklist lines

### Shaman section
- **Open** `sim/shaman/air_totems.go:50`, Windfury Totem's value is hard-coded in `sim/core/buffs.go`. The client
  lowers it to 95/179/246 attack power (Classic 122/229/315); the shaman table now holds those, but the buff the sim
  applies is core's 315 until `buffs.go` changes.
- **Resolved** `sim/shaman/lava_burst.go:11`: cast 2 -> 2.5 sec, cooldown 8 -> 10 sec, 10% base mana -> 265 flat, .5714
  -> .714, 158-187 -> 192-248 at rank 3, and the spell is three ranks.
- **Resolved** `sim/shaman/lightning_overload.go:25`: client curve 3/7/10, as the sim had.
- **Resolved** `sim/shaman/talents.go:68`, Improved Reincarnation: client curve 2/4%, unchanged.
- **Resolved** `sim/shaman/talents.go:139`, Elemental Alacrity: 0.17/0.34/0.51 -> 0.17/0.33/0.5 sec.
- **Resolved** `sim/shaman/talents.go:145`, Improved Fire Nova: client curve 10/20% and 2/4 sec, unchanged. It now
  shortens Fire Nova's 10 sec cooldown (was the totem's 15).
- **Resolved** `sim/shaman/talents.go:252`, Elemental Fury: client curve 20-100%, unchanged.
- **Resolved** `sim/shaman/talents.go:420`, Improved Stormstrike chances: client curve 50/100% for both rolls, unchanged.
  The regeneration was wrongly doubled at rank 2: 100% -> 50%.
- **Resolved** `sim/shaman/talents.go:428`, Improved Stormstrike window: the buff (1238931) lasts 15 sec at both ranks.
- **Open** `sim/shaman/talents.go:463`, Maelstrom Weapon proc rate: the 4% per point is confirmed by the curve, but
  408498 has no procs-per-minute entry and a 100% proc chance, so the client does not say how often it procs. The sim
  keeps 2 PPM per point.
- **Resolved** `sim/shaman/talents.go:470`, Maelstrom Weapon shape: 5 stacks (408498) and 30 sec (408505) at every rank.
- **Resolved** `sim/shaman/talents.go:515`, Rage of the Farseer: 3 min cooldown (425336), unchanged.
- **Resolved** `sim/shaman/totems.go:9`, Enhancing Totems baseline: not a flat 15%. Grace of Air is 89 (88 before),
  Strength of Earth is 53 (88 before).
- **Resolved** `sim/shaman/water_shield.go:18`: 3.5 sec between globes (`ProcCategoryRecovery`), unchanged; cost 6% base
  mana -> free, and a 15 sec cooldown added.
- **Open** `sim/shaman/water_totems.go:115`, Mana Spring's value is hard-coded in `sim/core/buffs.go`. The client leaves
  it at 10 mana every 2 sec, so only the structural TODO remains.
- **Open** `sim/shaman/windfury_weapon.go:78`: both clients give each imbued weapon its own 20% enchant proc (enchant
  1669 -> 439431), but which hand's attacks the extra swings use is server behaviour the data does not show.

### Talents the sim does not read
- **Open** Shaman / Restoration: Riptide stays not simulated. The client gives its numbers (rank 1 408521: 456-504
  heal at .214, 89 every 3 sec for 15 sec at .1, +25% to Chain Heal on the target, 6 sec cooldown, 245 mana; ranks 2-3
  1239242/1239243 at 50/60), but it is healing only and the restoration spec is not built (`sim/shaman/_restoration`).

### Baseline ability changes (spellbook diff, `--learned shaman` vs `--learned shaman --era`)
`SkillLineAbility` lists Season of Discovery rune and engraving spells in both builds, so only the differences count:
- **Removed**: Fire Nova Totem ranks 1-5 (the ids are gone from the client entirely). Replaced by Fire Nova, now modelled
  in `sim/shaman/fire_totems.go`. Also gone: the Windfury Totem Effect spells 8514/10607/10611 and Flurry buff ranks
  16277-16280 (the talent keeps one buff id).
- **New, damage**: Lava Burst ranks 2 and 3 (1238299 at 50, 1238300 at 60), modelled.
- **New, not modelled**: Totemic Recall (36936, level 20), Call of the Elements / Ancestors / Spirits (66842-66844, levels
  20/30/40, drop several totems in one 3 sec cast), Riptide ranks 2-3. All utility or healing.
- **Changed level**: Water Shield (level 20), Totemic Projection (level 22, 1 min cooldown, 25% base mana), Lava Burst rank
  1 (level 40). Shamanistic Rage's id 425336 is now Rage of the Farseer.
- Talent spells collapsed to single ids (Elemental Warding, Eye of the Storm, Ancestral Healing, Healing Way); Storm Reach
  is Elemental Reach and Healing Grace is Natural Grace.

## Core changes this needs (not made here)
- `sim/core/buffs.go`: Windfury Totem attack power 315 -> 246; Strength of Earth 77 -> 53 and Grace of Air's improved
  value 88 -> 89 for the raid-buff path (the shaman's own totems pass multipliers that land on these); the Strength of
  Earth and Grace of Air totem auras last 2 min where the client says 5 (the shaman package overrides the duration on
  its own copies).
- `ui/core/components/inputs/totem_inputs.ts` still offers Fire Nova Totem (11315) as a fire totem.

## DPS (Phase 1 Average-Default)

| Test | Before | After |
|---|---|---|
| TestElemental | 470.2 | 397.2 |
| TestForeverElemental | 483.6 | 409.2 |
| TestForeverStormcaller | 456.9 | 389.6 |
| TestEnhancement | 707.6 | 660.0 |
| TestForeverEnhancement | 767.8 | 720.2 |
