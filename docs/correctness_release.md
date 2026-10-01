# Correctness review

The September 30 [Hunter/Druid deep dive](https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid) and recent evidence review led to 45 correction groups. The [implementation ledger](correctness_workplan.md) gives code/source scope and focused checks; [feature accounting](question_review/correctness_followup.json) covers all 144 article feature points and 20 recent findings. Announced changes, numeric script gaps and conflicting sources remain distinct.

## Reproduced results

The matrix contains 31 builds and 201 race/build profiles, each with a 5,000-iteration baseline, Tier-off, +10% and +50% item-stat/weapon-damage replay: 804 scenarios. Gear, enchants, talents, pet choices, consumes and the 300-second reference encounter remain fixed. The 16 Survival/Pet-Melee profiles migrate only their rotation to maintain their own finite Hunter’s Mark; an external mark cannot provide their Expose Prey benefit.

Raw requests/results and exact selectable web profiles are regenerated together. Ordinary MP5 and natural item/talent hit remain in force. The catalog is unchanged by sensitivity runs. Scaling columns use equal race weights and conservative marginal bounds; finite-range amplification is not proof of exponential growth or stat synergy.

The [comparison record](correctness/comparison.json) preserves the previous per-profile DPS/OOM values. This table averages available races equally; its changes combine corrections and are not isolated effect measurements.

| Build | Races | Previous DPS | Corrected DPS | Difference |
|---|---:|---:|---:|---:|
| affliction | 5 | 804.82 | 798.86 | -5.96 |
| arcane | 6 | 727.05 | 726.94 | -0.11 |
| arcane_frost | 6 | 747.83 | 747.74 | -0.09 |
| arms | 10 | 843.15 | 780.94 | -62.21 |
| balance | 4 | 705.66 | 710.87 | +5.21 |
| beast_mastery | 8 | 1011.57 | 947.03 | -64.54 |
| combat | 9 | 738.06 | 689.16 | -48.90 |
| demonology | 5 | 887.39 | 880.64 | -6.75 |
| destruction | 5 | 801.48 | 795.06 | -6.43 |
| ds_ruin | 5 | 766.62 | 766.53 | -0.10 |
| elemental | 5 | 543.30 | 543.30 | +0.00 |
| enhancement | 5 | 780.36 | 763.60 | -16.76 |
| feral | 4 | 720.06 | 693.15 | -26.90 |
| feral_tank_druid | 4 | 642.05 | 609.66 | -32.40 |
| fire | 6 | 725.84 | 725.76 | -0.09 |
| frost | 6 | 705.60 | 705.54 | -0.06 |
| fury | 10 | 896.04 | 827.23 | -68.81 |
| fury_2h | 10 | 799.36 | 739.88 | -59.48 |
| fury_sunder | 10 | 860.67 | 794.84 | -65.83 |
| marksmanship | 8 | 826.07 | 774.07 | -52.00 |
| mutilate | 9 | 719.50 | 675.63 | -43.88 |
| pet_melee | 8 | 978.34 | 908.73 | -69.61 |
| protection_paladin | 3 | 551.62 | 532.46 | -19.16 |
| retribution | 3 | 1145.26 | 1101.10 | -44.16 |
| retribution_physical | 3 | 1013.42 | 981.42 | -32.00 |
| shadow | 6 | 829.08 | 829.03 | -0.05 |
| smite | 6 | 546.73 | 546.67 | -0.06 |
| stormcaller | 5 | 547.26 | 547.26 | +0.00 |
| subtlety | 9 | 670.51 | 636.36 | -34.15 |
| survival | 8 | 954.25 | 884.93 | -69.32 |
| tank_warrior | 10 | 390.13 | 357.60 | -32.54 |

## Tank workloads

All 17 tank profiles have matched 300-second single-attacker and 90-second three-attacker checks with their original damage/healing assumptions. Each workload uses two disjoint 5,000-iteration ranges starting at 20291951 and 20341951. Pooled modeled death probability reaches 0.38%; it is not zero for every tank. DPS continues after modeled death, so survival and per-target threat are reported separately.

The [tank checks](correctness/tank_checks.json.gz) retain complete requests/results. Corrected baseline single-attacker replay matches the matrix. Earlier selection and confidence evidence is historical; equipment is replayed rather than newly selected. The earlier confirmation ranges overlap, so their nominal iteration totals must not be mistaken for independent evidence.

## Provenance and limits

The raw matrix identifies replay phases and engine/input hashes. The initial phase supplies unaffected classes; final Rogue rows add the sourced Instant Poison AP correction, and two Hunter rows supply migrated owned-Mark duties. Subsequent source-equivalent rebuilds change descriptive metadata and expose explicit rank-ID tables to the existing source scanner; they change no spell values or event rules. This keeps completed, unaffected replays without pretending every record came from one binary. Complete requests are the portable replay inputs.

Earlier baseline/profile artifacts and tank controls are [archived](correctness/history/manifest.json). Historical build DPS means are preserved; overlapping-seed confidence is marked conservative rather than independent. Old gear-selection gains do not establish current gains or an optimum.

The structured [questions register](uncertainties.json) contains 145 records: 36 resolved, 6 provisional answers, 13 implementation tasks, 69 evidence questions, 8 modeling choices and 13 outside the current scenarios. Low-rank experiments do not automatically settle level-60 behavior.

- Flametongue’s current numeric script is still missing. Its exclusion rules and learn levels are sourced, but a Flametongue/Grace versus Windfury DPS comparison would require invented damage/scaling. The level-30 rank-one test is actionable; the full combination requires later access (SHA-006).
- Savage Strikes and caster Faerie Fire retain article/client version conflicts. Registered Savage Strikes uses the captured crit-family data, not an additional unsourced damage modifier.
- Pet families, Hawk guardian behavior, form variance/conditional proc details, Light’s Vigil healing and shield threat/overlap ordering retain explicit qualifications.
- Instant Poison uses the sourced 0.5% AP term before multipliers; Deadly’s tick/stack placement remains unresolved. Classic behavior is unchanged.

Private chat, account identities and raw personal logs are not included in this evidence packet.

The spellbook refresh joins all 201 current profiles to captured 69893 acquisition tables: 151 candidate damage families, 54 absent from these samples. Six earlier apparent gaps are already registered in the expanded roster. This does not establish newer acquisition data, inactive talents or all pet families.
