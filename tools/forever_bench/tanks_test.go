//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	googleProto "google.golang.org/protobuf/proto"
)

func TestPublishedTankControlsEnterBenchmark(t *testing.T) {
	profiles := 0
	for _, b := range builds() {
		if !b.isTank() {
			continue
		}
		for _, race := range b.races() {
			profiles++
			player := b.player(race)
			if err := validateGear(player); err != nil {
				t.Fatalf("%s/%s: %v", b.Key, raceName(race), err)
			}
			if err := b.validateTankIdentity(player); err != nil {
				t.Fatal(err)
			}
			if err := loadTalents(b).validate(player.TalentsString); err != nil {
				t.Fatalf("%s: %v", b.Key, err)
			}
		}
		control := tankControl(b.Key)
		player := googleProto.Clone(control.Raid.Parties[0].Players[0]).(*proto.Player)
		request := tankRequest(b, player, int(control.SimOptions.Iterations), control.SimOptions.RandomSeed, 1, 300)
		if !googleProto.Equal(control, request) {
			t.Fatalf("%s: baseline template changed its published scenario", b.Key)
		}
		multi := tankRequest(b, player, 10, 11, 3, 90)
		if len(multi.Encounter.Targets) != 3 || multi.Encounter.Duration != 90 ||
			multi.Raid.Parties[0].Players[0].HealingModel.Hps != 2500 ||
			multi.Encounter.Targets[0].MinBaseDamage != 1500 ||
			player.HealingModel.Hps != 1500 || len(control.Encounter.Targets) != 1 {
			t.Fatal("multi-target configuration mutated the frozen single-target control")
		}
		if b.Key != "feral_tank_druid" {
			player.Equipment.Items[15] = &proto.ItemSpec{}
			if b.validateTankIdentity(player) == nil {
				t.Fatal("shield tank accepted an empty off-hand")
			}
		}
	}
	if profiles != 17 {
		t.Fatalf("%d tank race profiles, want 17", profiles)
	}
}

func TestTankTargetThreatIgnoresSelfAndUsesUnitIndices(t *testing.T) {
	result := &proto.RaidSimResult{
		IterationsDone: 10, AvgIterationDuration: 100,
		RaidMetrics: &proto.RaidMetrics{Parties: []*proto.PartyMetrics{{
			Players: []*proto.UnitMetrics{{Actions: []*proto.ActionMetrics{{
				Targets: []*proto.TargetedActionMetrics{
					{UnitIndex: 5, Threat: 100000},
					{UnitIndex: 2, Threat: 300000},
					{UnitIndex: 7, Threat: 9000000},
				},
			}}}},
		}}},
		EncounterMetrics: &proto.EncounterMetrics{Targets: []*proto.UnitMetrics{
			{UnitIndex: 2}, {UnitIndex: 5},
		}},
	}
	metrics := tankResultMetrics(result)
	if metrics.TargetTPS[0] != 300 || metrics.TargetTPS[1] != 100 ||
		metrics.LeastTargetTPS != 100 || math.IsNaN(metrics.TPSStandardError) {
		t.Fatalf("incorrect per-target threat: %+v", metrics)
	}
}

func TestTankSelectionRejectsDPSFromWorseSurvivalOrThreatCoverage(t *testing.T) {
	row := resultRow{DPS: 500, Tank: &tankMetrics{
		TPS: 1000, DTPS: 500, TMI: 50, ChanceOfDeath: .01, LeastTargetTPS: 100,
	}}
	reference := tankPair{Single: row, Multi: row}
	for _, change := range []func(*tankMetrics){
		func(m *tankMetrics) { m.DTPS *= 1.02 },
		func(m *tankMetrics) { m.TMI *= 1.03 },
		func(m *tankMetrics) { m.ChanceOfDeath += .005 },
		func(m *tankMetrics) { m.TPS *= .98 },
		func(m *tankMetrics) { m.LeastTargetTPS *= .97 },
	} {
		candidate := reference
		metrics := *row.Tank
		candidate.Multi.Tank = &metrics
		candidate.Single.DPS *= 2
		change(candidate.Multi.Tank)
		if len(tankPolicy.failures(candidate, reference)) == 0 {
			t.Fatal("a large single-target gain bypassed a multi-target tank guard")
		}
	}
	if len(tankPolicy.failures(reference, reference)) != 0 {
		t.Fatal("identical controls must pass")
	}
}
