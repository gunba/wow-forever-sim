package hunter

import (
	"math"
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (hp *HunterPet) newDemoralizingScreech() *core.Spell {
	// Active ranks, not the Hunter's teaching records 24580/81/82.
	ranks := []struct {
		level, maxLevel               int32
		low, high, baseAP, apPerLevel float64
	}{
		{8, 24, 7, 9, 41, 1.4},
		{24, 48, 9, 13, 77, 1.4},
		{48, 56, 21, 27, 153, 1.4},
		{56, 60, 24, 42, 198, 1.6},
	}
	selected := -1
	for i, rank := range ranks {
		if hp.Level >= rank.level {
			selected = i
		}
	}
	if selected < 0 {
		return nil
	}
	rank := ranks[selected]
	spellID := []int32{24423, 24577, 24578, 24579}[selected]
	// Client per-level coefficient, capped by rank. Nearest-integer rounding
	// agrees with rank-max tooltips; lower-level server rounding remains open.
	reduction := math.Round(rank.baseAP + rank.apPerLevel*float64(min(hp.Level, rank.maxLevel)-rank.level))
	auras := hp.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.DemoralizingScreechAura(target, core.ActionID{SpellID: spellID}, reduction)
	})
	return hp.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{SpellID: spellID}, SpellCode: SpellCode_HunterPetScreech, RequiredLevel: int(rank.level), Rank: selected + 1,
		SpellSchool: core.SpellSchoolPhysical, DefenseType: core.DefenseTypeMelee, ProcMask: core.ProcMaskMeleeSpecial, Flags: core.SpellFlagMeleeMetrics,
		FocusCost:        core.FocusCostOptions{Cost: 20},
		Cast:             core.CastConfig{DefaultCast: core.Cast{GCD: PetGCD}, IgnoreHaste: true, CD: core.Cooldown{Timer: hp.NewTimer(), Duration: 10 * time.Second}},
		DamageMultiplier: 1, ThreatMultiplier: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			// Data: damage is single-target, the AP effect is caster-area (five yards).
			result := spell.CalcAndDealDamage(sim, target, sim.Roll(rank.low, rank.high), spell.OutcomeMeleeSpecialHitAndCrit)
			// Stationary enemies are represented as a melee cluster. Linked-area
			// avoidance remains unverified; use the landed primary application.
			if result.Landed() {
				for _, enemy := range sim.Encounter.TargetUnits {
					auras.Get(enemy).Activate(sim)
				}
			}
		},
	})
}
