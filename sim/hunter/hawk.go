package hunter

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	hawkLifetime   = 18 * time.Second // 1293248, duration 85.
	hawkSwingSpeed = 2.5

	// Melee DPS of a level-60 creature, from the game's creature damage table (wowsims/forever
	// CreatureAutoAttackDpsMaxLevel).
	creatureAutoAttackDps60 = 60.45291519165039
)

// What a hawk is, read off level-30 combat logs (wowsims/forever 81651da66 and b89878e6f: 2,862 swings at 407 to
// 1,023 owner ranged attack power, and 4,340 swings with Ferocity 5/5 against targets 2 levels below to 3 above):
//   - no armor or attack power of its own; 80% of the owner's ranged attack power and 30% of its armor;
//   - half the owner's crit plus Ferocity's, with none from Agility;
//   - the owner's hit and expertise (1% hit took a hawk from 1 miss in 32 swings to none in 515);
//   - Unleashed Fury, and the owner's damage multipliers as it is summoned.
//
// A swing before armor is 0.0815 of the creature's DPS over the 2.5 s, a third either side, plus 0.0179 a point of
// the owner's ranged attack power (112 attack power a point of DPS at 80% of it). Both are fitted, not stated: the
// row that scales a hawk, Hunter Guardian Hawk Pet Scaling 01 (1293586), is not in the client data. The logs had
// Unleashed Fury 5/5, so both parts are stated before it.
const hawkLogUnleashedFury = 1 + 0.03*5

var hawkSwingAverage = 0.0815 * creatureAutoAttackDps60 * hawkSwingSpeed / hawkLogUnleashedFury

// Hawk is a temporary guardian, separate from the Hunter's permanent companion.
type Hawk struct {
	core.Pet

	hunterOwner          *Hunter
	spawnedAt            time.Duration
	expiresAt            time.Duration
	lastOwnerTarget      *core.Unit
	inheritedRangedSpeed float64
	inheritedMultiplier  float64
}

func (hunter *Hunter) hawkStatInheritance() core.PetStatInheritance {
	ferocity := 2 * float64(hunter.Talents.Ferocity) * core.CritRatingPerCritChance
	return func(owner stats.Stats) stats.Stats {
		return stats.Stats{
			stats.Armor:       0.3 * owner[stats.Armor],
			stats.AttackPower: 0.8 * owner[stats.RangedAttackPower],
			// The pet's MeleeCrit is the owner's higher of melee and ranged crit (core.ForeverPetInheritance).
			stats.MeleeCrit: 0.5*owner[stats.MeleeCrit] + ferocity,
			stats.MeleeHit:  owner[stats.MeleeHit],
			stats.Expertise: owner[stats.Expertise],
		}
	}
}

