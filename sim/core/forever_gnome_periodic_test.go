package core

import (
	"math"
	"testing"
)

func TestEurekaExcludesPeriodicSnapshotAndApplicationTick(t *testing.T) {
	character := &Character{Unit: Unit{auraTracker: newAuraTracker()}}
	sim := &Simulation{}
	target := &Unit{}
	aura := character.RegisterAura(Aura{Label: "Eureka test", Duration: NeverExpires, MaxStacks: 3})
	aura.Activate(sim)
	aura.SetStacks(sim, 1)
	parent := &Spell{DamageMultiplier: 1}
	var direct, snapshot, initialTick, laterTick float64
	dot := &Dot{Aura: &Aura{Unit: target}}
	dot.OnSnapshot = func(_ *Simulation, _ *Unit, _ *Dot, _ bool) { snapshot = parent.DamageMultiplier }
	dot.OnTick = func(_ *Simulation, _ *Unit, _ *Dot) { initialTick = parent.DamageMultiplier }
	parent.dots = []*Dot{dot}
	parent.ApplyEffects = func(sim *Simulation, target *Unit, _ *Spell) {
		direct = parent.DamageMultiplier
		dot.OnSnapshot(sim, target, dot, false)
		dot.OnTick(sim, target, dot)
	}
	character.attachEurekaCast(aura, parent, ProcMaskSpellDamage)
	parent.ApplyEffects(sim, target, parent)
	laterTick = parent.DamageMultiplier
	if math.Abs(direct-1.1) > 1e-9 || math.Abs(snapshot-1) > 1e-9 || math.Abs(initialTick-1) > 1e-9 || math.Abs(laterTick-1) > 1e-9 || aura.IsActive() {
		t.Fatalf("direct=%v snapshot=%v initial periodic=%v later=%v active=%v", direct, snapshot, initialTick, laterTick, aura.IsActive())
	}
}

func TestEurekaIncludesChannelWithoutSeparateEffectSpell(t *testing.T) {
	character := &Character{Unit: Unit{auraTracker: newAuraTracker()}}
	sim := &Simulation{}
	target := &Unit{}
	aura := character.RegisterAura(Aura{Label: "Eureka channel test", Duration: NeverExpires, MaxStacks: 3})
	aura.Activate(sim)
	aura.SetStacks(sim, 1)
	parent := &Spell{DamageMultiplier: 1}
	var tickMultiplier float64
	dot := &Dot{Aura: &Aura{Unit: target}, isChanneled: true}
	dot.OnTick = func(_ *Simulation, _ *Unit, _ *Dot) { tickMultiplier = parent.DamageMultiplier }
	parent.dots = []*Dot{dot}
	parent.ApplyEffects = func(_ *Simulation, _ *Unit, _ *Spell) {}
	character.attachEurekaCast(aura, parent, ProcMaskSpellDamage)
	parent.ApplyEffects(sim, target, parent)
	for i := 0; i < 3; i++ {
		dot.OnTick(sim, target, dot)
		if math.Abs(tickMultiplier-1.1) > 1e-9 || math.Abs(parent.DamageMultiplier-1) > 1e-9 {
			t.Fatalf("channel tick %d multiplier=%v retained=%v", i, tickMultiplier, parent.DamageMultiplier)
		}
	}
	if aura.IsActive() {
		t.Fatal("channel failed to spend its single cast charge")
	}
}
