package warrior

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

func (warrior *Warrior) registerExecuteSpell() {

	rank := warrior.rank(ranks.Execute)
	if rank == nil {
		return
	}
	flatDamage := rank.Find(ranks.EffectDummy, 0).Low
	convertedRageDamage := ranks.ExecuteDamagePerRage(rank)
	spellID := int32(20662)

	var rageMetrics *core.ResourceMetrics
	warrior.Execute = warrior.RegisterSpell(BattleStance|BerserkerStance, core.SpellConfig{
		SpellCode:   SpellCode_WarriorExecute,
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | core.SpellFlagPassiveSpell | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			// Rank 2 takes 5 off, not the 6 that doubling rank 1 gives. Both the tree and
			// wowforevertalents read it this way; no beta tooltip past rank 1 has been seen.
			Cost:   15 - []float64{0, 3, 5}[warrior.Talents.ImprovedExecute],
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			if ExecuteRangeHook != nil {
				return ExecuteRangeHook(sim, target)
			}
			return sim.IsExecutePhase20()
		},

		CritDamageBonus: warrior.impale(),

		DamageMultiplier: 1,
		ThreatMultiplier: 1.25,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			extraRage := spell.Unit.CurrentRage()
			warrior.SpendRage(sim, extraRage, rageMetrics)
			// We must count this rage event if the spell itself cost 0,
			// otherwise we could end up with 0 events even though rage was spent.
			if spell.Cost.GetCurrentCost() > 0 {
				rageMetrics.Events--
			}

			baseDamage := flatDamage + convertedRageDamage*(extraRage)

			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)

			if !result.Landed() {
				spell.IssueRefund(sim)
			}
		},
	})
	rageMetrics = warrior.Execute.Cost.SpellCostFunctions.(*core.RageCost).ResourceMetrics
}
