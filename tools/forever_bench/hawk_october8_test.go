//go:build with_db

package main

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/hunter"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

func hawkFixture(talents map[string]int) (*core.Simulation, *hunter.Hunter) {
	if talents == nil {
		talents = map[string]int{"summonHawk": 1}
	}
	req := historyTalentFixture("survival", talents)
	player := req.Raid.Parties[0].Players[0]
	bow := googleProto.Clone(player.Equipment.Items[16]).(*proto.ItemSpec)
	player.Equipment = &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 17)}
	for i := range player.Equipment.Items {
		player.Equipment.Items[i] = &proto.ItemSpec{}
	}
	player.Equipment.Items[16] = bow
	player.Consumes = &proto.Consumes{}
	player.Buffs = &proto.IndividualBuffs{}
	player.ForeverTier1Bonuses = false
	player.GetHunter().Options.PetType = proto.Hunter_Options_PetNone
	player.GetHunter().Options.QuiverBonus = proto.Hunter_Options_Speed10
	req.Raid.Buffs = &proto.RaidBuffs{}
	req.Raid.Debuffs = &proto.Debuffs{JudgementOfWisdom: true}
	req.Raid.Parties[0].Buffs = &proto.PartyBuffs{}
	req.Encounter.Duration = 45
	req.Encounter.Targets[0].Stats = stats.Stats{}.ToFloatArray()
	req.Encounter.Targets[0].Level = 60
	req.Encounter.Targets = append(req.Encounter.Targets, googleProto.Clone(req.Encounter.Targets[0]).(*proto.Target))
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	for _, unit := range sim.Environment.AllUnits {
		unit.AutoAttacks.CancelAutoSwing(sim)
	}
	h := sim.Raid.Parties[0].Players[0].(hunter.HunterAgent).GetHunter()
	return sim, h
}

func hawkPets(h *hunter.Hunter) []*core.Pet {
	var hawks []*core.Pet
	for _, pet := range h.Pets {
		if strings.HasPrefix(pet.Name, "Hawk ") && pet.IsGuardian() {
			hawks = append(hawks, pet)
		}
	}
	return hawks
}

func advanceHawks(t *testing.T, sim *core.Simulation, at time.Duration) {
	t.Helper()
	core.StartDelayedAction(sim, core.DelayedActionOptions{DoAt: at + time.Nanosecond, OnAction: func(*core.Simulation) {}})
	for sim.CurrentTime <= at {
		if sim.Step() {
			t.Fatalf("fight ended before %s", at)
		}
	}
}

func TestHawkOctober8ProjectileTiming(t *testing.T) {
	sim, h := hawkFixture(nil)
	for _, tc := range []struct {
		distance float64
		want     time.Duration
	}{{0, time.Second}, {30, time.Second}, {70, 2 * time.Second}} {
		h.DistanceFromTarget = tc.distance
		if got := h.SummonHawk.TravelTime(); got != tc.want {
			t.Fatalf("Hawk travel at %gyards = %s, want %s", tc.distance, got, tc.want)
		}
	}
	h.DistanceFromTarget = 30
	spell := h.SummonHawk
	spell.BonusHitRating, spell.BonusCritRating = 10000, -10000
	spell.ApplyEffects(sim, h.CurrentTarget, spell)
	if spell.SpellMetrics[h.CurrentTarget.UnitIndex].TotalDamage != 0 {
		t.Fatal("Hawk opening dealt damage before arrival")
	}
	advanceHawks(t, sim, time.Second)
	if spell.SpellMetrics[h.CurrentTarget.UnitIndex].TotalDamage <= 0 {
		t.Fatal("Hawk opening did not deal damage on arrival")
	}
}

