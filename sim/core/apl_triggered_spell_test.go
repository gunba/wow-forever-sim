package core

import (
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"testing"
	"time"
)

func TestAPLRejectsTriggeredSpells(t *testing.T) {
	env := &Environment{Ruleset: proto.Ruleset_RulesetForever, Encounter: Encounter{Targets: []*Target{{}}}}
	u := &Unit{Env: env, Type: PlayerUnit, UnitIndex: 0, PseudoStats: stats.NewPseudoStats(), Metrics: NewUnitMetrics()}
	target := &Unit{Env: env, Type: EnemyUnit, UnitIndex: 1, PseudoStats: stats.NewPseudoStats(), Metrics: NewUnitMetrics()}
	u.CurrentTarget = target
	env.AllUnits = []*Unit{u, target}
	env.Raid = &Raid{AllUnits: []*Unit{u}, AllPlayerUnits: []*Unit{u}}
	rot := &APLRotation{unit: u}
	procCount := 0
	proc := u.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 1}, SpellSchool: SpellSchoolShadow, ProcMask: ProcMaskEmpty, DamageMultiplier: 1,
		ApplyEffects: func(*Simulation, *Unit, *Spell) { procCount++ }})
	u.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 2}, Flags: SpellFlagChanneled, SpellSchool: SpellSchoolShadow, ProcMask: ProcMaskEmpty, DamageMultiplier: 1})
	u.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 3}, SpellSchool: SpellSchoolShadow, ProcMask: ProcMaskEmpty, DamageMultiplier: 1,
		Dot: DotConfig{Aura: Aura{Label: "Triggered dot"}, NumberOfTicks: 1, TickLength: time.Second}})
	u.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 4}, SpellSchool: SpellSchoolHoly, ProcMask: ProcMaskEmpty, DamageMultiplier: 1,
		Shield: ShieldConfig{Aura: Aura{Label: "Triggered shield", Duration: time.Second}}})
	active := u.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 5}, Flags: SpellFlagAPL, SpellSchool: SpellSchoolShadow, ProcMask: ProcMaskEmpty, DamageMultiplier: 1})
	selfShield := u.RegisterSpell(SpellConfig{ActionID: ActionID{SpellID: 6}, Flags: SpellFlagAPL, SpellSchool: SpellSchoolHoly, ProcMask: ProcMaskEmpty, DamageMultiplier: 1, Shield: ShieldConfig{SelfOnly: true, Aura: Aura{Label: "Self shield", Duration: time.Second}}})
	if rot.newActionCastSpell(&proto.APLActionCastSpell{SpellId: proc.ActionID.ToProto()}) != nil {
		t.Error("imported APL can manually cast a proc")
	}
	if rot.newActionChannelSpell(&proto.APLActionChannelSpell{SpellId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 2}}, InstantInterrupt: true}) != nil {
		t.Error("channel action bypasses proc guard")
	}
	if rot.newActionMultidot(&proto.APLActionMultidot{SpellId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 3}}, MaxDots: 1}) != nil {
		t.Error("multidot action bypasses proc guard")
	}
	if rot.newActionMultishield(&proto.APLActionMultishield{SpellId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 4}}, MaxShields: 1}) != nil {
		t.Error("multishield action bypasses proc guard")
	}
	if rot.newActionCastSpell(&proto.APLActionCastSpell{SpellId: active.ActionID.ToProto()}) == nil {
		t.Error("active spell rejected")
	}
	if rot.newActionMultidot(&proto.APLActionMultidot{SpellId: active.ActionID.ToProto(), MaxDots: 1}) != nil {
		t.Error("multidot accepted a spell without a dot")
	}
	if rot.newActionMultishield(&proto.APLActionMultishield{SpellId: selfShield.ActionID.ToProto(), MaxShields: 1}) != nil {
		t.Error("multishield accepted a self-only shield")
	}
	if rot.GetAPLSpell(proc.ActionID.ToProto()) != proc {
		t.Error("read-only proc references rejected")
	}
	// Proc-triggered engine calls still work, without a player-APL action.
	sim := &Simulation{Environment: env, Duration: time.Second}
	proc.finalize()
	proc.Cast(sim, target)
	if procCount != 1 {
		t.Fatal("engine proc casting disabled")
	}
}
