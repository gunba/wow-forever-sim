package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const ShadowburnRanks = 6

func (warlock *Warlock) registerShadowBurnBaseConfig(rank int, timer *core.Timer) core.SpellConfig {
	// Beta client 1.60.1 values for every rank. The BlizzCon tooltip's 102 to 111 for rank 1 is not
	// what the client carries.
	spellId := [ShadowburnRanks + 1]int32{0, 17877, 18867, 18868, 18869, 18870, 18871}[rank]
	baseDamage := [ShadowburnRanks + 1][]float64{{0}, {65, 74}, {81, 91}, {119, 133}, {147, 164}, {201, 224}, {259, 288}}[rank]
	manaCost := [ShadowburnRanks + 1]float64{0, 105, 130, 190, 245, 305, 365}[rank]
	level := [ShadowburnRanks + 1]int{0, 15, 24, 32, 40, 48, 56}[rank]

	spellCoeff := 0.429

	return core.SpellConfig{
		ActionID:      core.ActionID{SpellID: spellId},
		SpellCode:     SpellCode_WarlockShadowburn,
		SpellSchool:   core.SpellSchoolShadow,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | core.SpellFlagResetAttackSwing | core.SpellFlagBinary | WarlockFlagDestruction,
		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    timer,
				Duration: time.Second * time.Duration(15),
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamage[0], baseDamage[1])
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
		},
	}
}

func (warlock *Warlock) registerShadowBurnSpell() {
	if !warlock.Talents.Shadowburn {
		return
	}

	warlock.Shadowburn = make([]*core.Spell, 0)
	timer := warlock.NewTimer()
	for rank := 1; rank <= ShadowburnRanks; rank++ {
		config := warlock.registerShadowBurnBaseConfig(rank, timer)

		if config.RequiredLevel <= int(warlock.Level) {
			warlock.Shadowburn = append(warlock.Shadowburn, warlock.GetOrRegisterSpell(config))
		}
	}
}
