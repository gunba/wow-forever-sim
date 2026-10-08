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
	"github.com/wowsims/classic/sim/warlock"
	googleProto "google.golang.org/protobuf/proto"
)

func havocFixture() (*core.Simulation, []*warlock.Warlock) {
	req := warlockEffectFixture(map[string]int{"baneOfHavoc": 1})
	first := req.Raid.Parties[0].Players[0]
	first.Name = "First Havoc"
	first.GetWarlock().Options.Summon = proto.WarlockOptions_Imp
	second := googleProto.Clone(first).(*proto.Player)
	second.Name = "Second Havoc"
	req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, second)
	req.Encounter.Targets = append(req.Encounter.Targets, googleProto.Clone(req.Encounter.Targets[0]).(*proto.Target))
	req.Encounter.Duration = 600
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	owners := make([]*warlock.Warlock, 2)
	for i, player := range sim.Raid.Parties[0].Players {
		owners[i] = player.(warlock.WarlockAgent).GetWarlock()
	}
	return sim, owners
}

func applyHavocBane(sim *core.Simulation, spell *core.Spell, target *core.Unit) {
	spell.BonusHitRating = 10000
	spell.ApplyEffects(sim, target, spell)
}

func havocDamage(owner *warlock.Warlock, target *core.Unit) float64 {
	copy := owner.GetSpell(core.ActionID{SpellID: 1243339})
	if copy == nil {
		return 0
	}
	return copy.SpellMetrics[target.UnitIndex].TotalDamage
}

// Deliver a result which has already passed the original spell's modifiers,
// critical outcome and mitigation, just as the damage callbacks receive it.
func deliverHavocSource(sim *core.Simulation, spell *core.Spell, target *core.Unit, damage float64, periodic bool) {
	result := spell.NewResult(target)
	result.Outcome = core.OutcomeCrit
	result.Damage = damage
	result.PreOutcomeDamage = damage * 4
	result.ResistanceMultiplier = .25
	result.Threat = spell.ThreatFromDamage(result.Outcome, damage)
	if periodic {
		spell.DealPeriodicDamage(sim, result)
	} else {
		spell.DealDamage(sim, result)
	}
}

func TestBaneOfHavocCopiesDealtDamage(t *testing.T) {
	sim, owners := havocFixture()
	owner := owners[0]
	source, marked := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
	applyHavocBane(sim, owner.BaneOfHavoc, marked)

	owner.PseudoStats.DamageDealtMultiplier *= 2
	owner.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexShadow] *= 3
	marked.PseudoStats.DamageTakenMultiplier *= 5
	marked.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexShadow] *= 7
	marked.AddStatsDynamic(sim, stats.Stats{stats.ShadowResistance: 1000})
	marked.DynamicDamageTakenModifiers = append(marked.DynamicDamageTakenModifiers,
		func(_ *core.Simulation, _ *core.Spell, result *core.SpellResult) {
			result.Damage = result.Damage*11 + 13
		})

	shadowBolt := owner.ShadowBolt[len(owner.ShadowBolt)-1]
	deliverHavocSource(sim, shadowBolt, source, 120, false)
	if got := havocDamage(owner, marked); math.Abs(got-18) > 1e-9 {
		t.Fatalf("Havoc copied %g damage, want 18 from the 120 actually dealt", got)
	}
	copy := owner.GetSpell(core.ActionID{SpellID: 1243339})
	if copy.SpellSchool != core.SpellSchoolShadow || copy.ProcMask != core.ProcMaskEmpty ||
		!copy.Flags.Matches(core.SpellFlagNoOnDamageDealt) {
		t.Fatal("Havoc copy must be the passive Shadow child without damage proc callbacks")
	}
	metrics := &copy.SpellMetrics[marked.UnitIndex]
	if metrics.Crits != 0 || metrics.TotalResistedDamage != 0 || metrics.TotalCritDamage != 0 {
		t.Fatal("the copied amount was crit or resisted again")
	}
	if want := 18 * owner.PseudoStats.ThreatMultiplier; math.Abs(metrics.TotalThreat-want) > 1e-9 {
		t.Fatalf("copy threat %g, want ordinary damage threat %g", metrics.TotalThreat, want)
	}

	corruption := owner.Corruption[len(owner.Corruption)-1]
	deliverHavocSource(sim, corruption, source, 90, true)
	if got := havocDamage(owner, marked); math.Abs(got-31.5) > 1e-9 {
		t.Fatalf("periodic damage brought Havoc to %g, want 31.5", got)
	}
	// Driver 1243338 includes outgoing weapon damage and may proc from procs;
	// it does not restrict the source to the Warlock class-family flags.
	weaponAttack := owner.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{SpellID: 6603}, SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee, ProcMask: core.ProcMaskMeleeMHAuto,
		DamageMultiplier: 1, ThreatMultiplier: 1,
	})
	deliverHavocSource(sim, weaponAttack, source, 30, false)
	ownedProc := owner.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{SpellID: 1, Tag: 77}, SpellSchool: core.SpellSchoolFire,
		ProcMask: core.ProcMaskSpellDamage, Flags: core.SpellFlagPassiveSpell,
		DamageMultiplier: 1, ThreatMultiplier: 1,
	})
	deliverHavocSource(sim, ownedProc, source, 20, false)
	if got := havocDamage(owner, marked); math.Abs(got-39) > 1e-9 {
		t.Fatalf("owned weapon/proc damage brought Havoc to %g, want 39", got)
	}
}

