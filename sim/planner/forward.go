package planner

import (
	"fmt"
	"math"
	"os"

	"github.com/wowsims/classic/sim/warrior/ranks"
)

// The forward model: a fast event simulation of one Warrior against the current target (and the other engaged
// enemies) over a short horizon. The top-level search runs it once per candidate first action and random sample;
// inside it, a base policy (immediate value minus a rage price) makes the later decisions.

const (
	actWait    = -1
	actNone    = -2
	actUnqueue = -3

	maxAbilities = 48
)

// Ctx is everything a decision's rollouts share: the model, the intent and the expected outcome tables at the
// decision's stats.
type Ctx struct {
	M             *Model
	P             ResolvedPolicy
	H             float64 // horizon, seconds
	AP            float64 // attack power without Battle Shout
	BSAP          float64 // Battle Shout attack power, if it was up at decision time
	Crit          float64 // percent, character sheet
	Targets       int     // enemies in melee range, including the target
	TLevel        int32
	NextTTK       float64 // life of the next enemy when the target dies
	TargetAttacks bool    // the target is attacking the player
	Front         bool    // attacking from the front (the target faces the player): parry and block apply
	TArmor        float64
	Elite         bool

	// Threat value per point: target and other enemies (hold), or a cost per point (group DPS).
	WThreatTarget, WThreatOther float64

	// Incoming melee from mobs attacking the player.
	Attackers   float64
	MobSpeed    float64
	MobHit      float64 // raw damage per landed mob hit, before Demoralizing Shout and blocks
	MobAPEquiv  float64 // damage share removed per point of Demoralizing Shout (engine's enemy AP coefficient)
	TTK         [3]float64
	TTKProb     [3]float64
	ExecuteFrac float64 // share of the target's life left when it enters execute range (health 20%)
	HealthPct   float64

	// Outcome tables.
	mhWhite, ohWhite, special, noAvoid, magic outcome
	glanceMult                                float64
	ThreatWeight                              float64 // weight applied to threat (1 when a group-DPS ceiling is active)
	allowed                                   [maxAbilities]bool
	armorDR                                   [6]float64 // damage multiplier after armor by Sunder stacks

	Mu       float64 // rage price for the base policy
	trace    bool
	Pr       Params
	sinkCost float64 // rage cost of the rage sink (Heroic Strike or Cleave)

	abil   []*Ability
	byName map[ranks.Ability]int
}

type outcome struct {
	miss, dodge, parry, glance, crit float64
	// expected damage multiplier of a roll (glance and crit included) and landed probability
	mult, land float64
}

// State is a rollout's mutable state. Times are seconds from the decision.
type State struct {
	t         float64
	rage      float64
	gcdAt     float64
	cdAt      [maxAbilities]float64
	stance    int
	queued    int     // ability index of a queued Heroic Strike/Cleave, or -1
	requeueAt float64 // after a cancel, the console presses the swing again no sooner than this
	queuedTop bool    // the queued swing is the decision's own (queued now or before); later ones are the policy's
	mhAt      float64
	ohAt      float64
	flurry    int
	mobAt     float64
	amAt      float64
	brEnd     float64
	brTick    float64
	bsEnd     float64
	enrEnd    float64
	dwEnd     float64 // Death Wish (manual; only if already up)
	reckEnd   float64
	sbEnd     float64
	sbLeft    int
	revEnd    float64
	opEnd     float64
	sunder    int
	sunEnd    float64
	rendEnd   float64
	rendNext  float64
	tcEnd     float64
	demoEnd   float64
	deathAt   float64
	execAt    float64
	ttkRef    float64 // the policy's own estimate of the target's death time
	targets   int     // enemies in melee, including the target
	attackers float64
	wt        float64 // threat value per point on the current target
	homeAt    float64 // when a borrowed stance returns home
	home      int

	dmg, threat, mitig float64
	rng                uint64
	draws              [rsCount + maxAbilities]uint32
}

func (c *Ctx) idx(a ranks.Ability) int {
	if i, ok := c.byName[a]; ok {
		return i
	}
	return -1
}

