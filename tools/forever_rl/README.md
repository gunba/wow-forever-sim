# Learning rotations and talents

An agent that picks talents point by point and then plays the five-minute benchmark fight in the real engine, choosing
every action. The reward is damage. Built for Fire first; the environment runs every benchmark build.

## Pieces

- `env/` (Go): the environment. Episodes run in a coroutine, so the simulation stops at each decision.
  - Talent phase (optional): 51 decisions, masked to legal points.
  - Fight phase: one action per decision from the build's castable spells (top rank, utility spells left out), the
    published APL's other entries as macros (move, channel), or wait.
  - Cooldowns the published APL leaves to Autocast Other Cooldowns stay automatic, unless `fullChoices` is set.
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
- `codesign.py`: talents and rotation together. One network, conditioned on the talents, is trained by PPO
  across many builds at once; a build search picks the build of every next fight from the same stream of fights:
  - a population scored by decayed fight statistics, with Thompson sampling over its leaders;
  - mutations of the leaders (talent points moved);
  - a five-network surrogate (build -> DPS) trained on the fights, proposing builds by mean + spread;
  - shared-seed evaluations of the leaders against the same network and the APL on the published build.
  It keeps the best checkpoint and anneals the learning rate and entropy bonus.
- `talent_screen.py`: every legal one-point talent move from a base build, played by an APL or a
  model; the leaders are confirmed on fresh fights and the base moves to the best significant change.
- `compare.py`: a paired table of (build, player) rows, where a player is the published APL, another rotation file,
  or a model.
- `explain.py`: a model against an APL on the same fights: agreement, the most common disagreements, actions per
  fight in and out of the execute phase, rage at each action, and cooldown timing.

Environment options (`Config`):

- `fullChoices`: major cooldowns become agent actions (the teacher uses them under their automatic rules), queued
  on-next-swing attacks can be unqueued, and 0.25 s and 1 s waits are added.
- Decisions where waiting is the only legal action are skipped (`keepWaitOnly` keeps them).
- `pollSeconds`: after a wait, the agent is asked again as soon as anything it sees changes, and otherwise after this
  long.
- Each environment can be given its own talents (`set_talents`), taking effect from its next episode.
- `rotation` replaces the published APL, both as teacher and as the source of the agent's actions.
- `labeledRands` gives one random stream per roll label. Paired fights of two builds then stay aligned: correlation
  rises from 0.12 to 0.60 on Fury, and changes that alter nothing give a difference of exactly zero.

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
  (about 890). `codesign.py` replaces it: one network across many builds, with a build search beside it.
- Rotation decisions are worth a few DPS each against about 25 DPS of luck once copies diverge, so labelled replays need
  many samples per decision.

## Fury Warrior (Orc, modelled dual-wield gear)

Paired results on 4,096 fresh fights (standard error about 0.5 DPS), with the published gear.
"Best talents" is `30305213002-25053005112010501-1`: two Enrage points moved to Improved Heroic Strike 3/3 and
Improved Bloodrage 1/2.

| Rotation | Talents | DPS | vs published |
|---|---|---|---|
| Published APL | published | 982.4 | |
| APL, round 1 rules | published | 997.6 | +15.2 |
| + round 2 rules without Demonic Rune | published | 1001.0 | +18.6 |
| + Demonic Rune | published | 1004.1 | +21.7 |
| APL, all rules | best | 1005.6 | +23.2 |
| Co-design network | published | 1010.8 | +28.3 |
| Co-design network | best | 1011.9 | +29.4 |

The network's habits, measured against the APL on the same fights with `explain.py` and written back as APL rules:

- Round 1:
  - Dance to Battle Stance for Overpower in the execute phase and during Recklessness, up to 35 rage (the published
    APL stops at 25, outside the execute phase and Recklessness).
  - Refresh Rend while in Battle Stance.
  - Queue Heroic Strike from 40 rage (was 30).
- Round 2:
  - No Rend in the execute phase.
  - Hamstring for Windfury only when the next main-hand swing is at least 1 s away.
  - Demonic Rune at up to 60 rage. Its self-damage gives rage, and the benchmark already carries it.

The tuned APL then taught the second co-design run. The network's remaining 6 DPS comes from finer timing that
thresholds did not capture: late Heroic Strike queueing lost DPS at every setting tried.

The same rules on the other Fury APLs, each against its published APL:

- Fury Sunder (no Rend refresh, 40 rage Heroic Strike): +18.6.
- 2H Fury (dance up to 60 rage, 80 rage Heroic Strike): +36.2.

Talents:

- Under every rotation tried, the published build is within 3 DPS of the best one-point neighbourhood.
- Enrage points beyond the tier gates are worth 0.33 DPS each.
- Booming Voice does nothing (the APL never shouts).
- Improved Berserker Rage does nothing, because Berserker Rage is a survival cooldown and is never cast. Casting it
  in idle GCDs breaks even: it displaces the Hamstring weave.
- Weaponmaster hybrids (25 Arms points, no Bloodthirst) lose 140-158 DPS; Bloodthrill loses 183.

Weapons: with racial weapon skills gone under Forever, weapon type matters only through Weaponmaster. A one-off
check on a Weaponmaster build put the Sword ahead of Axe and Mace. Arms, the only published Weaponmaster build,
already wields the Sword. Builds are not searched over weapons.

What did not work:

- An end-to-end talent head commits to a build before the rotation can judge it.
- A shared policy drifts on any one build without annealing and a kept best checkpoint. The first run peaked at +28
  and fell to +6.
- The imitation anchor must fade, but slowly. The first gains came only after it reached zero, at about 800
  iterations.
