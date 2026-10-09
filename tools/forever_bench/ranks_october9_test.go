//go:build with_db

package main

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
)

// SpellEffect/SpellLevels/SpellPower 1.60.1.70291; public patch context:
// https://us.forums.blizzard.com/en/wow/t/2360696/5
// Client variance spreads the unscaled base; capped level growth adds to the
// midpoint. Keep these floating values, rather than inventing integer endpoint
// rounding. The retained Classic columns independently guard the prior model.
func TestOctober9EarlySpellRanks(t *testing.T) {
	rows := []struct {
		family, key                                           string
		id, tag                                               int32
		level                                                 int
		low, high, classicLow, classicHigh, coefficient, cost float64
	}{
		{"Wrath", "balance", 5176, 0, 1, 14.64615381507, 16.95384620877, 15, 18, 0.429, 10},
		{"Wrath", "balance", 5177, 0, 6, 21.24444450439, 24.35555563864, 25, 28, 0.486, 20},
		{"Wrath", "balance", 5178, 0, 14, 30.49999992555, 35.50000007445, 32, 37, 0.571, 40},
		{"Wrath", "balance", 5179, 0, 22, 37.87941179414, 43.32058849194, 41, 46, 0.571, 50},
		{"Wrath", "balance", 5180, 0, 30, 46.2833333239, 52.11666653305, 47, 53, 0.571, 70},
		{"Frostbolt", "arcane", 116, 0, 4, 18.80000000446, 20.80000001938, 20, 22, 0.407, 25},
		{"Frostbolt", "arcane", 205, 0, 8, 28.44242434581, 31.95757574954, 33, 38, 0.489, 35},
		{"Frostbolt", "arcane", 837, 0, 14, 45.44444442538, 50.55555557462, 46, 53, 0.597, 50},
		{"Frostbolt", "arcane", 7322, 0, 20, 69.61025667193, 76.78974342343, 61, 68, 0.706, 65},
		{"Frostbolt", "arcane", 8406, 0, 26, 104.472726941, 113.9272726775, 97, 105, 0.814, 100},
		{"Fireball", "arcane", 133, 0, 1, 14.79999998208, 22.80000004176, 16, 25, 0.429, 30},
		{"Fireball", "arcane", 143, 0, 6, 27.3052632808, 39.09473681456, 32, 47, 0.571, 45},
		{"Fireball", "arcane", 145, 0, 12, 45.30476172269, 61.49523846803, 48, 65, 0.714, 65},
		{"Fireball", "arcane", 3140, 0, 18, 68.28000017991, 92.91999962937, 66, 91, 0.857, 95},
		{"Fireball", "arcane", 8400, 0, 24, 107.9312877656, 143.2687120436, 98, 130, 1.000, 140},
		{"Smite", "smite", 585, 0, 1, 13.49999990318, 17.50000011173, 15, 20, 0.429, 20},
		{"Smite", "smite", 591, 0, 6, 25.60714288807, 31.39285723114, 28, 34, 0.571, 30},
		{"Smite", "smite", 598, 0, 14, 44.39655173568, 50.60344826433, 47, 53, 0.714, 60},
		{"Smite", "smite", 984, 0, 22, 59.2142854632, 67.7857144176, 60, 68, 0.714, 95},
		{"Lightning Bolt", "elemental", 403, 0, 1, 14.42857138815, 16.57142862675, 15, 17, 0.429, 15},
		{"Lightning Bolt", "elemental", 529, 0, 8, 25.57142849271, 29.42857152219, 29, 34, 0.571, 30},
		{"Lightning Bolt", "elemental", 548, 0, 14, 42.82653074705, 50.17346937215, 44, 52, 0.714, 45},
		{"Lightning Bolt", "elemental", 915, 0, 20, 58.88764055074, 67.11235950886, 55, 63, 0.714, 60},
		{"Lightning Bolt", "elemental", 943, 0, 26, 78.99253702147, 90.00746297853, 70, 81, 0.714, 85},
		{"Shadow Bolt", "affliction", 686, 0, 1, 11.94285707174, 15.6571429521, 12, 16, 0.486, 25},
		{"Shadow Bolt", "affliction", 695, 0, 6, 21.84615389252, 27.15384622668, 25, 31, 0.629, 40},
		{"Shadow Bolt", "affliction", 705, 0, 12, 39.9230769276, 46.0769233108, 41, 48, 0.800, 70},
		{"Shadow Bolt", "affliction", 1088, 0, 20, 58.21739140161, 65.78260871759, 57, 64, 0.857, 110},
		{"AM children", "arcane", 5143, 1, 8, 24.80000001192, 24.80000001192, 26, 26, 0.286, 85},
		{"AM children", "arcane", 5144, 1, 16, 35.20000004768, 35.20000004768, 33, 33, 0.286, 140},
		{"AM children", "arcane", 5145, 1, 24, 48.60000002384, 48.60000002384, 46, 46, 0.286, 235},
	}
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		for _, row := range rows {
			t.Run(fmt.Sprintf("%v/%s/%d", ruleset, row.family, row.id), func(t *testing.T) {
				req := historyTalentFixture(row.key, nil)
				p := req.Raid.Parties[0].Players[0]
				p.Equipment, p.Rotation = &proto.EquipmentSpec{}, &proto.APLRotation{}
				p.ForeverTier1Bonuses, p.BonusStats = false, nil
				req.Encounter.Targets[0].Level = 60
				req.SimOptions.Ruleset, req.SimOptions.IsTest, req.SimOptions.RandomSeed = ruleset, true, 70291
				sim := core.NewSim(req, simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
				parent := unit.GetSpell(core.ActionID{SpellID: row.id})
				spell := unit.GetSpell(core.ActionID{SpellID: row.id, Tag: row.tag})
				if spell == nil || parent == nil || parent.RequiredLevel != row.level || parent.Cost.BaseCost != row.cost {
					t.Fatal("rank lost its registered identity, required level, or unchanged mana cost")
				}
				if math.Abs(spell.BonusCoefficient-row.coefficient) > 1e-9 || spell.ProcMask != core.ProcMaskSpellDamage || spell.DefenseType != core.DefenseTypeMagic {
					t.Fatal("unchanged spell-power coefficient or harmful spell flags changed")
				}
				low, high := row.low, row.high
				if ruleset == proto.Ruleset_RulesetClassic {
					low, high = row.classicLow, row.classicHigh
				}
				spell.BonusCoefficient, spell.MissileSpeed = 0, 0
				spell.BonusHitRating, spell.BonusCritRating = 10000, -10000
				spell.Flags |= core.SpellFlagIgnoreModifiers | core.SpellFlagIgnoreResists
				// The reference uses an independent, identically seeded Damage Roll stream.
				// Outcomes/procs cannot perturb that stream when IsTest is enabled.
				reference := core.NewSim(req, simsignals.Signals{})
				reference.Reset()
				for sample := 0; sample < 3; sample++ {
					want := reference.Roll(low, high)
					before := spell.SpellMetrics[target.UnitIndex].TotalDamage
					spell.ApplyEffects(sim, target, spell)
					advanceHawks(t, sim, sim.CurrentTime)
					got := spell.SpellMetrics[target.UnitIndex].TotalDamage - before
					if math.Abs(got-want) > 1e-7 {
						t.Fatalf("sample%d damage %.12g, want %.12g from [%.12g,%.12g]", sample, got, want, low, high)
					}
				}
			})
		}
	}
}

