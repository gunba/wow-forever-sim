# Auto attacks

Ordinary spellcasting retains the swing restrictions from
[wowsims/classic `7779ebbf79`](https://github.com/wowsims/classic/tree/7779ebbf79dc7f1341e6ab939b28a3402c9a730a).
The earlier role-wide permission to keep attacking through spellcasts was too
broad. Weapon-cast exceptions are now explicit rather than applying to every
spell cast by a melee character.
Hunter Auto Shot follows [the original Classic Hunter registration](https://github.com/wowsims/classic/blob/7779ebbf79dc7f1341e6ab939b28a3402c9a730a/sim/hunter/hunter.go#L249-L264):
a 500 ms base wind-up divided by ranged attack speed, the cast-time/no-GCD
flag, and a cast-in-progress guard. Only its Forever projectile speed differs.

| Action | Current behavior |
|---|---|
| Auto Shot | Original WoWSims Classic half-second base wind-up, shortened by ranged attack speed. It occupies the cast bar but is off the normal GCD; it cannot start while another shot is casting. Forever retains its sourced projectile speed. |
| Aimed Shot | Two-second base cast. Auto Shot waits until it finishes; an Auto Shot already winding up delays Aimed Shot. |
| Sniper Shot | Four-second base cast; Auto Shot waits for its completion. |
| Multi-Shot | Client-listed half-second cast; Auto Shot cannot start during it. |
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

Slam retains its explicit weapon-cast exception in Forever. Hunter special
shots no longer have that exception: their Auto Shot uses the original Classic
wind-up and cast gating. A ranged swing due exactly when a special shot ends
waits for the special's completion callback before winding up. Range, movement,
pet dismissal, stealth, form changes and equipment swaps retain their distinct
restrictions.

Regressions cover ordinary-cast versus weapon-cast readiness, caster melee
exclusion, full swing resets at zero/one/three/five Maelstrom stacks,
back-to-back Shaman hardcasts, Slam, and Hunter Auto Shot suppression during
the Aimed Shot/Multi-Shot/Sniper Shot cast bars. Chain Lightning neither
benefits from nor consumes Maelstrom stacks.

Results using the earlier instant Auto Shot model are superseded for all four
Hunter builds. The current benchmark retains their exact equipment, talents,
APLs and external settings while replacing the shot timing. The preceding
continuous-melee correction remains separate.

At the same 5,000-iteration seed, against the [previous published
results](https://github.com/gunba/wow-forever-sim/tree/forever/artifacts/history/e5b762523), the race-weighted
Marksmanship mean falls **917.42 → 806.63 DPS** (−100.69 to −117.92 per
race), while Beast Mastery falls **879.21 → 844.71 DPS** (−28.42 to −47.81).
Survival and Pet/Melee are unchanged because their equipped benchmark range
does not fire Auto Shot. This is a timing-only comparison, not an updated
Hunter gear search; the previous equipment selection is legal but no longer
established as best under the corrected timing.

## Earlier melee correction (historical)

Results using the earlier continuous-melee model were superseded for affected
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
