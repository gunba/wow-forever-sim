package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/wowsims/classic/sim/core/proto"
	googleProto "google.golang.org/protobuf/proto"
)

type talentTrial struct {
	Stage, Talents          string
	Round                   int
	Seed                    int64
	Iterations              int32
	DPS, StandardError, OOM float64
	Warnings                []string
}

type talentAudit struct {
	Build, Race, FinalTalents string
	AcceptanceStandardErrors  float64
	BaselinePlayer            json.RawMessage
	Scenario                  json.RawMessage
	Trials                    []talentTrial
}

func (a *talentAudit) record(stage string, round int, rng int64, result resultRow) {
	if a.BaselinePlayer == nil {
		a.BaselinePlayer, a.Scenario = result.BaselinePlayer, result.Request
	}
	a.Trials = append(a.Trials, talentTrial{
		Stage: stage, Round: round, Seed: rng, Talents: result.Talents,
		Iterations: result.Iterations, DPS: result.DPS, StandardError: result.StandardError,
		OOM: result.OOMSeconds, Warnings: result.Warnings,
	})
}

func (a *talentAudit) save() {
	if *output == "" {
		return
	}
	path := *output + "." + a.Build + "." + strings.ReplaceAll(a.Race, " ", "_") + ".talent-search.json"
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(path, append(data, '\n'), 0644); err != nil {
		panic(err)
	}
}

// Screen every legal one-point reallocation, then independently verify the best
// candidates. The input gear, APL, consumes, buffs and encounter do not change.
func optimizeTalents(b build, input *proto.Player) *proto.Player {
	if b.isTank() {
		panic("DPS-only talent optimization is not valid for tanks; use research_tanks.py and select_tanks.py with both tank scenarios")
	}
	config := loadTalents(b)
	if err := config.validate(input.TalentsString); err != nil {
		panic(err)
	}
	best := googleProto.Clone(input).(*proto.Player)
	audit := talentAudit{Build: b.Key, Race: raceName(best.Race), AcceptanceStandardErrors: 2.5}
	visited := map[string]bool{}
	for round := 0; round < *searchRounds; round++ {
		baseline := run(b, best, 250, *seed)
		audit.record("screen-baseline", round+1, *seed, baseline)
		allowedWarnings := map[string]bool{}
		for _, warning := range baseline.Warnings {
			allowedWarnings[warning] = true
		}
		type candidate struct {
			talents string
			dps     float64
		}
		var candidates []candidate
		for _, talents := range config.neighbors(best.TalentsString) {
			if visited[talents] {
				continue
			}
			visited[talents] = true
			p := googleProto.Clone(best).(*proto.Player)
			p.TalentsString = talents
			r := run(b, p, 250, *seed)
			audit.record("screen", round+1, *seed, r)
			compatible := true
			for _, warning := range r.Warnings {
				if !allowedWarnings[warning] {
					compatible = false
				}
			}
			if !compatible {
				continue
			}
			if r.DPS > baseline.DPS {
				candidates = append(candidates, candidate{talents, r.DPS})
			}
		}
		sort.Slice(candidates, func(i, j int) bool { return candidates[i].dps > candidates[j].dps })
		fmt.Printf("SEARCH %s/%s round %d screened %d, improving %d\n", b.Key, raceName(best.Race), round+1, len(visited), len(candidates))
		validationSeed := *seed + int64(10000+round)
		winner := run(b, best, 2000, validationSeed)
		audit.record("validate-baseline", round+1, validationSeed, winner)
		validationBaseline := winner.DPS
		next := googleProto.Clone(best).(*proto.Player)
		for _, candidate := range candidates[:min(5, len(candidates))] {
			p := googleProto.Clone(best).(*proto.Player)
			p.TalentsString = candidate.talents
			result := run(b, p, 2000, validationSeed)
			audit.record("validate", round+1, validationSeed, result)
			margin := 2.5 * math.Hypot(winner.StandardError, result.StandardError)
			if result.DPS > winner.DPS+margin {
				winner, next = result, p
			}
		}
		if next.TalentsString == best.TalentsString {
			audit.FinalTalents = best.TalentsString
			audit.save()
			break
		}
		fmt.Printf("KEEP %s %.2f -> %.2f %s\n", b.Key, validationBaseline, winner.DPS, next.TalentsString)
		best = next
		audit.FinalTalents = best.TalentsString
		audit.save()
	}
	return best
}
