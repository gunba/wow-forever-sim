package core

import (
	"math"
	"testing"
)

func TestEurekaReservesChargeBeforeNestedCast(t *testing.T) {
	character := &Character{Unit: Unit{auraTracker: newAuraTracker()}}
	sim := &Simulation{}
	aura := character.RegisterAura(Aura{Label: "Eureka test", Duration: NeverExpires, MaxStacks: 3})
	aura.Activate(sim)
	aura.SetStacks(sim, 1)
	var outerDamage, innerDamage float64
	inner := &Spell{DamageMultiplier: 1, ApplyEffects: func(_ *Simulation, _ *Unit, spell *Spell) {
		innerDamage = spell.DamageMultiplier
	}}
	outer := &Spell{DamageMultiplier: 1, ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
		outerDamage = spell.DamageMultiplier
		inner.ApplyEffects(sim, target, inner)
	}}
	character.attachEurekaCast(aura, inner, ProcMaskMeleeSpecial)
	character.attachEurekaCast(aura, outer, ProcMaskMeleeSpecial)
	outer.ApplyEffects(sim, nil, outer)
	if math.Abs(outerDamage-1.1) > 1e-9 || innerDamage != 1 || aura.IsActive() ||
		math.Abs(outer.DamageMultiplier-1) > 1e-9 || inner.DamageMultiplier != 1 {
		t.Fatalf("final charge was double-spent: outer=%v inner=%v stacks=%v", outerDamage, innerDamage, aura.GetStacks())
	}
}
