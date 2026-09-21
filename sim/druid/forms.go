package druid

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

type DruidForm uint8

const (
	Humanoid DruidForm = 1 << iota
	Bear
	Cat
	Moonkin
	Any = Humanoid | Bear | Cat | Moonkin
)

func (form DruidForm) Matches(other DruidForm) bool {
	return (form & other) != 0
}

func (druid *Druid) GetForm() DruidForm {
	return druid.form
}

func (druid *Druid) InForm(form DruidForm) bool {
	return druid.form.Matches(form)
}

func (druid *Druid) GetCatWeapon() core.Weapon {
	return core.Weapon{
		BaseDamageMin:        43.84,
		BaseDamageMax:        65.76,
		SwingSpeed:           1.0,
		NormalizedSwingSpeed: 1.0,
		AttackPowerPerDPS:    core.DefaultAttackPowerPerDPS,
	}
}

func (druid *Druid) GetBearWeapon() core.Weapon {
	return core.Weapon{
		BaseDamageMin:        109,
		BaseDamageMax:        165,
		SwingSpeed:           2.5,
		NormalizedSwingSpeed: 2.5,
		AttackPowerPerDPS:    core.DefaultAttackPowerPerDPS,
	}
}

// TODO: Class bonus stats for both cat and bear.
func (druid *Druid) GetFormShiftStats() stats.Stats {
	s := stats.Stats{
		stats.AttackPower: float64(druid.Talents.PredatoryStrikes) * 0.5 * float64(druid.Level),
		stats.MeleeCrit:   float64(druid.Talents.SharpenedClaws) * 3 * core.CritRatingPerCritChance,
	}
	/*
		if weapon := druid.GetMHWeapon(); weapon != nil {
			dps := (weapon.WeaponDamageMax+weapon.WeaponDamageMin)/2.0/weapon.SwingSpeed + druid.PseudoStats.BonusMHDps
			weapAp := weapon.Stats[stats.AttackPower] + weapon.Enchant.Stats[stats.AttackPower]
			fap := math.Floor((dps - 54.8) * 14)

			s[stats.AttackPower] += fap
			s[stats.AttackPower] += (fap + weapAp) * ((0.2 / 3) * float64(druid.Talents.PredatoryStrikes))
		}
	*/

	return s
}

func (druid *Druid) GetDynamicPredStrikeStats() stats.Stats {
	// Accounts for ap bonus for 'dynamic' enchants
	// just scourgebane currently, this is a bit hacky but is needed as the bonus varies based on current target
	// so has to be 'cached' differently
	s := stats.Stats{}
	if weapon := druid.GetMHWeapon(); weapon != nil {
		bonusAp := 0.0
		if weapon.Enchant.EffectID == 3247 && druid.CurrentTarget.MobType == proto.MobType_MobTypeUndead {
			bonusAp += 140
		}
		s[stats.AttackPower] += bonusAp * ((0.2 / 3) * float64(druid.Talents.PredatoryStrikes))
	}
	return s
}

