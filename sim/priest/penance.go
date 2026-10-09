package priest

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const PenanceTicks = 3

func (priest *Priest) registerPenanceSpell() {
	if !priest.Talents.Penance {
		return
	}

	type penanceDamageRank struct {
		spellID                                       int32
		level, maxLevel                               int
		mana, basePoints, pointsPerLevel, coefficient float64
	}
	// Preserve the prior, max-rank-only model outside Forever.
	ranks := []penanceDamageRank{{1316995, 60, 60, 355, 131, 0, .285}}
	if priest.Env.IsForever() {
		// SpellEffect/SpellLevels/SpellPower 1.60.1.70291.
		// https://us.forums.blizzard.com/en/wow/t/2360696/5
		// These are the damaging bolts, not the separately tuned healing children.
		ranks = []penanceDamageRank{
			{402174, 30, 39, 150, 42, .28999999166, .19},  // Bolt402284.
			{1240720, 40, 49, 220, 50, .33000001311, .19}, // Bolt1240727.
			{1240721, 50, 59, 270, 68, .43000000715, .19}, // Bolt1240730.
			{1316995, 60, 60, 385, 92, 0, .19},            // Bolt1316993.
		}
	}
	cooldown := priest.NewTimer()
	for index, rankData := range ranks {
		if rankData.level > int(priest.Level) {
			continue
		}
		rank := index + 1
		baseDamage := rankData.basePoints + float64(max(0, min(int(priest.Level), rankData.maxLevel)-rankData.level))*rankData.pointsPerLevel
		spellCoeff := rankData.coefficient
		label := "Penance"
		if rankData.spellID != 1316995 {
			label = fmt.Sprintf("Penance (Rank %d)", rank)
		}
		// Keep Penance as the highest registered rank; existing maximum-rank APLs
		// retain their spell identity, without selecting lower ranks automatically.
		priest.Penance = priest.RegisterSpell(core.SpellConfig{
			SpellCode:   SpellCode_PriestPenance,
			ActionID:    core.ActionID{SpellID: rankData.spellID},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,
			Flags:       SpellFlagPriest | core.SpellFlagAPL | core.SpellFlagChanneled,

			RequiredLevel: rankData.level,
			Rank:          rank,

			ManaCost: core.ManaCostOptions{
				FlatCost:   rankData.mana,
				Multiplier: 100 - 5*priest.Talents.ImprovedHealing,
			},

			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
				CD: core.Cooldown{
					Timer:    cooldown,
					Duration: time.Second * 12,
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: label,
				},

				// The first bolt lands on application; the two remaining bolts
				// land one and two seconds later (402261's one-second period).
				NumberOfTicks:       PenanceTicks - 1,
				TickLength:          time.Second,
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
					dot := spell.Dot(target)
					dot.Apply(sim)
					dot.OnTick(sim, target, dot)
				}
				spell.DealOutcome(sim, result)
			},

			ExpectedTickDamage: func(sim *core.Simulation, target *core.Unit, spell *core.Spell, _ bool) *core.SpellResult {
				return spell.CalcPeriodicDamage(sim, target, baseDamage, spell.OutcomeExpectedMagicAlwaysHit)
			},
		})
	}
}
