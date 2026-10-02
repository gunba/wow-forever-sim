import { ClassicPhase } from '../core/constants/other.js';
import { foreverTankSupport } from '../core/forever_tank_support.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	ArmorElixir,
	AttackPowerBuff,
	Class,
	Consumes,
	Flask,
	Food,
	HealthElixir,
	Potions,
	StrengthBuff,
	UnitReference,
	ZanzaBuff,
} from '../core/proto/common.js';
import { FeralTankDruid_Options as DruidOptions } from '../core/proto/druid.js';
import { SavedTalents } from '../core/proto/ui.js';
import ForeverBearApl from './apls/forever_bear.apl.json';
import LaunchGearJSON from './gear_sets/launch.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearLaunch = PresetUtils.makePresetGear('Modeled level 65', LaunchGearJSON);

export const GearPresets = {
	[ClassicPhase.Phase1]: [GearLaunch],
};

export const DefaultGear = GearPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLForeverBear = PresetUtils.makePresetAPLRotation('Forever Bear', ForeverBearApl);

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLForeverBear],
};

export const DefaultAPL = APLPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsBearTank = PresetUtils.makePresetTalents('Bear Tank 1/40/10', SavedTalents.create({ talentsString: '01-50022332120132012551-055' }));

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsBearTank],
};

export const DefaultTalents = TalentPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = DruidOptions.create({
	innervateTarget: UnitReference.create(),
	startingRage: 0,
});

// No weapon coating for Bear. The separate party Windfury buff is enabled.
export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	armorElixir: ArmorElixir.ElixirOfSuperiorDefense,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultPotion: Potions.GreaterStoneshieldPotion,
	dragonBreathChili: true,
	food: Food.FoodSmokedDesertDumpling,
	flask: Flask.FlaskOfTheTitans,
	healthElixir: HealthElixir.ElixirOfFortitude,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

const Support = foreverTankSupport(Class.ClassDruid);
export const DefaultRaidBuffs = Support.raid;
export const DefaultPartyBuffs = Support.party;
export const DefaultIndividualBuffs = Support.player;
export const DefaultDebuffs = Support.debuffs;
