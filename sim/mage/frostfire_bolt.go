package mage

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (mage *Mage) registerFrostfireBoltSpell() {
	if !mage.Env.IsForever() {
		return
	}

	// Beta 1.60.1.69893, SpellEffect/SpellLevels/SpellMisc. Rank 3 is
	// 270–314 at level 60; its static tooltip instead scales to level 68.
	ranks := []struct {
		id                     int32
		level                  int32
		low, high, growth, dot float64
		mana                   float64
	}{
		{401502, 40, 92, 108, 1.3, 9, 205},
		{1237312, 50, 168, 196, 1.7, 13, 285},
		{1237313, 60, 270, 314, 2.1, 19, 370},
	}
	mage.FrostfireBolt = make([]*core.Spell, len(ranks)+1)
	for index, rank := range ranks {
		if mage.Level < rank.level {
			continue
		}
		growth := rank.growth * float64(min(mage.Level-rank.level, 8))
		actionID := core.ActionID{SpellID: rank.id}
		mage.FrostfireBolt[index+1] = mage.RegisterSpell(core.SpellConfig{
			ActionID:     actionID,
			SpellCode:    SpellCode_MageFrostfireBolt,
			SpellSchool:  core.SpellSchoolFire | core.SpellSchoolFrost,
			DefenseType:  core.DefenseTypeMagic,
			ProcMask:     core.ProcMaskSpellDamage,
			Flags:        SpellFlagMage | SpellFlagChillSpell | core.SpellFlagAPL,
			MissileSpeed: 24,

			RequiredLevel: int(rank.level),
			Rank:          index + 1,
			ManaCost:      core.ManaCostOptions{FlatCost: rank.mana},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
					// Improved Fireball explicitly includes Frostfire Bolt;
					// Improved Frostbolt does not.
					CastTime: 3*time.Second - time.Duration(mage.Talents.ImprovedFireball)*100*time.Millisecond,
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: .814,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label:    fmt.Sprintf("Frostfire Bolt (Rank %d)", index+1),
					ActionID: actionID.WithTag(1),
				},
				NumberOfTicks: 3,
				TickLength:    3 * time.Second,
				// The client gives the DoT no spell-power coefficient.
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, rank.dot, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcDamage(sim, target, sim.Roll(rank.low+growth, rank.high+growth), spell.OutcomeMagicHitAndCrit)
				spell.WaitTravelTime(sim, func(sim *core.Simulation) {
					spell.DealDamage(sim, result)
					if result.Landed() {
						spell.Dot(target).Apply(sim)
					}
				})
			},
		})
	}
}
