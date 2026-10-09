//go:build with_db

package main

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
)

func revelationOctober9Fixture(key string, model *proto.ForeverRevelationModel, equipped bool, edit func(*proto.Player), settings ...func(*proto.RaidSimRequest)) (*core.Simulation, *core.Character, *proto.Player) {
	fields := map[string]int{}
	if key == "smite" {
		fields["penance"] = 1
	}
	req := historyTalentFixture(key, fields)
	p := req.Raid.Parties[0].Players[0]
	p.ForeverRevelationModel = model
	p.Equipment.Items[14].Enchant = 0
	if equipped {
		p.Equipment.Items[14].Enchant = 8217
	}
	if edit != nil {
		edit(p)
	}
	for _, setting := range settings {
		setting(req)
	}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	for _, u := range sim.Environment.AllUnits {
		u.AutoAttacks.CancelAutoSwing(sim)
		if u.Type == core.PlayerUnit {
			u.CancelGCDTimer(sim)
		}
	}
	return sim, sim.Raid.Parties[0].Players[0].GetCharacter(), p
}

func revelationOctober9Spell(t *testing.T, c *core.Character, id int32, tag int32) *core.Spell {
	t.Helper()
	s := c.GetSpell(core.ActionID{SpellID: id, Tag: tag})
	if s == nil {
		t.Fatalf("missing real spell %d/%d", id, tag)
	}
	return s
}

func revelationOctober9Activate(sim *core.Simulation, c *core.Character) *core.Aura {
	a := c.GetAura("Revelation")
	a.Activate(sim)
	a.SetStacks(sim, 1)
	return a
}

func TestOctober9RevelationCatalogAndHandler(t *testing.T) {
	if _, ok := core.EnchantsByEffectID[8217]; !ok {
		t.Fatal("source Revelation8217 missing from native catalog")
	}
	if !core.HasEnchantEffect(8217) {
		t.Fatal("equipped Revelation8217 has no handler/model validation")
	}
	for id, minimum := range map[int32]int32{8483: 0, 8486: 15, 8488: 25, 8491: 35} {
		e := core.EnchantsByEffectID[id]
		if e.RequiredLevel != 0 || e.ItemLevelMin != minimum {
			t.Fatalf("kit%d wearer/use-level confusion: %+v", id, e)
		}
		for _, slot := range []proto.ItemSlot{proto.ItemSlot_ItemSlotChest, proto.ItemSlot_ItemSlotLegs, proto.ItemSlot_ItemSlotHands, proto.ItemSlot_ItemSlotFeet} {
			item := core.Item{Type: proto.ItemType_ItemTypeChest, ArmorType: proto.ArmorType_ArmorTypeLeather, ItemLevel: minimum, Enchant: e}
			if err := core.ValidateEnchantRequirements(1, slot, item); err != nil {
				t.Fatalf("kit wearer minimum is source0: %v", err)
			}
			item.ItemLevel = minimum - 1
			if minimum > 0 && core.ValidateEnchantRequirements(60, slot, item) == nil {
				t.Fatal("kit item minimum ignored")
			}
		}
	}
}

func TestOctober9RevelationExplicitModelAndUnsupportedPayload(t *testing.T) {
	cases := []struct {
		key   string
		model *proto.ForeverRevelationModel
		want  string
	}{
		{"frost", nil, "explicitly enabled"}, {"frost", &proto.ForeverRevelationModel{}, "explicitly enabled"},
		{"frost", &proto.ForeverRevelationModel{Enabled: true, BaseChance: -.01}, "baseChance"},
		{"frost", &proto.ForeverRevelationModel{Enabled: true, BaseChance: 1.01}, "baseChance"},
		{"frost", &proto.ForeverRevelationModel{Enabled: true, BaseChance: math.NaN()}, "baseChance"},
		{"frost", &proto.ForeverRevelationModel{Enabled: true, BaseChance: math.Inf(1)}, "baseChance"},
		{"frost", &proto.ForeverRevelationModel{Enabled: true, CritExponent: -1}, "critExponent"},
		{"frost", &proto.ForeverRevelationModel{Enabled: true, CritExponent: math.Inf(1)}, "critExponent"},
		{"fury", &proto.ForeverRevelationModel{Enabled: true}, "unsupported source class payload"},
		{"combat", &proto.ForeverRevelationModel{Enabled: true}, "unsupported source class payload"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil || !strings.Contains(fmt.Sprint(r), tc.want) {
					t.Fatalf("error %v, want %s", r, tc.want)
				}
			}()
			revelationOctober9Fixture(tc.key, tc.model, true, nil)
		})
	}
	m := &proto.ForeverRevelationModel{Enabled: true, BaseChance: 0, CritExponent: 0}
	_, c, p := revelationOctober9Fixture("frost", m, true, nil)
	p.ForeverRevelationModel.BaseChance = 1
	if c.ForeverRevelationModel.BaseChance != 0 {
		t.Fatal("model pointer aliases mutable input")
	}
	// Saving controls on an unsupported class is legal until8217 is equipped.
	revelationOctober9Fixture("fury", &proto.ForeverRevelationModel{Enabled: true}, false, nil)
}

