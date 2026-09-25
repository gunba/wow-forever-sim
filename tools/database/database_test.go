package database

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestDuplicateEffectEnchantsHaveStableOrder(t *testing.T) {
	for _, reversed := range []bool{false, true} {
		db := NewWowDatabase()
		a := &proto.UIEnchant{EffectId: 66, SpellId: 7863, Name: "Boots"}
		b := &proto.UIEnchant{EffectId: 66, SpellId: 7457, Name: "Bracers"}
		if reversed {
			a, b = b, a
		}
		db.Enchants[EnchantToDBKey(a)] = a
		db.Enchants[EnchantToDBKey(b)] = b
		got := db.ToUIProto().Enchants
		if len(got) != 2 || got[0].SpellId != 7457 || got[1].SpellId != 7863 {
			t.Fatalf("insertion order %v: duplicate enchant order = %v", reversed, got)
		}
	}
}
