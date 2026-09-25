package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"sync"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

type gearRecord struct {
	ID, ItemLevel, RequiredLevel, RequiredSkill, RequiredSkillRank, RequiredAbility int32
	MaxCount, LimitCategory, LimitCategoryQuantity, SetID                           int32
	UniqueEquipped, ClientStatRecord                                                bool
	Effects                                                                         []json.RawMessage
	SetBonuses                                                                      []struct{ Pieces, SpellID int32 }
	Sources                                                                         []struct{ Kind, Faction string }
	Synthetic, BenchmarkEligible                                                    bool
	ModelVersion                                                                    int32
}

var readGearCatalog = sync.OnceValue(func() map[int32]gearRecord {
	var data struct{ Items []gearRecord }
	if err := json.Unmarshal(mustRead("assets/db_inputs/forever_gear_catalog.json"), &data); err != nil {
		panic(err)
	}
	out := map[int32]gearRecord{}
	for _, item := range data.Items {
		out[item.ID] = item
	}
	for _, path := range []string{
		"assets/db_inputs/forever_synthetic_gear.json",
		"assets/db_inputs/forever_synthetic_gear_v2.json",
	} {
		if err := json.Unmarshal(mustRead(path), &data); err != nil {
			panic(err)
		}
		for _, item := range data.Items {
			if _, exists := out[item.ID]; exists {
				panic(fmt.Sprintf("synthetic item %d collides with the real catalog", item.ID))
			}
			out[item.ID] = item
		}
	}
	return out
})

func fullyModeledV2(p *proto.Player) bool {
	catalog := readGearCatalog()
	count := 0
	for _, equipped := range p.GetEquipment().GetItems() {
		if equipped.GetId() == 0 {
			continue
		}
		count++
		if catalog[equipped.Id].ModelVersion != 2 {
			return false
		}
	}
	return count >= 16
}

var readGearReviews = sync.OnceValue(func() map[int32]string {
	var data struct {
		Items map[int32]struct{ Status string }
	}
	if err := json.Unmarshal(mustRead("assets/db_inputs/forever_gear_reviews.json"), &data); err != nil {
		panic(err)
	}
	out := map[int32]string{}
	for id, review := range data.Items {
		out[id] = review.Status
	}
	return out
})

func gearReviewError(item core.Item) error {
	record, ok := readGearCatalog()[item.ID]
	if !ok {
		return fmt.Errorf("%s is outside the current equipment catalog", item.Name)
	}
	if record.Synthetic {
		if !record.BenchmarkEligible {
			return fmt.Errorf("%s is a trinket sensitivity case, not benchmark gear", item.Name)
		}
		if len(record.Effects) != 0 || record.SetID != 0 {
			return fmt.Errorf("%s incorrectly includes a modeled effect or set", item.Name)
		}
		return nil
	}
	review := readGearReviews()[item.ID]
	if review == "implemented" || review == "verified" {
		return nil
	}
	if !record.ClientStatRecord || len(record.Effects) != 0 ||
		core.HasItemEffect(item.ID) || core.HasWeaponEffect(item.ID) {
		return fmt.Errorf("%s needs an item-effect review", item.Name)
	}
	return nil
}

func validateGear(p *proto.Player) error {
	if err := validateGearLayout(p, true); err != nil {
		return err
	}
	return validateGearEnchants(p)
}

