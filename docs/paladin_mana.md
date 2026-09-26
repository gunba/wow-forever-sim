# Protection Paladin mana correction

## What was wrong

The published rotation had two seal actions: a guarded refresh near expiry,
and an unconditional seal filler. The filler repeatedly recast an active
30-second Seal of Righteousness. Forever Judgement does not consume that seal.

In matched 2,000-iteration tests, this produced about **101 seal casts and
18,100 mana spent** over five minutes. Maintaining the seal instead required
**11 casts and 1,980 mana**. Removing only the filler gained about 12 DPS;
removing the seal completely lost more than 100 DPS. Removing Righteousness
from the published Ret and physical-Ret profiles also lost substantial damage.
Those two rotations and their inputs are unchanged.

## Rotation

- Cast the primary seal three seconds before the pull, then Holy Shield at
  −1.5 seconds. Both casts pay their normal costs.
- Keep health-gated Bulwark, Holy Shield and seal maintenance.
- Use Judgement, then full-rank Consecration at **30% mana or above**.
- Use rank-one Consecration as the **15% mana** fallback.
- Retain Hammer of Wrath, Holy Strike and the existing cooldown handling.

The seal is refreshed only near expiry, not used to fill empty globals.
Gear, talents, enchants, consumables, external support, natural hit and Tier 1
are unchanged. This is a focused correction and resource comparison, not a
claim of global rotation optimality.

## Independent checks

Two independent seeds, **20263237 and 20263283**, each use **100,000 iterations
per arm**, for all three races and both tank workloads. The final recipe passes
the damage, survival, total-threat and lowest-target-threat safeguards against
both the previous release and the original frozen tank controls.

| Race | Previous DPS | Corrected DPS | Gain | Conservative 95% bound on gain |
|---|---:|---:|---:|---:|
| Undead | 423.85 | 467.58 | +10.32% | +43.73 ±0.16 DPS |
| Human | 407.19 | 450.81 | +10.71% | +43.61 ±0.15 DPS |
| Dwarf | 400.07 | 443.24 | +10.79% | +43.17 ±0.15 DPS |

These are five-minute confirmation means. The matrix uses its usual separate
5,000-iteration seed, so its displayed numbers differ slightly.

The three-attacker workload retains the 90-second encounter and 2,500 HPS
support. The final lowest-target TPS is approximately **212.31 / 211.20 /
209.72** for Undead / Human / Dwarf.

An intermediate no-filler recipe improved damage and passed the survival
checks, but lost roughly 3% secondary-target threat. It was not selected.
The inherited engine credits non-exempt effective mana restoration with
0.5 threat **per enemy**: less mana waste can therefore reduce that credited
threat. Source-specific exemptions, ownership and allocation need verification
under **CORE-014**. No threat formula was changed to make this correction pass.

## Evidence

- [Final comparisons and safeguards](../artifacts/paladin_mana/validation.json)
- [Unchanged-profile accounting](../artifacts/paladin_mana/preservation.json)
- [Complete comparison archives](../artifacts/paladin_mana/archives.json)
- [Superseded Protection results](../artifacts/history/481ac9412f/paladin_before_mana_fix.json.gz)

The archives include complete players, requests and results. Each gzip contains
a `files` object keyed by source filename. `final-a` and `final-b` are the
accepted independent comparisons; the earlier confirmation archives retain
the rejected intermediate recipes. Extract a pair's `Single.Request` or
`Multi.Request` and replay it with:

```sh
go run -tags with_db ./tools/forever_tanks -request request.json -output replay.json
```

The comparison engine matches the combat implementation at
`481ac9412f3dcce9ac24ac33f3e4802dafdd9f1c`; the archives record its binary hash.
Only the three Protection profiles and their twelve matrix scenarios are
replaced. Other profiles and results are retained unchanged.
