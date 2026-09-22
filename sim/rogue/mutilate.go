package rogue

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (rogue *Rogue) registerMutilateSpell() {
	if !rogue.Talents.Mutilate {
		return
	}

	// The beta client ranks Mutilate: 23, 33, 48 and 67 flat added to each hand's weapon damage
	// before the 75%. Rank 1 is level 30, so level 25 borrows it.
	flatDamage := map[int32]float64{
		25: 23,
		40: 33,
		50: 48,
		60: 67,
	}[rogue.Level]

	spellID := map[int32]int32{
		25: 1310707,
		40: 399956,
		50: 1241582,
		60: 1241584,
	}[rogue.Level]

	actionID := core.ActionID{SpellID: spellID}

	rogue.mutilateOH = rogue.RegisterSpell(core.SpellConfig{
		ActionID:    actionID.WithTag(2),
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeOHSpecial,
		Flags:       SpellFlagBuilder | core.SpellFlagMeleeMetrics | core.SpellFlagNoOnCastComplete,

		// Puncturing Wounds' family mask includes both child strikes.
		BonusCritRating: 5 * core.CritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds),
		CritDamageBonus: rogue.lethality(),

		DamageMultiplier: rogue.AutoAttacks.OHConfig().DamageMultiplier * []float64{1, 1.05, 1.1}[rogue.Talents.Opportunity],
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := rogue.mutilateDamage(target, flatDamage, rogue.OHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target)))
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
		},
	})

	rogue.Mutilate = rogue.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_RogueMutilate,
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       rogue.builderFlags(),

		EnergyCost: core.EnergyCostOptions{
			Cost:   60,
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			// Every rank, and both hand strikes, require a Dagger in the beta client.
			return rogue.HasDagger(core.MainHand) && rogue.HasDagger(core.OffHand)
		},

		// Puncturing Wounds gives 5% per rank (beta client).
		BonusCritRating: 5 * core.CritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds),

		CritDamageBonus: rogue.lethality(),

		DamageMultiplier: []float64{1, 1.05, 1.1}[rogue.Talents.Opportunity],
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			rogue.BreakStealth(sim)

			// Cold Blood is spent on the main hand half, which is the larger of the two.
			baseDamage := rogue.mutilateDamage(target, flatDamage, rogue.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target)))
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			rogue.mutilateOH.Cast(sim, target)

			if result.Landed() {
				rogue.AddComboPoints(sim, 2, target, spell.ComboPointMetrics())
			} else {
				spell.IssueRefund(sim)
			}
		},
	})
}

// Each half strikes for 75% of weapon damage plus a flat bonus, and hits harder while one of
// the rogue's lingering poisons is on the target.
func (rogue *Rogue) mutilateDamage(target *core.Unit, flatDamage float64, weaponDamage float64) float64 {
	baseDamage := 0.75 * (flatDamage + weaponDamage)
	if rogue.isPoisoned(target) {
		baseDamage *= 1.2
	}
	return baseDamage
}

// Instant Poison leaves nothing behind, so only the two lingering poisons count.
func (rogue *Rogue) isPoisoned(target *core.Unit) bool {
	return rogue.deadlyPoisonTick.Dot(target).IsActive() || rogue.woundPoisonDebuffAuras.Get(target).IsActive()
}
