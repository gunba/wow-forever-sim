package core

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestResourceNoThreatPreservesGainsAndOrdinaryThreat(t *testing.T) {
	targets := []*Target{{Unit: Unit{Type: EnemyUnit, UnitIndex: 0}}, {Unit: Unit{Type: EnemyUnit, UnitIndex: 1}}}
	env := &Environment{Ruleset: proto.Ruleset_RulesetForever, Encounter: Encounter{Targets: targets, TargetUnits: []*Unit{&targets[0].Unit, &targets[1].Unit}}}
	unit := &Unit{Type: PlayerUnit, UnitIndex: 2, Env: env, PseudoStats: stats.NewPseudoStats(), Metrics: NewUnitMetrics()}
	unit.PseudoStats.ThreatMultiplier = 2
	unit.stats[stats.Mana] = 1000
	unit.manaBar.unit = unit
	unit.rageBar.unit, unit.rageBar.maxRage = unit, MaxRage
	mana := &Spell{ActionID: ActionID{OtherID: proto.OtherAction_OtherActionManaGain}, Unit: unit, SpellMetrics: make([]SpellMetrics, 3)}
	rage := &Spell{ActionID: ActionID{OtherID: proto.OtherAction_OtherActionRageGain}, Unit: unit, SpellMetrics: make([]SpellMetrics, 3)}
	unit.Spellbook = []*Spell{mana, rage}
	sim := &Simulation{Environment: env, CurrentTime: time.Second, Options: &proto.SimOptions{Interactive: true}, pendingActions: []*PendingAction{{NextActionAt: NeverExpires}}}

	freeMana := unit.NewManaMetrics(ActionID{SpellID: 11689})
	freeMana.NoThreat = true
	ordinaryMana := unit.NewManaMetrics(ActionID{SpellID: 2})
	unit.AddMana(sim, 100, freeMana)
	unit.AddMana(sim, 40, ordinaryMana)
	unit.manaBar.doneIteration(sim)
	if unit.CurrentMana() != 140 || freeMana.ActualGain != 100 || ordinaryMana.ActualGain != 40 {
		t.Fatal("threat policy changed mana gains")
	}
	for i := 0; i < 2; i++ {
		if got := mana.SpellMetrics[i].TotalThreat; got != 40 {
			t.Errorf("mana target %d: threat %v, want retained ordinary amount 40", i, got)
		}
	}
	freeMana.reset()
	if !freeMana.NoThreat || freeMana.ActualGainForCurrentIteration() != 0 {
		t.Fatal("resource policy did not persist independently of iteration accounting")
	}

	noThreatRage := unit.NewRageMetrics(ActionID{SpellID: 100}) // Tests the accounting contract, not a Charge registration.
	noThreatRage.NoThreat = true
	ordinaryRage := unit.NewRageMetrics(ActionID{SpellID: 2687})
	unit.AddRage(sim, 10, noThreatRage)
	unit.AddRage(sim, 5, ordinaryRage)
	unit.rageBar.doneIteration()
	if unit.CurrentRage() != 15 || noThreatRage.ActualGain != 10 || ordinaryRage.ActualGain != 5 {
		t.Fatal("threat policy changed Rage gains")
	}
	for i := 0; i < 2; i++ {
		if got := rage.SpellMetrics[i].TotalThreat; got != 25 {
			t.Errorf("Rage target %d: threat %v, want retained ordinary amount 25", i, got)
		}
	}
}
