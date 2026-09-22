package shaman

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const SearingTotemRanks = 6

// Forever beta client: the attack's damage is unchanged but its coefficient falls to 0.017 at every rank.
var SearingTotemSpellId = [SearingTotemRanks + 1]int32{0, 3599, 6363, 6364, 6365, 10437, 10438}
var SearingTotemAttackSpellId = [SearingTotemRanks + 1]int32{0, 3606, 6350, 6351, 6352, 10435, 10436}
var SearingTotemBaseDamage = [SearingTotemRanks + 1][]float64{{0}, {9, 11}, {13, 17}, {19, 25}, {26, 34}, {33, 45}, {40, 54}}
var SearingTotemSpellCoef = [SearingTotemRanks + 1]float64{0, .017, .017, .017, .017, .017, .017}
var SearingTotemManaCost = [SearingTotemRanks + 1]float64{0, 25, 45, 75, 110, 145, 170}
var SearingTotemDuration = [SearingTotemRanks + 1]int{0, 30, 35, 40, 45, 50, 55}
var SearingTotemLevel = [SearingTotemRanks + 1]int{0, 10, 20, 30, 40, 50, 60}

func (shaman *Shaman) registerSearingTotemSpell() {
	shaman.SearingTotem = make([]*core.Spell, SearingTotemRanks+1)

	for rank := 1; rank <= SearingTotemRanks; rank++ {
		config := shaman.newSearingTotemSpellConfig(rank)

		if config.RequiredLevel <= int(shaman.Level) {
			shaman.SearingTotem[rank] = shaman.RegisterSpell(config)
		}
	}

	shaman.FireTotems = append(
		shaman.FireTotems,
		core.FilterSlice(shaman.SearingTotem, func(spell *core.Spell) bool { return spell != nil })...,
	)
}

func (shaman *Shaman) newSearingTotemSpellConfig(rank int) core.SpellConfig {
	totemSpellId := SearingTotemSpellId[rank]
	baseDamageLow := SearingTotemBaseDamage[rank][0]
	baseDamageHigh := SearingTotemBaseDamage[rank][1]
	spellCoeff := SearingTotemSpellCoef[rank]
	manaCost := SearingTotemManaCost[rank]
	duration := time.Second * time.Duration(SearingTotemDuration[rank])
	level := SearingTotemLevel[rank]

	attackInterval := time.Millisecond * 2500

	attackSpell := shaman.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_ShamanSearingTotem,
		ActionID:    core.ActionID{SpellID: SearingTotemAttackSpellId[rank]},
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       SpellFlagTotem,

		DamageMultiplier: shaman.callOfFlameMultiplier(),
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageLow, baseDamageHigh)
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
		},
	})

	spell := core.SpellConfig{
		SpellCode:   SpellCode_ShamanSearingTotem,
		ActionID:    core.ActionID{SpellID: totemSpellId},
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       SpellFlagTotem | core.SpellFlagAPL,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost:   manaCost,
			Multiplier: shaman.totemManaMultiplier(),
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: totemGCD,
			},
			IgnoreHaste: true,
		},

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: fmt.Sprintf("Searing Totem (Rank %d)", rank),
			},
			// These are the real tick values, but searing totem doesn't start its next
			// cast until the previous missile hits the target. We don't have an option
			// for target distance yet so just pretend the tick rate is lower.
			// https://wotlk.wowhead.com/spell=25530/attack
			//TickLength:           time.Second * 2.2,
			NumberOfTicks: int32(duration / attackInterval),
			TickLength:    attackInterval,
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				attackSpell.Cast(sim, target)
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			if shaman.ActiveTotems[FireTotem] != nil {
				shaman.ActiveTotems[FireTotem].Dot(sim.GetTargetUnit(0)).Cancel(sim)
			}
			spell.Dot(sim.GetTargetUnit(0)).Apply(sim)
			// +1 needed because of rounding issues with totem tick time.
			shaman.TotemExpirations[FireTotem] = sim.CurrentTime + duration + 1
			shaman.ActiveTotems[FireTotem] = spell
		},
	}

	return spell
}

const MagmaTotemRanks = 4

// Forever beta client values for the pulse.
var MagmaTotemSpellId = [MagmaTotemRanks + 1]int32{0, 8190, 10585, 10586, 10587}
var MagmaTotemAoeSpellId = [MagmaTotemRanks + 1]int32{0, 8187, 10579, 10580, 10581}
var MagmaTotemBaseDamage = [MagmaTotemRanks + 1]float64{0, 20, 35, 52, 73}
var MagmaTotemSpellCoeff = [MagmaTotemRanks + 1]float64{0, .033, .033, .033, .033}
var MagmaTotemManaCost = [MagmaTotemRanks + 1]float64{0, 230, 360, 500, 650}
var MagmaTotemLevel = [MagmaTotemRanks + 1]int{0, 26, 36, 46, 56}

func (shaman *Shaman) registerMagmaTotemSpell() {
	shaman.MagmaTotem = make([]*core.Spell, MagmaTotemRanks+1)

	for rank := 1; rank <= MagmaTotemRanks; rank++ {
		config := shaman.newMagmaTotemSpellConfig(rank)

		if config.RequiredLevel <= int(shaman.Level) {
			shaman.MagmaTotem[rank] = shaman.RegisterSpell(config)
		}
	}

	shaman.FireTotems = append(
		shaman.FireTotems,
		core.FilterSlice(shaman.MagmaTotem, func(spell *core.Spell) bool { return spell != nil })...,
	)
}

