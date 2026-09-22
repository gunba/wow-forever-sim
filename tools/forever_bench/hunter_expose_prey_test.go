//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/hunter"
)

func TestExposePreyAttackMask(t *testing.T) {
	req := racialFixture("survival", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	req.Raid.Debuffs.HuntersMark = proto.TristateEffect_TristateEffectRegular
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	h := sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
	if h.Talents.ExposePrey != 2 {
		t.Fatal("fixture requires 2/2 Expose Prey")
	}
	trigger := h.GetAura("Expose Prey")
	target := sim.Encounter.TargetUnits[0]
	for _, tc := range []struct {
		name string
		mask core.ProcMask
		hit  core.HitOutcome
		want bool
	}{
		{"melee auto", core.ProcMaskMeleeMHAuto, core.OutcomeHit, true},
		{"ranged auto", core.ProcMaskRangedAuto, core.OutcomeHit, true},
		{"melee special", core.ProcMaskMeleeMHSpecial, core.OutcomeHit, true},
		{"ranged special", core.ProcMaskRangedSpecial, core.OutcomeHit, true},
		{"trap spell", core.ProcMaskSpellDamage, core.OutcomeHit, false},
		{"proc", core.ProcMaskSpellDamageProc, core.OutcomeHit, false},
		{"miss", core.ProcMaskMeleeMHAuto, core.OutcomeMiss, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			procs := 0
			for i := 0; i < 500; i++ {
				h.DefensiveState.Deactivate(sim)
				trigger.OnSpellHitDealt(trigger, sim, &core.Spell{ProcMask: tc.mask}, &core.SpellResult{Target: target, Outcome: tc.hit})
				if h.DefensiveState.IsActive() {
					procs++
				}
			}
			if (procs > 0) != tc.want {
				t.Fatalf("procs = %d; eligible = %v", procs, tc.want)
			}
		})
	}
}
