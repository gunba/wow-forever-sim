# Forever simulator: beta test procedures

Generated 2026-10-11 from [`uncertainties.json`](uncertainties.json). Each section below explains how to settle one open simulator question with an in-game test. Resolved questions, and questions that do not change any simulated DPS or tank result, are left out; the [full register](uncertainties.md) lists them.

There are 65 questions: 36 can be tested in the current level-30 beta, 14 can be tested in part now (low ranks) and finished at launch, 10 need level 31+ or later content, and 5 need code or a modelling decision rather than a game test.

## How to collect data

- Turn on *Advanced Combat Logging* (System › Network), then type `/combatlog` before the test and again after it. The file is `Logs/WoWCombatLog.txt` in the game folder. Upload it to foreverlogs.gg or share the raw file.
- Test on one mob type, at or below your level unless the test says otherwise, so armor, level and avoidance stay the same between runs.
- Change one thing at a time. Remove buffs, consumables, procs and talents that are not part of the test.
- For coefficients, compare **non-crit** hits from two gear sets that differ only in the stat under test. Coefficient = (average damage with set B − average with set A) ÷ (stat B − stat A).
- For proc chances, count procs and eligible events. Report both numbers, not just the percentage.
- For threat, use a threat meter that reads the game's own threat values (not a damage-based estimate).
- Record level, talents (export string), gear, race and the mob's name and level with every result.

## Summary

