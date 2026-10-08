//go:build with_db

package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/simsignals"
	"github.com/wowsims/classic/sim/core/stats"
	"github.com/wowsims/classic/sim/rogue"
	"google.golang.org/protobuf/encoding/protojson"
	googleProto "google.golang.org/protobuf/proto"
)

func roguePoisonRequest(t *testing.T, options string) *proto.RaidSimRequest {
	t.Helper()
	req := racialFixture("mutilate", proto.Race_RaceOrc)
	p := req.Raid.Parties[0].Players[0]
	p.TalentsString = ""
	p.Rotation = &proto.APLRotation{}
	p.ForeverTier1Bonuses = false
	p.GetRogue().Options = &proto.RogueOptions{}
	if err := protojson.Unmarshal([]byte(options), p.GetRogue().Options); err != nil {
		t.Fatalf("Rogue poison options cannot be imported: %v", err)
	}
	p.Consumes = &proto.Consumes{}
	req.SimOptions.IsTest = true
	req.SimOptions.RandomSeed = 1
	return req
}

func TestRoguePoisonAndTemporaryEnchantCoexist(t *testing.T) {
	req := roguePoisonRequest(t, `{"mainHandPoison":"InstantPoison","offHandPoison":"DeadlyPoison"}`)
	baseline := core.NewSim(req, simsignals.Signals{})
	base := baseline.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
	p := req.Raid.Parties[0].Players[0]
	p.Consumes.MainHandImbue = proto.WeaponImbue_DenseSharpeningStone
	p.Consumes.OffHandImbue = proto.WeaponImbue_BrilliantWizardOil
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	r := sim.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
	if r.AutoAttacks.MH().BaseDamageMin-base.AutoAttacks.MH().BaseDamageMin != 8 ||
		r.AutoAttacks.MH().BaseDamageMax-base.AutoAttacks.MH().BaseDamageMax != 8 {
		t.Fatal("poison suppressed the sharpening stone's weapon damage")
	}
	if r.GetStat(stats.SpellPower)-base.GetStat(stats.SpellPower) != 36 {
		t.Fatal("poison suppressed Wizard Oil's spell power")
	}
	if r.MaxEnergy() != base.MaxEnergy() || r.CurrentEnergy() != 100 || r.ComboPoints() != 0 || r.HasManaBar() {
		t.Fatal("separate poison/enchant inputs changed Rogue resources")
	}
	instant, deadly := r.GetAura("Instant Poison"), r.GetAura("Deadly Poison")
	if instant == nil || deadly == nil {
		t.Fatal("temporary enchants suppressed poison proc auras")
	}
	hit := &core.SpellResult{Target: r.CurrentTarget, Outcome: core.OutcomeHit}
	instant.OnSpellHitDealt(instant, sim, r.AutoAttacks.OHAuto(), hit)
	if r.InstantPoison.SpellMetrics[r.CurrentTarget.UnitIndex].Casts != 0 {
		t.Fatal("main-hand poison procced from the off hand")
	}
	var wantInstant, wantDeadly int32
	for i := 0; i < 64; i++ {
		if baseline.RandomFloat("Instant Poison") < r.GetInstantPoisonProcChance() {
			wantInstant++
		}
		if baseline.RandomFloat("Deadly Poison") < r.GetDeadlyPoisonProcChance() {
			wantDeadly++
		}
		instant.OnSpellHitDealt(instant, sim, r.AutoAttacks.MHAuto(), hit)
		deadly.OnSpellHitDealt(deadly, sim, r.AutoAttacks.OHAuto(), hit)
	}
	if wantInstant == 0 || wantDeadly == 0 || r.InstantPoison.SpellMetrics[r.CurrentTarget.UnitIndex].Casts != wantInstant || r.DeadlyPoison.SpellMetrics[r.CurrentTarget.UnitIndex].Casts != wantDeadly {
		t.Fatal("poisons did not retain their hand eligibility and proc chances alongside temporary enchants")
	}
}

func TestRoguePoisonInputValidation(t *testing.T) {
	cases := []struct {
		name, options, errorText string
		configure                func(*proto.Player)
	}{
		{"legacy instant", `{}`, "migrate this profile", func(p *proto.Player) { p.Consumes.MainHandImbue = proto.WeaponImbue_InstantPoison }},
		{"legacy deadly", `{}`, "migrate this profile", func(p *proto.Player) { p.Consumes.OffHandImbue = proto.WeaponImbue_DeadlyPoison }},
		{"legacy wound", `{}`, "migrate this profile", func(p *proto.Player) { p.Consumes.MainHandImbue = proto.WeaponImbue_WoundPoison }},
		{"empty main hand", `{"mainHandPoison":"InstantPoison"}`, "requires a weapon", func(p *proto.Player) { p.Equipment = &proto.EquipmentSpec{} }},
		{"empty off hand", `{"offHandPoison":"WoundPoison"}`, "requires a weapon", func(p *proto.Player) { p.Equipment.Items[15] = &proto.ItemSpec{} }},
		{"unknown poison", `{"mainHandPoison":99}`, "Invalid Rogue poison", func(*proto.Player) {}},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req := roguePoisonRequest(t, test.options)
			test.configure(req.Raid.Parties[0].Players[0])
			defer func() {
				if recovered := recover(); recovered == nil || !strings.Contains(fmt.Sprint(recovered), test.errorText) {
					t.Fatalf("invalid input rejection: %v, want %q", recovered, test.errorText)
				}
			}()
			core.NewSim(req, simsignals.Signals{})
		})
	}
}

func TestRoguePoisonClassicConsumesRemainValid(t *testing.T) {
	req := roguePoisonRequest(t, `{}`)
	req.SimOptions.Ruleset = proto.Ruleset_RulesetClassic
	p := req.Raid.Parties[0].Players[0]
	p.Consumes.MainHandImbue, p.Consumes.OffHandImbue = proto.WeaponImbue_InstantPoison, proto.WeaponImbue_DeadlyPoison
	sim := core.NewSim(req, simsignals.Signals{})
	r := sim.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
	if r.GetAura("Instant Poison") == nil || r.GetAura("Deadly Poison") == nil {
		t.Fatal("Classic poison consumes stopped applying poisons")
	}
}

func TestRoguePoisonRequestRoundTrip(t *testing.T) {
	req := roguePoisonRequest(t, `{"mainHandPoison":"WoundPoison","offHandPoison":"InstantPoison"}`)
	req.Raid.Parties[0].Players[0].Consumes.MainHandImbue = proto.WeaponImbue_ShadowOil
	data, err := protojson.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var restored proto.RaidSimRequest
	if err := protojson.Unmarshal(data, &restored); err != nil {
		t.Fatal(err)
	}
	if !googleProto.Equal(req, &restored) {
		t.Fatal("poison and temporary enchant inputs changed during request round trip")
	}
}
