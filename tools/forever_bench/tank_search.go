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

// These are selection tolerances, not game formulas. Every candidate is
// compared with the frozen published control, never a repeatedly weakened
// predecessor. Search screens use means; final independent-seed confirmation
// must separately assess uncertainty and may reject the search's winner.
type tankGuardPolicy struct {
	MaxDTPSIncrease, MaxTMIIncrease                     float64
	MaxDeathIncrease, MaxTPSLoss, MaxLeastTargetTPSLoss float64
}

var tankPolicy = tankGuardPolicy{
	MaxDTPSIncrease: .01, MaxTMIIncrease: .02, MaxDeathIncrease: .0025,
	MaxTPSLoss: .01, MaxLeastTargetTPSLoss: .02,
}

type tankPair struct {
	Single resultRow
	Multi  resultRow
}

func runTankPair(b build, p *proto.Player, count int, rng int64) tankPair {
	runEncounter := func(targets int, seconds float64) resultRow {
		return runWithRequest(b, p, count, rng, func(b build, p *proto.Player, n int, seed int64) *proto.RaidSimRequest {
			return tankRequest(b, p, n, seed, targets, seconds)
		})
	}
	return tankPair{Single: runEncounter(1, 300), Multi: runEncounter(3, 90)}
}

func (policy tankGuardPolicy) failures(candidate, reference tankPair) []string {
	var failures []string
	for _, scenario := range []struct {
		name string
		c, r resultRow
	}{{"single", candidate.Single, reference.Single}, {"multi", candidate.Multi, reference.Multi}} {
		c, r := scenario.c.Tank, scenario.r.Tank
		if c == nil || r == nil {
			return []string{"missing tank measurements"}
		}
		checks := []struct {
			name string
			ok   bool
		}{
			{"DTPS", c.DTPS <= r.DTPS*(1+policy.MaxDTPSIncrease)},
			{"TMI", c.TMI <= r.TMI*(1+policy.MaxTMIIncrease)},
			{"death", c.ChanceOfDeath <= r.ChanceOfDeath+policy.MaxDeathIncrease},
			{"TPS", c.TPS >= r.TPS*(1-policy.MaxTPSLoss)},
			{"least-target TPS", c.LeastTargetTPS >= r.LeastTargetTPS*(1-policy.MaxLeastTargetTPSLoss)},
			{"warnings", len(scenario.c.Warnings) == 0},
		}
		for _, check := range checks {
			if !check.ok {
				failures = append(failures, scenario.name+": "+check.name)
			}
		}
	}
	return failures
}

