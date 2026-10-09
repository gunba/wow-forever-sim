package priest

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const SmiteRanks = 8

var SmiteSpellId = [SmiteRanks + 1]int32{0, 585, 591, 598, 984, 1004, 6060, 10933, 10934}

// Forever beta client 1.60.1.69893: ranks 3 and up hit for less, and there is no downranking penalty.
var SmiteBaseDamage = [SmiteRanks + 1][]float64{{0}, {15, 20}, {28, 34}, {47, 53}, {60, 68}, {81, 91}, {93, 106}, {124, 139}, {166, 187}}
var SmiteSpellCoef = [SmiteRanks + 1]float64{0, 0.429, 0.571, 0.714, 0.714, 0.714, 0.714, 0.714, 0.714}
var SmiteCastTime = [SmiteRanks + 1]int{0, 1500, 2000, 2500, 2500, 2500, 2500, 2500, 2500}
var SmiteManaCost = [SmiteRanks + 1]float64{0, 20, 30, 60, 95, 140, 185, 230, 280}
var SmiteLevel = [SmiteRanks + 1]int{0, 1, 6, 14, 22, 30, 38, 46, 54}

// SpellEffect/SpellLevels 1.60.1.70291; changed low ranks only.
// https://us.forums.blizzard.com/en/wow/t/2360696/5
var foreverSmiteRankDamage = [...]struct {
	basePoints, variance, pointsPerLevel float64
	spellLevel, maxLevel                 int
}{
	{},
	{15, 0.26666668057, 0.10000000149, 1, 6},
	{27, 0.21428571641, 0.30000001192, 6, 11},
	{45, 0.13793103397, 0.5, 14, 19},
	{60, 0.14285714924, 0.69999998808, 22, 27},
}

func (priest *Priest) registerSmiteSpell() {
	priest.Smite = make([]*core.Spell, SmiteRanks+1)

	for rank := 1; rank <= SmiteRanks; rank++ {
		config := priest.getSmiteBaseConfig(rank)

		if config.RequiredLevel <= int(priest.Level) {
			priest.Smite[rank] = priest.GetOrRegisterSpell(config)
		}
	}
}

func (priest *Priest) getSmiteBaseConfig(rank int) core.SpellConfig {
	spellId := SmiteSpellId[rank]
	baseDamageLow := SmiteBaseDamage[rank][0]
	baseDamageHigh := SmiteBaseDamage[rank][1]
	if priest.Env.IsForever() && rank < len(foreverSmiteRankDamage) {
		damage := foreverSmiteRankDamage[rank]
		growth := float64(max(0, min(int(priest.Level), damage.maxLevel)-damage.spellLevel)) * damage.pointsPerLevel
		mean, spread := damage.basePoints+growth, damage.basePoints*damage.variance/2
		baseDamageLow, baseDamageHigh = mean-spread, mean+spread
	}
	spellCoeff := SmiteSpellCoef[rank]
	castTime := SmiteCastTime[rank]
	manaCost := SmiteManaCost[rank]
	level := SmiteLevel[rank]

	return core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellId},
		SpellCode:   SpellCode_PriestSmite,
		SpellSchool: core.SpellSchoolHoly,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagPriest | core.SpellFlagAPL,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*time.Duration(castTime) - time.Millisecond*100*time.Duration(priest.Talents.DivineFury),
			},
		},

		BonusCoefficient: spellCoeff,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageLow, baseDamageHigh)
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
		},
	}
}
