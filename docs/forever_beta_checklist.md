# Forever beta re-verification checklist

> Historical inventory. The [canonical register](uncertainties.md) replaces this
> file as the actionable checklist; old open counts and implementation claims
> below are not the current simulator state.

**The pass has happened.** This file was written before the beta client was datamined, when every number in the Forever ruleset came off a BlizzCon 2026 demo tooltip. Since 17 September the client has been the source: build `1.60.1.69913` on wago.tools, read against Classic Era `1.15.9.69722` and diffed spell by spell, with per-class write-ups in `docs/beta-pass/`. Most of what follows is history, kept because it records what each number used to rest on.

Where it stands, counted from `ui/core/spells/*.json`:

| | |
|---|---|
| on the client's numbers | 586 |
| unchanged from Classic, and checked | 246 |
| at least one number still assumed | 41 |
| not classified yet | 143 |

Fourteen `Open` lines remain across the per-class passes, and 141 `TODO`s are left in `sim/`. Those are the live work; the class sections below are the trail that led to them.

The demo mostly showed rank 1 of each talent, so where a talent had more ranks than the demo displayed, the implementation assumed a scaling and said so in a `TODO` beside the number. This file lists those `TODO`s, gathered from `grep -rn TODO sim/` and grouped by class. When a value is confirmed, fix the number if it moved, delete the comment, and regenerate the affected `.results` files with `make test && make update-tests`.

## How to run the pass

1. Re-export the talent trees from the beta client (the community calculators rebuild from it) and diff against `ui/core/talents/trees/*.json`: talent set, grid positions, rank counts and prerequisite arrows. `go test ./sim/ -run TestTalentTreesMatchTheirProtos` then pins the trees, protos and `TalentTreeSizes` together, and `-run TestPresetBuildsAreLegal` checks every shipped build still fits.
2. Work through the class sections below against the beta tooltips.
3. ~~The beta is capped at level 30, so it settles neither the level 60 ranks nor the coefficients.~~ **Wrong, corrected 17 September.** The level cap limits what a beta tester can play, not what the client ships. `SpellEffect.EffectBonusCoefficient` is populated for 11,398 effects, and each rank's damage scales to level 60 through the client's own `EffectRealPointsPerLevel` and `SpellLevels`. Both are settled; see `docs/spell-data-2026-09-17.md`. What the client genuinely does not carry is the server's side of a fight: proc chances that read as unset, and rules like whether melee-table Holy damage partially resists.
   The one reading still open is downranking. Era stored a reduced spell power coefficient on low ranks; Forever stores the full one on every rank from 3 up. Taken as written, a rank 4 Lightning Bolt does most of a rank 10 for a quarter of the mana, which is what the elemental rotation now does. Either Forever removed the penalty or it applies it somewhere the tables do not show, and the answer changes every caster in the sim.
4. The [race audit](forever_races.md) supersedes the demo's generic Eureka description. Client class variants specify 20% Energy, 40% Rage, 50% Mage/Warlock Mana, and 15% Priest Mana reductions, with three charges for 15 seconds and a two-minute cooldown. Remaining cast/periodic interactions and Read Ley Line regeneration are listed in [in-game checks](in_game_checks.md).
5. World buffs do not work inside Forever raids (reported 13 September from the demo; the sim ignores them under the Forever ruleset and hides the picker). Confirm on the beta client, and confirm whether the campsite buffs that replace them have combat numbers.
6. Re-run the DPS sweep across every spec and compare with the numbers recorded in the pull request history; anything that moves more than its change explains is worth a second look.

## The client's rank curves disagree with the beta tooltips (open, 2026-09-17)

The beta client stores per-rank talent values in a curve: `TraitDefinition` ->
`TraitDefinitionEffectPoints` -> `CurvePoint`, where each point is (rank, value). Read one
with `node tools/data_watch/trait_curve.mjs "<talent name>"`. `SpellEffect.EffectBasePointsF`
holds only the max-rank value and is sometimes stale, so it is the wrong field for this.

The method is sound - Ignite reads 8/16/24/32/40, Ruin 20..100 and Improved Life Tap 10/20,
all matching the tree exactly. On three talents the curve disagreed with the rank 1 value in the
talentsforever crawl:

| Talent | Crawl (rank 1) | Client curve, now in the sim |
|---|---|---|
| Druid / Moonglow | 3% | 8 / 17 / 25 |
| Druid / Moonfury | 1% | 2 / 4 / 6 / 8 / 10 |
| Rogue / Lethality | 6% | 4 / 8 / 12 / 16 / 20 |

