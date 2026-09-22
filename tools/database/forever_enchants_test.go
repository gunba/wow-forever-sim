package database

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/wowsims/classic/sim/core/stats"
)

func TestForeverHealingEnchantEffects(t *testing.T) {
	data, err := os.ReadFile("../../assets/db_inputs/forever_effect_audit.json")
	if err != nil {
		t.Fatal(err)
	}
	var audit struct {
		Enchants []struct {
			EnchantID int32 `json:"enchantId"`
			Stats     struct {
				HealingPower float64
				SpellDamage  float64
			} `json:"healingAndDamageStats"`
		}
	}
	if err := json.Unmarshal(data, &audit); err != nil {
		t.Fatal(err)
	}
	if len(audit.Enchants) != 10 {
		t.Fatalf("expected ten audited enchants, got %d", len(audit.Enchants))
	}
	for _, record := range audit.Enchants {
		found := false
		for _, enchant := range EnchantOverrides {
			if enchant.EffectId != record.EnchantID {
				continue
			}
			found = true
			s := stats.FromFloatArray(enchant.Stats)
			if s[stats.SpellPower] != 0 || s[stats.HealingPower] != record.Stats.HealingPower || s[stats.SpellDamage] != record.Stats.SpellDamage {
				t.Errorf("enchant %d disagrees with client healing/damage effects: %v", enchant.EffectId, s)
			}
		}
		if !found {
			t.Errorf("audited enchant %d missing", record.EnchantID)
		}
	}
}
