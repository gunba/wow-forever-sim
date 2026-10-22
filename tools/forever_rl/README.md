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
  It keeps the best checkpoint (paired evaluations against the published APL), anneals the learning rate and entropy
  bonus over a decision budget, and stops early once the best evaluation is old. It then picks talents among the
  leaders on paired fights and confirms on 4,096 fresh paired fights.
- `run_all.py`: every DPS build at its highest-simming race (tanks excluded), with budgets by fight length (melee
  500M decisions, casters 150M); writes `artifacts/ml/<key>.json` (learned and published DPS on the same fights,
  talent changes, actions, cooldowns, differences, a first-minute timeline) for the review portal's Learned column
  (`tools/forever_bench/ml_pages.py`). Resumable; `budgets.json` in the runs folder overrides budgets between builds.
- `talent_screen.py`: every legal one-point talent move from a base build, played by an APL or a
  model; the leaders are confirmed on fresh fights and the base moves to the best significant change.
- `apl_synth.py`: a learned policy turned into edits to the published APL, in iterations. The environment plays the
  current APL as its teacher (`teacherRotation`, which keeps the trained action space); the network labels its own
  states and the APL's. Rules are induced where the network prefers another action to the APL's choice (an action and
  up to `--literals` thresholds on observation features, each compiled to the APL value that reads the same quantity,
  the environment's `featureAPL`), and actions the network uses far more often are offered unconditionally. The
  simulator then keeps, on paired fights, only the removals of APL entries and the added rules that raise DPS by two
  standard errors, and tunes the thresholds. The result is checked on fresh paired fights against the published APL
  and the network.
- `synth_all.py`: `apl_synth.py` for every learned build, on the learned and (when they differ) the published
  talents, keeping the better profile; parallel jobs, resumable.
- `learned_bench.py`: the converted profiles benched like the matrix (and the published cells with the same binary);
  adds the network's paired gain and the profile to `artifacts/ml/<key>.json` and writes the gaining profiles for
  `export_ui_profiles.py --learned`.
- `compare.py`: a paired table of (build, player) rows, where a player is the published APL, another rotation file,
  or a model.
- `explain.py`: a model against an APL on the same fights: agreement, the most common disagreements, actions per
  fight in and out of the execute phase, rage at each action, and cooldown timing.

Environment options (`Config`):

- `fullChoices`: major cooldowns become agent actions (the teacher uses them under their automatic rules), queued
  on-next-swing attacks can be unqueued, and 0.25 s and 1 s waits are added.
- Decisions where waiting is the only legal action are skipped (`keepWaitOnly` keeps them).
- `pollSeconds` (0 for results): 0 asks the agent at every engine poll, the resolution the APL is checked at. A
  positive value asks after a wait only when something the agent sees changes or that long has passed; that loses
  timing fidelity and is not used for published results.
- Each environment can be given its own talents (`set_talents`), taking effect from its next episode.
- `rotation` replaces the published APL, both as teacher and as the source of the agent's actions.
- `labeledRands` gives one random stream per roll label. Paired fights of two builds then stay aligned: correlation
  rises from 0.12 to 0.60 on Fury, and changes that alter nothing give a difference of exactly zero.

Engine hooks: `APLRotation.SetPolicy`, `PriorityConfigIndexes`, `APLAction.CastSpell`, `Unit.LiveDamage` (a running
total of damage to opponents; the metrics only fold it in when the iteration ends).

Speed, with fights identical bit for bit (same DPS and decisions for the same seeds; byte-identical observations):

- the reward reads a running damage total instead of summing every spell's metrics at each decision;
- with a policy set, a cast entry whose spell is cooling down, or needs the global cooldown while it runs, is skipped
  before its condition is evaluated (Spell.CanCast would reject it anyway);
- talents are written into the observation once per fight.

Together 1.24-1.68x for the environment. The learner keeps 8,192 fights in flight on all threads with page-locked
buffers, draws actions with the Gumbel-max trick, keeps its per-fight bookkeeping in arrays, and spreads the fights
over the 300 s before training (each environment first plays a random number of the published rotation's decisions)
so every batch covers the whole fight. About 500k decisions a second end to end, against 118k before. At full
fidelity a melee decision is ~0.11 s of fight, so rollouts are 128 decisions (~14 s) with GAE lambda 0.97 (~3.6 s);
32-decision rollouts at lambda 0.95 collapsed once the imitation anchor faded.

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
