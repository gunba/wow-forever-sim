package druid

import "testing"

// The previous sparse 25/40/50/60 map produced an empty spell at level 30.
func TestPrimalBiteRankAtIntermediateLevels(t *testing.T) {
	for _, tc := range []struct {
		level int32
		rank  int
	}{
		{24, 0}, {25, 1}, {30, 1}, {35, 1}, {36, 2}, {47, 2}, {48, 3}, {59, 3}, {60, 4},
	} {
		if got := primalBiteRank(tc.level); got != tc.rank {
			t.Errorf("level %v: rank %v, want %v", tc.level, got, tc.rank)
		}
	}
}
