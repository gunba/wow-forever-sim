package core

import (
	"slices"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

type foreverEquipmentBonus struct {
	Supported   bool
	Stats       stats.Stats
	SpiritRegen float64
}

type foreverEquipmentSet struct {
	Name       string
	Profession proto.Profession
	Bonuses    map[int32]foreverEquipmentBonus
}

func ForeverEquipmentBonusSupported(id, pieces int32) bool {
	if isForeverTier1EquipmentSet(id) {
		return pieces >= 2 && pieces <= 5
	}
	return foreverEquipmentSets[id].Bonuses[pieces].Supported
}

func isForeverTier1EquipmentSet(id int32) bool {
	for spec := range proto.Spec_name {
		if id != 0 && ForeverTier1SetID(proto.Spec(spec)) == id {
			return true
		}
	}
	return false
}

func overridesForeverEquipmentSet(id int32) bool {
	_, exists := foreverEquipmentSets[id]
	return exists && !isForeverTier1EquipmentSet(id)
}

// A catalog set is isolated from obsolete Classic effects even when some of
// its current thresholds remain unimplemented.
func HasForeverEquipmentSetDefinition(id int32) bool {
	return overridesForeverEquipmentSet(id) || isForeverTier1EquipmentSet(id)
}

func (character *Character) foreverEquippedBonuses() []ActiveSetBonus {
	counts := map[int32]int32{}
	for _, item := range character.Equipment {
		counts[item.SetID]++
	}
	var ids []int32
	for id := range counts {
		ids = append(ids, id)
	}
	slices.Sort(ids)
	var active []ActiveSetBonus
	for _, id := range ids {
		set, exists := foreverEquipmentSets[id]
		if !exists || !overridesForeverEquipmentSet(id) {
			continue
		}
		if set.Profession != proto.Profession_ProfessionUnknown && !character.HasProfession(set.Profession) {
			continue
		}
		var pieces []int32
		for threshold, bonus := range set.Bonuses {
			if bonus.Supported && counts[id] >= threshold {
				pieces = append(pieces, threshold)
			}
		}
		slices.Sort(pieces)
		for _, threshold := range pieces {
			bonus := set.Bonuses[threshold]
			active = append(active, ActiveSetBonus{
				Name: set.Name, NumPieces: threshold,
				BonusEffect: func(agent Agent) {
					c := agent.GetCharacter()
					c.AddStats(bonus.Stats)
					c.PseudoStats.SpiritRegenRateCasting += bonus.SpiritRegen
				},
			})
		}
	}
	return active
}
