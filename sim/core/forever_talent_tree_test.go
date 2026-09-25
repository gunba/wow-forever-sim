package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestForeverTalentTreePositionsAfterPatch(t *testing.T) {
	paladin := &proto.PaladinTalents{}
	FillTalentsProto(paladin.ProtoReflect(), "5-0-000000000000000001", [3]int{17, 16, 18})
	if paladin.DivineStrength != 5 || !paladin.TwistOfLight {
		t.Fatalf("removed Holy Strike talent shifted Paladin fields: %+v", paladin)
	}

	shaman := &proto.ShamanTalents{}
	FillTalentsProto(shaman.ProtoReflect(), "000000030000005", [3]int{16, 18, 16})
	if shaman.ElementalAlacrity != 3 || shaman.ElementalFury != 5 {
		t.Fatalf("Elemental talent swap mapped to old fields: %+v", shaman)
	}

	warrior := &proto.WarriorTalents{}
	FillTalentsProto(warrior.ProtoReflect(), "--000000000000000351", [3]int{17, 18, 18})
	if warrior.FocusedRage != 3 || warrior.Bastion != 5 || !warrior.ShieldSlam {
		t.Fatalf("Protection talent swap mapped to old fields: %+v", warrior)
	}
}
