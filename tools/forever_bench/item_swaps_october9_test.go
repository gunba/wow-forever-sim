//go:build with_db

package main

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func swapOctober9Fixture(swap *proto.ItemSwap) (*core.Simulation, *core.Character) {
	req := historyTalentFixture("arms", nil)
	p := req.Raid.Parties[0].Players[0]
	p.Equipment.Items[14] = &proto.ItemSpec{Id: 920000360, Enchant: 1900}
	p.Equipment.Items[15] = &proto.ItemSpec{}
	p.EnableItemSwap = true
	p.ItemSwap = swap
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim, sim.Raid.Parties[0].Players[0].GetCharacter()
}

func TestOctober9SwapSameWeaponEnchantAndSuffix(t *testing.T) {
	for _, spec := range []*proto.ItemSpec{{Id: 920000360, Enchant: 241}, {Id: 920000360, Enchant: 1900, RandomSuffix: 5}} {
		sim, c := swapOctober9Fixture(&proto.ItemSwap{MhItem: spec})
		if !c.ItemSwap.IsEnabled() {
			t.Fatal("same-ID enchant/suffix change did not enable swap")
		}
		before := c.GetStats()
		c.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand})
		if c.MainHand().Enchant.EffectID != spec.Enchant || c.MainHand().RandomSuffix.ID != spec.RandomSuffix {
			t.Fatal("swap lost enchant or random suffix")
		}
		c.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand})
		if c.ItemSwap.IsSwapped() || c.MainHand().Enchant.EffectID != 1900 || c.MainHand().RandomSuffix.ID != 0 || c.GetStats() != before {
			t.Fatal("swap round-trip drifted")
		}
	}
}

func TestOctober9SwapAtomicAndSequentialWithinGCD(t *testing.T) {
	for _, sequential := range []bool{false, true} {
		sim, c := swapOctober9Fixture(&proto.ItemSwap{MhItem: &proto.ItemSpec{Id: 920000334, Enchant: 241}, OhItem: &proto.ItemSpec{Id: 920000383, Enchant: 929}})
		stats := c.GetStats()
		both := []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand, proto.ItemSlot_ItemSlotOffHand}
		if sequential {
			c.ItemSwap.SwapItems(sim, both[:1])
			c.ItemSwap.SwapItems(sim, both[1:])
		} else {
			c.ItemSwap.SwapItems(sim, both)
		}
		if c.MainHand().ID != 920000334 || c.OffHand().ID != 920000383 || !c.PseudoStats.CanBlock || !c.ItemSwap.IsSwapped() {
			t.Fatalf("sequential=%v: sword/shield or slot-state incorrect", sequential)
		}
		gcd := c.GCD.ReadyAt()
		sim.CurrentTime = 100 * time.Millisecond
		c.ItemSwap.SwapItems(sim, both)
		if c.MainHand().ID != 920000360 || c.OffHand().ID != 0 || c.PseudoStats.CanBlock || c.ItemSwap.IsSwapped() || c.GetStats() != stats {
			t.Fatal("2H round-trip not atomic or stats drifted")
		}
		if gcd != 1500*time.Millisecond || c.GCD.ReadyAt() != gcd {
			t.Fatal("repeated swap extended existing GCD")
		}
	}
}

func TestOctober9SwapIterationReset(t *testing.T) {
	sim, c := swapOctober9Fixture(&proto.ItemSwap{MhItem: &proto.ItemSpec{Id: 920000334, Enchant: 241}, OhItem: &proto.ItemSpec{Id: 920000383, Enchant: 929}})
	before := c.GetStats()
	c.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand, proto.ItemSlot_ItemSlotOffHand})
	sim.Cleanup()
	sim.Reset()
	if c.ItemSwap.IsSwapped() || c.MainHand().ID != 920000360 || c.GetStats() != before {
		t.Fatalf("swap reset drifted stats: %v", c.GetStats().Subtract(before))
	}
}

func TestOctober9SwapUniqueRollback(t *testing.T) {
	sim, c := swapOctober9Fixture(&proto.ItemSwap{MhItem: &proto.ItemSpec{Id: 920000334}, OhItem: &proto.ItemSpec{Id: 920000383}})
	*c.ItemSwap.GetItem(proto.ItemSlot_ItemSlotMainHand) = core.ItemsByID[11086]
	*c.ItemSwap.GetItem(proto.ItemSlot_ItemSlotOffHand) = core.ItemsByID[11086]
	before := c.GetStats()
	gcd := c.GCD.ReadyAt()
	var failure any
	func() {
		defer func() { failure = recover() }()
		c.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand, proto.ItemSlot_ItemSlotOffHand})
	}()
	if failure == nil {
		t.Fatal("unique-equipped duplicate accepted")
	}
	if c.MainHand().ID != 920000360 || c.OffHand().ID != 0 || c.GetStats() != before || c.GCD.ReadyAt() != gcd || c.ItemSwap.IsSwapped() {
		t.Fatal("invalid swap mutated live equipment/stats/GCD")
	}
}
