package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

const RevengeRanks = 6

var RevengeSpellId = [RevengeRanks + 1]int32{0, 6572, 6574, 7379, 11600, 11601, 25288}

// Forever raises every rank by about 70% over Classic's.
var RevengeBaseDamage = [RevengeRanks + 1][]float64{{0, 0}, {20, 24}, {31, 37}, {43, 53}, {73, 91}, {109, 133}, {138, 168}}
var RevengeLevel = [RevengeRanks + 1]int{0, 14, 24, 34, 44, 54, 60}

func (warrior *Warrior) registerRevengeSpell(cdTimer *core.Timer) {
	rank := warrior.rank(ranks.Revenge)
	if rank == nil {
		return
	}
	actionID := core.ActionID{SpellID: core.TernaryInt32(core.IncludeAQ, 25288, 11601)}
	damage := rank.Find(ranks.EffectSchoolDamage, 0)
	basedamageLow, basedamageHigh := damage.Low, damage.High
	revengeLevel := float64(rank.Level)

	warrior.revengeProcAura = warrior.RegisterAura(core.Aura{
		Label:    "Revenge",
		Duration: 5 * time.Second,
		ActionID: actionID,
	})

	warrior.RegisterAura(core.Aura{
		Label:    "Revenge Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.Outcome.Matches(core.OutcomeBlock | core.OutcomeDodge | core.OutcomeParry) {
				warrior.revengeProcAura.Activate(sim)
			}
		},
	})

	warrior.Revenge = warrior.RegisterSpell(DefensiveStance, core.SpellConfig{
		SpellCode:   SpellCode_WarriorRevenge,
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost:   5,
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: time.Second * 5,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return warrior.revengeProcAura.IsActive()
		},

		CritDamageBonus: warrior.impale(),

		DamageMultiplier: 1 + 0.2*float64(warrior.Talents.ImprovedRevenge),
		// Forever: threat is 2.25x damage with no flat part, and damage gains AP scaling
		// (published 'Warrior Threat Values' sheet, six-point fit at rank 2: 1.495 * 2.25 * (roll + 0.28 AP), residuals < 0.5;
		// 0.20 measured at rank 1). Higher ranks keep the rank-2 value until measured.
		ThreatMultiplier: 2.25,
		FlatThreatBonus:  core.TernaryFloat64(warrior.Env.IsForever(), 0, 2.25*2*revengeLevel),
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(basedamageLow, basedamageHigh)
			if sim.IsForever() {
				baseDamage += ranks.RevengeAPCoefficient * spell.MeleeAttackPower(target)
			}
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)

			if !result.Landed() {
				spell.IssueRefund(sim)
			}

			warrior.revengeProcAura.Deactivate(sim)
		},
	})
}
