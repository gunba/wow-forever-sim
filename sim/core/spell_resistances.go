package core

import (
	"math"

	"github.com/wowsims/classic/sim/core/stats"
)

func (result *SpellResult) applyResistances(sim *Simulation, spell *Spell, isPeriodic bool, attackTable *AttackTable) {
	resistanceMultiplier, outcome := spell.ResistanceMultiplier(sim, isPeriodic, attackTable)

	result.Damage *= resistanceMultiplier
	result.Outcome |= outcome

	result.ResistanceMultiplier = resistanceMultiplier
	result.PreOutcomeDamage = result.Damage
}

// Modifies damage based on Armor or Magic resistances, depending on the damage type.
func (spell *Spell) ResistanceMultiplier(sim *Simulation, isPeriodic bool, attackTable *AttackTable) (float64, HitOutcome) {
	if spell.Flags.Matches(SpellFlagIgnoreResists) {
		return 1, OutcomeEmpty
	}

	if spell.SpellSchool.Matches(SpellSchoolPhysical) {
		if spell.SchoolIndex == stats.SchoolIndexPhysical || MultiSchoolShouldUseArmor(spell, attackTable.Defender) {
			// All physical dots (Bleeds) ignore armor.
			if isPeriodic {
				return 1, OutcomeEmpty
			}

			// Physical resistance (armor).
			return attackTable.GetArmorDamageModifier(), OutcomeEmpty
		}
	}

	// Magical resistance. An expected-damage call takes no roll: finalizeExpectedDamage applies the average resist.
	if spell.expecting {
		return 1, OutcomeEmpty
	}
	if attackTable.Attacker.Env != nil && attackTable.Attacker.Env.IsForever() {
		if spell.Flags.Matches(SpellFlagNoPartialResist) {
			return 1, OutcomeEmpty
		}
		return attackTable.Defender.foreverPartialResist(sim, spell, attackTable.Attacker, isPeriodic && !spell.Flags.Matches(SpellFlagChanneled))
	}

	if spell.Flags.Matches(SpellFlagBinary) {
		return 1, OutcomeEmpty
	}

	resistanceRoll := sim.RandomFloat("Partial Resist")

	threshold00, threshold25, threshold50 := attackTable.GetPartialResistThresholds(spell, spell.Flags.Matches(SpellFlagPureDot))
	//if sim.Log != nil {
	//	sim.Log("Resist thresholds: %0.04f, %0.04f, %0.04f", threshold00, threshold25, threshold50)
	//}

	if resistanceRoll > threshold00 {
		// No partial resist.
		return 1, OutcomeEmpty
	} else if resistanceRoll > threshold25 {
		return 0.75, OutcomePartial1_4
	} else if resistanceRoll > threshold50 {
		return 0.5, OutcomePartial2_4
	} else {
		return 0.25, OutcomePartial3_4
	}
}

// Decide whether to use armor for physical multi school spells.
//
// TODO: This is most likely not accurate for the case: armor near resistance but not 0
//
// A short test showed that the game uses armor if it's far enough below resistance,
// but not simply if it's lower.
// 49 (and above) armor vs 57 res => used resistance
// 7 (and below) armor vs 57 res => used armor/no partials anymore
// If level based resist is used in this decission process is also not known as it was tested PvP.
//
// For most purposes this should work fine for now, but should be properly tested and fixed if
// spells using it become important and boss armor can actually go below (level based) resistance values.
func MultiSchoolShouldUseArmor(spell *Spell, target *Unit) bool {
	resistance := 100000.0
	lowestIsArmor := true
	for _, baseSchoolIndex := range spell.SchoolBaseIndices {
		resiVal := target.GetResistanceForSchool(baseSchoolIndex)
		if resiVal < resistance {
			resistance = resiVal
			lowestIsArmor = baseSchoolIndex == stats.SchoolIndexPhysical
		}
	}
	return lowestIsArmor
}

