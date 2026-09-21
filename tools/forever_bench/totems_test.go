//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
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
		grace := character.GetSpell(core.ActionID{SpellID: 25359})
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

func TestOwnTotemDoesNotRemoveExternalGrace(t *testing.T) {
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
		grace := character.GetSpell(core.ActionID{SpellID: 25359})
		windfury := character.GetSpell(core.ActionID{SpellID: 10614})
		grace.ApplyEffects(sim, target, grace)
		windfury.ApplyEffects(sim, target, windfury)
		if !character.GetAura("Grace of Air Totem").IsActive() {
			t.Fatal("replacing our totem removed another provider's permanent buff")
		}
	}
}
