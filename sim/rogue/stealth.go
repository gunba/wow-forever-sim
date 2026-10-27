package rogue

import (
	"github.com/wowsims/classic/sim/core"
)

// Stealth's ranks differ only in movement speed.
var stealthSpellID = [...]int32{1784, 1785, 1786, 1787}
var stealthLevel = [...]int32{1, 20, 34, 50}

// registerStealthSpell lets the Rogue enter Stealth out of combat, which in the simulation means before the pull;
// during the fight only Vanish leads into it. Every learned rank is castable.
func (rogue *Rogue) registerStealthSpell() {
	for i := range stealthSpellID {
		if rogue.Level < stealthLevel[i] {
			break
		}
		spell := rogue.RegisterSpell(core.SpellConfig{
			ActionID:      core.ActionID{SpellID: stealthSpellID[i]},
			Flags:         core.SpellFlagAPL,
			RequiredLevel: int(stealthLevel[i]),
			Cast: core.CastConfig{
				IgnoreHaste: true,
			},
			ExtraCastCondition: func(sim *core.Simulation, _ *core.Unit) bool {
				return sim.CurrentTime < 0 && !rogue.StealthAura.IsActive()
			},
			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
				rogue.AutoAttacks.CancelAutoSwing(sim)
				rogue.StealthAura.Activate(sim)
			},
		})
		rogue.Stealth = spell
	}
}

func (rogue *Rogue) registerStealthAura() {
	rogue.StealthAura = rogue.RegisterAura(core.Aura{
		Label:    "Stealth",
		ActionID: core.ActionID{SpellID: 1787},
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			// Stealth triggered auras
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
		},
		// Stealth breaks on damage taken (if not absorbed)
		// This may be desirable later, but not applicable currently
	})
}
