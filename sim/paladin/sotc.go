package paladin

import (
	"fmt"
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

func (paladin *Paladin) registerSealOfTheCrusader() {
	type judge struct {
		spellID int32
		bonus   float64
	}

	var ranks = []struct {
		level      int32
		spellID    int32
		manaCost   float64
		scaleLevel int32
		ap         float64
		scale      float64
		judge      judge
	}{
		{level: 6, spellID: 21082, manaCost: 25, scaleLevel: 12, ap: 31, scale: 0.7, judge: judge{spellID: 21183, bonus: 23}},
		{level: 12, spellID: 20162, manaCost: 40, scaleLevel: 20, ap: 51, scale: 1.1, judge: judge{spellID: 20188, bonus: 35}},
		{level: 22, spellID: 20305, manaCost: 65, scaleLevel: 30, ap: 94, scale: 1.7, judge: judge{spellID: 20300, bonus: 58}},
		{level: 32, spellID: 20306, manaCost: 90, scaleLevel: 40, ap: 145, scale: 2, judge: judge{spellID: 20301, bonus: 92}},
		{level: 42, spellID: 20307, manaCost: 125, scaleLevel: 50, ap: 221, scale: 2.2, judge: judge{spellID: 20302, bonus: 127}},
		{level: 52, spellID: 20308, manaCost: 160, scaleLevel: 60, ap: 306, scale: 2.4, judge: judge{spellID: 20303, bonus: 161}},
	}

	// Beta client 1.60.1.69893: Improved Seal of the Crusader's 15% is baked into the judgement
	// (rank 6 140 -> 161, which core.JudgementOfTheCrusaderAura builds as 140 x 1.15) but not into
	// the seal, whose attack power is Classic's at every rank. The debuff lasts 40 sec.
	const improvedSotC = 1.15

	var libramAp, libramBonus float64
	if paladin.Ranged().ID == LibramOfFervor {
		libramAp = 48
		libramBonus = 33
	}

	for i, rank := range ranks {
		rank := rank
		if paladin.Level < rank.level {
			break
		}

		debuffs := paladin.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			if paladin.Env.IsForever() {
				bonus := rank.judge.bonus + libramBonus
				aura := target.GetOrRegisterAura(core.Aura{
					Label:    fmt.Sprintf("Judgement of the Crusader %d %s", rank.judge.spellID, paladin.Label),
					ActionID: core.ActionID{SpellID: rank.judge.spellID},
					Tag:      core.JudgementAuraTag,
					Duration: 40 * time.Second,
					OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
						if spell.Unit == &paladin.Unit && result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) {
							aura.Refresh(sim)
						}
					},
				})
				aura.NewExclusiveEffect(core.JudgementOfTheCrusaderCategory, true, core.ExclusiveEffect{
					Priority: bonus,
					OnGain: func(_ *core.ExclusiveEffect, _ *core.Simulation) {
						target.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] += bonus
					},
					OnExpire: func(_ *core.ExclusiveEffect, _ *core.Simulation) {
						target.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] -= bonus
					},
				})
				return aura
			}
			return core.JudgementOfTheCrusaderAura(&paladin.Unit, target, improvedSotC, libramBonus)
		})

		judgeSpell := paladin.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: rank.judge.spellID},
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagMeleeMetrics,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealOutcome(sim, target, spell.OutcomeAlwaysHit)
				debuffs.Get(target).Activate(sim)
			},
		})

		ap := rank.ap + rank.scale*float64(min(paladin.Level, rank.scaleLevel)-rank.level)

		aura := paladin.RegisterAura(core.Aura{
			Label:    "Seal of the Crusader" + paladin.Label + strconv.Itoa(i+1),
			ActionID: core.ActionID{SpellID: rank.spellID},
			Duration: time.Second * 30,
			OnGain: func(_ *core.Aura, sim *core.Simulation) {
				paladin.MultiplyMeleeSpeed(sim, 1.4)
				paladin.AutoAttacks.MHAuto().DamageMultiplier /= 1.4
				paladin.AddStatDynamic(sim, stats.AttackPower, ap+libramAp)
			},
			OnExpire: func(_ *core.Aura, sim *core.Simulation) {
				paladin.MultiplyMeleeSpeed(sim, 1/1.4)
				paladin.AutoAttacks.MHAuto().DamageMultiplier *= 1.4
				paladin.AddStatDynamic(sim, stats.AttackPower, -(ap + libramAp))
			},
		})

		paladin.aurasSotC = append(paladin.aurasSotC, aura)

		paladin.RegisterSpell(core.SpellConfig{
			ActionID:    aura.ActionID,
			SpellSchool: core.SpellSchoolHoly,
			Flags:       core.SpellFlagAPL,

			RequiredLevel: int(rank.level),
			Rank:          i + 1,

			ManaCost: core.ManaCostOptions{
				FlatCost:   rank.manaCost - paladin.getLibramSealCostReduction(),
				Multiplier: paladin.sealCostMultiplier(),
			},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
			},

			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
				paladin.applySeal(aura, spell, judgeSpell, sim)
			},
		})

		paladin.spellsJotC = append(paladin.spellsJotC, judgeSpell)
	}
}
