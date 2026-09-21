package rogue

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func (rogue *Rogue) ApplyTalents() {
	rogue.applyRuthlessness()
	rogue.applyMurder()
	rogue.applyRelentlessStrikes()
	rogue.applySealFate()
	rogue.applyHackAndSlash()
	rogue.applyWeaponExpertise()
	rogue.applyInitiative()
	rogue.applySerratedBlades()
	rogue.applyCutthroat()
	rogue.applyThousandCuts()

	rogue.AddStat(stats.Dodge, 1*float64(rogue.Talents.LightningReflexes))
	rogue.AddStat(stats.Parry, 2*float64(rogue.Talents.Deflection))
	rogue.AddStat(stats.MeleeCrit, 1*float64(rogue.Talents.Malice))
	rogue.AddStat(stats.MeleeHit, 1*float64(rogue.Talents.Precision))
	// Malice and Precision cover Poisons in Forever, and those roll against the spell stats.
	rogue.AddStat(stats.SpellCrit, 1*float64(rogue.Talents.Malice))
	rogue.AddStat(stats.SpellHit, 1*float64(rogue.Talents.Precision))
	rogue.AutoAttacks.OHConfig().DamageMultiplier *= rogue.dwsMultiplier()

	rogue.registerColdBloodCD()
	rogue.registerBladeFlurryCD()
	rogue.registerAdrenalineRushCD()
	rogue.registerPreparationCD()
	rogue.registerPremeditation()
	rogue.registerGhostlyStrikeSpell()
	rogue.applyRiposte()
}

// dwsMultiplier returns the offhand damage multiplier
func (rogue *Rogue) dwsMultiplier() float64 {
	return 1 + 0.05*float64(rogue.Talents.DualWieldSpecialization)
}

func (rogue *Rogue) applyRuthlessness() {
	if rogue.Talents.Ruthlessness == 0 {
		return
	}

	procChance := 0.2 * float64(rogue.Talents.Ruthlessness)
	cpMetrics := rogue.NewComboPointMetrics(core.ActionID{SpellID: 14161})
	rogue.OnComboPointsSpent(func(sim *core.Simulation, spell *core.Spell, comboPoints int32) {
		if sim.Proc(procChance, "Ruthlessness") {
			rogue.AddComboPointsIgnoreTarget(sim, 1, cpMetrics)
		}
	})
}

// Murder talent
func (rogue *Rogue) applyMurder() {
	if rogue.Talents.Murder == 0 {
		return
	}

	// post finalize, since attack tables need to be setup
	rogue.Env.RegisterPostFinalizeEffect(func() {
		for _, t := range rogue.Env.Encounter.Targets {
			switch t.MobType {
			case proto.MobType_MobTypeHumanoid, proto.MobType_MobTypeGiant:
				multiplier := []float64{1, 1.02, 1.04}[rogue.Talents.Murder]
				for _, at := range rogue.AttackTables[t.UnitIndex] {
					at.DamageDealtMultiplier *= multiplier
					at.CritMultiplier *= multiplier
				}
			}
		}
	})
}

func (rogue *Rogue) applyRelentlessStrikes() {
	if !rogue.Talents.RelentlessStrikes {
		return
	}

	cpMetrics := rogue.NewEnergyMetrics(core.ActionID{SpellID: 14179})
	rogue.OnComboPointsSpent(func(sim *core.Simulation, spell *core.Spell, comboPoints int32) {
		if sim.Proc(0.2*float64(comboPoints), "RelentlessStrikes") {
			rogue.AddEnergy(sim, 25, cpMetrics)
		}
	})
}

// Cold Blood talent
func (rogue *Rogue) registerColdBloodCD() {
	if !rogue.Talents.ColdBlood {
		return
	}

	actionID := core.ActionID{SpellID: 14177}

	coldBloodAura := rogue.RegisterAura(core.Aura{
		Label:    "Cold Blood",
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range rogue.Spellbook {
				if spell.Flags.Matches(SpellFlagColdBlooded) {
					spell.BonusCritRating += 100 * core.CritRatingPerCritChance
				}
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range rogue.Spellbook {
				if spell.Flags.Matches(SpellFlagColdBlooded) {
					spell.BonusCritRating -= 100 * core.CritRatingPerCritChance
				}
			}
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.Flags.Matches(SpellFlagColdBlooded) {
				aura.Deactivate(sim)
			}
		},
	})

	rogue.ColdBlood = rogue.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    rogue.NewTimer(),
				Duration: time.Minute * 3,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			coldBloodAura.Activate(sim)
		},
	})

	rogue.AddMajorCooldown(core.MajorCooldown{
		Spell: rogue.ColdBlood,
		Type:  core.CooldownTypeDPS,
	})
}

