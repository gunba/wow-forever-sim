package paladin

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Consecration is baseline in the beta client (1.60.1.69893): every paladin trains all five ranks, and
// the Forever tree builds on top of it through Consecrated Ground and Holy Conduit.
//
// Each tick casts a separate damage spell (1280345-1280349) with two parts: a flat amount every enemy
// in the area takes, and a larger amount with the spell power coefficient that only the first
// $s3 = 4 enemies take. Classic's 48 a tick at 0.042 becomes 12 + 27 at 0.095 on the capped part.
func (paladin *Paladin) registerConsecration() {
	const cappedTargets = 4

	ranks := []struct {
		level    int32
		manaCost float64
		damage   float64 // every enemy, per tick
		capped   float64 // first 4 enemies, per tick, scales with spell power
	}{
		{level: 20, manaCost: 135, damage: 2, capped: 4},
		{level: 30, manaCost: 235, damage: 3, capped: 7},
		{level: 40, manaCost: 320, damage: 6, capped: 11},
		{level: 50, manaCost: 435, damage: 8, capped: 20},
		{level: 60, manaCost: 565, damage: 12, capped: 27},
	}

	cd := core.Cooldown{
		Timer:    paladin.NewTimer(),
		Duration: time.Second * 8,
	}

	for i, rank := range ranks {
		rank := rank
		spellID := []int32{26573, 20116, 20922, 20923, 20924}[i]
		if paladin.Level < rank.level {
			break
		}

		paladin.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: spellID},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,
			Flags:       core.SpellFlagPureDot | core.SpellFlagAPL,

			RequiredLevel: int(rank.level),
			Rank:          i + 1,

			SpellCode: SpellCode_PaladinConsecration,
			ManaCost: core.ManaCostOptions{
				FlatCost:   rank.manaCost,
				Multiplier: paladin.benediction() * paladin.holyConduit() / 100,
			},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
				CD: cd,
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			Dot: core.DotConfig{
				IsAOE: true,
				Aura: core.Aura{
					Label: "Consecration" + paladin.Label + strconv.Itoa(i+1),
				},
				NumberOfTicks: 8,
				TickLength:    time.Second * 1,

				BonusCoefficient: 0.042, // 0.095 until 70425 (October 15 notes)

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, rank.damage+rank.capped, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					// Consecration can miss, showing up as either a resist in logs or a
					// silent failure (missing damage tick).
					// ponytail: "first 4 to enter" is read as the first 4 targets in the encounter.
					for j, aoeTarget := range sim.Encounter.TargetUnits {
						if j < cappedTargets {
							dot.CalcAndDealPeriodicSnapshotDamage(sim, aoeTarget, dot.OutcomeMagicHitAndTick)
						} else {
							// The flat part alone, which carries no coefficient; the spell's own
							// BonusCoefficient stays zero so this does not pick one up.
							dot.Spell.CalcAndDealDamage(sim, aoeTarget, rank.damage, dot.OutcomeMagicHitAndTick)
						}
					}
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.AOEDot().Apply(sim)
			},
		})
	}
}
