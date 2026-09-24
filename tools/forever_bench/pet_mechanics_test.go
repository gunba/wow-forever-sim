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
)

func foreverPetFixture(build string, race proto.Race) *core.Simulation {
	req := racialFixture(build, race)
	p := req.Raid.Parties[0].Players[0]
	p.Equipment = &proto.EquipmentSpec{}
	p.ForeverTier1Bonuses = false
	p.Rotation = &proto.APLRotation{}
	p.Consumes = &proto.Consumes{}
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim
}

func activePet(t *testing.T, sim *core.Simulation) *core.Pet {
	t.Helper()
	for _, pet := range sim.Raid.Parties[0].Pets {
		if pet.GetPet().IsEnabled() {
			return pet.GetPet()
		}
	}
	t.Fatal("no active pet")
	return nil
}

func TestForeverPetInheritanceAndDynamicPowerChoice(t *testing.T) {
	owner := stats.Stats{
		stats.Stamina: 100, stats.Armor: 500, stats.AttackPower: 300,
		stats.RangedAttackPower: 500, stats.SpellPower: 400,
		stats.MeleeHit: 3, stats.SpellHit: 4,
		stats.MeleeCrit: 11, stats.SpellCrit: 12,
	}
	hunter := core.ForeverPetInheritance(owner, false)
	if hunter[stats.Health] != 200 || hunter[stats.Armor] != 150 ||
		hunter[stats.AttackPower] != 50 || hunter[stats.SpellPower] != 0 ||
		hunter[stats.MeleeHit] != 3 || hunter[stats.SpellHit] != 4 ||
		hunter[stats.MeleeCrit] != 11 || hunter[stats.SpellCrit] != 12 {
		t.Fatalf("hunter inheritance: %v", hunter)
	}
	warlock := core.ForeverPetInheritance(owner, true)
	if warlock[stats.SpellPower] != 40 || warlock[stats.AttackPower] != 50 {
		t.Fatalf("warlock inheritance: %v", warlock)
	}
	owner[stats.RangedAttackPower] = 250
	if got := core.ForeverPetInheritance(owner, false)[stats.AttackPower]; got != 30 {
		t.Fatalf("highest-power source did not switch: %v", got)
	}

	for _, tc := range []struct {
		build string
		race  proto.Race
	}{
		{"beast_mastery", proto.Race_RaceOrc},
		{"affliction", proto.Race_RaceUndead},
	} {
		t.Run(tc.build, func(t *testing.T) {
			sim := foreverPetFixture(tc.build, tc.race)
			pet := activePet(t, sim)
			owner := sim.Raid.Parties[0].Players[0].GetCharacter()
			before := pet.GetStat(stats.MeleeCrit)
			owner.AddStatDynamic(sim, stats.MeleeCrit, 5)
			if got := pet.GetStat(stats.MeleeCrit) - before; math.Abs(got-5) > 1e-9 {
				t.Fatalf("pet did not inherit live owner crit: %v", got)
			}
			baseAP := pet.GetStat(stats.AttackPower)
			owner.AddStatDynamic(sim, stats.AttackPower, 10000)
			boostedAP := pet.GetStat(stats.AttackPower)
			if boostedAP < baseAP+900 {
				t.Fatalf("pet did not switch to the owner's new higher power: %v -> %v", baseAP, boostedAP)
			}
			owner.AddStatDynamic(sim, stats.AttackPower, -10000)
			if got := pet.GetStat(stats.AttackPower); math.Abs(got-baseAP) > 1e-6 {
				t.Fatalf("pet did not restore its prior power after a swap: %v -> %v", baseAP, got)
			}
		})
	}
}

func TestForeverHunterPetFocusAndSpeed(t *testing.T) {
	sim := foreverPetFixture("survival", proto.Race_RaceOrc)
	pet := activePet(t, sim)
	if got := pet.CurrentFocusPerSecond(); got != 10 {
		t.Fatalf("focus regeneration is %v/second, want 10", got)
	}
	if got := pet.CurrentFocus(); got != 150 {
		t.Fatalf("Hunter Pet Scaling maximum Focus is %v, want 150", got)
	}
	pet.SpendFocus(sim, 100, pet.NewFocusMetrics(core.ActionID{SpellID: 3009}))
	sim.CurrentTime = 2250 * time.Millisecond
	if got := pet.CurrentFocus(); math.Abs(got-72.5) > 1e-6 {
		t.Fatalf("Focus at 2.25s is %v, want continuously accrued 72.5", got)
	}

	base := pet.AutoAttacks.MH().BaseDamageMin
	req := racialFixture("survival", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].GetHunter().Options.PetAttackSpeed = proto.Hunter_Options_Two
	req.Raid.Parties[0].Players[0].Equipment = &proto.EquipmentSpec{}
	req.Raid.Parties[0].Players[0].ForeverTier1Bonuses = false
	slow := core.NewSim(req, simsignals.Signals{})
	slow.Reset()
	slowPet := activePet(t, slow)
	if got := slowPet.AutoAttacks.MH().BaseDamageMin; got != base {
		t.Fatalf("1.2s and 2.0s pets have different Forever base hit damage: %v vs %v", base, got)
	}
	if fast, slow := pet.AutoAttacks.MH().CalculateAverageWeaponDamage(300),
		slowPet.AutoAttacks.MH().CalculateAverageWeaponDamage(300); fast != slow {
		t.Fatalf("speed changed a Forever pet's AP-added hit: fast %v, slow %v", fast, slow)
	}

	req = racialFixture("survival", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].GetHunter().Options.PetType = proto.Hunter_Options_WindSerpent
	req.Raid.Parties[0].Players[0].Equipment = &proto.EquipmentSpec{}
	req.Raid.Parties[0].Players[0].ForeverTier1Bonuses = false
	wind := core.NewSim(req, simsignals.Signals{})
	wind.Reset()
	breath := activePet(t, wind).GetSpell(core.ActionID{SpellID: 25012})
	if breath == nil || breath.BonusCoefficient != 0 {
		t.Fatalf("Forever Lightning Breath has an unsupported bonus coefficient: %v", breath)
	}
}

