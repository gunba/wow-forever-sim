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

func gnomeFixture(key string) *core.Simulation {
	req := racialFixture(key, proto.Race_RaceGnome)
	p := req.Raid.Parties[0].Players[0]
	p.TalentsString = ""
	p.Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	s := core.NewSim(req, simsignals.Signals{})
	s.Reset()
	return s
}

func TestForeverEurekaClassVariants(t *testing.T) {
	for _, tc := range []struct {
		key           string
		racial, spell int32
		discount      int32
	}{
		{"combat", 1259812, 11294, 10},
		{"fury", 1259813, core.TernaryInt32(core.IncludeAQ, 25286, 11567), 10},
		{"arcane", 1259817, 133, 10},
		{"ds_ruin", 1259821, 686, 10},
		{"smite", 1259823, 585, 10},
	} {
		t.Run(tc.key, func(t *testing.T) {
			s := gnomeFixture(tc.key)
			unit := s.Raid.AllPlayerUnits[0]
			racial := unit.GetSpell(core.ActionID{SpellID: tc.racial})
			action := unit.GetSpell(core.ActionID{SpellID: tc.spell})
			aura := unit.GetAura("Eureka!")
			if racial == nil || action == nil || action.Cost == nil {
				t.Fatal("missing racial or damaging ability")
			}
			if aura.Duration != 15*time.Second || racial.CD.Duration != 2*time.Minute || racial.Cost != nil {
				t.Fatal("incorrect Eureka timing or cost")
			}
			cost := action.Cost.Multiplier
			multiplier := action.DamageMultiplier
			globalDamage := unit.PseudoStats.DamageDealtMultiplier
			racial.ApplyEffects(s, unit, racial)
			if aura.GetStacks() != 3 || action.Cost.Multiplier != cost-tc.discount {
				t.Fatal("incorrect class cost reduction or charge count")
			}
			if unit.PseudoStats.DamageDealtMultiplier != globalDamage {
				t.Fatal("Eureka must not amplify white attacks or incidental damage")
			}
			for i := int32(0); i < 3; i++ {
				action.ApplyEffects(s, s.Encounter.TargetUnits[0], action)
				if aura.GetStacks() != 2-i {
					t.Fatalf("cast %d left %d charges", i+1, aura.GetStacks())
				}
			}
			if aura.IsActive() || action.Cost.Multiplier != cost || math.Abs(action.DamageMultiplier-multiplier) > 1e-10 {
				t.Fatal("third cast failed to remove Eureka or restore modifiers")
			}
		})
	}
}

func TestEurekaThirdChannelKeepsItsBonus(t *testing.T) {
	s := gnomeFixture("arcane")
	unit := s.Raid.AllPlayerUnits[0]
	parent := unit.GetSpell(core.ActionID{SpellID: 5143})
	tick := unit.GetSpell(core.ActionID{SpellID: 5143, Tag: 1})
	base := tick.DamageMultiplier
	var observed float64
	tick.ApplyEffects = func(_ *core.Simulation, _ *core.Unit, spell *core.Spell) {
		observed = spell.DamageMultiplier
	}
	aura := unit.GetAura("Eureka!")
	aura.Activate(s)
	aura.SetStacks(s, 1)
	target := s.Encounter.TargetUnits[0]
	parent.ApplyEffects(s, target, parent)
	if aura.IsActive() {
		t.Fatal("channel start did not consume the final charge")
	}
	dot := parent.Dot(target)
	dot.OnTick(s, target, dot)
	if math.Abs(observed/base-1.1) > 1e-10 || math.Abs(tick.DamageMultiplier-base) > 1e-10 {
		t.Fatal("channel lost its charged damage or leaked it to another cast")
	}
	parent.ApplyEffects(s, target, parent)
	dot.OnTick(s, target, dot)
	if math.Abs(observed-base) > 1e-10 {
		t.Fatal("uncharged channel retained the previous cast's bonus")
	}
}

func TestGnomeEnergyAndRagePools(t *testing.T) {
	rogue := gnomeFixture("combat").Raid.AllPlayerUnits[0]
	s := gnomeFixture("fury")
	warrior := s.Raid.AllPlayerUnits[0]
	warrior.AddRage(s, 1000, warrior.NewRageMetrics(core.ActionID{SpellID: 1259802}))
	if rogue.MaxEnergy() != 105 || warrior.CurrentRage() != 105 {
		t.Fatalf("Expansive Mind pools: energy %.2f, rage %.2f", rogue.MaxEnergy(), warrior.CurrentRage())
	}
}

func TestElunesLightClientSpell(t *testing.T) {
	req := racialFixture("balance", proto.Race_RaceNightElf)
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	spell := env.Raid.AllPlayerUnits[0].GetSpell(core.ActionID{SpellID: 1259799})
	if spell == nil || spell.CD.Duration != 3*time.Minute {
		t.Fatal("Elune's Light does not match its client spell")
	}
}

func TestAllianceReceivesForeverShamanBuffs(t *testing.T) {
	for _, race := range []proto.Race{proto.Race_RaceHuman, proto.Race_RaceOrc} {
		req := racialFixture("fury", race)
		_, before, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		req.Raid.Parties[0].Buffs = &proto.PartyBuffs{
			StrengthOfEarthTotem: proto.TristateEffect_TristateEffectRegular,
			GraceOfAirTotem:      proto.TristateEffect_TristateEffectRegular,
			ManaSpringTotem:      proto.TristateEffect_TristateEffectRegular,
		}
		_, after, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		b := stats.FromFloatArray(before.Parties[0].Players[0].FinalStats.Stats)
		a := stats.FromFloatArray(after.Parties[0].Players[0].FinalStats.Stats)
		for stat, want := range map[stats.Stat]float64{stats.Strength: 53, stats.Agility: 89, stats.MP5: 25} {
			if math.Abs(a[stat]-b[stat]-want) > 1e-8 {
				t.Fatalf("%s: stat %v changed by %.2f, want %.2f", race, stat, a[stat]-b[stat], want)
			}
		}
	}
}
