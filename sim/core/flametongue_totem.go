package core

import "fmt"

// The client encodes these as dummy-effect values, not direct damage. The
// division by 100 is a provisional interpretation of the speed-based tooltip;
// it is not a verified server formula. Spell power deliberately contributes 0.
var flametongueTotemDPS = [...]float64{0, 5.48, 7.81, 10.61, 13.63}
var flametongueTotemBuffIDs = [...]int32{0, 8230, 8250, 10521, 15036}
var flametongueTotemDamageIDs = [...]int32{0, 8253, 8248, 10523, 16389}

func hasFlametongueWeapon(character *Character) bool {
	switch character.MainHand().TempEnchant {
	case 5, 4, 3, 523, 1665, 1666:
		return true
	}
	return false
}

// FlametongueTotemAura models an external provider or a caster-owned totem.
// Main-hand melee proc eligibility and magic hit/crit behavior are provisional,
// as is using the active form's base swing speed for a shapeshifted recipient.
func FlametongueTotemAura(character *Character, rank int) *Aura {
	if rank < 1 || rank >= len(flametongueTotemDPS) {
		panic("invalid Flametongue Totem rank")
	}
	label := fmt.Sprintf("Flametongue Totem (Rank %d; provisional)", rank)
	if aura := character.GetAura(label); aura != nil {
		return aura
	}
	proc := character.RegisterSpell(SpellConfig{
		ActionID:         ActionID{SpellID: flametongueTotemDamageIDs[rank]},
		SpellSchool:      SpellSchoolFire,
		DefenseType:      DefenseTypeMagic,
		ProcMask:         ProcMaskSpellDamageProc,
		Flags:            SpellFlagPassiveSpell | SpellFlagNoOnCastComplete,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			// Base speed, not the haste-adjusted swing interval. Weapon swaps
			// and forms use the current main-hand weapon model.
			base := flametongueTotemDPS[rank] * character.AutoAttacks.MH().SwingSpeed
			spell.CalcAndDealDamage(sim, target, base, spell.OutcomeMagicHitAndCrit)
		},
	})
	var exclusive *ExclusiveEffect
	aura := character.RegisterAura(Aura{
		Label:    label,
		ActionID: ActionID{SpellID: flametongueTotemBuffIDs[rank]},
		Duration: NeverExpires,
		OnSpellHitDealt: func(_ *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			if !exclusive.IsActive() || !result.Landed() ||
				!spell.ProcMask.Matches(ProcMaskMeleeMH) ||
				spell.Flags.Matches(SpellFlagSuppressWeaponProcs|SpellFlagSuppressEquipProcs) ||
				!character.HasMHWeapon() || hasFlametongueWeapon(character) {
				return
			}
			// Windfury Weapon can coexist with this totem, but its triggered
			// hits do not trigger additional totem damage. Keep the exclusion
			// local: it says nothing about other weapon-proc eligibility.
			switch spell.ActionID.SpellID {
			case 8232, 8235, 10486, 16362:
				return
			}
			proc.Cast(sim, result.Target)
		},
	})
	// Windfury wins conflicting imported settings. Grace of Air occupies a
	// different category and can coexist with Flametongue.
	exclusive = aura.NewExclusiveEffect("ForeverWeaponTotem", false, ExclusiveEffect{Priority: float64(rank) / 10})
	return aura
}
