import { Player } from '../../player.js';
import { Ruleset } from '../../proto/api.js';
import { Class, ItemSlot, Spec, WeaponImbue } from '../../proto/common.js';
import { RogueOptions_Poison as Poison } from '../../proto/rogue.js';
import { ActionId } from '../../proto_utils/action_id.js';
import { isWeapon } from '../../proto_utils/utils.js';
import { TypedEvent } from '../../typed_event.js';
import * as InputHelpers from '../input_helpers.js';
import type { ConsumableInputConfig } from './consumables.js';

// Classic poisons occupy the same slot as temporary weapon enchants.
const classicRogue = (player: Player<any>) => player.getClass() === Class.ClassRogue && player.sim.getRuleset() !== Ruleset.RulesetForever;

export const ClassicInstantPoisonWeaponImbue: ConsumableInputConfig<WeaponImbue> = {
	actionId: () => ActionId.fromItemId(8928),
	value: WeaponImbue.InstantPoison,
	showWhen: classicRogue,
};

export const ClassicDeadlyPoisonWeaponImbue: ConsumableInputConfig<WeaponImbue> = {
	actionId: () => ActionId.fromItemId(20844),
	value: WeaponImbue.DeadlyPoison,
	showWhen: classicRogue,
};

export const ClassicWoundPoisonWeaponImbue: ConsumableInputConfig<WeaponImbue> = {
	actionId: () => ActionId.fromItemId(10922),
	value: WeaponImbue.WoundPoison,
	showWhen: classicRogue,
};

export function roguePoisonInput(slot: ItemSlot.ItemSlotMainHand | ItemSlot.ItemSlotOffHand) {
	const hasWeapon = (player: Player<Spec.SpecRogue>) => {
		const weapon = player.getEquippedItem(slot);
		return !!weapon && isWeapon(weapon.item.weaponType);
	};
	return InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecRogue, Poison>({
		fieldName: slot === ItemSlot.ItemSlotMainHand ? 'mainHandPoison' : 'offHandPoison',
		tooltip: slot === ItemSlot.ItemSlotMainHand ? 'Main-hand poison' : 'Off-hand poison',
		values: [
			{ value: Poison.NoPoison },
			{ value: Poison.InstantPoison, actionId: () => ActionId.fromItemId(8928), showWhen: hasWeapon },
			{ value: Poison.DeadlyPoison, actionId: () => ActionId.fromItemId(20844), showWhen: hasWeapon },
			{ value: Poison.WoundPoison, actionId: () => ActionId.fromItemId(10922), showWhen: hasWeapon },
		],
		changeEmitter: player => TypedEvent.onAny([player.specOptionsChangeEmitter, player.gearChangeEmitter, player.sim.rulesetChangeEmitter]),
		showWhen: player => player.sim.getRuleset() === Ruleset.RulesetForever,
	});
}

export function roguePoisonWarning(player: Player<Spec.SpecRogue>): string {
	if (player.sim.getRuleset() !== Ruleset.RulesetForever) return '';
	const consumes = player.getConsumes();
	const oldPoisons = [WeaponImbue.InstantPoison, WeaponImbue.DeadlyPoison, WeaponImbue.WoundPoison];
	if (oldPoisons.includes(consumes.mainHandImbue) || oldPoisons.includes(consumes.offHandImbue)) {
		return 'This profile stores poisons as temporary weapon enchants. Select them in the separate poison controls and clear the old weapon-imbue selections before simulating.';
	}
	const options = player.getSpecOptions();
	const hands: Array<[ItemSlot, Poison]> = [[ItemSlot.ItemSlotMainHand, options.mainHandPoison], [ItemSlot.ItemSlotOffHand, options.offHandPoison]];
	for (const [slot, poison] of hands) {
		const weapon = player.getEquippedItem(slot);
		if (poison !== Poison.NoPoison && (!weapon || !isWeapon(weapon.item.weaponType))) {
			return 'A selected poison requires a weapon in that hand.';
		}
	}
	return '';
}
