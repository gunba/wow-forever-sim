# In-game checks

Client data establishes many values but not every server-side interaction.
These are the checks most likely to change simulated damage or the preferred
rotation. The open checks below still need measurements.

## Available early

### T50 — Trainer and spellbook availability

**Priority:** medium · **Access:** level 20 trainer lists; repeat at the listed learning level

Capture the trainer's future-spell list and current spellbook, with spell IDs
where possible, for Hunter Black Arrow/active Lacerate, Mage Felfire, Warrior
Victory Rush and Paladin Hammer of the Righteous. Pet Lava Breath and other
family teaching records belong with T29. Retained class-skill rows do not
establish availability. Note any quest, tome, racial or talent teaching path.

**Resolves:** the availability gaps in the [spell coverage inventory](spell_coverage.md).
This is separate from acknowledged implementation gaps such as Cone of Cold
and Hellfire; a trainer screenshot alone does not implement an ability.

### T51 — Cat stealth openers

**Priority:** medium · **Access:** level 20 Prowl; Pounce/Ravage at their learning levels

Record whether Prowl can be applied before entering combat in Cat Form, then
the timing of its removal, the opening ability and the first white swing.
At later access, compare Pounce on stun-susceptible and stun-immune targets:
does an immune stun still apply its damage? Preserve energy cost, positional
requirements, spell rank and combo-point results.

**Current model:** Prowl, Pounce and Ravage are not implemented. The Feral
profile does not claim an opening-stealth damage contribution.

### T29 — Hunter pet inheritance and trained abilities

**Priority:** high · **Access:** level 20 initially; repeat at later levels

With a permanently tamed pet, record its family, level, trained abilities,
happiness, base swing interval and AP. Change owner melee AP and ranged AP
separately, crossing which value is higher. Repeat with a temporary buff,
resummoning and instance transfer. Check player haste separately from direct
pet haste. Preserve normal/critical white and spell damage events.

**Resolves:** the observed low-level ten-percent higher-AP inheritance and its
dynamic behavior. The engine currently inherits no owner stats. Later-level
measurements or a source establishing level scaling are needed for level 60.
Do not use quest-controlled NPC abilities as pet-training evidence.

### T30 — Paladin seals, Holy Strike and Consecration

**Priority:** high · **Access:** level 20 for early ranks; later for extra-attack sources

- Record seal duration through repeated Judgements; distinguish expiry from
  consumption.
- Compare Holy Strike ranks with asymmetric weapon damage and controlled AP/SP.
  Preserve weapon speed, talent points and normal/critical damage.
- On flat ground, compare Consecration with and without added SP on one
  through six targets. Repeat center/edge positioning; separate crits and
  the first four targets from later targets.
- At later access, separate Wisdom events from white swings, seal damage,
  Windfury/Reckoning and Twist of Light echoes. Test recursion and cooldowns
  rather than inferring unlimited mana from a historical proc chain.

The [matched Windfury comparison](windfury.md) separates white damage and Seal
gains. Specifically test whether Command damage can trigger the totem, whether
Righteousness damage cannot, and which attacks consume its AP charges. A large
combined white/Seal gain alone does not establish a duplicate buff.

### T32 — Inner Focus and periodic critical strikes

**Priority:** medium · **Access:** level 20 with Inner Focus

Compare SW:P casts with and without Inner Focus, recording mana, aura
consumption and every tick. An individual crit does not establish that the
whole DoT snapshots the bonus. Record enough casts to distinguish the crit
rates. Separately record Starshards cost, ticks and crits on a Night Elf.

### T27 — Fire Nova and its Fire totem

**Priority:** medium · **Access:** Shaman level 20

Try Fire Nova with no Fire totem, an active Searing Totem, and after that
totem expires or is destroyed. Record the cast result, mana/GCD and damage
source. The client says the Nova originates from an active Fire totem;
the simulator now enforces that prerequisite. A threat capture can separately
check attribution to the Shaman rather than the totem.

### T26 — Druid talent ranks and Faerie Fire in forms

**Priority:** high · **Access:** level 20 for the early talents and form tests; later access for Moonfury

- Compare zero, one and two Nature's Majesty points with the same equipment:
  record spell/melee crit and check yellow attacks. The captured tooltip says
  2/4%, but the raw client scalar needs rank/curve interpretation.
- Measure Nature's Reach hit and Moonglow mana costs at each rank. Include
  Faerie Fire and a heal as negative controls for Moonglow.
- Cast regular Faerie Fire in humanoid, Cat and Bear forms. Record spell ID,
  mana/Energy changes, GCD, cooldown, whether the form remains, and auto timing.
  The simulator now uses regular Faerie Fire in forms, following its client
  shapeshift mask, mana cost and GCD.