func (druid *Druid) registerCatFormSpell() {
	actionID := core.ActionID{SpellID: 768}

	statBonus := druid.GetFormShiftStats().Add(stats.Stats{
		stats.AttackPower: float64(druid.Level) * 2,
	})

	agiApDep := druid.NewDynamicStatDependency(stats.Agility, stats.AttackPower, 1)
	feralApDep := druid.NewDynamicStatDependency(stats.FeralAttackPower, stats.AttackPower, 1)

	var hotwDep *stats.StatDependency
	if druid.Talents.HeartOfTheWild > 0 {
		hotwDep = druid.NewDynamicMultiplyStat(stats.Strength, 1.0+0.02*float64(druid.Talents.HeartOfTheWild))
	}

	clawWeapon := druid.GetCatWeapon()

	predBonus := stats.Stats{}

	druid.CatFormAura = druid.RegisterAura(core.Aura{
		Label:      "Cat Form",
		ActionID:   actionID,
		Duration:   core.NeverExpires,
		BuildPhase: core.Ternary(druid.StartingForm.Matches(Cat), core.CharacterBuildPhaseBase, core.CharacterBuildPhaseNone),
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			if !druid.Env.MeasuringStats && druid.form != Humanoid {
				druid.CancelShapeshift(sim)
			}
			druid.form = Cat
			druid.SetCurrentPowerBar(core.EnergyBar)

			druid.AutoAttacks.SetMH(clawWeapon)

			druid.PseudoStats.ThreatMultiplier *= 0.71
			druid.SetShapeshift(aura)

			predBonus = druid.GetDynamicPredStrikeStats()
			druid.AddStatsDynamic(sim, predBonus)
			druid.AddStatsDynamic(sim, statBonus)
			druid.EnableDynamicStatDep(sim, agiApDep)
			druid.EnableDynamicStatDep(sim, feralApDep)
			if hotwDep != nil {
				druid.EnableDynamicStatDep(sim, hotwDep)
			}

			if !druid.Env.MeasuringStats {
				druid.AutoAttacks.SetReplaceMHSwing(nil)
				druid.AutoAttacks.EnableAutoSwing(sim)
				druid.manageCooldownsEnabled()
				druid.UpdateManaRegenRates()
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			druid.form = Humanoid
			druid.SetCurrentPowerBar(core.ManaBar)

			druid.TigersFuryAura.Deactivate(sim)

			druid.AutoAttacks.SetMH(druid.WeaponFromMainHand())

			druid.lastCatFormEnergy = druid.CurrentEnergy()
			druid.lastCatFormExitAt = sim.CurrentTime

			druid.PseudoStats.ThreatMultiplier /= 0.71
			druid.SetShapeshift(nil)

			druid.AddStatsDynamic(sim, predBonus.Invert())
			druid.AddStatsDynamic(sim, statBonus.Invert())
			druid.DisableDynamicStatDep(sim, agiApDep)
			druid.DisableDynamicStatDep(sim, feralApDep)
			if hotwDep != nil {
				druid.DisableDynamicStatDep(sim, hotwDep)
			}

			if !druid.Env.MeasuringStats {
				druid.AutoAttacks.SetReplaceMHSwing(nil)
				druid.AutoAttacks.EnableAutoSwing(sim)
				druid.manageCooldownsEnabled()
				druid.UpdateManaRegenRates()

				//druid.TigersFuryAura.Deactivate(sim)
			}
		},
	})

	energyMetrics := druid.NewEnergyMetrics(actionID)

	hasWolfheadBonus := false
	if head := druid.Equipment.Head(); !druid.Env.IsForever() && head != nil && head.ID == WolfsheadHelm {
		hasWolfheadBonus = true
	}

	druid.CatForm = druid.RegisterSpell(Any, core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.55,
			Multiplier: 100 - 10*druid.Talents.NaturalShapeshifter,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				if druid.CatFormAura.IsActive() {
					cast.GCD = 0
					spell.Cost.Multiplier -= 100
				}
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			if druid.CatFormAura.IsActive() {
				druid.CancelShapeshift(sim)
				spell.Cost.Multiplier += 100
			} else {
				maxShiftEnergy := druid.furorShiftEnergy(sim)
				maxShiftEnergy = core.TernaryFloat64(hasWolfheadBonus, maxShiftEnergy+20, maxShiftEnergy)
				energyDelta := maxShiftEnergy - druid.CurrentEnergy()

				if energyDelta > 0 {
					druid.AddEnergy(sim, energyDelta, energyMetrics)
				} else {
					druid.SpendEnergy(sim, -energyDelta, energyMetrics)
				}

				druid.CatFormAura.Activate(sim)
			}
		},
	})
}

