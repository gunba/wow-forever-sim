//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/hunter"
	"github.com/wowsims/classic/sim/mage"
	"github.com/wowsims/classic/sim/priest"
)

func TestClearcastingConsumptionAndCooldown(t *testing.T) {
	req := mageSpellFixture(map[string]int{"arcaneConcentration": 5})
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	unit := sim.Raid.AllPlayerUnits[0]
	aura := unit.GetAura("Clearcasting")
	trigger := unit.GetAura("Arcane Concentration")
	fireball := unit.GetSpell(core.ActionID{SpellID: 133})
	if aura.ActionID.SpellID != 12536 || trigger.Icd.Duration != time.Second {
		t.Fatal("Clearcasting must use the client aura and one-second proc cooldown")
	}
	paid := fireball.Cost.GetCurrentCost()
	aura.Activate(sim)
	if fireball.Cost.GetCurrentCost() != 0 {
		t.Fatal("Clearcasting did not waive the cost")
	}
	sim.CurrentTime = time.Second
	free := &core.Spell{Flags: mage.SpellFlagMage, ProcMask: core.ProcMaskSpellDamage}
	aura.OnCastComplete(aura, sim, free)
	if !aura.IsActive() {
		t.Fatal("an intrinsically free action consumed Clearcasting")
	}
	aura.OnCastComplete(aura, sim, fireball)
	if aura.IsActive() || fireball.Cost.GetCurrentCost() != paid {
		t.Fatal("the next paid damage spell must consume Clearcasting and restore costs")
	}
	// Multiple hits at the same timestamp cannot refresh the proc repeatedly.
	trigger.Icd.Use(sim)
	for i := 0; i < 500; i++ {
		trigger.OnSpellHitDealt(trigger, sim, fireball, &core.SpellResult{
			Target: unit.CurrentTarget, Outcome: core.OutcomeHit,
		})
	}
	if aura.IsActive() {
		t.Fatal("Arcane Concentration ignored its internal cooldown")
	}
	sim.CurrentTime += time.Second
	for i := 0; i < 500 && !aura.IsActive(); i++ {
		trigger.OnSpellHitDealt(trigger, sim, fireball, &core.SpellResult{
			Target: unit.CurrentTarget, Outcome: core.OutcomeHit,
		})
	}
	if !aura.IsActive() {
		t.Fatal("Arcane Concentration did not become eligible after its cooldown")
	}
}

func TestLightningShieldDoesNotResetShocks(t *testing.T) {
	for _, id := range []int32{324, 10432} {
		t.Run(core.ActionID{SpellID: id}.String(), func(t *testing.T) {
			req := historyTalentFixture("elemental", nil)
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			unit := sim.Raid.AllPlayerUnits[0]
			shock := unit.GetSpell(core.ActionID{SpellID: 10414})
			shield := unit.GetSpell(core.ActionID{SpellID: id})
			if !shock.Cast(sim, unit.CurrentTarget) {
				t.Fatal("Earth Shock failed to cast")
			}
			readyAt := shock.CD.ReadyAt()
			sim.CurrentTime = 2 * time.Second
			if shock.IsReady(sim) || !shield.Cast(sim, unit) {
				t.Fatal("expected a cooling-down shock and a castable Lightning Shield")
			}
			if shock.CD.ReadyAt() != readyAt || shock.IsReady(sim) {
				t.Fatal("applying Lightning Shield reset Earth Shock's cooldown")
			}
		})
	}
}

func TestStarshardsSharedCooldown(t *testing.T) {
	req := racialFixture("smite", proto.Race_RaceNightElf)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	p := sim.Raid.Parties[0].Players[0].(priest.PriestAgent).GetPriest()
	full := p.Starshards[priest.StarshardsRanks][0]
	if !full.Cast(sim, p.CurrentTarget) {
		t.Fatal("Starshards was not castable")
	}
	if dot := full.Dot(p.CurrentTarget); dot.IsActive() {
		dot.Cancel(sim)
	}
	sim.CurrentTime = 29 * time.Second
	for rank, variants := range p.Starshards {
		for tag, spell := range variants {
			if spell == nil {
				continue
			}
			if spell.CD.Duration != 30*time.Second || spell.CD.Timer != full.CD.Timer || spell.IsReady(sim) {
				t.Fatalf("rank %d tag %d bypassed the shared cooldown", rank, tag)
			}
		}
	}
	sim.CurrentTime = 30 * time.Second
	if !full.IsReady(sim) {
		t.Fatal("Starshards cooldown did not end at 30 seconds")
	}
}

