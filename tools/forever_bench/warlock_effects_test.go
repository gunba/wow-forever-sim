//go:build with_db

package main

import (
	"math"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/warlock/dps"
)

func warlockEffectFixture(talents map[string]int) *proto.RaidSimRequest {
	req := historyTalentFixture("demonology", talents)
	p := req.Raid.Parties[0].Players[0]
	p.Race = proto.Race_RaceHuman
	p.ForeverTier1Bonuses = false
	p.GetWarlock().Options.Sacrifice = proto.WarlockOptions_NoSummon
	p.Equipment.Items[14] = &proto.ItemSpec{Id: 272683}
	p.Equipment.Items[15] = &proto.ItemSpec{Id: 272685}
	return req
}

func TestForeverWarlockWeaponStones(t *testing.T) {
	makeUnit := func(stone proto.WarlockOptions_WeaponImbue, emptyHand bool) *core.Unit {
		req := warlockEffectFixture(nil)
		p := req.Raid.Parties[0].Players[0]
		p.GetWarlock().Options.WeaponImbue = stone
		if emptyHand {
			p.Equipment.Items[14] = &proto.ItemSpec{}
		}
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		return env.Raid.AllPlayerUnits[0]
	}
	base := makeUnit(proto.WarlockOptions_NoWeaponImbue, false)
	for _, stone := range []proto.WarlockOptions_WeaponImbue{proto.WarlockOptions_Firestone, proto.WarlockOptions_Spellstone} {
		unit := makeUnit(stone, false)
		want := stats.Stats{stats.FirePower: 21}
		if stone == proto.WarlockOptions_Firestone {
			want[stats.SpellCrit] = 2 * core.SpellCritRatingPerCritChance
		} else {
			want[stats.ShadowPower] = 21
		}
		for stat, diff := range unit.GetStats().Subtract(base.GetStats()) {
			if math.Abs(diff-want[stat]) > 1e-8 {
				t.Errorf("%v stat %v gain %v, want %v", stone, stats.Stat(stat), diff, want[stat])
			}
		}
		wantSpeed := base.PseudoStats.CastSpeedMultiplier
		if stone == proto.WarlockOptions_Spellstone {
			wantSpeed *= 1.02
		}
		if math.Abs(unit.PseudoStats.CastSpeedMultiplier-wantSpeed) > 1e-9 ||
			unit.SwingSpeed() != base.SwingSpeed() || unit.RangedSwingSpeed() != base.RangedSwingSpeed() {
			t.Fatal("stone haste was missing or leaked into attack speed")
		}
		if unit.GetAura("Firestone Proc") != nil {
			t.Fatal("obsolete Firestone melee proc registered")
		}
		emptyBase := makeUnit(proto.WarlockOptions_NoWeaponImbue, true)
		emptyStone := makeUnit(stone, true)
		if emptyBase.GetStats() != emptyStone.GetStats() ||
			emptyBase.PseudoStats.CastSpeedMultiplier != emptyStone.PseudoStats.CastSpeedMultiplier {
			t.Fatal("stone granted bonuses without a main-hand weapon")
		}
	}
}

func TestWarlockStoneRejectsOtherImbue(t *testing.T) {
	req := warlockEffectFixture(nil)
	p := req.Raid.Parties[0].Players[0]
	p.GetWarlock().Options.WeaponImbue = proto.WarlockOptions_Firestone
	p.Consumes.MainHandImbue = proto.WeaponImbue_BrilliantWizardOil
	defer func() {
		err := recover()
		if err == nil || !strings.Contains(err.(string), "cannot be combined") {
			t.Fatalf("expected a conflicting-imbue error, got %v", err)
		}
	}()
	core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
}

func TestDemonicBrandDamageAndTargetScope(t *testing.T) {
	for _, summon := range []proto.WarlockOptions_Summon{proto.WarlockOptions_Imp, proto.WarlockOptions_Succubus} {
		t.Run(summon.String(), func(t *testing.T) {
			req := warlockEffectFixture(map[string]int{"demonicBrand": 3, "unholyPower": 5, "masterDemonologist": 5})
			req.Raid.Parties[0].Players[0].GetWarlock().Options.Summon = summon
			req.Encounter.Targets = append(req.Encounter.Targets, req.Encounter.Targets[0])
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			w := sim.Raid.Parties[0].Players[0].(*dps.DpsWarlock).GetWarlock()
			pet := w.ActivePet
			target, other := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
			id, powerStat := int32(1293697), stats.ShadowPower
			attack := pet.GetSpell(core.ActionID{SpellID: 11780}) // Lash of Pain
			if summon == proto.WarlockOptions_Imp {
				id, powerStat = 1293698, stats.FirePower
				attack = pet.GetSpell(core.ActionID{SpellID: 11763})
			}
			brand := pet.GetSpell(core.ActionID{SpellID: id})
			brand.Flags |= core.SpellFlagIgnoreResists
			metrics := &brand.SpellMetrics[target.UnitIndex]
			pain := w.GetSpell(core.ActionID{SpellID: 17923})
			pain.CalcAndDealDamage(sim, target, 1, pain.OutcomeAlwaysHit)
			aura := target.GetAura("Demonic Brand-" + w.Label)
			if aura == nil || aura.GetStacks() != 6 || other.GetAura("Demonic Brand-"+w.Label).IsActive() {
				t.Fatal("Searing Pain did not brand only its target")
			}
			attack.CalcAndDealDamage(sim, other, 1, attack.OutcomeAlwaysHit)
			if aura.GetStacks() != 6 {
				t.Fatal("attacking a different target consumed the brand")
			}
			hit := func() float64 {
				before := metrics.TotalDamage
				attack.CalcAndDealDamage(sim, target, 1, attack.OutcomeAlwaysHit)
				return metrics.TotalDamage - before
			}
			// The client formula's two 10% pet multipliers must apply once each.
			const multiplier = 1.1 * 1.1
			power := w.GetStat(stats.SpellPower) + w.GetStat(stats.SpellDamage) + w.GetStat(powerStat)
			sim.Reseed(101)
			first := hit()
			if first < (65+.078*power)*multiplier || first > (68+.078*power)*multiplier || metrics.Misses != 0 {
				t.Fatalf("brand damage %v outside client range with power %v", first, power)
			}
			w.AddStatDynamic(sim, powerStat, 100)
			sim.Reseed(101)
			second := hit()
			if math.Abs(second-first-7.8*multiplier) > 1e-8 {
				t.Fatalf("matching-school coefficient or pet modifiers wrong: delta %v", second-first)
			}
			otherPower := stats.FirePower
			if powerStat == stats.FirePower {
				otherPower = stats.ShadowPower
			}
			w.AddStatDynamic(sim, otherPower, 500)
			sim.Reseed(101)
			if third := hit(); math.Abs(third-second) > 1e-8 {
				t.Fatal("brand incorrectly used the other school's spell power")
			}
			for i := 0; i < 3; i++ {
				hit()
			}
			if aura.IsActive() || metrics.Hits != 6 || metrics.Misses != 0 {
				t.Fatal("brand failed to consume exactly six charges without a second miss roll")
			}
			if hit() != 0 {
				t.Fatal("brand dealt damage with no charges")
			}
		})
	}
}
