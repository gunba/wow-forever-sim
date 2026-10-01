//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/hunter"
	googleProto "google.golang.org/protobuf/proto"
	"time"
)

func TestExposePreyMarkOwnership(t *testing.T) {
	req := racialFixture("survival", proto.Race_RaceOrc)
	p := req.Raid.Parties[0].Players[0]
	p.Rotation = &proto.APLRotation{}
	req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, googleProto.Clone(p).(*proto.Player))
	req.Raid.Debuffs.HuntersMark = proto.TristateEffect_TristateEffectRegular
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	a := sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
	b := sim.Raid.Parties[0].Players[1].(*hunter.Hunter)
	target := sim.Encounter.TargetUnits[0]
	check := func(h *hunter.Hunter, want bool) {
		proc := h.GetAura("Expose Prey")
		count := 0
		for i := 0; i < 500; i++ {
			h.DefensiveState.Deactivate(sim)
			proc.OnSpellHitDealt(proc, sim, &core.Spell{ProcMask: core.ProcMaskMeleeMHAuto}, &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
			if h.DefensiveState.IsActive() {
				count++
			}
		}
		if (count > 0) != want {
			t.Errorf("caster%d: external/other-owner mark eligibility incorrect: %d procs", h.UnitIndex, count)
		}
	}
	check(a, false)
	check(b, false)
	// External permanent support is a scenario provider, not this caster's mark.
	// Clear it before testing the legal self-maintained provider.
	target.GetAuraByID(core.ActionID{SpellID: 14325}).Deactivate(sim)
	markA := a.GetSpell(core.ActionID{SpellID: 14325})
	markB := b.GetSpell(core.ActionID{SpellID: 14325})
	if markA == nil || markB == nil {
		t.Fatal("Hunter's Mark not registered for ownership")
	}
	if markA.DefaultCast.GCD != 1500*time.Millisecond || markA.RequiredLevel != 58 {
		t.Fatal("Mark GCD/level mismatch")
	}
	before := a.CurrentMana()
	if !markA.Cast(sim, target) {
		t.Fatal("own mark failed")
	}
	if before-a.CurrentMana() != 60 {
		t.Fatal("Mark did not pay60Mana")
	}
	check(a, true)
	check(b, false)
	if !markB.Cast(sim, target) {
		t.Fatal("second owner's mark failed")
	}
	check(a, false)
	check(b, true)
	if aura := target.GetAuraByID(core.ActionID{SpellID: 14325, Tag: b.Index + 1}); aura == nil || aura.Duration != 2*time.Minute {
		t.Fatal("owned mark metadata missing")
	}
}

func TestHunterOwnedMarkDefaultDuty(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "survival" && b.Key != "pet_melee" {
			continue
		}
		p := prepare(b, b.races()[0])
		req := requestForBuild(b, p, 10, 20261002)
		if req.Raid.Debuffs.HuntersMark != proto.TristateEffect_TristateEffectMissing {
			t.Fatal("external mark still masks the owned duty")
		}
		statsResult := core.ComputeStats(&proto.ComputeStatsRequest{Raid: req.Raid, Encounter: req.Encounter, Ruleset: proto.Ruleset_RulesetForever})
		playerStats := statsResult.RaidStats.Parties[0].Players[0]
		for _, row := range append(playerStats.RotationStats.PrepullActions, playerStats.RotationStats.PriorityList...) {
			if len(row.Warnings) > 0 {
				t.Fatal(row.Warnings)
			}
		}
		result := core.RunRaidSim(requestForBuild(b, p, 10, 20261002))
		if result.Error != nil {
			t.Fatal(result.Error)
		}
		found := false
		for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
			if action.Id.GetSpellId() == 14325 {
				for _, target := range action.Targets {
					if target.Casts >= 20 {
						found = true
					}
				}
			}
		}
		if !found {
			t.Fatal("default did not cast and refresh its own Mark")
		}
	}
}

func TestExposePreyAttackMask(t *testing.T) {
	req := racialFixture("survival", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	req.Raid.Debuffs.HuntersMark = proto.TristateEffect_TristateEffectRegular
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	h := sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
	if h.Talents.ExposePrey != 2 {
		t.Fatal("fixture requires 2/2 Expose Prey")
	}
	trigger := h.GetAura("Expose Prey")
	target := sim.Encounter.TargetUnits[0]
	target.GetAuraByID(core.ActionID{SpellID: 14325}).Deactivate(sim)
	h.GetSpell(core.ActionID{SpellID: 14325}).Cast(sim, target)
	for _, tc := range []struct {
		name string
		mask core.ProcMask
		hit  core.HitOutcome
		want bool
	}{
		{"melee auto", core.ProcMaskMeleeMHAuto, core.OutcomeHit, true},
		{"ranged auto", core.ProcMaskRangedAuto, core.OutcomeHit, true},
		{"melee special", core.ProcMaskMeleeMHSpecial, core.OutcomeHit, true},
		{"ranged special", core.ProcMaskRangedSpecial, core.OutcomeHit, true},
		{"trap spell", core.ProcMaskSpellDamage, core.OutcomeHit, false},
		{"proc", core.ProcMaskSpellDamageProc, core.OutcomeHit, false},
		{"miss", core.ProcMaskMeleeMHAuto, core.OutcomeMiss, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			procs := 0
			for i := 0; i < 500; i++ {
				h.DefensiveState.Deactivate(sim)
				trigger.OnSpellHitDealt(trigger, sim, &core.Spell{ProcMask: tc.mask}, &core.SpellResult{Target: target, Outcome: tc.hit})
				if h.DefensiveState.IsActive() {
					procs++
				}
			}
			if (procs > 0) != tc.want {
				t.Fatalf("procs = %d; eligible = %v", procs, tc.want)
			}
		})
	}
}
