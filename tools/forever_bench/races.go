package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"sync"

	"github.com/wowsims/classic/sim/core/proto"
)

type raceDefinition struct {
	Proto, Name, Faction string
	Classes              []string
}

var raceCatalog = sync.OnceValue(func() map[proto.Race]raceDefinition {
	var data struct {
		Races []raceDefinition
	}
	if err := json.Unmarshal(mustRead("assets/db_inputs/forever_races.json"), &data); err != nil {
		panic(err)
	}
	out := make(map[proto.Race]raceDefinition)
	for _, race := range data.Races {
		id, ok := proto.Race_value[race.Proto]
		if !ok {
			panic(fmt.Sprintf("unknown race %q", race.Proto))
		}
		out[proto.Race(id)] = race
	}
	return out
})

func raceName(race proto.Race) string {
	if r, ok := raceCatalog()[race]; ok {
		return r.Name
	}
	panic(fmt.Sprintf("unsupported race %v", race))
}

func raceFaction(race proto.Race) string {
	if r, ok := raceCatalog()[race]; ok {
		return r.Faction
	}
	panic(fmt.Sprintf("unsupported race %v", race))
}

func (b build) races() []proto.Race {
	// Keep the existing representative Horde race first for seeded gear and fixtures.
	order := []proto.Race{
		proto.Race_RaceOrc, proto.Race_RaceTauren, proto.Race_RaceTroll, proto.Race_RaceUndead, proto.Race_RaceSkyborneWindshaper,
		proto.Race_RaceHuman, proto.Race_RaceDwarf, proto.Race_RaceNightElf, proto.Race_RaceGnome, proto.Race_RaceSkyborneHighOrder,
	}
	var eligible []proto.Race
	for _, race := range order {
		for _, class := range raceCatalog()[race].Classes {
			if id, ok := proto.Class_value["Class"+class]; ok && proto.Class(id) == b.Class {
				eligible = append(eligible, race)
			}
		}
	}
	return eligible
}

func (b build) allowsRace(race proto.Race) bool {
	return slices.Contains(b.races(), race)
}
