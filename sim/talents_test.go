package sim

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"unicode"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/druid"
	"github.com/wowsims/classic/sim/hunter"
	"github.com/wowsims/classic/sim/mage"
	"github.com/wowsims/classic/sim/paladin"
	"github.com/wowsims/classic/sim/priest"
	"github.com/wowsims/classic/sim/rogue"
	"github.com/wowsims/classic/sim/shaman"
	"github.com/wowsims/classic/sim/warlock"
	"github.com/wowsims/classic/sim/warrior"
)

// A talent string is positional: the i-th digit of a tree is the i-th talent of that
// tree in ui/core/talents/trees/<class>.json, and the sim reads it back through the
// proto's field numbers with FillTalentsProto. Three things have to agree for that to
// work, and nothing else checks them: the tree sizes each class package hardcodes, the
// order of the talents in the JSON, and the order of the fields in the proto. The UI
// also refuses a tree that is not in reading order, which it only says at page load.

type talentLocation struct {
	RowIdx int `json:"rowIdx"`
	ColIdx int `json:"colIdx"`
}

type talentNode struct {
	FieldName string          `json:"fieldName"`
	Location  talentLocation  `json:"location"`
	MaxPoints *int            `json:"maxPoints"`
	Prereq    *talentLocation `json:"prereqLocation"`

	// Read by TestConfirmedTalentRanksMatchTheSim, which fills the description's
	// placeholders from each rank's values to compare against what the game shows. A rank
	// holds numbers alongside the odd string, which is a pluralisation helper ("" or "s")
	// rather than a value, so the entries stay raw and only the numbers are read.
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Ranks       [][]json.RawMessage `json:"ranks"`

	// The picker warns that points spent here do nothing. Read by
	// TestUnsimulatedTalentsAreMarked, which checks the flag is on every talent the sim
	// does not implement.
	NotSimulated bool `json:"notSimulated"`
}

type talentTree struct {
	Name    string       `json:"name"`
	Talents []talentNode `json:"talents"`
}

type talentClass struct {
	name      string
	treeSizes [3]int
	talents   protoreflect.Message
	// The ui/ directories whose presets are written against this class's trees.
	specDirs []string
}

var talentClasses = []talentClass{
	{"druid", druid.TalentTreeSizes, (&proto.DruidTalents{}).ProtoReflect(), []string{"balance_druid", "feral_druid", "feral_tank_druid", "restoration_druid"}},
	{"hunter", hunter.TalentTreeSizes, (&proto.HunterTalents{}).ProtoReflect(), []string{"hunter"}},
	{"mage", mage.TalentTreeSizes, (&proto.MageTalents{}).ProtoReflect(), []string{"mage"}},
	{"paladin", paladin.TalentTreeSizes, (&proto.PaladinTalents{}).ProtoReflect(), []string{"holy_paladin", "protection_paladin", "retribution_paladin"}},
	{"priest", priest.TalentTreeSizes, (&proto.PriestTalents{}).ProtoReflect(), []string{"healing_priest", "shadow_priest", "smite_priest"}},
	{"rogue", rogue.TalentTreeSizes, (&proto.RogueTalents{}).ProtoReflect(), []string{"rogue"}},
	{"shaman", shaman.TalentTreeSizes, (&proto.ShamanTalents{}).ProtoReflect(), []string{"elemental_shaman", "enhancement_shaman", "restoration_shaman"}},
	{"warlock", warlock.TalentTreeSizes, (&proto.WarlockTalents{}).ProtoReflect(), []string{"warlock"}},
	{"warrior", warrior.TalentTreeSizes, (&proto.WarriorTalents{}).ProtoReflect(), []string{"tank_warrior", "warrior"}},
}

func loadTrees(t *testing.T, class string) []talentTree {
	data, err := os.ReadFile(filepath.Join("..", "ui", "core", "talents", "trees", class+".json"))
	if err != nil {
		t.Fatalf("%s: %v", class, err)
	}
	var trees []talentTree
	if err := json.Unmarshal(data, &trees); err != nil {
		t.Fatalf("%s: %v", class, err)
	}
	if len(trees) != 3 {
		t.Fatalf("%s: %d trees", class, len(trees))
	}
	return trees
}

func maxPoints(talent *talentNode) int {
	if talent.MaxPoints == nil {
		return 1
	}
	return *talent.MaxPoints
}

