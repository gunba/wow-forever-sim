package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const ConflagrateRanks = 6

func (warlock *Warlock) getConflagrateConfig(rank int, timer *core.Timer) core.SpellConfig {
	// Beta client 1.60.1: Forever adds two ranks below Classic's four (1293817 at 25, 1293818 at 32),
	// which makes Classic's 17962 rank 3, and every rank does about half Classic's damage.
	spellId := [ConflagrateRanks + 1]int32{0, 1293817, 1293818, 17962, 18930, 18931, 18932}[rank]
	baseDamageMin := [ConflagrateRanks + 1]float64{0, 88, 113, 134, 179, 220, 251}[rank]
	baseDamageMax := [ConflagrateRanks + 1]float64{0, 111, 142, 170, 222, 273, 313}[rank]
	manaCost := [ConflagrateRanks + 1]float64{0, 100, 130, 165, 200, 230, 255}[rank]
	level := [ConflagrateRanks + 1]int{0, 25, 32, 40, 48, 54, 60}[rank]

	spCoeff := 0.429

	// 20% per point, so at 5/5 Conflagrate stops consuming Immolate altogether. The demo
	// only showed rank 1 and the tree repeated its 20% at every rank, which is why this was
	// held flat; the beta has since confirmed 80% at rank 4 and 100% at rank 5.
	keepImmolateChance := 0.2 * float64(warlock.Talents.ShadowAndFlame)

	return core.SpellConfig{
		SpellCode:     SpellCode_WarlockConflagrate,
		ActionID:      core.ActionID{SpellID: spellId},
		SpellSchool:   core.SpellSchoolFire,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | WarlockFlagDestruction,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    timer,
				Duration: time.Second * 10,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return warlock.getActiveImmolateSpell(target) != nil
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseDamageMin, baseDamageMax)

			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			immoSpell := warlock.getActiveImmolateSpell(target)
			if immoSpell != nil && !sim.Proc(keepImmolateChance, "Shadow and Flame") {
				immoSpell.Dot(target).Deactivate(sim)
			}
		},
	}
}

func (warlock *Warlock) registerConflagrateSpell() {
	if !warlock.Talents.Conflagrate {
		return
	}

	warlock.Conflagrate = make([]*core.Spell, 0)
	timer := warlock.NewTimer()
	for rank := 1; rank <= ConflagrateRanks; rank++ {
		config := warlock.getConflagrateConfig(rank, timer)

		if config.RequiredLevel <= int(warlock.Level) {
			warlock.Conflagrate = append(warlock.Conflagrate, warlock.GetOrRegisterSpell(config))
		}
	}
}
