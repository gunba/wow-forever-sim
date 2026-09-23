package priest

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const MindFlayRanks = 6
const MindFlayTicks = 3

var MindFlaySpellId = [MindFlayRanks + 1]int32{0, 15407, 17311, 17312, 17313, 17314, 18807}
var MindFlayTickSpellId = [MindFlayRanks + 1]int32{0, 16568, 7378, 17316, 17317, 17318, 18808}

// Forever beta client 1.60.1.69893, the total over the three ticks. The demo's 119 at rank 1 is not the
// client's 21 a tick; every rank is a little below Classic's.
var MindFlayBaseDamage = [MindFlayRanks + 1]float64{0, 63, 102, 153, 225, 294, 390}
var MindFlayManaCost = [MindFlayRanks + 1]float64{0, 45, 70, 100, 135, 165, 205}
var MindFlayLevel = [MindFlayRanks + 1]int{0, 20, 28, 36, 44, 52, 60}

func (priest *Priest) registerMindFlay() {
	if !priest.Talents.MindFlay {
		return
	}

	priest.MindFlay = make([][]*core.Spell, MindFlayRanks+1)

	for rank := 1; rank <= MindFlayRanks; rank++ {
		priest.MindFlay[rank] = make([]*core.Spell, MindFlayTicks+1)

		var tick int32
		for tick = 0; tick < MindFlayTicks; tick++ {
			config := priest.newMindFlaySpellConfig(rank, tick)

			if config.RequiredLevel <= int(priest.Level) {
				priest.MindFlay[rank][tick] = priest.RegisterSpell(config)
			}
		}
	}
}

func (priest *Priest) newMindFlaySpellConfig(rank int, tickIdx int32) core.SpellConfig {
	ticks := tickIdx
	flags := SpellFlagPriest | core.SpellFlagChanneled | core.SpellFlagBinary
	if tickIdx == 0 {
		ticks = 3
		flags |= core.SpellFlagAPL
	}

	spellId := MindFlaySpellId[rank]
	baseDamage := MindFlayBaseDamage[rank] / float64(MindFlayTicks)
	manaCost := MindFlayManaCost[rank]
	level := MindFlayLevel[rank]

	spellCoeff := 0.167 // per tick, the Forever beta client's (Classic's .15 carried a penalty for the slow)

	tickLength := time.Second

	return core.SpellConfig{
		SpellCode:   SpellCode_PriestMindFlay,
		ActionID:    core.ActionID{SpellID: spellId}.WithTag(tickIdx),
		SpellSchool: core.SpellSchoolShadow,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       flags,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		// Improved Mind Flay, 10/20% in the beta client.
		DamageMultiplier: 1 + 0.1*float64(priest.Talents.ImprovedMindFlay),
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: fmt.Sprintf("MindFlay-%d-%d", rank, tickIdx),
			},
			NumberOfTicks:       ticks,
			TickLength:          tickLength,
			AffectedByCastSpeed: true,
			BonusCoefficient:    spellCoeff,
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, baseDamage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHit)
			if result.Landed() {
				priest.AddShadowWeavingStack(sim)
				spell.Dot(target).Apply(sim)
			}
			spell.DealOutcome(sim, result)
		},

		ExpectedTickDamage: func(sim *core.Simulation, target *core.Unit, spell *core.Spell, _ bool) *core.SpellResult {
			return spell.CalcPeriodicDamage(sim, target, baseDamage, spell.OutcomeExpectedMagicAlwaysHit)
		},
	}
}
