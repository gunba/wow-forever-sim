package warlock

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const BaneOfAgonyRanks = 6

func (warlock *Warlock) getBaneOfAgonyBaseConfig(rank int) core.SpellConfig {
	numTicks := int32(12)
	tickLength := time.Second * 2

	spellId := [BaneOfAgonyRanks + 1]int32{0, 980, 1014, 6217, 11711, 11712, 11713}[rank]
	// Beta client 1.60.1: 0.133 per tick at every rank, and the average tick roughly halved
	spellCoeff := [BaneOfAgonyRanks + 1]float64{0, .133, .133, .133, .133, .133, .133}[rank]
	baseDamage := [BaneOfAgonyRanks + 1]float64{0, 6, 10, 14, 21, 33, 46}[rank] * (1 + .05*float64(warlock.Talents.ImprovedBaneOfAgony))
	manaCost := [BaneOfAgonyRanks + 1]float64{0, 25, 50, 90, 130, 170, 215}[rank]
	level := [BaneOfAgonyRanks + 1]int{0, 8, 18, 28, 38, 48, 58}[rank]

	snapshotBaseDmgNoBonus := 0.0

	return core.SpellConfig{
		SpellCode:     SpellCode_WarlockBaneOfAgony,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolShadow,
		DefenseType:   core.DefenseTypeMagic,
		Flags:         core.SpellFlagAPL | core.SpellFlagResetAttackSwing | core.SpellFlagPureDot | WarlockFlagAffliction,
		ProcMask:      core.ProcMaskSpellDamage,
		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		CritDamageBonus: 0,

		DamageMultiplierAdditive: 1,
		DamageMultiplier:         1,
		ThreatMultiplier:         1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "BaneofAgony-" + warlock.Label + strconv.Itoa(rank),
			},
			NumberOfTicks:    numTicks,
			TickLength:       tickLength,
			BonusCoefficient: spellCoeff,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				baseDmg := baseDamage

				if warlock.AmplifyCurseAura.IsActive() {
					baseDmg *= 1.5
					warlock.AmplifyCurseAura.Deactivate(sim)
				}

				// BoA starts with 50% base damage, but bonus from spell power is not changed.
				// Every 4 ticks this base damage is added again, resulting in 150% base damage for the last 4 ticks
				snapshotBaseDmgNoBonus = baseDmg * 0.5

				dot.Snapshot(target, snapshotBaseDmgNoBonus, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				if dot.TickCount%4 == 0 { // BoA ramp up
					dot.SnapshotBaseDamage += snapshotBaseDmgNoBonus
				}
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
			if result.Landed() {
				dot := spell.Dot(target)

				if activeBane := warlock.ActiveBaneAura.Get(target); activeBane != nil && activeBane != dot.Aura {
					activeBane.Deactivate(sim)
				}

				dot.Apply(sim)
				warlock.ActiveBaneAura[target.UnitIndex] = dot.Aura
			}
			spell.DealOutcome(sim, result)
		},
	}
}

func (warlock *Warlock) registerBaneOfAgonySpell() {
	warlock.BaneOfAgony = make([]*core.Spell, 0)
	for rank := 1; rank <= BaneOfAgonyRanks; rank++ {
		config := warlock.getBaneOfAgonyBaseConfig(rank)

		if config.RequiredLevel <= int(warlock.Level) {
			warlock.BaneOfAgony = append(warlock.BaneOfAgony, warlock.GetOrRegisterSpell(config))
		}
	}
}

func (warlock *Warlock) registerCurseOfRecklessnessSpell() {
	playerLevel := warlock.Level

	warlock.CurseOfRecklessnessAuras = warlock.NewEnemyAuraArray(core.CurseOfRecklessnessAura)

	spellID := map[int32]int32{
		25: 704,
		40: 7658,
		50: 7659,
		60: 11717,
	}[playerLevel]

	rank := map[int32]int{
		25: 1,
		40: 2,
		50: 3,
		60: 4,
	}[playerLevel]

	manaCost := map[int32]float64{
		25: 35.0,
		40: 60.0,
		50: 90.0,
		60: 115.0,
	}[playerLevel]

	warlock.CurseOfRecklessness = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL | WarlockFlagAffliction,
		Rank:        rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		ThreatMultiplier: 1,
		FlatThreatBonus:  156,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
			if result.Landed() {
				aura := warlock.CurseOfRecklessnessAuras.Get(target)
				if activeCurse := warlock.ActiveCurseAura.Get(target); activeCurse != nil && activeCurse != aura {
					activeCurse.Deactivate(sim)
				}

				warlock.ActiveCurseAura[target.UnitIndex] = aura
				warlock.ActiveCurseAura.Get(target).Activate(sim)
			}
		},

		RelatedAuras: []core.AuraArray{warlock.CurseOfRecklessnessAuras},
	})
}