func TestHawkOctober8HasIndependentGuardians(t *testing.T) {
	sim, h := hawkFixture(nil)
	hawks := hawkPets(h)
	if len(hawks) != 2 {
		t.Fatalf("Hawk has %d guardian slots, want two independent pets", len(hawks))
	}
	h.DistanceFromTarget = 30
	first, second := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
	h.SummonHawk.ApplyEffects(sim, first, h.SummonHawk)
	advanceHawks(t, sim, time.Second)
	advanceHawks(t, sim, 6*time.Second)
	h.SummonHawk.ApplyEffects(sim, second, h.SummonHawk)
	advanceHawks(t, sim, 7*time.Second+time.Nanosecond)
	if !hawks[0].IsEnabled() || !hawks[1].IsEnabled() || hawks[0].CurrentTarget != first || hawks[1].CurrentTarget != second {
		t.Fatal("Hawks did not keep independent targets and overlapping lives")
	}
	advanceHawks(t, sim, 19*time.Second+time.Nanosecond)
	if hawks[0].IsEnabled() || !hawks[1].IsEnabled() {
		t.Fatal("a later summon refreshed the first Hawk's lifetime")
	}
}

func TestHawkOctober8AlwaysHitOpeningCanCrit(t *testing.T) {
	sim, h := hawkFixture(nil)
	h.DistanceFromTarget = 30
	target := h.CurrentTarget
	spell := h.SummonHawk
	table := h.AttackTables[target.UnitIndex][spell.CastType]
	table.BaseMissChance, table.BaseDodgeChance, table.BaseParryChance, table.BaseBlockChance = 1, 1, 1, 1
	spell.BonusCritRating = 10000
	want := 2 * (108 + .05*spell.RangedAttackPower(target, false))
	spell.ApplyEffects(sim, target, spell)
	advanceHawks(t, sim, time.Second)
	metrics := spell.SpellMetrics[target.UnitIndex]
	if metrics.Crits != 1 || metrics.Misses+metrics.Dodges+metrics.Parries+metrics.Blocks != 0 || math.Abs(metrics.TotalDamage-want) > 1e-9 {
		t.Fatalf("always-hit opening must crit for %g despite impossible hit table: %+v", want, metrics)
	}
	if !h.Hawks[0].IsEnabled() || h.Hawks[0].AutoAttacks.MHAuto().SpellMetrics[target.UnitIndex].Casts != 1 {
		t.Fatal("opening did not summon a guardian with an immediate arrival attack")
	}
	if h.SummonHawk.CD.Timer != h.ArcaneShot.CD.Timer || h.SummonHawk.CD.Duration != 6*time.Second ||
		h.SummonHawk.DefaultCast.GCD != 1500*time.Millisecond {
		t.Fatal("travel changed the shared Arcane Shot cooldown or Hunter GCD")
	}
}

func TestHawkOctober8ReplacesOldestGlobally(t *testing.T) {
	sim, h := hawkFixture(nil)
	h.DistanceFromTarget = 30
	first, second := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
	h.SummonHawk.ApplyEffects(sim, first, h.SummonHawk)
	advanceHawks(t, sim, time.Second)
	advanceHawks(t, sim, 6*time.Second)
	h.SummonHawk.ApplyEffects(sim, second, h.SummonHawk)
	advanceHawks(t, sim, 7*time.Second+time.Nanosecond)
	advanceHawks(t, sim, 12*time.Second)
	h.SummonHawk.ApplyEffects(sim, second, h.SummonHawk)
	advanceHawks(t, sim, 13*time.Second+time.Nanosecond)
	if h.Hawks[0].CurrentTarget != second || !h.Hawks[0].IsEnabled() || !h.Hawks[1].IsEnabled() || len(h.DynamicRangedSpeedPets) != 2 {
		t.Fatal("the third arrival did not replace the oldest of the two owner-global guardians")
	}
	advanceHawks(t, sim, 19*time.Second+2*time.Nanosecond)
	if !h.Hawks[0].IsEnabled() || !h.Hawks[1].IsEnabled() {
		t.Fatal("the replaced guardian's canceled timeout disabled a current guardian")
	}
	advanceHawks(t, sim, 25*time.Second+2*time.Nanosecond)
	if !h.Hawks[0].IsEnabled() || h.Hawks[1].IsEnabled() || len(h.DynamicRangedSpeedPets) != 1 {
		t.Fatal("the younger original guardian did not expire independently")
	}
	advanceHawks(t, sim, 31*time.Second+2*time.Nanosecond)
	if h.Hawks[0].IsEnabled() || len(h.DynamicRangedSpeedPets) != 0 {
		t.Fatal("replacement guardian or ranged-speed subscription survived its 18s lifetime")
	}
}

