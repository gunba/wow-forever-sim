package warrior

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"time"
)

var ItemSetBattlegearOfGlory = core.NewItemSet(core.ItemSet{
	ID: 2102, Name: "Battlegear of Glory",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HasteBonus,
		// Intimidating Shout is not modeled.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301095, proto.MobType_MobTypeHumanoid, 36, 0),
		5: core.ForeverTier1CooldownBonus(30*time.Second, func(s *core.Spell) bool { return s.SpellID == 1719 }),
	},
})
