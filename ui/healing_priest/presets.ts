import * as PresetUtils from '../core/preset_utils.js';
import {
	Consumes,
	Debuffs,
	Flask,
	Food,
	IndividualBuffs,
	RaidBuffs,
	TristateEffect,
	UnitReference,
} from '../core/proto/common.js';
import {
	HealingPriest_Options as Options,
} from '../core/proto/priest.js';
import { SavedTalents } from '../core/proto/ui.js';
import DiscApl from './apls/disc.apl.json';
import HolyApl from './apls/holy.apl.json';
import BlankGear from './gear_sets/blank.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const GearBlank = PresetUtils.makePresetGear('Blank', BlankGear);

export const ROTATION_PRESET_DISC = PresetUtils.makePresetAPLRotation('Disc', DiscApl);
export const ROTATION_PRESET_HOLY = PresetUtils.makePresetAPLRotation('Holy', HolyApl);

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.
export const TalentsHolyHealer = PresetUtils.makePresetTalents('Holy 19/32/0', SavedTalents.create({ talentsString: '00520302133-2350510323000053' }));
export const TalentsDisciplineHealer = PresetUtils.makePresetTalents('Discipline 35/16/0', SavedTalents.create({ talentsString: '005205031332011531-03505003' }));

export const DefaultOptions = Options.create({
	useInnerFire: true,

	powerInfusionTarget: UnitReference.create(),
});

export const DefaultConsumes = Consumes.create({
	flask: Flask.FlaskUnknown,
	food: Food.FoodUnknown,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectRegular,
	arcaneBrilliance: true,
	divineSpirit: true,
	moonkinAura: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
});

export const DefaultDebuffs = Debuffs.create({
});
