package priest

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

var TalentTreeSizes = [3]int{18, 17, 18}

const (
	SpellFlagPriest = core.SpellFlagAgentReserved1
)

const (
	SpellCode_PriestNone int32 = iota

	SpellCode_PriestDevouringPlague
	SpellCode_PriestFlashHeal
	SpellCode_PriestGreaterHeal
	SpellCode_PriestHeal
	SpellCode_PriestHolyFire
	SpellCode_PriestHolyNova
	SpellCode_PriestMindBlast
	SpellCode_PriestMindFlay
	SpellCode_PriestPenance
	SpellCode_PriestShadowWordPain
	SpellCode_PriestSmite
	SpellCode_PriestStarshards
	SpellCode_PriestVampiricTouch
	SpellCode_PriestShadowWordDeath
)

type Priest struct {
	core.Character
	Talents *proto.PriestTalents

	Latency float64

	CircleOfHealing *core.Spell
	DevouringPlague []*core.Spell
	EmpoweredRenew  *core.Spell
	FlashHeal       []*core.Spell
	GreaterHeal     []*core.Spell
	HolyFire        []*core.Spell
	HolyNova        *core.Spell
	InnerFocus      *core.Spell
	MindBlast       []*core.Spell
	MindFlay        [][]*core.Spell // 1 entry for each tick for each rank
	Penance         *core.Spell
	PowerWordShield []*core.Spell
	PrayerOfHealing []*core.Spell
	PrayerOfMending *core.Spell
	Renew           []*core.Spell
	Shadowform      *core.Spell
	ShadowWordPain  []*core.Spell
	ShadowWordDeath []*core.Spell
	Smite           []*core.Spell
	Starshards      [][]*core.Spell
	VampiricEmbrace *core.Spell

	InnerFocusAura    *core.Aura
	SearingLightAura  *core.Aura
	ShadowformAura    *core.Aura
	ShadowWeavingAura *core.Aura
	SpiritTapAura     *core.Aura

	VampiricEmbraceAuras core.AuraArray
	WeakenedSouls        core.AuraArray

	shadowWeavingProcChance float64

	ProcPrayerOfMending core.ApplySpellResults
}

func (priest *Priest) GetCharacter() *core.Character {
	return &priest.Character
}

func (priest *Priest) RegisterSpell(config core.SpellConfig) *core.Spell {
	// Shadowform allows Holy damage and non-healing utility, but not active
	// healing casts. Holy Nova also heals even though its parent is damage.
	activeHeal := config.ProcMask.Matches(core.ProcMaskSpellHealing) && !config.Flags.Matches(core.SpellFlagPassiveSpell)
	if activeHeal || config.SpellCode == SpellCode_PriestHolyNova {
		condition := config.ExtraCastCondition
		config.ExtraCastCondition = func(sim *core.Simulation, target *core.Unit) bool {
			return (priest.ShadowformAura == nil || !priest.ShadowformAura.IsActive()) &&
				(condition == nil || condition(sim, target))
		}
	}
	return priest.Character.RegisterSpell(config)
}

func (priest *Priest) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	// Divine Spirit and Improved Power Word: Fortitude are gone from the Forever trees. Both are
	// raid buffs the rest of the raid is built around, so they are assumed to have become baseline.
	// TODO: beta will confirm whether they were made baseline or removed outright.
	raidBuffs.ShadowProtection = true
	raidBuffs.DivineSpirit = true
	raidBuffs.PowerWordFortitude = proto.TristateEffect_TristateEffectImproved
}

func (priest *Priest) AddPartyBuffs(_ *proto.PartyBuffs) {
}

func (priest *Priest) Initialize() {
	priest.registerMindBlast()
	priest.registerMindFlay()
	priest.registerShadowWordPainSpell()
	if priest.Env.IsForever() {
		priest.registerShadowWordDeath()
	}
	// Devouring Plague is an Undead racial in Classic. The Forever beta client teaches it to
	// priests of every race (SkillLineAbility race mask -1), so it is baseline here.
	if priest.Env.IsForever() || priest.GetCharacter().Race == proto.Race_RaceUndead {
		priest.registerDevouringPlagueSpell()
	}
	if priest.GetCharacter().Race == proto.Race_RaceNightElf {
		priest.registerStarshardsSpell()
	}
	priest.registerSmiteSpell()
	priest.registerHolyFire()
	priest.registerHolyNovaSpell()
	priest.registerPenanceSpell()
	priest.registerDarkSacrifice()

	priest.registerPowerInfusionCD()
}

func (priest *Priest) RegisterHealingSpells() {
	// priest.registerFlashHealSpell()
	// priest.registerGreaterHealSpell()
	// priest.registerPowerWordShieldSpell()
	// priest.registerPrayerOfHealingSpell()
	// priest.registerRenewSpell()
}

func New(character *core.Character, talents string) *Priest {
	priest := &Priest{
		Character: *character,
		Talents:   &proto.PriestTalents{},
	}
	core.FillTalentsProto(priest.Talents.ProtoReflect(), talents, TalentTreeSizes)

	priest.EnableManaBar()

	priest.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	priest.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[priest.Class]*core.SpellCritRatingPerCritChance)

	// Set mana regen to 12.5 + Spirit/4 each 2s tick
	priest.SpiritManaRegenPerSecond = func() float64 {
		return 6.25 + priest.GetStat(stats.Spirit)/8
	}

	return priest
}

// Agent is a generic way to access underlying priest on any of the agents.
type PriestAgent interface {
	GetPriest() *Priest
}
