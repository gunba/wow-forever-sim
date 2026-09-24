#!/usr/bin/env python3
"""Generate the separately labelled, provisional level-65 modeled equipment pool.

Real Forever items remain in forever_gear_catalog.json; model hypotheses live here.
The full-passive trinket variants are sensitivity references, not benchmark options.
"""
from __future__ import annotations

import argparse
import json
from pathlib import Path

from forever_synthetic_gear import BUILD, BODY_SLOTS, BY_ID, ITEMS, ROOT, proposals

ID_START = 910_000_000


def icon_reference(source: dict, slot: int, material: str) -> dict:
    target_class = 4 if slot in (1,2,3,5,6,7,8,9,10,11,12,14,16,23) else 2
    armor_material = {'Cloth': 1, 'Leather': 2, 'Mail': 3, 'Plate': 4}.get(material)
    if source['class'] == target_class and (
        armor_material is None or source['subclass'] == armor_material
    ) and (
        source['inventoryType'] == slot or slot == 5 and source['inventoryType'] == 20
    ):
        return source
    candidates = [
        item for item in ITEMS
        if item['inventoryType'] == slot or slot == 5 and item['inventoryType'] == 20
        if item['icon'] and
        (armor_material is None or item['subclass'] == armor_material) and
        item['class'] == target_class
    ]
    if not candidates:
        return source
    return min(candidates, key=lambda item: (
        abs(item['itemLevel'] - 65), item['quality'] != 4,
        not item.get('clientStatRecord', False), item['id'],
    ))


def record(proposal: dict, index: int) -> dict:
    source = BY_ID[proposal['SourceID']]
    slot = proposal['SlotID']
    item_id = ID_START + index
    archetype = proposal['Archetype']
    material = proposal['Material']
    speed = proposal['Speed']
    label = f"Modeled: {archetype} — {proposal['Slot']}"
    if material in ('Cloth', 'Leather', 'Mail', 'Plate'):
        label += f' ({material})'
    if speed is not None:
        label += f' ({speed:g}s)'
    benchmark_eligible = not (proposal['SlotID'] == 12 and
                              proposal['Policy'].startswith('100%'))
    if not benchmark_eligible:
        label += ' [sensitivity only]'
    # The projected caster dagger takes its damage/speed/type from the weapon
    # reference and its allocations from the separate documented offhand.
    material = {'Cloth': 1, 'Leather': 2, 'Mail': 3, 'Plate': 4}.get(proposal['Material'])
    body = slot in BODY_SLOTS or slot == 16
    armor = slot in (1,2,3,5,6,7,8,9,10,11,12,14,16,23)
    icon_source = icon_reference(source, slot, proposal['Material'])
    result = {
        'id': item_id, 'name': label, 'icon': icon_source['icon'],
        'itemLevel': 65, 'requiredLevel': 60, 'quality': 4,
        'class': 4 if armor else 2,
        'subclass': (material if body and material else
                     6 if slot == 14 else 0 if armor else source['subclass']),
        'inventoryType': slot, 'classMask': 0,
        'raceMasks': [0, 0], 'requiredSkill': 0, 'requiredSkillRank': 0,
        'requiredAbility': 0, 'maxCount': 0, 'limitCategory': 0,
        'setId': 0, 'sources': [{'kind': 'synthetic'}], 'effects': [],
        'stats': proposal['Stats'], 'armor': proposal['Armor'] or 0,
        'clientStatRecord': False, 'statSource': 'modeled-client-table-projection',
        'synthetic': True, 'benchmarkEligible': benchmark_eligible,
        'modelID': proposal['ID'], 'modelReferenceItemID': proposal['SourceID'],
        'modelSecondaryReferenceID': proposal.get('ExtraSourceID', 0),
        'modelIconReferenceItemID': icon_source['id'],
        'modelSourceSlot': proposal['SourceSlot'], 'modelPolicy': proposal['Policy'],
        'modelConfidence': proposal['Confidence'], 'modelNotes': proposal['Notes'],
        'modelAllocations': proposal['Allocation'],
        'modelBudgetRatio': proposal['ProposedBudget'],
    }
    if speed is not None:
        result['weapon'] = {'min': proposal['Min'], 'max': proposal['Max'],
                            'speed': speed, 'school': source['weapon'].get('school', 0)}
    return result


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--output', type=Path,
                        default=ROOT/'assets/db_inputs/forever_synthetic_gear.json')
    parser.add_argument('--ui-metadata', type=Path,
                        default=ROOT/'ui/core/forever_synthetic_item_metadata.json')
    args = parser.parse_args()
    plans = proposals()
    items = [record(plan, i+1) for i, plan in enumerate(plans)]
    real_ids = set(BY_ID)
    assert not real_ids.intersection({item['id'] for item in items})
    assert all(item['itemLevel'] == 65 and item['requiredLevel'] <= 60 for item in items)
    assert all(not item['effects'] and not item['setId'] and
               not item['requiredSkill'] for item in items)
    payload = {
        'schemaVersion': 1, 'clientBuild': BUILD,
        'method': 'pinned client-table projection; special-stat cost hypotheses are not verified',
        'benchmarkPolicy': 'maximum modeled Stamina share 20%; passive trinket 80%-allocation case',
        'sourceCatalog': 'assets/db_inputs/forever_gear_catalog.json',
        'items': items,
    }
    args.output.parent.mkdir(parents=True, exist_ok=True)
    args.output.write_text(json.dumps(payload, indent=2, ensure_ascii=False)+'\n')
    ui_metadata = {
        str(item['id']): {
            'name': item['name'],
            'sourceItemId': item['modelReferenceItemID'],
            'sourceItemName': BY_ID[item['modelReferenceItemID']]['name'],
            'stats': item['stats'],
            'armor': item['armor'],
            'weapon': item.get('weapon'),
            'confidence': item['modelConfidence'],
            'notes': item['modelNotes'],
            'rankingEligible': item['benchmarkEligible'],
        } for item in items
    }
    args.ui_metadata.parent.mkdir(parents=True, exist_ok=True)
    args.ui_metadata.write_text(json.dumps(ui_metadata, separators=(',', ':'), ensure_ascii=False)+'\n')
    print(f"Generated {len(items)} labeled projections; "
          f"{sum(item['benchmarkEligible'] for item in items)} benchmark candidates")


if __name__ == '__main__':
    main()
