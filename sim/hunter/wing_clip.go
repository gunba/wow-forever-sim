package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (hunter *Hunter) getWingClipConfig(rank int) core.SpellConfig {
	spellId := [4]int32{0, 2974, 14267, 14268}[rank]
	baseDamage := [4]float64{0, 5, 25, 50}[rank]
	manaCost := [4]float64{0, 40, 60, 80}[rank]
	level := [4]int{0, 12, 38, 60}[rank]

	return core.SpellConfig{
		SpellCode:     SpellCode_HunterWingClip,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeMHSpecial,
		Flags:         core.SpellFlagMeleeMetrics | core.SpellFlagAPL | core.SpellFlagBinary,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance
		},

		CritDamageBonus:  hunter.mortalShots(),
		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
			if hunter.Env.IsForever() && result.Landed() {
				hunter.slowTarget(sim, target, spell.SpellID, 10*time.Second, []float64{0, .5, .55, .6}[rank])
				// Trait134455 / curve82786 is 7/13/20%, not 7/14/21%.
				// The registered level-60 ability is max rank; the server's
				// additional downrank proc penalty is not exported/assumed.
				if target.PseudoStats.CanBeRooted && sim.Proc([]float64{0, .07, .13, .20}[hunter.Talents.ImprovedWingClip], "Improved Wing Clip") {
					hunter.rootTarget(sim, target, 19229, 5*time.Second)
				}
			}
		},
	}
}

func (hunter *Hunter) registerWingClipSpell() {
	rank := map[int32]int{
		25: 1,
		40: 2,
		50: 3,
		60: 3,
	}[hunter.Level]

	config := hunter.getWingClipConfig(rank)
	hunter.WingClip = hunter.GetOrRegisterSpell(config)
}
