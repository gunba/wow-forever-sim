# Spell coverage

The client dump is a discovery source, not an implementation. A previous
class-spell export filtered only `ClassMask`; the four new Death ranks have
`ClassMask=0` and were absent from that export. They were separately known under
T37, but the automated coverage check did not catch their missing registration.
Discovery now includes both class masks and class skill lines, with a regression
test for that failure.

The audit joins client build **1.60.1.69893** spell names, effects, levels and
class-associated skill rows to the actual owner/pet spellbooks of all 171
profiles. It follows taught/triggered actions and known client/native aliases.
The subsequent Physical Ret row shares the already audited Paladin spellbook;
the original audit's profile count is historical.
Every unmatched damage family has an explicit disposition in
[`assets/spell_coverage_dispositions.json`](../assets/spell_coverage_dispositions.json).
The generated [full inventory](../artifacts/spell_coverage.json) includes individual
ranks and source-table hashes.

## Corrections and interpretation

- Death's four ranks are registered at levels 32/40/48/56. The damage coefficient,
  mana, shared cooldown and parent tooltip's maximum-health backlash are modeled.
  Early Demise adds 15/30 percentage points of Death crit in execute. T37 retains
  the unresolved script and server-interaction questions.
- Arcane Blast, Ice Lance and Wrack use legacy native action IDs; that is not
  evidence that their current client spells are missing.
- Trap child effects are recorded under trap placement actions.
- Pet teaching spells are not the pet's damage action IDs. Unselected pet
  families and Bear-only abilities do not appear in the sampled Cat/raid profiles.
- Many melee abilities intentionally register only the highest learned rank.
  Lower ranks are listed separately; they are not silently treated as supported.
- A registered spell is not a claim that every coefficient or interaction is
  verified. The in-game checklist remains applicable.

## Remaining coverage limits

These are explicit omissions, not a claim of a complete game spellbook:

| Area | Current limitation |
|---|---|
| Feral stealth | Prowl/opening-stealth state, Ravage and Pounce are not implemented. Their opener contribution is absent. |
| Mage proximity/control | Cone of Cold and Frost Nova are not implemented. The current caster profiles stand at 20 yards; in-range/AoE builds need these spells and their interactions. |
| Warlock proximity | Hellfire is not implemented. The current profiles stand at 20 yards. |
| Dwarf Priest | Chastise is not implemented; its Humanoid-only damage does not apply to the Dragonkin reference target. |
| Reactive/control attacks | Incoming-attack, parry, mana-burn, interrupt and control triggers are absent from the fixed DPS encounter. Individual omissions are listed in the inventory. |
| Hunter pets | New family teaching records exist, but trained availability and pet scaling remain T29 gaps. The existing model is not a complete new-pet simulator. |
| Retained records | Black Arrow, active Hunter Lacerate, Felfire, Lava Breath, Victory Rush and Hammer of the Righteous require current availability evidence rather than automatic import from retained records. |

The rankings compare the available profiles within these limits. They do not
establish that no unimplemented opener, in-range build or pet ability could win.

## Reproduce

```sh
forever-bench -spell-inventory -baseline-results /tmp/forever-profiles.json \
  -output /tmp/registered
python3 tools/data_watch/spell_coverage.py \
  --client-data /path/to/client-csv \
  --registered /tmp/registered.json \
  --output /tmp/coverage.json
```

`--learned` in `tools/data_watch/spell_client.py` now means class-associated
candidates. It also returns retained records; an association alone does not
prove a current trainer, racial or talent path.
