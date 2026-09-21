package core

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func energyTestUnit(ruleset proto.Ruleset) (*Unit, *Simulation) {
	unit := &Unit{Env: &Environment{Ruleset: ruleset}, PseudoStats: stats.NewPseudoStats()}
	unit.EnableEnergyBar(100)
	sim := &Simulation{Options: &proto.SimOptions{Interactive: true}}
	unit.nextEnergyTick = unit.energyTickDuration()
	return unit, sim
}

func TestEnergyRegenerationCadence(t *testing.T) {
	for _, tc := range []struct {
		ruleset proto.Ruleset
		tick    time.Duration
		energy  float64
	}{
		{proto.Ruleset_RulesetClassic, 2020 * time.Millisecond, 20.2},
		{proto.Ruleset_RulesetForever, 100 * time.Millisecond, 1},
	} {
		t.Run(tc.ruleset.String(), func(t *testing.T) {
			unit, sim := energyTestUnit(tc.ruleset)
			sim.CurrentTime = tc.tick - time.Nanosecond
			unit.energyBar.RunTask(sim)
			if unit.CurrentEnergy() != 0 {
				t.Fatal("regenerated before the scheduled update")
			}
			sim.CurrentTime = tc.tick
			next := unit.energyBar.RunTask(sim)
			if math.Abs(unit.CurrentEnergy()-tc.energy) > 1e-9 || next != 2*tc.tick {
				t.Fatalf("energy %v, next %v; want %v, %v", unit.CurrentEnergy(), next, tc.energy, 2*tc.tick)
			}
			if got := tc.energy / tc.tick.Seconds(); got != 10 {
				t.Fatalf("regen %v/sec, want 10", got)
			}
		})
	}
}

func TestForeverEnergyRegenMultiplierTransitions(t *testing.T) {
	unit, sim := energyTestUnit(proto.Ruleset_RulesetForever)
	sim.CurrentTime = 50 * time.Millisecond
	unit.ResetEnergyTick(sim)
	unit.EnergyTickMultiplier = 2
	sim.CurrentTime = 150 * time.Millisecond
	unit.energyBar.RunTask(sim)
	sim.CurrentTime = 200 * time.Millisecond
	unit.ResetEnergyTick(sim)
	unit.EnergyTickMultiplier = 1
	sim.CurrentTime = 300 * time.Millisecond
	unit.energyBar.RunTask(sim)
	// 50 ms at 10/sec, 150 ms at 20/sec, 100 ms at 10/sec.
	if got := unit.CurrentEnergy(); math.Abs(got-4.5) > 1e-9 {
		t.Fatalf("energy %v, want 4.5", got)
	}
}

func TestForeverEnergyCapAndWastedRegen(t *testing.T) {
	unit, sim := energyTestUnit(proto.Ruleset_RulesetForever)
	unit.currentEnergy = 99.5
	sim.CurrentTime = 100 * time.Millisecond
	unit.energyBar.RunTask(sim)
	if unit.CurrentEnergy() != 100 || unit.energyBar.regenMetrics.Gain != 1 || unit.energyBar.regenMetrics.ActualGain != .5 {
		t.Fatalf("energy/regen = %v / %+v", unit.CurrentEnergy(), unit.energyBar.regenMetrics)
	}
}

func TestEnergyRefundUsesPaidCost(t *testing.T) {
	for _, paid := range []float64{60, 48} {
		unit, sim := energyTestUnit(proto.Ruleset_RulesetForever)
		unit.currentEnergy = 100
		unit.SpendEnergy(sim, paid, unit.NewEnergyMetrics(ActionID{SpellID: 1241584}))
		cost := &EnergyCost{Refund: .8, RefundMetrics: unit.EnergyRefundMetrics}
		cost.IssueRefund(sim, &Spell{Unit: unit, CurCast: Cast{Cost: paid}})
		if got, want := unit.CurrentEnergy(), 100-.2*paid; math.Abs(got-want) > 1e-9 {
			t.Fatalf("paid %v: remaining energy %v, want %v", paid, got, want)
		}
	}
}

func TestForeverEnergyHaste(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetClassic, proto.Ruleset_RulesetForever} {
		unit, sim := energyTestUnit(ruleset)
		unit.stats[stats.MeleeHaste] = 20 * HasteRatingPerHastePercent
		unit.PseudoStats.EnergyHasteMultiplier = 1.1
		unit.regenHasteMultiplier = unit.energyHasteMultiplier()
		want := 1.0
		if ruleset == proto.Ruleset_RulesetForever {
			want = 1.32
		}
		if got := unit.regenHasteMultiplier; math.Abs(got-want) > 1e-9 {
			t.Fatalf("%s haste %v, want %v", ruleset, got, want)
		}
		unit.MultiplyMeleeSpeed(sim, 1.4) // Slice and Dice-like attack speed is not general haste.
		if got := unit.energyHasteMultiplier(); math.Abs(got-want) > 1e-9 {
			t.Fatalf("attack speed changed Energy haste: %v", got)
		}
	}
}

func TestForeverEnergyHasteTransitions(t *testing.T) {
	unit, sim := energyTestUnit(proto.Ruleset_RulesetForever)
	sim.CurrentTime = 50 * time.Millisecond
	unit.MultiplyEnergyHaste(sim, 1.2)
	sim.CurrentTime = 150 * time.Millisecond
	unit.energyBar.RunTask(sim)
	unit.MultiplyEnergyHaste(sim, 1/1.2) // Change at a just-completed tick: no second payment.
	sim.CurrentTime = 250 * time.Millisecond
	unit.energyBar.RunTask(sim)
	if got := unit.CurrentEnergy(); math.Abs(got-2.7) > 1e-9 {
		t.Fatalf("energy %v, want .5 + 1.2 + 1 = 2.7", got)
	}
}
