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
	"github.com/wowsims/classic/sim/druid"
	"github.com/wowsims/classic/sim/hunter"
	"github.com/wowsims/classic/sim/paladin"
	"github.com/wowsims/classic/sim/rogue"
	"github.com/wowsims/classic/sim/warlock"
	googleProto "google.golang.org/protobuf/proto"
)

func TestEvidenceInstantPoisonAttackPower(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		sample := func(ap float64) float64 {
			req := racialFixture("combat", proto.Race_RaceOrc)
			req.SimOptions.Ruleset = ruleset
			req.Raid.Buffs = nil
			req.Raid.Debuffs = nil
			req.Raid.Parties[0].Buffs = nil
			p := req.Raid.Parties[0].Players[0]
			p.Rotation = &proto.APLRotation{}
			p.TalentsString = ""
			p.Equipment = &proto.EquipmentSpec{}
			p.GetRogue().Options = &proto.RogueOptions{}
			p.Consumes = nil
			p.Buffs = nil
			p.ForeverTier1Bonuses = false
			p.BonusStats = &proto.UnitStats{Stats: stats.Stats{stats.AttackPower: ap, stats.SpellHit: 1000}.ToFloatArray()}
			req.Encounter.Targets[0].Level = 60
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			r := sim.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
			spell := r.InstantPoison
			spell.BonusCritRating = -100000
			spell.ApplyEffects(sim, sim.Encounter.TargetUnits[0], spell)
			damage := spell.SpellMetrics[sim.Encounter.TargetUnits[0].UnitIndex].TotalDamage
			if damage <= 0 {
				t.Fatal("poison sample did not land")
			}
			return damage
		}
		delta := sample(2000) - sample(0)
		want := 0.0
		if ruleset == proto.Ruleset_RulesetForever {
			want = 10
		}
		if math.Abs(delta-want) > 1e-7 {
			t.Errorf("ruleset %v: poison AP delta %v, want %v", ruleset, delta, want)
		}
	}
}

func TestEvidenceMinorArmorExclusivity(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		for _, curseFirst := range []bool{false, true} {
			req := racialFixture("fury", proto.Race_RaceOrc)
			req.SimOptions.Ruleset = ruleset
			req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
			req.Raid.Debuffs = &proto.Debuffs{FaerieFire: true, CurseOfRecklessness: true}
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			target := sim.Encounter.TargetUnits[0]
			ff, curse := target.GetAura("Faerie Fire"), target.GetAura("Curse of Recklessness")
			ff.Deactivate(sim)
			curse.Deactivate(sim)
			base, baseAP := target.GetStat(stats.Armor), target.GetStat(stats.AttackPower)
			first, second := ff, curse
			if curseFirst {
				first, second = curse, ff
			}
			first.Activate(sim)
			second.Activate(sim)
			wantReduction, wantAP := 505.0, baseAP
			if ruleset == proto.Ruleset_RulesetClassic {
				wantReduction, wantAP = 1010, baseAP+90
			}
			if math.Abs(target.GetStat(stats.Armor)-(base-wantReduction)) > 1e-9 || target.GetStat(stats.AttackPower) != wantAP {
				t.Errorf("ruleset %v, curse first %v: armor %v AP %v, want armor %v AP %v", ruleset, curseFirst, target.GetStat(stats.Armor), target.GetStat(stats.AttackPower), base-wantReduction, wantAP)
			}
			first.Deactivate(sim)
			if math.Abs(target.GetStat(stats.Armor)-(base-505)) > 1e-9 {
				t.Error("remaining reduction was lost or applied twice")
			}
			second.Deactivate(sim)
			if target.GetStat(stats.Armor) != base || target.GetStat(stats.AttackPower) != baseAP {
				t.Error("armor/AP effect leaked after both auras expired")
			}
		}
	}
}

func TestEvidenceMongooseFirstDodge(t *testing.T) {
	req := racialFixture("survival", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	h := sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
	for _, outcome := range []core.HitOutcome{core.OutcomeMiss, core.OutcomeHit, core.OutcomeDodge} {
		h.DefensiveState.Deactivate(sim)
		h.OnSpellHitTaken(sim, &core.Spell{ProcMask: core.ProcMaskMeleeMHAuto}, &core.SpellResult{Target: &h.Unit, Outcome: outcome})
		if h.DefensiveState.IsActive() != (outcome == core.OutcomeDodge) {
			t.Errorf("outcome %v: eligibility window %v", outcome, h.DefensiveState.IsActive())
		}
	}
}

func TestEvidenceMoonfireFullDamage(t *testing.T) {
	for _, extraSP := range []float64{0, 200} {
		var direct, tick [2]float64
		for arm, points := range []int{0, 2} {
			req := historyTalentFixture("balance", map[string]int{"improvedMoonfire": points})
			req.Raid.Parties[0].Players[0].ForeverTier1Bonuses = false
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			unit := sim.Raid.AllPlayerUnits[0]
			unit.AddStatDynamic(sim, stats.SpellPower, extraSP)
			spell := unit.GetSpell(core.ActionID{SpellID: 9835})
			target := unit.CurrentTarget
			direct[arm] = spell.CalcDamage(sim, target, 100, spell.OutcomeAlwaysHit).Damage
			dot := spell.Dot(target)
			dot.TakeSnapshot(sim, false)
			tick[arm] = dot.CalcSnapshotDamage(sim, target, spell.OutcomeAlwaysHit).Damage
		}
		if math.Abs(direct[1]/direct[0]-1.1) > 1e-9 || math.Abs(tick[1]/tick[0]-1.1) > 1e-9 {
			t.Errorf("extra SP %v: direct ratio %v, tick ratio %v; both must be 1.1", extraSP, direct[1]/direct[0], tick[1]/tick[0])
		}
	}
}

func TestEvidenceLacerateExistingStacks(t *testing.T) {
	req := historyTalentFixture("feral", nil)
	p := req.Raid.Parties[0].Players[0]
	p.Equipment = &proto.EquipmentSpec{}
	p.ForeverTier1Bonuses = false
	p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{Options: &proto.FeralTankDruid_Options{StartingRage: 50}}}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	bear := sim.Raid.Parties[0].Players[0].(interface{ GetDruid() *druid.Druid }).GetDruid()
	target := bear.CurrentTarget
	target.AddStatDynamic(sim, stats.Armor, -target.GetStat(stats.Armor))
	bear.AddStatDynamic(sim, stats.AttackPower, -bear.GetStat(stats.AttackPower))
	bear.AddStatDynamic(sim, stats.MeleeHit, 10000)
	bear.Lacerate.BonusCritRating = -10000
	for _, table := range bear.AttackTables[target.UnitIndex] {
		table.BaseDodgeChance, table.BaseParryChance, table.BaseBlockChance = 0, 0, 0
	}
	bear.AutoAttacks.SetMH(core.Weapon{BaseDamageMin: 100, BaseDamageMax: 100, SwingSpeed: 2.5, NormalizedSwingSpeed: 2.5, AttackPowerPerDPS: core.DefaultAttackPowerPerDPS})
	for existing := int32(0); existing <= 5; existing++ {
		sim.CurrentTime = time.Duration(existing) * 2 * time.Second
		metrics := &bear.Lacerate.SpellMetrics[target.UnitIndex]
		before := metrics.TotalDamage
		bear.Lacerate.ApplyEffects(sim, target, bear.Lacerate.Spell)
		want := 10 * float64(existing)
		if got := metrics.TotalDamage - before; math.Abs(got-want) > 1e-9 {
			t.Errorf("%d existing stacks: direct damage %v, want %v", existing, got, want)
		}
		if stacks := bear.LacerateBleed.Dot(target).GetStacks(); stacks != min(existing+1, 5) {
			t.Fatalf("application lost bleed stacks: %d", stacks)
		}
	}
	bear.LacerateBleed.Dot(target).Deactivate(sim)
	for _, table := range bear.AttackTables[target.UnitIndex] {
		table.BaseDodgeChance = 1
	}
	bear.ClearcastingAura.Deactivate(sim)
	before := bear.CurrentRage()
	if !bear.Lacerate.Cast(sim, target) {
		t.Fatal("avoided-hit control did not cast")
	}
	if got := before - bear.CurrentRage(); math.Abs(got-3) > 1e-9 || bear.LacerateBleed.Dot(target).IsActive() {
		t.Fatalf("avoided hit: spent %v Rage or added bleed; want 15 cost minus 12 refund and no bleed", got)
	}
	sim.CurrentTime += 2 * time.Second
	bear.ClearcastingAura.Activate(sim)
	before = bear.CurrentRage()
	if !bear.Lacerate.Cast(sim, target) || bear.CurrentRage() != before || bear.LacerateBleed.Dot(target).IsActive() {
		t.Fatalf("free avoided Lacerate generated Rage: %v -> %v", before, bear.CurrentRage())
	}
}

