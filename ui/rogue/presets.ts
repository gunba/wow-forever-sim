import { ClassicPhase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	Flask,
	Food,
	IndividualBuffs,
	Profession,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
	SapperExplosive,
} from '../core/proto/common.js';
import { RogueOptions } from '../core/proto/rogue.js';
import { SavedTalents } from '../core/proto/ui.js';
import BackstabAPL from './apls/combat_backstab.apl.json';
import BackstabSweatyAPL from './apls/combat_backstab_sweaty.apl.json';
import SinisterStrikeAPL from './apls/combat_sinister_strike.apl.json';
import SinisterStrikeSweatyAPL from './apls/combat_sinister_strike_sweaty.apl.json';
import SinisterStrikeIEAAPL from './apls/combat_sinister_strike_iea.apl.json';
import HemorrhageAPL from './apls/forever_hemorrhage.apl.json';
import MutilateAPL from './apls/forever_mutilate.apl.json';
import GearCombatJSON from './gear_sets/forever_combat.gear.json';
import GearMutilateJSON from './gear_sets/forever_mutilate.gear.json';
import GearSubtletyJSON from './gear_sets/forever_subtlety.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearCombat = PresetUtils.makePresetGear('Combat', GearCombatJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearMutilate = PresetUtils.makePresetGear('Mutilate', GearMutilateJSON, { tooltip: 'Level 60 Forever equipment.' });
export const GearSubtlety = PresetUtils.makePresetGear('Subtlety', GearSubtletyJSON, { tooltip: 'Level 60 Forever equipment.' });

export const GearPresets = [GearCombat, GearMutilate, GearSubtlety];
export const DefaultGear = GearCombat;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets[]
///////////////////////////////////////////////////////////////////////////

export const ROTATION_PRESET_BACKSTAB = PresetUtils.makePresetAPLRotation('Backstab', BackstabAPL, {});
export const ROTATION_PRESET_SINISTER_STRIKE = PresetUtils.makePresetAPLRotation('Sinister Strike', SinisterStrikeAPL, {});
export const ROTATION_PRESET_BACKSTAB_SWEATY = PresetUtils.makePresetAPLRotation('Backstab (Sweaty)', BackstabSweatyAPL, {});
export const ROTATION_PRESET_SINISTER_STRIKE_SWEATY = PresetUtils.makePresetAPLRotation('Sinister Strike (Sweaty)', SinisterStrikeSweatyAPL, {});
export const ROTATION_PRESET_SINISTER_STRIKE_IEA = PresetUtils.makePresetAPLRotation('Improved Expose Armor (SS)', SinisterStrikeIEAAPL, {});
export const ROTATION_PRESET_MUTILATE = PresetUtils.makePresetAPLRotation('Mutilate', MutilateAPL, {});
export const ROTATION_PRESET_HEMORRHAGE = PresetUtils.makePresetAPLRotation('Hemorrhage', HemorrhageAPL, {});

export const APLPresets = {
	[ClassicPhase.Phase1]: [
		ROTATION_PRESET_BACKSTAB,
		ROTATION_PRESET_SINISTER_STRIKE,
		ROTATION_PRESET_BACKSTAB_SWEATY,
		ROTATION_PRESET_SINISTER_STRIKE_SWEATY,
		ROTATION_PRESET_SINISTER_STRIKE_IEA,
		ROTATION_PRESET_MUTILATE,
		ROTATION_PRESET_HEMORRHAGE,
	],
	[ClassicPhase.Phase2]: [
		ROTATION_PRESET_BACKSTAB,
		ROTATION_PRESET_SINISTER_STRIKE,
		ROTATION_PRESET_BACKSTAB_SWEATY,
		ROTATION_PRESET_SINISTER_STRIKE_SWEATY,
		ROTATION_PRESET_SINISTER_STRIKE_IEA,
		ROTATION_PRESET_MUTILATE,
		ROTATION_PRESET_HEMORRHAGE,
	],
};

//Need to add main hand equip logic or talent/rotation logic to map to Auto APL
export const DefaultAPLs: Record<number, PresetUtils.PresetRotation> = {
	[0]: ROTATION_PRESET_SINISTER_STRIKE,
	[1]: ROTATION_PRESET_BACKSTAB,
};

export const DefaultAPLBackstab = APLPresets[ClassicPhase.Phase2][0];
export const DefaultAPLSinisterStrike = APLPresets[ClassicPhase.Phase2][1];
export const DefaultAPLIEA = APLPresets[ClassicPhase.Phase2][4];
export const DefaultAPLMutilate = APLPresets[ClassicPhase.Phase2][5];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

// Preset name must be unique. Ex: 'Backstab DPS' cannot be used as a name more than once

