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
PYTHONDONTWRITEBYTECODE=1 python3 tools/database/compile_forever_sets.py \
  --cache /tmp/forever-review/client-data
PYTHONDONTWRITEBYTECODE=1 python3 -m unittest discover \
  -s tools/database -p 'test_*forever*.py'
```

The inputs are the Forever gear-planner snapshot and the pinned Wago client
tables. The catalog records their SHA-256 hashes and table URLs. The planner
snapshot was checked against the live endpoint during this pass and matched:
`4c0981cdeab4036898c7a8010a277147b625dafc75963db4679d6a56c1abbf13`.

Acquisition comes from crafting skill/spell records, dungeon drops and quests,
exported Honor purchases, reviewed non-raid quartermasters, or the two
Darkmoon 1,200-ticket necklace exchanges. Known battleground quartermasters retain
their faction restrictions. The catalog pins vendor-export
hashes alongside the other inputs. Exported stats and weapon damage take
precedence over the client/planner representation.
Explicit old-raid zones and raid quest categories are excluded. A vendor
location alone does not establish a qualifying source; this excludes the
Naxxramas Waywatcher Bindings despite their Eastern Plaguelands vendor listing.
Item level and numerical item-ID ranges are not availability tests.

`forever_ilvl65_items.json` captures all 706 records in the Wowhead comparison
filter, not just its first 50 rendered rows. Every captured record has item level
65. The list includes unavailable raid rewards and omits cloaks, so it supplements
rather than replaces acquisition checks and the wider verified equipment catalog.

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

### Enchants

The enchant audit uses Forever tooltips and client SpellItemEnchantment /
SpellEffect records. Chromatic Mantle of the Dawn (2488) supplies a non-raid
shoulder option with five resistances; it adds no offensive stats.

Three speed enchants need different handling:

- Iron Counterweight (34 → 7217): 3% melee speed only.
- Arcanum of Rapidity (2543 → 22841): 1% melee and ranged speed per piece.
- Minor Haste (931 → 13928): attack and casting speed. Its general haste
  participates in the provisional haste-to-Energy model.

The first two are not Energy regeneration bonuses. Their effects are kept
outside the generic haste stat. Identical Arcanum copies retain additive
stacking; this interaction has not been measured in-game.
Ranged-only scope hit counts for ranged attacks, not an unused melee weapon;
Summon Hawk's unresolved melee hit model remains distinct.

Two new crafted enhancements are imported:

- [Wild Leather Armor Kit](https://www.wowhead.com/forever/item=279258):
  enchant 8719, +4 Defense and +10 Stamina on chest, legs, hands or feet,
  minimum item level 45.
- [SAF-T Ultra Precision Scope](https://www.wowhead.com/forever/item=279272):
  enchant 8720, +2 percentage points of ranged critical chance.
  SpellEquippedItems restricts its effect 1310308 to bows, guns and crossbows
  (mask 262156); it does not grant general melee or caster critical chance.

The equipment search fills every enchantable slot and compares legal
alternatives using native DPS runs. Raid-reward inscriptions are excluded.
Healing-only enchants currently inherit the engine's one-third damage fallback:
+24 healing bracers contribute eight modeled spell damage. This is retained
behavior, not a newly verified Forever rule; in-game check T22 covers it.

### Items

The current snapshot admits 1,958 items across the supported equipment slots,
including exported PvP items, reviewed quartermaster stock, dungeon quests
and zone-level drop records. Of these, 1,501 have client stat records;
457 use published Forever planner stats.
Eighty-one candidate records remain unresolved, mainly class-set templates
without published stats and weapons missing damage values.

The compiler currently admits 1,908 of these records. Forty-five race-restricted
quest items await an explicit race mapping; five have unsupported stat types.
Dungeon quests outside a named dungeon category need individual acquisition
review: an outdoor turn-in location alone cannot distinguish a dungeon chain
from a raid chain. Mutually exclusive quest rewards must also be checked before
selecting loadouts.

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

The original search seeds average roughly item level 55–59 and have no
enchants. They are seed loadouts, not optimized or equivalent equipment sets.
They use reviewed crafted/dungeon items and validate slots, proficiencies and equipment
limits. `-write-seed-gear` rebuilds these starting profiles with a deterministic
stat score; it is not an optimized best-in-slot search. The
[coordinate search](../tools/forever_bench/README.md#equipment-search) replaces
that heuristic with candidate DPS comparisons and independent validation.
Lower item levels can still win with verified Forever stats. For example,
Gloves of the Greatfather supply 24 spell damage plus 10 Nature damage in the
client, despite their item level of 38. The search does not replace a stronger
item just to raise the displayed average. Lower-level trinkets are compared
explicitly because the known level-65 Undermine trinkets share a one-item limit.

Ordinary set bonuses use current client effects, with actual piece counts and
profession requirements. Unsupported thresholds are reported in benchmark
results and do not fall back to Classic effects. They are separate from the
forced Tier 1 role bonuses.

[Weakness Analyzer](https://www.wowhead.com/forever/item=272438) and
[Serenity Field](https://www.wowhead.com/forever/item=272439) have implemented
on-use effects and their shared Undermine Trinkets limit. Adaptive Combat
Assistant's absorption/explosion and Defender's Grip Stabilizer's location
modifier remain unimplemented; those items are not admitted to the DPS search.

Four vendor relic effects are implemented from their current tooltips and
client effects: Howling Idol (272427), Swarming Idol (272430), Burning Totem
(272433), and Libram of Law (272435). They respectively reduce Tiger's Fury's
cooldown, extend Insect Swarm, extend Flame Shock, and increase Judgement damage.
Tests include their interactions with talents and the forced Tier 1 bonuses.
These effects are tested with fixed equipment, not mid-combat relic swaps.
Totem of the Storm (272432) remains excluded: its Lightning Bolt proc rule
depends on the unresolved Maelstrom Weapon proc model.

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
| Expertise: dodge/parry reduction | 10 |

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
