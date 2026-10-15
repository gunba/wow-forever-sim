package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

func (warrior *Warrior) registerThunderClapSpell() {
	rank := warrior.rank(ranks.ThunderClap)
	if rank == nil {
		return
	}
	spellID := int32(11581)
	baseDamage := rank.Find(ranks.EffectSchoolDamage, 0).Low
	has5pcConq := warrior.HasSetBonus(ItemSetConquerorsBattleGear, 5)
	// Forever doubles the slow to 20% and moves the cooldown from 4 to 6 sec.
	attackSpeedReduction := core.TernaryInt32(has5pcConq, 25, 20)
	// Forever lets Thunder Clap be used in Defensive Stance as well.
	stanceMask := BattleStance | DefensiveStance

	warrior.ThunderClapAuras = warrior.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.ThunderClapAura(target, spellID, attackSpeedReduction)
	})

	results := make([]*core.SpellResult, min(4, warrior.Env.GetNumTargets()))

	warrior.ThunderClap = warrior.RegisterSpell(stanceMask, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost: 20 - []float64{0, 2, 4, 6}[warrior.Talents.ImprovedThunderClap],
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: time.Second * 6,
			},
		},

		CritDamageBonus: warrior.impale(),

		DamageMultiplier: core.TernaryFloat64(has5pcConq, 1.5, 1),
		// The deep dive gives Thunder Clap "significant damage scaling from Attack Power" and
		// threat "rebalanced (reduced)" for Defensive Stance. Testers measured about 7% AP
		// (0.068) and threat equal to damage before stance modifiers.
		ThreatMultiplier: core.TernaryFloat64(warrior.Env.IsForever(), 1, 2.5),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := baseDamage
			if sim.IsForever() {
				damage += ranks.ThunderClapAPCoefficient * spell.MeleeAttackPower(target)
			}
			for idx := range results {
				results[idx] = spell.CalcDamage(sim, target, damage, spell.OutcomeMagicHitAndCrit)
				target = sim.Environment.NextTargetUnit(target)
			}

			for _, result := range results {
				spell.DealDamage(sim, result)
				if result.Landed() {
					warrior.ThunderClapAuras.Get(result.Target).Activate(sim)
				}
			}
		},

		RelatedAuras: []core.AuraArray{warrior.ThunderClapAuras},
	})
}
