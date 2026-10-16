package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/planner"
	googleProto "google.golang.org/protobuf/proto"
)

// benchFile describes a benchmark: characters, priority lists and scenarios, each run with the planner and with
// the priority list on the same seeds.
type benchFile struct {
	Characters map[string]string `json:"characters"` // name -> character message file
	APLs       map[string]string `json:"apls"`       // name -> SecretProbe profile JSON (its "apl") or .apl text
	Iterations int               `json:"iterations"`
	Seed       int64             `json:"seed"`
	Samples    int               `json:"samples"` // planner samples per decision (fixed, for reproducibility)
	Params     json.RawMessage   `json:"params"`  // overrides on planner.DefaultParams
	Scenarios  []benchScenario   `json:"scenarios"`
	Out        string            `json:"out"`
	Casts      bool              `json:"casts"`
	Calibrate  bool              `json:"calibrate"`
	Trace      string            `json:"trace"`   // scenario name: print one iteration's decisions for both choosers
	ACache     string            `json:"a_cache"` // directory caching the priority-list side by its inputs and this binary
}

type benchScenario struct {
	Name       string          `json:"name"`
	Character  string          `json:"character"`
	APL        string          `json:"apl"`
	APLB       string          `json:"apl_b"` // when set, this priority list replaces the planner side (A/B of two lists)
	Level      int32           `json:"level"` // 0: the character's
	Talents    string          `json:"talents"`
	Enemies    int             `json:"enemies"`
	EnemyLevel int32           `json:"enemy_level"` // 0: character level + 1
	Elite      bool            `json:"elite"`
	HPMult     float64         `json:"hp_mult"`
	DmgMult    float64         `json:"dmg_mult"`
	Partner    planner.Partner `json:"partner"`
	Mode       string          `json:"mode"`
	HomeStance int             `json:"home_stance"`
	HomeWeapon int             `json:"home_weapon"`
	TTKNoise   float64         `json:"ttk_noise"`
	Manual     []string        `json:"manual"` // extra abilities kept manual (ablation)
	Cap        float64         `json:"cap"`
	TTKUnknown float64         `json:"ttk_unknown"`
}

type sideStats struct {
	KillTime, Useful, Taken, LossTime, Swaps, Decisions, MsPerDecision, TPS, DPS float64
	Casts                                                                        map[string]float64
}

type benchResult struct {
	Name    string                `json:"name"`
	Planner sideStats             `json:"planner"`
	APL     sideStats             `json:"apl"`
	Diff    map[string][3]float64 `json:"diff"` // metric -> mean, CI low, CI high of (planner - APL)
	PctKill [3]float64            `json:"pct_kill_time"`
	N       int                   `json:"n"`
}

