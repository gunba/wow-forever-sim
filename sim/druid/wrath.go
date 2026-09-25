package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const WrathRanks = 8

var WrathSpellId = [WrathRanks + 1]int32{0, 5176, 5177, 5178, 5179, 5180, 6780, 8905, 9912}

// Client 1.60.1.70009 raises base damage and per-level growth on every rank.
// Ranges use the client variance and growth through each rank's max level.
var WrathBaseDamage = [WrathRanks + 1][]float64{{0}, {15, 18}, {25, 28}, {32, 37}, {41, 46}, {47, 53}, {56, 63}, {70, 77}, {92, 102}}
var WrathSpellCoeff = [WrathRanks + 1]float64{0, 0.429, 0.486, 0.571, 0.571, 0.571, 0.571, 0.571, 0.571}
var WrathManaCost = [WrathRanks + 1]float64{0, 10, 20, 40, 50, 70, 80, 100, 120}
var WrathCastTime = [WrathRanks + 1]int{0, 1500, 1700, 2000, 2000, 2000, 2000, 2000, 2000}
var WrathLevel = [WrathRanks + 1]int{0, 1, 6, 14, 22, 30, 38, 46, 54}

func (druid *Druid) registerWrathSpell() {
	druid.Wrath = make([]*DruidSpell, WrathRanks+1)

	for rank := 1; rank <= WrathRanks; rank++ {
		config := druid.newWrathSpellConfig(rank)

		if config.RequiredLevel <= int(druid.Level) {
			druid.Wrath[rank] = druid.RegisterSpell(Humanoid|Moonkin, config)
		}
	}
}

func (druid *Druid) newWrathSpellConfig(rank int) core.SpellConfig {
	spellId := WrathSpellId[rank]
	baseDamageLow := WrathBaseDamage[rank][0]
	baseDamageHigh := WrathBaseDamage[rank][1]
	spellCoeff := WrathSpellCoeff[rank]
	manaCost := WrathManaCost[rank]
	castTime := WrathCastTime[rank]
	level := WrathLevel[rank]

	return core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellId},
		SpellCode:   SpellCode_DruidWrath,
		SpellSchool: core.SpellSchoolNature,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL | core.SpellFlagResetAttackSwing,

		RequiredLevel: level,
		Rank:          rank,
		MissileSpeed:  20,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
			// Improved Wrath's beta client curves are 10-50% of the cost and 0.1-0.5 sec of the cast.
			Multiplier: 100 - 10*druid.Talents.ImprovedWrath,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*time.Duration(castTime) - time.Millisecond*100*time.Duration(druid.Talents.ImprovedWrath),
			},
		},

		DamageMultiplier: 1, // + core.Ternary(druid.Ranged().ID == IdolOfWrath, .02, 0),
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageLow, baseDamageHigh)
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			// NG procs when the cast finishes
			if result.DidCrit() {
				druid.procNaturesGrace(sim)
			}

			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
			})
		},
	}
}
