package core

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

func TestInnervateManaAttribution(t *testing.T) {
	for _, tc := range []struct {
		name     string
		casting  bool
		missing  float64
		baseline float64
	}{
		{"casting", true, 5000, 110},
		{"idle", false, 5000, 210},
		{"partly capped", true, 200, 110},
		{"fully capped", true, 0, 110},
	} {
		t.Run(tc.name, func(t *testing.T) {
			unit := &Unit{PseudoStats: stats.NewPseudoStats()}
			unit.stats[stats.Mana] = 10000
			unit.stats[stats.MP5] = 25
			unit.manaBar.unit = unit
			unit.currentMana = unit.MaxMana() - tc.missing
			unit.SpiritManaRegenPerSecond = func() float64 { return 100 }
			unit.PseudoStats.SpiritRegenRateCasting = .5
			unit.PseudoStats.SpiritRegenMultiplier = 5
			unit.PseudoStats.FullSpiritRegenSources = 1
			if tc.casting {
				unit.PseudoStats.FiveSecondRuleRefreshTime = 5 * time.Second
			}
			unit.manaCastingMetrics = unit.NewManaMetrics(ActionID{OtherID: proto.OtherAction_OtherActionManaRegen, Tag: 1})
			unit.manaNotCastingMetrics = unit.NewManaMetrics(ActionID{OtherID: proto.OtherAction_OtherActionManaRegen, Tag: 2})
			innervate := unit.NewManaMetrics(ActionID{SpellID: 29166})
			unit.innervateRegenMetrics = innervate
			unit.UpdateManaRegenRates()
			sim := &Simulation{CurrentTime: 2 * time.Second}
			before := unit.CurrentMana()
			unit.ManaTick(sim)

			ordinary := unit.manaNotCastingMetrics
			if tc.casting {
				ordinary = unit.manaCastingMetrics
			}
			total := 1010.0 // two seconds of 5 MP/s plus 5 * 100 Spirit MP/s
			actual := min(total, tc.missing)
			checks := map[string][2]float64{
				"ordinary gain":  {ordinary.Gain, tc.baseline},
				"Innervate gain": {innervate.Gain, total - tc.baseline},
				"actual gain":    {unit.CurrentMana() - before, actual},
				"gain metrics":   {ordinary.Gain + innervate.Gain, total},
				"actual metrics": {ordinary.ActualGain + innervate.ActualGain, actual},
				"mana gained":    {unit.Metrics.ManaGained, actual},
				"bonus actual":   {innervate.ActualGain, max(0, actual-tc.baseline)},
			}
			for name, pair := range checks {
				if math.Abs(pair[0]-pair[1]) > 1e-8 {
					t.Errorf("%s: got %.9f, want %.9f", name, pair[0], pair[1])
				}
			}
			// Reattribution must not create mana-gain threat. No environment is
			// needed unless doneIteration incorrectly attempts to award it.
			threat := &Spell{
				ActionID: ActionID{OtherID: proto.OtherAction_OtherActionManaGain},
				Unit:     unit, SpellMetrics: make([]SpellMetrics, 1),
			}
			unit.Spellbook = []*Spell{threat}
			unit.manaBar.doneIteration(sim)
			if threat.SpellMetrics[0].TotalThreat != 0 {
				t.Fatal("passive Innervate regeneration generated threat")
			}
			unit.manaBar.reset()
			if unit.innervateRegenMetrics != nil {
				t.Fatal("Innervate attribution leaked across iteration reset")
			}
		})
	}
}
