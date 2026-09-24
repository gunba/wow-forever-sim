import data from './forever_ranked_profiles.json';
import type { JsonValue } from '@protobuf-ts/runtime';
import { Spec } from './proto/common';
import { IndividualSimSettings } from './proto/ui';
import { playerToSpec } from './proto_utils/utils';

export interface RankedProfile {
	id: string;
	key: string;
	build: string;
	race: string;
	dps: number;
	settings: IndividualSimSettings;
	unmodeledSetBonuses: string[];
	caveats: string[];
	modeledGear: boolean;
}

let profiles: RankedProfile[] | undefined;

export function getRankedProfiles(spec: Spec): RankedProfile[] {
	if (!profiles) {
		profiles = (data.profiles as unknown as Array<Omit<RankedProfile, 'settings'> & { settings: JsonValue }>).map(profile => ({
			...profile,
			settings: IndividualSimSettings.fromJson(profile.settings),
		}));
	}
	return profiles
		.filter(profile => profile.settings.player && playerToSpec(profile.settings.player) === spec)
		.sort((a, b) => b.dps - a.dps || a.id.localeCompare(b.id));
}
