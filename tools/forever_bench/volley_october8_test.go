//go:build with_db

package main

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/hunter"
)

type volleyAutoEvent struct {
	start, complete time.Duration
	channeling      bool
}

func volleyOctober8Fixture(distance, haste float64, withPets bool) (*core.Simulation, *hunter.Hunter) {
	talents := map[string]int{}
	if withPets {
		talents["summonHawk"] = 1
	}
	req := historyTalentFixture("survival", talents)
	player := req.Raid.Parties[0].Players[0]
	player.Equipment = &proto.EquipmentSpec{}
	player.DistanceFromTarget = distance
	player.ForeverTier1Bonuses = false
	player.BonusStats = &proto.UnitStats{Stats: stats.Stats{
		stats.SpellHaste: haste * core.HasteRatingPerHastePercent,
		stats.SpellHit:   100,
	}.ToFloatArray()}
	player.GetHunter().Options.QuiverBonus = 0
	player.GetHunter().Options.PetType = proto.Hunter_Options_PetNone
	if withPets {
		player.GetHunter().Options.PetType = proto.Hunter_Options_Cat
	}
	req.Encounter.Duration = 12
	req.Encounter.Targets[0].Level = 60
	req.Encounter.Targets[0].Stats = stats.Stats{stats.Health: 1000000}.ToFloatArray()
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	h := sim.Raid.Parties[0].Players[0].(hunter.HunterAgent).GetHunter()
	// Fixed two-second weapons, without item procs or attack-speed talents.
	h.AutoAttacks.CancelAutoSwing(sim)
	h.AutoAttacks.SetMH(core.Weapon{BaseDamageMin: 100, BaseDamageMax: 100, SwingSpeed: 2, AttackPowerPerDPS: 14})
	h.AutoAttacks.SetRanged(core.Weapon{BaseDamageMin: 100, BaseDamageMax: 100, SwingSpeed: 2, AttackPowerPerDPS: 14})
	h.AutoAttacks.EnableAutoSwing(sim)
	return sim, h
}

func volleyOctober8Trace(spell *core.Spell, owner *hunter.Hunter) *[]volleyAutoEvent {
	events := new([]volleyAutoEvent)
	apply := spell.ApplyEffects
	spell.ApplyEffects = func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
		*events = append(*events, volleyAutoEvent{
			start: sim.CurrentTime - spell.CurCast.CastTime, complete: sim.CurrentTime,
			channeling: owner.IsChanneling(sim),
		})
		apply(sim, target, spell)
	}
	return events
}

func volleyOctober8Advance(t *testing.T, sim *core.Simulation, until time.Duration) {
	t.Helper()
	core.StartDelayedAction(sim, core.DelayedActionOptions{DoAt: until, OnAction: func(*core.Simulation) {}})
	for sim.CurrentTime < until {
		if sim.Step() {
			t.Fatalf("fight ended before %s", until)
		}
	}
}

func TestVolleyOctober8OwnerAutoLifecycle(t *testing.T) {
	for _, tc := range []struct {
		name            string
		distance, haste float64
		clip, nextDue   time.Duration
	}{
		{"melee_natural", 0, 0, 0, 0},
		{"melee_hasted", 0, 100, 0, 0},
		{"melee_clipped", 0, 100, 1500 * time.Millisecond, 0},
		{"ranged_natural", 20, 0, 0, 0},
		{"ranged_hasted", 20, 100, 0, 0},
		{"ranged_clipped", 20, 0, 1500 * time.Millisecond, 0},
		{"ranged_hasted_clipped", 20, 100, 1500 * time.Millisecond, 0},
		{"ranged_remaining_swing", 20, 100, 0, 4 * time.Second},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sim, h := volleyOctober8Fixture(tc.distance, tc.haste, false)
			auto := h.AutoAttacks.MHAuto()
			if tc.distance > 0 {
				auto = h.AutoAttacks.RangedAuto()
				h.AutoAttacks.DelayRangedUntil(sim, tc.nextDue)
			}
			events := volleyOctober8Trace(auto, h)
			if !h.Volley.Cast(sim, h.CurrentTarget) {
				t.Fatal("Volley was not castable")
			}
			dot := h.Volley.AOEDot()
			end := dot.Duration
			if tc.haste == 100 && end != 3*time.Second || tc.haste == 0 && end != 6*time.Second {
				t.Fatalf("unexpected Volley duration %s", end)
			}
			if auto.CanCast(sim, h.CurrentTarget) {
				t.Fatal("owner auto was castable during Volley")
			}
			if tc.clip > 0 {
				end = tc.clip
				core.StartDelayedAction(sim, core.DelayedActionOptions{DoAt: tc.clip, OnAction: func(sim *core.Simulation) { dot.Cancel(sim) }})
			}
			volleyOctober8Advance(t, sim, 8*time.Second)
			t.Logf("channel ended %s; actual auto casts: %+v", end, *events)
			if len(*events) == 0 {
				t.Fatal("owner never resumed auto attacks")
			}
			for _, event := range *events {
				if event.start < end || event.channeling {
					t.Fatalf("owner auto during channel: %+v, end %s", event, end)
				}
			}
			// The original global channel gate polls an overdue swing every100ms;
			// an unexpired remaining swing must not be shortened by the channel.
			want := max(end, tc.nextDue)
			first := (*events)[0]
			if first.start < want || first.start > want+100*time.Millisecond {
				t.Fatalf("first auto started %s, want %s..%s", first.start, want, want+100*time.Millisecond)
			}
			if tc.distance > 0 && first.complete-first.start != 500*time.Millisecond {
				t.Fatalf("Auto Shot wind-up %s, want original500ms", first.complete-first.start)
			}
			if tc.clip == 0 && dot.TickCount != 6 {
				t.Fatalf("natural Volley produced %d ticks, want6", dot.TickCount)
			}
		})
	}
}

func TestVolleyOctober8PetsKeepAttacking(t *testing.T) {
	sim, h := volleyOctober8Fixture(0, 100, true)
	companionHits, guardianHits := 0, 0
	previousDamage := make(map[*core.Pet]float64)
	// Deliver the real guardian launch without spending a second player GCD.
	h.SummonHawk.ApplyEffects(sim, h.CurrentTarget, h.SummonHawk)
	if !h.Volley.Cast(sim, h.CurrentTarget) {
		t.Fatal("Volley was not castable")
	}
	// Observe damage at each actual action time, not only an end-of-fight sum.
	for sim.CurrentTime < 3*time.Second {
		if sim.Step() {
			t.Fatal("fight ended during Volley")
		}
		for _, pet := range h.Pets {
			damage := 0.0
			for _, spell := range pet.Spellbook {
				for _, metrics := range spell.SpellMetrics {
					damage += metrics.TotalDamage
				}
			}
			if damage > previousDamage[pet] && h.IsChanneling(sim) {
				t.Logf("%s dealt %.3f damage at%s during owner channel", pet.Name, damage-previousDamage[pet], sim.CurrentTime)
				if pet.IsGuardian() {
					guardianHits++
				} else {
					companionHits++
				}
			}
			previousDamage[pet] = damage
		}
	}
	t.Logf("hits during owner channel: companion=%d, guardians=%d", companionHits, guardianHits)
	if companionHits == 0 || guardianHits == 0 {
		t.Fatal("Volley suppressed the independent companion or Hawk attacks")
	}
}
