package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (druid *Druid) registerFaerieFireSpell() {
	actionID := core.ActionID{SpellID: 9907}
	manaCostOptions := core.ManaCostOptions{
		FlatCost: 115,
	}
	flatThreatBonus := 2. * 54

	druid.FaerieFireAuras = druid.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.FaerieFireAura(target)
	})

	// Client 9907 allows Cat, Bear, Dire Bear and Moonkin (mask 1073741969).
	// It retains its mana cost and 1.5-second base GCD, with no six-second CD.
	druid.FaerieFire = druid.RegisterSpell(Humanoid|Moonkin|Cat|Bear, core.SpellConfig{
		SpellCode:   SpellCode_DruidFaerieFire,
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolNature,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCostOptions,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		ThreatMultiplier: 1,
		FlatThreatBonus:  flatThreatBonus,
		DamageMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
			if result.Landed() {
				druid.FaerieFireAuras.Get(target).Activate(sim)
			}

			if druid.InForm(Humanoid | Moonkin) {
				druid.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime, false)
			}
		},

		RelatedAuras: []core.AuraArray{druid.FaerieFireAuras},
	})
}

func (druid *Druid) ShouldFaerieFire(sim *core.Simulation, target *core.Unit) bool {
	if druid.FaerieFire == nil {
		return false
	}

	if !druid.FaerieFire.IsReady(sim) {
		return false
	}

	debuff := druid.FaerieFireAuras.Get(target)
	return !debuff.IsActive() || debuff.RemainingDuration(sim) < time.Second*4
}
