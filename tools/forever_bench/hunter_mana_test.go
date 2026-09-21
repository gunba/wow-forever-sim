//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestHunterBaselineManaSustain(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "beast_mastery" && b.Key != "marksmanship" {
			continue
		}
		t.Run(b.Key, func(t *testing.T) {
			r := run(b, b.player(proto.Race_RaceOrc), 200, 20260920)
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
