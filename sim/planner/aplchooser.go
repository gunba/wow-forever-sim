package planner

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wowsims/classic/sim/warrior/ranks"
)

// APLChooser runs a SecretProbe priority list (the console's "actions+=/Name,rank=auto,if=..." text) on the same
// decide message the planner gets, so the two are compared on identical information.
type APLChooser struct {
	D          *Driver
	Rows       []aplRow
	HomeStance int
	HomeWeapon int // 1 two-hand, 2 shield, 3 dual wield
}

type aplRow struct {
	name ranks.Ability
	cond aplNode
	src  string
}

var aplLine = regexp.MustCompile(`^actions\+=/([^,]+)(?:,rank=[a-z0-9]+)?(?:,if=(.*))?$`)

// ParseAPL reads the rows the planner can press; rows for unknown or manual abilities are skipped.
func ParseAPL(text string) ([]aplRow, error) {
	var rows []aplRow
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		m := aplLine.FindStringSubmatch(line)
		if m == nil {
			return nil, fmt.Errorf("unparsed APL line %q", line)
		}
		name := ranks.Ability(strings.ReplaceAll(m[1], " ", ""))
		var cond aplNode = aplConst(1)
		if m[2] != "" {
			p := &aplParser{toks: aplTokens(m[2])}
			n, err := p.or()
			if err != nil || p.i != len(p.toks) {
				return nil, fmt.Errorf("APL condition %q: %v", m[2], err)
			}
			cond = n
		}
		rows = append(rows, aplRow{name: name, cond: cond, src: line})
	}
	return rows, nil
}

func (ac *APLChooser) Choose(msg *DecideMsg) (int32, float64) {
	m := ac.D.model
	env := &aplEnv{msg: msg, c: ac, m: m}
	env.index()
	for _, r := range ac.Rows {
		ab := m.Get(r.name)
		if ab == nil || !ab.Modelled {
			continue
		}
		var act *ActionState
		for i := range msg.Actions {
			if msg.Actions[i].Spell == ab.GameID {
				act = &msg.Actions[i]
				break
			}
		}
		if ab.Kind != KindStance && act == nil {
			continue // not pressable (manual, untrained or no stance to borrow)
		}
		if r.cond.eval(env) == 0 {
			continue
		}
		if ab.Kind == KindStance {
			if env.stance() == ab.ToStance {
				continue
			}
			return ab.GameID, 0
		}
		if !env.castable(ab, act) {
			continue
		}
		return ab.GameID, 0
	}
	return 0, 0.1
}

type aplEnv struct {
	msg   *DecideMsg
	c     *APLChooser
	m     *Model
	auras map[string]AuraState
	taura map[string]AuraState
}

func (e *aplEnv) index() {
	e.auras, e.taura = map[string]AuraState{}, map[string]AuraState{}
	for _, a := range e.msg.Player.Auras {
		e.auras[auraKey(a.Spell)] = a
	}
	for _, a := range e.msg.Target.Auras {
		e.taura[auraKey(a.Spell)] = a
	}
}

func auraKey(id int32) string {
	if name, r := ranks.ByID(id); r != nil {
		return strings.ToLower(snake(string(name)))
	}
	switch id {
	case 2565:
		return "shield_block"
	case 1719:
		return "recklessness"
	case 12328:
		return "death_wish"
	}
	return strconv.Itoa(int(id))
}

// snake turns "BattleShout" into "battle_shout".
func snake(s string) string {
	var b strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				b.WriteByte('_')
			}
			r += 'a' - 'A'
		}
		b.WriteRune(r)
	}
	return b.String()
}

func (e *aplEnv) stance() int {
	return int(f(e.msg.Player.Stance, 0))
}

func (e *aplEnv) castable(ab *Ability, act *ActionState) bool {
	p := e.msg.Player
	if act.Usable != nil && !*act.Usable {
		return false
	}
	if act.Cooldown != nil && *act.Cooldown > 0.05 {
		return false
	}
	if f(p.Rage, 0) < ab.Cost {
		return false
	}
	if !ab.OffGCD && f(p.GCDRemaining, 0) > 0.05 {
		return false
	}
	if ab.Kind == KindQueued && p.QueuedSwing != nil && int32(*p.QueuedSwing) == ab.GameID {
		return false
	}
	if ab.NeedsShield && !e.m.Shield {
		return false
	}
	return true
}

