# Paladin beta pass (17 September 2026)

## Talent availability correction

Seal of Command requires its talent point. In client `1.60.1.69893`,
[`TraitDefinition`](https://wago.tools/db2/TraitDefinition/csv?build=1.60.1.69893)
135226 references spell 20375. Its
[`TraitNodeEntry`](https://wago.tools/db2/TraitNodeEntry/csv?build=1.60.1.69893)
130425 has one rank and is attached by link 127892 to live
[`TraitNode`](https://wago.tools/db2/TraitNode/csv?build=1.60.1.69893)
105696, in tree 1100 at `(10280, 3330)` — Retribution row 3.
This check uses the current Trait graph, not the leftover Classic Talent table.
The engine previously registered every rank unconditionally. That missing
availability check is corrected and covered by
`TestSealOfCommandRequiresTalent`. Search candidates that omitted the point
while casting Command are invalid, not DPS improvements.

Beta client `1.60.1.69893` against Classic Era `1.15.9.69722`, read with `tools/data_watch/spell_client.py`, the
talent curves with `tools/data_watch/trait_curve.mjs`, and `../beta/paladin.json`.

## How the numbers were read

Before changing a number the sim's value was checked against the Era client, so only what Forever moved was changed.
A Forever range is `EffectBasePointsF` with `Variance`; as in the mage pass, a rank's damage is that range plus
`EffectRealPointsPerLevel` for each level from the rank's base level to its max level (capped at 60), rounded down at
the low end and up at the high end. Tables that add the growth in Go (Exorcism, Seal of Righteousness) keep their base
and scale; tables that hold a single range (Hammer of Wrath, Holy Shock, Holy Strike) hold the grown value.

Splits the client and the sim make differently:
- **Consecration** now casts a damage spell each tick (1280345-1280349) with two parts: a flat amount every enemy in the
  area takes (no coefficient) and a larger amount at 0.095 a tick that only the first `$s3` = 4 enemies take. The sim
  keeps a dot on the parent spell, snapshots the sum for the first 4 targets and deals the flat part alone to the rest.
  "First 4 to enter" is read as the first 4 targets in the encounter.
- **Holy Strike** is a normalized weapon strike with a flat bonus (effect 121) and a weapon percentage (effect 31). The
  client multiplies the flat bonus by the percentage, as it does for Backstab (whose tooltip prints the multiplied
  225 for a 150 base), so rank 8 is 40% of (weapon + 81 to 105). The beta tooltip prints the raw `$s1`; the BlizzCon
  tooltip's 36 to 46 was already the multiplied figure. Spell power is added at 0.429 on top.
- **Seal of Righteousness' proc spells** (25742 ... 25713) carry 4 ... 35 in the beta client where Era carries 1. Those
  are the seal's value for a 2.2 speed one-hander (rank 8: 18.8 x 0.85 x 2.2 = 35), a tooltip figure; the sim's
  weapon-speed formula is unchanged.
- **Seal of Command** triggers 20424 from every rank in both clients. The sim's per-rank proc ids 20944-20947 were Era
  learn-spell dummies and do not exist in the beta client, so every rank now uses 20424.
- **Holy Shield's** proc ids 20955-20957 are Era learn-spell dummies too, absent from the beta client, where the block
  damage comes from the aura. They are kept only as metrics labels.

## Spells changed (old -> new, level 60 rank unless noted)

| Spell | Change |
|---|---|
| Holy Strike (8 ranks, new) | one guessed level 60 rank on Classic's 13953 (40% weapon + 36-46) -> ranks 679, 678, 1866, 680, 2495, 5569, 10332, 10333 from level 6; rank 8 20 mana, 12 sec, 40% x (weapon + 81-105), 0.429 confirmed. APLs now cast 10333. |
| Exorcism (6 ranks) | rank 6 505-563 -> 474-530, every rank about 10% lower; cost, cooldown, 0.429 unchanged |
| Hammer of Wrath (3 ranks) | rank 3 504-566 (Era reads 504-556) -> 473-523; ranks 1-2 316-348 / 412-455 -> 285-315 / 382-421 |
| Holy Shock (3 -> 4 ranks) | new rank 1 at level 30 (1311606, 160 mana, 128-140); old ranks become 2-4: 204-220 / 279-301 / 365-395 -> 175-189 / 248-268 / 334-362; cooldown 30 sec -> 10 |
| Consecration (5 ranks) | rank 5 48 a tick at 0.042 -> 12 a tick to all + 27 a tick at 0.095 to the first 4 (rank 1-4: 2+4, 3+7, 6+11, 8+20); cost and cooldown unchanged |
| Holy Shield (3 ranks) | 110 / 161 / 220 (ranks 2-3 guessed) -> 110 / 153 / 221; coefficient 0.05 -> 0.08; rank 1 trains at 40, not 30; 20% block and 4 charges confirmed |
| Seal of Righteousness ranks 1-3 | proc coefficient 0.029 / 0.063 / 0.093 -> 0.1 |
| Judgement of Righteousness ranks 1-3 | coefficient 0.144 / 0.312 / 0.462 -> 0.5; rank 1 flat 15 -> 14-16 |
| Judgement of the Crusader (6 ranks) | 20 ... 140 -> 23 ... 161 (Improved Seal of the Crusader's 15% folded in), 10 sec -> 40 sec |
| Seal of the Crusader | the seal's attack power no longer takes the 15% (the client keeps Classic's 306 at rank 6); also fixed the aura removing the Libram of Fervor's attack power with the wrong sign on expiry |
| Righteous Fury | Holy threat +60% -> +90% |
| Templar's Bulwark | now costs its 110 mana |
| Blessing of Sanctuary | gone from the beta client, so the sim no longer registers it under the Forever ruleset |

That is 43 spell ids across 11 spells whose numbers moved (Holy Strike 8, Exorcism 6, Judgement of the Crusader 6,
Consecration 5, Holy Shock 4, Holy Shield 3, Hammer of Wrath 3, Seal of Righteousness 3, Judgement of Righteousness 3,
Righteous Fury, Templar's Bulwark), plus Seal of Command's proc id and Blessing of Sanctuary's removal.

Unchanged in the client: Seal of Command and Judgement of Command (all ranks), Seal of the Crusader's attack power, Seal
and Judgement of Righteousness ranks 4-8, Holy Wrath, Judgement (10 sec, 6% of base mana), Divine Favor, Forbearance,
Lay on Hands (the 20 min cooldown is in the client at every rank).

`ui/core/spells/paladin.json`: every spell above is `forever` with a tooltip, the unchanged ranks are `classic`, and
Templar's Bulwark stays `assumed` only for the absorb being modelled as damage taken. The two remaining `unreviewed`
entries (14, 21) are not paladin spells. Ranked tables in `exorcism.go`, `hammer_of_wrath.go`, `holy_shock.go`,
`consecration.go`, `holy_strike.go`, `holy_shield.go` and `blessing_of_sanctuary.go` now keep their ids in an int table
so `sim/spell_sources_test.go` reads them, which also dropped 14 junk ids (mana costs and damage values the walk had
read as ids). `unreviewedSpellBudget` 762 -> 708, `unresolvedSpellSiteBudget` 19 -> 15.

## Talents changed

| Talent | Change |
|---|---|
| Vengeance | 1% a stack at every rank -> 1 / 2 / 3% a stack (5 stacks, 30 sec) |
| Vindication | 1% attack power at every rank -> 1 / 2 / 3%, 30 sec, 100% chance on a landed melee attack |
| Redoubt | 6% block and 10% chance at every rank -> 6 / 12 / 18 / 24 / 30% block and 2 / 4 / 6 / 8 / 10% chance (10 sec or 5 blocks). The tree read 10..50% chance with a flat 6%; it now reads the curve. |

## Checklist lines

### Paladin section
- **Resolved** `consecration.go:10`, Consecration baseline: every paladin trains all five ranks (SkillLineAbility,
  class mask 2). Its numbers moved, see above.
- **Resolved** `hammer_of_wrath.go:29`, Instrument of Law: client curve -500 / -1000 ms, so 0.5 / 1 sec as modelled.
- **Resolved** `holy_shield.go:18`, Holy Shield ranks: 110 / 161 / 220 -> 110 / 153 / 221, coefficient 0.05 -> 0.08.
- **Resolved** `holy_strike.go:15`, Holy Strike: eight ranks from level 6 on real ids; level 60 is 10333, 20 mana,
  12 sec, 40% x (weapon + 81-105). The 0.429 coefficient is in the client (`EffectBonusCoefficient`; the zero column
  the note looked at is `Coefficient`).
- **Open** `holy_strike.go:17`, whether melee-table Holy damage partial resists. That is a server combat rule the
  client data does not carry; it needs a combat log.
- **Resolved** `holy_strike.go:29`, Improved Holy Strike: -1000 / -2000 ms, as modelled.
- **Resolved** `holy_strike.go:41`, Iron Creed threat: 5 / 10 / 15 / 20 / 25%, as modelled.
- **Resolved** `holy_strike.go:94`, Iron Creed damage reduction: 2 / 4 / 6 / 8 / 10% for 6 sec (1311033), as modelled.
  The talent spell's own effects still carry 30 and 15; the curve is what the tooltips use.
- **Resolved** `sotc.go:34`, Improved Seal of the Crusader baseline: the judgement carries the 15% (161 at rank 6)
  and lasts 40 sec; the seal's attack power does not, so the sim stopped multiplying it.
- **Resolved** `swift_judgement.go:11`, Swift Judgement cooldown: 1 min in the client (1310994), as modelled.
- **Resolved** `talents.go:19`, Divine Precision: 6 / 12 / 18%, as modelled.
- **Resolved** `talents.go:33`, Holy Power: 1% a rank on every spell and 2% a rank more on Holy Shock, as modelled.
- **Resolved** `talents.go:37`, Sacred Duty stamina: 2 / 4%, as modelled.
- **Resolved** `talents.go:41`, Shield Specialization absorb: 10 / 20 / 30%, as modelled.
- **Resolved** `talents.go:47`, Champion of the Light: 33 / 66 / 100%, as modelled.
- **Resolved** `talents.go:89`, Redoubt: 6% a rank block, not flat; chance 2% a rank. One caveat: the chance comes from
  the curve on effect index 1, which the talent spell has no effect for. It is read as the proc chance because its rank
  5 value is the 10% `ProcChance` the spell carries. The spell's own numbers are not always the top rank's (Iron Creed's
  are stale), so an in-game proc count would settle it.
- **Resolved** `talents.go:170`, Shield Specialization mana: 33 / 66 / 100%, 6% of maximum mana (1310925), 3 sec
  internal cooldown (`ProcCategoryRecovery` 3000), as modelled.
- **Resolved** `talents.go:231`, Vengeance: 1 / 2 / 3% a stack, fixed.
- **Resolved** `talents.go:264`, Vindication: 1 / 2 / 3%, fixed. The target's lost attack power stays unmodelled.
- **Resolved (number), not modelled** `talents.go:301`, Consecrated Ground: 5 / 10%, and the cap is 4 enemies
  (`$20924s3`), not "4 or 8". The buff still sits on the paladin, so it only differs from the game on pulls of more
  than 4.
- **Resolved** `talents.go:337`, Instrument of Law threat: 10 / 20%, as modelled.
- **Resolved** `templars_bulwark.go:11`, Templar's Bulwark: 5 min cooldown, 110 mana (now charged), 8 sec, 100% of
  maximum health (1311015).
- **Resolved** `templars_bulwark.go:36`, Sacred Duty cooldown: -30 / -60 sec, as modelled.
- **Resolved** (paragraph) Infusion of Light rank 2: -500 / -1000 ms, so 0.5 / 1 sec.

### Talents the sim does not read
- **Voice of Truth** (1310897): 6 sec of silence and interrupt immunity on a 3 min cooldown. Nothing in a raid sim
  silences the paladin. Still not simulated.
- **Light's Vigil** (1310911, 1311590, 1311595): a 1.5 sec cast, 1340 mana at rank 3, that makes the next Holy Shock on
  the target skip its cooldown and deal 380-410 more Holy damage at 0.429 (or heal the party), refunding 75% of the
  cost. The numbers are all in the client, but it is a healer's tool: against an enemy it spends a global cooldown
  and 335 mana for less damage than one Exorcism, and no shipped spec casts Holy Shock at a boss. Still not simulated.
- **Infusion of Light**: shortens Holy Light, which the sim does not have. Still not simulated.
- **Improved Seal of Fury** (1314103): mana back when Seal of Fury's absorb breaks. Seal of Fury is not in the sim
  (below). Still not simulated.

### Baseline ability changes
Nothing was listed for the paladin. See the spellbook differences below.

## Spellbook: beta client against Classic Era (`--learned paladin`)

New:
- **Holy Strike**, 8 ranks (levels 6-60), implemented above.
- **Seal of Fury** 7 ranks (1311649, 1311656, 20163, 20419, 20421, 20422, 20423; levels 10-58) and its **Judgement of
  Fury** (1311650, 1311655, 20183, 20411-20414). A tank seal: rank 7 costs 200 mana, adds 35 Holy damage at 0.1 to
  each melee hit (20418), and while a shield is equipped turns 50% of that into an absorb; the judgement deals 146-160
  at 0.45. **Not implemented**: it needs a new `PaladinSeal` option in `proto/` for the primary seal, and its value is
  mostly the absorb, which the sim has no model for. Twist of Light's tooltip names it, so the echo would want it too.
  Some rank rows look unfinished (rank 4's seal effect carries 0.9 where the others carry 0.09, rank 7's carries none).
- **Hammer of the Righteous** (407632) at level 40, 6% of base mana, 6 sec: Holy damage equal to 3 times main hand
  weapon DPS to up to 3 targets. **Not implemented**: the spell and its text are Season of Discovery's ("Cooldown shared
  with Crusader Strike", which Forever does not have), with only the level moved, so it reads as a leftover rather
  than a designed Forever ability.
- **Holy Shock rank 1** (1311606, level 30), implemented.
- **Light's Vigil** 3 ranks, **Swift Judgement** (1310994), **Twist of Light** (1310735): talents the sim already
  lists.
- **Flash of Light** 5 new ranks (1313342-1313346) and **Holy Light** 5 new ranks (1313348-1313352); heals, not
  simulated.
- **Summon Warhorse** (1279399) at 40; no combat effect.

Removed:
- **Blessing of Sanctuary** and **Greater Blessing of Sanctuary**, gone from the client entirely. The sim skips it under
  Forever; `ui/protection_paladin` still offers it as the personal blessing and its preset still picks it.
- **Sanctity Aura** (20218) is no longer trained, though the spell still exists. The sim still applies it when the
  options pick it (`core.SanctityAuraAura`), which is not a paladin package change.
- Classic Era's Season of Discovery runes (Beacon of Light, Hand of Reckoning, Avenger's Shield, Crusader Strike,
  Divine Steed) and the S03 tuning passives.
- Several talents became rank-less spells (Holy Power, Anticipation, Vindication, Eye for an Eye, Improved Judgement,
  Pursuit of Justice, Improved Concentration Aura): the ranks moved into the trait curves.

## Changes needed outside the paladin package
- `sim/core/debuffs.go` `JudgementOfTheCrusaderAura`: duration 10 sec -> 40 sec (the value, 140 x 1.15 = 161, is right).
- `ui/protection_paladin` inputs and presets: drop Blessing of Sanctuary for Forever.
- `proto/` `PaladinSeal`: a Seal of Fury option, if Seal of Fury is wanted.

## DPS before -> after

Retribution (`TestRetribution`): Average 653.7 -> 709.8 (TPS 542.9 -> 587.7), mostly Vengeance at 3/3 now paying 3% a
stack. Long single target, full buffs (Human) 158.0 -> 162.9; long multi-target 228.7 -> 234.8.

Protection (`TestProtection`): Average 446.7 -> 476.9 (TPS 604.0 -> 728.6, Righteous Fury +90% and Holy Shield's
coefficient). Long single target, full buffs (Human) 105.8 -> 115.3; long multi-target (20 targets) 934.6 -> 414.9,
Consecration's capped part reaching only 4 of the 20.
