package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (warlock *Warlock) registerWrackSpell() {
	if !warlock.Talents.Wrack {
		return
	}

	// Beta client 1.60.1 (spell 1316697): 36 per tick with a 0.143 coefficient, 6 sec channel,
	// 200 mana. The client stores one value, not a per-level table, so it is used as is.
	numTicks := int32(6)
	tickLength := time.Second
	baseDamage := 36.0
	spellCoeff := 0.143
	manaCost := 200.0

	warlock.Wrack = warlock.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_WarlockWrack,
		ActionID:    core.ActionID{SpellID: 11704},
		SpellSchool: core.SpellSchoolShadow,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL | core.SpellFlagChanneled | core.SpellFlagResetAttackSwing | WarlockFlagAffliction,

		RequiredLevel: 60,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Wrack-" + warlock.Label,
			},
			NumberOfTicks:       numTicks,
			TickLength:          tickLength,
			AffectedByCastSpeed: true,
			BonusCoefficient:    spellCoeff,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, baseDamage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				result := dot.CalcSnapshotDamage(sim, target, dot.OutcomeTick)
				result.Damage *= warlock.improvedDrainsMultiplier(sim, target)
				dot.Spell.DealPeriodicDamage(sim, result)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
			if result.Landed() {
				dot := spell.Dot(target)
				dot.Apply(sim)
			}
			spell.DealOutcome(sim, result)
		},
	})

	// Other shadow dots on the target tick for 10% more while Wrack is on it
	for _, target := range warlock.Env.Encounter.TargetUnits {
		target.AddDynamicDamageTakenModifier(func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.Unit != &warlock.Unit || spell == warlock.Wrack {
				return
			}

			if spell.SpellSchool.Matches(core.SpellSchoolShadow) && len(spell.Dots()) > 0 && warlock.Wrack.Dot(result.Target).IsActive() {
				result.Damage *= 1.1
			}
		})
	}
}
