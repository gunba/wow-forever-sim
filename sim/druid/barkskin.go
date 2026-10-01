package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

func (druid *Druid) registerBarkskinCD() {
	if !druid.Env.IsForever() && !druid.InForm(Bear) {
		return
	}

	actionId := core.ActionID{SpellID: 22812}

	// Beta client 1.60.1.69893: 20% less Physical damage taken for 15 sec, no cost and no cast time penalty.
	druid.BarkskinAura = druid.RegisterAura(core.Aura{
		Label:    "Barkskin",
		ActionID: actionId,
		Duration: time.Second * 15,
	}).AttachMultiplicativePseudoStatBuff(&druid.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexPhysical], 0.8)

	cast := core.Cast{}
	school := core.SpellSchoolNone
	if druid.Env.IsForever() {
		cast.GCD = core.GCDDefault
		school = core.SpellSchoolNature
		druid.BarkskinAura.AttachAdditivePseudoStatBuff(&druid.PseudoStats.SpellPushbackReduction, 1)
	}
	druid.Barkskin = druid.RegisterSpell(Any, core.SpellConfig{
		ActionID:      actionId,
		SpellSchool:   school,
		RequiredLevel: 44,
		Flags:         core.SpellFlagAPL,
		Cast: core.CastConfig{
			DefaultCast: cast,
			CD: core.Cooldown{
				Timer:    druid.NewTimer(),
				Duration: time.Second * 60,
			},
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			druid.BarkskinAura.Activate(sim)
			if !druid.Env.IsForever() {
				druid.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime, false)
			}
		},
	})

	druid.AddMajorCooldown(core.MajorCooldown{
		Spell: druid.Barkskin.Spell,
		Type:  core.CooldownTypeSurvival,
	})
}
