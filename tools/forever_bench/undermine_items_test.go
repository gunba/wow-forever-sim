//go:build with_db

package main

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestWeaknessAnalyzerConsumptionAndEquipmentLimit(t *testing.T) {
	b := builds()[0]
	p := b.player(b.races()[0])
	p.Rotation = &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	p.Equipment.Items[12] = &proto.ItemSpec{Id: 272438}
	if err := validateGear(p); err != nil {
		t.Fatal(err)
	}
	p.Equipment.Items[13] = &proto.ItemSpec{Id: 272439}
	if err := validateGear(p); err == nil || !strings.Contains(err.Error(), "shared equipment category") {
		t.Fatalf("two Undermine trinkets accepted: %v", err)
	}
	p.Equipment.Items[13] = b.player(b.races()[0]).Equipment.Items[13]
	sim := core.NewSim(request(p, 1, 1), simsignals.Signals{})
	sim.Reset()
	c := sim.Raid.Parties[0].Players[0].GetCharacter()
	c.AutoAttacks.CancelAutoSwing(sim)
	aura := c.GetAura("ItemActive-272438")
	use := c.GetSpell(core.ActionID{ItemID: 272438})
	if use.CD.Duration != 90*time.Second || use.SharedCD.Duration != 20*time.Second {
		t.Fatal("incorrect item or shared cooldown")
	}
	before := c.GetStat(stats.MeleeCrit)
	if !use.Cast(sim, &c.Unit) || !aura.IsActive() || math.Abs(c.GetStat(stats.MeleeCrit)-before-5) > 1e-9 {
		t.Fatal("critical chance buff not applied")
	}
	if aura.OnPeriodicDamageDealt != nil {
		t.Fatal("periodic criticals must not consume the effect")
	}
	aura.OnSpellHitDealt(aura, sim, use, &core.SpellResult{Outcome: core.OutcomeHit})
	if !aura.IsActive() {
		t.Fatal("ordinary hit consumed effect")
	}
	aura.OnSpellHitDealt(aura, sim, use, &core.SpellResult{Outcome: core.OutcomeCrit})
	if aura.IsActive() || math.Abs(c.GetStat(stats.MeleeCrit)-before) > 1e-9 {
		t.Fatal("direct critical did not remove effect cleanly")
	}
}
