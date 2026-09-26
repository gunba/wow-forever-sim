package priest

import (
	"slices"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

func (priest *Priest) ApplyTalents() {
	// Discipline
	priest.registerInnerFocus()
	priest.applyPowerInLight()
	priest.applyTwinDisciplines()
	priest.applyHolyPrecision()
	priest.applyMentalAgility()

	if priest.Talents.SilentResolve > 0 {
		priest.OnSpellRegistered(func(spell *core.Spell) {
			if spell.Flags.Matches(SpellFlagPriest) && spell.SpellSchool.Matches(core.SpellSchoolHoly) {
				spell.ThreatMultiplier *= 1 - .1*float64(priest.Talents.SilentResolve)
			}
		})
	}

	// 17/33/50, not the 17/34/51 that multiplying rank 1 gives. Both the tree and
	// wowforevertalents read it this way; no beta tooltip past rank 1 has been seen.
	priest.PseudoStats.SpiritRegenRateCasting = []float64{0.0, 0.17, 0.33, 0.50}[priest.Talents.Meditation]

	if priest.Talents.MentalStrength > 0 {
		priest.MultiplyStat(stats.Intellect, 1.0+0.03*float64(priest.Talents.MentalStrength))
	}

	// Holy
	priest.applyInspiration()
	priest.applyHolySpecialization()
	priest.applySearingLight()

	priest.PseudoStats.SchoolDamageTakenMultiplier.MultiplyMagicSchools(1 - 0.02*float64(priest.Talents.SpellWarding))

	// The beta client's curves: healing 5% of Spirit per point, damage 1/3/5/6/8%.
	if priest.Talents.SpiritualGuidance > 0 {
		priest.AddStatDependency(stats.Spirit, stats.HealingPower, 0.05*float64(priest.Talents.SpiritualGuidance))
		priest.AddStatDependency(stats.Spirit, stats.SpellDamage, []float64{0, 0.01, 0.03, 0.05, 0.06, 0.08}[priest.Talents.SpiritualGuidance])
	}

	// Shadow Magic
	priest.registerVampiricEmbraceSpell()
	priest.registerShadowform()
	priest.applySpiritTap()
	priest.applyShadowAffinity()
	priest.applyShadowFocus()
	priest.applyShadowWeaving()
	priest.applyDarkness()
	priest.applyEarlyDemise()
}

func (priest *Priest) applyEarlyDemise() {
	if priest.Talents.EarlyDemise == 0 {
		return
	}
	// Client 1310076: 15/30 percentage points of Death crit at <=20% health.
	bonus := 15 * float64(priest.Talents.EarlyDemise) * core.CritRatingPerCritChance
	execute := priest.RegisterAura(core.Aura{
		Label: "Early Demise", ActionID: core.ActionID{SpellID: 1310076},
		Duration: core.NeverExpires,
		OnGain: func(_ *core.Aura, _ *core.Simulation) {
			for _, spell := range priest.ShadowWordDeath {
				if spell != nil {
					spell.BonusCritRating += bonus
				}
			}
		},
		OnExpire: func(_ *core.Aura, _ *core.Simulation) {
			for _, spell := range priest.ShadowWordDeath {
				if spell != nil {
					spell.BonusCritRating -= bonus
				}
			}
		},
	})
	core.MakePermanent(priest.RegisterAura(core.Aura{
		Label: "Early Demise Trigger",
		OnReset: func(_ *core.Aura, sim *core.Simulation) {
			sim.RegisterExecutePhaseCallback(func(sim *core.Simulation, phase int32) {
				if phase == 20 {
					execute.Activate(sim)
				}
			})
		},
	}))
}

// Smite and Penance hit harder while the target is burning from this priest's Holy Fire, 2% per point.
func (priest *Priest) applyPowerInLight() {
	if priest.Talents.PowerInLight == 0 {
		return
	}

	multiplier := 1 + 0.02*float64(priest.Talents.PowerInLight)
	affectedSpellCodes := []int32{SpellCode_PriestSmite, SpellCode_PriestPenance}

	for _, target := range priest.Env.Encounter.TargetUnits {
		target.AddDynamicDamageTakenModifier(func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.Unit == &priest.Unit && slices.Contains(affectedSpellCodes, spell.SpellCode) && priest.hasActiveHolyFire(result.Target) {
				result.Damage *= multiplier
			}
		})
	}
}

