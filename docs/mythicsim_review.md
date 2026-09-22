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
tests, metadata and integration have been reviewed. The inherited upstream
changes are being compared separately; this is not a completed line-by-line
audit of every generated spell table or talent migration.

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

### Review queue

Finish the imported upstream spell-data and modifier migrations against our
current implementations. Separate behavior changes from structural refactors,
and retain source gaps rather than treating generated tables as an availability
list. No MythicSim implementation has been copied into the released engine.
