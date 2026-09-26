# Protection Warrior queue correction

## Fault and repair

The Protection APL referenced the untagged Heroic Strike (25286) and Cleave
(20569) damage actions. The engine registered separate queue actions with
`tag: 1`, but did not reject direct APL casts of the damage components.
The result was instant extra attacks while ordinary swings continued to deal
damage and generate rage.

Both actions now use their queue entries in the prototype and all ten exact
race profiles. The engine marks queued damage components as on-next-swing
attacks and rejects direct APL casts with a queue-specific warning. This also
protects Maul imports. Ability-value queries can still inspect the damage
component; actual queued attacks still fire through the swing scheduler.

The roster audit found this fault in all ten Protection profiles, not Fury,
other Warrior DPS profiles or Bear. Maul was already correctly queued.

## Scope and evidence

Gear, talents, enchants, consumes, external support, rage formulas and encounter
settings are unchanged. All ten Protection profiles are replayed for the
5,000-iteration baseline and Tier-off/+10%/+50% scenarios using seed 20291951.
The other 191 profiles and their 764 scenario rows remain unchanged.

Separate fixed-input comparisons use seeds 20263441 and 20263493,
10,000 iterations per arm, in both the 300-second single-attacker and 90-second
three-attacker workloads. Historical direct-cast inputs run only on the
previous engine, strictly to measure the error; they are not legitimate
performance targets. The raw records include complete requests, player inputs,
resource/action metrics and engine hashes.

For Orc, the independent-seed mean changes from **691.78 to 386.62 DPS**;
Heroic Strike falls to about **53 casts per five minutes**. Across all ten races,
corrected means are **385.78–399.40 DPS**. These confirmation means are separate
from the common-seed matrix replay.

- [Comparison summary](../artifacts/protection_queue/validation.json)
- [Raw-record archive index](../artifacts/protection_queue/archives.json)
- [Unchanged-profile accounting](../artifacts/protection_queue/preservation.json)
- [Pre-correction results](../artifacts/history/a7fe4cd30b/protection_before_queue_fix.json.gz)

Earlier Protection gear/talent/APL selection gains and survival/threat
non-inferiority claims are superseded. Current gear and talents are retained,
not asserted to be optimal under the corrected queue model. Paladin and Bear
selection evidence remains separate.

Future Warrior gear searches use the separately frozen
`artifacts/tanks/queue_corrected_warrior_controls.json` inputs, never the
invalid historical APL. Search startup rejects any control with APL warnings.

## Rage interpretation

Queued Heroic Strike/Cleave spend rage and replace an ordinary main-hand swing;
the replacement does not award ordinary white-swing rage. Refunds and distinct
eligible procs are separate resource events. Frequent queueing by itself is not
proof of excessive rage.

Warrior's level-60 normalized outgoing and incoming rage models remain qualified
by WAR-001 through WAR-004. Bear instead retains inherited damage-based rage,
with its separate uncertainty tracked as DRU-012. No Warrior formula has been
silently applied to Bear.

## Regression coverage

`tools/forever_bench/queued_attacks_test.go` checks every current prototype and
ranked profile for proper next-swing action tags. Runtime tests reject direct
Heroic Strike, Cleave and Maul actions while confirming their queue actions
still produce attacks. Browser replay checks load the actual ranked Protection
profiles and compare native/WASM damage, threat and damage-taken metrics.