func TestForeverProjectileSpeeds(t *testing.T) {
	for _, tc := range []struct {
		key    string
		speeds map[int32]float64
	}{
		{"marksmanship", map[int32]float64{20904: 40, 14287: 40, 2643: 30, 25295: 40, 1310786: 60}},
		{"destruction", map[int32]float64{686: 20, 25307: 20, 1293813: 20}},
	} {
		t.Run(tc.key, func(t *testing.T) {
			req := racialFixture(tc.key, proto.Race_RaceOrc)
			sim := core.NewSim(req, simsignals.Signals{})
			unit := sim.Raid.AllPlayerUnits[0]
			for id, speed := range tc.speeds {
				spell := unit.GetSpell(core.ActionID{SpellID: id})
				if spell == nil || spell.MissileSpeed != speed {
					t.Errorf("spell %d: want projectile speed %v, got %+v", id, speed, spell)
				}
			}
			if tc.key == "marksmanship" && unit.AutoAttacks.RangedConfig().MissileSpeed != 40 {
				t.Fatal("Auto Shot has the wrong projectile speed")
			}
		})
	}
}

func TestHunterPetBiteDamageRange(t *testing.T) {
	req := racialFixture("marksmanship", proto.Race_RaceOrc)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	h := sim.Raid.Parties[0].Players[0].(*hunter.Hunter)
	pet := h.Pets[0]
	bite := pet.GetSpell(core.ActionID{SpellID: 17261})
	if bite == nil {
		t.Fatal("fixture requires rank-eight Bite")
	}
	minRoll, maxRoll := math.Inf(1), math.Inf(-1)
	for i := 0; i < 2000; i++ {
		metrics := &bite.SpellMetrics[pet.CurrentTarget.UnitIndex]
		hits, blocks, damage := metrics.Hits, metrics.Blocks, metrics.TotalDamage
		bite.ApplyEffects(sim, pet.CurrentTarget, bite)
		if metrics.Hits != hits+1 || metrics.Blocks != blocks {
			continue
		}
		dealt := metrics.TotalDamage - damage
		low := bite.CalcDamage(sim, pet.CurrentTarget, 81, bite.OutcomeAlwaysHit).Damage
		high := bite.CalcDamage(sim, pet.CurrentTarget, 99, bite.OutcomeAlwaysHit).Damage
		roll := 81 + 18*(dealt-low)/(high-low)
		minRoll, maxRoll = min(minRoll, roll), max(maxRoll, roll)
	}
	if minRoll < 81-1e-8 || minRoll > 81.5 || maxRoll < 98.5 || maxRoll > 99+1e-8 {
		t.Fatalf("Bite sampled range = %.3f–%.3f, want 81–99", minRoll, maxRoll)
	}
}

func TestImprovedStormstrikeManaTicks(t *testing.T) {
	req := historyTalentFixture("enhancement", map[string]int{"stormstrike": 1, "improvedStormstrike": 2})
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	unit := sim.Raid.AllPlayerUnits[0]
	aura := unit.GetAura("Improved Stormstrike")
	if aura == nil {
		t.Fatal("fixture requires Improved Stormstrike")
	}
	sim.CurrentTime = time.Second
	unit.PseudoStats.FiveSecondRuleRefreshTime = 10 * time.Second
	unit.SpendMana(sim, 1000, unit.NewManaMetrics(core.ActionID{SpellID: 990001}))
	base := unit.ManaRegenPerSecondWhileCasting()
	checkTick := func(want float64) {
		t.Helper()
		start := unit.CurrentMana()
		unit.ManaTick(sim)
		if got := unit.CurrentMana() - start; math.Abs(got-want) > 1e-8 {
			t.Fatalf("mana tick = %v, want %v", got, want)
		}
	}
	checkTick(2 * base)
	aura.Activate(sim)
	boosted := unit.ManaRegenPerSecondWhileCasting()
	if boosted <= base {
		t.Fatal("Improved Stormstrike did not increase casting regeneration")
	}
	checkTick(2 * boosted)
	aura.Deactivate(sim)
	checkTick(2 * base)
}

func TestSpiritTapManaTicks(t *testing.T) {
	req := historyTalentFixture("shadow", map[string]int{"spiritTap": 5})
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	unit := sim.Raid.AllPlayerUnits[0]
	aura := unit.GetAura("Spirit Tap")
	if aura == nil {
		t.Fatal("fixture requires Spirit Tap")
	}
	sim.CurrentTime = time.Second
	unit.PseudoStats.FiveSecondRuleRefreshTime = 10 * time.Second
	unit.SpendMana(sim, 1000, unit.NewManaMetrics(core.ActionID{SpellID: 990002}))
	base := unit.ManaRegenPerSecondWhileCasting()
	for _, active := range []bool{true, false} {
		if active {
			aura.Activate(sim)
		} else {
			aura.Deactivate(sim)
		}
		want := unit.ManaRegenPerSecondWhileCasting()
		if active && want <= base || !active && math.Abs(want-base) > 1e-8 {
			t.Fatalf("incorrect Spirit Tap regeneration transition: %v -> %v", base, want)
		}
		start := unit.CurrentMana()
		unit.ManaTick(sim)
		if got := unit.CurrentMana() - start; math.Abs(got-2*want) > 1e-8 {
			t.Fatalf("Spirit Tap active=%v: tick %v, want %v", active, got, 2*want)
		}
	}
}

