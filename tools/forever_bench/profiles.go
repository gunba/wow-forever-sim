package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

type build struct {
	Key, Name, Dir, Preset, APL, Gear string
	Class                             proto.Class
	MinimumTrees                      [3]int
	Required                          string
}

func builds() []build {
	return []build{
		{"balance", "Balance", "balance_druid", "TalentsMoonkin", "launch", "forever_balance", proto.Class_ClassDruid, [3]int{31, 0, 0}, "moonkinForm"},
		{"feral", "Feral", "feral_druid", "TalentsFeralCat", "feral", "forever_feral", proto.Class_ClassDruid, [3]int{0, 31, 0}, "berserk"},
		{"elemental", "Elemental", "elemental_shaman", "TalentsElemental", "default", "forever_elemental", proto.Class_ClassShaman, [3]int{31, 0, 0}, "lavaBurst"},
		{"stormcaller", "Stormcaller", "elemental_shaman", "TalentsStormcaller", "stormcaller", "forever_stormcaller", proto.Class_ClassShaman, [3]int{21, 20, 0}, ""},
		{"enhancement", "Enhancement", "enhancement_shaman", "TalentsEnhancement", "default", "forever_enhancement", proto.Class_ClassShaman, [3]int{0, 31, 0}, "rageOfTheFarseer"},
		{"beast_mastery", "Beast Mastery", "hunter", "TalentsBeastMastery", "beast_mastery", "forever_beast_mastery", proto.Class_ClassHunter, [3]int{31, 0, 0}, "bestialWrath"},
		{"marksmanship", "Marksmanship", "hunter", "TalentsP1", "p1", "forever_marksmanship", proto.Class_ClassHunter, [3]int{0, 31, 0}, ""},
		{"survival", "Survival", "hunter", "TalentsSurvival", "survival", "forever_survival", proto.Class_ClassHunter, [3]int{0, 0, 31}, "laceratingStrikes"},
		{"pet_melee", "Pet/Melee", "hunter", "TalentsPetMelee", "pet_melee", "forever_survival", proto.Class_ClassHunter, [3]int{15, 0, 20}, "summonHawk"},
		{"arcane", "Arcane", "mage", "TalentsArcane", "forever_arcane", "forever_arcane", proto.Class_ClassMage, [3]int{31, 0, 0}, "arcanePower"},
		{"fire", "Fire", "mage", "TalentsP1Fire", "forever_fire", "forever_fire", proto.Class_ClassMage, [3]int{0, 31, 0}, "combustion"},
		{"frost", "Frost", "mage", "TalentsFrost", "forever_frost", "forever_frost", proto.Class_ClassMage, [3]int{0, 0, 30}, "wintersChill"},
		{"arcane_frost", "Arcane–Frost", "mage", "TalentsArcaneFrost", "forever_arcane_frost", "forever_frost", proto.Class_ClassMage, [3]int{21, 0, 20}, "missileBarrage"},
		{"retribution", "Retribution", "retribution_paladin", "TalentsRetribution", "basic_ret", "forever_retribution", proto.Class_ClassPaladin, [3]int{0, 0, 31}, "twistOfLight"},
		{"retribution_physical", "Physical Ret", "retribution_paladin", "TalentsPhysicalRetribution", "basic_ret", "forever_retribution", proto.Class_ClassPaladin, [3]int{0, 0, 31}, "twistOfLight"},
		{"shadow", "Shadow", "shadow_priest", "TalentsP1Shadow", "p1", "forever_shadow", proto.Class_ClassPriest, [3]int{0, 0, 31}, "shadowform"},
		{"smite", "Smite", "smite_priest", "TalentsSmite", "launch", "forever_smite", proto.Class_ClassPriest, [3]int{0, 15, 0}, "searingLight"},
		{"combat", "Combat", "rogue", "CombatSinisterStrikeTalents", "combat_sinister_strike_sweaty", "forever_combat", proto.Class_ClassRogue, [3]int{0, 31, 0}, "adrenalineRush"},
		{"mutilate", "Mutilate", "rogue", "AssassinationMutilateTalents", "forever_mutilate", "forever_mutilate", proto.Class_ClassRogue, [3]int{31, 0, 0}, ""},
		{"subtlety", "Subtlety", "rogue", "TalentsSubtletyHemo", "forever_hemorrhage", "forever_subtlety", proto.Class_ClassRogue, [3]int{0, 0, 31}, "thousandCuts"},
		{"demonology", "Demonic Pact", "warlock", "TalentsPactOptimised", "forever_pact", "forever_demonology", proto.Class_ClassWarlock, [3]int{0, 31, 0}, "demonicPact"},
		{"affliction", "Affliction", "warlock", "TalentsDeepAffliction", "forever_affliction", "forever_affliction", proto.Class_ClassWarlock, [3]int{31, 0, 0}, ""},
		{"ds_ruin", "DS/Ruin", "warlock", "TalentsDSRuinPandemic", "forever_ds_ruin", "forever_ds_ruin", proto.Class_ClassWarlock, [3]int{0, 11, 15}, "demonicSacrifice"},
		{"destruction", "Destruction", "warlock", "TalentsShadowAndFlame", "forever_shadow_and_flame", "forever_destruction", proto.Class_ClassWarlock, [3]int{0, 0, 31}, "incinerate"},
		{"fury", "Fury", "warrior", "TalentsP1DPS", "forever_fury", "forever_fury", proto.Class_ClassWarrior, [3]int{0, 31, 0}, "bloodthirst"},
		{"fury_sunder", "Fury (Sunder)", "warrior", "TalentsP1DPS", "forever_fury_sunder", "forever_fury", proto.Class_ClassWarrior, [3]int{0, 31, 0}, "bloodthirst"},
		{"arms", "Arms", "warrior", "TalentsArms", "forever_arms", "forever_arms", proto.Class_ClassWarrior, [3]int{31, 0, 0}, "mortalStrike"},
		{"fury_2h", "2H Bloodthirst", "warrior", "TalentsFuryTwoHand", "forever_fury_2h", "forever_arms", proto.Class_ClassWarrior, [3]int{11, 31, 0}, "bloodthirst"},
	}
}