**Resolved 17 September for the client.** The crawl's own data export says it was read off BlizzCon
demo footage and Blizzard's slides, not the beta, so it is older than build 1.60.1.69893 rather than a
second reading of it. `assets/confirmed_talents.json` is now generated from the beta client by
`tools/forever_talents/apply_beta_tooltips.py`, and the Go reads the curves above.

Beware two traps when reading these tables. A talent name shared across classes returns
several curves - Deflection returns four - so check the definition count. And some effects
are stored in milliseconds or tenths, which is why Improved Stings appears twice, once as
`2/4/6` and once as `2000/4000/6000`.

## Druid (16)

- `sim/druid/berserk.go:16` — The tooltip didn't show a cooldown, the 3 minutes are taken from the Classic Berserk.
- `sim/druid/demoralizing_roar.go:24` — assumed baseline, beta will confirm - Feral Aggression is gone from the tree, so the attack power reduction is taken at full strength like the warrior's Demoralizing Shout.
- `sim/druid/druid.go:122` — Improved Mark of the Wild is gone from the tree and is assumed to be baseline, beta will confirm.
- `sim/druid/faerie_fire.go:26` — The feral version's talent is gone from the tree and is assumed to be baseline, beta will confirm.
- `sim/druid/forms.go:223` — Only rank 1 of Furor was seen and the demo repeated it at every rank, which would leave four of the five points doing nothing. The Energy carryover, the out of form regen and its cap are all assumed to scale linearly, and the tree carries that assumption rather than an observation.
- `sim/druid/forms.go:342` — Only rank 1 of Furor's Bear shift was seen, a 20% chance at 10 Rage. Classic scaled the chance 20% per point for a flat 10 Rage and the tree now reads that.
- `sim/druid/lacerate.go:16` — assumed from Season of Discovery, beta will confirm the cost, the damage and the threat. Shredding Attacks names Lacerate, so the bear has it, but no tooltip for it has been seen: it is absent from the level 38 beta spellbook, which fits an ability learned later. Treat the Season of Discovery numbers as weaker than they look — Blizzard told Gamespot on 16 September that a lot of Season of Discovery "is likely not coming to a Forever game mode", so its tuning is a last resort rather than a related source.
- `sim/druid/mangle.go:17` — Only the tooltip was seen, the Energy cost is taken from the Classic Mangle (Cat).
- `sim/druid/mangle.go:71` — Only the tooltip was seen, the Rage cost, the 6 sec cooldown and the threat are taken from the Classic Mangle (Bear).
- `sim/druid/talents.go:301` — Only rank 1 of Eclipse was seen at 0.17 sec and the reduction is assumed to scale linearly, so rank 3 is 0.51 sec rather than the round half second the community talent calculator rounded it to.
- `sim/druid/talents.go:404` — Only rank 1 of Primal Fury was seen at 50%. Classic's Primal Fury and the Blood Frenzy folded into it both went from half the time to every time at rank 2, which is how both halves are read here.
- `sim/druid/talents.go:453` — Only rank 1 of Natural Reaction was seen, the dodge chance is assumed to scale linearly. Beta will confirm.
- `sim/druid/talents.go:456` — Only rank 1 of Natural Reaction's Rage proc was seen at 20%, the tree's 20/40/60/80/100 comes from the community talent calculator rather than from a tooltip, and at 5/5 it makes the Rage certain on every dodge.
- `sim/druid/talents.go:522` — Only rank 1 of Naturalist was seen, the damage bonus is assumed to scale linearly. Beta will confirm.
- `sim/druid/tigers_fury.go` — **resolved 16 September.** Soda's Druid on 13 September showed the tooltip: instant, 30 sec cooldown, no Energy cost, "Increases Physical damage done by 15% for 6 sec". The assumed Wrath shape was right about the cooldown and the cost, and wrong about the damage: Forever pays a share of Physical damage, not Classic's flat +40. Only the level 60 rank was seen, so the lower ranks keep the flat damage until one is observed.
- `sim/druid/wrath.go:53` — Only rank 1 of Improved Wrath was seen, the mana cost reduction is assumed to scale linearly. Beta will confirm.

## Hunter (9)

