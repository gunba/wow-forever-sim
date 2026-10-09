package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const ShadowBoltRanks = 10

// SpellEffect/SpellLevels 1.60.1.70291; changed low ranks only.
// https://us.forums.blizzard.com/en/wow/t/2360696/5
var foreverShadowBoltRankDamage = [...]struct {
	basePoints, variance, pointsPerLevel float64
	spellLevel, maxLevel                 int
}{
	{},
	{13, 0.28571429849, 0.20000000298, 1, 5},
	{23, 0.23076923192, 0.30000001192, 6, 11},
	{40, 0.15384615958, 0.60000002384, 12, 17},
	{58, 0.13043478131, 0.80000001192, 20, 25},
}

func (warlock *Warlock) getShadowBoltBaseConfig(rank int) core.SpellConfig {
	// Beta client 1.60.1: every rank's damage moved and the low ranks lost their downranking penalty.
	// Damage is each rank's value at the level it stops scaling at (capped at 60), as the Classic table was.
	spellCoeff := [ShadowBoltRanks + 1]float64{0, .486, .629, .8, .857, .857, .857, .857, .857, .857, .857}[rank]
	baseDamage := [ShadowBoltRanks + 1][]float64{{0}, {12, 16}, {25, 31}, {41, 48}, {57, 64}, {79, 89}, {101, 113}, {140, 156}, {188, 210}, {237, 265}, {253, 283}}[rank]
	spellId := [ShadowBoltRanks + 1]int32{0, 686, 695, 705, 1088, 1106, 7641, 11659, 11660, 11661, 25307}[rank]
	manaCost := [ShadowBoltRanks + 1]float64{0, 25, 40, 70, 110, 160, 210, 265, 315, 370, 380}[rank]
	level := [ShadowBoltRanks + 1]int{0, 1, 6, 12, 20, 28, 36, 44, 52, 60, 60}[rank]
	castTime := [ShadowBoltRanks + 1]int32{0, 1700, 2200, 2800, 3000, 3000, 3000, 3000, 3000, 3000, 3000}[rank]
	if warlock.Env.IsForever() && rank < len(foreverShadowBoltRankDamage) {
		damage := foreverShadowBoltRankDamage[rank]
		growth := float64(max(0, min(int(warlock.Level), damage.maxLevel)-damage.spellLevel)) * damage.pointsPerLevel
		mean, spread := damage.basePoints+growth, damage.basePoints*damage.variance/2
		baseDamage = []float64{mean - spread, mean + spread}
	}

	return core.SpellConfig{
		SpellCode:     SpellCode_WarlockShadowBolt,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolShadow,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | core.SpellFlagResetAttackSwing | WarlockFlagDestruction,
		RequiredLevel: level,
		Rank:          rank,
		MissileSpeed:  20,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * time.Duration(castTime),
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcDamage(sim, target, sim.Roll(baseDamage[0], baseDamage[1]), spell.OutcomeMagicHitAndCrit)
			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
			})
		},
	}
}

func (warlock *Warlock) registerShadowBoltSpell() {
	warlock.ShadowBolt = make([]*core.Spell, 0)

	maxRank := core.TernaryInt(core.IncludeAQ, ShadowBoltRanks, ShadowBoltRanks-1)
	for rank := 1; rank <= maxRank; rank++ {
		config := warlock.getShadowBoltBaseConfig(rank)

		if config.RequiredLevel <= int(warlock.Level) {
			warlock.ShadowBolt = append(warlock.ShadowBolt, warlock.GetOrRegisterSpell(config))
		}
	}
}
