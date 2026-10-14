// Prints a request's final player-0 stats (JSON array) using the simulator's ComputeStats.
// With a second argument "spellhit", prints each player-0 spell's talent and school spell-hit bonus in percent instead.
// "hpstat effects ids.json" prints, per id: item effect, weapon effect, enchant effect, enchant in the database.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	sim.RegisterAll()
	if os.Args[1] == "effects" {
		var ids []int32
		data, err := os.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(data, &ids); err != nil {
			panic(err)
		}
		out := map[string][]bool{}
		for _, id := range ids {
			_, known := core.EnchantsByEffectID[id]
			out[fmt.Sprint(id)] = []bool{core.HasItemEffect(id), core.HasWeaponEffect(id), core.HasEnchantEffect(id), known}
		}
		b, _ := json.Marshal(out)
		fmt.Println(string(b))
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		panic(err)
	}
	raw, ok := wrapper["request"]
	if !ok {
		raw = data
	}
	req := &proto.RaidSimRequest{}
	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(raw, req); err != nil {
		panic(err)
	}
	if len(os.Args) > 2 && os.Args[2] == "spellhit" {
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Reset()
		unit := sim.Raid.AllPlayerUnits[0]
		out := map[string]map[string]float64{}
		for _, spell := range unit.Spellbook {
			if spell.ActionID.SpellID == 0 {
				continue
			}
			bonus := (spell.BonusHitRating + unit.GetSchoolBonusHitChance(spell)) / core.SpellHitRatingPerHitChance
			out[fmt.Sprint(spell.ActionID.SpellID)] = map[string]float64{"school": float64(spell.SchoolIndex), "bonus_hit_pct": bonus}
		}
		b, _ := json.Marshal(out)
		fmt.Println(string(b))
		return
	}
	res := core.ComputeStats(&proto.ComputeStatsRequest{Raid: req.Raid, Encounter: req.Encounter, Ruleset: req.SimOptions.Ruleset})
	if res.ErrorResult != "" {
		panic(res.ErrorResult)
	}
	out, _ := json.Marshal(res.RaidStats.Parties[0].Players[0].FinalStats.Stats)
	fmt.Println(string(out))
}
