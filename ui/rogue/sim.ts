import * as BuffDebuffInputs from '../core/components/inputs/buffs_debuffs';
import * as OtherInputs from '../core/components/other_inputs.js';
import { ClassicPhase } from '../core/constants/other.js';
import { IndividualSimUI, registerSpecConfig } from '../core/individual_sim_ui.js';
import { Player } from '../core/player.js';
import { Class, Faction, ItemSlot, PartyBuffs, PseudoStat, Race, Spec, Stat, Target, WeaponType } from '../core/proto/common.js';
import { Stats } from '../core/proto_utils/stats.js';
import { getSpecIcon } from '../core/proto_utils/utils.js';
import * as Presets from './presets.js';

const SPEC_CONFIG = registerSpecConfig(Spec.SpecRogue, {
	cssClass: 'rogue-sim-ui',
	cssScheme: 'rogue',
	// List any known bugs / issues here and they'll be shown on the site.
	knownIssues: ['Rotations are not fully optimized, especially for non-standard setups.'],
	warnings: [
		(simUI: IndividualSimUI<Spec.SpecRogue>) => {
			return {
				updateOn: simUI.sim.encounter.changeEmitter,
				getContent: () => {
					const hasNoArmor = !!(simUI.sim.encounter.targets ?? []).find((target: Target) => new Stats(target.stats).getStat(Stat.StatArmor) <= 0);
					if (hasNoArmor) {
						return 'One or more targets have no armor. Check advanced encounter settings.';
					} else {
						return '';
					}
				},
			};
		},
	],

	// All stats for which EP should be calculated.
	epStats: [
		// Attributes
		Stat.StatAgility,
		Stat.StatStrength,
		// Physical
		Stat.StatAttackPower,
		Stat.StatMeleeHit,
		Stat.StatMeleeCrit,
		// Spell
		Stat.StatSpellPower,
	],
	epPseudoStats: [PseudoStat.PseudoStatMainHandDps, PseudoStat.PseudoStatOffHandDps, PseudoStat.PseudoStatMeleeSpeedMultiplier],
	// Reference stat against which to calculate EP.
	epReferenceStat: Stat.StatAttackPower,
	// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
	displayStats: [
		// Attributes
		Stat.StatAgility,
		Stat.StatStrength,
		// Physical
		Stat.StatAttackPower,
		Stat.StatMeleeHit,
		Stat.StatMeleeCrit,
	],
	displayPseudoStats: [PseudoStat.PseudoStatMeleeSpeedMultiplier],

	defaults: {
		race: Race.RaceOrc,
		// Default equipped gear.
		gear: Presets.DefaultGear.gear,
		// Default EP weights for sorting gear in the gear picker.
		epWeights: Stats.fromMap(
			{
				[Stat.StatAgility]: 2.38,
				[Stat.StatStrength]: 1.26,
				[Stat.StatAttackPower]: 1.0,
				[Stat.StatSpellCrit]: 0.41,
				[Stat.StatSpellHit]: 0.94,
				[Stat.StatMeleeHit]: 29.44,
				[Stat.StatMeleeCrit]: 17.92,
				[Stat.StatFireResistance]: 0.5,
			},
			{
				[PseudoStat.PseudoStatMainHandDps]: 10.49,
				[PseudoStat.PseudoStatOffHandDps]: 3.74,
				[PseudoStat.PseudoStatMeleeSpeedMultiplier]: 18.56,
			},
		),

		// Default consumes settings.
		consumes: Presets.DefaultConsumes[ClassicPhase.Phase1],
		// Default talents.
		talents: Presets.DefaultTalents.data,
		// Default spec-specific settings.
		specOptions: Presets.DefaultOptions,
		other: Presets.OtherDefaults[ClassicPhase.Phase1],
		// Default raid/party buffs settings.
		raidBuffs: Presets.DefaultRaidBuffs[ClassicPhase.Phase1],
		partyBuffs: PartyBuffs.create({}),
		individualBuffs: Presets.DefaultIndividualBuffs[ClassicPhase.Phase1],
		debuffs: Presets.DefaultDebuffs[ClassicPhase.Phase1],
	},

	playerInputs: {
		inputs: [],
	},
	// IconInputs to include in the 'Player' section on the settings tab.
	playerIconInputs: [],
	// Buff and Debuff inputs to include/exclude, overriding the EP-based defaults.
	includeBuffDebuffInputs: [
		BuffDebuffInputs.SpellCritBuff,
		BuffDebuffInputs.SpellShadowWeavingDebuff,
		BuffDebuffInputs.SpellScorchDebuff,
		BuffDebuffInputs.PowerInfusion,
	],
	excludeBuffDebuffInputs: [],
	// Inputs to include in the 'Other' section on the settings tab.
	otherInputs: {
		inputs: [OtherInputs.TankAssignment, OtherInputs.InFrontOfTarget],
	},
	encounterPicker: {
		// Whether to include 'Execute Duration (%)' in the 'Encounter' section of the settings tab.
		showExecuteProportion: false,
	},

	presets: {
		talents: Presets.BuildPresets.map(build => build.talents!),
		rotations: Presets.BuildPresets.map(build => build.rotation!),
		gear: Presets.GearPresets,
		builds: Presets.BuildPresets,
	},

	autoRotation: player => {
		const talents = player.getTalents();
		if (talents.hemorrhage) {
			return Presets.ROTATION_PRESET_HEMORRHAGE.rotation.rotation!;
		}
		if (player.getEquippedItem(ItemSlot.ItemSlotMainHand)?._item.weaponType == WeaponType.WeaponTypeDagger) {
			if (talents.mutilate) {
				return Presets.ROTATION_PRESET_MUTILATE.rotation.rotation!;
			}
			return Presets.ROTATION_PRESET_BACKSTAB.rotation.rotation!;
		}
		return Presets.ROTATION_PRESET_SINISTER_STRIKE_SWEATY.rotation.rotation!;
	},

	raidSimPresets: [
		{
			spec: Spec.SpecRogue,
			tooltip: 'Assassination Rogue',
			defaultName: 'Assassination',
			iconUrl: getSpecIcon(Class.ClassRogue, 0),

			talents: Presets.DefaultTalentsAssassin.data,
			specOptions: Presets.DefaultOptions,
			consumes: Presets.DefaultConsumes[ClassicPhase.Phase1],
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceHuman,
				[Faction.Horde]: Race.RaceOrc,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: Presets.GearMutilate.gear,
				},
				[Faction.Horde]: {
					1: Presets.GearMutilate.gear,
				},
			},
		},
		{
			spec: Spec.SpecRogue,
			tooltip: 'Combat Rogue',
			defaultName: 'Combat',
			iconUrl: getSpecIcon(Class.ClassRogue, 1),

			talents: Presets.DefaultTalentsCombat.data,
			specOptions: Presets.DefaultOptions,
			consumes: Presets.DefaultConsumes[ClassicPhase.Phase1],
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceHuman,
				[Faction.Horde]: Race.RaceOrc,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: Presets.GearCombat.gear,
				},
				[Faction.Horde]: {
					1: Presets.GearCombat.gear,
				},
			},
		},
	],
});

export class RogueSimUI extends IndividualSimUI<Spec.SpecRogue> {
	constructor(parentElem: HTMLElement, player: Player<Spec.SpecRogue>) {
		super(parentElem, player, SPEC_CONFIG);
	}
}