- `sim/hunter/aimed_shot.go` — assumed baseline, beta will confirm. Separately, the cast time may be shorter than Classic's 3.5 sec: Xaryu's Hunter showed a 2 sec cast on rank 3, and reaching that from a 3.5 base needs 75% ranged haste, which no talent grants. Probably a lower base under Forever, but one observation on a character of unknown gear does not settle the number — an unbuffed tooltip will.
- `sim/hunter/aspects.go:13` — Only rank 1 of Deadly Aspects was observed. The 30% attack speed for 12 sec is held at every rank, as in the Classic Improved Aspect of the Hawk this talent is built from, where the proc chance is the only slot the points buy.
- `sim/hunter/aspects.go:55` — only rank 1 was observed, the proc chance is the one slot of Deadly Aspects assumed to scale per rank.
- `sim/hunter/rapid_fire.go:15` — only rank 1 of Rapid Killing was observed, the reduction is assumed to scale per rank. The buff a kill grants is not modelled, so nothing checks the 40 sec window or the 20% damage the tree reads at rank 2.
- `sim/hunter/serpent_sting.go:43` — only rank 1 of Improved Stings was observed, the damage bonus is assumed to scale per rank. The Viper Sting cooldown and the Scorpid Sting duration are not modelled, so nothing checks the rest of what the tree reads.
- `sim/hunter/talents.go:300` — only rank 1 was observed, the cost reduction and the proc chance of Resourcefulness are assumed to scale per rank. The 50% regeneration and the 30 sec window do not, which is what the tree reads.
- `sim/hunter/talents.go:354` — only rank 1 was observed, the Rapid Recuperation regeneration is assumed to scale per rank. The 15 sec window is held at both ranks: a duration read off a single tooltip is not extrapolated.
- `sim/hunter/talents.go:378` — only rank 1 of Expose Prey was observed, the proc chance is assumed to scale per rank. The 5 sec window does not, which is what the tree reads.
- `sim/core/buffs.go`, `BattleShoutAura` and `BlessingOfMightAura` — both grant melee attack power only, as in Classic. If Forever lets either reach ranged attack power every ranked hunter gains 10% to 12%.

## Mage (3)

- `sim/mage/fire_blast.go:39` — only rank 1 of Wake of Fire was shown, so mage.json copies its 1 sec into rank 2.
- `sim/mage/talents.go` — **resolved 17 September.** The beta tooltip for Fingers of Frost rank 2 settled what the demo could not: the proc chance does not scale, both ranks give Chill effects a 15% chance, and the second point buys a second charge ("treats your next 2 spells cast as if the target were Frozen"). The aura now carries a stack per point and spends one per cast. Frost lost 3.2%.
- `sim/mage/talents.go` — **resolved 17 September.** Shatter is three ranks at 17/33/50% in the beta client.
- `sim/mage/talents.go:674` — a cast already in progress when a chill lands is held out of Fingers of Frost, so it neither takes the Shatter crit nor spends the charge. Beta will confirm which cast the charge belongs to.

## Paladin (23)

- `sim/paladin/consecration.go:10` — assumed baseline, beta will confirm - Consecration is no longer a talent and the Forever tree builds on top of it through Consecrated Ground and Holy Conduit.
- `sim/paladin/hammer_of_wrath.go:29` — Only rank 1 of Instrument of Law was seen at 0.5 sec, the full second the tree reads at rank 2 comes from the community talent calculator rather than from a tooltip.
- `sim/paladin/holy_shield.go:18` — Only rank 1 was seen at 110, up from Classic's 65. The other ranks are scaled by the same ratio until the beta shows them.
- `sim/paladin/holy_strike.go:15` — assumed baseline, beta will confirm - only the level 60 rank is modelled, and the flat damage is taken from the published tooltip rather than from the game. Forever's own spell id for Holy Strike is 17143, which the item database does not carry, so the sim keeps Classic's unused 13953.
  - Checked against the client's `SpellEffect` on 2026-09-17: spell 17143 is `Effect=58`, weapon damage plus a flat amount, which is the shape the sim models. **The 0.429 spell power coefficient is still unconfirmed and the client cannot settle it** - `Coefficient` is zero on all 42,449 rows in build 1.60.1.69893, so the column is simply not populated. Do not read that zero as "no coefficient". A beta tooltip or a damage log is still what is needed.