func TestOctober9RevelationConditionalCurveAndCandidateReasons(t *testing.T) {
	m := &proto.ForeverRevelationModel{Enabled: true, BaseChance: .8, CritExponent: 2}
	for _, tc := range []struct{ crit, want float64 }{{-1, .8}, {0, .8}, {.25, .45}, {.5, .2}, {1, 0}, {2, 0}} {
		if got := common.ForeverRevelationProcChance(m, tc.crit); math.Abs(got-tc.want) > 1e-12 {
			t.Fatalf("conditional probability(%g)=%g want%g", tc.crit, got, tc.want)
		}
	}
	m.CritExponent = 0
	for _, crit := range []float64{0, .5, 1} {
		if common.ForeverRevelationProcChance(m, crit) != .8 {
			t.Fatal("exponent0 is not flat CONDITIONAL chance")
		}
	}
	for _, key := range []string{"frost", "fury"} {
		_, _, p := revelationOctober9Fixture(key, nil, false, nil)
		e := &proto.UIEnchant{EffectId: 8217, Type: proto.ItemType_ItemTypeWeapon}
		if !enchantFits(p, core.ItemsByID[p.Equipment.Items[14].Id], e) {
			t.Fatal("unsupported model fabricated item illegality")
		}
		reason := excludedEnchantCandidates(p, 14)[8217]
		want := "explicitly enabled"
		if key == "fury" {
			want = "unsupported source class payload"
		}
		if !strings.Contains(reason, want) {
			t.Fatalf("candidate missing qualified reason %q", reason)
		}
		if slices.ContainsFunc(legalEnchants(p, 14), func(e *proto.UIEnchant) bool { return e.EffectId == 8217 }) {
			t.Fatal("unconfigured candidate admitted")
		}
		p.Equipment.Items[14].Enchant = 8217
		if err := validateGearEnchants(p); err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("optimizer validation %v", err)
		}
		data, _ := json.Marshal(gearSearchReport{ExcludedEnchants: excludedEnchantsForPlayer(p)})
		if !strings.Contains(string(data), "ExcludedEnchants") || !strings.Contains(string(data), want) {
			t.Fatal("excluded enchant trace not serialized separately")
		}
		p.ForeverRevelationModel = &proto.ForeverRevelationModel{Enabled: true}
		if key == "frost" && !slices.ContainsFunc(legalEnchants(p, 14), func(e *proto.UIEnchant) bool { return e.EffectId == 8217 }) {
			t.Fatal("explicit configured zero chance excluded")
		}
	}
}

