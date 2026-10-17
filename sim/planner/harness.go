package planner

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/warrior"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

// Partner is the other group member in a scenario: a damage dealer competing with a tanking Warrior, or a tank
// holding aggro while the Warrior deals damage. Zero DPS means solo.
type Partner struct {
	Role     string  // "dps" or "tank"
	DPS      float64 // damage per second on the focused target
	TPS      float64 // threat per second on the focused target
	AoEShare float64 // share of TPS (and damage) also put on each other enemy
	Delay    float64 // seconds after the pull before the partner starts
	Noise    float64 // relative per-step noise of the partner's damage and threat
	Ranged   bool    // a ranged partner pulls aggro at 130%, a melee one at 110%
}

// Harness is the scenario's world: every enemy's health, who holds its aggro (the 110%/130% rule against the
// partner), and the fight's outcome. Enemies die at zero health and stop attacking; the fight ends when all are dead.
type Harness struct {
	W       *warrior.Warrior
	P       Partner
	HP      []float64
	targets []*core.Unit

	Dealt, Useful []float64 // the Warrior's damage per enemy, all and up to its death
	PartnerDealt  []float64
	My, Them      []float64
	leadHist      [][8]float64 // lead % over the last 2 s, one sample per tick
	leadAt        int
	lifeHist      [][24]float64 // health lost, sampled every tick over the last six seconds
	theyHave      []bool
	Dead          []bool

	KillTime     float64
	LossFocus    float64 // tank: seconds the partner held the player's current target
	LossTime     float64 // tank: enemy-seconds held by the partner; damage dealer: enemy-seconds held by the Warrior
	Swaps        int
	Taken        float64 // damage the Warrior took
	OnHit        func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult, periodic bool)
	ForceExecute bool
	DamageBy     map[int32]float64 // useful damage by spell ID (0 = white swings)
	Results      map[int32]int     // attack results (any outcome, per target) by spell ID
	events       []takenEvent
}

type takenEvent struct {
	at     time.Duration
	amount float64
}

func newHarness(w *warrior.Warrior, p Partner, hp []float64) *Harness {
	h := &Harness{W: w, P: p, HP: hp}
	w.RegisterAura(core.Aura{
		Label:    "Planner scenario",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
			h.reset(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			h.hit(sim, spell, result)
			if h.OnHit != nil {
				h.OnHit(sim, spell, result, false)
			}
		},
		OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			h.hit(sim, spell, result)
			if h.OnHit != nil {
				h.OnHit(sim, spell, result, true)
			}
		},
		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.Damage > 0 {
				h.Taken += result.Damage
				h.events = append(h.events, takenEvent{sim.CurrentTime, result.Damage})
			}
		},
	})
	return h
}

func (h *Harness) tankRole() bool { return h.P.Role != "tank" }

func (h *Harness) reset(sim *core.Simulation) {
	h.targets = sim.Encounter.TargetUnits
	n := len(h.targets)
	h.Dealt, h.Useful, h.PartnerDealt = make([]float64, n), make([]float64, n), make([]float64, n)
	h.My, h.Them, h.theyHave, h.Dead = make([]float64, n), make([]float64, n), make([]bool, n), make([]bool, n)
	h.leadHist, h.leadAt = make([][8]float64, n), 0
	h.lifeHist = make([][24]float64, n)
	h.KillTime, h.LossTime, h.Swaps, h.Taken, h.events, h.LossFocus = 0, 0, 0, 0, h.events[:0], 0
	h.DamageBy = map[int32]float64{}
	h.Results = map[int32]int{}
	for i := range h.targets {
		h.theyHave[i] = !h.tankRole()
		if len(h.HP) <= i {
			h.HP = append(h.HP, h.HP[len(h.HP)-1])
		}
	}
	h.W.CurrentTarget = h.targets[0]
	step := 250 * time.Millisecond
	core.StartPeriodicAction(sim, core.PeriodicActionOptions{
		Period: step,
		OnAction: func(sim *core.Simulation) {
			h.tick(sim, step.Seconds())
		},
	})
	if !h.tankRole() {
		// A tank partner has every enemy from the pull.
		sim.AddPendingAction(&core.PendingAction{NextActionAt: 0, OnAction: func(sim *core.Simulation) {
			for _, t := range h.targets {
				t.AutoAttacks.CancelAutoSwing(sim)
			}
		}})
	}
}

func (h *Harness) index(u *core.Unit) int {
	for i, t := range h.targets {
		if t == u {
			return i
		}
	}
	return -1
}

func (h *Harness) hit(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
	i := h.index(result.Target)
	if i < 0 {
		return
	}
	h.My[i] += result.Threat
	if h.Results != nil {
		h.Results[spellKey(spell)]++
	}
	if result.Damage > 0 {
		h.Dealt[i] += result.Damage
		if !h.Dead[i] {
			u := min(result.Damage, h.left(i))
			h.Useful[i] += u
			if h.DamageBy != nil {
				h.DamageBy[spellKey(spell)] += u
			}
			h.checkDeath(sim, i)
		}
	}
}

func (h *Harness) left(i int) float64 { return max(0, h.HP[i]-h.Useful[i]-h.PartnerDealt[i]) }

// HealthFrac is enemy i's remaining health share.
func (h *Harness) HealthFrac(i int) float64 { return h.left(i) / h.HP[i] }

