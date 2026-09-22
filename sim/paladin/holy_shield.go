package paladin

import (
	"strconv"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

var HolyShieldValues = []struct {
	level    int32
	manaCost float64
	damage   float64
}{
	// Beta client 1.60.1.69893: 110 / 153 / 221 (Classic 65 / 95 / 130), trained at 40 / 50 / 60 as in
	// Classic. The proc ids are Classic's learn-spell dummies, kept only to give the damage its own
	// metrics line; in the client the block damage comes from the aura itself.
	{level: 40, manaCost: 150, damage: 110},
	{level: 50, manaCost: 195, damage: 153},
	{level: 60, manaCost: 240, damage: 221},
}

func (paladin *Paladin) registerHolyShield() {
	if !paladin.Talents.HolyShield {
		return
	}

	// 4 charges and 20% block (Classic 30%), from the beta client.
	numCharges := int32(4)
	blockBonus := 20.0 * core.BlockRatingPerBlockChance

	timer := paladin.NewTimer() // All ranks share client cooldown category 931.
	for i, values := range HolyShieldValues {
		rank := i + 1
		level := values.level
		spellID := []int32{20925, 20927, 20928}[i]
		procID := []int32{20955, 20956, 20957}[i]
		manaCost := values.manaCost
		damage := values.damage

		if paladin.Level < level {
			break
		}

		paladin.holyShieldProc[i] = paladin.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: procID},
			SpellCode:   SpellCode_PaladinHolyShieldProc,
			SpellSchool: core.SpellSchoolHoly,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellDamage,

			RequiredLevel: int(level),
			Rank:          rank,

			DamageMultiplier: 1,
			ThreatMultiplier: 1.2,
			BonusCoefficient: 0.08, // Classic 0.05

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				// Spell damage from Holy Shield can crit, but does not miss.
				spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicCrit)
			},
		})

		paladin.holyShieldAura[i] = paladin.RegisterAura(core.Aura{
			Label:     "Holy Shield" + paladin.Label + strconv.Itoa(rank),
			ActionID:  core.ActionID{SpellID: spellID},
			Duration:  time.Second * 10,
			MaxStacks: numCharges,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.SetStacks(sim, numCharges)
				paladin.AddStatDynamic(sim, stats.Block, blockBonus)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				paladin.AddStatDynamic(sim, stats.Block, -blockBonus)
			},
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.DidBlock() {
					paladin.holyShieldProc[i].Cast(sim, spell.Unit)
					aura.RemoveStack(sim)
				}
			},
		})

		paladin.RegisterSpell(core.SpellConfig{
			ActionID:      core.ActionID{SpellID: spellID},
			SpellCode:     SpellCode_PaladinHolyShield,
			Flags:         core.SpellFlagAPL,
			RequiredLevel: int(level),
			Rank:          rank,
			ManaCost: core.ManaCostOptions{
				FlatCost:   manaCost,
				Multiplier: paladin.benediction(),
			},
			Cast: core.CastConfig{
				DefaultCast: core.Cast{
					GCD: core.GCDDefault,
				},
				CD: core.Cooldown{
					Timer:    timer,
					Duration: time.Second * 10,
				},
			},
			ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
				return paladin.OffHand().WeaponType == proto.WeaponType_WeaponTypeShield
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				paladin.holyShieldAura[i].Activate(sim)
			},
		})
	}
}
