import { BooleanPicker } from '../components/boolean_picker.js';
import { Player } from '../player.js';
import { Ruleset } from '../proto/api.js';
import { Class, ForeverIncomingRageModel, ForeverRevelationModel, Spec, UnitReference } from '../proto/common.js';
import { emptyUnitReference } from '../proto_utils/utils.js';
import { Sim } from '../sim.js';
import { EventID, TypedEvent } from '../typed_event.js';
import { EnumPicker } from './enum_picker.js';
import { Input, InputConfig } from './input.js';

interface ProvisionalNumberConfig extends InputConfig<Player<any>, number | undefined> {
	id: string;
	min: number;
	max?: number;
	defaultSource?: string;
}

// Unlike the general number picker, these assumptions must retain blank vs. zero
// and display their full precision. Invalid edits never replace the saved value.
class ProvisionalNumberPicker extends Input<Player<any>, number | undefined> {
	private readonly inputElem: HTMLInputElement;
	private readonly sourceElem: HTMLElement;
	private readonly defaultSource?: string;

	constructor(parent: HTMLElement, player: Player<any>, config: ProvisionalNumberConfig) {
		super(parent, 'number-picker-root', player, config);
		this.defaultSource = config.defaultSource;
		this.inputElem = document.createElement('input');
		this.inputElem.id = config.id;
		this.inputElem.type = 'number';
		this.inputElem.step = 'any';
		this.inputElem.min = String(config.min);
		if (config.max !== undefined) this.inputElem.max = String(config.max);
		this.inputElem.required = config.defaultSource === undefined;
		this.inputElem.placeholder = config.defaultSource === undefined ? '' : 'Default (blank)';
		this.inputElem.classList.add('form-control');
		this.sourceElem = document.createElement('small');
		this.sourceElem.id = `${config.id}-source`;
		this.sourceElem.classList.add('text-muted');
		this.inputElem.setAttribute('aria-describedby', this.sourceElem.id);
		this.rootElem.append(this.inputElem, this.sourceElem);
		this.inputElem.addEventListener('change', () => {
			const value = this.getInputValue();
			this.inputElem.setCustomValidity(value !== undefined && !Number.isFinite(value) ? 'Enter a finite number.' : '');
			if (!this.inputElem.reportValidity()) return;
			this.inputChanged(TypedEvent.nextEventID());
		}, { signal: this.signal });
		this.init();
	}

	getInputElem(): HTMLInputElement {
		return this.inputElem;
	}

	getInputValue(): number | undefined {
		return this.inputElem.value === '' ? undefined : this.inputElem.valueAsNumber;
	}

	setInputValue(value: number | undefined) {
		this.inputElem.value = value === undefined ? '' : String(value);
		this.inputElem.setCustomValidity('');
		this.sourceElem.textContent = value === undefined ? `Source: ${this.defaultSource}` : 'Source: manual provisional assumption.';
	}
}

