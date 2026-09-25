package paladin

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func (paladin *Paladin) ApplyTalents() {
	// Precision reads "all spells and attacks" under Forever, where Classic's only covered melee.
	paladin.AddStat(stats.MeleeHit, float64(paladin.Talents.Precision)*core.MeleeHitRatingPerHitChance)
	paladin.AddStat(stats.SpellHit, float64(paladin.Talents.Precision)*core.SpellHitRatingPerHitChance)

	paladin.AddStat(stats.MeleeCrit, float64(paladin.Talents.Conviction)*core.CritRatingPerCritChance)
	// TODO: paladin.AddStat(stats.RangedCrit, float64(paladin.Talents.Conviction)*core.CritRatingPerCritChance)

	// Divine Precision: 6/12/18%, confirmed by the beta client's talent curve.
	paladin.PseudoStats.SchoolBonusHitChance[stats.SchoolIndexHoly] += 6 * float64(paladin.Talents.DivinePrecision) * core.SpellHitRatingPerHitChance

	if paladin.Talents.Toughness > 0 {
		paladin.ApplyEquipScaling(stats.Armor, 1.0+0.02*float64(paladin.Talents.Toughness))
	}

	// These are no-op if untalented.
	paladin.MultiplyStat(stats.Strength, 1.0+0.02*float64(paladin.Talents.DivineStrength))
	paladin.MultiplyStat(stats.Intellect, 1.0+0.02*float64(paladin.Talents.DivineIntellect))
	paladin.AddStat(stats.Defense, 4*float64(paladin.Talents.Anticipation))
	paladin.AddStat(stats.Parry, 1*float64(paladin.Talents.Deflection))
	// Holy Power gives every spell 1% crit per point here; Holy Shock's larger share is the
	// extra 2% per point added on the spell itself, see holy_shock.go.
	// The beta client's curves are 1% a rank for every spell and 2% a rank more on Holy Shock.
	paladin.AddStat(stats.SpellCrit, float64(paladin.Talents.HolyPower)*core.SpellCritRatingPerCritChance)
	paladin.PseudoStats.SpiritRegenRateCasting += 0.1 * float64(paladin.Talents.Reverence)

	// Sacred Duty: 2% a rank, confirmed by the beta client.
	paladin.MultiplyStat(stats.Stamina, 1.0+0.02*float64(paladin.Talents.SacredDuty))

	// Shield Specialization: 10% a rank, confirmed by the beta client.
	// NOTE: Total SBV will be inflated until
	// https://github.com/wowsims/sod/issues/1025 gets resolved.
	paladin.PseudoStats.BlockValueMultiplier += 0.1 * float64(paladin.Talents.ShieldSpecialization)

	// Champion of the Light: 33/66/100%, confirmed by the beta client.
	if paladin.Talents.ChampionOfTheLight > 0 {
		paladin.AddStatDependency(stats.Intellect, stats.SpellPower, []float64{0, 0.33, 0.66, 1.00}[paladin.Talents.ChampionOfTheLight])
	}

	paladin.applyWeaponSpecialization()
	paladin.applyCrusade()
	paladin.applyVengeance()
	paladin.applyVindication()
	paladin.applyRedoubt()
	paladin.applyReckoning()
	paladin.applyShieldSpecialization()
	paladin.applyConsecratedGround()
	paladin.applyInstrumentOfLaw()
	paladin.applySanctifiedJudgement()
}

// Improved Seals raises the damage of every seal and of the judgement it powers.
func (paladin *Paladin) improvedSeals() float64 {
	return 1 + 0.05*float64(paladin.Talents.ImprovedSeals)
}

func (paladin *Paladin) benediction() int32 {
	return 100 - 2*paladin.Talents.Benediction
}

// Holy Conduit only discounts Consecration, Holy Wrath, Exorcism and Hammer of Wrath.
func (paladin *Paladin) holyConduit() int32 {
	return 100 - 20*paladin.Talents.HolyConduit
}

// Purifying Power shortens the Exorcism and Holy Wrath cooldowns.
func (paladin *Paladin) purifyingPower(duration time.Duration) time.Duration {
	return time.Duration(float64(duration) * (1 - []float64{0, .17, .33}[paladin.Talents.PurifyingPower]))
}

