package ranks

// Server-side numbers the client tables do not carry. Each is sourced or marked provisional.

// HeroicStrikeThreat is the flat threat of a Heroic Strike rank: 2.54 per spell level, measured with an addon reading
// the threat table against the combat log (wowsims/forever #264); rank 4 (level 24) gives 61, as the published Warrior
// threat sheet measured at level 30. Rank 8 (level 56) 142.2.
func HeroicStrikeThreat(r *Rank) float64 { return 2.54 * float64(r.Level) }

// CleaveThreat is the flat threat of a Cleave rank, 1.53 per spell level (rank 2, level 30: 46 in the threat sheet;
// rank 5: 91.8). It is split between the targets the cleave lands on (wowsims/forever #264, #385).
func CleaveThreat(r *Rank) float64 { return 1.53 * float64(r.Level) }

// SunderThreatAP is the attack power coefficient of Sunder Armor's threat: (flat + 0.05 AP) x stance at every rank,
// measured on rank 1 across 132-180 attack power in both stances (wowsims/forever #264).
func SunderThreatAP(_ *Rank) float64 { return 0.05 }

// ExecuteDamagePerRage is Classic's extra damage per Rage converted: 3 per rank.
func ExecuteDamagePerRage(r *Rank) float64 { return 3 * float64(r.Rank) }

// DemoralizingThreatScale scales the client's flat threat of Demoralizing Shout to the measured value: rank 2 lists
// 19 and was measured at 22.7 before modifiers (threat sheet). Provisional for other ranks.
const DemoralizingThreatScale = 22.74 / 19

// RevengeAPCoefficient: measured at rank 2 (threat sheet fit); used for every rank.
const RevengeAPCoefficient = 0.28

// RendAPCoefficient: 2% of attack power per tick, measured at rank 2 (Fight Club logs); used for every rank.
const RendAPCoefficient = 0.02

// ThunderClapAPCoefficient: about 7% of attack power (testers); used for every rank.
const ThunderClapAPCoefficient = 0.07
