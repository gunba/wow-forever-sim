package core

import (
	"math"
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

type DebuffName int32

const (
	// General Buffs
	DemoralizingShout DebuffName = iota
)

func applyDebuffEffects(target *Unit, targetIdx int, debuffs *proto.Debuffs, raid *proto.Raid) {
	if debuffs.JudgementOfWisdom && targetIdx == 0 {
		jowAura := JudgementOfWisdomAura(target)
		if jowAura != nil {
			MakePermanent(jowAura)
		}
	}

	if targetIdx == 0 {
		if debuffs.JudgementOfTheCrusader == proto.TristateEffect_TristateEffectRegular {
			MakePermanent(JudgementOfTheCrusaderAura(nil, target, 1, 0))
		} else if debuffs.JudgementOfTheCrusader == proto.TristateEffect_TristateEffectImproved {
			MakePermanent(JudgementOfTheCrusaderAura(nil, target, 1.15, 0))
		}
	}

	// Forever's Improved Shadow Bolt and Stormstrike only raise the damage their own
	// caster deals, so another raid member running one is worth nothing here. The
	// warlock and the shaman apply their own in sim/warlock and sim/shaman.
	if debuffs.ImprovedShadowBolt && targetIdx == 0 && !target.Env.IsForever() {
		ExternalIsbCaster(debuffs, target)
	}

	// Forever turned Shadow Weaving into a buff on the priest rather than a debuff on the
	// target, so another raid member running one is worth nothing here - same as Improved
	// Shadow Bolt above. The priest applies its own in sim/priest/talents.go.
	if debuffs.ShadowWeaving && !target.Env.IsForever() {
		aura := ShadowWeavingAura(target)
		SchedulePeriodicDebuffApplication(aura, PeriodicActionOptions{
			Period:          time.Millisecond * 1500,
			NumTicks:        5,
			TickImmediately: true,
			Priority:        ActionPriorityDOT, // High prio
			OnAction: func(sim *Simulation) {
				aura.Activate(sim)
				if aura.IsActive() {
					aura.AddStack(sim)
				}
			},
		}, raid)
	}

	if debuffs.CurseOfElements {
		MakePermanent(CurseOfElementsAura(target))
	}

	if debuffs.CurseOfShadow && !target.Env.IsForever() {
		MakePermanent(CurseOfShadowAura(target))
	}

	if debuffs.ImprovedScorch && targetIdx == 0 && !target.Env.IsForever() {
		aura := ImprovedScorchAura(target)
		SchedulePeriodicDebuffApplication(aura, PeriodicActionOptions{
			Period:          time.Millisecond * 1500,
			NumTicks:        5,
			TickImmediately: true,
			Priority:        ActionPriorityDOT, // High prio
			OnAction: func(sim *Simulation) {
				aura.Activate(sim)
				if aura.IsActive() {
					aura.AddStack(sim)
				}
			},
		}, raid)
	}

	if debuffs.WintersChill && targetIdx == 0 && !target.Env.IsForever() {
		aura := WintersChillAura(target)
		SchedulePeriodicDebuffApplication(aura, PeriodicActionOptions{
			Period:          time.Millisecond * 1500,
			NumTicks:        5,
			TickImmediately: true,
			Priority:        ActionPriorityDOT, // High prio
			OnAction: func(sim *Simulation) {
				aura.Activate(sim)
				if aura.IsActive() {
					aura.AddStack(sim)
				}
			},
		}, raid)
	}

	if debuffs.Stormstrike && targetIdx == 0 && !target.Env.IsForever() {
		ExternalStormstrikeCaster(debuffs, target)
	}

	if debuffs.GiftOfArthas {
		MakePermanent(GiftOfArthasAura(target))
	}

	/* if debuffs.Mangle {
		MakePermanent(MangleAura(target, level))
	} */

	if debuffs.CrystalYield {
		MakePermanent(CrystalYieldAura(target))
	}

	// Major Armor Debuffs
	if targetIdx == 0 {
		if debuffs.ExposeArmor != proto.TristateEffect_TristateEffectMissing {
			aura := ExposeArmorAura(target, TernaryInt32(debuffs.ExposeArmor == proto.TristateEffect_TristateEffectRegular, 0, 2))
			SchedulePeriodicDebuffApplication(aura, PeriodicActionOptions{
				Period:   time.Second * 3,
				NumTicks: 1,
				OnAction: func(sim *Simulation) {
					aura.Activate(sim)
				},
			}, raid)
		}

		if debuffs.SunderArmor {
			// Sunder Armor
			aura := SunderArmorAura(target)
			SchedulePeriodicDebuffApplication(aura, PeriodicActionOptions{
				Period:          time.Millisecond * 200,
				NumTicks:        5,
				TickImmediately: true,
				Priority:        ActionPriorityDOT, // High prio so it comes before actual warrior sunders.
				OnAction: func(sim *Simulation) {
					aura.Activate(sim)
					if aura.IsActive() {
						aura.AddStack(sim)
					}
				},
			}, raid)
		}
	}

	if debuffs.CurseOfRecklessness {
		MakePermanent(CurseOfRecklessnessAura(target))
	}

	if debuffs.FaerieFire {
		MakePermanent(FaerieFireAura(target))
	}

	if debuffs.CurseOfWeakness != proto.TristateEffect_TristateEffectMissing {
		MakePermanent(CurseOfWeaknessAura(target, GetTristateValueInt32(debuffs.CurseOfWeakness, 0, 3)))
	}

	if debuffs.DemoralizingRoar != proto.TristateEffect_TristateEffectMissing {
		MakePermanent(DemoralizingRoarAura(target))
	}
	if debuffs.DemoralizingShout != proto.TristateEffect_TristateEffectMissing {
		MakePermanent(DemoralizingShoutAura(target))
	}
	if debuffs.HuntersMark != proto.TristateEffect_TristateEffectMissing {
		MakePermanent(HuntersMarkAura(target, GetTristateValueInt32(debuffs.HuntersMark, 0, 5)))
	}

	// Atk spd reduction
	if debuffs.ThunderClap != proto.TristateEffect_TristateEffectMissing {
		// +5% from Warrior's Conqueror's Battlegear 5pc
		MakePermanent(ThunderClapAura(target, 8205, GetTristateValueInt32(debuffs.ThunderClap, 10, 15)))
	}
	if debuffs.Thunderfury {
		MakePermanent(ThunderfuryASAura(target))
	}

	// Miss
	if debuffs.InsectSwarm && targetIdx == 0 {
		MakePermanent(InsectSwarmAura(target))
	}
	if debuffs.ScorpidSting && targetIdx == 0 {
		MakePermanent(ScorpidStingAura(target))
	}
}

type StormstrikeConfig struct {
	stormstrikeFrequency     float64
	natureAttackersFrequency float64
}

func (character *Character) createStormstrikeConfig(player *proto.Player) {
	character.StormstrikeConfig = StormstrikeConfig{
		stormstrikeFrequency:     player.StormstrikeFrequency,
		natureAttackersFrequency: player.StormstrikeNatureAttackerFrequency,
	}
	// Defaults if not configured
	if character.StormstrikeConfig.stormstrikeFrequency == 0.0 {
		character.StormstrikeConfig.stormstrikeFrequency = 20.0
	}
}

const (
	StormstrikeCooldown = time.Second * 20
)

func ExternalStormstrikeCaster(_ *proto.Debuffs, target *Unit) {
	stormstrikeConfig := target.Env.Raid.Parties[0].Players[0].GetCharacter().StormstrikeConfig
	stormstrikeAura := StormstrikeAura(target)
	var pa *PendingAction
	MakePermanent(target.GetOrRegisterAura(Aura{
		Label: "Stormstrike External Proc Aura",
		OnGain: func(aura *Aura, sim *Simulation) {
			pa = NewPeriodicAction(sim, PeriodicActionOptions{
				Period:          DurationFromSeconds(stormstrikeConfig.stormstrikeFrequency),
				TickImmediately: true,
				OnAction: func(s *Simulation) {
					stormstrikeAura.Activate(sim)
					stormstrikeAura.SetStacks(sim, stormstrikeAura.MaxStacks)
				},
			})
			sim.AddPendingAction(pa)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			pa.Cancel(sim)
		},
	}))
}

func StormstrikeAura(unit *Unit) *Aura {
	stormstrikeConfig := unit.Env.Raid.Parties[0].Players[0].GetCharacter().StormstrikeConfig

	aura := unit.GetOrRegisterAura(Aura{
		Label:     "Stormstrike",
		ActionID:  ActionID{SpellID: 17364},
		Duration:  time.Second * 12,
		MaxStacks: 2,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexNature] *= 1.20
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexNature] /= 1.20
		},
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			// In Forever the debuff lasts its full duration instead of being consumed by Nature damage
			if sim.IsForever() {
				return
			}

			if aura.GetStacks() > 0 && spell.SpellSchool.Matches(SpellSchoolNature) && result.Landed() && result.Damage > 0 {
				aura.RemoveStack(sim)
			}
		},
	})

	// External attacks using nature strike
	if stormstrikeConfig.natureAttackersFrequency > 0 {
		aura.OnReset = func(aura *Aura, sim *Simulation) {
			sim.AddPendingAction(
				NewPeriodicAction(sim, PeriodicActionOptions{
					Period: DurationFromSeconds(stormstrikeConfig.natureAttackersFrequency),
					OnAction: func(s *Simulation) {
						if aura.GetStacks() > 0 {
							aura.RemoveStack(sim)
						}
					},
				}),
			)
		}
	}

	return aura
}

