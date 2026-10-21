// Package main builds the reinforcement-learning environment as a C shared library for the Python learner:
//
//	go build -tags with_db -buildmode=c-shared -o libforeverrl.so ./tools/forever_rl/capi
//
// The caller owns every buffer; each call fills them in place.
package main

/*
#include <stdint.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"github.com/wowsims/classic/tools/forever_rl/env"
)

type handle struct {
	vec *env.VecEnv
	n   int
}

var (
	mu      sync.Mutex
	handles = map[int]*handle{}
	nextID  = 1
	lastErr string
)

func setErr(err any) C.int {
	mu.Lock()
	lastErr = fmt.Sprint(err)
	mu.Unlock()
	return -1
}

//export RLLastError
func RLLastError(buf *C.char, size C.int) C.int {
	mu.Lock()
	defer mu.Unlock()
	return copyString(lastErr, buf, size)
}

func copyString(s string, buf *C.char, size C.int) C.int {
	dst := unsafe.Slice((*byte)(unsafe.Pointer(buf)), int(size))
	n := copy(dst, s)
	if n < int(size) {
		dst[n] = 0
	}
	return C.int(len(s))
}

// RLCreate builds a vector of n environments from a JSON env.Config; it returns a handle or -1.
//
//export RLCreate
func RLCreate(configJSON *C.char, n C.int) (h C.int) {
	defer func() {
		if r := recover(); r != nil {
			h = setErr(r)
		}
	}()
	var cfg env.Config
	if err := json.Unmarshal([]byte(C.GoString(configJSON)), &cfg); err != nil {
		return setErr(err)
	}
	if cfg.Threads > 0 {
		runtime.GOMAXPROCS(cfg.Threads + 1)
	}
	spec, err := env.NewSpec(cfg)
	if err != nil {
		return setErr(err)
	}
	vec := env.NewVecEnv(spec, int(n))
	mu.Lock()
	id := nextID
	nextID++
	handles[id] = &handle{vec: vec, n: int(n)}
	mu.Unlock()
	return C.int(id)
}

func get(h C.int) *handle {
	mu.Lock()
	defer mu.Unlock()
	return handles[int(h)]
}

//export RLSpec
func RLSpec(h C.int, buf *C.char, size C.int) C.int {
	hd := get(h)
	if hd == nil {
		return setErr("bad handle")
	}
	return copyString(string(hd.vec.Spec.SpecJSON()), buf, size)
}

func buffers(hd *handle, obs *C.float, mask *C.uint8_t, teacher *C.int32_t, reward *C.float, done *C.uint8_t, info *C.float) *env.Buffers {
	s := hd.vec.Spec
	b := &env.Buffers{
		Obs:     unsafe.Slice((*float32)(unsafe.Pointer(obs)), hd.n*s.ObsDim),
		Mask:    unsafe.Slice((*uint8)(unsafe.Pointer(mask)), hd.n*s.ActDim),
		Teacher: unsafe.Slice((*int32)(unsafe.Pointer(teacher)), hd.n),
	}
	if reward != nil {
		b.Reward = unsafe.Slice((*float32)(unsafe.Pointer(reward)), hd.n)
		b.Done = unsafe.Slice((*uint8)(unsafe.Pointer(done)), hd.n)
		b.Info = unsafe.Slice((*float32)(unsafe.Pointer(info)), hd.n*4)
	}
	return b
}

//export RLReset
func RLReset(h C.int, obs *C.float, mask *C.uint8_t, teacher *C.int32_t) (rc C.int) {
	defer func() {
		if r := recover(); r != nil {
			rc = setErr(r)
		}
	}()
	hd := get(h)
	if hd == nil {
		return setErr("bad handle")
	}
	hd.vec.Reset(buffers(hd, obs, mask, teacher, nil, nil, nil))
	return 0
}

//export RLStep
func RLStep(h C.int, actions *C.int32_t, obs *C.float, mask *C.uint8_t, teacher *C.int32_t, reward *C.float, done *C.uint8_t, info *C.float) (rc C.int) {
	defer func() {
		if r := recover(); r != nil {
			rc = setErr(r)
		}
	}()
	hd := get(h)
	if hd == nil {
		return setErr("bad handle")
	}
	acts := unsafe.Slice((*int32)(unsafe.Pointer(actions)), hd.n)
	hd.vec.Step(acts, buffers(hd, obs, mask, teacher, reward, done, info))
	return 0
}

// RLRestart starts every environment's next episode with the given fight seeds (nonzero) and fills the first
// decision's buffers.
//
//export RLRestart
func RLRestart(h C.int, seeds *C.int64_t, obs *C.float, mask *C.uint8_t, teacher *C.int32_t) (rc C.int) {
	defer func() {
		if r := recover(); r != nil {
			rc = setErr(r)
		}
	}()
	hd := get(h)
	if hd == nil {
		return setErr("bad handle")
	}
	s := unsafe.Slice((*int64)(unsafe.Pointer(seeds)), hd.n)
	hd.vec.Restart(s, buffers(hd, obs, mask, teacher, nil, nil, nil))
	return 0
}

// RLStepReseed is RLStep with a per-environment reseed applied right after each action (0: none).
//
//export RLStepReseed
func RLStepReseed(h C.int, actions *C.int32_t, reseeds *C.int64_t, obs *C.float, mask *C.uint8_t, teacher *C.int32_t, reward *C.float, done *C.uint8_t, info *C.float) (rc C.int) {
	defer func() {
		if r := recover(); r != nil {
			rc = setErr(r)
		}
	}()
	hd := get(h)
	if hd == nil {
		return setErr("bad handle")
	}
	acts := unsafe.Slice((*int32)(unsafe.Pointer(actions)), hd.n)
	seeds := unsafe.Slice((*int64)(unsafe.Pointer(reseeds)), hd.n)
	hd.vec.StepReseed(acts, seeds, buffers(hd, obs, mask, teacher, reward, done, info))
	return 0
}

// RLSetTalents makes environment i play its following episodes on this build.
//
//export RLSetTalents
func RLSetTalents(h, i C.int, talents *C.char) C.int {
	hd := get(h)
	if hd == nil || int(i) < 0 || int(i) >= hd.n {
		return setErr("bad handle or index")
	}
	if err := hd.vec.SetTalents(int(i), C.GoString(talents)); err != nil {
		return setErr(err)
	}
	return 0
}

// RLTalents writes the talent string of environment i's last finished episode.
//
//export RLTalents
func RLTalents(h, i C.int, buf *C.char, size C.int) C.int {
	hd := get(h)
	if hd == nil || int(i) < 0 || int(i) >= hd.n {
		return setErr("bad handle or index")
	}
	return copyString(hd.vec.Envs[int(i)].LastTalents, buf, size)
}

//export RLClose
func RLClose(h C.int) {
	mu.Lock()
	hd := handles[int(h)]
	delete(handles, int(h))
	mu.Unlock()
	if hd != nil {
		hd.vec.Close()
	}
}

func main() {}
