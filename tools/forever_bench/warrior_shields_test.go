//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func TestWarriorShieldTalentsFollowEquipment(t *testing.T) {
	for _, startShield := range []bool{false, true} {
		req := racialFixture("fury", proto.Race_RaceOrc)
		p := req.Raid.Parties[0].Players[0]
		p.ForeverTier1Bonuses = false
		p.Rotation = &proto.APLRotation{}
		// Isolate the three equipment conditions, not a legal talent build.
		for _, b := range builds() {
			if b.Key != "fury" {
				continue
			}
			config := loadTalents(b)
			points, _ := config.decode("")
			index := 0
			ranks := map[string]int{"defiance": 3, "bastion": 5, "masterOfDefense": 2}
			for _, tree := range config.Trees {
				for _, talent := range tree.Talents {
					points[index] = ranks[talent.Field]
					index++
				}
			}
			p.TalentsString = config.encode(points)
		}
		p.Equipment = &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 17)}
		for i := range p.Equipment.Items {
			p.Equipment.Items[i] = &proto.ItemSpec{}
		}
		p.Equipment.Items[14] = &proto.ItemSpec{Id: 279261}
		equipped, swapped := int32(279261), int32(279262)
		if startShield {
			equipped, swapped = swapped, equipped
		}
		p.Equipment.Items[15] = &proto.ItemSpec{Id: equipped}
		p.EnableItemSwap = true
		p.ItemSwap = &proto.ItemSwap{OhItem: &proto.ItemSpec{Id: swapped}}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		character := sim.Raid.Parties[0].Players[0].GetCharacter()
		defensive := character.GetAura("Defensive Stance")
		bastion := character.GetAura("Bastion")
		master := character.GetAura("Master of Defense")
		for _, name := range []string{"Battle Stance", "Defensive Stance", "Berserker Stance", "Bastion"} {
			character.GetAura(name).Deactivate(sim)
		}
		baseDamage := character.PseudoStats.DamageDealtMultiplier
		baseThreat := character.PseudoStats.ThreatMultiplier
		bastion.OnReset(bastion, sim)
		defensive.Activate(sim)
		for swap := 0; swap < 3; swap++ {
			shield := startShield != (swap%2 == 1)
			wantThreat, wantDamage := baseThreat*1.3, baseDamage*.9
			if shield {
				wantThreat *= 1.15
				wantDamage *= 1.1
			}
			if math.Abs(character.PseudoStats.ThreatMultiplier-wantThreat) > 1e-9 ||
				math.Abs(character.PseudoStats.DamageDealtMultiplier-wantDamage) > 1e-9 {
				t.Fatalf("start shield=%v, swap=%d: threat/damage = %.4f/%.4f, want %.4f/%.4f",
					startShield, swap, character.PseudoStats.ThreatMultiplier,
					character.PseudoStats.DamageDealtMultiplier, wantThreat, wantDamage)
			}
			before := character.CurrentRage()
			master.OnSpellHitTaken(master, sim, nil, &core.SpellResult{Outcome: core.OutcomeDodge})
			gain := 0.0
			if shield {
				gain = 5
			}
			if got := character.CurrentRage() - before; got != gain {
				t.Fatalf("shield=%v: Master of Defense gave %v rage, want %v", shield, got, gain)
			}
			character.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotOffHand})
		}
		defensive.Deactivate(sim)
		if math.Abs(character.PseudoStats.ThreatMultiplier-baseThreat) > 1e-9 {
			t.Fatal("Defiance left a threat multiplier after leaving Defensive Stance")
		}
	}
}
