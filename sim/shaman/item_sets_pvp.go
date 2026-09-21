package shaman

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

var ItemSetChampionsEarthshaker = core.NewItemSet(core.ItemSet{
	Name: "Champion's Earthshaker",
	Bonuses: map[int32]core.ApplyEffect{
		// +40 Attack Power.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
		// Improves your chance to get a critical strike with all Shock spells by 2%.
		4: func(agent core.Agent) {
			shaman := agent.(ShamanAgent).GetShaman()
			shaman.GetOrRegisterAura(core.Aura{
				Label:    "Shaman Shock Crit Bonus",
				ActionID: core.ActionID{SpellID: 22804},
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range core.Flatten([][]*core.Spell{shaman.EarthShock, shaman.FlameShock, shaman.FrostShock}) {
						if spell != nil {
							spell.BonusCritRating += 2 * core.CritRatingPerCritChance
						}
					}
				},
			})
		},
		// +20 Stamina (14467, was 15).
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetChampionsStormcaller = core.NewItemSet(core.ItemSet{
	Name: "Champion's Stormcaller",
	Bonuses: map[int32]core.ApplyEffect{
		// +40 Attack Power.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
		// Improves your chance to get a critical strike with all Shock spells by 2%.
		4: func(agent core.Agent) {
			shaman := agent.(ShamanAgent).GetShaman()
			shaman.GetOrRegisterAura(core.Aura{
				Label:    "Shaman Shock Crit Bonus",
				ActionID: core.ActionID{SpellID: 22804},
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range core.Flatten([][]*core.Spell{shaman.EarthShock, shaman.FlameShock, shaman.FrostShock}) {
						if spell != nil {
							spell.BonusCritRating += 2 * core.CritRatingPerCritChance
						}
					}
				},
			})
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetChampionsThunderfist = core.NewItemSet(core.ItemSet{
	ID:   2094,
	Name: "Champion's Thunderfist",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases damage and healing done by magical spells and effects by up to 23.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
		// Improves your chance to get a critical strike with all Shock spells by 2%.
		4: func(agent core.Agent) {
			shaman := agent.(ShamanAgent).GetShaman()
			shaman.GetOrRegisterAura(core.Aura{
				Label:    "Shaman Shock Crit Bonus",
				ActionID: core.ActionID{SpellID: 22804},
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range core.Flatten([][]*core.Spell{shaman.EarthShock, shaman.FlameShock, shaman.FrostShock}) {
						if spell != nil {
							spell.BonusCritRating += 2 * core.CritRatingPerCritChance
						}
					}
				},
			})
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetWarlordsEarthshaker = core.NewItemSet(core.ItemSet{
	Name: "Warlord's Earthshaker",
	Bonuses: map[int32]core.ApplyEffect{
		// +20 Stamina.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
		// Improves your chance to get a critical strike with all Shock spells by 2%.
		4: func(agent core.Agent) {
			shaman := agent.(ShamanAgent).GetShaman()
			shaman.GetOrRegisterAura(core.Aura{
				Label:    "Shaman Shock Crit Bonus",
				ActionID: core.ActionID{SpellID: 22804},
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range core.Flatten([][]*core.Spell{shaman.EarthShock, shaman.FlameShock, shaman.FrostShock}) {
						if spell != nil {
							spell.BonusCritRating += 2 * core.CritRatingPerCritChance
						}
					}
				},
			})
		},
		// +40 Attack Power.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
	},
})
