# Change-history review

> This review describes its pinned revisions. Current dispositions, including
> later resolutions of H10, H11, H20 and H29, are in the
> [canonical register](uncertainties.md).

## Boundaries and coverage

The common ancestor of this checkout and `wowsims/classic` is
`7779ebbf79dc7f1341e6ab939b28a3402c9a730a`.
The actual imported ElliotWood/Forever revision is
`a7622df37e2d046010cc207cf6da3419f5c9ba9b`, not the later upstream revision
that was previously present in the remote-tracking branch.

The imported range has 345 first-parent commits. Its commit inventory has
been reviewed, followed by the cumulative core, Mage, Priest, Shaman, Rogue,
Warrior, Hunter, Paladin, Warlock and Druid runtime diffs. Follow-up findings
are tracked below. Local review also covers the benchmark runner, hit ledger,
profile/default import, sensitivity accounting, vendor/catalog compiler and
enchant replay changes from the four local commits. This is
not a claim that every mechanic or every historical intermediate version
has been validated.

Later upstream changes through
`12f81fc56d183ca98a0f8e1161f1adc426ac1ce9` are leads, not authority for
game behavior. In particular, a new implementation of Warrior rage does
not establish its level-60 or off-hand coefficients.

Local changes reviewed so far:

| Commit | Review scope |
|---|---|
| `79fcaeccb` | Initial engine changes: shared ratings, resource cadence, racials, Tier 1 overrides, channels, autos, item effects and class additions |
| `91bba4199` | Browser preservation of cross-faction buffs and replay regression |
| `15d28d946` | Equipment/set/enchant runtime, weapon-layout validation and totem handling |
| `29a06bce2` | Enchant preparation and replay changes; generated numerical output is historical evidence, not new mechanics |

Commit metadata does not establish which model produced the first local
commit. No model attribution is inferred from the author field.

## Findings

“Fixed” below means included in the published correctness release
`c65503774`, with targeted regression coverage. It does not mean that every
associated game mechanic has been measured in the beta.

