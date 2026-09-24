package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	googleProto "google.golang.org/protobuf/proto"
)

var (
	searchGear   = flag.Bool("search-gear", false, "coordinate search through the complete ilvl-65 pool and verified level-60+ catalog gear")
	gearScreen   = flag.Int("gear-screen", 100, "iterations per gear candidate")
	gearValidate = flag.Int("gear-validate", 1000, "independent iterations for each slot's shortlisted candidates")
	gearPasses   = flag.Int("gear-passes", 3, "maximum coordinate-search passes")
)

type gearTrial struct {
	Pass, Slot, Iterations int
	Seed                   int64
	Equipment              *proto.EquipmentSpec
	Profession2            proto.Profession
	DPS, StandardError     float64
	Stage                  string
	Warnings               []string
}

type gearSearchReport struct {
	Build, Race     string
	Baseline, Final resultRow
	Passes          int
	Converged       bool
	Trials          []gearTrial
	Accepted        []gearTrial
	PoolIDs         []int32
	Excluded        map[int32]string
}

func comparisonGearPool(b build, p *proto.Player) ([]core.Item, map[int32]string) {
	var source struct {
		Count int
		Items []struct {
			ID    int32 `json:"id"`
			Level int   `json:"level"`
		}
	}
	if err := json.Unmarshal(mustRead("assets/db_inputs/forever_ilvl65_items.json"), &source); err != nil {
		panic(err)
	}
	if source.Count != 706 || len(source.Items) != 706 {
		panic("incomplete ilvl-65 comparison pool")
	}
	var pool []core.Item
	excluded := map[int32]string{}
	catalog := readGearCatalog()
	ids := map[int32]bool{}
	for _, row := range source.Items {
		if row.Level != 65 {
			panic("comparison pool contains a different item level")
		}
		ids[row.ID] = true
	}
	// The supplied filter omits cloaks. The verified level-65 trinkets also
	// share an equipment limit, so compare lower-level trinkets for slot two.
	for id, record := range catalog {
		if record.Synthetic || record.ItemLevel >= 60 || core.ItemsByID[id].Type == proto.ItemType_ItemTypeTrinket {
			ids[id] = true
		}
	}
	// Keep original and resumed gear in the pool across interchangeable slots:
	// replacing trinket one must not remove it from consideration for slot two.
	for _, player := range []*proto.Player{b.player(p.Race), p} {
		for _, item := range player.Equipment.Items {
			if item.GetId() != 0 {
				ids[item.Id] = true
			}
		}
	}
	for id := range ids {
		item, exists := core.ItemsByID[id]
		if _, admitted := catalog[id]; !exists || !admitted {
			excluded[id] = "not in source-verified compiled equipment catalog"
		} else if catalog[id].Synthetic && !catalog[id].BenchmarkEligible {
			excluded[id] = "full-capacity passive trinket is a sensitivity case, not ranking gear"
		} else if !classCanEquip(p.Class, item) {
			excluded[id] = "class cannot equip"
		} else if catalog[id].Synthetic && item.ArmorType != 0 &&
			item.Type != proto.ItemType_ItemTypeBack &&
			item.ArmorType != classPreferredArmor(p.Class) {
			// The same projection exists in the class's ordinary armor type.
			// Avoid selecting a visually inappropriate lower-armor duplicate.
			excluded[id] = "modeled body armor has a class-appropriate equivalent"
		} else if err := gearReviewError(item); err != nil {
			excluded[id] = err.Error()
		} else if b.Key == "retribution_physical" && !physicalRetItemEligible(item) {
			excluded[id] = "physical Ret excludes spell-power and pure Intellect gear"
		} else {
			pool = append(pool, item)
		}
	}
	sort.Slice(pool, func(i, j int) bool { return pool[i].ID < pool[j].ID })
	return pool, excluded
}

