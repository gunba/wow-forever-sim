package mage

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterMage()
}

func TestP1Mage(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassMage,
			Phase:      1,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceGnome},

			Talents:     P1FrostTalents,
			GearSet:     core.GetGearSet("../../ui/mage/gear_sets", "p0.bis"),
			Rotation:    core.GetAplRotation("../../ui/mage/apls", "forever_frost"),
			Buffs:       core.ForeverBuffs,
			Consumes:    P1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "DPS", SpecOptions: PlayerOptions},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

func TestP1MageArcane(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassMage,
			Phase: 1,
			Race:  proto.Race_RaceTroll,

			Talents:     P1ArcaneTalents,
			GearSet:     core.GetGearSet("../../ui/mage/gear_sets", "p0.bis"),
			Rotation:    core.GetAplRotation("../../ui/mage/apls", "forever_arcane"),
			Buffs:       core.ForeverBuffs,
			Consumes:    P1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "DPS", SpecOptions: PlayerOptions},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

func TestP1MageFire(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassMage,
			Phase: 1,
			Race:  proto.Race_RaceTroll,

			Talents:     P1FireTalents,
			GearSet:     core.GetGearSet("../../ui/mage/gear_sets", "p0.bis"),
			Rotation:    core.GetAplRotation("../../ui/mage/apls", "forever_fire"),
			Buffs:       core.ForeverBuffs,
			Consumes:    P1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "DPS", SpecOptions: PlayerOptions},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

// The community builds the rankings page runs, on the same raid preset as the P1 tests.
func TestForeverMageArcane(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassMage,
			Phase: 1,
			Race:  proto.Race_RaceGnome,

			Talents:     ForeverArcaneTalents,
			GearSet:     core.GetGearSet("../../ui/mage/gear_sets", "p0.bis"),
			Rotation:    core.GetAplRotation("../../ui/mage/apls", "forever_arcane"),
			Buffs:       core.ForeverBuffs,
			Consumes:    P1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "DPS", SpecOptions: PlayerOptions},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

func TestForeverMageFire(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassMage,
			Phase: 1,
			Race:  proto.Race_RaceGnome,

			Talents:     ForeverFireTalents,
			GearSet:     core.GetGearSet("../../ui/mage/gear_sets", "p0.bis"),
			Rotation:    core.GetAplRotation("../../ui/mage/apls", "forever_fire"),
			Buffs:       core.ForeverBuffs,
			Consumes:    P1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "DPS", SpecOptions: PlayerOptions},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

func TestForeverMageFrost(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassMage,
			Phase: 1,
			Race:  proto.Race_RaceGnome,

			Talents:     ForeverFrostTalents,
			GearSet:     core.GetGearSet("../../ui/mage/gear_sets", "p0.bis"),
			Rotation:    core.GetAplRotation("../../ui/mage/apls", "forever_frost"),
			Buffs:       core.ForeverBuffs,
			Consumes:    P1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "DPS", SpecOptions: PlayerOptions},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

var P1FrostTalents = "0502050030003--055500033100030024"
var P1ArcaneTalents = "050215003100311531-2305003202003-"
var P1FireTalents = "0502252000003-23550000130133051-"

// Arcane 35/0/16, Fire 0/35/16 and Frost 14/0/37 from ui/mage/presets.ts.
var ForeverArcaneTalents = "055005023100311531--005500033"
var ForeverFireTalents = "-03552020130133151-005500033"
var ForeverFrostTalents = "050005013--0555003301001301251"

var PlayerOptions = &proto.Player_Mage{
	Mage: &proto.Mage{
		Options: &proto.Mage_Options{
			Armor: proto.Mage_Options_MageArmor,
		},
	},
}

var P1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:  proto.Potions_MajorManaPotion,
		Flask:          proto.Flask_FlaskOfSupremePower,
		FirePowerBuff:  proto.FirePowerBuff_ElixirOfGreaterFirepower,
		FrostPowerBuff: proto.FrostPowerBuff_ElixirOfFrostPower,
		Food:           proto.Food_FoodRunnTumTuberSurprise,
		MainHandImbue:  proto.WeaponImbue_BrilliantWizardOil,
		SpellPowerBuff: proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeSword,
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
	proto.Stat_StatSpellPower,
	proto.Stat_StatArcanePower,
	proto.Stat_StatFirePower,
	proto.Stat_StatFrostPower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}
