// The spell manifest, loaded for the browser.
//
// Internal action IDs can differ from Forever's spell IDs. The manifest supplies
// display names, current-ID mappings and audited tooltip notes without changing
// the IDs used by rotations or simulation metrics.

import commonJson from './common.json';
import coreJson from './core.json';
import druidJson from './druid.json';
import encountersJson from './encounters.json';
import hunterJson from './hunter.json';
import mageJson from './mage.json';
import paladinJson from './paladin.json';
import priestJson from './priest.json';
import rogueJson from './rogue.json';
import shamanJson from './shaman.json';
import warlockJson from './warlock.json';
import warriorJson from './warrior.json';

export type SpellSource = {
	ability: string;
	file: string;
	source: 'classic' | 'forever' | 'assumed' | 'unreviewed';
	foreverId?: number;
	tooltip?: string;
	note?: string;
	assumptions?: Array<string>;
};

const files: Array<Record<string, unknown>> = [
	commonJson,
	coreJson,
	druidJson,
	encountersJson,
	hunterJson,
	mageJson,
	paladinJson,
	priestJson,
	rogueJson,
	shamanJson,
	warlockJson,
	warriorJson,
];

const bySpellId = new Map<number, SpellSource>();
for (const file of files) {
	for (const [id, entry] of Object.entries(file)) {
		bySpellId.set(parseInt(id), entry as SpellSource);
	}
}

export function spellSource(spellId: number): SpellSource | undefined {
	return bySpellId.get(spellId);
}
