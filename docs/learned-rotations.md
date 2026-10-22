# Learned rotations and talents

The review matrix's **Learned** column shows, for each DPS build at its highest-simming race, what a reinforcement-learning
agent reached in this simulator when it chose every action itself and searched the talents at the same time. Each
value links to a page with the learned talents, what the agent did differently from the APL, and a first-minute
timeline. Tanks are not included: their benchmark balances threat and survival, not damage.

## What it is

- A neural network chose every action in the real engine: casts, major cooldowns, Heroic Strike/Cleave queueing and
  unqueueing, and waits. It could act at every point the APL is checked (the engine's 50 ms poll). It saw talents,
  time, resources, swing timers, every action's readiness and cooldown, and every timed aura on itself and the target.
- It started by copying the published APL. It was then trained with PPO on the total damage of five-minute fights.
- One network learned across thousands of talent builds at once. A build search, fed by the same fights, decided which
  builds it played: a population of leaders, mutations of them, and a surrogate model of build DPS. The page shows
  the build that won paired fights against the published talents. If none won clearly, the published talents stay.
- Every number comes from paired fights: the published APL and the learned policy meet the same 4,096 fights, with one
  random stream per roll label, so the difference is measured, not two separate averages. Standard errors are
  typically 0.3-0.6 DPS.

## What it is not

- **A player target.** The policy reacts instantly and knows every timer exactly. It is this simulator's ceiling for the
  build under its current mechanics.
- **An APL.** It cannot be loaded into the simulator. The published setup for the same race is linked from each page.
- **Independent of the model.** It uses whatever the simulator allows (for example Demonic Rune's self-damage for
  rage), so any modelling error it leans on shows up as a gain. Large gains were checked against the code. The Hunter
  weave builds gain mostly from Forever's instant, 6-second Multi-Shot: they fire it, and Arcane Shot, while
  stepping between melee and range. Fired on the move it is legal and modelled with continuous movement, the dead
  zone and the swing timers.
- **Pet control.** Pets act on their own.

## Reproducing

[tools/forever_rl](../tools/forever_rl/README.md): `run_all.py` runs every build and writes `artifacts/ml/<key>.json`;
`codesign.py` trains one build; `explain.py` writes the action report; `apl_synth.py` turns a learned policy into
APL overrides that the simulator can run.
