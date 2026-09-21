# Priest beta pass (17 September 2026)

Beta client `1.60.1.69893` against Classic Era `1.15.9.69722`, read with `tools/data_watch/spell_client.py`, the
talent curves with `tools/data_watch/trait_curve.mjs`, and `../beta/priest.json`.

## How the numbers were read

Same rule as the mage pass: a rank's damage is the client's base plus `EffectRealPointsPerLevel` for every level from
the rank's base level to its max level (capped at 60), low end rounded down, high end up. Against the Era client that
reproduces every max rank the sim had and most lower ranks. The ranks that never matched Era (Smite 4, 6, 7; Holy Fire
2, 4, 6; Mind Blast 7) now follow the rule too; none is cast at 60. Dots are stored in the sim as a total, so the
client's per tick value was multiplied by its tick count (SWP 6, Devouring Plague 8, Mind Flay 3, Starshards 6, Holy
Fire 5). Every coefficient below is per tick for dots, which is how both the client and the sim store it.

## Spells changed (old -> new, level 60 rank unless noted)

| Spell | Change |
|---|---|
| Smite (8 ranks) | rank 8 384-429 -> 166-187, lower from rank 3 up; coefficients ranks 1-3 .123/.271/.554 -> .429/.571/.714 |
| Holy Fire (8 ranks) | rank 8 355-449 -> 184-232, dot 145 -> 75; about half at every rank. Coefficients .75/.05 unchanged |
| Mind Blast (9 ranks) | rank 9 508-537 -> 477-504; coefficients ranks 1-2 .268/.364 -> .429 |
| Shadow Word: Pain (8 ranks) | rank 8 852 -> 762; .2 a tick at every rank (was .167, ranks 1-3 .067/.104/.154) |
| Mind Flay (6 ranks) | rank 6 676 -> 390 (demo-scaled 119 at rank 1 -> the client's 63; Era was 426); .15 -> .167 a tick |
| Devouring Plague (6 ranks) | rank 6 904 -> 848; .063 -> .1 a tick; cooldown 3 min -> 1 min |
| Starshards (7 ranks) | rank 7 936 -> 1800; .167 a tick at every rank. Also fixed: ticks 1-5 of the channel divided the total by the tick index instead of 6 |
| Holy Nova (6 ranks + 6 healing) | was rank 1 only, 30-35 / 64-73 at 22% base mana, .143/.286 -> six ranks, highest known registered: rank 6 (27801) 174-200 damage, (27805) 288-334 healing, 750 mana, .107 both |
| Penance | 47540 (Wrath id), 93 a bolt, .268, 16% base mana, 10 sec -> rank 4 1316995: 131 a bolt (1316993), .285, 355 mana, 12 sec. APL `ui/smite_priest/apls/launch.apl.json` now casts 1316995 |
| Power Infusion | 16% -> 20% of base mana; 3 min, +20% unchanged |
| Vampiric Embrace | gains the client's 1 min cooldown (Classic 10 sec, sim had none); 30 sec and 20% confirmed |

Unchanged in the client and left alone: every cast time and mana cost on the ranked spells, Mind Blast's 8 sec
cooldown, Inner Focus (100% / 25% / 3 min), Shadow Word: Pain's 6 ticks every 3 sec.

`ui/core/spells/priest.json` has no unreviewed entries left: the spells above are `forever` with tooltips, Penance
stays `assumed` (bolt timing and the unsimulated heal), Spirit Tap's five aura ids are `classic`.

## Talents changed from the client curves

| Talent | Go before | Client |
|---|---|---|
| Spiritual Guidance (damage half) | 1/2/3/4/5% of Spirit | 1/3/5/6/8% |
| Mental Agility | 3/6/9% | 3/7/10% |
| Searing Light | 2/4% Holy damage, free Holy Nova never expires | 2/5%, Holy Purpose (1284536) lasts 10 sec |
| Shadow Weaving | 33/66/100% | 33/67/100% (the tree read 66 too; `ui/core/talents/trees/priest.json` now reads 67) |
| Darkness | 2% a point on five named spells (base damage only on Mind Blast and Devouring Plague) | 15259 is aura 79: +2% a point to all Shadow damage done, now a school multiplier |

Confirmed as they were: Power in Light 2..10, Twin Disciplines 1..5, Holy Precision 6/12/18, Silent Resolve
10/20/30, Meditation 17/33/50, Mental Strength 3..15, Holy Specialization 1..5, Spell Warding 2..10, Divine Fury
0.1..0.5 sec, Shadow Focus 1..5, Shadow Affinity 10/20/30, Improved Shadow Word: Pain 3/6 sec, Improved Mind Blast
0.5..2.5 sec, Improved Mind Flay 10/20, Devouring Contagion 25/50, Shadowform 10/50/100/15.

## Checklist lines

### Priest section
- **Resolved** `sim/priest/devouring_plague.go:53` Devouring Contagion rank 2: curve -25/-50, the sim's 50% stands.
- **Resolved** `sim/priest/holy_nova.go:13` higher ranks and mana cost: six ranks, 185..750 flat mana (not 22% of base), see table.
- **Resolved** `sim/priest/mind_flay.go:81` Improved Mind Flay: curve 10/20.
- **Resolved** `sim/priest/penance.go:18` cost, cooldown, level 60 damage: 355 mana, 12 sec, 131 a bolt (rank 4).
- **Open** `sim/priest/power_infusion.go:10` casting Power Infusion on a raid member: a sim feature, not a number
  the client holds. The cost half of the neighbouring line is settled.
- **Resolved** `sim/priest/power_infusion.go:25` cost and cooldown: 20% of base mana (was 16%), 3 min.
- **Open** `sim/priest/priest.go:81` Divine Spirit / Improved Power Word: Fortitude baseline or removed. Divine Spirit
  (14752-27841) and Prayer of Fortitude are still in the priest's SkillLineAbility with no race mask, and Divine
  Spirit's duration was retuned 30 -> 60 min, so it was not removed; but SkillLineAbility lists talent-taught spells
  the same way as trained ones, so it cannot say whether Divine Spirit is trained. Improved Power Word: Fortitude's
  passives (14749, 14767: +15/30%) still exist but nothing in the Forever tree or spellbook grants them. The buff
  values live in `sim/core/buffs.go` (out of scope); the client's Divine Spirit rank 4 is still +40 Spirit.
- **Resolved** `sim/priest/priest.go:97` Devouring Plague race lock: SkillLineAbility race mask is Undead (16) in Era
  and -1 (every race) in Forever.
- **Resolved** `sim/priest/talents.go:40` Spiritual Guidance: 5/10/15/20/25% healing, 1/3/5/6/8% damage (the damage
  half was wrong in Go).
- **Resolved** `sim/priest/talents.go:57` Power in Light: 2% a point.
- **Resolved** `sim/priest/talents.go:167` Searing Light rank 2: 5% Holy damage (not 4%), 10% chance.
- **Resolved** `sim/priest/talents.go:400` Shadowform mana discount: 15473 effect 3 is -50.
- **Open** the healing spellbook paragraph. The commented-out healing spells stay commented; the client's numbers for
  when they come back: Improved Power Word: Shield 7/14/20 (the checklist's tree says 7/14/21), Spiritual Healing
  3/7/10, Silent Resolve healing threat 10/20/30.

### Talents the sim does not read (Priest lines)
All still **open, not simulated**, and none became something a raid rotation reads. The client's rank values, for the
record: Wand Specialization 13/25% (the sim has no wand attacks), Martyrdom 50/100% after a crit taken, Improved
Inner Fire 15/30/45% armor and 4/8/12 charges (Inner Fire not in the sim), Soul Warding -4 sec / -15% Power Word:
Shield, Improved Mana Burn 0.5/1 sec (no Mana Burn), Renewed Hope 2..10% heal crit, Divine Aegis 5/10/15% of crit
heals, Twilight Focus 23/47/70% pushback avoidance, Blessed Recovery 8/17/25%, Holy Reach 10/20% range, Binding Heal
(heal), Litany of Light 5/10% mana on alternating heals, Spirit of Redemption (death), Blackout 2..10% stun (bosses
immune), Spirit Tap 20..100% on a kill (nothing dies), Shadow Reach 10/20%, Improved Psychic Scream 2/4 sec,
Improved Fade 3/6 sec, Silence, Early Demise +15/30% Shadow Word: Death crit at or below 20% (Shadow Word: Death
401955 is still a rune spell, AcquireMethod 3 in both clients, not trained).

### Baseline ability changes
No Priest line in that section beyond the spellbook diff below.

## Spellbook, Forever against Era (`--learned priest`)

New in Forever:
- **Penance ranks 2-4** (1240720, 1240721, 1316995 with bolts 1240727, 1240730, 1316993 and heals), learned at 40,
  50, 60; rank 1 (402174) is now learned at 30 with 100 mana instead of SoD's 16% of base. Implemented at rank 4.
- **Binding Heal ranks 2-6** (1240770-1240774) and **Prayer of Mending ranks 2-3** (1240826, 1240827): healing, not simulated.
- New race-locked priest spells, replacing Classic's racial priest spells alongside them:
  - **Chastise** (Dwarf, 1277331-1277335): instant, rank 5 272-306 Holy at .143, 225 mana, 2 min cooldown, 2 sec
    root, "Only works against Humanoids". A damage spell, but it only works on humanoids, and the sim has no
    humanoid check on a target, so it is not implemented; Dwarf builds would need it if a boss is humanoid.
  - **Dark Sacrifice** (Undead, 1277324-1277328): rank 5 turns 1600 health into 1600 mana over 15 sec, 10 min cooldown.
    Implemented for all five ranks with automatic mana-cooldown use. Self-damage currently uses unmodified client
    amounts; its interaction with damage buffs and resistances remains unverified.
  - **Divine Grace** (Human, 1277370-1277378): instant heal on a target below 50%. Healing.
  - **Contingency Plan** (Gnome, 1277462-1277640) and **Confounding Flash** (Gnome, 1277455): ward and crowd control.
  - **Expansive Mind** (20591) and the Gnome's Eureka! (1259823), Undead Touch of the Grave (1260201) show up as racials.
- Journeyman Riding, Equip Transmog Outfit: not abilities.

Gone from the list: Command (Orc racial), the SoD tuning passives, and every per-rank talent spell id (Darkness 2-5,
Shadow Weaving 2-5, Spirit Tap 2-5, Improved Mind Blast 2-5, ...): Forever keeps one spell per talent and reads the
ranks from a curve. Starshards (Night Elf, 8), Elune's Grace (Night Elf), Feedback (Human), Touch of Weakness,
Hex of Weakness, Shadowguard, Desperate Prayer and Mana Burn are all still there.

## Surprising in the data

- **Penance rank 4's bolt (131) is smaller than rank 3's (180)** while its heal keeps climbing (184, 291, 482, 673).
  The sim uses 131 as the client says; worth a look on the live beta.
- **Holy Fire's dot is not monotone**: 4, 5, 6, 7, 13, 10, 13, 15 a tick. Rank 5 out-ticks rank 6.
- **Holy Fire lost about half its damage** (rank 8 355-449 -> 184-232) while Mind Blast lost 6%.
- **Mind Flay's demo 119 was wrong**: the client's rank 1 is 21 a tick, 63 in total, and every rank sits below Classic.
- **Starshards nearly doubled** and Devouring Plague's cooldown went from 3 min to 1 min.
- Shadow Weaving's buff (15258) is now a self buff on the priest (aura 270, 2%), confirming the sim's Forever model.
