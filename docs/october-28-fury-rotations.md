# October 28 Fury rotations

Fury and Fury (Sunder) get new talents and rotations, found by a rotation search on the simulator. Their gear was then
searched again under the new rotations. Nothing else changes. All 904 matrix scenarios (seed 1793600001, 5,000 fights
each) are re-run; only the 20 Fury and Fury (Sunder) rows move. The previous results are archived under
[artifacts/history/afc5d384a](../artifacts/history/afc5d384a/manifest.json).

Evidence:

- [Selection](../artifacts/correctness/october_28_fury_rotations/selection.json): the search's validation, the
  confirmation runs, the Fury (Sunder) variants and their Sunder coverage.
- [Gear search](../artifacts/correctness/october_28_fury_rotations/gear_search.json.gz) and its
  [confirmation](../artifacts/correctness/october_28_fury_rotations/gear_validation.json.gz).
- [Comparison](../artifacts/correctness/october_28_fury_rotations/comparison.json) with the published results.

## Talents

Fury moves from 19/31/1 to **20/31/0** (`30305213102-25053005112010501`): the Improved Bloodrage point becomes
Spearing Strike. Spearing Strike costs 15 rage, has a 20-second cooldown and deals 40% normalized weapon damage; the
simulator allows it only in Battle Stance. The 31 Fury points, including Bloodthirst, are unchanged.

## Rotation

The Fury rotation keeps its October 21 shape and changes these rules:

- **Before the pull:** Bloodrage at 10 seconds instead of 1, and Death Wish at 1.5 seconds instead of 2.5. Bloodrage
  keeps the Warrior in combat while it lasts, so its rage does not decay; Bloodrage any earlier stays excluded
  (WAR-030).
- **Heroic Strike:** queued from 42 rage instead of 45, checked before the other abilities. It replaces the next swing
  and does not use the global cooldown.
- **Overpower:** the swap to Battle Stance during an Overpower window now needs 35 rage or less instead of 40.
- **Spearing Strike:** cast in Battle Stance during an Overpower window, swapping under the same 35-rage limit. It is
  cast about nine times a fight.
- **Rend:** previously refreshed only when the Warrior was already in Battle Stance. It now swaps to Battle Stance to
  refresh Rend at 30 rage or less outside execute, and while Rend is down it stays there up to 75 rage instead of
  returning to Berserker Stance.

Fury (Sunder) uses the same rules, with its two Sunder Armor upkeep rules first. These keep five stacks and wait for
rage when a stack is due. It also takes Fury's Hamstring timing: Bloodthirst and Whirlwind at least 1 s away and the
next main-hand swing at least 0.7 s away, instead of 1.5 s, 1.5 s and 1 s. Its Demonic Rune limit falls from 60 rage
to 40, as in Fury.

## How it was found and checked

Five agents per round searched the rotation on the published Windshaper profile for five rounds. Changes were scored
on 1,000-fight banks and kept only after an independent 5,000-fight holdout. A sixth round on the updated Undead gear
found nothing better. The search stopped when the remaining candidates could not beat the incumbent by 0.25%; that is
a plateau, not proof of the best rotation.

On 10,000 paired fights per race on the October 28 follow-up's gear, the result beat the published Fury preset by
+17.6 to +19.0 DPS. The bench gives the same result on two seeds that share no fights: +17.5 to +19.8 DPS on every race.

Fury (Sunder) shares Fury's talent preset, so it needs a matching rotation. Gains over the published Fury (Sunder)
rotation on every race, on the new talents:

| Fury (Sunder) rotation | Gain |
|---|---|
| Published Sunder rotation | -1.8 to +0.2 DPS |
| Published Sunder rotation plus the Spearing Strike dance | +4.5 to +6.3 DPS |
| New Fury rotation with Sunder first, no Rend dance | +7.8 to +10.5 DPS |
| **New Fury rotation with Sunder first (selected)** | **+27.6 to +30.4 DPS** |

The selected rotation keeps five Sunder stacks for 293 of 300 seconds, reaching them at 6.6 seconds, the same as the
published rotation.

## Gear

The new rotations changed the gear of 4 of 20 profiles. All four changes held on an independent seed (30296071).

| Profile | Change | Gain |
|---|---|---|
| Fury (Sunder) Troll | Fist off-hand | +5.8 DPS |
| Fury (Sunder) Dwarf | Waist | +3.7 DPS |
| Fury Dwarf | Feet | +3.1 DPS |
| Fury Troll | Waist | +2.9 DPS |

## Results

| Race | Fury | Change | Fury (Sunder) | Change |
|---|---|---|---|---|
| Windshaper / High Order | 1015.6 | +18.4 | 978.5 | +27.7 |
| Undead | 1015.6 | +18.4 | 983.1 | +28.3 |
| Dwarf | 1011.0 | +21.0 | 964.7 | +33.0 |
| Orc | 1007.4 | +18.1 | 972.4 | +29.9 |
| Gnome | 1007.3 | +17.9 | 970.7 | +29.5 |
| Troll | 1007.0 | +21.4 | 959.5 | +35.1 |
| Night Elf | 1006.6 | +20.3 | 971.6 | +31.0 |
| Human | 1004.9 | +18.2 | 976.8 | +29.8 |
| Tauren | 1004.3 | +18.1 | 966.4 | +29.0 |

Fury moves from ninth to seventh of 35 builds, ahead of Retribution (1008) and Fire (1000). Keeping Sunder Armor up now
costs a Fury Warrior about 32 DPS instead of 42.

Spearing Strike, Overpower and Rend all need a stance swap. The rage that Tactical Mastery keeps through the swap
follows the provisional Warrior rage model (WAR-028).

Spearing Strike deals three times its damage to the Dragonkin reference boss. The [October 28 Fury follow-up](october-28-fury-follow-up.md)
keeps this build for the benchmark and the Giants/Dragonkin presets and gives the site's general presets the same rules
without Spearing Strike.
