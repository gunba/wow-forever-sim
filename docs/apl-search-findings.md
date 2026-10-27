# APL and simulator findings from DPS searches

Developer handoff, started 27 October 2026. The experiment uses the Forever ruleset, level-60 characters with modeled item-level-65 gear, and 300-second encounters. The requests omit player level; [`PlayerLevel`](../sim/core/base_stats.go) therefore uses the level cap, currently 60. The native simulator binary used for the observations had SHA-256 `f9c7ce22c3c9b2c1af428de8bf2f38dfe86dd47844b52b1d889a50051d60d1f1`. The saved requests below are self-contained `RaidSimRequest` files with one logged iteration; the full search records and statistical comparisons remain in the original run directories.

The exact submitted agent reports, including later reports, are preserved in [`apl-search-reports.jsonl`](apl-search-reports.jsonl). Duplicate coordinator-forwarded copies and a diagnostic-only test report are omitted. The findings below distinguish behavior reproduced in the simulator from rules that still require in-game verification. The builds to repeat after relevant fixes are tracked in [`apl-search-rerun-queue.md`](apl-search-rerun-queue.md).

## Resolution in the simulator (27 October 2026)

The findings below are kept as reported. The simulator now handles them as follows; the regression tests are in [`tools/forever_bench/apl_search_findings_test.go`](../tools/forever_bench/apl_search_findings_test.go) and use the requests in [`apl-search-repros`](apl-search-repros).

| Finding | Status |
|---|---|
| 1. Hostile effects before the pull | **Flagged.** Anything that reaches an enemy before 0 (a hit, miss or resist, a DoT tick, or an aura put on it such as a curse) gives a warning on the pre-pull action responsible. The simulated fight still starts at 0, and damage before 0 was already left out of the metrics, but DoTs, debuffs and timers started early carry into the fight, so a flagged result is not valid; the search evaluator rejects it. Exempt: Hunter's Mark, Battle Shout (its hits on the target stand in for party threat), placing a fire totem (Jordan confirmed it does not start the fight; a totem's attacks before 0 are flagged) and a trap laid before the pull, which goes off at 0. The rule goes by what the simulation does, not by spell properties: casting each of 556 APL spells alone before the pull, on every build, showed that a rule based on the defense type would have refused Life Tap and Holy Wrath, which do not reach a non-undead target, and missed curses, Faerie Fire, Sunder Armor and Demoralizing Shout. In the published matrix only BM, MM and SV Weave for Windshaper and High Order are flagged: their faster Aimed Shot, cast at −2 s, hits at −0.03 s. Their presets move to −1.9 s with the next benchmark run. |
| 2 and 15. Ranks of one DoT overlap | **Fixed.** Jordan confirmed that one caster's ranks do not stack. A new rank of a DoT now ends the caster's other ranks of it on that target (all ranked DoTs sharing a spell code, such as Corruption, Frostfire Bolt, Immolate and Moonfire; channels already end each other). |
| 3. Warnings not visible to the search | **Fixed.** The search evaluator returns the rotation warnings, and the benchmark counts them. Pre-pull failures (for example a stance the character is already in, or Aimed Shot inside its minimum range) are found by running the pre-pull up to 0, so they now show too. Jordan confirmed the AQ ranks (Arcane Missiles 8, Frostbolt 11, Shadow Bolt 10) cannot be learned, as modelled. Every learned rank of Arcane Wrath can now be cast; the ranks share the cooldown. |
| 4. Tiger's Fury | No change needed: Jordan confirmed it is removed for Forever Feral, as modelled. |
| 5. Pre-pull Autocast Other Cooldowns | **Fixed.** Refused with a warning. |
| 6. Obsolete `doAt` in the APL documentation | **Fixed.** |
| 7. Curse of the Elements | No change needed (the raid debuff was already on). |
| 8 and 13. Unresolved conditions | **Fixed.** An aura the unit cannot have is never active and has no stacks or time left, and a spell the character does not know cannot be cast, is never ready and has no DoT up; both give a warning. Any other part of a condition that cannot be evaluated (such as target health in a fight without a health bar) stops its action from ever running, with a warning. The Wing Clip rank 1 guard is now false, and Soul Fire on target health never runs. |
| 9. Range cliff | Unchanged (the 12-yard minimum is the modelled rule); movement messages now print the distance. |
| 10. No distance value | **Pre-trapping added.** Jordan confirmed Hunters can pre-trap: a trap laid before the pull goes off at the pull, wherever the Hunter stands, if it was laid within its minute on the ground. During the fight the target does not move, so a trap still only reaches it from melee range. There is still no distance value. |
| 11. Pet mana after a resummon | **Fixed.** The current-pet mana values read whichever demon is active, and are 0 with none. |
| 12. Pre-combat Stealth | **Fixed.** Every learned Stealth rank can be cast before the pull, without a GCD; it stops auto attacks. |
| 14. Fire Nova origin | No change needed: Jordan confirmed Fire Nova needs a Fire totem, as modelled; the code comment now says so. |
| 16. Rage before the pull | Starting with 50 rage is confirmed. Out-of-combat decay is unresolved, so the model is unchanged and Bloodrage before −10 s stays excluded. |

