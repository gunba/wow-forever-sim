//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/druid"
)

// Shifting into and out of Bear Form during a fight keeps the fraction of maximum health. A build without Heart of
// the Wild once crashed here: Bear Form's entry gained a rounding-sized negative amount of health.
func TestBearFormKeepsHealthFraction(t *testing.T) {
	for _, talents := range []string{"", "55322021101011--5542322120001"} {
		req := racialFixture("feral", proto.Race_RaceTauren)
		p := req.Raid.Parties[0].Players[0]
		p.Rotation = &proto.APLRotation{}
		if talents != "" {
			p.TalentsString = talents
		}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		d := sim.Raid.Parties[0].Players[0].(interface{ GetDruid() *druid.Druid }).GetDruid()
		sim.Reset()
		sim.CurrentTime = time.Second
		d.CancelShapeshift(sim)
		for pct := 5; pct <= 100; pct++ {
			d.GainHealth(sim, d.MaxHealth()-d.CurrentHealth(), d.DamageTakenHealthMetrics)
			d.RemoveHealth(sim, (1-float64(pct)/100)*d.MaxHealth())
			want := d.CurrentHealthPercent()
			d.BearFormAura.Activate(sim)
			if got := d.CurrentHealthPercent(); math.Abs(got-want) > 1e-12 {
				t.Fatalf("talents %q at %d%%: entering Bear Form left %v of maximum health, want %v", p.TalentsString, pct, got, want)
			}
			d.CancelShapeshift(sim)
			if got := d.CurrentHealthPercent(); math.Abs(got-want) > 1e-12 {
				t.Fatalf("talents %q at %d%%: leaving Bear Form left %v of maximum health, want %v", p.TalentsString, pct, got, want)
			}
		}
	}
}
