package common

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/common/itemhelpers"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// Ordered by ID
const (
	Destiny                   = 647
	ShortswordOfVengeance     = 754
	FieryWarAxe               = 870
	Bloodrazor                = 809
	HammerOfTheNorthernWind   = 810
	FlurryAxe                 = 871
	SkullflameShield          = 1168
	TeebusBlazingLongsword    = 1728
	Nightblade                = 1982
	Shadowblade               = 2163
	GutRipper                 = 2164
	HandOfEdwardTheOdd        = 2243
	BowOfSearingArrows        = 2825
	Gutwrencher               = 5616
	SwordOfZeal               = 6622
	Ravager                   = 7717
	HanzoSword                = 8190
	TheJackhammer             = 9423
	PendulumOfDoom            = 9425
	ManualCrowdPummeler       = 9449
	BloodletterScalpel        = 9511
	TheHandOfAntusul          = 9639
	GryphonRidersStormhammer  = 9651
	Ragehammer                = 10626
	EnchantedAzshariteSword   = 10696
	ColdrageDagger            = 10761
	Firebreather              = 10797
	DragonsCall               = 10847
	VilerendSlicer            = 11603
	HookfangShanker           = 11635
	Naglering                 = 11669
	Ironfoe                   = 11684
	Bloodfist                 = 11744
	FlameWrath                = 11809
	HandOfJustice             = 11815
	LordGeneralsSword         = 11817
	SecondWind                = 11819
	BurstOfKnowledge          = 11832
	WraithScythe              = 11920
	LinkensSwordOfMastery     = 11902
	SearingNeedle             = 12531
	KerisOfZulSerak           = 12582
	BlackhandDoomsaw          = 12583
	BlackbladeOfShahram       = 12592
	BarmanShanker             = 12791
	VolcanicHammer            = 12792
	SeepingWillow             = 12969
	Felstriker                = 12590
	PipsSkinner               = 12709
	BlazingRapier             = 12777
	ArcaniteChampion          = 12790
	MasterworkStormhammer     = 12794
	BloodTalon                = 12795
	Frostguard                = 12797
	Annihlator                = 12798
	SerpentSlicer             = 13035
	TheNeedler                = 13060
	Chillpike                 = 13148
	Venomspitter              = 13183
	Bashguuder                = 13204
	SealOfTheDawn             = 13209
	SmolderwebsEye            = 13213
	FangOfTheCrystalSpider    = 13218
	ArgentAvenger             = 13246
	Rivenspike                = 13286
	SkullforgeReaver          = 13361
	TheCruelHandOfTimmy       = 13401
	RunebladeOfBaronRivendare = 13505
	// HeadmastersCharge      = 13937
	GravestoneWarAxe           = 13983
	Darrowspike                = 13984
	Frightalon                 = 14024
	BonechillHammer            = 14487
	EbonHiltOfMarduk           = 14576
	FrightskullShaft           = 14531
	BarovianFamilySword        = 14541
	CloudkeeperLegplates       = 14554
	AlcorsSunrazor             = 14555
	HameyasSlayer              = 15814
	JoonhosMercy               = 17054
	DrillborerDisk             = 17066
	Deathbringer               = 17068
	GutgoreRipper              = 17071
	Shadowstrike               = 17074
	ViskagTheBloodletter       = 17075
	BonereaversEdge            = 17076
	BlazefuryMedallion         = 17111
	EmpyreanDemolisher         = 17112
	SulfurasHandOfRagnaros     = 17182
	SulfuronHammer             = 17193
	Thunderstrike              = 17223
	ThrashBlade                = 17705
	SatyrsLash                 = 17752
	MarkOfTheChosen            = 17774
	BladeOfEternalDarkness     = 17780
	ForceReactiveDisk          = 18168
	EskhandarsLeftClaw         = 18202
	EskhandarsRightClaw        = 18203
	FiendishMachete            = 18310
	RazorGauntlets             = 18326
	QuelSerrar                 = 18348
	BaronCharrsSceptre         = 18671
	TalismanOfEphemeralPower   = 18820
	EssenceOfThePureFlame      = 18815
	PerditionsBlade            = 18816
	Thunderfury                = 19019
	GlacialBlade               = 19099
	ElectrifiedDagger          = 19100
	Nightfall                  = 19169
	EbonHand                   = 19170
	DarkmoonCardHeroism        = 19287
	DarkmoonCardBlueDragon     = 19288
	DarkmoonCardMaelstrom      = 19289
	TheLobotomizer             = 19324
	RuneOfTheDawn              = 19812
	DraconicInfusedEmblem      = 22268
	HeartOfWyrmthalak          = 22321
	GauntletsOfUndeadSlaying   = 23078
	HandwrapsOfUndeadSlaying   = 23081
	HandguardsOfUndeadSlaying  = 23082
	GlovesOfUndeadCleansing    = 23084
	RobeOfUndeadCleansing      = 23085
	BreastplateOfUndeadSlaying = 23087
	ChestguardOfUndeadSlaying  = 23088
	TunicOfUndeadSlaying       = 23089
	BracersOfUndeadSlaying     = 23090
	BracersOfUndeadCleansing   = 23091
	WristguardsOfUndeadSlaying = 23092
	WristwrapsOfUndeadSlaying  = 23093
	MarkOfTheChampionPhys      = 23206
	MarkOfTheChampionSpell     = 23207
)

