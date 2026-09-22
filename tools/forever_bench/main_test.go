//go:build with_db

package main

import (
	"math"
	"os"
	"reflect"
	"testing"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestMain(m *testing.M) {
	if err := os.Chdir("../.."); err != nil {
		panic(err)
	}
	sim.RegisterAll()
	os.Exit(m.Run())
}

func TestBuildsAreLegal(t *testing.T) {
	for _, b := range builds() {
		if err := loadTalents(b).validate(b.presetTalents()); err != nil {
			t.Errorf("%s: %v", b.Key, err)
		}
	}
}

func TestFireAndElementalTalentCorrections(t *testing.T) {
	want := map[string]map[string]int{
		"fire":      {"improvedFireball": 4, "wakeOfFire": 2, "hotStreak": 1},
		"elemental": {"ancestralKnowledge": 2, "improvedLightningShield": 0, "waterShield": 0},
	}
	for _, b := range builds() {
		fields, ok := want[b.Key]
		if !ok {
			continue
		}
		config := loadTalents(b)
		points, err := config.decode(b.presetTalents())
		if err != nil {
			t.Fatal(err)
		}
		index := 0
		for _, tree := range config.Trees {
			for _, talent := range tree.Talents {
				if expected, checked := fields[talent.Field]; checked && points[index] != expected {
					t.Errorf("%s/%s = %d, want %d", b.Key, talent.Field, points[index], expected)
				}
				index++
			}
		}
	}
}

func TestRejectIllegalTalentRows(t *testing.T) {
	invalid := map[string]string{
		"balance":    "5532220115301341-05-005003",
		"feral":      "020022-5500002123032213051-055",
		"survival":   "-005355000050305-500200031000020151",
		"affliction": "2535002013521105--0540005002",
		"ds_ruin":    "25220010135201-0025003001-0540005002",
	}
	for _, b := range builds() {
		if s, ok := invalid[b.Key]; ok && loadTalents(b).validate(s) == nil {
			t.Errorf("accepted illegal %s", b.Key)
		}
	}
}

func TestHitCapAndTaurenRacial(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "enhancement" && b.Key != "fury" {
			continue
		}
		var additions []hitAdjustment
		for _, race := range []proto.Race{proto.Race_RaceOrc, proto.Race_RaceTauren} {
			p := b.player(race)
			normalized, hit, err := capHit(b, p)
			if err != nil {
				t.Fatal(err)
			}
			req := request(normalized, 1, 1)
			_, rs, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
			final := rs.Parties[0].Players[0].FinalStats.Stats
			for stat, want := range map[stats.Stat]float64{stats.MeleeHit: hit.MeleeFinal, stats.SpellHit: hit.SpellFinal} {
				if math.Abs(final[stat]-want) > 1e-8 {
					t.Errorf("%s/%s stat %v = %v, want %v", b.Key, race, stat, final[stat], want)
				}
			}
			_, again, err := capHit(b, p)
			if err != nil || !reflect.DeepEqual(hit, again) {
				t.Errorf("hit normalization accumulated: %v -> %v", hit, again)
			}
			if b.Key == "fury" && math.Abs(hit.MeleeFinal-9) > 1e-8 {
				t.Errorf("Fury should pay for the physical special cap, not unused magic: %+v", hit)
			}
			if b.Key == "enhancement" && math.Abs(hit.SpellFinal-16) > 1e-8 {
				t.Errorf("Enhancement should cap its damaging magic: %+v", hit)
			}
			additions = append(additions, hit)
		}
		if additions[0].MeleeAdded-additions[1].MeleeAdded != 1 || additions[0].SpellAdded-additions[1].SpellAdded != 1 {
			t.Errorf("%s: Tauren racial not subtracted: %v", b.Key, additions)
		}
		if math.Abs(additions[0].RawHitDelta-additions[1].RawHitDelta-10) > 1e-8 {
			t.Errorf("%s: Tauren should retain ten offensive budget points", b.Key)
		}
	}
}

func TestVendorDatabaseRegression(t *testing.T) {
	cleaver := core.ItemsByID[272592]
	if cleaver.Stats[stats.MeleeCrit]+cleaver.Stats[stats.SpellCrit] != 1 {
		t.Fatal("cleaver must give 1% crit")
	}
	for _, skill := range cleaver.WeaponSkills {
		if skill != 0 {
			t.Fatal("cleaver has fabricated weapon skill")
		}
	}
	if core.ItemsByID[272683].Stats[stats.SpellPower] != 94 {
		t.Fatal("spellblade missing spell power")
	}
	crossbow := core.ItemsByID[272595]
	if crossbow.WeaponDamageMin != 85 || crossbow.WeaponDamageMax != 129 || crossbow.Stats[stats.RangedAttackPower] != 32 {
		t.Fatal("crossbow differs from vendor export")
	}
}

