package druid

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
)

const InsectSwarmRanks = 5

var InsectSwarmSpellId = [InsectSwarmRanks + 1]int32{0, 5570, 24974, 24975, 24976, 24977}

// Beta client 1.60.1.69893: total over 12 sec, the client's tick times 6 (rank 5 54 -> 31 a tick). The .158 per tick
// coefficient and the costs are Classic's.
var InsectSwarmBaseDamage = [InsectSwarmRanks + 1]float64{0, 48, 90, 120, 150, 186}
var InsectSwarmManaCost = [InsectSwarmRanks + 1]float64{0, 45, 85, 100, 140, 160}
var InsectSwarmLevel = [InsectSwarmRanks + 1]int{0, 20, 30, 40, 50, 60}

func (druid *Druid) registerInsectSwarmSpell() {
	if !druid.Talents.InsectSwarm {
		return
	}

	druid.InsectSwarm = make([]*DruidSpell, InsectSwarmRanks+1)

	druid.InsectSwarmAuras = druid.NewEnemyAuraArray(core.InsectSwarmAura)

	for rank := 1; rank <= InsectSwarmRanks; rank++ {
		level := InsectSwarmLevel[rank]
		if int32(level) <= druid.Level {
			numTicks := int32(6)
			tickLength := time.Second * 2

			spellID := InsectSwarmSpellId[rank]
			baseDamage := InsectSwarmBaseDamage[rank] / float64(numTicks)
			manaCost := InsectSwarmManaCost[rank]
			spellCoef := .158
			durationRemainder := time.Duration(0)
			if druid.HasSetBonus(ItemSetGrovekeeperEclipse, 5) {
				// 1301242 adds three seconds, not three ticks. Retain the
				// engine's full-tick rule and the full fifteen-second aura.
				numTicks++
				durationRemainder = time.Second
			}

			druid.InsectSwarm[rank] = druid.RegisterSpell(Humanoid|Moonkin, core.SpellConfig{
				SpellCode:   SpellCode_DruidInsectSwarm,
				ActionID:    core.ActionID{SpellID: spellID},
				SpellSchool: core.SpellSchoolNature,
				DefenseType: core.DefenseTypeMagic,
				ProcMask:    core.ProcMaskSpellDamage,
				Flags:       core.SpellFlagAPL | core.SpellFlagBinary,

				ManaCost: core.ManaCostOptions{
					FlatCost: manaCost,
				},
				Cast: core.CastConfig{
					DefaultCast: core.Cast{
						GCD: core.GCDDefault,
					},
				},

				DamageMultiplier: 1,
				ThreatMultiplier: 1,

				Dot: core.DotConfig{
					Aura: core.Aura{
						Label: fmt.Sprintf("Insect Swarm (Rank %d)", rank),
						OnGain: func(aura *core.Aura, sim *core.Simulation) {
							druid.InsectSwarmAuras.Get(aura.Unit).Activate(sim)
						},
						OnExpire: func(aura *core.Aura, sim *core.Simulation) {
							insectSwarmAura := druid.InsectSwarmAuras.Get(aura.Unit)
							if !insectSwarmAura.IsPermanent() {
								insectSwarmAura.Deactivate(sim)
							}
						},
					},

					NumberOfTicks:     numTicks,
					TickLength:        tickLength,
					DurationRemainder: durationRemainder,
					BonusCoefficient:  spellCoef,

					OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
						dot.Snapshot(target, baseDamage, isRollover)
					},
					OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
						dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
					},
				},

				ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
					result := spell.CalcOutcome(sim, target, spell.OutcomeMagicHitNoHitCounter)
					if result.Landed() {
						spell.Dot(target).Apply(sim)
					}
					spell.DealOutcome(sim, result)
				},

				RelatedAuras: []core.AuraArray{druid.InsectSwarmAuras},
			})
		}
	}
}
