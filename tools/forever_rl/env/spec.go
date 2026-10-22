// Package env is a reinforcement-learning environment over the simulator: one episode chooses a talent build point by
// point (optional), then plays one benchmark fight in the real engine, choosing every action. The reward is damage.
package env

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

// Config selects the build and how episodes run.
type Config struct {
	Repo           string  `json:"repo"`           // repository root
	Results        string  `json:"results"`        // results file with complete requests (relative to Repo)
	Build          string  `json:"build"`          // build key, e.g. "fire"
	Race           string  `json:"race"`           // race as in the results ("Human")
	LearnTalents   bool    `json:"learnTalents"`   // choose the talents in the episode
	Talents        string  `json:"talents"`        // fixed talents when not learning them (default: the row's)
	FightsPerBuild int     `json:"fightsPerBuild"` // fights played on one chosen build before choosing again
	Seed           int64   `json:"seed"`
	RewardScale    float64 `json:"rewardScale"`    // reward = damage * RewardScale
	DiscoverBuilds int     `json:"discoverBuilds"` // random builds sampled to find every spell and aura a build can have
	CachedSims     int     `json:"cachedSims"`     // simulations kept per environment, one per talent build
	Duration       float64 `json:"duration"`       // fight length override in seconds (0: the request's)
	Threads        int     `json:"threads"`        // worker goroutines for a vector of environments
	LabeledRands   bool    `json:"labeledRands"`   // one random stream per roll label, so replays with one changed decision stay aligned
	// FullChoices gives the agent every choice a player has: major cooldowns (otherwise automatic), unqueueing a
	// queued on-next-swing attack, and longer waits.
	FullChoices bool `json:"fullChoices"`
	// KeepWaitOnly asks the agent even when waiting is its only legal action (default: those points are skipped).
	KeepWaitOnly bool `json:"keepWaitOnly"`
	// PollSeconds: after a plain wait the agent is asked again as soon as anything it sees changes (an action becomes
	// usable or not, a whole point of rage/energy or 1% of mana, an aura appearing or fading), and otherwise only after
	// this long. 0 asks at every engine poll (50 ms).
	PollSeconds float64 `json:"pollSeconds"`
	// Rotation replaces the published APL (the teacher and the source of the agent's actions): APLRotation JSON.
	Rotation string `json:"rotation"`
	// TeacherRotation replaces only the teacher, after the agent's actions are found from the published (or Rotation)
	// APL, so the spaces stay those a trained policy expects. Its entries for actions the agent lacks count as waits.
	TeacherRotation string `json:"teacherRotation"`
}

func (c *Config) defaults() {
	if c.Results == "" {
		c.Results = "artifacts/modelled_gear/forever_dps_5min.json"
	}
	if c.FightsPerBuild <= 0 {
		c.FightsPerBuild = 1
	}
	if c.RewardScale == 0 {
		c.RewardScale = 1e-3
	}
	if c.DiscoverBuilds == 0 {
		c.DiscoverBuilds = 40
	}
	if c.CachedSims <= 0 {
		c.CachedSims = 4
	}
}

const (
	kindAuto = iota
	kindTeacher
	kindAgent
)

type auraKey struct {
	Target bool
	Label  string
}

// auraMeta is what an APL needs to name a tracked aura and read its observation features back in game units.
type auraMeta struct {
	ID        core.ActionID
	Duration  time.Duration
	MaxStacks int32
}

// Spec is everything the environments of one build share. It is read-only once built.
type Spec struct {
	Cfg         Config
	request     *proto.RaidSimRequest
	teacher     *proto.APLRotation
	candidate   *proto.APLRotation
	configKind  []int // per candidate config entry
	configAgent []int // per candidate config entry: its agent action (teachers: the agent action of the same spell, or -1)
	trees       *talentTrees
	spellNames  map[int32]string
	itemNames   map[int32]string
	basePoints  []int
	BaseTalents string
	agentSpells []core.ActionID
	cooldowns   map[core.ActionID]bool // agent spells that are otherwise automatic major cooldowns
	macros      []*proto.APLAction     // the teacher's other entries (move, channel, sequences...) without their conditions
	macroNames  []string
	queueAuras  map[core.ActionID]bool
	waits       []float64 // seconds of the agent's longer waits, before the plain wait
	auras       []auraKey
	auraMeta    map[auraKey]auraMeta
	featureAPL  []any // per observation feature, its APL value and scale (nil: none)
	actionAPL   []any // per agent action, its APL action (nil: plain wait)
	Duration    float64

	T, K, ObsDim, ActDim int // talents, agent actions (spells + wait), observation and action sizes
	ActionNames          []string
	FeatureNames         []string
	TalentNames          []string
	fightOffset          int
}

