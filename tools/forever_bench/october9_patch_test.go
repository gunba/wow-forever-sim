//go:build with_db

package main

import (
	"math"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	googleProto "google.golang.org/protobuf/proto"
)

func TestOctober9LastStandHealthGain(t *testing.T) {
	for _, race := range []proto.Race{proto.Race_RaceHuman, proto.Race_RaceTauren} {
		for _, mining := range []bool{false, true} {
			req := historyTalentFixture("tank_warrior", map[string]int{"lastStand": 1})
			p := req.Raid.Parties[0].Players[0]
			p.Race = race
			p.ForeverTier1Bonuses = false
			if mining {
				p.Profession1 = proto.Profession_Mining
			}
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			c := sim.Raid.Parties[0].Players[0].GetCharacter()
			base := c.MaxHealth()
			c.RemoveHealth(sim, base/2)
			if !c.GetSpell(core.ActionID{SpellID: 12975}).Cast(sim, c.CurrentTarget) {
				t.Fatal("Last Stand did not cast")
			}
			if math.Abs(c.MaxHealth()-base*1.3) > 1e-7 || math.Abs(c.CurrentHealth()-base*.8) > 1e-7 {
				t.Errorf("%v mining=%v: max %v/current %v, want %v/%v", race, mining, c.MaxHealth(), c.CurrentHealth(), base*1.3, base*.8)
			}
			c.GetAura("Last Stand").Deactivate(sim)
			if math.Abs(c.MaxHealth()-base) > 1e-7 || math.Abs(c.CurrentHealth()-base*.5) > 1e-7 {
				t.Fatal("Last Stand retained temporary maximum/current health after expiration")
			}
			sim = core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			c = sim.Raid.Parties[0].Players[0].GetCharacter()
			if !c.GetSpell(core.ActionID{SpellID: 12975}).Cast(sim, c.CurrentTarget) {
				t.Fatal("Last Stand did not cast after reset")
			}
			c.RemoveHealth(sim, c.CurrentHealth()-1)
			c.GetAura("Last Stand").Deactivate(sim)
			if c.CurrentHealth() != 1 {
				t.Fatal("temporary-health expiration killed or resurrected the living tank")
			}
		}
	}
}

// The official patch proves nonzero threat, not its amount. Defaults below
// retain the existing Classic-derived rank convention as a provisional scenario;
// explicit zero remains a control, not a claim about the patched server.
func TestOctober9DemoralizingThreatModel(t *testing.T) {
	for _, tc := range []struct {
		key     string
		id      int32
		central float64
	}{
		{"tank_warrior", 11556, 43.2},
		{"feral_tank_druid", 9898, 100},
	} {
		for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
			for _, override := range []*float64{nil, googleProto.Float64(0), googleProto.Float64(75)} {
				req := historyTalentFixture(tc.key, nil)
				req.SimOptions.Ruleset = ruleset
				req.Raid.Parties[0].Players[0].ForeverDemoralizingThreat = override
				for len(req.Encounter.Targets) < 3 {
					req.Encounter.Targets = append(req.Encounter.Targets, googleProto.Clone(req.Encounter.Targets[0]).(*proto.Target))
				}
				sim := core.NewSim(req, simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				c := sim.Raid.Parties[0].Players[0].GetCharacter()
				spell := c.GetSpell(core.ActionID{SpellID: tc.id})
				if spell == nil {
					t.Fatal("Demoralizing fixture is not registered")
				}
				want := tc.central
				if ruleset == proto.Ruleset_RulesetForever && override != nil {
					want = *override
				}
				for i := 0; i < 100; i++ {
					spell.ApplyEffects(sim, c.CurrentTarget, spell)
				}
				models := provisionalModelsForCharacter(c, req.Raid.Parties[0].Players[0])
				if ruleset == proto.Ruleset_RulesetForever {
					parameters, ok := models["incomingRage"].(core.ForeverIncomingRageParameters)
					if !ok || parameters.ExpectedHealth != 4365 || parameters.Coefficient != 10 || models["demoralizingThreatPerTarget"] != want || models["demoralizingThreatOverride"] != (override != nil) {
						t.Fatal("replay metadata conceals resolved provisional parameters or optional presence")
					}
				} else if models != nil {
					t.Fatal("Forever assumption metadata applied to a Classic control")
				}
				for _, target := range sim.Encounter.TargetUnits {
					metrics := spell.SpellMetrics[target.UnitIndex]
					if metrics.Hits == 0 || metrics.TotalDamage != 0 {
						t.Fatal("Demoralizing effect invented damage or did not land")
					}
					got := metrics.TotalThreat / float64(metrics.Hits) / c.PseudoStats.ThreatMultiplier
					if math.Abs(got-want) > 1e-7 {
						t.Errorf("%s %v override=%v target=%d: base threat %v, want %v", tc.key, ruleset, override, target.UnitIndex, got, want)
					}
				}
			}
		}
	}
}

func TestOctober9ImpaleRequiresDeepWounds(t *testing.T) {
	var b build
	for _, candidate := range builds() {
		if candidate.Key == "fury" {
			b = candidate
			break
		}
	}
	config := loadTalents(b)
	points, err := config.decode(b.presetTalents())
	if err != nil {
		t.Fatal(err)
	}
	if err := config.validate(config.encode(points)); err != nil {
		t.Fatal(err)
	}
	deep, impale, recipient := -1, -1, -1
	for i, talent := range config.Trees[0].Talents {
		switch talent.Field {
		case "deepWounds":
			deep = i
		case "impale":
			impale = i
		default:
			if recipient < 0 && talent.Location.Row < 3 && points[i] < talent.Max && !talent.NotSimulated {
				recipient = i
			}
		}
	}
	if deep < 0 || impale < 0 || recipient < 0 || points[deep] != 3 || points[impale] == 0 {
		t.Fatal("fixture does not isolate the restored prerequisite")
	}
	points[deep]--
	points[recipient]++
	if err := config.validate(config.encode(points)); err == nil || !strings.Contains(err.Error(), "deepWounds") {
		t.Fatalf("Impale without full Deep Wounds accepted: %v", err)
	}
}

func TestOctober9BearOwnsFaerieFire(t *testing.T) {
	if core.ForeverTankSupport(proto.Class_ClassDruid).Debuffs.FaerieFire {
		t.Fatal("external permanent Faerie Fire masks the Bear's explicit self-maintenance duty")
	}
	for _, class := range []proto.Class{proto.Class_ClassWarrior, proto.Class_ClassPaladin} {
		if !core.ForeverTankSupport(class).Debuffs.FaerieFire {
			t.Fatal("Bear duty correction changed another tank's declared external support")
		}
	}
	for _, b := range builds() {
		if b.Key != "feral_tank_druid" {
			continue
		}
		for _, race := range b.races() {
			for _, targets := range []int{1, 3} {
				duration := 300.0
				if targets == 3 {
					duration = 90
				}
				request := tankRequest(b, b.player(race), 1, 1800400001, targets, duration)
				if request.Raid.Debuffs.FaerieFire {
					t.Fatal("native tank request reintroduced free Faerie Fire")
				}
				result := core.RunRaidSim(request)
				if result.Error != nil {
					t.Fatal(result.Error)
				}
				casts := int32(0)
				for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
					if action.Id.GetSpellId() == 9907 {
						for _, target := range action.Targets {
							casts += target.Casts
						}
					}
				}
				if casts == 0 {
					t.Fatalf("%s/%d attackers: self Faerie Fire was never cast", raceName(race), targets)
				}
			}
		}
	}
}
