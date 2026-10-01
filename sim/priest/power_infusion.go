package priest

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// The Priest owns the cost/cooldown; the selected friendly player owns the aura.
// An omitted option keeps the self-target default. An explicit unassigned or
// invalid recipient disables the owned cooldown, as in the raid assignment UI.
func (priest *Priest) registerPowerInfusionCD() {
	if !priest.Talents.PowerInfusion {
		return
	}

	actionID := core.ActionID{SpellID: 10060, Tag: priest.Index}
	recipient := &priest.Unit
	if priest.Env.IsForever() && priest.PowerInfusionTarget != nil {
		if priest.PowerInfusionTarget.Type != proto.UnitReference_Player && priest.PowerInfusionTarget.Type != proto.UnitReference_Self {
			return
		}
		recipient = priest.GetUnit(priest.PowerInfusionTarget)
		if recipient == nil || recipient.Type != core.PlayerUnit {
			return
		}
	}
	powerInfusionAura := core.PowerInfusionAura(recipient, actionID.Tag)

	piSpell := priest.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    SpellFlagPriest | core.SpellFlagHelpful | core.SpellFlagAPL,

		// 20% of base mana and a 3 min cooldown in the Forever beta client, as in Classic.
		ManaCost: core.ManaCostOptions{
			BaseCost: 0.20,
		},
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    priest.NewTimer(),
				Duration: core.PowerInfusionCD,
			},
		},

		ExtraCastCondition: func(sim *core.Simulation, _ *core.Unit) bool {
			return !priest.Env.IsForever() || (recipient.IsActive() && !recipient.HasActiveAuraWithTag(core.PowerInfusionAuraTag))
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			powerInfusionAura.Activate(sim)
		},
	})

	priest.AddMajorCooldown(core.MajorCooldown{
		Spell:    piSpell,
		Priority: core.CooldownPriorityDefault,
		Type:     core.CooldownTypeDPS,
	})
}
