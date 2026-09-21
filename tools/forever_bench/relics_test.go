//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/paladin"
)

func relicTestSim(t *testing.T, key string, itemID int32, tier bool) *core.Simulation {
	t.Helper()
	for _, b := range builds() {
		if b.Key != key {
			continue
		}
		p := b.player(b.races()[0])
		p.Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
		p.ForeverTier1Bonuses = tier
		p.Equipment.Items[16] = &proto.ItemSpec{Id: itemID}
		sim := core.NewSim(request(p, 1, 1), simsignals.Signals{})
		sim.Reset()
		return sim
	}
	t.Fatalf("unknown build %s", key)
	return nil
}

func TestForeverRelicCooldownsAndDurations(t *testing.T) {
	sim := relicTestSim(t, "feral", 272427, true)
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	// Thirty seconds, minus three from Tier 1 and three from the idol.
	if got := c.GetSpell(core.ActionID{SpellID: 9846}).CD.Duration; got != 24*time.Second {
		t.Fatalf("Tiger's Fury cooldown: %s", got)
	}
	for _, tier := range []bool{false, true} {
		base := relicTestSim(t, "balance", 0, tier)
		baseDot := base.Raid.Parties[0].Players[0].GetCharacter().GetSpell(core.ActionID{SpellID: 24977}).Dot(base.Encounter.TargetUnits[0])
		sim = relicTestSim(t, "balance", 272430, tier)
		c = sim.Raid.Parties[0].Players[0].GetCharacter()
		dot := c.GetSpell(core.ActionID{SpellID: 24977}).Dot(sim.Encounter.TargetUnits[0])
		if dot.NumberOfTicks != baseDot.NumberOfTicks+1 || dot.Duration != baseDot.Duration+2*time.Second ||
			dot.OriginalNumberOfTicks != baseDot.OriginalNumberOfTicks+1 {
			t.Fatalf("Insect Swarm tier=%t: %d ticks, %s", tier, dot.NumberOfTicks, dot.Duration)
		}
	}
	sim = relicTestSim(t, "elemental", 272433, true)
	c = sim.Raid.Parties[0].Players[0].GetCharacter()
	dot := c.GetSpell(core.ActionID{SpellID: 29228}).Dot(sim.Encounter.TargetUnits[0])
	if dot.NumberOfTicks != 5 || dot.Duration != 15*time.Second {
		t.Fatalf("Flame Shock: %d ticks, %s", dot.NumberOfTicks, dot.Duration)
	}
}

func TestLibramOfLawAffectsBothJudgements(t *testing.T) {
	base := relicTestSim(t, "retribution", 0, true).Raid.Parties[0].Players[0].GetCharacter()
	equipped := relicTestSim(t, "retribution", 272435, true).Raid.Parties[0].Players[0].GetCharacter()
	checked := 0
	for _, spell := range base.Spellbook {
		if spell.SpellCode != paladin.SpellCode_PaladinJudgementOfCommand && spell.SpellCode != paladin.SpellCode_PaladinJudgementOfRighteousness {
			continue
		}
		got := equipped.GetSpell(spell.ActionID).DamageMultiplier / spell.DamageMultiplier
		if math.Abs(got-1.04) > 1e-9 {
			t.Fatalf("%s multiplier = %.8f", spell.ActionID, got)
		}
		checked++
	}
	if checked < 2 {
		t.Fatal("judgement spells not registered")
	}
}
