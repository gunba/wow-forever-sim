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
	"github.com/wowsims/classic/sim/shaman"
	googleProto "google.golang.org/protobuf/proto"
)

func shamanOctober8Fixture(talents map[string]int, ruleset proto.Ruleset) (*core.Simulation, *shaman.Shaman) {
	req := historyTalentFixture("elemental", talents)
	player := req.Raid.Parties[0].Players[0]
	player.Equipment = &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 17)}
	for i := range player.Equipment.Items {
		player.Equipment.Items[i] = &proto.ItemSpec{}
	}
	player.ForeverTier1Bonuses = false
	req.SimOptions.Ruleset = ruleset
	req.Encounter.Duration = 45
	req.Encounter.Targets[0].Stats = stats.Stats{}.ToFloatArray()
	req.Encounter.Targets[0].Level = 60
	for len(req.Encounter.Targets) < 3 {
		req.Encounter.Targets = append(req.Encounter.Targets, googleProto.Clone(req.Encounter.Targets[0]).(*proto.Target))
	}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	for _, unit := range sim.Environment.AllUnits {
		if unit.Type == core.PlayerUnit {
			unit.CancelGCDTimer(sim)
		}
		unit.AutoAttacks.CancelAutoSwing(sim)
	}
	s := sim.Raid.Parties[0].Players[0].(interface{ GetShaman() *shaman.Shaman }).GetShaman()
	return sim, s
}

func TestShamanOctober8ElementalDevastationExcludesOverload(t *testing.T) {
	sim, s := shamanOctober8Fixture(map[string]int{"lightningOverload": 3, "elementalDevastation": 3}, proto.Ruleset_RulesetForever)
	proc := s.GetAura("Elemental Devastation Proc")
	target := s.CurrentTarget
	for _, overload := range []*core.Spell{s.LightningBoltOverload[10], s.ChainLightningOverload[4]} {
		if !overload.Flags.Matches(core.SpellFlagPassiveSpell) || overload.ProcMask != core.ProcMaskSpellDamage {
			t.Fatal("fixture lost the actual triggered damage spell")
		}
		overload.BonusCritRating = 10000
		overload.CalcAndDealDamage(sim, target, 10, overload.OutcomeMagicCrit)
		if proc.IsActive() {
			t.Fatalf("triggered Overload %v incorrectly granted Elemental Devastation", overload.ActionID)
		}
	}
}

func TestShamanOctober8ChainOverloadOwnDamageRange(t *testing.T) {
	sim, s := shamanOctober8Fixture(map[string]int{"lightningOverload": 3}, proto.Ruleset_RulesetForever)
	child := s.ChainLightningOverload[4]
	child.BonusCritRating = -10000
	target := s.CurrentTarget
	low, high := 63.23678181311, 70.36321856833 // 408484, 70245 at level60; unrounded.
	if child.GetBonusDamage(target) != 0 {
		t.Fatal("unexpected spell power in the range fixture")
	}
	for i := 0; i < 64; i++ {
		before := child.SpellMetrics[target.UnitIndex].TotalDamage
		child.ApplyEffects(sim, target, child)
		damage := child.SpellMetrics[target.UnitIndex].TotalDamage - before
		if damage != 0 && (damage < low-1e-7 || damage > high+1e-7) {
			t.Fatalf("CL Overload dealt %g, outside own client range [%g,%g]", damage, low, high)
		}
	}
}

func TestShamanOctober8ElementalDevastationActiveAndClassic(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		sim, s := shamanOctober8Fixture(map[string]int{"lightningOverload": 3, "elementalDevastation": 3}, ruleset)
		proc := s.GetAura("Elemental Devastation Proc")
		target := s.CurrentTarget
		baseCrit := s.GetStat(stats.MeleeCrit)
		active := s.LightningBolt[10]
		active.CalcAndDealDamage(sim, target, 10, active.OutcomeAlwaysHit)
		if proc.IsActive() {
			t.Fatal("noncrit activated Elemental Devastation")
		}
		active.BonusCritRating = 10000
		active.CalcAndDealDamage(sim, target, 10, active.OutcomeMagicCrit)
		if !proc.IsActive() || math.Abs(s.GetStat(stats.MeleeCrit)-baseCrit-9*core.CritRatingPerCritChance) > 1e-7 {
			t.Fatal("active spell crit failed to grant the existing 9% melee crit aura")
		}
		sim.CurrentTime = 2 * time.Second
		overload := s.LightningBoltOverload[10]
		overload.BonusCritRating = 10000
		overload.CalcAndDealDamage(sim, target, 10, overload.OutcomeMagicCrit)
		wantRemaining := 8 * time.Second
		if ruleset == proto.Ruleset_RulesetClassic {
			wantRemaining = 10 * time.Second
		}
		if proc.RemainingDuration(sim) != wantRemaining {
			t.Fatalf("triggered crit refresh for ruleset %v: got %s, want %s", ruleset, proc.RemainingDuration(sim), wantRemaining)
		}
		shock := s.EarthShock[7]
		shock.BonusCritRating = 10000
		shock.CalcAndDealDamage(sim, target, 10, shock.OutcomeMagicCrit)
		if proc.RemainingDuration(sim) != 10*time.Second {
			t.Fatal("another active offensive spell crit did not refresh Elemental Devastation")
		}
		proc.Deactivate(sim)
		if math.Abs(s.GetStat(stats.MeleeCrit)-baseCrit) > 1e-7 {
			t.Fatal("Elemental Devastation did not restore melee crit")
		}
	}
}

