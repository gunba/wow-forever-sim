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
	"github.com/wowsims/classic/sim/mage"
	"github.com/wowsims/classic/sim/warlock"
	googleproto "google.golang.org/protobuf/proto"
)

func upstreamFixSim(key string, fields map[string]int) *core.Simulation {
	req := historyTalentFixture(key, fields)
	p := req.Raid.Parties[0].Players[0]
	p.ForeverTier1Bonuses = false
	p.Equipment = &proto.EquipmentSpec{}
	p.BonusStats = nil
	if key == "smite" {
		p.Race = proto.Race_RaceNightElf
	}
	if p.GetWarlock() != nil {
		p.GetWarlock().Options.Sacrifice = proto.WarlockOptions_NoSummon
	}
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim
}

func codeSpell(t *testing.T, unit *core.Unit, code int32) *core.Spell {
	t.Helper()
	for _, spell := range unit.Spellbook {
		if spell.SpellCode == code {
			return spell
		}
	}
	t.Fatalf("missing spell code %d on %s", code, unit.Label)
	return nil
}

func TestUpstreamHunterTalentMasks(t *testing.T) {
	talents := map[string]int{
		"aimedShot": 1, "sniperShot": 1, "summonHawk": 1, "striderKick": 1, "laceratingStrikes": 1,
	}
	baseSim := upstreamFixSim("survival", talents)
	talents["efficiency"], talents["resourcefulness"], talents["predatorsEdge"] = 5, 2, 5
	boostSim := upstreamFixSim("survival", talents)
	base, boost := baseSim.Raid.AllPlayerUnits[0], boostSim.Raid.AllPlayerUnits[0]

	for _, tc := range []struct {
		code int32
		cost int32
		crit float64
	}{
		{hunter.SpellCode_HunterSniperShot, 0, 0},
		{hunter.SpellCode_HunterStriderKick, 0, .30},
		{hunter.SpellCode_HunterSummonHawk, 15, 0},
		{hunter.SpellCode_HunterVolley, 15, 0},
		{hunter.SpellCode_HunterAimedShot, 15, 0},
		{hunter.SpellCode_HunterSerpentSting, 15, 0},
		{hunter.SpellCode_HunterWingClip, 75, .30},
		{hunter.SpellCode_HunterMongooseBite, 75, .30},
		{hunter.SpellCode_HunterExplosiveTrap, 60, 0},
		{hunter.SpellCode_HunterLaceratingStrikes, 0, .30},
	} {
		b, a := codeSpell(t, base, tc.code), codeSpell(t, boost, tc.code)
		if a.Cost != nil && b.Cost.Multiplier-a.Cost.Multiplier != tc.cost {
			t.Errorf("%s cost reduction %d, want %d", a.ActionID, b.Cost.Multiplier-a.Cost.Multiplier, tc.cost)
		}
		if got := a.CritDamageBonus - b.CritDamageBonus; math.Abs(got-tc.crit) > 1e-9 {
			t.Errorf("%s crit damage gained %v, want %v", a.ActionID, got, tc.crit)
		}
	}
	if boost.AutoAttacks.MHAuto().CritDamageBonus != base.AutoAttacks.MHAuto().CritDamageBonus ||
		boost.AutoAttacks.OHAuto().CritDamageBonus != base.AutoAttacks.OHAuto().CritDamageBonus {
		t.Fatal("Predator's Edge increased white-swing crit damage")
	}
}

func TestUpstreamFocusedFireOwnerAndPetLifecycle(t *testing.T) {
	sim := upstreamFixSim("beast_mastery", map[string]int{"focusedFire": 2})
	owner := sim.Raid.AllPlayerUnits[0]
	pet := activePet(t, sim)
	aura := owner.GetAura("Focused Fire")
	if aura == nil || !aura.IsActive() {
		t.Fatal("Focused Fire not active with the pet")
	}
	ownerActive, petActive := owner.PseudoStats.DamageDealtMultiplier, pet.PseudoStats.DamageDealtMultiplier
	var agent core.PetAgent
	for _, p := range sim.Raid.Parties[0].Pets {
		if p.GetPet() == pet {
			agent = p
		}
	}
	for i := 0; i < 2; i++ {
		pet.Disable(sim)
		if aura.IsActive() ||
			math.Abs(owner.PseudoStats.DamageDealtMultiplier*1.02-ownerActive) > 1e-9 ||
			math.Abs(pet.PseudoStats.DamageDealtMultiplier*1.02-petActive) > 1e-9 {
			t.Fatal("Focused Fire persisted after the pet was disabled")
		}
		pet.Enable(sim, agent)
		if !aura.IsActive() || math.Abs(owner.PseudoStats.DamageDealtMultiplier-ownerActive) > 1e-9 ||
			math.Abs(pet.PseudoStats.DamageDealtMultiplier-petActive) > 1e-9 {
			t.Fatal("Focused Fire missing or stacked after re-enabling the pet")
		}
	}
}

