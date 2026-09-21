//go:build with_db

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

func TestRetainedPresetsMatchBenchmark(t *testing.T) {
	checked := 0
	for _, b := range builds() {
		path := filepath.Join("artifacts", "optimization", b.Key+".json")
		data, err := os.ReadFile(path)
		if os.IsNotExist(err) {
			continue // This build has not completed its optimization pass.
		}
		if err != nil {
			t.Fatal(err)
		}
		var bundle struct {
			Retained struct {
				Results []resultRow
			}
		}
		if err := json.Unmarshal(data, &bundle); err != nil {
			t.Fatalf("%s: %v", path, err)
		}
		if len(bundle.Retained.Results) != len(b.races()) {
			t.Fatalf("%s: incomplete retained race coverage", b.Key)
		}
		for _, row := range bundle.Retained.Results {
			player := &proto.Player{}
			if err := protojson.Unmarshal(row.BaselinePlayer, player); err != nil {
				t.Fatal(err)
			}
			// Compare the pre-normalization input. Paid hit adjustments belong
			// in the recorded request, not the equipment catalog or UI presets.
			current := prepare(b, player.Race)
			if !googleProto.Equal(current, player) {
				t.Errorf("%s/%s: current preset differs from the validated input", b.Key, row.Race)
			}
			checked++
		}
	}
	if checked == 0 {
		t.Fatal("no retained profiles found")
	}
}
