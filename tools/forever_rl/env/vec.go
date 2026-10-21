package env

import (
	"runtime"
	"sync"
)

// VecEnv steps many environments at once. Each worker goroutine creates and always resumes its own slice of
// environments: a coroutine resumed from another goroutine takes the runtime's slow path.
type VecEnv struct {
	Spec    *Spec
	Envs    []*Env
	workers []*worker
	wg      sync.WaitGroup
}

type command struct {
	kind    int // 0 reset, 1 step, 2 close, 3 restart with seeds
	actions []int32
	seeds   []int64
	buf     *Buffers
}

type worker struct {
	lo, hi int
	cmds   chan command
}

func NewVecEnv(spec *Spec, n int) *VecEnv {
	w := spec.Cfg.Threads
	if w <= 0 {
		w = runtime.GOMAXPROCS(0)
	}
	if w > n {
		w = n
	}
	v := &VecEnv{Spec: spec, Envs: make([]*Env, n)}
	var ready sync.WaitGroup
	for k := 0; k < w; k++ {
		wk := &worker{lo: k * n / w, hi: (k + 1) * n / w, cmds: make(chan command)}
		v.workers = append(v.workers, wk)
		ready.Add(1)
		go v.loop(wk, &ready)
	}
	ready.Wait()
	return v
}

func (v *VecEnv) loop(wk *worker, ready *sync.WaitGroup) {
	for i := wk.lo; i < wk.hi; i++ {
		v.Envs[i] = newEnv(v.Spec, i)
	}
	ready.Done()
	for cmd := range wk.cmds {
		for i := wk.lo; i < wk.hi; i++ {
			e := v.Envs[i]
			switch cmd.kind {
			case 0:
				e.Advance()
				v.copyOut(i, e, cmd.buf)
			case 1:
				e.Action = cmd.actions[i]
				if cmd.seeds != nil {
					e.Reseed = cmd.seeds[i]
				}
				e.Advance()
				v.copyOut(i, e, cmd.buf)
			case 2:
				e.Close()
			case 3:
				e.Restart(cmd.seeds[i])
				e.Advance()
				v.copyOut(i, e, cmd.buf)
			}
		}
		v.wg.Done()
	}
}

// Buffers are flat: obs n*ObsDim, mask n*ActDim, teacher/reward/done n, info n*4.
type Buffers struct {
	Obs     []float32
	Mask    []uint8
	Teacher []int32
	Reward  []float32
	Done    []uint8
	Info    []float32
}

func (v *VecEnv) broadcast(cmd command) {
	v.wg.Add(len(v.workers))
	for _, wk := range v.workers {
		wk.cmds <- cmd
	}
	v.wg.Wait()
}

func (v *VecEnv) copyOut(i int, e *Env, b *Buffers) {
	d, a := v.Spec.ObsDim, v.Spec.ActDim
	copy(b.Obs[i*d:(i+1)*d], e.Obs)
	copy(b.Mask[i*a:(i+1)*a], e.Mask)
	b.Teacher[i] = e.Teacher
	if b.Reward != nil {
		b.Reward[i] = float32(e.Reward)
		b.Done[i] = 0
		if e.Done {
			b.Done[i] = 1
		}
		copy(b.Info[i*4:(i+1)*4], e.Info[:])
	}
}

// Reset runs every environment to its first decision.
func (v *VecEnv) Reset(b *Buffers) { v.broadcast(command{kind: 0, buf: b}) }

// Step applies one action per environment and runs each to its next decision.
func (v *VecEnv) Step(actions []int32, b *Buffers) {
	v.broadcast(command{kind: 1, actions: actions, buf: b})
}

// StepReseed is Step with a per-environment reseed applied right after the action (0: none).
func (v *VecEnv) StepReseed(actions []int32, reseeds []int64, b *Buffers) {
	v.broadcast(command{kind: 1, actions: actions, seeds: reseeds, buf: b})
}

// Restart starts a new episode in every environment with the given fight seeds and runs each to its first decision.
func (v *VecEnv) Restart(seeds []int64, b *Buffers) {
	v.broadcast(command{kind: 3, seeds: seeds, buf: b})
}

// SetTalents makes environment i play every following episode on this build. Call it between steps.
func (v *VecEnv) SetTalents(i int, talents string) error {
	if _, err := v.Spec.trees.decode(talents); err != nil {
		return err
	}
	v.Envs[i].NextTalents = talents
	return nil
}

func (v *VecEnv) Close() {
	v.broadcast(command{kind: 2})
	for _, wk := range v.workers {
		close(wk.cmds)
	}
}
