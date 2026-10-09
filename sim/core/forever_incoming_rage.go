package core

import (
	"fmt"
	"math"

	"github.com/wowsims/classic/sim/core/proto"
)

// ForeverIncomingRageParameters is an explicitly provisional scenario, not a
// decoded or verified server rule. October 8's official notes establish expected
// creature health, shield-independent damage, and a varying reference Armor;
// coefficient 10, player-level lookup, this curve and its operation are assumptions.
// https://us.forums.blizzard.com/en/wow/t/2360696/5
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

const ForeverIncomingRageAssumptions = "Provisional: coefficient 10; player-level ExpectedStat CreatureHealth; reference Armor clamp(CreatureArmor/(CreatureArmor+ArmorConstant), 0.20, 0.40); post-actual-mitigation damage including absorbs divided by (1-reference Armor). Not verified server behavior."

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
