package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (warrior *Warrior) registerWhirlwindSpell() {
	results := make([]*core.SpellResult, min(4, warrior.Env.GetNumTargets()))
	ohResults := make([]*core.SpellResult, len(results))

	// October 1 makes the off-hand attack baseline for dual wielding.
	var whirlwindOh *core.Spell
	if (warrior.Env.IsForever() || warrior.Talents.RagingBlows) && warrior.AutoAttacks.IsDualWielding {
		whirlwindOh = warrior.registerWhirlwindOffHandSpell()
	}

	warrior.Whirlwind = warrior.RegisterSpell(BerserkerStance, core.SpellConfig{
		SpellCode:   SpellCode_WarriorWhirlwind,
		ActionID:    core.ActionID{SpellID: 1680},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost: 25 - core.TernaryFloat64(warrior.Env.IsForever() && warrior.Talents.RagingBlows, 3, 0),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: time.Second * 10,
			},
		},
		CritDamageBonus: warrior.impale(),

		DamageMultiplier: 1,
		ThreatMultiplier: 1.25,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			mhTarget := target
			for idx := range results {
				baseDamage := spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target))
				results[idx] = spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
				target = sim.Environment.NextTargetUnit(target)
			}

			if whirlwindOh != nil {
				target = mhTarget
				for idx := range ohResults {
					baseDamage := whirlwindOh.Unit.OHNormalizedWeaponDamage(sim, whirlwindOh.MeleeAttackPower(target))
					ohResults[idx] = whirlwindOh.CalcDamage(sim, target, baseDamage, whirlwindOh.OutcomeMeleeWeaponSpecialHitAndCrit)
					target = sim.Environment.NextTargetUnit(target)
				}
			}

			for _, result := range results {
				spell.DealDamage(sim, result)
			}

			if whirlwindOh != nil {
				for _, result := range ohResults {
					whirlwindOh.DealDamage(sim, result)
				}
			}
		},
	})
}

func (warrior *Warrior) registerWhirlwindOffHandSpell() *core.Spell {
	return warrior.RegisterSpell(BerserkerStance, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 1680}.WithTag(2),
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeOHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

		CritDamageBonus: warrior.impale(),

		DamageMultiplier: warrior.AutoAttacks.OHConfig().DamageMultiplier,
		ThreatMultiplier: 1.25,
		BonusCoefficient: 1,
	}).Spell
}
