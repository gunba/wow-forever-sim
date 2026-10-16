package priest

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const DevouringPlagueRanks = 6

var DevouringPlagueSpellId = [DevouringPlagueRanks + 1]int32{0, 2944, 19276, 19277, 19278, 19279, 19280}

// Forever beta client 1.60.1.69893.
var DevouringPlagueBaseDamage = [DevouringPlagueRanks + 1]float64{0, 128, 232, 344, 488, 656, 848}
var DevouringPlagueManaCost = [DevouringPlagueRanks + 1]float64{0, 215, 350, 495, 645, 810, 985}
var DevouringPlagueLevel = [DevouringPlagueRanks + 1]int{0, 20, 28, 36, 44, 52, 60}

func (priest *Priest) registerDevouringPlagueSpell() {
	//TO DO: Implement race requirement
	priest.DevouringPlague = make([]*core.Spell, DevouringPlagueRanks+1)
	cdTimer := priest.NewTimer()

	for rank := 1; rank <= DevouringPlagueRanks; rank++ {
		config := priest.getDevouringPlagueConfig(rank, cdTimer)

		if config.RequiredLevel <= int(priest.Level) {
			priest.DevouringPlague[rank] = priest.GetOrRegisterSpell(config)
		}
	}
}

func (priest *Priest) getDevouringPlagueConfig(rank int, cdTimer *core.Timer) core.SpellConfig {

	var ticks int32 = 8

	spellId := DevouringPlagueSpellId[rank]
	baseDotDamage := (DevouringPlagueBaseDamage[rank] / float64(ticks))
	manaCost := DevouringPlagueManaCost[rank]
	level := DevouringPlagueLevel[rank]

	spellCoeff := 0.1 // per tick

	return core.SpellConfig{
		SpellCode:   SpellCode_PriestDevouringPlague,
		ActionID:    core.ActionID{SpellID: spellId},
		SpellSchool: core.SpellSchoolShadow,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagPriest | core.SpellFlagAPL | core.SpellFlagDisease | core.SpellFlagPureDot,

		Rank:          rank,
		RequiredLevel: level,

		// Devouring Contagion, 25/50% in the beta client.
		ManaCost: core.ManaCostOptions{
			FlatCost:   manaCost,
			Multiplier: 100 - 25*priest.Talents.DevouringContagion,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: 24 * time.Second, // 1 min until client 1.60.1.70425; 3 min in Classic
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: fmt.Sprintf("Devouring Plague (Rank %d)", rank),
			},

			NumberOfTicks:    ticks,
			TickLength:       time.Second * 3,
			BonusCoefficient: spellCoeff,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, baseDotDamage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
			if result.Landed() {
				priest.AddShadowWeavingStack(sim)
				spell.Dot(target).Apply(sim)
			}
			spell.DealOutcome(sim, result)
		},
	}
}
