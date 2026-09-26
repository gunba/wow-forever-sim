package core

import (
	"fmt"
	"testing"
)

func TestAuraRegistryScalesWithRaidEffects(t *testing.T) {
	unit := &Unit{}
	for i := 0; i < 2001; i++ {
		label := fmt.Sprintf("rank-and-attacker-%d", i)
		aura := unit.RegisterAura(Aura{Label: label})
		if unit.GetAura(label) != aura {
			t.Fatalf("registered aura %d was lost", i)
		}
	}
	// Removing the size guard does not permit duplicate registrations.
	defer func() {
		if recover() == nil {
			t.Fatal("duplicate aura registration did not fail")
		}
	}()
	unit.RegisterAura(Aura{Label: "rank-and-attacker-0"})
}
