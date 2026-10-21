package env

import (
	"iter"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/wowsims/classic/sim/core"
)

// One simulation per talent build, with the lookups the policy and the observation need.
type simEntry struct {
	sim    *core.Simulation
	char   *core.Character
	target *core.Unit
	kind   []int // per parsed priority-list entry
	agent  []int
	spells []*core.Spell // per agent spell, nil when this build lacks it
	auras  []*core.Aura  // per tracked aura, nil when absent
	// Automatic-cooldown rules for agent-owned cooldowns, highest priority first: the teacher uses them as the
	// automatic set would.
	cdRules []cdRule
	usedAt  int64
	talents string
}

// Env is one environment. Its episodes run inside a coroutine, so the simulation stops at every decision.
type Env struct {
	spec *Spec
	id   int
	rng  *rand.Rand

	points     []int
	spent      int
	fightsLeft int
	talents    string
	newBuild   bool
	cache      map[string]*simEntry
	cur        *simEntry
	clock      int64

	Obs         []float32
	Mask        []uint8
	Teacher     int32
	Action      int32
	Reward      float64
	Done        bool
	Info        [4]float32 // at an episode end: DPS, decisions, new build (1/0), fight seconds
	LastTalents string

	readyPos      []int
	lastDamage    float64
	decisions     int
	sameTimeCount int
	lastDecisionT time.Duration
	episodes      int64
	SimsBuilt     int
	SimSeconds    float64
	forcedSeed    int64 // seed of the next fight when restarted (0: automatic)
	inFight       bool
	Reseed        int64         // when nonzero, reseed the fight's random streams right after the next action (then cleared)
	NextTalents   string        // when set, the talents of every following episode (not when talents are learned)
	waitUntil     time.Duration // a longer wait in progress
	waitSig       uint64        // state signature at the last plain wait
	waitAskedAt   time.Duration
	waiting       bool

	yield func(struct{}) bool
	next  func() (struct{}, bool)
	stop  func()
}

type stopSignal struct{}

type cdRule struct {
	agent int
	mcd   core.MajorCooldown
}

func newEnv(spec *Spec, id int) *Env {
	e := &Env{spec: spec, id: id, rng: rand.New(rand.NewSource(spec.Cfg.Seed*7919 + int64(id))),
		points: make([]int, spec.T), cache: map[string]*simEntry{},
		Obs: make([]float32, spec.ObsDim), Mask: make([]uint8, spec.ActDim), readyPos: make([]int, spec.K)}
	e.next, e.stop = iter.Pull(e.run)
	return e
}

// Advance resumes the episode with Action and runs to the next decision. The first call needs no action.
func (e *Env) Advance() {
	e.Reward, e.Done = 0, false
	e.Info = [4]float32{}
	if _, ok := e.next(); !ok {
		panic("environment coroutine ended")
	}
}

func (e *Env) Close() { e.stop() }

// Restart abandons the current episode and starts a new one whose fight uses seed (talents are chosen again when
// learned). A replay that feeds the same actions reaches the same states. Call it from the owning goroutine.
func (e *Env) Restart(seed int64) {
	e.stop()
	if e.inFight && e.cur != nil {
		// A fight abandoned at a decision leaves auras up; the engine only resets clean simulations.
		delete(e.cache, e.cur.talents)
		e.cur = nil
		e.inFight = false
	}
	e.forcedSeed = seed
	e.fightsLeft = 0
	e.Reward, e.Done = 0, false
	e.next, e.stop = iter.Pull(e.run)
}

func (e *Env) run(yield func(struct{}) bool) {
	e.yield = yield
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(stopSignal); !ok {
				panic(r)
			}
		}
	}()
	for {
		if e.spec.Cfg.LearnTalents && e.fightsLeft == 0 {
			e.chooseBuild()
			e.fightsLeft = e.spec.Cfg.FightsPerBuild
			e.newBuild = true
		} else if !e.spec.Cfg.LearnTalents && (e.talents == "" || (e.NextTalents != "" && e.NextTalents != e.talents)) {
			e.talents = e.spec.Cfg.Talents
			if e.NextTalents != "" {
				e.talents = e.NextTalents
			}
			pts, err := e.spec.trees.decode(e.talents)
			if err != nil {
				panic(err)
			}
			copy(e.points, pts)
			e.newBuild = true
		}
		e.fightsLeft--
		e.fight()
		e.newBuild = false
	}
}

func (e *Env) suspend() {
	if !e.yield(struct{}{}) {
		panic(stopSignal{})
	}
}

