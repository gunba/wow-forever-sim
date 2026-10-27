package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

const ShoutExpirationThreshold = time.Second * 3

func (warrior *Warrior) newShoutSpellConfig(actionID core.ActionID, rank int32, allyAuras core.AuraArray) *WarriorSpell {
	// Use extra hits to simulate buffing your party for threat
	extraHits := 5 - len(allyAuras)

	return warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID: actionID,
		// The hits on the target only stand in for the threat of buffing the party, so they do not start the fight.
		Flags: core.SpellFlagNoOnCastComplete | core.SpellFlagAPL | core.SpellFlagHelpful | core.SpellFlagPullSafe,

		RageCost: core.RageCostOptions{
			Cost: 10 * core.TernaryFloat64(warrior.Env.IsForever(), 1-.05*float64(warrior.Talents.BoomingVoice), 1),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		// Forever buffs generate no threat (Blizzard: Not A Bug, forever-bugs#64; Battle Shout measured 0).
		FlatThreatBonus: core.TernaryFloat64(warrior.Env.IsForever(), 0, float64(core.BattleShoutLevel[rank])),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for _, aura := range allyAuras {
				spell.CalcAndDealOutcome(sim, aura.Unit, spell.OutcomeAlwaysHit)
				aura.Activate(sim)
			}

			for i := 0; i < extraHits; i++ {
				spell.CalcAndDealOutcome(sim, target, spell.OutcomeAlwaysHit)
			}
		},

		RelatedAuras: []core.AuraArray{allyAuras},
	})
}

func (warrior *Warrior) registerBattleShout() {
	trained := warrior.rank(ranks.BattleShout)
	if trained == nil {
		return
	}
	topRank := core.TernaryInt32(core.IncludeAQ, 7, 6)
	rank := min(trained.Rank, topRank)
	actionId := core.BattleShoutSpellId[topRank]
	has3pcWrath := warrior.HasSetBonus(ItemSetBattleGearOfWrath, 3)

	warrior.BattleShout = warrior.newShoutSpellConfig(core.ActionID{SpellID: actionId}, rank, warrior.NewPartyAuraArray(func(unit *core.Unit) *core.Aura {
		// Improved Battle Shout is gone from the Forever tree and did not become baseline: the beta
		// client's rank 7 gives 139 attack power for 3 min, 60% of Classic's 232 for 2 min. Booming
		// Voice widens the radius and, since October 1, discounts the Rage cost. The base value and duration live in core.BattleShoutAura.
		return core.BattleShoutAuraRank(unit, rank, 0, 0, has3pcWrath)
	}))
}

func (warrior *Warrior) registerShouts() {
	warrior.registerBattleShout()
}
