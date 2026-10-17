package planner

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/warrior/ranks"
)

// Params are the planner's tunable settings (the simulator's sweeps tune them).
type Params struct {
	Horizon        float64 // seconds
	MinSamples     int
	MaxSamples     int
	HoldDecay      float64  // lead-% scale over which threat value fades above the secure lead
	HoldFloor      float64  // threat value per point (hold) when the lead is far above secure
	HoldLook       float64  // seconds the lead trend is projected over (capped by the time to die)
	CeilingK       float64  // group-DPS threat cost at the ceiling, per threat point
	CeilingStart   float64  // share of the ceiling where the group-DPS threat cost starts
	UnknownLead    float64  // threat value per point on an enemy whose lead is unknown (hold)
	LooseWeight    float64  // threat value per point (hold) on an enemy someone else has
	ContestWeight  float64  // threat value per point (hold) on an enemy the player is tanking without the lead
	ArmorScale     float64  // scale on the armor-reduction value of Sunder Armor beyond what the rollout sees
	SunderToDeath  bool     // value Sunder stacks up at the horizon as kept up (refreshed) until the target dies
	SwingLossScale float64  // scale on the value of off-hand swings a cast pushes back (untalented Slam)
	SwingLossMH    float64  // the same for the main hand
	NextBoost      float64  // extra threat value on other enemies as the target's death nears (hold)
	NextScale      float64  // seconds of the target's remaining life over which that boost fades
	RageKeep       float64  // share of the rage price leftover rage keeps at the horizon
	SinkFloor      float64  // the policy queues a rage sink only above this share of the rage price per rage
	WaitBias       float64  // percent of its value added to waiting
	CandFloor      float64  // candidates need immediate value net of rage above this share of their rage cost's price (-1 = off)
	CDValue        float64  // weight of the cooldown time still running at the horizon
	QueueGate      float64  // Heroic Strike/Cleave may be queued only this many seconds before the swing (<= 0: any time)
	Unqueue        bool     // the planner may cancel a queued swing (contract v1.4)
	CancelLead     float64  // seconds before the swing the console can still be relied on to cancel the swing queued now
	UnqueueBias    float64  // percent of its value taken off cancelling (a toggle must clearly win, not by rollout noise)
	GCDFirst       float64  // percent within which a global-cooldown ability goes before a queue or cancel
	Disable        []string `json:",omitempty"` // abilities the planner never uses (diagnostics)
	// FixedSamples, when set, replaces the time budget with a fixed sample count (reproducible benchmarks).
	FixedSamples int `json:",omitempty"`
}

var DefaultParams = Params{Horizon: 7, MinSamples: 6, MaxSamples: 60, HoldDecay: 40, HoldFloor: 1, HoldLook: 8, CeilingK: 3, CeilingStart: 0.7, UnknownLead: 0.7,
	LooseWeight: 3, ContestWeight: 2, ArmorScale: 0.5, SwingLossScale: 4, SwingLossMH: 4, NextBoost: 12, NextScale: 8, RageKeep: 0.9, SinkFloor: 0.5, WaitBias: 0, CandFloor: -0.25, CDValue: 1,
	QueueGate: 0, Unqueue: true, CancelLead: 0.3, UnqueueBias: 5, GCDFirst: 1, SunderToDeath: true}

// enemyAPCoefficient is the simulator's enemy attack power coefficient (core.EnemyAutoAttackAPCoefficient): the
// share of base damage one point of attack power adds, used to value Demoralizing Shout.
const enemyAPCoefficient = 1.0 / (14.0 * 177.0)

func f(p *float64, def float64) float64 {
	if p == nil || math.IsNaN(*p) {
		return def
	}
	return *p
}

// holdWeight is the value of a point of threat on an enemy while holding aggro, from its threat state. The lead is
// projected over the look-ahead with its trend; above the secure lead the value falls towards HoldFloor (threat
// banked now still guards against the partner's later threat), never to zero.
func (pr Params) holdWeight(t *ThreatState, targetsPlayer *bool, secure float64, lookahead float64) float64 {
	if targetsPlayer != nil && !*targetsPlayer {
		return pr.LooseWeight // loose: someone else has it
	}
	if t == nil || t.LeadPct == nil || t.Status == nil {
		return pr.UnknownLead
	}
	if *t.Status < 3 {
		if *t.Status < 2 {
			return pr.LooseWeight
		}
		return pr.ContestWeight
	}
	lead := min(*t.LeadPct, 400) + min(f(t.LeadTrendPctPer, 0), 0)*lookahead
	if lead <= secure {
		return 1
	}
	return pr.HoldFloor + (1-pr.HoldFloor)*math.Exp(-(lead-secure)/pr.HoldDecay)
}