- `sim/paladin/holy_strike.go:17` — Holy damage on the melee hit table, so it rolls partial resists the way every other Holy ability here does. Whether a melee-table Holy strike actually partial resists is unknown; if it does not, it wants `SpellFlagIgnoreResists`. Raised by AdamRC from the demo, 16 September.
- `sim/paladin/holy_strike.go:29` — Only rank 1 of Improved Holy Strike was seen, the second second of cooldown is assumed to scale linearly.
- `sim/paladin/holy_strike.go:41` — Only rank 1 of Iron Creed's threat was seen at 5%, the 5% per rank the tree reads comes from the community talent calculator rather than from a tooltip.
- `sim/paladin/holy_strike.go:94` — Only rank 1 of Iron Creed's damage reduction was seen at 2%, the 2% per rank the tree reads comes from the community talent calculator rather than from a tooltip. The 6 seconds is flat at every rank.
- `sim/paladin/sotc.go:34` — assumed baseline, beta will confirm - Improved Seal of the Crusader is gone from the tree and the raid reads the improved Judgement of the Crusader through Debuffs either way.
- `sim/paladin/swift_judgement.go:11` — assumed baseline, beta will confirm - the tooltip carries no cooldown, so it is given a minute, long enough that it buys one extra Judgement rather than a second rotation.
- `sim/paladin/talents.go:19` — Only rank 1 of Divine Precision was seen, ranks 2 and 3 are extrapolated from it.
- `sim/paladin/talents.go:33` — Only rank 1 of Holy Power was seen, ranks 2 to 5 are extrapolated from it. The talent is split across two files: every spell gets 1% crit per point here and Holy Shock picks up another 2% per point in `holy_shock.go`, which together are the 3% per point the tree reads for it. Confirm both halves, and keep them in step if either moves.
- `sim/paladin/talents.go:37` — Only rank 1 of Sacred Duty was seen, the 2% per rank the tree reads comes from the community talent calculator rather than from a tooltip.
- `sim/paladin/talents.go:41` — Only rank 1 of Shield Specialization's absorb was seen, the 10% per rank the tree reads comes from the community talent calculator rather than from a tooltip.
- `sim/paladin/talents.go:47` — Only rank 1 of Champion of the Light was seen, and the extrapolated ranks 2 and 3 are a large chunk of a Forever paladin's spell power.
- `sim/paladin/talents.go:89` — Every rank of Redoubt reads the same 10% chance for 6% block, so ranks 2-5 do nothing. The tree says the same thing, so the picker and the sim agree; neither has a source, because the generator could not line Classic's Redoubt up against Forever's wording and copied rank 1 into ranks 2 to 5.
- `sim/paladin/talents.go:170` — Only rank 1's 33% chance was seen, the tree's 33/66/100 comes from the community talent calculator rather than from a tooltip. The 6% of maximum mana does not scale.
- `sim/paladin/talents.go:231` — Every rank of Vengeance reads 1% per stack up to 5 stacks, so ranks 2 and 3 do nothing. The tree says the same thing; as with Redoubt the generator copied rank 1 into the later ranks rather than extrapolating.
- `sim/paladin/talents.go:264` — The Vindication self buff reads 1% at every rank, and so does the tree, again from a copied rank 1. The 42 attack power the target loses is not modelled, nothing in the sim reads an enemy's attack power.
- `sim/paladin/talents.go:301` — The tooltip caps the bonus at the first 4 or 8 enemies to enter the Consecration, which is not modelled here - everything standing in it gets the bonus.
- `sim/paladin/talents.go:337` — Only rank 1's 10% was seen, the tree's second rank comes from the community talent calculator rather than from a tooltip.
- `sim/paladin/templars_bulwark.go:11` — assumed baseline, beta will confirm - the tooltip carries no cooldown, so it shares the 5 minutes of the two Forbearance abilities Sacred Duty shortens alongside it.
- `sim/paladin/templars_bulwark.go:36` — Only rank 1 of Sacred Duty was seen at 30 sec, the tree's second rank comes from the community talent calculator rather than from a tooltip.

Infusion of Light is not on that list because there is nothing to check against: the sim has no Holy Light for it to shorten, so it is one of the talents below that the sim never reads. Its rank 2 half second still came from the community talent calculator, so the beta pass should read the tooltip for it rather than treating the tree as settled.

## Priest (12)

- `sim/priest/devouring_plague.go:53` — only rank 1 of Devouring Contagion was shown, beta will confirm the rank 2 value
- `sim/priest/holy_nova.go:13` — beta will confirm the higher ranks and the mana cost.
- `sim/priest/mind_flay.go:81` — Only rank 1 of Improved Mind Flay was seen at 10%, the 20% the tree reads at rank 2 comes from the community talent calculator rather than from a tooltip.
- `sim/priest/penance.go:18` — beta will confirm the cost, the cooldown and the level 60 damage.
- `sim/priest/power_infusion.go:10` — let the option pick a raid member once buffing another player is modelled.
- `sim/priest/power_infusion.go:25` — beta will confirm the cost and the cooldown.
- `sim/priest/priest.go:81` — beta will confirm whether they were made baseline or removed outright.
- `sim/priest/priest.go:97` — beta will confirm that Devouring Plague is no longer race locked.
- `sim/priest/talents.go:40` — only rank 1 was shown, beta will confirm that the two halves scale at 5% and 1% per point
- `sim/priest/talents.go:57` — only rank 1 was shown, beta will confirm the 2% per point
- `sim/priest/talents.go:167` — Only rank 1 of Searing Light was shown, so rank 2's 4% Holy damage and 10% Holy Nova refund chance are both extrapolated from it. Beta will confirm them.
- `sim/priest/talents.go:400` — beta will confirm the 50%.

