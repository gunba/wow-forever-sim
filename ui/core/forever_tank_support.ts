import Support from '../../sim/core/forever_tank_support.json';
import { Class, Debuffs, IndividualBuffs, PartyBuffs, RaidBuffs } from './proto/common.js';

export function foreverTankSupport(playerClass: Class) {
	const warrior = playerClass === Class.ClassWarrior;
	return {
		raid: RaidBuffs.fromJson({
			...Support.raid,
			battleShout: warrior ? 'TristateEffectMissing' : Support.raid.battleShout,
		}),
		party: PartyBuffs.fromJson(Support.party),
		player: IndividualBuffs.fromJson(Support.player),
		debuffs: Debuffs.fromJson({
			...Support.debuffs,
			exposeArmor: warrior ? 'TristateEffectMissing' : Support.debuffs.exposeArmor,
		}),
	};
}
