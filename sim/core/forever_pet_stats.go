package core

import "github.com/wowsims/classic/sim/core/stats"

// ForeverPetInheritance uses the shared owner-scaling coefficients reported
// for Forever Hunter pets. The client has separate Hunter (415429) and
// Warlock (416189) scaling auras but does not export their server-side
// coefficients; the same coefficients are provisionally used for both.
// Spell-power-to-pet-power is the Warlock-specific interpretation.
func ForeverPetInheritance(owner stats.Stats, casterOwner bool) stats.Stats {
	offensivePower := max(owner[stats.AttackPower], owner[stats.RangedAttackPower])
	spellPower := owner[stats.SpellPower] + owner[stats.SpellDamage] +
		max(owner[stats.ArcanePower], owner[stats.FirePower],
			owner[stats.FrostPower], owner[stats.HolyPower],
			owner[stats.NaturePower], owner[stats.ShadowPower])
	if casterOwner {
		offensivePower = max(offensivePower, spellPower)
	}

	inherited := stats.Stats{
		stats.Health:      2 * owner[stats.Stamina],
		stats.Armor:       0.30 * owner[stats.Armor],
		stats.AttackPower: 0.10 * offensivePower,
		stats.MeleeHit:    owner[stats.MeleeHit],
		stats.SpellHit:    owner[stats.SpellHit],
		// Since client 1.60.1.70425 Hunter pets take the higher of the owner's melee and ranged crit.
		// The sim keeps one physical crit stat, and its ranged-only modifiers only lower ranged crit
		// (Elemental Sharpening Stone), so MeleeCrit is that higher value.
		stats.MeleeCrit: owner[stats.MeleeCrit],
		stats.SpellCrit: owner[stats.SpellCrit],
	}
	if !casterOwner {
		// Hunter pets benefit from the owner's Expertise since client 1.60.1.70425 (October 15 notes).
		inherited[stats.Expertise] = owner[stats.Expertise]
	}
	if casterOwner {
		inherited[stats.SpellPower] = 0.10 * spellPower
	}
	return inherited
}
