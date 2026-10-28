# October 28 follow-up

This follow-up adds a ranged Marksmanship Hunter and re-selects every DPS build's gear from the modeled level-65 item
pool on the current engine. All 904 matrix scenarios (226 race/build profiles, seed 1793600001, 5,000 fights each) are
re-run. Detailed numbers are in the [comparison](../artifacts/correctness/october_28_follow_up/comparison.json), the
[gear-search evidence](../artifacts/modelled_gear_search/current/summary.json) and the
[matrix manifest](../artifacts/correctness/october_28_follow_up/matrix-manifest.json). The previous results are archived
under [artifacts/history/92120c303](../artifacts/history/92120c303/manifest.json).

## Ranged Marksmanship

MM Weave is a hybrid: 0/27/24 talents, melee weaving, and a damage mix almost the same as SV Weave's. The new
**Marksmanship** build is the ranged version: at least 31 Marksmanship points with Trueshot Aura, Auto Shot from 12
yards, and no melee. MM Weave stays as it was.

- **Talents:** 7/37/7 (`502-0053552001523051-50002`): Trueshot Aura, Sniper Shot, Ranged Weapon Specialization,
  Barrage, Mortal Shots and Efficiency, with Deadly Aspects and Focused Fire for the pet.
- **Pet:** a Cat. It beat Lone Wolf without a pet by about 5 DPS, and Summon Hawk by about 45.
- **Rotation:**
  - Before the pull: Aspect of the Hawk, an Immolation Trap where the boss will stand, and an Aimed Shot that lands
    at the pull.
  - In the fight: Rapid Fire on cooldown, Aimed Shot, Serpent Sting, Arcane Shot and Sniper Shot. Multi-Shot shares
    Aimed Shot's cooldown, so it fires only when Aimed Shot cannot be cast.
  - Thorium Grenades fill the gaps, and the build drinks Major Mana Potions.
- **Gear:** searched separately for this build on every race.

The hunter deals no melee damage on any race; the Cat supplies about 15%.

| Race | DPS |
|---|---|
| Night Elf | 1042 |
| Windshaper / High Order | 1041 |
| Orc | 1039 |
| Human | 1039 |
| Troll | 1030 |
| Dwarf | 1024 |
| Tauren | 1022 |

That puts it sixth of 35 builds, behind SV Weave (1386), MM Weave (1328), Survival (1280), BM Weave (1202) and
Demonic Pact (1108).

The benchmark raid already has Trueshot Aura, so the talent adds no DPS here. In a raid where this Hunter is the only
source, it would raise everyone's ranged attack power.

The talents and rotation came from a coordinate search of the rotation's options alternating with the bench's talent
optimiser, under the build's rules.

## Gear

Each DPS race/build's equipment was searched again from the modeled pool. Talents, rotations, consumables, race and
enchants stayed fixed. Tanks keep their separately selected equipment.

The search changed 87 of 209 profiles. A change is kept only if it holds on a second, independent seed (30296071):
the mean gain over both seeds must exceed 1.96 times the summed standard errors. 86 changes passed. Combat Gnome
failed (+0.82 DPS against a 0.84 bound) and keeps its gear.

The search changed weapons without changing their stones. That left a Dense Weightstone on nine edged weapons
(Survival's new off-hand sword on every race, and Fury Dwarf's new main-hand sword). Those are now Dense
Sharpening Stones, which give the same +8 weapon damage, so DPS is unchanged. The search now changes the stone
with the weapon.

| Build | Races changed | Gain |
|---|---|---|
| Survival | 8 | +8.5% to +9.7% |
| Stormcaller | 5 | +1.4% to +4.6% |
| Retribution | 3 | +1.2% to +3.2% |
| Destruction | 5 | +1.2% to +2.4% |
| Pet/Melee | 8 | +0.9% to +1.9% |
| Demonic Pact | 5 | +0.6% to +1.3% |
| DS/Ruin | 5 | +0.6% to +1.0% |
| Fury | 5 | +0.3% to +1.0% |
| Fire | 5 | +0.4% to +1.0% |
| Shadow | 4 | +0.7% to +1.0% |
| Affliction | 5 | +0.3% to +1.0% |
| Fury (Sunder) | 8 | +0.4% to +0.9% |
| Physical Ret | 2 | +0.4% to +0.8% |
| Enhancement | 3 | +0.3% to +0.5% |
| Others | 1-6 each | under +0.4% |

Survival gains most because its October 27 rotation fires Auto Shot between melee attacks. It now carries two slow
2.9-second swords instead of a fast off-hand, and a 3.0-second crossbow instead of a 2.5-second one. MM Weave,
SV Weave, Arms, 2H Bloodthirst, Shockadin, Frostfire, Feral, Balance, Arcane–Frost, Smite and the tanks are unchanged.

## Best race

| Build | Best race | Was |
|---|---|---|
| Retribution | Undead 1008 | Human 990 |
| Destruction | Orc 938 | Undead 919 |
| Pet/Melee | Orc 935 | Windshaper 921 |
| Affliction | Undead 930 | Orc 923 |
| Physical Ret | Undead 940 | Human 935 |
| Fury | Undead 997 | Windshaper 995 |

## Other changes

- The October 11 gear-search record moves to `artifacts/modelled_gear_search/october_11`.
- The site's Marksmanship build preset, talents and rotation are the new build. MM Weave's gear preset is now labelled
  MM Weave.
