# Forever DPS benchmarks

The benchmark contains 23 level-60 builds and 147 race/build combinations:
71 Alliance and 76 Horde. The [client roster and racial audit](../../docs/forever_races.md)
record availability and remaining interaction questions.
Profiles use the [crafted/dungeon catalog](../../docs/forever_gear_data.md),
paid shared-hit normalization, and the role's complete
[Forever Tier 1 bonuses](../../docs/forever_tier1.md).
They are starting builds, not a claim of optimal talents, rotations or gear.

Run from the repository root:

```sh
python3 tools/database/compile_forever_equipment.py
go run ./tools/database/gen_db -outDir=assets -gen=db
python3 -m unittest discover -s tools/database -p 'test_*forever*.py'
go test -tags with_db ./tools/forever_bench
go run -tags with_db ./tools/forever_bench -baseline-results artifacts/forever_dps_5min.json \
  -iterations 5000 -seed 20291951 -output /tmp/forever-replay
python3 tools/forever_bench/fetch_icons.py
python3 tools/forever_bench/chart.py artifacts/forever_dps_5min.json
```

Chart rendering requires Python's `matplotlib`. Go builds can use
`-buildvcs=false` when source-control metadata is unavailable.
The matrix has fixed race columns and spec rows ordered by their peak mean DPS,
highest first, with class labels and spec icons. A dash means the race/class
combination is unavailable; a missing available result is an error.
Use `-faction horde` or `-faction alliance` to limit the benchmark; the chart
accepts the matching `--faction` filter.

### Equipment search

`assets/db_inputs/forever_ilvl65_items.json` contains all 706 distinct entries
from the Wowhead item-level-65 list, including the rows beyond the first rendered
page. It is a discovery pool, not an availability whitelist. Raid rewards and
unresolved item effects remain excluded. Verified catalog equipment at item level
60 and above supplements the list, which omits cloaks. Source-verified lower-level
trinkets also remain eligible: the known level-65 Undermine trinkets share a
one-item limit. Original and resumed equipment stays available across matching
slots, so replacing one ring or trinket does not discard it from the other slot.

```sh
go build -tags with_db -o /tmp/forever-bench ./tools/forever_bench
/tmp/forever-bench -build enhancement -race Orc -search-gear \
  -gear-screen 100 -gear-validate 1000 -gear-passes 3 \
  -iterations 5000 -seed 20291931 -output /tmp/gear-search/enhancement
```

Run or resume the complete roster with bounded parallelism:

```sh
python3 tools/forever_bench/search_gear.py --binary /tmp/forever-bench \
  --output /tmp/gear-search/all
```

`--representatives` selects one race per build for an initial coverage check.
The driver fingerprints the binary and runtime inputs before reusing completed
jobs. Its summary lists failed or unconverged searches separately.
It defaults to one independent native process per available CPU; `--workers`
can override that limit.
The native search retains its original starting profiles; replay published
loadouts with `-baseline-results` rather than rerunning those starting profiles.
The search driver also accepts `--baseline-results` to continue from saved
winners, and `--builds balance,feral` to restrict a follow-up pass.

Replay selected equipment and all paired sensitivities in parallel:

```sh
python3 tools/forever_bench/run_matrix.py \
  --binary /tmp/forever-bench --profiles /tmp/gear-search/all/results.json \
  --original-baselines /tmp/gear-search/all/baseline.json \
  --output /tmp/forever-final-matrix
```

The final replay fills missing enchants without overriding legal simulated
choices. Initial enchant scoring respects the inherited healing-to-damage
conversion, including healing enchants; that conversion remains an in-game
validation question. Proc and haste enchants use their simulated comparison.

The search compares each slot against the current loadout, with legal weapon
layouts compared together. It then compares legal enchants and the second
profession; Engineering remains available for the existing consumables.
Each candidate pays for hit again. The strongest three improving screen results
receive independent-seed validation, with conservative two-standard-error
acceptance bounds. Passes stop when no meaningful gain survives validation.
This is coordinate search, not proof of a global optimum.

The per-build/race `.gear-search.json` contains pool exclusions, every trial's
equipment, second profession, seed and iteration count, accepted changes, and
complete baseline/final requests and results. Coordinate IDs 0–16 identify item
slots; 17–33 identify enchants on slots 0–16; 34 identifies the second profession.

