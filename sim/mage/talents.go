package mage

import (
	"slices"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

func (mage *Mage) ApplyTalents() {
	mage.applyArcaneTalents()
	mage.applyFireTalents()
	mage.applyFrostTalents()
}

func (mage *Mage) applyArcaneTalents() {
	mage.applyArcaneConcentration()
	mage.applyMissileBarrage()
	mage.registerPresenceOfMindCD()
	mage.registerArcanePowerCD()

	// Arcane Subtlety, 8 and 15 spell penetration and 15% threat reduction per point in the beta client.
	if mage.Talents.ArcaneSubtlety > 0 {
		threatMultiplier := 1 - .15*float64(mage.Talents.ArcaneSubtlety)
		mage.AddStat(stats.SpellPenetration, []float64{0, 8, 15}[mage.Talents.ArcaneSubtlety])
		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolArcane) && spell.Flags.Matches(SpellFlagMage) {
				spell.ThreatMultiplier *= threatMultiplier
			}
		})
	}

	// Arcane Focus
	if mage.Talents.ArcaneFocus > 0 {
		bonusHit := 1 * float64(mage.Talents.ArcaneFocus) * core.SpellHitRatingPerHitChance
		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolArcane) && spell.Flags.Matches(SpellFlagMage) {
				spell.BonusHitRating += bonusHit
			}
		})
	}

	// Magic Absorption
	if mage.Talents.MagicAbsorption > 0 {
		magicAbsorptionBonus := 5 * float64(mage.Talents.MagicAbsorption)
		mage.AddResistances(magicAbsorptionBonus)
	}

	// Arcane Resilience
	if mage.Talents.ArcaneResilience > 0 {
		mage.AddStatDependency(stats.Intellect, stats.Armor, .25*float64(mage.Talents.ArcaneResilience))
	}

	// Arcane Impact
	if mage.Talents.ArcaneImpact > 0 {
		bonusCrit := 2 * float64(mage.Talents.ArcaneImpact) * core.SpellCritRatingPerCritChance
		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolArcane) && spell.Flags.Matches(SpellFlagMage) {
				spell.BonusCritRating += bonusCrit
			}
		})
	}

	// Arcane Meditation
	mage.PseudoStats.SpiritRegenRateCasting += []float64{0, .17, .33, .50}[mage.Talents.ArcaneMeditation]

	// Arcane Mind
	if mage.Talents.ArcaneMind > 0 {
		critBonus := .20 * float64(mage.Talents.ArcaneMind)

		mage.MultiplyStat(stats.Intellect, 1.0+0.02*float64(mage.Talents.ArcaneMind))
		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolArcane) && spell.Flags.Matches(SpellFlagMage) {
				spell.CritDamageBonus += critBonus
			}
		})
	}

	// Arcane Instability
	if mage.Talents.ArcaneInstability > 0 {
		bonusDamageMultiplierAdditive := .01 * float64(mage.Talents.ArcaneInstability)
		bonusCritRating := 1 * float64(mage.Talents.ArcaneInstability) * core.SpellCritRatingPerCritChance

		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.Flags.Matches(SpellFlagMage) {
				spell.DamageMultiplierAdditive += bonusDamageMultiplierAdditive
				spell.BonusCritRating += bonusCritRating
			}
		})
	}
}

func (mage *Mage) applyFireTalents() {
	mage.applyIgnite()
	mage.applyImprovedScorch()
	mage.applyMasterOfElements()
	mage.applyHotStreak()

	mage.registerCombustionCD()

	// Incineration
	if mage.Talents.Incineration > 0 {
		bonusCrit := 2 * float64(mage.Talents.Incineration) * core.SpellCritRatingPerCritChance
		affectedSpellCodes := []int32{SpellCode_MageArcaneBlast, SpellCode_MageFireBlast, SpellCode_MageIceLance, SpellCode_MageScorch}
		mage.OnSpellRegistered(func(spell *core.Spell) {
			if slices.Contains(affectedSpellCodes, spell.SpellCode) {
				spell.BonusCritRating += bonusCrit
			}
		})
	}

	// Burning Soul
	if mage.Talents.BurningSoul > 0 {
		threatMultiplier := 1 - .10*float64(mage.Talents.BurningSoul)
		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolFire) && spell.Flags.Matches(SpellFlagMage) {
				spell.ThreatMultiplier *= threatMultiplier
			}
		})
	}

	// Critical Mass
	if mage.Talents.CriticalMass > 0 {
		bonusCrit := 2 * float64(mage.Talents.CriticalMass) * core.SpellCritRatingPerCritChance
		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolFire) && spell.Flags.Matches(SpellFlagMage) {
				spell.BonusCritRating += bonusCrit
			}
		})
	}

	// Fire Power
	if mage.Talents.FirePower > 0 {
		bonusDamageMultiplierAdditive := 0.02 * float64(mage.Talents.FirePower)
		mage.OnSpellRegistered(func(spell *core.Spell) {
			// Fire Power buffs pretty much all mage fire spells EXCEPT ignite
			if spell.SpellSchool.Matches(core.SpellSchoolFire) && spell.Flags.Matches(SpellFlagMage) && spell.SpellCode != SpellCode_MageIgnite {
				spell.DamageMultiplierAdditive += bonusDamageMultiplierAdditive
			}
		})
	}
}

