package mage

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

type frostAreaRank struct {
	id                             int32
	level                          int
	mana, base, perLevel, variance float64
}

// Client 1.60.1.70009. Variance spreads the unscaled base; level growth
// adds to the midpoint and stops five levels after learning each rank.
var coneOfColdRanks = []frostAreaRank{
	{120, 26, 210, 97, .8, .09708737581968307},
	{8492, 34, 290, 144, 1, .09150326997041702},
	{10159, 42, 380, 204, 1.2, .09389671683311462},
	{10160, 50, 465, 267, 1.3, .09386281669139862},
	{10161, 58, 555, 340, 1.5, .08571428805589676},
}

var frostNovaRanks = []frostAreaRank{
	{122, 10, 55, 20, .5, .10000000149011612},
	{865, 26, 85, 34, .5, .11428570747375488},
	{6131, 40, 115, 53, .5, .1090909093618393},
	{10230, 54, 145, 73, .5, .1066666692495346},
}

func (mage *Mage) registerFrostAreaSpells() {
	if !mage.Env.IsForever() {
		return
	}
	mage.ConeOfCold = mage.registerFrostAreaRanks(coneOfColdRanks, SpellCode_MageConeOfCold,
		.129, 10*time.Second, 1+[]float64{0, .12, .23, .35}[mage.Talents.ImprovedConeOfCold], true)
	mage.FrostNova = mage.registerFrostAreaRanks(frostNovaRanks, SpellCode_MageFrostNova,
		.029, time.Duration(25-2*mage.Talents.ImprovedFrostNova)*time.Second, 1, false)
}

func (mage *Mage) registerFrostAreaRanks(ranks []frostAreaRank, code int32, coefficient float64, cooldown time.Duration, multiplier float64, chill bool) []*core.Spell {
	spells := make([]*core.Spell, len(ranks)+1)
	timer := mage.NewTimer()
	for i, rank := range ranks {
		if rank.level > int(mage.Level) {
			continue
		}
		midpoint := rank.base + rank.perLevel*float64(min(int(mage.Level), rank.level+5)-rank.level)
		spread := rank.base * rank.variance / 2
		flags := SpellFlagMage | core.SpellFlagAPL | core.SpellFlagBinary
		if chill {
			flags |= SpellFlagChillSpell
		}
		spells[i+1] = mage.RegisterSpell(core.SpellConfig{
			ActionID:      core.ActionID{SpellID: rank.id},
			SpellCode:     code,
			SpellSchool:   core.SpellSchoolFrost,
			DefenseType:   core.DefenseTypeMagic,
			ProcMask:      core.ProcMaskSpellDamage,
			Flags:         flags,
			RequiredLevel: rank.level,
			Rank:          i + 1,
			ManaCost:      core.ManaCostOptions{FlatCost: rank.mana},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{GCD: core.GCDDefault},
				CD:          core.Cooldown{Timer: timer, Duration: cooldown},
			},
			DamageMultiplier: multiplier,
			ThreatMultiplier: 1,
			BonusCoefficient: coefficient,
			ExtraCastCondition: func(_ *core.Simulation, _ *core.Unit) bool {
				return mage.DistanceFromTarget <= 10
			},
			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
				// Like the other point-blank/cone AOEs, targets are assumed to
				// be clustered. Root/slow control and cone geometry are not
				// modeled; in particular, Nova does not freeze a raid boss.
				for _, target := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, target, sim.Roll(midpoint-spread, midpoint+spread), spell.OutcomeMagicHitAndCrit)
				}
			},
		})
	}
	return spells
}