## 1. Hostile effects can land before the encounter starts

**Reproduced in the simulator; encounter legality needs enforcement or explicit definition.** The APL can apply hostile effects at negative encounter times while the encounter clock and DPS denominator still start at zero.

- [`prepull-bane-of-doom.json`](apl-search-repros/prepull-bane-of-doom.json): Bane of Doom (`603`) is cast and applied at **−59 s**; its 60-second timer completes at **+1 s**, producing damage during the nominal five-minute fight. Other hostile effects are also applied before zero.
- [`prepull-immolate.json`](apl-search-repros/prepull-immolate.json): Immolate (`11668`) begins at **−5 s**, lands and deals direct damage at **−3.5 s**, and ticks at **−0.5 s**. Its remaining DoT then continues into the nominal fight. The candidate was selected by the original search wrapper, so its reported improvement must not be treated as a valid combat result without resolving pull legality.
- [`prepull-starfire-haste.json`](apl-search-repros/prepull-starfire-haste.json): Starfire (`9876`) starts at **−3 s** and hits at **−0.09 s** because its actual cast time is shorter than its nominal three seconds. The experiment's smoke-trace safeguard flagged this candidate; a −2.5-second cast was observed landing after the pull. Legitimate pre-casts need to account for cast-time modifiers as well as projectile travel.

Relevant code: [`sim/core/apl.go`](../sim/core/apl.go), [`sim/core/sim.go`](../sim/core/sim.go), [`sim/warlock/curses.go`](../sim/warlock/curses.go), [`sim/warlock/immolate.go`](../sim/warlock/immolate.go), and [`sim/hunter/hunters_mark.go`](../sim/hunter/hunters_mark.go). `SpellFlagEncounterOnly` is described in [`sim/core/flags.go`](../sim/core/flags.go) as an APL UI signal rather than an enforced rule.

**Suggested check:** define when a damaging or hostile effect starts combat and ensure the simulation cannot obtain uncounted pre-pull combat time. Preserve legitimate self-buffs, Hunter's Mark before combat, and normal pre-casts whose effects land at or after the pull. Other non-damaging pre-pull target debuffs may have different in-game rules and need separate verification.

Hunter's Mark is a permitted pre-combat exception, not an issue: Jordan confirmed it does not trigger combat. [`prepull-hunters-mark.json`](apl-search-repros/prepull-hunters-mark.json) records its non-damaging target aura at **−1.5 s**. This does not establish the legality of other pre-pull target effects.

A Searing Totem trace needs separate interpretation: [`prepull-searing-totem.json`](apl-search-repros/prepull-searing-totem.json) places the totem and records a target aura at **−1.5 s**, but its first attack hits at **+1.0 s**. [`sim/shaman/fire_totems.go`](../sim/shaman/fire_totems.go) uses a target-bound DoT internally to schedule totem attacks; that log entry alone is not evidence that the game applies a hostile target debuff or starts combat on placement. Confirm the in-game placement and aggro timing. An agent's otherwise identical combat-only placement was about **3.9 DPS lower** on its 1,000-fight search bank, so this uncertainty can matter to Shaman results.

## 2. One Warlock can maintain several Corruption ranks on one target