func (mage *Mage) applyFrostTalents() {
	mage.registerColdSnapCD()
	mage.registerIceBarrierSpell()
	mage.applyFingersOfFrost()
	mage.applyWintersChill()

	// Elemental Precision
	if mage.Talents.ElementalPrecision > 0 {
		bonusHit := 1 * float64(mage.Talents.ElementalPrecision) * core.SpellHitRatingPerHitChance

		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.Flags.Matches(SpellFlagMage) && (spell.SpellSchool.Matches(core.SpellSchoolFire) || spell.SpellSchool.Matches(core.SpellSchoolFrost)) {
				spell.BonusHitRating += bonusHit
			}
		})
	}

	// Ice Shards
	if mage.Talents.IceShards > 0 {
		critBonus := .20 * float64(mage.Talents.IceShards)

		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolFrost) && spell.Flags.Matches(SpellFlagMage) {
				spell.CritDamageBonus += critBonus
			}
		})
	}

	// Piercing Ice
	if mage.Talents.PiercingIce > 0 {
		bonusDamageMultiplierAdditive := 0.02 * float64(mage.Talents.PiercingIce)

		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolFrost) && spell.Flags.Matches(SpellFlagMage) {
				spell.DamageMultiplierAdditive += bonusDamageMultiplierAdditive
			}
		})
	}

	// Frost Channeling
	if mage.Talents.FrostChanneling > 0 {
		manaCostMultiplier := 5 * mage.Talents.FrostChanneling
		threatMultiplier := 1 - .10*float64(mage.Talents.FrostChanneling)
		mage.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellSchool.Matches(core.SpellSchoolFrost) && spell.Flags.Matches(SpellFlagMage) {
				spell.Cost.Multiplier -= manaCostMultiplier
				spell.ThreatMultiplier *= threatMultiplier
			}
		})
	}
}

func (mage *Mage) applyArcaneConcentration() {
	if mage.Talents.ArcaneConcentration == 0 {
		return
	}

	procChance := 0.02 * float64(mage.Talents.ArcaneConcentration)
	icd := core.Cooldown{Timer: mage.NewTimer(), Duration: time.Second}

	mage.ClearcastingAura = mage.RegisterAura(core.Aura{
		Label:    "Clearcasting",
		ActionID: core.ActionID{SpellID: 12536},
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.SchoolCostMultiplier.AddToMagicSchools(-100)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.SchoolCostMultiplier.AddToMagicSchools(100)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			// OnCastComplete is called after OnSpellHitDealt / etc, so don't deactivate if it was just activated.
			if aura.RemainingDuration(sim) == aura.Duration {
				return
			}
			if !spell.Flags.Matches(SpellFlagMage) {
				return
			}
			// Clearcasting itself makes the current cost zero. Eligibility is
			// based on the underlying paid damage spell, not its discounted cost.
			if spell.Cost == nil || spell.Cost.BaseCost == 0 || !spell.ProcMask.Matches(core.ProcMaskSpellDamage) {
				return
			}
			aura.Deactivate(sim)
		},
	})

	mage.RegisterAura(core.Aura{
		Label:    "Arcane Concentration",
		Duration: core.NeverExpires,
		Icd:      &icd,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() || !spell.Flags.Matches(SpellFlagMage) || spell.SpellCode == SpellCode_MageArcaneMissiles || !icd.IsReady(sim) {
				return
			}

			// TODO: Classic verify arcane missile proc chance
			// Arcane Missile ticks can proc CC, just at a low rate of about 1.5% with 5/5 Arcane Concentration
			// if spell == mage.ArcaneMissilesTickSpell {
			// 	procChance *= 0.15
			// }

			if sim.Proc(procChance, "Arcane Concentration") {
				icd.Use(sim)
				mage.ClearcastingAura.Activate(sim)
			}
		},
	})
}