// Random streams. Each event kind draws from its own stream by its own count, so two candidates' rollouts see
// the same rolls for the same swings and mob attacks even after their actions differ (common random numbers).
const (
	rsMH = iota
	rsOH
	rsMob
	rsStrike
	rsCount
)

// draw returns stream's next uniform [0,1).
func (s *State) draw(stream int) float64 {
	n := s.draws[stream]
	s.draws[stream]++
	x := s.rng + uint64(stream+1)*0xD1B54A32D192ED03 + uint64(n)*0x9E3779B97F4A7C15
	x ^= x >> 30
	x *= 0xBF58476D1CE4E5B9
	x ^= x >> 27
	x *= 0x94D049BB133111EB
	x ^= x >> 31
	return float64(x>>11) / float64(1<<53)
}

// rand returns a uniform [0,1) from the state's xorshift generator (sample setup only).
func (s *State) rand() float64 {
	s.rng ^= s.rng << 13
	s.rng ^= s.rng >> 7
	s.rng ^= s.rng << 17
	return float64(s.rng>>11) / float64(1<<53)
}

// ---------------------------------------------------------------- outcome tables

func (c *Ctx) build() {
	m := c.M
	c.abil = m.Abilities
	c.byName = map[ranks.Ability]int{}
	for i, ab := range c.abil {
		c.byName[ab.Name] = i
	}
	pl := float64(m.Level)
	tl := float64(c.TLevel)
	skill := pl * 5
	def := tl * 5
	diff := def - skill
	missBase := 0.05 + diff*0.001
	if diff > 10 {
		missBase = 0.05 + diff*0.002
	}
	dodge := max(0, 0.05+diff*0.001-m.ExpertisePct/100)
	parry := max(0, 0.05+diff*0.001-m.ExpertisePct/100)
	if diff > 10 {
		parry = max(0, 0.05+diff*0.006-m.ExpertisePct/100)
	}
	glance := max(0, 0.1+diff*0.02)
	glanceMult := (max(min(1.3-0.05*diff, 0.91), 0.01) + max(min(1.2-0.03*diff, 0.99), 0.2)) / 2
	c.glanceMult = glanceMult
	suppress := diff * 0.002
	if diff <= 0 {
		suppress = diff * 0.0004
	}
	crit := max(0, c.Crit/100-suppress)
	hit := m.HitPct / 100
	// Attacking from the front adds the enemy's parry and block (enemies block for no damage, but a blocked
	// weapon attack cannot crit). White swings use one roll; specials roll crit separately on what lands.
	block := 0.0
	if !c.Front {
		parry = 0
	} else {
		block = 0.05
	}
	white := func(miss, crit, critMult float64) outcome {
		miss = max(0, miss)
		land := max(0, 1-miss-dodge-parry)
		crit = min(crit, max(0, land-glance-block))
		o := outcome{miss: miss, dodge: dodge, parry: parry, glance: glance, crit: crit, land: land}
		o.mult = (land - glance - crit) + glance*glanceMult + crit*critMult
		return o
	}
	special := func(miss, dodge, parry, block, crit, critMult float64) outcome {
		miss = max(0, miss)
		land := max(0, 1-miss-dodge-parry)
		crit = max(0, land-block) * min(crit, 1)
		o := outcome{miss: miss, dodge: dodge, parry: parry, crit: crit, land: land}
		o.mult = (land - crit) + crit*critMult
		return o
	}
	dwPenalty := 0.0
	if m.DW {
		dwPenalty = 0.19
	}
	c.mhWhite = white(missBase+dwPenalty-hit, crit+m.MH.CritBonus/100, m.WhiteCritMult)
	c.ohWhite = white(missBase+dwPenalty-hit-m.OH.HitBonus/100, crit, m.WhiteCritMult)
	c.special = special(missBase-hit, dodge, parry, block, crit, m.SpecialCritMult)
	c.noAvoid = special(missBase-hit, 0, 0, 0, crit, m.SpecialCritMult)
	spellMiss := []float64{0.04, 0.05, 0.06, 0.17}[int(min(max(tl-pl, 0), 3))]
	c.magic = special(spellMiss, 0, 0, 0, crit, 1.5*max(m.CritScale, 1))
	k := 400 + 85*pl
	sunder := 0.0
	if ab := m.Get(ranks.SunderArmor); ab != nil {
		sunder = ab.ArmorPerStack
	}
	for s := 0; s <= 5; s++ {
		armor := max(0, c.TArmor*(1-m.ArmorIgnore)-sunder*float64(s))
		c.armorDR[s] = 1 - armor/(armor+k)
	}
}

