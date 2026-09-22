package priest

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const StarshardsRanks = 7
const StarshardsTicks = 6

var StarshardsSpellId = [StarshardsRanks + 1]int32{0, 10797, 19296, 19299, 19302, 19303, 19304, 19305}
var StarshardsTickSpellId = [StarshardsRanks + 1]int32{0, 19350, 19351, 19352, 19353, 19354, 19355, 19356}

// Forever beta client 1.60.1.69893, about double Classic's, at .167 a tick for every rank.
var StarshardsBaseDamage = [StarshardsRanks + 1]float64{0, 162, 300, 528, 762, 1068, 1440, 1800}
var StarshardsManaCost = [StarshardsRanks + 1]float64{0, 50, 85, 140, 190, 245, 300, 350}
var StarshardsLevel = [StarshardsRanks + 1]int{0, 10, 18, 26, 34, 42, 50, 58}

func (priest *Priest) registerStarshardsSpell() {
	if priest.Race != proto.Race_RaceNightElf {
		return
	}

	priest.Starshards = make([][]*core.Spell, StarshardsRanks+1)
	// SpellCooldowns: all ranks share a 30-second category cooldown.
	timer := priest.NewTimer()

	for rank := 1; rank <= StarshardsRanks; rank++ {
		priest.Starshards[rank] = make([]*core.Spell, StarshardsTicks+1)

		var tick int32
		for tick = 0; tick < StarshardsTicks; tick++ {
			config := priest.newStarshardsSpellConfig(rank, tick, timer)

			if config.RequiredLevel <= int(priest.Level) {
				priest.Starshards[rank][tick] = priest.RegisterSpell(config)
			}
		}
	}
}

func (priest *Priest) newStarshardsSpellConfig(rank int, tickIdx int32, timer *core.Timer) core.SpellConfig {
	ticks := tickIdx
	flags := SpellFlagPriest | core.SpellFlagChanneled | core.SpellFlagBinary
	if tickIdx == 0 {
		ticks = 6
		flags |= core.SpellFlagAPL
	}

	spellId := StarshardsSpellId[rank]
	baseDamage := StarshardsBaseDamage[rank] / StarshardsTicks
	manaCost := StarshardsManaCost[rank]
	level := StarshardsLevel[rank]

	spellCoeff := 0.167

	tickLength := time.Second

	return core.SpellConfig{
		SpellCode:   SpellCode_PriestStarshards,
		ActionID:    core.ActionID{SpellID: spellId}.WithTag(tickIdx),
		SpellSchool: core.SpellSchoolArcane,
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
			CD: core.Cooldown{Timer: timer, Duration: 30 * time.Second},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: fmt.Sprintf("Starshards-%d-%d", rank, tickIdx),
			},
			NumberOfTicks:       ticks,
			TickLength:          tickLength,
			AffectedByCastSpeed: false,
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
				spell.Dot(target).Apply(sim)
			}
			spell.DealOutcome(sim, result)
		},

		ExpectedTickDamage: func(sim *core.Simulation, target *core.Unit, spell *core.Spell, _ bool) *core.SpellResult {
			result := spell.CalcPeriodicDamage(sim, target, baseDamage, spell.OutcomeExpectedMagicAlwaysHit)
			return result
		},
	}
}
