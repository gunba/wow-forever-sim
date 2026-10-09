//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestOctober9RecoveryFormsAndSharedCooldown(t *testing.T) {
	for _, key := range []string{"fury", "feral", "feral_tank_druid"} {
		t.Run(key, func(t *testing.T) {
			enchants := map[proto.ItemSlot]int32{proto.ItemSlot_ItemSlotMainHand: 8721}
			if key == "fury" {
				enchants[proto.ItemSlot_ItemSlotOffHand] = 8721
			}
			sim, c := enchantFixture(key, enchants)
			aura := c.GetAura("Recovery")
			if aura == nil || !aura.IsActive() {
				t.Fatal("Recovery not registered/active")
			}
			c.RemoveHealth(sim, c.MaxHealth()*.5)
			hp, mana, form := c.CurrentHealth(), c.CurrentMana(), c.ActiveShapeShift
			fire := func(mask core.ProcMask, outcome core.HitOutcome) {
				aura.OnSpellHitDealt(aura, sim, &core.Spell{Unit: &c.Unit, ProcMask: mask}, &core.SpellResult{Target: c.CurrentTarget, Outcome: outcome})
			}
			for _, outcome := range []core.HitOutcome{core.OutcomeHit, core.OutcomeCrit, core.OutcomeMiss, core.OutcomeBlock} {
				fire(core.ProcMaskMeleeMHAuto, outcome)
			}
			fire(core.ProcMaskSpellDamage, core.OutcomeDodge)
			fire(core.ProcMaskRangedAuto, core.OutcomeDodge)
			if c.CurrentHealth() != hp {
				t.Fatal("Recovery triggered outside melee dodge/parry")
			}
			fire(core.ProcMaskMeleeMHAuto, core.OutcomeDodge)
			want := hp + c.MaxHealth()*.05
			if math.Abs(c.CurrentHealth()-want) > 1e-6 {
				t.Fatalf("heal %g, want5%% maxhealth %g", c.CurrentHealth()-hp, c.MaxHealth()*.05)
			}
			fire(core.ProcMaskMeleeOHSpecial, core.OutcomeParry)
			if c.CurrentHealth() != want {
				t.Fatal("offhand/duplicate enchant bypassed shared ICD")
			}
			sim.CurrentTime = 10 * time.Second
			fire(core.ProcMaskMeleeOHSpecial, core.OutcomeParry)
			if math.Abs(c.CurrentHealth()-(want+c.MaxHealth()*.05)) > 1e-6 {
				t.Fatal("equip aura failed offhand avoidance after ICD")
			}
			if c.ActiveShapeShift != form || (form != nil && !form.IsActive()) || c.CurrentMana() != mana {
				t.Fatal("Recovery changed form or charged mana")
			}
			c.Equipment[proto.ItemSlot_ItemSlotMainHand].Enchant = core.Enchant{}
			c.Equipment[proto.ItemSlot_ItemSlotOffHand].Enchant = core.Enchant{}
			sim.CurrentTime = 20 * time.Second
			hp = c.CurrentHealth()
			fire(core.ProcMaskMeleeMHAuto, core.OutcomeParry)
			if c.CurrentHealth() != hp {
				t.Fatal("unequipped Recovery still healed")
			}
		})
	}
}
