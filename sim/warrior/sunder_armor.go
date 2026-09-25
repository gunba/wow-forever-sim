package warrior

import (
	"github.com/wowsims/classic/sim/core"
)

func (warrior *Warrior) registerSunderArmorSpell() {
	warrior.SunderArmorAuras = warrior.NewEnemyAuraArray(core.SunderArmorAura)

	spellID := int32(11597)

	// Client 1.60.1.70009 gives rank 5 a flat 206 threat (was 1013).
	// The newly noted extra Attack Power contribution is server-side and its
	// coefficient is not exposed in SpellEffect; do not fabricate one here.
	threat := 2.25 * 2 * float64(warrior.Level)
	if warrior.Env.IsForever() {
		threat = 206
	}

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
