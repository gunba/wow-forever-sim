package core

import (
	"fmt"

	"github.com/wowsims/classic/sim/core/proto"
)

// ClassMaxArmorType is the class's armor proficiency ceiling, not a
// recommendation to wear that material in every slot.
func ClassMaxArmorType(class proto.Class) proto.ArmorType {
	switch class {
	case proto.Class_ClassMage, proto.Class_ClassPriest, proto.Class_ClassWarlock:
		return proto.ArmorType_ArmorTypeCloth
	case proto.Class_ClassDruid, proto.Class_ClassRogue:
		return proto.ArmorType_ArmorTypeLeather
	case proto.Class_ClassHunter, proto.Class_ClassShaman:
		return proto.ArmorType_ArmorTypeMail
	case proto.Class_ClassPaladin, proto.Class_ClassWarrior:
		return proto.ArmorType_ArmorTypePlate
	default:
		return proto.ArmorType_ArmorTypeUnknown
	}
}

func ValidateEquipmentArmor(class proto.Class, equipment Equipment) error {
	for slot, item := range equipment {
		if item.ID != 0 && item.ArmorType > ClassMaxArmorType(class) {
			return fmt.Errorf("%s cannot wear %s (%s) in %s", class, item.Name, item.ArmorType, proto.ItemSlot(slot))
		}
	}
	return nil
}
