# Stat weights for gear addons

`generate.py` turns the published level-60 benchmark into stat weights and weapon-speed models for every build
the simulator supports, written as one JSON file per class.

```sh
python3 tools/forever_weights/generate.py                      # all classes, about an hour on 22 cores
python3 tools/forever_weights/generate.py --classes druid,paladin
python3 tools/forever_weights/generate.py --smoke              # 100 iterations, checks every build runs
```

- Output: `~/.cache/forever-sim-audit/statw/exports/<class>.json` (`--out` to change). The ForeverGear addon imports these.
- Work cache: `~/.cache/forever-weights/<engine hash>/`. Runs are keyed by the `forever_tanks` binary, so a rebuilt engine
  recomputes everything and an interrupted run resumes. Delete old engine folders when they are no longer needed.
- Go: uses `go` from `PATH`, `$GO`, or `~/.local/share/go/bin/go`, offline (`GOPROXY=off`, `GOTOOLCHAIN=local`).
- Inputs: `artifacts/modelled_gear/forever_dps_5min.json` (DPS requests), `forever_input_profiles.json` (tank players),
  `fury_tank.json` here (the dual-wield Fury tank: SecretProbe tank rotation plus Bloodthirst, Fury 33 / Prot 18).
  Rerun after regenerating the benchmark.

## Method

- One race per build (Tauren, Orc, Troll, Undead, Human, ... first available).
- **Weights:** each stat is moved up and down by a fixed amount with the same random seed and the change in DPS (or the
  tank score) is divided by the change in the stat (central differences, 20,000 iterations). Weights are normalised so the
  profile's main stat is 1: Strength or Agility for physical profiles (Strength for Warriors and Paladins, otherwise the larger),
  spell power for casters.
- **Hit:** probed at 4% and 17% total hit (spell hit: 8%), giving the value below the 8% special cap and, for dual wield,
  between the special and white caps. `talent_hit_pct` is the non-gear hit already in the total (talents, race: Forever
  Tauren get +1%).
- **Capped stats (`cap_models`):** crit, expertise and weapon skill for melee profiles, plus Defense and (shield) block for
  tanks. The character's total is set to each of a few points and the value per point is taken between neighbours, then made
  non-increasing so gear optimisers can treat it as concave. `base` is the non-gear part of the total.
- **Weapon skill:** forward differences, +0 to +5 and +5 to +10, on every melee skill (ranged skills for Hunters).
- **Tanks:** benchmark workloads (one boss for 300 s, three attackers for 90 s). `_lambda_X` = % threat + X × % longer
  time-to-die (max health / damage taken per second), averaged over both workloads.
- **Weapon speed:** a grid of speeds × weapon DPS (60/100/140%) × attack power (+0/+600; ranged attack power for the ranged
  slot), 10,000 iterations, fitted per speed to `h_weapon_dps·D + h_ap14·AP/14 + h_const`. Forever Rage per swing depends
  only on speed, so DPS fits are exact; tank fits have small residuals. `_nowf` models repeat the grid without Windfury Totem.

## Limits

- Level 60 only, modelled level-65 benchmark gear, Horde raid support (Windfury Totem for melee groups).
- No healer profiles: the simulator does not model healing throughput.
- Feral forms use weapon DPS but not speed, so feral profiles have no speed model.
