package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sync"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

var tankPairOutput = flag.Bool("tank-pair", false, "record the fixed single/multi-target tank scenarios together")
var tankNeighbors = flag.Bool("tank-neighbors", false, "write legal one-point tank talent alternatives without simulation")

// Frozen published controls retain each tank's actual external support,
// consumables and healing scenario. They are not interchangeable with the
// non-attacking DPS encounter or an assertion of equal support across classes.
var tankControls = sync.OnceValue(func() map[string]*proto.RaidSimRequest {
	controls := map[string]*proto.RaidSimRequest{}
	for _, key := range []string{"tank_warrior", "protection_paladin", "feral_tank_druid"} {
		var saved struct {
			Request json.RawMessage `json:"request"`
		}
		path := "artifacts/tanks/" + key + "_1t_20261993_selected.json"
		if err := json.Unmarshal(mustRead(path), &saved); err != nil {
			panic(err)
		}
		request := &proto.RaidSimRequest{}
		if err := protojson.Unmarshal(saved.Request, request); err != nil {
			panic(err)
		}
		controls[key] = request
	}
	return controls
})

func (b build) isTank() bool {
	return b.Key == "tank_warrior" || b.Key == "protection_paladin" || b.Key == "feral_tank_druid"
}

func tankControl(key string) *proto.RaidSimRequest {
	request := tankControls()[key]
	if request == nil {
		panic("unknown tank control: " + key)
	}
	return request
}

// The original Warrior controls used an invalid direct-cast rotation. Freeze
// the corrected per-race inputs separately; never move the guard with a trial.
func tankGuardPlayer(b build, race proto.Race) *proto.Player {
	if b.Key == "tank_warrior" {
		var controls struct {
			Results []struct {
				Race           string
				BaselinePlayer json.RawMessage
			}
		}
		if err := json.Unmarshal(mustRead("artifacts/tanks/queue_corrected_warrior_controls.json"), &controls); err != nil {
			panic(err)
		}
		for _, row := range controls.Results {
			if row.Race == raceName(race) {
				player := &proto.Player{}
				if err := protojson.Unmarshal(row.BaselinePlayer, player); err != nil {
					panic(err)
				}
				return player
			}
		}
		panic("missing queue-corrected Warrior guard control")
	}
	player := googleProto.Clone(tankControl(b.Key).Raid.Parties[0].Players[0]).(*proto.Player)
	player.Race = race
	return player
}

func tankRequest(b build, p *proto.Player, count int, rng int64, targets int, seconds float64) *proto.RaidSimRequest {
	if targets < 1 || targets > 10 {
		panic("tank targets must be between 1 and 10")
	}
	req := googleProto.Clone(tankControl(b.Key)).(*proto.RaidSimRequest)
	req.Raid.Parties[0].Players[0] = googleProto.Clone(p).(*proto.Player)
	req.Encounter.Duration = seconds
	target := req.Encounter.Targets[0]
	req.Encounter.Targets = nil
	for i := 0; i < targets; i++ {
		copy := googleProto.Clone(target).(*proto.Target)
		if targets > 1 {
			copy.MinBaseDamage = 1500
		}
		req.Encounter.Targets = append(req.Encounter.Targets, copy)
	}
	player := req.Raid.Parties[0].Players[0]
	if targets > 1 {
		player.HealingModel.Hps = 2500
	}
	req.SimOptions.Iterations, req.SimOptions.RandomSeed = int32(count), rng
	req.SimOptions.Ruleset = proto.Ruleset_RulesetForever
	return req
}

func (b build) validateTankIdentity(p *proto.Player) error {
	if !b.isTank() {
		return nil
	}
	if !p.InFrontOfTarget || p.HealingModel == nil {
		return fmt.Errorf("%s requires front-facing tank position and an explicit healing model", b.Key)
	}
	mainHand := core.ItemsByID[p.Equipment.Items[14].GetId()]
	offHand := core.ItemsByID[p.Equipment.Items[15].GetId()]
	if b.Key == "feral_tank_druid" {
		if p.GetFeralTankDruid() == nil {
			return fmt.Errorf("Bear requires the tank specialization, not Cat")
		}
	} else if offHand.WeaponType != proto.WeaponType_WeaponTypeShield ||
		mainHand.HandType == proto.HandType_HandTypeTwoHand {
		return fmt.Errorf("%s requires a one-hand weapon and shield", b.Key)
	}
	return nil
}

type tankMetrics struct {
	TPS, TPSStandardError                float64
	DTPS, DTPSStandardError              float64
	TMI, TMIStandardError, ChanceOfDeath float64
	TargetTPS                            []float64
	LeastTargetTPS                       float64
	EncounterMetrics                     json.RawMessage
}

func tankResultMetrics(result *proto.RaidSimResult) *tankMetrics {
	player := result.RaidMetrics.Parties[0].Players[0]
	scale := math.Sqrt(float64(result.IterationsDone))
	out := &tankMetrics{
		TPS: player.GetThreat().GetAvg(), TPSStandardError: player.GetThreat().GetStdev() / scale,
		DTPS: player.GetDtps().GetAvg(), DTPSStandardError: player.GetDtps().GetStdev() / scale,
		TMI: player.GetTmi().GetAvg(), TMIStandardError: player.GetTmi().GetStdev() / scale,
		ChanceOfDeath: player.ChanceOfDeath,
	}
	// Action target indices are unit indices, not positions in the target
	// list. Helpful/self actions may share these arrays and must not become
	// a fabricated extra enemy or inflate a least-target coverage result.
	for _, target := range result.EncounterMetrics.Targets {
		total := 0.0
		for _, action := range player.Actions {
			for _, hit := range action.Targets {
				if hit.UnitIndex == target.UnitIndex {
					total += hit.Threat
				}
			}
		}
		out.TargetTPS = append(out.TargetTPS, total/float64(result.IterationsDone)/result.AvgIterationDuration)
	}
	out.LeastTargetTPS = math.Inf(1)
	for _, value := range out.TargetTPS {
		out.LeastTargetTPS = min(out.LeastTargetTPS, value)
	}
	var err error
	out.EncounterMetrics, err = protojson.Marshal(result.EncounterMetrics)
	if err != nil {
		panic(err)
	}
	return out
}

func writeTankPair(b build, p *proto.Player) {
	if !b.isTank() || *output == "" || *targetCount != 1 || *duration != 300 {
		panic("tank-pair requires a tank, output and the standard encounter")
	}
	pair := runTankPair(b, p, *iterations, *seed)
	data, err := json.MarshalIndent(pair, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Dir(*output), 0755); err != nil {
		panic(err)
	}
	if err := os.WriteFile(*output+".pair.json", append(data, '\n'), 0644); err != nil {
		panic(err)
	}
	fmt.Printf("%s/%s: %.2f DPS / %.2f TPS / %.2f DTPS; multi %.2f DPS / %.2f minimum TPS\n",
		b.Key, raceName(p.Race), pair.Single.DPS, pair.Single.Tank.TPS,
		pair.Single.Tank.DTPS, pair.Multi.DPS, pair.Multi.Tank.LeastTargetTPS)
}
