import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	FirePowerBuff,
	Flask,
	Food,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	Race,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { EnhancementShaman_Options as EnhancementShamanOptions } from '../core/proto/shaman.js';
import { SavedTalents } from '../core/proto/ui.js';
import DefaultAPLJSON from './apls/default.apl.json';
import GearEnhancementJSON from './gear_sets/forever_enhancement.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearEnhancement = PresetUtils.makePresetGear('Enhancement', GearEnhancementJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearEnhancement];
export const DefaultGear = GearEnhancement;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLDefault = PresetUtils.makePresetAPLRotation('Default', DefaultAPLJSON);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLDefault],
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

export const TalentsLevel60 = PresetUtils.makePresetTalents('Level 60', SavedTalents.create({ talentsString: '5505301-053030031005112251' }));

export const TalentsEnhancement = PresetUtils.makePresetTalents('Enhancement 20/31/0', SavedTalents.create({ talentsString: '550133102-053030031005102251' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsLevel60, TalentsEnhancement],
	[ClassicPhase.Phase2]: [],
	[ClassicPhase.Phase3]: [],
	[ClassicPhase.Phase4]: [],
	[ClassicPhase.Phase5]: [],
	[ClassicPhase.Phase6]: [],
};

export const DefaultTalents = TalentsEnhancement;

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = EnhancementShamanOptions.create();

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultPotion: Potions.MajorManaPotion,
	defaultConjured: Conjured.ConjuredDemonicRune,
	dragonBreathChili: true,
	firePowerBuff: FirePowerBuff.ElixirOfFirepower,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodBlessSunfruit,
	mainHandImbue: WeaponImbue.WindfuryWeapon,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	// saygesFortune: SaygesFortune.SaygesDamage,
	// spiritOfZandalar: true,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	sunderArmor: true,
});

export const OtherDefaults = {
	profession1: Profession.Engineering,
	profession2: Profession.Enchanting,
	race: Race.RaceOrc,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Enhancement', {
		gear: GearEnhancement,
		talents: TalentsEnhancement,
		rotation: APLDefault,
		options: DefaultOptions,
		distance: 5,
	}),
];