var registerOnce sync.Once

type resultRow struct {
	Key, Race string
	Request   json.RawMessage
}

// NewSpec loads the build's benchmark request and finds the spells and auras its talent builds can have.
func NewSpec(cfg Config) (*Spec, error) {
	cfg.defaults()
	registerOnce.Do(sim.RegisterAll)
	data, err := os.ReadFile(filepath.Join(cfg.Repo, cfg.Results))
	if err != nil {
		return nil, err
	}
	var results struct{ Results []resultRow }
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}
	var row *resultRow
	for i := range results.Results {
		if results.Results[i].Key == cfg.Build && (cfg.Race == "" || results.Results[i].Race == cfg.Race) {
			row = &results.Results[i]
			break
		}
	}
	if row == nil {
		return nil, fmt.Errorf("no %s/%s row in %s", cfg.Build, cfg.Race, cfg.Results)
	}
	req := &proto.RaidSimRequest{}
	if err := protojson.Unmarshal(row.Request, req); err != nil {
		return nil, err
	}
	if cfg.Duration > 0 {
		req.Encounter.Duration = cfg.Duration
	}
	req.Encounter.DurationVariation = 0
	if req.SimOptions == nil {
		req.SimOptions = &proto.SimOptions{}
	}
	// Keep the ruleset (Forever) and every other option; only the iteration count and seed change.
	req.SimOptions.Iterations, req.SimOptions.RandomSeed = 1, cfg.Seed
	req.SimOptions.UseLabeledRands = cfg.LabeledRands
	player := req.Raid.Parties[0].Players[0]
	class := strings.ToLower(strings.TrimPrefix(player.Class.String(), "Class"))
	trees, err := loadTalentTrees(cfg.Repo, class)
	if err != nil {
		return nil, err
	}
	if cfg.Rotation != "" {
		rot := &proto.APLRotation{}
		if err := protojson.Unmarshal([]byte(cfg.Rotation), rot); err != nil {
			return nil, fmt.Errorf("rotation: %w", err)
		}
		player.Rotation = rot
	}
	s := &Spec{Cfg: cfg, request: req, teacher: player.Rotation, trees: trees, BaseTalents: player.TalentsString,
		Duration: req.Encounter.Duration}
	if s.basePoints, err = trees.decode(player.TalentsString); err != nil {
		return nil, err
	}
	if cfg.Talents == "" {
		s.Cfg.Talents = player.TalentsString
	}
	s.spellNames, s.itemNames = loadSpellNames(cfg.Repo)
	if err := s.discover(); err != nil {
		return nil, err
	}
	s.buildCandidate()
	s.layout()
	if cfg.TeacherRotation != "" {
		rot := &proto.APLRotation{}
		if err := protojson.Unmarshal([]byte(cfg.TeacherRotation), rot); err != nil {
			return nil, fmt.Errorf("teacherRotation: %w", err)
		}
		s.teacher = rot
		s.candidate = s.assemble(rot, false)
	}
	return s, nil
}

func (s *Spec) newSim(talents string, rotation *proto.APLRotation, seed int64) (*core.Simulation, *core.Character) {
	req := googleProto.Clone(s.request).(*proto.RaidSimRequest)
	p := req.Raid.Parties[0].Players[0]
	p.TalentsString = talents
	p.Rotation = rotation
	req.SimOptions.RandomSeed = seed
	sim := core.NewSim(req, simsignals.Signals{})
	return sim, sim.Raid.Parties[0].Players[0].GetCharacter()
}

// randomBuild spends every point at random among the legal choices.
func (s *Spec) randomBuild(rng *rand.Rand) []int {
	points := make([]int, len(s.trees.Nodes))
	legal := make([]int, 0, len(points))
	for spent := 0; spent < 51; spent++ {
		legal = legal[:0]
		for i := range points {
			if s.trees.canAdd(points, i) {
				legal = append(legal, i)
			}
		}
		if len(legal) == 0 {
			break
		}
		points[legal[rng.Intn(len(legal))]]++
	}
	return points
}

