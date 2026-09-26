package shaman

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const HealingStreamTotemRanks = 5

// Forever beta client values for the heal each tick.
var HealingStreamTotemSpellId = [HealingStreamTotemRanks + 1]int32{0, 5394, 6375, 6377, 10462, 10463}
var HealingStreamTotemHealId = [HealingStreamTotemRanks + 1]int32{0, 5672, 6371, 6372, 10460, 10461}
var HealingStreamTotemBaseHealing = [HealingStreamTotemRanks + 1]float64{0, 5, 6, 7, 9, 11}
var HealingStreamTotemSpellCoeff = [HealingStreamTotemRanks + 1]float64{0, .022, .022, .022, .022, .022}
var HealingStreamTotemManaCost = [HealingStreamTotemRanks + 1]float64{0, 40, 50, 60, 70, 80}
var HealingStreamTotemLevel = [HealingStreamTotemRanks + 1]int{0, 20, 30, 40, 50, 60}

func (shaman *Shaman) registerHealingStreamTotemSpell() {
	shaman.HealingStreamTotem = make([]*core.Spell, HealingStreamTotemRanks+1)

	for rank := 1; rank <= HealingStreamTotemRanks; rank++ {
		config := shaman.newHealingStreamTotemSpellConfig(rank)

		if config.RequiredLevel <= int(shaman.Level) {
			shaman.HealingStreamTotem[rank] = shaman.RegisterSpell(config)
		}
	}

	shaman.WaterTotems = append(
		shaman.WaterTotems,
		core.FilterSlice(shaman.HealingStreamTotem, func(spell *core.Spell) bool { return spell != nil })...,
	)
}

func (shaman *Shaman) newHealingStreamTotemSpellConfig(rank int) core.SpellConfig {
	spellId := HealingStreamTotemSpellId[rank]
	healId := HealingStreamTotemHealId[rank]
	baseHealing := HealingStreamTotemBaseHealing[rank]*shaman.purificationHealingModifier() + shaman.restorativeTotemsModifier()
	spellCoeff := HealingStreamTotemSpellCoeff[rank]
	manaCost := HealingStreamTotemManaCost[rank]
	level := HealingStreamTotemLevel[rank]

	duration := totemDuration
	healInterval := time.Second * 2

	config := shaman.newTotemSpellConfig(manaCost, spellId)
	config.RequiredLevel = level
	config.Rank = rank

	healSpell := shaman.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: healId},
		SpellSchool: core.SpellSchoolNature,
		ProcMask:    core.ProcMaskSpellHealing,
		Flags:       core.SpellFlagHelpful | core.SpellFlagNoOnCastComplete | core.SpellFlagNoLogs | core.SpellFlagNoMetrics,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealHealing(sim, target, baseHealing, spell.OutcomeHealing)
		},
	})

	config.Hot = core.DotConfig{
		Aura: core.Aura{
			Label: fmt.Sprintf("Healing Stream HoT (Rank %d)", rank),
		},
		NumberOfTicks: int32(duration / healInterval),
		TickLength:    healInterval,
		OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
			healSpell.Cast(sim, target)
		},
	}

	config.ApplyEffects = func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
		shaman.replaceWaterTotem(sim, spell, duration)

		for _, agent := range shaman.Party.Players {
			spell.Hot(&agent.GetCharacter().Unit).Activate(sim)
		}
	}

	return config
}

const ManaSpringTotemRanks = 4

var ManaSpringTotemSpellId = [ManaSpringTotemRanks + 1]int32{0, 5675, 10495, 10496, 10497}
var ManaSpringTotemManaRestore = [ManaSpringTotemRanks + 1]int32{0, 4, 6, 8, 10}
var ManaSpringTotemManaCost = [ManaSpringTotemRanks + 1]float64{0, 40, 60, 80, 100}
var ManaSpringTotemLevel = [ManaSpringTotemRanks + 1]int{0, 26, 36, 46, 56}

func (shaman *Shaman) registerManaSpringTotemSpell() {
	shaman.ManaSpringTotem = make([]*core.Spell, ManaSpringTotemRanks+1)

	for rank := 1; rank <= ManaSpringTotemRanks; rank++ {
		config := shaman.newManaSpringTotemSpellConfig(rank)

		if config.RequiredLevel <= int(shaman.Level) {
			shaman.ManaSpringTotem[rank] = shaman.RegisterSpell(config)
		}
	}

	shaman.WaterTotems = append(
		shaman.WaterTotems,
		core.FilterSlice(shaman.ManaSpringTotem, func(spell *core.Spell) bool { return spell != nil })...,
	)
}

