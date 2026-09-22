package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const DeathCoilRanks = 3

func (warlock *Warlock) getDeathCoilBaseConfig(rank int, timer *core.Timer) core.SpellConfig {
	spellId := [DeathCoilRanks + 1]int32{0, 6789, 17925, 17926}[rank]
	// Beta client 1.60.1 values: slightly less damage, slightly more mana
	baseDamage := [DeathCoilRanks + 1]float64{0, 285, 375, 460}[rank]
	manaCost := [DeathCoilRanks + 1]float64{0, 435, 525, 600}[rank]
	level := [DeathCoilRanks + 1]int{0, 42, 50, 58}[rank]
	spellCoeff := 0.214

	healingSpell := warlock.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellId}.WithTag(1),
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskSpellHealing,
		Flags:       core.SpellFlagPassiveSpell | core.SpellFlagHelpful,

		DamageMultiplier: 1,
		ThreatMultiplier: 0,
	})

	return core.SpellConfig{
		SpellCode:     SpellCode_WarlockDeathCoil,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolShadow,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | core.SpellFlagResetAttackSwing | core.SpellFlagBinary | WarlockFlagAffliction,
		RequiredLevel: level,
		Rank:          rank,
		MissileSpeed:  24,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    timer,
				Duration: time.Minute * 2,
			},
		},

		DamageMultiplierAdditive: 1,
		DamageMultiplier:         1,
		ThreatMultiplier:         1,
		BonusCoefficient:         spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			results := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			spell.WaitTravelTime(sim, func(s *core.Simulation) {
				spell.DealDamage(sim, results)
				if results.Landed() {
					healingSpell.CalcAndDealHealing(sim, healingSpell.Unit, results.Damage, healingSpell.OutcomeHealing)
				}
			})
		},
	}
}

func (warlock *Warlock) registerDeathCoilSpell() {
	warlock.DeathCoil = make([]*core.Spell, 0)
	timer := warlock.NewTimer()
	for rank := 1; rank <= DeathCoilRanks; rank++ {
		config := warlock.getDeathCoilBaseConfig(rank, timer)

		if config.RequiredLevel <= int(warlock.Level) {
			warlock.DeathCoil = append(warlock.DeathCoil, warlock.GetOrRegisterSpell(config))
		}
	}
}
