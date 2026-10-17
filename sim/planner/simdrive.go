package planner

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/warrior"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

// The simulator side: a console emulator that builds the same decide message the SecretProbe console builds, from the
// engine's state, and presses the chosen action the way the console does (borrowing a stance, queuing Heroic
// Strike, returning home). A Chooser picks the action: the planner, or the console's priority list for comparison.

// Chooser picks an action from a decide message: a game spell ID, or 0 to wait (until > 0: seconds).
type Chooser interface {
	Choose(msg *DecideMsg) (spell int32, until float64)
}

// DefaultManual are the spells the console keeps manual unless the player changes it.
var DefaultManual = []ranks.Ability{ranks.ShieldWall, ranks.LastStand, ranks.Recklessness, ranks.DeathWish, ranks.Retaliation,
	ranks.ChallengingShout, ranks.IntimidatingShout, ranks.Charge, ranks.Intercept, ranks.Pummel, ranks.ShieldBash,
	ranks.Hamstring, ranks.Disarm, ranks.PiercingHowl, ranks.MockingBlow, ranks.BerserkerRage, ranks.Taunt, ranks.SweepingStrikes}

// DriveConfig is a scenario's console setup.
type DriveConfig struct {
	Intent     Intent
	HomeStance int
	Harness    *Harness
	TTKNoise   float64 // relative spread given to the time-to-die estimate
	PriorDPS   float64 // the console's damage-rate guess before the fight shows one
	Elite      bool
	TTKUnknown float64 // seconds after an enemy first takes damage before its time to die is known (0 = always known)
}

// Driver is one Warrior's console emulator in the simulator.
type Driver struct {
	cfg     DriveConfig
	w       *warrior.Warrior
	model   *Model
	chooser Chooser
	spells  map[ranks.Ability]*warrior.WarriorSpell
	manual  map[ranks.Ability]bool

	lastT       time.Duration
	waitUntil   time.Duration
	borrowedAt  time.Duration
	borrowed    bool
	seq         int64
	lastRage    float64
	hurtAt      map[int]time.Duration // when each enemy was first seen damaged, this fight
	noQueueAt   time.Duration         // after a cancel the console presses the queued swing again no sooner than this
	Decisions   int
	Sim         *core.Simulation
	Casts       map[ranks.Ability]int
	Trace       func(sim *core.Simulation, msg *DecideMsg, spell int32, until float64)
	ChooseNanos int64
}

var (
	driversMu sync.Mutex
	drivers   = map[*warrior.Warrior]*Driver{}
)

// newDriver registers the driver's auras on a Warrior during Initialize (warrior.SetupHook).
func newDriver(w *warrior.Warrior, cfg DriveConfig) *Driver {
	d := &Driver{cfg: cfg, w: w, manual: map[ranks.Ability]bool{}}
	for _, a := range DefaultManual {
		d.manual[a] = true
	}
	// The live console asks every frame in combat (about every 50 ms). While the global cooldown is ready the rotation
	// loop already runs at that cadence; while it runs, this asks for off-GCD actions (Heroic Strike, cancelling it).
	w.RegisterResetEffect(func(sim *core.Simulation) {
		d.lastT, d.waitUntil, d.borrowed, d.lastRage, d.noQueueAt = 0, 0, false, -1, 0
		core.StartPeriodicAction(sim, core.PeriodicActionOptions{
			Period: ConsoleCadence,
			OnAction: func(sim *core.Simulation) {
				if !w.GCD.IsReady(sim) {
					d.act(sim, true)
				}
			},
		})
	})
	driversMu.Lock()
	drivers[w] = d
	driversMu.Unlock()
	return d
}

func hook(w *warrior.Warrior, sim *core.Simulation) {
	driversMu.Lock()
	d := drivers[w]
	driversMu.Unlock()
	if d != nil {
		d.act(sim, false)
	}
}

