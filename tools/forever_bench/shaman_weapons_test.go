//go:build with_db

package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestShamanProfilesUseLegalWeapons(t *testing.T) {
	for _, b := range builds() {
		if b.Class != proto.Class_ClassShaman {
			continue
		}
		p := b.player(proto.Race_RaceOrc)
		if err := validateGear(p); err != nil {
			t.Fatalf("%s: %v", b.Key, err)
		}
		req := request(p, 1, 1)
		env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
		if env.Raid.Parties[0].Players[0].GetCharacter().AutoAttacks.IsDualWielding {
			t.Fatalf("%s has off-hand auto attacks", b.Key)
		}
		for _, swapped := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/swap=%t", b.Key, swapped), func(t *testing.T) {
				p := b.player(proto.Race_RaceOrc)
				mh, oh := &proto.ItemSpec{Id: 279261}, &proto.ItemSpec{Id: 22384}
				if swapped {
					p.EnableItemSwap = true
					p.ItemSwap = &proto.ItemSwap{MhItem: mh, OhItem: oh}
				} else {
					p.Equipment.Items[14], p.Equipment.Items[15] = mh, oh
					if err := validateGear(p); err == nil || !strings.Contains(err.Error(), "cannot dual wield") {
						t.Fatalf("benchmark accepted dual wield: %v", err)
					}
				}
				defer func() {
					if caught := recover(); caught == nil || !strings.Contains(fmt.Sprint(caught), "cannot dual wield") {
						t.Errorf("engine did not reject dual wield: %v", caught)
					}
				}()
				req := request(p, 1, 1)
				core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
			})
		}
	}
}