func assertHawkRemainingScale(t *testing.T, sim *core.Simulation, h *hunter.Hunter, factor float64, change func()) {
	t.Helper()
	before := [2]time.Duration{}
	for i, hawk := range h.Hawks {
		before[i] = hawk.AutoAttacks.MainhandSwingAt() - sim.CurrentTime
	}
	change()
	for i, hawk := range h.Hawks {
		want := time.Duration(float64(before[i]) / factor)
		got := hawk.AutoAttacks.MainhandSwingAt() - sim.CurrentTime
		if math.Abs(float64(got-want)) > 3 || math.Abs(hawk.SwingSpeed()-h.RangedSwingSpeed()) > 1e-9 {
			t.Fatalf("Hawk %d remaining %s, want %s; speed %g, owner ranged %g", i, got, want, hawk.SwingSpeed(), h.RangedSwingSpeed())
		}
	}
}

func TestHawkOctober8DynamicRangedSpeed(t *testing.T) {
	sim, h := hawkFixture(nil)
	h.DistanceFromTarget = 30
	h.SummonHawk.ApplyEffects(sim, h.CurrentTarget, h.SummonHawk)
	h.SummonHawk.ApplyEffects(sim, h.CurrentTarget, h.SummonHawk)
	advanceHawks(t, sim, time.Second)
	for _, hawk := range h.Hawks {
		if got := hawk.AutoAttacks.MainhandSwingSpeed(); math.Abs(float64(got)-float64(2500*time.Millisecond)/1.1) > 1 {
			t.Fatalf("initial Hawk interval %s did not inherit 1.10 ranged speed", got)
		}
	}
	advanceHawks(t, sim, 1500*time.Millisecond)
	assertHawkRemainingScale(t, sim, h, 1.3, func() { h.MultiplyRangedSpeed(sim, 1.3) })
	before := [2]time.Duration{h.Hawks[0].AutoAttacks.MainhandSwingAt(), h.Hawks[1].AutoAttacks.MainhandSwingAt()}
	h.MultiplyMeleeSpeed(sim, 2)
	h.AddStatDynamic(sim, stats.SpellHaste, 1000)
	for i, hawk := range h.Hawks {
		if hawk.AutoAttacks.MainhandSwingAt() != before[i] {
			t.Fatal("melee-only speed or casting haste altered a Hawk assault")
		}
	}
	assertHawkRemainingScale(t, sim, h, 1.2, func() { h.MultiplyAttackSpeed(sim, 1.2) })
	haste := float64(20 * core.HasteRatingPerHastePercent)
	assertHawkRemainingScale(t, sim, h, 1.2, func() { h.AddStatDynamic(sim, stats.MeleeHaste, haste) })
	assertHawkRemainingScale(t, sim, h, 1/1.2, func() { h.AddStatDynamic(sim, stats.MeleeHaste, -haste) })
	if len(h.DynamicRangedSpeedPets) != 2 || len(h.DynamicMeleeSpeedPets) != 0 {
		t.Fatal("Hawks duplicated speed callbacks or subscribed to melee-only changes")
	}
	disabled := h.Hawks[0]
	disabled.Disable(sim)
	oldSpeed := disabled.SwingSpeed()
	h.MultiplyRangedSpeed(sim, 2)
	if len(h.DynamicRangedSpeedPets) != 1 || disabled.SwingSpeed() != oldSpeed {
		t.Fatal("a disabled Hawk retained a ranged-speed callback")
	}
	advanceHawks(t, sim, 19*time.Second)
	if h.Hawks[1].IsEnabled() {
		t.Fatal("haste extended the guardian's lifetime")
	}
}

