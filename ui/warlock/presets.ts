import { Player } from '../core/player.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	Alcohol,
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
	RaidBuffs,
	SaygesFortune,
	ShadowPowerBuff,
	SpellPowerBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common';
import { SavedTalents } from '../core/proto/ui.js';
import {
	WarlockOptions as WarlockOptions,
	WarlockOptions_Armor as Armor,
	WarlockOptions_Summon as Summon,
	WarlockOptions_WeaponImbue as WarlockWeaponImbue,
} from '../core/proto/warlock.js';
// apls
import AfflictionApl from './apls/forever_affliction.apl.json';
import DSRuinApl from './apls/forever_ds_ruin.apl.json';
import DemonicPactApl from './apls/forever_pact.apl.json';
import ShadowAndFlameApl from './apls/forever_shadow_and_flame.apl.json';
// gear
import GearDemonologyJSON from './gear_sets/forever_demonology.gear.json';
import GearAfflictionJSON from './gear_sets/forever_affliction.gear.json';
import GearDsRuinJSON from './gear_sets/forever_ds_ruin.gear.json';
import GearDestructionJSON from './gear_sets/forever_destruction.gear.json';

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearDemonology = PresetUtils.makePresetGear('Demonic Pact', GearDemonologyJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearAffliction = PresetUtils.makePresetGear('Affliction', GearAfflictionJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearDsRuin = PresetUtils.makePresetGear('DS/Ruin', GearDsRuinJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearDestruction = PresetUtils.makePresetGear('Destruction', GearDestructionJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearDemonology, GearAffliction, GearDsRuin, GearDestruction];
export const DefaultGear = GearDsRuin;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

// P1
export const RotationDemonicPact = PresetUtils.makePresetAPLRotation('Demonic Pact', DemonicPactApl);
export const RotationAffliction = PresetUtils.makePresetAPLRotation('Affliction', AfflictionApl);
export const RotationDSRuin = PresetUtils.makePresetAPLRotation('DS/Ruin', DSRuinApl);
export const RotationShadowAndFlame = PresetUtils.makePresetAPLRotation('Shadow and Flame', ShadowAndFlameApl);

export const APLPresets = [RotationDemonicPact, RotationAffliction, RotationDSRuin, RotationShadowAndFlame];

export const DefaultAPL = RotationDSRuin;

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsPactOptimised = PresetUtils.makePresetTalents(
	'Demonic Pact 2/31/18',
	SavedTalents.create({ talentsString: '011-0005003221220311351-0550005003' }),
);
export const TalentsDeepAffliction = PresetUtils.makePresetTalents(
	'Deep Affliction 31/0/20',
	SavedTalents.create({ talentsString: '2525000013520105--0550015103' }),
);
export const TalentsDSRuinPandemic = PresetUtils.makePresetTalents(
	'DS/Ruin Pandemic 21/11/19',
	SavedTalents.create({ talentsString: '252200001351-0025003001-0550005103' }),
);
export const TalentsShadowAndFlame = PresetUtils.makePresetTalents(
	'Shadow and Flame 13/5/33',
	SavedTalents.create({ talentsString: '2521000003-0005-0050355103101351' }),
);

export const TalentPresets = [
	TalentsPactOptimised,
	TalentsDeepAffliction,
	TalentsDSRuinPandemic,
	TalentsShadowAndFlame,
];

export const DefaultTalents = TalentsDSRuinPandemic;

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = WarlockOptions.create({
	armor: Armor.DemonArmor,
	summon: Summon.Imp,
	weaponImbue: WarlockWeaponImbue.NoWeaponImbue,
});

// Without pet talents the Succubus out-damages the Imp, so the Affliction builds run one; a
// sacrificing rotation summons and sacrifices its own Imp, which leaves Shadow damage in Forever.
export const AfflictionOptions = WarlockOptions.create({
	armor: Armor.DemonArmor,
	summon: Summon.Succubus,
	weaponImbue: WarlockWeaponImbue.NoWeaponImbue,
});

// Demonic Pact keeps the sacrifice when another demon is out. The Succubus stays out for Master
// Demonologist and Soul Link, and the Voidwalker is sacrificed for the mana, which Forever moved
// from the Felhunter to the Voidwalker.
export const DemonicPactOptions = WarlockOptions.create({
	armor: Armor.DemonArmor,
	summon: Summon.Succubus,
	sacrifice: Summon.Voidwalker,
	weaponImbue: WarlockWeaponImbue.NoWeaponImbue,
});

export const DefaultConsumes = Consumes.create({
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	defaultPotion: Potions.MajorManaPotion,
	defaultConjured: Conjured.ConjuredDemonicRune,
	flask: Flask.FlaskOfSupremePower,
	firePowerBuff: FirePowerBuff.ElixirOfFirepower,
	food: Food.FoodRunnTumTuberSurprise,
	// mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	shadowPowerBuff: ShadowPowerBuff.ElixirOfShadowPower,
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
	powerWordFortitude: TristateEffect.TristateEffectImproved,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	// saygesFortune: SaygesFortune.SaygesDamage,
	// spiritOfZandalar: true,
});

export const DefaultDebuffs = Debuffs.create({
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	judgementOfWisdom: true,
	sunderArmor: true,
});

///////////////////////////////////////////////////////////////////////////
//                                 Builds
///////////////////////////////////////////////////////////////////////////

// The community builds with the pet setup and rotation each one is measured with.

export const OtherDefaults = {
	distanceFromTarget: 20,
	profession1: Profession.Engineering,
	profession2: Profession.Tailoring,
	channelClipDelay: 150,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Demonic Pact', {
		gear: GearDemonology,
		talents: TalentsPactOptimised,
		rotation: RotationDemonicPact,
		options: DemonicPactOptions,
		distance: 20,
	}),
	PresetUtils.makePresetBuild('Affliction', {
		gear: GearAffliction,
		talents: TalentsDeepAffliction,
		rotation: RotationAffliction,
		options: AfflictionOptions,
		distance: 20,
	}),
	PresetUtils.makePresetBuild('DS/Ruin', {
		gear: GearDsRuin,
		talents: TalentsDSRuinPandemic,
		rotation: RotationDSRuin,
		options: DefaultOptions,
		distance: 20,
	}),
	PresetUtils.makePresetBuild('Destruction', {
		gear: GearDestruction,
		talents: TalentsShadowAndFlame,
		rotation: RotationShadowAndFlame,
		options: AfflictionOptions,
		distance: 20,
	}),
];