func (paladin *Paladin) applyRedoubt() {
	if paladin.Talents.Redoubt == 0 {
		return
	}

	// Beta client 1.60.1.69893, talent curves: 6% block a rank and a 2% chance a rank, 10 sec or 5
	// blocks. The trees had the block flat at 6% and the chance at 10% a rank. The chance is the
	// curve on effect index 1, which the talent spell has no effect for; its rank 5 value is the
	// 10% ProcChance the spell carries, the way a spell's own number is always the top rank's.
	blockBonus := 6.0 * float64(paladin.Talents.Redoubt) * core.BlockRatingPerBlockChance

	paladin.redoubtAura = paladin.RegisterAura(core.Aura{
		Label:     "Redoubt",
		ActionID:  core.ActionID{SpellID: 20134},
		Duration:  time.Second * 10,
		MaxStacks: 5,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			paladin.AddStatDynamic(sim, stats.Block, blockBonus)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			paladin.AddStatDynamic(sim, stats.Block, -blockBonus)
		},
		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.DidBlock() {
				aura.RemoveStack(sim)
			}
		},
	})

	// Forever moved the trigger from taking a crit to any melee attack that lands.
	core.MakeProcTriggerAura(&paladin.Unit, core.ProcTrigger{
		Name:       "Redoubt Trigger",
		Callback:   core.CallbackOnSpellHitTaken,
		Outcome:    core.OutcomeLanded,
		ProcMask:   core.ProcMaskMelee,
		ProcChance: 0.02 * float64(paladin.Talents.Redoubt),
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			paladin.redoubtAura.Activate(sim)
			paladin.redoubtAura.SetStacks(sim, 5)
		},
	})
}

func (paladin *Paladin) applyReckoning() {

	if paladin.Talents.Reckoning == 0 {
		return
	}

	procID := core.ActionID{SpellID: 20178} // Reckoning Proc ID

	core.MakeProcTriggerAura(&paladin.Unit, core.ProcTrigger{
		Name:       "Reckoning Crit Trigger",
		Callback:   core.CallbackOnSpellHitTaken,
		Outcome:    core.OutcomeCrit,
		ProcMask:   core.ProcMaskMeleeOrRanged,
		ProcChance: 0.2 * float64(paladin.Talents.Reckoning),
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			paladin.AutoAttacks.ExtraMHAttack(sim, 1, procID, spell.ActionID)
		},
	})

	// Forever also gives Reckoning a smaller chance to fire off a block.
	core.MakeProcTriggerAura(&paladin.Unit, core.ProcTrigger{
		Name:       "Reckoning Block Trigger",
		Callback:   core.CallbackOnSpellHitTaken,
		Outcome:    core.OutcomeBlock,
		ProcMask:   core.ProcMaskMelee,
		ProcChance: 0.08 * float64(paladin.Talents.Reckoning),
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			paladin.AutoAttacks.ExtraMHAttack(sim, 1, procID, spell.ActionID)
		},
	})
}

// Shield Specialization returns mana on block, on top of the absorb applied in ApplyTalents.
func (paladin *Paladin) applyShieldSpecialization() {
	if paladin.Talents.ShieldSpecialization == 0 {
		return
	}

	actionID := core.ActionID{SpellID: 20148}
	manaMetrics := paladin.NewManaMetrics(actionID)

	icd := core.Cooldown{
		Timer:    paladin.NewTimer(),
		Duration: time.Second * 3,
	}

	// 33/66/100%, confirmed by the beta client. The 6% of maximum mana does not scale.
	core.MakeProcTriggerAura(&paladin.Unit, core.ProcTrigger{
		Name:       "Shield Specialization Trigger",
		Callback:   core.CallbackOnSpellHitTaken,
		Outcome:    core.OutcomeBlock,
		ProcMask:   core.ProcMaskMelee,
		ProcChance: []float64{0, 0.33, 0.66, 1}[paladin.Talents.ShieldSpecialization],
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !icd.IsReady(sim) {
				return
			}
			icd.Use(sim)
			paladin.AddMana(sim, 0.06*paladin.MaxMana(), manaMetrics)
		},
	})
}

func (paladin *Paladin) getWeaponSpecializationModifier() float64 {
	handType := paladin.MainHand().HandType
	if handType == proto.HandType_HandTypeMainHand || handType == proto.HandType_HandTypeOneHand {
		return 1. + []float64{0, .03, .07, .10}[paladin.Talents.OneHandedWeaponSpecialization]
	} else if handType == proto.HandType_HandTypeTwoHand {
		return 1. + 0.02*float64(paladin.Talents.TwoHandedWeaponSpecialization)
	} else {
		return 1.
	}
}

// Affects all physical damage or spells that can be rolled as physical.
func (paladin *Paladin) applyWeaponSpecialization() {
	paladin.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexPhysical] *= paladin.getWeaponSpecializationModifier()
}

func (paladin *Paladin) applyCrusade() {
	if paladin.Talents.Crusade == 0 {
		return
	}

	multiplier := 1 + 0.01*float64(paladin.Talents.Crusade)
	paladin.PseudoStats.DamageDealtMultiplier *= multiplier

	// The same bonus again, but only against Demon and Undead targets. It is damage dealt and
	// nothing else, so the crit multiplier is left alone.
	paladin.Env.RegisterPostFinalizeEffect(func() {
		for _, target := range paladin.Env.Encounter.Targets {
			if target.MobType != proto.MobType_MobTypeDemon && target.MobType != proto.MobType_MobTypeUndead {
				continue
			}
			for _, at := range paladin.AttackTables[target.UnitIndex] {
				at.DamageDealtMultiplier *= multiplier
			}
		}
	})
}

