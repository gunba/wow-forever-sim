package common

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

func init() {
	// Item 272440 / spell 1291105: ordinary outside-city raid use. The
	// Stronghold/City doubling is not part of the supported raid encounters.
	registerDefender := core.MakeTemporaryStatsOnUseCDRegistration(
		"Defender's Grip Stabilizer",
		stats.Stats{stats.Block: 8 * core.BlockRatingPerBlockChance},
		15*time.Second,
		core.SpellConfig{
			ActionID: core.ActionID{ItemID: 272440},
			Flags:    core.SpellFlagDefensiveEquipment | core.SpellFlagNoOnCastComplete,
		},
		func(character *core.Character) core.Cooldown {
			return core.Cooldown{Timer: character.NewTimer(), Duration: 6 * time.Minute}
		},
		func(character *core.Character) core.Cooldown {
			// Item category 1141 is shared with the other stat-use trinkets;
			// defensive equipment flags do not create a separate category.
			return core.Cooldown{Timer: character.GetOffensiveTrinketCD(), Duration: 15 * time.Second}
		},
	)
	core.NewItemEffect(272440, func(agent core.Agent) {
		if agent.GetCharacter().Env.IsForever() {
			registerDefender(agent)
		}
	})
}
