# Simulator review — 26 September 2026

## Scope and progress

The pinned release is `51cc3d6e2c22b341b15d23a7e37be78099e90b21`.
The review window starts at 18 September, 00:00 AWST. It contains 142
reachable commits: 40 local and 102 inherited, including merge and changelog
commits. There are 117 first-parent entries. The boundary before the window is
`8bce57ac55fcc1d47040fafeb2f68ad6478eee00`.

The commit and [Flurry evidence](flurry_review.md) reviews are complete. The
[142-row coverage ledger](weekly_review_commits.csv) records dispositions.
Review coverage is not a claim that every server mechanic is verified.

Reproduce the inventory:

```sh
git log --reverse --since='2026-09-18T00:00:00+08:00' \
  --format='%H %aI %s' 51cc3d6e2c22b341b15d23a7e37be78099e90b21
```

| Area | Status |
|---|---|
| Commit inventory and local/inherited boundary | Complete |
| Core resource, damage, pet inheritance, cast and periodic-damage changes | Reviewed; confirmed defects and explicit assumptions separated below |
| Buff/debuff and consumable cumulative changes | Reviewed; Expose Armor and Blood Pact defects reproduced |
| All nine class engine diffs and shared item/set-effect diffs | Reviewed; inherited proc/scaling assumptions remain disclosed |
| Tank defensive abilities and magical utility GCDs | Reviewed; confirmed fixes below, incomplete tank features disclosed |
| Equipment import and projected-gear selection | Armor import and missing shield Block corrected; generator and current profile accounting checked |
| Benchmark/UI integration and non-engine commits | Current paths reviewed; private-data tooling corrections below |
| Class source checks and item-effect coverage | Dispositions recorded; unsupported effects are not silently enabled |
| Updated benchmark | All 736 DPS scenarios and 24 paired tank requests replayed |

The current equipment/default cross-check covers all 184 DPS profiles: their
embedded source hash matches `artifacts/modelled_gear/forever_dps_5min.json`,
all manual stat bonuses and hit exchanges are zero, and no Feral profile has
mail or plate. The 504-item projection and UI metadata reproduced exactly
before the shield correction. Equal projected cost remains a disclosed
budget assumption, not proof of equal DPS or a verified Blizzard formula.

## Priority and affected paths

- **Tank correctness:** Shield Block recasts and missing intrinsic shield
  block value changed mitigation; Bulwark's inherited GCD changed rotation
  timing (W1, W2, W8).
- **Shared engine correctness:** native armor validation now rejects illegal
  imports; Blood Pact's level-60 amount and automatic party scope are fixed
  (W4, W14).
- **Damage/resource correctness:** utility spell GCDs, refreshed Expose
  Armor, multi-target Agony and Unbridled Wrath eligibility are corrected
  (W3, W5, W6, W15).
- **Incomplete equipment effects:** Nature's Bounty, nine ordinary-set MP5
  thresholds and three retained weapon-proc amounts are repaired. These are
  not new effects on the projected ranking gear (W7, W11, W13).
- **Tooling and privacy:** database builds, cache-name scrubbing and failed
  client-data comparisons are corrected (W9, W10, W12).

The frozen benchmark is being replayed, not retuned around these fixes.
Old research comparisons remain evidence for their recorded engine
revision, not newly validated gains on this one.

## Confirmed corrections

### W1 — Shield Block refresh lost charges

**Fixed.** After spending one charge, casting Shield Block
again at six seconds refreshed its seven-second aura but left only one charge.
Charges were assigned in `OnGain`, which is not called when an active aura is
refreshed. Each successful cast now restores both charges without adding the
75% block bonus twice.

- Affected code: `sim/warrior/shield_block.go`.
- The Forever two-charge/seven-second configuration was introduced in
  `b865c9235`; the five-second cooldown exposes the refresh problem.
- Client `SpellAuraOptions` 1.60.1.69977, spell 2565: two charges.
- Regression: `TestShieldBlockRecastRestoresCharges`.
- Isolating this fix before the later shield/stat corrections, the Prot
  Warrior golden scenario changed from 528.074 to 533.082
  DPS, 1138.188 to 1149.495 TPS, and 528.735 to 520.577 DTPS. These are that
  test's results, not a replacement tank benchmark.

### W2 — Templar's Bulwark incorrectly used a global cooldown

**Fixed.** The ability inherited a 1.5-second GCD. Client
spell 1311015 has zero `StartRecoveryTime` and no GCD category. It can now be
cast while the ordinary GCD is running, without changing that timer. The
absorb, cooldown and Forbearance still apply.