func TestEvidenceHunterPetEndurance(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		petStats := func(rank int) (float64, float64) {
			req := historyTalentFixture("beast_mastery", map[string]int{"enduranceTraining": rank})
			req.SimOptions.Ruleset = ruleset
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			pet := activePet(t, sim)
			return pet.MaxHealth(), pet.GetStat(stats.Armor) * pet.PseudoStats.ArmorMultiplier
		}
		baseHealth, baseArmor := petStats(0)
		for rank := 1; rank <= 5; rank++ {
			health, armor := petStats(rank)
			factor := 1 + .03*float64(rank)
			wantArmor := baseArmor
			if ruleset == proto.Ruleset_RulesetForever {
				wantArmor *= factor
			}
			if math.Abs(health-baseHealth*factor) > 1e-7 || math.Abs(armor-wantArmor) > 1e-7 {
				t.Errorf("%v rank %d: Health/Armor %v/%v, want %v/%v", ruleset, rank, health, armor, baseHealth*factor, wantArmor)
			}
		}
	}
}

func TestEvidenceIntimidationSpellCrit(t *testing.T) {
	req := historyTalentFixture("beast_mastery", map[string]int{"intimidation": 1})
	req.Raid.Parties[0].Players[0].GetHunter().Options.PetType = proto.Hunter_Options_WindSerpent
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	h := sim.Raid.Parties[0].Players[0].(hunter.HunterAgent).GetHunter()
	pet := activePet(t, sim)
	before := pet.GetStats()
	h.IntimidationPetAura.Activate(sim)
	for stat, bonus := range map[stats.Stat]float64{stats.MeleeCrit: 100 * core.CritRatingPerCritChance, stats.SpellCrit: 100 * core.SpellCritRatingPerCritChance} {
		if got := pet.GetStat(stat) - before[stat]; math.Abs(got-bonus) > 1e-7 {
			t.Errorf("%v bonus %v, want %v", stat, got, bonus)
		}
	}
	spell := pet.GetSpell(core.ActionID{SpellID: 25012})
	pet.OnSpellHitDealt(sim, spell, &core.SpellResult{Target: pet.CurrentTarget, Outcome: core.OutcomeMiss})
	if !h.IntimidationPetAura.IsActive() {
		t.Fatal("miss consumed next-attack bonus")
	}
	pet.AddStatDynamic(sim, stats.SpellHit, 10000)
	spell.ApplyEffects(sim, pet.CurrentTarget, spell)
	metrics := spell.SpellMetrics[pet.CurrentTarget.UnitIndex]
	if metrics.Crits+metrics.ResistedCrits != 1 || h.IntimidationPetAura.IsActive() {
		t.Fatal("successful Lightning Breath did not crit and consume Intimidation")
	}
	if math.Abs(pet.GetStat(stats.MeleeCrit)-before[stats.MeleeCrit]) > 1e-7 || math.Abs(pet.GetStat(stats.SpellCrit)-before[stats.SpellCrit]) > 1e-7 {
		t.Fatal("consumption retained crit bonus")
	}
}

func TestEvidenceAimedShotLearning(t *testing.T) {
	req := historyTalentFixture("marksmanship", nil)
	sim := core.NewSim(req, simsignals.Signals{})
	h := sim.Raid.Parties[0].Players[0].(hunter.HunterAgent).GetHunter()
	for rank, id := range []int32{19434, 20900, 20901, 20902, 20903, 20904} {
		spell := h.GetSpell(core.ActionID{SpellID: id})
		want := []int{20, 28, 36, 44, 52, 60}[rank]
		if spell == nil || spell.RequiredLevel != want || spell.Rank != rank+1 || spell.DefaultCast.CastTime != 2*time.Second {
			t.Fatalf("Aimed Shot %d has wrong learning/cast metadata", id)
		}
	}
}

func TestEvidenceOrdinaryBombsRejectAnimalForms(t *testing.T) {
	for _, form := range []string{"cat", "bear", "moonkin"} {
		for _, itemID := range []int32{10646, 18641} {
			t.Run(fmt.Sprintf("%s/%d", form, itemID), func(t *testing.T) {
				req := historyTalentFixture("feral", nil)
				if form == "moonkin" {
					req = historyTalentFixture("balance", map[string]int{"moonkinForm": 1})
				}
				p := req.Raid.Parties[0].Players[0]
				if form == "bear" {
					p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{Options: &proto.FeralTankDruid_Options{}}}
				}
				p.DistanceFromTarget = 5
				p.Profession1 = proto.Profession_Engineering
				p.ForeverTier1Bonuses = false
				p.Consumes = &proto.Consumes{SapperExplosive: proto.SapperExplosive_SapperGoblinSapper, FillerExplosive: proto.Explosive_ExplosiveDenseDynamite}
				sim := core.NewSim(req, simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				d := sim.Raid.Parties[0].Players[0].(druid.DruidAgent).GetDruid()
				spell := d.GetSpell(core.ActionID{ItemID: itemID})
				if spell == nil || !d.IsShapeshifted() {
					t.Fatal("fixture missing bomb or form")
				}
				if spell.Cast(sim, d.CurrentTarget) || !d.IsShapeshifted() {
					t.Fatal("ordinary bomb accepted or silently cancelled form")
				}
				d.CancelShapeshift(sim)
				if !spell.CanCast(sim, d.CurrentTarget) {
					t.Fatal("ordinary bomb remained blocked after explicit unshift")
				}
			})
		}
	}
}