func TestHawkOctober8AssaultCalibrationAndIsolation(t *testing.T) {
	sim, h := hawkFixture(map[string]int{"summonHawk": 1, "unleashedFury": 5, "ferocity": 5, "improvedTracking": 5, "resourcefulness": 2})
	h.DistanceFromTarget = 30
	target := h.CurrentTarget
	guardian := h.Hawks[0]
	assault := guardian.AutoAttacks.MHAuto()
	table := guardian.AttackTables[target.UnitIndex][assault.CastType]
	table.BaseMissChance, table.BaseDodgeChance, table.BaseGlanceChance = 0, 0, 0
	assault.BonusCritRating = -10000
	ownerProcs, victimProcs := 0, 0
	trigger := h.GetAura("Resourcefulness Trigger")
	original := trigger.OnSpellHitDealt
	trigger.OnSpellHitDealt = func(a *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		if spell.Unit.Type == core.PetUnit {
			ownerProcs++
		}
		original(a, sim, spell, result)
	}
	wisdom := target.GetAura("Judgement of Wisdom")
	originalWisdom := wisdom.OnSpellHitTaken
	wisdom.OnSpellHitTaken = func(a *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		if spell.Unit.Type == core.PetUnit {
			victimProcs++
		}
		originalWisdom(a, sim, spell, result)
	}
	h.SummonHawk.ApplyEffects(sim, target, h.SummonHawk)
	advanceHawks(t, sim, time.Second)
	want := .35 * 108 * 1.15 // UF once; rank 60 is an extrapolation.
	if got := assault.SpellMetrics[target.UnitIndex].TotalDamage; math.Abs(got-want) > 1e-9 {
		t.Fatalf("guardian dealt %g, want calibrated %g", got, want)
	}
	beforeParent := h.SummonHawk.SpellMetrics[target.UnitIndex].TotalDamage
	h.PseudoStats.DamageDealtMultiplier *= 3
	h.AddStatsDynamic(sim, stats.Stats{stats.RangedAttackPower: 10000, stats.AttackPower: 10000, stats.MeleeCrit: 10000, stats.SpellDamage: 10000})
	assault.ApplyEffects(sim, target, assault)
	if got := assault.SpellMetrics[target.UnitIndex].TotalDamage; math.Abs(got-2*want) > 1e-9 ||
		guardian.GetStat(stats.AttackPower) != 0 || guardian.GetStat(stats.Agility) != 0 || guardian.GetStat(stats.SpellDamage) != 0 ||
		math.Abs(guardian.GetStat(stats.MeleeCrit)-1.5*core.CritRatingPerCritChance) > 1e-9 {
		t.Fatal("guardian inherited owner damage, power, Ferocity or crit instead of its own calibrated state")
	}
	if ownerProcs != 0 || victimProcs != 0 || assault.ProcMask != core.ProcMaskEmpty || !assault.Flags.Matches(core.SpellFlagNoOnDamageDealt) ||
		h.SummonHawk.SpellMetrics[target.UnitIndex].TotalDamage != beforeParent || len(h.SummonHawk.Dots()) != 0 {
		t.Fatal("guardian damage was attributed to the owner or created a proc/DoT opportunity")
	}
	sim.Cleanup()
	metrics := h.GetMetricsProto()
	if len(metrics.Pets) != 2 || metrics.Pets[0].Name != "Hawk 1" || len(metrics.Pets[0].Actions) == 0 || metrics.Pets[0].Dps.Avg <= 0 {
		t.Fatal("guardian assaults were not emitted as real pet metrics")
	}
}

