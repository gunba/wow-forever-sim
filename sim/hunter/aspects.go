package hunter

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

// Utility function to create the Deadly Aspects haste aura
func (hunter *Hunter) createDeadlyAspectsAura(auraLabel string, actionID core.ActionID, melee bool) *core.Aura {
	// Every rank of Deadly Aspects triggers the same Quick Shots (6150): 30% for 12 sec. The
	// points buy only the proc chance.
	bonusMultiplier := 1.3
	return hunter.GetOrRegisterAura(core.Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: time.Second * 12,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			if melee {
				aura.Unit.MultiplyMeleeSpeed(sim, bonusMultiplier)
			} else {
				aura.Unit.MultiplyRangedSpeed(sim, bonusMultiplier)
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			if melee {
				aura.Unit.MultiplyMeleeSpeed(sim, 1/bonusMultiplier)
			} else {
				aura.Unit.MultiplyRangedSpeed(sim, 1/bonusMultiplier)
			}
		},
	})
}

// Function to get the maximum attack power for Aspect of the Hawk based on rank
func (hunter *Hunter) getMaxAspectOfTheHawkAttackPower(rank int) float64 {
	// Rank 6 is 55 in the beta client, down from 110, while ranks 5 and 7 are unchanged.
	attackPower := [8]float64{0, 20, 35, 50, 70, 90, 55, 120}

	if rank < 1 || rank > 7 {
		return 0.0
	}

	return attackPower[rank]
}

func (hunter *Hunter) getMaxHawkRank() int {
	maxRank := core.TernaryInt(core.IncludeAQ, 7, 6)

	for i := maxRank; i > 0; i-- {
		config := hunter.getAspectOfTheHawkSpellConfig(i)
		if config.RequiredLevel <= int(hunter.Level) {
			return i
		}
	}
	return 1
}

func (hunter *Hunter) getAspectOfTheHawkSpellConfig(rank int) core.SpellConfig {
	var deadlyAspectsAura *core.Aura
	// Deadly Aspects: 2/4/6/8/10% (client curve).
	deadlyAspectsProcChance := 0.02 * float64(hunter.Talents.DeadlyAspects)

	spellIds := [8]int32{0, 13165, 14318, 14319, 14320, 14321, 14322, 25296}
	levels := [8]int{0, 10, 18, 28, 38, 48, 58, 60}

	spellId := spellIds[rank]
	level := levels[rank]

	if hunter.Talents.DeadlyAspects > 0 {
		deadlyAspectsAura = hunter.createDeadlyAspectsAura(
			"Quick Shots",
			core.ActionID{SpellID: 6150},
			false,
		)
	}
	// Use utility function to get the attack power based on rank
	rap := hunter.getMaxAspectOfTheHawkAttackPower(rank)

	actionID := core.ActionID{SpellID: spellId}
	aspectOfTheHawkAura := hunter.GetOrRegisterAura(core.Aura{
		Label:    "Aspect of the Hawk" + strconv.Itoa(rank),
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.RangedAttackPower, rap*hunter.AspectOfTheHawkAPMultiplier)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.RangedAttackPower, -rap*hunter.AspectOfTheHawkAPMultiplier)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !spell.ProcMask.Matches(core.ProcMaskRangedAuto) {
				return
			}

			if deadlyAspectsAura != nil && sim.Proc(deadlyAspectsProcChance, "Deadly Aspects") {
				deadlyAspectsAura.Activate(sim)
			}
		},
	})

	aspectOfTheHawkAura.NewExclusiveEffect("Aspect", true, core.ExclusiveEffect{})

	return core.SpellConfig{
		ActionID:      actionID,
		Flags:         core.SpellFlagAPL,
		Rank:          rank,
		RequiredLevel: level,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return !aspectOfTheHawkAura.IsActive()
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			aspectOfTheHawkAura.Activate(sim)
		},
	}
}

func (hunter *Hunter) registerAspectOfTheHawkSpell() {
	hunter.AspectOfTheHawkAPMultiplier = 1.0
	maxRank := hunter.getMaxHawkRank()
	config := hunter.getAspectOfTheHawkSpellConfig(maxRank)
	hunter.GetOrRegisterSpell(config)
}

func (hunter *Hunter) registerAspectOfTheBeastSpell() {
	if !hunter.Env.IsForever() || hunter.Level < 60 {
		return
	}

	// Level-60 rank: beta client 1.60.1.69893 and Forever spell 1299447.
	// Quick Strikes (1299448) grants 30% melee haste for 12 seconds.
	actionID := core.ActionID{SpellID: 1299447}
	var quickStrikes *core.Aura
	if hunter.Talents.DeadlyAspects > 0 {
		quickStrikes = hunter.createDeadlyAspectsAura("Quick Strikes", core.ActionID{SpellID: 1299448}, true)
	}
	procChance := 0.02 * float64(hunter.Talents.DeadlyAspects)
	aspect := hunter.RegisterAura(core.Aura{
		Label:    "Aspect of the Beast4",
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			hunter.AddStatDynamic(sim, stats.AttackPower, 110)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			hunter.AddStatDynamic(sim, stats.AttackPower, -110)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if quickStrikes != nil && spell.ProcMask.Matches(core.ProcMaskMeleeWhiteHit) && sim.Proc(procChance, "Deadly Aspects") {
				quickStrikes.Activate(sim)
			}
		},
	})
	aspect.NewExclusiveEffect("Aspect", true, core.ExclusiveEffect{})
	hunter.RegisterSpell(core.SpellConfig{
		ActionID:      actionID,
		SpellSchool:   core.SpellSchoolNature,
		Flags:         core.SpellFlagAPL,
		Rank:          4,
		RequiredLevel: 60,
		ManaCost:      core.ManaCostOptions{FlatCost: 110},
		Cast:          core.CastConfig{DefaultCast: core.Cast{GCD: core.GCDDefault}, IgnoreHaste: true},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return !aspect.IsActive()
		},
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			aspect.Activate(sim)
		},
	})
}
