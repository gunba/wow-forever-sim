package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

var (
	buildFilter      = flag.String("build", "", "comma-separated build keys; empty selects all")
	raceFilter       = flag.String("race", "", "race name; empty selects all available races")
	factionFilter    = flag.String("faction", "all", "all, alliance, or horde")
	iterations       = flag.Int("iterations", 5000, "iterations per result")
	seed             = flag.Int64("seed", 20260920, "random seed")
	output           = flag.String("output", "", "output prefix for CSV and complete JSON results")
	talentsOverride  = flag.String("talents", "", "talent string override for one build")
	aplOverride      = flag.String("apl", "", "APL JSON override for one build")
	gearOverride     = flag.String("gear", "", "equipment JSON override for one build")
	playerOverride   = flag.String("player", "", "complete player JSON override for one build")
	distanceOverride = flag.Float64("distance", -1, "distance from target override in yards")
	spellPower       = flag.Float64("spell-power", 0, "additional spell power for sensitivity analysis")
	duration         = flag.Float64("duration", 300, "fight duration in seconds")
	targetArmor      = flag.Float64("armor", 3731, "starting target armor before debuffs")
	demon            = flag.Bool("demon", false, "use a demon target for a separate encounter sensitivity check")
	tier1            = flag.Bool("tier1", true, "apply the role's full Forever Tier 1 bonuses independently of gear")
	equipmentScale   = flag.Float64("equipment-scale", 1, "hypothetical item-stat and weapon-damage multiplier; enchants and effects stay fixed")
	baselineResults  = flag.String("baseline-results", "", "read exact unnormalized BaselinePlayer profiles from a previous results file")
	optimize         = flag.Bool("optimize", false, "search legal one-point talent reallocations")
	searchRounds     = flag.Int("rounds", 3, "talent search rounds")
	seedGear         = flag.Bool("write-seed-gear", false, "write reviewed crafted/dungeon starting gearsets")
	refreshEnchants  = flag.Bool("refresh-enchants", false, "fill missing enchants while preserving legal simulated choices")
)

type resultRow struct {
	BaselinePlayer                         json.RawMessage
	Key, Build, Race, Faction, Talents     string
	DPS, StdDev, StandardError, OOMSeconds float64
	Iterations                             int32
	Hit                                    hitAdjustment
	Stats                                  []float64
	Sets                                   []string
	Warnings                               []string
	UnmodeledSetBonuses                    []string
	Request                                json.RawMessage
	Metrics                                json.RawMessage
}

func encounter() *proto.Encounter {
	mob := proto.MobType_MobTypeUnknown
	if *demon {
		mob = proto.MobType_MobTypeDemon
	}
	return &proto.Encounter{
		Duration: *duration, ExecuteProportion_20: .2, ExecuteProportion_25: .25, ExecuteProportion_35: .35,
		Targets: []*proto.Target{{Level: 63, MobType: mob, Stats: stats.Stats{stats.Armor: *targetArmor}.ToFloatArray()}},
	}
}

func request(player *proto.Player, count int, rng int64) *proto.RaidSimRequest {
	// Clone because agent initialization can modify options, buffs and consumes.
	return googleProto.Clone(&proto.RaidSimRequest{
		Raid:       core.SinglePlayerRaidProto(player, core.ForeverBuffs.Party, core.ForeverBuffs.Raid, core.ForeverBuffs.Debuffs),
		Encounter:  encounter(),
		SimOptions: &proto.SimOptions{Iterations: int32(count), RandomSeed: rng, Ruleset: proto.Ruleset_RulesetForever},
	}).(*proto.RaidSimRequest)
}