func (d *Driver) ensureModel(sim *core.Simulation) {
	if d.spells != nil {
		return
	}
	w := d.w
	d.spells = map[ranks.Ability]*warrior.WarriorSpell{
		ranks.HeroicStrike: w.HeroicStrikeQueue, ranks.Cleave: w.CleaveQueue, ranks.SunderArmor: w.SunderArmor, ranks.Revenge: w.Revenge,
		ranks.ThunderClap: w.ThunderClap, ranks.DemoralizingShout: w.DemoralizingShout, ranks.BattleShout: w.BattleShout,
		ranks.Bloodthirst: w.Bloodthirst, ranks.MortalStrike: w.MortalStrike, ranks.Whirlwind: w.Whirlwind, ranks.Slam: w.Slam,
		ranks.Overpower: w.Overpower, ranks.Execute: w.Execute, ranks.Rend: w.Rend, ranks.ShieldSlam: w.ShieldSlam,
		ranks.Hamstring: w.Hamstring, ranks.Bloodrage: w.Bloodrage, ranks.ShieldBlock: w.ShieldBlock,
		ranks.BattleStance: w.BattleStance, ranks.DefensiveStance: w.DefensiveStance, ranks.BerserkerStance: w.BerserkerStance,
	}
	for k, v := range d.spells {
		if v == nil {
			delete(d.spells, k)
		}
	}
}

// act is one console frame: return home if a borrow is over, else ask the chooser and press.
func (d *Driver) act(sim *core.Simulation, offGCD bool) {
	d.Sim = sim
	if sim.CurrentTime < d.lastT {
		d.waitUntil, d.borrowed = 0, false
		clear(d.hurtAt)
	}
	d.lastT = sim.CurrentTime
	d.ensureModel(sim)
	w := d.w
	if w.Hardcast.Expires > sim.CurrentTime {
		return
	}
	// Return to the home stance once the borrowed press's global cooldown has passed.
	if d.borrowed && sim.CurrentTime >= d.borrowedAt+1500*time.Millisecond {
		if home := d.stanceSpell(d.cfg.HomeStance); home != nil && !w.StanceMatches(stanceBit(d.cfg.HomeStance)) {
			if home.CanCast(sim, w.CurrentTarget) {
				home.Cast(sim, w.CurrentTarget)
				d.borrowed = false
				return
			}
		} else {
			d.borrowed = false
		}
	}
	rage := w.CurrentRage()
	if !offGCD && sim.CurrentTime < d.waitUntil && rage == d.lastRage {
		return
	}
	d.lastRage = rage
	msg := d.message(sim)
	start := time.Now()
	spell, until := d.chooser.Choose(msg)
	d.ChooseNanos += time.Since(start).Nanoseconds()
	d.Decisions++
	if d.Trace != nil {
		d.Trace(sim, msg, spell, until)
	}
	if spell == 0 {
		if until > 0 && !offGCD {
			d.waitUntil = sim.CurrentTime + min(time.Duration(until*float64(time.Second)), ConsoleCadence) // the console asks again every frame
		}
		return
	}
	if spell == unqueueSpell {
		if w.CancelQueuedSwing(sim) {
			d.count(ranks.Ability("Unqueue"))
			d.noQueueAt = sim.CurrentTime + 400*time.Millisecond
		}
		return
	}
	d.press(sim, spell)
}

func stanceBit(s int) warrior.Stance {
	switch s {
	case StanceDefensive:
		return warrior.DefensiveStance
	case StanceBerserker:
		return warrior.BerserkerStance
	}
	return warrior.BattleStance
}

func (d *Driver) stanceSpell(s int) *warrior.WarriorSpell {
	switch s {
	case StanceDefensive:
		return d.w.DefensiveStance
	case StanceBerserker:
		return d.w.BerserkerStance
	case StanceBattle:
		return d.w.BattleStance
	}
	return nil
}