// Forever reworks powershifting: instead of a chance at a flat 40 energy, shifting
// into Cat Form carries over a share of the energy you left the form with, plus a
// small amount for every second spent out of form.
// The beta client's curve is 20-100 and the text builds all three from it: that share of the energy, a tenth of it a
// second, and the whole of it as the cap, so 20%, 2 a second and 20 Energy per point.
func (druid *Druid) furorShiftEnergy(sim *core.Simulation) float64 {
	if druid.Talents.Furor == 0 {
		return 0
	}

	points := float64(druid.Talents.Furor)
	carryOver := druid.lastCatFormEnergy * 0.2 * points
	outOfForm := 0.0
	if druid.lastCatFormExitAt > 0 {
		outOfForm = min(20*points, 2*points*(sim.CurrentTime-druid.lastCatFormExitAt).Seconds())
	}

	return min(druid.MaxEnergy(), carryOver+outOfForm)
}

// Dire Bear Form: 180 attack power, 1240 health, 360% more armor from items and 30%
// more threat. Rage replaces mana while it lasts.
const BearFormArmorMultiplier = 4.6
const BearFormThreatMultiplier = 1.3

func (druid *Druid) registerBearFormSpell() {
	actionID := core.ActionID{SpellID: 9634}
	healthMetrics := druid.NewHealthMetrics(actionID)

	statBonus := druid.GetFormShiftStats().Add(stats.Stats{
		stats.AttackPower: 3 * float64(druid.Level),
		stats.Health:      1240,
	})

	feralApDep := druid.NewDynamicStatDependency(stats.FeralAttackPower, stats.AttackPower, 1)

	var hotwDep *stats.StatDependency
	if druid.Talents.HeartOfTheWild > 0 {
		hotwDep = druid.NewDynamicMultiplyStat(stats.Stamina, 1.0+0.04*float64(druid.Talents.HeartOfTheWild))
	}

	clawWeapon := druid.GetBearWeapon()
	predBonus := stats.Stats{}

	druid.BearFormAura = druid.RegisterAura(core.Aura{
		Label:      "Bear Form",
		ActionID:   actionID,
		Duration:   core.NeverExpires,
		BuildPhase: core.Ternary(druid.StartingForm.Matches(Bear), core.CharacterBuildPhaseBase, core.CharacterBuildPhaseNone),
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			if !druid.Env.MeasuringStats && druid.form != Humanoid {
				druid.CancelShapeshift(sim)
			}
			druid.form = Bear
			druid.SetCurrentPowerBar(core.RageBar)

			druid.AutoAttacks.SetMH(clawWeapon)

			druid.PseudoStats.ThreatMultiplier *= BearFormThreatMultiplier
			druid.SetShapeshift(aura)

			predBonus = druid.GetDynamicPredStrikeStats()
			druid.AddStatsDynamic(sim, predBonus)
			druid.AddStatsDynamic(sim, statBonus)
			druid.ApplyDynamicEquipScaling(sim, stats.Armor, BearFormArmorMultiplier)
			druid.EnableDynamicStatDep(sim, feralApDep)

			// Preserve fraction of max health when shifting. The shift at the start of
			// every iteration is left out of the metrics, which are compared across threads.
			healthFrac := druid.CurrentHealth() / druid.MaxHealth()
			if hotwDep != nil {
				druid.EnableDynamicStatDep(sim, hotwDep)
			}
			if sim.CurrentTime > 0 {
				druid.GainHealth(sim, healthFrac*druid.MaxHealth()-druid.CurrentHealth(), healthMetrics)
			}

			if !druid.Env.MeasuringStats {
				druid.AutoAttacks.SetReplaceMHSwing(druid.ReplaceBearMHFunc)
				druid.AutoAttacks.EnableAutoSwing(sim)
				druid.manageCooldownsEnabled()
				druid.UpdateManaRegenRates()
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			druid.form = Humanoid
			druid.SetCurrentPowerBar(core.ManaBar)

			druid.AutoAttacks.SetMH(druid.WeaponFromMainHand())

			druid.PseudoStats.ThreatMultiplier /= BearFormThreatMultiplier
			druid.SetShapeshift(nil)

			druid.AddStatsDynamic(sim, predBonus.Invert())
			druid.AddStatsDynamic(sim, statBonus.Invert())
			druid.RemoveDynamicEquipScaling(sim, stats.Armor, BearFormArmorMultiplier)
			druid.DisableDynamicStatDep(sim, feralApDep)

			healthFrac := druid.CurrentHealth() / druid.MaxHealth()
			if hotwDep != nil {
				druid.DisableDynamicStatDep(sim, hotwDep)
			}
			if sim.CurrentTime > 0 {
				druid.RemoveHealth(sim, druid.CurrentHealth()-healthFrac*druid.MaxHealth())
			}

			if !druid.Env.MeasuringStats {
				druid.AutoAttacks.SetReplaceMHSwing(nil)
				druid.AutoAttacks.EnableAutoSwing(sim)
				druid.manageCooldownsEnabled()
				druid.UpdateManaRegenRates()
				druid.EnrageAura.Deactivate(sim)
				druid.MaulQueueAura.Deactivate(sim)
			}
		},
	})

	rageMetrics := druid.NewRageMetrics(actionID)

	// The Bear half of Furor is the Classic one, a chance at 10 Rage on the shift.
	// The beta client keeps it: a 20-100% curve, and 17057 grants 10 Rage.
	furorProcChance := 0.2 * float64(druid.Talents.Furor)

	druid.BearForm = druid.RegisterSpell(Any, core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.55,
			Multiplier: 100 - 10*druid.Talents.NaturalShapeshifter,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				if druid.BearFormAura.IsActive() {
					cast.GCD = 0
					spell.Cost.Multiplier -= 100
				}
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			if druid.BearFormAura.IsActive() {
				druid.CancelShapeshift(sim)
				spell.Cost.Multiplier += 100
				return
			}

			rageDelta := 0 - druid.CurrentRage()
			if sim.Proc(furorProcChance, "Furor") {
				rageDelta += 10
			}
			if rageDelta > 0 {
				druid.AddRage(sim, rageDelta, rageMetrics)
			} else if rageDelta < 0 {
				druid.SpendRage(sim, -rageDelta, rageMetrics)
			}
			druid.BearFormAura.Activate(sim)
		},
	})
}

