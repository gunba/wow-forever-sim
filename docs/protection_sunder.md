# Protection Warrior Sunder maintenance

The old rotation could spend its remaining rage on other actions while Sunder
was about to expire. Having a refresh action in the priority list was not
enough: five stacks sometimes fell off and had to be rebuilt.

The revised rotation:

- Begins refreshing five stacks with eight seconds remaining, ahead of ordinary
  Thunder Clap/Demoralizing Shout refreshes.
- Keeps emergency defenses, Shield Block and Bloodrage ahead of that refresh.
- Reserves the current rage costs of Sunder and Shield Block before discretionary
  damage or Heroic Strike/Cleave queueing while building stacks or nearing expiry.
- Skips refreshing when the existing Sunder duration already covers the rest
  of the encounter.

The reserve uses actual spell costs, including talent discounts, rather than
assuming that every profile pays the same amount. Existing shout, Thunder Clap,
Demoralizing Shout and five-stack setup duties remain. Sunder is supplied by the
Warrior on the main target, not by a free external debuff.
Custom scenarios with an existing equal-or-stronger exclusive armor debuff do
not reserve rage for a Sunder application that the engine would reject.

## Stack coverage

Each arm covers 40,000 fights per workload: all ten races, two non-overlapping
seed ranges, 2,000 iterations per race/seed. These figures measure **five
stacks**, not merely the presence of the Sunder aura.

| Workload | Measure | Previous | Revised |
|---|---|---:|---:|
| One attacker, 300 seconds | Fights with a post-ramp stack loss | 42.22% | 0.57% |
| One attacker | Mean time to five stacks | 13.92s | 12.54s |
| One attacker | Mean time below five after the initial ramp | 4.86s | 0.053s |
| Three attackers, 90 seconds | Fights with a post-ramp stack loss | 8.80% | 0.038% |
| Three attackers | Mean time to five stacks | 13.69s | 12.27s |
| Three attackers | Mean time below five after the initial ramp | 0.81s | 0.003s |

This is not guaranteed uptime: avoidance and resource pressure can still cause
a rare lapse. Expiration exactly at encounter end is not counted as a failure.

## Performance and scope

Matched comparisons use 10,000 iterations per arm at seeds 20275001 and
20300001 for every race. Orc uses 25,000 per arm/seed after its initial DTPS
confidence interval straddled the safeguard. Even those longer seed ranges do
not overlap. All races pass the existing conservative survival and threat
safeguards in both workloads.

Single-target changes are small: approximately +0.35 to +1.35 DPS. Human's
personal DPS gain is not statistically established. The correction's purpose
is reliable armor reduction, not a claim of a new optimal damage rotation.

Only Protection Warrior profiles and their 40 baseline/Tier-off/+10%/+50%
scenario rows change. The other 191 profiles and 764 scenario rows are
preserved. Equipment, talents, external support, rage mechanics and workload
settings are unchanged.

## Unbridled Wrath

Before this rotation correction, Human Protection averaged 42.1 landed ordinary
white swings and 22.6 landed Windfury extra swings per five minutes.
At one talent point, `64.6 × 12% × 1 rage ≈ 7.75 rage`, matching the recorded
7.753 with no cap waste.

The model excludes Heroic Strike/Cleave replacements, consistent with the
[available observations](https://github.com/ClassicWoWCommunity/forever-bugs/issues/105).
Windfury extra white attacks are included; their Unbridled Wrath eligibility
has not been independently established from Forever logs. The issue's “Fixed”
label accompanies a comment about a future beta build and does not by itself
establish the live deployment date. No rage/proc formula changes are made here.

## Evidence and replay

- [Survival/threat comparisons](../artifacts/protection_sunder/validation.json)
- [Full-stack coverage](../artifacts/protection_sunder/coverage.json)
- [Complete request/result archives](../artifacts/protection_sunder/archives.json)
- [Unchanged-row accounting](../artifacts/protection_sunder/preservation.json)
- [Previous profiles/results](../artifacts/history/91b669cc2aa/protection_before_sunder_maintenance.json.gz)

Extract a request from an archived pair or coverage record:

```sh
go build -tags with_db -o /tmp/tank-replay ./tools/forever_tanks
/tmp/tank-replay -request request.json -output result.json
go build -tags with_db -o /tmp/sunder-audit ./tools/sunder_audit
/tmp/sunder-audit -request request.json -output coverage.json
```

The observer preserves the aura's armor effect and does not add actions or draw
random numbers. A targeted regression reproduces the previous Night Elf lapse
and checks that the revised APL retains its stacks in that same scenario.

### Historical seed qualification

The simulator advances the seed by one per iteration. Older reports that used
nearby starting seeds, including the preceding tank-support and queue reviews,
therefore reused most of the same iteration seeds. Their runs must not be
counted as independent replications. Raw results and individual-run means
remain available; their pooled uncertainty needs separate review (SCEN-016).
The comparisons in this report use disjoint seed ranges.
