package core

import (
	"fmt"
	"slices"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

type OnSwapItem func(*Simulation)

const offset = proto.ItemSlot_ItemSlotMainHand

type ItemSwap struct {
	character       *Character
	onSwapCallbacks []OnSwapItem

	// Which slots to actually swap.
	slots []proto.ItemSlot

	// Holds items that are currently not equipped
	unEquippedItems [3]Item
	initialItems    [3]Item
}

// TODO All the extra parameters here and the code in multiple places for handling the Weapon struct is really messy,
//  we'll need to figure out something cleaner as this will be quite error-prone

func (character *Character) enableItemSwap(itemSwap *proto.ItemSwap) {
	var slots []proto.ItemSlot
	hasMhSwap := itemSwap.MhItem != nil && itemSwap.MhItem.Id != 0
	hasOhSwap := itemSwap.OhItem != nil && itemSwap.OhItem.Id != 0
	hasRangedSwap := itemSwap.RangedItem != nil && itemSwap.RangedItem.Id != 0

	mainItems := [3]Item{
		character.Equipment[proto.ItemSlot_ItemSlotMainHand],
		character.Equipment[proto.ItemSlot_ItemSlotOffHand],
		character.Equipment[proto.ItemSlot_ItemSlotRanged],
	}
	swapItems := [3]Item{
		toItem(itemSwap.MhItem),
		toItem(itemSwap.OhItem),
		toItem(itemSwap.RangedItem),
	}
	if !hasMhSwap {
		swapItems[0] = mainItems[0]
	}
	if !hasRangedSwap {
		swapItems[2] = mainItems[2]
	}
	if err := ValidateWeaponLayout(character.Class, swapItems[0], swapItems[1]); err != nil {
		panic(err)
	}

	// Handle MH and OH together, because present MH + empty OH --> swap MH and unequip OH
	if hasMhSwap || hasOhSwap {
		if !sameSwapItem(swapItems[0], mainItems[0]) {
			slots = append(slots, proto.ItemSlot_ItemSlotMainHand)
		}
		if !sameSwapItem(swapItems[1], mainItems[1]) {
			slots = append(slots, proto.ItemSlot_ItemSlotOffHand)
		}
	}
	if hasRangedSwap {
		if !sameSwapItem(swapItems[2], mainItems[2]) {
			slots = append(slots, proto.ItemSlot_ItemSlotRanged)
		}
	}

	if len(slots) == 0 {
		return
	}

	character.ItemSwap = ItemSwap{
		character:       character,
		slots:           slots,
		unEquippedItems: swapItems,
		initialItems:    mainItems,
	}
}

func (swap *ItemSwap) initialize(character *Character) {
	swap.character = character
}

func (character *Character) RegisterOnItemSwap(callback OnSwapItem) {
	if character == nil || !character.ItemSwap.IsEnabled() {
		return
	}

	character.ItemSwap.onSwapCallbacks = append(character.ItemSwap.onSwapCallbacks, callback)
}

// Helper for handling Effects that use PPMManager to toggle the aura on/off
func (swap *ItemSwap) RegisterOnSwapItemForEffectWithPPMManager(effectID int32, ppm float64, ppmm *PPMManager, aura *Aura) {
	character := swap.character
	character.RegisterOnItemSwap(func(sim *Simulation) {
		procMask := character.GetProcMaskForEnchant(effectID)
		*ppmm = character.AutoAttacks.NewPPMManager(ppm, procMask)

		if ppmm.Chance(procMask) == 0 {
			aura.Deactivate(sim)
		} else {
			aura.Activate(sim)
		}
	})

}

// Helper for handling Effects that use the effectID to toggle the aura on and off
func (swap *ItemSwap) RegisterOnSwapItemForEffect(effectID int32, aura *Aura) {
	character := swap.character
	character.RegisterOnItemSwap(func(sim *Simulation) {
		procMask := character.GetProcMaskForEnchant(effectID)

		if procMask == ProcMaskUnknown {
			aura.Deactivate(sim)
		} else {
			aura.Activate(sim)
		}
	})
}

func (swap *ItemSwap) IsEnabled() bool {
	return swap.character != nil && len(swap.slots) > 0
}

func (swap *ItemSwap) IsSwapped() bool {
	if !swap.IsEnabled() {
		return false
	}
	for _, slot := range swap.slots {
		if !sameSwapItem(swap.character.Equipment[slot], swap.initialItems[slot-offset]) {
			return true
		}
	}
	return false
}

func (swap *ItemSwap) GetItem(slot proto.ItemSlot) *Item {
	if slot-offset < 0 {
		panic("Not able to swap Item " + slot.String() + " not supported")
	}
	return &swap.unEquippedItems[slot-offset]
}

// Compare the complete per-instance identity, not just the item ID.
func sameSwapItem(a, b Item) bool {
	return a.ID == b.ID && a.Enchant.EffectID == b.Enchant.EffectID && a.RandomSuffix.ID == b.RandomSuffix.ID
}

func ValidateEquipmentUnique(equipment Equipment) error {
	counts := map[int32]int{}
	for _, item := range equipment {
		if item.ID == 0 {
			continue
		}
		counts[item.ID]++
		if item.Unique && counts[item.ID] > 1 {
			return fmt.Errorf("%s is unique-equipped", item.Name)
		}
	}
	return nil
}

// Compute and validate the entire layout before changing equipment or stats.
// A requested MH change to a two-hander also unequips the off hand.
func (swap *ItemSwap) plannedEquipment(slots []proto.ItemSlot) Equipment {
	next := swap.character.Equipment
	mainRequested := false
	for _, slot := range slots {
		if slot < offset || slot > proto.ItemSlot_ItemSlotRanged {
			panic("unsupported weapon swap slot")
		}
		if !slices.Contains(swap.slots, slot) {
			continue
		}
		next[slot] = *swap.GetItem(slot)
		mainRequested = mainRequested || slot == proto.ItemSlot_ItemSlotMainHand
	}
	if mainRequested && next[proto.ItemSlot_ItemSlotMainHand].HandType == proto.HandType_HandTypeTwoHand {
		next[proto.ItemSlot_ItemSlotOffHand] = Item{}
	}
	if err := ValidateWeaponLayout(swap.character.Class, next[proto.ItemSlot_ItemSlotMainHand], next[proto.ItemSlot_ItemSlotOffHand]); err != nil {
		panic(err)
	}
	if err := ValidateEquipmentUnique(next); err != nil {
		panic(err)
	}
	for slot, item := range next {
		if err := ValidateEnchantRequirements(swap.character.Level, proto.ItemSlot(slot), item); err != nil {
			panic(err)
		}
	}
	return next
}

func (swap *ItemSwap) CalcStatChanges(slots []proto.ItemSlot) stats.Stats {
	next := swap.plannedEquipment(slots)
	delta := stats.Stats{}
	for slot := offset; slot <= proto.ItemSlot_ItemSlotRanged; slot++ {
		delta = delta.Add(swap.getItemStats(next[slot]).Subtract(swap.getItemStats(swap.character.Equipment[slot])))
	}
	return delta
}

func (swap *ItemSwap) SwapItems(sim *Simulation, slots []proto.ItemSlot) {
	if !swap.IsEnabled() {
		return
	}
	character := swap.character
	next := swap.plannedEquipment(slots)
	newStats := stats.Stats{}
	var changed []proto.ItemSlot
	meleeWeaponSwapped := false
	for slot := offset; slot <= proto.ItemSlot_ItemSlotRanged; slot++ {
		old := character.Equipment[slot]
		if sameSwapItem(old, next[slot]) {
			continue
		}
		newStats = newStats.Add(swap.getItemStats(next[slot]).Subtract(swap.getItemStats(old)))
		swap.unEquippedItems[slot-offset] = old
		character.Equipment[slot] = next[slot]
		changed = append(changed, slot)
		meleeWeaponSwapped = meleeWeaponSwapped || slot != proto.ItemSlot_ItemSlotRanged
	}
	if len(changed) == 0 {
		return
	}
	character.AddStatsDynamic(sim, newStats)
	for _, slot := range changed {
		swap.swapWeapon(slot)
	}
	if sim.Log != nil {
		sim.Log("Item Swap Stats: %v", newStats)
	}
	for _, onSwap := range swap.onSwapCallbacks {
		onSwap(sim)
	}
	if character.AutoAttacks.AutoSwingMelee && meleeWeaponSwapped && sim.CurrentTime > 0 {
		character.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime, false)
	}
	if character.GCD.IsReady(sim) {
		character.SetGCDTimer(sim, sim.CurrentTime+1500*time.Millisecond)
	}
}

