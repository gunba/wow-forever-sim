package paladin

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

// Libram IDs
const (
	SanctifiedOrb      = 20512
	LibramOfHope       = 22401
	LibramOfFervor     = 23203
	LibramOfInvocation = 249442
	LibramOfInfusion   = 279248
)

func init() {
	core.NewSimpleStatOffensiveTrinketEffect(SanctifiedOrb, stats.Stats{stats.MeleeCrit: 3 * core.CritRatingPerCritChance, stats.SpellCrit: 3 * core.CritRatingPerCritChance}, time.Second*25, time.Minute*3)
	core.NewItemEffect(LibramOfInfusion, func(agent core.Agent) {
		paladin := agent.(PaladinAgent).GetPaladin()
		if paladin.Env.IsForever() {
			paladin.OnSpellRegistered(func(spell *core.Spell) {
				if spell.SpellCode == SpellCode_PaladinHolyShock {
					spell.BonusCritRating += 6 * core.SpellCritRatingPerCritChance
				}
			})
		}
	})
}

func (paladin *Paladin) sealCostMultiplier() int32 {
	cost := paladin.benediction()
	if paladin.Env.IsForever() && paladin.Ranged().ID == LibramOfInvocation {
		cost -= 5
	}
	return cost
}
