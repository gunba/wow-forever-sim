# October 11 log audit

Baseline: published `e7f94795f` ([October 11 audit](october-11-audit.md)).

This round tests open register entries against public combat logs, community test reports and client
data, then corrects the engine where the evidence is clear. It also classifies every remaining open
entry by whether the current level-30 beta can settle it, and gives a test procedure for each one in
[Beta test procedures](beta-test-procedures.md).

Evidence:
- **Public combat logs.** The 97 public foreverlogs.gg reports with encounter data (5,863 encounters,
  7.2 million events, beta characters up to level 30). The public feed has event types, outcomes
  (hit, crit, miss, dodge, parry, block), spell IDs and timestamps. It has no attack power, spell
  power, talent ranks, hand or glancing flag, so it can settle timing and proc questions but not
  coefficients.
- **Community test reports.** Class Discord test channels (Hunter, Warlock, Fight Club, Druid) and
  the official October 1 beta notes.
- **Client data.** 1.60.1.70338 `SpellAuraOptions`, talent curves and enchant records.

Register dispositions are in [Simulator questions and coverage](uncertainties.md). Raw evidence is in
[`artifacts/correctness/october_11_logs/`](../artifacts/correctness/october_11_logs/). The previous
results are [archived](../artifacts/history/e7f94795f/manifest.json).

## Corrections

| Change | Evidence | Register |
|---|---|---|
| Shaman Flurry lasts for one white swing, not three. Windfury and Stormstrike attacks do not end it. | All 28 Shamans with Flurry windows lost the buff to the first white swing after its last application or refresh (1,151 of 1,171 windows). The client lists three charges. | SHA-009 |
| After its 10-second cooldown, Omen of Clarity procs from white hits at 4 per minute of attack speed and from abilities, spells and heals at 20%. | 691 Druid proc windows: the cooldown is exact. White hits procced 5.5-7.5% in Cat Form (1.0 s attacks) and 14.5-18.5% in Bear Form (2.5 s), the same 3.5-4.5 per minute; other events 18-23%. The client's 100% is not literal. | DRU-001 |
| Frostbolt and Frostfire Bolt roll Fingers of Frost when the cast completes (if not missed), before the bolt lands. | 121 of 121 Frostbolt-triggered procs at the cast-success timestamp; none at impact. | MAG-002 |
| Fireball, Frostbolt and Frostfire Bolt roll Missile Barrage at cast completion. Arcane Missiles spends it when the channel starts. | 11 of 11 Fireball/Frostbolt procs at cast success; removals at Arcane Missiles cast success. | MAG-005 |
| Redoubt also rolls on missed, dodged and parried attacks. | 10.0% of isolated landed and 9.8% of isolated avoided attacks on Protection Paladins (10% at 5/5). | PAL-011 |
| Shadow Bolt grants Decimation at cast completion, not when the bolt lands. | Warlock testers' combat-log traces. | WL-012 |
| The Imp waits about 0.43 s between Firebolt casts (was 0.2 s). | 2,443 cast cycles: 2.0 s cast, median 432 ms pause (10th-90th percentile 389-467 ms). | WL-006 |
| A queued Raptor Strike no longer removes the off-hand dual-wield miss penalty. | Hunter testers (September 28); the October 1 notes removed the same effect from Heroic Strike. | HUN-015 |
| Maelstrom Weapon procs 10 times per minute at every rank (5/5 unchanged). | A 1-point Shaman with a 3.6 s weapon gained a stack on 55% of landed melee hits; 2 procs per minute per point would give 12%. | SHA-001 |

## Confirmed without change

- **Windfury Weapon** (SHA-003): 22% per landed white swing (1,283 procs), two extra attacks, a 1.5 s
  cooldown and full melee outcomes on the extra attacks.
- **Unbridled Wrath** (WAR-007): 57-59% of eligible white swings for four Fury Warriors, matching 12% per point at 5/5.
- **Queued Heroic Strike** (WAR-005): the October 1 notes match the existing model; the queue does
  not change off-hand hit.
- **Seal of Command:** about 7 procs per minute for Paladins with two-handed weapons, matching the model.
- **Hack and Slash** (ROG-005): no two extra attacks closer than 824 ms in one encounter.
- **Summon Hawk** (HUN-004): 1,312 Hawks: two concurrent Hawks, at most 18 s of attacks, an opening dive about
  1 s after the cast that never missed, and assault attacks that use the full melee table.
- **Poison charges** (ROG-003): 175-180 charges per application cannot run out in five minutes.

## Register

- 7 entries resolved and 3 given working answers; 3 new resolved entries (SHA-009, WL-012, HUN-015).
- 28 entries that change no simulated DPS or tank result moved outside current scenarios (for
  example healing spells, crowd control, items no profile uses and multi-target-only actions).
- The remaining 49 open entries and 16 working answers that affect results each have a test
  procedure, grouped by class and by when they can be tested: 36 now, 14 partly now (low ranks),
  10 at launch, and 5 that need code or a modelling decision.

## Selections

Builds whose results changed were re-searched (52 DPS profiles, 3 coordinate passes, seed
1793700001) and every change was confirmed on two independent seeds (1793800001 and 1793900001). One
change was retained: Enhancement Orc (+0.7%). A second search on the final engine (seed 1794100001)
found no further changes. Tank gear and talents are unchanged.

## Benchmark

All 1,005 scenarios were replayed (seed 1793400001, 5,000 iterations) and all 68 tank workload
checks (seeds 1793140001 and 1793160001). Builds that changed by at least 0.05%, averaged over races:

| Build | Published | Corrected engine, old setup | Released | Change |
|---|---:|---:|---:|---:|
| Enhancement | 761 | 726 | 727 | -4.5% |
| Feral | 729 | 703 | 703 | -3.6% |
| Arcane–Frost | 732 | 743 | 743 | +1.4% |
| Frost | 683 | 690 | 690 | +0.9% |
| Bear | 545 | 540 | 540 | -0.8% |
| Protection Paladin | 600 | 603 | 603 | +0.5% |
| Arcane | 731 | 732 | 732 | +0.2% |
| Balance | 710 | 710 | 710 | -0.1% |

Tank checks: Protection Paladin single-target threat +0.5% (more Redoubt blocks feed Holy Shield);
Bear single-target threat -1.0% (fewer Omen procs).

The largest open question that still affects results is Warrior Flurry (WAR-006). The Warrior model
keeps three charges, but if Warrior Flurry lasts one swing like Shaman Flurry, Fury and Fury (Sunder)
lose about 7% and 2H Bloodthirst about 6%. Level 30 gives 26 talent points, so 1/5 Flurry can be
tested now. The next largest are Revenge's rank-6 AP coefficient (Protection Warrior), pet stat
inheritance and per-hit damage (Hunter and Warlock pets), and Rogue poison and finisher scaling. All
can be tested at least in part in the current beta.

Register tests that wrongly assumed these talents need launch were corrected: Flurry, Maelstrom
Weapon, Hack and Slash, Weaponmaster, Vindication and Demonic Knowledge are all reachable with the
26 points available at level 30.
