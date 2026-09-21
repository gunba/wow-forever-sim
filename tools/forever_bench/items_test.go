//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/shaman"
)

func itemFixture(key string, slot proto.ItemSlot, id int32) *core.Simulation {
	req := tierRequest(key, true)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	req.Raid.Parties[0].Players[0].Equipment.Items[slot] = &proto.ItemSpec{Id: id}
	s := core.NewSim(req, simsignals.Signals{})
	s.Reset()
	return s
}

func TestForeverHeartTrinkets(t *testing.T) {
	s := tierSim("arcane", true)
	u := s.Raid.AllPlayerUnits[0]
	frozen := u.GetSpell(core.ActionID{ItemID: 249469})
	molten := u.GetSpell(core.ActionID{ItemID: 249470})
	if frozen == nil || molten == nil || frozen.SharedCD.Timer != molten.SharedCD.Timer {
		t.Fatal("heart trinkets must share their on-use lockout")
	}
	for _, tc := range []struct {
		spell   *core.Spell
		label   string
		bonuses stats.Stats
	}{
		{frozen, "ItemActive-249469", stats.Stats{stats.FrostPower: 29, stats.ShadowPower: 29}},
		{molten, "ItemActive-249470", stats.Stats{stats.AttackPower: 55, stats.RangedAttackPower: 55}},
	} {
		aura := u.GetAura(tc.label)
		if tc.spell.CD.Duration != 5*time.Minute || tc.spell.SharedCD.Duration != 15*time.Second || aura.Duration != 15*time.Second {
			t.Fatal("incorrect heart duration or cooldown")
		}
		before := u.GetStats()
		aura.Activate(s)
		for stat, amount := range tc.bonuses {
			if amount != 0 && math.Abs(u.GetStat(stats.Stat(stat))-before[stat]-amount) > 1e-8 {
				t.Fatalf("%s incorrect stat %d", tc.label, stat)
			}
		}
		aura.Deactivate(s)
		if !u.GetStats().Equals(before) {
			t.Fatal("heart bonus did not expire")
		}
	}
}

func TestForeverRelics(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		id := int32(0)
		if enabled {
			id = 220606
		}
		s := itemFixture("feral", proto.ItemSlot_ItemSlotRanged, id)
		dot := s.Raid.AllPlayerUnits[0].GetSpell(core.ActionID{SpellID: 9896}).Dot(s.Encounter.TargetUnits[0])
		want := int32(6)
		if enabled {
			want = 7
		}
		if dot.NumberOfTicks != want || dot.TickPeriod() != 2*time.Second {
			t.Fatal("Idol of the Dream did not add one Rip tick")
		}
	}
	a := itemFixture("elemental", proto.ItemSlot_ItemSlotRanged, 0).Raid.AllPlayerUnits[0]
	b := itemFixture("elemental", proto.ItemSlot_ItemSlotRanged, 228176).Raid.AllPlayerUnits[0]
	for _, spell := range a.Spellbook {
		if spell.SpellCode == shaman.SpellCode_ShamanLightningBolt {
			if math.Abs(b.GetSpell(spell.ActionID).BonusCritRating-spell.BonusCritRating-core.SpellCritRatingPerCritChance) > 1e-8 {
				t.Fatal("Totem of Thunder must give Lightning Bolt one percentage point of crit")
			}
		}
	}
	a = itemFixture("retribution", proto.ItemSlot_ItemSlotRanged, 0).Raid.AllPlayerUnits[0]
	b = itemFixture("retribution", proto.ItemSlot_ItemSlotRanged, 249442).Raid.AllPlayerUnits[0]
	for _, id := range []int32{20154, 20375, 21082} {
		x, y := a.GetSpell(core.ActionID{SpellID: id}), b.GetSpell(core.ActionID{SpellID: id})
		if x == nil || y == nil || x.Cost.Multiplier-y.Cost.Multiplier != 5 {
			t.Fatalf("Libram seal-cost reduction missing for %d", id)
		}
	}
}

func TestForeverWolfsheadEnergy(t *testing.T) {
	var energy []float64
	for _, id := range []int32{0, 8345} {
		s := itemFixture("feral", proto.ItemSlot_ItemSlotHead, id)
		u := s.Raid.AllPlayerUnits[0]
		u.SpendEnergy(s, u.CurrentEnergy(), u.NewEnergyMetrics(core.ActionID{SpellID: 1}))
		tf := u.GetSpell(core.ActionID{SpellID: 9846})
		tf.ApplyEffects(s, u, tf)
		energy = append(energy, u.CurrentEnergy())
	}
	if energy[1]-energy[0] != 20 {
		t.Fatalf("Wolfshead must add 20 Tiger's Fury energy: %v", energy)
	}
}
