//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func TestInnervateReportsNetMana(t *testing.T) {
	req := racialFixture("balance", proto.Race_RaceTauren)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	id := core.ActionID{SpellID: 29166}
	innervate := c.GetSpell(id)
	c.SpendMana(sim, c.CurrentMana()-500, c.NewManaMetrics(core.ActionID{OtherID: proto.OtherAction_OtherActionWait}))
	ordinaryTick := c.ManaRegenPerSecondWhileCasting() * 2
	before := c.CurrentMana()
	if !innervate.Cast(sim, &c.Unit) {
		t.Fatal("Innervate failed to cast")
	}
	cost := before - c.CurrentMana()
	if math.Abs(cost-.05*c.BaseMana) > 1e-8 {
		t.Fatalf("Innervate cost = %v, want 5%% base mana", cost)
	}
	fullTick := c.ManaRegenPerSecondWhileCasting() * 2
	sim.CurrentTime = 2 * time.Second
	c.ManaTick(sim)
	var gain float64
	for _, m := range c.Metrics.ToProto().Resources {
		if m.Id.GetSpellId() == 29166 {
			gain += m.Gain
		}
	}
	if want := fullTick - ordinaryTick - cost; want <= 0 || math.Abs(gain-want) > 1e-8 {
		t.Fatalf("reported net Innervate mana = %v, want %v", gain, want)
	}
	aura := c.GetAuraByID(id)
	aura.Deactivate(sim)
	if math.Abs(c.ManaRegenPerSecondWhileCasting()*2-ordinaryTick) > 1e-8 || c.PseudoStats.FullSpiritRegenSources != 0 {
		t.Fatal("Innervate left regeneration modifiers behind")
	}
}

func TestInnervateEvocationOverlap(t *testing.T) {
	for _, expireFirst := range []string{"Innervate", "Evocation"} {
		t.Run(expireFirst, func(t *testing.T) {
			req := racialFixture("arcane", proto.Race_RaceGnome)
			p := req.Raid.Parties[0].Players[0]
			p.Rotation = &proto.APLRotation{}
			p.Buffs.Innervates = 1
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			c := sim.Raid.Parties[0].Players[0].GetCharacter()
			innervate := c.GetAuraByID(core.ActionID{SpellID: 29166, Tag: -1})
			evocation := c.GetAura("Evocation Regen")
			first, second := innervate, evocation
			if expireFirst == "Evocation" {
				first, second = evocation, innervate
			}
			baseline := c.ManaRegenPerSecondWhileCasting()
			second.Activate(sim)
			secondOnly := c.ManaRegenPerSecondWhileCasting()
			first.Activate(sim)
			first.Deactivate(sim)
			if c.PseudoStats.FullSpiritRegenSources != 1 || math.Abs(c.ManaRegenPerSecondWhileCasting()-secondOnly) > 1e-8 {
				t.Fatal("expiring one effect disabled the other's full casting regeneration")
			}
			second.Deactivate(sim)
			if c.PseudoStats.FullSpiritRegenSources != 0 || math.Abs(c.ManaRegenPerSecondWhileCasting()-baseline) > 1e-8 {
				t.Fatal("overlapping effects did not restore baseline regeneration")
			}
		})
	}
}
