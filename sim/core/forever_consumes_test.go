package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestForeverFoodStats(t *testing.T) {
	for _, tc := range []struct {
		food             proto.Food
		forever, classic stats.Stats
	}{
		{proto.Food_FoodGrilledSquid, stats.Stats{stats.MeleeCrit: CritRatingPerCritChance, stats.SpellCrit: SpellCritRatingPerCritChance}, stats.Stats{stats.Agility: 10}},
		{proto.Food_FoodNightfinSoup, stats.Stats{stats.SpellDamage: 22}, stats.Stats{stats.MP5: 8}},
		{proto.Food_FoodRunnTumTuberSurprise, stats.Stats{stats.Intellect: 15}, stats.Stats{stats.Intellect: 10}},
	} {
		for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
			t.Run(tc.food.String()+"/"+ruleset.String(), func(t *testing.T) {
				c := &Character{Unit: Unit{Env: &Environment{Ruleset: ruleset}}}
				applyFoodConsumes(c, &proto.Consumes{Food: tc.food})
				want := tc.classic
				if ruleset == proto.Ruleset_RulesetForever {
					want = tc.forever
				}
				if got := c.GetStats(); got != want {
					t.Fatalf("food stats = %v, want %v", got, want)
				}
			})
		}
	}
}
