//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestWarlockImpFireboltRequiresActivePet(t *testing.T) {
	req := racialFixture("ds_ruin", proto.Race_RaceOrc)
	req.Encounter.Duration = 30
	// The ranked DS/Ruin APL sacrifices its Imp at -5 seconds; remove that
	// prepull action to check the learned Firebolt while the pet is active.
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	for _, pet := range result.RaidMetrics.Parties[0].Players[0].Pets {
		if pet.Name == "Imp" {
			if pet.Dps.Avg <= 0 {
				t.Fatal("Active Imp made no attacks over 30 seconds")
			}
			return
		}
	}
	t.Fatal("Imp was not in the raid's pet metrics")
}
