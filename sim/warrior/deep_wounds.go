package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func (warrior *Warrior) applyDeepWounds() {
	if warrior.Talents.DeepWounds == 0 {
		return
	}

	spellID := map[int32]int32{
		1: 12834,
		2: 12849,
		3: 12867,
	}[warrior.Talents.DeepWounds]

	flags := core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | core.SpellFlagNoPeriodicCrit
	if warrior.Env.IsForever() {
		// Child 412613 excludes caster damage modifiers. This is a weapon
		// payload, not damage copied from the critical strike.
		flags |= core.SpellFlagIgnoreAttackerModifiers
	}

	warrior.DeepWounds = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		SpellCode:   SpellCode_WarriorDeepWounds,
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskEmpty,
		// The triggered Forever bleed child 412613 is explicitly unable to
		// crit. The blanket Forever periodic rule must not grant it crits.
		Flags: flags,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: core.TernaryFloat64(warrior.Env.IsForever(), 0, 1),

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Deep Wounds",
			},
			NumberOfTicks: 4,
			TickLength:    time.Second * 3,

			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				if sim.IsForever() {
					// The rolled, unpaid weapon payload is the exception to ordinary
					// dynamic DoTs. Target modifiers still apply once at each tick.
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, dot.SnapshotBaseDamage, dot.OutcomeTick)
					return
				}
				attackTable := warrior.AttackTables[target.UnitIndex][proto.CastType_CastTypeMainHand]
				dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(attackTable, true) // Double dips on attackers mods
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			if sim.IsForever() {
				spell.Dot(target).ApplyOrRefresh(sim)
			} else {
				spell.Dot(target).Apply(sim)
			}
			spell.CalcAndDealOutcome(sim, target, spell.OutcomeAlwaysHitNoHitCounter)
		},
	})

	core.MakePermanent(warrior.RegisterAura(core.Aura{
		Label: "Deep Wounds Talent",
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskEmpty) || !spell.SpellSchool.Matches(core.SpellSchoolPhysical) {
				return
			}

			// Ravager doesn't proc Deep Wounds
			if spell.ActionID.SpellID == 9633 {
				return
			}

			if result.Outcome.Matches(core.OutcomeCrit) {
				warrior.procDeepWounds(sim, result.Target, spell.IsOH())
			}
		},
	}))
}

func (warrior *Warrior) procDeepWounds(sim *core.Simulation, target *core.Unit, isOh bool) {
	dot := warrior.DeepWounds.Dot(target)

	if sim.IsForever() {
		outstanding := 0.0
		if dot.IsActive() {
			outstanding = dot.SnapshotBaseDamage * float64(dot.MaxTicksRemaining())
		}
		// Low-level measurements support the main-hand weapon-only payload.
		// Preserve that reference until off-hand attribution is measured.
		newDamage := warrior.AutoAttacks.MH().AverageDamage() * .2 * float64(warrior.Talents.DeepWounds)
		dot.SnapshotBaseDamage = (outstanding + newDamage) / float64(dot.NumberOfTicks)
		dot.SnapshotAttackerMultiplier = 1
		warrior.DeepWounds.Cast(sim, target)
		return
	}

	var awd float64
	if isOh {
		attackTableOh := warrior.AttackTables[target.UnitIndex][proto.CastType_CastTypeOffHand]
		adm := warrior.AutoAttacks.OHAuto().AttackerDamageMultiplier(attackTableOh, true)
		awd = warrior.AutoAttacks.OH().CalculateAverageWeaponDamage(dot.Spell.MeleeAttackPower(target)) * 0.5 * adm
	} else { // MH
		attackTableMh := warrior.AttackTables[target.UnitIndex][proto.CastType_CastTypeMainHand]
		adm := warrior.AutoAttacks.MHAuto().AttackerDamageMultiplier(attackTableMh, true)
		awd = warrior.AutoAttacks.MH().CalculateAverageWeaponDamage(dot.Spell.MeleeAttackPower(target)) * adm
	}

	newDamage := awd * 0.2 * float64(warrior.Talents.DeepWounds) // 60% of average attackers damage

	dot.SnapshotBaseDamage = newDamage / 4.0 // spread over 4 ticks of the dot
	dot.SnapshotAttackerMultiplier = 1

	warrior.DeepWounds.Cast(sim, target)
}
