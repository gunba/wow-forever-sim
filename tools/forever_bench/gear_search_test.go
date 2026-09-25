//go:build with_db

package main

import (
	"encoding/json"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestGearSearchFullPoolAndWeaponLegality(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "enhancement" && b.Key != "mutilate" {
			continue
		}
		p := b.player(b.races()[0])
		pool, excluded := comparisonGearPool(b, p)
		var exported struct{ Items []struct{ ID int32 } }
		if err := json.Unmarshal(mustRead("assets/db_inputs/forever_ilvl65_items.json"), &exported); err != nil {
			t.Fatal(err)
		}
		accounted := map[int32]bool{}
		for _, item := range pool {
			accounted[item.ID] = true
		}
		for id := range excluded {
			accounted[id] = true
		}
		for _, row := range exported.Items {
			if !accounted[row.ID] {
				t.Fatalf("missing exported item %d", row.ID)
			}
		}
		foundLastPageCraft := false
		for _, item := range pool {
			foundLastPageCraft = foundLastPageCraft || item.ID == 279253 || item.ID == 279260
		}
		if !foundLastPageCraft {
			t.Fatal("items beyond the first rendered page are missing")
		}
		candidates := gearCandidates(b, p, 14, pool)
		if len(candidates) == 0 {
			t.Fatalf("%s has no weapon candidates", b.Key)
		}
		for _, candidate := range candidates {
			if err := validateGear(candidate); err != nil {
				t.Fatal(err)
			}
			if candidate.Class == proto.Class_ClassShaman && candidate.Consumes.GetOffHandImbue() != 0 {
				t.Fatal("Shaman gained an off-hand imbue")
			}
		}
	}
}

func TestGearSearchCanMoveReplacedTrinketToSecondSlot(t *testing.T) {
	b := builds()[0]
	p := b.player(b.races()[0])
	p.Equipment.Items[12] = &proto.ItemSpec{Id: 272438}
	p.Equipment.Items[13] = &proto.ItemSpec{Id: 249470}
	pool, _ := comparisonGearPool(b, p)
	for _, candidate := range gearCandidates(b, p, 13, pool) {
		if candidate.Equipment.Items[12].Id == 272438 && candidate.Equipment.Items[13].Id == 249469 {
			return
		}
	}
	t.Fatal("lower-level Frozen Heart vanished after replacing the first trinket")
}

func TestModeledGearPoolKeepsTrinketSensitivitySeparate(t *testing.T) {
	*modeledOnly = true
	defer func() { *modeledOnly = false }()
	b := builds()[0]
	p := b.player(b.races()[0])
	pool, excluded := comparisonGearPool(b, p)
	modeled, passive := 0, 0
	for _, item := range pool {
		if readGearCatalog()[item.ID].ModelVersion != 2 {
			t.Fatalf("noncurrent or real item entered the modeled-only pool: %d", item.ID)
		}
		if item.ArmorType != 0 && item.Type != proto.ItemType_ItemTypeBack &&
			item.ArmorType != classPreferredArmor(p.Class) {
			t.Fatalf("modeled body item %d is the wrong class armor type", item.ID)
		}
		modeled++
		if readGearCatalog()[item.ID].BenchmarkEligible && item.Type == proto.ItemType_ItemTypeTrinket {
			passive++
		}
	}
	if modeled == 0 || passive != 11 {
		t.Fatalf("expected the modeled pool and eleven conservative trinkets; got %d/%d", modeled, passive)
	}
	rejected := 0
	for id, reason := range excluded {
		if readGearCatalog()[id].ModelVersion == 2 && !readGearCatalog()[id].BenchmarkEligible {
			if reason != "full-capacity passive trinket is a sensitivity case, not ranking gear" {
				t.Fatalf("unclear rejection reason for modeled item %d: %s", id, reason)
			}
			rejected++
		}
	}
	if rejected != 8 {
		t.Fatalf("expected eight full-capacity trinkets outside rankings, got %d", rejected)
	}
}

func TestModeledSeedReplacesEveryEquippedRealItem(t *testing.T) {
	*modeledOnly = true
	defer func() { *modeledOnly = false }()
	representatives := map[string]bool{
		"arcane": true, "feral": true, "enhancement": true,
		"retribution": true, "beast_mastery": true, "mutilate": true,
		"fury": true,
	}
	for _, b := range builds() {
		if !representatives[b.Key] {
			continue
		}
		p := b.player(b.races()[0])
		pool, _ := comparisonGearPool(b, p)
		selected := seedModeledGear(b, p, pool)
		for slot, item := range selected.Equipment.Items {
			if item.GetId() == 0 {
				if slot == 15 && selected.Equipment.Items[14].GetId() != 0 {
					continue
				}
				t.Fatalf("%s: slot %d empty", b.Key, slot)
			}
			if readGearCatalog()[item.Id].ModelVersion != 2 {
				t.Fatalf("%s: real gear in slot %d: %d", b.Key, slot, item.Id)
			}
		}
		if !fullyModeledV2(selected) {
			t.Fatalf("%s: modeled seed was not recognized as an equipment-only profile", b.Key)
		}
		if _, report, err := equippedHitOnly(b, selected); err != nil ||
			report.RawHitDelta != 0 || report.Model != "equipped-only-v2" {
			t.Fatalf("%s: modeled seed has a failed or paid hit adjustment: %v %+v", b.Key, err, report)
		}
	}
}
