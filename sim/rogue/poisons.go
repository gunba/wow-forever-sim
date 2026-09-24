package rogue

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

/**
Forever beta client 1.60.1.69893. Proc chances are Classic's; charges are about 1.5x Classic's and
still not modelled.

Instant Poison: 20% proc chance
25: 13-17 damage, 8679 ID, 60 charges
40: 29-37 damage, 8688 ID, 105 charges
50: 45-57 damage, 11338 ID, 130 charges
60: 76-100 damage, 11340 ID, 175 charges

Deadly Poison: 30% proc chance, 5 stacks, damage per stack over 12 sec
40: 36 damage, 2824 ID, 115 charges
50: 56 damage, 11355 ID, 135 charges
60: 92 damage, 25347 ID, 180 charges (Rank 5, by book)

Wound Poison: 30% proc chance, 5 stacks
25: x damage, x ID (none, first rank is level 32)
40: -75 healing, 11325 ID, 75 charges (Rank 2)
50: -105 healing, 13226 ID, 90 charges (Rank 3)
60: -135 healing, 13227 ID, 105 charges (Rank 4)
*/

// TODO: Add charges to poisons

type PoisonProcSource int

const (
	NormalProc PoisonProcSource = iota
)

func (rogue *Rogue) GetInstantPoisonProcChance() float64 {
	return 0.2 + rogue.improvedPoisons() + rogue.additivePoisonBonusChance
}

func (rogue *Rogue) GetDeadlyPoisonProcChance() float64 {
	return 0.3 + rogue.improvedPoisons() + rogue.additivePoisonBonusChance
}

func (rogue *Rogue) GetWoundPoisonProcChance() float64 {
	return 0.3 + rogue.improvedPoisons() + rogue.additivePoisonBonusChance
}

func (rogue *Rogue) improvedPoisons() float64 {
	return []float64{0, 0.02, 0.04, 0.06, 0.08, 0.1}[rogue.Talents.ImprovedPoisons]
}

func (rogue *Rogue) getPoisonDamageMultiplier() float64 {
	return []float64{1, 1.04, 1.08, 1.12, 1.16, 1.2}[rogue.Talents.VilePoisons]
}

// Venom lands after the poisons are registered, so the spells are scaled directly.
// Forever Deadly Poison evaluates this multiplier on each tick. Classic retains
// its application-time snapshot and needs the existing dot adjusted as well.
//
// TODO: The tooltip doesn't say whether Venom reaches a Deadly Poison that is already on
// the target or only the stacks applied while it is up. Beta will confirm.
func (rogue *Rogue) multiplyPoisonDamage(multiplier float64) {
	for _, spell := range rogue.Spellbook {
		if spell.Flags.Matches(SpellFlagRoguePoison) {
			spell.DamageMultiplier *= multiplier
		}
	}

	if !rogue.Env.IsForever() {
		for _, target := range rogue.Env.Encounter.TargetUnits {
			if dot := rogue.deadlyPoisonTick.Dot(target); dot.IsActive() {
				dot.SnapshotAttackerMultiplier *= multiplier
			}
		}
	}
}

///////////////////////////////////////////////////////////////////////////
//                               Apply Poisons
///////////////////////////////////////////////////////////////////////////

func (rogue *Rogue) applyPoisons() {
	rogue.applyDeadlyPoison()
	rogue.applyInstantPoison()
	rogue.applyWoundPoison()
}

// Apply Instant Poison to weapon and enable procs
func (rogue *Rogue) applyInstantPoison() {
	procMask := rogue.getImbueProcMask(proto.WeaponImbue_InstantPoison)
	if procMask == core.ProcMaskUnknown {
		return
	}

	rogue.RegisterAura(core.Aura{
		Label:    "Instant Poison",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() || !spell.ProcMask.Matches(procMask) {
				return
			}

			if sim.RandomFloat("Instant Poison") < rogue.GetInstantPoisonProcChance() {
				rogue.InstantPoison.Cast(sim, result.Target)
			}
		},
	})
}

// Apply Deadly Poison to weapon and enable procs
func (rogue *Rogue) applyDeadlyPoison() {
	procMask := rogue.getImbueProcMask(proto.WeaponImbue_DeadlyPoison)
	if procMask == core.ProcMaskUnknown {
		return
	}

	rogue.RegisterAura(core.Aura{
		Label:    "Deadly Poison",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() || !spell.ProcMask.Matches(procMask) {
				return
			}
			if sim.RandomFloat("Deadly Poison") < rogue.GetDeadlyPoisonProcChance() {
				rogue.DeadlyPoison.Cast(sim, result.Target)
			}
		},
	})
}

// Apply Wound Poison to weapon and enable procs
func (rogue *Rogue) applyWoundPoison() {
	procMask := rogue.getImbueProcMask(proto.WeaponImbue_WoundPoison)
	if procMask == core.ProcMaskUnknown {
		return
	}

	rogue.RegisterAura(core.Aura{
		Label:    "Wound Poison",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !result.Landed() || !spell.ProcMask.Matches(procMask) {
				return
			}

			if sim.RandomFloat("Wound Poison") < rogue.GetWoundPoisonProcChance() {
				rogue.WoundPoison.Cast(sim, result.Target)
			}
		},
	})
}

