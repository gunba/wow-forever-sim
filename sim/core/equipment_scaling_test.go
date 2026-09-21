package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/stats"
)

func TestEquipmentScaling(t *testing.T) {
	original := Equipment{Item{
		ID: 1, Stats: stats.Stats{stats.Strength: 40, stats.MeleeHit: 1, stats.MeleeCrit: 2},
		RandomSuffix:    RandomSuffix{ID: 2, Stats: stats.Stats{stats.Agility: 20}},
		Enchant:         Enchant{EffectID: 3, Stats: stats.Stats{stats.Strength: 10}},
		WeaponDamageMin: 100, WeaponDamageMax: 200, SwingSpeed: 2,
		BonusPhysicalDamage: 5,
	}}
	scaled := original.Scaled(1.2)
	if scaled[0].Stats[stats.Strength] != 48 || scaled[0].Stats[stats.MeleeHit] != 1.2 ||
		scaled[0].RandomSuffix.Stats[stats.Agility] != 24 ||
		scaled[0].WeaponDamageMin != 120 || scaled[0].WeaponDamageMax != 240 ||
		scaled[0].BonusPhysicalDamage != 6 {
		t.Fatalf("wrong scaled item: %+v", scaled[0])
	}
	if scaled[0].ID != original[0].ID || scaled[0].Enchant != original[0].Enchant ||
		scaled[0].SwingSpeed != original[0].SwingSpeed || scaled[0].WeaponSkills != original[0].WeaponSkills {
		t.Fatal("scaling changed an item identity, enchant, speed or weapon skill")
	}
	if original[0].Stats[stats.Strength] != 40 || original[0].WeaponDamageMin != 100 {
		t.Fatal("scaled equipment mutated the source")
	}
	if original.Scaled(0)[0].Stats != original[0].Stats || original.Scaled(1)[0].Stats != original[0].Stats {
		t.Fatal("default scale changed equipment")
	}
}