The priest healing spellbook is commented out in `RegisterHealingSpells`, so Flash Heal, Greater Heal, Power Word: Shield, Prayer of Healing and Renew never reach a sim. Every talent whose only effect is on those spells is therefore unmeasurable at any rank, whatever the picker promises: Improved Power Word: Shield, Improved Renew, Improved Healing, Spiritual Healing and Inspiration. The commented code also predates the current trees and reads Improved Power Word: Shield at 5/10/15% against the tree's 7/14/21%, Spiritual Healing at 2/4/6% against the tree's 3/6/9%, and Silent Resolve's healing threat at 7/14/20% against the tree's 10/20/30%. Fix those alongside whatever brings the spellbook back rather than in isolation, or they come back wrong.

## Rogue (11)

- `sim/rogue/backstab.go:28` — Only rank 1 of Puncturing Wounds was seen, so both the extra combo point chance and the Backstab crit chance are assumed to scale linearly. Beta will confirm.
- `sim/rogue/expose_armor.go:11` — assumed baseline, beta will confirm. The raid reads this debuff, so the Classic 2/2 armor value is treated as baseline rather than deleted.
- `sim/rogue/expose_armor.go:36` — Only rank 1's 5 Energy was seen. The tree's second rank is a linear extrapolation of it rather than an observation, and the refund and the 5 combo point trigger stay put. Beta will confirm.
- `sim/rogue/hack_and_slash.go:14` — Only rank 1 was seen, and the data behind the tree copies it into every other rank rather than observing them, so the linear scaling of all three effects is an assumption. Beta will confirm.
- `sim/rogue/mutilate.go:43` — The tooltip showed no Energy cost, the 60 is taken from the Classic Mutilate.
- `sim/rogue/mutilate.go:55` — The tooltip didn't repeat the Classic dagger requirement, it's assumed to still apply.
- `sim/rogue/mutilate.go:59` — Only rank 1 of Puncturing Wounds was seen, the crit chance it gives Mutilate is assumed to scale linearly. Beta will confirm.
- `sim/rogue/poisons.go:63` — The tooltip doesn't say whether Venom reaches a Deadly Poison that is already on the target or only the stacks applied while it is up. Beta will confirm.
- `sim/rogue/talents.go:273` — Only rank 1 of Cutthroat's 3% was seen, and the data behind the tree copies it into every other rank rather than observing them, so the linear scaling is an assumption. The 10 sec duration is held at rank 1. Beta will confirm.
- `sim/rogue/talents.go:352` — Only rank 1 of Quietus was seen and the damage bonus is assumed to scale linearly. The health threshold cannot be extrapolated alongside it, so rank 1's 35% is used for every rank and the tree now reads the same. Beta will confirm.
- `sim/rogue/venom.go:47` — The tooltip showed no Energy cost, the 25 matches the other Rogue finishers.

## Shaman (16)

- `sim/shaman/air_totems.go:50` — The sim won't respect the value of a totem dropped via the APL. It uses hard-coded values from buffs.go bonusDamage := WindfuryTotemBonusDamage[rank]
- `sim/shaman/lava_burst.go:11` — Only the damage range and the Flame Shock bonus were on the tooltip. The cast time, cooldown, mana cost and coefficient are taken from the spell of the same name, beta will confirm them.
- `sim/shaman/lightning_overload.go:25` — Only rank 1 was seen, beta will confirm that the ranks go up in steps of 3%.
- `sim/shaman/talents.go:68` — Only rank 1 of Improved Reincarnation's 2% health was seen, the 4% the tree reads at rank 2 comes from the community talent calculator rather than from a tooltip.
- `sim/shaman/talents.go:139` — Only rank 1 of Elemental Alacrity's 0.17 sec was seen, the 0.34 and 0.51 the tree reads come from the community talent calculator rather than from a tooltip.
- `sim/shaman/talents.go:145` — Only rank 1 of Improved Fire Nova's 10% and 2 sec were seen, the doubled rank 2 the tree reads comes from the community talent calculator rather than from a tooltip.
- `sim/shaman/talents.go:252` — Only rank 1 of Elemental Fury's 20% was seen, the steps up to 100% the tree reads come from the community talent calculator rather than from a tooltip.
- `sim/shaman/talents.go:420` — Only rank 1 of Improved Stormstrike's 50% was seen, the doubling to a certainty at rank 2 comes from the community talent calculator rather than from a tooltip, and a talent that makes two separate rolls certain is worth a second look.
- `sim/shaman/talents.go:428` — Improved Stormstrike's 15 sec window is rank 1's and is applied at both ranks. The community talent calculator reads 30 sec at rank 2, but it extrapolates every number in a tooltip and a buff whose duration grows with the talent would be unusual.
- `sim/shaman/talents.go:463` — Maelstrom Weapon's tooltip never showed a proc rate and only rank 1's 4% was seen, beta will confirm both.
- `sim/shaman/talents.go:470` — Maelstrom Weapon's five stacks and 30 sec are rank 1's and are applied at every rank, because five stacks of 4% per point reach exactly a free instant cast at 5/5 and the calculator's extrapolated 25 stacks over 150 sec overshoot it several times over.
- `sim/shaman/talents.go:515` — The tooltip showed no cooldown, beta will confirm it. 3 minutes matches the other class cooldowns of this size.
- `sim/shaman/totems.go:9` — Assumed baseline rather than deleted, beta will confirm it.
- `sim/shaman/water_shield.go:18` — "Only one globe will activate every few seconds", the tooltip never said how long.
- `sim/shaman/water_totems.go:115` — The sim won't respect the value of a totem dropped via the APL. It uses hard-coded values from buffs.go manaRestoreBase := ManaSpringTotemManaRestore[rank]
- `sim/shaman/windfury_weapon.go:78` — Classic lets both weapons carry the imbue and gives the extra attacks to the hand that procced, beta will confirm that Forever kept both halves of that.

