package paladin

import (
	"math"
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// PROVISIONAL modeling convention, not an independently verified server script:
// one replacement pool,30s from each proc. Keep this separate from the sourced
// seal's30s duration. Controlled sensitivity tests may vary SelfShield().Duration
// on the registered tagged shield spell, without changing seal uptime.
const sealOfFuryShieldDuration = 30 * time.Second

type furyRank struct {
	level, cap                                  int32
	seal, proc, judge                           int32
	cost, damage, judgeBase, judgePPL, variance float64
}

var furyRanks = []furyRank{
	{10, 16, 1311649, 1311647, 1311650, 40, 6, 23, 1.71000003815, .10000000149},
	{18, 24, 1311656, 1311654, 1311655, 60, 9, 36.90000152588, 2.16000008583, .10000000149},
	{25, 31, 20163, 20231, 20183, 90, 14, 54, 2.51999998093, .10000000149},
	{34, 40, 20419, 20415, 20411, 120, 19, 74, 2.78999996185, .09756097198},
	{42, 48, 20421, 20416, 20412, 140, 25, 96, 3.42000007629, .09756097198},
	{50, 56, 20422, 20417, 20413, 170, 32, 123, 3.69000005722, .08759099990},
	{58, 64, 20423, 20418, 20414, 200, 35, 153, 3.69000005722, .08759099990},
}

func (paladin *Paladin) registerSealOfFury() {
	if !paladin.Env.IsForever() {
		return
	}
	var highest int32
	for _, r := range furyRanks {
		if paladin.Level >= r.level {
			highest = r.seal
		}
	}
	if highest == 0 {
		return
	}
	duration := sealOfFuryShieldDuration
	if seconds := paladin.Options.SealOfFuryShieldDurationSeconds; seconds != nil {
		if math.IsNaN(*seconds) || math.IsInf(*seconds, 0) || *seconds < 0 || *seconds > 3600 {
			panic("Provisional Seal of Fury shield duration must be finite and within0..3600seconds")
		}
		duration = time.Duration(*seconds * float64(time.Second))
	}
	// One shared pool across every registered rank: switching/downranking cannot
	// accumulate independent shields. There is no invented Judgement shield/cap.
	shield := paladin.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{SpellID: highest, Tag: 1}, SpellSchool: core.SpellSchoolHoly,
		ProcMask: core.ProcMaskEmpty, Flags: core.SpellFlagPassiveSpell | core.SpellFlagNoOnCastComplete,
		DamageMultiplier: 1, ThreatMultiplier: 0,
		Shield: core.ShieldConfig{SelfOnly: true, Aura: core.Aura{Label: "Seal of Fury Shield (provisional lifetime)", Duration: duration},
			OnDepleted: func(sim *core.Simulation, incoming *core.Spell, _ *core.SpellResult) {
				if !paladin.Talents.ImprovedSealOfFury {
					return
				}
				levels := min(int32(3), max(int32(0), incoming.Unit.Level-paladin.Level))
				paladin.AddMana(sim, float64(paladin.Level)*(1+.15*float64(levels)), paladin.furyManaMetrics)
			}},
	}).SelfShield()
	paladin.furyShield = shield
	if paladin.Talents.ImprovedSealOfFury {
		paladin.furyManaMetrics = paladin.NewManaMetrics(core.ActionID{SpellID: 1314104})
		paladin.furyManaMetrics.NoThreat = true
	}
	taunts := paladin.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return target.GetOrRegisterAura(core.Aura{Label: "Judgement of Fury Taunt-" + strconv.Itoa(int(paladin.Index)), ActionID: core.ActionID{SpellID: 20232}, Duration: 4 * time.Second})
	})
	for index, rank := range furyRanks {
		if paladin.Level < rank.level {
			break
		}
		mean := rank.judgeBase + rank.judgePPL*float64(min(paladin.Level, rank.cap)-rank.level)
		judgement := paladin.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: rank.judge}, SpellSchool: core.SpellSchoolHoly, DefenseType: core.DefenseTypeMelee, ProcMask: core.ProcMaskMeleeMHSpecial,
			RequiredLevel: int(rank.level), Rank: index + 1,
			Flags:            core.SpellFlagMeleeMetrics | core.SpellFlagSuppressWeaponProcs | core.SpellFlagSuppressEquipProcs | core.SpellFlagBinary,
			DamageMultiplier: paladin.improvedSeals(), ThreatMultiplier: 1, BonusCoefficient: .44999998808,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcDamage(sim, target, sim.Roll(mean*(1-rank.variance/2), mean*(1+rank.variance/2)), spell.OutcomeMeleeSpecialNoBlockDodgeParry)
				spell.DealDamage(sim, result)
				if result.Landed() {
					taunts.Get(target).Activate(sim)
				}
			},
		})
		proc := paladin.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: rank.proc}, SpellSchool: core.SpellSchoolHoly, DefenseType: core.DefenseTypeMelee, ProcMask: core.ProcMaskMeleeMHSpecial,
			RequiredLevel: int(rank.level), Rank: index + 1,
			Flags:            core.SpellFlagMeleeMetrics | core.SpellFlagPassiveSpell | core.SpellFlagSuppressEquipProcs,
			DamageMultiplier: paladin.improvedSeals(), ThreatMultiplier: 1, BonusCoefficient: .10000000149,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcDamage(sim, target, rank.damage, spell.OutcomeMeleeSpecialCritOnly)
				spell.DealDamage(sim, result)
				if shield.Duration == 0 {
					shield.Deactivate(sim)
				} else if paladin.OffHand().WeaponType == proto.WeaponType_WeaponTypeShield && result.Damage > 0 {
					shield.Apply(sim, .5*result.Damage)
				}
			},
		})
		aura := paladin.RegisterAura(core.Aura{Label: "Seal of Fury" + paladin.Label + strconv.Itoa(index+1), ActionID: core.ActionID{SpellID: rank.seal}, Duration: 30 * time.Second,
			OnSpellHitDealt: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMeleeWhiteHit) {
					proc.Cast(sim, result.Target)
				}
			},
		})
		paladin.aurasSoF = append(paladin.aurasSoF, aura)
		paladin.registerSealProc(aura, proc, echoOfFury)
		paladin.sealOfFury = paladin.RegisterSpell(core.SpellConfig{ActionID: aura.ActionID, SpellSchool: core.SpellSchoolHoly, Flags: core.SpellFlagAPL, RequiredLevel: int(rank.level), Rank: index + 1,
			ManaCost: core.ManaCostOptions{FlatCost: rank.cost - paladin.getLibramSealCostReduction(), Multiplier: paladin.sealCostMultiplier()}, Cast: core.CastConfig{DefaultCast: core.Cast{GCD: core.GCDDefault}},
			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
				paladin.applySeal(aura, spell, judgement, sim)
			},
		})
		paladin.spellsJoF = append(paladin.spellsJoF, judgement)
	}
}
