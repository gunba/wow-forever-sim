import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	Explosive,
	FirePowerBuff,
	Flask,
	Food,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	RaidBuffs,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { PaladinAura, PaladinOptions as RetributionPaladinOptions, PaladinSeal } from '../core/proto/paladin.js';
import { SavedTalents } from '../core/proto/ui.js';
import APLBasicRetJson from './apls/basic_ret.apl.json';
import GearRetributionJSON from './gear_sets/forever_retribution.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearRetribution = PresetUtils.makePresetGear('Retribution', GearRetributionJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearRetribution];
export const DefaultGear = GearRetribution;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLBasicRet = PresetUtils.makePresetAPLRotation('Basic Ret', APLBasicRetJson);

export const APLPresets = {
	[ClassicPhase.Phase1]: [],
	[ClassicPhase.Phase2]: [],
	[ClassicPhase.Phase3]: [],
	[ClassicPhase.Phase4]: [APLBasicRet],
	[ClassicPhase.Phase5]: [],
};

export const DefaultAPL = APLBasicRet;

///////////////////////////////////////////////////////////////////////////
//                                 Talent presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsRetribution = PresetUtils.makePresetTalents('Retribution 13/6/32', SavedTalents.create({ talentsString: '55003-231-05225331001330301' }));
export const TalentsPhysicalRetribution = PresetUtils.makePresetTalents(
	'Physical Ret 12/6/33',
	SavedTalents.create({ talentsString: '52003002-231-05225331001330311' }),
);

// Shockadin (Holy Shock Retribution), wowsims/forever's Holy Shock preset.
export const TalentsShockadin = PresetUtils.makePresetTalents('Shockadin 23/0/28', SavedTalents.create({ talentsString: '550030032101012--052252300012303' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [],
	[ClassicPhase.Phase2]: [],
	[ClassicPhase.Phase3]: [],
	[ClassicPhase.Phase4]: [TalentsRetribution, TalentsPhysicalRetribution, TalentsShockadin],
};

export const DefaultTalents = TalentsRetribution;

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = RetributionPaladinOptions.create({
	aura: PaladinAura.NoPaladinAura,
	primarySeal: PaladinSeal.Command,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	boglingRoot: false,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	fillerExplosive: Explosive.ExplosiveUnknown,
	firePowerBuff: FirePowerBuff.ElixirOfGreaterFirepower,
	food: Food.FoodBlessSunfruit,
	flask: Flask.FlaskOfSupremePower,
	//mainHandImbue: WeaponImbue.WildStrikes,
	//offHandImbue: WeaponImbue.MagnificentTrollshine,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

export const PhysicalRetConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	food: Food.FoodBlessSunfruit,
	mainHandImbue: WeaponImbue.DenseSharpeningStone,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
	moonkinAura: true,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	faerieFire: true,
	giftOfArthas: true,
	sunderArmor: true,
	judgementOfWisdom: true,
	judgementOfTheCrusader: TristateEffect.TristateEffectImproved,
});

export const OtherDefaults = {
	profession1: Profession.Engineering,
	profession2: Profession.Enchanting,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Retribution', {
		gear: GearRetribution,
		talents: TalentsRetribution,
		rotation: APLBasicRet,
		options: DefaultOptions,
		distance: 5,
	}),
	PresetUtils.makePresetBuild('Physical Ret', {
		gear: GearRetribution,
		consumes: PhysicalRetConsumes,
		talents: TalentsPhysicalRetribution,
		rotation: APLBasicRet,
		options: DefaultOptions,
		distance: 5,
	}),
];
