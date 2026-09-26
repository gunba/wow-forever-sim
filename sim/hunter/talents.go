package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func (hunter *Hunter) ApplyTalents() {
	if hunter.pet != nil {
		hunter.applyFrenzy()
		hunter.registerBestialWrathCD()
		hunter.registerIntimidationCD()

		hunter.pet.AddStat(stats.MeleeCrit, core.CritRatingPerCritChance*2*float64(hunter.Talents.Ferocity))
		hunter.pet.AddStat(stats.SpellCrit, core.SpellCritRatingPerCritChance*2*float64(hunter.Talents.Ferocity))

		hunter.pet.PseudoStats.DamageDealtMultiplier *= 1 + 0.03*float64(hunter.Talents.UnleashedFury)

		if hunter.Talents.EnduranceTraining > 0 {
			hunter.pet.MultiplyStat(stats.Health, 1+(0.03*float64(hunter.Talents.EnduranceTraining)))
		}

		if hunter.Talents.FocusedFire > 0 {
			multiplier := 1 + 0.01*float64(hunter.Talents.FocusedFire)
			if hunter.Env.IsForever() {
				// 1223755: owner and pet damage, while the pet is active.
				aura := hunter.RegisterAura(core.Aura{
					Label:    "Focused Fire",
					ActionID: core.ActionID{SpellID: 1223755},
					Duration: core.NeverExpires,
					OnGain: func(_ *core.Aura, _ *core.Simulation) {
						hunter.PseudoStats.DamageDealtMultiplier *= multiplier
						hunter.pet.PseudoStats.DamageDealtMultiplier *= multiplier
					},
					OnExpire: func(_ *core.Aura, _ *core.Simulation) {
						hunter.PseudoStats.DamageDealtMultiplier /= multiplier
						hunter.pet.PseudoStats.DamageDealtMultiplier /= multiplier
					},
				})
				hunter.pet.ApplyOnPetEnable(func(sim *core.Simulation) { aura.Activate(sim) })
				hunter.pet.ApplyOnPetDisable(func(sim *core.Simulation) { aura.Deactivate(sim) })
			} else {
				hunter.PseudoStats.DamageDealtMultiplier *= multiplier
			}
		}
	} else if hunter.Talents.LoneWolf {
		hunter.PseudoStats.DamageDealtMultiplier *= 1.2
	}

	if hunter.Talents.ImprovedTracking > 0 {
		// Client 24293 modifies tracking's second effect (damage aura 168).
		// It does not retain the old slaying talents' separate crit-damage aura.
		// Everything a raid encounter can be is trackable apart from Mechanical.
		multiplier := 1 + 0.01*float64(hunter.Talents.ImprovedTracking)
		hunter.Env.RegisterPostFinalizeEffect(func() {
			for _, t := range hunter.Env.Encounter.Targets {
				switch t.MobType {
				case proto.MobType_MobTypeBeast, proto.MobType_MobTypeDemon, proto.MobType_MobTypeDragonkin,
					proto.MobType_MobTypeElemental, proto.MobType_MobTypeGiant, proto.MobType_MobTypeHumanoid,
					proto.MobType_MobTypeUndead:
					for _, at := range hunter.AttackTables[t.UnitIndex] {
						at.DamageDealtMultiplier *= multiplier
					}
				}
			}
		})
	}

	if hunter.Talents.BestialDiscipline > 0 {
		hunter.PseudoStats.SpiritRegenRateCasting += 0.25 * float64(hunter.Talents.BestialDiscipline)

		core.MakePermanent(hunter.RegisterAura(core.Aura{
			Label: "Bestial Discipline",
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				if hunter.pet != nil {
					hunter.pet.AddFocusRegenMultiplier(0.1 * float64(hunter.Talents.BestialDiscipline))
				}
			},
		}))
	}

	hunter.AddStat(stats.MeleeHit, float64(hunter.Talents.Surefooted)*1*core.MeleeHitRatingPerHitChance)
	hunter.AddStat(stats.SpellHit, float64(hunter.Talents.Surefooted)*1*core.SpellHitRatingPerHitChance)

	hunter.AddStat(stats.MeleeCrit, float64(hunter.Talents.LethalAttacks)*1*core.CritRatingPerCritChance)
	hunter.AddStat(stats.SpellCrit, float64(hunter.Talents.LethalAttacks)*1*core.SpellCritRatingPerCritChance)

	hunter.AddStat(stats.Parry, 2*float64(hunter.Talents.Deflection))

	if hunter.Talents.CarefulAim > 0 {
		// The tooltip only says Attack Power, but melee and ranged attack power are separate stats
		// here and every other attack power buff feeds both, so Careful Aim does too.
		apPerInt := 0.2 * float64(hunter.Talents.CarefulAim)
		hunter.AddStatDependency(stats.Intellect, stats.AttackPower, apPerInt)
		hunter.AddStatDependency(stats.Intellect, stats.RangedAttackPower, apPerInt)
	}

	if hunter.Talents.RangedWeaponSpecialization > 0 {
		mult := 1 + 0.01*float64(hunter.Talents.RangedWeaponSpecialization)
		hunter.OnSpellRegistered(func(spell *core.Spell) {
			if spell.ProcMask.Matches(core.ProcMaskRanged) && spell.SpellCode != SpellCode_HunterSerpentSting {
				spell.DamageMultiplier *= mult
			}
		})
	}

	if hunter.Talents.Survivalist > 0 {
		hunter.MultiplyStat(stats.Health, 1.0+0.02*float64(hunter.Talents.Survivalist))
	}

	if hunter.Talents.LightningReflexes > 0 {
		agiBonus := 0.02 * float64(hunter.Talents.LightningReflexes)
		hunter.MultiplyStat(stats.Agility, 1.0+agiBonus)
	}

	hunter.applyEfficiency()
	hunter.applyResourcefulness()
	hunter.applySurvivalTactics()
	hunter.applyCleverTraps()
	hunter.applySurvivalistsDiscipline()
	hunter.applyPredatorsEdge()
	hunter.applyRapidRecuperation()
	hunter.applyExposePrey()
}

