import { Spec } from './proto/common.js';
import { getRankedProfiles } from './ranked_profiles';

// Fallback menu for pages without published ranking profiles.
export const communityBuilds: Record<Spec, string[]> = {
	[Spec.SpecBalanceDruid]: ['Moonkin 38/0/13'],
	[Spec.SpecFeralDruid]: ['Feral Cat 9/35/7'],
	[Spec.SpecFeralTankDruid]: ['Bear Tank 0/31/20'],
	[Spec.SpecRestorationDruid]: [],
	[Spec.SpecElementalShaman]: ['Elemental 31/6/14', 'Stormcaller 28/23/0'],
	[Spec.SpecEnhancementShaman]: ['Enhancement 16/35/0'],
	[Spec.SpecRestorationShaman]: [],
	[Spec.SpecHunter]: ['Beast Mastery 35/16/0', 'Marksmanship 0/39/12', 'Survival 0/15/36'],
	[Spec.SpecMage]: ['Fire 0/35/16', 'Frost 14/0/37', 'Arcane 35/0/16'],
	[Spec.SpecRogue]: ['Combat Dual-Wield 15/33/3', 'Assassination Mutilate 31/20/0', 'Subtlety Hemo 15/0/36'],
	[Spec.SpecHolyPaladin]: [],
	[Spec.SpecProtectionPaladin]: ['Protection 0/45/6'],
	[Spec.SpecRetributionPaladin]: ['Retribution 10/0/41'],
	[Spec.SpecHealingPriest]: [],
	[Spec.SpecShadowPriest]: ['Shadow 15/0/36'],
	[Spec.SpecSmitePriest]: ['Smite 31/17/3'],
	[Spec.SpecWarlock]: ['Demonic Pact 2/31/18', 'Deep Affliction 35/0/16', 'DS/Ruin Pandemic 24/11/16', 'Shadow and Flame 13/11/27'],
	[Spec.SpecWarrior]: ['Fury 17/34/0', 'Arms 39/12/0'],
	[Spec.SpecTankWarrior]: ['Protection 1/0/50'],
};

export function getCommunityBuilds(spec: Spec): string[] {
	const ranked = getRankedProfiles(spec);
	if (ranked.length) return [...new Set(ranked.map(profile => profile.build))];
	return communityBuilds[spec] || [];
}