func TestOctober9RevelationSourcePayloadsAndDirectQualification(t *testing.T) {
	for _, tc := range []struct {
		key                         string
		payload, eligible, excluded int32
	}{
		{"frost", 1248808, 25304, 25306}, {"smite", 1323377, 10934, 10894},
		{"destruction", 1323392, 25307, 25309}, {"feral", 1323400, 9912, 9835},
		{"survival", 1323410, 14287, 2643}, {"elemental", 1323418, 15208, 29228},
		{"retribution", 1323419, 20286, 20924},
	} {
		t.Run(tc.key, func(t *testing.T) {
			sim, c, _ := revelationOctober9Fixture(tc.key, &proto.ForeverRevelationModel{Enabled: true, BaseChance: 1}, true, nil)
			buff, driver := c.GetAura("Revelation"), c.GetAura("Revelation Driver")
			if buff.ActionID.SpellID != tc.payload || buff.Duration != 15*time.Second || buff.MaxStacks != 1 || driver.ActionID.SpellID != 1248806 {
				t.Fatal("source payload metadata mismatch")
			}
			s := revelationOctober9Spell(t, c, tc.eligible, 0)
			bad := revelationOctober9Spell(t, c, tc.excluded, 0)
			fire := func(spell *core.Spell, result *core.SpellResult) { driver.OnSpellHitDealt(driver, sim, spell, result) }
			r := &core.SpellResult{Target: c.CurrentTarget, Outcome: core.OutcomeHit, DirectDamageOrHealing: true, EffectiveCritChance: .2}
			fire(bad, r)
			if buff.IsActive() {
				t.Fatal("unmasked spell proc")
			}
			fire(s, &core.SpellResult{Target: c.CurrentTarget, Outcome: core.OutcomeHit})
			fire(s, &core.SpellResult{Target: c.CurrentTarget, Outcome: core.OutcomeCrit, DirectDamageOrHealing: true})
			fire(s, &core.SpellResult{Target: c.CurrentTarget, Outcome: core.OutcomeMiss, DirectDamageOrHealing: true})
			if buff.IsActive() {
				t.Fatal("container/critical/miss proc")
			}
			before := c.GetStats()
			fire(s, r)
			if !buff.IsActive() || buff.GetStacks() != 1 || s.BonusDirectCritRating != 100*core.SpellCritRatingPerCritChance || bad.BonusDirectCritRating != 0 || c.GetStats() != before {
				t.Fatal("source family bonus qualification/global stat bleed")
			}
			buff.Deactivate(sim)
			c.ForeverRevelationModel.TriggerOnMiss = true
			fire(s, &core.SpellResult{Target: c.CurrentTarget, Outcome: core.OutcomeMiss, DirectDamageOrHealing: true})
			if !buff.IsActive() {
				t.Fatal("explicit miss trigger policy ignored")
			}
			buff.Deactivate(sim)
			c.ForeverRevelationModel.BaseChance = 0
			fire(s, r)
			if buff.IsActive() {
				t.Fatal("zero chance replaced by hidden fallback")
			}
		})
	}
}

func TestOctober9RevelationChargeSamplingPeriodicAndEffectiveCrit(t *testing.T) {
	sim, c, _ := revelationOctober9Fixture("frost", &proto.ForeverRevelationModel{Enabled: true}, true, nil)
	s := revelationOctober9Spell(t, c, 25304, 0)
	target := c.CurrentTarget
	s.BonusCritRating += (.25 - s.SpellCritChance(target)) * core.SpellCritRatingPerCritChance * 100
	buff := revelationOctober9Activate(sim, c)
	baseCrit := s.SpellCritChance(target)
	if math.Abs(baseCrit-.25) > 1e-12 {
		t.Fatal("baseline crit setup failed")
	}
	// Outcome-only and non-crit-capable direct calculations never spend charges.
	for _, calc := range []func() *core.SpellResult{
		func() *core.SpellResult { return s.CalcOutcome(sim, target, s.OutcomeMagicCrit) },
		func() *core.SpellResult { return s.CalcDamage(sim, target, 10, s.OutcomeAlwaysHit) },
		func() *core.SpellResult { return s.CalcDamage(sim, target, 10, s.OutcomeAlwaysMiss) },
		func() *core.SpellResult { return s.CalcDamage(sim, target, 10, s.OutcomeExpectedMagicHitAndCrit) },
	} {
		r := calc()
		if r.DirectCritBonusApplied || !buff.IsActive() {
			t.Fatal("non-crit/miss/container/expected query consumed charge")
		}
		s.DisposeResult(r)
	}
	r := s.CalcDamage(sim, target, 10, s.OutcomeMagicCrit)
	if !r.DidCrit() || !r.DirectCritBonusApplied || math.Abs(r.EffectiveCritChance-baseCrit) > 1e-12 || buff.IsActive() {
		t.Fatal("direct sample did not reserve once or captured boosted crit")
	}
	s.DisposeResult(r)
	if s.BonusDirectCritRating != 0 {
		t.Fatal("one charge removal left crit bonus")
	}
	// Fresh OnHit proc cannot consume its own charge in that same event.
	c.ForeverRevelationModel.BaseChance = 1
	r = s.CalcDamage(sim, target, 10, s.OutcomeAlwaysHit)
	s.DealDamage(sim, r)
	if !buff.IsActive() || r.DirectCritBonusApplied {
		t.Fatal("fresh hit proc self-consumed")
	}
	c.ForeverRevelationModel.BaseChance = 0
	// Flamestrike's initial direct impact qualifies, its ordinary DoT does not.
	dotSpell := revelationOctober9Spell(t, c, 10216, 0)
	dot := dotSpell.DotOrAOEDot(target)
	dot.Snapshot(target, 10, false)
	if dot.SnapshotCritChance >= 1 {
		t.Fatal("direct crit bonus leaked into snapshot")
	}
	r = dot.CalcSnapshotDamage(sim, target, dot.OutcomeTick)
	if r.DirectDamageOrHealing || r.DirectCritBonusApplied || !buff.IsActive() {
		t.Fatal("ordinary DoT used charge")
	}
	dotSpell.DisposeResult(r)
	wand := c.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 116, Tag: 98}, SpellSchool: core.SpellSchoolFrost, DefenseType: core.DefenseTypeRanged, ProcMask: core.ProcMaskRangedAuto})
	if wand.BonusDirectCritRating != 0 {
		t.Fatal("wand alias gained bonus")
	}
	r = wand.CalcDamage(sim, target, 10, wand.OutcomeRangedCritOnly)
	if r.DirectCritBonusApplied || !buff.IsActive() {
		t.Fatal("wand used charge")
	}
	wand.DisposeResult(r)
	// Direct healing qualifies, periodic healing and healing snapshots do not.
	heal := c.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 116, Tag: 97}, SpellSchool: core.SpellSchoolFrost, ProcMask: core.ProcMaskSpellHealing, Flags: core.SpellFlagHelpful})
	hot := &core.Dot{Spell: heal, DamageMultiplier: 1}
	hot.SnapshotHeal(&c.Unit, 1, false)
	r = hot.CalcSnapshotHealing(sim, &c.Unit, heal.OutcomeHealingCrit)
	if r.DirectDamageOrHealing || r.DirectCritBonusApplied || !buff.IsActive() {
		t.Fatal("healing snapshot used direct bonus")
	}
	heal.DisposeResult(r)
	heal.CalcAndDealPeriodicHealing(sim, &c.Unit, 1, heal.OutcomeHealingCrit)
	if !buff.IsActive() {
		t.Fatal("periodic healing used direct bonus")
	}
	r = heal.CalcHealing(sim, &c.Unit, 1, heal.OutcomeHealingCrit)
	if !r.DirectCritBonusApplied || buff.IsActive() {
		t.Fatal("direct healing did not use charge")
	}
	heal.DisposeResult(r)
}