func loadAPL(path string) string {
	data, err := os.ReadFile(path)
	must(err)
	if strings.HasSuffix(path, ".json") {
		var prof struct {
			APL string `json:"apl"`
		}
		must(json.Unmarshal(data, &prof))
		return prof.APL
	}
	return string(data)
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runBench(path string) {
	data, err := os.ReadFile(path)
	must(err)
	var bf benchFile
	must(json.Unmarshal(data, &bf))
	if bf.Iterations <= 0 {
		bf.Iterations = 200
	}
	if bf.Samples <= 0 {
		bf.Samples = 24
	}
	params := planner.DefaultParams
	if len(bf.Params) > 0 {
		must(json.Unmarshal(bf.Params, &params))
	}
	params.FixedSamples = bf.Samples
	players := map[string]*proto.Player{}
	for name, file := range bf.Characters {
		raw, err := os.ReadFile(file)
		must(err)
		var msg planner.CharacterMsg
		must(json.Unmarshal(raw, &msg))
		p, warn, err := planner.CharacterPlayer(&msg)
		must(err)
		if len(warn) > 0 {
			fmt.Fprintf(os.Stderr, "%s: %v\n", name, warn)
		}
		players[name] = p
	}
	apls := map[string]string{}
	for name, file := range bf.APLs {
		apls[name] = loadAPL(file)
	}
	workers := runtime.GOMAXPROCS(0)
	type job struct {
		sc      int
		planner bool
		from    int
		to      int
	}
	type part struct {
		sc      int
		planner bool
		from    int
		outs    []planner.Outcome
	}
	// The priority-list side of an A/B depends only on its own inputs, so a sweep of many B lists runs it once.
	aKeys := make([]string, len(bf.Scenarios))
	cachedA := make([][]planner.Outcome, len(bf.Scenarios))
	if bf.ACache != "" {
		must(os.MkdirAll(bf.ACache, 0o755))
		exe, err := os.Executable()
		must(err)
		bin, err := os.ReadFile(exe)
		must(err)
		binSum := sha256.Sum256(bin)
		for i, s := range bf.Scenarios {
			key := s
			key.Name, key.APLB = "", ""
			char, err := os.ReadFile(bf.Characters[s.Character])
			must(err)
			raw, err := json.Marshal([]any{key, apls[s.APL], string(char), bf.Iterations, bf.Seed, hex.EncodeToString(binSum[:])})
			must(err)
			sum := sha256.Sum256(raw)
			aKeys[i] = filepath.Join(bf.ACache, hex.EncodeToString(sum[:12])+".json")
			if data, err := os.ReadFile(aKeys[i]); err == nil {
				var outs []planner.Outcome
				if json.Unmarshal(data, &outs) == nil && len(outs) == bf.Iterations {
					cachedA[i] = outs
				}
			}
		}
	}
	var jobs []job
	chunk := max(bf.Iterations/max(1, workers/max(1, 2*len(bf.Scenarios))), 10)
	for i := range bf.Scenarios {
		for _, pl := range []bool{true, false} {
			if !pl && cachedA[i] != nil {
				continue
			}
			for from := 0; from < bf.Iterations; from += chunk {
				jobs = append(jobs, job{i, pl, from, min(from+chunk, bf.Iterations)})
			}
		}
	}
	scen := make([]planner.Scenario, len(bf.Scenarios))
	models := make([]*planner.Model, len(bf.Scenarios))
	for i, s := range bf.Scenarios {
		p, ok := players[s.Character]
		if !ok {
			must(fmt.Errorf("scenario %s: unknown character %s", s.Name, s.Character))
		}
		p = cloneWithStance(p, s)
		m, err := planner.ModelFor(p)
		must(err)
		lvl := p.Level
		if lvl == 0 {
			lvl = 60
		}
		el := s.EnemyLevel
		if el == 0 {
			el = lvl + 1
		}
		scen[i] = planner.Scenario{Name: s.Name, Player: p, Enemies: s.Enemies, EnemyLevel: el, Elite: s.Elite, HPMult: def(s.HPMult, 1),
			DmgMult: def(s.DmgMult, 1), Partner: s.Partner, HomeStance: s.HomeStance, HomeWeapon: s.HomeWeapon, TTKNoise: s.TTKNoise, Cap: def(s.Cap, 180), TTKUnknown: s.TTKUnknown,
			Intent: planner.Intent{Mode: s.Mode, Source: "override"}, Front: s.Partner.Role != "tank", Manual: s.Manual}
		if s.Partner.Role == "tank" {
			yes := true
			scen[i].Intent.Context = &planner.IntentContext{OtherTankPresent: &yes}
		}
		models[i] = m
	}
	if bf.Calibrate {
		for i, s := range bf.Scenarios {
			fmt.Printf("== %s (stance %d)\n%s\n", s.Name, s.HomeStance, planner.FormatCalibration(planner.Calibrate(scen[i], models[i], 40)))
		}
		return
	}
	if bf.Trace != "" {
		for i, s := range bf.Scenarios {
			if s.Name == bf.Trace {
				trace(scen[i], models[i], s, params, apls, bf.Seed)
			}
		}
		return
	}
	results := make(chan part, len(jobs))
	var wg sync.WaitGroup
	queue := make(chan job)
	start := time.Now()
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range queue {
				s := bf.Scenarios[j.sc]
				var b *planner.Bench
				var err error
				if j.planner && s.APLB == "" {
					b, err = planner.NewBench(scen[j.sc], models[j.sc], func(d *planner.Driver) planner.Chooser {
						return &planner.PlannerChooser{D: d, Params: params}
					})
				} else {
					name := s.APL
					if j.planner {
						name = s.APLB
					}
					text, ok := apls[name]
					if !ok {
						must(fmt.Errorf("scenario %s: unknown APL %s", s.Name, name))
					}
					rows, perr := planner.ParseAPL(text)
					must(perr)
					b, err = planner.NewBench(scen[j.sc], models[j.sc], func(d *planner.Driver) planner.Chooser {
						return &planner.APLChooser{D: d, Rows: rows, HomeStance: s.HomeStance, HomeWeapon: s.HomeWeapon}
					})
				}
				must(err)
				p := part{sc: j.sc, planner: j.planner, from: j.from}
				for it := j.from; it < j.to; it++ {
					p.outs = append(p.outs, b.Run(bf.Seed+int64(it)))
				}
				b.Release()
				results <- p
			}
		}()
	}
	go func() {
		for _, j := range jobs {
			queue <- j
		}
		close(queue)
	}()
	wg.Wait()
	close(results)
	pl := make([][]planner.Outcome, len(scen))
	ap := make([][]planner.Outcome, len(scen))
	for i := range scen {
		pl[i] = make([]planner.Outcome, bf.Iterations)
		ap[i] = make([]planner.Outcome, bf.Iterations)
	}
	for p := range results {
		dst := ap[p.sc]
		if p.planner {
			dst = pl[p.sc]
		}
		copy(dst[p.from:], p.outs)
	}
	for i := range scen {
		if cachedA[i] != nil {
			ap[i] = cachedA[i]
		} else if aKeys[i] != "" {
			data, err := json.Marshal(ap[i])
			must(err)
			must(os.WriteFile(aKeys[i], data, 0o644))
		}
	}
	var out []benchResult
	for i, s := range bf.Scenarios {
		r := benchResult{Name: s.Name, N: bf.Iterations, Diff: map[string][3]float64{}}
		r.Planner, r.APL = summarize(pl[i]), summarize(ap[i])
		metrics := map[string]func(o planner.Outcome) float64{
			"kill_time":  func(o planner.Outcome) float64 { return o.KillTime },
			"useful":     func(o planner.Outcome) float64 { return o.Useful },
			"taken":      func(o planner.Outcome) float64 { return o.Taken },
			"loss_time":  func(o planner.Outcome) float64 { return o.LossTime },
			"loss_focus": func(o planner.Outcome) float64 { return o.LossFocus },
		}
		for name, fn := range metrics {
			r.Diff[name] = pairedCI(pl[i], ap[i], fn)
		}
		d := r.Diff["kill_time"]
		base := r.APL.KillTime
		r.PctKill = [3]float64{100 * d[0] / base, 100 * d[1] / base, 100 * d[2] / base}
		out = append(out, r)
	}
	sort.SliceStable(out, func(a, b int) bool { return out[a].Name < out[b].Name })
	fmt.Printf("%-28s %8s %8s %18s %8s %8s %16s %8s %8s %6s %13s %13s\n", "scenario", "kill P", "kill A", "kill% (95% CI)", "loss P", "loss A", "loss diff CI", "taken P", "taken A", "ms/d", "DPS P/A", "TPS P/A")
	for _, r := range out {
		lc := r.Diff["loss_time"]
		fmt.Printf("%-28s %8.2f %8.2f %+6.2f [%+5.2f,%+5.2f] %8.2f %8.2f [%+6.2f,%+6.2f] %8.0f %8.0f %6.3f %6.1f/%6.1f %6.1f/%6.1f\n", r.Name, r.Planner.KillTime, r.APL.KillTime,
			r.PctKill[0], r.PctKill[1], r.PctKill[2], r.Planner.LossTime, r.APL.LossTime, lc[1], lc[2], r.Planner.Taken, r.APL.Taken, r.Planner.MsPerDecision,
			r.Planner.DPS, r.APL.DPS, r.Planner.TPS, r.APL.TPS)
	}
	if bf.Casts {
		for _, r := range out {
			fmt.Printf("\n%s casts per fight (planner / APL):", r.Name)
			keys := map[string]bool{}
			for k := range r.Planner.Casts {
				keys[k] = true
			}
			for k := range r.APL.Casts {
				keys[k] = true
			}
			var ks []string
			for k := range keys {
				ks = append(ks, k)
			}
			sort.Strings(ks)
			for _, k := range ks {
				fmt.Printf(" %s %.1f/%.1f", k, r.Planner.Casts[k], r.APL.Casts[k])
			}
			fmt.Println()
		}
	}
	fmt.Fprintf(os.Stderr, "bench: %d scenarios x %d iterations in %s\n", len(out), bf.Iterations, time.Since(start).Round(time.Second))
	if bf.Out != "" {
		data, _ := json.MarshalIndent(out, "", " ")
		must(os.WriteFile(bf.Out, data, 0o644))
	}
}

