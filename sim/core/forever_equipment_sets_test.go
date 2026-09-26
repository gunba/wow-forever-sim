package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestForeverEquippedSetManaRegeneration(t *testing.T) {
	for _, tc := range []struct {
		id, pieces int32
		mp5        float64
	}{
		{181, 6, 8}, {182, 6, 8}, {183, 6, 8}, {185, 6, 8},
		{186, 6, 8}, {187, 6, 8}, {188, 6, 8}, {443, 3, 12}, {490, 2, 3},
	} {
		bonus := foreverEquipmentSets[tc.id].Bonuses[tc.pieces]
		if !bonus.Supported || bonus.Stats[stats.MP5] != tc.mp5 {
			t.Fatalf("set %d/%d: missing %g MP5", tc.id, tc.pieces, tc.mp5)
		}
		c := &Character{Unit: Unit{Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}}}
		for slot := int32(0); slot < tc.pieces; slot++ {
			c.Equipment[slot] = Item{SetID: tc.id}
		}
		found := 0
		for _, active := range c.foreverEquippedBonuses() {
			if active.NumPieces == tc.pieces {
				found++
			}
		}
		if found != 1 {
			t.Fatalf("set %d: expected the mana bonus once, found %d", tc.id, found)
		}
		c.Equipment[0] = Item{}
		for _, active := range c.foreverEquippedBonuses() {
			if active.NumPieces == tc.pieces {
				t.Fatalf("set %d: mana bonus survived removing the required piece", tc.id)
			}
		}
	}
}

func TestForeverOrdinarySetThresholds(t *testing.T) {
	c := &Character{Unit: Unit{Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}}}
	for slot := 0; slot < 3; slot++ {
		c.Equipment[slot] = Item{SetID: 489, SetName: "Black Dragon Mail"}
	}
	if got := len(c.GetActiveSetBonuses()); got != 2 {
		t.Fatalf("three equipped pieces must activate exactly two bonuses, got %d", got)
	}
	two := foreverEquipmentSets[489].Bonuses[2]
	if two.Stats[stats.MeleeHit] != 1 || two.Stats[stats.SpellHit] != 0 {
		t.Fatal("post-equipment set hit must not be mirrored into spell hit")
	}
	if ForeverEquipmentBonusSupported(489, 5) {
		t.Fatal("nonexistent bonus reported as supported")
	}
}

func TestForeverOrdinarySetProfession(t *testing.T) {
	c := &Character{Unit: Unit{Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}}}
	for slot := 0; slot < 3; slot++ {
		c.Equipment[slot] = Item{SetID: 421, SetName: "Bloodvine Garb"}
	}
	if len(c.GetActiveSetBonuses()) != 0 {
		t.Fatal("Bloodvine set bonus requires Tailoring")
	}
	c.professions[1] = proto.Profession_Tailoring
	if len(c.GetActiveSetBonuses()) != 1 {
		t.Fatal("Tailoring did not enable the equipped Bloodvine bonus")
	}
}

func TestForeverTierSetNotAppliedByOrdinarySetAdapter(t *testing.T) {
	c := &Character{Unit: Unit{Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}}}
	for slot := 0; slot < 4; slot++ {
		c.Equipment[slot] = Item{SetID: 2111}
	}
	if overridesForeverEquipmentSet(2111) || len(c.foreverEquippedBonuses()) != 0 {
		t.Fatal("ordinary equipment adapter must not duplicate the role-specific Tier 1 implementation")
	}
}

func TestForeverTankTierSetMappings(t *testing.T) {
	for _, tc := range []struct {
		spec proto.Spec
		id   int32
	}{
		{proto.Spec_SpecTankWarrior, 2103},
		{proto.Spec_SpecProtectionPaladin, 2108},
		{proto.Spec_SpecFeralTankDruid, 2113},
	} {
		if got := ForeverTier1SetID(tc.spec); got != tc.id {
			t.Fatalf("%v mapped to Tier 1 set %d, want %d", tc.spec, got, tc.id)
		}
		c := &Character{Unit: Unit{Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}}, Spec: tc.spec, ForeverTier1Bonuses: true}
		if got := c.forcedForeverTier1SetID(); got != tc.id {
			t.Fatalf("%v forced set %d, want %d", tc.spec, got, tc.id)
		}
		for slot := 0; slot < 5; slot++ {
			c.Equipment[slot] = Item{SetID: tc.id}
		}
		if overridesForeverEquipmentSet(tc.id) || len(c.foreverEquippedBonuses()) != 0 {
			t.Fatalf("tank Tier 1 %d must not be applied by the ordinary equipment adapter", tc.id)
		}
	}
}
