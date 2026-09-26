//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestLevel60BloodPactIncludesLevelGrowth(t *testing.T) {
	var baseline float64
	for _, tc := range []struct {
		buff proto.TristateEffect
		gain float64
	}{
		{proto.TristateEffect_TristateEffectMissing, 0},
		{proto.TristateEffect_TristateEffectRegular, 54},
		{proto.TristateEffect_TristateEffectImproved, 70},
	} {
		req := racialFixture("arcane", proto.Race_RaceHuman)
		req.Raid.Buffs.BloodPact = tc.buff
		req.Raid.Parties[0].Buffs.BloodPact = proto.TristateEffect_TristateEffectMissing
		req.Raid.Parties[0].Players[0].Buffs = &proto.IndividualBuffs{}
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		stamina := env.Raid.AllPlayerUnits[0].GetStat(stats.Stamina)
		if tc.buff == proto.TristateEffect_TristateEffectMissing {
			baseline = stamina
		}
		if math.Abs(stamina-baseline-tc.gain) > 1e-8 {
			t.Fatalf("%v adds %g Stamina, want %g", tc.buff, stamina-baseline, tc.gain)
		}
	}
}

func TestImpBloodPactDoesNotLeakAcrossParties(t *testing.T) {
	req := racialFixture("arcane", proto.Race_RaceHuman)
	second := racialFixture("arcane", proto.Race_RaceHuman)
	petOwner := racialFixture("affliction", proto.Race_RaceOrc).Raid.Parties[0].Players[0]
	petOwner.TalentsString = ""
	petOwner.GetWarlock().Options.Summon = proto.WarlockOptions_Imp
	req.Raid.Buffs.BloodPact = proto.TristateEffect_TristateEffectMissing
	req.Raid.Parties = append(req.Raid.Parties, second.Raid.Parties[0])
	for _, party := range req.Raid.Parties {
		party.Buffs.BloodPact = proto.TristateEffect_TristateEffectMissing
		party.Players[0].Buffs = &proto.IndividualBuffs{}
	}
	req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, petOwner)
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	first := env.Raid.Parties[0].Players[0].GetCharacter()
	other := env.Raid.Parties[1].Players[0].GetCharacter()
	if difference := first.GetStat(stats.Stamina) - other.GetStat(stats.Stamina); math.Abs(difference-54) > 1e-8 {
		t.Fatalf("Imp party Stamina advantage = %g, want 54", difference)
	}
}