// Seal Fate talent
func (rogue *Rogue) applySealFate() {
	if rogue.Talents.SealFate == 0 {
		return
	}

	procChance := 0.2 * float64(rogue.Talents.SealFate)
	cpMetrics := rogue.NewComboPointMetrics(core.ActionID{SpellID: 14195})

	icd := core.Cooldown{
		Timer:    rogue.NewTimer(),
		Duration: 500 * time.Millisecond,
	}

	rogue.RegisterAura(core.Aura{
		Label:    "Seal Fate",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !spell.Flags.Matches(SpellFlagBuilder) {
				return
			}

			if !result.Outcome.Matches(core.OutcomeCrit) {
				return
			}

			if icd.IsReady(sim) && sim.Proc(procChance, "Seal Fate") {
				rogue.AddComboPoints(sim, 1, result.Target, cpMetrics)
				icd.Use(sim)
			}
		},
	})
}

// Initiative talent
func (rogue *Rogue) applyInitiative() {
	if rogue.Talents.Initiative == 0 {
		return
	}

	// The beta rounds rank 2 up to 67% rather than doubling rank 1's 33%.
	procChance := []float64{0, 0.33, 0.67, 1.00}[rogue.Talents.Initiative]
	cpMetrics := rogue.NewComboPointMetrics(core.ActionID{SpellID: 13980})

	rogue.RegisterAura(core.Aura{
		Label:    "Initiative",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell == rogue.Garrote || spell == rogue.Ambush {
				if result.Landed() {
					if sim.Proc(procChance, "Initiative") {
						rogue.AddComboPoints(sim, 1, result.Target, cpMetrics)
					}
				}
			}
		},
	})
}

// Weapon Expertise no longer grants weapon skill, it takes the chance for the rogue's
// attacks to be dodged or parried off the attack table directly.
func (rogue *Rogue) applyWeaponExpertise() {
	if rogue.Talents.WeaponExpertise == 0 {
		return
	}

	rogue.AddStat(stats.Expertise, float64(rogue.Talents.WeaponExpertise)*core.ExpertiseRatingPerExpertiseChance)
}

// Serrated Blades ignores a share of the target's Armor rather than a flat amount, 3% per
// rank in the beta client. The Rupture bonus lives on the spell itself.
func (rogue *Rogue) applySerratedBlades() {
	if rogue.Talents.SerratedBlades == 0 {
		return
	}

	rogue.PseudoStats.ArmorIgnorePercent += 0.03 * float64(rogue.Talents.SerratedBlades)
}

// Cutthroat lets Ambush be used outside of Stealth for a short while after a Backstab.
func (rogue *Rogue) applyCutthroat() {
	if rogue.Talents.Cutthroat == 0 {
		return
	}

	// 3% per rank, and a 10 sec window at every rank (beta client).
	procChance := 0.03 * float64(rogue.Talents.Cutthroat)

	rogue.CutthroatAura = rogue.RegisterAura(core.Aura{
		Label:    "Cutthroat",
		Duration: time.Second * 10,
	})

	rogue.RegisterAura(core.Aura{
		Label:    "Cutthroat Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.SpellCode != SpellCode_RogueBackstab || !result.Landed() {
				return
			}

			if sim.Proc(procChance, "Cutthroat") {
				rogue.CutthroatAura.Activate(sim)
			}
		},
	})
}

// Thousand Cuts discounts the next Hemorrhage or Backstab as Rupture ticks.
func (rogue *Rogue) applyThousandCuts() {
	if !rogue.Talents.ThousandCuts {
		return
	}

	var affectedSpells []*core.Spell
	rogue.ThousandCutsAura = rogue.RegisterAura(core.Aura{
		Label:     "Thousand Cuts",
		Duration:  time.Second * 10,
		MaxStacks: 5,
		OnInit: func(aura *core.Aura, sim *core.Simulation) {
			affectedSpells = core.FilterSlice([]*core.Spell{rogue.Backstab, rogue.Hemorrhage}, func(spell *core.Spell) bool {
				return spell != nil
			})
		},
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			for _, spell := range affectedSpells {
				spell.Cost.FlatModifier -= 3 * (newStacks - oldStacks)
			}
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.SpellCode == SpellCode_RogueBackstab || spell.SpellCode == SpellCode_RogueHemorrhage {
				aura.Deactivate(sim)
			}
		},
	})

	rogue.RegisterAura(core.Aura{
		Label:    "Thousand Cuts Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.SpellCode != SpellCode_RogueRupture {
				return
			}

			rogue.ThousandCutsAura.Activate(sim)
			rogue.ThousandCutsAura.AddStack(sim)
		},
	})
}