// Shared by individual and raid player settings. Returns its disposal callback.
export function makeForeverProvisionalModelInputs(parent: HTMLElement, player: Player<any>): () => void {
	const details = document.createElement('details');
	details.classList.add('mt-3', 'forever-provisional-models');
	const summary = document.createElement('summary');
	summary.textContent = 'Advanced: provisional models';
	details.appendChild(summary);
	parent.appendChild(details);
	const pickers: Array<Input<Player<any>, any>> = [];
	const changedEvent = (p: Player<any>) => p.miscOptionsChangeEmitter;
	const addNote = (text: string) => {
		const note = document.createElement('p');
		note.classList.add('small', 'mt-2');
		note.textContent = text;
		details.appendChild(note);
	};
	addNote('Unverified model assumptions for sensitivity analysis, not established game formulas. Settings are saved with the profile.');

	if ([Spec.SpecWarrior, Spec.SpecTankWarrior, Spec.SpecFeralTankDruid].includes(player.spec)) {
		addNote('Incoming rage: leave a field blank to use its declared default; enter a value for a manual assumption. Zero coefficient disables incoming rage. Bounds are model validation limits, not game claims.');
		const addRageInput = (field: keyof ForeverIncomingRageModel, label: string, min: number, max: number, defaultSource: string) => {
			pickers.push(new ProvisionalNumberPicker(details, player, {
				id: `forever-incoming-rage-${field}`,
				label,
				min,
				max,
				defaultSource,
				changedEvent,
				getValue: p => p.getForeverIncomingRageModel()?.[field],
				setValue: (eventID, p, value) => {
					const model = p.getForeverIncomingRageModel() ?? ForeverIncomingRageModel.create();
					model[field] = value;
					p.setForeverIncomingRageModel(eventID, model);
				},
			}));
		};
		addRageInput('coefficient', 'Incoming rage coefficient (0–1000)', 0, 1000, 'declared central coefficient 10.');
		addRageInput('referenceArmorOverride', 'Reference armor fraction (0–0.95)', 0, 0.95,
			'clamp(ExpectedStatCreatureArmor / (Armor + ArmorConstant), 0.20, 0.40).');
		addRageInput('expectedHealthOverride', 'Expected creature health (1–1,000,000,000)', 1, 1e9,
			'ExpectedStatCreatureHealth at the player level.');
		pickers.push(new ProvisionalNumberPicker(details, player, {
			id: 'forever-demoralizing-threat',
			label: 'Provisional Demoralizing threat per target',
			labelTooltip: 'Flat threat per affected target for Demoralizing Shout/Roar, before stance/global modifiers. The exact Forever amount is unpublished. Blank retains the Classic-derived rank convention; explicit zero is a legacy-zero sensitivity. The finite 0–1,000,000 bounds are model validation limits, not game claims.',
			min: 0,
			max: 1e6,
			defaultSource: 'retained Classic-derived rank convention (not a verified Forever amount).',
			changedEvent,
			getValue: p => p.getForeverDemoralizingThreat(),
			setValue: (eventID, p, value) => p.setForeverDemoralizingThreat(eventID, value),
		}));
	}

	addNote('Revelation (enchant 8217) requires an explicitly enabled provisional event model; no hidden proc chance is supplied (initial base chance 0%). For eligible noncritical direct events: p = baseChance × (1 − clamp(effectiveCrit, 0, 1))^exponent. Exponent 0 is flat conditional chance, not per cast. Source-backed channel damage children qualify through legacy tick aliases; ordinary DoT/HoT ticks, wands and channel containers do not. The first eligible crit-roll sample reserves one charge before projectile travel, not every target of a spell. A miss that never rolls crit preserves the charge; this is separate from allowing misses to trigger a new proc. The law, miss policy and event-sampling/consumption conventions are unverified script assumptions.');
	if (player.getClass() === Class.ClassWarrior || player.getClass() === Class.ClassRogue) {
		addNote('No Revelation class payload is present in the source for Warrior/Rogue. These controls can be saved, but equipping enchant 8217 on this class produces an explicit simulation error; this is an unsupported payload, not an item class restriction.');
	}
	pickers.push(new BooleanPicker(details, player, {
		id: 'forever-revelation-enabled',
		label: 'Enable provisional Revelation model',
		changedEvent,
		getValue: p => p.getForeverRevelationModel()?.enabled ?? false,
		setValue: (eventID, p, enabled) => {
			const model = p.getForeverRevelationModel() ?? ForeverRevelationModel.create();
			model.enabled = enabled;
			p.setForeverRevelationModel(eventID, model);
		},
	}));
	const enableWhen = (p: Player<any>) => p.getForeverRevelationModel()?.enabled === true;
	pickers.push(new ProvisionalNumberPicker(details, player, {
		id: 'forever-revelation-base-chance',
		label: 'Revelation base chance % (0–100)',
		min: 0,
		max: 100,
		changedEvent,
		enableWhen,
		getValue: p => (p.getForeverRevelationModel()?.baseChance ?? 0) * 100,
		setValue: (eventID, p, value) => {
			if (value === undefined) return;
			const model = p.getForeverRevelationModel() ?? ForeverRevelationModel.create();
			model.baseChance = value / 100;
			p.setForeverRevelationModel(eventID, model);
		},
	}));
	pickers.push(new ProvisionalNumberPicker(details, player, {
		id: 'forever-revelation-crit-exponent',
		label: 'Revelation crit exponent (≥0)',
		min: 0,
		changedEvent,
		enableWhen,
		getValue: p => p.getForeverRevelationModel()?.critExponent ?? 0,
		setValue: (eventID, p, value) => {
			if (value === undefined) return;
			const model = p.getForeverRevelationModel() ?? ForeverRevelationModel.create();
			model.critExponent = value;
			p.setForeverRevelationModel(eventID, model);
		},
	}));
	pickers.push(new BooleanPicker(details, player, {
		id: 'forever-revelation-trigger-on-miss',
		label: 'Allow Revelation on eligible misses',
		changedEvent,
		enableWhen,
		getValue: p => p.getForeverRevelationModel()?.triggerOnMiss ?? false,
		setValue: (eventID, p, triggerOnMiss) => {
			const model = p.getForeverRevelationModel() ?? ForeverRevelationModel.create();
			model.triggerOnMiss = triggerOnMiss;
			p.setForeverRevelationModel(eventID, model);
		},
	}));
	const updateVisibility = () => { details.hidden = player.sim.getRuleset() !== Ruleset.RulesetForever; };
	const rulesetListener = player.sim.rulesetChangeEmitter.on(updateVisibility);
	updateVisibility();
	return () => {
		rulesetListener.dispose();
		pickers.forEach(picker => picker.dispose());
	};
}

