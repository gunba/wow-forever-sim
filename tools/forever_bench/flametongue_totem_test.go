//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"google.golang.org/protobuf/encoding/protojson"
)

func flametongueFixture(t *testing.T, windfury bool) (*core.Simulation, *core.Character, *core.Aura, *core.Spell) {
	t.Helper()
	req := racialFixture("retribution", proto.Race_RaceHuman)
	p := req.Raid.Parties[0].Players[0]
	p.ForeverTier1Bonuses = false
	p.Consumes = &proto.Consumes{}
	p.Buffs = &proto.IndividualBuffs{}
	p.Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	req.Raid.Buffs = &proto.RaidBuffs{}
	req.Raid.Debuffs = &proto.Debuffs{}
	req.Raid.Parties[0].Buffs = &proto.PartyBuffs{FlametongueTotem: true, WindfuryTotem: windfury}
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	proc := c.GetSpell(core.ActionID{SpellID: 16389})
	if proc == nil {
		t.Fatal("Flametongue damage proc was not registered")
	}
	proc.BonusHitRating = 10000
	proc.BonusCritRating = -10000
	return sim, c, c.GetAura("Flametongue Totem (Rank 4; provisional)"), proc
}

func TestFlametongueTotemProvisionalSpeedAndZeroSP(t *testing.T) {
	for _, speed := range []float64{1.3, 2.6, 4} {
		var reference float64
		for _, sp := range []float64{0, 1000} {
			sim, c, aura, proc := flametongueFixture(t, false)
			c.AutoAttacks.MH().SwingSpeed = speed
			c.AddStatsDynamic(sim, stats.Stats{stats.SpellPower: sp, stats.FirePower: sp})
			target := c.CurrentTarget
			aura.OnSpellHitDealt(aura, sim, c.AutoAttacks.MHAuto(), &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
			damage := proc.SpellMetrics[target.UnitIndex].TotalDamage
			if proc.BonusCoefficient != 0 || math.Abs(damage-13.63*speed) > 1e-6 {
				t.Fatalf("speed %v SP %v: damage %v coefficient %v", speed, sp, damage, proc.BonusCoefficient)
			}
			if sp == 0 {
				reference = damage
			} else if damage != reference {
				t.Fatalf("spell power changed proc damage: %v versus %v", damage, reference)
			}
		}
	}
}

func TestFlametongueTotemProcEligibilityAndConflicts(t *testing.T) {
	for _, tc := range []struct {
		name     string
		mask     core.ProcMask
		outcome  core.HitOutcome
		flags    core.SpellFlag
		windfury bool
		enchant  int32
		want     int32
	}{
		{"main_hand", core.ProcMaskMeleeMHAuto, core.OutcomeHit, 0, false, 0, 1},
		{"queued_main_hand", core.ProcMaskMeleeMHAuto | core.ProcMaskMeleeMHSpecial, core.OutcomeHit, 0, false, 0, 1},
		{"off_hand", core.ProcMaskMeleeOHAuto, core.OutcomeHit, 0, false, 0, 0},
		{"miss", core.ProcMaskMeleeMHAuto, core.OutcomeMiss, 0, false, 0, 0},
		{"spell", core.ProcMaskSpellDamage, core.OutcomeHit, 0, false, 0, 0},
		{"suppressed", core.ProcMaskMeleeMHAuto, core.OutcomeHit, core.SpellFlagSuppressWeaponProcs, false, 0, 0},
		{"windfury", core.ProcMaskMeleeMHAuto, core.OutcomeHit, 0, true, 0, 0},
		{"flametongue_weapon", core.ProcMaskMeleeMHAuto, core.OutcomeHit, 0, false, 1666, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sim, c, aura, proc := flametongueFixture(t, tc.windfury)
			c.MainHand().TempEnchant = tc.enchant
			aura.OnSpellHitDealt(aura, sim, &core.Spell{ProcMask: tc.mask, Flags: tc.flags}, &core.SpellResult{Target: c.CurrentTarget, Outcome: tc.outcome})
			if got := proc.SpellMetrics[c.CurrentTarget.UnitIndex].Casts; got != tc.want {
				t.Fatalf("proc casts %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFlametongueTotemRequestRoundTrip(t *testing.T) {
	original := &proto.PartyBuffs{FlametongueTotem: true, GraceOfAirTotem: proto.TristateEffect_TristateEffectRegular}
	data, err := protojson.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	restored := &proto.PartyBuffs{}
	if err := protojson.Unmarshal(data, restored); err != nil {
		t.Fatal(err)
	}
	if !restored.FlametongueTotem || restored.GraceOfAirTotem != original.GraceOfAirTotem || restored.WindfuryTotem {
		t.Fatalf("request lost legal Flametongue/Grace combination: %s", data)
	}
}
