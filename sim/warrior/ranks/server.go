package ranks

// Server-side numbers the client tables do not carry. Each is sourced or marked provisional.

// HeroicStrikeThreat is the flat threat of a Heroic Strike rank. Measured: rank 4 adds 61 (published Warrior threat
// sheet, level 30); ranks 8 and 9 keep the simulator's long-standing 145 and 173. Other ranks are provisional:
// below rank 4 the measured value is scaled by the rank's bonus damage, between 4 and 8 interpolated by bonus damage.
func HeroicStrikeThreat(r *Rank) float64 {
	bonus := r.Find(EffectWeaponDamageBonus, 0).Low
	switch {
	case r.Rank >= 9:
		return 173
	case r.Rank == 8:
		return 145
	case r.Rank <= 4:
		return 61 * bonus / 44
	default:
		return 61 + (145-61)*(bonus-44)/(138-44)
	}
}

// CleaveThreat is the flat threat of a Cleave rank, per target. Measured: rank 2 adds 46 (threat sheet); rank 5 keeps
// the simulator's 100. Other ranks are provisional, interpolated by bonus damage.
func CleaveThreat(r *Rank) float64 {
	bonus := r.Find(EffectWeaponDamageBonus, 0).Low
	switch {
	case r.Rank >= 5:
		return 100
	case r.Rank <= 2:
		return 46 * max(bonus, 1) / 10
	default:
		return 46 + (100-46)*(bonus-10)/(50-10)
	}
}

// SunderThreatAP is the attack power coefficient of Sunder Armor's threat: 0.0433 at rank 1 and 0.0476 at rank 2
// (threat sheet); 0.05 at rank 5 is the simulator's provisional value, and ranks 3-4 keep rank 2's.
func SunderThreatAP(r *Rank) float64 {
	switch {
	case r.Rank <= 1:
		return 0.0433
	case r.Rank >= 5:
		return 0.05
	default:
		return 0.0476
	}
}

// ExecuteDamagePerRage is Classic's extra damage per Rage converted: 3 per rank.
func ExecuteDamagePerRage(r *Rank) float64 { return 3 * float64(r.Rank) }

// DemoralizingThreatScale scales the client's flat threat of Demoralizing Shout to the measured value: rank 2 lists
// 19 and was measured at 22.7 before modifiers (threat sheet). Provisional for other ranks.
const DemoralizingThreatScale = 22.74 / 19

// RevengeAPCoefficient: measured at rank 2 (threat sheet fit); used for every rank.
const RevengeAPCoefficient = 0.28

// RendAPCoefficient: 2% of attack power per tick, measured at rank 2 (Fight Club logs); used for every rank.
const RendAPCoefficient = 0.02

// ThunderClapAPCoefficient: about 7% of attack power (testers); used for every rank.
const ThunderClapAPCoefficient = 0.07
