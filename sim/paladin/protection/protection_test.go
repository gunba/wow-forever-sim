package protection

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterProtectionPaladin()
}

func TestProtection(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassPaladin,
			Phase:      1,
			Race:       proto.Race_RaceHuman,
			OtherRaces: []proto.Race{proto.Race_RaceDwarf},

			Talents:     ProtectionTalents,
			GearSet:     core.GetGearSet("../../../ui/protection_paladin/gear_sets", "launch"),
			Rotation:    core.GetAplRotation("../../../ui/protection_paladin/apls", "forever_protection"),
			Buffs:       core.ForeverBuffs,
			Consumes:    ProtectionConsumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Forever Protection", SpecOptions: PlayerOptionsSealofRighteousness},

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

var ProtectionTalents = "-5521513321301551-15002"

var ProtectionConsumes = core.ConsumesCombo{
	Label: "Forever Protection",
	Consumes: &proto.Consumes{
		DefaultPotion:     proto.Potions_MajorManaPotion,
		AgilityElixir:     proto.AgilityElixir_ElixirOfTheMongoose,
		AttackPowerBuff:   proto.AttackPowerBuff_JujuMight,
		Flask:             proto.Flask_FlaskOfTheTitans,
		SpellPowerBuff:    proto.SpellPowerBuff_GreaterArcaneElixir,
		DragonBreathChili: true,
		Food:              proto.Food_FoodSmokedDesertDumpling,
		StrengthBuff:      proto.StrengthBuff_JujuPower,
	},
}

var PlayerOptionsSealofCommand = &proto.Player_ProtectionPaladin{
	ProtectionPaladin: &proto.ProtectionPaladin{
		Options: optionsSealOfCommand,
	},
}

var PlayerOptionsSealofRighteousness = &proto.Player_ProtectionPaladin{
	ProtectionPaladin: &proto.ProtectionPaladin{
		Options: optionsSealOfRighteousness,
	},
}

var optionsSealOfCommand = &proto.PaladinOptions{
	PrimarySeal:   proto.PaladinSeal_Command,
	RighteousFury: true,
}

var optionsSealOfRighteousness = &proto.PaladinOptions{
	PrimarySeal:   proto.PaladinSeal_Righteousness,
	RighteousFury: true,
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeAxe,
		proto.WeaponType_WeaponTypeSword,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypePolearm,
		proto.WeaponType_WeaponTypeShield,
	},
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeLibram,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatHealth,
	proto.Stat_StatMana,
	proto.Stat_StatStrength,
	proto.Stat_StatStamina,
	proto.Stat_StatAgility,
	proto.Stat_StatIntellect,
	proto.Stat_StatAttackPower,
	proto.Stat_StatMeleeHit,
	proto.Stat_StatMeleeCrit,
	proto.Stat_StatMeleeHaste,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
	proto.Stat_StatSpellPower,
	proto.Stat_StatHolyPower,
	proto.Stat_StatHealingPower,
	proto.Stat_StatArmor,
	proto.Stat_StatBonusArmor,
	proto.Stat_StatDefense,
	proto.Stat_StatDodge,
	proto.Stat_StatParry,
	proto.Stat_StatBlock,
	proto.Stat_StatBlockValue,
	proto.Stat_StatFireResistance,
	proto.Stat_StatNatureResistance,
	proto.Stat_StatShadowResistance,
	proto.Stat_StatFrostResistance,
	proto.Stat_StatArcaneResistance,
}
