//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/hunter"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

type hunterUtilitySettings struct {
	second bool
	anchor proto.CrowdControlDrResetAnchor
}

func hunterUtilityFixture(fields map[string]int, control, dr bool, settings ...hunterUtilitySettings) (*core.Simulation, *hunter.Hunter) {
	req := historyTalentFixture("survival", fields)
	p := req.Raid.Parties[0].Players[0]
	p.ForeverTier1Bonuses = false
	p.GetHunter().Options.PetType = proto.Hunter_Options_PetNone
	p.Consumes, p.Buffs = &proto.Consumes{}, &proto.IndividualBuffs{}
	req.Raid.Buffs, req.Raid.Debuffs = &proto.RaidBuffs{}, &proto.Debuffs{}
	req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
	req.Encounter.Targets[0].CanBeStunned = control
	req.Encounter.Targets[0].CanBeRooted = control
	req.Encounter.Targets[0].CanBeSlowed = control
	req.Encounter.Targets[0].UseCrowdControlDiminishingReturns = dr
	if len(settings) > 0 {
		req.Encounter.Targets[0].CrowdControlDrResetAnchor = settings[0].anchor
		if settings[0].second {
			req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, googleProto.Clone(p).(*proto.Player))
		}
	}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	for _, unit := range sim.Environment.AllUnits {
		unit.AutoAttacks.CancelAutoSwing(sim)
		if unit.Type == core.PlayerUnit {
			unit.CancelGCDTimer(sim)
		}
	}
	return sim, sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
}

func utilityAdvance(sim *core.Simulation, until time.Duration) {
	sim.AddPendingAction(&core.PendingAction{NextActionAt: until, OnAction: func(*core.Simulation) {}})
	for sim.CurrentTime < until {
		sim.Step()
	}
}

func TestOctober9HunterUtilityBossDefaultsAndRegistration(t *testing.T) {
	sim, h := hunterUtilityFixture(map[string]int{"improvedWingClip": 3, "improvedConcussiveShot": 5, "entrapment": 5}, false, false)
	target := h.CurrentTarget
	if target.PseudoStats.CanBeStunned || target.PseudoStats.CanBeRooted || target.PseudoStats.CanBeSlowed || target.PseudoStats.UseCrowdControlDiminishingReturns {
		t.Fatal("boss baseline silently gained CC eligibility/DR")
	}
	if h.ConcussiveShot == nil || h.FrostTrap == nil || h.Disengage == nil {
		t.Fatal("source-supported utility spells not registered")
	}
	if h.ConcussiveShot.CD.Duration != 12*time.Second || h.FrostTrap.CD.Duration != 30*time.Second || h.Disengage.CD.Duration != 5*time.Second {
		t.Fatal("utility cooldowns differ from client")
	}
	h.DistanceFromTarget = 0
	for i := 0; i < 100; i++ {
		h.WingClip.ApplyEffects(sim, target, h.WingClip)
	}
	if target.IsRooted() || target.PseudoStats.Stunned || target.MovementHandler.MoveSpeed != 7 {
		t.Fatal("immune raid target was controlled")
	}
	h.ConcussiveShot.BonusHitRating = 10000
	h.DistanceFromTarget = 30
	h.ConcussiveShot.ApplyEffects(sim, target, h.ConcussiveShot)
	utilityAdvance(sim, time.Second)
	if target.PseudoStats.Stunned || target.MovementHandler.MoveSpeed != 7 {
		t.Fatal("immune target received Concussive control")
	}
	if h.ConcussiveShot.SpellMetrics[target.UnitIndex].TotalDamage != 0 {
		t.Fatal("utility shot invented direct damage")
	}
}

