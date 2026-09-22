# Auto attacks

Ordinary spellcasting retains the swing restrictions from
[wowsims/classic `7779ebbf79`](https://github.com/wowsims/classic/tree/7779ebbf79dc7f1341e6ab939b28a3402c9a730a).
The earlier role-wide permission to keep attacking through spellcasts was too
broad. Weapon-cast exceptions are now explicit rather than applying to every
spell cast by a melee character.

| Action | Current behavior |
|---|---|
| Auto Shot | Independent instant weapon event on its normal hasted swing schedule; does not occupy the cast bar or GCD. |
| Aimed Shot | Two-second base cast; Auto Shot continues. |
| Sniper Shot | Four-second base cast; Auto Shot continues. |
| Multi-Shot | Client-listed half-second cast; Auto Shot continues. |
| Volley | Restores the original channel's ranged-swing delay. |
| Slam | Does not reset or suspend weapon swings, including without Improved Slam. |
| Enhancement Lightning Bolt, Chain Lightning and Lava Burst | Reset the next melee swing to cast completion plus a full current weapon swing duration. |
| Ordinary hardcasts and channels | Block melee autos and queued swing replacements; retain each spell's original reset/delay behavior. |
| Caster roles | No melee weaving during or between casts. |

`StopMeleeUntil` does not simply release a queued swing when the cast finishes.
It starts a fresh full swing timer from that point, preventing free melee
swings between consecutive Shaman hardcasts. The original Shaman helper calls
it even when the calculated cast time is zero: a five-stack instant Lightning
Bolt still resets that timer in this model. Instant abilities without a reset
hook are not given a new blanket reset.

Slam, Aimed Shot, Sniper Shot and Multi-Shot carry an explicit weapon-cast
exception in Forever. They permit ordinary autos, not simultaneous special
casts. The exception is not applied to Classic or to arbitrary spells merely
because their caster is a melee specialization. Range, movement, pet dismissal,
stealth, form changes and equipment swaps retain their distinct restrictions.

Regressions cover ordinary-cast versus weapon-cast readiness, caster melee
exclusion, full swing resets at zero/one/three/five Maelstrom stacks,
back-to-back Shaman hardcasts, Slam, and equal Hunter Auto Shot counts during
Aimed Shot/Multi-Shot/Sniper Shot. Chain Lightning neither benefits from nor
consumes Maelstrom stacks.

Results using the earlier continuous-melee model are superseded for affected
profiles. Both sides of a comparison must use the corrected engine, including
an all-instant candidate compared with a hardcasting baseline. The benchmark
JSON records its auto-attack, shot-cast and Energy models explicitly.

The preceding complete benchmark and sensitivities remain available at
[revision `8d4127a06`](https://github.com/gunba/wow-forever-sim/tree/8d4127a06/artifacts).
The corrected replay covers all 147 profiles and all four scenarios, 588 runs
at 5,000 iterations each. Every request and unnormalized profile is unchanged.
Only the five Enhancement races change DPS; Orc falls from 692.18 to 601.88
in the main scenario. No research rotation has been substituted into these
rankings.
