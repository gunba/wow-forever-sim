package priest

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"time"
)

var ItemSetVestmentsOfConviction = core.NewItemSet(core.ItemSet{
	ID: 2104, Name: "Vestments of Conviction",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(a core.Agent) { a.GetCharacter().AddStat(stats.Spirit, 10) },
		// Fear Ward is not modeled.
		3: func(core.Agent) {},
		4: func(a core.Agent) {
			a.GetCharacter().AddStats(stats.Stats{stats.HealingPower: 26, stats.SpellDamage: 9})
		},
		// Prayer of Mending is not modeled.
		5: core.ForeverTier1CooldownBonus(time.Second, func(s *core.Spell) bool { return s.SpellCode == SpellCode_PriestPenance }),
	},
})

var ItemSetRaimentsOfConviction = core.NewItemSet(core.ItemSet{
	ID: 2105, Name: "Raiments of Conviction",
	Bonuses: map[int32]core.ApplyEffect{
		2: core.ForeverTier1HitBonus,
		// Shackle Undead is not modeled.
		3: func(core.Agent) {},
		4: core.ForeverTier1CreatureBonus(1301087, proto.MobType_MobTypeUndead, 0, 21),
		5: core.ForeverTier1CooldownBonus(time.Minute, func(s *core.Spell) bool { return s.SpellCode == SpellCode_PriestDevouringPlague }),
	},
})
