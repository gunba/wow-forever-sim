//go:build with_db

package main

import (
	"encoding/json"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"math"
	"testing"
	"time"
)

func TestPaladinOctober9FuryExplicitLifetime(t *testing.T) {
	for _, seconds := range []float64{0, 5, 30, 60} {
		sim, p := paladinOctober9Configured(t, nil, nil, func(player *proto.Player) {
			o := player.GetProtectionPaladin().Options
			o.PrimarySeal = proto.PaladinSeal_Fury
			o.SealOfFuryShieldDurationSeconds = &seconds
		})
		action := p.NewAPLAction(nil, &proto.APLAction{Action: &proto.APLAction_CastPaladinPrimarySeal{CastPaladinPrimarySeal: &proto.APLActionCastPaladinPrimarySeal{}}})
		action.Execute(sim)
		if p.GetSpell(core.ActionID{SpellID: 20423}).SpellMetrics[p.CurrentTarget.UnitIndex].Casts != 1 {
			t.Fatal("Fury primary not selectable")
		}
		proc := p.GetSpell(core.ActionID{SpellID: 20418})
		proc.BonusCritRating = -100000
		proc.ApplyEffects(sim, p.CurrentTarget, proc)
		shield := p.GetSpell(core.ActionID{SpellID: 20423, Tag: 1}).SelfShield()
		if seconds == 0 {
			if shield.IsActive() || shield.RemainingAbsorb() != 0 {
				t.Fatal("zero scenario absorbed")
			}
			if proc.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage == 0 {
				t.Fatal("zerodisabled sourceddamage")
			}
		} else if !shield.IsActive() || shield.ExpiresAt() != time.Duration(seconds*float64(time.Second)) {
			t.Fatal("configuredlifetimeignored")
		}
	}
	for _, value := range []float64{-1, math.Inf(1), math.NaN(), 3601} {
		func() {
			defer func() {
				if recover() == nil {
					t.Fatal("invalidmodelaccepted")
				}
			}()
			paladinOctober9Configured(t, nil, nil, func(player *proto.Player) {
				player.GetProtectionPaladin().Options.SealOfFuryShieldDurationSeconds = &value
			})
		}()
	}
	zero := 0.0
	o := &proto.PaladinOptions{PrimarySeal: proto.PaladinSeal_Fury, SealOfFuryShieldDurationSeconds: &zero}
	raw, err := protojson.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}
	var back proto.PaladinOptions
	if err = protojson.Unmarshal(raw, &back); err != nil || back.SealOfFuryShieldDurationSeconds == nil || *back.SealOfFuryShieldDurationSeconds != 0 || back.PrimarySeal != proto.PaladinSeal_Fury {
		t.Fatal("explicitzero model lost")
	}
	var fields map[string]any
	json.Unmarshal(raw, &fields)
	if fields["sealOfFuryShieldDurationSeconds"] != float64(0) {
		t.Fatal("zero missing in export")
	}
}
func TestPaladinOctober9FuryEligibilityAndShieldRequirement(t *testing.T) {
	sim, p := paladinOctober9Fixture(t, nil, nil)
	seal := p.GetSpell(core.ActionID{SpellID: 20423})
	seal.ApplyEffects(sim, p.CurrentTarget, seal)
	aura := p.GetAura("Seal of Fury" + p.Label + "7")
	proc := p.GetSpell(core.ActionID{SpellID: 20418})
	before := proc.SpellMetrics[p.CurrentTarget.UnitIndex].Casts
	special := p.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 991}, ProcMask: core.ProcMaskMeleeMHSpecial})
	white := p.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 992}, ProcMask: core.ProcMaskMeleeMHAuto})
	for _, event := range []struct {
		s *core.Spell
		o core.HitOutcome
	}{{special, core.OutcomeHit}, {white, core.OutcomeMiss}} {
		aura.OnSpellHitDealt(aura, sim, event.s, &core.SpellResult{Target: p.CurrentTarget, Outcome: event.o})
	}
	if proc.SpellMetrics[p.CurrentTarget.UnitIndex].Casts != before {
		t.Fatal("nonwhite/missproc")
	}
	aura.OnSpellHitDealt(aura, sim, white, &core.SpellResult{Target: p.CurrentTarget, Outcome: core.OutcomeHit})
	if proc.SpellMetrics[p.CurrentTarget.UnitIndex].Casts != before+1 {
		t.Fatal("whiteprocnot100percent")
	}
	pool := p.GetSpell(core.ActionID{SpellID: 20423, Tag: 1}).SelfShield()
	pool.Deactivate(sim)
	p.Equipment[proto.ItemSlot_ItemSlotOffHand] = core.Item{}
	beforeDamage := proc.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
	proc.ApplyEffects(sim, p.CurrentTarget, proc)
	if pool.IsActive() || proc.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage <= beforeDamage {
		t.Fatal("shieldrequirementchangedproc damage/absorb")
	}
}
