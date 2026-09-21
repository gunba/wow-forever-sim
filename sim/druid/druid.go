package druid

import (
	"time"

	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

const SpellFlagBuilder = core.SpellFlagAgentReserved1

var TalentTreeSizes = [3]int{16, 19, 16}

const (
	SpellCode_DruidNone int32 = iota

	SpellCode_DruidClaw
	SpellCode_DruidFaerieFire
	SpellCode_DruidFaerieFireFeral
	SpellCode_DruidFerociousBite
	SpellCode_DruidHurricane
	SpellCode_DruidInsectSwarm
	SpellCode_DruidLacerate
	SpellCode_DruidMangle
	SpellCode_DruidMaul
	SpellCode_DruidMoonfire
	SpellCode_DruidRake
	SpellCode_DruidRip
	SpellCode_DruidShred
	SpellCode_DruidStarfire
	SpellCode_DruidSwipe
	SpellCode_DruidWrath
)

type Druid struct {
	core.Character
	SelfBuffs

	Talents *proto.DruidTalents

	DruidSpells []*DruidSpell

	StartingForm DruidForm

	RebirthTiming     float64
	BleedsActive      int
	AssumeBleedActive bool

	ReplaceBearMHFunc core.ReplaceMHSwing

	Barkskin             *DruidSpell
	Berserk              *DruidSpell
	DemoralizingRoar     *DruidSpell
	Enrage               *DruidSpell
	FaerieFire           *DruidSpell
	FerociousBite        *DruidSpell
	ForceOfNature        *DruidSpell
	FrenziedRegeneration *DruidSpell
	GiftOfTheWild        *DruidSpell
	Hurricane            []*DruidSpell
	Innervate            *DruidSpell
	InsectSwarm          []*DruidSpell
	Lacerate             *DruidSpell
	LacerateBleed        *DruidSpell
	Languish             *DruidSpell
	MangleBear           *DruidSpell
	MangleCat            *DruidSpell
	Maul                 *DruidSpell
	MaulQueueSpell       *DruidSpell
	Moonfire             []*DruidSpell
	Rebirth              *DruidSpell
	Rake                 *DruidSpell
	Rip                  *DruidSpell
	Shred                *DruidSpell
	Claw                 *DruidSpell
	Starfire             []*DruidSpell
	SwipeBear            *DruidSpell
	TigersFury           *DruidSpell
	Wrath                []*DruidSpell

	BearForm    *DruidSpell
	CatForm     *DruidSpell
	MoonkinForm *DruidSpell

	BarkskinAura             *core.Aura
	BearFormAura             *core.Aura
	BerserkAura              *core.Aura
	CatFormAura              *core.Aura
	DemoralizingRoarAuras    core.AuraArray
	EclipseAura              *core.Aura
	EnrageAura               *core.Aura
	FaerieFireAuras          core.AuraArray
	FrenziedRegenerationAura *core.Aura
	FurorAura                *core.Aura
	InsectSwarmAuras         core.AuraArray
	MaulQueueAura            *core.Aura
	MoonkinFormAura          *core.Aura
	NaturesGraceHasteAura    *core.Aura
	TigersFuryAura           *core.Aura

	BleedCategories core.ExclusiveCategoryArray

	form         DruidForm
	disabledMCDs []*core.MajorCooldown

	// Energy carried out of Cat Form, for the Forever version of Furor.
	lastCatFormEnergy float64
	lastCatFormExitAt time.Duration
}

type SelfBuffs struct {
	InnervateTarget *proto.UnitReference
}

func (druid *Druid) GetCharacter() *core.Character {
	return &druid.Character
}

func (druid *Druid) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	// Improved Mark of the Wild is baseline: the beta client's Mark and Gift of the Wild give 385 armor, 16 stats and
	// 27 resistances, Classic's 285 / 12 / 20 raised by 35%.
	if raidBuffs.GiftOfTheWild == proto.TristateEffect_TristateEffectRegular {
		raidBuffs.GiftOfTheWild = proto.TristateEffect_TristateEffectImproved
	}

}

func (druid *Druid) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	// These auras affect the druid's five-player subgroup in Forever. Keep
	// them on PartyBuffs rather than promoting them to raid-wide effects.
	if druid.InForm(Moonkin) {
		partyBuffs.MoonkinAura = true
	}

	if druid.InForm(Cat|Bear) && druid.Talents.LeaderOfThePack {
		partyBuffs.LeaderOfThePack = true
	}
}

func (druid *Druid) TryMaul(sim *core.Simulation, mhSwingSpell *core.Spell) *core.Spell {
	return druid.MaulReplaceMH(sim, mhSwingSpell)
}