func run(b build, p *proto.Player, count int, rng int64) resultRow {
	if !b.allowsRace(p.Race) {
		panic(fmt.Errorf("%s: race %v is not available", b.Key, p.Race))
	}
	if err := loadTalents(b).validate(p.TalentsString); err != nil {
		panic(fmt.Errorf("%s: %w", b.Key, err))
	}
	if err := validateGear(p); err != nil {
		panic(err)
	}
	baseline, err := protojson.Marshal(p)
	if err != nil {
		panic(err)
	}
	p, hit, err := capHit(b, p)
	if err != nil {
		panic(err)
	}
	req := request(p, count, rng)
	_, rs, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	ps := rs.Parties[0].Players[0]
	var warnings []string
	for _, a := range append(ps.RotationStats.GetPrepullActions(), ps.RotationStats.GetPriorityList()...) {
		warnings = append(warnings, a.Warnings...)
	}
	// Fresh request keeps the saved replay independent of the stats probe.
	req = request(p, count, rng)
	saved, err := protojson.Marshal(req)
	if err != nil {
		panic(err)
	}
	result := core.RunRaidSim(req)
	if result.Error != nil {
		panic(result.Error.Message)
	}
	metrics, err := protojson.Marshal(result.RaidMetrics.Parties[0].Players[0])
	if err != nil {
		panic(err)
	}
	dps := result.RaidMetrics.Dps
	return resultRow{
		Key: b.Key, Build: b.Name, Race: raceName(p.Race), Faction: raceFaction(p.Race), Talents: p.TalentsString,
		DPS: dps.Avg, StdDev: dps.Stdev, StandardError: dps.Stdev / math.Sqrt(float64(result.IterationsDone)),
		OOMSeconds: result.RaidMetrics.Parties[0].Players[0].SecondsOomAvg,
		Iterations: result.IterationsDone, Hit: hit, Stats: ps.FinalStats.Stats, Sets: ps.Sets,
		Warnings: warnings, UnmodeledSetBonuses: unmodeledSetBonuses(p),
		BaselinePlayer: baseline, Request: saved, Metrics: metrics,
	}
}

func selectBuilds() []build {
	var selected []build
	for _, b := range builds() {
		if *buildFilter == "" || slices.Contains(strings.Split(*buildFilter, ","), b.Key) {
			selected = append(selected, b)
		}
	}
	if len(selected) == 0 {
		panic("no matching builds")
	}
	if len(selected) != 1 && (*talentsOverride != "" || *aplOverride != "" || *gearOverride != "" || *playerOverride != "") {
		panic("overrides require exactly one build")
	}
	return selected
}

func prepare(b build, race proto.Race) *proto.Player {
	p := b.player(race)
	if *playerOverride != "" {
		p = readPlayer(*playerOverride)
		p.Race = race
	}
	if *talentsOverride != "" {
		p.TalentsString = *talentsOverride
	}
	if *aplOverride != "" {
		p.Rotation = core.APLRotationFromJsonString(string(mustRead(*aplOverride)))
	}
	if *gearOverride != "" {
		p.Equipment = core.EquipmentSpecFromJsonString(string(mustRead(*gearOverride)))
	}
	if *distanceOverride >= 0 {
		p.DistanceFromTarget = *distanceOverride
	}
	if p.BonusStats == nil {
		p.BonusStats = &proto.UnitStats{}
	}
	bonus := stats.FromFloatArray(p.BonusStats.Stats)
	bonus[stats.SpellPower] += *spellPower
	p.BonusStats.Stats = bonus.ToFloatArray()
	p.ForeverTier1Bonuses = *tier1
	if *equipmentScale != 1 {
		p.EquipmentScale = *equipmentScale
	}
	return p
}

