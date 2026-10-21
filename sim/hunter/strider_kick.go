package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Strider Kick from the beta client (1317257): 100% normalized melee weapon damage, 8 sec
// cooldown, 5.81% of base mana.
func (hunter *Hunter) registerStriderKickSpell() {
	// Effect 3 is A_MOD_INCREASE_SPEED 30 on the caster for 3 sec ("Strider Kick now increases movement speed by
	// 30% for 3 seconds", 24 September notes): it shortens a weave's walk back out to Auto Shot range.
	var speedAura *core.Aura
	if hunter.Env.IsForever() {
		speedAura = hunter.RegisterAura(core.Aura{
			Label:    "Strider Kick Speed",
			ActionID: core.ActionID{SpellID: 1317257, Tag: 1},
			Duration: time.Second * 3,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddMoveSpeedModifierDynamic(sim, &aura.ActionID, 1.30)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.RemoveMoveSpeedModifierDynamic(sim, &aura.ActionID)
			},
		})
	}
	if !hunter.Talents.StriderKick {
		return
	}

	hunter.StriderKick = hunter.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_HunterStriderKick,
		ActionID:    core.ActionID{SpellID: 1317257},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost: 0.0581,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 8,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance
		},

		BonusCritRating:  hunter.savageStrikesLegacyCrit(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := hunter.AutoAttacks.MH().CalculateNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target))
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			if speedAura != nil {
				speedAura.Activate(sim)
			}
		},
	})
}
