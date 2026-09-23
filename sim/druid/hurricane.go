package druid

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (druid *Druid) registerHurricaneSpell() {
	ranks := []struct {
		level      int32
		spellID    int32
		manaCost   float64
		scaleLevel int32
		damage     float64
		scale      float64
	}{
		// Beta client 1.60.1.69893: each second the storm casts a separate damage spell (1278965, 1278968, 1278759),
		// whose damage is what the tick deals here, 2 less than Classic's at every rank with the same .03 coefficient.
		// Written as floats so the spell id walk in sim/spell_sources_test.go does not read them as ids.
		{level: 40, spellID: 16914, manaCost: 880, scaleLevel: 46, damage: 68.0, scale: 0.2},
		{level: 50, spellID: 17401, manaCost: 1180, scaleLevel: 56, damage: 98.0, scale: 0.2},
		{level: 60, spellID: 17402, manaCost: 1495, scaleLevel: 66, damage: 132.0, scale: 0.3},
	}

	for i, rank := range ranks {
		if druid.Level < rank.level {
			break
		}

		damage := rank.damage + float64(min(druid.Level, rank.scaleLevel)-rank.level)*rank.scale
		spell := druid.RegisterSpell(Humanoid|Moonkin, core.SpellConfig{
			SpellCode:   SpellCode_DruidHurricane,
			ActionID:    core.ActionID{SpellID: rank.spellID},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,
			Flags:       core.SpellFlagChanneled | core.SpellFlagBinary | core.SpellFlagAPL,

			RequiredLevel: int(rank.level),
			Rank:          i + 1,

			ManaCost: core.ManaCostOptions{
				FlatCost: rank.manaCost,
			},
			Cast: core.CastConfig{
				// Forever drops Classic's 1 min cooldown: the client's Hurricane category has no recovery time.
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				IsAOE: true,
				Aura: core.Aura{
					Label: "Hurricane" + druid.Label + strconv.Itoa(i+1),
				},
				NumberOfTicks:       10,
				TickLength:          time.Second * 1,
				AffectedByCastSpeed: true,

				BonusCoefficient: 0.03,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, damage, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					for _, aoeTarget := range sim.Encounter.TargetUnits {
						dot.CalcAndDealPeriodicSnapshotDamage(sim, aoeTarget, dot.OutcomeTick)
					}
				},
			},

			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
				druid.AutoAttacks.CancelAutoSwing(sim)
				spell.AOEDot().Apply(sim)
			},
		})

		druid.Hurricane = append(druid.Hurricane, spell)
	}
}
