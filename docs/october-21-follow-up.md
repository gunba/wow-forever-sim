# October 21 follow-up

A second pass on the same day as the [October 21 update](october-21-update.md), on client 1.60.1.70450. It works through
the wowsims/forever items that update left unchecked, makes the ranged Hunter trees weave, takes a fix Forever shipped
the same day, adds the party buffs wowsims offers, and re-runs all 872 matrix scenarios (seed 1793600001) and the
68 tank checks. Detailed numbers are in the [comparison](../artifacts/correctness/october_21_follow_up/comparison.json)
and the [selection evidence](../artifacts/correctness/october_21_follow_up/selection-evidence.json).

## Rotations

- **Marksmanship and Beast Mastery weave.** Both trees now step from 12 to 5 yards for Raptor Strike (and Mongoose Bite,
  Strider Kick and Immolation Trap where known) while the next Auto Shot is more than 1.5 s away, then step back out, as
  wowsims/forever's MM Weave and BM Weave presets do. They start from the SV Weave gear and consumables. MM Weave takes
  wowsims' Marksmanship/Survival talents, re-searched to 0/28/23 (Expose Prey and Strider Kick, so it keeps its own
  Hunter's Mark); BM Weave keeps its 31/20/0 talents (wowsims' BM Weave talents were 12.8 DPS worse here) with the Cat
  and two Hawks. The gear search kept six Beast Mastery changes (+5 to +6 DPS) and none for Marksmanship. The ranged
  presets stay in the simulator as "Marksmanship (ranged)" and "Beast Mastery (ranged)". (HUN-019)
