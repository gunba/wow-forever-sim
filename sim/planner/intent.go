package planner

// Intent modes and their default weights (PLANNER.md v1.3, "Modes"). The simulator's sweeps tune these.
// group_dps and survive are retired: optimal damage ignores the tank's threat, and there is no survival mode.
const (
	ModeKill = "kill"
	ModeHold = "hold"
	ModeAuto = "auto"
)

var DefaultWeights = map[string]IntentWeights{
	ModeKill: {Damage: 1},
	ModeHold: {Damage: 0.25, Threat: 1, Mitigation: 0.5},
}

var retiredModes = map[string]bool{"group_dps": true, "survive": true}

const DefaultSecureLeadPct = 130

// ResolvedPolicy is the intent the planner optimises: weights plus the threat shape.
type ResolvedPolicy struct {
	Mode          string
	Weights       IntentWeights
	SecureLeadPct float64
	CeilingPct    float64 // 0 = no ceiling; only an explicit ceiling_pct sets one (v1.3 always sends null)
	Reason        string
}

// ResolveIntent turns the message's intent into weights. "auto", a missing mode and the retired modes follow the
// contract's table; the console's own reason is kept for a mode it resolved.
func ResolveIntent(in Intent) ResolvedPolicy {
	mode, reason := in.Mode, in.Reason
	if reason == "" {
		if in.Source == "override" {
			reason = "override"
		} else {
			reason = "resolved by the console"
		}
	}
	switch {
	case mode == ModeAuto || mode == "":
		mode, reason = resolveAuto(in.Context)
	case retiredModes[mode]:
		m, r := resolveAuto(in.Context)
		mode, reason = m, "retired mode "+in.Mode+"; "+r
	}
	if _, ok := DefaultWeights[mode]; !ok {
		mode, reason = ModeKill, "unknown mode "+in.Mode+", using kill"
	}
	w := DefaultWeights[mode]
	if in.Weights != nil && !retiredModes[in.Mode] {
		w = *in.Weights
	}
	p := ResolvedPolicy{Mode: mode, Weights: w, SecureLeadPct: DefaultSecureLeadPct, Reason: reason}
	if in.Threat != nil {
		if in.Threat.SecureLeadPct != nil {
			p.SecureLeadPct = *in.Threat.SecureLeadPct
		}
		if in.Threat.CeilingPct != nil {
			p.CeilingPct = *in.Threat.CeilingPct
		}
	}
	return p
}

// resolveAuto is the v1.3 table: solo -> kill; in a group, the assigned tank -> hold; another assigned tank -> kill;
// no assigned tank -> hold in home Defensive Stance, otherwise kill.
func resolveAuto(c *IntentContext) (string, string) {
	if c == nil || c.InGroup == nil || !*c.InGroup {
		if c == nil || c.InGroup == nil {
			return ModeKill, "no group context, treated as solo"
		}
		return ModeKill, "solo"
	}
	if c.AssignedTank != nil && *c.AssignedTank {
		return ModeHold, "group, assigned tank"
	}
	if c.OtherTankPresent != nil && *c.OtherTankPresent {
		return ModeKill, "group, another member is the assigned tank"
	}
	if c.HomeStance != nil && int(*c.HomeStance) == 2 {
		return ModeHold, "group, no assigned tank, home Defensive Stance"
	}
	return ModeKill, "group, no assigned tank, home stance not Defensive"
}

func (p ResolvedPolicy) Answer() ResolvedIntent {
	return ResolvedIntent{Mode: p.Mode, Weights: p.Weights, Reason: p.Reason}
}