func ExternalIsbCaster(_ *proto.Debuffs, target *Unit) {
	isbConfig := target.Env.Raid.Parties[0].Players[0].GetCharacter().IsbConfig
	isbAura := ImprovedShadowBoltAura(target, 5)
	isbCrit := isbConfig.casterCrit / 100.0
	var pa *PendingAction
	MakePermanent(target.GetOrRegisterAura(Aura{
		Label: "Isb External Proc Aura",
		OnGain: func(aura *Aura, sim *Simulation) {
			pa = NewPeriodicAction(sim, PeriodicActionOptions{
				Period: DurationFromSeconds(isbConfig.shadowBoltFrequency),
				OnAction: func(s *Simulation) {
					for i := 0; i < int(isbConfig.isbWarlocks); i++ {
						if sim.Proc(isbCrit, "External Isb Crit") {
							isbAura.Activate(sim)
							isbAura.SetStacks(sim, isbAura.MaxStacks)
						} else if isbAura.IsActive() {
							isbAura.RemoveStack(sim)
						}
					}
				},
			})
			sim.AddPendingAction(pa)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			pa.Cancel(sim)
		},
	}))
}

type IsbConfig struct {
	shadowBoltFrequency float64
	casterCrit          float64
	isbWarlocks         int32
	isbShadowPriests    int32
}

