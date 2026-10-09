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
	"github.com/wowsims/classic/sim/druid"
	"github.com/wowsims/classic/sim/hunter"
)

func october9DruidFixture(talents map[string]int, helm bool, ruleset proto.Ruleset, bear ...bool) (*core.Simulation, *druid.Druid) {
	req := historyTalentFixture("feral", talents)
	p := req.Raid.Parties[0].Players[0]
	p.ForeverTier1Bonuses = false
	if len(bear) > 0 && bear[0] {
		p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{Options: &proto.FeralTankDruid_Options{}}}
	}
	p.Equipment = &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 17)}
	for i := range p.Equipment.Items {
		p.Equipment.Items[i] = &proto.ItemSpec{}
	}
	if helm {
		p.Equipment.Items[0].Id = druid.WolfsheadHelm
	}
	p.Buffs, p.Consumes = &proto.IndividualBuffs{}, &proto.Consumes{}
	req.Raid.Buffs, req.Raid.Parties[0].Buffs = &proto.RaidBuffs{}, &proto.PartyBuffs{}
	if len(bear) > 1 && bear[1] {
		req.Raid.Parties[0].Buffs.FlametongueTotem = true
		p.Equipment.Items[14].Id = 7946
	}
	req.SimOptions.Ruleset = ruleset
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	d := sim.Raid.Parties[0].Players[0].(interface{ GetDruid() *druid.Druid }).GetDruid()
	return sim, d
}

func TestOctober9BearThreatRuleset(t *testing.T) {
	for _, tc := range []struct {
		ruleset proto.Ruleset
		want    float64
	}{{proto.Ruleset_RulesetForever, 1.5}, {proto.Ruleset_RulesetClassic, 1.3}} {
		sim, d := october9DruidFixture(nil, false, tc.ruleset, true)
		d.CancelShapeshift(sim)
		baseline := d.PseudoStats.ThreatMultiplier
		d.BearFormAura.Activate(sim)
		if got := d.PseudoStats.ThreatMultiplier / baseline; math.Abs(got-tc.want) > 1e-9 {
			t.Errorf("ruleset %v: Bear threat multiplier %g, want %g", tc.ruleset, got, tc.want)
		}
		d.CancelShapeshift(sim)
		if math.Abs(d.PseudoStats.ThreatMultiplier-baseline) > 1e-9 {
			t.Fatal("Bear threat leaked after form exit")
		}
	}
}

func TestOctober9ShiftingPowerRejectsFullEnergy(t *testing.T) {
	sim, d := october9DruidFixture(map[string]int{"shiftingPower": 1}, false, proto.Ruleset_RulesetForever)
	d.AddEnergy(sim, d.MaxEnergy()-d.CurrentEnergy(), d.NewEnergyMetrics(core.ActionID{SpellID: 1322605}))
	mana, gcd := d.CurrentMana(), d.GCD.ReadyAt()
	if d.ShiftingPower.CanCast(sim, &d.Unit) || d.ShiftingPower.Cast(sim, &d.Unit) {
		t.Fatal("Shifting Power accepted full Energy")
	}
	if d.CurrentMana() != mana || d.GCD.ReadyAt() != gcd || !d.ShiftingPower.CD.IsReady(sim) {
		t.Fatal("rejected Shifting Power charged Mana/GCD/cooldown")
	}
}

func TestOctober9WolfsheadActiveAbilities(t *testing.T) {
	for _, helm := range []bool{false, true} {
		sim, d := october9DruidFixture(map[string]int{"shiftingPower": 1}, helm, proto.Ruleset_RulesetForever)
		d.SpendEnergy(sim, d.CurrentEnergy(), d.NewEnergyMetrics(core.ActionID{SpellID: 1322605}))
		if !d.ShiftingPower.Cast(sim, &d.Unit) {
			t.Fatal("Shifting Power failed")
		}
		want := 40.0
		if helm {
			want = 45
		}
		if d.CurrentEnergy() != want {
			t.Errorf("helm %v: Shifting Power Energy %g, want %g", helm, d.CurrentEnergy(), want)
		}
		sim, d = october9DruidFixture(nil, helm, proto.Ruleset_RulesetForever, true)
		before := d.CurrentRage()
		d.Enrage.ApplyEffects(sim, &d.Unit, d.Enrage.Spell)
		wantRage := 10.0
		if helm {
			wantRage = 15
		}
		if d.CurrentRage()-before != wantRage {
			t.Fatal("Wolfshead's unchanged Enrage Rage bonus was lost")
		}
	}
}