func TestShamanOctober8AllOverloadClientRowsAndClassic(t *testing.T) {
	// Independent expected ranges from the cached 70245 SpellEffect/SpellLevels
	// capture, not the engine's arrays or an assumed half-parent relationship.
	rows := []struct {
		sourceID               int32
		learnLevel             int
		low, high, coefficient float64
	}{
		{408439, 1, 7.5000005140949995, 8.499999515705, 0.21449999511},
		{408440, 8, 14.71428566801, 16.78571433199, 0.28549998999},
		{408441, 14, 22.163268756162502, 25.8367313630375, 0.35699999332},
		{408442, 20, 27.61236611014, 31.38763400906, 0.35699999332},
		{408443, 26, 35.3320960998, 40.167903840600005, 0.35699999332},
		{408472, 32, 54.61538463844, 61.38461542116, 0.35699999332},
		{408473, 38, 71.3750000298, 80.6250000298, 0.35699999332},
		{408474, 44, 78.879194647135, 88.120805352865, 0.35699999332},
		{408475, 50, 86.179347768295, 96.820652231705, 0.35699999332},
		{408477, 56, 94.85701363551, 105.94298655521, 0.35699999332},
		{408479, 32, 44.73039200904, 50.76960799096, 0.28549998999},
		{408481, 40, 51.43877536055, 57.56122440105, 0.28549998999},
		{408482, 48, 58.02999970308, 64.46999982012, 0.28549998999},
		{408484, 56, 63.23678181311, 70.36321856833, 0.28549998999},
	}
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		sim, s := shamanOctober8Fixture(map[string]int{"lightningOverload": 3}, ruleset)
		target := s.CurrentTarget
		for i, row := range rows {
			rank := i + 1
			var parent, child *core.Spell
			low, high, coefficient := row.low, row.high, row.coefficient
			if i >= shaman.LightningBoltRanks {
				rank = i - shaman.LightningBoltRanks + 1
				parent, child = s.ChainLightning[rank], s.ChainLightningOverload[rank]
			} else {
				parent, child = s.LightningBolt[rank], s.LightningBoltOverload[rank]
			}
			if parent == nil || child == nil || child.Rank != rank || child.RequiredLevel != row.learnLevel || parent.RequiredLevel != row.learnLevel {
				t.Fatalf("source row%d lost its parent rank/learn-level mapping", row.sourceID)
			}
			if child.ActionID.SpellID != parent.ActionID.SpellID || child.ActionID.Tag != shaman.CastTagLightningOverload || child.Cost != nil || child.DefaultCast.GCD != 0 || child.DefaultCast.CastTime != 0 || child.CD.Timer != nil || child.ThreatMultiplier != 0 ||
				child.Flags.Matches(core.SpellFlagAPL) || !child.Flags.Matches(core.SpellFlagPassiveSpell) || !child.Flags.Matches(core.SpellFlagNoOnCastComplete) || child.ProcMask != core.ProcMaskSpellDamage {
				t.Fatalf("source row%d changed free/instant/triggered/no-threat or metric identity", row.sourceID)
			}
			if ruleset == proto.Ruleset_RulesetClassic {
				coefficient = parent.BonusCoefficient * .5
				if i < shaman.LightningBoltRanks {
					low, high = shaman.LightningBoltBaseDamage[rank][0]*.5, shaman.LightningBoltBaseDamage[rank][1]*.5
				} else {
					low, high = shaman.ChainLightningBaseDamage[rank][0]*.5, shaman.ChainLightningBaseDamage[rank][1]*.5
				}
			}
			if math.Abs(child.BonusCoefficient*child.DamageMultiplier-coefficient) > 1e-10 {
				t.Fatalf("row%d effective SP coefficient %g, want %g", row.sourceID, child.BonusCoefficient*child.DamageMultiplier, coefficient)
			}
			child.BonusCritRating = -10000
			child.MissileSpeed = 0 // Synchronous numeric check; real travel is checked separately.
			landed := 0
			for sample := 0; sample < 128; sample++ {
				before := child.SpellMetrics[target.UnitIndex].TotalDamage
				child.ApplyEffects(sim, target, child)
				// WaitTravelTime queues even a zero-duration projectile.
				advanceHawks(t, sim, sim.CurrentTime)
				damage := child.SpellMetrics[target.UnitIndex].TotalDamage - before
				if damage == 0 {
					continue
				}
				landed++
				if damage < low-1e-7 || damage > high+1e-7 {
					t.Fatalf("ruleset%v row%d damage%g outside[%g,%g]", ruleset, row.sourceID, damage, low, high)
				}
			}
			if landed == 0 {
				t.Fatalf("ruleset%v row%d: no landed overload range samples", ruleset, row.sourceID)
			}
			s.AddStatDynamic(sim, stats.SpellPower, 100)
			result := child.CalcDamage(sim, target, 0, child.OutcomeAlwaysHit)
			if math.Abs(result.Damage-coefficient*100) > 1e-6 {
				t.Fatalf("row%d SP scaled twice or used parent coefficient: %g, want %g", row.sourceID, result.Damage, coefficient*100)
			}
			s.AddStatDynamic(sim, stats.SpellPower, -100)
		}
	}
}