type spellGroup struct {
	name string
	code int32
	id   core.ActionID
}

// Spell names from the database group the ranks of one spell; spells it does not name group by spell code or id.
func loadSpellNames(repo string) (map[int32]string, map[int32]string) {
	spells, items := map[int32]string{}, map[int32]string{}
	data, err := os.ReadFile(filepath.Join(repo, "assets/database/db.json"))
	if err != nil {
		return spells, items
	}
	var db struct {
		SpellIcons []struct {
			ID   int32  `json:"id"`
			Name string `json:"name"`
		} `json:"spellIcons"`
		ItemIcons []struct {
			ID   int32  `json:"id"`
			Name string `json:"name"`
		} `json:"itemIcons"`
	}
	if json.Unmarshal(data, &db) == nil {
		for _, x := range db.SpellIcons {
			spells[x.ID] = x.Name
		}
		for _, x := range db.ItemIcons {
			items[x.ID] = x.Name
		}
	}
	return spells, items
}

// Spells that never deal damage or restore resources in a boss fight.
var utilitySpells = map[string]bool{"Counterspell": true, "Frost Nova": true, "Blink": true, "Polymorph": true, "Remove Lesser Curse": true,
	"Ice Block": true, "Mana Shield": true, "Fire Ward": true, "Frost Ward": true, "Slow Fall": true, "Kick": true, "Pummel": true,
	"Shield Bash": true, "Earth Shock (interrupt)": true, "Taunt": true, "Growl": true, "Feint": true, "Vanish": true}

func (s *Spec) groupOf(spell *core.Spell) spellGroup {
	if spell.ActionID.SpellID != 0 {
		if name, ok := s.spellNames[spell.ActionID.SpellID]; ok {
			return spellGroup{name: name, id: core.ActionID{Tag: spell.ActionID.Tag}}
		}
	}
	if spell.SpellCode != 0 {
		return spellGroup{code: spell.SpellCode, id: core.ActionID{Tag: spell.ActionID.Tag}}
	}
	return spellGroup{id: spell.ActionID}
}

// deepBuild fills one tree first (31 to 41 points), then spends the rest at random.
func (s *Spec) deepBuild(rng *rand.Rand, tree int) []int {
	points := make([]int, len(s.trees.Nodes))
	target := 31 + rng.Intn(11)
	legal := make([]int, 0, len(points))
	for spent := 0; spent < 51; spent++ {
		legal = legal[:0]
		inTree := spent < target
		for i, n := range s.trees.Nodes {
			if s.trees.canAdd(points, i) && (!inTree || n.Tree == tree) {
				legal = append(legal, i)
			}
		}
		if len(legal) == 0 && inTree {
			for i := range points {
				if s.trees.canAdd(points, i) {
					legal = append(legal, i)
				}
			}
		}
		if len(legal) == 0 {
			break
		}
		points[legal[rng.Intn(len(legal))]]++
	}
	return points
}

