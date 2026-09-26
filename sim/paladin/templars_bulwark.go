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

	var remainingAbsorb float64
	bulwarkAura := paladin.RegisterAura(core.Aura{
		Label:    "Templar's Bulwark",
		ActionID: actionID,
		Duration: time.Second * 8,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			remainingAbsorb = paladin.MaxHealth()
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			remainingAbsorb = 0
		},
	})
	paladin.AddDynamicDamageTakenModifier(func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		if !bulwarkAura.IsActive() || result.Damage <= 0 {
			return
		}
		absorbed := min(remainingAbsorb, result.Damage)
		result.Damage -= absorbed
		remainingAbsorb -= absorbed
		if remainingAbsorb <= 0 {
			bulwarkAura.Deactivate(sim)
		}
	})

	// Sacred Duty: 30 sec a rank, confirmed by the beta client's talent data.
	cooldown := time.Minute*5 - time.Second*30*time.Duration(paladin.Talents.SacredDuty)

	bulwark := paladin.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolHoly,
		Flags:       core.SpellFlagAPL | SpellFlag_Forbearance,

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

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			bulwarkAura.Activate(sim)
		},
	})

	paladin.AddMajorCooldown(core.MajorCooldown{
		Spell: bulwark,
		Type:  core.CooldownTypeSurvival,
	})
}
