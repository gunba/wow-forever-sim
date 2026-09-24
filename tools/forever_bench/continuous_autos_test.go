//go:build with_db

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/shaman"
)

func TestForeverHunterAutoShotUsesClassicWindup(t *testing.T) {
	var hunter build
	for _, b := range builds() {
		if b.Key == "marksmanship" {
			hunter = b
		}
	}
	casts := func(spellID int32) int32 {
		player := hunter.player(proto.Race_RaceOrc)
		// Register Sniper Shot without random attack-speed talents.
		player.TalentsString = historyTalentFixture("marksmanship", map[string]int{"sniperShot": 1}).Raid.Parties[0].Players[0].TalentsString
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
		var count, special, specialHits int32
		for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
			for _, target := range action.Targets {
				if action.Id.GetOtherId() == proto.OtherAction_OtherActionShoot {
					count += target.Casts
				}
				if action.Id.GetSpellId() == spellID && spellID != 0 {
					special += target.Casts
					specialHits += target.Hits + target.Crits
				}
			}
		}
		if spellID != 0 && (special == 0 || specialHits == 0) {
			t.Fatalf("special shot %d started %d times but landed %d hits; its completion may have been overwritten", spellID, special, specialHits)
		}
		return count
	}
	idle := casts(0)
	for _, id := range []int32{20904, 2643, 1310786} {
		if got := casts(id); got > idle || got == 0 {
			t.Errorf("%d: Auto Shots %d, idle %d; hardcasts cannot add free shots", id, got, idle)
		}
	}

	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetClassic, proto.Ruleset_RulesetForever} {
		req := request(hunter.player(proto.Race_RaceOrc), 1, 321)
		req.SimOptions.Ruleset = ruleset
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		character := sim.Raid.Parties[0].Players[0].GetCharacter()
		auto := character.AutoAttacks.RangedAuto()
		aimed := character.GetSpell(core.ActionID{SpellID: 20904})
		if aimed.DefaultCast.CastTime != 2*time.Second ||
			auto.DefaultCast.CastTime != 500*time.Millisecond ||
			!auto.Flags.Matches(core.SpellFlagCastTimeNoGCD) ||
			auto.ExtraCastCondition == nil {
			t.Fatalf("%v: lost Classic Auto Shot wind-up or special-shot timing", ruleset)
		}
		want := time.Duration(float64(500*time.Millisecond) / character.RangedSwingSpeed())
		if auto.CastTime() != want {
			t.Fatalf("%v: wind-up %s, want speed-scaled %s", ruleset, auto.CastTime(), want)
		}
		if !aimed.Cast(sim, character.CurrentTarget) {
			t.Fatalf("%v: Aimed Shot could not start", ruleset)
		}
		if auto.CanCast(sim, character.CurrentTarget) {
			t.Fatalf("%v: Auto Shot started during Aimed Shot", ruleset)
		}
		sim.CurrentTime = character.Hardcast.Expires
		if auto.CanCast(sim, character.CurrentTarget) {
			t.Fatalf("%v: Auto Shot replaced the special shot at its completion timestamp", ruleset)
		}
	}
}

func TestShamanCastsRestoreClassicSwingReset(t *testing.T) {
	for _, stacks := range []int32{0, 1, 3, 5} {
		for _, spellID := range []int32{15208, 10605} {
			t.Run(fmt.Sprintf("%d/%d-stacks", spellID, stacks), func(t *testing.T) {
				req := historyTalentFixture("enhancement", map[string]int{"maelstromWeapon": 5})
				sim := core.NewSim(req, simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				s := sim.Raid.Parties[0].Players[0].(shaman.ShamanAgent).GetShaman()
				if stacks > 0 {
					s.MaelstromWeaponAura.Activate(sim)
					s.MaelstromWeaponAura.SetStacks(sim, stacks)
				}
				spell := s.GetSpell(core.ActionID{SpellID: spellID})
				castTime := spell.CastTime()
				if !spell.Cast(sim, s.CurrentTarget) {
					t.Fatal("spell was not cast")
				}
				want := sim.CurrentTime + castTime + s.AutoAttacks.MainhandSwingSpeed()
				if got := s.AutoAttacks.MainhandSwingAt(); got != want {
					t.Fatalf("next swing %s, want cast end + full swing timer %s", got, want)
				}
				if castTime > 0 && s.AutoAttacks.MHAuto().CanCast(sim, s.CurrentTarget) {
					t.Fatal("melee was allowed inside an ordinary hardcast")
				}
				if spellID == 10605 && s.MaelstromWeaponAura.GetStacks() != stacks {
					t.Fatal("Chain Lightning consumed Maelstrom stacks")
				}
			})
		}
	}
}

func TestForeverSlamKeepsMeleeSwings(t *testing.T) {
	req := historyTalentFixture("arms", nil)
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	unit := sim.Raid.AllPlayerUnits[0]
	unit.AddRage(sim, 100, unit.NewRageMetrics(core.ActionID{SpellID: 1464}))
	slam := unit.GetSpell(core.ActionID{SpellID: 11605})
	before := unit.AutoAttacks.MainhandSwingAt()
	if !slam.Cast(sim, unit.CurrentTarget) || !unit.IsCasting(sim) {
		t.Fatal("Slam hardcast did not start")
	}
	if unit.AutoAttacks.MainhandSwingAt() != before || !unit.AutoAttacks.MHAuto().CanCast(sim, unit.CurrentTarget) {
		t.Fatal("Slam delayed or blocked a melee auto")
	}
}

func TestShamanConsecutiveHardcastsDoNotWeaveMelee(t *testing.T) {
	req := historyTalentFixture("enhancement", nil)
	p := req.Raid.Parties[0].Players[0]
	p.Rotation = core.APLRotationFromJsonString(`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":15208}}}}]}`)
	req.Encounter.Duration = 20
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	var bolts, swings int32
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		for _, target := range action.Targets {
			if action.Id.GetSpellId() == 15208 {
				bolts += target.Casts
			}
			if action.Id.GetOtherId() == proto.OtherAction_OtherActionAttack {
				swings += target.Casts
			}
		}
	}
	// The opening swing at t=0 can precede the first cast.
	if bolts < 4 || swings > 1 {
		t.Fatalf("back-to-back casts: bolts=%d melee=%d", bolts, swings)
	}
}
