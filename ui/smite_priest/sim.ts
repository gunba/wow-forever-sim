import * as BuffDebuffInputs from '../core/components/inputs/buffs_debuffs';
import * as OtherInputs from '../core/components/other_inputs.js';
import * as Mechanics from '../core/constants/mechanics.js';
import { ClassicPhase } from '../core/constants/other.js';
import { IndividualSimUI, registerSpecConfig } from '../core/individual_sim_ui.js';
import { Player } from '../core/player.js';
import { Class, Faction, PartyBuffs, PseudoStat, Race, Spec, Stat } from '../core/proto/common.js';
import { Stats } from '../core/proto_utils/stats.js';
import { getSpecIcon, specNames } from '../core/proto_utils/utils.js';
import * as SmitePriestInputs from './inputs.js';
import * as Presets from './presets.js';

const SPEC_CONFIG = registerSpecConfig(Spec.SpecSmitePriest, {
	cssClass: 'smite-priest-sim-ui',
	cssScheme: 'priest',
	// List any known bugs / issues here and they'll be shown on the site.
	knownIssues: ['Power in Light, Holy Precision and Mental Agility were all shown at rank 1 only, so their per-rank scaling is extrapolated.'],

	// All stats for which EP should be calculated.
	epStats: [
		// Attributes
		Stat.StatIntellect,
		Stat.StatSpirit,
		// Spell
		Stat.StatSpellPower,
		Stat.StatSpellDamage,
		Stat.StatHolyPower,
		Stat.StatSpellHit,
		Stat.StatSpellCrit,
		Stat.StatSpellHaste,
		Stat.StatMP5,
	],
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
		Stat.StatHolyPower,
		Stat.StatSpellHit,
		Stat.StatSpellCrit,
		Stat.StatMP5,
	],
	displayPseudoStats: [PseudoStat.PseudoStatCastSpeedMultiplier],

	modifyDisplayStats: (player: Player<Spec.SpecSmitePriest>) => {
		let stats = new Stats();
		stats = stats.addPseudoStat(PseudoStat.PseudoStatSchoolHitHoly, player.getTalents().holyPrecision * 6 * Mechanics.SPELL_HIT_RATING_PER_HIT_CHANCE);

		return {
			talents: stats,
		};
	},

	defaults: {
		race: Race.RaceTroll,
		// Default equipped gear.
		gear: Presets.DefaultGear.gear,
		// Default EP weights for sorting gear in the gear picker.
		epWeights: Stats.fromMap({
			[Stat.StatIntellect]: 0.16,
			// Spiritual Guidance turns spirit into spell damage as well as regen in
			// Forever, so it is worth more here than to the other casters.
			[Stat.StatSpirit]: 0.12,
			[Stat.StatSpellPower]: 1,
			[Stat.StatSpellDamage]: 1,
			[Stat.StatHolyPower]: 1,
			[Stat.StatSpellHit]: 5.51,
			[Stat.StatSpellCrit]: 6.5,
			[Stat.StatSpellHaste]: 1.65,
			[Stat.StatMP5]: 0.1,
			[Stat.StatFireResistance]: 0.5,
		}),
		// Default consumes settings.
		consumes: Presets.DefaultConsumes,
		// Default talents.
		talents: Presets.DefaultTalents.data,
		// Default spec-specific settings.
		specOptions: Presets.DefaultOptions,
		// Default raid/party buffs settings.
		raidBuffs: Presets.DefaultRaidBuffs,

		partyBuffs: PartyBuffs.create({}),

		individualBuffs: Presets.DefaultIndividualBuffs,

		debuffs: Presets.DefaultDebuffs,

		other: Presets.OtherDefaults,
	},

	// IconInputs to include in the 'Player' section on the settings tab.
	playerIconInputs: [SmitePriestInputs.ArmorInput],
	// Buff and Debuff inputs to include/exclude, overriding the EP-based defaults.
	includeBuffDebuffInputs: [BuffDebuffInputs.BlessingOfWisdom, BuffDebuffInputs.ManaSpringTotem, BuffDebuffInputs.StaminaBuff],
	excludeBuffDebuffInputs: [],
	// Inputs to include in the 'Other' section on the settings tab.
	otherInputs: {
		inputs: [OtherInputs.TankAssignment, OtherInputs.ChannelClipDelay, OtherInputs.DistanceFromTarget],
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

	autoRotation: _ => {
		return Presets.DefaultAPL.rotation.rotation!;
	},

	raidSimPresets: [
		{
			spec: Spec.SpecSmitePriest,
			tooltip: specNames[Spec.SpecSmitePriest],
			defaultName: 'Smite',
			iconUrl: getSpecIcon(Class.ClassPriest, 0),

			talents: Presets.DefaultTalents.data,
			specOptions: Presets.DefaultOptions,
			consumes: Presets.DefaultConsumes,
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceDwarf,
				[Faction.Horde]: Race.RaceUndead,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: Presets.DefaultGear.gear,
				},
				[Faction.Horde]: {
					1: Presets.DefaultGear.gear,
				},
			},
		},
	],
});

export class SmitePriestSimUI extends IndividualSimUI<Spec.SpecSmitePriest> {
	constructor(parentElem: HTMLElement, player: Player<Spec.SpecSmitePriest>) {
		super(parentElem, player, SPEC_CONFIG);
	}
}
