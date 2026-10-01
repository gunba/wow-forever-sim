package common

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func applyStrikingEnchant(agent core.Agent, slot proto.ItemSlot, bonus float64) {
	character := agent.GetCharacter()
	if character.Env.IsForever() {
		// Read the actual equipped enchant, including when a different swap
		// enchant's registration is being visited during initialization.
		if slot == proto.ItemSlot_ItemSlotOffHand {
			character.AutoAttacks.SetOH(character.WeaponFromOffHand())
		} else {
			character.AutoAttacks.SetMH(character.WeaponFromMainHand())
		}
		return
	}
	weapon := character.AutoAttacks.MH()
	if slot == proto.ItemSlot_ItemSlotOffHand {
		weapon = character.AutoAttacks.OH()
	}
	weapon.BaseDamageMin += bonus
	weapon.BaseDamageMax += bonus
}

func crusaderStrengthBonus(level int32) float64 {
	return max(0, 100-4*float64(max(0, level-60)))
}

func init() {
	core.AddEffectsToTest = false

	///////////////////////////////////////////////////////////////////////////
	//                        All effects ordered by ID
	///////////////////////////////////////////////////////////////////////////

	// Ranged Scopes
	core.AddWeaponEffect(32, func(agent core.Agent, _ proto.ItemSlot) {
		w := agent.GetCharacter().AutoAttacks.Ranged()
		w.BaseDamageMin += 2
		w.BaseDamageMax += 2
	})

	// Accurate Scope
	core.AddWeaponEffect(33, func(agent core.Agent, _ proto.ItemSlot) {
		w := agent.GetCharacter().AutoAttacks.Ranged()
		w.BaseDamageMin += 3
		w.BaseDamageMax += 3
	})

	// Weapon - Fiery Blaze
	core.NewEnchantEffect(36, func(agent core.Agent) {
		character := agent.GetCharacter()

		procMask := character.GetProcMaskForEnchant(36)
		procChance := 0.15

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 6296},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					damage := sim.Roll(9, 13)
					spell.CalcAndDealDamage(sim, aoeTarget, damage, spell.OutcomeMagicHitAndCrit)
				}

			},
		})

		aura := character.GetOrRegisterAura(core.Aura{
			Label:    "Fiery Blaze",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || !spell.ProcMask.Matches(procMask) || spell.Flags.Matches(core.SpellFlagSuppressWeaponProcs) {
					return
				}

				if sim.RandomFloat("Fiery Blaze") < procChance {
					procSpell.Cast(sim, result.Target)
				}
			},
		})

		character.ItemSwap.RegisterOnSwapItemForEffect(36, aura)
	})

	// Weapon - Lesser Striking
	core.AddWeaponEffect(241, func(agent core.Agent, slot proto.ItemSlot) {
		applyStrikingEnchant(agent, slot, 2)
	})

	// Weapon - Beast Slaying
	core.AddWeaponEffect(249, func(agent core.Agent, slot proto.ItemSlot) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeBeast {
			w := character.AutoAttacks.MH()
			if slot == proto.ItemSlot_ItemSlotOffHand {
				w = character.AutoAttacks.OH()
			}

			w.BaseDamageMin += 2
			w.BaseDamageMax += 2

			w = character.AutoAttacks.Ranged()
			w.BaseDamageMin += 2
			w.BaseDamageMax += 2
		}
	})

	// Weapon - Minor Striking
	core.AddWeaponEffect(250, func(agent core.Agent, slot proto.ItemSlot) {
		applyStrikingEnchant(agent, slot, 1)
	})

	// Deadly Scope
	core.AddWeaponEffect(663, func(agent core.Agent, _ proto.ItemSlot) {
		w := agent.GetCharacter().AutoAttacks.Ranged()
		w.BaseDamageMin += 5
		w.BaseDamageMax += 5
	})

	// Sniper Scope
	core.AddWeaponEffect(664, func(agent core.Agent, _ proto.ItemSlot) {
		w := agent.GetCharacter().AutoAttacks.Ranged()
		w.BaseDamageMin += 7
		w.BaseDamageMax += 7
	})

	// Weapon - Fiery Weapon
	core.AddWeaponEffect(803, func(agent core.Agent, _ proto.ItemSlot) {
		character := agent.GetCharacter()

		procMask := character.GetProcMaskForEnchant(803)
		ppmm := character.AutoAttacks.NewPPMManager(6.0, procMask)

		procMaskOnAuto := core.ProcMaskDamageProc     // Both spell and melee proc combo
		procMaskOnSpecial := core.ProcMaskSpellDamage // TODO: check if core.ProcMaskSpellDamage remains on special

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 13897},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    procMaskOnAuto,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 40, spell.OutcomeMagicHitAndCrit)
			},
		})

		aura := core.MakePermanent(character.GetOrRegisterAura(core.Aura{
			Label: "Fiery Weapon",
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || spell.Flags.Matches(core.SpellFlagSuppressWeaponProcs) {
					return
				}
				if ppmm.Proc(sim, spell.ProcMask, "Fiery Weapon") {
					if spell.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
						procSpell.ProcMask = procMaskOnSpecial
					} else {
						procSpell.ProcMask = procMaskOnAuto
					}
					procSpell.Cast(sim, result.Target)
				}
			},
		}))

		character.ItemSwap.RegisterOnSwapItemForEffectWithPPMManager(803, 6.0, &ppmm, aura)
	})

	// Weapon - Greater Striking
	core.AddWeaponEffect(805, func(agent core.Agent, slot proto.ItemSlot) {
		applyStrikingEnchant(agent, slot, 4)
	})

	// Weapon - Lesser Beastslayer
	core.AddWeaponEffect(853, func(agent core.Agent, slot proto.ItemSlot) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeBeast {
			w := character.AutoAttacks.MH()
			if slot == proto.ItemSlot_ItemSlotOffHand {
				w = character.AutoAttacks.OH()
			}

			w.BaseDamageMin += 6
			w.BaseDamageMax += 6

			w = character.AutoAttacks.Ranged()
			w.BaseDamageMin += 6
			w.BaseDamageMax += 6
		}
	})

	// Weapon - Lesser Elemental Slayer
	core.AddWeaponEffect(854, func(agent core.Agent, slot proto.ItemSlot) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeElemental {
			w := character.AutoAttacks.MH()
			if slot == proto.ItemSlot_ItemSlotOffHand {
				w = character.AutoAttacks.OH()
			}

			w.BaseDamageMin += 6
			w.BaseDamageMax += 6

			w = character.AutoAttacks.Ranged()
			w.BaseDamageMin += 6
			w.BaseDamageMax += 6
		}
	})

	// Boots - Minor Speed
	core.NewEnchantEffect(911, func(agent core.Agent) {
		character := agent.GetCharacter()

		character.RegisterAura(core.Aura{
			Label: "Minor Speed",
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				character.AddMoveSpeedModifier(&core.ActionID{SpellID: 13889}, 1.08)
			},
		})
	})

	// Weapon - Striking
	core.AddWeaponEffect(943, func(agent core.Agent, slot proto.ItemSlot) {
		applyStrikingEnchant(agent, slot, 3)
	})

	// Weapon - Superior Striking
	core.AddWeaponEffect(1897, func(agent core.Agent, slot proto.ItemSlot) {
		applyStrikingEnchant(agent, slot, 5)
	})

	// Weapon - Lifestealing
	core.AddWeaponEffect(1898, func(agent core.Agent, slot proto.ItemSlot) {
		character := agent.GetCharacter()

		procMask := character.GetProcMaskForEnchant(1898)
		ppmm := character.AutoAttacks.NewPPMManager(6.66, procMask)
		healthMetrics := character.NewHealthMetrics(core.ActionID{SpellID: 20004})

		procMaskOnAuto := core.ProcMaskDamageProc     // Both spell and melee proc combo
		procMaskOnSpecial := core.ProcMaskSpellDamage // TODO: check if core.ProcMaskSpellDamage remains on special

		procSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 20004},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    procMaskOnAuto,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 30, spell.OutcomeMagicHitAndCrit)
				if character.Env.IsForever() {
					// Effect 697642 is health leech, not damage alone.
					character.GainHealth(sim, result.Damage, healthMetrics)
				}
			},
		})

		aura := core.MakePermanent(character.GetOrRegisterAura(core.Aura{
			Label: "Lifestealing Weapon",
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && !spell.Flags.Matches(core.SpellFlagSuppressWeaponProcs) && ppmm.Proc(sim, spell.ProcMask, "Lifestealing Weapon") {
					if spell.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
						procSpell.ProcMask = procMaskOnSpecial
					} else {
						procSpell.ProcMask = procMaskOnAuto
					}
					procSpell.Cast(sim, result.Target)
				}
			},
		}))

		character.ItemSwap.RegisterOnSwapItemForEffectWithPPMManager(1898, 6.66, &ppmm, aura)
	})

	// TODO: Crusader, Mongoose, and Executioner could also be modelled as AddWeaponEffect instead
	// ApplyCrusaderEffect will be applied twice if there is two weapons with this enchant.
	//   However, it will automatically overwrite one of them, so it should be ok.
	//   A single application of the aura will handle both mh and oh procs.
	core.NewEnchantEffect(1900, func(agent core.Agent) {
		character := agent.GetCharacter()

		procMask := character.GetProcMaskForEnchant(1900)
		ppmm := character.AutoAttacks.NewPPMManager(1.0, procMask)

		// -4 Strength per level above 60, never a bonus below 60.
		strBonus := crusaderStrengthBonus(character.Level)
		healthMetrics := character.NewHealthMetrics(core.ActionID{SpellID: 20007})
		mhAura := character.NewTemporaryStatsAura("Crusader Enchant MH", core.ActionID{SpellID: 20007, Tag: 1}, stats.Stats{stats.Strength: strBonus}, time.Second*15)
		ohAura := character.NewTemporaryStatsAura("Crusader Enchant OH", core.ActionID{SpellID: 20007, Tag: 2}, stats.Stats{stats.Strength: strBonus}, time.Second*15)

		aura := character.GetOrRegisterAura(core.Aura{
			Label:    "Crusader Enchant",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if !result.Landed() || spell.Flags.Matches(core.SpellFlagSuppressWeaponProcs) {
					return
				}
				if ppmm.Proc(sim, spell.ProcMask, "Crusader") {
					if character.Env.IsForever() {
						// Effect 697923: 100 mean, 0.5 variance. Match existing
						// item health-return accounting; proc healing threat remains unverified.
						character.GainHealth(sim, sim.Roll(75, 125), healthMetrics)
					}
					if spell.IsMH() {
						mhAura.Activate(sim)
					} else {
						ohAura.Activate(sim)
					}
				}
			},
		})

		character.ItemSwap.RegisterOnSwapItemForEffectWithPPMManager(1900, 1.0, &ppmm, aura)
	})

	// Biznicks 247x128 Accurascope
	core.AddWeaponEffect(2523, func(agent core.Agent, _ proto.ItemSlot) {
		character := agent.GetCharacter()
		character.AddBonusRangedHitRating(3)
	})

	// Gloves - Threat
	core.NewEnchantEffect(2613, func(agent core.Agent) {
		character := agent.GetCharacter()

		character.RegisterAura(core.Aura{
			Label: "Threat +2%",
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.ThreatMultiplier *= 1.02
			},
		})
	})

	// Cloak - Subtlety
	core.NewEnchantEffect(2621, func(agent core.Agent) {
		character := agent.GetCharacter()

		character.RegisterAura(core.Aura{
			Label: "Subtlety",
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.ThreatMultiplier /= 1.02
			},
		})
	})

	core.AddEffectsToTest = true
}
