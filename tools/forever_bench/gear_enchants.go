package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

var benchmarkEnchants = sync.OnceValue(func() []*proto.UIEnchant {
	var db proto.UIDatabase
	if err := protojson.Unmarshal(mustRead("assets/database/db.json"), &db); err != nil {
		panic(err)
	}
	var enchants []*proto.UIEnchant
	for _, enchant := range db.Enchants {
		// Enchanting recipes are crafting. Item-based enhancements also include
		// unavailable raid rewards, so those require a separate source list.
		crafted := enchant.ItemId == 0 && strings.HasPrefix(enchant.Name, "Enchant ")
		// These ItemIds identify Enchanting formula items, not consumable
		// reward enchants. In particular, the existing tank shield's +7
		// Stamina enchant must not be rejected as an unavailable raid reward.
		craftedShield := enchant.EnchantType == proto.EnchantType_EnchantTypeShield &&
			slices.Contains([]int32{13464, 13689, 13817, 13933, 20017, 20016}, enchant.SpellId)
		nonraidItem := slices.Contains([]int32{
			30, 32, 33, 34, 663, 664, 2523, // crafted scopes / counterweight
			15, 16, 17, 18, 2503, 8719, 8720, // crafted armor kits and new scope
			1483, 1503, 1504, 1505, 1506, 1507, 1508, 1509, 1510, // librams
			2488, 2543, 2544, 2545, // Argent Dawn and Dire Maul
		}, enchant.EffectId)
		if crafted || craftedShield || nonraidItem {
			enchants = append(enchants, enchant)
		}
	}
	sort.SliceStable(enchants, func(i, j int) bool {
		if enchants[i].EffectId == enchants[j].EffectId {
			return enchants[i].Type < enchants[j].Type
		}
		return enchants[i].EffectId < enchants[j].EffectId
	})
	return enchants
})

func enchantFits(p *proto.Player, item core.Item, e *proto.UIEnchant) bool {
	if e.Type != item.Type && !slices.Contains(e.ExtraTypes, item.Type) {
		return false
	}
	if len(e.ClassAllowlist) != 0 && !slices.Contains(e.ClassAllowlist, p.Class) {
		return false
	}
	if e.RequiredProfession != 0 && e.RequiredProfession != p.Profession1 && e.RequiredProfession != p.Profession2 {
		return false
	}
	if e.EffectId == 8719 && readGearCatalog()[item.ID].ItemLevel < 45 {
		return false
	}
	switch e.EnchantType {
	case proto.EnchantType_EnchantTypeTwoHand:
		return item.HandType == proto.HandType_HandTypeTwoHand
	case proto.EnchantType_EnchantTypeShield:
		return item.WeaponType == proto.WeaponType_WeaponTypeShield
	case proto.EnchantType_EnchantTypeStaff:
		return item.WeaponType == proto.WeaponType_WeaponTypeStaff
	}
	if item.Type == proto.ItemType_ItemTypeWeapon {
		return item.IsWeapon()
	}
	if item.Type == proto.ItemType_ItemTypeRanged {
		return slices.Contains([]proto.RangedWeaponType{
			proto.RangedWeaponType_RangedWeaponTypeBow,
			proto.RangedWeaponType_RangedWeaponTypeGun,
			proto.RangedWeaponType_RangedWeaponTypeCrossbow,
		}, item.RangedWeaponType)
	}
	return true
}

func legalEnchants(p *proto.Player, slot int) []*proto.UIEnchant {
	item := core.ItemsByID[p.Equipment.Items[slot].GetId()]
	if item.ID == 0 {
		return nil
	}
	var out []*proto.UIEnchant
	seen := map[int32]bool{}
	for _, e := range benchmarkEnchants() {
		if !seen[e.EffectId] && enchantFits(p, item, e) {
			out = append(out, e)
			seen[e.EffectId] = true
		}
	}
	return out
}

func enchantSeedScore(b build, e *proto.UIEnchant, slot int) float64 {
	s := stats.FromFloatArray(e.Stats)
	return seedScore(b, core.Item{Stats: s}, slot)
}

func physicalRetEnchantEligible(e *proto.UIEnchant) bool {
	s := stats.FromFloatArray(e.Stats)
	if s[stats.SpellPower]+s[stats.SpellDamage]+s[stats.ArcanePower]+s[stats.FirePower]+
		s[stats.FrostPower]+s[stats.HolyPower]+s[stats.NaturePower]+s[stats.ShadowPower] > 0 {
		return false
	}
	return s[stats.Intellect] <= s[stats.Strength]+s[stats.Agility]+s[stats.AttackPower]
}

// Fill every enchantable slot before the DPS search. The heuristic only
// supplies a starting point; native simulations compare the alternatives.
func prepareGearEnchants(b build, p *proto.Player) {
	for slot, spec := range p.Equipment.Items {
		choices := legalEnchants(p, slot)
		if len(choices) == 0 {
			spec.Enchant = 0
			continue
		}
		if slices.ContainsFunc(choices, func(e *proto.UIEnchant) bool { return e.EffectId == spec.Enchant }) {
			continue
		}
		best, score := int32(0), math.Inf(-1)
		for _, e := range choices {
			if b.Key == "retribution_physical" && !physicalRetEnchantEligible(e) {
				continue
			}
			s := stats.FromFloatArray(e.Stats)
			value := enchantSeedScore(b, e, slot)
			// Prefer real defensive benefits over an empty enhancement when
			// none of the available choices has a modeled damage stat.
			for _, v := range s {
				value += v * .0001
			}
			value += float64(e.Quality) * .000001
			if value > score {
				best, score = e.EffectId, value
			}
		}
		spec.Enchant = best
	}
}

func validateGearEnchants(p *proto.Player) error {
	for slot, spec := range p.Equipment.Items {
		if spec.GetEnchant() == 0 {
			continue
		}
		if !slices.ContainsFunc(legalEnchants(p, slot), func(e *proto.UIEnchant) bool { return e.EffectId == spec.Enchant }) {
			return fmt.Errorf("enchant %d is unavailable or illegal for equipment slot %d", spec.Enchant, slot)
		}
	}
	return nil
}

func enchantCandidates(b build, p *proto.Player, slot int) []*proto.Player {
	var out []*proto.Player
	for _, e := range legalEnchants(p, slot) {
		if e.EffectId == p.Equipment.Items[slot].GetEnchant() ||
			b.Key == "retribution_physical" && !physicalRetEnchantEligible(e) {
			continue
		}
		candidate := googleProto.Clone(p).(*proto.Player)
		candidate.Equipment.Items[slot].Enchant = e.EffectId
		out = append(out, candidate)
	}
	return out
}
