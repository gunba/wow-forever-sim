# DPS check review

> Historical evidence and dispositions. Use the [canonical register](uncertainties.md)
> for current status: several later fixes supersede the models described below.
> Earlier T50/T51/T52 references are qualified there rather than silently renumbered.

## Poison, periodic damage and mana review, 23 September

The pinned Forever client **1.60.1.69913** exposes the level-60
[Instant Poison VI](https://www.wowhead.com/forever/spell=11340/instant-poison-vi)
trigger **11337** with mean 88 and variance giving 76–100 damage; the
[Deadly Poison V](https://www.wowhead.com/forever/spell=25347/deadly-poison-v)
trigger **25349** has 23 damage per three-second tick per stack, or 92 over
12 seconds. Both triggered effects have zero recorded bonus-SP and AP
coefficients in `SpellEffect`. These values already matched `sim/rogue/poisons.go`;
no additional poison scaling was inferred from a different expansion.
The fifth Deadly Poison recipe is present in the client, but its **learning
source in Forever** remains unresolved: the Classic handbook came from the
unavailable Ruins of Ahn'Qiraj. Client presence is not proof a player can
learn it. Poison charges are not consumed in this five-minute benchmark;
that shortcut still needs checking for unusually high proc-rate builds.

Forever's ordinary periodic damage now reads current offensive power and
modifiers each tick; Classic still snapshots. Applied combo-point count stays
fixed while Rip and Rupture re-read current AP. Deep Wounds reads current
main-hand weapon/AP each tick in Forever, matching the
[independent WoWSims implementation](https://github.com/wowsims/forever/commit/36cfc58328e3)
under the reported no-snapshot rule. Its main-hand choice and exact
server-side behavior remain testable assumptions. Ignite retains the pool
created by its triggering hit; this is not a fresh direct-damage calculation
on every tick. Periodic crit was already dynamic.

The local 23 September Mage capture contains two Berserking windows:
the first three Arcane Missile events in each still occur approximately one
second apart. The unbuffed channel likewise has approximately one-second
spacing. This supports the existing distinction between **casting speed**
and channel haste; event gaps across separate channels or target changes
cannot establish a channel coefficient. In one gap between channels, mana
rose by about **70 over 12.79 seconds** after accounting for the next
85-mana channel cost. The log does not identify a flat-MP5 effect or a
Spirit-only control, so it cannot establish MP5-per-second, passive tick
timing or a level-60 mana rate. Private combat-log identifiers remain local.

[Judgement of Wisdom rank 3](https://www.wowhead.com/forever/spell=20355/judgement-of-wisdom)
still specifies **59 mana** per proc in the client. Its **50%** eligible-event
chance is inherited from Classic and also used by the independent WoWSims
Forever implementation; the client aura's proc trigger does **not** specify
the server script's effective chance. At that assumption, the expected
return is **29.5 mana per eligible attack or spell event**, not per second.
The previously published Arcane Orc result averages roughly **22.9 mana/sec**
from Wisdom, and physical Ret about **37 mana/sec**, with the debuff forced
on the single boss throughout the encounter. No duplicated mana source was
found in `sim/core/debuffs.go`: periodic effects and proc-generated damage
are excluded, and a missed spell is excluded. The return can still be large
because frequent separate landed actions can each roll. Actual Forever proc
chance/eligibility still needs a direct combat-log check; the amount itself
is sourced. These figures describe the prior release and must not be reused
as the new benchmark results.

Reviewed 22 September 2026 against release
[`8ce872cfddae0560675b918a0850cc5b5434c497`](https://github.com/gunba/wow-forever-sim/commit/8ce872cfddae0560675b918a0850cc5b5434c497).
The [active checklist](in_game_checks.md) separates immediate priorities from
later-access and conditional work. This register covers **T01–T51 without
discarding unresolved questions**.

“Resolved” means that the stated question has a sufficient source or an
established benchmark rule—not that every possible server override has been
measured. “Open” identifies a particular missing interaction or formula.
“Conditional” preserves a credible alternative-build question without making
it a prerequisite for all current rankings.

The original 22 September review changed no engine or result data; subsequent
corrections are noted below. Private chat,
identifying metadata and original combat logs remain outside the repository.
The [previous full checklist](https://github.com/gunba/wow-forever-sim/blob/8ce872cfddae0560675b918a0850cc5b5434c497/docs/in_game_checks.md)
preserves the original test descriptions.

## Source findings

### Druid rank curves, not rankless tooltips

The active Druid trait tree is **1089** in captured client **1.60.1.69893**.
Following `TraitNode → TraitNodeEntry → TraitDefinition →
TraitDefinitionEffectPoints → CurvePoint` resolves the conflicting amounts:

| Talent | Node / entry / definition | Curve | Values by rank |
|---|---|---|---|
| Moonglow | 104925 / 129585 / 134386 | 85328 | −8, −17, −25% cost |
| Nature's Majesty | 104927 / 129587 / 134388 | 85326 | 2, 4% crit |
| Nature's Reach | 104929 / 129589 / 134390 | 85329 and 102877 | 2, 4% spell and physical hit |
| Improved Moonfire | 104931 / 129591 / 134392 | 85327, 111698, 111699 | 5, 10% modifiers |
| Moonfury | 104936 / 129596 / 134397 | 85335 | 2, 4, 6, 8, 10% damage |

Primary tables:
[definitions](https://wago.tools/db2/TraitDefinition/csv?build=1.60.1.69893),
[effect curves](https://wago.tools/db2/TraitDefinitionEffectPoints/csv?build=1.60.1.69893),
[curve points](https://wago.tools/db2/CurvePoint/csv?build=1.60.1.69893).
`tools/data_watch/trait_curve.mjs` reads this chain. The actual tree connection
matters because legacy definitions coexist with active definitions.

These match the present talent amounts in `sim/druid/talents.go`. Nature's
Majesty's standalone spell scalar of 1 is **not** proof that the two-rank
talent should grant 1/2%. Likewise, a guide's 5% Moonfury claim does not override
the active curve. The remaining T26 question is damage calculation order:
Moonfury's school-percent effect and Improved Moonfire's damage modifiers
versus the engine's base-only multiplication.

### Consecration

The present Ret profiles use rank 1, spell **26573**, whose description refers
to child **1280345**. The child's two damage effects are:

- effect 0: flat 2, zero SP coefficient;
- effect 1: flat 4, coefficient .095, for the first four targets.

The wrapper supplies eight one-second ticks. Thus the single target receives
`2 + 4 + .095 × SP` per tick before modifiers, matching
`sim/paladin/consecration.go`. Rank 5's child **1280349** similarly supplies
12 + 27 and .095. Source:
[SpellEffect](https://wago.tools/db2/SpellEffect/csv?build=1.60.1.69893) and
[Consecration](https://www.wowhead.com/forever/spell=26573/consecration).
Ground placement and the identity of the fifth entrant cannot change this
stationary single-target calculation. T30 no longer requests those tests.

### Penance: corrected timing

[Penance 1316995](https://www.wowhead.com/forever/spell=1316995/penance)
explicitly describes an immediate bolt and subsequent bolts every second for
two seconds. The captured description references **402261**, whose periodic
effect is 1000 ms and duration is 2000 ms.

The earlier engine delayed all three ticks at 2/3-second intervals. It now
fires immediately and at one and two seconds; the focused regression asserts
all three instants. The full channel still has three bolts. Whether haste
changes this timing remains T40; the tooltip alone does not establish it.

### Pet abilities and resource evidence, 23 September

The [69977 SpellEffect](https://wago.tools/db2/SpellEffect/csv?build=1.60.1.69977),
[SpellPower](https://wago.tools/db2/SpellPower/csv?build=1.60.1.69977)
and [SpellCooldowns](https://wago.tools/db2/SpellCooldowns/csv?build=1.60.1.69977)
exports match 69913 byte-for-byte. At level 60, the selected Hunter Cat's
Claw **3009** remains 43–59 for 25 Focus; Bite **17261** is 81–99 for 35
Focus and has a ten-second cooldown. Both already match the engine.
The selected Warlock Succubus's Lash **11780** remains 50 base damage for
160 mana and a twelve-second cooldown, also already modeled. Imp Firebolt
**11763** has mean base damage **44**, variance **.113636**, 115 mana and
a .571 coefficient. The inherited rank-7 engine range of 43–48 was high;
all Firebolt rank ranges now use the captured mean and variance. The exact
server-side rounding of fractional damage bounds remains unknown. The ranked
DS/Ruin APL **sacrifices its Imp at −5 seconds**, so this ability correction
does not change that row; a focused test confirms that an unsacrificed Imp
does cast and deal damage. The Hunter
pet's 1.6-second AI action interval and the Imp's extra 200 ms cast delay are
timing approximations, not new claims about Forever spell cooldowns. Owner
hit/AP inheritance is still unresolved (T29).

In [69977 SpellEffect](https://wago.tools/db2/SpellEffect/csv?build=1.60.1.69977),
Maelstrom **408498** has base effect 20, dummy effects 50 and 5; the rank
curve changes effect 20 to 16/12/8/4 at ranks 1–4. Child **408505**
discounts cast time and mana by 20% per stack and masks Lightning Bolt,
not Chain Lightning. [AuraOptions](https://wago.tools/db2/SpellAuraOptions/csv?build=1.60.1.69977)
records an eligible proc trigger and a 100% trigger field, but no PPM link.
Nothing in this snapshot establishes how the server scripts the actual
proc frequency or what dummy 50 controls. The 2 PPM per talent point in
`sim/shaman/talents.go` remains an explicit model assumption, not a verified
rate. A level-30 controlled proc count is still useful.

The supplied local Rogue combat capture is client **1.60.1.69893**. It
contains resource snapshots on melee and casts, not a passive Energy
`SPELL_ENERGIZE` stream. In the short session labeled 1% haste, two
snapshots 4.728 seconds apart differ by three Energy after one known
45-Energy Sinister Strike; inferred gain is **48 Energy**, or roughly
**10.15/sec**. One-point snapshot quantization alone spans approximately
**9.94–10.36/sec**, covering both 10 and 10.1. Several shorter subintervals
have inconsistent observed increments, so this is not a reliable haste
coefficient. The [69977 PowerType table](https://wago.tools/db2/PowerType/csv?build=1.60.1.69977)
still lists base 10/sec, with no general-haste rule. No character identifiers
or raw log lines are published here.

### Blizzard's beta notice

The [official September 18 known-issues post](https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-known-issues-september-18/2352687)
changes the evidence status of several checks:

- **T01/T06:** higher-level melee glancing damage and displayed tables are
  flagged as incorrect; the post supplies no replacement formula.
- **T29:** pet hit inherits from the owner. The coefficient is not stated;
  the shared owner-stat implementation below provisionally assumes 100%.
- **T30:** periodic crits activating Vengeance are a bug. The current engine
  already listens to direct-hit crits, not periodic damage, in
  `Paladin.applyVengeance`.
- **T32:** Starshards should benefit from Twilight Focus. That talent protects
  against incoming-damage pushback, which this encounter does not apply.

These are official statements of intent/known defects, not measurements of a
corrected level-60 build. Do not calibrate the simulator to a known broken beta
display.

### Pet formula source audit

The [Pet Attack Formulas](https://github.com/classic-hunter/forever-hunter/wiki/Pet-Attack-Formulas)
and [Pet Stat Mechanics](https://github.com/classic-hunter/forever-hunter/wiki/Pet-Stat-Mechanics)
pages explicitly describe **original Classic**, not Forever. Both were
created September 16, before beta. Their "no owner inheritance" statement
contradicts both Blizzard's T29 notice and that wiki's later
[Forever Beta Changes](https://github.com/classic-hunter/forever-hunter/wiki/Forever-Beta-Changes)
page. The Classic page's 5.2-second Focus ticks also contradict its later
Forever report of continuous regeneration. Do not import the Classic
formula's Orc Command or AQ-set multiplier: the former was replaced in
Forever and the latter is unavailable in our benchmark.

The current beta-client export has hidden **Hunter Pet Scaling 415429** and
**Warlock Pet Scaling 416189**, with separate owner-linked health, power,
physical/spell hit and crit, speed and resistance aura slots. Their zero
effect base points and Wowhead's displayed "1" placeholders do **not**
provide the owner-to-pet coefficients. The Warlock aura also carries
Intellect, Mana regeneration and damage/healing-related slots; it is not safe to reuse
a Hunter coefficient **as an established client fact**. In the absence of
contrary measured values, the current model deliberately uses the Hunter
coefficients as shared beta assumptions for both classes. The Warlock's
owner spell power is provisionally mapped into pet spell power and pet
attack power at the same 10% offensive-power rate; its additional
Intellect and Mana-regeneration aura coefficients remain unresolved.
The client `PowerType` row for Focus in build 69977 supplies a concrete
**10 Focus/second in combat**, and Hunter scaling 415429 adds 50 maximum
Focus to the base 100. `sim/core/focus.go` now accrues this rate continuously
for Forever and preserves Classic's old 26.25/5.25-second ticks.
Bestial Discipline scales the corrected rate.

The later wiki reports Hunter inheritance of 10% of the higher attack
power source, owner Crit, 30% Armor and two pet Health per owner Stamina.
These are beta-community observations, not coefficients in the exported
scaling aura; an independent [beta guide](https://www.warcrafttavern.com/forever/guides/hunter/)
repeats the AP/Crit and continuous-Focus claims. The wiki's owner-Hit
statement is **still Classic-specific**; official beta notes establish
owner-linked hit for both classes but no coefficient. The new shared
`core.ForeverPetInheritance` uses the beta wiki's reported rates for
both classes, including a provisionally full owner-hit transfer. The wiki
also reports that pet damage is not normalized by swing speed. A separate
beta Discord report says pets do the same damage per hit at different
attack speeds. The Forever Hunter pet weapon now preserves both the base
roll and the AP-added part of a hit at the selected pet's attack speed;
its actual swing interval still varies. The fixed two-second reference AP
conversion for unequal owner AP totals remains a modeling assumption,
not an isolated beta AP measurement.

The later wiki and a [direct beta report](https://eu.forums.blizzard.com/en/wow/t/pets-cant-be-buffed/629896)
say that external blessings, scrolls and similar buffs cannot be placed on
Hunter pets. Forever Hunter and Warlock pets now share the no-direct-buff
rule in `sim/core/buffs.go:applyPetBuffEffects`, and the pet AP/Agility/
Strength consumables no longer apply or appear in the Forever Hunter
defaults. An owner buff can still reach the pet once via stat inheritance.
The source-supported Classic mechanics that remain useful are the
14 AP-per-DPS white-attack term, pet-family physical multiplier, happiness
multiplier, pet Agility conversion and family-specific attack speeds. Our
Hunter code already models those. The beta client gives Lightning Breath
rank 6 (25012) **zero** bonus-power coefficient, now honored by the
Forever ability.
The companion correction has not yet been included in a published
benchmark; historical rankings still use the prior pet model.

### Rage evidence

The previously reviewed level-18 capture supports damage-independent rage:
29 conservatively isolated landed swings, or 54 with the documented relaxed
timestamp exclusion, gain 9.3/9.4 rage across normal hits, glances and crits.
The latter sample spans 15–78 damage. With the recorder's stated 2.7-second
weapon, its mean is about 3.4602 rage per weapon-second. Level and speed are
reported context, not parsed equipment metadata.

The seven public files at
[`tzcnt/forever-data` commit c7d1746](https://github.com/tzcnt/forever-data/tree/c7d17462c50d1eb0103aa5e2aff52f77f33e3418/raw-logs)
lack matching equipment/level metadata. Their 478 eligible snapshot candidates
are not 478 verified weapon-rate observations. This does not establish the
level-60, off-hand, avoidance or extra-attack formula. Repeating an already
known low-level damage-versus-rage comparison is less useful than filling those
specific gaps.

A later research screenshot reports 5.17 rage at 1.5 seconds, 5.54 at 1.6,
7.27 at 2.1 for one-handed weapons, and 13.50 at 3.0 and 14.40 at 3.2 for
two-handers. These approximate `speed × 3.46` for the first three and
`speed × 4.5` for the last two, consistent with the independently reviewed
2.7-second level-18 sample above. Unlike the screenshot alone, the
[researcher's later issue](https://github.com/magey/forever-warrior/issues/3)
links twelve raw level-8/9 logs with weapon speeds, armor and health context.
Those logs support the two-hand rate and a geared incoming-hit rule; they
still do not establish a level-60 or off-hand formula. A comparison
to Cataclysm at its maximum level is a proposed analogy, **not evidence
that Forever retains this rule at level 60**. The inherited Classic formula
would instead change rage by a factor of
roughly five between 15- and 78-damage swings; the captured constant
9.3/9.4 already falsifies that low-level prediction. The level-60 rule
remains unresolved. Avoided outgoing swings are also disputed: the screenshot
reports no rage for miss/dodge/parry, whereas the inherited Classic rule
generates rage for dodge/parry. Record outcomes separately to validate the
new provisional Forever branch.

The Forever Warrior branch now uses `base weapon speed × 4.5` for a landed
two-handed autoattack, or `base weapon speed × (4.5 / 1.3)` for a landed
one-handed autoattack, regardless of damage, glancing or crit. Misses,
dodges and parries award none. The off-hand provisionally gets **half**
the one-hand rate before the sourced Dual Wield Specialization rage
multiplier. The half-rate rule is from [Blizzard's Cataclysm design
statement](https://www.bluetracker.gg/wow/topic/us-en/24038830344-what-about-rage-when-off-tanking/),
not a verified Forever off-hand measurement. Rage from damage taken uses
`unmitigated damage × 10 / maximum health` for
geared Warriors, as fitted to the linked level-8/9 logs. Non-Forever rules
remain unchanged. These outgoing rates are an explicit benchmark assumption extrapolated to level
60, not proof of a level-60 or extra-swing server formula.

The low-level logs also show roughly **double** the geared incoming Rage
rate below 130–170 armor. This may be a beta defect; we do not extrapolate
that threshold to level 60. The reference boss does not attack DPS players,
though self-damage can still generate some rage. A tank/solo comparison at
higher levels, with Berserker Rage both on and off, is still needed.

## Relevance to the published profiles

Measured from the prior published
`artifacts/history/a8c419e91/forever_dps_5min.json`, SHA-256
`e1c3c9e3f259d90d2d21b506848b0baab2fe9c88dbba427245c2db959090812b`.
For each action, sum enemy-target damage (`unitIndex = 0`), divide by
`Iterations × 300`, then divide by that profile's DPS for its share. Ranges
span the supported races. Do not include Death's self-target backlash.
These historical contributions are **not predictions of the size of a
correction**; use the current results for revised per-build totals.

| Component | DPS range | Share of its build |
|---|---:|---:|
| Ret rank-1 Consecration | 82.50–86.82 | 7.79–8.14% |
| Smite rank-2 fallback | 90.38–105.10 | 16.05–18.89% |
| Stormcaller rank-2 Lightning Bolt | 88.59–92.35 | 16.55–17.24% |
| Smite Penance | 124.91–130.31 | 22.18–23.29% |
| Frost Ice Lance | 186.86–188.77 | 27.27–27.70% |
| Arcane–Frost Ice Lance | 129.25–130.49 | 18.14–18.29% |
| Balance Insect Swarm | 85.16–86.36 | 14.12–14.30% |
| BM Summon Hawk | 66.65–67.66 | 7.61–7.66% |
| Pet/Melee Summon Hawk | 71.28–72.61 | 8.23–8.30% |
| Destruction Incinerate | 256.00–262.10 | 32.44–33.13% |

No current profile casts Wrack, Frostfire Bolt or Starshards. Their questions
remain relevant to those alternative builds, not to a claimed contribution
in the present rows. None equips the three formerly queried armor IDs
272505/272556/272749. All 24 Mage requests specify `MageArmor`.

## Complete disposition register

| ID | Disposition | Evidence and remaining scope |
|---|---|---|
| T01 | Open; merged with T06 | Official glancing/display defects invalidate the old screenshot-only plan. `sim/core/attack.go` and outcome tables still use inherited rules. Need corrected-build outcome evidence, not a guessed 25% boss penalty. |
| T02 | Provisional model; later test needed | Low-level normalized rage is supported by logs. The new speed-based Warrior branch extrapolates to level 60, assumes Cataclysm's half-rate off-hand and treats avoided swings as zero. Verify those rules and extra attacks; see the evidence section above and `sim/core/rage.go`. |
| T03 | Partly resolved; open interaction | Class-specific costs and three charges are implemented in `sim/core/racials.go` from `assets/db_inputs/forever_races.json`. Keep last-charge channel/DoT lifetime scope; remove redundant tooltip, healing and utility checks. |
| T04 | Conditional resource feature | [1259705](https://www.wowhead.com/forever/spell=1259705/read-ley-line) resolves cast/CD and 15-second versus 15-minute duration. Buffs 1270842/1259691 have +100% mana-regeneration aura 110, not a grant of full casting Spirit regeneration. Effect on Spirit/flat regeneration under the five-second rule remains unimplemented and unmeasured. High Order only; no favorable location assumed. |
| T05 | Source watch; merged into T50 | Existing supported racial/class spell implementations remain. Retained class-skill rows alone do not prove new teaching paths. Dark Sacrifice/utility acquisition is not an automatic DPS test; use `docs/spell_coverage.md` for particular omissions. |
| T06 | Open; merged with T01 | Do not ask for the same physical-table evidence twice. Separate single-weapon/DW miss, glancing frequency and damage after the reported beta defect is addressed. |
| T07 | Conditional alternative build | Wrack is implemented in `sim/warlock/wrack.go` but has zero current casts. Keep cancellation/DoT amplification scope for a Wrack build; its six-second base channel and .143 tick coefficient are already captured. |
| T08 | Open DPS arithmetic | [1293813](https://www.wowhead.com/forever/spell=1293813/incinerate) gives .714 and the 25% Immolate bonus; a dummy effect does not expose its arithmetic. `sim/warlock/incinerate.go` applies the bonus before adding SP. Whole-hit versus base-only remains material. |
| T09 | Resolved | [Demonic Pact 425464](https://www.wowhead.com/forever/spell=425464/demonic-pact) explicitly permits a different demon and excludes resummoning the sacrificed demon. `sim/warlock/warlock.go`, pet enabling/reset and existing full-run checks implement this. The no-Pact guard allegation was retracted; see M13 in `mythicsim_review.md`. |
| T10 | Conditional alternative build | No current ranked APL casts FFB. Keep Missile Barrage/FoF/T1 script interactions for a future FFB recipe. The present Arcane–Frost row uses Frostbolt/Missiles/Lance, so T49, not FFB, is its immediate uncertainty. |
| T11 | Lower-priority DPS edge | `forever_tier1.md` sources the duration extension. Seven full two-second ticks plus an aura tail are modeled; a fractional final tick is not established. Only that tail is disputed, not all Swarm damage. |
| T12 | Open; high impact | Client-listed low-rank coefficients are retained. A further server-side downranking penalty would affect substantial fallback damage. Check SP slopes, not just tooltip damage or OOM. |
| T13 | Resolved benchmark rule | Aimed 2s, Sniper 4s and Multi-Shot .5s casts with continued ranged autos are the adopted rule, covered by existing Hunter regressions. No new recording prerequisite. See `auto_attack_audit.md`. |
| T14 | Open guardian model | `sim/hunter/summon_hawk.go` approximates guardian attacks; permanent-pet speed confirmation does not establish Hawk damage, inheritance or overlapping guardians. It contributes to two current rows. |
| T15 | Partly resolved; open haste | `PowerType` supplies 10 Energy/sec; usable smooth recovery has player evidence. General-haste scaling is explicitly assumed; attack-speed-only enchant scopes are sourced. Furor's low-rank shifting edge is conditional on a shifting APL, not a new baseline-regeneration question. See `energy_audit.md`. |
| T16 | Open; merged with T35 | Costs and paid-cost refund accounting are implemented. Server refund fractions on avoided builders/finishers are not provided by those cost records. One combined check replaces two requests. |
| T17 | Open; merged into haste check | `sim/priest/mind_flay.go` uses three unhasted one-second ticks. Client base period does not establish haste scaling; T40 collects the relevant timings. |
| T18 | Open | `sim/druid/forms.go` retains the Classic form weapon. No reviewed primary source supplies a Forever weapon-DPS-to-Cat conversion. This is separate from Energy or gear stats. |
| T19 | Timing resolved; engine corrected | Current Penance text and referenced period establish 0/1/2-second bolts. `sim/priest/penance.go` now implements them, with a focused regression. Channel haste remains under T40. |
| T20 | Removed from damage testing | The queried armor IDs are absent from current equipment. Vendor stock and faction acquisition remain catalog provenance work, not tests of damage. Existing verified eligibility and vendor-stat precedence remain unchanged. |
| T21 | Open | Maelstrom's proc frequency and Totem of the Storm/free-cast interactions remain material to Enhancement. LB-only eligibility is already sourced; do not retest CL as a proc spender. |
| T22 | Resolved data correction | `assets/db_inputs/forever_effect_audit.json` and `tools/database/enchant_overrides.go` retain explicit hybrid enchant effects. The generic healing-to-damage fallback is gone. There is no affirmative evidence for an extra hidden conversion requiring a standing gameplay task. |
| T23 | Resolved charges; accepted timing model | Forever aura **400625**, not obsolete 48108, supplies the stack/charge model (`forever_effect_audit.json`; the engine uses action alias 44445). `Mage.applyHotStreak` consumes it on completed Pyroblast. Interrupted-cast cases are outside the current stationary APLs; completion/in-flight ordering remains a disclosed edge, not proof of persistent acceleration. |
| T24 | Conditional consumable interaction | Item class eligibility, 100 Energy and five-minute cooldown are sourced. Current five-minute profiles do not require repeated post-pull Tea uses. The inherited shared-conjured category is still not proved by that tooltip; revisit before relying on cross-category or prepull sequencing. |
| T25 | Partly resolved; narrowed | [16864](https://www.wowhead.com/forever/spell=16864/omen-of-clarity) confirms the ten-second cooldown and Wrath/free-action exclusions. `SpellAuraOptions` rows 129438/129841 give 100%/one charge; 16870 lasts 15s. Keep effective proc probability, not redundant charge/exclusion tests. |
| T26 | Rank disputes resolved; damage scope open | Active curves above settle Majesty/Reach/Moonglow/Moonfury amounts. Regular Faerie Fire form permissions follow client shapeshift masks. Keep base-only versus full-SP damage scope; remove Thick Hide defense and redundant rank/cost measurements. |
| T27 | Resolved; threat removed | Client Fire-totem requirement is enforced by `sim/shaman/fire_totems.go`. Threat ownership cannot change this no-incoming-damage DPS encounter. |
| T28 | Open; combined with T30 | Weapon/totem exclusion and advertised chance are sourced. Actual eligible triggers, extra-attack resolution and the inherited ICD still need event evidence. `windfury.md` separates the already measured simulator effects. |
| T29 | Existence resolved; coefficients open | The official notice establishes pet hit inheritance. Both class inheritance functions return empty stats. Exact hit/AP/SP/level formulas remain required; this is not permission to import retail pet scaling. |
| T30 | Split | Consecration single-target constants resolved; geometry/multi-target tests removed. Holy Strike's weapon-plus-damage structure is sourced in 10333; later rank changes require a newer source capture, not a generic spell retest. Keep Seal/Wisdom/Windfury proc eligibility with T28. Vengeance already excludes periodic events. |
| T31 | Open | Retain precise melee crit/proc scope for Predator's Edge and Expose Prey. Do not extend generic gear-crit behavior to every talent or treat a miss as a landed trigger. |
| T32 | Conditional; pushback removed | Inner Focus periodic snapshotting matters to such alternative casts, not every Priest rotation. No current Starshards casts. Its base cost/periods are sourced; Twilight Focus's incoming-damage interaction is outside this encounter. |
| T33 | Mostly resolved; utility archived | Shadowform 15473's modifiers, mana/GCD and implemented Holy-damage permissions follow the reviewed records and code. Healing/shield restrictions are not raid-DPS tests. Any new Divine Spirit teaching evidence belongs to the consolidated availability inventory. |
| T34 | Resolved charge model; edges conditional | Mage Clearcasting 12536 has one charge; the prior duration-wide free-casting error is fixed. Utility-spell consumption is not part of current DPS validation. Revisit cast/channel edge ordering only with specific contrary evidence; no blanket retest requested. |
| T35 | Merged into T16 | Later upstream 80%-refund changes are implementation choices, not proof of Forever finishers. Preserve the unresolved refund question once. |
| T36 | Unsupported availability; excluded | Sanctity Aura lacks a current trait/training path in the reviewed client. No 10% Holy buff is granted. Reopen on an actual acquisition source; a retained tooltip is not grounds for a mandatory test. |
| T37 | Narrow conditional script question | Four learnable Death ranks, .429 SP, 15s CD and stated 10%-max-health backlash are implemented. Retain unexplained script value 150 and any execute-damage implications. Remove absorb/healer-survival cases from this encounter's required list. See `spell_coverage.md`. |
| T38 | Open | Keep Mutilate hand ordering, off-hand flat bonus and Cold Blood consumption; fixed dagger gear and simultaneous hits make inference from aggregate crit damage unreliable. |
| T39 | Open; threat removed | Child damage/coefficient records are captured. Brand power attribution and pet-family trigger/crit rules are not established by those constants. Threefold threat is irrelevant to the DPS comparison. |
| T40 | Open | General haste currently changes cast time but not ordinary spell GCD. Dedicated GCD modifiers do not prove a general haste rule. Includes channel timing questions from T17/T19. |
| T41 | Partly resolved; lower priority | Nature's Grace 16886 establishes separate 10% cast/GCD effects for three seconds. Only completion-versus-impact proc timing remains; do not re-ask the sourced magnitude/duration. |
| T42 | Partly resolved; conditional | Actual stone stat values, legal occupied-offhand use and imbue replacement were corrected. No profile selects a stone. Fire+Shadow mask versus Shadow-only Spellstone wording remains a specific source conflict for a stone comparison. See H23 in `history_review.md`. |
| T43 | Unsupported availability; excluded | Fel Armor 403619's retained acquisition-method-3 row does not establish current training. No bonus is granted. Keep an acquisition watch rather than asking players to measure a spell they may not be able to learn. |
| T44 | Lower-priority assumption check | Ordinary MP5 remains the default; fivefold MP5 is opt-in. Current Mage resource waits are much smaller after Mage Armor/build corrections. The old OOM comparison cannot establish the server rate. See `mana_regeneration.md`. |
| T45 | Outside present encounter | Seal of Fury shield-break mana is tank/incoming-damage behavior, absent from current two-handed Ret. Preserve the feature gap in `mythicsim_review.md`; do not invent mana for a player receiving no attacks. |
| T46 | Source watch; merged into T50 | Hammer of the Righteous 407632 is not established as currently learnable. Teaching/CD/scaling work is conditional on availability evidence, not a current ranked-rotation requirement. |
| T47 | Resolved code cleanup | The unsupported Lightning Shield callback is removed. Its old individual Earth Shock reset never bypassed the shared shock CD. No demonstrated recycling DPS gain remains to test. |
| T48 | Accepted benchmark rule | Original Classic spell/swing handling was explicitly adopted. Five-stack instant LB's full reset remains part of that rule; its independent Forever validation is not claimed. Preserve the disclosure, but do not list the adopted rule as an outstanding release prerequisite. |
| T49 | Open; high impact | Ice Lance's .143 coefficient is assumed, not supplied by the captured damage effect. Whole-hit frozen multiplier and in-flight FoF timing materially affect the current Frost/Arcane–Frost rows. |
| T50 | Consolidated source watch | Availability of Black Arrow, active Lacerate, Felfire, Victory Rush, HoR and unusual pet teaching needs affirmative current acquisition evidence. `spell_coverage.md` distinguishes these from supported-but-unimplemented damage spells. Not a request to grind characters for every retained record. |
| T51 | Implementation backlog | Prowl/Pounce/Ravage are absent from the engine. Ordinary stealth existence is not a missing experiment. Stun-immune Pounce application is a conditional opener question when implementing that feature, not a current Feral damage source. |

## Scope of the result

The active list now has **10 immediate topics and eight later-access topics**,
with conditional follow-ups separated. Several topics combine duplicate old
IDs; this count is not a claim that only 18 server details could ever differ.

The strongest newly resolved questions are the active Druid rank values,
Consecration's single-target formula, Pact coexistence and Penance's corrected
schedule. The most important outstanding implementation gap is pet hit
inheritance. Remaining resource, coefficient and proc-script questions are not
declared solved merely because another simulator makes the same assumption.