func (at *AttackTable) GetArmorDamageModifier() float64 {
	armorPenRating := at.Attacker.stats[stats.ArmorPenetration]
	// The share the attacker ignores comes off first, then the flat rating.
	ignoredArmor := at.Defender.Armor() * at.Attacker.PseudoStats.ArmorIgnorePercent
	defenderArmor := max(at.Defender.Armor()-ignoredArmor-armorPenRating, 0.0)
	return 1 - defenderArmor/(defenderArmor+400+85*float64(at.Attacker.Level))
}

func (at *AttackTable) GetPartialResistThresholds(spell *Spell, pureDot bool) (float64, float64, float64) {
	return at.Defender.partialResistRollThresholds(spell, at.Attacker, pureDot)
}

func (at *AttackTable) GetBinaryHitChance(spell *Spell) float64 {
	return at.Defender.binaryHitChance(spell, at.Attacker)
}

// Only for base schools!
func (unit *Unit) GetResistanceForSchool(schoolIndex stats.SchoolIndex) float64 {
	switch schoolIndex {
	case stats.SchoolIndexNone:
		return 0
	case stats.SchoolIndexPhysical:
		return unit.GetStat(stats.Armor)
	case stats.SchoolIndexArcane:
		return unit.GetStat(stats.ArcaneResistance)
	case stats.SchoolIndexFire:
		return unit.GetStat(stats.FireResistance)
	case stats.SchoolIndexFrost:
		return unit.GetStat(stats.FrostResistance)
	case stats.SchoolIndexHoly:
		return 0 // Holy resistance doesn't exist.
	case stats.SchoolIndexNature:
		return unit.GetStat(stats.NatureResistance)
	case stats.SchoolIndexShadow:
		return unit.GetStat(stats.ShadowResistance)
	default:
		return 0
	}
}

// All of the following calculations are based on this guide:
// https://royalgiraffe.github.io/resist-guide

func (unit *Unit) resistCoeff(spell *Spell, attacker *Unit, binary bool, pureDot bool) float64 {
	if spell.SchoolIndex <= stats.SchoolIndexPhysical {
		return 0
	}

	var resistance float64

	if spell.SchoolIndex.IsMultiSchool() {
		// Multi school: Choose the lowest resistance available.
		resistance = 1000.0
		for _, baseSchoolIndex := range spell.SchoolBaseIndices {
			resiVal := unit.GetResistanceForSchool(baseSchoolIndex)
			if resiVal < resistance {
				resistance = resiVal
			}
		}
	} else {
		resistance = unit.GetResistanceForSchool(spell.SchoolIndex)
	}

	resistance = max(0, resistance-attacker.stats[stats.SpellPenetration])

	resistanceCap := float64(attacker.Level * 5)
	resistanceCoef := resistance / resistanceCap

	// Pre-TBC all dots that don't have an initial damage component
	// use a 1/10 of the resistance score
	if pureDot {
		resistanceCoef /= 10
	}

	if !binary && unit.Type == EnemyUnit && unit.Level > attacker.Level {
		avgMitigationAdded := AverageMagicPartialResistPerLevelMultiplier * float64(unit.Level-attacker.Level)
		// coef is scaled 0 to 1, not 0 to 0.75
		resistanceCoef += avgMitigationAdded * 1 / 0.75
	}

	return min(1, resistanceCoef)
}

func (unit *Unit) binaryHitChance(spell *Spell, attacker *Unit) float64 {
	if attacker.Env != nil && attacker.Env.IsForever() {
		// A binary spell is resisted in full at the average resist (wowsims/forever #293).
		return 1 - max(0, unit.foreverAverageResist(spell, attacker, false))
	}
	resistCoeff := unit.resistCoeff(spell, attacker, true, false)
	return 1 - 0.75*resistCoeff
}

