# In-game checks

Client data establishes many values but not every server-side interaction.
These are the checks most likely to change simulated damage or the preferred
rotation. The open checks below still need measurements.

## Available early

### T22 — Healing-only gear and enchant damage

**Priority:** high · **Access:** level 20 if a suitable item or enchant is available · **Status:** open

Record spell damage and bonus healing before and after equipping a healing-only
item. Repeat with a healing enchant, keeping the underlying item unchanged.
Record item/enchant IDs and check all spell schools, for example with
`GetSpellBonusDamage`. Exclude items whose client data already has explicit
spell damage.

**Resolves:** the inherited simulator rule adds one spell damage per three
healing power when a record has no explicit damage component. It also applies
to enchants: +24 healing bracers currently provide eight modeled spell damage.
The presence of that rule is not confirmation that Forever applies it to
every healing-only item or enchant.

### T20 — PvP vendor stock and level-65 class armor

**Priority:** high · **Access:** vendor access on either faction; Alliance stock especially needed · **Status:** open

Inspect the weapon, armor and accessory quartermasters and export their item
IDs, costs and tooltips. Buying level-60 equipment is not necessary.
Compare the Premier items with the Horde exports, especially weapon damage.
Check the level-65 class sets as well: for example Premier Champion's Magus
Handguards (272505), Premier Champion's Felweave Gloves (272556), and Premier
Lieutenant Commander's Silk Gloves (272749). Their current database entries
have no acquisition source, unlike the lower-level Premier items already
captured from vendors. Record the NPC and item ID rather than matching names.

**Resolves:** the item catalog has unrestricted client race masks and tooltips
that name both Grand Marshal and High Warlord, but the captured merchant stock
is from Horde vendors. Equip legality does not itself establish Alliance
acquisition or prove that faction counterparts have identical server stats.

### T15 — Energy recovery and haste

**Priority:** high · **Access:** Rogue or level-20 Cat · **Status:** open

Spend Energy below 40, stop spending, and record the numeric bar recovering.
Try an ability as soon as its cost is reached, to distinguish usable regeneration
from a smoothly animated bar. Record race, talents and buffs.

This macro captures the client's reported recovery rates and haste:

```lua
/run local a,b=GetPowerRegen(); print("Energy",UnitPower("player",3),"Regen",a,b,"Haste",GetHaste())
```

Repeat with Slice and Dice or another available attack-speed/haste effect, naming
the exact buff. Slice and Dice tests that particular effect; it does not establish
what haste rating does. A Skyborne character's racial is another useful case.

**Resolves:** usable recovery cadence, the 10 Energy/sec baseline, and whether
particular haste effects increase regeneration. The simulator now approximates
smooth recovery with 100-ms updates and assumes general haste increases Energy
recovery. Attack-speed-only effects are excluded from that assumption.
The client table confirms the base rate, not either server-side interaction.

An **Iron Counterweight** on a two-handed weapon is a useful equipment check
if one can be crafted at the available profession cap. Compare melee speed,
ranged speed (Hunter), and Energy recovery (Cat). Client spell 7217 specifies
melee attack speed, not ranged speed or general haste. The model now reflects
that distinction. Arcanum of Rapidity needs later access; its client effect
adds melee and ranged speed, but not casting speed.

### T16 — Energy costs and failed-attack refunds

**Priority:** medium · **Access:** Rogue or level-20 Cat · **Status:** open

Record the tooltip cost and Energy immediately before/after a landed builder,
a missed/dodged builder, and a missed/dodged damaging finisher. Include timestamps
so natural regeneration can be separated from the refund. Note the exact rank
and cost-reduction talents. For Gnomes, repeat a builder under Eureka.

**Resolves:** the engine's inherited 80% builder refund and zero damaging-finisher
refund. The simulator refunds a fraction of the amount actually paid, including
temporary discounts; client cost records alone do not verify these server rules.

### T01 — Weapon-skill and combat-table tooltips

**Priority:** high · **Effort:** a few screenshots · **Status:** open