export function makeShow1hWeaponsSelector(parent: HTMLElement, sim: Sim): BooleanPicker<Sim> {
	parent.classList.remove('hide');
	return new BooleanPicker<Sim>(parent, sim, {
		id: 'show-1h-weapons-selector',
		extraCssClasses: ['show-1h-weapons-selector', 'mb-0'],
		label: '1H',
		inline: true,
		changedEvent: (sim: Sim) => sim.filtersChangeEmitter,
		getValue: (sim: Sim) => sim.getFilters().oneHandedWeapons,
		setValue: (eventID: EventID, sim: Sim, newValue: boolean) => {
			const filters = sim.getFilters();
			filters.oneHandedWeapons = newValue;
			sim.setFilters(eventID, filters);
		},
	});
}

export function makeShow2hWeaponsSelector(parent: HTMLElement, sim: Sim): BooleanPicker<Sim> {
	parent.classList.remove('hide');
	return new BooleanPicker<Sim>(parent, sim, {
		id: 'show-2h-weapons-selector',
		extraCssClasses: ['show-2h-weapons-selector', 'mb-0'],
		label: '2H',
		inline: true,
		changedEvent: (sim: Sim) => sim.filtersChangeEmitter,
		getValue: (sim: Sim) => sim.getFilters().twoHandedWeapons,
		setValue: (eventID: EventID, sim: Sim, newValue: boolean) => {
			const filters = sim.getFilters();
			filters.twoHandedWeapons = newValue;
			sim.setFilters(eventID, filters);
		},
	});
}

export function makeShowEPValuesSelector(parent: HTMLElement, sim: Sim): BooleanPicker<Sim> {
	return new BooleanPicker<Sim>(parent, sim, {
		id: 'show-ep-values-selector',
		extraCssClasses: ['show-ep-values-selector', 'input-inline', 'mb-0'],
		label: 'Show EP',
		inline: true,
		changedEvent: (sim: Sim) => sim.showEPValuesChangeEmitter,
		getValue: (sim: Sim) => sim.getShowEPValues(),
		setValue: (eventID: EventID, sim: Sim, newValue: boolean) => {
			sim.setShowEPValues(eventID, newValue);
		},
	});
}

export function makePhaseSelector(parent: HTMLElement, sim: Sim): EnumPicker<Sim> {
	return new EnumPicker<Sim>(parent, sim, {
		id: 'phase-selector',
		extraCssClasses: ['phase-selector'],
		values: [
			{ name: 'Phase 6 - Naxx', value: 6 },
			{ name: 'Phase 5 - AQ', value: 5 },
			{ name: 'Phase 4 - ZG/AB', value: 4 },
			{ name: 'Phase 3 - BWL/R14', value: 3 },
			{ name: 'Phase 2 - DM/WSG/AV', value: 2 },
			{ name: 'Phase 1 - MC/Ony', value: 1 },
		],
		changedEvent: (sim: Sim) => sim.phaseChangeEmitter,
		getValue: (sim: Sim) => sim.getPhase(),
		setValue: (eventID: EventID, sim: Sim, newValue: number) => {
			sim.setPhase(eventID, newValue);
		},
	});
}

