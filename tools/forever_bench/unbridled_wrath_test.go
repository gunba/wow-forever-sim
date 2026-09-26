//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/warrior"
)

func TestForeverUnbridledWrathExcludesQueuedSpecials(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	s := core.NewSim(req, simsignals.Signals{})
	s.Options.Interactive = true
	s.Reset()
	w := s.Raid.Parties[0].Players[0].(warrior.WarriorAgent).GetWarrior()
	aura := w.GetAura("Unbridled Wrath")
	if aura == nil || w.Talents.UnbridledWrath == 0 {
		t.Fatal("fixture must have Unbridled Wrath")
	}
	for _, tc := range []struct {
		name    string
		spell   *core.Spell
		outcome core.HitOutcome
		want    bool
	}{
		{"main hand", w.AutoAttacks.MHAuto(), core.OutcomeHit, true},
		{"off hand", w.AutoAttacks.OHAuto(), core.OutcomeHit, true},
		{"miss", w.AutoAttacks.MHAuto(), core.OutcomeMiss, false},
		{"Heroic Strike", w.HeroicStrike.Spell, core.OutcomeHit, false},
		{"Cleave", w.Cleave.Spell, core.OutcomeCrit, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w.SpendRage(s, w.CurrentRage(), w.NewRageMetrics(core.ActionID{SpellID: 12964}))
			hit := &core.SpellResult{Target: s.Encounter.TargetUnits[0], Outcome: tc.outcome}
			for i := 0; i < 200; i++ {
				aura.OnSpellHitDealt(aura, s, tc.spell, hit)
			}
			if got := w.CurrentRage() > 0; got != tc.want {
				t.Fatalf("Unbridled Wrath generated %g rage; eligible=%t", w.CurrentRage(), tc.want)
			}
		})
	}
}
