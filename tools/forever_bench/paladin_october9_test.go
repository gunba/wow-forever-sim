//go:build with_db

package main

import (
	"math"
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/paladin"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

func paladinOctober9Fixture(t *testing.T, talents map[string]int, buffs *proto.RaidBuffs, extra ...*proto.Player) (*core.Simulation, *paladin.Paladin) {
	return paladinOctober9Configured(t, talents, buffs, nil, extra...)
}
func paladinOctober9Configured(t *testing.T, talents map[string]int, buffs *proto.RaidBuffs, configure func(*proto.Player), extra ...*proto.Player) (*core.Simulation, *paladin.Paladin) {
	t.Helper()
	for _, b := range builds() {
		if b.Key != "protection_paladin" {
			continue
		}
		p := b.player(proto.Race_RaceHuman)
		if configure != nil {
			configure(p)
		}
		cfg := loadTalents(b)
		points, _ := cfg.decode(p.TalentsString)
		offset := 0
		for _, tree := range cfg.Trees {
			for i, node := range tree.Talents {
				points[offset+i] = talents[node.Field]
			}
			offset += len(tree.Talents)
		}
		p.TalentsString = cfg.encode(points)
		p.ForeverTier1Bonuses = false
		p.Consumes = nil
		p.Buffs = nil
		p.Equipment = &proto.EquipmentSpec{Items: make([]*proto.ItemSpec, 17)}
		for i := range p.Equipment.Items {
			p.Equipment.Items[i] = &proto.ItemSpec{}
		}
		p.Equipment.Items[14].Id = 920000332
		p.Equipment.Items[15].Id = 920000384
		req := tankRequest(b, p, 1, 1793100001, 1, 300)
		req.Raid.Parties[0].Players = append(req.Raid.Parties[0].Players, extra...)
		req.Raid.Buffs = buffs
		req.Raid.Debuffs = nil
		req.Raid.Parties[0].Buffs = nil
		req.Encounter.Targets[0].Level = 60
		req.Encounter.Targets[0].Stats = nil
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		return sim, sim.Raid.Parties[0].Players[0].(paladin.PaladinAgent).GetPaladin()
	}
	t.Fatal("missing Protection build")
	return nil, nil
}
func TestPaladinOctober9ReckoningSharedGate(t *testing.T) {
	sim, p := paladinOctober9Fixture(t, map[string]int{"reckoning": 5}, nil)
	crit, block := p.GetAura("Reckoning Crit Trigger"), p.GetAura("Reckoning Block Trigger")
	if crit.Icd == nil || crit.Icd != block.Icd || crit.Icd.Duration != 1500*time.Millisecond {
		t.Fatal("Reckoning outcomes do not share1500ms gate")
	}
	incoming := p.CurrentTarget.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 1}, ProcMask: core.ProcMaskMeleeMHAuto, SpellSchool: core.SpellSchoolPhysical, DamageMultiplier: 1, ThreatMultiplier: 1})
	hit := &core.SpellResult{Target: &p.Unit, Outcome: core.OutcomeCrit, Damage: 1}
	crit.OnSpellHitTaken(crit, sim, incoming, hit)
	if crit.Icd.IsReady(sim) {
		t.Fatal("crit did not consume gate")
	}
	ready := crit.Icd.ReadyAt()
	for i := 0; i < 100; i++ {
		hit.Outcome = core.OutcomeBlock
		block.OnSpellHitTaken(block, sim, incoming, hit)
	}
	if crit.Icd.ReadyAt() != ready {
		t.Fatal("block bypassed/restarted crit gate")
	}
	sim.CurrentTime = 1500 * time.Millisecond
	for i := 0; i < 100 && block.Icd.IsReady(sim); i++ {
		block.OnSpellHitTaken(block, sim, incoming, hit)
	}
	if block.Icd.IsReady(sim) {
		t.Fatal("block never consumed ready shared gate")
	}
	if crit.Icd.ReadyAt() != 3*time.Second {
		t.Fatal("block did not lockcrit")
	}
	sim.Cleanup()
	sim.Reset()
	if !crit.Icd.IsReady(sim) {
		t.Fatal("ICD did not reset")
	}
}
func TestPaladinOctober9ShieldManaNoThreat(t *testing.T) {
	sim, p := paladinOctober9Fixture(t, map[string]int{"shieldSpecialization": 3}, nil)
	p.SpendMana(sim, p.CurrentMana()/2, p.NewManaMetrics(core.ActionID{SpellID: 2}))
	before := p.CurrentMana()
	a := p.GetAura("Shield Specialization Trigger")
	incoming := p.CurrentTarget.RegisterSpell(core.SpellConfig{ActionID: core.ActionID{SpellID: 3}, ProcMask: core.ProcMaskMeleeMHAuto})
	a.OnSpellHitTaken(a, sim, incoming, &core.SpellResult{Target: &p.Unit, Outcome: core.OutcomeBlock, Damage: 1})
	if math.Abs(p.CurrentMana()-before-.06*p.MaxMana()) > 1e-8 {
		t.Fatal("mana amount changed")
	}
	sim.Cleanup()
	manaGain := p.GetSpell(core.ActionID{OtherID: proto.OtherAction_OtherActionManaGain})
	for _, m := range manaGain.SpellMetrics {
		if m.TotalThreat != 0 {
			t.Fatal("shield mana generated threat")
		}
	}
}
func TestPaladinOctober9ReflectionsExternalAndOwned(t *testing.T) {
	for _, owned := range []bool{false, true} {
		buffs := &proto.RaidBuffs{RetributionAura: proto.TristateEffect_TristateEffectImproved, RetributionAuraProviderSpellPower: 100}
		if owned {
			buffs.RetributionAuraProvider = &proto.UnitReference{Type: proto.UnitReference_Self}
		}
		sim, p := paladinOctober9Fixture(t, nil, buffs)
		p.AddStatsDynamic(sim, stats.Stats{stats.SpellPower: 200, stats.SpellDamage: 30, stats.HolyPower: 70, stats.SpellCrit: 10000})
		s := p.GetSpell(core.ActionID{SpellID: 10301})
		s.BonusHitRating = 10000
		before := s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
		s.ApplyEffects(sim, p.CurrentTarget, s)
		want := 39.0
		if owned {
			want = (20 + .06*(p.GetStat(stats.SpellPower)+p.GetStat(stats.SpellDamage)+p.GetStat(stats.HolyPower))) * 1.5
		}
		got := s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage - before
		if math.Abs(got-want) > 1e-8 {
			t.Fatalf("owned%v damage%g expected%g", owned, got, want)
		}
		if s.SpellMetrics[p.CurrentTarget.UnitIndex].Crits != 0 || s.ProcMask != core.ProcMaskEmpty {
			t.Fatal("reflection crit/proc eligibility changed")
		}
		p.AddStatDynamic(sim, stats.SpellPower, 100)
		before = s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage
		s.ApplyEffects(sim, p.CurrentTarget, s)
		if owned {
			want += 9
		}
		if math.Abs(s.SpellMetrics[p.CurrentTarget.UnitIndex].TotalDamage-before-want) > 1e-8 {
			t.Fatal("provider power notdynamic or receiver power leaked")
		}
	}
}
func TestPaladinOctober9ProviderProtoRoundTrip(t *testing.T) {
	b := &proto.RaidBuffs{Thorns: proto.TristateEffect_TristateEffectRegular, ThornsProvider: &proto.UnitReference{Type: proto.UnitReference_Player, Index: 7}, ThornsProviderSpellPower: 321, RetributionAuraProviderSpellPower: 456}
	json, err := protojson.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	back := &proto.RaidBuffs{}
	if err = protojson.Unmarshal(json, back); err != nil || !googleProto.Equal(b, back) {
		t.Fatal("provider JSON roundtrip", err)
	}
	raw, err := googleProto.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if err = googleProto.Unmarshal(raw, back); err != nil || !googleProto.Equal(b, back) {
		t.Fatal("provider binary roundtrip", err)
	}
}
