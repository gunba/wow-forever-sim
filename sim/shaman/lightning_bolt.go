package shaman

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const LightningBoltRanks = 10

// Damage, coefficients, cast times and costs are the Forever beta client's (build 1.60.1.70009). Damage is the
// client's range scaled to level 60 by its per-level points, capped at the rank's max level, low end rounded
// down and high end up. Forever halves the upper ranks, drops the cast to 2.5 sec and gives every rank from 3 up
// the full 0.714 coefficient.
var LightningBoltSpellId = [LightningBoltRanks + 1]int32{0, 403, 529, 548, 915, 943, 6041, 10391, 10392, 15207, 15208}
var LightningBoltBaseDamage = [LightningBoltRanks + 1][]float64{{0}, {15, 17}, {29, 34}, {44, 52}, {55, 63}, {70, 81}, {108, 122}, {142, 160}, {157, 177}, {172, 194}, {190, 212}}
var LightningBoltSpellCoef = [LightningBoltRanks + 1]float64{0, .429, .571, .714, .714, .714, .714, .714, .714, .714, .714}
var LightningBoltCastTime = [LightningBoltRanks + 1]int32{0, 1500, 2000, 2500, 2500, 2500, 2500, 2500, 2500, 2500, 2500}
var LightningBoltManaCost = [LightningBoltRanks + 1]float64{0, 15, 30, 45, 60, 85, 110, 135, 160, 190, 220}
var LightningBoltLevel = [LightningBoltRanks + 1]int{0, 1, 8, 14, 20, 26, 32, 38, 44, 50, 56}

func (shaman *Shaman) registerLightningBoltSpell() {
	shaman.LightningBolt = make([]*core.Spell, LightningBoltRanks+1)
	shaman.LightningBoltOverload = make([]*core.Spell, LightningBoltRanks+1)

	for rank := 1; rank <= LightningBoltRanks; rank++ {
		config := shaman.newLightningBoltSpellConfig(rank)

		if config.RequiredLevel <= int(shaman.Level) {
			// The overload gets a config of its own so that the two casts share no state.
			shaman.LightningBoltOverload[rank] = shaman.registerOverloadSpell(shaman.newLightningBoltSpellConfig(rank))
			shaman.LightningBolt[rank] = shaman.RegisterSpell(config)
		}
	}
}

func (shaman *Shaman) newLightningBoltSpellConfig(rank int) core.SpellConfig {
	spellId := LightningBoltSpellId[rank]
	baseDamageLow := LightningBoltBaseDamage[rank][0]
	baseDamageHigh := LightningBoltBaseDamage[rank][1]
	spellCoeff := LightningBoltSpellCoef[rank]
	castTime := LightningBoltCastTime[rank]
	manaCost := LightningBoltManaCost[rank]
	level := LightningBoltLevel[rank]

	spell := shaman.newElectricSpellConfig(
		core.ActionID{SpellID: spellId},
		manaCost,
		time.Millisecond*time.Duration(castTime),
	)
	spell.SpellCode = SpellCode_ShamanLightningBolt
	spell.MissileSpeed = 20
	spell.RequiredLevel = level
	spell.Rank = rank
	spell.BonusCoefficient = spellCoeff

	spell.ApplyEffects = func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
		baseDamage := sim.Roll(baseDamageLow, baseDamageHigh)
		result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

		spell.WaitTravelTime(sim, func(sim *core.Simulation) {
			spell.DealDamage(sim, result)
		})

		shaman.tryLightningOverload(sim, target, spell, shaman.LightningBoltOverload[rank])
	}

	return spell
}
