package mage

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const IceBarrierRanks = 4

var IceBarrierSpellId = [IceBarrierRanks + 1]int32{0, 11426, 13031, 13032, 13033}
var IceBarrierManaCost = [IceBarrierRanks + 1]float64{0, 305, 360, 420, 480}
var IceBarrierLevel = [IceBarrierRanks + 1]int{0, 40, 46, 52, 58}

// Frost damage each of Ice Barrier's two buffs adds, from the server hotfix (SpellEffect 1364979-1365015, game cache
// push 112753): "increasing your Frost damage by $1326891s1% for $d. While shielded, your Frost damage is increased by
// an additional $s2%". Both are SPELLMOD percent rows of 3 on direct and periodic damage. The boss in a DPS sim never
// strikes the mage, so the shield holds for its whole minute.
const IceBarrierFrostDamage = 0.03

func (mage *Mage) iceBarrierFrostSpells() []*core.Spell {
	var spells []*core.Spell
	for _, spell := range mage.Spellbook {
		if spell.Flags.Matches(SpellFlagMage) && spell.SpellSchool.Matches(core.SpellSchoolFrost) && spell.ProcMask.Matches(core.ProcMaskSpellDamage) {
			spells = append(spells, spell)
		}
	}
	return spells
}

func (mage *Mage) frostDamageAura(label string, spellID int32, duration time.Duration) *core.Aura {
	var spells []*core.Spell
	return mage.RegisterAura(core.Aura{
		ActionID: core.ActionID{SpellID: spellID},
		Label:    label,
		Duration: duration,
		OnInit: func(aura *core.Aura, sim *core.Simulation) {
			spells = mage.iceBarrierFrostSpells()
		},
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range spells {
				spell.DamageMultiplierAdditive += IceBarrierFrostDamage
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			for _, spell := range spells {
				spell.DamageMultiplierAdditive -= IceBarrierFrostDamage
			}
		},
	})
}

func (mage *Mage) registerIceBarrierSpell() {
	mage.IceBarrier = make([]*core.Spell, IceBarrierRanks+1)
	mage.IceBarrierAuras = make([]*core.Aura, IceBarrierRanks+1)
	if !mage.Talents.IceBarrier {
		return
	}
	if mage.Env.IsForever() {
		mage.frigidPowerAura = mage.frostDamageAura("Frigid Power", 1326891, time.Minute)
	}

	cdTimer := mage.NewTimer()

	for rank := 1; rank <= IceBarrierRanks; rank++ {
		config := mage.newIceBarrierSpellConfig(rank, cdTimer)

		if config.RequiredLevel <= int(mage.Level) {
			mage.IceBarrier[rank] = mage.GetOrRegisterSpell(config)
		}
	}
}

func (mage *Mage) newIceBarrierSpellConfig(rank int, cdTimer *core.Timer) core.SpellConfig {
	spellID := IceBarrierSpellId[rank]
	manaCost := IceBarrierManaCost[rank]
	level := IceBarrierLevel[rank]

	cooldown := time.Second * 30

	if mage.Env.IsForever() {
		mage.IceBarrierAuras[rank] = mage.frostDamageAura(fmt.Sprintf("Ice Barrier (Rank %d)", rank), spellID, time.Minute)
	} else {
		mage.IceBarrierAuras[rank] = mage.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: spellID},
			Label:    fmt.Sprintf("Ice Barrier (Rank %d)", rank),
			Duration: time.Minute,
		})
	}

	return core.SpellConfig{
		ActionID:      core.ActionID{SpellID: spellID},
		SpellSchool:   core.SpellSchoolFrost,
		Flags:         core.SpellFlagAPL | core.SpellFlagNoOnCastComplete,
		Rank:          rank,
		RequiredLevel: level,
		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: cooldown,
			},
		},
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			// Disable an existing barrier
			if mage.activeBarrier != nil && mage.activeBarrier.IsActive() {
				mage.activeBarrier.Deactivate(sim)
				mage.activeBarrier = mage.IceBarrierAuras[rank]
			}
			mage.IceBarrierAuras[rank].Activate(sim)
			if mage.frigidPowerAura != nil {
				mage.frigidPowerAura.Activate(sim)
			}
		},
	}
}