export const ReactionTime = {
	id: 'reaction-time',
	type: 'number' as const,
	label: 'Reaction Time',
	labelTooltip: "Reaction time of the player, in milliseconds. Used with certain APL values (such as 'Aura Is Active With Reaction Time').",
	changedEvent: (player: Player<any>) => player.miscOptionsChangeEmitter,
	getValue: (player: Player<any>) => player.getReactionTime(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setReactionTime(eventID, newValue);
	},
};

export const ChannelClipDelay = {
	id: 'channel-clip-delay',
	type: 'number' as const,
	label: 'Channel Clip Delay',
	labelTooltip:
		'Clip delay following channeled spells, in milliseconds. This delay occurs following any full or partial channel ending after the GCD becomes available, due to the player not being able to queue the next spell.',
	changedEvent: (player: Player<any>) => player.miscOptionsChangeEmitter,
	getValue: (player: Player<any>) => player.getChannelClipDelay(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setChannelClipDelay(eventID, newValue);
	},
};

export const InFrontOfTarget = {
	id: 'in-front-of-target',
	type: 'boolean' as const,
	label: 'In Front of Target',
	labelTooltip: 'Stand in front of the target, causing Blocks and Parries to be included in the attack table.',
	changedEvent: (player: Player<any>) => player.inFrontOfTargetChangeEmitter,
	getValue: (player: Player<any>) => player.getInFrontOfTarget(),
	setValue: (eventID: EventID, player: Player<any>, newValue: boolean) => {
		player.setInFrontOfTarget(eventID, newValue);
	},
};

export const DistanceFromTarget = {
	id: 'distance-from-target',
	type: 'number' as const,
	label: 'Distance From Target',
	labelTooltip: 'Distance from targets, in yards. Used to calculate travel time for certain spells.',
	changedEvent: (player: Player<any>) => player.distanceFromTargetChangeEmitter,
	getValue: (player: Player<any>) => player.getDistanceFromTarget(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setDistanceFromTarget(eventID, newValue);
	},
};

export const IsbSbFrequencey = {
	id: 'isb-sb-frequency',
	type: 'number' as const,
	label: 'SB Frequency',
	labelTooltip: 'How often a Shadow Bolt is cast by the external warlock.',
	float: true,
	defaultValue: 3.0,
	inline: true,
	changedEvent: (player: Player<any>) => TypedEvent.onAny([player.changeEmitter, player.getRaid()!.debuffsChangeEmitter]),
	getValue: (player: Player<any>) => player.getIsbSbFrequency(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setIsbSbFrequency(eventID, newValue);
	},
	showWhen: (player: Player<any>) => player.getRaid()?.getDebuffs().improvedShadowBolt === true,
};

export const IsbCrit = {
	id: 'isb-sb-crit',
	type: 'number' as const,
	label: 'SB Crit',
	labelTooltip: 'How often a Shadow Bolt from external warlock is a crit.',
	float: true,
	defaultValue: 25.0,
	inline: true,
	changedEvent: (player: Player<any>) => TypedEvent.onAny([player.changeEmitter, player.getRaid()!.debuffsChangeEmitter]),
	getValue: (player: Player<any>) => player.getIsbCrit(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setIsbCrit(eventID, newValue);
	},
	showWhen: (player: Player<any>) => player.getRaid()?.getDebuffs().improvedShadowBolt === true,
};

export const IsbWarlocks = {
	id: 'isb-warlock',
	type: 'number' as const,
	label: 'SB Warlocks',
	labelTooltip: 'Number of ISB warlocks.',
	defaultValue: 1.0,
	inline: true,
	changedEvent: (player: Player<any>) => TypedEvent.onAny([player.changeEmitter, player.getRaid()!.debuffsChangeEmitter]),
	getValue: (player: Player<any>) => player.getIsbWarlocks(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setIsbWarlocks(eventID, newValue);
	},
	showWhen: (player: Player<any>) => player.getRaid()?.getDebuffs().improvedShadowBolt === true,
};

export const IsbSpriests = {
	id: 'isb-sb-priests',
	type: 'number' as const,
	label: 'Shadow Priests',
	labelTooltip: 'Number of other shadow priests.',
	inline: true,
	changedEvent: (player: Player<any>) => TypedEvent.onAny([player.changeEmitter, player.getRaid()!.debuffsChangeEmitter]),
	getValue: (player: Player<any>) => player.getIsbSpriests(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setIsbSpriests(eventID, newValue);
	},
	showWhen: (player: Player<any>) => player.getRaid()?.getDebuffs().improvedShadowBolt === true,
};

