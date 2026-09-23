//go:build with_db

package main

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/rogue"
)

func TestVenomCanRefreshAfterExpiration(t *testing.T) {
	req := racialFixture("mutilate", proto.Race_RaceOrc)
	player := req.Raid.Parties[0].Players[0]
	player.TalentsString = "00530310341401051-302303202014"
	player.Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	r := sim.Raid.Parties[0].Players[0].(interface{ GetRogue() *rogue.Rogue }).GetRogue()
	if r.Venom == nil {
		t.Fatal("Venom was not registered")
	}
	metrics := r.NewComboPointMetrics(core.ActionID{SpellID: 1310703})
	sim.CurrentTime = 0
	r.AddComboPoints(sim, 3, r.CurrentTarget, metrics)
	if !r.Venom.Cast(sim, r.CurrentTarget) {
		t.Fatal("initial Venom cast failed")
	}
	sim.CurrentTime = 20 * time.Second
	r.VenomAura.Deactivate(sim)
	r.AddComboPoints(sim, 3, r.CurrentTarget, metrics)
	if !r.Venom.Cast(sim, r.CurrentTarget) {
		t.Fatal("Venom could not be reapplied after expiration")
	}
}
