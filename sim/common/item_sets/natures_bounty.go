package item_sets

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// ApplyNaturesBounty is shared by Wildheart and Feralheart. Forever's 450608
// describes three distinct triggers. The script's proc rates are unresolved;
// retain the existing 2% assumption rather than treating its dummy aura's 100%
// ProcChance as a guaranteed resource return.
func ApplyNaturesBounty(agent core.Agent) {
	c := agent.GetCharacter()
	if c.GetAura("Nature's Bounty (Mana)") != nil {
		return
	}
	actionID := core.ActionID{SpellID: 450608}
	manaMetrics := c.NewManaMetrics(core.ActionID{SpellID: 27782})
	rageMetrics := c.NewRageMetrics(core.ActionID{SpellID: 27783})
	energyMetrics := c.NewEnergyMetrics(core.ActionID{SpellID: 27784})
	energyRestore := c.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 27784},
		SpellSchool: core.SpellSchoolNature,
		Flags:       core.SpellFlagPassiveSpell | core.SpellFlagNoOnCastComplete,
		Hot: core.DotConfig{
			Aura: core.Aura{Label: "Nature's Bounty Energy"},
			// Client 70009: periodic energize 4, period 1000 ms; duration 5s.
			NumberOfTicks: 5,
			TickLength:    time.Second,
			OnTick: func(sim *core.Simulation, _ *core.Unit, _ *core.Dot) {
				c.AddEnergy(sim, 4, energyMetrics)
			},
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			spell.Hot(&c.Unit).Apply(sim)
		},
	})
	core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
		ActionID: actionID, Name: "Nature's Bounty (Mana)",
		Callback:   core.CallbackOnCastComplete,
		ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
		ProcChance: 0.02,
		Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
			if c.HasManaBar() {
				c.AddMana(sim, 200, manaMetrics)
			}
		},
	})
	core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
		ActionID: actionID, Name: "Nature's Bounty (Energy)",
		Callback: core.CallbackOnSpellHitDealt,
		// Proc mask 0x1402c includes outgoing white melee, not melee specials.
		ProcMask:   core.ProcMaskMeleeWhiteHit,
		Outcome:    core.OutcomeLanded,
		ProcChance: 0.02,
		Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
			if c.HasEnergyBar() {
				energyRestore.Cast(sim, &c.Unit)
			}
		},
	})
	core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
		ActionID: actionID, Name: "Nature's Bounty (Rage)",
		Callback:   core.CallbackOnSpellHitTaken,
		ProcMask:   core.ProcMaskMelee,
		Outcome:    core.OutcomeLanded,
		ProcChance: 0.02,
		Handler: func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
			if c.HasRageBar() {
				c.AddRage(sim, 10, rageMetrics)
			}
		},
	})
}