// ---------------------------------------------------------------- damage helpers

func (c *Ctx) ap(s *State) float64 {
	ap := c.AP
	if s.t < s.bsEnd {
		ap += c.BSAP
	}
	return ap
}

// physMult is the multiplier on physical damage dealt at the state's time, armor included.
func (c *Ctx) physMult(s *State) float64 {
	m := c.M.PhysMult * c.armorDR[s.sunder]
	if s.stance == StanceDefensive {
		m *= 0.9
	}
	if s.t < s.enrEnd {
		m *= 1 + c.M.EnragePct
	}
	if s.t < s.dwEnd {
		m *= 1.2
	}
	return m
}

func (c *Ctx) stanceThreat(s *State) float64 {
	if s.stance == StanceDefensive {
		return c.M.DefStanceThreat
	}
	return c.M.OtherStanceMult
}

func swingDamage(w Weapon, ap float64, normalized bool) float64 {
	speed := w.Speed
	if normalized {
		speed = w.NormSpeed
	}
	return w.Avg() + ap*speed/14
}

func (c *Ctx) table(ab *Ability) *outcome {
	switch ab.Table {
	case TableNoAvoid:
		return &c.noAvoid
	case TableMagic:
		return &c.magic
	default:
		return &c.special
	}
}

// abilityRaw is an ability's damage per target before the roll and multipliers other than its own.
func (c *Ctx) abilityRaw(ab *Ability, ap float64) float64 {
	d := ab.Flat + ab.APCoef*ap
	if ab.WeaponMult > 0 {
		d += ab.WeaponMult * swingDamage(c.M.MH, ap, ab.Normalized)
	}
	return d * ab.DamageMult
}

// credit adds damage and threat dealt at the state's time to enemy e (0 = target).
func (c *Ctx) credit(s *State, damage, threat float64, other bool) {
	if !other {
		if s.t >= s.deathAt {
			return
		}
		s.dmg += damage
		s.threat += threat * s.wt
		return
	}
	s.dmg += damage
	s.threat += threat * c.WThreatOther
}

// ---------------------------------------------------------------- events

func (c *Ctx) whiteSwing(s *State, oh bool) {
	m := c.M
	w, o := m.MH, &c.mhWhite
	stream := rsMH
	if oh {
		w, o, stream = m.OH, &c.ohWhite, rsOH
	}
	r, r2 := s.draw(stream), s.draw(stream)
	ap := c.ap(s)
	switch {
	case r < o.miss:
	case r < o.miss+o.dodge:
		s.opEnd = s.t + 5
	case r < o.miss+o.dodge+o.parry:
	default:
		crit := r >= 1-o.crit
		glance := !crit && r < o.miss+o.dodge+o.parry+o.glance
		d := swingDamage(w, ap, false) * w.DamageMult * c.physMult(s)
		rage := w.RagePerLanded
		if crit {
			d *= m.WhiteCritMult
			rage *= 2
			if m.FlurryHaste > 0 {
				s.flurry = 3
			}
			if m.DeepWoundsPct > 0 {
				c.deepWounds(s)
			}
		} else if glance {
			d *= c.glanceMult
		}
		s.rage = min(100, s.rage+rage)
		if m.UWChance > 0 && r2 < m.UWChance {
			s.rage = min(100, s.rage+1)
		}
		c.credit(s, d, d*c.stanceThreat(s), false)
	}
}

