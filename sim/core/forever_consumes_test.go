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

func TestOffHandImbuesRequireWeapon(t *testing.T) {
	for _, slot := range []struct {
		name  string
		item  Item
		valid bool
	}{
		{"empty", Item{}, false},
		{"shield", Item{ID: 2, WeaponType: proto.WeaponType_WeaponTypeShield}, false},
		{"frill", Item{ID: 3, WeaponType: proto.WeaponType_WeaponTypeOffHand}, false},
		{"weapon", Item{ID: 4, WeaponType: proto.WeaponType_WeaponTypeDagger}, true},
	} {
		for _, imbue := range []struct {
			value proto.WeaponImbue
			bonus stats.Stats
		}{
			{proto.WeaponImbue_ElementalSharpeningStone, stats.Stats{stats.MeleeCrit: 2 * CritRatingPerCritChance}},
			{proto.WeaponImbue_BrilliantWizardOil, stats.Stats{stats.SpellPower: 36, stats.SpellCrit: SpellCritRatingPerCritChance}},
		} {
			t.Run(slot.name+"/"+imbue.value.String(), func(t *testing.T) {
				c := &Character{Unit: Unit{PseudoStats: stats.NewPseudoStats()}}
				c.Equipment[proto.ItemSlot_ItemSlotOffHand] = slot.item
				applyWeaponImbueConsumes(c, &proto.Consumes{OffHandImbue: imbue.value})
				want := stats.Stats{}
				if slot.valid {
					want = imbue.bonus
				}
				if got := c.GetStats(); got != want {
					t.Fatalf("off-hand imbue stats = %v, want %v", got, want)
				}
			})
		}
	}
}
