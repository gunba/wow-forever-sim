package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestForeverEquipmentAndBonusStats(t *testing.T) {
	c := &Character{Unit: Unit{Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}}}
	for i := range c.itemStatMultipliers {
		c.itemStatMultipliers[i] = 1
	}
	c.Equipment[proto.ItemSlot_ItemSlotMainHand] = Item{
		Stats: stats.Stats{stats.MeleeCrit: 1, stats.MeleeHit: 1},
	}
	c.bonusStats = stats.Stats{stats.MeleeHit: 8, stats.SpellHit: 15, stats.SpellCrit: 2}
	got := c.EquipStats()
	for stat, want := range map[stats.Stat]float64{
		stats.MeleeHit: 9, stats.SpellHit: 16,
		stats.MeleeCrit: 1, stats.SpellCrit: 3,
	} {
		if got[stat] != want {
			t.Errorf("stat %v = %v, want %v", stat, got[stat], want)
		}
	}
}

func TestForeverExplicitHealingDamage(t *testing.T) {
	c := &Character{Unit: Unit{Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}}}
	mace := Item{Stats: stats.Stats{stats.HealingPower: 176, stats.SpellDamage: 58}}
	if got := c.itemStats(mace, true)[stats.SpellDamage]; got != 58 {
		t.Fatalf("Battle Mace spell damage = %v, want 58", got)
	}
	mace.Enchant.Stats = stats.Stats{stats.HealingPower: 30}
	mace.RandomSuffix.Stats = stats.Stats{stats.HealingPower: 24}
	if got := c.itemStats(mace, true)[stats.SpellDamage]; got != 58 {
		t.Fatalf("healing-only components added damage: %v, want 58", got)
	}
	mace.Enchant.Stats = stats.Stats{stats.HealingPower: 35, stats.SpellDamage: 12}
	if got := c.itemStats(mace, true)[stats.SpellDamage]; got != 70 {
		t.Fatalf("explicit enchant damage = %v, want 70", got)
	}
	healingOnly := Item{Stats: stats.Stats{stats.HealingPower: 42}}
	if got := c.itemStats(healingOnly, true)[stats.SpellDamage]; got != 0 {
		t.Fatalf("healing-only item added %v spell damage", got)
	}
}

func TestClassicEquipmentKeepsSeparateStats(t *testing.T) {
	c := &Character{Unit: Unit{Env: &Environment{}}}
	item := Item{Stats: stats.Stats{stats.MeleeCrit: 1, stats.HealingPower: 30}}
	got := c.itemStats(item, true)
	if got[stats.SpellCrit] != 0 || got[stats.SpellDamage] != 0 || got[stats.MeleeCrit] != 1 {
		t.Fatalf("Classic item was transformed: %v", got)
	}
}
