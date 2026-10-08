package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

var summonHawkBaseDamage = [5]float64{0, 32, 47, 85, 108}

func (hunter *Hunter) summonHawkRank() int {
	switch {
	case hunter.Level >= 60:
		return 4
	case hunter.Level >= 48:
		return 3
	case hunter.Level >= 36:
		return 2
	default:
		return 1
	}
}

func (hunter *Hunter) registerSummonHawkSpell(timer *core.Timer) {
	if !hunter.Talents.SummonHawk {
		return
	}

	rank := hunter.summonHawkRank()
	spellId := [5]int32{0, 1293241, 1293525, 1293526, 1293527}[rank]
	baseDamage := summonHawkBaseDamage[rank]
	manaCost := [5]float64{0, 80, 105, 135, 190}[rank]

	hunter.SummonHawk = hunter.RegisterSpell(core.SpellConfig{
		SpellCode:     SpellCode_HunterSummonHawk,
		ActionID:      core.ActionID{SpellID: spellId},
		Rank:          rank,
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeRanged,
		CastType:      proto.CastType_CastTypeRanged,
		ProcMask:      core.ProcMaskEmpty,
		Flags:         core.SpellFlagAPL,
		MissileSpeed:  35,
		MinTravelTime: time.Second,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true, // Hunter GCD is locked at 1.5s.
			CD: core.Cooldown{
				Timer:    timer,
				Duration: time.Second * 6,
			},
		},

		BonusCritRating:  2 * float64(hunter.Talents.Ferocity) * core.CritRatingPerCritChance,
		DamageMultiplier: 1 + .03*float64(hunter.Talents.UnleashedFury),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			// Client 70245's ALWAYS_HIT permits crits, but no failed owner hit
			// roll can suppress the opening or its independent summon effect.
			damage := baseDamage + .05*spell.RangedAttackPower(target, false)
			result := spell.CalcDamage(sim, target, damage, spell.OutcomeMeleeSpecialCritOnly)
			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
				hunter.summonHawkGuardian(sim, target)
			})
		},
	})
}
