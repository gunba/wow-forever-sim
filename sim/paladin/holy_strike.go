package paladin

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Holy Strike is new in Forever and baseline from level 6. Iron Creed sharpens its
// threat and Sacred Arbiter its damage; the former cooldown talent is now baseline.
//
// Beta client 1.60.1.70009 gives it eight ranks on ids Classic used for NPC spells. Each is a
// normalized weapon strike plus a flat amount (effect 121) and a weapon damage percentage (effect 31),
// with a 10 sec cooldown and the 0.429 coefficient. The client multiplies the flat
// amount by the percentage the same way it does Backstab's, so rank 8 is 50% of
// (weapon + 81 to 105).
// Damage is each rank's at its max level.
// TODO: beta will confirm - Holy damage on the melee hit table, so it rolls partial resists the
// way every other Holy ability here does. Whether a melee-table Holy strike actually partial
// resists is unknown; if it does not, it wants SpellFlagIgnoreResists.
const holyStrikeCooldown = time.Second * 10

func (paladin *Paladin) registerHolyStrike() {
	cd := core.Cooldown{
		Timer:    paladin.NewTimer(),
		Duration: holyStrikeCooldown,
	}

	// Sacred Arbiter also refreshes the paladin's Judgement effects. Judgement of the Crusader
	// is the only Judgement that leaves anything behind, and it already refreshes off every
	// melee attack the paladin lands, so that half of the talent needs nothing here.
	damageMultiplier := paladin.getWeaponSpecializationModifier()
	if paladin.Talents.SacredArbiter {
		damageMultiplier *= 1.2
	}

	// 5% per rank, confirmed on the beta at ranks 2, 3 and 4: 10%, 15% and 20%.
	threatMultiplier := 1 + 0.05*float64(paladin.Talents.IronCreed)

	ironCreedAura := paladin.registerIronCreedAura()

	ranks := []struct {
		level     int32
		manaCost  float64
		weapon    float64
		minDamage float64
		maxDamage float64
	}{
		{level: 6, manaCost: 5, weapon: 0.25, minDamage: 11, maxDamage: 14},
		{level: 12, manaCost: 9, weapon: 0.29, minDamage: 15, maxDamage: 20},
		{level: 20, manaCost: 12, weapon: 0.32, minDamage: 17, maxDamage: 23},
		{level: 28, manaCost: 14, weapon: 0.36, minDamage: 22, maxDamage: 29},
		{level: 36, manaCost: 16, weapon: 0.39, minDamage: 32, maxDamage: 40},
		{level: 44, manaCost: 17, weapon: 0.43, minDamage: 53, maxDamage: 68},
		{level: 52, manaCost: 19, weapon: 0.46, minDamage: 73, maxDamage: 91},
		{level: 60, manaCost: 20, weapon: 0.50, minDamage: 81, maxDamage: 105},
	}

	for i, rank := range ranks {
		rank := rank
		spellID := []int32{679, 678, 1866, 680, 2495, 5569, 10332, 10333}[i]
		if paladin.Level < rank.level {
			break
		}

		paladin.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: spellID},
			SpellCode:   SpellCode_PaladinHolyStrike,
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskMeleeMHSpecial,
			Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

			RequiredLevel: int(rank.level),
			Rank:          i + 1,

			ManaCost: core.ManaCostOptions{
				FlatCost:   rank.manaCost,
				Multiplier: paladin.benediction(),
			},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
				IgnoreHaste: true,
				CD:          cd,
			},

			DamageMultiplier: damageMultiplier,
			BonusCritRating: float64(paladin.Talents.HolyPower) *
				core.CritRatingPerCritChance * 3,
			ThreatMultiplier: threatMultiplier,
			// Holy damage, so spell power feeds it on top of the weapon share and the flat roll.
			BonusCoefficient: 0.429,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				if ironCreedAura != nil {
					ironCreedAura.Activate(sim)
				}

				// A share of weapon damage, so it takes the normalized swing the way every other
				// percentage-of-weapon strike in the sim does.
				baseDamage := rank.weapon * (spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target)) +
					sim.Roll(rank.minDamage, rank.maxDamage))
				spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)
			},
		})
	}
}

// The half of Iron Creed that is not threat: Holy Strike shaves the damage the paladin takes,
// but only while Righteous Fury is up. Ardent Defender's spell id stands in for the buff.
func (paladin *Paladin) registerIronCreedAura() *core.Aura {
	if paladin.Talents.IronCreed == 0 || !paladin.Options.RighteousFury {
		return nil
	}

	// 2% per rank, confirmed on the beta at ranks 2, 3 and 4: 4%, 6% and 8%. The 6 seconds
	// is flat at every rank, which those same tooltips show.
	damageTaken := 1 - 0.02*float64(paladin.Talents.IronCreed)

	return paladin.RegisterAura(core.Aura{
		Label:    "Iron Creed",
		ActionID: core.ActionID{SpellID: 31850},
		Duration: time.Second * 6,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			paladin.PseudoStats.DamageTakenMultiplier *= damageTaken
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			paladin.PseudoStats.DamageTakenMultiplier /= damageTaken
		},
	})
}
