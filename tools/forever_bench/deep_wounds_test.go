//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/warrior"
)

func TestForeverDeepWoundsCannotCrit(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.SimOptions.Iterations = 250
	req.SimOptions.RandomSeed = 20295719
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		if action.Id.GetSpellId() != 12867 {
			continue
		}
		var ticks, critTicks int32
		for _, target := range action.Targets {
			ticks += target.Ticks
			critTicks += target.CritTicks
		}
		if ticks == 0 || critTicks != 0 {
			t.Fatalf("Deep Wounds: %d normal ticks, %d critical ticks", ticks, critTicks)
		}
		return
	}
	t.Fatal("Fury did not trigger Deep Wounds")
}

func deepWoundsSampleUnit() (*core.Simulation, *warrior.Warrior, *core.Unit) {
	req := historyTalentFixture("fury", map[string]int{"deepWounds": 1})
	req.Raid.Buffs, req.Raid.Debuffs, req.Raid.Parties[0].Buffs = nil, nil, nil
	p := req.Raid.Parties[0].Players[0]
	p.ForeverTier1Bonuses, p.Consumes, p.Buffs = false, nil, nil
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	w := sim.Raid.Parties[0].Players[0].(warrior.WarriorAgent).GetWarrior()
	return sim, w, sim.Encounter.TargetUnits[0]
}

func triggerSampleDeepWounds(sim *core.Simulation, w *warrior.Warrior, target *core.Unit) {
	w.OnSpellHitDealt(sim, w.AutoAttacks.MHAuto(), &core.SpellResult{Target: target, Outcome: core.OutcomeCrit})
}

func TestForeverDeepWoundsWeaponOnlyAndModifiers(t *testing.T) {
	sample := func(ap, damageMultiplier, targetMultiplier float64) float64 {
		sim, w, target := deepWoundsSampleUnit()
		w.AddStatDynamic(sim, stats.AttackPower, ap)
		w.PseudoStats.DamageDealtMultiplier *= damageMultiplier
		target.PseudoStats.DamageTakenMultiplier *= targetMultiplier
		triggerSampleDeepWounds(sim, w, target)
		sim.CurrentTime = 3 * time.Second
		w.DeepWounds.Dot(target).TickOnce(sim)
		return w.DeepWounds.SpellMetrics[target.UnitIndex].TotalDamage
	}
	base := sample(0, 1, 1)
	if base <= 0 {
		t.Fatal("Deep Wounds did not tick")
	}
	for _, variant := range []struct {
		name                     string
		ap, caster, target, want float64
	}{
		{"attack power", 2000, 1, 1, base},
		{"caster damage", 0, 1.4, 1, base},
		{"target damage", 0, 1, 1.5, base * 1.5},
	} {
		if got := sample(variant.ap, variant.caster, variant.target); math.Abs(got-variant.want) > 1e-8 {
			t.Errorf("%s: tick %v, want %v", variant.name, got, variant.want)
		}
	}
}

func TestForeverDeepWoundsRollsAndPreservesTick(t *testing.T) {
	sim, w, target := deepWoundsSampleUnit()
	weaponDamage := w.AutoAttacks.MH().AverageDamage() * .2
	triggerSampleDeepWounds(sim, w, target)
	dot := w.DeepWounds.Dot(target)
	if math.Abs(dot.SnapshotBaseDamage-weaponDamage/4) > 1e-8 {
		t.Errorf("initial tick payload %v, want %v", dot.SnapshotBaseDamage, weaponDamage/4)
	}
	sim.CurrentTime = 3 * time.Second
	dot.TickOnce(sim)
	dot.TickCount = 1 // TickOnce is the direct callback, not the scheduled ticker.
	sim.CurrentTime = 4 * time.Second
	triggerSampleDeepWounds(sim, w, target)
	wantTick := (weaponDamage + 3*weaponDamage/4) / 4
	if math.Abs(dot.SnapshotBaseDamage-wantTick) > 1e-8 {
		t.Errorf("rolled tick payload %v, want %v", dot.SnapshotBaseDamage, wantTick)
	}
	if got := dot.NextTickAt(); got != 6*time.Second {
		t.Errorf("refresh moved next tick to %v, want 6s", got)
	}
	if dot.MaxTicksRemaining() != 4 {
		t.Errorf("refresh has %d ticks, want 4", dot.MaxTicksRemaining())
	}
}
