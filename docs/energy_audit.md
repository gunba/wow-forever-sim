# Energy mechanics

## Evidence and model

| Mechanic | Evidence | Implementation |
|---|---|---|
| Base recovery | `PowerType.db2` gives 10 Energy/sec both in and out of combat, zero regeneration-interrupt time, and a base maximum of 100. Same records in Forever 1.60.1.69893 and 1.60.1.69913. | 10/sec before explicit regeneration effects. |
| Recovery timing | [Rogue beta feedback](https://us.forums.blizzard.com/en/wow/t/rogue-changes-wishlist/2352824) and [Feral beta feedback](https://us.forums.blizzard.com/en/wow/t/feral-druid-in-wow-forever-has-a-design-problem-not-just-a-tuning-problem/2354035) describe smooth, usable regeneration rather than Classic ticks. These are player reports, not Blizzard specifications. | Forever uses 100-ms integration steps, with 1 Energy per step before modifiers. This is not a claimed server tick interval. |
| Haste scaling | Model assumption, not established by client data. The energy records, including flags, are identical to Classic Era; that does not settle how the server applies haste. | General haste multiplies regeneration. Rating, Skyborne haste, Berserking and applicable Tier 1 haste are included. Attack-speed-only effects such as Slice and Dice are not. In-game check T15 remains open. |
| Adrenaline Rush | Spell 13750 specifies +100% regeneration for 15 seconds. | Doubles recovery. Partial intervals are settled at the old rate when the buff starts or ends. |
| Miss refunds | Client costs do not specify the general refund rule. | Existing 80% builder refunds and zero damaging-finisher refunds retained pending T16. Refunds use paid cost, not undiscounted cost. |

The previous engine delivered 20.2 Energy every 2.02 seconds in Forever, just
as in Classic. The average rate was already 10/sec, but ability readiness and
energy-cap waste were wrong for smooth recovery. Fixing timing does not amount
to giving the classes additional baseline resources.

Source records and spell inputs are preserved in
[`assets/db_inputs/forever_energy.json`](../assets/db_inputs/forever_energy.json).
Regenerate them with:

```sh
python3 tools/database/import_forever_energy.py --cache /tmp/forever-client-data
```

The primary table is
[Forever PowerType](https://wago.tools/db2/PowerType/csv?build=1.60.1.69913).

## Spending and resource talents

The audited client costs match the engine before talents and set bonuses:

| Ability | Base Energy |
|---|---:|
| Shred / Backstab / Mutilate | 60 |
| Claw / Sinister Strike | 45 |
| Rake | 40 |
| Ferocious Bite / Hemorrhage / Eviscerate | 35 |
| Rip | 30 |
| Slice and Dice / Rupture / Venom | 25 |
| Tiger's Fury / Adrenaline Rush | 0 |

Existing talent reductions remain: for example, Shredding Attacks brings Shred
to 42, Ferocity brings Claw to 40 and Rake to 35, and Improved Sinister Strike
brings Sinister Strike to 40. Flawless Execution reduces Eviscerate by 10; the
Forever Tier 1 five-piece reduces it by another 5 when enabled.

Relentless Strikes already restores 25 Energy with a 20%-per-combo-point chance.
Vigor increases the maximum pool by 5/10; its verified effects do not justify
adding retail Vigor's regeneration bonus. Gnome Expansive Mind and Eureka are
already modeled separately. Tiger's Fury restores Energy through King of the
Jungle; its client records do not establish an additional baseline 60-Energy
grant without that talent.

Seal Fate's 500-ms internal cooldown is present in the Forever client's
`SpellAuraOptions` record for spell 14186 and already matches the engine.
The two simultaneous Mutilate strikes therefore do not receive two separate
Seal Fate procs in this model; that cooldown is not merely inherited guesswork.

Ferocious Bite still consumes remaining Energy on a landed hit. Its extra-energy
behavior and general refund rules are not being replaced with retail rules.
All four benchmark profiles use explicit APLs, not the legacy scripted Cat
rotation's Classic pooling estimates.

## Verification

`sim/core/energy_test.go` covers cadence, average recovery, caps and wasted
recovery, regeneration-rate transitions, and refunds under a discounted cost.
The existing racial tests cover the Gnome pool and class-specific discounts.
Haste gain/expiry settles the elapsed partial interval at its old rate; a change
on an already-paid tick must not award that interval twice. Adrenaline Rush
multiplies the haste-adjusted rate.

### Historical timing-only comparison

All 31 supported race/build combinations were rerun at 5,000 iterations, using
the exact frozen baseline profiles. No equipment, talents, rotation, buffs or
hit-budget inputs changed. The new runs produced no APL warnings.

| Build | Representative race | Previous DPS | Smooth-recovery DPS | Change across races |
|---|---|---:|---:|---:|
| Feral | Tauren | 523.25 | 519.01 | −0.81% to −0.70% |
| Combat | Orc | 496.58 | 490.47 | −1.40% to −1.18% |
| Mutilate | Orc | 482.13 | 482.28 | −1.70% to +0.03% |
| Subtlety | Orc | 418.95 | 418.64 | −0.22% to −0.04% |

Very small changes are within sampling error. The largest Mutilate decrease is
Gnome; it should be revisited during race-specific rotation testing.

These are **timing-correction comparisons, not newly optimized rankings**.
An APL's energy thresholds can behave differently with small resource increments.
For example, the unchanged Feral baseline spent more Energy on Ferocious Bite
and less on Rip after the correction. Smooth recovery is not automatically a
damage increase with an unchanged rotation.

Exact requests, metrics and comparisons:
[`energy_cadence_5min.json`](../artifacts/energy_cadence_5min.json) and
[`energy_cadence_5min.csv`](../artifacts/energy_cadence_5min.csv).
The reference remains [`forever_baseline_5min.json`](../artifacts/forever_baseline_5min.json).

The retained Feral build was then revalidated separately: 534.44–535.24 DPS,
3.03–3.14% above the corresponding smooth-recovery baselines, with no mana-limited
time or APL warnings. A fresh rotation screen found no further validated gain.
[`optimization/feral.json`](../artifacts/optimization/feral.json) labels historical
Classic-tick searches separately from the new comparison and rotation screen.
These figures precede haste-scaled regeneration. The current engine and its
retained profiles require new comparisons; historical figures are not the final
ranking.