// ceilingCost is the cost of a point of threat for a damage dealer whose threat is lead% of the aggro holder's.
func (pr Params) ceilingCost(t *ThreatState, ceiling, look float64) float64 {
	// The lead projected over the look-ahead with its rising trend: a pull comes from where the lead is heading.
	lead := *t.LeadPct + max(f(t.LeadTrendPctPer, 0), 0)*look
	start := pr.CeilingStart * ceiling
	if lead < start {
		return 0
	}
	x := (lead - start) / (ceiling - start)
	if x >= 1 {
		return pr.CeilingK * (1 + (x-1)*5)
	}
	return pr.CeilingK * x * x
}

// NewCtx builds a decision context and the initial rollout state from a decide message.
func NewCtx(m *Model, msg *DecideMsg, pr Params) (*Ctx, *State, []string) {
	var warn []string
	pol := ResolveIntent(msg.Intent)
	c := &Ctx{M: m, P: pol, H: pr.Horizon, ThreatWeight: pol.Weights.Threat, Pr: pr}
	tg := msg.Target
	c.TLevel = m.Level
	c.Front = tg.TargetsPlayer == nil || *tg.TargetsPlayer
	if tg.Level != nil {
		c.TLevel = int32(*tg.Level)
	}
	lvl := min(max(c.TLevel, 1), int32(len(expectedCreatures)-1))
	ec := expectedCreatures[lvl]
	c.TArmor = ec.armor
	c.Elite = tg.Classification != nil && (*tg.Classification == "elite" || *tg.Classification == "rareelite" || *tg.Classification == "worldboss")
	c.HealthPct = f(tg.HealthPct, 100)
	// Death time quantiles. Unknown: a normal mob's health over a typical solo damage rate.
	if tg.TTK != nil && tg.TTK.P50 != nil {
		p50 := *tg.TTK.P50
		c.TTK = [3]float64{f(tg.TTK.P10, p50*0.6), p50, f(tg.TTK.P90, p50*1.6)}
	} else {
		base := 15.0 * c.HealthPct / 100
		if c.Elite {
			base *= 3
		}
		c.TTK = [3]float64{base * 0.6, base, base * 1.7}
	}
	for i := range c.TTK {
		c.TTK[i] = max(c.TTK[i], 0.5)
	}
	// Enemies in melee for area attacks.
	c.Targets = 1
	if n := msg.EnemyCount.Melee; n != nil && *n >= 1 {
		c.Targets = int(*n)
	} else if n := msg.EnemyCount.Manual; n != nil && *n >= 1 {
		c.Targets = int(*n)
	}
	// Attackers on the player and their damage.
	attackers := 0.0
	if !tg.Exists {
		attackers = 0
	} else if tg.TargetsPlayer == nil || *tg.TargetsPlayer {
		attackers = 1
	}
	for _, e := range msg.Enemies {
		if e.TargetsPlayer != nil && *e.TargetsPlayer {
			attackers++
		}
	}
	c.Attackers, c.MobSpeed = attackers, 2.0
	c.TargetAttacks = tg.Exists && (tg.TargetsPlayer == nil || *tg.TargetsPlayer)
	// The next enemy's life when the target dies: the others' median, or the target's own.
	c.NextTTK = 0
	nn := 0
	for _, e := range msg.Enemies {
		if e.TTK != nil && e.TTK.P50 != nil {
			c.NextTTK += *e.TTK.P50
			nn++
		}
	}
	if nn > 0 {
		c.NextTTK /= float64(nn)
	} else {
		c.NextTTK = c.TTK[1] / max(c.HealthPct/100, 0.2)
	}
	landed := 1 - m.MissPct/100 - m.DodgePct/100 - m.ParryPct/100
	if in := msg.Player.IncomingPctPerS; in != nil && in.Mean != nil && *in.Mean > 0 && attackers > 0 && m.Health > 0 {
		c.MobHit = *in.Mean / 100 * m.Health * c.MobSpeed / attackers / max(landed, 0.3)
	} else {
		dps := ec.autoAttackDPS
		if c.Elite {
			dps *= 2.5
		}
		c.MobHit = dps * c.MobSpeed
	}
	c.MobAPEquiv = enemyAPCoefficient / 1.17
	// Stats: live values first.
	p := msg.Player
	c.AP = f(p.AttackPower, m.AP)
	c.Crit = f(p.CritPct, m.CritPct)
	if p.Swing != nil {
		if sw, ok := p.Swing["main"]; ok && sw.Speed != nil && *sw.Speed > 0 {
			m2 := *m
			m2.MH.Speed = *sw.Speed
			if sw2, ok := p.Swing["off"]; ok && sw2.Speed != nil && *sw2.Speed > 0 && m.DW {
				m2.OH.Speed = *sw2.Speed
			}
			c.M = &m2
		}
	}
	s := &State{queued: -1, stance: StanceBattle, mhAt: -1, ohAt: -1, mobAt: -1, amAt: -1}
	s.rage = f(p.Rage, 0)
	s.gcdAt = f(p.GCDRemaining, 0)
	if p.Stance != nil && *p.Stance >= 1 && *p.Stance <= 3 {
		s.stance = int(*p.Stance)
	}
	if p.Swing != nil {
		if sw, ok := p.Swing["main"]; ok && sw.Known && sw.Remaining != nil {
			s.mhAt = max(0, *sw.Remaining-f(sw.AgeMs, 0)/1000)
		}
		if sw, ok := p.Swing["off"]; ok && sw.Known && sw.Remaining != nil {
			s.ohAt = max(0, *sw.Remaining-f(sw.AgeMs, 0)/1000)
		}
	}
	c.build()
	// A normal enemy cannot outlive what the player alone would take to kill it: cap the console's time to die
	// (its health-trend estimate runs long on an enemy just picked up) with the model's own damage rate.
	if !c.Elite {
		if dps := c.physDPS(s); dps > 0 {
			hp := ec.health * c.HealthPct / 100
			capQ := [3]float64{0.8, 1, 1.3}
			for i := range c.TTK {
				c.TTK[i] = max(0.5, min(c.TTK[i], hp/dps*capQ[i]))
			}
			c.NextTTK = min(c.NextTTK, ec.health/dps)
		}
	}
	for i, ab := range c.abil {
		if i >= maxAbilities {
			warn = append(warn, "more abilities than the planner tracks")
			break
		}
		_ = ab
	}
	if p.QueuedSwing != nil {
		if ab := m.ByGameID(int32(*p.QueuedSwing)); ab != nil && ab.Kind == KindQueued {
			s.queued = c.idx(ab.Name)
		} else if name, r := ranks.ByID(int32(*p.QueuedSwing)); r != nil {
			s.queued = c.idx(name)
		}
		s.queuedTop = s.queued >= 0
	}
	// Auras by spell, any rank.
	auraLeft := func(a AuraState) float64 { return f(a.Remaining, 999) - f(a.AgeMs, 0)/1000 }
	for _, a := range p.Auras {
		name, _ := ranks.ByID(a.Spell)
		switch {
		case name == ranks.BattleShout:
			s.bsEnd = auraLeft(a)
			if ab := m.Get(ranks.BattleShout); ab != nil {
				c.BSAP = ab.BuffAP
			}
		case a.Spell == 29131:
			s.brEnd, s.brTick = auraLeft(a), 1
		case a.Spell == 12880:
			s.enrEnd = auraLeft(a)
		case a.Spell == 12966:
			s.flurry = int(f(a.Stacks, 3))
		case a.Spell == 12328:
			s.dwEnd = auraLeft(a)
		case a.Spell == 1719:
			s.reckEnd = auraLeft(a)
		case a.Spell == 2565:
			s.sbEnd, s.sbLeft = auraLeft(a), 2
		case a.Spell == 1282733:
			s.opEnd = auraLeft(a)
		}
	}
	if c.BSAP > 0 && s.bsEnd > 0 {
		c.AP -= c.BSAP
	}
	for _, a := range tg.Auras {
		name, _ := ranks.ByID(a.Spell)
		switch name {
		case ranks.SunderArmor:
			s.sunder, s.sunEnd = int(min(5, f(a.Stacks, 1))), auraLeft(a)
		case ranks.Rend:
			s.rendEnd, s.rendNext = auraLeft(a), math.Mod(auraLeft(a), 3)
		case ranks.ThunderClap:
			s.tcEnd = auraLeft(a)
		case ranks.DemoralizingShout:
			s.demoEnd = auraLeft(a)
		}
	}
	// Threat shape.
	secure := pol.SecureLeadPct
	switch {
	case pol.CeilingPct > 0:
		// The ceiling only applies while someone else holds aggro on the target.
		// It applies while someone else holds aggro, and also when the player has pulled it off another tank
		// (then the tank should get it back). With no other tank, a player holding aggro is the tank.
		ctx := msg.Intent.Context
		tankPresent := ctx != nil && ctx.OtherTankPresent != nil && *ctx.OtherTankPresent
		othersHold := (tg.TargetsPlayer != nil && !*tg.TargetsPlayer) || (tg.Threat != nil && tg.Threat.Status != nil && *tg.Threat.Status < 3)
		if (othersHold || tankPresent) && tg.Threat != nil && tg.Threat.LeadPct != nil {
			c.WThreatTarget = -pr.ceilingCost(tg.Threat, pol.CeilingPct, min(pr.HoldLook, c.TTK[1]))
			c.ThreatWeight = 1
		}
		cost, n := 0.0, 0
		for _, e := range msg.Enemies {
			n++
			hold := tankPresent || (e.TargetsPlayer != nil && !*e.TargetsPlayer) || (e.Threat != nil && e.Threat.Status != nil && *e.Threat.Status < 3)
			if hold && e.Threat != nil && e.Threat.LeadPct != nil {
				cost += pr.ceilingCost(e.Threat, pol.CeilingPct, min(pr.HoldLook, c.TTK[1]))
			}
		}
		if n > 0 && cost > 0 {
			c.WThreatOther = -cost / float64(n)
			c.ThreatWeight = 1
		}
	case pol.Weights.Threat > 0:
		look := min(pr.HoldLook, c.TTK[1])
		c.WThreatTarget = pr.holdWeight(tg.Threat, tg.TargetsPlayer, secure, look)
		other, n := 0.0, 0
		for _, e := range msg.Enemies {
			other += pr.holdWeight(e.Threat, e.TargetsPlayer, secure, look)
			n++
		}
		if n > 0 {
			c.WThreatOther = other / float64(n)
			// The next enemy becomes the focus when the target dies, and the group's damage follows it: threat
			// banked on it before then is worth more as the target's death nears.
			c.WThreatOther += pr.NextBoost * math.Exp(-c.TTK[1]/max(pr.NextScale, 0.1))
		} else {
			c.WThreatOther = pr.UnknownLead
		}
	}
	// Which abilities the player may press: listed, not manual-only. Cooldowns from the message.
	manual := map[int32]bool{}
	for _, id := range msg.ManualOnly {
		manual[id] = true
	}
	for _, a := range msg.Actions {
		ab := m.ByGameID(a.Spell)
		if ab == nil || manual[a.Spell] {
			continue
		}
		i := c.idx(ab.Name)
		if i < 0 || i >= maxAbilities {
			continue
		}
		c.allowed[i] = true
		for _, d := range pr.Disable {
			if string(ab.Name) == d {
				c.allowed[i] = false
			}
		}
		if a.Cooldown != nil {
			s.cdAt[i] = *a.Cooldown
		}
		if ab.NeedsWindow == windowRevenge && a.Usable != nil && *a.Usable {
			s.revEnd = 2.5 // window open, time left unknown
		}
		if ab.NeedsWindow == windowOverpower && a.Usable != nil && *a.Usable && s.opEnd == 0 {
			s.opEnd = 2.5
		}
	}
	// The decision-time state uses the median death time; each rollout draws its own.
	s.deathAt, s.ttkRef = c.TTK[1], c.TTK[1]
	s.targets, s.attackers, s.wt = c.Targets, c.Attackers, c.WThreatTarget
	if c.HealthPct > 20 {
		s.execAt = s.deathAt * (1 - 20/c.HealthPct)
	}
	c.Mu = c.ragePrice(s)
	return c, s, warn
}

