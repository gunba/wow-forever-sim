package core

// Global configuration settings

// Whether to include spell ranks taught only by Ahn'Qiraj books (Fireball 12, Frostbolt 11, Arcane
// Missiles 8, Heroic Strike 9, Battle Shout 7, Blessing of Might 7, Backstab 9, Eviscerate 9, Shadow
// Bolt 10 and the like). Forever's trainers stop one rank short ("Rank 11 is last trainer taught",
// wowsims/forever #538; Book of Ferocious Bite V 24101), and the books need raids, which open no
// earlier than 9 December 2026. Applies to the player's own ranks and the raid buffs alike.
const IncludeAQ = false
