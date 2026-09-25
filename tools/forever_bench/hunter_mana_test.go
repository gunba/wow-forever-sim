//go:build with_db

package main

import (
	"encoding/json"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestHunterBaselineManaSustain(t *testing.T) {
	var saved struct{ Results []resultRow }
	if err := json.Unmarshal(mustRead("artifacts/modelled_gear/forever_input_profiles.json"), &saved); err != nil {
		t.Fatal(err)
	}
	players := map[string]json.RawMessage{}
	for _, row := range saved.Results {
		if row.Race == "Orc" {
			players[row.Key] = row.BaselinePlayer
		}
	}
	for _, b := range builds() {
		if b.Key != "beast_mastery" && b.Key != "marksmanship" {
			continue
		}
		t.Run(b.Key, func(t *testing.T) {
			// Exercise the actual ranking loadout, not the old seed gear.
			p := &proto.Player{}
			if err := protojson.Unmarshal(players[b.Key], p); err != nil {
				t.Fatal(err)
			}
			r := run(b, p, 200, 20260920)
			if r.OOMSeconds > 5 {
				t.Fatalf("mana-limited time = %.2fs in 300s", r.OOMSeconds)
			}
			metrics := &proto.UnitMetrics{}
			if err := protojson.Unmarshal(r.Metrics, metrics); err != nil {
				t.Fatal(err)
			}
			for _, itemID := range []int32{13444, 12662} {
				used := false
				for _, resource := range metrics.Resources {
					if resource.Id.GetItemId() == itemID && resource.ActualGain > 0 {
						used = true
					}
				}
				if !used {
					t.Errorf("mana consumable %d was not used", itemID)
				}
			}
		})
	}
}
