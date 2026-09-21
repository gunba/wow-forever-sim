//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestSealOfCommandRequiresTalent(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "retribution" {
			continue
		}
		for _, talented := range []bool{false, true} {
			p := b.player(proto.Race_RaceUndead)
			config := loadTalents(b)
			points, err := config.decode(p.TalentsString)
			if err != nil {
				t.Fatal(err)
			}
			offset := 0
			for _, tree := range config.Trees {
				for i, talent := range tree.Talents {
					if talent.Field == "sealOfCommand" {
						points[offset+i] = core.TernaryInt(talented, 1, 0)
					}
				}
				offset += len(tree.Talents)
			}
			p.TalentsString = config.encode(points)
			req := request(p, 1, 123)
			env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
			character := env.Raid.Parties[0].Players[0].GetCharacter()
			for _, id := range []int32{20375, 20915, 20918, 20919, 20920} {
				if got := character.GetSpell(core.ActionID{SpellID: id}) != nil; got != talented {
					t.Errorf("rank %d available=%v with talent=%v", id, got, talented)
				}
			}
		}
	}
}