func TestOctober9RevelationArcaneMissilesTravelReservation(t *testing.T) {
	sim, c, _ := revelationOctober9Fixture("arcane", &proto.ForeverRevelationModel{Enabled: true}, true, nil)
	child := revelationOctober9Spell(t, c, 25345, 1)
	parent := revelationOctober9Spell(t, c, 25345, 0)
	child.BonusHitRating = 10000
	child.BonusCritRating -= child.SpellCritChance(c.CurrentTarget) * core.SpellCritRatingPerCritChance * 100
	buff := revelationOctober9Activate(sim, c)
	parent.CalcAndDealOutcome(sim, c.CurrentTarget, parent.OutcomeMagicHit)
	if !buff.IsActive() {
		t.Fatal("AM container used charge")
	}
	child.ApplyEffects(sim, c.CurrentTarget, child)
	if buff.IsActive() {
		t.Fatal("real AMtag1 did not reserve charge before impact")
	}
	child.ApplyEffects(sim, c.CurrentTarget, child)
	if child.SpellMetrics[c.CurrentTarget.UnitIndex].Crits != 1 {
		t.Fatal("concurrent missiles shared a charge")
	}
	if child.SpellMetrics[c.CurrentTarget.UnitIndex].TotalDamage != 0 {
		t.Fatal("missile travel was bypassed")
	}
	// A newly created charge while the old projectile is in flight cannot be
	// consumed again when that already-sampled critical projectile arrives.
	revelationOctober9Activate(sim, c)
	utilityAdvance(sim, 2*time.Second)
	if child.SpellMetrics[c.CurrentTarget.UnitIndex].TotalDamage <= 0 || !buff.IsActive() {
		t.Fatal("travel delivery consumed a newer charge")
	}
}