func TestEvidenceDivinePrecisionFamily(t *testing.T) {
	for rank := 1; rank <= 3; rank++ {
		req := historyTalentFixture("retribution", map[string]int{"divinePrecision": rank, "holyShock": 1, "holyShield": 1, "sealOfCommand": 1})
		req.Raid.Parties[0].Players[0].ForeverTier1Bonuses = false
		sim := core.NewSim(req, simsignals.Signals{})
		u := sim.Raid.AllPlayerUnits[0]
		for _, id := range []int32{10333, 20924, 10314, 10318, 20930} {
			spell := u.GetSpell(core.ActionID{SpellID: id})
			if spell == nil || spell.BonusHitRating != 6*float64(rank) {
				t.Errorf("rank %d: family spell %d lacks its hit modifier", rank, id)
			}
		}
		for _, id := range []int32{24239, 20286, 20966, 20424} {
			spell := u.GetSpell(core.ActionID{SpellID: id})
			if spell == nil || spell.BonusHitRating != 0 || u.GetSchoolBonusHitChance(spell) != 0 {
				t.Errorf("rank %d: unrelated Holy spell %d received hit", rank, id)
			}
		}
		spell := u.GetSpell(core.ActionID{SpellID: 10333})
		table := u.AttackTables[sim.Encounter.TargetUnits[0].UnitIndex][spell.CastType]
		sim.Options.Interactive = true
		sim.Reset()
		u.AddStatDynamic(sim, stats.MeleeHit, 10000)
		table.BaseDodgeChance = 1
		result := spell.CalcOutcome(sim, sim.Encounter.TargetUnits[0], spell.OutcomeMeleeSpecialHitAndCrit)
		if !result.Outcome.Matches(core.OutcomeDodge) {
			t.Fatal("hit modifier removed dodge")
		}
	}
}

func TestEvidenceHolyStrikePowerTerms(t *testing.T) {
	for rank, id := range []int32{679, 678, 1866, 680, 2495, 5569, 10332, 10333} {
		damage := func(playerSP, targetBonus float64) float64 {
			req := historyTalentFixture("retribution", nil)
			p := req.Raid.Parties[0].Players[0]
			p.Equipment = &proto.EquipmentSpec{}
			p.ForeverTier1Bonuses = false
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			u := sim.Raid.AllPlayerUnits[0]
			target := sim.Encounter.TargetUnits[0]
			u.AddStatDynamic(sim, stats.AttackPower, -u.GetStat(stats.AttackPower))
			u.AddStatDynamic(sim, stats.SpellPower, playerSP-u.GetStat(stats.SpellPower))
			u.AddStatDynamic(sim, stats.SpellDamage, -u.GetStat(stats.SpellDamage))
			u.AddStatDynamic(sim, stats.HolyPower, -u.GetStat(stats.HolyPower))
			u.AddStatDynamic(sim, stats.MeleeHit, 10000)
			target.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] = targetBonus
			u.AutoAttacks.SetMH(core.Weapon{BaseDamageMin: 100, BaseDamageMax: 100, SwingSpeed: 2.4, NormalizedSwingSpeed: 2.4, AttackPowerPerDPS: core.DefaultAttackPowerPerDPS})
			spell := u.GetSpell(core.ActionID{SpellID: id})
			spell.BonusCritRating = -10000
			spell.Flags |= core.SpellFlagIgnoreResists // Partial resistance is a separate unresolved rule.
			for _, table := range u.AttackTables[target.UnitIndex] {
				table.BaseDodgeChance, table.BaseParryChance, table.BaseBlockChance = 0, 0, 0
			}
			spell.ApplyEffects(sim, target, spell)
			return spell.SpellMetrics[target.UnitIndex].TotalDamage
		}
		base := damage(0, 0)
		weaponShare := []float64{.25, .29, .32, .36, .39, .43, .46, .50}[rank]
		if got, want := damage(100, 0)-base, 42.9*weaponShare; math.Abs(got-want) > 1e-7 {
			t.Errorf("rank %d: player SP gain %v, want %v", rank+1, got, want)
		}
		if got := damage(0, 100) - base; math.Abs(got-42.9) > 1e-7 {
			t.Errorf("rank %d: target bonus gain %v, want full 42.9", rank+1, got)
		}
	}
}

func TestEvidenceFeralWeaponsAndStones(t *testing.T) {
	for _, bearForm := range []bool{false, true} {
		for _, stone := range []struct {
			imbue proto.WeaponImbue
			flat  float64
		}{{proto.WeaponImbue_WeaponImbueUnknown, 0}, {proto.WeaponImbue_SolidWeightstone, 6}, {proto.WeaponImbue_DenseWeightstone, 8}} {
			t.Run(fmt.Sprintf("bear-%v/%v", bearForm, stone.imbue), func(t *testing.T) {
				req := historyTalentFixture("feral", nil)
				p := req.Raid.Parties[0].Players[0]
				p.ForeverTier1Bonuses = false
				p.Consumes = &proto.Consumes{MainHandImbue: stone.imbue}
				p.EnableItemSwap = true
				if p.Database == nil {
					p.Database = &proto.SimDatabase{}
				}
				p.Database.Items = append(p.Database.Items, &proto.SimItem{Id: 1900000001, Name: "Form regression staff", Type: proto.ItemType_ItemTypeWeapon, WeaponType: proto.WeaponType_WeaponTypeStaff, HandType: proto.HandType_HandTypeTwoHand, WeaponDamageMin: 120, WeaponDamageMax: 180, WeaponSpeed: 3})
				p.ItemSwap = &proto.ItemSwap{MhItem: &proto.ItemSpec{Id: 1900000001}}
				if bearForm {
					p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{Options: &proto.FeralTankDruid_Options{}}}
				}
				sim := core.NewSim(req, simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				d := sim.Raid.Parties[0].Players[0].(druid.DruidAgent).GetDruid()
				interval := core.TernaryFloat64(bearForm, 2.5, 1)
				check := func() {
					t.Helper()
					item := d.GetMHWeapon()
					want := ((item.WeaponDamageMin+item.WeaponDamageMax)/2+stone.flat)/item.SwingSpeed + d.PseudoStats.BonusMHDps
					weapon := d.AutoAttacks.MH()
					if math.Abs(weapon.DPS()-want) > 1e-7 || weapon.SwingSpeed != interval || weapon.NormalizedSwingSpeed != interval {
						t.Errorf("form weapon DPS/interval %v/%v/%v, want %v/%v/%v", weapon.DPS(), weapon.SwingSpeed, weapon.NormalizedSwingSpeed, want, interval, interval)
					}
					if got := weapon.CalculateAverageWeaponDamage(140) / interval; math.Abs(got-want-10) > 1e-7 {
						t.Error("form normalization changed the AP contribution")
					}
				}
				check()
				d.PseudoStats.BonusMHDps += 5
				aura := d.CatFormAura
				if bearForm {
					aura = d.BearFormAura
				}
				aura.Deactivate(sim)
				aura.Activate(sim)
				check()
				d.RegisterOnItemSwap(func(_ *core.Simulation) {
					if d.AutoAttacks.MH().SwingSpeed != interval {
						t.Error("swap callback observed equipment speed instead of form speed")
					}
				})
				sim.CurrentTime = time.Second
				d.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand})
				check()
				aura.Deactivate(sim)
				raw := d.WeaponFromMainHand()
				if *d.AutoAttacks.MH() != raw {
					t.Error("leaving form did not restore the actual weapon")
				}
			})
		}
	}
}

func TestEvidenceFeralWeaponImbueEligibility(t *testing.T) {
	for _, tc := range []struct {
		imbue   proto.WeaponImbue
		spellID int32
	}{{proto.WeaponImbue_ElementalSharpeningStone, 0}, {proto.WeaponImbue_ShadowOil, 1382}, {proto.WeaponImbue_FrostOil, 1191}} {
		req := historyTalentFixture("feral", nil)
		p := req.Raid.Parties[0].Players[0]
		p.ForeverTier1Bonuses = false
		p.Consumes = &proto.Consumes{MainHandImbue: tc.imbue}
		sim := core.NewSim(req, simsignals.Signals{})
		if tc.spellID != 0 && sim.Raid.AllPlayerUnits[0].GetSpell(core.ActionID{SpellID: tc.spellID}) == nil {
			t.Errorf("%v proc remains excluded in form", tc.imbue)
		}
		if tc.spellID == 0 {
			baseline := historyTalentFixture("feral", nil)
			baseline.Raid.Parties[0].Players[0].ForeverTier1Bonuses = false
			base := core.NewSim(baseline, simsignals.Signals{})
			got := sim.Raid.AllPlayerUnits[0].GetStat(stats.MeleeCrit) - base.Raid.AllPlayerUnits[0].GetStat(stats.MeleeCrit)
			if math.Abs(got-2*core.CritRatingPerCritChance) > 1e-7 {
				t.Errorf("Elemental Stone crit %v, want 2%%", got)
			}
		}
	}
}

