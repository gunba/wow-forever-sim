package planner

// Intent modes and their default weights (PLANNER.md, "Intent modes"). The simulator's sweeps tune these.
const (
	ModeKill     = "kill"
	ModeHold     = "hold"
	ModeGroupDPS = "group_dps"
	ModeSurvive  = "survive"
	ModeAuto     = "auto"
)

var DefaultWeights = map[string]IntentWeights{
	ModeKill:     {Damage: 1},
	ModeHold:     {Damage: 0.25, Threat: 1, Mitigation: 0.5},
	ModeGroupDPS: {Damage: 1},
	ModeSurvive:  {Damage: 0.25, Mitigation: 1},
}

const (
	DefaultSecureLeadPct = 130
	DefaultMeleeCeiling  = 110
)

// ResolvedPolicy is the intent the planner optimises: weights plus the threat shape.
type ResolvedPolicy struct {
	Mode          string
	Weights       IntentWeights
	SecureLeadPct float64
	CeilingPct    float64 // 0 = no ceiling
	Reason        string
}

// ResolveIntent turns the message's intent into weights. "auto" follows the contract's table, applied in order:
// override, not in a group -> kill, assigned tank -> hold or group_dps, otherwise home Defensive Stance -> hold.
func ResolveIntent(in Intent) ResolvedPolicy {
	mode, reason := in.Mode, ""
	if in.Source == "" || in.Source == "auto" {
		reason = "resolved by the console"
	} else {
		reason = "override"
	}
	if mode == ModeAuto || mode == "" {
		mode, reason = resolveAuto(in.Context)
	}
	if _, ok := DefaultWeights[mode]; !ok {
		mode, reason = ModeKill, "unknown mode "+in.Mode+", using kill"
	}
	w := DefaultWeights[mode]
	if in.Weights != nil {
		w = *in.Weights
	}
	p := ResolvedPolicy{Mode: mode, Weights: w, SecureLeadPct: DefaultSecureLeadPct, Reason: reason}
	if mode == ModeGroupDPS {
		p.CeilingPct = DefaultMeleeCeiling
	}
	if in.Threat != nil {
		if in.Threat.SecureLeadPct != nil {
			p.SecureLeadPct = *in.Threat.SecureLeadPct
		}
		if in.Threat.CeilingPct != nil {
			p.CeilingPct = *in.Threat.CeilingPct
		} else if mode != ModeGroupDPS {
			p.CeilingPct = 0
		}
	}
	return p
}

func resolveAuto(c *IntentContext) (string, string) {
	if c == nil || c.InGroup == nil || !*c.InGroup {
		if c == nil || c.InGroup == nil {
			return ModeKill, "no group context, treated as solo"
		}
		return ModeKill, "not in a group"
	}
	if c.AssignedTank != nil && *c.AssignedTank {
		return ModeHold, "assigned tank"
	}
	if c.OtherTankPresent != nil && *c.OtherTankPresent {
		return ModeGroupDPS, "another member is the assigned tank"
	}
	if c.HomeStance != nil && int(*c.HomeStance) == 2 {
		return ModeHold, "home Defensive Stance, no assigned tank"
	}
	return ModeGroupDPS, "no assigned tank, home stance is not Defensive"
}

func (p ResolvedPolicy) Answer() ResolvedIntent {
	return ResolvedIntent{Mode: p.Mode, Weights: p.Weights, Reason: p.Reason}
}
