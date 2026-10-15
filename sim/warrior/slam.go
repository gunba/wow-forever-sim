package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

func (warrior *Warrior) registerSlamSpell() {
	rank := warrior.rank(ranks.Slam)
	if rank == nil {
		return
	}
	requiredLevel := int(rank.Level)
	spellID := int32(11605)
	flatDamageBonus := rank.Find(ranks.EffectWeaponDamageBonus, 0).Low

	// Improved Slam takes the same amount off the global cooldown as it does the cast time.
	castReduction := time.Millisecond * 250 * time.Duration(warrior.Talents.ImprovedSlam)
	slamCooldown := core.Cooldown{}
	if warrior.Env.IsForever() {
		// Client 70291+: an 18-second category cooldown on every rank. Improved Slam's cooldown
		// modifier (12862 effect 2, op 11) is -3000 ms at both ranks (talent curve [0.25, 3], [0.5, 3]).
		reduction := time.Duration(0)
		if warrior.Talents.ImprovedSlam > 0 {
			reduction = time.Second * 3
		}
		slamCooldown = core.Cooldown{
			Timer:    warrior.NewTimer(),
			Duration: time.Second*18 - reduction,
		}
	}

	// Untalented Slam is an ordinary weapon hardcast in Forever as in Classic: its ranks carry the
	// Combat interrupt flag (SpellInterrupts 15) and lack SPELL_ATTR2_DO_NOT_RESET_COMBAT_TIMERS, so
	// no auto-attacks land during the cast and both swing timers restart when it completes.
	// Improved Slam overrides every rank with 1310196-1310200, which drop the Combat flag (7) and
	// set DO_NOT_RESET_COMBAT_TIMERS: "Slam no longer interrupts or delays your melee swing".
	flags := core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive
	if warrior.Env.IsForever() && warrior.Talents.ImprovedSlam > 0 {
		flags |= core.SpellFlagAllowAutoAttacks
	}

	warrior.Slam = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		SpellCode:   SpellCode_WarriorSlam,
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       flags,

		RequiredLevel: requiredLevel,

		RageCost: core.RageCostOptions{
			Cost:   15,
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      max(core.GCDDefault-castReduction, core.GCDMin),
				CastTime: time.Millisecond*1500 - castReduction,
			},
			CD: slamCooldown,
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				if castTime := warrior.ApplyCastSpeedForSpell(cast.CastTime, spell); warrior.Talents.ImprovedSlam == 0 && castTime > 0 {
					warrior.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime+castTime, true)
				}
			},
		},

		CritDamageBonus: warrior.impale(),

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		FlatThreatBonus:  140, // Should this be 54 or the old 140 value from before SoD?
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := flatDamageBonus + spell.Unit.MHWeaponDamage(sim, spell.MeleeAttackPower(target))

			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			if !result.Landed() {
				spell.IssueRefund(sim)
			}
		},
	})
}
