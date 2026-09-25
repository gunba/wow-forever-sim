package warrior

import (
	"github.com/wowsims/classic/sim/core"
)

func (warrior *Warrior) registerDemoralizingShoutSpell() {
	rank := int32(5)
	actionId := core.DemoralizingShoutSpellId[rank]

	warrior.DemoralizingShoutAuras = warrior.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		// Improved Demoralizing Shout is gone from the Forever tree and is baseline at full strength:
		// the beta client's rank 5 reduces attack power by 196, Classic's 140 plus 40%, for 45 sec.
		// Forever's Booming Voice only widens the radius.
		return core.DemoralizingShoutAura(target)
	})
	threatMultiplier := 0.4
	flatThreatBonus := 0.4 * 2 * float64(core.DemoralizingShoutLevel[rank])
	if warrior.Env.IsForever() {
		// Forever's tank guide says Demoralizing Shout no longer generates threat.
		threatMultiplier = 0
		flatThreatBonus = 0
	}

	warrior.DemoralizingShout = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: actionId},
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost: 10,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ThreatMultiplier: threatMultiplier,
		FlatThreatBonus:  flatThreatBonus,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for _, aoeTarget := range sim.Encounter.TargetUnits {
				result := spell.CalcAndDealOutcome(sim, aoeTarget, spell.OutcomeMagicHit)
				if result.Landed() {
					warrior.DemoralizingShoutAuras.Get(aoeTarget).Activate(sim)
				}
			}
		},

		RelatedAuras: []core.AuraArray{warrior.DemoralizingShoutAuras},
	})
}