func TestOctober9PenanceNumbers(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		for _, healing := range []int{0, 3} {
			t.Run(fmt.Sprintf("%v/discount%d", ruleset, healing), func(t *testing.T) {
				req := historyTalentFixture("smite", map[string]int{"penance": 1, "improvedHealing": healing})
				p := req.Raid.Parties[0].Players[0]
				p.Equipment, p.Rotation = &proto.EquipmentSpec{}, &proto.APLRotation{}
				p.ForeverTier1Bonuses, p.BonusStats = false, nil
				req.Encounter.Targets[0].Level = 60
				req.SimOptions.Ruleset = ruleset
				sim := core.NewSim(req, simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
				rows := []struct {
					id                      int32
					rank, level             int
					base, cost, coefficient float64
				}{
					{402174, 1, 30, 44.60999992494, 150, .19},  // Damage bolt402284, cap39.
					{1240720, 2, 40, 52.97000011799, 220, .19}, // Bolt1240727, cap49.
					{1240721, 3, 50, 71.87000006435, 270, .19}, // Bolt1240730, cap59.
					{1316995, 4, 60, 92, 385, .19},             // Bolt1316993.
				}
				if ruleset == proto.Ruleset_RulesetClassic {
					rows = rows[3:]
					rows[0].base, rows[0].cost, rows[0].coefficient, rows[0].rank = 131, 355, .285, 1
				}
				max := unit.GetSpell(core.ActionID{SpellID: 1316995})
				for _, row := range rows {
					t.Run(fmt.Sprintf("rank%d", row.rank), func(t *testing.T) {
						spell := unit.GetSpell(core.ActionID{SpellID: row.id})
						if spell == nil || spell.Cost.BaseCost != row.cost || spell.Rank != row.rank || spell.RequiredLevel != row.level {
							t.Fatalf("incorrect Penance rank%d identity/cost", row.rank)
						}
						if spell.CD.Timer != max.CD.Timer || spell.CD.Duration != 12*time.Second {
							t.Fatal("Penance ranks must share their12-second cooldown")
						}
						if math.Abs(spell.Cost.GetCurrentCost()-row.cost*(1-.05*float64(healing))) > 1e-7 {
							t.Fatal("Improved Healing did not use the new base mana cost")
						}
						dot := spell.Dot(target)
						if dot.BonusCoefficient != row.coefficient || dot.NumberOfTicks != 2 || dot.TickLength != time.Second {
							t.Fatal("Penance coefficient or retained two-second/three-bolt timing changed")
						}
						spell.Flags |= core.SpellFlagIgnoreModifiers | core.SpellFlagIgnoreResists
						spell.BonusHitRating = 10000
						spell.ApplyEffects(sim, target, spell)
						want := row.base + row.coefficient*spell.GetBonusDamage(target)
						if math.Abs(dot.SnapshotBaseDamage-want) > 1e-7 {
							t.Fatalf("rank%d bolt snapshot %g, want %g", row.rank, dot.SnapshotBaseDamage, want)
						}
						first := spell.SpellMetrics[target.UnitIndex].TotalDamage
						if math.Abs(first-want) > 1e-7 {
							t.Fatalf("rank%d first bolt %g, want %g", row.rank, first, want)
						}
						unit.AddStatDynamic(sim, stats.SpellPower, 100)
						expected := spell.CalcPeriodicDamage(sim, target, row.base, spell.OutcomeAlwaysHit)
						if math.Abs(expected.Damage-(want+row.coefficient*100)) > 1e-7 {
							t.Fatal("Penance SP coefficient changed outside its source slope")
						}
						unit.AddStatDynamic(sim, stats.SpellPower, -100)
						if dot.Duration != 2*time.Second {
							t.Fatal("Penance duration was changed by a numerical rank update")
						}
						dot.Deactivate(sim)
					})
				}
				if ruleset == proto.Ruleset_RulesetClassic && unit.GetSpell(core.ActionID{SpellID: 402174}) != nil {
					t.Fatal("new Forever ranks leaked into Classic")
				}
			})
		}
	}
}

