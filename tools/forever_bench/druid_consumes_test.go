//go:build with_db

package main

import (
	"encoding/json"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/druid"
)

func TestDemonicRuneDamagesOnlyItsUser(t *testing.T) {
	var shadow build
	for _, candidate := range builds() {
		if candidate.Key == "shadow" {
			shadow = candidate
			break
		}
	}
	row := run(shadow, prepare(shadow, proto.Race_RaceUndead), 100, 20296421)
	var metrics struct {
		Actions []struct {
			ID struct {
				SpellID int32
			}
			Targets []struct {
				Hits           int32
				Damage         float64
				ResistedDamage float64
			}
		}
	}
	if err := json.Unmarshal(row.Metrics, &metrics); err != nil {
		t.Fatal(err)
	}
	for _, action := range metrics.Actions {
		if action.ID.SpellID != 16666 {
			continue
		}
		if len(action.Targets) < 2 || action.Targets[0].Damage != 0 || action.Targets[1].Hits == 0 {
			t.Fatalf("rune must hit the player, not the boss: %+v", action.Targets)
		}
		// Client spell 16666 rolls 800 ±25% before Shadow resistance.
		rolled := (action.Targets[1].Damage + action.Targets[1].ResistedDamage) / float64(action.Targets[1].Hits)
		if rolled < 600 || rolled > 1000 {
			t.Fatalf("rune self-hit mean outside the client range: %.2f", rolled)
		}
		return
	}
	t.Fatal("Demonic Rune did not record its self-damage effect")
}

func TestForeverManaConsumablesPreserveForms(t *testing.T) {
	for _, key := range []string{"balance", "feral", "bear"} {
		fixture := key
		if key == "bear" {
			fixture = "feral"
		}
		req := racialFixture(fixture, proto.Race_RaceTauren)
		p := req.Raid.Parties[0].Players[0]
		p.Rotation = &proto.APLRotation{}
		p.Consumes.DefaultPotion = proto.Potions_MajorManaPotion
		p.Consumes.DefaultConjured = proto.Conjured_ConjuredDemonicRune
		if key == "bear" {
			p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{
				Options: &proto.FeralTankDruid_Options{},
			}}
		}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		d := sim.Raid.Parties[0].Players[0].(interface{ GetDruid() *druid.Druid }).GetDruid()
		form := map[string]druid.DruidForm{"balance": druid.Moonkin, "feral": druid.Cat, "bear": druid.Bear}[key]
		for _, item := range []int32{13444, 12662} {
			id := core.ActionID{ItemID: item}
			d.SpendMana(sim, d.CurrentMana(), d.NewManaMetrics(id))
			cd := d.GetMajorCooldown(id)
			if cd == nil || !cd.ShouldActivate(sim, &d.Character) {
				t.Fatalf("%s cannot auto-use mana item %d in form", key, item)
			}
			swing := d.AutoAttacks.MainhandSwingAt()
			if !cd.Spell.Cast(sim, &d.Unit) || d.CurrentMana() <= 0 || !d.InForm(form) {
				t.Fatalf("%s mana item %d failed, restored no mana or unshifted", key, item)
			}
			if d.AutoAttacks.MainhandSwingAt() != swing {
				t.Fatal("mana consumable changed the swing timer")
			}
		}
	}
}
