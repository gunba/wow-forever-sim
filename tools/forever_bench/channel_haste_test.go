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
	"github.com/wowsims/classic/sim/hunter"
	"github.com/wowsims/classic/sim/mage"
	"github.com/wowsims/classic/sim/priest"
	"github.com/wowsims/classic/sim/rogue"
	"github.com/wowsims/classic/sim/shaman"
)

func channelHasteFixture(buildKey string, ruleset proto.Ruleset, haste float64) *core.Simulation {
	req := racialFixture(buildKey, proto.Race_RaceUndead)
	player := req.Raid.Parties[0].Players[0]
	player.Equipment = &proto.EquipmentSpec{}
	player.ForeverTier1Bonuses = false
	player.Rotation = &proto.APLRotation{}
	player.BonusStats = &proto.UnitStats{Stats: stats.Stats{
		stats.SpellHaste: haste * core.HasteRatingPerHastePercent,
		stats.SpellHit:   100,
	}.ToFloatArray()}
	req.SimOptions.Ruleset = ruleset
	req.Encounter.Targets[0].Level = 60
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	return sim
}

func TestForeverMindFlayChannelAndSpellGCDScaleWithHaste(t *testing.T) {
	for _, tc := range []struct {
		name    string
		ruleset proto.Ruleset
		haste   float64
		want    time.Duration
		gcd     time.Duration
	}{
		{"forever_no_haste", proto.Ruleset_RulesetForever, 0, 3 * time.Second, 1500 * time.Millisecond},
		{"forever_100_percent", proto.Ruleset_RulesetForever, 100, 1500 * time.Millisecond, time.Second},
		{"classic_100_percent", proto.Ruleset_RulesetClassic, 100, 3 * time.Second, 1500 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sim := channelHasteFixture("shadow", tc.ruleset, tc.haste)
			p := sim.Raid.Parties[0].Players[0].(priest.PriestAgent).GetPriest()
			spell := p.MindFlay[6][0]
			if !spell.Cast(sim, p.CurrentTarget) {
				t.Fatal("Mind Flay was not castable")
			}
			dot := spell.Dot(p.CurrentTarget)
			if dot.Duration != tc.want || p.GCD.ReadyAt() != tc.gcd {
				t.Fatalf("channel %s, GCD %s; want %s and %s", dot.Duration, p.GCD.ReadyAt(), tc.want, tc.gcd)
			}
			for sim.CurrentTime < tc.want {
				if sim.Step() {
					t.Fatalf("fight ended before final channel tick at %s", tc.want)
				}
			}
			if dot.TickCount != 3 {
				t.Fatalf("Mind Flay delivered %d ticks; want 3", dot.TickCount)
			}
		})
	}
}

func TestForeverAPLChannelClipTracksHastedTickCadence(t *testing.T) {
	var affliction build
	for _, candidate := range builds() {
		if candidate.Key == "affliction" {
			affliction = candidate
			break
		}
	}
	if affliction.Key == "" {
		t.Fatal("no Affliction build")
	}
	var baselineSeconds, hastedSeconds float64
	for _, haste := range []float64{0, 100} {
		player := affliction.player(proto.Race_RaceOrc)
		player.TalentsString = "25350020135201051--05000551"
		player.BonusStats = &proto.UnitStats{Stats: stats.Stats{
			stats.SpellHaste: haste * core.HasteRatingPerHastePercent,
		}.ToFloatArray()}
		player.Rotation = core.APLRotationFromJsonString(`{"type":"TypeAPL","priorityList":[
			{"action":{"channelSpell":{"spellId":{"spellId":11704},"allowRecast":true,
				"interruptIf":{"cmp":{"lhs":{"spellChanneledTicks":{"spellId":{"spellId":11704}}},
				"op":"OpGe","rhs":{"const":{"val":"2"}}}}}}}
		]}`)
		req := request(player, 10, 1)
		req.Encounter.Duration = 40
		result := core.RunRaidSim(req)
		if result.Error != nil {
			t.Fatal(result.Error.Message)
		}
		for _, action := range result.RaidMetrics.Parties[0].Players[0].Actions {
			if action.Id.GetSpellId() != 11704 || action.Targets[0].Casts == 0 {
				continue
			}
			seconds := action.Targets[0].CastTimeMs / float64(action.Targets[0].Casts) / 1000
			if haste == 0 {
				baselineSeconds = seconds
			} else {
				hastedSeconds = seconds
			}
		}
	}
	if baselineSeconds < 1.5 || baselineSeconds > 2.1 ||
		hastedSeconds < .75 || hastedSeconds > 1.3 {
		t.Fatalf("two-tick APL clip should follow haste: baseline %.3fs, +100%% haste %.3fs", baselineSeconds, hastedSeconds)
	}
}

