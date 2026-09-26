//go:build with_db

package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	sim.RegisterAll()
}

func TestIntrinsicShieldBlockSurvivesDatabaseImport(t *testing.T) {
	for _, id := range []int32{272591, 278469, 279262, 920000383, 920000384} {
		item := core.ItemsByID[id]
		if item.Stats[stats.BlockValue] != 44 {
			t.Errorf("shield %d has %g Block, want 44", id, item.Stats[stats.BlockValue])
		}
	}
}

func tankMechanicsFixture(t *testing.T, role string) *core.Simulation {
	t.Helper()
	data, err := os.ReadFile("../../artifacts/tanks/" + role + "_1t_20261993_selected.json")
	if err != nil {
		t.Fatal(err)
	}
	var saved struct {
		Request json.RawMessage `json:"request"`
	}
	if err := json.Unmarshal(data, &saved); err != nil {
		t.Fatal(err)
	}
	request := &proto.RaidSimRequest{}
	if err := protojson.Unmarshal(saved.Request, request); err != nil {
		t.Fatal(err)
	}
	request.SimOptions.Iterations = 1
	request.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	s := core.NewSim(request, simsignals.Signals{})
	s.Options.Interactive = true
	s.Reset()
	return s
}

func TestShieldBlockRecastRestoresCharges(t *testing.T) {
	s := tankMechanicsFixture(t, "tank_warrior")
	unit := s.Raid.AllPlayerUnits[0]
	spell := unit.GetSpell(core.ActionID{SpellID: 2565})
	aura := unit.GetAura("Shield Block")
	unit.AddRage(s, 100, unit.NewRageMetrics(core.ActionID{SpellID: 2565}))
	blockBefore := unit.GetStat(stats.Block)
	if !spell.Cast(s, unit) || aura.GetStacks() != 2 {
		t.Fatal("initial cast must grant two blocks")
	}
	aura.RemoveStack(s)
	s.CurrentTime = 6 * time.Second // original aura is still active; cooldown has expired
	if !spell.Cast(s, unit) {
		t.Fatal("Shield Block should be ready before the old aura expires")
	}
	if aura.GetStacks() != 2 {
		t.Fatalf("recast left %d block charges, want 2", aura.GetStacks())
	}
	if got := unit.GetStat(stats.Block) - blockBefore; got != 75*core.BlockRatingPerBlockChance {
		t.Fatalf("refresh changed the block bonus to %g", got)
	}
}

func TestTemplarsBulwarkDoesNotUseGlobalCooldown(t *testing.T) {
	s := tankMechanicsFixture(t, "protection_paladin")
	unit := s.Raid.AllPlayerUnits[0]
	spell := unit.GetSpell(core.ActionID{SpellID: 1311015})
	unit.SetGCDTimer(s, time.Second)
	if !spell.Cast(s, unit) {
		t.Fatal("Bulwark must be usable during the global cooldown")
	}
	if unit.GCD.ReadyAt() != time.Second {
		t.Fatal("Bulwark changed the existing global cooldown")
	}
	if !unit.GetAura("Templar's Bulwark").IsActive() || !unit.GetAura("Forbearance").IsActive() {
		t.Fatal("Bulwark did not apply its absorb and Forbearance")
	}
}

func TestHolyShieldUsesHastedHolyGlobalCooldown(t *testing.T) {
	for _, id := range []int32{20925, 20927, 20928} {
		s := tankMechanicsFixture(t, "protection_paladin")
		unit := s.Raid.AllPlayerUnits[0]
		unit.MultiplyCastSpeed(1.2)
		spell := unit.GetSpell(core.ActionID{SpellID: id})
		if !spell.Cast(s, unit) {
			t.Fatalf("Holy Shield %d failed to cast", id)
		}
		want := unit.ApplyCastSpeed(core.GCDDefault)
		if spell.SpellSchool != core.SpellSchoolHoly || unit.GCD.ReadyAt() != want {
			t.Fatalf("Holy Shield %d: school=%v GCD=%s, want Holy / %s", id, spell.SpellSchool, unit.GCD.ReadyAt(), want)
		}
	}
}