func (priest *Priest) hasActiveHolyFire(target *core.Unit) bool {
	for _, spell := range priest.HolyFire {
		if spell != nil && spell.Dot(target).IsActive() {
			return true
		}
	}
	return false
}

func (priest *Priest) applyTwinDisciplines() {
	if priest.Talents.TwinDisciplines == 0 {
		return
	}

	points := float64(priest.Talents.TwinDisciplines)
	priest.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagPriest) && spell.DefaultCast.CastTime == 0 {
			spell.DamageMultiplierAdditive += 0.01 * points
		}
	})
}

func (priest *Priest) applyHolyPrecision() {
	if priest.Talents.HolyPrecision == 0 {
		return
	}

	bonusHit := 6 * float64(priest.Talents.HolyPrecision) * core.SpellHitRatingPerHitChance
	priest.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagPriest) && spell.SpellSchool.Matches(core.SpellSchoolHoly) {
			spell.BonusHitRating += bonusHit
		}
	})
}

func (priest *Priest) applyMentalAgility() {
	if priest.Talents.MentalAgility == 0 {
		return
	}

	affectedSpellCodes := []int32{SpellCode_PriestSmite, SpellCode_PriestHolyFire}
	priest.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Cost == nil || !spell.Flags.Matches(SpellFlagPriest) {
			return
		}

		if spell.DefaultCast.CastTime == 0 || slices.Contains(affectedSpellCodes, spell.SpellCode) {
			spell.Cost.Multiplier -= []int32{0, 3, 7, 10}[priest.Talents.MentalAgility]
		}
	})
}

func (priest *Priest) applyHolySpecialization() {
	if priest.Talents.HolySpecialization == 0 {
		return
	}

	priest.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagPriest) && spell.SpellSchool.Matches(core.SpellSchoolHoly) {
			spell.BonusCritRating += 1 * float64(priest.Talents.HolySpecialization) * core.CritRatingPerCritChance
		}
	})
}

func (priest *Priest) applyInspiration() {
	if priest.Talents.Inspiration == 0 {
		return
	}

	auras := make([]*core.Aura, len(priest.Env.AllUnits))
	for _, unit := range priest.Env.AllUnits {
		if !priest.IsOpponent(unit) {
			aura := core.InspirationAura(unit, priest.Talents.Inspiration)
			auras[unit.UnitIndex] = aura
		}
	}

	priest.RegisterAura(core.Aura{
		Label:    "Inspiration Talent",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnHealDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if slices.Contains([]int32{SpellCode_PriestFlashHeal, SpellCode_PriestHeal, SpellCode_PriestGreaterHeal}, spell.SpellCode) {
				auras[result.Target.UnitIndex].Activate(sim)
			}
		},
	})
}

// Searing Light now buffs every Holy spell and lets Holy Fire ticks refund the next Holy Nova.
// The beta client reads 2/5% Holy damage and a 5/10% chance, and the free Holy Nova (Holy
// Purpose, 1284536) lasts 10 sec.
func (priest *Priest) applySearingLight() {
	if priest.Talents.SearingLight == 0 {
		return
	}

	points := float64(priest.Talents.SearingLight)
	priest.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly] *= []float64{1, 1.02, 1.05}[priest.Talents.SearingLight]

	priest.SearingLightAura = priest.RegisterAura(core.Aura{
		Label:    "Searing Light",
		ActionID: core.ActionID{SpellID: 14909},
		Duration: time.Second * 10,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			if priest.HolyNova != nil {
				priest.HolyNova.Cost.Multiplier -= 100
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			if priest.HolyNova != nil {
				priest.HolyNova.Cost.Multiplier += 100
			}
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell.SpellCode == SpellCode_PriestHolyNova {
				aura.Deactivate(sim)
			}
		},
	})

	procChance := 0.05 * points
	core.MakePermanent(priest.RegisterAura(core.Aura{
		Label: "Searing Light Trigger",
		OnPeriodicDamageDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.SpellCode == SpellCode_PriestHolyFire && sim.Proc(procChance, "Searing Light") {
				priest.SearingLightAura.Activate(sim)
			}
		},
	}))
}

