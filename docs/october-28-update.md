# October 28 update

On October 27 Blizzard added two launch changes to the
[October 26 racial notes](https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-october-26/2360696/8).
This update models them and re-runs all 872 matrix scenarios (seed 1793600001, 5,000 fights each):

- **Orc:** Blood Fury's cooldown is two minutes (was three). A five-minute fight now has three casts instead of two.
- **Undead:** Touch of the Grave's internal cooldown is 15 seconds (was 25).

Neither change is in a client build yet: 70532, the last beta build, still has three minutes and 25 seconds.

The same post clarifies that the Skyborne's one-hour Charged and Attuned buffs are removed on entering a raid. The
simulator uses neither Skysight nor Read Leyline, so nothing changes (CORE-010).

Talents, rotations, gear and consumables are as published. Every result for the other eight races is identical to the
October 27 update. Detailed numbers are in the
[comparison](../artifacts/correctness/october_28/comparison.json) and the
[matrix manifest](../artifacts/correctness/october_28/matrix-manifest.json). The previous results are archived under
[artifacts/history/571d71465](../artifacts/history/571d71465/manifest.json).

## Results

| Race | Builds | Mean change | Range |
|---|---|---|---|
| Orc | 25 | +0.27% | 0.00% to +0.48% |
| Undead | 23 | +0.69% | +0.44% to +1.32% |

- **Orc:** Warlocks gain most (about +0.45%); Arcane and Fury (Sunder) gain least (+0.14%). The Orc Protection Warrior
  is unchanged because its rotation does not cast Blood Fury.
- **Undead:** tanks and Rogues gain most (Protection Paladin +1.32%, Protection Warrior +1.25%, Rogues about +0.88%).
  The drain is 5% of maximum health, so it is worth more to characters with more health and less damage.

## Best race

Eleven builds have a new best race:

| Build | Best race | Was | Gain |
|---|---|---|---|
| Warrior Fury (Sunder) | Undead 952.2 | Windshaper 946.1 | +6.1 |
| Warlock DS/Ruin | Undead 868.1 | Human 864.5 | +3.7 |
| Rogue Mutilate | Undead 735.2 | Windshaper 732.1 | +3.1 |
| Priest Smite | Undead 503.6 | Gnome 500.5 | +3.1 |
| Hunter Survival | Orc 1180.2 | Night Elf 1178.6 | +1.6 |
| Mage Fire | Undead 994.5 | Human 993.0 | +1.5 |
| Shaman Stormcaller | Orc 638.3 | Dwarf 636.9 | +1.5 |
| Warlock Destruction | Undead 919.0 | Troll 917.9 | +1.1 |
| Mage Arcane–Frost | Undead 794.7 | Human 794.2 | +0.5 |
| Rogue Subtlety | Undead 718.5 | Windshaper 718.2 | +0.3 |
| Warlock Affliction | Orc 923.3 | Human 923.0 | +0.3 |

The last three leads are within about one standard error (0.3-0.5 DPS), so those races are effectively tied.

## Register

- **CORE-021** (Touch of the Grave) now gives the 15-second cooldown.
- **CORE-010** (Read Leyline) notes that Attuned is removed on entering a raid.
- The racial table in [forever_rules.md](forever_rules.md) lists the launch values.
