package core

import (
	"slices"
	"strconv"
)

type ShieldConfig struct {
	SelfOnly bool // Set to true to only create the self-shield.

	Spell *Spell

	// Called only after an actual incoming damage result fully consumes a pool.
	// Expiration, replacement and explicit deactivation do not invoke it.
	OnDepleted func(*Simulation, *Spell, *SpellResult)
	Aura
}

// Represents a finite all-school absorption effect, e.g. Power Word: Shield.
type Shield struct {
	Spell *Spell

	// Embed Aura so we can use IsActive/Refresh/etc directly.
	*Aura

	remainingAbsorb float64
	onDepleted      func(*Simulation, *Spell, *SpellResult)
}

func (shield *Shield) RemainingAbsorb() float64 { return shield.remainingAbsorb }

func (shield *Shield) Apply(sim *Simulation, shieldAmount float64) {
	caster := shield.Spell.Unit
	target := shield.Aura.Unit
	//attackTable := caster.AttackTables[target.UnitIndex][shield.Spell.CastType]

	// Shields are not affected by healing pseudostats the same way heals are.
	// So we only apply the spell-specific multiplier and shield-specific multiplier.
	shieldAmount = max(0, shieldAmount*shield.Spell.DamageMultiplier*caster.PseudoStats.ShieldDealtMultiplier)

	shield.Aura.Deactivate(sim)
	shield.remainingAbsorb = shieldAmount
	if shieldAmount > 0 {
		shield.Aura.Activate(sim)
		target.activeShields = append(target.activeShields, shield)
	}

	threat := 0.0 // TODO
	shield.Spell.SpellMetrics[target.UnitIndex].TotalThreat += threat
	shield.Spell.SpellMetrics[target.UnitIndex].TotalShielding += shieldAmount
	shield.Spell.SpellMetrics[target.UnitIndex].Hits++

	if sim.Log != nil {
		caster.Log(sim, "%s %s Hit for %0.3f shielding. (Threat: %0.3f)", target.LogLabel(), shield.Spell.ActionID, shieldAmount, threat)
	}
}

func newShield(config Shield) *Shield {
	shield := &Shield{}
	*shield = config
	previousExpire := shield.Aura.OnExpire
	shield.Aura.OnExpire = func(aura *Aura, sim *Simulation) {
		shield.remainingAbsorb = 0
		aura.Unit.activeShields = slices.DeleteFunc(aura.Unit.activeShields, func(active *Shield) bool { return active == shield })
		if previousExpire != nil {
			previousExpire(aura, sim)
		}
	}

	return shield
}

// Pools are consumed after mitigation and only during damage delivery. In the
// absence of a verified multi-shield priority rule, overlaps use application
// order. Refreshing a pool replaces it and moves it to the end of that order.
func (unit *Unit) absorbDamage(sim *Simulation, incoming *Spell, result *SpellResult) {
	for result.Damage > 0 && len(unit.activeShields) > 0 {
		shield := unit.activeShields[0]
		if shield.ExpiresAt() <= sim.CurrentTime {
			shield.Deactivate(sim)
			continue
		}
		amount := min(result.Damage, shield.remainingAbsorb)
		shield.remainingAbsorb -= amount
		result.Damage -= amount
		result.AbsorbedDamage += amount
		shield.Spell.SpellMetrics[unit.UnitIndex].TotalAbsorbedShielding += amount
		if sim.Log != nil {
			shield.Spell.Unit.Log(sim, "%s %s absorbed %0.3f damage (%0.3f remaining).", unit.LogLabel(), shield.Spell.ActionID, amount, shield.remainingAbsorb)
		}
		if shield.remainingAbsorb <= 0 {
			// Remove the consumed pool first; keep the incoming caster/result
			// available to the callback without deactivating a newly applied pool.
			shield.Deactivate(sim)
			if shield.onDepleted != nil {
				shield.onDepleted(sim, incoming, result)
			}
		}
	}
}

type ShieldArray []*Shield

func (shields ShieldArray) Get(target *Unit) *Shield {
	return shields[target.UnitIndex]
}

func (spell *Spell) createShields(config ShieldConfig) {
	if config.Aura.Label == "" {
		return
	}

	if config.Spell == nil {
		config.Spell = spell
	}
	shield := Shield{
		Spell:      config.Spell,
		onDepleted: config.OnDepleted,
	}

	auraConfig := config.Aura
	if auraConfig.ActionID.IsEmptyAction() {
		auraConfig.ActionID = shield.Spell.ActionID
	}

	caster := shield.Spell.Unit
	if config.SelfOnly {
		shield.Aura = caster.GetOrRegisterAura(auraConfig)
		spell.selfShield = newShield(shield)
	} else {
		auraConfig.Label += "-" + strconv.Itoa(int(caster.UnitIndex))
		if spell.shields == nil {
			spell.shields = make([]*Shield, len(caster.Env.AllUnits))
		}
		for _, target := range caster.Env.AllUnits {
			if !caster.IsOpponent(target) {
				shield.Aura = target.GetOrRegisterAura(auraConfig)
				spell.shields[target.UnitIndex] = newShield(shield)
			}
		}
	}
}