- Affected code: `sim/paladin/templars_bulwark.go`.
- The finite-absorb correction in `b865c9235` retained the older cast timing.
- Sources: `SpellCooldowns` and `SpellCategories` 1.60.1.69977, spell 1311015;
  `SpellMisc` 1.60.1.70009 identifies the school as Holy.
- [Forever spell entry](https://www.wowhead.com/forever/spell=1311015).
- Regression: `TestTemplarsBulwarkDoesNotUseGlobalCooldown`.

### W3 — Missing schools bypassed the spell-GCD rule

**Fixed.** Several utility actions had
`SpellSchoolNone`. The school-based Forever GCD policy introduced in
`374c75267` therefore treated them like fixed-GCD physical abilities.

The corrected registrations are:

| Ability | School | Spell IDs |
|---|---|---|
| Innervate | Nature | 29166 |
| Lightning Shield | Nature | 324, 325, 905, 945, 8134, 10431, 10432 |
| Aspect of the Hawk, level-60 rank | Nature | 25296 |
| Evocation | Arcane | 12051 |
| Shadowform | Shadow | 15473 |
| Soul Link | Shadow | 19028 |
| Holy Shield | Holy | 20925, 20927, 20928 |

Client schools come from `SpellMisc` 1.60.1.70009. Each listed action has a
1500-ms `StartRecoveryTime` in `SpellCooldowns` 1.60.1.69977. Regression tests
cast them with increased casting speed and check their effective GCD.
This does not change physical GCDs, the one-second floor, or channel-haste
policy.

- Regressions: `TestForeverUtilitySpellSchoolsAndGCDs`,
  `TestHolyShieldUsesHastedHolyGlobalCooldown`.
- Physical-school forms and Warrior buffs were not relabeled as magical
  merely because their original registration omitted a school.

### W4 — Raw imports could still put mail or plate on a Druid

**Fixed.** The picker and gear-search filters respected
class armor limits, but native/WASM character construction did not. A direct
request containing a modeled mail helm on Feral successfully simulated.
Character construction now enforces the same armor proficiency ceiling.
The gear-search ceiling uses that shared definition.

- Affected code: `sim/core/character.go`,
  `sim/core/equipment_validation.go`.
- This was a pre-existing engine validation gap left uncovered by the
  material/selector correction in `755c08a29`, not evidence that the
  current selected Feral profiles contain plate.
- Regression: `TestFeralImportRejectsMailAndPlate`.
- This check does not claim to validate every item restriction, profession,
  enchant, weapon proficiency or low-level training requirement.

### W5 — Upgrading Expose Armor did not upgrade the actual armor reduction

**Fixed.** Replacing one-combo-point Expose Armor with a
five-point application changed the effect's priority field without notifying
the active effect system. Against 3731 armor, the target stayed at 3281
instead of dropping to 1481. Expiration could then restore the larger amount
that had never been removed.

The spell now calls `SetPriority`, which removes the old reduction and applies
the new one. Expiration restores exactly the original armor.

- Affected code: `sim/rogue/expose_armor.go`.
- Origin: `a791b03fc` (4 March 2024), before this fork. This is an inherited
  defect found while reviewing the recent armor/debuff work, not a new local
  regression.
- Regression: `TestExposeArmorUpgradeChangesReductionAndRestoresArmor`.

### W6 — Agony's ramp could borrow another target's application bonus

**Fixed.** Each rank held one shared ramp increment for
all its targets. Applying ordinary Bane of Agony to a second target replaced
the increment previously calculated for an amplified application on the
first target. The first target's later ticks therefore lost part of Amplify
Curse's bonus.

The increment is now stored per target's Dot instance. The regression applies
an amplified and an ordinary Bane, then verifies both damage ramps separately.
At the tested rank, the amplified first target incorrectly advanced to 57.5
and 80.5 base damage instead of 69 and 103.5.

- Affected code: `sim/warlock/curses.go`.
- Origin: `5cd498570` (4 April 2024), before this fork.
- Regression: `TestAgonyRampRetainsEachTargetsAmplifyBonus`.
- This is shared-state isolation, not a new coefficient or snapshot rule.
  A single-target benchmark does not exercise this failure.

### W7 — Nature's Bounty fallback used the wrong resource events and timing

**Fallback corrected; catalog proc coverage remains incomplete.** The
Wildheart fallback granted all available resources when struck, instead of
separating spellcasts, outgoing melee attacks and incoming melee attacks.
Feralheart separated those events but granted its 20 Energy instantly.

Both now share the distinct triggers and a five-tick Energy effect: 4 Energy
each second. They retain the existing, explicitly unverified 2% proc-rate
assumption. The client dummy aura's 100% proc field is not used as a guarantee.
Any additional server-side active-form restrictions remain unverified.

- Origins: dungeon-set revisions `3d4229afc` and `3571a9207`.
- Sources: [Forever spell 450608](https://www.wowhead.com/forever/spell=450608);
  captured `SpellEffect` 70009, spell 27784 (periodic energize, 4, 1000 ms);
  `SpellMisc` 70009 duration index 28; `SpellDuration` 69913 (5000 ms).
- Regression: `TestNaturesBountyRoutesResourcesAndTicksEnergy`, covering both
  name-based set registrations.

Importantly, current catalog set IDs take a separate, conservative route in
`sim/core/forever_equipment_sets.go`: unsupported proc thresholds are omitted,
even where an older name-based registration exists. This correction does
**not** enable that unresolved proc in the catalog or change modeled-gear
rankings. Confirmed trigger timing and unknown proc probability must remain
separate. A visible implementation in a class file is not proof that current
item records actually activate it.

### W8 — Imported shields lost their intrinsic Block value

**Fixed.** `GetItemStats` and the saved gear-planner
metadata omit the standalone Block line. The importer therefore assigned
zero base Block to all 16 usable catalog shields, and the gear projector
copied that omission into both modeled shields. This reduces mitigation
and the Block contribution to Shield Slam.

The two level-65 PvP shield exports explicitly contain **44 Block**.
[Shield Wall](https://www.wowhead.com/forever/item=272591),
[Barricade](https://www.wowhead.com/forever/item=278469), and
[Evergreen Shield](https://www.wowhead.com/forever/item=279262) independently
show the same value. All 17 catalog shield tooltips were captured; Jagged
Obsidian Shield still remains excluded for its unrelated unresolved stat.
Evidence URLs and response hashes are in
`assets/db_inputs/forever_shield_block.json`.

The fix preserves intrinsic Block separately from allocated bonus Block:
it is not charged against the discretionary stat budget, just as base
armor is not. Vendor tooltip values take precedence on overlapping items.
Both modeled shields retain their original stats and budget, gain 44 Block,
and display it in their names and tooltips. No proc was inferred or added.

- Origin: the catalog/vendor paths introduced in `79fcaeccb`; propagated
  into projections by `ff7b01f28`; newly important for `b865c9235` tanks.
- Tests cover vendor precedence, base versus bonus Block, projection
  preservation, and the final native item database. The native test first
  reproduced zero Block on all five checked real/modeled level-65 shields.
- Database comparison: only Block changes on 16 real and two current
  modeled items; only the two modeled item names additionally change.

### W9 — The database Make recipe included test files in `go run`

**Fixed.** Adding `gen_db/forever_equipment_test.go` in
`c65503774` made the existing `go run tools/database/gen_db/*.go` command fail
with “cannot run *_test.go files.” Normal site builds did not exercise that
recipe. It now runs the package (`go run ./tools/database/gen_db`); database
generation and its targeted tests pass.

### W10 — Cache scrubbing missed short and non-ASCII names

**Fixed.** The inherited browser and Python readers
required at least three ASCII characters. Two-character and UTF-8 names
were silently left intact. Synthetic records reproduced both failures;
the readers now recognize printable UTF-8 and preserve **byte** lengths,
not character counts, when replacing names. Damage records and offsets
remain unchanged.

The fork had already disabled the upstream upload endpoint, so this was
not an automatic upload of private data. However, the page still promised
private-bucket uploads and complete name removal. It now offers a local
download and explicitly warns that unknown records, pet names and other
identifiers can remain. The unused upload controls/stub are removed.

- Origins: Python cache reader `f01728bbb`, browser port `cbcff5df9`, upload
  workflow `9a7c5ec9e`.
- Regressions use invented short/UTF-8 names only; no private log is copied
  into tests or the repository.
- The existing Cloudflare backend is not deployed or enabled by this change.

### W11 — Nine equipped-set mana bonuses were unnecessarily disabled

**Fixed.** The set compiler accepted passive
power-regeneration aura 85 only when `EffectAuraPeriod` was 5000. These
passive stat effects actually have a zero period; they are not periodic
energize auras.

The captured client effects and Forever tooltips agree:
[18378](https://www.wowhead.com/forever/spell=18378) grants 8 MP5,
[21636](https://www.wowhead.com/forever/spell=21636) grants 12 MP5, and
[21625](https://www.wowhead.com/forever/spell=21625) grants 3 MP5.
This enables the six-piece bonuses on Magister's, Devout, Dreadmist,
Wildheart, Beaststalker, Elements and Lightforge, plus three-piece Bloodsoul
and two-piece Green Dragon Mail.

- Origin: the ordinary-set compiler in `15d28d946`.
- Actual piece counts are required. These are not added to the forced
  Tier 1 setting, and modeled items do not acquire ordinary set IDs.
- All nine generated thresholds are checked; no other generated effect
  changes. Proc chances remain unresolved and are not enabled by this fix.
- Both the compiler test and native support/threshold test reproduced the
  omission before the correction.

### W12 — The client-change report could hide gaps or mislabel spells

**Fixed.** Reproduced problems affected the
research tooling, not automatic engine imports:

- All failed fetches could produce a headline saying “no changes,” above
  individual error messages; an HTML error response could also be read as
  an empty table. CSV validation and incomplete-comparison labels now
  prevent these false clean results.
- SpellEffect row IDs could be looked up as spell IDs, attaching an
  unrelated ability name. Labels now follow the row's `SpellID`.
- The table list omitted spell costs/categories, target restrictions,
  talent effect curves, and item/set effect links. Those tables are now
  included.

This helper predates the review window (`3119b2d35`, September 17). Its
limitations matter when interpreting later client-change reports. Existing
reports have not been retroactively treated as complete scans, and no
mechanic is changed merely because a table diff finds a row.

### W13 — Three weapon proc amounts missed the 70009 patch

**Fixed.** Comparing the inherited implementations
with both captured client builds found three additional numeric changes:

| Effect | Previous implementation | Build 1.60.1.70009 |
|---|---:|---:|
| Baron Charr's Sceptre, Firebolt 13442 | 35 | 90 |
| Hand of Timmy 17505 | 210 | 85 |
| Linken's Sword, Lightning Bolt 18089 | 45–75 | Mean 98, variance 0.5 (73.5–122.5 in the simulator) |

Source rows are `SpellEffect` 691018, 695399 and 695940.
[Hand of Timmy's current tooltip](https://www.wowhead.com/forever/spell=17505)
also states 85. Timmy's redesign in `8306a8173` was correct for the older
client but stale after the patch. Proc rates have not been guessed afresh.
All three registered effects failed the new isolated damage regression
before correction.

These items are absent from the current catalog and modeled profiles.
The test supplies temporary fixtures; it does not make the items available
or grant these procs to projected weapons.

### W14 — Blood Pact omitted level growth and leaked across parties

**Fixed.** `88fd19d64` changed the shared buff to its
level-50 base, 49 Stamina, without its 0.5-per-level growth. Rank 5 at 60
provides **54 Stamina**, or 70 with the existing fully improved modifier.
Both captured `SpellEffect`/`SpellLevels` rows and the
[level-60 tooltip](https://www.wowhead.com/forever/spell=11767/blood-pact)
agree.

The older automatic Imp provider also wrote into raid buffs even though
the spell is a party aura. It now supplies its own party only. Explicit
external raid-wide coverage settings are unchanged. Regressions reproduce
the 49-versus-54 error and the missing difference between two parties.
Partial talent ranks and pet summon/death transitions are not newly
implemented by this static provider correction.

### W15 — Unbridled Wrath also triggered on queued specials

**Fixed.** Heroic Strike and Cleave deliberately carry
both auto and special proc masks. Matching any white-hit bit therefore
granted them Unbridled Wrath rolls, contrary to the linked low-level
observations. Forever now excludes special-tagged replacement attacks from
this talent without changing the masks used by other proc systems.

The new regression reproduces both illegal triggers and keeps ordinary
main/off-hand hits eligible. The intended 12%-per-point rate is retained;
the acknowledged lower-rate beta bug is disclosed separately. See the
[Flurry review](flurry_review.md#separate-defect-found-in-the-linked-research)
for sources and scope.

## Checks that did not justify a mechanic change

- Holy Shield's four charges, 20% block bonus, and rank-three 221 damage plus
  0.08 coefficient match the captured client records. The GCD registration
  was wrong; these values were not.
- Shield Slam's rank-four client effect really contains coefficient 1.
  Its presence alone is not evidence of an accidentally copied coefficient.
- Ordinary Forever periodic crit checks use current crit chance. Reading the
  retained snapshot fields alone would incorrectly suggest otherwise.
- The shared crit aura is combined with an OR, not added twice when both
  provider flags are present.
- Demonic Rune self-damage is excluded from encounter damage and outgoing
  damage/threat totals while retaining damage-taken behavior.
- The 23 merge commits have no independent combined-resolution diff. Their
  contributing changes were reviewed in the ledger; a merge is not counted
  as a second implementation.
- Another 36 entries only change the changelog. The remaining 83 entries
  comprise 38 with simulator code and 45 integration, data, tooling or
  documentation entries. The 45 integration entries now have reviewed
  dispositions, including superseded homepage/launch presets, exact-profile
  preservation, historical enchant assumptions and the findings above.
  The 38 engine-related entries now have source/implementation dispositions
  in the coverage ledger. These counts are not a claim that every beta
  mechanic is verified.

## Follow-up evidence questions

- **Nature's Grace and GCD stacking:** the current implementation applies
  the explicit client GCD modifier and also runs through the general
  cast-speed/GCD rule. Those are two separate reductions in the model.
  The [spell entry](https://www.wowhead.com/forever/spell=16886) and captured
  effects identify the modifier, but do not establish whether the server
  also applies the casting-speed aura to GCDs in this case. No speculative
  removal has been made.
- **Dungeon-set proc coverage:** 44 thresholds remain disabled after the
  stat-effect reconciliation. Several need proc/script identifiers or
  effect amounts, not just an item-set description. Nature's Bounty's
  retained legacy proc chance also still needs a Forever source. Legacy
  code alone was not used to enable every remaining effect.
- **Seal of Righteousness scaling (T57):** the retained implementation adds
  its SP contribution twice and uses a two-handed coefficient multiplier
  justified by old SoD testing. The exported coefficient alone cannot
  verify that server behavior. This is a high-priority unresolved scaling
  assumption, not a newly verified Forever formula.

## Evidence and publication limits

On 26 September, both the GitHub API and Git transport returned “not found”
for `wowsims/forever`. Search results still list the repository, which is not
proof that its latest source was retrieved. The locally captured upstream
reference is `ea5412873474d5fbb057a751a97c35bf725dce22`; no claim about newer
upstream commits is made here.

The review has not established a complete Seal of Fury implementation or
resolved the already documented server-script, pet, rage and item-effect
assumptions. Those remain separate from reproduced programming errors.

## Refreshed results and verification

- All **184 baseline profiles and 552 paired sensitivities** completed at
  5,000 iterations with the recorded seed, zero warnings and unchanged
  requests. Gear, talents, rotations, consumables and encounter inputs were
  not reselected.
- Fury's mean change across races is **−2.04 DPS**, with individual changes
  from −2.69 to −1.24. This refresh is a correctness change, not a new build
  search or a claim that all remaining assumptions are settled.
- All **24 tank requests** were replayed unchanged with two independent
  seeds. All **17 race/tank combinations** also passed a native smoke run.
  The updated tank table is in [Tank simulations](tank_benchmark.md).
- Previous requests/results, sensitivities and tank comparisons are retained
  under [`artifacts/history/51cc3d6e2`](../artifacts/history/51cc3d6e2/).
- Targeted core, class, benchmark, tank and database checks pass with
  `with_db`. Source-registration checks pass, along with 44 database,
  14 benchmark-tool and six data-watch Python tests.
- The old Fury test fixture still selected the removed Windfury weapon
  imbue and had stale expectations. It now selects a legal stone; its
  generated expectations and the affected tank fixtures were refreshed.
  These fixtures are not the ranking loadouts. Their suites are now part
  of deployment CI.
- Web/WASM builds and browser checks cover page loading, bundled icons,
  the +50% Gnome Fury replay, rejected Shaman dual wield and fresh defaults
  for all three tank routes. Tank DPS, TPS and DTPS match native results.
  The cache-scrubbing browser check uses synthetic names and verifies local
  download without an upload request.
- The workbook includes intrinsic shield Block. Nature's Bounty's three
  resource effects now have declared tooltips and bundled display metadata,
  with their modeled resource amounts and proc-chance caveat.

Deployment status is available in the repository's
[Build and Deploy workflow](https://github.com/gunba/wow-forever-sim/actions/workflows/deploy.yml).