func (druid *Druid) RegisterSpell(formMask DruidForm, config core.SpellConfig) *DruidSpell {
	prev := config.ExtraCastCondition
	prevModify := config.Cast.ModifyCast

	ds := &DruidSpell{FormMask: formMask}
	config.ExtraCastCondition = func(sim *core.Simulation, target *core.Unit) bool {
		// Check if we're in allowed form to cast
		// Allow 'humanoid' auto unshift casts
		if (ds.FormMask != Any && !druid.InForm(ds.FormMask)) && !ds.FormMask.Matches(Humanoid) {
			if sim.Log != nil {
				sim.Log("Failed cast to spell %s, wrong form", ds.ActionID)
			}
			return false
		}
		return prev == nil || prev(sim, target)
	}
	config.Cast.ModifyCast = func(sim *core.Simulation, s *core.Spell, c *core.Cast) {
		if !druid.InForm(ds.FormMask) && ds.FormMask.Matches(Humanoid) {
			druid.CancelShapeshift(sim)
		}
		if prevModify != nil {
			prevModify(sim, s, c)
		}
	}

	ds.Spell = druid.Unit.RegisterSpell(config)
	druid.DruidSpells = append(druid.DruidSpells, ds)

	return ds
}

func (druid *Druid) Initialize() {
	druid.BleedCategories = druid.GetEnemyExclusiveCategories(core.BleedEffectCategory)

	druid.registerFaerieFireSpell()
	druid.registerInnervateCD()
}

func (druid *Druid) RegisterBalanceSpells() {
	druid.registerHurricaneSpell()
	druid.registerInsectSwarmSpell()
	druid.registerMoonfireSpell()
	druid.registerStarfireSpell()
	druid.registerWrathSpell()
}

// TODO: Classic feral
func (druid *Druid) RegisterFeralCatSpells() {
	druid.registerCatFormSpell()
	// druid.registerBearFormSpell()
	// druid.registerEnrageSpell()
	druid.registerFerociousBiteSpell()
	druid.registerMangleCatSpell()
	// druid.registerMangleBearSpell()
	// druid.registerMaulSpell()
	druid.registerRakeSpell()
	druid.registerRipSpell()
	druid.registerShredSpell()
	druid.registerClawSpell()
	// druid.registerSwipeBearSpell()
	druid.registerTigersFurySpell()
	druid.registerBerserkCD()
}

func (druid *Druid) RegisterFeralTankSpells() {
	druid.registerBearFormSpell()
	druid.registerBarkskinCD()
	druid.registerBerserkCD()
	druid.registerDemoralizingRoarSpell()
	druid.registerEnrageSpell()
	druid.registerFrenziedRegenerationCD()
	druid.registerLacerateSpell()
	druid.registerMangleBearSpell()
	druid.registerMaulSpell()
	druid.registerSwipeBearSpell()
}

func (druid *Druid) Reset(_ *core.Simulation) {
	druid.BleedsActive = 0
	druid.form = druid.StartingForm
	druid.disabledMCDs = []*core.MajorCooldown{}
}

func New(character *core.Character, form DruidForm, selfBuffs SelfBuffs, talents string) *Druid {
	druid := &Druid{
		Character:    *character,
		SelfBuffs:    selfBuffs,
		Talents:      &proto.DruidTalents{},
		StartingForm: form,
		form:         form,
	}
	core.FillTalentsProto(druid.Talents.ProtoReflect(), talents, TalentTreeSizes)
	druid.EnableManaBar()

	druid.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	druid.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritPerAgiAtLevel[character.Class]*core.CritRatingPerCritChance)
	druid.AddStatDependency(stats.Agility, stats.Dodge, core.DodgePerAgiAtLevel[character.Class]*core.DodgeRatingPerDodgeChance)
	druid.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[character.Class]*core.SpellCritRatingPerCritChance)
	druid.AddStatDependency(stats.BonusArmor, stats.Armor, 1)

	// Druids get extra melee haste
	// druid.PseudoStats.MeleeHasteRatingPerHastePercent /= 1.3

	guardians.ConstructGuardians(&druid.Character)

	return druid
}

type DruidSpell struct {
	*core.Spell
	FormMask DruidForm
}

func (ds *DruidSpell) IsReady(sim *core.Simulation) bool {
	if ds == nil {
		return false
	}
	return ds.Spell.IsReady(sim)
}

func (ds *DruidSpell) CanCast(sim *core.Simulation, target *core.Unit) bool {
	if ds == nil {
		return false
	}
	return ds.Spell.CanCast(sim, target)
}

func (ds *DruidSpell) IsEqual(s *core.Spell) bool {
	if ds == nil || s == nil {
		return false
	}
	return ds.Spell == s
}

// Agent is a generic way to access underlying druid on any of the agents (for example balance druid.)
type DruidAgent interface {
	GetDruid() *Druid
}