func (c *Ctx) deepWounds(s *State) {
	// The bleed pays the main-hand average weapon damage share over 12 s; only the part before death counts.
	alive := min(12, max(0, s.deathAt-s.t)) / 12
	d := c.M.MH.Avg() * c.M.DeepWoundsPct * alive * c.armorFreeMult(s)
	c.credit(s, d, d*c.stanceThreat(s), false)
}

// armorFreeMult: bleeds ignore armor.
func (c *Ctx) armorFreeMult(s *State) float64 {
	return c.physMult(s) / c.armorDR[s.sunder]
}

func (c *Ctx) queuedSwing(s *State) {
	ab := c.abil[s.queued]
	if c.trace {
		fmt.Fprintf(os.Stderr, "      %.2f queued swing %s rage %.1f\n", s.t, ab.Name, s.rage)
	}
	s.queued = -1
	top := s.queuedTop
	s.queuedTop = false
	// Short of rage it is a white swing. With cancelling allowed, the swing is also cancelled (a white swing) when
	// it would leave too little rage for an ability coming off its cooldown before the next swing: later decisions
	// can cancel it, except the decision's own swing when it lands too soon for another one.
	if s.rage < ab.Cost || c.Pr.Unqueue && (!top || s.t >= c.Pr.CancelLead) && s.rage-ab.Cost < c.reserve(s) {
		c.whiteSwing(s, false)
		return
	}
	s.rage -= ab.Cost
	o := &c.special
	ap := c.ap(s)
	targets := 1
	if ab.Targets > 1 {
		targets = min(ab.Targets, s.targets)
	}
	for k := 0; k < targets; k++ {
		var r float64
		if k == 0 {
			r = s.draw(rsMH)
			s.draw(rsMH)
		} else {
			r = s.draw(rsStrike + ab.Index)
		}
		if r < o.miss+o.dodge+o.parry {
			if r >= o.miss && r < o.miss+o.dodge && k == 0 {
				s.opEnd = s.t + 5
			}
			if k == 0 {
				s.rage = min(100, s.rage+ab.Cost*ab.Refund)
			}
			continue
		}
		d := c.abilityRaw(ab, ap) * c.physMult(s)
		if r >= 1-o.crit {
			d *= ab.CritMult
			if c.M.FlurryHaste > 0 {
				s.flurry = 3
			}
			if c.M.DeepWoundsPct > 0 && k == 0 {
				c.deepWounds(s)
			}
		}
		th := (d*ab.ThreatMult + ab.FlatThreat) * c.stanceThreat(s)
		if c.trace {
			fmt.Fprintf(os.Stderr, "        hit %d dmg %.1f r %.3f\n", k, d, r)
		}
		c.credit(s, d, th, k > 0)
	}
}

func (c *Ctx) mobAttack(s *State) {
	m := c.M
	interval := c.MobSpeed / max(s.attackers, 1e-9)
	if s.t < s.tcEnd {
		interval *= 1.2
	}
	s.mobAt = s.t + interval
	r, r2 := s.draw(rsMob), s.draw(rsMob)
	pMiss, pDodge, pParry := m.MissPct/100, m.DodgePct/100, m.ParryPct/100
	pBlock := 0.0
	if m.Shield {
		pBlock = m.BlockPct / 100
		if s.t < s.sbEnd && s.sbLeft > 0 {
			pBlock += 0.75
		}
	}
	raw := c.MobHit
	slowSaved := 0.0
	if s.t < s.tcEnd {
		slowSaved = raw * 0.2
	}
	demoSaved := 0.0
	if s.t < s.demoEnd {
		if ab := c.M.Get(ranks.DemoralizingShout); ab != nil {
			demoSaved = raw * ab.APReduction * c.MobAPEquiv
		}
	}
	switch {
	case r < pMiss:
		return
	case r < pMiss+pDodge+pParry:
		s.revEnd = s.t + 5
		return
	}
	taken := raw - demoSaved
	saved := demoSaved + slowSaved
	if r < pMiss+pDodge+pParry+pBlock {
		blocked := min(taken, m.BlockValue)
		taken -= blocked
		saved += blocked
		s.revEnd = s.t + 5
		if s.t < s.sbEnd && s.sbLeft > 0 {
			s.sbLeft--
		}
		if m.ShieldSpecRage > 0 {
			s.rage = min(100, s.rage+m.ShieldSpecRage)
		}
	}
	if s.stance == StanceDefensive {
		taken *= 0.9
	}
	s.mitig += saved
	s.rage = min(100, s.rage+taken*m.IncomingRagePerD)
	if m.EnrageChance > 0 && taken > 0 && r2 < m.EnrageChance {
		s.enrEnd = s.t + 12
	}
}