func TestUpstreamDemonicKnowledgeAddsExplicitPetBonus(t *testing.T) {
	base := upstreamFixSim("demonology", nil)
	boost := upstreamFixSim("demonology", map[string]int{"demonicKnowledge": 3})
	bOwner, owner := base.Raid.AllPlayerUnits[0], boost.Raid.AllPlayerUnits[0]
	bPet, pet := activePet(t, base), activePet(t, boost)
	bWarlock := base.Raid.Parties[0].Players[0].(warlock.WarlockAgent).GetWarlock()
	aWarlock := boost.Raid.Parties[0].Players[0].(warlock.WarlockAgent).GetWarlock()
	inheritance := core.ForeverPetInheritance(owner.GetStats(), true).Subtract(
		core.ForeverPetInheritance(bOwner.GetStats(), true))
	if got := owner.GetStat(stats.SpellPower) - bOwner.GetStat(stats.SpellPower); math.Abs(got-60) > 1e-9 {
		t.Fatalf("owner gain %v, want 60", got)
	}
	if got, want := pet.GetStat(stats.SpellPower)-bPet.GetStat(stats.SpellPower), 60+inheritance[stats.SpellPower]; math.Abs(got-want) > 1e-9 {
		t.Fatalf("pet gain %v, want %v including ordinary inheritance", got, want)
	}
	for i := 0; i < 2; i++ {
		aWarlock.ActivePet.Disable(boost, false)
		bWarlock.ActivePet.Disable(base, false)
		if owner.GetAura("Demonic Knowledge").IsActive() || pet.GetAura("Demonic Knowledge").IsActive() {
			t.Fatal("Demonic Knowledge survived pet dismissal")
		}
		if pet.GetStat(stats.SpellPower) != bPet.GetStat(stats.SpellPower) {
			t.Fatal("dismissed pet retained spell power")
		}
		bWarlock.ActivePet.Enable(base, bWarlock.ActivePet)
		aWarlock.ActivePet.Enable(boost, aWarlock.ActivePet)
		if got, want := pet.GetStat(stats.SpellPower)-bPet.GetStat(stats.SpellPower), 60+inheritance[stats.SpellPower]; math.Abs(got-want) > 1e-9 {
			t.Fatalf("reset pet gain %v, want %v", got, want)
		}
	}
}

func TestUpstreamMoonfuryScalesSpellPowerAndSubtletyFlatThreat(t *testing.T) {
	base := upstreamFixSim("balance", nil)
	boost := upstreamFixSim("balance", map[string]int{"moonfury": 5, "subtlety": 3})
	b, a := base.Raid.AllPlayerUnits[0], boost.Raid.AllPlayerUnits[0]
	for _, school := range []stats.SchoolIndex{stats.SchoolIndexNature, stats.SchoolIndexArcane} {
		if math.Abs(a.PseudoStats.SchoolDamageDealtMultiplier[school]/b.PseudoStats.SchoolDamageDealtMultiplier[school]-1.1) > 1e-9 {
			t.Fatalf("Moonfury did not multiply school %v", school)
		}
	}
	b.AddStatDynamic(base, stats.SpellPower, 1000)
	a.AddStatDynamic(boost, stats.SpellPower, 1000)
	before, after := codeSpell(t, b, druid.SpellCode_DruidWrath), codeSpell(t, a, druid.SpellCode_DruidWrath)
	if after.BaseDamageMultiplierAdditive != before.BaseDamageMultiplierAdditive {
		t.Fatal("Moonfury also modified the base damage")
	}
	bResult := before.CalcDamage(base, b.CurrentTarget, 100, before.OutcomeAlwaysHit)
	aResult := after.CalcDamage(boost, a.CurrentTarget, 100, after.OutcomeAlwaysHit)
	if math.Abs(aResult.Damage/bResult.Damage-1.1) > 1e-9 {
		t.Fatalf("full Wrath damage ratio %v, want 1.1", aResult.Damage/bResult.Damage)
	}
	var checked bool
	for _, spell := range a.Spellbook {
		if spell.SpellSchool.Matches(core.SpellSchoolNature|core.SpellSchoolArcane) && spell.FlatThreatBonus != 0 {
			old := b.GetSpell(spell.ActionID)
			if old == nil {
				t.Fatalf("missing original %s", spell.ActionID)
			}
			got := spell.ThreatFromDamage(core.OutcomeHit, 100)
			want := .7 * old.ThreatFromDamage(core.OutcomeHit, 100)
			if math.Abs(got-want) > 1e-9 {
				t.Fatalf("%s threat %v, want %v", spell.ActionID, got, want)
			}
			checked = true
		}
	}
	if !checked {
		t.Fatal("no flat-threat Nature/Arcane spell was checked")
	}
}

