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
	SpellPowerBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { SmitePriest_Options as Options } from '../core/proto/priest.js';
import { SavedTalents } from '../core/proto/ui.js';
import LaunchAPL from './apls/launch.apl.json';
import GearSmiteJSON from './gear_sets/forever_smite.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearSmite = PresetUtils.makePresetGear('Smite', GearSmiteJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearSmite];
export const DefaultGear = GearSmite;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLLaunch = PresetUtils.makePresetAPLRotation('Smite', LaunchAPL);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLLaunch],
};

export const DefaultAPL = APLLaunch;

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Thirty-one points of Discipline for Power Infusion, twenty of Holy for the pieces
// that scale Smite. Power in Light and Searing Light are the two that make the build:
// the first pays for keeping Holy Fire on the target, the second pays for casting it.
// Divine Fury is what makes Smite castable at all.
export const TalentsLaunch = PresetUtils.makePresetTalents('Smite', SavedTalents.create({ talentsString: '515330031305001001-30505113002' }));

export const TalentsSmite = PresetUtils.makePresetTalents('Smite 31/17/3', SavedTalents.create({ talentsString: '515030031305001031-00505023002-003' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsLaunch, TalentsSmite],
};

export const DefaultTalents = TalentsSmite;

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
	PresetUtils.makePresetBuild('Smite', { gear: GearSmite, talents: TalentsSmite, rotation: APLLaunch, options: DefaultOptions, distance: 20 }),
];
