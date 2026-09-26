package warlock

import (
	"math"
	"testing"
)

func TestForeverLifeTapRankAndLevelConversion(t *testing.T) {
	for _, tc := range []struct {
		rank, level int
		base        float64
	}{
		{1, 6, 20}, {1, 10, 24}, {1, 16, 30}, {1, 60, 30},
		{2, 16, 65}, {2, 60, 75}, {3, 60, 140}, {4, 60, 220},
		{5, 60, 310}, {6, 56, 420}, {6, 60, 424},
	} {
		if got := foreverLifeTapConversion(tc.rank, int32(tc.level), 50, 2); math.Abs(got-(tc.base+50)*1.2) > 1e-9 {
			t.Errorf("rank %d level %d: %v, want %v", tc.rank, tc.level, got, (tc.base+50)*1.2)
		}
	}
}
