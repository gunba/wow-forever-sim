package priest

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const PenanceTicks = 3

func (priest *Priest) registerPenanceSpell() {
	if !priest.Talents.Penance {
		return
	}

	// Rank 4 (1316995), the level 60 rank in the Forever beta client 1.60.1.69893: 355 mana, a 12 sec
	// cooldown and 131 Holy damage a bolt (1316993) at .285. Ranks 1-3 learn at 30, 40 and 50. The
	// client's rank 3 bolt (180) is larger than rank 4's; the numbers are taken as they are.
	baseDamage := 131.0
	spellCoeff := 0.285

	priest.Penance = priest.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_PriestPenance,
		ActionID:    core.ActionID{SpellID: 1316995},
		SpellSchool: core.SpellSchoolHoly,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagPriest | core.SpellFlagAPL | core.SpellFlagChanneled,

		RequiredLevel: 60,
		Rank:          1,

		ManaCost: core.ManaCostOptions{
			FlatCost: 355,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    priest.NewTimer(),
				Duration: time.Second * 12,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Penance",
			},

			// The first bolt lands on application; the two remaining bolts
			// land one and two seconds later (402261's one-second period).
			NumberOfTicks:    PenanceTicks - 1,
			TickLength:       time.Second,
			BonusCoefficient: spellCoeff,

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