func TestEvidenceWrackFamilyAndCancellation(t *testing.T) {
	req := historyTalentFixture("affliction", map[string]int{"wrack": 1})
	req.Raid.Debuffs = &proto.Debuffs{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	w := sim.Raid.Parties[0].Players[0].(warlock.WarlockAgent).GetWarlock()
	target := sim.Encounter.TargetUnits[0]
	if w.Wrack.RequiredLevel != 40 {
		t.Error("Wrack is not available at its level-40 acquisition")
	}
	check := func(id int32, multiplier float64, otherOwner bool) {
		t.Helper()
		spell := w.GetSpell(core.ActionID{SpellID: id})
		if spell == nil {
			t.Fatalf("missing test spell %v", id)
		}
		copy := *spell
		if otherOwner {
			copy.Unit = &core.Unit{AttackTables: w.AttackTables}
		}
		result := &core.SpellResult{Target: target, Damage: 100}
		copy.ApplyPostOutcomeDamageModifiers(sim, result)
		if math.Abs(result.Damage-100*multiplier) > 1e-7 {
			t.Errorf("spell %v other=%v: Wrack modifier %v, want %v", id, otherOwner, result.Damage/100, multiplier)
		}
	}
	w.Wrack.Dot(target).Apply(sim)
	for _, id := range []int32{25311, 11713} {
		check(id, 1.1, false)
		check(id, 1, true)
	}
	for _, id := range []int32{603, 11700, 11675} {
		check(id, 1, false)
	}
	w.Wrack.Dot(target).Cancel(sim)
	check(25311, 1, false)
	check(11713, 1, false)
}

func TestEvidenceImprovedImpFullPower(t *testing.T) {
	damage := func(rank int32, sp float64) float64 {
		req := historyTalentFixture("demonology", map[string]int{"improvedImp": int(rank)})
		p := req.Raid.Parties[0].Players[0]
		p.ForeverTier1Bonuses = false
		p.GetWarlock().Options.Summon = proto.WarlockOptions_Imp
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		pet := sim.Raid.Parties[0].Players[0].(warlock.WarlockAgent).GetWarlock().Imp
		pet.AddStatDynamic(sim, stats.SpellPower, sp-pet.GetStat(stats.SpellPower))
		spell := pet.GetSpell(core.ActionID{SpellID: 11763})
		spell.BonusCritRating = -10000
		spell.Flags |= core.SpellFlagIgnoreResists
		spell.ApplyEffects(sim, sim.Encounter.TargetUnits[0], spell)
		return spell.SpellMetrics[sim.Encounter.TargetUnits[0].UnitIndex].TotalDamage
	}
	base, power := damage(0, 0), damage(0, 100)
	for rank := int32(1); rank <= 3; rank++ {
		multiplier := 1 + 0.1*float64(rank)
		if got := damage(rank, 0); math.Abs(got-base*multiplier) > 1e-7 {
			t.Errorf("rank %v did not scale base damage", rank)
		}
		if got := damage(rank, 100) - damage(rank, 0); math.Abs(got-(power-base)*multiplier) > 1e-7 {
			t.Errorf("rank %v SP term %v, want %v", rank, got, (power-base)*multiplier)
		}
	}
}

func TestEvidenceCrusaderStrengthAndHeal(t *testing.T) {
	for _, level := range []int32{60} {
		req := historyTalentFixture("fury", nil)
		p := req.Raid.Parties[0].Players[0]
		p.TalentsString, p.ForeverTier1Bonuses = "", false
		p.Buffs = &proto.IndividualBuffs{}
		req.Raid.Buffs = &proto.RaidBuffs{}
		req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
		p.Equipment.Items[14].Enchant = 1900
		p.Equipment.Items[15].Enchant = 0
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		c := sim.Raid.Parties[0].Players[0].GetCharacter()
		c.RemoveHealth(sim, 500)
		before, strength := c.CurrentHealth(), c.GetStat(stats.Strength)
		proc, buff := c.GetAura("Crusader Enchant"), c.GetAura("Crusader Enchant MH")
		if proc == nil || buff == nil {
			t.Fatal("missing Crusader fixture")
		}
		result := &core.SpellResult{Target: sim.Encounter.TargetUnits[0], Outcome: core.OutcomeHit}
		for i := 0; i < 1000 && !buff.IsActive(); i++ {
			proc.OnSpellHitDealt(proc, sim, c.AutoAttacks.MHAuto(), result)
		}
		if !buff.IsActive() {
			t.Fatal("deterministic proc fixture never triggered")
		}
		if got := c.GetStat(stats.Strength) - strength; math.Abs(got-100) > 1e-7 {
			t.Errorf("level %v: Strength %v, want 100", level, got)
		}
		if got := c.CurrentHealth() - before; got < 75 || got > 125 {
			t.Errorf("level %v: self-heal %v outside 75-125", level, got)
		}
		c.GainHealth(sim, c.MaxHealth(), c.NewHealthMetrics(core.ActionID{SpellID: 1}))
		buff.Deactivate(sim)
		for i := 0; i < 1000 && !buff.IsActive(); i++ {
			proc.OnSpellHitDealt(proc, sim, c.AutoAttacks.MHAuto(), result)
		}
		if c.CurrentHealth() != c.MaxHealth() {
			t.Error("proc at full health violated the health cap")
		}
		for _, spell := range c.Spellbook {
			for _, metrics := range spell.SpellMetrics {
				if metrics.TotalDamage != 0 {
					t.Error("Crusader self-heal contributed outgoing damage")
				}
			}
		}
	}
}

func TestEvidenceLifestealingHealthReturn(t *testing.T) {
	req := historyTalentFixture("fury", nil)
	p := req.Raid.Parties[0].Players[0]
	p.ForeverTier1Bonuses = false
	p.Equipment.Items[14].Enchant = 1898
	p.Equipment.Items[15].Enchant = 0
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	spell := c.GetSpell(core.ActionID{SpellID: 20004})
	if spell == nil {
		t.Fatal("missing Lifestealing proc")
	}
	spell.Flags |= core.SpellFlagIgnoreResists
	spell.BonusHitRating, spell.BonusCritRating = 10000, -10000
	c.RemoveHealth(sim, 100)
	before := c.CurrentHealth()
	spell.ApplyEffects(sim, c.CurrentTarget, spell)
	damage := spell.SpellMetrics[c.CurrentTarget.UnitIndex].TotalDamage
	if got := c.CurrentHealth() - before; math.Abs(got-damage) > 1e-7 || damage <= 0 {
		t.Errorf("leech healed %v from %v damage", got, damage)
	}
}

func TestEvidenceRetributionAuraBase(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		req := historyTalentFixture("tank_warrior", nil)
		req.SimOptions.Ruleset = ruleset
		req.Raid.Parties[0].Players[0].Race = proto.Race_RaceHuman
		req.Raid.Buffs = &proto.RaidBuffs{RetributionAura: proto.TristateEffect_TristateEffectRegular}
		req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
		req.Raid.Debuffs = &proto.Debuffs{}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		u := sim.Raid.AllPlayerUnits[0]
		spell := u.GetSpell(core.ActionID{SpellID: 10301})
		spell.Flags |= core.SpellFlagIgnoreModifiers | core.SpellFlagIgnoreResists
		spell.BonusHitRating = 10000
		spell.ApplyEffects(sim, u.CurrentTarget, spell)
		// 70291 source base is 20; explicit external provider defaults to 0 SP.
		want := 20.0
		if got := spell.SpellMetrics[u.CurrentTarget.UnitIndex].TotalDamage; math.Abs(got-want) > 1e-7 {
			t.Errorf("%v: Ret Aura %v, want %v", ruleset, got, want)
		}
	}
}