// ragePrice is the base policy's value of a point of rage: what the best rage sink (Heroic Strike or Cleave,
// against the white swing it replaces) turns a point into, or a small floor without one.
func (c *Ctx) ragePrice(s *State) float64 {
	best := 0.0
	saved := c.Mu
	c.Mu = 0
	for i, ab := range c.abil {
		if ab.Kind != KindQueued || !ab.Modelled || !c.allowed[i] {
			continue
		}
		v := c.immediate(s, i)
		_, _, wr := c.whiteSwingValue(s)
		if v/(ab.Cost+wr) > best {
			best, c.sinkCost = v/(ab.Cost+wr), ab.Cost
		}
	}
	c.Mu = saved
	if best == 0 {
		// No sink: price rage by the best rage spender's value per rage.
		for i, ab := range c.abil {
			if !ab.Modelled || !c.allowed[i] || ab.Cost <= 0 || ab.Kind == KindQueued {
				continue
			}
			best = max(best, 0.5*c.immediate(s, i)/ab.Cost)
		}
	}
	return best
}

// Decide answers a decide message.
func Decide(m *Model, msg *DecideMsg, pr Params) DecideAnswer {
	start := time.Now()
	ans := DecideAnswer{Type: "decide", ID: msg.ID, Warnings: []string{}, Ranked: []RankedAction{}}
	c, s, warn := NewCtx(m, msg, pr)
	ans.Warnings = append(ans.Warnings, warn...)
	ans.IntentResolved = c.P.Answer()
	cands := c.candidates(s, msg)
	if debugDecide {
		fmt.Fprintf(os.Stderr, "special %+v mhwhite %+v front %v\n", c.special, c.mhWhite, c.Front)
		fmt.Fprintf(os.Stderr, "decide t=%.2f rage %.1f mu %.3f targets %d attackers %.1f wt %.2f wo %.2f ttk %.1f next %.1f\n", msg.T, s.rage, c.Mu, s.targets, s.attackers, s.wt, c.WThreatOther, c.TTK[1], c.NextTTK)
		for i, ab := range c.abil {
			if i < maxAbilities && c.allowed[i] && ab.Modelled && ab.Kind != KindStance {
				fmt.Fprintf(os.Stderr, "   %-18s usable %-5v imm %7.1f net %7.1f\n", ab.Name, c.usable(s, i), c.immediate(s, i), c.immediate(s, i)-c.Mu*ab.Cost)
			}
		}
	}
	budget := time.Duration(max(msg.BudgetMs, 1) * 0.7 * float64(time.Millisecond))
	seed := uint64(msg.ID)*0x2545F4914F6CDD1D + 1
	lo, hi := pr.MinSamples, pr.MaxSamples
	if pr.FixedSamples > 0 {
		lo, hi, budget = pr.FixedSamples, pr.FixedSamples, time.Hour
	}
	res := c.Search(s, cands, budget, lo, hi, seed)
	res = c.gcdFirst(s, res)
	for _, r := range res {
		ans.Ranked = append(ans.Ranked, r.toRanked(c))
	}
	if len(res) > 0 {
		ans.Best = res[0].toRanked(c).Action
		if len(res) > 1 {
			ans.Margin = round3(res[0].Value - res[1].Value)
		}
	}
	if debugDecide {
		for _, r := range res {
			name := "wait"
			if r.Cand.Act == actUnqueue {
				name = "unqueue"
			}
			if r.Cand.Act >= 0 {
				name = string(c.abil[r.Cand.Act].Name)
			}
			fmt.Fprintf(os.Stderr, "   => %-18s value %8.1f dmg %8.1f thr %8.1f mit %7.1f term %8.1f n %d\n", name, r.Value, r.Terms.Damage, r.Terms.Threat, r.Terms.Mitigation, r.Terms.Terminal, r.Samples)
		}
	}
	ans.Depth = 1
	ans.ElapsedMs = float64(time.Since(start).Microseconds()) / 1000
	return ans
}