func TestMutilatePuncturingWoundsBothHands(t *testing.T) {
	var base [2]float64
	for points := 0; points <= 3; points++ {
		req := historyTalentFixture("mutilate", map[string]int{"mutilate": 1, "puncturingWounds": points})
		sim := core.NewSim(req, simsignals.Signals{})
		unit := sim.Raid.AllPlayerUnits[0]
		for i, tag := range []int32{0, 2} {
			spell := unit.GetSpell(core.ActionID{SpellID: 1241584, Tag: tag})
			if spell == nil {
				t.Fatalf("missing Mutilate hand %d", tag)
			}
			if points == 0 {
				base[i] = spell.BonusCritRating
			}
			want := base[i] + float64(points)*5*core.CritRatingPerCritChance
			if spell.BonusCritRating != want {
				t.Fatalf("rank %d hand %d: crit bonus %v, want %v", points, tag, spell.BonusCritRating, want)
			}
		}
	}
}

func TestHunterTemporaryManaRegenTicks(t *testing.T) {
	req := historyTalentFixture("marksmanship", map[string]int{"resourcefulness": 2, "rapidRecuperation": 2})
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	unit := sim.Raid.AllPlayerUnits[0]
	sim.CurrentTime = time.Second
	unit.PseudoStats.FiveSecondRuleRefreshTime = 10 * time.Second
	unit.SpendMana(sim, 2000, unit.NewManaMetrics(core.ActionID{SpellID: 990003}))
	base := unit.ManaRegenPerSecondWhileCasting()
	for _, label := range []string{"Resourcefulness", "Rapid Recuperation"} {
		aura := unit.GetAura(label)
		if aura == nil {
			t.Fatalf("fixture missing %s", label)
		}
		for _, active := range []bool{true, false} {
			if active {
				aura.Activate(sim)
			} else {
				aura.Deactivate(sim)
			}
			want := unit.ManaRegenPerSecondWhileCasting()
			if active && want <= base || !active && math.Abs(want-base) > 1e-8 {
				t.Fatalf("%s regeneration transition: %v -> %v", label, base, want)
			}
			start := unit.CurrentMana()
			unit.ManaTick(sim)
			if got := unit.CurrentMana() - start; math.Abs(got-2*want) > 1e-8 {
				t.Fatalf("%s active=%v: tick %v, want %v", label, active, got, 2*want)
			}
		}
	}
}

func TestPaladinRanksShareCooldowns(t *testing.T) {
	for _, ids := range [][]int32{
		{1311606, 20473, 20929, 20930},
		{879, 5614, 5615, 10312, 10313, 10314},
		{2812, 10318},
		{20925, 20927, 20928},
	} {
		req := historyTalentFixture("retribution", map[string]int{"holyShock": 1, "holyShield": 1})
		p := req.Raid.Parties[0].Players[0]
		p.Equipment.Items[14] = &proto.ItemSpec{Id: 279261}
		p.Equipment.Items[15] = &proto.ItemSpec{Id: 279262}
		req.Encounter.Targets[0].MobType = proto.MobType_MobTypeDemon
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		unit := sim.Raid.AllPlayerUnits[0]
		highest := unit.GetSpell(core.ActionID{SpellID: ids[len(ids)-1]})
		if highest == nil || !highest.Cast(sim, sim.Encounter.TargetUnits[0]) {
			t.Fatalf("could not cast highest rank in %v", ids)
		}
		for _, id := range ids {
			spell := unit.GetSpell(core.ActionID{SpellID: id})
			if spell == nil || spell.CD.Timer != highest.CD.Timer || spell.CD.IsReady(sim) {
				t.Fatalf("rank %d bypassed its shared cooldown", id)
			}
		}
	}
}

