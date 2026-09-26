# Flurry evidence review

Checked on 26 September 2026 against the working tree based on
`51cc3d6e2c22b341b15d23a7e37be78099e90b21`.

## Conclusion

[Issue 12](https://github.com/magey/forever-warrior/issues/12) is an open
research checklist, last updated September 24, with no comments or Flurry
combat logs. The linked [compendium](https://ppach-warriorcompendium.share.connect.posit.cloud/flurry.html)
explicitly calls its Flurry results a model: the beta level cap prevents
testing the talent. It therefore does not establish a new Flurry bug.
No Flurry behavior was changed on this evidence alone.

## What this simulator actually does

| Area | Current implementation | Evidence status |
|---|---|---|
| Amount | 5/10/15/20/25% melee speed by talent rank | Uses the talent values; the separate triggered aura's 30% remains a source inconsistency |
| Proc | Direct melee crits, including Bloodthirst, Whirlwind and queued attacks, activate or refresh it | Forever event eligibility still needs logs |
| Refresh | Consume the old charge first, then refill to three on a crit | A triggering white crit leaves three, not two |
| Shared charges | Main- and off-hand use the same aura | Modeled; not independently confirmed in Forever |
| Consumption | Each white-tagged hit event consumes a charge, including avoided swings; no 500 ms gate in Forever | Existing same-batch regression covers the implementation, not the game |
| Queued specials | Heroic Strike and Cleave carry both auto and special masks, so they consume charges; each Cleave target can invoke consumption | Important unverified case; do not infer its answer from ordinary yellow attacks |
| Ordinary specials | Bloodthirst/Whirlwind do not consume charges, but their crits refresh them | Matches the linked model's broad assumption |
| Swing timing | Activation and expiry rescale the remaining melee swing timer | Not just a change to future complete swings |
| Stacking | Multiplicative with other melee-speed effects | 25% faster means an interval divided by 1.25, not an interval 25% shorter |
| Rage | Uses the unhasted weapon's speed, not the shortened interval | Supported by low-level tests of other haste effects; still extrapolated for Flurry/level 60 |

Code: `sim/warrior/talents.go` (`applyFlurry`,
`makeFlurryConsumptionTrigger`), `sim/warrior/heroic_strike_cleave.go`,
`sim/core/unit.go` (`MultiplyMeleeSpeed`), `sim/core/attack.go`
(`UpdateSwingTimers`) and `sim/core/rage.go`.

The compendium's transition example refills then consumes the triggering
swing, leaving two charges. Our callback order leaves three. It is not a
server measurement, so matching that arithmetic is not a reason to change
the engine. Its aggregate uptime/rage estimates also use different hit,
crit and rotation assumptions; they are not benchmark targets.

## Supporting haste evidence

The raw attachment in
[discussion 14](https://github.com/magey/forever-warrior/discussions/14)
was inspected independently. Four usable outgoing hits under Berserking
each add 6.9 Rage; comparable unbuffed deltas are 6.9–7.0. The hasted hit
intervals are approximately 1.78–1.84 seconds for the reported two-second
weapon. Incoming-hit resource snapshots were included, and the long
out-of-combat gap was excluded.

This supports the engine's use of base weapon speed for Rage. It does not
measure Flurry, its charges, or level-60 behavior. The downloaded capture
and its identifying metadata remain local.

## Separate defect found in the linked research

Unbridled Wrath used the same broad white-hit mask and therefore also
triggered on queued Heroic Strike/Cleave. This is **not** a Flurry change.
[The public bug report](https://github.com/ClassicWoWCommunity/forever-bugs/issues/105)
reports no Unbridled Wrath procs from about 250 such hits; its auto-only
eligibility is also recorded in
[issue 4](https://github.com/magey/forever-warrior/issues/4).

Forever now excludes special-tagged replacement swings from Unbridled
Wrath only. The regression reproduced both illegal triggers before the
fix and still permits ordinary main/off-hand hits.

The model retains the intended **12% per talent point**. The reported
lower beta proc rate was acknowledged as a game bug with a future fix;
that is not evidence that 8% is the intended level-60 talent. The exact
deployment of that server fix remains unverified.