func (h *Harness) checkDeath(sim *core.Simulation, i int) {
	if h.Dead[i] || h.left(i) > 0 {
		return
	}
	h.Dead[i] = true
	h.targets[i].AutoAttacks.CancelAutoSwing(sim)
	alive := -1
	for j := range h.targets {
		if !h.Dead[j] {
			alive = j
			break
		}
	}
	if alive < 0 {
		h.KillTime = sim.CurrentTime.Seconds()
		sim.EndCombatNow()
		return
	}
	if h.W.CurrentTarget == h.targets[i] {
		h.W.CurrentTarget = h.targets[alive]
	}
}

func (h *Harness) tick(sim *core.Simulation, dt float64) {
	if h.KillTime > 0 {
		return
	}
	t := sim.CurrentTime.Seconds()
	focus := h.index(h.W.CurrentTarget)
	if h.P.DPS > 0 && t >= h.P.Delay {
		for i := range h.targets {
			if h.Dead[i] {
				continue
			}
			share := h.P.AoEShare
			if i == focus {
				share = 1
			}
			if share == 0 {
				continue
			}
			noise := 1 + h.P.Noise*(2*sim.RandomFloat("partner")-1)
			h.Them[i] += h.P.TPS * share * noise * dt
			h.PartnerDealt[i] += h.P.DPS * share * noise * dt
			h.checkDeath(sim, i)
			if h.KillTime > 0 {
				return
			}
		}
	}
	for i := range h.targets {
		_, lead, _ := h.ThreatState(i)
		h.leadHist[i][h.leadAt%8] = lead
		h.lifeHist[i][h.leadAt%24] = h.HP[i] - h.left(i)
	}
	h.leadAt++
	pull := 1.1
	if h.P.Ranged {
		pull = 1.3
	}
	for i, target := range h.targets {
		if h.Dead[i] {
			continue
		}
		if h.theyHave[i] {
			if h.tankRole() {
				h.LossTime += dt
				if target == h.W.CurrentTarget {
					h.LossFocus += dt
				}
			}
			if h.My[i] > pull*h.Them[i] && h.My[i] > 0 {
				if debugPress {
					fmt.Printf("  ** %.2f regained enemy %d\n", sim.CurrentTime.Seconds(), i)
				}
				h.theyHave[i] = false
				h.Swaps++
				target.AutoAttacks.EnableAutoSwing(sim)
			}
		} else {
			if !h.tankRole() {
				h.LossTime += dt
			}
			if h.P.DPS > 0 && h.Them[i] > pull*h.My[i] && h.Them[i] > 0 {
				if debugPress {
					fmt.Printf("  ** %.2f partner takes enemy %d (focus %v): mine %.0f theirs %.0f\n", sim.CurrentTime.Seconds(), i, target == h.W.CurrentTarget, h.My[i], h.Them[i])
				}
				h.theyHave[i] = true
				h.Swaps++
				target.AutoAttacks.CancelAutoSwing(sim)
			}
		}
	}
}

// ThreatState is what the bus reports for enemy i: status 0-3, lead % and whether it targets the player.
func (h *Harness) ThreatState(i int) (status float64, lead float64, targetsPlayer bool) {
	my, them := h.My[i], h.Them[i]
	lead = 999
	if them > 0 {
		lead = min(999, my/them*100)
	}
	if !h.theyHave[i] {
		if my >= them {
			return 3, lead, true
		}
		return 2, lead, true
	}
	if my > them {
		return 1, lead, false
	}
	return 0, lead, false
}

// UsefulDamage is the Warrior's damage that went into living enemies.
func (h *Harness) UsefulDamage() float64 {
	s := 0.0
	for _, v := range h.Useful {
		s += v
	}
	return s
}

// LeadTrend is enemy i's lead change in % per second over the last two seconds, as the console estimates it.
func (h *Harness) LeadTrend(i int) float64 {
	if h.leadAt < 8 {
		return 0
	}
	now := h.leadHist[i][(h.leadAt-1)%8]
	then := h.leadHist[i][h.leadAt%8]
	return (min(now, 400) - min(then, 400)) / 1.75
}

func spellKey(spell *core.Spell) int32 {
	// Heroic Strike and Cleave replace the main-hand swing and carry its proc mask; keep them apart.
	if a, r := ranks.ByID(spell.ActionID.SpellID); r != nil && (a == ranks.HeroicStrike || a == ranks.Cleave) {
		return spell.ActionID.SpellID
	}
	if spell.ProcMask.Matches(core.ProcMaskMeleeMHAuto) {
		return -1
	}
	if spell.ProcMask.Matches(core.ProcMaskMeleeOHAuto) {
		return -2
	}
	return spell.ActionID.SpellID
}

// RecentRate is enemy i's health loss per second over up to the last six seconds, and the window's length.
func (h *Harness) RecentRate(i int) (float64, float64) {
	if h.leadAt == 0 {
		return 0, 0
	}
	n := min(h.leadAt, 24)
	now := h.lifeHist[i][(h.leadAt-1)%24]
	then := 0.0
	if h.leadAt > 24 {
		then = h.lifeHist[i][h.leadAt%24]
	}
	w := float64(n) * 0.25
	return (now - then) / w, w
}