func (character *Character) createIsbConfig(player *proto.Player) {
	character.IsbConfig = IsbConfig{
		shadowBoltFrequency: player.IsbSbFrequency,
		casterCrit:          player.IsbCrit,
		isbWarlocks:         player.IsbWarlocks,
		isbShadowPriests:    player.IsbSpriests,
	}
	//Defaults if not configured
	if character.IsbConfig.shadowBoltFrequency == 0.0 {
		character.IsbConfig.shadowBoltFrequency = 3.0
	}
	if character.IsbConfig.casterCrit == 0.0 {
		character.IsbConfig.casterCrit = 25.0
	}
	if character.IsbConfig.isbWarlocks == 0 {
		character.IsbConfig.isbWarlocks = 1
	}
}

const (
	ISBNumStacksBase = 4
)

func ImprovedShadowBoltAura(unit *Unit, rank int32) *Aura {
	isbLabel := "Improved Shadow Bolt"
	if unit.GetAura(isbLabel) != nil {
		return unit.GetAura(isbLabel)
	}

	isbConfig := unit.Env.Raid.Parties[0].Players[0].GetCharacter().IsbConfig

	priestGcds := []bool{false, true, true, true, true, true}
	priestCurGcd := 0
	externalShadowPriests := isbConfig.isbShadowPriests
	var priestPa *PendingAction

	damageMulti := 1. + 0.04*float64(rank)
	aura := unit.GetOrRegisterAura(Aura{
		Label:     isbLabel,
		ActionID:  ActionID{SpellID: 17800},
		Duration:  12 * time.Second,
		MaxStacks: ISBNumStacksBase,
		OnReset: func(aura *Aura, sim *Simulation) {
			// External shadow priests simulation
			if externalShadowPriests > 0 {
				priestCurGcd = 0
				priestPa = NewPeriodicAction(sim, PeriodicActionOptions{
					Period: GCDDefault,
					OnAction: func(s *Simulation) {
						if priestGcds[priestCurGcd] {
							for i := 0; i < int(externalShadowPriests); i++ {
								if aura.IsActive() {
									aura.RemoveStack(sim)
								}
							}
						}
						priestCurGcd++
						if priestCurGcd >= len(priestGcds) {
							priestCurGcd = 0
						}
					},
				})
				sim.AddPendingAction(priestPa)
			}
		},
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexShadow] *= damageMulti
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexShadow] /= damageMulti
		},
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			if spell.SpellSchool.Matches(SpellSchoolShadow) && result.Landed() && result.Damage > 0 {
				aura.RemoveStack(sim)
			}
		},
	})

	return aura
}

