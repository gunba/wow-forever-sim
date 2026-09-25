#!/usr/bin/env python3
"""Turn saved tank-sim defaults into native raid requests without changing their setup."""

import argparse
import copy
import json
from pathlib import Path


def request(settings, iterations, seed, targets):
    encounter = copy.deepcopy(settings['encounter'])
    if targets != 1:
        encounter['duration'] = 90
        encounter['targets'] = [copy.deepcopy(encounter['targets'][0]) for _ in range(targets)]
        for target in encounter['targets']:
            target['minBaseDamage'] = 1500
            target.pop('id', None)
            target.pop('name', None)
    return {
        'raid': {
            'parties': [{'players': [copy.deepcopy(settings['player'])], 'buffs': settings.get('partyBuffs', {})}],
            'buffs': settings.get('raidBuffs', {}),
            'debuffs': settings.get('debuffs', {}),
            'tanks': settings['tanks'],
        },
        'encounter': encounter,
        'simOptions': {
            'iterations': iterations,
            'randomSeed': seed,
            'ruleset': settings['settings']['ruleset'],
        },
    }


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--settings-dir', type=Path, required=True)
    parser.add_argument('--output-dir', type=Path, required=True)
    parser.add_argument('--iterations', type=int, default=5000)
    parser.add_argument('--seed', type=int, default=20261993)
    parser.add_argument('--targets', type=int, choices=(1, 3), default=1)
    args = parser.parse_args()
    args.output_dir.mkdir(parents=True, exist_ok=True)
    for role in ('tank_warrior', 'protection_paladin', 'feral_tank_druid'):
        settings = json.loads((args.settings_dir / f'{role}.settings.json').read_text())
        value = request(settings, args.iterations, args.seed, args.targets)
        path = args.output_dir / f'{role}_{args.targets}t_{args.seed}.request.json'
        path.write_text(json.dumps(value, indent=2) + '\n')
        print(path)


if __name__ == '__main__':
    main()