**Reproduced in the simulator; whether Forever permits it is unverified.** [`corruption-multiple-ranks.json`](apl-search-repros/corruption-multiple-ranks.json) applies Corruption IDs `11672`, `11671`, `7648`, `6223`, `6222` and `172` at 1.47, 4.44, 5.91, 7.38, 8.85 and 10.32 seconds respectively, before the earlier ranks expire. Search results exploiting the overlap were much higher than the ordinary rotation and should not be treated as game-valid until checked.

[`sim/warlock/corruption.go`](../sim/warlock/corruption.go) gives each rank a separate aura label and applies the selected rank without deactivating other ranks. The same pattern is worth auditing for other ranked DoTs; it is not evidence that every such spell is wrong. By comparison, [`sim/warlock/curses.go`](../sim/warlock/curses.go) explicitly deactivates the prior active Bane.

Classic-era players describe a single caster's different ranks overwriting each other ([Blizzard forum discussion](https://us.forums.blizzard.com/en/wow/t/warlock-in-classic/246765), [older discussion](https://gamefaqs.gamespot.com/boards/534914-world-of-warcraft/43952087)); these are not authoritative evidence of Forever's current behavior. Verify in game, including different casters and targets, before changing the model.

## 3. The search runner does not surface unresolved APL spell warnings

**Reproduced at the search interface; distinguish this from a missing core warning.** In the present build, `core.IncludeAQ` is false. The registration loops therefore omit Arcane Missiles rank 8 (`25345`) and Frostbolt rank 11 (`25304`) despite their entries in the rank tables. These requests use level-60 characters, not level 65 as the agent originally assumed. Whether that gating is intended for Forever should be checked.

- [`unavailable-arcane-missiles.json`](apl-search-repros/unavailable-arcane-missiles.json) and [`unavailable-frostbolt.json`](apl-search-repros/unavailable-frostbolt.json) were scored by the search with an empty warnings list, while their rotations attempted these unavailable IDs and performed very poorly.
- [`unavailable-shadow-bolt.json`](apl-search-repros/unavailable-shadow-bolt.json) attempts Shadow Bolt rank 10 (`25307`). The level-60 table includes it, but [`registerShadowBoltSpell`](../sim/warlock/shadowbolt.go) limits registration to rank 9 when `IncludeAQ` is false. Removing the inert rank-10 action produced an identical combat log and DPS on the same seed; the search reported no warning. Verify whether that gate matches Forever's actual spell availability.
- An Arcane-Frost agent also tested an unavailable Arcane Wrath action (`1326080`) that acted as an inert APL entry in its scenario. Removing that action from the otherwise identical candidate produced exactly the same per-fight DPS values on 1,000 shared seeds. Original candidate: `forever-dps-sweep-01--arcane_frost-r02/evaluations/c64c83c4fd5feca0c38da0a5ce79bd7781623d85e8019891a17a029a941c5767/candidate.json` under the run root noted below.
- [`unavailable-arcane-wrath-rank.json`](apl-search-repros/unavailable-arcane-wrath-rank.json) uses the Arcane Wrath talent and requests rank 3 (`1326081`) at level 60. No rank-3 cast occurs; removing that action yields the identical combat log and DPS on the same seed. [`registerArcaneWrathSpell`](../sim/mage/arcane_blast.go) selects and registers only the highest eligible rank—rank 4 (`1326082`) at level 60—even though it contains tables for four ranks. Whether lower ranks should be available in Forever needs verification; currently the search interface does not surface the unresolved action.
- [`prepull-positive-time-ignored.json`](apl-search-repros/prepull-positive-time-ignored.json) contains Rapid Fire at `+1s` in the pre-pull list. [`sim/core/apl.go`](../sim/core/apl.go) deliberately ignores positive pre-pull times and records a validation warning, but the search result did not expose it. Removing only that invalid entry produced identical DPS for all 1,000 shared-seed fights. The separate in-combat Rapid Fire action still works. The raw candidate and clean alternative are retained under `forever-dps-sweep-01--marksmanship-r04/evaluations/` as `c1ca91cae1851bc46d397bfa4ea91d966cea94614657bef5b1e104f2760944dc` and `8a69cd1dfb0f1ff5931bf2ad24ad4a4f244e3a6de870bec3381ae97cd6745170`, respectively.