| ID | Question | Affects | When | State |
|---|---|---|---|---|
| [WAR-002](#war-002) | Off-hand and extra-swing rage | DPS rankings | Now | Open |
| [WAR-006](#war-006) | Flurry amount and charge events | DPS rankings | Now | Open |
| [WAR-003](#war-003) | Incoming rage and mitigation | Tank results | Now | Working answer |
| [WAR-018](#war-018) | Demoralizing Shout and Roar flat threat | Tank results | Now | Working answer |
| [WAR-010](#war-010) | Weaponmaster hand-specific armor ignore | Small | Now | Open |
| [PAL-001](#pal-001) | Seal of Righteousness spell-power formula | DPS rankings | Now | Open |
| [PAL-005](#pal-005) | Seal of Fury and its judgement | Tank results | Now | Working answer |
| [PAL-006](#pal-006) | Seal of Fury shield replacement and mana | Tank results | Now | Open |
| [PAL-008](#pal-008) | Vindication's target AP reduction | Tank results | Now | Open |
| [PAL-010](#pal-010) | Retribution Aura provider scaling | Tank results | Now | Open |
| [PAL-004](#pal-004) | Holy melee partial resistance | Small | Now | Open |
| [HUN-001](#hun-001) | Pet owner-stat inheritance | DPS rankings | Now | Open |
| [HUN-002](#hun-002) | Pet per-hit damage versus speed | DPS rankings | Now | Open |
| [HUN-004](#hun-004) | Summon Hawk guardian model | DPS rankings | Now | Working answer |
| [HUN-003](#hun-003) | Pet external buffs and action cadence | Small | Now | Open |
| [HUN-012](#hun-012) | Savage Strikes article/client version conflict | Small | Now | Open |
| [ROG-001](#rog-001) | Mutilate off-hand flat damage and Cold Blood | DPS rankings | Now | Open |
| [ROG-002](#rog-002) | Rogue finisher and poison scaling | DPS rankings | Now | Open |
| [ROG-005](#rog-005) | Hack and Slash hand scope and ICD | DPS rankings | Now | Open |
| [PRI-003](#pri-003) | Inner Focus periodic crit behavior | Small | Now | Open |
| [SHA-001](#sha-001) | Maelstrom Weapon proc rate | DPS rankings | Now | Working answer |
| [MAG-005](#mag-005) | Clearcasting and Missile Barrage event ordering | DPS rankings | Now | Working answer |
| [WL-001](#wl-001) | Demonic Brand power source and threat | DPS rankings | Now | Open |
| [WL-006](#wl-006) | Improved Imp's extra dummy effect | DPS rankings | Now | Open |
| [WL-009](#wl-009) | Demonic Knowledge ranked values | DPS rankings | Now | Open |
| [DRU-001](#dru-001) | Omen of Clarity proc frequency | DPS rankings | Now | Working answer |
| [DRU-003](#dru-003) | Form weapon damage and flat imbues | DPS rankings | Now | Open |
| [DRU-005](#dru-005) | Bear threat multipliers | Tank results | Now | Open |
| [DRU-012](#dru-012) | Bear outgoing and incoming rage formulas | Tank results | Now | Open |
| [DRU-008](#dru-008) | Nature's Grace GCD stacking and timing | Small | Now | Open |
| [CORE-003](#core-003) | Energy refunds on avoided attacks | DPS rankings | Now | Working answer |
| [CORE-006](#core-006) | Optional MP5-per-second interpretation | DPS rankings | Now | Working answer |
| [CORE-016](#core-016) | Common player Spirit regeneration | DPS rankings | Now | Working answer |
| [CORE-009](#core-009) | Eureka cast and periodic bonuses | Small | Now | Open |
| [CORE-015](#core-015) | Weapon-racial crit and ranged attacks | Small | Now | Open |
| [DATA-001](#data-001) | Server downranking penalties | DPS rankings | Now | Open |
| [WAR-020](#war-020) | Rend and Thunder Clap Attack Power scaling and Thunder Clap threat | DPS rankings | Partly now | Working answer |
| [WAR-009](#war-009) | Sunder Armor AP threat coefficient | Tank results | Partly now | Working answer |
| [WAR-021](#war-021) | Revenge Attack Power scaling and threat | Tank results | Partly now | Working answer |
| [HUN-005](#hun-005) | Arcane Shot, Serpent Sting and Volley SP | DPS rankings | Partly now | Working answer |
| [MAG-001](#mag-001) | Ice Lance coefficient and frozen multiplier | DPS rankings | Partly now | Working answer |
| [MAG-004](#mag-004) | Hot Streak interruption and new-charge ordering | Small | Partly now | Open |
| [DRU-011](#dru-011) | Proc enchants in animal forms | DPS rankings | Partly now | Open |
| [DRU-006](#dru-006) | Lacerate direct and periodic formula | Tank results | Partly now | Open |
| [CORE-001](#core-001) | Boss attack tables | DPS rankings | Partly now | Open |
| [CORE-011](#core-011) | Base attributes and class conversions | DPS rankings | Partly now | Open |
| [CORE-002](#core-002) | Energy regeneration and haste | Small | Partly now | Open |
| [CORE-004](#core-004) | Spell GCD versus physical GCD | Small | Partly now | Model choice |
| [DATA-006](#data-006) | Legacy item and enchant proc assumptions | Small | Partly now | Open |
| [SCEN-005](#scen-005) | General absorb-pool and shielding threat | Tank results | Partly now | Open |
| [WAR-001](#war-001) | Level-60 outgoing rage normalization | DPS rankings | Launch | Open |
| [WAR-008](#war-008) | Shield attacks and weapon procs | Tank results | Launch | Open |
| [PAL-002](#pal-002) | Seal Echo coexistence and consumption | DPS rankings | Launch | Open |
| [PAL-003](#pal-003) | Windfury Totem proc and charge rules | DPS rankings | Launch | Open |
| [HUN-011](#hun-011) | Hunter’s Mark rank-value anomaly | Small | Launch | Open |
| [ROG-004](#rog-004) | Deadly Poison V acquisition | Small | Launch | Open |
| [PRI-001](#pri-001) | Shadow Word: Death script value 150 | DPS rankings | Launch | Open |
| [WL-002](#wl-002) | Incinerate's Immolate bonus scope | DPS rankings | Launch | Open |
| [WL-004](#wl-004) | Spellstone's Fire/Shadow mask | Small | Launch | Working answer |
| [CORE-007](#core-007) | Judgement of Wisdom trigger rules | DPS rankings | Launch | Open |
| [SHA-004](#sha-004) | Owned buff-totem effects and replacement | DPS rankings | No game test | Code gap |
| [DRU-010](#dru-010) | Feral stealth openers | Small | No game test | Code gap |
| [DATA-002](#data-002) | Projected gear budget and stat prices | DPS rankings | No game test | Model choice |
| [SCEN-001](#scen-001) | Tank boss damage and external healing | Tank results | No game test | Model choice |
| [SCEN-003](#scen-003) | Tank target selection and threat coverage | Tank results | No game test | Code gap |

## Quick tests for a level-30 Warrior

Each of these takes 10–20 minutes on a Warrior at level 28–30 and answers a question that changes Warrior results. Level 30 gives 26 talent points, so any talent in the first six rows of a tree can be tested now; the sixth row needs 25 points in that tree first.

- **[WAR-006](#war-006)** Flurry charges (highest value: about 7% of Fury DPS depends on it): move to 26 Fury points with 1/5 Flurry (Fury 25053005040010100). Auto-attack for three minutes, then count white swings between each Flurry gain and fade in the combat log.
- **[WAR-002](#war-002)** Off-hand rage: dual wield, auto-attack only for 3 minutes, Dual Wield Specialization 0/5 then 5/5. Note the Rage change after each off-hand swing.
- **[WAR-021](#war-021)** Revenge AP slope: tank one mob with two gear sets that differ only in Attack Power. Cast Revenge after every block, dodge or parry. 50 non-crit Revenges per set.
- **[WAR-020](#war-020)** Rend and Thunder Clap AP slope: the same two gear sets, Improved Rend 0/3. 50 Rend ticks and 30 Thunder Claps per set.
- **[WAR-009](#war-009)** Sunder threat: the same two gear sets, five Sunders on a fresh mob, threat meter on. 20 Sunders per set.
- **[WAR-018](#war-018)** Demoralizing Shout threat: Battle Stance, one mob then three mobs, threat meter on. 10 casts each.
- **[WAR-003](#war-003)** Incoming rage: stand still with auto-attack off and let one mob hit you for two minutes, with two armor sets of equal Stamina.
- **[CORE-001](#core-001)** Attack table: 2,000 white swings from behind a level-33 dungeon elite; count misses, dodges, glancing blows and crits.

## Testable now (level-30 beta)

### Warrior

<a id="war-002"></a>
#### WAR-002: Off-hand and extra-swing rage

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Off-hand and critical-swing Rage drive every dual-wield Warrior's Heroic Strike budget.
- **What the simulator assumes now:** Off-hand swings give half the main-hand rate (1.73 Rage per weapon-second), raised 10% per Dual Wield Specialization point; a white crit adds a second share; misses, dodges and parries give nothing.
- **Setup:** Dual-wield Warrior, any level up to 30. Two runs: Dual Wield Specialization 0/5 and 5/5. No Unbridled Wrath, Bloodrage or Berserker Rage. A non-casting mob at or below your level.
- **Steps:**
  1. Start below 50 Rage and only auto-attack (no abilities) for two to three minutes per run.
  2. Repeat with the second talent setting.
  3. Keep the same two weapons in the same hands for both runs.
- **Record:** Each off-hand swing's outcome and the Rage change that follows it (advanced combat log 'current power' field).
- **Sample size:** At least 150 landed off-hand swings per run, including about 20 crits.
- **How the result is used:** Sets the off-hand rate, the DWS multiplier and whether off-hand crits add the extra share.

<a id="war-006"></a>
#### WAR-006: Flurry amount and charge events

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Flurry is the main Fury haste source and drives rage income, because each swing gives fixed Rage. If Warrior Flurry lasts one swing, as Shaman Flurry does in the logs (SHA-009), Fury and Fury (Sunder) lose about 7% and 2H Bloodthirst about 6%.
- **What the simulator assumes now:** A melee crit gives 3 charges of 5-25% attack speed; every white swing (including avoided ones and queued Heroic Strikes) spends a charge; charges are shared between hands.
- **Setup:** Level-30 Warrior with 26 Fury points: 25 in rows 1-5 including Death Wish, plus 1/5 Flurry (for example Fury 25053005040010100, with no Arms or Protection points). Advanced combat logging on.
- **Steps:**
  1. Auto-attack a mob at or below your level for three minutes with one weapon, using no abilities.
  2. Repeat while dual wielding.
  3. Repeat while queueing Heroic Strike on every main-hand swing.
  4. During each run, glance at the Flurry buff icon after a crit and note whether it shows a charge count of 3.
- **Record:** Each Flurry gain and fade in the combat log, and every white swing (hit, crit, miss, dodge, parry) between them.
- **Sample size:** About 60 Flurry procs per setup (about three minutes each).
- **How the result is used:** Swings per Flurry window give the charge count: almost always 1 means one charge, as for Shamans; up to 3 means the current model is right. The dual-wield and Heroic Strike runs show what spends a charge.

<a id="war-003"></a>
#### WAR-003: Incoming rage and mitigation

*Affects: Tank results. State: Working answer.*

- **Why it matters:** Protection Warrior Rage from damage taken.
- **What the simulator assumes now:** Rage = (damage before armor + absorbed) x 10 / level-based expected health / (1 - reference armor).
- **Setup:** Warrior tank at level 30. Two armor sets with the same Stamina; two Stamina sets with the same armor.
- **Steps:**
  1. Stand still with auto-attack off and let one mob hit you for two minutes per set.
  2. Repeat with a shield absorb active (e.g. Power Word: Shield) for one set.
- **Record:** Each incoming hit's raw and final damage, absorbed amount, and the Rage change.
- **Sample size:** 100 hits per set.
- **How the result is used:** Separates the armor, maximum-health and absorb terms of the formula.

<a id="war-018"></a>
#### WAR-018: Demoralizing Shout and Roar flat threat

*Affects: Tank results. State: Working answer.*

- **Why it matters:** Protection threat from Demoralizing Shout.
- **What the simulator assumes now:** 43.2 threat per target at rank 5 (Classic-derived).
- **Setup:** Warrior at level 30 (rank 2-3), threat meter addon.
- **Steps:** Cast Demoralizing Shout on one mob, then on three mobs, in Battle Stance (no Defiance).
- **Record:** Threat added per affected mob.
- **Sample size:** 10 casts per setup.
- **How the result is used:** Gives the per-rank flat value and whether it is split between targets.

<a id="war-010"></a>
#### WAR-010: Weaponmaster hand-specific armor ignore

*Affects: Small. State: Open.*

- **Why it matters:** Weaponmaster's armor ignore with a mace.
- **What the simulator assumes now:** A mace in either hand grants the character-wide armor ignore.
- **Setup:** Warrior with 5/5 Weaponmaster (20 Arms points first; 25 of the 26 points available at level 30), one mace and one non-mace one-hander.
- **Steps:** Hit the same mob type with the mace in the main hand, then in the off hand.
- **Record:** Main-hand non-crit damage in each layout.
- **Sample size:** 100 hits per layout.
- **How the result is used:** If main-hand damage rises only with the mace in the main hand, the effect is hand-specific.

### Paladin

<a id="pal-001"></a>
#### PAL-001: Seal of Righteousness spell-power formula

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Seal of Righteousness damage for Retribution.
- **What the simulator assumes now:** Spell power is added twice and two-handers get a 1.1 factor (inherited).
- **Setup:** Paladin with Seal of Righteousness, two spell power sets, one-hander then two-hander.
- **Steps:** Auto-attack a mob with the seal active, no Judgement of the Crusader.
- **Record:** Seal proc damage (non-crit).
- **Sample size:** 100 procs per setting.
- **How the result is used:** Fixes the SP coefficient and the two-hand factor.

<a id="pal-005"></a>
#### PAL-005: Seal of Fury and its judgement

*Affects: Tank results. State: Working answer.*

- **Why it matters:** Protection Paladin Seal of Fury.
- **What the simulator assumes now:** Damage, judgement and shield as implemented; one shared shield pool.
- **Setup:** Protection Paladin with Seal of Fury.
- **Steps:** Tank one mob and record seal procs and shield absorbs.
- **Record:** Proc damage, shield size, absorbs.
- **Sample size:** 100 procs.
- **How the result is used:** Validates damage and shield size.

<a id="pal-006"></a>
#### PAL-006: Seal of Fury shield replacement and mana

*Affects: Tank results. State: Open.*

- **Why it matters:** Seal of Fury shield behaviour.
- **What the simulator assumes now:** New procs replace the shield; 30-second lifetime.
- **Setup:** Protection Paladin with Seal of Fury and Improved Seal of Fury.
- **Steps:** Let several procs happen before the shield is used up; then let it expire.
- **Record:** Shield size after each proc, absorbs, expiry, mana returns.
- **Sample size:** 50 procs.
- **How the result is used:** Replace versus add, lifetime, mana trigger.

<a id="pal-008"></a>
#### PAL-008: Vindication's target AP reduction

*Affects: Tank results. State: Open.*

- **Why it matters:** Vindication's enemy AP reduction.
- **What the simulator assumes now:** 68/136/204 AP at level 60.
- **Setup:** Paladin with Vindication, a mob whose hit size is stable.
- **Steps:** Compare the mob's hits on you with and without the debuff.
- **Record:** Incoming damage.
- **Sample size:** 100 hits each.
- **How the result is used:** Confirms the AP reduction.

<a id="pal-010"></a>
#### PAL-010: Retribution Aura provider scaling

*Affects: Tank results. State: Open.*

- **Why it matters:** Retribution Aura damage on tanks.
- **What the simulator assumes now:** Aura damage + 0.06 x the provider's spell power.
- **Setup:** Paladin providing the aura, two spell power sets.
- **Steps:** Let a mob hit an aura-protected player.
- **Record:** Aura damage per hit.
- **Sample size:** 50 hits per set.
- **How the result is used:** Confirms the coefficient.

<a id="pal-004"></a>
#### PAL-004: Holy melee partial resistance

*Affects: Small. State: Open.*

- **Why it matters:** Partial resists on Holy melee damage.
- **What the simulator assumes now:** Holy melee damage can be partially resisted.
- **Setup:** Paladin with Seal of Command.
- **Steps:** Fight mobs and inspect seal damage events.
- **Record:** Resisted amounts on Holy events.
- **Sample size:** 200 events.
- **How the result is used:** Removes partial resists if none appear.

### Hunter and pets

<a id="hun-001"></a>
#### HUN-001: Pet owner-stat inheritance

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Pet damage for Beast Mastery, Survival and Warlock rows.
- **What the simulator assumes now:** Pets get 10% of the owner's higher melee or ranged AP, plus hit, crit, 2x Stamina and 30% armor.
- **Setup:** Level 30 Hunter and Warlock with pets; gear sets that change one owner stat at a time.
- **Steps:**
  1. Read the pet's stats (pet tab) after each gear change, then dismiss and re-summon and read again.
  2. Fight a mob and compare pet hit damage.
- **Record:** Pet AP, spell power, crit, health and armor; pet damage per hit.
- **Sample size:** One reading per stat change; 100 pet hits per AP setting.
- **How the result is used:** Sets the transfer coefficients and whether they update live or only on summon.

<a id="hun-002"></a>
#### HUN-002: Pet per-hit damage versus speed

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Beast Mastery pet damage depends on per-hit damage at the pet's speed.
- **What the simulator assumes now:** Every pet family hits equally hard per hit; faster pets therefore do more damage per second.
- **Setup:** Two pets of different base attack speed, same owner gear.
- **Steps:** Let each pet auto-attack the same mob type; repeat with higher owner AP.
- **Record:** Pet white-hit damage (non-crit) and swing interval.
- **Sample size:** 200 hits per pet per AP setting.
- **How the result is used:** Shows whether damage per hit scales with speed (normalized) and the AP-to-damage conversion.

<a id="hun-004"></a>
#### HUN-004: Summon Hawk guardian model

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Summon Hawk damage for the Pet/Melee and Beast Mastery options.
- **What the simulator assumes now:** Two 18-second Hawks; opening dive always hits about 1 s after the cast; assault attacks use the full melee table; 1.5% base crit; Ferocity transfer unknown (logs show 5-8% crit for Beast Mastery Hunters of unknown Ferocity rank).
- **Setup:** Hunter with Summon Hawk (16 Beast Mastery points); one run with 0/5 Ferocity and one with Ferocity points, same gear.
- **Steps:**
  1. Summon the Hawk repeatedly on the same mob type in each talent setup.
  2. Repeat one setup with a second ranged AP gear set.
- **Record:** Hawk opening hit and later attack damage, crits, despawns.
- **Sample size:** 100 Hawk attacks per set.
- **How the result is used:** Fixes RAP transfer and crit; the third summon shows replacement.

<a id="hun-003"></a>
#### HUN-003: Pet external buffs and action cadence

*Affects: Small. State: Open.*

- **Why it matters:** How quickly pets use their abilities.
- **What the simulator assumes now:** Hunter pet abilities are used as soon as they are ready and affordable. (Warlock Imp pause measured at 432 ms.)
- **Setup:** Hunter pet with Bite and Claw, at full Focus and when Focus-starved.
- **Steps:** Fight a mob for two minutes in each state.
- **Record:** Pet ability timestamps and Focus.
- **Sample size:** 50 casts per ability.
- **How the result is used:** Sets the delay between an ability becoming ready and the pet using it.

<a id="hun-012"></a>
#### HUN-012: Savage Strikes article/client version conflict

*Affects: Small. State: Open.*

- **Why it matters:** Survival Savage Strikes.
- **What the simulator assumes now:** Adds crit chance to Raptor Strike, Mongoose Bite and similar abilities (client data), not melee damage (article).
- **Setup:** Survival Hunter with 2/2 Savage Strikes.
- **Steps:** Use Raptor Strike and Mongoose Bite on mobs with and without the talent.
- **Record:** Crit rate and non-crit damage of those abilities.
- **Sample size:** 100 uses each setting.
- **How the result is used:** Decides crit versus damage.

### Rogue

<a id="rog-001"></a>
#### ROG-001: Mutilate off-hand flat damage and Cold Blood

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Mutilate's off-hand strike damage in the Mutilate row.
- **What the simulator assumes now:** Both strikes share one avoidance roll; the flat bonus is applied to both hands as in the current code.
- **Setup:** Rogue with Mutilate (20 Assassination points, reachable at 30), two daggers, two Attack Power sets.
- **Steps:**
  1. Cast Mutilate repeatedly on the same mob type with each AP set.
  2. Swap the daggers between hands and repeat.
- **Record:** Main-hand and off-hand strike damage (non-crit), with Cold Blood on and off.
- **Sample size:** 40 casts per setup.
- **How the result is used:** The off-hand intercept shows whether the flat bonus is halved with the off-hand penalty.

<a id="rog-002"></a>
#### ROG-002: Rogue finisher and poison scaling

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Poison and finisher damage for all Rogue rows.
- **What the simulator assumes now:** Instant Poison 0.5% AP and Deadly Poison 0.45% AP per the developer reply; finishers keep Classic AP terms.
- **Setup:** Level 30 Rogue with Deadly Poison (trained at 30) and Instant Poison, two Attack Power sets.
- **Steps:**
  1. Fight the same mob type with each AP set, keeping Deadly Poison at one stack, then at five stacks.
  2. Use Eviscerate and Rupture at 1 and 5 combo points.
- **Record:** Poison hit and tick damage, finisher damage and combo points.
- **Sample size:** 100 poison hits/ticks and 20 finishers per setting.
- **How the result is used:** Fixes the poison per-stack placement and the finisher AP terms.

<a id="rog-005"></a>
#### ROG-005: Hack and Slash hand scope and ICD

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Combat Rogue extra attacks from Hack and Slash.
- **What the simulator assumes now:** 200 ms cooldown (public logs show no two extra attacks closer than 824 ms); crit and armor parts follow the current hand rules.
- **Setup:** Rogue with Hack and Slash (20 Combat points), a sword in one hand and a non-sword in the other.
- **Steps:** Auto-attack with the sword in the main hand, then in the off hand.
- **Record:** Crit rate and damage per hand, extra-attack events.
- **Sample size:** 300 swings per layout.
- **How the result is used:** Decides whether the crit/armor part applies only to the hand holding the sword.

### Priest

<a id="pri-003"></a>
#### PRI-003: Inner Focus periodic crit behavior

*Affects: Small. State: Open.*

- **Why it matters:** Inner Focus with DoTs.
- **What the simulator assumes now:** Only spells in the client's crit mask get the crit bonus.
- **Setup:** Priest with Inner Focus.
- **Steps:** Apply Shadow Word: Pain with Inner Focus consumed by it, and without it.
- **Record:** Tick crits.
- **Sample size:** 200 ticks each.
- **How the result is used:** Decides whether DoT ticks keep the crit bonus.

### Shaman

<a id="sha-001"></a>
#### SHA-001: Maelstrom Weapon proc rate

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Maelstrom Weapon drives Enhancement and Stormcaller Lightning Bolts.
- **What the simulator assumes now:** 10 procs per minute at every rank (logs: 55% per landed hit with a 3.6-second weapon at 1 point).
- **Setup:** Enhancement Shaman with 1/5 Maelstrom Weapon (25 points, reachable with 26 at level 30). One slow and one fast weapon.
- **Steps:** Melee a mob for three minutes with each weapon without casting Lightning Bolt.
- **Record:** Maelstrom stack gains and the attacks that preceded them (white, Stormstrike, Windfury).
- **Sample size:** 150 stack gains per weapon.
- **How the result is used:** If the rate per hit rises with weapon speed, it is procs per minute; if it does not, it is a flat chance.

### Mage

<a id="mag-005"></a>
#### MAG-005: Clearcasting and Missile Barrage event ordering

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Arcane rotation resource use.
- **What the simulator assumes now:** Missile Barrage now rolls at Fireball/Frostbolt cast completion (logs); Clearcasting has a 1-second cooldown and is spent at cast completion.
- **Setup:** Arcane Mage with Arcane Concentration and Missile Barrage.
- **Steps:** Channel Arcane Missiles while Clearcasting is active, and note whether a Clearcasting proc from a missile is spent by the same channel.
- **Record:** Clearcasting gain and removal times against Arcane Missiles ticks; mana spent.
- **Sample size:** 40 channels.
- **How the result is used:** Fixes Clearcasting consumption order.

### Warlock

<a id="wl-001"></a>
#### WL-001: Demonic Brand power source and threat

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Demonic Brand damage for Demonology.
- **What the simulator assumes now:** (level-26) x 1.5 + 14-17 + 0.078 x owner school power per pet attack; can crit.
- **Setup:** Warlock with 3/3 Demonic Brand (15 points), two spell power sets.
- **Steps:** Cast Searing Pain, then let the pet attack; repeat with each set.
- **Record:** Brand damage, crits.
- **Sample size:** 60 Brand hits per set.
- **How the result is used:** Confirms the coefficient and crit.

<a id="wl-006"></a>
#### WL-006: Improved Imp's extra dummy effect

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Improved Imp's extra effect on Firebolt.
- **What the simulator assumes now:** Firebolt has a 2.0-second cast and the Imp waits about 0.43 s between casts (logs). Improved Imp is assumed not to change either.
- **Setup:** Warlock with 3/3 Improved Imp and 0/3, same gear.
- **Steps:** Let the Imp cast Firebolt continuously for two minutes each.
- **Record:** Firebolt cast start/success times.
- **Sample size:** 50 casts per setting.
- **How the result is used:** Shows whether the talent shortens the cast or the pause.

<a id="wl-009"></a>
#### WL-009: Demonic Knowledge ranked values

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Demonic Knowledge spell power.
- **What the simulator assumes now:** 33/67/100% of character level.
- **Setup:** Warlock with Demonic Knowledge (20 points) at 1/3, 2/3 and 3/3.
- **Steps:** Read owner and pet spell power at each rank, re-summoning after each change.
- **Record:** Spell power.
- **Sample size:** One reading per rank.
- **How the result is used:** Confirms the ranked values.

### Druid

<a id="dru-001"></a>
#### DRU-001: Omen of Clarity proc frequency

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Omen of Clarity frequency for Feral and Balance.
- **What the simulator assumes now:** After a 10-second cooldown: white hits at 4 procs per minute of attack speed, abilities/spells/heals at 20% (fitted from 691 logged proc windows).
- **Setup:** Druid in Cat Form, then in caster form (Balance), with no other Clearcasting source.
- **Steps:** Attack or cast continuously for five minutes in each form.
- **Record:** Clearcasting gains and every attack or cast between them.
- **Sample size:** 60 procs per form.
- **How the result is used:** Gives the chance per event after the cooldown and whether it depends on attack speed.

<a id="dru-003"></a>
#### DRU-003: Form weapon damage and flat imbues

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Feral damage from the equipped weapon.
- **What the simulator assumes now:** Cat uses a 1-second weapon and Bear a 2.5-second weapon built from the equipped weapon's DPS; stones and weapon enchants apply.
- **Setup:** Druid in Cat Form with two weapons of different DPS, same AP.
- **Steps:** Auto-attack with each weapon; add a sharpening stone to one.
- **Record:** White and Shred/Claw damage (non-crit).
- **Sample size:** 100 hits per weapon.
- **How the result is used:** Confirms the weapon-DPS and flat-bonus routing.

<a id="dru-005"></a>
#### DRU-005: Bear threat multipliers

*Affects: Tank results. State: Open.*

- **Why it matters:** Bear tank threat.
- **What the simulator assumes now:** Form 1.5x; Primal Bite 3x, Lacerate 3.33x, Swipe 2x ability multipliers.
- **Setup:** Bear Druid, threat meter addon.
- **Steps:** Use each ability on one mob and on three mobs.
- **Record:** Threat added versus damage dealt.
- **Sample size:** 20 uses per ability.
- **How the result is used:** Sets each ability multiplier.

<a id="dru-012"></a>
#### DRU-012: Bear outgoing and incoming rage formulas

*Affects: Tank results. State: Open.*

- **Why it matters:** Bear Rage income.
- **What the simulator assumes now:** Classic damage-based outgoing Rage with the Forever crit bonus.
- **Setup:** Bear Druid, two AP sets.
- **Steps:** Auto-attack only, then take hits with auto-attack off.
- **Record:** Rage per landed hit and per hit taken.
- **Sample size:** 100 each.
- **How the result is used:** Decides whether Bear uses speed-based Rage like Warriors.

<a id="dru-008"></a>
#### DRU-008: Nature's Grace GCD stacking and timing

*Affects: Small. State: Open.*

- **Why it matters:** Balance Nature's Grace timing.
- **What the simulator assumes now:** Applies 10% cast speed and -10% GCD; Wrath triggers it on cast completion.
- **Setup:** Balance Druid with Nature's Grace (20 points).
- **Steps:** Cast Wrath and Starfire at long range.
- **Record:** Nature's Grace gain time versus cast end and impact; the next GCD length.
- **Sample size:** 40 procs.
- **How the result is used:** Fixes the trigger point and GCD stacking.

### Shared mechanics

<a id="core-003"></a>
#### CORE-003: Energy refunds on avoided attacks

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Energy returned when a builder or finisher is avoided.
- **What the simulator assumes now:** 80% refund on miss, dodge or parry for builders and finishers.
- **Setup:** Level 30 Rogue, attack a mob from the front (to get parries).
- **Steps:** Use Sinister Strike and Eviscerate repeatedly below the Energy cap.
- **Record:** Energy before and after each avoided ability (advanced combat log power field), separately for misses, dodges and parries.
- **Sample size:** 20 avoided casts of each type.
- **How the result is used:** Sets the refund fraction and whether misses refund.

<a id="core-006"></a>
#### CORE-006: Optional MP5-per-second interpretation

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Mana for every caster; several rows run out of mana.
- **What the simulator assumes now:** MP5 restores mana per five seconds (the per-second reading is an optional switch).
- **Setup:** Any mana user with MP5 gear and a set without it, same Spirit.
- **Steps:** Cast continuously to stay inside the five-second rule, then stop and read mana regeneration.
- **Record:** Mana per tick with each set.
- **Sample size:** 20 ticks per set.
- **How the result is used:** Decides per-second versus per-five-seconds.

<a id="core-016"></a>
#### CORE-016: Common player Spirit regeneration

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Spirit regeneration for every caster.
- **What the simulator assumes now:** 6.25 + Spirit/8 mana per second for all classes (fitted from low-level samples).
- **Setup:** Casters of several classes at different levels, two Spirit sets.
- **Steps:** Out of the five-second rule, sit idle and read regeneration with each set.
- **Record:** Mana per tick, Spirit, level, class.
- **Sample size:** 20 ticks per set.
- **How the result is used:** Confirms the formula across classes and levels.

<a id="core-009"></a>
#### CORE-009: Eureka cast and periodic bonuses

*Affects: Small. State: Open.*

- **Why it matters:** Gnome Eureka on channels and DoTs.
- **What the simulator assumes now:** The consuming cast keeps the bonus on all its ticks; ordinary periodic damage does not get it.
- **Setup:** Gnome caster with a channel (Mind Flay, Drain Life) and a DoT.
- **Steps:** Cast each with Eureka active, without it, and as the last charge.
- **Record:** Tick damage in each case.
- **Sample size:** 30 ticks per case.
- **How the result is used:** Settles which ticks get the bonus.

<a id="core-015"></a>
#### CORE-015: Weapon-racial crit and ranged attacks

*Affects: Small. State: Open.*

- **Why it matters:** Orc/Human weapon racials on ranged attacks.
- **What the simulator assumes now:** The racial's crit also applies to ranged attacks.
- **Setup:** Orc Hunter with an axe, or Human Hunter with a sword.
- **Steps:** Compare ranged crit in the character sheet and in logged Auto Shots with and without the matching melee weapon equipped.
- **Record:** Auto Shot crit rate.
- **Sample size:** 400 Auto Shots per setting.
- **How the result is used:** Removes or keeps the ranged crit.

### Items and data

<a id="data-001"></a>
#### DATA-001: Server downranking penalties

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Lower ranks used by Smite and Elemental rows.
- **What the simulator assumes now:** No hidden downranking penalty.
- **Setup:** Caster with two learned ranks, two spell power sets.
- **Steps:** Cast each rank with each set on the same mob type.
- **Record:** Non-crit damage.
- **Sample size:** 40 casts per rank and set.
- **How the result is used:** Compares the spell-power slope of each rank.

## Partly testable now (low ranks now, top ranks at launch)

### Warrior

<a id="war-020"></a>
#### WAR-020: Rend and Thunder Clap Attack Power scaling and Thunder Clap threat

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Arms Rend upkeep and Protection Thunder Clap.
- **What the simulator assumes now:** Rend ticks add 2% of current AP; Thunder Clap adds 7% AP and its threat equals its damage.
- **Setup:** Level 30 Warrior, two gear sets differing only in Attack Power. Improved Rend 0/3.
- **Steps:** Apply Rend and Thunder Clap to the same mob type with each set.
- **Record:** Each Rend tick and Thunder Clap hit (non-crit).
- **Sample size:** 50 ticks and 30 Thunder Claps per set.
- **How the result is used:** Confirms the low-rank AP slopes; repeat at 60 for the top ranks.

<a id="war-009"></a>
#### WAR-009: Sunder Armor AP threat coefficient

*Affects: Tank results. State: Working answer.*

- **Why it matters:** Protection Warrior Sunder threat.
- **What the simulator assumes now:** Threat = client flat value + 0.05 x Attack Power.
- **Setup:** Warrior at level 30 (Sunder rank 2), two gear sets differing only in Attack Power.
- **Steps:** Apply Sunder five times on a fresh mob with each set, using a threat meter that reads UnitDetailedThreatSituation.
- **Record:** Threat added by each Sunder (subtract stance and talent multipliers).
- **Sample size:** 20 Sunders per set.
- **How the result is used:** Gives the AP slope at rank 2; repeat at 60 for rank 5.

<a id="war-021"></a>
#### WAR-021: Revenge Attack Power scaling and threat

*Affects: Tank results. State: Working answer.*

- **Why it matters:** Revenge is Protection Warrior's largest threat source; the rank-6 coefficient is the most sensitive value in the sim.
- **What the simulator assumes now:** Damage = rank roll + 0.28 x AP; threat = 2.25 x damage.
- **Setup:** Level 30 Warrior (Revenge rank 2), two gear sets differing only in Attack Power; later level 60.
- **Steps:** Tank a mob and cast Revenge after every block, dodge or parry with each set.
- **Record:** Revenge non-crit damage and threat.
- **Sample size:** 50 non-crit Revenges per set.
- **How the result is used:** Rank-2 slope now; rank 6 at launch replaces the extrapolated 0.28.

### Hunter and pets

<a id="hun-005"></a>
#### HUN-005: Arcane Shot, Serpent Sting and Volley SP

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Spell-power gear for Hunters.
- **What the simulator assumes now:** Arcane Shot: rank base + 10% RAP, no spell power. Serpent Sting: rank base + 3% current RAP per tick, no spell power. Volley untested.
- **Setup:** Level-30 Hunter with Arcane Shot and Serpent Sting; two gear sets differing only in spell power, then two differing only in ranged AP. Volley needs level 40 (launch).
- **Steps:** Cast Arcane Shot and apply Serpent Sting on the same mob type with each set.
- **Record:** Arcane Shot non-crit damage and Serpent Sting tick damage.
- **Sample size:** 40 shots and 50 ticks per set.
- **How the result is used:** Confirms zero spell-power scaling and the RAP slopes.

### Mage

<a id="mag-001"></a>
#### MAG-001: Ice Lance coefficient and frozen multiplier

*Affects: DPS rankings. State: Working answer.*

- **Why it matters:** Ice Lance is in the Frost and Arcane-Frost rotations.
- **What the simulator assumes now:** Provisional 0.10 spell-power coefficient; the frozen bonus covers the whole hit.
- **Setup:** Level-30 Mage with Ice Lance, two spell power sets.
- **Steps:** Cast Ice Lance on an unfrozen target and on a frozen target (Frost Nova or Fingers of Frost) with each set.
- **Record:** Non-crit damage in each case.
- **Sample size:** 40 casts per case and set.
- **How the result is used:** Gives the coefficient at this rank and whether the frozen multiplier applies to the spell-power part.

<a id="mag-004"></a>
#### MAG-004: Hot Streak interruption and new-charge ordering

*Affects: Small. State: Open.*

- **Why it matters:** Fire Hot Streak edge cases.
- **What the simulator assumes now:** Completed Pyroblast consumes the stacks.
- **Setup:** Fire Mage with Hot Streak.
- **Steps:** Start an instant Pyroblast while another Fire projectile is in flight.
- **Record:** Hot Streak stack events at impact and Pyroblast completion.
- **Sample size:** 30 cases.
- **How the result is used:** Decides whether the in-flight stack is consumed.

### Druid

<a id="dru-011"></a>
#### DRU-011: Proc enchants in animal forms

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Weapon enchant procs in Cat and Bear Form.
- **What the simulator assumes now:** Procs work in form at the form's weapon speed.
- **Setup:** Druid with a proc enchant (e.g. Fiery Weapon) in Cat Form and in caster form.
- **Steps:** Auto-attack for three minutes in each.
- **Record:** Proc events per hit.
- **Sample size:** 40 procs per form.
- **How the result is used:** Decides form eligibility and which speed sets the rate.

<a id="dru-006"></a>
#### DRU-006: Lacerate direct and periodic formula

*Affects: Tank results. State: Open.*

- **Why it matters:** Bear Lacerate damage.
- **What the simulator assumes now:** 10% form weapon damage per stack plus a level-based bleed per stack.
- **Setup:** Bear Druid with Lacerate.
- **Steps:** Stack Lacerate 1 to 5 times on a mob and keep it up.
- **Record:** Direct and bleed damage at each stack count.
- **Sample size:** 20 casts per stack count.
- **How the result is used:** Fixes the per-stack formula.

### Shared mechanics

<a id="core-001"></a>
#### CORE-001: Boss attack tables

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Hit, dodge, parry, crit and glancing against bosses.
- **What the simulator assumes now:** No hit suppression; Classic dodge/parry/glancing chances and glancing damage.
- **Setup:** Level 30 vs level 33 dungeon elites (current beta), level 60 vs 63 at launch.
- **Steps:** Auto-attack and use specials from behind and from the front.
- **Record:** Outcome counts; glancing damage against normal hits.
- **Sample size:** 2,000 swings per setup.
- **How the result is used:** Confirms the table; glancing damage once Blizzard fixes it.

<a id="core-011"></a>
#### CORE-011: Base attributes and class conversions

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Base attributes for every row.
- **What the simulator assumes now:** Classic race/class tables with Forever racial changes.
- **Setup:** Naked characters at level 60 (launch); level 30 now as a partial check.
- **Steps:** Remove all gear and buffs; read the character sheet.
- **Record:** Strength, Agility, Stamina, Intellect, Spirit, health, mana.
- **Sample size:** One reading per race/class.
- **How the result is used:** Replaces the tables.

<a id="core-002"></a>
#### CORE-002: Energy regeneration and haste

*Affects: Small. State: Open.*

- **Why it matters:** Energy regeneration with general haste.
- **What the simulator assumes now:** 10 Energy per second, raised by general haste.
- **Setup:** Rogue or Cat Druid, a known general-haste effect (e.g. a haste potion or racial).
- **Steps:** From low Energy, wait without attacking until near the cap, with and without the haste effect.
- **Record:** Energy over time.
- **Sample size:** 10 recovery cycles each.
- **How the result is used:** Decides whether general haste speeds Energy.

<a id="core-004"></a>
#### CORE-004: Spell GCD versus physical GCD

*Affects: Small. State: Model choice.*

- **Why it matters:** Spell GCD with haste.
- **What the simulator assumes now:** Casting speed shortens magical GCDs to a 1-second floor.
- **Setup:** Caster with a known haste effect.
- **Steps:** Chain instant spells with and without haste.
- **Record:** Time between casts.
- **Sample size:** 30 casts each.
- **How the result is used:** Confirms the rule.

### Items and data

<a id="data-006"></a>
#### DATA-006: Legacy item and enchant proc assumptions

*Affects: Small. State: Open.*

- **Why it matters:** Proc rates of legacy enchants and items.
- **What the simulator assumes now:** Classic rates.
- **Setup:** Weapons with the enchants used in profiles.
- **Steps:** Melee for five minutes per weapon speed.
- **Record:** Procs per landed hit.
- **Sample size:** 40 procs per enchant.
- **How the result is used:** Replaces assumed rates.

### Scenarios

<a id="scen-005"></a>
#### SCEN-005: General absorb-pool and shielding threat

*Affects: Tank results. State: Open.*

- **Why it matters:** Threat from absorb shields.
- **What the simulator assumes now:** Absorbs consume damage; threat rules inherited.
- **Setup:** Priest or Paladin shielding a tank, threat meter.
- **Steps:** Shield the tank and let mobs hit it.
- **Record:** Threat on shield application, absorption and expiry.
- **Sample size:** 20 shields.
- **How the result is used:** Fixes shield threat.

## Launch only (needs level 31+ or later content)

### Warrior

<a id="war-001"></a>
#### WAR-001: Level-60 outgoing rage normalization

*Affects: DPS rankings. State: Open.*

- **Why it matters:** All Warrior rows depend on Rage per swing at level 60.
- **What the simulator assumes now:** Low-level rates (3.46 Rage per weapon-second one-hand, 4.5 two-hand) are assumed unchanged at 60.
- **Setup:** Level 60 Warrior (launch). Uncapped Rage. One weapon at a time.
- **Steps:**
  1. Auto-attack with a one-hander, then a two-hander, without abilities.
  2. Repeat with a weapon of different damage but the same speed.
- **Record:** Rage gained per landed white swing.
- **Sample size:** 100 landed swings per weapon.
- **How the result is used:** Confirms or replaces the rates; a change in damage at fixed speed tests whether damage matters.

<a id="war-008"></a>
#### WAR-008: Shield attacks and weapon procs

*Affects: Tank results. State: Open.*

- **Why it matters:** Protection Warrior procs from Shield Slam and Shield Bash.
- **What the simulator assumes now:** Shield Slam is a main-hand special and can trigger weapon procs.
- **Setup:** Level 40+ Protection Warrior with a proc weapon (launch).
- **Steps:** Use Shield Slam and Shield Bash repeatedly, with no white swings in between (stop auto-attack).
- **Record:** Proc events after each Shield Slam/Bash.
- **Sample size:** 100 casts each.
- **How the result is used:** Decides whether shield attacks can trigger weapon effects.

### Paladin

<a id="pal-002"></a>
#### PAL-002: Seal Echo coexistence and consumption

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Retribution seal twisting.
- **What the simulator assumes now:** Separate one-charge Echoes per seal.
- **Setup:** Paladin with Twist of Light (31 Retribution points; launch).
- **Steps:** Swap seals before the next swing.
- **Record:** Echo auras and what spends them.
- **Sample size:** 40 swaps.
- **How the result is used:** Coexistence and consumption.

<a id="pal-003"></a>
#### PAL-003: Windfury Totem proc and charge rules

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Windfury Totem for every melee row.
- **What the simulator assumes now:** Inherited Classic proc and charge rules.
- **Setup:** Melee player with a level 32+ Shaman's Windfury Totem (launch).
- **Steps:** Melee a mob with the totem buff.
- **Record:** Windfury procs per white and special attack, AP charges, swing timing.
- **Sample size:** 200 procs.
- **How the result is used:** Proc events, cooldown and charges.

### Hunter and pets

<a id="hun-011"></a>
#### HUN-011: Hunter’s Mark rank-value anomaly

*Affects: Small. State: Open.*

- **Why it matters:** Hunter's Mark ranged AP bonus at ranks 3 and 4.
- **What the simulator assumes now:** Captured values: rank 3 = 98 RAP, rank 4 = 71 RAP (an apparent data error).
- **Setup:** Level 40 and level 58 Hunters (launch).
- **Steps:** Apply each rank and read your ranged AP against the marked target.
- **Record:** Ranged AP with and without the mark.
- **Sample size:** One reading per rank.
- **How the result is used:** Replaces the anomalous rank-4 value.

### Rogue

<a id="rog-004"></a>
#### ROG-004: Deadly Poison V acquisition

*Affects: Small. State: Open.*

- **Why it matters:** Whether Rogue rows may use the highest Deadly Poison rank.
- **What the simulator assumes now:** The top rank is not assumed available without evidence of how it is learned.
- **Setup:** Level-60 Rogue (launch).
- **Steps:** Check the poison trainer and drop sources for Deadly Poison V.
- **Record:** How the rank is learned.
- **Sample size:** One check.
- **How the result is used:** Enables or excludes the rank.

### Priest

<a id="pri-001"></a>
#### PRI-001: Shadow Word: Death script value 150

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Shadow Word: Death damage.
- **What the simulator assumes now:** No hidden execute multiplier.
- **Setup:** Level 32+ Priest (launch).
- **Steps:** Cast Death above and below the execute threshold.
- **Record:** Damage and backlash.
- **Sample size:** 40 casts each.
- **How the result is used:** Settles the unexplained script value 150.

### Warlock

<a id="wl-002"></a>
#### WL-002: Incinerate's Immolate bonus scope

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Destruction Incinerate damage with Immolate.
- **What the simulator assumes now:** The 25% bonus multiplies the base damage only.
- **Setup:** Level 40+ Warlock (launch), two spell power sets.
- **Steps:** Cast Incinerate on targets with and without Immolate.
- **Record:** Non-crit damage.
- **Sample size:** 40 casts per case.
- **How the result is used:** Base-only versus whole-hit.

<a id="wl-004"></a>
#### WL-004: Spellstone's Fire/Shadow mask

*Affects: Small. State: Working answer.*

- **Why it matters:** Optional Warlock weapon imbue.
- **What the simulator assumes now:** Spellstone increases both Fire and Shadow damage, despite Shadow-only wording.
- **Setup:** Level-36+ Warlock with a Spellstone (launch).
- **Steps:** Cast Fire and Shadow spells with and without the Spellstone, same gear otherwise.
- **Record:** Non-crit damage per school.
- **Sample size:** 40 casts per school and setting.
- **How the result is used:** Confirms or removes the Fire bonus.

### Shared mechanics

<a id="core-007"></a>
#### CORE-007: Judgement of Wisdom trigger rules

*Affects: DPS rankings. State: Open.*

- **Why it matters:** Judgement of Wisdom mana for melee casters.
- **What the simulator assumes now:** 50% chance per hit, 59 mana at level 60.
- **Setup:** Level 38+ Paladin (launch).
- **Steps:** Judge Wisdom and let melee players attack.
- **Record:** Mana returns per white hit, special and spell.
- **Sample size:** 200 events.
- **How the result is used:** Proc chance and eligible events.

## No game test (code or modelling decision)

### Shaman

<a id="sha-004"></a>
#### SHA-004: Owned buff-totem effects and replacement

*Affects: DPS rankings. State: Code gap.*

- **Why it matters:** Self-sufficient Shaman scenarios cast their own totems.
- **What the simulator assumes now:** Owned Mana Spring does not yet restore mana; raid buffs supply the ranking effects.
- **Setup:** Code change, no game test.
- **Steps:** Make owned buff totems apply their rank-correct party effect and remove the previous totem of the same element.
- **Record:** Mana restored and buffs active with owned totems.
- **How the result is used:** Removes the gap without double counting raid buffs.

### Druid

<a id="dru-010"></a>
#### DRU-010: Feral stealth openers

*Affects: Small. State: Code gap.*

- **Why it matters:** Feral openers matter in short fights.
- **What the simulator assumes now:** No stealth state, Ravage or Pounce; their damage is absent rather than estimated.
- **Setup:** Code change, then a level-30 check of Ravage/Pounce damage.
- **Steps:**
  1. Add a one-time stealth opener state with Prowl, Ravage and Pounce.
  2. In game, record Ravage and Pounce damage from stealth at two AP values.
- **Record:** Opener damage and energy.
- **Sample size:** 20 openers per AP value.
- **How the result is used:** Adds the opener with measured values.

### Items and data

<a id="data-002"></a>
#### DATA-002: Projected gear budget and stat prices

*Affects: DPS rankings. State: Model choice.*

- **Why it matters:** All cross-class rankings use hypothetical level-65 items built from a stat budget.
- **What the simulator assumes now:** Sourced quantity, armor and weapon tables with assumed allocation budgets and special-stat prices.
- **Setup:** No game test until level-60+ loot exists.
- **Steps:**
  1. When real level-60/65 items are datamined, compare their stat totals with the model's budget for the same item level and slot.
  2. Refit the special-stat prices (hit, crit, spell power) from the real items.
- **Record:** Item level, slot, stats and model budget per item.
- **Sample size:** At least 30 items per armor class and slot group.
- **How the result is used:** Replaces the assumed budgets; rerun the gear search.

### Scenarios

<a id="scen-001"></a>
#### SCEN-001: Tank boss damage and external healing

*Affects: Tank results. State: Model choice.*

- **Why it matters:** Tank comparisons depend on how hard the boss hits and how much healing arrives.
- **What the simulator assumes now:** Fixed stress scenarios: one 300-second boss or three 90-second attackers with set swing, AP, armor and healing.
- **Setup:** Public raid logs once level-60 raids exist.
- **Steps:** Measure boss swing timer, hit size against known armor, and healing received by tanks in real raid logs.
- **Record:** Boss damage per swing, swing interval, tank healing per second.
- **Sample size:** Several kills per boss.
- **How the result is used:** Sets realistic scenario parameters; until then the stress scenarios stay as stated assumptions.

<a id="scen-003"></a>
#### SCEN-003: Tank target selection and threat coverage

*Affects: Tank results. State: Code gap.*

- **Why it matters:** Multi-target tank checks sum threat across targets and can hide a neglected add.
- **What the simulator assumes now:** Tank presets attack the first target; there is no lowest-threat target selector.
- **Setup:** Code change, no game test.
- **Steps:**
  1. Add an APL target selector that picks the enemy with the lowest threat margin.
  2. Report minimum threat margin per target in the Multi workload.
- **Record:** Per-target threat margin.
- **How the result is used:** Makes multi-target tank results meaningful.
