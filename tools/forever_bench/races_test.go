//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestBothFactionRoster(t *testing.T) {
	counts := map[string]int{}
	for _, b := range builds() {
		seen := map[proto.Race]bool{}
		for _, race := range b.races() {
			if seen[race] {
				t.Fatalf("%s repeats race %s", b.Key, raceName(race))
			}
			seen[race] = true
			counts[raceFaction(race)]++
			p := b.player(race)
			if err := validateGear(p); err != nil {
				t.Errorf("%s/%s: %v", b.Key, raceName(race), err)
			}
			req := requestForBuild(b, p, 1, 1)
			_, result, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
			rotation := result.Parties[0].Players[0].RotationStats
			for _, action := range append(rotation.GetPrepullActions(), rotation.GetPriorityList()...) {
				for _, warning := range action.Warnings {
					t.Errorf("%s/%s APL: %s", b.Key, raceName(race), warning)
				}
			}
		}
	}
	if counts["Horde"] != 102 || counts["Alliance"] != 99 {
		t.Fatalf("unexpected build/race coverage: %v", counts)
	}
	for _, b := range builds() {
		if b.Class == proto.Class_ClassPaladin && b.allowsRace(proto.Race_RaceOrc) {
			t.Fatal("Orc Paladin is not in the client roster")
		}
		if b.Class == proto.Class_ClassMage && !b.allowsRace(proto.Race_RaceGnome) {
			t.Fatal("Gnome Mage is missing from the client roster")
		}
	}
}
