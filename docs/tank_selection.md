# Tank selection

## Objective and controls

The primary objective is single-target DPS with tank performance preserved in
both workloads:

| Workload | Duration | Attackers | Base swing damage | External healing |
|---|---:|---:|---:|---:|
| Reference | 300 seconds | 1 | 3,000 | 1,500 HPS, three-second cadence |
| Multiple attackers | 90 seconds | 3 | 1,500 each | 2,500 HPS, three-second cadence |

Both use level-63 Dragonkin, 3,731 starting armor, two-second enemy swings,
parry haste and a six-second burst window. These are stress scenarios, not
measured boss/healer behavior.

Paladin/Bear controls are their classes' previously published selected tank
setups, with only the race changed, from the
`artifacts/tanks/*_1t_20261993_selected.json` requests.
After the queue correction, future Warrior searches instead use the ten
frozen `artifacts/tanks/queue_corrected_warrior_controls.json` profiles.
These controls do not move with candidate trials. Their actual buffs,
debuff ownership, professions and consumables are retained. Support differs
between tank classes and from the non-attacking DPS scenario: the matrix must
not imply an equal-encounter survival ranking.

## Screening safeguards

Every proposed selection is compared with the **frozen control**, not only
the previous candidate. The permitted losses cannot accumulate over passes.
The same limits apply independently to both workloads:

| Measurement | Maximum permitted change |
|---|---:|
| Damage taken per second | +1% |
| Tankiness Metric Index | +2% |
| Modeled death probability | +0.25 percentage points |
| Total threat per second | −1% |
| Lowest per-target threat per second | −2% |

These are practical selection tolerances, not verified game constants.
Screens use sample means and are not final acceptance evidence. Small
fluctuations can eliminate a viable candidate, so the search does not establish
a global optimum. Final independent-seed runs must report the uncertainty,
including the binomial uncertainty of death probability; ambiguous safety
comparisons require more iterations or retaining the control.
Per-target threat has no exported variance estimate. Its safeguard must pass
in each independent seed; it is not a formal non-inferiority confidence bound.

One-target and three-target threat are measured against actual enemy unit
indices. Self/helpful targets must not become an extra enemy in the minimum.
Threat coverage is still a proxy: these encounters do not simulate an entire
party competing for aggro, changing targets or interrupting scripted casts.
The engine records a modeled death but continues the iteration. DPS is not
reduced by post-death downtime; survival therefore remains a separate mandatory
selection check, not something already priced into the DPS column.
The matrix's Tier and gear-scaling cells still measure **DPS**, not total tank
value. Defensive bonuses can change incoming damage and resource generation.
The paired raw results retain their survival and threat measurements as well.

## Equipment and build selection

- Search the existing legal, normalized v2 pool slot by slot, including coupled
  weapon layouts and supported enchants. No new defensive stat prices are
  inferred.
- Shield tanks retain a one-hand weapon and shield. Bear retains its tank
  specialization and class-legal equipment. Deep tank-tree identity is
  validated separately from general talent legality.
- Natural hit remains natural hit. The selector neither purchases hit nor
  converts excess hit into damage stats; it preserves existing rotational hit
  coverage.
- Gear comparisons keep the proposed talents/APL fixed. Talent/APL experiments
  must preserve defensive and debuff maintenance rather than exploiting an
  omission in the survival model.
- An independently tested slot gain must exceed a conservative DPS noise bound.
  Accepted slot swaps are still proposals until the final frozen-control and
  independent-seed checks pass.

The native reports retain both scenarios, guard limits, rejection reasons and
candidate/selected equipment. A final control failure explicitly reverts the
selection; it is not reported as an accepted improvement.

The current pool has no separate Defense/Dodge/Block-rating armor alternatives.
The search trades its existing stats and supported enchants, not an unrestricted
future tank-loot catalog. This is a coverage limit, not a claim that offensive
armor is universally the best tank gear.

## Selected profiles

The Bear selection checks use two independent seeds per arm with 20,000
iterations per seed; Paladin uses 100,000 to resolve small differences in modeled
death probability. Both pass the stated safeguards in both workloads.

**Protection Warrior's previous selection evidence is invalid.** Its APL
directly cast Heroic Strike/Cleave damage rather than queuing replacements.
The [queue correction](protection_queue.md) retains the existing gear/talents
and replays both workloads with 10,000 iterations on each of two independent
seeds. It is a correctness repair, not a new optimization or a claim of
non-inferiority against an invalid control. Earlier Warrior search results and
their 42.2–45.8% gain claims are superseded.

| Tank | Races | Single-target DPS gain over the original control |
|---|---:|---:|
| Protection Warrior | 10 | Not established after queue correction |
| Protection Paladin | 3 | 11.9–12.5% |
| Bear | 4 | 17.7–18.2% |

