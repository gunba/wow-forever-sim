package paladin

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Templar's Bulwark is new in Forever. Client spell 1311015 confirms 110 mana, the 5 min
// cooldown, 8 sec and an absorb of 100% of maximum health.
// The 5 minute cooldown the sim assumed is confirmed: the BlizzCon "Paladin Class Change"
// talent slide reads "110 Mana, Instant, 5 min cooldown", and the same tooltip appears in
// Joardee's VOD. It applies Forbearance for 1 min, which is what sim/paladin/forbearance.go
// already grants.
func (paladin *Paladin) registerTemplarsBulwark() {
	if !paladin.Talents.TemplarsBulwark {
		return
	}

	actionID := core.ActionID{SpellID: 1311015}

	// Sacred Duty: 30 sec a rank, confirmed by the beta client's talent data.
	cooldown := time.Minute*5 - time.Second*30*time.Duration(paladin.Talents.SacredDuty)

	bulwark := paladin.RegisterSpell(core.SpellConfig{
		ActionID:         actionID,
		SpellSchool:      core.SpellSchoolHoly,
		Flags:            core.SpellFlagAPL | SpellFlag_Forbearance,
		ProcMask:         core.ProcMaskEmpty,
		DamageMultiplier: 1,
		Shield: core.ShieldConfig{SelfOnly: true, Aura: core.Aura{
			Label: "Templar's Bulwark", ActionID: actionID, Duration: 8 * time.Second,
		}},

		ManaCost: core.ManaCostOptions{
			FlatCost:   110,
			Multiplier: paladin.benediction(),
		},

		Cast: core.CastConfig{
			// SpellCooldowns 1311015: no StartRecoveryTime / GCD category.
			DefaultCast: core.Cast{},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    paladin.NewTimer(),
				Duration: cooldown,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			spell.SelfShield().Apply(sim, paladin.MaxHealth())
		},
	})

	paladin.AddMajorCooldown(core.MajorCooldown{
		Spell: bulwark,
		Type:  core.CooldownTypeSurvival,
	})
}
