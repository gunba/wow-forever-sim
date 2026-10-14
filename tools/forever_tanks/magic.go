package main

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// Optional magic damage for the tank workloads. A target with this ID keeps its melee
// swings and also casts a spell at the tank every two seconds, cycling Fire, Frost,
// Shadow, Nature and Arcane. Spells use the magic hit table and partial resists from the
// tank's resistances; they cannot be dodged, parried or blocked and ignore armor.
const magicSourceID = 990001
const magicDPSLabel = "Magic DPS"

var magicSchools = []core.SpellSchool{core.SpellSchoolFire, core.SpellSchoolFrost, core.SpellSchoolShadow,
	core.SpellSchoolNature, core.SpellSchoolArcane}

type magicSourceAI struct {
	target *core.Target
	spells []*core.Spell
	damage float64
	period time.Duration
	next   int
	due    time.Duration
}

func registerMagicSource() {
	core.AddPresetTarget(&core.PresetTarget{
		PathPrefix: "Forever tank workloads",
		Config:     &proto.Target{Id: magicSourceID, Name: "Magic source"},
		AI:         func() core.TargetAI { return &magicSourceAI{} },
	})
}

func (ai *magicSourceAI) Initialize(target *core.Target, config *proto.Target) {
	ai.target = target
	ai.period = 2 * time.Second
	for _, input := range config.TargetInputs {
		if input.Label == magicDPSLabel {
			ai.damage = input.NumberValue * ai.period.Seconds()
		}
	}
	for i, school := range magicSchools {
		ai.spells = append(ai.spells, target.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: magicSourceID*10 + int32(i)},
			SpellSchool:      school,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskSpellDamage,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, tank *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, tank, ai.damage, spell.OutcomeMagicHit)
			},
		}))
	}
}

func (ai *magicSourceAI) Reset(*core.Simulation) {
	ai.next, ai.due = 0, 0
}

func (ai *magicSourceAI) ExecuteCustomRotation(sim *core.Simulation) {
	tank := ai.target.CurrentTarget
	if tank == nil {
		tank = &ai.target.Env.Raid.Parties[0].Players[0].GetCharacter().Unit
	}
	// The rotation can be called more often than the period (after casts and other
	// events), so casts follow their own schedule.
	if ai.damage > 0 && sim.CurrentTime >= ai.due {
		ai.spells[ai.next%len(ai.spells)].Cast(sim, tank)
		ai.next++
		ai.due = sim.CurrentTime + ai.period
	}
	ai.target.WaitUntil(sim, max(ai.due, sim.CurrentTime+time.Millisecond))
}
