//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/warlock"
	googleProto "google.golang.org/protobuf/proto"
)

func TestAgonyRampRetainsEachTargetsAmplifyBonus(t *testing.T) {
	req := warlockEffectFixture(map[string]int{"amplifyCurse": 1})
	req.Encounter.Targets = append(req.Encounter.Targets, googleProto.Clone(req.Encounter.Targets[0]).(*proto.Target))
	s := core.NewSim(req, simsignals.Signals{})
	s.Options.Interactive = true
	s.Reset()
	w := s.Raid.Parties[0].Players[0].(warlock.WarlockAgent).GetWarlock()
	spell := w.BaneOfAgony[len(w.BaneOfAgony)-1]
	targets := s.Encounter.TargetUnits
	dots := []*core.Dot{spell.Dot(targets[0]), spell.Dot(targets[1])}
	w.AmplifyCurseAura.Activate(s)
	dots[0].Apply(s)
	first := dots[0].SnapshotRawBaseDamage
	dots[1].Apply(s)
	second := dots[1].SnapshotRawBaseDamage
	if first != second*1.5 {
		t.Fatalf("fixture lost Amplify Curse: %g / %g", first, second)
	}
	for i, base := range []float64{first, second} {
		dots[i].TickCount = 4
		dots[i].OnTick(s, targets[i], dots[i])
		if got := dots[i].SnapshotRawBaseDamage; got != base*2 {
			t.Errorf("target %d middle ramp = %g, want %g", i, got, base*2)
		}
		dots[i].TickCount = 8
		dots[i].OnTick(s, targets[i], dots[i])
		if got := dots[i].SnapshotRawBaseDamage; got != base*3 {
			t.Errorf("target %d final ramp = %g, want %g", i, got, base*3)
		}
	}
}