// chooseBuild spends the 51 points one decision at a time.
func (e *Env) chooseBuild() {
	for i := range e.points {
		e.points[i] = 0
	}
	t := e.spec.trees
	for e.spent = 0; e.spent < 51; e.spent++ {
		for i := range e.Mask {
			e.Mask[i] = 0
		}
		any := false
		for i := range t.Nodes {
			if t.canAdd(e.points, i) {
				e.Mask[i] = 1
				any = true
			}
		}
		if !any {
			break
		}
		teacher := t.teacherPoint(e.points, e.spec.basePoints)
		if teacher < 0 {
			for i := range t.Nodes {
				if e.Mask[i] == 1 {
					teacher = i
					break
				}
			}
		}
		e.Teacher = int32(teacher)
		e.writeCommon(1)
		e.clearFight()
		e.suspend()
		a := int(e.Action)
		if a < 0 || a >= e.spec.T || e.Mask[a] == 0 {
			a = teacher
		}
		e.points[a]++
	}
	e.talents = t.encode(e.points)
}

func (e *Env) entry() *simEntry {
	if ent, ok := e.cache[e.talents]; ok {
		e.clock++
		ent.usedAt = e.clock
		return ent
	}
	if len(e.cache) >= e.spec.Cfg.CachedSims {
		var oldest *simEntry
		for _, ent := range e.cache {
			if oldest == nil || ent.usedAt < oldest.usedAt {
				oldest = ent
			}
		}
		delete(e.cache, oldest.talents)
	}
	t0 := time.Now()
	sim, char := e.spec.newSim(e.talents, e.spec.candidate, e.spec.Cfg.Seed)
	e.SimsBuilt++
	defer func() { e.SimSeconds += time.Since(t0).Seconds() }()
	ent := &simEntry{sim: sim, char: char, target: sim.Encounter.TargetUnits[0], talents: e.talents}
	for _, cfgIdx := range char.Rotation.PriorityConfigIndexes() {
		ent.kind = append(ent.kind, e.spec.configKind[cfgIdx])
		ent.agent = append(ent.agent, e.spec.configAgent[cfgIdx])
	}
	ent.spells = make([]*core.Spell, len(e.spec.agentSpells))
	for k, id := range e.spec.agentSpells {
		ent.spells[k] = char.GetSpell(id)
	}
	ent.auras = make([]*core.Aura, len(e.spec.auras))
	for j, a := range e.spec.auras {
		unit := &char.Unit
		if a.Target {
			unit = ent.target
		}
		ent.auras[j] = unit.GetAura(a.Label)
	}
	if e.spec.Cfg.FullChoices {
		for _, mcd := range char.GetRemovedMajorCooldowns() {
			if mcd.Type.Matches(core.CooldownTypeSurvival) {
				continue
			}
			for k, id := range e.spec.agentSpells {
				if mcd.Spell.SameAction(id) {
					ent.cdRules = append(ent.cdRules, cdRule{agent: k, mcd: mcd})
				}
			}
		}
		sort.SliceStable(ent.cdRules, func(i, j int) bool { return ent.cdRules[i].mcd.Priority > ent.cdRules[j].mcd.Priority })
	}
	char.Rotation.SetPolicy(func(sim *core.Simulation, ready []int) int { return e.policy(sim, ready) })
	e.clock++
	ent.usedAt = e.clock
	e.cache[e.talents] = ent
	return ent
}

// damage is the damage the player and its pets have dealt so far this fight.
func (e *Env) damage() float64 {
	d := e.cur.char.LiveDamage()
	for _, pet := range e.cur.char.Pets {
		d += pet.LiveDamage()
	}
	return d
}

// fight plays one benchmark fight with the policy choosing every action.
func (e *Env) fight() {
	e.cur = e.entry()
	e.lastDamage, e.decisions, e.sameTimeCount = 0, 0, 0
	e.waitUntil, e.waiting = 0, false
	e.lastDecisionT = -1
	seed := e.spec.Cfg.Seed + int64(e.id)*1_000_003 + e.episodes*7_919
	if e.forcedSeed != 0 {
		seed, e.forcedSeed = e.forcedSeed, 0
	}
	e.episodes++
	e.inFight = true
	e.cur.sim.RunIteration(seed)
	e.inFight = false
	// The iteration's own total, pets included, as the benchmark reports it.
	total := e.cur.char.Metrics.IterationDamage()
	e.Reward += (total - e.lastDamage) * e.spec.Cfg.RewardScale
	e.Done = true
	seconds := e.cur.sim.Duration.Seconds()
	nb := float32(0)
	if e.newBuild {
		nb = 1
	}
	e.Info = [4]float32{float32(total / seconds), float32(e.decisions), nb, float32(seconds)}
	e.LastTalents = e.talents
	// The next decision (the next episode's first) carries this reward and Done.
}

