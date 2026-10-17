package planner

import (
	"fmt"
	"os"
	"sort"
	"time"
)

// Candidate is a first action the search compares: an ability (possibly borrowing a stance), or waiting.
type Candidate struct {
	Act       int     // ability index, actWait or actUnqueue
	WaitUntil float64 // for actWait, seconds from now
	Borrow    int     // stance to switch to first (0 = none)
}

type Result struct {
	Cand    Candidate
	Value   float64
	Terms   Terms
	Samples int
}

// ttkWeights are the probabilities given to the p10, p50 and p90 death times.
var ttkWeights = [3]float64{0.3, 0.4, 0.3}

// Search runs rollouts for every candidate with common random numbers until the budget is spent (at least
// minSamples per candidate), and returns the candidates by value. init must not be modified by the caller during
// the search. Sample k uses the death time p10/p50/p90 in turn and its own random stream, the same for every
// candidate.
func (c *Ctx) Search(init *State, cands []Candidate, budget time.Duration, minSamples, maxSamples int, seed uint64) []Result {
	start := time.Now()
	res := make([]Result, len(cands))
	weights := make([]float64, len(cands))
	for i := range cands {
		res[i].Cand = cands[i]
	}
	for k := 0; k < maxSamples; k++ {
		if k >= minSamples && k%3 == 0 && time.Since(start) > budget {
			break
		}
		j := k % 3
		w := ttkWeights[j]
		sampleSeed := seed ^ (uint64(k+1) * 0x9E3779B97F4A7C15)
		for i := range cands {
			s := *init
			s.rng = sampleSeed | 1
			c.prepareSample(&s, j)
			c.trace = debugRollout && k < 2
			if c.trace {
				fmt.Fprintf(os.Stderr, "   rollout cand %d act %d mh %.2f death %.1f\n", i, cands[i].Act, s.mhAt, s.deathAt)
			}
			c.runCandidate(&s, cands[i])
			if c.trace {
				fmt.Fprintf(os.Stderr, "   end dmg %.1f rage %.1f t %.2f targets %d\n", s.dmg, s.rage, s.t, s.targets)
			}
			c.trace = false
			term := c.terminal(&s)
			r := &res[i]
			r.Terms.Damage += w * c.P.Weights.Damage * s.dmg
			r.Terms.Threat += w * c.ThreatWeight * s.threat
			r.Terms.Mitigation += w * c.P.Weights.Mitigation * s.mitig
			r.Terms.Terminal += w * term
			r.Samples++
			weights[i] += w
		}
	}
	for i := range res {
		r := &res[i]
		if weights[i] > 0 {
			r.Terms.Damage /= weights[i]
			r.Terms.Threat /= weights[i]
			r.Terms.Mitigation /= weights[i]
			r.Terms.Terminal /= weights[i]
		}
		r.Value = r.Terms.Damage + r.Terms.Threat + r.Terms.Mitigation + r.Terms.Terminal
		if r.Cand.Act == actWait {
			// Acting must beat waiting by a margin: close calls are mostly rollout noise.
			r.Value += c.Pr.WaitBias / 100 * abs(r.Value)
		}
		if r.Cand.Act == actUnqueue {
			r.Value -= c.Pr.UnqueueBias / 100 * abs(r.Value)
		}
	}
	sort.SliceStable(res, func(a, b int) bool { return res[a].Value > res[b].Value })
	return res
}

// prepareSample draws a sample's unknowns: the death time, and swing and mob phases the state leaves unknown (<0).
func (c *Ctx) prepareSample(s *State, j int) {
	s.deathAt = c.TTK[j]
	if c.HealthPct > 20 {
		s.execAt = s.deathAt * (1 - 20/c.HealthPct)
	} else {
		s.execAt = 0
	}
	if s.mhAt < 0 {
		s.mhAt = s.rand() * c.M.MH.Speed
	}
	if c.M.DW && s.ohAt < 0 {
		s.ohAt = s.rand() * c.M.OH.Speed
	}
	if s.mobAt < 0 {
		s.mobAt = s.rand() * c.MobSpeed / max(s.attackers, 1)
	}
	if c.M.AngerManagement && s.amAt < 0 {
		s.amAt = s.rand() * 3
	}
}

func (c *Ctx) runCandidate(s *State, cand Candidate) {
	if cand.Act == actWait {
		c.run(s, actWait, cand.WaitUntil)
		return
	}
	if cand.Act == actUnqueue {
		s.queued, s.queuedTop, s.requeueAt = -1, false, s.t+0.4
		s.press()
		c.run(s, actNone, 0)
		return
	}
	if cand.Borrow != 0 && cand.Borrow != s.stance {
		home := s.stance
		s.stance = cand.Borrow
		s.rage = min(s.rage, c.M.TacticalMastery)
		s.press()
		s.home, s.homeAt = home, 1.5 // the console returns home after the global cooldown
	}
	c.run(s, cand.Act, 0)
}

var debugRollout = os.Getenv("PLANNER_DEBUG_ROLLOUT") != ""

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