## Warlock (12)

- `sim/warlock/conflagrate.go:13` — The Forever tooltip puts Conflagrate rank 1 at 109 to 132, less than half Classic's 249 to 316, and Incinerate at 125 to 140 against the 380 to 440 used here. Neither spell's higher ranks were shown, so both keep their Classic tables.
- `sim/warlock/conflagrate.go:23` — Only rank 1 of Shadow and Flame was seen and every rank of the tree reads the same 20% chance not to consume Immolate, so the chance is held there. Read per point it would be certain at 5/5 and Conflagrate would stop consuming Immolate at all.
- `sim/warlock/immolate.go:74` — Only rank 1 of Aftermath was seen. The initial Immolate damage is read per point; the Daze half is not modelled, and its chance, its slow and its duration are held at rank 1's values because per point 5/5 would Daze on every Conflagrate and slow by 250%.
- `sim/warlock/shadowburn.go:12` — The Forever tooltip puts Shadowburn rank 1 at 102 to 111 rather than Classic's 91 to 104. The other five ranks were never shown, so the Classic table is kept.
- `sim/warlock/soul_fire.go:53` — Only rank 1 of Decimation was seen, so the 45% Soul Fire cooldown reduction is held at both ranks. Read per point the second rank would leave Soul Fire on a six second cooldown.
- `sim/warlock/talents.go:168` — Only rank 1 of Improved Drains was seen. The bonus per Affliction effect and its cap are read per point, the 20% health the Drain Soul bonus triples below is a threshold rather than a magnitude and is held there at every rank.
- `sim/warlock/talents.go:359` — Only rank 1 of Decimation was seen. The 3% damage and the 20% cast time reduction are read per point, the health threshold and the ten second window are held at rank 1's values. The aura's 63165 names no Forever spell and looks like a transposition of the tree's 63156.
- `sim/warlock/talents.go:414` — Only rank 1 of Demonic Brand was seen. The Searing Pain threat reduction is read per point, the ten second brand, the two pet attacks it arms and their 39 to 42 damage are held at rank 1's values, so points two and three buy threat alone.
- `sim/warlock/talents.go:476` — Beta will show whether the 33% of level is per rank or the full value
- `sim/warlock/talents.go:806` — Only rank 1 of Improved Shadow Bolt was seen. The extra Shadow damage taken is read per point, the twelve seconds it lasts are held at rank 1's value.
- `sim/warlock/talents.go:916` — Only rank 1 of Shadow and Flame was seen and every rank of the tree reads the same 2% for 20 sec, so the damage bonus is held there rather than read per point.
- `ui/core/talents/trees/warlock.json` — Soul Harvesting, Improved Health Funnel, Demonic Aegis, Improved Voidwalker, Improved Felhunter, Destructive Reach, Molten Skin, Pyroclasm, Fel Concentration and Intensity have no implementation in the sim, so nothing checks their per-rank values from the sim's side. Four of them carry per-rank values stamped as matching the scaling the sim applies, which the sim never applied.

## Warrior (11)

