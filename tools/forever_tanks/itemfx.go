package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

// Generic item effects for items the simulator does not implement, described by
// tools/forever_weights/item_effects.py from the client spell tables. Each spec is
// registered as an ordinary item effect, so it only acts while that item is equipped.

type fxEffect struct {
	Kind       string             `json:"kind"` // damage, dot, stats, haste, extra_attacks, armor_debuff
	Spell      int32              `json:"spell"`
	School     string             `json:"school"`
	Min        float64            `json:"min"`
	Range      float64            `json:"range"`
	Tick       float64            `json:"tick"`
	Ticks      int32              `json:"ticks"`
	PeriodMs   int                `json:"period_ms"`
	Stats      map[string]float64 `json:"stats"` // stat index -> amount
	Pct        float64            `json:"pct"`
	N          int32              `json:"n"`
	Amount     float64            `json:"amount"`
	DurationMs int                `json:"duration_ms"`
	Stacks     int32              `json:"stacks"`
}

type fxSpec struct {
	Item       int32      `json:"item"`
	Name       string     `json:"name"`
	Trigger    string     `json:"trigger"` // coh, melee_hit, ranged_hit, struck, spell_hit, use
	PPM        float64    `json:"ppm"`
	Chance     float64    `json:"chance"`
	ICDMs      int        `json:"icd_ms"`
	CooldownMs int        `json:"cooldown_ms"`
	Effects    []fxEffect `json:"effects"`
}

var fxSchools = map[string]core.SpellSchool{"physical": core.SpellSchoolPhysical, "arcane": core.SpellSchoolArcane,
	"fire": core.SpellSchoolFire, "frost": core.SpellSchoolFrost, "holy": core.SpellSchoolHoly,
	"nature": core.SpellSchoolNature, "shadow": core.SpellSchoolShadow}

func ms(v int) time.Duration { return time.Duration(v) * time.Millisecond }

func registerItemFx(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var specs []fxSpec
	if err := json.Unmarshal(data, &specs); err != nil {
		panic(err)
	}
	for _, spec := range specs {
		if core.HasItemEffect(spec.Item) {
			continue
		}
		spec := spec
		core.NewItemEffect(spec.Item, func(agent core.Agent) { applyItemFx(agent.GetCharacter(), spec) })
	}
}

type fxAction func(sim *core.Simulation, target *core.Unit, trigger *core.Spell)

func applyItemFx(character *core.Character, spec fxSpec) {
	var actions []fxAction
	for i, effect := range spec.Effects {
		if action := buildFxAction(character, spec, i, effect); action != nil {
			actions = append(actions, action)
		}
	}
	run := func(sim *core.Simulation, target *core.Unit, trigger *core.Spell) {
		for _, action := range actions {
			action(sim, target, trigger)
		}
	}
	if spec.Trigger == "use" {
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{ItemID: spec.Item},
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{CD: core.Cooldown{Timer: character.NewTimer(),
				Duration: max(ms(spec.CooldownMs), time.Second)}},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				run(sim, target, spell)
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{Spell: spell, Type: core.CooldownTypeDPS})
		return
	}
	trigger := core.ProcTrigger{
		Name:       spec.Name + " " + strconv.Itoa(int(spec.Item)),
		ActionID:   core.ActionID{ItemID: spec.Item},
		Callback:   core.CallbackOnSpellHitDealt,
		Outcome:    core.OutcomeLanded,
		PPM:        spec.PPM,
		ProcChance: spec.Chance,
		ICD:        ms(spec.ICDMs),
	}
	switch spec.Trigger {
	case "coh":
		trigger.ProcMask = character.GetProcMaskForItem(spec.Item)
		trigger.SpellFlagsExclude = core.SpellFlagSuppressWeaponProcs
	case "melee_hit":
		trigger.ProcMask = core.ProcMaskMelee
		trigger.SpellFlagsExclude = core.SpellFlagSuppressEquipProcs
	case "ranged_hit":
		trigger.ProcMask = core.ProcMaskRanged
		trigger.SpellFlagsExclude = core.SpellFlagSuppressEquipProcs
	case "spell_hit":
		trigger.ProcMask = core.ProcMaskSpellDamage
	case "struck":
		trigger.Callback = core.CallbackOnSpellHitTaken
		trigger.ProcMask = core.ProcMaskMelee
	default:
		panic("unknown item effect trigger " + spec.Trigger)
	}
	if trigger.PPM == 0 && trigger.ProcChance == 0 {
		trigger.ProcChance = 1
	}
	trigger.Handler = func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		target := result.Target
		if spec.Trigger == "struck" {
			target = spell.Unit
		}
		run(sim, target, spell)
	}
	core.MakeProcTriggerAura(&character.Unit, trigger)
}

