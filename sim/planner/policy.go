package planner

import (
	"github.com/wowsims/classic/sim/warrior/ranks"
)

// The base policy used inside rollouts: take the action with the highest immediate value net of a rage price (Mu),
// or wait. Its values are expectations at the current state; the rollouts themselves are sampled.

// value weights a damage, threat and mitigation triple by the intent.
func (c *Ctx) value(d, threat, mitig float64) float64 {
	return c.P.Weights.Damage*d + c.ThreatWeight*threat + c.P.Weights.Mitigation*mitig
}

// ttkLeft is the policy's own estimate of the target's remaining life (it does not see the sampled death time).
func (c *Ctx) ttkLeft(s *State) float64 {
	return max(0, s.ttkRef-s.t)
}

// whiteSwingValue is a main-hand white swing's expected damage, threat and rage at the state.
func (c *Ctx) whiteSwingValue(s *State) (d, threat, rage float64) {
	o := &c.mhWhite
	d = swingDamage(c.M.MH, c.ap(s), false) * c.physMult(s) * o.mult
	rage = c.M.MH.RagePerLanded * (o.land + o.crit)
	return d, d * c.stanceThreat(s), rage
}

// immediate is an ability's expected weighted value now, not counting its rage cost.
func (c *Ctx) immediate(s *State, i int) float64 {
	ab := c.abil[i]
	m := c.M
	ap := c.ap(s)
	st := c.stanceThreat(s)
	ttk := c.ttkLeft(s)
	wt, wo := s.wt, c.WThreatOther
	switch ab.Kind {
	case KindStrike:
		o := c.table(ab)
		d := c.abilityRaw(ab, ap) * c.physMult(s) * (o.land - o.crit + o.crit*ab.CritMult)
		th := (d*ab.ThreatMult + (ab.FlatThreat+ab.ThreatAP*ap)*o.land) * st
		return c.value(d, th*wt, 0) - c.swingLoss(s, ab)
	case KindExecute:
		o := c.table(ab)
		raw := (ab.Flat + ab.PerRage*max(0, s.rage-ab.Cost)) * ab.DamageMult
		d := raw * c.physMult(s) * (o.land - o.crit + o.crit*ab.CritMult)
		return c.value(d, d*ab.ThreatMult*st*wt, 0) - c.Mu*max(0, s.rage-ab.Cost)
	case KindQueued:
		o := &c.special
		n := 1
		if ab.Targets > 1 {
			n = min(ab.Targets, s.targets)
		}
		d := c.abilityRaw(ab, ap) * c.physMult(s) * (o.land - o.crit + o.crit*ab.CritMult)
		th := (d*ab.ThreatMult + ab.FlatThreat*o.land) * st
		wd, wth, wr := c.whiteSwingValue(s)
		v := c.value(d, th*wt, 0) - c.value(wd, wth*wt, 0) - c.Mu*wr
		if n > 1 {
			v += c.value(d, th*wo, 0) * float64(n-1)
		}
		return v
	case KindAoE:
		n := min(max(ab.Targets, 1), s.targets)
		o := c.table(ab)
		mult := c.physMult(s)
		d := c.abilityRaw(ab, ap) * mult * (o.land - o.crit + o.crit*ab.CritMult)
		if ab.Name == ranks.Whirlwind && m.DW {
			d += swingDamage(m.OH, ap, true) * m.OH.DamageMult * mult * (o.land - o.crit + o.crit*ab.CritMult)
		}
		th := d * ab.ThreatMult * st
		v := c.value(d, th*wt, 0) + c.value(d, th*wo, 0)*float64(n-1)
		if ab.Name == ranks.ThunderClap && s.t >= s.tcEnd {
			v += c.value(0, 0, c.incomingRate(s)*(1-1/1.2)*min(ab.Duration, max(ttk, 6)))
		}
		return v
	case KindDot:
		if s.t < s.rendEnd {
			return 0
		}
		ticks := min(float64(ab.Ticks), ttk/3)
		d := (ab.TickDamage + ab.TickAPCoef*ap) * ab.DamageMult * c.armorFreeMult(s) * c.tickCrit() * ticks * c.special.land
		return c.value(d, d*st*wt, 0)
	case KindDebuff:
		if ab.Name == ranks.DemoralizingShout {
			th := ab.FlatThreat * st * (wt + wo*float64(s.targets-1))
			mit := 0.0
			if s.t >= s.demoEnd-3 {
				mit = c.incomingRate(s) * ab.APReduction * c.MobAPEquiv * min(ab.Duration, max(ttk, 6))
			}
			return c.value(0, th, mit)
		}
		// Sunder Armor: threat now, and the next stack's armor reduction over the target's remaining life.
		o := &c.special
		th := (ab.FlatThreat + ab.ThreatAP*ap) * st * o.land
		stacks := s.sunder
		if s.t >= s.sunEnd {
			stacks = 0
		}
		gain := 0.0
		if stacks < 5 {
			gain = (c.armorDR[stacks+1]/c.armorDR[stacks] - 1) * c.Pr.ArmorScale
		}
		d := c.physDPS(s) * gain * min(ab.Duration, ttk) * o.land
		return c.value(d, (th+d*st)*wt, 0)
	case KindBuff:
		left := s.bsEnd - s.t
		if left > 3 {
			return 0
		}
		gainAP := ab.BuffAP
		dur := min(ab.Duration, max(ttk, 0)+20)
		d := c.dpsPerAP(s) * gainAP * dur
		return c.value(d, d*st*wt, 0)
	case KindRage:
		if s.rage > 60 {
			return 0
		}
		return c.Mu * (ab.RageGain + ab.RagePerSec*ab.Duration) * 0.8
	case KindBlock:
		if s.t < s.sbEnd || s.attackers <= 0 {
			return 0
		}
		// The block charge lands if an attack comes in the window and is not avoided first.
		hits := min(ab.Duration*s.attackers/c.MobSpeed, float64(max(ab.BlockCharges, 1))) * (1 - m.MissPct/100 - m.DodgePct/100 - m.ParryPct/100)
		blocked := min(c.MobHit, m.BlockValue) * hits
		rev := 0.0
		if ri := c.idx(ranks.Revenge); ri >= 0 && s.t >= s.revEnd && c.allowed[ri] {
			if rv := c.abil[ri]; rv.Stances == 0 || rv.Stances&(1<<s.stance) != 0 {
				rev = hits * 0.75 * max(0, c.immediate(s, ri)-c.Mu*rv.Cost)
			}
		}
		return c.value(0, 0, blocked) + rev
	}
	return 0
}