func (druid *Druid) manageCooldownsEnabled() {
	// Disable cooldowns not usable in form and/or delay others
	if druid.StartingForm.Matches(Cat | Bear) {
		for _, mcd := range druid.disabledMCDs {
			mcd.Enable()
		}
		druid.disabledMCDs = nil

		if druid.InForm(Humanoid) {
			// Disable cooldown that incurs a gcd, so we dont get stuck out of form when we dont need to (Greater Drums)
			for _, mcd := range druid.GetMajorCooldowns() {
				if mcd.Spell.DefaultCast.GCD > 0 {
					mcd.Disable()
					druid.disabledMCDs = append(druid.disabledMCDs, mcd)
				}
			}
		}
	}
}

// Moonkin Form: 360% more armor from items, and 3% critical strike chance for the party,
// which the raid buff carries. Periodic crits are the Forever rule for every class, not
// something the form grants.
const MoonkinFormArmorMultiplier = 4.6

func (druid *Druid) registerMoonkinFormSpell() {
	if !druid.Talents.MoonkinForm {
		return
	}

	actionID := core.ActionID{SpellID: 24858}

	druid.MoonkinFormAura = druid.RegisterAura(core.Aura{
		Label:      "Moonkin Form",
		ActionID:   actionID,
		Duration:   core.NeverExpires,
		BuildPhase: core.Ternary(druid.StartingForm.Matches(Moonkin), core.CharacterBuildPhaseBase, core.CharacterBuildPhaseNone),
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			if !druid.Env.MeasuringStats && druid.form != Humanoid {
				druid.CancelShapeshift(sim)
			}
			druid.form = Moonkin
			druid.ApplyDynamicEquipScaling(sim, stats.Armor, MoonkinFormArmorMultiplier)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			druid.form = Humanoid
			druid.RemoveDynamicEquipScaling(sim, stats.Armor, MoonkinFormArmorMultiplier)
		},
	})

	druid.MoonkinForm = druid.RegisterSpell(Any, core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost:   0.35,
			Multiplier: 100 - 10*druid.Talents.NaturalShapeshifter,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			druid.MoonkinFormAura.Activate(sim)
		},
	})
}
