import {
	Consumes,
	Flask,
	Food,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';

import {
	PaladinAura,
	PaladinOptions as HolyPaladinOptions,
} from '../core/proto/paladin.js';

import * as PresetUtils from '../core/preset_utils.js';

import BlankGear from './gear_sets/blank.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const DefaultGear = PresetUtils.makePresetGear('Blank', BlankGear);

// Legal level-60 Holy allocation. This is not a validated healer optimum;
// full healing-model coverage remains separate from the tank/DPS benchmarks.

export const StandardTalents = {
	name: 'Standard',
	data: SavedTalents.create({
		talentsString: '05323213225121051-55021',
	}),
};

export const TalentsHolyHealer = PresetUtils.makePresetTalents('Holy 38/13/0', SavedTalents.create({ talentsString: '05323213225121051-55021' }));

export const DefaultOptions = HolyPaladinOptions.create({
	aura: PaladinAura.DevotionAura,
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskUnknown,
	food: Food.FoodUnknown,
});
