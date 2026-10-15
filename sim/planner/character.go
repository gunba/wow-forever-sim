package planner

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/warrior"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

//go:embed items.json.gz
var itemsGz []byte

type itemRec struct {
	Class    int                `json:"c"`
	Subclass int                `json:"s"`
	Inv      int                `json:"inv"`
	ILvl     int                `json:"il"`
	Quality  int                `json:"q"`
	Stats    map[string]float64 `json:"st"`
	Weapon   []float64          `json:"w"`
	Ranged   []float64          `json:"r"`
}

type itemTables struct {
	Items    map[string]itemRec              `json:"items"`
	Bonus    map[string][][2]float64         `json:"bonus"`
	RandProp map[string]map[string][]float64 `json:"randProp"`
}

var (
	itemsOnce sync.Once
	items     itemTables
)

func loadItems() {
	itemsOnce.Do(func() {
		r, err := gzip.NewReader(bytes.NewReader(itemsGz))
		if err != nil {
			panic(err)
		}
		if err := json.NewDecoder(r).Decode(&items); err != nil {
			panic(err)
		}
	})
}

// WoW equipment slot (1-19) -> simulator item slot.
var gameSlot = map[string]proto.ItemSlot{"1": proto.ItemSlot_ItemSlotHead, "2": proto.ItemSlot_ItemSlotNeck, "3": proto.ItemSlot_ItemSlotShoulder,
	"5": proto.ItemSlot_ItemSlotChest, "6": proto.ItemSlot_ItemSlotWaist, "7": proto.ItemSlot_ItemSlotLegs, "8": proto.ItemSlot_ItemSlotFeet,
	"9": proto.ItemSlot_ItemSlotWrist, "10": proto.ItemSlot_ItemSlotHands, "11": proto.ItemSlot_ItemSlotFinger1, "12": proto.ItemSlot_ItemSlotFinger2,
	"13": proto.ItemSlot_ItemSlotTrinket1, "14": proto.ItemSlot_ItemSlotTrinket2, "15": proto.ItemSlot_ItemSlotBack,
	"16": proto.ItemSlot_ItemSlotMainHand, "17": proto.ItemSlot_ItemSlotOffHand, "18": proto.ItemSlot_ItemSlotRanged}

var raceNames = map[string]proto.Race{"human": proto.Race_RaceHuman, "dwarf": proto.Race_RaceDwarf, "nightelf": proto.Race_RaceNightElf,
	"gnome": proto.Race_RaceGnome, "orc": proto.Race_RaceOrc, "tauren": proto.Race_RaceTauren, "troll": proto.Race_RaceTroll,
	"undead": proto.Race_RaceUndead, "scourge": proto.Race_RaceUndead, "skybornehighorder": proto.Race_RaceSkyborneHighOrder,
	"highorder": proto.Race_RaceSkyborneHighOrder, "skybornewindshaper": proto.Race_RaceSkyborneWindshaper, "windshaper": proto.Race_RaceSkyborneWindshaper}

func itemType(inv int) proto.ItemType {
	switch inv {
	case 1:
		return proto.ItemType_ItemTypeHead
	case 2:
		return proto.ItemType_ItemTypeNeck
	case 3:
		return proto.ItemType_ItemTypeShoulder
	case 5, 20:
		return proto.ItemType_ItemTypeChest
	case 6:
		return proto.ItemType_ItemTypeWaist
	case 7:
		return proto.ItemType_ItemTypeLegs
	case 8:
		return proto.ItemType_ItemTypeFeet
	case 9:
		return proto.ItemType_ItemTypeWrist
	case 10:
		return proto.ItemType_ItemTypeHands
	case 11:
		return proto.ItemType_ItemTypeFinger
	case 12:
		return proto.ItemType_ItemTypeTrinket
	case 16:
		return proto.ItemType_ItemTypeBack
	case 13, 14, 17, 21, 22, 23:
		return proto.ItemType_ItemTypeWeapon
	case 15, 25, 26, 28:
		return proto.ItemType_ItemTypeRanged
	}
	return proto.ItemType_ItemTypeUnknown
}

func weaponType(class, sub int) proto.WeaponType {
	if class == 4 && sub == 6 {
		return proto.WeaponType_WeaponTypeShield
	}
	switch sub {
	case 0, 1:
		return proto.WeaponType_WeaponTypeAxe
	case 4, 5:
		return proto.WeaponType_WeaponTypeMace
	case 6:
		return proto.WeaponType_WeaponTypePolearm
	case 7, 8:
		return proto.WeaponType_WeaponTypeSword
	case 10:
		return proto.WeaponType_WeaponTypeStaff
	case 13:
		return proto.WeaponType_WeaponTypeFist
	case 15:
		return proto.WeaponType_WeaponTypeDagger
	}
	return proto.WeaponType_WeaponTypeUnknown
}

