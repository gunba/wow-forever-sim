package warrior

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

func (warrior *Warrior) registerSunderArmorSpell() {
	rank := warrior.rank(ranks.SunderArmor)
	if rank == nil {
		return
	}
	armorPerStack := -rank.Find(ranks.EffectApplyAura, ranks.AuraModArmorPercentOrNeg).Low
	if rank.Rank >= 5 {
		warrior.SunderArmorAuras = warrior.NewEnemyAuraArray(core.SunderArmorAura)
	} else {
		warrior.SunderArmorAuras = warrior.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return core.SunderArmorAuraWithReduction(target, armorPerStack)
		})
	}

	spellID := int32(11597)

	// Client 70338 gives rank 5 a flat 206 threat (effect 63; was 1013). The deep dive says
	// Sunder's "Threat generated now scales with Attack Power"; the coefficient is server-side.
	// A published 'Warrior Threat Values' test sheet solves 0.0433 AP (rank 1) and 0.0476 (rank 2) on the client's
	// flat values (34/75); 0.05 is the provisional level-60 value.
	threat := 2.25 * 2 * float64(warrior.Level)
	if warrior.Env.IsForever() {
		threat = rank.Find(ranks.EffectThreat, 0).Low
	}
	threatAP := ranks.SunderThreatAP(rank)

	var canApplySunder bool

	warrior.SunderArmor = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost:   15 - float64(warrior.Talents.ImprovedSunderArmor),
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			sa := warrior.SunderArmorAuras.Get(target)
			if sa.IsActive() {
				canApplySunder = true
			} else if sa.ExclusiveEffects[0].Category.AnyActive() {
				canApplySunder = false
			} else {
				canApplySunder = true
			}
			return canApplySunder
		},

		ThreatMultiplier: 1,
		FlatThreatBonus:  threat,

		RelatedAuras: []core.AuraArray{warrior.SunderArmorAuras},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			if sim.IsForever() {
				spell.FlatThreatBonus = threat + threatAP*spell.MeleeAttackPower(target)
			}
			result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeWeaponSpecialNoCrit) // Cannot be blocked
			if !result.Landed() {
				spell.IssueRefund(sim)
				return
			}

			if canApplySunder {
				sa := warrior.SunderArmorAuras.Get(target)
				sa.Activate(sim)
				sa.AddStack(sim)
			}
		},
	})
}
