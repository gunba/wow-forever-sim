import * as Tooltips from '../../constants/tooltips';
import { Encounter } from '../../encounter';
import { IndividualSimUI, InputSection } from '../../individual_sim_ui';
import { Player } from '../../player';
import { Ruleset } from '../../proto/api';
import { Consumes, Debuffs, HealingModel, IndividualBuffs, ItemSwap, PartyBuffs, Profession, RaidBuffs, Spec } from '../../proto/common';
import { SavedEncounter, SavedSettings } from '../../proto/ui';
import { professionNames, raceNames } from '../../proto_utils/names';
import { isHealingSpec, isTankSpec, specToEligibleRaces } from '../../proto_utils/utils';
import { EventID, TypedEvent } from '../../typed_event';
import { getEnumValues } from '../../utils';
import { BooleanPicker } from '../boolean_picker';
import { ContentBlock } from '../content_block';
import { EncounterPicker } from '../encounter_picker';
import { EnumPicker } from '../enum_picker';
import { IconEnumPicker } from '../icon_enum_picker';
import * as IconInputs from '../icon_inputs';
import { Input } from '../input';
import * as BuffDebuffInputs from '../inputs/buffs_debuffs';
import { relevantStatOptions } from '../inputs/stat_options';
import { MultiIconPicker, MultiIconPickerItemConfig } from '../multi_icon_picker';
import { NumberPicker } from '../number_picker';
import { SavedDataManager } from '../saved_data_manager';
import { SimTab } from '../sim_tab';
import { IsbConfig, StormstrikeConfig } from './../other_inputs';
import { ConsumesPicker } from './consumes_picker';
import { ItemSwapPicker } from './item_swap_picker';
import { PresetConfigurationPicker } from './preset_configuration_picker';

export class SettingsTab extends SimTab {
	protected simUI: IndividualSimUI<Spec>;

	readonly leftPanel: HTMLElement;
	readonly rightPanel: HTMLElement;

	readonly column1: HTMLElement = this.buildColumn(1, 'settings-left-col');
	readonly column2: HTMLElement = this.buildColumn(2, 'settings-left-col');
	readonly column3: HTMLElement = this.buildColumn(3, 'settings-left-col');
	readonly column4?: HTMLElement;

	constructor(parentElem: HTMLElement, simUI: IndividualSimUI<Spec>) {
		super(parentElem, simUI, { identifier: 'settings-tab', title: 'Settings' });
		this.simUI = simUI;

		this.leftPanel = document.createElement('div');
		this.leftPanel.classList.add('settings-tab-left', 'tab-panel-left');

		this.leftPanel.appendChild(this.column1);
		this.leftPanel.appendChild(this.column2);
		this.leftPanel.appendChild(this.column3);

		// The 4th column is only used in the raid sim player editor to spread out player settings
		if (this.simUI.isWithinRaidSim) {
			this.column4 = this.buildColumn(4, 'settings-left-col');
			this.leftPanel.appendChild(this.column4);
		}

		this.rightPanel = document.createElement('div');
		this.rightPanel.classList.add('settings-tab-right', 'tab-panel-right', 'within-raid-sim-hide');

		this.contentContainer.appendChild(this.leftPanel);
		this.contentContainer.appendChild(this.rightPanel);

		this.buildTabContent();
	}

	protected buildTabContent() {
		if (!this.simUI.isWithinRaidSim) {
			this.buildEncounterSettings();
		}

		this.simUI.sim.waitForInit().then(() => {
			this.buildPlayerSettings();
			this.buildCustomSettingsSections();
			this.buildConsumesSection();
			this.buildOtherSettings();
			this.buildIsbSettings();
			this.buildStormstrikeSettings();

			if (!this.simUI.isWithinRaidSim) {
				this.buildBuffsSettings();
				this.buildWorldBuffsSettings();
				this.buildDebuffsSettings();
				this.buildPresetConfigurationPicker();
				this.buildSavedDataPickers();
			}
		});
	}

	private buildEncounterSettings() {
		const contentBlock = new ContentBlock(this.column1, 'encounter-settings', {
			header: { title: 'Encounter' },
		});

		new EncounterPicker(contentBlock.bodyElement, this.simUI.sim.encounter, this.simUI.individualConfig.encounterPicker, this.simUI);
	}