[`GetAPLSpell`](../sim/core/apl_helpers.go) does call `ValidationWarning` for an unknown spell. The likely integration gap is that [`tools/forever_arena/evaluate.go`](../tools/forever_arena/evaluate.go) returns `RaidSimResult`, while [`proto/api.proto`](../proto/api.proto) places `rotation_stats` and its APL warnings in the `ComputeStats` response rather than `RaidSimResult`. The search wrapper scans the simulation result for warnings, so it cannot be assumed to receive that separate channel. Check the complete UI and API paths independently before attributing the omission to core APL validation.

**Suggested check:** make unresolved or unavailable APL actions visible to the search evaluator and user, and decide whether such a candidate should be rejected. Verify optional-rank registration and the actual Forever spell availability independently.

## 4. Tiger's Fury is unavailable in the Forever Feral setup

**Observed in the configured simulator; intended Forever game behavior is unverified.** Feral agents attempted Tiger's Fury IDs `9846`, `9845`, `6793` and `5217` without an effect or a surfaced search warning. [`feral-tigers-fury.json`](apl-search-repros/feral-tigers-fury.json) contains an APL attempt to cast `9846`; its one-fight log has no such cast.

The immediate reason is not the agents' original level-map explanation: these characters are level 60, not 65. [`Druid.RegisterFeralCatSpells`](../sim/druid/druid.go) registers **Shifting Power** in Forever and calls `registerTigersFurySpell` only in the non-Forever branch. [`sim/druid/tigers_fury.go`](../sim/druid/tigers_fury.go) separately indexes exact levels 25/40/50/60, but level 60 is present in that map. Confirm whether Tiger's Fury is replaced, should coexist, or is intentionally unavailable before making a simulator change.

## 5. Pre-pull automatic cooldown action causes a nil-pointer error

**Reproduced.** [`prepull-autocooldowns-crash.json`](apl-search-repros/prepull-autocooldowns-crash.json) uses `autocastOtherCooldowns` as a pre-pull action at −10 seconds. The native evaluator exits with a simulation error, and the stack trace shows a nil `MajorCooldown` passed from [`APLActionAutocastOtherCooldowns.Execute`](../sim/core/apl_actions_casting.go) to `tryActivateHelper` in [`major_cooldown.go`](../sim/core/major_cooldown.go). This is an error returned by the native evaluator; it did not terminate the entire search service in the observed run.

**Suggested check:** either validate and reject this action in pre-pull with an actionable warning, or safely determine readiness before execution. Preserve normal in-combat automatic-cooldown behavior.

## 6. The APL documentation uses an obsolete pre-pull field

**Confirmed documentation mismatch.** [`tools/APL.md`](../tools/APL.md) shows `"doAt": "-1s"`, but the current [`proto/apl.proto`](../proto/apl.proto) represents this as `do_at_value`, serialized as `doAtValue`, for example `"doAtValue": {"const": {"val": "-1s"}}`. An agent encountered this while investigating the pre-pull action above.

## 7. Reported inert Curse of Elements needs interpretation

**Reported behavior, with an important scenario explanation.** An Affliction agent found that removing its `1311680` Curse of Elements APL action produced identical per-fight results and suspected the action was unavailable. The scenario already sets `raid.debuffs.curseOfElements: true`, and the APL casts only when that aura is absent. The observed equality therefore does not by itself establish an unavailable spell. [`registerCurseOfElementsSpell`](../sim/warlock/curses.go) maps level 60 to `1311680`; these requests use level-60 characters. If this is investigated, compare with the raid debuff disabled and inspect the APL warning channel separately.

## 8. A missing aura condition can make an action unconditional

**Reproduced.** [`missing-aura-unconditional.json`](apl-search-repros/missing-aura-unconditional.json) uses an Arcane–Frost talent build without Fingers of Frost and places an Ice Lance action guarded by `auraIsActive` for the absent aura `44543` before the filler. A logged fight cast Ice Lance 102 times. Removing only that action's condition produced the **identical combat log and DPS** on the same seed. The agent's original 1,000-fight search scored about 610 DPS; the search interface reported no warnings.

The same precaution applies to other talent-dependent auras: a Fury agent noted that Bloodthrill (`56638`) is unavailable in its baseline because that talent is not allocated; [`applyBloodthrill`](../sim/warrior/talents.go) registers the aura only when the talent has points. The Fury baseline does not reference that aura, so this observation alone is not evidence of a faulty Fury result.

