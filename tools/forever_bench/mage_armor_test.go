//go:build with_db

package main

import (
	"math"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func TestMageArmorDefaultAndRegeneration(t *testing.T) {
	castingRegen := map[proto.Mage_Options_ArmorType]float64{}
	for _, armor := range []proto.Mage_Options_ArmorType{proto.Mage_Options_NoArmor, proto.Mage_Options_MageArmor} {
		req := racialFixture("fire", proto.Race_RaceOrc)
		player := req.Raid.Parties[0].Players[0]
		if player.GetMage().Options.Armor != proto.Mage_Options_MageArmor {
			t.Fatal("Mage default does not select a supported armor")
		}
		player.GetMage().Options.Armor = armor
		player.Rotation = &proto.APLRotation{}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		castingRegen[armor] = sim.Raid.AllPlayerUnits[0].PseudoStats.SpiritRegenRateCasting
	}
	if math.Abs(castingRegen[proto.Mage_Options_MageArmor]-castingRegen[proto.Mage_Options_NoArmor]-.5) > 1e-8 {
		t.Fatal("Mage Armor did not add 50 percentage points of casting Spirit regeneration")
	}
}