// discover builds the published build and random ones, and keeps every castable APL spell (top rank of each) and
// every aura with a duration on the player or the target.
func (s *Spec) discover() error {
	// A fixed seed: every environment of a build gets the same spaces whatever its own seed (evaluation environments
	// use other seeds). It is the seed training environments used before (1,000,011 + 7), so their models still fit.
	rng := rand.New(rand.NewSource(1_000_018))
	builds := [][]int{s.basePoints}
	for i := 0; i < s.Cfg.DiscoverBuilds; i++ {
		if i%4 == 3 {
			builds = append(builds, s.randomBuild(rng))
		} else {
			builds = append(builds, s.deepBuild(rng, i%len(s.trees.Sizes)))
		}
	}
	spells := map[spellGroup]*core.Spell{}
	ranks := map[spellGroup]int{}
	s.auraMeta = map[auraKey]auraMeta{}
	s.queueAuras = map[core.ActionID]bool{}
	auras := map[auraKey]bool{}
	// Cooldowns the published rotation leaves to Autocast Other Cooldowns stay automatic: listing them as agent
	// actions would take them out of the automatic set and away from the teacher.
	autocast := map[core.ActionID]bool{}
	for _, points := range builds {
		sim, char := s.newSim(s.trees.encode(points), s.teacher, 1)
		for _, id := range char.GetMajorCooldownIDs() {
			autocast[core.ProtoToActionID(id)] = true
		}
		for _, spell := range char.Spellbook {
			if !spell.Flags.Matches(core.SpellFlagAPL) || utilitySpells[s.spellNames[spell.ActionID.SpellID]] {
				continue
			}
			g := s.groupOf(spell)
			if cur, ok := spells[g]; !ok || spell.Rank > ranks[g] || (spell.Rank == ranks[g] && spell.ActionID.SpellID > cur.ActionID.SpellID) {
				spells[g], ranks[g] = spell, spell.Rank
			}
		}
		for _, aura := range char.GetAuras() {
			if aura.Duration == core.NeverExpires && aura.ActionID.Tag != 0 && strings.Contains(aura.Label, "Queue") {
				s.queueAuras[aura.ActionID] = true
			}
		}
		for _, unit := range []*core.Unit{&char.Unit, sim.Encounter.TargetUnits[0]} {
			for _, aura := range unit.GetAuras() {
				if aura.Duration > 0 && aura.Duration != core.NeverExpires && aura.Label != "" {
					k := auraKey{Target: unit != &char.Unit, Label: aura.Label}
					auras[k] = true
					if _, ok := s.auraMeta[k]; !ok {
						s.auraMeta[k] = auraMeta{ID: aura.ActionID, Duration: aura.Duration, MaxStacks: aura.MaxStacks}
					}
				}
			}
		}
	}
	// The teacher's own spells, items included, must be agent actions too.
	for _, item := range s.teacher.PriorityList {
		if cast := item.Action.GetCastSpell(); cast != nil {
			id := core.ProtoToActionID(cast.SpellId)
			found := false
			for _, spell := range spells {
				if spell.ActionID == id {
					found = true
				}
			}
			if !found {
				spells[spellGroup{id: id}] = &core.Spell{ActionID: id}
			}
		}
	}
	s.cooldowns = map[core.ActionID]bool{}
	for _, spell := range spells {
		if !autocast[spell.ActionID] {
			s.agentSpells = append(s.agentSpells, spell.ActionID)
		} else if s.Cfg.FullChoices {
			s.agentSpells = append(s.agentSpells, spell.ActionID)
			s.cooldowns[spell.ActionID] = true
		}
	}
	if s.Cfg.FullChoices {
		// Major cooldowns the spellbook does not flag for the APL (items, racials) are agent actions too.
		for id := range autocast {
			found := false
			for _, a := range s.agentSpells {
				if a == id {
					found = true
				}
			}
			if !found {
				s.agentSpells = append(s.agentSpells, id)
				s.cooldowns[id] = true
			}
		}
	}
	sort.Slice(s.agentSpells, func(i, j int) bool { return s.agentSpells[i].String() < s.agentSpells[j].String() })
	for k := range auras {
		s.auras = append(s.auras, k)
	}
	sort.Slice(s.auras, func(i, j int) bool {
		if s.auras[i].Target != s.auras[j].Target {
			return !s.auras[i].Target
		}
		return s.auras[i].Label < s.auras[j].Label
	})
	return nil
}

// buildCandidate makes the rotation every environment runs (see assemble) and settles the agent's actions.
func (s *Spec) buildCandidate() {
	s.candidate = s.assemble(s.teacher, true)
}