- At later access, compare Moonfury's zero/five-point damage on the same
  Arcane/Nature spell. The reviewed sources disagree between 5% and 10%.
- With Improved Moonfire available now, compare normal direct hits and ticks
  at two known SP values, with and without the talent. This separates a
  base-damage-only bonus from one that also multiplies SP contribution.
  Repeat for Moonfury when accessible.
- Compare Thick Hide at each rank and form with known bonus Defense. Its form
  condition is explicit, but the Defense coefficient's raw scalar and earlier
  rank text disagree; the inherited coefficient is not yet validated.

**Resolves:** actual rank scaling, shared-hit scope and confirmation of the
source-derived Faerie Fire behavior. Current positive/negative talent-family tests prevent
Moonglow from discounting unrelated mana spells.

### T25 — Omen of Clarity proc frequency and consumption

**Priority:** high · **Access:** Druid level 20 with Omen trained · **Status:** provisional proc model

Capture several minutes of Cat autos, consuming each Clearcasting proc with a
Claw. Separately log repeated Wrath casts and direct healing, then test periodic
damage/healing alone. Record proc/refresh timestamps and eligible events after
each ten-second interval; distinguish an internal cooldown from an additional
random roll. A humanoid weapon-speed comparison would help identify PPM behavior.

Check that Wrath, Faerie Fire, Tiger's Fury and other free actions preserve the charge,
while a paid offensive/healing action becomes free and consumes it. Also check
a proc arriving during a cast, a canceled cast and a missed free attack.

