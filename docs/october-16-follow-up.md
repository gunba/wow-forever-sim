# October 16 follow-up

This follow-up to the [October 15 update](october-15-update.md) stays on client 1.60.1.70425. It changes how weapon
imbues and relics are modelled, adds Rake's Attack Power scaling, reviews every class's rotation after the October 15
changes, and re-runs the full benchmark: a gear re-search of every DPS profile with two-seed confirmation, all 1,005
matrix scenarios and the 68 tank checks.

## Weapon imbues and temporary enchants

In the client, every Shaman imbue and Rogue poison rank applies its enchantment through its own enchant slot (effect
360), while sharpening stones, weightstones and oils use the temporary-enchant slot (effect 54). Rogue and Shaman
testers have reported poisons, stones and Windfury on the same weapon since September.

- Shaman imbues (Rockbiter, Flametongue, Frostbrand, Windfury) are now Elemental, Enhancement and Restoration spec
  options for each hand. The Consumes weapon slot holds a stone or oil, so both can be used at once. Classic keeps one
  shared temporary-enchant slot, and an imbue left in Consumes still works.
- Rogue poisons were already class options beside the Consumes stone or oil.
- Feral and Bear can select stones and oils in the web simulator (weapon enchants work in forms since the 70291 notes).
- Enhancement now runs Windfury Weapon with Shadow Oil on the main hand (+1.2% over an Elemental Sharpening Stone).
  Feral and Bear keep an Elemental Sharpening Stone. (CORE-023)

## Relics

The benchmark no longer equips full-capacity modeled relics. Every real Forever idol, libram and totem whose effect the
simulator implements has a projected copy: the real item's own effect plus a quarter of the modeled level-65 relic's
stats (physical: 2 Stamina, 1 Agility, 6.5 Attack Power; spell: 2 Stamina, 1.75 Intellect, 2.25 Spell Power). The gear
search picks the best one for each profile. (DATA-018)

New relic effects: Enraged Idol (Enrage +1 Rage a tick), Idol of the Ursine Twins (4% for Lacerate to reset Primal
Bite), Sentinel's Libram (Swift Judgement cooldown 50 s), Steadfast Libram (block value x1.15 while Holy Shield is
active) and Totem of the Storm (Lightning Bolt can trigger Maelstrom Weapon at half the main-hand per-hit chance; the
rate is provisional, SHA-011).

| Build | Relic |
|---|---|
| Feral | Idol of the Dream (+1.2% over Howling Idol) |
| Bear | Enraged Idol |
| Balance | Swarming Idol |
| Elemental, Stormcaller | Polished Driftwood Icon (Spirit regeneration while casting; chosen by the search) |
| Enhancement | Totem of the Storm (the totem choice is within 0.1%) |
| Retribution | Libram of Law; Libram of Invocation for two Physical Retribution races |
| Protection Paladin | Libram of Law (about 5 threat per second over Steadfast Libram) |

The smaller relic budget costs 0.2-0.8% where the relic choice did not change.

## Feral

- **Rake now scales with Attack Power.** The September 30 class deep dive says "Rake: Now gains increased damage from
  Attack Power". The client carries no coefficient (the term is server-side, like Swipe's), so it is fitted to public
  beta logs from October 2-10: about 6.5% of Attack Power on the hit and 4.5% on each of the three ticks, snapshotted
  on application. Keeping Rake up is then worth +6.1%; at two thirds of the fitted values it is still +3.1%, and without
  the term it loses 2.9%. Rake goes before Shifting Power in the rotation. (DRU-017, with a beta test)
- **Bear-weaving loses.** With Furor keeping Cat Energy through Bear Form, every weave variant (Enrage with Primal Bite
  or Lacerate, energy, cooldown and mana gates, with Mangle, Natural Shapeshifter, Enraged Idol and Idol of the Ursine
  Twins) lost 0.4-33%: each weave pays two shapeshifts that would otherwise fund Shifting Power, and Energy does not
  regenerate in Bear Form. (DRU-015, DRU-016)