// Hybrid rows inherit their comparison profile's equipment and scenario model.
// Talent identity and rotation remain specific to the hybrid.
func (b build) modelKey() string {
	switch b.Key {
	case "pet_melee":
		return "survival"
	case "arcane_frost":
		return "frost"
	case "fury_2h":
		return "arms"
	case "fury_sunder":
		return "fury"
	case "retribution_physical":
		return "retribution"
	default:
		return b.Key
	}
}

func (b build) presetTalents() string {
	data := mustRead(filepath.Join("ui", b.Dir, "presets.ts"))
	pattern := regexp.MustCompile(`(?s)export const ` + regexp.QuoteMeta(b.Preset) + `\s*=.*?talentsString:\s*'([0-9-]+)'`)
	m := pattern.FindSubmatch(data)
	if m == nil {
		panic(fmt.Sprintf("missing %s/%s", b.Dir, b.Preset))
	}
	return string(m[1])
}

func (b build) player(race proto.Race) *proto.Player {
	p := &proto.Player{
		Name: b.Name, Class: b.Class, Race: race,
		Equipment:     core.GetGearSet(filepath.Join("ui", b.Dir, "gear_sets"), b.Gear).GearSet,
		Rotation:      core.GetAplRotation(filepath.Join("ui", b.Dir, "apls"), b.APL).Rotation,
		TalentsString: b.presetTalents(), Buffs: core.ForeverIndividualBuffs,
		Consumes: b.consumes(), DistanceFromTarget: 20,
		Profession1:         proto.Profession_Engineering,
		ForeverTier1Bonuses: true,
	}
	switch b.modelKey() {
	case "balance":
		core.WithSpec(p, &proto.Player_BalanceDruid{BalanceDruid: &proto.BalanceDruid{Options: &proto.BalanceDruid_Options{}}})
	case "feral":
		p.DistanceFromTarget = 5
		core.WithSpec(p, &proto.Player_FeralDruid{FeralDruid: &proto.FeralDruid{Options: &proto.FeralDruid_Options{LatencyMs: 100}}})
	case "elemental", "stormcaller":
		core.WithSpec(p, &proto.Player_ElementalShaman{ElementalShaman: &proto.ElementalShaman{Options: &proto.ElementalShaman_Options{}}})
	case "enhancement":
		p.DistanceFromTarget = 5
		core.WithSpec(p, &proto.Player_EnhancementShaman{EnhancementShaman: &proto.EnhancementShaman{Options: &proto.EnhancementShaman_Options{}}})
	case "beast_mastery", "marksmanship", "survival":
		p.DistanceFromTarget = 12
		if b.modelKey() == "survival" {
			p.DistanceFromTarget = 5
		}
		core.WithSpec(p, &proto.Player_Hunter{Hunter: &proto.Hunter{Options: &proto.Hunter_Options{
			Ammo: proto.Hunter_Options_ThoriumHeadedArrow, QuiverBonus: proto.Hunter_Options_Speed15,
			PetType: proto.Hunter_Options_Cat, PetUptime: 1, PetAttackSpeed: proto.Hunter_Options_OneTwo,
		}}})
		if b.modelKey() == "marksmanship" {
			p.GetHunter().Options.PetType = proto.Hunter_Options_PetNone
		}
	case "arcane", "fire", "frost":
		core.WithSpec(p, &proto.Player_Mage{Mage: &proto.Mage{Options: &proto.Mage_Options{Armor: proto.Mage_Options_MageArmor}}})
	case "retribution":
		p.DistanceFromTarget = 5
		core.WithSpec(p, &proto.Player_RetributionPaladin{RetributionPaladin: &proto.RetributionPaladin{Options: &proto.PaladinOptions{
			Aura: proto.PaladinAura_NoPaladinAura, PrimarySeal: proto.PaladinSeal_Command,
		}}})
	case "shadow":
		core.WithSpec(p, &proto.Player_ShadowPriest{ShadowPriest: &proto.ShadowPriest{Options: &proto.ShadowPriest_Options{}}})
	case "smite":
		core.WithSpec(p, &proto.Player_SmitePriest{SmitePriest: &proto.SmitePriest{Options: &proto.SmitePriest_Options{}}})
	case "combat", "mutilate", "subtlety":
		p.DistanceFromTarget = 5
		core.WithSpec(p, &proto.Player_Rogue{Rogue: &proto.Rogue{Options: &proto.RogueOptions{}}})
	case "demonology", "affliction", "ds_ruin", "destruction":
		summon, sacrifice := proto.WarlockOptions_Succubus, proto.WarlockOptions_NoSummon
		if b.Key == "demonology" {
			sacrifice = proto.WarlockOptions_Voidwalker
		} else if b.Key == "ds_ruin" {
			summon = proto.WarlockOptions_Imp
		}
		core.WithSpec(p, &proto.Player_Warlock{Warlock: &proto.Warlock{Options: &proto.WarlockOptions{
			Armor: proto.WarlockOptions_DemonArmor, Summon: summon, Sacrifice: sacrifice,
		}}})
	case "fury", "arms":
		p.DistanceFromTarget = 5
		core.WithSpec(p, &proto.Player_Warrior{Warrior: &proto.Warrior{Options: &proto.Warrior_Options{
			StartingRage: 50, QueueDelay: 250, Shout: proto.WarriorShout_WarriorShoutBattle, Stance: proto.WarriorStance_WarriorStanceBerserker,
		}}})
	}
	return p
}

