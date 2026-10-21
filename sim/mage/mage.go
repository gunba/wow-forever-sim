package mage

import (
	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"time"
)

const (
	SpellFlagMage       = core.SpellFlagAgentReserved1
	SpellFlagChillSpell = core.SpellFlagAgentReserved2
)

const (
	SpellCode_MageNone int32 = iota
	SpellCode_MageArcaneBlast
	SpellCode_MageArcaneExplosion
	SpellCode_MageArcaneMissiles
	SpellCode_MageArcaneMissilesTick
	SpellCode_MageBlastWave
	SpellCode_MageFireball
	SpellCode_MageFireBlast
	SpellCode_MageFlamestrike
	SpellCode_MageFrostbolt
	SpellCode_MageIceLance
	SpellCode_MageIgnite
	SpellCode_MagePyroblast
	SpellCode_MageScorch
	SpellCode_MageFrostfireBolt
	SpellCode_MageConeOfCold
	SpellCode_MageFrostNova
	SpellCode_MageArcaneWrath
)

var TalentTreeSizes = [3]int{19, 17, 19}

// ForeverTalentTreeFields is the tree after the server hotfix of 20 October 2026 (TraitNode 105771-105776, TraitNodeEntry
// 130507, game cache push 112753): Piercing Ice is gone, Frostbite and Permafrost trade places, Frost Channeling and
// Improved Blizzard move left, and Elemental Precision has three ranks.
var ForeverTalentTreeFields = [3][]string{
	{"wandSpecialization", "arcaneFocus", "improvedChanneling", "missileBarrage", "arcaneSubtlety", "magicAbsorption", "arcaneConcentration", "arcaneResilience", "arcaneGeometry", "arcaneImpact", "arcaneBlast", "arcaneShielding", "improvedCounterspell", "arcaneMeditation", "arcaneWrath", "presenceOfMind", "arcaneMind", "arcaneInstability", "arcanePower"},
	{"improvedFireBlast", "incineration", "improvedFireball", "ignite", "flameThrowing", "impact", "burningSoul", "improvedFlamestrike", "pyroblast", "improvedScorch", "improvedFireWard", "hotStreak", "masterOfElements", "criticalMass", "blastWave", "firePower", "combustion"},
	{"frostWarding", "improvedFrostbolt", "elementalPrecision", "iceShards", "frostbite", "improvedFrostNova", "permafrost", "frostChanneling", "improvedBlizzard", "iceLance", "arcticReach", "iceBlock", "shatter", "improvedConeOfCold", "coldSnap", "fingersOfFrost", "wintersChill", "iceBarrier"},
}

func RegisterMage() {
	core.RegisterAgentFactory(
		proto.Player_Mage{},
		proto.Spec_SpecMage,
		func(character *core.Character, options *proto.Player) core.Agent {
			return NewMage(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Mage)
			if !ok {
				panic("Invalid spec value for Mage!")
			}
			player.Spec = playerSpec
		},
	)
}

type Mage struct {
	core.Character

	Talents *proto.MageTalents
	Options *proto.Mage_Options

	activeBarrier *core.Aura

	ArcaneBlast             *core.Spell
	ArcaneWrath             *core.Spell
	pomUsedBy               *core.Spell
	pomUsedAt               time.Duration
	ArcaneExplosion         []*core.Spell
	ArcaneMissiles          []*core.Spell
	ArcaneMissilesTickSpell []*core.Spell
	BlastWave               []*core.Spell
	Blizzard                []*core.Spell
	Counterspell            *core.Spell
	ConeOfCold              []*core.Spell
	Evocation               *core.Spell
	Fireball                []*core.Spell
	FireBlast               []*core.Spell
	Flamestrike             []*core.Spell
	Frostbolt               []*core.Spell
	FrostfireBolt           []*core.Spell
	FrostNova               []*core.Spell
	IceBarrier              []*core.Spell
	IceLance                *core.Spell
	Ignite                  *core.Spell
	igniteTick              *core.Spell
	ManaGem                 []*core.Spell
	PresenceOfMind          *core.Spell
	Pyroblast               []*core.Spell
	Scorch                  []*core.Spell

	ArcaneBlastAura    *core.Aura
	ArcanePowerAura    *core.Aura
	ClearcastingAura   *core.Aura
	CombustionAura     *core.Aura
	FingersOfFrostAura *core.Aura
	castResultHooks    []func(sim *core.Simulation, spell *core.Spell)
	HotStreakAura      *core.Aura
	IceArmorAura       *core.Aura
	IceBarrierAuras    []*core.Aura
	ImprovedScorchAura *core.Aura
	ImpactAuras        core.AuraArray
	MageArmorAura      *core.Aura
	MissileBarrageAura *core.Aura
	WintersChillAura   *core.Aura
	frigidPowerAura    *core.Aura
	talentsString      string
	empoweredLances    int
}

