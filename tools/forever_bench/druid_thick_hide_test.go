//go:build with_db

package main

import (
	"math"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/druid"
)

func TestThickHideFollowsFormAndDefense(t *testing.T) {
	for _, form := range []druid.DruidForm{druid.Cat, druid.Bear, druid.Moonkin} {
		key, multiplier := "feral", 1.0
		if form == druid.Moonkin {
			key, multiplier = "balance", druid.MoonkinFormArmorMultiplier
		} else if form == druid.Bear {
			multiplier = druid.BearFormArmorMultiplier
		}
		req := racialFixture(key, proto.Race_RaceTauren)
		p := req.Raid.Parties[0].Players[0]
		p.Rotation = &proto.APLRotation{}
		p.Equipment = &proto.EquipmentSpec{}
		p.ForeverTier1Bonuses = false
		p.TalentsString = "-500003" // five Ferocity, three Thick Hide
		if form == druid.Moonkin {
			for _, b := range builds() {
				if b.Key == "balance" {
					p.TalentsString = strings.Split(b.player(p.Race).TalentsString, "-")[0] + "-500003"
				}
			}
		} else if form == druid.Bear {
			p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{Options: &proto.FeralTankDruid_Options{}}}
		}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		d := sim.Raid.Parties[0].Players[0].(interface{ GetDruid() *druid.Druid }).GetDruid()
		aura := d.CatFormAura
		if form == druid.Bear {
			aura = d.BearFormAura
		} else if form == druid.Moonkin {
			aura = d.MoonkinFormAura
		}
		formed := d.GetStat(stats.Armor)
		d.CancelShapeshift(sim)
		humanoid := d.GetStat(stats.Armor)
		if math.Abs(formed-humanoid-180*multiplier) > 1e-7 {
			t.Fatalf("form %v retained or mis-scaled Thick Hide armor", form)
		}
		d.AddStatDynamic(sim, stats.Defense, 10)
		if d.GetStat(stats.Armor) != humanoid {
			t.Fatal("Defense added Thick Hide armor in humanoid form")
		}
		aura.Activate(sim)
		// This tests dependency updates, not validation of the provisional coefficient.
		if math.Abs(d.GetStat(stats.Armor)-formed-10*.67*3*multiplier) > 1e-7 {
			t.Fatal("form did not use the current Defense stat")
		}
		d.AddStatDynamic(sim, stats.Defense, 10)
		if math.Abs(d.GetStat(stats.Armor)-formed-20*.67*3*multiplier) > 1e-7 {
			t.Fatal("active form ignored a dynamic Defense change")
		}
		d.CancelShapeshift(sim)
		if math.Abs(d.GetStat(stats.Armor)-humanoid) > 1e-7 {
			t.Fatal("form exit leaked Thick Hide armor")
		}
	}
}
