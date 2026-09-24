# Gear comparisons

The 27 current builds were screened across all 174 available race/build combinations. The slot-coordinate search ran **87,044 trials**, including screening and slot validation. Every profile converged. Of the 174 loadouts, **19 changed** and survived an independent 5,000-iteration comparison; 155 kept their previous equipment and enchants.

Talents, rotations, pet choices and each existing build's consumables and external buffs were fixed throughout. The separate Physical Ret row begins with the corresponding race's proven physical Warrior gear, a legal Paladin relic and its own non-caster consumables. The encounter remains a 300-second level-63 Dragonkin with 3,731 starting armor, ordinary MP5, paid shared-hit accounting and complete role-specific Tier 1.

## Retained changes

The gains compare each race with its own **post-haste, pre-search** loadout, not with a different race or another Ret build.

| Build | Changed races | Confirmed DPS gain |
|---|---|---:|
| Retribution | Undead, Human, Dwarf | +12.31–13.52 |
| Physical Ret | Human, Dwarf | +7.61–16.87 |
| 2H Bloodthirst | Orc, Undead, Troll, Human, Dwarf, Gnome, Night Elf | +5.51–7.03 |
| Fury | Orc, Gnome | +5.27–7.21 |
| Enhancement | Windshaper, Dwarf | +4.60–5.51 |
| Beast Mastery | Tauren, Dwarf, Night Elf | +1.62–1.77 |

Physical Ret spends no points in Divine Intellect or Champion of the Light. It instead takes Reverence, Holy Conduit and the third Vindication rank while retaining 31 points in Retribution. The gear search rejects spell-power gear and Intellect-only items, but permits genuinely mixed physical/Intellect pieces; spell-power-focused enchants and caster flasks/elixirs are also excluded. The retained physical loadouts happen to have **zero item Intellect and zero item spell power** across all three races. They still receive the same external raid buffs and can deal Holy damage with Strength/AP-oriented gear. This row is a comparison archetype, not a claim that physical gearing beats mixed Ret.

The smallest retained gain, Night Elf Beast Mastery, is +1.62 DPS against a conservative ±0.81 DPS comparison bound. Full matched results include the complete player, profession, enchant, paid-hit ledger, source-limited pool and every simulated candidate. The final chart uses a third seed, so its displayed means can differ from the search comparisons.

## Method and limits

The main pool is the complete 706-item exact-ilvl-65 export, supplemented by verified available crafted, dungeon, PvP/vendor, cloak and trinket candidates. Class, faction, profession, unique and equipment restrictions are enforced. Old-raid and token rewards remain excluded. A lower item level is not disqualifying when a legal item's simulated DPS is higher.

Each item slot is screened against the current loadout, with legal weapon layouts compared together; enchants and the second profession are also compared. Every candidate pays for hit again, preserving Tauren's racial and the unchanged catalog. The strongest three improving screen candidates receive independent slot validation; passes continue until no supported improvement remains. The confirmation uses two 5,000-iteration pairs with independent seeds **20296317** and **20296383**. A gain must exceed `1.96 × (baseline SE + candidate SE)`; this conservative bound does not assume independent race/arm outcomes. All 19 changes passed.

The search identifies a local choice under the modeled rules, not a global optimum. Haste-channel and spell-GCD behavior follows a Mage community report pending direct timing logs. Warrior rage, pet inheritance, proc behavior and other limitations remain listed in the [gameplay checks](in_game_checks.md).

The subsequent restoration of Classic's half-second Hunter Auto Shot wind-up
changed Hunter rotations without another gear-coordinate search. Their gear
remains legal and is replayed with the corrected timing, but the earlier search
does **not** establish the best Hunter gear under this rule. Other classes'
gear selections are unaffected.

## Evidence

- [Per-race decisions and trial counts](../artifacts/gear_search/current/summary.json)
- [Complete candidate trials](https://gunba.github.io/wow-forever-sim/classic/review/gear_search/search.json.gz)
- [Independent comparison requests/results](https://gunba.github.io/wow-forever-sim/classic/review/gear_search/validation.json.gz)
- [Post-haste starting loadouts](https://gunba.github.io/wow-forever-sim/classic/review/gear_search/before.json.gz)
- [Previous published gear and results](../artifacts/gear_search/history/7fe1b5366/summary.json)
- [Replay procedure](../tools/forever_bench/README.md)
