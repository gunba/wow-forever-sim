//go:build with_db

package main

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"math"
	"testing"
	"time"
)

func TestPaladinOctober9FuryProcPoolAndTalentMana(t *testing.T) {
	for _, talented := range []bool{false, true} {
		points := map[string]int{}
		if talented {
			points["improvedSealOfFury"] = 1
		}
		sim, p := paladinOctober9Fixture(t, points, nil)
		proc := p.GetSpell(core.ActionID{SpellID: 20418})
		pool := p.GetSpell(core.ActionID{SpellID: 20423, Tag: 1}).SelfShield()
		proc.BonusCritRating = -100000
		before := proc.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
		proc.ApplyEffects(sim, p.CurrentTarget, proc)
		damage := proc.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage - before
		if math.Abs(damage-(35+.10000000149*proc.GetBonusDamage(p.CurrentTarget))) > 1e-8 || math.Abs(pool.RemainingAbsorb()-damage*.5) > 1e-8 || pool.ExpiresAt() != 30*time.Second {
			t.Fatalf("damage%g pool%g expires%s", damage, pool.RemainingAbsorb(), pool.ExpiresAt())
		}
		low := p.GetSpell(core.ActionID{SpellID: 1311647})
		low.BonusCritRating = -100000
		low.ApplyEffects(sim, p.CurrentTarget, low)
		if math.Abs(pool.RemainingAbsorb()-.5*(6+.10000000149*low.GetBonusDamage(p.CurrentTarget))) > 1e-8 {
			t.Fatal("rank pools stacked ratherthanreplaced")
		}
		p.SpendMana(sim, 1000, p.NewManaMetrics(core.ActionID{SpellID: 4}))
		beforeMana := p.CurrentMana()
		incoming := p.CurrentTarget.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 5}, ProcMask: core.ProcMaskEmpty, SpellSchool: core.SpellSchoolHoly, DamageMultiplier: 1, ThreatMultiplier: 1})
		p.CurrentTarget.Level = 65
		selected := p.CurrentTarget
		p.CurrentTarget = &p.Unit
		incoming.CalcAndDealDamage(sim, &p.Unit, 10, incoming.OutcomeAlwaysHit)
		p.CurrentTarget = selected
		want := 0.0
		if talented {
			want = 87
		}
		if math.Abs(p.CurrentMana()-beforeMana-want) > 1e-8 {
			t.Fatalf("talented%v mana%g want%g", talented, p.CurrentMana()-beforeMana, want)
		}
		if pool.IsActive() {
			t.Fatal("depletedpoolactive")
		}
		pool.Duration = 5 * time.Second
		low.ApplyEffects(sim, p.CurrentTarget, low)
		if pool.ExpiresAt() != 5*time.Second {
			t.Fatal("controlledlifetimesensitivityignored")
		}
		beforeMana = p.CurrentMana()
		sim.CurrentTime = 5 * time.Second
		incoming.CalcAndDealDamage(sim, &p.Unit, 3, incoming.OutcomeAlwaysHit)
		if p.CurrentMana() != beforeMana {
			t.Fatal("expirygrantedmana")
		}
		sim.Cleanup()
		manaGain := p.GetSpell(core.ActionID{OtherID: proto.OtherAction_OtherActionManaGain})
		for _, m := range manaGain.SpellMetrics {
			if m.TotalThreat != 0 {
				t.Fatal("Fury mana generated threat")
			}
		}
		sim.Reset()
		if pool.IsActive() || pool.RemainingAbsorb() != 0 {
			t.Fatal("poolstalereset")
		}
	}
}
func TestPaladinOctober9FuryEchoAndJudgement(t *testing.T) {
	sim, p := paladinOctober9Fixture(t, map[string]int{"twistOfLight": 1}, nil)
	fury := p.GetSpell(core.ActionID{SpellID: 20423})
	fury.ApplyEffects(sim, p.CurrentTarget, fury)
	judge := p.GetSpell(core.ActionID{SpellID: 20414})
	judge.BonusHitRating = 100000
	judge.BonusCritRating = -100000
	judge.ApplyEffects(sim, p.CurrentTarget, judge)
	taunt := p.CurrentTarget.GetAura("Judgement of Fury Taunt-0")
	if taunt == nil || !taunt.IsActive() || taunt.ExpiresAt() != 4*time.Second {
		t.Fatal("known4stauntmissing")
	}
	pool := p.GetSpell(core.ActionID{SpellID: 20423, Tag: 1}).SelfShield()
	if pool.IsActive() {
		t.Fatal("inventedjudgementshield")
	}
	sor := p.GetSpell(core.ActionID{SpellID: 20293})
	sor.ApplyEffects(sim, p.CurrentTarget, sor)
	echo := p.GetAura("Echo of Fury")
	if !echo.IsActive() || echo.GetStacks() != 1 {
		t.Fatal("Furynotbanked")
	}
	judge.ApplyEffects(sim, p.CurrentTarget, judge)
	if !echo.IsActive() {
		t.Fatal("judgementconsumedEcho")
	}
	white := p.GetSpell(core.ActionID{OtherID: proto.OtherAction_OtherActionAttack})
	if white == nil {
		white = p.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 7}, ProcMask: core.ProcMaskMeleeMHAuto})
	}
	echo.OnSpellHitDealt(echo, sim, white, &core.SpellResult{Target: p.CurrentTarget, Outcome: core.OutcomeMiss})
	if !echo.IsActive() {
		t.Fatal("missconsumedEcho")
	}
	echo.OnSpellHitDealt(echo, sim, white, &core.SpellResult{Target: p.CurrentTarget, Outcome: core.OutcomeHit})
	if echo.IsActive() || !pool.IsActive() {
		t.Fatal("Echoeffects/consumptionmissing")
	}
	sim.Cleanup()
	sim.Reset()
	if echo.IsActive() {
		t.Fatal("Echoreset")
	}
}
func TestPaladinOctober9FuryDamageCritSPOnce(t *testing.T) {
	sim, p := paladinOctober9Fixture(t, map[string]int{"improvedSeals": 5}, nil)
	p.AddStatDynamic(sim, stats.SpellPower, 100)
	proc := p.GetSpell(core.ActionID{SpellID: 20418})
	proc.BonusCritRating = 100000
	before := proc.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
	proc.ApplyEffects(sim, p.CurrentTarget, proc)
	damage := proc.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage - before
	want := (35 + .10000000149*proc.GetBonusDamage(p.CurrentTarget)) * 1.25 * 2
	if math.Abs(damage-want) > 1e-7 {
		t.Fatalf("critSPdamage%g want%g", damage, want)
	}
	pool := p.GetSpell(core.ActionID{SpellID: 20423, Tag: 1}).SelfShield()
	if math.Abs(pool.RemainingAbsorb()-damage*.5) > 1e-8 {
		t.Fatal("absorbdoublemodified")
	}
}
