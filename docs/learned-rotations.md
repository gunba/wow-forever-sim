# Learned rotations and talents

The review matrix's **Learned** column shows, for each DPS build at its highest-simming race, what a reinforcement-learning
agent reached in this simulator when it chose every action itself and searched the talents at the same time. Each
value links to a page with the learned talents, what the agent did differently from the APL, and a first-minute
timeline. Tanks are not included: their benchmark balances threat and survival, not damage.

Each cell has up to two results:

- **The network** (the number): the agent's own DPS. It always shows, whether or not its play could be written as an
  APL. "no gain" means it did not beat the published setup by two standard errors.
- **The simmable profile** ("sim" under the number): the published setup for that race with the chosen talents and the
  published APL edited by rules converted from the network. It is an ordinary profile in the simulator's Ranked builds
  selector, benched like every other cell, and the link opens it. It appears only where it beats the published setup.
  "no APL" means no conversion kept a gain yet; the page then lists the network's talents and actions, the place to
  start investigating what it does.

## The network

- A neural network chose every action in the real engine: casts, major cooldowns, Heroic Strike/Cleave queueing and
  unqueueing, and waits. It could act at every point the APL is checked (the engine's 50 ms poll). It saw talents,
  time, resources, swing timers, every action's readiness, cooldown and cost, and every timed aura on itself and the
  target.
- It started by copying the published APL. It was then trained with PPO on the total damage of five-minute fights.
- One network learned across thousands of talent builds at once. A build search, fed by the same fights, decided which
  builds it played: a population of leaders, mutations of them, and a surrogate model of build DPS. The page shows
  the build that won paired fights against the published talents. If none won clearly, the published talents stay.
- Its gain is measured on paired fights: the network and the published setup meet the same 2,048 fights, with one
  random stream per roll label, so the difference is measured, not two separate averages. The published setup is
  played by the engine's own APL path. (Training compared the network with an emulation of the APL that uses major
  cooldowns at decision points; for eight builds that emulation differs from the engine by 0.1-0.6%, so training's
  own figures are not used.)

## The simmable profile

[`apl_synth.py`](../tools/forever_rl/apl_synth.py) edits the published APL in a few iterations:

1. The network and the current APL each play fights. In every state the network's preferences and the APL's choice are
   logged.
2. Candidate rules are induced where the network prefers another action to the APL's choice: an action and up to two
   thresholds on what the network sees, such as "cast Frostbolt if Arcane Blast's cost is at least 24% of current
   mana and Missile Barrage is not active". Each threshold compiles to the APL value that reads the same quantity.
   Actions the network uses far more often than the APL are offered unconditionally.
3. The simulator decides. Removing each published entry, and adding each candidate rule above the APL, is tested on
   paired fights, and an edit is kept only if it raises DPS by two standard errors. Thresholds are then tuned the
   same way. The result is the next iteration's APL.

Builds whose learned talents differ from the published ones are converted on both, and the better profile is kept:
an APL may not exploit the learned talents the way the network does. The page lists every added and removed entry
and the share of the network's gain the profile keeps.

## What it is not

- **A player target.** The network reacts instantly and knows every timer exactly. It is this simulator's ceiling for
  the build under its current mechanics; the profile keeps the part of it that simple APL rules express.
- **Independent of the model.** It uses whatever the simulator allows (for example Demonic Rune's self-damage for
  rage), so any modelling error it leans on shows up as a gain. Large gains were checked against the code. The Hunter
  weave builds gain mostly from Forever's instant, 6-second Multi-Shot: they fire it, and Arcane Shot, while
  stepping between melee and range. Fired on the move it is legal and modelled with continuous movement, the dead
  zone and the swing timers.
- **Pet control.** Pets act on their own.

## Reproducing

[tools/forever_rl](../tools/forever_rl/README.md): `run_all.py` runs every build and writes `artifacts/ml/<key>.json`;
`codesign.py` trains one build; `explain.py` writes the action report; `synth_all.py` converts every build with
`apl_synth.py`; `learned_bench.py` benches the profiles like the matrix and adds them to the reports.