func (r Result) toRanked(c *Ctx) RankedAction {
	out := RankedAction{Value: round3(r.Value), Terms: Terms{Damage: round3(r.Terms.Damage), Threat: round3(r.Terms.Threat),
		Mitigation: round3(r.Terms.Mitigation), Terminal: round3(r.Terms.Terminal)}}
	if r.Cand.Act == actUnqueue {
		out.Action = json.RawMessage(`"unqueue"`)
		return out
	}
	if r.Cand.Act == actWait {
		out.Action = json.RawMessage(`"wait"`)
		u := round3(r.Cand.WaitUntil)
		out.Until = &u
		return out
	}
	ab := c.abil[r.Cand.Act]
	out.Action = json.RawMessage(strconv.Itoa(int(ab.GameID)))
	return out
}

func round3(v float64) float64 { return math.Round(v*1000) / 1000 }

// candidates are the actions the player can take now, plus waiting.
func (c *Ctx) candidates(s *State, msg *DecideMsg) []Candidate {
	var out []Candidate
	borrow := map[int]int{}
	for _, a := range msg.Actions {
		if a.Borrow != nil && a.Borrow.Stance != nil {
			if ab := c.M.ByGameID(a.Spell); ab != nil {
				if i := c.idx(ab.Name); i >= 0 {
					borrow[i] = int(*a.Borrow.Stance)
				}
			}
		}
	}
	tg := msg.Target
	engaged := tg.Exists && ((tg.InMelee != nil && *tg.InMelee) || (tg.Threat != nil && tg.Threat.Status != nil) ||
		(msg.Player.GCDRemaining != nil) || (msg.Player.IncomingPctPerS != nil && msg.Player.IncomingPctPerS.Mean != nil))
	for i, ab := range c.abil {
		if i >= maxAbilities || !c.allowed[i] || ab.Kind == KindStance {
			continue
		}
		if !engaged && ab.Kind != KindBuff {
			continue // out of combat: only buffs; the console engages first
		}
		// With a gate, Heroic Strike and Cleave are queued only shortly before the swing.
		if ab.Kind == KindQueued && c.Pr.QueueGate > 0 && s.mhAt > c.Pr.QueueGate {
			continue
		}
		// An action clearly worth less than the rage it costs only wins a rollout comparison by noise.
		if c.Pr.CandFloor > -1 && ab.Kind != KindQueued && ab.Kind != KindRage && ab.Cost > 0 {
			if c.immediate(s, i)-c.Mu*ab.Cost < c.Pr.CandFloor*c.Mu*ab.Cost {
				continue
			}
		}
		cand := Candidate{Act: i}
		if b, ok := borrow[i]; ok && ab.Stances != 0 && ab.Stances&(1<<s.stance) == 0 {
			cand.Borrow = b
			trial := *s
			trial.stance = b
			trial.rage = min(trial.rage, c.M.TacticalMastery)
			if !c.usable(&trial, i) {
				continue
			}
		} else if !c.usable(s, i) {
			continue
		}
		out = append(out, cand)
	}
	// Wait until the next thing that can change the choice.
	until := 1.0
	if s.mhAt >= 0 {
		until = min(until, s.mhAt+0.05)
	}
	if s.gcdAt > 0 {
		until = min(until, s.gcdAt)
	}
	for i, ab := range c.abil {
		if i < maxAbilities && c.allowed[i] && ab.Modelled && s.cdAt[i] > 0 {
			until = min(until, s.cdAt[i])
		}
	}
	// A queued swing costs nothing until it lands, so cancelling it is an off-GCD action like any other.
	if c.Pr.Unqueue && s.queued >= 0 && engaged {
		out = append(out, Candidate{Act: actUnqueue})
	}
	out = append(out, Candidate{Act: actWait, WaitUntil: max(until, 0.1)})
	return out
}

