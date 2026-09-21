//go:build with_db

package main

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
)

func tierRequest(key string, enabled bool) *proto.RaidSimRequest {
	for _, b := range builds() {
		if b.Key == key {
			req := racialFixture(key, b.races()[0])
			req.Raid.Parties[0].Players[0].ForeverTier1Bonuses = enabled
			return req
		}
	}
	panic(key)
}

func tierSim(key string, enabled bool) *core.Simulation {
	s := core.NewSim(tierRequest(key, enabled), simsignals.Signals{})
	s.Reset()
	return s
}

func TestForeverTier1EveryBuildAndNoDoubleCount(t *testing.T) {
	for _, b := range builds() {
		t.Run(b.Key, func(t *testing.T) {
			for _, enabled := range []bool{false, true} {
				req := tierRequest(b.Key, enabled)
				env, result, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
				character := env.Raid.Parties[0].Players[0].GetCharacter()
				got := character.GetActiveSetBonuses()
				want := 0
				if enabled {
					want = 4
				}
				if len(got) != want || len(result.Parties[0].Players[0].Sets) != want {
					t.Fatalf("enabled=%v: bonuses=%v, want %d", enabled, got, want)
				}
				if !enabled {
					continue
				}
				for i, bonus := range got {
					if bonus.NumPieces != int32(i+2) {
						t.Fatalf("wrong bonus threshold: %v", bonus)
					}
				}
				// Simulate real tier pieces coexisting with the override. The same
				// lookup drives effect application, so each threshold occurs once.
				for i := 0; i < 6; i++ {
					character.Equipment[i].SetID = core.ForeverTier1SetID(character.Spec)
					character.Equipment[i].SetName = got[0].Name
				}
				if len(character.GetActiveSetBonuses()) != 4 {
					t.Fatal("equipped pieces doubled the virtual set")
				}
				character.ForeverTier1Bonuses = false
				if len(character.GetActiveSetBonuses()) != 4 {
					t.Fatal("ordinary equipped set stopped working")
				}
			}
		})
	}
	req := tierRequest("arcane", true)
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetClassic, false)
	if len(env.Raid.Parties[0].Players[0].GetCharacter().GetActiveSetBonuses()) != 0 {
		t.Fatal("Forever override leaked into Classic")
	}
}

func TestForeverTier1StatsAndCooldowns(t *testing.T) {
	for _, key := range []string{"arcane", "balance", "shadow", "elemental", "combat", "affliction"} {
		off, on := tierSim(key, false), tierSim(key, true)
		a, b := off.Raid.AllPlayerUnits[0], on.Raid.AllPlayerUnits[0]
		for _, stat := range []stats.Stat{stats.MeleeHit, stats.SpellHit} {
			if math.Abs(b.GetStat(stat)-a.GetStat(stat)-1) > 1e-8 {
				t.Errorf("%s: missing shared hit bonus for %v", key, stat)
			}
		}
	}
	for _, key := range []string{"fury", "feral", "enhancement", "retribution", "marksmanship"} {
		a, b := tierSim(key, false).Raid.AllPlayerUnits[0], tierSim(key, true).Raid.AllPlayerUnits[0]
		for _, pair := range [][2]float64{
			{a.PseudoStats.MeleeSpeedMultiplier, b.PseudoStats.MeleeSpeedMultiplier},
			{a.PseudoStats.RangedSpeedMultiplier, b.PseudoStats.RangedSpeedMultiplier},
			{a.PseudoStats.CastSpeedMultiplier, b.PseudoStats.CastSpeedMultiplier},
		} {
			if math.Abs(pair[1]/pair[0]-1.01) > 1e-8 {
				t.Errorf("%s: incorrect haste bonus: %v", key, pair)
			}
		}
	}
	for _, tc := range []struct {
		key       string
		id        int32
		reduction time.Duration
	}{
		{"arcane", 2139, 5 * time.Second}, {"marksmanship", 20904, time.Second},
		{"marksmanship", 2643, time.Second}, {"feral", 9846, 3 * time.Second},
		{"fury", 1719, 30 * time.Second}, {"retribution", 20271, 500 * time.Millisecond},
		{"enhancement", 17364, 500 * time.Millisecond}, {"elemental", 1238300, time.Second},
		{"smite", 1316995, time.Second}, {"shadow", 19280, time.Minute},
	} {
		a := tierSim(tc.key, false).Raid.AllPlayerUnits[0].GetSpell(core.ActionID{SpellID: tc.id})
		b := tierSim(tc.key, true).Raid.AllPlayerUnits[0].GetSpell(core.ActionID{SpellID: tc.id})
		if a == nil || b == nil || a.CD.Duration-b.CD.Duration != tc.reduction {
			t.Errorf("%s/%d: incorrect cooldown reduction", tc.key, tc.id)
		}
	}
}

