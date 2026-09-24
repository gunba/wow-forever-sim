import * as BuffDebuffInputs from '../core/components/inputs/buffs_debuffs';
import * as ConsumesInputs from '../core/components/inputs/consumables';
import * as OtherInputs from '../core/components/other_inputs';
import { ClassicPhase } from '../core/constants/other';
import { IndividualSimUI, registerSpecConfig } from '../core/individual_sim_ui';
import { Player } from '../core/player';
import { Class, Faction, ItemSlot, PartyBuffs, PseudoStat, Race, Spec, Stat } from '../core/proto/common';
import { Stats } from '../core/proto_utils/stats';
import { getSpecIcon, specNames } from '../core/proto_utils/utils';
import * as Presets from './presets';

const SPEC_CONFIG = registerSpecConfig(Spec.SpecElementalShaman, {
	cssClass: 'elemental-shaman-sim-ui',
	cssScheme: 'shaman',
	// List any known bugs / issues here and they'll be shown on the site.
	knownIssues: [],
	warnings: [],

	// All stats for which EP should be calculated.
	epStats: [
		// Primary
		Stat.StatMana,
		// Attributes
		Stat.StatIntellect,
		Stat.StatSpirit,
		// Spell
		Stat.StatSpellPower,
		Stat.StatSpellDamage,
		Stat.StatFirePower,
		Stat.StatNaturePower,
		Stat.StatSpellHit,
		Stat.StatSpellCrit,
		Stat.StatSpellHaste,
		Stat.StatMP5,
	],
	epPseudoStats: [],
	// Reference stat against which to calculate EP. I think all classes use either spell power or attack power.
	epReferenceStat: Stat.StatSpellPower,
	// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
	displayStats: [
		// Primary
		Stat.StatMana,
		// Attributes
		Stat.StatIntellect,
		Stat.StatSpirit,
		// Spell
		Stat.StatSpellDamage,
		Stat.StatNaturePower,
		Stat.StatFirePower,
		Stat.StatSpellHit,
		Stat.StatSpellCrit,
		Stat.StatSpellHaste,
		Stat.StatMP5,
	],
	displayPseudoStats: [PseudoStat.PseudoStatCastSpeedMultiplier],

	defaults: {
		race: Race.RaceOrc,
		// Default equipped gear.
		gear: Presets.DefaultGear.gear,
		// Default EP weights for sorting gear in the gear picker.
		epWeights: Stats.fromMap({
			[Stat.StatIntellect]: 0.14,
			[Stat.StatSpellPower]: 1,
			[Stat.StatSpellDamage]: 1,
			[Stat.StatFirePower]: 0.1, // Eles don't really use Fire much except for Niche situations or AOE
			[Stat.StatNaturePower]: 1.0,
			[Stat.StatSpellHit]: 12.37,
			[Stat.StatSpellCrit]: 7.57,
			[Stat.StatSpellHaste]: 1.49,
			[Stat.StatMP5]: 0.02,
			[Stat.StatStrength]: 0.01,
			[Stat.StatAttackPower]: 0.01,
			[Stat.StatFireResistance]: 0.01,
		}),
		// Default consumes settings.
		consumes: Presets.DefaultConsumes,
		// Default talents.
		talents: Presets.DefaultTalents.data,
		// Default spec-specific settings.
		specOptions: Presets.DefaultOptions,
		other: Presets.OtherDefaults,
		// Default raid/party buffs settings.
		raidBuffs: Presets.DefaultRaidBuffs,
		partyBuffs: PartyBuffs.create({}),
		individualBuffs: Presets.DefaultIndividualBuffs,
		debuffs: Presets.DefaultDebuffs,
	},
	// IconInputs to include in the 'Player' section on the settings tab.
	playerIconInputs: [],
	// Buff and Debuff inputs to include/exclude, overriding the EP-based defaults.
	includeBuffDebuffInputs: [BuffDebuffInputs.SpellWintersChillDebuff, ConsumesInputs.ElixirOfFrostPower],
	excludeBuffDebuffInputs: [],
	// Inputs to include in the 'Other' section on the settings tab.
	otherInputs: {
		inputs: [OtherInputs.DistanceFromTarget],
	},
	itemSwapConfig: {
		itemSlots: [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand],
	},
	customSections: [
		// TotemsSection,
	],
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
		return (player.getTalents().lavaBurst ? Presets.APLDefault : Presets.APLStormcaller).rotation.rotation!;
	},

	raidSimPresets: [
		{
			spec: Spec.SpecElementalShaman,
			tooltip: specNames[Spec.SpecElementalShaman],
			defaultName: 'Elemental',
			iconUrl: getSpecIcon(Class.ClassShaman, 0),

			talents: Presets.DefaultTalents.data,
			specOptions: Presets.DefaultOptions,
			consumes: Presets.DefaultConsumes,
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceUnknown,
				[Faction.Horde]: Race.RaceTroll,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: Presets.DefaultGear.gear,
					2: Presets.DefaultGear.gear,
				},
				[Faction.Horde]: {
					1: Presets.DefaultGear.gear,
					2: Presets.DefaultGear.gear,
				},
			},
		},
	],
});

export class ElementalShamanSimUI extends IndividualSimUI<Spec.SpecElementalShaman> {
	constructor(parentElem: HTMLElement, player: Player<Spec.SpecElementalShaman>) {
		super(parentElem, player, SPEC_CONFIG);
	}
}