func validateGearLayout(p *proto.Player, complete bool) error {
	catalog := readGearCatalog()
	counts, groups, sets := map[int32]int32{}, map[int32]int32{}, map[int32]int32{}
	if p.Equipment == nil || len(p.Equipment.Items) != 17 {
		return fmt.Errorf("a complete level-60 equipment layout is required")
	}
	for slot, spec := range p.Equipment.Items {
		id := spec.GetId()
		if id == 0 {
			if !complete {
				continue
			}
			if slot == int(proto.ItemSlot_ItemSlotOffHand) &&
				core.ItemsByID[p.Equipment.Items[proto.ItemSlot_ItemSlotMainHand].GetId()].HandType == proto.HandType_HandTypeTwoHand {
				continue
			}
			return fmt.Errorf("empty equipment slot %d", slot)
		}
		item, exists := core.ItemsByID[id]
		if !exists {
			return fmt.Errorf("missing database item %d", id)
		}
		if err := gearReviewError(item); err != nil {
			return err
		}
		record := catalog[id]
		available := false
		for _, source := range record.Sources {
			available = available || source.Faction == "" || source.Faction == raceFaction(p.Race)
		}
		if !available {
			return fmt.Errorf("%s is not available from this faction's sources", item.Name)
		}
		if record.RequiredLevel > 60 || record.RequiredSkillRank > 300 || record.RequiredAbility != 0 {
			return fmt.Errorf("%s has an unsupported level/skill/ability requirement", item.Name)
		}
		if record.RequiredSkill != 0 {
			profession := map[int32]proto.Profession{
				171: proto.Profession_Alchemy, 164: proto.Profession_Blacksmithing,
				333: proto.Profession_Enchanting, 202: proto.Profession_Engineering,
				165: proto.Profession_Leatherworking, 197: proto.Profession_Tailoring,
			}[record.RequiredSkill]
			if profession == proto.Profession_ProfessionUnknown ||
				(profession != p.Profession1 && profession != p.Profession2) {
				return fmt.Errorf("%s requires profession skill %d", item.Name, record.RequiredSkill)
			}
		}
		if len(item.ClassAllowlist) != 0 && !slices.Contains(item.ClassAllowlist, p.Class) {
			return fmt.Errorf("%s cannot equip %s", p.Class, item.Name)
		}
		if err := validateEquipmentSlot(p, slot, item); err != nil {
			return err
		}
		if !classCanEquip(p.Class, item) {
			return fmt.Errorf("%s lacks proficiency for %s", p.Class, item.Name)
		}
		counts[id]++
		if (record.UniqueEquipped || record.MaxCount == 1) && counts[id] > 1 {
			return fmt.Errorf("%s is unique-equipped", item.Name)
		}
		if record.LimitCategory != 0 {
			groups[record.LimitCategory]++
			if groups[record.LimitCategory] > record.LimitCategoryQuantity {
				return fmt.Errorf("%s exceeds shared equipment category %d", item.Name, record.LimitCategory)
			}
		}
		if record.SetID != 0 {
			sets[record.SetID]++
		}
	}
	// Unimplemented current bonuses are reported separately; they must never
	// activate an obsolete Classic implementation or prevent legal equipment.
	for _, spec := range p.Equipment.Items {
		record := catalog[spec.GetId()]
		for _, bonus := range record.SetBonuses {
			if sets[record.SetID] >= bonus.Pieces && !core.HasForeverEquipmentSetDefinition(record.SetID) {
				return fmt.Errorf("set %d lacks a Forever equipment definition", record.SetID)
			}
		}
	}
	return nil
}

func unmodeledSetBonuses(p *proto.Player) []string {
	catalog := readGearCatalog()
	counts := map[int32]int32{}
	records := map[int32]gearRecord{}
	for _, spec := range p.Equipment.Items {
		record := catalog[spec.GetId()]
		if record.SetID != 0 {
			counts[record.SetID]++
			records[record.SetID] = record
		}
	}
	var missing []string
	for id, count := range counts {
		for _, bonus := range records[id].SetBonuses {
			if bonus.Pieces <= count && !core.ForeverEquipmentBonusSupported(id, bonus.Pieces) {
				missing = append(missing, fmt.Sprintf("Set %d, %d pieces, spell %d", id, bonus.Pieces, bonus.SpellID))
			}
		}
	}
	slices.Sort(missing)
	return missing
}

// The current UI's ordinary proficiencies. New Forever extensions are not
// inferred from an item's otherwise unrestricted class mask.
func classCanEquip(class proto.Class, item core.Item) bool {
	if item.ArmorType > classPreferredArmor(class) {
		return false
	}
	if item.Type == proto.ItemType_ItemTypeRanged {
		allowed := map[proto.Class][]proto.RangedWeaponType{
			proto.Class_ClassDruid: {4}, proto.Class_ClassHunter: {1, 2, 3},
			proto.Class_ClassMage: {8}, proto.Class_ClassPaladin: {5}, proto.Class_ClassPriest: {8},
			proto.Class_ClassRogue: {1, 2, 3, 6}, proto.Class_ClassShaman: {7},
			proto.Class_ClassWarlock: {8}, proto.Class_ClassWarrior: {1, 2, 3, 6},
		}
		return slices.Contains(allowed[class], item.RangedWeaponType)
	}
	if item.Type != proto.ItemType_ItemTypeWeapon {
		return true
	}
	allowed := map[proto.Class][]proto.WeaponType{
		proto.Class_ClassDruid: {2, 3, 4, 5, 8}, proto.Class_ClassHunter: {1, 2, 3, 5, 6, 8, 9},
		proto.Class_ClassMage: {2, 5, 8, 9}, proto.Class_ClassPaladin: {1, 4, 5, 6, 7, 9},
		proto.Class_ClassPriest: {2, 4, 5, 8}, proto.Class_ClassRogue: {2, 3, 4, 5, 9},
		proto.Class_ClassShaman:  {1, 2, 3, 4, 5, 7, 8},
		proto.Class_ClassWarlock: {2, 5, 8, 9}, proto.Class_ClassWarrior: {1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	if !slices.Contains(allowed[class], item.WeaponType) {
		return false
	}
	if item.HandType == proto.HandType_HandTypeTwoHand {
		if class == proto.Class_ClassRogue {
			return false
		}
		if class == proto.Class_ClassMage || class == proto.Class_ClassPriest || class == proto.Class_ClassWarlock {
			return item.WeaponType == proto.WeaponType_WeaponTypeStaff
		}
	}
	return true
}
