# Questions review

The starting register has 137 entries at
`35cab052456a320358e998c12416687d69cae142`: 84 evidence questions,
27 implementation tasks, five model choices, one out-of-scope item and
20 resolved records. This review has not changed combat calculations or
the published benchmark.

All **137 of 137 entries** now have individual adjudications. Newly identified
Swipe AP-script, Improved Imp damage-scope and weapon-racial crit-scope gaps are recorded separately
rather than counted as answers to existing questions.

The [structured review](../artifacts/question_review/review.json) records
individual answers, confidence, remaining scope, implementation consequences
and test prerequisites. The [canonical register](uncertainties.md) now combines
those findings with the original questions and stable IDs.

## Accounting

| Status | Before | Original entries now | With three additions |
|---|---:|---:|---:|
| Needs game/source evidence | 84 | 61 | 63 |
| Implementation work | 27 | 26 | 27 |
| Model/scenario choices | 5 | 8 | 8 |
| Accepted working answers | 0 | 5 | 5 |
| Resolved | 20 | 24 | 24 |
| Outside current scenarios | 1 | 13 | 13 |

The register separates **21 source-supported implementation follow-ups** from
remaining unknown behavior. None has been applied by this publication.
Pending queues contain 39 questions with current-beta checks, 16 with level-30
checks, 41 needing later access and 39 with source/code work. These overlap:
a low-rank test does not necessarily settle its level-60 counterpart.
Accepted-answer corroboration is optional, not included in those pending queues.

## Findings

- **Older confirmation runs were not independent.** The
  [seed-range inventory](../artifacts/question_review/seed_ranges.json)
  identifies overlap in gear, build and tank confirmations. For example,
  the current gear-selection confirmation has 5,001 distinct iteration seeds,
  not 10,000. The latest Sunder ranges are disjoint. This does not make every
  reported gain false, but independence-based confidence claims need correction.
- **Some catalog unknowns do not affect the current scenario.** The projected
  gear has no ordinary set procs, item guardians or unresolved real-item effects.
  Those stay in the catalog backlog rather than the blocking combat-test list.
  Weapon enchants are different: all 17 tanks equip Crusader, whose sourced
  self-heal is absent from its current handler. Its effective impact has not
  been simulated in this pass.
- **Old talent-source conflicts are mostly stale history.** The September 17
  audit contains a later reversal, and current selectable trees already use
  the corrected rank values and removals. The retained video-extraction JSON
  is not the current UI tree. A small inherited-fixture probe also distinguished
  obsolete test expectations from a genuinely stale Paladin tree-size constant.
- **Several open features need code, not beta access.** These include generic
  absorb pools, targeting another player with owned Power Infusion, item-picker
  shared-stat valuation, raw haste-rating representation and a refreshed
  spell-coverage census. Defender's Grip's ordinary outside-city effect also
  has enough client data for implementation.
- **Weapon-racial crit has a newly isolated scope conflict.** The Hunter wiki
  excludes ranged crit, whereas the current engine's shared physical-crit path
  includes it. A matching/nonmatching melee-weapon control can check this at
  level 20. This is not a proposal to change generic equipment crit.

- **Divine Spirit availability is answered.**
  [Blizzard explicitly confirms it as baseline](https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap).
  A future trainer screenshot is not needed to justify the intended buff.
- **Two supposedly inherited cooldowns are actually sourced.** Hack and
  Slash has a 200-ms recovery in both its parent and child records. Windfury
  Weapon's actual enchant points to a dummy aura with a 1.5-second recovery.
  Windfury's attack-resolution script remains open; its cooldown does not.
- **Wrack's modifier is too broad.** The current client family mask selects
  Corruption and Agony, not every Shadow DoT. This is an implementation
  correction, not a question that needs a future raid test.
- **Improved Imp exposes two separate discrepancies.** Its documented
  Firebolt multiplier should include the SP contribution; the engine applies
  it only to base damage. It no longer improves Blood Pact, but the engine and
  184 DPS requests still select the old improved Stamina bonus. The 17 tank
  requests do not select Blood Pact. The undocumented negative dummy remains
  a separate level-20 test.
- **Poison charges are a lower-priority accounting gap.** Existing five-minute
  Rogue results average at most 92.29 Instant attempts and 96.05 Deadly
  applications against 175/180 charges. This is useful headroom, not a
  per-iteration guarantee or permission to ignore arbitrary long fights.
