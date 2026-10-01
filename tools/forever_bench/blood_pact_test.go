//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/warlock"
)

func TestLevel60BloodPactIncludesLevelGrowth(t *testing.T) {
	var baseline float64
	for _, tc := range []struct {
		buff proto.TristateEffect
		gain float64
	}{
		{proto.TristateEffect_TristateEffectMissing, 0},
		{proto.TristateEffect_TristateEffectRegular, 54},
		{proto.TristateEffect_TristateEffectImproved, 54},
	} {
		req := racialFixture("arcane", proto.Race_RaceHuman)
		req.Raid.Buffs.BloodPact = tc.buff
		req.Raid.Parties[0].Buffs.BloodPact = proto.TristateEffect_TristateEffectMissing
		req.Raid.Parties[0].Players[0].Buffs = &proto.IndividualBuffs{}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		stamina := sim.Raid.AllPlayerUnits[0].GetStat(stats.Stamina)
		if tc.buff == proto.TristateEffect_TristateEffectMissing {
			baseline = stamina
		}
		if math.Abs(stamina-baseline-tc.gain) > 1e-8 {
			t.Fatalf("%v adds %g Stamina, want %g", tc.buff, stamina-baseline, tc.gain)
		}
	}
}

func TestClassicImprovedBloodPactUnchanged(t *testing.T) {
	var base float64
	for _, enabled := range []bool{false, true} {
		req := racialFixture("arcane", proto.Race_RaceHuman)
		req.SimOptions.Ruleset = proto.Ruleset_RulesetClassic
		req.Raid.Buffs = &proto.RaidBuffs{}
		req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
		if enabled {
			req.Raid.Buffs.BloodPact = proto.TristateEffect_TristateEffectImproved
		}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		got := sim.Raid.AllPlayerUnits[0].GetStat(stats.Stamina)
		if !enabled {
			base = got
		} else if math.Abs(got-base-70) > 1e-7 {
			t.Fatal("Forever correction changed Classic improved Blood Pact")
		}
	}
}

func TestOwnedBloodPactFollowsEnabledImp(t *testing.T) {
	for _, external := range []bool{false, true} {
		for rank := 0; rank <= 3; rank++ {
			req := historyTalentFixture("affliction", map[string]int{"improvedImp": rank})
			p := req.Raid.Parties[0].Players[0]
			p.GetWarlock().Options.Summon = proto.WarlockOptions_Imp
			p.GetWarlock().Options.Sacrifice = proto.WarlockOptions_NoSummon
			p.Buffs = &proto.IndividualBuffs{}
			req.Raid.Buffs = &proto.RaidBuffs{}
			req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
			if external {
				req.Raid.Parties[0].Buffs.BloodPact = proto.TristateEffect_TristateEffectRegular
			}
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			w := sim.Raid.Parties[0].Players[0].(warlock.WarlockAgent).GetWarlock()
			before := w.GetStat(stats.Stamina)
			if w.CurrentHealth() != w.MaxHealth() {
				t.Error("initial health excludes the enabled Imp")
			}
			w.Imp.Disable(sim, false)
			if w.CurrentHealth() > w.MaxHealth() {
				t.Error("disabling Imp left health above the new maximum")
			}
			want := core.TernaryFloat64(external, 0, 54)
			if got := before - w.GetStat(stats.Stamina); math.Abs(got-want) > 1e-7 {
				t.Errorf("rank %v external %v: disabling Imp removes %v Stamina, want %v", rank, external, got, want)
			}
			w.Imp.Enable(sim, w.Imp)
			if math.Abs(w.GetStat(stats.Stamina)-before) > 1e-7 {
				t.Error("enabling Imp lost or duplicated Blood Pact")
			}
		}
	}
}

func TestBloodPactMultipleOwnersAndIterations(t *testing.T) {
	req := historyTalentFixture("affliction", nil)
	second := historyTalentFixture("affliction", nil).Raid.Parties[0].Players[0]
	req.Raid.Buffs = &proto.RaidBuffs{}
	req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
	req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, second)
	for _, p := range req.Raid.Parties[0].Players {
		p.Buffs = &proto.IndividualBuffs{}
		p.Rotation = &proto.APLRotation{}
		p.GetWarlock().Options.Summon = proto.WarlockOptions_Imp
		p.GetWarlock().Options.Sacrifice = proto.WarlockOptions_NoSummon
	}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	first := sim.Raid.Parties[0].Players[0].(warlock.WarlockAgent).GetWarlock()
	other := sim.Raid.Parties[0].Players[1].(warlock.WarlockAgent).GetWarlock()
	before := first.GetStat(stats.Stamina)
	first.Imp.Disable(sim, true)
	if first.GetStat(stats.Stamina) != before {
		t.Error("sacrifice removed a second Imp's provider")
	}
	other.Imp.Disable(sim, false)
	if math.Abs(first.GetStat(stats.Stamina)-before+54) > 1e-7 {
		t.Error("multiple Imps stacked or retained an absent provider")
	}
	req.SimOptions.Iterations = 3
	req.Encounter.Duration, req.Encounter.DurationVariation = 1, 0
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatalf("iteration reset: %v", result.Error)
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
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	first := sim.Raid.Parties[0].Players[0].GetCharacter()
	other := sim.Raid.Parties[1].Players[0].GetCharacter()
	if difference := first.GetStat(stats.Stamina) - other.GetStat(stats.Stamina); math.Abs(difference-54) > 1e-8 {
		t.Fatalf("Imp party Stamina advantage = %g, want 54", difference)
	}
}
