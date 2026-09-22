package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const SoulFireRanks = 2
const SoulFireCastTime = time.Millisecond * 6000

func (warlock *Warlock) getSoulFireBaseConfig(rank int, timer *core.Timer) core.SpellConfig {
	spellId := [SoulFireRanks + 1]int32{0, 6353, 17924}[rank]
	// Beta client 1.60.1 values
	baseDamage := [SoulFireRanks + 1][]float64{{0, 0}, {344, 430}, {390, 487}}[rank]
	manaCost := [SoulFireRanks + 1]float64{0, 305, 335}[rank]
	level := [SoulFireRanks + 1]int{0, 48, 56}[rank]
	spellCoeff := 1.0

	config := core.SpellConfig{
		SpellCode:     SpellCode_WarlockSoulFire,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolFire,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | core.SpellFlagResetAttackSwing | WarlockFlagDestruction,
		RequiredLevel: level,
		Rank:          rank,
		MissileSpeed:  24,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: SoulFireCastTime,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := sim.Roll(baseDamage[0], baseDamage[1])
			results := spell.CalcDamage(sim, target, damage, spell.OutcomeMagicHitAndCrit)
			spell.WaitTravelTime(sim, func(s *core.Simulation) {
				spell.DealDamage(sim, results)
			})
		},
	}

	// Decimation: 45% per point in the beta client, so 2/2 leaves Soul Fire a six second cooldown
	cooldownReduction := 0.45 * float64(warlock.Talents.Decimation)

	config.Cast.CD = core.Cooldown{
		Timer:    timer,
		Duration: time.Duration(float64(time.Minute) * (1 - cooldownReduction)),
	}

	return config
}

func (warlock *Warlock) registerSoulFireSpell() {
	warlock.SoulFire = make([]*core.Spell, 0)
	timer := warlock.NewTimer()
	for rank := 1; rank <= SoulFireRanks; rank++ {
		config := warlock.getSoulFireBaseConfig(rank, timer)

		if config.RequiredLevel <= int(warlock.Level) {
			warlock.SoulFire = append(warlock.SoulFire, warlock.GetOrRegisterSpell(config))
		}
	}
}
