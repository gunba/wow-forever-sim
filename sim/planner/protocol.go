// Package planner chooses a character's next action by valuing each available action over a short horizon, from the
// character's own ability model and the current intent, instead of an ordered priority list. The simulator calls it
// through a custom rotation and the SecretProbe console runs it as a local process (tools/forever_planner); both
// build the same decide state. The message types follow the planner contract (PLANNER.md, v1.2).
package planner

import "encoding/json"

// Envelope reads a message's type before decoding the rest.
type Envelope struct {
	Type string `json:"type"`
	ID   *int64 `json:"id,omitempty"`
}

type TalentPick struct {
	Node  int32 `json:"node"`
	Entry int32 `json:"entry"`
	Rank  int32 `json:"rank"`
}

type GearPiece struct {
	Item    int32   `json:"item"`
	Enchant int32   `json:"enchant"`
	Bonus   []int32 `json:"bonus"`
}

type CharacterMsg struct {
	Type       string               `json:"type"`
	Generation int64                `json:"generation"`
	Class      string               `json:"class"`
	Race       string               `json:"race"`
	Level      int32                `json:"level"`
	Talents    []TalentPick         `json:"talents"`
	Spells     []int32              `json:"spells"`
	Gear       map[string]GearPiece `json:"gear"`
}

type AbilityInfo struct {
	Spell    int32  `json:"spell"`
	Name     string `json:"name"`
	Modelled bool   `json:"modelled"`
	Reason   string `json:"reason,omitempty"`
}

type CharacterAnswer struct {
	Type       string        `json:"type"`
	Generation int64         `json:"generation"`
	OK         bool          `json:"ok"`
	Abilities  []AbilityInfo `json:"abilities"`
	Warnings   []string      `json:"warnings"`
	Error      string        `json:"error,omitempty"`
}

// Decide message. Pointers are nullable: nil always means unknown.

type IntentWeights struct {
	Damage     float64 `json:"damage"`
	Threat     float64 `json:"threat"`
	Mitigation float64 `json:"mitigation"`
}

type IntentThreat struct {
	SecureLeadPct *float64 `json:"secure_lead_pct"`
	CeilingPct    *float64 `json:"ceiling_pct"`
}

type IntentContext struct {
	InGroup          *bool    `json:"in_group"`
	GroupSize        *float64 `json:"group_size"`
	AssignedTank     *bool    `json:"assigned_tank"`
	OtherTankPresent *bool    `json:"other_tank_present"`
	HomeStance       *float64 `json:"home_stance"`
}

type Intent struct {
	Mode    string         `json:"mode"`
	Source  string         `json:"source"`
	Context *IntentContext `json:"context,omitempty"`
	Weights *IntentWeights `json:"weights,omitempty"`
	Threat  *IntentThreat  `json:"threat,omitempty"`
}

type SwingTimer struct {
	Remaining *float64 `json:"remaining"`
	Speed     *float64 `json:"speed"`
	Known     bool     `json:"known"`
	AgeMs     *float64 `json:"age_ms"`
}

type AuraState struct {
	Spell     int32    `json:"spell"`
	Remaining *float64 `json:"remaining"`
	Stacks    *float64 `json:"stacks"`
	AgeMs     *float64 `json:"age_ms"`
}

type MeanSD struct {
	Mean *float64 `json:"mean"`
	SD   *float64 `json:"sd"`
}

type PlayerState struct {
	Rage            *float64              `json:"rage"`
	RageMax         *float64              `json:"rage_max"`
	HealthPct       *float64              `json:"health_pct"`
	Stance          *float64              `json:"stance"`
	GCDRemaining    *float64              `json:"gcd_remaining"`
	QueuedSwing     *float64              `json:"queued_swing"`
	Moving          *bool                 `json:"moving"`
	Swing           map[string]SwingTimer `json:"swing"`
	AttackPower     *float64              `json:"attack_power"`
	CritPct         *float64              `json:"crit_pct"`
	Auras           []AuraState           `json:"auras"`
	IncomingPctPerS *MeanSD               `json:"incoming_pct_per_s"`
}

type TTK struct {
	P10 *float64 `json:"p10"`
	P50 *float64 `json:"p50"`
	P90 *float64 `json:"p90"`
}

type ThreatState struct {
	Status          *float64 `json:"status"`
	LeadPct         *float64 `json:"lead_pct"`
	LeadTrendPctPer *float64 `json:"lead_trend_pct_per_s"`
}

type TargetState struct {
	Exists         bool         `json:"exists"`
	HealthPct      *float64     `json:"health_pct"`
	Level          *float64     `json:"level"`
	Classification *string      `json:"classification"`
	TTK            *TTK         `json:"ttk"`
	InMelee        *bool        `json:"in_melee"`
	TargetsPlayer  *bool        `json:"targets_player"`
	Threat         *ThreatState `json:"threat"`
	Auras          []AuraState  `json:"auras"`
}

type EnemyState struct {
	ID            string       `json:"id"`
	HealthPct     *float64     `json:"health_pct"`
	TTK           *TTK         `json:"ttk"`
	InMelee       *bool        `json:"in_melee"`
	TargetsPlayer *bool        `json:"targets_player"`
	Threat        *ThreatState `json:"threat"`
}

type EnemyCount struct {
	Melee   *float64 `json:"melee"`
	Engaged *float64 `json:"engaged"`
	Manual  *float64 `json:"manual"`
}

type Borrow struct {
	Stance    *float64 `json:"stance"`
	Equipment bool     `json:"equipment"`
}

type ActionState struct {
	Spell    int32    `json:"spell"`
	Usable   *bool    `json:"usable"`
	Cooldown *float64 `json:"cooldown"`
	InRange  *bool    `json:"in_range"`
	Borrow   *Borrow  `json:"borrow"`
}

type DecideMsg struct {
	Type       string        `json:"type"`
	ID         int64         `json:"id"`
	Generation int64         `json:"generation"`
	T          float64       `json:"t"`
	BudgetMs   float64       `json:"budget_ms"`
	Intent     Intent        `json:"intent"`
	Player     PlayerState   `json:"player"`
	Target     TargetState   `json:"target"`
	Enemies    []EnemyState  `json:"enemies"`
	EnemyCount EnemyCount    `json:"enemy_count"`
	Actions    []ActionState `json:"actions"`
	ManualOnly []int32       `json:"manual_only"`
}

type Terms struct {
	Damage     float64 `json:"damage"`
	Threat     float64 `json:"threat"`
	Mitigation float64 `json:"mitigation"`
	Terminal   float64 `json:"terminal"`
}

type RankedAction struct {
	Action json.RawMessage `json:"action"` // spell ID or "wait"
	Value  float64         `json:"value"`
	Until  *float64        `json:"until,omitempty"`
	Terms  Terms           `json:"terms"`
}

type ResolvedIntent struct {
	Mode    string        `json:"mode"`
	Weights IntentWeights `json:"weights"`
	Reason  string        `json:"reason"`
}

type DecideAnswer struct {
	Type           string          `json:"type"`
	ID             int64           `json:"id"`
	Best           json.RawMessage `json:"best"`
	Ranked         []RankedAction  `json:"ranked"`
	Margin         float64         `json:"margin"`
	ElapsedMs      float64         `json:"elapsed_ms"`
	Depth          int             `json:"depth"`
	IntentResolved ResolvedIntent  `json:"intent_resolved"`
	Warnings       []string        `json:"warnings"`
	Error          string          `json:"error,omitempty"`
}
