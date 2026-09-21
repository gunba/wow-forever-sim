package druid

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

var ItemSetChampionsRefuge = core.NewItemSet(core.ItemSet{
	Name:            "Champion's Refuge",
	AlternativeName: "Champion's Sanctuary",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases healing done by up to 44 and damage done by up to 15 for all magical spells and
		// effects (467550).
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.HealingPower: 44,
				stats.SpellDamage:  15,
			})
		},
		// Increases your movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetLieutenantCommandersWildhide = core.NewItemSet(core.ItemSet{
	ID:   2081,
	Name: "Lieutenant Commander's Wildhide",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases damage and healing done by magical spells and effects by up to 23.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
		// Increases movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetLieutenantCommandersRefuge = core.NewItemSet(core.ItemSet{
	Name: "Lieutenant Commander's Refuge",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases healing done by up to 44 and damage done by up to 15 for all magical spells and
		// effects (467550).
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.HealingPower: 44,
				stats.SpellDamage:  15,
			})
		},
		// Increases your movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetFieldMarshalsSanctuary = core.NewItemSet(core.ItemSet{
	Name: "Field Marshal's Sanctuary",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Increases your movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// +40 Attack Power.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.AttackPower, 40)
			c.AddStat(stats.RangedAttackPower, 40)
		},
	},
})

var ItemSetWarlordsSanctuary = core.NewItemSet(core.ItemSet{
	Name: "Warlord's Sanctuary",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Increases your movement speed by 15% while in Bear, Cat, or Travel Form. Only active outdoors.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// +40 Attack Power.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.AttackPower, 40)
			c.AddStat(stats.RangedAttackPower, 40)
		},
	},
})
