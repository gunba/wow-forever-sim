package mage

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const FireballRanks = 12

var FireballSpellId = [FireballRanks + 1]int32{0, 133, 143, 145, 3140, 8400, 8401, 8402, 10148, 10149, 10150, 10151, 25306}

// Beta client 1.60.1.69893. Damage is the client's base plus its per level growth up to the rank's
// max level (capped at 60), the same way the Classic numbers were read. Forever lowered every rank
// from 2 up and dropped the downranking penalty from the coefficients. The dot is the client's per
// tick damage times its tick count; the sim spreads that total over 4 ticks at every rank.
var FireballBaseDamage = [FireballRanks + 1][]float64{{0}, {16, 25}, {32, 47}, {48, 65}, {66, 91}, {98, 130}, {138, 183}, {171, 222}, {212, 275}, {271, 348}, {338, 431}, {397, 505}, {425, 541}}
var FireballDotDamage = [FireballRanks + 1]float64{0, 2, 3, 6, 12, 16, 24, 24, 32, 40, 48, 56, 60}
var FireballSpellCoeff = [FireballRanks + 1]float64{0, .429, .571, .714, .857, 1, 1, 1, 1, 1, 1, 1, 1}
var FireballCastTime = [FireballRanks + 1]int32{0, 1500, 2000, 2500, 3000, 3500, 3500, 3500, 3500, 3500, 3500, 3500, 3500}
var FireballManaCost = [FireballRanks + 1]float64{0, 30, 45, 65, 95, 140, 185, 220, 260, 305, 350, 395, 410}
var FireballLevel = [FireballRanks + 1]int{0, 1, 6, 12, 18, 24, 30, 36, 42, 48, 54, 60, 60}

// SpellEffect/SpellLevels 1.60.1.70291; patch context:
// https://us.forums.blizzard.com/en/wow/t/2360696/5
// Variance spreads the unscaled base. Capped level growth shifts the midpoint;
// retain the client floats rather than inventing integer endpoint rounding.
type mageSpellRankDamage struct {
	basePoints, variance, pointsPerLevel float64
	spellLevel, maxLevel                 int
}

func (damage mageSpellRankDamage) rangeAtLevel(level int32) (float64, float64) {
	growth := float64(max(0, min(int(level), damage.maxLevel)-damage.spellLevel)) * damage.pointsPerLevel
	mean, spread := damage.basePoints+growth, damage.basePoints*damage.variance/2
	return mean - spread, mean + spread
}

var foreverFireballRankDamage = [...]mageSpellRankDamage{
	{},
	{18, 0.44444444776, 0.20000000298, 1, 5},
	{32, 0.36842104793, 0.30000001192, 6, 10},
	{51, 0.31746032834, 0.60000002384, 12, 16},
	{77, 0.31999999285, 0.89999997616, 18, 22},
	{120, 0.29447853565, 1.39999997616, 24, 28},
}

func (mage *Mage) registerFireballSpell() {
	mage.Fireball = make([]*core.Spell, FireballRanks+1)

	maxRank := core.TernaryInt(core.IncludeAQ, FireballRanks, FireballRanks-1)
	for rank := 1; rank <= maxRank; rank++ {
		config := mage.newFireballSpellConfig(rank)

		if config.RequiredLevel <= int(mage.Level) {
			mage.Fireball[rank] = mage.GetOrRegisterSpell(config)
		}
	}
}

func (mage *Mage) newFireballSpellConfig(rank int) core.SpellConfig {
	numTicks := int32(4)
	// 70291 SpellMisc/SpellDuration: ranks1/2/3 last4/6/6s.
	// Periodic effects679021/679023/679025 specify2000ms and1/1/2
	// damage per tick, confirming the retained totals2/3/6.
	if mage.Env.IsForever() && rank <= 3 {
		numTicks = []int32{0, 2, 3, 3}[rank]
	}
	tickLength := time.Second * 2

	spellId := FireballSpellId[rank]
	baseDamageLow := FireballBaseDamage[rank][0]
	baseDamageHigh := FireballBaseDamage[rank][1]
	if mage.Env.IsForever() && rank < len(foreverFireballRankDamage) {
		baseDamageLow, baseDamageHigh = foreverFireballRankDamage[rank].rangeAtLevel(mage.Level)
	}
	baseDotDamage := FireballDotDamage[rank] / float64(numTicks)
	spellCoeff := FireballSpellCoeff[rank]
	castTime := FireballCastTime[rank]
	manaCost := FireballManaCost[rank]
	level := FireballLevel[rank]

	actionID := core.ActionID{SpellID: spellId}

	return core.SpellConfig{
		ActionID:     actionID,
		SpellCode:    SpellCode_MageFireball,
		SpellSchool:  core.SpellSchoolFire,
		DefenseType:  core.DefenseTypeMagic,
		ProcMask:     core.ProcMaskSpellDamage,
		Flags:        core.SpellFlagAPL | SpellFlagMage,
		MissileSpeed: 24,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*time.Duration(castTime) - time.Millisecond*100*time.Duration(mage.Talents.ImprovedFireball),
			},
		},

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label:    fmt.Sprintf("Fireball (Rank %d)", rank),
				ActionID: actionID.WithTag(1),
			},
			NumberOfTicks: numTicks,
			TickLength:    tickLength,
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, baseDotDamage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageLow, baseDamageHigh)
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)

				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			})
		},
	}
}
