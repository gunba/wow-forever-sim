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
	"github.com/wowsims/classic/sim/mage"
)

// Client 1.60.1.70245 SpellAuraOptions/SpellEffect and TraitDefinitionEffectPoints.
// Public context: https://github.com/ElliotWood/Forever/pull/660 and /pull/615.
// These guard the exported client rules, not a claim about server deployment.
func TestOctober8MageMissileBarrageImpact(t *testing.T) {
	sim := core.NewSim(mageSpellFixture(map[string]int{"missileBarrage": 1, "arcaneBlast": 1}), simsignals.Signals{})
	unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
	trigger, barrage := unit.GetAura("Missile Barrage Trigger"), unit.GetAura("Missile Barrage")
	frostfire := unit.GetSpell(core.ActionID{SpellID: 1237313})

	t.Run("not_before_impact", func(t *testing.T) {
		barrage.Deactivate(sim)
		for i := 0; i < 100; i++ {
			if trigger.OnCastComplete != nil {
				trigger.OnCastComplete(trigger, sim, frostfire)
			}
		}
		if barrage.IsActive() {
			t.Fatal("Missile Barrage procced at cast completion before projectile impact")
		}
	})
	t.Run("miss", func(t *testing.T) {
		if trigger.OnSpellHitDealt == nil {
			t.Fatal("Missile Barrage has no harmful spell-hit callback")
		}
		barrage.Deactivate(sim)
		for i := 0; i < 100; i++ {
			trigger.OnSpellHitDealt(trigger, sim, frostfire, &core.SpellResult{Target: target, Outcome: core.OutcomeMiss})
		}
		if barrage.IsActive() {
			t.Fatal("missed Frostfire Bolt triggered Missile Barrage")
		}
	})
	for _, id := range []int32{133, 116, 1237313, 30451} {
		for _, outcome := range []core.HitOutcome{core.OutcomeHit, core.OutcomeCrit} {
			t.Run(fmt.Sprintf("landed/%d/%v", id, outcome), func(t *testing.T) {
				if trigger.OnSpellHitDealt == nil {
					t.Fatal("missing Missile Barrage hit callback")
				}
				barrage.Deactivate(sim)
				spell := unit.GetSpell(core.ActionID{SpellID: id})
				spell.CurCast.Cost = 0 // A free, landed nuke is still eligible.
				for i := 0; i < 100 && !barrage.IsActive(); i++ {
					trigger.OnSpellHitDealt(trigger, sim, spell, &core.SpellResult{Target: target, Outcome: outcome})
				}
				if !barrage.IsActive() {
					t.Fatal("eligible landed nuke did not trigger Missile Barrage")
				}
			})
		}
	}
	t.Run("ineligible_spell", func(t *testing.T) {
		if trigger.OnSpellHitDealt == nil {
			t.Fatal("missing Missile Barrage hit callback")
		}
		barrage.Deactivate(sim)
		arcaneMissiles := unit.GetSpell(core.ActionID{SpellID: 5143})
		for i := 0; i < 100; i++ {
			trigger.OnSpellHitDealt(trigger, sim, arcaneMissiles, &core.SpellResult{Target: target, Outcome: core.OutcomeCrit})
		}
		if barrage.IsActive() {
			t.Fatal("Arcane Missiles triggered Missile Barrage")
		}
	})
}

func TestOctober8MageMissileBarrageClassicCallback(t *testing.T) {
	req := mageSpellFixture(map[string]int{"missileBarrage": 1})
	req.SimOptions.Ruleset = proto.Ruleset_RulesetClassic
	sim := core.NewSim(req, simsignals.Signals{})
	trigger := sim.Raid.AllPlayerUnits[0].GetAura("Missile Barrage Trigger")
	if trigger.OnCastComplete == nil || trigger.OnSpellHitDealt != nil {
		t.Fatal("Forever proc timing changed the retained Classic callback")
	}
}