func (hunter *Hunter) newHawk(slot int) *Hawk {
	inheritance := hunter.hawkStatInheritance()
	hawk := &Hawk{
		Pet:                  core.NewPet(fmt.Sprintf("Hawk %d", slot+1), &hunter.Character, stats.Stats{}, inheritance, false, true),
		hunterOwner:          hunter,
		inheritedRangedSpeed: 1,
		inheritedMultiplier:  1,
	}
	hawk.EnableAutoAttacks(hawk, core.AutoAttackOptions{
		MainHand: core.Weapon{
			BaseDamageMin:     hawkSwingAverage * 2 / 3,
			BaseDamageMax:     hawkSwingAverage * 4 / 3,
			SwingSpeed:        hawkSwingSpeed,
			AttackPowerPerDPS: 112 * hawkLogUnleashedFury,
		},
		AutoSwingMelee: true,
	})
	config := hawk.AutoAttacks.MHConfig()
	config.ProcMask = core.ProcMaskEmpty
	config.Flags |= core.SpellFlagPassiveSpell | core.SpellFlagNoOnDamageDealt
	config.BonusCoefficient = 0
	// Unleashed Fury names hawks; the owner's multipliers come through the hawk's own (OnPetEnable).
	config.DamageMultiplier = 1 + .03*float64(hunter.Talents.UnleashedFury)
	config.ExtraCastCondition = func(sim *core.Simulation, _ *core.Unit) bool {
		hawk.updateTarget()
		return hawk.IsEnabled() && sim.CurrentTime < hawk.expiresAt && hawkTargetActive(hawk.CurrentTarget)
	}
	config.ApplyEffects = func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
		baseDamage := spell.Unit.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
		spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWhite)
	}

	hawk.OnPetEnable = func(sim *core.Simulation) {
		hawk.spawnedAt = sim.CurrentTime
		hawk.expiresAt = sim.CurrentTime + hawkLifetime
		hawk.lastOwnerTarget = hunter.CurrentTarget
		hawk.DistanceFromTarget = 0
		// The owner's stats reach the hawk while it is up, its damage multipliers as it is summoned.
		hawk.EnableDynamicStats(inheritance)
		hawk.inheritedMultiplier = hunter.PseudoStats.DamageDealtMultiplier
		hawk.PseudoStats.DamageDealtMultiplier *= hawk.inheritedMultiplier
		hawk.syncRangedSpeed(sim, hunter.RangedSwingSpeed())
		hawk.EnableDynamicRangedSpeed(func(ownerRangedSpeed float64) {
			hawk.syncRangedSpeed(sim, ownerRangedSpeed)
		})
		// Enable computes the current swing duration. A reused slot must attack
		// on arrival, not retain the previous guardian's unfinished swing.
		hawk.AutoAttacks.EnableAutoSwing(sim)
		hawk.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime-hawk.AutoAttacks.MainhandSwingSpeed(), false)
	}
	hawk.OnPetDisable = func(_ *core.Simulation) {
		hawk.PseudoStats.DamageDealtMultiplier /= hawk.inheritedMultiplier
		hawk.inheritedMultiplier = 1
	}
	hunter.AddPet(hawk)
	return hawk
}

func (hawk *Hawk) syncRangedSpeed(sim *core.Simulation, ownerRangedSpeed float64) {
	// Change only the inherited factor: normalizing the whole swing speed
	// would erase independent guardian-local buffs/slows. Disable retains
	// both the applied factor and unit speed; Reset returns both to baseline.
	hawk.MultiplyMeleeSpeed(sim, ownerRangedSpeed/hawk.inheritedRangedSpeed)
	hawk.inheritedRangedSpeed = ownerRangedSpeed
}

func hawkTargetActive(target *core.Unit) bool {
	return target.IsEnabled() && (!target.HasHealthBar() || target.CurrentHealth() > 0)
}

func (hawk *Hawk) updateTarget() {
	ownerTarget := hawk.hunterOwner.CurrentTarget
	// Logs show retargeting, but not its exact command rule. Preserve the
	// summon target until the owner switches focus or that target is inactive.
	if hawkTargetActive(ownerTarget) && (ownerTarget != hawk.lastOwnerTarget || !hawkTargetActive(hawk.CurrentTarget)) {
		hawk.CurrentTarget = ownerTarget
	}
	hawk.lastOwnerTarget = ownerTarget
}

func (hawk *Hawk) GetPet() *core.Pet { return &hawk.Pet }
func (hawk *Hawk) Initialize()       {}
func (hawk *Hawk) Reset(sim *core.Simulation) {
	hawk.Disable(sim)
	hawk.spawnedAt, hawk.expiresAt = 0, 0
	hawk.lastOwnerTarget = nil
	hawk.inheritedRangedSpeed = 1
	hawk.inheritedMultiplier = 1
}
func (hawk *Hawk) ExecuteCustomRotation(sim *core.Simulation) {
	hawk.WaitUntil(sim, hawk.AutoAttacks.NextAttackAt())
}

func (hunter *Hunter) summonHawkGuardian(sim *core.Simulation, target *core.Unit) {
	hawk := hunter.Hawks[0]
	for _, candidate := range hunter.Hawks {
		if !candidate.IsEnabled() {
			hawk = candidate
			break
		}
		if candidate.spawnedAt < hawk.spawnedAt {
			hawk = candidate
		}
	}
	hawk.Disable(sim)
	hawk.EnableWithTimeout(sim, hawk, hawkLifetime)
	hawk.CurrentTarget = target
}
