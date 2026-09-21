package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/hunter"
	googleProto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type budgetEntry struct {
	Stat              string
	GearAmount, Delta float64
	CostPerUnit       float64
	BudgetDelta       float64
}

type hitRequirement struct {
	Action            string
	Kind              string
	AdditionalPercent float64
}

type hitAdjustment struct {
	MeleeAdded, SpellAdded, MeleeFinal, SpellFinal float64
	Model                                          string
	RatingPerPercent                               float64
	RawHitDelta, OffensiveBudgetBefore             float64
	Balance                                        float64
	Entries                                        []budgetEntry
	Requirements                                   []hitRequirement
	Exclusions                                     []string
}

var readHitBudgetRates = sync.OnceValue(func() map[string]float64 {
	var data struct {
		Level        int
		Coefficients map[string]float64
	}
	if err := json.Unmarshal(mustRead("assets/db_inputs/forever_combat_ratings.json"), &data); err != nil {
		panic(err)
	}
	c := data.Coefficients
	if data.Level != 60 || c["Hit - Melee"] <= 0 ||
		c["Hit - Melee"] != c["Hit - Ranged"] || c["Hit - Melee"] != c["Hit - Spell"] ||
		c["Crit - Melee"] <= 0 || c["Crit - Melee"] != c["Crit - Spell"] {
		panic("hit-budget model requires the level-60 shared hit/crit coefficients")
	}
	return c
})

func casterBuild(b build) bool {
	return b.Class == proto.Class_ClassMage || b.Class == proto.Class_ClassWarlock ||
		b.Class == proto.Class_ClassPriest || b.Key == "balance" ||
		b.Key == "elemental" || b.Key == "stormcaller"
}

// Collect casts rather than every registered spell: an unused spell from another
// school must not erase a rotational hit talent's item-budget benefit.
func rotationCastIDs(rotation *proto.APLRotation) (map[core.ActionID]bool, bool) {
	ids := map[core.ActionID]bool{}
	cooldowns := false
	if rotation == nil {
		return ids, cooldowns
	}
	var visit func(protoreflect.Message)
	visit = func(message protoreflect.Message) {
		switch action := message.Interface().(type) {
		case *proto.APLListItem:
			if action.Hide {
				return
			}
		case *proto.APLPrepullAction:
			if action.Hide {
				return
			}
		case *proto.APLActionCastSpell:
			ids[core.ProtoToActionID(action.SpellId).WithTag(0)] = true
		case *proto.APLActionChannelSpell:
			ids[core.ProtoToActionID(action.SpellId).WithTag(0)] = true
		case *proto.APLActionMultidot:
			ids[core.ProtoToActionID(action.SpellId).WithTag(0)] = true
		case *proto.APLActionAutocastOtherCooldowns:
			cooldowns = true
		}
		message.Range(func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
			if field.Message() == nil || field.IsMap() {
				return true
			}
			if field.IsList() {
				for i := 0; i < value.List().Len(); i++ {
					visit(value.List().Get(i).Message())
				}
			} else {
				visit(value.Message())
			}
			return true
		})
	}
	visit(rotation.ProtoReflect())
	return ids, cooldowns
}