func TestProfileEquipmentIsLegal(t *testing.T) {
	for _, b := range builds() {
		p := b.player(b.races()[0])
		if err := validateGear(p); err != nil {
			t.Errorf("%s: %v", b.Key, err)
		}
		if b.Key == "balance" {
			p.Equipment.Items[proto.ItemSlot_ItemSlotMainHand].Id = 272603
			if validateGear(p) == nil {
				t.Error("two-handed staff was allowed with an off-hand tome")
			}
		}
	}
}

func TestExternalShamanBuffsStayPermanent(t *testing.T) {
	for _, b := range builds() {
		if b.Class != proto.Class_ClassShaman {
			continue
		}
		req := request(b.player(proto.Race_RaceOrc), 1, 1)
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		unit := env.Raid.AllPlayerUnits[0]
		for _, label := range []string{"Strength of Earth Totem", "Grace of Air Totem"} {
			aura := unit.GetAura(label)
			if aura == nil || aura.Duration != core.NeverExpires {
				t.Errorf("%s: external %s was replaced by a temporary self-buff", b.Key, label)
			}
		}
	}
}

func TestForeverCurseCoversAllMagic(t *testing.T) {
	b := builds()[0]
	req := request(b.player(b.races()[0]), 1, 1)
	req.Raid.Debuffs = &proto.Debuffs{CurseOfElements: true, CurseOfShadow: true}
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	target := env.Encounter.TargetUnits[0]
	if target.GetAura("Curse of Shadow") != nil {
		t.Fatal("removed Curse of Shadow was applied")
	}
	aura := target.GetAura("Curse of Elements")
	if aura == nil || aura.ActionID.SpellID != 1311680 {
		t.Fatal("missing Forever Curse of the Elements")
	}
	sim := &core.Simulation{Environment: env}
	before := target.GetStats()
	aura.Activate(sim)
	for school := stats.SchoolIndexArcane; school < stats.SchoolLen; school++ {
		if got := target.PseudoStats.SchoolDamageTakenMultiplier[school]; math.Abs(got-1.1) > 1e-9 {
			t.Errorf("school %v multiplier = %v", school, got)
		}
	}
	for _, stat := range []stats.Stat{stats.ArcaneResistance, stats.FireResistance, stats.FrostResistance, stats.NatureResistance, stats.ShadowResistance} {
		if got := target.GetStat(stat) - before[stat]; got != -75 {
			t.Errorf("resistance %v changed by %v, want -75", stat, got)
		}
	}
	aura.Deactivate(sim)
	if target.GetStats() != before {
		t.Fatal("curse resistance changes did not reverse")
	}
}

func TestForeverRecklessnessHasNoAttackPower(t *testing.T) {
	b := builds()[0]
	req := request(b.player(b.races()[0]), 1, 1)
	req.Raid.Debuffs = &proto.Debuffs{CurseOfRecklessness: true}
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	target := env.Encounter.TargetUnits[0]
	aura := core.CurseOfRecklessnessAura(target)
	before := target.GetStats()
	sim := &core.Simulation{Environment: env}
	aura.Activate(sim)
	if delta := target.GetStat(stats.Armor) - before[stats.Armor]; delta != -505 {
		t.Errorf("armor changed by %v, want -505", delta)
	}
	if target.GetStat(stats.AttackPower) != before[stats.AttackPower] {
		t.Fatal("Curse of Recklessness granted attack power")
	}
	aura.Deactivate(sim)
	if target.GetStats() != before {
		t.Fatal("Curse of Recklessness did not restore the original stats")
	}
}

