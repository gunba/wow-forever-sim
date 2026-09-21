package core

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// ForeverTier1SetID maps the simulator's DPS roles to the beta ItemSet records.
// Smite uses the Holy/Discipline set, not the Shadow set.
func ForeverTier1SetID(spec proto.Spec) int32 {
	switch spec {
	case proto.Spec_SpecMage:
		return 2098
	case proto.Spec_SpecRogue:
		return 2099
	case proto.Spec_SpecWarlock:
		return 2100
	case proto.Spec_SpecHunter:
		return 2101
	case proto.Spec_SpecWarrior:
		return 2102
	case proto.Spec_SpecSmitePriest:
		return 2104
	case proto.Spec_SpecShadowPriest:
		return 2105
	case proto.Spec_SpecRetributionPaladin:
		return 2106
	case proto.Spec_SpecEnhancementShaman:
		return 2110
	case proto.Spec_SpecElementalShaman:
		return 2111
	case proto.Spec_SpecBalanceDruid:
		return 2114
	case proto.Spec_SpecFeralDruid:
		return 2115
	default:
		return 0
	}
}

func (character *Character) forcedForeverTier1SetID() int32 {
	if !character.ForeverTier1Bonuses || character.Env == nil || !character.Env.IsForever() {
		return 0
	}
	return ForeverTier1SetID(character.Spec)
}

func ForeverTier1HitBonus(agent Agent) {
	agent.GetCharacter().AddStats(stats.Stats{
		stats.MeleeHit: MeleeHitRatingPerHitChance,
		stats.SpellHit: SpellHitRatingPerHitChance,
	})
}

func ForeverTier1HasteBonus(agent Agent) {
	c := agent.GetCharacter()
	c.PseudoStats.MeleeSpeedMultiplier *= 1.01
	c.PseudoStats.RangedSpeedMultiplier *= 1.01
	c.PseudoStats.CastSpeedMultiplier *= 1.01
	c.PseudoStats.EnergyHasteMultiplier *= 1.01
}

func ForeverTier1CooldownBonus(reduction time.Duration, matches func(*Spell) bool) ApplyEffect {
	return func(agent Agent) {
		agent.GetCharacter().OnSpellRegistered(func(spell *Spell) {
			if matches(spell) && spell.CD.Timer != nil {
				spell.CD.Duration = max(0, spell.CD.Duration-reduction)
			}
		})
	}
}

// The creature restriction belongs on each attack table, not on the player's
// global stats: mixed encounters must not grant the bonus against other targets.
func ForeverTier1CreatureBonus(spellID int32, mobType proto.MobType, attackPower, spellPower float64) ApplyEffect {
	return func(agent Agent) {
		c := agent.GetCharacter()
		apply := func(multiplier float64) {
			for _, target := range c.Env.Encounter.TargetUnits {
				if target.MobType != mobType {
					continue
				}
				for _, table := range c.AttackTables[target.UnitIndex] {
					table.BonusAttackPowerTaken += multiplier * attackPower
					table.BonusSpellDamageTaken += multiplier * spellPower
				}
			}
		}
		MakePermanent(c.RegisterAura(Aura{
			Label:    fmt.Sprintf("Tier 1 creature bonus %d", spellID),
			ActionID: ActionID{SpellID: spellID},
			OnGain:   func(_ *Aura, _ *Simulation) { apply(1) },
			OnExpire: func(_ *Aura, _ *Simulation) { apply(-1) },
		}))
	}
}