// Arcane Blast feeds Missile Barrage at twice the rate of the other nukes, so the two of them
// are the backbone of the Forever arcane rotation.
func (mage *Mage) applyMissileBarrage() {
	if !mage.Talents.MissileBarrage {
		return
	}
	tierFrostfireChance := core.TernaryFloat64(mage.HasSetBonus(ItemSetManaflareRegalia, 5), .10, 0)

	var arcaneMissiles []*core.Spell
	mage.OnSpellRegistered(func(spell *core.Spell) {
		if spell.SpellCode == SpellCode_MageArcaneMissiles {
			arcaneMissiles = append(arcaneMissiles, spell)
		}
	})

	mage.MissileBarrageAura = mage.RegisterAura(core.Aura{
		Label:    "Missile Barrage",
		ActionID: core.ActionID{SpellID: 44404},
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range arcaneMissiles {
				spell.Cost.Multiplier -= 100
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range arcaneMissiles {
				spell.Cost.Multiplier += 100
			}
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell.SpellCode == SpellCode_MageArcaneMissiles {
				aura.Deactivate(sim)
			}
		},
	})

	mage.RegisterAura(core.Aura{
		Label:    "Missile Barrage Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			procChance := 0.0
			switch spell.SpellCode {
			case SpellCode_MageArcaneBlast:
				procChance = .40
			case SpellCode_MageFireball, SpellCode_MageFrostbolt, SpellCode_MageFrostfireBolt:
				procChance = .20
			default:
				return
			}
			if spell.SpellCode == SpellCode_MageFrostfireBolt {
				procChance += tierFrostfireChance
			}

			if sim.Proc(procChance, "Missile Barrage") {
				mage.MissileBarrageAura.Activate(sim)
			}
		},
	})
}

func (mage *Mage) registerPresenceOfMindCD() {
	if !mage.Talents.PresenceOfMind {
		return
	}

	actionID := core.ActionID{SpellID: 12043}
	cooldown := time.Second * 180

	affectedSpells := []*core.Spell{}
	pomAura := mage.RegisterAura(core.Aura{
		Label:    "Presence of Mind",
		ActionID: actionID,
		Duration: time.Second * 15,
		OnInit: func(aura *core.Aura, sim *core.Simulation) {
			for spellIdx := range mage.Spellbook {
				if spell := mage.Spellbook[spellIdx]; spell.DefaultCast.CastTime > 0 {
					affectedSpells = append(affectedSpells, spell)
				}
			}
		},
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			core.Each(affectedSpells, func(spell *core.Spell) {
				spell.CastTimeMultiplier -= 1
			})
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			core.Each(affectedSpells, func(spell *core.Spell) {
				spell.CastTimeMultiplier += 1
			})
			mage.PresenceOfMind.CD.Use(sim)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if !slices.Contains(affectedSpells, spell) {
				return
			}

			aura.Deactivate(sim)
		},
	})

	mage.PresenceOfMind = mage.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete,
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    mage.NewTimer(),
				Duration: cooldown,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return mage.GCD.IsReady(sim)
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			pomAura.Activate(sim)
		},
	})

	mage.AddMajorCooldown(core.MajorCooldown{
		Spell: mage.PresenceOfMind,
		Type:  core.CooldownTypeDPS,
	})
}

