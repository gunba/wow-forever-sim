package core

import (
	_ "embed"
	"encoding/json"

	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

// Shared by native benchmark setup and the web presets. This is a default
// scenario, not an override applied to arbitrary user requests.
//
//go:embed forever_tank_support.json
var foreverTankSupportJSON []byte

func ForeverTankSupport(class proto.Class) BuffsCombo {
	if class != proto.Class_ClassWarrior && class != proto.Class_ClassPaladin && class != proto.Class_ClassDruid {
		panic("Forever tank support requires a supported tank class")
	}
	var data map[string]json.RawMessage
	if err := json.Unmarshal(foreverTankSupportJSON, &data); err != nil {
		panic(err)
	}
	buffs := BuffsCombo{
		Raid: &proto.RaidBuffs{}, Party: &proto.PartyBuffs{},
		Player: &proto.IndividualBuffs{}, Debuffs: &proto.Debuffs{},
	}
	for key, message := range map[string]googleProto.Message{
		"raid": buffs.Raid, "party": buffs.Party,
		"player": buffs.Player, "debuffs": buffs.Debuffs,
	} {
		if err := protojson.Unmarshal(data[key], message); err != nil {
			panic(err)
		}
	}
	// Warrior supplies its own shout and major armor debuff. No tank gets
	// external Thunder Clap or Demoralizing Shout/Roar: retain those duties.
	if class == proto.Class_ClassWarrior {
		buffs.Raid.BattleShout = proto.TristateEffect_TristateEffectMissing
		buffs.Debuffs.ExposeArmor = proto.TristateEffect_TristateEffectMissing
	}
	return buffs
}