func (warlock *Warlock) registerCurseOfElementsSpell() {
	playerLevel := warlock.Level
	if playerLevel < 40 {
		return
	}

	// Beta client 1.60.1: Curse of Shadow is gone from the spellbook and Curse of the Elements took it
	// over, reducing every magic resistance and raising all magic damage taken (school mask 126 against
	// Classic's Fire and Frost). Its ranks are new ids learned at 30, 40 and 50, the last at Classic's
	// top rank values of 75 resistance and 10%. The raid debuffs in core only split that into Fire and
	// Frost plus Shadow and Arcane, so the curse applies both; Nature and Holy are not covered.
	spellID := map[int32]int32{
		40: 1311677,
		50: 1311680,
		60: 1311680,
	}[playerLevel]

	rank := map[int32]int{
		40: 3,
		50: 4,
		60: 4,
	}[playerLevel]

	manaCost := map[int32]float64{
		40: 150.0,
		50: 200.0,
		60: 200.0,
	}[playerLevel]

	elementsAuras := warlock.NewEnemyAuraArray(core.CurseOfElementsAura)
	shadowAuras := warlock.NewEnemyAuraArray(core.CurseOfShadowAura)
	warlock.CurseOfElementsAuras = warlock.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
		debuffs := []*core.Aura{elementsAuras.Get(unit)}
		if !warlock.Env.IsForever() {
			debuffs = append(debuffs, shadowAuras.Get(unit))
		}
		return unit.RegisterAura(core.Aura{
			Label:    "Curse of the Elements-" + warlock.Label,
			ActionID: core.ActionID{SpellID: spellID},
			Duration: time.Minute * 5,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				for _, debuff := range debuffs {
					debuff.Activate(sim)
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				// A raid debuff made permanent by the encounter settings is not this curse's to remove
				for _, debuff := range debuffs {
					if debuff.Duration != core.NeverExpires {
						debuff.Deactivate(sim)
					}
				}
			},
		})
	})

	warlock.CurseOfElements = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL | WarlockFlagAffliction,
		Rank:        rank,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		ThreatMultiplier: 1,
		FlatThreatBonus:  156,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
			if result.Landed() {
				aura := warlock.CurseOfElementsAuras.Get(target)
				if activeCurse := warlock.ActiveCurseAura.Get(target); activeCurse != nil && activeCurse != aura {
					activeCurse.Deactivate(sim)
				}

				warlock.ActiveCurseAura[target.UnitIndex] = aura
				warlock.ActiveCurseAura.Get(target).Activate(sim)
			}
		},

		RelatedAuras: []core.AuraArray{warlock.CurseOfElementsAuras},
	})
}

func (warlock *Warlock) registerAmplifyCurseSpell() {
	if !warlock.Talents.AmplifyCurse {
		return
	}

	actionID := core.ActionID{SpellID: 18288}

	warlock.AmplifyCurseAura = warlock.GetOrRegisterAura(core.Aura{
		Label:    "Amplify Curse",
		ActionID: actionID,
		Duration: time.Second * 30,
	})

	warlock.AmplifyCurse = warlock.GetOrRegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		Flags:       core.SpellFlagAPL | WarlockFlagAffliction,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    warlock.NewTimer(),
				Duration: 3 * time.Minute,
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			warlock.AmplifyCurseAura.Activate(sim)
		},
	})
}

func (warlock *Warlock) registerBaneOfDoomSpell() {
	if warlock.Level < 60 {
		return
	}

	warlock.BaneOfDoom = warlock.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_WarlockBaneOfDoom,
		ActionID:    core.ActionID{SpellID: 603},
		SpellSchool: core.SpellSchoolShadow,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL | WarlockFlagAffliction,

		RequiredLevel: 60,

		ManaCost: core.ManaCostOptions{
			FlatCost: 300,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    warlock.NewTimer(),
				Duration: time.Second * 60,
			},
		},

		CritDamageBonus: 0,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		FlatThreatBonus:  160,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "BaneofDoom",
			},
			NumberOfTicks: 1,
			TickLength:    time.Minute,
			// Beta client 1.60.1: 1742 damage with a 4.0 spell power coefficient. Classic's 3200 carried
			// no coefficient of its own, and the spell level 1 this file used to set never reached the dot.
			BonusCoefficient: 4,
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				dot.Snapshot(target, 1742, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
			if result.Landed() {
				dot := spell.Dot(target)
				if activeBane := warlock.ActiveBaneAura.Get(target); activeBane != nil && activeBane != dot.Aura {
					activeBane.Deactivate(sim)
				}

				dot.Apply(sim)
				warlock.ActiveBaneAura[target.UnitIndex] = dot.Aura
			}
		},
	})
}

func (warlock *Warlock) registerBaneOfHavocSpell() {
	if !warlock.Talents.BaneOfHavoc {
		return
	}

	actionID := core.ActionID{SpellID: 80240}

	warlock.BaneOfHavocAuras = warlock.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
		return unit.RegisterAura(core.Aura{
			Label:    "Bane of Havoc-" + warlock.Label,
			ActionID: actionID,
			Duration: time.Minute * 5,
		})
	})

	// Only marks the target for now, the 15% damage copy needs a second target to matter
	warlock.BaneOfHavoc = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL | WarlockFlagDestruction,

		// 5% of base mana in the beta client (1225228), not a flat 300
		ManaCost: core.ManaCostOptions{
			BaseCost: 0.05,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
			if result.Landed() {
				aura := warlock.BaneOfHavocAuras.Get(target)
				if activeBane := warlock.ActiveBaneAura.Get(target); activeBane != nil && activeBane != aura {
					activeBane.Deactivate(sim)
				}

				warlock.ActiveBaneAura[target.UnitIndex] = aura
				aura.Activate(sim)
			}
		},

		RelatedAuras: []core.AuraArray{warlock.BaneOfHavocAuras},
	})
}
