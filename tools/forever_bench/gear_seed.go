package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

// Initial ranking only. Final choices must be compared in the five-minute sim.
func seedScore(b build, item core.Item, slot int) float64 {
	s := item.Stats
	var score float64
	if casterBuild(b) {
		power := s[stats.SpellPower] + s[stats.SpellDamage]
		switch b.Key {
		case "balance":
			power += .5 * (s[stats.NaturePower] + s[stats.ArcanePower])
		case "elemental", "stormcaller":
			power += .8*s[stats.NaturePower] + .2*s[stats.FirePower]
		case "fire":
			power += s[stats.FirePower]
		case "frost":
			power += s[stats.FrostPower]
		case "arcane":
			power += s[stats.ArcanePower]
		case "smite":
			power += s[stats.HolyPower]
		case "shadow":
			power += s[stats.ShadowPower]
		case "destruction":
			power += s[stats.FirePower]
		default:
			power += .8*s[stats.ShadowPower] + .2*s[stats.FirePower]
		}
		score = power + .4*s[stats.Intellect] + .1*s[stats.Spirit] + .7*s[stats.MP5] +
			12*(s[stats.MeleeCrit]+s[stats.SpellCrit]) + (70.0/6.0)*(s[stats.MeleeHit]+s[stats.SpellHit])
	} else {
		score = s[stats.AttackPower] + 2*s[stats.Strength] + 1.2*s[stats.Agility] +
			22*(s[stats.MeleeCrit]+s[stats.SpellCrit]) + 20*(s[stats.MeleeHit]+s[stats.SpellHit])
		if b.Class == proto.Class_ClassHunter {
			score = s[stats.RangedAttackPower] + 2.2*s[stats.Agility] + .2*s[stats.Intellect] +
				22*(s[stats.MeleeCrit]+s[stats.SpellCrit]) + 20*(s[stats.MeleeHit]+s[stats.SpellHit])
			if b.Key == "survival" {
				score += s[stats.Strength]
			}
		}
		if b.Key == "feral" {
			score += 1.3 * s[stats.Agility]
		}
		if b.Key == "enhancement" || b.Key == "retribution" {
			score += .4*(s[stats.SpellPower]+s[stats.SpellDamage]) + .2*s[stats.Intellect] + .5*s[stats.MP5]
		}
	}
	if item.SwingSpeed > 0 {
		dps := (item.WeaponDamageMin + item.WeaponDamageMax) / (2 * item.SwingSpeed)
		switch slot {
		case 14:
			if !casterBuild(b) && b.Key != "feral" && b.Key != "beast_mastery" && b.Key != "marksmanship" {
				score += 6*dps + 8*item.SwingSpeed
			}
		case 15:
			if !casterBuild(b) && b.Key != "beast_mastery" && b.Key != "marksmanship" {
				score += 3 * dps
			}
		case 16:
			if b.Key == "beast_mastery" || b.Key == "marksmanship" {
				score += 14 * dps
			}
		}
	}
	for _, source := range readGearCatalog()[item.ID].Sources {
		if source.Kind == "crafted" {
			score += .0001 // Prefer crafted when the initial stat score is tied.
			break
		}
	}
	if slot == 16 && ((b.Key == "balance" && item.ID == 249396) ||
		(b.Key == "feral" && item.ID == 220606) ||
		(b.Key == "retribution" && item.ID == 249442)) {
		score += 5
	}
	return score
}

func seedEquipment(b build, p *proto.Player) error {
	original := googleProto.Clone(p.Equipment).(*proto.EquipmentSpec)
	p.Equipment = &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 17)}
	for slot := range p.Equipment.Items {
		p.Equipment.Items[slot] = &proto.ItemSpec{}
		if slot < len(original.Items) {
			p.Equipment.Items[slot].Enchant = original.Items[slot].GetEnchant()
		}
	}
	// Existing sapper choices require Engineering. It supplies no free combat
	// stats, and also permits the catalog's engineering equipment.
	p.Profession1 = proto.Profession_Engineering
	var pool []core.Item
	for id := range readGearCatalog() {
		item, exists := core.ItemsByID[id]
		if exists && classCanEquip(p.Class, item) && gearReviewError(item) == nil {
			pool = append(pool, item)
		}
	}
	sort.Slice(pool, func(i, j int) bool { return pool[i].ID < pool[j].ID })
	for _, slot := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 16} {
		best, score := int32(0), math.Inf(-1)
		for _, item := range pool {
			p.Equipment.Items[slot].Id = item.ID
			if validateGearLayout(p, false) == nil {
				value := seedScore(b, item, slot)
				if value > score {
					best, score = item.ID, value
				}
			}
		}
		if best == 0 {
			return fmt.Errorf("%s: no reviewed item for slot %d", b.Key, slot)
		}
		p.Equipment.Items[slot].Id = best
	}
	bestMH, bestOH, score := int32(0), int32(0), math.Inf(-1)
	twoHand := b.Key == "arms" || b.Key == "retribution" || b.Key == "feral"
	dual := b.Class == proto.Class_ClassRogue || b.Key == "fury" || b.Key == "enhancement"
	offhands := append([]core.Item{{}}, pool...)
	for _, mh := range pool {
		if mh.Type != proto.ItemType_ItemTypeWeapon || mh.HandType == proto.HandType_HandTypeOffHand {
			continue
		}
		if twoHand != (mh.HandType == proto.HandType_HandTypeTwoHand) && (twoHand || dual) {
			continue
		}
		if b.Key == "mutilate" && mh.WeaponType != proto.WeaponType_WeaponTypeDagger {
			continue
		}
		p.Equipment.Items[14].Id = mh.ID
		for _, oh := range offhands {
			if (oh.ID == 0) != (mh.HandType == proto.HandType_HandTypeTwoHand) {
				continue
			}
			if oh.ID != 0 && oh.Type != proto.ItemType_ItemTypeWeapon {
				continue
			}
			if dual && oh.WeaponDamageMin == 0 {
				continue
			}
			if b.Key == "mutilate" && oh.WeaponType != proto.WeaponType_WeaponTypeDagger {
				continue
			}
			p.Equipment.Items[15].Id = oh.ID
			if validateGear(p) != nil {
				continue
			}
			value := seedScore(b, mh, 14) + seedScore(b, oh, 15)
			if value > score {
				bestMH, bestOH, score = mh.ID, oh.ID, value
			}
		}
	}
	if bestMH == 0 {
		return fmt.Errorf("%s: no reviewed weapon layout", b.Key)
	}
	p.Equipment.Items[14].Id, p.Equipment.Items[15].Id = bestMH, bestOH
	if bestOH == 0 {
		p.Equipment.Items[15] = &proto.ItemSpec{}
	}
	return validateGear(p)
}

func writeSeedGear() {
	for _, b := range selectBuilds() {
		p := b.player(b.races()[0])
		if err := seedEquipment(b, p); err != nil {
			panic(err)
		}
		data, err := (protojson.MarshalOptions{Indent: "  "}).Marshal(p.Equipment)
		if err != nil {
			panic(err)
		}
		path := filepath.Join("ui", b.Dir, "gear_sets", "forever_"+b.Key+".gear.json")
		if err := os.WriteFile(path, append(data, '\n'), 0644); err != nil {
			panic(err)
		}
		fmt.Println(path)
	}
}