// gcdFirst moves a global-cooldown ability ahead of a queue or cancel that only matches or narrowly beats it: the
// console presses one key a frame, and the swing can be queued or cancelled on a later frame while the global
// cooldown runs, unless the swing lands before then.
func (c *Ctx) gcdFirst(s *State, res []Result) []Result {
	if len(res) < 2 || s.mhAt >= 0 && s.mhAt <= 2*ConsoleCadence.Seconds() {
		return res
	}
	top := res[0].Cand.Act
	if top != actUnqueue && (top < 0 || c.abil[top].Kind != KindQueued) {
		return res
	}
	wait := math.Inf(-1)
	for _, r := range res {
		if r.Cand.Act == actWait {
			wait = r.Value
		}
	}
	for k := 1; k < len(res); k++ {
		a := res[k].Cand.Act
		if a < 0 || !c.abil[a].GCD {
			continue
		}
		if res[k].Value >= wait && res[0].Value-res[k].Value <= c.Pr.GCDFirst/100*abs(res[0].Value) {
			out := append([]Result{res[k]}, res[:k]...)
			return append(out, res[k+1:]...)
		}
		break
	}
	return res
}

// OnlyWait reports whether waiting is the only candidate (nothing usable now), without running the search.
func OnlyWait(m *Model, msg *DecideMsg, pr Params) bool {
	c, s, _ := NewCtx(m, msg, pr)
	cands := c.candidates(s, msg)
	return len(cands) == 1 && cands[0].Act == actWait
}

func (c *Ctx) String() string {
	return fmt.Sprintf("ctx(level %d vs %d, targets %d, attackers %.0f, mu %.2f)", c.M.Level, c.TLevel, c.Targets, c.Attackers, c.Mu)
}

var debugDecide = os.Getenv("PLANNER_DEBUG_DECIDE") != ""
