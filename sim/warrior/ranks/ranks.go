// Package ranks holds every rank of every trained Warrior ability, generated from the client spell tables, and the
// server-side numbers the client does not carry. The simulator's spells and the rotation planner both read it, so the
// two use the same numbers at every level.
package ranks

import "sort"

type Ability string

const (
	BattleShout       Ability = "BattleShout"
	BattleStance      Ability = "BattleStance"
	BerserkerRage     Ability = "BerserkerRage"
	BerserkerStance   Ability = "BerserkerStance"
	Bloodrage         Ability = "Bloodrage"
	Bloodthirst       Ability = "Bloodthirst"
	ChallengingShout  Ability = "ChallengingShout"
	Charge            Ability = "Charge"
	Cleave            Ability = "Cleave"
	ConcussionBlow    Ability = "ConcussionBlow"
	DeathWish         Ability = "DeathWish"
	DefensiveStance   Ability = "DefensiveStance"
	DemoralizingShout Ability = "DemoralizingShout"
	Disarm            Ability = "Disarm"
	Execute           Ability = "Execute"
	Hamstring         Ability = "Hamstring"
	HeroicStrike      Ability = "HeroicStrike"
	Intercept         Ability = "Intercept"
	IntimidatingShout Ability = "IntimidatingShout"
	LastStand         Ability = "LastStand"
	MockingBlow       Ability = "MockingBlow"
	MortalStrike      Ability = "MortalStrike"
	Overpower         Ability = "Overpower"
	PiercingHowl      Ability = "PiercingHowl"
	Pummel            Ability = "Pummel"
	Recklessness      Ability = "Recklessness"
	Rend              Ability = "Rend"
	Retaliation       Ability = "Retaliation"
	Revenge           Ability = "Revenge"
	ShieldBash        Ability = "ShieldBash"
	ShieldBlock       Ability = "ShieldBlock"
	ShieldSlam        Ability = "ShieldSlam"
	ShieldWall        Ability = "ShieldWall"
	Slam              Ability = "Slam"
	SpearingStrike    Ability = "SpearingStrike"
	SunderArmor       Ability = "SunderArmor"
	SweepingStrikes   Ability = "SweepingStrikes"
	Taunt             Ability = "Taunt"
	ThunderClap       Ability = "ThunderClap"
	VictoryRush       Ability = "VictoryRush"
	Whirlwind         Ability = "Whirlwind"
)

// Effect is one client spell effect of a rank. Low and High are the effect's base points range.
type Effect struct {
	Effect, Aura  int32
	Low, High     float64
	Misc          int32
	PeriodMs      int32
	Trigger       int32
	APCoefficient float64
}

type Rank struct {
	ID         int32
	Rank       int32
	Level      int32
	Cost       float64 // rage
	CooldownMs int32
	GCDMs      int32
	DurationMs int32
	Effects    []Effect
}

// Find returns the first effect of the given type and aura, or a zero effect.
func (r *Rank) Find(effect, aura int32) Effect {
	for _, e := range r.Effects {
		if e.Effect == effect && (aura == 0 || e.Aura == aura) {
			return e
		}
	}
	return Effect{}
}

// Client effect and aura types used by the Warrior abilities.
const (
	EffectSchoolDamage       = 2
	EffectDummy              = 3
	EffectApplyAura          = 6
	EffectWeaponDamageBonus  = 17 // flat damage added to a weapon strike
	EffectThreat             = 63
	EffectNormalizedWeapon   = 121
	AuraPeriodicDamage       = 3
	AuraModAttackPower       = 99
	AuraModResistance        = 22
	AuraModArmorPercentOrNeg = 674 // Forever's armor reduction aura on Sunder Armor
)

// All returns every rank of an ability, lowest first.
func All(a Ability) []Rank { return generated[a] }

// AtLevel is the highest rank a character of this level can have trained, or nil.
func AtLevel(a Ability, level int32) *Rank {
	rs := generated[a]
	for i := len(rs) - 1; i >= 0; i-- {
		if rs[i].Level <= level {
			return &rs[i]
		}
	}
	return nil
}

// Known is the highest rank among known spell IDs, or nil when none is known.
func Known(a Ability, known map[int32]bool) *Rank {
	rs := generated[a]
	for i := len(rs) - 1; i >= 0; i-- {
		if known[rs[i].ID] {
			return &rs[i]
		}
	}
	return nil
}

// ByID finds the ability and rank of a client spell ID.
func ByID(id int32) (Ability, *Rank) {
	for a, rs := range generated {
		for i := range rs {
			if rs[i].ID == id {
				return a, &rs[i]
			}
		}
	}
	return "", nil
}

// Abilities lists every ability in the table, sorted.
func Abilities() []Ability {
	out := make([]Ability, 0, len(generated))
	for a := range generated {
		out = append(out, a)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// Selector picks the rank of each ability a character has: from a known-spell list when one is given, otherwise the
// highest rank trainable at the character's level.
type Selector struct {
	Level int32
	Known map[int32]bool
}

func (s Selector) Rank(a Ability) *Rank {
	if s.Known != nil {
		return Known(a, s.Known)
	}
	return AtLevel(a, s.Level)
}