func casterConsumes() *proto.Consumes {
	return &proto.Consumes{
		Flask: proto.Flask_FlaskOfSupremePower, Food: proto.Food_FoodRunnTumTuberSurprise,
		DefaultPotion: proto.Potions_MajorManaPotion, DefaultConjured: proto.Conjured_ConjuredDemonicRune,
		ManaRegenElixir: proto.ManaRegenElixir_MagebloodPotion, SpellPowerBuff: proto.SpellPowerBuff_GreaterArcaneElixir,
		MainHandImbue: proto.WeaponImbue_BrilliantWizardOil, ZanzaBuff: proto.ZanzaBuff_CerebralCortexCompound,
	}
}

func (b build) consumes() *proto.Consumes {
	physicalRet := b.Key == "retribution_physical"
	b.Key = b.modelKey()
	c := casterConsumes()
	switch b.Class {
	case proto.Class_ClassMage, proto.Class_ClassShaman:
		c.FirePowerBuff, c.FrostPowerBuff = proto.FirePowerBuff_ElixirOfFirepower, proto.FrostPowerBuff_ElixirOfFrostPower
	case proto.Class_ClassWarlock:
		c.FirePowerBuff, c.ShadowPowerBuff = proto.FirePowerBuff_ElixirOfFirepower, proto.ShadowPowerBuff_ElixirOfShadowPower
	case proto.Class_ClassPriest:
		if b.Key == "shadow" {
			c.ShadowPowerBuff = proto.ShadowPowerBuff_ElixirOfShadowPower
		}
	}
	if b.Key == "elemental" || b.Key == "stormcaller" {
		c.StrengthBuff = proto.StrengthBuff_JujuPower
	}
	switch b.Key {
	case "enhancement", "retribution", "beast_mastery", "marksmanship", "survival", "feral", "combat", "mutilate", "subtlety", "fury", "arms":
		c.AgilityElixir, c.AttackPowerBuff, c.StrengthBuff = proto.AgilityElixir_ElixirOfTheMongoose, proto.AttackPowerBuff_JujuMight, proto.StrengthBuff_JujuPower
		c.ZanzaBuff, c.Food = proto.ZanzaBuff_GroundScorpokAssay, proto.Food_FoodGrilledSquid
		c.MainHandImbue, c.OffHandImbue = proto.WeaponImbue_WeaponImbueUnknown, proto.WeaponImbue_WeaponImbueUnknown
	}
	switch b.Key {
	case "enhancement", "retribution":
		c.ZanzaBuff, c.Food = proto.ZanzaBuff_ROIDS, proto.Food_FoodBlessSunfruit
		c.FirePowerBuff, c.FrostPowerBuff = proto.FirePowerBuff_ElixirOfFirepower, proto.FrostPowerBuff_FrostPowerBuffUnknown
		c.DragonBreathChili = true
		if b.Key == "enhancement" {
			c.MainHandImbue, c.OffHandImbue = proto.WeaponImbue_WindfuryWeapon, proto.WeaponImbue_WeaponImbueUnknown
		} else {
			c.MainHandImbue = proto.WeaponImbue_Windfury
		}
		if physicalRet {
			c.Flask = proto.Flask_FlaskUnknown
			c.SpellPowerBuff = proto.SpellPowerBuff_SpellPowerBuffUnknown
			c.FirePowerBuff = proto.FirePowerBuff_FirePowerBuffUnknown
		}
	case "beast_mastery", "marksmanship", "survival":
		c.Food = proto.Food_FoodSmokedDesertDumpling
		c.MainHandImbue, c.OffHandImbue = proto.WeaponImbue_Windfury, proto.WeaponImbue_ElementalSharpeningStone
	case "feral":
		c.Flask = proto.Flask_FlaskOfDistilledWisdom
		c.SpellPowerBuff, c.ManaRegenElixir = 0, 0
		c.DefaultConjured = proto.Conjured_ConjuredRogueThistleTea
		c.DragonBreathChili, c.SapperExplosive = true, proto.SapperExplosive_SapperGoblinSapper
	case "combat", "mutilate", "subtlety":
		c.DefaultPotion, c.DefaultConjured, c.ManaRegenElixir = 0, proto.Conjured_ConjuredRogueThistleTea, 0
		c.MainHandImbue, c.OffHandImbue = proto.WeaponImbue_InstantPoison, proto.WeaponImbue_DeadlyPoison
		c.DragonBreathChili, c.SapperExplosive = true, proto.SapperExplosive_SapperGoblinSapper
	case "fury", "arms":
		c.Flask, c.SpellPowerBuff, c.ManaRegenElixir = 0, 0, 0
		c.Food, c.ZanzaBuff = proto.Food_FoodSmokedDesertDumpling, proto.ZanzaBuff_ROIDS
		c.ArmorElixir, c.HealthElixir, c.Alcohol = proto.ArmorElixir_ElixirOfSuperiorDefense, proto.HealthElixir_ElixirOfFortitude, proto.Alcohol_AlcoholRumseyRumBlackLabel
		c.DefaultPotion = proto.Potions_MightyRagePotion
		c.MainHandImbue, c.OffHandImbue = proto.WeaponImbue_Windfury, proto.WeaponImbue_ElementalSharpeningStone
		c.DragonBreathChili, c.SapperExplosive = true, proto.SapperExplosive_SapperGoblinSapper
		if b.Key == "arms" {
			c.OffHandImbue = proto.WeaponImbue_WeaponImbueUnknown
		}
	}
	return c
}

