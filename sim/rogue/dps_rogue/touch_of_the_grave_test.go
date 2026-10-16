package dpsrogue

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	googleProto "google.golang.org/protobuf/proto"
)

// The undead's Forever racial replaced Shadow Resistance with Touch of the Grave, and the
// sim carried the removal without the replacement, so an undead player had no racial at
// all and nothing in the log to show for it (#145). Only Orc and Human rogues are covered
// by the DPS suites, which is why it went unnoticed; this asserts the racial actually
// fires and is attributed to its own spell id.
func TestTouchOfTheGraveIsLogged(t *testing.T) {
	player := &proto.Player{
		Class:         proto.Class_ClassRogue,
		Race:          proto.Race_RaceUndead,
		Equipment:     core.GetGearSet("../../../ui/rogue/gear_sets", "combat_sinister_strike_prebis").GearSet,
		TalentsString: CombatSwordsTalents,
		Rotation:      core.GetAplRotation("../../../ui/rogue/apls", "combat_sinister_strike").Rotation,
		Consumes:      Phase1Consumes.Consumes,
		Spec:          DefaultRogue,
	}

	// The racial only exists under the Forever ruleset; under Classic the undead keep
	// Shadow Resistance instead, so the options have to say which game this is.
	simOptions := googleProto.Clone(core.DefaultSimTestOptions).(*proto.SimOptions)
	simOptions.Ruleset = proto.Ruleset_RulesetForever

	request := &proto.RaidSimRequest{
		Raid:       core.SinglePlayerRaidProto(player, nil, core.ForeverBuffs.Raid, core.ForeverBuffs.Debuffs),
		Encounter:  core.MakeSingleTargetEncounter(0),
		SimOptions: simOptions,
	}

	result := core.RunRaidSim(request)
	if result.Error != nil {
		t.Fatalf("sim failed: %s", result.Error.Message)
	}

	var damage float64
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		if action.Id.GetSpellId() != 1260198 { // the triggered drain
			continue
		}
		for _, target := range action.Targets {
			damage += target.Damage
		}
	}

	if damage <= 0 {
		t.Error("Touch of the Grave never dealt damage for an undead rogue; the racial is missing from the sim")
	}
}
