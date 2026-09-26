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
	googleproto "google.golang.org/protobuf/proto"
)

func ownedTideFixture(t *testing.T, perSecond bool) (*core.Simulation, *shaman.Shaman) {
	t.Helper()
	req := historyTalentFixture("elemental", map[string]int{"manaTideTotem": 1, "totemicFocus": 5})
	p := req.Raid.Parties[0].Players[0]
	p.Equipment, p.BonusStats, p.ForeverTier1Bonuses = &proto.EquipmentSpec{}, nil, false
	p.ForeverMp5PerSecond = perSecond
	p.Name = "Owner"
	ally := googleproto.Clone(p).(*proto.Player)
	ally.Name = "Same party"
	outsider := googleproto.Clone(p).(*proto.Player)
	outsider.Name = "Other party"
	req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, ally)
	req.Raid.Parties = append(req.Raid.Parties, &proto.Party{Players: []*proto.Player{outsider}, Buffs: &proto.PartyBuffs{}})
	req.Raid.NumActiveParties = 2
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	for _, unit := range sim.Raid.AllPlayerUnits {
		unit.AddStatsDynamic(sim, stats.Stats{stats.Mana: 20000, stats.SpellHaste: 50 * core.HasteRatingPerHastePercent})
		unit.SpendMana(sim, unit.CurrentMana()-500, unit.NewManaMetrics(core.ActionID{SpellID: 1}))
	}
	return sim, sim.Raid.Parties[0].Players[0].(shaman.ShamanAgent).GetShaman()
}

func advanceTide(t *testing.T, sim *core.Simulation, at time.Duration) {
	t.Helper()
	done := false
	sim.AddPendingAction(&core.PendingAction{
		NextActionAt: at,
		Priority:     core.ActionPriorityLow,
		OnAction:     func(_ *core.Simulation) { done = true },
	})
	for !done {
		if sim.Step() {
			t.Fatalf("simulation ended before %s", at)
		}
	}
}

func tideMana(unit *core.Unit, id core.ActionID) float64 {
	var total float64
	for _, metric := range unit.Metrics.ToProto().Resources {
		if metric.Type == proto.ResourceType_ResourceTypeMana && id.SameAction(core.ProtoToActionID(metric.Id)) {
			total += metric.ActualGain
		}
	}
	return total
}

func TestOwnedManaTideRanksPartyTimingAndModes(t *testing.T) {
	for _, perSecond := range []bool{false, true} {
		sim, s := ownedTideFixture(t, perSecond)
		for rank, want := range []float64{0, 7.5, 22.5, 45} {
			if rank == 0 {
				continue
			}
			spell := s.ManaTideTotem[rank]
			if spell == nil || spell.Cost.GetCurrentCost() != want || spell.CD.Timer != s.ManaTideTotem[1].CD.Timer {
				t.Fatalf("rank %d cost or shared cooldown incorrect", rank)
			}
		}
		tide := s.ManaTideTotem[3]
		if tide.CD.Duration != 5*time.Minute || tide.DefaultCast.GCD != time.Second {
			t.Fatal("Mana Tide cooldown/GCD incorrect")
		}
		if !tide.Cast(sim, s.CurrentTarget) {
			t.Fatal("owned Mana Tide did not cast")
		}
		if s.CurrentMana() != 455 || s.ActiveTotems[shaman.WaterTotem] != tide || s.TotemExpirations[shaman.WaterTotem] != 12*time.Second {
			t.Fatal("owned Mana Tide did not pay mana/occupy the water slot")
		}
		ally := &sim.Raid.Parties[0].Players[1].GetCharacter().Unit
		other := &sim.Raid.Parties[1].Players[0].GetCharacter().Unit
		initial := tideMana(&s.Unit, tide.ActionID)
		advanceTide(t, sim, 2*time.Second)
		if tideMana(&s.Unit, tide.ActionID) != initial || tideMana(ally, tide.ActionID) != 0 {
			t.Fatal("Mana Tide pulsed before three seconds")
		}
		for tick := 1; tick <= 4; tick++ {
			advanceTide(t, sim, time.Duration(tick)*3*time.Second)
			want := 290 * float64(tick)
			ownerGain, allyGain, otherGain := tideMana(&s.Unit, tide.ActionID)-initial, tideMana(ally, tide.ActionID), tideMana(other, tide.ActionID)
			if math.Abs(ownerGain-want) > 1e-9 || math.Abs(allyGain-want) > 1e-9 || otherGain != 0 {
				t.Fatalf("mode %v tick %d: owner %v ally %v other party %v", perSecond, tick,
					ownerGain, allyGain, otherGain)
			}
		}
		advanceTide(t, sim, 15*time.Second)
		if math.Abs(tideMana(&s.Unit, tide.ActionID)-initial-1160) > 1e-9 || tide.Hot(&s.Unit).IsActive() {
			t.Fatal("Mana Tide continued after four pulses")
		}
	}
}

func TestOwnedWaterTotemsReplaceEachOther(t *testing.T) {
	sim, s := ownedTideFixture(t, false)
	healing, tide, spring := s.HealingStreamTotem[5], s.ManaTideTotem[3], s.ManaSpringTotem[4]
	healing.ApplyEffects(sim, s.CurrentTarget, healing)
	if !healing.Hot(&s.Unit).IsActive() {
		t.Fatal("Healing Stream did not activate")
	}
	if !tide.Cast(sim, s.CurrentTarget) {
		t.Fatal("Mana Tide could not replace Healing Stream")
	}
	for _, agent := range s.Party.Players {
		if healing.Hot(&agent.GetCharacter().Unit).IsActive() {
			t.Fatal("Healing Stream survived a water-slot replacement")
		}
	}
	initial := tideMana(&s.Unit, tide.ActionID)
	advanceTide(t, sim, 3*time.Second)
	if math.Abs(tideMana(&s.Unit, tide.ActionID)-initial-290) > 1e-9 {
		t.Fatal("first Tide pulse missing")
	}
	if !spring.Cast(sim, s.CurrentTarget) {
		t.Fatal("Mana Spring could not replace Mana Tide")
	}
	advanceTide(t, sim, 15*time.Second)
	if math.Abs(tideMana(&s.Unit, tide.ActionID)-initial-290) > 1e-9 || tide.Hot(&s.Unit).IsActive() || s.ActiveTotems[shaman.WaterTotem] != spring {
		t.Fatal("replaced Mana Tide kept restoring mana")
	}
}