func hitRequirements(b build, p *proto.Player, character *core.Character, target *core.Unit) ([]hitRequirement, []string) {
	unit := &character.Unit
	ids, cooldowns := rotationCastIDs(p.Rotation)
	if cooldowns {
		for _, id := range character.GetMajorCooldownIDs() {
			ids[core.ProtoToActionID(id).WithTag(0)] = true
		}
	}
	autos := map[*core.Spell]bool{}
	if !casterBuild(b) {
		// The benchmark has a fixed position. A ranged hunter's unused melee
		// weapon must not cancel the benefit of a ranged-only scope.
		if unit.AutoAttacks.AutoSwingMelee && unit.DistanceFromTarget <= core.MaxMeleeAttackDistance {
			autos[unit.AutoAttacks.MHAuto()] = true
			if unit.AutoAttacks.IsDualWielding {
				autos[unit.AutoAttacks.OHAuto()] = true
			}
		}
		if unit.AutoAttacks.AutoSwingRanged && unit.DistanceFromTarget >= core.MinRangedAttackDistance {
			autos[unit.AutoAttacks.RangedAuto()] = true
		}
	}
	hasPoison := p.Consumes != nil &&
		(p.Consumes.MainHandImbue == proto.WeaponImbue_InstantPoison ||
			p.Consumes.MainHandImbue == proto.WeaponImbue_DeadlyPoison ||
			p.Consumes.OffHandImbue == proto.WeaponImbue_InstantPoison ||
			p.Consumes.OffHandImbue == proto.WeaponImbue_DeadlyPoison)
	var requirements []hitRequirement
	var exclusions []string
	for _, spell := range unit.Spellbook {
		poison := b.Class == proto.Class_ClassRogue && hasPoison && spell.Flags.Matches(core.SpellFlagPoison)
		if !ids[spell.ActionID.WithTag(0)] && !autos[spell] && !poison {
			continue
		}
		if spell.Flags.Matches(core.SpellFlagHelpful) || spell.DefenseType == core.DefenseTypeNone {
			continue
		}
		if b.Class == proto.Class_ClassHunter && spell.Flags.Matches(hunter.SpellFlagTrap) {
			exclusions = append(exclusions, spell.ActionID.String()+": current trap implementation ignores gear hit")
			continue
		}
		table := unit.AttackTables[target.Index][spell.CastType]
		if table == nil {
			continue
		}
		required := hitRequirement{Action: spell.ActionID.String()}
		if spell.DefenseType == core.DefenseTypeMagic {
			required.Kind = "spell"
			required.AdditionalPercent = 100*(table.BaseSpellMissChance-.01) -
				unit.GetStat(stats.SpellHit) - spell.BonusHitRating -
				unit.GetSchoolBonusHitChance(spell) - target.PseudoStats.BonusSpellHitRatingTaken
		} else {
			required.Kind = "physical-special"
			// Deliberately exclude the dual-wield white miss penalty and dodge.
			required.AdditionalPercent = 100*(table.BaseMissChance+table.HitSuppression) -
				unit.GetStat(stats.MeleeHit) - spell.BonusHitRating -
				target.PseudoStats.BonusMeleeHitRatingTaken
		}
		requirements = append(requirements, required)
	}
	return requirements, exclusions
}

type budgetDonor struct {
	name   string
	amount float64
	cost   float64
	apply  func(*stats.Stats, float64)
}

// Only the item and its selected random suffix supply budget. Base attributes,
// enchants, consumes, buffs, racials and explicit sensitivity stats cannot fund it.
func offensiveGearBudget(b build, gear stats.Stats, critCoefficient float64) []budgetDonor {
	var donors []budgetDonor
	add := func(name string, amount, cost float64, apply func(*stats.Stats, float64)) {
		if amount > 0 {
			donors = append(donors, budgetDonor{name, amount, cost, apply})
		}
	}
	single := func(stat stats.Stat, cost float64) {
		add(stat.StatName(), gear[stat], cost, func(delta *stats.Stats, amount float64) {
			delta[stat] += amount
		})
	}
	physical := !casterBuild(b)
	magic := !physical || b.Class == proto.Class_ClassPaladin || b.Class == proto.Class_ClassShaman ||
		b.Class == proto.Class_ClassHunter
	if physical {
		// Imported generic AP occupies both AP fields but represents one budget.
		shared := min(gear[stats.AttackPower], gear[stats.RangedAttackPower])
		add("AttackPower", shared, .5, func(delta *stats.Stats, amount float64) {
			delta[stats.AttackPower] += amount
			delta[stats.RangedAttackPower] += amount
		})
		if b.Class == proto.Class_ClassHunter {
			add("RangedAttackPower", gear[stats.RangedAttackPower]-shared, .5, func(delta *stats.Stats, amount float64) {
				delta[stats.RangedAttackPower] += amount
			})
		} else {
			add("MeleeAttackPower", gear[stats.AttackPower]-shared, .5, func(delta *stats.Stats, amount float64) {
				delta[stats.AttackPower] += amount
			})
		}
		single(stats.Agility, 1)
		if b.Class != proto.Class_ClassHunter || b.Key == "survival" {
			single(stats.Strength, 1)
		}
	}
	if magic {
		single(stats.SpellPower, 6.0/7.0)
		single(stats.SpellDamage, 6.0/7.0)
		single(stats.Intellect, 1)
	}
	add("CritRating", (gear[stats.MeleeCrit]+gear[stats.SpellCrit])*critCoefficient, 1,
		func(delta *stats.Stats, amount float64) {
			delta[stats.MeleeCrit] += amount / critCoefficient
			delta[stats.SpellCrit] += amount / critCoefficient
		})
	return donors
}

