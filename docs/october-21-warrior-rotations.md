# October 21 Warrior rotations

A third pass on 21 October. A reinforcement-learning agent played the five-minute Fury benchmark in the engine and chose
every action, including the cooldowns the rotation otherwise leaves to Autocast. Where it beat the published rotation, its
habits were written back as APL rules and confirmed on paired fights. Fury, Fury (Sunder) and 2H Bloodthirst now use those
rules, and Fury's talents move two Enrage points. All 872 matrix scenarios are re-run (seed 1793600001); no tank input or
engine rule changed, so the tank checks stand. Detailed numbers are in the
[comparison](../artifacts/correctness/october_21_warrior/comparison.json) and the
[selection evidence](../artifacts/correctness/october_21_warrior/selection-evidence.json).

## Rotations

Measured on 4,096 paired fights (shared seeds, one random stream per roll label, standard errors about 0.5 DPS) against
each build's previous APL, with the previous talents:

- **Overpower in the execute phase.** The rotation now dances to Battle Stance for an Overpower window up to 35 rage
  (60 for 2H Bloodthirst), also in the execute phase and during Recklessness. It used to stop at 25 rage, the most a
  swap keeps with Improved Tactical Mastery 5/5, and skip both. The Overpower is worth the rage a swap above 25 loses.
  On its own this was worth +6.9 DPS on Fury.
- **Rend in Battle Stance.** While the dance is in Battle Stance, Rend is refreshed if it has fallen off, outside the
  execute phase. Fury (Sunder) leaves this out: its Sunder upkeep needs those globals.
- **Heroic Strike thresholds.** Fury and Fury (Sunder) queue Heroic Strike from 40 rage (was 30), 2H Bloodthirst from 80
  (was 50). Lower thresholds starve Bloodthirst and Whirlwind; queueing at lower rage just before a swing lost DPS at every
  setting tried.
- **Hamstring timing.** The Windfury Hamstring weave (WAR-027) now also waits until the next main-hand swing is at least
  1 s away.
- **Demonic Rune.** The benchmark already gives Warriors Demonic Runes. Their 600-1000 self-damage (client spell 16666)
  is damage taken and gives rage, so the rotation uses one at up to 60 rage, three times a fight. Whether the game agrees
  is registered for a check (WAR-029).

| Build | Rules | Previous rotation | New rotation | Change |
|---|---|---|---|---|
| Fury (Orc) | all | 982.4 | 1004.1 | +21.7 |
| Fury (Sunder) (Orc) | all, no Rend refresh | 934.4 | 953.0 | +18.6 |
| 2H Bloodthirst (Orc) | all, own thresholds | 832.4 | 868.6 | +36.2 |

(WAR-028)

## Talents

Fury and Fury (Sunder) move to 19/31/1 (`30305213002-25053005112010501-1`): two Enrage points go to Improved Heroic
Strike 3/3 and Improved Bloodrage 1/2. Beyond the points the tier gates need, Enrage is worth 0.33 DPS a point (it procs
only from damage taken). Every legal one-point move was screened on paired fights under both the old and the new rotation,
and only these two cleared 2.5 standard errors: +1.5 DPS together under the new rotation. Booming Voice and Improved
Berserker Rage do nothing here (the rotation never shouts; Berserker Rage is a survival cooldown and casting it in idle
globals breaks even). Builds that trade Bloodthirst and Flurry for Weaponmaster or Bloodthrill lose 140 to 183 DPS.

## How the rules were found

The agent ([tools/forever_rl](../tools/forever_rl/README.md)) sees the talents, resources, swing timers and every timed
aura, and can also unqueue Heroic Strike or wait. One network conditioned on the talents was trained across thousands of
Fury builds while a build search, fed by the same fights, chose which builds it played. Its choices were then compared with
the APL's on the same fights; the habits with the largest share of the difference became rules, were tuned on paired
fights, and the tuned APL taught a second run. That run's network is still about 6 DPS ahead of the new rules on Fury
(1011.9 against 1005.6), from finer swing timing that thresholds did not capture; it is a network, not an APL, and is not
used for the benchmark.

## Not changed

Weapons: with racial weapon skills gone under Forever, weapon type matters only through Weaponmaster, and Arms, the only
Weaponmaster build, already wields the Sword. Arms, Protection, Fury (Adaptive AoE) and Fury (Giants/Dragonkin) keep
their rotations; they were not tested.

## Results

Best race per Warrior build, 5-minute benchmark (seed 1793600001), against the October 21 follow-up. Every other
build replays exactly as before.

| Build | Best DPS | Was | Mean change |
|---|---|---|---|
| Warrior Fury | 1008 (Windshaper) | 988 | +2.2% |
| Warrior Fury (Sunder) | 972 (Human) | 953 | +2.1% |
| Warrior 2H Bloodthirst | 870 (Orc) | 834 | +4.2% |
