// Command rlbench measures the reinforcement-learning environment without a learner: environment decisions per
// second, decisions and seconds per fight, the cost of building a simulation for a new talent build, and whether
// replaying the teacher (the published rotation and talents) reproduces the published DPS.
//
//	go run -tags with_db ./tools/forever_rl/rlbench -build fire -race Human -envs 64 -threads 8 -fights 2000
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/wowsims/classic/tools/forever_rl/env"
)

func main() {
	build := flag.String("build", "fire", "build key")
	race := flag.String("race", "", "race (default: the first row of the build)")
	envs := flag.Int("envs", 64, "parallel environments")
	threads := flag.Int("threads", 8, "worker goroutines")
	fights := flag.Int("fights", 1000, "fights to play per measurement")
	learn := flag.Bool("learn-talents", false, "choose talents in each episode")
	policy := flag.String("policy", "teacher", "teacher or random")
	fightsPerBuild := flag.Int("fights-per-build", 1, "fights per chosen talent build")
	specOnly := flag.Bool("spec", false, "print the spaces and exit")
	jsonOut := flag.Bool("json", false, "print one JSON line of results")
	cpuProfile := flag.String("cpuprofile", "", "write a CPU profile of the stepping loop")
	flag.Parse()
	runtime.GOMAXPROCS(*threads)
	repo, _ := os.Getwd()

	t0 := time.Now()
	spec, err := env.NewSpec(env.Config{Repo: repo, Build: *build, Race: *race, LearnTalents: *learn, Seed: 20291021,
		Threads: *threads, FightsPerBuild: *fightsPerBuild})
	if err != nil {
		panic(err)
	}
	specSeconds := time.Since(t0).Seconds()
	if *specOnly {
		var m map[string]any
		json.Unmarshal(spec.SpecJSON(), &m)
		fmt.Printf("obs %v act %v talents %v agent actions %v\n", m["obsDim"], m["actDim"], m["talents"], m["agentActions"])
		for _, a := range m["actionNames"].([]any)[spec.T:] {
			fmt.Println("  ", a)
		}
		for _, line := range spec.TeacherMap() {
			fmt.Println("  teacher", line)
		}
		tp := time.Now()
		fmt.Println("plain APL DPS (RunIteration, 500 fights):", spec.PlainDPS(500), "ms/fight", time.Since(tp).Seconds()*2)
		tf := time.Now()
		fmt.Println("full simulation DPS (500 iterations):", spec.FullDPS(500), "ms/fight", time.Since(tf).Seconds()*2)
		return
	}

	var mem0 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&mem0)
	t1 := time.Now()
	v := env.NewVecEnv(spec, *envs)
	n := *envs
	b := &env.Buffers{Obs: make([]float32, n*spec.ObsDim), Mask: make([]uint8, n*spec.ActDim), Teacher: make([]int32, n),
		Reward: make([]float32, n), Done: make([]uint8, n), Info: make([]float32, n*4)}
	v.Reset(b)
	firstSeconds := time.Since(t1).Seconds()
	var mem1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&mem1)

	rng := rand.New(rand.NewSource(1))
	actions := make([]int32, n)
	steps, done := 0, 0
	dpsSum, dpsSq, decSum, secSum, builds := 0.0, 0.0, 0.0, 0.0, 0
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	t2 := time.Now()
	for done < *fights {
		for i := 0; i < n; i++ {
			if *policy == "teacher" {
				actions[i] = b.Teacher[i]
				continue
			}
			mask := b.Mask[i*spec.ActDim : (i+1)*spec.ActDim]
			legal := 0
			for _, m := range mask {
				legal += int(m)
			}
			pick := rng.Intn(legal)
			for a, m := range mask {
				if m == 1 {
					if pick == 0 {
						actions[i] = int32(a)
						break
					}
					pick--
				}
			}
		}
		v.Step(actions, b)
		steps += n
		for i := 0; i < n; i++ {
			if b.Done[i] == 1 {
				done++
				dpsSum += float64(b.Info[i*4])
				dpsSq += float64(b.Info[i*4]) * float64(b.Info[i*4])
				decSum += float64(b.Info[i*4+1])
				builds += int(b.Info[i*4+2])
				secSum += float64(b.Info[i*4+3])
			}
		}
	}
	elapsed := time.Since(t2).Seconds()
	built, buildSeconds := 0, 0.0
	for _, e := range v.Envs {
		built += e.SimsBuilt
		buildSeconds += e.SimSeconds
	}
	out := map[string]any{
		"build": *build, "race": spec.Cfg.Race, "policy": *policy, "learnTalents": *learn, "envs": n, "threads": *threads,
		"obsDim": spec.ObsDim, "actDim": spec.ActDim, "agentActions": spec.K, "specSeconds": specSeconds,
		"firstDecisionSeconds": firstSeconds, "simMemoryMBPerEnv": float64(mem1.HeapAlloc-mem0.HeapAlloc) / float64(n) / 1e6,
		"fights": done, "decisionsPerSecond": float64(steps) / elapsed, "fightsPerSecond": float64(done) / elapsed,
		"meanDPS": dpsSum / float64(done), "stdDPS": math.Sqrt(math.Max(0, dpsSq/float64(done)-(dpsSum/float64(done))*(dpsSum/float64(done)))), "decisionsPerFight": decSum / float64(done), "fightSeconds": secSum / float64(done),
		"newBuilds": builds, "newBuildsPerSecond": float64(builds) / elapsed, "simsBuilt": built,
		"msPerSimBuild": 1000 * buildSeconds / float64(max(built, 1)),
	}
	if *jsonOut {
		j, _ := json.Marshal(out)
		fmt.Println(string(j))
	} else {
		for _, k := range []string{"build", "race", "policy", "learnTalents", "envs", "threads", "obsDim", "actDim", "agentActions", "specSeconds",
			"firstDecisionSeconds", "simMemoryMBPerEnv", "fights", "decisionsPerSecond", "fightsPerSecond", "meanDPS", "decisionsPerFight",
			"fightSeconds", "newBuilds", "newBuildsPerSecond"} {
			fmt.Printf("%-22s %v\n", k, out[k])
		}
	}
	v.Close()
}