// ForeverResistanceConstant is K in a spell's average resist R / (|R| + K), per caster level. Forever
// combat logs with fixed Spell Piercing at caster levels 20-30 fit 9L (wowsims/forever #436); the logs
// show no level-based resist against higher-level targets at any gap (#350).
const ForeverResistanceConstant = 9

// foreverAverageResist is the average share of a magic hit resisted. R is the target's resistance, which
// stops at 0 (a debuff that lowers it further adds nothing), less the caster's Spell Piercing; piercing past
// the resistance makes R negative and adds damage (#293). A DoT tick, channels aside, reads a tenth of the
// share (#288).
func (unit *Unit) foreverAverageResist(spell *Spell, attacker *Unit, tick bool) float64 {
	if spell.SchoolIndex <= stats.SchoolIndexPhysical {
		return 0
	}
	var resistance float64
	if spell.SchoolIndex.IsMultiSchool() {
		resistance = 1000.0
		for _, baseSchoolIndex := range spell.SchoolBaseIndices {
			resistance = min(resistance, unit.GetResistanceForSchool(baseSchoolIndex))
		}
	} else {
		resistance = unit.GetResistanceForSchool(spell.SchoolIndex)
	}
	resistance = max(0, resistance) - attacker.stats[stats.SpellPenetration]
	k := ForeverResistanceConstant * float64(attacker.Level)
	if k <= 0 {
		return 0
	}
	average := resistance / (math.Abs(resistance) + k)
	if tick {
		average /= 10
	}
	return average
}

// foreverPartialResist rolls a magic hit's partial resist in 10% steps (Forever logs: every partial on
// a 10% step, never 25/50/75). From an average a of 0.1 or more the steps follow the triangle
// P(x) = 0.5 - 2.5|x - a|; below it the roll is [1 - 7.5a, 5a, 2.5a] over 0/10/20%. A negative a
// (Spell Piercing past the target's resistance) is the same roll with the steps adding damage.
func (unit *Unit) foreverPartialResist(sim *Simulation, spell *Spell, attacker *Unit, tick bool) (float64, HitOutcome) {
	a := unit.foreverAverageResist(spell, attacker, tick)
	if a == 0 {
		return 1, OutcomeEmpty
	}
	sign := 1.0
	if a < 0 {
		sign, a = -1, -a
	}
	a = min(a, 1)
	roll := sim.RandomFloat("Partial Resist")
	steps := 0
	if a < 0.1 {
		switch {
		case roll < 2.5*a:
			steps = 2
		case roll < 7.5*a:
			steps = 1
		}
	} else {
		cumulative := 0.0
		for k := 0; k <= 10; k++ {
			x := 0.1 * float64(k)
			cumulative += max(0, 0.5-2.5*math.Abs(x-a))
			if roll < cumulative {
				steps = k
				break
			}
			steps = k
		}
	}
	if steps == 0 {
		return 1, OutcomeEmpty
	}
	multiplier := 1 - sign*0.1*float64(steps)
	if sign < 0 {
		return multiplier, OutcomeEmpty
	}
	switch {
	case steps <= 3:
		return multiplier, OutcomePartial1_4
	case steps <= 6:
		return multiplier, OutcomePartial2_4
	default:
		return multiplier, OutcomePartial3_4
	}
}

// Roll threshold for each type of partial resist.
func (unit *Unit) partialResistRollThresholds(spell *Spell, attacker *Unit, pureDot bool) (float64, float64, float64) {
	resistCoeff := unit.resistCoeff(spell, attacker, false, pureDot)

	// Based on the piecewise linear regression estimates at https://royalgiraffe.github.io/partial-resist-table.
	if val := resistCoeff * 3; val <= 1 {
		return 0.76 * val, 0.21 * val, 0.03 * val
	} else if val <= 2 {
		val -= 1
		return 0.76 + 0.24*val, 0.21 + 0.57*val, 0.03 + 0.19*val
	} else {
		val -= 2
		return 1, 0.78 + 0.18*val, 0.22 + 0.58*val
	}
}