// assemble makes the rotation every environment runs: the automatic cooldown entry, the teacher's own entries (they
// only say what the teacher would do), then one unconditioned entry per agent action. With grow, teacher actions the
// agent lacks become new agent actions (macros); without it the action set is fixed and such entries count as waits.
func (s *Spec) assemble(teacher *proto.APLRotation, grow bool) *proto.APLRotation {
	s.configKind, s.configAgent = nil, nil
	c := &proto.APLRotation{Type: proto.APLRotation_TypeAPL, PrepullActions: teacher.PrepullActions}
	add := func(item *proto.APLListItem, kind, agent int) {
		c.PriorityList = append(c.PriorityList, item)
		s.configKind = append(s.configKind, kind)
		s.configAgent = append(s.configAgent, agent)
	}
	add(&proto.APLListItem{Action: &proto.APLAction{Action: &proto.APLAction_AutocastOtherCooldowns{AutocastOtherCooldowns: &proto.APLActionAutocastOtherCooldowns{}}}}, kindAuto, -1)
	nSpells := len(s.agentSpells)
	for _, item := range teacher.PriorityList {
		if item.Hide || item.Action.GetAutocastOtherCooldowns() != nil {
			continue
		}
		agent := -1
		switch item.Action.Action.(type) {
		case *proto.APLAction_CastSpell:
			id := core.ProtoToActionID(item.Action.GetCastSpell().SpellId)
			for k, a := range s.agentSpells {
				if a == id {
					agent = k
				}
			}
		case *proto.APLAction_Wait, *proto.APLAction_WaitUntil, *proto.APLAction_Schedule:
			// The agent's own wait.
		default:
			macro := &proto.APLAction{Action: googleProto.Clone(item.Action).(*proto.APLAction).Action}
			for k, m := range s.macros {
				if googleProto.Equal(m, macro) {
					agent = nSpells + k
				}
			}
			if agent < 0 && grow {
				s.macros = append(s.macros, macro)
				s.macroNames = append(s.macroNames, macroName(macro, s))
				agent = nSpells + len(s.macros) - 1
			}
		}
		add(item, kindTeacher, agent)
	}
	for k, id := range s.agentSpells {
		add(&proto.APLListItem{Action: &proto.APLAction{Action: &proto.APLAction_CastSpell{CastSpell: &proto.APLActionCastSpell{SpellId: id.ToProto()}}}}, kindAgent, k)
	}
	if grow && s.Cfg.FullChoices {
		ids := make([]core.ActionID, 0, len(s.queueAuras))
		for id := range s.queueAuras {
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i].String() < ids[j].String() })
		for _, id := range ids {
			s.macros = append(s.macros, &proto.APLAction{Action: &proto.APLAction_CancelAura{CancelAura: &proto.APLActionCancelAura{AuraId: id.ToProto()}}})
			s.macroNames = append(s.macroNames, "unqueue:"+s.actionName(core.ActionID{SpellID: id.SpellID}))
		}
		s.waits = []float64{0.25, 1}
	}
	for k, m := range s.macros {
		add(&proto.APLListItem{Action: m}, kindAgent, nSpells+k)
	}
	return c
}

func macroName(m *proto.APLAction, s *Spec) string {
	switch a := m.Action.(type) {
	case *proto.APLAction_Move:
		return "move:" + strings.Join(strings.Fields(protojson.Format(a.Move.RangeFromTarget)), "")
	case *proto.APLAction_ChannelSpell:
		return "channel:" + s.actionName(core.ProtoToActionID(a.ChannelSpell.SpellId))
	}
	j := protojson.Format(m)
	if len(j) > 60 {
		j = j[:60]
	}
	return "macro:" + strings.Join(strings.Fields(j), "")
}

func (s *Spec) layout() {
	s.T = len(s.trees.Nodes)
	s.K = len(s.agentSpells) + len(s.macros) + len(s.waits) + 1
	s.ActDim = s.T + s.K
	for _, n := range s.trees.Nodes {
		s.TalentNames = append(s.TalentNames, n.Name)
		s.ActionNames = append(s.ActionNames, "talent:"+n.Name)
	}
	for _, id := range s.agentSpells {
		s.ActionNames = append(s.ActionNames, "cast:"+s.actionName(id))
	}
	s.ActionNames = append(s.ActionNames, s.macroNames...)
	for _, w := range s.waits {
		s.ActionNames = append(s.ActionNames, fmt.Sprintf("wait:%gs", w))
	}
	s.ActionNames = append(s.ActionNames, "wait")
	f := []string{"phase"}
	for _, n := range s.trees.Nodes {
		f = append(f, "talent:"+n.Name)
	}
	for _, name := range s.trees.Names {
		f = append(f, "spent:"+name)
	}
	f = append(f, "spent:total")
	s.fightOffset = len(f)
	f = append(f, "elapsed", "remaining", "remaining_s", "execute20", "execute25", "execute35", "gcd", "mana_frac", "mana",
		"rage", "energy", "combo", "distance", "mh_swing", "oh_swing", "ranged_swing", "damage_rate")
	for _, id := range s.agentSpells {
		for _, p := range []string{"known", "ready", "cooldown", "cast_time", "cost"} {
			f = append(f, s.actionName(id)+":"+p)
		}
	}
	for _, name := range s.macroNames {
		f = append(f, name+":ready")
	}
	for _, a := range s.auras {
		owner := "self"
		if a.Target {
			owner = "target"
		}
		for _, p := range []string{"active", "remaining", "stacks"} {
			f = append(f, owner+":"+a.Label+":"+p)
		}
	}
	s.FeatureNames = f
	s.ObsDim = len(f)
	s.mapToAPL()
}