// Quietus is an execute range bonus, so it's rolled in as the strike is cast.
func (rogue *Rogue) quietusMultiplier(sim *core.Simulation) float64 {
	if rogue.Talents.Quietus == 0 || !sim.IsExecutePhase35() {
		return 1
	}

	// 2% per rank, below 35% health at every rank (beta client).
	return 1 + 0.02*float64(rogue.Talents.Quietus)
}

func (rogue *Rogue) registerBladeFlurryCD() {
	if !rogue.Talents.BladeFlurry {
		return
	}

	// TODO verify that this double dips from damage modifiers

	var curDmg float64
	bfHit := rogue.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 22482},
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskEmpty, // No proc mask, so it won't proc itself.
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagNoOnCastComplete,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealDamage(sim, target, curDmg, spell.OutcomeAlwaysHit)
		},
	})

	rogue.BladeFlurryAura = rogue.RegisterAura(core.Aura{
		Label:    "Blade Flurry",
		ActionID: core.ActionID{SpellID: 13877},
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, 1.2)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, 1/1.2)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if sim.GetNumTargets() < 2 {
				return
			}

			if result.Damage == 0 || !spell.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}

			// Undo armor reduction to get the raw damage value.
			curDmg = result.Damage / result.ResistanceMultiplier

			bfHit.Cast(sim, rogue.Env.NextTargetUnit(result.Target))
			bfHit.SpellMetrics[result.Target.UnitIndex].Casts--
		},
	})

	cooldownDur := time.Minute * 2
	rogue.BladeFlurry = rogue.RegisterSpell(core.SpellConfig{
		SpellCode: SpellCode_RogueBladeFlurry,
		ActionID:  core.ActionID{SpellID: 13877},
		Flags:     core.SpellFlagAPL,

		EnergyCost: core.EnergyCostOptions{
			Cost: 25,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    rogue.NewTimer(),
				Duration: cooldownDur,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			rogue.BreakStealth(sim)
			rogue.BladeFlurryAura.Activate(sim)
		},
	})

	rogue.AddMajorCooldown(core.MajorCooldown{
		Spell:    rogue.BladeFlurry,
		Type:     core.CooldownTypeDPS,
		Priority: core.CooldownPriorityDefault,
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			if sim.GetRemainingDuration() > cooldownDur+time.Second*15 {
				// We'll have enough time to cast another BF, so use it immediately to make sure we get the 2nd one.
				return true
			}

			// Since this is our last BF, wait until we have SND / procs up.
			sndTimeRemaining := rogue.SliceAndDiceAura.RemainingDuration(sim)
			return sndTimeRemaining >= time.Second
		},
	})
}

var AdrenalineRushActionID = core.ActionID{SpellID: 13750}

func (rogue *Rogue) registerAdrenalineRushCD() {
	if !rogue.Talents.AdrenalineRush {
		return
	}

	rogue.AdrenalineRushAura = rogue.RegisterAura(core.Aura{
		Label:    "Adrenaline Rush",
		ActionID: AdrenalineRushActionID,
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			rogue.ApplyEnergyTickMultiplier(sim, 1.0)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			rogue.ApplyEnergyTickMultiplier(sim, -1.0)
		},
	})

	rogue.AdrenalineRush = rogue.RegisterSpell(core.SpellConfig{
		SpellCode: SpellCode_RogueAdrenalineRush,
		ActionID:  AdrenalineRushActionID,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    rogue.NewTimer(),
				Duration: time.Minute * 5,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			rogue.BreakStealth(sim)
			rogue.AdrenalineRushAura.Activate(sim)
		},
	})

	rogue.AddMajorCooldown(core.MajorCooldown{
		Spell:    rogue.AdrenalineRush,
		Type:     core.CooldownTypeDPS,
		Priority: core.CooldownPriorityBloodlust,
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return rogue.CurrentEnergy() <= 45.0
		},
	})
}

func (rogue *Rogue) lethality() float64 {
	return 0.04 * float64(rogue.Talents.Lethality)
}
