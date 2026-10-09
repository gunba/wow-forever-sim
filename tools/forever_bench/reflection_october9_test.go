//go:build with_db

package main

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"math"
	"testing"
)

func TestPaladinOctober9ThornsCasterPower(t *testing.T) {
	var druid *proto.Player
	for _, b := range builds() {
		if b.Key == "balance" {
			druid = b.player(proto.Race_RaceTauren)
			break
		}
	}
	if druid == nil {
		t.Fatal("missingDruid")
	}
	druid.ForeverTier1Bonuses = false
	druid.Consumes = nil
	druid.Buffs = nil
	sim, p := paladinOctober9Fixture(t, nil, &proto.RaidBuffs{Thorns: proto.TristateEffect_TristateEffectRegular, ThornsProvider: &proto.UnitReference{Type: proto.UnitReference_Player, Index: 1}}, druid)
	owner := sim.Raid.Parties[0].Players[1].GetCharacter()
	owner.AddStatsDynamic(sim, stats.Stats{stats.SpellPower: 100, stats.NaturePower: 50, stats.SpellDamage: 20})
	p.AddStatDynamic(sim, stats.SpellPower, 1000)
	s := p.GetSpell(core.ActionID{SpellID: 9910})
	s.BonusHitRating = 10000
	before := s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
	s.ApplyEffects(sim, p.CurrentTarget, s)
	want := 18 + .06*(owner.GetStat(stats.SpellPower)+owner.GetStat(stats.NaturePower)+owner.GetStat(stats.SpellDamage))
	if math.Abs(s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage-before-want) > 1e-8 {
		t.Fatalf("caster SP/receiverSP ownership damage%g want%g", s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage-before, want)
	}
	owner.AddStatDynamic(sim, stats.SpellPower, 100)
	before = s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
	s.ApplyEffects(sim, p.CurrentTarget, s)
	if math.Abs(s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage-before-want-6) > 1e-8 {
		t.Fatal("ownedDruidstatsdidnotupdate")
	}
}
func TestPaladinOctober9ProviderRejectsInvalid(t *testing.T) {
	for _, ref := range []*proto.UnitReference{{Type: proto.UnitReference_Player, Index: 99}, {Type: proto.UnitReference_Target}, {Type: proto.UnitReference_Self}} {
		func() {
			defer func() {
				if recover() == nil {
					t.Fatal("invalidDruidprovideraccepted")
				}
			}()
			paladinOctober9Fixture(t, nil, &proto.RaidBuffs{Thorns: proto.TristateEffect_TristateEffectRegular, ThornsProvider: ref})
		}()
	}
}
