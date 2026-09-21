package rogue

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

var ItemSetGrimstitchArmor = core.NewItemSet(core.ItemSet{
	ID: 2099, Name: "Grimstitch Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HitBonus,
		// Kidney Shot is not modeled.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301088, proto.MobType_MobTypeHumanoid, 36, 0),
		// Eviscerate cost is set at registration. Envenom is not modeled.
		5: func(core.Agent) {},
	},
})
