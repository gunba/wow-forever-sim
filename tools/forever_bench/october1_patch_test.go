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
	"github.com/wowsims/classic/sim/druid"
)

func TestOctober1WarriorAndBearCriticalSwingRage(t *testing.T) {
	for _, tc := range []struct {
		key  string
		race proto.Race
		want float64
	}{{"fury", proto.Race_RaceOrc, 2}, {"feral_tank_druid", proto.Race_RaceTauren, 1.75}} {
		req := racialFixture(tc.key, tc.race)
		req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		c := sim.Raid.Parties[0].Players[0].GetCharacter()
		spell := c.AutoAttacks.MHAuto()
		aura := c.GetAura("RageBar")
		clear := func() { c.SpendRage(sim, c.CurrentRage(), c.NewRageMetrics(core.ActionID{SpellID: 1})) }
		clear()
		aura.OnSpellHitDealt(aura, sim, spell, &core.SpellResult{Target: c.CurrentTarget, Outcome: core.OutcomeHit, Damage: 100})
		normal := c.CurrentRage()
		clear()
		critDamage := 100 * spell.CritMultiplier(c.AttackTables[c.CurrentTarget.UnitIndex][spell.CastType])
		aura.OnSpellHitDealt(aura, sim, spell, &core.SpellResult{Target: c.CurrentTarget, Outcome: core.OutcomeCrit, Damage: critDamage})
		if got := c.CurrentRage() / normal; math.Abs(got-tc.want) > 1e-6 {
			t.Errorf("%s critical Rage ratio=%v want=%v", tc.key, got, tc.want)
		}
	}
}

func TestOctober1QueuedAttackDoesNotChangeOffHandMissTable(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	id := core.TernaryInt32(core.IncludeAQ, 25286, 11567)
	queue := c.GetSpell(core.ActionID{SpellID: id, Tag: 1})
	queue.ApplyEffects(sim, c.CurrentTarget, queue)
	if c.PseudoStats.DisableDWMissPenalty {
		t.Fatal("Heroic Strike still gives the off hand a special hit table")
	}
}

func TestOctober1CombustionExpiresOnThirdCrit(t *testing.T) {
	req := racialFixture("fire", proto.Race_RaceHuman)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	combustion := c.GetSpell(core.ActionID{SpellID: 11129})
	combustion.ApplyEffects(sim, c.CurrentTarget, combustion)
	aura := c.GetAura("Combustion")
	fire := c.GetSpell(core.ActionID{SpellID: 10149})
	for i := 1; i <= 3; i++ {
		aura.OnSpellHitDealt(aura, sim, fire, &core.SpellResult{Outcome: core.OutcomeCrit, Target: c.CurrentTarget})
		if aura.IsActive() != (i < 3) {
			t.Fatalf("Combustion active=%v after %d critical strikes", aura.IsActive(), i)
		}
	}
}

func TestOctober1PlagueCritAndInnerFocusScope(t *testing.T) {
	req := racialFixture("shadow", proto.Race_RaceUndead)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	dp := c.GetSpell(core.ActionID{SpellID: 19280})
	before := dp.BonusCritRating
	focus := c.GetSpell(core.ActionID{SpellID: 14751})
	focus.ApplyEffects(sim, c.CurrentTarget, focus)
	if dp.BonusCritRating != before {
		t.Fatal("Inner Focus still adds periodic Plague crit")
	}
	dp.BonusHitRating, dp.BonusCritRating = 10000, 10000
	dp.ApplyEffects(sim, c.CurrentTarget, dp)
	dot := dp.Dot(c.CurrentTarget)
	dot.OnTick(sim, c.CurrentTarget, dot)
	if dp.SpellMetrics[c.CurrentTarget.UnitIndex].CritTicks != 1 {
		t.Fatal("Devouring Plague tick cannot crit")
	}
}

func TestOctober1ShiftingPowerManaEnergyAndTierCooldown(t *testing.T) {
	req := racialFixture("feral", proto.Race_RaceTauren)
	p := req.Raid.Parties[0].Players[0]
	// Move Naturalist's three points into Natural Shapeshifter for this cost check.
	p.TalentsString = p.TalentsString[:len(p.TalentsString)-3] + "05003"
	p.Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	power := c.GetSpell(core.ActionID{SpellID: 1322605})
	if power == nil || c.GetSpell(core.ActionID{SpellID: 9846}) != nil {
		t.Fatal("obsolete Tiger's Fury survived the cutover")
	}
	if power.CD.Duration != 7*time.Second {
		t.Fatalf("Shifting cooldown=%v want 16-8-1 seconds", power.CD.Duration)
	}
	c.SpendEnergy(sim, c.CurrentEnergy(), c.NewEnergyMetrics(core.ActionID{SpellID: 1}))
	mana := c.CurrentMana()
	if !power.Cast(sim, c.CurrentTarget) {
		t.Fatal("Shifting Power cannot cast")
	}
	if math.Abs(c.CurrentEnergy()-40) > 1e-6 {
		t.Fatalf("Shifting returned %v Energy", c.CurrentEnergy())
	}
	want := c.BaseMana * .55 * .7
	if got := mana - c.CurrentMana(); math.Abs(got-want) > 1e-6 {
		t.Fatalf("Shifting spent %v Mana, want %v", got, want)
	}
	power.CD.Timer.Reset()
	c.GCD.Reset()
	sim.Raid.Parties[0].Players[0].(druid.DruidAgent).GetDruid().CancelShapeshift(sim)
	if power.CanCast(sim, c.CurrentTarget) {
		t.Fatal("Shifting Power allowed outside Cat Form")
	}
}

