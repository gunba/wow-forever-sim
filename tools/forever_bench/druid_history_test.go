//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/druid"
)

func TestNaturesGraceGCDIncludesInstantSpells(t *testing.T) {
	req := historyTalentFixture("balance", map[string]int{"naturesGrace": 1, "insectSwarm": 1, "moonkinForm": 1})
	req.Raid.Parties[0].Players[0].ForeverTier1Bonuses = false
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	unit := sim.Raid.AllPlayerUnits[0]
	aura := unit.GetAura("Natures Grace")
	ids := []int32{5176, 2912, 8921, 5570, 16914, 9907}
	before := make(map[int32]time.Duration)
	for _, id := range ids {
		before[id] = unit.GetSpell(core.ActionID{SpellID: id}).DefaultCast.GCD
	}
	form := unit.GetSpell(core.ActionID{SpellID: 24858})
	formGCD := form.DefaultCast.GCD
	castSpeed := unit.PseudoStats.CastSpeedMultiplier
	aura.Activate(sim)
	for _, id := range ids {
		if gcd := unit.GetSpell(core.ActionID{SpellID: id}).DefaultCast.GCD; gcd != before[id]-before[id]/10 {
			t.Fatalf("spell %d GCD %v: want exactly 10%% reduction", id, gcd)
		}
	}
	if form.DefaultCast.GCD != formGCD || math.Abs(unit.PseudoStats.CastSpeedMultiplier/castSpeed-1.1) > 1e-9 {
		t.Fatal("form GCD changed or cast haste is incorrect")
	}
	moonfire := unit.GetSpell(core.ActionID{SpellID: 8921})
	if !moonfire.Cast(sim, unit.CurrentTarget) || moonfire.CurCast.GCD != 1350*time.Millisecond {
		t.Fatal("instant Moonfire did not use its reduced GCD")
	}
	aura.Deactivate(sim)
	for _, id := range ids {
		if gcd := unit.GetSpell(core.ActionID{SpellID: id}).DefaultCast.GCD; gcd != before[id] {
			t.Fatal("GCD reduction leaked after expiry")
		}
	}
}

func TestMangleTalentDoesNotInventCatSpell(t *testing.T) {
	req := historyTalentFixture("feral", map[string]int{"mangle": 1})
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	for _, id := range []int32{33876, 407993} {
		if env.Raid.AllPlayerUnits[0].GetSpell(core.ActionID{SpellID: id}) != nil {
			t.Fatalf("unsupported Cat Mangle %d was registered", id)
		}
	}
}

func TestBerserkCleaveOnlyAffectsMangle(t *testing.T) {
	req := historyTalentFixture("feral", map[string]int{"mangle": 1, "berserk": 1})
	p := req.Raid.Parties[0].Players[0]
	p.ForeverTier1Bonuses = false
	p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{Options: &proto.FeralTankDruid_Options{}}}
	req.Encounter.Targets = append(req.Encounter.Targets, req.Encounter.Targets[0], req.Encounter.Targets[0])
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	bear := sim.Raid.Parties[0].Players[0].(interface{ GetDruid() *druid.Druid }).GetDruid()
	bear.BerserkAura.Activate(sim)
	for i := 0; i < 5; i++ {
		bear.Lacerate.ApplyEffects(sim, bear.CurrentTarget, bear.Lacerate.Spell)
		bear.MangleBear.ApplyEffects(sim, bear.CurrentTarget, bear.MangleBear.Spell)
	}
	for i, target := range sim.Encounter.TargetUnits {
		if bear.MangleBear.SpellMetrics[target.UnitIndex].TotalDamage == 0 {
			t.Fatal("Berserk did not allow Mangle to cleave")
		}
		if i > 0 && (bear.Lacerate.SpellMetrics[target.UnitIndex].TotalDamage != 0 ||
			bear.LacerateBleed.Dot(target).IsActive()) {
			t.Fatal("Berserk incorrectly made Lacerate cleave")
		}
	}
}