func TestEvidenceFlatWeaponEnchantsFollowEquipment(t *testing.T) {
	for _, build := range []string{"feral", "feral_tank_druid", "fury", "arcane"} {
		for _, effect := range []int32{250, 241, 943, 805, 1897} {
			req := historyTalentFixture(build, nil)
			p := req.Raid.Parties[0].Players[0]
			p.ForeverTier1Bonuses = false
			p.Consumes = &proto.Consumes{}
			p.Equipment.Items[proto.ItemSlot_ItemSlotMainHand].Enchant = effect
			p.EnableItemSwap = true
			if p.Database == nil {
				p.Database = &proto.SimDatabase{}
			}
			weaponType := proto.WeaponType_WeaponTypeMace
			if build == "arcane" {
				weaponType = proto.WeaponType_WeaponTypeDagger
			}
			p.Database.Items = append(p.Database.Items, &proto.SimItem{Id: 1900000001, Name: "Enchant regression weapon", Type: proto.ItemType_ItemTypeWeapon, WeaponType: weaponType, HandType: proto.HandType_HandTypeOneHand, WeaponDamageMin: 120, WeaponDamageMax: 180, WeaponSpeed: 3})
			p.ItemSwap = &proto.ItemSwap{MhItem: &proto.ItemSpec{Id: 1900000001, Enchant: 250}}
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			c := sim.Raid.Parties[0].Players[0].GetCharacter()
			check := func(bonus float64) {
				t.Helper()
				if build == "arcane" && (c.AutoAttacks.MHAuto() != nil || c.AutoAttacks.AutoSwingMelee) {
					t.Fatal("Striking enabled caster melee autos")
				}
				item := c.GetMHWeapon()
				want := ((item.WeaponDamageMin+item.WeaponDamageMax)/2+bonus)/item.SwingSpeed + c.PseudoStats.BonusMHDps
				got := c.AutoAttacks.MH().DPS()
				if build == "arcane" {
					// There is no melee handler to refresh after a caster swap;
					// the equipped weapon remains the authoritative damage source.
					weapon := c.EquippedMainHandWeapon()
					got = weapon.DPS()
				}
				if math.Abs(got-want) > 1e-7 {
					t.Errorf("%s enchant %d: %v DPS, want %v", build, effect, got, want)
				}
			}
			check(map[int32]float64{250: 1, 241: 2, 943: 3, 805: 4, 1897: 5}[effect])
			sim.CurrentTime = time.Second
			c.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand})
			check(1)
			if build == "feral" || build == "feral_tank_druid" {
				d := sim.Raid.Parties[0].Players[0].(druid.DruidAgent).GetDruid()
				aura := d.CatFormAura
				if build == "feral_tank_druid" {
					aura = d.BearFormAura
				}
				aura.Deactivate(sim)
				aura.Activate(sim)
				check(1)
			}
		}
	}
}

func TestEvidenceVindicationTargetAP(t *testing.T) {
	for rank := 1; rank <= 3; rank++ {
		req := historyTalentFixture("retribution", map[string]int{"vindication": rank})
		req.Raid.Debuffs = &proto.Debuffs{DemoralizingShout: proto.TristateEffect_TristateEffectRegular}
		p := req.Raid.Parties[0].Players[0]
		p.Rotation = &proto.APLRotation{}
		p.ForeverTier1Bonuses = false
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		c := sim.Raid.Parties[0].Players[0].GetCharacter()
		target := sim.Encounter.TargetUnits[0]
		demo := target.GetAura("DemoralizingShout")
		demo.Deactivate(sim)
		target.AddStatDynamic(sim, stats.AttackPower, 805-target.GetStat(stats.AttackPower))
		beforeAP := c.GetStat(stats.AttackPower)
		enemyWeapon := core.Weapon{BaseDamageMin: 100}
		beforeDamage := enemyWeapon.EnemyWeaponDamage(sim, 805, 0)
		trigger := c.GetAura("Vindication Talent")
		trigger.OnSpellHitDealt(trigger, sim, c.AutoAttacks.MHAuto(), &core.SpellResult{Target: target, Outcome: core.OutcomeHit, Damage: 1})
		want := 805 - float64(rank)*68
		if got := target.GetStat(stats.AttackPower); math.Abs(got-want) > 1e-7 {
			t.Errorf("rank %d: target AP %v, want %v", rank, got, want)
		}
		if got := c.GetStat(stats.AttackPower); math.Abs(got-beforeAP*(1+.01*float64(rank))) > 1e-7 {
			t.Error("target reduction changed owner AP gain")
		}
		if got := enemyWeapon.EnemyWeaponDamage(sim, target.GetStat(stats.AttackPower), 0); got >= beforeDamage {
			t.Error("Vindication did not reduce incoming weapon damage")
		}
		aura := target.GetAura(fmt.Sprintf("Vindication-%d", rank))
		if aura == nil {
			t.Fatal("missing target aura")
		}
		demo.Activate(sim)
		if got := target.GetStat(stats.AttackPower); math.Abs(got-601) > 1e-7 {
			t.Error("Vindication stacked with strongest AP reduction")
		}
		demo.Deactivate(sim)
		if got := target.GetStat(stats.AttackPower); math.Abs(got-want) > 1e-7 {
			t.Error("Vindication did not resume when stronger aura expired")
		}
		aura.Deactivate(sim)
		if got := target.GetStat(stats.AttackPower); got != 805 {
			t.Error("target AP did not restore on expiry")
		}
	}
}

func TestEvidenceConsecratedGroundTargetCap(t *testing.T) {
	for rank := 1; rank <= 2; rank++ {
		req := historyTalentFixture("retribution", map[string]int{"consecratedGround": rank})
		p := req.Raid.Parties[0].Players[0]
		p.ForeverTier1Bonuses = false
		req.Raid.Debuffs = &proto.Debuffs{}
		for len(req.Encounter.Targets) < 6 {
			req.Encounter.Targets = append(req.Encounter.Targets, googleProto.Clone(req.Encounter.Targets[0]).(*proto.Target))
		}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		c := sim.Raid.Parties[0].Players[0].GetCharacter()
		holy := c.GetSpell(core.ActionID{SpellID: 10333})
		holy.BonusCoefficient = 0
		holy.Flags |= core.SpellFlagIgnoreResists
		physical := c.AutoAttacks.MHAuto()
		baseHoly, basePhysical := make([]float64, 6), make([]float64, 6)
		for i, target := range sim.Encounter.TargetUnits {
			baseHoly[i] = holy.CalcDamage(sim, target, 100, holy.OutcomeAlwaysHit).Damage
			basePhysical[i] = physical.CalcDamage(sim, target, 100, physical.OutcomeAlwaysHit).Damage
		}
		aura := c.GetAura("Consecrated Ground")
		aura.Activate(sim)
		for i, target := range sim.Encounter.TargetUnits {
			multiplier := 1.0
			if i < 4 {
				multiplier += .05 * float64(rank)
			}
			if got := holy.CalcDamage(sim, target, 100, holy.OutcomeAlwaysHit).Damage; math.Abs(got-baseHoly[i]*multiplier) > 1e-7 {
				t.Errorf("rank %d target %d: Holy %v, want %v", rank, i, got, baseHoly[i]*multiplier)
			}
			if got := physical.CalcDamage(sim, target, 100, physical.OutcomeAlwaysHit).Damage; math.Abs(got-basePhysical[i]) > 1e-7 {
				t.Error("Ground buffed physical damage")
			}
		}
		aura.Deactivate(sim)
		for i, target := range sim.Encounter.TargetUnits {
			if got := holy.CalcDamage(sim, target, 100, holy.OutcomeAlwaysHit).Damage; math.Abs(got-baseHoly[i]) > 1e-7 {
				t.Error("Ground bonus outlasted its window")
			}
		}
	}
}