func handType(inv int) proto.HandType {
	switch inv {
	case 13:
		return proto.HandType_HandTypeOneHand
	case 21:
		return proto.HandType_HandTypeMainHand
	case 22, 14, 23:
		return proto.HandType_HandTypeOffHand
	case 17:
		return proto.HandType_HandTypeTwoHand
	}
	return proto.HandType_HandTypeUnknown
}

func randPropIndex(inv int) int {
	switch inv {
	case 1, 5, 7, 17, 20:
		return 0
	case 3, 6, 8, 10, 12:
		return 1
	case 13, 21:
		return 3
	case 15, 25, 26:
		return 4
	}
	return 2
}

// suffixStats are the stats of an item's random-suffix bonus lists: allocation x the item level's budget.
func suffixStats(rec itemRec, bonus []int32) stats.Stats {
	var out stats.Stats
	q := map[int]string{2: "Good", 3: "Superior", 4: "Epic", 5: "Epic"}[rec.Quality]
	if q == "" {
		q = "Good"
	}
	budget := 0.0
	if rp, ok := items.RandProp[strconv.Itoa(rec.ILvl)]; ok {
		budget = rp[q][randPropIndex(rec.Inv)]
	}
	for _, b := range bonus {
		for _, sa := range items.Bonus[strconv.Itoa(int(b))] {
			out[int(sa[0])] += math.Floor(sa[1]/10000*budget + 0.5)
		}
	}
	return out
}

// CharacterPlayer turns a character message into a simulator player. Unknown items are reported in the warnings.
func CharacterPlayer(msg *CharacterMsg) (*proto.Player, []string, error) {
	loadItems()
	var warn []string
	if !strings.EqualFold(msg.Class, "WARRIOR") {
		return nil, nil, fmt.Errorf("class %s is not supported yet (Warrior only)", msg.Class)
	}
	race, ok := raceNames[strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(msg.Race, " ", ""), "-", ""))]
	if !ok {
		race = proto.Race_RaceHuman
		warn = append(warn, "unknown race "+msg.Race+", using Human")
	}
	// Talents -> the simulator's talent string.
	trees := [3][]byte{make([]byte, 17), make([]byte, 18), make([]byte, 18)}
	for i := range trees {
		for j := range trees[i] {
			trees[i][j] = '0'
		}
	}
	slots := traitEntries["WARRIOR"]
	for _, t := range msg.Talents {
		sl, ok := slots[t.Entry]
		if !ok {
			for _, cand := range slots {
				if cand.Node == t.Node {
					sl, ok = cand, true
					break
				}
			}
		}
		if !ok {
			warn = append(warn, fmt.Sprintf("unknown talent node %d entry %d", t.Node, t.Entry))
			continue
		}
		if sl.Index < len(trees[sl.Tree]) {
			trees[sl.Tree][sl.Index] = byte('0' + min(t.Rank, sl.MaxRank))
		}
	}
	talents := strings.Join([]string{strings.TrimRight(string(trees[0]), "0"), strings.TrimRight(string(trees[1]), "0"), strings.TrimRight(string(trees[2]), "0")}, "-")

	eq := &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 17)}
	for i := range eq.Items {
		eq.Items[i] = &proto.ItemSpec{}
	}
	db := &proto.SimDatabase{}
	keys := make([]string, 0, len(msg.Gear))
	for k := range msg.Gear {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		g := msg.Gear[k]
		slot, ok := gameSlot[k]
		if !ok || g.Item == 0 {
			continue
		}
		spec := &proto.ItemSpec{Id: g.Item}
		if _, known := core.EnchantsByEffectID[g.Enchant]; known && g.Enchant != 0 {
			spec.Enchant = g.Enchant
		}
		rec, have := items.Items[strconv.Itoa(int(g.Item))]
		_, simHas := core.ItemsByID[g.Item]
		if simHas && len(g.Bonus) == 0 {
			eq.Items[slot] = spec
			continue
		}
		if !have {
			if simHas {
				eq.Items[slot] = spec
				continue
			}
			warn = append(warn, fmt.Sprintf("unknown item %d in slot %s", g.Item, k))
			continue
		}
		var st stats.Stats
		for idx, v := range rec.Stats {
			if i, err := strconv.Atoi(idx); err == nil && i < len(st) {
				st[i] += v
			}
		}
		if len(g.Bonus) > 0 {
			st = st.Add(suffixStats(rec, g.Bonus))
		}
		item := &proto.SimItem{Id: g.Item, Name: fmt.Sprintf("item %d", g.Item), Type: itemType(rec.Inv), Stats: st.ToFloatArray(), Ilvl: int32(rec.ILvl)}
		if rec.Class == 2 || (rec.Class == 4 && rec.Subclass == 6) {
			item.Type = proto.ItemType_ItemTypeWeapon
			item.WeaponType = weaponType(rec.Class, rec.Subclass)
			item.HandType = handType(rec.Inv)
			if slot == proto.ItemSlot_ItemSlotRanged {
				item.Type = proto.ItemType_ItemTypeRanged
				item.RangedWeaponType = proto.RangedWeaponType_RangedWeaponTypeBow
			}
		}
		if len(rec.Weapon) == 3 {
			item.WeaponDamageMin, item.WeaponDamageMax, item.WeaponSpeed = rec.Weapon[0], rec.Weapon[1], rec.Weapon[2]
		} else if len(rec.Ranged) == 3 {
			item.WeaponDamageMin, item.WeaponDamageMax, item.WeaponSpeed = rec.Ranged[0], rec.Ranged[1], rec.Ranged[2]
		}
		if simHas {
			// Keep the simulator's own item (and its effects) but add the suffix stats as a separate item id.
			item.Id = 1_700_000_000 + g.Item
			spec.Id = item.Id
		}
		db.Items = append(db.Items, item)
		eq.Items[slot] = spec
	}
	player := &proto.Player{
		Name: "planner", Race: race, Class: proto.Class_ClassWarrior, Level: msg.Level, KnownSpells: msg.Spells,
		Equipment: eq, TalentsString: talents, Database: db, Consumes: &proto.Consumes{}, Buffs: &proto.IndividualBuffs{},
		Rotation: &proto.APLRotation{},
		Spec: &proto.Player_TankWarrior{TankWarrior: &proto.TankWarrior{Options: &proto.TankWarrior_Options{
			Stance: proto.WarriorStance_WarriorStanceBattle}}},
	}
	return player, warn, nil
}

