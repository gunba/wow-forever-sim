package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (hunter *Hunter) getMongooseBiteConfig(rank int) core.SpellConfig {
	spellId := [5]int32{0, 1495, 14269, 14270, 14271}[rank]
	baseDamage := [5]float64{0, 15, 22, 37, 57}[rank]
	manaCost := [5]float64{0, 30, 40, 50, 65}[rank]
	level := [5]int{0, 16, 30, 44, 58}[rank]

	spellConfig := core.SpellConfig{
		SpellCode:     SpellCode_HunterMongooseBite,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 5,
			},
		},

		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance && hunter.DefensiveState.IsActive()
		},

		BonusCritRating:  hunter.savageStrikesLegacyCrit(),
		CritDamageBonus:  hunter.mortalShots(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			hunter.DefensiveState.Deactivate(sim)
			// Forever: normalized melee weapon damage plus a smaller flat amount, where Classic dealt the flat amount alone.
			damage := baseDamage + hunter.AutoAttacks.MH().CalculateNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target))
			result := spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)

			if hunter.LaceratingStrikes != nil && result.Landed() {
				hunter.procLaceratingStrikes(sim, result)
			}
		},
	}

	return spellConfig
}

func (hunter *Hunter) registerMongooseBiteSpell() {
	// Aura is only used as a pre-requisite for Mongoose Bite
	hunter.DefensiveState = hunter.RegisterAura(core.Aura{
		Label:    "Defensive State",
		ActionID: core.ActionID{SpellID: 5302},
		Duration: time.Second * 5,
	})
	core.MakePermanent(hunter.RegisterAura(core.Aura{
		Label: "Mongoose Bite Dodge Trigger",
		OnSpellHitTaken: func(_ *core.Aura, sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
			if result.DidDodge() {
				hunter.DefensiveState.Activate(sim)
			}
		},
	}))

	rank := map[int32]int{
		25: 1,
		40: 2,
		50: 3,
		60: 4,
	}[hunter.Level]

	config := hunter.getMongooseBiteConfig(rank)
	hunter.MongooseBite = hunter.GetOrRegisterSpell(config)
}