// mapToAPL records, for each observation feature, the APL value that reads the same quantity and the factor from
// the feature to the value's units, and for each agent action the APL action that performs it. A learned policy's
// conditions on features then compile to APL conditions that mean the same thing.
func (s *Spec) mapToAPL() {
	raw := func(m googleProto.Message) json.RawMessage {
		b, _ := protojson.Marshal(m)
		return b
	}
	val := func(v *proto.APLValue, kind string, scale float64) any {
		return map[string]any{"value": raw(v), "kind": kind, "scale": scale}
	}
	self := &proto.UnitReference{Type: proto.UnitReference_Self}
	target := &proto.UnitReference{Type: proto.UnitReference_CurrentTarget}
	swing := func(attack proto.APLValueAutoTimeToNext_AttackType, sw proto.APLValueAutoSwingTime_SwingType) any {
		return val(&proto.APLValue{Value: &proto.APLValue_Math{Math: &proto.APLValueMath{Op: proto.APLValueMath_OpDiv,
			Lhs: &proto.APLValue{Value: &proto.APLValue_AutoTimeToNext{AutoTimeToNext: &proto.APLValueAutoTimeToNext{AutoType: attack}}},
			Rhs: &proto.APLValue{Value: &proto.APLValue_AutoSwingTime{AutoSwingTime: &proto.APLValueAutoSwingTime{AutoType: sw}}}}}}, "num", 1)
	}
	exec := func(t proto.APLValueIsExecutePhase_ExecutePhaseThreshold) any {
		return val(&proto.APLValue{Value: &proto.APLValue_IsExecutePhase{IsExecutePhase: &proto.APLValueIsExecutePhase{Threshold: t}}}, "bool", 1)
	}
	fixed := map[string]any{
		"elapsed":      val(&proto.APLValue{Value: &proto.APLValue_CurrentTimePercent{CurrentTimePercent: &proto.APLValueCurrentTimePercent{}}}, "num", 1),
		"remaining":    val(&proto.APLValue{Value: &proto.APLValue_RemainingTimePercent{RemainingTimePercent: &proto.APLValueRemainingTimePercent{}}}, "num", 1),
		"remaining_s":  val(&proto.APLValue{Value: &proto.APLValue_RemainingTime{RemainingTime: &proto.APLValueRemainingTime{}}}, "time", 300),
		"execute20":    exec(proto.APLValueIsExecutePhase_E20),
		"execute25":    exec(proto.APLValueIsExecutePhase_E25),
		"execute35":    exec(proto.APLValueIsExecutePhase_E35),
		"gcd":          val(&proto.APLValue{Value: &proto.APLValue_GcdTimeToReady{GcdTimeToReady: &proto.APLValueGCDTimeToReady{}}}, "time", 1),
		"mana_frac":    val(&proto.APLValue{Value: &proto.APLValue_CurrentManaPercent{CurrentManaPercent: &proto.APLValueCurrentManaPercent{}}}, "num", 1),
		"mana":         val(&proto.APLValue{Value: &proto.APLValue_CurrentMana{CurrentMana: &proto.APLValueCurrentMana{}}}, "num", 10000),
		"rage":         val(&proto.APLValue{Value: &proto.APLValue_CurrentRage{CurrentRage: &proto.APLValueCurrentRage{}}}, "num", 100),
		"energy":       val(&proto.APLValue{Value: &proto.APLValue_CurrentEnergy{CurrentEnergy: &proto.APLValueCurrentEnergy{}}}, "num", 100),
		"combo":        val(&proto.APLValue{Value: &proto.APLValue_CurrentComboPoints{CurrentComboPoints: &proto.APLValueCurrentComboPoints{}}}, "num", 5),
		"mh_swing":     swing(proto.APLValueAutoTimeToNext_MainHand, proto.APLValueAutoSwingTime_MainHand),
		"oh_swing":     swing(proto.APLValueAutoTimeToNext_OffHand, proto.APLValueAutoSwingTime_OffHand),
		"ranged_swing": swing(proto.APLValueAutoTimeToNext_Ranged, proto.APLValueAutoSwingTime_Ranged),
	}
	s.featureAPL = make([]any, 0, len(s.FeatureNames))
	for i := 0; i < s.fightOffset; i++ {
		s.featureAPL = append(s.featureAPL, nil) // phase, talents, points spent: fixed for a fight
	}
	for _, name := range s.FeatureNames[s.fightOffset : s.fightOffset+17] {
		s.featureAPL = append(s.featureAPL, fixed[name])
	}
	for _, id := range s.agentSpells {
		sid := id.ToProto()
		s.featureAPL = append(s.featureAPL,
			val(&proto.APLValue{Value: &proto.APLValue_SpellIsKnown{SpellIsKnown: &proto.APLValueSpellIsKnown{SpellId: sid}}}, "bool", 1),
			val(&proto.APLValue{Value: &proto.APLValue_SpellCanCast{SpellCanCast: &proto.APLValueSpellCanCast{SpellId: sid}}}, "bool", 1),
			val(&proto.APLValue{Value: &proto.APLValue_SpellTimeToReady{SpellTimeToReady: &proto.APLValueSpellTimeToReady{SpellId: sid}}}, "time", 30),
			val(&proto.APLValue{Value: &proto.APLValue_SpellCastTime{SpellCastTime: &proto.APLValueSpellCastTime{SpellId: sid}}}, "time", 3.5),
			// Cost over current mana, as observed (0 without a mana bar; there the feature is constant and unused).
			val(&proto.APLValue{Value: &proto.APLValue_Math{Math: &proto.APLValueMath{Op: proto.APLValueMath_OpDiv,
				Lhs: &proto.APLValue{Value: &proto.APLValue_SpellCurrentCost{SpellCurrentCost: &proto.APLValueSpellCurrentCost{SpellId: sid}}},
				Rhs: &proto.APLValue{Value: &proto.APLValue_CurrentMana{CurrentMana: &proto.APLValueCurrentMana{}}}}}}, "num", 1))
	}
	for range s.macros {
		s.featureAPL = append(s.featureAPL, nil)
	}
	for _, k := range s.auras {
		m, ok := s.auraMeta[k]
		if !ok {
			s.featureAPL = append(s.featureAPL, nil, nil, nil)
			continue
		}
		unit := self
		if k.Target {
			unit = target
		}
		aid := m.ID.ToProto()
		s.featureAPL = append(s.featureAPL,
			val(&proto.APLValue{Value: &proto.APLValue_AuraIsActive{AuraIsActive: &proto.APLValueAuraIsActive{SourceUnit: unit, AuraId: aid}}}, "bool", 1),
			val(&proto.APLValue{Value: &proto.APLValue_AuraRemainingTime{AuraRemainingTime: &proto.APLValueAuraRemainingTime{SourceUnit: unit, AuraId: aid}}}, "time", m.Duration.Seconds()))
		if m.MaxStacks > 0 {
			s.featureAPL = append(s.featureAPL,
				val(&proto.APLValue{Value: &proto.APLValue_AuraNumStacks{AuraNumStacks: &proto.APLValueAuraNumStacks{SourceUnit: unit, AuraId: aid}}}, "num", float64(m.MaxStacks)))
		} else {
			s.featureAPL = append(s.featureAPL, nil)
		}
	}
	s.actionAPL = make([]any, 0, s.K)
	for _, id := range s.agentSpells {
		s.actionAPL = append(s.actionAPL, raw(&proto.APLAction{Action: &proto.APLAction_CastSpell{CastSpell: &proto.APLActionCastSpell{SpellId: id.ToProto()}}}))
	}
	for _, m := range s.macros {
		s.actionAPL = append(s.actionAPL, raw(m))
	}
	for _, w := range s.waits {
		s.actionAPL = append(s.actionAPL, raw(&proto.APLAction{Action: &proto.APLAction_Wait{Wait: &proto.APLActionWait{
			Duration: &proto.APLValue{Value: &proto.APLValue_Const{Const: &proto.APLValueConst{Val: fmt.Sprintf("%gs", w)}}}}}}))
	}
	s.actionAPL = append(s.actionAPL, nil)
}

