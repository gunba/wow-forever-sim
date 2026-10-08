//go:build with_db

package main

import (
	"fmt"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestForeverNoThreatManaConversions(t *testing.T) {
	for _, tc := range []struct {
		name, build    string
		spend, restore int32
	}{
		{"Life Tap", "demonology", 25307, 11689},
		{"Dark Sacrifice", "shadow", 10894, 1277328},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := historyTalentFixture(tc.build, map[string]int{"demonicEnergies": 2})
			p := req.Raid.Parties[0].Players[0]
			p.ForeverTier1Bonuses = false
			if w := p.GetWarlock(); w != nil {
				w.Options.Sacrifice = proto.WarlockOptions_NoSummon
				w.Options.Summon = proto.WarlockOptions_Imp
			} else {
				p.Race = proto.Race_RaceUndead
			}
			p.Rotation = core.APLRotationFromJsonString(fmt.Sprintf(`{"type":"TypeAPL","priorityList":[{"action":{"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":%d}}},{"castSpell":{"spellId":{"spellId":%d}}}]}}}]}`, tc.spend, tc.restore))
			req.Encounter.Duration = 18
			req.SimOptions.Iterations = 1
			result := core.RunRaidSim(req)
			if result.Error != nil {
				t.Fatal(result.Error)
			}
			player := result.RaidMetrics.Parties[0].Players[0]
			var gain float64
			for _, metrics := range append([]*proto.UnitMetrics{player}, player.Pets...) {
				for _, resource := range metrics.Resources {
					if resource.Id.GetSpellId() == tc.restore && resource.Type == proto.ResourceType_ResourceTypeMana {
						gain += resource.ActualGain
					}
				}
				for _, action := range metrics.Actions {
					if action.Id.GetOtherId() == proto.OtherAction_OtherActionManaGain {
						for _, target := range action.Targets {
							if target.Threat != 0 {
								t.Errorf("%s gained %v phantom resource threat", metrics.Name, target.Threat)
							}
						}
					}
				}
			}
			if gain <= 0 {
				t.Fatal("fixture did not restore mana")
			}
		})
	}
}
