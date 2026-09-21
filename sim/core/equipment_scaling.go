package core

import "math"

// Scaled returns per-character item copies; the item catalog remains unchanged.
func (equipment Equipment) Scaled(multiplier float64) Equipment {
	if multiplier == 0 || multiplier == 1 {
		return equipment
	}
	if multiplier < 0 || math.IsNaN(multiplier) || math.IsInf(multiplier, 0) {
		panic("equipment scale must be finite and positive")
	}
	for i := range equipment {
		item := &equipment[i]
		item.Stats = item.Stats.Multiply(multiplier)
		item.RandomSuffix.Stats = item.RandomSuffix.Stats.Multiply(multiplier)
		item.WeaponDamageMin *= multiplier
		item.WeaponDamageMax *= multiplier
		item.BonusPhysicalDamage *= multiplier
	}
	return equipment
}
