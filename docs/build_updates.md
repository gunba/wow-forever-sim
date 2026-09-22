# Build updates

Nine established builds have updated talents or rotations. Three hybrids have
separate rows. The original Arcane, Fire, Frost, Arms and Survival identities
remain available.

Equipment, enchants, consumables and external buffs are unchanged from each
comparison profile. These results use ordinary MP5, the corrected casting/swing
rules, full role-specific Tier 1 and recalculated paid hit.

## Validation

Each comparison uses two independent seeds, 20295211 and 20295271, with 5,000
iterations per arm per seed. There are 310 unique complete runs and 168 matched
input checks, including the additional Stormcaller control. Every retained
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
| Survival | 772.75 | 828.25 | +7.13–7.37% |
| Demonic Pact | 810.09 | 843.46 | +3.78–4.44% |
| Marksmanship | 847.34 | 874.41 | +3.08–3.29% |
| Retribution | 1011.08 | 1039.84 | +2.54–3.00% |
| Subtlety | 548.56 | 560.20 | +1.36–3.20% |
| Stormcaller | 526.89 | 536.91 | +1.73–1.97% |
| Smite | 552.28 | 559.37 | +1.11–1.39% |
| Shadow | 768.03 | 770.44 | +0.26–0.46% |

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
- **Marksmanship:** 17/34 Hawk/Sniper. The Hawk guardian and pet model remain
  approximate.
- **Retribution:** 13/7/31 with lower-rank seal/Consecration choices. Its current
  gain is smaller than the earlier research result under the old casting model.
  Low-rank scaling and proc interactions still need confirmation.
- **Subtlety:** unchanged talents; three-combo-point Rupture maintenance,
  Ghostly Strike and improved Tea handling.
- **Stormcaller / Smite:** rank-2 Lightning Bolt / Smite fallbacks. These depend
  on the modeled low-rank coefficients; see T12 in the in-game checks.
- **Shadow:** a modest 20/0/31 resource allocation improvement, not a new hybrid.

### Separate hybrid rows

| Build | Reference profile | Reference mean DPS | Hybrid mean DPS | Gain across races |
|---|---|---:|---:|---:|
| Pet/Melee Hunter — 16/10/25 | Original Survival | 772.75 | 830.80 | +7.39–7.62% |
| Arcane–Frost Mage — 28/0/23 | Frost | 632.73 | 710.49 | +8.77–13.39% |
| 2H Bloodthirst Warrior — 20/31/0 | Arms | 774.90 | 811.64 | +4.46–5.06% |

Each hybrid retains its reference profile's equipment and options. The Hunter
is only a few DPS ahead of the revised Survival rotation, not another 7% ahead.
The Mage uses full-rank spells, including Ice Lance; its result depends on the
assumed 0.143 Ice Lance coefficient and unresolved proc timing. It is not a
Fire build. The Warrior is a two-handed Fury-style build without Mortal Strike;
its inherited damage-based rage model remains unverified at level 60.

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
recipe identifiers and the original comparison profiles. The simulator's
mechanics are unchanged from a21ec1117; the native harness adds the new build
definitions and their parent-profile rules.

`tools/forever_bench/apply_research_profiles.py` prepares the complete 171-player
input set from this archive. `run_matrix.py` regenerates the 684 baseline,
Tier-off, +10% and +50% scenarios. See the
[benchmark instructions](../tools/forever_bench/README.md).
