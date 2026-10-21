# October 21 update

This update stays on client 1.60.1.70450 (no class spell or talent rows changed from 70425 in the CDN tables). It brings
the simulator up to date with [wowsims/forever](https://github.com/wowsims/forever), public again, for everything merged
since our last review on 23 September (pull requests #39 to #618), reads the game's own server hotfix cache, takes the
in-game findings from MythicSim's fork, adds three builds (SV Weave, Frostfire and Shockadin) and Forever consumables,
and re-runs the benchmark: talents and consumables re-searched for every DPS build, a gear search for the new and
reworked builds, all 872 matrix scenarios and the 68 tank checks.

## Server hotfixes

The client's CDN tables do not carry server hotfixes, so the hotfix rows were decoded from the game's own cache
(`DBCache.bin`, build 70450, pushes up to 112985) and compared with the CDN tables. (CORE-028)

| Change | Before | Now |
|---|---|---|
| Champion of the Light | 10/20/30% of Intellect | 20/40/60% |
| Hunter's Mark rank 4 | 71 ranged attack power (below rank 3's 98) | 143 |
| Bloodthrill's Overpower window | 6 s | 5 s |
| Mage Frost tree | 19 talents | Piercing Ice removed, Frostbite and Permafrost swap places, Frost Channeling and Improved Blizzard move left |
| Elemental Precision | 5 ranks, 1% each | 3 ranks, 2/4/6% |
| Winter's Chill | any Frost spell, +40% to the next Cone of Cold a stack | Frostbolt, Frost Blast and Fingers of Frost Ice Lance only; +5% damage and +0.07 spell power coefficient a stack on the next Cone of Cold; stacks to the talent rank |
| Ice Barrier | absorb only (not simulated) | +3% Frost damage while the shield holds and Frigid Power +3% for 1 min; now cast by the rotation when talented |
| Missile Barrage | Arcane Missiles hits also proc it | Arcane Missiles no longer proc it |
| Frost Blast | 0.343 coefficient | 0.271 (not simulated) |

Mage talent strings use the hotfixed Frost tree; existing Mage profiles were remapped and re-searched.

## Engine changes from wowsims/forever

- **Spell resistance.** Forever combat logs show no level-based partial resists against higher-level targets, partial
  resists in 10% steps, and Spell Piercing past a target's resistance adding damage. The average resist is now
  R / (|R| + 9 x caster level), with the target's resistance floored at zero before piercing, and DoT ticks take a tenth.
  The 1% spell miss floor is gone (boss spell hit cap 17%). Against a resistance-free boss this removes the old ~6%
  level-based resist from every caster. (CORE-024)
- **Ahn'Qiraj book ranks are off.** Fireball 12, Frostbolt 11, Arcane Missiles 8, Heroic Strike 9, Revenge 6, Battle
  Shout 7, Blessing of Might 7, Shadow Bolt 10, Corruption 7, Immolate 8, Serpent Sting 9, Aspect of the Hawk 7,
  Starfire 7, Ferocious Bite 5, Backstab 9, Eviscerate 9, Deadly Poison V and the top Grace of Air and Strength of Earth
  ranks come from raid books; raids open no earlier than 9 December. Players and raid buffs use the last trainer rank
  (Battle Shout 111 attack power, Blessing of Might 112). (CORE-025)
- **A completed cast restarts the melee swing** unless the spell is marked Do Not Reset Combat Timers (Aimed Shot,
  Multi-Shot). Thorium Grenades, Hammer of Wrath and Holy Wrath now restart it. (CORE-026)
- **Race offsets** from level-1 character sheets: for example Tauren +2 Strength, -2 Agility, +2 Stamina, -2 Intellect
  (was +5/-5/+2/-5 and +2 Spirit); no race changes Spirit; the Skyborne are -1 Strength, +1 Agility, -1 Stamina,
  +1 Intellect. (CORE-027)
- **Beast Slaying, Big Game Hunter and Elemental Insight** raise damage only, not critical damage.
- **Threat.** Heroic Strike adds 2.54 threat per spell level and Cleave 1.53 per spell level, split between the targets
  it lands on; Sunder Armor adds 5% of attack power at every rank; rage threat counts the threat multiplier when the rage
  arrives and is split between the enemies, as is mana threat. (WAR-025)
- **Bloodthrill** procs only from main-hand white swings and the queued Heroic Strike or Cleave (Forever logs: no
  procs from off-hand swings, Hamstring, Overpower or Execute). (WAR-026)
- **Seal of Righteousness** per hit: value/100 x speed x 1.2 on a two-hander (0.12 of spell power), or x (0.85 - 0.2 x
  max(0, 2 - speed)) on a one-hander (0.10), as measured on eighteen weapons. (PAL-017)
- **Not taken:** Anger Management stays 1 Rage per 3 seconds. wowsims reads its aura as 1.5 per 5 seconds, but Classic
  Era carries the identical aura row and gives 1 per 3 seconds. Windfury Weapon's strikes still do not proc Flametongue
  Totem (Forever logs in wowsims #278).

## MythicSim in-game findings

- Immolation Trap and Explosive Trap's explosion take 13% of the higher of ranged and melee attack power (was 10% on the
  burn of both). (HUN-018)
- Strider Kick speeds the hunter up 30% for 3 seconds; Mongoose Bite is a main-hand attack; Lacerating Strikes rolls the
  running bleed's unticked damage into the new one.
- Earth Shock takes Stormstrike's bonus without spending the charge. (SHA-012)
- One-Handed Weapon Specialization raises every school except Consecration, Retribution Aura and Holy Shock.
  Judgement of Righteousness and Holy Strike keep the client rows' coefficients; the in-game readings are disputed.

## Consumables

The Merchant's Favor alchemy recipes and Forever cooking are in, chosen per build by simulation (CORE-029):

- Elixir of Ferocity (18 Strength and Agility), Elixir of Cunning (25 Agility and Intellect), Elixir of the Grizzly
  (25 Strength, 2% crit), Elixir of the Owl (25 Intellect, 2% crit), Elixir of Sages (25 Spirit, 2% crit).
  Only one elixir's crit applies, to attacks and spells alike.
- Major Spellblasting Potion (+47 spell damage) and Major Frenzy Potion (+80 attack power), 30 seconds on the potion
  cooldown.
- Swiftstrike Steak (20 Agility) and Prehistoric Pulled Raptor (20 Intellect).
- Scroll of Imbue Precision (+1% crit) and Accuracy (+0.5% hit) beside a stone or oil.
- Elemental Sharpening Stone is a Molten Core recipe and is no longer offered; Dense stones replace it.

Typical picks: casters take the Owl and a Precision scroll, keep Major Mana Potions (Balance and Shadow prefer
Spellblasting); hunters take Swiftstrike Steak, Cunning and Ferocity; rogues take Mongoose, Ferocity and Precision on both hands,
warriors Mongoose with Ferocity or Juju Power; Retribution takes Cunning.

## New builds

| Build | Best race | DPS | Setup |
|---|---|---:|---|
| SV Weave (Hunter) | Human | 1258 | `-00530501105-530200230050121151`; Auto Shot from 12 yards, steps to 5 yards for Raptor Strike, Mongoose Bite, Strider Kick and Immolation Trap while the next Auto Shot is over 1.5 s away; no pet |
| Shockadin (Paladin) | Human | 879 | `550030032101012--052252300012303`; Seal of Righteousness, Judgement, Consecration, Hammer of Wrath, Holy Strike, Holy Shock |
| Frostfire (Mage) | Human | 802 | `-1355000013013502-2035010311013002`; Scorch to five Fire Vulnerability stacks, Pyroblast on three Heating Up, Frostfire Bolt |

The rotations follow wowsims/forever's and MythicSim's presets. SV Weave tops the matrix, 28% ahead of Survival melee on
the same equipment model. The gear search ran for all three and for Arcane, Frost and Arcane–Frost; 23 changes were
confirmed on a second seed.

## Results

On the published setups, before any re-search, the engine changes move casters up 3-10% (the resist model), Arcane down
15% and Frost 12% (Missile Barrage, Piercing Ice, trainer ranks), warriors down 3-9% (trainer ranks, lower raid
buffs, Bloodthrill, grenade swing resets) and rogues down 2-4%. After the talent, consumable and gear re-search:

| Build | Change | Build | Change |
|---|---:|---|---:|
| Elemental | +19.4% | Arms | -5.0% |
| Fire | +17.0% | Fury (2H) | -5.1% |
| Stormcaller | +12.4% | Frost | -5.8% |
| Smite | +10.0% | Arcane | -8.3% |
| Balance | +9.8% | Protection Warrior (DPS row) | -5.3% |
| Shadow, warlocks | +8 to +9% | Fury | -2.8% |
| Marksmanship | +8.0% | Rogues | -3 to 0% |
| Retribution | +7.7% | Feral | -2.4% |

Elemental's gain includes a talent change the search found (Shamanistic Focus over a Restoration point, +6%).
Tank checks: Protection Paladin threat +4-6% (One-Handed Weapon Specialization), Protection Warrior -6% single target
and -18% on the weakest of three targets (trainer Heroic Strike, Cleave and rage threat split), Bear -4%.

[Matrix](../artifacts/modelled_gear/forever_dps_5min.png) ·
[results](../artifacts/modelled_gear/forever_dps_5min.json) ·
[comparison with the previous release](../artifacts/correctness/october_21/comparison.json) ·
[tank checks](../artifacts/tanks/current/summary.json) ·
[previous published files](../artifacts/history/03a5b49d5/manifest.json)