func init() {
	core.AddEffectsToTest = false

	// ! Please keep items ordered alphabetically within a given category !

	// Many referenced from Classic WoW Armaments - https://discord.gg/w95B2hXfBF

	///////////////////////////////////////////////////////////////////////////
	//                                 Weapons
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=14555/alcors-sunrazor
	// Chance on hit: Blasts a target for 75 to 105 Fire damage.
	// 1 PPM from Armaments Discord
	itemhelpers.CreateWeaponCoHProcDamage(AlcorsSunrazor, "Alcor's Sunrazor", 1.0, 18833, core.SpellSchoolFire, 75, 30, 0, core.DefenseTypeMagic)

	//https://www.wowhead.com/classic/item=12798/annihilator
	// Beta client 1.60.1 (spell 16928): 165 armor a stack, down from Era's 200. Still 3 stacks.
	// 1 PPM from Armaments Discord but may be higher
	itemhelpers.CreateWeaponProcSpell(Annihlator, "Annihlator", 1.0, func(character *core.Character) *core.Spell {
		armorShatterAuras := character.NewEnemyAuraArray(ArmorShatterAuras)

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16928},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				activateAura := armorShatterAuras.Get(target)
				activateAura.Activate(sim)

				if activateAura.IsActive() {
					activateAura.AddStack(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/spell=16916/strength-of-the-champion
	// Chance on hit: Heal self for 270 to 450 and Increases Strength by 120 for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(ArcaniteChampion, "Arcanite Champion", 1.0, func(character *core.Character) *core.Aura {
		actionID := core.ActionID{SpellID: 16916}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.GetOrRegisterAura(core.Aura{
			Label:    "Strength of the Champion",
			ActionID: actionID,
			Duration: time.Second * 30,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.GainHealth(sim, sim.Roll(270, 450), healthMetrics)
				character.AddStatDynamic(sim, stats.Strength, 120)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.Strength, -120)
			},
		})
	})

	// https://www.wowhead.com/classic/item=13246/argent-avenger
	// Beta client 1.60.1 (spell 17352): "Increases Attack Power by 100, and Attack Power against
	// Undead by an additional 100 for 10 sec" - so it is now worth something against everything,
	// where Era's 200 only counted against Undead. The client also carries 100 ranged attack power.
	// 1 PPM from Armaments Discord
	itemhelpers.CreateWeaponProcAura(ArgentAvenger, "Argent Avenger", 1.0, func(character *core.Character) *core.Aura {
		matchingTargets := core.FilterSlice(
			character.Env.Encounter.TargetUnits,
			func(unit *core.Unit) bool { return unit.MobType == proto.MobType_MobTypeUndead },
		)

		return character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 17352},
			Label:    "Argent Avenger",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{stats.AttackPower: 100, stats.RangedAttackPower: 100})
				for _, target := range matchingTargets {
					for _, at := range character.AttackTables[target.UnitIndex] {
						at.BonusAttackPowerTaken += 100
					}
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{stats.AttackPower: -100, stats.RangedAttackPower: -100})
				for _, target := range matchingTargets {
					for _, at := range character.AttackTables[target.UnitIndex] {
						at.BonusAttackPowerTaken -= 100
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=12791/barman-shanker
	// Chance on hit: Wounds the target causing them to bleed for 100 damage over 30 sec.
	// Assumed 1 PPM
	itemhelpers.CreateWeaponProcSpell(BarmanShanker, "Barman Shanker", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13318},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Barman Shanker)",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 10, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=18671/baron-charrs-sceptre
	// Chance on hit: Blasts a target for 35 Fire damage.
	// 1 PPM assumed
	itemhelpers.CreateWeaponCoHProcDamage(BaronCharrsSceptre, "Baron Charr's Sceptre", 1.0, 13442, core.SpellSchoolFire, 35, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=14541/barovian-family-sword
	// Chance on hit: Deals 30 Shadow damage every 3 sec for 15 sec. All damage done is then transferred to the caster.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(BarovianFamilySword, "Barovian Family Sword", 0.5, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 18652}

		// Keep track of damage taken by each enemy
		enemyDamageTaken := map[int32]float64{}
		for _, target := range character.Env.Encounter.TargetUnits {
			enemyDamageTaken[target.UnitIndex] = 0
		}

		healthMetrics := character.NewHealthMetrics(actionID)

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,

			Dot: core.DotConfig{
				NumberOfTicks: 5,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Siphon Health (Barovian Family Sword)",
				},
				OnSnapshot: func(_ *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					enemyDamageTaken[target.UnitIndex] = 0
					dot.Snapshot(target, 30, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					result := dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
					enemyDamageTaken[target.UnitIndex] += result.Damage
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				enemyDamageTaken[target.UnitIndex] = 0
				spell.Dot(target).Apply(sim)
			},
		})

		// The healing is applied at the end of the DoT and can crit according to old comments
		for _, dot := range spell.Dots() {
			if dot != nil {
				unit := dot.Unit
				dot.ApplyOnExpire(func(aura *core.Aura, sim *core.Simulation) {
					// TODO: This may not be quite correct but it's close enough
					result := spell.CalcDamage(sim, unit, enemyDamageTaken[unit.UnitIndex], spell.OutcomeHealingCrit)
					character.GainHealth(sim, result.Damage, healthMetrics)
				})
			}
		}

		return spell
	})

	// https://www.wowhead.com/classic/item=13204/bashguuder
	// Chance on hit: Punctures target's armor lowering it by 200. Can be applied up to 3 times.
	// 2 PPM from Armaments Discord - same proc as Rivenspike
	itemhelpers.CreateWeaponProcSpell(Bashguuder, "Bashguuder", 2.0, func(character *core.Character) *core.Spell {
		punctureArmorAuras := character.NewEnemyAuraArray(PunctureArmorAura)

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17315},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				activateAura := punctureArmorAuras.Get(target)
				activateAura.Activate(sim)

				if activateAura.IsActive() {
					activateAura.AddStack(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=228606/blackblade-of-shahram
	// Chance on hit: Summons the infernal spirit of Shahram.
	// Summons an NPC "Shahram" who has an equal chance to cast one of 6 spells:
	// Curse of Shahram: -50% movement speed and -25% attack speed on all enemies within 10 yards of Shahram for 10 seconds.
	// Might of Shahram: 5-second stun on all enemies within 10 yards of Shahram.
	// Fist of Shahram: +30% Melee Attack Speed for all party members within 30 yards of Shahram for 8 seconds.
	// Blessing of Shahram: Restores 50 health and mana every 5 seconds for all party members within 30 yards of Shahram for 20 seconds. The Healing portion of this effect scales at 100% of self-healing buffs such as Amplify Magic.
	// Will of Shahram: +50 all stats for yourself for 20 seconds.
	// Flames of Shahram: Deals 100-150 Fire damage to all enemies within 10 yards of Shahram. Damage scales at 100% with +spelldmg debuffs placed on enemies such as Flame Buffet.
	//
	// Implementing this without the guardian as it seems to just cast a spell and depart and guardians are expensive
	// All spells use ProcMaskEmpty because they're not actually cast by the player
	core.NewItemEffect(BlackbladeOfShahram, func(agent core.Agent) {
		character := agent.GetCharacter()

		curseOfShahramAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			aura := target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16597},
				Label:    "Curse of Shahram",
				Duration: time.Second * 10,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.MultiplyAttackSpeed(sim, 1/1.25)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.MultiplyAttackSpeed(sim, 1.25)
				},
			})
			core.AtkSpeedReductionEffect(aura, 1.25)
			return aura
		})

		curseOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16597},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				curseOfShahramAuras.Get(target).Activate(sim)
			},
		})

		mightOfShahramAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16600},
				Label:    "Might of Shahram",
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.Stunned = true
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.Stunned = false
				},
			})
		})

		mightOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16600},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					mightOfShahramAuras.Get(aoeTarget).Activate(sim)
				}
			},
		})

		// This isn't explicit in-game but using a safe value that will likely never be hit
		numFistOfShahramAuras := 8
		fistOfShahramAuras := []*core.Aura{}
		for i := 0; i < numFistOfShahramAuras; i++ {
			fistOfShahramAuras = append(fistOfShahramAuras, character.RegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16601},
				Label:    fmt.Sprintf("Fist of Shahram (%d)", i),
				Duration: time.Second * 8,
			}).AttachMultiplyAttackSpeed(&character.Unit, 1.3))
		}

		fistOfShahram := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16601},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for i := 0; i < numFistOfShahramAuras; i++ {
					if aura := fistOfShahramAuras[i]; !aura.IsActive() {
						aura.Activate(sim)
						break
					}
				}
			},
		})

		blessingOfShahramManaMetrics := character.NewPartyManaMetrics(core.ActionID{SpellID: 16599})
		blessingOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16599},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagIgnoreAttackerModifiers | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			Hot: core.DotConfig{
				Aura: core.Aura{
					Label: "Blessing of Shahram",
				},
				NumberOfTicks: 4,
				TickLength:    time.Second * 5,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, _ bool) {
					dot.SnapshotBaseDamage = 50
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotHealing(sim, target, dot.OutcomeTick)
					if target.HasManaBar() {
						target.AddMana(sim, 50, blessingOfShahramManaMetrics[target.UnitIndex])
					}
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, agent := range character.Party.PlayersAndPets {
					spell.Hot(&agent.GetCharacter().Unit).Apply(sim)
				}
			},
		})

		// This isn't explicit in-game but using a safe value that will likely never be hit
		numWillOfShahramAuras := 8
		willOfShahramAuras := []*core.Aura{}
		willOfShahramStats := stats.Stats{
			stats.Agility:   50,
			stats.Intellect: 50,
			stats.Stamina:   50,
			stats.Spirit:    50,
			stats.Strength:  50,
		}

		for i := 0; i < numWillOfShahramAuras; i++ {
			willOfShahramAuras = append(willOfShahramAuras, character.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16598},
				Label:    fmt.Sprintf("Will of Shahram (%d)", i),
				Duration: time.Second * 20,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					character.AddStatsDynamic(sim, willOfShahramStats)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					character.AddStatsDynamic(sim, willOfShahramStats.Invert())
				},
			}))
		}

		willOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16598},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for i := 0; i < numWillOfShahramAuras; i++ {
					if aura := willOfShahramAuras[i]; !aura.IsActive() {
						aura.Activate(sim)
						break
					}
				}
			},
		})

		flamesOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16596},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagIgnoreAttackerModifiers | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, 90, spell.OutcomeMagicCrit)
				}
			},
		})

		castableSpells := []*core.Spell{curseOfShahram, mightOfShahram, fistOfShahram, blessingOfShahram, willOfShahram, flamesOfShahram}
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Summon Shahram",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				spellIdx := int32(sim.Roll(0, 6))
				castableSpells[spellIdx].Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=12583/blackhand-doomsaw
	// Chance on hit: Wounds the target for 324 to 540 damage.
	// TODO: Proc rate based on the original item
	itemhelpers.CreateWeaponCoHProcDamage(BlackhandDoomsaw, "Blackhand Doomsaw", 0.4, 16549, core.SpellSchoolPhysical, 324, 216, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=17780/blade-of-eternal-darkness
	// Equip: Chance on landing a damaging spell to deal 100 Shadow damage and restore 100 mana to you. (Proc chance: 10%)
	core.NewItemEffect(BladeOfEternalDarkness, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 27860}
		manaMetrics := character.NewManaMetrics(actionID)

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 100, spell.OutcomeAlwaysHit)
				character.AddMana(sim, 100, manaMetrics)
			},
		})

		handler := func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.Damage > 0 {
				procSpell.Cast(sim, character.CurrentTarget)
			}
		}

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Engulfing Shadows",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskSpellDamage,
			ProcChance: .10,
			Handler:    handler,
		})
	})

	// https://www.wowhead.com/classic/item=12777/blazing-rapier
	// Beta client 1.60.1 (spell 16898): 20 a tick every 3 sec, 200 over 30 sec. Era is 10 a tick.
	// 1 PPM Assumed
	itemhelpers.CreateWeaponProcSpell(BlazingRapier, "Blazing Rapier", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16898},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Blaze",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 20, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=11744/bloodfist
	// Chance on hit: Wounds the target for 20 damage.
	// 4 PPM from Armaments Discord
	itemhelpers.CreateWeaponCoHProcDamage(Bloodfist, "Bloodfist", 4, 16433, core.SpellSchoolPhysical, 20, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=9511/bloodletter-scalpel
	itemhelpers.CreateWeaponCoHProcDamage(BloodletterScalpel, "Bloodletter Scalpel", 1.0, 18081, core.SpellSchoolPhysical, 60, 10, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=809/bloodrazor
	itemhelpers.CreateWeaponProcSpell(Bloodrazor, "Bloodrazor", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17504},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Bloodrazor)",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 12, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=12795/blood-talon
	// Chance on hit: Wounds the target causing them to bleed for 100 damage over 30 sec.
	// Assumed 1 PPM
	itemhelpers.CreateWeaponProcSpell(BloodTalon, "Blood Talon", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13318},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Blood Talon)",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 10, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=14487/bonechill-hammer
	// Chance on hit: Blasts a target for 90 Frost damage.
	// 1 PPM from Armaments Discord
	itemhelpers.CreateWeaponCoHProcDamage(BonechillHammer, "Bonechill Hammer", 1.0, 18276, core.SpellSchoolFrost, 90, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17076/bonereavers-edge
	// Chance on hit: Your attacks ignore 700 of your enemies' armor for 10 sec. This effect stacks up to 3 times.
	itemhelpers.CreateWeaponProcSpell(BonereaversEdge, "Bonereaver's Edge", 2.0, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 21153}
		buffAura := character.RegisterAura(core.Aura{
			ActionID:  actionID,
			Label:     "Bonereaver's Edge",
			Duration:  time.Second * 10,
			MaxStacks: 3,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				for _, target := range sim.Encounter.TargetUnits {
					target.AddStatDynamic(sim, stats.Armor, 700*float64(oldStacks))
					target.AddStatDynamic(sim, stats.Armor, -700*float64(newStacks))
				}
			},
		})
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
				buffAura.AddStack(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=2825/bow-of-searing-arrows
	// Beta client 1.60.1 (spell 29638): 23 to 33 Fire damage, up from Era's 18 to 26.
	itemhelpers.CreateWeaponProcSpell(BowOfSearingArrows, "Bow of Searing Arrows", 3.35, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 29638},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeRanged,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(23, 33)
				spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeRangedCritOnly)
			},
		})
	})
	// https://www.wowhead.com/classic/item=13148/chillpike
	// Chance on hit: Blasts a target for 160 to 250 Frost damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponCoHProcDamage(Chillpike, "Chillpike", 1.0, 19260, core.SpellSchoolFrost, 160, 90, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=10761/coldrage-dagger
	// Chance on hit: Launches a bolt of frost at the enemy causing 20 to 30 Frost damage and slowing movement speed by 50% for 5 sec.
	// 2.2 PPM from Armaments Discord
	itemhelpers.CreateWeaponCoHProcDamage(ColdrageDagger, "Coldrage Dagger", 2.2, 13439, core.SpellSchoolFrost, 20, 10, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=13984/darrowspike
	// Chance on hit: Blasts a target for 90 Frost damage.
	// 1 PPM from Armaments Discord
	itemhelpers.CreateWeaponCoHProcDamage(Darrowspike, "Darrowspike", 1.0, 18276, core.SpellSchoolFrost, 90, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17068/deathbringer
	// Chance on hit: Sends a shadowy bolt at the enemy causing 110 to 140 Shadow damage.
	itemhelpers.CreateWeaponCoHProcDamage(Deathbringer, "Deathbringer", 1.0, 18138, core.SpellSchoolShadow, 110, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=647/destiny
	itemhelpers.CreateWeaponProcAura(Destiny, "Destiny", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 17152},
			Label:    "Destiny",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.Strength, 200)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.Strength, -200)
			},
		})
	})

	// https://www.wowhead.com/classic/item=10847/dragons-call
	// Chance on hit: Calls forth an Emerald Dragon Whelp to protect you in battle for a short period of time.
	core.NewItemEffect(DragonsCall, func(agent core.Agent) {
		character := agent.GetCharacter()

		procMask := character.GetProcMaskForItem(DragonsCall)

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Emerald Dragon Whelp Proc",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          procMask,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               1.0, // Reported by armaments discord
			//ICD:               time.Minute * 1,  Removed ICD due to comments on multiple from Classic but am not implementing multiple pets
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				for _, petAgent := range character.PetAgents {
					if whelp, ok := petAgent.(*guardians.EmeraldDragonWhelp); ok {
						whelp.EnableWithTimeout(sim, whelp, time.Second*15)
						break
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=19170/ebon-hand
	// Chance on hit: Sends a shadowy bolt at the enemy causing 125 to 275 Shadow damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponCoHProcDamage(EbonHand, "Ebon Hand", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=14576/ebon-hilt-of-marduk
	// Chance on hit: Corrupts the target, causing 210 damage over 3 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(EbonHiltOfMarduk, "Ebon Hilt of Marduk", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18656},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Corruption (Ebon Hilt of Marduk)",
				},
				TickLength:    time.Second,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 70, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classicaitem=19100/electrified-dagger
	// Chance on hit: Blasts a target for 45 Nature damage.
	// 1.4 PPM from Armaments Discord - assumed same as horde Glacial Dagger
	itemhelpers.CreateWeaponCoHProcDamage(ElectrifiedDagger, "Electrified Dagger", 1.4, 23592, core.SpellSchoolNature, 45, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17112/empyrean-demolisher
	// Chance on hit: Increases your attack speed by 20% for 10 sec.
	itemhelpers.CreateWeaponProcAura(EmpyreanDemolisher, "Empyrean Demolisher", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "Empyrean Demolisher Haste Aura",
			ActionID: core.ActionID{SpellID: 21165},
			Duration: time.Second * 10,
		}).AttachMultiplyAttackSpeed(&character.Unit, 1.2)
	})

	// https://www.wowhead.com/classic/item=10696/enchanted-azsharite-felbane-sword
	// Equip: +33 Attack Power when fighting Demons.
	core.NewMobTypeAttackPowerEffect(EnchantedAzshariteSword, []proto.MobType{proto.MobType_MobTypeDemon}, 33)

	// https://www.wowhead.com/classic/item=18202/eskhandars-left-claw
	// Beta client 1.60.1 (spell 22639): 21 a tick every 1 sec for 5 sec, 105 in all. Era's is
	// 15 every 3 sec for 30 sec, 150 in all - a slower, longer bleed for more total damage.
	// TODO: Proc rate untested
	itemhelpers.CreateWeaponProcSpell(EskhandarsLeftClaw, "Eskhandar's Left Claw", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 22639},
			SpellSchool: core.SpellSchoolPhysical,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Eskhandar's Rake",
				},
				TickLength:    time.Second * 1,
				NumberOfTicks: 5,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 21, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=18203/eskhandars-right-claw
	// Beta client 1.60.1 (spell 22640): 10% for 6 sec, down from Era's 30% for 5 sec.
	itemhelpers.CreateWeaponProcAura(EskhandarsRightClaw, "Eskhandar's Right Claw", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "Eskhandar's Rage",
			ActionID: core.ActionID{SpellID: 22640},
			Duration: time.Second * 6,
		}).AttachMultiplyAttackSpeed(&character.Unit, 1.1)
	})

	// https://www.wowhead.com/classic/item=13218/fang-of-the-crystal-spider
	// Chance on hit: Slows target enemy's casting speed and increases the time between melee and ranged attacks by 10% for 10 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(FangOfTheCrystalSpider, func(agent core.Agent) {
		character := agent.GetCharacter()

		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
			aura := unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 17331},
				Label:    "Fang of the Crystal Spider",
				Duration: time.Second * 10,
			})
			core.AtkSpeedReductionEffect(aura, 1.10)
			return aura
		})

		procMask := character.GetProcMaskForItem(FangOfTheCrystalSpider)
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Fang of the Crystal Spider Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          procMask,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuras.Get(result.Target).Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=12590/felstriker
	// Chance on hit: All attacks are guaranteed to land and will be critical strikes for the next 3 sec.
	core.NewItemEffect(Felstriker, func(agent core.Agent) {
		character := agent.GetCharacter()

		effectAura := character.NewTemporaryStatsAura("Felstriker", core.ActionID{SpellID: 16551}, stats.Stats{stats.MeleeCrit: 100 * core.CritRatingPerCritChance, stats.MeleeHit: 100 * core.MeleeHitRatingPerHitChance}, time.Second*3)
		procMask := character.GetProcMaskForItem(Felstriker)
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Felstriker Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          procMask,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				effectAura.Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=18310/fiendish-machete
	// Equip: +36 Attack Power when fighting Elementals.
	core.NewMobTypeAttackPowerEffect(FiendishMachete, []proto.MobType{proto.MobType_MobTypeElemental}, 36)

	// https://www.wowhead.com/classic/item=870/fiery-war-axe
	itemhelpers.CreateWeaponProcSpell(FieryWarAxe, "Fiery War Axe", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18796},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fiery War Axe Fireball",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(155, 197)
				result := spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(Firebreather, "Firebreather", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16413},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 70, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 3,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Fireball",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 3, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=11809/flame-wrath
	// Chance on hit: Envelops the caster with a Fire shield for 15 sec and shoots a ring of fire dealing 130 to 170 damage to all nearby enemies.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(FlameWrath, "Flame Wrath", 1.0, func(character *core.Character) *core.Spell {
		shieldActionID := core.ActionID{SpellID: 461152}
		shieldSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         shieldActionID,
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 1, // Only the shield portion has scaling
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 10, spell.OutcomeAlwaysHit)
			},
		})
		shieldAura := character.RegisterAura(core.Aura{
			ActionID: shieldActionID,
			Label:    "Flame Wrath",
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.FireResistance, 30)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.FireResistance, -30)
			},
			OnSpellHitTaken: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() {
					shieldSpell.Cast(sim, spell.Unit)
				}
			},
		})
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461151},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				shieldAura.Activate(sim)

				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, sim.Roll(130, 170), spell.OutcomeMagicHit)
				}
			},
		})
	})

	// PPM from Armaments discord
	itemhelpers.CreateWeaponProcSpell(FlurryAxe, "Flurry Axe", 1.9, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 18797},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 18797}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=14024/frightalon
	// Not modelled. Era's proc is "Lowers all attributes of target by 10 for 1 min" and the sim
	// applied it as a stat debuff; the beta client's 19755 reads "Increase the duration of Fear
	// effects applied to the target by 25%", which has nothing to do with damage and nothing the
	// sim can spend. Leaving the old debuff in place would have been a straight invention.

	// https://www.wowhead.com/classic/item=14531/frightskull-shaft
	// Chance on hit: Deals 8 Shadow damage every 2 sec for 30 sec and lowers their Strength for the duration of the disease.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(FrightskullShaft, "Frightskull Shaft", 0.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18633},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot | core.SpellFlagDisease,

			Dot: core.DotConfig{
				NumberOfTicks: 15,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Weakening Disease",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Strength, -50)
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Strength, 50)
					},
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=12797/frostguard#comments
	// Chance on hit: Target's movement slowed by 30% and increasing the time between attacks by 25% for 5 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(Frostguard, func(agent core.Agent) {
		character := agent.GetCharacter()
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
			aura := unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16927},
				Label:    "Chilled (Frostguard)",
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddMoveSpeedModifier(&aura.ActionID, 0.30)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.RemoveMoveSpeedModifier(&aura.ActionID)
				},
			})
			core.AtkSpeedReductionEffect(aura, 1.25)
			return aura
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Frostguard",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMeleeMH,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuras.Get(result.Target).Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19099/glacial-blade
	// Chance on hit: Blasts a target for 45 Frost damage.
	// 1.4 PPM from Armaments Discord
	itemhelpers.CreateWeaponCoHProcDamage(GlacialBlade, "Glacial Blade", 1.4, 18398, core.SpellSchoolFrost, 45, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=13983/gravestone-war-axe
	// Chance on hit: Diseases target enemy for 55 Nature damage every 3 sec for 15 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(GravestoneWarAxe, "Gravestone War Axe", 0.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18289},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagDisease | core.SpellFlagPureDot,

			Dot: core.DotConfig{
				NumberOfTicks: 15,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Creeping Mold",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 55, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=9651/gryphon-riders-stormhammer
	itemhelpers.CreateWeaponCoHProcDamage(GryphonRidersStormhammer, "Gryphon Rider's Stormhammer", 1.0, 18081, core.SpellSchoolNature, 91, 34, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17071/gutgore-ripper
	// Chance on hit: Sends a shadowy bolt at the enemy causing 75 Shadow damage and lowering all stats by 25 for 30 sec.
	itemhelpers.CreateWeaponProcSpell(GutgoreRipper, "Gutgore Ripper", 1.0, func(character *core.Character) *core.Spell {
		procAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 461682},
				Label:    "Gutgore Ripper",
				Duration: time.Second * 30,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   -25,
						stats.Intellect: -25,
						stats.Stamina:   -25,
						stats.Spirit:    -25,
						stats.Strength:  -25,
					})
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   25,
						stats.Intellect: 25,
						stats.Stamina:   25,
						stats.Spirit:    25,
						stats.Strength:  25,
					})
				},
			})
		})

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461682},
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 75, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					procAuras.Get(target).Activate(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=5616/gutwrencher
	itemhelpers.CreateWeaponProcSpell(Gutwrencher, "Gutwrencher", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16406},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Gutwrencher)",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=2164/gut-ripper
	itemhelpers.CreateWeaponCoHProcDamage(GutRipper, "Gut Ripper", 1.0, 18107, core.SpellSchoolPhysical, 95, 26, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=15814/hameyas-slayer
	// Chance on hit: Wounds the target causing them to bleed for 80 damage over 30 sec.
	// Assumed 1 PPM
	itemhelpers.CreateWeaponProcSpell(HameyasSlayer, "Hameya's Slayer", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16406},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Hameya's Slayer)",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 8, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=810/hammer-of-the-northern-wind
	itemhelpers.CreateWeaponCoHProcDamage(HammerOfTheNorthernWind, "Hammer of the Northern Wind", 3.5, 13439, core.SpellSchoolFrost, 20, 10, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=2243/hand-of-edward-the-odd
	// Chance on hit: Next spell cast within 4 sec will cast instantly.
	itemhelpers.CreateWeaponProcAura(HandOfEdwardTheOdd, "Hand of Edward the Odd", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 18803},
			Label:    "Focus (Hand of Edward the Odd)",
			Duration: time.Second * 4,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyCastSpeed(100000)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyCastSpeed(1 / 100000.0)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				aura.Deactivate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=8190/hanzo-sword
	itemhelpers.CreateWeaponCoHProcDamage(HanzoSword, "Hanzo Sword", 1.0, 16405, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=13937/headmasters-charge
	// Use: Gives 20 additional intellect to party members within 30 yards. (10 Min Cooldown)
	// Originally did not stack with Arcane Intellect, but is reported to stack in SoD
	/* core.NewItemEffect(HeadmastersCharge, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 18264}

		buffAura := character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Headmaster's Charge",
			Duration: time.Minute * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.Intellect, 25)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.Intellect, -25)
			},
		})
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 10,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	}) */

	// https://www.wowhead.com/classic/item=11635/hookfang-shanker
	itemhelpers.CreateWeaponProcSpell(HookfangShanker, "Hookfang Shanker", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13526},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Corrosive Poison",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Armor: -50})
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Armor: 50})
					},
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 7, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=11684/ironfoe
	// Chance on hit: Grants 2 extra attacks on your next swing.
	itemhelpers.CreateWeaponProcSpell(Ironfoe, "Ironfoe", 0.8, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 15494},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 2, core.ActionID{SpellID: 15494}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=17054/joonhos-mercy
	itemhelpers.CreateWeaponCoHProcDamage(JoonhosMercy, "Joonho's Mercy", 1.0, 20883, core.SpellSchoolArcane, 70, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=12582/keris-of-zulserak
	// Chance on hit: Inflicts numbing pain that deals 10 Nature damage every 2 sec and increases time between target's attacks by 10% for 10 sec.
	// 1 PPM assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(KerisOfZulSerak, "Keris of Zul'Serak", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16528},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 5,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Numbing Pain",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						core.AtkSpeedReductionEffect(aura, 1.10)
					},
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=11902/linkens-sword-of-mastery
	itemhelpers.CreateWeaponCoHProcDamage(LinkensSwordOfMastery, "Linken's Sword of Mastery", 1.0, 18089, core.SpellSchoolNature, 45, 30, 0, core.DefenseTypeMagic)

	//https://www.wowhead.com/classic/item=19324/the-lobotomizer
	// Chance on hit: Wounds the target for 200 to 300 damage and lowers Intellect of target by 25 for 30 sec.
	// 0.4 PPM from Armaments Disc
	itemhelpers.CreateWeaponProcSpell(TheLobotomizer, "The Lobotomizer", .4, func(character *core.Character) *core.Spell {
		procAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 24388},
				Label:    "Brain Damage",
				Duration: time.Second * 30,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Intellect: -25,
					})
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Intellect: 25,
					})
				},
			})
		})

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 24388},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				damageRoll := sim.Roll(200, 300)
				result := spell.CalcAndDealDamage(sim, target, damageRoll, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					procAuras.Get(target).Activate(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=11817/lord-generals-sword
	// Beta client 1.60.1 (spell 15602): 84 attack power for 10 sec, up from Era's 50 for 30 sec.
	// // TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(LordGeneralsSword, "Lord General's Sword", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 15602},
			Label:    "Lord General's Sword",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.AttackPower:       84,
					stats.RangedAttackPower: 84,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.AttackPower:       -84,
					stats.RangedAttackPower: -84,
				})
			},
		})
	})

	// https://www.wowhead.com/classic/item=9449/manual-crowd-pummeler

	core.NewItemEffect(ManualCrowdPummeler, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 13494}
		duration := time.Second * 30

		mcpAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Haste",
			ActionID: actionID,
			Duration: duration,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.5)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.0/1.5)
			},
		})

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Second * 30,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				mcpAura.Activate(sim)
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=12794/masterwork-stormhammer
	// Chance on hit: Blasts up to 3 targets for 105 to 145 Nature damage.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(MasterworkStormhammer, "Masterwork Stormhammer", 0.5, func(character *core.Character) *core.Spell {
		maxHits := int(min(3, character.Env.GetNumTargets()))
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 463946},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for numHits := 0; numHits < maxHits; numHits++ {
					spell.CalcAndDealDamage(sim, target, sim.Roll(105, 145), spell.OutcomeMagicHitAndCrit)
					target = character.Env.NextTargetUnit(target)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=1982/nightblade
	itemhelpers.CreateWeaponCoHProcDamage(Nightblade, "Nightblade", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=19169/nightfall
	// Chance on hit: Spell damage taken by target increased by 15% for 5 sec.
	core.NewItemEffect(Nightfall, func(agent core.Agent) {
		character := agent.GetCharacter()

		procAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				Label:    "Spell Vulnerability",
				ActionID: core.ActionID{SpellID: 23605},
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexArcane] *= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] *= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFrost] *= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] *= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexNature] *= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexShadow] *= 1.15
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexArcane] /= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] /= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFrost] /= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] /= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexNature] /= 1.15
					aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexShadow] /= 1.15
				},
			})
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Nightfall Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               2,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procAuras.Get(result.Target).Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=9425/pendulum-of-doom
	itemhelpers.CreateWeaponCoHProcDamage(PendulumOfDoom, "Pendulum of Doom", 0.5, 10373, core.SpellSchoolPhysical, 250, 100, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=12709/pips-skinner
	core.NewMobTypeAttackPowerEffect(PipsSkinner, []proto.MobType{proto.MobType_MobTypeBeast}, 45)

	// https://www.wowhead.com/classic/item=18816/perditions-blade
	// Chance on hit: Blasts a target for 40 to 56 Fire damage.
	itemhelpers.CreateWeaponCoHProcDamage(PerditionsBlade, "Perdition's Blade", 2.8, 23267, core.SpellSchoolFire, 40, 56, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=18348/quelserrar
	// Chance on hit: When active, grants the wielder 13 defense and 300 armor for 10 sec.
	// Proc rate estimated based on data from WoW Armaments Discord for the original item
	itemhelpers.CreateWeaponProcAura(QuelSerrar, "Quel'Serrar", 2.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 463105},
			Label:    "Sanctuary",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Defense:    13,
					stats.BonusArmor: 300,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Defense:    -13,
					stats.BonusArmor: -300,
				})
			},
		})
	})

	// https://www.wowhead.com/classic/item=10626/ragehammer
	// Beta client 1.60.1 (spell 12686): "Increases Attack Power by 102 and melee attack speed
	// by 6% for 15 sec". Era's reads "Increases damage done by 20", which is why this was flat
	// physical damage - Forever changed the wording along with the number, so 102 is attack
	// power and not 102 damage a swing.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(Ragehammer, "Ragehammer", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 12686},
			Label:    "Enrage (12686)",
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{stats.AttackPower: 102, stats.RangedAttackPower: 102})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{stats.AttackPower: -102, stats.RangedAttackPower: -102})
			},
		}).AttachMultiplyAttackSpeed(&character.Unit, 1.06)
	})

	// https://www.wowhead.com/classic/item=7717/ravager
	itemhelpers.CreateWeaponProcAura(Ravager, "Ravager", 1.0, func(character *core.Character) *core.Aura {
		tickActionID := core.ActionID{SpellID: 9633}
		procActionID := core.ActionID{SpellID: 9632}
		//Used as part of a canceling ravager APL
		auraActionID := core.ActionID{SpellID: 433801}

		ravegerBladestormTickSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    tickActionID,
			SpellSchool: core.SpellSchoolPhysical,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskMeleeMHSpecial,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			BonusCoefficient: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				damage := 5.0 + spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target))
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, damage, spell.OutcomeMeleeSpecialHitAndCrit)
				}
			},
		})

		character.GetOrRegisterSpell(core.SpellConfig{
			SpellSchool: core.SpellSchoolPhysical,
			ActionID:    procActionID,
			ProcMask:    core.ProcMaskMeleeMHSpecial,
			Flags:       core.SpellFlagChanneled,
			Dot: core.DotConfig{
				IsAOE: true,
				Aura: core.Aura{
					Label: "Ravager Whirlwind",
				},
				NumberOfTicks:       3,
				TickLength:          time.Second * 3,
				AffectedByCastSpeed: false,
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					ravegerBladestormTickSpell.Cast(sim, target)
				},
			},
		})

		return character.GetOrRegisterAura(core.Aura{
			Label:    "Ravager Bladestorm",
			ActionID: auraActionID,
			Duration: time.Second * 9,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AutoAttacks.CancelAutoSwing(sim)
				dotSpell := character.GetSpell(procActionID)
				dotSpell.AOEDot().Apply(sim)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AutoAttacks.EnableAutoSwing(sim)
				dotSpell := character.GetSpell(procActionID)
				dotSpell.AOEDot().Cancel(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=13286/rivenspike
	// Beta client 1.60.1 (spell 17315): 100 armor a stack, halved from Era's 200. Still 3 stacks.
	// 2 PPM - Armaments Discord has 1 PPM recorded before it could record refreshes.  Bashguuder with same effect is recorded at 2PPM so setting to match
	itemhelpers.CreateWeaponProcSpell(Rivenspike, "Rivenspike", 2.0, func(character *core.Character) *core.Spell {
		punctureArmorAuras := character.NewEnemyAuraArray(PunctureArmorAura)

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17315},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				activateAura := punctureArmorAuras.Get(target)
				activateAura.Activate(sim)

				if activateAura.IsActive() {
					activateAura.AddStack(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=13505/runeblade-of-baron-rivendare
	// Equip: Increases movement speed and life regeneration rate.
	// TODO: Movement speed not implemented
	core.NewItemEffect(RunebladeOfBaronRivendare, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 17625}
		healthMetrics := character.NewHealthMetrics(actionID)
		character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Unholy Aura",
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   time.Second * 5,
					Priority: core.ActionPriorityAuto,
					OnAction: func(sim *core.Simulation) {
						character.GainHealth(sim, 20, healthMetrics)
					},
				})
			},
		})
	})

	// https://www.wowhead.com/classic/item=17752/satyrs-lash
	itemhelpers.CreateWeaponCoHProcDamage(SatyrsLash, "Satyr's Lash", 1.0, 18205, core.SpellSchoolShadow, 55, 30, 0, core.DefenseTypeMagic)

	// TODO Searing Needle adds an "Apply Aura: Mod Damage Done (Fire): 10" aura to the /target/, buffing it; not currently modelled
	itemhelpers.CreateWeaponCoHProcDamage(SearingNeedle, "Searing Needle", 1.0, 16454, core.SpellSchoolFire, 60, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=12969/seeping-willow
	// Beta client 1.60.1 (spell 17196): "Deals 51 Nature damage every 2 sec to all enemies within
	// an 8 yard radius of the caster for 10 sec" - 5 ticks for 255. Era's also lowered all stats
	// by 20 and ran 20 a tick every 3 sec for 30 sec; the stat debuff is gone entirely.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(SeepingWillow, "Seeping Willow", 0.5, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 17196},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Seeping Willow Poison",
				},
				NumberOfTicks: 5,
				TickLength:    time.Second * 2,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 51, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					result := spell.CalcAndDealOutcome(sim, aoeTarget, spell.OutcomeMagicHit)
					if result.Landed() {
						spell.Dot(aoeTarget).Apply(sim)
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=13035/serpent-slicer
	// Chance on hit: Poisons target for 8 Nature damage every 2 sec for 20 sec.
	itemhelpers.CreateWeaponProcSpell(SerpentSlicer, "Serpent Slicer", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17511},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Poison (Serpent Slicer)",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=2163/shadowblade
	itemhelpers.CreateWeaponCoHProcDamage(Shadowblade, "Shadowblade", 1.0, 18138, core.SpellSchoolShadow, 110, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17074/shadowstrike
	// Chance on hit: Steals 100 to 180 life from target enemy.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(Shadowstrike, "Shadowstrike", 2.2, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 461683}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: 1.0,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(100, 180), spell.OutcomeMagicHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	// https://www.wowhead.com/classic/item=754/shortsword-of-vengeance
	itemhelpers.CreateWeaponCoHProcDamage(ShortswordOfVengeance, "Shortsword of Vengeance", 1.0, 13519, core.SpellSchoolHoly, 30, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=13361/skullforge-reaver
	// Beta client 1.60.1 (spell 17484): 21 a tick every 1 sec for 10 sec, 210 in all. Era is 2 a
	// tick for 30 sec, 60 in all.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(SkullforgeReaver, "Skullforge Reaver", 1.7, func(character *core.Character) *core.Spell {
		procMask := character.GetProcMaskForItem(SkullforgeReaver)
		actionID := core.ActionID{SpellID: 17484}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    procMask,
			Flags:       core.SpellFlagPureDot,
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second,
				Aura: core.Aura{
					Label: "Skullforge Brand",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 21, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					result := dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
					character.GainHealth(sim, result.Damage, healthMetrics)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=17182/sulfuras-hand-of-ragnaros
	// Chance on hit: Hurls a fiery ball that causes 273 to 333 Fire damage and an additional 75 damage over 10 sec.
	// Equip: Deals 5 Fire damage to anyone who strikes you with a melee attack.
	core.NewItemEffect(SulfurasHandOfRagnaros, func(agent core.Agent) {
		character := agent.GetCharacter()

		immolationActionID := core.ActionID{SpellID: 21142}

		immolationSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    immolationActionID,
			SpellSchool: core.SpellSchoolFire,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 5, spell.OutcomeMagicHit)
			},
		})

		character.GetOrRegisterAura(core.Aura{
			ActionID: immolationActionID,
			Label:    "Immolation (Hand of Ragnaros)",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) {
					immolationSpell.Cast(sim, spell.Unit)
				}
			},
		})

		fireballSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 21162},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fireball (Hand of Ragnaros)",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 5,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 15, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(273, 333), spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Hand of Ragnaros Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				fireballSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=17193/sulfuron-hammer
	// Chance on hit: Hurls a fiery ball that causes 83 to 101 Fire damage and an additional 16 damage over 8 sec.
	core.NewItemEffect(SulfuronHammer, func(agent core.Agent) {
		character := agent.GetCharacter()

		fireballSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 21159},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fireball (Sulfuron Hammer)",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 4,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 4, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(83, 101), spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Sulfuron Hammer Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1, // TODO: Armaments Discord didn't have any data on Sulfuron Hammer
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				fireballSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=6622/sword-of-zeal
	// Chance on hit: A burst of energy fills the caster, increasing his damage by 10 and armor by 150 for 15 sec.
	// 1.8 PPM from Armaments discord
	itemhelpers.CreateWeaponProcAura(SwordOfZeal, "Sword of Zeal", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 8191},
			Label:    "Zeal",
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.BonusPhysicalDamage += 10
				character.AddStatsDynamic(sim, stats.Stats{stats.Armor: 150})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.BonusPhysicalDamage -= 10
				character.AddStatsDynamic(sim, stats.Stats{stats.Armor: -150})
			},
		})
	})

	// https://www.wowhead.com/classic/item=1728/teebus-blazing-longsword
	// Chance on hit: Blasts a target for 150 Fire damage.
	// Chance on Hit Assumed: 1 PPM
	itemhelpers.CreateWeaponCoHProcDamage(TeebusBlazingLongsword, "Teebu's Blazing Longsword", 1.0, 18086, core.SpellSchoolFire, 150, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=13401/the-cruel-hand-of-timmy
	// Beta client 1.60.1 (spell 17505): "Steals 210 life from target enemy". Era's is "Lowers all
	// attributes of target by 15 for 1 min", so this is a different ability rather than a retune -
	// the attribute debuff is gone and it deals Shadow damage instead.
	// 0.65 PPM from Armaments Discord
	itemhelpers.CreateWeaponProcSpell(TheCruelHandOfTimmy, "The Cruel Hand of Timmy", 0.65, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 17505}
		healthMetrics := character.NewHealthMetrics(actionID)

		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 0,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 210, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	// https://www.wowhead.com/classic/item=9639/the-hand-of-antusul
	itemhelpers.CreateWeaponProcSpell(TheHandOfAntusul, "The Hand of Antu'sul", 1.0, func(character *core.Character) *core.Spell {
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
			aura := unit.GetOrRegisterAura(core.Aura{
				Label:    "ThunderClap-Antu'sul",
				ActionID: core.ActionID{SpellID: 13532},
				Duration: time.Second * 10,
			})
			core.AtkSpeedReductionEffect(aura, 1.11)
			return aura
		})

		results := make([]*core.SpellResult, min(4, character.Env.GetNumTargets()))

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13532},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for idx := range results {
					results[idx] = spell.CalcDamage(sim, target, 7, spell.OutcomeMagicHitAndCrit)
					target = character.Env.NextTargetUnit(target)
				}
				for _, result := range results {
					spell.DealDamage(sim, result)
					if result.Landed() {
						debuffAuras.Get(result.Target).Activate(sim)
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=9423/the-jackhammer
	itemhelpers.CreateWeaponProcAura(TheJackhammer, "The Jackhammer", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "The Jackhammer Haste Aura",
			ActionID: core.ActionID{SpellID: 13533},
			Duration: time.Second * 10,
			// Beta client 1.60.1 (spell 13533): 15%, halved from Era's 30%.
		}).AttachMultiplyAttackSpeed(&character.Unit, 1.15)
	})

	// https://www.wowhead.com/classic/item=13060/the-needler
	itemhelpers.CreateWeaponCoHProcDamage(TheNeedler, "The Needler", 3.0, 13060, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=17705/thrash-blade
	itemhelpers.CreateWeaponProcSpell(ThrashBlade, "Thrash Blade", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 21919},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 21919}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19019/thunderfury-blessed-blade-of-the-windseeker
	// Chance on hit: Blasts your enemy with lightning, dealing 300 Nature damage and then jumping to additional nearby enemies.
	// Each jump reduces that victim's Nature resistance by 25. Affects 5 targets.
	// Your primary target is also consumed by a cyclone, slowing its attack speed by 20% for 12 sec.
	core.NewItemEffect(Thunderfury, func(agent core.Agent) {
		character := agent.GetCharacter()

		procMask := character.GetProcMaskForItem(Thunderfury)
		ppmm := character.AutoAttacks.NewPPMManager(6.0, procMask)
		thunderfuryASAuras := character.NewEnemyAuraArray(core.ThunderfuryASAura)
		procActionID := core.ActionID{SpellID: 21992}

		singleTargetSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    procActionID.WithTag(1),
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellProc | core.ProcMaskSpellDamageProc,
			Flags:       core.SpellFlagIgnoreAttackerModifiers,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			FlatThreatBonus:  126,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 300, spell.OutcomeMagicHitAndCrit)
				thunderfuryASAuras.Get(target).Activate(sim)
			},
		})

		debuffAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				Label:    "Thunderfury",
				ActionID: procActionID,
				Duration: time.Second * 12,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					target.AddStatDynamic(sim, stats.NatureResistance, -25)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					target.AddStatDynamic(sim, stats.NatureResistance, 25)
				},
			})
		})

		results := make([]*core.SpellResult, min(5, character.Env.GetNumTargets()))

		bounceSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    procActionID.WithTag(2),
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskEmpty,

			ThreatMultiplier: 1,
			FlatThreatBonus:  126,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for idx := range results {
					results[idx] = spell.CalcDamage(sim, target, 0, spell.OutcomeMagicHit)
					target = sim.Environment.NextTargetUnit(target)
				}
				for _, result := range results {
					if result.Landed() {
						debuffAuras[result.Target.Index].Activate(sim)
					}
					spell.DealDamage(sim, result)
				}
			},
		})

		core.MakePermanent(character.GetOrRegisterAura(core.Aura{
			Label: "Thunderfury Trigger",
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && ppmm.Proc(sim, spell.ProcMask, "Thunderfury") {
					singleTargetSpell.Cast(sim, result.Target)
					bounceSpell.Cast(sim, result.Target)
				}
			},
		}))
	})

	// https://www.wowhead.com/classic/item=17223/thunderstrike
	// Chance on hit: Blasts up to 3 targets for 150 to 250 Nature damage. Each target after the first takes less damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(Thunderstrike, "Thunderstrike", 1.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461686},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				initialResult := spell.CalcOutcome(sim, target, spell.OutcomeMagicHit)
				// Only the initial hit can be fully resisted according to a wowhead comment
				if initialResult.Landed() {
					damageMultiplier := 1.0
					for numHits := 0; numHits < 3; numHits++ {
						spell.CalcAndDealDamage(sim, target, sim.Roll(150, 250)*damageMultiplier, spell.OutcomeMagicCrit)
						numHits++
						target = character.Env.NextTargetUnit(target)
						// TODO: Couldn't find information on what the multiplier actually is
						damageMultiplier *= .65
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=13183/venomspitter
	// Chance on hit: Poisons target for 7 Nature damage every 2 sec for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(Venomspitter, "Venomspitter", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18203},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Poison (Venomspitter)",
				},
				TickLength:    time.Second * 2,
				NumberOfTicks: 15,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 7, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=11603/vilerend-slicer
	itemhelpers.CreateWeaponCoHProcDamage(VilerendSlicer, "Vilerend Slicer", 1.0, 16405, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=17075/viskag-the-bloodletter
	itemhelpers.CreateWeaponCoHProcDamage(ViskagTheBloodletter, "Vis'kag the Bloodletter", 0.6, 21140, core.SpellSchoolPhysical, 240, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=12792/volcanic-hammer
	// Beta client 1.60.1 (spell 18082): 200 to 256 Fire damage and 12 a tick over 6 sec, both
	// doubled from Era's 100 to 128 and 6 a tick.
	// Assumed 1 PPM
	itemhelpers.CreateWeaponProcSpell(VolcanicHammer, "Volcanic Hammer", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18082},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fireball (Volcanic Hammer)",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 12, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(200, 256)
				result := spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=11920/wraith-scythe
	// Beta client 1.60.1 (spell 16414): steals 49 life, and the 0.3 spell power coefficient Era
	// carried is gone.
	itemhelpers.CreateWeaponProcSpell(WraithScythe, "Wraith Scythe", 1.0, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 16414}
		healthMetrics := character.NewHealthMetrics(actionID)

		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 0,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 49, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	///////////////////////////////////////////////////////////////////////////
	//                                 Trinkets
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=11832/burst-of-knowledge
	// Use: Reduces mana cost of all spells by 100 for 10 sec. (5 Min Cooldown)
	core.NewItemEffect(BurstOfKnowledge, func(agent core.Agent) {
		character := agent.GetCharacter()

		aura := character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{ItemID: BurstOfKnowledge},
			Label:    "Burst of Knowledge",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range aura.Unit.Spellbook {
					if spell.Cost != nil && spell.Cost.CostType() == core.CostTypeMana {
						spell.Cost.FlatModifier -= 100
					}
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range aura.Unit.Spellbook {
					if spell.Cost != nil && spell.Cost.CostType() == core.CostTypeMana {
						spell.Cost.FlatModifier += 100
					}
				}
			},
		})

		spell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{ItemID: BurstOfKnowledge},
			ProcMask: core.ProcMaskEmpty,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 15,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				aura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeMana,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=22268/draconic-infused-emblem
	// Use: Increases your spell damage by up to 100 and your healing by up to 190 for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(DraconicInfusedEmblem, stats.Stats{stats.SpellDamage: 100, stats.HealingPower: 190}, time.Second*15, time.Second*90)

	// https://www.wowhead.com/classic/item=19288/darkmoon-card-blue-dragon
	// Equip: 2% chance on successful spellcast to allow 100% of your Mana regeneration to continue while casting for 15 sec.
	core.NewItemEffect(DarkmoonCardBlueDragon, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23688}

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Aura of the Blue Dragon",
			ActionID: actionID,
			Duration: time.Second * 15,
		}).AttachSpiritRegenRateCastingBuff(1)

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Aura of the Blue Dragon Trigger",
			Callback:   core.CallbackOnCastComplete,
			ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
			ProcChance: .02,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procAura.Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19287/darkmoon-card-heroism
	// Equip: Sometimes heals bearer of 120 to 180 damage when damaging an enemy in melee.
	core.NewItemEffect(DarkmoonCardHeroism, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23689}
		healthMetrics := character.NewHealthMetrics(actionID)

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolHoly,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.GainHealth(sim, sim.Roll(120, 180), healthMetrics)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Heroism Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               2,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19289/darkmoon-card-maelstrom
	// Equip: Chance to strike your melee target with lightning for 200 to 300 Nature damage.
	core.NewItemEffect(DarkmoonCardMaelstrom, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23686}

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(200, 300), spell.OutcomeMagicHitAndCrit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Lightning Strike Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               1.0,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=18815/essence-of-the-pure-flame
	// Equip: When struck in combat inflicts 13 Fire damage to the attacker.
	core.NewItemEffect(EssenceOfThePureFlame, func(agent core.Agent) {
		character := agent.GetCharacter()

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 23266},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 13, spell.OutcomeAlwaysHit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Fiery Aura Proc",
			Callback: core.CallbackOnSpellHitTaken,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee, // TODO: Unsure if this means melee attacks or all attacks
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})
	})

	// Forever halves it. Era's tooltip hardcodes "2% chance on melee hit" and SpellAuraOptions
	// holds 2 for 15600; Forever's reads "${$h/3}% chance on Melee hit to gain $s1 extra attack.
	// Attacks against Dwarves are $s2 times as likely", with the chance field at 3 and $s2 at 3.
	// So the field is still a percent and the division is new: 1% here, 3% against a Dwarf, which
	// a raid boss never is.
	core.NewItemEffect(HandOfJustice, func(agent core.Agent) {
		character := agent.GetCharacter()
		if !character.AutoAttacks.AutoSwingMelee {
			return
		}

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 2,
		}

		character.GetOrRegisterAura(core.Aura{
			Label:    "Hand of Justice",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell.Flags.Matches(core.SpellFlagSuppressEquipProcs) {
					return
				}
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) && icd.IsReady(sim) && sim.Proc(0.01, "HandOfJustice") {
					icd.Use(sim)
					aura.Unit.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 15600}, spell)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=22321/heart-of-wyrmthalak
	// Equip: Chance to bathe your melee target in flames for 120 to 180 Fire damage.
	// TODO: Proc rate assumed from a wowhead comment and needs testing
	core.NewItemEffect(HeartOfWyrmthalak, func(agent core.Agent) {
		character := agent.GetCharacter()
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 27656},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(120, 180), spell.OutcomeMagicHitAndCrit)
			},
		})
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Heart of Wyrmthalak Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               0.4,
			Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				spell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=17774/mark-of-the-chosen
	core.NewItemEffect(MarkOfTheChosen, func(agent core.Agent) {
		character := agent.GetCharacter()
		// Beta client 1.60.1 (spell 21970): 21 to all stats, down from Era's 25.
		statIncrease := float64(21)
		markProcChance := 0.02

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Mark of the Chosen Effect",
			ActionID: core.ActionID{SpellID: 21970},
			Duration: time.Minute,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:   statIncrease,
					stats.Agility:   statIncrease,
					stats.Strength:  statIncrease,
					stats.Intellect: statIncrease,
					stats.Spirit:    statIncrease,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:   -statIncrease,
					stats.Agility:   -statIncrease,
					stats.Strength:  -statIncrease,
					stats.Intellect: -statIncrease,
					stats.Spirit:    -statIncrease,
				})
			},
		})

		core.MakePermanent(character.GetOrRegisterAura(core.Aura{
			Label: "Mark of the Chosen",
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) && sim.RandomFloat("Mark of the Chosen") < markProcChance {
					procAura.Activate(sim)
				}
			},
		}))
	})

	// https://www.wowhead.com/classic/item=19812/rune-of-the-dawn
	// Equip: Increases damage done to Undead by magical spells and effects by up to 48.
	core.NewMobTypeSpellPowerEffect(RuneOfTheDawn, []proto.MobType{proto.MobType_MobTypeUndead}, 48)

	// https://www.wowhead.com/classic/item=11819/second-wind
	// Beta client 1.60.1 (spell 15604): 63 mana a tick, up from Era's 30. The tooltip's
	// "doubled in Mountainous areas" has nothing to sim against.
	core.NewItemEffect(SecondWind, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 15604}
		manaMetrics := character.NewManaMetrics(actionID)
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			ProcMask: core.ProcMaskEmpty,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 15,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   time.Second * 1,
					NumTicks: 10,
					Priority: core.ActionPriorityAuto,
					OnAction: func(sim *core.Simulation) {
						character.AddMana(sim, 63, manaMetrics)
					},
				})
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=13213/smolderwebs-eye
	// Not modelled. Era's use-effect is "Poisons target for 20 Nature damage every 2 sec for
	// 20 sec" and the sim registered it as a DPS major cooldown; the beta client's 17330 reads
	// "Web yourself to the ground, preventing jumping and negating the next Knockback effect
	// taken", which is a self-debuff with no damage in it at all.

	// https://www.wowhead.com/classic/item=13209/seal-of-the-dawn
	// Equip: +81 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(SealOfTheDawn, []proto.MobType{proto.MobType_MobTypeUndead}, 81)

	// https://www.wowhead.com/classic/item=23078/gauntlets-of-undead-slaying
	// +60 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(GauntletsOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 60)

	// https://www.wowhead.com/classic/item=23081/handwraps-of-undead-slaying
	// +60 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(HandwrapsOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 60)

	// https://www.wowhead.com/classic/item=23082/handguards-of-undead-slaying
	// +60 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(HandguardsOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 60)

	// https://www.wowhead.com/classic/item=23084/gloves-of-undead-cleansing
	// Increases damage done to Undead by magical spells and effects by up to 35.
	core.NewMobTypeSpellPowerEffect(GlovesOfUndeadCleansing, []proto.MobType{proto.MobType_MobTypeUndead}, 35)

	// https://www.wowhead.com/classic/item=23085/robe-of-undead-cleansing
	// Increases damage done to Undead by magical spells and effects by up to 48.
	core.NewMobTypeSpellPowerEffect(RobeOfUndeadCleansing, []proto.MobType{proto.MobType_MobTypeUndead}, 48)

	// https://www.wowhead.com/classic/item=23087/breastplate-of-undead-slaying
	// +81 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(BreastplateOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 81)

	// https://www.wowhead.com/classic/item=23088/chestguard-of-undead-slaying
	// +81 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(ChestguardOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 81)

	// https://www.wowhead.com/classic/item=23089/tunic-of-undead-slaying
	// +81 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(TunicOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 81)

	// https://www.wowhead.com/classic/item=23090/bracers-of-undead-slaying
	// +45 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(BracersOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 45)

	// https://www.wowhead.com/classic/item=23091/bracers-of-undead-cleansing
	// Increases damage done to Undead by magical spells and effects by up to 26.
	core.NewMobTypeSpellPowerEffect(BracersOfUndeadCleansing, []proto.MobType{proto.MobType_MobTypeUndead}, 26)

	// https://www.wowhead.com/classic/item=23092/wristguards-of-undead-slaying
	// +45 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(WristguardsOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 45)

	// https://www.wowhead.com/classic/item=23093/wristwraps-of-undead-slaying
	// +45 Attack Power when fighting Undead.
	core.NewMobTypeAttackPowerEffect(WristwrapsOfUndeadSlaying, []proto.MobType{proto.MobType_MobTypeUndead}, 45)

	// https://www.wowhead.com/classic/item=18820/talisman-of-ephemeral-power
	// Use: Increases damage and healing done by magical spells and effects by up to 175 for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(TalismanOfEphemeralPower, stats.Stats{stats.SpellPower: 175}, time.Second*15, time.Second*90)

	///////////////////////////////////////////////////////////////////////////
	//                                 Other
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=17111/blazefury-medallion
	// Equip: Adds 2 fire damage to your melee attacks.
	core.NewItemEffect(BlazefuryMedallion, func(agent core.Agent) {
		character := agent.GetCharacter()

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 7712},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskMeleeDamageProc,
			Flags:            core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 2, spell.OutcomeMagicCrit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Blazefury Medallion Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
					procSpell.ProcMask = core.ProcMaskEmpty
				} else {
					procSpell.ProcMask = core.ProcMaskDamageProc // Both spell and melee procs
				}
				procSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=14554/cloudkeeper-legplates
	// Use: Increases Attack Power by 100 for 30 sec. (15 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(CloudkeeperLegplates, stats.Stats{stats.AttackPower: 100, stats.RangedAttackPower: 100}, time.Second*30, time.Minute*15)

	// https://www.wowhead.com/classic/item=228266/drillborer-disk
	// Equip: When struck in combat inflicts 3 Arcane damage to the attacker.
	core.NewItemEffect(DrillborerDisk, func(agent core.Agent) {
		thornsArcaneDamageEffect(agent, DrillborerDisk, "Drillborer Disk", 3)
	})

	// https://www.wowhead.com/classic/item=18168/force-reactive-disk
	// Equip: When the shield blocks it releases an electrical charge that damages all nearby enemies. (1s cooldown)
	core.NewItemEffect(ForceReactiveDisk, func(agent core.Agent) {
		character := agent.GetCharacter()

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{ItemID: ForceReactiveDisk},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, 25, spell.OutcomeMagicHitAndCrit)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Force Reactive Disk",
			Callback: core.CallbackOnSpellHitTaken,
			ProcMask: core.ProcMaskMelee,
			Outcome:  core.OutcomeBlock,
			ICD:      time.Second,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})
	})

	// https://www.wowhead.com/classic/item=11669/naglering
	// Equip: When struck in combat inflicts 3 Arcane damage to the attacker.
	core.NewItemEffect(Naglering, func(agent core.Agent) {
		thornsArcaneDamageEffect(agent, Naglering, "Naglering", 3)
	})

	// https://www.wowhead.com/classic/item=18326/razor-gauntlets
	// Equip: When struck in combat inflicts 3 Arcane damage to the attacker.
	core.NewItemEffect(RazorGauntlets, func(agent core.Agent) {
		thornsArcaneDamageEffect(agent, RazorGauntlets, "Razor Gauntlets", 3)
	})

	// https://www.wowhead.com/classic/item=1168/skullflame-shield
	// Beta client 1.60.1 (spell 18817): steals 270 life, not Era's 35, and the spell power
	// coefficient Era carried is gone.
	// Equip: When struck in combat has a 1% chance of dealing 75 to 125 Fire damage to all targets around you. (Proc chance: 1%)
	core.NewItemEffect(SkullflameShield, func(agent core.Agent) {
		character := agent.GetCharacter()

		drainLifeActionID := core.ActionID{SpellID: 18817}
		healthMetrics := character.NewHealthMetrics(drainLifeActionID)
		drainLifeSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    drainLifeActionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: 0,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 270, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})

		flamestrikeSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18818},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					// Beta client 1.60.1 (spell 18818): 90 to 140, up from Era's 75 to 125.
					spell.CalcAndDealDamage(sim, aoeTarget, sim.Roll(90, 140), spell.OutcomeMagicHit)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Drain Life Trigger",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: 0.03,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				drainLifeSpell.Cast(sim, spell.Unit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Flamestrike Trigger",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: 0.01,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				flamestrikeSpell.Cast(sim, spell.Unit)
			},
		})
	})

	core.AddEffectsToTest = true
}