func TestForeverPetConsumesDoNotBuffPetsDirectly(t *testing.T) {
	for _, tc := range []struct {
		build string
		race  proto.Race
	}{
		{"beast_mastery", proto.Race_RaceOrc},
		{"affliction", proto.Race_RaceUndead},
	} {
		t.Run(tc.build, func(t *testing.T) {
			sim := foreverPetFixture(tc.build, tc.race)
			base := activePet(t, sim).GetStats()
			req := racialFixture(tc.build, tc.race)
			p := req.Raid.Parties[0].Players[0]
			p.Equipment = &proto.EquipmentSpec{}
			p.ForeverTier1Bonuses = false
			p.Rotation = &proto.APLRotation{}
			p.Consumes = &proto.Consumes{
				PetAttackPowerConsumable: 1,
				PetAgilityConsumable:     1,
				PetStrengthConsumable:    1,
			}
			req.Encounter.Targets[0].Level = 60
			withConsumes := core.NewSim(req, simsignals.Signals{})
			withConsumes.Reset()
			got := activePet(t, withConsumes).GetStats()
			for _, stat := range []stats.Stat{stats.AttackPower, stats.Strength, stats.Agility} {
				if base[stat] != got[stat] {
					t.Fatalf("direct pet consumable changed %v: %v -> %v", stat, base[stat], got[stat])
				}
			}
		})
	}
}

func TestForeverPetReceivesOwnerBuffOnlyThroughInheritance(t *testing.T) {
	plain := foreverPetFixture("survival", proto.Race_RaceOrc)
	owner0 := plain.Raid.Parties[0].Players[0].GetCharacter().GetStats()
	pet0 := activePet(t, plain).GetStats()

	req := racialFixture("survival", proto.Race_RaceOrc)
	p := req.Raid.Parties[0].Players[0]
	p.Equipment = &proto.EquipmentSpec{}
	p.ForeverTier1Bonuses = false
	p.Rotation = &proto.APLRotation{}
	p.Consumes = &proto.Consumes{}
	req.Raid.Buffs.TrueshotAura = true
	req.Encounter.Targets[0].Level = 60
	withBuff := core.NewSim(req, simsignals.Signals{})
	withBuff.Reset()
	owner1 := withBuff.Raid.Parties[0].Players[0].GetCharacter().GetStats()
	pet1 := activePet(t, withBuff).GetStats()
	power := func(s stats.Stats) float64 { return max(s[stats.AttackPower], s[stats.RangedAttackPower]) }
	want := 0.1 * (power(owner1) - power(owner0))
	if got := pet1[stats.AttackPower] - pet0[stats.AttackPower]; math.Abs(got-want) > 1e-6 {
		t.Fatalf("owner AP buff directly applied to pet: pet gained %v, inherited amount %v", got, want)
	}
}

func TestClassicPetFocusAndSpeedUnchanged(t *testing.T) {
	req := racialFixture("survival", proto.Race_RaceOrc)
	req.SimOptions.Ruleset = proto.Ruleset_RulesetClassic
	p := req.Raid.Parties[0].Players[0]
	p.Equipment = &proto.EquipmentSpec{}
	p.ForeverTier1Bonuses = false
	p.Rotation = &proto.APLRotation{}
	p.GetHunter().Options.PetAttackSpeed = proto.Hunter_Options_OneTwo
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Reset()
	pet := activePet(t, sim)
	if pet.CurrentFocus() != 100 || pet.CurrentFocusPerSecond() != 5 {
		t.Fatalf("Classic pet Focus changed: cap %v, regen %v/s", pet.CurrentFocus(), pet.CurrentFocusPerSecond())
	}
	if got := pet.AutoAttacks.MH().BaseDamageMin; math.Abs(got-18.17*1.2) > 1e-6 {
		t.Fatalf("Classic pet base hit no longer scales with swing time: %v", got)
	}
	if pet.AutoAttacks.MH().APScalingSpeed != 0 {
		t.Fatal("Classic pet AP scaling was overridden")
	}
}
