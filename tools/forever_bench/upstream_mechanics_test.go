//go:build with_db

package main

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/warrior"
)

func TestForeverFlurrySpendsEachWhiteSwing(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	w := sim.Raid.Parties[0].Players[0].(warrior.WarriorAgent).GetWarrior()
	if w.Talents.Flurry == 0 {
		t.Fatal("Fury profile has no Flurry")
	}
	id := []int32{12319, 12971, 12972, 12973, 12974}[w.Talents.Flurry-1]
	flurry := w.GetAura(fmt.Sprintf("Flurry Proc (%d)", id))
	trigger := w.GetAura(fmt.Sprintf("Flurry Consume Trigger - %d", id))
	if flurry == nil || trigger == nil {
		t.Fatal("the Fury profile is missing its Flurry auras")
	}
	flurry.Activate(sim)
	flurry.SetStacks(sim, 3)
	for i := 0; i < 2; i++ {
		trigger.OnSpellHitDealt(trigger, sim, &core.Spell{ProcMask: core.ProcMaskMeleeMHAuto},
			&core.SpellResult{Outcome: core.OutcomeHit, Target: sim.Encounter.TargetUnits[0]})
	}
	if got := flurry.GetStacks(); got != 1 {
		t.Errorf("two same-batch white swings left %d Flurry stacks, want one", got)
	}
}

func TestForeverRecklessnessCoversEveryDamageSchool(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	w := sim.Raid.Parties[0].Players[0].(warrior.WarriorAgent).GetWarrior()
	aura := w.GetAura("Recklessness")
	if aura == nil {
		t.Fatal("Recklessness aura not registered")
	}
	for _, stat := range []stats.Stat{stats.MeleeCrit, stats.SpellCrit} {
		before := w.GetStat(stat)
		aura.Activate(sim)
		if got, want := w.GetStat(stat)-before, float64(100*core.CritRatingPerCritChance); math.Abs(got-want) > 1e-7 {
			t.Errorf("%v gained %g crit from Recklessness, want %g", stat, got, want)
		}
		aura.Deactivate(sim)
		if got := w.GetStat(stat); math.Abs(got-before) > 1e-7 {
			t.Errorf("%v retained %g crit after Recklessness expired", stat, got-before)
		}
	}
}

func TestForeverQueuedWarriorSwingChangesOffHandHitTable(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	w := sim.Raid.Parties[0].Players[0].(warrior.WarriorAgent).GetWarrior()
	if w.PseudoStats.DisableDWMissPenalty {
		t.Fatal("dual-wield miss penalty disabled without a queued attack")
	}

	for _, queued := range []*warrior.WarriorSpell{w.HeroicStrike, w.Cleave} {
		aura := w.GetAura("HS/Cleave Queue Aura-" + queued.ActionID.String())
		if aura == nil {
			t.Fatalf("no queue aura for %s", queued.ActionID)
		}
		aura.Activate(sim)
		if !w.PseudoStats.DisableDWMissPenalty {
			t.Fatalf("off-hand still pays dual-wield miss penalty while %s is queued", queued.ActionID)
		}
		aura.Deactivate(sim)
		if w.PseudoStats.DisableDWMissPenalty {
			t.Fatalf("off-hand retains queued hit table after %s expires", queued.ActionID)
		}
	}
}

