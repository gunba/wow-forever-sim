package core

import (
	"math"
	"testing"
)

func TestEurekaBoostsEveryCastWhileActive(t *testing.T) {
	character := &Character{Unit: Unit{auraTracker: newAuraTracker()}}
	sim := &Simulation{}
	aura := character.RegisterAura(Aura{Label: "Eureka test", Duration: NeverExpires})
	aura.Activate(sim)
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
	if math.Abs(outerDamage-1.1) > 1e-9 || math.Abs(innerDamage-1.1) > 1e-9 || !aura.IsActive() ||
		math.Abs(outer.DamageMultiplier-1) > 1e-9 || math.Abs(inner.DamageMultiplier-1) > 1e-9 {
		t.Fatalf("an 8-second Eureka should boost both casts and stay up: outer=%v inner=%v", outerDamage, innerDamage)
	}
}