func TestTalentTreesMatchTheirProtos(t *testing.T) {
	for _, class := range talentClasses {
		trees := loadTrees(t, class.name)

		fields := class.talents.Descriptor().Fields()
		var protoOrder []string
		for i := 0; i < fields.Len(); i++ {
			field := fields.Get(i)
			protoOrder = append(protoOrder, field.JSONName())
		}

		var treeOrder []string
		for treeIdx, tree := range trees {
			if len(tree.Talents) != class.treeSizes[treeIdx] {
				t.Errorf("%s/%s: %d talents in the tree but TalentTreeSizes says %d", class.name, tree.Name, len(tree.Talents), class.treeSizes[treeIdx])
			}
			for i, talent := range tree.Talents {
				treeOrder = append(treeOrder, talent.FieldName)
				if i == 0 {
					continue
				}
				prev := tree.Talents[i-1].Location
				cur := talent.Location
				if cur.RowIdx < prev.RowIdx || (cur.RowIdx == prev.RowIdx && cur.ColIdx <= prev.ColIdx) {
					t.Errorf("%s/%s: %s is out of reading order", class.name, tree.Name, talent.FieldName)
				}
			}
		}

		if strings.Join(protoOrder, " ") != strings.Join(treeOrder, " ") {
			t.Errorf("%s: proto fields and tree talents disagree on order\n proto: %v\n trees: %v", class.name, protoOrder, treeOrder)
		}
	}
}

// No two talents may share a grid slot. The reading-order check above only notices a
// collision when the duplicate happens to sort out of order, so a pair landing in the
// same cell can otherwise pass silently. A wago.tools Trait-table export proposed exactly
// that for warrior Protection (Bastion onto Vitality's cell); see
// docs/talent-check-2026-09-17.md.
func TestTalentsDoNotShareASlot(t *testing.T) {
	for _, class := range talentClasses {
		for _, tree := range loadTrees(t, class.name) {
			seen := map[talentLocation]string{}
			for _, talent := range tree.Talents {
				if other, taken := seen[talent.Location]; taken {
					t.Errorf("%s/%s: %s and %s both sit at row %d, col %d",
						class.name, tree.Name, other, talent.FieldName,
						talent.Location.RowIdx, talent.Location.ColIdx)
					continue
				}
				seen[talent.Location] = talent.FieldName
			}
		}
	}
}

var talentsStringRegex = regexp.MustCompile(`talentsString: '([0-9-]*)'`)
var presetNameRegex = regexp.MustCompile(`makePresetTalents\(\s*'([^']+)'`)
var buildLinkRegex = regexp.MustCompile(`href="([a-z_]+)/\?build=([^"]+)"`)
var communityBuildRegex = regexp.MustCompile(`\[Spec\.Spec(\w+)\]:\s*\[([^\]]*)\]`)
var quotedNameRegex = regexp.MustCompile(`'([^']*)'`)

// Every build the UI ships has to be one a player could actually click together:
// no talent past its rank cap, five points spent above every row a point sits in,
// prerequisites filled, and no more than 51 in total.
func TestPresetBuildsAreLegal(t *testing.T) {
	for _, class := range talentClasses {
		trees := loadTrees(t, class.name)

		for _, dir := range class.specDirs {
			data, err := os.ReadFile(filepath.Join("..", "ui", dir, "presets.ts"))
			if err != nil {
				t.Fatalf("%s: %v", dir, err)
			}
			for _, match := range talentsStringRegex.FindAllStringSubmatch(string(data), -1) {
				checkBuild(t, dir, trees, match[1])
			}
		}
	}
}