func TestOctober8MageMasterOfElementsBaseCost(t *testing.T) {
	for rank := 1; rank <= 3; rank++ {
		for _, paidFraction := range []float64{0, .5, 1.3} {
			t.Run(fmt.Sprintf("rank%d/paid%g", rank, paidFraction), func(t *testing.T) {
				sim := core.NewSim(mageSpellFixture(map[string]int{"masterOfElements": rank}), simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
				unit.SpendMana(sim, unit.CurrentMana(), unit.NewManaMetrics(core.ActionID{SpellID: 1}))
				spell := unit.GetSpell(core.ActionID{SpellID: 133})
				spell.CurCast.Cost = spell.Cost.BaseCost * paidFraction
				aura := unit.GetAura("Master of Elements")
				aura.OnSpellHitDealt(aura, sim, spell, &core.SpellResult{Target: target, Outcome: core.OutcomeCrit})
				want := spell.Cost.BaseCost * .1 * float64(rank)
				if got := unit.CurrentMana(); math.Abs(got-want) > 1e-8 {
					t.Fatalf("refund %g, want base-cost refund %g", got, want)
				}
			})
		}
	}
}

func TestOctober8MageMasterOfElementsBatch(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		t.Run(ruleset.String(), func(t *testing.T) {
			req := mageSpellFixture(map[string]int{"masterOfElements": 3, "ignite": 1})
			req.SimOptions.Ruleset = ruleset
			for i := 0; i < 2; i++ {
				target := *req.Encounter.Targets[0]
				req.Encounter.Targets = append(req.Encounter.Targets, &target)
			}
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			unit := sim.Raid.AllPlayerUnits[0]
			unit.SpendMana(sim, unit.CurrentMana(), unit.NewManaMetrics(core.ActionID{SpellID: 1}))
			spell := unit.GetSpell(core.ActionID{SpellID: 120}) // Cone of Cold, three target hits in one batch.
			if ruleset == proto.Ruleset_RulesetClassic {
				// The Forever-only AoE registration is absent in Classic;
				// retain its existing independent paid-crit callback behavior.
				spell = unit.GetSpell(core.ActionID{SpellID: 133})
			}
			spell.CurCast.Cost = spell.Cost.BaseCost
			aura := unit.GetAura("Master of Elements")
			refund := spell.Cost.BaseCost * .3
			for _, target := range sim.Encounter.TargetUnits {
				aura.OnSpellHitDealt(aura, sim, spell, &core.SpellResult{Target: target, Outcome: core.OutcomeCrit})
			}
			want := refund
			if ruleset == proto.Ruleset_RulesetClassic {
				want *= 3
			}
			if got := unit.CurrentMana(); math.Abs(got-want) > 1e-8 {
				t.Errorf("batch refund %g, want %g", got, want)
			}
			if ruleset == proto.Ruleset_RulesetClassic {
				spell.CurCast.Cost = 0
				aura.OnSpellHitDealt(aura, sim, spell, &core.SpellResult{Target: sim.Encounter.TargetUnits[0], Outcome: core.OutcomeCrit})
				if math.Abs(unit.CurrentMana()-want) > 1e-8 {
					t.Fatal("Classic free cast unexpectedly gained a refund")
				}
				return
			}
			if aura.Icd == nil || aura.Icd.Duration != 9*time.Millisecond {
				t.Error("Master of Elements ICD must be 9 milliseconds, not seconds")
			}
			sim.CurrentTime = 8 * time.Millisecond
			aura.OnSpellHitDealt(aura, sim, spell, &core.SpellResult{Target: sim.Encounter.TargetUnits[0], Outcome: core.OutcomeCrit})
			if math.Abs(unit.CurrentMana()-want) > 1e-8 {
				t.Error("ICD did not suppress a crit at 8ms")
			}
			sim.CurrentTime = 9 * time.Millisecond
			for _, outcome := range []core.HitOutcome{core.OutcomeMiss, core.OutcomeHit} {
				aura.OnSpellHitDealt(aura, sim, spell, &core.SpellResult{Target: sim.Encounter.TargetUnits[0], Outcome: outcome})
			}
			// A costless triggered fire effect cannot refund or consume the ICD.
			ignite := unit.GetSpell(core.ActionID{SpellID: 12654})
			if ignite != nil {
				aura.OnSpellHitDealt(aura, sim, ignite, &core.SpellResult{Target: sim.Encounter.TargetUnits[0], Outcome: core.OutcomeCrit})
			}
			spell.CurCast.Cost = 0
			aura.OnSpellHitDealt(aura, sim, spell, &core.SpellResult{Target: sim.Encounter.TargetUnits[0], Outcome: core.OutcomeCrit})
			if got := unit.CurrentMana(); math.Abs(got-2*refund) > 1e-8 {
				t.Errorf("free crit at 9ms refunded %g total, want %g", got, 2*refund)
			}
		})
	}
}