func TestOctober9NaturalInstinctHealing(t *testing.T) {
	for rank := 1; rank <= 2; rank++ {
		sim, d := october9DruidFixture(map[string]int{"predatoryInstincts": rank}, false, proto.Ruleset_RulesetForever)
		fraction := []float64{0, .12, .25}[rank]
		if got := d.GetStat(stats.HealingPower); math.Abs(got-d.GetStat(stats.Intellect)*fraction) > 1e-7 {
			t.Errorf("rank %d: Intellect healing %g, want %g", rank, got, d.GetStat(stats.Intellect)*fraction)
		}
		before := d.GetStats()
		d.AddStatDynamic(sim, stats.Intellect, 100)
		if got := d.GetStat(stats.HealingPower) - before[stats.HealingPower]; math.Abs(got-100*fraction) > 1e-7 {
			t.Errorf("rank %d: dynamic Intellect healing %g, want %g", rank, got, 100*fraction)
		}
		if d.GetStat(stats.SpellPower) != before[stats.SpellPower] || d.GetStat(stats.SpellDamage) != before[stats.SpellDamage] {
			t.Fatal("healing-only Intellect conversion increased spell damage")
		}
		_, baseline := october9DruidFixture(nil, false, proto.Ruleset_RulesetForever)
		if math.Abs(d.Shred.CritDamageBonus-baseline.Shred.CritDamageBonus-.1*float64(rank)) > 1e-9 {
			t.Fatal("existing melee critical damage bonus changed")
		}
	}
}