func (e *aplEnv) value(id string) float64 {
	msg := e.msg
	p, tg := msg.Player, msg.Target
	b := func(v bool) float64 {
		if v {
			return 1
		}
		return 0
	}
	status := -1.0
	if tg.Threat != nil && tg.Threat.Status != nil {
		status = *tg.Threat.Status
	}
	pack := func() (secure, loose, contested float64) {
		add := func(t *ThreatState) {
			if t == nil || t.Status == nil {
				return
			}
			switch s := *t.Status; {
			case s >= 3:
				secure++
			case s == 2:
				contested++
			default:
				loose++
			}
		}
		add(tg.Threat)
		for _, en := range msg.Enemies {
			add(en.Threat)
		}
		return
	}
	switch id {
	case "true":
		return 1
	case "false":
		return 0
	case "in_combat", "target.can_attack", "target.exists":
		return 1
	case "moving":
		return 0
	case "rotation.pvp":
		return 0
	case "rotation.tanking":
		return b(e.c.HomeStance == StanceDefensive)
	case "rotation.dps":
		return b(e.c.HomeStance != StanceDefensive)
	case "home.stance.defensive":
		return b(e.c.HomeStance == StanceDefensive)
	case "home.stance.battle":
		return b(e.c.HomeStance == StanceBattle)
	case "home.stance.berserker":
		return b(e.c.HomeStance == StanceBerserker)
	case "home.weapons.two_hand":
		return b(e.c.HomeWeapon == 1)
	case "home.weapons.shield":
		return b(e.c.HomeWeapon == 2)
	case "home.weapons.dual_wield":
		return b(e.c.HomeWeapon == 3)
	case "player.stance.battle":
		return b(e.stance() == StanceBattle)
	case "player.stance.defensive":
		return b(e.stance() == StanceDefensive)
	case "player.stance.berserker":
		return b(e.stance() == StanceBerserker)
	case "player.rage":
		return f(p.Rage, 0)
	case "player.health_pct":
		return f(p.HealthPct, 100)
	case "player.shield_equipped":
		return b(e.m.Shield)
	case "target.health_pct":
		return f(tg.HealthPct, 100)
	case "target.count":
		return f(msg.EnemyCount.Melee, 1)
	case "target.targets_player":
		return b(tg.TargetsPlayer == nil || *tg.TargetsPlayer)
	case "target.time_to_die":
		if tg.TTK != nil && tg.TTK.P50 != nil {
			return *tg.TTK.P50
		}
		return 0
	case "target.time_to_die_known":
		return b(tg.TTK != nil && tg.TTK.P50 != nil)
	case "target.elite", "target.classification.elite":
		// Elite, rare elite and world boss; a rare is not elite.
		return b(tg.Classification != nil && (*tg.Classification == "elite" || *tg.Classification == "rareelite" || *tg.Classification == "worldboss"))
	case "target.threat.known":
		return b(status >= 0)
	case "target.threat.secure":
		return b(status == 3)
	case "target.threat.tanking":
		return b(status >= 2)
	case "target.threat.unstable":
		return b(status >= 0 && status < 3)
	case "target.threat.lost":
		return b(status >= 0 && status < 2)
	case "pack.secure_count":
		s, _, _ := pack()
		return s
	case "pack.loose":
		_, l, _ := pack()
		return l
	case "pack.contested":
		_, _, c := pack()
		return c
	}
	parts := strings.Split(id, ".")
	if len(parts) == 4 && (parts[0] == "player" || parts[0] == "target") && (parts[1] == "buff" || parts[1] == "debuff") {
		set := e.auras
		if parts[0] == "target" {
			set = e.taura
		}
		a, ok := set[parts[2]]
		switch parts[3] {
		case "up", "react", "ticking":
			return b(ok)
		case "remains", "remaining":
			if !ok {
				return 0
			}
			return f(a.Remaining, 999)
		case "stacks":
			if !ok {
				return 0
			}
			return f(a.Stacks, 1)
		}
	}
	if len(parts) == 3 && parts[0] == "spell" {
		name := ranks.Ability(strings.ReplaceAll(strings.Title(strings.ReplaceAll(parts[1], "_", " ")), " ", ""))
		ab := e.m.Get(name)
		var act *ActionState
		if ab != nil {
			for i := range msg.Actions {
				if msg.Actions[i].Spell == ab.GameID {
					act = &msg.Actions[i]
				}
			}
		}
		switch parts[2] {
		case "on_bar":
			return b(ab != nil && ab.Modelled)
		case "usable":
			if act == nil {
				return 0
			}
			return b(act.Usable == nil || *act.Usable)
		}
	}
	panic(fmt.Sprintf("APL identifier %q not supported", id))
}

