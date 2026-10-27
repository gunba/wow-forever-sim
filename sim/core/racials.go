package core

import (
	"fmt"
	"slices"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func applyRaceEffects(agent Agent) {
	character := agent.GetCharacter()

	// Forever drops every +10 resistance racial and, since client 1.60.1.70425, the weapon racials.
	// The launch values are Blizzard's October 26 tuning pass, which is not in any client build yet:
	// https://www.wowhead.com/forever/news/racial-ability-tuning-pass-for-wow-forever-383597
	forever := character.Env.IsForever()

	switch character.Race {
	case proto.Race_RaceDwarf:
		if forever {
			// Artisan's Precision: +1% hit with spells and attacks, in place of Earth Fury's critical damage.
			character.AddStat(stats.MeleeHit, 1*MeleeHitRatingPerHitChance)
			character.AddStat(stats.SpellHit, 1*SpellHitRatingPerHitChance)
			character.beastSlayingAura(1.05)
		} else {
			character.AddStat(stats.FrostResistance, 10)
			character.GunSpecializationAura()
		}

		actionID := ActionID{SpellID: 20594}

		statDep := character.NewDynamicMultiplyStat(stats.Armor, 1.1)
		stoneFormAura := character.NewTemporaryStatsAuraWrapped("Stoneform", actionID, stats.Stats{}, time.Second*8, func(aura *Aura) {
			aura.ApplyOnGain(func(aura *Aura, sim *Simulation) {
				aura.Unit.EnableDynamicStatDep(sim, statDep)
			})
			aura.ApplyOnExpire(func(aura *Aura, sim *Simulation) {
				aura.Unit.DisableDynamicStatDep(sim, statDep)
			})
		})

		spell := character.RegisterSpell(SpellConfig{
			ActionID: actionID,
			Flags:    SpellFlagNoOnCastComplete,
			Cast: CastConfig{
				CD: Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 3,
				},
			},
			ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
				stoneFormAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(MajorCooldown{
			Spell: spell,
			Type:  CooldownTypeSurvival,
			ShouldActivate: func(s *Simulation, c *Character) bool {
				// Only castable with manual APL Action
				return false
			},
		})
	case proto.Race_RaceGnome:
		if forever {
			// Client 20591 / 1259802 / 1259803 increase the resource pools, not Intellect.
			character.MultiplyStat(stats.Mana, 1.05)
			if character.HasEnergyBar() {
				character.energyBar.maxEnergy *= 1.05
			}
			if character.HasRageBar() {
				character.rageBar.maxRage *= 1.05
			}
			character.registerEureka()
		} else {
			character.AddStat(stats.ArcaneResistance, 10)
			character.MultiplyStat(stats.Intellect, 1.05)
		}
	case proto.Race_RaceHuman:
		if forever {
			// The Human Spirit: +10% Spirit. Ruthless (20597): +1% critical strike chance with all spells and attacks.
			character.MultiplyStat(stats.Spirit, 1.1)
			character.AddStat(stats.MeleeCrit, 1*CritRatingPerCritChance)
			character.AddStat(stats.SpellCrit, 1*SpellCritRatingPerCritChance)
		} else {
			character.MultiplyStat(stats.Spirit, 1.05)
			character.SwordSpecializationAura()
			character.MaceSpecializationAura()
		}
	case proto.Race_RaceNightElf:
		if forever {
			character.registerElunesLight()
		} else {
			character.AddStat(stats.NatureResistance, 10)
		}
		character.AddStat(stats.Dodge, 1)
		// TODO: Shadowmeld?
	case proto.Race_RaceOrc:
		// Forever's Orc has no weapon or crit racial (Savage was removed). Shatter Curse, a curse
		// dispel, is nothing the sim measures.
		if !forever {
			character.AxeSpecializationAura()
		}

		if character.Class == proto.Class_ClassHunter || character.Class == proto.Class_ClassWarlock {
			// Command: damage dealt by Hunter and Warlock pets increased by 5%
			for _, pet := range character.Pets {
				if !pet.IsGuardian() {
					pet.PseudoStats.DamageDealtMultiplier *= 1.05
				}
			}
		}
		if forever {
			character.registerForeverBloodFury()
			break
		}

		// Classic Blood Fury
		actionID := ActionID{SpellID: 20572}
		var bloodFuryAP float64
		bloodFuryAura := character.RegisterAura(Aura{
			Label:    "Blood Fury",
			ActionID: actionID,
			Duration: time.Second * 15,
			// Tooltip is misleading; ap bonus is base AP plus AP from current strength, does not include +attackpower on items/buffs
			OnGain: func(aura *Aura, sim *Simulation) {
				bloodFuryAP = (character.GetBaseStats()[stats.AttackPower] + (character.GetStat(stats.Strength) * APPerStrength[character.Class]) + (character.GetStat(stats.Agility) * APPerAgility[character.Class])) * 0.25
				character.AddStatDynamic(sim, stats.AttackPower, bloodFuryAP)
			},

			OnExpire: func(aura *Aura, sim *Simulation) {
				character.AddStatDynamic(sim, stats.AttackPower, -bloodFuryAP)
			},
		})

		spell := character.RegisterSpell(SpellConfig{
			ActionID: actionID,
			Flags:    SpellFlagNoOnCastComplete,
			Cast: CastConfig{
				DefaultCast: Cast{
					GCD: GCDDefault,
				},
				CD: Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},
			ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
				bloodFuryAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(MajorCooldown{
			Spell: spell,
			Type:  CooldownTypeDPS,
		})
	case proto.Race_RaceTauren:
		// Endurance (20550): 5% health. Forever's hit, then critical damage, rider is gone.
		character.MultiplyStat(stats.Health, 1.05)
		if !forever {
			character.AddStat(stats.NatureResistance, 10)
		}
	case proto.Race_RaceTroll:
		// Forever's troll has four racials and neither ranged weapon specialization is
		// among them; Rapid Regeneration and Regeneration took their place.
		if !forever {
			character.BowSpecializationAura()
			character.ThrownSpecializationAura()
		}

		character.beastSlayingAura(1.05)

		// Berserking
		berserkingTimer := character.NewTimer()
		if forever {
			character.registerForeverBerserking(berserkingTimer)
		} else {
			// Baseline cooldown
			makeBerserkingCooldown(character, 0, berserkingTimer)
			// Hard-coded percentage cooldown options
			makeBerserkingCooldown(character, .1, berserkingTimer)
			makeBerserkingCooldown(character, .15, berserkingTimer)
			makeBerserkingCooldown(character, .2, berserkingTimer)
			makeBerserkingCooldown(character, .25, berserkingTimer)
			makeBerserkingCooldown(character, .3, berserkingTimer)
		}
	case proto.Race_RaceUndead:
		if !forever {
			character.AddStat(stats.ShadowResistance, 10)
		} else {
			character.registerTouchOfTheGrave()
		}
	case proto.Race_RaceSkyborneHighOrder, proto.Race_RaceSkyborneWindshaper:
		// The Skyborne are a Forever race, so under Classic rules they have no racials to
		// grant. A saved setting can still name one - the race picker offers whatever the
		// spec allows - and without this guard that saved race would carry Forever's haste
		// and elemental damage into a Classic sim.
		if !forever {
			break
		}

		// Wind Blessed: 2% haste (3% in client 70425, 1% before).
		character.PseudoStats.MeleeSpeedMultiplier *= 1.02
		character.PseudoStats.RangedSpeedMultiplier *= 1.02
		character.PseudoStats.CastSpeedMultiplier *= 1.02
		character.PseudoStats.EnergyHasteMultiplier *= 1.02

		// Elemental Insight
		character.mobTypeDamageAura(proto.MobType_MobTypeElemental, 1.05)

		// Skysight grants movement speed, not attack/spell power. Read Ley Line
		// belongs to High Order; neither is used in this stationary Horde model.
	}
}

func (character *Character) registerForeverBloodFury() {
	actionID := ActionID{SpellID: 20572}
	var deps []*stats.StatDependency
	for _, stat := range []stats.Stat{
		stats.AttackPower, stats.RangedAttackPower, stats.SpellPower, stats.SpellDamage, stats.HealingPower,
		stats.ArcanePower, stats.FirePower, stats.FrostPower,
		stats.HolyPower, stats.NaturePower, stats.ShadowPower,
	} {
		deps = append(deps, character.NewDynamicMultiplyStat(stat, 1.1))
	}
	aura := character.RegisterAura(Aura{
		Label: "Blood Fury", ActionID: actionID, Duration: time.Second * 15,
		OnGain: func(aura *Aura, sim *Simulation) {
			for _, dep := range deps {
				character.EnableDynamicStatDep(sim, dep)
			}
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			for _, dep := range deps {
				character.DisableDynamicStatDep(sim, dep)
			}
		},
	})
	spell := character.RegisterSpell(SpellConfig{
		ActionID: actionID, Flags: SpellFlagNoOnCastComplete,
		// Three minutes since 70425 (was two).
		Cast:         CastConfig{CD: Cooldown{Timer: character.NewTimer(), Duration: time.Minute * 3}},
		ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) { aura.Activate(sim) },
	})
	character.AddMajorCooldown(MajorCooldown{Spell: spell, Type: CooldownTypeDPS})
}

func (character *Character) registerForeverBerserking(timer *Timer) {
	actionID := ActionID{SpellID: 20554}
	aura := character.RegisterAura(Aura{
		Label: "Berserking", ActionID: actionID, Duration: time.Second * 10,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.MultiplyCastSpeed(1.1)
			character.MultiplyAttackSpeed(sim, 1.1)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			character.MultiplyCastSpeed(1 / 1.1)
			character.MultiplyAttackSpeed(sim, 1/1.1)
		},
	})
	spell := character.RegisterSpell(SpellConfig{
		ActionID: actionID, Flags: SpellFlagNoOnCastComplete,
		Cast:         CastConfig{CD: Cooldown{Timer: timer, Duration: time.Minute * 3}},
		ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) { aura.Activate(sim) },
	})
	character.AddMajorCooldown(MajorCooldown{Spell: spell, Type: CooldownTypeDPS})
}

// SkillLineAbility assigns 1260189 to Warrior/Paladin/Rogue and 1260201 to
// Priest/Mage/Warlock. Since 70425 SpellAuraOptions gives 20%/30% and a 25-second ICD (was
// 5%/10% and one second). The triggered drain, 1260198, is 5% of maximum health, Shadow, and
// cannot crit; it no longer scales with Shadow damage modifiers on the caster or the target.
func (character *Character) registerTouchOfTheGrave() {
	actionID := ActionID{SpellID: 1260198}
	passiveID, procChance := int32(1260189), 0.2
	switch character.Class {
	case proto.Class_ClassPriest, proto.Class_ClassMage, proto.Class_ClassWarlock:
		passiveID, procChance = 1260201, 0.3
	}
	icd := Cooldown{Timer: character.NewTimer(), Duration: 25 * time.Second}
	healthMetrics := character.NewHealthMetrics(actionID)

	drain := character.RegisterSpell(SpellConfig{
		ActionID:    actionID,
		SpellSchool: SpellSchoolShadow,
		DefenseType: DefenseTypeMagic,
		ProcMask:    ProcMaskEmpty,
		// The drain is a proc off another hit, so it must not feed the procs that spawned
		// it or two undead attacks would chain into each other.
		Flags: SpellFlagNoOnCastComplete | SpellFlagPassiveSpell | SpellFlagIgnoreModifiers,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			maxHealth := character.MaxHealth()
			result := spell.CalcAndDealDamage(sim, target, maxHealth*0.05, spell.OutcomeMagicHit)

			// Only the specs that track a health bar can be healed; for everyone else the
			// drain is still damage, it just has nothing to return the health to.
			if result.Landed() && character.HasHealthBar() {
				character.GainHealth(sim, result.Damage, healthMetrics)
			}
		},
	})

	MakePermanent(character.RegisterAura(Aura{
		Label:    "Touch of the Grave",
		ActionID: ActionID{SpellID: passiveID},
		Icd:      &icd,
		OnSpellHitDealt: func(_ *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			// The client now requires an enemy-facing damaging ability. Dot
			// applications such as Shadow Word: Pain count once; their ticks use
			// OnPeriodicDamageDealt and cannot trigger this aura. Utility casts
			// and self-damage (e.g. Demonic Rune) do not qualify.
			if !result.Landed() || result.Target.Type != EnemyUnit ||
				(result.Damage <= 0 && !spell.Flags.Matches(SpellFlagPureDot)) ||
				spell == drain || !icd.IsReady(sim) {
				return
			}
			if sim.Proc(procChance, "Touch of the Grave") {
				icd.Use(sim)
				drain.Cast(sim, result.Target)
			}
		},
	}))
}