// ---------------------------------------------------------------- casting

// usable reports whether ability i can be used now (ignoring the rage price).
func (c *Ctx) usable(s *State, i int) bool {
	ab := c.abil[i]
	if !ab.Modelled || !c.allowed[i] || s.t < s.cdAt[i] || (ab.GCD && s.t < s.gcdAt) {
		return false
	}
	if ab.Stances != 0 && ab.Stances&(1<<s.stance) == 0 {
		return false
	}
	if ab.NeedsShield && !c.M.Shield {
		return false
	}
	switch ab.Kind {
	case KindQueued:
		// One swing is queued at a time (cancel it first to queue the other).
		if s.queued >= 0 || s.rage < ab.Cost {
			return false
		}
	case KindExecute:
		if s.t < s.execAt || s.rage < ab.Cost {
			return false
		}
	case KindStance:
		return ab.ToStance != s.stance && s.t >= s.cdAt[i]
	default:
		if s.rage < ab.Cost {
			return false
		}
	}
	switch ab.NeedsWindow {
	case windowRevenge:
		return s.t < s.revEnd
	case windowOverpower:
		return s.t < s.opEnd
	}
	// Refreshing early only throws away what is left: debuffs and buffs are renewed when they are about to lapse.
	switch {
	case ab.Kind == KindDot && s.rendEnd-s.t > 3:
		return false
	case ab.Name == ranks.SunderArmor && s.sunder >= 5 && s.sunEnd-s.t > 6:
		return false
	case ab.Name == ranks.ThunderClap && s.targets <= 1 && s.tcEnd-s.t > 2:
		return false
	case ab.Name == ranks.DemoralizingShout && s.demoEnd-s.t > 4:
		return false
	case ab.Kind == KindBuff && s.bsEnd-s.t > 3:
		return false
	}
	if s.t >= s.deathAt && ab.Kind != KindBuff && ab.Kind != KindRage && ab.Kind != KindBlock && ab.Kind != KindStance {
		return false
	}
	return true
}

