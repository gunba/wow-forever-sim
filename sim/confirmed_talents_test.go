package sim

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// The sourced rank ledger retains older client values and identifies later
// note/client-derived updates. This check catches display drift; matching the
// ledger does not independently verify server-script behavior. Update source
// provenance alongside changed values in assets/confirmed_talents.json.
type confirmedTalents struct {
	Generated string                            `json:"generated"`
	Talents   map[string]map[string][][]float64 `json:"talents"`
}

var placeholderRegex = regexp.MustCompile(`\{(\d+)\}`)
var numberRegex = regexp.MustCompile(`\d+(\.\d+)?`)

func TestConfirmedTalentRanksMatchTheSim(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "assets", "confirmed_talents.json"))
	if err != nil {
		t.Fatal(err)
	}
	var confirmed confirmedTalents
	if err := json.Unmarshal(data, &confirmed); err != nil {
		t.Fatal(err)
	}
	if len(confirmed.Talents) == 0 {
		t.Fatal("assets/confirmed_talents.json has no talents; has the refresh tool changed shape?")
	}

	compared := 0
	for _, class := range talentClasses {
		byName := map[string][][]float64{}
		for className, talents := range confirmed.Talents {
			if strings.EqualFold(className, class.name) {
				for name, ranks := range talents {
					byName[normalizeTalentName(name)] = ranks
				}
			}
		}
		if len(byName) == 0 {
			continue
		}

		for _, tree := range loadTrees(t, class.name) {
			for i := range tree.Talents {
				talent := &tree.Talents[i]
				ranks, ok := byName[normalizeTalentName(talent.Name)]
				if !ok {
					continue
				}

				for rank, values := range ranks {
					if rank >= len(talent.Ranks) {
						break
					}
					compared++
					simNumbers := numbersIn(fillPlaceholders(talent.Description, talent.Ranks[rank]))
					if !sameNumbers(simNumbers, values) {
						t.Errorf("%s %s rank %d: the sim reads %v, but the game shows %v",
							class.name, talent.Name, rank+1, simNumbers, values)
					}
				}
			}
		}
	}

	if compared == 0 {
		t.Fatal("no confirmed ranks lined up with a talent in the trees; the names have probably drifted")
	}
	t.Logf("checked %d confirmed ranks against the sim (source generated %s)", compared, confirmed.Generated)
}

// Talent names are compared loosely because the two sources punctuate differently.
func normalizeTalentName(name string) string {
	var b strings.Builder
	lastSpace := true
	for _, r := range strings.ToLower(name) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			lastSpace = false
		} else if !lastSpace {
			b.WriteRune(' ')
			lastSpace = true
		}
	}
	return strings.TrimSpace(b.String())
}

// The trees hold one description with {0}-style placeholders and a value per rank. A value
// is usually a number but occasionally a pluralisation string, which contributes nothing to
// compare and is substituted as-is.
func fillPlaceholders(description string, values []json.RawMessage) string {
	return placeholderRegex.ReplaceAllStringFunc(description, func(match string) string {
		var idx int
		fmt.Sscanf(match, "{%d}", &idx)
		if idx >= len(values) {
			return match
		}
		var number float64
		if err := json.Unmarshal(values[idx], &number); err == nil {
			return formatNumber(number)
		}
		var text string
		if err := json.Unmarshal(values[idx], &text); err == nil {
			return text
		}
		return match
	})
}

func formatNumber(v float64) string {
	if v == float64(int64(v)) {
		return fmt.Sprintf("%d", int64(v))
	}
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", v), "0"), ".")
}

func numbersIn(text string) []float64 {
	var out []float64
	for _, match := range numberRegex.FindAllString(text, -1) {
		var v float64
		fmt.Sscanf(match, "%g", &v)
		out = append(out, v)
	}
	return out
}

// A rank states the same numbers even when the two sources order them differently, so the
// comparison is on the multiset rather than the sequence.
//
// The two also pick different units for the same duration - Elusiveness reads "90 sec"
// here and "1.5 min" on the site - so a value that matches another once multiplied by 60
// counts as equal. That costs a little strictness, since it would also accept a genuine
// 60-fold error, but no talent value is plausibly 60 times another.
func sameNumbers(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	x := append([]float64(nil), a...)
	y := append([]float64(nil), b...)
	sort.Float64s(x)
	sort.Float64s(y)
	for i := range x {
		if x[i] == y[i] || x[i] == y[i]*60 || y[i] == x[i]*60 {
			continue
		}
		return false
	}
	return true
}