// Optimize the published single-target row without sacrificing the separate
// three-attacker workload. Both scenarios gate every shortlisted choice.
func optimizeTankGear(b build, initial *proto.Player) *proto.Player {
	if !*modeledOnly || *targetCount != 1 || *duration != 300 || *targetArmor != 3731 || *demon {
		panic("tank gear search requires modeled-only and the canonical 300s/3731-armor Dragonkin scenario")
	}
	p := googleProto.Clone(initial).(*proto.Player)
	if !fullyModeledV2(p) {
		panic("tank search must start from a complete modeled-v2 control")
	}
	pool, excluded := comparisonGearPool(b, p)
	referencePlayer := tankGuardPlayer(b, p.Race)
	anchor := runTankPair(b, referencePlayer, *iterations, *seed)
	if len(anchor.Single.Warnings)+len(anchor.Multi.Warnings) != 0 {
		panic("tank gear search requires a valid frozen control without APL warnings")
	}
	initialPair := runTankPair(b, p, *iterations, *seed)
	report := gearSearchReport{
		Build: b.Key, Race: raceName(p.Race), Baseline: initialPair.Single,
		Excluded: excluded, TankGuardPolicy: &tankPolicy, TankControl: &anchor,
	}
	for _, item := range pool {
		report.PoolIDs = append(report.PoolIDs, item.ID)
	}
	evaluate := func(player *proto.Player, pass, slot, count int, rng int64, stage string, reference tankPair) tankPair {
		pair := runTankPair(b, player, count, rng)
		report.Trials = append(report.Trials, gearTrial{
			Pass: pass, Slot: slot, Iterations: count, Seed: rng,
			Equipment: googleProto.Clone(player.Equipment).(*proto.EquipmentSpec), Profession2: player.Profession2,
			DPS: pair.Single.DPS, StandardError: pair.Single.StandardError, Stage: stage,
			Tank: pair.Single.Tank, TankMulti: pair.Multi.Tank,
			Warnings:    append(append([]string{}, pair.Single.Warnings...), pair.Multi.Warnings...),
			RejectedFor: tankPolicy.failures(pair, reference),
		})
		return pair
	}
	for pass := 1; pass <= *gearPasses; pass++ {
		report.Passes = pass
		changed := false
		slots := []int{0, 2, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16}
		for slot := 0; slot < 17; slot++ {
			slots = append(slots, 17+slot)
		}
		for _, slot := range slots {
			var candidates []*proto.Player
			if slot >= 17 {
				candidates = enchantCandidates(b, p, slot-17)
			} else {
				candidates = gearCandidates(b, p, slot, pool)
			}
			if len(candidates) == 0 {
				continue
			}
			rng := *seed + int64(pass*1000+slot)
			reference := runTankPair(b, referencePlayer, *gearScreen, rng)
			current := evaluate(p, pass, slot, *gearScreen, rng, "screen-baseline", reference)
			type choice struct {
				player *proto.Player
				pair   tankPair
			}
			var shortlist []choice
			for _, candidate := range candidates {
				pair := evaluate(candidate, pass, slot, *gearScreen, rng, "screen", reference)
				if pair.Single.DPS > current.Single.DPS && len(tankPolicy.failures(pair, reference)) == 0 &&
					doesNotWorsenHit(pair.Single.Hit, current.Single.Hit) {
					shortlist = append(shortlist, choice{candidate, pair})
				}
			}
			sort.SliceStable(shortlist, func(i, j int) bool { return shortlist[i].pair.Single.DPS > shortlist[j].pair.Single.DPS })
			if len(shortlist) == 0 {
				continue
			}
			rng += 1000000
			reference = runTankPair(b, referencePlayer, *gearValidate, rng)
			current = evaluate(p, pass, slot, *gearValidate, rng, "validation-baseline", reference)
			best, bestPair := p, current
			for _, candidate := range shortlist[:min(3, len(shortlist))] {
				pair := evaluate(candidate.player, pass, slot, *gearValidate, rng, "validation", reference)
				gainBound := 2 * (pair.Single.StandardError + current.Single.StandardError)
				if pair.Single.DPS > bestPair.Single.DPS && pair.Single.DPS-current.Single.DPS > gainBound &&
					len(tankPolicy.failures(pair, reference)) == 0 &&
					doesNotWorsenHit(pair.Single.Hit, current.Single.Hit) {
					best, bestPair = candidate.player, pair
				}
			}
			if best != p {
				p, changed = best, true
				report.Accepted = append(report.Accepted, gearTrial{
					Pass: pass, Slot: slot, Iterations: *gearValidate, Seed: rng,
					Equipment: googleProto.Clone(p.Equipment).(*proto.EquipmentSpec), Profession2: p.Profession2,
					DPS: bestPair.Single.DPS, StandardError: bestPair.Single.StandardError,
					Tank: bestPair.Single.Tank, TankMulti: bestPair.Multi.Tank, Stage: "accepted",
				})
				fmt.Printf("%s/%s pass %d slot %d: %.2f -> %.2f; tank guards passed\n",
					b.Key, report.Race, pass, slot, current.Single.DPS, bestPair.Single.DPS)
			}
		}
		if !changed {
			report.Converged = true
			break
		}
	}
	final := runTankPair(b, p, *iterations, *seed)
	report.TankProposed = &final
	failures := tankPolicy.failures(final, anchor)
	if final.Single.DPS < initialPair.Single.DPS || math.IsNaN(final.Single.DPS) {
		failures = append(failures, "final single-target DPS")
	}
	selected := p
	report.FinalDisposition = "retained-pending-independent-confirmation"
	if len(failures) != 0 {
		report.FinalDisposition, report.FinalRejection = "reverted", failures
		proposed := final
		report.TankProposed = &proposed
		selected, final = initial, initialPair
		fmt.Printf("%s/%s: final control check failed (%v); retaining input\n", b.Key, report.Race, failures)
	}
	report.Final, report.TankFinal = final.Single, &final
	path := *output + "." + b.Key + "." + strings.ReplaceAll(strings.ToLower(report.Race), " ", "-") + ".gear-search.json"
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(path, append(data, '\n'), 0644); err != nil {
		panic(err)
	}
	return selected
}