export const IsbConfig = {
	tooltip: 'Improved Shadow Bolt debuff configuration',
	inputs: [IsbSbFrequencey, IsbCrit, IsbWarlocks, IsbSpriests],
};

export const StormstrikeFrequency = {
	id: 'stormstrike-frequency',
	type: 'number' as const,
	label: 'Stormstrike Cast Frequency',
	labelTooltip: 'How often (in Seconds) that Stormstrike is cast by an external Enhancement Shaman.',
	float: true,
	defaultValue: 20.0,
	inline: true,
	changedEvent: (player: Player<any>) => TypedEvent.onAny([player.changeEmitter, player.getRaid()!.debuffsChangeEmitter]),
	getValue: (player: Player<any>) => player.getStormstrikeFrequency(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setStormstrikeFrequency(eventID, newValue);
	},
	showWhen: (player: Player<any>) => player.getRaid()?.getDebuffs().stormstrike === true,
};

export const StormstrikeNatureAttackersFrequencey = {
	id: 'stormstrike-attackers-frequency',
	type: 'number' as const,
	label: 'Other Nature Attacks Frequency',
	labelTooltip: 'How often (in Seconds) that external attackers deal Nature damage to the target.',
	float: true,
	defaultValue: 4.0,
	inline: true,
	changedEvent: (player: Player<any>) => TypedEvent.onAny([player.changeEmitter, player.getRaid()!.debuffsChangeEmitter]),
	getValue: (player: Player<any>) => player.getStormstrikeNatureAttackerFrequency(),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		player.setStormstrikeNatureAttackerFrequency(eventID, newValue);
	},
	showWhen: (player: Player<any>) => player.getRaid()?.getDebuffs().stormstrike === true,
};

export const StormstrikeConfig = {
	tooltip: 'Stormstrike debuff configuration',
	inputs: [StormstrikeFrequency, StormstrikeNatureAttackersFrequencey],
};

export const TankAssignment = {
	id: 'tank-assignment',
	type: 'enum' as const,
	extraCssClasses: ['tank-selector', 'threat-metrics', 'within-raid-sim-hide'],
	label: 'Tank Assignment',
	labelTooltip:
		'Determines which mobs will be tanked. Most mobs default to targeting the Main Tank, but in preset multi-target encounters this is not always true.',
	values: [
		{ name: 'None', value: -1 },
		{ name: 'Main Tank', value: 0 },
		{ name: 'Tank 2', value: 1 },
		{ name: 'Tank 3', value: 2 },
		{ name: 'Tank 4', value: 3 },
	],
	changedEvent: (player: Player<any>) => player.getRaid()!.tanksChangeEmitter,
	getValue: (player: Player<any>) => (player.getRaid()?.getTanks() || []).findIndex(tank => UnitReference.equals(tank, player.makeUnitReference())),
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		const newTanks = [];
		if (newValue != -1) {
			for (let i = 0; i < newValue; i++) {
				newTanks.push(emptyUnitReference());
			}
			newTanks.push(player.makeUnitReference());
		}
		player.getRaid()!.setTanks(eventID, newTanks);
	},
};

export const IncomingHps = {
	id: 'incoming-hps',
	type: 'number' as const,
	label: 'Incoming HPS',
	labelTooltip: `
		<p>Average amount of healing received per second. Used for calculating chance of death.</p>
		<p class="mb-0>If set to 0, defaults to 17.5% of the primary target's base DPS.</p>
	`,
	changedEvent: (player: Player<any>) => player.getRaid()!.changeEmitter,
	getValue: (player: Player<any>) => player.getHealingModel().hps,
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		const healingModel = player.getHealingModel();
		healingModel.hps = newValue;
		player.setHealingModel(eventID, healingModel);
	},
	enableWhen: (player: Player<any>) => (player.getRaid()?.getTanks() || []).find(tank => UnitReference.equals(tank, player.makeUnitReference())) != null,
};

