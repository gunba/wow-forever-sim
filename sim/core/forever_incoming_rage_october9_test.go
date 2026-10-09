package core

import (
	"math"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"google.golang.org/protobuf/encoding/protojson"
	gproto "google.golang.org/protobuf/proto"
)

// This exercises the actual registered incoming-hit callback, not only a formula.
func october9IncomingRageWithModel(t *testing.T, ruleset proto.Ruleset, warrior bool, level int32, health float64, model *proto.ForeverIncomingRageModel, result *SpellResult) float64 {
	t.Helper()
	parameters, err := ResolveForeverIncomingRageModel(level, model)
	if err != nil {
		t.Fatal(err)
	}
	unit := &Unit{
		Level:                         level,
		Env:                           &Environment{Ruleset: ruleset},
		auraTracker:                   newAuraTracker(),
		Metrics:                       NewUnitMetrics(),
		PseudoStats:                   stats.NewPseudoStats(),
		foreverIncomingRageParameters: parameters,
	}
	unit.stats[stats.Health] = health
	unit.EnableHealthBar()
	unit.EnableRageBar(RageBarOptions{DamageDealtMultiplier: 1, DamageTakenMultiplier: 1, ForeverWarriorRage: warrior, ForeverBearCriticalRage: !warrior})
	sim := &Simulation{Options: &proto.SimOptions{Interactive: true}, pendingActions: []*PendingAction{{NextActionAt: NeverExpires}}}
	aura := unit.GetAura("RageBar")
	aura.OnSpellHitTaken(aura, sim, &Spell{Unit: &Unit{Level: 63}}, result)
	return unit.CurrentRage()
}

func october9IncomingRage(t *testing.T, ruleset proto.Ruleset, warrior bool, health, damage, absorbed, mitigation float64) float64 {
	return october9IncomingRageWithModel(t, ruleset, warrior, 60, health, nil, &SpellResult{Outcome: OutcomeHit, Damage: damage, AbsorbedDamage: absorbed, ResistanceMultiplier: mitigation})
}

func TestForeverIncomingRageOctober9CentralContract(t *testing.T) {
	// Explicit provisional central scenario, NOT a decoded server coefficient.
	want := 100.0 * 10 / 4365 / (1 - 3566.0/(3566+5500))
	for _, warrior := range []bool{true, false} {
		for _, health := range []float64{3000, 9000} {
			for _, absorbed := range []float64{0, 40, 100} {
				got := october9IncomingRage(t, proto.Ruleset_RulesetForever, warrior, health, 100-absorbed, absorbed, .5)
				if math.Abs(got-want) > 1e-12 {
					t.Errorf("warrior=%t HP=%g absorbed=%g: incoming Rage %g, want provisional %g", warrior, health, absorbed, got, want)
				}
			}
		}
	}
}

func TestForeverIncomingRageOctober9RetainsActualArmor(t *testing.T) {
	for _, warrior := range []bool{true, false} {
		naked := october9IncomingRage(t, proto.Ruleset_RulesetForever, warrior, 3000, 100, 0, 1)
		armored := october9IncomingRage(t, proto.Ruleset_RulesetForever, warrior, 3000, 50, 0, .5)
		if math.Abs(armored-naked/2) > 1e-12 {
			t.Errorf("warrior=%t: armor-undone Rage %g; want half of naked %g", warrior, armored, naked)
		}
	}
}

func TestForeverIncomingRageOctober9ClassicRetained(t *testing.T) {
	for _, warrior := range []bool{true, false} {
		got := october9IncomingRage(t, proto.Ruleset_RulesetClassic, warrior, 3000, 100, 50, .5)
		want := 100 * 2.5 / GetRageConversion(63)
		if math.Abs(got-want) > 1e-12 {
			t.Errorf("warrior=%t: Classic Rage %g, want unchanged %g", warrior, got, want)
		}
	}
}

