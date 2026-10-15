package planner

import (
	"fmt"
	"strings"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/warrior"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

// Calibration compares the forward model's expected outcome of each ability with the engine's: the planner presses
// one ability whenever it can, with rage topped up and its window forced open, against an enemy that does not die
// or attack, and the engine's average damage and threat per use are set against the model's.

type CalibrationRow struct {
	Name                      string
	Uses                      int
	EngineDamage, ModelDamage float64
	EngineThreat, ModelThreat float64
}

type calibChooser struct {
	d   *Driver
	ab  *Ability
	fix func(sim *core.Simulation, w *warrior.Warrior)
}

func (cc *calibChooser) Choose(msg *DecideMsg) (int32, float64) {
	d := cc.d
	w := d.w
	sim := d.Sim
	if r := w.CurrentRage(); r < 100 {
		w.AddRage(sim, 100-r, calibMetrics(w))
	}
	if cc.fix != nil {
		cc.fix(sim, w)
	}
	return cc.ab.GameID, 0
}

var calibRage = map[*warrior.Warrior]*core.ResourceMetrics{}

func calibMetrics(w *warrior.Warrior) *core.ResourceMetrics {
	if m, ok := calibRage[w]; ok {
		return m
	}
	m := w.NewRageMetrics(core.ActionID{OtherID: 1})
	calibRage[w] = m
	return m
}

// Calibrate runs every modelled damage or threat ability of the scenario's character.
func Calibrate(sc Scenario, model *Model, iters int) []CalibrationRow {
	var rows []CalibrationRow
	sc.Enemies, sc.HPMult, sc.DmgMult, sc.Partner, sc.Cap = 1, 1e4, 0, Partner{}, 60
	for _, ab := range model.Abilities {
		if !ab.Modelled {
			continue
		}
		switch ab.Kind {
		case KindStance, KindBuff, KindRage, KindBlock:
			continue
		}
		s := sc
		s.HomeStance = sc.HomeStance
		if ab.Stances != 0 && ab.Stances&(1<<s.HomeStance) == 0 {
			for st := StanceBattle; st <= StanceBerserker; st++ {
				if ab.Stances&(1<<st) != 0 {
					s.HomeStance = st
					break
				}
			}
		}
		player := cloneStance(sc.Player, s.HomeStance)
		s.Player = player
		var cc *calibChooser
		b, err := NewBench(s, model, func(d *Driver) Chooser {
			cc = &calibChooser{d: d, ab: ab}
			return cc
		})
		if err != nil {
			continue
		}
		w := b.driver.w
		switch ab.NeedsWindow {
		case windowRevenge:
			cc.fix = func(sim *core.Simulation, w *warrior.Warrior) { w.RevengeWindow().Activate(sim) }
		case windowOverpower:
			cc.fix = func(sim *core.Simulation, w *warrior.Warrior) { w.OverpowerAura.Activate(sim) }
		}
		if ab.Kind == KindExecute {
			b.harness.ForceExecute = true
		}
		var dmg, threat float64
		uses := 0
		target := ab.SpellID
		b.harness.OnHit = func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult, periodic bool) {
			if spell.ActionID.SpellID != target {
				return
			}
			dmg += result.Damage
			threat += result.Threat
			if !periodic && ab.Kind != KindDot || periodic && ab.Kind == KindDot {
				uses++
			}
		}
		presses := 0
		for i := 0; i < iters; i++ {
			o := b.Run(int64(1000 + i))
			presses += o.Casts[string(ab.Name)]
		}
		b.Release()
		_ = w
		n := uses
		if ab.Kind == KindDebuff {
			n = presses
		}
		if n == 0 {
			rows = append(rows, CalibrationRow{Name: string(ab.Name)})
			continue
		}
		md, mt := expectedOutcome(model, s, ab)
		rows = append(rows, CalibrationRow{Name: string(ab.Name), Uses: n, EngineDamage: dmg / float64(n), EngineThreat: threat / float64(n),
			ModelDamage: md, ModelThreat: mt})
	}
	// White swings: nothing pressed.
	b, err := NewBench(sc, model, func(d *Driver) Chooser { return waitChooser{} })
	if err == nil {
		var mhD, ohD float64
		var mhN, ohN int
		b.harness.OnHit = func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult, periodic bool) {
			switch {
			case spell.ProcMask.Matches(core.ProcMaskMeleeMHAuto):
				mhD += result.Damage
				mhN++
			case spell.ProcMask.Matches(core.ProcMaskMeleeOHAuto):
				ohD += result.Damage
				ohN++
			}
		}
		for i := 0; i < iters; i++ {
			b.Run(int64(2000 + i))
		}
		b.Release()
		stance := float64(sc.HomeStance)
		lvl := float64(sc.EnemyLevel)
		msg := &DecideMsg{Intent: Intent{Mode: "kill"}, Player: PlayerState{Stance: &stance}, Target: TargetState{Exists: true, Level: &lvl, TargetsPlayer: bp(sc.Front)}}
		c, s, _ := NewCtx(model, msg, DefaultParams)
		md := swingDamage(model.MH, c.ap(s), false) * model.MH.DamageMult * c.physMult(s) * c.mhWhite.mult
		if mhN > 0 {
			rows = append(rows, CalibrationRow{Name: "white main hand", Uses: mhN, EngineDamage: mhD / float64(mhN), ModelDamage: md})
		}
		if ohN > 0 {
			od := swingDamage(model.OH, c.ap(s), false) * model.OH.DamageMult * c.physMult(s) * c.ohWhite.mult
			rows = append(rows, CalibrationRow{Name: "white off hand", Uses: ohN, EngineDamage: ohD / float64(ohN), ModelDamage: od})
		}
	}
	return rows
}

