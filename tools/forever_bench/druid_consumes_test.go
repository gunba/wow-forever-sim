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

func TestDemonicRuneSelfHitIsNotOutgoingDPS(t *testing.T) {
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
		if action.Targets[1].Damage != 0 || action.Targets[1].ResistedDamage != 0 {
			t.Fatalf("self-hit entered outgoing DPS metrics: %+v", action.Targets[1])
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

func TestConfirmedExplosivesRespectRangeAndCatForm(t *testing.T) {
	for _, tc := range []struct {
		key, race string
		grenade   bool
		sapper    bool
	}{
		{"feral", "Tauren", false, false},
		{"beast_mastery", "Orc", true, false},
		{"marksmanship", "Orc", true, false},
		{"fire", "Gnome", false, false},
		{"combat", "Orc", true, true},
		{"enhancement", "Orc", true, true},
	} {
		for _, b := range builds() {
			if b.Key != tc.key {
				continue
			}
			var race proto.Race
			for _, candidate := range b.races() {
				if raceName(candidate) == tc.race {
					race = candidate
				}
			}
			if race == proto.Race_RaceUnknown {
				t.Fatalf("%s/%s unavailable", tc.key, tc.race)
			}
			c := b.consumes(race)
			if (c.FillerExplosive == proto.Explosive_ExplosiveThoriumGrenade) != tc.grenade ||
				(c.SapperExplosive == proto.SapperExplosive_SapperGoblinSapper) != tc.sapper {
				t.Fatalf("%s/%s: unexpected explosives: %+v", tc.key, tc.race, c)
			}
			break
		}
	}
}
