//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestPhysicalRetIdentityAndGearPool(t *testing.T) {
	var b build
	for _, candidate := range builds() {
		if candidate.Key == "retribution_physical" {
			b = candidate
			break
		}
	}
	if b.Key == "" {
		t.Fatal("missing physical Ret build")
	}
	p := b.player(proto.Race_RaceHuman)
	if err := loadTalents(b).validate(p.TalentsString); err != nil {
		t.Fatal(err)
	}
	if p.Consumes.Flask != 0 || p.Consumes.SpellPowerBuff != 0 || p.Consumes.FirePowerBuff != 0 {
		t.Fatal("physical Ret started with caster-focused consumables")
	}
	pool, excluded := comparisonGearPool(b, p)
	for _, item := range pool {
		if !physicalRetItemEligible(item) {
			t.Fatalf("caster item %d admitted to physical Ret's pool", item.ID)
		}
	}
	if excluded[19682] == "" || !physicalRetItemEligible(core.ItemsByID[252484]) {
		t.Fatal("Bloodvine caster chest or Timbermaw physical chest was classified incorrectly")
	}
	for _, e := range benchmarkEnchants() {
		power := stats.FromFloatArray(e.Stats)[stats.SpellPower]
		if power > 0 && physicalRetEnchantEligible(e) {
			t.Fatalf("spell-power enchant %d admitted to physical Ret", e.EffectId)
		}
	}
}
