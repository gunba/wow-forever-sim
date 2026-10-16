package dpswarrior

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// Two racial bugs reached master in a row - the undead had no Forever racial at all
// (#145), and the Skyborne carried theirs into Classic - because the DPS suites pick one
// or two races per spec and nothing ever builds the rest. The warrior is the one class
// every race can be, so sweeping it covers all ten under both rulesets.
//
// The eligible races are read out of the UI rather than copied, so a race added there is
// swept from the next run without touching this file.

var warriorRacesRegex = regexp.MustCompile(`(?m)^const warriorRaces = \[([^\]]*)\]`)
var skyborneRegex = regexp.MustCompile(`(?m)^const skyborne = \[([^\]]*)\]`)
var raceNameRegex = regexp.MustCompile(`Race\.(Race\w+)`)

// What each race's passives are worth in a stat sweep, per ruleset. Cooldowns need a
// rotation to fire and are left to the spec suites; these are the always-on effects,
// which is the kind that went missing in both bugs.
type racialExpectation struct {
	stat  proto.PseudoStat
	value float64
}

var foreverExpectations = map[proto.Race]racialExpectation{
	// Wind Blessed: 3% melee, ranged and cast haste since client 1.60.1.70425 (was 1%).
	proto.Race_RaceSkyborneHighOrder:  {proto.PseudoStat_PseudoStatMeleeSpeedMultiplier, 1.03},
	proto.Race_RaceSkyborneWindshaper: {proto.PseudoStat_PseudoStatMeleeSpeedMultiplier, 1.03},
}

// Under Classic the Skyborne do not exist, so nothing may be granted for them.
var classicExpectations = map[proto.Race]racialExpectation{
	proto.Race_RaceSkyborneHighOrder:  {proto.PseudoStat_PseudoStatMeleeSpeedMultiplier, 1},
	proto.Race_RaceSkyborneWindshaper: {proto.PseudoStat_PseudoStatMeleeSpeedMultiplier, 1},
}

func TestEveryWarriorRaceBuildsUnderBothRulesets(t *testing.T) {
	races := eligibleWarriorRaces(t)
	if len(races) < 10 {
		t.Errorf("expected every race to be warrior-eligible, got %d", len(races))
	}

	for _, race := range races {
		for _, ruleset := range []proto.Ruleset{proto.Ruleset_RulesetForever, proto.Ruleset_RulesetClassic} {
			t.Run(race.String()+"/"+ruleset.String(), func(t *testing.T) {
				stats := buildWarrior(t, race, ruleset)
				if stats == nil {
					return
				}

				expectations := foreverExpectations
				if ruleset == proto.Ruleset_RulesetClassic {
					expectations = classicExpectations
				}
				if want, ok := expectations[race]; ok {
					if got := stats[want.stat]; got != want.value {
						t.Errorf("%v = %v, want %v", want.stat, got, want.value)
					}
				}
			})
		}
	}
}

// Reads the warrior race list out of the UI so this test cannot drift from what players
// are actually offered.
func eligibleWarriorRaces(t *testing.T) []proto.Race {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("..", "..", "..", "ui", "core", "proto_utils", "utils.ts"))
	if err != nil {
		t.Fatal(err)
	}
	source := string(data)

	skyborne := ""
	if m := skyborneRegex.FindStringSubmatch(source); m != nil {
		skyborne = m[1]
	}
	match := warriorRacesRegex.FindStringSubmatch(source)
	if match == nil {
		t.Fatal("no warriorRaces list in ui/core/proto_utils/utils.ts; has the shape changed?")
	}

	var races []proto.Race
	body := strings.ReplaceAll(match[1], "...skyborne", skyborne)
	for _, name := range raceNameRegex.FindAllStringSubmatch(body, -1) {
		value, ok := proto.Race_value[name[1]]
		if !ok {
			t.Errorf("utils.ts names a race the proto does not have: %s", name[1])
			continue
		}
		races = append(races, proto.Race(value))
	}
	return races
}

// Returns the character's final pseudo stats. A race the sim cannot build is the failure
// this sweep exists to catch, so a panic is reported rather than allowed to abort the run.
func buildWarrior(t *testing.T, race proto.Race, ruleset proto.Ruleset) []float64 {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("building %v under %v panicked: %v", race, ruleset, r)
		}
	}()

	result := core.ComputeStats(&proto.ComputeStatsRequest{
		Raid: core.SinglePlayerRaidProto(&proto.Player{
			Class:         proto.Class_ClassWarrior,
			Race:          race,
			Equipment:     core.GetGearSet("../../../ui/warrior/gear_sets", "p0.bis").GearSet,
			TalentsString: P1Talents,
			Spec:          PlayerOptionsArms,
		}, nil, nil, nil),
		Ruleset: ruleset,
	})

	if len(result.RaidStats.Parties) == 0 || len(result.RaidStats.Parties[0].Players) == 0 {
		t.Errorf("building %v under %v produced no player stats", race, ruleset)
		return nil
	}
	return result.RaidStats.Parties[0].Players[0].FinalStats.PseudoStats
}
