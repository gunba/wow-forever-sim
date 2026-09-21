package shaman

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"time"
)

var ItemSetSpiritcallersRage = core.NewItemSet(core.ItemSet{
	ID: 2110, Name: "The Spiritcaller's Rage",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HasteBonus,
		// Beneficial totem range does not constrain the stationary benchmark.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301090, proto.MobType_MobTypeElemental, 36, 0),
		5: core.ForeverTier1CooldownBonus(500*time.Millisecond, func(s *core.Spell) bool { return s.SpellCode == SpellCode_ShamanStormstrike }),
	},
})

var ItemSetSpiritcallersStorm = core.NewItemSet(core.ItemSet{
	ID: 2111, Name: "The Spiritcaller's Storm",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HitBonus,
		// Earthbind Totem is not modeled.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301092, proto.MobType_MobTypeElemental, 0, 21),
		5: core.ForeverTier1CooldownBonus(time.Second, func(s *core.Spell) bool { return s.SpellCode == SpellCode_ShamanLavaBurst }),
	},
})
