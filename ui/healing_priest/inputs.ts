import * as InputHelpers from '../core/components/input_helpers.js';
import { Player } from '../core/player.js';
import { Spec, UnitReference, UnitReference_Type as UnitType } from '../core/proto/common.js';
import { ActionId } from '../core/proto_utils/action_id.js';
import { EventID } from '../core/typed_event.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const SelfPowerInfusion = InputHelpers.makeSpecOptionsBooleanIconInput<Spec.SpecHealingPriest>({
	fieldName: 'powerInfusionTarget',
	actionId: () => ActionId.fromSpellId(10060),
	extraCssClasses: [
		'within-raid-sim-hide',
	],
	getValue: (player: Player<Spec.SpecHealingPriest>) => {
		const target = player.getSpecOptions().powerInfusionTarget;
		return !target || target.type == UnitType.Self || (target.type == UnitType.Player && target.index == player.getRaidIndex());
	},
	setValue: (eventID: EventID, player: Player<Spec.SpecHealingPriest>, newValue: boolean) => {
		const newOptions = player.getSpecOptions();
		newOptions.powerInfusionTarget = UnitReference.create({
			type: newValue ? UnitType.Player : UnitType.Unknown,
			index: player.getRaidIndex(),
		});
		player.setSpecOptions(eventID, newOptions);
	},
});

export const InnerFire = InputHelpers.makeSpecOptionsBooleanIconInput<Spec.SpecHealingPriest>({
	fieldName: 'useInnerFire',
	actionId: () => ActionId.fromSpellId(48168),
});