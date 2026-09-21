# Hunter beta pass (17 September 2026)

The shot wind-up convention below is historical and has been removed.
See the current [auto-attack model](../auto_attack_audit.md).

Beta client `1.60.1.69893` against Classic Era `1.15.9.69722`, read with `tools/data_watch/spell_client.py`, the
talent curves with `tools/data_watch/trait_curve.mjs`, and `../beta/hunter.json`.

## How the numbers were read

Every rank table in `sim/hunter` was checked against both clients. Where the sim matched Era, only what Forever moved
was changed. Two conventions carry over from the Classic sim:
- **Shot wind-up.** The sim adds 0.5 sec to a ranged shot's client cast time (Classic Aimed Shot is 3 sec in the client
  and 3.5 in the sim, Multi-Shot 0 and 0.5). Forever's Aimed Shot is 2 sec in the client, so 2.5 in the sim; Sniper
  Shot 4 sec, so 4.5. Forever's Multi-Shot now shows 0.5 sec in the client itself; that is read as the same wind-up
  made visible, so the sim keeps 0.5 rather than 1.0.
- **Dot totals.** Serpent Sting's table is the total over 5 ticks, as before (client per tick x 5).

Three spells lost their spell power coefficient in the client entirely: Arcane Shot, Serpent Sting and Volley's tick
spell. Volley's channel carries .03 on a dummy effect, the same placeholder Blizzard and Rain of Fire carry. Classic's
coefficients are kept for all three and the manifest marks them `assumed`. For a hunter they only matter with spell
power gear.

## Spells changed (old -> new, level 60 rank unless noted)

| Spell | Change |
|---|---|
| Aimed Shot (6 ranks) | cast 3 -> 2 sec (sim 3.5 -> 2.5); bonus 70/125/200/330/460/600 -> 20/34/55/89/125/166. Mana and the 6 sec cooldown unchanged. |
| Arcane Shot (8 ranks) | 13/21/33/59/83/115/145/183 -> 20/26/39/65/94/134/170/217. No coefficient in the client (kept). |
| Multi-Shot | ranks 2-5 (14288, 14289, 14290, 25294) are gone from the spellbook. One rank, 2643: no flat bonus (was 150), 13.9% of base mana (was 230 flat), 0.5 sec cast, 6 sec cooldown shared with Aimed Shot. The APL now casts 2643. |
| Serpent Sting (9 ranks) | totals 20/40/80/140/210/290/385/490/555 -> 10/30/60/110/170/240/320/415/555. Rank 9 unchanged. No coefficient in the client (kept). |
| Volley (3 ranks) | damage now comes from a tick spell (1279721, 1279719, 1279715): 50/65/80 -> 70/91/112 a tick. The 1 min cooldown is gone. Renataki's Charm no longer resets it (nothing to reset). |
| Raptor Strike (8 ranks) | rank 1 mana 15 -> 10; bonus ranks 4-8 34/50/80/110/140 -> 30/35/40/55/70; rank 7 mana 80 -> 85 (the sim had 80, both clients say 85). |
| Mongoose Bite (4 ranks) | flat 25/45/75/115 -> normalized melee weapon damage plus 15/22/37/57. |
| Aspect of the Hawk | rank 6 110 -> 55 ranged attack power. Ranks 1-5 and 7 unchanged. Only cast when AQ content is off. |
| Rapid Fire | 40% melee attack speed as well as ranged. |
| Explosive, Immolation, Freezing Trap | spell ids moved from Season of Discovery's 4095xx (not in the beta client) to 13813/14316/14317, 13795/14302-14305 and 1499. Damage and cost unchanged; the 30 sec shared cooldown already in the sim is confirmed. Immolation Trap ticks every 3 sec in both clients; the sim had 1.5 sec. |
| Sniper Shot (talent) | guessed Steady Shot shape (1.5 sec, no cooldown, 110 mana, +160, borrowed id 56641) -> client ranks 1310687/1310785/1310786: 4 sec cast (sim 4.5), 15 sec cooldown, 365 mana, +160/225/295. The APL now casts 1310786. |
| Summon Hawk (talent) | borrowed id 131894 -> 1293241/1293525/1293526/1293527; dive bomb 53 -> 32/47/85/108 plus 5% of ranged attack power; mana 190 -> 80/105/135/190. The hawk lasts 18 sec (1293248). |
| Intimidation | 3% -> 8% of base mana (both clients say 8%), buff window 10 -> 15 sec. |
| Pet Lightning Breath | 36-41 / 78-91 / 99-113 -> 32-36 / 71-81 / 86-98; no growth per level. |
| Pet Screech (now Demoralizing Screech) | 12-16 / 19-25 / 26-46 -> 9-13 / 21-27 / 24-42, and a new 10 sec cooldown. |
| Pet Scorpid Poison | 3/6/8 -> 2/4/5 a tick. |

Unchanged in the client and left alone: Wing Clip, Aspect of the Hawk ranks 1-5 and 7, Rapid Fire's cost and cooldown,
Bestial Wrath, Quick Shots, pet Claw and Bite, trap damage and costs, Defensive State.

Counted by spell id: 70 ranks changed (4 of them removed), plus Strider Kick added.

## Checklist lines

