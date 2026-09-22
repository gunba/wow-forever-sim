package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

type Stance uint8

const (
	BattleStance Stance = 1 << iota
	DefensiveStance
	BerserkerStance

	AnyStance = BattleStance | DefensiveStance | BerserkerStance
)

func (stance Stance) Matches(other Stance) bool {
	return (stance & other) != 0
}

var StanceCodes = []int32{SpellCode_WarriorStanceBattle, SpellCode_WarriorStanceDefensive, SpellCode_WarriorStanceBerserker}

const stanceEffectCategory = "Stance"

func (warrior *Warrior) StanceMatches(other Stance) bool {
	return warrior.Stance.Matches(other)
}

func (warrior *Warrior) makeStanceSpell(stance Stance, aura *core.Aura, stanceCD *core.Timer) *WarriorSpell {
	spellCode := map[Stance]int32{
		BattleStance:    SpellCode_WarriorStanceBattle,
		DefensiveStance: SpellCode_WarriorStanceDefensive,
		BerserkerStance: SpellCode_WarriorStanceBerserker,
	}[stance]
	actionID := aura.ActionID
	// Tactical Mastery is a baseline passive in the Arms tab under Forever, not a talent,
	// and keeps 10 Rage on its own; Improved Tactical Mastery adds 3 per point on top.
	// Read off Soda's Warrior video: the level 38 spellbook lists it as a passive, and the
	// note with it reads "keeps up to 10 Rage through a stance change by itself; the Arms
	// talent adds up to 15 more".
	baselineRetainedRage := core.TernaryFloat64(warrior.Env.IsForever(), 10, 0)
	maxRetainedRage := baselineRetainedRage + 3*float64(warrior.Talents.ImprovedTacticalMastery)
	rageMetrics := warrior.NewRageMetrics(actionID)

	stanceSpell := warrior.RegisterSpell(AnyStance, core.SpellConfig{
		SpellCode: spellCode,
		ActionID:  actionID,
		Flags:     core.SpellFlagAPL,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    stanceCD,
				Duration: time.Second,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return !warrior.StanceMatches(stance)
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			if warrior.CurrentRage() > maxRetainedRage {
				warrior.SpendRage(sim, warrior.CurrentRage()-maxRetainedRage, rageMetrics)
			}

			if warrior.WarriorInputs.StanceSnapshot {
				// Delayed, so same-GCD casts are affected by the current aura.
				//  Alternatively, those casts could just (artificially) happen before the stance change.
				core.StartDelayedAction(sim, core.DelayedActionOptions{
					DoAt:     sim.CurrentTime + 10*time.Millisecond,
					OnAction: aura.Activate,
				})
			} else {
				aura.Activate(sim)
			}

			warrior.PreviousStance = warrior.Stance
			warrior.Stance = stance
		},
	})

	warrior.Stances = append(warrior.Stances, stanceSpell)

	return stanceSpell
}

func (warrior *Warrior) registerBattleStanceAura() {
	warrior.BattleStanceAura = warrior.RegisterAura(core.Aura{
		Label:    "Battle Stance",
		ActionID: core.ActionID{SpellID: 2457},
		Duration: core.NeverExpires,
	})
	warrior.BattleStanceAura.NewExclusiveEffect(stanceEffectCategory, true, core.ExclusiveEffect{
		OnGain: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.PseudoStats.ThreatMultiplier *= 0.8
		},
		OnExpire: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.PseudoStats.ThreatMultiplier /= 0.8
		},
	})
}

func (warrior *Warrior) registerDefensiveStanceAura() {
	threatMultiplier := func() float64 {
		defiance := core.TernaryFloat64(warrior.PseudoStats.CanBlock, 0.05*float64(warrior.Talents.Defiance), 0)
		return 1.3 * (1 + defiance)
	}
	warrior.defensiveStanceThreatMultiplier = threatMultiplier()

	warrior.DefensiveStanceAura = warrior.RegisterAura(core.Aura{
		Label:    "Defensive Stance",
		ActionID: core.ActionID{SpellID: 71},
		Duration: core.NeverExpires,
	})
	warrior.DefensiveStanceAura.NewExclusiveEffect(stanceEffectCategory, true, core.ExclusiveEffect{
		OnGain: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			warrior.defensiveStanceThreatMultiplier = threatMultiplier()
			ee.Aura.Unit.PseudoStats.ThreatMultiplier *= warrior.defensiveStanceThreatMultiplier
			ee.Aura.Unit.PseudoStats.DamageDealtMultiplier *= 0.9
			ee.Aura.Unit.PseudoStats.DamageTakenMultiplier *= 0.9
		},
		OnExpire: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.PseudoStats.ThreatMultiplier /= warrior.defensiveStanceThreatMultiplier
			ee.Aura.Unit.PseudoStats.DamageDealtMultiplier /= 0.9
			ee.Aura.Unit.PseudoStats.DamageTakenMultiplier /= 0.9
		},
	})
	warrior.RegisterOnItemSwap(func(sim *core.Simulation) {
		next := threatMultiplier()
		if warrior.DefensiveStanceAura.IsActive() {
			warrior.PseudoStats.ThreatMultiplier *= next / warrior.defensiveStanceThreatMultiplier
		}
		warrior.defensiveStanceThreatMultiplier = next
	})
}

func (warrior *Warrior) registerBerserkerStanceAura() {
	warrior.BerserkerStanceAura = warrior.RegisterAura(core.Aura{
		Label:    "Berserker Stance",
		ActionID: core.ActionID{SpellID: 2458},
		Duration: core.NeverExpires,
	})
	warrior.BerserkerStanceAura.NewExclusiveEffect(stanceEffectCategory, true, core.ExclusiveEffect{
		OnGain: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.PseudoStats.ThreatMultiplier *= 0.8
			ee.Aura.Unit.PseudoStats.DamageTakenMultiplier *= 1.1
			ee.Aura.Unit.AddStatDynamic(sim, stats.MeleeCrit, core.CritRatingPerCritChance*3)
		},
		OnExpire: func(ee *core.ExclusiveEffect, sim *core.Simulation) {
			ee.Aura.Unit.PseudoStats.ThreatMultiplier /= 0.8
			ee.Aura.Unit.PseudoStats.DamageTakenMultiplier /= 1.1
			ee.Aura.Unit.AddStatDynamic(sim, stats.MeleeCrit, -core.CritRatingPerCritChance*3)
		},
	})
}

func (warrior *Warrior) registerStances() {
	warrior.Stances = make([]*WarriorSpell, 0)
	stanceCD := warrior.NewTimer()
	warrior.registerBattleStanceAura()
	warrior.registerDefensiveStanceAura()
	warrior.registerBerserkerStanceAura()
	warrior.BattleStance = warrior.makeStanceSpell(BattleStance, warrior.BattleStanceAura, stanceCD)
	warrior.DefensiveStance = warrior.makeStanceSpell(DefensiveStance, warrior.DefensiveStanceAura, stanceCD)
	warrior.BerserkerStance = warrior.makeStanceSpell(BerserkerStance, warrior.BerserkerStanceAura, stanceCD)
}
