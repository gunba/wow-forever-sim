//go:build with_db

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// The requests in docs/apl-search-repros reproduce what APL searches found (docs/apl-search-findings.md); each is one
// logged fight.

func aplSearchRepro(t *testing.T, name string) *proto.RaidSimRequest {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("docs", "apl-search-repros", name))
	if err != nil {
		t.Fatal(err)
	}
	req := &proto.RaidSimRequest{}
	if err := protojson.Unmarshal(raw, req); err != nil {
		t.Fatal(err)
	}
	req.SimOptions.Iterations, req.SimOptions.Debug = 1, true
	return req
}

func runAPLSearchRepro(t *testing.T, req *proto.RaidSimRequest) *proto.RaidSimResult {
	t.Helper()
	result := core.RunRaidSim(req)
	if result.Error != nil {
		t.Fatal(result.Error.Message)
	}
	return result
}

// rotationWarnings are the warnings the UI shows for the player's rotation.
func rotationWarnings(req *proto.RaidSimRequest) []string {
	stats := core.ComputeStats(&proto.ComputeStatsRequest{Raid: req.Raid, Encounter: req.Encounter, Ruleset: req.SimOptions.Ruleset})
	rs := stats.RaidStats.Parties[0].Players[0].RotationStats
	var warnings []string
	for _, action := range append(rs.GetPrepullActions(), rs.GetPriorityList()...) {
		warnings = append(warnings, action.Warnings...)
	}
	return warnings
}

func playerCasts(result *proto.RaidSimResult, spellID int32) int32 {
	casts := int32(0)
	for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
		if action.Id.GetSpellId() == spellID {
			for _, target := range action.Targets {
				casts += target.Casts
			}
		}
	}
	return casts
}

func hasWarning(warnings []string, part string) bool {
	for _, w := range warnings {
		if strings.Contains(w, part) {
			return true
		}
	}
	return false
}

// Conditions on things the character lacks used to be dropped, which made their actions run every time. An aura it
// cannot have is never active, and a condition that cannot be evaluated at all stops its action.
func TestUnresolvedConditionsDoNotRunTheirActions(t *testing.T) {
	for _, c := range []struct {
		repro   string
		spell   int32
		max     int32
		warning string
	}{
		// Ice Lance on a Fingers of Frost proc the build lacks; the other Ice Lance entry is for the last 3 seconds.
		{"missing-aura-unconditional.json", 30455, 3, "No aura found"},
		// Soul Fire guarded by the target's health, which a duration-based target does not have.
		{"target-health-percent-20.json", 17924, 0, "cannot be evaluated"},
	} {
		req := aplSearchRepro(t, c.repro)
		if casts := playerCasts(runAPLSearchRepro(t, req), c.spell); casts > c.max {
			t.Errorf("%s: %d casts of %d, want at most %d", c.repro, casts, c.spell, c.max)
		}
		if !hasWarning(rotationWarnings(req), c.warning) {
			t.Errorf("%s: no %q warning", c.repro, c.warning)
		}
	}
}

// Autocast Other Cooldowns before the pull crashed the simulation; it is now refused with a warning.
func TestPrepullAutocastOtherCooldownsIsRefused(t *testing.T) {
	req := aplSearchRepro(t, "prepull-autocooldowns-crash.json")
	runAPLSearchRepro(t, req)
	if !hasWarning(rotationWarnings(req), "Autocast Other Cooldowns only works in the priority list") {
		t.Error("no warning for Autocast Other Cooldowns before the pull")
	}
}

// The current-pet mana values follow the demon summoned before the pull, as if it had been the configured one.
func TestWarlockPetManaFollowsTheActiveDemon(t *testing.T) {
	replaced := playerCasts(runAPLSearchRepro(t, aplSearchRepro(t, "pet-mana-after-replacement.json")), 11689)
	initial := playerCasts(runAPLSearchRepro(t, aplSearchRepro(t, "pet-mana-initial-imp.json")), 11689)
	if replaced == 0 || replaced != initial {
		t.Errorf("Life Tap on the summoned Imp's mana: %d casts, %d with the Imp configured", replaced, initial)
	}
}

// A Rogue can enter Stealth before the pull and open from it without spending Vanish.
func TestRogueStealthBeforeThePull(t *testing.T) {
	result := runAPLSearchRepro(t, aplSearchRepro(t, "rogue-1784-opener.json"))
	if playerCasts(result, 1784) != 1 || playerCasts(result, 11290) == 0 {
		t.Errorf("Stealth cast %d times and Garrote %d times, want Stealth once and a Garrote opener",
			playerCasts(result, 1784), playerCasts(result, 11290))
	}
}

