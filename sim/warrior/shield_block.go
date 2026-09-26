package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

func (warrior *Warrior) RegisterShieldBlockCD() {
	actionID := core.ActionID{SpellID: 2565}
	cooldownDur := time.Second * 5
	duration := time.Second * 5
	maxBlocks := int32(1)
	if warrior.Env.IsForever() {
		// Forever's level-16 rank lasts seven seconds or two successful blocks.
		duration = time.Second * 7
		maxBlocks = 2
	}

	warrior.ShieldBlockAura = warrior.RegisterAura(core.Aura{
		Label:     "Shield Block",
		ActionID:  actionID,
		Duration:  duration,
		MaxStacks: maxBlocks,

		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			warrior.AddStatDynamic(sim, stats.Block, 75*core.BlockRatingPerBlockChance)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			warrior.AddStatDynamic(sim, stats.Block, -75*core.BlockRatingPerBlockChance)
		},
		OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.DidBlock() {
				aura.RemoveStack(sim)
			}
		},
	})

	warrior.ShieldBlock = warrior.RegisterSpell(DefensiveStance, core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,

		RageCost: core.RageCostOptions{
			Cost: 10,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{},
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: cooldownDur,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return warrior.PseudoStats.CanBlock
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			warrior.ShieldBlockAura.Activate(sim)
			// The five-second cooldown is shorter than the Forever aura.
			// Refreshing an active aura does not call OnGain again.
			warrior.ShieldBlockAura.SetStacks(sim, warrior.ShieldBlockAura.MaxStacks)
		},
	})

	warrior.AddMajorCooldown(core.MajorCooldown{
		Spell:    warrior.ShieldBlock.Spell,
		Priority: core.CooldownPriorityDefault,
		Type:     core.CooldownTypeSurvival,
		ShouldActivate: func(s *core.Simulation, c *core.Character) bool {
			// Only castable with manual APL Action
			return false
		},
	})
}
