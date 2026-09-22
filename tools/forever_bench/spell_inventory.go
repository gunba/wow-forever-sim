package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"slices"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

var spellInventory = flag.Bool("spell-inventory", false, "export registered spell IDs for the selected profiles without running fights")

type registeredSpellInventory struct {
	Build, Class, Race, Talents string
	SpellIDs                    []int32
	PetSpellIDs                 []int32
}

func inventorySpells(b build, player *proto.Player) registeredSpellInventory {
	req := request(player, 1, 1)
	env, _, _ := core.NewEnvironment(req.Raid, req.Encounter, proto.Ruleset_RulesetForever, false)
	ids, petIDs := []int32{}, []int32{}
	for _, unit := range env.Raid.AllUnits {
		for _, spell := range unit.Spellbook {
			if spell.SpellID > 0 {
				if unit.Type == core.PetUnit {
					petIDs = append(petIDs, spell.SpellID)
				} else {
					ids = append(ids, spell.SpellID)
				}
			}
		}
	}
	slices.Sort(ids)
	slices.Sort(petIDs)
	return registeredSpellInventory{
		Build: b.Key, Class: b.Class.String(), Race: raceName(player.Race),
		Talents: player.TalentsString, SpellIDs: slices.Compact(ids), PetSpellIDs: slices.Compact(petIDs),
	}
}

func writeSpellInventory(rows []registeredSpellInventory) {
	if len(rows) == 0 || *output == "" {
		panic("spell-inventory requires matching profiles and an output prefix")
	}
	if err := os.MkdirAll(filepath.Dir(*output), 0755); err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(*output+".json", append(data, '\n'), 0644); err != nil {
		panic(err)
	}
}