Tier 1 stays enabled throughout the search. Ordinary equipment sets use current
client-derived effects and actual piece/profession requirements rather than the
inherited Classic registrations. Unimplemented active thresholds are omitted and
listed in `UnmodeledSetBonuses`; they do not prevent equipping otherwise legal
items. Item proc gaps remain separate from these equipment-set gaps.

### Tier and gear sensitivity

The matrix's three gain columns use equal-weight arithmetic means of each
available race's percentage gain. Tier 1 gain is `100 × (on / off − 1)`.
Gear gains are `100 × (scaled / baseline − 1)`, with Tier 1 still enabled.
All comparisons retain the same talents, APL, encounter, seed and 5,000 iterations.
They measure the current build's response, not a reoptimized build at each gear level.

`Player.equipment_scale` scales item and random-suffix stats, weapon min/max damage
and item flat bonus damage in per-character copies. It does not alter catalog
records, enchants, weapon speeds/skills, proc effects, set effects or external
buffs. Hit normalization is recalculated against the scaled offensive item budget.
This is a hypothetical proportional-upgrade model, not actual future equipment or
individual stat weights. In particular, Cat-form weapon-DPS contribution is still
an unresolved mechanic, and resource thresholds can make gains nonlinear.

Reproduce the comparisons from the exact saved, unnormalized baseline players:

```sh
go build -tags with_db -o /tmp/forever-bench ./tools/forever_bench
/tmp/forever-bench -baseline-results artifacts/forever_dps_5min.json -seed 20291951 -tier1=false -output artifacts/sensitivity/tier1_off
/tmp/forever-bench -baseline-results artifacts/forever_dps_5min.json -seed 20291951 -equipment-scale 1.1 -output artifacts/sensitivity/gear_110
/tmp/forever-bench -baseline-results artifacts/forever_dps_5min.json -seed 20291951 -equipment-scale 1.2 -output artifacts/sensitivity/gear_120
python3 tools/forever_bench/sensitivity.py
python3 tools/forever_bench/chart.py artifacts/forever_dps_5min.json --sensitivity artifacts/forever_sensitivity.json
python3 tools/forever_bench/build_review_site.py
```

The summary includes source hashes, individual race comparisons and conservative
95% Monte Carlo bounds. Since common seeds correlate the runs and races, these
bounds sum marginal standard-error contributions instead of assuming independent
errors. They do not account for uncertain game mechanics.

The Tier-off run exposed a Eureka charge underflow in nested Warrior attacks.
The engine now reserves a charge before resolving an attack's callbacks, so a
nested attack cannot spend the same final charge. The baseline and all three
scenarios were rerun after this correction. Earlier outputs are retained under
`artifacts/sensitivity/before_charge_fix.*`; in-game Eureka scope remains an open check.

The benchmark uses level-60 characters, a level-63 target with no creature-type
bonuses, 3,731 starting armor, five minutes without duration variation, a fixed
external buff/debuff package, and no world buffs. Armor is a benchmark assumption, not
a measured beta boss value. Hit adjustments cover physical specials/ranged shots
and rotational magic with one shared allocation; they do not cap dual-wield white
attacks or remove dodge. Actual talents and racials count toward each cap.
These caps follow the fork's combat tables, not measured beta miss rates.
Hunter traps retain the fork's special rule that ignores gear hit; their
remaining miss chance cannot be removed by the normalization.

Every profile has a complete legal loadout; only the off-hand slot is empty
when a two-handed weapon occupies both hands. Crafted, dungeon and verified
PvP/vendor equipment are eligible; old-raid-derived gear is excluded.
Current client/planner records replace provisional
armor. Exported vendor values retain precedence if an overlapping item is
used. Generic AP applies to melee and ranged attacks.

Tier 1 is an explicit override, independent of the items equipped. It grants
the role's 2/3/4/5-piece bonuses once, without adding raid-item stats.
Creature restrictions remain active. `-tier1=false` provides a comparison
without the override. Ordinary equipment-set bonuses still require their
actual pieces; the starting profiles avoid unreviewed active equipment sets.

