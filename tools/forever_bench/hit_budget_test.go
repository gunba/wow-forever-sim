//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	googleProto "google.golang.org/protobuf/proto"
)

func TestRangedScopeReducesPaidHunterHit(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "beast_mastery" && b.Key != "marksmanship" {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		// Isolate ranged requirements. Both current Hunter presets can summon
		// Hawk, whose unresolved melee hit model otherwise sets the overall cap.
		p.Rotation = core.APLRotationFromJsonString(`{
			"type":"TypeAPL","priorityList":[
				{"action":{"castSpell":{"spellId":{"spellId":25295}}}}
			]}`)
		_, before, err := capHit(b, p)
		if err != nil {
			t.Fatal(err)
		}
		p.Equipment.Items[proto.ItemSlot_ItemSlotRanged].Enchant = 2523
		_, after, err := capHit(b, p)
		if err != nil {
			t.Fatal(err)
		}
		// The scope cannot refund more hit budget than remained payable; the
		// rest of its rating becomes overcap and does not create offensive stats.
		want := math.Min(30, before.RawHitDelta)
		if got := before.RawHitDelta - after.RawHitDelta; math.Abs(got-want) > 1e-8 {
			t.Errorf("%s: scope saved %v raw hit, want %v; requirements %+v", b.Key, got, want, after.Requirements)
		}
	}
}

func TestHitBudgetAllBuilds(t *testing.T) {
	for _, b := range builds() {
		t.Run(b.Key, func(t *testing.T) {
			p := b.player(b.races()[0])
			original := googleProto.Clone(p)
			normalized, report, err := capHit(b, p)
			if err != nil {
				t.Fatal(err)
			}
			if !googleProto.Equal(p, original) {
				t.Fatal("normalization mutated the search baseline")
			}
			if !googleProto.Equal(p.Equipment, normalized.Equipment) ||
				!googleProto.Equal(p.Consumes, normalized.Consumes) ||
				!googleProto.Equal(p.Buffs, normalized.Buffs) {
				t.Fatal("normalization altered equipment, enchants, consumes or buffs")
			}
			if report.MeleeAdded != report.SpellAdded || math.Abs(report.Balance) > 1e-8 {
				t.Fatalf("non-shared or unbalanced exchange: %+v", report)
			}
			var balance float64
			for _, entry := range report.Entries {
				balance += entry.Delta * entry.CostPerUnit
				if entry.GearAmount+entry.Delta < -1e-8 {
					t.Fatalf("spent more than the item's stat pool: %+v", entry)
				}
			}
			if math.Abs(balance) > 1e-8 {
				t.Fatalf("ledger is not budget neutral: %v", balance)
			}
			req := request(normalized, 1, 1)
			env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
			requirements, _ := hitRequirements(b, normalized, env.Raid.Parties[0].Players[0].GetCharacter(), env.Encounter.TargetUnits[0])
			for _, requirement := range requirements {
				if requirement.AdditionalPercent > 1e-8 {
					t.Errorf("not hit capped: %+v", requirement)
				}
			}
		})
	}
}

func TestOffensiveBudgetPricesAndSharing(t *testing.T) {
	b := build{Key: "enhancement", Class: proto.Class_ClassShaman}
	gear := stats.Stats{
		stats.AttackPower: 20, stats.RangedAttackPower: 20,
		stats.SpellPower: 14, stats.Intellect: 5, stats.Agility: 6, stats.Strength: 7,
		stats.MeleeCrit: 1, stats.SpellCrit: .5, stats.Stamina: 100, stats.Spirit: 100,
	}
	var total float64
	var delta stats.Stats
	for _, donor := range offensiveGearBudget(b, gear, 14) {
		total += donor.amount * donor.cost
		donor.apply(&delta, -donor.amount)
	}
	// 20 AP is ten points, 14 SP is twelve, and 1.5% crit is 21 raw rating.
	if total != 10+12+5+6+7+21 {
		t.Fatalf("wrong or duplicated prices: %v", total)
	}
	if delta[stats.AttackPower] != -20 || delta[stats.RangedAttackPower] != -20 ||
		delta[stats.MeleeCrit] != -1.5 || delta[stats.SpellCrit] != -1.5 ||
		delta[stats.Stamina] != 0 || delta[stats.Spirit] != 0 {
		t.Fatalf("wrong simulator deltas: %v", delta)
	}
}