// Return a normalized clone. Talent-search inputs never acquire a previous
// candidate's hit bonus or stat debit.
func capHit(b build, input *proto.Player) (*proto.Player, hitAdjustment, error) {
	p := googleProto.Clone(input).(*proto.Player)
	if p.BonusStats == nil {
		p.BonusStats = &proto.UnitStats{}
	}
	bonus := stats.FromFloatArray(p.BonusStats.Stats)
	if bonus[stats.MeleeHit] != 0 || bonus[stats.SpellHit] != 0 {
		return nil, hitAdjustment{}, fmt.Errorf("hit normalization requires a baseline without manual hit overrides")
	}
	probe := request(p, 1, 1)
	env, _, _ := core.NewEnvironment(probe.Raid, probe.Encounter, proto.Ruleset_RulesetForever, false)
	unit, target := env.Raid.AllPlayerUnits[0], env.Encounter.TargetUnits[0]
	character := env.Raid.Parties[0].Players[0].GetCharacter()
	rates := readHitBudgetRates()
	report := hitAdjustment{Model: "linear-v1", RatingPerPercent: rates["Hit - Melee"]}
	report.Requirements, report.Exclusions = hitRequirements(b, p, character, target)
	if len(report.Requirements) == 0 {
		return nil, report, fmt.Errorf("%s: no rotational hit requirements found", b.Key)
	}
	additional := math.Inf(-1)
	for _, requirement := range report.Requirements {
		additional = max(additional, requirement.AdditionalPercent)
	}
	var gear stats.Stats
	for _, item := range character.Equipment {
		gear = gear.Add(item.Stats).Add(item.RandomSuffix.Stats)
	}
	gearHit := (gear[stats.MeleeHit] + gear[stats.SpellHit]) * report.RatingPerPercent
	// Excess item hit can be exchanged back, but a talent/racial by itself
	// cannot be cashed out into free offensive stats.
	report.RawHitDelta = max(additional*report.RatingPerPercent, -gearHit)
	donors := offensiveGearBudget(b, gear, rates["Crit - Melee"])
	for _, donor := range donors {
		report.OffensiveBudgetBefore += donor.amount * donor.cost
	}
	if report.RawHitDelta > report.OffensiveBudgetBefore+1e-8 {
		return nil, report, fmt.Errorf("%s: need %.2f hit budget, only %.2f offensive item budget available",
			b.Key, report.RawHitDelta, report.OffensiveBudgetBefore)
	}
	if report.OffensiveBudgetBefore == 0 && math.Abs(report.RawHitDelta) > 1e-8 {
		return nil, report, fmt.Errorf("%s: no offensive allocation to receive the hit-budget exchange", b.Key)
	}
	report.Entries = append(report.Entries, budgetEntry{
		Stat: "HitRating", GearAmount: gearHit, Delta: report.RawHitDelta,
		CostPerUnit: 1, BudgetDelta: report.RawHitDelta,
	})
	report.Balance = report.RawHitDelta
	if report.OffensiveBudgetBefore > 0 {
		// A uniform proportional exchange avoids class-specific cherry-picking
		// of cheap-to-lose stats. It is a benchmark rule, not an in-game reforge.
		fraction := report.RawHitDelta / report.OffensiveBudgetBefore
		for _, donor := range donors {
			amount := -donor.amount * fraction
			donor.apply(&bonus, amount)
			entry := budgetEntry{donor.name, donor.amount, amount, donor.cost, amount * donor.cost}
			report.Entries = append(report.Entries, entry)
			report.Balance += entry.BudgetDelta
		}
	}
	if math.Abs(report.Balance) > 1e-6 {
		return nil, report, fmt.Errorf("hit-budget ledger does not balance: %v", report.Balance)
	}
	report.MeleeAdded = report.RawHitDelta / report.RatingPerPercent
	report.SpellAdded = report.MeleeAdded
	bonus[stats.MeleeHit] += report.MeleeAdded
	bonus[stats.SpellHit] += report.SpellAdded
	report.MeleeFinal = unit.GetStat(stats.MeleeHit) + report.MeleeAdded
	report.SpellFinal = unit.GetStat(stats.SpellHit) + report.SpellAdded
	p.BonusStats.Stats = bonus.ToFloatArray()
	return p, report, nil
}