func TestHawkOctober8EffectiveCritWithBossWhiteTable(t *testing.T) {
	sim, h := hawkFixture(nil)
	h.DistanceFromTarget = 30
	target := h.CurrentTarget
	target.Level = 63
	guardian := h.Hawks[0]
	assault := guardian.AutoAttacks.MHAuto()
	table := core.NewAttackTable(&guardian.Unit, target, nil)
	guardian.AttackTables[target.UnitIndex][assault.CastType] = table
	originalSuppression := table.MeleeCritSuppression
	genericTable := core.NewAttackTable(&h.Unit, target, nil)
	genericCrit := h.SummonHawk.PhysicalCritChance(genericTable)
	h.SummonHawk.ApplyEffects(sim, target, h.SummonHawk)
	advanceHawks(t, sim, time.Second)
	for i := 0; i < 2000; i++ {
		assault.ApplyEffects(sim, target, assault)
	}
	metrics := assault.SpellMetrics[target.UnitIndex]
	if metrics.Crits == 0 || metrics.Crits > 100 || metrics.Glances == 0 || metrics.Misses == 0 || metrics.Dodges == 0 {
		t.Fatalf("effective crit or standard boss white outcomes were lost: %+v", metrics)
	}
	if originalSuppression <= 0 || table.MeleeCritSuppression != originalSuppression || h.SummonHawk.PhysicalCritChance(genericTable) != genericCrit {
		t.Fatal("Hawk's local effective-crit scenario changed a shared attack table")
	}
}

func TestHawkOctober8RetargetsWithoutResummoning(t *testing.T) {
	sim, h := hawkFixture(nil)
	h.DistanceFromTarget = 30
	first, second := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
	h.SummonHawk.ApplyEffects(sim, first, h.SummonHawk)
	advanceHawks(t, sim, time.Second)
	guardian := h.Hawks[0]
	h.CurrentTarget = second
	advanceHawks(t, sim, guardian.AutoAttacks.MainhandSwingAt())
	if guardian.CurrentTarget != second || !guardian.IsEnabled() || h.Hawks[1].IsEnabled() {
		t.Fatal("owner target change did not retarget the existing guardian")
	}
	// An original target may die while the owner's focus has not changed.
	h.SummonHawk.ApplyEffects(sim, first, h.SummonHawk)
	advanceHawks(t, sim, sim.CurrentTime+time.Second)
	first.EnableHealthBar()
	if first.CurrentHealth() != 0 {
		t.Fatal("test target should be dead")
	}
	advanceHawks(t, sim, h.Hawks[1].AutoAttacks.MainhandSwingAt())
	if h.Hawks[1].CurrentTarget != second || !h.Hawks[1].IsEnabled() {
		t.Fatal("guardian did not leave its dead summon target for the owner's focus")
	}
	advanceHawks(t, sim, 19*time.Second)
	if guardian.IsEnabled() {
		t.Fatal("retargeting refreshed the guardian's lifetime")
	}
}

