//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	googleProto "google.golang.org/protobuf/proto"
)

func TestProtectionDoesNotPoolForExternallyBlockedSunder(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "tank_warrior" {
			continue
		}
		player := b.rankedPlayer(proto.Race_RaceHuman)
		player.Rotation = core.GetAplRotation("ui/tank_warrior/apls", "forever_protection").Rotation
		request := tankRequest(b, player, 1, 20291951, 1, 300)
		request.Raid.Debuffs.ExposeArmor = proto.TristateEffect_TristateEffectImproved
		control := googleProto.Clone(request).(*proto.RaidSimRequest)
		// In this custom scenario, permanent Expose prevents Sunder. Compare
		// against the same priorities without discretionary rage reservations.
		for _, row := range control.Raid.Parties[0].Players[0].Rotation.PriorityList {
			action := row.Action
			if action.Condition.GetOr() != nil {
				action.Condition = nil
			} else if and := action.Condition.GetAnd(); and != nil && len(and.Vals) == 2 &&
				and.Vals[1].GetOr() != nil {
				action.Condition = and.Vals[0]
			}
		}
		actual, unreserved := core.RunRaidSim(request), core.RunRaidSim(control)
		if actual.Error != nil || unreserved.Error != nil {
			t.Fatalf("simulation failed: %v / %v", actual.Error, unreserved.Error)
		}
		if math.Abs(actual.RaidMetrics.Dps.Avg-unreserved.RaidMetrics.Dps.Avg) > 1e-9 {
			t.Fatalf("blocked Sunder starved damage: %g vs %g DPS",
				actual.RaidMetrics.Dps.Avg, unreserved.RaidMetrics.Dps.Avg)
		}
	}
}

func sunderMaintenance(request *proto.RaidSimRequest) (time.Duration, time.Duration, int) {
	sim := core.NewSim(request, simsignals.Signals{})
	aura := sim.Encounter.TargetUnits[0].GetAura("Sunder Armor")
	original := aura.OnStacksChange
	var firstFive, lostAt, downtime time.Duration
	drops := 0
	lost := false
	aura.OnStacksChange = func(a *core.Aura, s *core.Simulation, old, next int32) {
		original(a, s, old, next)
		if s.CurrentTime >= s.Duration {
			return
		}
		if next == 5 {
			if firstFive == 0 {
				firstFive = s.CurrentTime
			} else if old != 5 {
				downtime += s.CurrentTime - lostAt
			}
			lost = false
		} else if old == 5 {
			drops++
			lostAt = s.CurrentTime
			lost = true
		}
	}
	sim.Reset()
	sim.PrePull()
	for !sim.Step() {
	}
	if lost {
		downtime += sim.Duration - lostAt
	}
	sim.Cleanup()
	return firstFive, downtime, drops
}

func TestProtectionMaintainsSunderUnderRagePressure(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "tank_warrior" {
			continue
		}
		player := b.rankedPlayer(proto.Race_RaceNightElf)
		preset := core.GetAplRotation("ui/tank_warrior/apls", "forever_protection").Rotation
		if !googleProto.Equal(player.Rotation, preset) {
			t.Fatal("ranked default must contain the maintained Sunder rotation")
		}
		request := tankRequest(b, player, 1, 20291951, 1, 300)
		if request.Raid.Debuffs.SunderArmor || request.Raid.Debuffs.ExposeArmor != 0 {
			t.Fatal("maintenance must not be supplied by an external armor debuff")
		}
		first, downtime, drops := sunderMaintenance(request)
		if first == 0 || first > 20*time.Second || drops != 0 || downtime != 0 {
			t.Fatalf("first five stacks %v; %d subsequent drops / %v below five", first, drops, downtime)
		}
	}
}
