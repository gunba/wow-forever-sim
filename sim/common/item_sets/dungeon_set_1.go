package item_sets

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

// Forever rebuilt dungeon set 1. Classic's 2/4/6/8 became 2/3/4/5/6, so the whole set pays out at
// six pieces instead of eight, and every 4 piece is now a PvP break the sim has no use for. Read
// from beta client 1.60.1.69893 (ItemSetSpell), Era 1.15.9.69722 for what it replaced.
//
//	2  Increased All Resist 08 (18679), every set, where Classic gave +200 armor
//	3  a stat block: 9336 +30 attack power, 9346 +18 spell damage and healing, 13198 +15 Strength
//	4  break a snare / root / disarm, a movement speed burst, or punish the attacker - not simulated
//	5  the proc that used to sit at 6
//	6  18378 +8 mana every 5 sec, or Vitality (21347) on the two rage/energy sets
//
// Two proc values moved with it: Crusader's Wrath and The Furious Storm are 65 spell power (were 95),
// and Rogue Armor Energize gives 20 energy (was 35).

// resistBonus is every set's 2 piece in Forever.
func resistBonus(agent core.Agent) {
	agent.GetCharacter().AddResistances(8)
}

// manaPerFive is the 6 piece on the seven caster-facing sets (18378).
func manaPerFive(agent core.Agent) {
	agent.GetCharacter().AddStat(stats.MP5, 8)
}

// vitality is the 6 piece on Shadowcraft and Battlegear of Valor (21347), 15 health every 5 sec.
// The sim has no health regen stat and a boss fight never idles, so it is left unmodelled.
func vitality(_ core.Agent) {}

// suddenInsight is the 5 piece shared by Magister's Regalia and Vestments of the Devout: a
// spellcast has a 5% chance to restore 200 mana. The two sets carry it under their own ids, so the
// ActionID is built at the call site and the spell manifest can find it in the source.
func suddenInsight(agent core.Agent, actionID core.ActionID) {
	c := agent.GetCharacter()
	manaMetrics := c.NewManaMetrics(actionID)

	core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
		ActionID:   actionID,
		Name:       "Sudden Insight",
		Callback:   core.CallbackOnCastComplete,
		ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
		ProcChance: 0.05,
		Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
			if c.HasManaBar() {
				c.AddMana(sim, 200, manaMetrics)
			}
		},
	})
}

var ItemSetWildheartRaiment = core.NewItemSet(core.ItemSet{
	Name: "Wildheart Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +30 Attack Power, and +18 damage and healing done by magical spells and effects.
		3: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       30,
				stats.RangedAttackPower: 30,
				stats.SpellDamage:       18,
				stats.HealingPower:      18,
			})
		},
		// Wild Heart: a movement speed burst when struck.
		4: func(_ core.Agent) {},
		5: ApplyNaturesBounty,
		6: manaPerFive,
	},
})

var ItemSetBeaststalkerArmor = core.NewItemSet(core.ItemSet{
	Name: "Beaststalker Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +30 Attack Power.
		3: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       30,
				stats.RangedAttackPower: 30,
			})
		},
		// Beast Unleashed: breaks a Root when struck.
		4: func(_ core.Agent) {},
		// Melee and ranged autoattacks have a 5% chance to restore 200 mana. Classic had this at
		// 6 pieces, 4%, and ranged only.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 450577}
			manaMetrics := c.NewManaMetrics(actionID)

			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				ActionID:   actionID,
				Name:       "Hunter Armor Energize",
				Callback:   core.CallbackOnSpellHitDealt,
				Outcome:    core.OutcomeLanded,
				ProcMask:   core.ProcMaskWhiteHit,
				ProcChance: 0.05,
				Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
					if c.HasManaBar() {
						c.AddMana(sim, 200, manaMetrics)
					}
				},
			})
		},
		6: manaPerFive,
	},
})

var ItemSetMagistersRegalia = core.NewItemSet(core.ItemSet{
	Name: "Magister's Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +18 damage and healing done by magical spells and effects.
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 18)
		},
		// Freeze: roots the attacker when struck.
		4: func(_ core.Agent) {},
		// Spellcasts have a 5% chance to restore 200 mana.
		5: func(agent core.Agent) {
			suddenInsight(agent, core.ActionID{SpellID: 450527})
		},
		6: manaPerFive,
	},
})

var ItemSetLightforgeArmor = core.NewItemSet(core.ItemSet{
	Name: "Lightforge Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +18 damage and healing done by magical spells and effects.
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 18)
		},
		// Rebuke: silences a caster that lands a harmful spell on you.
		4: func(_ core.Agent) {},
		// Crusader's Wrath, moved down from 6 pieces and cut from 95 to 65 spell power.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 450625}
			procAura := c.NewTemporaryStatsAura("Crusader's Wrath", core.ActionID{SpellID: 27499}, stats.Stats{stats.SpellPower: 65}, time.Second*10)

			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				ActionID: actionID,
				Name:     "Item - Crusader's Wrath Proc - Lightforge Armor",
				Callback: core.CallbackOnSpellHitDealt,
				Outcome:  core.OutcomeLanded,
				ProcMask: core.ProcMaskMeleeWhiteHit,
				PPM:      3, // The client stores no chance for it; 3 PPM is what the sim already used.
				Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
					procAura.Activate(sim)
				},
			})
		},
		6: manaPerFive,
	},
})

