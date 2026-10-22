# October 22 learned rotations

The review matrix has a new first column, **Learned**. For every DPS build at its highest-simming race it shows the best
a reinforcement-learning agent reached: the agent's own network, or a simmable profile converted from it where that
scores higher. The published profiles, rotations and matrix numbers do not change: all 218 published cells replay
exactly with this release's code (seed 1793600001). Tanks are not included. The method is in
[Learned rotations and talents](learned-rotations.md).

- **The network** chose every action in the engine and searched the talents at the same time. Its DPS shows for every
  build, with its talents, what it did differently from the APL and a first-minute timeline on the build's page.
- **The simmable profile** is the published setup for that race with the chosen talents and the published APL edited by
  rules converted from the network. Each edit was kept only if it raised DPS on paired fights in the simulator. The
  profiles are benched like the other cells, appear as "<Build> (learned)" in each simulator's Ranked builds selector,
  and "sim" in a Learned cell opens them. Where no conversion kept a gain the cell says "no APL"; the build's page
  then lists the network's talents and actions to investigate.

## Results

Best race per build, five-minute benchmark. "Published" is the matrix cell. The network's gain is measured against the
published setup on the same 2,048 paired fights; the profile is benched like the matrix (5,000 iterations, seed
1793600001) and its gain is against the published cell with the same code. "Share kept" is the profile's paired gain
as a share of the network's on the same fights; "all" means it matched or passed the network.

| Build | Race | Published | Network (gain) | Simmable profile (gain) | Share kept | Talents | APL edits |
|---|---|---|---|---|---|---|---|
| Mage Arcane | Gnome | 698 | 809 (+112.0) | 777 (+79.1) | 71% | learned | 11 added, 0 removed |
| Hunter MM Weave | Human | 1182 | 1353 (+170.9) | 1219 (+37.6) | 22% | published | 3 added, 1 removed |
| Hunter SV Weave | Human | 1258 | 1439 (+179.9) | 1304 (+46.0) | 25% | learned | 3 added, 1 removed |
| Mage Frost | Gnome | 610 | 677 (+68.5) | 689 (+79.6) | all | learned | 5 added, 2 removed |
| Warrior Arms | Human | 837 | 916 (+79.0) | 870 (+33.1) | 43% | published | 8 added, 2 removed |
| Mage Frostfire | Human | 803 | 865 (+61.4) | 893 (+90.4) | all | learned | 3 added, 1 removed |
| Shaman Stormcaller | Tauren | 531 | 562 (+32.8) | 562 (+31.8) | 94% | learned | 4 added, 2 removed |
| Warrior 2H Bloodthirst | Orc | 870 | 915 (+46.4) | 883 (+13.4) | 29% | learned | 3 added, 1 removed |
| Mage Arcane–Frost | High Order | 797 | 833 (+37.7) | 832 (+34.7) | 97% | learned | 6 added, 1 removed |
| Priest Smite | Gnome | 503 | 519 (+17.3) | 522 (+19.8) | all | learned | 2 added, 1 removed |
| Paladin Retribution | Human | 958 | 991 (+32.3) | 1003 (+45.4) | all | learned | 6 added, 1 removed |
| Hunter BM Weave | Human | 1166 | 1201 (+34.4) | 1180 (+13.9) | 37% | published | 2 added, 1 removed |
| Paladin Physical Ret | Human | 890 | 910 (+17.2) | 928 (+38.0) | all | learned | 5 added, 0 removed |
| Mage Fire | Human | 930 | 948 (+17.6) | 958 (+28.3) | all | learned | 6 added, 1 removed |
| Hunter Pet/Melee | Windshaper | 932 | 948 (+15.8) | 946 (+14.0) | 87% | learned | 5 added, 0 removed |
| Warrior Fury (Sunder) | Human | 972 | 988 (+13.1) | 976 (+3.5) | 22% | learned | 1 added, 1 removed |
| Warlock Affliction | Human | 915 | 926 (+11.6) | 926 (+11.1) | all | learned | 0 added, 0 removed |
| Warrior Fury | Windshaper | 1008 | 1020 (+10.9) | 1013 (+4.4) | 35% | learned | 2 added, 0 removed |
| Rogue Combat | Windshaper | 754 | 760 (+6.2) | 760 (+6.4) | all | learned | 2 added, 0 removed |
| Shaman Enhancement | Windshaper | 813 | 821 (+6.3) | 820 (+6.9) | all | learned | 1 added, 2 removed |
| Druid Feral | Windshaper | 832 | 837 (+5.4) | 832 (+0.6) | 10% | learned | 1 added, 0 removed |
| Warlock Demonic Pact | Human | 1006 | 1010 (+5.0) | 1013 (+7.7) | all | learned | 0 added, 2 removed |
| Rogue Mutilate | Windshaper | 736 | 739 (+3.0) | 743 (+7.1) | all | learned | 1 added, 0 removed |
| Warlock DS/Ruin | Human | 879 | 881 (+2.4) | 881 (+1.7) | 78% | learned | 1 added, 0 removed |
| Druid Balance | Windshaper | 796 | 797 (+1.1) | 798 (+1.7) | all | learned | 0 added, 0 removed |
| Shaman Elemental | Orc | 636 | 636 (+0.6) | 642 (+5.8) | all | learned | 0 added, 1 removed |
| Warlock Destruction | Human | 885 | no gain (+0.3) | — |  |  |  |
| Hunter Survival | Human | 985 | no gain (+0.1) | — |  |  |  |
| Priest Shadow | Human | 913 | no gain (-0.0) | 915 (+1.9) |  | learned | 0 added, 0 removed |
| Paladin Shockadin | Human | 880 | no gain (-3.9) | — |  |  |  |
| Rogue Subtlety | Windshaper | 718 | no gain (-6.4) | — |  |  |  |