func TestOctober9ExposePreySeparateWindows(t *testing.T) {
	sim := upstreamFixSim("survival", map[string]int{"exposePrey": 2})
	h := sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
	target := h.CurrentTarget
	if external := target.GetAuraByID(core.ActionID{SpellID: 14325}); external != nil {
		external.Deactivate(sim)
	}
	h.GetSpell(core.ActionID{SpellID: 14325}).ApplyEffects(sim, target, h.GetSpell(core.ActionID{SpellID: 14325}))
	trigger := h.GetAura("Expose Prey")
	h.DistanceFromTarget = 0
	for i := 0; i < 500 && !h.MongooseBite.CanCast(sim, target); i++ {
		trigger.OnSpellHitDealt(trigger, sim, h.AutoAttacks.MHAuto(), &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
	}
	if !h.MongooseBite.CanCast(sim, target) {
		t.Fatal("owned mark did not grant a Mongoose opportunity")
	}
	for _, unit := range sim.Environment.AllUnits {
		unit.AutoAttacks.CancelAutoSwing(sim)
		if unit.Type == core.PlayerUnit {
			unit.CancelGCDTimer(sim)
		}
	}
	sim.AddPendingAction(&core.PendingAction{NextActionAt: 6 * time.Second, OnAction: func(*core.Simulation) {}})
	for sim.CurrentTime < 6*time.Second {
		sim.Step()
	}
	if !h.MongooseBite.CanCast(sim, target) {
		t.Fatal("Expose Prey opportunity expired with the unrelated five-second dodge window")
	}
	if h.DefensiveState.IsActive() {
		t.Fatal("Expose Prey incorrectly extended the ordinary dodge state")
	}
	if !h.MongooseBite.Cast(sim, target) || h.MongooseBite.CanCast(sim, target) {
		t.Fatal("Mongoose did not consume its opportunity")
	}
}

func TestOctober9ThickHideActualGearSwapInForms(t *testing.T) {
	for _, form := range []druid.DruidForm{druid.Cat, druid.Bear, druid.Moonkin} {
		req := historyTalentFixture("feral", map[string]int{"thickHide": 3})
		p := req.Raid.Parties[0].Players[0]
		multiplier := 1.0
		if form == druid.Bear {
			p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{Options: &proto.FeralTankDruid_Options{}}}
			multiplier = druid.BearFormArmorMultiplier
		} else if form == druid.Moonkin {
			req = historyTalentFixture("balance", map[string]int{"thickHide": 3, "moonkinForm": 1})
			p = req.Raid.Parties[0].Players[0]
			multiplier = druid.MoonkinFormArmorMultiplier
		}
		p.ForeverTier1Bonuses = false
		p.Equipment = &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 17)}
		for i := range p.Equipment.Items {
			p.Equipment.Items[i] = &proto.ItemSpec{}
		}
		p.Equipment.Items[14].Id = 7946 // Runed Mithril Hammer: legal Druid Defense weapon.
		p.EnableItemSwap = true
		p.ItemSwap = &proto.ItemSwap{MhItem: &proto.ItemSpec{Id: 10804}} // Fist of the Damned.
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		d := sim.Raid.Parties[0].Players[0].(interface{ GetDruid() *druid.Druid }).GetDruid()
		before := d.GetStat(stats.Armor)
		delta := core.ItemsByID[10804].Stats.Subtract(core.ItemsByID[7946].Stats)
		want := multiplier * (delta[stats.Armor] + 2*delta[stats.Defense])
		if delta[stats.Defense] == 0 {
			t.Fatal("actual gear swap fixture lost Armor/Defense changes")
		}
		sim.CurrentTime = time.Second
		d.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand})
		if got := d.GetStat(stats.Armor) - before; math.Abs(got-want) > 1e-7 {
			t.Fatalf("form %v: swapped Armor/Defense changed armor by %g, want %g", form, got, want)
		}
		sim.CurrentTime += 2 * time.Second
		d.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand})
		if math.Abs(d.GetStat(stats.Armor)-before) > 1e-7 {
			t.Fatal("swapping back leaked form/Thick Hide armor")
		}
	}
}

func TestOctober9UnleashedFuryAppliedOnce(t *testing.T) {
	base := upstreamFixSim("beast_mastery", map[string]int{"summonHawk": 1})
	boost := upstreamFixSim("beast_mastery", map[string]int{"summonHawk": 1, "unleashedFury": 5})
	basePet, boostPet := activePet(t, base), activePet(t, boost)
	if math.Abs(boostPet.PseudoStats.DamageDealtMultiplier/basePet.PseudoStats.DamageDealtMultiplier-1.15) > 1e-9 {
		t.Fatal("permanent companion did not receive exactly one Unleashed Fury multiplier")
	}
	a := base.Raid.Parties[0].Players[0].(*hunter.Hunter)
	b := boost.Raid.Parties[0].Players[0].(*hunter.Hunter)
	if math.Abs(b.Hawks[0].AutoAttacks.MHAuto().DamageMultiplier/a.Hawks[0].AutoAttacks.MHAuto().DamageMultiplier-1.15) > 1e-9 ||
		math.Abs(b.SummonHawk.DamageMultiplier/a.SummonHawk.DamageMultiplier-1.15) > 1e-9 {
		t.Fatal("Hawk impact/assault did not receive exactly one Unleashed Fury multiplier")
	}
}

