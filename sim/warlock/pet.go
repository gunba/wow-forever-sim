package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

type OnPetDisable func(sim *core.Simulation, isSacrifice bool)

type WarlockPet struct {
	core.Pet

	OnPetDisable OnPetDisable

	owner *Warlock

	primaryAbility   *core.Spell
	secondaryAbility *core.Spell

	SoulLinkAura *core.Aura

	manaPooling bool
}

type PetConfig struct {
	Name          string
	PowerModifier float64 // GetUnitPowerModifier("pet")
	Stats         stats.Stats
	AutoAttacks   core.AutoAttackOptions
}

func (warlock *Warlock) petFromSummon(summon proto.WarlockOptions_Summon) *WarlockPet {
	switch summon {
	case proto.WarlockOptions_Imp:
		return warlock.Imp
	case proto.WarlockOptions_Felhunter:
		return warlock.Felhunter
	case proto.WarlockOptions_Succubus:
		return warlock.Succubus
	case proto.WarlockOptions_Voidwalker:
		return warlock.Voidwalker
	}
	return nil
}

func (warlock *Warlock) setDefaultActivePet() {
	warlock.ActivePet = warlock.petFromSummon(warlock.Options.Summon)
}

func (warlock *Warlock) changeActivePet(sim *core.Simulation, newPet *WarlockPet, isSacrifice bool) {
	if warlock.ActivePet != nil {
		warlock.ActivePet.Disable(sim, isSacrifice)

		// Sacrificed pets lose all buffs
		if isSacrifice {
			for _, aura := range warlock.ActivePet.GetAuras() {
				aura.Deactivate(sim)
			}
		}

		if warlock.MasterDemonologistAura != nil {
			warlock.MasterDemonologistAura.Deactivate(sim)
		}
	}

	warlock.ActivePet = newPet

	if newPet != nil {
		newPet.Enable(sim, newPet)
	}
}

func (warlock *Warlock) registerPets() {
	warlock.Felhunter = warlock.makeFelhunter()
	warlock.Imp = warlock.makeImp()
	warlock.Succubus = warlock.makeSuccubus()
	warlock.Voidwalker = warlock.makeVoidwalker()

	warlock.BasePets = []*WarlockPet{warlock.Felhunter, warlock.Imp, warlock.Succubus, warlock.Voidwalker}
}

func (warlock *Warlock) makePet(cfg PetConfig, enabledOnStart bool) *WarlockPet {
	wp := &WarlockPet{
		Pet:          core.NewPet(cfg.Name, &warlock.Character, cfg.Stats, warlock.makeStatInheritance(), enabledOnStart, false),
		owner:        warlock,
		OnPetDisable: func(sim *core.Simulation, isSacrifice bool) {},
	}

	wp.EnableManaBarWithModifier(cfg.PowerModifier)

	if cfg.Name == "Imp" {
		// Imp gets 1mp/5 non casting regen per spirit
		wp.PseudoStats.SpiritRegenMultiplier = 1
		wp.PseudoStats.SpiritRegenRateCasting = 0
		wp.SpiritManaRegenPerSecond = func() float64 {
			// 1mp5 per spirit
			return wp.GetStat(stats.Spirit) / 5
		}

		// Mage spell crit scaling for imp
		wp.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[proto.Class_ClassMage]*core.SpellCritRatingPerCritChance)
	} else {
		// Warrior scaling for all other pets
		wp.AddStat(stats.AttackPower, -20)
		wp.AddStatDependency(stats.Strength, stats.AttackPower, 2)

		// Warrior crit scaling
		wp.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritPerAgiAtLevel[proto.Class_ClassWarrior]*core.CritRatingPerCritChance)
		wp.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[proto.Class_ClassWarrior]*core.SpellCritRatingPerCritChance)

		// Imps generally don't melee
		wp.EnableAutoAttacks(wp, cfg.AutoAttacks)
	}

	warlock.AddPet(wp)

	return wp
}

func (warlock *Warlock) registerPetAbilities() {
	warlock.Imp.registerImpFireboltSpell()
	warlock.Succubus.registerSuccubusLashOfPainSpell()
}

func (wp *WarlockPet) GetPet() *core.Pet {
	return &wp.Pet
}

func (wp *WarlockPet) Initialize() {
	core.ApplyPetConsumeEffects(&wp.Pet, wp.owner.Consumes)
}

func (wp *WarlockPet) Reset(_ *core.Simulation) {
	wp.manaPooling = false
}

func (wp *WarlockPet) Disable(sim *core.Simulation, isSacrifice bool) {
	wp.Pet.Disable(sim)

	if wp.OnPetDisable != nil {
		wp.OnPetDisable(sim, isSacrifice)
	}
}

func (wp *WarlockPet) ApplyOnPetDisable(newOnPetDisable OnPetDisable) {
	oldOnPetDisable := wp.OnPetDisable
	if oldOnPetDisable == nil {
		wp.OnPetDisable = oldOnPetDisable
	} else {
		wp.OnPetDisable = func(sim *core.Simulation, isSacrifice bool) {
			oldOnPetDisable(sim, isSacrifice)
			newOnPetDisable(sim, isSacrifice)
		}
	}
}

func (wp *WarlockPet) ExecuteCustomRotation(sim *core.Simulation) {
	if !wp.IsEnabled() || wp.primaryAbility == nil {
		return
	}

	if wp.manaPooling {
		maxPossibleCasts := sim.GetRemainingDuration().Seconds() / wp.primaryAbility.CurCast.CastTime.Seconds()

		if wp.CurrentMana() > (maxPossibleCasts*wp.primaryAbility.CurCast.Cost)*0.75 {
			wp.manaPooling = false
			wp.WaitUntil(sim, sim.CurrentTime+10*time.Millisecond)
			return
		}

		if wp.CurrentMana() >= wp.MaxMana()*0.94 {
			wp.manaPooling = false
			wp.WaitUntil(sim, sim.CurrentTime+10*time.Millisecond)
			return
		}

		if wp.manaPooling {
			return
		}
	}

	if !wp.primaryAbility.IsReady(sim) {
		wp.WaitUntil(sim, wp.primaryAbility.CD.ReadyAt())
		return
	}

	if wp.Unit.CurrentMana() >= wp.primaryAbility.CurCast.Cost {
		wp.primaryAbility.Cast(sim, wp.CurrentTarget)
	} else if !wp.owner.Options.PetPoolMana {
		wp.manaPooling = true
	}
}

func (warlock *Warlock) makeStatInheritance() core.PetStatInheritance {
	return func(ownerStats stats.Stats) stats.Stats {
		if !warlock.Env.IsForever() {
			return stats.Stats{}
		}
		return core.ForeverPetInheritance(ownerStats, true)
	}
}