func TestForeverBossArmorShred(t *testing.T) {
	b := builds()[0]
	for _, startingArmor := range []float64{4638, 3731, 3009} {
		req := request(b.player(b.races()[0]), 1, 1)
		req.Encounter.Targets[0].Stats = stats.Stats{stats.Armor: startingArmor}.ToFloatArray()
		req.Raid.Debuffs = &proto.Debuffs{
			SunderArmor: true, ExposeArmor: proto.TristateEffect_TristateEffectImproved,
			CurseOfRecklessness: true, FaerieFire: true,
		}
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		target := env.Encounter.TargetUnits[0]
		sim := &core.Simulation{Environment: env}
		sunder := target.GetAura("Sunder Armor")
		sunder.Activate(sim)
		sunder.SetStacks(sim, 5)
		target.GetAura("Curse of Recklessness").Activate(sim)
		target.GetAura("Faerie Fire").Activate(sim)
		want := startingArmor - 3260 // 2250 major + 505 curse + 505 Faerie Fire.
		if got := target.GetStat(stats.Armor); got != want {
			t.Fatalf("boss armor %v: got %v, want %v", startingArmor, got, want)
		}
		if got := target.Armor(); got != math.Max(want, 0) {
			t.Fatalf("effective armor %v: got %v", startingArmor, got)
		}
		// Expose Armor replaces Sunder; its talent no longer multiplies shred.
		target.GetAura("ExposeArmor").Activate(sim)
		if got := target.GetStat(stats.Armor); got != want {
			t.Fatalf("Expose stacked with Sunder or used Classic scaling: %v, want %v", got, want)
		}
	}
}

func TestHunterPetKeepsSelectedAttackSpeed(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "beast_mastery" {
			continue
		}
		for _, tc := range []struct {
			option proto.Hunter_Options_PetAttackSpeed
			speed  float64
		}{
			{proto.Hunter_Options_One, 1},
			{proto.Hunter_Options_OneTwo, 1.2},
			{proto.Hunter_Options_Two, 2},
		} {
			p := b.player(proto.Race_RaceOrc)
			p.GetHunter().Options.PetAttackSpeed = tc.option
			req := request(p, 1, 1)
			env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
			pet := env.Raid.Parties[0].Players[0].GetCharacter().Pets[0]
			if got := pet.AutoAttacks.MH().SwingSpeed; got != tc.speed {
				t.Errorf("pet swing speed = %v, want %v", got, tc.speed)
			}
		}
		return
	}
	t.Fatal("missing Beast Mastery fixture")
}

func TestSacrificedPetDoesNotAttack(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "ds_ruin" {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		// Melee pets expose a deferred auto-swing event that an Imp does not.
		p.Rotation = core.APLRotationFromJsonString(`{"type":"TypeAPL","prepullActions":[
			{"action":{"castSpell":{"spellId":{"spellId":712}}},"doAtValue":{"const":{"val":"-16s"}}},
			{"action":{"castSpell":{"spellId":{"spellId":18788}}},"doAtValue":{"const":{"val":"-5s"}}}
		],"priorityList":[{"action":{"castSpell":{"spellId":{"spellId":25307}}}}]}`)
		req := request(p, 2, 1)
		req.Encounter.Duration = 30
		result := core.RunRaidSim(req)
		if result.Error != nil {
			t.Fatal(result.Error.Message)
		}
		for _, pet := range result.RaidMetrics.Parties[0].Players[0].Pets {
			if pet.Dps.Avg != 0 {
				t.Errorf("sacrificed/dismissed pet %s still did %v DPS", pet.Name, pet.Dps.Avg)
			}
		}
	}
}

func TestChannelAPLInterruptsWrack(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "affliction" {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		p.TalentsString = "25350020135201051--05000551"
		for _, clipped := range []bool{false, true} {
			if clipped {
				p.Rotation = core.APLRotationFromJsonString(`{"type":"TypeAPL","priorityList":[{"action":{"channelSpell":{
					"spellId":{"spellId":11704},"allowRecast":true,
					"interruptIf":{"cmp":{"lhs":{"spellChanneledTicks":{"spellId":{"spellId":11704}}},"op":"OpGe","rhs":{"const":{"val":"2"}}}}
				}}}]}`)
			} else {
				p.Rotation = core.APLRotationFromJsonString(`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":11704}}}}]}`)
			}
			req := request(p, 20, 1)
			req.Encounter.Duration = 60
			result := core.RunRaidSim(req)
			if result.Error != nil {
				t.Fatal(result.Error.Message)
			}
			found := false
			for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
				if action.Id.GetSpellId() != 11704 {
					continue
				}
				found = true
				metrics := action.Targets[0]
				seconds := metrics.CastTimeMs / float64(metrics.Casts) / 1000
				if clipped && (seconds < 1.5 || seconds > 2.05) {
					t.Errorf("two-tick channels averaged %.3f seconds", seconds)
				}
				if !clipped && seconds < 4 {
					t.Errorf("ordinary channels were unexpectedly clipped: %.3f seconds", seconds)
				}
			}
			if !found {
				t.Fatal("Wrack was never cast")
			}
		}
	}
}