Capture the expanded weapon-skill tooltip and character combat stats at level 20,
including any table of outcomes against higher-level targets. Include weapon
type and current/max skill. If practical, capture a second weapon with a
different skill value.

**Resolves:** the hit/crit suppression rules and any stated glancing penalty.
The engine still uses inherited Classic combat tables; flat defence penalties
and the smaller boss glancing penalty have not been established in this model.
Tooltip evidence alone will not establish the observed outcome distribution.

### T02 — Rage gained from outgoing white hits

**Priority:** high · **Access:** Warrior · **Status:** open

Use a target being held by someone else, or a training target, so incoming damage
does not also generate rage. Record level, weapon damage/speed, attack power,
talents, and racial/buff effects. Start from zero rage and record several ordinary
hits with the damage and rage change visible. Repeat with a substantially faster
or slower weapon. At level 20, an off-hand comparison is also useful if dual wield
is available.

A short recording is useful: a combat log may not expose white-hit rage changes
with enough precision on its own.

**Resolves:** whether outgoing rage follows damage alone, or contains a
weapon-speed/normalisation term. The level-60 conversion constant must **not** be
applied directly to a level-20 result. Critical hits, dodges and off-hand hits
are useful additional samples, identified separately.

### T03 — Eureka charges, scope and channels

**Priority:** high · **Access:** Gnome · **Status:** open

Capture the racial tooltip for the tested class. Activate Eureka and record:

1. Resource cost and charge count across three damaging casts.
2. Whether white swings, wand shots, a non-damaging ability, or a proc use a charge
   or receive the damage bonus.
3. For a caster, use a channel as the **third** charged ability. Compare its
   individual ticks with an unbuffed channel.
4. For a DoT, compare the ticks of a charged cast after the racial aura expires.

Useful early abilities include Sinister Strike, Fireball/Arcane Missiles, and
Corruption/Drain Life, depending on class and what the beta spellbook offers.
Record mana-reduction talents when comparing costs.

**Resolves:** charge consumption, modifier stacking, projectile/channel timing,
and DoT snapshots. Client data has separate class versions and resource-cost
reductions; the previous generic racial implementation is not sufficient.

### T04 — Read Ley Line during combat

**Priority:** medium · **Access:** Skyborne with the racial · **Status:** open

Away from a visible ley line, record the cast time, buff duration and resource
ticks after using it. Compare an idle period with continuous casting, recording
Spirit, mana/5, and the resource bar. If a ley line can be reached, capture the
same observations there separately.

**Resolves:** whether the bonus affects Spirit regeneration, flat mana/5, or both,
and whether it bypasses the five-second rule. Its ordinary and ley-line
durations must not be conflated. No favourable ley-line location is assumed
for the benchmark.

### T05 — Priest spellbook differences

**Priority:** medium · **Access:** any Priest · **Status:** open

Capture race, level, the trainer's available spells and the relevant spellbook
pages. Particularly useful are Devouring Plague, Starshards, Shadowguard,
Hex of Weakness and Dark Sacrifice, if present.

**Resolves:** class/race restrictions and replaced abilities. Retained
SkillLineAbility records are not by themselves proof that a spell can be learned.
This check does not require levelling a separate character solely for it.

### T12 — Low-rank spell-power scaling

**Priority:** high · **Access:** level-20 Shaman with some spell-power gear · **Status:** open

Compare Lightning Bolt ranks 3 and 4 with two known spell-power totals, keeping
talents and the target unchanged. Record spell tooltips and roughly 30 ordinary,
unresisted, non-critical hits for each rank at each spell-power total.
Changing only spell-power equipment makes the comparison easier to interpret.

For each rank, divide the change in average damage by the change in spell power.
The current client records a 0.714 coefficient for both ranks. A remaining
server-side low-level/downranking penalty would make the measured coefficients
differ. Damage talents can multiply both measurements and should be recorded.

**Resolves:** whether the low-rank fillers used for mana efficiency retain their
client-listed scaling. This affects Elemental and Stormcaller rotation choices
and potentially other caster classes. The model currently uses the client values.

