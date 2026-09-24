package shaman

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const FlameShockRanks = 6

// Forever beta client values. The client stores the dot as damage per tick (7, 8, 14, 21, 34, 44) and the table
// holds the four ticks' total; its 0.1 coefficient is per tick, which is how the dot config applies it.
var FlameShockSpellId = [FlameShockRanks + 1]int32{0, 8050, 8052, 8053, 10447, 10448, 29228}
var FlameShockBaseDamage = [FlameShockRanks + 1]float64{0, 24, 38, 49, 89, 137, 166}
var FlameShockBaseDotDamage = [FlameShockRanks + 1]float64{0, 28, 32, 56, 84, 136, 176}
var FlameShockBaseSpellCoef = [FlameShockRanks + 1]float64{0, .214, .214, .214, .214, .214, .214}
var FlameShockDotSpellCoef = [FlameShockRanks + 1]float64{0, .1, .1, .1, .1, .1, .1}
var FlameShockManaCost = [FlameShockRanks + 1]float64{0, 55, 95, 160, 250, 345, 410}
var FlameShockLevel = [FlameShockRanks + 1]int{0, 10, 18, 28, 40, 52, 60}

func (shaman *Shaman) registerFlameShockSpell(shockTimer *core.Timer) {
	shaman.FlameShock = make([]*core.Spell, FlameShockRanks+1)

	for rank := 1; rank <= FlameShockRanks; rank++ {
		if FlameShockLevel[rank] <= int(shaman.Level) {
			shaman.FlameShock[rank] = shaman.RegisterSpell(shaman.newFlameShockSpell(rank, shockTimer))
		}
	}
}

func (shaman *Shaman) newFlameShockSpell(rank int, shockTimer *core.Timer) core.SpellConfig {
	numTicks := 4
	tickDuration := time.Second * 3

	spellId := FlameShockSpellId[rank]
	baseDamage := FlameShockBaseDamage[rank]
	baseDotDamage := FlameShockBaseDotDamage[rank] / float64(numTicks)
	baseSpellCoeff := FlameShockBaseSpellCoef[rank]
	dotSpellCoeff := FlameShockDotSpellCoef[rank]
	manaCost := FlameShockManaCost[rank]
	level := FlameShockLevel[rank]

	spell := shaman.newShockSpellConfig(
		core.ActionID{SpellID: spellId},
		core.SpellSchoolFire,
		manaCost,
		shockTimer,
	)

	spell.SpellCode = SpellCode_ShamanFlameShock
	spell.RequiredLevel = level
	spell.Rank = rank
	// Call of Flame names Flame Shock alongside the fire totems under Forever.
	spell.DamageMultiplier *= shaman.callOfFlameMultiplier()

	spell.BonusCoefficient = baseSpellCoeff

	spell.Dot = core.DotConfig{
		Aura: core.Aura{
			Label: fmt.Sprintf("Flame Shock (Rank %d)", rank),
		},

		NumberOfTicks:    int32(numTicks),
		TickLength:       tickDuration,
		BonusCoefficient: dotSpellCoeff,

		OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
			dot.Snapshot(target, baseDotDamage, isRollover)
		},

		OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
			dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
		},
	}

	spell.ApplyEffects = func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
		result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
		if result.Landed() {
			spell.Dot(result.Target).Apply(sim)
		}
	}

	return spell
}
