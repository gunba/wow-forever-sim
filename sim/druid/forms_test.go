package druid

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestClassicFormWeaponsRemainFixed(t *testing.T) {
	d := &Druid{}
	d.Env = &core.Environment{Ruleset: proto.Ruleset_RulesetClassic}
	d.Equipment[proto.ItemSlot_ItemSlotMainHand] = core.Item{ID: 1, WeaponType: proto.WeaponType_WeaponTypeStaff, SwingSpeed: 3, WeaponDamageMin: 120, WeaponDamageMax: 180}
	cat, bear := d.GetCatWeapon(), d.GetBearWeapon()
	if cat.BaseDamageMin != 43.84 || cat.BaseDamageMax != 65.76 || bear.BaseDamageMin != 109 || bear.BaseDamageMax != 165 {
		t.Fatal("Forever normalization changed Classic paw weapons")
	}
}

func TestFurorRankCapAndBearExclusion(t *testing.T) {
	for rank := int32(0); rank <= 5; rank++ {
		d := &Druid{Talents: &proto.DruidTalents{Furor: rank}}
		d.Env = &core.Environment{Ruleset: proto.Ruleset_RulesetForever}
		d.EnableEnergyBar(100)
		d.lastCatFormEnergy, d.lastCatFormExitAt = 100, 0
		sim := &core.Simulation{CurrentTime: 5 * time.Second}
		if got, want := d.furorShiftEnergy(sim), float64(rank)*20; got != want {
			t.Errorf("rank %d: %v Energy exceeds combined carry/regen cap %v", rank, got, want)
		}
		d.form = Bear
		if got := d.furorShiftEnergy(sim); got != 0 {
			t.Errorf("rank %d: recovered %v Energy while in Bear", rank, got)
		}
	}
}

func TestFurorBearHistoryDiscardAndResume(t *testing.T) {
	d := &Druid{Talents: &proto.DruidTalents{Furor: 5}}
	d.Env = &core.Environment{Ruleset: proto.Ruleset_RulesetForever}
	d.EnableEnergyBar(100)
	d.lastCatFormEnergy, d.lastCatFormExitAt = 90, time.Second
	d.form = Bear
	d.resetFurorHistory(core.NeverExpires)
	if d.lastCatFormEnergy != 0 || d.lastCatFormExitAt != core.NeverExpires {
		t.Fatal("Bear entry retained Cat history")
	}
	sim := &core.Simulation{CurrentTime: 100 * time.Second}
	if got := d.furorShiftEnergy(sim); got != 0 {
		t.Fatalf("long Bear stay regenerated %v Energy", got)
	}
	d.form = Humanoid
	d.resetFurorHistory(sim.CurrentTime)
	sim.CurrentTime += time.Second
	if got := d.furorShiftEnergy(sim); got != 10 {
		t.Fatalf("Bear exit: got %v, want only the subsequent humanoid second's 10 Energy", got)
	}
}

func TestFurorExitTimesAndIterationReset(t *testing.T) {
	d := &Druid{Talents: &proto.DruidTalents{Furor: 5}}
	d.EnableEnergyBar(100)
	for _, exitAt := range []time.Duration{-2 * time.Second, 0, time.Second} {
		d.lastCatFormEnergy, d.lastCatFormExitAt = 20, exitAt
		sim := &core.Simulation{CurrentTime: exitAt + time.Second}
		if got := d.furorShiftEnergy(sim); got != 30 {
			t.Fatalf("exit at %v: got %v Energy, want 20 carried + 10 regenerated", exitAt, got)
		}
		d.Reset(sim)
		if got := d.furorShiftEnergy(sim); got != 0 {
			t.Fatalf("iteration reset retained Furor history: %v", got)
		}
	}
}