A Priest can also test Smite rank 3 at level 20 using two spell-power totals.
Its client coefficient is 0.714, before damage talents. This directly checks the
low-rank filler considered by the Smite build.

### T17 — Mind Flay and channel haste

**Priority:** medium · **Access:** level-20 Shadow Priest · **Status:** open

Record Mind Flay's cast bar and tick timestamps with and without a named haste
effect. Troll Berserking is useful if available. Distinguish a shorter global
cooldown from a shorter channel or faster ticks.

**Resolves:** the engine currently keeps Mind Flay at three one-second ticks,
without channel haste. Client base duration alone does not establish the
server's haste behavior.

### T18 — Cat attacks and weapon damage

**Priority:** high · **Access:** level-20 Cat · **Status:** open

Compare Cat-form white attacks using two weapons with substantially different
weapon DPS, ideally plain weapons without attributes or procs. Keep the target
and other equipment unchanged; record the Cat-form attack-power and damage
tooltips for both. Collect ordinary, non-critical, non-glancing hits separately.

**Resolves:** whether weapon DPS contributes to Cat damage or attack power in
Forever. The current model inherits Classic's form weapon, rather than adding
retail-style weapon-DPS scaling. This is separate from Energy regeneration.

### T13 — Aimed Shot cast time and Auto Shot timing

**Status:** reported gameplay rule adopted

Aimed Shot uses the client's two-second base cast without an additional
half-second wind-up. Auto Shot continues during it. The engine now implements
this rule; a regression test compares the shot count with an idle-auto run.
Sniper Shot likewise uses its four-second client cast. Multi-Shot retains its
client-listed half-second cast, without suspending autos.

## Possible if suitable targets are reachable

### T06 — White miss and glancing outcomes

**Priority:** high · **Status:** open, target access uncertain

At maximum weapon skill, collect white attacks against targets of known level:
same level and, if reachable, three levels higher. Keep hit bonuses and equipment
fixed within each sample; separate single-weapon and dual-wield samples.
Several hundred swings are useful for outcome rates. Preserve individual
glancing damage values rather than only the average.

**Resolves:** dual-wield miss, glancing frequency and the glancing damage range.
Level-20 targets cannot directly validate a level-63 raid boss, but relative-level
comparisons can distinguish several competing formulas.

## Deferred beyond the current level limit

| ID | Question | Why it matters | Access needed |
|---|---|---|---|
| T07 | Which DoTs Wrack amplifies, and exactly when amplification ends after cancelling | Affliction channel timing and DoT selection | Wrack's 31-point talent, normally level 40 |
| T08 | Whether Incinerate's Immolate bonus also multiplies spell-power damage | Destruction scaling and filler choice | Incinerate's 31-point talent |
| T09 | Demonic Pact and sacrifice-buff coexistence with a different active demon | Demonology build identity and damage | Demonic Pact's 31-point talent |
| T10 | Frostfire interactions with Missile Barrage, Fingers of Frost and Tier 1 | Mage proc rates and rotation | Relevant deeper talents and the actual set effect |
| T11 | Tier 1 Insect Swarm's final second: partial tick, delayed tick, or aura tail | Balance refresh timing | Actual Tier 1 bonus |
| T14 | Summon Hawk guardian attacks, two-hawk coexistence and scaling | Beast Mastery damage and talent value; the engine approximates one guardian with periodic damage | Summon Hawk, normally level 25 or later |
| T19 | Penance bolt timing and haste | The engine spreads three ticks over two seconds, using Forever's periodic critical-hit rules; do not rely on partial-channel optimizations until this is checked | Penance, normally level 30 or later |
| T21 | Maelstrom Weapon proc rate and Totem of the Storm | Count procs per landed melee attack with two weapon speeds, then per Lightning Bolt cast while out of melee range with item 272432. Its tooltip specifies half the ordinary chance, but does not define a spell conversion from PPM. | Maelstrom Weapon talent; level 60 for the totem |

## Result record

For a completed check, retain the ID, client build, race/class/level, relevant
talents and equipment, target level/type, observations, and any log or recording.
The list will distinguish measured results from remaining interpretations.