// Classic's Shadow Weaving is five ranks, 15257 and 15331 to 15334, and the raid debuff is the
// top one. Forever deleted 15331 to 15334 outright - the talent is a single spell, 15257, whose
// three tree ranks only move the proc chance - so the rank table has no meaning there and this
// aura is Classic only. The priest's own buff is 15258, in sim/priest/talents.go.
func ShadowWeavingAura(unit *Unit) *Aura {
	return unit.GetOrRegisterAura(Aura{
		Label:     "Shadow Weaving",
		ActionID:  ActionID{SpellID: 15334},
		Duration:  time.Second * 15,
		MaxStacks: 5,
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexShadow] /= 1.0 + 0.02*float64(oldStacks)
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexShadow] *= 1.0 + 0.02*float64(newStacks)
		},
	})
}

func SchedulePeriodicDebuffApplication(aura *Aura, options PeriodicActionOptions, _ *proto.Raid) {
	aura.OnReset = func(aura *Aura, sim *Simulation) {
		aura.Duration = NeverExpires
		StartPeriodicAction(sim, options)
	}
}

const JudgementAuraTag = "Judgement"

// TODO: Classic verify logic
func JudgementOfWisdomAura(target *Unit) *Aura {
	actionID := ActionID{SpellID: 20355}

	jowMana := 59.0

	return target.GetOrRegisterAura(Aura{
		Label:    "Judgement of Wisdom",
		ActionID: actionID,
		Tag:      JudgementAuraTag,
		Duration: time.Second * 40,
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			unit := spell.Unit
			if !unit.HasManaBar() {
				return
			}

			if spell.ProcMask.Matches(ProcMaskEmpty|ProcMaskProc|ProcMaskSpellDamageProc) && !spell.Flags.Matches(SpellFlagNotAProc) {
				return // Phantom spells (Romulo's, Lightning Capacitor, etc.) don't proc JoW.
			}

			if !spell.ProcMask.Matches(ProcMaskDirect) {
				return
			}

			// melee auto attacks don't even need to land
			if !result.Landed() && !spell.ProcMask.Matches(ProcMaskMeleeWhiteHit) {
				return
			}

			if sim.RandomFloat("jow") < 0.5 {
				if unit.JowManaMetrics == nil {
					unit.JowManaMetrics = unit.NewManaMetrics(actionID)
				}
				// JoW returns flat mana
				unit.AddMana(sim, jowMana, unit.JowManaMetrics)
			}
		},
	})
}

func JudgementOfLightAura(target *Unit) *Aura {
	actionID := ActionID{SpellID: 20271}

	return target.GetOrRegisterAura(Aura{
		Label:    "Judgement of Light",
		ActionID: actionID,
		Tag:      JudgementAuraTag,
		Duration: time.Second * 10,
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			if !spell.ProcMask.Matches(ProcMaskMelee) || !result.Landed() {
				return
			}
		},
	})
}

func JudgementOfTheCrusaderAura(caster *Unit, target *Unit, mult float64, extraBonus float64) *Aura {
	var spellId int32 = 20303
	var bonus float64 = 140

	bonus *= mult
	bonus += extraBonus

	return target.GetOrRegisterAura(Aura{
		Label:    "Judgement of the Crusader",
		ActionID: ActionID{SpellID: spellId},
		Tag:      JudgementAuraTag,
		Duration: 10 * time.Second,

		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] += bonus
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] -= bonus
		},
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			if spell.Unit != caster { // caster is nil for permanent auras
				return
			}
			if result.Landed() && spell.ProcMask.Matches(ProcMaskMelee) {
				aura.Refresh(sim)
			}
		},
	})
}

