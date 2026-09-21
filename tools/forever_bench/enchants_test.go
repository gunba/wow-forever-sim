//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func enchantFixture(key string, enchants map[proto.ItemSlot]int32) (*core.Simulation, *core.Character) {
	for _, b := range builds() {
		if b.Key != key {
			continue
		}
		p := b.player(proto.Race_RaceTauren)
		p.Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
		for slot, id := range enchants {
			p.Equipment.Items[slot].Enchant = id
		}
		sim := core.NewSim(request(p, 1, 1), simsignals.Signals{})
		sim.Reset()
		character := sim.Raid.Parties[0].Players[0].GetCharacter()
		character.AutoAttacks.CancelAutoSwing(sim)
		return sim, character
	}
	panic("unknown enchant fixture")
}

func TestForeverEnchantHasteScopes(t *testing.T) {
	_, base := enchantFixture("arms", nil)
	_, counterweight := enchantFixture("arms", map[proto.ItemSlot]int32{proto.ItemSlot_ItemSlotMainHand: 34})
	if math.Abs(counterweight.SwingSpeed()/base.SwingSpeed()-1.03) > 1e-9 {
		t.Fatal("counterweight did not increase melee speed by 3%")
	}
	if counterweight.RangedSwingSpeed() != base.RangedSwingSpeed() {
		t.Fatal("counterweight increased ranged speed")
	}
	_, rapidity := enchantFixture("arms", map[proto.ItemSlot]int32{
		proto.ItemSlot_ItemSlotHead: 2543, proto.ItemSlot_ItemSlotLegs: 2543,
	})
	if math.Abs(rapidity.SwingSpeed()/base.SwingSpeed()-1.02) > 1e-9 ||
		math.Abs(rapidity.RangedSwingSpeed()/base.RangedSwingSpeed()-1.02) > 1e-9 {
		t.Fatal("two Arcanums did not grant 2% melee/ranged attack speed")
	}
}

func TestRapidityDoesNotIncreaseEnergyRegeneration(t *testing.T) {
	recovery := func(enchants map[proto.ItemSlot]int32) float64 {
		sim, character := enchantFixture("feral", enchants)
		character.SpendEnergy(sim, 90, character.EnergyRefundMetrics)
		start := character.CurrentEnergy()
		for sim.CurrentTime < time.Second {
			sim.Step()
		}
		return character.CurrentEnergy() - start
	}
	base := recovery(nil)
	rapidity := recovery(map[proto.ItemSlot]int32{
		proto.ItemSlot_ItemSlotHead: 2543, proto.ItemSlot_ItemSlotLegs: 2543,
	})
	if math.Abs(rapidity-base) > 1e-9 {
		t.Fatalf("attack-speed enchants changed Energy recovery: %v vs %v", rapidity, base)
	}
	general := recovery(map[proto.ItemSlot]int32{proto.ItemSlot_ItemSlotHands: 931})
	if math.Abs(general/base-1.01) > 1e-9 {
		t.Fatalf("Minor Haste did not follow the general-haste Energy model: %v vs %v", general, base)
	}
}

func TestForeverPrecisionScopeOnlyAffectsRangedCrit(t *testing.T) {
	_, base := enchantFixture("marksmanship", nil)
	sim, scoped := enchantFixture("marksmanship", map[proto.ItemSlot]int32{proto.ItemSlot_ItemSlotRanged: 8720})
	for reset := 0; reset < 2; reset++ {
		ranged := 0
		for _, spell := range scoped.Spellbook {
			original := base.GetSpell(spell.ActionID)
			if original == nil {
				continue
			}
			want := original.BonusCritRating
			if spell.CastType == proto.CastType_CastTypeRanged {
				want += 2
				ranged++
			}
			if spell.BonusCritRating != want {
				t.Fatalf("%v crit %v, want %v", spell.ActionID, spell.BonusCritRating, want)
			}
		}
		if ranged == 0 {
			t.Fatal("scope fixture has no ranged attacks")
		}
		aura := scoped.GetAura("SAF-T Ultra Precision Scope")
		aura.Deactivate(sim)
		aura.Activate(sim)
	}
}

func TestGearEnchantCoverageAndLegality(t *testing.T) {
	for _, b := range builds() {
		p := b.player(b.races()[0])
		prepareGearEnchants(b, p)
		if err := validateGearEnchants(p); err != nil {
			t.Fatalf("%s: %v", b.Key, err)
		}
		for slot, spec := range p.Equipment.Items {
			if len(legalEnchants(p, slot)) > 0 && spec.GetEnchant() == 0 {
				t.Fatalf("%s slot %d missing enchant", b.Key, slot)
			}
		}
		p.Equipment.Items[proto.ItemSlot_ItemSlotNeck].Enchant = 2504
		if validateGearEnchants(p) == nil {
			t.Fatal("weapon enchant accepted on a necklace")
		}
	}
	for _, enchant := range benchmarkEnchants() {
		if enchant.EffectId == 2588 || enchant.EffectId == 2717 {
			t.Fatal("raid reward enchant entered the benchmark pool")
		}
	}
}

func TestEnchantPreparationReplacesHealingOnlyCasterBracers(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "fire" {
			continue
		}
		p := b.player(b.races()[0])
		p.Equipment.Items[proto.ItemSlot_ItemSlotWrist].Enchant = 2566
		prepareGearEnchants(b, p)
		if p.Equipment.Items[proto.ItemSlot_ItemSlotWrist].Enchant != 1883 {
			t.Fatal("expected seven Intellect instead of healing-only bracers")
		}
	}
}