- **Gear MP5 is a reported beta bug, not a universal mana rule.**
  [The report](https://github.com/ClassicWoWCommunity/forever-bugs/issues/54)
  includes gear controls and distinguishes ordinary Wisdom/Mageblood behavior.
  The optional uniform fivefold mode is a sensitivity scenario, not an exact
  reproduction of that source-specific bug.
- **Poison AP scaling has a published answer.**
  [The developer response](https://github.com/ClassicWoWCommunity/forever-bugs/issues/100)
  gives Instant and Deadly ratios absent from the current engine. Instant's
  direct-hit term is source-ready; Deadly's tick/stack placement needs
  reconciliation. This is narrower than treating all Rogue scaling as unknown.
- **Ordinary off-hand rage now has Forever evidence.**
  [New public logs and analysis](https://ppach-warriorcompendium.share.connect.posit.cloud/rage.html)
  support the half-rate model. One independently inspected sample reproduces
  approximately 1.73 versus 3.455 rage per weapon-second. This does not validate
  level-60, rage-talent or proc-extra-swing behavior.
- **Thick Hide's intended Defense coefficient is in the active rank curves.**
  It is 0.67/1.33/2.00, not a rankless scalar or 0.67 multiplied by every rank.
  The engine's small rounding discrepancy is implementation work. A
  [separate live-beta bug](https://github.com/ClassicWoWCommunity/forever-bugs/issues/151)
  does not replace the intended formula.
- **Generic pet buff exclusion is answered for both pet classes.**
  [The explicit response](https://github.com/ClassicWoWCommunity/forever-bugs/issues/27)
  identifies this as intentional. Demon inheritance coefficients and AI cadence
  are separate questions.
- **Consecrated Ground was misclassified.** It grants extra Holy damage to
  the first four covered enemies, not mitigation. Its current owner-wide
  approximation needs target scoping for larger encounters; that cap does not
  distinguish the stationary one-/three-target workloads.
- **Retribution Aura still has a sourced base-value discrepancy.**
  The captured rank-5 effect is 30; the engine uses 20. The
  [developer response](https://github.com/ClassicWoWCommunity/forever-bugs/issues/117#issuecomment-5822983107)
  also confirms intended SP scaling for Ret Aura and Thorns, but does not
  provide their numeric coefficients.
- **Swipe has a new, specific follow-up.**
  [The server-side AP script was accidentally bypassed](https://github.com/ClassicWoWCommunity/forever-bugs/issues/125#issuecomment-5841232704).
  An internal fix was announced for a later build. Its coefficient is not in
  that response, so importing a remembered SoD number would not resolve it.
- **Judgement of Wisdom has conflicting source layers.** Wowhead displays a
  50% chance, while the captured client parent has 100% and invokes a dummy
  intermediate. Neither establishes the effective server-script probability.
  The 59-mana max-rank return itself is sourced.
- **Some tests were scheduled too early.** Warrior Flurry and Maelstrom
  Weapon first become available at level 35, and their fifth rank at 39.
  They are not level-30 beta experiments. Conversely, Inner Focus can be
  tested now at level 20.

## Implementation follow-up

The [class evidence ledger](../artifacts/question_review/class_evidence.json)
and [scenario/data ledger](../artifacts/question_review/scenario_data_evidence.json)
record selected raw fields, table provenance and existing-result calculations.
The structured review separates source-ready values from rules that still need
an experiment. Examples of source/code work include the Ret Aura base amount,
Thick Hide rank rounding, Improved Moonfire's damage modifier, target-scoped
Consecrated Ground and missing supported spell families such as Seal of Fury
and Light's Vigil. These have **not** been applied to the published simulation
in this research pass.

The existing register also contains historical fixes. A resolved queue or
seal-maintenance bug should remain resolved; a separate uncertainty about rage
or threat does not invalidate the reproduction of that bug. Conversely, older
nearby random seeds must not be presented as independent confirmations.

## Reading a disposition

An evidence-backed answer is not necessarily already implemented. A credible
working answer need not be a measured level-60 server formula. A known beta
bug is not automatically the intended launch rule, and a closed issue is not
proof that its fix has shipped.

“Not material” is limited to a stated scenario or unsupported capability.
Damage, threat, resource supply and survival all count when assessing tanks.
Missing damage abilities and proc enchants are not dismissed merely because
the current rotation does not select them.

Current-beta tests assume the level-20 cap. “Launch” means beyond the announced
level-30 beta or requiring max-level content; that label must move if access
expands. Each remaining test must identify the required ability, talent rank,
item or target rather than simply saying “test in game.”

## Evidence limits

- The captured client is a versioned collection of tables, not one complete
  build-70009 export. Each table's actual build and hash are retained in the
  source ledger. Client dummy effects do not expose every server script.
- Several Hunter issue records contain only Discord links, without publicly
  readable test results. Their existence is not independent confirmation.
- A fresh fetch of missing objects from `wowsims/forever` returned
  “Repository not found.” Previously captured, pinned material was usable;
  this review does not claim a fresh audit of unavailable upstream objects.
- Low-level measurements, announced internal fixes and current intended
  mechanics have different scopes. A closed issue does not prove a fix has
  shipped, and a level-20 result does not automatically settle a later rank.
- The older overlapping confirmation runs remain intact. Their independence
  claims need correction; this publication does not replace their results or
  imply that every retained build gain disappears.