func TestOctober9HunterImprovedConcussiveStun(t *testing.T) {
	sim, h := hunterUtilityFixture(map[string]int{"improvedConcussiveShot": 5}, true, false)
	target := h.CurrentTarget
	h.DistanceFromTarget = 30
	h.ConcussiveShot.BonusHitRating = 10000
	for i := 0; i < 100 && !target.PseudoStats.Stunned; i++ {
		h.ConcussiveShot.ApplyEffects(sim, target, h.ConcussiveShot)
		utilityAdvance(sim, sim.CurrentTime+time.Second)
	}
	aura := target.GetAuraByID(core.ActionID{SpellID: 19410, Tag: h.Index + 1})
	if aura == nil || !aura.IsActive() || aura.Duration != 3*time.Second {
		t.Fatal("Improved Concussive Shot did not apply sourced three-second stun")
	}
	if target.MovementHandler.MoveSpeed != 3.5 {
		t.Fatal("Concussive Shot did not apply 50% slow")
	}
	utilityAdvance(sim, sim.CurrentTime+5*time.Second)
	if target.PseudoStats.Stunned || target.MovementHandler.MoveSpeed != 7 {
		t.Fatal("stun/slow expiry failed to restore state")
	}
}

func TestOctober9HunterControlOptInRoundTrip(t *testing.T) {
	original := &proto.Target{CanBeStunned: true, CanBeRooted: true, CanBeSlowed: true, UseCrowdControlDiminishingReturns: true, CrowdControlDrResetAnchor: proto.CrowdControlDrResetAnchor_CrowdControlDrResetAfterApplication}
	data, err := protojson.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	restored := &proto.Target{}
	if err := protojson.Unmarshal(data, restored); err != nil {
		t.Fatal(err)
	}
	if !restored.CanBeStunned || !restored.CanBeRooted || !restored.CanBeSlowed || !restored.UseCrowdControlDiminishingReturns || restored.CrowdControlDrResetAnchor != original.CrowdControlDrResetAnchor {
		t.Fatal("CC opt-in did not survive native export/import")
	}
}

func TestOctober9HunterWingClipAndInitialEntrapment(t *testing.T) {
	sim, h := hunterUtilityFixture(map[string]int{"improvedWingClip": 3, "entrapment": 5}, true, true)
	target := h.CurrentTarget
	h.DistanceFromTarget = 0
	h.WingClip.BonusHitRating = 10000
	for i := 0; i < 500 && !target.IsRooted(); i++ {
		h.WingClip.ApplyEffects(sim, target, h.WingClip)
	}
	if !target.IsRooted() || math.Abs(target.MovementHandler.MoveSpeed-2.8) > 1e-7 {
		t.Fatal("max-rank Wing Clip did not slow/root opted-in target")
	}
	wing := target.GetAuraByID(core.ActionID{SpellID: 19229, Tag: h.Index + 1})
	if wing == nil || wing.Duration != 5*time.Second {
		t.Fatal("Wing Clip root duration is not five seconds")
	}
	wing.Deactivate(sim)
	utilityAdvance(sim, 16*time.Second)
	h.FrostTrap.BonusHitRating = 10000
	h.FrostTrap.ApplyEffects(sim, target, h.FrostTrap)
	root := target.GetAuraByID(core.ActionID{SpellID: 19185, Tag: h.Index + 1})
	if root == nil || !root.IsActive() || root.Duration != 5*time.Second {
		t.Fatal("initial Frost Trap activation did not apply Entrapment")
	}
	firstExpiry := root.ExpiresAt()
	utilityAdvance(sim, sim.CurrentTime+4*time.Second)
	if root.ExpiresAt() != firstExpiry {
		t.Fatal("Frost Trap periodic field refreshed Entrapment/DR")
	}
	utilityAdvance(sim, sim.CurrentTime+2*time.Second)
	if target.IsRooted() {
		t.Fatal("Entrapment did not expire without periodic reapplication")
	}
}

func TestOctober9HunterDisengageDoubledThreat(t *testing.T) {
	sim, h := hunterUtilityFixture(nil, false, false)
	h.DistanceFromTarget = 0
	h.Disengage.BonusHitRating = 10000
	target := h.CurrentTarget
	for i := 0; i < 100; i++ {
		h.Disengage.ApplyEffects(sim, target, h.Disengage)
	}
	metrics := h.Disengage.SpellMetrics[target.UnitIndex]
	landed := metrics.Hits
	if landed == 0 || math.Abs(metrics.TotalThreat/float64(landed)/h.PseudoStats.ThreatMultiplier+840) > 1e-7 {
		t.Fatal("level60 rank3 Disengage did not use doubled base810 + capped30 level scaling")
	}
	if metrics.TotalDamage != 0 {
		t.Fatal("Disengage invented damage")
	}
}

