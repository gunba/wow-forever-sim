# Tank simulations

Three Forever tank pages are available: [Protection Warrior](../ui/tank_warrior/), [Protection Paladin](../ui/protection_paladin/) and [Feral Bear](../ui/feral_tank_druid/). Their selectable defaults contain the rotation, legal level-60 talents, projected level-65 equipment, enchants, consumes, buffs and encounter. The role-specific Forever Tier 1 bonuses are enabled separately from equipped items. These are **tank scenarios**, not additions to the DPS ranking or evidence that one tank is universally better.

## Scenario and comparisons

The default is one level-63 Dragonkin with 3,731 armor, 805 attack power and a 3,000-damage base swing every two seconds, with a 33.33% damage spread and parry haste. Fight length is 300 seconds. External healing is modeled as 1,500 HPS in three-second intervals, with a six-second burst window. These swing and healing values are explicit stress-test assumptions, **not** observed boss or healer values. Three-target checks use 90 seconds, three enemies with 1,500-damage base swings each, and 2,500 external HPS. Each comparison uses two independent seeds, 5,000 iterations per arm, identical race, talents, gear and buffs. Only the Protection Paladin comparison also changes its potion from Greater Stoneshield to Major Mana. Threat per second (TPS), damage taken per second (DTPS), death rate and the *least-threatened target* matter more than personal DPS.

| Tank | Scenario | Previous TPS | Current TPS | Current DTPS | Least target: previous → current TPS |
|---|---|---:|---:|---:|---:|
| Protection Warrior (Human) | One boss, 300 s | 1,005 | 1,031 | 520 | — |
| Protection Warrior (Human) | Three enemies, 90 s | 1,163 | 1,178 | 772 | 37 → 37 |
| Protection Paladin (Human) | One boss, 300 s | 545 | 612 | 780 | — |
| Protection Paladin (Human) | Three enemies, 90 s | 840 | 1,101 | 1,137 | 121 → 202 |
| Feral Bear (Tauren) | One boss, 300 s | 1,450 | 1,450 | 755 | — |
| Feral Bear (Tauren) | Three enemies, 90 s | 1,703 | 1,656 | 1,095 | 107 → 210 |

Bear's new Swipe priority trades some *total* threat for much better threat on the least-threatened add. Warrior's three-target rotation remains **single-target focused**: its 37 TPS on the least-threatened add cannot hold an active damage dealer. The simulator does not yet select or cycle to the enemy with the least threat; tab-targeting is necessary in an actual pull. Do not use its summed three-target TPS to claim reliable pack control. The previous and current warrior rotations both keep Shield Block and attack-speed/power debuffs; removing these raises threat but measurably worsens mitigation. Full paired requests and results are in [`artifacts/tanks/`](../artifacts/tanks/).

Protection Warrior maintains Battle Shout when absent, keeps Shield Block active, refreshes Thunder Clap, Demoralizing Shout and five Sunder Armor stacks, uses available Revenge and Shield Slam, and spends surplus rage on Heroic Strike or Cleave. Health-gated Shield Wall and Last Stand remain defensive choices. Rend was tested as a one-target 12-second-or-longer maintenance action: approximately 15–17 casts lowered TPS by 17–25 in this boss setup, so it is omitted. The source-backed Shield Block lasts seven seconds or two blocks in Forever (five seconds or one block in Classic), and Demoralizing Shout's debuff generates no threat in Forever. See the current [Forever Warrior abilities](https://www.wowhead.com/forever/guide/classes/warrior/protection/tank-abilities) and [Shield Block](https://www.wowhead.com/forever/spell=2565/shield-block).

Protection Paladin keeps Holy Shield, Judgment and Holy Strike, uses Templar's Bulwark when appropriate, and selects a sustainable Consecration rank from target count and mana. The potion change reduces estimated single-target out-of-mana time from about 27 to 2 seconds. Templar's Bulwark is an eight-second maximum-health absorb, not unlimited damage reduction. **Seal of Fury and its shield, taunt and Improved Seal of Fury mana return are not implemented**; the preset uses Seal of Righteousness. This is a major reason not to treat its current relative threat or survival as final. See the [Forever Protection Paladin overview](https://www.wowhead.com/forever/guide/classes/paladin/protection/overview-pve-tank).

Feral Bear keeps Faerie Fire and Demoralizing Roar, favors Primal Bite on one enemy and uses Swipe against multiple enemies. The strength of Primal Bite's bonus threat and Swipe's inherited multiplier are not confirmed by Forever threat logs. The Bear Form health adjustment is applied safely when the aura expires.

## Data limits and validation

- The equipment is **modeled**, not confirmed obtainable tank gear. It uses the published level-65 projection library; the current pool lacks a defensively optimized set, especially for Paladin. These results must not be read as a fair cross-tank ranking.
- Protection Warrior's speed-based autoattack rage at level 60, Fury/Seal of Fury scripted effects, Paladin block/absorb order, Bear ability threat and multitarget targeting remain unverified. The level-20 beta can check Shield Block's two charges, Demoralizing Shout threat, Paladin's seal shield and Bulwark duration, and per-target Swipe/Primal Bite threat; level-60 gear and rage rates require later access.
- External healing keeps most, but not all, 300-second modeled fights alive. Current single-target Paladin and Bear still show roughly 1–2% and 1% deaths respectively under independent seeds. A dead tank cannot be judged from mean TPS alone. Neither 1,500 HPS nor the boss's 3,000-damage swing is a claim about an actual raid boss.
- The same source-backed three tank roles appear in the [Forever class/role listing](https://wowforever.gg/en/classes/). No Shaman tank has been asserted.
- All 17 supported race/tank combinations completed a native smoke run without rotation warnings. Fresh web defaults for all three tank pages reproduced native DPS, TPS and DTPS in WebAssembly with matched seeds.

To reproduce a selected build from actual web defaults, start the staged site locally and run `node tools/forever_tanks/export_defaults.mjs /tmp/tank-defaults`, then `python3 tools/forever_tanks/build_requests.py --settings-dir /tmp/tank-defaults --output-dir /tmp/tank-requests --iterations 5000 --seed 20261993`. Build the native runner with `go build -tags with_db -o /tmp/tank-bench ./tools/forever_tanks`, then pass a request using `/tmp/tank-bench -request /tmp/tank-requests/tank_warrior_1t_20261993.request.json -output /tmp/tank-result.json`. The paired files retain the exact baseline and chosen APLs and all inputs; the main DPS matrix is unchanged.