// cast applies ability i at the state's time.
func (c *Ctx) cast(s *State, i int) {
	ab := c.abil[i]
	if c.trace {
		fmt.Fprintf(os.Stderr, "      %.2f cast %s rage %.1f mh %.2f q %d\n", s.t, ab.Name, s.rage, s.mhAt, s.queued)
	}
	m := c.M
	if ab.GCD {
		s.gcdAt = s.t + 1.5 + ab.Cast
	}
	s.cdAt[i] = s.t + ab.CD
	ap := c.ap(s)
	switch ab.Kind {
	case KindQueued:
		s.queued, s.queuedTop = i, false
		return
	case KindRage:
		s.rage = min(100, s.rage+ab.RageGain)
		s.brEnd, s.brTick = s.t+ab.Duration, s.t+1
		return
	case KindBlock:
		s.rage -= ab.Cost
		s.sbEnd, s.sbLeft = s.t+ab.Duration, ab.BlockCharges
		return
	case KindStance:
		if s.homeAt == 0 && s.home == 0 {
			s.home = s.stance
		}
		s.stance = ab.ToStance
		s.rage = min(s.rage, m.TacticalMastery)
		return
	case KindBuff:
		s.rage -= ab.Cost
		s.bsEnd = s.t + ab.Duration
		return
	}
	s.rage -= ab.Cost
	if ab.Cast > 0 && !ab.KeepSwings {
		// Untalented Slam stops auto-attacks during the cast and restarts both swing timers when it completes
		// (the off hand half a swing later).
		s.mhAt = s.t + ab.Cast + m.MH.Speed
		if m.DW {
			s.ohAt = s.t + ab.Cast + 1.5*m.OH.Speed
		}
	}
	switch ab.Kind {
	case KindDebuff:
		if ab.Name == ranks.DemoralizingShout {
			s.demoEnd = s.t + ab.Duration
			th := (ab.FlatThreat) * c.stanceThreat(s)
			c.credit(s, 0, th, false)
			if s.targets > 1 {
				c.credit(s, 0, th*float64(s.targets-1), true)
			}
			return
		}
		// Sunder Armor: miss, dodge or parry refund most of the cost.
		r := s.draw(rsStrike + ab.Index)
		o := &c.special
		if r < o.miss+o.dodge+o.parry {
			s.rage += ab.Cost * ab.Refund
			if r >= o.miss && r < o.miss+o.dodge {
				s.opEnd = s.t + 5
			}
			return
		}
		if s.t >= s.sunEnd {
			s.sunder = 0
		}
		s.sunder = min(5, s.sunder+1)
		s.sunEnd = s.t + ab.Duration
		c.credit(s, 0, (ab.FlatThreat+ab.ThreatAP*ap)*c.stanceThreat(s), false)
		return
	case KindDot:
		r := s.draw(rsStrike + ab.Index)
		o := &c.special
		if r < o.miss+o.dodge+o.parry {
			s.rage += ab.Cost * ab.Refund
			return
		}
		s.rendEnd, s.rendNext = s.t+ab.Duration, s.t+3
		return
	case KindExecute:
		extra := s.rage
		s.rage = 0
		c.strike(s, ab, ab.Flat+ab.PerRage*extra, 1, false)
		return
	case KindAoE:
		n := min(max(ab.Targets, 1), s.targets)
		base := c.abilityRaw(ab, ap)
		if ab.Name == ranks.ThunderClap {
			s.tcEnd = s.t + ab.Duration
		}
		for k := 0; k < n; k++ {
			c.strike(s, ab, base, 1, k > 0)
		}
		if ab.Name == ranks.Whirlwind && m.DW {
			// The off-hand Whirlwind strike (October 1, dual wield).
			ohBase := swingDamage(m.OH, ap, true) * m.OH.DamageMult
			for k := 0; k < n; k++ {
				c.strike(s, ab, ohBase, 1, k > 0)
			}
		}
		return
	}
	if ab.NeedsWindow == windowRevenge {
		s.revEnd = 0
	}
	if ab.NeedsWindow == windowOverpower {
		s.opEnd = 0
	}
	c.strike(s, ab, c.abilityRaw(ab, ap), 1, false)
}

// strike resolves one hit of an ability's base damage on one enemy.
func (c *Ctx) strike(s *State, ab *Ability, base float64, scale float64, other bool) {
	o := c.table(ab)
	r := s.draw(rsStrike + ab.Index)
	if r < o.miss+o.dodge+o.parry {
		if !other {
			s.rage += ab.Cost * ab.Refund
			if r >= o.miss && r < o.miss+o.dodge {
				s.opEnd = s.t + 5
			}
		}
		return
	}
	d := base * scale * c.physMult(s)
	if r >= 1-o.crit {
		d *= ab.CritMult
		if ab.Table != TableMagic {
			if c.M.FlurryHaste > 0 && !other {
				s.flurry = 3
			}
			if c.M.DeepWoundsPct > 0 && !other {
				c.deepWounds(s)
			}
		}
	}
	th := (d*ab.ThreatMult + ab.FlatThreat + ab.ThreatAP*c.ap(s)) * c.stanceThreat(s)
	c.credit(s, d, th, other)
}

// ---------------------------------------------------------------- the loop