**Resolves:** the missing passive is restored, but its effective proc rate
could substantially change Feral Energy and Balance mana availability.
Current client 16864 lists 100% proc chance, ten-second recovery and no PPM
entry; 16870 lists one charge and a 15-second duration. The simulator currently
interprets those entries literally. They do not rule out server-side overrides.
Source: [Forever Omen of Clarity](https://www.wowhead.com/forever/spell=16864/omen-of-clarity).

### T24 — Thistle Tea and shared cooldowns

**Priority:** medium · **Access:** Rogue or level-20 Cat with tea · **Status:** class/effect sourced; interaction test open

Use tea below ten Energy and record the resource gain, remaining form and
cooldown. Check potion and other consumable cooldowns before and after use.
Repeat while out of Cat form if possible, separating usability from whether
the restored resource survives a subsequent form change.

**Resolves:** server-side form behavior and shared cooldown categories.
The current Forever item permits Rogue and Druid and restores 100 Energy on
a five-minute item cooldown. Both energy-using builds are now supported.
The existing two-minute shared-conjured timer is inherited, not established
by the item's five-minute tooltip.

### T22 — Healing-only gear and enchant damage

**Priority:** high · **Access:** level 20 if a suitable item or enchant is available · **Status:** open

Record spell damage and bonus healing before and after equipping a healing-only
item. Repeat with a healing enchant, keeping the underlying item unchanged.
Record item/enchant IDs and check all spell schools, for example with
`GetSpellBonusDamage`. Exclude items whose client data already has explicit
spell damage.

**Resolves:** whether any genuinely healing-only record has an unlisted server
effect. The unsupported blanket conversion has been removed. Bracer Healing
Power is not a suitable healing-only test: its client effects explicitly give
24 healing and eight damage. Gloves give 35/12 and the weapon enchant gives
55/19. These hybrids retain their actual damage; they do not establish a
universal conversion rule.

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

For Furor, leave Cat at a recorded Energy value, wait known intervals and
return at one through five talent points. Test the combined carry/regeneration
cap at low ranks. The engine now counts exits at the pull and before it, and
clears this state between iterations; the low-rank cap interpretation still
needs gameplay confirmation.

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

**Priority:** high · **Access:** Warrior · **Status:** low-level evidence available; level scaling and off-hand unresolved

Use a target being held by someone else, or a training target, so incoming damage
does not also generate rage. Record level, weapon damage/speed, attack power,
talents, and racial/buff effects. Start from zero rage and record several ordinary
hits with the damage and rage change visible. Repeat with a substantially faster
or slower weapon. At level 20, an off-hand comparison is also useful if dual wield
is available.

Enable advanced combat logging and preserve the matching equipment/level
SavedVariables from [Forever State](https://github.com/tzcnt/forever-data).
The advanced log exposes resource snapshots in tenths of a rage point.
Separate incoming hits, spending, resource procs, caps, gear changes and
out-of-combat decay from each measured swing.

Reported level-8–10 measurements suggest weapon-speed-only generation:
4.5 rage per weapon-second for two-handed weapons and about 3.46 for
one-handed main-hand attacks, with ordinary rage on crits/glances and none
on misses/dodges. The separately supplied level-18, 2.7-second one-hand capture
has now been checked. A conservative selection excludes nearby spells/incoming
attacks within 250 ms and retains 29 landed swings; allowing close timestamps
while retaining the other exclusions gives 54. Every retained swing gains
9.3 or 9.4 rage, across normal hits, glances and crits. The 54-swing sample spans
15–78 damage and averages about 3.4602 rage per reported weapon-second.
This supports normalized rather than damage-proportional generation.

The earlier screenshot's 78-swing selection is not used as an independently
reproduced count; its exclusions were not supplied. Weapon speed and level
remain the recorder's stated context, not parsed equipment metadata.
Avoided swings have no immediate resource snapshot, so this analysis does
not independently establish their zero-rage behavior.

The seven public logs at commit `c7d17462c50d1eb0103aa5e2aff52f77f33e3418`
parse successfully, but lack the matching level/equipment/speed metadata.
There are 478 otherwise eligible snapshot candidates, not 478 verified
weapon-rate measurements. These support investigating damage-independent
rage, but do not reproduce the exact one-hand/two-hand coefficients.

**Resolves:** base versus hasted weapon speed, level dependence, off-hand
generation, avoided attacks, and whether extra attacks or on-next-swing
specials follow the same rule. The engine still has inherited
damage-proportional rage and dodge/parry generation; those remain a material
Warrior uncertainty, not validated Forever behavior. Do not extend the
low-level coefficients to level 60 or invent an off-hand factor without
additional evidence or an explicit provisional model.

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

Compare Lightning Bolt ranks 2, 3 and 4 with two known spell-power totals, keeping
talents and the target unchanged. Record spell tooltips and roughly 30 ordinary,
unresisted, non-critical hits for each rank at each spell-power total.
Changing only spell-power equipment makes the comparison easier to interpret.

For each rank, divide the change in average damage by the change in spell power.
The model uses the client coefficients: 0.571 for rank 2 and 0.714 for ranks
3 and 4. A remaining
server-side low-level/downranking penalty would make the measured coefficients
differ. Damage talents can multiply both measurements and should be recorded.

**Resolves:** whether the low-rank fillers used for mana efficiency retain their
client-listed scaling. This affects Elemental and Stormcaller rotation choices
and potentially other caster classes. The model currently uses the client values.

A Priest can also test Smite ranks 2 and 3 at level 20 using two spell-power totals.
Their modeled coefficients are 0.571 and 0.714 respectively, before damage
talents. Rank 2 is the fallback in the current Smite rotation.

### T49 — Ice Lance coefficient and frozen-target multiplier

**Priority:** high · **Access:** level 20 with Ice Lance; repeat at higher ranks

Use two known spell-power totals and record ordinary, non-critical Ice Lance
hits against the same unfrozen target. Record rank, level, talents and buffs.
Divide the change in average damage by the change in spell power, accounting
for damage talents. The engine currently assumes a 0.143 coefficient because
the captured damage effect does not supply one.

Repeat against a target frozen by Frost Nova, separating normal hits from
crits. The engine's Fingers of Frost case multiplies the entire hit, including
spell power, by four.
A rank-1 measurement does not establish the level-60 rank's coefficient.

At later access, separate Fingers of Frost activation, Ice Lance launch,
impact and charge consumption, including a Frostbolt already in flight.
This affects both the ordinary Frost and Arcane–Frost rows. T10 covers the
related deeper Mage proc interactions.

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

### T34 — Mage Clearcasting consumption

**Priority:** high · **Access:** Arcane Concentration talent, available by level 20

Record a proc, cast a paid damage spell, then immediately cast another.
Capture the buff disappearing and the mana charged for both casts. Repeat
with Arcane Missiles and, separately, a utility spell. Distinguish the proc
from a new proc caused by the consuming cast.

**Current model:** one eligible damage cast, not fifteen seconds of free
casts. The one-second proc cooldown follows client metadata; missile-trigger
and utility-spell edge cases still need a server observation.

### T35 — Finisher Energy refunds

**Priority:** high · **Access:** Rogue Eviscerate/Expose Armor/Rupture and
Druid Rip by level 20; Ferocious Bite needs later access

Record Energy immediately before and after a successful finisher and an
avoided one, retaining hit outcome, talent points, combo points and natural
regeneration events. Use a higher-level target if necessary to obtain
dodges. A miss and a dodge may not have the same behavior.

**Resolves:** whether these finishers refund 80% of paid Energy. A later
upstream commit changes this, but the commit alone is not server evidence.
The current engine retains the older no-refund behavior for these finishers.

### T36 — Sanctity Aura availability

**Priority:** high · **Access:** trainer/talent inspection now; cast test if obtainable

Check whether a Forever Paladin can actually learn Sanctity Aura, and record
the teaching source, spell ID and required level. If available, compare a
party member's Holy damage with and without it.

**Current model:** spell 20218 still has legacy spell/talent records but no
current trait or `SkillLineAbility` learn path. The benchmark now excludes
its 10% Holy-damage buff. An actual acquisition source would warrant revisiting
that exclusion; a database tooltip alone does not establish availability.

### T40 — Haste and the global cooldown

**Priority:** high · **Access:** level 20 with a suitable haste effect

On a Troll caster, record an instant spell's GCD duration and a hardcast's
duration before and during Berserking. Record the actual buff and health
state rather than assuming a fixed haste amount. Use client cooldown timings,
not only the delay between manually pressed buttons.

**Current model:** general spell haste shortens casts but does not shorten the
default GCD. Explicit GCD modifiers, such as Nature's Grace, are separate.
This is an important input to the scaling chart and needs a Forever measurement.

### T43 — Fel Armor availability

**Priority:** medium · **Access:** trainer/spellbook inspection now; later if level-gated

Record whether Fel Armor is learnable, its spell ID, teaching source and
required level. A spell or skill row alone is insufficient: the retained
403619 record uses a different acquisition method from trained Demon Armor.
If obtainable, record SP and healing before/after at two Spirit values.
The engine currently grants no Fel Armor bonus.

### T44 — Flat mana regeneration rate

**Priority:** high · **Access:** level 20 with a known flat mana-regeneration effect

Record the tooltip, spell/item ID, client build and mana changes for at least
60 seconds with and without one flat mana-regeneration effect. Keep Spirit,
Intellect and other regeneration sources unchanged. Start below full mana and
separate continuous casting from five-second-rule recovery. Exclude potions,
mana refunds and direct restoration procs from the comparison.

If the effect advertises `X` mana per five seconds, the additional recovery
should distinguish `12 × X` from `60 × X` over 60 seconds. Record tick spacing
as well as the total; a displayed per-second rate does not by itself establish
one-second ticks.

**Current model:** the default uses MP5 / 5 per second. An opt-in provisional
setting uses MP5 per second, without multiplying Spirit or direct mana returns.
Both retain the existing two-second tick cadence.

### T45 — Seal of Fury and shield-break mana

**Priority:** medium · **Access:** level 20, if the early trained ranks are available

Record the teaching source and spell IDs. With a shield equipped, compare
isolated white attacks, specials and Judgement; then repeat without the shield.
Track which attacks produce Holy damage and an absorb. Proc twice before the
first shield is consumed to distinguish replacement, stacking and refresh.

With Improved Seal of Fury, compare partial absorption, complete depletion,
expiry and replacement. Record mana against same-level and higher-level
attackers. Keep incoming damage and other mana sources identifiable.

**Current model:** Fury and its shield-break restoration are not implemented.
Another engine implements them but assumes replacement and white-hit-only
triggers. The client effects establish a useful starting point, not every
server interaction.

### T47 — Lightning Shield and shock cooldowns

**Priority:** low · **Access:** level 20

With no special gear effects, cast Earth Shock, then rank-1 Lightning Shield
while the shock is still cooling down. Once the global cooldown ends, check
whether Earth Shock remains unavailable until its original cooldown finishes.
Repeat with rank 2 and a fresh versus replaced shield. Record spell IDs and
cooldown timestamps.

**Current model:** applying Lightning Shield does not reset shocks. The old
SoD callback reset Earth Shock's individual timer on the initial three charges,
but the separate shared shock cooldown still prevented early recasts. The
callback has been removed; Forever's captured shield effects do not specify it.

### T48 — Spellcasting and melee swing timers

**Priority:** high · **Access:** level 20 for ordinary Lightning Bolt; later for Maelstrom

With a known weapon speed and no haste procs, record white swings, Lightning
Bolt start/completion, and the first swing after casting. Start casts at several
points in the swing cycle, then chain casts without a gap. Compare genuinely
short casts with casts that occupy the whole global cooldown. Record interrupted
casts separately.

At later access, repeat at one, three and five Maelstrom stacks. In particular,
measure whether the instant five-stack Lightning Bolt resets the timer. Compare
Chain Lightning as an ordinary cast, not as a Maelstrom spender.

**Current model:** the original Classic Shaman helper schedules the next swing
at cast completion plus one full swing duration, including zero-time casts.
It does not release a delayed swing immediately at cast completion. Slam and
Hunter weapon shots have explicit separate exceptions; the earlier blanket
permission for melee during spellcasting has been removed.

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
| T21 | Maelstrom Weapon proc rate, free-cast interactions and Totem of the Storm | Count procs per landed melee attack with two weapon speeds, separating Windfury extra attacks. Record whether a five-stack free Lightning Bolt consumes or can generate Clearcasting. Then test Lightning Bolt out of melee range with item 272432: its tooltip specifies half the ordinary chance, not a spell conversion from PPM. The current 2 PPM per talent point remains unverified. | Maelstrom Weapon talent; level 60 for the totem |
| T23 | Hot Streak charge consumption and interrupted casts | At one, two and three stacks, record Pyroblast's cast time and remaining buff; repeat with a cancelled cast and a filler projectile landing during the Pyroblast cast. The client confirms one charge and three stacks; the engine consumes the buff on completed Pyroblast. | Hot Streak talent, beyond level 20 |
| T28 | Windfury Weapon and Totem trigger rules | Separate weapon and totem procs, log proc spacing and extra-attack outcomes, and compare attacks with/without the weapon imbue while receiving the totem. The client states 20% weapon chance and excludes the matching totem benefit; the inherited 1.5-second ICD and extra-attack resolution still need measurement. | Windfury Weapon/Totem, beyond level 20 |
| T31 | Hunter melee proc and crit scope | Compare Predator's Edge normal/critical autos and specials separately. For Expose Prey, distinguish melee/ranged attacks, traps, misses, refreshes and Mongoose cooldown state on a marked target. | Relevant deeper Hunter talents |
| T33 | Shadowform and Priest buff access | Confirm Shadowform's mana/GCD and Holy damage, shield, heal and Holy Nova permissions. Record Divine Spirit's trainer acquisition after talent removal rather than inferring availability from a database skill entry. | Shadowform at level 40; relevant Divine Spirit ranks |
| T37 | Shadow Word: Death damage and backlash | Record hits/crits above and below 20% target health, surviving versus killing blows, and misses/absorbs. Repeat with damage-increase/reduction buffs to check backlash modifiers. The implemented client values are four ranks, .429 SP, a shared 15-second cooldown and 10% maximum-health backlash; script value 150 is unresolved and no extra execute multiplier is assumed. Verify damage-range rounding/level growth too. DPS-only fights record self-damage but do not model healer survival constraints. | Shadow Word: Death, level 32 or later |
| T38 | Mutilate's two-hand resolution | With distinguishable dagger damage, capture each hand's normal and critical hits across controlled AP values. Repeat with Cold Blood and with avoided main-hand strikes. Check whether Cold Blood guarantees both crits and whether the off-hand penalty also reduces the flat bonus. | Mutilate and Cold Blood, beyond level 20 |
| T39 | Demonic Brand damage and proc rules | Change Fire and Shadow SP separately; compare Imp Firebolt, Succubus melee/Lash, Voidwalker and Felhunter attacks. Record brand spell IDs, target, charges, crits/misses and threat if measurable. Change SP after applying the brand to distinguish snapshotting from power at hit time and owner from pet scaling. Current child formulas are 65–68 at level 60 plus 7.8% matching-school power; power attribution, other-pet schools, crit behavior and 3× threat remain provisional. | Demonic Brand, normally level 25 or later |
| T41 | Nature's Grace timing | Separate Wrath cast completion from projectile impact, recording when the haste aura appears. Check Moonfire/Swarm/Faerie Fire GCDs and whether a second crit refreshes the three-second duration. The model applies separate 10% cast-haste and GCD effects but retains Wrath's cast-completion proc timing. | Nature's Grace, normally level 30 |
| T42 | Weapon-stone effects | Compare Firestone's spell crit and Fire power with Spellstone's casting speed and Shadow power, then test whether Spellstone also increases Fire damage as its effect mask says. Verify that stones work with an off-hand equipped and replace other weapon imbues. | Firestone from level 28; Spellstone from 36 |
| T46 | Hammer of the Righteous availability | Confirm that spell 407632 is actually learnable, its teaching source, and whether it shares a cooldown with Holy Strike. Do not infer availability from a retained class-skill row or tooltip. If available, compare weapon-DPS scaling and target count. | Reported required level 40; trainer inspection may be possible earlier |

## Result record

For a completed check, retain the ID, client build, race/class/level, relevant
talents and equipment, target level/type, observations, and any log or recording.
The list will distinguish measured results from remaining interpretations.