func TestEvidenceSavageStrikesFamily(t *testing.T) {
	for _, rank := range []int{1, 2} {
		req := historyTalentFixture("survival", map[string]int{"savageStrikes": rank, "laceratingStrikes": 1, "striderKick": 1})
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		h := sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
		for _, spell := range []*core.Spell{h.RaptorStrikeHit, h.MongooseBite, h.WingClip, h.LaceratingStrikes, h.StriderKick} {
			if spell == nil {
				t.Fatal("missing family ability")
			}
			if want := float64(rank) * 2 * core.CritRatingPerCritChance; spell.BonusCritRating != want {
				t.Errorf("rank%d %s crit %v, want %v", rank, spell.ActionID, spell.BonusCritRating, want)
			}
		}
		for _, spell := range []*core.Spell{h.ArcaneShot, h.MultiShot, h.ImmolationTrap} {
			if spell.BonusCritRating != 0 {
				t.Error("Savage Strikes leaked outside melee family")
			}
		}
	}
}

func TestEvidenceScreechFamily(t *testing.T) {
	for _, family := range []proto.Hunter_Options_PetType{proto.Hunter_Options_CarrionBird, proto.Hunter_Options_Bat, proto.Hunter_Options_Owl} {
		req := historyTalentFixture("beast_mastery", nil)
		req.Raid.Parties[0].Players[0].GetHunter().Options.PetType = family
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		h := sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
		pet := h.GetCharacter().PetAgents[0].GetPet()
		var screech *core.Spell
		for _, spell := range pet.Spellbook {
			if spell.SpellCode == hunter.SpellCode_HunterPetScreech {
				screech = spell
			}
		}
		if family != proto.Hunter_Options_CarrionBird {
			if screech != nil {
				t.Errorf("family %v illegally has Screech", family)
			}
			continue
		}
		if screech == nil || screech.SpellID != 24579 || screech.RequiredLevel != 56 {
			t.Fatal("Carrion Bird missing active rank-four Screech")
		}
	}
}

func TestEvidenceScreechAP(t *testing.T) {
	req := historyTalentFixture("beast_mastery", nil)
	req.Raid.Parties[0].Players[0].GetHunter().Options.PetType = proto.Hunter_Options_CarrionBird
	req.Raid.Debuffs = nil
	req.Encounter.Targets = append(req.Encounter.Targets, googleProto.Clone(req.Encounter.Targets[0]).(*proto.Target))
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	pet := sim.Raid.Parties[0].Players[0].GetCharacter().PetAgents[0].GetPet()
	var screech *core.Spell
	for _, spell := range pet.Spellbook {
		if spell.SpellCode == hunter.SpellCode_HunterPetScreech {
			screech = spell
		}
	}
	target := sim.Encounter.TargetUnits[0]
	// Isolate the two sourced effects from primary-hit avoidance randomness.
	at := pet.AttackTables[target.UnitIndex][proto.CastType_CastTypeMainHand]
	at.BaseMissChance = 0
	at.BaseDodgeChance = 0
	at.BaseParryChance = 0
	ap := []float64{target.GetStat(stats.AttackPower), sim.Encounter.TargetUnits[1].GetStat(stats.AttackPower)}
	if !screech.Cast(sim, target) {
		t.Fatal("Screech could not cast")
	}
	for i, enemy := range sim.Encounter.TargetUnits {
		if enemy.GetStat(stats.AttackPower) != ap[i]-204 {
			t.Errorf("target%d did not lose204AP", i)
		}
	}
	if screech.SpellMetrics[sim.Encounter.TargetUnits[1].UnitIndex].TotalDamage != 0 {
		t.Fatal("Screech incorrectly dealt AoE damage")
	}
	aura := target.GetAura("DemoralizingScreech-24579-204")
	if aura.Duration != 30*time.Second || screech.CD.Duration != 10*time.Second || pet.CurrentFocus() != 130 {
		t.Fatal("Screech duration/cooldown/focus mismatch")
	}
	aura.Deactivate(sim)
	if target.GetStat(stats.AttackPower) != ap[0] {
		t.Fatal("Screech AP did not expire")
	}
}

func TestEvidenceBulwarkDamageProbes(t *testing.T) {
	req := historyTalentFixture("protection_paladin", map[string]int{"templarsBulwark": 1})
	req.Raid.Tanks = []*proto.UnitReference{{Type: proto.UnitReference_Player, Index: 0}}
	req.Raid.Parties[0].Players[0].HealingModel = &proto.HealingModel{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	bulwark := c.GetSpell(core.ActionID{SpellID: 1311015})
	if bulwark == nil || !bulwark.Cast(sim, &c.Unit) {
		t.Fatal("Bulwark unavailable")
	}
	if bulwark.CurCast.Cost != 110 || bulwark.CD.TimeToReady(sim) != 5*time.Minute || c.GetAura("Templar's Bulwark").Duration != 8*time.Second {
		t.Fatal("Bulwark cost, cooldown or duration changed")
	}
	hpBefore := c.CurrentHealth()
	incoming := sim.Encounter.TargetUnits[0].RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 1}, SpellSchool: core.SpellSchoolShadow, ProcMask: core.ProcMaskSpellDamage, DamageMultiplier: 1, ThreatMultiplier: 1})
	for i := 0; i < 2; i++ {
		result := incoming.CalcDamage(sim, &c.Unit, c.MaxHealth()*2, incoming.OutcomeAlwaysHit)
		incoming.DisposeResult(result)
	}
	if !c.GetAura("Templar's Bulwark").IsActive() {
		t.Fatal("damage-only probes depleted Bulwark")
	}
	result := incoming.NewResult(&c.Unit)
	result.Outcome = core.OutcomeHit
	result.Damage = c.MaxHealth() + 100
	incoming.DealDamage(sim, result)
	if math.Abs(result.Damage-100) > 1e-7 || c.GetAura("Templar's Bulwark").IsActive() {
		t.Fatal("actual hit did not use exactly one maximum-Health pool")
	}
	if math.Abs(c.CurrentHealth()-(hpBefore-100)) > 1e-7 {
		t.Fatal("absorbed damage reached the Health bar")
	}
	metrics := bulwark.SpellMetrics[c.UnitIndex]
	if metrics.TotalShielding != c.MaxHealth() || metrics.TotalAbsorbedShielding != c.MaxHealth() || metrics.TotalDamage != 0 {
		t.Fatal("Bulwark metrics omitted absorption or credited it as damage")
	}
}