// run advances the state to the horizon, letting the base policy act from the first decision on. If first >= 0 or
// first == actWait, that is the decision at t=0 (actWait holds until waitUntil).
func (c *Ctx) run(s *State, first int, waitUntil float64) {
	nextDecision, holdUntil := 0.0, 0.0
	if first == actWait {
		holdUntil = waitUntil
	} else if first >= 0 {
		c.cast(s, first)
		if c.abil[first].Kind == KindQueued {
			s.queuedTop = true
		}
	}
	actions := 0
	for {
		// Decide whenever the policy may act.
		if s.t >= nextDecision && s.t >= holdUntil && s.t < c.H {
			act := actNone
			if actions < 400 {
				act = c.basePolicy(s)
			}
			if act >= 0 {
				c.cast(s, act)
				actions++
				continue // another action may follow (off-GCD actions)
			}
			nextDecision = math.Inf(1)
		}
		// Next event.
		next := c.H
		kind := 0
		if s.mhAt < next {
			next, kind = s.mhAt, 1
		}
		if c.M.DW && s.ohAt < next {
			next, kind = s.ohAt, 2
		}
		if s.attackers > 0 && s.mobAt < next {
			next, kind = s.mobAt, 3
		}
		if s.brTick > 0 && s.brTick <= s.brEnd && s.brTick < next {
			next, kind = s.brTick, 4
		}
		if c.M.AngerManagement && s.amAt < next {
			next, kind = s.amAt, 5
		}
		if s.rendNext > 0 && s.rendNext <= s.rendEnd && s.rendNext < next {
			next, kind = s.rendNext, 6
		}
		if s.gcdAt > s.t && s.gcdAt < next {
			next, kind = s.gcdAt, 7
		}
		if s.homeAt > s.t && s.homeAt < next {
			next, kind = s.homeAt, 8
		}
		if holdUntil > s.t && holdUntil < next {
			next, kind = holdUntil, 9
		}
		if s.deathAt > s.t && s.deathAt < next {
			next, kind = s.deathAt, 10
		}
		if next >= c.H {
			s.t = c.H
			return
		}
		s.t = next
		switch kind {
		case 1:
			haste := 1.0
			if s.flurry > 0 {
				s.flurry--
				haste = 1 + c.M.FlurryHaste
			}
			if s.queued >= 0 {
				c.queuedSwing(s)
			} else {
				c.whiteSwing(s, false)
			}
			s.mhAt = s.t + c.M.MH.Speed/haste
		case 2:
			haste := 1.0
			if s.flurry > 0 {
				haste = 1 + c.M.FlurryHaste
			}
			c.whiteSwing(s, true)
			s.ohAt = s.t + c.M.OH.Speed/haste
		case 3:
			c.mobAttack(s)
		case 4:
			if ab := c.M.Get(ranks.Bloodrage); ab != nil {
				s.rage = min(100, s.rage+ab.RagePerSec)
			}
			s.brTick += 1
		case 5:
			s.rage = min(100, s.rage+1)
			s.amAt += 3
		case 6:
			if ab := c.M.Get(ranks.Rend); ab != nil {
				d := (ab.TickDamage + ab.TickAPCoef*c.ap(s)) * ab.DamageMult * c.armorFreeMult(s) * c.tickCrit()
				c.credit(s, d, d*c.stanceThreat(s), false)
			}
			s.rendNext += 3
		case 10:
			c.targetDied(s)
		case 8:
			if s.home != 0 {
				s.stance = s.home
				s.rage = min(s.rage, c.M.TacticalMastery)
				s.home = 0
			}
			s.homeAt = 0
		}
		nextDecision = s.t
	}
}

// targetDied moves the player to the next enemy when the target dies: its debuffs are gone, the next one has its
// own life and threat weight. With no enemy left, nothing more is credited.
func (c *Ctx) targetDied(s *State) {
	if c.TargetAttacks {
		s.attackers = max(0, s.attackers-1)
	}
	if s.targets <= 1 {
		return
	}
	s.targets--
	s.sunder, s.sunEnd, s.rendEnd, s.rendNext = 0, 0, 0, 0
	s.deathAt = s.t + c.NextTTK
	s.ttkRef = s.deathAt
	s.execAt = s.t + 0.8*c.NextTTK
	s.wt = c.WThreatOther
}