func TestOctober9CrowdControlDRSharedRefreshBreakAndReset(t *testing.T) {
	sim, h := hunterUtilityFixture(map[string]int{"entrapment": 5}, true, true, hunterUtilitySettings{second: true})
	target := h.CurrentTarget
	other := sim.Raid.Parties[0].Players[1].(*hunter.Hunter)
	a := target.GetAuraByID(core.ActionID{SpellID: 19185, Tag: h.Index + 1})
	b := target.GetAuraByID(core.ActionID{SpellID: 19185, Tag: other.Index + 1})
	if !target.ApplyCrowdControl(sim, a, 8*time.Second, core.CrowdControlDRRoot) || a.Duration != 8*time.Second {
		t.Fatal("first control not full duration")
	}
	sim.CurrentTime = time.Second
	if !target.ApplyCrowdControl(sim, a, 8*time.Second, core.CrowdControlDRRoot) || a.Duration != 4*time.Second {
		t.Fatal("refresh did not halve next application")
	}
	if !target.ApplyCrowdControl(sim, b, 8*time.Second, core.CrowdControlDRRoot) || b.Duration != 2*time.Second {
		t.Fatal("other caster did not share target DR")
	}
	if target.ApplyCrowdControl(sim, b, 8*time.Second, core.CrowdControlDRRoot) {
		t.Fatal("fourth application was not immune")
	}
	a.Deactivate(sim) // Explicit break; overlapping root still prevents movement/reset.
	if !target.IsRooted() {
		t.Fatal("one break removed another caster's root")
	}
	sim.CurrentTime = 2 * time.Second
	b.Deactivate(sim)
	if target.IsRooted() {
		t.Fatal("last effect end retained root")
	}
	sim.CurrentTime = 16 * time.Second
	if target.ApplyCrowdControl(sim, a, 8*time.Second, core.CrowdControlDRRoot) {
		t.Fatal("qualified after-end reset anchor expired too early")
	}
	sim.CurrentTime = 17 * time.Second
	if !target.ApplyCrowdControl(sim, a, 8*time.Second, core.CrowdControlDRRoot) || a.Duration != 8*time.Second {
		t.Fatal("immunity did not reset fifteen seconds after final effect end")
	}
	a.Deactivate(sim)
	for _, unit := range sim.Environment.AllUnits {
		for _, aura := range unit.GetAuras() {
			if aura.IsActive() {
				aura.Deactivate(sim)
			}
		}
	}
	sim.Reset()
	if !target.ApplyCrowdControl(sim, a, 8*time.Second, core.CrowdControlDRRoot) || a.Duration != 8*time.Second {
		t.Fatal("DR leaked across iterations")
	}
}