func CurseOfElementsAura(target *Unit) *Aura {
	resistance := 75.0
	dmgMod := 1.1
	forever := target.Env != nil && target.Env.IsForever()
	spellID := int32(11722)
	schools := []stats.SchoolIndex{stats.SchoolIndexFire, stats.SchoolIndexFrost}
	if forever {
		spellID = 1311680
		schools = []stats.SchoolIndex{stats.SchoolIndexArcane, stats.SchoolIndexFire, stats.SchoolIndexFrost, stats.SchoolIndexHoly, stats.SchoolIndexNature, stats.SchoolIndexShadow}
	}

	aura := target.GetOrRegisterAura(Aura{
		Label:    "Curse of Elements",
		ActionID: ActionID{SpellID: spellID},
		Duration: time.Minute * 5,
	})
	for _, school := range schools {
		spellSchoolDamageEffect(aura, school, dmgMod, 0.0, false)
		if school != stats.SchoolIndexHoly {
			spellSchoolResistanceEffect(aura, school, resistance, 0.0, false)
		}
	}

	return aura
}

func CurseOfShadowAura(target *Unit) *Aura {
	resistance := 75.0
	dmgMod := 1.1

	aura := target.GetOrRegisterAura(Aura{
		Label:    "Curse of Shadow",
		ActionID: ActionID{SpellID: 17937},
		Duration: time.Minute * 5,
	})
	spellSchoolDamageEffect(aura, stats.SchoolIndexArcane, dmgMod, 0.0, false)
	spellSchoolDamageEffect(aura, stats.SchoolIndexShadow, dmgMod, 0.0, false)

	spellSchoolResistanceEffect(aura, stats.SchoolIndexArcane, resistance, 0.0, false)
	spellSchoolResistanceEffect(aura, stats.SchoolIndexShadow, resistance, 0.0, false)

	return aura
}

func spellSchoolDamageEffect(aura *Aura, school stats.SchoolIndex, multiplier float64, extraPriority float64, exclusive bool) *ExclusiveEffect {
	return aura.NewExclusiveEffect("spellDamage"+strconv.Itoa(int(school)), exclusive, ExclusiveEffect{
		Priority: multiplier + extraPriority,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[school] *= multiplier
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[school] /= multiplier
		},
	})
}

func spellSchoolResistanceEffect(aura *Aura, school stats.SchoolIndex, amount float64, extraPriority float64, exclusive bool) *ExclusiveEffect {
	resistanceStat := map[stats.SchoolIndex]stats.Stat{
		stats.SchoolIndexArcane: stats.ArcaneResistance,
		stats.SchoolIndexFire:   stats.FireResistance,
		stats.SchoolIndexFrost:  stats.FrostResistance,
		stats.SchoolIndexNature: stats.NatureResistance,
		stats.SchoolIndexShadow: stats.ShadowResistance,
	}[school]
	return aura.NewExclusiveEffect("resistance"+strconv.Itoa(int(school)), exclusive, ExclusiveEffect{
		Priority: amount + extraPriority,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			aura.Unit.AddStatDynamic(sim, resistanceStat, -amount)
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			aura.Unit.AddStatDynamic(sim, resistanceStat, amount)
		},
	})
}

func GiftOfArthasAura(target *Unit) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:    "Gift of Arthas",
		ActionID: ActionID{SpellID: 11374},
		Duration: time.Minute * 3,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexPhysical] += 8
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexPhysical] -= 8
		},
	})
}

func MangleAura(target *Unit) *Aura {
	return bleedDamageAura(target, Aura{
		Label:    "Mangle",
		ActionID: ActionID{SpellID: 409828},
		Duration: time.Minute,
	}, 1.3)
}

// Bleed Damage Multiplier category
const BleedEffectCategory = "BleedDamage"

func bleedDamageAura(target *Unit, config Aura, multiplier float64) *Aura {
	aura := target.GetOrRegisterAura(config)
	aura.NewExclusiveEffect(BleedEffectCategory, true, ExclusiveEffect{
		Priority: multiplier,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.PseudoStats.BleedDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.PseudoStats.BleedDamageTakenMultiplier /= multiplier
		},
	})
	return aura
}

const SpellFirePowerEffectCategory = "spellFirePowerdebuff"