func october8ConcentrationFixture(targets int, ruleset proto.Ruleset) *core.Simulation {
	req := mageSpellFixture(map[string]int{"arcaneConcentration": 5, "arcaneBlast": 1, "wintersChill": 5, "fingersOfFrost": 1, "improvedBlizzard": 3})
	req.SimOptions.Ruleset = ruleset
	for len(req.Encounter.Targets) < targets {
		target := *req.Encounter.Targets[0]
		req.Encounter.Targets = append(req.Encounter.Targets, &target)
	}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim
}

// 70245 Blizzard has a magic-defense, area-targeted E_DUMMY on the paid
// parent; its area-trigger children lack NOT_A_PROC. Public context: PR643/647.
func TestOctober8MageBlizzardInitialConcentration(t *testing.T) {
	const trials = 3000
	for _, targets := range []int{1, 3} {
		t.Run(fmt.Sprintf("targets%d", targets), func(t *testing.T) {
			sim := october8ConcentrationFixture(targets, proto.Ruleset_RulesetForever)
			ref := october8ConcentrationFixture(targets, proto.Ruleset_RulesetForever)
			unit, refUnit := sim.Raid.AllPlayerUnits[0], ref.Raid.AllPlayerUnits[0]
			spell, refSpell := unit.GetSpell(core.ActionID{SpellID: 10187}), refUnit.GetSpell(core.ActionID{SpellID: 10187})
			clear, refClear := unit.GetAura("Clearcasting"), refUnit.GetAura("Clearcasting")
			refTrigger := refUnit.GetAura("Arcane Concentration")
			genericHits := 0
			observer := unit.GetAura("Fingers of Frost Trigger")
			originalHit := observer.OnSpellHitDealt
			observer.OnSpellHitDealt = func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				genericHits++
				originalHit(aura, sim, spell, result)
			}
			procs, mismatches := 0, 0
			for i := 0; i < trials; i++ {
				sim.CurrentTime, ref.CurrentTime = time.Duration(i)*2*time.Second, time.Duration(i)*2*time.Second
				clear.Deactivate(sim)
				refClear.Deactivate(ref)
				spell.ApplyEffects(sim, sim.Encounter.TargetUnits[0], spell)
				for _, target := range ref.Encounter.TargetUnits {
					result := refSpell.CalcOutcome(ref, target, refSpell.OutcomeMagicHitNoHitCounter)
					refTrigger.OnSpellHitDealt(refTrigger, ref, refSpell, result)
					refSpell.DisposeResult(result)
				}
				if clear.IsActive() {
					procs++
				}
				if clear.IsActive() != refClear.IsActive() {
					mismatches++
				}
			}
			if mismatches != 0 {
				t.Errorf("Blizzard initial AC disagreed with per-enemy landed-hit/ICD controls in %d trials", mismatches)
			}
			hitChance := 1 - spell.SpellChanceToMiss(unit.AttackTables[sim.Encounter.TargetUnits[0].UnitIndex][spell.CastType])
			want := 1 - math.Pow(1-.1*hitChance, float64(targets))
			if got := float64(procs) / trials; math.Abs(got-want) > .03 {
				t.Errorf("initial proc chance %g, want %g with %d targets", got, want, targets)
			}
			if genericHits != 0 {
				t.Error("initial Blizzard application emitted generic hit/proc callbacks")
			}
			for _, metrics := range spell.SpellMetrics {
				if metrics.Hits != 0 || metrics.Misses != 0 || metrics.TotalDamage != 0 || metrics.TotalThreat != 0 {
					t.Error("initial AC roll invented a counted hit, damage or threat")
				}
			}
		})
	}
}

