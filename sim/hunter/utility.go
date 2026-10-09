package hunter

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// Utility control is opt-in per target/mechanic. Raid targets remain immune;
// no target level or creature-type heuristic is used to manufacture eligibility.
func (hunter *Hunter) registerSlowAura(target *core.Unit, id int32, duration time.Duration, penalty float64) *core.Aura {
	return target.GetOrRegisterAura(core.Aura{
		Label:    fmt.Sprintf("Hunter slow %d-%d", id, hunter.UnitIndex),
		ActionID: core.ActionID{SpellID: id, Tag: hunter.Index + 1},
		Duration: duration,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			target.AddMoveSpeedModifierDynamic(sim, &aura.ActionID, penalty)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			target.RemoveMoveSpeedModifierDynamic(sim, &aura.ActionID)
		},
	})
}

func (hunter *Hunter) slowTarget(sim *core.Simulation, target *core.Unit, id int32, duration time.Duration, penalty float64) {
	if target.PseudoStats.CanBeSlowed {
		hunter.registerSlowAura(target, id, duration, penalty).Activate(sim)
	}
}

func hunterRootDRGroup(id int32) int32 {
	// 70291 SpellCategories: Entrapment19185 mask1; Wing Clip19229 mask0.
	if id == 19185 {
		return core.CrowdControlDRRoot
	}
	return 0
}

func (hunter *Hunter) registerRootAura(target *core.Unit, id int32, duration time.Duration) *core.Aura {
	return target.RegisterCrowdControlAura(core.Aura{
		Label:    fmt.Sprintf("Hunter root %d-%d", id, hunter.UnitIndex),
		ActionID: core.ActionID{SpellID: id, Tag: hunter.Index + 1},
		Duration: duration,
		OnGain:   func(_ *core.Aura, sim *core.Simulation) { target.AddRoot(sim) },
		OnExpire: func(_ *core.Aura, sim *core.Simulation) { target.RemoveRoot(sim) },
	}, hunterRootDRGroup(id))
}

func (hunter *Hunter) rootTarget(sim *core.Simulation, target *core.Unit, id int32, duration time.Duration) {
	if target.PseudoStats.CanBeRooted {
		target.ApplyCrowdControl(sim, hunter.registerRootAura(target, id, duration), duration, hunterRootDRGroup(id))
	}
}

func (hunter *Hunter) registerStunAura(target *core.Unit, id int32, duration time.Duration) *core.Aura {
	return target.RegisterCrowdControlAura(core.Aura{
		Label:    fmt.Sprintf("Hunter stun %d-%d", id, hunter.UnitIndex),
		ActionID: core.ActionID{SpellID: id, Tag: hunter.Index + 1},
		Duration: duration,
		OnGain:   func(_ *core.Aura, sim *core.Simulation) { target.AddStun(sim) },
		OnExpire: func(_ *core.Aura, sim *core.Simulation) { target.RemoveStun(sim) },
	}, core.CrowdControlDRStunProc)
}

func (hunter *Hunter) stunTarget(sim *core.Simulation, target *core.Unit, id int32, duration time.Duration) {
	if target.PseudoStats.CanBeStunned {
		target.ApplyCrowdControl(sim, hunter.registerStunAura(target, id, duration), duration, core.CrowdControlDRStunProc)
	}
}

func (hunter *Hunter) registerUtilityAuras() {
	if !hunter.Env.IsForever() {
		return
	}
	for _, target := range hunter.Env.Encounter.TargetUnits {
		hunter.registerSlowAura(target, 5116, 4*time.Second, .5)
		hunter.registerSlowAura(target, 13810, time.Duration(float64(30*time.Second)*(1+.15*float64(hunter.Talents.CleverTraps))), .6)
		for rank, id := range []int32{2974, 14267, 14268} {
			hunter.registerSlowAura(target, id, 10*time.Second, []float64{.5, .55, .6}[rank])
		}
		hunter.registerRootAura(target, 19229, 5*time.Second)
		hunter.registerRootAura(target, 19185, time.Duration(hunter.Talents.Entrapment)*time.Second)
		hunter.registerStunAura(target, 19410, 3*time.Second)
	}
}

