# Questions review

The original review covered all 137 entries at `35cab052456a320358e998c12416687d69cae142` and added three specific gaps. The correctness follow-up reviews code at `4235b6feda425a82e76f1af741fa61e87dcabcc8`, the September 30 Hunter/Druid deep dive and 20 recent evidence findings. Captured client tables extend to `1.60.1.70009`.

The [canonical questions register](uncertainties.md) is the tracking list. Each entry keeps its stable ID, answer, confidence, remaining scope, sources and concrete test prerequisites. The [structured review](../artifacts/question_review/review.json) records the adjudications. The [follow-up coverage](../artifacts/question_review/correctness_followup.json) accounts for all 144 article feature points and all 20 recent findings without publishing private chat or identifying metadata.

## Accounting

| Status | Original 137 | Initial review, 140 | Correctness follow-up, 145 |
|---|---:|---:|---:|
| Needs game/source evidence | 84 | 63 | 69 |
| Implementation work | 27 | 27 | 13 |
| Model/scenario choices | 5 | 8 | 8 |
| Accepted working answers | 0 | 5 | 6 |
| Resolved | 20 | 24 | 36 |
| Outside current scenarios | 1 | 13 | 13 |

Added records distinguish Hunter talent decoding, Hunter’s Mark rank values, Savage Strikes, Faerie Fire’s announced rules and Flametongue’s missing numeric model. A fixed implementation does not close an unrelated coefficient or script question.

Pending queues include 39 entries with level-20 checks, 17 with level-30 checks, 42 with later-access checks and 33 with source/code work. These overlap: a low-rank test may not settle its level-60 counterpart. The level stages describe prerequisites, not an independently verified statement that every spell, item or server build is currently accessible.

## Correctness follow-up

The [implementation ledger](correctness_workplan.md) records 45 correction groups with focused checks. They include:

- **Combat and resources:** competing minor armor reductions; instant-versus-hardcast swing handling; free-cast Rage refunds; caster-side pushback; finite delivered-damage absorbs and generated/absorbed/unused shield accounting.
- **Druid/Hunter:** full Moonfire scaling, exact Thick Hide coefficients, Furor carryover, form weapon DPS/stones/enchant reconstruction, form-safe cast eligibility, Barkskin registration, trap timers, Intimidation crit scope, first-dodge Mongoose, owned Mark/Expose Prey, Carrion Screech and the sourced Savage Strikes crit family.
- **Paladin/Warlock/support:** Holy Strike/Divine Precision scope, Vindication, four-target Consecrated Ground, enemy Light’s Vigil, Wrack ownership/family, full Improved Imp damage, Blood Pact lifecycle, PI targeting/non-stacking, Crusader/Lifestealing health returns and Defender’s Grip.
- **Data and legality:** Hunter’s Careful Aim/Lethal Attacks decode swap, legal legacy presets, current metadata fixtures and actual menu/landing contracts.
- **Rogue:** [Instant Poison’s sourced 0.5% AP term](https://github.com/ClassicWoWCommunity/forever-bugs/issues/100#issuecomment-5804701503). Deadly’s tick/stack placement remains open; Classic receives no new AP term.

The [correctness review](correctness_release.md) records the regenerated 804 scenarios, 201 exact web defaults and 68 matched tank checks. Old requests/results are archived rather than relabeled as corrected. These are fixed-profile correctness comparisons, not optimized-build claims; publication also requires native/WASM/browser and live deployment checks.

## Partial answers that remain material

| Topic | Supported answer | Remaining question |
|---|---|---|
| Hunter pets | Screech’s active ranks/AP effect and Carrion assignment are corrected; basic Cat attacks remain. | Other unique abilities, new families, active metadata, coefficients and family baselines. See HUN-007/HUN-009. |
| Summon Hawk | Its current single refreshing DoT is an approximation. | Actual concurrent guardians, stats, damage and proc behavior. HUN-004. |
| Savage Strikes | Captured data says family-specific crit; registered scope is repaired. | The article announces melee damage instead. HUN-012 preserves the version conflict. |
| Faerie Fire | Caster and form records are distinct. | The announced free/six-second rule conflicts with captured base spell data. DRU-014. |
| Swipe / reactive auras | AP scaling for Swipe and provider-SP scaling for Thorns/Ret Aura are supported intentions. | Numeric server coefficients and whether the announced Swipe fix is live. DRU-013/PAL-010. |
| Light’s Vigil | Enemy damage, one owned mark, refund and linked cooldown controls are implemented. | Party healing and the exact linked hit/proc sequence. PAL-013. |
| Absorbs | Finite all-school pools and actual absorption are modeled. | Shielding threat and overlapping-shield priority; application order is a convention. SCEN-005. |
| Poisons | Instant’s direct AP term is sourced and implemented. | Deadly placement, separate finisher/bleed questions and charge accounting. ROG-002/ROG-003. |
| Shaman instant casts | Current no-reset evidence supports instant Bolt preserving swings. | Later-access combat traces; old reset-based research remains historical. SHA-002. |
| Flametongue | Official notes prohibit stacking with Windfury Totem or Flametongue Weapon and remove lingering buffs. | A numeric damage/speed/SP/proc rule. A dummy field is not a damage formula. SHA-006. |

The new Feral talents and row moves explicitly announced for the following build are not silently merged into current strings. Roots, utility/healer spellbooks, movement and kill-triggered features have separate coverage limits; their absence is not mistaken for an unknown coefficient in an already implemented spell.

## Historical evidence

The [seed-range inventory](../artifacts/question_review/seed_ranges.json) identifies overlapping historical confirmations. Distinct starting seeds do not imply disjoint iteration ranges. New validation guards compare complete ranges; retained historical means remain usable descriptions, but independence-based confidence labels require qualification. Common-seed baseline/sensitivity replays are deliberately paired and are not independent confirmations of game mechanics.

The modeled catalog has no ordinary set procs, item guardians or unresolved real-item effects. Those remain a catalog backlog, rather than blockers for this hypothetical-equipment matrix. Weapon enchants are active effects and have been reviewed separately.

Useful source-backed answers include [baseline Divine Spirit](https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap), [pet external-buff exclusions](https://github.com/ClassicWoWCommunity/forever-bugs/issues/27), [ordinary off-hand rage evidence](https://ppach-warriorcompendium.share.connect.posit.cloud/rage.html) and [the source-specific MP5 bug report](https://github.com/ClassicWoWCommunity/forever-bugs/issues/54). None validates every related server script or level-60 interaction.

## Test priorities

- **Level 20:** separate flat MP5 from Spirit/other returns; weapon-racial melee/ranged crit controls; low-rank Energy/avoidance behavior; accessible pet buff/ability controls. Record builds and isolate one variable.
- **Level 30:** one-stack/five-stack Deadly AP controls; rank-one Flametongue at level 28; accessible later pet ranks. Grace of Air is not available for the full support comparison at this stage.
- **Later access:** full-rank poison/form/pet effects, Maelstrom timing, high-rank Mark, guardian behavior, tank/server coefficients and full Flametongue + Grace versus Windfury comparisons.
- **Sources/code:** shipped client updates for announced redesigns, missing active spell/family joins, historical confidence qualification and known implementation tasks. Do not wait for an in-game experiment to repair a deterministic code defect.

The searchable register provides each test’s actual prerequisites and observations. A report, another simulator or a 100%-chance client wrapper is not automatically a verified effective server rule.
