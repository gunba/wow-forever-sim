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
				// Client 1.60.1.70425 splits them again under Forever too: Leader of the Pack is melee
				// and ranged crit only, Moonkin Aura spell crit only.
				wantMelee, wantSpell := 0.0, 0.0
				if providers&10 != 0 {
					wantMelee = 3 * core.CritRatingPerCritChance
				}
				if providers&5 != 0 {
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

// Raid-wide since client 1.60.1.70425: a Leader of the Pack in one party reaches the next, and
// stays within its party under Classic.
func TestCritAuraReachesTheWholeRaid(t *testing.T) {
	for _, tc := range []struct {
		ruleset proto.Ruleset
		want    float64
	}{{proto.Ruleset_RulesetForever, 0}, {proto.Ruleset_RulesetClassic, 3 * core.CritRatingPerCritChance}} {
		req := racialFixture("arcane", proto.Race_RaceHuman)
		second := racialFixture("arcane", proto.Race_RaceHuman)
		req.Raid.Parties[0].Buffs.LeaderOfThePack = true
		req.Raid.Parties = append(req.Raid.Parties, second.Raid.Parties[0])
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, tc.ruleset, false)
		first, other := env.Raid.AllPlayerUnits[0], env.Raid.AllPlayerUnits[1]
		if diff := first.GetStat(stats.MeleeCrit) - other.GetStat(stats.MeleeCrit); math.Abs(diff-tc.want) > 1e-8 {
			t.Fatalf("%v: party melee crit difference %g, want %g", tc.ruleset, diff, tc.want)
		}
		if diff := first.GetStat(stats.SpellCrit) - other.GetStat(stats.SpellCrit); math.Abs(diff) > 1e-8 {
			t.Fatalf("%v: Leader of the Pack changed spell crit by %g", tc.ruleset, diff)
		}
	}
}
