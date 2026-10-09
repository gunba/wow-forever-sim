package main

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// Numeric assumptions accompany each replay, rather than hiding mutable defaults
// behind a general 'provisional' label. Exact provider references and inputs are
// additionally retained in Request. This is reporting, not server validation.
func provisionalModelsForCharacter(character *core.Character, player *proto.Player) map[string]any {
	if character.Env == nil || !character.Env.IsForever() {
		return nil
	}
	models := map[string]any{}
	if character.HasRageBar() {
		models["incomingRage"] = character.ForeverIncomingRageParameters()
		models["incomingRageQualification"] = core.ForeverIncomingRageAssumptions
	}
	for _, id := range []int32{11556, 9898} {
		if spell := character.GetSpell(core.ActionID{SpellID: id}); spell != nil {
			models["demoralizingThreatPerTarget"] = spell.FlatThreatBonus
			models["demoralizingThreatOverride"] = player.ForeverDemoralizingThreat != nil
			models["demoralizingThreatQualification"] = "Nonzero threat is sourced; the default Classic-derived rank amount is not measured Forever behavior. This is base threat before the unit/stance multiplier."
		}
	}
	var paladinOptions *proto.PaladinOptions
	if paladin := player.GetProtectionPaladin(); paladin != nil {
		paladinOptions = paladin.Options
	} else if paladin := player.GetRetributionPaladin(); paladin != nil {
		paladinOptions = paladin.Options
	}
	if paladinOptions != nil {
		lifetime := 30.0
		if paladinOptions.SealOfFuryShieldDurationSeconds != nil {
			lifetime = *paladinOptions.SealOfFuryShieldDurationSeconds
		}
		models["sealOfFuryShield"] = map[string]any{
			"lifetimeSeconds":     lifetime,
			"lifetimeOverride":    paladinOptions.SealOfFuryShieldDurationSeconds != nil,
			"primarySealSelected": paladinOptions.PrimarySeal == proto.PaladinSeal_Fury,
			"replacementPool":     true,
			"qualification":       "ONE shared replacement pool and its independent lifetime are provisional. Zero lifetime disables absorption, not sourced proc damage. A non-Fury primary option does not imply an explicit Fury APL action was cast.",
		}
	}
	if model := player.ForeverRevelationModel; model != nil {
		models["revelation"] = map[string]any{
			"enabled":       model.Enabled,
			"baseChance":    model.BaseChance,
			"critExponent":  model.CritExponent,
			"triggerOnMiss": model.TriggerOnMiss,
			"qualification": "Explicit conditional noncritical-event curve and first eligible outcome-sample charge reservation; not a verified server script. Configuration alone does not equip Revelation.",
		}
	}
	if len(models) == 0 {
		return nil
	}
	return models
}
