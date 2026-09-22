//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestWindfurySourcesRegisterOnce(t *testing.T) {
	var baseline float64
	for _, sources := range []struct{ imbue, party bool }{{true, false}, {false, true}, {true, true}} {
		req := racialFixture("retribution", proto.Race_RaceHuman)
		req.SimOptions.Iterations, req.SimOptions.RandomSeed = 1, 20295711
		player := req.Raid.Parties[0].Players[0]
		player.Consumes = &proto.Consumes{}
		player.Rotation = &proto.APLRotation{}
		if sources.imbue {
			player.Consumes.MainHandImbue = proto.WeaponImbue_Windfury
		}
		req.Raid.Parties[0].Buffs = &proto.PartyBuffs{WindfuryTotem: sources.party}
		result := core.RunRaidSim(req)
		if result.Error != nil {
			t.Fatal(result.Error)
		}
		metrics := result.RaidMetrics.Parties[0].Players[0]
		found := false
		for _, aura := range metrics.Auras {
			if aura.Id.GetSpellId() == 21919 {
				found = true
				if aura.ProcsAvg != 1 {
					t.Fatalf("%+v: extra-attack tracking activated %v times at reset", sources, aura.ProcsAvg)
				}
			}
		}
		if !found {
			t.Fatal("extra-attack tracking missing")
		}
		if baseline == 0 {
			baseline = result.RaidMetrics.Dps.Avg
		} else if baseline != result.RaidMetrics.Dps.Avg {
			t.Fatalf("%+v: source selection changed damage", sources)
		}
	}
}
