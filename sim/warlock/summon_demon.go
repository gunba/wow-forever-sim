package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (warlock *Warlock) registerSummonDemon() {
	manaCost := core.ManaCostOptions{
		FlatCost: warlock.BaseMana,
	}
	for _, pet := range warlock.BasePets {
		pet.summonStunAura = pet.RegisterAura(core.Aura{
			Label: "Summoning another demon", Duration: core.NeverExpires,
			OnGain:   func(_ *core.Aura, sim *core.Simulation) { pet.AddStun(sim) },
			OnExpire: func(_ *core.Aura, sim *core.Simulation) { pet.RemoveStun(sim) },
		})
	}
	// Keep the existing Forever pet and its owned buffs until replacement.
	cast := core.CastConfig{
		DefaultCast: core.Cast{
			GCD:      core.GCDDefault,
			CastTime: time.Second * 10,
		},
		OnCastStart: func(sim *core.Simulation, spell *core.Spell) {
			if !warlock.Env.IsForever() {
				warlock.changeActivePet(sim, nil, false)
				return
			}
			if spell.CurCast.CastTime > 0 && warlock.ActivePet != nil && warlock.ActivePet.IsEnabled() {
				warlock.summoningPet = warlock.ActivePet
				warlock.summoningPet.summonStunAura.Activate(sim)
			}
		},
		OnCastEnd: func(sim *core.Simulation, _ *core.Spell, _ bool) {
			if warlock.summoningPet != nil {
				warlock.summoningPet.summonStunAura.Deactivate(sim)
				warlock.summoningPet = nil
			}
		},
	}

	// Felhunter
	warlock.SummonDemonSpells = append(warlock.SummonDemonSpells, warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 691},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCost,
		Cast:     cast,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.changeActivePet(sim, warlock.Felhunter, false)
		},
	}))

	// Imp
	warlock.SummonDemonSpells = append(warlock.SummonDemonSpells, warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 688},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCost,
		Cast:     cast,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.changeActivePet(sim, warlock.Imp, false)
		},
	}))

	// Succubus
	warlock.SummonDemonSpells = append(warlock.SummonDemonSpells, warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 712},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCost,
		Cast:     cast,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.changeActivePet(sim, warlock.Succubus, false)
		},
	}))

	// Voidwalker
	warlock.SummonDemonSpells = append(warlock.SummonDemonSpells, warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 697},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCost,
		Cast:     cast,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.changeActivePet(sim, warlock.Voidwalker, false)
		},
	}))
}