func TestHawkOctober8ResetAndUntalentedCompanion(t *testing.T) {
	sim, h := hawkFixture(nil)
	for iteration := 0; iteration < 2; iteration++ {
		if iteration > 0 {
			sim.Reset()
			for _, unit := range sim.Environment.AllUnits {
				unit.AutoAttacks.CancelAutoSwing(sim)
			}
		}
		if h.Hawks[0].IsEnabled() || h.Hawks[1].IsEnabled() || len(h.DynamicRangedSpeedPets) != 0 {
			t.Fatal("reset retained a Hawk or subscription")
		}
		h.DistanceFromTarget = 30
		h.SummonHawk.ApplyEffects(sim, h.CurrentTarget, h.SummonHawk)
		advanceHawks(t, sim, time.Second)
		if !h.Hawks[0].IsEnabled() || h.Hawks[1].IsEnabled() || len(h.DynamicRangedSpeedPets) != 1 || math.Abs(h.Hawks[0].SwingSpeed()-1.1) > 1e-9 {
			t.Fatal("repeated iteration retained stale speed or duplicated guardians")
		}
		h.MultiplyRangedSpeed(sim, 2)
		sim.Cleanup()
		if h.Hawks[0].IsEnabled() || len(h.DynamicRangedSpeedPets) != 0 {
			t.Fatal("iteration cleanup retained a guardian")
		}
	}
	req := historyTalentFixture("beast_mastery", map[string]int{"ferocity": 5})
	ordinary := core.NewSim(req, simsignals.Signals{})
	ordinary.Options.Interactive = true
	ordinary.Reset()
	owner := ordinary.Raid.Parties[0].Players[0].(hunter.HunterAgent).GetHunter()
	if len(hawkPets(owner)) != 0 || len(owner.Pets) != 1 || owner.Pets[0].IsGuardian() {
		t.Fatal("untalented Hunter's ordinary companion changed")
	}
	companionSpeed, companionCrit := owner.Pets[0].SwingSpeed(), owner.Pets[0].GetStat(stats.MeleeCrit)
	owner.MultiplyRangedSpeed(ordinary, 2)
	if owner.Pets[0].SwingSpeed() != companionSpeed || owner.Pets[0].GetStat(stats.MeleeCrit) != companionCrit || len(owner.DynamicRangedSpeedPets) != 0 {
		t.Fatal("the new guardian ranged-speed hook changed an ordinary companion")
	}
}

func hawkAPLValues(h *hunter.Hunter) (core.APLValue, core.APLValue) {
	count := h.NewAPLValue(h.Rotation, &proto.APLValue{Value: &proto.APLValue_HunterHawkCount{HunterHawkCount: &proto.APLValueHunterHawkCount{}}})
	remaining := h.NewAPLValue(h.Rotation, &proto.APLValue{Value: &proto.APLValue_HunterHawkRemainingTime{HunterHawkRemainingTime: &proto.APLValueHunterHawkRemainingTime{}}})
	return count, remaining
}

func assertHawkAPLState(t *testing.T, sim *core.Simulation, count, remaining core.APLValue, wantCount int32, wantExpiry time.Duration) {
	t.Helper()
	wantRemaining := time.Duration(0)
	if wantCount > 0 {
		wantRemaining = wantExpiry - sim.CurrentTime
	}
	if count.Type() != proto.APLValueType_ValueTypeInt || remaining.Type() != proto.APLValueType_ValueTypeDuration {
		t.Fatal("Hawk APL values have incorrect native types")
	}
	if got := count.GetInt(sim); got != wantCount {
		t.Fatalf("active Hawks %d, want %d", got, wantCount)
	}
	if got := remaining.GetDuration(sim); got != wantRemaining {
		t.Fatalf("earliest Hawk remaining %s, want %s", got, wantRemaining)
	}
}

func TestHawkOctober8APLOwnerGlobalState(t *testing.T) {
	sim, h := hawkFixture(nil)
	count, remaining := hawkAPLValues(h)
	h.DistanceFromTarget = 30
	first, second := sim.Encounter.TargetUnits[0], sim.Encounter.TargetUnits[1]
	assertHawkAPLState(t, sim, count, remaining, 0, 0)
	h.SummonHawk.ApplyEffects(sim, first, h.SummonHawk)
	advanceHawks(t, sim, 500*time.Millisecond)
	assertHawkAPLState(t, sim, count, remaining, 0, 0) // In flight, not active.
	advanceHawks(t, sim, time.Second)
	assertHawkAPLState(t, sim, count, remaining, 1, 19*time.Second)
	advanceHawks(t, sim, 6*time.Second)
	h.SummonHawk.ApplyEffects(sim, second, h.SummonHawk)
	assertHawkAPLState(t, sim, count, remaining, 1, 19*time.Second)
	advanceHawks(t, sim, 7*time.Second+time.Nanosecond)
	assertHawkAPLState(t, sim, count, remaining, 2, 19*time.Second)
	h.CurrentTarget = second
	assertHawkAPLState(t, sim, count, remaining, 2, 19*time.Second)
	advanceHawks(t, sim, 12*time.Second)
	h.SummonHawk.ApplyEffects(sim, second, h.SummonHawk)
	advanceHawks(t, sim, 13*time.Second+time.Nanosecond)
	assertHawkAPLState(t, sim, count, remaining, 2, 25*time.Second+time.Nanosecond)
	advanceHawks(t, sim, 25*time.Second+2*time.Nanosecond)
	assertHawkAPLState(t, sim, count, remaining, 1, 31*time.Second+time.Nanosecond)
	advanceHawks(t, sim, 31*time.Second+2*time.Nanosecond)
	assertHawkAPLState(t, sim, count, remaining, 0, 0)
	sim.Cleanup()
	sim.Reset()
	assertHawkAPLState(t, sim, count, remaining, 0, 0)
	untalentedSim, untalented := hawkFixture(map[string]int{})
	untalentedCount, untalentedRemaining := hawkAPLValues(untalented)
	assertHawkAPLState(t, untalentedSim, untalentedCount, untalentedRemaining, 0, 0)
}

