package core

import "testing"

// `^` is XOR in Go, not a power, so the level term used to read 60^2 = 62 instead of 3600 and
// every warrior generated 16% too much rage.
func TestRageConversionAtLevel60(t *testing.T) {
	if got := GetRageConversion(60); got < 230.5 || got > 230.7 {
		t.Fatalf("rage conversion at 60 = %v, want 230.6", got)
	}
}

func TestForeverWarriorRagePerSwing(t *testing.T) {
	for _, tc := range []struct {
		speed            float64
		twoHand, offHand bool
		want             float64
	}{
		{1.5, false, false, 1.5 * 4.5 / 1.3},
		{2.7, false, false, 2.7 * 4.5 / 1.3},
		{3.2, true, false, 14.4},
		{1.6, false, true, 1.6 * 4.5 / 1.3 / 2},
		{0, false, false, 0},
	} {
		got := foreverWarriorRagePerSwing(tc.speed, tc.twoHand, tc.offHand)
		if delta := got - tc.want; delta < -0.000001 || delta > 0.000001 {
			t.Errorf("speed %.1f two-hand %t off-hand %t: got %.4f, want %.4f", tc.speed, tc.twoHand, tc.offHand, got, tc.want)
		}
	}
}

func TestForeverWarriorDamageTakenUsesPreArmorHealthRatio(t *testing.T) {
	if got := foreverWarriorDamageTakenRage(6, .5, 200); got != .6 {
		t.Fatalf("mitigated 12 damage against 200 health gave %g rage, want 0.6", got)
	}
	if got := foreverWarriorDamageTakenRage(0, .5, 200); got != 0 {
		t.Fatalf("missed hit generated %g rage", got)
	}
}