// This row tests a physical equipment budget, not another Holy/Intellect
// item search. Mixed Strength/Intellect pieces remain available when their
// physical offensive stats justify them; pure caster gear does not.
func physicalRetItemEligible(item core.Item) bool {
	s := item.Stats
	if s[stats.SpellPower]+s[stats.SpellDamage]+s[stats.ArcanePower]+s[stats.FirePower]+
		s[stats.FrostPower]+s[stats.HolyPower]+s[stats.NaturePower]+s[stats.ShadowPower] > 0 {
		return false
	}
	physical := s[stats.Strength] + s[stats.Agility] + s[stats.AttackPower] + s[stats.MeleeHit] +
		s[stats.MeleeCrit] + s[stats.MeleeHaste]
	return s[stats.Intellect] == 0 || physical > 0
}

// Ordinary slots vary independently. Weapon layouts also cover transitions
// between two-handers and legal main/off-hand combinations.
func gearCandidates(b build, p *proto.Player, slot int, pool []core.Item) []*proto.Player {
	var candidates []*proto.Player
	appendCandidate := func(ids map[int]int32) {
		candidate := googleProto.Clone(p).(*proto.Player)
		for index, id := range ids {
			if candidate.Equipment.Items[index].GetId() != id {
				candidate.Equipment.Items[index] = &proto.ItemSpec{Id: id, Enchant: candidate.Equipment.Items[index].GetEnchant()}
			}
			if id == 0 {
				candidate.Equipment.Items[index] = &proto.ItemSpec{}
			}
			skill := readGearCatalog()[id].RequiredSkill
			profession := map[int32]proto.Profession{
				171: proto.Profession_Alchemy, 164: proto.Profession_Blacksmithing,
				333: proto.Profession_Enchanting, 202: proto.Profession_Engineering,
				165: proto.Profession_Leatherworking, 197: proto.Profession_Tailoring,
			}[skill]
			if profession != 0 && profession != candidate.Profession1 {
				candidate.Profession2 = profession
			}
		}
		prepareGearEnchants(b, candidate)
		if googleProto.Equal(candidate.Equipment, p.Equipment) || validateGear(candidate) != nil {
			return
		}
		candidates = append(candidates, candidate)
	}
	if slot != 14 {
		for _, item := range pool {
			appendCandidate(map[int]int32{slot: item.ID})
		}
		return candidates
	}
	twoHand := b.modelKey() == "arms" || b.modelKey() == "retribution" || b.modelKey() == "feral"
	dual := b.Class == proto.Class_ClassRogue || b.Key == "fury"
	weapons := append([]core.Item{}, pool...)
	for _, index := range []int{14, 15} {
		id := p.Equipment.Items[index].GetId()
		if id != 0 {
			found := false
			for _, item := range weapons {
				found = found || item.ID == id
			}
			if !found {
				weapons = append(weapons, core.ItemsByID[id])
			}
		}
	}
	for _, mh := range weapons {
		if mh.Type != proto.ItemType_ItemTypeWeapon || mh.HandType == proto.HandType_HandTypeOffHand {
			continue
		}
		if twoHand != (mh.HandType == proto.HandType_HandTypeTwoHand) && (twoHand || dual) {
			continue
		}
		if b.Key == "mutilate" && mh.WeaponType != proto.WeaponType_WeaponTypeDagger {
			continue
		}
		if mh.HandType == proto.HandType_HandTypeTwoHand {
			appendCandidate(map[int]int32{14: mh.ID, 15: 0})
			continue
		}
		for _, oh := range weapons {
			if oh.Type != proto.ItemType_ItemTypeWeapon || dual && oh.WeaponDamageMin == 0 {
				continue
			}
			if b.Key == "mutilate" && oh.WeaponType != proto.WeaponType_WeaponTypeDagger {
				continue
			}
			appendCandidate(map[int]int32{14: mh.ID, 15: oh.ID})
		}
	}
	return candidates
}

