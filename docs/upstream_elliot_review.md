# Recent ElliotWood/Forever review

Current questions and corrections to earlier interpretations are in the
[canonical register](uncertainties.md). In particular, **DONE-010** corrects
the Seal of Command comparison below: this engine's final multiplier already
puts personal SP inside the 70% term. Moving fields is not an additional
damage correction; Seal of Righteousness remains a separate unresolved formula.

## Scope and release status

Compared [ElliotWood/Forever at `eccf6aa87d`](https://github.com/ElliotWood/Forever/tree/eccf6aa87ddc030ed6d5a0c4842121027ae89bb9) with this simulator's `367ac97df8` release.

The window is **September 24, 03:33:34 UTC through September 26, 03:33:34 UTC, 2026**: 161 reachable commits, including 83 first-parent commits. These include changelogs, generated data and imported WoWSims work; they are not 161 independent mechanics fixes. Commit metadata was screened across the window, with the relevant implementation diffs and corresponding local code examined below. This is not a fresh audit of every file in the other simulator.

The preceding [weekly review](weekly_review.md) and [Flurry review](flurry_review.md) remain the baseline for this follow-up. Their results are preserved in [the historical archive](../artifacts/history/367ac97df8/). The tables below describe findings against that baseline; the implementation and refreshed results are summarized here.

### Implementation status

The supported corrections are implemented at [`3b2d47961`](https://github.com/gunba/wow-forever-sim/commit/3b2d47961). All **184 fixed profiles** and their Tier-off/+10%/+50% scenarios were replayed: **736 runs**, 5,000 iterations each, seed 20291951, zero warnings. Requests and baseline players match the preceding release exactly. Gear selections, talents and APLs were not reoptimized.

Implemented:

- Spirit-based Life Tap, with level-capped rank bases, separate health payment and mana return, no damage/proc scaling, and rank-correct pet mana metrics. Improved Life Tap follows the current parent spell's health-and-mana formula; Tier 1 multiplies mana only. Non-tanking DPS retains its existing external-healing abstraction.
- Demonic Knowledge's explicit pet bonus, and Focused Fire on both the owner and active pet, with dismissal/resummon handling. Existing general pet inheritance is preserved, not introduced by this change.
- Client-mask Hunter cost/crit modifiers. Verification also found missing **Efficiency discounts for Hawk and Volley**. The fixed Marksmanship profile now reports about **5.5 seconds** of mana limitation in the 200-iteration regression, rather than receiving an unsupported Sniper Shot discount; its APL is unchanged.
- Stronger legal Trueshot rank, Thorns' 22 base damage, whole-school Moonfury, flat-threat Subtlety, Improved Healing on Penance and the separate Inner Focus crit mask.
- All five Cone of Cold and four Frost Nova damage ranks, shared rank cooldowns, relevant talents, Cold Snap and 10-yard range checks. Root/slow control and cone geometry are not simulated. At level 60, Cone of Cold rank 5 has a **343** midpoint, not the tooltip's level-63-capped **347.5**. Frost Nova rank 4 caps at a **75.5** midpoint. Variance spreads the unscaled base in the captured data/tooltip, rather than multiplying the level-adjusted midpoint.
- Owned Mana Tide's three legal ranks: **10/30/60 mana**, restoring **88/197/290 per pulse**, respectively. Four pulses at three-second intervals, party scope, one-second GCD, five-minute shared rank cooldown and Totemic Focus cost reduction. Dropping another owned water totem cancels remaining Tide pulses; Tide cancels owned Healing Stream. It is not added as a free extra external cooldown. The manually selected Mana Spring/Wisdom buff remains an external-provider approximation, not a second owned totem; automation of the owned Mana Spring buff itself remains a pre-existing coverage gap.
- Incomplete APL casts return no action instead of panicking. The arbitrary aura-count ceiling is removed; duplicate-registration and finalized-environment guards remain.

The qualified Seal/shield/Windfury/form differences remain unverified server behavior and are not imported. Spell metadata is updated alongside the engine; no gear search or talent/APL campaign accompanies these corrections.

Verification uses the captured `SpellEffect_1.60.1.70009.csv`
(SHA-256 `fdecddfd93a3c27b48d7bbbebc8cc733a502fa8d58de623646550f9a5cffeb00`),
the pinned upstream store/trait data and current Forever spell pages.
Regression coverage is in `tools/forever_bench/upstream_followup_test.go`,
`tools/forever_bench/mana_tide_test.go`, `sim/warlock/lifetap_test.go`
and the core APL/aura tests. The source scanner also follows typed rank-slice
parameters, so the new Mage ranks cannot disappear from the coverage check.

The legacy `TestP1Hunter` and `TestP1Mage` suites reference removed item IDs
272491 and 272457 and stop before exercising these changes. Current-profile
regressions, not those obsolete gear fixtures, cover the affected agents.
The runnable Warrior/Protection snapshot suites were also refreshed for
the 25-RAP Trueshot increase and stronger Thorns. A test-only control
restoring just those two old values reproduced all three old suites;
the engine and fixed-profile benchmark inputs were not reverted.

### Fixed-profile results

Equal-race-weight mean DPS across each build's supported races:

| Build | Previous | Corrected | Change |
|---|---:|---:|---:|
| Balance | 663.46 | 705.66 | +6.36% |
| Beast Mastery | 1,000.25 | 1,011.57 | +1.13% |
| Marksmanship | 825.55 | 826.07 | +0.06% |
| Survival | 989.12 | 954.25 | −3.53% |
| Pet/Melee Hunter | 1,014.88 | 978.34 | −3.60% |
| Demonology | 896.73 | 887.39 | −1.04% |
| Affliction | 831.60 | 804.82 | −3.22% |
| DS/Ruin | 799.42 | 766.62 | −4.10% |
| Destruction | 825.25 | 801.48 | −2.88% |

These are combined engine-change effects, not isolated gains from individual
fixes. Other build means changed by less than 0.01 DPS. The 24 paired tank
requests and 17 race/tank smoke checks also passed; the tank comparison
numbers are unchanged. Full requests/results remain in the
[current benchmark](../artifacts/modelled_gear/forever_dps_5min.json) and its
paired sensitivities.

## Source-supported gaps

Most of these appear in the upstream [talent confidence pass, `8dc1353e26`](https://github.com/ElliotWood/Forever/commit/8dc1353e26e5b65cf11293cfd1399668020cac34). The findings were compared with client effects, descriptions and our implementations rather than accepted solely because that commit calls them fixes.

| Area | Finding in this simulator | Proposed correction and evidence |
|---|---|---|
| **Warlock: Life Tap** | `sim/warlock/lifetap.go` still calculates the return from a self-targeted damage result with a spell-power coefficient. | Use the current Spirit-based conversion, not spell power or outgoing damage modifiers. [Life Tap 11689](https://www.wowhead.com/forever/spell=11689/life-tap) explicitly names Spirit. Separate the mana formula from the health-cost/healing policy; see the important upstream mismatch below. |
| **Warlock: Demonic Knowledge** | `sim/warlock/talents.go` gave the owner the level-based spell damage, but not the active demon. | Apply the explicit pet bonus as well. [412732](https://www.wowhead.com/forever/spell=412732/demonic-knowledge) names both owner and demon. The existing engine already has general owner-stat inheritance; this explicit talent bonus is additional and must not count inherited stats twice. |
| **Hunter: Focused Fire** | `sim/hunter/talents.go` applies the active-pet damage bonus only to the Hunter. | Apply the 1/2% bonus to the pet too. [1223755](https://www.wowhead.com/forever/spell=1223755/focused-fire) explicitly covers both. |
| **Hunter: Predator's Edge** | The crit-damage modifier uses every melee-defense-type spell, which is broader than the client's affected-spell mask. | Limit that part to the supported Hunter abilities, excluding ordinary auto-attacks and the modeled Hawk strike. Keep the separate off-hand damage effect separate. SpellEffect 1340259 is a masked spell modifier, not a general melee-crit modifier. |
| **Hunter: Efficiency / Resourcefulness** | Broad shot/melee filters discount Sniper Shot and Strider Kick beyond the client masks. | Efficiency's mask excludes both; Resourcefulness's cost mask excludes Strider Kick. Compare effects 696992 and 1134467 with the abilities' family flags. The resource proc is a separate effect and must not be changed along with the cost filter. |
| **Hunter: Trueshot Aura** | `sim/core/buffs.go` supplies 50 ranged AP from rank 5. | The learnable level-50 rank, [20905](https://www.wowhead.com/forever/spell=20905/trueshot-aura), gives **75 ranged AP**; level-60 rank 20906 gives 50. Use the stronger legal rank rather than automatically choosing the highest rank. Upstream [a35f6b19fb](https://github.com/ElliotWood/Forever/commit/a35f6b19fb7fbae6e31fe9029586d31b92aedf26) handles this deliberately. |
| **Druid: Moonfury** | `applyMoonfury` adds to base damage on a limited spell list; `CalcDamage` adds spell power afterward. | Effect 694153 is aura 79 with school mask 72: Arcane and Nature damage. A school damage multiplier covers the spell-power component and the proper school scope. This provides concrete data for the previously open base-only question, T26. |
| **Druid / external buff: Thorns** | `ThornsAura` still uses 18 damage. | Effect 687684 for [rank-6 spell 9910](https://www.wowhead.com/forever/spell=9910/thorns) is **22**, with no per-level or spell-power coefficient. Correct the base value without inventing scaling. This primarily affects characters being attacked. |
| **Druid: Subtlety threat** | It scales `spell.ThreatMultiplier`, but `ThreatFromDamage` adds `FlatThreatBonus` after that multiplication. | The Arcane/Nature threat reduction must also reach applicable flat threat, such as Faerie Fire. This is a threat correction, not a general DPS gain. |
| **Priest: Improved Healing** | The Penance registration does not apply this talent's cost reduction. | Apply the sourced **5/10/15%** discount to Penance. [14912](https://www.wowhead.com/forever/spell=14912/improved-healing); upstream [00f3ec38a5](https://github.com/ElliotWood/Forever/commit/00f3ec38a542ee1c033d8fd653e025ce9c2b1b76). |
| **Priest: Inner Focus** | Every mana-costed Priest spell receives the crit bonus along with the cost reduction. | Use the distinct client masks for the two effects. The crit mask excludes Mind Flay, Shadow Word: Death and Starshards; their treatment cannot be inferred from eligibility for the free cast. Effects 692356 and 692357 have different masks. |
| **Mage: Cone of Cold / Frost Nova** | These damage abilities and their talent effects are not registered in our Mage implementation. | Complete their sourced rank registrations and the corresponding talents: Improved Cone of Cold **12/23/35%**, Improved Frost Nova **−2/4 seconds**. [11190](https://www.wowhead.com/forever/spell=11190/improved-cone-of-cold) and upstream [a5b3fe56d4](https://github.com/ElliotWood/Forever/commit/a5b3fe56d41e924d74e25fd5279a002f205c4a61). This is spellbook/AoE coverage, not evidence of a five-minute single-target gain. |
| **Shaman: owned Mana Tide** | The talent's own registration remains commented out, although the external party effect exists. | Add an actual owned cast, including its cost, cooldown, GCD and water-totem interaction. Upstream [ddc37179dc](https://github.com/ElliotWood/Forever/commit/ddc37179dcd04cb0e9814e129cc32d5e9a73df15) adds an external-provider count instead; that shortcut does not model those costs or interruption of Mana Spring. Our existing external helper already supplies four direct mana pulses and can remain distinct from MP5. |
| **APL editor/import: missing spell ID** | `GetAPLSpell(nil)` calls `ProtoToActionID(nil)`, which panics before the cast-action constructor can reject the missing spell. | Port the early missing-ID rejection from [f8ecd9b572](https://github.com/ElliotWood/Forever/commit/f8ecd9b572857237650127a9a9cb5570194f8441), with a reproduction using an empty cast action. This is a concrete usability bug; it does not change valid benchmark rotations. |

The most consequential first batch is Life Tap, Hunter talent scope, pet bonuses and Moonfury. Those alter damage or resources in current profiles. Trueshot and Priest corrections follow. Spellbook completion and the tank/threat-only changes have a narrower or different scope.

### Life Tap: do not copy the upstream implementation unchanged

Upstream's code uses Spirit for mana but a flat base health payment. Its current captured description, and the current Wowhead description, put the Spirit/talent expression on **both** health and mana. Thus the upstream change correctly identifies our spell-power mistake but is not itself a complete authority for the replacement.

The rank-6 effect is 420 with 1 point per level, spell level 56 and cap 66: **424 at character level 60**, before Spirit and the relevant talent. Wowhead's displayed 430 must not simply be copied into a level-60 constant. Tier bonuses that specifically increase mana returned must remain separate from health cost. The existing treatment of external healing in DPS scenarios is also a separate modeling choice.

### Mask cross-check

The relevant raw `SpellEffect_1.60.1.70009.csv` capture has SHA-256:

`fdecddfd93a3c27b48d7bbbebc8cc733a502fa8d58de623646550f9a5cffeb00`

Its Predator's Edge, Efficiency, Resourcefulness, Inner Focus, Moonfury and Thorns rows agree with the audited upstream store. Family-flag comparisons use that store's ability records. Examples:

- Sniper Shot: Hunter family, word 3 `0x00800000`; absent from Efficiency's mask.
- Strider Kick: Hunter family, word 2 `0x00200000`; absent from both cost-reduction masks.
- Inner Focus's cost and crit masks are different; a cost reduction is not evidence of a crit bonus.
- Moonfury's effect is a school-percent aura, with Arcane/Nature mask `72`, rather than a base-damage-only spell modifier.

## Useful differences that still need qualification

| Change | Comparison and disposition |
|---|---|
| **Seal of Command's spell-power placement** — [3144902f6e](https://github.com/ElliotWood/Forever/commit/3144902f6e495d8c0f32702c80dc4f5f5ebb7414), [af747a5c29](https://github.com/ElliotWood/Forever/commit/af747a5c2924dadadecac1a6cc1d4245251d337f) | Their formula puts personal spell power inside the weapon-percent term, while target Holy-damage bonuses remain outside it. The commit reports a beta measurement of about 5.5 damage from 27 spell power, consistent with `0.29 × 0.70`. Our `soc.go` differs. This deserves a focused port/measurement review, including Improved Seals; the underlying measurement is not included with the diff. It does **not** settle Seal of Righteousness. |
| **Seal of Righteousness** — current upstream Paladin implementation | Upstream adds its coefficient once and does not use our extra two-hand `1.1` factor. Our `sor.go` deliberately carries an older double-add model and a SoD-derived factor. This strengthens the priority of **T57**, but a different simulator implementation is not proof of the server formula. Its rank data, target Holy bonuses and one-/two-hand behavior must be resolved together. |
| **Shield Slam / Shield Bash as off-hand attacks** — [b6ce30c972](https://github.com/ElliotWood/Forever/commit/b6ce30c972472c37efb67787663614ff07b1ee2c) | Ours treats them as main-hand specials. Upstream moves them off-hand, preventing main-hand Windfury/weapon procs while avoiding inappropriate off-hand weapon modifiers. The commit has engine tests and simulated comparisons, but no supplied in-game proc evidence. This is a useful tank-proc lead, not something to reproduce by changing one mask without checking the resulting hit/modifier behavior. |
| **Windfury charge handling** — [5d5b7018f9](https://github.com/ElliotWood/Forever/commit/5d5b7018f9), [c4c4ef243c](https://github.com/ElliotWood/Forever/commit/c4c4ef243cb91d400b6e3a2ab25938012c71ad63) | Their generated driver accidentally used the 20% proc chance for charge spending too. Ours does **not** have that bug: qualifying attacks spend charges without another 20% roll, and landed main-hand specials can trigger the buff. Remaining differences include avoided-swing consumption, AP-buff duration and extra-attack handling. Their temporary-enchant wording is not a reason to undo our separate Totem buff or air-totem exclusivity. |
| **Hawk inheritance, Ice Lance coefficient, Maelstrom details** — [fa0cd37596](https://github.com/ElliotWood/Forever/commit/fa0cd37596b780529259b1672de5a4501422a1f3) | This is largely an evidence/assumption update, not a set of verified numerical replacements. Hawk aura 1293586 is useful evidence about haste/damage/crit, but does not alone establish the complete AP-inheritance rule. Ice Lance's `0.143` remains an assumption there too. The Maelstrom/Windfury effect interpretation still needs server confirmation. |
| **Form weapons retaining an imbue's flat damage** — [443904a6af](https://github.com/ElliotWood/Forever/commit/443904a6af) | This repairs how their weapon-based imbue representation survives a form-weapon replacement. Our representation differs. It is not, by itself, in-game evidence that every flat weapon enchant should improve Cat/Bear attacks; that eligibility must be distinguished from preserving an already-approved bonus. |
| **Removing the aura-count cap** — [a1b51cd255](https://github.com/ElliotWood/Forever/commit/a1b51cd255) | Their former 200-aura cap can reject a large caster raid. Ours already raises that guard to 1,000, so the specific 200-aura failure is addressed, but the implementations are not identical. Removing the arbitrary cap entirely remains a useful robustness improvement. |

## Changes already represented here

- **Penance's initial bolt, then bolts at one and two seconds:** already present. Our implementation also preserves the separately established channel-haste behavior; replacing the whole upstream file would lose that distinction.
- **Demonic Brand's level-60 base/coefficient:** already present, including the relevant demon's damage school and our sourced always-hit handling.
- **Demonic Pact with a different active demon, and Voidwalker's Fel Energy:** already implemented. The previous apparent reset-order problem was disproved by full-run checks.
- **Master Demonologist affecting the demon too:** our implementation registers both owner and pet auras.
- **Berserk's three-target Bear Mangle:** already implemented.
- **Fire Nova receiving Clearcasting; Lightning Shield receiving Elemental Fury:** our spell predicates already cover them.
- **Blood Craze not inheriting physical damage modifiers:** our helpful spell ignores attacker modifiers.
- **Generic external pet-buff exclusions and Forever Battle Shout's non-talent-scaled value:** already represented. Explicit pet talents remain separate exceptions.

These do not need duplicate patches. Newly registered abilities still need to respect the same predicates and exception rules.

## What the latest “hotfix data” commit actually changes

In [e7223eb328](https://github.com/ElliotWood/Forever/commit/e7223eb3283306a86d32acac21b77c964d762722), the compressed spell-store file changes bytes, but decompressing and comparing its parsed content shows **no semantic spell-store changes**.

The generated UI database does change: it selects the stronger Trueshot rank, adds the Mana Tide display entry, revises item 5181 and adds item 284272. Those item rows are not our selected projected equipment. This commit is not evidence of an additional batch of new spell-number hotfixes.

The data-driven spell store, rank handling and proc infrastructure are useful longer-term work, but their schemas and engine differ substantially from ours. A wholesale merge would also disturb established equipment, pet, haste, resource and Totem rules. The focused findings above are the actionable output of this comparison.
