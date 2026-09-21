//go:build with_db

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestForeverHunterContinuousAutoShot(t *testing.T) {
	var hunter build
	for _, b := range builds() {
		if b.Key == "marksmanship" {
			hunter = b
		}
	}
	casts := func(spellID int32) int32 {
		player := hunter.player(proto.Race_RaceOrc)
		player.TalentsString = "" // No random attack-speed procs.
		apl := `{"type":"TypeAPL"}`
		if spellID != 0 {
			apl = fmt.Sprintf(`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":%d}}}}]}`, spellID)
		}
		player.Rotation = core.APLRotationFromJsonString(apl)
		req := request(player, 1, 321)
		req.Encounter.Duration = 30
		result := core.RunRaidSim(req)
		if result.Error != nil {
			t.Fatal(result.Error)
		}
		var count, special int32
		for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
			for _, target := range action.Targets {
				if action.Id.GetOtherId() == proto.OtherAction_OtherActionShoot {
					count += target.Casts
				}
				if action.Id.GetSpellId() == spellID && spellID != 0 {
					special += target.Casts
				}
			}
		}
		if spellID != 0 && special == 0 {
			t.Fatalf("test never cast %d", spellID)
		}
		return count
	}
	idle := casts(0)
	for _, id := range []int32{20904, 2643} {
		if got := casts(id); got != idle || got == 0 {
			t.Errorf("%d: Auto Shots %d, idle %d", id, got, idle)
		}
	}
	player := hunter.player(proto.Race_RaceOrc)
	req := request(player, 1, 321)
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	character := env.Raid.Parties[0].Players[0].GetCharacter()
	aimed := character.GetSpell(core.ActionID{SpellID: 20904})
	if aimed.DefaultCast.CastTime != 2*time.Second || character.AutoAttacks.RangedAuto().DefaultCast.CastTime != 0 {
		t.Fatal("shot wind-up remains in the Forever configuration")
	}
}
