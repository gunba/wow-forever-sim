package shaman

import (
	"time"

	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

var TalentTreeSizes = [3]int{16, 18, 16}

const (
	SpellFlagShaman    = core.SpellFlagAgentReserved1
	SpellFlagTotem     = core.SpellFlagAgentReserved2
	SpellFlagLightning = core.SpellFlagAgentReserved3
)

// mainHandImbue and offHandImbue are the Forever spec-option imbues. Forever imbues have their own enchant
// slot (client effect 360) next to the temporary weapon enchant (effect 54: stones, weightstones and oils),
// so they no longer share Consumes. A legacy profile that still names a Shaman imbue in Consumes keeps it.
func NewShaman(character *core.Character, talents string, mainHandImbue, offHandImbue proto.ShamanWeaponImbue) *Shaman {
	shaman := &Shaman{
		Character: *character,
		Talents:   &proto.ShamanTalents{},
	}
	shaman.optionImbues = [2]proto.ShamanWeaponImbue{mainHandImbue, offHandImbue}

	core.FillTalentsProto(shaman.Talents.ProtoReflect(), talents, TalentTreeSizes)
	shaman.EnableManaBar()

	// Add Shaman stat dependencies
	shaman.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	shaman.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritPerAgiAtLevel[character.Class]*core.CritRatingPerCritChance)
	shaman.AddStatDependency(stats.Agility, stats.Dodge, core.DodgePerAgiAtLevel[character.Class]*core.DodgeRatingPerDodgeChance)
	shaman.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[character.Class]*core.SpellCritRatingPerCritChance)
	shaman.AddStatDependency(stats.BonusArmor, stats.Armor, 1)
	shaman.PseudoStats.BlockValuePerStrength = .05 // 20 str = 1 block

	guardians.ConstructGuardians(&shaman.Character)

	return shaman
}

var shamanImbues = map[proto.ShamanWeaponImbue]proto.WeaponImbue{
	proto.ShamanWeaponImbue_RockbiterImbue:   proto.WeaponImbue_RockbiterWeapon,
	proto.ShamanWeaponImbue_FlametongueImbue: proto.WeaponImbue_FlametongueWeapon,
	proto.ShamanWeaponImbue_FrostbrandImbue:  proto.WeaponImbue_FrostbrandWeapon,
	proto.ShamanWeaponImbue_WindfuryImbue:    proto.WeaponImbue_WindfuryWeapon,
}

// classImbue returns the Shaman imbue on one hand: the spec option if set, otherwise a Shaman imbue
// left in the Consumes weapon-imbue field.
func classImbue(option proto.ShamanWeaponImbue, consumable proto.WeaponImbue) proto.WeaponImbue {
	if imbue, ok := shamanImbues[option]; ok {
		return imbue
	}
	for _, imbue := range shamanImbues {
		if consumable == imbue {
			return imbue
		}
	}
	return proto.WeaponImbue_WeaponImbueUnknown
}

func (shaman *Shaman) getImbueProcMask(imbue proto.WeaponImbue) core.ProcMask {
	mask := core.ProcMaskUnknown
	if shaman.HasMHWeapon() && shaman.imbues[0] == imbue {
		mask |= core.ProcMaskMeleeMH
	}
	if shaman.HasOHWeapon() && shaman.imbues[1] == imbue {
		mask |= core.ProcMaskMeleeOH
	}
	return mask
}

// resolveImbues picks each hand's imbue once the ruleset is known. Classic imbues are temporary weapon
// enchants that share the Consumes slot with stones and oils, so the spec options are ignored there.
func (shaman *Shaman) ResolveRulesetInputs() {
	options := shaman.optionImbues
	if !shaman.Env.IsForever() {
		options = [2]proto.ShamanWeaponImbue{}
	}
	shaman.imbues = [2]proto.WeaponImbue{
		classImbue(options[0], shaman.Consumes.GetMainHandImbue()),
		classImbue(options[1], shaman.Consumes.GetOffHandImbue()),
	}
	shaman.ApplyRockbiterImbue(shaman.getImbueProcMask(proto.WeaponImbue_RockbiterWeapon))
	shaman.ApplyFlametongueImbue(shaman.getImbueProcMask(proto.WeaponImbue_FlametongueWeapon))
	shaman.ApplyFrostbrandImbue(shaman.getImbueProcMask(proto.WeaponImbue_FrostbrandWeapon))
	shaman.ApplyWindfuryImbue(shaman.getImbueProcMask(proto.WeaponImbue_WindfuryWeapon))
}

// Indexes into NextTotemDrops for self buffs
const (
	AirTotem int = iota
	EarthTotem
	FireTotem
	WaterTotem
)

const (
	SpellCode_ShamanNone int32 = iota

	SpellCode_ShamanChainHeal
	SpellCode_ShamanChainLightning
	SpellCode_ShamanEarthShock
	SpellCode_ShamanFireNova
	SpellCode_ShamanFlameShock
	SpellCode_ShamanFrostShock
	SpellCode_ShamanHealingWave
	SpellCode_ShamanLavaBurst
	SpellCode_ShamanLesserHealingWave
	SpellCode_ShamanLightningBolt
	SpellCode_ShamanLightningShield
	SpellCode_ShamanMagmaTotem
	SpellCode_ShamanSearingTotem
	SpellCode_ShamanStormstrike
	SpellCode_ShamanWaterShield
)