func TestUpstreamPriestInnerFocusAndPenanceDiscount(t *testing.T) {
	sim := upstreamFixSim("smite", map[string]int{"innerFocus": 1, "penance": 1, "mindFlay": 1, "improvedHealing": 3})
	unit := sim.Raid.AllPlayerUnits[0]
	penance := unit.GetSpell(core.ActionID{SpellID: 1316995})
	if got := penance.Cost.GetCurrentCost(); math.Abs(got-355*.85) > 1e-9 {
		t.Fatalf("Penance cost %v, want %v", got, 355*.85)
	}
	aura := unit.GetAura("Inner Focus")
	beforeCrit, beforeCost := map[*core.Spell]float64{}, map[*core.Spell]float64{}
	for _, spell := range unit.Spellbook {
		beforeCrit[spell] = spell.BonusCritRating
		if spell.Cost != nil {
			beforeCost[spell] = spell.Cost.GetCurrentCost()
		}
	}
	aura.Activate(sim)
	checked := map[int32]bool{}
	for _, spell := range unit.Spellbook {
		wantCrit := -1.0
		switch spell.SpellID {
		case 18807, 1309636, 19305: // Mind Flay, Death, Starshards
			wantCrit = 0
		case 10947, 10934, 1316995: // Mind Blast, Smite, Penance's modeled bolts
			wantCrit = 25 * core.SpellCritRatingPerCritChance
		}
		if wantCrit < 0 {
			continue
		}
		checked[spell.SpellID] = true
		if got := spell.BonusCritRating - beforeCrit[spell]; math.Abs(got-wantCrit) > 1e-9 {
			t.Errorf("%s Inner Focus crit %v, want %v", spell.ActionID, got, wantCrit)
		}
		if spell.Cost != nil && spell.Cost.GetCurrentCost() != 0 {
			t.Errorf("%s not free during Inner Focus", spell.ActionID)
		}
	}
	for _, id := range []int32{18807, 1309636, 19305, 10947, 10934, 1316995} {
		if !checked[id] {
			t.Errorf("spell %d was not checked", id)
		}
	}
	aura.Deactivate(sim)
	for _, spell := range unit.Spellbook {
		if spell.BonusCritRating != beforeCrit[spell] ||
			(spell.Cost != nil && math.Abs(spell.Cost.GetCurrentCost()-beforeCost[spell]) > 1e-9) {
			t.Errorf("%s did not restore crit/cost after Inner Focus", spell.ActionID)
		}
	}
}

