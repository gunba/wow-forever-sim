package paladin

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

type lightsVigilMark struct {
	aura    *core.Aura
	damage  *core.Spell
	refund  float64
	metrics *core.ResourceMetrics
	trigger *bool
}

// The enemy branch is sourced by 1310911/1311590/1311595 and their children.
// Party healing is not represented by this damage-only registration.
func (paladin *Paladin) registerLightsVigil() {
	if !paladin.Env.IsForever() || !paladin.Talents.LightsVigil {
		return
	}
	ranks := []struct {
		level          int
		cost, min, max float64
	}{
		{40, 730, 175, 189},
		{50, 1000, 268, 288},
		{60, 1340, 380, 410},
	}
	timer := paladin.NewTimer() // Client cooldown category 2576.
	for i, spellID := range []int32{1310911, 1311590, 1311595} {
		rank := ranks[i]
		markID := []int32{1310910, 1311594, 1311599}[i]
		damageID := []int32{1310914, 1311592, 1311598}[i]
		if int(paladin.Level) < rank.level {
			continue
		}
		actionID := core.ActionID{SpellID: spellID}
		metrics := paladin.NewManaMetrics(actionID)
		triggered := false
		damage := paladin.RegisterSpell(core.SpellConfig{
			ActionID:           core.ActionID{SpellID: damageID},
			SpellSchool:        core.SpellSchoolHoly,
			DefenseType:        core.DefenseTypeMagic,
			ProcMask:           core.ProcMaskSpellDamage,
			Flags:              core.SpellFlagPassiveSpell | core.SpellFlagNoOnCastComplete | core.SpellFlagMeleeMetrics,
			RequiredLevel:      rank.level,
			Rank:               i + 1,
			DamageMultiplier:   1,
			ThreatMultiplier:   1,
			BonusCoefficient:   0.429,
			ExtraCastCondition: func(_ *core.Simulation, _ *core.Unit) bool { return triggered },
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(rank.min, rank.max), spell.OutcomeMagicHitAndCrit)
			},
		})
		marks := make([]lightsVigilMark, len(paladin.Env.Encounter.TargetUnits))
		for j, target := range paladin.Env.Encounter.TargetUnits {
			mark := &marks[j]
			mark.damage, mark.metrics, mark.trigger = damage, metrics, &triggered
			mark.aura = target.GetOrRegisterAura(core.Aura{
				Label:    fmt.Sprintf("Light's Vigil-%d-%d", paladin.Index, i+1),
				ActionID: core.ActionID{SpellID: markID, Tag: paladin.Index},
				Duration: 30 * time.Second,
				OnExpire: func(_ *core.Aura, _ *core.Simulation) {
					if paladin.activeVigil == mark {
						paladin.activeVigil = nil
					}
				},
			})
		}
		paladin.RegisterSpell(core.SpellConfig{
			ActionID:      actionID,
			SpellCode:     SpellCode_PaladinLightsVigil,
			SpellSchool:   core.SpellSchoolHoly,
			DefenseType:   core.DefenseTypeMagic,
			ProcMask:      core.ProcMaskEmpty,
			Flags:         core.SpellFlagAPL | core.SpellFlagResetAttackSwing,
			RequiredLevel: rank.level,
			Rank:          i + 1,
			ManaCost:      core.ManaCostOptions{FlatCost: rank.cost},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{GCD: core.GCDDefault, CastTime: 1500 * time.Millisecond},
				CD:          core.Cooldown{Timer: timer, Duration: 6 * time.Second},
			},
			ExtraCastCondition: func(_ *core.Simulation, target *core.Unit) bool {
				return target != nil && target.Type == core.EnemyUnit && (!target.HasHealthBar() || target.IsActive())
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				if !spell.CalcOutcome(sim, target, spell.OutcomeMagicHit).Landed() {
					return
				}
				if paladin.activeVigil != nil {
					paladin.activeVigil.aura.Deactivate(sim)
				}
				mark := &marks[target.Index]
				mark.refund = 0.75 * spell.CurCast.Cost
				mark.aura.Activate(sim)
				paladin.activeVigil = mark
			},
		})
	}
}

func (paladin *Paladin) consumeLightsVigil(sim *core.Simulation, target *core.Unit, shock *core.Spell) {
	mark := paladin.activeVigil
	if mark == nil || !mark.aura.IsActive() || mark.aura.Unit != target {
		return
	}
	mark.aura.Deactivate(sim)
	*mark.trigger = true
	mark.damage.Cast(sim, target)
	*mark.trigger = false
	paladin.AddMana(sim, mark.refund, mark.metrics)
	// CanCast still requires any pre-existing Shock cooldown to have expired.
	// Only the new cooldown started by this consuming cast is removed.
	shock.CD.Timer.Reset()
}
