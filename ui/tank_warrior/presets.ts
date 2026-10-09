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
	Profession,
	Race,
	StrengthBuff,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';
import { TankWarrior_Options as TankWarriorOptions, WarriorShout, WarriorStance } from '../core/proto/warrior.js';
import APLForeverProtectionJSON from './apls/forever_protection.apl.json';
import APLQueueThirtyJSON from './apls/forever_protection_tauren.apl.json';
import APLQueueFortyFiveJSON from './apls/forever_protection_orc.apl.json';
import LaunchGearJSON from './gear_sets/launch.gear.json';

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

export const APLForeverProtection = PresetUtils.makePresetAPLRotation('Protection · queue 20', APLForeverProtectionJSON, { customCondition: player => [Race.RaceHuman, Race.RaceUndead].includes(player.getRace()) });
export const APLQueueThirty = PresetUtils.makePresetAPLRotation('Protection · queue 30', APLQueueThirtyJSON, { customCondition: player => ![Race.RaceHuman, Race.RaceUndead, Race.RaceOrc, Race.RaceDwarf].includes(player.getRace()) });
export const APLQueueFortyFive = PresetUtils.makePresetAPLRotation('Protection · queue 45', APLQueueFortyFiveJSON, { customCondition: player => [Race.RaceOrc, Race.RaceDwarf].includes(player.getRace()) });

export const APLPresets = {
	[ClassicPhase.Phase1]: [APLForeverProtection, APLQueueThirty, APLQueueFortyFive],
};

export const DefaultAPL = APLPresets[ClassicPhase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsProtection = PresetUtils.makePresetTalents('Protection 8/6/37', SavedTalents.create({ talentsString: '35-0501-251333120230021351' }));

// Race-specific selected recipes; full ranked profiles also retain exact gear and APL.
export const TalentsProtectionOrc = PresetUtils.makePresetTalents('Protection · Orc', SavedTalents.create({ talentsString: '31-0501-255233120330021351' }), { customCondition: player => player.getRace() === Race.RaceOrc });
export const TalentsProtectionTauren = PresetUtils.makePresetTalents('Protection · Tauren', SavedTalents.create({ talentsString: '31-0501-255333120330011351' }), { customCondition: player => player.getRace() === Race.RaceTauren });
export const TalentsProtectionTroll = PresetUtils.makePresetTalents('Protection · Troll', SavedTalents.create({ talentsString: '31-0501-255333120330011351' }), { customCondition: player => player.getRace() === Race.RaceTroll });
export const TalentsProtectionUndead = PresetUtils.makePresetTalents('Protection · Undead', SavedTalents.create({ talentsString: '35-0501-251333120230021351' }), { customCondition: player => player.getRace() === Race.RaceUndead });
export const TalentsProtectionWindshaper = PresetUtils.makePresetTalents('Protection · Windshaper', SavedTalents.create({ talentsString: '31-0501-255333120230021351' }), { customCondition: player => player.getRace() === Race.RaceSkyborneWindshaper });
export const TalentsProtectionHuman = PresetUtils.makePresetTalents('Protection · Human', SavedTalents.create({ talentsString: '35-0501-251333120230021351' }), { customCondition: player => player.getRace() === Race.RaceHuman });
export const TalentsProtectionDwarf = PresetUtils.makePresetTalents('Protection · Dwarf', SavedTalents.create({ talentsString: '31-0501-255333120330011351' }), { customCondition: player => player.getRace() === Race.RaceDwarf });
export const TalentsProtectionNightElf = PresetUtils.makePresetTalents('Protection · Night Elf', SavedTalents.create({ talentsString: '31-0501-255333120230021351' }), { customCondition: player => player.getRace() === Race.RaceNightElf });
export const TalentsProtectionGnome = PresetUtils.makePresetTalents('Protection · Gnome', SavedTalents.create({ talentsString: '31-0501-255333120230021351' }), { customCondition: player => player.getRace() === Race.RaceGnome });
export const TalentsProtectionHighOrder = PresetUtils.makePresetTalents('Protection · High Order', SavedTalents.create({ talentsString: '31-0501-255333120230021351' }), { customCondition: player => player.getRace() === Race.RaceSkyborneHighOrder });

export const TalentPresets = {
	[ClassicPhase.Phase1]: [TalentsProtectionOrc, TalentsProtectionTauren, TalentsProtectionTroll, TalentsProtectionUndead, TalentsProtectionWindshaper, TalentsProtectionHuman, TalentsProtectionDwarf, TalentsProtectionNightElf, TalentsProtectionGnome, TalentsProtectionHighOrder],
};

export const DefaultTalents = TalentPresets[ClassicPhase.Phase1][0];

export const PresetBuildTanky = PresetUtils.makePresetBuild('Tanky', { gear: DefaultGear, talents: TalentsProtection, rotation: DefaultAPL });

///////////////////////////////////////////////////////////////////////////
//                                 Options Presets
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = TankWarriorOptions.create({
	queueDelay: 250,
	startingRage: 0,
	shout: WarriorShout.WarriorShoutBattle,
	stance: WarriorStance.WarriorStanceDefensive,
});

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
	mainHandImbue: WeaponImbue.ElementalSharpeningStone,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

const Support = foreverTankSupport(Class.ClassWarrior);
export const DefaultRaidBuffs = Support.raid;
export const DefaultPartyBuffs = Support.party;
export const DefaultIndividualBuffs = Support.player;
export const DefaultDebuffs = Support.debuffs;

export const OtherDefaults = {
	profession1: Profession.Blacksmithing,
	profession2: Profession.Enchanting,
	race: Race.RaceHuman,
};
