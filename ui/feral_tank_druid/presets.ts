import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	ArmorElixir,
	AttackPowerBuff,
	Consumes,
	Debuffs,
	Flask,
	Food,
	HealthElixir,
	IndividualBuffs,
	Potions,
	RaidBuffs,
	StrengthBuff,
	TristateEffect,
	UnitReference,
	ZanzaBuff,
} from '../core/proto/common.js';
import { FeralTankDruid_Options as DruidOptions } from '../core/proto/druid.js';
import { SavedTalents } from '../core/proto/ui.js';
import ForeverBearApl from './apls/forever_bear.apl.json';
import LaunchGearJSON from './gear_sets/launch.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearLaunch = PresetUtils.makePresetGear('Modeled level 65', LaunchGearJSON);

export const GearPresets = {
	[ClassicPhase.Phase1]: [GearLaunch],
};

export const DefaultGear = GearPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLForeverBear = PresetUtils.makePresetAPLRotation('Forever Bear', ForeverBearApl);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLForeverBear],
};

export const DefaultAPL = APLPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsBearTank = PresetUtils.makePresetTalents('Bear Tank 1/40/10', SavedTalents.create({ talentsString: '01-5002232123132210551-055' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsBearTank],
};

export const DefaultTalents = TalentPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = DruidOptions.create({
	innervateTarget: UnitReference.create(),
	startingRage: 0,
});

// No weapon imbue: sharpening stones and Windfury do nothing for a bear's paws.
export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	armorElixir: ArmorElixir.ElixirOfSuperiorDefense,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultPotion: Potions.GreaterStoneshieldPotion,
	dragonBreathChili: true,
	food: Food.FoodSmokedDesertDumpling,
	flask: Flask.FlaskOfTheTitans,
	healthElixir: HealthElixir.ElixirOfFortitude,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	battleShout: TristateEffect.TristateEffectImproved,
	devotionAura: TristateEffect.TristateEffectRegular,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	stoneskinTotem: TristateEffect.TristateEffectRegular,
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
