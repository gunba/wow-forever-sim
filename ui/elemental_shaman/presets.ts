import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	FirePowerBuff,
	Flask,
	Food,
	FrostPowerBuff,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { ElementalShaman_Options as ElementalShamanOptions } from '../core/proto/shaman.js';
import { SavedTalents } from '../core/proto/ui.js';
import DefaultAPLJson from './apls/default.apl.json';
import StormcallerAPLJson from './apls/stormcaller.apl.json';
import GearElementalJSON from './gear_sets/forever_elemental.gear.json';
import GearStormcallerJSON from './gear_sets/forever_stormcaller.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearElemental = PresetUtils.makePresetGear('Elemental', GearElementalJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearStormcaller = PresetUtils.makePresetGear('Stormcaller', GearStormcallerJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearElemental, GearStormcaller];
export const DefaultGear = GearElemental;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLDefault = PresetUtils.makePresetAPLRotation('Default', DefaultAPLJson);
export const APLStormcaller = PresetUtils.makePresetAPLRotation('Stormcaller', StormcallerAPLJson);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLDefault, APLStormcaller],
	[ClassicPhase.Phase2]: [],
	[ClassicPhase.Phase3]: [],
	[ClassicPhase.Phase4]: [],
	[ClassicPhase.Phase5]: [],
	[ClassicPhase.Phase6]: [],
};

export const DefaultAPL = APLDefault;

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

export const TalentsLevel60 = PresetUtils.makePresetTalents('Level 60', SavedTalents.create({ talentsString: '5505301500103031--503352001' }));

export const TalentsElemental = PresetUtils.makePresetTalents(
	'Elemental 31/6/14',
	SavedTalents.create({ talentsString: '5502301500123031-0500001-053050001' }),
);
export const TalentsStormcaller = PresetUtils.makePresetTalents(
	'Stormcaller 28/23/0',
	SavedTalents.create({ talentsString: '550032150010303-055030031004002' }),
);

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsLevel60, TalentsElemental, TalentsStormcaller],
	[ClassicPhase.Phase2]: [],
	[ClassicPhase.Phase3]: [],
	[ClassicPhase.Phase4]: [],
	[ClassicPhase.Phase5]: [],
	[ClassicPhase.Phase6]: [],
};

export const DefaultTalents = TalentsElemental;

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = ElementalShamanOptions.create({});

export const DefaultConsumes = Consumes.create({
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	firePowerBuff: FirePowerBuff.ElixirOfFirepower,
	frostPowerBuff: FrostPowerBuff.ElixirOfFrostPower,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodRunnTumTuberSurprise,
	// Not available until Phase 4
	// mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.CerebralCortexCompound,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	moonkinAura: true,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	// saygesFortune: SaygesFortune.SaygesDamage,
	// spiritOfZandalar: true,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfElements: true,
	stormstrike: true,
});

export const OtherDefaults = {
	distanceFromTarget: 20,
	profession2: Profession.Alchemy,
	profession1: Profession.Engineering,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Elemental', { gear: GearElemental, talents: TalentsElemental, rotation: APLDefault, options: DefaultOptions, distance: 20 }),
	PresetUtils.makePresetBuild('Stormcaller', {
		gear: GearStormcaller,
		talents: TalentsStormcaller,
		rotation: APLStormcaller,
		options: DefaultOptions,
		distance: 20,
	}),
];