func def(v, d float64) float64 {
	if v == 0 {
		return d
	}
	return v
}

func cloneWithStance(p *proto.Player, s benchScenario) *proto.Player {
	q := googleProto.Clone(p).(*proto.Player)
	if s.Level > 0 {
		q.Level = s.Level
	}
	if s.Talents != "" {
		q.TalentsString = s.Talents
	}
	st := proto.WarriorStance(s.HomeStance)
	if st == 0 {
		st = proto.WarriorStance_WarriorStanceBattle
	}
	q.Spec = &proto.Player_TankWarrior{TankWarrior: &proto.TankWarrior{Options: &proto.TankWarrior_Options{Stance: st}}}
	return q
}

func summarize(os []planner.Outcome) sideStats {
	s := sideStats{Casts: map[string]float64{}}
	var ns float64
	for _, o := range os {
		for k, v := range o.Casts {
			s.Casts[k] += float64(v) / float64(len(os))
		}
		for k, v := range o.DamageBy {
			s.Casts["dmg:"+k] += v / float64(len(os))
		}
		s.KillTime += o.KillTime
		s.TPS += o.Threat / max(o.KillTime, 1)
		s.DPS += o.Useful / max(o.KillTime, 1)
		s.Useful += o.Useful
		s.Taken += o.Taken
		s.LossTime += o.LossTime
		s.Swaps += float64(o.Swaps)
		s.Decisions += float64(o.Decisions)
		ns += float64(o.ChooseNanos)
	}
	n := float64(len(os))
	s.KillTime /= n
	s.TPS /= n
	s.DPS /= n
	s.Useful /= n
	s.Taken /= n
	s.LossTime /= n
	s.Swaps /= n
	if s.Decisions > 0 {
		s.MsPerDecision = ns / 1e6 / s.Decisions
	}
	s.Decisions /= n
	return s
}

