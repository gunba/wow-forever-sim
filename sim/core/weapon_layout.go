package core

import (
	"fmt"

	"github.com/wowsims/classic/sim/core/proto"
)

func ClassCanDualWield(class proto.Class) bool {
	return class == proto.Class_ClassHunter || class == proto.Class_ClassRogue || class == proto.Class_ClassWarrior
}

func ValidateWeaponLayout(class proto.Class, mainHand, offHand Item) error {
	if offHand.ID == 0 {
		return nil
	}
	if mainHand.HandType == proto.HandType_HandTypeTwoHand {
		return fmt.Errorf("a two-handed weapon cannot be equipped with an off-hand item")
	}
	if offHand.WeaponType != proto.WeaponType_WeaponTypeShield &&
		offHand.WeaponType != proto.WeaponType_WeaponTypeOffHand && !ClassCanDualWield(class) {
		return fmt.Errorf("%s cannot dual wield; use a two-handed weapon or a one-handed weapon with a shield or held off-hand item", class)
	}
	return nil
}
