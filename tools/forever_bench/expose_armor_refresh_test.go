//go:build with_db

package main

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/rogue"
)

func TestExposeArmorUpgradeChangesReductionAndRestoresArmor(t *testing.T) {
	req := racialFixture("combat", proto.Race_RaceOrc)
	req.Raid.Debuffs = &proto.Debuffs{}
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	s := core.NewSim(req, simsignals.Signals{})
	s.Options.Interactive = true
	s.Reset()
	r := s.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
	target := r.CurrentTarget
	startArmor := target.GetStat(stats.Armor)
	// Isolate the refresh mechanics from avoidance.
	r.ExposeArmor.BonusHitRating = 100 * core.MeleeHitRatingPerHitChance
	target.PseudoStats.CanParry = false
	r.PseudoStats.DodgeReduction = 1
	cp := r.NewComboPointMetrics(core.ActionID{SpellID: 11198})
	for i, points := range []int32{1, 5} {
		s.CurrentTime = time.Duration(i*2) * time.Second
		r.AddEnergy(s, 100, r.NewEnergyMetrics(core.ActionID{SpellID: 11198}))
		r.AddComboPoints(s, points-r.ComboPoints(), target, cp)
		if !r.ExposeArmor.Cast(s, target) {
			t.Fatalf("%d-point Expose Armor did not cast", points)
		}
		want := startArmor - 450*float64(points)
		if got := target.GetStat(stats.Armor); got != want {
			t.Fatalf("%d-point Expose Armor: armor %g, want %g", points, got, want)
		}
	}
	r.ExposeArmorAuras.Get(target).Deactivate(s)
	if got := target.GetStat(stats.Armor); got != startArmor {
		t.Fatalf("Expose Armor expiry left armor %g, want %g", got, startArmor)
	}
}
