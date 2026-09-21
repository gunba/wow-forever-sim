//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
)

func mageSpellFixture(pointsByName map[string]int) *proto.RaidSimRequest {
	req := racialFixture("arcane", proto.Race_RaceOrc)
	p := req.Raid.Parties[0].Players[0]
	p.Equipment = &proto.EquipmentSpec{}
	p.Rotation = &proto.APLRotation{}
	for _, b := range builds() {
		if b.Key != "arcane" {
			continue
		}
		config := loadTalents(b)
		points, _ := config.decode("")
		index := 0
		for _, tree := range config.Trees {
			for _, talent := range tree.Talents {
				points[index] = pointsByName[talent.Field]
				index++
			}
		}
		// Isolate the named modifiers, without spending unrelated filler points.
		// This fixture is not an optimizer candidate.
		p.TalentsString = config.encode(points)
	}
	req.Encounter.Targets[0].Level = 60
	return req
}

func TestForeverFrostfireBolt(t *testing.T) {
	req := mageSpellFixture(nil)
	sim := core.NewSim(req, simsignals.Signals{})
	unit := sim.Raid.AllPlayerUnits[0]
	target := sim.Encounter.TargetUnits[0]
	spell := unit.GetSpell(core.ActionID{SpellID: 1237313})
	if spell == nil || spell.SpellSchool != core.SpellSchoolFire|core.SpellSchoolFrost ||
		spell.DefaultCast.CastTime != 3*time.Second || spell.Cost.BaseCost != 370 ||
		spell.BonusCoefficient != .814 || spell.MissileSpeed != 24 {
		t.Fatal("incorrect Frostfire Bolt registration")
	}
	before := spell.CalcDamage(sim, target, 270, spell.OutcomeAlwaysHit).Damage
	unit.AddStatDynamic(sim, stats.SpellPower, 100)
	after := spell.CalcDamage(sim, target, 270, spell.OutcomeAlwaysHit).Damage
	if math.Abs(after-before-81.4) > 1e-7 {
		t.Fatalf("direct spell-power contribution = %v, want 81.4", after-before)
	}
	dot := spell.Dot(target)
	dot.Snapshot(target, 19, false)
	if dot.NumberOfTicks != 3 || dot.TickLength != 3*time.Second ||
		dot.SnapshotBaseDamage != 19 {
		t.Fatal("Frostfire DoT must be three 19-point ticks without SP scaling")
	}
	// School-specific power must use the higher applicable school, not their sum.
	unit.AddStatDynamic(sim, stats.FirePower, 80)
	unit.AddStatDynamic(sim, stats.FrostPower, 120)
	if math.Abs(spell.GetBonusDamage(target)-220) > 1e-7 {
		t.Fatal("Frostfire double-counted school spell power")
	}
}

func TestForeverFrostfireTalents(t *testing.T) {
	req := mageSpellFixture(map[string]int{
		"improvedFireball": 5, "improvedFrostbolt": 5,
		"elementalPrecision": 3, "iceShards": 5, "firePower": 5,
		"piercingIce": 3, "frostChanneling": 3,
		"hotStreak": 1, "missileBarrage": 1, "wintersChill": 5,
	})
	sim := core.NewSim(req, simsignals.Signals{})
	unit := sim.Raid.AllPlayerUnits[0]
	spell := unit.GetSpell(core.ActionID{SpellID: 1237313})
	if spell.DefaultCast.CastTime != 2500*time.Millisecond {
		t.Fatal("only Improved Fireball should shorten Frostfire Bolt")
	}
	if spell.BonusHitRating != 3*core.SpellHitRatingPerHitChance ||
		spell.CritDamageBonus != 2 || spell.Cost.Multiplier != 85 ||
		math.Abs(spell.DamageMultiplierAdditive-1.16) > 1e-7 {
		t.Fatalf("incorrect Fire/Frost modifiers: hit %v, crit bonus %v, cost %v, damage %v",
			spell.BonusHitRating, spell.CritDamageBonus, spell.Cost.Multiplier, spell.DamageMultiplierAdditive)
	}
	critBefore := spell.BonusCritRating
	wc := unit.GetAura("Winter's Chill")
	wc.Activate(sim)
	wc.SetStacks(sim, 5)
	if spell.BonusCritRating != critBefore {
		t.Fatal("Winter's Chill only grants crit to Frostbolt and Ice Lance")
	}
	unit.GetAura("Hot Streak Trigger").OnSpellHitDealt(
		unit.GetAura("Hot Streak Trigger"), sim, spell,
		&core.SpellResult{Target: sim.Encounter.TargetUnits[0], Outcome: core.OutcomeCrit},
	)
	if unit.GetAura("Hot Streak").GetStacks() != 1 {
		t.Fatal("Frostfire direct crit did not grant Hot Streak")
	}
	trigger := unit.GetAura("Missile Barrage Trigger")
	for i := 0; i < 100 && !unit.GetAura("Missile Barrage").IsActive(); i++ {
		trigger.OnCastComplete(trigger, sim, spell)
	}
	if !unit.GetAura("Missile Barrage").IsActive() {
		t.Fatal("Frostfire casts did not trigger Missile Barrage")
	}
}

func TestFrostboltRecordsMisses(t *testing.T) {
	req := mageSpellFixture(nil)
	req.Encounter.Targets[0].Level = 63
	req.Encounter.Duration = 6
	req.SimOptions.Iterations = 100
	req.Raid.Parties[0].Players[0].Rotation = core.APLRotationFromJsonString(
		`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":116}}}}]}`)
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error.Message)
	}
	var misses int32
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		if action.Id.GetSpellId() == 116 {
			for _, target := range action.Targets {
				misses += target.Misses
			}
		}
	}
	if misses == 0 {
		t.Fatal("Frostbolt misses disappeared from action metrics")
	}
}

func TestForeverFrostfireRuntime(t *testing.T) {
	req := mageSpellFixture(map[string]int{"ignite": 5, "hotStreak": 1})
	req.Encounter.Duration = 30
	req.SimOptions.Iterations = 100
	req.Raid.Parties[0].Players[0].Rotation = core.APLRotationFromJsonString(
		`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":1237313}}}}]}`)
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error.Message)
	}
	var direct, periodic, ignites int32
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		for _, target := range action.Targets {
			switch action.Id.GetSpellId() {
			case 1237313:
				direct += target.Hits + target.Crits
				periodic += target.Ticks + target.CritTicks
			case 12654:
				ignites += target.Ticks
			}
		}
	}
	if direct == 0 || periodic == 0 || ignites == 0 {
		t.Fatalf("Frostfire runtime missing direct (%d), periodic (%d), or Ignite (%d) damage", direct, periodic, ignites)
	}
}
