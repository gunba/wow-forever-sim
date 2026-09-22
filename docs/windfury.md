# Windfury and Retribution

## Matched comparison

Human Retribution, 300 seconds against the level-63 Dragonkin reference target.
Equipment, enchants, talents, APL, other buffs, ordinary MP5 and the paid hit
adjustment are identical. Removing Windfury does not add a replacement imbue.
Each arm has two independent seeds with 5,000 iterations per seed.

| Damage per second | Windfury off | Windfury on | Change |
|---|---:|---:|---:|
| Total | 920.01 | 1,035.18 | +115.18 / +12.52% |
| Ordinary and extra auto-attacks combined | 193.11 | 256.78 | +63.67 / +32.97% |
| Seal of Command | 192.40 | 208.24 | +15.84 |
| Seal of Righteousness | 190.34 | 223.45 | +33.12 |

About 46 extra swings occur per fight. Ordinary swing count falls slightly
under the inherited extra-attack timing, so the full 69.13 DPS attributed to
extra swings is not the net white-damage gain. The remaining damage difference
comes from small changes in ability usage and resource/proc outcomes.

The total gain with Tier 1 disabled is 12.90%. Tier 1 is therefore not the main
explanation for this effect. These figures describe this loadout and rotation,
not a universal Windfury percentage.

## Implementation checks

- The modeled totem has a 20% chance on eligible landed main-hand attacks,
  adds one extra attack, and supplies the current rank's 246 AP.
- Holy Strike and Judgement of Command are eligible melee specials. Seal of
  Command's proc is also eligible in the inherited model. Seal of Righteousness
  damage is excluded by its equipment-proc suppression flag.
- Extra white swings can produce Seal damage. The 1.5-second Windfury ICD stops
  immediate Windfury recursion; Command also has its own one-second ICD.
- Selecting the party buff, the weapon setting, or both produces identical
  damage and action results with matched seeds. A duplicate permanent-aura
  initialization inflated the *tracking aura's activation count*, not damage.
  Registration is now idempotent and has a regression test.

The AP value is client-supported. The inherited ICD, eligible Seal interactions,
batching and AP-charge resolution are not thereby established as live Forever
behavior. They remain [T28 and T30](in_game_checks.md), not reasons to alter the
model until measured.

## Reproduction

[`artifacts/windfury/runs.json.gz`](../artifacts/windfury/runs.json.gz) contains
the complete requests and results, including the single-fight event trace.
[`summary.json`](../artifacts/windfury/summary.json) retains per-action accounting.
The request files can be replayed with the native CLI:

```sh
go build -tags=with_db -o /tmp/wowsimcli ./cmd/wowsimcli
/tmp/wowsimcli sim --infile request.json --outfile result.json
```