func ImprovedScorchAura(target *Unit) *Aura {
	aura := target.GetOrRegisterAura(Aura{
		Label:     "Improved Scorch",
		ActionID:  ActionID{SpellID: 12873},
		Duration:  time.Second * 30,
		MaxStacks: 5,
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexFire] /= 1 + .03*float64(oldStacks)
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexFire] *= 1 + .03*float64(newStacks)
		},
	})

	return aura
}

const SpellCritEffectCategory = "spellcritdebuff"

func WintersChillAura(target *Unit) *Aura {
	aura := target.GetOrRegisterAura(Aura{
		Label:     "Winter's Chill",
		ActionID:  ActionID{SpellID: 28593},
		Duration:  time.Second * 15,
		MaxStacks: 5,
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks, newStacks int32) {
			aura.Unit.PseudoStats.SchoolCritTakenChance[stats.SchoolIndexFrost] -= 0.02 * float64(oldStacks)
			aura.Unit.PseudoStats.SchoolCritTakenChance[stats.SchoolIndexFrost] += 0.02 * float64(newStacks)
		},
	})

	// effect = aura.NewExclusiveEffect(SpellCritEffectCategory, true, ExclusiveEffect{
	// 	Priority: 0,
	// 	OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
	// 		ee.Aura.Unit.PseudoStats.BonusSpellCritRatingTaken += ee.Priority * CritRatingPerCritChance
	// 	},
	// 	OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
	// 		ee.Aura.Unit.PseudoStats.BonusSpellCritRatingTaken -= ee.Priority * CritRatingPerCritChance
	// 	},
	// })
	return aura
}

var majorArmorReductionEffectCategory = "MajorArmorReduction"
var minorArmorReductionEffectCategory = "MinorArmorReduction"

func SunderArmorAura(target *Unit) *Aura {
	arpen := 450.0

	var effect *ExclusiveEffect
	aura := target.GetOrRegisterAura(Aura{
		Label:     "Sunder Armor",
		ActionID:  ActionID{SpellID: 11597},
		Duration:  time.Second * 30,
		MaxStacks: 5,
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			effect.SetPriority(sim, arpen*float64(newStacks))
		},
	})

	effect = aura.NewExclusiveEffect(majorArmorReductionEffectCategory, true, ExclusiveEffect{
		Priority: 0,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.Armor, -ee.Priority)
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.Armor, ee.Priority)
		},
	})

	return aura
}

func ExposeArmorAura(target *Unit, improvedEA int32) *Aura {
	spellID := int32(11198)
	arpen := 1700.0

	arpen *= []float64{1, 1.25, 1.5}[improvedEA]
	if target.Env != nil && target.Env.IsForever() {
		// Five combo points at 450 armor each. Forever's talent changes cost
		// and combo-point refunds, not the armor reduction.
		arpen = 2250
	}

	aura := target.GetOrRegisterAura(Aura{
		Label:    "ExposeArmor",
		ActionID: ActionID{SpellID: spellID},
		Duration: time.Second * 30,
	})

	aura.NewExclusiveEffect(majorArmorReductionEffectCategory, true, ExclusiveEffect{
		Priority: arpen,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.Armor, -ee.Priority)
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.Armor, ee.Priority)
		},
	})

	return aura
}

func CurseOfRecklessnessAura(target *Unit) *Aura {
	arpen := float64(505)
	ap := float64(90)
	if target.Env != nil && target.Env.IsForever() {
		ap = 0
	}

	aura := target.GetOrRegisterAura(Aura{
		Label:    "Curse of Recklessness",
		ActionID: ActionID{SpellID: 11717},
		Duration: time.Minute * 2,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.Armor, -arpen)
			aura.Unit.AddStatDynamic(sim, stats.AttackPower, ap)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.AddStatDynamic(sim, stats.Armor, arpen)
			aura.Unit.AddStatDynamic(sim, stats.AttackPower, -ap)
		},
	})
	return aura
}

// Decreases the armor of the target by X for 40 sec.
// Improved: Your Faerie Fire and Faerie Fire (Feral) also increase the chance for all attacks to hit that target by 1% for 40 sec.
func FaerieFireAura(target *Unit) *Aura {
	return faerieFireAuraInternal(target, "Faerie Fire", 9907)
}

