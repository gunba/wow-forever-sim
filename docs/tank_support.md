# Tank raid support and Judgement timing

All 17 tank profiles now use one [raid-support definition](../sim/core/forever_tank_support.json)
in native requests, web defaults and the ranking picker. Gear, talents,
professions, workload/healing settings and rage formulas are unchanged.
The 184 non-tank profiles and their results are unchanged.

## Corrected settings

- All tanks receive the same available raid buffs, Kings/Might/Wisdom, shared
  3% crit aura, Curse of Elements and the existing tank debuff support.
  Curse of Elements includes Holy damage in Forever.
- Paladin now uses Brilliant Wizard Oil; Warrior retains its sharpening stone.
  These coexist with the party Windfury buff. Bear's weapon coating remains unset.
- Paladin also receives the defensive support previously enabled for the other
  tanks. Redundant Fire Resistance Totem and Bear's redundant external Sunder
  selection are removed. Obsolete improvement labels for Fortitude, Gift of the
  Wild, Might, Wisdom and Battle Shout are normalized to their current base effects.
- Warrior supplies Battle Shout and Sunder instead of receiving external
  Battle Shout/Expose Armor. External Thunder Clap and Demoralizing Shout/Roar
  remain disabled: existing self-maintenance is preserved.

The scenario assumes enough raid/group providers for its buffs, not one
character casting every aura or totem. Windfury and Grace of Air are never
combined; only one shared crit aura source is selected.

## Swift Judgement

The engine already reset Judgement and made the next cast free. The defect was
rotation timing: automatic cooldown use could reset a nearly finished timer
and let another cast intervene. The explicit priority is now Judgement, then
Swift when at least six seconds of Judgement cooldown remain, then the existing
cooldown/defensive/maintenance priorities. Both actions are off the GCD.
This avoids spending the reset immediately before natural cooldown recovery.

The targeted regression verifies five resets in five minutes, each followed
immediately by a zero-cost Judgement, without returning to seal-refresh spam.
Matched APL-only comparisons keep all support settings identical.
They gain 5.4–5.9 DPS across the three Paladin races and pass the existing
survival/total-threat/least-target-threat checks in both workloads.

Representative five-minute means from two overlapping-seed 10,000-iteration runs
per arm (these are validation seeds, not the chart's seed):

| Profile | Previous support/APL | Corrected |
|---|---:|---:|
| Human Protection Paladin | 450.82 | 548.27 |
| Undead Protection Paladin | 467.42 | 567.52 |
| Dwarf Protection Paladin | 443.24 | 539.42 |
| Orc Protection Warrior | 386.62 | 386.79 |
| Tauren Bear | 640.98 | 638.82 |

Warrior's change is negligible. Bear gains the previously missing Insect Swarm
support: fewer landed boss attacks also reduce incoming rage in its current
model, so better defensive support need not increase its personal DPS.

## Flasks and unresolved mechanics

Titans remains the main tank flask. Supreme Power is already selectable as an
offensive tradeoff: it exchanges 1,200 base health for 150 SP. It is not a
missing free buff, and the damage advantage does not establish a survivability
advantage. This simulator continues counting damage after a modeled death.

Bear still uses the inherited, unverified damage-based rage model (DRU-012).
Maul correctly replaces its white swing; frequent Maul is not the previously
fixed Warrior direct-cast bug. No Warrior rage coefficients were inferred for
Bear. Threat, shield, enchant/form and boss/healer assumptions remain in the
[question register](uncertainties.md).

## Evidence and replay

- [Historical paired comparisons](../artifacts/tank_support/validation.json)
  cover every race and both tank workloads, with separate Paladin APL-only controls.
- [Archived complete requests/results](../artifacts/tank_support/archives.json)
  include the 68 refreshed baseline/Tier-off/+10%/+50% runs.
- [Preservation checks](../artifacts/tank_support/preservation.json) retain all
  736 non-tank scenario rows unchanged.
- [Historical tank settings/results](../artifacts/history/807db6aaea/tanks_before_shared_support.json.gz)
  remain separate. Earlier gear/talent selection gains are not current gains
  under the corrected support.

Archived `.pair.json` files contain `Single.Request` and `Multi.Request`.
Extract either request and replay it with:

```sh
go build -tags with_db -o /tmp/forever-tank-replay ./tools/forever_tanks
/tmp/forever-tank-replay -request request.json -output result.json
```

Current searches use the fixed [shared-support controls](../artifacts/tanks/shared_support_controls.json),
not the candidate itself or the historical unequal-support profiles. These
corrections do not claim a new global gear/talent optimum.

**Later review:** these two starting seeds differ by only 38. Because the engine
increments the seed every iteration, the runs overlap and are not independent
replications; pooled uncertainty needs review (SCEN-016). Protection Warrior's
rotation and current results are superseded by the
[Sunder maintenance correction](protection_sunder.md).
