package common

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

func init() {
	core.NewSimpleStatOffensiveTrinketEffect(272439,
		stats.Stats{stats.Spirit: 72}, 15*time.Second, 90*time.Second)
	core.NewSimpleStatOffensiveTrinketEffectWithOtherEffects(272438,
		stats.Stats{stats.MeleeCrit: 5, stats.SpellCrit: 5}, 20*time.Second, 90*time.Second,
		func(agent core.Agent) {
			aura := agent.GetCharacter().GetAura("ItemActive-272438")
			consume := func(aura *core.Aura, sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				if result.DidCrit() {
					aura.Deactivate(sim)
				}
			}
			aura.OnSpellHitDealt = consume
			aura.OnHealDealt = consume
		})
	// ItemEffect supplies the five-minute cooldown and 15-second shared category.
	// SpellEffect supplies the ordinary-area values; no snowy/volcanic bonus is
	// assumed for the neutral benchmark encounter.
	core.NewSimpleStatOffensiveTrinketEffect(249469,
		stats.Stats{stats.FrostPower: 29, stats.ShadowPower: 29},
		15*time.Second, 5*time.Minute)
	core.NewSimpleStatOffensiveTrinketEffect(249470,
		stats.Stats{stats.AttackPower: 55, stats.RangedAttackPower: 55},
		15*time.Second, 5*time.Minute)
}