// press casts the chosen ability, borrowing the first allowed stance when needed.
func (d *Driver) press(sim *core.Simulation, gameID int32) {
	ab := d.model.ByGameID(gameID)
	if ab == nil {
		return
	}
	ws := d.spells[ab.Name]
	if ws == nil {
		return
	}
	if ab.Kind == KindQueued && sim.CurrentTime < d.noQueueAt {
		return
	}
	w := d.w
	if ab.Kind == KindStance {
		if ws.CanCast(sim, w.CurrentTarget) {
			ws.Cast(sim, w.CurrentTarget)
			d.borrowed = false // a stance chosen outright is not a borrow
			d.count(ab.Name)
		} else if debugPress {
			fmt.Printf("  stance %s not castable: ready %v cd %v casting %v gcd %v\n", ab.Name, ws.IsReady(sim), ws.TimeToReady(sim), w.IsCasting(sim), w.GCD.IsReady(sim))
		}
		return
	}
	if ab.Stances != 0 && ab.Stances&(1<<d.stanceNow()) == 0 {
		for s := StanceBattle; s <= StanceBerserker; s++ {
			if ab.Stances&(1<<s) != 0 {
				sp := d.stanceSpell(s)
				if sp == nil || !sp.CanCast(sim, w.CurrentTarget) {
					return
				}
				sp.Cast(sim, w.CurrentTarget)
				d.borrowed, d.borrowedAt = true, sim.CurrentTime
				break
			}
		}
	}
	if ws.CanCast(sim, w.CurrentTarget) {
		ws.Cast(sim, w.CurrentTarget)
		d.count(ab.Name)
		if d.borrowed {
			d.borrowedAt = sim.CurrentTime
		}
	}
}

func (d *Driver) stanceNow() int {
	switch {
	case d.w.StanceMatches(warrior.DefensiveStance):
		return StanceDefensive
	case d.w.StanceMatches(warrior.BerserkerStance):
		return StanceBerserker
	}
	return StanceBattle
}

func fp(v float64) *float64 { return &v }
func bp(v bool) *bool       { return &v }

