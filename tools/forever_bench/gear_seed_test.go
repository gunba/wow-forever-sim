//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestSeedScoreUsesCatalogRatingPool(t *testing.T) {
	item := core.Item{Stats: stats.Stats{
		stats.MeleeHit: 1, stats.MeleeCrit: 1,
	}}
	for _, test := range []struct {
		build build
		want  float64
	}{
		{build{Key: "fury", Class: proto.Class_ClassWarrior}, 42},
		{build{Key: "arcane", Class: proto.Class_ClassMage}, 12 + 70.0/6.0},
	} {
		if got := seedScore(test.build, item, 0); math.Abs(got-test.want) > 1e-9 {
			t.Errorf("%s: shared rating score %v, want %v", test.build.Key, got, test.want)
		}
	}
}
