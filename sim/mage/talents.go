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
	mage.applyImpact()
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

	mage.applyFirePower()
}

// Fire Power, redesigned in client 1.60.1.70425: 1/2/3% more Fire damage per Fire effect of the mage
// on the target, counting Ignite, Improved Scorch, the Fireball and Pyroblast damage over time and
// Impact, up to four. Mixed-school spells (Frostfire Bolt) benefit but do not count. Ignite itself is
// left out as before.
func (mage *Mage) applyFirePower() {
	if mage.Talents.FirePower == 0 {
		return
	}
	perEffect := 0.01 * float64(mage.Talents.FirePower)
	dotActive := func(spell *core.Spell, target *core.Unit) bool {
		if spell == nil || int(target.UnitIndex) >= len(spell.Dots()) {
			return false
		}
		dot := spell.Dot(target)
		return dot != nil && dot.IsActive()
	}
	anyDot := func(spells []*core.Spell, target *core.Unit) bool {
		for _, spell := range spells {
			if dotActive(spell, target) {
				return true
			}
		}
		return false
	}
	effects := func(target *core.Unit) int {
		n := 0
		if dotActive(mage.igniteTick, target) { // the Ignite DoT lives on its tick spell
			n++
		}
		if mage.ImprovedScorchAura != nil && mage.ImprovedScorchAura.IsActive() {
			n++
		}
		if anyDot(mage.Fireball, target) {
			n++
		}
		if anyDot(mage.Pyroblast, target) {
			n++
		}
		if len(mage.ImpactAuras) > 0 && mage.ImpactAuras.Get(target) != nil && mage.ImpactAuras.Get(target).IsActive() {
			n++
		}
		return min(n, 4)
	}
	for _, target := range mage.Env.Encounter.TargetUnits {
		target.AddDynamicDamageTakenModifier(func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.Unit != &mage.Unit || !spell.SpellSchool.Matches(core.SpellSchoolFire) || !spell.Flags.Matches(SpellFlagMage) ||
				spell.SpellCode == SpellCode_MageIgnite {
				return
			}
			if n := effects(result.Target); n > 0 {
				result.Damage *= 1 + perEffect*float64(n)
			}
		})
	}
}

// Client 1.60.1.70291: node105792/definition135322 curve83091 is3/7/10.
// Spell11103 triggers the2s stun12355 on direct hits (proc mask0x15550,
// no periodic flag). CAN_PROC_FROM_PROCS permits direct Mage-family fire procs.
// https://us.forums.blizzard.com/en/wow/t/2360696/5
func (mage *Mage) applyImpact() {
	if !mage.Env.IsForever() || mage.Talents.Impact == 0 {
		return
	}
	chance := []float64{0, .03, .07, .10}[mage.Talents.Impact]
	mage.ImpactAuras = mage.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return target.RegisterCrowdControlAura(core.Aura{
			Label:    "Impact-" + mage.Label,
			ActionID: core.ActionID{SpellID: 12355},
			Duration: 2 * time.Second,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStun(sim)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.RemoveStun(sim)
			},
		}, core.CrowdControlDRStunProc)
	})
	core.MakePermanent(mage.RegisterAura(core.Aura{
		Label:    "Impact Trigger",
		ActionID: core.ActionID{SpellID: 11103},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() || !result.Target.PseudoStats.CanBeStunned ||
				!spell.Flags.Matches(SpellFlagMage) || !spell.SpellSchool.Matches(core.SpellSchoolFire) ||
				!spell.ProcMask.Matches(core.ProcMaskSpellOrSpellProc|core.ProcMaskSpellDamageProc) {
				return
			}
			if stun := mage.ImpactAuras.Get(result.Target); stun != nil && sim.Proc(chance, "Impact") {
				result.Target.ApplyCrowdControl(sim, stun, 2*time.Second, core.CrowdControlDRStunProc)
			}
		},
	}))
}

