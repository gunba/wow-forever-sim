package planner

import (
	"fmt"
	"strings"

	"github.com/wowsims/classic/sim/warrior/ranks"
)

// Kind is how an ability acts, which decides how the forward model resolves it.
type Kind int

const (
	KindStrike  Kind = iota // instant yellow melee attack on the target
	KindQueued              // replaces the next main-hand swing (Heroic Strike, Cleave)
	KindAoE                 // hits up to Targets enemies (Thunder Clap, Whirlwind)
	KindDot                 // applies a damage-over-time (Rend)
	KindDebuff              // applies a debuff and may hit (Sunder Armor, Demoralizing Shout)
	KindBuff                // self or party buff (Battle Shout)
	KindRage                // grants rage (Bloodrage)
	KindBlock               // Shield Block
	KindStance              // a stance change
	KindExecute             // Execute: consumes all rage, target below 20%
)

// Table is the attack table an ability rolls on.
type Table int

const (
	TableSpecial  Table = iota // miss, dodge, parry, block, crit
	TableNoAvoid               // miss and crit only (Overpower)
	TableMagic                 // spell miss and crit (Thunder Clap)
	TableNoDamage              // no damage roll; miss/dodge/parry for debuff landing (Sunder)
	TableNone
)

const (
	StanceBattle    = 1
	StanceDefensive = 2
	StanceBerserker = 3
)

// Ability is the planner's model of one trained ability for one character.
type Ability struct {
	Name     ranks.Ability
	Index    int   // position in Model.Abilities (its random stream in rollouts)
	SpellID  int32 // the simulator's action ID (top rank)
	GameID   int32 // the rank the character has, as the game names it
	Rank     *ranks.Rank
	Modelled bool
	Reason   string

	Kind       Kind
	Table      Table
	Stances    uint8 // bit (1 << stance); 0 = any
	Cost       float64
	Refund     float64 // share of the cost returned on a miss, dodge or parry
	CD         float64 // seconds
	GCD        bool
	Cast       float64 // cast time, seconds (Slam)
	OffGCD     bool    // usable while the global cooldown runs
	KeepSwings bool    // a cast that leaves the swing timers running (Improved Slam)

	// Damage: (Flat + APCoef*AP + WeaponMult*weapon swing damage) * DamageMult. Weapon swings use the
	// normalized speed when Normalized is set.
	Flat, APCoef, WeaponMult float64
	Normalized               bool
	DamageMult               float64
	BonusCrit                float64 // crit chance added, percent
	CritMult                 float64
	Targets                  int // enemies hit; 0 or 1 = target only

	// Threat: (damage*ThreatMult + FlatThreat + ThreatAP*AP) before stance multipliers.
	ThreatMult, FlatThreat, ThreatAP float64

	// Effects.
	ArmorPerStack float64 // Sunder Armor
	MaxStacks     int
	Duration      float64 // debuff/buff/dot duration
	APReduction   float64 // Demoralizing Shout
	SlowPct       float64 // Thunder Clap attack speed slow
	BuffAP        float64 // Battle Shout
	TickDamage    float64 // Rend per tick before AP
	TickAPCoef    float64
	Ticks         int
	RageGain      float64 // Bloodrage instant
	RagePerSec    float64 // Bloodrage over Duration
	BlockPct      float64 // Shield Block
	BlockCharges  int
	PerRage       float64 // Execute
	NeedsWindow   int     // windowRevenge, windowOverpower
	NeedsShield   bool
	ToStance      int // stance spells
}

const (
	windowNone = iota
	windowRevenge
	windowOverpower
)

type Weapon struct {
	Present        bool
	Min, Max       float64
	Speed          float64
	NormSpeed      float64
	TwoHand        bool
	DamageMult     float64 // off-hand 0.5 x Dual Wield Specialization
	RagePerLanded  float64 // Forever rage per landed swing (before crit doubling)
	HitBonus       float64 // percent, hand-specific (Furious Precision)
	CritBonus      float64 // percent, hand-specific (Weaponmaster)
	WeaponSkillAdd float64
}

func (w Weapon) Avg() float64 { return (w.Min + w.Max) / 2 }