export const HealingCadence = {
	id: 'healing-cadence',
	type: 'number' as const,
	float: true,
	label: 'Healing Cadence',
	labelTooltip: `
		<p>How often the incoming heal 'ticks', in seconds. Generally, longer durations favor Effective Hit Points (EHP) for minimizing Chance of Death, while shorter durations favor avoidance.</p>
		<p>Example: if Incoming HPS is set to 1000 and this is set to 1s, then every 1s a heal will be received for 1000. If this is instead set to 2s, then every 2s a heal will be received for 2000.</p>
		<p class="mb-0">If set to 0, defaults to 1.5 times the primary target's base swing timer, and half that for dual wielding targets.</p>
	`,
	changedEvent: (player: Player<any>) => player.getRaid()!.changeEmitter,
	getValue: (player: Player<any>) => player.getHealingModel().cadenceSeconds,
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		const healingModel = player.getHealingModel();
		healingModel.cadenceSeconds = newValue;
		player.setHealingModel(eventID, healingModel);
	},
	enableWhen: (player: Player<any>) => (player.getRaid()?.getTanks() || []).find(tank => UnitReference.equals(tank, player.makeUnitReference())) != null,
};

export const HealingCadenceVariation = {
	id: 'healing-cadence-variation',
	type: 'number' as const,
	float: true,
	label: 'Cadence +/-',
	labelTooltip: `
		<p>Magnitude of random variation in healing intervals, in seconds.</p>
		<p>Example: if Healing Cadence is set to 1s with 0.5s variation, then the interval between successive heals will vary uniformly between 0.5 and 1.5s. If the variation is instead set to 2s, then 50% of healing intervals will fall between 0s and 1s, and the other 50% will fall between 1s and 3s.</p>
		<p class="mb-0">The amount of healing per 'tick' is automatically scaled up or down based on the randomized time since the last tick, so as to keep HPS constant.</p>
	`,
	changedEvent: (player: Player<any>) => player.getRaid()!.changeEmitter,
	getValue: (player: Player<any>) => player.getHealingModel().cadenceVariation,
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		const healingModel = player.getHealingModel();
		healingModel.cadenceVariation = newValue;
		player.setHealingModel(eventID, healingModel);
	},
	enableWhen: (player: Player<any>) => (player.getRaid()?.getTanks() || []).find(tank => UnitReference.equals(tank, player.makeUnitReference())) != null,
};

export const BurstWindow = {
	id: 'burst-window',
	type: 'number' as const,
	float: false,
	label: 'TMI Burst Window',
	labelTooltip: `
		<p>Size in whole seconds of the burst window for calculating TMI. It is important to use a consistent setting when comparing this metric.</p>
		<p>Default is 6 seconds. If set to 0, TMI calculations are disabled.</p>
	`,
	changedEvent: (player: Player<any>) => player.getRaid()!.changeEmitter,
	getValue: (player: Player<any>) => player.getHealingModel().burstWindow,
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		const healingModel = player.getHealingModel();
		healingModel.burstWindow = newValue;
		player.setHealingModel(eventID, healingModel);
	},
	enableWhen: (player: Player<any>) => (player.getRaid()?.getTanks() || []).find(tank => UnitReference.equals(tank, player.makeUnitReference())) != null,
};

export const HpPercentForDefensives = {
	id: 'hp-percent-for-defensives',
	type: 'number' as const,
	float: true,
	label: 'HP % for Defensive CDs',
	labelTooltip: `
		<p>% of Maximum Health, below which defensive cooldowns are allowed to be used.</p>
		<p class="mb-0">If set to 0, this restriction is disabled.</p>
	`,
	changedEvent: (player: Player<any>) => player.rotationChangeEmitter,
	getValue: (player: Player<any>) => player.getSimpleCooldowns().hpPercentForDefensives * 100,
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		const cooldowns = player.getSimpleCooldowns();
		cooldowns.hpPercentForDefensives = newValue / 100;
		player.setSimpleCooldowns(eventID, cooldowns);
	},
};

export const InspirationUptime = {
	id: 'inspiration-uptime',
	type: 'number' as const,
	float: true,
	label: 'Inspiration % Uptime',
	labelTooltip: `
		<p>% average of Encounter Duration, during which you have the Inspiration buff.</p>
		<p class="mb-0">If set to 0, the buff isn't applied.</p>
	`,
	changedEvent: (player: Player<any>) => player.healingModelChangeEmitter,
	getValue: (player: Player<any>) => player.getHealingModel().inspirationUptime * 100,
	setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
		const healingModel = player.getHealingModel();
		healingModel.inspirationUptime = newValue / 100;
		player.setHealingModel(eventID, healingModel);
	},
};
