# Speed effects and unresolved GCD rules

Client source: cached `SpellEffect`, `SpellMisc`, `SpellCategories` and
`SpellCooldowns` CSVs for Forever **1.60.1.69893**, with selected race and
Tier 1 records in `assets/db_inputs/forever_races.json` and
`assets/db_inputs/forever_tier1_bonuses.json`. These are client descriptions
and flags, not a capture of server timing.

| Source spell | Client aura types | Modeled distinction |
|---|---|---|
| Berserking 20554 | 319 melee speed, 140 ranged speed, 65 casting speed; 10% each | Speeds attacks and hardcasts. A reported Arcane Missiles observation rules out treating its casting-speed aura as generic channel haste. No Energy regeneration bonus. |
| Wind Blessed 1259710 | 342 combined melee/ranged haste, 65 casting speed; 1% each | Speeds attacks and hardcasts. Its assumed Energy effect remains T15. The presence of aura 65 alone does not establish channel timing. |
| Hunter Tier 1 1300945 and the other DPS haste sets | 342, 65; 1% each | Same provisional distinction as Wind Blessed. These forced set bonuses explain some non-Skyborne profiles' 1% base speed. |
| Minor Haste 13928 (gloves) | 342, 65; 1% each | The catalog's melee-haste stat applies once to both autoattack speeds; its separate casting-speed aura affects hardcasts, not the SpellHaste stat. |

`SpellCooldowns` records **1,500 ms** and `SpellCategories` records GCD
category **133** for Earth Shock 10414, Flame Shock 29228, Frost Shock 10473,
Mortal Strike 12294 and Arcane Missiles 25345. The first three shocks share
the same magical defense type. The old Flame Shock `IgnoreHaste` flag was
inherited from Classic and made its GCD differ from the other shocks only
after a Forever-specific spell-GCD rule was added; the difference is removed.

The exception audit found other nonphysical-school spells with a normal
1.5-second GCD and inherited `IgnoreHaste`: Hunter Arcane Shot, Serpent
Sting, traps, aspects and Summon Hawk; Paladin Holy Strike; and Succubus
Lash of Pain. Some of these use a melee/ranged *hit table* despite a Holy,
Arcane or Nature *school*. Pet-GCD, one-second-GCD, off-GCD and pure
physical-school entries are a different case. Forever GCD haste now follows
the **spell school**, rather than the hit table or an inherited
`IgnoreHaste` setting that also suppresses cast-time haste. An in-sim
regression covers Holy Strike, Arcane Shot, the shocks and physical
Stormstrike.

Neither `StartRecoveryTime`, `StartRecoveryCategory` nor the spell's defense
type **specifies whether the server hastes its GCD**. `SpellMisc` marks many
physical abilities with the `Is Ability` bit, but does not provide a complete
per-effect GCD-haste policy. WoWSims Forever at `ea5412873474` uses
`SpellSchoolPhysical` to set `IgnoreHaste` when parsing client spells
(`sim/core/spelldata/resolve_spell.go`, lines 173–184), **not**
`DefenseTypeMelee`. Its Holy Strike retains a Holy school/melee hit table
and no haste exemption. This independent implementation supports using
school consistently, but is not proof of the server rule. In particular,
pure-physical 1.5-second abilities remain fixed in both models pending the
matched level-20 timing comparison in
[`in_game_checks.md`](in_game_checks.md) (T40).

The model now separates hardcast speed from channel speed. The client
Arcane Missiles rank 8 retains its five one-second triggered missiles
and does not carry a client flag proving Berserking should compress the channel.
The available observation says it does not. A separately represented
`SpellHaste` stat continues to scale channels **provisionally**; the current
gear catalog has no item carrying that stat. Missile Barrage's explicit
half-second missile period is independent of either source.

`assets/db_inputs/forever_combat_ratings.json` records **10 raw haste
rating per percentage point** at level 60 for melee, ranged and spells.
The inherited engine constant still treats a haste *stat* point as one
percentage point. Current selected equipment contains no haste-rating item;
its Minor Haste enchant is instead a sourced flat 1% aura. Do not use the
inherited rating constant to predict future rating-bearing gear. This
conversion needs a separate flat-percent versus raw-rating migration before
new haste-rating items enter the equipment pool.
