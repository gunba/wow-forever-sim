//go:build with_db

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/rogue"
)

func rogueMutilateSim(t *testing.T, fields map[string]int, ruleset proto.Ruleset) (*core.Simulation, *rogue.Rogue) {
	t.Helper()
	req := historyTalentFixture("mutilate", fields)
	p := req.Raid.Parties[0].Players[0]
	weapons := p.Equipment.Items[14:16]
	p.Equipment = &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 16)}
	for i := range p.Equipment.Items {
		p.Equipment.Items[i] = &proto.ItemSpec{}
	}
	for i, weapon := range weapons {
		p.Equipment.Items[14+i] = &proto.ItemSpec{Id: weapon.Id}
	}
	p.ForeverTier1Bonuses = false
	p.BonusStats = nil
	req.Encounter.Targets[0].Level = 60
	req.SimOptions.Ruleset = ruleset
	req.SimOptions.IsTest = true
	req.SimOptions.RandomSeed = 1
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	r := sim.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
	table := r.AttackTables[r.CurrentTarget.UnitIndex][r.Mutilate.CastType]
	table.BaseMissChance, table.BaseDodgeChance, table.BaseParryChance, table.BaseBlockChance = 0, 0, 0, 0
	for _, spell := range r.Spellbook {
		spell.BonusCritRating = -r.GetStat(stats.MeleeCrit)
	}
	return sim, r
}

func mutilateHand(t *testing.T, r *rogue.Rogue, mask core.ProcMask) *core.Spell {
	t.Helper()
	for _, spell := range r.Spellbook {
		if spell.ActionID.SpellID == r.Mutilate.ActionID.SpellID && spell.ProcMask == mask {
			return spell
		}
	}
	t.Fatalf("missing Mutilate hand with proc mask %v", mask)
	return nil
}

func TestRogueMutilateAvoidedCastHasNoHandStrikes(t *testing.T) {
	for _, outcome := range []string{"miss", "dodge", "parry"} {
		t.Run(outcome, func(t *testing.T) {
			sim, r := rogueMutilateSim(t, map[string]int{"mutilate": 1, "coldBlood": 1}, proto.Ruleset_RulesetForever)
			table := r.AttackTables[r.CurrentTarget.UnitIndex][r.Mutilate.CastType]
			r.PseudoStats.InFrontOfTarget = true
			switch outcome {
			case "miss":
				table.BaseMissChance = 1
			case "dodge":
				table.BaseDodgeChance = 1
			case "parry":
				table.BaseParryChance = 1
			}
			r.ColdBlood.Cast(sim, r.CurrentTarget)
			if !r.Mutilate.Cast(sim, r.CurrentTarget) {
				t.Fatal("Mutilate was not cast")
			}
			for _, mask := range []core.ProcMask{core.ProcMaskMeleeMHSpecial, core.ProcMaskMeleeOHSpecial} {
				spell := mutilateHand(t, r, mask)
				if got := spell.SpellMetrics[r.CurrentTarget.UnitIndex].Casts; got != 0 {
					t.Errorf("avoided parent cast executed %d %s strikes", got, spell.ActionID)
				}
			}
			if r.ComboPoints() != 0 || r.CurrentEnergy() != 88 {
				t.Errorf("avoided cast: combo points %d, energy %v; want 0 and 88", r.ComboPoints(), r.CurrentEnergy())
			}
			if r.GetAura("Cold Blood").IsActive() {
				t.Error("avoided Mutilate did not consume Cold Blood")
			}
		})
	}
}

func TestRogueMutilateHandStrikesDoNotRollAvoidance(t *testing.T) {
	sim, r := rogueMutilateSim(t, map[string]int{"mutilate": 1}, proto.Ruleset_RulesetForever)
	r.PseudoStats.InFrontOfTarget = true
	table := r.AttackTables[r.CurrentTarget.UnitIndex][r.Mutilate.CastType]
	table.BaseMissChance, table.BaseDodgeChance, table.BaseParryChance = 1, 1, 1
	for _, mask := range []core.ProcMask{core.ProcMaskMeleeMHSpecial, core.ProcMaskMeleeOHSpecial} {
		spell := mutilateHand(t, r, mask)
		spell.Cast(sim, r.CurrentTarget)
		metrics := spell.SpellMetrics[r.CurrentTarget.UnitIndex]
		if metrics.Misses+metrics.Dodges+metrics.Parries != 0 || metrics.Hits != 1 {
			t.Fatalf("%s rerolled avoidance: hits %d, misses %d, dodges %d, parries %d", spell.ActionID, metrics.Hits, metrics.Misses, metrics.Dodges, metrics.Parries)
		}
	}
}

