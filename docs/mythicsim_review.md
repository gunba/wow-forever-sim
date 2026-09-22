# MythicSim engine comparison

## Revisions and scope

Compared [MythicSim revision `e6d999e927`](https://github.com/sage3648/mythicsim-forever-engine/tree/e6d999e9270c652db0820bee08c68f9f22a9079a)
with this project's published `3bf610850` release.

- The common ancestor with our imported engine is `a7622df37e2d046010cc207cf6da3419f5c9ba9b`.
- MythicSim's first independent commit, `38dc1ca717`, directly follows
  ElliotWood/Forever `0eb921c12cd3ce6501a915cd602726100bc46a78`.
- Its later merge `4000cfcf04` imports upstream through
  `2ac3dd8e7cdf3401f497b06ecf93889ccaee86c8`, including 48 commits after
  that initial fork point.
- Relative to that imported upstream revision, ten reachable commits remain,
  including merges and changelog updates. Their cumulative change affects
  15 files and is principally Seal of Fury, its tests, metadata and UI.
- The later upstream revisions `59255c67d5` and `12f81fc56d1`, which supplied
  several leads for our previous review, are not ancestors of the pinned
  MythicSim revision.

The independent Seal of Fury implementation, its follow-up combat handling,
tests, metadata and integration have been reviewed. The comparison also covers
the nine classes' talent-modifier migrations, core runtime changes, item-proc
changes and selected spell-data migrations inherited from upstream. Generated
tables were not treated as a new verification of every spell or its availability.

## Findings

| ID | Finding | Disposition |
|---|---|---|
| M01 | Seal of Fury, Judgement of Fury and the shield-break mana talent are implemented there but absent from our runtime. Our talent schema already exposes Improved Seal of Fury. | Confirmed implementation gap. Live Forever tooltips and captured client records support the ability and its basic effects; proc eligibility, shield replacement and some combat-table details still need tests. No implementation imported yet. |
| M02 | Their local shield pool consumes incoming damage and awards mana only when fully depleted. It excludes partial depletion, expiry, replacement and iteration reset. | Useful implementation and regression pattern, not proof of server behavior. Our generic shield helper still records shielding without a general finite-absorb consumption model; see H20. |
| M03 | Their imported upstream adds Hammer of the Righteous, spell 407632, with a shared cooldown with Holy Strike. Our runtime lacks it. | Availability lead, not yet a verified omission. A retained spell or class-skill row is insufficient to establish that this former rune ability is actually learnable. Confirm trainer access before adding it. |
| M04 | Imported upstream `35792e1fe5` checks positive net mana spending before estimating time to OOM. Our code divides first and handles a negative duration afterwards. | Numerical-hardening lead. Avoids division by zero or negative spending; not evidence of a DPS change or of a bad current published OOM duration. |
| M05 | The same commit replaces repeated textual action/resource searches in concurrent-result merging with typed-key indexes. | Performance lead, not a game-mechanic fix. Existing numerical comparisons do not justify attributing any DPS difference to it. |
| M06 | MythicSim uses speed-based Warrior rage at every level, including unverified off-hand handling, with 3.5 rage/weapon-second for one-hand and 4.5 for two-hand. | Intentional modeling difference. Low-level evidence supports normalization, but neither this implementation nor its tests establish level-60 coefficients. Our measured one-hand result is about 3.46, not a verified universal 3.5. T02 remains open. |
| M07 | The pinned engine retains persistent Hot Streak acceleration and the blanket one-third healing-to-damage conversion. | Do not import. Our release separately corrected both. Its Mage implementation also still omits Frostfire from Hot Streak's trigger list. |
| M08 | Upstream `e8d4270a60` and `692958311b` update legacy weapon-proc amounts and periods that remain old in our handlers. | Source-supported follow-up outside the selected equipment. Independent client checks are listed below. Not imported as a blanket patch; proc chances, equipment availability and other behavior are separate questions. |
| M09 | Both engines reset Earth Shock's individual timer when a freshly cast Lightning Shield gains its initial three charges. The separate shared shock cooldown remains intact. | Fixed here: remove the unsupported individual-timer reset. It originated in pre-Forever SoD `cd3e60f49`; `37d72dd6f` moved Rolling Thunder's handling but left this callback. Current shield effects do not grant that reset. Real-cast regressions detect the timer mutation for ranks 1 and 7; full-rotation checks do **not** demonstrate extra casts or a DPS exploit. |
| M10 | Upstream changes the Emerald Dragon Whelp's Acid Spit to a flat 374 while retaining assumed guardian stats. | Not verified as a complete formula. Our captured effect row has base 438.523071, variance 0.293333 and coefficient 1, not a literal flat 374. Guardian level scaling, distribution and the inherited 220 spell-damage stat need independent evidence. Dragon's Call is absent from our current catalog and selected gear. |
| M11 | Talent and aura modifiers across nine classes migrate callbacks and spell-code lists to static/dynamic SpellMods. | Primarily structural, not nine new sets of verified mechanics. The reviewed filters and additive/multiplicative operations mostly preserve earlier behavior, including unresolved assumptions. Do not replace our independently corrected handlers merely to match the new API. |
| M12 | Upstream adds APL action groups/variables, generated spell-table plumbing and a presimulation early exit. | Feature/performance work, not evidence that existing profiles deal incorrect damage. Current published APLs do not require these new constructs. |
| M13 | A suspected Demonic Sacrifice reset guard appears absent when reading only the owner's reset function. | Already correct: `Character.reset` resets the owner before its pets; `OnPetEnable` cancels sacrifice without Demonic Pact. A no-Pact full-run check retains the Succubus but records zero sacrifice uptime and zero sacrifice mana. The transient reset proc is not a combat benefit. |

### Inherited changes reviewed

- Talent modifiers: `afe03c6c5f` (Mage), `4065bef8a6` (Priest),
  `e50470f4c1` (Warlock), `625ef4d018` (Rogue), `c52f60acd0` (Shaman),
  `c9534eafab` (Druid), `804f484304` (Paladin), `5ab5fa6269` (Warrior),
  and `c18304461c` (Hunter).
- Runtime: `35792e1fe5`, `0f75974f28`, the SpellMod infrastructure and
  `c6156ba564`'s APL integration. No performance refactor was imported.
- Selected data migrations through `2ac3dd8e7c`: rank/cost/cooldown lookups,
  Mage damage variance, Rogue cooldowns and Paladin triggered-aura durations.
  Replacing literals with generated rows does not resolve coefficients,
  server scripts or learnability. This was not an exhaustive re-extraction
  of every generated spell row.
- Clearcasting (`d66773b84b`), Sanctity Aura (`7fe7a34532`), Expose Armor,
  Curse of Recklessness and Firestone overlap corrections already covered
  by this project's earlier review. Shadow Word: Death and its backlash
  remain H11/T37 rather than importing only its damage component.

### Legacy item effects

The captured 1.60.1.69893 `SpellEffect` rows independently support these
amounts and periods. Numbers are base effects, not final damage after
modifiers or proof of proc frequency.

| Item / effect spell | Current handler | Captured client effect |
|---|---|---|
| Barovian Family Sword / 18652 | 30 every 3 seconds | 49 every second |
| Blade of Eternal Darkness / 27860 | 100 damage, 100 mana | 112 damage, 104 mana |
| Ebon Hilt of Marduk / 18656 | 70 every second | 28 every 3 seconds |
| Frightskull Shaft / 18633 | 8 every 2 seconds, −50 Strength | 15 every 2 seconds, −92 attack power |
| Gravestone War Axe / 18289 | 55 every 3 seconds | 42 every 3 seconds |
| Hookfang Shanker / 13526 | 7 every 3 seconds | 13 every 2 seconds; −50 armor |
| Keris of Zul'Serak / 16528 | 8 every 2 seconds | 10 every second |
| Runeblade of Baron Rivendare / 17625 | 20 health every 5 seconds | 60 health every 5 seconds |
| The Hand of Antu'sul / 13532 | 7 direct damage | 42 direct damage |
| Venomspitter / 18203 | 7 every 2 seconds | 7 every second |
| Firebreather / 16413 | 70 direct plus 3 every 2 seconds | 77 direct plus 2 every second |

Only The Hand of Antu'sul (9639, item level 50) is present in our current
catalog; none of these weapons is equipped in the 147 published profiles.
These are recorded omissions, not corrected item effects. Their inherited
proc rates and full duration/debuff handling must not acquire “verified”
status from the amount checks.

### Seal of Fury evidence

[Seal of Fury rank 7](https://www.wowhead.com/forever/spell=20423/seal-of-fury)
describes additional Holy damage, a shield equal to 50% of that damage while
a shield is equipped, and a four-second taunt on Judgement.
[Improved Seal of Fury's restoration spell](https://www.wowhead.com/forever/spell=1314104/improved-seal-of-fury)
describes 60 mana at level 60, plus 15% per attacker level above the Paladin,
capped at 45%. These tooltips were retrieved on 22 September 2026.

An independent check of the captured **1.60.1.69893** tables found:

- Proc 20418: 35 base Holy damage and 0.1 spell-power coefficient.
- Judgement 20414: base 153, variance 0.087591, 3.69 per scaling level,
  0.45 spell-power coefficient; learned at 58 with maximum scaling level 64.
- Seal 20423: 200 mana; a Paladin class-skill entry with acquisition method 0.
- Talent 1314103: level-based restoration and the 15%/three-level modifiers.

MythicSim cites build **1.60.1.69913**. The local table check above is explicitly
the earlier captured build, not an independent extraction of that later build.
The current tooltip does not print every flat damage component.

Its follow-up `414c84b184` changes the seal proc to melee crit handling and
adds weapon-specialization scaling. Those choices are explained through client
flags and the existing Righteousness implementation. They are not a substitute
for observing which attacks trigger Fury and which modifiers affect it.
The implementation also uses magic hit/crit handling for Judgement, while the
captured `SpellCategories` record has defense type 2. That mapping needs checking
before treating the entire implementation as verified.

The shield's replacement rule is expressly provisional in
[their implementation notes](https://github.com/sage3648/mythicsim-forever-engine/blob/e6d999e9270c652db0820bee08c68f9f22a9079a/docs/mythicsim-seal-of-fury.md).
Its taunt does not change encounter targeting in the simulator.

### Impact on the published matrix

None of the current 23 DPS profiles casts Seal of Fury or Hammer of the
Righteous. Identifying those missing registrations does not itself change
their recorded DPS or establish that either ability improves those builds.
Shield-break mana is especially relevant to tanking and incoming damage.
It must not be granted to a two-handed Ret profile that is not taking hits.

Lightning Shield is also absent from the published Shaman APLs; their initial
shield setup precedes combat. Even an explicit Earth Shock → rank-1 Lightning
Shield loop produces the same three Earth Shocks over the inclusive 12-second
test window before and after the fix: the shared cooldown already prevents
early recasts. The fix removes unsupported timer mutation, not demonstrated
extra damage.
Matched 5,000-iteration replays at seed 20291951 reproduce the Orc Elemental,
Stormcaller and Enhancement DPS exactly (510.58, 529.44 and 692.18), including
their action/resource metrics after ignoring collection order.

### Remaining work

The comparison identifies bounded follow-ups, not complete feature parity.
Fury's absorb behavior, Hammer availability, guardian formulas and the legacy
weapon effects above remain open. No MythicSim implementation has been copied
wholesale into this engine. See T45–T47 and H11/H20 for actionable checks.
