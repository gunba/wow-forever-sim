# Auto attacks

Forever's auto-attack roles keep their weapon events running during ordinary
ability casts. The engine no longer imports Classic shot pauses or an extra
half-second wind-up.

| Action | Current behavior |
|---|---|
| Auto Shot | Independent instant weapon event on its normal hasted swing schedule; does not occupy the cast bar or GCD. |
| Aimed Shot | Two-second base cast; Auto Shot continues. |
| Sniper Shot | Four-second base cast; Auto Shot continues. |
| Multi-Shot | Client-listed half-second cast; Auto Shot continues. |
| Volley | Channel no longer explicitly delays ranged swings. |
| Slam | Does not reset or suspend weapon swings, including without Improved Slam. |
| Enhancement spell casts | Do not reset melee swings. |
| Other ordinary casts in weapon roles | Do not block white attacks or queued swing replacements. |
| Caster roles | No melee weaving during or between casts. |

Continuous swings are a reported gameplay rule. Volley and the other ordinary
weapon-role casts follow the same general model; they have not each been
separately measured. Range, movement/range transitions, pet dismissal, stealth,
form changes and equipment swaps retain their distinct restrictions.

The auto-attack role classification includes Hunter, Rogue, Warrior, Feral,
Enhancement and Retribution, plus the corresponding modeled tank roles. It does
not make ordinary spells castable during another spell.

`sim/core/continuous_autos_test.go` checks readiness without allowing simultaneous
special casts or caster melee. The benchmark's Hunter regression checks equal
Auto Shot counts with and without repeated Aimed Shot/Multi-Shot, and checks
the cast-time configuration.

Earlier Hunter comparisons used the interrupted-shot model. They must not be
presented as current results. The updated benchmark JSON records its auto-attack,
shot-cast and Energy models explicitly.
