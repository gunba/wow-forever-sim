//go:build with_db

package main

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func TestNaturesBountyRoutesResourcesAndTicksEnergy(t *testing.T) {
	for _, set := range []struct {
		name string
		ids  map[int]int32
	}{
		{"Wildheart", map[int]int32{0: 16720, 2: 16718, 4: 16706, 8: 16719, 9: 16715}},
		{"Feralheart", map[int]int32{0: 16720, 2: 16718, 4: 16706, 8: 16719}},
	} {
		t.Run(set.name, func(t *testing.T) {
			req := racialFixture("feral", proto.Race_RaceTauren)
			p := req.Raid.Parties[0].Players[0]
			p.Rotation = &proto.APLRotation{}
			for slot, id := range set.ids {
				// Exercise the name-based fallback. Catalog set IDs separately
				// suppress unresolved procs; this test does not enable them.
				item := core.ItemsByID[id]
				item.ID = 990000 + int32(slot)
				item.SetName, item.SetID = set.name+" Raiment", 0
				id = item.ID
				core.ItemsByID[id] = item
				t.Cleanup(func() { delete(core.ItemsByID, id) })
				p.Equipment.Items[slot] = &proto.ItemSpec{Id: id}
			}
			s := core.NewSim(req, simsignals.Signals{})
			s.Options.Interactive = true
			s.Reset()
			u := s.Raid.AllPlayerUnits[0]
			mana := u.GetAura("Nature's Bounty (Mana)")
			energy := u.GetAura("Nature's Bounty (Energy)")
			rage := u.GetAura("Nature's Bounty (Rage)")
			if mana == nil || energy == nil || rage == nil {
				t.Fatal("missing independent spellcast, outgoing-melee and incoming-melee triggers")
			}
			id := core.ActionID{SpellID: 27784}
			u.SpendMana(s, 1000, u.NewManaMetrics(id))
			u.SpendEnergy(s, u.CurrentEnergy(), u.NewEnergyMetrics(id))
			beforeMana, beforeRage := u.CurrentMana(), u.CurrentRage()
			hit := &core.SpellResult{Target: u, Outcome: core.OutcomeHit}
			melee := u.AutoAttacks.MHAuto()
			for i := 0; i < 10000 && u.CurrentRage() == beforeRage; i++ {
				rage.OnSpellHitTaken(rage, s, melee, hit)
			}
			if u.CurrentRage() != beforeRage+10 || u.CurrentEnergy() != 0 || u.CurrentMana() != beforeMana {
				t.Fatal("incoming melee did not grant only ten Rage")
			}
			magic := &core.Spell{ProcMask: core.ProcMaskSpellDamage}
			for i := 0; i < 10000 && u.CurrentMana() == beforeMana; i++ {
				mana.OnCastComplete(mana, s, magic)
			}
			if u.CurrentMana() != beforeMana+200 || u.CurrentEnergy() != 0 {
				t.Fatal("spellcast did not grant only 200 Mana")
			}
			restore := u.GetSpell(id)
			if restore == nil {
				t.Fatal("missing periodic Energy restoration")
			}
			dot := restore.Hot(u)
			hit.Target = s.Encounter.TargetUnits[0]
			for i := 0; i < 10000 && !dot.IsActive(); i++ {
				energy.OnSpellHitDealt(energy, s, melee, hit)
			}
			if !dot.IsActive() || u.CurrentEnergy() != 0 || dot.TickLength != time.Second || dot.NumberOfTicks != 5 {
				t.Fatal("Energy proc must start five one-second ticks, not an instant return")
			}
			for i := 1; i <= 5; i++ {
				dot.TickOnce(s)
				if u.CurrentEnergy() != float64(i*4) {
					t.Fatalf("Energy after tick %d = %g", i, u.CurrentEnergy())
				}
			}
		})
	}
}
