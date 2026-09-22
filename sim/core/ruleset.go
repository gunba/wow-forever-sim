package core

import (
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// Simulation embeds Environment, so sim.IsForever() resolves here too.
func (env *Environment) IsForever() bool {
	return env.Ruleset == proto.Ruleset_RulesetForever
}

// Periodic damage rolls for crits under the Forever ruleset. Spells that
// should keep ticking for flat damage opt out with SpellFlagNoPeriodicCrit.
func (dot *Dot) canCrit(sim *Simulation) bool {
	return sim.IsForever() && !dot.Spell.Flags.Matches(SpellFlagNoPeriodicCrit)
}

// Ticks roll against the caster's crit chance at the time of the tick instead of the
// chance snapshotted when the dot went up. Dots that are applied by hand, like Deep
// Wounds, never snapshot one at all, so rolling live is also the only way for them to
// crit at the right rate.
func (dot *Dot) critCheck(sim *Simulation, target *Unit, attackTable *AttackTable) bool {
	if dot.Spell.SchoolIndex == stats.SchoolIndexPhysical {
		return dot.Spell.PhysicalCritCheck(sim, attackTable)
	}
	return dot.Spell.MagicCritCheck(sim, target)
}

func (character *Character) itemStats(item Item, includeEnchant bool) stats.Stats {
	parts := []stats.Stats{item.Stats, item.RandomSuffix.Stats}
	if includeEnchant {
		parts = append(parts, item.Enchant.Stats)
	}
	var total stats.Stats
	for _, part := range parts {
		if character.Env.IsForever() {
			part = character.unifyEquipHitAndCrit(part)
		}
		total = total.Add(part)
	}
	return total
}

// Forever pays out hit and critical strike from gear against every kind of attack
// rather than splitting them into a melee and a spell pool. Attribute conversions are
// untouched: only the hit and crit an item spells out become universal.
func (character *Character) unifyEquipHitAndCrit(equipStats stats.Stats) stats.Stats {
	hit := equipStats[stats.MeleeHit] + equipStats[stats.SpellHit]
	crit := equipStats[stats.MeleeCrit] + equipStats[stats.SpellCrit]

	equipStats[stats.MeleeHit] = hit
	equipStats[stats.SpellHit] = hit
	equipStats[stats.MeleeCrit] = crit
	equipStats[stats.SpellCrit] = crit

	return equipStats
}
