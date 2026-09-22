//go:build with_db

package main

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/shaman"
)

func TestFireNovaRequiresActiveFireTotem(t *testing.T) {
	req := racialFixture("enhancement", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	s := sim.Raid.Parties[0].Players[0].(interface{ GetShaman() *shaman.Shaman }).GetShaman()
	nova := s.FireNova[5]
	target := sim.Encounter.TargetUnits[0]
	mana := s.CurrentMana()
	if nova.Cast(sim, target) || s.CurrentMana() != mana || !nova.CD.IsReady(sim) {
		t.Fatal("Fire Nova cast/charged resources without a Fire totem")
	}
	if !s.SearingTotem[6].Cast(sim, target) {
		t.Fatal("Searing Totem failed")
	}
	sim.CurrentTime = 2 * time.Second
	if !nova.Cast(sim, target) {
		t.Fatal("Fire Nova failed with an active Fire totem")
	}
	if nova.Unit != &s.Unit || nova.ThreatMultiplier <= 0 {
		t.Fatal("Fire Nova damage/threat is not attributed to the Shaman")
	}
	sim.CurrentTime = s.TotemExpirations[shaman.FireTotem] + time.Nanosecond
	if nova.CanCast(sim, target) {
		t.Fatal("expired totem pointer still permits Fire Nova")
	}
}

func TestWindfuryWeaponExcludesTotemBenefit(t *testing.T) {
	req := racialFixture("enhancement", proto.Race_RaceOrc)
	p := req.Raid.Parties[0].Players[0]
	p.Rotation = &proto.APLRotation{}
	p.Consumes.MainHandImbue = proto.WeaponImbue_WindfuryWeapon
	req.Raid.Parties[0].Buffs = &proto.PartyBuffs{WindfuryTotem: true}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	trigger, buff := c.GetAura("Windfury"), c.GetAura("Windfury Buff")
	if trigger == nil || buff == nil || c.MainHand().TempEnchant != 1669 {
		t.Fatal("missing Windfury fixture")
	}
	hit := &core.SpellResult{Target: sim.Encounter.TargetUnits[0], Outcome: core.OutcomeHit}
	for _, enchant := range []int32{283, 284, 525, 1669} {
		c.MainHand().TempEnchant = enchant
		for i := 0; i < 100; i++ {
			sim.CurrentTime += 2 * time.Second
			trigger.OnSpellHitDealt(trigger, sim, c.AutoAttacks.MHAuto(), hit)
			if buff.IsActive() {
				t.Fatal("Windfury Weapon received the Windfury Totem proc")
			}
		}
	}
	c.MainHand().TempEnchant = 0
	for i := 0; i < 100 && !buff.IsActive(); i++ {
		sim.CurrentTime += 2 * time.Second
		trigger.OnSpellHitDealt(trigger, sim, c.AutoAttacks.MHAuto(), hit)
	}
	if !buff.IsActive() {
		t.Fatal("totem remained disabled after removing the conflicting imbue")
	}
	c.MainHand().TempEnchant = 1669
	trigger.OnSpellHitDealt(trigger, sim, c.AutoAttacks.MHAuto(), hit)
	if buff.IsActive() {
		t.Fatal("conflicting imbue retained the totem's AP buff")
	}
}