func FaerieFireFeralAura(target *Unit) *Aura {
	return faerieFireAuraInternal(target, "Faerie Fire (Feral)", 17392)
}

func faerieFireAuraInternal(target *Unit, label string, spellID int32) *Aura {
	arPen := float64(505)

	aura := target.GetOrRegisterAura(Aura{
		Label:    label,
		ActionID: ActionID{SpellID: spellID},
		Duration: time.Second * 40,
	})

	aura.NewExclusiveEffect(minorArmorReductionEffectCategory, true, ExclusiveEffect{
		Priority: arPen,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.AddStatDynamic(sim, stats.Armor, -arPen)
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.AddStatDynamic(sim, stats.Armor, arPen)
		},
	})

	return aura
}

func CurseOfWeaknessAura(target *Unit, points int32) *Aura {
	modDmgReduction := -37.0

	modDmgReduction *= []float64{1, 1.06, 1.13, 1.20}[points]
	modDmgReduction = math.Floor(modDmgReduction)

	aura := target.GetOrRegisterAura(Aura{
		Label:    "Curse of Weakness" + strconv.Itoa(int(points)),
		ActionID: ActionID{SpellID: 11708},
		Duration: time.Minute * 2,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusPhysicalDamage += modDmgReduction
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusPhysicalDamage -= modDmgReduction
		},
	})
	return aura
}

const HuntersMarkAuraTag = "HuntersMark"

func HuntersMarkAura(target *Unit, points int32) *Aura {
	bonus := 71.0

	bonus *= 1 + 0.03*float64(points)

	aura := target.GetOrRegisterAura(Aura{
		Label:    "HuntersMark-" + strconv.Itoa(int(bonus)),
		Tag:      HuntersMarkAuraTag,
		ActionID: ActionID{SpellID: 14325},
		Duration: time.Minute * 2,
	})

	aura.NewExclusiveEffect("HuntersMark", true, ExclusiveEffect{
		Priority: bonus,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			for _, unit := range sim.Environment.Raid.AllUnits {
				if rangedAttackTable := unit.AttackTables[target.UnitIndex][proto.CastType_CastTypeRanged]; rangedAttackTable != nil {
					rangedAttackTable.BonusAttackPowerTaken += bonus
				}
			}
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			for _, unit := range sim.Environment.Raid.AllUnits {
				if rangedAttackTable := unit.AttackTables[target.UnitIndex][proto.CastType_CastTypeRanged]; rangedAttackTable != nil {
					rangedAttackTable.BonusAttackPowerTaken -= bonus
				}
			}
		},
	})

	return aura
}

func ExposeWeaknessAura(target *Unit) *Aura {
	bonus := 450.0

	aura := target.GetOrRegisterAura(Aura{
		ActionID: ActionID{SpellID: 23577},
		Label:    "Expose Weakness",
		Duration: time.Second * 7,
		OnGain: func(aura *Aura, sim *Simulation) {
			for _, unit := range sim.Environment.Raid.AllUnits {
				if rangedAttackTable := unit.AttackTables[target.UnitIndex][proto.CastType_CastTypeRanged]; rangedAttackTable != nil {
					rangedAttackTable.BonusAttackPowerTaken += bonus
				}
			}
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			for _, unit := range sim.Environment.Raid.AllUnits {
				if rangedAttackTable := unit.AttackTables[target.UnitIndex][proto.CastType_CastTypeRanged]; rangedAttackTable != nil {
					rangedAttackTable.BonusAttackPowerTaken -= bonus
				}
			}
		},
	})

	return aura
}

// Forever folds Improved Demoralizing Roar into the base value: the beta client's rank 5 takes
// 204 attack power at level 60 where Era's takes 138, and 204/138 is 1.48, exactly the 1+0.08x6
// the talent used to add. So the points must not be applied a second time on top.
func DemoralizingRoarAura(target *Unit) *Aura {
	baseAPReduction := 204.0

	aura := target.GetOrRegisterAura(Aura{
		Label:    "DemoralizingRoar",
		ActionID: ActionID{SpellID: 9898},
		Duration: time.Second * 30,
	})
	apReductionEffect(aura, baseAPReduction)
	return aura
}

const DemoralizingShoutRanks = 5