func (s *Spec) actionName(id core.ActionID) string {
	if id.ItemID != 0 {
		if n, ok := s.itemNames[id.ItemID]; ok {
			return n
		}
	} else if n, ok := s.spellNames[id.SpellID]; ok {
		if id.Tag != 0 {
			return fmt.Sprintf("%s#%d", n, id.Tag)
		}
		return n
	}
	return id.String()
}

// PlainDPS runs the published rotation with no policy through the same per-iteration path, for parity checks.
func (s *Spec) PlainDPS(fights int) float64 {
	sim, char := s.newSim(s.Cfg.Talents, s.teacher, s.Cfg.Seed)
	total := 0.0
	for i := 0; i < fights; i++ {
		sim.RunIteration(s.Cfg.Seed + int64(i))
		total += char.Metrics.IterationDamage() / sim.Duration.Seconds()
	}
	return total / float64(fights)
}

// FullDPS runs the published request through the normal simulation path, for parity checks.
func (s *Spec) FullDPS(fights int) float64 {
	req := googleProto.Clone(s.request).(*proto.RaidSimRequest)
	req.SimOptions.Iterations, req.SimOptions.RandomSeed = int32(fights), s.Cfg.Seed
	res := core.RunRaidSim(req)
	if res.Error != nil {
		panic(res.Error.Message)
	}
	return res.RaidMetrics.Parties[0].Players[0].Dps.Avg
}

