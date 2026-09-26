package main

import (
	"fmt"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func classPreferredArmor(class proto.Class) proto.ArmorType {
	return core.ClassMaxArmorType(class)
}

func validateEquipmentSlot(p *proto.Player, slot int, item core.Item) error {
	slotTypes := []proto.ItemType{
		proto.ItemType_ItemTypeHead, proto.ItemType_ItemTypeNeck, proto.ItemType_ItemTypeShoulder,
		proto.ItemType_ItemTypeBack, proto.ItemType_ItemTypeChest, proto.ItemType_ItemTypeWrist,
		proto.ItemType_ItemTypeHands, proto.ItemType_ItemTypeWaist, proto.ItemType_ItemTypeLegs,
		proto.ItemType_ItemTypeFeet, proto.ItemType_ItemTypeFinger, proto.ItemType_ItemTypeFinger,
		proto.ItemType_ItemTypeTrinket, proto.ItemType_ItemTypeTrinket, proto.ItemType_ItemTypeWeapon,
		proto.ItemType_ItemTypeWeapon, proto.ItemType_ItemTypeRanged,
	}
	if slot >= len(slotTypes) || item.Type != slotTypes[slot] {
		return fmt.Errorf("%s does not fit slot %d", item.Name, slot)
	}
	if item.Type != proto.ItemType_ItemTypeWeapon {
		return nil
	}
	if slot == int(proto.ItemSlot_ItemSlotMainHand) && item.HandType == proto.HandType_HandTypeOffHand {
		return fmt.Errorf("%s is off-hand only", item.Name)
	}
	if slot == int(proto.ItemSlot_ItemSlotOffHand) {
		if item.HandType != proto.HandType_HandTypeOffHand && item.HandType != proto.HandType_HandTypeOneHand {
			return fmt.Errorf("%s cannot be equipped in the off hand", item.Name)
		}
		if core.ItemsByID[p.Equipment.Items[proto.ItemSlot_ItemSlotMainHand].GetId()].HandType == proto.HandType_HandTypeTwoHand {
			return fmt.Errorf("an off-hand item cannot accompany a two-handed weapon")
		}
		if item.WeaponType != proto.WeaponType_WeaponTypeOffHand && item.WeaponType != proto.WeaponType_WeaponTypeShield {
			if !core.ClassCanDualWield(p.Class) {
				return fmt.Errorf("%s cannot dual wield in this build", p.Class)
			}
		}
	}
	return nil
}