func optimizeGear(b build, initial *proto.Player) *proto.Player {
	if *output == "" || *gearScreen < 1 || *gearValidate < 1 || *gearPasses < 1 {
		panic("gear search requires -output and positive iteration/pass limits")
	}
	if !initial.ForeverTier1Bonuses {
		panic("gear comparison requires the fixed Tier 1 override")
	}
	p := googleProto.Clone(initial).(*proto.Player)
	prepareGearEnchants(b, p)
	pool, excluded := comparisonGearPool(b, p)
	report := gearSearchReport{Build: b.Key, Race: raceName(p.Race), Excluded: excluded}
	for _, item := range pool {
		report.PoolIDs = append(report.PoolIDs, item.ID)
	}
	report.Baseline = run(b, initial, *iterations, *seed)
	evaluate := func(candidate *proto.Player, pass, slot, count int, rng int64, stage string) resultRow {
		row := run(b, candidate, count, rng)
		report.Trials = append(report.Trials, gearTrial{
			Pass: pass, Slot: slot, Iterations: count, Seed: rng,
			Equipment:   googleProto.Clone(candidate.Equipment).(*proto.EquipmentSpec),
			Profession2: candidate.Profession2,
			DPS:         row.DPS, StandardError: row.StandardError, Stage: stage, Warnings: row.Warnings,
		})
		return row
	}
	for pass := 1; pass <= *gearPasses; pass++ {
		report.Passes = pass
		changed := false
		slots := []int{0, 2, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16}
		for _, slot := range append(slots, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34) {
			screenSeed := *seed + int64(pass*1000+slot)
			type screened struct {
				player *proto.Player
				row    resultRow
			}
			var shortlist []screened
			var candidates []*proto.Player
			if slot == 34 {
				for _, profession := range []proto.Profession{
					proto.Profession_Tailoring, proto.Profession_Leatherworking,
					proto.Profession_Blacksmithing, proto.Profession_Enchanting, proto.Profession_Alchemy,
				} {
					if profession == p.Profession1 || profession == p.Profession2 {
						continue
					}
					candidate := googleProto.Clone(p).(*proto.Player)
					candidate.Profession2 = profession
					if validateGear(candidate) == nil {
						candidates = append(candidates, candidate)
					}
				}
			} else if slot >= 17 {
				candidates = enchantCandidates(b, p, slot-17)
			} else {
				candidates = gearCandidates(b, p, slot, pool)
			}
			if len(candidates) == 0 {
				continue
			}
			current := evaluate(p, pass, slot, *gearScreen, screenSeed, "screen-baseline")
			for _, candidate := range candidates {
				row := evaluate(candidate, pass, slot, *gearScreen, screenSeed, "screen")
				if row.DPS > current.DPS && len(row.Warnings) == 0 {
					shortlist = append(shortlist, screened{candidate, row})
				}
			}
			sort.SliceStable(shortlist, func(i, j int) bool { return shortlist[i].row.DPS > shortlist[j].row.DPS })
			if len(shortlist) == 0 {
				continue
			}
			validationSeed := screenSeed + 1000000
			current = evaluate(p, pass, slot, *gearValidate, validationSeed, "validation-baseline")
			best, bestRow := p, current
			for _, candidate := range shortlist[:min(3, len(shortlist))] {
				row := evaluate(candidate.player, pass, slot, *gearValidate, validationSeed, "validation")
				// Conservative bound: paired covariance is not assumed.
				if row.DPS > bestRow.DPS && len(row.Warnings) == 0 &&
					row.DPS-current.DPS > 2*(row.StandardError+current.StandardError) {
					best, bestRow = candidate.player, row
				}
			}
			if best != p {
				p, changed = best, true
				report.Accepted = append(report.Accepted, gearTrial{
					Pass: pass, Slot: slot, Iterations: *gearValidate, Seed: validationSeed,
					Equipment:   googleProto.Clone(p.Equipment).(*proto.EquipmentSpec),
					Profession2: p.Profession2,
					DPS:         bestRow.DPS, StandardError: bestRow.StandardError, Stage: "accepted",
				})
				fmt.Printf("%s/%s pass %d slot %d: %.2f -> %.2f\n", b.Key, report.Race, pass, slot, current.DPS, bestRow.DPS)
			}
		}
		if !changed {
			report.Converged = true
			break
		}
	}
	report.Final = run(b, p, *iterations, *seed)
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
	return p
}
