# October 8 overnight update: build 70291

## Status and source boundary

**Verified 70291 release package. Local verification and result review are complete.** Actual publication status and its receipt are in the [deploy.yml workflow](https://github.com/gunba/wow-forever-sim/actions/workflows/deploy.yml), not inferred from readiness or a push. Supported fixes, catalog/source corrections, final comparisons and bounded selections retain the provisional server-rule qualifications below.

The [official development notes, post 5](https://us.forums.blizzard.com/en/wow/t/2360696/5) were published on **October 8, 2026 at 21:32:38 UTC** (October 9, 05:32 Perth). The previous public simulator revision, `e7d0c2a6856ac1a6c947490f2f72784ccf44b400`, was released at 16:49 UTC, before these notes. [The earlier October 8 review](october-8-review.md) describes that earlier release; its results are not measurements of this overnight update.

Current client evidence is **1.60.1.70291**, created October 8 at 20:25 UTC. Primary tables include [SpellEffect](https://wago.tools/db2/SpellEffect?build=1.60.1.70291), [SpellPower](https://wago.tools/db2/SpellPower?build=1.60.1.70291), [SpellDuration](https://wago.tools/db2/SpellDuration?build=1.60.1.70291), [SpellAuraOptions](https://wago.tools/db2/SpellAuraOptions?build=1.60.1.70291), [TraitEdge](https://wago.tools/db2/TraitEdge?build=1.60.1.70291), [CurvePoint](https://wago.tools/db2/CurvePoint?build=1.60.1.70291), [ExpectedStat](https://wago.tools/db2/ExpectedStat?build=1.60.1.70291) and [SpellDiminish](https://wago.tools/db2/SpellDiminish?build=1.60.1.70291). Static data establish recorded values, not every server script or the deployment of a later hotfix. Another simulator is corroboration, not game authority.

Completed source-backed implementation changes have passed scoped native checks, including low-rank Fireball timing and the final registered-spell source guard. Remaining unimplemented features and unpublished scripts are explicitly distinguished below. Configurable assumptions remain assumptions. The [uncertainty register](uncertainties.md), including its preserved earlier answers, distinguishes these from implementation gaps and resolved source questions.

## Combat and class coverage

| Official change | Current disposition and implementation |
|---|---|
| Incoming Rage: expected creature health, crit/crush handling, absorbs ignored, reference Armor varying approximately 20–40% | Warrior and Bear use the explicit provisional model below on the **direct hit-callback pipeline only**. Periodic-only damage/logical channel children represented as periodic do not generate this modeled Rage; server eligibility remains unverified. Shield-restored post-outcome damage is retained **after actual Armor/other mitigation**; player maximum health is never the denominator. The notes do not publish the coefficient, level key or normalization operation. [WAR-003](uncertainties.md#war-003), [DRU-012](uncertainties.md#dru-012). |
| Druid mana potions in Bear/Dire Bear/Cat | Already supported without canceling form; active form-path checks corroborate it. No duplicate buff or form-removal change. |
| Bear outgoing critical Rage: +100%, previously +75% | Critical white-swing ratio corrected to 2. Bear's base outgoing/avoidance interpretation remains qualified; Maul is a queued replacement, not a white-Rage grant. Warrior outgoing normalization is not changed by the incoming-Rage note. |
| Thorns and Retribution Aura bases, 6% SP, correct casting provider | Registered maximum bases are now **18 and 20**, respectively, with **0.06 × provider school damage power**. Owned providers use explicit dynamic caster references; external providers default to declared **0 SP**. Recipient SP is never substituted. Party/provider coverage and receiver-owned reflection attribution remain explicit conventions. [PAL-010](uncertainties.md#pal-010). |
| Cat maximum health; Thick Hide gear changes while shifted | No local active-path defect was reproduced. Legal gear swaps exercise the existing Defense-to-Armor dependency in shifted forms; health/form lifecycle controls pass. This is an already-correct path, not a second armor or health buff. |
| Bear/Dire Bear threat: 50%, previously 30% | Form multiplier corrected from 1.3 to **1.5**, reversed on exit; Classic remains 1.3. Other ability-specific threat assumptions are not certified by this note. [DRU-005](uncertainties.md#dru-005). |
| Shapeshift appearance overrides | Game presentation/transform eligibility; not a new stat or damage multiplier. |
| Demoralizing Roar and Shout now generate threat on affected targets | Shout's former zero-threat branch is removed. Unset scenario inputs use **43.2 Shout / 100 Roar** at maximum registered rank, before stance/global modifiers. These are **Classic-derived conventions, not published Forever amounts**. Explicit per-target override supports zero or another finite amount; each landed affected target is handled independently. |
| Wrath; Frostbolt; Fireball; Arcane Missiles; Smite; Lightning Bolt; Shadow Bolt early-rank smoothing | **31 changed damage ranks** imported, including Lightning Bolt rank 5: Wrath 1–5, Frostbolt 1–5, Fireball 1–5, Missiles 1–3, Smite 1–4, Bolt 1–5, Shadow Bolt 1–4. Source variance and capped per-level growth are retained. Higher/max-rank values and SP ratios are unchanged; hidden downrank penalties and integer server rounding are not invented. A subsequent **confirmed duration discrepancy** is isolated to Fireball ranks 1/2/3: current client durations are **4/6/6 seconds**, not the inherited uniform 8 seconds. The Forever fix uses the sourced per-tick damage 1/1/2 every two seconds over 2/3/3 ticks, preserving totals 2/3/6. Actual timing/total/no-extra-tick controls pass; Classic and higher/max-rank 8-second timing remain unchanged. This is sourced timing, not an unverified duration assumption. |
| Predatory Instincts renamed Natural Instinct; 12/25% Intellect healing | Name and separate healing dependency corrected, preserving its existing damage/crit behavior and stable talent field. |
| Shifting Power at full Energy; Wolfshead Helm | Full-Energy cast rejected before spending resources or starting timers. Wolfshead adds **5 Energy**, not 20, to Shifting Power; its Rage portion and active Enrage/form bonuses remain unchanged. Dead Classic/unused branches are not treated as new fixes. |
| Hunter Disengage doubled threat reduction | Three learned ranks registered. Current client amounts are already doubled: −280/−560/−810 plus capped −3 per growth level; maximum rank at 60 is **−840**. No second doubling or damage payload. |
| Hunter low-health pet health accumulation; Unleashed Fury | Existing enable/inheritance teardown and damage multiplier paths pass active lifecycle controls. Unleashed Fury applies once to pet/Hawk damage; no duplicate 15% buff. |
| Improved Concussive Shot; Improved Wing Clip | Registered slow/control paths use sourced **4/8/12/16/20%** Concussive and **7/13/20%** Wing Clip chances. Target control eligibility is opt-in. Wing Clip's additional downrank penalty is unpublished; only the highest trained registered rank is modeled, not a fake penalty formula. |
| Expose Prey: 10 seconds | Separate 10-second opportunity state; the ordinary dodge-triggered Mongoose window remains 5 seconds. Mark ownership and attack-mask boundaries remain intact. |
| Entrapment DR; Frost Trap initial activation only | Entrapment triggers on successful initial trap activation, never periodic ticks. Explicit target/DR controls use the sourced groups and qualified reset choices below. Frost area geometry and Freezing Trap's actual freeze/break payload are not a complete CC engine. |
| Mage Impact excludes Flamestrike periodic damage | Previously unmodeled Impact now uses owner Mage-family direct landed Fire events only, including eligible direct procs, with **3/7/10%** current-tree chances and a two-second stun. Periodic Flamestrike does not trigger it. Default targets remain immune; shared stun DR is separately opt-in. The old 3/6/9 generator prior is not primary curve evidence. |
| Reckoning: 1.5-second cooldown | One shared **1.5-second gate** across critical-hit and block triggers, not two independent timers. Classic unchanged. |
| Seal of Fury mana without its talent; Improved Seal of Fury/Shield Specialization mana threat | Seal of Fury's seven learned ranks, white-hit damage, shield requirement, judgement and Echo are now registered. Only talented **actual full depletion** returns mana, using the actual incoming attacker level; replacement, expiry and cancellation do not refund. Both restoration effects use explicit NoThreat accounting. Shield lifetime/replacement is qualified below. Righteous Fury is a different spell. |
| Penance two-second channel, damage/healing and mana changes | All four **damaging** ranks are registered with shared 12-second cooldown and unchanged immediate/1s/2s bolts. At level 60 their per-bolt bases are approximately 44.61/52.97/71.87/92, SP coefficient **0.19**, mana **150/220/270/385**. Maximum formerly used 131/0.285/355. Helpful Penance is not registered. |
| Lesser Heal/Heal smoothing; Penance healing ranks 2/4 | Complete changed healing rows are captured, but **Lesser Heal, Heal and Penance's healing branch remain unregistered feature gaps**. The notes describe healing increases at Penance ranks 2/4 while exported overall bases fall (rank 4: 673→425, coefficient .285→.19); no guessed extra healing multiplier resolves this conflict. [PRI-005](uncertainties.md#pri-005). |
| Flametongue Totem shapeshift attack speed | Already reads active Cat/Bear weapon speed (1/2.5 seconds), not equipped weapon speed. This corroborates speed routing, not the still-provisional Flametongue damage coefficient. [SHA-006](uncertainties.md#sha-006). |
| Elemental Focus and Fire Nova | Existing just-activated guard keeps a fresh charge; a later eligible cast consumes it. Actual Nova controls pass; no extra Clearcasting charge was added. |
| Water Shield amount/cooldown tooltip; Mana Tide visual | Game display fixes. Existing modeled mana/cooldown behavior is not changed because a tooltip falsely said 15 seconds or a pulse visual was absent. |
| Warlock old pet stunned while summoning a replacement | Successful cast start stuns the old pet; failed validation does not dismiss it. Cancel/completion/overlap/reset lifecycles restore or replace correctly. A permanent pet does not receive a recurring penalty. |
| Demonic Brand high threat only for Shadow | Already corrected in the prior public release: Fire ordinary threat, Shadow retained high-threat model. No duplicate overnight threat change. [WL-001](uncertainties.md#wl-001). |
| Impale requires Deep Wounds; Deep Wounds excludes off-hand AP and rolls | Restored sourced tree prerequisite; published Warrior allocations already satisfy it. Existing weapon-only/no-AP rolling bleed with preserved tick timing is corroborated, not buffed again. Full server payload/hand attribution remains qualified. [WAR-016](uncertainties.md#war-016). |
| Last Stand health activation | Corrected health-multiplier double application. A cast snapshots a 30% temporary final-health grant, compensating raw Health for Tauren/Mining factors; later item changes do not grow that snapshot. Expiration removes the temporary maximum **and its health grant**, rather than retaining a free heal or creating fake attack/TMI damage. Dead units stay dead; a living one-HP floor is a **Classic temporary-health convention**, not a measured Forever floor. |

## Items, professions and catalog boundaries

- **AP armor kits:** four source-backed records 8483/8486/8488/8491 grant AP/RAP once on chest/legs/hands/feet across Cloth/Leather/Mail/Plate. Enchant wearer level is **0**; item-level minima are **0/15/25/35**, distinct from source crafting/use levels. Leatherworking is crafting, not a wearer restriction. Maximum 8491 grants **40 bonus Armor, 10 AP and 10 RAP**, not 20 RAP. The former extra RAP grant is not a reason to halve every Hunter's existing AP.
- **Recovery:** equipped enchant 8721 follows its sourced dodge/parry route, heals 5% current maximum health, has a shared ten-second ICD and works in forms. Miss/ranged/magic/unequip controls are excluded; its crafting recipe does not require the wearer to be an Enchanter.
- **Revelation:** form retention, 15-second payload, one charge, crit bonus and eligible family are sourced. Proc probability remains unpublished. Its seven class-specific family payloads are an exact whitelist, not every spell: Fireball, Pyroblast, Aimed Shot and Multi-Shot are not automatically eligible. Represented channels qualify only through an actual matching SCHOOL_DAMAGE child; a Revelation-only logical alias does not globally reclassify periodic damage. Reservation of the first eligible calculated target versus a whole next cast is another model convention. The explicit profile-owned model is now integrated locally: its nine-test suite and all five represented channel bridges passed, with separate healing-architecture and TypeScript preservation checks. No universal hidden 5% is supplied; final benchmark/selection/sensitivity evidence is recorded below, while the server proc law remains unverified and deployment is separately evidenced. [DATA-015](uncertainties.md#data-015).
- **Chillknife:** temporary enchant 8698's source duration is **3600 seconds**, previously zero, with Mage/dagger/use-level-5 restrictions and slow child 1296223. Proc cadence and acquisition remain unestablished; no guessed PPM, permanent-enchant optimizer entry or validated DPS contribution is supplied. This is an explicit remaining feature/evidence gap, not an unfinished duration audit.
- **Weapon swaps:** actual failures involving same-ID/different-enchant weapons, suffix identity, sequential MH/OH transitions, unique-item validation and iteration-reset stat drift are corrected. Multiple swaps use the existing GCD semantics; this does not create extra attacks or a free speculative DPS bonus.
- **Real-item audit:** all 153 changed and 21 added records (none removed) have leaf dispositions. Only 11 changed IDs were already compiled: seven vendor-price-only records and **four Spiritcaller pieces (276538–276541) now restricted to Shaman**, exposing a real missing class-import validation path. The final source overlay, compiler/DB generation and all four actual native class-restriction checks pass; ordinary UI imports have separate controls. This resolves the identified catalog/import fix, not every live acquisition route. Two added reviewed quest items, knife 251485 and cuffs 251486, have sourced stats; the other 19 additions are consumables/relics/bag/key records. Thirty-eight changed unlisted combat items still lack supported acquisition rather than being silently granted. [DATA-016](uncertainties.md#data-016).
- **Dungeon rewards:** rarity/stat/source changes require real-item catalog and acquisition review. The 32 numerical reference IDs in the prior modeled gear inputs did not change; that does not establish optimal new loot or availability. Truthseeker's Bow is now level 40. Desolace reward tuning and the *announced upcoming* level-35 equip hotfix are distinguished from already applied records: Abandoned Ferocity, Blade of the Magram Clan, Bludgeon of Betrayed Virtues, Centaur Spear, Ceremonial Centaur Blanket, Desert Crawler's Claw, Desperate Barrier, Greatstaff of the Necrokhans, Kolkar Hunter's Belt, Kolkar Marauder Chain, Scavenged Magram Armament, Soulsplatter Mace, Staff of Revelations, Thrice-Damned Effigy, Traitor's Finger and Wanderer's Broadsword. A static item record does not prove current quest completion/training access.
- Shoulder visual size, Transformative Cocoon's debuff display and Buckshot's removed knockback are game/catalog changes, not new stationary PvE damage bonuses. General knockback/CC-break behavior is not modeled by these reference encounters; Discolored Healing Potion's no-CC-break fix does not grant extra damage or utility credit.
- Vendor-only changes: Peace Tea 10c; Minor Wizard Oil 1s; Novice's Practice Wand 5c; Lesser Magic Wand 1s10c; Greater Magic Wand 2s; Lesser/Greater Mystic Wand 4s50c/6s; Twisted Nether/Dreambough/Lesser Eternal/Greater Eternal Wand 10s/15s/20s/25s. Prices are not combat effects.

## Coverage statements and historical reports

An earlier generated build caveat incorrectly said Seal of Fury was implemented. The previous public Protection profiles actually used Seal of Righteousness; the new SoF implementation belongs to this **70291 release package**. Historical benchmark/selection documents and the sampled spell inventory are not current runtime proof. Setup verification also exposed a real default-support defect: permanent external Faerie Fire masked the Bear APL's explicit self-maintenance. Native and web tank support now remove that external aura for Bear only; all four races actually cast Faerie Fire in both workloads. The earlier free-aura Bear study phase is superseded, not silently reused. Warrior and Paladin external support is unchanged. Current source selection now supports SoF subject to its explicit shield model, but neither that nor captured helpful spell rows means all healing/utility spells are implemented. Generated profile caveats and static inventory dispositions are reconciled. All six registered-source guards pass with **1128 entries, zero missing/orphan/errors, 44 unreviewed and 13 unresolved source labels**; source labels are not a claim of verified server mechanics. The public register retains **160 stable IDs, 42 resolved and 12 needs-code** questions: closing the bounded selection workflow does not resolve its provisional numerical models. [40 official dispositions and two observed integration regressions](../artifacts/correctness/october_9/official-dispositions.json).

## Explicit assumptions, not resolved scripts

| Model | Exposed scope and limits |
|---|---|
| Incoming Rage | **Direct callback only**; periodic received damage is not covered. Central coefficient **10**; player-level ExpectedStat CreatureHealth; reference Armor `clamp(CreatureArmor/(CreatureArmor+ArmorConstant), .20, .40)`; `(Damage + AbsorbedDamage) × coefficient / expectedHealth / (1-referenceArmor)`. The coefficient, level selector, curve and operation are assumptions. Optional coefficient/reference-Armor/expected-health overrides retain absence versus explicit zero; finite validation bounds 0–1000 / 0–.95 / 1–1e9 are input limits, not game laws. |
| Seal of Fury | One **shared replacement pool**, not rank-specific accumulation. Shield equals 50% actual dealt proc Holy damage when a shield is equipped. Optional lifetime defaults to **30 seconds**, explicit zero disables absorption only, finite 0–3600. The seal's 30-second duration does not prove shield expiry. No accumulation mode, invented judgement absorb, capacity cap or complete aggro system. Four-second judgement taunt aura only; no highest-threat equalization/AI target switching. Canonical default profiles have not been switched to SoF. |
| Demoralizing threat | Optional per-target scalar; unset uses qualified 43.2/100 maximum-rank conventions, explicit zero is a sensitivity control. Finite 0–1e6. Nonzero behavior is official; these exact amounts are not. |
| Crowd-control DR | Target stun/root/slow eligibility and DR are opt-in, off in default scenarios. Entrapment mask1→Root1; Impact/Improved Concussive mask512→proc Stun10. Primary values: **15 seconds, ×0.5, immunity after three applications**. Reset after final effect end versus after application is exposed and unverified; target-policy field remains undecoded. Wing Clip group0 and slows have no invented DR. |
| Revelation | Integrated explicit conditional noncrit model `baseChance × (1-effectiveCrit)^critExponent`, with eligibility/miss policy separately declared. Exponent zero is a flat **conditional** probability, not flat per-cast probability. No inferred universal chance or optimized unknown curve. |
| Temporary health and reactive providers | Last Stand's living one-HP expiry floor, reflection hit/threat attribution and configured buff range are conventions. External reflection SP defaults to zero; owned provider references are explicit and dynamic. These are not receiver-SP inference. |

## Regression references and current comparisons

Scoped passing checks are recorded by the current test sources, not by old benchmark totals:

- [Incoming Rage and native profile validation](../sim/core/forever_incoming_rage_october9_test.go): stamina independence, crit/crush, partial/full absorbs, actual versus reference Armor, levels/overrides, Classic and binary/JSON optional presence.
- [Patch/tree/Last Stand/Demoralizing checks](../tools/forever_bench/october9_patch_test.go): legal prerequisite boundary, outgoing Bear ratio, health activation/expiration and multi-target threat/override/Classic controls.
- [Early ranks and Penance](../tools/forever_bench/ranks_october9_test.go): changed source ranges, ranks/costs/shared timers and retained timing/Classic.
- [Mage Impact](../tools/forever_bench/mage_impact_october9_test.go); [Druid/Hunter/Shaman paths](../tools/forever_bench/druid_hunter_shaman_october9_test.go): direct/periodic boundaries, immunity/shared DR and actual form/pet/trigger paths.
- [Paladin](../tools/forever_bench/paladin_october9_test.go), [reflections](../tools/forever_bench/reflection_october9_test.go), [SoF](../tools/forever_bench/sof_october9_test.go), [SoF model](../tools/forever_bench/sof_model_october9_test.go) and [finite depletion](../sim/core/shield_depletion_test.go): gate, provider identity, NoThreat, rank/shield/talent/miss and lifetime boundaries.
- [Recovery/catalog](../tools/forever_bench/enchant_october9_test.go), [summon stun](../tools/forever_bench/warlock_summon_october9_test.go) and [weapon swaps](../tools/forever_bench/item_swaps_october9_test.go) cover their scoped regressions. [Revelation](../tools/forever_bench/revelation_october9_test.go) now has passing explicit-model/whitelist/charge/form and all five channel-bridge controls; these validate implementation, not a server probability law. Profile/SavedSettings optional-input JSON/binary roundtrips passed **458 checks**; this is configuration evidence, not game-mechanic measurement.

The [exact previous-release archive](../artifacts/history/e7d0c2a68/published-before-70291.tar.gz) and [manifest](../artifacts/history/e7d0c2a68/manifest.json) preserve the earlier inputs/results and immutable research references. They are historical controls, not new-engine outputs.

## Current results, setup decisions and conditional limits

**Level-60 forecasts with hypothetical modeled-v2 equipment—not level-30 beta loot, measured logs or a verified global optimum.** Full 51-point allocations, ordinary MP5, role Tier 1 and natural hit remain. No extra 16 Legacy talent points are added.

The final refresh contains **1,005** native scenarios: all **201 profiles / 31 builds**, each at baseline, Tier 1 off, +10% equipment, +50% equipment and corrected-engine original-loadout control. Every complete baseline player and actual request player is audited; Tier/scaling flags are exact, warnings are zero, and equipment scaling leaves enchants/effects/recipes fixed. Seed **1815000001**, **5,000 iterations** each. Applicable rows retain resolved `ProvisionalModels` values and override-presence flags; unset versus explicit zero are not collapsed.

[Current chart](../artifacts/modelled_gear/forever_dps_5min.png) · [Current raw 201 requests/results](../artifacts/modelled_gear/forever_dps_5min.json) · [Tier/equipment sensitivities](../artifacts/modelled_gear/forever_sensitivity.json) · [Corrected original-loadout controls](../artifacts/correctness/october_9/original.json.gz) · [Full comparison](../artifacts/correctness/october_9/comparison.json) · [Input/matrix audit](../artifacts/correctness/october_9/matrix-audit.json) · [Selection evidence](../artifacts/correctness/october_9/selection-evidence.json).

Exactly **23 players change**, limited to **5 equipment, 5 talent strings and 18 rotations**, overlapping; profession 2 and all other player fields are exact originals. Ten Warriors and three Paladins change; four Bear players are retained. Bear **support** explicitly migrates external Faerie Fire true to absent/false, so the APL performs its own 9907 duty. Prior release1 Bear study phases are quarantined as invalid-duty evidence. Historical requests are not falsely called unchanged; all 13 Warrior/Paladin external Faerie Fire settings remain unchanged.

### Before, corrected controls and selected profiles

Equal-weight available-race mean DPS. Historical published values use pinned `e7d0c2a68`; their difference from fresh controls also includes Monte Carlo sampling at a new seed, not a same-seed causal engine estimate. Setup deltas compare current selected versus current original loadouts at the same final-matrix seed.

| Build | Races | Earlier published | Corrected original loadout | Current selected | Setup delta |
|---|---:|---:|---:|---:|---:|
| affliction | 5 | 824.31 | 824.00 | 824.00 | +0.00 |
| arcane | 6 | 730.75 | 730.95 | 730.95 | +0.00 |
| arcane_frost | 6 | 732.58 | 732.36 | 732.36 | +0.00 |
| arms | 10 | 832.00 | 832.33 | 832.33 | +0.00 |
| balance | 4 | 710.87 | 710.61 | 710.61 | +0.00 |
| beast_mastery | 8 | 902.22 | 902.29 | 902.29 | +0.00 |
| combat | 9 | 710.91 | 711.15 | 711.15 | +0.00 |
| demonology | 5 | 917.72 | 917.36 | 917.36 | +0.00 |
| destruction | 5 | 815.11 | 815.12 | 815.12 | +0.00 |
| ds_ruin | 5 | 793.50 | 793.50 | 793.50 | +0.00 |
| elemental | 5 | 491.81 | 491.45 | 491.45 | +0.00 |
| enhancement | 5 | 812.10 | 812.07 | 812.07 | +0.00 |
| feral | 4 | 723.10 | 723.26 | 725.03 | +1.77 |
| feral_tank_druid | 4 | 609.07 | 544.72 | 544.72 | +0.00 |
| fire | 6 | 729.45 | 729.34 | 729.34 | +0.00 |
| frost | 6 | 684.00 | 683.53 | 683.53 | +0.00 |
| fury | 10 | 946.61 | 946.99 | 946.99 | +0.00 |
| fury_2h | 10 | 829.88 | 829.34 | 829.34 | +0.00 |
| fury_sunder | 10 | 917.24 | 917.43 | 917.43 | +0.00 |
| marksmanship | 8 | 777.41 | 776.65 | 776.65 | +0.00 |
| mutilate | 9 | 700.94 | 701.11 | 701.11 | +0.00 |
| pet_melee | 8 | 898.69 | 905.33 | 905.33 | +0.00 |
| protection_paladin | 3 | 595.08 | 594.06 | 599.46 | +5.40 |
| retribution | 3 | 1053.68 | 1054.01 | 1054.01 | +0.00 |
| retribution_physical | 3 | 958.91 | 959.51 | 959.51 | +0.00 |
| shadow | 6 | 829.14 | 829.08 | 829.08 | +0.00 |
| smite | 6 | 470.28 | 427.08 | 449.77 | +22.69 |
| stormcaller | 5 | 456.27 | 455.74 | 455.74 | +0.00 |
| subtlety | 9 | 656.20 | 656.39 | 656.39 | +0.00 |
| survival | 8 | 884.84 | 891.78 | 891.78 | +0.00 |
| tank_warrior | 10 | 367.90 | 370.89 | 377.24 | +6.35 |


### Tanks: central conditional selection, not global robustness

Selection's two disjoint **5,000-iteration** windows begin **1802000001 / 1803000001**: **152 workload results / 38 unique players**, including originals, proposals and three intentional native Paladin defensive equipment anchors. Single DPS is the objective. Each changed selection is positive in both windows and exceeds `1.96 × sqrt(sum((SEcandidate + SEbaseline)^2)) / 2`. All five **point-estimate** limits must pass both workloads in each window: DTPS +1%, TMI +2%, death +0.0025, TPS −1%, least-target TPS −2%. This is **not confidence-proof of no survival loss**.

Largest selected per-window increases are DTPS **+0.3405%**, TMI **+0.4893%** and death **+0.0012**. Undead Warrior gains **14.415 Single DPS** but loses **14.175 Multi DPS**; Orc/Dwarf Multi losses are **3.24 / 3.44 DPS**. Multi TPS guards still pass. Paladins retain **SoR, Brilliant Wizard Oil, Swift and Holy Shield**, not SoF.

Fresh **post-source-sync native-default verification**, not reselection, uses new **1817000001 / 1819000001**, 5,000 each: **68 selected workload rows plus 68 exact corrected-original controls**. Selected CLI paths use no baseline/player override. All 17 races match exact selected players; all five guards pass both workloads in each window, with zero warnings. [Current tank validation](../artifacts/tanks/current/validation.json) · [All fresh paired requests/results](../artifacts/correctness/october_9/tank-default-checks.json.gz).

Incoming Rage's central scenario remains **UNVERIFIED**: coefficient **10**, player level **60**, expected creature HP **4365**, reference Armor **0.39333774542245753**. Actual damage/mitigation/absorbs are not replaced by reference Armor. Demoralizing defaults **43.2 Shout / 100 Roar** remain provisional conventions. The fixed-input 1,000-iteration sensitivity grid evaluates 240 cases; no unknown parameter is optimized. **Six alternative-model guard failures are material:**

| Warrior race | Incoming coefficient | Reference Armor | Failed guard | Single DPS change |
|---|---:|---:|---|---:|
| Orc | 20 | 0.20 | Single TMI | +1.040 |
| Orc | 20 | central 0.3933377454 | Single DTPS; gain also negative | −0.242 |
| Undead | 20 | 0.20 | Multi TPS | +16.304 |
| Undead | 20 | central 0.3933377454 | Multi TPS | +14.215 |
| Undead | 20 | 0.40 | Multi TPS | +16.154 |
| Dwarf | 20 | 0.40 | Single TMI | +0.085 |

These failures prevent a global-robustness claim. They are not alternate optimal recommendations or server calibration.

All 10 Warrior races build five Sunders by **10.5–13.5 seconds** in both workload diagnostic traces; no midfight stack decline and successful-refresh/fight-end gap at most **28.3 seconds**. All four Bears actually cast 9907 **8 times / 300 seconds** or **3 times / 90 seconds**, first around 3 seconds, without external Faerie Fire. These are actual one-iteration stack/cast traces, **not proof of duty perfection across 5,000 iterations**. Canonical 17 native modeled gear searches retain **9,676 trials**, all converged in 1–3 passes; **1,527 superseded Bear trials** remain historical. All four Forceful kits and Recovery were tried, not promoted; Revelation exclusion states explicit model/payload requirements, not fake item illegality.

All three Paladin SoF lifetime **0/5/30/60-second** fixed-recipe cases lose about **50 Single DPS** and fail TPS guards. At 5/30/60 shields deplete quickly; zero disables absorption/mana return and increases DTPS. One shared replacement pool and its 30-second unset lifetime remain provisional. The sourced four-second taunt aura is **not aggro equalization or a target-switch engine**.

### DPS: ten rotation-only migrations and costs

Six Smite profiles remove routine Penance: five retain the highest-rank finisher only at remaining time ≤2.5 seconds; Undead has no Penance action. Holy Fire rank 8/prepull, maintenance and **Inner Focus + Smite** remain. Both independent 5,000-iteration windows (**1807000001 / 1809000001**) confirm gains **20.88–29.15 DPS**, above the stronger sum-SE bound. Correctly sequenced Inner Focus + Penance loses to selected IF + Smite in all six profiles (**0.42–2.79 DPS**). The initial priority-starved pilot is explicitly non-material. Old Penance was **131 / 0.285 / 355 mana**; current is **92 / 0.19 / 385**. Six bounded native Smite gear searches contain **2,308 trials**, converged at pass 1 with no gear/talent promotion.

Four Feral Cat profiles move the existing finisher block before Shifting Power and use energy ≤60. Gains **1.92–2.14 DPS** pass both independent windows/bound. The energy grant stays **40 per event**, at unchanged source mana cost. **OOM increases**: Tauren **1.092→2.701 seconds**, Night Elf **0.943→2.433**, Windshaper/High Order **0.832→1.866**. This is a DPS-positive resource tradeoff, **not mana efficiency**. No current head is Wolfshead.

All seven final Forceful-kit proposals fail positive-both/stronger-bound acceptance; none is promoted. **114 AP/RAP profiles** receive four-slot screens; **70 caster profiles** have explicit analytic exclusions. Warlock pets inherit **max(AP, RAP, SP + school SP)**, not “pets never use AP”; current caster SP remains dominant by at least **181.2** even after all four kits. Fourteen duplicate final control arms are retained but are **not additional independence**.

### Revelation: unknown-law sensitivity, no canonical promotion

**112 arms**, seven represented classes × eight cases × two **1,000-iteration** windows (**1811000001 / 1813000001**), preserve exact original recipes except declared main-hand 8217 and explicit model. Smite is the **original routine-Penance research control**, not the revised default. All requests are warning-free with full player/support preservation. Cases cover zero, base .05/.10/.20 at exponent 1, exponent 0/2, miss policy and original enchant. Exponent zero is flat **conditional noncrit** chance, not flat per cast. Source crit/family/event/AoE/reservation conventions are fixed, not proven by this grid.

Balance's rune comparison changes from **−12.98 DPS at base .05** to **+9.79 at .20**; other classes differ. This dependence prevents an honest universal recommendation. No server curve is calibrated, no Rune 8217 default is promoted; Warrior/Rogue source payloads remain unsupported with explicit native errors. Helpful Priest spellbooks, Ghost Wolf and Chillknife PPM/acquisition remain concrete feature/evidence gaps, not a complete healer or utility simulation.

### Reproducibility and runtime boundary

Final benchmark SHA256 `512b673d7b5c86433feadab6e1aa3ef805f817f9b0f4638aaac1df9bcdb96290`; generic replayer `350a24878f176b9255a09a2a389843ed9709b4a860eacdd3db707ca85b579056`. Release2 mechanics and binaries remain separately preserved. Final source runtime change is race-default selection, not combat math. Five captured explicit-input class/role comparisons have exact requests and direct metrics after only unordered Go action/aura/resource-array normalization. One derived Warrior target-TPS sum differs **849.3994299187007 vs 849.3994299187005** from map summation order; only those identified derived sum fields are disclosed, not globally rounded or used to relax any guard.

Native defaults, source race gear/APLs/talents and actual UI ranked profiles are synchronized, not merely matrix override JSON. Native study archives retain commands/full inputs/results/per-seed statistics and phase audits: [tank study](../artifacts/correctness/october_9/tank-selection-native.tar.gz), [DPS study](../artifacts/correctness/october_9/dps-setups-native.tar.gz), [Revelation study](../artifacts/correctness/october_9/revelation-sensitivity-native.tar.gz), [archive hash/privacy receipts](../artifacts/correctness/october_9/native-evidence-receipts.json) and [raw engine-equivalence pairs](../artifacts/correctness/october_9/engine-equivalence-raw.tar.gz). [Runtime/input provenance](../artifacts/correctness/october_9/runtime-manifest.json) · [Explicit-input equivalence](../artifacts/correctness/october_9/engine-equivalence.json). Private client/community/chat/account/GUID material is not published; evidence packets contain native synthetic build names, commands, exact requests/results and declared phase statistics only.

### Reproduction

Build the current source with Go 1.24 and the checked-in database; no binary is committed. Extract the complete original-loadout players from `original.json.gz` into a separate input file, preserving each `Key`, `Race` and `BaselinePlayer` and `GearScenario: modeled-65-v2`. That request's current support intentionally removes external Faerie Fire for Bears; do not restore the historical free aura.

```sh
go build -tags with_db -o /tmp/forever-bench ./tools/forever_bench
go build -tags with_db -o /tmp/forever-replay ./tools/forever_tanks
python3 tools/forever_bench/run_matrix.py --binary /tmp/forever-bench \
  --profiles artifacts/modelled_gear/forever_input_profiles.json \
  --original-baselines /tmp/original-70291-loadouts.json --natural-hit \
  --iterations 5000 --seed 1815000001 --workers 24 --output /tmp/70291-matrix
# Representative fresh native-default pair: no baseline/player override.
/tmp/forever-bench -build tank_warrior -race Orc -natural-hit -tank-pair \
  -iterations 5000 -seed 1817000001 -output /tmp/70291-orc-default
```

The fresh tank receipt lists every race, both 1817000001/1819000001 windows and exact corrected-original control commands. Archived study requests use their own preserved runtime phase hashes, not silently substituted later defaults.

**Publication gate: `make verify` passed.** Native/core/class suites, all six source guards, Python/catalog/register checks, scrubbing and all built pages pass. Browser checks audit **all 201 exact race/build bindings**, replay **all 31 builds and all 17 tank races** against native DPS/TPS/DTPS/TMI, and exercise fresh tank defaults and their owned duties. Identified integration issues—stale source assertions, unique-equipped test fixture, old tank goldens, missing real recipe icon and literal APL-label regex handling—were corrected without changing combat math or weakening guards. [Golden causal restoration](../artifacts/correctness/october_9/golden-causality.json) preserves the private old-source-overlay proof.

The [deployment workflow receipt](https://github.com/gunba/wow-forever-sim/actions/workflows/deploy.yml) and actual live resource/WASM verification determine publication; completed local checks and a push alone do not.


## Concrete beta and launch follow-up

Current testing is bounded by the **level-30 beta** and actual learned ranks/talent prerequisites. The new **16 Legacy Points** improve access to tests; they are not additional class points and do not create a new 51-point allocation. The **City of Dalaran** supplies a new dungeon/party context, not a measured level-60 raid boss.

1. **Incoming Rage:** Warrior and Bear, no autos/spending/passive Rage, exact build and fractional power trace; independent Stamina/Armor changes, ordinary/crit/crush hits and partial/full shields. Repeat against different attacker levels to separate the expected-health level key. A one-level fit establishes only an effective coefficient, not both coefficient and reference curve. Add a **DoT-only then isolated channeled-bolt eligibility** experiment: disable outgoing autos, Bloodrage/Berserker Rage and other on-hit bonuses; record actual damage/absorbs and uncapped unmodified fractional power over enough ticks to avoid an integer-display false zero. The current callback omission is clear code coverage, not proof that live DoTs should generate or should not generate Rage.
2. **SoF:** level 20+ with the talent, shield and learned ranks; successive procs before damage, rank switching, partial/full depletion, expiry/replacement and equal/up-to-three-level-higher attackers. Record capacity and mana independently. Compare configured no-absorb/short/30-second scenarios without claiming one is the decoded script.
3. **Reflections and demoralizing threat:** distinct caster/recipient; change caster SP after application without changing recipient gear. Isolate Shout/Roar against one and several targets, misses and zero-resource controls; determine flat amounts before stance multipliers. Dalaran packs can expand this test, but uncontrolled pack damage is not a coefficient sample.
4. **Hunter/Mage controls:** eligible cooperating targets, then immune controls. Test initial trap activation versus later ticks, cross-caster Root1/Stun10 reductions, breaks and 15-second reset after application versus effect end. Wing Clip downranking needs an independently available lower-rank cast; current highest-rank registration cannot validate the missing penalty.
5. **Recovery/Revelation:** obtainable recipes/enchants and correct item requirements first. Recovery can isolate dodge/parry and form retention. Revelation needs noncrit/crit/miss and direct/periodic family-separated samples, current SP/crit, charge consumption and shapeshift retention; first establish availability before estimating a hidden proc law.
6. **Ranks, training and prerequisites:** validate learned low ranks at level 30; Penance rank 1 starts at 30, higher ranks at 40/50/60. Full Deep Wounds is required before Impale. Static higher-rank/book records are not trainer/drop proof. Beyond-30 and level-60 coefficient, gear and acquisition tests remain launch/later-access work, not silently moved into today's beta.
7. **Last Stand:** record temporary maximum/current health before, during and after expiry, including a near-death expiry. Confirm the living one-HP convention separately from the source-backed removal of temporary health.

## Complete non-combat and game-only dispositions

The remaining official bullets do not change the current stationary level-60 PvE damage/threat/survival workload:

- Classic empty-corpse looting, large-creature effect flicker and unintended long-distance pulls; Dalaran availability; dungeon XP approximately +20%, individual-level party XP allocation and the retained vastly-higher-party-member no-XP condition.
- Cooldown Manager support for classes/racials/trinkets, with consumable support still future; gamepad held/pressed indicators, post-loot input, map focus, raid traversal, edit-mode support/layout storage/settings migration.
- Legacy Challenge grant of 16 points; beta respec minimum 1s and one-respec-per-hour cost recovery. No extra class talent budget.
- PvP critical efficacy and Shipwreck Cove capture-point location. No corresponding PvE crit multiplier is invented. Undead Will of the Forsaken sleep/charm usability is game control behavior, not a general modeled fear/sleep/charm encounter.
- Blood Tithe starting conditions, Night Watchman's Torch visual, replacement capital-city barbers; Barrens Echeyakee horn reuse and Principal Source marker; Desolace Valley of Bones minimum level35 and reward tuning.
- Dun Morogh Far Sight input/cancel; Durotar egg supply and Lazy Peons marker/behavior; hostile Stitches; Wetlands follow-up Titan Relic restoration; Hillsbrad Elixir of Pain XP redistribution; Loch Modan ambush frequency; Mulgore markers/respawn/XP; Orgrimmar Neeru marker/tooltip; Redridge supply respawns.
- Rogue Plundering markers; Shaman Storm/Earth/Fire spawn/drop sources and Eye of the Tempest multi-drop; Stonetalon quest turn-in availability; Teldrassil reset/minion/escort timing; Tirisfal quest/creature level escalation; Warrior Ulag trigger reactivation.
- Discord voice/group-finder voice selection, inspect layout, combo-point displays, quest types/level requirements, buff-stat highlighting, macro tabs, Legacy navigation, View Talents and character-pane class/level display.
- Known issues: future Skyborne/Zephras/mount/pet texture updates, **Ferocity pet-sheet crit display only**, and broken profession icons. The visual Ferocity issue is not evidence to remove its actual pet crit effect.

This covers the complete post, including announcements and known issues. It does not imply every game system, healing spellbook, PvP control rule or newly obtainable item has a simulator implementation.