func (priest *Priest) applySpiritTap() {
	if priest.Talents.SpiritTap == 0 {
		return
	}

	spellID := []int32{0, 15270, 15335, 15336, 15337, 15338}[priest.Talents.SpiritTap]
	statDep := priest.NewDynamicMultiplyStat(stats.Spirit, 2.0)

	priest.SpiritTapAura = priest.RegisterAura(core.Aura{
		ActionID: core.ActionID{SpellID: spellID},
		Label:    "Spirit Tap",
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			priest.EnableDynamicStatDep(sim, statDep)
			priest.PseudoStats.SpiritRegenRateCasting += 0.50
			priest.UpdateManaRegenRates()
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			priest.DisableDynamicStatDep(sim, statDep)
			priest.PseudoStats.SpiritRegenRateCasting -= 0.50
			priest.UpdateManaRegenRates()
		},
	})
}

func (priest *Priest) applyShadowAffinity() {
	if priest.Talents.ShadowAffinity == 0 {
		return
	}

	priest.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagPriest) && spell.SpellSchool.Matches(core.SpellSchoolShadow) {
			spell.ThreatMultiplier *= 1 - 0.1*float64(priest.Talents.ShadowAffinity)
		}
	})
}

func (priest *Priest) applyShadowFocus() {
	if priest.Talents.ShadowFocus == 0 {
		return
	}

	bonusHit := 1 * float64(priest.Talents.ShadowFocus) * core.SpellHitRatingPerHitChance
	priest.OnSpellRegistered(func(spell *core.Spell) {
		if spell.Flags.Matches(SpellFlagPriest) && spell.SpellSchool.Matches(core.SpellSchoolShadow) {
			spell.BonusHitRating += bonusHit
		}
	})
}

// The raid debuff version of Shadow Weaving is a Classic mechanic, in Forever it buffs the priest instead.
func (priest *Priest) applyShadowWeaving() {
	if priest.Talents.ShadowWeaving == 0 {
		return
	}

	priest.shadowWeavingProcChance = []float64{0, 0.33, 0.67, 1.00}[priest.Talents.ShadowWeaving]

	priest.ShadowWeavingAura = priest.RegisterAura(core.Aura{
		Label: "Shadow Weaving",
		// The stacking buff is 15258 - aura 270 at 2 per stack, "increase the Shadow damage you
		// deal". Reading the talent's own rank table here put Classic's 15331 and 15332 on it,
		// two spells Forever deleted, so a two or three point priest wore an id with nothing
		// behind it.
		ActionID:  core.ActionID{SpellID: 15258},
		Duration:  time.Second * 15,
		MaxStacks: 5,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexShadow] /= 1 + 0.02*float64(oldStacks)
			aura.Unit.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexShadow] *= 1 + 0.02*float64(newStacks)
		},
	})
}

func (priest *Priest) AddShadowWeavingStack(sim *core.Simulation) {
	if priest.ShadowWeavingAura == nil || !sim.Proc(priest.shadowWeavingProcChance, "Shadow Weaving") {
		return
	}

	priest.ShadowWeavingAura.Activate(sim)
	priest.ShadowWeavingAura.AddStack(sim)
}

