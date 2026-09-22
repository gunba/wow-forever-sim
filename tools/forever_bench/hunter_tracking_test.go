//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/hunter"
)

func TestImprovedTrackingDoesNotDoubleDipCrits(t *testing.T) {
	tables := map[proto.MobType]*core.AttackTable{}
	var rank int32
	for _, mob := range []proto.MobType{proto.MobType_MobTypeUnknown, proto.MobType_MobTypeMechanical, proto.MobType_MobTypeDragonkin} {
		req := racialFixture("survival", proto.Race_RaceOrc)
		req.Encounter.Targets[0].MobType = mob
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		player := env.Raid.Parties[0].Players[0].(*hunter.Hunter)
		rank = player.Talents.ImprovedTracking
		tables[mob] = player.AttackTables[env.Encounter.TargetUnits[0].UnitIndex][proto.CastType_CastTypeRanged]
	}
	if rank == 0 {
		t.Fatal("fixture does not have Improved Tracking")
	}
	base := tables[proto.MobType_MobTypeUnknown]
	for mob, table := range tables {
		want := 1.0
		if mob == proto.MobType_MobTypeDragonkin {
			want += .01 * float64(rank)
		}
		if math.Abs(table.DamageDealtMultiplier/base.DamageDealtMultiplier-want) > 1e-8 {
			t.Fatalf("%v: wrong tracking damage multiplier", mob)
		}
		if table.CritMultiplier != base.CritMultiplier {
			t.Fatalf("%v: tracking added a separate critical multiplier", mob)
		}
	}
}
