package hunter

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	hawkLifetime              = 18 * time.Second // 1293248, duration 85.
	hawkAssaultBaseSwingSpeed = 2.5

	// Oct 2 reports 2695/2701 calibrate rank 1, not a decoded server formula
	// or a measured level-60 guardian. Damage sensitivity: .30-.40 of rank
	// base. Effective crit sensitivity: .01-.03; no owner/Ferocity crit transfer.
	hawkAssaultDamageCoefficient   = 0.35
	hawkAssaultEffectiveCritChance = 0.015
)

// Hawk is a temporary guardian, separate from the Hunter's permanent companion.
type Hawk struct {
	core.Pet

	hunterOwner          *Hunter
	spawnedAt            time.Duration
	expiresAt            time.Duration
	lastOwnerTarget      *core.Unit
	inheritedRangedSpeed float64
}

func (hunter *Hunter) newHawk(slot int) *Hawk {
	hawk := &Hawk{
		Pet: core.NewPet(fmt.Sprintf("Hawk %d", slot+1), &hunter.Character,
			stats.Stats{stats.MeleeCrit: hawkAssaultEffectiveCritChance * 100 * core.CritRatingPerCritChance},
			func(stats.Stats) stats.Stats { return stats.Stats{} }, false, true),
		hunterOwner:          hunter,
		inheritedRangedSpeed: 1,
	}
	baseDamage := summonHawkBaseDamage[hunter.summonHawkRank()] * hawkAssaultDamageCoefficient
	hawk.EnableAutoAttacks(hawk, core.AutoAttackOptions{
		MainHand: core.Weapon{
			BaseDamageMin: baseDamage,
			BaseDamageMax: baseDamage,
			SwingSpeed:    hawkAssaultBaseSwingSpeed,
		},
		AutoSwingMelee: true,
	})
	config := hawk.AutoAttacks.MHConfig()
	config.ProcMask = core.ProcMaskEmpty
	config.Flags |= core.SpellFlagPassiveSpell | core.SpellFlagNoOnDamageDealt | core.SpellFlagIgnoreAttackerModifiers
	config.BonusCoefficient = 0
	// UF explicitly names Hawks. Apply it once here, not through owner damage
	// inheritance. The generic assault's transfer remains an intended-model choice.
	config.DamageMultiplier = 1 + .03*float64(hunter.Talents.UnleashedFury)
	config.ExtraCastCondition = func(sim *core.Simulation, _ *core.Unit) bool {
		hawk.updateTarget()
		return hawk.IsEnabled() && sim.CurrentTime < hawk.expiresAt && hawkTargetActive(hawk.CurrentTarget)
	}
	config.ApplyEffects = func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
		spell.CalcAndDealDamage(sim, target, baseDamage, hawk.assaultOutcome)
	}

	hawk.OnPetEnable = func(sim *core.Simulation) {
		hawk.spawnedAt = sim.CurrentTime
		hawk.expiresAt = sim.CurrentTime + hawkLifetime
		hawk.lastOwnerTarget = hunter.CurrentTarget
		hawk.DistanceFromTarget = 0
		hawk.syncRangedSpeed(sim, hunter.RangedSwingSpeed())
		hawk.EnableDynamicRangedSpeed(func(ownerRangedSpeed float64) {
			hawk.syncRangedSpeed(sim, ownerRangedSpeed)
		})
		// Enable computes the current swing duration. A reused slot must attack
		// on arrival, not retain the previous guardian's unfinished swing.
		hawk.AutoAttacks.EnableAutoSwing(sim)
		hawk.AutoAttacks.StopMeleeUntil(sim, sim.CurrentTime-hawk.AutoAttacks.MainhandSwingSpeed(), false)
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

func (hawk *Hawk) assaultOutcome(sim *core.Simulation, result *core.SpellResult, table *core.AttackTable) {
	spell := hawk.AutoAttacks.MHAuto()
	// This effective calibration is not an intrinsic crit stat subsequently
	// erased by boss suppression. The ordinary white miss/dodge/glance table
	// is provisional; the captured target levels and glancing flags are unknown.
	effectiveTable := *table
	effectiveTable.MeleeCritSuppression = 0
	spell.OutcomeMeleeWhite(sim, result, &effectiveTable)
}

func (hawk *Hawk) GetPet() *core.Pet { return &hawk.Pet }
func (hawk *Hawk) Initialize()       {}
func (hawk *Hawk) Reset(sim *core.Simulation) {
	hawk.Disable(sim)
	hawk.spawnedAt, hawk.expiresAt = 0, 0
	hawk.lastOwnerTarget = nil
	hawk.inheritedRangedSpeed = 1
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
