package tank

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common" // imported to get item effects included.
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterFeralTankDruid()
}

func TestP1FeralTank(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassDruid,
			Phase:      1,
			Race:       proto.Race_RaceTauren,
			OtherRaces: []proto.Race{proto.Race_RaceNightElf},

			Talents:     P1Talents,
			GearSet:     core.GetGearSet("../../../ui/feral_tank_druid/gear_sets", "launch"),
			Rotation:    core.GetAplRotation("../../../ui/feral_tank_druid/apls", "forever_bear"),
			Buffs:       core.ForeverBuffs,
			Consumes:    P1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Default", SpecOptions: PlayerOptionsDefault},

			ItemFilter: ItemFilters,
			// Without this the boss never attacks, so nothing the spec does in response to
			// being hit can fire and the damage taken metrics are all zero.
			IsTank:          true,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

var P1Talents = "01-5002232123132210551-055"

var PlayerOptionsDefault = &proto.Player_FeralTankDruid{
	FeralTankDruid: &proto.FeralTankDruid{
		Options: &proto.FeralTankDruid_Options{
			InnervateTarget: &proto.UnitReference{}, // no Innervate
			StartingRage:    20,
		},
	},
}

var P1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:     proto.AgilityElixir_ElixirOfTheMongoose,
		ArmorElixir:       proto.ArmorElixir_ElixirOfSuperiorDefense,
		AttackPowerBuff:   proto.AttackPowerBuff_JujuMight,
		DefaultPotion:     proto.Potions_GreaterStoneshieldPotion,
		DragonBreathChili: true,
		Flask:             proto.Flask_FlaskOfTheTitans,
		Food:              proto.Food_FoodSmokedDesertDumpling,
		HealthElixir:      proto.HealthElixir_ElixirOfFortitude,
		StrengthBuff:      proto.StrengthBuff_JujuPower,
	},
}

var ItemFilters = core.ItemFilter{
	ArmorType: proto.ArmorType_ArmorTypeLeather,

	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypeOffHand,
		proto.WeaponType_WeaponTypeStaff,
		proto.WeaponType_WeaponTypePolearm,
	},
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeIdol,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatStamina,
	proto.Stat_StatStrength,
	proto.Stat_StatAgility,
	proto.Stat_StatAttackPower,
	proto.Stat_StatArmor,
	proto.Stat_StatDodge,
	proto.Stat_StatDefense,
}