func TestRogueMutilateBlockAndCrit(t *testing.T) {
	for _, inFront := range []bool{true, false} {
		t.Run(fmt.Sprintf("front%v", inFront), func(t *testing.T) {
			sim, r := rogueMutilateSim(t, map[string]int{"mutilate": 1, "coldBlood": 1, "sealFate": 5}, proto.Ruleset_RulesetForever)
			r.PseudoStats.InFrontOfTarget = inFront
			table := r.AttackTables[r.CurrentTarget.UnitIndex][r.Mutilate.CastType]
			table.BaseBlockChance = 1
			r.CurrentTarget.AddStatDynamic(sim, stats.BlockValue, 10000)
			r.ColdBlood.Cast(sim, r.CurrentTarget)
			if !r.Mutilate.Cast(sim, r.CurrentTarget) {
				t.Fatal("Mutilate was not cast")
			}
			wantCP := int32(3)
			if inFront {
				wantCP = 2
			}
			if r.ComboPoints() != wantCP || r.GetAura("Cold Blood").IsActive() {
				t.Fatalf("blocked/critical cast: combo points %d, Cold Blood %v", r.ComboPoints(), r.GetAura("Cold Blood").IsActive())
			}
			for _, mask := range []core.ProcMask{core.ProcMaskMeleeMHSpecial, core.ProcMaskMeleeOHSpecial} {
				spell := mutilateHand(t, r, mask)
				metrics := spell.SpellMetrics[r.CurrentTarget.UnitIndex]
				if inFront && (metrics.Blocks != 1 || metrics.Crits != 0 || metrics.TotalDamage != 0) ||
					!inFront && (metrics.Blocks != 0 || metrics.Crits != 1 || metrics.TotalDamage <= 0) {
					t.Fatalf("%s: blocks %d, crits %d, damage %v", spell.ActionID, metrics.Blocks, metrics.Crits, metrics.TotalDamage)
				}
				spell.CalcDamage(sim, r.CurrentTarget, 100, spell.OutcomeMeleeSpecialBlockAndCritNoHitCounter)
				if metrics != spell.SpellMetrics[r.CurrentTarget.UnitIndex] {
					t.Fatal("uncounted outcome changed hand metrics")
				}
			}
		})
	}
}

func TestRogueMutilateColdBloodBothHandsOneCast(t *testing.T) {
	sim, r := rogueMutilateSim(t, map[string]int{"mutilate": 1, "coldBlood": 1}, proto.Ruleset_RulesetForever)
	if !r.ColdBlood.Cast(sim, r.CurrentTarget) || !r.Mutilate.Cast(sim, r.CurrentTarget) {
		t.Fatal("Cold Blood/Mutilate was not cast")
	}
	for _, mask := range []core.ProcMask{core.ProcMaskMeleeMHSpecial, core.ProcMaskMeleeOHSpecial} {
		spell := mutilateHand(t, r, mask)
		if spell.Flags.Matches(core.SpellFlagAPL) || !spell.Flags.Matches(core.SpellFlagNoOnCastComplete) {
			t.Errorf("%s is not a triggered hand strike", spell.ActionID)
		}
		if got := spell.SpellMetrics[r.CurrentTarget.UnitIndex].Crits; got != 1 {
			t.Errorf("%s crits %d, want 1", spell.ActionID, got)
		}
	}
	parentMetrics := r.Mutilate.SpellMetrics[r.CurrentTarget.UnitIndex]
	if r.Mutilate.ActionID.Tag != 0 || !r.Mutilate.Flags.Matches(core.SpellFlagAPL) || r.Mutilate.ProcMask != core.ProcMaskEmpty ||
		parentMetrics.Casts != 1 || parentMetrics.Hits != 1 || parentMetrics.TotalDamage != 0 {
		t.Fatal("Mutilate parent is not a selectable, non-damaging cast gate")
	}
	if r.GetAura("Cold Blood").IsActive() || r.ComboPoints() != 2 || r.CurrentEnergy() != 40 {
		t.Fatalf("landed cast: Cold Blood %v, combo points %d, energy %v; want false, 2, 40", r.GetAura("Cold Blood").IsActive(), r.ComboPoints(), r.CurrentEnergy())
	}
	sim.CurrentTime = time.Second
	r.AddEnergy(sim, 60, r.Mutilate.EnergyMetrics())
	if !r.Mutilate.Cast(sim, r.CurrentTarget) {
		t.Fatal("second Mutilate was not cast")
	}
	for _, mask := range []core.ProcMask{core.ProcMaskMeleeMHSpecial, core.ProcMaskMeleeOHSpecial} {
		spell := mutilateHand(t, r, mask)
		if got := spell.SpellMetrics[r.CurrentTarget.UnitIndex].Crits; got != 1 {
			t.Errorf("%s retained Cold Blood for the next cast: %d crits", spell.ActionID, got)
		}
	}
}

