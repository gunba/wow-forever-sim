import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	ArmorElixir,
	AttackPowerBuff,
	Consumes,
	Debuffs,
	Food,
	HealthElixir,
	IndividualBuffs,
	Potions,
	Profession,
	Race,
	RaidBuffs,
	SapperExplosive,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';
import { Warrior_Options as WarriorOptions, WarriorShout, WarriorStance } from '../core/proto/warrior.js';
import APLFuryJSON from './apls/forever_fury.apl.json';
import APLArmsJSON from './apls/forever_arms.apl.json';
import GearFuryJSON from './gear_sets/forever_fury.gear.json';
import GearArmsJSON from './gear_sets/forever_arms.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearFury = PresetUtils.makePresetGear('Fury', GearFuryJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearArms = PresetUtils.makePresetGear('Arms', GearArmsJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearFury, GearArms];
export const DefaultGear = GearFury;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLFury = PresetUtils.makePresetAPLRotation('Fury', APLFuryJSON);
export const APLArms = PresetUtils.makePresetAPLRotation('Arms', APLArmsJSON);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLFury, APLArms],
};

export const DefaultAPLs = [APLPresets[ClassicPhase.Phase1][0]];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1DPS = PresetUtils.makePresetTalents('Fury 17/34/0', SavedTalents.create({ talentsString: '20305113002-050520035151010051' }));

export const TalentsArms = PresetUtils.makePresetTalents('Arms 34/17/0', SavedTalents.create({ talentsString: '20305213132515001-550500000002' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsP1DPS, TalentsArms],
};

export const DefaultTalents = TalentsP1DPS;

///////////////////////////////////////////////////////////////////////////
//                                 Options Presets
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = WarriorOptions.create({
	queueDelay: 250,
	startingRage: 50,
	shout: WarriorShout.WarriorShoutBattle,
	stance: WarriorStance.WarriorStanceBerserker,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	armorElixir: ArmorElixir.ElixirOfSuperiorDefense,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultPotion: Potions.MightyRagePotion,
	dragonBreathChili: true,
	food: Food.FoodSmokedDesertDumpling,
	healthElixir: HealthElixir.ElixirOfFortitude,
	mainHandImbue: WeaponImbue.Windfury,
	offHandImbue: WeaponImbue.ElementalSharpeningStone,
	sapperExplosive: SapperExplosive.SapperGoblinSapper,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	battleShout: TristateEffect.TristateEffectImproved,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	giftOfArthas: true,
	sunderArmor: true,
});

export const OtherDefaults = {
	profession1: Profession.Engineering,
	profession2: Profession.Engineering,
	race: Race.RaceHuman,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Fury', { gear: GearFury, talents: TalentsP1DPS, rotation: APLFury, options: DefaultOptions, distance: 5 }),
	PresetUtils.makePresetBuild('Arms', { gear: GearArms, talents: TalentsArms, rotation: APLArms, options: DefaultOptions, distance: 5 }),
];
