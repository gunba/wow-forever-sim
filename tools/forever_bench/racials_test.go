//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func racialFixture(key string, race proto.Race) *proto.RaidSimRequest {
	for _, b := range builds() {
		if b.Key != key {
			continue
		}
		p := b.player(race)
		p.Buffs = &proto.IndividualBuffs{}
		p.Consumes = &proto.Consumes{}
		req := request(p, 1, 1)
		req.Raid.Buffs = &proto.RaidBuffs{}
		req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
		req.Raid.Debuffs = &proto.Debuffs{}
		return req
	}
	panic(key)
}

func TestForeverBloodFury(t *testing.T) {
	req := racialFixture("arcane", proto.Race_RaceOrc)
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	unit := env.Raid.AllPlayerUnits[0]
	sim := &core.Simulation{Environment: env}
	spell := unit.GetSpell(core.ActionID{SpellID: 20572})
	if spell.DefaultCast.GCD != 0 || spell.Cost != nil || spell.CD.Duration != 2*time.Minute {
		t.Fatal("Blood Fury must be free, off the GCD, on a two-minute cooldown")
	}
	before := unit.GetStats()
	aura := unit.GetAura("Blood Fury")
	unit.AddStatDynamic(sim, stats.FirePower, 100)
	aura.Activate(sim)
	for _, stat := range []stats.Stat{stats.AttackPower, stats.RangedAttackPower, stats.SpellPower} {
		if math.Abs(unit.GetStat(stat)-before[stat]*1.1) > 1e-7 {
			t.Errorf("stat %v did not gain 10%%", stat)
		}
	}
	if math.Abs(unit.GetStat(stats.FirePower)-(before[stats.FirePower]+100)*1.1) > 1e-7 {
		t.Fatal("school-specific spell power did not gain 10%")
	}
	// A percentage aura must also affect stats gained after activation.
	unit.AddStatDynamic(sim, stats.SpellPower, 100)
	if math.Abs(unit.GetStat(stats.SpellPower)-(before[stats.SpellPower]+100)*1.1) > 1e-7 {
		t.Fatal("Blood Fury snapshotted spell power")
	}
	aura.Deactivate(sim)
	unit.AddStatDynamic(sim, stats.SpellPower, -100)
	unit.AddStatDynamic(sim, stats.FirePower, -100)
	for stat, value := range before {
		if math.Abs(unit.GetStat(stats.Stat(stat))-value) > 1e-7 {
			t.Errorf("stat %v did not reverse", stat)
		}
	}
}

func TestForeverBerserking(t *testing.T) {
	for _, key := range []string{"arcane", "fury", "combat"} {
		t.Run(key, func(t *testing.T) {
			req := racialFixture(key, proto.Race_RaceTroll)
			env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
			unit := env.Raid.AllPlayerUnits[0]
			sim := &core.Simulation{Environment: env}
			spell := unit.GetSpell(core.ActionID{SpellID: 20554})
			if spell == nil || spell.Cost != nil || spell.DefaultCast.GCD != 0 || spell.CD.Duration != 3*time.Minute {
				t.Fatal("Berserking must be free and off the GCD for every resource type")
			}
			before := unit.PseudoStats
			aura := unit.GetAura("Berserking")
			aura.Activate(sim)
			for _, ratio := range []float64{
				unit.PseudoStats.MeleeSpeedMultiplier / before.MeleeSpeedMultiplier,
				unit.PseudoStats.RangedSpeedMultiplier / before.RangedSpeedMultiplier,
				unit.PseudoStats.CastSpeedMultiplier / before.CastSpeedMultiplier,
			} {
				if math.Abs(ratio-1.1) > 1e-7 {
					t.Errorf("speed multiplier = %v, want 1.1", ratio)
				}
			}
			aura.Deactivate(sim)
			if math.Abs(unit.PseudoStats.MeleeSpeedMultiplier-before.MeleeSpeedMultiplier) > 1e-7 {
				t.Fatal("Berserking did not reverse")
			}
		})
	}
}

func TestForeverWindshaperHasNoPowerCooldown(t *testing.T) {
	req := racialFixture("balance", proto.Race_RaceSkyborneWindshaper)
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	unit := env.Raid.AllPlayerUnits[0]
	if unit.GetSpell(core.ActionID{SpellID: 460530}) != nil || unit.GetAura("Windshaper") != nil {
		t.Fatal("unsupported Windshaper power cooldown remains registered")
	}
	if math.Abs(unit.PseudoStats.CastSpeedMultiplier-1.01) > 1e-7 {
		t.Fatal("Wind Blessed's passive haste is missing")
	}
}