func TestForeverTier1CreatureRestrictions(t *testing.T) {
	for _, tc := range []struct {
		key    string
		mob    proto.MobType
		ap, sp float64
	}{
		{"fury", proto.MobType_MobTypeHumanoid, 36, 0},
		{"marksmanship", proto.MobType_MobTypeBeast, 36, 0},
		{"enhancement", proto.MobType_MobTypeElemental, 36, 0},
		{"retribution", proto.MobType_MobTypeDemon, 36, 0},
		{"arcane", proto.MobType_MobTypeElemental, 0, 21},
		{"affliction", proto.MobType_MobTypeDemon, 0, 21},
		{"shadow", proto.MobType_MobTypeUndead, 0, 21},
	} {
		req := tierRequest(tc.key, true)
		req.Encounter.Targets = append(req.Encounter.Targets, &proto.Target{Level: 63, MobType: tc.mob})
		s := core.NewSim(req, simsignals.Signals{})
		unit := s.Raid.AllPlayerUnits[0]
		for _, aura := range unit.GetAuras() {
			if strings.HasPrefix(aura.Label, "Tier 1 creature bonus ") {
				aura.Activate(s)
				for i, target := range s.Encounter.TargetUnits {
					wantAP, wantSP := 0.0, 0.0
					if i == 1 {
						wantAP, wantSP = tc.ap, tc.sp
					}
					for _, table := range unit.AttackTables[target.UnitIndex] {
						if table.BonusAttackPowerTaken != wantAP || table.BonusSpellDamageTaken != wantSP {
							t.Fatalf("%s target %d: incorrect conditional power", tc.key, i)
						}
					}
				}
				aura.Deactivate(s)
				aura.Activate(s)
				for _, table := range unit.AttackTables[s.Encounter.TargetUnits[1].UnitIndex] {
					if table.BonusAttackPowerTaken != tc.ap || table.BonusSpellDamageTaken != tc.sp {
						t.Fatal("conditional power accumulated")
					}
				}
			}
		}
	}
}

func TestForeverTier1ResourceAndDotBonuses(t *testing.T) {
	for _, key := range []string{"combat", "mutilate", "subtlety"} {
		a, b := tierSim(key, false).Raid.AllPlayerUnits[0], tierSim(key, true).Raid.AllPlayerUnits[0]
		for _, spell := range a.Spellbook {
			if spell.SpellID == 11300 || spell.SpellID == 31016 {
				if spell.Cost.GetCurrentCost()-b.GetSpell(spell.ActionID).Cost.GetCurrentCost() != 5 {
					t.Fatalf("%s: Eviscerate cost did not drop by five", key)
				}
			}
		}
	}
	var restored []float64
	for _, enabled := range []bool{false, true} {
		s := tierSim("affliction", enabled)
		u := s.Raid.AllPlayerUnits[0]
		u.SpendMana(s, u.CurrentMana(), u.NewManaMetrics(core.ActionID{SpellID: 1}))
		spell := u.GetSpell(core.ActionID{SpellID: 11689})
		spell.ApplyEffects(s, u, spell)
		restored = append(restored, u.CurrentMana())

		balance := tierSim("balance", enabled)
		dot := balance.Raid.AllPlayerUnits[0].GetSpell(core.ActionID{SpellID: 24977}).Dot(balance.Encounter.TargetUnits[0])
		// This baseline also has Nature's Splendor: one extra two-second tick.
		duration, ticks := 14*time.Second, int32(7)
		if enabled {
			duration, ticks = 17*time.Second, 8
		}
		dot.RecomputeAuraDuration()
		if dot.Duration != duration || dot.NumberOfTicks != ticks || dot.TickPeriod() != 2*time.Second || dot.BonusCoefficient != .158 {
			t.Fatalf("Insect Swarm duration/ticks changed incorrectly: %+v", dot)
		}
	}
	if restored[0] <= 0 || math.Abs(restored[1]/restored[0]-1.2) > 1e-8 {
		t.Fatalf("Life Tap mana multiplier: %v", restored)
	}
}