func TestHawkOctober8APLJSONRoundTrip(t *testing.T) {
	for _, input := range []string{`{"hunterHawkCount":{}}`, `{"hunterHawkRemainingTime":{}}`} {
		value := &proto.APLValue{}
		if err := protojson.Unmarshal([]byte(input), value); err != nil {
			t.Fatal(err)
		}
		encoded, err := protojson.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		roundTrip := &proto.APLValue{}
		if err := protojson.Unmarshal(encoded, roundTrip); err != nil || !googleProto.Equal(value, roundTrip) {
			t.Fatalf("Hawk value JSON round trip failed: %s, %v", encoded, err)
		}
		binary, err := googleProto.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		if err := googleProto.Unmarshal(binary, roundTrip); err != nil || !googleProto.Equal(value, roundTrip) {
			t.Fatalf("Hawk value binary round trip failed: %v", err)
		}
	}
}

func TestHawkOctober8APLComparisonsBuildWithoutWarnings(t *testing.T) {
	rotation := &proto.APLRotation{}
	if err := protojson.Unmarshal([]byte(`{
		"type":"TypeAPL",
		"priorityList":[
			{"action":{"condition":{"cmp":{"op":"OpLt","lhs":{"hunterHawkCount":{}},"rhs":{"const":{"val":"1"}}}},"castSpell":{"spellId":{"spellId":1293527}}}},
			{"action":{"condition":{"cmp":{"op":"OpLt","lhs":{"hunterHawkCount":{}},"rhs":{"const":{"val":"2"}}}},"castSpell":{"spellId":{"spellId":1293527}}}},
			{"action":{"condition":{"cmp":{"op":"OpLe","lhs":{"hunterHawkRemainingTime":{}},"rhs":{"spellTravelTime":{"spellId":{"spellId":1293527}}}}},"castSpell":{"spellId":{"spellId":1293527}}}}
		]
	}`), rotation); err != nil {
		t.Fatal(err)
	}
	req := historyTalentFixture("survival", map[string]int{"summonHawk": 1})
	req.Raid.Parties[0].Players[0].Rotation = rotation
	_, raidStats, _ := core.NewEnvironment(req.Raid, req.Encounter, req.SimOptions.Ruleset, false)
	rotationStats := raidStats.Parties[0].Players[0].RotationStats
	if rotationStats == nil || len(rotationStats.PriorityList) != 3 {
		t.Fatal("Hawk APL predicates did not build")
	}
	for i, action := range rotationStats.PriorityList {
		if len(action.Warnings) != 0 {
			t.Fatalf("Hawk predicate %d: %v", i, action.Warnings)
		}
	}
}

