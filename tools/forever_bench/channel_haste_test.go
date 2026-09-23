//go:build with_db

package main

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/mage"
	"github.com/wowsims/classic/sim/priest"
	"github.com/wowsims/classic/sim/rogue"
)

func channelHasteFixture(buildKey string, ruleset proto.Ruleset, haste float64) *core.Simulation {
	req := racialFixture(buildKey, proto.Race_RaceUndead)
	player := req.Raid.Parties[0].Players[0]
	player.Equipment = &proto.EquipmentSpec{}
	player.ForeverTier1Bonuses = false
	player.Rotation = &proto.APLRotation{}
	player.BonusStats = &proto.UnitStats{Stats: stats.Stats{
		stats.SpellHaste: haste * core.HasteRatingPerHastePercent,
		stats.SpellHit:   100,
	}.ToFloatArray()}
	req.SimOptions.Ruleset = ruleset
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim
}

func TestForeverMindFlayChannelAndSpellGCDScaleWithHaste(t *testing.T) {
	for _, tc := range []struct {
		name    string
		ruleset proto.Ruleset
		haste   float64
		want    time.Duration
		gcd     time.Duration
	}{
		{"forever_no_haste", proto.Ruleset_RulesetForever, 0, 3 * time.Second, 1500 * time.Millisecond},
		{"forever_100_percent", proto.Ruleset_RulesetForever, 100, 1500 * time.Millisecond, time.Second},
		{"classic_100_percent", proto.Ruleset_RulesetClassic, 100, 3 * time.Second, 1500 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sim := channelHasteFixture("shadow", tc.ruleset, tc.haste)
			p := sim.Raid.Parties[0].Players[0].(priest.PriestAgent).GetPriest()
			spell := p.MindFlay[6][0]
			if !spell.Cast(sim, p.CurrentTarget) {
				t.Fatal("Mind Flay was not castable")
			}
			dot := spell.Dot(p.CurrentTarget)
			if dot.Duration != tc.want || p.GCD.ReadyAt() != tc.gcd {
				t.Fatalf("channel %s, GCD %s; want %s and %s", dot.Duration, p.GCD.ReadyAt(), tc.want, tc.gcd)
			}
			for sim.CurrentTime < tc.want {
				if sim.Step() {
					t.Fatalf("fight ended before final channel tick at %s", tc.want)
				}
			}
			if dot.TickCount != 3 {
				t.Fatalf("Mind Flay delivered %d ticks; want 3", dot.TickCount)
			}
		})
	}
}

func TestForeverMissileBarrageTickCountWithHaste(t *testing.T) {
	sim := channelHasteFixture("arcane", proto.Ruleset_RulesetForever, 100)
	m := sim.Raid.Parties[0].Players[0].(mage.MageAgent).GetMage()
	spell := m.ArcaneMissiles[len(m.ArcaneMissiles)-1]
	m.MissileBarrageAura.Activate(sim)
	if !spell.Cast(sim, m.CurrentTarget) {
		t.Fatal("Missile Barrage was not castable")
	}
	dot := spell.Dot(m.CurrentTarget)
	want := time.Duration(dot.NumberOfTicks) * 250 * time.Millisecond
	if dot.Duration != want {
		t.Fatalf("Barrage channel %s, want %s", dot.Duration, want)
	}
	for sim.CurrentTime < want {
		if sim.Step() {
			t.Fatal("fight ended before the final missile")
		}
	}
	if dot.TickCount != dot.NumberOfTicks {
		t.Fatalf("Barrage delivered %d ticks, want %d", dot.TickCount, dot.NumberOfTicks)
	}
}

func TestForeverPenanceKeepsImmediateBoltAtHaste(t *testing.T) {
	sim := channelHasteFixture("smite", proto.Ruleset_RulesetForever, 100)
	p := sim.Raid.Parties[0].Players[0].(priest.PriestAgent).GetPriest()
	spell := p.Penance
	if !spell.Cast(sim, p.CurrentTarget) {
		t.Fatal("Penance was not castable")
	}
	target := p.CurrentTarget.UnitIndex
	damage := spell.SpellMetrics[target].TotalDamage
	if damage == 0 || spell.Dot(p.CurrentTarget).Duration != time.Second {
		t.Fatal("the first bolt must land immediately, with two hasted bolts remaining")
	}
	for _, at := range []time.Duration{500 * time.Millisecond, time.Second} {
		for sim.CurrentTime < at {
			if sim.Step() {
				t.Fatalf("fight ended before Penance bolt at %s", at)
			}
		}
		next := spell.SpellMetrics[target].TotalDamage
		if sim.CurrentTime != at || next <= damage {
			t.Fatalf("Penance did not deal a bolt at %s", at)
		}
		damage = next
	}
}

func TestRogueGCDDoesNotInheritSpellHaste(t *testing.T) {
	sim := channelHasteFixture("combat", proto.Ruleset_RulesetForever, 100)
	r := sim.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
	if !r.SinisterStrike.Cast(sim, r.CurrentTarget) {
		t.Fatal("Sinister Strike was not castable")
	}
	if r.GCD.ReadyAt() != time.Second {
		t.Fatalf("Rogue GCD %s, want 1s", r.GCD.ReadyAt())
	}
}