// message builds the decide message the console would send now.
func (d *Driver) message(sim *core.Simulation) *DecideMsg {
	w := d.w
	m := d.model
	now := sim.CurrentTime
	d.seq++
	msg := &DecideMsg{Type: "decide", ID: d.seq, T: now.Seconds(), BudgetMs: 3, Intent: d.cfg.Intent}
	stance := float64(d.stanceNow())
	gcd := max(0, (w.GCD.ReadyAt() - now).Seconds())
	var queued *float64
	if q := w.QueuedSwingSpell(); q != nil {
		if q == w.Cleave || q == w.CleaveQueue {
			queued = fp(float64(m.Get(ranks.Cleave).GameID))
		} else {
			queued = fp(float64(m.Get(ranks.HeroicStrike).GameID))
		}
	}
	p := PlayerState{Rage: fp(w.CurrentRage()), RageMax: fp(100), Stance: &stance, GCDRemaining: &gcd, QueuedSwing: queued, Moving: bp(false),
		AttackPower: fp(w.GetStat(stats.AttackPower)), CritPct: fp(w.GetStat(stats.MeleeCrit) / core.CritRatingPerCritChance),
		HealthPct: fp(100 * w.CurrentHealthPercent()), Swing: map[string]SwingTimer{}}
	mhLeft := (w.AutoAttacks.MainhandSwingAt() - now).Seconds()
	p.Swing["main"] = SwingTimer{Remaining: fp(max(0, mhLeft)), Speed: fp(w.AutoAttacks.MainhandSwingSpeed().Seconds()), Known: true, AgeMs: fp(0)}
	if w.AutoAttacks.IsDualWielding {
		p.Swing["off"] = SwingTimer{Remaining: fp(max(0, (w.AutoAttacks.OffhandSwingAt() - now).Seconds())),
			Speed: fp(w.AutoAttacks.OffhandSwingSpeed().Seconds()), Known: true, AgeMs: fp(0)}
	}
	aura := func(a *core.Aura, gameID int32) {
		if a != nil && a.IsActive() {
			st := float64(a.GetStacks())
			p.Auras = append(p.Auras, AuraState{Spell: gameID, Remaining: fp(a.RemainingDuration(sim).Seconds()), Stacks: &st, AgeMs: fp(0)})
		}
	}
	aura(w.GetAura("Battle Shout"), 25289)
	aura(w.BloodrageAura, 29131)
	aura(w.EnrageAura, 12880)
	aura(w.GetAura("Death Wish"), 12328)
	aura(w.GetAura("Recklessness"), 1719)
	aura(w.ShieldBlockAura, 2565)
	aura(w.OverpowerAura, 1282733)
	for _, a := range w.GetAuras() {
		if a.IsActive() && strings.HasPrefix(a.Label, "Flurry Proc (") {
			aura(a, 12966)
		}
	}
	h := d.cfg.Harness
	// Incoming damage over the last five seconds.
	cut := now - 5*time.Second
	k, sum := 0, 0.0
	for _, e := range h.events {
		if e.at >= cut {
			h.events[k] = e
			k++
			sum += e.amount
		}
	}
	h.events = h.events[:k]
	window := min(5, max(now.Seconds(), 1))
	if w.MaxHealth() > 0 {
		p.IncomingPctPerS = &MeanSD{Mean: fp(100 * sum / window / w.MaxHealth()), SD: fp(0)}
	}
	msg.Player = p

	targets := sim.Encounter.TargetUnits
	// Time to die as the console estimates it: remaining health over the damage rate so far, with noise.
	ttk := func(i int) *TTK {
		if d.cfg.TTKUnknown > 0 {
			// The console reports no estimate until an enemy has been losing health for a few seconds.
			if d.hurtAt == nil {
				d.hurtAt = map[int]time.Duration{}
			}
			at, seen := d.hurtAt[i]
			if !seen {
				if h.left(i) >= h.HP[i] {
					return nil
				}
				d.hurtAt[i], at = sim.CurrentTime, sim.CurrentTime
			}
			if (sim.CurrentTime - at).Seconds() < d.cfg.TTKUnknown {
				return nil
			}
		}
		// The console's estimate: health lost over the last six seconds, blended with a prior while the
		// window is short or the enemy is barely touched.
		rate, w := h.RecentRate(i)
		prior := max(d.cfg.PriorDPS, 1)
		rate = (rate*w + prior*1.5) / (w + 1.5)
		left := h.left(i) / max(rate, 1)
		n := d.cfg.TTKNoise
		return &TTK{P10: fp(left * (1 - n)), P50: fp(left), P90: fp(left * (1 + 1.5*n))}
	}
	ti := h.index(w.CurrentTarget)
	lvl := float64(w.CurrentTarget.Level)
	class := "normal"
	if d.cfg.Elite {
		class = "elite"
	}
	threat := func(i int) (*ThreatState, *bool) {
		st, lead, tp := h.ThreatState(i)
		trend := h.LeadTrend(i)
		return &ThreatState{Status: fp(st), LeadPct: fp(lead), LeadTrendPctPer: fp(trend)}, bp(tp)
	}
	tg := TargetState{Exists: true, HealthPct: fp(100 * h.HealthFrac(ti)), Level: &lvl, Classification: &class, TTK: ttk(ti), InMelee: bp(true)}
	tg.Threat, tg.TargetsPlayer = threat(ti)
	target := w.CurrentTarget
	taura := func(a *core.Aura, gameID int32) {
		if a != nil && a.IsActive() {
			st := float64(a.GetStacks())
			tg.Auras = append(tg.Auras, AuraState{Spell: gameID, Remaining: fp(a.RemainingDuration(sim).Seconds()), Stacks: &st, AgeMs: fp(0)})
		}
	}
	if len(w.SunderArmorAuras) > 0 {
		taura(w.SunderArmorAuras.Get(target), 11597)
	}
	if w.Rend != nil {
		if dot := w.Rend.Dot(target); dot != nil {
			taura(dot.Aura, 11574)
		}
	}
	if len(w.ThunderClapAuras) > 0 {
		taura(w.ThunderClapAuras.Get(target), 11581)
	}
	if len(w.DemoralizingShoutAuras) > 0 {
		taura(w.DemoralizingShoutAuras.Get(target), 11556)
	}
	msg.Target = tg
	alive := 1.0
	for i, u := range targets {
		if i == ti || h.Dead[i] {
			continue
		}
		alive++
		e := EnemyState{ID: u.Label, HealthPct: fp(100 * h.HealthFrac(i)), TTK: ttk(i), InMelee: bp(true)}
		e.Threat, e.TargetsPlayer = threat(i)
		msg.Enemies = append(msg.Enemies, e)
	}
	msg.EnemyCount = EnemyCount{Melee: fp(alive), Engaged: fp(alive), Manual: fp(alive)}
	// Pressable actions.
	for _, ab := range m.Abilities {
		if !ab.Modelled || d.manual[ab.Name] || ab.Kind == KindStance {
			continue
		}
		ws := d.spells[ab.Name]
		if ws == nil {
			continue
		}
		cd := ws.TimeToReady(sim).Seconds()
		a := ActionState{Spell: ab.GameID, Cooldown: &cd, InRange: bp(true)}
		switch ab.NeedsWindow {
		case windowRevenge:
			a.Usable = bp(w.RevengeWindow().IsActive())
		case windowOverpower:
			a.Usable = bp(w.OverpowerAura.IsActive())
		}
		if ab.Kind == KindExecute {
			a.Usable = bp(h.ForceExecute || h.HealthFrac(ti) <= 0.2)
		}
		if ab.Stances != 0 && ab.Stances&(1<<int(stance)) == 0 {
			for s := StanceBattle; s <= StanceBerserker; s++ {
				if ab.Stances&(1<<s) != 0 && d.stanceSpell(s) != nil {
					st := float64(s)
					a.Borrow = &Borrow{Stance: &st}
					// The console reports a borrowed action as cooling down while the stance bar is.
					cd = max(cd, d.stanceSpell(s).TimeToReady(sim).Seconds())
					a.Cooldown = &cd
					break
				}
			}
			if a.Borrow == nil {
				continue
			}
		}
		msg.Actions = append(msg.Actions, a)
	}
	return msg
}