| ID | Finding and provenance | Disposition |
|---|---|---|
| H01 | `77983c00a` introduced persistent Hot Streak cast reduction. The local `79fcaeccb` added Frostfire eligibility but retained that assumption. Client aura 400625 has one charge and three stacks. | Fixed: completed Pyroblast consumes the aura. |
| H02 | `2c34ea1b2` inferred one damage per three healing. `79fcaeccb` avoided some double-counting but retained the fallback for individual item/enchant components; `29a06bce2` also used it for enchant scoring. | Fixed: no inferred conversion. Actual hybrid effects are imported separately. |
| H03 | Mage Clearcasting checked the already-discounted cost and therefore did not expire after a free spell. The condition originates in pre-fork `a9f4e029f`, not a Forever-specific change. Later upstream `d66773b84` exposed it. | Fixed: check underlying paid damage-spell eligibility. Use client aura 12536 and its trigger's one-second internal cooldown. |
| H04 | The Priest audit `fe554d98a` updated Starshards damage but retained its missing cooldown. Later upstream `12f81fc56` identified the omission. `SpellCooldowns` specifies 30 seconds. | Fixed: one timer shared by every rank and partial-channel variant; cancellation does not bypass it. |
| H05 | Hunter pet Bite rank 8 retained a maximum of 91. Client effect 17261 has base 90 and variance 0.2: 81–99. | Fixed: damage range and sampled-result regression. Pet owner-stat inheritance is a separate unresolved issue. |
| H06 | Inherited projectile speeds were reused for updated Hunter shots; Shadow Bolt used the travel helper without a projectile speed. | Fixed from `SpellMisc`: Auto/Aimed/Arcane/Sting 40, Multi-Shot 30, Sniper Shot 60, Shadow Bolt/Incinerate 20 yards per second. |
| H07 | `bfc763dad` changed the Improved Stormstrike regeneration multiplier without updating cached mana ticks. Subsequent unrelated stat changes could make the bonus appear or linger inconsistently. | Fixed: refresh regeneration on both activation and expiration. |
| H08 | Spirit Tap had the same cached-regeneration ordering problem, inherited from pre-fork `a185f837f`. | Fixed: refresh after both stat and casting-regeneration changes. It does not create a new proc source or kill event in the benchmark. |
| H09 | Upstream later changed five Feral/Rogue finishers to 80% Energy refunds. A code change alone does not verify the server refund rule. | Awaiting source/test; do not adopt from the commit description alone. |
| H10 | Penance deliberately distributes three bolts evenly across two seconds instead of the stated immediate/one-second/two-second timing. | Awaiting T19; no partial-channel optimization uses this approximation. |
| H11 | Shadow Word: Death is missing. Later upstream adds four client-backed ranks but leaves its script effect and health backlash unmodeled. | Awaiting complete effect interpretation; do not silently copy an incomplete damage-only implementation. |
| H12 | Sanctity Aura was supplied globally, although neither the current trait tree nor `SkillLineAbility` provides a learn path for spell 20218. Legacy spell/talent records still exist. Later upstream `7fe7a3453` suppresses it; the imported Paladin audit already noted its removal. | Fixed: exclude the unavailable buff from Forever raid, party and personal settings and defaults. T36 can establish a currently unknown acquisition path. Divine Spirit differs: it retains Priest class-skill entries, though trainer access after talent removal still needs confirmation. |
| H13 | `6ec35f115` omitted Puncturing Wounds from Mutilate's off-hand strike. Client talent 1224716 effect 2 targets family mask 6, including both Mutilate child spells 1241586 and 1241590 (mask 2). | Fixed: both hands receive 5% crit per talent point; regression covers ranks 0–3. |
| H14 | Mutilate consumes Cold Blood on the main-hand result and does not apply the off-hand penalty to its flat damage component. Those choices are not established by the parent spell's tooltip. | Awaiting a two-hand damage/crit capture; retain current behavior rather than infer it from another expansion. |
| H15 | Warrior Whirlwind already calls normalized weapon damage. Heroic Strike/Cleave carry the auto proc mask as well as the special mask, so Unbridled Wrath includes them already. | Already correct for the cited implementation claims. This does not validate the separate level-60 rage-generation model or every proc rule. |
| H16 | `8c9d313c2` attached Resourcefulness/Rapid Recuperation as bare pseudostat changes, leaving cached mana ticks stale. The same helper was used by Blue Dragon's proc. | Fixed: a regeneration-aware aura helper updates ticks immediately on gain/expiration. Hunter regressions cover both transitions for both talents. |
| H17 | Exorcism, Holy Wrath, Holy Shock and Holy Shield allocated one cooldown timer per rank. These predate the Forever fork; Holy Shock's independent timer is visible in `d7f33bb0a`. Client categories 19, 35, 892 and 931 explicitly group the ranks. | Fixed: one timer per ability; real casts prevent every lower rank from bypassing it. |
| H18 | Holy Shield did not require an equipped shield. Every current rank's `SpellEquippedItems` entry requires armor subclass mask 64. | Fixed: enforce the current off-hand shield at cast time. Positive and negative equipment cases are covered. |
| H19 | Summon Hawk approximates an undocumented guardian with periodic copies of the dive-bomb base and models only one Hawk, despite a two-Hawk limit in the tooltip. | Awaiting T14; not a verified guardian model. This existing approximation remains disclosed rather than replaced with another invented attack profile. |
| H20 | Templar's Bulwark is represented as 99% damage reduction instead of a finite absorb of maximum health. | Unsupported as an exact tank model; not used by the DPS benchmark. The current core shield helper records shielding but does not consume a finite absorb pool. Tank-survival validation needs that separate implementation. |
| H21 | Soul Fire, Death Coil, Shadowburn and Conflagrate also allocated separate rank cooldowns, despite client categories 631, 633, 651 and 672. | Fixed: one timer per ability, with real-cast regressions for every registered rank. A structural check across all 23 baseline spellbooks found no remaining split timers among 85 positive-cooldown category groups, including 49 multi-rank/type groups. Shocks were already correct through their shared cooldown, despite also having individual timers. |
| H22 | Decimation, introduced in `6ec709410` and revised in `16c01ffd1`, tied Shadow Bolt/Searing Pain damage to the ten-second Soul Fire proc. Client 440870 makes that damage conditional on the execute threshold, not on the proc being active. | Fixed: the damage bonus starts at the 35% phase, including the first qualifying hit. Proc activation/expiry changes only Soul Fire's cast time. Regression covers both transitions independently. |
| H23 | Firestone retained an old weapon proc and omitted spell crit; its wrapper required an empty off-hand. Spellstone incorrectly granted crit instead of its current power/haste effects. | Fixed: legal main-hand enchants with occupied off-hands; no stacking with another imbue. At level 60 Firestone grants 21 Fire power and 2% spell crit. Spellstone grants 21 power with the client's Fire+Shadow mask and 2% casting haste. The latter mask conflicts with its Shadow-only tooltip: T42 remains open. No saved profile selects either stone. |
| H24 | `6ec709410` added Demonic Brand with flat 39–42 demo damage. Description variables 1016/1020 contain its display formula; child variables 1017/1018 use the matching school's power. Its pet aura also let attacks on other targets consume charges and excluded pet spells. | Fixed supported components: level-60 base 65–68 plus 7.8% matching-school power; pet modifiers once, target-scoped charges, direct pet spells included and no second miss roll. Owner power at hit time, noncritical proc handling, Voidwalker/Felhunter school assignment and the 3× threat multiplier remain provisional: T39. |
| H25 | `87c1a41c6` introduced Cat Mangle under TBC ID 33876; the later audit acknowledged that its Forever source was missing but retained it. The port also extended Berserk's Mangle cleave to Lacerate, whose family mask is not selected. | Fixed: remove the unsupported Cat spell and restrict the cleave to Mangle. The saved Feral profile uses neither Cat Mangle nor the Mangle talent, so its rotation is unchanged. |
| H26 | `bebe4b3bc` reduced Nature's Grace GCDs only for hardcasts, by division through 1.1. Client 16886 has a distinct −10% GCD modifier covering instant Balance spells and Faerie Fire too. | Fixed: exact 10% reduction for the supported matching spells, separately from 10% cast haste; expiry restores the original values. Actual Moonfire casts are tested. Wrath's inherited cast-completion proc timing remains open in T41. |
| H27 | The Warlock UI offered a Fel Armor tooltip that actually selected the same Demon Armor enum as the other entry. Fel Armor retains a spell and acquisition-method-3 skill row, but that does not establish Forever availability. | Fixed the misleading duplicate selector. No new Fel Armor effect is granted without an acquisition source; T43 records the check. |
| H28 | Furor's out-of-form timer ignored exits at time zero or before the pull. Its saved Energy and exit timestamp also survived iteration reset. | Fixed: explicit no-exit sentinel and per-iteration reset, with regression cases for negative, zero and positive exit times. Lower-rank carry/cap semantics still need T15; the saved Feral build has no Furor points. |
| H29 | Improved Moonfire and Moonfury retain base-damage-only multipliers. General haste still does not shorten the engine's default spell GCD. Neither rule is established merely by the spell's familiar name or its similarity to another expansion. | Awaiting controlled SP/talent and haste/GCD measurements (T26, T40). Do not silently switch either rule to retail behavior. |
| H30 | Spell-source validation exposed integration defects: the AST walker treated non-ID values in Frostfire/Dark Sacrifice rank arrays as spell IDs and lost conditional Brand/Eureka IDs. The manifest also lacked descriptions and had inconsistent assumption labels. | Fixed: positional fields are resolved by name, branch alternatives are retained, and literal package-helper arguments are followed. The metadata/coverage suite passes. Thirteen generic registration sites remain unresolved by this static scanner; 44 metadata entries remain explicitly unreviewed. This is not a claim that every tooltip is verified. |
| H31 | The review generator described old starting-gear baselines as same-engine comparisons without checking their provenance. The preset regression also compared current inputs to historical optimization bundles. | Fixed: current review tables report DPS, standard error and resource limitations, without mixing historical search gains into corrected results. The regression now checks all 147 actual web defaults against current recorded requests and verifies their talent/APL sources and gear legality. |
| H32 | `79fcaeccb` retained vendor-only records outside the reviewed catalog for older comparisons. This left 64 provisional armor records and two cosmetic tabards in the live database after better source-filtered equipment became available. | Fixed: the live equipment pool is the 1,908-record reviewed catalog, including its authoritative vendor overlaps. All 94 distinct selected item records are unchanged; 147 deterministic replays preserve requests, stats, DPS, resources and canonical metrics after this availability cleanup. Historical requests remain archived with their code revision. |

