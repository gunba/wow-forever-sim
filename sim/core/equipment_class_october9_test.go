//go:build with_db

package core

import (
	"slices"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestOctober9SpiritcallerClassRestrictions(t *testing.T) {
	for _, id := range []int32{276538, 276539, 276540, 276541} {
		item := ItemsByID[id]
		if !slices.Equal(item.ClassAllowlist, []proto.Class{proto.Class_ClassShaman}) {
			t.Fatalf("%d: source class restriction missing: %v", id, item.ClassAllowlist)
		}
		p := &proto.Player{
			Name: "Spiritcaller import", Class: proto.Class_ClassHunter, Race: proto.Race_RaceOrc,
			Spec:      &proto.Player_Hunter{Hunter: &proto.Hunter{}},
			Equipment: &proto.EquipmentSpec{Items: []*proto.ItemSpec{{Id: id}}},
		}
		var failure any
		func() {
			defer func() { failure = recover() }()
			NewCharacter(&Party{Index: 0}, 0, p)
		}()
		if failure == nil {
			t.Fatalf("Hunter imported Shaman-only %s", item.Name)
		}
		if err, ok := failure.(error); !ok || !strings.Contains(err.Error(), "class restriction") {
			t.Fatalf("unexpected import failure: %v", failure)
		}
		p.Class = proto.Class_ClassShaman
		p.Spec = &proto.Player_ElementalShaman{ElementalShaman: &proto.ElementalShaman{}}
		c := NewCharacter(&Party{Index: 0}, 0, p)
		if c.Equipment[ItemTypeToSlot(item.Type)].ID != id {
			t.Fatalf("Shaman lost allowed %s", item.Name)
		}
		// No class allowlist still means unrestricted, not forbidden.
		item.ClassAllowlist = nil
		var equipment Equipment
		equipment.EquipItem(item)
		if err := ValidateEquipmentArmor(proto.Class_ClassHunter, equipment); err != nil {
			t.Fatalf("unrestricted mail unexpectedly rejected: %v", err)
		}
	}
}
