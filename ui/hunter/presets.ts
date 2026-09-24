import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	Flask,
	Food,
	HealthElixir,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	Race,
	RaidBuffs,
	SapperExplosive,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import {
	Hunter_Options as HunterOptions,
	Hunter_Options_Ammo as Ammo,
	Hunter_Options_PetAttackSpeed as PetAttackSpeed,
	Hunter_Options_PetType as PetType,
	Hunter_Options_QuiverBonus,
} from '../core/proto/hunter.js';
import { SavedTalents } from '../core/proto/ui.js';
import P1APL from './apls/p1.apl.json';
import BeastMasteryAPL from './apls/beast_mastery.apl.json';
import SurvivalAPL from './apls/survival.apl.json';
import PetMeleeAPL from './apls/pet_melee.apl.json';
import GearBeastMasteryJSON from './gear_sets/forever_beast_mastery.gear.json';
import GearMarksmanshipJSON from './gear_sets/forever_marksmanship.gear.json';
import GearSurvivalJSON from './gear_sets/forever_survival.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearBeastMastery = PresetUtils.makePresetGear('Beast Mastery', GearBeastMasteryJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearMarksmanship = PresetUtils.makePresetGear('Marksmanship', GearMarksmanshipJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearSurvival = PresetUtils.makePresetGear('Survival', GearSurvivalJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearBeastMastery, GearMarksmanship, GearSurvival];
export const DefaultGear = GearMarksmanship;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLP1 = PresetUtils.makePresetAPLRotation('Marksmanship', P1APL);
export const APLBeastMastery = PresetUtils.makePresetAPLRotation('Beast Mastery', BeastMasteryAPL);
export const APLSurvival = PresetUtils.makePresetAPLRotation('Survival', SurvivalAPL);
export const APLPetMelee = PresetUtils.makePresetAPLRotation('Pet/Melee', PetMeleeAPL);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLP1, APLBeastMastery, APLSurvival, APLPetMelee],
};

export const DefaultAPL = APLP1;

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1 = PresetUtils.makePresetTalents('Marksmanship 5/35/11', SavedTalents.create({ talentsString: '5-0050552011523051-50005001' }));

export const TalentsBeastMastery = PresetUtils.makePresetTalents(
	'Beast Mastery 31/20/0',
	SavedTalents.create({ talentsString: '5320001505101251-00531510005' }),
);
export const TalentsSurvival = PresetUtils.makePresetTalents('Survival 7/11/33', SavedTalents.create({ talentsString: '502-0050051-230230230250022151' }));
export const TalentsPetMelee = PresetUtils.makePresetTalents('Pet/Melee 16/10/25', SavedTalents.create({ talentsString: '53200005001-005005-5302002300502201' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsP1, TalentsBeastMastery, TalentsSurvival, TalentsPetMelee],
};

export const DefaultTalents = TalentsP1;

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const PetOptions = HunterOptions.create({
	ammo: Ammo.ThoriumHeadedArrow,
	quiverBonus: Hunter_Options_QuiverBonus.Speed15,
	petAttackSpeed: PetAttackSpeed.OneTwo,
	petType: PetType.Cat,
	petUptime: 1,
});

export const DefaultOptions = HunterOptions.create({ ...PetOptions, petType: PetType.PetNone });

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodSmokedDesertDumpling,
	healthElixir: HealthElixir.ElixirOfFortitude,
	mainHandImbue: WeaponImbue.Windfury,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,
	offHandImbue: WeaponImbue.ElementalSharpeningStone,
	sapperExplosive: SapperExplosive.SapperUnknown,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.GroundScorpokAssay,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectImproved,
	leaderOfThePack: false,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	fengusFerocity: false,
	moldarsMoxie: false,
	slipkiksSavvy: false,
	spiritOfZandalar: false,
	warchiefsBlessing: false,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	huntersMark: TristateEffect.TristateEffectRegular,
	judgementOfWisdom: true,
	stormstrike: false,
	sunderArmor: true,
});

export const OtherDefaults = {
	distanceFromTarget: 12,
	profession1: Profession.Engineering,
	profession2: Profession.Engineering,
	race: Race.RaceTroll,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Beast Mastery', {
		gear: GearBeastMastery,
		talents: TalentsBeastMastery,
		rotation: APLBeastMastery,
		options: PetOptions,
		distance: 12,
	}),
	PresetUtils.makePresetBuild('Marksmanship', { gear: GearMarksmanship, talents: TalentsP1, rotation: APLP1, options: DefaultOptions, distance: 12 }),
	PresetUtils.makePresetBuild('Survival', { gear: GearSurvival, talents: TalentsSurvival, rotation: APLSurvival, options: PetOptions, distance: 5 }),
	PresetUtils.makePresetBuild('Pet/Melee', { gear: GearSurvival, talents: TalentsPetMelee, rotation: APLPetMelee, options: PetOptions, distance: 5 }),
];
