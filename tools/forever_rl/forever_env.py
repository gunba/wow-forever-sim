"""ctypes wrapper for the simulator's reinforcement-learning environment (tools/forever_rl/capi)."""
import ctypes
import json
import os
from pathlib import Path

import numpy as np

REPO = Path(__file__).resolve().parents[2]
DEFAULT_LIB = Path(os.environ.get("FOREVER_RL_LIB", Path.home() / ".cache/forever-sim-audit/rl/libforeverrl.so"))


class VecEnv:
    """n environments stepped together. Every array is reused between calls; copy what you keep."""

    def __init__(self, config: dict, n: int, lib: Path = DEFAULT_LIB):
        self.lib = ctypes.CDLL(str(lib))
        L = self.lib
        f32p, u8p, i32p = (ctypes.POINTER(t) for t in (ctypes.c_float, ctypes.c_uint8, ctypes.c_int32))
        L.RLCreate.argtypes, L.RLCreate.restype = [ctypes.c_char_p, ctypes.c_int], ctypes.c_int
        L.RLSpec.argtypes, L.RLSpec.restype = [ctypes.c_int, ctypes.c_char_p, ctypes.c_int], ctypes.c_int
        L.RLReset.argtypes, L.RLReset.restype = [ctypes.c_int, f32p, u8p, i32p], ctypes.c_int
        L.RLStep.argtypes, L.RLStep.restype = [ctypes.c_int, i32p, f32p, u8p, i32p, f32p, u8p, f32p], ctypes.c_int
        L.RLTalents.argtypes, L.RLTalents.restype = [ctypes.c_int, ctypes.c_int, ctypes.c_char_p, ctypes.c_int], ctypes.c_int
        L.RLLastError.argtypes, L.RLLastError.restype = [ctypes.c_char_p, ctypes.c_int], ctypes.c_int
        L.RLClose.argtypes, L.RLClose.restype = [ctypes.c_int], None
        L.RLStepReseed.argtypes, L.RLStepReseed.restype = [ctypes.c_int, i32p, ctypes.POINTER(ctypes.c_int64), f32p, u8p, i32p, f32p, u8p, f32p], ctypes.c_int
        L.RLRestart.argtypes, L.RLRestart.restype = [ctypes.c_int, ctypes.POINTER(ctypes.c_int64), f32p, u8p, i32p], ctypes.c_int
        cfg = {"repo": str(REPO), **config}
        self.h = L.RLCreate(json.dumps(cfg).encode(), n)
        if self.h < 0:
            raise RuntimeError(self._error())
        buf = ctypes.create_string_buffer(1 << 20)
        L.RLSpec(self.h, buf, len(buf))
        self.spec = json.loads(buf.value.decode())
        self.n = n
        D, A = self.spec["obsDim"], self.spec["actDim"]
        self.obs = np.zeros((n, D), np.float32)
        self.mask = np.zeros((n, A), np.uint8)
        self.teacher = np.zeros(n, np.int32)
        self.reward = np.zeros(n, np.float32)
        self.done = np.zeros(n, np.uint8)
        self.info = np.zeros((n, 4), np.float32)
        self._actions = np.zeros(n, np.int32)

    def _error(self):
        buf = ctypes.create_string_buffer(4096)
        self.lib.RLLastError(buf, len(buf))
        return buf.value.decode()

    @staticmethod
    def _p(a, t):
        return a.ctypes.data_as(ctypes.POINTER(t))

    def reset(self):
        rc = self.lib.RLReset(self.h, self._p(self.obs, ctypes.c_float), self._p(self.mask, ctypes.c_uint8),
                              self._p(self.teacher, ctypes.c_int32))
        if rc < 0:
            raise RuntimeError(self._error())
        return self.obs, self.mask, self.teacher

    def step(self, actions):
        self._actions[:] = actions
        rc = self.lib.RLStep(self.h, self._p(self._actions, ctypes.c_int32), self._p(self.obs, ctypes.c_float),
                             self._p(self.mask, ctypes.c_uint8), self._p(self.teacher, ctypes.c_int32),
                             self._p(self.reward, ctypes.c_float), self._p(self.done, ctypes.c_uint8),
                             self._p(self.info, ctypes.c_float))
        if rc < 0:
            raise RuntimeError(self._error())
        return self.obs, self.mask, self.teacher, self.reward, self.done, self.info

    def step_reseed(self, actions, reseeds):
        """step() that reseeds each environment's random streams right after its action (0: no reseed)."""
        self._actions[:] = actions
        self._reseeds = np.ascontiguousarray(reseeds, dtype=np.int64)
        rc = self.lib.RLStepReseed(self.h, self._p(self._actions, ctypes.c_int32), self._p(self._reseeds, ctypes.c_int64),
                                   self._p(self.obs, ctypes.c_float), self._p(self.mask, ctypes.c_uint8),
                                   self._p(self.teacher, ctypes.c_int32), self._p(self.reward, ctypes.c_float),
                                   self._p(self.done, ctypes.c_uint8), self._p(self.info, ctypes.c_float))
        if rc < 0:
            raise RuntimeError(self._error())
        return self.obs, self.mask, self.teacher, self.reward, self.done, self.info

    def restart(self, seeds):
        """Start every environment's next episode with these fight seeds (nonzero int64) and return the first decision."""
        self._seeds = np.ascontiguousarray(seeds, dtype=np.int64)
        rc = self.lib.RLRestart(self.h, self._p(self._seeds, ctypes.c_int64), self._p(self.obs, ctypes.c_float),
                                self._p(self.mask, ctypes.c_uint8), self._p(self.teacher, ctypes.c_int32))
        if rc < 0:
            raise RuntimeError(self._error())
        return self.obs, self.mask, self.teacher

    def talents(self, i):
        buf = ctypes.create_string_buffer(256)
        self.lib.RLTalents(self.h, i, buf, len(buf))
        return buf.value.decode()

    def close(self):
        if self.h >= 0:
            self.lib.RLClose(self.h)
            self.h = -1
