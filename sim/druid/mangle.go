package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Berserk widens Mangle to a cleave.
const MangleBerserkTargets = 3

// Beta client 1.60.1.70009: Primal Bite (formerly Mangle) is 20 Rage, a 6 sec cooldown, and 100% weapon damage plus a bonus that grows by
// rank (407995, 1238069, 1238070, 1238073 at levels 25, 36, 48, 60). The client does not carry threat, so the 1.5x is
// still Season of Discovery's.
func (druid *Druid) registerMangleBearSpell() {
	if !druid.Talents.Mangle {
		return
	}

	flatDamageBonus := map[int32]float64{25: 26, 40: 38, 50: 59, 60: 77}[druid.Level]
	spellID := map[int32]int32{25: 407995, 40: 1238069, 50: 1238070, 60: 1238073}[druid.Level]
	results := make([]*core.SpellResult, min(MangleBerserkTargets, druid.Env.GetNumTargets()))

	druid.MangleBear = druid.RegisterSpell(Bear, core.SpellConfig{
		SpellCode:   SpellCode_DruidMangle,
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		RageCost: core.RageCostOptions{
			Cost:   20 - float64(druid.Talents.Ferocity),
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    druid.NewTimer(),
				Duration: time.Second * 6,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1.5,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			numHits := 1
			if druid.BerserkAura.IsActive() {
				numHits = len(results)
			}

			for idx := 0; idx < numHits; idx++ {
				baseDamage := flatDamageBonus + spell.Unit.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
				results[idx] = spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)
				target = sim.Environment.NextTargetUnit(target)
			}

			for idx := 0; idx < numHits; idx++ {
				spell.DealDamage(sim, results[idx])
			}

			if !results[0].Landed() {
				spell.IssueRefund(sim)
			}

			if druid.BerserkAura.IsActive() {
				spell.CD.Reset()
			}
		},
	})
}
