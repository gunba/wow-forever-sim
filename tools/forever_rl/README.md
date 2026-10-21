# Learning rotations and talents

An agent that picks talents point by point and then plays the five-minute benchmark fight in the real engine, choosing
every action. The reward is damage. Built for Fire first; the environment runs every benchmark build.

## Pieces

- `env/` (Go): the environment. Episodes run in a coroutine, so the simulation stops at each decision.
  - Talent phase (optional): 51 decisions, masked to legal points.
  - Fight phase: one action per decision from the build's castable spells (top rank, utility spells left out), the
    published APL's other entries as macros (move, channel), or wait.
  - Cooldowns the published APL leaves to Autocast Other Cooldowns stay automatic.
  - Observation: talents, time, resources, swing timers, each action's readiness, cooldown, cast time and cost, and each
    timed aura on the player and target.
  - The published rotation's choice is reported at every decision (the teacher).
- `rlbench/` (Go): environment speed, decisions per fight, and teacher replay against the published DPS.
- `capi/` (Go): the environment as a C library for Python:
  `go build -tags with_db -buildmode=c-shared -o ~/.cache/forever-sim-audit/rl/libforeverrl.so ./tools/forever_rl/capi`.
- `forever_env.py`: ctypes wrapper.
- `train.py`: masked PPO with a behaviour-cloning warm start, an optional decaying imitation term (`--bc-coef`) and a KL
  guard.
- `crn_train.py`: approximate policy iteration on exact replays with labelled random streams, reseeded at the labelled
  decision.
- `estimate.py`: training-time estimates per build from `rlbench -json` data.

Engine hooks: `APLRotation.SetPolicy`, `PriorityConfigIndexes`, `APLAction.CastSpell`, `Unit.LiveDamage` (per-iteration
damage is otherwise only folded into the metrics when the iteration ends).

## Measured (Fire, RTX 4070 Ti, 12 of 24 threads)

- Teacher replay matches the published DPS on all 34 builds, within about 1%.
- Environment alone, 8 threads: 4,000-5,000 Fire fights per second.
- Learning loop: 317k-410k decisions per second across builds.
- Behaviour cloning reaches 99.8% agreement in under a minute.
- Fixed talents, evaluation fights where the published rotation scores 927.6:
  - Cloned network: 947.
  - PPO, 86M decisions (5 min): 954.
- End-to-end talents plus rotation in 8 minutes: the talent head commits to a build before the rotation can judge it
  (about 890). A rotation network trained across many builds, then a seeded build search, is the next step.
- Rotation decisions are worth a few DPS each against about 25 DPS of luck once copies diverge, so labelled replays need
  many samples per decision.
