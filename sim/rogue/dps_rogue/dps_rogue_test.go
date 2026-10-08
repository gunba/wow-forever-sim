package dpsrogue

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterDpsRogue()
}

func TestCombatSinisterStrike(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassRogue,
			Race:       proto.Race_RaceHuman,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     CombatSwordsTalents,
			GearSet:     core.GetGearSet("../../../ui/rogue/gear_sets", "combat_sinister_strike_prebis"),
			Rotation:    core.GetAplRotation("../../../ui/rogue/apls", "combat_sinister_strike"),
			Buffs:       core.ForeverBuffs,
			Consumes:    Phase1Consumes,
			Phase:       1,
			SpecOptions: core.SpecOptionsCombo{Label: "Windfury MH", SpecOptions: DefaultRogue},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

func TestCombatDaggers(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassRogue,
			Race:       proto.Race_RaceHuman,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     CombatDaggersTalents,
			GearSet:     core.GetGearSet("../../../ui/rogue/gear_sets", "combat_backstab_prebis"),
			Rotation:    core.GetAplRotation("../../../ui/rogue/apls", "combat_backstab"),
			Buffs:       core.ForeverBuffs,
			Consumes:    Phase1Consumes,
			Phase:       1,
			SpecOptions: core.SpecOptionsCombo{Label: "Windfury MH", SpecOptions: DefaultRogue},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

func TestAssassinationMutilate(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassRogue,
			Race:       proto.Race_RaceHuman,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     AssassinationMutilateTalents,
			GearSet:     core.GetGearSet("../../../ui/rogue/gear_sets", "combat_backstab_prebis"),
			Rotation:    core.GetAplRotation("../../../ui/rogue/apls", "forever_mutilate"),
			Buffs:       core.ForeverBuffs,
			Consumes:    Phase1PoisonConsumes,
			Phase:       1,
			SpecOptions: core.SpecOptionsCombo{Label: "Poisons", SpecOptions: DualPoisonRogue},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

func TestSubtletyHemorrhage(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassRogue,
			Race:       proto.Race_RaceHuman,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     SubtletyHemorrhageTalents,
			GearSet:     core.GetGearSet("../../../ui/rogue/gear_sets", "combat_backstab_prebis"),
			Rotation:    core.GetAplRotation("../../../ui/rogue/apls", "forever_hemorrhage"),
			Buffs:       core.ForeverBuffs,
			Consumes:    Phase1PoisonConsumes,
			Phase:       1,
			SpecOptions: core.SpecOptionsCombo{Label: "Poisons", SpecOptions: DualPoisonRogue},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,

			Ruleset: proto.Ruleset_RulesetForever,
		},
	}))
}

var CombatSwordsTalents = "00530310501-32003311201515231"
var CombatDaggersTalents = "005302005-30230320201515231-102"
var AssassinationMutilateTalents = "00530310551021051-302303202004"
var SubtletyHemorrhageTalents = "125320101--5320003310013211551"

var DefaultRogue = &proto.Player_Rogue{
	Rogue: &proto.Rogue{
		Options: &proto.RogueOptions{OffHandPoison: proto.RogueOptions_InstantPoison},
	},
}

var DualPoisonRogue = &proto.Player_Rogue{
	Rogue: &proto.Rogue{
		Options: &proto.RogueOptions{MainHandPoison: proto.RogueOptions_InstantPoison, OffHandPoison: proto.RogueOptions_DeadlyPoison},
	},
}

var ItemFilters = core.ItemFilter{
	ArmorType: proto.ArmorType_ArmorTypeLeather,
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeFist,
		proto.WeaponType_WeaponTypeSword,
		proto.WeaponType_WeaponTypeMace,
	},
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeBow,
		proto.RangedWeaponType_RangedWeaponTypeCrossbow,
		proto.RangedWeaponType_RangedWeaponTypeGun,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatAttackPower,
	proto.Stat_StatAgility,
	proto.Stat_StatStrength,
	proto.Stat_StatMeleeHit,
	proto.Stat_StatMeleeCrit,
}

// Combat runs Windfury on the main hand and a poison off hand; the Assassination combo
// below is the one that poisons both.
var Phase1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:   proto.AgilityElixir_ElixirOfTheMongoose,
		MainHandImbue:   proto.WeaponImbue_Windfury,
		StrengthBuff:    proto.StrengthBuff_JujuPower,
		AttackPowerBuff: proto.AttackPowerBuff_JujuMight,
	},
}

// Venom and Mutilate both key off the poisons, so the Assassination build runs them on both weapons.
var Phase1PoisonConsumes = core.ConsumesCombo{
	Label: "P1-Poison-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:   proto.AgilityElixir_ElixirOfTheMongoose,
		StrengthBuff:    proto.StrengthBuff_JujuPower,
		AttackPowerBuff: proto.AttackPowerBuff_JujuMight,
	},
}