func TestRotationHitSelection(t *testing.T) {
	rotation := core.APLRotationFromJsonString(`{
		"priorityList":[
			{"action":{"condition":{"spellIsReady":{"spellId":{"spellId":99}}},
				"castSpell":{"spellId":{"spellId":1}}}},
			{"action":{"channelSpell":{"spellId":{"spellId":2}}}},
			{"action":{"multidot":{"spellId":{"spellId":3},"maxDots":1}}},
			{"hide":true,"action":{"castSpell":{"spellId":{"spellId":4}}}},
			{"action":{"castSpell":{"spellId":{"itemId":5}}}},
			{"action":{"autocastOtherCooldowns":{}}}
		]}`)
	ids, cooldowns := rotationCastIDs(rotation)
	if !ids[core.ActionID{SpellID: 1}] || !ids[core.ActionID{SpellID: 2}] ||
		!ids[core.ActionID{SpellID: 3}] || !ids[core.ActionID{ItemID: 5}] ||
		ids[core.ActionID{SpellID: 4}] || ids[core.ActionID{SpellID: 99}] ||
		ids[core.ActionID{}] || !cooldowns {
		t.Fatalf("wrong rotational spell selection: %v, %v", ids, cooldowns)
	}
}

func TestHitBudgetRejectsManualHit(t *testing.T) {
	b := builds()[0]
	p := b.player(b.races()[0])
	p.BonusStats = &proto.UnitStats{Stats: stats.Stats{stats.MeleeHit: 1}.ToFloatArray()}
	if _, _, err := capHit(b, p); err == nil {
		t.Fatal("manual/free hit was accepted")
	}
}

func TestHitBudgetEquipmentScaling(t *testing.T) {
	for _, b := range builds() {
		t.Run(b.Key, func(t *testing.T) {
			p := b.player(b.races()[0])
			_, original, err := capHit(b, p)
			if err != nil {
				t.Fatal(err)
			}
			for _, scale := range []float64{1.1, 1.5} {
				p.EquipmentScale = scale
				normalized, report, err := capHit(b, p)
				if err != nil {
					t.Fatal(err)
				}
				if math.Abs(report.OffensiveBudgetBefore-original.OffensiveBudgetBefore*scale) > 1e-8 ||
					math.Abs(report.Balance) > 1e-8 {
					t.Fatalf("scaled item budget is wrong: %+v", report)
				}
				req := request(normalized, 1, 1)
				env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
				requirements, _ := hitRequirements(b, normalized, env.Raid.Parties[0].Players[0].GetCharacter(), env.Encounter.TargetUnits[0])
				for _, requirement := range requirements {
					if requirement.AdditionalPercent > 1e-8 {
						t.Errorf("scaled profile is not hit capped: %+v", requirement)
					}
				}
			}
		})
	}
}

func TestUnusedSpellSchoolsDoNotEraseHitTalents(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "arcane" {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		p.TalentsString = "05"
		p.Rotation = core.APLRotationFromJsonString(`{
			"priorityList":[{"action":{"channelSpell":{"spellId":{"spellId":25345}}}}]
		}`)
		_, report, err := capHit(b, p)
		if err != nil {
			t.Fatal(err)
		}
		if math.Abs(report.SpellFinal-11) > 1e-8 {
			t.Fatalf("five points of Arcane Focus should reduce generic hit to 11%%: %+v", report)
		}
	}
}