func (mage *Mage) registerArcanePowerCD() {
	if !mage.Talents.ArcanePower {
		return
	}

	actionID := core.ActionID{SpellID: 12042}

	affectedSpells := []*core.Spell{}

	mage.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagMage) {
			affectedSpells = append(affectedSpells, spell)
		}
	})

	mage.ArcanePowerAura = mage.RegisterAura(core.Aura{
		Label:    "Arcane Power",
		ActionID: actionID,
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range affectedSpells {
				spell.DamageMultiplierAdditive += 0.3
				if spell.Cost != nil {
					spell.Cost.Multiplier += 30
				}
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range affectedSpells {
				spell.DamageMultiplierAdditive -= 0.3
				if spell.Cost != nil {
					spell.Cost.Multiplier -= 30
				}
			}
		},
	})
	core.RegisterPercentDamageModifierEffect(mage.ArcanePowerAura, 1.3)

	spell := mage.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete,
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    mage.NewTimer(),
				Duration: time.Second * 180,
			},
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			mage.ArcanePowerAura.Activate(sim)
		},
	})

	mage.AddMajorCooldown(core.MajorCooldown{
		Spell: spell,
		Type:  core.CooldownTypeDPS,
	})
}

// The raid debuff version of Improved Scorch is a Classic mechanic, in Forever the fire
// vulnerability only raises the damage the mage who stacked it deals.
func (mage *Mage) applyImprovedScorch() {
	if mage.Talents.ImprovedScorch == 0 {
		return
	}

	mage.ImprovedScorchAura = mage.RegisterAura(core.Aura{
		Label:     "Improved Scorch",
		ActionID:  core.ActionID{SpellID: 12873},
		Duration:  time.Second * 30,
		MaxStacks: 5,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexFire] /= 1 + .03*float64(oldStacks)
			aura.Unit.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexFire] *= 1 + .03*float64(newStacks)
		},
	})
}

func (mage *Mage) applyMasterOfElements() {
	if mage.Talents.MasterOfElements == 0 {
		return
	}

	refundCoeff := 0.1 * float64(mage.Talents.MasterOfElements)
	manaMetrics := mage.NewManaMetrics(core.ActionID{SpellID: 29076})

	mage.RegisterAura(core.Aura{
		Label:    "Master of Elements",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !spell.SpellSchool.Matches(core.SpellSchoolFire | core.SpellSchoolFrost) {
				return
			}
			if spell.CurCast.Cost == 0 {
				return
			}
			if result.DidCrit() {
				mage.AddMana(sim, spell.Cost.BaseCost*refundCoeff, manaMetrics)
			}
		},
	})
}

// Forever aura 400625 has three cumulative stacks but one proc charge:
// the next Pyroblast consumes the whole buff, not one stack per cast.
// See assets/db_inputs/forever_effect_audit.json. Consumption is modeled on
// cast completion; interrupted-cast timing still needs a server-side check.
func (mage *Mage) applyHotStreak() {
	if !mage.Talents.HotStreak {
		return
	}

	triggerSpellCodes := []int32{SpellCode_MageFireball, SpellCode_MageFireBlast, SpellCode_MageScorch, SpellCode_MageFrostfireBolt}

	var pyroblasts []*core.Spell
	mage.OnSpellRegistered(func(spell *core.Spell) {
		if spell.SpellCode == SpellCode_MagePyroblast {
			pyroblasts = append(pyroblasts, spell)
		}
	})

	mage.HotStreakAura = mage.RegisterAura(core.Aura{
		Label:     "Hot Streak",
		ActionID:  core.ActionID{SpellID: 44445},
		Duration:  time.Second * 20,
		MaxStacks: 3,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			castTimeMultiplier := .25 * float64(newStacks-oldStacks)
			for _, spell := range pyroblasts {
				spell.CastTimeMultiplier -= castTimeMultiplier
			}
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell.SpellCode == SpellCode_MagePyroblast {
				aura.Deactivate(sim)
			}
		},
	})

	mage.RegisterAura(core.Aura{
		Label:    "Hot Streak Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.DidCrit() || !slices.Contains(triggerSpellCodes, spell.SpellCode) {
				return
			}

			mage.HotStreakAura.Activate(sim)
			mage.HotStreakAura.AddStack(sim)
		},
	})
}

// Number of non-periodic fire crits Combustion lasts for, up from 3 in Classic.
const CombustionCrits = 4