func TestOctober9RevelationLogicalChannelChildren(t *testing.T) {
	for _, tc := range []struct {
		key           string
		parent, child int32
		tick          time.Duration
	}{
		{"frost", 10187, 1279949, time.Second}, {"smite", 1316995, 1316993, time.Second},
		{"balance", 17402, 1278759, time.Second}, {"survival", 14295, 1279715, time.Second},
		{"destruction", 11678, 1282385, 2 * time.Second},
	} {
		t.Run(tc.key, func(t *testing.T) {
			sim, c, _ := revelationOctober9Fixture(tc.key, &proto.ForeverRevelationModel{Enabled: true}, true, nil)
			s := revelationOctober9Spell(t, c, tc.parent, 0)
			if s.RevelationPeriodicDirectChildID != tc.child {
				t.Fatalf("logical child %d want%d", s.RevelationPeriodicDirectChildID, tc.child)
			}
			dot := s.DotOrAOEDot(c.CurrentTarget)
			if dot.TickLength != tc.tick {
				t.Fatalf("timing changed: %s", dot.TickLength)
			}
			s.BonusCritRating -= s.SpellCritChance(c.CurrentTarget) * core.SpellCritRatingPerCritChance * 100
			buff := revelationOctober9Activate(sim, c)
			r := s.CalcOutcome(sim, c.CurrentTarget, s.OutcomeMagicHit)
			if r.DirectCritBonusApplied || !buff.IsActive() {
				t.Fatal("logical channel container consumed")
			}
			s.DisposeResult(r)
			dot.OnSnapshot(sim, c.CurrentTarget, dot, false)
			r = dot.CalcSnapshotDamage(sim, c.CurrentTarget, dot.OutcomeTick)
			if !r.DirectDamageOrHealing || !r.DidCrit() || !r.DirectCritBonusApplied || buff.IsActive() || math.Abs(r.EffectiveCritChance) > 1e-12 {
				t.Fatalf("logical child result %+v buff%v", r, buff.IsActive())
			}
			s.DealPeriodicDamage(sim, r)
			c.ForeverRevelationModel.BaseChance = 1
			r = dot.CalcSnapshotDamage(sim, c.CurrentTarget, dot.OutcomeTick)
			s.DealPeriodicDamage(sim, r)
			if !buff.IsActive() || r.DirectCritBonusApplied {
				t.Fatal("logical child's fresh periodic callback proc self-consumed")
			}
			// Same seed and gear: an explicitly zero model with no active bonus
			// must retain the original channel's numeric and timing behavior.
			baseSim, base, _ := revelationOctober9Fixture(tc.key, nil, false, nil)
			zeroSim, zero, _ := revelationOctober9Fixture(tc.key, &proto.ForeverRevelationModel{Enabled: true}, true, nil)
			baseS, zeroS := revelationOctober9Spell(t, base, tc.parent, 0), revelationOctober9Spell(t, zero, tc.parent, 0)
			baseDot, zeroDot := baseS.DotOrAOEDot(base.CurrentTarget), zeroS.DotOrAOEDot(zero.CurrentTarget)
			baseDot.OnSnapshot(baseSim, base.CurrentTarget, baseDot, false)
			zeroDot.OnSnapshot(zeroSim, zero.CurrentTarget, zeroDot, false)
			for i := 0; i < 3; i++ {
				a := baseDot.CalcSnapshotDamage(baseSim, base.CurrentTarget, baseDot.OutcomeTick)
				b := zeroDot.CalcSnapshotDamage(zeroSim, zero.CurrentTarget, zeroDot.OutcomeTick)
				if a.Damage != b.Damage || a.Outcome != b.Outcome || baseDot.TickLength != zeroDot.TickLength {
					t.Fatal("inactive alias changed ordinary channel simulation")
				}
				baseS.DisposeResult(a)
				zeroS.DisposeResult(b)
			}
		})
	}
}

func TestOctober9RevelationAoEEventScopeAndClassicRetention(t *testing.T) {
	sim, c, _ := revelationOctober9Fixture("frost", &proto.ForeverRevelationModel{Enabled: true}, true, nil, func(req *proto.RaidSimRequest) {
		req.Encounter.Targets = append(req.Encounter.Targets, req.Encounter.Targets[0])
	})
	s := revelationOctober9Spell(t, c, 10202, 0)
	s.BonusHitRating = 10000
	s.BonusCritRating -= s.SpellCritChance(c.CurrentTarget) * core.SpellCritRatingPerCritChance * 100
	buff := revelationOctober9Activate(sim, c)
	s.ApplyEffects(sim, c.CurrentTarget, s)
	targets := sim.Encounter.TargetUnits
	if buff.IsActive() || s.SpellMetrics[targets[0].UnitIndex].Crits != 1 || s.SpellMetrics[targets[1].UnitIndex].Crits != 0 {
		t.Fatal("event model did not reserve for first calculated eligible AoE target only")
	}
	_, classic, _ := revelationOctober9Fixture("frost", nil, true, nil, func(req *proto.RaidSimRequest) {
		req.SimOptions.Ruleset = proto.Ruleset_RulesetClassic
	})
	if classic.GetAura("Revelation") != nil || classic.GetAura("Revelation Driver") != nil {
		t.Fatal("Forever-only enchant required a model or granted a Classic buff")
	}
	for _, spell := range classic.Spellbook {
		if spell.BonusDirectCritRating != 0 || spell.RevelationPeriodicDirectChildID != 0 {
			t.Fatal("Revelation direct-only hooks leaked into Classic")
		}
	}
}

