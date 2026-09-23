//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestForeverWarriorAutoRageIgnoresDamage(t *testing.T) {
	measure := func(bonusAP float64) (int32, float64, float64) {
		req := racialFixture("arms", proto.Race_RaceOrc)
		req.Encounter.Duration = 12
		req.SimOptions.Iterations = 1
		req.SimOptions.RandomSeed = 1927
		player := req.Raid.Parties[0].Players[0]
		player.Rotation = &proto.APLRotation{}
		if player.BonusStats == nil {
			player.BonusStats = &proto.UnitStats{Stats: stats.Stats{}.ToFloatArray()}
		}
		player.BonusStats.Stats[stats.AttackPower] += bonusAP
		weapon := core.ItemsByID[player.Equipment.Items[proto.ItemSlot_ItemSlotMainHand].Id]
		if weapon.HandType != proto.HandType_HandTypeTwoHand {
			t.Fatal("test profile no longer has a two-handed weapon")
		}
		result := core.RunRaidSim(req)
		if result.Error != nil {
			t.Fatal(result.Error)
		}
		for _, resource := range result.RaidMetrics.Parties[0].Players[0].Resources {
			if resource.Id.GetOtherId() == proto.OtherAction_OtherActionAttack && resource.Id.GetTag() == 1 {
				return resource.Events, resource.Gain, weapon.SwingSpeed
			}
		}
		t.Fatal("no main-hand autoattack rage")
		return 0, 0, 0
	}

	events, rage, speed := measure(0)
	if events == 0 || math.Abs(rage-float64(events)*speed*4.5) > 0.001 {
		t.Fatalf("%d landed 2H swings at %.2fs gained %.3f rage, expected %.3f", events, speed, rage, float64(events)*speed*4.5)
	}
	boostedEvents, boostedRage, _ := measure(2000)
	if events != boostedEvents || math.Abs(rage-boostedRage) > 0.001 {
		t.Fatalf("adding AP changed normalized auto rage: %d/%.3f versus %d/%.3f", events, rage, boostedEvents, boostedRage)
	}
}
