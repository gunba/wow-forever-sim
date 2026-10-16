package smite

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common" // imported to get caster sets included.
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterSmitePriest()
}

func TestSmitePriest(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassPriest,
			Phase:      1,
			Race:       proto.Race_RaceUndead,
			OtherRaces: []proto.Race{proto.Race_RaceDwarf},

			Talents:     LaunchTalents,
			GearSet:     core.GetGearSet("../../../ui/smite_priest/gear_sets", "launch"),
			Rotation:    core.GetAplRotation("../../../ui/smite_priest/apls", "launch"),
			Buffs:       core.ForeverBuffs,
			Consumes:    LaunchConsumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Basic", SpecOptions: PlayerOptionsBasic},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

// The community Smite 31/17/3 the rankings page ranks, so the golden measures the build
// people actually run. Thirty-one Discipline for Power Infusion, seventeen Holy for the
// talents that scale Smite; Power in Light and Searing Light are what make the build.
var LaunchTalents = "515035031030001031-00505023002-003"

var LaunchConsumes = core.ConsumesCombo{
	Label: "Launch-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:  proto.Potions_MajorManaPotion,
		Flask:          proto.Flask_FlaskOfSupremePower,
		Food:           proto.Food_FoodRunnTumTuberSurprise,
		MainHandImbue:  proto.WeaponImbue_WizardOil,
		SpellPowerBuff: proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var PlayerOptionsBasic = &proto.Player_SmitePriest{
	SmitePriest: &proto.SmitePriest{
		Options: &proto.SmitePriest_Options{
			Armor: proto.SmitePriest_Options_InnerFire,
		},
	},
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypeOffHand,
		proto.WeaponType_WeaponTypeStaff,
	},
	ArmorType: proto.ArmorType_ArmorTypeCloth,
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeWand,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatIntellect,
	proto.Stat_StatSpirit,
	proto.Stat_StatSpellPower,
	proto.Stat_StatHolyPower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}
