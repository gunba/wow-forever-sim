# Mage beta pass (17 September 2026)

Beta client `1.60.1.69893` against Classic Era `1.15.9.69722`, read with `tools/data_watch/spell_client.py`, the
talent curves with `tools/data_watch/trait_curve.mjs`, and `../beta/mage.json`.

## How the numbers were read

A rank's damage in the sim is the client's base plus `EffectRealPointsPerLevel` for every level from the rank's base
level up to its max level, capped at 60, rounded down for the low end and up for the high end. That rule reproduces the
sim's Classic tables for every max rank and most lower ranks against the Era client, so the Forever tables were built
the same way. About 25 lower ranks never matched Era under any rule (Fireball 5 and 9, Frostbolt 7 and 9, Scorch 1 and
4, Fire Blast 3, 5, 6, Pyroblast 2, 4, 6, Arcane Missiles 3, 5, 6, Arcane Explosion 2, 4, 5, Flamestrike 2, 4, 5, Blast
Wave 2); they now follow the rule too. None of them is cast at level 60.

Two splits the client and the sim make differently:
- **Blizzard and Flamestrike's burn** are area triggers in Forever that cast a separate damage spell each tick
  (Blizzard 1279976-1279980 and 1279949, Flamestrike 1279983-1279990). The sim keeps a dot on the parent spell, so its
  total is that spell's damage times the tick count (8 and 4), and its per tick coefficient is that spell's. Blizzard's
  parent spell also carries 0.03 on a dummy effect; the tooltip text builds the damage from the tick spell, so the 0.03
  is not used.
- **Fireball's dot** has 2, 3, 3 ticks at ranks 1-3 in both clients; the sim spreads the total over 4 ticks at every
  rank, as it did before.

## Spells changed (old -> new, level 60 rank unless noted)

| Spell | Change |
|---|---|
| Fireball (12 ranks) | rank 12 596-760 -> 425-541, dot 76 -> 60; every rank from 2 lower; coefficients ranks 1-4 .123/.271/.5/.793 -> .429/.571/.714/.857 |
| Frostbolt (11 ranks) | rank 11 515-555 -> 457-493; ranks 3+ lower; coefficients ranks 1-3 .163/.269/.463 -> .407/.489/.597 |
| Scorch (7 ranks) | rank 7 237-280 -> 166-197, about 30% lower at every rank |
| Fire Blast (7 ranks) | rank 7 446-524 -> 417-489; coefficients ranks 1-2 .204/.332 -> .429 |
| Pyroblast (8 ranks) | rank 8 716-890 -> 520-646, dot 364 (demo guess; Era 268) -> 212. The demo's 76 is rank 3. |
| Arcane Missiles (8 ranks) | rank 8 230 -> 209 a missile; coefficient .24 (ranks 1-2 .132/.204) -> .286 at every rank |
| Arcane Explosion (6 ranks) | rank 6 249-270 -> 238-259; rank 1 coefficient .111 -> .143 |
| Blizzard (6 ranks) | rank 6 1192 -> 1168 over 8 sec; .042 a tick unchanged |
| Flamestrike (6 ranks) | direct unchanged, rank 1 coefficient .134 -> .157; burn 340 -> 332, coefficient .02 (.017) -> .032 a tick |
| Blast Wave (5 ranks) | rank 5 462-544 -> 453-533 |
| Arcane Blast | guessed 435-505, 240 mana -> rank 5 (1239700) 364-424, 15% of base mana; .714 and 2.5 sec confirmed |
| Ice Lance | guessed 115-133 -> rank 6 (1240047) 136-161; 160 mana and 300% frozen bonus confirmed |
| Mage Armor | 30% -> 50% of mana regeneration while casting |

Unchanged in the client and left alone: cast times and mana costs of every ranked spell, Evocation, Counterspell,
Frost and Ice Armor, Arcane Power, Presence of Mind, Cold Snap, Combustion's 10% and 4 charges, mana gems, Ignite's
4 sec in two ticks (now spell 412538). Ice Barrier absorbs 7 less at every rank, which the sim does not model.

`ui/core/spells/mage.json` has no unreviewed entries left: the spells above are `forever` with tooltips, Frost and Ice
Armor are `classic`, Ice Lance and Fingers of Frost stay `assumed` for the reasons below.

## Checklist lines

### Mage section
- **Resolved** `sim/mage/fire_blast.go:39`, Wake of Fire rank 2: the client curve is 1000/2000 ms, so rank 2 takes
  2 sec off, not 1 sec.
- **Resolved** (already marked) Fingers of Frost rank 2: the client agrees, 15% at both ranks (400647), 2 charges at
  rank 2 (400670 stacks to 2), 15 sec.
- **Resolved** (already marked) Shatter: client curve 17/33/50.
- **Open** `sim/mage/talents.go:674`, whether a cast already in progress when the chill lands spends a Fingers of Frost
  charge. The client gives the aura (262, two charges) but not when the server consumes a charge relative to a cast
  that started earlier; that needs testing in game.

### Other TODOs in `sim/mage` settled from the client
- **Resolved** Arcane Subtlety spell penetration: curve 8/15, the sim took 8/16.
- **Resolved** Pyroblast periodic damage (see table).
- **Resolved** Missile Barrage duration: 15 sec (400589), with -50% channel, -100% cost, 40/20% procs as modelled.
- **Resolved, not in the checklist** Winter's Chill: the client stacks it once per talent point (`Stacks up to $s1
  times`, curve 1..5, debuff 12579 with 5 stacks) at 2% crit a stack. The sim had one stack; it now stacks.
- **Open** Ice Lance coefficient: the client's damage effect carries no spell power coefficient at all, the same way
  Arcane Shot, Serpent Sting and Consecration lost theirs. .143 stays, marked assumed.

### Talent icons
- **Open** `arcaneGeometry` still draws Flame Throwing's icon. Not a number the spell tables settle; the icon file id
  would come from the talent's spell in `SpellMisc`.

### Talents the sim does not read / Baseline ability changes
No mage lines in either section.

## Spellbook, Forever against Era (`--learned mage`)

New in Forever:
- Arcane Blast ranks 2-5 (1239696, 1239697, 1239699, 1239700) and Ice Lance ranks 1, 3-6 (1312002, 1240044-1240047):
  the talent spells gained ranks. The sim casts the level 60 rank of each.
- **Frostfire Bolt ranks 2 and 3** (1237312 level 50, 1237313 level 60; rank 1 401502 was already in Era's data from
  Season of Discovery). Rank 3: 3 sec cast, 370 mana, 270-314 plus growth, coefficient .814, a 40% slow, and 19 a tick
  every 3 sec for 9 sec; it counts as both Frost and Fire and is named by Missile Barrage, Hot Streak and Improved
  Fireball. **Implemented** in `sim/mage/frostfire_bolt.go`, including those explicit talent interactions.
  Its damage at level 60 has no rank-3 growth yet. All three ranks are available; rotation/talent optimization remains.
- Teleport: Dalaran, Study and Comprehend Scroll: no modelled combat use.
- Expansive Mind, Touch of the Grave and Eureka affect combat and are handled in core; see [the race audit](../forever_races.md).

Gone in Forever:
- Detect Magic.
- The per rank talent spells (Improved Fireball rank 2, Shatter ranks 2-5, Winter's Chill ranks 2-5 and so on) and
  Improved Blizzard ranks 2-3: talents moved to trait curves, so each talent keeps a single spell.