	private buildPlayerSettings() {
		const contentBlock = new ContentBlock(this.column1, 'player-settings', {
			header: { title: 'Player' },
		});

		const playerIconGroup = Input.newGroupContainer();
		playerIconGroup.classList.add('player-icon-group', 'icon-group');
		contentBlock.bodyElement.appendChild(playerIconGroup);

		this.configureIconSection(
			playerIconGroup,
			this.simUI.individualConfig.playerIconInputs.map(iconInput => IconInputs.buildIconInput(playerIconGroup, this.simUI.player, iconInput)),
			true,
		);

		const races = specToEligibleRaces[this.simUI.player.spec];
		new EnumPicker(contentBlock.bodyElement, this.simUI.player, {
			id: 'player-race',
			label: 'Race',
			values: races.map(race => {
				return {
					name: raceNames.get(race)!,
					value: race,
				};
			}),
			changedEvent: player => player.raceChangeEmitter,
			getValue: player => player.getRace(),
			setValue: (eventID, player, newValue) => player.setRace(eventID, newValue),
		});

		if (this.simUI.individualConfig.playerInputs?.inputs.length) {
			this.configureInputSection(contentBlock.bodyElement, this.simUI.individualConfig.playerInputs);
		}

		const professionGroup = Input.newGroupContainer();
		contentBlock.bodyElement.appendChild(professionGroup);

		const professions = getEnumValues(Profession) as Array<Profession>;
		new EnumPicker(professionGroup, this.simUI.player, {
			id: 'player-profession-1',
			label: 'Profession 1',
			values: professions.map(p => {
				return {
					name: professionNames.get(p)!,
					value: p,
				};
			}),
			changedEvent: player => player.professionChangeEmitter,
			getValue: player => player.getProfession1(),
			setValue: (eventID, player, newValue) => player.setProfession1(eventID, newValue),
		});

		new EnumPicker(professionGroup, this.simUI.player, {
			id: 'player-profession-2',
			label: 'Profession 2',
			values: professions.map(p => {
				return {
					name: professionNames.get(p)!,
					value: p,
				};
			}),
			changedEvent: player => player.professionChangeEmitter,
			getValue: player => player.getProfession2(),
			setValue: (eventID, player, newValue) => player.setProfession2(eventID, newValue),
		});
	}

	private buildCustomSettingsSections() {
		(this.simUI.individualConfig.customSections || []).forEach(customSection => {
			const section = customSection(this.column2, this.simUI);
			section.rootElem.classList.add('custom-section');
		});
	}

	private buildConsumesSection() {
		const column = this.simUI.isWithinRaidSim ? this.column3 : this.column2;
		const contentBlock = new ContentBlock(column, 'consumes-settings', {
			header: { title: 'Consumables' },
		});

		new ConsumesPicker(contentBlock.bodyElement, this.simUI);
	}

