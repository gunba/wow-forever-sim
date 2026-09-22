//go:build with_db

package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestCritAuraProvidersAndStacking(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetClassic, proto.Ruleset_RulesetForever} {
		var baseline stats.Stats
		for providers := 0; providers < 16; providers++ {
			t.Run(fmt.Sprintf("%v/%04b", ruleset, providers), func(t *testing.T) {
				req := racialFixture("arcane", proto.Race_RaceHuman)
				req.Raid.Buffs.MoonkinAura = providers&1 != 0
				req.Raid.Buffs.LeaderOfThePack = providers&2 != 0
				req.Raid.Parties[0].Buffs.MoonkinAura = providers&4 != 0
				req.Raid.Parties[0].Buffs.LeaderOfThePack = providers&8 != 0
				env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, ruleset, false)
				got := env.Raid.AllPlayerUnits[0].GetStats()
				if providers == 0 {
					baseline = got
				}
				wantMelee, wantSpell := 0.0, 0.0
				if providers&10 != 0 || (ruleset == proto.Ruleset_RulesetForever && providers != 0) {
					wantMelee = 3 * core.CritRatingPerCritChance
				}
				if providers&5 != 0 || (ruleset == proto.Ruleset_RulesetForever && providers != 0) {
					wantSpell = 3 * core.SpellCritRatingPerCritChance
				}
				if math.Abs(got[stats.MeleeCrit]-baseline[stats.MeleeCrit]-wantMelee) > 1e-8 ||
					math.Abs(got[stats.SpellCrit]-baseline[stats.SpellCrit]-wantSpell) > 1e-8 {
					t.Fatalf("crit gain melee=%g spell=%g, want %g/%g",
						got[stats.MeleeCrit]-baseline[stats.MeleeCrit], got[stats.SpellCrit]-baseline[stats.SpellCrit], wantMelee, wantSpell)
				}
			})
		}
	}
}

func TestCritAuraStaysWithinItsParty(t *testing.T) {
	req := racialFixture("arcane", proto.Race_RaceHuman)
	second := racialFixture("arcane", proto.Race_RaceHuman)
	req.Raid.Parties[0].Buffs.LeaderOfThePack = true
	req.Raid.Parties = append(req.Raid.Parties, second.Raid.Parties[0])
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	first, other := env.Raid.AllPlayerUnits[0], env.Raid.AllPlayerUnits[1]
	for _, stat := range []stats.Stat{stats.MeleeCrit, stats.SpellCrit} {
		if math.Abs(first.GetStat(stat)-other.GetStat(stat)-3) > 1e-8 {
			t.Fatalf("party crit difference for %v = %g", stat, first.GetStat(stat)-other.GetStat(stat))
		}
	}
}