func (shaman *Shaman) newManaSpringTotemSpellConfig(rank int) core.SpellConfig {
	spellId := ManaSpringTotemSpellId[rank]
	// TODO: The sim won't respect the value of a totem dropped via the APL. It uses hard-coded values from buffs.go
	// manaRestoreBase := ManaSpringTotemManaRestore[rank]
	manaCost := ManaSpringTotemManaCost[rank]
	level := ManaSpringTotemLevel[rank]

	duration := totemDuration

	spell := shaman.newTotemSpellConfig(manaCost, spellId)
	spell.RequiredLevel = level
	spell.Rank = rank
	spell.ApplyEffects = func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
		shaman.replaceWaterTotem(sim, spell, duration)
	}
	return spell
}

func (shaman *Shaman) replaceWaterTotem(sim *core.Simulation, next *core.Spell, duration time.Duration) {
	if old := shaman.ActiveTotems[WaterTotem]; old != nil && len(old.Dots()) > 0 {
		for _, agent := range shaman.Party.Players {
			if hot := old.Hot(&agent.GetCharacter().Unit); hot != nil {
				hot.Cancel(sim)
			}
		}
	}
	shaman.TotemExpirations[WaterTotem] = sim.CurrentTime + duration
	shaman.ActiveTotems[WaterTotem] = next
}

var ManaTideTotemSpellId = []int32{0, 16190, 17354, 17359}
var ManaTideTotemManaCost = []float64{0, 10, 30, 60}
var ManaTideTotemManaRestore = []float64{0, 88, 197, 290}
var ManaTideTotemLevel = []int{0, 40, 48, 58}

func (shaman *Shaman) registerManaTideTotemSpell() {
	if !shaman.Env.IsForever() || !shaman.Talents.ManaTideTotem {
		return
	}
	shaman.ManaTideTotem = make([]*core.Spell, len(ManaTideTotemSpellId))
	timer := shaman.NewTimer()
	var highest int
	for rank := 1; rank < len(ManaTideTotemSpellId); rank++ {
		if ManaTideTotemLevel[rank] > int(shaman.Level) {
			continue
		}
		spellID, restore := ManaTideTotemSpellId[rank], ManaTideTotemManaRestore[rank]
		config := shaman.newTotemSpellConfig(ManaTideTotemManaCost[rank], spellID)
		config.SpellSchool = core.SpellSchoolNature
		config.Flags |= core.SpellFlagHelpful
		config.RequiredLevel, config.Rank = ManaTideTotemLevel[rank], rank
		config.Cast.CD = core.Cooldown{Timer: timer, Duration: 5 * time.Minute}
		config.ExtraCastCondition = func(_ *core.Simulation, _ *core.Unit) bool {
			return !shaman.HasActiveAuraWithTag(core.ManaTideTotemAuraTag)
		}
		metrics := make(map[*core.Character]*core.ResourceMetrics)
		for _, agent := range shaman.Party.Players {
			character := agent.GetCharacter()
			if character.HasManaBar() {
				metrics[character] = character.NewManaMetrics(core.ActionID{SpellID: spellID})
			}
		}
		config.Hot = core.DotConfig{
			Aura: core.Aura{
				Label: fmt.Sprintf("Owned Mana Tide (Rank %d)", rank),
				Tag:   core.ManaTideTotemAuraTag,
			},
			NumberOfTicks: 4,
			TickLength:    3 * time.Second,
			OnTick: func(sim *core.Simulation, _ *core.Unit, _ *core.Dot) {
				// Direct periodic mana, not MP5: no haste or MP5-mode scaling.
				for _, agent := range shaman.Party.Players {
					character := agent.GetCharacter()
					if metric := metrics[character]; metric != nil {
						character.AddMana(sim, restore, metric)
					}
				}
			},
		}
		config.ApplyEffects = func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			shaman.replaceWaterTotem(sim, spell, 12*time.Second)
			spell.Hot(&shaman.Unit).Apply(sim)
		}
		shaman.ManaTideTotem[rank] = shaman.RegisterSpell(config)
		shaman.WaterTotems = append(shaman.WaterTotems, shaman.ManaTideTotem[rank])
		highest = rank
	}
	if highest > 0 {
		shaman.AddMajorCooldown(core.MajorCooldown{
			Spell: shaman.ManaTideTotem[highest],
			Type:  core.CooldownTypeMana,
			ShouldActivate: func(_ *core.Simulation, character *core.Character) bool {
				return character.MaxMana()-character.CurrentMana() >= 4*ManaTideTotemManaRestore[highest]
			},
		})
	}
}