func TestHawkOctober8LocalSpeedModifierSurvivesInheritance(t *testing.T) {
	sim, h := hawkFixture(nil)
	h.DistanceFromTarget = 30
	h.SummonHawk.ApplyEffects(sim, h.CurrentTarget, h.SummonHawk)
	h.SummonHawk.ApplyEffects(sim, h.CurrentTarget, h.SummonHawk)
	advanceHawks(t, sim, 1500*time.Millisecond)
	first, second := h.Hawks[0], h.Hawks[1]
	first.MultiplyMeleeSpeed(sim, .8)
	oldSpeed := first.SwingSpeed()
	oldRemaining := first.AutoAttacks.MainhandSwingAt() - sim.CurrentTime
	h.MultiplyRangedSpeed(sim, 1.3)
	if got := first.SwingSpeed(); math.Abs(got-oldSpeed*1.3) > 1e-9 {
		t.Fatalf("owner ranged inheritance erased guardian-local .8 modifier: got %g, want %g", got, oldSpeed*1.3)
	}
	if got, want := first.AutoAttacks.MainhandSwingAt()-sim.CurrentTime, time.Duration(float64(oldRemaining)/1.3); math.Abs(float64(got-want)) > 3 {
		t.Fatalf("owner haste rescaled the locally slowed swing twice: got %s, want %s", got, want)
	}
	if math.Abs(second.SwingSpeed()-h.RangedSwingSpeed()) > 1e-9 || len(h.DynamicRangedSpeedPets) != 2 {
		t.Fatal("local modifier leaked to another guardian or duplicated inheritance")
	}
	h.MultiplyRangedSpeed(sim, 1/1.3)
	if math.Abs(first.SwingSpeed()-oldSpeed) > 1e-9 {
		t.Fatal("removing owner haste erased the local slow")
	}
	h.MultiplyAttackSpeed(sim, 1.2)
	haste := float64(20 * core.HasteRatingPerHastePercent)
	h.AddStatDynamic(sim, stats.MeleeHaste, haste)
	if math.Abs(first.SwingSpeed()-h.RangedSwingSpeed()*.8) > 1e-9 {
		t.Fatal("general speed or haste-stat inheritance erased the local slow")
	}
	h.AddStatDynamic(sim, stats.MeleeHaste, -haste)
	h.MultiplyAttackSpeed(sim, 1/1.2)
	first.MultiplyMeleeSpeed(sim, 1/.8)
	if math.Abs(first.SwingSpeed()-h.RangedSwingSpeed()) > 1e-9 {
		t.Fatal("removing a local modifier did not restore the inherited clock")
	}

	// Disable stops inheritance, but does not undo the last applied factor.
	// Reusing the slot must update that factor once, not normalize all speed.
	first.MultiplyMeleeSpeed(sim, .8)
	h.MultiplyRangedSpeed(sim, 1.3)
	first.Disable(sim)
	disabledSpeed := first.SwingSpeed()
	h.MultiplyRangedSpeed(sim, 1/1.3)
	if first.SwingSpeed() != disabledSpeed {
		t.Fatal("disabled guardian received an owner speed update")
	}
	h.SummonHawk.ApplyEffects(sim, h.CurrentTarget, h.SummonHawk)
	advanceHawks(t, sim, 2500*time.Millisecond+time.Nanosecond)
	if math.Abs(first.SwingSpeed()-h.RangedSwingSpeed()*.8) > 1e-9 || len(h.DynamicRangedSpeedPets) != 2 {
		t.Fatal("slot reuse lost a local modifier or applied inherited speed twice")
	}
	sim.Cleanup()
	sim.Reset()
	if math.Abs(first.SwingSpeed()-1) > 1e-9 || len(h.DynamicRangedSpeedPets) != 0 {
		t.Fatal("reset retained guardian-local speed or callbacks")
	}
	h.SummonHawk.ApplyEffects(sim, h.CurrentTarget, h.SummonHawk)
	advanceHawks(t, sim, time.Second)
	if math.Abs(first.SwingSpeed()-h.RangedSwingSpeed()) > 1e-9 || len(h.DynamicRangedSpeedPets) != 1 {
		t.Fatal("new iteration reused a stale inherited factor or local modifier")
	}
}