These gains combine equipment, talents and APL changes; they are not isolated
talent effects or evidence of live-game balance. Each race has its exact saved
selection. The matrix uses a separate common-seed 5,000-iteration replay, so its
rounded DPS can differ slightly from the confirmation estimates.

- **Warrior:** retained 4/6/41, with race-specific Protection point choices. Improved
  Heroic Strike, Cruelty and Unbridled Wrath replace some utility/threat points.
  Heroic Strike/Cleave use actual queue actions with the existing thresholds;
  the earlier threshold comparison was invalid. Shield Block availability,
  debuffs and health-gated defensive cooldowns remain unchanged. Thunder Clap is also
  prioritized for three or more attackers. Fixed-target pack threat remains
  uneven: preserving the least-target control does not establish reliable
  aggro against a damage-dealing party. Charge/interrupt utility is not exercised
  by these stationary, noncasting targets.
- **Paladin:** 0/43/8. Two points move from Precision/Guardian's Favor to
  Conviction; equipment remains unchanged. The
  [mana correction](paladin_mana.md) removes redundant seal refreshes,
  pre-buffs Righteousness/Holy Shield and uses full-rank Consecration above
  a mana reserve with a cheaper fallback. It passes both the original and
  previous-release controls on two new independent seeds. Earlier Paladin
  rotation comparisons used the wasteful seal filler and are superseded;
  they do not establish that full-rank Consecration is unsuitable.
  Seal of Fury remains a major implementation gap.
- **Bear:** 1/40/10, adding damage/resource talents and removing restoration
  investments not used by this rotation. Frenzied Regeneration is explicitly
  health-gated, Barkskin is proactive, and Maul starts at 25 rage. Captured
  Gift of Nature masks do not select Frenzied Regeneration; Genesis includes
  Lacerate's periodic parent. The retained [scope records](../artifacts/tanks/current/talent_scope_evidence.json)
  distinguish the client snapshots used for those checks.

**Bear's enchant result is conditional:** Crusader's current form eligibility
and inherited proc-rate model are not verified by Forever logs (DRU-011).
The earlier static-Agility configuration already gained roughly 15% in the
independent checks; the additional enchant/gear gain is not proof that Crusader
works in game. Existing Night Elf Feral DPS uses the same unresolved handler.
Bear also retains inherited damage-based outgoing/incoming rage; this has not
been established for Forever (DRU-012). Warrior measurements are not a verified
replacement formula for Bear.

The [selection ledger](../artifacts/tanks/current/validation.json) includes
the frozen controls, per-seed checks, uncertainty bounds, rejection reasons and
pooled per-target threat. Compressed research records preserve the candidate
inputs and both encounter results. No combat-engine mechanic was changed for
these selections.
The ledger records seeds per build: Paladin's later correction has separate
[validation and archives](paladin_mana.md). Mana-restoration threat remains an
inherited assumption tracked as CORE-014, not verified Forever threat behavior.

## Reproduction

The [archive index](../artifacts/tanks/current/archives.json) lists every screening,
gear-search and confirmation record with checksums. Each gzip contains a `files`
object keyed by the original filename. `final-confirmation-a` and `-b` contain
complete players, requests and results for both independent seeds.

For example, replay the retained Orc Warrior and both workloads:

```sh
go build -tags with_db -o /tmp/forever-bench ./tools/forever_bench
python3 - <<'PY'
import gzip, json
from pathlib import Path
ledger = json.loads(Path('artifacts/tanks/current/validation.json').read_text())
row = next(r for r in ledger['selections']
           if r['Key'] == 'tank_warrior' and r['Race'] == 'Orc')
with gzip.open('artifacts/tanks/current/final-confirmation-a.json.gz', 'rt') as f:
    records = json.load(f)['files']
Path('/tmp/tank-player.json').write_text(json.dumps(records[row['SelectedID'] + '.player.json']))
PY
/tmp/forever-bench -build tank_warrior -race Orc -natural-hit \
  -player /tmp/tank-player.json -tank-pair -iterations 20000 \
  -seed 20262501 -output /tmp/tank-confirmation
```

Use the ledger's `ControlID` to extract and replay the matched control; use
seed `20262763` for the other confirmation. Paladin runs use 100,000 iterations
per seed. To replay the complete chart and all three sensitivities:

```sh
python3 tools/forever_bench/run_matrix.py --binary /tmp/forever-bench \
  --profiles artifacts/modelled_gear/forever_input_profiles.json --natural-hit \
  --iterations 5000 --seed 20291951 --workers 24 --output /tmp/tank-matrix
```

See the [uncertainty register](uncertainties.md), particularly the class tank
entries and scenario/model limitations, before treating these measurements as
predictions of live tank balance.

Release checks cover the benchmark, maintained mechanics/source tests, all
three tank agents and browser replay. The full inherited suite still has stale
fixtures and metadata expectations; that separate coverage gap is SCEN-014.