func mustRead(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func readPlayer(path string) *proto.Player {
	p := &proto.Player{}
	if err := protojson.Unmarshal(mustRead(path), p); err != nil {
		panic(err)
	}
	return p
}

// The checked-in exact profiles are the native defaults as well as the source
// of the web's ranking picker. The preset gear files remain useful as a seed
// when introducing a new build before its first full benchmark.
var rankedDefaultPlayers = sync.OnceValue(func() map[string]*proto.Player {
	var source struct {
		Results []struct {
			Key, Race      string
			BaselinePlayer json.RawMessage
		}
	}
	if err := json.Unmarshal(mustRead("artifacts/modelled_gear/forever_input_profiles.json"), &source); err != nil {
		panic(err)
	}
	players := make(map[string]*proto.Player, len(source.Results))
	for _, row := range source.Results {
		player := &proto.Player{}
		if err := protojson.Unmarshal(row.BaselinePlayer, player); err != nil {
			panic(err)
		}
		key := row.Key + "/" + row.Race
		if _, exists := players[key]; exists {
			panic("duplicate ranked player: " + key)
		}
		players[key] = player
	}
	return players
})

func (b build) rankedPlayer(race proto.Race) *proto.Player {
	if p := rankedDefaultPlayers()[b.Key+"/"+raceName(race)]; p != nil {
		return googleProto.Clone(p).(*proto.Player)
	}
	return nil
}
