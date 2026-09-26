//go:build with_db

package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestPublishedNextSwingAttacksUseQueueActions(t *testing.T) {
	for _, b := range builds() {
		for _, race := range b.races() {
			for name, player := range map[string]*proto.Player{
				"prototype": b.player(race),
				"ranked":    b.rankedPlayer(race),
			} {
				for _, row := range player.Rotation.GetPriorityList() {
					id := row.GetAction().GetCastSpell().GetSpellId()
					switch id.GetSpellId() {
					case 11567, 25286, 20569, 9881:
						if id.GetTag() != 1 {
							t.Errorf("%s/%s/%s directly casts next-swing damage: %v", b.Key, raceName(race), name, id)
						}
					}
				}
			}
			if b.isTank() {
				player := tankGuardPlayer(b, race)
				request := tankRequest(b, player, 1, 1, 1, 300)
				_, computed, _ := core.NewEnvironment(request.Raid, request.Encounter, proto.Ruleset_RulesetForever, false)
				for _, action := range computed.Parties[0].Players[0].RotationStats.PriorityList {
					if len(action.Warnings) != 0 {
						t.Errorf("%s/%s guard control: %v", b.Key, raceName(race), action.Warnings)
					}
				}
			}
		}
	}
}

func TestNextSwingDamageCannotBeDirectlyCastByAPL(t *testing.T) {
	for _, b := range builds() {
		var ids []int32
		race := proto.Race_RaceOrc
		switch b.Key {
		case "tank_warrior":
			ids = []int32{25286, 20569}
		case "feral_tank_druid":
			ids = []int32{9881}
			race = proto.Race_RaceTauren
		default:
			continue
		}
		for _, id := range ids {
			for _, tag := range []int32{0, 1} {
				t.Run(fmt.Sprintf("%s/%d/tag%d", b.Key, id, tag), func(t *testing.T) {
					player := b.rankedPlayer(race)
					player.Rotation = core.APLRotationFromJsonString(fmt.Sprintf(
						`{"type":"TypeAPL","priorityList":[{"action":{"castSpell":{"spellId":{"spellId":%d,"tag":%d}}}}]}`, id, tag))
					request := tankRequest(b, player, 2, 20263413, 1, 60)
					_, computed, _ := core.NewEnvironment(request.Raid, request.Encounter, proto.Ruleset_RulesetForever, false)
					warnings := computed.Parties[0].Players[0].RotationStats.PriorityList[0].Warnings
					if tag == 0 {
						if len(warnings) == 0 || !strings.Contains(strings.Join(warnings, " "), "queue action") {
							t.Fatalf("direct next-swing attack was not rejected: %v", warnings)
						}
					} else if len(warnings) != 0 {
						t.Fatalf("valid queue action rejected: %v", warnings)
					}
					result := core.RunRaidSim(request)
					if result.Error != nil {
						t.Fatal(result.Error.Message)
					}
					var casts int32
					for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
						if action.Id.GetSpellId() == id && action.Id.GetTag() == 0 {
							for _, target := range action.Targets {
								casts += target.Casts
							}
						}
					}
					if tag == 0 && casts != 0 || tag == 1 && casts == 0 {
						t.Fatalf("damage casts=%d with APL action tag=%d", casts, tag)
					}
				})
			}
		}
	}
}
