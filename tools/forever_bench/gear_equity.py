#!/usr/bin/env python3
"""Audit the selected modeled equipment's cost capacity and offensive allocations."""

import argparse
import csv
import json
from pathlib import Path
from statistics import mean
import sys

from build_display import BUILDS

sys.path.insert(0, str(Path(__file__).resolve().parents[1] / 'database'))
from forever_synthetic_gear import scale

STATS = ('Strength', 'Agility', 'Intellect', 'Spirit', 'Stamina', 'AttackPower',
         'RangedAttackPower', 'SpellPower', 'SpellDamage', 'HealingPower',
         'CritRating', 'HitRating', 'MP5')


def inspect(results: dict, catalog: dict) -> tuple[list[dict], list[dict]]:
    if results.get('GearScenario') != 'modeled-65-v2':
        raise ValueError('Only the current complete modeled roster can be audited')
    items = {item['id']: item for item in catalog['items']}
    classes = {key: cls for key, cls, _, _ in BUILDS}
    rows = []
    for result in results['Results']:
        equipped = [entry.get('id', 0) for entry in
                    result['BaselinePlayer']['equipment']['items']]
        unknown = set(equipped) - set(items) - {0}
        if unknown or any(item_id < 920_000_001 for item_id in equipped if item_id):
            raise ValueError(f"{result['Key']}/{result['Race']} contains real or older gear: {unknown}")
        if result['Hit']['Model'] != 'equipped-only-v2' or result['Hit']['RawHitDelta']:
            raise ValueError('This modeled profile received a paid hit exchange')
        selected = [items[item_id] for item_id in equipped if item_id]
        stats = {stat: sum(item['stats'].get(stat, 0) for item in selected) for stat in STATS}
        row = {
            'Build': result['Key'], 'Race': result['Race'],
            'Class': classes[result['Key']], 'DPS': round(result['DPS'], 2),
            'ItemCount': len(selected),
            'MeanItemLevel': round(mean(item['itemLevel'] for item in selected), 2),
            'EstimatedBudgetUnits': round(sum(item['modelBudgetRatio'] for item in selected), 3),
            'AssumedCapacityPoints': round(sum(
                item['modelBudgetRatio'] * scale(item['inventoryType'])
                for item in selected), 3),
            'MeleeHitPercent': round(result['Hit']['MeleeFinal'], 3),
            'SpellHitPercent': round(result['Hit']['SpellFinal'], 3),
            'WorstHitShortfallPercent': round(max(
                [0] + [entry['AdditionalPercent'] for entry in
                       result['Hit']['Requirements']]), 3),
            **stats,
        }
        rows.append(row)
    by_build = {}
    for row in rows:
        by_build.setdefault(row['Build'], []).append(row)
    totals = [{
        'Build': build, 'Races': len(group),
        **{f'Average{stat}': round(mean(row[stat] for row in group), 3)
           for stat in ('DPS', 'ItemCount', 'MeanItemLevel', 'EstimatedBudgetUnits',
                        'AssumedCapacityPoints', 'MeleeHitPercent', 'SpellHitPercent',
                        'WorstHitShortfallPercent', *STATS)},
    } for build, group in sorted(by_build.items())]
    return rows, totals


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--results', type=Path, required=True)
    parser.add_argument('--catalog', type=Path, default=Path(
        'assets/db_inputs/forever_synthetic_gear_v2.json'))
    parser.add_argument('--output', type=Path, required=True)
    args = parser.parse_args()
    rows, summary = inspect(json.loads(args.results.read_text()), json.loads(args.catalog.read_text()))
    args.output.mkdir(parents=True, exist_ok=True)
    for filename, records in [('selected_gear_equity.csv', rows),
                              ('spec_gear_equity.csv', summary)]:
        with (args.output / filename).open('w') as handle:
            writer = csv.DictWriter(handle, fieldnames=list(records[0]), lineterminator="\n")
            writer.writeheader()
            writer.writerows(records)
    (args.output / 'gear_equity.json').write_text(json.dumps({
        'warning': 'Estimated budget units depend on unverified stat prices; equal cost is not equal DPS.',
        'profiles': rows, 'builds': summary,
    }, indent=2) + '\n')
    print(f'Audited {len(rows)} fully modeled profiles and {len(summary)} builds')


if __name__ == '__main__':
    main()