func pairedCI(a, b []planner.Outcome, fn func(planner.Outcome) float64) [3]float64 {
	n := float64(len(a))
	var sum, sq float64
	for i := range a {
		d := fn(a[i]) - fn(b[i])
		sum += d
		sq += d * d
	}
	mean := sum / n
	sd := math.Sqrt(max(0, sq/n-mean*mean) * n / max(n-1, 1))
	h := 1.96 * sd / math.Sqrt(n)
	return [3]float64{mean, mean - h, mean + h}
}

func trace(sc planner.Scenario, m *planner.Model, s benchScenario, params planner.Params, apls map[string]string, seed int64) {
	names := map[int32]string{}
	for _, ab := range m.Abilities {
		names[ab.GameID] = string(ab.Name)
	}
	for _, usePlanner := range []bool{true, false} {
		var pc *planner.PlannerChooser
		b, err := planner.NewBench(sc, m, func(d *planner.Driver) planner.Chooser {
			if usePlanner {
				pc = &planner.PlannerChooser{D: d, Params: params}
				return pc
			}
			rows, err := planner.ParseAPL(apls[s.APL])
			must(err)
			return &planner.APLChooser{D: d, Rows: rows, HomeStance: s.HomeStance, HomeWeapon: s.HomeWeapon}
		})
		must(err)
		label := map[bool]string{true: "PLANNER", false: "APL"}[usePlanner]
		b.Driver().Trace = func(sim *core.Simulation, msg *planner.DecideMsg, spell int32, until float64) {
			if spell == 0 && !usePlanner {
				return
			}
			act := names[spell]
			if spell == 0 {
				act = fmt.Sprintf("wait %.2f", until)
			}
			extra := ""
			if pc != nil {
				for k, r := range pc.Last.Ranked {
					if k >= 4 {
						break
					}
					var id int32
					json.Unmarshal(r.Action, &id)
					n := names[id]
					if n == "" {
						n = "wait"
					}
					extra += fmt.Sprintf(" %s=%.0f", n, r.Value)
				}
			}
			fmt.Printf("%s %6.2f rage %5.1f gcd %.2f mh %.2f hp %5.1f%% stance %.0f q %v -> %-16s%s\n", label, msg.T, *msg.Player.Rage,
				*msg.Player.GCDRemaining, *msg.Player.Swing["main"].Remaining, *msg.Target.HealthPct, *msg.Player.Stance, msg.Player.QueuedSwing != nil, act, extra)
		}
		o := b.Run(seed)
		fmt.Printf("%s outcome: %+v\n\n", label, o)
		b.Release()
	}
}
