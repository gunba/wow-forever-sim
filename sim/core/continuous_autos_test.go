package core

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestForeverContinuousAutoReadiness(t *testing.T) {
	unit := &Unit{Type: PlayerUnit, Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}}
	unit.AutoAttacks = AutoAttacks{mh: WeaponAttack{unit: unit}, continuousWhileCasting: true}
	unit.Hardcast.Expires = 2 * time.Second
	sim := &Simulation{}
	for _, mask := range []ProcMask{ProcMaskMeleeMHAuto, ProcMaskMeleeOHAuto, ProcMaskRangedAuto, ProcMaskMeleeMHAuto | ProcMaskMeleeMHSpecial} {
		spell := &Spell{Unit: unit, ProcMask: mask}
		if !spell.CanCast(sim, nil) {
			t.Errorf("auto %v was blocked during a cast", mask)
		}
	}
	special := &Spell{Unit: unit, ProcMask: ProcMaskRangedSpecial}
	if special.CanCast(sim, nil) {
		t.Fatal("ordinary special became castable during another cast")
	}
	unit.AutoAttacks.continuousWhileCasting = false
	melee := &Spell{Unit: unit, ProcMask: ProcMaskMeleeMHAuto}
	if melee.CanCast(sim, nil) {
		t.Fatal("caster melee was allowed during a cast")
	}
	unit.Hardcast.Expires = 0
	if melee.CanCast(sim, nil) {
		t.Fatal("caster melee weaving was allowed between casts")
	}
}