// Whatever reaches the enemy before the pull, which in the game starts the fight early, is flagged on the pre-pull
// action responsible. Hunter's Mark, placing a fire totem whose first attack comes after the pull, and Life Tap, which
// is cast at the target but only affects the Warlock, are not.
func TestPrepullEffectsReachingTheEnemyAreFlagged(t *testing.T) {
	lifeTap := aplSearchRepro(t, "prepull-immolate.json")
	action := &proto.APLPrepullAction{}
	if err := protojson.Unmarshal([]byte(`{"action": {"castSpell": {"spellId": {"spellId": 11689}}}, "doAtValue": {"const": {"val": "-5s"}}}`), action); err != nil {
		t.Fatal(err)
	}
	lifeTap.Raid.Parties[0].Players[0].Rotation.PrepullActions = []*proto.APLPrepullAction{action}

	for _, c := range []struct {
		name    string
		req     *proto.RaidSimRequest
		spell   int32
		flagged bool
	}{
		// Hits at -3.5s.
		{"prepull-immolate.json", aplSearchRepro(t, "prepull-immolate.json"), 11668, true},
		// A curse, which puts an aura on the enemy without a hit, at -59s.
		{"prepull-bane-of-doom.json", aplSearchRepro(t, "prepull-bane-of-doom.json"), 603, true},
		// Cast at -3s, but haste brings the hit to -0.09s.
		{"prepull-starfire-haste.json", aplSearchRepro(t, "prepull-starfire-haste.json"), 9876, true},
		{"prepull-hunters-mark.json", aplSearchRepro(t, "prepull-hunters-mark.json"), 14325, false},
		{"prepull-searing-totem.json", aplSearchRepro(t, "prepull-searing-totem.json"), 10438, false},
		{"Life Tap at -5s", lifeTap, 11689, false},
	} {
		flagged := hasWarning(rotationWarnings(c.req), fmt.Sprintf("{SpellID: %d} reaches the enemy before the pull", c.spell))
		if flagged != c.flagged {
			t.Errorf("%s: flagged %v, want %v", c.name, flagged, c.flagged)
		}
	}
}

// mostActiveAtOnce is the most of the given auras the log shows on unit at the same time.
func mostActiveAtOnce(result *proto.RaidSimResult, unit string, ids []int32) int {
	active, most := map[string]bool{}, 0
	for _, line := range strings.Split(result.Logs, "\n") {
		for _, id := range ids {
			// The aura's id may carry a tag: {SpellID: 1237313, Tag: 1}.
			for _, aura := range []string{fmt.Sprintf("{SpellID: %d}", id), fmt.Sprintf("{SpellID: %d,", id)} {
				switch {
				case strings.Contains(line, "] ["+unit+"] Aura gained: "+aura):
					active[aura] = true
				case strings.Contains(line, "] ["+unit+"] Aura faded: "+aura):
					delete(active, aura)
				}
			}
		}
		most = max(most, len(active))
	}
	return most
}

// One caster's ranks of a DoT do not stack: a new rank replaces the others on the target.
func TestDotRanksReplaceEachOther(t *testing.T) {
	for _, c := range []struct {
		repro string
		ids   []int32
	}{
		{"corruption-multiple-ranks.json", []int32{172, 6222, 6223, 7648, 11671, 11672}},
		{"frostfire-multiple-ranks.json", []int32{401502, 1237312, 1237313}},
	} {
		result := runAPLSearchRepro(t, aplSearchRepro(t, c.repro))
		if most := mostActiveAtOnce(result, "Target 1", c.ids); most != 1 {
			t.Errorf("%s: up to %d ranks on the target at once, want 1", c.repro, most)
		}
	}
}

// Every learned rank of Arcane Wrath can be cast, not only the highest.
func TestArcaneWrathLowerRanksCanBeCast(t *testing.T) {
	if casts := playerCasts(runAPLSearchRepro(t, aplSearchRepro(t, "unavailable-arcane-wrath-rank.json")), 1326081); casts == 0 {
		t.Error("Arcane Wrath rank 3 was never cast")
	}
}

// A trap laid before the pull goes off when the fight starts, though the Hunter stands at range.
func TestTrapLaidBeforeThePullGoesOffAtThePull(t *testing.T) {
	req := aplSearchRepro(t, "hunter-trap-at-range.json")
	action := &proto.APLPrepullAction{}
	if err := protojson.Unmarshal([]byte(`{"action": {"castSpell": {"spellId": {"spellId": 14305}}}, "doAtValue": {"const": {"val": "-5s"}}}`), action); err != nil {
		t.Fatal(err)
	}
	rotation := req.Raid.Parties[0].Players[0].Rotation
	rotation.PrepullActions = append(rotation.PrepullActions, action)
	result := runAPLSearchRepro(t, req)
	if !strings.Contains(result.Logs, "[0.00] [Target 1] Aura gained: {SpellID: 14305}") {
		t.Error("the Immolation Trap laid at -5s did not go off at the pull")
	}
	if hasWarning(rotationWarnings(req), "{SpellID: 14305} reaches the enemy before the pull") {
		t.Error("the trap laid before the pull was flagged as reaching the enemy before it")
	}
}