func (mage *Mage) registerCombustionCD() {
	if !mage.Talents.Combustion {
		return
	}

	actionID := core.ActionID{SpellID: 11129}
	cd := core.Cooldown{
		Timer:    mage.NewTimer(),
		Duration: time.Minute * 3,
	}

	var fireSpells []*core.Spell
	mage.OnSpellRegistered(func(spell *core.Spell) {
		if spell.SpellSchool.Matches(core.SpellSchoolFire) && spell.Flags.Matches(SpellFlagMage) {
			fireSpells = append(fireSpells, spell)
		}
	})

	numCrits := 0
	critPerStack := 10.0 * core.SpellCritRatingPerCritChance
	tierFrostfireCrit := core.TernaryFloat64(mage.HasSetBonus(ItemSetManaflareRegalia, 5), 10*core.SpellCritRatingPerCritChance, 0)

	mage.CombustionAura = mage.RegisterAura(core.Aura{
		Label:     "Combustion",
		ActionID:  actionID,
		Duration:  core.NeverExpires,
		MaxStacks: 20,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			numCrits = 0
			for _, spell := range fireSpells {
				if spell.SpellCode == SpellCode_MageFrostfireBolt {
					spell.BonusCritRating += tierFrostfireCrit
				}
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range fireSpells {
				if spell.SpellCode == SpellCode_MageFrostfireBolt {
					spell.BonusCritRating -= tierFrostfireCrit
				}
			}
			cd.Use(sim)
			mage.UpdateMajorCooldowns()
		},
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			bonusCrit := critPerStack * float64(newStacks-oldStacks)
			for _, spell := range fireSpells {
				spell.BonusCritRating += bonusCrit
			}
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() || numCrits >= CombustionCrits || !spell.SpellSchool.Matches(core.SpellSchoolFire) || !spell.Flags.Matches(SpellFlagMage) {
				return
			}

			// Ignite never consumes a crit stack, its damage isn't a cast of its own.
			if spell.SpellCode == SpellCode_MageIgnite {
				return
			}

			// TODO: This wont work properly with flamestrike
			aura.AddStack(sim)

			if result.DidCrit() {
				numCrits++
				if numCrits == CombustionCrits {
					aura.Deactivate(sim)
				}
			}
		},
	})

	spell := mage.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete,
		Cast: core.CastConfig{
			CD: cd,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return !mage.CombustionAura.IsActive()
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			mage.CombustionAura.Activate(sim)
			mage.CombustionAura.AddStack(sim)
		},
	})

	mage.AddMajorCooldown(core.MajorCooldown{
		Spell: spell,
		Type:  core.CooldownTypeDPS,
	})
}

func (mage *Mage) registerColdSnapCD() {
	if !mage.Talents.ColdSnap {
		return
	}

	// Grab all frost spells with a CD > 0
	var affectedSpells = []*core.Spell{}
	mage.OnSpellRegistered(func(spell *core.Spell) {
		if spell.SpellSchool.Matches(core.SpellSchoolFrost) && spell.CD.Duration > 0 {
			affectedSpells = append(affectedSpells, spell)
		}
	})

	spell := mage.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{SpellID: 12472},
		Flags:    core.SpellFlagNoOnCastComplete,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    mage.NewTimer(),
				Duration: time.Duration(time.Minute * 10),
			},
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			for _, spell := range affectedSpells {
				spell.CD.Reset()
			}
		},
	})

	mage.AddMajorCooldown(core.MajorCooldown{
		Spell: spell,
		Type:  core.CooldownTypeDPS,
	})
}

