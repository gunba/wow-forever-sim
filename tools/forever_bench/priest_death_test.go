//go:build with_db

package main

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/priest"
)

func deathFixture(talents ...string) (*core.Simulation, *priest.Priest) {
	req := racialFixture("shadow", proto.Race_RaceTroll)
	player := req.Raid.Parties[0].Players[0]
	player.TalentsString = ""
	if len(talents) > 0 {
		player.TalentsString = talents[0]
	}
	player.Equipment = &proto.EquipmentSpec{}
	player.Consumes = nil
	player.Buffs = nil
	player.ForeverTier1Bonuses = false
	player.BonusStats = &proto.UnitStats{Stats: stats.Stats{stats.SpellHit: 100, stats.SpellCrit: -100}.ToFloatArray()}
	player.Rotation = &proto.APLRotation{}
	req.Raid.Buffs, req.Raid.Debuffs, req.Raid.Parties[0].Buffs = nil, nil, nil
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim, sim.Raid.Parties[0].Players[0].(priest.PriestAgent).GetPriest()
}

func TestPenanceImmediateThenOneSecondBolts(t *testing.T) {
	req := racialFixture("smite", proto.Race_RaceTroll)
	player := req.Raid.Parties[0].Players[0]
	player.Rotation = &proto.APLRotation{}
	player.BonusStats = &proto.UnitStats{Stats: stats.Stats{stats.SpellHit: 100}.ToFloatArray()}
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	p := sim.Raid.Parties[0].Players[0].(priest.PriestAgent).GetPriest()
	spell := p.Penance
	if spell == nil {
		t.Fatal("Penance was not registered")
	}
	sim.CurrentTime = 0
	if !spell.Cast(sim, p.CurrentTarget) {
		t.Fatal("Penance was not castable")
	}
	damage := spell.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
	if damage <= 0 {
		t.Fatal("first Penance bolt did not land immediately")
	}
	for _, at := range []time.Duration{time.Second, 2 * time.Second} {
		for sim.CurrentTime < at {
			if sim.Step() {
				t.Fatalf("fight ended before Penance tick at %s", at)
			}
		}
		next := spell.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
		if next <= damage || sim.CurrentTime != at {
			t.Fatalf("Penance had no bolt at %s (time %s, damage %.1f → %.1f)", at, sim.CurrentTime, damage, next)
		}
		damage = next
	}
}

func TestEarlyDemiseExecuteScopeAndReset(t *testing.T) {
	sim, p := deathFixture("--" + strings.Repeat("0", 15) + "2")
	death := p.ShadowWordDeath[4]
	base := death.BonusCritRating
	execute := p.GetAura("Early Demise")
	if execute == nil || execute.IsActive() {
		t.Fatal("Early Demise missing or active before execute")
	}
	for !sim.IsExecutePhase20() {
		if sim.Step() {
			t.Fatal("fight ended before execute")
		}
	}
	if !execute.IsActive() || math.Abs(death.BonusCritRating-base-30*core.CritRatingPerCritChance) > 1e-8 {
		t.Fatal("Death did not gain 30% crit in execute")
	}
	if p.GetSpell(core.ActionID{SpellID: 1309598}).BonusCritRating != 0 {
		t.Fatal("Early Demise affected backlash")
	}
	sim.Cleanup()
	sim.Reset()
	if execute.IsActive() || death.BonusCritRating != base {
		t.Fatal("execute crit leaked into the next iteration")
	}
}

func TestShadowWordDeathRanksAndSharedCooldown(t *testing.T) {
	sim, p := deathFixture()
	for rank := 1; rank <= priest.ShadowWordDeathRanks; rank++ {
		spell := p.ShadowWordDeath[rank]
		if spell == nil {
			t.Fatalf("missing learnable rank %d", rank)
		}
		if spell.RequiredLevel != []int{0, 32, 40, 48, 56}[rank] ||
			spell.BonusCoefficient != .429 || spell.CD.Duration != 15*time.Second ||
			spell.DefaultCast.CastTime != 0 || spell.DefaultCast.GCD != core.GCDDefault {
			t.Fatalf("rank %d: wrong spell configuration", rank)
		}
	}
	before := p.CurrentMana()
	if !p.ShadowWordDeath[4].Cast(sim, p.CurrentTarget) {
		t.Fatal("rank 4 was not castable")
	}
	if math.Abs(before-p.CurrentMana()-340) > 1e-8 {
		t.Fatal("rank 4 did not cost 340 mana")
	}
	sim.CurrentTime = 2 * time.Second
	if p.ShadowWordDeath[1].CanCast(sim, p.CurrentTarget) {
		t.Fatal("downranking bypassed Death's shared cooldown")
	}
}

func TestShadowWordDeathBacklash(t *testing.T) {
	for _, lethal := range []bool{false, true} {
		sim, p := deathFixture()
		p.PseudoStats.DamageDealtMultiplier *= 2
		p.PseudoStats.DamageTakenMultiplier *= .5
		if lethal {
			sim.Encounter.EndFightAtHealth = 1
		}
		if !p.ShadowWordDeath[4].Cast(sim, p.CurrentTarget) {
			t.Fatal("Death was not castable")
		}
		dealt := p.ShadowWordDeath[4].SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
		if dealt <= 0 {
			t.Fatal("fixture did not land Death")
		}
		backlash := p.GetSpell(core.ActionID{SpellID: 1309598})
		wantHits := int32(1)
		if lethal {
			wantHits = 0
		}
		if got := backlash.SpellMetrics[p.UnitIndex].Hits; got != wantHits {
			t.Fatalf("lethal=%v: backlash hits %v, want %v", lethal, got, wantHits)
		}
		for _, target := range backlash.SpellMetrics {
			if target.TotalDamage != 0 {
				t.Fatal("backlash entered outgoing DPS metrics")
			}
		}
	}
}