- **Fury Hamstring weave.** Fury, Fury (Sunder) and 2H Bloodthirst cast Hamstring when Bloodthirst and Whirlwind are at
  least 1.5 s away, outside execute, for the extra Windfury Totem procs every landed special rolls (wowsims/forever #243).
  Arms is unchanged: Hamstring cost it 30 to 115 DPS in every placement tried, against its Slam. (WAR-027)
- **Retribution twisting.** Forever keeps one Twist of Light Echo at a time: gaining an Echo drops any other, and casting
  a seal drops its own Echo (wowsims/forever #218, from an in-game clip, and #124). The old rotation re-twisted at once
  and banked Command and Righteousness Echoes together. Under the corrected rule that dropped the banked Echo before a
  swing could spend it (Retribution 1078 to 829); the rotation now waits until the banked Echo is spent, which recovers
  most of it (956). Three other orders were tested and rejected. (PAL-002)

## Mechanics

- **Rage from damage taken** follows the rule wowsims/forever fitted to Forever combat logs (#356): 2.5 x (1 - A) x the
  hit before armor, blocks and absorbs / c(attacker level), with crits and crushing blows counted and A the armor
  reduction the game assumes for the attacker's level (20% up to 42, 30% at 44, 36% at 46, 38% at 50, 40% at 63). It
  replaces the provisional coefficient-10 model, which gave a well-armoured tank about a third of the rage against a boss;
  a manual setting in the simulator still selects the old formula. The planner's rage forecast uses the same rule.
  (WAR-003)
- **Race offsets.** Forever fixed its race stats on 21 October: the original races take the original game's offsets
  again (Tauren +5 Strength, -5 Agility, +2 Stamina, -5 Intellect, +2 Spirit; Undead +5 Spirit; Night Elf +5 Agility),
  undoing the modern offsets the October 21 update took from level-1 sheets. The Skyborne keep -1/+1/-1/+1 until their
  new offsets show ([forever-bugs#414](https://github.com/ClassicWoWCommunity/forever-bugs/issues/414),
  wowsims/forever #619). (CORE-027)
- **Paladin.** Consecration's ticks can crit (#124). Sanctified Judgement returns its share of the seal's base cost times
  (2 - the percent cost modifier), 59.4 mana on a 90-mana seal with Benediction 5/5 (#507). Consecrated Ground's mark
  lasts until 0.6 s after the last tick (#516). Vindication procs 5 times a minute on the weapon's speed from swings and
  specials, and its debuff rolls spell hit (#409). (PAL-008, PAL-018)
- **Mage.** Arcane Missiles' missiles no longer roll Clearcasting (a log has none from 34 casts); Blizzard and Flamestrike
  ticks can crit (#489). Arcane loses about 11%. Arcane-Frost's talents were re-searched and moved a point from
  Frostbite to Frost Channeling. (MAG-013)

## Party buffs

Dark Inspiration (a Shadow Priest's talent ranks, +3% Spirit a rank), Mana Tide Totem (count), Darkspear's Ancestor's
Vengeance (4% crit at the 16% uptime one 3.20-speed wielder gives) and Admiral's Hat (+10 Stamina) are new inputs. The
benchmark does not use them. Concentration Aura is not offered: the simulator does not model interrupts. (CORE-030)

## Checked and already matching

Hunter pets' crit, hit and expertise inheritance; Multi-Shot instant; Expose Prey; Resourcefulness; Arcane Blast
charges on a miss; Combustion's first stack; Master of Elements on Clearcast crits; Seal of Command's Judgement of the
Crusader share after the damage multipliers; the Warrior Glory sets; Fingers of Frost's 15% at the cast; Forever's
Arcane Blast row (15% of Intellect as spell damage a stack, not wowsims' older Arcane Charges).

## Not changed

Demoralizing Shout's rank-5 threat (43 vs the client's 47 a target); Fingers of Frost's cast/impact split with
Frostbite (same 15% total, slightly different timing); Seal of the Crusader's effect on Vindication's chance.

## Results

Best race per build, 5-minute benchmark (seed 1793600001), against the October 21 update:

| Build | Best DPS | Was | Mean change |
|---|---|---|---|
| Hunter SV Weave | 1258 (Human) | 1258 | -0.1% |
| Hunter MM Weave | 1182 (Human) | 1051 | +12.0% |
| Hunter BM Weave | 1166 (Human) | 1024 | +13.7% |
| Warlock Demonic Pact | 1006 (Human) | 1006 | -0.0% |
| Warrior Fury | 988 (High Order) | 969 | +1.8% |
| Hunter Survival | 985 (Human) | 985 | -0.1% |
| Paladin Retribution | 958 (Human) | 1083 | -11.4% |
| Warrior Fury (Sunder) | 953 (High Order) | 942 | +1.2% |
| Hunter Pet/Melee | 932 (High Order) | 933 | -0.1% |
| Mage Fire | 930 (Human) | 930 | +0.1% |
| Warlock Affliction | 915 (Human) | 915 | -0.0% |
| Priest Shadow | 913 (Human) | 913 | -0.0% |
| Paladin Physical Ret | 890 (Human) | 1017 | -12.5% |
| Warlock Destruction | 885 (Human) | 885 | +0.0% |
| Paladin Shockadin | 880 (Human) | 879 | +0.0% |
| Warlock DS/Ruin | 879 (Human) | 880 | +0.1% |
| Warrior Arms | 837 (Human) | 837 | -0.1% |
| Warrior 2H Bloodthirst | 834 (Orc) | 803 | +3.9% |
| Druid Feral | 832 (High Order) | 832 | +0.0% |
| Shaman Enhancement | 813 (Windshaper) | 814 | -0.1% |
| Mage Frostfire | 803 (Human) | 802 | +0.1% |
| Mage Arcane–Frost | 797 (High Order) | 823 | -3.0% |
| Druid Balance | 796 (High Order) | 795 | +0.1% |
| Rogue Combat | 754 (High Order) | 753 | -0.0% |
| Rogue Mutilate | 736 (High Order) | 737 | -0.1% |
| Rogue Subtlety | 718 (High Order) | 718 | +0.0% |
| Mage Arcane | 698 (Gnome) | 787 | -11.6% |
| Shaman Elemental | 636 (Orc) | 636 | +0.0% |
| Mage Frost | 610 (Gnome) | 609 | +0.1% |
| Druid Bear · Tank | 602 (High Order) | 534 | +12.7% |
| Paladin Protection · Tank | 559 (Undead) | 559 | -0.1% |
| Warrior Protection · Tank | 542 (Undead) | 492 | +8.2% |
| Shaman Stormcaller | 531 (Tauren) | 533 | -0.1% |
| Priest Smite | 503 (Gnome) | 503 | +0.1% |

Tanks (two seeds, 5,000 iterations each, both workloads): Protection Warrior threat +12.5% single target and +19.8% against three, damage taken +2.1%; Bear threat +15.1% single target and +5.4% against three, damage taken +0.0%; Protection Paladin threat -0.1% single target and -0.0% against three, damage taken +0.1%.
