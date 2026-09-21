package common

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

func init() {
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
