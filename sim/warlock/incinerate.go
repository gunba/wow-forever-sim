package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (warlock *Warlock) registerIncinerateSpell() {
	if !warlock.Talents.Incinerate {
		return
	}

	// Beta client 1.60.1: three ranks, learned at 40 (412758), 50 (1293812) and 60 (1293813), 2.5 sec
	// cast, 0.714 coefficient and 25% more damage on a target with Immolate. The highest rank the
	// character's level allows is the one registered.
	if warlock.Level < 40 {
		return
	}

	spellID := map[int32]int32{40: 412758, 50: 1293812, 60: 1293813}[warlock.Level]
	rank := map[int32]int{40: 1, 50: 2, 60: 3}[warlock.Level]
	baseDamage := map[int32][]float64{40: {100, 114}, 50: {146, 168}, 60: {201, 233}}[warlock.Level]
	manaCost := map[int32]float64{40: 205, 50: 265, 60: 325}[warlock.Level]
	spellCoeff := 0.714
	castTime := time.Millisecond * 2500

	warlock.Incinerate = warlock.RegisterSpell(core.SpellConfig{
		SpellCode:     SpellCode_WarlockIncinerate,
		ActionID:      core.ActionID{SpellID: spellID},
		SpellSchool:   core.SpellSchoolFire,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | core.SpellFlagResetAttackSwing | WarlockFlagDestruction,
		RequiredLevel: int(warlock.Level),
		Rank:          rank,
		MissileSpeed:  20,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: castTime,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := sim.Roll(baseDamage[0], baseDamage[1])
			if warlock.getActiveImmolateSpell(target) != nil {
				damage *= 1.25
			}

			result := spell.CalcDamage(sim, target, damage, spell.OutcomeMagicHitAndCrit)
			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
			})
		},
	})
}
