//go:build with_db

package main

import (
	"encoding/json"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestActiveSourceAPLsUseHighestRegisteredRanks(t *testing.T) {
	for _, b := range builds() {
		checkHighestRegisteredRanks(t, b, b.player(b.races()[0]))
	}
}

func TestActiveRankedAPLsUseHighestRegisteredRanks(t *testing.T) {
	for _, b := range builds() {
		for _, race := range b.races() {
			checkHighestRegisteredRanks(t, b, b.rankedPlayer(race))
		}
	}
}

func checkHighestRegisteredRanks(t *testing.T, b build, p *proto.Player) {
	t.Helper()
	inventory := inventorySpells(b, p)
	best := map[int32]registeredSpellRank{}
	for _, s := range inventory.Ranks {
		if s.Castable && s.Code != 0 && s.Rank > best[s.Code].Rank {
			best[s.Code] = s
		}
	}
	lower := map[int32]int32{}
	for _, s := range inventory.Ranks {
		if s.Castable && s.Code != 0 && s.Rank > 0 && s.Rank < best[s.Code].Rank {
			lower[s.ID] = best[s.Code].ID
		}
	}
	data, err := protojson.Marshal(p.Rotation)
	if err != nil {
		t.Fatal(err)
	}
	var rotation any
	if err := json.Unmarshal(data, &rotation); err != nil {
		t.Fatal(err)
	}
	var walk func(any)
	walk = func(value any) {
		switch v := value.(type) {
		case map[string]any:
			for k, child := range v {
				if id, ok := child.(float64); k == "spellId" && ok {
					if top, exists := lower[int32(id)]; exists {
						t.Errorf("%s references lower-rank spell %v; highest registered rank is %v", b.Key, id, top)
					}
				} else {
					walk(child)
				}
			}
		case []any:
			for _, child := range v {
				walk(child)
			}
		}
	}
	walk(rotation)
}
