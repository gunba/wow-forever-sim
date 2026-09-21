package priest

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func (priest *Priest) registerDarkSacrifice() {
	if !priest.Env.IsForever() || priest.Race != proto.Race_RaceUndead {
		return
	}

	// SpellEffect and SpellLevels, beta 1.60.1.69893. Each rank grows by
	// one point per level, capped eight levels after it is learned.
	ranks := []struct {
		id    int32
		level int32
		tick  float64
	}{
		{1277324, 20, 80},
		{1277325, 30, 136},
		{1277326, 40, 196},
		{1277327, 50, 258},
		{1277328, 60, 320},
	}
	timer := priest.NewTimer()
	var highest *core.Spell
	var totalConversion float64
	for index, rank := range ranks {
		if priest.Level < rank.level {
			continue
		}
		amount := rank.tick + float64(min(priest.Level-rank.level, 8))
		actionID := core.ActionID{SpellID: rank.id}
		manaMetrics := priest.NewManaMetrics(actionID)
		aura := priest.RegisterAura(core.Aura{
			Label:    fmt.Sprintf("Dark Sacrifice (Rank %d)", index+1),
			ActionID: actionID,
			Duration: 15 * time.Second,
			OnGain: func(_ *core.Aura, sim *core.Simulation) {
				// Use the client base amounts. Self-damage modifier interactions
				// remain a server-side verification gap.
				// The client marks it non-cancelable. Schedule all five transfers
				// independently of APL aura cancellation.
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   3 * time.Second,
					NumTicks: 5,
					OnAction: func(sim *core.Simulation) {
						priest.RemoveHealth(sim, amount)
						priest.AddMana(sim, amount, manaMetrics)
					},
				})
			},
		})
		highest = priest.RegisterSpell(core.SpellConfig{
			ActionID:      actionID,
			SpellSchool:   core.SpellSchoolShadow,
			Flags:         core.SpellFlagAPL | core.SpellFlagHelpful | core.SpellFlagNoOnCastComplete,
			Rank:          index + 1,
			RequiredLevel: int(rank.level),
			Cast: core.CastConfig{
				DefaultCast: core.Cast{GCD: core.GCDDefault},
				CD:          core.Cooldown{Timer: timer, Duration: 10 * time.Minute},
			},
			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
				aura.Activate(sim)
			},
		})
		totalConversion = amount * 5
	}
	if highest != nil {
		priest.AddMajorCooldown(core.MajorCooldown{
			Spell: highest,
			Type:  core.CooldownTypeMana,
			ShouldActivate: func(sim *core.Simulation, _ *core.Character) bool {
				return priest.CurrentHealth() > totalConversion &&
					priest.MaxMana()-priest.CurrentMana() >= totalConversion &&
					sim.GetRemainingDuration() > 15*time.Second
			},
		})
	}
}