func TestOctober8MageBlizzardConcentrationControls(t *testing.T) {
	t.Run("miss_does_not_gate_channel", func(t *testing.T) {
		sim := october8ConcentrationFixture(3, proto.Ruleset_RulesetForever)
		unit := sim.Raid.AllPlayerUnits[0]
		spell := unit.GetSpell(core.ActionID{SpellID: 10187})
		spell.BonusHitRating = -1000 * core.SpellHitRatingPerHitChance
		unit.AddStatDynamic(sim, stats.SpellCrit, -1000*core.SpellCritRatingPerCritChance)
		manaBefore := unit.CurrentMana()
		if !spell.Cast(sim, unit.CurrentTarget) {
			t.Fatal("Blizzard was not castable")
		}
		if unit.GetAura("Clearcasting").IsActive() {
			t.Error("missed initial application procced Clearcasting")
		}
		if math.Abs(manaBefore-unit.CurrentMana()-spell.CurCast.Cost) > 1e-8 {
			t.Error("initial application altered the channel's mana charge")
		}
		dot := spell.AOEDot()
		if !dot.IsActive() || dot.Duration != 8*time.Second || unit.GCD.ReadyAt() != core.GCDDefault {
			t.Error("initial application changed channel duration or GCD")
		}
		for sim.CurrentTime < 8*time.Second {
			if sim.Step() {
				t.Fatal("fight ended before Blizzard channel completed")
			}
		}
		if dot.TickCount != 8 || spell.SpellMetrics[unit.CurrentTarget.UnitIndex].Ticks != 8 {
			t.Error("initial application miss blocked or altered channel ticks")
		}
		if spell.SpellMetrics[unit.CurrentTarget.UnitIndex].TotalDamage <= 0 {
			t.Error("Blizzard's existing periodic damage was blocked")
		}
		if unit.GetAura("Clearcasting").IsActive() || unit.GetAura("Fingers of Frost").IsActive() || unit.GetAura("Winter's Chill").GetStacks() != 0 {
			t.Error("Blizzard ticks invented an AC, FoF or Winter's Chill proc")
		}
	})
	t.Run("shared_icd", func(t *testing.T) {
		sim := october8ConcentrationFixture(3, proto.Ruleset_RulesetForever)
		unit := sim.Raid.AllPlayerUnits[0]
		trigger := unit.GetAura("Arcane Concentration")
		blizzard := unit.GetSpell(core.ActionID{SpellID: 10187})
		trigger.Icd.Use(sim)
		for i := 0; i < 100; i++ {
			blizzard.ApplyEffects(sim, unit.CurrentTarget, blizzard)
		}
		if unit.GetAura("Clearcasting").IsActive() {
			t.Error("Blizzard ignored Arcane Concentration's existing one-second ICD")
		}
	})
	t.Run("classic", func(t *testing.T) {
		sim := october8ConcentrationFixture(3, proto.Ruleset_RulesetClassic)
		unit := sim.Raid.AllPlayerUnits[0]
		blizzard := unit.GetSpell(core.ActionID{SpellID: 10187})
		for i := 0; i < 100; i++ {
			sim.CurrentTime = time.Duration(i) * 2 * time.Second
			blizzard.ApplyEffects(sim, unit.CurrentTarget, blizzard)
		}
		if unit.GetAura("Clearcasting").IsActive() {
			t.Error("Forever initial Blizzard application changed Classic")
		}
	})
}