// Elune's Light, the night elf's Forever racial cooldown: 10% critical strike for 15
// seconds on a three minute cooldown.
func (character *Character) registerElunesLight() {
	actionID := ActionID{SpellID: 1259799}

	aura := character.RegisterAura(Aura{
		Label:    "Elune's Light",
		ActionID: actionID,
		Duration: time.Second * 15,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.AddStatsDynamic(sim, stats.Stats{
				stats.MeleeCrit: 10 * CritRatingPerCritChance,
				stats.SpellCrit: 10 * SpellCritRatingPerCritChance,
			})
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			character.AddStatsDynamic(sim, stats.Stats{
				stats.MeleeCrit: -10 * CritRatingPerCritChance,
				stats.SpellCrit: -10 * SpellCritRatingPerCritChance,
			})
		},
	})

	spell := character.RegisterSpell(SpellConfig{
		ActionID: actionID,
		Flags:    SpellFlagNoOnCastComplete,
		Cast: CastConfig{
			CD: Cooldown{
				Timer:    character.NewTimer(),
				Duration: time.Minute * 3,
			},
		},
		ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
			aura.Activate(sim)
		},
	})

	character.AddMajorCooldown(MajorCooldown{
		Spell: spell,
		Type:  CooldownTypeDPS,
	})
}