func thornsArcaneDamageEffect(agent core.Agent, itemID int32, itemName string, damage float64) {
	character := agent.GetCharacter()

	procSpell := character.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{ItemID: itemID},
		SpellSchool: core.SpellSchoolArcane,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagBinary | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicHit)
		},
	})

	core.MakePermanent(character.GetOrRegisterAura(core.Aura{
		Label: fmt.Sprintf("Thorns (%s)", itemName),
		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) {
				procSpell.Cast(sim, spell.Unit)
			}
		},
	}))
}

var minorArmorReductionEffectCategory = "MinorArmorReduction"

func PunctureArmorAura(target *core.Unit) *core.Aura {
	// Beta client 1.60.1 (spell 17315): 100, halved from Era's 200.
	arpen := float64(100)

	var effect *core.ExclusiveEffect
	aura := target.GetOrRegisterAura(core.Aura{
		Label:     "Puncture Armor",
		ActionID:  core.ActionID{SpellID: 17315},
		Duration:  time.Second * 30,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			effect.SetPriority(sim, arpen*float64(newStacks))
		},
	})

	effect = aura.NewExclusiveEffect(minorArmorReductionEffectCategory, true, core.ExclusiveEffect{
		Priority: 0,
		OnGain: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.AddStatDynamic(sim, stats.Armor, -ee.Priority)
		},
		OnExpire: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.AddStatDynamic(sim, stats.Armor, ee.Priority)
		},
	})

	return aura
}

func ArmorShatterAuras(target *core.Unit) *core.Aura {
	// Beta client 1.60.1 (spell 16928): 165, down from Era's 200.
	arpen := float64(165)

	var effect *core.ExclusiveEffect
	aura := target.GetOrRegisterAura(core.Aura{
		Label:     "Armor Shatter",
		ActionID:  core.ActionID{SpellID: 16928},
		Duration:  time.Second * 45,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			effect.SetPriority(sim, arpen*float64(newStacks))
		},
	})

	effect = aura.NewExclusiveEffect(minorArmorReductionEffectCategory, true, core.ExclusiveEffect{
		Priority: 0,
		OnGain: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.AddStatDynamic(sim, stats.Armor, -ee.Priority)
		},
		OnExpire: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.AddStatDynamic(sim, stats.Armor, ee.Priority)
		},
	})

	return aura
}
