# Build reviews

These summaries cover completed build comparisons. They are simulation results, not independent confirmation of server mechanics or proof of a global optimum.

The tables below use each build's independent validation run. The [final matrix](../artifacts/forever_dps_5min.png) uses a separate common-seed run of all 147 combinations, so small Monte Carlo differences are expected. The final matrix also includes the later Eureka nested-charge correction; the optimization validation tables remain historical records.

The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific Tier 1 bonuses, and paid shared-hit normalization. [Scenario and exchange model](../tools/forever_bench/README.md) · [In-game checks](in_game_checks.md)

## Paladin — Retribution

**Talents:** 12/0/39 · `250003002--052253312012331321`

[Requests, results and search evidence](../artifacts/optimization/retribution.json)

- The retained 12/0/39 allocation moves Deflection 5→0 into Vindication 3/3 and Reverence 2/3. Seal of Command and Twist of Light remain talented.
- Prioritize Exorcism, retain Judgement and the existing Command/Righteousness switching, and omit Consecration on the fixed single target. The existing 15% mana condition on seal switching remains.
- An earlier talent search was invalidated because the engine incorrectly granted Seal of Command without its point. The availability check is now enforced from the current client talent records; the final build pays the point and has been rerun.
- All three Paladin races passed 5,000-iteration comparisons under the corrected engine, with negligible mana-limited time and no APL warnings. No gear, external buffs or consumables changed.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Undead | 864.91 | 896.42 | +3.64% | 0.00 |
| Human | 849.41 | 881.89 | +3.82% | 0.01 |
| Dwarf | 847.91 | 880.68 | +3.86% | 0.01 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -1.5s: [Spell 20920 (rank 5)](https://www.wowhead.com/forever/spell=20920).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Holy Strike (rank 8)](https://www.wowhead.com/forever/spell=10333).
3. Cast [Judgement of Light](https://www.wowhead.com/forever/spell=20271) when `{"spellCanCast":{"spellId":{"spellId":20271}}}`.
4. Cast [Spell 20920 (rank 5)](https://www.wowhead.com/forever/spell=20920) when `{"currentSealRemainingTime":{}}` ≤ 1.
5. Cast [Spell 20293 (rank 8)](https://www.wowhead.com/forever/spell=20293) when (Mana fraction ≥ 15% AND NOT [Twist of Light](https://www.wowhead.com/forever/spell=77485) active AND [Spell 20920 (rank 5)](https://www.wowhead.com/forever/spell=20920) active).
6. Cast [Spell 20920 (rank 5)](https://www.wowhead.com/forever/spell=20920) when (Mana fraction ≥ 15% AND NOT [Twist of Light](https://www.wowhead.com/forever/spell=77485) active AND [Spell 20293 (rank 8)](https://www.wowhead.com/forever/spell=20293) active).
7. Cast [Hammer of Wrath (rank 3)](https://www.wowhead.com/forever/spell=24239).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Seal of Command](https://www.wowhead.com/forever/spell=20424) | 215.00 |
| Auto-attack (tag 1) | 211.53 |
| [Seal of Righteousness](https://www.wowhead.com/forever/spell=25713) | 170.68 |
| Auto-attack (tag 3) | 78.69 |
| [Holy Strike](https://www.wowhead.com/forever/spell=10333) | 76.23 |
| [Judgement of Command](https://www.wowhead.com/forever/spell=20966) | 48.76 |
| [Judgement of Righteousness](https://www.wowhead.com/forever/spell=20286) | 35.21 |
| [Hammer of Wrath](https://www.wowhead.com/forever/spell=24239) | 24.51 |

## Warrior — Fury

**Talents:** 17/34/0 · `20305113002-050520035151010051`

[Requests, results and search evidence](../artifacts/optimization/fury.json)

- The retained 17/34/0 allocation moves one point from Improved Heroic Strike into Improved Overpower. All talent prerequisites, including Enrage for Flurry, remain funded.
- Allow ordinary, untalented Slam at 30 rage rather than excluding its 1.5-second cast. Continuous melee swings make it useful without Improved Slam. Use Recklessness on the pull with the existing cooldown package.
- The existing Overpower stance dance, Execute priority and 30-rage Heroic Strike condition outside execute remain. Rage thresholds, queue priority, major-ability order and alternative Recklessness windows were screened.
- All ten races passed 5,000-iteration comparisons. The off-hand retains its dual-wield white miss chance while Heroic Strike is queued; no queue/cancel hit exploit is used. Rage and combat-table uncertainties remain T01/T02/T06.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 773.35 | 799.48 | +3.38% | 0.00 |
| Tauren | 772.22 | 797.36 | +3.26% | 0.00 |
| Troll | 770.01 | 796.90 | +3.49% | 0.00 |
| Undead | 784.36 | 810.58 | +3.34% | 0.00 |
| Windshaper | 771.82 | 795.57 | +3.08% | 0.00 |
| Human | 781.45 | 806.80 | +3.24% | 0.00 |
| Dwarf | 772.93 | 797.07 | +3.12% | 0.00 |
| Night Elf | 773.23 | 795.00 | +2.82% | 0.00 |
| Gnome | 764.22 | 789.62 | +3.32% | 0.00 |
| High Order | 771.82 | 795.57 | +3.08% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719) when Time remaining ≤ 300s.
2. Use ready automatic cooldowns.
3. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when ([Overpower](https://www.wowhead.com/forever/spell=11585) active AND Rage ≤ 25 AND NOT `{"isExecutePhase":{"threshold":"E20"}}` AND NOT [Recklessness](https://www.wowhead.com/forever/spell=1719) active).
4. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
5. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when NOT [Overpower](https://www.wowhead.com/forever/spell=11585) active.
6. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
7. Cast [Bloodthirst](https://www.wowhead.com/forever/spell=23894).
8. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
9. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 30.
10. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 30 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 220.23 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 122.50 |
| Auto-attack (tag 2) | 119.73 |
| [Bloodthirst](https://www.wowhead.com/forever/spell=23894) | 92.27 |
| Auto-attack (tag 1) | 61.31 |
| Auto-attack (tag 3) | 46.14 |
| [Whirlwind](https://www.wowhead.com/forever/spell=1680) | 37.49 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 26.74 |

## Hunter — Beast Mastery

**Talents:** 31/20/0 · `5320001505101251-00531510005`

[Requests, results and search evidence](../artifacts/optimization/beast_mastery.json)

- The retained 31/20 build moves Efficiency 5→1 to fund Improved Stings 0→3 and Mortal Shots 4→5 relative to the frozen baseline. Pet settings and equipment are unchanged.
- Rank-4 Aimed Shot replaces Multi-Shot as the main shared-cooldown shot and has no swing-window wait. Arcane Shot retains its 15% mana floor; Intimidation requires more than 80% mana.
- All Aimed Shot ranks, shot windows, dot omissions, mana floors and end-of-fight spending were screened under continuous Auto Shot. All eight races passed 5,000-iteration comparisons, with under one second of mana-limited time.
- Summon Hawk guardian coexistence and scaling remain an approximation awaiting T14. Historical search entries predate the auto-attack correction.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 751.55 | 772.13 | +2.74% | 0.88 |
| Tauren | 749.50 | 770.17 | +2.76% | 0.97 |
| Troll | 751.96 | 771.95 | +2.66% | 0.91 |
| Windshaper | 752.21 | 773.11 | +2.78% | 0.70 |
| Human | 748.94 | 769.52 | +2.75% | 0.56 |
| Dwarf | 747.22 | 767.41 | +2.70% | 0.85 |
| Night Elf | 757.24 | 778.15 | +2.76% | 0.78 |
| High Order | 752.21 | 773.11 | +2.78% | 0.70 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -10s: [Aspect of the Hawk](https://www.wowhead.com/forever/spell=25296).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Intimidation](https://www.wowhead.com/forever/spell=19577) when Mana fraction > 80%.
3. Cast [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) when NOT [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) DoT active.
4. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=25295) when NOT [Serpent Sting](https://www.wowhead.com/forever/spell=25295) DoT active.
5. Cast [Aimed Shot](https://www.wowhead.com/forever/spell=20902) when `{"autoTimeToNext":{"autoType":"Ranged"}}` ≥ 0s.
6. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287) when Mana fraction ≥ 15%.

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| Shoot | 325.28 |
| Cat: Auto-attack (tag 1) | 148.90 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20902) | 111.20 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 67.54 |
| [Summon Hawk](https://www.wowhead.com/forever/spell=1293527) | 63.58 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 38.60 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 12.57 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 10.48 |

## Mage — Fire

**Talents:** 17/31/3 · `0501252000002-23450000130133051-003`

[Requests, results and search evidence](../artifacts/optimization/fire.json)

- The 17/31/3 build adds Elemental Precision 3/3, funded by one point each from Arcane Subtlety, Improved Fireball and Arcane Meditation. The hit talent reduces the paid hit-budget deduction rather than granting free item stats.
- The fallback filler is maximum-rank Scorch instead of rank 1. Fireball uses a 150-mana-per-remaining-second reserve, retaining the existing Hot Streak Pyroblast, Fire Blast, Scorch debuff, Mana Gem and Evocation priorities.
- Fireball ranks, Frostfire ranks, mana reserves and Fire Blast/Pyroblast omission were tested before and after talent changes. No unverified Frostfire interaction was added.
- All six races passed independent 5,000-iteration comparisons. Mana-limited time is 1.43–2.29 seconds; existing gear, potion/rune choices and external buffs remain fixed.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 698.70 | 759.46 | +8.70% | 2.22 |
| Troll | 694.03 | 754.07 | +8.65% | 2.34 |
| Undead | 708.53 | 766.60 | +8.20% | 2.24 |
| Human | 695.67 | 757.07 | +8.83% | 2.06 |
| Gnome | 699.73 | 762.36 | +8.95% | 1.58 |
| High Order | 696.84 | 757.01 | +8.64% | 2.16 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 80%.
2. Use ready automatic cooldowns.
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 20% AND Time remaining > 25s).
4. Cast [Pyroblast (rank 8)](https://www.wowhead.com/forever/spell=18809) when [Hot Streak](https://www.wowhead.com/forever/spell=44445) stacks = 3.
5. Cast [Fire Blast (rank 7)](https://www.wowhead.com/forever/spell=10199) when Mana ≥ Time remaining × 10.
6. Cast [Scorch (rank 7)](https://www.wowhead.com/forever/spell=10207) when ([Improved Scorch](https://www.wowhead.com/forever/spell=12873) stacks < 5 OR `{"auraRemainingTime":{"auraId":{"spellId":12873}}}` ≤ 5s).
7. Cast [Fireball](https://www.wowhead.com/forever/spell=25306) when Mana ≥ Time remaining × 150.
8. Cast [Scorch](https://www.wowhead.com/forever/spell=10207).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Pyroblast](https://www.wowhead.com/forever/spell=18809) | 428.52 |
| [Scorch](https://www.wowhead.com/forever/spell=10207) | 131.92 |
| [Fire Blast](https://www.wowhead.com/forever/spell=10199) | 95.74 |
| [Ignite](https://www.wowhead.com/forever/spell=12654) | 79.93 |
| [Fireball](https://www.wowhead.com/forever/spell=25306) | 18.86 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 11.63 |

## Hunter — Marksmanship

**Talents:** 18/33/0 · `5023000503-0053451001503051`

[Requests, results and search evidence](../artifacts/optimization/marksmanship.json)

- The retained 18/33/0 allocation includes Sniper Shot and its Trueshot Aura prerequisite, with Ferocity 3/5. Existing external buffs remain fixed; no new raid buff is credited for taking the prerequisite.
- Use Sniper Shot above 40% mana, maintain Serpent Sting, then use rank-3 Aimed Shot and Arcane Shot without swing-window waits. Multi-Shot is omitted on the single target.
- Sniper Shot became viable after removing inherited shot pauses and added wind-up. Initial versions lost 19–31 seconds to failed mana checks; the lower Aimed Shot rank and explicit Sniper mana floor reduce that to 1.72–2.30 seconds across races.
- All eight races passed 5,000-iteration comparisons against the same-engine frozen baseline, gaining 4.28–4.39%. All Aimed Shot ranks and several shot/mana priorities were screened; no gear or consumable changed.
- Historical search entries used the interrupted-shot model and are not evidence against the current Sniper build.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 702.66 | 732.75 | +4.28% | 2.05 |
| Tauren | 700.47 | 730.45 | +4.28% | 2.30 |
| Troll | 702.53 | 732.61 | +4.28% | 2.21 |
| Windshaper | 702.84 | 733.66 | +4.39% | 1.74 |
| Human | 699.54 | 729.93 | +4.34% | 1.72 |
| Dwarf | 697.57 | 727.60 | +4.30% | 1.90 |
| Night Elf | 708.30 | 738.88 | +4.32% | 1.80 |
| High Order | 702.84 | 733.66 | +4.39% | 1.74 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -10s: [Aspect of the Hawk](https://www.wowhead.com/forever/spell=25296).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) when Mana fraction ≥ 40%.
3. Cast [Serpent Sting](https://www.wowhead.com/forever/spell=25295) when NOT [Serpent Sting](https://www.wowhead.com/forever/spell=25295) DoT active.
4. Cast [Aimed Shot](https://www.wowhead.com/forever/spell=20901) when `{"autoTimeToNext":{"autoType":"Ranged"}}` ≥ 0s.
5. Cast [Arcane Shot](https://www.wowhead.com/forever/spell=14287) when `{"autoTimeToNext":{"autoType":"Ranged"}}` ≥ 0s.

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| Shoot | 341.70 |
| [Aimed Shot](https://www.wowhead.com/forever/spell=20901) | 113.23 |
| Cat: Auto-attack (tag 1) | 103.14 |
| [Serpent Sting](https://www.wowhead.com/forever/spell=25295) | 62.46 |
| [Arcane Shot](https://www.wowhead.com/forever/spell=14287) | 56.90 |
| [Sniper Shot](https://www.wowhead.com/forever/spell=1310786) | 45.26 |
| Cat: [Claw](https://www.wowhead.com/forever/spell=3009) | 12.56 |
| Cat: [Bite](https://www.wowhead.com/forever/spell=17261) | 3.64 |

## Hunter — Survival

**Talents:** 7/11/33 · `502-0050051-230230230250022151`

[Requests, results and search evidence](../artifacts/optimization/survival.json)

- The retained 7/11/33 build adds Deadly Aspects 5/5 and Focused Fire 2/2, funded by Efficiency 4→0 and Mortal Shots 3→0.
- Surefooted 1→3 is funded by Improved Tracking 3→2 and Deterrence 1→0. Its hit benefit reduces the paid hit-budget deduction; catalog gear stats are unchanged.
- Deadly Aspects enables the modeled 30% melee-haste proc under Aspect of the Beast. Ten independently validated one-point talent moves were retained; a subsequent search found no further validated move.
- Raptor Strike, Strider Kick, Mongoose Bite and Immolation Trap priorities were retested after the talent changes. The original APL remained best within the tested choices.
- All eight races passed 5,000-iteration comparisons with zero mana-limited time. The fixed melee distance, weapon and pet remain unchanged.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 587.57 | 662.37 | +12.73% | 0.00 |
| Tauren | 585.30 | 659.66 | +12.70% | 0.00 |
| Troll | 585.93 | 660.88 | +12.79% | 0.00 |
| Windshaper | 586.61 | 661.36 | +12.74% | 0.00 |
| Human | 595.30 | 671.17 | +12.75% | 0.00 |
| Dwarf | 582.39 | 656.28 | +12.69% | 0.00 |
| Night Elf | 591.15 | 666.37 | +12.72% | 0.00 |
| High Order | 586.61 | 661.36 | +12.74% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -4s: [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447).
- -2s: `{"move": {"rangeFromTarget": {"const": {"val": "5"}}}}`.

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Raptor Strike](https://www.wowhead.com/forever/spell=14266).
3. Cast [Strider Kick](https://www.wowhead.com/forever/spell=1317257).
4. Cast [Mongoose Bite](https://www.wowhead.com/forever/spell=14271).
5. Cast [Immolation Trap](https://www.wowhead.com/forever/spell=14305).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| Auto-attack (tag 2) | 135.49 |
| Auto-attack (tag 1) | 123.00 |
| Cat: Auto-attack (tag 1) | 83.70 |
| [Raptor Strike](https://www.wowhead.com/forever/spell=14266) | 77.93 |
| [Mongoose Bite](https://www.wowhead.com/forever/spell=14271) | 72.96 |
| [Strider Kick](https://www.wowhead.com/forever/spell=1317257) | 56.66 |
| [Immolation Trap](https://www.wowhead.com/forever/spell=14305) | 52.03 |
| Auto-attack (tag 3) | 40.18 |

## Warrior — Arms

**Talents:** 34/17/0 · `20305213132515001-550500000002`

[Requests, results and search evidence](../artifacts/optimization/arms.json)

- The retained 34/17/0 allocation moves Deflection 2/2, Improved Slam 2/2 and one Improved Heroic Strike point into Unbridled Wrath 5/5 and Improved Execute 2/2. Mortal Strike’s Sweeping Strikes prerequisite and Anger Management’s Tactical Mastery prerequisite remain funded.
- A coupled talent/APL experiment permits untalented Slam at 15 rage while reallocating its two talent points. It preserves continuous autos; merely removing Improved Slam without replacing the old cast-time condition would disable Slam.
- Add the already-talented Spearing Strike as a low-priority action at 15 rage. Rend, Overpower, Mortal Strike, execute handling, the 70-rage Heroic Strike condition and final-window Recklessness remain.
- The three execute/Slam allocations were compared explicitly, followed by another legal talent search and rotation screen. All ten races passed independent 5,000-iteration comparisons. No equipment or external effect changed.
- Rage generation and combat-table uncertainties remain T01/T02/T06; these results use the same inherited rules as the frozen baseline.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 626.67 | 654.23 | +4.40% | 0.00 |
| Tauren | 625.34 | 652.84 | +4.40% | 0.00 |
| Troll | 623.54 | 650.57 | +4.33% | 0.00 |
| Undead | 635.80 | 663.17 | +4.30% | 0.00 |
| Windshaper | 623.18 | 650.72 | +4.42% | 0.00 |
| Human | 619.36 | 646.66 | +4.41% | 0.00 |
| Dwarf | 619.32 | 646.97 | +4.47% | 0.00 |
| Night Elf | 624.73 | 651.19 | +4.23% | 0.00 |
| Gnome | 622.60 | 648.87 | +4.22% | 0.00 |
| High Order | 623.18 | 650.72 | +4.42% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Cast [Berserker Stance](https://www.wowhead.com/forever/spell=2458) when Time remaining ≤ 15s.
2. Cast [Recklessness](https://www.wowhead.com/forever/spell=1719) when Time remaining ≤ 15s.
3. Use ready automatic cooldowns.
4. Cast [Battle Stance](https://www.wowhead.com/forever/spell=2457) when (NOT Time remaining ≤ 15s AND NOT [Recklessness](https://www.wowhead.com/forever/spell=1719) active).
5. Cast [Overpower](https://www.wowhead.com/forever/spell=11585).
6. Cast [Rend](https://www.wowhead.com/forever/spell=11574) when NOT [Rend](https://www.wowhead.com/forever/spell=11574) DoT active.
7. Cast [Execute](https://www.wowhead.com/forever/spell=20662).
8. Cast [Mortal Strike](https://www.wowhead.com/forever/spell=21553).
9. Cast [Whirlwind](https://www.wowhead.com/forever/spell=1680).
10. Cast [Slam](https://www.wowhead.com/forever/spell=11605) when Rage ≥ 15.
11. Cast [Heroic Strike](https://www.wowhead.com/forever/spell=25286) when (Rage ≥ 70 AND NOT `{"isExecutePhase":{"threshold":"E20"}}`).
12. Cast [Spearing Strike](https://www.wowhead.com/forever/spell=1310222) when Rage ≥ 15.

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 148.02 |
| [Overpower](https://www.wowhead.com/forever/spell=11585) | 111.49 |
| [Execute](https://www.wowhead.com/forever/spell=20662) | 102.55 |
| [Mortal Strike](https://www.wowhead.com/forever/spell=21553) | 102.14 |
| Auto-attack (tag 3) | 64.33 |
| [Heroic Strike](https://www.wowhead.com/forever/spell=25286) | 52.19 |
| [Slam](https://www.wowhead.com/forever/spell=11605) | 25.60 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 17.33 |

## Warlock — Destruction

**Talents:** 13/5/33 · `2521000003-0005-0050355103101351`

[Requests, results and search evidence](../artifacts/optimization/destruction.json)

- The retained 13/5/33 build replaces Improved Shadow Bolt 5/5 and Fel Vitality 3/3 with four additional Suppression points, Pandemic 3/3 and Cataclysm 3/3. Incinerate remains the filler; no Shadow Bolt debuff is credited without its talent.
- Refresh Immolate ahead of Conflagrate and other DoTs, and use Conflagrate whenever its available Immolate allows it instead of waiting for the last two seconds. The existing curse, Corruption and Shadowburn actions remain.
- All five races passed independent 5,000-iteration comparisons. Suppression returns paid-hit budget; the gain is not entirely from rotation changes. Pet choice, gear and consumes remain fixed.
- Incinerate’s Immolate bonus remains base-damage-only pending T08. Life Tap mana sustain assumes healing support.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 599.12 | 645.39 | +7.72% | 0.00 |
| Troll | 594.28 | 639.30 | +7.58% | 0.00 |
| Undead | 605.23 | 650.98 | +7.56% | 0.00 |
| Human | 594.60 | 640.40 | +7.70% | 0.00 |
| Gnome | 601.96 | 646.15 | +7.34% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 75%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 60%).
3. Use ready automatic cooldowns.
4. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871) when Time remaining ≤ 1.5.
5. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
6. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 10%.
7. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
8. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) when [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309,"rank":8}}}`.
9. Cast [Conflagrate](https://www.wowhead.com/forever/spell=18932) when [Immolate](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ 15s.
10. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
11. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`).
12. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) when [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311,"rank":7}}}`.
13. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871).
14. Cast [Incinerate](https://www.wowhead.com/forever/spell=1293813).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Incinerate](https://www.wowhead.com/forever/spell=1293813) | 195.94 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 92.43 |
| Succubus: Auto-attack (tag 1) | 88.98 |
| [Conflagrate](https://www.wowhead.com/forever/spell=18932) | 77.05 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 66.09 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 61.35 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 40.64 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 11.18 |

## Priest — Shadow

**Talents:** 20/0/31 · `005300231303--505120501201300051`

[Requests, results and search evidence](../artifacts/optimization/shadow.json)

- The existing 20/0/31 talents are retained. Legal one-point alternatives did not pass independent validation.
- Inner Focus is explicitly paired with Mind Blast rather than sitting behind the ordinary Devouring Plague cast. The retained dots, spell ranks and full Mind Flay channel remain unchanged.
- Dot omissions, end-of-fight Pain thresholds, channel interruption conditions and Inner Focus targets were screened. The retained change passed 5,000-iteration comparisons on all six Priest races.
- Priest spellbook availability still needs in-game confirmation (T05); Mind Flay haste remains a separate unresolved interaction (T17).
- Channel readiness conditions were retested after fixing spellCanCast inside interruptIf. Neither Mind Blast-only nor DoT-aware interruption improved the retained full-channel rotation; the earlier no-op interrupt trials are superseded by shadow-channel-fixed.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Troll | 635.98 | 637.21 | +0.19% | 0.00 |
| Undead | 644.63 | 646.32 | +0.26% | 0.00 |
| Human | 636.31 | 637.98 | +0.26% | 0.00 |
| Dwarf | 641.62 | 643.28 | +0.26% | 0.00 |
| Night Elf | 641.52 | 643.15 | +0.25% | 0.00 |
| Gnome | 641.23 | 642.90 | +0.26% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -1s: [Shadowform](https://www.wowhead.com/forever/spell=15473).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) when NOT [Devouring Plague (rank 6)](https://www.wowhead.com/forever/spell=19280) DoT active.
3. Cast [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) when (NOT [Shadow Word: Pain (rank 8)](https://www.wowhead.com/forever/spell=10894) DoT active AND Time remaining ≥ 10).
4. `{"condition":{},"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10947}}}]}}` when `{}`.
5. Cast [Mind Blast (rank 9)](https://www.wowhead.com/forever/spell=10947).
6. Cast [Mind Flay (rank 6)](https://www.wowhead.com/forever/spell=18807).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Mind Flay](https://www.wowhead.com/forever/spell=18807) | 307.45 |
| [Mind Blast](https://www.wowhead.com/forever/spell=10947) | 131.72 |
| [Shadow Word: Pain](https://www.wowhead.com/forever/spell=10894) | 115.49 |
| [Devouring Plague](https://www.wowhead.com/forever/spell=19280) | 80.90 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 10.76 |

## Warlock — Demonic Pact

**Talents:** 2/31/18 · `011-0005003221220311351-0550005003`

[Requests, results and search evidence](../artifacts/optimization/demonology.json)

- The retained 2/31/18 build moves Improved Life Tap 1/2 into Suppression 1/5. The paid-hit model converts the saved hit requirement into offensive budget; this is not free hit.
- The existing APL and Succubus/Voidwalker Pact configuration remain unchanged. DoT ordering, late-fight cutoffs, Life Tap timing, Soul Fire under Decimation and channel fillers were screened; no rotation gain survived independent confirmation.
- Five races passed 5,000-iteration comparisons. Demonic Pact sacrifice/pet persistence remains subject to T09; pet settings were not changed. Life Tap sustain assumes healing support.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 631.10 | 635.03 | +0.62% | 0.00 |
| Troll | 627.70 | 631.18 | +0.56% | 0.00 |
| Undead | 636.53 | 639.32 | +0.44% | 0.00 |
| Human | 637.56 | 641.24 | +0.58% | 0.00 |
| Gnome | 638.10 | 640.85 | +0.43% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -1s: [Soul Link](https://www.wowhead.com/forever/spell=19028).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 75%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 60%).
3. Use ready automatic cooldowns.
4. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
5. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 10%.
6. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
7. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
8. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`).
9. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) when [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311,"rank":7}}}`.
10. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) when [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309,"rank":8}}}`.
11. Cast [Shadow Bolt (rank 9)](https://www.wowhead.com/forever/spell=25307).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 316.08 |
| Succubus: Auto-attack (tag 1) | 91.82 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 75.63 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 69.08 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 67.69 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 13.06 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 5.75 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 2.13 |

## Warlock — Affliction

**Talents:** 31/0/20 · `2525000013520105--0550015103`

[Requests, results and search evidence](../artifacts/optimization/affliction.json)

- The retained 31/0/20 build raises Suppression 2→5, funded by Improved Corruption 3→2 and Improved Bane of Agony 2→0. The saved paid-hit cost is returned to the same offensive item budget.
- Stop refreshing Agony, Corruption, Siphon Life and Immolate with less than twelve seconds remaining. Doom and Shadow Bolt remain in the rotation; pet and consumables are unchanged.
- Two legal Wrack allocations were tested, including Improved Drains 3/3. Full channels and DoT-aware clipping remained substantially behind Shadow Bolt after repairing channel-condition readiness. Wrack is implemented, but not selected for this profile.
- The current model ends Wrack’s 10% bonus when its channel ends; server-side scope and cancellation behavior remain T07. The rejected channel trials are preserved with their exact talent allocations and requests.
- All five races passed independent 5,000-iteration comparisons. Life Tap sustain assumes healing support.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 606.85 | 623.13 | +2.68% | 0.03 |
| Troll | 603.08 | 619.44 | +2.71% | 0.02 |
| Undead | 611.04 | 626.28 | +2.49% | 0.02 |
| Human | 613.66 | 629.53 | +2.59% | 0.03 |
| Gnome | 611.66 | 628.31 | +2.72% | 0.01 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -5s: [Amplify Curse](https://www.wowhead.com/forever/spell=18288).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 75%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 60%).
3. Use ready automatic cooldowns.
4. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871) when Time remaining ≤ 1.5.
5. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
6. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 10%.
7. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
8. Cast [Amplify Curse](https://www.wowhead.com/forever/spell=18288).
9. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
10. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when ((NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`) AND Time remaining ≥ 12s).
11. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) when ([Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311,"rank":7}}}` AND Time remaining ≥ 12s).
12. Cast [Siphon Life (rank 4)](https://www.wowhead.com/forever/spell=18881) when ([Siphon Life (rank 4)](https://www.wowhead.com/forever/spell=18881) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":18881,"rank":4}}}` AND Time remaining ≥ 12s).
13. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) when ([Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309,"rank":8}}}` AND Time remaining ≥ 12s).
14. Cast [Shadow Bolt (rank 9)](https://www.wowhead.com/forever/spell=25307).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 271.79 |
| Succubus: Auto-attack (tag 1) | 81.07 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 80.35 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 73.75 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 68.33 |
| [Siphon Life](https://www.wowhead.com/forever/spell=18881) | 32.70 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 13.96 |
| Succubus: [Lash of Pain](https://www.wowhead.com/forever/spell=11780) | 3.84 |

## Shaman — Enhancement

**Talents:** 20/31/0 · `05043305-055030031005102051`

[Requests, results and search evidence](../artifacts/optimization/enhancement.json)

- The retained 20/31 talents include Shamanistic Focus and Reverberation 4/5. The flat 45% shock/shield discount is used; no Classic crit-triggered Shamanistic Focus effect is added.
- Stormstrike takes priority over Maelstrom casts. At three or more Maelstrom stacks, cast Chain Lightning when ready, then Lightning Bolt; ordinary casts no longer reset or suspend melee swings.
- This rotation was re-screened after correcting the inherited swing resets. The comparison uses a same-engine baseline, so the mechanics correction itself is not reported as a talent/APL gain.
- All five races passed 5,000-iteration comparisons. Maximum mana-limited time is below three seconds; weapon, imbues, gear, buffs and consumes are unchanged.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 598.33 | 625.46 | +4.53% | 2.66 |
| Tauren | 596.64 | 623.66 | +4.53% | 2.79 |
| Troll | 596.11 | 623.00 | +4.51% | 2.55 |
| Windshaper | 598.71 | 625.87 | +4.54% | 2.40 |
| Dwarf | 600.64 | 627.87 | +4.53% | 2.56 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -4.5s: [Strength of Earth Totem (rank 4)](https://www.wowhead.com/forever/spell=10442).
- -3s: [Grace of Air Totem (rank 2)](https://www.wowhead.com/forever/spell=10627).
- -1.5s: [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438).

### Rotation priorities

1. Cast [Strength of Earth Totem (rank 4)](https://www.wowhead.com/forever/spell=10442) when Earth totem time remaining ≤ 0s.
2. Cast [Grace of Air Totem (rank 2)](https://www.wowhead.com/forever/spell=10627) when Air totem time remaining ≤ 0s.
3. Use ready automatic cooldowns.
4. Cast [Stormstrike (rank 1)](https://www.wowhead.com/forever/spell=17364).
5. Cast [Chain Lightning](https://www.wowhead.com/forever/spell=10605) when [Maelstrom Weapon](https://www.wowhead.com/forever/spell=408505) stacks ≥ 3.
6. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when [Maelstrom Weapon](https://www.wowhead.com/forever/spell=408505) stacks ≥ 3.
7. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Fire totem time remaining ≤ 0s AND Time remaining ≥ 5s).
8. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active.
9. Cast [Earth Shock (rank 7)](https://www.wowhead.com/forever/spell=10414).

### Damage breakdown — Dwarf

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 161.25 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 85.92 |
| Auto-attack (tag 2) | 80.09 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 60.75 |
| [Earth Shock](https://www.wowhead.com/forever/spell=10414) | 48.29 |
| [Stormstrike](https://www.wowhead.com/forever/spell=17364) | 45.56 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 43.50 |
| [Windfury Weapon](https://www.wowhead.com/forever/spell=16362) | 33.94 |

## Mage — Arcane

**Talents:** 33/3/15 · `053005023100311531-03-005500032`

[Requests, results and search evidence](../artifacts/optimization/arcane.json)

- Arcane Focus 3→5 is funded by Improved Channeling 5→3. The allocation remains 33/3/15; the stationary benchmark has no spell pushback.
- The extra talent hit reduces the paid hit-budget deduction. Gear records, enchants and external buffs remain unchanged.
- Arcane Blast stack thresholds, mana reserves, Frostbolt/Fireball/Frostfire/Missiles resets and single-spell fillers were screened. The original APL remained best among the validated choices.
- All six Mage races passed independent 5,000-iteration comparisons. Mana-limited time is 0.11–0.24 seconds; High Order Read Ley Line regeneration remains unmodeled pending T04.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 561.59 | 570.85 | +1.65% | 0.19 |
| Troll | 560.15 | 569.15 | +1.61% | 0.20 |
| Undead | 573.73 | 582.76 | +1.57% | 0.21 |
| Human | 567.97 | 576.89 | +1.57% | 0.21 |
| Gnome | 564.92 | 573.47 | +1.51% | 0.11 |
| High Order | 562.46 | 571.18 | +1.55% | 0.21 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 15% AND Time remaining > 20s).
3. Cast [Arcane Missiles (rank 8)](https://www.wowhead.com/forever/spell=25345) when [Missile Barrage](https://www.wowhead.com/forever/spell=44404) active.
4. Cast [Frostbolt (rank 11)](https://www.wowhead.com/forever/spell=25304) when ([Arcane Blast](https://www.wowhead.com/forever/spell=30451) stacks ≥ 4 OR Mana fraction < 20%).
5. Cast [Arcane Blast](https://www.wowhead.com/forever/spell=30451).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Arcane Missiles](https://www.wowhead.com/forever/spell=25345) | 281.24 |
| [Arcane Blast](https://www.wowhead.com/forever/spell=30451) | 258.42 |
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 28.05 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 15.04 |

## Warlock — DS/Ruin

**Talents:** 21/11/19 · `252200001351-0025003001-0550005103`

[Requests, results and search evidence](../artifacts/optimization/ds_ruin.json)

- The retained 21/11/19 build raises Suppression 1→5, funded by Malediction 5→2 and Improved Bane of Agony 1→0. This gain includes the lower paid-hit cost, with no alteration to catalog gear.
- Retain the existing sacrifice, curse/DoT priorities and Shadow Bolt filler. Rotation variants did not clear the independent validation margin.
- All five races passed 5,000-iteration comparisons, with negligible failed mana checks. The sacrificed Imp remains absent in combat; Life Tap sustain assumes healing support.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 553.88 | 570.05 | +2.92% | 0.00 |
| Troll | 549.58 | 566.27 | +3.04% | 0.00 |
| Undead | 559.52 | 575.83 | +2.91% | 0.01 |
| Human | 560.29 | 576.93 | +2.97% | 0.00 |
| Gnome | 558.43 | 575.01 | +2.97% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -5s: [Demonic Sacrifice](https://www.wowhead.com/forever/spell=18788).
- -5s: [Amplify Curse](https://www.wowhead.com/forever/spell=18288).

### Rotation priorities

1. Cast OtherActionPotion when (Time remaining ≥ 15s AND Mana fraction ≤ 75%).
2. Cast [Item 12662](https://www.wowhead.com/forever/item=12662) when (Time remaining ≥ 15s AND Mana fraction ≤ 60%).
3. Use ready automatic cooldowns.
4. Cast [Shadowburn (rank 6)](https://www.wowhead.com/forever/spell=18871) when Time remaining ≤ 1.5.
5. Cast [Searing Pain (rank 6)](https://www.wowhead.com/forever/spell=17923) when Time remaining ≤ 3.5.
6. Cast [Life Tap (rank 6)](https://www.wowhead.com/forever/spell=11689) when Mana fraction < 10%.
7. Cast [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) when NOT [Curse of the Elements](https://www.wowhead.com/forever/spell=1311680) active.
8. Cast [Amplify Curse](https://www.wowhead.com/forever/spell=18288).
9. Cast [Bane of Doom](https://www.wowhead.com/forever/spell=603) when (Time remaining ≥ 61s AND NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active).
10. Cast [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) when (NOT [Bane of Doom](https://www.wowhead.com/forever/spell=603) DoT active AND [Bane of Agony (rank 6)](https://www.wowhead.com/forever/spell=11713) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":11713,"rank":6}}}`).
11. Cast [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) when [Corruption (rank 7)](https://www.wowhead.com/forever/spell=25311) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25311,"rank":7}}}`.
12. Cast [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) when [Immolate (rank 8)](https://www.wowhead.com/forever/spell=25309) DoT time remaining ≤ `{"spellCastTime":{"spellId":{"spellId":25309,"rank":8}}}`.
13. Cast [Shadow Bolt (rank 9)](https://www.wowhead.com/forever/spell=25307).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Shadow Bolt](https://www.wowhead.com/forever/spell=25307) | 323.21 |
| [Corruption](https://www.wowhead.com/forever/spell=25311) | 88.46 |
| [Bane of Doom](https://www.wowhead.com/forever/spell=603) | 79.52 |
| [Immolate](https://www.wowhead.com/forever/spell=25309) | 65.69 |
| [Bane of Agony](https://www.wowhead.com/forever/spell=11713) | 15.66 |
| [Shadowburn](https://www.wowhead.com/forever/spell=18871) | 2.68 |
| [Searing Pain](https://www.wowhead.com/forever/spell=17923) | 1.71 |

## Mage — Frost

**Talents:** 11/3/37 · `050005001-03-0555003321001301251`

[Requests, results and search evidence](../artifacts/optimization/frost.json)

- No tested change passed the independent improvement threshold; the 11/3/37 baseline and its APL are retained.
- Frostbolt and Frostfire ranks, Ice Lance on any Fingers of Frost proc versus two stacks, and Ice Lance omission were screened on Orc and Gnome. Both races also completed legal one-point talent searches without a validated gain.
- The unchanged baseline already has 5,000-iteration results for all six supported Mage races. Eureka/channel interactions and Frostfire/Tier 1 behavior remain listed for in-game review; unverified effects were not added.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 568.14 | 567.86 | -0.05% | 0.12 |
| Troll | 569.05 | 568.89 | -0.03% | 0.12 |
| Undead | 573.49 | 572.95 | -0.09% | 0.12 |
| Human | 573.81 | 573.43 | -0.07% | 0.10 |
| Gnome | 570.65 | 570.72 | +0.01% | 0.05 |
| High Order | 570.28 | 570.09 | -0.03% | 0.12 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Cast [Item 8008](https://www.wowhead.com/forever/item=8008) when Mana fraction ≤ 80%.
2. Use ready automatic cooldowns.
3. Cast [Evocation](https://www.wowhead.com/forever/spell=12051) when (Mana fraction < 15% AND Time remaining > 25s).
4. Cast [Ice Lance](https://www.wowhead.com/forever/spell=30455) when [Fingers of Frost](https://www.wowhead.com/forever/spell=44543) active.
5. Cast [Frostbolt (rank 11)](https://www.wowhead.com/forever/spell=25304).

### Damage breakdown — Human

| Action | DPS |
|---|---:|
| [Frostbolt](https://www.wowhead.com/forever/spell=25304) | 409.07 |
| [Ice Lance](https://www.wowhead.com/forever/spell=30455) | 164.35 |

## Druid — Feral

**Talents:** 9/34/8 · `050022-3521002023032213041-053`

[Requests, results and search evidence](../artifacts/optimization/feral.json)

- Naturalist 5/5 to 3/5; Nature's Reach 0/2 to 2/2. Client spell 16819 has separate unmasked physical/spell hit effects (auras 54/55), each worth 4 percentage points at rank 2; only its range modifier has a Balance spell mask.
- The hit talent recovers offensive stats from the paid-hit allocation. Equipment records and the exchange model remain unchanged.
- Maintain Rip from 4 combo points and use Shred as the builder. Routine Ferocious Bite was worse; retain it only with at least 3 combo points in the final 4 seconds.
- Claw, Rake maintenance, 5-point Rip, Bite energy thresholds and alternative finish timings were tested. The retained change passed 5,000-iteration comparisons on all four Druid races.
- Revalidated after the smooth-Energy correction: all four races passed independent 5,000-iteration baseline comparisons. A fresh rotation screen found no further validated gain. Earlier talent and rotation trials used Classic-style ticks and are marked separately.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Tauren | 521.58 | 537.35 | +3.02% | 0.00 |
| Windshaper | 523.00 | 538.37 | +2.94% | 0.00 |
| Night Elf | 522.00 | 537.41 | +2.95% | 0.00 |
| High Order | 523.00 | 538.37 | +2.94% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Tiger's Fury](https://www.wowhead.com/forever/spell=9846) when Energy ≤ 40.
3. Cast [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) when (Time remaining ≤ 4s AND Combo points ≥ 3).
4. Cast [Rip](https://www.wowhead.com/forever/spell=9896) when (Combo points ≥ 4 AND NOT [Rip](https://www.wowhead.com/forever/spell=9896) DoT active AND Time remaining ≥ 8s).
5. Cast [Shred](https://www.wowhead.com/forever/spell=9830).

### Damage breakdown — Windshaper

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 199.18 |
| [Shred](https://www.wowhead.com/forever/spell=9830) | 175.61 |
| [Rip](https://www.wowhead.com/forever/spell=9896) | 156.90 |
| [Ferocious Bite](https://www.wowhead.com/forever/spell=31018) | 4.16 |
| [Dragonbreath Chili (proc)](https://www.wowhead.com/forever/spell=15851) | 2.51 |

## Rogue — Mutilate

**Talents:** 31/20/0 · `0053031035140105-302303202014`

[Requests, results and search evidence](../artifacts/optimization/mutilate.json)

- The 31/20 allocation trades Venom for Flawless Execution, and Lethality 5→3 plus Vigor 2→0 for Improved Poisons 4/5. Mutilate remains the builder.
- Eviscerate is used at three or more combo points. Thistle Tea is explicitly used at Energy ≤10. Slice and Dice and the existing Cold Blood condition remain.
- Venom, Rupture, builder pooling, Slice and Dice refreshes, finisher thresholds and Cold Blood timing were tested. A separate Gnome APL screen retained the same rotation.
- All nine races passed independent 5,000-iteration comparisons against the corrected Energy baseline. Seal Fate retains the client-supported 500-ms internal cooldown; no equipment or regeneration multiplier changed.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 482.28 | 517.24 | +7.25% | 0.00 |
| Troll | 476.16 | 518.21 | +8.83% | 0.00 |
| Undead | 500.76 | 535.99 | +7.04% | 0.00 |
| Windshaper | 476.82 | 518.67 | +8.78% | 0.00 |
| Human | 479.27 | 514.26 | +7.30% | 0.00 |
| Dwarf | 478.29 | 513.23 | +7.31% | 0.00 |
| Night Elf | 485.63 | 520.92 | +7.27% | 0.00 |
| Gnome | 481.34 | 523.70 | +8.80% | 0.00 |
| High Order | 476.82 | 518.67 | +8.78% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND `{"auraRemainingTime":{"auraId":{"spellId":6774,"rank":2}}}` < 3 AND Time remaining > 9)).
3. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
4. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when (Combo points ≥ 4 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active).
5. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 3.
6. Cast [Mutilate](https://www.wowhead.com/forever/spell=1241584).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 133.45 |
| Auto-attack (tag 2) | 77.03 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 65.63 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 56.30 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 51.18 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 45.56 |
| [Mutilate](https://www.wowhead.com/forever/spell=1241584) | 32.15 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 25.61 |

## Rogue — Combat

**Talents:** 18/33/0 · `005303103012-32003311201515231`

[Requests, results and search evidence](../artifacts/optimization/combat.json)

- Lethality 5→3 funds Improved Poisons 0→2; the allocation remains 18/33/0.
- Thistle Tea now uses an Energy threshold of 10 instead of a Classic tick-timing gate that is unreachable under smooth regeneration. Adrenaline Rush uses Energy ≤40.
- Eviscerate is used at three or more combo points; Cold Blood is explicitly aligned with that finisher instead of its automatic timing. The existing Sinister Strike swing-window gate and Slice and Dice priority remain.
- Finisher thresholds, builder pooling, Slice and Dice refreshes, Rupture and Cold Blood timing were tested. All nine races passed 5,000-iteration comparisons against the corrected Energy baseline; gear and consumes are unchanged.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 490.56 | 506.63 | +3.28% | 0.00 |
| Troll | 491.76 | 508.32 | +3.37% | 0.00 |
| Undead | 502.48 | 519.98 | +3.48% | 0.00 |
| Windshaper | 491.32 | 509.21 | +3.64% | 0.00 |
| Human | 495.62 | 511.82 | +3.27% | 0.00 |
| Dwarf | 490.00 | 506.13 | +3.29% | 0.00 |
| Night Elf | 493.12 | 509.40 | +3.30% | 0.00 |
| Gnome | 493.61 | 510.21 | +3.36% | 0.00 |
| High Order | 491.32 | 509.21 | +3.64% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -1s: [Blade Flurry](https://www.wowhead.com/forever/spell=13877).

### Rotation priorities

1. Cast [Item 7676](https://www.wowhead.com/forever/item=7676) when Energy ≤ 10.
2. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND `{"auraRemainingTime":{"auraId":{"spellId":6774,"rank":2}}}` < 3 AND Time remaining > 9)).
3. Cast [Adrenaline Rush](https://www.wowhead.com/forever/spell=13750) when Energy ≤ 40.
4. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
5. Cast [Cold Blood](https://www.wowhead.com/forever/spell=14177) when Combo points ≥ 3.
6. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when Combo points ≥ 3.
7. Cast [Sinister Strike (rank 8)](https://www.wowhead.com/forever/spell=11294) when (`{"autoTimeToNext":{"autoType":"MainHand"}}` > `{"autoSwingTime":{"autoType":"MainHand"}}` − 0.5 OR Time remaining < 6 OR Energy ≥ 79).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 151.45 |
| [Sinister Strike](https://www.wowhead.com/forever/spell=11294) | 110.42 |
| Auto-attack (tag 2) | 94.47 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 81.30 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 23.50 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 18.27 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 18.05 |
| Auto-attack (tag 3) | 10.36 |

## Priest — Smite

**Talents:** 31/17/3 · `515030031305001031-00505023002-003`

[Requests, results and search evidence](../artifacts/optimization/smite.json)

- The existing 31/17/3 talents are retained; screened one-point reallocations did not produce a validated improvement.
- Penance takes first spell priority, followed by Holy Fire maintenance. Remove Shadow Word: Pain, spend on maximum-rank Smite above 20 mana per remaining second, and use rank-3 Smite otherwise. Inner Focus remains paired with maximum-rank Smite.
- Removing Pain also removes the Shadow-school hit requirement. Holy Precision already caps the retained Holy spells, so the paid-hit model returns offensive budget. The gain is not solely a fixed-stat rotation improvement.
- The rank-3 filler uses the current client coefficient of 0.714; a server-side downranking penalty is still an open in-game check (T12). Penance uses complete channels, not a partial-channel timing exploit (T19).
- All six Priest races passed independent 5,000-iteration comparisons without mana-limited time or APL warnings.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Troll | 430.66 | 468.31 | +8.74% | 0.00 |
| Undead | 445.62 | 474.84 | +6.56% | 0.00 |
| Human | 431.22 | 467.33 | +8.37% | 0.00 |
| Dwarf | 431.57 | 468.56 | +8.57% | 0.00 |
| Night Elf | 431.99 | 469.40 | +8.66% | 0.00 |
| Gnome | 436.56 | 473.31 | +8.42% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -3s: [Holy Fire (rank 8)](https://www.wowhead.com/forever/spell=15261).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Penance](https://www.wowhead.com/forever/spell=1316995).
3. Cast [Holy Fire (rank 8)](https://www.wowhead.com/forever/spell=15261) when [Holy Fire (rank 8)](https://www.wowhead.com/forever/spell=15261) DoT time remaining < 3s.
4. `{"strictSequence":{"actions":[{"castSpell":{"spellId":{"spellId":14751}}},{"castSpell":{"spellId":{"spellId":10934,"rank":8}}}]}}`.
5. Cast [Smite](https://www.wowhead.com/forever/spell=10934) when Mana ≥ Time remaining × 20.
6. Cast [Smite](https://www.wowhead.com/forever/spell=598).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| [Smite](https://www.wowhead.com/forever/spell=10934) | 174.64 |
| [Penance](https://www.wowhead.com/forever/spell=1316995) | 108.30 |
| [Holy Fire](https://www.wowhead.com/forever/spell=15261) | 98.38 |
| [Smite](https://www.wowhead.com/forever/spell=598) | 85.85 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 7.66 |

## Rogue — Subtlety

**Talents:** 20/0/31 · `115303101104--5320003310013211501`

[Requests, results and search evidence](../artifacts/optimization/subtlety.json)

- The 20/0/31 allocation adds Improved Slice and Dice 3/3, Improved Poisons 4/5 and Vile Poisons 1/5. These are funded by Remorseless Attacks 2→1, Murder 2→0 and Cutthroat 5→0.
- Murder does not benefit the fixed neutral target and Remorseless Attacks has no kill trigger in this encounter. Cutthroat loses to sustained poison/maintenance talents in the tested five-minute rotation; this is not a general PvP or leveling recommendation.
- The original APL is retained. Finisher thresholds, Slice and Dice refreshes, Rupture, the stealth cycle and Ghostly Strike were tested; no rotation change passed independent validation.
- All nine races passed 5,000-iteration comparisons against the corrected Energy baseline. Hemorrhage and Thousand Cuts remain required specialization features; no gear or resource rules changed.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 418.64 | 449.20 | +7.30% | 0.00 |
| Troll | 414.34 | 444.31 | +7.23% | 0.00 |
| Undead | 431.70 | 463.08 | +7.27% | 0.00 |
| Windshaper | 415.03 | 445.08 | +7.24% | 0.00 |
| Human | 422.90 | 453.93 | +7.34% | 0.00 |
| Dwarf | 418.15 | 448.87 | +7.34% | 0.00 |
| Night Elf | 420.85 | 451.59 | +7.30% | 0.00 |
| Gnome | 420.48 | 450.88 | +7.23% | 0.00 |
| High Order | 415.03 | 445.08 | +7.24% | 0.00 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Rotation priorities

1. Cast [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) when ((Combo points ≥ 1 AND NOT [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining ≥ 6) OR (Combo points ≥ 5 AND `{"auraRemainingTime":{"auraId":{"spellId":6774,"rank":2}}}` < 3 AND Time remaining > 9)).
2. Cast [Vanish](https://www.wowhead.com/forever/spell=1856) when (Energy ≥ 60 AND Combo points ≤ 1 AND [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active AND Time remaining > 10).
3. Cast [Premeditation](https://www.wowhead.com/forever/spell=14183).
4. Cast [Ambush](https://www.wowhead.com/forever/spell=11269).
5. Cast [Rupture](https://www.wowhead.com/forever/spell=11275) when (Combo points ≥ 4 AND NOT [Rupture](https://www.wowhead.com/forever/spell=11275) DoT active AND Time remaining > 8).
6. Use ready automatic cooldowns when [Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active.
7. Cast [Spell 31016](https://www.wowhead.com/forever/spell=31016) when (Combo points ≥ 4 AND ([Slice and Dice (rank 2)](https://www.wowhead.com/forever/spell=6774) active OR Energy ≥ 79 OR Time remaining < 6)).
8. Cast [Hemorrhage](https://www.wowhead.com/forever/spell=16511).

### Damage breakdown — Undead

| Action | DPS |
|---|---:|
| Auto-attack (tag 1) | 141.71 |
| [Hemorrhage](https://www.wowhead.com/forever/spell=16511) | 109.91 |
| Auto-attack (tag 2) | 70.48 |
| [Rupture](https://www.wowhead.com/forever/spell=11275) | 47.32 |
| [Spell 25347](https://www.wowhead.com/forever/spell=25347) | 26.12 |
| [Instant Poison VI](https://www.wowhead.com/forever/spell=11340) | 21.81 |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198) | 18.50 |
| [Spell 31016](https://www.wowhead.com/forever/spell=31016) | 11.12 |

## Druid — Balance

**Talents:** 38/0/13 · `5232220115501351--505003`

[Requests, results and search evidence](../artifacts/optimization/balance.json)

- Genesis 4/5 to 2/5; Nature's Reach 0/2 to 2/2. Nature's Reach grants 4 percentage points of shared hit at rank 2.
- The gain includes recovery of offensive stats previously spent on paid hit normalization. Catalog equipment and the normalization exchange prices are unchanged.
- Retained the baseline Eclipse/DoT APL. Filler, DoT omission and priority-order variants failed independent validation.
- Read Ley Line regeneration remains unmodelled for High Order; see docs/in_game_checks.md T04.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Tauren | 442.36 | 458.28 | +3.60% | 0.02 |
| Windshaper | 439.86 | 456.25 | +3.73% | 0.03 |
| Night Elf | 442.00 | 458.37 | +3.70% | 0.03 |
| High Order | 439.86 | 456.25 | +3.73% | 0.03 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -2s: [Moonkin Form](https://www.wowhead.com/forever/spell=24858).

### Rotation priorities

1. Use ready automatic cooldowns.
2. Cast [Insect Swarm (rank 5)](https://www.wowhead.com/forever/spell=24977) when NOT [Insect Swarm (rank 5)](https://www.wowhead.com/forever/spell=24977) DoT active.
3. Cast [Moonfire (rank 10)](https://www.wowhead.com/forever/spell=9835) when NOT [Moonfire (rank 10)](https://www.wowhead.com/forever/spell=9835) DoT active.
4. Cast [Starfire (rank 7)](https://www.wowhead.com/forever/spell=25298) when [Eclipse](https://www.wowhead.com/forever/spell=48518) stacks ≥ 1.
5. Cast [Wrath (rank 8)](https://www.wowhead.com/forever/spell=9912).

### Damage breakdown — Night Elf

| Action | DPS |
|---|---:|
| [Starfire](https://www.wowhead.com/forever/spell=25298) | 259.68 |
| [Moonfire](https://www.wowhead.com/forever/spell=9835) | 71.94 |
| [Insect Swarm](https://www.wowhead.com/forever/spell=24977) | 63.83 |
| [Wrath](https://www.wowhead.com/forever/spell=9912) | 62.92 |

## Shaman — Stormcaller

**Talents:** 28/23/0 · `550032150010303-055030031004002`

[Requests, results and search evidence](../artifacts/optimization/stormcaller.json)

- Retained the baseline talents and APL. Rank/reserve, Chain Lightning, Flame Shock and totem/cooldown alternatives were tested on Orc and Tauren; none gave a validated DPS gain.
- A higher mana reserve and rank-6 fallback reduced mana-limited time but did not improve DPS in independent validation.
- Low-rank spell-power scaling remains an in-game check (T12); these results use the client-listed coefficients.
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 414.06 | 414.05 | -0.00% | 1.72 |
| Tauren | 414.26 | 414.42 | +0.04% | 1.70 |
| Troll | 411.57 | 411.33 | -0.06% | 1.69 |
| Windshaper | 412.86 | 413.46 | +0.15% | 1.75 |
| Dwarf | 414.39 | 414.48 | +0.02% | 1.70 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -2s: [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Rotation priorities

1. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Target count = 1 AND NOT [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) DoT active).
2. Cast [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) when (Target count ≥ 2 AND NOT [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) DoT active).
3. Use ready automatic cooldowns.
4. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active.
5. Cast [Chain Lightning (rank 4)](https://www.wowhead.com/forever/spell=10605) when ([Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active OR Target count ≥ 2).
6. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when Mana ≥ Time remaining × 10.
7. Cast [Lightning Bolt (rank 4)](https://www.wowhead.com/forever/spell=915).

### Damage breakdown — Dwarf

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 194.21 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=915) | 82.49 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 70.67 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 31.12 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 21.06 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 9.75 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=915) | 4.13 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 1.04 |

## Shaman — Elemental

**Talents:** 31/6/14 · `5502301500123031-0500001-053050001`

[Requests, results and search evidence](../artifacts/optimization/elemental.json)

- Retained the baseline talents and APL. Spell-rank, mana-reserve, Chain Lightning, Flame Shock, totem and cooldown-order alternatives produced no independently validated gain on Orc or Tauren.
- A 31/10/10 Intellect-focused alternative without Water Shield produced 386.85 DPS on Orc, below the roughly 395.6 DPS baseline, and was rejected.
- Lava Burst uses the model's Forever 20% Flame Shock damage bonus, not retail's guaranteed critical strike. Low-rank spell-power scaling remains an in-game check (T12).
- Revalidated at 5,000 iterations per race with continuous weapon-role autos and assumed general-haste Energy scaling. The prior retained talents and APL passed without changes; the comparison baseline uses this same engine.

### Results

| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |
|---|---:|---:|---:|---:|
| Orc | 395.60 | 395.42 | -0.05% | 0.61 |
| Tauren | 395.94 | 395.77 | -0.04% | 0.62 |
| Troll | 393.45 | 393.51 | +0.01% | 0.67 |
| Windshaper | 394.11 | 393.88 | -0.06% | 0.69 |
| Dwarf | 395.47 | 395.34 | -0.03% | 0.65 |

Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.

### Before the pull

- -2s: [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208).

### Rotation priorities

1. Cast [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) when (Target count = 1 AND NOT [Searing Totem (rank 6)](https://www.wowhead.com/forever/spell=10438) DoT active).
2. Cast [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) when (Target count ≥ 2 AND NOT [Magma Totem (rank 4)](https://www.wowhead.com/forever/spell=10587) DoT active).
3. Use ready automatic cooldowns.
4. Cast [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) when NOT [Flame Shock (rank 6)](https://www.wowhead.com/forever/spell=29228) DoT active.
5. Cast [Lava Burst (rank 3)](https://www.wowhead.com/forever/spell=1238300).
6. Cast [Chain Lightning (rank 4)](https://www.wowhead.com/forever/spell=10605) when ([Clearcasting (Elemental Focus)](https://www.wowhead.com/forever/spell=16246) active OR Target count ≥ 2).
7. Cast [Lightning Bolt (rank 10)](https://www.wowhead.com/forever/spell=15208) when Mana ≥ Time remaining × 35.
8. Cast [Lightning Bolt (rank 6)](https://www.wowhead.com/forever/spell=6041).

### Damage breakdown — Tauren

| Action | DPS |
|---|---:|
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 158.73 |
| [Lava Burst](https://www.wowhead.com/forever/spell=1238300) | 70.11 |
| [Flame Shock](https://www.wowhead.com/forever/spell=29228) | 66.88 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 43.45 |
| [Attack](https://www.wowhead.com/forever/spell=10436) | 30.65 |
| [Chain Lightning](https://www.wowhead.com/forever/spell=10605) | 14.98 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=6041) | 8.02 |
| [Lightning Bolt](https://www.wowhead.com/forever/spell=15208) | 2.19 |
