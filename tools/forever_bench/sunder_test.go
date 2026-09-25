//go:build with_db

package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestFurySunderProvidesOwnArmorReduction(t *testing.T) {
	var provider build
	for _, b := range builds() {
		if b.Key == "fury_sunder" {
			provider = b
			break
		}
	}
	if provider.Key == "" {
		t.Fatal("missing Fury Sunder build")
	}
	var original build
	for _, b := range builds() {
		if b.Key == "fury" {
			original = b
			break
		}
	}
	if original.Key == "" {
		t.Fatal("missing ordinary Fury build")
	}
	unburdened := requestForBuild(original, original.player(proto.Race_RaceOrc), 1, 1)
	if !unburdened.Raid.Debuffs.SunderArmor {
		t.Fatal("ordinary Fury must keep its established external raid debuffs")
	}
	input, err := os.ReadFile("artifacts/modelled_gear/forever_input_profiles.json")
	if err != nil {
		t.Fatal(err)
	}
	var archive struct {
		Results []struct {
			Key, Race      string
			BaselinePlayer json.RawMessage
		}
	}
	if err := json.Unmarshal(input, &archive); err != nil {
		t.Fatal(err)
	}
	exact := map[string]*proto.Player{}
	for _, row := range archive.Results {
		if row.Key != provider.Key {
			continue
		}
		p := &proto.Player{}
		if err := protojson.Unmarshal(row.BaselinePlayer, p); err != nil {
			t.Fatal(err)
		}
		exact[row.Race] = p
	}
	for _, race := range provider.races() {
		player := exact[raceName(race)]
		if player == nil {
			t.Fatalf("no Sunder profile for %s", raceName(race))
		}
		paid, _, err := capHit(provider, player)
		if err != nil {
			t.Fatal(err)
		}
		for _, seed := range []int64{20299021, 20299047, 20299073, 20299101, 20299127, 20299163, 20299191, 20299229, 20299257, 20299311} {
			req := requestForBuild(provider, paid, 1, seed)
			if req.Raid.Debuffs.SunderArmor ||
				req.Raid.Debuffs.ExposeArmor != proto.TristateEffect_TristateEffectMissing {
				t.Fatal("Sunder provider cannot start with a competing major armor debuff")
			}
			req.SimOptions.Debug = true
			result := core.RunRaidSim(req)
			if result.Error != nil {
				t.Fatal(result.Error.Message)
			}
			var sunderCasts float64
			for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
				if action.Id.GetSpellId() == 11597 {
					sunderCasts = float64(action.Targets[0].Casts)
				}
			}
			// There can only be one five-to-zero reset: cleanup at fight end.
			// A second reset means the required five-stack debuff fell off.
			if sunderCasts < 14 || sunderCasts > 35 ||
				!strings.Contains(result.Logs, "{SpellID: 11597} stacks: 4 --> 5") ||
				strings.Count(result.Logs, "{SpellID: 11597} stacks: 5 --> 0") != 1 {
				var drops []string
				for _, line := range strings.Split(result.Logs, "\n") {
					if strings.Contains(line, "{SpellID: 11597} stacks: 5 --> 0") {
						drops = append(drops, line)
					}
				}
				t.Errorf("%s seed %d: Sunder uptime failed, %.0f casts, %d resets",
					raceName(race), seed, sunderCasts, len(drops))
				t.Log(drops)
			}
		}
	}
}
