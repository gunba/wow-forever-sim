package mage

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const BlizzardRanks = 6

var BlizzardSpellId = [BlizzardRanks + 1]int32{0, 10, 6141, 8427, 10185, 10186, 10187}

// Beta client 1.60.1.69893. Forever's Blizzard is an area trigger that casts a damage spell every
// second (1279976 ... 1279949), 8 times; these are 8 times that spell's base. The coefficient below
// is the damage spell's, per tick; the parent spell's dummy effect carries 0.03, which is not used.
var BlizzardBaseDamage = [BlizzardRanks + 1]float64{0, 192, 336, 496, 696, 912, 1168}
var BlizzardManaCost = [BlizzardRanks + 1]float64{0, 320, 520, 720, 935, 1160, 1400}
var BlizzardLevel = [BlizzardRanks + 1]int{0, 20, 28, 36, 44, 52, 60}

func (mage *Mage) registerBlizzardSpell() {
	mage.Blizzard = make([]*core.Spell, BlizzardRanks+1)

	for rank := 1; rank <= BlizzardRanks; rank++ {
		config := mage.newBlizzardSpellConfig(rank)

		if config.RequiredLevel <= int(mage.Level) {
			mage.Blizzard[rank] = mage.GetOrRegisterSpell(config)
		}
	}
}

func (mage *Mage) newBlizzardSpellConfig(rank int) core.SpellConfig {
	numTicks := int32(8)
	tickLength := time.Second * 1

	spellId := BlizzardSpellId[rank]
	baseDamage := BlizzardBaseDamage[rank] / float64(numTicks)
	manaCost := BlizzardManaCost[rank]
	level := BlizzardLevel[rank]

	spellCoeff := .042

	var improvedBlizzardProcApplication *core.Spell
	if mage.Talents.ImprovedBlizzard > 0 {
		impId := []int32{0, 11185, 12487, 12488}[mage.Talents.ImprovedBlizzard]
		auras := mage.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
			return unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: impId},
				Label:    "Improved Blizzard",
				Duration: time.Millisecond * 1500,
			})
		})
		improvedBlizzardProcApplication = mage.RegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{SpellID: impId},
			ProcMask: core.ProcMaskSpellProc,
			Flags:    SpellFlagMage | core.SpellFlagNoLogs | SpellFlagChillSpell,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				auras.Get(target).Activate(sim)
			},
		})
	}

	return core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellId},
		SpellSchool: core.SpellSchoolFrost,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagMage | core.SpellFlagChanneled | core.SpellFlagAPL,

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

		Dot: core.DotConfig{
			IsAOE: true,
			Aura: core.Aura{
				Label: fmt.Sprintf("Blizzard (Rank %d)", rank),
			},
			NumberOfTicks:       numTicks,
			TickLength:          tickLength,
			AffectedByCastSpeed: true,
			BonusCoefficient:    spellCoeff,
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, baseDamage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, aoeTarget, dot.OutcomeTick)

					if improvedBlizzardProcApplication != nil {
						improvedBlizzardProcApplication.Cast(sim, aoeTarget)
					}
				}
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.AOEDot().Apply(sim)
		},
	}
}