- `sim/warrior/demoralizing_shout.go:15` — assumed baseline, beta will confirm
- `sim/warrior/shield_wall.go` — only rank 1 of Improved Shield Wall was shown, beta will confirm that rank 2 is another 5.5 minutes. The spell's own numbers are no longer Classic's: Esfand's Warrior showed a 15 min cooldown reducing damage taken from all attacks by 60% for 12 sec, against Classic's 30 min, 75% and 10 sec. All three are applied under Forever.
- `sim/warrior/shouts.go:55` — assumed baseline, beta will confirm
- `sim/warrior/talents.go:76` — only rank 1 of Weaponmaster was shown. The 1% crit / 3% armor / 1% extra attack per point the tree lists comes from the community talent calculator, which is rebuilt from the same rank 1 tooltip, so it agrees with the sim without confirming it.
- `sim/warrior/talents.go:177` — only rank 1 of Unbridled Wrath was shown, beta will confirm the 12% per point. The tree now lists 12/24/36/48/60% instead of repeating rank 1 at every rank.
- `sim/warrior/talents.go:206` — every rank of Dual Wield Specialization reads the same 5% damage, 20% Rage and 2% hit; the damage is Classic's per point value, so all three are read per point. Beta will confirm the Rage and the hit, which together are the largest single source of a dual wielding Fury build's Rage income.
- `sim/warrior/talents.go:234` — only rank 1 of Enrage was shown at a 30% chance. Read per point the chance passes 100% at 4/5, so the sim caps it and the fifth point buys nothing; the tree reads 30/60/90/100/100 to match. Classic ranked the damage and left the chance flat, so it offers no slope for the half Forever put the ranks on. Beta will confirm the chance at each rank.
- `sim/warrior/talents.go:376` — only rank 1 was shown, beta will confirm that rank 2 doubles both the chance and the rage.
- `sim/warrior/talents.go:402` — only rank 1 was shown, beta will confirm the 2% per point.
- `sim/warrior/talents.go:413` — only rank 1 was shown, beta will confirm the 2% per point.
- `sim/warrior/talents.go:422` — only rank 1 was shown, beta will confirm the 1 Rage per point.

## Talent icons still standing in (3)

Three talents draw an icon that is not their own. Their real icon was only ever seen in the
demo video, and the frames it was cut from are gone, so they keep the icon of the Classic
talent whose spell ids the tree carried. The name, description and per-rank numbers on each
are the talent's own; only the picture is borrowed. Re-cut them from the beta client.

- `arcaneGeometry` (mage, Arcane) draws Flame Throwing's icon.
- `divinePrecision` (paladin, Holy) draws Precision's icon.
- `twinDisciplines` (priest, Discipline) draws Spiritual Healing's icon.

## Baseline ability changes

**Done, 17 September.** Every class's spellbook has been diffed against Classic Era with `tools/data_watch/spell_client.py --learned <class>`, and what changed is written up per class in `docs/beta-pass/`. The section below is what the panel said before that, kept because it is what each of these lines used to rest on.

Forever also changes abilities that are not talents. None of their tooltips were shown at BlizzCon; what follows comes from the panel and the coverage of it, so every line is a claim to check against the beta spellbook rather than a number to confirm.

Warrior, the only class with concrete changes reported:

- Slam no longer resets the swing timer — modelled in `sim/warrior/slam.go`.
- Thunder Clap can be used in Defensive Stance — modelled in `sim/warrior/thunder_clap.go`, and the protection rotation casts it on cooldown.
- Improved Shield Wall shortens the cooldown instead of lengthening the duration — modelled in `sim/warrior/shield_wall.go`.
- `sim/warrior/stances.go:41` — Tactical Mastery is baseline with Improved Tactical Mastery on top, but the baseline was never shown a number, so only the talent's own 3 Rage per point is modelled and an untalented warrior keeps nothing across a stance change. Beta will show what the baseline retains.
- Victory Rush is baseline — not modelled. It needs a killing blow, which a boss encounter never gives before the fight ends.

~~Other classes: the panel spoke of baseline changes across every class without listing them, and nothing more specific has been published. When the beta client is datamined, diff each class spellbook against Classic Era and add every changed ability here with the file that models it, or the reason it is left out.~~ **Done** — each class's "Spellbook diff" section in `docs/beta-pass/` lists what it found and what was left out.

## Talents the sim does not read (37)

These are in the trees and the picker marks them as not simulated; spending points in them changes nothing. Most are utility or PvP talents the Classic sim never modelled either. They are listed so the beta pass can confirm none of them turned into something a raid rotation cares about.

