# Gear comparisons

The finalized 26 builds were screened across all 171 supported race/build
combinations. The search ran **89,979 trials**, including comparison baselines
and validation runs. Every profile converged; 24 changed loadouts survived
independent confirmation and 147 retained their equipment and enchants.

Talents, rotations, pet choices, consumables and external buffs stayed fixed.
The encounter remains a 300-second level-63 Dragonkin with 3,731 starting
armor, ordinary MP5 and complete role-specific Tier 1.

## Retained changes

The damage ranges below are pooled gains against each race's own previous
loadout, not gains over another race.

| Build | Changed races | Confirmed gain |
|---|---|---:|
| Retribution | All three | +14.82–32.24 DPS |
| 2H Bloodthirst | Orc, Dwarf, Tauren, Windshaper, High Order | +2.87–7.43 DPS |
| Destruction | Orc | +5.48 DPS |
| Fire | Orc, Troll, High Order | +3.49–5.01 DPS |
| Pet/Melee | Orc, Human, Night Elf, Windshaper, High Order | +2.67–3.60 DPS |
| Survival | Orc, Tauren, Human, Dwarf, Night Elf | +3.20–3.47 DPS |
| Marksmanship | Tauren | +2.44 DPS |
| Enhancement | Orc | +1.93 DPS |

Notable changes:

- Ret replaces Blackfury with race-specific PvP weapons and uses the glove
  Healing Power enchant. That enchant has a **verified 35 healing / 12 spell
  damage** split; only its explicit 12 damage contributes to DPS. It is not a
  healing-to-damage conversion.
- Orc Enhancement and 2H Bloodthirst select the Battle Axe; Dwarf 2H Bloodthirst
  selects the Pulverizer. Several other changes are enchants rather than items.
- Several melee Hunters prefer Crusader over Lifestealing. Other races keep
  their previous enchant; gains were not assumed to transfer between races.
- Tauren MM selects Bloodvine Boots. Armor type alone is not an offensive
  requirement: these are legal cloth boots, and the full loadout is compared
  after paying for hit.
- Fire's changed races use Orb of the Darkmoon. Orc Destruction changes its
  second Heart of the Mountain trinket.

The smaller gains remain modest. The final matrix uses another seed, so its
displayed means can differ slightly from these paired comparisons.

## Method and limits

The comparison pool starts with the complete 706-item exact-ilvl-65 export.
Class/equipment restrictions and source eligibility remove unusable entries.
Verified catalog cloaks and trinket alternatives cover the original filter's
omissions and equipment limits; this is not a claim that every slot is ilvl 65.
Old-raid/token rewards remain excluded.

Each slot is screened against the current loadout, with legal weapon layouts
compared together, followed by enchant and second-profession comparisons.
Every candidate recalculates paid shared hit. Catalog records are not rewritten
by that adjustment. The strongest three improving screen candidates receive
independent slot validation; further passes stop when no accepted gain remains.

Changed final loadouts have two independent 5,000-iteration pairs, seeds
20295941 and 20296071. The first is separate from the search's screening and
slot-validation seeds. Acceptance requires a pooled gain larger than
`1.96 × (baseline SE + candidate SE)`, without assuming paired-arm covariance.
All 24 passed. The smallest margin is Orc Enhancement: +1.93 DPS against a
conservative ±1.59 DPS noise bound.

This is a finite coordinate search under a beta model, not proof of a global
optimum. Proc, pet, rage, low-rank scaling and spell-coverage uncertainties
remain in the [in-game checks](in_game_checks.md) and [spell audit](spell_coverage.md).

## Evidence

- [Summary and per-race decisions](../artifacts/gear_search/current/summary.json)
- [All search trials](https://gunba.github.io/wow-forever-sim/classic/review/gear_search/search.json.gz)
- [Complete independent pairs](https://gunba.github.io/wow-forever-sim/classic/review/gear_search/validation.json.gz)
- [Previous benchmark](https://gunba.github.io/wow-forever-sim/classic/review/gear_search/before.json.gz)
- [Replay instructions](../tools/forever_bench/README.md)

Earlier research-build comparisons used their original fixed gear and remain
distinct from this gear pass.