// BuildCharacter builds the planner model for a character message by constructing the simulator's Warrior.
func BuildCharacter(msg *CharacterMsg) (*Model, CharacterAnswer) {
	ans := CharacterAnswer{Type: "character", Generation: msg.Generation, Abilities: []AbilityInfo{}, Warnings: []string{}}
	player, warn, err := CharacterPlayer(msg)
	ans.Warnings = append(ans.Warnings, warn...)
	if err != nil {
		ans.Error = err.Error()
		return nil, ans
	}
	model, err := buildFromPlayer(player)
	if err != nil {
		ans.Error = err.Error()
		return nil, ans
	}
	known := map[int32]bool{}
	for _, id := range msg.Spells {
		known[id] = true
	}
	for _, id := range msg.Spells {
		ab := model.ByGameID(id)
		if ab == nil {
			name, _ := ranks.ByID(id)
			if name == "" {
				ans.Abilities = append(ans.Abilities, AbilityInfo{Spell: id, Name: "", Modelled: false, Reason: "not an ability the planner chooses (passive, racial, utility or unknown)"})
			} else {
				ans.Abilities = append(ans.Abilities, AbilityInfo{Spell: id, Name: string(name), Modelled: false, Reason: "not available to the simulator for this character"})
			}
			continue
		}
		info := AbilityInfo{Spell: id, Name: string(ab.Name), Modelled: ab.Modelled, Reason: ab.Reason}
		if ab.Kind == KindStance && ab.Modelled {
			info.Reason = "chosen only as part of borrowing a stance"
		}
		ans.Abilities = append(ans.Abilities, info)
	}
	ans.OK = true
	return model, ans
}

func buildFromPlayer(player *proto.Player) (model *Model, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("building the character: %v", r)
		}
	}()
	level := core.PlayerLevel(player)
	ec := expectedCreatures[min(int(level)+1, len(expectedCreatures)-1)]
	raid := core.SinglePlayerRaidProto(player, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{})
	enc := &proto.Encounter{Duration: 60, Targets: []*proto.Target{{Level: level + 1, MobType: proto.MobType_MobTypeHumanoid,
		Stats: stats.Stats{stats.Armor: ec.armor, stats.Health: ec.health}.ToFloatArray(), MinBaseDamage: ec.autoAttackDPS * 2 / 1.17,
		DamageSpread: 0.3333, SwingSpeed: 2}}}
	env, _, _ := core.NewEnvironment(raid, enc, proto.Ruleset_RulesetForever, false)
	agent := env.Raid.Parties[0].Players[0]
	wa, ok := agent.(warrior.WarriorAgent)
	if !ok {
		return nil, fmt.Errorf("not a Warrior agent")
	}
	known := map[int32]bool{}
	for _, id := range player.KnownSpells {
		known[id] = true
	}
	if len(known) == 0 {
		known = nil
	}
	return BuildWarriorModel(wa.GetWarrior(), known), nil
}

// talentSlot is a talent's position in the simulator's talent string.
type talentSlot struct {
	Tree, Index int
	MaxRank     int32
	Node        int32
	Field       string
}