	private buildOtherSettings() {
		// const column = this.simUI.isWithinRaidSim ? this.column4 : this.column2;
		const settings = this.simUI.individualConfig.otherInputs?.inputs.filter(
			inputs => !inputs.extraCssClasses || !inputs.extraCssClasses?.includes('within-raid-sim-hide'),
		) ?? [];

		const itemSwapConfig = this.simUI.individualConfig.itemSwapConfig;
		const tierSupported = !isTankSpec(this.simUI.player.spec) && !isHealingSpec(this.simUI.player.spec);
		const mp5Supported = ![Spec.SpecWarrior, Spec.SpecTankWarrior, Spec.SpecRogue].includes(this.simUI.player.spec);

		if (settings.length || itemSwapConfig?.itemSlots.length || tierSupported || mp5Supported) {
			const contentBlock = new ContentBlock(this.column2, 'other-settings', {
				header: { title: 'Other' },
			});

			if (settings.length) {
				this.configureInputSection(contentBlock.bodyElement, this.simUI.individualConfig.otherInputs);
				contentBlock.bodyElement.querySelectorAll('.input-root').forEach(elem => {
					elem.classList.add('input-inline');
				});
			}

			if (itemSwapConfig?.itemSlots.length) {
				new ItemSwapPicker(contentBlock.bodyElement, this.simUI, this.simUI.player, itemSwapConfig);
			}
			if (mp5Supported) {
				new BooleanPicker(contentBlock.bodyElement, this.simUI.player, {
					id: 'forever-mp5-per-second',
					label: 'MP5 acts per second',
					labelTooltip: 'Provisional model: each point of the MP5 stat restores one mana per second instead of one per five seconds. Applies to static and temporary MP5 from gear, buffs and consumes. Spirit regeneration and direct mana returns are unchanged. Not verified in the beta; off in the published rankings.',
					changedEvent: player => TypedEvent.onAny([player.miscOptionsChangeEmitter, player.sim.rulesetChangeEmitter]),
					getValue: player => player.getForeverMp5PerSecond(),
					setValue: (eventID, player, value) => player.setForeverMp5PerSecond(eventID, value),
					showWhen: player => player.sim.getRuleset() == Ruleset.RulesetForever,
				});
			}
			if (tierSupported) {
				new BooleanPicker(contentBlock.bodyElement, this.simUI.player, {
					id: 'forever-tier1-bonuses',
					label: 'Tier 1 bonuses',
					labelTooltip: 'Apply all four bonuses from this role’s Forever Tier 1 set independently of gear. Adds no raid-item stats. Creature-type restrictions still apply.',
					changedEvent: player => TypedEvent.onAny([player.miscOptionsChangeEmitter, player.sim.rulesetChangeEmitter]),
					getValue: player => player.getForeverTier1Bonuses(),
					setValue: (eventID, player, value) => player.setForeverTier1Bonuses(eventID, value),
					showWhen: player => player.sim.getRuleset() == Ruleset.RulesetForever,
				});
			}
			const scalingNote = document.createElement('p');
			scalingNote.classList.add('small', 'text-warning', 'mt-2');
			const updateScalingNote = () => {
				const scale = this.simUI.player.getEquipmentScale();
				scalingNote.hidden = scale === 1;
				scalingNote.textContent = `Hypothetical equipment multiplier: ×${scale.toFixed(2)}. Item stats and weapon damage are scaled; tooltips show original items. Paid hit must be recalculated outside the UI.`;
			};
			this.simUI.player.miscOptionsChangeEmitter.on(updateScalingNote);
			updateScalingNote();
			contentBlock.bodyElement.appendChild(scalingNote);
			if ([Spec.SpecFeralDruid, Spec.SpecRogue].includes(this.simUI.player.spec)) {
				const note = document.createElement('p');
				note.classList.add('small', 'text-muted', 'mt-2');
				note.textContent = 'Energy regeneration scales with general haste as a model assumption. Attack-speed-only effects do not increase it.';
				const update = () => { note.hidden = this.simUI.sim.getRuleset() != Ruleset.RulesetForever; };
				this.simUI.sim.rulesetChangeEmitter.on(update);
				update();
				contentBlock.bodyElement.appendChild(note);
			}
		}
	}

	private buildIsbSettings() {
		if (!this.simUI.isWithinRaidSim) {
			const contentBlock = new ContentBlock(this.column1, 'other-settings', {
				header: { title: 'Improved Shadow Bolt' },
			});

			this.configureInputSection(contentBlock.bodyElement, IsbConfig);

			TypedEvent.onAny([this.simUI.player.talentsChangeEmitter, this.simUI.player.getRaid()!.debuffsChangeEmitter]).on(() => {
				const externalIsb = this.simUI.player.getRaid()?.getDebuffs()?.improvedShadowBolt == true;
				if (externalIsb) {
					contentBlock.rootElem.classList.remove('hide');
				} else {
					contentBlock.rootElem.classList.add('hide');
				}
			});
		}
	}

	private buildStormstrikeSettings() {
		if (!this.simUI.isWithinRaidSim) {
			const contentBlock = new ContentBlock(this.column1, 'other-settings', {
				header: { title: 'Stormstrike' },
			});

			this.configureInputSection(contentBlock.bodyElement, StormstrikeConfig);

			this.simUI.player.getRaid()!.debuffsChangeEmitter.on(() => {
				if (this.simUI.player.getRaid()?.getDebuffs()?.stormstrike) {
					contentBlock.rootElem.classList.remove('hide');
				} else {
					contentBlock.rootElem.classList.add('hide');
				}
			});
		}
	}

	private buildBuffsSettings() {
		const buffOptions = relevantStatOptions(BuffDebuffInputs.RAID_BUFFS_CONFIG, this.simUI);
		const miscBuffOptions = relevantStatOptions(BuffDebuffInputs.MISC_BUFFS_CONFIG, this.simUI);

		const contentBlock = new ContentBlock(this.column3, 'buffs-settings', {
			header: { title: 'Raid Buffs', tooltip: Tooltips.BUFFS_SECTION },
		});

		this.configureIconSection(
			contentBlock.bodyElement,
			buffOptions.map(options => options.picker && new options.picker(contentBlock.bodyElement, this.simUI.player, options.config as any, this.simUI)),
		);

		if (miscBuffOptions.length) {
			new MultiIconPicker(
				contentBlock.bodyElement,
				this.simUI.player,
				{
					values: miscBuffOptions.map(options => options.config) as Array<MultiIconPickerItemConfig<Player<Spec>>>,
					label: 'Misc Buffs',
				},
				this.simUI,
			);
		}
	}

