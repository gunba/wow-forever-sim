package rogue

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Improved Expose Armor no longer scales the armor reduction, it discounts the finisher
// and hands combo points back on a full spend. The beta client raised the spell's own armor
// per combo point instead (rank 5 340 -> 450, a little under Classic's talented 510), so no
// talent multiplier is applied.
func (rogue *Rogue) registerExposeArmorSpell() {
	rogue.ExposeArmorAuras = rogue.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.ExposeArmorAura(target, 0)
	})

	spellID := map[int32]int32{
		25: 8647,
		40: 8650,
		50: 11197,
		60: 11198,
	}[rogue.Level]

	arpenPerCombo := map[int32]float64{
		25: 90,
		40: 270,
		50: 360,
		60: 450,
	}[rogue.Level]

	// Improved Expose Armor takes 5 Energy off and hands 1 combo point back per rank, on a 5
	// point spend.
	energyCost := 25.0 - 5*float64(rogue.Talents.ImprovedExposeArmor)
	cpMetrics := rogue.NewComboPointMetrics(core.ActionID{SpellID: 14169})

	// share ExtraCastCondition() state with ApplyEffects()
	var arpen float64
	var eaAura *core.Aura

	rogue.ExposeArmor = rogue.RegisterSpell(core.SpellConfig{
		SpellCode:    SpellCode_RogueExposeArmor,
		ActionID:     core.ActionID{SpellID: spellID},
		SpellSchool:  core.SpellSchoolPhysical,
		DefenseType:  core.DefenseTypeMelee,
		ProcMask:     core.ProcMaskMeleeMHSpecial,
		Flags:        rogue.finisherFlags(),
		MetricSplits: 6,

		EnergyCost: core.EnergyCostOptions{
			Cost:   energyCost,
			Refund: 0,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				spell.SetMetricsSplit(spell.Unit.ComboPoints())
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			if rogue.ComboPoints() == 0 {
				return false
			}

			eaAura = rogue.ExposeArmorAuras.Get(target)
			arpen = float64(rogue.ComboPoints()) * arpenPerCombo

			if curActive := eaAura.ExclusiveEffects[0].Category.GetActiveEffect(); curActive != nil {
				return arpen >= curActive.Priority
			}
			return true
		},

		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			rogue.BreakStealth(sim)

			comboPoints := rogue.ComboPoints()

			result := spell.CalcOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
			if result.Landed() {
				eaAura.ExclusiveEffects[0].SetPriority(sim, arpen)
				eaAura.Activate(sim)
				rogue.SpendComboPoints(sim, spell)
				if rogue.Talents.ImprovedExposeArmor > 0 && comboPoints == 5 {
					rogue.AddComboPoints(sim, rogue.Talents.ImprovedExposeArmor, target, cpMetrics)
				}
			} else {
				spell.IssueRefund(sim)
			}
			spell.DealOutcome(sim, result)
		},

		RelatedAuras: []core.AuraArray{rogue.ExposeArmorAuras},
	})
	rogue.Finishers = append(rogue.Finishers, rogue.ExposeArmor)
}
