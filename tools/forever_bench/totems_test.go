//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestForeverAirTotemsDoNotTwist(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "enhancement" {
			continue
		}
		req := request(b.player(proto.Race_RaceOrc), 1, 1)
		req.Raid.Buffs = &proto.RaidBuffs{}
		req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Reset()
		character := sim.Raid.Parties[0].Players[0].GetCharacter()
		target := sim.Encounter.TargetUnits[0]
		windfury := character.GetSpell(core.ActionID{SpellID: 10614})
		grace := character.GetSpell(core.ActionID{SpellID: 10627})
		windwall := character.GetSpell(core.ActionID{SpellID: 15112})
		windfury.ApplyEffects(sim, target, windfury)
		wf := character.GetAura("Windfury (Rank 3)")
		// Represent the pulse without running the character's other scheduled actions.
		wf.Activate(sim)
		grace.ApplyEffects(sim, target, grace)
		if wf.IsActive() {
			t.Fatal("Windfury lingered after replacing its totem")
		}
		if !character.GetAura("Grace of Air Totem").IsActive() {
			t.Fatal("Grace of Air did not replace Windfury")
		}
		windwall.ApplyEffects(sim, target, windwall)
		if character.GetAura("Grace of Air Totem").IsActive() {
			t.Fatal("Grace of Air lingered after replacement")
		}
	}
}

// Since client 1.60.1.70425 Grace of Air is no longer exclusive with Windfury Totem, so another
// Shaman's Grace stays in force under this Shaman's own Windfury.
func TestOwnWindfuryKeepsExternalGrace(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "enhancement" {
			continue
		}
		req := request(b.player(proto.Race_RaceOrc), 1, 1)
		req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Reset()
		character := sim.Raid.Parties[0].Players[0].GetCharacter()
		target := sim.Encounter.TargetUnits[0]
		grace := character.GetSpell(core.ActionID{SpellID: 10627})
		windfury := character.GetSpell(core.ActionID{SpellID: 10614})
		grace.ApplyEffects(sim, target, grace)
		withGrace := character.GetStat(stats.Agility)
		windfury.ApplyEffects(sim, target, windfury)
		character.GetAura("Windfury (Rank 3)").Activate(sim)
		if !character.GetAura("Grace of Air Totem").IsActive() {
			t.Fatal("external Grace of Air should remain configured")
		}
		if got := character.GetStat(stats.Agility); got != withGrace {
			t.Fatalf("Windfury suppressed external Grace: agility %.1f, want %.1f", got, withGrace)
		}
	}
}
