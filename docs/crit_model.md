# Critical strike model

## Druid crit aura

In Forever, Moonkin Aura and Leader of the Pack supply the same non-stacking
3 percentage points of melee, ranged and spell critical chance. Either provider
is sufficient. The correction does not change party coverage or make a party
aura raid-wide. Classic retains its separate physical/spell effects.

The web control is **Crit Aura (3%)**, representing one active aura, not two
independently selectable buffs. Published setups select the Moonkin provider.
If multiple providers are present in a group, they still produce only one
effective aura.

Previously, Moonkin supplied only spell crit and Leader only physical crit.
The preceding 147 benchmark requests selected both source flags, so that
combination already received the correct 3 points in each pool, not 6. The
corrected setups select just one provider. This is a representation correction,
not an additional crit bonus.

## Generic equipment crit

The current import already treats generic item crit as shared:

- The pinned level-60 client rating table uses **14 rating per percentage
  point** for melee, ranged and spell crit.
- `tools/database/compile_forever_equipment.py` stores generic `CritRating`
  once. `sim/core/ruleset.go` shares that equipment contribution between the
  physical and spell pools.
- Outlaw's Collar's 21 rating therefore contributes 1.5 percentage points to
  each, not separate additive copies of the same stat.
- The source catalog's 123 flat-crit records use generic `CritRating`.
  The current loadouts use 34 distinct flat-crit items, each represented once;
  none of their equipped enchants contributes a flat crit stat.

Sources: `assets/db_inputs/forever_combat_ratings.json`,
`assets/db_inputs/forever_gear_catalog.json`, and the exact requests in
`artifacts/forever_dps_5min.json`.

## Item-picker EP limitation

The browser's item EP calculation in `ui/core/player.ts` takes a dot product
of the stored item fields and EP weights without the engine's shared-crit
projection. Mage's default weights value SpellCrit but not MeleeCrit.
Consequently a generic crit bonus stored in the latter field can contribute
zero to that EP estimate even though it correctly increases spell crit in
the simulation. This is an item-picker valuation issue, not missing crit in
simulated DPS. It was identified but not changed in this pass.

## Scope and remaining risk

This is not a merge of every crit modifier. Attribute conversions, ability
bonuses, school-specific bonuses, ranged-only effects, crit suppression and
effects that cannot crit remain distinct. The explicit benchmark stat
adjustments also bypass equipment sharing to avoid counting them twice.

The equipment helper sums the two stored crit fields before sharing them.
That representation would be unsafe for an explicitly scoped flat-crit item
or enchant, or a legacy record storing the same generic bonus in both fields.
No such case was found among the current loadouts' flat-crit records. Other
legacy catalog entries and individual proc effects are not certified by that
check; their exact source effects need review before changing their scope.
No blanket gear, enchant or talent conversion was made in this pass.