	private buildWorldBuffsSettings() {
		const contentBlock = new ContentBlock(this.column3, 'world-buffs-settings', {
			header: { title: 'World Buffs', tooltip: Tooltips.WORLD_BUFFS_SECTION },
		});

		// World buffs do not work in Forever raids, so the section is only offered under Classic
		// rules. The engine ignores them under Forever as well, so a saved setting cannot sneak
		// them back in.
		const showForRuleset = () => {
			contentBlock.rootElem.hidden = this.simUI.sim.getRuleset() != Ruleset.RulesetClassic;
		};
		showForRuleset();
		this.simUI.sim.rulesetChangeEmitter.on(showForRuleset);

		const saygesOptions = relevantStatOptions(BuffDebuffInputs.SAYGES_CONFIG, this.simUI);
		new IconEnumPicker(contentBlock.bodyElement, this.simUI.player, BuffDebuffInputs.SaygesDarkFortune(saygesOptions));

		const worldBuffOptions = relevantStatOptions(BuffDebuffInputs.WORLD_BUFFS_CONFIG, this.simUI);
		this.configureIconSection(
			contentBlock.bodyElement,
			worldBuffOptions.map(
				options => options.picker && new options.picker(contentBlock.bodyElement, this.simUI.player, options.config as any, this.simUI),
			),
		);
	}

	private buildDebuffsSettings() {
		const debuffOptions = relevantStatOptions(BuffDebuffInputs.DEBUFFS_CONFIG, this.simUI);
		const miscDebuffOptions = relevantStatOptions(BuffDebuffInputs.MISC_DEBUFFS_CONFIG, this.simUI);

		if (!debuffOptions.length && !miscDebuffOptions.length) return;

		const contentBlock = new ContentBlock(this.column3, 'debuffs-settings', {
			header: { title: 'Debuffs', tooltip: Tooltips.DEBUFFS_SECTION },
		});

		this.configureIconSection(
			contentBlock.bodyElement,
			debuffOptions.map(options => options.picker && new options.picker(contentBlock.bodyElement, this.simUI.player, options.config as any, this.simUI)),
		);

		if (miscDebuffOptions.length) {
			new MultiIconPicker(
				contentBlock.bodyElement,
				this.simUI.player,
				{
					values: miscDebuffOptions.map(options => options.config) as Array<MultiIconPickerItemConfig<Player<Spec>>>,
					label: 'Misc Debuffs',
				},
				this.simUI,
			);
		}

		// In case no debuffs are active, this will fire a change event to update the pickers
		this.simUI.player.getRaid()?.debuffsChangeEmitter.emit(TypedEvent.nextEventID());
	}

	private buildPresetConfigurationPicker() {
		new PresetConfigurationPicker(this.rightPanel, this.simUI, ['encounter', 'options']);
	}

