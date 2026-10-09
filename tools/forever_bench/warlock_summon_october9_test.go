//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/warlock"
)

func summonOctober9Fixture() (*core.Simulation, *warlock.Warlock) {
	req := warlockEffectFixture(map[string]int{"improvedImp": 3})
	req.Raid.Parties[0].Players[0].GetWarlock().Options.Summon = proto.WarlockOptions_Imp
	req.SimOptions.IsTest = true
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim, sim.Raid.Parties[0].Players[0].(warlock.WarlockAgent).GetWarlock()
}

func TestOctober9SummonRejectedCastKeepsPet(t *testing.T) {
	sim, w := summonOctober9Fixture()
	old := w.ActivePet
	w.SetGCDTimer(sim, time.Second)
	if w.SummonDemonSpells[0].Cast(sim, w.CurrentTarget) {
		t.Fatal("cast succeeded during GCD")
	}
	if w.ActivePet != old || !old.IsEnabled() || old.PseudoStats.Stunned {
		t.Fatal("rejected summon dismissed or stunned the existing pet")
	}
}

func advanceSummonOctober9(t *testing.T, sim *core.Simulation, until time.Duration) {
	t.Helper()
	core.StartDelayedAction(sim, core.DelayedActionOptions{DoAt: until, OnAction: func(*core.Simulation) {}})
	for sim.CurrentTime < until {
		if sim.Step() {
			t.Fatal("fight ended during summon fixture")
		}
	}
}

func TestOctober9SummonCancelAndCompletion(t *testing.T) {
	sim, w := summonOctober9Fixture()
	old, hp, mana := w.ActivePet, w.MaxHealth(), w.CurrentMana()
	cast := w.SummonDemonSpells[0]
	if !cast.Cast(sim, w.CurrentTarget) {
		t.Fatal("summon failed")
	}
	if w.MaxHealth() != hp {
		t.Fatal("summon stun removed owned Blood Pact")
	}
	advanceSummonOctober9(t, sim, 500*time.Millisecond)
	w.InterruptCast(sim)
	if w.ActivePet != old || !old.IsEnabled() || old.PseudoStats.Stunned || w.CurrentMana() != mana || w.GCD.IsReady(sim) {
		t.Fatal("cancel must restore the exact old pet, preserve unexpired GCD and spend no completion mana")
	}
	advanceSummonOctober9(t, sim, 2*time.Second)
	if !cast.Cast(sim, w.CurrentTarget) {
		t.Fatal("repeat summon failed")
	}
	finish := w.Hardcast.Expires
	advanceSummonOctober9(t, sim, finish)
	if w.ActivePet != w.Felhunter || !w.Felhunter.IsEnabled() || old.IsEnabled() || old.PseudoStats.Stunned {
		t.Fatal("completed summon did not replace exactly once")
	}
	if got := w.MaxMana() - w.CurrentMana(); math.Abs(got-cast.CurCast.Cost) > 1e-6 {
		t.Fatalf("summon cost %g, want %g", got, cast.CurCast.Cost)
	}
	w.InterruptCast(sim)
	if old.IsEnabled() {
		t.Fatal("late cancellation restarted replaced pet")
	}
	w.AddMana(sim, w.MaxMana(), w.NewManaMetrics(core.ActionID{SpellID: 688}))
	if !w.SummonDemonSpells[1].Cast(sim, w.CurrentTarget) {
		t.Fatal("return summon failed")
	}
	advanceSummonOctober9(t, sim, w.Hardcast.Expires)
	if w.ActivePet != old || !old.IsEnabled() || w.MaxHealth() != hp {
		t.Fatal("resummoning Imp duplicated/lost Blood Pact")
	}
	sim.Cleanup()
	sim.Reset()
	if w.ActivePet != old || old.PseudoStats.Stunned {
		t.Fatal("summon state leaked across reset")
	}
}

func TestOctober9SummonNestedStunAndPetCastCancel(t *testing.T) {
	sim, w := summonOctober9Fixture()
	old := w.ActivePet
	firebolt := old.GetSpell(core.ActionID{SpellID: 11763})
	if firebolt == nil || !firebolt.Cast(sim, old.CurrentTarget) {
		t.Fatal("Imp hardcast unavailable")
	}
	old.AddStun(sim)
	if old.IsCasting(sim) || firebolt.CanCast(sim, old.CurrentTarget) || firebolt.Cast(sim, old.CurrentTarget) {
		t.Fatal("stun did not interrupt/gate pet cast")
	}
	if !w.SummonDemonSpells[0].Cast(sim, w.CurrentTarget) {
		t.Fatal("summon failed")
	}
	w.InterruptCast(sim)
	if !old.PseudoStats.Stunned || !old.IsEnabled() {
		t.Fatal("summon cancellation erased another stun")
	}
	old.RemoveStun(sim)
	if old.PseudoStats.Stunned {
		t.Fatal("last stun did not release")
	}
	old.PseudoStats.Stunned = true
	old.AddStun(sim)
	old.RemoveStun(sim)
	if !old.PseudoStats.Stunned {
		t.Fatal("preexisting stun flag erased")
	}
}

func TestOctober9SummonAlreadyDisabledPet(t *testing.T) {
	sim,w:=summonOctober9Fixture()
	old:=w.ActivePet
	disabled:=0
	old.ApplyOnPetDisable(func(*core.Simulation,bool){disabled++})
	if !w.SummonDemonSpells[0].Cast(sim,w.CurrentTarget) {t.Fatal("summon failed")}
	old.Disable(sim,false)
	advanceSummonOctober9(t,sim,w.Hardcast.Expires)
	if disabled!=1 || old.IsEnabled() || old.PseudoStats.Stunned || w.ActivePet!=w.Felhunter {t.Fatalf("old pet lifecycle repeated or restored: disable callbacks=%d",disabled)}
}

func TestOctober9SummonKeepsExistingPetWhileCasting(t *testing.T) {
	sim, w := summonOctober9Fixture()
	old := w.ActivePet
	if !w.SummonDemonSpells[0].Cast(sim, w.CurrentTarget) {
		t.Fatal("summon failed")
	}
	if w.ActivePet != old || !old.IsEnabled() || !old.PseudoStats.Stunned {
		t.Fatal("summoning must stun the existing pet, not dismiss it")
	}
}
