//go:build with_db

package main

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func TestForeverThistleTeaRogueAndCat(t *testing.T) {
	for _, key := range []string{"combat", "feral", "balance", "fury"} {
		race := proto.Race_RaceOrc
		if key == "feral" || key == "balance" {
			race = proto.Race_RaceTauren
		}
		req := racialFixture(key, race)
		p := req.Raid.Parties[0].Players[0]
		p.Rotation = &proto.APLRotation{}
		p.Consumes.DefaultConjured = proto.Conjured_ConjuredRogueThistleTea
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		character := sim.Raid.Parties[0].Players[0].GetCharacter()
		actionID := core.ActionID{ItemID: 7676}
		spell := character.GetSpell(actionID)
		if key == "balance" || key == "fury" {
			if spell != nil {
				t.Fatalf("%s registered an unusable energy consumable", key)
			}
			continue
		}
		count := 0
		for _, registered := range character.Spellbook {
			if registered.ActionID == actionID {
				count++
			}
		}
		if count != 1 || spell == nil || spell.CD.Duration != 5*time.Minute {
			t.Fatalf("%s tea: %d registrations, incorrect cooldown or missing spell", key, count)
		}
		character.SpendEnergy(sim, character.CurrentEnergy(), character.NewEnergyMetrics(actionID))
		if !spell.Cast(sim, &character.Unit) || character.CurrentEnergy() != 100 {
			t.Fatalf("%s tea did not restore 100 Energy", key)
		}
		if spell.IsReady(sim) {
			t.Fatalf("%s tea ignored its item cooldown", key)
		}
	}
}
