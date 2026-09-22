package shaman

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// The talent teaches rank 1 and the spellbook adds ranks 2 and 3 at 50 and 60. Everything here is the beta
// client's: a 2.5 sec cast, a 10 sec cooldown, flat mana costs, the full 0.714 coefficient and 20% more damage
// on a target carrying the shaman's Flame Shock. Damage is scaled to level 60 like Lightning Bolt's.
const LavaBurstRanks = 3
const LavaBurstFlameShockBonus = .2

var LavaBurstSpellId = [LavaBurstRanks + 1]int32{0, 408490, 1238299, 1238300}
var LavaBurstBaseDamage = [LavaBurstRanks + 1][]float64{{0}, {105, 135}, {165, 211}, {192, 248}}
var LavaBurstSpellCoef = [LavaBurstRanks + 1]float64{0, .714, .714, .714}
var LavaBurstManaCost = [LavaBurstRanks + 1]float64{0, 165, 230, 265}
var LavaBurstLevel = [LavaBurstRanks + 1]int{0, 40, 50, 60}

func (shaman *Shaman) registerLavaBurstSpell() {
	shaman.LavaBurst = make([]*core.Spell, LavaBurstRanks+1)

	if !shaman.Talents.LavaBurst {
		return
	}

	cdTimer := shaman.NewTimer()

	for rank := 1; rank <= LavaBurstRanks; rank++ {
		if LavaBurstLevel[rank] <= int(shaman.Level) {
			shaman.LavaBurst[rank] = shaman.RegisterSpell(shaman.newLavaBurstSpellConfig(rank, cdTimer))
		}
	}
}

func (shaman *Shaman) newLavaBurstSpellConfig(rank int, cdTimer *core.Timer) core.SpellConfig {
	spellId := LavaBurstSpellId[rank]
	baseDamageLow := LavaBurstBaseDamage[rank][0]
	baseDamageHigh := LavaBurstBaseDamage[rank][1]
	spellCoeff := LavaBurstSpellCoef[rank]
	manaCost := LavaBurstManaCost[rank]
	level := LavaBurstLevel[rank]

	castTime := time.Millisecond * 2500

	return core.SpellConfig{
		SpellCode:   SpellCode_ShamanLavaBurst,
		ActionID:    core.ActionID{SpellID: spellId},
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagShaman | core.SpellFlagAPL,

		RequiredLevel: level,
		Rank:          rank,

		MissileSpeed: 20,

		ManaCost: core.ManaCostOptions{
			FlatCost:   manaCost,
			Multiplier: 100 - 2*shaman.Talents.Convection,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				CastTime: castTime - shaman.elementalAlacrityReduction(),
				GCD:      core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: time.Second * 10,
			},
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				castTime := shaman.ApplyCastSpeedForSpell(cast.CastTime, spell)
				shaman.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime+castTime, false)
			},
		},

		DamageMultiplier: shaman.callOfFlameMultiplier(),
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			if shaman.hasActiveFlameShock(target) {
				spell.DamageMultiplier *= 1 + LavaBurstFlameShockBonus
				defer func() { spell.DamageMultiplier /= 1 + LavaBurstFlameShockBonus }()
			}

			baseDamage := sim.Roll(baseDamageLow, baseDamageHigh)
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
			})
		},
	}
}

func (shaman *Shaman) hasActiveFlameShock(target *core.Unit) bool {
	for _, spell := range shaman.FlameShock {
		if spell != nil && spell.Dot(target).IsActive() {
			return true
		}
	}
	return false
}