func TestUpstreamTrueshotAndThorns(t *testing.T) {
	base := upstreamFixSim("beast_mastery", nil)
	req := historyTalentFixture("beast_mastery", nil)
	p := req.Raid.Parties[0].Players[0]
	p.Equipment, p.BonusStats, p.ForeverTier1Bonuses = &proto.EquipmentSpec{}, nil, false
	req.Raid.Buffs.TrueshotAura = true
	req.Raid.Buffs.Thorns = proto.TristateEffect_TristateEffectRegular
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	unit := sim.Raid.AllPlayerUnits[0]
	old := base.Raid.AllPlayerUnits[0]
	if got := unit.GetStat(stats.RangedAttackPower) - old.GetStat(stats.RangedAttackPower); got != 75 {
		t.Fatalf("Trueshot gave %v RAP, want 75", got)
	}
	if unit.GetStat(stats.AttackPower) != old.GetStat(stats.AttackPower) || unit.GetAura("Trueshot Aura").ActionID.SpellID != 20905 {
		t.Fatal("Trueshot modified melee AP or used the weaker rank")
	}
	unit.AddStatDynamic(sim, stats.SpellHit, 100)
	spell := unit.GetSpell(core.ActionID{SpellID: 9910})
	spell.Cast(sim, unit.CurrentTarget)
	if got := spell.SpellMetrics[unit.CurrentTarget.UnitIndex].TotalDamage; math.Abs(got-22) > 1e-9 {
		t.Fatalf("Thorns damage %v, want 22", got)
	}
}

func TestUpstreamLifeTapSpiritPaymentAndPetMana(t *testing.T) {
	for _, tier := range []bool{false, true} {
		req := historyTalentFixture("demonology", map[string]int{"improvedLifeTap": 2, "demonicEnergies": 1})
		p := req.Raid.Parties[0].Players[0]
		p.ForeverTier1Bonuses, p.Equipment, p.BonusStats = tier, &proto.EquipmentSpec{}, nil
		p.GetWarlock().Options.Sacrifice = proto.WarlockOptions_NoSummon
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		owner := sim.Raid.Parties[0].Players[0].GetCharacter()
		pet := activePet(t, sim)
		tap := owner.GetSpell(core.ActionID{SpellID: 11689})
		spend := owner.NewManaMetrics(core.ActionID{SpellID: 1})
		petSpend := pet.NewManaMetrics(core.ActionID{SpellID: 1})
		multiplier := 1.0
		if tier {
			multiplier = 1.2
		}
		pay := func(tanking bool) {
			t.Helper()
			owner.SpendMana(sim, owner.CurrentMana(), spend)
			pet.SpendMana(sim, pet.CurrentMana(), petSpend)
			health := owner.CurrentHealth()
			wantCost := (424 + owner.GetStat(stats.Spirit)) * 1.2
			tap.ApplyEffects(sim, owner.CurrentTarget, tap)
			if math.Abs(owner.CurrentMana()-wantCost*multiplier) > 1e-9 {
				t.Fatalf("Life Tap mana %v, want %v", owner.CurrentMana(), wantCost*multiplier)
			}
			if math.Abs(pet.CurrentMana()-wantCost*multiplier*.5) > 1e-9 {
				t.Fatalf("pet Life Tap mana %v, want %v", pet.CurrentMana(), wantCost*multiplier*.5)
			}
			wantHealth := health
			if tanking {
				wantHealth -= wantCost
			}
			if math.Abs(owner.CurrentHealth()-wantHealth) > 1e-9 {
				t.Fatalf("Life Tap health %v, want %v", owner.CurrentHealth(), wantHealth)
			}
		}
		pay(false)
		// Neither spell damage nor outgoing/incoming damage multipliers are
		// inputs to a resource conversion.
		owner.AddStatsDynamic(sim, stats.Stats{stats.SpellPower: 1000, stats.ShadowPower: 300})
		owner.PseudoStats.DamageDealtMultiplier *= 3
		owner.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexShadow] *= 2
		owner.PseudoStats.DamageTakenMultiplier *= .1
		pay(false)
		owner.AddStatDynamic(sim, stats.Spirit, 100)
		pay(false)
		sim.Encounter.TargetUnits[0].CurrentTarget = &owner.Unit
		pay(true)
		for _, metrics := range tap.SpellMetrics {
			if metrics.TotalDamage != 0 || metrics.TotalThreat != 0 {
				t.Fatal("Life Tap generated damage/threat")
			}
		}
		if tap.ProcMask != core.ProcMaskEmpty {
			t.Fatal("Life Tap is still eligible as a damage proc")
		}
		owner.SpendHealth(sim, owner.CurrentHealth()-1, owner.NewHealthMetrics(core.ActionID{SpellID: 1}))
		if tap.ExtraCastCondition(sim, owner.CurrentTarget) {
			t.Fatal("tanking Life Tap was allowed without enough health")
		}
	}
}

