# Race coverage

The roster is captured from client **1.60.1.69893**:
[CharBaseInfo](https://wago.tools/db2/CharBaseInfo?build=1.60.1.69893),
ChrRaces and ChrClasses. It contains 56 race/class pairs. Across the 23 DPS
builds this gives **147 results: 71 Alliance and 76 Horde**.

| Class | Alliance | Horde |
|---|---|---|
| Druid | Night Elf, High Order | Tauren, Windshaper |
| Hunter | Human, Dwarf, Night Elf, High Order | Orc, Tauren, Troll, Windshaper |
| Mage | Human, Gnome, High Order | Orc, Troll, Undead |
| Paladin | Human, Dwarf | Undead |
| Priest | Human, Dwarf, Night Elf, Gnome | Troll, Undead |
| Rogue | Human, Dwarf, Night Elf, Gnome, High Order | Orc, Troll, Undead, Windshaper |
| Shaman | Dwarf | Orc, Tauren, Troll, Windshaper |
| Warlock | Human, Gnome | Orc, Troll, Undead |
| Warrior | Human, Dwarf, Night Elf, Gnome, High Order | Orc, Tauren, Troll, Undead, Windshaper |

High Order and Windshaper are the two Skyborne entries. Their client IDs are
95/96 and playable-mask bits 32/33; those are not the simulator's enum values.
The original gearsets have no known faction-locked pieces, so both factions
use the same equipment and external buff package. Both now have Paladins and
Shaman.

The old Horde-only gates on Strength of Earth, Grace of Air, Stoneskin and
Mana Spring are disabled for Forever's Alliance characters. Keeping those Era
gates would incorrectly remove 53 Strength and 89 Agility from the Alliance
benchmark before stat multipliers.

## Alliance racial checks

Values and raw spell effects are retained in
`assets/db_inputs/forever_races.json`. Rebuild that capture with:

```sh
python3 tools/database/import_forever_races.py --cache /tmp/forever-client-data
```

### Eureka

SkillLineAbility identifies a different spell for each Gnome class. All have
three charges, a 15-second window and a two-minute cooldown.

| Class | Spell | Resource-cost reduction | Damage/healing bonus |
|---|---:|---:|---:|
| Rogue | 1259812 | 20% Energy | 10% damage |
| Warrior | 1259813 | 40% Rage | 10% damage |
| Mage | 1259817 | 50% Mana | 10% damage |
| Warlock | 1259821 | 50% Mana | 10% damage |
| Priest | 1259823 | 15% Mana | 10% damage/healing |

The old generic implementation incorrectly applied a global damage bonus and
could leave melee attacks benefiting without spending charges. The replacement
applies the modifier to charged damaging/healing casts, including tagged effects,
and spends a charge even when a spell suppresses its cast-complete event.

Cost reductions use the engine's additive cost-modifier convention; the
damage/healing bonus is multiplicative. DoTs retain their cast snapshot and
tagged channel ticks retain the charged cast's bonus. These interaction rules
need the in-game comparison in **T03**; the client values alone do not prove
every stacking or charge-consumption edge case.

### Other passives and cooldowns

- Expansive Mind raises maximum Mana, Rage or Energy by 5%, not Intellect.
  Rage and Energy were missing from the prior implementation.
- Elune's Light uses spell 1259799: 10 percentage points of crit for 15 seconds,
  with a three-minute cooldown.
- Human Sword Specialization is 2 percentage points of all attack/spell crit
  while a sword is equipped. Dwarf Mace Specialization is 1 point while a mace
  is equipped; holding two qualifying weapons does not double either bonus.
- Quickness's client dodge value remains 1 percentage point. Its separate
  movement-speed value is 2%; they must not be read as the same stat.
- Wind Blessed and Elemental Insight retain their haste and creature-type
  effects for both Skyborne entries.

### Read Ley Line

SkillLineAbility assigns spell 1259705 only to **High Order**. It has a
two-second cast, a 1.5-second GCD and a two-minute cooldown. The ordinary buff
lasts 15 seconds; a nearby ley line can instead give 15 minutes. Both buff
records specify 100% health/mana regeneration.

The benchmark does not assume a nearby ley line, and the active regeneration
effect is not yet used. Its relationship to Spirit regeneration, mana/5 and the
five-second rule is tracked as **T04**, rather than silently granting the
favourable location buff.

The actionable tests, including checks feasible at level 20, are in
[In-game checks](in_game_checks.md).
