import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	Consumes,
	Debuffs,
	Flask,
	Food,
	IndividualBuffs,
	ManaRegenElixir,
	PartyBuffs,
	Potions,
	Profession,
	RaidBuffs,
	SpellPowerBuff,
	TristateEffect,
	UnitReference,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { BalanceDruid_Options as BalanceDruidOptions } from '../core/proto/druid.js';
import { SavedTalents } from '../core/proto/ui.js';
import Balance from './apls/balance.apl.json';
import LaunchAPL from './apls/launch.apl.json';
import GearBalanceJSON from './gear_sets/forever_balance.gear.json';
// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearBalance = PresetUtils.makePresetGear('Balance', GearBalanceJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearBalance];
export const DefaultGear = GearBalance;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

// export const APLP1Balance = PresetUtils.makePresetAPLRotation('Balance', P1APL);
export const APLLaunch = PresetUtils.makePresetAPLRotation('Launch', LaunchAPL);
export const DefaultBalance = PresetUtils.makePresetAPLRotation('Default Balance', Balance);
export const APLPresets = {
	[ClassicPhase.Phase1]: [APLLaunch, DefaultBalance],
};

export const DefaultAPL = APLLaunch;

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

export const TalentsP1Balance = PresetUtils.makePresetTalents('Balance', SavedTalents.create({ talentsString: '5532220115001351--505302' }));

export const TalentsMoonkin = PresetUtils.makePresetTalents('Moonkin 38/0/13', SavedTalents.create({ talentsString: '5232220115501351--505003' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsP1Balance, TalentsMoonkin],
};

export const DefaultTalents = TalentsMoonkin;

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = BalanceDruidOptions.create({
	innervateTarget: UnitReference.create(),
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.MajorManaPotion,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodRunnTumTuberSurprise,
	mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,

	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	zanzaBuff: ZanzaBuff.CerebralCortexCompound,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	moonkinAura: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
});

export const DefaultPartyBuffs = PartyBuffs.create({});

export const DefaultDebuffs = Debuffs.create({
	faerieFire: true,
	judgementOfWisdom: true,
	stormstrike: true,
});

export const OtherDefaults = {
	distanceFromTarget: 20,
	profession1: Profession.Engineering,
	profession2: Profession.Tailoring,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Balance', { gear: GearBalance, talents: TalentsMoonkin, rotation: APLLaunch, options: DefaultOptions, distance: 20 }),
];