func (paladin *Paladin) applyVengeance() {
	if paladin.Talents.Vengeance == 0 {
		return
	}

	// Beta client 1.60.1.70009: 1% a stack per rank, up to 3 stacks, 30 sec (the trees had copied
	// rank 1 into every rank).
	perStack := 0.01 * float64(paladin.Talents.Vengeance)
	procAura := paladin.RegisterAura(core.Aura{
		Label:     "Vengeance Proc",
		ActionID:  core.ActionID{SpellID: 20059},
		Duration:  time.Second * 30,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			multiplier := (1 + perStack*float64(newStacks)) / (1 + perStack*float64(oldStacks))
			aura.Unit.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly] *= multiplier
			aura.Unit.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexPhysical] *= multiplier
		},
	})

	paladin.RegisterAura(core.Aura{
		Label:    "Vengeance",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			// Periodic critical strikes use OnPeriodicDamageDealt instead.
			if result.Target.Type == core.EnemyUnit && result.DidCrit() {
				procAura.Activate(sim)
				procAura.AddStack(sim)
			}
		},
	})
}

func (paladin *Paladin) applyVindication() {
	if paladin.Talents.Vindication == 0 {
		return
	}

	// Beta client 1.60.1.69893: 1% attack power a rank for 30 sec, on every damaging melee attack that
	// lands (100% proc chance). The attack power the target loses is not modelled, nothing in the sim
	// reads an enemy's attack power.
	attackPowerMultiplier := paladin.NewDynamicMultiplyStat(stats.AttackPower, 1+0.01*float64(paladin.Talents.Vindication))

	vindicationAura := paladin.RegisterAura(core.Aura{
		Label:    "Vindication Proc",
		ActionID: core.ActionID{SpellID: 26021},
		Duration: time.Second * 30,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			paladin.EnableDynamicStatDep(sim, attackPowerMultiplier)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			paladin.DisableDynamicStatDep(sim, attackPowerMultiplier)
		},
	})

	paladin.RegisterAura(core.Aura{
		Label:    "Vindication Talent",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) {
				vindicationAura.Activate(sim)
			}
		},
	})
}

// Consecrated Ground buffs Holy damage while the paladin's Consecration is on the ground.
func (paladin *Paladin) applyConsecratedGround() {
	if paladin.Talents.ConsecratedGround == 0 {
		return
	}

	// The tooltip caps the bonus at the first 4 enemies to enter the Consecration (Consecration's
	// $s3 in the beta client), which is not modelled: the buff sits on the paladin, so every
	// target takes it. It only differs from the game on a pull of more than 4.
	multiplier := 1 + 0.05*float64(paladin.Talents.ConsecratedGround)

	buffAura := paladin.RegisterAura(core.Aura{
		Label:    "Consecrated Ground",
		ActionID: core.ActionID{SpellID: 26573},
		Duration: time.Second * 8,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			paladin.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly] *= multiplier
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			paladin.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly] /= multiplier
		},
	})

	paladin.RegisterAura(core.Aura{
		Label:    "Consecrated Ground Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell.SpellCode == SpellCode_PaladinConsecration {
				buffAura.Activate(sim)
			}
		},
	})
}

// Instrument of Law also shaves the Hammer of Wrath cast time, see hammer_of_wrath.go.
func (paladin *Paladin) applyInstrumentOfLaw() {
	if paladin.Talents.InstrumentOfLaw == 0 || paladin.Options.RighteousFury {
		return
	}

	// 10% a rank, confirmed by the beta client.
	paladin.PseudoStats.ThreatMultiplier *= 1 - 0.1*float64(paladin.Talents.InstrumentOfLaw)
}

// Sanctified Judgement refunds part of the mana spent on the seal that Judgement consumes.
func (paladin *Paladin) applySanctifiedJudgement() {
	if paladin.Talents.SanctifiedJudgement == 0 {
		return
	}

	manaMetrics := paladin.NewManaMetrics(core.ActionID{SpellID: 31876})

	procChance := []float64{0, 0.33, 0.66, 1.00}[paladin.Talents.SanctifiedJudgement]
	refund := 0.2 * float64(paladin.Talents.SanctifiedJudgement)

	paladin.RegisterAura(core.Aura{
		Label:    "Sanctified Judgement",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell != paladin.judgement || paladin.currentSealSpell == nil {
				return
			}

			if sim.Proc(procChance, "Sanctified Judgement") {
				paladin.AddMana(sim, refund*paladin.currentSealSpell.Cost.GetCurrentCost(), manaMetrics)
			}
		},
	})
}