func (e *Env) policy(sim *core.Simulation, ready []int) int {
	ent := e.cur
	for k, idx := range ready {
		if ent.kind[idx] == kindAuto {
			return k
		}
	}
	if sim.CurrentTime < e.waitUntil {
		return -1
	}
	spec := e.spec
	wait := spec.T + spec.K - 1
	for i := range e.Mask {
		e.Mask[i] = 0
	}
	e.Mask[wait] = 1
	for w := range spec.waits {
		e.Mask[wait-len(spec.waits)+w] = 1
	}
	teacher, teacherSet := wait, false
	anyAction := false
	for k, idx := range ready {
		switch ent.kind[idx] {
		case kindTeacher:
			if !teacherSet {
				teacherSet = true
				if a := ent.agent[idx]; a >= 0 {
					teacher = spec.T + a
				}
			}
		case kindAgent:
			a := ent.agent[idx]
			e.Mask[spec.T+a] = 1
			e.readyPos[a] = k
			anyAction = true
		}
	}
	if !anyAction && !spec.Cfg.KeepWaitOnly {
		// Waiting is the only choice: no decision to make.
		return -1
	}
	sig := e.signature(sim)
	if e.waiting && spec.Cfg.PollSeconds > 0 && sig == e.waitSig && sim.CurrentTime-e.waitAskedAt < core.DurationFromSeconds(spec.Cfg.PollSeconds) {
		return -1
	}
	e.waiting = false
	// The automatic cooldowns the agent owns come first for the teacher, under their own activation rules.
	for _, r := range ent.cdRules {
		if e.Mask[spec.T+r.agent] == 1 && (r.mcd.ShouldActivate == nil || r.mcd.ShouldActivate(sim, ent.char)) {
			teacher = spec.T + r.agent
			break
		}
	}
	if e.Mask[teacher] == 0 {
		teacher = wait
	}
	// Guard the rotation's loop limit: many choices at one instant end in a wait.
	if sim.CurrentTime == e.lastDecisionT {
		e.sameTimeCount++
		if e.sameTimeCount > 40 {
			return -1
		}
	} else {
		e.sameTimeCount = 0
		e.lastDecisionT = sim.CurrentTime
	}
	e.Teacher = int32(teacher)
	d := e.damage()
	e.Reward += (d - e.lastDamage) * spec.Cfg.RewardScale
	e.lastDamage = d
	e.decisions++
	e.writeCommon(0)
	e.writeFight(sim)
	e.suspend()
	if e.Reseed != 0 {
		// Copies of one state that try different actions then share every later random roll.
		sim.Reseed(e.Reseed)
		e.Reseed = 0
	}
	a := int(e.Action) - spec.T
	if a < 0 || a >= spec.K || e.Mask[spec.T+a] == 0 || a == spec.K-1 {
		e.waiting, e.waitSig, e.waitAskedAt = true, sig, sim.CurrentTime
		return -1
	}
	if w := a - (spec.K - 1 - len(spec.waits)); w >= 0 {
		e.waitUntil = sim.CurrentTime + core.DurationFromSeconds(spec.waits[w])
		return -1
	}
	return e.readyPos[a]
}

// signature hashes what the agent sees change in steps: which actions are usable, whole resource points and which
// tracked auras are up.
func (e *Env) signature(sim *core.Simulation) uint64 {
	h := uint64(14695981039346656037)
	mix := func(v uint64) {
		h ^= v
		h *= 1099511628211
	}
	for i := e.spec.T; i < len(e.Mask); i++ {
		mix(uint64(e.Mask[i]))
	}
	u := &e.cur.char.Unit
	if u.HasRageBar() {
		mix(uint64(u.CurrentRage()))
	}
	if u.HasEnergyBar() {
		mix(uint64(u.CurrentEnergy()))
		mix(uint64(u.ComboPoints()))
	}
	if u.HasManaBar() {
		mix(uint64(u.CurrentManaPercent() * 100))
	}
	for j, aura := range e.cur.auras {
		if aura != nil && aura.IsActive() {
			mix(uint64(j + 1))
		}
	}
	return h
}

