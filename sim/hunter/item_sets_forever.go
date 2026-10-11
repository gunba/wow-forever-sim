package hunter

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"time"
)

var ItemSetWildstalkerArmor = core.NewItemSet(core.ItemSet{
	ID: 2101, Name: "Wildstalker Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HasteBonus,
		// Frost and Freezing Trap are not modeled; this does not affect damage traps.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301078, proto.MobType_MobTypeBeast, 36, 0),
		// The tooltip names Aimed Shot and Multi-Shot, but 1301252's cooldown modifier carries class
		// mask 131072 (Aimed Shot) only; Multi-Shot is 4096. Spell modifiers apply by mask.
		5: core.ForeverTier1CooldownBonus(time.Second, func(s *core.Spell) bool {
			return s.SpellCode == SpellCode_HunterAimedShot
		}),
	},
})