func TestForeverTier1MageCombustion(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		s := tierSim("fire", enabled)
		u := s.Raid.AllPlayerUnits[0]
		ffb := u.GetSpell(core.ActionID{SpellID: 1237313})
		fireball := u.GetSpell(core.ActionID{SpellID: 10151})
		ffbBefore, fireballBefore := ffb.BonusCritRating, fireball.BonusCritRating
		aura := u.GetAura("Combustion")
		aura.Activate(s)
		want := 0.0
		if enabled {
			want = 10 * core.SpellCritRatingPerCritChance
		}
		if ffb.BonusCritRating-ffbBefore != want || fireball.BonusCritRating != fireballBefore {
			t.Fatal("Combustion's Tier 1 crit bonus must affect only Frostfire Bolt")
		}
		aura.Deactivate(s)
		if ffb.BonusCritRating != ffbBefore || fireball.BonusCritRating != fireballBefore {
			t.Fatal("Combustion tier crit did not expire")
		}
	}
}

func TestForeverTier1MageProcChances(t *testing.T) {
	const trials = 10000
	for _, tc := range []struct {
		key, trigger, proc string
		spell              int32
		base               float64
		hit                bool
	}{
		{"arcane", "Missile Barrage Trigger", "Missile Barrage", 1237313, .20, false},
		{"arcane", "Missile Barrage Trigger", "Missile Barrage", 133, .20, false},
		{"frost", "Fingers of Frost Trigger", "Fingers of Frost", 1237313, .15, true},
		{"frost", "Fingers of Frost Trigger", "Fingers of Frost", 116, .15, true},
	} {
		for _, enabled := range []bool{false, true} {
			s := tierSim(tc.key, enabled)
			u := s.Raid.AllPlayerUnits[0]
			trigger, proc := u.GetAura(tc.trigger), u.GetAura(tc.proc)
			spell := u.GetSpell(core.ActionID{SpellID: tc.spell})
			if trigger == nil || proc == nil || spell == nil {
				t.Fatalf("missing %s fixture", tc.key)
			}
			hits := 0
			for i := 0; i < trials; i++ {
				if tc.hit {
					trigger.OnSpellHitDealt(trigger, s, spell, &core.SpellResult{Target: s.Encounter.TargetUnits[0], Outcome: core.OutcomeHit})
				} else {
					trigger.OnCastComplete(trigger, s, spell)
				}
				if proc.IsActive() {
					hits++
					proc.Deactivate(s)
				}
			}
			want := tc.base
			if enabled && tc.spell == 1237313 {
				want += .10
			}
			if got := float64(hits) / trials; math.Abs(got-want) > .02 {
				t.Fatalf("%s/%d tier=%v: proc rate %.3f, want %.3f", tc.key, tc.spell, enabled, got, want)
			}
		}
	}
}

func TestForeverTier1HitReducesPaidBudget(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "arcane" {
			continue
		}
		p := b.player(b.races()[0])
		p.ForeverTier1Bonuses = false
		_, without, err := capHit(b, p)
		if err != nil {
			t.Fatal(err)
		}
		p.ForeverTier1Bonuses = true
		_, with, err := capHit(b, p)
		if err != nil {
			t.Fatal(err)
		}
		if math.Abs(without.RawHitDelta-with.RawHitDelta-10) > 1e-8 {
			t.Fatal("Tier 1 hit did not save ten raw hit-rating budget points")
		}
	}
}