func buildFxAction(character *core.Character, spec fxSpec, index int, effect fxEffect) fxAction {
	label := spec.Name + " " + strconv.Itoa(int(spec.Item)) + " " + strconv.Itoa(index)
	actionID := core.ActionID{SpellID: effect.Spell, Tag: int32(index)}
	school, ok := fxSchools[effect.School]
	if !ok {
		school = core.SpellSchoolPhysical
	}
	physical := school == core.SpellSchoolPhysical
	switch effect.Kind {
	case "damage":
		config := core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      school,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
		}
		if physical {
			config.DefenseType = core.DefenseTypeMelee
			config.ProcMask = core.ProcMaskMeleeSpecial
			config.Flags |= core.SpellFlagSuppressEquipProcs | core.SpellFlagSuppressWeaponProcs
		}
		config.ApplyEffects = func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := effect.Min + effect.Range*sim.RandomFloat(label)
			if physical {
				spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMeleeSpecialHitAndCrit)
			} else {
				spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicHitAndCrit)
			}
		}
		spell := character.RegisterSpell(config)
		return func(sim *core.Simulation, target *core.Unit, _ *core.Spell) { spell.Cast(sim, target) }
	case "dot":
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      school,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			Dot: core.DotConfig{
				NumberOfTicks: max(effect.Ticks, 1),
				TickLength:    max(ms(effect.PeriodMs), time.Second),
				Aura:          core.Aura{Label: label},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, effect.Tick, dot.OutcomeTick)
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				outcome := spell.OutcomeMagicHit
				if physical {
					outcome = spell.OutcomeAlwaysHit
				}
				if result := spell.CalcAndDealOutcome(sim, target, outcome); result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
		return func(sim *core.Simulation, target *core.Unit, _ *core.Spell) { spell.Cast(sim, target) }
	case "stats":
		var bonus stats.Stats
		for key, amount := range effect.Stats {
			index, err := strconv.Atoi(key)
			if err != nil || index < 0 || index >= len(bonus) {
				panic("bad stat index " + key)
			}
			bonus[index] += amount
		}
		aura := character.NewTemporaryStatsAura(label, actionID, bonus, max(ms(effect.DurationMs), time.Second))
		return func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) { aura.Activate(sim) }
	case "haste":
		multiplier := 1 + effect.Pct/100
		aura := character.GetOrRegisterAura(core.Aura{
			Label:    label,
			ActionID: actionID,
			Duration: max(ms(effect.DurationMs), time.Second),
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, multiplier)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1/multiplier)
			},
		})
		return func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) { aura.Activate(sim) }
	case "extra_attacks":
		return func(sim *core.Simulation, _ *core.Unit, trigger *core.Spell) {
			if trigger != nil && trigger.ProcMask.Matches(core.ProcMaskMelee) {
				character.AutoAttacks.ExtraMHAttackProc(sim, max(effect.N, 1), actionID, trigger)
			} else {
				character.AutoAttacks.ExtraMHAttack(sim, max(effect.N, 1), actionID, core.ActionID{ItemID: spec.Item})
			}
		}
	case "armor_debuff":
		auras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				Label:     label,
				ActionID:  actionID,
				Duration:  max(ms(effect.DurationMs), time.Second),
				MaxStacks: max(effect.Stacks, 1),
				OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
					aura.Unit.AddStatDynamic(sim, stats.Armor, -effect.Amount*float64(newStacks-oldStacks))
				},
			})
		})
		return func(sim *core.Simulation, target *core.Unit, _ *core.Spell) {
			if target == nil || target.Type != core.EnemyUnit {
				return
			}
			aura := auras.Get(target)
			aura.Activate(sim)
			aura.AddStack(sim)
		}
	}
	return nil
}