var ItemSetVestmentsOfTheDevout = core.NewItemSet(core.ItemSet{
	Name: "Vestments of the Devout",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +18 damage and healing done by magical spells and effects.
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 18)
		},
		// Divine Protection: a damage absorb when struck.
		4: func(_ core.Agent) {},
		// Spellcasts have a 5% chance to restore 200 mana.
		5: func(agent core.Agent) {
			suddenInsight(agent, core.ActionID{SpellID: 450576})
		},
		6: manaPerFive,
	},
})

var ItemSetShadowcraftArmor = core.NewItemSet(core.ItemSet{
	Name: "Shadowcraft Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +30 Attack Power.
		3: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.AttackPower:       30,
				stats.RangedAttackPower: 30,
			})
		},
		// Crafted Shadows: breaks a Snare when struck.
		4: func(_ core.Agent) {},
		// Chance on melee attack to restore energy, moved down from 6 pieces and cut 35 -> 20.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 27787}
			energyMetrics := c.NewEnergyMetrics(actionID)

			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				ActionID: actionID,
				Name:     "Rogue Armor Energize",
				Callback: core.CallbackOnSpellHitDealt,
				Outcome:  core.OutcomeLanded,
				ProcMask: core.ProcMaskMeleeWhiteHit,
				PPM:      1,
				Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
					if c.HasEnergyBar() {
						c.AddEnergy(sim, 20, energyMetrics)
					}
				},
			})
		},
		6: vitality,
	},
})

var ItemSetTheElements = core.NewItemSet(core.ItemSet{
	Name: "The Elements",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +18 damage and healing done by magical spells and effects.
		3: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStats(stats.Stats{
				stats.SpellDamage:  18,
				stats.HealingPower: 18,
			})
		},
		// Electrocute: disarms an attacker.
		4: func(_ core.Agent) {},
		// The Furious Storm, moved down from 6 pieces and cut from 95 to 65 spell power.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 450626}
			procAura := c.NewTemporaryStatsAura("The Furious Storm", core.ActionID{SpellID: 27775}, stats.Stats{stats.SpellPower: 65}, time.Second*10)

			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				ActionID:   actionID,
				Name:       "Item - The Furious Storm Proc",
				Callback:   core.CallbackOnCastComplete,
				ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
				ProcChance: 0.04, // No chance in the client; Classic's 4% is kept.
				Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
					procAura.Activate(sim)
				},
			})
		},
		6: manaPerFive,
	},
})

var ItemSetDreadmistRaiment = core.NewItemSet(core.ItemSet{
	Name: "Dreadmist Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +18 damage and healing done by magical spells and effects.
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 18)
		},
		// Corrupted Fear: the attacker flees when you are struck.
		4: func(_ core.Agent) {},
		// Spellcasts have a 5% chance to heal you for 270 to 330.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 450585}
			healthMetrics := c.NewHealthMetrics(actionID)

			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				ActionID:   actionID,
				Name:       "Dark Reward",
				Callback:   core.CallbackOnCastComplete,
				ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
				ProcChance: 0.05,
				Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
					c.GainHealth(sim, sim.Roll(270, 331), healthMetrics)
				},
			})
		},
		6: manaPerFive,
	},
})

var ItemSetBattlegearOfValor = core.NewItemSet(core.ItemSet{
	Name: "Battlegear of Valor",
	Bonuses: map[int32]core.ApplyEffect{
		2: resistBonus,
		// +15 Strength.
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.Strength, 15)
		},
		// Moment of Valor: breaks a Disarm when struck.
		4: func(_ core.Agent) {},
		// Warrior's Resolve, moved down from 6 pieces: it now returns 10 Rage with the heal.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			actionID := core.ActionID{SpellID: 450587}
			healthMetrics := c.NewHealthMetrics(core.ActionID{SpellID: 450589})
			rageMetrics := c.NewRageMetrics(core.ActionID{SpellID: 450589})

			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				ActionID: actionID,
				Name:     "Warrior's Resolve",
				Callback: core.CallbackOnSpellHitDealt,
				Outcome:  core.OutcomeLanded,
				ProcMask: core.ProcMaskMelee,
				PPM:      1,
				Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
					c.GainHealth(sim, sim.Roll(88, 133), healthMetrics)
					if c.HasRageBar() {
						c.AddRage(sim, 10, rageMetrics)
					}
				},
			})
		},
		6: vitality,
	},
})
