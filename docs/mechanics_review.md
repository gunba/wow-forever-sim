# Mechanics review

## Corrections

- **Healing and damage:** removed the blanket healing-to-damage conversion.
  Only explicit damage effects contribute. Current hybrid enchants retain
  their sourced values; see [gear data](forever_gear_data.md).
- **Hot Streak:** the next completed Pyroblast consumes all stacks. Client aura
  400625 has three cumulative stacks but one charge. Interrupted-cast timing
  remains an open check.
- **Elemental:** the two unused shield points move to Ancestral Knowledge,
  adding 4% Intellect. The existing rotation and equipment are unchanged.
- **Warrior shield talents:** Defiance, Bastion and Master of Defense now
  follow the currently equipped shield, rather than freezing their conditions
  from the starting loadout. Tests cover equipping/removing the shield and
  leaving Defensive Stance.
- **Thistle Tea:** Cat Druids can now select and use it. The
  [current item](https://www.wowhead.com/forever/item=7676/thistle-tea) permits
  both Rogues and Druids and restores 100 Energy on a five-minute cooldown.
  Existing benchmark consume choices are unchanged. The inherited
  cross-consumable shared cooldown still needs verification.

- **Omen of Clarity:** restored the missing trained Druid passive. The provisional
  model uses the current client entry: a direct attack/spell/heal can grant one
  Clearcasting charge after a ten-second cooldown. It waives the next eligible
  action's resource cost, but not Wrath or resource-free actions. The listed
  100% chance is **not a measured server proc rate**; see T25 before treating
  Feral or Balance results as validated against in-game resource generation.
- **Moonglow:** its mana discount now follows the damage-spell family mask.
  It no longer incorrectly discounts Faerie Fire, shapeshifting or healing.
- **Enrage:** Dire Bear now uses its stated 16% armor penalty, not ordinary
  Bear's 27%. The existing immediate ten Rage plus twenty over ten seconds
  was already implemented.
- **Faerie Fire:** replaced the obsolete Feral spell with regular Faerie Fire
  in all supported forms. Current client records allow it in Cat/Bear and
  retain 115 mana, a 1.5-second base GCD and no separate cooldown. The Bear APL
  now references the real spell; Cat/Bear autos continue.
- **Moonkin form:** now participates in the shared shapeshift tracking, so
  canceling the form actually removes its form state and armor multiplier.
- **Mana supplies in forms:** the supported mana potions, Demonic Rune and
  Minor Recombobulator no longer inherit a blanket shapeshift block or force
  an unshift in Forever. Their client effects have no form exclusion.
  Minor Recombobulator also uses its five-minute item cooldown, rather than
  the helper's default two minutes. This does not grant blanket permission
  to every consumable or explosive.
- **Innervate reporting:** its extra regeneration is now credited to Innervate
  rather than ordinary mana regeneration. Previously the Innervate row showed
  only its casting cost, producing an apparent negative return. This reallocates
  the existing gain; it adds no mana or regeneration threat. Ordinary
  regeneration fills missing mana first, and any excess bonus is shown as
  wasted. Innervate's cost remains payable. Overlapping Innervate and Evocation
  no longer disable each other's full casting regeneration when one expires.
- **Fire Nova:** requires the caster's active Fire totem; an expired totem does
  not qualify. Damage and threat remain attributed to the Shaman.
- **Windfury:** Windfury Totem is a separate, exclusive party air buff, not
  an imbue available to other classes. Melee reference groups use it instead
  of Grace of Air; ranged Hunters use Grace. Windfury Weapon still excludes
  the Shaman owner's Totem proc, including after weapon swaps. Flametongue
  Weapon does not exclude Windfury Totem.
- **Shadowform:** casting now costs 40% base mana and triggers a 1.5-second
  GCD. Its prepull starts 1.5 seconds before combat. Holy damage and non-healing
  utility remain allowed; active heals and Holy Nova are blocked in form.
  This does not implement the missing Priest healing spellbook.
- **Expose Prey:** only melee/ranged attacks can trigger it, matching the
  client proc mask. Generic trap spell hits no longer open a Mongoose Bite
  window.
- **Thick Hide:** armor now follows the active Cat/Bear/Moonkin form and
  current bonus Defense instead of being frozen from starting equipment and
  form. The inherited Defense coefficient remains provisional; the current
  raw scalar and earlier rank text need reconciliation.
- **Mage Clearcasting:** no longer persists because its own discount made
  the consumption check see zero mana cost. The next eligible paid damage
  spell consumes it; the trigger follows the client's one-second cooldown.
- **Starshards:** all ranks and partial channels now share the client's
  30-second cooldown.
- **Hunter Bite and projectiles:** corrected rank-eight Bite to 81–99 and
  replaced inherited projectile speeds with the corresponding client values.
  Shadow Bolt now has its missing travel time; Incinerate uses 20 yards/sec.
- **Temporary casting regeneration:** Improved Stormstrike and Spirit Tap
  now update the cached mana-tick amounts immediately on gain and expiration.
  Resourcefulness, Rapid Recuperation and Blue Dragon's proc now do so too.
- **Mutilate:** Puncturing Wounds now applies its critical-strike bonus to
  both hands, matching the client's spell-family mask, rather than only the
  main-hand strike.
- **Paladin cooldowns and equipment:** Exorcism, Holy Wrath, Holy Shock and
  Holy Shield ranks share their client-defined cooldown categories. Holy
  Shield also requires an equipped shield instead of being castable without one.
- **Warlock cooldowns:** Soul Fire, Death Coil, Shadowburn and Conflagrate
  now share a timer across ranks instead of permitting a lower-rank bypass.
- **Sanctity Aura:** excluded from Forever. Its legacy spell/talent record
  survives, but the current trait tree and class skill list provide no access.
  Old raid/party/personal settings no longer grant an unsupported 10% Holy bonus.

The client records behind the enchant, Hot Streak, shield and Omen conditions are
captured in `assets/db_inputs/forever_effect_audit.json`.

## Focused talent check

Both comparisons used the corrected engine, fixed Orc equipment and rotation,
and 5,000 iterations with independent validation seed 20292237.

| Build | Previous allocation | Tested allocation | Decision |
|---|---:|---:|---|
| Elemental: shield points → Ancestral Knowledge | 508.84 DPS | 510.56 DPS | Retain |
| Fire: one Wake of Fire point → fifth Improved Fireball point | 714.86 DPS | 703.57 DPS | Revert |

Fire already had four Improved Fireball points. These checks do not establish
a global optimum or justify changing every rotation.
The Fire comparison predates the Clearcasting consumption correction and
must not be read as a current resource validation.

## Druid resource check

Matched Tauren profiles, 5,000 iterations, seed 20292541:

| Build | Previous engine | Corrected engine | Change |
|---|---:|---:|---:|
| Feral | 602.13 DPS | 671.28 DPS | +11.5% |
| Balance | 593.56 DPS | 600.83 DPS | +1.2% |

The paired requests are identical. These checks include the provisional Omen
model and the related Druid corrections. They are resource-model
checks, not revised all-race rankings or a rotation optimization.

An additional 200-iteration accounting check (seed 20292564) used identical
Balance requests with one scheduled Innervate. DPS remained 599.21 and threat
remained 616.40 per second. Innervate's reported net gain changed from
−62.20 to +2,577.02 mana, of which +1,480.68 was effective after capping;
ordinary regeneration decreased by the corresponding amount. Total mana
and damage were unchanged.

## Already represented

The code already separates banes from curses, resets Immolate's duration and
tick timer on reapplication, removes Curse of Recklessness's target AP bonus,
and omits the obsolete Drain Soul execute/drain-haste rules. Rogue Flawless
Execution, dagger requirements and Backstab positioning are also present.
These findings do not require new buffs or additional damage effects.

## Material uncertainties

The supplied low-level Warrior capture supports speed-normalized landed-swing
rage and equal gains on normal hits, crits and glances. It does not establish
the level-60 or off-hand coefficient. The engine now models outgoing Warrior
rage as speed-normalized, with a provisional Cataclysm-like off-hand rate.
This extrapolation is not a validated level-60 Forever rule. See
[the rage check](check_dispositions.md#rage-evidence).

At the time of this review Hunter pets still used **zero owner-stat
inheritance**. The local post-review pet correction now applies the
Forever-specific wiki rates to Hunter and Warlock pets provisionally, along
with client Focus regeneration; it has not yet been rerun or published in
the rankings. See [the pet formula audit](check_dispositions.md#pet-formula-source-audit).

Divine Spirit's baseline availability after talent removal is assumed rather
than confirmed from a trainer. Seal of Fury and the Priest healing spellbook
are not implemented. These omissions are not free effects in the DPS profiles.

Pet inheritance, several proc-trigger edge cases, periodic snapshot behavior
and other unresolved interactions remain on the [in-game test list](in_game_checks.md).
The [change-history review](history_review.md) records provenance, coverage,
additional availability gaps and matched resource checks.