func TestHolyShieldRequiresShield(t *testing.T) {
	req := historyTalentFixture("retribution", map[string]int{"holyShield": 1})
	p := req.Raid.Parties[0].Players[0]
	p.Equipment.Items[15] = &proto.ItemSpec{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	for _, id := range []int32{20925, 20927, 20928} {
		spell := sim.Raid.AllPlayerUnits[0].GetSpell(core.ActionID{SpellID: id})
		if spell == nil || spell.CanCast(sim, sim.Encounter.TargetUnits[0]) {
			t.Fatalf("Holy Shield %d available without a shield", id)
		}
	}
}

func TestWarlockRanksShareCooldowns(t *testing.T) {
	for _, ids := range [][]int32{
		{6353, 17924},
		{6789, 17925, 17926},
		{17877, 18867, 18868, 18869, 18870, 18871},
		{1293817, 1293818, 17962, 18930, 18931, 18932},
	} {
		req := historyTalentFixture("destruction", map[string]int{"conflagrate": 1, "shadowburn": 1})
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		unit := sim.Raid.AllPlayerUnits[0]
		target := sim.Encounter.TargetUnits[0]
		unit.GetSpell(core.ActionID{SpellID: 25309}).Dot(target).Apply(sim)
		highest := unit.GetSpell(core.ActionID{SpellID: ids[len(ids)-1]})
		if highest == nil || !highest.Cast(sim, target) {
			t.Fatalf("could not cast highest rank in %v", ids)
		}
		for _, id := range ids {
			spell := unit.GetSpell(core.ActionID{SpellID: id})
			if spell == nil || spell.CD.Timer != highest.CD.Timer || spell.CD.IsReady(sim) {
				t.Fatalf("rank %d bypassed its shared cooldown", id)
			}
		}
	}
}

func TestForeverDoesNotGrantSanctityAura(t *testing.T) {
	req := racialFixture("retribution", proto.Race_RaceHuman)
	req.Raid.Parties[0].Players[0].Rotation = &proto.APLRotation{}
	req.Raid.Parties[0].Players[0].GetRetributionPaladin().Options.Aura = proto.PaladinAura_SanctityAura
	req.Raid.Buffs.SanctityAura = true
	req.Raid.Parties[0].Buffs.SanctityAura = true
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	if aura := sim.Raid.AllPlayerUnits[0].GetAura("Sanctity Aura"); aura != nil {
		t.Fatal("unavailable Sanctity Aura registered from raid, party or personal settings")
	}
}

func TestDecimationDamageFollowsExecutePhase(t *testing.T) {
	req := historyTalentFixture("demonology", map[string]int{"decimation": 2})
	req.Raid.Parties[0].Players[0].ForeverTier1Bonuses = false
	req.Encounter.Duration = 10
	req.Encounter.ExecuteProportion_35 = .9
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	unit := sim.Raid.AllPlayerUnits[0]
	bolt := unit.GetSpell(core.ActionID{SpellID: 25307})
	pain := unit.GetSpell(core.ActionID{SpellID: 17923})
	soul := unit.GetSpell(core.ActionID{SpellID: 17924})
	proc := unit.GetAura("Decimation")
	baseBolt, basePain := bolt.DamageMultiplierAdditive, pain.DamageMultiplierAdditive
	baseCast := soul.CastTimeMultiplier
	proc.Activate(sim)
	if bolt.DamageMultiplierAdditive != baseBolt || pain.DamageMultiplierAdditive != basePain ||
		math.Abs(soul.CastTimeMultiplier-(baseCast-.4)) > 1e-9 {
		t.Fatal("Soul Fire proc incorrectly controlled the damage bonus")
	}
	proc.Deactivate(sim)
	for !sim.IsExecutePhase35() {
		if sim.Step() {
			t.Fatal("execute phase not reached")
		}
	}
	if proc.IsActive() {
		t.Fatal("fixture unexpectedly cast a triggering spell")
	}
	if math.Abs(bolt.DamageMultiplierAdditive-(baseBolt+.06)) > 1e-9 ||
		math.Abs(pain.DamageMultiplierAdditive-(basePain+.06)) > 1e-9 {
		t.Fatal("the first execute-range cast would miss its damage bonus")
	}
	proc.Activate(sim)
	proc.Deactivate(sim)
	if math.Abs(bolt.DamageMultiplierAdditive-(baseBolt+.06)) > 1e-9 ||
		math.Abs(pain.DamageMultiplierAdditive-(basePain+.06)) > 1e-9 ||
		math.Abs(soul.CastTimeMultiplier-baseCast) > 1e-9 {
		t.Fatal("proc expiry removed the conditional damage bonus or retained cast acceleration")
	}
}

// Isolate the named modifiers without spending unrelated filler points.
// These fixtures are not benchmark candidates.
func historyTalentFixture(key string, fields map[string]int) *proto.RaidSimRequest {
	for _, b := range builds() {
		if b.Key != key {
			continue
		}
		req := racialFixture(key, b.races()[0])
		p := req.Raid.Parties[0].Players[0]
		p.Rotation = &proto.APLRotation{}
		config := loadTalents(b)
		points, _ := config.decode("")
		index := 0
		for _, tree := range config.Trees {
			for _, talent := range tree.Talents {
				points[index] = fields[talent.Field]
				index++
			}
		}
		p.TalentsString = config.encode(points)
		return req
	}
	panic("unknown fixture build: " + key)
}
