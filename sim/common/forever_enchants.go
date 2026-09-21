package common

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func init() {
	// 1310308 is a 2% crit aura restricted by SpellEquippedItems to bows,
	// guns and crossbows (subclass mask 262156), not general melee/spell crit.
	core.NewEnchantEffect(8720, func(agent core.Agent) {
		c := agent.GetCharacter()
		const label = "SAF-T Ultra Precision Scope"
		if c.GetAura(label) != nil {
			return
		}
		equipped := func() bool { return c.Equipment[proto.ItemSlot_ItemSlotRanged].Enchant.EffectID == 8720 }
		adjust := func(amount float64) {
			for _, spell := range c.Spellbook {
				if spell.CastType == proto.CastType_CastTypeRanged {
					spell.BonusCritRating += amount
				}
			}
		}
		applied := false
		aura := c.RegisterAura(core.Aura{
			Label: label, ActionID: core.ActionID{ItemID: 279272},
			Duration: core.NeverExpires, BuildPhase: core.CharacterBuildPhaseGear,
			OnGain: func(_ *core.Aura, _ *core.Simulation) {
				applied = equipped()
				if applied {
					adjust(2)
				}
			},
			OnExpire: func(_ *core.Aura, _ *core.Simulation) {
				if applied {
					adjust(-2)
				}
				applied = false
			},
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				if equipped() {
					aura.Activate(sim)
				}
			},
		})
		c.OnSpellRegistered(func(spell *core.Spell) {
			if applied && spell.CastType == proto.CastType_CastTypeRanged {
				spell.BonusCritRating += 2
			}
		})
		c.RegisterOnItemSwap(func(sim *core.Simulation) {
			aura.Deactivate(sim)
			if equipped() {
				aura.Activate(sim)
			}
		})
	})
	// Client 69893: 22841 has melee/ranged speed auras 319/140;
	// 7217 has only melee speed. Neither is general haste or Energy haste.
	registerAttackSpeedEnchant(2543, 22841, "Arcanum of Rapidity", 1, true)
	registerAttackSpeedEnchant(34, 7217, "Iron Counterweight", 3, false)
	// Minor Haste (13928) also has casting-speed aura 65. Its melee haste
	// remains in the enchant's stats and participates in the Energy model.
	core.NewEnchantEffect(931, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.Env.IsForever() {
			character.AddStat(stats.SpellHaste, 1)
		}
	})
}

func registerAttackSpeedEnchant(id, spellID int32, label string, percent float64, ranged bool) {
	core.NewEnchantEffect(id, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.GetAura(label) != nil {
			return
		}
		count := func() float64 {
			n := 0.0
			for _, item := range character.Equipment {
				if item.Enchant.EffectID == id {
					n++
				}
			}
			return n
		}
		var haste, multiplier float64
		aura := character.RegisterAura(core.Aura{
			Label: label, ActionID: core.ActionID{SpellID: spellID},
			Duration: core.NeverExpires, BuildPhase: core.CharacterBuildPhaseGear,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				if count() > 0 {
					aura.Activate(sim)
				}
			},
			OnGain: func(_ *core.Aura, sim *core.Simulation) {
				haste = percent * count()
				multiplier = 1 + haste/100
				if !character.Env.IsForever() {
					character.AddStatDynamic(sim, stats.MeleeHaste, haste)
					return
				}
				character.MultiplyMeleeSpeed(sim, multiplier)
				if ranged {
					character.MultiplyRangedSpeed(sim, multiplier)
				}
			},
			OnExpire: func(_ *core.Aura, sim *core.Simulation) {
				if !character.Env.IsForever() {
					character.AddStatDynamic(sim, stats.MeleeHaste, -haste)
					return
				}
				character.MultiplyMeleeSpeed(sim, 1/multiplier)
				if ranged {
					character.MultiplyRangedSpeed(sim, 1/multiplier)
				}
			},
		})
		character.RegisterOnItemSwap(func(sim *core.Simulation) {
			aura.Deactivate(sim)
			if count() > 0 {
				aura.Activate(sim)
			}
		})
	})
}
