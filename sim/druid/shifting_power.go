package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Client 70170: spell 1322605, effect 1358458, power 314986 and cooldown 101953.
func (druid *Druid) registerShiftingPowerSpell() {
	if !druid.Talents.ShiftingPower {
		return
	}
	actionID := core.ActionID{SpellID: 1322605}
	energy := druid.NewEnergyMetrics(actionID.WithTag(1))
	druid.ShiftingPower = druid.RegisterSpell(Cat, core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagAPL | core.SpellFlagHelpful,
		ManaCost: core.ManaCostOptions{
			BaseCost:   .55,
			Multiplier: 100 - 10*druid.Talents.NaturalShapeshifter,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: time.Second},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    druid.NewTimer(),
				Duration: time.Duration(16-4*druid.Talents.ImprovedShiftingPower) * time.Second,
			},
		},
		ExtraCastCondition: func(_ *core.Simulation, _ *core.Unit) bool {
			return druid.CurrentEnergy() < druid.MaxEnergy()
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			amount := 40.0
			if druid.Equipment.Head().ID == WolfsheadHelm {
				amount += 5
			}
			druid.AddEnergy(sim, amount, energy)
		},
	})
}
