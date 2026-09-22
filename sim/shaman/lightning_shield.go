package shaman

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const LightningShieldRanks = 7

// Forever beta client values for the orb (the proc spell): lower damage from rank 2 up, and every rank carries the
// full 0.267. Rank 7's 370 mana was missing from the table, which left it free.
var LightningShieldSpellId = [LightningShieldRanks + 1]int32{0, 324, 325, 905, 945, 8134, 10431, 10432}
var LightningShieldProcSpellId = [LightningShieldRanks + 1]int32{0, 26364, 26365, 26366, 26367, 26369, 26370, 26363}
var LightningShieldBaseDamage = [LightningShieldRanks + 1]float64{0, 13, 24, 40, 64, 96, 134, 178}
var LightningShieldSpellCoef = [LightningShieldRanks + 1]float64{0, .267, .267, .267, .267, .267, .267, .267}
var LightningShieldManaCost = [LightningShieldRanks + 1]float64{0, 45, 80, 125, 180, 240, 305, 370}
var LightningShieldLevel = [LightningShieldRanks + 1]int{0, 8, 16, 24, 32, 40, 48, 56}

func (shaman *Shaman) registerLightningShieldSpell() {
	shaman.LightningShield = make([]*core.Spell, LightningShieldRanks+1)
	shaman.LightningShieldProcs = make([]*core.Spell, LightningShieldRanks+1)
	shaman.LightningShieldAuras = make([]*core.Aura, LightningShieldRanks+1)

	for rank := 1; rank <= LightningShieldRanks; rank++ {
		level := LightningShieldLevel[rank]

		if level <= int(shaman.Level) {
			shaman.registerNewLightningShieldSpell(rank)
		}
	}
}

func (shaman *Shaman) registerNewLightningShieldSpell(rank int) {
	impLightningShieldBonus := 1 + []float64{0, .05, .10, .15}[shaman.Talents.ImprovedLightningShield]

	spellId := LightningShieldSpellId[rank]
	procSpellId := LightningShieldProcSpellId[rank]
	baseDamage := LightningShieldBaseDamage[rank] * impLightningShieldBonus
	spellCoeff := LightningShieldSpellCoef[rank]
	manaCost := LightningShieldManaCost[rank]
	level := LightningShieldLevel[rank]

	baseCharges := int32(3)
	maxCharges := int32(3)

	shaman.LightningShieldProcs[rank] = shaman.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: procSpellId},
		SpellSchool: core.SpellSchoolNature,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | SpellFlagShaman | SpellFlagLightning,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeAlwaysHit)
			shaman.ActiveShieldAura.RemoveStack(sim)
		},
	})

	// Both clients give the shield a 3.5 sec proc recovery.
	icd := core.Cooldown{
		Timer:    shaman.NewTimer(),
		Duration: time.Millisecond * 3500,
	}

	shaman.LightningShieldAuras[rank] = shaman.RegisterAura(core.Aura{
		Label:     fmt.Sprintf("Lightning Shield (Rank %d)", rank),
		ActionID:  core.ActionID{SpellID: spellId},
		Duration:  time.Minute * 10,
		MaxStacks: maxCharges,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.SetStacks(sim, baseCharges)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			if shaman.ActiveShieldAura.ActionID == aura.ActionID {
				shaman.ActiveShieldAura = nil
				shaman.ActiveShield = nil
			}
		},
		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskMelee) && result.Landed() && icd.IsReady(sim) {
				icd.Use(sim)
				shaman.LightningShieldProcs[rank].Cast(sim, spell.Unit)
			}
		},
	})

	shaman.LightningShield[rank] = shaman.RegisterSpell(core.SpellConfig{
		ActionID:  core.ActionID{SpellID: spellId},
		SpellCode: SpellCode_ShamanLightningShield,
		ProcMask:  core.ProcMaskEmpty,
		Flags:     core.SpellFlagAPL | SpellFlagShaman | SpellFlagLightning,

		RequiredLevel: level,
		Rank:          rank,

		ManaCost: core.ManaCostOptions{
			FlatCost:   manaCost,
			Multiplier: 100 - shaman.shamanisticFocusReduction(),
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			if shaman.ActiveShieldAura != nil {
				shaman.ActiveShieldAura.Deactivate(sim)
			}
			shaman.ActiveShield = spell
			shaman.ActiveShieldAura = shaman.LightningShieldAuras[rank]
			shaman.ActiveShieldAura.Activate(sim)
		},
	})
}