[`GetAPLAura`](../sim/core/apl_helpers.go) issues a validation warning for a missing aura, while [`newValueAuraIsActive`](../sim/core/apl_values_aura.go) returns `nil` when it cannot find the aura. [`APLAction.IsReady`](../sim/core/apl_action.go) treats a `nil` condition as an unconditional action. As in finding 3, check whether the core warning is visible to each API/UI caller; the search interface did not surface it.

[`unresolved-wing-clip-condition.json`](apl-search-repros/unresolved-wing-clip-condition.json) demonstrates the same danger inside an `and`: a Survival Weave trap action checks `spellCanCast` for Wing Clip rank 1 (`2974`), which is not registered at level 60. [`newValueSpellCanCast`](../sim/core/apl_values_spell.go) returns `nil`; [`newValueAnd`](../sim/core/apl_values_operators.go) filters that operand out, leaving the other condition active. Removing only the unavailable Wing Clip operand produced the identical combat log and DPS on the same seed. Replacing it with the registered rank 3 (`14268`) changed the trap behavior and score. This affected a search finalist, though not the leading candidates. A safer handling should avoid silently dropping a specified condition when it cannot be resolved. Check the behavior of other nested operators and API warning channels too.

## 9. Hunter range cliff matches the simulator's dead-zone rule

**Reproduced; not established as a simulator bug.** A Beast Mastery agent reported that changing only the weave's return distance from 12 to 11 yards dropped its 1,000-fight score from about 1,168 to 360 DPS. [`hunter-weave-range-12.json`](apl-search-repros/hunter-weave-range-12.json) and [`hunter-weave-range-11.json`](apl-search-repros/hunter-weave-range-11.json) reproduce the two rotations on the same seed. In the logged 11-yard fight, the hunter moved to 5, then to 11 yards, and did not move again; the 12-yard rotation continued weaving.

The cliff is consistent with the explicit `MinRangedAttackDistance = 12` and `MaxMeleeAttackDistance = 5` in [`sim/core/constants.go`](../sim/core/constants.go): ranged attacks are blocked below 12, and 11 is also outside melee range. This explains the observed outcome in this APL, but the intended Forever in-game distance and hitbox convention should be verified before treating 12 as authoritative game behavior.

The same logs reveal a separate formatting defect: [`APLActionMove.Execute`](../sim/core/apl_actions_misc.go) formats a `float64` range with `%s`, so movement messages appear as `Moving to %!s(float64=11)` instead of the numeric distance. This is a logging issue, not evidence of incorrect movement.

## 10. APL cannot directly query current target distance for trap weaving

**Confirmed interface limitation; intended trap behavior needs verification.** A Beast Mastery agent could not express a condition that casts a trap only when the hunter is within five yards. The [`APLActionMove`](../proto/apl.proto) exposes a destination through `range_from_target`, but the APL value types provide no current distance or generic in-range value. The simulator itself tracks `DistanceFromTarget`.

[`hunter-trap-at-range.json`](apl-search-repros/hunter-trap-at-range.json) starts the hunter at 12 yards and places an unconditional rank-5 Immolation Trap first in the priority list. At **+0.43 s** it starts moving toward five yards and casts the trap, spending **245 mana**; the log contains no Immolation Trap aura or damage. The next cast is at **+31.41 s**. [`Immolation Trap`](../sim/hunter/immolation_trap.go) and [`Explosive Trap`](../sim/hunter/explosive_trap.go) return without applying their effects if `DistanceFromTarget > 5`, after the cast has consumed resources and cooldown. Check whether the intended game model should leave a trap on the ground or otherwise distinguish placement from triggering.

A `spellCanCast` check on a registered melee ability such as Wing Clip can be a conservative proxy for being in melee range, because its [`ExtraCastCondition`](../sim/hunter/wing_clip.go) requires range ≤ 5. In one logged fight using that proxy, an agent observed Immolation Trap cast at +4.00 s, its target aura applied at +4.20 s, and damage ticks from +7.20 to +19.20 s. The candidate is `forever-dps-sweep-01--beast_mastery-r03/evaluations/2def06c713c15bf8706427b965a00a73be9f1dc5a57096d2e35308b8424c5ea7/candidate.json` under the run root below. One trace does not establish behavior for all seeds or in game. The proxy is not a general distance predicate: spell availability, resources, readiness and other cast restrictions also affect it. A direct range or distance value would let the APL express the intended rule without those additional conditions.

