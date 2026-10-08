package shaman

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Totem Item IDs
const (
	StormfuryTotem           = 31031
	TotemOfAncestralGuidance = 32330
	TotemOfStorms            = 23199
	TotemOfTheVoid           = 28248
	TotemOfHex               = 40267
	VentureCoLightningRod    = 38361
	ThunderfallTotem         = 45255
)

// The current 70245 client retains separate Overload damage rows for all parent
// ranks. Their levels/caps and Shaman spell families match the parents; the
// current three-rank 408438 trait, not a rune grant, supplies Overload. Source
// row IDs below are provenance; metrics retain the parent ID plus Overload tag.
// The server's implicit row selection is corroborated by upstream PR659, not
// independently decoded. Keep this data substitution Forever-only.
type electricOverloadDamage struct {
	sourceSpellID                                     int32
	basePoints, variance, pointsPerLevel, coefficient float64
	spellLevel, maxLevel                              int
}

func (damage electricOverloadDamage) rangeAtLevel(level int32) (float64, float64) {
	growth := float64(max(0, min(int(level), damage.maxLevel)-damage.spellLevel)) * damage.pointsPerLevel
	spread := damage.basePoints * damage.variance / 2
	mean := damage.basePoints + growth
	return mean - spread, mean + spread
}

// Shared precomputation logic for LB and CL.
func (shaman *Shaman) newElectricSpellConfig(actionID core.ActionID, baseCost float64, baseCastTime time.Duration) core.SpellConfig {
	spell := core.SpellConfig{
		ActionID:     actionID,
		SpellSchool:  core.SpellSchoolNature,
		DefenseType:  core.DefenseTypeMagic,
		ProcMask:     core.ProcMaskSpellDamage,
		Flags:        SpellFlagShaman | SpellFlagLightning | core.SpellFlagAPL,
		MetricSplits: 6,

		ManaCost: core.ManaCostOptions{
			FlatCost:   baseCost,
			Multiplier: 100 - 2*shaman.Talents.Convection,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				CastTime: baseCastTime - shaman.elementalAlacrityReduction(),
				GCD:      core.GCDDefault,
			},
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				castTime := shaman.ApplyCastSpeedForSpell(cast.CastTime, spell)
				// Both spells have client attribute 6, 0x02000000:
				// instant casts preserve the swing timer. Partial reductions
				// remain ordinary hardcasts and reset the full swing.
				if castTime > 0 || !shaman.Env.IsForever() {
					shaman.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime+castTime, false)
				}
			},
		},

		BonusCritRating: core.TernaryFloat64(shaman.Talents.CallOfThunder, 3, 0) * core.SpellCritRatingPerCritChance,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
	}

	return spell
}