func TestChannelInterruptConditionCanCheckSpellReadiness(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "affliction" {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		p.TalentsString = "25350020135201051--05000551"
		p.Rotation = core.APLRotationFromJsonString(`{"type":"TypeAPL","priorityList":[
			{"action":{"condition":{"cmp":{"lhs":{"currentTime":{}},"op":"OpGe","rhs":{"const":{"val":"3s"}}}},
				"castSpell":{"spellId":{"spellId":17923}}}},
			{"action":{"channelSpell":{"spellId":{"spellId":11704},
				"interruptIf":{"spellCanCast":{"spellId":{"spellId":17923}}}}}}
		]}`)
		req := request(p, 1, 1)
		req.Encounter.Duration = 5
		result := core.RunRaidSim(req)
		if result.Error != nil {
			t.Fatal(result.Error.Message)
		}
		for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
			if action.Id.GetSpellId() == 17923 && action.Targets[0].Casts == 1 {
				return
			}
		}
		t.Fatal("spellCanCast in interruptIf must evaluate readiness after cancelling the channel")
	}
}

func TestWrackDebuffEndsWithChannel(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "affliction" {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		p.TalentsString = "25350020135201051--05000551"
		req := request(p, 1, 1)
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		unit, target := env.Raid.AllPlayerUnits[0], env.Encounter.TargetUnits[0]
		wrack := unit.GetSpell(core.ActionID{SpellID: 11704})
		corruption := unit.GetSpell(core.ActionID{SpellID: 25311})
		bolt := unit.GetSpell(core.ActionID{SpellID: 25307})
		sim := &core.Simulation{Environment: env}
		modified := func(spell *core.Spell) float64 {
			result := &core.SpellResult{Target: target, Damage: 100}
			for _, modifier := range target.DynamicDamageTakenModifiers {
				modifier(sim, spell, result)
			}
			return result.Damage
		}
		wrack.SpellMetrics = make([]core.SpellMetrics, len(env.AllUnits))
		// Isolate the damage modifier; channel scheduling is tested above.
		wrack.Dot(target).Aura.OnGain = nil
		wrack.Dot(target).Aura.Activate(sim)
		for spell, want := range map[*core.Spell]float64{corruption: 110, wrack: 100, bolt: 100} {
			if got := modified(spell); math.Abs(got-want) > 1e-9 {
				t.Errorf("%v damage = %v, want %v", spell.ActionID, got, want)
			}
		}
		wrack.Dot(target).Cancel(sim)
		if modified(corruption) != 100 {
			t.Fatal("Wrack's damage bonus persisted after cancellation")
		}
	}
}

func TestForeverAspectOfTheBeast(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "beast_mastery" {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		p.TalentsString = "5320001505101251-00531510005"
		req := request(p, 1, 1)
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		unit := env.Raid.AllPlayerUnits[0]
		sim := &core.Simulation{Environment: env}
		before := unit.GetStats()
		hawk, beast := unit.GetAura("Aspect of the Hawk7"), unit.GetAura("Aspect of the Beast4")
		hawk.Activate(sim)
		beast.Activate(sim)
		if hawk.IsActive() || !beast.IsActive() {
			t.Fatal("hunter aspects were not mutually exclusive")
		}
		want := unit.ApplyStatDependencies(stats.Stats{stats.AttackPower: 110})[stats.AttackPower]
		if got := unit.GetStat(stats.AttackPower) - before[stats.AttackPower]; math.Abs(got-want) > 1e-8 {
			t.Errorf("Beast AP = %v, want %v", got, want)
		}
		if got := unit.GetStat(stats.RangedAttackPower); math.Abs(got-before[stats.RangedAttackPower]) > 1e-8 {
			t.Fatal("Hawk's ranged AP remained active with Beast")
		}
		melee, ranged := unit.PseudoStats.MeleeSpeedMultiplier, unit.PseudoStats.RangedSpeedMultiplier
		quick := unit.GetAura("Quick Strikes")
		quick.Activate(sim)
		if math.Abs(unit.PseudoStats.MeleeSpeedMultiplier/melee-1.3) > 1e-8 || unit.PseudoStats.RangedSpeedMultiplier != ranged {
			t.Fatal("Quick Strikes must grant only 30% melee haste")
		}
		quick.Deactivate(sim)
		beast.Deactivate(sim)
		if math.Abs(unit.PseudoStats.MeleeSpeedMultiplier-melee) > 1e-8 || unit.GetStats() != before {
			t.Fatal("aspect or haste changes did not reverse")
		}
	}
}