// PlannerChooser is the planner as a Chooser.
type PlannerChooser struct {
	D      *Driver
	Params Params
	Last   DecideAnswer
}

// ConsoleCadence is how often the live console asks in combat.
const ConsoleCadence = 50 * time.Millisecond

// unqueueSpell is the chooser's code for cancelling the queued swing.
const unqueueSpell int32 = -1

func (pc *PlannerChooser) Choose(msg *DecideMsg) (int32, float64) {
	if OnlyWait(pc.D.model, msg, pc.Params) {
		return 0, ConsoleCadence.Seconds() // nothing usable: the answer is wait, without the search
	}
	ans := Decide(pc.D.model, msg, pc.Params)
	pc.Last = ans
	if len(ans.Ranked) == 0 {
		return 0, 0.2
	}
	best := ans.Ranked[0]
	if best.Until != nil {
		return 0, *best.Until
	}
	if string(best.Action) == `"unqueue"` {
		return unqueueSpell, 0
	}
	var id float64
	if err := jsonNumber(best.Action, &id); err != nil {
		return 0, 0.2
	}
	return int32(id), 0
}

func jsonNumber(raw []byte, out *float64) error {
	v := 0.0
	for _, c := range raw {
		if c < '0' || c > '9' {
			return errNotNumber
		}
		v = v*10 + float64(c-'0')
	}
	*out = v
	return nil
}

type notNumber struct{}

func (notNumber) Error() string { return "not a number" }

var errNotNumber = notNumber{}

var debugPress = os.Getenv("PLANNER_DEBUG") != ""

func (d *Driver) count(a ranks.Ability) {
	if d.Casts == nil {
		d.Casts = map[ranks.Ability]int{}
	}
	d.Casts[a]++
}
