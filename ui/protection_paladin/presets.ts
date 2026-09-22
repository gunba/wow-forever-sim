import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	ArmorElixir,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	Explosive,
	FirePowerBuff,
	Flask,
	Food,
	HealthElixir,
	IndividualBuffs,
	Potions,
	Profession,
	RaidBuffs,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { Blessings, PaladinAura, PaladinOptions as ProtectionPaladinOptions, PaladinSeal } from '../core/proto/paladin.js';
import { SavedTalents } from '../core/proto/ui.js';
import APLBasicProtectionJson from './apls/basic_prot.apl.json';
import APLP5ProtJson from './apls/p5prot.apl.json';
import BlankGear from './gear_sets/blank.gear.json';
import LaunchGearJSON from './gear_sets/launch.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearLaunch = PresetUtils.makePresetGear('Launch', LaunchGearJSON);
export const GearBlank = PresetUtils.makePresetGear('Blank', BlankGear);

export const GearPresets = {
	[ClassicPhase.Phase1]: [GearLaunch],
};

export const DefaultGear = GearPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLP5Prot = PresetUtils.makePresetAPLRotation('P5 Prot', APLP5ProtJson);
export const APLBasicProt = PresetUtils.makePresetAPLRotation('Basic Prot', APLBasicProtectionJson);

export const APLPresets = {
	[ClassicPhase.Phase1]: [],
	[ClassicPhase.Phase2]: [],
	[ClassicPhase.Phase3]: [],
	[ClassicPhase.Phase4]: [APLBasicProt, APLP5Prot],
	[ClassicPhase.Phase5]: [APLP5Prot, APLP5Prot],
};

export const DefaultAPL = APLPresets[ClassicPhase.Phase5][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const P4ProtTalents = {
	name: 'P4 Prot',
	data: SavedTalents.create({
		talentsString: '052003003-5530513321301501',
	}),
};

export const P5ProtTalents = {
	name: 'P5 Prot',
	data: SavedTalents.create({
		talentsString: '055003-5530513321301501',
	}),
};

export const TalentsProtection = PresetUtils.makePresetTalents('Protection 0/45/6', SavedTalents.create({ talentsString: '-5532513321301551-15' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [],
	[ClassicPhase.Phase2]: [],
	[ClassicPhase.Phase3]: [],
	[ClassicPhase.Phase4]: [P4ProtTalents],
	[ClassicPhase.Phase5]: [P5ProtTalents, TalentsProtection],
};

export const DefaultTalents = TalentPresets[ClassicPhase.Phase5][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = ProtectionPaladinOptions.create({
	aura: PaladinAura.NoPaladinAura,
	primarySeal: PaladinSeal.Righteousness,
	personalBlessing: Blessings.BlessingOfSanctuary,
	righteousFury: true,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	healthElixir: HealthElixir.ElixirOfFortitude,
	armorElixir: ArmorElixir.ElixirOfSuperiorDefense,
	defaultPotion: Potions.GreaterStoneshieldPotion,
	dragonBreathChili: true,
	food: Food.FoodTenderWolfSteak,
	flask: Flask.FlaskOfTheTitans,
	firePowerBuff: FirePowerBuff.ElixirOfGreaterFirepower,
	fillerExplosive: Explosive.ExplosiveDenseDynamite,
	//mainHandImbue: WeaponImbue.WildStrikes,
	//offHandImbue: WeaponImbue.MagnificentTrollshine,

	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredDemonicRune,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	faerieFire: true,
	giftOfArthas: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	judgementOfWisdom: true,
	judgementOfTheCrusader: TristateEffect.TristateEffectImproved,
});

export const OtherDefaults = {
	distanceFromTarget: 5, // Max melee range
	profession1: Profession.Blacksmithing,
	profession2: Profession.Engineering,
};
