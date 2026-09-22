package druid

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestFurorExitTimesAndIterationReset(t *testing.T) {
	d := &Druid{Talents: &proto.DruidTalents{Furor: 5}}
	d.EnableEnergyBar(100)
	for _, exitAt := range []time.Duration{-2 * time.Second, 0, time.Second} {
		d.lastCatFormEnergy, d.lastCatFormExitAt = 20, exitAt
		sim := &core.Simulation{CurrentTime: exitAt + time.Second}
		if got := d.furorShiftEnergy(sim); got != 30 {
			t.Fatalf("exit at %v: got %v Energy, want 20 carried + 10 regenerated", exitAt, got)
		}
		d.Reset(sim)
		if got := d.furorShiftEnergy(sim); got != 0 {
			t.Fatalf("iteration reset retained Furor history: %v", got)
		}
	}
}