func TestRogueSealFateOneRollPerMutilate(t *testing.T) {
	for _, rank := range []int{1, 3, 5} {
		for _, crits := range [][2]bool{{true, false}, {false, true}, {true, true}, {false, false}} {
			t.Run(fmt.Sprintf("rank%d/%v", rank, crits), func(t *testing.T) {
				sim, r := rogueMutilateSim(t, map[string]int{"mutilate": 1, "sealFate": rank}, proto.Ruleset_RulesetForever)
				oracle, _ := rogueMutilateSim(t, map[string]int{"mutilate": 1}, proto.Ruleset_RulesetForever)
				trigger := r.GetAura("Seal Fate")
				hands := []*core.Spell{mutilateHand(t, r, core.ProcMaskMeleeMHSpecial), mutilateHand(t, r, core.ProcMaskMeleeOHSpecial)}
				for cast := 0; cast < 24; cast++ {
					sim.CurrentTime = time.Duration(cast) * time.Second
					r.SpendComboPoints(sim, r.Eviscerate)
					want := int32(0)
					if (crits[0] || crits[1]) && oracle.Proc(.2*float64(rank), "Seal Fate") {
						want = 1
					}
					for i, hand := range hands {
						outcome := core.OutcomeHit
						if crits[i] {
							outcome = core.OutcomeCrit
						}
						trigger.OnSpellHitDealt(trigger, sim, hand, &core.SpellResult{Target: r.CurrentTarget, Outcome: outcome})
					}
					if got := r.ComboPoints(); got != want {
						t.Fatalf("cast %d: Seal Fate awarded %d combo points, want one-roll result %d", cast, got, want)
					}
				}
			})
		}
	}
}

func TestRogueColdBloodAndSealFateOrdinaryBuilder(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		t.Run(ruleset.String(), func(t *testing.T) {
			sim, r := rogueMutilateSim(t, map[string]int{"mutilate": 1, "coldBlood": 1, "sealFate": 5}, ruleset)
			r.ColdBlood.Cast(sim, r.CurrentTarget)
			if !r.SinisterStrike.Cast(sim, r.CurrentTarget) {
				t.Fatal("Sinister Strike was not cast")
			}
			if r.GetAura("Cold Blood").IsActive() || r.ComboPoints() != 2 || r.SinisterStrike.SpellMetrics[r.CurrentTarget.UnitIndex].Crits != 1 {
				t.Fatal("ordinary builder did not consume Cold Blood, crit, and grant its Seal Fate point")
			}
			sim, r = rogueMutilateSim(t, map[string]int{"mutilate": 1, "sealFate": 1}, ruleset)
			oracle, _ := rogueMutilateSim(t, map[string]int{"mutilate": 1}, ruleset)
			trigger := r.GetAura("Seal Fate")
			for cast := 0; cast < 24; cast++ {
				sim.CurrentTime = time.Duration(cast) * time.Second
				r.SpendComboPoints(sim, r.Eviscerate)
				want := int32(0)
				if oracle.Proc(.2, "Seal Fate") {
					want = 1
				}
				trigger.OnSpellHitDealt(trigger, sim, r.SinisterStrike, &core.SpellResult{Target: r.CurrentTarget, Outcome: core.OutcomeCrit})
				if r.ComboPoints() != want {
					t.Fatal("ordinary builder changed its single Seal Fate roll")
				}
			}
		})
	}
}
