package hunter

import (
	"github.com/wowsims/classic/sim/core"
)

func (hunter *Hunter) registerHuntersMark() {
	if !hunter.Env.IsForever() {
		return
	}
	for i, spellID := range []int32{1130, 14323, 14324, 14325} {
		rank := []struct {
			level     int32
			mana, rap float64
		}{{6, 15, 26}, {22, 30, 59}, {40, 45, 98}, {58, 60, 71}}[i]
		if hunter.Level < rank.level {
			continue
		}
		// Preserve captured effects literally. Rank3's larger RAP than rank4 is a
		// source anomaly, not grounds for extrapolating a new rank4 value.
		auras := hunter.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return core.HuntersMarkAura(target, &hunter.Unit, spellID, rank.rap)
		})
		hunter.HuntersMarkAuras = append(hunter.HuntersMarkAuras, auras)
		hunter.RegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{SpellID: spellID}, Rank: i + 1, RequiredLevel: int(rank.level), SpellSchool: core.SpellSchoolShadow,
			ProcMask: core.ProcMaskEmpty, Flags: core.SpellFlagAPL, ManaCost: core.ManaCostOptions{FlatCost: rank.mana},
			Cast:               core.CastConfig{DefaultCast: core.Cast{GCD: core.GCDDefault}},
			ExtraCastCondition: func(_ *core.Simulation, target *core.Unit) bool { return target.Type == core.EnemyUnit },
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, _ *core.Spell) {
				// A Hunter can mark one enemy, at any learned rank, at a time.
				for _, own := range hunter.HuntersMarkAuras {
					for _, enemy := range sim.Encounter.TargetUnits {
						if aura := own.Get(enemy); aura.IsActive() {
							aura.Deactivate(sim)
						}
					}
				}
				auras.Get(target).Activate(sim)
			},
		})
	}
}

func (hunter *Hunter) hasOwnHuntersMark(target *core.Unit) bool {
	for _, auras := range hunter.HuntersMarkAuras {
		if auras.Get(target).IsActive() {
			return true
		}
	}
	return false
}