// Model is a character's ability model: everything the forward model needs that does not change within a fight.
type Model struct {
	Level  int32
	Class  string
	Race   string
	MH, OH Weapon
	Shield bool
	DW     bool
	Health float64 // max health estimate
	Armor  float64

	// Defaults when the decide message leaves them unknown.
	AP, CritPct float64

	HitPct       float64 // melee hit from gear and talents, percent
	ExpertisePct float64
	ArmorIgnore  float64

	PhysMult        float64 // static physical damage multiplier (talents, race)
	WhiteCritMult   float64
	SpecialCritMult float64
	CritScale       float64 // racial critical damage, 1 without one

	// Defensive stats for incoming attacks.
	DodgePct, ParryPct, BlockPct, BlockValue float64
	MissPct                                  float64 // chance an equal-level mob misses

	// Talents that act in the forward model.
	FlurryHaste      float64 // e.g. 0.25 for 5/5; 0 = none
	EnragePct        float64 // physical damage bonus while Enrage is up; 0 = none
	EnrageChance     float64
	DeepWoundsPct    float64 // share of average weapon damage over 12 s per crit
	UWChance         float64 // Unbridled Wrath chance per landed swing
	AngerManagement  bool
	TacticalMastery  float64 // rage kept through a stance change
	DefStanceThreat  float64 // Defensive Stance threat multiplier (1.3, more with Defiance and a shield)
	OtherStanceMult  float64 // Battle/Berserker threat multiplier
	ShieldSpecRage   float64 // rage per block
	IncomingRagePerD float64 // rage per point of damage taken under a manual (provisional) incoming rage model
	// IncomingRageManual: the character overrides Forever's measured rage-from-damage-taken rule.
	IncomingRageManual bool

	Abilities []*Ability
	byName    map[ranks.Ability]*Ability
	byGameID  map[int32]*Ability
	Warnings  []string
}

func (m *Model) Get(a ranks.Ability) *Ability {
	if m == nil || m.byName == nil {
		return nil
	}
	return m.byName[a]
}

// ByGameID maps a spell ID from the game (any rank) to the character's ability.
func (m *Model) ByGameID(id int32) *Ability {
	if ab, ok := m.byGameID[id]; ok {
		return ab
	}
	if name, _ := ranks.ByID(id); name != "" {
		return m.byName[name]
	}
	return nil
}

func (m *Model) index() {
	m.byName = map[ranks.Ability]*Ability{}
	m.byGameID = map[int32]*Ability{}
	for i, ab := range m.Abilities {
		ab.Index = i
		m.byName[ab.Name] = ab
		m.byGameID[ab.GameID] = ab
		m.byGameID[ab.SpellID] = ab
		for _, r := range ranks.All(ab.Name) {
			m.byGameID[r.ID] = ab
		}
	}
}

// Describe is a readable summary of the model, for checking a character build.
func (m *Model) Describe() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s %s level %d: AP %.0f, crit %.2f%%, hit %.1f%%, expertise %.1f%%, health %.0f, armor %.0f\n", m.Race, m.Class, m.Level, m.AP, m.CritPct, m.HitPct, m.ExpertisePct, m.Health, m.Armor)
	fmt.Fprintf(&b, "main hand %.0f-%.0f @ %.2f (2H %v), rage/landed %.2f; off hand %v %.0f-%.0f @ %.2f x%.3f, rage/landed %.2f\n",
		m.MH.Min, m.MH.Max, m.MH.Speed, m.MH.TwoHand, m.MH.RagePerLanded, m.DW, m.OH.Min, m.OH.Max, m.OH.Speed, m.OH.DamageMult, m.OH.RagePerLanded)
	fmt.Fprintf(&b, "phys mult %.3f, crit mult white %.2f special %.2f, dodge %.2f parry %.2f block %.2f (shield %v, value %.0f)\n",
		m.PhysMult, m.WhiteCritMult, m.SpecialCritMult, m.DodgePct, m.ParryPct, m.BlockPct, m.Shield, m.BlockValue)
	fmt.Fprintf(&b, "flurry %.2f, enrage %.2f, deep wounds %.2f, UW %.2f, anger management %v, tactical mastery %.0f, def threat %.3f, incoming rage/dmg %.4f\n",
		m.FlurryHaste, m.EnragePct, m.DeepWoundsPct, m.UWChance, m.AngerManagement, m.TacticalMastery, m.DefStanceThreat, m.IncomingRagePerD)
	for _, ab := range m.Abilities {
		rank := int32(0)
		if ab.Rank != nil {
			rank = ab.Rank.Rank
		}
		if !ab.Modelled {
			fmt.Fprintf(&b, "  %-18s rank %d  game %d  NOT MODELLED: %s\n", ab.Name, rank, ab.GameID, ab.Reason)
			continue
		}
		fmt.Fprintf(&b, "  %-18s rank %d  game %-6d cost %4.1f cd %4.1f gcd %-5v stances %03b flat %6.1f ap %.3f weap %.0f norm %v dmgx %.3f threatx %.2f flatthreat %.1f threatAP %.4f\n",
			ab.Name, rank, ab.GameID, ab.Cost, ab.CD, ab.GCD, ab.Stances>>1, ab.Flat, ab.APCoef, ab.WeaponMult, ab.Normalized, ab.DamageMult, ab.ThreatMult, ab.FlatThreat, ab.ThreatAP)
	}
	return b.String()
}
