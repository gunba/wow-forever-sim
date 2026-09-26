package warlock

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

const LifeTapRanks = 6

var LifeTapSpellId = [LifeTapRanks + 1]int32{0, 1454, 1455, 1456, 11687, 11688, 11689}

var LifeTapBaseDamage = [LifeTapRanks + 1]float64{0, 30, 75, 140, 220, 310, 424}

// Dummy effect base + one point per level, capped ten levels after learning.
// The current description applies Spirit and Improved Life Tap to both health
// and mana. Tier 1's separate mana-only bonus is applied by the caller.
func foreverLifeTapConversion(rank int, level int32, spirit float64, improved int32) float64 {
	learned := int32(6 + 10*(rank-1))
	base := [LifeTapRanks + 1]float64{0, 20, 65, 130, 210, 300, 420}[rank]
	base += float64(max(0, min(level, learned+10)-learned))
	return (base + spirit) * (1 + .1*float64(improved))
}

func (warlock *Warlock) getLifeTapBaseConfig(rank int) core.SpellConfig {
	spellId := LifeTapSpellId[rank]
	baseDamage := LifeTapBaseDamage[rank]
	spellCoef := [LifeTapRanks + 1]float64{0, 0.68, 0.8, 0.8, 0.8, 0.8, 0.8}[rank]

	level := [LifeTapRanks + 1]int{0, 6, 16, 26, 36, 46, 56}[rank]

	actionID := core.ActionID{SpellID: spellId}
	petManaShare := 0.5 * float64(warlock.Talents.DemonicEnergies)
	tierManaMultiplier := core.TernaryFloat64(warlock.HasSetBonus(ItemSetDemonheartRaiment, 5), 1.2, 1)

	manaMetrics := warlock.NewManaMetrics(actionID)
	petManaMetrics := make(map[*WarlockPet]*core.ResourceMetrics)
	for _, pet := range warlock.BasePets {
		petManaMetrics[pet] = pet.NewManaMetrics(actionID)
	}

	config := core.SpellConfig{
		ActionID:      actionID,
		SpellSchool:   core.SpellSchoolShadow,
		SpellCode:     SpellCode_WarlockLifeTap,
		DefenseType:   core.DefenseTypeMagic,
		ProcMask:      core.ProcMaskSpellDamage,
		Flags:         core.SpellFlagAPL | core.SpellFlagResetAttackSwing | core.SpellFlagBinary | WarlockFlagAffliction,
		RequiredLevel: level,
		Rank:          rank,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		BonusCoefficient: spellCoef,

		DamageMultiplier: 1 + 0.1*float64(warlock.Talents.ImprovedLifeTap),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcDamage(sim, spell.Unit, baseDamage, spell.OutcomeAlwaysHit)
			// Tier 1 increases mana returned, not the damage/health cost.
			restore := result.Damage * tierManaMultiplier

			if warlock.IsTanking() {
				spell.DealDamage(sim, result)
			}

			warlock.AddMana(sim, restore, manaMetrics)

			if petManaShare > 0 && warlock.ActivePet != nil {
				warlock.ActivePet.AddMana(sim, restore*petManaShare, petManaMetrics[warlock.ActivePet])
			}
		},
	}
	if warlock.Env.IsForever() {
		healthMetrics := warlock.NewHealthMetrics(actionID)
		conversion := func() float64 {
			return foreverLifeTapConversion(rank, warlock.Level, warlock.GetStat(stats.Spirit), warlock.Talents.ImprovedLifeTap)
		}
		config.ProcMask = core.ProcMaskEmpty
		config.Flags &^= core.SpellFlagBinary
		config.BonusCoefficient = 0
		config.DamageMultiplier = 1
		config.ThreatMultiplier = 0
		config.ExtraCastCondition = func(_ *core.Simulation, _ *core.Unit) bool {
			// Non-tanking DPS still uses the established external-healing
			// abstraction rather than adding an unconfigured survival model.
			return !warlock.IsTanking() || warlock.CurrentHealth() > conversion()
		}
		config.ApplyEffects = func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			payment := conversion()
			if warlock.IsTanking() {
				warlock.SpendHealth(sim, payment, healthMetrics)
			}
			restore := payment * tierManaMultiplier
			warlock.AddMana(sim, restore, manaMetrics)
			if petManaShare > 0 && warlock.ActivePet != nil && warlock.ActivePet.IsEnabled() {
				warlock.ActivePet.AddMana(sim, restore*petManaShare, petManaMetrics[warlock.ActivePet])
			}
		}
	}
	return config
}

func (warlock *Warlock) registerLifeTapSpell() {
	warlock.LifeTap = make([]*core.Spell, 0)
	for i := 1; i <= LifeTapRanks; i++ {
		config := warlock.getLifeTapBaseConfig(i)

		if config.RequiredLevel <= int(warlock.Level) {
			warlock.LifeTap = append(warlock.LifeTap, warlock.GetOrRegisterSpell(config))
		}
	}
}