export const CombatBackstabTalents = PresetUtils.makePresetTalents('Backstab', SavedTalents.create({ talentsString: '005302005-30230320201515231-102' }));
export const CombatSinisterStrikeTalents = PresetUtils.makePresetTalents(
	'Sinister Strike',
	SavedTalents.create({ talentsString: '005303103012-32003311201515231' }),
);
export const CombatSinisterStrikeIEATalents = PresetUtils.makePresetTalents(
	'Improved Expose Armor (SS)',
	SavedTalents.create({ talentsString: '005303125-32003311201515131' }),
);
export const AssassinationMutilateTalents = PresetUtils.makePresetTalents('Mutilate', SavedTalents.create({ talentsString: '0053031035140105-302303202014' }));

export const TalentsCombatDualWield = PresetUtils.makePresetTalents(
	'Combat Dual-Wield 15/33/3',
	SavedTalents.create({ talentsString: '1053231-22530300001515231-012' }),
);
export const TalentsAssassinationMutilate = PresetUtils.makePresetTalents(
	'Assassination Mutilate 31/20/0',
	SavedTalents.create({ talentsString: '02532010531201051-225303000005' }),
);
export const TalentsSubtletyHemo = PresetUtils.makePresetTalents(
	'Subtlety Hemo 20/0/31',
	SavedTalents.create({ talentsString: '115303101104--5320003310013211501' }),
);

export const TalentPresets = {
	[ClassicPhase.Phase1]: [
		CombatBackstabTalents,
		CombatSinisterStrikeTalents,
		CombatSinisterStrikeIEATalents,
		AssassinationMutilateTalents,
		TalentsCombatDualWield,
		TalentsAssassinationMutilate,
		TalentsSubtletyHemo,
	],
	[ClassicPhase.Phase2]: [
		CombatBackstabTalents,
		CombatSinisterStrikeTalents,
		CombatSinisterStrikeIEATalents,
		AssassinationMutilateTalents,
		TalentsCombatDualWield,
		TalentsAssassinationMutilate,
		TalentsSubtletyHemo,
	],
};

export const DefaultTalentsAssassin = AssassinationMutilateTalents;
export const DefaultTalentsCombat = CombatSinisterStrikeTalents;
export const DefaultTalentsSubtlety = TalentsSubtletyHemo;

export const DefaultTalentsBackstab = TalentPresets[ClassicPhase.Phase2][0];
export const DefaultTalentsSinisterStrike = TalentPresets[ClassicPhase.Phase2][1];
export const DefaultTalentsIEA = TalentPresets[ClassicPhase.Phase2][2];
export const DefaultTalentsMutilate = TalentPresets[ClassicPhase.Phase2][3];

export const DefaultTalents = CombatSinisterStrikeTalents;

///////////////////////////////////////////////////////////////////////////
//                                Build Presets
///////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = RogueOptions.create({});

///////////////////////////////////////////////////////////////////////////
//                         Consumes/Buffs/Debuffs
///////////////////////////////////////////////////////////////////////////

export const P1Consumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredRogueThistleTea,
	dragonBreathChili: true,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodGrilledSquid,
	mainHandImbue: WeaponImbue.InstantPoison,
	offHandImbue: WeaponImbue.DeadlyPoison,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.GroundScorpokAssay,
	sapperExplosive: SapperExplosive.SapperGoblinSapper,
});

export const DefaultConsumes = {
	[ClassicPhase.Phase1]: P1Consumes,
};

export const P1RaidBuffs = RaidBuffs.create({
	battleShout: TristateEffect.TristateEffectImproved,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
	trueshotAura: true,
});

export const DefaultRaidBuffs = {
	[ClassicPhase.Phase1]: P1RaidBuffs,
};

export const P1IndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	//saygesFortune: SaygesFortune.SaygesDamage,
	spiritOfZandalar: false,
});

export const DefaultIndividualBuffs = {
	[ClassicPhase.Phase1]: P1IndividualBuffs,
};

export const P1DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	faerieFire: true,
	sunderArmor: true,
});

export const DefaultDebuffs = {
	[ClassicPhase.Phase1]: P1DefaultDebuffs,
};

export const P1OtherDefaults = {
	profession1: Profession.Engineering,
	profession2: Profession.ProfessionUnknown,
};

export const OtherDefaults = {
	[ClassicPhase.Phase1]: P1OtherDefaults,
};

export const BuildPresets = [
	PresetUtils.makePresetBuild('Combat', {
		gear: GearCombat,
		talents: CombatSinisterStrikeTalents,
		rotation: ROTATION_PRESET_SINISTER_STRIKE_SWEATY,
		options: DefaultOptions,
		distance: 5,
	}),
	PresetUtils.makePresetBuild('Mutilate', {
		gear: GearMutilate,
		talents: AssassinationMutilateTalents,
		rotation: ROTATION_PRESET_MUTILATE,
		options: DefaultOptions,
		distance: 5,
	}),
	PresetUtils.makePresetBuild('Subtlety', {
		gear: GearSubtlety,
		talents: TalentsSubtletyHemo,
		rotation: ROTATION_PRESET_HEMORRHAGE,
		options: DefaultOptions,
		distance: 5,
	}),
];