The local import review found no additional nonzero vendor stats silently
discarded: the only field outside the stat map is weapon DPS, whose authoritative
damage range and speed are imported separately. Shared hit/crit are imported
into one equipment pool; explicit benchmark adjustments are applied after
sharing, avoiding a second conversion. Item/suffix stats and weapon damage
scale in the sensitivity scenarios while enchants and proc effects stay fixed.
The frozen 147 profiles preserve all selected gear and enchants. Their only
profile changes are five Elemental talent strings, six Shadowform prepull
timings and three Ret aura selections. No selected enchant is healing-only.

Client records supporting these corrections are retained in
`assets/db_inputs/forever_effect_audit.json`; tests are in
`tools/forever_bench/history_fixes_test.go` and the class-specific regression
files. The [mechanics review](mechanics_review.md) covers the other corrections
identified through source and class review.

For Demonic Brand, the child-spell `Attributes_3` flag `0x40000` is decoded as
[Always Hit](https://github.com/TrinityCore/TrinityCore/blob/master/src/server/game/Miscellaneous/SharedDefines.h).
This reference supplies the client-field label, not a private-server damage
formula. The damage formulas and proc mask come from the captured Forever tables.

## Resource checks

Matched saved profiles, 5,000 iterations, seed 20292581. Equipment, talents,
APL, buffs, consumes and encounter were held constant. These compare the
pre-history-review working engine with the first history corrections, not
the older published release.
The table below predates the Sanctity Aura exclusion and Hunter regeneration
repair; it is kept as a record of those isolated checks, not the current matrix.

| Build / race | Before DPS | After DPS | After OOM time |
|---|---:|---:|---:|
| Arcane / Orc | 705.33 | 681.71 | 1.28 s |
| Fire / Orc | 715.08 | 605.90 | 31.90 s |
| Frost / Orc | 697.87 | 627.67 | 24.10 s |
| Marksmanship / Orc | 851.55 | 852.03 | 0.03 s |
| Survival / Orc | 791.36 | 779.92 | 0.00 s |
| Destruction / Orc | 784.87 | 785.17 | 0.00 s |
| Smite / Night Elf | 600.13 | 600.13 | 0.00 s |
| Mutilate / Orc | 608.68 | 618.25 | 0.00 s |

The Mage change reveals resource pressure previously hidden by repeated free
casts. Fire and Frost already use mana potions, gems and Evocation; those
supplies are not missing from the checked profiles. No fight shortening or
extra mana was added. Survival includes the earlier Expose Prey restriction
as well as the projectile/Bite corrections. The saved Smite rotation does not
use Starshards, so its cooldown repair does not change this paired result.
The Mutilate pair isolates the subsequent off-hand Puncturing Wounds repair
on top of those changes: +9.57 DPS, or 1.57%.

These checks are not final rankings and do not validate a new rotation.

### Buff and regeneration follow-up

Matched 5,000-iteration checks, seed 20292582, with the same saved profiles.
These add Sanctity Aura exclusion, corrected Hunter mana ticks and Paladin
cooldown/equipment checks:

| Build / race | Before DPS | After DPS | After OOM time |
|---|---:|---:|---:|
| Retribution / Human | 1069.75 | 1004.71 | 0.00 s |
| Smite / Night Elf | 600.13 | 552.63 | 0.00 s |
| Marksmanship / Orc | 852.03 | 852.03 | 0.03 s |
| Survival / Orc | 779.92 | 779.92 | 0.00 s |
| Beast Mastery / Orc | 882.58 | 882.58 | 0.00 s |

The Holy-damage builds lose the unsupported buff. The checked Hunter profiles
do not gain DPS from the regeneration repair; it is not presented as a
performance improvement merely because the implementation was wrong.

### Warlock follow-up

Matched 5,000-iteration checks, seed 20292583, retain the saved Orc profiles.
The changes are shared rank cooldowns, Decimation's execute damage and the
Warlock effect corrections:

| Build | Before DPS | After DPS | After OOM time |
|---|---:|---:|---:|
| Affliction | 804.50 | 804.50 | 0.01 s |
| Demonology | 811.80 | 812.47 | 0.00 s |
| DS/Ruin | 764.71 | 764.71 | 0.00 s |
| Destruction | 785.17 | 785.17 | 0.00 s |

The fixed profiles do not select weapon stones. These are matched correctness
checks, not a search for stronger talent or consume combinations.

### Druid follow-up

Matched 5,000-iteration Tauren checks, seed 20292584:

| Build | Before DPS | After DPS | After OOM time |
|---|---:|---:|---:|
| Balance | 600.89 | 604.01 | 0.00 s |
| Feral | 671.34 | 671.34 | 0.00 s |

Balance reflects the Nature's Grace GCD correction. The Feral profile uses
neither Mangle nor Furor; those corrections are covered by targeted mechanic
tests rather than an expected change in this benchmark.

### Auto-attack follow-up

The role-wide casting exception introduced in `79fcaeccb` also bypassed
ordinary spell swing resets. This was broader than the intended weapon-cast
exceptions. The [auto-attack audit](auto_attack_audit.md) records the restoration
of the original Classic reset helper, explicit Slam/Hunter shot exceptions,
regression checks and the matched 588-run replay. Both hardcast and instant
Lightning Bolt follow the original Shaman reset hook; the latter remains an
explicit in-game check in T48.

The local rebuild also exposed a build dependency omission: imported UI JSON
files did not trigger rebundling. They now do, so regenerated ranking data
cannot leave stale DPS labels in the selectable defaults.
