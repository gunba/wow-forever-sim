// Tank encounter replay harness. This does not alter the published DPS roster.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

func main() {
	sim.RegisterAll()
	playerPath := flag.String("player", "", "exact tank player proto JSON")
	requestPath := flag.String("request", "", "replay an exact web/native RaidSimRequest JSON instead of building a scenario")
	aplPath := flag.String("apl", "", "optional APL JSON override")
	out := flag.String("output", "", "complete request/result JSON output")
	race := flag.String("race", "", "override player race enum")
	duration := flag.Float64("duration", 300, "encounter seconds")
	targets := flag.Int("targets", 1, "number of identical attackers, 1-10")
	swing := flag.Float64("swing", 3000, "unmitigated base damage per boss swing")
	healing := flag.Float64("healing", 1500, "external healing per second")
	iterations := flag.Int("iterations", 5000, "iterations")
	seed := flag.Int64("seed", 20260925, "random seed")
	tier := flag.Bool("tier1", true, "enable forced Forever tank Tier 1")
	flag.Parse()
	if *requestPath != "" {
		data, err := os.ReadFile(*requestPath)
		if err != nil {
			panic(err)
		}
		request := &proto.RaidSimRequest{}
		if err := protojson.Unmarshal(data, request); err != nil {
			panic(err)
		}
		if *race != "" {
			id, ok := proto.Race_value[*race]
			if !ok {
				panic("unknown race " + *race)
			}
			request.Raid.Parties[0].Players[0].Race = proto.Race(id)
		}
		if *aplPath != "" {
			data, err := os.ReadFile(*aplPath)
			if err != nil {
				panic(err)
			}
			request.Raid.Parties[0].Players[0].Rotation = core.APLRotationFromJsonString(string(data))
		}
		runRequest(request, *out)
		return
	}
	if *playerPath == "" || *targets < 1 || *targets > 10 || *iterations < 1 {
		panic("-player required; targets 1-10; iterations >=1")
	}
	input, err := os.ReadFile(*playerPath)
	if err != nil {
		panic(err)
	}
	player := &proto.Player{}
	if err = protojson.Unmarshal(input, player); err != nil {
		panic(err)
	}
	if *race != "" {
		id, ok := proto.Race_value[*race]
		if !ok {
			panic("unknown race " + *race)
		}
		player.Race = proto.Race(id)
	}
	if *aplPath != "" {
		data, e := os.ReadFile(*aplPath)
		if e != nil {
			panic(e)
		}
		player.Rotation = core.APLRotationFromJsonString(string(data))
	}
	player.ForeverTier1Bonuses = *tier
	player.HealingModel = &proto.HealingModel{Hps: *healing, CadenceSeconds: 3, BurstWindow: 6}
	raid := core.SinglePlayerRaidProto(player, core.ForeverBuffs.Party, core.ForeverBuffs.Raid, core.ForeverBuffs.Debuffs)
	raid = googleProto.Clone(raid).(*proto.Raid)
	raid.Tanks = []*proto.UnitReference{{Type: proto.UnitReference_Player, Index: 0}}
	// Forever's air totems are exclusive. The tank is in a melee group.
	raid.Buffs.GraceOfAirTotem = proto.TristateEffect_TristateEffectMissing
	raid.Parties[0].Buffs.WindfuryTotem = true
	if player.GetTankWarrior() != nil {
		// This tank supplies Battle Shout, Sunder, Thunder Clap and Demoralizing Shout.
		raid.Buffs.BattleShout = proto.TristateEffect_TristateEffectMissing
		raid.Debuffs.SunderArmor = false
		raid.Debuffs.ExposeArmor = proto.TristateEffect_TristateEffectMissing
		raid.Debuffs.ThunderClap = proto.TristateEffect_TristateEffectMissing
		raid.Debuffs.DemoralizingShout = proto.TristateEffect_TristateEffectMissing
		raid.Debuffs.DemoralizingRoar = proto.TristateEffect_TristateEffectMissing
	} else if player.GetFeralTankDruid() != nil {
		raid.Debuffs.DemoralizingRoar = proto.TristateEffect_TristateEffectMissing
		raid.Debuffs.DemoralizingShout = proto.TristateEffect_TristateEffectMissing
	}
	// Keep the default site's level-63 target armor, AP and health. The tank
	// scenario enables its swings and fixed external healing.
	encounter := &proto.Encounter{Duration: *duration, Targets: make([]*proto.Target, *targets)}
	for i := range encounter.Targets {
		encounter.Targets[i] = &proto.Target{Level: 63, MobType: proto.MobType_MobTypeDragonkin,
			Stats:         stats.Stats{stats.Armor: 3731, stats.AttackPower: 805, stats.Health: 127393}.ToFloatArray(),
			MinBaseDamage: *swing, DamageSpread: .3333, SwingSpeed: 2, ParryHaste: true, TankIndex: 0}
	}
	request := &proto.RaidSimRequest{Raid: raid, Encounter: encounter, SimOptions: &proto.SimOptions{Iterations: int32(*iterations), RandomSeed: *seed, Ruleset: proto.Ruleset_RulesetForever}}
	runRequest(request, *out)
}

func runRequest(request *proto.RaidSimRequest, outputPath string) {
	// The environment compiler reports invalid APL actions before the run.
	probe := googleProto.Clone(request).(*proto.RaidSimRequest)
	_, rs, _ := core.NewEnvironment(probe.Raid, probe.Encounter, request.SimOptions.Ruleset, false)
	var warnings []string
	for _, action := range append(rs.Parties[0].Players[0].RotationStats.GetPrepullActions(), rs.Parties[0].Players[0].RotationStats.GetPriorityList()...) {
		warnings = append(warnings, action.Warnings...)
	}
	result := core.RunRaidSim(googleProto.Clone(request).(*proto.RaidSimRequest))
	if result.Error != nil {
		panic(result.Error.Message)
	}
	m := result.RaidMetrics.Parties[0].Players[0]
	if math.IsNaN(m.Threat.Avg) || math.IsNaN(m.Dtps.Avg) {
		panic("non-finite tank metrics")
	}
	rawRequest, err := protojson.Marshal(request)
	if err != nil {
		panic(err)
	}
	rawResult, err := protojson.Marshal(result)
	if err != nil {
		panic(err)
	}
	payload := map[string]interface{}{"request": json.RawMessage(rawRequest), "result": json.RawMessage(rawResult), "warnings": warnings}
	if outputPath != "" {
		data, e := json.MarshalIndent(payload, "", "  ")
		if e != nil {
			panic(e)
		}
		if e = os.MkdirAll(filepath.Dir(outputPath), 0755); e != nil {
			panic(e)
		}
		if e = os.WriteFile(outputPath, append(data, '\n'), 0644); e != nil {
			panic(e)
		}
	}
	player := request.Raid.Parties[0].Players[0]
	fmt.Printf("%s %s %dt %.0fs: %.2f DPS %.2f TPS %.2f DTPS %.2f TMI %.3f death, %.2f OOMs; warnings %d\n", player.Name, player.Race, len(request.Encounter.Targets), request.Encounter.Duration, m.Dps.Avg, m.Threat.Avg, m.Dtps.Avg, m.Tmi.Avg, m.ChanceOfDeath, m.SecondsOomAvg, len(warnings))
	for _, w := range warnings {
		fmt.Println("APL warning:", w)
	}
}