func (hunter *Hunter) applyFrenzy() {
	if hunter.Talents.Frenzy == 0 {
		return
	}

	procChance := 0.2 * float64(hunter.Talents.Frenzy)

	procAura := hunter.pet.RegisterAura(core.Aura{
		Label:    "Frenzy Proc",
		ActionID: core.ActionID{SpellID: 19625},
		Duration: time.Second * 8,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.MultiplyAttackSpeed(sim, 1.3)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.MultiplyAttackSpeed(sim, 1/1.3)
		},
	})

	hunter.pet.RegisterAura(core.Aura{
		Label:    "Frenzy",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellResult *core.SpellResult) {
			if !spellResult.Outcome.Matches(core.OutcomeCrit) {
				return
			}
			if procChance == 1 || sim.RandomFloat("Frenzy") < procChance {
				procAura.Activate(sim)
			}
		},
	})
}

func (hunter *Hunter) registerBestialWrathCD() {
	if !hunter.Talents.BestialWrath {
		return
	}

	actionID := core.ActionID{SpellID: 19574}

	hunter.BestialWrathPetAura = hunter.pet.RegisterAura(core.Aura{
		Label:    "Bestial Wrath Pet",
		ActionID: actionID,
		Duration: time.Second * 18,
	}).AttachMultiplicativePseudoStatBuff(&hunter.pet.PseudoStats.DamageDealtMultiplier, 1.5)

	bwSpell := hunter.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost: 0.12,
		},

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Minute * 2,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			hunter.BestialWrathPetAura.Activate(sim)
		},
	})

	hunter.AddMajorCooldown(core.MajorCooldown{
		Spell: bwSpell,
		Type:  core.CooldownTypeDPS,
	})
}