### Hunter section
- **Resolved** `sim/hunter/aimed_shot.go`, baseline: every Aimed Shot rank is on the hunter with the Multi-Shot
  sentence in its tooltip. Cast time: the client has 2 sec, down from 3 (sim 3.5 -> 2.5). The damage change was not
  on the checklist and is larger than the cast change (rank 6 600 -> 166).
- **Resolved** `sim/hunter/aspects.go:13`, Deadly Aspects: every rank triggers the same Quick Shots (6150), 30% for
  12 sec. Held flat, as modelled.
- **Resolved** `sim/hunter/aspects.go:55`, Deadly Aspects proc chance: curve 2/4/6/8/10, as modelled.
- **Resolved** `sim/hunter/rapid_fire.go:15`, Rapid Killing: curve -60000/-120000 ms, so 1 then 2 min, as modelled.
  The kill buff (415407) is 20% on the next Shot for **20 sec**, not the 40 sec the checklist names. It is still not
  modelled because nothing dies in a boss fight.
- **Resolved** `sim/hunter/serpent_sting.go:43`, Improved Stings: curve 6/13/20 (1310661 carries 20 at rank 3), as
  the sim already had. Viper Sting and Scorpid Sting are not in the sim.
- **Resolved** `sim/hunter/talents.go:300`, Resourcefulness: cost curve 30/60 as modelled, **proc chance curve
  50/100, not 30/60** (fixed, tree and override updated). The buff (1242688) is 50% for 30 sec at both ranks; the sim
  now uses that id instead of the borrowed 34491.
- **Resolved** `sim/hunter/talents.go:354`, Rapid Recuperation: curve 25/50 for Serpent Sting, 50/100 for Rapid
  Killing. The buff (1242512) lasts 15 sec at both ranks; the sim now uses that id instead of the borrowed 53232.
- **Resolved** `sim/hunter/talents.go:378`, Expose Prey: curve 5/10, and the window (1310726) is 5 sec at both ranks.
- **Resolved, needs a core change** `sim/core/buffs.go`, Battle Shout and Blessing of Might: both are still aura 99,
  melee attack power only, so hunters still get nothing from them. Their values did move and are out of scope here:
  Battle Shout rank 7 232 -> 139 AP (2 min -> 3 min), Blessing of Might and Greater Blessing of Might rank 7
  185 -> 133 AP (both last 60 min now).

### Talents the sim does not read
- **Resolved** Strider Kick: 1317257, 100% normalized melee weapon damage, 8 sec cooldown, 5.81% of base mana.
  Implemented (`sim/hunter/strider_kick.go`), no longer marked not simulated. No shipped APL casts it.

### Baseline ability changes
No hunter lines existed. The spellbook diff below covers it.

## Spellbook, Forever against Era (`--learned hunter`)

New in Forever:
- **Aspect of the Beast ranks 2-4** (1299445-1299447, levels 40/50/60): 70/90/110 melee attack power, plus Quick
  Strikes (1299448, 30% melee speed) under Deadly Aspects. The level-60 rank is now implemented, mutually exclusive
  with Hawk, and used by the melee Survival APL. See [shared integration](core.md).
- **Summon Hawk ranks 1-4** (see above, implemented).
- **Counterattack rank 2** (1242634, level 30): 50% weapon damage plus 40, talent, still not simulated (needs a parry).
- **Strider Kick** (talent, implemented).
- **Pet family abilities**: Swipe, Pinch, Dismember, Tendon Rip, Mine!, Savage Rend, Web, Dust Cloud (5 ranks each,
  1264727-1265913). Each only teaches a spell; not implemented, the sim's pet families use Claw, Bite, Lightning
  Breath, Screech and Scorpid Poison.
- Hunter's Mark (in core, out of scope): ranks 1-3 up (20/45/75 -> 26/59/98 ranged AP), **rank 4 down 110 -> 71**.
- Riding, transmog, a test dummy: no combat use.

Gone in Forever:
- Multi-Shot ranks 2-5.
- Scorpid Sting ranks 2-4.
- Wyvern Sting (all ranks; no longer a talent).
- Season of Discovery's trap ids (409510-409535), Frost Trap 409520, Explosive Shot 409552/409554 and the S03 tuning
  passives, which the Era client carried but Forever does not.
- Per-rank talent spells (Lightning Reflexes, Surefooted, Survivalist, Trap Mastery, Improved Aspect of the Monkey and
  so on): talents moved to trait curves.
- Screech renamed Demoralizing Screech.

## Still open
- Arcane Shot, Serpent Sting and Volley coefficients: the client has none; Classic's are kept.
- Multi-Shot's 0.5 sec: the sim reads the client's new 0.5 sec cast as the old wind-up, not an extra 0.5.
- Summon Hawk's assault: the hawk is a guardian, the client gives its 18 sec but not its swings. The sim keeps ticking
  the rank's dive bomb base damage every 3 sec, one hawk at a time.
- Mongoose Bite rank 1 (1495) is claimed by `ui/core/spells/druid.json` (Hurricane's mana cost table reads as a spell
  id there), so the hunter manifest cannot declare it.
- Aspect of the Hawk rank 6 at 55 and Hunter's Mark rank 4 at 71 look like data slips in the beta client (each is below
  the rank before it), but they are what the client says.