func TestForeverMissileBarrageTickCountWithHaste(t *testing.T) {
	sim := channelHasteFixture("arcane", proto.Ruleset_RulesetForever, 100)
	m := sim.Raid.Parties[0].Players[0].(mage.MageAgent).GetMage()
	spell := m.ArcaneMissiles[len(m.ArcaneMissiles)-1]
	m.MissileBarrageAura.Activate(sim)
	if !spell.Cast(sim, m.CurrentTarget) {
		t.Fatal("Missile Barrage was not castable")
	}
	dot := spell.Dot(m.CurrentTarget)
	want := time.Duration(dot.NumberOfTicks) * 250 * time.Millisecond
	if dot.Duration != want {
		t.Fatalf("Barrage channel %s, want %s", dot.Duration, want)
	}
	for sim.CurrentTime < want {
		if sim.Step() {
			t.Fatal("fight ended before the final missile")
		}
	}
	if dot.TickCount != dot.NumberOfTicks {
		t.Fatalf("Barrage delivered %d ticks, want %d", dot.TickCount, dot.NumberOfTicks)
	}
}

func TestBerserkingDoesNotShortenArcaneMissiles(t *testing.T) {
	req := racialFixture("arcane", proto.Race_RaceTroll)
	player := req.Raid.Parties[0].Players[0]
	player.Equipment = &proto.EquipmentSpec{}
	player.ForeverTier1Bonuses = false
	player.Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	m := sim.Raid.Parties[0].Players[0].(mage.MageAgent).GetMage()
	m.GetAura("Berserking").Activate(sim)
	if got := m.ApplyCastSpeed(time.Second); got != time.Second*10/11 {
		t.Fatalf("Berserking should shorten hardcasts; second became %s", got)
	}
	spell := m.ArcaneMissiles[len(m.ArcaneMissiles)-1]
	if !spell.Cast(sim, m.CurrentTarget) {
		t.Fatal("Arcane Missiles was not castable")
	}
	if got := spell.Dot(m.CurrentTarget).Duration; got != 5*time.Second {
		t.Fatalf("Berserking shortened Arcane Missiles to %s; want 5s", got)
	}
}

func TestForeverBerserkingSpeedIsNotEnergyHaste(t *testing.T) {
	req := racialFixture("combat", proto.Race_RaceTroll)
	player := req.Raid.Parties[0].Players[0]
	player.ForeverTier1Bonuses = false
	player.Rotation = &proto.APLRotation{}
	sim := core.NewSim(req, simsignals.Signals{})
	sim.Options.Interactive = true
	sim.Reset()
	r := sim.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
	energySpeed := r.PseudoStats.EnergyHasteMultiplier
	meleeSpeed := r.SwingSpeed()
	r.GetAura("Berserking").Activate(sim)
	if r.PseudoStats.EnergyHasteMultiplier != energySpeed {
		t.Fatal("Berserking's attack/casting-speed auras changed Energy regeneration")
	}
	if got := r.SwingSpeed() / meleeSpeed; math.Abs(got-1.1) > 1e-9 {
		t.Fatalf("Berserking multiplied melee speed %.6fx, want 1.1x", got)
	}
}

func TestForeverShocksShareHastedSpellGCD(t *testing.T) {
	for _, id := range []int32{10414, 29228} {
		req := racialFixture("enhancement", proto.Race_RaceSkyborneWindshaper)
		player := req.Raid.Parties[0].Players[0]
		player.Equipment = &proto.EquipmentSpec{}
		player.ForeverTier1Bonuses = false
		player.Rotation = &proto.APLRotation{}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		s := sim.Raid.Parties[0].Players[0].(shaman.ShamanAgent).GetShaman()
		spell := s.GetSpell(core.ActionID{SpellID: id})
		if spell == nil || !spell.Cast(sim, s.CurrentTarget) {
			t.Fatalf("shock %d was not castable", id)
		}
		baseGCD := core.GCDDefault
		want := time.Duration(float64(baseGCD) / 1.01)
		if got := s.GCD.ReadyAt(); got != want {
			t.Fatalf("shock %d GCD %s, want %s", id, got, want)
		}
	}
}

