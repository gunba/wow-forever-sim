# Forever Tier 1

The **Tier 1 bonuses** setting applies the selected DPS role's 2-, 3-, 4- and
5-piece bonuses independently of equipment. It adds no raid-item stats. The
setting is saved in player exports and saved settings, and is ignored under the
Classic ruleset.

The benchmark enables it by default; `-tier1=false` disables it. Raw requests
record `foreverTier1Bonuses`, and the results also record the active set bonuses.
Equipping pieces of the same set does not apply its effects a second time.
Other equipped sets still use their actual piece counts.

## Set assignments

| Builds | Set |
|---|---|
| Balance | Grovekeeper Eclipse |
| Feral | Grovekeeper Ferocity |
| Beast Mastery, Marksmanship, Survival | Wildstalker Armor |
| Arcane, Fire, Frost | Manaflare Regalia |
| Retribution | Justice Battlegear |
| Shadow | Raiments of Conviction |
| Smite | Vestments of Conviction |
| Combat, Mutilate, Subtlety | Grimstitch Armor |
| Elemental, Stormcaller | The Spiritcaller's Storm |
| Enhancement | The Spiritcaller's Rage |
| All Warlock builds | Demonheart Raiment |
| Arms, Fury | Battlegear of Glory |

Creature-specific four-piece bonuses remain conditional. The neutral benchmark
target does not receive bonuses against demons, undead, beasts, humanoids or
elementals. Mixed-target encounters evaluate the restriction separately for each
target.

The one-percent hit bonuses reduce the amount of hit purchased by the benchmark
normalization. They are not part of the offensive item-stat pool charged for
that purchase.

## Data and implementation limits

`assets/db_inputs/forever_tier1_bonuses.json` preserves all eighteen datamined
role-set definitions, spell descriptions and raw effect records from client
**1.60.1.69893**. The twelve assignments above cover the current DPS profiles.

Rebuild the source capture with:

```sh
python3 tools/database/import_forever_tier1.py --cache /path/to/client-data
```

The damage, resource, stat and cooldown effects used by these profiles are
implemented. Utility bonuses for unmodeled crowd-control spells remain
descriptive: they do not invent extra damage. Totem range is not simulated.
Envenom is not currently implemented; Grimstitch's cost reduction applies to
the supported Eviscerate.

Manaflare adds ten percentage points to Frostfire Bolt's Missile Barrage and
Fingers of Frost proc chances, provided the relevant talent is learned. Its
ten-point critical-strike bonus applies only to Frostfire Bolt and only while
Combustion is active. The proc-chance stacking interpretation follows the
tooltip; the client dummy effects do not expose the server script.

Grovekeeper Eclipse extends Insect Swarm from twelve to fifteen seconds. Its
two-second tick interval and damage per tick stay unchanged. The engine retains
its full-tick rule: seven ticks, then one second of remaining aura duration.
Whether Forever awards a fractional final tick still needs gameplay evidence.

Demonheart increases Life Tap's final mana return by twenty percent without
increasing its health cost. This multiplies the existing Improved Life Tap
calculation; server-side stacking is not exposed by the dummy effect.

Raiments of Conviction removes the current one-minute Devouring Plague
cooldown. The Shadow APL checks the existing DoT before recasting it.