Each JSON result includes the original player, full normalized simulation request,
hit-budget ledger, final stats, active sets, APL warnings, action/resource metrics,
and out-of-mana time. CSV results include Monte Carlo standard errors and budget
totals. Build definitions read talent presets and APLs from the UI rather than
duplicating talent strings.

Current JSON outputs also include a `Mechanics` record. The
[auto-attack model](../../docs/auto_attack_audit.md) keeps weapon-role autos
running during ordinary casts and uses client shot cast times without added
wind-up. Energy regeneration uses smooth integration and assumes general haste
scaling; attack-speed-only effects are excluded. That assumption remains an
open in-game check.

`artifacts/forever_mechanics_baseline_5min.json` is the current frozen-input
reference: 147 combinations at 5,000 iterations, with no APL warnings and paid
hit-budget balances within floating-point rounding. Its CSV and matrix chart
are adjacent. These are baseline rotations, including historical shot-window
conditions; they are not the final optimized results.

## Saved Horde baseline

`artifacts/horde_dps_5min.json` preserves 5,000 iterations for each of the
76 Horde combinations before the all-race optimization pass. All
requests completed without APL warnings, and every profile reports its four
role-specific Tier 1 thresholds. Hit-budget balances differ from zero by less
than 0.000000000001 budget points.

Mana limits are included, not disabled. Hunters use the sustain changes described
below; each result records its exact mana-limited time. This metric measures
failed mana-cost checks, not a period of zero damage. Auto-attacks and pets
continue. These results do not establish a damage ceiling for the builds.

`artifacts/forever_optimization_baselines.json` freezes the 23 unadjusted input
profiles. Comparisons must rerun those inputs under the same engine as the
candidates, so a racial or buff correction is not counted as an optimization gain.
The [in-game checks](../../docs/in_game_checks.md) separate level-20 tests from
later checks for class reviewers.

The [energy audit](../../docs/energy_audit.md) replaces Classic-style Energy
ticks with smooth-recovery integration. `artifacts/energy_cadence_5min.json`
preserves the timing-only comparison for 31 Feral/Rogue race combinations.
It predates haste-scaled regeneration and is not the current comparison
baseline. Rerun frozen inputs under the candidate engine when assessing gains.

Across all four Horde Hunter races, the saved Horde run records 0.77–1.05 seconds
of mana-limited time for Beast Mastery and 1.39–2.66 seconds for Marksmanship,
down from roughly 31–32 and 20–24 seconds respectively. The encounter therefore
remains five minutes for every build.

The web UI defaults to a 300-second encounter and offers matching gear, talent
and rotation presets. Its ordinary character settings do **not** silently apply
the benchmark's hit-budget adjustments or replace its class-specific buff
defaults. Use the saved requests for exact benchmark reproduction.

### Browser replay

Export a results file into settings accepted by the web simulator's **Import →
JSON** dialog:

```sh
python3 tools/forever_bench/export_ui_profiles.py \
  --results artifacts/forever_dps_5min.json --output artifacts/ui_profiles \
  --bundle ui/core/forever_ranked_profiles.json
python3 tools/forever_bench/build_review_site.py
```

Each file includes the race, talents/APL, equipment, consumes, buffs, encounter,
seed, Tier 1 setting and paid-hit adjustment. Load it in the matching class/spec
page. `index.json` records each file's expected mean DPS and hit ledger.
Browser worker partitioning can produce small Monte Carlo differences.

The hit adjustment is stored as fixed bonus stats. Changing gear, talents or race
does **not** automatically recalculate it in the UI; rerun the native benchmark
to obtain a new, fairly normalized comparison. Do not treat these bonus stats as
additional obtainable gear.

Run the review-site step **after** the web build. It stages a clickable matrix,
replay downloads, raw results and audit notes at `/classic/review/` in the built
site. The same access protection must cover this directory and all simulator
assets when deployed.

### Hunter mana handling

These changes describe the historical Horde snapshot. The
[build reviews](../../docs/build_reviews.md) describe the current retained builds.

- **Beast Mastery:** Efficiency increases from 1/5 to 5/5, funded by three
  Improved Stings points and one Mortal Shots point. The build remains 31/20/0.
  Intimidation is used above 80% mana instead of automatically on every cooldown.
