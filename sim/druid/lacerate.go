package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const LacerateMaxStacks int32 = 5

// Forever trains Lacerate at 42, 50 and 58 (414644, 1235826, 1235827). Beta client 1.60.1.69893: 15 Rage, 10% weapon
// damage per stack on the hit, and a bleed of 10 / 12 / 15 a tick per stack over 15 sec that no longer scales with
// attack power. The client does not carry threat, so the 3.33x is still Season of Discovery's.
func (druid *Druid) registerLacerateSpell() {
	druid.registerLacerateBleedSpell()

	druid.Lacerate = druid.RegisterSpell(Bear, core.SpellConfig{
		SpellCode:   SpellCode_DruidLacerate,
		ActionID:    core.ActionID{SpellID: 414644},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		RageCost: core.RageCostOptions{
			Cost:   15 - float64(druid.Talents.ShreddingAttacks),
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 3.33,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			// Berserk's extra-target modifier names only Mangle (family mask 0x40),
			// not Lacerate (0x100).
			stacks := min(druid.LacerateBleed.Dot(target).GetStacks()+1, LacerateMaxStacks)
			baseDamage := spell.Unit.MHWeaponDamage(sim, spell.MeleeAttackPower(target)) * 0.1 * float64(stacks)
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)
			if result.Landed() {
				druid.LacerateBleed.Cast(sim, target)
			}
			if !result.Landed() {
				spell.IssueRefund(sim)
			}
		},
	})
}

func (druid *Druid) registerLacerateBleedSpell() {
	tickDamage := 10.0
	if druid.Level >= 58 {
		tickDamage = 15
	} else if druid.Level >= 50 {
		tickDamage = 12
	}

	druid.LacerateBleed = druid.RegisterSpell(Bear, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 414647},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagNoOnCastComplete,

		DamageMultiplier: 1,
		ThreatMultiplier: 3.33,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label:     "Lacerate",
				MaxStacks: LacerateMaxStacks,
				Duration:  time.Second * 15,
			},
			NumberOfTicks: 5,
			TickLength:    time.Second * 3,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, tickDamage*float64(dot.Aura.GetStacks()), isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			dot := spell.Dot(target)
			if dot.IsActive() {
				dot.Refresh(sim)
				dot.AddStack(sim)
			} else {
				dot.Apply(sim)
				dot.SetStacks(sim, 1)
			}
			// Snapshot again once the stacks are in, since the damage grows with them.
			dot.TakeSnapshot(sim, false)
		},
	})
}