// 70291 Fireball periodic effects679021/679023/679025 are per-tick
// base1/1/2, period2000ms. Duration35/32/32 is4000/6000/6000ms:
// totals2/3/6, not four ticks spread over8s. Higher ranks stay four ticks.
func TestOctober9FireballLowRankDot(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		for _, row := range []struct {
			id    int32
			ticks int32
			total float64
		}{{133, 2, 2}, {143, 3, 3}, {145, 3, 6}, {3140, 4, 12}, {25306, 4, 60}} {
			t.Run(fmt.Sprintf("%v/%d", ruleset, row.id), func(t *testing.T) {
				req := historyTalentFixture("arcane", nil)
				p := req.Raid.Parties[0].Players[0]
				p.Equipment, p.Rotation = &proto.EquipmentSpec{}, &proto.APLRotation{}
				p.ForeverTier1Bonuses, p.BonusStats = false, nil
				req.Encounter.Targets[0].Level = 60
				req.SimOptions.Ruleset, req.SimOptions.IsTest = ruleset, true
				sim := core.NewSim(req, simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
				spell := unit.GetSpell(core.ActionID{SpellID: row.id})
				if spell == nil {
					t.Fatal("registered Fireball rank missing")
				}
				ticks := row.ticks
				if ruleset == proto.Ruleset_RulesetClassic {
					ticks = 4
				}
				dot := spell.Dot(target)
				if dot.NumberOfTicks != ticks || dot.TickLength != 2*time.Second || dot.Duration != time.Duration(ticks)*2*time.Second {
					t.Fatalf("DoT %d ticks/%v duration; want%d ticks/%v", dot.NumberOfTicks, dot.Duration, ticks, time.Duration(ticks)*2*time.Second)
				}
				spell.BonusCoefficient, spell.MissileSpeed = 0, 0
				spell.BonusHitRating, spell.BonusCritRating = 10000, -10000
				spell.Flags |= core.SpellFlagIgnoreModifiers | core.SpellFlagIgnoreResists
				spell.ApplyEffects(sim, target, spell)
				advanceHawks(t, sim, sim.CurrentTime)
				perTick := row.total / float64(ticks)
				if math.Abs(dot.SnapshotBaseDamage-perTick) > 1e-9 {
					t.Fatalf("per-tick source base %g, want%g", dot.SnapshotBaseDamage, perTick)
				}
				direct := spell.SpellMetrics[target.UnitIndex].TotalDamage
				for tick := int32(1); tick <= ticks; tick++ {
					at := time.Duration(tick) * 2 * time.Second
					// The helper steps through at+1ns; keep the pre-tick probe
					// strictly before the tick rather than tying its priority.
					advanceHawks(t, sim, at-time.Microsecond)
					before := spell.SpellMetrics[target.UnitIndex].TotalDamage - direct
					if math.Abs(before-float64(tick-1)*perTick) > 1e-9 {
						t.Fatalf("tick%d arrived early: %g", tick, before)
					}
					advanceHawks(t, sim, at)
					got := spell.SpellMetrics[target.UnitIndex].TotalDamage - direct
					if math.Abs(got-float64(tick)*perTick) > 1e-9 {
						t.Fatalf("tick%d damage %g", tick, got)
					}
				}
				advanceHawks(t, sim, 10*time.Second)
				if dot.IsActive() || math.Abs(spell.SpellMetrics[target.UnitIndex].TotalDamage-direct-row.total) > 1e-9 {
					t.Fatal("unexpected extra tick/DoT damage after source duration")
				}
			})
		}
	}
}