// Raid bosses can't be chilled or frozen, so Fingers of Frost is the only thing that gets
// Shatter and the Ice Lance bonus going on one. Shatter is folded in here because the two
// talents only ever fire together.
func (mage *Mage) applyFingersOfFrost() {
	if mage.Talents.FingersOfFrost == 0 {
		return
	}

	// The beta tooltip for rank 2 settles what the demo could not: the proc chance does not
	// scale. Both ranks give Chill effects a 15% chance; the second point buys a second
	// charge, "treats your next 2 spells cast as if the target were Frozen".
	procChance := core.TernaryFloat64(mage.Talents.FingersOfFrost > 0, .15, 0)
	tierFrostfireChance := core.TernaryFloat64(mage.HasSetBonus(ItemSetManaflareRegalia, 5), .10, 0)
	// Beta client 1.60.1: three ranks, 17/33/50%.
	shatterCrit := []float64{0, 17, 33, 50}[mage.Talents.Shatter] * core.SpellCritRatingPerCritChance

	var affectedSpells []*core.Spell
	mage.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagMage) {
			affectedSpells = append(affectedSpells, spell)
		}
	})

	// Chill effects land while the mage is already part way through the next cast. That cast is
	// not the "next spell cast" the talent grants, so it is held out of the Shatter bonus and
	// doesn't spend the charge either; the cast after it gets both.
	// TODO: the demo tooltip only says "your next 1 spell cast", beta will confirm whether a cast
	// already in progress when the chill lands counts as that one.
	var inFlight *core.Spell

	mage.FingersOfFrostAura = mage.RegisterAura(core.Aura{
		Label:    "Fingers of Frost",
		ActionID: core.ActionID{SpellID: 44543},
		Duration: time.Second * 15,
		// One charge per point: rank 2 treats the next two spells as if the target were
		// frozen rather than raising the proc chance.
		MaxStacks: int32(mage.Talents.FingersOfFrost),
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range affectedSpells {
				spell.BonusCritRating += shatterCrit
			}

			inFlight = nil
			if mage.IsCasting(sim) {
				for _, spell := range affectedSpells {
					if spell.ActionID.SameAction(mage.Hardcast.ActionID) {
						spell.BonusCritRating -= shatterCrit
						inFlight = spell
						break
					}
				}
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			if inFlight != nil {
				inFlight.BonusCritRating += shatterCrit
				inFlight = nil
			}

			for _, spell := range affectedSpells {
				spell.BonusCritRating -= shatterCrit
			}
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if !spell.Flags.Matches(SpellFlagMage) || !spell.ProcMask.Matches(core.ProcMaskSpellDamage) {
				return
			}

			if spell == inFlight {
				spell.BonusCritRating += shatterCrit
				inFlight = nil
				return
			}

			// OnCastComplete runs after the damage is rolled, so the consuming cast keeps the
			// bonus. Each cast spends one charge; the aura falls when the last one goes.
			aura.RemoveStack(sim)
		},
	})

	mage.RegisterAura(core.Aura{
		Label:    "Fingers of Frost Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() || !spell.Flags.Matches(SpellFlagChillSpell) {
				return
			}

			chance := procChance
			if spell.SpellCode == SpellCode_MageFrostfireBolt {
				chance += tierFrostfireChance
			}
			if sim.Proc(chance, "Fingers of Frost") {
				mage.FingersOfFrostAura.Activate(sim)
				mage.FingersOfFrostAura.SetStacks(sim, mage.FingersOfFrostAura.MaxStacks)
			}
		},
	})
}

// IsTargetFrozen reports whether the mage's next spell is treated as hitting a frozen target.
func (mage *Mage) IsTargetFrozen() bool {
	return mage.FingersOfFrostAura != nil && mage.FingersOfFrostAura.IsActive()
}

// The raid debuff version of Winter's Chill is a Classic mechanic, in Forever it only helps the
// mage's own Frostbolt and Ice Lance. The beta client gives 2% crit a stack, stacking once per
// talent point: "Stacks up to 5 times" at 5/5.
func (mage *Mage) applyWintersChill() {
	if mage.Talents.WintersChill == 0 {
		return
	}

	procChance := .20 * float64(mage.Talents.WintersChill)
	critPerStack := 2.0 * core.SpellCritRatingPerCritChance
	affectedSpellCodes := []int32{SpellCode_MageFrostbolt, SpellCode_MageIceLance}

	var affectedSpells []*core.Spell
	mage.OnSpellRegistered(func(spell *core.Spell) {
		if slices.Contains(affectedSpellCodes, spell.SpellCode) {
			affectedSpells = append(affectedSpells, spell)
		}
	})

	mage.WintersChillAura = mage.RegisterAura(core.Aura{
		Label:     "Winter's Chill",
		ActionID:  core.ActionID{SpellID: 28593},
		Duration:  time.Second * 15,
		MaxStacks: int32(mage.Talents.WintersChill),
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			bonusCrit := critPerStack * float64(newStacks-oldStacks)
			for _, spell := range affectedSpells {
				spell.BonusCritRating += bonusCrit
			}
		},
	})

	mage.RegisterAura(core.Aura{
		Label:    "Winters Chill Talent",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() || !spell.SpellSchool.Matches(core.SpellSchoolFrost) {
				return
			}

			if sim.Proc(procChance, "Winters Chill") {
				mage.WintersChillAura.Activate(sim)
				mage.WintersChillAura.AddStack(sim)
			}
		},
	})
}
