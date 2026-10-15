// forever_planner answers the SecretProbe console's planner messages (PLANNER.md): one JSON object per line on stdin,
// one answer per line on stdout. A "character" message builds the character's ability model; each "decide" message
// is answered from the model with the matching generation.
//
//	go build -tags with_db -o dist/forever_planner ./tools/forever_planner
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/planner"
)

func main() {
	params := flag.String("params", "", "optional JSON file of planner parameters (planner.Params)")
	inspect := flag.String("inspect", "", "print the ability model built from a character message file and exit")
	bench := flag.String("bench", "", "run the scenario benchmark described by this JSON file and exit")
	flag.Parse()
	sim.RegisterAll()
	if *bench != "" {
		runBench(*bench)
		return
	}
	if *inspect != "" {
		data, err := os.ReadFile(*inspect)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		var msg planner.CharacterMsg
		if err := json.Unmarshal(data, &msg); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		model, ans := planner.BuildCharacter(&msg)
		if model == nil {
			fmt.Fprintln(os.Stderr, ans.Error)
			os.Exit(1)
		}
		fmt.Print(model.Describe())
		return
	}
	pr := planner.DefaultParams
	if *params != "" {
		data, err := os.ReadFile(*params)
		if err == nil {
			err = json.Unmarshal(data, &pr)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "params:", err)
			os.Exit(2)
		}
	}
	debug.SetGCPercent(400)
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 1<<20), 1<<24)
	out := bufio.NewWriter(os.Stdout)
	enc := json.NewEncoder(out)
	models := map[int64]*planner.Model{}
	var latest int64 = -1
	for in.Scan() {
		line := in.Bytes()
		if len(line) == 0 {
			continue
		}
		var env planner.Envelope
		if err := json.Unmarshal(line, &env); err != nil {
			enc.Encode(map[string]any{"type": "error", "error": "bad JSON: " + err.Error()})
			out.Flush()
			continue
		}
		switch env.Type {
		case "character":
			var msg planner.CharacterMsg
			if err := json.Unmarshal(line, &msg); err != nil {
				enc.Encode(planner.CharacterAnswer{Type: "character", Error: err.Error(), Abilities: []planner.AbilityInfo{}, Warnings: []string{}})
				break
			}
			model, ans := planner.BuildCharacter(&msg)
			if model != nil {
				models[msg.Generation] = model
				latest = msg.Generation
				for g := range models { // keep the last few generations
					if g < msg.Generation-4 {
						delete(models, g)
					}
				}
			}
			enc.Encode(ans)
		case "decide":
			var msg planner.DecideMsg
			if err := json.Unmarshal(line, &msg); err != nil {
				enc.Encode(planner.DecideAnswer{Type: "decide", Error: err.Error(), Warnings: []string{}})
				break
			}
			model := models[msg.Generation]
			warn := []string{}
			if model == nil && latest >= 0 {
				model = models[latest]
				warn = append(warn, fmt.Sprintf("no model for generation %d; using %d", msg.Generation, latest))
			}
			if model == nil {
				enc.Encode(planner.DecideAnswer{Type: "decide", ID: msg.ID, Error: "no character message yet", Warnings: warn})
				break
			}
			ans := decideSafe(model, &msg, pr)
			ans.Warnings = append(warn, ans.Warnings...)
			enc.Encode(ans)
		default:
			enc.Encode(map[string]any{"type": "error", "error": "unknown message type " + env.Type})
		}
		out.Flush()
	}
}

func decideSafe(model *planner.Model, msg *planner.DecideMsg, pr planner.Params) (ans planner.DecideAnswer) {
	defer func() {
		if r := recover(); r != nil {
			ans = planner.DecideAnswer{Type: "decide", ID: msg.ID, Error: fmt.Sprintf("planner failure: %v", r), Warnings: []string{}}
		}
	}()
	return planner.Decide(model, msg, pr)
}