func TestBaneOfHavocExcludesOtherDamageAndProcChains(t *testing.T) {
	sim, owners := havocFixture()
	owner, other := owners[0], owners[1]
	first, second := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
	applyHavocBane(sim, owner.BaneOfHavoc, second)
	applyHavocBane(sim, other.BaneOfHavoc, first)
	bolt := owner.ShadowBolt[len(owner.ShadowBolt)-1]

	deliverHavocSource(sim, bolt, second, 100, false) // already marked
	deliverHavocSource(sim, bolt, &owner.Unit, 100, false)
	deliverHavocSource(sim, bolt, &other.Unit, 100, false)
	deliverHavocSource(sim, owner.Imp.GetSpell(core.ActionID{SpellID: 11763}), first, 100, false)
	bolt.CalcAndDealHealing(sim, &owner.Unit, 100, bolt.OutcomeHealing)
	result := bolt.NewResult(first)
	result.Outcome = core.OutcomeMiss
	bolt.DealDamage(sim, result)
	if havocDamage(owner, second) != 0 || havocDamage(other, first) != 0 {
		t.Fatal("marked, self, friendly, pet, healing or missed damage triggered Havoc")
	}

	copyCallbacks := 0
	driver := owner.GetAura("Bane of Havoc - Copy")
	if driver == nil {
		t.Fatal("missing Havoc damage driver")
	}
	originalCallback := driver.OnSpellHitDealt
	driver.OnSpellHitDealt = func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		if spell.ActionID.SpellID == 1243339 {
			copyCallbacks++
		}
		originalCallback(aura, sim, spell, result)
	}
	deliverHavocSource(sim, bolt, first, 100, false)
	deliverHavocSource(sim, other.ShadowBolt[len(other.ShadowBolt)-1], second, 80, false)
	if havocDamage(owner, second) != 15 || havocDamage(other, first) != 12 || copyCallbacks != 0 {
		t.Fatal("Havoc duplicated damage, crossed caster ownership, or triggered a new damage proc")
	}
}

func TestBaneOfHavocTargetAndBaneExclusivity(t *testing.T) {
	sim, owners := havocFixture()
	owner, other := owners[0], owners[1]
	first, second := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
	agony := owner.BaneOfAgony[len(owner.BaneOfAgony)-1]
	applyHavocBane(sim, agony, first)
	applyHavocBane(sim, owner.BaneOfDoom, second)
	applyHavocBane(sim, owner.BaneOfHavoc, first)
	if agony.Dot(first).IsActive() || !owner.BaneOfDoom.Dot(second).IsActive() {
		t.Fatal("Havoc did not replace only the Bane on its target")
	}
	applyHavocBane(sim, owner.BaneOfHavoc, second)
	if owner.BaneOfHavocAuras.Get(first).IsActive() || owner.BaneOfDoom.Dot(second).IsActive() ||
		!owner.BaneOfHavocAuras.Get(second).IsActive() {
		t.Fatal("moving Havoc did not clear its old target and replace the new target's Bane")
	}

	applyHavocBane(sim, owner.BaneOfDoom, second)
	applyHavocBane(sim, other.BaneOfHavoc, second)
	applyHavocBane(sim, owner.BaneOfHavoc, first)
	if !owner.BaneOfDoom.Dot(second).IsActive() || !other.BaneOfHavocAuras.Get(second).IsActive() {
		t.Fatal("a stale Havoc target canceled its replacement Bane or another caster's Havoc")
	}
	applyHavocBane(sim, owner.BaneOfHavoc, second)
	applyHavocBane(sim, owner.CurseOfElements, second)
	if owner.BaneOfDoom.Dot(second).IsActive() || !other.BaneOfHavocAuras.Get(second).IsActive() ||
		!owner.BaneOfHavocAuras.Get(second).IsActive() {
		t.Fatal("Havoc or a utility Curse crossed caster/Bane exclusivity")
	}
	applyHavocBane(sim, other.BaneOfHavoc, first)
	if !owner.BaneOfHavocAuras.Get(second).IsActive() || other.BaneOfHavocAuras.Get(second).IsActive() {
		t.Fatal("moving the second caster's Havoc altered the first caster's target")
	}
	applyHavocBane(sim, agony, second)
	applyHavocBane(sim, owner.BaneOfHavoc, first)
	if !agony.Dot(second).IsActive() || !other.BaneOfHavocAuras.Get(first).IsActive() {
		t.Fatal("moving Havoc after Agony replacement canceled unrelated active Banes")
	}
}

func TestBaneOfHavocCancellationAndExpiration(t *testing.T) {
	for _, expire := range []bool{false, true} {
		name := "cancel"
		if expire {
			name = "expire"
		}
		t.Run(name, func(t *testing.T) {
			sim, owners := havocFixture()
			owner := owners[0]
			first, second := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
			applyHavocBane(sim, owner.BaneOfHavoc, second)
			aura := owner.BaneOfHavocAuras.Get(second)
			if aura.Duration != 5*time.Minute {
				t.Fatal("Havoc must last five minutes")
			}
			if expire {
				for _, unit := range sim.Environment.AllUnits {
					unit.AutoAttacks.CancelAutoSwing(sim)
				}
				for _, w := range owners {
					w.ActivePet.Disable(sim, false)
				}
				core.StartDelayedAction(sim, core.DelayedActionOptions{DoAt: aura.ExpiresAt(), OnAction: func(*core.Simulation) {}})
				for sim.CurrentTime < aura.ExpiresAt() {
					if sim.Step() {
						t.Fatal("fight ended before Havoc expired")
					}
				}
			} else {
				aura.Deactivate(sim)
			}
			if aura.IsActive() {
				t.Fatal("Havoc remained active after cancellation/expiration")
			}
			deliverHavocSource(sim, owner.ShadowBolt[len(owner.ShadowBolt)-1], first, 100, false)
			if havocDamage(owner, second) != 0 {
				t.Fatal("Havoc copied damage after cancellation/expiration")
			}
		})
	}
}
