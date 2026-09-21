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
		5: core.ForeverTier1CooldownBonus(time.Second, func(s *core.Spell) bool {
			return s.SpellCode == SpellCode_HunterAimedShot || s.SpellCode == SpellCode_HunterMultiShot
		}),
	},
})