func TestForeverSpellGCDHasteFollowsSchoolNotDefenseType(t *testing.T) {
	for _, tc := range []struct {
		build string
		race  proto.Race
		id    int32
		want  time.Duration
	}{
		{"retribution", proto.Race_RaceUndead, 10333, time.Second},          // Holy Strike: Holy school, melee hit table.
		{"marksmanship", proto.Race_RaceOrc, 14287, time.Second},            // Arcane Shot: Arcane school, ranged hit table.
		{"enhancement", proto.Race_RaceOrc, 17364, 1500 * time.Millisecond}, // Stormstrike: Physical school.
	} {
		t.Run(tc.build, func(t *testing.T) {
			req := racialFixture(tc.build, tc.race)
			player := req.Raid.Parties[0].Players[0]
			player.ForeverTier1Bonuses = false
			player.Rotation = &proto.APLRotation{}
			player.BonusStats = &proto.UnitStats{Stats: stats.Stats{stats.SpellHaste: 100 * core.HasteRatingPerHastePercent}.ToFloatArray()}
			sim := core.NewSim(req, simsignals.Signals{})
			sim.Options.Interactive = true
			sim.Reset()
			unit := sim.Raid.Parties[0].Players[0].GetCharacter()
			spell := unit.GetSpell(core.ActionID{SpellID: tc.id})
			if spell == nil || !spell.Cast(sim, unit.CurrentTarget) {
				t.Fatalf("spell %d was not castable", tc.id)
			}
			if got := unit.GCD.ReadyAt(); got != tc.want {
				t.Fatalf("spell %d GCD %s; want %s", tc.id, got, tc.want)
			}
		})
	}
}

func TestMinorHasteEnchantCountsOnceForAttackSpeeds(t *testing.T) {
	attackSpeeds := func(withEnchant bool) (float64, float64, time.Duration) {
		req := racialFixture("marksmanship", proto.Race_RaceTroll)
		player := req.Raid.Parties[0].Players[0]
		player.ForeverTier1Bonuses = false
		player.Rotation = &proto.APLRotation{}
		gloves := player.Equipment.Items[proto.ItemSlot_ItemSlotHands]
		gloves.Enchant = 0
		if withEnchant {
			gloves.Enchant = 931
		}
		sim := core.NewSim(req, simsignals.Signals{})
		sim.Options.Interactive = true
		sim.Reset()
		h := sim.Raid.Parties[0].Players[0].(hunter.HunterAgent).GetHunter()
		return h.SwingSpeed(), h.RangedSwingSpeed(), h.ApplyCastSpeed(time.Second)
	}
	meleeWithout, rangedWithout, castWithout := attackSpeeds(false)
	meleeWith, rangedWith, castWith := attackSpeeds(true)
	for name, got := range map[string]float64{"melee": meleeWith / meleeWithout, "ranged": rangedWith / rangedWithout} {
		if math.Abs(got-1.01) > 1e-9 {
			t.Fatalf("Minor Haste multiplied %s speed %.6fx, want 1.01x", name, got)
		}
	}
	if got := float64(castWithout) / float64(castWith); math.Abs(got-1.01) > 1e-8 {
		t.Fatalf("Minor Haste multiplied hardcast speed %.6fx, want 1.01x", got)
	}
}

func TestForeverPenanceKeepsImmediateBoltAtHaste(t *testing.T) {
	sim := channelHasteFixture("smite", proto.Ruleset_RulesetForever, 100)
	p := sim.Raid.Parties[0].Players[0].(priest.PriestAgent).GetPriest()
	spell := p.Penance
	if !spell.Cast(sim, p.CurrentTarget) {
		t.Fatal("Penance was not castable")
	}
	target := p.CurrentTarget.UnitIndex
	damage := spell.SpellMetrics[target].TotalDamage
	if damage == 0 || spell.Dot(p.CurrentTarget).Duration != time.Second {
		t.Fatal("the first bolt must land immediately, with two hasted bolts remaining")
	}
	for _, at := range []time.Duration{500 * time.Millisecond, time.Second} {
		for sim.CurrentTime < at {
			if sim.Step() {
				t.Fatalf("fight ended before Penance bolt at %s", at)
			}
		}
		next := spell.SpellMetrics[target].TotalDamage
		if sim.CurrentTime != at || next <= damage {
			t.Fatalf("Penance did not deal a bolt at %s", at)
		}
		damage = next
	}
}

func TestRogueGCDDoesNotInheritSpellHaste(t *testing.T) {
	sim := channelHasteFixture("combat", proto.Ruleset_RulesetForever, 100)
	r := sim.Raid.Parties[0].Players[0].(rogue.RogueAgent).GetRogue()
	if !r.SinisterStrike.Cast(sim, r.CurrentTarget) {
		t.Fatal("Sinister Strike was not castable")
	}
	if r.GCD.ReadyAt() != time.Second {
		t.Fatalf("Rogue GCD %s, want 1s", r.GCD.ReadyAt())
	}
}
