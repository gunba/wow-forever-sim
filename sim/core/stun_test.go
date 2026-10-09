package core

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestStunCastLifecycleCallbacks(t *testing.T) {
	u := &Unit{Type: PlayerUnit, enabled: true, Env: &Environment{Ruleset: proto.Ruleset_RulesetForever}, GCD: new(Timer), PseudoStats: stats.NewPseudoStats()}
	u.gcdAction = &PendingAction{OnAction: func(*Simulation) {}}
	u.MovementHandler = &MovementHandler{unit: u, moveAura: &Aura{}}
	sim := &Simulation{Duration: time.Minute, Options: &proto.SimOptions{Interactive: true}, pendingActions: []*PendingAction{{NextActionAt: NeverExpires}}}
	target := &Unit{}
	started, ended, interrupted := 0, 0, 0
	cast := Cast{GCD: GCDDefault, CastTime: 2 * time.Second}
	spell := &Spell{Unit: u, DefaultCast: cast, Flags: SpellFlagNoOnCastComplete, SpellMetrics: make([]SpellMetrics, 1), ApplyEffects: func(*Simulation, *Unit, *Spell) {}}
	spell.castFn = spell.makeCastFunc(CastConfig{DefaultCast: cast, IgnoreHaste: true,
		OnCastStart: func(*Simulation, *Spell) { started++ },
		OnCastEnd: func(_ *Simulation, _ *Spell, cancel bool) {
			ended++
			if cancel {
				interrupted++
			}
		},
	})
	u.GCD.Set(time.Second)
	if spell.Cast(sim, target) || started != 0 || ended != 0 {
		t.Fatal("rejected attempt invoked lifecycle hooks")
	}
	u.GCD.Set(0)
	if !spell.Cast(sim, target) || started != 1 {
		t.Fatal("accepted cast missing start")
	}
	sim.CurrentTime = 500 * time.Millisecond
	u.AddStun(sim)
	if ended != 1 || interrupted != 1 || u.IsCasting(sim) || u.GCD.ReadyAt() != GCDDefault || spell.Cast(sim, target) || spell.CanCast(sim, target) {
		t.Fatal("stun/cancellation lifecycle or remaining GCD incorrect")
	}
	u.InterruptCast(sim)
	u.AddStun(sim)
	u.RemoveStun(sim)
	if !u.PseudoStats.Stunned || ended != 1 {
		t.Fatal("nested stun ended too early or callback repeated")
	}
	u.RemoveStun(sim)
	if u.PseudoStats.Stunned {
		t.Fatal("last stun not removed")
	}
	sim.CurrentTime = 2 * time.Second
	if !spell.Cast(sim, target) {
		t.Fatal("could not recast after release")
	}
	sim.CurrentTime = u.Hardcast.Expires
	hc := u.Hardcast
	u.Hardcast.Expires = startingCDTime
	hc.OnComplete(sim, target)
	u.InterruptCast(sim)
	if started != 2 || ended != 2 || interrupted != 1 {
		t.Fatalf("callbacks start/end/cancel=%d/%d/%d", started, ended, interrupted)
	}
}

func TestStunAcrossPullDefersAutos(t *testing.T) {
	u := &Unit{Type: EnemyUnit, enabled: true, PseudoStats: stats.NewPseudoStats()}
	u.AutoAttacks = AutoAttacks{AutoSwingMelee: true, mh: WeaponAttack{unit: u, Weapon: Weapon{SwingSpeed: 2}}}
	sim := &Simulation{Duration: time.Minute}
	u.AddStun(sim)
	u.AutoAttacks.startPull(sim)
	if u.AutoAttacks.enabled || len(sim.weaponAttacks) != 0 {
		t.Fatal("pull started swings while stunned")
	}
	u.RemoveStun(sim)
	if !u.AutoAttacks.enabled || len(sim.weaponAttacks) != 1 {
		t.Fatal("unstun lost deferred pull autos")
	}
	u.AddStun(sim)
	u.enabled = false
	u.RemoveStun(sim)
	if u.AutoAttacks.enabled || len(sim.weaponAttacks) != 0 {
		t.Fatal("unstun restarted a disabled actor")
	}
}

func TestStunPreservesPriorFlagAndDisabledState(t *testing.T) {
	u := &Unit{enabled: true, PseudoStats: stats.NewPseudoStats()}
	sim := &Simulation{Duration: time.Minute}
	u.PseudoStats.Stunned = true
	u.AddStun(sim)
	u.RemoveStun(sim)
	if !u.PseudoStats.Stunned {
		t.Fatal("preexisting flag erased")
	}
	u.PseudoStats.Stunned = false
	u.AddStun(sim)
	u.enabled = false
	u.RemoveStun(sim)
	if u.enabled || u.PseudoStats.Stunned {
		t.Fatal("release resurrected a disabled unit")
	}
}
