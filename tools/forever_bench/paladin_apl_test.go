//go:build with_db

package main

import (
	"regexp"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestProtectionSwiftJudgementImmediatelyPaysForASecondJudgement(t *testing.T) {
	casts := regexp.MustCompile(`(?m)^\[([0-9.]+)\].* Casting \{SpellID: (20271|53671)\} \(Cost = ([0-9.]+)`)
	for _, b := range builds() {
		if b.Key != "protection_paladin" {
			continue
		}
		for name, player := range map[string]*proto.Player{
			"prototype": b.player(proto.Race_RaceHuman),
			"ranked":    b.rankedPlayer(proto.Race_RaceHuman),
		} {
			t.Run(name, func(t *testing.T) {
				request := tankRequest(b, player, 1, 20263511, 1, 300)
				request.SimOptions.Debug = true
				result := core.RunRaidSim(request)
				if result.Error != nil {
					t.Fatal(result.Error.Message)
				}
				events := casts.FindAllStringSubmatch(result.Logs, -1)
				resets, judgements := 0, 0
				for i, event := range events {
					if event[2] == "20271" {
						judgements++
						continue
					}
					resets++
					if i+1 == len(events) || events[i+1][2] != "20271" ||
						events[i+1][1] != event[1] || events[i+1][3] != "0.000" {
						t.Fatalf("reset did not immediately produce a free Judgement: %v", events[i:])
					}
				}
				if resets != 5 || judgements < 34 {
					t.Fatalf("%d resets, %d Judgements: cooldown opportunity was wasted", resets, judgements)
				}
			})
		}
	}
}

func TestProtectionSealMaintenanceDoesNotSpamAnActiveSeal(t *testing.T) {
	for _, b := range builds() {
		if b.Key != "protection_paladin" {
			continue
		}
		for name, player := range map[string]*proto.Player{
			"prototype": b.player(proto.Race_RaceHuman),
			"ranked":    b.rankedPlayer(proto.Race_RaceHuman),
		} {
			t.Run(name, func(t *testing.T) {
				result := core.RunRaidSim(tankRequest(b, player, 8, 20263011, 1, 300))
				if result.Error != nil {
					t.Fatal(result.Error.Message)
				}
				var casts float64
				for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
					switch action.Id.GetSpellId() {
					case 20154, 20287, 20288, 20289, 20290, 20291, 20292, 20293:
						for _, target := range action.Targets {
							casts += float64(target.Casts)
						}
					}
				}
				// A 30-second seal with a two-second refresh window needs at
				// most 11 applications here. Judgement does not consume it.
				average := casts / float64(result.IterationsDone)
				if average < 1 || average > 12 {
					t.Fatalf("%.2f seal casts per fight: maintain the seal without recasting it as filler", average)
				}
			})
		}
	}
}