func TestOctober1FuriousPrecisionIsOffHandOnly(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	mh, oh := c.AutoAttacks.MHAuto(), c.AutoAttacks.OHAuto()
	if math.Abs(oh.BonusHitRating-mh.BonusHitRating-10*core.MeleeHitRatingPerHitChance) > 1e-6 {
		t.Fatalf("MH hit=%v OH hit=%v", mh.BonusHitRating, oh.BonusHitRating)
	}
}

func TestOctober1ChampionUsesDamageOnlyIntellectConversion(t *testing.T) {
	req := racialFixture("retribution", proto.Race_RaceHuman)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	before := c.GetStats()
	c.AddStatsDynamic(sim, stats.Stats{stats.Intellect: 100})
	after := c.GetStats()
	want := .6 * (after[stats.Intellect] - before[stats.Intellect])
	if got := after[stats.SpellDamage] - before[stats.SpellDamage]; math.Abs(got-want) > 1e-6 {
		t.Fatalf("Intellect produced %v damage power, want %v", got, want)
	}
	if after[stats.SpellPower] != before[stats.SpellPower] || after[stats.HealingPower] != before[stats.HealingPower] {
		t.Fatal("Champion increased healing-compatible power")
	}
}

func TestOctober1HolyShieldAddsThirtyPercentBlock(t *testing.T) {
	req := racialFixture("protection_paladin", proto.Race_RaceHuman)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	shield := c.GetSpell(core.ActionID{SpellID: 20928})
	before := c.GetStat(stats.Block)
	shield.ApplyEffects(sim, c.CurrentTarget, shield)
	if got := c.GetStat(stats.Block) - before; math.Abs(got-30*core.BlockRatingPerBlockChance) > 1e-6 {
		t.Fatalf("Holy Shield added %v block rating", got)
	}
}

func TestOctober1BloodthirstDoesNotTriggerBloodCraze(t *testing.T) {
	req := racialFixture("fury_2h", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	bt := c.GetSpell(core.ActionID{SpellID: 23894})
	bt.BonusHitRating = 10000
	c.AddStatDynamic(sim, stats.Expertise, 10000)
	trigger := c.GetAura("Blood Craze Trigger")
	if trigger == nil {
		t.Fatal("fixture lacks Blood Craze")
	}
	bt.ApplyEffects(sim, c.CurrentTarget, bt)
	hot := c.GetSpell(core.ActionID{SpellID: 16488}).SelfHot()
	if hot.IsActive() {
		t.Fatal("Bloodthirst still activated Blood Craze")
	}
	trigger.OnSpellHitTaken(trigger, sim, bt, &core.SpellResult{Outcome: core.OutcomeCrit, Damage: 1})
	if !hot.IsActive() {
		t.Fatal("critical incoming damage no longer triggers Blood Craze")
	}
}

func TestOctober1BloodthirstAttackPowerCoefficient(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	p := req.Raid.Parties[0].Players[0]
	p.Rotation = &proto.APLRotation{}
	p.Consumes = &proto.Consumes{}
	p.ForeverTier1Bonuses = false
	req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
	req.Raid.Debuffs = &proto.Debuffs{}
	req.Encounter.Targets[0].Level = 60
	req.Encounter.Targets[0].Stats[stats.Armor] = 0
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	bt := c.GetSpell(core.ActionID{SpellID: 23894})
	bt.BonusHitRating, bt.BonusCritRating = 10000, -10000
	c.AddStatDynamic(sim, stats.Expertise, 10000)
	beforeAP := bt.MeleeAttackPower(c.CurrentTarget)
	bt.ApplyEffects(sim, c.CurrentTarget, bt)
	before := bt.SpellMetrics[c.CurrentTarget.UnitIndex].TotalDamage
	c.AddStatDynamic(sim, stats.AttackPower, 100)
	expected := .45 * (bt.MeleeAttackPower(c.CurrentTarget) - beforeAP)
	expected *= bt.DamageMultiplier * bt.DamageMultiplierAdditive * c.PseudoStats.DamageDealtMultiplier * c.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexPhysical]
	bt.ApplyEffects(sim, c.CurrentTarget, bt)
	got := bt.SpellMetrics[c.CurrentTarget.UnitIndex].TotalDamage - 2*before
	if math.Abs(got-expected) > 1e-6 {
		t.Fatalf("Bloodthirst gained %v damage, want %v from 45%% AP", got, expected)
	}
}
