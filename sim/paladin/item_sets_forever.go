package paladin

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"time"
)

var ItemSetJusticeBattlegear = core.NewItemSet(core.ItemSet{
	ID: 2106, Name: "Justice Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HasteBonus,
		// Hammer of Justice is not modeled.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301084, proto.MobType_MobTypeDemon, 36, 0),
		5: core.ForeverTier1CooldownBonus(500*time.Millisecond, func(s *core.Spell) bool { return s.SpellID == 20271 }),
	},
})

var ItemSetJusticeBattleplate = core.NewItemSet(core.ItemSet{
	ID: 2108, Name: "Justice Battleplate",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1DefenseBonus,
		// Turn Undead cast time and Forbearance duration are handled where
		// those spells/auras are registered; the former has no boss target.
		3: func(core.Agent) {},
		4: core.ForeverTier1ExpertiseBonus,
		5: func(core.Agent) {},
	},
})