func TestForeverTouchOfTheGrave(t *testing.T) {
	for _, key := range []string{"arcane", "fury", "retribution", "combat", "shadow", "affliction"} {
		req := racialFixture(key, proto.Race_RaceUndead)
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		unit := env.Raid.AllPlayerUnits[0]
		aura := unit.GetAura("Touch of the Grave")
		wantID := int32(1260189)
		if key == "arcane" || key == "shadow" || key == "affliction" {
			wantID = 1260201
		}
		if aura.ActionID.SpellID != wantID || aura.Icd == nil || aura.Icd.Duration != time.Second {
			t.Errorf("%s: incorrect racial variant or ICD", key)
		}
	}

	req := racialFixture("arcane", proto.Race_RaceUndead)
	p := req.Raid.Parties[0].Players[0]
	p.TalentsString = ""
	p.Rotation = core.APLRotationFromJsonString(`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":133}}}}]}`)
	req.SimOptions.Iterations = 500
	req.Encounter.Duration = 60
	req.Encounter.Targets[0].Level = 60
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	maxHealth := env.Raid.AllPlayerUnits[0].MaxHealth()
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error.Message)
	}
	var sourceHits, procs float64
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		for _, m := range action.Targets {
			switch action.Id.GetSpellId() {
			case 133:
				sourceHits += float64(m.Hits + m.Crits)
			case 1260198:
				procs += float64(m.Hits + m.Misses)
				if m.Crits != 0 || math.Abs(m.Damage-float64(m.Hits)*maxHealth*.05) > 1e-5 {
					t.Fatalf("drain damage must be 5%% of max health without crits: %+v, health %v", m, maxHealth)
				}
			}
		}
	}
	if sourceHits == 0 || math.Abs(procs/sourceHits-.1) > .01 {
		t.Fatalf("caster racial proc rate = %v / %v", procs, sourceHits)
	}
}

func TestForeverDarkSacrifice(t *testing.T) {
	for _, race := range []proto.Race{proto.Race_RaceUndead, proto.Race_RaceTroll} {
		req := racialFixture("smite", race)
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		unit := env.Raid.AllPlayerUnits[0]
		spell := unit.GetSpell(core.ActionID{SpellID: 1277328})
		if race == proto.Race_RaceTroll {
			if spell != nil {
				t.Fatal("Dark Sacrifice must be restricted to Undead")
			}
			continue
		}
		if spell == nil || spell.Cost != nil || spell.DefaultCast.GCD != core.GCDDefault ||
			spell.CD.Duration != 10*time.Minute {
			t.Fatal("incorrect Dark Sacrifice cost/GCD/cooldown")
		}
		if unit.GetSpell(core.ActionID{SpellID: 1277324}).CD.Timer != spell.CD.Timer {
			t.Fatal("Dark Sacrifice ranks must share one cooldown")
		}
	}

	// A full simulation exercises tick timing, mana gains, health costs, and
	// pending-action cleanup between iterations. The 60-only rank gives 320
	// per tick, not the level-68 value rendered by Wowhead's static preview.
	req := racialFixture("smite", proto.Race_RaceUndead)
	req.Encounter.Duration = 20
	req.SimOptions.Iterations = 2
	req.Raid.Parties[0].Players[0].TalentsString = ""
	req.Raid.Parties[0].Players[0].Rotation = core.APLRotationFromJsonString(`{
		"type":"TypeAPL","priorityList":[
			{"action":{"castSpell":{"spellId":{"spellId":1277328}}}},
			{"action":{"castSpell":{"spellId":{"spellId":10934}}}}
		]}`)
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error.Message)
	}
	var mana, healthCost float64
	var ticks int32
	for _, resource := range result.RaidMetrics.Parties[0].Players[0].Resources {
		if resource.Id.GetSpellId() == 1277328 {
			mana += resource.Gain
			ticks += resource.Events
		}
		if resource.Id.GetOtherId() == proto.OtherAction_OtherActionDamageTaken &&
			resource.Type == proto.ResourceType_ResourceTypeHealth {
			healthCost -= resource.Gain
		}
	}
	if mana != 3200 || healthCost != 3200 || ticks != 10 {
		t.Fatalf("Dark Sacrifice transfer: mana %v, health %v, ticks %v", mana, healthCost, ticks)
	}
}
