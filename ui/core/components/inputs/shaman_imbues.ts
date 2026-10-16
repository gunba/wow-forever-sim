import { Player } from '../../player.js';
import { Ruleset } from '../../proto/api.js';
import { Class, ItemSlot, Spec, WeaponImbue } from '../../proto/common.js';
import { ShamanWeaponImbue } from '../../proto/shaman.js';
import { ActionId } from '../../proto_utils/action_id';
import { isWeapon } from '../../proto_utils/utils';
import { TypedEvent } from '../../typed_event.js';
import * as InputHelpers from '../input_helpers.js';
import { ConsumableInputConfig } from './consumables';

// Classic Shaman imbues are temporary weapon enchants and share the slot with stones and oils. Forever gives
// them their own slot (client effect 360), so they are spec options and coexist with the Consumes enchant.
const classicShaman = (player: Player<any>) => player.isClass(Class.ClassShaman) && player.sim.getRuleset() !== Ruleset.RulesetForever;

// Shaman Imbues
export const RockbiterWeaponImbue = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromSpellId(16316),
		value: WeaponImbue.RockbiterWeapon,
		showWhen: player => {
			if (!classicShaman(player)) return false;

			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

export const FlametongueWeaponImbue = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromSpellId(16342),
		value: WeaponImbue.FlametongueWeapon,
		showWhen: player => {
			if (!classicShaman(player)) return false;
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

export const FrostbrandWeaponImbue = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromSpellId(16356),
		value: WeaponImbue.FrostbrandWeapon,
		showWhen: player => {
			if (!classicShaman(player)) return false;
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

export const WindfuryWeaponImbue = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromSpellId(16362),
		value: WeaponImbue.WindfuryWeapon,
		showWhen: player => {
			if (!classicShaman(player)) return false;
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

type ShamanSpec = Spec.SpecElementalShaman | Spec.SpecEnhancementShaman;

export function shamanImbueInput<SpecType extends ShamanSpec>(slot: ItemSlot.ItemSlotMainHand | ItemSlot.ItemSlotOffHand) {
	const hasWeapon = (player: Player<SpecType>) => {
		const weapon = player.getEquippedItem(slot);
		return !!weapon && isWeapon(weapon.item.weaponType);
	};
	return InputHelpers.makeSpecOptionsEnumIconInput<SpecType, ShamanWeaponImbue>({
		fieldName: (slot === ItemSlot.ItemSlotMainHand ? 'mainHandImbue' : 'offHandImbue') as any,
		tooltip: slot === ItemSlot.ItemSlotMainHand ? 'Main-hand Shaman imbue' : 'Off-hand Shaman imbue',
		values: [
			{ value: ShamanWeaponImbue.NoShamanImbue },
			{ value: ShamanWeaponImbue.WindfuryImbue, actionId: () => ActionId.fromSpellId(16362), showWhen: hasWeapon },
			{ value: ShamanWeaponImbue.FlametongueImbue, actionId: () => ActionId.fromSpellId(16342), showWhen: hasWeapon },
			{ value: ShamanWeaponImbue.RockbiterImbue, actionId: () => ActionId.fromSpellId(16316), showWhen: hasWeapon },
			{ value: ShamanWeaponImbue.FrostbrandImbue, actionId: () => ActionId.fromSpellId(16356), showWhen: hasWeapon },
		],
		changeEmitter: player => TypedEvent.onAny([player.specOptionsChangeEmitter, player.gearChangeEmitter, player.sim.rulesetChangeEmitter]),
		showWhen: player => player.sim.getRuleset() === Ruleset.RulesetForever && (slot === ItemSlot.ItemSlotMainHand || hasWeapon(player)),
	});
}

export function shamanImbueWarning<SpecType extends ShamanSpec>(player: Player<SpecType>): string {
	if (player.sim.getRuleset() !== Ruleset.RulesetForever) return '';
	const consumes = player.getConsumes();
	const oldImbues = [WeaponImbue.RockbiterWeapon, WeaponImbue.FlametongueWeapon, WeaponImbue.FrostbrandWeapon, WeaponImbue.WindfuryWeapon];
	if (oldImbues.includes(consumes.mainHandImbue) || oldImbues.includes(consumes.offHandImbue)) {
		return 'This profile stores a Shaman imbue as a temporary weapon enchant. Forever imbues have their own slot: select the imbue in the Shaman imbue controls and use the weapon-enchant slot for a stone or oil.';
	}
	return '';
}
