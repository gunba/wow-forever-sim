package priest

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const ShadowWordDeathRanks = 4

// Client 1.60.1.69893: SkillLineAbility, SpellLevels, SpellEffect,
// SpellPower and SpellCooldowns. The parent tooltip specifies 10% maximum
// health backlash, not the old child tooltip's damage-reflection formula.
var ShadowWordDeathSpellIDs = [...]int32{0, 1309595, 1309633, 1309635, 1309636}
var ShadowWordDeathLevels = [...]int32{0, 32, 40, 48, 56}

func (priest *Priest) registerShadowWordDeath() {
	priest.ShadowWordDeath = make([]*core.Spell, ShadowWordDeathRanks+1)
	if priest.Level < ShadowWordDeathLevels[1] {
		return
	}
	backlash := priest.RegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: 1309598},
		SpellSchool:      core.SpellSchoolShadow,
		DefenseType:      core.DefenseTypeMagic,
		ProcMask:         core.ProcMaskEmpty,
		Flags:            core.SpellFlagPassiveSpell | core.SpellFlagNoOnCastComplete | core.SpellFlagIgnoreModifiers | core.SpellFlagIgnoreResists,
		DamageMultiplier: 1,
		ThreatMultiplier: 0,
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			// Self-target damage is recorded as damage taken, not DPS. Health
			// depletion follows the core's existing tank/healing-model rules.
			spell.CalcAndDealDamage(sim, &priest.Unit, .10*priest.MaxHealth(), spell.OutcomeAlwaysHit)
		},
	})
	timer := priest.NewTimer()
	for rank := 1; rank <= ShadowWordDeathRanks; rank++ {
		if priest.Level >= ShadowWordDeathLevels[rank] {
			priest.ShadowWordDeath[rank] = priest.RegisterSpell(priest.shadowWordDeathConfig(rank, timer, backlash))
		}
	}
}

func (priest *Priest) shadowWordDeathConfig(rank int, timer *core.Timer, backlash *core.Spell) core.SpellConfig {
	level := ShadowWordDeathLevels[rank]
	base := [...]float64{0, 295, 370, 403, 448}[rank]
	growth := [...]float64{0, 1.5, 1.9, 2.2, 2.5}[rank]
	mean := base + growth*float64(min(priest.Level-level, 5))
	halfRange := base * .0605 / 2
	cost := [...]float64{0, 175, 205, 250, 340}[rank]
	return core.SpellConfig{
		SpellCode:     SpellCode_PriestShadowWordDeath,
		ActionID:      core.ActionID{SpellID: ShadowWordDeathSpellIDs[rank]},
		SpellSchool:   core.SpellSchoolShadow,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         SpellFlagPriest | core.SpellFlagAPL,
		RequiredLevel: int(level),
		Rank:          rank,
		ManaCost:      core.ManaCostOptions{FlatCost: cost},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			CD:          core.Cooldown{Timer: timer, Duration: 15 * time.Second},
		},
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: .429,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcDamage(sim, target, sim.Roll(mean-halfRange, mean+halfRange), spell.OutcomeMagicHitAndCrit)
			landed := result.Landed()
			// Duration encounters have no individual target deaths. A killing
			// blow is observable in a single-target health-based encounter.
			kills := len(sim.Encounter.Targets) == 1 && sim.Encounter.EndFightAtHealth > 0 &&
				sim.Encounter.DamageTaken+result.Damage >= sim.Encounter.EndFightAtHealth
			if landed {
				priest.AddShadowWeavingStack(sim)
			}
			spell.DealDamage(sim, result)
			if landed && !kills {
				backlash.Cast(sim, &priest.Unit)
			}
		},
		ExpectedInitialDamage: func(sim *core.Simulation, target *core.Unit, spell *core.Spell, _ bool) *core.SpellResult {
			return spell.CalcDamage(sim, target, mean, spell.OutcomeExpectedMagicHitAndCrit)
		},
	}
}