func TestUpstreamMageFrostAreaSpellbook(t *testing.T) {
	req := historyTalentFixture("frost", map[string]int{
		"improvedConeOfCold": 3, "improvedFrostNova": 2, "frostChanneling": 3,
		"elementalPrecision": 3, "iceShards": 5, "coldSnap": 1,
	})
	p := req.Raid.Parties[0].Players[0]
	p.Equipment, p.BonusStats, p.ForeverTier1Bonuses = &proto.EquipmentSpec{}, nil, false
	req.Encounter.Targets[0].Level = 60
	for i := 0; i < 2; i++ {
		req.Encounter.Targets = append(req.Encounter.Targets, googleproto.Clone(req.Encounter.Targets[0]).(*proto.Target))
	}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	m := sim.Raid.Parties[0].Players[0].(mage.MageAgent).GetMage()
	for _, id := range []int32{122, 865, 6131, 10230, 120, 8492, 10159, 10160, 10161} {
		if spell := m.GetSpell(core.ActionID{SpellID: id}); spell == nil || !spell.Flags.Matches(core.SpellFlagAPL) {
			t.Fatalf("missing selectable rank %d", id)
		}
	}
	cone, nova := m.ConeOfCold[5], m.FrostNova[4]
	m.DistanceFromTarget = 20
	if cone.ExtraCastCondition(sim, m.CurrentTarget) || nova.ExtraCastCondition(sim, m.CurrentTarget) {
		t.Fatal("short-range Frost spells were allowed at 20 yards")
	}
	m.DistanceFromTarget = 10
	if !cone.ExtraCastCondition(sim, m.CurrentTarget) || !nova.ExtraCastCondition(sim, m.CurrentTarget) {
		t.Fatal("short-range Frost spells were rejected at 10 yards")
	}
	if cone.CD.Duration != 10*time.Second || nova.CD.Duration != 21*time.Second ||
		cone.CD.Timer != m.ConeOfCold[1].CD.Timer || nova.CD.Timer != m.FrostNova[1].CD.Timer {
		t.Fatal("Frost AOE cooldown/rank sharing incorrect")
	}
	if cone.DamageMultiplier != 1.35 || cone.BonusCoefficient != .129 || nova.BonusCoefficient != .029 ||
		math.Abs(cone.Cost.GetCurrentCost()-555*.85) > 1e-9 ||
		math.Abs(nova.Cost.GetCurrentCost()-145*.85) > 1e-9 ||
		cone.BonusHitRating != 3*core.SpellHitRatingPerHitChance || cone.CritDamageBonus != 2 {
		t.Fatal("Frost AOE coefficients or talent modifiers incorrect")
	}
	if !cone.Flags.Matches(mage.SpellFlagChillSpell) || nova.Flags.Matches(mage.SpellFlagChillSpell) {
		t.Fatal("a root was treated as a chilling slow")
	}
	m.AddStatsDynamic(sim, stats.Stats{stats.SpellHit: 100, stats.SpellCrit: -10000})
	for _, spell := range []*core.Spell{cone, nova} {
		spell.ApplyEffects(sim, m.CurrentTarget, spell)
		for _, target := range sim.Encounter.TargetUnits {
			metrics := spell.SpellMetrics[target.UnitIndex]
			if metrics.Hits+metrics.Crits+metrics.Misses != 1 {
				t.Fatalf("%s did not attempt damage on each AOE target", spell.ActionID)
			}
			if metrics.TotalDamage > 0 {
				low, high := 71.6066, 79.3934 // Frost Nova's level-59 cap.
				if spell == cone {
					// Level 60 adds 3, not the level-63 maximum of 7.5.
					low, high = (343-14.5715)*1.35, (343+14.5715)*1.35
				}
				if metrics.TotalDamage < low || metrics.TotalDamage > high {
					t.Errorf("%s damage %v outside %v–%v", spell.ActionID, metrics.TotalDamage, low, high)
				}
			}
		}
	}
	cone.CD.Use(sim)
	nova.CD.Use(sim)
	m.GetSpell(core.ActionID{SpellID: 12472}).Cast(sim, m.CurrentTarget)
	if !cone.CD.IsReady(sim) || !nova.CD.IsReady(sim) {
		t.Fatal("Cold Snap did not reset the added Frost spells")
	}
}
