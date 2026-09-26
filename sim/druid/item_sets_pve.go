package druid

import (
	"github.com/wowsims/classic/sim/common/item_sets"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

var ItemSetFeralheartRaiment = core.NewItemSet(core.ItemSet{
	Name: "Feralheart Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		// (2) Set : +8 All Resistances.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddResistances(8)
		},
		// (3) Set : Restores 8 mana per 5 sec.
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MP5, 8)
		},
		4: item_sets.ApplyNaturesBounty,
		// (5) Set : Wild Heart, a movement speed burst when struck.
		5: func(_ core.Agent) {},
		// (6) Set : +23 damage and healing done by magical spells and effects, and +40 Attack Power.
		// Classic gave 15 and 26, at eight and six pieces respectively.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.SpellDamage:       23,
				stats.HealingPower:      23,
				stats.AttackPower:       40,
				stats.RangedAttackPower: 40,
			})
		},
	},
})

var ItemSetCenarionRaiment = core.NewItemSet(core.ItemSet{
	Name: "Cenarion Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		// (3) Set : Damage dealt by Thorns increased by 4 and duration increased by 50%.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// (5) Set : Improves your chance to get a critical strike with spells by 2%.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellCrit, 2*core.SpellCritRatingPerCritChance)
		},
		// (8) Set : Reduces the cooldown of your Tranquility and Hurricane spells by 50%.
		8: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

var ItemSetStormrageRaiment = core.NewItemSet(core.ItemSet{
	Name: "Stormrage Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		// (3) Set : Allows 15% of your Mana regeneration to continue while casting.
		3: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.PseudoStats.SpiritRegenRateCasting += .15
		},
		// (5) Set : Reduces the casting time of your Regrowth spell by 0.2 sec.
		5: func(agent core.Agent) {
			// Nothing to do.
		},
		// (8) Set : Increases the duration of your Rejuvenation spell by 3 sec.
		8: func(agent core.Agent) {
			// Nothing to do.
		},
	},
})

var ItemSetSymbolsOfUnendingLife = core.NewItemSet(core.ItemSet{
	Name: "Symbols of Unending Life",
	Bonuses: map[int32]core.ApplyEffect{
		// (3) Set : Your finishing moves now refund 30 energy on a Miss, Dodge, Block, or Parry.
		3: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 26107}
			energyMetrics := c.NewEnergyMetrics(actionID)
			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				Name:     "Symbols of Unending Life Finisher Bonus",
				Callback: core.CallbackOnSpellHitDealt,
				Outcome:  core.OutcomeMiss | core.OutcomeDodge | core.OutcomeBlock | core.OutcomeParry,
				Handler: func(sim *core.Simulation, spell *core.Spell, _ *core.SpellResult) {
					if spell.SpellCode == SpellCode_DruidFerociousBite || spell.SpellCode == SpellCode_DruidRip {
						c.AddEnergy(sim, 30, energyMetrics)
					}
				},
			})
		},
	},
})
