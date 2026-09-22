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
)

func TestForeverOmenClientModel(t *testing.T) {
	for _, key := range []string{"feral", "balance", "bear"} {
		fixture := key
		if key == "bear" {
			fixture = "feral"
		}
		req := racialFixture(fixture, proto.Race_RaceTauren)
		p := req.Raid.Parties[0].Players[0]
		if key == "bear" {
			p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{
				Options: &proto.FeralTankDruid_Options{},
			}}
		}
		p.ForeverTier1Bonuses = false
		p.Rotation = &proto.APLRotation{}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		unit := sim.Raid.AllPlayerUnits[0]
		trigger := unit.GetAura("Omen of Clarity")
		clear := unit.GetAura("Clearcasting")
		if trigger == nil || clear == nil || clear.Duration != 15*time.Second ||
			trigger.Icd.Duration != 10*time.Second || trigger.OnPeriodicDamageDealt != nil {
			t.Fatalf("%s: missing or incorrect Omen registration", key)
		}
		var paid, wrath, free, faerieFire *core.Spell
		for _, spell := range unit.Spellbook {
			if (key == "feral" && spell.SpellCode == druid.SpellCode_DruidShred) ||
				(key == "bear" && spell.SpellCode == druid.SpellCode_DruidMaul) ||
				(key == "balance" && spell.SpellCode == druid.SpellCode_DruidStarfire) {
				paid = spell
			}
			if spell.SpellCode == druid.SpellCode_DruidWrath {
				wrath = spell
			}
			if spell.SpellCode == druid.SpellCode_DruidFaerieFire {
				faerieFire = spell
			}
			if spell.Cost == nil && spell.Flags.Matches(core.SpellFlagAPL) {
				free = spell
			}
		}
		if paid == nil || free == nil {
			t.Fatal("incomplete Druid fixture")
		}
		baseCost := paid.Cost.GetCurrentCost()
		ffCost := 0.0
		if faerieFire != nil {
			ffCost = faerieFire.Cost.GetCurrentCost()
		}
		if baseCost <= 0 {
			t.Fatal("fixture needs a normally paid action")
		}
		source := paid
		if key != "balance" {
			source = unit.AutoAttacks.MHAuto()
		}
		target := sim.Encounter.TargetUnits[0]
		sim.CurrentTime = 0
		trigger.OnSpellHitDealt(trigger, sim, source, &core.SpellResult{Target: target, Outcome: core.OutcomeMiss})
		if clear.IsActive() {
			t.Fatal("miss triggered Omen")
		}
		trigger.OnSpellHitDealt(trigger, sim, source, &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
		if !clear.IsActive() || paid.Cost.GetCurrentCost() != 0 {
			t.Fatalf("%s: proc did not waive the eligible resource cost", key)
		}
		clear.OnCastComplete(clear, sim, free)
		if faerieFire != nil {
			if faerieFire.Cost.GetCurrentCost() != ffCost {
				t.Fatal("Faerie Fire incorrectly received Clearcasting")
			}
			clear.OnCastComplete(clear, sim, faerieFire)
		}
		if wrath != nil {
			if wrath.Cost.GetCurrentCost() == 0 {
				t.Fatal("Wrath incorrectly received Clearcasting")
			}
			clear.OnCastComplete(clear, sim, wrath)
		}
		if !clear.IsActive() {
			t.Fatal("Wrath/resource-free action consumed Clearcasting")
		}
		if source == paid {
			paid.CurCast.Cost = baseCost
			clear.OnCastComplete(clear, sim, paid)
			if !clear.IsActive() {
				t.Fatal("the paid cast that triggered Omen consumed it")
			}
			sim.CurrentTime = time.Second
		}
		if key == "feral" {
			unit.SpendEnergy(sim, unit.CurrentEnergy(), unit.NewEnergyMetrics(core.ActionID{SpellID: 16870}))
			if !paid.Cast(sim, target) || unit.CurrentEnergy() != 0 {
				t.Fatal("Clearcasting did not allow a real Shred at zero Energy")
			}
		} else if key == "bear" {
			if !paid.Cast(sim, target) || paid.CurCast.Cost != 0 {
				t.Fatal("Clearcasting did not waive Maul's Rage cost")
			}
		} else {
			paid.CurCast.Cost = 0
			clear.OnCastComplete(clear, sim, paid)
		}
		if clear.IsActive() || paid.Cost.GetCurrentCost() != baseCost {
			t.Fatal("next eligible action did not consume and restore the modifier")
		}
		// The Cat action above happens at the white swing's timestamp: it must
		// consume the proc even though the aura still has its full duration.
		sim.CurrentTime = 9999 * time.Millisecond
		trigger.OnSpellHitDealt(trigger, sim, source, &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
		if clear.IsActive() {
			t.Fatal("Omen ignored its ten-second ICD")
		}
		sim.CurrentTime = 10 * time.Second
		trigger.OnSpellHitDealt(trigger, sim, source, &core.SpellResult{Target: target, Outcome: core.OutcomeHit})
		if !clear.IsActive() {
			t.Fatal("Omen did not become available after ten seconds")
		}
		clear.Deactivate(sim)
		if paid.Cost.GetCurrentCost() != baseCost {
			t.Fatal("Omen expiration accumulated a cost modifier")
		}
	}
}

func TestDruidMoonglowExcludesUtility(t *testing.T) {
	req := racialFixture("balance", proto.Race_RaceTauren)
	p := req.Raid.Parties[0].Players[0]
	p.ForeverTier1Bonuses = false
	p.Equipment = &proto.EquipmentSpec{}
	p.Rotation = &proto.APLRotation{}
	talents := &proto.DruidTalents{}
	core.FillTalentsProto(talents.ProtoReflect(), p.TalentsString, druid.TalentTreeSizes)
	sim := core.NewSim(req, simsignals.Signals{})
	unit := sim.Raid.AllPlayerUnits[0]
	paidDamage, utility := 0, 0
	for _, spell := range unit.Spellbook {
		if spell.Cost == nil {
			continue
		}
		switch spell.SpellCode {
		case druid.SpellCode_DruidWrath, druid.SpellCode_DruidStarfire,
			druid.SpellCode_DruidMoonfire, druid.SpellCode_DruidInsectSwarm, druid.SpellCode_DruidHurricane:
			paidDamage++
			expected := int32(75)
			if spell.SpellCode == druid.SpellCode_DruidWrath {
				expected -= 10 * talents.ImprovedWrath
			}
			if spell.Cost.Multiplier != expected {
				t.Fatalf("%s: Moonglow did not apply its 25%% cost reduction: %v", spell.ActionID, spell.Cost.Multiplier)
			}
		case druid.SpellCode_DruidFaerieFire:
			utility++
			if spell.Cost.Multiplier != 100 {
				t.Fatalf("%s: Moonglow discounted a nonmatching family", spell.ActionID)
			}
		}
		if spell.ActionID.SpellID == 24858 && spell.Cost.Multiplier != 100 {
			t.Fatal("Moonglow discounted Moonkin Form")
		}
	}
	if paidDamage == 0 || utility == 0 {
		t.Fatal("missing spells in Moonglow fixture")
	}
}

func TestForeverDireBearEnrageArmor(t *testing.T) {
	req := racialFixture("feral", proto.Race_RaceTauren)
	p := req.Raid.Parties[0].Players[0]
	p.Spec = &proto.Player_FeralTankDruid{FeralTankDruid: &proto.FeralTankDruid{
		Options: &proto.FeralTankDruid_Options{},
	}}
	p.Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	character := sim.Raid.Parties[0].Players[0].GetCharacter()
	before := character.BaseEquipStats()[stats.Armor]
	enrage := character.GetAura("Enrage")
	if before <= 0 || enrage == nil {
		t.Fatal("incomplete armored Bear fixture")
	}
	enrage.Activate(sim)
	if got := character.BaseEquipStats()[stats.Armor]; math.Abs(got-before*.84) > 1e-7 {
		t.Fatalf("Dire Bear armor: got %v, want %v", got, before*.84)
	}
	enrage.Deactivate(sim)
	if got := character.BaseEquipStats()[stats.Armor]; math.Abs(got-before) > 1e-7 {
		t.Fatal("Enrage expiration did not restore armor")
	}
}
