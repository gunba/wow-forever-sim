package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

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

func TestForeverRageRefundUsesPaidCost(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		for _, paid := range []float64{0, 5, 15} {
			unit := &Unit{Env: &Environment{Ruleset: ruleset}}
			unit.rageBar = rageBar{unit: unit, currentRage: 20, maxRage: 100, RageRefundMetrics: &ResourceMetrics{}}
			spell := &Spell{Unit: unit}
			spell.Cost = newRageCost(spell, RageCostOptions{Cost: 15, Refund: .8})
			spell.CurCast.Cost = paid
			sim := &Simulation{Options: &proto.SimOptions{Interactive: true}, pendingActions: []*PendingAction{{NextActionAt: NeverExpires}}}
			spell.Cost.SpendCost(sim, spell)
			spell.Cost.IssueRefund(sim, spell)
			refund := 12.0
			if ruleset == proto.Ruleset_RulesetForever {
				refund = .8 * paid
			}
			if got, want := unit.CurrentRage(), 20-paid+refund; got != want {
				t.Errorf("%v paid %v: Rage %v, want %v", ruleset, paid, got, want)
			}
		}
	}
}

func TestForeverIncomingRageReferenceNormalization(t *testing.T) {
	// This is an explicit sensitivity scenario, not a server measurement.
	parameters := ForeverIncomingRageParameters{Coefficient: 10, ReferenceArmor: .5, ExpectedHealth: 200}
	if got := parameters.rageFromDamageTaken(&SpellResult{Damage: 6}); got != .6 {
		t.Fatalf("6 delivered damage at reference Armor .5 gave %g Rage, want .6", got)
	}
	if got := parameters.rageFromDamageTaken(&SpellResult{}); got != 0 {
		t.Fatalf("missed hit generated %g Rage", got)
	}
}
