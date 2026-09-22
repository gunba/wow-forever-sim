package hunter

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (hunter *Hunter) registerVolleySpell() {
	ranks := 3

	for i := ranks; i >= 0; i-- {
		config := hunter.getVolleyConfig(i)

		if config.RequiredLevel <= int(hunter.Level) {
			hunter.Volley = hunter.GetOrRegisterSpell(config)
			break
		}
	}
}

func (hunter *Hunter) getVolleyConfig(rank int) core.SpellConfig {
	spellId := [4]int32{0, 1510, 14294, 14295}[rank]
	// Forever fires a separate damage spell each tick (1279721, 1279719, 1279715).
	baseDamage := [4]float64{0, 70, 91, 112}[rank]
	manaCost := [4]float64{0, 350, 420, 490}[rank]
	level := [4]int{0, 40, 50, 58}[rank]

	return core.SpellConfig{
		SpellCode:   SpellCode_HunterVolley,
		ActionID:    core.ActionID{SpellID: spellId},
		SpellSchool: core.SpellSchoolArcane,
		DefenseType: core.DefenseTypeRanged,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagChanneled | core.SpellFlagAPL,

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
				Label: fmt.Sprintf("Volley (Rank %d)", rank),
			},
			NumberOfTicks: 6,
			TickLength:    time.Second * 1,
			// The tick spell has no coefficient and the channel's dummy effect carries .03, the
			// same placeholder Blizzard and Rain of Fire carry, so Classic's .056 a tick stands.
			BonusCoefficient: .056,
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				damage := baseDamage
				dot.Snapshot(target, damage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, aoeTarget, dot.OutcomeTick)
				}
			},
		},

		CritDamageBonus:  hunter.mortalShots(),
		DamageMultiplier: 1 + []float64{0, .03, .07, .10}[hunter.Talents.Barrage],
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			hunter.Unit.AutoAttacks.DelayRangedUntil(sim, sim.CurrentTime+(time.Second*6))
			spell.AOEDot().Apply(sim)
		},
	}
}