func TestHitBudgetNeverCreditsExcessGearHit(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "fury" {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		// An isolated fixture; no real catalog record is changed.
		const fixtureID = 999999
		fixture := core.ItemsByID[p.Equipment.Items[0].Id]
		fixture.ID = fixtureID
		fixture.Stats[stats.MeleeHit] += 20
		core.ItemsByID[fixtureID] = fixture
		defer delete(core.ItemsByID, fixtureID)
		p.Equipment.Items[0].Id = fixtureID
		_, orc, err := capHit(b, p)
		if err != nil {
			t.Fatal(err)
		}
		p.Race = proto.Race_RaceTauren
		_, tauren, err := capHit(b, p)
		if err != nil {
			t.Fatal(err)
		}
		if orc.RawHitDelta != 0 || tauren.RawHitDelta != 0 {
			t.Fatalf("overcapped hit was converted into offensive stats: Orc %+v, Tauren %+v", orc, tauren)
		}
		for _, report := range []hitAdjustment{orc, tauren} {
			if report.MeleeFinal < 9 || math.Abs(report.Balance) > 1e-8 {
				t.Fatalf("incorrect hit cap or free budget: %+v", report)
			}
			for _, entry := range report.Entries {
				if entry.Delta != 0 {
					t.Fatalf("excess hit produced a budget exchange: %+v", entry)
				}
			}
		}
	}
}

func TestTaurenRacialReducesPaidHitWithoutCreatingBudget(t *testing.T) {
	var fury build
	for _, b := range builds() {
		if b.Key == "fury" {
			fury = b
			break
		}
	}
	p := fury.player(proto.Race_RaceOrc)
	for slot, equipped := range p.Equipment.Items {
		item := core.ItemsByID[equipped.GetId()]
		if item.Stats[stats.MeleeHit]+item.Stats[stats.SpellHit] == 0 {
			continue
		}
		fixtureID := int32(999900 + slot)
		item.ID = fixtureID
		item.Stats[stats.MeleeHit], item.Stats[stats.SpellHit] = 0, 0
		core.ItemsByID[fixtureID] = item
		defer delete(core.ItemsByID, fixtureID)
		equipped.Id = fixtureID
	}
	_, orc, err := capHit(fury, p)
	if err != nil {
		t.Fatal(err)
	}
	p.Race = proto.Race_RaceTauren
	_, tauren, err := capHit(fury, p)
	if err != nil {
		t.Fatal(err)
	}
	if orc.RawHitDelta <= 0 || tauren.RawHitDelta < 0 ||
		math.Abs(orc.RawHitDelta-tauren.RawHitDelta-10) > 1e-8 {
		t.Fatalf("racial hit did not save its paid budget: Orc %+v, Tauren %+v", orc, tauren)
	}
}

func TestEquippedHitOnlyPreservesStatsAndAppliesTaurenRacial(t *testing.T) {
	var fury build
	for _, b := range builds() {
		if b.Key == "fury" {
			fury = b
			break
		}
	}
	*modeledOnly = true
	defer func() { *modeledOnly = false }()
	p := fury.player(proto.Race_RaceOrc)
	pool, _ := comparisonGearPool(fury, p)
	p = seedModeledGear(fury, p, pool)
	before := googleProto.Clone(p)
	unchanged, orc, err := equippedHitOnly(fury, p)
	if err != nil {
		t.Fatal(err)
	}
	if !googleProto.Equal(p, before) || !googleProto.Equal(p, unchanged) ||
		orc.Model != "equipped-only-v2" || orc.RawHitDelta != 0 ||
		orc.MeleeAdded != 0 || orc.SpellAdded != 0 || len(orc.Entries) != 0 {
		t.Fatalf("modeled hit generated or exchanged stats: %+v", orc)
	}
	p.Race = proto.Race_RaceTauren
	_, tauren, err := equippedHitOnly(fury, p)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(tauren.MeleeFinal-orc.MeleeFinal-1) > 1e-8 ||
		tauren.RawHitDelta != 0 {
		t.Fatalf("Tauren racial did not grant one natural hit without conversion: %+v %+v", orc, tauren)
	}
}
