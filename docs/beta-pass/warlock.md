# Warlock: beta client pass (17 September 2026)

Beta build `1.60.1.69893` against Classic Era `1.15.9.69722`, read with `tools/data_watch/spell_client.py` and
`tools/data_watch/trait_curve.mjs`. Every number below is the beta client's.

## How the rank tables were read

- Direct damage is each rank's value at the level it stops scaling (`SpellLevels.MaxLevel`, capped at 60) with
  `EffectRealPointsPerLevel`. That rule reproduces the sim's existing Classic tables exactly from the Era client for
  Shadow Bolt, Searing Pain, Shadowburn, Immolate, Firebolt and Lash of Pain, so the Forever tables are built the
  same way. Era rounds the low end down and the high end up; Forever's float ranges are rounded to nearest.
- Dots keep the sim's existing split: Corruption, Immolate and Drain Soul tables are totals (client per tick x ticks),
  Bane of Agony, Drain Life, Siphon Life and Rain of Fire are per tick. Dot coefficients are per tick in both the
  client and the sim, so they carry over unconverted.
- Where the sim did not match Era before this pass: Soul Fire rank 1 (628-789, Era 640-801), Conflagrate 18930
  (319-400, Era 326-407), Death Coil rank 2 (375, Era 391), Rain of Fire rank 2 (92, Era 96), Bane of Doom (the
  spell level coefficient of 1 never reached the dot) and Incinerate (380-440 scaled from a demo tooltip).

## Spells changed (93 ranks)

