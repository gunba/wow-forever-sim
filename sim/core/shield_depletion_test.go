package core

import (
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"testing"
	"time"
)

func TestShieldDepletionCallback(t *testing.T) {
	env := &Environment{Ruleset: proto.Ruleset_RulesetForever}
	target := &Unit{Env: env, Type: PlayerUnit, UnitIndex: 0, PseudoStats: stats.NewPseudoStats(), Metrics: NewUnitMetrics()}
	enemy := &Unit{Env: env, Type: EnemyUnit, UnitIndex: 1, Level: 62, PseudoStats: stats.NewPseudoStats(), Metrics: NewUnitMetrics()}
	env.AllUnits = []*Unit{target, enemy}
	sim := &Simulation{Environment: env, Duration: 30 * time.Second}
	count := 0
	var source *Spell
	s := target.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 777}, ProcMask: ProcMaskEmpty, SpellSchool: SpellSchoolHoly, DamageMultiplier: 1, Shield: ShieldConfig{SelfOnly: true, Aura: Aura{Label: "Depletioncallback", Duration: time.Second}, OnDepleted: func(_ *Simulation, incoming *Spell, result *SpellResult) {
		count++
		source = incoming
		if result.Target != target || incoming.Unit.Level != 62 {
			t.Fatal("lostactualattacker/recipient")
		}
	}}})
	incoming := enemy.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 778}, SpellSchool: SpellSchoolHoly, ProcMask: ProcMaskEmpty, DamageMultiplier: 1, ThreatMultiplier: 1})
	s.finalize()
	incoming.finalize()
	shield := s.SelfShield()
	hit := func(amount float64) {
		result := incoming.NewResult(target)
		result.Outcome = OutcomeHit
		result.Damage = amount
		incoming.DealDamage(sim, result)
	}
	shield.Apply(sim, 10)
	shield.Apply(sim, 20)
	if count != 0 {
		t.Fatal("replacementproc")
	}
	hit(5)
	if count != 0 || shield.RemainingAbsorb() != 15 {
		t.Fatal("partialdepletionproc")
	}
	hit(15)
	if count != 1 || source != incoming || shield.IsActive() {
		t.Fatal("fulldepletion")
	}
	hit(20)
	if count != 1 {
		t.Fatal("duplicateproc")
	}
	shield.Apply(sim, 20)
	shield.Deactivate(sim)
	if count != 1 {
		t.Fatal("manualdeactivateproc")
	}
	shield.Apply(sim, 20)
	sim.CurrentTime = time.Second
	hit(20)
	if count != 1 {
		t.Fatal("expiryproc")
	}
}
