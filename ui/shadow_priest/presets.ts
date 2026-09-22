import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	Conjured,
	Consumes,
	Debuffs,
	Flask,
	Food,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	RaidBuffs,
	ShadowPowerBuff,
	SpellPowerBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { ShadowPriest_Options as Options } from '../core/proto/priest.js';
import { SavedTalents } from '../core/proto/ui.js';
import P1APL from './apls/p1.apl.json';
import GearShadowJSON from './gear_sets/forever_shadow.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearShadow = PresetUtils.makePresetGear('Shadow', GearShadowJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearShadow];
export const DefaultGear = GearShadow;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLP1Shadow = PresetUtils.makePresetAPLRotation('Shadow', P1APL);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLP1Shadow],
};

export const DefaultAPL = APLP1Shadow;

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1Shadow = PresetUtils.makePresetTalents('Shadow 20/0/31', SavedTalents.create({ talentsString: '305030001305--504020501201302051' }));

export const TalentsShadow = PresetUtils.makePresetTalents('Shadow 15/0/36', SavedTalents.create({ talentsString: '0253000311--550022501201302251' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsP1Shadow, TalentsShadow],
};

export const DefaultTalents = TalentsP1Shadow;

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = Options.create({});

export const DefaultConsumes = Consumes.create({
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodRunnTumTuberSurprise,
	mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,

	shadowPowerBuff: ShadowPowerBuff.ElixirOfShadowPower,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	zanzaBuff: ZanzaBuff.CerebralCortexCompound,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	manaSpringTotem: TristateEffect.TristateEffectImproved,
	moonkinAura: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
});

export const DefaultDebuffs = Debuffs.create({
	judgementOfWisdom: true,
});

export const OtherDefaults = {
	channelClipDelay: 100,
	distanceFromTarget: 20,
	profession1: Profession.Engineering,
	profession2: Profession.Enchanting,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Shadow', { gear: GearShadow, talents: TalentsP1Shadow, rotation: APLP1Shadow, options: DefaultOptions, distance: 20 }),
];
