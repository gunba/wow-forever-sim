# Build updates

The comparison table below documents the fixed-gear research pass preceding
the [current gear comparison](gear_updates.md). Its paired DPS figures used
the earlier 26-build rules and gear; they are **historical evidence**, not
the current 27-build chart's DPS. The retained recipes and focused
corrections remain in the current ranking defaults.

## 23 September focused changes

Matched comparisons used the exact previously published race-specific gear,
enchants, external buffs, 300-second target and role Tier 1. Two independent
seeds with 5,000 iterations per arm and all available races appear with
complete native requests in
[`targeted_2026_09_23.json`](../artifacts/research_builds/targeted_2026_09_23.json).
The fixed paid-hit adjustment was recalculated for each candidate.

| Build | Change | Paired DPS difference across races |
|---|---|---:|
| Feral | Replace Demonic Rune with usable Thistle Tea; cast at ≤10 Energy | +6.39 to +7.19 |
| Shadow | 2/2 Early Demise, 3/5 rather than 5/5 Mental Strength | +1.58 to +1.79 |
| Mutilate | Venom with one Vile Poisons point moved and a ≥5-CP priority | −2.77 to −3.08; rejected |

The [Forever Thistle Tea item](https://www.wowhead.com/forever/item=7676/thistle-tea)
explicitly permits Rogues and Druids, restores 100 Energy and has a five-minute
cooldown. Feral used it once per five-minute fight; merely equipping it without an
explicit item action produced **no cast and no gain**. Shadow's Early Demise
increases Death crit chance only during execute; dropping two Intellect
talent points pays for it without removing Darkness. The Venom screen also
tried Poison, Vile Poisons and Seal Fate reallocations and 1–5-combo-point
thresholds; frequent one-point Venom upkeep lost substantially more damage.
These are comparisons on the present gear, not proof that no Venom build can
ever work.

Smite's existing Penance now hits at 0/1/2 seconds rather than 0.667/1.333/2,
following its client period and stated immediate bolt. Mage external
Innervate now waits for enough missing mana to receive a full 20-second
regeneration window instead of using a fixed fraction of maximum mana;
none of the ranked Mage profiles has an external Innervate. Corrected Imp
Firebolt ranges do not alter the DS/Ruin row because its Imp is sacrificed
before combat.

### Warrior rage model

The current Warrior comparison uses speed-normalized, damage-independent
rage on successful white swings: 4.5 rage per base weapon-second for
two-handers and `4.5 / 1.3` for one-handers. This matches low-level Forever
observations, including the reviewed level-18 combat log. It is an
**unverified level-60 extrapolation**, not a datamined max-level formula.
The off-hand gets half the one-hand rate before the sourced Dual Wield
Specialization multiplier; that half-rate is from
[Blizzard's Cataclysm explanation](https://www.bluetracker.gg/wow/topic/us-en/24038830344-what-about-rage-when-off-tanking/),
not from a Forever off-hand test. Avoided swings grant no autoattack rage
in the provisional model. Rage from damage taken retains the separate
inherited rule; the boss does not attack DPS, though self-damage can still
award a little rage. Classic
rules outside Forever are unchanged.

All three Warrior rows have been rerun. Comparisons that relied on the
old damage-proportional Warrior model are historical, not independent
validation of the revised numbers. See [T02](in_game_checks.md) for the
missing off-hand, extra-swing, avoidance and high-level measurements.

| Warrior build | Previous equal-race mean | Current equal-race mean | Change |
|---|---:|---:|---:|
| Fury | 937.56 | 786.64 | −16.10% |
| Arms | 785.91 | 733.07 | −6.72% |
| 2H Bloodthirst | 799.88 | 736.90 | −7.87% |

These are model-change effects, **not measured in-game DPS losses**.
Talents and APLs were held fixed, so they have not been reoptimized for
the new rage supply.
The separate 2H Bloodthirst row remains available for review, but its
earlier talent comparison used the superseded rage model; under the
current model it does not beat Arms for every race.

Nine established builds have updated talents or rotations. Three hybrids have
separate rows. The original Arcane, Fire, Frost, Arms and Survival identities
remain available.

Equipment, enchants, consumables and external buffs in the earlier
comparisons were unchanged from each reference profile, except that
Marksmanship selected no pet. Those results
use a Dragonkin reference boss, ordinary MP5, corrected casting/swing rules,
full role-specific Tier 1 and recalculated paid hit. Mage comparisons use
supported Mage Armor; the former Molten Armor setting applied no armor effect.

## Validation

Each comparison uses two independent seeds, 20295211 and 20295271, with 5,000
iterations per arm per seed. There are 338 unique complete runs, including
superseded recipes and the additional Stormcaller control. Every retained
candidate beats its original reference on every available race beyond the
conservative 95% Monte Carlo bound.

The bound is `1.96 × (baseline SE + candidate SE)`, pooling the independent
seeds first. It covers simulation noise, not uncertainty about game mechanics.
Means below weight races equally; the percentage ranges show individual races.
The published matrix uses a third seed, 20291951, so its numbers differ slightly.

### Updated builds

| Build | Prior mean DPS | New mean DPS | Gain across races |
|---|---:|---:|---:|
| Enhancement | 600.86 | 736.07 | +22.06–22.85% |
| Survival | 786.27 | 842.88 | +7.15–7.39% |
| Demonic Pact | 810.09 | 843.46 | +3.78–4.44% |
| Marksmanship | 847.34 | 917.15 | +7.98–8.54% |
| Retribution | 1011.08 | 1039.84 | +2.54–3.00% |
| Subtlety | 548.56 | 560.20 | +1.36–3.20% |
| Stormcaller | 526.89 | 536.91 | +1.73–1.97% |
| Smite | 552.28 | 559.37 | +1.11–1.39% |
| Shadow | 768.03 | 804.66 | +4.65–4.83% |

- **Enhancement:** 20/31 Nova/Maelstrom. Prioritize Clearcasting Fire Nova,
  maintain Searing, use Stormstrike and shocks, and cast Lightning Bolt only at
  five stacks. No Chain Lightning. Instant Lightning Bolt still incurs the
  modeled full-swing reset. The mana-limited time is nearly eliminated without
  changing the mana supplies.
- **Survival:** unchanged talents; rank-3 Wing Clip fills idle globals above
  40% mana. Its extra proc opportunities depend on the existing Windfury,
  Expose Prey and Judgement of Wisdom models.
- **Demonic Pact:** unchanged talents; Brand-first ordering and cast-time-aware
  DoT refreshes. Brand's pet/power attribution remains a model assumption.
- **Marksmanship:** 5/35/11 Lone Wolf, with no pet or Hawk. Full-rank Aimed Shot
  precedes Sniper Shot. The focused comparison also tested 3/3 Surefooted; the
  retained 1/3 allocation wins with this gear and paid-hit model. Hit is
  recalculated after talents, not before them. Matching tracking is active on
  the Dragonkin target; it supplies damage, not a separate crit multiplier.
- **Retribution:** 13/7/31 with lower-rank seal/Consecration choices. Its current
  gain is smaller than the earlier research result under the old casting model.
  Low-rank scaling and proc interactions still need confirmation.
- **Subtlety:** unchanged talents; three-combo-point Rupture maintenance,
  Ghostly Strike and improved Tea handling.
- **Stormcaller / Smite:** rank-2 Lightning Bolt / Smite fallbacks. These depend
  on the modeled low-rank coefficients; see T12 in the in-game checks.
- **Shadow (earlier comparison):** 20/0/31 with Death before Mind Blast. The
  current default is **18/0/33** with 2/2 Early Demise as shown above. Four
  learned Death ranks are implemented. The extra script value 150 remains unresolved;
  no additional execute multiplier is inferred. Conditional two-tick Mind Flay
  clipping is available as an alternate APL, but did not reliably beat the new
  default. It only interrupts for a ready higher-priority cast, never simply
  because two ticks elapsed. Smite retains its existing APL: adding Death to its
  ordinary priorities lost damage in the resource-controlled comparison.

### Separate hybrid rows

| Build | Reference profile | Reference mean DPS | Hybrid mean DPS | Gain across races |
|---|---|---:|---:|---:|
| Pet/Melee Hunter — 16/10/25 | Original Survival | 786.27 | 866.76 | +10.11–10.34% |
| Arcane–Frost Mage — 28/0/23 | Frost | 681.65 | 711.89 | +4.17–5.03% |
| 2H Bloodthirst Warrior — 20/31/0 | Arms | 786.10 | 797.03 | +1.05–1.77% |

Each hybrid retains its reference profile's equipment and options. The Hunter
benefits more from matching tracking than the revised Survival allocation.
The Mage uses full-rank spells, including Ice Lance; its result depends on the
assumed 0.143 Ice Lance coefficient and unresolved proc timing. It is not a
Fire build. The Warrior is a two-handed Fury-style build without Mortal Strike;
the historical gains above used the superseded damage-based rage model.
The current row uses provisional speed-normalized rage, and its level-60
performance still needs testing.

These are distinct architectures, not evidence that every hybrid beats every
other build. Their model caveats also appear in the web profile picker.

## Not adopted

- The 29/22 Lightning allocation was checked on actual Stormcaller equipment.
  It improves the old profile but has no clear overall advantage over the
  retained rotation. It does not replace pure Elemental.
- Combat-dagger respecs do not replace Mutilate or Subtlety, and are not
  portable to the current non-dagger Combat equipment.
- Frostfire downranking remains a conditional research result, not a default.
  The separate Arcane–Frost row does not rely on downranked Frostfire.
- Nova/shock Enhancement remains behind the retained Nova/Maelstrom build.
- Shockadin and Penance/Shadow did not beat the stronger same-profile
  alternatives. Historical cast-heavy Enhancement results remain superseded.
- Other original rows retain their existing builds; no unvalidated replacement
  was inferred.

## Evidence and replay

- [Validation summary](../artifacts/research_builds/summary.json)
- [Complete validation archive](https://gunba.github.io/wow-forever-sim/classic/review/research_builds/validation.json.gz)
- [Current build details](build_reviews.md) and [in-game checks](in_game_checks.md),
  especially T12, T14, T21, T29, T30, T39, T48 and T49.
- [Pre-update benchmark at a21ec1117](https://github.com/gunba/wow-forever-sim/blob/a21ec111784b515c4f35a9244ae83356285c32d2/artifacts/forever_dps_5min.json)

The archive contains complete requests/results, unnormalized players, seeds,
recipe identifiers and original comparison profiles. Both sides receive the
same armor, off-hand and encounter corrections. The additional focused MM and
Priest comparisons use separate seeds and are retained in
`artifacts/profile_corrections/`. No old no-armor Mage or invalid off-hand
comparison is presented as current evidence.

See the [spell coverage audit](spell_coverage.md), [Windfury comparison](windfury.md)
and [in-game checks](in_game_checks.md) for implementation gaps and unresolved
server behavior. The preceding release and research snapshots remain in Git history.

`tools/forever_bench/apply_research_profiles.py` prepares the complete 171-player
input set from this archive. `run_matrix.py` regenerates the 684 baseline,
Tier-off, +10% and +50% scenarios. See the
[benchmark instructions](../tools/forever_bench/README.md).