- Druid / Balance: Overgrowth
- Druid / Restoration: Gift of the Earthmother
- Druid / Restoration: Wild Growth
- Hunter / Survival: Strider Kick
- Paladin / Holy: Voice of Truth
- Paladin / Holy: Light's Vigil
- Paladin / Holy: Infusion of Light
- Paladin / Protection: Improved Seal of Fury
- Priest / Discipline: Wand Specialization (the sim has no wand attacks at all, so a shadow build's two points here are idle)
- Priest / Discipline: Martyrdom
- Priest / Discipline: Improved Inner Fire (Inner Fire itself is not in the spellbook)
- Priest / Discipline: Soul Warding
- Priest / Discipline: Improved Mana Burn (Mana Burn is not in the spellbook, so neither rank does anything)
- Priest / Discipline: Renewed Hope
- Priest / Discipline: Divine Aegis
- Priest / Holy: Twilight Focus (the sim never interrupts a cast, so pushback resistance has nothing to resist)
- Priest / Holy: Blessed Recovery (the Smite build spends two points here)
- Priest / Holy: Holy Reach
- Priest / Holy: Binding Heal
- Priest / Holy: Litany of Light
- Priest / Holy: Spirit of Redemption
- Priest / Shadow Magic: Blackout (a stun, and a raid boss is immune; the Shadow build spends five points here)
- Priest / Shadow Magic: Spirit Tap (it needs a kill, and nothing dies in a raid encounter; the Smite build spends three points here)
- Priest / Shadow Magic: Shadow Reach
- Priest / Shadow Magic: Improved Psychic Scream
- Priest / Shadow Magic: Improved Fade
- Priest / Shadow Magic: Silence
- Priest / Shadow Magic: Early Demise (Shadow Word: Death is not in the spellbook at all)
- Rogue / Subtlety: Improved Distract
- Shaman / Restoration: Riptide
- Warlock / Demonology: Demonic Aegis
- Warlock / Demonology: Improved Felhunter
- Warlock / Destruction: Molten Skin
- Warrior / Arms: Spearing Strike
- Warrior / Fury: Blood Craze
- Warrior / Protection: Concussion Blow (a stun, and a raid boss is immune)
- Warrior / Protection: Vanguard

## Item stats: the gear planner disagrees with the item database (found 17 September)

`assets/db_inputs/wowhead_forever_gearplanner.txt`, the snapshot the data watcher takes
from Wowhead's Forever gear planner, carries stats for 6,533 items. Comparing agility,
strength, intellect, spirit, stamina and armor against `assets/database/db.json`:

- 13,896 values compared on items both carry
- **603 differ, across 482 items**, and 200 of those are item level 60 or higher
- **19 of them appear in shipped gear sets**, mostly the PvP epics: Field Marshal's and
  Marshal's Dragonhide pieces, Onyxia Tooth Pendant, Hammer of Bestial Fury

The differences run both ways - 179 where only Wowhead carries a value, 190 where only the
sim does - which argues against a Forever retune, since that would push one direction.
Onyxia Tooth Pendant reads 12 agility and 9 stamina here against Wowhead's 7 and 8, which
looks like two eras of the same item rather than a change.

**Not acted on.** Nothing in the snapshot says which client version it was built from, and
the sim's database came from a separate scrape, so neither side is obviously right. Taking
Wowhead's numbers would change 19 items a player actually sims on a guess.

What would settle it: an item tooltip read off the beta client for any of the 19, or a
version marker on the gear planner payload. The scaling tables in the same snapshot all
agree with the sim - base stat offsets for all ten races, spell crit per intellect for all
seven casting classes - so the snapshot is not wrong in general.

**Settled 2026-09-17: the sim is right, no change needed.** `ItemSparse` from the beta
client (wago.tools build 1.60.1.69893, the table `tools/data_watch/wago_db2_diff.py`
already watches) carries the stat block for 2,316 of the sim's items. Onyxia Tooth
Pendant - the clearest of the 19 - reads agility, stamina, fire resistance, hit and crit
in the client, exactly the stats the sim gives it. The gear planner's 7/8 was the outlier.

Reading that table takes care, and three passes were wrong before the fourth was right:

- `StatPercentEditor` is an allocation budget, not the displayed value, so only *which*
  stats an item carries can be compared, never the numbers.
- The mod ids are the retail `ITEM_MOD_*` set: 45 is spell power (not 31), 38 attack
  power, 31/32 generic hit and crit. Guessing 31 for spell power invented 821 retunes.
- The resistance ids are 51 fire, 52 frost, 54 shadow, 55 nature, 56 arcane - pinned by
  the items carrying them (Fiery Cloak, Icy Cloak, Ring of the Shadow, Dragonscale). An
  off-by-one here turned every nature resist into a fake arcane one.
- The client splits no hit/crit into melee and spell variants; the sim does, and picks by
  item type. Fold the sim's back down before comparing.

What is left after those corrections is ~350 items, and the ones checked are all the same
known representation difference rather than a retune: Classic delivered weapon spell power
through an equip effect, which does not live in `ItemSparse`. Grand Marshal's Mageblade
holds its 72 spell power that way, Hammer of the Gathering Storm its 53. The script is
`review/compare-itemsparse.mjs`.
