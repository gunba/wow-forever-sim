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

// SpellEffect/SpellLevels 1.60.1.70245: own base, variance, level growth and SP
// coefficient, without integer rounding or a second 50% damage multiplier.
var lightningBoltOverloadDamage = [LightningBoltRanks + 1]electricOverloadDamage{
	{},
	{408439, 7, .14285700023, .20000000298, .21449999511, 1, 6},
	{408440, 14.5, .14285714924, .25, .28549998999, 8, 13},
	{408441, 22.5, .16326500475, .30000001192, .35699999332, 14, 19},
	{408442, 28, .13483099639, .30000001192, .35699999332, 20, 25},
	{408443, 36, .1343279928, .34999999404, .35699999332, 26, 31},
	{408472, 56, .12087912112, .40000000596, .35699999332, 32, 37},
	{408473, 74, .125, .40000000596, .35699999332, 38, 43},
	{408474, 81, .11409395933, .5, .35699999332, 44, 49},
	{408475, 89, .11956521869, .5, .35699999332, 50, 55},
	{408477, 98, .11312217265, .60000002384, .35699999332, 56, 61},
}

// Build 1.60.1.70291 parent ranks1–5; overload children keep their own,
// unchanged rows. No rounded endpoints or additional half-damage multiplier.
// https://us.forums.blizzard.com/en/wow/t/2360696/5
var foreverLightningBoltRankDamage = [...]electricOverloadDamage{
	{},
	{403, 15, 0.14285714924, 0.10000000149, 0.42899999022, 1, 6},
	{529, 27, 0.14285714924, 0.10000000149, 0.57099997997, 8, 13},
	{548, 45, 0.16326530278, 0.30000001192, 0.71399998665, 14, 19},
	{915, 61, 0.13483145833, 0.40000000596, 0.71399998665, 20, 25},
	{943, 82, 0.13432836533, 0.5, 0.71399998665, 26, 31},
}

func (shaman *Shaman) registerLightningBoltSpell() {
	shaman.LightningBolt = make([]*core.Spell, LightningBoltRanks+1)
	shaman.LightningBoltOverload = make([]*core.Spell, LightningBoltRanks+1)

	for rank := 1; rank <= LightningBoltRanks; rank++ {
		config := shaman.newLightningBoltSpellConfig(rank, false)

		if config.RequiredLevel <= int(shaman.Level) {
			// The overload gets a config of its own so that the two casts share no state.
			shaman.LightningBoltOverload[rank] = shaman.registerOverloadSpell(shaman.newLightningBoltSpellConfig(rank, true))
			shaman.LightningBolt[rank] = shaman.RegisterSpell(config)
		}
	}
}

func (shaman *Shaman) newLightningBoltSpellConfig(rank int, isOverload bool) core.SpellConfig {
	spellId := LightningBoltSpellId[rank]
	baseDamageLow := LightningBoltBaseDamage[rank][0]
	baseDamageHigh := LightningBoltBaseDamage[rank][1]
	spellCoeff := LightningBoltSpellCoef[rank]
	castTime := LightningBoltCastTime[rank]
	manaCost := LightningBoltManaCost[rank]
	level := LightningBoltLevel[rank]
	if !isOverload && shaman.Env.IsForever() && rank < len(foreverLightningBoltRankDamage) {
		baseDamageLow, baseDamageHigh = foreverLightningBoltRankDamage[rank].rangeAtLevel(shaman.Level)
	}
	if isOverload && shaman.Env.IsForever() {
		damage := lightningBoltOverloadDamage[rank]
		baseDamageLow, baseDamageHigh = damage.rangeAtLevel(shaman.Level)
		spellCoeff = damage.coefficient
	}

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