///////////////////////////////////////////////////////////////////////////
//                              Register Poisons
///////////////////////////////////////////////////////////////////////////

func (rogue *Rogue) registerInstantPoisonSpell() {
	rogue.InstantPoison = rogue.makeInstantPoison()
}

func (rogue *Rogue) registerDeadlyPoisonSpell() {
	// Beta client 1.60.1.69893: about a third less a tick at every rank (rank 5 34 -> 23).
	baseDamageTick := map[int32]float64{
		25: 6,
		40: 9,
		50: 14,
		60: core.TernaryFloat64(core.IncludeAQ, 23, 18),
	}[rogue.Level]
	spellID := map[int32]int32{
		25: 2823,
		40: 2824,
		50: 11355,
		60: core.TernaryInt32(core.IncludeAQ, 25347, 11356),
	}[rogue.Level]

	rogue.deadlyPoisonTick = rogue.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID, Tag: 100},
		SpellSchool: core.SpellSchoolNature,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamageProc,
		Flags:       core.SpellFlagPoison | core.SpellFlagPassiveSpell | SpellFlagRoguePoison,

		DamageMultiplier: rogue.getPoisonDamageMultiplier(),
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label:     "DeadlyPoison",
				MaxStacks: 5,
				Duration:  time.Second * 12,
			},
			NumberOfTicks: 4,
			TickLength:    time.Second * 3,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, applyStack bool) {
				if !applyStack {
					return
				}

				// only the first stack snapshots the multiplier
				if dot.GetStacks() == 1 {
					attackTable := dot.Spell.Unit.AttackTables[target.UnitIndex][dot.Spell.CastType]
					dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(attackTable, true)
					dot.SnapshotBaseDamage = 0
				}

				dot.SnapshotBaseDamage += baseDamageTick
			},

			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				if sim.IsForever() {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, dot.SnapshotBaseDamage, dot.OutcomeTick)
				} else {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				}
			},
		},
	})

	rogue.DeadlyPoison = rogue.makeDeadlyPoison()
}

func (rogue *Rogue) registerWoundPoisonSpell() {
	woundPoisonDebuffAura := core.Aura{
		Label:     "WoundPoison-" + strconv.Itoa(int(rogue.Index)),
		ActionID:  core.ActionID{SpellID: 13219},
		MaxStacks: 5,
		Duration:  time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			// all healing effects used on target reduced by x, stacks 5 times
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			// undo reduced healing effects used on targets
		},
	}

	rogue.woundPoisonDebuffAuras = rogue.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return target.RegisterAura(woundPoisonDebuffAura)
	})
	rogue.WoundPoison = rogue.makeWoundPoison()
}

///////////////////////////////////////////////////////////////////////////
//                              Make Poisons
///////////////////////////////////////////////////////////////////////////

// Make a source based variant of Instant Poison
func (rogue *Rogue) makeInstantPoison() *core.Spell {
	// Beta client 1.60.1.69893: about a third less at every rank (rank 6 112-148 -> 76-100).
	baseDamageByLevel := map[int32]float64{
		25: 13,
		40: 29,
		50: 45,
		60: 76,
	}[rogue.Level]

	damageVariance := map[int32]float64{
		25: 4,
		40: 8,
		50: 12,
		60: 24,
	}[rogue.Level]

	spellID := map[int32]int32{
		25: 8679,
		40: 8688,
		50: 11338,
		60: 11340,
	}[rogue.Level]

	return rogue.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolNature,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamageProc,
		Flags:       core.SpellFlagPoison | core.SpellFlagPassiveSpell | SpellFlagRoguePoison,

		DamageMultiplier: rogue.getPoisonDamageMultiplier(),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageByLevel, baseDamageByLevel+damageVariance)
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
		},
	})
}

func (rogue *Rogue) makeDeadlyPoison() *core.Spell {
	return rogue.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{SpellID: rogue.deadlyPoisonTick.SpellID},
		Flags:    core.SpellFlagPoison | core.SpellFlagPassiveSpell | SpellFlagRoguePoison,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)

			if !result.Landed() {
				return
			}

			dot := rogue.deadlyPoisonTick.Dot(target)

			dot.ApplyOrRefresh(sim)
			if dot.GetStacks() < dot.MaxStacks {
				dot.AddStack(sim)
				// snapshotting only takes place when adding a stack
				dot.TakeSnapshot(sim, true)
			}
		},
	})
}

// Make a source based variant of Wound Poison
func (rogue *Rogue) makeWoundPoison() *core.Spell {
	return rogue.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 13219},
		SpellSchool: core.SpellSchoolNature,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamageProc,
		Flags:       core.SpellFlagPoison | core.SpellFlagPassiveSpell | SpellFlagRoguePoison,

		DamageMultiplier: rogue.getPoisonDamageMultiplier(),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)

			if !result.Landed() {
				return
			}

			aura := rogue.woundPoisonDebuffAuras.Get(target)
			if !aura.IsActive() {
				aura.Activate(sim)
				aura.SetStacks(sim, 1)
				return
			}

			if aura.GetStacks() < 5 {
				aura.Refresh(sim)
				aura.AddStack(sim)
				return
			}
			aura.Refresh(sim)
		},
	})
}
