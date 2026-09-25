# DPS verification

Reviewed against the current **27-build / 174-profile benchmark** and updated
on 23 September 2026. This is a list of remaining DPS uncertainties,
not a request to retest every datamined number. Client build **1.60.1.69977**
has identical Spell, SpellEffect, SpellAuraOptions, TraitDefinition,
TraitDefinitionEffectPoints, CurvePoint and PowerType exports to 69913; the
[official update](https://us.forums.blizzard.com/en/wow/t/beta-client-update-september-22/2358655)
does not claim a change to these abilities.

The [disposition register](check_dispositions.md) accounts for **all 51 original
checks**, including resolved facts, implementation work and deferred features.
It also records sources and the damage contributions used to set priorities.

## Answers available without another gameplay test

- **Mage Armor is live in all 24 Mage profiles**, including Arcane–Frost.
  Molten Armor is not an available option.
- **Healing-only bonuses do not grant inferred damage.** The healing enchants
  that also add damage have explicit client effects: bracers 24 healing / 8
  damage, gloves 35 / 12, weapon 55 / 19. No evidence supports an additional
  blanket conversion.
- **Druid rank values are resolved from the active trait curves:** Nature's
  Majesty 2/4%; Nature's Reach 2/4% hit; Moonglow 8/17/25% cost reduction;
  Moonfury 2/4/6/8/10%. A standalone rankless tooltip is not the talent curve.
  Damage-bonus scope is a separate question below.
- **Demonic Pact explicitly permits a different active demon alongside the
  sacrifice bonus.** Resummoning the sacrificed demon cancels it. The reported
  no-Pact coexistence bug was disproved by the pet-reset order.
- **Fire Nova requires a Fire totem.** This is sourced and enforced.
  Its threat attribution is not a DPS-validation task.
- **Consecration's current single-target formula is sourced:** rank 1 has
  eight ticks of `2 + 4 + 0.095 × SP` before modifiers. Geometry and fifth-target
  behavior do not affect this encounter. There is no outstanding coefficient
  discrepancy requiring a Consecration test here.
- Hot Streak and Mage Clearcasting have sourced charge limits; neither grants
  a whole buff-duration window of free/accelerated casts.

### Code follow-ups, not questions for a player to rediscover

**Penance timing:** the earlier engine emitted at 0.667, 1.333 and 2 seconds.
It now emits the client-stated immediate/one-second/two-second bolts without
haste, covered by a focused regression. Its channel period scales with
casting haste in the Forever model; the level-20 check remains T40.

**Pet inheritance:** Blizzard's
[beta known-issues post](https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-known-issues-september-18/2352687)
confirms owner-derived pet hit, but gives no coefficient. Both Hunter and
Warlock now use the shared beta-wiki owner-stat assumptions. The existence
of inheritance is no longer an open question; the hit and Warlock-specific
power coefficients still are.

## Highest-value checks

Keep spell IDs, client build, level, equipment and relevant buffs with any
recording. A level-20 result establishes that case, not automatically level 60.

| ID | Remaining question and why it matters | Smallest useful check |
|---|---|---|
| **T01 / T06** | Physical miss, crit suppression and higher-level glancing damage. The engine retains Classic tables. Blizzard flags higher-level melee glancing damage and displayed tables as incorrect; an old character-sheet screenshot is not a reliable replacement formula. | On a build where that defect is fixed, log maximum-skill white attacks against equal- and higher-level targets, separating single-weapon and dual-wield attacks. Preserve glancing damage, not just counts. |
| **T02** | **High priority: confirm the provisional level-60 Warrior rage model.** [Level-8/9 raw logs](https://github.com/magey/forever-warrior/issues/3) support ~3.46 rage per weapon-second one-handed and 4.5 two-handed, no outgoing damage term, and incoming unmitigated damage × 10 / maximum health in geared cases. The engine extrapolates these results to 60 and assumes a half-rate off-hand. Low-armor incoming hits doubled rage at level 8/9, but their cause and relevance at 60 are unknown. | Supply level, weapon speeds, off-hand identity and matching resource/equipment events for one-handers, two-handers, off-hand and extra swings. Distinguish hit/crit/glance/dodge/parry/miss. For incoming damage, record maximum health, armor, pre-/post-armor damage, Rage, and Berserker Rage state. Repeat higher up before trusting the Warrior ranking or tank formula. [Evidence and arithmetic](check_dispositions.md#rage-evidence). |
| **T50** | Queued Heroic Strike/Cleave and off-hand miss. A [level-20 test](https://github.com/magey/forever-warrior/issues/2) found 4 misses in 77 queued off-hand swings versus 18.27% without a queue. The engine provisionally removes the dual-wield miss penalty while queued; the client still marks these attacks on-next-swing, contrary to a newer WoWSims instant-special change. | At 20, confirm that Heroic Strike waits for the next main-hand swing; log more queued and unqueued off-hand swings against the **same target level** with unchanged hit and skill. Repeat later to establish the level-60 rule. |
| **T51** | Forever base attributes by level/race/class. The [WoWSims branch](https://github.com/wowsims/forever/commit/0c8e81e0ab) explicitly still uses level-70 primary attributes at 60; our different level-60 rows are not independently established as Forever values. This affects every spec and race. | At level 20, record naked, untalented Str/Agi/Sta/Int/Spi and health/mana for several race/class pairs; remove racials' displayed multipliers before comparing base values. Repeat at level 60 when accessible. |
| **T29** | Both classes have client pet-scaling auras and Blizzard confirms owner-linked pet hit, but the export does not encode the coefficients. The shared model uses the Hunter wiki's 10% highest owner offensive power, full owner Crit and Hit, 30% Armor and two pet Health per owner Stamina for both classes. Warlock owner spell power feeds pet AP and spell power provisionally at 10%; its client aura also carries unsolved Intellect/Mana fields. Client Focus regenerates at 10/second and the wiki reports continuous delivery. Forever Hunter pet hit damage is held constant across pet attack speeds, consistent with the wiki and a separate Discord report; its two-second reference AP contribution remains an assumption. | At level 20, log pet stats/outcomes before and after changing owner **one stat at a time** (melee AP, ranged AP, Crit, hit, Stamina, Armor; Warlock spell power and Intellect), with family/level/happiness fixed. Check live updates against resummoning. Spend Focus and capture power every 0.1 seconds for 10 seconds; verify the scaling aura's 150-Focus cap. Try a Blessing, scroll/Juju and owner-only buff on both classes; check pet aura/stats. Compare 1.2- and 2.0-speed same-family pets at two owner AP values to refine the AP-added part of a hit. Repeat coefficients at 60; these checks refine the shared model, not gate Warlock inclusion. |
| **T12** | Any server-side downranking penalty beyond the client coefficients. Rank-2 Smite contributes 16–19% of Smite DPS; rank-2 Lightning Bolt contributes about 17% of Stormcaller DPS. | At level 20, compare the same low rank at two known SP totals, using noncritical/unresisted hits. Client coefficients are .571 for rank 2 and .714 for ranks 3/4 of Lightning Bolt; Smite ranks 2/3 are .571/.714. Record damage talents. |
| **T49** | Ice Lance's missing SP coefficient and Fingers of Frost impact/charge timing. Lance contributes about 27–28% of Frost and 18% of Arcane–Frost DPS. | At level 20, compare unfrozen hits at two SP totals, then frozen hits. The modeled .143 coefficient is an assumption. Later test FoF with an already airborne Frostbolt; repeat higher ranks. |
| **T15 / T16 / T35** | Haste-to-Energy scaling and avoided-attack refunds. PowerType in build 69977 still says 10/sec. The supplied low-level Rogue log contains too few unsaturated snapshots to distinguish 10 from 10.1/sec at 1% haste: one 4.728-second segment, corrected for a 45-Energy cast, gives 48 points gained (~10.15/sec), with at least one-point quantization uncertainty. Other subsegments disagree, so this is not a new measured rate. General-haste scaling, 80% builder refunds and zero finisher refunds remain assumptions. Berserking's separate attack/casting-speed auras (319/140/65) no longer count as Energy haste; Wind Blessed and applicable Tier 1 use the different 342/65 combination, whose Energy effect is still unverified. | At level 20, log at least a minute of unsaturated Energy with and without Wind Blessed or another named general-haste change; avoid unlogged resource procs and compare actual power snapshots. Test Berserking separately to check the distinction, and do not count attack-speed-only Slice and Dice. Separately test miss/dodge costs, subtracting natural regeneration. Gnomes should include a discounted cast. |
| **T40 / T17** | The pinned client separates Berserking's melee/ranged/casting-speed auras (319/140/65) from Wind Blessed's and haste Tier 1's combined melee/ranged haste plus casting-speed auras (342/65). The supplied Mage combat log independently shows approximately one-second Arcane Missile event spacing both with and without Berserking, supporting the no-channel-speed distinction for that effect. The model now separates casting-speed auras from the still-provisional channel response to the separately recorded SpellHaste stat. The client gives Earth, Flame and Frost Shock the same 1.5-second GCD and category; an inherited Flame Shock `IgnoreHaste` exception was removed. Other inherited exceptions on nonphysical spells were corrected with the same **spell-school** rule that WoWSims Forever uses: Holy Strike and Arcane Shot qualify despite melee/ranged hit tables; physical Stormstrike, Sunder Armor and Mortal Strike do not. The client does **not** establish the server-side GCD or channel response to each haste effect; the school distinction, one-second floor, Mind Flay, Evocation and later Penance channel response need targeted checks. | At level 20, measure hardcast, GCD and Mind Flay channels with and without an identified general-haste effect such as Wind Blessed or actual haste rating; the local log already checks Berserking versus Arcane Missiles, so do not repeat that control by default. Compare Troll Warrior Sunder Armor (1.5-second physical GCD) against Troll Shaman Earth and Flame Shock under the same Berserking effect; compare Paladin Holy Strike or Hunter Arcane Shot as nonphysical-school abilities that use physical hit tables. Record the GCD via client API or sufficient repeated casts rather than assuming `StartRecoveryTime` encodes haste scaling. Evocation/Penance require later access. |
| **T18** | Whether Cat damage inherits weapon DPS. The current model uses the Classic form weapon. | At level 20, compare two plain weapons with substantially different DPS, holding attributes and target fixed. Record Cat AP and normal white-hit damage. |
| **T25** | Omen's effective proc chance after its ten-second cooldown. Charges, duration and Wrath/free-action exclusions are already sourced. | At level 20, log several minutes of eligible attacks, promptly spending each proc. Does the first eligible event after each cooldown always proc, as the current literal 100% client model predicts? Do not spend time retesting the known free-action exclusions. |
| **T03** | Eureka's channel/DoT damage after its last charge expires. The current 10% cost discount and three-charge limit are sourced. | On a Gnome, use a channel as the third charged spell; compare ticks with an unbuffed channel. Then compare a charged DoT before/after aura expiry. Utility, healing and wand tests are not required for the current profiles. |

## Later-access, build-specific checks

| ID | Remaining DPS question | Access / focused observation |
|---|---|---|
| **T14** | Hawk guardian damage, scaling and overlapping summons. The approximation contributes roughly 8% of BM/Pet-Melee DPS. | Summon Hawk talent: record its individual attacks, lifetime, owner-stat changes and a second summon. Do not assume permanent-pet inheritance also applies to a guardian. |
| **T21** | Maelstrom's actual proc rate and its item/free-cast interactions. In build 69977, talent 408498 still has a 20% cast/cost effect per stack at rank five, five stacks and a proc trigger mask, but no linked PPM record. Its dummy effect `50` does **not** by itself prove a 50% proc chance. The current 2 PPM per talent point remains an assumption. | With the talent at level 30, compare two weapon speeds, counting landed white swings and Windfury extras separately; check Clearcasting around five-stack LB. Totem of the Storm needs later access. The 408505 cast/cost masks name Lightning Bolt, **not** Chain Lightning. |
| **T28 / T30** | Windfury/Seal/Wisdom trigger chains and ICDs—not the already explained combined white/Seal damage increase. | Log which attack triggers each extra attack, seal event and mana return; separate Command from Righteousness. Test weapon and totem separately. [Current on/off accounting](windfury.md). Consecration positioning is excluded. |
| **T26** | Whether Moonfury and Improved Moonfire multiply the SP contribution. Their rank amounts are now settled; the engine still uses base-only damage bonuses. | Compare normal hits/ticks at two SP totals with and without the talent. Improved Moonfire is available early; Moonfury later. The percent-effect records favor a broader reading, but do not by themselves reproduce the server's calculation order. |
| **T08** | Incinerate's additional 25% with Immolate: whole hit or base damage only? The implementation increases only base damage. Incinerate supplies about a third of Destruction DPS. | With Incinerate, compare two SP totals against otherwise identical targets with/without Immolate. The client dummy value establishes 25%, not the script's arithmetic. |
| **T39** | Demonic Brand's owner/pet power attribution, school and trigger eligibility. | With Brand, vary Fire/Shadow power separately and compare pet families and attacks. Preserve child spell IDs and proc counts. Threat measurements are unnecessary. |
| **T31** | Predator's Edge crit scope and Expose Prey trigger eligibility. | With those talents, separate normal/critical autos and melee specials; distinguish landed melee, ranged and periodic events on the marked target. |
| **T38** | Mutilate's two-hand/Cold Blood resolution and off-hand flat bonus. | With Mutilate/Cold Blood, use distinguishable daggers; separate both hands and an avoided main-hand strike. |
| **T53** | Twist of Light's “next melee attack” trigger and coexistence of different Echoes. Client 1311703/1311704 give Command and Righteousness separate one-charge auras; [WoWSims consumes them on landed white hits](https://github.com/wowsims/forever/blob/08dfb6be22a0/sim/paladin/talents_retribution.go#L382-L421). Treating a melee-classified Judgement as eligible changed Ret DPS substantially. | At level 40 with Twist of Light, bank each Echo before the next white swing, then cast a Judgement or Holy Strike. Record whether both remain and which event consumes them. Repeat with a missed or avoided white swing if possible. |
| **T54** | Flurry's charge timing and rank-dependent haste. The client applies a three-charge aura with no recorded proc lockout, and [WoWSims spends charges on every landed white swing](https://github.com/wowsims/forever/blob/08dfb6be22a0/sim/warrior/talents_fury.go#L150-L191). Its triggered buff says 30% speed while the talent ranks say 5–25%. We use the talent amount. | With Flurry at a reachable level, log simultaneous or closely spaced main/off-hand swings and the three charges. Compare attack speed at different talent ranks; repeat at 60 before treating the haste scaling as settled. |
| **T52** | The Forever model now recalculates ordinary periodic damage and Deep Wounds each tick, following the reported dynamic rule; exact server behavior and Deep Wounds' main-hand choice remain unverified. | At 20, apply Rend and another accessible DoT, then gain/lose a measured AP or damage buff mid-bleed without reapplying; compare normal ticks on the same target. Test Deep Wounds later with a weapon/AP change after the triggering crit. |
| **T55** | The client confirms 59 mana for rank-3 Judgement of Wisdom, but does not establish its effective server-script proc chance or which separate attacks can each trigger it. The 50% eligible-event model can return 23–37 mana/sec in prior caster/Ret profiles because the raid debuff is forced on continuously. | When Judgement is available, record at least 100 landed direct spells and 100 white swings against a judged target, separately counting actual 59-mana returns, misses and periodic/proc hits. Avoid Blessing/Seal of Wisdom and other returns. Recheck rank-3 amount and proc rate later. |
| **T56** | Deadly Poison V's level-60 effect is in the client, but its learning source in Forever is not proven; the Classic handbook came from an unavailable raid. | At level 60, confirm a trainer, vendor, quest or obtainable book teaches spell 25347; otherwise the rank-4 poison must replace it in profiles. Check whether high-frequency poison setups exhaust charges during a five-minute fight. |

## Lower-priority or conditional follow-ups

These are not prerequisites for rechecking every current row.

- **T11:** a partial final Insect Swarm tick under Tier 1. Swarm supplies about
  14% of Balance DPS, but the disputed part is only its final fractional tick.
- **T41:** Nature's Grace proc timing at Wrath completion versus projectile
  impact. Its 10% cast/GCD effects and three-second duration are sourced.
- **T04:** High Order's ordinary 15-second Read Ley Line regeneration scope.
  It is not used in the ranking APLs, but could matter to a resource-limited
  alternative. Test away from a ley line; do not grant the 15-minute variant.
- **T07 / T10:** Wrack cancellation/DoT scope and the deeper Frostfire/Tier 1
  proc interactions. Neither Wrack nor Frostfire Bolt is cast by the current
  ranked profiles. Keep these for those alternative builds, not as present
  ranking blockers.
- **T32:** Inner Focus DoT snapshotting and Starshards channel behavior when
  evaluating a rotation that actually uses them. No current profile casts
  Starshards. Incoming-damage pushback is outside this encounter.
- **T37:** Death's unexplained script value 150 could affect damage; retain
  that narrow question. Its four ranks, .429 coefficient, cooldown and stated
  backlash are implemented, and the current Shadow profile now spends two
  points on Early Demise. Absorb/healer-survival tests are not part of this
  DPS encounter.
- **T42:** Spellstone's Fire+Shadow effect mask versus its Shadow-only wording.
  No current profile equips either Warlock stone; reopen before comparing them.
- **T50:** new, affirmative trainer/quest/tome evidence for the spellbook
  watch list in [spell coverage](spell_coverage.md). Retained rune/NPC records
  are not missing damage spells simply because a tooltip exists.

### T44 — Flat mana regeneration rate

The default remains MP5 / 5 per second. The optional fivefold interpretation is
unverified, not a correction established by OOM results. With Mage Armor and
the current builds, average recorded mana-limited time is **0.06s Arcane,
2.97s Fire, 0.13s Frost and 0s Arcane–Frost** over 300 seconds. This is much
less urgent than the historical Mage results suggested.

If testing the option, compare 60 seconds with/without one flat MP5 effect,
holding Spirit and other returns fixed. An advertised `X` per five seconds
means an extra `12 × X`, versus `60 × X` under the alternate. Separate active
casting from five-second-rule recovery; exclude potions and direct procs.

## Not on the raid-DPS test list

Consecration geometry/multi-target ordering, Thick Hide defense, threat
attribution, healing/absorb survival and tank shield-break mana are archived in
the register. The old uncertain PvP armor IDs are not equipped in the current
174 profiles. Acquisition questions remain catalog work, not damage tests.

Prowl/Pounce/Ravage are acknowledged implementation gaps, not a request to
prove that ordinary stealth exists. Sanctity Aura, Fel Armor and Hammer of the
Righteous remain excluded pending an actual current acquisition path.

The adopted auto/swing rules, no Shaman dual wield, shared crit aura, Fire Nova
prerequisite and removed Lightning Shield shock-reset callback remain in place.
The five-stack LB swing reset is still a disclosed modeling rule, not newly
claimed gameplay evidence. See [auto attacks](auto_attack_audit.md).
