package shaman

import (
	"github.com/wowsims/classic/sim/core"
)

const CastTagLightningOverload = 1

// Overloads are free instant triggered spells with no threat. Classic halves
// the parent; Forever configurations already carry the separate child row's
// damage and SP coefficient, so they must not be halved a second time.
func (shaman *Shaman) registerOverloadSpell(config core.SpellConfig) *core.Spell {
	if shaman.Talents.LightningOverload == 0 {
		return nil
	}

	config.ActionID.Tag = CastTagLightningOverload
	config.Flags = config.Flags&^core.SpellFlagAPL | core.SpellFlagPassiveSpell | core.SpellFlagNoOnCastComplete
	config.Cast = core.CastConfig{}
	config.ManaCost = core.ManaCostOptions{}
	if !shaman.Env.IsForever() {
		config.DamageMultiplier *= .5
	}
	config.ThreatMultiplier = 0

	return shaman.RegisterSpell(config)
}

// 3, 7 and 10% are the beta client's talent curve.
func (shaman *Shaman) lightningOverloadChance() float64 {
	return []float64{0, .03, .07, .10}[shaman.Talents.LightningOverload]
}

func (shaman *Shaman) tryLightningOverload(sim *core.Simulation, target *core.Unit, spell *core.Spell, overload *core.Spell) {
	if overload == nil || spell == overload {
		return
	}

	if sim.Proc(shaman.lightningOverloadChance(), "Lightning Overload") {
		overload.Cast(sim, target)
	}
}
