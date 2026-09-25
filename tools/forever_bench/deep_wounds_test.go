//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestForeverDeepWoundsCannotCrit(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.SimOptions.Iterations = 250
	req.SimOptions.RandomSeed = 20295719
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		if action.Id.GetSpellId() != 12867 {
			continue
		}
		var ticks, critTicks int32
		for _, target := range action.Targets {
			ticks += target.Ticks
			critTicks += target.CritTicks
		}
		if ticks == 0 || critTicks != 0 {
			t.Fatalf("Deep Wounds: %d normal ticks, %d critical ticks", ticks, critTicks)
		}
		return
	}
	t.Fatal("Fury did not trigger Deep Wounds")
}
