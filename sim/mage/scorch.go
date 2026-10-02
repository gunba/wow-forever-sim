package mage

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const ScorchRanks = 7

var ScorchSpellId = [ScorchRanks + 1]int32{0, 2948, 8444, 8445, 8446, 10205, 10206, 10207}

// Beta client 1.60.1.69893, about 30% below Classic at every rank.
var ScorchBaseDamage = [ScorchRanks + 1][]float64{{0}, {38, 47}, {54, 64}, {67, 79}, {89, 106}, {111, 132}, {143, 169}, {166, 197}}
var ScorchManaCost = [ScorchRanks + 1]float64{0, 50, 65, 80, 100, 115, 135, 150}
var ScorchLevel = [ScorchRanks + 1]int{0, 22, 28, 34, 40, 46, 52, 58}

func (mage *Mage) registerScorchSpell() {
	mage.Scorch = make([]*core.Spell, ScorchRanks+1)

	for rank := 1; rank <= ScorchRanks; rank++ {
		config := mage.getScorchConfig(rank)

		if config.RequiredLevel <= int(mage.Level) {
			mage.Scorch[rank] = mage.GetOrRegisterSpell(config)
		}
	}
}

func (mage *Mage) getScorchConfig(rank int) core.SpellConfig {
	spellId := ScorchSpellId[rank]
	baseDamageLow := ScorchBaseDamage[rank][0]
	baseDamageHigh := ScorchBaseDamage[rank][1]
	manaCost := ScorchManaCost[rank]
	level := ScorchLevel[rank]

	spellCoeff := .429
	debuffProcChance := []float64{0, .33, .66, 1}[mage.Talents.ImprovedScorch]

	return core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellId},
		SpellCode:   SpellCode_MageScorch,
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL | SpellFlagMage,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 1500,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageLow, baseDamageHigh)
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			// The landed Scorch supplies the hit check; the aura has no second resist roll.
			if result.Landed() && sim.RandomFloat("Improved Scorch") < debuffProcChance {
				mage.ImprovedScorchAura.Activate(sim)
				mage.ImprovedScorchAura.AddStack(sim)
			}
		},
	}
}
