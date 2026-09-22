import { ClassicPhase } from '../core/constants/other';
import * as PresetUtils from '../core/preset_utils';
import {
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
	SpellPowerBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common';
import { Mage_Options as MageOptions, Mage_Options_ArmorType as ArmorType } from '../core/proto/mage';
import { SavedTalents } from '../core/proto/ui';
import ArcaneAPL from './apls/forever_arcane.apl.json';
import FireAPL from './apls/forever_fire.apl.json';
import ArcaneFrostAPL from './apls/forever_arcane_frost.apl.json';
import FrostAPL from './apls/forever_frost.apl.json';
import GearArcaneJSON from './gear_sets/forever_arcane.gear.json';
import GearFireJSON from './gear_sets/forever_fire.gear.json';
import GearFrostJSON from './gear_sets/forever_frost.gear.json';

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearArcane = PresetUtils.makePresetGear('Arcane', GearArcaneJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearFire = PresetUtils.makePresetGear('Fire', GearFireJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearFrost = PresetUtils.makePresetGear('Frost', GearFrostJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearArcane, GearFire, GearFrost];
export const DefaultGear = GearFrost;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLFrost = PresetUtils.makePresetAPLRotation('Frost', FrostAPL);
export const APLArcane = PresetUtils.makePresetAPLRotation('Arcane', ArcaneAPL);
export const APLFire = PresetUtils.makePresetAPLRotation('Fire', FireAPL);
export const APLArcaneFrost = PresetUtils.makePresetAPLRotation('Arcane–Frost', ArcaneFrostAPL);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLFrost, APLArcane, APLFire, APLArcaneFrost],
};

export const DefaultAPL = APLFrost;

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1Frost = PresetUtils.makePresetTalents('Frost DPS', SavedTalents.create({ talentsString: '0502050030003--055500033100030024' }));
export const TalentsP1Arcane = PresetUtils.makePresetTalents('Arcane DPS', SavedTalents.create({ talentsString: '050215003100311531-2305003202003-' }));
export const TalentsP1Fire = PresetUtils.makePresetTalents('Fire 17/31/3', SavedTalents.create({ talentsString: '0501252000002-23450000130133051-003' }));

export const TalentsFire = PresetUtils.makePresetTalents('Fire 0/35/16', SavedTalents.create({ talentsString: '-03552020130133151-005500033' }));
export const TalentsFrost = PresetUtils.makePresetTalents('Frost 11/3/37', SavedTalents.create({ talentsString: '050005001-03-0555003321001301251' }));
export const TalentsArcane = PresetUtils.makePresetTalents('Arcane 33/3/15', SavedTalents.create({ talentsString: '053005023100311531-03-005500032' }));
export const TalentsArcaneFrost = PresetUtils.makePresetTalents('Arcane–Frost 28/0/23', SavedTalents.create({ talentsString: '05020500310031053--05550002010003002' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsP1Frost, TalentsP1Arcane, TalentsP1Fire, TalentsFire, TalentsFrost, TalentsArcane, TalentsArcaneFrost],
};

export const DefaultTalents = TalentsFrost;

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = MageOptions.create({
	armor: ArmorType.MageArmor,
});

export const DefaultConsumes = Consumes.create({
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	firePowerBuff: FirePowerBuff.ElixirOfGreaterFirepower,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodRunnTumTuberSurprise,
	frostPowerBuff: FrostPowerBuff.ElixirOfFrostPower,
	mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,

	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	zanzaBuff: ZanzaBuff.CerebralCortexCompound,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	moonkinAura: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
});

// Improved Scorch and Winter's Chill only help the mage that applied them in Forever, so they
// are no longer raid debuffs anyone else supplies.
export const DefaultDebuffs = Debuffs.create({
	judgementOfWisdom: true,
});

export const OtherDefaults = {
	distanceFromTarget: 20,
	profession1: Profession.Engineering,
	profession2: Profession.Tailoring,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Arcane', { gear: GearArcane, talents: TalentsArcane, rotation: APLArcane, options: DefaultOptions, distance: 20 }),
	PresetUtils.makePresetBuild('Fire', { gear: GearFire, talents: TalentsP1Fire, rotation: APLFire, options: DefaultOptions, distance: 20 }),
	PresetUtils.makePresetBuild('Frost', { gear: GearFrost, talents: TalentsFrost, rotation: APLFrost, options: DefaultOptions, distance: 20 }),
];
