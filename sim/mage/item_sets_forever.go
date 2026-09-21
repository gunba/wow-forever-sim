package mage

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"time"
)

// ItemSetSpell / SpellEffect, beta 1.60.1.69893.
var ItemSetManaflareRegalia = core.NewItemSet(core.ItemSet{
	ID: 2098, Name: "Manaflare Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HitBonus,
		3: core.ForeverTier1CooldownBonus(5*time.Second, func(s *core.Spell) bool { return s.SpellID == 2139 }),
		4: core.ForeverTier1CreatureBonus(1301079, proto.MobType_MobTypeElemental, 0, 21),
		// Frostfire Bolt's three talent interactions are applied in talents.go.
		5: func(core.Agent) {},
	},
})