## 11. Warlock pet-mana conditions can remain bound to a replaced demon

**Reproduced.** [`pet-mana-after-replacement.json`](apl-search-repros/pet-mana-after-replacement.json) configures an initial Succubus, summons an Imp before combat, then gates Life Tap on `warlockCurrentPetManaPercent < 95%`. The Imp spends 115 mana at +2.00 s, falling from 2,074.37 to 1,959.37 (below 95%), but the player casts no Life Taps during the 90-second trace. In the otherwise equivalent [`pet-mana-initial-imp.json`](apl-search-repros/pet-mana-initial-imp.json), the Imp is also the initial configured pet; the player casts Life Tap 24 times. The reproduction holds the summon, APL, and seed fixed while changing the initial configured demon.

[`newValueWarlockCurrentPetMana`](../sim/warlock/apl_values.go) and `newValueWarlockCurrentPetManaPercent` capture `warlock.ActivePet` when the APL is constructed, and their value objects keep that pet pointer. [`changeActivePet`](../sim/warlock/pet.go) later replaces `ActivePet`, but the captured pointer is not updated. If no initial pet exists, the value returns nil and can also interact with the unresolved-condition behavior in finding 8. The two requests demonstrate the percentage value; the absolute-mana implementation has the same static-pointer structure. Check that both values read the currently active demon and define the no-active-pet behavior without silently dropping a condition.

## 12. Rogue cannot enter ordinary Stealth before the pull

**Reproduced; Jordan confirmed that pre-combat Stealth should be available.** [`rogue-1784-opener.json`](apl-search-repros/rogue-1784-opener.json) attempts the normal Stealth spell (`1784`) at −1 s and places Garrote first in the priority list. The log contains no Stealth cast or aura and no Garrote. In the otherwise identical [`rogue-1856-opener.json`](apl-search-repros/rogue-1856-opener.json), using Vanish (`1856`) at −1 s activates Stealth and permits Garrote at the pull, but spends Vanish. The search does not surface a warning for the unavailable Stealth action.

[`registerStealthAura`](../sim/rogue/stealth.go) registers an aura, not an APL-castable ordinary Stealth spell. [`registerVanishSpell`](../sim/rogue/vanish.go) activates that aura; [`Rogue.Reset`](../sim/rogue/rogue.go) does not start in Stealth. Add an ordinary out-of-combat Stealth action or a suitable starting-state option, preserving Vanish for combat, then rerun the affected Rogue builds. No simulator mechanic was changed during the experiment.

## 13. An unresolved target-health condition can make Soul Fire unconditional

**Reproduced in the fixed-duration encounter.** [`target-health-percent-20.json`](apl-search-repros/target-health-percent-20.json) guards Soul Fire (`17924`) with current target health ≤ 20%. Removing only that condition, as in [`target-health-unconditional.json`](apl-search-repros/target-health-unconditional.json), produces the **identical combat log and 964.38 DPS** on the saved seed. The original 1,000-fight search candidates using 20% and 40% produced identical per-fight results; an additional 1,000-fight run with the condition removed also matched every fight. Soul Fire averaged 26.53 casts per fight rather than being confined to execute. The search returned no warnings.

[`newValueCurrentHealthPercent`](../sim/core/apl_values_resources.go) returns `nil` if the referenced unit is unavailable or has no health bar, and warns in the latter case. The encounter is duration-based and the submitted target's health stat is zero. As in finding 8, an unresolved condition can turn its action unconditional; examine both target-unit resolution and the separate APL warning channel before choosing a fix. A duration-based approximation of execute time may be appropriate for some rotations, but is not equivalent to a working target-health predicate.

## 14. Fire Nova's documented origin and implementation disagree

**Confirmed code inconsistency; intended Forever mechanics need in-game verification.** The comment above the Fire Nova tables in [`sim/shaman/fire_totems.go`](../sim/shaman/fire_totems.go) describes a caster-centred nova hitting nearby enemies. The cast condition below instead requires a live Fire totem and describes the nova as originating from that totem. The effect iterates over every encounter target without checking distance from either the caster or a totem. This can change which casts and targets contribute damage, particularly in ranged or multi-target scenarios. Confirm the actual spell's totem requirement, origin and radius before changing the model; the current single-target search alone does not establish the intended behavior.