func TestOctober8MageConcentrationActiveSpellControls(t *testing.T) {
	for _, id := range []int32{133, 116, 30451} {
		t.Run(fmt.Sprint(id), func(t *testing.T) {
			sim := october8ConcentrationFixture(1, proto.Ruleset_RulesetForever)
			unit := sim.Raid.AllPlayerUnits[0]
			trigger, clear := unit.GetAura("Arcane Concentration"), unit.GetAura("Clearcasting")
			spell := unit.GetSpell(core.ActionID{SpellID: id})
			spell.CurCast.Cost = 0
			for i := 0; i < 100; i++ {
				sim.CurrentTime = time.Duration(i) * 2 * time.Second
				trigger.OnSpellHitDealt(trigger, sim, spell, &core.SpellResult{Target: unit.CurrentTarget, Outcome: core.OutcomeMiss})
			}
			if clear.IsActive() {
				t.Fatal("missed active spell procced Clearcasting")
			}
			for i := 0; i < 100 && !clear.IsActive(); i++ {
				sim.CurrentTime += 2 * time.Second
				trigger.OnSpellHitDealt(trigger, sim, spell, &core.SpellResult{Target: unit.CurrentTarget, Outcome: core.OutcomeCrit})
			}
			if !clear.IsActive() {
				t.Fatal("landed active spell could not proc Clearcasting")
			}
			clear.Deactivate(sim)
			spell.Flags &^= mage.SpellFlagMage // Ordinary item damage is not a Mage cast.
			for i := 0; i < 100; i++ {
				sim.CurrentTime += 2 * time.Second
				trigger.OnSpellHitDealt(trigger, sim, spell, &core.SpellResult{Target: unit.CurrentTarget, Outcome: core.OutcomeHit})
			}
			if clear.IsActive() {
				t.Error("non-Mage proc damage entered the Mage talent callback")
			}
		})
	}
}

func TestOctober8MageFrostfirePeriodicTalentScopes(t *testing.T) {
	cases := []struct {
		name           string
		talents        map[string]int
		direct, period float64
		power          bool
	}{
		{"arcane_power", map[string]int{"arcanePower": 1}, 1.3, 1, true},
		{"instability", map[string]int{"arcaneInstability": 3}, 1.03, 1, false},
		{"piercing1", map[string]int{"piercingIce": 1}, 1.02, 1.02, false},
		{"piercing2", map[string]int{"piercingIce": 2}, 1.04, 1.02, false},
		{"piercing3", map[string]int{"piercingIce": 3}, 1.06, 1.02, false},
		{"combined", map[string]int{"arcanePower": 1, "arcaneInstability": 3, "piercingIce": 3, "firePower": 5}, 1.49, 1.12, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sim := core.NewSim(mageSpellFixture(tc.talents), simsignals.Signals{})
			unit, target := sim.Raid.AllPlayerUnits[0], sim.Encounter.TargetUnits[0]
			spell := unit.GetSpell(core.ActionID{SpellID: 1237313})
			unit.AddStatDynamic(sim, stats.SpellCrit, -1000*core.SpellCritRatingPerCritChance)
			power := unit.GetAura("Arcane Power")
			if tc.power {
				power.Activate(sim)
			}
			if math.Abs(spell.DamageMultiplierAdditive-tc.direct) > 1e-8 {
				t.Fatalf("direct multiplier %g, want %g", spell.DamageMultiplierAdditive, tc.direct)
			}
			dot := spell.Dot(target)
			dot.OnSnapshot(sim, target, dot, false)
			base := spell.AttackerDamageMultiplier(unit.AttackTables[target.UnitIndex][spell.CastType], false) / tc.direct
			if got := dot.SnapshotAttackerMultiplier / base; math.Abs(got-tc.period) > 1e-8 {
				t.Errorf("periodic snapshot multiplier %g, want %g", got, tc.period)
			}
			before := spell.SpellMetrics[target.UnitIndex].TotalDamage
			dot.OnTick(sim, target, dot)
			if got := spell.SpellMetrics[target.UnitIndex].TotalDamage - before; math.Abs(got-19*base*tc.period) > 1e-8 {
				t.Errorf("periodic tick damage %g, want %g", got, 19*base*tc.period)
			}
			if tc.power {
				power.Deactivate(sim)
				before = spell.SpellMetrics[target.UnitIndex].TotalDamage
				dot.OnTick(sim, target, dot)
				if got := spell.SpellMetrics[target.UnitIndex].TotalDamage - before; math.Abs(got-19*base*tc.period) > 1e-8 {
					t.Errorf("Arcane Power expiry changed excluded periodic damage to %g", got)
				}
			}
		})
	}
}
