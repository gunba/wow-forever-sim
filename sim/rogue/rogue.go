package rogue

import (
	"time"

	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	SpellFlagBuilder      = core.SpellFlagAgentReserved1
	SpellFlagColdBlooded  = core.SpellFlagAgentReserved2
	SpellFlagDeadlyBrewed = core.SpellFlagAgentReserved3
	SpellFlagCarnage      = core.SpellFlagAgentReserved4 // for Carnage
	SpellFlagRoguePoison  = core.SpellFlagAgentReserved5 // RogueT1
)

const (
	SpellCode_RogueNone int32 = iota

	SpellCode_RogueAmbush
	SpellCode_RogueAdrenalineRush
	SpellCode_RogueBackstab
	SpellCode_RogueBladeFlurry
	SpellCode_RogueEviscerate
	SpellCode_RogueExposeArmor
	SpellCode_RogueGarrote
	SpellCode_RogueGhostlyStrike
	SpellCode_RogueHemorrhage
	SpellCode_RogueMutilate
	SpellCode_RogueRupture
	SpellCode_RogueSinisterStrike
	SpellCode_RogueSliceandDice
	SpellCode_RogueVanish
	SpellCode_RogueVenom
)

var TalentTreeSizes = [3]int{17, 17, 19}

type Rogue struct {
	core.Character

	Talents *proto.RogueTalents
	Options *proto.RogueOptions

	sliceAndDiceDurations [6]time.Duration

	AdrenalineRush *core.Spell
	Backstab       *core.Spell
	BladeFlurry    *core.Spell
	Feint          *core.Spell
	Garrote        *core.Spell
	Ambush         *core.Spell
	Hemorrhage     *core.Spell
	GhostlyStrike  *core.Spell
	Mutilate       *core.Spell
	mutilateMH     *core.Spell
	mutilateOH     *core.Spell
	SinisterStrike *core.Spell
	Shadowstep     *core.Spell
	Preparation    *core.Spell
	Premeditation  *core.Spell
	ColdBlood      *core.Spell
	Vanish         *core.Spell

	Eviscerate   *core.Spell
	ExposeArmor  *core.Spell
	Rupture      *core.Spell
	SliceAndDice *core.Spell
	Venom        *core.Spell
	Finishers    []*core.Spell

	Evasion *core.Spell

	DeadlyPoison     *core.Spell
	deadlyPoisonTick *core.Spell
	InstantPoison    *core.Spell
	WoundPoison      *core.Spell

	additivePoisonBonusChance float64

	AdrenalineRushAura *core.Aura
	BladeFlurryAura    *core.Aura
	CutthroatAura      *core.Aura
	ExposeArmorAuras   core.AuraArray
	EvasionAura        *core.Aura
	HemorrhageAuras    core.AuraArray
	SliceAndDiceAura   *core.Aura
	StealthAura        *core.Aura
	ThousandCutsAura   *core.Aura
	VanishAura         *core.Aura
	VenomAura          *core.Aura

	woundPoisonDebuffAuras core.AuraArray
}

func (rogue *Rogue) GetCharacter() *core.Character {
	return &rogue.Character
}

func (rogue *Rogue) GetRogue() *Rogue {
	return rogue
}

func (rogue *Rogue) AddRaidBuffs(_ *proto.RaidBuffs)   {}
func (rogue *Rogue) AddPartyBuffs(_ *proto.PartyBuffs) {}

func (rogue *Rogue) finisherFlags() core.SpellFlag {
	return core.SpellFlagMeleeMetrics | core.SpellFlagAPL
}

func (rogue *Rogue) builderFlags() core.SpellFlag {
	return SpellFlagBuilder | SpellFlagColdBlooded | core.SpellFlagMeleeMetrics | core.SpellFlagAPL
}

func (rogue *Rogue) Initialize() {
	rogue.validatePoisonInputs()
	rogue.applyPoisons()
	rogue.registerBackstabSpell()
	rogue.registerEviscerate()
	rogue.registerExposeArmorSpell()
	rogue.registerFeintSpell()
	rogue.registerGarrote()
	rogue.registerHemorrhageSpell()
	rogue.registerMutilateSpell()
	rogue.registerRupture()
	rogue.registerSinisterStrikeSpell()
	rogue.registerSliceAndDice()
	rogue.registerAmbushSpell()
	rogue.registerVenom()

	// Poisons
	rogue.registerInstantPoisonSpell()
	rogue.registerDeadlyPoisonSpell()
	rogue.registerWoundPoisonSpell()

	// Stealth
	rogue.registerStealthAura()
	rogue.registerVanishSpell()
}

func (rogue *Rogue) ApplyEnergyTickMultiplier(sim *core.Simulation, multiplier float64) {
	if rogue.Env.IsForever() {
		// Settle the elapsed fraction at the old rate before Adrenaline Rush
		// starts or ends, rather than applying the new rate retroactively.
		rogue.ResetEnergyTick(sim)
	}
	rogue.EnergyTickMultiplier += multiplier
}

func (rogue *Rogue) Reset(_ *core.Simulation) {
	for _, mcd := range rogue.GetMajorCooldowns() {
		mcd.Disable()
	}
}

