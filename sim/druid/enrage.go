package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

// Generates 20 Rage over 10 sec, reducing base armor by 27% in Bear or 16% in Dire Bear.
// Forever adds 10 Rage up front
// (beta client 1.60.1.69893, effect 1 of 5229).
func (druid *Druid) registerEnrageSpell() {
	actionID := core.ActionID{SpellID: 5229}
	rageMetrics := druid.NewRageMetrics(actionID)

	armorMultiplier := 1 - 0.27
	if druid.BearForm != nil && druid.BearForm.ActionID.SpellID == 9634 {
		armorMultiplier = 1 - 0.16
	}

	druid.EnrageAura = druid.RegisterAura(core.Aura{
		Label:    "Enrage",
		ActionID: actionID,
		Duration: time.Second * 10,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			druid.ApplyDynamicEquipScaling(sim, stats.Armor, armorMultiplier)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			druid.RemoveDynamicEquipScaling(sim, stats.Armor, armorMultiplier)
		},
	})

	druid.Enrage = druid.RegisterSpell(Bear, core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagAPL,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    druid.NewTimer(),
				Duration: time.Minute,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			if druid.Env.IsForever() {
				rage := 10.0
				if druid.Equipment.Head().ID == WolfsheadHelm {
					rage += 5
				}
				druid.AddRage(sim, rage, rageMetrics)
			}

			core.StartPeriodicAction(sim, core.PeriodicActionOptions{
				NumTicks: 10,
				Period:   time.Second,
				OnAction: func(sim *core.Simulation) {
					if druid.EnrageAura.IsActive() {
						druid.AddRage(sim, 2, rageMetrics)
					}
				},
			})

			druid.EnrageAura.Activate(sim)
		},
	})

	druid.AddMajorCooldown(core.MajorCooldown{
		Spell: druid.Enrage.Spell,
		Type:  core.CooldownTypeDPS,
	})
}