// Classic's Darkness raised the damage of five named Shadow spells. The Forever beta client's (15259)
// raises all Shadow damage done, 2% per point.
func (priest *Priest) applyDarkness() {
	if priest.Talents.Darkness == 0 {
		return
	}

	priest.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexShadow] *= 1 + 0.02*float64(priest.Talents.Darkness)
}

func (priest *Priest) registerInnerFocus() {
	if !priest.Talents.InnerFocus {
		return
	}

	actionID := core.ActionID{SpellID: 14751}

	critAffected := func(spell *core.Spell) bool {
		if !priest.Env.IsForever() {
			return spell.Flags.Matches(SpellFlagPriest) && spell.Cost != nil
		}
		// Effect 692357's crit mask differs from effect 692356's free-cast
		// mask. Penance's modeled bolts share their parent's SpellCode.
		switch spell.SpellCode {
		case SpellCode_PriestSmite, SpellCode_PriestMindBlast, SpellCode_PriestShadowWordPain,
			SpellCode_PriestDevouringPlague, SpellCode_PriestHolyFire, SpellCode_PriestHolyNova,
			SpellCode_PriestPenance, SpellCode_PriestFlashHeal, SpellCode_PriestGreaterHeal,
			SpellCode_PriestHeal, SpellCode_PriestVampiricTouch:
			return true
		default:
			return false
		}
	}

	priest.InnerFocusAura = priest.RegisterAura(core.Aura{
		Label:    "Inner Focus",
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range priest.Spellbook {
				if spell.Flags.Matches(SpellFlagPriest) && spell.Cost != nil {
					spell.Cost.Multiplier -= 100
				}
				if critAffected(spell) {
					spell.BonusCritRating += 25 * core.SpellCritRatingPerCritChance
				}
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range priest.Spellbook {
				if spell.Flags.Matches(SpellFlagPriest) && spell.Cost != nil {
					spell.Cost.Multiplier += 100
				}
				if critAffected(spell) {
					spell.BonusCritRating -= 25 * core.SpellCritRatingPerCritChance
				}
			}
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell.Flags.Matches(SpellFlagPriest) {
				// Remove the buff and put skill on CD
				aura.Deactivate(sim)
				priest.InnerFocus.CD.Use(sim)
				priest.UpdateMajorCooldowns()
			}
		},
	})

	priest.InnerFocus = priest.RegisterSpell(core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagAPL,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    priest.NewTimer(),
				Duration: time.Minute * 3,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			priest.InnerFocusAura.Activate(sim)
		},
	})

	priest.AddMajorCooldown(core.MajorCooldown{
		Spell: priest.InnerFocus,
		Type:  core.CooldownTypeDPS,
	})
}

func (priest *Priest) registerShadowform() {
	if !priest.Talents.Shadowform {
		return
	}

	actionID := core.ActionID{SpellID: 15473}

	// The beta client's 15473: +10% Shadow damage, -50% Shadow mana cost, +100% Shadow critical
	// strike damage bonus, -15% Physical damage taken.
	priest.ShadowformAura = priest.RegisterAura(core.Aura{
		Label:    "Shadowform",
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexShadow] *= 1.10
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexPhysical] *= 0.85
			for _, spell := range priest.shadowformSpells() {
				spell.CritDamageBonus += 1
				if spell.Cost != nil {
					spell.Cost.Multiplier -= 50
				}
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexShadow] /= 1.10
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexPhysical] /= 0.85
			for _, spell := range priest.shadowformSpells() {
				spell.CritDamageBonus -= 1
				if spell.Cost != nil {
					spell.Cost.Multiplier += 50
				}
			}
		},
	})

	priest.Shadowform = priest.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{BaseCost: 0.40},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			priest.ShadowformAura.Activate(sim)
		},
	})
}

func (priest *Priest) shadowformSpells() []*core.Spell {
	return core.FilterSlice(priest.Spellbook, func(spell *core.Spell) bool {
		return spell.Flags.Matches(SpellFlagPriest) && spell.SpellSchool.Matches(core.SpellSchoolShadow)
	})
}
