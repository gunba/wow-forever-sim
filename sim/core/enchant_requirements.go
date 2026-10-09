package core

import (
	"fmt"
	"slices"

	"github.com/wowsims/classic/sim/core/proto"
)

// ValidateEnchantRequirements covers the sourced Forever kits/enchant added in
// build 70291. Crafting skill belongs to the recipe, not to the wearer.
func ValidateEnchantRequirements(level int32, slot proto.ItemSlot, item Item) error {
	e := item.Enchant
	switch e.EffectID {
	case 8483, 8486, 8488, 8491:
		if !slices.Contains([]proto.ItemSlot{proto.ItemSlot_ItemSlotChest, proto.ItemSlot_ItemSlotLegs, proto.ItemSlot_ItemSlotHands, proto.ItemSlot_ItemSlotFeet}, slot) {
			return fmt.Errorf("armor kit %d cannot be applied to slot %s", e.EffectID, slot)
		}
		if item.ArmorType < proto.ArmorType_ArmorTypeCloth || item.ArmorType > proto.ArmorType_ArmorTypePlate || e.ArmorTypeMask&(1<<uint32(item.ArmorType)) == 0 {
			return fmt.Errorf("armor kit %d requires cloth, leather, mail or plate", e.EffectID)
		}
	case 8721, 8217:
		if !item.IsWeapon() || (slot != proto.ItemSlot_ItemSlotMainHand && slot != proto.ItemSlot_ItemSlotOffHand) {
			return fmt.Errorf("enchant %d requires a melee weapon", e.EffectID)
		}
	default:
		return nil
	}
	if level < e.RequiredLevel {
		return fmt.Errorf("enchant %d requires level %d", e.EffectID, e.RequiredLevel)
	}
	if item.ItemLevel < e.ItemLevelMin {
		return fmt.Errorf("enchant %d requires item level %d", e.EffectID, e.ItemLevelMin)
	}
	return nil
}
