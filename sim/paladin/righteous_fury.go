package paladin

import (
	"github.com/wowsims/classic/sim/core"
)

func (paladin *Paladin) registerRighteousFury() {
	if !paladin.Options.RighteousFury {
		return
	}
	actionID := core.ActionID{SpellID: 25780}

	// Beta client 1.60.1.70009: Holy threat +60%.
	paladin.OnSpellRegistered(func(spell *core.Spell) {
		if spell.SpellSchool.Matches(core.SpellSchoolHoly) {
			spell.ThreatMultiplier *= 1.6
		}
	})

	// Improved Righteous Fury no longer adds threat, it cuts damage taken while Righteous Fury is up.
	paladin.PseudoStats.DamageTakenMultiplier *= 1 - 0.02*float64(paladin.Talents.ImprovedRighteousFury)

	rfAura := core.MakePermanent(&core.Aura{Label: "Righteous Fury", ActionID: actionID})
	paladin.RegisterAura(*rfAura)
}
