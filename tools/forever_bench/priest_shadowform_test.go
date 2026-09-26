//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/priest"
)

func TestShadowformCostAndRestrictions(t *testing.T) {
	req := racialFixture("shadow", proto.Race_RaceTroll)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	p := sim.Raid.Parties[0].Players[0].(priest.PriestAgent).GetPriest()
	before := p.CurrentMana()
	if !p.Shadowform.Cast(sim, &p.Unit) || math.Abs(before-p.CurrentMana()-.4*p.BaseMana) > 1e-8 {
		t.Fatal("Shadowform did not charge 40% base mana")
	}
	if p.GCD.TimeToReady(sim) != p.ApplyCastSpeed(core.GCDDefault) {
		t.Fatal("Shadowform did not trigger its hasted spell GCD")
	}
	sim.CurrentTime = 2 * time.Second
	if !p.Smite[priest.SmiteRanks].CanCast(sim, p.CurrentTarget) || !p.HolyFire[priest.HolyFireRanks].CanCast(sim, p.CurrentTarget) {
		t.Fatal("Shadowform blocked Holy damage")
	}
	// Healing spells are not part of the supported DPS spellbook. Exercise
	// the registration restriction without introducing an invented heal.
	for _, tc := range []struct {
		name  string
		code  int32
		mask  core.ProcMask
		flags core.SpellFlag
		allow bool
	}{
		{"heal", 0, core.ProcMaskSpellHealing, core.SpellFlagHelpful, false},
		{"Holy Nova", priest.SpellCode_PriestHolyNova, core.ProcMaskSpellDamage, 0, false},
		{"shield or utility", 0, core.ProcMaskEmpty, core.SpellFlagHelpful, true},
		{"passive heal", 0, core.ProcMaskSpellHealing, core.SpellFlagHelpful | core.SpellFlagPassiveSpell, true},
	} {
		spell := p.RegisterSpell(core.SpellConfig{
			ActionID:  core.ActionID{SpellID: 990000, Tag: int32(len(p.Spellbook))},
			SpellCode: tc.code, ProcMask: tc.mask, Flags: tc.flags,
			SpellSchool: core.SpellSchoolHoly,
		})
		if spell.CanCast(sim, &p.Unit) != tc.allow {
			t.Errorf("%s cast permission is wrong", tc.name)
		}
		if tc.allow {
			p.OnCastComplete(sim, spell)
		}
		if !p.ShadowformAura.IsActive() {
			t.Fatal("checking an action cancelled Shadowform")
		}
	}
	p.ShadowformAura.Deactivate(sim)
	if !p.Spellbook[len(p.Spellbook)-4].CanCast(sim, &p.Unit) {
		t.Fatal("healing remained blocked after leaving Shadowform")
	}
}
