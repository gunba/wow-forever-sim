//go:build with_db

package main

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestFeralImportRejectsMailAndPlate(t *testing.T) {
	for _, armor := range []proto.ArmorType{proto.ArmorType_ArmorTypeMail, proto.ArmorType_ArmorTypePlate} {
		req := racialFixture("feral", proto.Race_RaceTauren)
		req.SimOptions.Iterations = 1
		var head int32
		for id, item := range core.ItemsByID {
			if id >= 920000000 && item.Type == proto.ItemType_ItemTypeHead && item.ArmorType == armor && (head == 0 || id < head) {
				head = id
			}
		}
		if head == 0 {
			t.Fatal("missing modeled armor fixture")
		}
		req.Raid.Parties[0].Players[0].Equipment.Items[proto.ItemSlot_ItemSlotHead] = &proto.ItemSpec{Id: head}
		result := core.RunRaidSim(req)
		if result.Error == nil {
			t.Fatalf("native import accepted %s on Feral", armor)
		}
	}
}