func TestForeverIncomingRageOctober9OutcomesAndOverrides(t *testing.T) {
	for _, warrior := range []bool{true, false} {
		model := &proto.ForeverIncomingRageModel{Coefficient: gproto.Float64(10), ReferenceArmorOverride: gproto.Float64(0), ExpectedHealthOverride: gproto.Float64(1000)}
		for _, tc := range []struct {
			outcome                HitOutcome
			damage, absorbed, want float64
		}{
			{OutcomeHit, 100, 0, 1},
			{OutcomeCrit, 200, 0, 2},
			{OutcomeCrush, 150, 0, 1.5},
			{OutcomeCrit, 120, 80, 2},
			{OutcomeCrush, 0, 150, 1.5},
		} {
			got := october9IncomingRageWithModel(t, proto.Ruleset_RulesetForever, warrior, 60, 5000, model, &SpellResult{Outcome: tc.outcome, Damage: tc.damage, AbsorbedDamage: tc.absorbed, PreOutcomeDamage: 100, ResistanceMultiplier: .5})
			if math.Abs(got-tc.want) > 1e-12 {
				t.Errorf("warrior=%t outcome=%v: got %g, want %g", warrior, tc.outcome, got, tc.want)
			}
		}
		model.Coefficient = gproto.Float64(0)
		got := october9IncomingRageWithModel(t, proto.Ruleset_RulesetForever, warrior, 60, 5000, model, &SpellResult{Outcome: OutcomeCrit, Damage: 200, AbsorbedDamage: 100})
		if got != 0 {
			t.Errorf("explicit zero coefficient gave %g Rage", got)
		}
		model.Coefficient = gproto.Float64(20)
		model.ReferenceArmorOverride = gproto.Float64(.2)
		got = october9IncomingRageWithModel(t, proto.Ruleset_RulesetForever, warrior, 60, 5000, model, &SpellResult{Damage: 100, ResistanceMultiplier: .5})
		if math.Abs(got-2.5) > 1e-12 {
			t.Errorf("numeric overrides gave %g, want 2.5", got)
		}
		got = october9IncomingRageWithModel(t, proto.Ruleset_RulesetClassic, warrior, 60, 5000, model, &SpellResult{Damage: 100, AbsorbedDamage: 100})
		if want := 100 * 2.5 / GetRageConversion(63); math.Abs(got-want) > 1e-12 {
			t.Errorf("overrides changed Classic: %g, want %g", got, want)
		}
	}
}

func TestForeverIncomingRageOctober9LevelsAndValidation(t *testing.T) {
	for _, tc := range []struct {
		level             int32
		health, reference float64
	}{
		{1, 42, .2}, {8, 156, 322.0 / (322 + 1080)}, {30, 1242, 1212.0 / (1212 + 2950)}, {60, 4365, 3566.0 / (3566 + 5500)},
	} {
		parameters, err := ResolveForeverIncomingRageModel(tc.level, nil)
		if err != nil {
			t.Fatal(err)
		}
		if parameters.ExpectedHealth != tc.health || math.Abs(parameters.ReferenceArmor-tc.reference) > 1e-12 {
			t.Errorf("level %d: %+v", tc.level, parameters)
		}
		got := october9IncomingRageWithModel(t, proto.Ruleset_RulesetForever, true, tc.level, 9999, nil, &SpellResult{Damage: 100})
		if want := 100 * 10 / tc.health / (1 - tc.reference); math.Abs(got-want) > 1e-12 {
			t.Errorf("level %d: %g, want %g", tc.level, got, want)
		}
	}
	for _, level := range []int32{-1, 0, 61} {
		if _, err := ResolveForeverIncomingRageModel(level, nil); err == nil {
			t.Errorf("accepted level %d", level)
		}
	}
	for _, tc := range []struct {
		name  string
		model *proto.ForeverIncomingRageModel
	}{
		{"negative coefficient", &proto.ForeverIncomingRageModel{Coefficient: gproto.Float64(-1)}},
		{"too large coefficient", &proto.ForeverIncomingRageModel{Coefficient: gproto.Float64(1001)}},
		{"NaN coefficient", &proto.ForeverIncomingRageModel{Coefficient: gproto.Float64(math.NaN())}},
		{"infinite coefficient", &proto.ForeverIncomingRageModel{Coefficient: gproto.Float64(math.Inf(1))}},
		{"negative Armor", &proto.ForeverIncomingRageModel{ReferenceArmorOverride: gproto.Float64(-.1)}},
		{"too large Armor", &proto.ForeverIncomingRageModel{ReferenceArmorOverride: gproto.Float64(.951)}},
		{"NaN Armor", &proto.ForeverIncomingRageModel{ReferenceArmorOverride: gproto.Float64(math.NaN())}},
		{"infinite Armor", &proto.ForeverIncomingRageModel{ReferenceArmorOverride: gproto.Float64(math.Inf(-1))}},
		{"zero health", &proto.ForeverIncomingRageModel{ExpectedHealthOverride: gproto.Float64(0)}},
		{"too large health", &proto.ForeverIncomingRageModel{ExpectedHealthOverride: gproto.Float64(1e9 + 1)}},
		{"NaN health", &proto.ForeverIncomingRageModel{ExpectedHealthOverride: gproto.Float64(math.NaN())}},
		{"infinite health", &proto.ForeverIncomingRageModel{ExpectedHealthOverride: gproto.Float64(math.Inf(1))}},
	} {
		if _, err := ResolveForeverIncomingRageModel(60, tc.model); err == nil {
			t.Errorf("accepted %s", tc.name)
		}
	}
	for _, model := range []*proto.ForeverIncomingRageModel{nil, {}, {Coefficient: gproto.Float64(0), ReferenceArmorOverride: gproto.Float64(0), ExpectedHealthOverride: gproto.Float64(1)}, {Coefficient: gproto.Float64(1000), ReferenceArmorOverride: gproto.Float64(.95), ExpectedHealthOverride: gproto.Float64(1e9)}} {
		if _, err := ResolveForeverIncomingRageModel(60, model); err != nil {
			t.Errorf("rejected valid boundary: %v", err)
		}
	}
}

