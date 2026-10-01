package core

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestFiniteShieldDelivery(t *testing.T) {
	env := &Environment{Ruleset: proto.Ruleset_RulesetForever}
	target := &Unit{Env: env, Type: PlayerUnit, UnitIndex: 0, PseudoStats: stats.NewPseudoStats(), Metrics: NewUnitMetrics()}
	enemy := &Unit{Env: env, Type: EnemyUnit, UnitIndex: 1, PseudoStats: stats.NewPseudoStats(), Metrics: NewUnitMetrics()}
	env.AllUnits = []*Unit{target, enemy}
	sim := &Simulation{Environment: env, Duration: 30 * time.Second}
	shieldSpell := target.RegisterSpell(SpellConfig{
		ActionID: ActionID{SpellID: 1291097}, DamageMultiplier: 1, ProcMask: ProcMaskEmpty, SpellSchool: SpellSchoolPhysical,
		Shield: ShieldConfig{SelfOnly: true, Aura: Aura{Label: "Finite shield regression", Duration: 15 * time.Second}},
	})
	incoming := enemy.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 1}, SpellSchool: SpellSchoolShadow, ProcMask: ProcMaskSpellDamage, Flags: SpellFlagNoOnDamageDealt, DamageMultiplier: 1, ThreatMultiplier: 1})
	shieldSpell.finalize()
	incoming.finalize()
	shield := shieldSpell.SelfShield()
	hit := func(amount float64, periodic bool) float64 {
		result := incoming.NewResult(target)
		result.Outcome = OutcomeHit
		result.Damage = amount
		if periodic {
			incoming.DealPeriodicDamage(sim, result)
		} else {
			incoming.DealDamage(sim, result)
		}
		return result.Damage
	}
	shield.Apply(sim, 450)
	if got := hit(100, false); got != 0 {
		t.Fatalf("shield let %v damage through a 450 pool", got)
	}
	if got := hit(300, true); got != 0 || !shield.IsActive() {
		t.Fatal("periodic hit did not leave 50 capacity")
	}
	if got := hit(100, false); got != 50 || shield.IsActive() {
		t.Fatal("depletion did not absorb only remaining 50")
	}
	shield.Apply(sim, 450)
	hit(100, false)
	shield.Apply(sim, 200)
	if got := hit(250, false); got != 50 {
		t.Fatal("overwrite accumulated old capacity")
	}
	shield.Apply(sim, 450)
	shield.Deactivate(sim)
	if got := hit(100, false); got != 100 {
		t.Fatal("expired shield still absorbed")
	}
	metrics := shieldSpell.SpellMetrics[0]
	if metrics.TotalShielding != 1550 || metrics.TotalAbsorbedShielding != 750 {
		t.Fatalf("capacity/absorption accounting %v / %v", metrics.TotalShielding, metrics.TotalAbsorbedShielding)
	}
	target.Metrics.addSpellMetrics(shieldSpell, shieldSpell.ActionID, shieldSpell.SpellMetrics)
	raw := target.Metrics.actions[shieldSpell.ActionID].Targets[0].ToProto(0)
	if raw.Shielding != 1550 || raw.AbsorbedShielding != 750 || raw.UnusedShielding != 800 || target.Metrics.hps.Total != 750 || target.Metrics.dps.Total != 0 {
		t.Fatal("generated/unused shielding was credited as effective healing or damage")
	}
	shield.Apply(sim, 450)
	sim.CurrentTime = 15 * time.Second
	if got := hit(100, false); got != 100 || shield.IsActive() {
		t.Fatal("timed expiry kept absorbing")
	}
	second := target.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 1291099}, DamageMultiplier: 1, ProcMask: ProcMaskEmpty, SpellSchool: SpellSchoolPhysical, Shield: ShieldConfig{SelfOnly: true, Aura: Aura{Label: "Second finite shield", Duration: 15 * time.Second}}})
	second.finalize()
	shield.Apply(sim, 450)
	second.SelfShield().Apply(sim, 200)
	if got := hit(500, false); got != 0 || shield.IsActive() || second.SelfShield().RemainingAbsorb() != 150 {
		t.Fatal("overlapping pool application order/capacity mismatch")
	}
	target.auraTracker.doneIteration(sim)
	if len(target.activeShields) != 0 || second.SelfShield().RemainingAbsorb() != 0 {
		t.Fatal("shield capacity leaked across iteration cleanup")
	}
}