func (mage *Mage) applyFrostTalents() {
	mage.registerColdSnapCD()
	mage.registerIceBarrierSpell()
	mage.applyFingersOfFrost()
	mage.applyWintersChill()

	// Elemental Precision
	if mage.Talents.ElementalPrecision > 0 {
		bonusHit := 1 * float64(mage.Talents.ElementalPrecision) * core.SpellHitRatingPerHitChance
		if mage.Env.IsForever() {
			// Three ranks of 2/4/6% since the server hotfix (CurvePoint 241098-241100 on curve 83078).
			bonusHit = 2 * float64(mage.Talents.ElementalPrecision) * core.SpellHitRatingPerHitChance
		}

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

	var lastSpell *core.Spell
	var lastAt time.Duration = -1
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
			// Forever: the missiles don't roll Clearcasting either (a log has 0 gains from 34 Arcane Missiles casts,
			// wowsims/forever #489).
			if mage.Env.IsForever() && spell.SpellCode == SpellCode_MageArcaneMissilesTick {
				return
			}
			// Client 1.60.1.70425: one roll per cast, not one per target struck.
			if mage.Env.IsForever() {
				if lastSpell == spell && lastAt == sim.CurrentTime {
					return
				}
				lastSpell, lastAt = spell, sim.CurrentTime
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

// Blizzard's client E_DUMMY is an area-targeted magic application on the
// paid parent. Its triggered damage is not a new Arcane Concentration roll.
// Keep this application local to AC: no synthetic damage, hit counters or
// generic callbacks for Winter's Chill, Fingers of Frost, raid or gear procs.
func (mage *Mage) rollBlizzardArcaneConcentration(sim *core.Simulation, spell *core.Spell) {
	if !mage.Env.IsForever() || mage.Talents.ArcaneConcentration == 0 {
		return
	}
	// Client 1.60.1.70425: Clearcasting no longer rolls for every target a spell strikes. The
	// trigger rolls once per cast, on the first target the wave lands on.
	trigger := mage.GetAura("Arcane Concentration")
	for _, target := range sim.Encounter.TargetUnits {
		result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
		trigger.OnSpellHitDealt(trigger, sim, spell, result)
		spell.DisposeResult(result)
	}
}

// Arcane Blast feeds Missile Barrage at twice the rate of the other nukes, so the two of them
// are the backbone of the Forever arcane rotation.
func (mage *Mage) applyMissileBarrage() {
	if mage.Talents.MissileBarrage == 0 {
		return
	}
	// Client 1.60.1.70425: two points; Arcane Blast 20/40%, and Arcane Wrath, Fireball, Frostbolt, Frost Blast
	// and Frostfire Bolt 10/20%. The missiles proc nothing (hotfixed tooltip).
	points := float64(mage.Talents.MissileBarrage)
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

	tryProc := func(sim *core.Simulation, spell *core.Spell) {
		procChance := 0.0
		switch spell.SpellCode {
		case SpellCode_MageArcaneBlast:
			procChance = .20 * points
		case SpellCode_MageArcaneMissilesTick:
			if mage.Env.IsForever() {
				// A developer confirmed Arcane Missiles does not proc it; the hotfixed tooltip drops the
				// missiles (wowsims/forever #479).
				return
			}
			procChance = .05 * points
		case SpellCode_MageArcaneWrath, SpellCode_MageFireball, SpellCode_MageFrostbolt, SpellCode_MageFrostfireBolt:
			procChance = .10 * points
		default:
			return
		}
		if spell.SpellCode == SpellCode_MageFrostfireBolt {
			procChance += tierFrostfireChance
		}
		if sim.Proc(procChance, "Missile Barrage") {
			mage.MissileBarrageAura.Activate(sim)
		}
	}

	trigger := core.Aura{
		Label:    "Missile Barrage Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
	}
	if mage.Env.IsForever() {
		// Client 1.60.1.70245, SpellAuraOptions 400588: harmful spell hit,
		// not cast completion. A projectile must land before it can proc.
		trigger.OnSpellHitDealt = func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.Landed() && spell.ProcMask.Matches(core.ProcMaskSpellDamage) && !mage.ProcsAtCast(sim, spell) {
				tryProc(sim, spell)
			}
		}
		mage.castResultHooks = append(mage.castResultHooks, tryProc)
	} else {
		trigger.OnCastComplete = func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			tryProc(sim, spell)
		}
	}
	mage.RegisterAura(trigger)
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

			mage.pomUsedBy, mage.pomUsedAt = spell, sim.CurrentTime
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
	var icd *core.Cooldown
	if mage.Env.IsForever() {
		// SpellAuraOptions 29074: ProcCategoryRecovery is milliseconds.
		// Its 9ms ICD coalesces an AoE batch; Arcane Concentration is 1000ms.
		icd = &core.Cooldown{Timer: mage.NewTimer(), Duration: 9 * time.Millisecond}
	}

	mage.RegisterAura(core.Aura{
		Label:    "Master of Elements",
		Duration: core.NeverExpires,
		Icd:      icd,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !spell.SpellSchool.Matches(core.SpellSchoolFire|core.SpellSchoolFrost) ||
				!result.Landed() || !result.DidCrit() || spell.Cost == nil || spell.Cost.BaseCost == 0 {
				return
			}
			if !mage.Env.IsForever() && spell.CurCast.Cost == 0 {
				return
			}
			if icd != nil {
				if !icd.IsReady(sim) {
					return
				}
				icd.Use(sim)
			}
			// The base-cost refund also applies to a free Forever cast.
			mage.AddMana(sim, spell.Cost.BaseCost*refundCoeff, manaMetrics)
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
		Label:     "Heating Up",
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
			// Client 1.60.1.70425: a Pyroblast cast instantly by Presence of Mind keeps the stacks.
			if spell.SpellCode == SpellCode_MagePyroblast && !mage.presenceOfMindUsedBy(sim, spell) {
				aura.Deactivate(sim)
			}
		},
	})

	mage.RegisterAura(core.Aura{
		Label:    "Heating Up Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !slices.Contains(triggerSpellCodes, spell.SpellCode) {
				return
			}
			// Client 1.60.1.70425: Fire Blast adds a stack every time it deals damage, so a critical Fire
			// Blast adds two.
			stacks := int32(0)
			if result.DidCrit() {
				stacks++
			}
			if spell.SpellCode == SpellCode_MageFireBlast && result.Landed() {
				stacks++
			}
			if stacks == 0 {
				return
			}
			mage.HotStreakAura.Activate(sim)
			mage.HotStreakAura.SetStacks(sim, min(mage.HotStreakAura.GetStacks()+stacks, mage.HotStreakAura.MaxStacks))
		},
	})
}

