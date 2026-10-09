package core

import (
	"math"
	"time"
)

// SpellCategories.DiminishType is a mask into SpellDiminish, not a target-type
// enum. Current Forever client 70291 rows 1 and 10 supply these group IDs.
const (
	CrowdControlDRRoot     int32 = 1
	CrowdControlDRStunProc int32 = 10
)

const crowdControlDRResetDuration = 15 * time.Second

type controlDiminishingReturns struct {
	applications  int
	activeEffects int
	resetAt       time.Duration
}

func (unit *Unit) resetCrowdControlDiminishingReturns() {
	for _, state := range unit.controlDiminishingReturns {
		*state = controlDiminishingReturns{}
	}
}

func (unit *Unit) crowdControlDRState(group int32) *controlDiminishingReturns {
	if group != CrowdControlDRRoot && group != CrowdControlDRStunProc {
		panic("unsupported crowd-control DR group")
	}
	if unit.controlDiminishingReturns == nil {
		unit.controlDiminishingReturns = make(map[int32]*controlDiminishingReturns)
	}
	if unit.controlDiminishingReturns[group] == nil {
		unit.controlDiminishingReturns[group] = &controlDiminishingReturns{}
	}
	return unit.controlDiminishingReturns[group]
}

// RegisterCrowdControlAura shares DR across casters/effects on this target.
// The opt-in model's reset clock starts after the final group effect ends
// (including an explicit break) by default. The target can explicitly select
// after-application instead. Both anchors are declared modeling assumptions,
// not measured Forever rules; the client target-policy field remains undecoded.
func (unit *Unit) RegisterCrowdControlAura(config Aura, group int32) *Aura {
	if group == 0 {
		return unit.GetOrRegisterAura(config)
	}
	if existing := unit.GetAura(config.Label); existing != nil {
		return existing
	}
	state := unit.crowdControlDRState(group)
	onGain, onExpire := config.OnGain, config.OnExpire
	config.OnGain = func(aura *Aura, sim *Simulation) {
		state.activeEffects++
		if !unit.PseudoStats.CrowdControlDRResetAfterApplication {
			state.resetAt = NeverExpires
		}
		if onGain != nil {
			onGain(aura, sim)
		}
	}
	config.OnExpire = func(aura *Aura, sim *Simulation) {
		if onExpire != nil {
			onExpire(aura, sim)
		}
		if state.activeEffects > 0 {
			state.activeEffects--
		}
		if state.activeEffects == 0 && !unit.PseudoStats.CrowdControlDRResetAfterApplication {
			state.resetAt = sim.CurrentTime + crowdControlDRResetDuration
		}
	}
	return unit.RegisterAura(config)
}

// ApplyCrowdControl uses sourced 0.5 scaling and immunity after three successes
// ONLY when the scenario explicitly enables DR. Eligibility is a separate,
// per-mechanic check owned by the ability; existing boss scenarios opt into neither.
func (unit *Unit) ApplyCrowdControl(sim *Simulation, aura *Aura, baseDuration time.Duration, group int32) bool {
	duration := baseDuration
	if group != 0 && unit.Env.IsForever() && unit.PseudoStats.UseCrowdControlDiminishingReturns {
		state := unit.crowdControlDRState(group)
		if (state.activeEffects == 0 || unit.PseudoStats.CrowdControlDRResetAfterApplication) && sim.CurrentTime >= state.resetAt {
			state.applications = 0
		}
		if state.applications >= 3 {
			return false
		}
		duration = time.Duration(float64(baseDuration) * math.Pow(.5, float64(state.applications)))
		state.applications++
		if unit.PseudoStats.CrowdControlDRResetAfterApplication {
			state.resetAt = sim.CurrentTime + crowdControlDRResetDuration
		}
	}
	aura.Duration = duration
	aura.Activate(sim)
	return true
}
