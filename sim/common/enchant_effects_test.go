package common

import "testing"

func TestCrusaderStrengthLevelClamp(t *testing.T) {
	for _, tc := range []struct {
		level int32
		want  float64
	}{{20, 100}, {60, 100}, {65, 80}, {85, 0}} {
		if got := crusaderStrengthBonus(tc.level); got != tc.want {
			t.Errorf("level %v: Strength %v, want %v", tc.level, got, tc.want)
		}
	}
}