## 15. One Mage can maintain several Frostfire Bolt ranks on one target

**Reproduced in the simulator; whether Forever permits it is unverified.** In [`frostfire-multiple-ranks.json`](apl-search-repros/frostfire-multiple-ranks.json), rank 3 (`1237313`) hits at **207.33 s** and rank 2 (`1237312`) hits at **209.83 s**. The rank-3 DoT continues ticking at **210.33 s** and **213.33 s**, while the rank-2 DoT also ticks at **212.83 s**. [`sim/mage/frostfire_bolt.go`](../sim/mage/frostfire_bolt.go) assigns a separate aura label to each rank and applies the chosen rank without deactivating the others. Verify whether the game overwrites a caster's earlier Frostfire rank before changing the simulator.

A search finalist relied on the overlap. It was excluded, and the agent's best single-rank alternative was compared with the other valid finalists on a fresh 5,000-fight bank. The original result and correction are preserved in `forever-dps-sweep-01--frostfire-r01/pre-rank-review-*` and `rank-review.json` under the run root below. The search guard conservatively quarantines candidates that cast multiple Frostfire Bolt ranks, including potentially safe non-overlapping alternation, until the game rule is verified.

## 16. Rage does not decay during a long pre-pull

**Reproduced in the simulator; exact Forever behavior needs an in-game check.** [`prepull-bloodrage-decay.json`](apl-search-repros/prepull-bloodrage-decay.json) starts with the scenario's configured 50 rage and casts Bloodrage at **−40 s**. Its aura expires at **−30 s**, after the final tick brings rage to 75. The next logged rage change is the first Sunder Armor at **0 s**, which spends 15 from **75**, so neither the initial rage nor the Bloodrage-generated rage decays in that 30-second interval. [`sim/core/rage.go`](../sim/core/rage.go) contains rage gains and spends but no apparent out-of-combat decay; [`bloodrage.go`](../sim/warrior/bloodrage.go) implements its ten-second generation. The configured 50 starting rage is a separate scenario assumption, not itself evidence of a simulator defect.

