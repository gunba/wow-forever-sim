import { ClassicPhase } from '../core/constants/other.js';
import { foreverTankSupport } from '../core/forever_tank_support.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	ArmorElixir,
	AttackPowerBuff,
	Class,
	Consumes,
	Explosive,
	FirePowerBuff,
	Flask,
	Food,
	HealthElixir,
	Potions,
	Profession,
	Race,
	SpellPowerBuff,
	StrengthBuff,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { Blessings, PaladinAura, PaladinOptions as ProtectionPaladinOptions, PaladinSeal } from '../core/proto/paladin.js';
import { SavedTalents } from '../core/proto/ui.js';
import APLForeverProtJson from './apls/forever_protection.apl.json';
import BlankGear from './gear_sets/blank.gear.json';
import LaunchGearJSON from './gear_sets/launch.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearLaunch = PresetUtils.makePresetGear('Modeled level 65', LaunchGearJSON);
export const GearBlank = PresetUtils.makePresetGear('Blank', BlankGear);

export const GearPresets = {
	[ClassicPhase.Phase1]: [GearLaunch],
};

export const DefaultGear = GearPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLForeverProt = PresetUtils.makePresetAPLRotation('Forever Protection', APLForeverProtJson);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLForeverProt],
};

export const DefaultAPL = APLPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsProtection = PresetUtils.makePresetTalents('Protection 0/42/9', SavedTalents.create({ talentsString: '-5520513321301551-15102' }));

// Race-specific selected recipes; full ranked profiles also retain exact gear and APL.
export const TalentsProtectionUndead = PresetUtils.makePresetTalents('Protection · Undead', SavedTalents.create({ talentsString: '-5520513321301551-15102' }), { customCondition: player => player.getRace() === Race.RaceUndead });
export const TalentsProtectionHuman = PresetUtils.makePresetTalents('Protection · Human', SavedTalents.create({ talentsString: '-5521513321301551-152' }), { customCondition: player => player.getRace() === Race.RaceHuman });
export const TalentsProtectionDwarf = PresetUtils.makePresetTalents('Protection · Dwarf', SavedTalents.create({ talentsString: '-5520513321301551-15003' }), { customCondition: player => player.getRace() === Race.RaceDwarf });

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsProtectionUndead, TalentsProtectionHuman, TalentsProtectionDwarf],
};

export const DefaultTalents = TalentPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = ProtectionPaladinOptions.create({
	aura: PaladinAura.DevotionAura,
	primarySeal: PaladinSeal.Righteousness,
	personalBlessing: Blessings.BlessingUnknown,
	righteousFury: true,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	healthElixir: HealthElixir.ElixirOfFortitude,
	armorElixir: ArmorElixir.ElixirOfSuperiorDefense,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	food: Food.FoodTenderWolfSteak,
	flask: Flask.FlaskOfTheTitans,
	firePowerBuff: FirePowerBuff.ElixirOfGreaterFirepower,
	fillerExplosive: Explosive.ExplosiveDenseDynamite,
	mainHandImbue: WeaponImbue.BrilliantWizardOil,

	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
});

const Support = foreverTankSupport(Class.ClassPaladin);
export const DefaultRaidBuffs = Support.raid;
export const DefaultPartyBuffs = Support.party;
export const DefaultIndividualBuffs = Support.player;
export const DefaultDebuffs = Support.debuffs;

export const OtherDefaults = {
	distanceFromTarget: 5, // Max melee range
	profession1: Profession.Blacksmithing,
	profession2: Profession.Engineering,
};