func NewRogue(character *core.Character, options *proto.Player, rogueOptions *proto.RogueOptions) *Rogue {
	rogue := &Rogue{
		Character: *character,
		Talents:   &proto.RogueTalents{},
		Options:   rogueOptions,
	}
	core.FillTalentsProto(rogue.Talents.ProtoReflect(), options.TalentsString, TalentTreeSizes)

	// Passive rogue threat reduction: https://wotlk.wowhead.com/spell=21184/rogue-passive-dnd
	rogue.PseudoStats.ThreatMultiplier *= 0.71
	// TODO: Be able to Parry based on results
	rogue.PseudoStats.CanParry = true
	maxEnergy := 100.0 + 5*float64(rogue.Talents.Vigor)
	rogue.EnableEnergyBar(maxEnergy)

	rogue.EnableAutoAttacks(rogue, core.AutoAttackOptions{
		MainHand:       rogue.WeaponFromMainHand(),
		OffHand:        rogue.WeaponFromOffHand(),
		Ranged:         rogue.WeaponFromRanged(),
		AutoSwingMelee: true,
	})
	rogue.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	rogue.AddStatDependency(stats.Agility, stats.AttackPower, 1)
	rogue.AddStatDependency(stats.Agility, stats.RangedAttackPower, 1)
	rogue.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritPerAgiAtLevel[character.Class]*core.CritRatingPerCritChance)
	rogue.AddStatDependency(stats.Agility, stats.Dodge, core.DodgePerAgiAtLevel[character.Class]*core.DodgeRatingPerDodgeChance)
	rogue.AddStatDependency(stats.BonusArmor, stats.Armor, 1)

	guardians.ConstructGuardians(&rogue.Character)

	return rogue
}

// Deactivate Stealth if it is active. This must be added to all abilities that cause Stealth to fade.
func (rogue *Rogue) BreakStealth(sim *core.Simulation) {
	if rogue.StealthAura.IsActive() {
		rogue.StealthAura.Deactivate(sim)
		rogue.AutoAttacks.EnableAutoSwing(sim)
	}
}

// Does the rogue have a dagger equipped in the specified hand (main or offhand)?
func (rogue *Rogue) HasDagger(hand core.Hand) bool {
	if hand == core.MainHand {
		return rogue.MainHand().WeaponType == proto.WeaponType_WeaponTypeDagger
	}
	return rogue.OffHand().WeaponType == proto.WeaponType_WeaponTypeDagger
}

// Check if the rogue is considered in "stealth" for the purpose of casting abilities
func (rogue *Rogue) IsStealthed() bool {
	return rogue.StealthAura.IsActive()
}

// Agent is a generic way to access underlying rogue on any of the agents.
type RogueAgent interface {
	GetRogue() *Rogue
}

func (rogue *Rogue) validatePoisonInputs() {
	if !rogue.Env.IsForever() {
		return
	}
	for _, imbue := range []proto.WeaponImbue{rogue.Consumes.MainHandImbue, rogue.Consumes.OffHandImbue} {
		switch imbue {
		case proto.WeaponImbue_InstantPoison, proto.WeaponImbue_DeadlyPoison, proto.WeaponImbue_WoundPoison:
			panic("Forever Rogue poisons belong in RogueOptions main_hand_poison/off_hand_poison, not Consumes weapon imbues; migrate this profile")
		}
	}
	for hand, poison := range []proto.RogueOptions_Poison{rogue.Options.GetMainHandPoison(), rogue.Options.GetOffHandPoison()} {
		if _, valid := proto.RogueOptions_Poison_name[int32(poison)]; !valid {
			panic("Invalid Rogue poison selection")
		}
		if poison != proto.RogueOptions_NoPoison && ((hand == 0 && !rogue.HasMHWeapon()) || (hand == 1 && !rogue.HasOHWeapon())) {
			panic("Rogue poison requires a weapon in the selected hand")
		}
	}
}

func (rogue *Rogue) getPoisonProcMask(poison proto.RogueOptions_Poison) core.ProcMask {
	mh, oh := rogue.Options.GetMainHandPoison() == poison, rogue.Options.GetOffHandPoison() == poison
	if !rogue.Env.IsForever() {
		// Classic poisons still occupy the temporary weapon-enchant slot.
		imbue := map[proto.RogueOptions_Poison]proto.WeaponImbue{
			proto.RogueOptions_InstantPoison: proto.WeaponImbue_InstantPoison,
			proto.RogueOptions_DeadlyPoison:  proto.WeaponImbue_DeadlyPoison,
			proto.RogueOptions_WoundPoison:   proto.WeaponImbue_WoundPoison,
		}[poison]
		mh, oh = rogue.Consumes.MainHandImbue == imbue, rogue.Consumes.OffHandImbue == imbue
	}
	var mask core.ProcMask
	if rogue.HasMHWeapon() && mh {
		mask |= core.ProcMaskMeleeMH
	}
	if rogue.HasOHWeapon() && oh {
		mask |= core.ProcMaskMeleeOH
	}
	return mask
}