func TestEvidenceCastPushbackUsesCasterProtection(t *testing.T) {
	for _, protection := range []struct {
		name          string
		own, incoming float64
		delay         time.Duration
	}{
		{"own", 1, 0, 0}, {"incoming", 0, 1, time.Second},
	} {
		t.Run(protection.name, func(t *testing.T) {
			req := historyTalentFixture("balance", nil)
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			c := sim.Raid.Parties[0].Players[0].GetCharacter()
			own := c.GetSpell(core.ActionID{SpellID: 5176})
			own.PushbackReduction = protection.own
			c.Hardcast = core.Hardcast{Expires: 3 * time.Second, ActionID: own.ActionID, Pushback: 1}
			incoming := &core.Spell{ProcMask: core.ProcMaskMeleeMHAuto, PushbackReduction: protection.incoming}
			aura := c.GetAura("Spell Pushback")
			aura.OnSpellHitTaken(aura, sim, incoming, &core.SpellResult{Target: &c.Unit, Outcome: core.OutcomeHit, Damage: 100})
			if c.Hardcast.Expires != 3*time.Second+protection.delay {
				t.Errorf("cast expires %v, want %v", c.Hardcast.Expires, 3*time.Second+protection.delay)
			}
		})
	}
}

func TestEvidenceBarkskinForms(t *testing.T) {
	for _, key := range []string{"balance", "feral", "feral_tank_druid"} {
		req := historyTalentFixture(key, nil)
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		c := sim.Raid.Parties[0].Players[0].GetCharacter()
		bark := c.GetSpell(core.ActionID{SpellID: 22812})
		if bark == nil {
			t.Errorf("%s has no Barkskin", key)
			continue
		}
		if bark.DefaultCast.GCD != core.GCDDefault {
			t.Error("Barkskin missing client GCD")
		}
		before := c.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexPhysical]
		magic := c.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexShadow]
		swing := c.AutoAttacks.NextAttackAt()
		if !bark.Cast(sim, &c.Unit) {
			t.Fatal("Barkskin cast failed")
		}
		if c.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexPhysical] != before*.8 || c.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexShadow] != magic {
			t.Error("Barkskin mitigation not physical-only")
		}
		if c.PseudoStats.SpellPushbackReduction != 1 || c.AutoAttacks.NextAttackAt() != swing {
			t.Error("Barkskin missing pushback protection or reset a melee swing")
		}
		if key == "balance" {
			wrath := c.GetSpell(core.ActionID{SpellID: 5176})
			c.Hardcast = core.Hardcast{Expires: 3 * time.Second, ActionID: wrath.ActionID, Pushback: 1}
			incoming := &core.Spell{ProcMask: core.ProcMaskMeleeMHAuto}
			aura := c.GetAura("Spell Pushback")
			aura.OnSpellHitTaken(aura, sim, incoming, &core.SpellResult{Target: &c.Unit, Outcome: core.OutcomeHit, Damage: 100})
			if c.Hardcast.Expires != 3*time.Second {
				t.Error("Barkskin did not prevent actual cast pushback")
			}
		}
		c.GetAura("Barkskin").Deactivate(sim)
		if c.PseudoStats.SpellPushbackReduction != 0 {
			t.Error("expired Barkskin kept pushback immunity")
		}
	}
}

func TestEvidenceDefendersGrip(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		req := historyTalentFixture("tank_warrior", nil)
		req.SimOptions.Ruleset = ruleset
		p := req.Raid.Parties[0].Players[0]
		p.Equipment.Items[proto.ItemSlot_ItemSlotTrinket2] = &proto.ItemSpec{Id: 272440}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		c := sim.Raid.Parties[0].Players[0].GetCharacter()
		use := c.GetSpell(core.ActionID{ItemID: 272440})
		if ruleset == proto.Ruleset_RulesetClassic {
			if use != nil {
				t.Error("Forever item use registered in Classic")
			}
			continue
		}
		if use == nil {
			t.Fatal("Defender's Grip use unregistered")
		}
		before := c.GetStat(stats.Block)
		if !use.Cast(sim, &c.Unit) {
			t.Fatal("item use failed")
		}
		if got := c.GetStat(stats.Block) - before; got != 8*core.BlockRatingPerBlockChance {
			t.Errorf("outside-city block %v, want 8%%", got)
		}
		if use.SharedCD.Timer != c.GetOffensiveTrinketCD() {
			t.Error("Defender split client category 1141 from other stat trinkets")
		}
		if use.CD.TimeToReady(sim) != 6*time.Minute || use.SharedCD.TimeToReady(sim) != 15*time.Second {
			t.Error("client item/shared cooldown mismatch")
		}
		aura := c.GetAura("Defender's Grip Stabilizer")
		if aura == nil || aura.Duration != 15*time.Second {
			t.Fatal("wrong buff duration")
		}
		aura.Deactivate(sim)
		if got := c.GetStat(stats.Block); math.Abs(got-before) > 1e-7 {
			t.Errorf("item block %v, want restored %v", got, before)
		}
	}
}

func TestEvidencePaladinTreeMetadata(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "retribution" {
			continue
		}
		trees := loadTalents(b).Trees
		var lengths [3]int
		for i, tree := range trees {
			lengths[i] = len(tree.Talents)
		}
		if paladin.TalentTreeSizes != lengths {
			t.Errorf("native Paladin sizes %v disagree with current tree %v", paladin.TalentTreeSizes, lengths)
		}
		if n := lengths[0] + lengths[1] + lengths[2]; n != (&proto.PaladinTalents{}).ProtoReflect().Descriptor().Fields().Len() {
			t.Error("current talent fields/tree positions disagree")
		}
		return
	}
	t.Fatal("missing Paladin build")
}

func TestEvidenceLightsVigilEnemyBranch(t *testing.T) {
	for i, id := range []int32{1310911, 1311590, 1311595} {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			req := historyTalentFixture("retribution", map[string]int{"holyShock": 1, "lightsVigil": 1})
			req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
			req.Raid.Debuffs = &proto.Debuffs{}
			req.Encounter.Targets = append(req.Encounter.Targets, googleProto.Clone(req.Encounter.Targets[0]).(*proto.Target))
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			c := sim.Raid.Parties[0].Players[0].GetCharacter()
			mark := c.GetSpell(core.ActionID{SpellID: id})
			if mark == nil {
				t.Fatalf("Light's Vigil rank %d unregistered", i+1)
			}
			cost := []float64{730, 1000, 1340}[i]
			if mark.RequiredLevel != 40+10*i || mark.DefaultCast.CastTime != 1500*time.Millisecond || mark.DefaultCast.Cost != cost || mark.CD.Duration != 6*time.Second {
				t.Fatal("Vigil rank metadata does not match client")
			}
			mark.BonusHitRating = 1e9
			target, other := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
			auraID := []int32{1310910, 1311594, 1311599}[i]
			label := fmt.Sprintf("Light's Vigil-%d-%d", c.Index, i+1)
			aura, otherAura := target.GetAura(label), other.GetAura(label)
			if aura.ActionID.SpellID != auraID || aura.Duration != 30*time.Second {
				t.Error("wrong enemy mark metadata")
			}
			finishMark := func(target *core.Unit) {
				c.AddMana(sim, c.MaxMana()-c.CurrentMana(), c.NewManaMetrics(core.ActionID{OtherID: proto.OtherAction_OtherActionManaRegen}))
				if !mark.Cast(sim, target) {
					t.Fatal("Vigil cast failed")
				}
				sim.CurrentTime += mark.CurCast.CastTime
				c.Hardcast.Expires = sim.CurrentTime
				c.Hardcast.OnComplete(sim, target)
			}
			finishMark(target)
			if !aura.IsActive() {
				t.Fatal("mark not applied")
			}
			sim.CurrentTime = 10 * time.Second
			finishMark(other)
			if aura.IsActive() || !otherAura.IsActive() {
				t.Error("target change retained more than one owned mark")
			}
			shock := c.GetSpell(core.ActionID{SpellID: 20930})
			shock.BonusHitRating = 1e9
			shock.BonusCritRating = -1e9
			// A consuming Shock cannot bypass an already running cooldown.
			shock.CD.Timer.Set(sim.CurrentTime + time.Second)
			if shock.CanCast(sim, other) {
				t.Error("mark bypassed existing Shock cooldown")
			}
			sim.CurrentTime += time.Second
			mana := c.CurrentMana()
			markCooldown := mark.CD.TimeToReady(sim)
			if !shock.Cast(sim, other) {
				t.Fatal("consuming Shock unavailable")
			}
			if otherAura.IsActive() || !shock.CD.IsReady(sim) || mark.CD.TimeToReady(sim) != markCooldown {
				t.Error("bad mark consumption/cooldown state")
			}
			wantMana := mana - shock.CurCast.Cost + .75*cost
			if math.Abs(c.CurrentMana()-wantMana) > 1e-7 {
				t.Errorf("mana %v want %v", c.CurrentMana(), wantMana)
			}
			damage := c.GetSpell(core.ActionID{SpellID: []int32{1310914, 1311592, 1311598}[i]})
			if damage.BonusCoefficient != .429 || damage.SpellMetrics[other.UnitIndex].Casts != 1 {
				t.Error("sourced child not triggered once")
			}
			if damage.CanCast(sim, other) {
				t.Error("Vigil damage child can be cast without its trigger")
			}

			sim.CurrentTime += 2 * time.Second
			if !shock.Cast(sim, other) {
				t.Fatal("ordinary follow-up Shock unavailable")
			}
			if damage.SpellMetrics[other.UnitIndex].Casts != 1 || shock.CD.IsReady(sim) {
				t.Error("consumed mark retriggered or cancelled an ordinary Shock cooldown")
			}
		})
	}
}

