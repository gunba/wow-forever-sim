package paladin

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (paladin *Paladin) registerHolyShock() {
	if !paladin.Talents.HolyShock {
		return
	}

	// Beta client 1.60.1.69893 adds a rank 1 at level 30 (1311606, damage spell 1311604), which
	// renumbers Classic's three ranks 2-4, lowers their damage (rank 4 365-395 -> 334-362) and cuts
	// the cooldown from 30 sec to 10. Costs and the 0.429 are Classic's.
	ranks := []struct {
		level     int32
		manaCost  float64
		minDamage float64
		maxDamage float64
	}{
		{level: 30, manaCost: 160, minDamage: 128, maxDamage: 140},
		{level: 40, manaCost: 225, minDamage: 175, maxDamage: 189},
		{level: 48, manaCost: 275, minDamage: 248, maxDamage: 268},
		{level: 56, manaCost: 325, minDamage: 334, maxDamage: 362},
	}

	timer := paladin.NewTimer() // All ranks share client cooldown category 892.
	for i, rank := range ranks {
		rank := rank
		spellID := []int32{1311606, 20473, 20929, 20930}[i]
		if paladin.Level < rank.level {
			break
		}

		paladin.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: spellID},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,
			Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

			RequiredLevel: int(rank.level),
			Rank:          i + 1,

			SpellCode: SpellCode_PaladinHolyShock,

			ManaCost: core.ManaCostOptions{
				FlatCost:   rank.manaCost,
				Multiplier: paladin.benediction(),
			},

			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
				CD: core.Cooldown{
					Timer:    timer,
					Duration: time.Second * 10,
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: 0.429,

			// Holy Power is worth an extra 2% crit per point on Holy Shock specifically. It stacks
			// on top of the 1% per point every spell gets in ApplyTalents, which is how Holy Shock
			// reaches the 3% per point the tree reads for it.
			BonusCritRating: 2 * float64(paladin.Talents.HolyPower) * core.SpellCritRatingPerCritChance,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				baseDamage := sim.Roll(rank.minDamage, rank.maxDamage)
				spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
			},
		})
	}
}
