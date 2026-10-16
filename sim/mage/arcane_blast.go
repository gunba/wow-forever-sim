package mage

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

const ArcaneBlastMaxStacks = 4

func (mage *Mage) registerArcaneBlastSpell() {
	if !mage.Talents.ArcaneBlast {
		return
	}

	// Client 1.60.1.70425 redesign: rank 5 (1239700) is a 1.5 sec cast for 218-254 at 0.429 (damage and
	// coefficient cut 40%). Each cast grants 15% of Intellect as Spell Damage and raises Arcane Blast's
	// cost 175%, stacking four times; the buff (400573) lasts 15 sec from its first application and is
	// no longer spent by other spells. 30451 is kept as the action id the APLs already name.
	baseDamage := []float64{218, 254}
	spellCoeff := .429
	baseManaCost := .15
	castTime := time.Millisecond * 1500

	actionID := core.ActionID{SpellID: 30451}

	var deps [ArcaneBlastMaxStacks + 1]*stats.StatDependency
	for i := 1; i <= ArcaneBlastMaxStacks; i++ {
		deps[i] = mage.NewDynamicStatDependency(stats.Intellect, stats.SpellDamage, 0.15*float64(i))
	}

	mage.ArcaneBlastAura = mage.RegisterAura(core.Aura{
		Label:     "Arcane Blast",
		ActionID:  actionID,
		Duration:  time.Second * 15,
		MaxStacks: ArcaneBlastMaxStacks,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			if oldStacks > 0 {
				mage.DisableDynamicStatDep(sim, deps[oldStacks])
			}
			if newStacks > 0 {
				mage.EnableDynamicStatDep(sim, deps[newStacks])
			}
			mage.ArcaneBlast.Cost.Multiplier += 175 * (newStacks - oldStacks)
		},
	})

	mage.ArcaneBlast = mage.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_MageArcaneBlast,
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolArcane,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagMage | core.SpellFlagAPL,

		RequiredLevel: 60,
		Rank:          1,

		ManaCost: core.ManaCostOptions{
			BaseCost: baseManaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: castTime,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := sim.Roll(baseDamage[0], baseDamage[1])
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicHitAndCrit)

			// The duration runs from the first stack; later casts add stacks without refreshing it.
			if !mage.ArcaneBlastAura.IsActive() {
				mage.ArcaneBlastAura.Activate(sim)
			}
			mage.ArcaneBlastAura.AddStack(sim)
		},
	})
}

// Arcane Missiles no longer spends the Arcane Blast stacks (client 1.60.1.70425).
func (mage *Mage) spendArcaneBlastStacks(_ *core.Simulation) {}

// Arcane Wrath (client 1.60.1.70425), the 16-point Arcane talent: a 3.5 sec cast, 6 sec cooldown, that
// gains Spell Damage equal to 5% of maximum mana. Rank 4 (1326082): 406-496 at coefficient 1.0, 395 mana.
const ArcaneWrathRanks = 4

var ArcaneWrathSpellId = [ArcaneWrathRanks + 1]int32{0, 1326079, 1326080, 1326081, 1326082}
var ArcaneWrathBaseDamage = [ArcaneWrathRanks + 1][]float64{{0, 0}, {108, 132}, {168, 206}, {270, 330}, {406, 496}}
var ArcaneWrathDamagePerLevel = [ArcaneWrathRanks + 1]float64{0, 1.4, 2.3, 2.4, 3.0}
var ArcaneWrathManaCost = [ArcaneWrathRanks + 1]float64{0, 140, 220, 305, 395}
var ArcaneWrathLevel = [ArcaneWrathRanks + 1]int32{0, 25, 36, 48, 60}

func (mage *Mage) registerArcaneWrathSpell() {
	if !mage.Talents.ArcaneWrath {
		return
	}
	rank := 0
	for r := 1; r <= ArcaneWrathRanks; r++ {
		if mage.Level >= ArcaneWrathLevel[r] {
			rank = r
		}
	}
	if rank == 0 {
		return
	}
	cap := mage.Level
	if rank < ArcaneWrathRanks {
		cap = min(cap, ArcaneWrathLevel[rank+1]-1)
	}
	scale := ArcaneWrathDamagePerLevel[rank] * float64(max(0, cap-ArcaneWrathLevel[rank]))
	low, high := ArcaneWrathBaseDamage[rank][0]+scale, ArcaneWrathBaseDamage[rank][1]+scale

	mage.ArcaneWrath = mage.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_MageArcaneWrath,
		ActionID:    core.ActionID{SpellID: ArcaneWrathSpellId[rank]},
		SpellSchool: core.SpellSchoolArcane,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagMage | core.SpellFlagAPL,

		RequiredLevel: int(ArcaneWrathLevel[rank]),
		Rank:          rank,

		ManaCost: core.ManaCostOptions{FlatCost: ArcaneWrathManaCost[rank]},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault, CastTime: 3500 * time.Millisecond},
			CD:          core.Cooldown{Timer: mage.NewTimer(), Duration: 6 * time.Second},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := sim.Roll(low, high) + 0.05*mage.MaxMana()*spell.BonusCoefficient
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicHitAndCrit)
		},
	})
}
