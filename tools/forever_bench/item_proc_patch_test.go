//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func TestClient70009WeaponProcDamage(t *testing.T) {
	for _, tc := range []struct {
		item, spell int32
		min, max    float64
	}{
		{18671, 13442, 90, 90},
		{13401, 17505, 85, 85},
		{11902, 18089, 73.5, 122.5}, // 98 mean, 0.5 variance.
	} {
		t.Run(core.ActionID{SpellID: tc.spell}.String(), func(t *testing.T) {
			req := racialFixture("fury", proto.Race_RaceOrc)
			p := req.Raid.Parties[0].Players[0]
			p.Rotation = &proto.APLRotation{}
			// Exercise registered definitions without making excluded items
			// available in the live catalog or projected-gear pool.
			item := core.ItemsByID[p.Equipment.Items[proto.ItemSlot_ItemSlotMainHand].Id]
			old, existed := core.ItemsByID[tc.item]
			item.ID = tc.item
			core.ItemsByID[tc.item] = item
			t.Cleanup(func() {
				if existed {
					core.ItemsByID[tc.item] = old
				} else {
					delete(core.ItemsByID, tc.item)
				}
			})
			p.Equipment.Items[proto.ItemSlot_ItemSlotMainHand] = &proto.ItemSpec{Id: tc.item}
			s := core.NewSim(req, simsignals.Signals{})
			s.Options.Interactive = true
			s.Reset()
			u := s.Raid.AllPlayerUnits[0]
			spell := u.GetSpell(core.ActionID{SpellID: tc.spell})
			if spell == nil {
				t.Fatal("missing weapon proc")
			}
			spell.Flags |= core.SpellFlagIgnoreModifiers | core.SpellFlagIgnoreResists
			spell.BonusCritRating = -10000
			spell.BonusHitRating = 10000
			target := s.Encounter.TargetUnits[0]
			landed := 0
			for i := 0; i < 100; i++ {
				before := spell.SpellMetrics[target.UnitIndex].TotalDamage
				misses := spell.SpellMetrics[target.UnitIndex].Misses
				spell.Cast(s, target)
				if spell.SpellMetrics[target.UnitIndex].Misses > misses {
					continue // Magic retains an unavoidable miss chance.
				}
				landed++
				damage := spell.SpellMetrics[target.UnitIndex].TotalDamage - before
				if damage < tc.min-1e-9 || damage > tc.max+1e-9 {
					t.Fatalf("proc damage %g outside client range [%g,%g]", damage, tc.min, tc.max)
				}
			}
			if landed < 80 {
				t.Fatalf("only %d proc hits observed", landed)
			}
		})
	}
}