func writeResults(rows []resultRow) {
	if *output == "" {
		return
	}
	if err := os.MkdirAll(filepath.Dir(*output), 0755); err != nil {
		panic(err)
	}
	payload := struct {
		Duration, StartingArmor, AddedSpellPower float64
		ProvisionalArmor                         bool
		Tier1Bonuses                             bool
		EquipmentScale                           float64
		Mechanics                                map[string]string
		Results                                  []resultRow
	}{*duration, *targetArmor, *spellPower, false, *tier1, *equipmentScale, map[string]string{
		"autoAttacks":  "continuous-weapon-roles-no-caster-weaving",
		"hunterCasts":  "client-duration-no-added-windup",
		"energy":       "10-per-second-100ms-integration-general-haste-assumed",
		"eureka":       "charge-reserved-before-nested-effects",
		"omen":         "provisional-client-100pct-direct-procs-10s-icd-no-wrath-consumption",
		"healing":      "explicit-damage-only-no-inferred-conversion",
		"hotStreak":    "one-charge-next-pyro-completion-consumes-all-stacks",
		"clearcasting": "paid-base-cost-consumption-12536-one-second-icd",
		"sanctityAura": "excluded-no-current-trait-or-class-skill",
		"demonicBrand": "client-school-formulas-owner-power-at-hit-pet-multipliers-once-target-scoped",
		"weaponStones": "client-effects-exclusive-main-hand-imbue-spellstone-school-mask36",
		"naturesGrace": "10pct-cast-haste-separate-10pct-gcd-reduction-including-instants",
		"rage":         "inherited-damage-based-level60-and-offhand-normalization-unresolved",
	}, rows}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(*output+".json", append(data, '\n'), 0644); err != nil {
		panic(err)
	}
	f, err := os.Create(*output + ".csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Write([]string{"build", "race", "faction", "dps", "stdev", "standard_error", "iterations", "oom_seconds", "spell_hit_adjustment", "melee_hit_adjustment", "spell_hit_final", "melee_hit_final", "hit_budget_model", "hit_rating_delta", "offensive_budget_before", "budget_balance", "talents"})
	for _, r := range rows {
		w.Write([]string{r.Build, r.Race, r.Faction, fmt.Sprint(r.DPS), fmt.Sprint(r.StdDev), fmt.Sprint(r.StandardError),
			fmt.Sprint(r.Iterations), fmt.Sprint(r.OOMSeconds), fmt.Sprint(r.Hit.SpellAdded), fmt.Sprint(r.Hit.MeleeAdded),
			fmt.Sprint(r.Hit.SpellFinal), fmt.Sprint(r.Hit.MeleeFinal), r.Hit.Model,
			fmt.Sprint(r.Hit.RawHitDelta), fmt.Sprint(r.Hit.OffensiveBudgetBefore), fmt.Sprint(r.Hit.Balance), r.Talents})
	}
	w.Flush()
	if err := w.Error(); err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	if *searchGear && (*optimize || *equipmentScale != 1) {
		panic("gear search cannot be combined with talent search or scaled equipment")
	}
	if *equipmentScale <= 0 || math.IsInf(*equipmentScale, 0) || math.IsNaN(*equipmentScale) {
		panic("equipment-scale must be finite and positive")
	}
	if !slices.Contains([]string{"all", "alliance", "horde"}, strings.ToLower(*factionFilter)) {
		panic("faction must be all, alliance, or horde")
	}
	sim.RegisterAll()
	if *seedGear {
		writeSeedGear()
		return
	}
	var rows []resultRow
	var savedPlayers map[string]json.RawMessage
	if *baselineResults != "" {
		if *playerOverride != "" || *talentsOverride != "" || *aplOverride != "" || *gearOverride != "" || *optimize {
			panic("baseline-results cannot be combined with profile overrides or optimization")
		}
		var saved struct{ Results []resultRow }
		if err := json.Unmarshal(mustRead(*baselineResults), &saved); err != nil {
			panic(err)
		}
		savedPlayers = map[string]json.RawMessage{}
		for _, row := range saved.Results {
			key := row.Key + "/" + row.Race
			if _, exists := savedPlayers[key]; exists {
				panic("duplicate baseline: " + key)
			}
			savedPlayers[key] = row.BaselinePlayer
		}
	}
	for _, b := range selectBuilds() {
		for _, race := range b.races() {
			if !strings.EqualFold(*factionFilter, "all") && !strings.EqualFold(*factionFilter, raceFaction(race)) {
				continue
			}
			if *raceFilter != "" && *raceFilter != raceName(race) {
				continue
			}
			p := prepare(b, race)
			if savedPlayers != nil {
				data, ok := savedPlayers[b.Key+"/"+raceName(race)]
				if !ok {
					panic("missing saved baseline: " + b.Key + "/" + raceName(race))
				}
				p = &proto.Player{}
				if err := protojson.Unmarshal(data, p); err != nil {
					panic(err)
				}
				if p.Race != race || p.EquipmentScale != 0 && p.EquipmentScale != 1 {
					panic("saved baseline has a different race or already scaled gear")
				}
				p.ForeverTier1Bonuses = *tier1
				if *equipmentScale != 1 {
					p.EquipmentScale = *equipmentScale
				}
			}
			if *refreshEnchants {
				prepareGearEnchants(b, p)
			}
			if *searchGear {
				p = optimizeGear(b, p)
			}
			resultSeed := *seed
			if *optimize {
				p = optimizeTalents(b, p)
				resultSeed += 30000
			}
			r := run(b, p, *iterations, resultSeed)
			rows = append(rows, r)
			fmt.Printf("%-15s %-10s %8.2f DPS SE %.2f OOM %.2fs hit %.2f/%.2f talents %s warnings %d\n",
				b.Name, r.Race, r.DPS, r.StandardError, r.OOMSeconds, r.Hit.MeleeFinal, r.Hit.SpellFinal, r.Talents, len(r.Warnings))
		}
	}
	if len(rows) == 0 {
		panic("no matching races")
	}
	writeResults(rows)
}