func TestOctober9HunterPetLowHealthLifecycle(t *testing.T) {
	sim := upstreamFixSim("beast_mastery", map[string]int{"enduranceTraining": 5})
	pet := activePet(t, sim)
	owner := sim.Raid.Parties[0].Players[0].GetCharacter()
	var agent core.PetAgent
	for _, candidate := range sim.Raid.Parties[0].Pets {
		if candidate.GetPet() == pet {
			agent = candidate
		}
	}
	pet.RemoveHealth(sim, .9*pet.MaxHealth())
	health, maximum := pet.CurrentHealth(), pet.MaxHealth()
	for i := 0; i < 20; i++ {
		owner.AddStatDynamic(sim, stats.Stamina, 10)
		owner.AddStatDynamic(sim, stats.Stamina, -10)
		owner.AddStatDynamic(sim, stats.SpellPower, 1)
		owner.AddStatDynamic(sim, stats.SpellPower, -1)
		pet.Enable(sim, agent) // Already enabled: must not reapply inheritance/heal.
		if pet.CurrentHealth() != health || math.Abs(pet.MaxHealth()-maximum) > 1e-7 {
			t.Fatal("low-health pet repeatedly gained current/max health")
		}
	}
	for i := 0; i < 3; i++ {
		pet.Disable(sim)
		pet.Enable(sim, agent)
		if pet.CurrentHealth() != health || math.Abs(pet.MaxHealth()-maximum) > 1e-7 {
			t.Fatal("pet re-enable stacked inherited health or healed at low health")
		}
	}
}

func TestOctober9FireNovaFreshClearcasting(t *testing.T) {
	sim, s := shamanOctober8Fixture(map[string]int{"elementalFocus": 1}, proto.Ruleset_RulesetForever)
	target := s.CurrentTarget
	s.SearingTotem[6].ApplyEffects(sim, target, s.SearingTotem[6])
	nova := s.FireNova[1]
	for i := 0; i < 100 && !s.ClearcastingAura.IsActive(); i++ {
		sim.CurrentTime += 2 * time.Second
		nova.CD.Timer.Set(sim.CurrentTime)
		s.GCD.Set(sim.CurrentTime)
		s.AddMana(sim, s.MaxMana(), s.NewManaMetrics(core.ActionID{SpellID: 16246}))
		if !nova.Cast(sim, target) {
			t.Fatal("Fire Nova did not cast from its active fire totem")
		}
	}
	if !s.ClearcastingAura.IsActive() || s.ClearcastingAura.GetStacks() != 1 {
		t.Fatal("Fire Nova immediately consumed its freshly generated Clearcasting")
	}
	if nova.Cost.GetCurrentCost() != 0 {
		t.Fatal("fresh Clearcasting did not waive the next Fire Nova cost")
	}
	s.GetAura("Elemental Focus Trigger").Deactivate(sim) // Prevent an unrelated refresh on the consuming cast.
	sim.CurrentTime += 2 * time.Second
	nova.CD.Timer.Set(sim.CurrentTime)
	s.GCD.Set(sim.CurrentTime)
	before := s.CurrentMana()
	if !nova.Cast(sim, target) || s.CurrentMana() != before || s.ClearcastingAura.IsActive() || nova.Cost.GetCurrentCost() != 95 {
		t.Fatal("the next Fire Nova did not consume one mature free-cast charge and restore its cost")
	}
}

func TestOctober9FlametongueActualFormSpeed(t *testing.T) {
	for _, bear := range []bool{false, true} {
		sim, d := october9DruidFixture(nil, false, proto.Ruleset_RulesetForever, bear, true)
		// A distinctly slower equipped weapon must not replace active form speed.
		d.MainHand().SwingSpeed = 4
		aura := d.GetAura("Flametongue Totem (Rank 4; provisional)")
		proc := d.GetSpell(core.ActionID{SpellID: 16389})
		proc.BonusHitRating, proc.BonusCritRating = 10000, -10000
		before := proc.SpellMetrics[d.CurrentTarget.UnitIndex].TotalDamage
		aura.OnSpellHitDealt(aura, sim, d.AutoAttacks.MHAuto(), &core.SpellResult{Target: d.CurrentTarget, Outcome: core.OutcomeHit})
		wantSpeed := 1.0
		if bear {
			wantSpeed = 2.5
		}
		if got := proc.SpellMetrics[d.CurrentTarget.UnitIndex].TotalDamage - before; math.Abs(got-13.63*wantSpeed) > 1e-7 {
			t.Fatalf("bear %v: equipped speed leaked into Flametongue proc: %g", bear, got)
		}
	}
}