func TestOctober9RevelationPhysicalShotExpiryFormsAndSwap(t *testing.T) {
	sim, c, _ := revelationOctober9Fixture("survival", &proto.ForeverRevelationModel{Enabled: true}, true, nil)
	s := revelationOctober9Spell(t, c, 14287, 0)
	table := c.AttackTables[c.CurrentTarget.UnitIndex][s.CastType]
	s.BonusCritRating -= s.PhysicalCritChance(table) * core.CritRatingPerCritChance * 100
	buff := revelationOctober9Activate(sim, c)
	r := s.CalcDamage(sim, c.CurrentTarget, 10, s.OutcomeRangedCritOnly)
	if !r.DidCrit() || !r.DirectCritBonusApplied || math.Abs(r.EffectiveCritChance) > 1e-12 || buff.IsActive() {
		t.Fatal("eligible Hunter physical crit hook failed")
	}
	s.DisposeResult(r)
	for _, key := range []string{"feral", "feral_tank_druid"} {
		t.Run(key, func(t *testing.T) {
			sim, c, _ := revelationOctober9Fixture(key, &proto.ForeverRevelationModel{Enabled: true, BaseChance: 1}, true, func(p *proto.Player) {
				p.EnableItemSwap = true
				p.ItemSwap = &proto.ItemSwap{MhItem: &proto.ItemSpec{Id: p.Equipment.Items[14].Id, Enchant: 0}}
			})
			buff := revelationOctober9Activate(sim, c)
			if c.ActiveShapeShift == nil || !c.ActiveShapeShift.IsActive() {
				t.Fatal("form fixture missing")
			}
			// These specs register their own form only. Shift to humanoid and
			// back through the actual form aura lifecycle, retaining the payload.
			originalForm := c.ActiveShapeShift
			originalForm.Deactivate(sim)
			if !buff.IsActive() {
				t.Fatal("leaving form removed Revelation")
			}
			originalForm.Activate(sim)
			if !buff.IsActive() || c.ActiveShapeShift != originalForm {
				t.Fatal("entering form removed Revelation")
			}
			form, mana := c.ActiveShapeShift, c.CurrentMana()
			wrath := c.GetSpell(core.ActionID{SpellID: 9912})
			if wrath == nil {
				// Tank currently registers no offensive caster spells. A sourced
				// late registration must still inherit the active class payload.
				wrath = c.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 9912}, SpellSchool: core.SpellSchoolNature, DefenseType: core.DefenseTypeMagic, ProcMask: core.ProcMaskSpellDamage})
			}
			before := wrath.BonusDirectCritRating
			sim.CurrentTime = 5 * time.Second
			driver := c.GetAura("Revelation Driver")
			driver.OnSpellHitDealt(driver, sim, wrath, &core.SpellResult{Target: c.CurrentTarget, Outcome: core.OutcomeHit, DirectDamageOrHealing: true})
			if buff.ExpiresAt() != 20*time.Second || wrath.BonusDirectCritRating != before || buff.GetStacks() != 1 || c.ActiveShapeShift != form || c.CurrentMana() != mana {
				t.Fatal("refresh stacked bonus or broke form/mana")
			}
			c.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand})
			if buff.IsActive() || wrath.BonusDirectCritRating != 0 {
				t.Fatal("unequip kept charge")
			}
			c.ItemSwap.SwapItems(sim, []proto.ItemSlot{proto.ItemSlot_ItemSlotMainHand})
			buff = revelationOctober9Activate(sim, c)
			utilityAdvance(sim, 21*time.Second)
			if buff.IsActive() || wrath.BonusDirectCritRating != 0 {
				t.Fatal("15sec expiry left bonus")
			}
			revelationOctober9Activate(sim, c)
			sim.Cleanup()
			sim.Reset()
			if buff.IsActive() || wrath.BonusDirectCritRating != 0 {
				t.Fatal("iteration reset retained temporary charge")
			}
		})
	}
}
