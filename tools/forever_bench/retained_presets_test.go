//go:build with_db

package main

import (
	"encoding/json"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

func TestRetainedPresetsMatchBenchmark(t *testing.T) {
	var results struct{ Results []resultRow }
	if err := json.Unmarshal(mustRead("artifacts/forever_dps_5min.json"), &results); err != nil {
		t.Fatal(err)
	}
	expected := map[string]resultRow{}
	for _, row := range results.Results {
		expected[row.Key+"/"+row.Race] = row
	}
	definitions := map[string]build{}
	expectedCount := 0
	for _, b := range builds() {
		definitions[b.Key] = b
		for _, race := range b.races() {
			expectedCount++
			if _, ok := expected[b.Key+"/"+raceName(race)]; !ok {
				t.Errorf("missing current profile %s/%s", b.Key, raceName(race))
			}
		}
	}
	var bundle struct {
		Profiles []struct {
			Key, Race string
			Settings  json.RawMessage
		}
	}
	if err := json.Unmarshal(mustRead("ui/core/forever_ranked_profiles.json"), &bundle); err != nil {
		t.Fatal(err)
	}
	if len(expected) != expectedCount || len(bundle.Profiles) != expectedCount {
		t.Fatal("incomplete current ranking/default coverage")
	}
	for _, profile := range bundle.Profiles {
		key := profile.Key + "/" + profile.Race
		row, ok := expected[key]
		if !ok {
			t.Fatalf("unknown or duplicate web profile %s", key)
		}
		delete(expected, key)
		settings, req, baseline := &proto.IndividualSimSettings{}, &proto.RaidSimRequest{}, &proto.Player{}
		for message, data := range map[googleProto.Message]json.RawMessage{
			settings: profile.Settings, req: row.Request, baseline: row.BaselinePlayer,
		} {
			if err := protojson.Unmarshal(data, message); err != nil {
				t.Fatal(err)
			}
		}
		party := req.Raid.Parties[0]
		if !googleProto.Equal(settings.Player, party.Players[0]) ||
			!googleProto.Equal(settings.Encounter, req.Encounter) ||
			!googleProto.Equal(settings.RaidBuffs, req.Raid.Buffs) ||
			!googleProto.Equal(settings.PartyBuffs, party.Buffs) ||
			!googleProto.Equal(settings.Debuffs, req.Raid.Debuffs) {
			t.Errorf("%s: web default differs from its recorded request", key)
		}
		current := prepare(definitions[profile.Key], baseline.Race)
		if current.TalentsString != baseline.TalentsString || !googleProto.Equal(current.Rotation, baseline.Rotation) {
			t.Errorf("%s: talent/APL preset differs from the current benchmark", key)
		}
		if !googleProto.Equal(current, baseline) {
			t.Errorf("%s: native default differs from the complete recorded player", key)
		}
		if err := validateGear(baseline); err != nil {
			t.Fatalf("%s: %v", key, err)
		}
	}
}