func TestForeverIncomingRageOctober9NativeRoundtrip(t *testing.T) {
	for _, model := range []*proto.ForeverIncomingRageModel{nil, {}, {Coefficient: gproto.Float64(0)}, {ReferenceArmorOverride: gproto.Float64(0)}, {Coefficient: gproto.Float64(20), ReferenceArmorOverride: gproto.Float64(.3), ExpectedHealthOverride: gproto.Float64(2000)}} {
		for _, threat := range []*float64{nil, gproto.Float64(0), gproto.Float64(77)} {
			input := &proto.Player{ForeverIncomingRageModel: model, ForeverDemoralizingThreat: threat}
			wire, err := gproto.Marshal(input)
			if err != nil {
				t.Fatal(err)
			}
			fromWire := &proto.Player{}
			if err := gproto.Unmarshal(wire, fromWire); err != nil {
				t.Fatal(err)
			}
			jsonBytes, err := protojson.Marshal(fromWire)
			if err != nil {
				t.Fatal(err)
			}
			fromJSON := &proto.Player{}
			if err := protojson.Unmarshal(jsonBytes, fromJSON); err != nil {
				t.Fatal(err)
			}
			if !gproto.Equal(input, fromWire) || !gproto.Equal(input, fromJSON) {
				t.Errorf("optional presence lost: %v -> %v -> %v", input, fromWire, fromJSON)
			}
		}
	}
}

func october9Player(model *proto.ForeverIncomingRageModel, threat *float64) *proto.Player {
	return &proto.Player{Class: proto.Class_ClassWarrior, Race: proto.Race_RaceHuman, Equipment: &proto.EquipmentSpec{}, Spec: &proto.Player_Warrior{Warrior: &proto.Warrior{}}, ForeverIncomingRageModel: model, ForeverDemoralizingThreat: threat}
}

func TestForeverIncomingRageOctober9CharacterValidationAndCopy(t *testing.T) {
	model := &proto.ForeverIncomingRageModel{Coefficient: gproto.Float64(20), ReferenceArmorOverride: gproto.Float64(.3), ExpectedHealthOverride: gproto.Float64(2000)}
	threat := 77.0
	character := NewCharacter(&Party{}, 0, october9Player(model, &threat))
	*model.Coefficient = 5
	*model.ReferenceArmorOverride = .1
	*model.ExpectedHealthOverride = 1000
	threat = 88
	parameters := character.ForeverIncomingRageParameters()
	if parameters.Coefficient != 20 || parameters.ReferenceArmor != .3 || parameters.ExpectedHealth != 2000 || !parameters.CoefficientOverride || !parameters.ReferenceArmorOverride || !parameters.ExpectedHealthOverride {
		t.Errorf("resolved values changed: %+v", parameters)
	}
	if character.ForeverDemoralizingThreat == nil || *character.ForeverDemoralizingThreat != 77 {
		t.Errorf("Demoralizing threat was not copied")
	}
	for _, player := range []*proto.Player{
		october9Player(&proto.ForeverIncomingRageModel{Coefficient: gproto.Float64(math.NaN())}, nil),
		october9Player(&proto.ForeverIncomingRageModel{ReferenceArmorOverride: gproto.Float64(1)}, nil),
		october9Player(&proto.ForeverIncomingRageModel{ExpectedHealthOverride: gproto.Float64(0)}, nil),
		october9Player(nil, gproto.Float64(-1)), october9Player(nil, gproto.Float64(1e6+1)), october9Player(nil, gproto.Float64(math.Inf(1))), october9Player(nil, gproto.Float64(math.NaN())),
	} {
		t.Run("invalid", func(t *testing.T) {
			defer func() {
				rejection := recover()
				if rejection == nil {
					t.Error("invalid imported profile was not rejected")
				} else if err, ok := rejection.(error); !ok || !strings.Contains(err.Error(), "provisional") {
					t.Errorf("unexpected failure, not model validation: %v", rejection)
				}
			}()
			NewCharacter(&Party{}, 0, player)
		})
	}
	for _, threat := range []*float64{nil, gproto.Float64(0), gproto.Float64(1e6)} {
		character := NewCharacter(&Party{}, 0, october9Player(nil, threat))
		if (character.ForeverDemoralizingThreat == nil) != (threat == nil) {
			t.Error("threat presence changed")
		}
	}
}