// October 1 restores three critical strikes per Combustion activation.
const CombustionCrits = 3

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
	var roll func(sim *core.Simulation, spell *core.Spell)
	// Forever grants the charges when Frostbolt or Frostfire Bolt finishes casting (see fingersOfFrostOnCastResult);
	// that cast must not spend them.
	var grantSpell *core.Spell
	var grantAt time.Duration

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
			if spell == grantSpell && sim.CurrentTime == grantAt {
				return
			}

			// OnCastComplete runs after the damage is rolled, so the consuming cast keeps the
			// bonus. Each cast spends one charge; the aura falls when the last one goes.
			if spell.SpellCode == SpellCode_MageIceLance {
				mage.empoweredLances++
			}
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
			if !result.Landed() || !spell.Flags.Matches(SpellFlagChillSpell) || mage.ProcsAtCast(sim, spell) {
				return
			}
			roll(sim, spell)
		},
	})

	roll = func(sim *core.Simulation, spell *core.Spell) {
		chance := procChance
		if spell.SpellCode == SpellCode_MageFrostfireBolt {
			chance += tierFrostfireChance
		}
		if sim.Proc(chance, "Fingers of Frost") {
			grantSpell, grantAt = spell, sim.CurrentTime
			mage.FingersOfFrostAura.Activate(sim)
			mage.FingersOfFrostAura.SetStacks(sim, mage.FingersOfFrostAura.MaxStacks)
		}
	}
	mage.castResultHooks = append(mage.castResultHooks, func(sim *core.Simulation, spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagChillSpell) {
			roll(sim, spell)
		}
	})
}