// Shaman represents a shaman character.
type Shaman struct {
	core.Character

	Talents *proto.ShamanTalents

	// Spells
	ChainHeal              []*core.Spell
	ChainLightning         []*core.Spell
	ChainLightningOverload []*core.Spell
	EarthShield            *core.Spell
	EarthShock             []*core.Spell
	FireNova               []*core.Spell
	FlameShock             []*core.Spell
	FrostShock             []*core.Spell
	GraceOfAirTotem        []*core.Spell
	HealingStreamTotem     []*core.Spell
	HealingWave            []*core.Spell
	LavaBurst              []*core.Spell
	LesserHealingWave      []*core.Spell
	LightningBolt          []*core.Spell
	LightningBoltOverload  []*core.Spell
	LightningShield        []*core.Spell
	LightningShieldProcs   []*core.Spell // The damage component of lightning shield is a separate spell
	MagmaTotem             []*core.Spell
	ManaSpringTotem        []*core.Spell
	ManaTideTotem          []*core.Spell
	SearingTotem           []*core.Spell
	StoneskinTotem         []*core.Spell
	Stormstrike            *core.Spell
	StrengthOfEarthTotem   []*core.Spell
	TremorTotem            *core.Spell
	WaterShield            *core.Spell
	WindfuryTotem          []*core.Spell
	WindfuryWeaponMH       *core.Spell
	WindfuryWeaponOH       *core.Spell
	WindwallTotem          []*core.Spell

	// Auras
	ClearcastingAura     *core.Aura
	LightningShieldAuras []*core.Aura
	MaelstromWeaponAura  *core.Aura
	WaterShieldAura      *core.Aura

	// Totems
	ActiveTotems     [4]*core.Spell
	ActiveTotemBuffs [4]*core.Aura
	TotemExpirations [4]time.Duration // The expiration time of each totem (earth, air, fire, water).

	EarthTotems []*core.Spell
	FireTotems  []*core.Spell
	WaterTotems []*core.Spell
	AirTotems   []*core.Spell
	Totems      *proto.ShamanTotems

	WindfuryTotemPeriodicActions      []*core.PendingAction
	ActiveWindfuryTotemPeriodicAction *core.PendingAction

	// Shield
	ActiveShield     *core.Spell // Tracks the Shaman's active shield spell
	ActiveShieldAura *core.Aura

	ChainLightningBounceCoefficient float64

	// The weapon imbue on each hand, and the spec-option imbues it came from.
	imbues       [2]proto.WeaponImbue
	optionImbues [2]proto.ShamanWeaponImbue
}

// Implemented by each Shaman spec.
type ShamanAgent interface {
	core.Agent

	// The Shaman controlled by this Agent.
	GetShaman() *Shaman
}

func (shaman *Shaman) GetCharacter() *core.Character {
	return &shaman.Character
}

func (shaman *Shaman) AddRaidBuffs(_ *proto.RaidBuffs) {
	// Buffs are handled explicitly through APLs now
}

func (shaman *Shaman) Initialize() {

	// Core abilities
	shaman.registerChainLightningSpell()
	shaman.registerLavaBurstSpell()
	shaman.registerLightningBoltSpell()
	shaman.registerLightningShieldSpell()
	shaman.registerShocks()
	shaman.registerFireNovaSpell()
	shaman.registerStormstrikeSpell()
	shaman.registerWaterShieldSpell()

	// Imbues
	// In the Initialize due to frost brand adding the aura to the enemy
	shaman.RegisterRockbiterImbue(shaman.getImbueProcMask(proto.WeaponImbue_RockbiterWeapon))
	shaman.RegisterFlametongueImbue(shaman.getImbueProcMask(proto.WeaponImbue_FlametongueWeapon))
	shaman.RegisterWindfuryImbue(shaman.getImbueProcMask(proto.WeaponImbue_WindfuryWeapon))
	shaman.RegisterFrostbrandImbue(shaman.getImbueProcMask(proto.WeaponImbue_FrostbrandWeapon))

	// Totems
	shaman.registerStrengthOfEarthTotemSpell()
	shaman.registerStoneskinTotemSpell()
	shaman.registerTremorTotemSpell()
	shaman.registerSearingTotemSpell()
	shaman.registerMagmaTotemSpell()
	shaman.registerHealingStreamTotemSpell()
	shaman.registerManaSpringTotemSpell()
	shaman.registerManaTideTotemSpell()
	shaman.registerWindfuryTotemSpell()
	shaman.registerGraceOfAirTotemSpell()
	shaman.registerWindwallTotemSpell()
}

func (shaman *Shaman) Reset(_ *core.Simulation) {
	shaman.ActiveShield = nil
	shaman.ActiveShieldAura = nil

	for i := range []int{EarthTotem, FireTotem, WaterTotem, AirTotem} {
		shaman.ActiveTotems[i] = nil
		shaman.TotemExpirations[i] = 0
		shaman.ActiveTotemBuffs[i] = nil
	}
}
