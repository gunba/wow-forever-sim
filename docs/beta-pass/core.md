# Shared effects and integration

The class passes checked client spell tables, but several changes had not reached
shared auras, preset flags, or tooltip overrides. A class pass marked “resolved”
does not establish that its raid-buff equivalent was also updated.

## Level-60 buff values

The following values are verified against Forever tooltips and covered by
`TestForeverBuffValues` in `tools/forever_bench`. The test measures applied
character stats, including stat dependencies, with both legacy input states.
Removed improvement talents must not multiply these values again.

| Effect | Value | Source |
|---|---|---|
| Mark/Gift of the Wild | 385 armor, 16 attributes, 27 resistances | [Mark of the Wild](https://www.wowhead.com/forever/spell=9885/mark-of-the-wild) |
| Power Word: Fortitude | 70 Stamina | [Fortitude](https://www.wowhead.com/forever/spell=10938/power-word-fortitude) |
| Strength of Earth | 53 Strength | [Strength of Earth](https://www.wowhead.com/forever/spell=25361/strength-of-earth-totem) |
| Grace of Air | 89 Agility | [Grace of Air](https://www.wowhead.com/forever/spell=25359/grace-of-air-totem) |
| Battle Shout | 139 melee AP at level 60, three minutes | [Battle Shout](https://www.wowhead.com/forever/spell=25289/battle-shout), client SpellEffect/SpellLevels |
| Blessing of Might | 133 melee AP | [Might](https://www.wowhead.com/forever/spell=25291/blessing-of-might) |
| Blessing of Wisdom | 40 mana per five seconds | [Wisdom](https://www.wowhead.com/forever/spell=25290/blessing-of-wisdom) |

The static Wowhead preview scales Battle Shout to the rank's maximum level, 61,
and displays 140. Its tooltip embeds `ppl60:61:139:60`: base level 60, maximum
61, base value 139, growth 0.6 per level. The level-60 benchmark uses 139.
The “Requires level 60” heading is not the preview's scaling level.
Both Battle Shout and Might explicitly say **melee** AP.
The old improvement talents are absent from the Forever trees. Booming Voice
changes shout radius, not duration.

Strength of Earth and Grace of Air last five minutes. An externally maintained
benchmark buff remains permanent when the shaman registers its own totem.

### Raid buff sources

Campfire benefits are not enabled in the raid benchmark. Ordinary food buffs
remain separate. The [Basic Campfire](https://www.wowhead.com/forever/spell=1307227/basic-campfire)
also has explicit raid/instance placement restrictions.

Forever air-totem replacement removes the old totem's lingering Windfury
tracking buff: a single shaman cannot twist Windfury and Grace of Air.
Replacing one's own totem does not remove a permanent buff assigned to another
provider. If both external effects are selected, they require separate
providers, not twisting. Dynamic party-wide totem distribution is still
represented by the party/raid buff settings, rather than the shaman's tracking
auras alone.

## Debuffs

- [Expose Armor](https://www.wowhead.com/forever/spell=11198/expose-armor):
  450 armor per combo point, 2,250 at five. Its talent changes Energy cost and
  combo-point refunds, not armor reduction. The shared aura previously used
  Classic's 1,700 × 1.5. `TestForeverExposeArmor` covers application and reversal.
- [Curse of Recklessness](https://www.wowhead.com/forever/spell=11717/curse-of-recklessness):
  505 armor reduction, no AP bonus. Covered by
  `TestForeverRecklessnessHasNoAttackPower`.
- Curse of the Elements covers all six magic schools. Resistance reductions
  apply once to each resistible school. The old Curse of Shadow toggle has no
  effect under Forever. Covered by `TestForeverCurseCoversAllMagic`.

The complete armor package is **2,250 major + 505 Curse of Recklessness +
505 Faerie Fire = 3,260**. Sunder and Expose replace each other; they do not
stack. Starting armor of 4,638 / 3,731 / 3,009 therefore becomes
1,378 / 471 / 0, with armor clamped at zero. `TestForeverBossArmorShred`
covers these cases and competing major debuffs.
This confirms the armor totals, not a change to the mitigation equation.
At attacker level 60, 471 armor gives 7.89% reduction under the current formula.

## Ability integration

- [Aspect of the Beast](https://www.wowhead.com/forever/spell=1299447/aspect-of-the-beast):
  level-60 rank implemented at 110 melee AP and 110 Mana, mutually exclusive
  with Hawk. Deadly Aspects enables
  [Quick Strikes](https://www.wowhead.com/forever/spell=1299448/quick-strikes):
  30% melee haste for 12 seconds. Regression checks cover stats, exclusivity,
  haste school, and reversal.
- Wrack channel interruption previously failed because readiness checks rejected
  every prospective spell while still channeling. The readiness probe now
  accounts for cancellation. Separate tests cover actual channel clipping and
  the existing damage modifier ending when Wrack is cancelled.
- Pet dismissal now cancels a queued combat-start auto-attack. Sacrificed pets
  no longer resume attacking at the pull.
- Hunter pets retain their selected base attack speed, rather than being
  forced to two seconds. `TestHunterPetKeepsSelectedAttackSpeed` covers
  1.0, 1.2 and 2.0 seconds. The current Cat presets select 1.2 seconds.

## Horde racials

Checked against client `1.60.1.69893` SpellEffect, SpellCooldowns,
SpellAuraOptions and SkillLineAbility, plus the linked Forever spell pages.

| Racial | Implementation |
|---|---|
| [Blood Fury](https://www.wowhead.com/forever/spell=20572/blood-fury) | Dynamic 10% melee AP, ranged AP and spell power for 15 seconds; free, no GCD, two-minute cooldown. Previously omitted ranged AP, snapshotted stats and consumed a GCD. |
| [Berserking](https://www.wowhead.com/forever/spell=20554/berserking) | Free, no GCD, exactly 10% melee/ranged/casting speed for ten seconds, three-minute cooldown. Removed Classic resource costs and mana-user 11.1% haste. |
| [Touch of the Grave](https://www.wowhead.com/forever/spell=1260198/touch-of-the-grave) | Triggered drain is 5% of maximum health before Shadow mitigation and cannot crit. One-second ICD; Warrior/Paladin/Rogue use passive 1260189 at 5%, Priest/Mage/Warlock use 1260201 at 10%. Removed the guessed random half-to-full damage range. |
| [Skysight](https://www.wowhead.com/forever/spell=1259686/skysight) | Movement speed, not a damage cooldown. Removed the unsupported Windshaper AP/SP cooldown; Wind Blessed's 1% haste and Elemental Insight remain. |
| [Endurance](https://www.wowhead.com/forever/spell=20550/endurance) | 5% maximum health and 1% physical/spell hit; existing implementation matches. |
| [Axe Specialization](https://www.wowhead.com/forever/spell=20574/axe-specialization) | 1% crit to spells and abilities while an axe is equipped; existing implementation matches both crit pools. |

Runtime tests cover Blood Fury's stat changes and later stat gains, all resource
types' Berserking speed/cost, removal of the unsupported power cooldown, and
Touch of the Grave's variants, ICD, damage and approximate proc rate.

## Open integration work

This is not a completed audit.

| Area | Remaining work |
|---|---|
| Shared buffs | Finish remaining aura/talent checks and remove obsolete improvement choices from UI presets and tooltip overrides. |
| Warrior resources/combat table | `sim/core/rage.go` still generates auto-attack rage from post-mitigation damage, using `damage × 7.5 / GetRageConversion(level)`. The DPS Warrior enables this with multiplier 1; there is no Forever normalization branch. A [September 18 beta report](https://us.forums.blizzard.com/en/wow/t/rage-nerfs/2354302) says crits give no extra rage, but supplies neither controlled measurements nor a replacement formula. Classic WarriorSim formulas are a baseline, not independent Forever evidence. Verify rage, miss/crit suppression and glancing against beta evidence before treating Warrior rankings as validated. |
| Racials | Damage-relevant Horde checks above are complete; Alliance-only abilities have not had this pass. |
| Mage | [Frostfire Bolt](https://www.wowhead.com/forever/spell=1237313/frostfire-bolt) now has all three ranks, direct and periodic damage, Fire/Frost school handling, and its explicit Improved Fireball, Hot Streak and Missile Barrage interactions. Winter's Chill's crit bonus remains limited to Frostbolt and Ice Lance. Rotations/talents still need optimization with the new spell. Frostbolt also now records misses instead of silently dropping them from metrics. |
| Priest | Undead [Dark Sacrifice](https://www.wowhead.com/forever/spell=1277328/dark-sacrifice) is implemented, with all five ranks, a shared ten-minute cooldown and automatic mana-cooldown use. At level 60 it converts five ticks of 320 health to Mana over 15 seconds. The static tooltip displays the rank's level-68 maximum, 1,640, not the level-60 total of 1,600. Self-damage currently uses unmodified base points; damage-buff/resistance interactions remain unverified. |
| Druid | Removed Cat Mangle and Feral Faerie Fire remain registered. The new benchmark APL does not cast them, but their registration and other UI rotations still need reconciliation. |
| Hunter | Summon Hawk models one guardian although the tooltip allows two; guardian attack damage/timing remains an approximation. Pet-family abilities and aspect/proc interactions need review. |
| Warlock | Incinerate's Immolate-bonus scope and Demonic Brand damage remain uncertain. Wrack's tooltip says other Shadow DoTs, while Wowhead's affected-spell list does not clearly establish every interaction; do not treat Doom-specific amplification as verified. |
| Other server behaviour | Fingers of Frost charge timing, Venom's effect on an existing poison DoT, Maelstrom Weapon proc frequency, and certain missing spell coefficients still need evidence beyond client tables. |

Old benchmark charts predate these corrections. Simulation sampling error does
not measure uncertainty in the underlying mechanics or provisional armor.