func TestShamanOctober8OverloadParentModifiersAndMaelstrom(t *testing.T) {
	sim, s := shamanOctober8Fixture(map[string]int{"lightningOverload": 3, "concussion": 5, "callOfThunder": 1, "elementalFury": 5, "maelstromWeapon": 5}, proto.Ruleset_RulesetForever)
	for _, pair := range [][2]*core.Spell{{s.LightningBolt[10], s.LightningBoltOverload[10]}, {s.ChainLightning[4], s.ChainLightningOverload[4]}} {
		parent, child := pair[0], pair[1]
		if parent.DamageMultiplierAdditive != child.DamageMultiplierAdditive || child.DamageMultiplierAdditive != 1.05 || parent.BonusCritRating != child.BonusCritRating || parent.CritDamageBonus != child.CritDamageBonus ||
			parent.MissileSpeed != child.MissileSpeed || parent.DefaultCast.GCD != core.GCDDefault {
			t.Fatal("child own-row substitution lost existing Concussion/Call of Thunder/Fury/travel or parent GCD")
		}
	}
	if s.LightningBolt[10].DefaultCast.CastTime != 2500*time.Millisecond || s.ChainLightning[4].DefaultCast.CastTime != 2*time.Second || s.ChainLightning[4].CD.Duration != 6*time.Second {
		t.Fatal("own overload rows changed active parent casting/cooldown")
	}
	s.MaelstromWeaponAura.Activate(sim)
	s.MaelstromWeaponAura.SetStacks(sim, 5)
	if !s.LightningBoltOverload[10].Cast(sim, s.CurrentTarget) || s.MaelstromWeaponAura.GetStacks() != 5 {
		t.Fatal("triggered overload consumed Maelstrom Weapon")
	}
	if !s.LightningBolt[10].Cast(sim, s.CurrentTarget) || s.MaelstromWeaponAura.IsActive() {
		t.Fatal("active Lightning Bolt stopped consuming Maelstrom Weapon")
	}
	if s.AutoAttacks.IsDualWielding {
		t.Fatal("overload substitution enabled Shaman dual wield")
	}
}

func TestShamanOctober8OverloadBoltTravelAndOrdinaryFlurry(t *testing.T) {
	sim, s := shamanOctober8Fixture(map[string]int{"lightningOverload": 3, "flurry": 5}, proto.Ruleset_RulesetForever)
	child := s.LightningBoltOverload[10]
	s.DistanceFromTarget = 30
	if child.MissileSpeed != 20 || child.TravelTime() != 1500*time.Millisecond {
		t.Fatal("overload lost existing bolt projectile timing")
	}
	deliveries := 0
	trigger := s.GetAura("Flurry Proc Trigger")
	original := trigger.OnSpellHitDealt
	trigger.OnSpellHitDealt = func(a *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		if spell == child {
			deliveries++
		}
		original(a, sim, spell, result)
	}
	if !child.Cast(sim, s.CurrentTarget) {
		t.Fatal("free instant overload did not cast")
	}
	if deliveries != 0 {
		t.Fatal("projectile delivered its hit/miss event before arrival")
	}
	advanceHawks(t, sim, 1500*time.Millisecond)
	if deliveries != 1 {
		t.Fatalf("projectile delivery events %d, want1 on arrival", deliveries)
	}
	flurry := s.GetAura("Flurry Proc (16280)")
	auto := s.AutoAttacks.MHAuto()
	if auto.Flags.Matches(core.SpellFlagPassiveSpell) {
		t.Fatal("ordinary/extra main-hand autos were reclassified as proc spells")
	}
	auto.BonusCritRating = 10000
	auto.CalcAndDealDamage(sim, s.CurrentTarget, 1, auto.OutcomeMeleeSpecialCritOnly)
	if flurry.GetStacks() != 3 {
		t.Fatal("ordinary melee crit stopped granting Flurry")
	}
	flurry.SetStacks(sim, 1)
	// An ordinary extra-auto packet uses the same attack spell, not a new proc
	// spell. WF's separate yellow surrogate remains explicitly unvalidated.
	auto.CalcAndDealDamage(sim, s.CurrentTarget, 1, auto.OutcomeMeleeSpecialCritOnly)
	if flurry.GetStacks() != 3 {
		t.Fatal("another ordinary auto crit stopped refreshing Flurry")
	}
}
