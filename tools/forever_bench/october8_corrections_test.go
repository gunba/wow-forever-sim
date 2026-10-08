//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/warlock/dps"
)

func TestOctober8RogueOneHandAxes(t *testing.T) {
	for _, hand := range []proto.HandType{proto.HandType_HandTypeOneHand, proto.HandType_HandTypeMainHand, proto.HandType_HandTypeOffHand, proto.HandType_HandTypeTwoHand} {
		item := core.Item{Type: proto.ItemType_ItemTypeWeapon, WeaponType: proto.WeaponType_WeaponTypeAxe, HandType: hand}
		want := hand != proto.HandType_HandTypeTwoHand
		if got := classCanEquip(proto.Class_ClassRogue, item); got != want {
			t.Errorf("Rogue axe hand %v: allowed %v, want %v", hand, got, want)
		}
	}
}

func TestOctober8ShamanWeaponMastery(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		for _, armed := range []bool{true, false} {
			req := historyTalentFixture("enhancement", map[string]int{"stormstrike": 1})
			req.SimOptions.Ruleset = ruleset
			req.Raid.Buffs, req.Raid.Debuffs, req.Raid.Parties[0].Buffs = nil, nil, nil
			p := req.Raid.Parties[0].Players[0]
			p.ForeverTier1Bonuses = false
			p.Consumes = &proto.Consumes{MainHandImbue: proto.WeaponImbue_WindfuryWeapon}
			if !armed {
				p.Equipment = &proto.EquipmentSpec{}
				p.Consumes = nil
			}
			sim := core.NewSim(req, simsignals.Signals{})
			c := sim.Raid.Parties[0].Players[0].GetCharacter()
			want := 1.0
			if ruleset == proto.Ruleset_RulesetForever && armed {
				want = 1.10
			}
			seenWF := false
			for _, spell := range c.Spellbook {
				if spell.SpellSchool != core.SpellSchoolPhysical || !spell.ProcMask.Matches(core.ProcMaskMelee) {
					continue
				}
				if spell.ActionID.SpellID == 16362 {
					seenWF = true
				}
				spellWant := want
				if spell.ProcMask.Matches(core.ProcMaskMeleeOH) {
					spellWant = 1 // No legal off-hand weapon in this fixture.
				}
				if math.Abs(spell.DamageMultiplier-spellWant) > 1e-9 {
					t.Errorf("%v armed %v spell %v multiplier %v, want %v", ruleset, armed, spell.ActionID, spell.DamageMultiplier, spellWant)
				}
			}
			if shock := c.GetSpell(core.ActionID{SpellID: 10414}); shock == nil || shock.DamageMultiplier != 1 {
				t.Fatal("Weapon Mastery altered Earth Shock")
			}
			if armed && !seenWF {
				t.Fatal("fixture did not register Windfury Weapon")
			}
		}
	}
}

func TestOctober8DemonicBrandImpThreat(t *testing.T) {
	req := warlockEffectFixture(map[string]int{"demonicBrand": 3})
	sim := core.NewSim(req, simsignals.Signals{})
	w := sim.Raid.Parties[0].Players[0].(*dps.DpsWarlock).GetWarlock()
	for _, pet := range w.BasePets {
		id, want := int32(1293697), 3.0
		if pet == w.Imp {
			id, want = 1293698, 1
		}
		spell := pet.GetSpell(core.ActionID{SpellID: id})
		if spell == nil {
			t.Fatalf("missing brand for %s", pet.Label)
		}
		if math.Abs(spell.ThreatMultiplier-want) > 1e-9 {
			t.Errorf("%s Brand threat multiplier %v, want %v", pet.Label, spell.ThreatMultiplier, want)
		}
	}
}
