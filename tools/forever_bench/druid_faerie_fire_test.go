//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/druid"
)

func TestForeverFaerieFireForms(t *testing.T) {
	for _, form := range []druid.DruidForm{druid.Humanoid, druid.Moonkin, druid.Cat, druid.Bear} {
		key := "balance"
		if form.Matches(druid.Cat | druid.Bear) {
			key = "feral"
		}
		req := racialFixture(key, proto.Race_RaceTauren)
		p := req.Raid.Parties[0].Players[0]
		p.Rotation = &proto.APLRotation{}
		p.Equipment = &proto.EquipmentSpec{}
		p.Consumes = &proto.Consumes{}
		p.ForeverTier1Bonuses = false
		req.Raid.Debuffs = &proto.Debuffs{}
		if form == druid.Bear {
			p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{
				Options: &proto.FeralTankDruid_Options{},
			}}
		}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		d := sim.Raid.Parties[0].Players[0].(interface{ GetDruid() *druid.Druid }).GetDruid()
		if form == druid.Humanoid {
			d.CancelShapeshift(sim)
		}
		ff := d.FaerieFire
		if ff == nil || ff.SpellID != 9907 || d.GetSpell(core.ActionID{SpellID: 17392}) != nil {
			t.Fatal("Faerie Fire did not replace the obsolete Feral action")
		}
		if ff.Cost.GetCurrentCost() != 115 || ff.DefaultCast.GCD != core.GCDDefault || ff.CD.Duration != 0 {
			t.Fatalf("incorrect Faerie Fire cost/timing for form %v", form)
		}
		mana := d.CurrentMana()
		swing := d.AutoAttacks.MainhandSwingAt()
		energy := 0.0
		if d.HasEnergyBar() {
			energy = d.CurrentEnergy()
		}
		if !ff.Cast(sim, sim.Encounter.TargetUnits[0]) || !d.InForm(form) {
			t.Fatalf("Faerie Fire failed or changed form %v", form)
		}
		if math.Abs(mana-d.CurrentMana()-115) > 1e-7 {
			t.Fatalf("form %v did not pay the mana cost", form)
		}
		if d.HasEnergyBar() && d.CurrentEnergy() != energy {
			t.Fatal("Faerie Fire spent Energy rather than mana")
		}
		if form.Matches(druid.Cat|druid.Bear) && d.AutoAttacks.MainhandSwingAt() != swing {
			t.Fatal("Faerie Fire reset a form autoattack")
		}
	}
}