func (e *Env) writeCommon(phase float32) {
	o := e.Obs
	o[0] = phase
	t := e.spec.trees
	for i, n := range t.Nodes {
		o[1+i] = float32(e.points[i]) / float32(n.Max)
	}
	off := 1 + len(t.Nodes)
	total := 0
	for ti := range t.Sizes {
		s := t.treeSpent(e.points, ti)
		total += s
		o[off+ti] = float32(s) / 51
	}
	o[off+len(t.Sizes)] = float32(total) / 51
}

func (e *Env) clearFight() {
	for i := e.spec.fightOffset; i < len(e.Obs); i++ {
		e.Obs[i] = 0
	}
}

func clip(v, hi float64) float32 {
	if v > hi {
		v = hi
	}
	if v < 0 || math.IsNaN(v) {
		v = 0
	}
	return float32(v)
}

func (e *Env) writeFight(sim *core.Simulation) {
	o := e.Obs[e.spec.fightOffset:]
	ent := e.cur
	u := &ent.char.Unit
	dur := sim.Duration.Seconds()
	now := sim.CurrentTime.Seconds()
	rem := dur - now
	o[0], o[1], o[2] = clip(now/dur, 1), clip(rem/dur, 1), clip(rem/300, 2)
	b := func(v bool) float32 {
		if v {
			return 1
		}
		return 0
	}
	o[3], o[4], o[5] = b(sim.IsExecutePhase20()), b(sim.IsExecutePhase25()), b(sim.IsExecutePhase35())
	o[6] = clip(u.GCD.TimeToReady(sim).Seconds(), 2)
	o[7], o[8], o[9], o[10], o[11] = 0, 0, 0, 0, 0
	if u.HasManaBar() {
		o[7] = clip(u.CurrentManaPercent(), 1)
		o[8] = clip(u.CurrentMana()/10000, 2)
	}
	if u.HasRageBar() {
		o[9] = clip(u.CurrentRage()/100, 1)
	}
	if u.HasEnergyBar() {
		o[10] = clip(u.CurrentEnergy()/100, 1.2)
		o[11] = clip(float64(u.ComboPoints())/5, 1)
	}
	o[12] = clip(u.DistanceFromTarget/30, 2)
	o[13], o[14], o[15] = 0, 0, 0
	if aa := &u.AutoAttacks; aa.MH() != nil || aa.Ranged() != nil {
		if mh := aa.MH(); mh != nil && mh.SwingSpeed > 0 {
			o[13] = clip((aa.MainhandSwingAt()-sim.CurrentTime).Seconds()/mh.SwingSpeed, 2)
		}
		if oh := aa.OH(); oh != nil && oh.SwingSpeed > 0 {
			o[14] = clip((aa.OffhandSwingAt()-sim.CurrentTime).Seconds()/oh.SwingSpeed, 2)
		}
		if r := aa.Ranged(); r != nil && r.SwingSpeed > 0 {
			o[15] = clip((aa.NextRangedAttackAt()-sim.CurrentTime).Seconds()/r.SwingSpeed, 2)
		}
	}
	if now > 1 {
		o[16] = clip(e.lastDamage/now/2000, 2)
	} else {
		o[16] = 0
	}
	p := 17
	for k, spell := range ent.spells {
		if spell == nil {
			o[p], o[p+1], o[p+2], o[p+3], o[p+4] = 0, 0, 0, 0, 0
		} else {
			o[p] = 1
			o[p+1] = float32(e.Mask[e.spec.T+k])
			o[p+2] = clip(spell.TimeToReady(sim).Seconds()/30, 1)
			o[p+3] = clip(spell.CastTime().Seconds()/3.5, 2)
			o[p+4] = 0
			if spell.Cost != nil && u.HasManaBar() && u.CurrentMana() > 0 {
				o[p+4] = clip(spell.Cost.GetCurrentCost()/u.CurrentMana(), 2)
			}
		}
		p += 5
	}
	for k := range e.spec.macros {
		o[p] = float32(e.Mask[e.spec.T+len(e.spec.agentSpells)+k])
		p++
	}
	for _, aura := range ent.auras {
		if aura == nil || !aura.IsActive() {
			o[p], o[p+1], o[p+2] = 0, 0, 0
		} else {
			o[p] = 1
			if aura.Duration > 0 && aura.Duration != core.NeverExpires {
				o[p+1] = clip(aura.RemainingDuration(sim).Seconds()/aura.Duration.Seconds(), 1)
			} else {
				o[p+1] = 1
			}
			if aura.MaxStacks > 0 {
				o[p+2] = clip(float64(aura.GetStacks())/float64(aura.MaxStacks), 1)
			}
		}
		p += 3
	}
}
