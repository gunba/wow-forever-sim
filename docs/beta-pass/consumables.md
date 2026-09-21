# Consumable checks

These checks use Forever client 1.60.1.69893 and the linked Forever tooltips.
They cover the listed effects, not every consumable in the editor.

## Food corrections

| Item | Forever effect | Client chain |
|---|---|---|
| [Grilled Squid](https://www.wowhead.com/forever/item=13928) | 1 percentage point of critical-strike chance, physical and spell; not 10 Agility | 1249522 → 1249523, aura 290 |
| [Runn Tum Tuber Surprise](https://www.wowhead.com/forever/item=18254) | 15 Intellect, not 10 | 1248396 → 1248421 |
| [Nightfin Soup](https://www.wowhead.com/forever/item=13931) | 22 spell damage, not 8 MP5 | 1249513 → 1249520 |

The eating spell supplies the amount to the triggered Well Fed aura. Its dummy
one-point base value is not the food's final stat amount. Nightfin's damage
effect covers the six magic schools and does not add healing power.

[Smoked Desert Dumplings](https://www.wowhead.com/forever/item=20452) still gives
20 Strength, and [Blessed Sunfruit](https://www.wowhead.com/forever/item=13810)
still gives 10 Strength.

## Mana consumables

[Mageblood Elixir](https://www.wowhead.com/forever/item=20007) still supplies
12 MP5. [Flask of Distilled Wisdom](https://www.wowhead.com/forever/item=13511)
supplies 2,000 maximum mana, replacing rather than stacking with another flask.

The [Major Mana Potion](https://www.wowhead.com/forever/item=13444) tooltip
displays 1,800 mana and the [Demonic Rune](https://www.wowhead.com/forever/item=12662)
tooltip displays 1,200 mana / 800 health. Those are central values:
client spells 17531 and 16666 still have **Variance 0.5**. The retained ranges
are therefore 1,350–2,250 mana, and 900–1,500 mana / 600–1,000 health.
The item tooltips specify two-minute cooldowns; the spell-table category value
alone is not the complete item cooldown.

Hunter baselines already select both mana consumables. The simulator's
`secondsOomAvg` counts periods of failed mana-cost checks. It is not a duration
with zero damage: ranged auto-attacks and pet attacks continue.
