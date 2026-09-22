package paladin

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func (paladin *Paladin) registerExorcism() {
	// Beta client 1.60.1.69893: every rank hits a little softer (rank 6 505-563 -> 474-530); cost,
	// cooldown, per level growth and the 0.429 coefficient are Classic's.
	ranks := []struct {
		level      int32
		manaCost   float64
		scaleLevel int32
		minDamage  float64
		maxDamage  float64
		scale      float64
	}{
		{level: 20, manaCost: 85, scaleLevel: 25, minDamage: 73, maxDamage: 85, scale: 1.2},
		{level: 28, manaCost: 135, scaleLevel: 33, minDamage: 132, maxDamage: 150, scale: 1.6},
		{level: 36, manaCost: 180, scaleLevel: 41, minDamage: 189, maxDamage: 215, scale: 2.0},
		{level: 44, manaCost: 235, scaleLevel: 49, minDamage: 273, maxDamage: 309, scale: 2.4},
		{level: 52, manaCost: 285, scaleLevel: 57, minDamage: 362, maxDamage: 406, scale: 2.8},
		{level: 60, manaCost: 345, scaleLevel: 60, minDamage: 474, maxDamage: 530, scale: 3.2},
	}

	timer := paladin.NewTimer() // All ranks share client cooldown category 19.
	for i, rank := range ranks {
		rank := rank
		spellID := []int32{879, 5614, 5615, 10312, 10313, 10314}[i]
		if paladin.Level < rank.level {
			break
		}

		minDamage := rank.minDamage + float64(min(paladin.Level, rank.scaleLevel)-rank.level)*rank.scale
		maxDamage := rank.maxDamage + float64(min(paladin.Level, rank.scaleLevel)-rank.level)*rank.scale

		spell := paladin.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: spellID},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,
			Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | core.SpellFlagBinary, //Logs show it never has partial resists, No clue why, still misses

			RequiredLevel: int(rank.level),
			Rank:          i + 1,

			SpellCode: SpellCode_PaladinExorcism,
			ManaCost: core.ManaCostOptions{
				FlatCost:   rank.manaCost,
				Multiplier: paladin.benediction() * paladin.holyConduit() / 100,
			},

			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
				CD: core.Cooldown{
					Timer:    timer,
					Duration: paladin.purifyingPower(time.Second * 15),
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BonusCoefficient: 0.429,

			ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
				return target.MobType == proto.MobType_MobTypeDemon || target.MobType == proto.MobType_MobTypeUndead
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(minDamage, maxDamage), spell.OutcomeMagicHitAndCrit)
			},
		})

		paladin.exorcism = append(paladin.exorcism, spell)
	}
}
