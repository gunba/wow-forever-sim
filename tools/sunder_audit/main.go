// Sunder stack coverage from exact simulation requests. Observing an aura does
// not add actions, consume randomness or modify its armor-reduction effect.
package main

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"google.golang.org/protobuf/encoding/protojson"
)

type coverage struct {
	FirstFiveSeconds, FullStackSeconds, BelowFiveSecondsAfterRamp float64
	Drops, NeverFive, Iterations, FightsWithDrop                  int
}

func measure(request *proto.RaidSimRequest) coverage {
	if request.SimOptions.Iterations <= 0 {
		panic("positive iteration count required")
	}
	s := core.NewSim(request, simsignals.Signals{})
	a := s.Encounter.TargetUnits[0].GetAura("Sunder Armor")
	if a == nil {
		panic("request must register Sunder Armor on the main target")
	}
	original := a.OnStacksChange
	result := coverage{Iterations: int(request.SimOptions.Iterations)}
	var first, last, full time.Duration
	drops, hadFive := 0, false
	a.OnStacksChange = func(a *core.Aura, s *core.Simulation, before, after int32) {
		original(a, s, before, after)
		now := min(s.CurrentTime, s.Duration)
		if before == 5 {
			full += now - last
		}
		last = now
		// A natural expiration exactly at fight end completes the previous
		// full-stack interval; it is not a maintenance failure.
		if now >= s.Duration {
			return
		}
		if after == 5 && !hadFive {
			first, hadFive = now, true
		}
		if before == 5 {
			drops++
		}
	}
	for i := 0; i < result.Iterations; i++ {
		first, last, full, hadFive, drops = 0, 0, 0, false, 0
		s.Reseed(int64(i))
		s.Reset()
		s.PrePull()
		for !s.Step() {
		}
		if a.GetStacks() == 5 {
			full += s.Duration - last
		}
		if !hadFive {
			result.NeverFive++
		} else {
			result.FirstFiveSeconds += first.Seconds()
			result.BelowFiveSecondsAfterRamp += (s.Duration - first - full).Seconds()
		}
		result.FullStackSeconds += full.Seconds()
		result.Drops += drops
		if drops > 0 {
			result.FightsWithDrop++
		}
		s.Cleanup()
	}
	n := float64(result.Iterations)
	result.FirstFiveSeconds /= n
	result.FullStackSeconds /= n
	result.BelowFiveSecondsAfterRamp /= n
	return result
}

func main() {
	input := flag.String("request", "", "exact RaidSimRequest JSON")
	output := flag.String("output", "", "coverage JSON output")
	flag.Parse()
	sim.RegisterAll()
	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	request := &proto.RaidSimRequest{}
	if err = protojson.Unmarshal(data, request); err != nil {
		panic(err)
	}
	data, err = json.MarshalIndent(measure(request), "", "  ")
	if err != nil {
		panic(err)
	}
	if err = os.WriteFile(*output, append(data, '\n'), 0644); err != nil {
		panic(err)
	}
}