// incomingRate is raw incoming melee damage per second.
func (c *Ctx) incomingRate(s *State) float64 {
	if s.attackers <= 0 {
		return 0
	}
	return c.MobHit * s.attackers / c.MobSpeed * (1 - c.M.MissPct/100 - c.M.DodgePct/100 - c.M.ParryPct/100)
}

// physDPS is a rough physical damage rate for valuing armor reduction: white swings plus a third for specials.
func (c *Ctx) physDPS(s *State) float64 {
	ap := c.ap(s)
	d := swingDamage(c.M.MH, ap, false) * c.mhWhite.mult / c.M.MH.Speed
	if c.M.DW {
		d += swingDamage(c.M.OH, ap, false) * c.M.OH.DamageMult * c.ohWhite.mult / c.M.OH.Speed
	}
	return d * 1.33 * c.physMult(s)
}

func (c *Ctx) dpsPerAP(s *State) float64 {
	per := c.mhWhite.mult / 14
	if c.M.DW {
		per += c.ohWhite.mult * c.M.OH.DamageMult / 14
	}
	return per * 1.33 * c.physMult(s)
}

// basePolicy picks the next action, or actNone to wait for the next event.
func (c *Ctx) basePolicy(s *State) int {
	best, bestV := actNone, 0.0
	// A queued Heroic Strike or Cleave keeps its rage until it lands.
	held := 0.0
	if s.queued >= 0 {
		held = c.abil[s.queued].Cost
	}
	for i, ab := range c.abil {
		if ab.Kind == KindStance || ab.Kind == KindQueued || !c.usable(s, i) {
			continue
		}
		if ab.Cost > 0 && s.rage-ab.Cost < held {
			continue
		}
		v := c.immediate(s, i) - c.Mu*ab.Cost
		if v > bestV {
			best, bestV = i, v
		}
	}
	if best >= 0 {
		return best
	}
	// Heroic Strike or Cleave: the rage sink. Queue the better one, when it beats the white swing it replaces and
	// the rage left covers the next GCD's likely spend. (Its rage is not priced here: the price comes from the sink.)
	if s.queued < 0 && s.t < s.deathAt && s.mhAt-s.t <= 0.5 {
		reserve := c.reserve(s)
		for i, ab := range c.abil {
			if ab.Kind != KindQueued || !c.usable(s, i) || s.rage-ab.Cost < reserve {
				continue
			}
			v := c.immediate(s, i)
			if v > bestV {
				best, bestV = i, v
			}
		}
	}
	return best
}