func TestEvidencePowerInfusionRecipient(t *testing.T) {
	for _, build := range []string{"smite", "shadow"} {
		for _, recipient := range []struct {
			name  string
			ref   *proto.UnitReference
			index int
		}{
			{"default", nil, 0}, {"unassigned", &proto.UnitReference{}, -1}, {"self", &proto.UnitReference{Type: proto.UnitReference_Player, Index: 0}, 0},
			{"other", &proto.UnitReference{Type: proto.UnitReference_Player, Index: 1}, 1},
			{"invalid", &proto.UnitReference{Type: proto.UnitReference_Player, Index: 99}, -1},
			{"enemy", &proto.UnitReference{Type: proto.UnitReference_Target, Index: 0}, -1},
		} {
			t.Run(build+"/"+recipient.name, func(t *testing.T) {
				req := historyTalentFixture(build, map[string]int{"powerInfusion": 1})
				p := req.Raid.Parties[0].Players[0]
				p.Buffs = &proto.IndividualBuffs{}
				p.ForeverTier1Bonuses = false
				if build == "smite" {
					p.GetSmitePriest().Options.PowerInfusionTarget = recipient.ref
				} else {
					p.GetShadowPriest().Options.PowerInfusionTarget = recipient.ref
				}
				other := racialFixture("arcane", proto.Race_RaceHuman).Raid.Parties[0].Players[0]
				other.Buffs = &proto.IndividualBuffs{}
				other.Rotation = &proto.APLRotation{}
				req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, other)
				sim := core.NewSim(req, simsignals.Signals{})
				sim.Options.Interactive = true
				sim.Reset()
				owner := sim.Raid.Parties[0].Players[0].GetCharacter()
				pi := owner.GetSpell(core.ActionID{SpellID: 10060})
				if recipient.index < 0 {
					if pi != nil {
						t.Fatal("PI registered without a valid friendly recipient")
					}
					return
				}
				target := sim.Raid.Parties[0].Players[recipient.index].GetCharacter()
				before := []float64{}
				for _, agent := range sim.Raid.Parties[0].Players {
					before = append(before, agent.GetCharacter().PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly])
				}
				mana := owner.CurrentMana()
				if !pi.Cast(sim, owner.CurrentTarget) {
					t.Fatal("PI cast failed")
				}
				if got := mana - owner.CurrentMana(); math.Abs(got-.2*owner.BaseMana) > 1e-7 {
					t.Errorf("owner paid %v mana, want %v", got, .2*owner.BaseMana)
				}
				if pi.CD.TimeToReady(sim) != core.PowerInfusionCD {
					t.Error("PI cooldown not retained on owner")
				}
				for i, agent := range sim.Raid.Parties[0].Players {
					want := before[i]
					if i == recipient.index {
						want *= 1.2
					}
					if got := agent.GetCharacter().PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly]; math.Abs(got-want) > 1e-7 {
						t.Errorf("player %d has %v damage multiplier, want %v", i, got, want)
					}
				}
				if !target.HasActiveAuraWithTag(core.PowerInfusionAuraTag) {
					t.Error("selected recipient lacks aura")
				}
			})
		}
	}
}

func TestEvidencePowerInfusionCoverageDoesNotStack(t *testing.T) {
	req := historyTalentFixture("smite", map[string]int{"powerInfusion": 1})
	p := req.Raid.Parties[0].Players[0]
	p.Buffs = &proto.IndividualBuffs{PowerInfusions: 1}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	external := c.GetSpell(core.ActionID{SpellID: 10060, Tag: -1})
	owned := c.GetSpell(core.ActionID{SpellID: 10060})
	if !external.Cast(sim, c.CurrentTarget) {
		t.Fatal("external PI unavailable")
	}
	mana := c.CurrentMana()
	if owned.CanCast(sim, c.CurrentTarget) {
		t.Error("owned PI can overwrite active external PI")
	}
	if c.CurrentMana() != mana || !owned.CD.IsReady(sim) {
		t.Error("failed overlap consumed owned resources")
	}
	damage := c.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly]
	healing := c.PseudoStats.HealingDealtMultiplier
	// Direct activations also cannot double the same effect across providers.
	externalAura := c.GetAura("PowerInfusion-" + external.ActionID.String())
	ownedAura := c.GetAura("PowerInfusion-" + owned.ActionID.String())
	ownedAura.Activate(sim)
	if c.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly] != damage || c.PseudoStats.HealingDealtMultiplier != healing {
		t.Error("simultaneous providers stacked")
	}
	externalAura.Deactivate(sim)
	if c.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly] != damage {
		t.Error("remaining owned provider lost its damage effect")
	}
	ownedAura.Deactivate(sim)
	if math.Abs(c.PseudoStats.HealingDealtMultiplier-healing/1.2) > 1e-7 {
		t.Error("PI healing modifier did not restore")
	}
	if !owned.CanCast(sim, c.CurrentTarget) {
		t.Error("owned cooldown still blocked after external expiry")
	}
}

func TestEvidenceHunterTrapCategories(t *testing.T) {
	for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
		req := racialFixture("survival", proto.Race_RaceOrc)
		req.SimOptions.Ruleset = ruleset
		req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, ruleset, false)
		h := env.Raid.Parties[0].Players[0].(*hunter.Hunter)
		if h.ExplosiveTrap.CD.Timer != h.ImmolationTrap.CD.Timer {
			t.Error("fire traps do not share a timer")
		}
		if separate := h.ExplosiveTrap.CD.Timer != h.FreezingTrap.CD.Timer; separate != (ruleset == proto.Ruleset_RulesetForever) {
			t.Errorf("ruleset %v: fire/frost separation %v", ruleset, separate)
		}
	}
}
