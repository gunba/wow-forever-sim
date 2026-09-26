//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func TestForeverUtilitySpellSchoolsAndGCDs(t *testing.T) {
	// SpellMisc 1.60.1.70009 and SpellCooldowns 1.60.1.69977:
	// these are magical, 1500 ms-GCD actions, not physical abilities.
	for _, tc := range []struct {
		build  string
		ids    []int32
		school core.SpellSchool
	}{
		{"balance", []int32{29166}, core.SpellSchoolNature},
		{"elemental", []int32{324, 325, 905, 945, 8134, 10431, 10432}, core.SpellSchoolNature},
		{"beast_mastery", []int32{25296}, core.SpellSchoolNature},
		{"arcane", []int32{12051}, core.SpellSchoolArcane},
		{"shadow", []int32{15473}, core.SpellSchoolShadow},
		{"demonology", []int32{19028}, core.SpellSchoolShadow},
	} {
		for _, id := range tc.ids {
			t.Run(core.ActionID{SpellID: id}.String(), func(t *testing.T) {
				var req *proto.RaidSimRequest
				for _, build := range builds() {
					if build.Key == tc.build {
						req = racialFixture(tc.build, build.races()[0])
						break
					}
				}
				if req == nil {
					t.Fatalf("missing fixture %s", tc.build)
				}
				req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
				s := core.NewSim(req, simsignals.Signals{})
				s.Options.Interactive = true
				s.Reset()
				unit := s.Raid.AllPlayerUnits[0]
				spell := unit.GetSpell(core.ActionID{SpellID: id})
				// Innervate is tagged with the provider's raid index.
				if spell == nil && id == 29166 {
					for _, candidate := range unit.Spellbook {
						if candidate.SpellID == id {
							spell = candidate
							break
						}
					}
				}
				if spell == nil || spell.SpellSchool != tc.school {
					t.Fatalf("spell %d must be registered with school %v", id, tc.school)
				}
				unit.MultiplyCastSpeed(1.2)
				if !spell.Cast(s, unit.CurrentTarget) {
					t.Fatalf("spell %d failed to cast", id)
				}
				if got, want := spell.CurCast.GCD, unit.ApplyCastSpeed(core.GCDDefault); got != want {
					t.Fatalf("spell %d GCD %s, want %s", id, got, want)
				}
			})
		}
	}
}
