# October 11 audit

Baseline: published `9c05fc890` (client 1.60.1.70291 integration). Evidence: client 1.60.1.70338 (the
latest build; it changes no simulated table after 70291), the 70338 hotfix cache, the official
class deep dives, ClassicWoWCommunity/forever-bugs and magey/forever-warrior, and class Discord test
channels. Register dispositions are in [Simulator questions and coverage](uncertainties.md).

## How the client was checked

- Every spell and aura registered by all 201 benchmark profiles was dumped from constructed
  simulations and compared with 70338 cast time, cooldown, GCD, cost and duration after applying the
  client's own passive talent and Tier 1 spell modifiers (aura 107/108 by class mask, trait curves).
  Remaining mismatches: Shield Wall, Intimidation and Stoneform GCDs; Sweeping Strikes duration; the
  Hunter Tier 1 Multi-Shot mask.
- Spell attributes were decoded with TrinityCore's names and diffed between Era 1.15.9 and Forever.
  This exposed Slam's swing reset, finisher Energy refunds, per-spell periodic crit, Arcane Shot and
  Serpent Sting spell-power removal and the absence of channel haste.
- The 70338 `DBCache.bin` hotfixes touch only NPC and quest spells plus a byte-identical Cold Blood
  row.
- PlayerExpectedStat confirms every class's level-60 crit per Agility, spell crit per Intellect, base
  mana and health per Stamina.
- Talent trees and curves are unchanged since 70291 and match the vendored trees.
- A race-detector run of all 31 builds in concurrent mode found no data races.

## Corrections

| Change | Evidence | Register |
|---|---|---|
| Untalented Slam blocks autos during its cast and restarts both swings; Improved Slam does not, and removes 3 s of cooldown at either rank | Base ranks: InterruptFlags 15, no `DO_NOT_RESET_COMBAT_TIMERS`; Improved Slam overrides with 1310196–1310200 (flags 7, attribute set); deep dive and talent text | WAR-019 |
| Forever has no hit suppression | 7.04% misses at 1% hit against +3 targets (95% CI 6.31–7.85%); magey/forever-warrior#1 | CORE-001 |
| Forever finishers refund 80% Energy when avoided | `DISCOUNT_POWER_ON_MISS` on every player rank of Eviscerate, Rupture, Expose Armor, Rip and Ferocious Bite (the earlier review read NPC ids) | CORE-003 |
| Siphon Life cannot crit | No `PERIODIC_CAN_CRIT`; forever-bugs#303 (219 ticks, no crits) | WL-011 |
| Hunter Tier 1 five-piece reduces Aimed Shot only | Bonus 1301252 mask 131072 | HUN-013 |
| Shield Wall and Intimidation trigger the GCD; Sweeping Strikes lasts 20 s | Client StartRecoveryTime and duration (Era and Forever) | CORE-018 |
| Arcane Shot +10% RAP, Serpent Sting +3% current RAP per tick, no spell power | Client coefficients zeroed (Era 0.429 / 0.2); Hunter Lodge measurements | HUN-005 |
| Rend +2% current AP per tick; Thunder Clap +7% AP and threat equal to damage | Deep dive (Priest and Warrior); Fight Club measurements | WAR-020 |
| Sunder threat + 0.05 AP | Deep dive; published threat sheet | WAR-009 |
| Revenge + 0.28 AP, threat 2.25× damage without a flat bonus | Six-point fit in the published threat sheet (rank 2); 0.20 at rank 1 | WAR-021 |

## Paired benchmark impact

Seed 20261011, 5,000 iterations, every race, unchanged gear, talents and rotations:

| Build | Change |
|---|---|
| Protection Warrior | +34% DPS (Revenge 233 → 976 per cast; sensitive to WAR-021) |
| Fury / Fury (Sunder) | −6.4% / −6.1% with the published rotations, −2.2% / −2.1% without Slam |
| 2H Bloodthirst | −4.6%; −3.9% with Slam gated at 50 Rage |
| Arms | −2.9%; Slam gated at 50 Rage recovers about 2% |
| Marksmanship / Beast Mastery | +2.9% / +1.4% |
| Affliction | −1.2% |
| Mutilate / Subtlety / Combat / Feral | +1.1% / +1.0% / +0.6% / +0.4% |
| Others | unchanged |

Hit-capped profiles were selected for the old 9% cap and are now 1% over the true 8% cap. Their gain
requires gear re-selection. Hunter flask and elixir choices that add spell power no longer add damage.
The published rankings, profiles and Warrior rotations have not been regenerated.
