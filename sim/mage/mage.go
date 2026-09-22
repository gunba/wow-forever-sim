package mage

import (
	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
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
)

var TalentTreeSizes = [3]int{18, 17, 19}

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
	ArcaneExplosion         []*core.Spell
	ArcaneMissiles          []*core.Spell
	ArcaneMissilesTickSpell []*core.Spell
	BlastWave               []*core.Spell
	Blizzard                []*core.Spell
	Counterspell            *core.Spell
	Evocation               *core.Spell
	Fireball                []*core.Spell
	FireBlast               []*core.Spell
	Flamestrike             []*core.Spell
	Frostbolt               []*core.Spell
	FrostfireBolt           []*core.Spell
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
	HotStreakAura      *core.Aura
	IceArmorAura       *core.Aura
	IceBarrierAuras    []*core.Aura
	ImprovedScorchAura *core.Aura
	MageArmorAura      *core.Aura
	MissileBarrageAura *core.Aura
	WintersChillAura   *core.Aura
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

	mage.registerEvocationCD()
	mage.registerManaGemCD()
	mage.registerCounterspellSpell()
}

func (mage *Mage) Reset(sim *core.Simulation) {
}

func NewMage(character *core.Character, options *proto.Player) *Mage {
	mageOptions := options.GetMage()

	mage := &Mage{
		Character: *character,
		Talents:   &proto.MageTalents{},
		Options:   mageOptions.Options,
	}
	core.FillTalentsProto(mage.Talents.ProtoReflect(), options.TalentsString, TalentTreeSizes)

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
