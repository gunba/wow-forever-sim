package warlock

import (
	"math"
	"testing"
)

func TestForeverImpFireboltRankSevenBaseRange(t *testing.T) {
	low, high := impFireboltBaseDamage(7)[0], impFireboltBaseDamage(7)[1]
	if math.Abs(low-41.5) > 0.001 || math.Abs(high-46.5) > 0.001 {
		t.Fatalf("Forever rank-7 Firebolt base range %.3f–%.3f, want 41.5–46.5", low, high)
	}
}
