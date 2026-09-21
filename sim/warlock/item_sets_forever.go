package warlock

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

var ItemSetDemonheartRaiment = core.NewItemSet(core.ItemSet{
	ID: 2100, Name: "Demonheart Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HitBonus,
		// Banish is not modeled.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301094, proto.MobType_MobTypeDemon, 0, 21),
		// Life Tap's mana-only multiplier is applied in lifetap.go.
		5: func(core.Agent) {},
	},
})
