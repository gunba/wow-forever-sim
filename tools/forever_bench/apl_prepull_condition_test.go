package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

// A prepull action has no condition in the editor. One supplied in imported
// JSON used to be ignored silently, so a false-conditioned Blood Fury still
// fired before the pull.
func TestPrepullActionWithConditionIsRejected(t *testing.T) {
	var fury build
	for _, b := range builds() {
		if b.Key == "fury" {
			fury = b
		}
	}
	plain := fury.player(proto.Race_RaceOrc)
	conditioned := fury.player(proto.Race_RaceOrc)
	conditioned.Rotation.PrepullActions = append(conditioned.Rotation.PrepullActions, &proto.APLPrepullAction{
		Action: &proto.APLAction{
			Condition: &proto.APLValue{Value: &proto.APLValue_Const{Const: &proto.APLValueConst{Val: "false"}}},
			Action: &proto.APLAction_CastSpell{CastSpell: &proto.APLActionCastSpell{
				SpellId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 20572}}}},
		},
		DoAtValue: &proto.APLValue{Value: &proto.APLValue_Const{Const: &proto.APLValueConst{Val: "-1s"}}},
	})

	req := requestForBuild(fury, conditioned, 20, 1)
	_, env, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	prepull := env.Parties[0].Players[0].RotationStats.PrepullActions
	want := "Prepull actions cannot have a condition, ignoring this Prepull Action"
	if got := prepull[len(prepull)-1].Warnings; len(got) != 1 || got[0] != want {
		t.Fatalf("warnings = %q, want [%q]", got, want)
	}

	dps := func(p *proto.Player) float64 {
		return core.RunSim(requestForBuild(fury, p, 20, 1), nil, simsignals.Signals{}).RaidMetrics.Dps.Avg
	}
	if got, want := dps(conditioned), dps(plain); got != want {
		t.Errorf("DPS with the rejected action = %v, want %v as without it", got, want)
	}
}