// Bosses are immune to the stun, but the pet's next attack still gets the crit bonus, so
// Intimidation is worth pressing on cooldown.
func (hunter *Hunter) registerIntimidationCD() {
	if !hunter.Talents.Intimidation {
		return
	}

	actionID := core.ActionID{SpellID: 19577}
	bonusCrit := 100.0 * core.CritRatingPerCritChance

	hunter.IntimidationPetAura = hunter.pet.RegisterAura(core.Aura{
		Label:    "Intimidation",
		ActionID: actionID,
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.MeleeCrit, bonusCrit)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.MeleeCrit, -bonusCrit)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.Landed() {
				aura.Deactivate(sim)
			}
		},
	})

	intimidation := hunter.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagAPL,

		// 8% of base mana and a 1 min cooldown in both clients (19577).
		ManaCost: core.ManaCostOptions{
			BaseCost: 0.08,
		},

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Minute,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			hunter.IntimidationPetAura.Activate(sim)
		},
	})

	hunter.AddMajorCooldown(core.MajorCooldown{
		Spell: intimidation,
		Type:  core.CooldownTypeDPS,
	})
}

func (hunter *Hunter) mortalShots() float64 {
	return 0.06 * float64(hunter.Talents.MortalShots)
}

func (hunter *Hunter) applySurvivalTactics() {
	if hunter.Talents.SurvivalTactics == 0 {
		return
	}

	hunter.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagTrap) {
			spell.BonusHitRating += 5 * float64(hunter.Talents.SurvivalTactics)
		}
	})
}

func (hunter *Hunter) applyCleverTraps() {
	if hunter.Talents.CleverTraps == 0 {
		return
	}

	hunter.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagTrap) {
			spell.DamageMultiplier *= 1 + 0.15*float64(hunter.Talents.CleverTraps)
		}
	})
}

func (hunter *Hunter) applySurvivalistsDiscipline() {
	if hunter.Talents.SurvivalistsDiscipline == 0 {
		return
	}

	multiplier := 1 - 0.2*float64(hunter.Talents.SurvivalistsDiscipline)

	hunter.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagTrap) {
			spell.CD.Duration = time.Duration(float64(spell.CD.Duration) * multiplier)
		}
	})
}

func (hunter *Hunter) applyEfficiency() {
	if hunter.Talents.Efficiency == 0 {
		return
	}

	hunter.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Cost == nil {
			return
		}
		affected := spell.Flags.Matches(SpellFlagSting|SpellFlagShot) || spell.ProcMask.Matches(core.ProcMaskMeleeSpecial)
		if hunter.Env.IsForever() {
			// 19416, effect 696992. Includes Hawk and Volley, not Sniper or Strider.
			switch spell.SpellCode {
			case SpellCode_HunterAimedShot, SpellCode_HunterArcaneShot, SpellCode_HunterMultiShot,
				SpellCode_HunterSerpentSting, SpellCode_HunterRaptorStrike, SpellCode_HunterRaptorStrikeHit,
				SpellCode_HunterMongooseBite, SpellCode_HunterWingClip, SpellCode_HunterLaceratingStrikes,
				SpellCode_HunterSummonHawk, SpellCode_HunterVolley:
				affected = true
			default:
				affected = false
			}
		}
		if affected {
			spell.Cost.Multiplier -= 3 * hunter.Talents.Efficiency
		}
	})
}