	private buildSavedDataPickers() {
		const savedEncounterManager = new SavedDataManager<Encounter, SavedEncounter>(this.rightPanel, this.simUI.sim.encounter, {
			label: 'Encounter',
			header: { title: 'Saved Encounters' },
			storageKey: this.simUI.getSavedEncounterStorageKey(),
			getData: (encounter: Encounter) => SavedEncounter.create({ encounter: encounter.toProto() }),
			setData: (eventID: EventID, encounter: Encounter, newEncounter: SavedEncounter) => encounter.fromProto(eventID, newEncounter.encounter!),
			changeEmitters: [this.simUI.sim.encounter.changeEmitter],
			equals: (a: SavedEncounter, b: SavedEncounter) => SavedEncounter.equals(a, b),
			toJson: (a: SavedEncounter) => SavedEncounter.toJson(a),
			fromJson: (obj: any) => SavedEncounter.fromJson(obj),
		});

		const savedSettingsManager = new SavedDataManager<IndividualSimUI<any>, SavedSettings>(this.rightPanel, this.simUI, {
			label: 'Settings',
			header: { title: 'Saved Settings' },
			storageKey: this.simUI.getSavedSettingsStorageKey(),
			getData: (simUI: IndividualSimUI<any>) => {
				const player = simUI.player;
				return SavedSettings.create({
					raidBuffs: simUI.sim.raid.getBuffs(),
					partyBuffs: player.getParty()?.getBuffs() || PartyBuffs.create(),
					playerBuffs: player.getBuffs(),
					debuffs: simUI.sim.raid.getDebuffs(),
					consumes: player.getConsumes(),
					race: player.getRace(),
					professions: player.getProfessions(),
					enableItemSwap: player.getEnableItemSwap(),
					itemSwap: player.getItemSwapGear().toProto(),
					reactionTimeMs: player.getReactionTime(),
					channelClipDelayMs: player.getChannelClipDelay(),
					inFrontOfTarget: player.getInFrontOfTarget(),
					foreverTier1Bonuses: player.getForeverTier1Bonuses(),
					foreverMp5PerSecond: player.getForeverMp5PerSecond(),
					distanceFromTarget: player.getDistanceFromTarget(),
					healingModel: player.getHealingModel(),
				});
			},
			setData: (eventID: EventID, simUI: IndividualSimUI<any>, newSettings: SavedSettings) => {
				TypedEvent.freezeAllAndDo(() => {
					simUI.sim.raid.setBuffs(eventID, newSettings.raidBuffs || RaidBuffs.create());
					simUI.sim.raid.setDebuffs(eventID, newSettings.debuffs || Debuffs.create());
					const party = simUI.player.getParty();
					if (party) {
						party.setBuffs(eventID, newSettings.partyBuffs || PartyBuffs.create());
					}
					simUI.player.setBuffs(eventID, newSettings.playerBuffs || IndividualBuffs.create());
					simUI.player.setConsumes(eventID, newSettings.consumes || Consumes.create());
					simUI.player.setRace(eventID, newSettings.race);
					simUI.player.setProfessions(eventID, newSettings.professions);
					simUI.player.setEnableItemSwap(eventID, newSettings.enableItemSwap);
					simUI.player.setItemSwapGear(eventID, simUI.sim.db.lookupItemSwap(newSettings.itemSwap || ItemSwap.create()));
					simUI.player.setReactionTime(eventID, newSettings.reactionTimeMs);
					simUI.player.setChannelClipDelay(eventID, newSettings.channelClipDelayMs);
					simUI.player.setInFrontOfTarget(eventID, newSettings.inFrontOfTarget);
					simUI.player.setForeverTier1Bonuses(eventID, newSettings.foreverTier1Bonuses);
					simUI.player.setForeverMp5PerSecond(eventID, newSettings.foreverMp5PerSecond);
					simUI.player.setDistanceFromTarget(eventID, newSettings.distanceFromTarget);
					simUI.player.setHealingModel(eventID, newSettings.healingModel || HealingModel.create());
				});
			},
			changeEmitters: [
				this.simUI.sim.raid.buffsChangeEmitter,
				this.simUI.sim.raid.debuffsChangeEmitter,
				this.simUI.player.getParty()!.buffsChangeEmitter,
				this.simUI.player.buffsChangeEmitter,
				this.simUI.player.consumesChangeEmitter,
				this.simUI.player.raceChangeEmitter,
				this.simUI.player.professionChangeEmitter,
				this.simUI.player.itemSwapChangeEmitter,
				this.simUI.player.miscOptionsChangeEmitter,
				this.simUI.player.inFrontOfTargetChangeEmitter,
				this.simUI.player.distanceFromTargetChangeEmitter,
				this.simUI.player.healingModelChangeEmitter,
			],
			equals: (a: SavedSettings, b: SavedSettings) => SavedSettings.equals(a, b),
			toJson: (a: SavedSettings) => SavedSettings.toJson(a),
			fromJson: (obj: any) => SavedSettings.fromJson(obj),
		});

		this.simUI.sim.waitForInit().then(() => {
			savedEncounterManager.loadUserData();
			savedSettingsManager.loadUserData();
		});
	}

	private configureInputSection(sectionElem: HTMLElement, sectionConfig: InputSection) {
		sectionConfig.inputs.forEach(inputConfig => {
			if (inputConfig.type == 'number') {
				new NumberPicker(sectionElem, this.simUI.player, inputConfig);
			} else if (inputConfig.type == 'boolean') {
				new BooleanPicker(sectionElem, this.simUI.player, { ...inputConfig, reverse: true });
			} else if (inputConfig.type == 'enum') {
				new EnumPicker(sectionElem, this.simUI.player, inputConfig);
			}
		});
	}

	private configureIconSection(sectionElem: HTMLElement, iconPickers: Array<any>, adjustColumns?: boolean) {
		if (iconPickers.length == 0) {
			sectionElem.classList.add('hide');
		} else if (adjustColumns) {
			if (iconPickers.length <= 4) {
				sectionElem.style.gridTemplateColumns = `repeat(${iconPickers.length}, 1fr)`;
			} else if (iconPickers.length > 4 && iconPickers.length < 8) {
				sectionElem.style.gridTemplateColumns = `repeat(${Math.ceil(iconPickers.length / 2)}, 1fr)`;
			}
		}
	}
}