func (shaman *Shaman) newMagmaTotemSpellConfig(rank int) core.SpellConfig {
	spellId := MagmaTotemSpellId[rank]
	baseDamage := MagmaTotemBaseDamage[rank]
	spellCoeff := MagmaTotemSpellCoeff[rank]
	manaCost := MagmaTotemManaCost[rank]
	level := MagmaTotemLevel[rank]

	duration := time.Second * 20
	attackInterval := time.Second * 2

	aoeSpell := shaman.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_ShamanMagmaTotem,
		ActionID:    core.ActionID{SpellID: MagmaTotemAoeSpellId[rank]},
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       SpellFlagTotem,

		DamageMultiplier: shaman.callOfFlameMultiplier(),
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for _, aoeTarget := range sim.Encounter.TargetUnits {
				spell.CalcAndDealDamage(sim, aoeTarget, baseDamage, spell.OutcomeMagicHitAndCrit)
			}
		},
	})

	spell := core.SpellConfig{
		SpellCode:   SpellCode_ShamanMagmaTotem,
		ActionID:    core.ActionID{SpellID: spellId},
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       SpellFlagTotem | core.SpellFlagAPL,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost:   manaCost,
			Multiplier: shaman.totemManaMultiplier(),
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: totemGCD,
			},
			IgnoreHaste: true,
		},

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: fmt.Sprintf("Magma Totem (Rank %d)", rank),
			},
			NumberOfTicks: int32(duration / attackInterval),
			TickLength:    attackInterval,

			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				aoeSpell.Cast(sim, target)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			if shaman.ActiveTotems[FireTotem] != nil {
				shaman.ActiveTotems[FireTotem].Dot(sim.GetTargetUnit(0)).Cancel(sim)
			}
			spell.Dot(sim.GetTargetUnit(0)).Apply(sim)
			// +1 needed because of rounding issues with totem tick time.
			shaman.TotemExpirations[FireTotem] = sim.CurrentTime + duration + 1
			shaman.ActiveTotems[FireTotem] = spell
		},
	}

	return spell
}

// Forever has no Fire Nova Totem: its spell ids are gone from the beta client, and the Fire Nova the spellbook
// teaches in its place is the caster-centred nova, a 10 sec cooldown that bursts on every nearby enemy at once.
// Mana costs and levels are the totem's; damage and the 0.214 coefficient are the beta client's, scaled to level
// 60 like Lightning Bolt's.
const FireNovaRanks = 5

var FireNovaSpellId = [FireNovaRanks + 1]int32{0, 408341, 408342, 408343, 408344, 408345}
var FireNovaBaseDamage = [FireNovaRanks + 1][]float64{{0, 0}, {51, 60}, {103, 117}, {182, 206}, {280, 316}, {397, 443}}
var FireNovaSpellCoeff = [FireNovaRanks + 1]float64{0, .214, .214, .214, .214, .214}
var FireNovaManaCost = [FireNovaRanks + 1]float64{0, 95, 170, 280, 395, 520}
var FireNovaLevel = [FireNovaRanks + 1]int{0, 12, 22, 32, 42, 52}

func (shaman *Shaman) registerFireNovaSpell() {
	shaman.FireNova = make([]*core.Spell, FireNovaRanks+1)
	cdTimer := shaman.NewTimer()

	for rank := 1; rank <= FireNovaRanks; rank++ {
		config := shaman.newFireNovaSpellConfig(rank, cdTimer)

		if config.RequiredLevel <= int(shaman.Level) {
			shaman.FireNova[rank] = shaman.RegisterSpell(config)
		}
	}
}

func (shaman *Shaman) newFireNovaSpellConfig(rank int, cdTimer *core.Timer) core.SpellConfig {
	spellId := FireNovaSpellId[rank]
	baseDamageLow := FireNovaBaseDamage[rank][0]
	baseDamageHigh := FireNovaBaseDamage[rank][1]
	spellCoeff := FireNovaSpellCoeff[rank]
	cooldown := time.Second*10 - shaman.improvedFireNovaCooldownReduction()
	manaCost := FireNovaManaCost[rank]
	level := FireNovaLevel[rank]

	return core.SpellConfig{
		SpellCode:   SpellCode_ShamanFireNova,
		ActionID:    core.ActionID{SpellID: spellId},
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagShaman | core.SpellFlagAPL,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},

		// Fire Nova originates from this Shaman's active Fire totem, not the
		// character. The totem pointer remains after expiration, so check time.
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return shaman.ActiveTotems[FireTotem] != nil && shaman.TotemExpirations[FireTotem] > sim.CurrentTime
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: cooldown,
			},
		},

		DamageMultiplier: shaman.callOfFlameMultiplier() * shaman.improvedFireNovaMultiplier(),
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageLow, baseDamageHigh)
			for _, aoeTarget := range sim.Encounter.TargetUnits {
				spell.CalcAndDealDamage(sim, aoeTarget, baseDamage, spell.OutcomeMagicHitAndCrit)
			}
		},
	}
}
