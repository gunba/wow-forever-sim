package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestWeaponLayoutProficiencies(t *testing.T) {
	oneHand := Item{ID: 1, HandType: proto.HandType_HandTypeOneHand, WeaponType: proto.WeaponType_WeaponTypeMace}
	twoHand := Item{ID: 2, HandType: proto.HandType_HandTypeTwoHand, WeaponType: proto.WeaponType_WeaponTypeAxe}
	shield := Item{ID: 3, HandType: proto.HandType_HandTypeOffHand, WeaponType: proto.WeaponType_WeaponTypeShield}
	focus := Item{ID: 4, HandType: proto.HandType_HandTypeOffHand, WeaponType: proto.WeaponType_WeaponTypeOffHand}
	for _, class := range []proto.Class{proto.Class_ClassHunter, proto.Class_ClassRogue, proto.Class_ClassWarrior, proto.Class_ClassShaman} {
		err := ValidateWeaponLayout(class, oneHand, oneHand)
		if (err == nil) != ClassCanDualWield(class) {
			t.Errorf("%s dual-wield validation: %v", class, err)
		}
	}
	for _, offHand := range []Item{shield, focus, {}} {
		if err := ValidateWeaponLayout(proto.Class_ClassShaman, oneHand, offHand); err != nil {
			t.Fatal(err)
		}
	}
	if err := ValidateWeaponLayout(proto.Class_ClassShaman, twoHand, Item{}); err != nil {
		t.Fatal(err)
	}
	if ValidateWeaponLayout(proto.Class_ClassShaman, twoHand, shield) == nil {
		t.Fatal("accepted a two-handed weapon with a shield")
	}
}
