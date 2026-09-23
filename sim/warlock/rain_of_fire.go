package warlock

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const RainOfFireRanks = 4

func (warlock *Warlock) getRainOfFireBaseConfig(rank int) core.SpellConfig {
	spellId := [RainOfFireRanks + 1]int32{0, 5740, 6219, 11677, 11678}[rank]
	// Beta client 1.60.1: each tick is now its own damage spell (1282380, 1282383, 1282384, 1282385)
	// carrying the per tick damage and a 0.083 coefficient. The 0.03 on the channel's dummy effect is
	// not the damage coefficient.
	spellCoeff := [RainOfFireRanks + 1]float64{0, 0.083, 0.083, 0.083, 0.083}[rank]
	baseDamage := [RainOfFireRanks + 1]float64{0, 40, 91, 149, 220}[rank]
	manaCost := [RainOfFireRanks + 1]float64{0, 295, 605, 885, 1185}[rank]
	level := [RainOfFireRanks + 1]int{0, 20, 34, 46, 58}[rank]

	flags := core.SpellFlagAPL | core.SpellFlagResetAttackSwing | WarlockFlagDestruction | core.SpellFlagChanneled

	config := core.SpellConfig{
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolFire,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         flags,
		RequiredLevel: level,
		Rank:          rank,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		Dot: core.DotConfig{
			IsAOE: true,
			Aura: core.Aura{
				Label: "RainOfFire-" + warlock.Label + strconv.Itoa(rank),
			},
			NumberOfTicks:       4,
			TickLength:          time.Second * 2,
			AffectedByCastSpeed: true,
			BonusCoefficient:    spellCoeff,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, baseDamage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, aoeTarget, dot.OutcomeTick)
				}

			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.AOEDot().Apply(sim)
		},
	}

	return config
}

func (warlock *Warlock) registerRainOfFireSpell() {
	warlock.RainOfFire = make([]*core.Spell, 0)
	for rank := 1; rank <= RainOfFireRanks; rank++ {
		config := warlock.getRainOfFireBaseConfig(rank)

		if config.RequiredLevel <= int(warlock.Level) {
			warlock.RainOfFire = append(warlock.RainOfFire, warlock.GetOrRegisterSpell(config))
		}
	}
}
