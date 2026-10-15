package planner

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/warrior"
	"github.com/wowsims/classic/sim/warrior/ranks"
)

// BuildWarriorModel reads a constructed (not yet reset) simulator Warrior: its registered spells carry the talented
// costs, cooldowns and multipliers, and its unit carries the static stats. The planner's damage formulas mirror the
// spells' ApplyEffects; tests compare the two.
func BuildWarriorModel(w *warrior.Warrior, known map[int32]bool) *Model {
	t := w.Talents
	m := &Model{
		Level:           w.Level,
		Class:           "WARRIOR",
		Race:            w.Race.String(),
		AP:              w.GetStat(stats.AttackPower),
		CritPct:         w.GetStat(stats.MeleeCrit) / core.CritRatingPerCritChance,
		HitPct:          w.GetStat(stats.MeleeHit) / core.MeleeHitRatingPerHitChance,
		ExpertisePct:    w.GetStat(stats.Expertise) / core.ExpertiseRatingPerExpertiseChance,
		ArmorIgnore:     w.PseudoStats.ArmorIgnorePercent,
		PhysMult:        w.PseudoStats.DamageDealtMultiplier * w.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexPhysical],
		WhiteCritMult:   2,
		SpecialCritMult: 2 + 2*0.1*float64(t.Impale),
		Health:          w.GetStat(stats.Health),
		Armor:           w.GetStat(stats.Armor),
		DodgePct:        w.GetStat(stats.Dodge) / core.DodgeRatingPerDodgeChance,
		ParryPct:        w.GetStat(stats.Parry) / core.ParryRatingPerParryChance,
		BlockPct:        w.GetStat(stats.Block) / core.BlockRatingPerBlockChance,
		BlockValue:      w.BlockValue(),
		MissPct:         5,
		TacticalMastery: 10 + 3*float64(t.ImprovedTacticalMastery),
		OtherStanceMult: 0.8,
		AngerManagement: t.AngerManagement,
	}
	m.Shield = w.PseudoStats.CanBlock
	defiance := 0.0
	if m.Shield {
		defiance = 0.05 * float64(t.Defiance)
	}
	m.DefStanceThreat = 1.3 * (1 + defiance)
	if t.Flurry > 0 {
		m.FlurryHaste = 0.05 * float64(t.Flurry)
	}
	if t.Enrage > 0 {
		m.EnragePct, m.EnrageChance = 0.02*float64(t.Enrage), 0.3
	}
	m.DeepWoundsPct = 0.2 * float64(t.DeepWounds)
	m.UWChance = 0.12 * float64(t.UnbridledWrath)
	if m.Shield {
		m.ShieldSpecRage = 5 * 0.2 * float64(t.ShieldSpecialization)
	}
	if p := w.ForeverIncomingRageParameters(); p.ExpectedHealth > 0 {
		m.IncomingRagePerD = p.Coefficient / p.ExpectedHealth / (1 - p.ReferenceArmor)
	}

	mh := w.AutoAttacks.MH()
	m.MH = Weapon{Present: mh != nil && mh.SwingSpeed > 0, DamageMult: 1, CritBonus: 0}
	if m.MH.Present {
		m.MH.Min, m.MH.Max, m.MH.Speed, m.MH.NormSpeed = mh.BaseDamageMin, mh.BaseDamageMax, mh.SwingSpeed, mh.NormalizedSwingSpeed
		if item := w.GetMHWeapon(); item != nil {
			m.MH.TwoHand = item.HandType == proto.HandType_HandTypeTwoHand
		}
		rate := 4.5 / 1.3
		if m.MH.TwoHand {
			rate = 4.5
		}
		m.MH.RagePerLanded = m.MH.Speed * rate
	}
	if w.AutoAttacks.IsDualWielding {
		oh := w.AutoAttacks.OH()
		m.DW = true
		m.OH = Weapon{Present: true, Min: oh.BaseDamageMin, Max: oh.BaseDamageMax, Speed: oh.SwingSpeed, NormSpeed: oh.NormalizedSwingSpeed,
			DamageMult: 0.5 * (1 + 0.05*float64(t.DualWieldSpecialization)),
			HitBonus:   []float64{0, 4, 7, 10}[min(t.FuriousPrecision, 3)]}
		m.OH.RagePerLanded = m.OH.Speed * 4.5 / 1.3 * 0.5 * (1 + 0.1*float64(t.DualWieldSpecialization))
	}

	add := func(name ranks.Ability, ws *warrior.WarriorSpell, fill func(ab *Ability)) {
		rank := (ranks.Selector{Level: w.Level, Known: known}).Rank(name)
		ab := &Ability{Name: name, Rank: rank, CritMult: m.SpecialCritMult, DamageMult: 1, ThreatMult: 1, GCD: true, Modelled: true}
		if rank != nil {
			ab.GameID = rank.ID
		}
		if ws == nil || ws.Spell == nil {
			if rank == nil {
				return // not trained: not an ability of this character
			}
			ab.Modelled = false
			ab.Reason = "not available to the simulator for this character"
			m.Abilities = append(m.Abilities, ab)
			return
		}
		sp := ws.Spell
		ab.SpellID = sp.ActionID.SpellID
		if rank == nil {
			ab.GameID = ab.SpellID
		}
		if sp.Cost != nil {
			ab.Cost = sp.Cost.ApplyCostModifiers(sp.Cost.BaseCost)
			if rc, ok := sp.Cost.SpellCostFunctions.(*core.RageCost); ok {
				ab.Refund = rc.Refund
			}
		}
		ab.CD = sp.CD.Duration.Seconds()
		ab.CD = max(ab.CD, sp.SharedCD.Duration.Seconds())
		ab.GCD = sp.DefaultCast.GCD > 0
		ab.Cast = sp.DefaultCast.CastTime.Seconds()
		ab.DamageMult = sp.DamageMultiplier
		ab.ThreatMult = sp.ThreatMultiplier
		ab.FlatThreat = sp.FlatThreatBonus
		ab.BonusCrit = sp.BonusCritRating / core.CritRatingPerCritChance
		mask := ws.GetStanceMask()
		if mask != warrior.AnyStance {
			for s, bit := range map[int]warrior.Stance{StanceBattle: warrior.BattleStance, StanceDefensive: warrior.DefensiveStance, StanceBerserker: warrior.BerserkerStance} {
				if mask.Matches(bit) {
					ab.Stances |= 1 << s
				}
			}
		}
		fill(ab)
		m.Abilities = append(m.Abilities, ab)
	}
	weapon := func(ab *Ability, bonus float64, normalized bool) {
		ab.Flat, ab.WeaponMult, ab.Normalized = bonus, 1, normalized
	}

	add(ranks.HeroicStrike, w.HeroicStrike, func(ab *Ability) {
		ab.Kind, ab.GCD, ab.OffGCD = KindQueued, false, true
		weapon(ab, ab.Rank.Find(ranks.EffectWeaponDamageBonus, 0).Low, false)
	})
	add(ranks.Cleave, w.Cleave, func(ab *Ability) {
		ab.Kind, ab.GCD, ab.OffGCD, ab.Targets = KindQueued, false, true, 2
		weapon(ab, ab.Rank.Find(ranks.EffectWeaponDamageBonus, 0).Low, false)
	})
	add(ranks.SunderArmor, w.SunderArmor, func(ab *Ability) {
		ab.Kind, ab.Table, ab.MaxStacks, ab.Duration = KindDebuff, TableNoDamage, 5, 30
		ab.ArmorPerStack = -ab.Rank.Find(ranks.EffectApplyAura, ranks.AuraModArmorPercentOrNeg).Low
		ab.FlatThreat = ab.Rank.Find(ranks.EffectThreat, 0).Low
		ab.ThreatAP = ranks.SunderThreatAP(ab.Rank)
	})
	add(ranks.Revenge, w.Revenge, func(ab *Ability) {
		d := ab.Rank.Find(ranks.EffectSchoolDamage, 0)
		ab.Kind, ab.Flat, ab.APCoef, ab.NeedsWindow = KindStrike, (d.Low+d.High)/2, ranks.RevengeAPCoefficient, windowRevenge
	})
	add(ranks.ThunderClap, w.ThunderClap, func(ab *Ability) {
		ab.Kind, ab.Table, ab.Targets = KindAoE, TableMagic, 4
		ab.Flat, ab.APCoef, ab.Duration, ab.SlowPct = ab.Rank.Find(ranks.EffectSchoolDamage, 0).Low, ranks.ThunderClapAPCoefficient, 30, 20
		ab.CritMult = 1 + (1.5-1)*sp(w.ThunderClap).CritDamageBonus
	})
	add(ranks.DemoralizingShout, w.DemoralizingShout, func(ab *Ability) {
		ab.Kind, ab.Table, ab.Targets, ab.Duration = KindDebuff, TableNone, 99, 45
		ab.APReduction = -ab.Rank.Find(ranks.EffectApplyAura, ranks.AuraModAttackPower).Low
	})
	add(ranks.BattleShout, w.BattleShout, func(ab *Ability) {
		ab.Kind, ab.Table, ab.Duration = KindBuff, TableNone, 180
		ab.BuffAP = ab.Rank.Find(ranks.EffectApplyAura, ranks.AuraModAttackPower).Low
	})
	add(ranks.Bloodthirst, w.Bloodthirst, func(ab *Ability) {
		ab.Kind, ab.Flat, ab.APCoef = KindStrike, ab.Rank.Find(ranks.EffectSchoolDamage, 0).Low, bloodthirstAP
	})
	add(ranks.MortalStrike, w.MortalStrike, func(ab *Ability) {
		ab.Kind = KindStrike
		weapon(ab, ab.Rank.Find(ranks.EffectNormalizedWeapon, 0).Low, true)
	})
	add(ranks.Whirlwind, w.Whirlwind, func(ab *Ability) {
		ab.Kind, ab.Targets = KindAoE, 4
		weapon(ab, 0, true)
	})
	add(ranks.Slam, w.Slam, func(ab *Ability) {
		ab.Kind = KindStrike
		weapon(ab, ab.Rank.Find(ranks.EffectWeaponDamageBonus, 0).Low, false)
		ab.KeepSwings = w.Talents.ImprovedSlam > 0
	})
	add(ranks.Overpower, w.Overpower, func(ab *Ability) {
		ab.Kind, ab.Table, ab.NeedsWindow = KindStrike, TableNoAvoid, windowOverpower
		weapon(ab, ab.Rank.Find(ranks.EffectNormalizedWeapon, 0).Low, true)
	})
	add(ranks.Execute, w.Execute, func(ab *Ability) {
		ab.Kind, ab.Flat, ab.PerRage = KindExecute, ab.Rank.Find(ranks.EffectDummy, 0).Low, ranks.ExecuteDamagePerRage(ab.Rank)
	})
	add(ranks.Rend, w.Rend, func(ab *Ability) {
		tick := ab.Rank.Find(ranks.EffectApplyAura, ranks.AuraPeriodicDamage)
		ab.Kind, ab.TickDamage, ab.TickAPCoef = KindDot, tick.Low, ranks.RendAPCoefficient
		ab.Ticks = int(ab.Rank.DurationMs / max(tick.PeriodMs, 1))
		ab.Duration = float64(ab.Rank.DurationMs) / 1000
	})
	add(ranks.ShieldSlam, w.ShieldSlam, func(ab *Ability) {
		d := ab.Rank.Find(ranks.EffectSchoolDamage, 0)
		ab.Kind, ab.Flat, ab.NeedsShield = KindStrike, (d.Low+d.High)/2+m.BlockValue, true
	})
	add(ranks.Hamstring, w.Hamstring, func(ab *Ability) {
		ab.Kind, ab.Flat = KindStrike, ab.Rank.Find(ranks.EffectSchoolDamage, 0).Low
	})
	add(ranks.Bloodrage, w.Bloodrage, func(br *Ability) {
		br.Kind, br.Table, br.GCD, br.OffGCD = KindRage, TableNone, false, true
		mult := 1 + 0.25*float64(t.ImprovedBloodrage)
		br.RageGain, br.RagePerSec, br.Duration = 10*mult, 1*mult, 10
	})
	add(ranks.ShieldBlock, w.ShieldBlock, func(ab *Ability) {
		ab.Kind, ab.Table, ab.GCD, ab.OffGCD, ab.NeedsShield = KindBlock, TableNone, false, true, true
		ab.BlockPct, ab.BlockCharges, ab.Duration = 75, 2, 7
	})
	for _, st := range []struct {
		name ranks.Ability
		ws   *warrior.WarriorSpell
		to   int
	}{{ranks.BattleStance, w.BattleStance, StanceBattle}, {ranks.DefensiveStance, w.DefensiveStance, StanceDefensive}, {ranks.BerserkerStance, w.BerserkerStance, StanceBerserker}} {
		st := st
		add(st.name, st.ws, func(ab *Ability) {
			ab.Kind, ab.Table, ab.GCD, ab.OffGCD, ab.ToStance, ab.CD = KindStance, TableNone, false, true, st.to, 1
		})
	}
	// Trained abilities the forward model does not resolve: reported, never chosen.
	for _, name := range []ranks.Ability{ranks.VictoryRush, ranks.Charge, ranks.Intercept, ranks.SpearingStrike, ranks.Pummel, ranks.ShieldBash,
		ranks.Disarm, ranks.PiercingHowl, ranks.MockingBlow, ranks.Taunt, ranks.ChallengingShout, ranks.IntimidatingShout,
		ranks.BerserkerRage, ranks.Recklessness, ranks.Retaliation, ranks.ShieldWall, ranks.ConcussionBlow, ranks.SweepingStrikes,
		ranks.DeathWish, ranks.LastStand} {
		if m.Get(name) != nil {
			continue
		}
		rank := (ranks.Selector{Level: w.Level, Known: known}).Rank(name)
		talentOnly := map[ranks.Ability]bool{ranks.DeathWish: !t.DeathWish, ranks.LastStand: !t.LastStand, ranks.SweepingStrikes: !t.SweepingStrikes,
			ranks.ConcussionBlow: !t.ConcussionBlow, ranks.SpearingStrike: !t.SpearingStrike}
		if rank == nil || talentOnly[name] {
			continue
		}
		m.Abilities = append(m.Abilities, &Ability{Name: name, GameID: rank.ID, Rank: rank, Modelled: false,
			Reason: "manual or not modelled by the planner (utility, movement, kills or a major cooldown)"})
		m.index()
	}
	m.index()
	return m
}

// bloodthirstAP is the Forever attack power coefficient of Bloodthirst (October 1), read from the engine's formula.
const bloodthirstAP = 0.45

func sp(ws *warrior.WarriorSpell) *core.Spell {
	if ws == nil {
		return &core.Spell{}
	}
	return ws.Spell
}

func boolf(b bool) float64 {
	if b {
		return 1
	}
	return 0
}