type waitChooser struct{}

func (waitChooser) Choose(*DecideMsg) (int32, float64) { return 0, 1 }

// expectedOutcome is the forward model's expected damage and raw threat (before intent weights) per use.
func expectedOutcome(m *Model, sc Scenario, ab *Ability) (float64, float64) {
	stance := float64(sc.HomeStance)
	lvl := float64(sc.EnemyLevel)
	msg := &DecideMsg{Intent: Intent{Mode: "kill"}, Player: PlayerState{Stance: &stance, Rage: fp(100)}, Target: TargetState{Exists: true, Level: &lvl, TargetsPlayer: bp(sc.Front)}}
	c, s, _ := NewCtx(m, msg, DefaultParams)
	c.WThreatTarget, c.WThreatOther = 1, 1
	s.deathAt, s.execAt = 1e9, 0
	ap := c.ap(s)
	st := c.stanceThreat(s)
	o := c.table(ab)
	switch ab.Kind {
	case KindStrike, KindAoE, KindQueued:
		if ab.Kind == KindQueued {
			o = &c.special
		}
		mult := c.physMult(s)
		d := c.abilityRaw(ab, ap) * mult * (o.land - o.crit + o.crit*ab.CritMult)
		th := (d*ab.ThreatMult + (ab.FlatThreat+ab.ThreatAP*ap)*o.land) * st
		if ab.Kind == KindQueued {
			th = (d*ab.ThreatMult + ab.FlatThreat*o.land) * st
		}
		return d, th
	case KindExecute:
		raw := (ab.Flat + ab.PerRage*(100-ab.Cost)) * ab.DamageMult
		d := raw * c.physMult(s) * (o.land - o.crit + o.crit*ab.CritMult)
		return d, d * ab.ThreatMult * st
	case KindDot:
		d := (ab.TickDamage + ab.TickAPCoef*ap) * ab.DamageMult * c.armorFreeMult(s) * c.tickCrit() // per tick
		return d, d * st
	case KindDebuff:
		if ab.Name == ranks.DemoralizingShout {
			return 0, ab.FlatThreat * st
		}
		return 0, (ab.FlatThreat + ab.ThreatAP*ap) * st * c.special.land
	}
	return 0, 0
}

// FormatCalibration prints the rows.
func FormatCalibration(rows []CalibrationRow) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%-18s %6s %10s %10s %7s %10s %10s %7s\n", "ability", "uses", "eng dmg", "model dmg", "ratio", "eng thr", "model thr", "ratio")
	for _, r := range rows {
		rd, rt := 0.0, 0.0
		if r.EngineDamage > 0 {
			rd = r.ModelDamage / r.EngineDamage
		}
		if r.EngineThreat > 0 {
			rt = r.ModelThreat / r.EngineThreat
		}
		fmt.Fprintf(&b, "%-18s %6d %10.1f %10.1f %7.3f %10.1f %10.1f %7.3f\n", r.Name, r.Uses, r.EngineDamage, r.ModelDamage, rd, r.EngineThreat, r.ModelThreat, rt)
	}
	return b.String()
}