func (hunter *Hunter) applyResourcefulness() {
	if hunter.Talents.Resourcefulness == 0 {
		return
	}

	// Client curves: 30/60% cost and a 50/100% proc chance. The buff (1242688) is 50% for 30 sec
	// at both ranks.
	costReduction := 30 * hunter.Talents.Resourcefulness
	procChance := 0.5 * float64(hunter.Talents.Resourcefulness)

	hunter.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Cost == nil {
			return
		}
		affected := spell.Flags.Matches(SpellFlagTrap) || spell.ProcMask.Matches(core.ProcMaskMeleeSpecial)
		if hunter.Env.IsForever() {
			// 440529, effect 1134467: traps, Raptor, Wing Clip and Mongoose.
			switch spell.SpellCode {
			case SpellCode_HunterExplosiveTrap, SpellCode_HunterFreezingTrap, SpellCode_HunterImmolationTrap,
				SpellCode_HunterRaptorStrike, SpellCode_HunterRaptorStrikeHit,
				SpellCode_HunterMongooseBite, SpellCode_HunterWingClip:
				affected = true
			default:
				affected = false
			}
		}
		if affected {
			spell.Cost.Multiplier -= costReduction
		}
	})

	procAura := hunter.RegisterAura(core.Aura{
		Label:    "Resourcefulness",
		ActionID: core.ActionID{SpellID: 1242688},
		Duration: time.Second * 30,
	}).AttachSpiritRegenRateCastingBuff(0.5)

	core.MakePermanent(hunter.RegisterAura(core.Aura{
		Label: "Resourcefulness Trigger",
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.DidCrit() && sim.Proc(procChance, "Resourcefulness") {
				procAura.Activate(sim)
			}
		},
	}))
}

func (hunter *Hunter) applyPredatorsEdge() {
	if hunter.Talents.PredatorsEdge == 0 {
		return
	}

	critDamageBonus := 0.06 * float64(hunter.Talents.PredatorsEdge)
	ohMultiplier := 1 + 0.1*float64(hunter.Talents.PredatorsEdge)

	hunter.OnSpellRegistered(func(spell *core.Spell) {
		critAffected := spell.DefenseType == core.DefenseTypeMelee
		if hunter.Env.IsForever() {
			// 1310627, effect 1340259 is a spell-family crit-damage modifier.
			switch spell.SpellCode {
			case SpellCode_HunterRaptorStrike, SpellCode_HunterRaptorStrikeHit, SpellCode_HunterMongooseBite,
				SpellCode_HunterWingClip, SpellCode_HunterLaceratingStrikes, SpellCode_HunterStriderKick:
				critAffected = true
			default:
				critAffected = false
			}
		}
		if critAffected {
			spell.CritDamageBonus += critDamageBonus
		}
		if spell.ProcMask.Matches(core.ProcMaskMeleeOH) && spell.BonusCoefficient > 0 {
			spell.DamageMultiplier *= ohMultiplier
		}
	})
}

// Only the Serpent Sting half is modelled, nothing dies mid fight to hand out Rapid Killing.
func (hunter *Hunter) applyRapidRecuperation() {
	if hunter.Talents.RapidRecuperation == 0 {
		return
	}

	// Client curve 25/50% for Serpent Sting; the buff (1242512) lasts 15 sec at both ranks.
	procAura := hunter.RegisterAura(core.Aura{
		Label:    "Rapid Recuperation",
		ActionID: core.ActionID{SpellID: 1242512},
		Duration: time.Second * 15,
	}).AttachSpiritRegenRateCastingBuff(0.25 * float64(hunter.Talents.RapidRecuperation))

	core.MakePermanent(hunter.RegisterAura(core.Aura{
		Label: "Rapid Recuperation Trigger",
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.SpellCode == SpellCode_HunterSerpentSting && result.Landed() {
				procAura.Activate(sim)
			}
		},
	}))
}

func (hunter *Hunter) applyExposePrey() {
	if hunter.Talents.ExposePrey == 0 {
		return
	}

	// Client curve 5/10%; the Mongoose Bite window (1310726) is 5 sec at both ranks.
	procChance := 0.05 * float64(hunter.Talents.ExposePrey)

	core.MakePermanent(hunter.RegisterAura(core.Aura{
		Label: "Expose Prey",
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			// Client 1310532 uses proc mask 0x154: melee/ranged autos and
			// attacks, not generic spell hits such as trap damage.
			if !result.Landed() || !spell.ProcMask.Matches(core.ProcMaskMelee|core.ProcMaskRanged) ||
				!result.Target.HasActiveAuraWithTag(core.HuntersMarkAuraTag) {
				return
			}
			if sim.Proc(procChance, "Expose Prey") {
				hunter.DefensiveState.Activate(sim)
			}
		},
	}))
}