// reserve is rage to keep for on-GCD abilities coming off cooldown within the next swing.
func (c *Ctx) reserve(s *State) float64 {
	horizon := s.t + c.M.MH.Speed
	r := 0.0
	for i, ab := range c.abil {
		if !ab.GCD || ab.Kind == KindBuff || ab.Kind == KindDebuff || ab.Kind == KindDot || !ab.Modelled || ab.CD == 0 {
			continue
		}
		if ab.Stances != 0 && ab.Stances&(1<<s.stance) == 0 {
			continue
		}
		if ab.NeedsWindow == windowRevenge && s.t >= s.revEnd {
			continue
		}
		if ab.NeedsWindow == windowOverpower && s.t >= s.opEnd {
			continue
		}
		if s.cdAt[i] <= horizon {
			r = max(r, ab.Cost)
		}
	}
	return r
}

// terminal values what the state carries past the horizon: rage, the rest of a Rend, and debuffs and buffs.
func (c *Ctx) terminal(s *State) float64 {
	alive := max(0, s.deathAt-s.t)
	// Leftover rage is worth the sink's price for what can still be spent before the enemies are dead: one sink
	// per main-hand swing and a GCD ability every 1.5 s at most.
	left := alive
	if s.targets > 1 {
		left += float64(s.targets-1) * c.NextTTK
	}
	if c.P.Weights.Threat > 0 {
		left = max(left, 10) // holding aggro: the pull goes on
	}
	spendable := left * (c.sinkCost/max(c.M.MH.Speed, 1) + 10/1.5)
	v := c.Mu * min(s.rage, spendable, 70)
	if ab := c.M.Get(ranks.Rend); ab != nil && s.rendEnd > s.t {
		ticks := min(s.rendEnd-s.t, alive) / 3
		d := (ab.TickDamage + ab.TickAPCoef*c.ap(s)) * ab.DamageMult * c.armorFreeMult(s) * c.tickCrit() * ticks
		v += c.value(d, d*c.stanceThreat(s)*s.wt, 0)
	}
	if s.sunder > 0 && s.sunEnd > s.t {
		gain := (c.armorDR[s.sunder]/c.armorDR[0] - 1) * c.Pr.ArmorScale
		d := c.physDPS(s) / c.armorDR[s.sunder] * c.armorDR[0] * gain * min(s.sunEnd-s.t, alive)
		v += c.value(d, d*c.stanceThreat(s)*s.wt, 0)
	}
	if s.tcEnd > s.t {
		v += c.value(0, 0, c.incomingRate(s)*(1-1/1.2)*min(s.tcEnd-s.t, alive))
	}
	if s.demoEnd > s.t {
		if ab := c.M.Get(ranks.DemoralizingShout); ab != nil {
			v += c.value(0, 0, c.incomingRate(s)*ab.APReduction*c.MobAPEquiv*min(s.demoEnd-s.t, alive))
		}
	}
	if ab := c.M.Get(ranks.BattleShout); ab != nil && s.bsEnd > s.t {
		d := c.dpsPerAP(s) * ab.BuffAP * min(s.bsEnd-s.t, alive+20)
		v += c.value(d, 0, 0)
	}
	return v
}

// swingLoss is the value of the white swings a cast that restarts the swing timers pushes back.
func (c *Ctx) swingLoss(s *State, ab *Ability) float64 {
	if ab.Cast <= 0 || ab.KeepSwings {
		return 0
	}
	m := c.M
	wd, wth, wr := c.whiteSwingValue(s)
	per := c.value(wd, wth*s.wt, 0) + c.Mu*wr
	mhAt := s.mhAt
	if mhAt < s.t {
		mhAt = s.t + m.MH.Speed/2
	}
	loss := max(0, s.t+ab.Cast+m.MH.Speed-mhAt) / m.MH.Speed * per
	if m.DW {
		ohAt := s.ohAt
		if ohAt < s.t {
			ohAt = s.t + m.OH.Speed/2
		}
		ohPer := per * m.OH.DamageMult * m.OH.Avg() / max(m.MH.Avg(), 1)
		loss += max(0, s.t+ab.Cast+1.5*m.OH.Speed-ohAt) / m.OH.Speed * ohPer
	}
	return loss * c.Pr.SwingLossScale
}

// tickCrit is the expected multiplier of Forever's critting Rend ticks.
func (c *Ctx) tickCrit() float64 {
	return 1 + c.special.crit/max(c.special.land, 1e-9)*(c.M.SpecialCritMult-1)
}
