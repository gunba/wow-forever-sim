# October 28 Fury follow-up

This follow-up corrects the site's Fury presets after the [October 28 Fury rotations](october-28-fury-rotations.md). It
also adds racial rules for Fury Orc and Night Elf and fixes an APL validation gap. All 904 matrix scenarios (seed
1793600001, 5,000 fights each) are re-run; only Fury Orc and Fury Night Elf move. The previous results are archived
under [artifacts/history/8222aecb8](../artifacts/history/8222aecb8/manifest.json). The evidence is in
[selection.json](../artifacts/correctness/october_28_fury_follow_up/selection.json) and the
[comparison](../artifacts/correctness/october_28_fury_follow_up/comparison.json).

## Spearing Strike and the target

Spearing Strike deals 40% normalized weapon damage, and three times that to Giants and Dragonkin. The benchmark's
reference boss is a level-63 Dragonkin, so the October 28 Fury build gets the full bonus. Against a target without
the bonus, the same build gains only a quarter as much over the October 21 rules:

| Fury build | Dragonkin boss | Other targets |
|---|---|---|
| October 21 (19/31/1) | baseline | baseline |
| October 28 Spearing Strike build (20/31/0) | +18.6 DPS | +4.6 DPS |
| October 28 rules without Spearing Strike (19/31/1) | +9.0 DPS | +9.0 DPS |

The benchmark rows keep the Spearing Strike build, as Hunters keep Monster Slaying against the same boss. The
October 28 release also gave that build to the site's general **Fury** preset, which the site describes as
general-purpose. That is corrected:

- **Fury** and **Fury (Adaptive AoE)** go back to the 19/31/1 talents. The **Fury** rotation keeps the October 28
  rules except Spearing Strike. It beats the October 21 rotation by +8.4 to +9.6 DPS on every race, against any
  target.
- The **Fury (Sunder)** rotation does the same: the October 28 rules with Sunder upkeep first and without Spearing
  Strike, on 19/31/1. It beats the October 21 rotation by +20.4 to +21.6 DPS on every race.
- **Fury (Giants/Dragonkin)** now uses the searched Spearing Strike build (20/31/0 with the October 28 rotation).
  The new **Fury (Sunder, Giants/Dragonkin)** rotation is its Sunder version. The previous Giants/Dragonkin build
  (18/33/0) was about 103 DPS behind it, even against Dragonkin.
- Against targets without the bonus, the general presets beat the Spearing Strike builds by about 4.6 DPS (Fury) and
  5.3 DPS (Fury (Sunder)).

The ranked benchmark profiles open the Spearing Strike builds, as before.

## Racial rules

A further search tried racial timing on the published Fury rows, one race at a time. Two changes held:

- **Orc:** Blood Fury 0.1 seconds before the pull, then whenever rage is at least 50, instead of at the first chance.
  +0.68 DPS on block-paired fights (+0.70 in the search's 10,000-fight validation).
- **Night Elf:** Elune's Light from 15 seconds, then on cooldown. This avoids most of the overlap with the opening
  Recklessness. +2.2 DPS on block-paired fights (+2.7 in the search's validation).

Troll Berserking, Gnome Eureka and Undead Touch of the Grave timing showed no confirmed gain, even over 50,000 fights.
The same two rules cost Fury (Sunder) 1.7 DPS (Orc) and 0.7 DPS (Night Elf), so they apply to Fury only.

| Race | Fury before | Fury now |
|---|---|---|
| Night Elf | 1006.6 | 1009.4 |
| Orc | 1007.4 | 1008.1 |

Night Elf passes Orc, Gnome and Troll. Windshaper, High Order and Undead still lead at 1015.6.

## Prepull actions with a condition

The rotation editor offers no condition for a prepull action, but imported JSON could carry one. The simulator ignored
it without a warning, so a prepull action marked never to run still ran. Such an action now raises a warning and is
skipped, like a prepull action with an invalid time. No published rotation had one, and no result changes.