func (hunter *Hunter) registerConcussiveShotSpell() {
	if !hunter.Env.IsForever() || hunter.Level < 8 {
		return
	}
	hunter.ConcussiveShot = hunter.RegisterSpell(core.SpellConfig{
		ActionID:      core.ActionID{SpellID: 5116},
		SpellCode:     SpellCode_HunterConcussiveShot,
		SpellSchool:   core.SpellSchoolArcane,
		DefenseType:   core.DefenseTypeRanged,
		ProcMask:      core.ProcMaskRangedSpecial,
		Flags:         core.SpellFlagAPL | SpellFlagShot,
		CastType:      proto.CastType_CastTypeRanged,
		RequiredLevel: 8,
		MissileSpeed:  40,
		ManaCost:      core.ManaCostOptions{BaseCost: .08},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			IgnoreHaste: true,
			CD:          core.Cooldown{Timer: hunter.NewTimer(), Duration: 12 * time.Second},
		},
		ExtraCastCondition: func(_ *core.Simulation, _ *core.Unit) bool {
			return hunter.HasRangedWeapon() && hunter.DistanceFromTarget >= core.MinRangedAttackDistance
		},
		ThreatMultiplier: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcOutcome(sim, target, spell.OutcomeRangedHit)
			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealOutcome(sim, result)
				if !result.Landed() {
					return
				}
				hunter.slowTarget(sim, target, 5116, 4*time.Second, .5)
				// Trait 134477 / curve 82813: exact 4/8/12/16/20%.
				if target.PseudoStats.CanBeStunned && sim.Proc(.04*float64(hunter.Talents.ImprovedConcussiveShot), "Improved Concussive Shot") {
					hunter.stunTarget(sim, target, 19410, 3*time.Second)
				}
			})
		},
	})
}

func (hunter *Hunter) registerDisengageSpell() {
	if !hunter.Env.IsForever() {
		return
	}
	timer := hunter.NewTimer()
	for rank := 1; rank <= 3; rank++ {
		level := []int{0, 20, 34, 48}[rank]
		if int(hunter.Level) < level {
			continue
		}
		// 70291 effect63 already contains the doubled reduction. Its -3
		// per-level adjustment is capped by SpellLevels MaxLevel, not doubled again.
		reduction := []float64{0, 280, 560, 810}[rank] + 3*float64(min(10, int(hunter.Level)-level))
		hunter.Disengage = hunter.RegisterSpell(core.SpellConfig{
			ActionID:      core.ActionID{SpellID: []int32{0, 781, 14272, 14273}[rank]},
			SpellCode:     SpellCode_HunterDisengage,
			SpellSchool:   core.SpellSchoolPhysical,
			DefenseType:   core.DefenseTypeMelee,
			ProcMask:      core.ProcMaskMeleeMHSpecial,
			Flags:         core.SpellFlagAPL | core.SpellFlagMeleeMetrics,
			Rank:          rank,
			RequiredLevel: level,
			ManaCost:      core.ManaCostOptions{FlatCost: []float64{0, 50, 100, 150}[rank]},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{GCD: core.GCDDefault}, IgnoreHaste: true,
				CD: core.Cooldown{Timer: timer, Duration: 5 * time.Second},
			},
			ExtraCastCondition: func(_ *core.Simulation, _ *core.Unit) bool {
				return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance
			},
			ThreatMultiplier: 1,
			FlatThreatBonus:  -reduction,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
			},
		})
	}
}

func (hunter *Hunter) procEntrapment(sim *core.Simulation, target *core.Unit) {
	if hunter.Env.IsForever() && hunter.Talents.Entrapment > 0 {
		// Trait curve82791 specifies 1..5 seconds, not a random proc chance.
		// Invoked only on initial successful trap activation, never periodic ticks.
		// DR is a separate scenario opt-in; its target policy/reset anchor
		// remain explicitly qualified by the shared control helper.
		hunter.rootTarget(sim, target, 19185, time.Duration(hunter.Talents.Entrapment)*time.Second)
	}
}

func (hunter *Hunter) registerFrostTrapSpell(timer *core.Timer) {
	if !hunter.Env.IsForever() || hunter.Level < 28 {
		return
	}
	hunter.FrostTrap = hunter.RegisterSpell(core.SpellConfig{
		ActionID:      core.ActionID{SpellID: 13809},
		SpellCode:     SpellCode_HunterFrostTrap,
		SpellSchool:   core.SpellSchoolFrost,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | SpellFlagTrap,
		RequiredLevel: 28,
		ManaCost:      core.ManaCostOptions{FlatCost: 60},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault}, IgnoreHaste: true,
			CD: core.Cooldown{Timer: timer, Duration: 30 * time.Second},
		},
		ExtraCastCondition: func(_ *core.Simulation, _ *core.Unit) bool { return hunter.DistanceFromTarget <= 5 },
		ThreatMultiplier:   1,
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			for _, target := range hunter.Env.Encounter.TargetUnits {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					hunter.slowTarget(sim, target, 13810, time.Duration(float64(30*time.Second)*(1+.15*float64(hunter.Talents.CleverTraps))), .6)
					hunter.procEntrapment(sim, target)
				}
			}
		},
	})
}
