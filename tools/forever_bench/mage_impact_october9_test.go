//go:build with_db

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/mage"
	googleproto "google.golang.org/protobuf/proto"
)

// Current70291 node105792/entry130521/definition135322: curve83091
// is3/7/10, not the UI's old proportional3/6/9 prior. Spell11103 has
// direct-hit proc flags, CAN_PROC_FROM_PROCS and a2s stun12355 (DR group10).
func mageImpactFixture(rank int, eligible, dr bool, ruleset proto.Ruleset, otherMage ...bool) *core.Simulation {
	req := mageSpellFixture(map[string]int{"impact": rank})
	req.Raid.Parties[0].Players[0].ForeverTier1Bonuses = false
	req.Encounter.Targets[0].CanBeStunned = eligible
	req.Encounter.Targets[0].UseCrowdControlDiminishingReturns = dr
	req.SimOptions.Ruleset, req.SimOptions.IsTest, req.SimOptions.RandomSeed = ruleset, true, 12355
	if len(otherMage) > 0 && otherMage[0] {
		other := googleproto.Clone(req.Raid.Parties[0].Players[0]).(*proto.Player)
		other.Name = "Other Mage"
		req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, other)
	}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim
}

func impactAuras(t *testing.T, sim *core.Simulation) (*core.Aura, *core.Aura) {
	t.Helper()
	unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
	trigger := unit.GetAura("Impact Trigger")
	stun := target.GetAura("Impact-" + unit.Label)
	if trigger == nil || trigger.OnSpellHitDealt == nil || stun == nil {
		t.Fatal("Impact is missing from the active talent initialization path")
	}
	if trigger.OnPeriodicDamageDealt != nil || trigger.OnCastComplete != nil {
		t.Fatal("Impact has a periodic or cast-complete trigger")
	}
	return trigger, stun
}

func TestOctober9MageImpactCurve(t *testing.T) {
	for rank, chance := range []float64{.03, .07, .10} {
		t.Run(fmt.Sprintf("rank%d", rank+1), func(t *testing.T) {
			sim := mageImpactFixture(rank+1, true, false, proto.Ruleset_RulesetForever)
			reference := mageImpactFixture(rank+1, true, false, proto.Ruleset_RulesetForever)
			trigger, stun := impactAuras(t, sim)
			unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
			fireball := unit.GetSpell(core.ActionID{SpellID: 133})
			procs := 0
			for sample := 0; sample < 500; sample++ {
				want := reference.Proc(chance, "Impact")
				trigger.OnSpellHitDealt(trigger, sim, fireball, &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
				if stun.IsActive() != want || target.PseudoStats.Stunned != want {
					t.Fatalf("sample%d disagrees with current curve chance%g", sample, chance)
				}
				if want {
					procs++
					if stun.Duration != 2*time.Second || stun.ActionID.SpellID != 12355 {
						t.Fatal("incorrect source stun duration/identity")
					}
					stun.Deactivate(sim)
				}
			}
			if procs == 0 || procs == 500 {
				t.Fatal("Impact chance control was not exercised")
			}
		})
	}
}

func TestOctober9MageImpactDirectScope(t *testing.T) {
	sim := mageImpactFixture(3, true, false, proto.Ruleset_RulesetForever)
	trigger, stun := impactAuras(t, sim)
	unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
	fireball := unit.GetSpell(core.ActionID{SpellID: 133})
	for _, outcome := range []core.HitOutcome{core.OutcomeMiss, core.OutcomeDodge} {
		for sample := 0; sample < 100; sample++ {
			trigger.OnSpellHitDealt(trigger, sim, fireball, &core.SpellResult{Target: target, Outcome: outcome})
		}
		if stun.IsActive() {
			t.Fatal("unlanded fire spell triggered Impact")
		}
	}
	for _, id := range []int32{116, 5143} {
		spell := unit.GetSpell(core.ActionID{SpellID: id})
		for sample := 0; sample < 100; sample++ {
			trigger.OnSpellHitDealt(trigger, sim, spell, &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
		}
		if stun.IsActive() {
			t.Fatal("non-Fire spell triggered Impact")
		}
	}
	// A foreign item proc is not a Mage-family spell. In contrast, the client's
	// CAN_PROC_FROM_PROCS bit permits eligible direct Mage-family fire procs.
	for _, mageFamily := range []bool{false, true} {
		flags := core.SpellFlagPassiveSpell
		if mageFamily {
			flags |= mage.SpellFlagMage
		}
		spell := unit.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 12355, Tag: int32(10 + core.TernaryInt(mageFamily, 1, 0))}, SpellSchool: core.SpellSchoolFire, ProcMask: core.ProcMaskSpellDamage | core.ProcMaskSpellProc, Flags: flags})
		for sample := 0; sample < 100; sample++ {
			trigger.OnSpellHitDealt(trigger, sim, spell, &core.SpellResult{Target: target, Outcome: core.OutcomeCrit})
		}
		if stun.IsActive() != mageFamily {
			t.Fatal("direct proc eligibility ignored source Mage family/CAN_PROC_FROM_PROCS")
		}
		stun.Deactivate(sim)
	}
	flamestrike := unit.GetSpell(core.ActionID{SpellID: 2120})
	dot := flamestrike.AOEDot()
	dot.Apply(sim)
	for tick := 0; tick < 100; tick++ {
		dot.OnTick(sim, target, dot)
	}
	if stun.IsActive() || target.PseudoStats.Stunned {
		t.Fatal("periodic Flamestrike damage triggered Impact")
	}
	if flamestrike.SpellMetrics[target.UnitIndex].TotalDamage <= 0 {
		t.Fatal("periodic control did not actually deal damage")
	}
	for _, id := range []int32{133, 2136, 2948, 2120, 1237313} {
		spell := unit.GetSpell(core.ActionID{SpellID: id})
		for sample := 0; sample < 100 && !stun.IsActive(); sample++ {
			trigger.OnSpellHitDealt(trigger, sim, spell, &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
		}
		if !stun.IsActive() {
			t.Fatalf("direct fire spell%d cannot proc Impact", id)
		}
		stun.Deactivate(sim)
	}
}