- **Marksmanship:** Aimed Shot uses rank 1. Forever keeps normalized weapon
  damage at every rank but reduces the flat bonuses: rank 1 adds 20 for 75 mana;
  rank 6 adds 166 for 310 mana, before Efficiency. Rank 1 preserves most of the
  weapon hit at a much lower cost. It is a learned rank, not an altered spell.
- Gear, hit budgets, potion/rune choices and their restoration rules are
  unchanged. Both builds already used Major Mana Potions and Demonic Runes.
  Earlier manual consumable timing and alternative flask/talent configurations
  did not provide a better overall trade-off and were not retained.

## Hit-budget model

`linear-v1` is an explicit benchmark model, **not a verified Blizzard item-price
formula or an in-game reforging system**:

| Stat quantity | Budget points |
|---|---:|
| 1 raw hit rating | 1 |
| 2 attack power | 1 |
| 7/6 spell power or explicit spell damage | 1 |
| 1 relevant Strength, Agility or Intellect | 1 |
| 1 raw crit rating | 1 |

The [client table](../../docs/forever_gear_data.md#hit-budget-accounting) supplies
the separate rating conversions: ten hit rating per percentage point and
fourteen crit rating per percentage point.

The cost is distributed proportionally across eligible offensive **item**
allocations, not chosen from the least useful stat for each class. Casters use
Intellect, spell power/damage and crit; physical builds use attack power, Agility,
Strength and crit. Hunters exclude Strength except for Survival. Hunters,
Enhancement and Retribution also include their spell power/damage and Intellect.
School-specific power, Spirit, Stamina and unpriced secondary stats are not
exchanged. Enchants, base attributes, buffs, consumes, racials and sensitivity
bonuses never fund the budget.

Generic hit is bought once and added to both attack types. Generic AP and crit
are likewise charged once even though they populate two simulator fields.
Excess **item** hit is exchanged back into the same offensive pool; no more than
the item's original hit allocation can be refunded. Tauren's hit therefore
reduces the debit, or increases the refund, rather than being wasted.

Required hit follows casts in the APL, automated major cooldowns and equipped
Rogue poisons. Unused registered spells do not erase school-specific hit talents.
Incidental racial/item procs and pets are not independently normalized. The
ledger lists the selected requirements and trap exclusions. Hybrids can have
unavoidable physical overcap when their spell cap requires more shared hit.

Normalization returns a clone, leaving talent-search inputs and item records
unchanged. A normalized request can be replayed directly. To run another search,
use `BaselinePlayer`, not the already-adjusted player from `Request`.

Examples:

```sh
go run -tags with_db ./tools/forever_bench -build enhancement -race Orc -iterations 2000
go run -tags with_db ./tools/forever_bench -build enhancement -race Orc -optimize -rounds 3 -output /tmp/enhancement-search
go run -tags with_db ./tools/forever_bench -spell-power 50 -output /tmp/forever-sp50
```

Talent search preserves build identity and rejects illegal point allocations.
It screens single-point reallocations, then validates promising changes with
an independent seed. It returns candidates; it does not rewrite presets.
These are local searches, not proof of a global optimum.
The final `-optimize` result uses the requested seed plus 30,000, separate from
both screening and validation; the saved request records that actual seed.
When `-output` is set, a `.talent-search.json` sidecar records every trial's
talents, seed, iteration count, DPS, standard error, mana-limited time and warnings,
plus the original player and benchmark scenario. Trial inputs are reconstructed
from that original player with the recorded talents, allowing hit normalization
to be recalculated rather than reusing another build's adjusted stats.

`screen_rotations.py` runs the APL families in `rotation_candidates.py`.
It keeps all other player fields unchanged, screens with 250 iterations,
validates up to five candidates with an independent 2,000-iteration seed,
and confirms a proposed change with a third seed at 5,000 iterations.
Both searches require a gain above 2.5 combined standard errors to retain a
candidate. Final cross-race comparisons remain a separate step.

```sh
go build -buildvcs=false -tags with_db -o /tmp/forever-bench ./tools/forever_bench
python3 tools/forever_bench/screen_rotations.py \
  --binary /tmp/forever-bench --build balance --race Tauren \
  --output /tmp/balance-rotations
```

The output directory contains candidate inputs, complete requests/results and a
`ledger.json`; its path must be new. Retained candidates do not replace UI presets
automatically.

`record_build.py` combines the search evidence with final cross-race comparisons
in `artifacts/optimization/`. It checks fixed player inputs and requires changed
builds to clear the acceptance margin on every included race.
`build_reviews.py` generates [class-review summaries](../../docs/build_reviews.md)
with talent choices, prepull actions, rotation priorities, damage breakdowns and
mana-limited time. The [frozen all-race baseline](../../artifacts/forever_baseline_5min.json)
contains 5,000 iterations for each of the 147 combinations under the audited
racial and faction-buff rules.

```sh
python3 -m unittest discover -s tools/forever_bench -p 'test_candidates.py'
python3 tools/forever_bench/build_reviews.py
```

`-apl`, `-gear`, `-talents` and `-player` allow single-build experiments.
Spell-power sensitivity is an explicit stat adjustment to every build, not
an assertion about unreleased gear or trinket effects.

## Evidence and assumptions

- The [shared-effects integration audit](../../docs/beta-pass/core.md) tracks
  runtime fixes and unresolved omissions. The earlier Horde snapshot is
  `artifacts/horde_dps_5min.{json,csv,svg,png}`. Earlier files with
  `vendor_gear` in their names describe the previous equipment scenario.
- Vendor tooltips establish the item conversions and weapon ranges. The
  overlapping class armor validates the planner conversion, but does not
  independently verify every other class's armor.
- `docs/beta-pass/warlock.md` records the client's all-magic Curse of the
  Elements and removal of Curse of Shadow. The shared aura now includes Holy
  and Nature, and each resistance changes once rather than once per school.
- [Curse of Recklessness](https://www.wowhead.com/forever/spell=11717/curse-of-recklessness)
  reduces armor by 505 and grants no attack power. Its metadata and regression
  test match the Forever tooltip.
- Pet dismissal cancels a queued combat-start auto-swing event. This does not
  remove Demonic Pact's separate ability to retain a sacrifice buff with a
  different demon.
- APL channel interruption evaluates both `interruptIf` and the next spell's
  readiness after hypothetical cancellation. The next-action check was repaired
  first; a separate regression then exposed the same problem in conditions using
  `spellCanCast`. Both now permit interruption at a legal tick/GCD boundary.
  Ordinary channels, GCDs and cast restrictions are unchanged.
- Warlock sustain assumes healing can cover Life Tap. The non-tank damage model
  records its mana return without deducting health; the mana result is not a
  claim that tapping has no healing cost.
- Incinerate retains the fork's base-damage-only Immolate bonus. Whether that
  bonus should also multiply spell-power damage needs beta confirmation.
- Existing low-rank spell coefficients, trap hit rules and Summon Hawk's
  guardian approximation are retained, not replaced with Classic or retail
  expectations. They remain limitations of the model.
- Warrior auto-attack rage is still damage-based Classic rage. The exact
  Forever formula and combat-table changes remain unverified.
- Frostfire Bolt now has its three ranks and explicit talent interactions.
  Undead Priest Dark Sacrifice now supplies its five mana ticks and debits
  health; its self-damage modifier/resistance interactions remain unverified.
- No new class weapon-proficiency rules are imposed by this tool. Candidate
  weapons use the existing UI's ordinary proficiencies; possible Forever
  extensions have not been independently verified.
- Tier dummy-effect stacking and Insect Swarm's fractional final tick remain
  assumptions described in the [Tier 1 notes](../../docs/forever_tier1.md).
  Item and dungeon/crafted-set effects outside the reviewed starting profiles
  have not all been audited.
- UI links, icon lookups and external tooltips now use Wowhead's Forever
  database, including its `CLASSICPLUS` widget environment (16), rather than
  Classic Era. Local mechanic notes still identify unresolved values.
- [Consumable checks](../../docs/beta-pass/consumables.md) record the corrected
  Grilled Squid, Runn Tum and Nightfin effects. Nightfin no longer grants MP5.
- The engine retains an inherited one-third spell-damage fallback for
  healing-only records. It does not apply to the selected healing items:
  each already has an explicit damage component.
