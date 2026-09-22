# Mana regeneration

## Optional interpretation

The **MP5 acts per second** option is under **Settings → Other**. It changes
the flat MP5 contribution from `MP5 / 5` mana per second to `MP5` mana per
second. It is an unverified beta assumption, off by default and off in the
published ranking profiles.

This covers the MP5 stat from equipment, buffs and consumables, including
temporary stat changes. Spirit regeneration, casting regeneration modifiers,
potions, Evocation and other direct or percentage-based returns retain their
existing rules. Mana still arrives on the existing two-second ticks; the
option changes the rate, not the tick interval. It has no effect under the
Classic ruleset.

The setting is stored in the player profile and survives export/import and
saved settings. Loading a ranked profile restores that profile's recorded
setting. Item records, their tooltips and the paid-hit adjustment do not change.

## Matched Mage comparison

Five-minute single target, 5,000 iterations per arm, seed 20293001, six races
per build. The only request difference is the explicit setting. Equipment,
enchants, talents, APL, consumes, buffs, Tier 1 and paid hit are identical.
There are 36 completed runs with no APL warnings.

| Build | Default DPS | Alternate DPS | Mean relative gain | Mana-limited time, default → alternate |
|---|---:|---:|---:|---:|
| Arcane | 683.30 | 713.52 | +4.42% | 1.21 → 0.00 seconds |
| Fire | 608.78 | 714.28 | +17.36% | 30.21 → 0.00 seconds |
| Frost | 633.25 | 700.60 | +10.66% | 22.02 → 0.00 seconds |

DPS and time columns average races equally; the gain column averages each
race's relative change. Mana-limited time records failed mana-cost checks,
not an equal amount of time doing zero damage or the timestamp of first OOM.
The alternate removes these waits in every tested Mage profile. It does not
establish that the game uses the alternate rate.

The tested Arcane example has 65 MP5: its flat regeneration changes from
13 to 65 mana per second. A dynamic addition of five MP5 changes both casting
and noncasting regeneration by one mana per second in default mode and five
in alternate mode; Spirit is unchanged. Under Classic, either setting retains
the original one-mana-per-second increment.

Raw requests, metrics, hit ledgers and uncertainty are in:

- [Default-mode results](https://gunba.github.io/wow-forever-sim/classic/review/mana_regen/mp5_off.json)
- [Alternate-mode results](https://gunba.github.io/wow-forever-sim/classic/review/mana_regen/mp5_per_second.json)
- [Paired comparisons and source hashes](https://gunba.github.io/wow-forever-sim/classic/review/mana_regen/summary.json)

See [T44](in_game_checks.md#t44--flat-mana-regeneration-rate) for an in-game test
that separates flat mana recovery from Spirit and direct mana gains.