func TestForeverOverpowerUsesItsTriggeredWindow(t *testing.T) {
	req := racialFixture("fury", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	w := sim.Raid.Parties[0].Players[0].(warrior.WarriorAgent).GetWarrior()
	if got := w.OverpowerAura.ActionID.SpellID; got != 1282733 {
		t.Errorf("Overpower window uses %d, want the client's triggered aura 1282733", got)
	}
	if got := w.OverpowerAura.Duration; got != 5*time.Second {
		t.Errorf("Overpower window lasts %s, want five seconds", got)
	}
}

func TestForeverJudgementsUseOneMeleeHitRoll(t *testing.T) {
	req := racialFixture("retribution", proto.Race_RaceDwarf)
	player := req.Raid.Parties[0].Players[0]
	player.Rotation = &proto.APLRotation{}
	player.BonusStats = &proto.UnitStats{Stats: stats.Stats{}.ToFloatArray()}
	// Deliberately separate hit types to expose an accidental spell-hit roll.
	player.BonusStats.Stats[stats.MeleeHit] = 20
	player.BonusStats.Stats[stats.SpellHit] = -100
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	p := sim.Raid.Parties[0].Players[0].GetCharacter()
	target := sim.Encounter.TargetUnits[0]

	for _, id := range []int32{20965, 20286} {
		spell := p.GetSpell(core.ActionID{SpellID: id})
		if spell == nil {
			t.Fatalf("missing Judgement %d", id)
		}
		if spell.DefenseType != core.DefenseTypeMelee {
			t.Fatalf("Judgement %d defense type %v, want melee", id, spell.DefenseType)
		}
		for i := 0; i < 256; i++ {
			spell.ApplyEffects(sim, target, spell)
		}
		metrics := spell.SpellMetrics[target.UnitIndex]
		if metrics.Misses != 0 || metrics.Hits+metrics.Crits == 0 {
			t.Errorf("Judgement %d: %d misses, %d hits and %d crits with capped melee hit and no spell hit",
				id, metrics.Misses, metrics.Hits, metrics.Crits)
		}
	}
}

func TestForeverTwistEchoWaitsForWhiteSwing(t *testing.T) {
	req := racialFixture("retribution", proto.Race_RaceDwarf)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	character := sim.Raid.Parties[0].Players[0].GetCharacter()
	command := character.GetAura("Echo of Command")
	righteousness := character.GetAura("Echo of Righteousness")
	if command == nil || righteousness == nil {
		t.Fatal("the Retribution profile is missing its distinct Echo auras")
	}
	for _, id := range []int32{20919, 20293, 20919} {
		seal := character.GetSpell(core.ActionID{SpellID: id})
		if seal == nil {
			t.Fatalf("missing seal %d", id)
		}
		seal.ApplyEffects(sim, sim.Encounter.TargetUnits[0], seal)
	}
	if !command.IsActive() || !righteousness.IsActive() {
		t.Fatal("replacing two different seals did not bank both Echoes")
	}
	command.OnSpellHitDealt(command, sim, &core.Spell{ProcMask: core.ProcMaskMeleeMHSpecial},
		&core.SpellResult{Outcome: core.OutcomeHit, Target: sim.Encounter.TargetUnits[0]})
	if !command.IsActive() || !righteousness.IsActive() {
		t.Fatal("a melee-classified Judgement spent the seal Echo before an autoattack")
	}
	command.OnSpellHitDealt(command, sim, &core.Spell{ProcMask: core.ProcMaskMeleeMHAuto},
		&core.SpellResult{Outcome: core.OutcomeHit, Target: sim.Encounter.TargetUnits[0]})
	if command.IsActive() || !righteousness.IsActive() {
		t.Fatal("spending Command's Echo also spent Righteousness's distinct charge")
	}
}

func TestForeverJudgementOfCrusaderUsesActualRank(t *testing.T) {
	req := racialFixture("retribution", proto.Race_RaceUndead)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	req.Raid.Debuffs.JudgementOfTheCrusader = proto.TristateEffect_TristateEffectMissing
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	character := sim.Raid.Parties[0].Players[0].GetCharacter()
	target := sim.Encounter.TargetUnits[0]
	start := target.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly]
	for _, test := range []struct {
		spellID int32
		bonus   float64
	}{
		{21183, 23},
		{20303, 161},
	} {
		spell := character.GetSpell(core.ActionID{SpellID: test.spellID})
		if spell == nil {
			t.Fatalf("Judgement rank %d is missing", test.spellID)
		}
		spell.ApplyEffects(sim, target, spell)
		if got := target.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] - start; math.Abs(got-test.bonus) > 0.001 {
			t.Fatalf("Judgement rank %d contributed %.1f flat Holy damage, want %.1f", test.spellID, got, test.bonus)
		}
	}
}