// TeacherMap lists each teacher entry with the agent action it maps to.
func (s *Spec) TeacherMap() []string {
	var out []string
	for i, kind := range s.configKind {
		if kind != kindTeacher {
			continue
		}
		item := s.candidate.PriorityList[i]
		target := "wait"
		if a := s.configAgent[i]; a >= 0 {
			target = s.ActionNames[s.T+a]
		}
		desc := "other"
		if cast := item.Action.GetCastSpell(); cast != nil {
			desc = s.actionName(core.ProtoToActionID(cast.SpellId))
		}
		out = append(out, fmt.Sprintf("%d %s -> %s", i, desc, target))
	}
	return out
}

func (s *Spec) talentMax() []int {
	out := make([]int, len(s.trees.Nodes))
	for i, n := range s.trees.Nodes {
		out[i] = n.Max
	}
	return out
}

func (s *Spec) talentRows() []int {
	out := make([]int, len(s.trees.Nodes))
	for i, n := range s.trees.Nodes {
		out[i] = n.Row
	}
	return out
}

func (s *Spec) talentTreeIdx() []int {
	out := make([]int, len(s.trees.Nodes))
	for i, n := range s.trees.Nodes {
		out[i] = n.Tree
	}
	return out
}

func (s *Spec) talentPrereqs() []int {
	out := make([]int, len(s.trees.Nodes))
	for i, n := range s.trees.Nodes {
		out[i] = n.Prereq
	}
	return out
}

// SpecJSON describes the spaces for the learner.
func (s *Spec) SpecJSON() []byte {
	out, _ := json.Marshal(map[string]any{
		"obsDim": s.ObsDim, "actDim": s.ActDim, "talents": s.T, "agentActions": s.K, "actionNames": s.ActionNames,
		"featureNames": s.FeatureNames, "talentNames": s.TalentNames, "baseTalents": s.BaseTalents, "duration": s.Duration,
		"learnTalents": s.Cfg.LearnTalents, "fixedTalents": s.Cfg.Talents, "build": s.Cfg.Build, "race": s.Cfg.Race,
		"treeSizes": s.trees.Sizes, "treeNames": s.trees.Names, "fightsPerBuild": s.Cfg.FightsPerBuild,
		"fullChoices": s.Cfg.FullChoices, "talentMax": s.talentMax(), "talentRow": s.talentRows(), "talentTree": s.talentTreeIdx(),
		"talentPrereq": s.talentPrereqs(), "featureAPL": s.featureAPL, "actionAPL": s.actionAPL,
	})
	return out
}