// Troll Beast Slaying, and Dwarf Big Game Hunter under Forever.
func (character *Character) beastSlayingAura(multiplier float64) {
	character.mobTypeDamageAura(proto.MobType_MobTypeBeast, multiplier)
}

func (character *Character) mobTypeDamageAura(mobType proto.MobType, multiplier float64) {
	character.Env.RegisterPostFinalizeEffect(func() {
		for _, t := range character.Env.Encounter.Targets {
			if t.MobType == mobType {
				// Damage only: Beast Slaying 20557, Big Game Hunter 1259721 and Elemental Insight 1259707
				// are aura 168 alone; the slaying rows that raise critical damage carry aura 547 too
				// (wowsims/forever #536).
				for _, at := range character.AttackTables[t.UnitIndex] {
					at.DamageDealtMultiplier *= multiplier
				}
			}
		}
	})
}

// If customPercentage is 0, use the baseline Berserking calculations from health missing
// otherwise create a cooldown hard-coded to the custom percentage.
func makeBerserkingCooldown(character *Character, customPercentage float64, timer *Timer) {
	actionID := ActionID{SpellID: 26297, Tag: int32(customPercentage * 20)}

	label := "Berserking"
	if customPercentage != 0 {
		label = fmt.Sprintf("%s (%d)", label, int(customPercentage*100))
	}

	calcBerserkingPct := func() float64 {
		if customPercentage != 0 {
			return customPercentage
		}
		// from 10% at full health to 30% at 40% or less health
		switch hp := character.CurrentHealthPercent(); {
		case hp >= 1:
			return 0.1
		case hp <= 0.4:
			return 0.3
		default:
			return 0.1 + (1-hp)/3
		}
	}

	var berserkingAura *Aura
	var berserkingHaste float64
	if character.HasManaBar() {
		// Mana-using classes gain a flat % reduction in attack and cast speed
		berserkingAura = character.RegisterAura(Aura{
			Label:    label,
			ActionID: actionID,
			Duration: time.Second * 10,
			OnGain: func(aura *Aura, sim *Simulation) {
				berserkingHaste = 1 / (1 - calcBerserkingPct())

				character.MultiplyCastSpeed(berserkingHaste)
				character.MultiplyAttackSpeed(sim, berserkingHaste)

				if sim.Log != nil {
					character.Log(sim, "Berserking increased attack and casting speed by %.2f%% (%.2f%% hp)", berserkingHaste*100-100, character.CurrentHealthPercent()*100)
				}
			},
			OnExpire: func(aura *Aura, sim *Simulation) {
				character.MultiplyCastSpeed(1 / berserkingHaste)
				character.MultiplyAttackSpeed(sim, 1/berserkingHaste)
			},
		})
	} else {
		// Non-mana bar classes gain a flat % reduction in attack and cast speed
		berserkingAura = character.RegisterAura(Aura{
			Label:    label,
			ActionID: actionID,
			Duration: time.Second * 10,
			OnGain: func(aura *Aura, sim *Simulation) {
				berserkingHaste = 1 + calcBerserkingPct()

				character.MultiplyAttackSpeed(sim, berserkingHaste)

				if sim.Log != nil {
					character.Log(sim, "Berserking increased attack speed by %.2f%% (%.2f%% hp)", berserkingHaste*100-100, character.CurrentHealthPercent()*100)
				}
			},
			OnExpire: func(aura *Aura, sim *Simulation) {
				character.MultiplyAttackSpeed(sim, 1/berserkingHaste)
			},
		})
	}

	config := SpellConfig{
		ActionID: actionID,

		Cast: CastConfig{
			CD: Cooldown{
				Timer:    timer,
				Duration: time.Minute * 3,
			},
		},

		ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
			berserkingAura.Activate(sim)
		},
	}

	switch {
	case character.HasManaBar():
		config.ManaCost = ManaCostOptions{BaseCost: 0.07}
	case character.HasRageBar():
		config.RageCost = RageCostOptions{Cost: 5}
	case character.HasEnergyBar():
		config.EnergyCost = EnergyCostOptions{Cost: 10}
	}

	berserkingSpell := character.RegisterSpell(config)

	character.AddMajorCooldown(MajorCooldown{
		Spell: berserkingSpell,
		Type:  CooldownTypeDPS,
	})
}

func (character *Character) GetFaction() proto.Faction {
	if slices.Contains([]proto.Race{proto.Race_RaceHuman, proto.Race_RaceDwarf, proto.Race_RaceGnome, proto.Race_RaceNightElf, proto.Race_RaceSkyborneHighOrder}, character.Race) {
		return proto.Faction_Alliance
	} else if slices.Contains([]proto.Race{proto.Race_RaceOrc, proto.Race_RaceTroll, proto.Race_RaceTauren, proto.Race_RaceUndead, proto.Race_RaceSkyborneWindshaper}, character.Race) {
		return proto.Faction_Horde
	} else {
		return proto.Faction_Unknown
	}
}