| Spell | Classic (sim before) | Forever beta |
|---|---|---|
| Shadow Bolt r1-r10 | 13-18 ... 482-538; coef .14/.299/.56 on r1-3 | 12-16, 25-31, 41-48, 57-64, 79-89, 101-113, 140-156, 188-210, 237-265, 253-283; coef .486/.629/.8 on r1-3, .857 above |
| Searing Pain r1-r6 | 38-47 ... 208-244; r1 coef .396 | 24-29, 34-40, 45-53, 62-73, 84-98, 107-126; .429 at every rank |
| Shadowburn r1-r6 | 91-104 ... 462-514 | 65-74, 81-91, 119-133, 147-164, 201-224, 259-288 |
| Soul Fire r1-r2 | 628-789, 715-894 | 344-430, 390-487 |
| Death Coil r1-r3 | 301/375/476, 430/495/565 mana | 285/375/460, 435/525/600 mana |
| Immolate r1-r8 | direct 11 ... 279, dot 20 ... 510; r1-2 coef .058/.125 and .037/.081 | direct 11, 21, 38, 64, 80, 116, 146, 158; dot 15, 30, 60, 95, 125, 190, 260, 275; .2 / .13 at every rank |
| Corruption r1-r7 | 40 ... 822 total; .08/.155/.167 per tick | 40, 65, 132, 168, 240, 342, 438; .2 per tick at every rank |
| Bane of Agony r1-r6 | 7 ... 87 per tick; .046/.077/.083 | 6, 10, 14, 21, 33, 46; .133 at every rank |
| Drain Life r1-r6 | 10 ... 71 per tick; r1 .078 | 10, 14, 22, 28, 39, 51; .1 at every rank |
| Drain Soul r1-r4 | 55/155/295/455 total; r1 .063 | 85/170/270/420; .1 at every rank |
| Siphon Life r1-r4 | 15/22/33/45 per tick | 11/19/29/41 |
| Rain of Fire r1-r4 | 42/92/155/226 per tick | 40/91/149/220 per tick, coefficient still .083 (each tick is its own damage spell, 1282380-1282385; the .03 the spell-data note lists is the channel's dummy effect) |
| Conflagrate | 4 ranks, 249-316 ... 447-557 | 6 ranks: new 1293817 (25) and 1293818 (32) below Classic's four; 88-111, 113-142, 134-170, 179-222, 220-273, 251-313; 100/130 mana on the new ranks |
| Incinerate | 29722, one rank, 380-440, 300 mana | 412758 (40), 1293812 (50), 1293813 (60): 100-114, 146-168, 201-233; 205/265/325 mana; 2.5 sec, .714, +25% on Immolate |
| Bane of Doom | 3200, no coefficient | 1742 with a 4.0 coefficient (now on the dot, where it applies) |
| Bane of Havoc | 300 mana | 5% of base mana |
| Fel Domination | 15 min cooldown | 5 min |
| Curse of the Elements | 1490/11721/11722 (Fire, Frost) plus Curse of Shadow 17862/17937 | Curse of Shadow removed; Curse of the Elements 1311677 (40) / 1311680 (50, 60) covers every magic school at 60/75 resistance and 8/10%. The shared aura now covers all six magic schools; the APLs cast 1311680 |
| Firebolt r1-r7 (Imp) | 7-10 ... 85-96 | 4-5, 7-9, 12-14, 18-19, 26-28, 36-39, 43-48 |
| Lash of Pain r1-r6 (Succubus) | 33 ... 99 | 16, 22, 30, 36, 43, 50 |

Life Tap's current description converts `(level-adjusted base + Spirit) × Improved Life Tap`
health into mana; it is not the inherited spell-power-scaled damage calculation.
Tier 1 multiplies the mana return only. The non-tanking DPS model retains its external-healing
abstraction; tanking pays the health as a resource cost, without damage procs.

Unchanged in the earlier client comparison: Demon Armor, Amplify Curse, Soul Link, the summons, Siphon Life's and Wrack's
coefficients, every cast time except Incinerate's, and every mana cost except Death Coil's and the new ranks.

## Checklist lines

**Resolved**

- `conflagrate.go:13` Conflagrate and Incinerate tables: see above. The BlizzCon tooltips (109-132, 125-140) are not
  what the client carries.
- `conflagrate.go:23` Shadow and Flame's chance not to consume Immolate: already 20% per point in the code; the curve
  (20/40/60/80/100) confirms it.
- `immolate.go:74` Aftermath: Immolate initial damage 10% per point confirmed; the Daze chance is 20% per point
  (curve 20..100), the 50% slow and 5 sec are single values (18118). The Daze changes no damage and stays unmodelled;
  the tree's ranks now read 20/40/60/80/100.
- `shadowburn.go:12` Shadowburn: all six ranks from the client (see above), 91-104 -> 65-74 at rank 1.
- `soul_fire.go:53` Decimation's Soul Fire cooldown reduction: 45% held flat -> 45% per point (90% at 2/2).
- `talents.go:168` Improved Drains: the client's talent is a flat 7/13/20% to Drain Life, Drain Soul and Wrack. The
  per-effect bonus the sim had read into it is Soul Siphon's, at 4/8/12% per other Affliction effect up to three
  (was 2% per point per effect, tripled below 20% health). Soul Siphon's faster drains (17/34/50%) and Drain Life
  healing penalty (10/15/20%) do not exist in the client and are gone. Wrack now takes both talents.
- `talents.go:359` Decimation: damage 3/6%, Soul Fire cast time 20/40% and cooldown 45/90% all scale; the 35%
  threshold and 10 sec buff are single values. The aura id 63165 -> 440873, the client's Decimation buff.
- `talents.go:476` Demonic Knowledge: 33/67/100% of level confirmed by the curve.
- `talents.go:806` Improved Shadow Bolt: 4% per point, the 12 sec debuff (17794) is a single value.
- `talents.go:916` Shadow and Flame damage: already 2% per point in the code; the curve (2..10) and 20 sec buffs
  (1293816, 426311) confirm it.
- `ui/core/talents/trees/warlock.json` unimplemented talents: every one of the ten reads the same per-rank values in
  the client curve as in the tree (Soul Harvesting 50/100, Improved Health Funnel 20/40, 15/30, 50/100, Demonic Aegis
  15/30, Improved Voidwalker 10/20/30, Improved Felhunter 10/20/30 and 2/4/6 sec, Destructive Reach 10/20, Molten Skin
  2..10, Pyroclasm 13/26, Fel Concentration and Intensity 23/47/70). Molten Skin and Demonic Aegis are now
  implemented (see below); the rest change no raid damage, healing or threat.
- Talents the sim does not read, **Demonic Aegis**: implemented, Demon Armor's armor and Shadow resistance +15% per
  point. It only matters to a tanking warlock, but the number is the client's and it is two lines.
- Talents the sim does not read, **Molten Skin**: implemented, 2% less damage taken per point.

**Partly resolved / open**

- `talents.go:414` Demonic Brand: threat 17/34/51% -> 17/33/50%, branded attacks 2 -> 2/4/6 per rank, brand 10 sec
  confirmed. Open: the pet hit is written `$<minDam> to $<maxDam>`, a formula the exported DB2 tables do not carry, so
  the BlizzCon 39-42 stays.
- Talents the sim does not read, **Improved Felhunter**: open by design. Tainted Blood, Devour Magic, Paranoia and
  Spell Lock are utility; the client's 10/20/30% and 2/4/6 sec change no raid damage.
- Soul Harvesting's tree tooltip carries a 10/20 sec duration the client writes as `<other spell d>`; not simulated.

## Also found and fixed (not on the checklist)

- **Demonic Sacrifice pairings are the reverse of Classic's**, and the sim had Classic's. The client's text and buffs
  agree: Imp -> +15% Shadow (18789, renamed Burning Shadow), Succubus -> +15% Fire (18791, Touch of Fire), Voidwalker ->
  2% mana every 4 sec (18792), Felhunter -> 3% health every 4 sec (18790). DS/Ruin sacrifices the Imp;
  the current Destruction benchmark keeps a live Succubus. Demonic Pact uses the Voidwalker sacrifice for mana.
- Cataclysm 3/6/9% -> 3/6/10%; Agonizing Flames 3/6/9% -> 3/7/10% (both ints the scaling test does not read).
- Master Demonologist's Voidwalker branch reduces Physical damage taken, not all damage.
- Nightfall also procs from Wrack.

## Spellbook (`--learned warlock` against Era)

New: Conflagrate ranks 1-2 (1293817, 1293818) and Incinerate ranks 1-3 (412758, 1293812, 1293813), both
implemented; Wrack (1316697) and Bane of Havoc (1225228), already in the sim; Curse of the Elements 1311676 (30),
1311677 (40), 1311680 (50), which carry ClassMask 0 on SkillLine 355 and so do not show in `--learned`, implemented;
racials Eureka!, Touch of the Grave, Expansive Mind (core); Journeyman Riding, Equip Transmog Outfit and a test spell
(4504 "Jeff Dummy 1").

Removed: Curse of Shadow (17862, 17937) and the Classic Curse of the Elements ids (1490, 11721, 11722), replaced
above; Dark Pact (18220, 18937, 18938; the sim never had it); Summon Felsteed and Summon Dreadsteed; the Master
Demonologist buff ids 23822-23844 (the sim still uses 23825 as a label); Drain Mana 18394; the Orc racial Command
(already handled in core); Season of Discovery's Haunt and tuning passive.

Renamed: Curse of Agony -> Bane of Agony, Curse of Doom -> Bane of Doom; stones and Detect Invisibility gained rank
labels.

## Shared integration follow-up

- Completed: `CurseOfElementsAura` now applies one aura over Fire, Frost, Arcane, Shadow, Nature and Holy.
  Each resistance changes once, and Curse of Shadow is ignored under Forever and removed from the picker.
- Completed: `CurseOfRecklessnessAura` applies -505 armor with no attack power bonus under Forever.
  The rank-4 tooltip metadata and the obsolete Curse of Shadow/Elements notes are corrected.
- See [shared effects](core.md) for runtime regression tests and remaining integration gaps.

## Oddities in the data

- Bane of Doom's 4.0 coefficient is the largest on any warlock spell.
- Improved Imp carries an undocumented third effect of -300/-700/-1000 (a dummy aura, likely Firebolt cast time in
  milliseconds) that its tooltip does not mention; not modelled.
- The Forever talent tooltips for Shadowburn, Conflagrate and Incinerate print 66, 95 and 97, matching none of the
  rank tables.
- The sacrificed Succubus attack bug is fixed: dismissal cancels its deferred pull-time auto-attack callback.