// Agent is a generic way to access underlying mage on any of the agents.
type MageAgent interface {
	GetMage() *Mage
}

func (mage *Mage) GetCharacter() *core.Character {
	return &mage.Character
}

func (mage *Mage) GetMage() *Mage {
	return mage
}

func (mage *Mage) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	// Arcane Brilliance has never been a talent, so the Forever trees don't change what the mage
	// hands the rest of the raid. Improved Scorch and Winter's Chill are personal now, see talents.go.
	raidBuffs.ArcaneBrilliance = true
}
func (mage *Mage) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (mage *Mage) Initialize() {
	mage.registerArcaneBlastSpell()
	mage.registerArcaneWrathSpell()
	mage.registerArcaneMissilesSpell()
	mage.registerFireballSpell()
	mage.registerFireBlastSpell()
	mage.registerFrostboltSpell()
	mage.registerFrostfireBoltSpell()
	mage.registerIceLanceSpell()
	mage.registerPyroblastSpell()
	mage.registerScorchSpell()

	mage.registerArcaneExplosionSpell()
	mage.registerBlastWaveSpell()
	mage.registerBlizzardSpell()
	mage.registerFlamestrikeSpell()
	mage.registerFrostAreaSpells()

	mage.registerEvocationCD()
	mage.registerManaGemCD()
	mage.registerCounterspellSpell()
}

func (mage *Mage) Reset(sim *core.Simulation) {
}

// ResolveTalentTree fills the talents once the ruleset is known (core.Character.initialize).
func (mage *Mage) ResolveTalentTree() {
	mage.Talents = &proto.MageTalents{}
	if mage.Env != nil && mage.Env.IsForever() {
		core.FillTalentsProtoByName(mage.Talents.ProtoReflect(), mage.talentsString, ForeverTalentTreeFields)
	} else {
		core.FillTalentsProto(mage.Talents.ProtoReflect(), mage.talentsString, TalentTreeSizes)
	}
}

func NewMage(character *core.Character, options *proto.Player) *Mage {
	mageOptions := options.GetMage()

	mage := &Mage{
		Character: *character,
		Talents:   &proto.MageTalents{},
		Options:   mageOptions.Options,
	}
	// The tree layout depends on the ruleset, attached after construction: see ResolveTalentTree.
	mage.talentsString = options.TalentsString

	mage.EnableManaBar()

	mage.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	mage.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[mage.Class]*core.SpellCritRatingPerCritChance)

	switch mage.Options.Armor {
	case proto.Mage_Options_NoArmor:
	case proto.Mage_Options_IceArmor:
		mage.applyFrostIceArmor()
	case proto.Mage_Options_MageArmor:
		mage.applyMageArmor()
	default:
		panic("unsupported Mage armor selection; choose No Armor, Ice Armor, or Mage Armor")
	}

	// Set mana regen to 12.5 + Spirit/4 each 2s tick
	mage.SpiritManaRegenPerSecond = func() float64 {
		return 6.25 + mage.GetStat(stats.Spirit)/8
	}

	guardians.ConstructGuardians(&mage.Character)

	return mage
}
