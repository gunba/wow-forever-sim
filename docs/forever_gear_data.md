# Forever gear data

## Simulator translation

`tools/database/compile_forever_equipment.py` translates the source catalog into
`assets/db_inputs/forever_equipment.json`. Generic hit/crit enters one equipment
pool, AP also supplies ranged AP, school-specific power stays school-specific,
and explicit spell damage remains separate from healing. Translation does not
certify proc or set-bonus implementations.

The published armor field is **base armor**, not base plus bonus armor.
[Cloak of Warding](https://www.wowhead.com/forever/item=18413/cloak-of-warding)
shows 44 armor and 170 bonus armor;
[Warden's Leather Waistguard](https://www.wowhead.com/forever/item=252475/wardens-leather-waistguard)
shows 89 armor and 102 bonus armor. Those components must remain separate.
Vendor-export stat handling is independent of this planner representation.

The catalog also retains `ItemLimitCategory` quantities and flags.
The four Spiritcaller armor pieces share **Artisan's Tier**, category 708:
only **one** can be equipped. Greenhammer has its own one-equipped limit,
category 712. Per-item uniqueness alone does not enforce a shared category.

`assets/db_inputs/forever_gear_catalog.json` is the source catalog, not a selected
loadout or a declaration that every listed effect is implemented.

## Rebuild

```sh
PYTHONDONTWRITEBYTECODE=1 python3 tools/database/forever_gear_catalog.py \
  --cache /tmp/forever-review/client-data
PYTHONDONTWRITEBYTECODE=1 python3 tools/database/compile_forever_equipment.py
PYTHONDONTWRITEBYTECODE=1 python3 -m unittest discover \
  -s tools/database -p 'test_*forever*.py'
```

The inputs are the Forever gear-planner snapshot and the pinned Wago client
tables. The catalog records their SHA-256 hashes and table URLs. The planner
snapshot was checked against the live endpoint during this pass and matched:
`4c0981cdeab4036898c7a8010a277147b625dafc75963db4679d6a56c1abbf13`.

Acquisition comes from crafting skill/spell records or named dungeon sources.
Explicit old-raid zones and raid quest categories are excluded. A vendor
location alone does not establish a qualifying source; this excludes the
Naxxramas Waywatcher Bindings despite their Eastern Plaguelands vendor listing.
Item level and numerical item-ID ranges are not availability tests.

## Stats that Wowhead omits

Client `1.60.1.69893` uses item mods 84–89 for Holy, Fire, Nature, Frost,
Shadow and Arcane damage. The matching strings are in GlobalStrings
58795–58806. Known school-specific items corroborate the ordering.

For an item with a client stat record:

```
displayed stat = round(RandPropPoints[quality, itemLevel, slot]
                      * StatPercentEditor / 10000)
```

The importer checks every overlapping stat against the planner and quarantines
conflicts. Shields use the neck/wrist/offhand allocation group, not the
one-handed weapon group.

Examples absent from both the current planner stats and tooltip rendering:

| Item | Client school power |
|---|---:|
| [Robes of Fiery Devastation](https://www.wowhead.com/forever/item=279266) | 26 Fire |
| [Mana-infused Cord](https://www.wowhead.com/forever/item=279267) | 20 Arcane |
| [Rime-encrusted Boots](https://www.wowhead.com/forever/item=279268) | 20 Frost |
| [Fel Cape](https://www.wowhead.com/forever/item=279269) | 15 Shadow |
| [Cloak of Earth and Sky](https://www.wowhead.com/forever/item=279270) | 15 each Nature, Arcane and Fire |

These values are reconstructed from client records, not guessed from the
item's name or a proposed spell-power budget. For example, the robe has an
allocation of 5545 and an epic level-65 chest scalar of 47: round(26.0615) = 26.

## Coverage and remaining work

The current snapshot admits 1,488 items across the supported equipment slots.
1,203 have client stat records; 285 use published Forever planner stats.
Seventy-eight candidate records remain unresolved, mainly class-set templates
without published stats and weapons missing damage values.

Absence from static ItemSparse is **not** evidence that an item is unavailable.
Shadowcraft Cap and many other dungeon items lack a static record but have
published Forever stats. These retain a distinct `statSource`. The current
published pool is especially sparse for rings, necks and dungeon trinkets;
missing items are not filled with Classic stats.

Item effects, set spells, skill requirements, unique limits and source records
are retained for subsequent eligibility/effect checks. An unrecognized stat
modifier remains visible as `UnmappedItemMod...`, rather than disappearing.
The compiled catalog now supplies the benchmark's `forever_<build>.gear.json`
profiles. The database generator replaces the old loot pool with these records
and the imported vendor records retained for comparison; benchmark eligibility
is narrower than that database pool.

The starting profiles use reviewed crafted/dungeon items, preserve their
existing per-slot enchants, and validate all slots, proficiencies and equipment
limits. `-write-seed-gear` rebuilds these starting profiles with a deterministic
stat score; it is not an optimized best-in-slot search. Unreviewed item effects
and active equipment-set bonuses are not admitted by that selector.

[Tier 1 bonuses](forever_tier1.md) are a separate simulator setting. They do not
require raid gear or alter the catalog's item stats.

## Hit-budget accounting

The catalog stores raw ratings. `forever_combat_ratings.json` now records the
level-60 coefficients from Blizzard's `GameTables/CombatRatings.txt`, build
**1.60.1.69913**, FileDataID **1391669**. The original table is retained as
`forever_combat_ratings_1.60.1.69913.txt`, with its SHA-256 in the JSON.

| Rating | Raw rating per percentage point |
|---|---:|
| Melee, ranged and spell hit | 10 |
| Melee, ranged and spell crit | 14 |
| Melee, ranged and spell haste | 10 |
| Dodge | 12 |
| Parry | 15 |
| Block | 5 |

The same coefficients appear at every level from 1 through 60. In particular,
spell hit is **not** TBC's 8 rating per point. Vendor crit/dodge percentages
independently corroborate those two conversions. Generic item hit/crit must
enter only one simulator stat pool: the equipment pass shares it with spells.
These coefficients do not establish combat-table miss caps.

The source was extracted with
[TACTTool 0.2.0-alpha1](https://github.com/wowdev/TACTSharp/releases/tag/0.2.0-alpha1):

```sh
TACTTool -p wow_classic_beta \
  -b 6c0df97e8e481a9a41600e373367c200 \
  -c 5525ea1ce6668e895569c89c2d6a154c \
  -m fdid -i 1391669 -o CombatRatings.txt
python3 tools/database/import_forever_ratings.py CombatRatings.txt \
  --build 1.60.1.69913 \
  --build-config 6c0df97e8e481a9a41600e373367c200 \
  --cdn-config 5525ea1ce6668e895569c89c2d6a154c
```

Rating-to-percentage conversion and the benchmark's offensive-stat budget
ledger are separate operations.
`StatPercentEditor` and RandPropPoints reconstruct quantities; they do not
establish equal-cost exchanges between attack power, spell power and hit.
Consequently the catalog importer does not implement or assume a 1:1 exchange.
The benchmark now uses a separate
[linear budget model](../tools/forever_bench/README.md#hit-budget-model), retaining
item records unchanged and reporting every debit or surplus-hit refund. Those
exchange prices are benchmark assumptions, not coefficients from the client table.
