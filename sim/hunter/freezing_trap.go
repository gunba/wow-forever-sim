package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (hunter *Hunter) getFreezingTrapConfig(timer *core.Timer) core.SpellConfig {

	return core.SpellConfig{
		SpellCode:     SpellCode_HunterFreezingTrap,
		ActionID:      core.ActionID{SpellID: 1499},
		SpellSchool:   core.SpellSchoolFrost,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | SpellFlagTrap,
		RequiredLevel: 20,
		MissileSpeed:  24,

		ManaCost: core.ManaCostOptions{
			FlatCost: 50,
		},
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer: timer,
				// Forever doubles the shared trap cooldown to 30 sec. Seen on every trap tooltip
				// from the demo streams (Savix, Xaryu and Soda, 12-13 September).
				Duration: core.TernaryDuration(hunter.Env.IsForever(), time.Second*30, time.Second*15),
			},
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true, // Hunter GCD is locked at 1.5s
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			if hunter.Env.IsForever() && hunter.DistanceFromTarget <= 5 {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					hunter.procEntrapment(sim, target)
				}
			}
		},
	}
}

func (hunter *Hunter) registerFreezingTrapSpell(timer *core.Timer) {
	config := hunter.getFreezingTrapConfig(timer)

	if config.RequiredLevel <= int(hunter.Level) {
		hunter.FreezingTrap = hunter.GetOrRegisterSpell(config)
	}
}
