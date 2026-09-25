package tankwarrior

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestDemoralizingShoutThreatByRuleset(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetClassic, proto.Ruleset_RulesetForever} {
		t.Run(ruleset.String(), func(t *testing.T) {
			player := &proto.Player{
				Class: proto.Class_ClassWarrior, Race: proto.Race_RaceHuman,
				Equipment: core.GetGearSet("../../../ui/tank_warrior/gear_sets", "launch").GearSet,
				Spec:      &proto.Player_TankWarrior{TankWarrior: &proto.TankWarrior{Options: &proto.TankWarrior_Options{StartingRage: 20}}},
				Rotation:  core.APLRotationFromJsonString(`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":11556}}}}]}`),
			}
			result := core.RunRaidSim(&proto.RaidSimRequest{
				Raid:       core.SinglePlayerRaidProto(player, nil, nil, nil),
				Encounter:  &proto.Encounter{Duration: 20, Targets: []*proto.Target{{Level: 63}}},
				SimOptions: &proto.SimOptions{Iterations: 50, RandomSeed: 19788, Ruleset: ruleset},
			})
			if result.Error != nil {
				t.Fatal(result.Error)
			}
			for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
				if action.Id.GetSpellId() != 11556 {
					continue
				}
				if action.Targets[0].Casts == 0 {
					t.Fatal("Demoralizing Shout was not cast")
				}
				threat := action.Targets[0].Threat
				if ruleset == proto.Ruleset_RulesetForever && threat != 0 {
					t.Fatalf("Forever Demoralizing Shout caused %.2f threat", threat)
				}
				if ruleset == proto.Ruleset_RulesetClassic && threat <= 0 {
					t.Fatalf("Classic Demoralizing Shout lost its threat: %.2f", threat)
				}
				return
			}
			t.Fatal("missing Demoralizing Shout metrics")
		})
	}
}

func TestShieldBlockChargeAndDurationByRuleset(t *testing.T) {
	for _, tc := range []struct {
		ruleset proto.Ruleset
		blocks  int32
		length  time.Duration
	}{
		{proto.Ruleset_RulesetClassic, 1, 5 * time.Second},
		{proto.Ruleset_RulesetForever, 2, 7 * time.Second},
	} {
		t.Run(tc.ruleset.String(), func(t *testing.T) {
			player := &proto.Player{
				Class: proto.Class_ClassWarrior, Race: proto.Race_RaceHuman,
				Equipment: core.GetGearSet("../../../ui/tank_warrior/gear_sets", "launch").GearSet,
				Spec:      &proto.Player_TankWarrior{TankWarrior: &proto.TankWarrior{Options: &proto.TankWarrior_Options{}}},
			}
			environment, _, _ := core.NewEnvironment(core.SinglePlayerRaidProto(player, nil, nil, nil), core.MakeSingleTargetEncounter(0), tc.ruleset, false)
			aura := environment.Raid.Parties[0].Players[0].(*TankWarrior).ShieldBlockAura
			if aura == nil || aura.MaxStacks != tc.blocks || aura.Duration != tc.length {
				t.Fatalf("Shield Block: got %+v, want %d blocks for %s", aura, tc.blocks, tc.length)
			}
		})
	}
}