func TestOctober9CrowdControlProvisionalAnchorSensitivity(t *testing.T) {
	for _, row := range []struct {
		name    string
		enabled bool
		anchor  proto.CrowdControlDrResetAnchor
		want    time.Duration
	}{
		{"off", false, proto.CrowdControlDrResetAnchor_CrowdControlDrResetAnchorDefault, 5 * time.Second},
		{"after-end", true, proto.CrowdControlDrResetAnchor_CrowdControlDrResetAfterEffectEnd, 2500 * time.Millisecond},
		{"after-application", true, proto.CrowdControlDrResetAnchor_CrowdControlDrResetAfterApplication, 5 * time.Second},
	} {
		t.Run(row.name, func(t *testing.T) {
			sim, h := hunterUtilityFixture(map[string]int{"entrapment": 5}, true, row.enabled, hunterUtilitySettings{anchor: row.anchor})
			target := h.CurrentTarget
			a := target.GetAuraByID(core.ActionID{SpellID: 19185, Tag: h.Index + 1})
			target.ApplyCrowdControl(sim, a, 5*time.Second, core.CrowdControlDRRoot)
			utilityAdvance(sim, 16*time.Second)
			if !target.ApplyCrowdControl(sim, a, 5*time.Second, core.CrowdControlDRRoot) || a.Duration != row.want {
				t.Fatalf("declared anchor sensitivity duration=%v want=%v", a.Duration, row.want)
			}
			// Wing Clip's current client category mask is zero, not Root1.
			wing := target.GetAuraByID(core.ActionID{SpellID: 19229, Tag: h.Index + 1})
			for i := 0; i < 4; i++ {
				if !target.ApplyCrowdControl(sim, wing, 5*time.Second, 0) || wing.Duration != 5*time.Second {
					t.Fatal("mask-zero Wing Clip acquired invented DR")
				}
			}
		})
	}
}

func TestOctober9MovementMidMoveSlowRootAndStun(t *testing.T) {
	sim, h := hunterUtilityFixture(nil, true, false)
	target := h.CurrentTarget
	target.CurrentTarget = &h.Unit
	target.DistanceFromTarget = 0
	target.MoveTo(14, sim)
	utilityAdvance(sim, time.Second)
	id := core.ActionID{SpellID: 5116}
	target.AddMoveSpeedModifierDynamic(sim, &id, .5)
	if math.Abs(target.DistanceFromTarget-7) > 1e-6 {
		t.Fatal("mid-move slow did not settle elapsed normal-speed distance")
	}
	utilityAdvance(sim, 2*time.Second)
	target.RemoveMoveSpeedModifierDynamic(sim, &id)
	if math.Abs(target.DistanceFromTarget-10.5) > 1e-6 {
		t.Fatal("slow did not change remaining movement timing")
	}
	target.AddRoot(sim)
	target.AddRoot(sim)
	before := target.DistanceFromTarget
	utilityAdvance(sim, 3*time.Second)
	if target.DistanceFromTarget != before || target.IsMoving() {
		t.Fatal("root moved/scheduled an invalid zero-speed action")
	}
	target.AddStun(sim)
	target.RemoveRoot(sim)
	target.RemoveRoot(sim)
	utilityAdvance(sim, 4*time.Second)
	if target.DistanceFromTarget != before || target.IsMoving() {
		t.Fatal("root expiry resumed movement while stunned")
	}
	target.RemoveStun(sim)
	utilityAdvance(sim, 4501*time.Millisecond)
	if math.Abs(target.DistanceFromTarget-14) > 1e-6 || target.IsMoving() || target.MovementHandler.MoveSpeed != 7 {
		t.Fatalf("control expiry did not restore/finish movement: distance=%g moving=%t speed=%g", target.DistanceFromTarget, target.IsMoving(), target.MovementHandler.MoveSpeed)
	}
	target.AddRoot(sim)
	target.MoveTo(0, sim)
	utilityAdvance(sim, 5*time.Second)
	if target.DistanceFromTarget != 14 || target.IsMoving() {
		t.Fatal("rooted new movement did not reject without divide-by-zero")
	}
	target.RemoveRoot(sim)
	strong := core.ActionID{SpellID: 13810}
	target.AddMoveSpeedModifierDynamic(sim, &id, .5)
	target.AddMoveSpeedModifierDynamic(sim, &strong, .6)
	if math.Abs(target.MovementHandler.MoveSpeed-2.8) > 1e-7 {
		t.Fatal("overlapping slows stacked instead of using strongest")
	}
	target.RemoveMoveSpeedModifierDynamic(sim, &strong)
	if math.Abs(target.MovementHandler.MoveSpeed-3.5) > 1e-7 {
		t.Fatal("strong slow expiry removed the weaker slow")
	}
	target.RemoveMoveSpeedModifierDynamic(sim, &id)
	if target.MovementHandler.MoveSpeed != 7 {
		t.Fatal("final slow expiry did not restore speed")
	}
}
