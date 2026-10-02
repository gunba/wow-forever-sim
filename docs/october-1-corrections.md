# October 1 corrections

## Sources and interpretation

The [official development notes](https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-october-1/2360696/4), last edited October 2, are the primary change record. Static client tables are pinned to **1.60.1.70170**. They do not include every hotfix: Bloodthirst and Champion of the Light still carry older values in some static rows.

Current tree identities, positions, prerequisites and rank curves are corroborated by the [ForeverDB class dataset](https://foreverdb.net/data/classes.json) and [ForeverChanges Warrior calculator](https://foreverchanges.pro/talents/warrior). These are client-derived implementation sources, not independent measurements of server scripts. Detailed ForeverDB spell shards were inaccessible; no conclusions depend on pretending those downloads succeeded.

### Important conflicts

- **Warrior critical-swing Rage:** the later Warrior-specific note increases the earlier 75% bonus to **100%**. Warrior outgoing white criticals therefore generate twice normal Rage. Bear retains the explicitly stated **75%** bonus.
- **Dual Wield Specialization:** the notes remove the old 20%-per-point off-hand Rage bonus. Current client-derived rank curves give **10% per point**, alongside 5% off-hand damage per point. Furious Precision separately gives 4/7/10% off-hand hit. The known-issue report of stale DWS hit is not modeled as intended behavior.
- **Downranking:** the notes describe penalties, not a ban. Active source presets use their highest registered damage ranks. The exact server SP/proc penalties remain unresolved; the engine does not reject all low-rank imports.
- **Primal Bite threat:** the notes say approximately twice the previous threat, but do not supply an absolute coefficient. The model applies a 2× relative increase to the inherited 1.5× value, giving a provisional 3× ability multiplier. Neither the old absolute value nor an exact new 3× coefficient has been verified.

## Class coverage

### Druid

- Bear white criticals gain 75% more outgoing Rage after separating the damage critical multiplier. The existing Bear baseline Rage formula is unchanged and remains qualified; this is not a new normalization formula.
- Swipe gains 3% Attack Power. Faerie Fire no longer resets the Forever melee timer.
- Tiger's Fury and King of the Jungle are removed from the current tree/preset. Shifting Power requires Cat Form, costs 55% base Mana before shapeshift-cost discounts, returns 40 Energy, has a 1-second GCD and a 16-second cooldown. Improved Shifting Power subtracts 4/8 seconds.
- Current [Wolfshead Helm](https://www.wowhead.com/forever/item=8345/wolfshead-helm) adds 20 Energy to Shifting Power and retains its separate 5 Rage on Enrage. Client spell 17768 corroborates both effects.
- [Howling Idol](https://www.wowhead.com/forever/item=272427) and client spell 1291059 now subtract 1 second from Shifting Power. The full Cat Tier 1 override separately subtracts 1 second; neither is the old three-second Tiger's Fury reduction.
- Primal Bite selects ranks at 25/36/48/60 rather than relying on sparse character-level keys. Threat uses the qualified relative update above.
- Form/disarm ordering needs a controllable disarm encounter and observations of pre-shift versus post-shift disarms. The reference encounter has no disarm; it is not evidence that this interaction is implemented.

### Hunter

- Deflection grants 1% Parry per point instead of 2%.
- Sniper Shot's 45-yard range and three-shot range extension are relevant to range/movement encounters. The fixed in-range reference does not gain damage from them; range-state coverage must remain separate from the existing shot damage/cast model.
- Tame-level pet-rank fixes change acquisition legality, not proven new level-60 damage coefficients. Selected pet abilities still need to respect learned level/rank availability. Aggressive mode and icon changes are not numerical pet buffs.

### Mage

- Heating Up is the current display name for the consumed Pyroblast proc; the stable internal talent field is unchanged.
- Combustion expires after three criticals rather than four.
- Improved Scorch applies only after its parent hits, without another resistance check. Winter's Chill already uses the parent result without an extra resistance roll; its legitimate talent-rank proc chance is not removed.
- Moving-cast restrictions on comprehension scrolls require movement support. Current stationary profiles do not cast those scrolls.

### Paladin

- Redoubt grants 4% Block per point instead of 6%; Holy Shield grants 30% instead of 20%.
- Champion of the Light converts 20/40/60% of Intellect to **spell damage**, not healing-compatible Spell Power. The damage-only stat dependency preserves sourced healing/damage separation.

### Priest

- Devouring Plague already critically strikes through the shared Forever periodic outcome. No redundant or snapshot-only critical rule is added.
- Early Demise already restricts the Death damage bonus to execute conditions.
- Inner Focus no longer grants periodic critical chance to pure periodic spells. Shared direct/periodic families retain their direct eligibility; free-cast consumption is a separate rule.
- Shadow Weaving uses landed parent results without another resist check. The notes fixing failed applications do not establish that lower talent ranks should lose their intended application chance.

### Rogue

- Expose Armor already spends combo points only on a landed application.
- Setup remains an explicitly unimplemented reactive talent. A correct implementation must require the attacker to be the current target and distinguish fully resisted spells from misses/partial resists. The engine's partial-resist flag alone is not sufficient. Current DPS profiles have no incoming attacks that could establish a gain.

### Shaman

- Disease Cleansing duration and Totemic Recall returns are utility/resource coverage, not unannounced damage buffs. Registration, duration and active-slot replacement need separate accounting.
- Ghost Wolf visibility does not affect the stationary damage model.
- External Flametongue modeling is described below; an external party toggle is not a claim that every caster-owned totem action is registered.

### Warlock

- Drain Soul uses the shared interruptible channel path. Starting another cast must cancel remaining ticks; this requires a focused interruption replay, not merely a tooltip comparison.
- Hellfire's periodic critical permission is source-supported. Hellfire remains a known missing family requiring its enemy/self-effect chain and an explicitly in-range AoE scenario; it is not silently added to the 20-yard reference rotation.
- Soul Harvest's current display metadata gives 50/100% increased Mana regeneration for 10 seconds after a qualifying Drain Soul kill, and 50/100% regeneration while casting. Its kill-triggered implementation remains absent and marked unsupported. A fixed-duration single boss does not provide repeated qualifying kills; the talent is not treated as a permanent regeneration buff.
- Incubus model resolution and aggressive-mode UI do not change the selected pet's supported ability values.

### Warrior

- Queued Heroic Strike/Cleave no longer remove the off-hand dual-wield miss penalty. The queued main-hand replacement still uses its own appropriate outcome table.
- Booming Voice reduces Shout Rage costs by 5% per point. Raging Blows subtracts 3 Rage from Cleave/Whirlwind; off-hand Whirlwind is baseline when legally dual wielding.
- Unbridled Wrath gives one Rage on two-handed procs as well as one-handed procs. Its existing supported proc model is not arbitrarily increased from the ambiguous bug-fix note.
- Bloodthirst uses 45% Attack Power. It no longer activates Blood Craze; incoming critical/large-hit triggers remain.
- Spearing Strike requires Battle Stance, without a two-handed requirement. Berserker Rage becomes available at level 30 in Forever.
- Removed Precision, Boundless Rage, Improved Cleave and Toughness do not contribute hidden effects. Current tree decoding is name-based and independent of Classic proto declaration order. It resolves after environment attachment but before item/set effects inspect talents.
- Source presets preserve surviving allocations and legal identity while reallocating removed points. They are legal migrations, not a claim of newly optimized talents.
- Lingering Rage concerns out-of-combat decay, which this combat-only model does not simulate. Gore Drinker is exposed as **not simulated** pending its triggered aura, refresh/charge consumption and multi-target event semantics; no healing is fabricated from guessed proc ordering.
- Enemy tapping and the game's cooldown manager UI do not change the reference damage/threat equations.

## Shared and non-combat changes

Eureka no longer amplifies ordinary periodic damage through a charged cast or an immediate application tick. Direct channel ticks retain their direct eligibility, including channels represented by a single spell. Its resource discount and one-charge-per-cast accounting remain separate. Periodic healing needs healer-specific coverage; no claim of complete healer support follows from damage checks.

Beta level 30 changes which experiments are now possible. It does not validate level-60 coefficients automatically. Gamepad controls, pet stable behavior, profession-service pricing and spatial visibility are outside current damage/threat/survival accounting.

## Flametongue comparison

The opt-in external Flametongue model deliberately contributes **zero Spell Power scaling**. Provisional per-weapon-second values are 5.48/7.81/10.61/13.63, obtained by dividing the client dummy-effect values by 100. A landed eligible main-hand hit deals that value times the active weapon's base speed. Haste changes attack frequency, not the per-hit weapon speed used here.

This division, magic hit/critical behavior, eligible specials and shapeshift interpretation are **not measured server rules**. Windfury Totem excludes Flametongue and Grace; Grace plus Flametongue is a legal comparison setting. Flametongue Weapon excludes the totem proc. Conflicting imported Windfury settings take deterministic precedence.

Two non-overlapping iteration-seed ranges starting at 2026100141 and 2026110141, 5,000 iterations per arm, on fixed current Human profiles give:

| Build | Windfury DPS | Grace + provisional Flametongue DPS | Difference |
|---|---:|---:|---:|
| Retribution | 1037.39 | 974.89 | -62.49 (-6.02%) |
| Fury | 893.77 | 871.63 | -22.14 (-2.48%) |

No gear, talents, rotation, consumes, Tier 1 or encounter settings changed between arms. These are scoped support comparisons, not evidence that the provisional damage formula is correct or that every melee spec favors Windfury. Canonical support is not silently replaced.

## Release state

The corrected run covers **201 exact profiles and 804 baseline/Tier-off/+10%/+50% scenarios**, each with 5,000 iterations and no APL warnings. Equipment, enchants, options, support and encounters are unchanged. Only documented talent/rotation migrations differ: 58 talent allocations and 37 rotations changed, with four profiles changing both.

A separate **68-run tank check** covers all 17 tank race profiles in both original workloads and two non-overlapping 5,000-iteration seed ranges. It preserves workload-specific healing rather than substituting the single-target healing rate into multi-target runs. These are correctness replays, not a renewed tank selection campaign.

Max-rank migrations materially affect Mana: BM records about 38–41 mana-limited seconds, Smite 44–51 and Stormcaller 65–70. These resource limitations remain visible; no new rotation or gear optimization hides them. Feral records about 1.5–2 seconds under Shifting Power.

[Before/after profiles and results](../artifacts/correctness/october_1/comparison.json), [superseded complete evidence](../artifacts/correctness/october_1/before.json.gz), [tank checks](../artifacts/correctness/october_1/tank_summary.json) and [provisional support comparison](../artifacts/correctness/october_1/support_summary.json) preserve the distinction between phases. Historical gear/talent selection controls do not validate the new mechanics. Targeted native mechanics and source-ledger checks pass. Browser profile loading reproduces native DPS, the +50% replay retains its inputs, all 31 build icon groups decode and the categorized question controls pass. Final publication and actual live verification remain pending.