func TestOctober9MageImpactImmunityAndClassic(t *testing.T) {
	sim := mageImpactFixture(3, false, false, proto.Ruleset_RulesetForever)
	trigger, stun := impactAuras(t, sim)
	unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
	for sample := 0; sample < 500; sample++ {
		trigger.OnSpellHitDealt(trigger, sim, unit.GetSpell(core.ActionID{SpellID: 133}), &core.SpellResult{Target: target, Outcome: core.OutcomeCrit})
	}
	if stun.IsActive() || target.PseudoStats.Stunned {
		t.Fatal("default PvE target was stunned without opt-in")
	}
	for _, pair := range []struct {
		rank    int
		ruleset proto.Ruleset
	}{{0, proto.Ruleset_RulesetForever}, {3, proto.Ruleset_RulesetClassic}} {
		sim := mageImpactFixture(pair.rank, true, false, pair.ruleset)
		if sim.Raid.AllPlayerUnits[0].GetAura("Impact Trigger") != nil {
			t.Fatal("Impact changed zero-rank or retained Classic initialization")
		}
	}
}

func TestOctober9MageImpactSharedDR(t *testing.T) {
	sim := mageImpactFixture(3, true, true, proto.Ruleset_RulesetForever, true)
	trigger, stun := impactAuras(t, sim)
	unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
	fireball := unit.GetSpell(core.ActionID{SpellID: 133})
	other := target.GetAura("Impact-" + sim.Raid.AllPlayerUnits[1].Label)
	if other == nil {
		t.Fatal("other caster's initialized Impact aura missing")
	}
	if !target.ApplyCrowdControl(sim, other, 2*time.Second, core.CrowdControlDRStunProc) {
		t.Fatal("initial shared-group control failed")
	}
	proc := func() {
		for sample := 0; sample < 500 && !stun.IsActive(); sample++ {
			trigger.OnSpellHitDealt(trigger, sim, fireball, &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
		}
	}
	for _, duration := range []time.Duration{time.Second, 500 * time.Millisecond} {
		proc()
		if !stun.IsActive() || stun.Duration != duration {
			t.Fatal("Impact did not share Stun10 diminution with another effect/caster")
		}
		stun.Deactivate(sim)
		if !target.PseudoStats.Stunned {
			t.Fatal("Impact expiration erased overlapping external stun")
		}
	}
	proc()
	if stun.IsActive() {
		t.Fatal("fourth shared-group application refreshed an immune Impact")
	}
	other.Deactivate(sim)
	if target.PseudoStats.Stunned {
		t.Fatal("final overlapping stun did not release")
	}
	// Shared helper's declared opt-in convention: reset15s after final group end.
	sim.CurrentTime = 15 * time.Second
	proc()
	if !stun.IsActive() || stun.Duration != 2*time.Second {
		t.Fatal("Impact did not use shared DR reset")
	}
	stun.Deactivate(sim)
}
