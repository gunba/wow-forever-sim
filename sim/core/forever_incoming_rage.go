package core

import (
	"fmt"
	"math"

	"github.com/wowsims/classic/sim/core/proto"
)

// ForeverIncomingRageParameters holds the optional manual incoming-Rage scenario. With no override the sim uses the
// rule measured on Forever combat logs (foreverMeasuredRage); any override switches to the older provisional
// formula (coefficient 10 / expected creature health / reference Armor), kept for sensitivity runs.
// Copied resolved values remain stable when health/Stamina changes during a fight.
type ForeverIncomingRageParameters struct {
	PlayerLevel            int32
	Coefficient            float64
	ReferenceArmor         float64
	ExpectedHealth         float64
	CoefficientOverride    bool
	ReferenceArmorOverride bool
	ExpectedHealthOverride bool
}

const ForeverIncomingRageAssumptions = "Measured (wowsims/forever #356, Forever combat logs): Rage = 2.5 x (1 - A) x pre-mitigation damage / c(attacker level), c(L) = 0.0091107836 L^2 + 3.225598133 L + 4.2652911, A = 20% up to attacker level 42, 30% at 44, 36% at 46, 38% at 50, 40% at 63 (interpolated; 63 unmeasured). Pre-mitigation damage includes crits (x2) and crushing blows (x1.5), before armor, blocks and absorbs. A manual override switches to the provisional coefficient-10 formula."

func validateForeverModelNumber(name string, value, minimum, maximum float64) error {
	if math.IsNaN(value) || math.IsInf(value, 0) || value < minimum || value > maximum {
		return fmt.Errorf("%s must be finite and between %g and %g, got %g", name, minimum, maximum, value)
	}
	return nil
}

// ResolveForeverIncomingRageModel retains optional scalar presence: an explicit
// coefficient zero disables the damage-based component, while absence chooses 10.
// Validation bounds protect scenario inputs; they are not claims about game values.
func ResolveForeverIncomingRageModel(level int32, model *proto.ForeverIncomingRageModel) (ForeverIncomingRageParameters, error) {
	row, ok := foreverExpectedStatAtLevel(level)
	if !ok {
		return ForeverIncomingRageParameters{}, fmt.Errorf("provisional incoming Rage requires a player level between 1 and 60, got %d", level)
	}
	parameters := ForeverIncomingRageParameters{
		PlayerLevel:    level,
		Coefficient:    10,
		ReferenceArmor: max(.20, min(.40, row.creatureArmor/(row.creatureArmor+row.armorConstant))),
		ExpectedHealth: row.creatureHealth,
	}
	if model == nil {
		return parameters, nil
	}
	if model.Coefficient != nil {
		if err := validateForeverModelNumber("provisional incoming Rage coefficient", *model.Coefficient, 0, 1000); err != nil {
			return ForeverIncomingRageParameters{}, err
		}
		parameters.Coefficient = *model.Coefficient
		parameters.CoefficientOverride = true
	}
	if model.ReferenceArmorOverride != nil {
		if err := validateForeverModelNumber("provisional incoming Rage reference Armor override", *model.ReferenceArmorOverride, 0, .95); err != nil {
			return ForeverIncomingRageParameters{}, err
		}
		parameters.ReferenceArmor = *model.ReferenceArmorOverride
		parameters.ReferenceArmorOverride = true
	}
	if model.ExpectedHealthOverride != nil {
		if err := validateForeverModelNumber("provisional incoming Rage expected health override", *model.ExpectedHealthOverride, 1, 1e9); err != nil {
			return ForeverIncomingRageParameters{}, err
		}
		parameters.ExpectedHealth = *model.ExpectedHealthOverride
		parameters.ExpectedHealthOverride = true
	}
	return parameters, nil
}

// Actual Armor, resistance and other mitigation are already included in Damage.
// AbsorbedDamage restores damage delivered to a shield without stripping the
// critical/crushing outcome. Never undo actual Armor or use player maximum health.
// manual reports whether the scenario overrides the measured rule with the provisional formula.
func (parameters ForeverIncomingRageParameters) manual() bool {
	return parameters.CoefficientOverride || parameters.ReferenceArmorOverride || parameters.ExpectedHealthOverride
}

// ForeverAssumedArmorReduction is the armor reduction Forever assumes for Rage from damage taken, by attacker level:
// the defender's own armor, level and health play no part (wowsims/forever #356, fitted to combat logs).
func ForeverAssumedArmorReduction(attackerLevel int32) float64 {
	points := [][2]float64{{42, .20}, {44, .30}, {46, .36}, {50, .38}, {63, .40}}
	l := float64(attackerLevel)
	if l <= points[0][0] {
		return points[0][1]
	}
	for i := 1; i < len(points); i++ {
		if l <= points[i][0] {
			a, b := points[i-1], points[i]
			return a[1] + (b[1]-a[1])*(l-a[0])/(b[0]-a[0])
		}
	}
	return points[len(points)-1][1]
}

// ForeverRagePerPreMitigationDamage is the Rage a point of pre-mitigation damage taken from an attacker of this level
// gives: 2.5 x (1 - A) / c(attacker level).
func ForeverRagePerPreMitigationDamage(attackerLevel int32) float64 {
	l := float64(attackerLevel)
	c := 0.0091107836*l*l + 3.225598133*l + 4.2652911
	return 2.5 * (1 - ForeverAssumedArmorReduction(attackerLevel)) / c
}

// foreverMeasuredRage: the hit with its outcome (crit, crushing blow) before armor, blocks and absorbs.
func foreverMeasuredRage(spell *Spell, result *SpellResult) float64 {
	if !result.Landed() || result.ResistanceMultiplier <= 0 {
		return 0
	}
	damage := result.RawDamage()
	switch {
	case result.Outcome.Matches(OutcomeCrit):
		damage *= spell.CritMultiplier(spell.Unit.AttackTables[result.Target.UnitIndex][spell.CastType])
	case result.Outcome.Matches(OutcomeCrush):
		damage *= 1.5
	}
	return damage * ForeverRagePerPreMitigationDamage(spell.Unit.Level)
}

func (parameters ForeverIncomingRageParameters) rageFromDamageTaken(result *SpellResult) float64 {
	damage := result.Damage + result.AbsorbedDamage
	if damage <= 0 || parameters.Coefficient == 0 {
		return 0
	}
	return damage / parameters.ExpectedHealth * parameters.Coefficient / (1 - parameters.ReferenceArmor)
}

// ForeverIncomingRageParameters returns a copy for scenario/metadata reporting.
func (unit *Unit) ForeverIncomingRageParameters() ForeverIncomingRageParameters {
	return unit.foreverIncomingRageParameters
}
