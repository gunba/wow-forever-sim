//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func TestWindfurySourcesRegisterOnce(t *testing.T) {
	for _, party := range []bool{false, true} {
		req := racialFixture("retribution", proto.Race_RaceHuman)
		req.SimOptions.Iterations, req.SimOptions.RandomSeed = 1, 20295711
		player := req.Raid.Parties[0].Players[0]
		player.Consumes = &proto.Consumes{}
		player.Rotation = &proto.APLRotation{}
		req.Raid.Parties[0].Buffs = &proto.PartyBuffs{WindfuryTotem: party}
		result := core.RunRaidSim(req)
		if result.Error != nil {
			t.Fatal(result.Error)
		}
		metrics := result.RaidMetrics.Parties[0].Players[0]
		count := 0
		for _, aura := range metrics.Auras {
			if aura.Id.GetSpellId() == 21919 {
				count++
				if aura.ProcsAvg != 1 {
					t.Fatalf("party=%v: extra-attack tracking activated %v times at reset", party, aura.ProcsAvg)
				}
			}
		}
		want := 0
		if party {
			want = 1
		}
		if count != want {
			t.Fatalf("party=%v: got %d extra-attack trackers, want %d", party, count, want)
		}
	}
}

func TestForeverRejectsLegacyWindfuryImbue(t *testing.T) {
	req := racialFixture("retribution", proto.Race_RaceHuman)
	req.Raid.Parties[0].Players[0].Consumes.MainHandImbue = proto.WeaponImbue_Windfury
	defer func() {
		if recover() == nil {
			t.Fatal("Forever silently accepted Windfury Totem as a personal weapon imbue")
		}
	}()
	core.NewSim(req, simsignals.Signals{})
}

func TestForeverFeralCanReceivePartyWindfury(t *testing.T) {
	req := racialFixture("feral", proto.Race_RaceTauren)
	req.Raid.Parties[0].Buffs = &proto.PartyBuffs{WindfuryTotem: true}
	req.Raid.Buffs.GraceOfAirTotem = proto.TristateEffect_TristateEffectMissing
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Reset()
	character := sim.Raid.Parties[0].Players[0].GetCharacter()
	if !character.PseudoStats.FeralCombatEnabled || character.GetAura("Windfury") == nil {
		t.Fatal("Forever's party Windfury did not register on the Feral combat form")
	}
}

func TestReferenceAirTotemSelection(t *testing.T) {
	melee := map[string]bool{
		"feral": true, "enhancement": true, "survival": true, "pet_melee": true,
		"retribution": true, "retribution_physical": true,
		"combat": true, "mutilate": true, "subtlety": true,
		"fury": true, "fury_sunder": true, "fury_2h": true, "arms": true,
		"tank_warrior": true, "protection_paladin": true, "feral_tank_druid": true,
	}
	for _, b := range builds() {
		race := b.races()[0]
		req := requestForBuild(b, b.player(race), 1, 20295719)
		hasWindfury := req.Raid.Parties[0].Buffs.WindfuryTotem
		hasGrace := req.Raid.Buffs.GraceOfAirTotem != proto.TristateEffect_TristateEffectMissing
		if hasWindfury != melee[b.Key] || hasGrace != (b.Key == "beast_mastery" || b.Key == "marksmanship") || hasWindfury && hasGrace {
			t.Errorf("%s: Windfury=%v Grace=%v", b.Key, hasWindfury, hasGrace)
		}
	}
}