Classic references describe Bloodrage as preventing rage decay for its ten-second duration, followed by out-of-combat decay ([Classic Bloodrage aura](https://www.wowhead.com/classic/spell=29131/bloodrage), [Bloodrage history](https://warcraft.wiki.gg/wiki/Bloodrage)). Confirm Forever's actual decay rate, tick timing and combat state, including starting rage and Anger Management, before implementing it. Several search finalists used Bloodrage earlier than −10 seconds, and the retained Arms candidate uses −15 seconds; these timings may overstate available rage. Long-prepull candidates are being excluded from Warrior finalist promotion and the affected starting builds must be checked or rerun after a fix. A −10-second or later timing avoids the extra post-aura interval, but is not a substitute for verifying the game rules.

For Arms, a separate timing check compared the retained rotation with Bloodrage at −10 and −5 seconds on 1,000 fresh search fights; the −5-second version was selected. On 5,000 further independent fights it scored **909.62 DPS** against the original Arms starting build's **830.09 DPS** (paired gain **79.53 ± 0.61 DPS** standard error). The candidate and seed details are saved in `forever-dps-sweep-01--arms-r15/rage-timing-audit.json` under the run root. This does not implement rage decay or verify the real-game starting-rage assumption.

The Fury Sunder round-two and Fury 2H round-three automatic winners used early Bloodrage. For each agent, the highest-scoring conservative proposal from the existing search bank was selected before a fresh 5,000-fight confirmation; the original summaries and winners are preserved as `pre-rage-review-*`. The reviewed winners scored **969.16 DPS** (Fury Sunder, +5.22 paired DPS) and **893.02 DPS** (Fury 2H, +5.35 paired DPS). Details, exclusions and independent seeds are in each run's `rage-review.json`; the review seed is included in subsequent round histories.

## Report not substantiated: omitted current-mana source

A Survival Weave agent reported identical scores for several `currentMana` thresholds and suspected that omitting `sourceUnit` made the condition unconditional. The agent later withdrew that interpretation. In a matched 100-fight comparison, an omitted source and explicit `Self` produced identical per-fight results at both 0 and 5,000 mana; changing the threshold from 0 to 5,000 changed the results and the logged Arcane Shot casts from 32 to 19. [`GetSourceUnit`](../sim/core/apl_helpers.go) defaults an omitted source to Self. This report is retained in the raw log, but is not an established simulator issue.

## Report not substantiated: no next-melee-swing APL timer

A Retribution agent reported that `autoSwingTime` returns the swing period rather than time until the next swing, and inferred that the latter is unavailable. The first part is correct, but the APL also has `autoTimeToNext` with `autoType: MainHand`. [`APLValueAutoTimeToNext`](../sim/core/apl_values_auto_attacks.go) returns `MainhandSwingAt() - CurrentTime`, clamped to zero. Existing rotations use this value. The agent has been given the predicate for testing seal timing; the missing-interface claim is not substantiated.

## Report not substantiated: Devouring Plague available to Gnomes

A Smite agent questioned why a Gnome could cast Devouring Plague and pointed to a stale TODO in [`devouring_plague.go`](../sim/priest/devouring_plague.go). Registration is gated by its caller in [`Priest.Initialize`](../sim/priest/priest.go): Classic restricts it to Undead, while Forever deliberately registers it for every race. The existing [Priest beta-data audit](beta-pass/priest.md) records the Forever SkillLineAbility race mask as −1, compared with Undead-only mask 16 in Era. This report does not establish a Forever simulator bug; the TODO should not be read in isolation.

## Reproduction

From the simulator repository, build the native evaluator and run any saved request:

```sh
go build -buildvcs=false -tags with_db -o /tmp/forever-arena-sim ./tools/forever_arena
/tmp/forever-arena-sim \
  -request docs/apl-search-repros/prepull-immolate.json \
  -output /tmp/apl-repro-result.json -threads 1
python3 -c 'import json; print(json.load(open("/tmp/apl-repro-result.json"))["logs"])'
```

Twenty-five requests completed against the pinned binary. The remaining request, `prepull-autocooldowns-crash.json`, reproducibly returns the nil-pointer error described above. For successful runs, the JSON result's `logs` field contains timestamps and spell IDs; the unknown-spell cases additionally require checking the APL/`ComputeStats` warning channel, not just the returned `RaidSimResult`.

The original run root is `~/.cache/forever-sim-audit/rl/feral-pilot/`. Search and holdout seeds, complete request/result files, and candidate identities are retained there while the cache remains available.

## Experiment-side precautions

No simulator mechanic has been patched for this experiment. Current and future agents are asked not to rely on early hostile effects other than Hunter's Mark, overlapping ranks of the same DoT, unavailable Mage ranks (including lower Arcane Wrath ranks), unresolved aura or target-health conditions, or pre-pull automatic cooldowns, and to report further issues. Hunter's Mark may be applied before combat. When changing talents, check that any aura referenced by the rotation is still registered: an unresolved condition can leave its action unconditional. The search wrapper separately quarantines mixed Corruption and Frostfire Bolt ranks, the observed early Bane of Doom, Corruption and Immolate pre-pulls, Bloodrage earlier than −10 seconds, `spellCanCast` guards on unavailable Wing Clip ranks 1 and 2, and target-health conditions from automatic promotion in newly launched searches. For newly launched searches, a logged first smoke-test fight also flags observed nonzero damage before time zero. These checks are intentionally conservative, can flag harmless rank alternation, and do **not** cover every possible hostile pre-pull or rank-stacking issue; one sampled trace cannot prove that every fight is clean.

A fresh valid pre-cast that lands at or after time zero remains allowed. Do not interpret the quarantine as a verified statement about Forever's in-game behavior.

The experiment's original Smite-theme guard accepted a single incidental Smite cast. One Shadow/Mind Flay finalist passed that guard but was excluded from the Smite line; the agent's sustained-Smite alternative was confirmed on a fresh independent bank. The original result and the correction are preserved in the run's `pre-theme-review-*` files and `theme-review.json`. The search guard now requires sustained Smite, not just an opener, and rejects the observed pattern of replacing it with Mind Flay for an entire half-fight. This conservative check does not prove that every alternative APL preserves its theme. A Mutilate search also produced a Backstab-only finalist; it was excluded from theme-compliant comparisons, while the confirmed round winner retained sustained Mutilate. The Mutilate guard now requires its talent and sustained use as the primary builder. These were experiment-validation problems, not reported simulator mechanic bugs.
