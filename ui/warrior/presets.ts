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
import APLFurySpearingJSON from './apls/forever_fury_spearing.apl.json';
import APLFurySunderJSON from './apls/forever_fury_sunder.apl.json';
import APLArmsJSON from './apls/forever_arms.apl.json';
import APLFuryTwoHandJSON from './apls/forever_fury_2h.apl.json';
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
export const APLFurySpearing = PresetUtils.makePresetAPLRotation('Fury (Giants/Dragonkin)', APLFurySpearingJSON);
APLFurySpearing.tooltip = 'Spearing Strike build for Giant or Dragonkin targets. Not a general-purpose Fury rotation.';
export const APLFurySunder = PresetUtils.makePresetAPLRotation('Fury (Sunder)', APLFurySunderJSON);
APLFurySunder.tooltip = 'Disable external Sunder Armor and Expose Armor, or load the exact ranked Fury (Sunder) profile.';
export const APLArms = PresetUtils.makePresetAPLRotation('Arms', APLArmsJSON);
export const APLFuryTwoHand = PresetUtils.makePresetAPLRotation('2H Bloodthirst', APLFuryTwoHandJSON);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLFury, APLFurySpearing, APLFurySunder, APLArms, APLFuryTwoHand],
};

export const DefaultAPLs = [APLPresets[ClassicPhase.Phase1][0]];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1DPS = PresetUtils.makePresetTalents('Fury 17/34/0', SavedTalents.create({ talentsString: '20305113002-050520035151010051' }));
export const TalentsFurySpearing = PresetUtils.makePresetTalents(
	'Fury Spearing Strike 18/33/0',
	SavedTalents.create({ talentsString: '20305113102-050520035150010051' }),
);

export const TalentsArms = PresetUtils.makePresetTalents('Arms 34/17/0', SavedTalents.create({ talentsString: '20305213132515001-550500000002' }));
export const TalentsFuryTwoHand = PresetUtils.makePresetTalents(
	'2H Bloodthirst 20/31/0',
	SavedTalents.create({ talentsString: '30304203032-050500320051310051' }),
);

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsP1DPS, TalentsFurySpearing, TalentsArms, TalentsFuryTwoHand],
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
	mainHandImbue: WeaponImbue.ElementalSharpeningStone,
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
	PresetUtils.makePresetBuild('Fury (Giants/Dragonkin)', {
		gear: GearFury, talents: TalentsFurySpearing, rotation: APLFurySpearing, options: DefaultOptions, distance: 5,
	}),
	PresetUtils.makePresetBuild('Arms', { gear: GearArms, talents: TalentsArms, rotation: APLArms, options: DefaultOptions, distance: 5 }),
];