// Forever rolls the caster's on-hit procs from Fireball, Frostbolt and Frostfire Bolt when the cast
// completes and its roll does not miss, before the projectile lands. Public foreverlogs.gg reports show
// every Frostbolt-triggered Fingers of Frost (121 of 121 outside encounter-start snapshots) and every
// Frostbolt/Fireball Missile Barrage (11 of 11) at the cast-success timestamp, none at impact.
func (mage *Mage) ProcsAtCast(sim *core.Simulation, spell *core.Spell) bool {
	return sim.IsForever() && (spell.SpellCode == SpellCode_MageFrostbolt ||
		spell.SpellCode == SpellCode_MageFrostfireBolt || spell.SpellCode == SpellCode_MageFireball)
}

func (mage *Mage) CastResultProcs(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	if !result.Landed() || !mage.ProcsAtCast(sim, spell) {
		return
	}
	for _, hook := range mage.castResultHooks {
		hook(sim, spell)
	}
}

// IsTargetFrozen reports whether the mage's next spell is treated as hitting a frozen target.
func (mage *Mage) IsTargetFrozen() bool {
	return mage.FingersOfFrostAura != nil && mage.FingersOfFrostAura.IsActive()
}

// Winter's Chill after the server hotfix (game cache push 112694-112753): "Gives your Frostbolt, Frost Blast, and
// Fingers of Frost empowered Ice Lance spells a $m2% chance to apply Winter's Chill, increasing the target's damage
// taken from your next Cone of Cold by $12579s1% and it gains an additional $s3% of your spell damage, stacking up to
// $s1 times." 20% a point; 5% damage and 0.07 spell power coefficient a stack (12579 effect 1 5, effect 2 7; 11180
// effect 3 7); stacks to the talent's rank. The next Cone of Cold spends them.
func (mage *Mage) applyWintersChill() {
	if mage.Talents.WintersChill == 0 {
		return
	}

	procChance := .20 * float64(mage.Talents.WintersChill)
	damagePerStack, coefficientPerStack := 0.40, 0.0
	maxStacks := int32(5)
	if mage.Env.IsForever() {
		damagePerStack, coefficientPerStack = 0.05, 0.07
		maxStacks = mage.Talents.WintersChill
	}
	var coneOfCold []*core.Spell
	mage.OnSpellRegistered(func(spell *core.Spell) {
		if spell.SpellCode == SpellCode_MageConeOfCold {
			coneOfCold = append(coneOfCold, spell)
		}
	})

	mage.WintersChillAura = mage.RegisterAura(core.Aura{
		Label:     "Winter's Chill",
		ActionID:  core.ActionID{SpellID: 12579},
		Duration:  time.Second * 15,
		MaxStacks: maxStacks,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			for _, spell := range coneOfCold {
				spell.DamageMultiplierAdditive += damagePerStack * float64(newStacks-oldStacks)
				spell.BonusCoefficient += coefficientPerStack * float64(newStacks-oldStacks)
			}
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell.SpellCode == SpellCode_MageConeOfCold {
				aura.Deactivate(sim)
			}
		},
	})

	applies := func(spell *core.Spell) bool {
		if !mage.Env.IsForever() {
			return spell.SpellSchool.Matches(core.SpellSchoolFrost) && spell.SpellCode != SpellCode_MageConeOfCold
		}
		return spell.SpellCode == SpellCode_MageFrostbolt
	}

	mage.RegisterAura(core.Aura{
		Label:    "Winters Chill Talent",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			mage.empoweredLances = 0
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			eligible := applies(spell)
			if spell.SpellCode == SpellCode_MageIceLance && mage.empoweredLances > 0 {
				// An Ice Lance cast while Fingers of Frost was up.
				mage.empoweredLances--
				eligible = mage.Env.IsForever()
			}
			if !result.Landed() || !eligible {
				return
			}

			if sim.Proc(procChance, "Winters Chill") {
				mage.WintersChillAura.Activate(sim)
				mage.WintersChillAura.AddStack(sim)
			}
		},
	})
}

// presenceOfMindUsedBy reports whether Presence of Mind made this cast instant, whichever of the two
// cast-complete callbacks runs first.
func (mage *Mage) presenceOfMindUsedBy(sim *core.Simulation, spell *core.Spell) bool {
	return mage.pomUsedBy == spell && mage.pomUsedAt == sim.CurrentTime
}
