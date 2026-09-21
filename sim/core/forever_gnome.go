package core

import (
	"time"

	"github.com/wowsims/classic/sim/core/proto"
)

// SkillLineAbility assigns the five Eureka spells by class. Their durations,
// cost modifiers and three proc charges are captured in forever_races.json.
func (character *Character) registerEureka() {
	var id, discount int32
	switch character.Class {
	case proto.Class_ClassRogue:
		id, discount = 1259812, 20
	case proto.Class_ClassWarrior:
		id, discount = 1259813, 40
	case proto.Class_ClassMage:
		id, discount = 1259817, 50
	case proto.Class_ClassWarlock:
		id, discount = 1259821, 50
	case proto.Class_ClassPriest:
		id, discount = 1259823, 15
	default:
		return
	}
	actionID := ActionID{SpellID: id}
	var affected []*Spell
	aura := character.RegisterAura(Aura{
		Label: "Eureka!", ActionID: actionID, Duration: 15 * time.Second, MaxStacks: 3,
		OnInit: func(aura *Aura, _ *Simulation) {
			mask := ProcMaskMeleeSpecial | ProcMaskRangedSpecial | ProcMaskSpellDamage
			if character.Class == proto.Class_ClassPriest {
				mask |= ProcMaskSpellHealing
			}
			for _, spell := range character.Spellbook {
				if spell.Cost == nil || !spell.ProcMask.Matches(mask) || spell.Flags.Matches(SpellFlagPassiveSpell) {
					continue
				}
				affected = append(affected, spell)
				character.attachEurekaCast(aura, spell, mask)
			}
		},
		OnGain: func(_ *Aura, _ *Simulation) {
			for _, spell := range affected {
				spell.Cost.Multiplier -= discount
			}
		},
		OnExpire: func(_ *Aura, _ *Simulation) {
			for _, spell := range affected {
				spell.Cost.Multiplier += discount
			}
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, _ int32, stacks int32) {
			if stacks == 0 {
				aura.Deactivate(sim)
			}
		},
	})
	spell := character.RegisterSpell(SpellConfig{
		ActionID: actionID, Flags: SpellFlagNoOnCastComplete,
		Cast: CastConfig{CD: Cooldown{Timer: character.NewTimer(), Duration: 2 * time.Minute}},
		ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
			aura.Activate(sim)
			aura.SetStacks(sim, 3)
		},
	})
	character.AddMajorCooldown(MajorCooldown{Spell: spell, Type: CooldownTypeDPS})
}

func (character *Character) attachEurekaCast(aura *Aura, parent *Spell, mask ProcMask) {
	// Off-hand Mutilate hits and Arcane Missiles ticks use tagged effect spells.
	// They belong to the charged cast, not additional charge-consuming casts.
	effects := []*Spell{parent}
	var children []*Spell
	for _, spell := range character.Spellbook {
		if spell != parent && spell.SpellID == parent.SpellID && spell.Cost == nil && spell.ProcMask.Matches(mask) {
			effects = append(effects, spell)
			children = append(children, spell)
		}
	}
	multiply := func(spells []*Spell, factor float64) {
		for _, spell := range spells {
			spell.DamageMultiplier *= factor
		}
	}
	dots := parent.Dots()
	chargedDots := make(map[*Dot]bool)
	applying := false
	for _, dot := range dots {
		if dot == nil || !dot.isChanneled || len(children) == 0 || dot.OnTick == nil {
			continue
		}
		onTick := dot.OnTick
		dot.OnTick = func(sim *Simulation, target *Unit, dot *Dot) {
			charged := chargedDots[dot] && !applying
			if charged {
				multiply(children, 1.1)
			}
			onTick(sim, target, dot)
			if charged {
				multiply(children, 1/1.1)
			}
		}
	}
	apply := parent.ApplyEffects
	parent.ApplyEffects = func(sim *Simulation, target *Unit, spell *Spell) {
		charged := aura.IsActive() && aura.GetStacks() > 0
		for _, dot := range dots {
			if dot != nil && dot.Unit == target {
				chargedDots[dot] = charged
			}
		}
		if charged {
			multiply(effects, 1.1)
			// Reserve this cast's charge before damage/resource callbacks can
			// trigger another eligible cast. The final charge belongs to this
			// cast, not simultaneously to every nested attack.
			aura.RemoveStack(sim)
		}
		applying = true
		apply(sim, target, spell)
		applying = false
		if charged {
			multiply(effects, 1/1.1)
		}
	}
}