// Iteration reset already restores the initial stats. Restore equipment/cache
// first, without applying another stat delta or activating auras before reset.
func (swap *ItemSwap) restoreInitialEquipment() bool {
	if !swap.IsEnabled() || !swap.IsSwapped() {
		return false
	}
	var changed []proto.ItemSlot
	for _, slot := range swap.slots {
		if sameSwapItem(swap.character.Equipment[slot], swap.initialItems[slot-offset]) {
			continue
		}
		swap.unEquippedItems[slot-offset] = swap.character.Equipment[slot]
		swap.character.Equipment[slot] = swap.initialItems[slot-offset]
		changed = append(changed, slot)
	}
	for _, slot := range changed {
		swap.swapWeapon(slot)
	}
	return len(changed) != 0
}

func (swap *ItemSwap) getItemStats(item Item) stats.Stats {
	return swap.character.itemStats(item, true).DotProduct(swap.character.itemStatMultipliers)
}

func (swap *ItemSwap) swapWeapon(slot proto.ItemSlot) {
	character := swap.character

	switch slot {
	case proto.ItemSlot_ItemSlotMainHand:
		if character.AutoAttacks.AutoSwingMelee {
			character.AutoAttacks.SetMH(character.WeaponFromMainHand())
		}
	case proto.ItemSlot_ItemSlotOffHand:
		if character.AutoAttacks.AutoSwingMelee {
			weapon := character.WeaponFromOffHand()
			character.AutoAttacks.SetOH(weapon)

			character.AutoAttacks.IsDualWielding = weapon.SwingSpeed != 0
			character.PseudoStats.CanBlock = character.OffHand().WeaponType == proto.WeaponType_WeaponTypeShield
		}
	case proto.ItemSlot_ItemSlotRanged:
		if character.AutoAttacks.AutoSwingRanged {
			character.AutoAttacks.SetRanged(character.WeaponFromRanged())
		}
	}
}

func (swap *ItemSwap) reset(sim *Simulation) {
	if !swap.IsEnabled() || !swap.IsSwapped() {
		return
	}

	var changed []proto.ItemSlot
	for _, slot := range swap.slots {
		if !sameSwapItem(swap.character.Equipment[slot], swap.initialItems[slot-offset]) {
			changed = append(changed, slot)
		}
	}
	swap.SwapItems(sim, changed)
}

func getInitialEquippedItems(character *Character) [3]Item {
	var items [3]Item

	for i := range items {
		items[i] = character.Equipment[i+int(offset)]
	}

	return items
}

func toItem(itemSpec *proto.ItemSpec) Item {
	if itemSpec == nil {
		return Item{}
	}

	return NewItem(ItemSpec{
		ID: itemSpec.Id,

		Enchant:      itemSpec.Enchant,
		RandomSuffix: itemSpec.RandomSuffix,
	})
}
