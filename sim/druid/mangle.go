package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Berserk widens Mangle to a cleave.
const MangleBerserkTargets = 3

// Client 70170: Primal Bite is 20 Rage, a 6 sec cooldown, and 100%
// weapon damage plus 26/38/59/77 at levels 25/36/48/60. The October 1
// notes roughly double threat; this applies a 2x relative change to the
// inherited, still-unverified 1.5x coefficient, not a measured absolute value.
func primalBiteRank(level int32) int {
	rank := 0
	for candidate, learned := range []int32{25, 36, 48, 60} {
		if level >= learned {
			rank = candidate + 1
		}
	}
	return rank
}

func (druid *Druid) registerMangleBearSpell() {
	if !druid.Talents.Mangle || druid.Level < 25 {
		return
	}

	rank := primalBiteRank(druid.Level)
	flatDamageBonus := [...]float64{0, 26, 38, 59, 77}[rank]
	spellID := [...]int32{0, 407995, 1238069, 1238070, 1238073}[rank]
	threatMultiplier := 1.5
	if druid.Env.IsForever() {
		threatMultiplier *= 2
	}
	results := make([]*core.SpellResult, min(MangleBerserkTargets, druid.Env.GetNumTargets()))

	druid.MangleBear = druid.RegisterSpell(Bear, core.SpellConfig{
		SpellCode:     SpellCode_DruidMangle,
		ActionID:      core.ActionID{SpellID: spellID},
		Rank:          rank,
		RequiredLevel: int([]int32{25, 36, 48, 60}[rank-1]),
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeMHSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

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
		ThreatMultiplier: threatMultiplier,
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