// Expression tree.
type aplNode interface{ eval(e *aplEnv) float64 }

type aplConst float64
type aplVar string
type aplNot struct{ x aplNode }
type aplBin struct {
	op   string
	l, r aplNode
}

func (c aplConst) eval(*aplEnv) float64 { return float64(c) }
func (v aplVar) eval(e *aplEnv) float64 { return e.value(string(v)) }
func (n aplNot) eval(e *aplEnv) float64 {
	if n.x.eval(e) != 0 {
		return 0
	}
	return 1
}
func (n aplBin) eval(e *aplEnv) float64 {
	t := func(v bool) float64 {
		if v {
			return 1
		}
		return 0
	}
	switch n.op {
	case "&":
		return t(n.l.eval(e) != 0 && n.r.eval(e) != 0)
	case "|":
		return t(n.l.eval(e) != 0 || n.r.eval(e) != 0)
	}
	a, b := n.l.eval(e), n.r.eval(e)
	switch n.op {
	case "<":
		return t(a < b)
	case "<=":
		return t(a <= b)
	case ">":
		return t(a > b)
	case ">=":
		return t(a >= b)
	case "=", "==":
		return t(a == b)
	case "!=":
		return t(a != b)
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return 0
		}
		return a / b
	}
	return 0
}

var aplTok = regexp.MustCompile(`\s*(<=|>=|!=|==|[&|!()<>=+\-*/]|[0-9]+(?:\.[0-9]+)?|[A-Za-z_][A-Za-z0-9_.]*)`)

func aplTokens(s string) []string {
	var out []string
	for _, m := range aplTok.FindAllStringSubmatch(s, -1) {
		out = append(out, m[1])
	}
	return out
}

type aplParser struct {
	toks []string
	i    int
}

func (p *aplParser) peek() string {
	if p.i < len(p.toks) {
		return p.toks[p.i]
	}
	return ""
}

func (p *aplParser) or() (aplNode, error) {
	l, err := p.and()
	for err == nil && p.peek() == "|" {
		p.i++
		var r aplNode
		r, err = p.and()
		l = aplBin{"|", l, r}
	}
	return l, err
}

func (p *aplParser) and() (aplNode, error) {
	l, err := p.cmp()
	for err == nil && p.peek() == "&" {
		p.i++
		var r aplNode
		r, err = p.cmp()
		l = aplBin{"&", l, r}
	}
	return l, err
}

func (p *aplParser) cmp() (aplNode, error) {
	l, err := p.sum()
	if err != nil {
		return nil, err
	}
	switch op := p.peek(); op {
	case "<", "<=", ">", ">=", "=", "==", "!=":
		p.i++
		r, err := p.sum()
		return aplBin{op, l, r}, err
	}
	return l, nil
}

func (p *aplParser) sum() (aplNode, error) {
	l, err := p.unary()
	for err == nil && (p.peek() == "+" || p.peek() == "-" || p.peek() == "*" || p.peek() == "/") {
		op := p.peek()
		p.i++
		var r aplNode
		r, err = p.unary()
		l = aplBin{op, l, r}
	}
	return l, err
}

func (p *aplParser) unary() (aplNode, error) {
	switch t := p.peek(); {
	case t == "!":
		p.i++
		x, err := p.unary()
		return aplNot{x}, err
	case t == "(":
		p.i++
		x, err := p.or()
		if err != nil {
			return nil, err
		}
		if p.peek() != ")" {
			return nil, fmt.Errorf("missing )")
		}
		p.i++
		return x, nil
	case t == "":
		return nil, fmt.Errorf("unexpected end")
	case t[0] >= '0' && t[0] <= '9':
		p.i++
		v, _ := strconv.ParseFloat(t, 64)
		return aplConst(v), nil
	default:
		p.i++
		return aplVar(strings.ToLower(t)), nil
	}
}