func TestForeverBuffValues(t *testing.T) {
	var b build
	for _, candidate := range builds() {
		if candidate.Key == "arcane" {
			b = candidate
		}
	}
	type buffSetter func(*proto.RaidSimRequest, proto.TristateEffect)
	snapshot := func(set buffSetter, value proto.TristateEffect) ([]float64, *core.Unit) {
		req := request(b.player(proto.Race_RaceUndead), 1, 1)
		req.Raid.Buffs = &proto.RaidBuffs{}
		req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
		req.Raid.Parties[0].Players[0].Buffs = &proto.IndividualBuffs{}
		req.Raid.Debuffs = &proto.Debuffs{}
		if set != nil {
			set(req, value)
		}
		env, rs, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		return rs.Parties[0].Players[0].FinalStats.Stats, env.Raid.AllPlayerUnits[0]
	}
	baseline, unit := snapshot(nil, 0)
	cases := []struct {
		name string
		set  buffSetter
		want stats.Stats
	}{
		{"Mark of the Wild", func(r *proto.RaidSimRequest, v proto.TristateEffect) { r.Raid.Buffs.GiftOfTheWild = v },
			stats.Stats{stats.BonusArmor: 385, stats.Strength: 16, stats.Agility: 16, stats.Stamina: 16, stats.Intellect: 16, stats.Spirit: 16,
				stats.ArcaneResistance: 27, stats.FireResistance: 27, stats.FrostResistance: 27, stats.NatureResistance: 27, stats.ShadowResistance: 27}},
		{"Fortitude", func(r *proto.RaidSimRequest, v proto.TristateEffect) { r.Raid.Buffs.PowerWordFortitude = v }, stats.Stats{stats.Stamina: 70}},
		{"Strength of Earth", func(r *proto.RaidSimRequest, v proto.TristateEffect) { r.Raid.Buffs.StrengthOfEarthTotem = v }, stats.Stats{stats.Strength: 53}},
		{"Grace of Air", func(r *proto.RaidSimRequest, v proto.TristateEffect) { r.Raid.Buffs.GraceOfAirTotem = v }, stats.Stats{stats.Agility: 89}},
		{"Battle Shout", func(r *proto.RaidSimRequest, v proto.TristateEffect) { r.Raid.Buffs.BattleShout = v }, stats.Stats{stats.AttackPower: 139}},
		{"Might", func(r *proto.RaidSimRequest, v proto.TristateEffect) {
			r.Raid.Parties[0].Players[0].Buffs.BlessingOfMight = v
		}, stats.Stats{stats.AttackPower: 133}},
		{"Wisdom", func(r *proto.RaidSimRequest, v proto.TristateEffect) {
			r.Raid.Parties[0].Players[0].Buffs.BlessingOfWisdom = v
		}, stats.Stats{stats.MP5: 40}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			want := unit.ApplyStatDependencies(tc.want)
			for _, mode := range []proto.TristateEffect{proto.TristateEffect_TristateEffectRegular, proto.TristateEffect_TristateEffectImproved} {
				got, _ := snapshot(tc.set, mode)
				for stat, bonus := range want {
					if math.Abs(got[stat]-baseline[stat]-bonus) > 1e-7 {
						t.Errorf("mode %v stat %d: bonus %v, want %v", mode, stat, got[stat]-baseline[stat], bonus)
					}
				}
			}
		})
	}
}

func TestForeverExposeArmor(t *testing.T) {
	b := builds()[0]
	req := request(b.player(b.races()[0]), 1, 1)
	req.Raid.Debuffs.ExposeArmor = proto.TristateEffect_TristateEffectImproved
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	target := env.Encounter.TargetUnits[0]
	aura := target.GetAura("ExposeArmor")
	sim := &core.Simulation{Environment: env}
	before := target.GetStat(stats.Armor)
	aura.Activate(sim)
	if target.GetStat(stats.Armor) != before-2250 {
		t.Fatal("Expose Armor must remove 2250 armor at five combo points")
	}
	aura.Deactivate(sim)
	if target.GetStat(stats.Armor) != before {
		t.Fatal("Expose Armor did not reverse")
	}
}