26 builds where the network beat the published setup, 27 simmable profiles that beat it, 31 builds.

Where a profile passes its network, the simulator-checked edits went beyond what the network found: the network
pointed at a choice, and the selection kept the version of it that scored best. Notable edits:

- **Frostfire:** the published APL spends spare mana on Fire Blast (rank 7, about 22 casts a fight) whenever mana
  exceeds ten times the seconds remaining. Dropping that rule, using a Mana Jade and keeping Fire Blast for the last
  three seconds gives +90 DPS.
- **Arcane–Frost:** one rule, "cast Frostbolt when Arcane Blast would cost over a quarter of current mana", is most of
  the +35. It needs the cost of a spell as a share of current mana, which the converter now reads from the APL's
  existing values (spell cost over current mana).
- **Frost:** the learned 18/2/31 talents take Arcane's Missile Barrage; with Arcane Missiles on Missile Barrage and no
  Evocation or Ice Lance, +80.
- **Retribution and Physical Ret:** Hammer of Wrath whenever it is usable, Consecration early in a Seal of Command
  window, and Seal of Righteousness between Holy Strikes, +45 and +38.
- **Elemental and Stormcaller:** no Chain Lightning at rank 4, +6 and (with Clearcasting rules) +32.
- **Hunters:** the weave builds keep a quarter to a third of the network's gain (Multi-Shot on cooldown, a better
  Immolation Trap rule, no Summon Hawk recasts for Beast Mastery). The rest comes from the network's weave timing, which
  threshold rules have not captured; their pages show it.

These edits are candidates for the published APLs; they have not been adopted, so the matrix stays comparable with
earlier releases.

## Fixes found on the way

- **The training environment's copy of the published APL** ran Autocast Other Cooldowns ahead of every APL entry and
  without its global-cooldown rule, and applied the automatic rules to cooldowns the APL casts itself. It drifted from
  the engine by up to 0.8% (Arcane–Frost +6, Fury (Sunder) −5.6, Beast Mastery −4.4, Demonology +4.4 DPS). It now
  matches the engine within noise, and every gain above is measured against the engine's own APL path.
- **Stoneskin Totem** registered its aura during the fight, which the engine rejects: casting it crashed Enhancement.
  The aura is now registered at setup.
- **Druid forms and energy refunds:** Cat Form now needs an energy bar and Bear Form a rage bar, and energy refunds fall
  back to the bar's metrics; an agent casting Cat Form as Balance crashed the engine. Feral, Bear and Balance results are
  unchanged.
- **APL Wait:** a 50 ms Wait rule that gained 17 DPS in the environment halved Beast Mastery's engine DPS (1184 against
  638). The APL's Wait adds its own pending action, which runs the rotation again on its own schedule. The converter no
  longer writes Wait rules; whether the Wait action itself misbehaves is open.