var DemoralizingShoutSpellId = [DemoralizingShoutRanks + 1]int32{0, 1160, 6190, 11554, 11555, 11556}
var DemoralizingShoutBaseAP = [DemoralizingShoutRanks + 1]float64{0, 63, 78, 106, 155, 204}
var DemoralizingShoutLevel = [DemoralizingShoutRanks + 1]int{0, 14, 24, 34, 44, 54}

// Same fold-in as Demoralizing Roar above. The beta client's rank 5 takes 204 attack power at
// level 60 where Era's takes 146, and 204/146 is 1.4000, exactly the 1+0.08x5 Improved
// Demoralizing Shout used to add. Booming Voice only widens the radius in Forever.
func DemoralizingShoutAura(target *Unit) *Aura {
	rank := int32(5)
	spellId := DemoralizingShoutSpellId[rank]
	baseAPReduction := DemoralizingShoutBaseAP[rank]

	aura := target.GetOrRegisterAura(Aura{
		Label:    "DemoralizingShout",
		ActionID: ActionID{SpellID: spellId},
		Duration: time.Second * 45,
	})
	apReductionEffect(aura, baseAPReduction)
	return aura
}

func VindicationAura(target *Unit, points int32, _ int32) *Aura {
	aura := target.GetOrRegisterAura(Aura{
		Label:    "Vindication",
		ActionID: ActionID{SpellID: 26016},
		Duration: time.Second * 10,
	})
	return aura
}

func apReductionEffect(aura *Aura, apReduction float64) *ExclusiveEffect {
	statReduction := stats.Stats{stats.AttackPower: -apReduction}
	return aura.NewExclusiveEffect("APReduction", false, ExclusiveEffect{
		Priority: apReduction,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.AddStatsDynamic(sim, statReduction)
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.AddStatsDynamic(sim, statReduction.Invert())
		},
	})
}

func ThunderClapAura(target *Unit, spellID int32, atkSpeedReductionPercent int32) *Aura {
	aura := target.GetOrRegisterAura(Aura{
		Label:    "ThunderClap-" + strconv.Itoa(int(atkSpeedReductionPercent)),
		ActionID: ActionID{SpellID: spellID},
		Duration: time.Second * 30,
	})
	AtkSpeedReductionEffect(aura, 1+0.01*float64(atkSpeedReductionPercent))
	return aura
}

func ThunderfuryASAura(target *Unit) *Aura {
	aura := target.GetOrRegisterAura(Aura{
		Label:    "Thunderfury",
		ActionID: ActionID{SpellID: 21992},
		Duration: time.Second * 12,
	})
	AtkSpeedReductionEffect(aura, 1.2)
	return aura
}

func AtkSpeedReductionEffect(aura *Aura, speedMultiplier float64) *ExclusiveEffect {
	return aura.NewExclusiveEffect("AtkSpdReduction", false, ExclusiveEffect{
		Priority: speedMultiplier,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.MultiplyAttackSpeed(sim, 1/speedMultiplier)
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.MultiplyAttackSpeed(sim, speedMultiplier)
		},
	})
}

func InsectSwarmAura(target *Unit) *Aura {
	aura := target.GetOrRegisterAura(Aura{
		Label:    "InsectSwarmMiss",
		ActionID: ActionID{SpellID: 24977},
		Duration: time.Second * 12,
	})
	increasedMissEffect(aura, 0.02)
	return aura
}

func ScorpidStingAura(target *Unit) *Aura {
	aura := target.GetOrRegisterAura(Aura{
		Label:    "Scorpid Sting",
		ActionID: ActionID{SpellID: 3043},
		Duration: time.Second * 20,
	})
	return aura
}

func increasedMissEffect(aura *Aura, increasedMissChance float64) *ExclusiveEffect {
	return aura.NewExclusiveEffect("IncreasedMiss", false, ExclusiveEffect{
		Priority: increasedMissChance,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.PseudoStats.IncreasedMissChance += increasedMissChance
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.PseudoStats.IncreasedMissChance -= increasedMissChance
		},
	})
}

func CrystalYieldAura(target *Unit) *Aura {
	return target.GetOrRegisterAura(Aura{
		Label:    "Crystal Yield",
		ActionID: ActionID{SpellID: 15235},
		Duration: 2 * time.Minute,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.stats[stats.Armor] -= 200
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.stats[stats.Armor] += 200
		},
	})
}