func checkBuild(t *testing.T, dir string, trees []talentTree, build string) {
	total := 0
	for treeIdx, treeStr := range strings.Split(build, "-") {
		tree := trees[treeIdx]
		if len(treeStr) > len(tree.Talents) {
			t.Errorf("%s %q: %d digits for the %d talents of %s", dir, build, len(treeStr), len(tree.Talents), tree.Name)
			continue
		}

		points := make(map[talentLocation]int)
		perRow := make(map[int]int)
		for i := range tree.Talents {
			talent := &tree.Talents[i]
			p := 0
			if i < len(treeStr) {
				p = int(treeStr[i] - '0')
			}
			if p > maxPoints(talent) {
				t.Errorf("%s %q: %d points in %s, which has %d ranks", dir, build, p, talent.FieldName, maxPoints(talent))
			}
			points[talent.Location] = p
			perRow[talent.Location.RowIdx] += p
			total += p
		}

		spentAbove := 0
		for row := 0; row < 7; row++ {
			if perRow[row] > 0 && spentAbove < 5*row {
				t.Errorf("%s %q: %d points in row %d of %s with only %d spent above it", dir, build, perRow[row], row, tree.Name, spentAbove)
			}
			spentAbove += perRow[row]
		}

		for i := range tree.Talents {
			talent := &tree.Talents[i]
			if talent.Prereq == nil || points[talent.Location] == 0 {
				continue
			}
			for j := range tree.Talents {
				parent := &tree.Talents[j]
				if parent.Location == *talent.Prereq && points[parent.Location] < maxPoints(parent) {
					t.Errorf("%s %q: %s needs %s filled first", dir, build, talent.FieldName, parent.FieldName)
				}
			}
		}
	}
	if total > 51 {
		t.Errorf("%s %q: %d points", dir, build, total)
	}
}

// The landing page redirects to the generated review portal and exact profiles.
func TestLandingPageUsesReviewPortal(t *testing.T) {
	page, err := os.ReadFile(filepath.Join("..", "ui", "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	for _, link := range []string{`url=./review/`, `href="./review/"`} {
		if !strings.Contains(string(page), link) {
			t.Errorf("landing page missing review portal link %q", link)
		}
	}
}

// The sim pages carry the same list in their class dropdown, so a player can reach the
// community builds without going back to the landing page. ui/core cannot read the
// presets to find them - every page bundles ui/core, and a spec's presets drag in that
// whole sim - so ui/core/community_builds.ts writes the list out again, and nothing at
// runtime notices when the two copies drift: the dropdown just quietly stops offering a
// build, or offers one whose name no preset answers to and opens the sim on its
// defaults.
func TestSimTitleDropdownBuildsNamePresets(t *testing.T) {
	presets := loadPresetNames(t)
	dropdown := dropdownBuilds(t)

	for dir := range presets {
		if _, ok := dropdown[dir]; !ok {
			t.Errorf("ui/core/community_builds.ts has no entry for %s", dir)
		}
	}

	for dir, names := range dropdown {
		if _, ok := presets[dir]; !ok {
			t.Errorf("ui/core/community_builds.ts lists %s, which is not a spec with talent presets", dir)
			continue
		}
		for _, name := range names {
			if !presets[dir][name] {
				t.Errorf("ui/core/community_builds.ts gives %s a build named %q, which its presets do not have", dir, name)
			}
		}
	}
}

func loadPresetNames(t *testing.T) map[string]map[string]bool {
	t.Helper()

	presets := make(map[string]map[string]bool)
	for _, class := range talentClasses {
		for _, dir := range class.specDirs {
			data, err := os.ReadFile(filepath.Join("..", "ui", dir, "presets.ts"))
			if err != nil {
				t.Fatalf("%s: %v", dir, err)
			}
			presets[dir] = make(map[string]bool)
			for _, match := range presetNameRegex.FindAllStringSubmatch(string(data), -1) {
				presets[dir][match[1]] = true
			}
		}
	}
	return presets
}

// The same, as the sim pages' dropdown has them.
func dropdownBuilds(t *testing.T) map[string][]string {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("..", "ui", "core", "community_builds.ts"))
	if err != nil {
		t.Fatal(err)
	}

	entries := communityBuildRegex.FindAllStringSubmatch(string(data), -1)
	if len(entries) == 0 {
		t.Fatal("no specs in ui/core/community_builds.ts")
	}

	builds := make(map[string][]string)
	for _, entry := range entries {
		names := []string{}
		for _, name := range quotedNameRegex.FindAllStringSubmatch(entry[2], -1) {
			names = append(names, name[1])
		}
		builds[specDirName(entry[1])] = names
	}
	return builds
}

// 'BalanceDruid' -> 'balance_druid', the directory the sim is built into and the one
// getSpecSiteUrl points a build link at.
func specDirName(spec string) string {
	var dir strings.Builder
	for i, r := range spec {
		if i > 0 && unicode.IsUpper(r) {
			dir.WriteRune('_')
		}
		dir.WriteRune(unicode.ToLower(r))
	}
	return dir.String()
}
