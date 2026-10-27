# October 27 update

This update re-runs all 872 matrix scenarios (seed 1793600001, 5,000 fights each) on the current engine. It brings in:

- the launch racial values;
- the engine fixes from the October 26 APL search sweep;
- Jordan's answers on spell ranks and pre-trapping;
- Mage-only imbue scrolls, pet hit at 17:8, explosives during the global cooldown and the logged Summon Hawk;
- new talents and rotations for 25 DPS builds.

The review matrix no longer has a Learned column. Detailed numbers are in the
[comparison](../artifacts/correctness/october_27/comparison.json), the
[selection evidence](../artifacts/correctness/october_27/selection-evidence.json) and the
[matrix manifest](../artifacts/correctness/october_27/matrix-manifest.json). The previous results are archived under
[artifacts/history/e0fb5a249](../artifacts/history/e0fb5a249/manifest.json).

## Engine

- **Racials at launch values** (Blizzard's October 26 tuning pass):
  - Human: 1% crit and 10% Spirit.
  - Dwarf: 1% melee and spell hit, instead of 2% crit damage.
  - Tauren: health only.
  - Orc: Command is +5% damage for Hunter and Warlock pets (not guardians); the 1% crit is gone.
  - Skyborne: 2% haste.

  These are not yet in a client build. Builds lose 0.6-1.35%, except on Gnome.
- **Imbue scrolls are Mage-only.** Client data gives the Scrolls of Imbue (277502, 277494) AllowableClass 128. Other classes lose
  1% crit per weapon: about 0.8% DPS with one weapon and 1.6% dual-wielding (CORE-031).
- **Pet hit at 17:8.** Hunter pets take the Hunter's ranged hit as melee hit and 17/8 of it as spell hit. Demons take the
  Warlock's spell hit and 8/17 of it as melee hit, so the Succubus's melee hit falls from about 16% to 7.5% (HUN-020).
- **Explosives during the global cooldown.** Autocast throws an explosive while the global cooldown runs if what is left
  of it covers the cast: Fire +1.05%, 2H Bloodthirst +0.56%, Frostfire +0.42% (CORE-032).
- **Summon Hawk from combat logs.** The hawk gets:
  - 80% of the Hunter's ranged attack power;
  - half the Hunter's crit, plus Ferocity;
  - the Hunter's hit, expertise and damage multipliers.

  Beast Mastery hawks rise from 26 to 41 DPS and Pet/Melee hawks from 28 to 34 (HUN-004).
- **Damage-over-time ranks replace each other.** Applying one rank of a caster damage-over-time spell removes any other
  rank of the same spell from the target. Channels are not affected (CORE-033).
- **Pre-trapping.** A trap laid at -60 s or later goes off at 0 at any distance; its cooldown starts when it is laid
  (HUN-021).
- **Arcane Wrath.** Every rank can be cast, and all ranks share one 6-second cooldown (MAG-014).
- **Pre-pull contact.** Anything that reaches the enemy before 0 (a hit, a miss, a tick, or an aura such as a curse) now
  warns on the pre-pull action responsible, because in the game it starts the fight. Hunter's Mark, Battle Shout's threat
  and fire totem placement are exempt. Damage before 0 was never counted (SCEN-017).
- **APL fixes from the sweep:**
  - A condition on an aura the character cannot have is never active, and one on an unknown spell can never cast.
    Other conditions that cannot be evaluated stop their action with a warning; they used to vanish, so the action ran
    unconditionally.
  - Autocast Other Cooldowns is refused before the pull.
  - Warlock current-pet mana values read the active demon.
  - Rogues can Stealth before the pull.
  - Pre-pull haste no longer brings auto attacks before 0.
  - Multiplying or dividing a duration by a fraction keeps the fraction: dividing by 1.4 used to divide by 1, and
    dividing by 0.5 crashed.

  [Findings](https://github.com/gunba/wow-forever-sim/blob/forever/docs/apl-search-findings.md) records each with its reproduction.
- **Expected damage** averages the partial resist without rolling it; this only moves random streams.
- **Bear Form** keeps the health fraction it had before the shift. This was deployed on 24 October, and these are the
  first matrix results that include it.

## New presets

The talents and rotations come from the ML agent's re-search on the racial-pass engine:

- five independent agents per build, keeping the build's gear and theme;
- 1,000 paired fights per candidate;
- a fresh 5,000-fight paired confirmation of each round's finalists.

Every final candidate was then validated on 5,000 fights on this engine, with no warnings and the same per-fight results
on the previous engine. A debug sample of 64 fights per build and static haste probes found nothing reaching the enemy
before 0. The search found improvements where it stopped; it does not establish a global optimum.

Each candidate replaced its build's talents and rotation for every race of the build. Race, gear, consumables, raid and
encounter stay as published. On the current engine each new preset beats the previous one on every race. Figures are
from 2,000 fights per race, seed 1793600001:

| Build | Talents | Mean gain (DPS) | Per race | Mean |
|---|---|---|---|---|
| Mage Arcane | 33/3/15 → 32/3/16 | +158 | +149 to +165 | +23.0% |
| Shaman Stormcaller | 29/22/0 → 29/22/0 | +116 | +113 to +119 | +22.4% |
| Hunter Survival | 7/11/33 → 7/12/32 | +211 | +204 to +217 | +22.0% |
| Mage Frostfire | 0/29/22 → 0/29/22 | +134 | +126 to +138 | +16.9% |
| Hunter MM Weave | 0/28/23 → 0/27/24 | +167 | +158 to +178 | +14.5% |
| Hunter SV Weave | 0/20/31 → 0/20/31 | +153 | +143 to +163 | +12.6% |
| Warlock Demonic Pact | 2/31/18 → 2/31/18 | +97 | +95 to +100 | +9.8% |
| Warrior Arms | 35/16/0 → 36/13/2 | +79 | +76 to +82 | +9.8% |
| Mage Fire | 18/31/2 → 18/31/2 | +65 | +55 to +85 | +7.0% |
| Paladin Physical Ret | 12/6/33 → 16/1/34 | +54 | +52 to +55 | +6.1% |
| Warlock Destruction | 13/5/33 → 20/0/31 | +45 | +41 to +47 | +5.1% |
| Shaman Elemental | 31/11/9 → 31/11/9 | +29 | +26 to +34 | +4.7% |
| Paladin Retribution | 13/6/32 → 16/2/33 | +38 | +24 to +45 | +4.0% |
| Warrior 2H Bloodthirst | 20/31/0 → 20/31/0 | +34 | +31 to +35 | +3.9% |
| Hunter BM Weave | 31/20/0 → 31/20/0 | +38 | +37 to +40 | +3.3% |
| Rogue Subtlety | 20/0/31 → 20/0/31 | +19 | +16 to +20 | +2.7% |
| Druid Feral | 7/34/10 → 7/34/10 | +22 | +21 to +23 | +2.7% |
| Warlock Affliction | 31/0/20 → 30/0/21 | +24 | +22 to +25 | +2.7% |
| Shaman Enhancement | 20/31/0 → 19/31/1 | +21 | +17 to +23 | +2.7% |
| Rogue Mutilate | 31/20/0 → 30/21/0 | +13 | +12 to +15 | +1.9% |
| Druid Balance | 38/0/13 → 38/0/13 | +12 | +11 to +12 | +1.5% |
| Rogue Combat | 18/33/0 → 18/33/0 | +9 | +8 to +10 | +1.2% |
| Warrior Fury | 19/31/1 → 19/31/1 | +11 | +9 to +13 | +1.2% |
| Priest Shadow | 18/0/33 → 18/0/33 | +9 | +2 to +12 | +1.0% |
| Paladin Shockadin | 23/0/28 → 23/0/28 | +2 | +1 to +3 | +0.3% |

### Changes made on integration

- **Lower ranks.** Presets may now cast lower spell ranks, as long as each rank's coefficient is the client's. Six of
  the new presets do, mostly to save mana, and every rank they use matches client build 1.60.1.70450:
  - Demonology and Destruction: rank-2 Shadow Bolt (0.629), Demonology whenever mana is at or below 80%.
  - Elemental: Lightning Bolt rank 7 (0.714). Stormcaller: rank 7, with rank 2 (0.571) as the filler when mana runs short.
  - Frostfire: the lowest Scorch to keep Fire Vulnerability up, and the lowest Fire Blast (both 0.429).
  - Arcane: the lowest Arcane Wrath, which scales with mana rather than a coefficient.

  The game has no other downranking penalty in this model. If one exists (DATA-001), these are the presets it would
  move. In this simulator the lower ranks are worth about 270 DPS to Demonology, 80 to Stormcaller, 55 to Frostfire,
  10 to Elemental, 8 to Destruction and 2 to Arcane.
- **Retribution and Physical Ret** left 2 and 1 talent points unspent. They go into Toughness, which adds armor only;
  DPS is identical on the validation fights.
- **Feral** keeps its published main hand. The candidate's had the same stats with a 3.0 s swing instead of 3.8 s, and
  weapons are not changed.
- **Fire** no longer has a separate High Order talent variant, because the new talents are legal for every race.
- **Bench build identity** follows the accepted builds:
  - Affliction and Mutilate need 30 points in their main tree, down from 31.
  - Balance no longer requires Moonkin Form; its talents take Genesis instead.
  - Destruction no longer requires Incinerate; it now casts Shadow Bolt, Conflagrate and Shadowburn.
- **Physical Ret** gets its own rotation preset; it used to share Retribution's.
- **Openers.** Balance, Frostfire, Marksmanship and Shadow time their pre-pull cast from the spell's own cast and travel
  time:
  - Balance and Frostfire land at +0.05 s.
  - Marksmanship accounts for Rapid Fire (5/7 of the cast time).
  - Shadow lands exactly at 0, with no margin for latency.

  The timing is set when the rotation loads. A haste effect used before the pull, or different gear, needs the opener
  rechecked. The rotation editor's Do At field now takes any value, not only a constant, so these openers load and can
  be edited there.
- **Bloodrage.** Bloodrage before -10 s stays excluded, because out-of-combat rage decay is unknown (WAR-030). The
  Warrior openers use it at -5 s (Arms), -1 s (Fury) and -10 s (2H Bloodthirst).
- **Weaving openers.** BM, MM and SV Weave used to cast Aimed Shot at -2 s, which reached the enemy at -0.03 s for
  Windshaper and High Order. Their new openers land after 0 on every race.

### Kept at their published presets

The rule: a changed build is kept if it stays broadly similar to its theme and is not a clone of another build. A build
that changes theme is kept only if it becomes its class's best.

- **Smite.** The candidate is a deep-Shadow priest: Smite is 10% of its damage, and Mind Flay and Mind Blast are most
  of it. That is close to the Shadow build and below it (851 against 910 DPS).
- **Frost.** The candidate drops Winter's Chill and deep Frost for Arcane Blast, Missile Barrage and Arcane Mind, and
  Arcane Missiles becomes its largest source, as in Arcane–Frost. It is not the best Mage build (790 against Fire's 993).
- **Arcane–Frost.** The candidate takes Arcane's talents (Arcane Wrath, Presence of Mind, Arcane Power), drops Shatter
  and Fingers of Frost, and matches Arcane's damage mix (879 against Fire's 993).
- **Pet/Melee.** The candidate drops every Beast Mastery point, Summon Hawk included, and ranged-weaves with a Cat like
  the new Survival (1153 against SV Weave's 1386).
- **DS/Ruin.** The candidate keeps its demon with Demonic Pact and Master Demonologist instead of sacrificing it. That
  makes it a copy of Demonology, the best Warlock build (1063 against 1092).
- **Fury (Sunder).** Sunder Armor lapses, once as early as 162 s, and this build supplies Sunder for the raid
  (978 against Fury's 994).

### Close pairs

- **Survival and SV Weave.** Survival now weaves Auto Shot with its melee, as SV Weave does. It keeps its Cat (SV Weave
  has no pet in this scenario) and different talents (7/12/32 against 0/20/31).
- **Retribution and Physical Ret.** They differ in gear (spell power against Strength), one Retribution talent point
  and their priorities. Judgement of Command averages about 10 casts a fight for Retribution against 35 for Physical Ret.

## Results

Best race per build, five-minute benchmark (seed 1793600001). "Was" is the published October 21 result, from before this
update's engine changes. "Kept" builds change only through the engine.

| Build | Best DPS | Was | Mean change | Preset |
|---|---|---|---|---|
| Hunter SV Weave | 1386 (Windshaper) | 1258 | +10.1% | new |
| Hunter MM Weave | 1328 (Windshaper) | 1182 | +12.2% | new |
| Hunter BM Weave | 1197 (Orc) | 1166 | +2.7% | new |
| Hunter Survival | 1179 (Night Elf) | 985 | +19.3% | new |
| Warlock Demonic Pact | 1091 (Undead) | 1006 | +8.7% | new |
| Warrior Fury | 995 (Windshaper) | 1008 | -1.3% | new |
| Mage Fire | 993 (Human) | 930 | +7.6% | new |
| Paladin Retribution | 990 (Human) | 958 | +2.7% | new |
| Warrior Fury (Sunder) | 946 (Windshaper) | 972 | -2.5% | kept |
| Mage Frostfire | 938 (Human) | 803 | +16.9% | new |
| Paladin Physical Ret | 935 (Human) | 890 | +4.9% | new |
| Warlock Affliction | 923 (Human) | 915 | +1.5% | new |
| Hunter Pet/Melee | 921 (Windshaper) | 932 | -1.1% | kept |
| Warlock Destruction | 918 (Troll) | 885 | +4.1% | new |
| Priest Shadow | 911 (Human) | 913 | -0.0% | new |
| Warrior Arms | 903 (Undead) | 837 | +8.3% | new |
| Warrior 2H Bloodthirst | 891 (Undead) | 870 | +2.5% | new |
| Paladin Shockadin | 874 (Undead) | 880 | -0.9% | new |
| Warlock DS/Ruin | 864 (Human) | 879 | -1.3% | kept |
| Mage Arcane | 852 (Gnome) | 698 | +22.7% | new |
| Druid Feral | 841 (Windshaper) | 832 | +1.2% | new |
| Shaman Enhancement | 828 (Windshaper) | 813 | +1.5% | new |
| Druid Balance | 794 (Windshaper) | 796 | -0.2% | new |
| Mage Arcane–Frost | 794 (Human) | 797 | -0.3% | kept |
| Rogue Combat | 747 (Windshaper) | 754 | -0.4% | new |
| Rogue Mutilate | 732 (Windshaper) | 736 | -0.1% | new |
| Rogue Subtlety | 718 (Windshaper) | 718 | +0.6% | new |
| Shaman Elemental | 658 (Windshaper) | 636 | +3.3% | new |
| Shaman Stormcaller | 637 (Dwarf) | 531 | +20.9% | new |
| Mage Frost | 610 (Gnome) | 610 | -0.3% | kept |
| Druid Bear | 599 (Windshaper) | 602 | -0.5% | kept |
| Paladin Protection Paladin | 561 (Undead) | 559 | -0.0% | kept |
| Warrior Protection Warrior | 542 (Undead) | 542 | -0.4% | kept |
| Priest Smite | 500 (Gnome) | 503 | -0.5% | kept |

## Register

- **New entries:** CORE-031 to CORE-033, HUN-020 to HUN-023, MAG-014, SCEN-017 and WAR-030.
- **Open questions added:** WAR-030 (rage decay before the pull) and two Hunter questions from the upstream review:
  - HUN-022: does haste shorten Auto Shot's wind-up?
  - HUN-023: does the off hand wait half a swing after Attack is pressed? This is worth about 136 DPS to SV Weave.
- **DATA-001** now records which presets use lower ranks and how much they gain from them. The test that kept presets
  at their highest ranks is removed.
- **HUN-004** describes the log-based hawk.
