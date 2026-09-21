package druid

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"time"
)

var ItemSetGrovekeeperEclipse = core.NewItemSet(core.ItemSet{
	ID: 2114, Name: "Grovekeeper Eclipse",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HitBonus,
		// Nature's Grasp is not modeled.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301073, proto.MobType_MobTypeDemon, 0, 21),
		// Insect Swarm duration is set at registration.
		5: func(core.Agent) {},
	},
})

var ItemSetGrovekeeperFerocity = core.NewItemSet(core.ItemSet{
	ID: 2115, Name: "Grovekeeper Ferocity",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HasteBonus,
		// Hibernate is not modeled.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301075, proto.MobType_MobTypeDemon, 36, 0),
		5: core.ForeverTier1CooldownBonus(3*time.Second, func(s *core.Spell) bool {
			return s.SpellID == 5217 || s.SpellID == 6793 || s.SpellID == 9845 || s.SpellID == 9846
		}),
	},
})
