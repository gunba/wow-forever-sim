import { Spec } from './proto/common.js';
import { getRankedProfiles } from './ranked_profiles';

// Fallback menu for pages without published ranking profiles.
export const communityBuilds: Record<Spec, string[]> = {
	[Spec.SpecBalanceDruid]: ['Moonkin 38/0/13'],
	[Spec.SpecFeralDruid]: ['Feral Cat 9/34/8'],
	[Spec.SpecFeralTankDruid]: ['Bear Tank 1/40/10'],
	[Spec.SpecRestorationDruid]: [],
	[Spec.SpecElementalShaman]: ['Elemental 31/7/13', 'Stormcaller 28/23/0'],
	[Spec.SpecEnhancementShaman]: ['Enhancement 20/31/0'],
	[Spec.SpecRestorationShaman]: [],
	[Spec.SpecHunter]: ['Beast Mastery 31/20/0', 'Marksmanship 5/35/11', 'Survival 7/11/33', 'Pet/Melee 16/10/25'],
	[Spec.SpecMage]: ['Fire 0/35/16', 'Frost 11/3/37', 'Arcane 33/3/15', 'Arcane–Frost 28/0/23'],
	[Spec.SpecRogue]: ['Combat Dual-Wield 15/33/3', 'Assassination Mutilate 31/20/0', 'Subtlety Hemo 20/0/31'],
	[Spec.SpecHolyPaladin]: [],
	[Spec.SpecProtectionPaladin]: ['Protection 0/43/8'],
	[Spec.SpecRetributionPaladin]: ['Retribution 13/7/31', 'Physical Ret 13/7/31'],
	[Spec.SpecHealingPriest]: [],
	[Spec.SpecShadowPriest]: ['Shadow 15/0/36'],
	[Spec.SpecSmitePriest]: ['Smite 31/17/3'],
	[Spec.SpecWarlock]: ['Demonic Pact 2/31/18', 'Deep Affliction 31/0/20', 'DS/Ruin Pandemic 21/11/19', 'Shadow and Flame 13/5/33'],
	[Spec.SpecWarrior]: ['Fury 17/34/0', 'Fury Spearing Strike 18/33/0', 'Arms 34/17/0', '2H Bloodthirst 20/31/0'],
	[Spec.SpecTankWarrior]: ['Protection 4/6/41'],
};

export function getCommunityBuilds(spec: Spec): string[] {
	const ranked = getRankedProfiles(spec);
	if (ranked.length) return [...new Set(ranked.map(profile => profile.build))];
	return communityBuilds[spec] || [];
}
