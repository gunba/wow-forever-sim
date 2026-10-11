# October 11 audit

Baseline: published `9c05fc890` (client 1.60.1.70291 integration).

Evidence:
- Client 1.60.1.70338, the latest build, which changes no simulated table after 70291.
- The 70338 hotfix cache.
- The official class deep dives. The Priest and Warrior article had not been cited before.
- ClassicWoWCommunity/forever-bugs, including its "Not A Bug" and fix replies, and magey/forever-warrior.
- About 175,000 messages from Discord class test channels (WoWSims Forever, Fight Club, Classic Hunter Lodge, Warlock, Priest, Shadow Priest, Rogue, and the WoW Forever US/NA theorycrafting and class channels), plus a published Warrior threat test sheet and a Druid community FAQ.

Register dispositions are in [Simulator questions and coverage](uncertainties.md). Raw evidence is in
[`artifacts/correctness/october_11/`](../artifacts/correctness/october_11/). The previous results
are [archived](../artifacts/history/9c05fc890/manifest.json).

## How the client was checked

- **Spell parameters.** Every spell and aura registered by all 201 benchmark profiles was dumped
  from constructed simulations and compared with 70338's cast time, cooldown, GCD, cost and
  duration. The comparison first applies the client's own passive talent and Tier 1 spell
  modifiers (aura 107/108, by class mask, using trait curves).
- **Talent scope.** Each benchmark build was rebuilt without each talent it takes (583 variants).
  The spells whose modifiers changed were compared with the spells the talent's client masks cover.
  See [before](../artifacts/correctness/october_11/talent-scope-before.txt) and
  [after](../artifacts/correctness/october_11/talent-scope-after.txt); the remaining entries are
  spells that cannot crit or deal no damage.
- **Spell attributes and interrupt flags.** These were decoded with TrinityCore's names and
  semantics and diffed between Era 1.15.9 and Forever.
- **Hotfixes.** The 70338 `DBCache.bin` hotfixes touch only NPC and quest spells, plus a
  byte-identical Cold Blood row.
- **Stat conversions.** PlayerExpectedStat confirms every class's level-60 crit per Agility, spell
  crit per Intellect, base mana and health per Stamina.
- **Concurrency.** A race-detector run of all 31 builds in concurrent mode found no data races.

## Corrections

| Change | Evidence | Register |
|---|---|---|
| Untalented Slam blocks autos during its cast and restarts both swings. Improved Slam doesn't, and removes 3 s of cooldown at either rank. | Base ranks have InterruptFlags 15 and lack `DO_NOT_RESET_COMBAT_TIMERS`. Improved Slam overrides them with 1310196–1310200 (flags 7, attribute set). The deep dive and talent text agree. | WAR-019 |
| Forever has no hit suppression. | 7.04% misses at 1% hit against +3 targets (95% CI 6.31–7.85%); magey/forever-warrior#1 | CORE-001 |
| Forever finishers refund 80% of their Energy when avoided. | `DISCOUNT_POWER_ON_MISS` is on every player rank of Eviscerate, Rupture, Expose Armor, Rip and Ferocious Bite | CORE-003 |
| Siphon Life cannot crit. | It lacks `PERIODIC_CAN_CRIT`; forever-bugs#303 logged 219 ticks with no crits | WL-011 |
| The Hunter Tier 1 five-piece reduces Aimed Shot only. | The bonus's class mask | HUN-013 |
| Shield Wall and Intimidation trigger the GCD; Sweeping Strikes lasts 20 s. | Client recovery time and duration | CORE-018 |
| Arcane Shot scales with 10% RAP; Serpent Sting with 3% of current RAP per tick. Neither scales with spell power. | The client zeroed both coefficients; Hunter Lodge measurements; forever-hunter wiki | HUN-005 |
| Auto Shot fires during Aimed, Multi- and Sniper Shot casts. | forever-bugs#76 confirmation; wiki; Aimed Shot's interrupt flags | HUN-006 |
| Rend gains 2% of current AP per tick. Thunder Clap gains 7% AP, and its threat equals its damage. | Deep dive; Fight Club measurements | WAR-020 |
| Sunder threat gains 0.05 × AP. | Deep dive; published threat sheet | WAR-009 |
| Revenge gains 0.28 × AP, and its threat is 2.25 × damage with no flat bonus. | Six-point fit at rank 2; 0.20 measured at rank 1 | WAR-021 |
| Battle Shout generates no threat. | forever-bugs#64 (Not A Bug); measured 0 | WAR-022 |
| Shaman Weapon Mastery is not baseline. | Blizzard retracted the deep dive's claim (forever-bugs#280) | SHA-007 |
| Fire Nova does not trigger Elemental Devastation. | Blizzard: only direct harmful spell crits, no totems (forever-bugs#257) | SHA-008 |
| Twin Disciplines, Holy Precision, Holy Specialization, Mental Agility, Mortal Shots and Impale follow the client class masks. | Differential scope audit; tester reports for Mind Flay and Penance | CORE-019 |

## Selections on the corrected engine

- **Warrior rotations.** Fury and Fury (Sunder) drop Slam and maintain Rend. Arms casts Slam at 50+
  Rage. 2H Bloodthirst does both. Each was screened on every race and confirmed on an independent
  seed (WAR-023).
- **Hunter consumables.** Hunters use Flask of Distilled Wisdom and drop Greater Arcane Elixir
  (HUN-014).
- **Gear.** All 184 DPS profiles were re-searched. Affected builds were searched again after the
  class-mask corrections, and every change was confirmed on an independent seed pair:
  - 36 changes retained, 1 reverted.
  - The retained changes cover all Marksmanship profiles, most Beast Mastery profiles, and some
    Fury, Fury (Sunder), Combat, Subtlety, Feral and Pet/Melee profiles.
  - Tank gear keeps its guarded October 8 selections.

## Benchmark

All 1,005 scenarios were replayed: baseline, Tier 1 off, +10% and +50% equipment, and the previous
setups on the corrected engine. Seed 1793400001, 5,000 iterations each. All 68 tank workload checks
were also replayed. Builds that changed by at least 0.2%, averaged over races:

| Build | Published | Corrected engine, old setup | Released | Change |
|---|---:|---:|---:|---:|
| Protection Warrior | 377 | 506 | 506 | +34.3% |
| Marksmanship | 777 | 905 | 955 | +22.9% |
| Beast Mastery | 902 | 954 | 980 | +8.7% |
| Subtlety | 656 | 663 | 664 | +1.1% |
| Mutilate | 701 | 708 | 708 | +1.1% |
| Combat | 711 | 715 | 716 | +0.6% |
| Feral | 725 | 728 | 729 | +0.5% |
| Pet/Melee | 905 | 909 | 909 | +0.5% |
| Survival | 892 | 895 | 895 | +0.4% |
| Arms | 832 | 811 | 825 | -0.9% |
| Fury | 947 | 886 | 938 | -1.0% |
| Fury (Sunder) | 917 | 862 | 907 | -1.1% |
| Affliction | 824 | 814 | 814 | -1.2% |
| Shadow | 829 | 809 | 809 | -2.4% |
| 2H Bloodthirst | 829 | 791 | 799 | -3.7% |
| Enhancement | 812 | 761 | 761 | -6.3% |

The most sensitive provisional value is Revenge's rank-6 AP coefficient. Two low-rank measurements
differ (0.20 and 0.28), and the Protection Warrior row depends on it.