- **Talents** with Rake: `05002-45010032021032212051-055` (Feral Cat 7/34/10), +0.7% over the October 15 screen's best.
- **Proc rates in forms.** The form attack is the equipped weapon's damage per second at the form's speed (Cat 1.0,
  Bear 2.5). Speed-scaled on-hit damage already used that speed, and beta logs confirm it: Flametongue Totem hits from a
  3.3-speed staff deal 18-19 in caster form and 5-6 in Cat Form, and a 2.0-speed Cat's 642 non-crit procs, including
  those from special attacks, are all 9-11. Proc-per-minute effects were inconsistent on special attacks: 15
  proc-trigger effects used the equipped weapon's speed (a Classic rule), while Crusader, Fiery Weapon and the weapon
  procs used the form speed. Forever now uses the form speed for all of them. No log shows a Cat PPM proc, so this is
  inferred (DRU-018, with a beta test). The ranked profiles carry no proc effects, so their results do not change.

Feral moves from 757 to 847 DPS (+12%).

## Rotation review after October 15

Each change below was screened on two seeds and confirmed on every race of the build at 5,000 iterations on two
further seeds.

| Build | Change | Confirmed |
|---|---|---|
| Frost | Cone of Cold at 5 Winter's Chill stacks (each +40%), standing at 10 yards; Winter's Chill 4/5, Arcane Subtlety 1/2 | +5.0% (MAG-011: positioning is a model choice) |
| Shadow | Mind Flay channelled in full (Improved Mind Flay makes the last tick strongest) | +0.4% |
| Smite | Shadow Word: Pain kept up after Holy Fire | +1.3% |
| Enhancement | Shadow Oil with Windfury Weapon | +1.2% |
| Feral | Rake, talents and Idol of the Dream (above) | +12% |

Screened and kept: Fire (Fire Blast and Pyroblast variants), Arcane and Arcane-Frost (published on October 15),
Marksmanship and Beast Mastery Multi-Shot (no gain), Retribution and Protection Paladin Consecration (removing it
loses 2.7% and 12%), Enhancement Fire Nova and Maelstrom Weapon at 5 stacks (4 stacks loses 2.4%), the Bear relic and
the Feral weapon stone.

## Results

The gear search then re-ran on every DPS profile (seed 1793710101); 28 loadout changes held on an independent
two-seed confirmation (Elemental and Stormcaller 5 each, Combat and pet Hunter 4, Fury with Sunder 3, Physical
Retribution 2, one each for Enhancement, Frost, Demonology, Destruction and Arcane-Frost). All 1,005 matrix scenarios
were replayed with no warnings.

| Build | October 15 | October 16 | Change |
|---|---|---|---|
| Feral | 756.5 | 847.1 | +12.0% |
| Frost | 607.8 | 638.7 | +5.1% |
| Enhancement | 801.0 | 814.3 | +1.7% |
| Smite | 444.1 | 449.9 | +1.3% |
| Bear (DPS) | 543.7 | 550.1 | +1.2% |
| Elemental | 525.4 | 529.8 | +0.9% |
| Stormcaller | 463.9 | 466.8 | +0.6% |
| Retribution | 1004.5 | 1001.5 | -0.3% |

Other builds moved by less than 0.2%. The 68 tank checks replayed on new seeds: Bear threat +1.4-1.6% (Enraged Idol),
Protection Paladin threat +0.1-0.2% (Libram of Law), Protection Warrior unchanged.

## Open questions

| Entry | Question |
|---|---|
| DRU-017 | Rake's Attack Power coefficients (fitted to logs) |
| DRU-018 | PPM rates on Cat special attacks |
| SHA-011 | Totem of the Storm's Maelstrom Weapon chance |
| DATA-018 | Stats on level-65 relics at launch |
| MAG-011 | Whether raid encounters allow 10-yard Frost positioning |

[Beta test procedures](beta-test-procedures.md) lists how to settle each one.

## Evidence

- [`comparison.json`](../artifacts/correctness/october_16/comparison.json): published, original-setup and selected
  DPS for all 201 profiles.
- [`selection-evidence.json`](../artifacts/correctness/october_16/selection-evidence.json): gear search and
  confirmation, talent searches, the confirmed selections and the Rake screens.
- The October 15 results this release supersedes are [archived](../artifacts/history/3c8abeb0c/manifest.json).
