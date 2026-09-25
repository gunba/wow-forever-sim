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
	Profession,
	Race,
	RaidBuffs,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';
import { TankWarrior_Options as TankWarriorOptions, WarriorShout, WarriorStance } from '../core/proto/warrior.js';
import APLNoReckJSON from './apls/dps_no_reck.apl.json';
import APLProtectionJSON from './apls/protection.apl.json';
import APLReckJSON from './apls/dps_reck.apl.json';
import LaunchGearJSON from './gear_sets/launch.gear.json';
import P0BISGear from './gear_sets/p0.bis.gear.json';
import P1BISGear from './gear_sets/p1.bis.gear.json';

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearLaunch = PresetUtils.makePresetGear('Launch', LaunchGearJSON);
export const GearP0BIS = PresetUtils.makePresetGear('Pre-BiS', P0BISGear);
export const GearP1BIS = PresetUtils.makePresetGear('P1 BiS', P1BISGear);

export const GearPresets = {
	[ClassicPhase.Phase1]: [GearLaunch, GearP0BIS, GearP1BIS],
};

export const DefaultGear = GearPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLProtection = PresetUtils.makePresetAPLRotation('Protection', APLProtectionJSON);
export const AplReck = PresetUtils.makePresetAPLRotation('DPS (With Reck)', APLReckJSON);
export const APLNoReck = PresetUtils.makePresetAPLRotation('DPS (No Reck)', APLNoReckJSON);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLProtection, APLNoReck, AplReck],
};

export const DefaultAPL = APLPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1Prot = PresetUtils.makePresetTalents('Protection', SavedTalents.create({ talentsString: '31--552531233330012351' }));

export const TalentsProtection = PresetUtils.makePresetTalents('Protection 1/0/50', SavedTalents.create({ talentsString: '1--552531233331212351' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsP1Prot, TalentsProtection],
};

export const DefaultTalents = TalentPresets[ClassicPhase.Phase1][0];

export const PresetBuildTanky = PresetUtils.makePresetBuild('Tanky', { gear: DefaultGear, talents: TalentsP1Prot, rotation: DefaultAPL });

///////////////////////////////////////////////////////////////////////////
//                                 Options Presets
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = TankWarriorOptions.create({
	queueDelay: 250,
	startingRage: 0,
	shout: WarriorShout.WarriorShoutBattle,
	stance: WarriorStance.WarriorStanceDefensive,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	armorElixir: ArmorElixir.ElixirOfSuperiorDefense,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultPotion: Potions.MightyRagePotion,
	dragonBreathChili: true,
	food: Food.FoodSmokedDesertDumpling,
	flask: Flask.FlaskOfTheTitans,
	healthElixir: HealthElixir.ElixirOfFortitude,
	mainHandImbue: WeaponImbue.ElementalSharpeningStone,
	offHandImbue: WeaponImbue.ElementalSharpeningStone,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	battleShout: TristateEffect.TristateEffectImproved,
	devotionAura: TristateEffect.TristateEffectRegular,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
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
	insectSwarm: true,
	sunderArmor: true,
});

export const OtherDefaults = {
	profession1: Profession.Blacksmithing,
	profession2: Profession.Enchanting,
	race: Race.RaceHuman,
};
