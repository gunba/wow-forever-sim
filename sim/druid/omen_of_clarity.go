package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (druid *Druid) registerOmenOfClarity() {
	if !druid.Env.IsForever() || druid.Level < 20 {
		return
	}

	// Provisional client-data model: 16864 lists 100% proc chance, a 10-second
	// ICD and no PPM entry. A server-side chance override has not been measured.
	// 16870 has one charge, 15-second duration and a -100% resource-cost mod.
	// See assets/db_inputs/forever_effect_audit.json and in-game check T25.
	icd := core.Cooldown{Timer: druid.NewTimer(), Duration: 10 * time.Second}
	eligible := map[*core.Spell]bool{}
	var grantSpell *core.Spell
	var grantAt time.Duration
	druid.OnSpellRegistered(func(spell *core.Spell) {
		// Wrath and Faerie Fire are absent from 16870's affected family mask.
		if spell.SpellCode != SpellCode_DruidWrath && spell.SpellCode != SpellCode_DruidFaerieFire &&
			spell.Cost != nil && spell.Cost.BaseCost > 0 &&
			!spell.Flags.Matches(core.SpellFlagPassiveSpell) &&
			spell.ProcMask.Matches(core.ProcMaskMeleeSpecial|core.ProcMaskSpellDamage|core.ProcMaskSpellHealing) {
			eligible[spell] = true
		}
	})
	consume := func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
		if !eligible[spell] || spell.CurCast.Cost != 0 {
			return
		}
		// Do not consume the proc this same event just granted/refreshed.
		// A different instant action at the same timestamp must still consume it.
		if grantSpell == spell && grantAt == sim.CurrentTime {
			return
		}
		aura.Deactivate(sim)
	}
	druid.ClearcastingAura = druid.RegisterAura(core.Aura{
		Label:    "Clearcasting",
		ActionID: core.ActionID{SpellID: 16870},
		Duration: 15 * time.Second,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			for spell := range eligible {
				spell.Cost.Multiplier -= 100
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			for spell := range eligible {
				spell.Cost.Multiplier += 100
			}
		},
		OnCastComplete: consume,
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			// Maul replaces a swing and deliberately has no cast-complete callback.
			if spell.Flags.Matches(core.SpellFlagNoOnCastComplete) {
				consume(aura, sim, spell)
			}
		},
	})

	proc := func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		if !result.Landed() || spell.Flags.Matches(core.SpellFlagPassiveSpell) ||
			!spell.ProcMask.Matches(core.ProcMaskMelee|core.ProcMaskSpellDamage|core.ProcMaskSpellHealing) ||
			!icd.IsReady(sim) {
			return
		}
		icd.Use(sim)
		grantSpell, grantAt = spell, sim.CurrentTime
		druid.ClearcastingAura.Activate(sim)
	}
	core.MakePermanent(druid.RegisterAura(core.Aura{
		Label:    "Omen of Clarity",
		ActionID: core.ActionID{SpellID: 16864},
		Icd:      &icd,
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			proc(sim, spell, result)
		},
		OnHealDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			proc(sim, spell, result)
		},
	}))
}
