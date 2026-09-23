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

func TestFeralDefaultActuallyUsesThistleTea(t *testing.T) {
	var player *proto.Player
	for _, b := range builds() {
		if b.Key == "feral" {
			player = b.player(proto.Race_RaceTauren)
			break
		}
	}
	if player == nil {
		t.Fatal("Feral build is missing")
	}
	if player.Consumes.GetDefaultConjured() != proto.Conjured_ConjuredRogueThistleTea {
		t.Fatal("Feral default does not equip Thistle Tea")
	}
	req := request(player, 1, 1)
	req.SimOptions.Iterations = 1
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		if action.Id.GetItemId() == 7676 {
			if action.Targets[0].Casts == 1 {
				return
			}
			t.Fatalf("Feral used Thistle Tea %d times, want once", action.Targets[0].Casts)
		}
	}
	t.Fatal("Feral default equipped Thistle Tea but never used it")
}
