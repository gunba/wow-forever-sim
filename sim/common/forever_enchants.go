package common

import (
	"fmt"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	"math"
	"time"
)

func init() {
	core.NewEnchantEffect(8721, applyRecoveryEnchant)
	core.NewEnchantEffect(8217, applyRevelationEnchant)
	// 1310308 is a 2% crit aura restricted by SpellEquippedItems to bows,
	// guns and crossbows (subclass mask 262156), not general melee/spell crit.
	core.NewEnchantEffect(8720, func(agent core.Agent) {
		c := agent.GetCharacter()
		const label = "SAF-T Ultra Precision Scope"
		if c.GetAura(label) != nil {
			return
		}
		equipped := func() bool { return c.Equipment[proto.ItemSlot_ItemSlotRanged].Enchant.EffectID == 8720 }
		adjust := func(amount float64) {
			for _, spell := range c.Spellbook {
				if spell.CastType == proto.CastType_CastTypeRanged {
					spell.BonusCritRating += amount
				}
			}
		}
		applied := false
		aura := c.RegisterAura(core.Aura{
			Label: label, ActionID: core.ActionID{ItemID: 279272},
			Duration: core.NeverExpires, BuildPhase: core.CharacterBuildPhaseGear,
			OnGain: func(_ *core.Aura, _ *core.Simulation) {
				applied = equipped()
				if applied {
					adjust(2)
				}
			},
			OnExpire: func(_ *core.Aura, _ *core.Simulation) {
				if applied {
					adjust(-2)
				}
				applied = false
			},
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				if equipped() {
					aura.Activate(sim)
				}
			},
		})
		c.OnSpellRegistered(func(spell *core.Spell) {
			if applied && spell.CastType == proto.CastType_CastTypeRanged {
				spell.BonusCritRating += 2
			}
		})
		c.RegisterOnItemSwap(func(sim *core.Simulation) {
			aura.Deactivate(sim)
			if equipped() {
				aura.Activate(sim)
			}
		})
	})
	// Client 69893: 22841 has melee/ranged speed auras 319/140;
	// 7217 has only melee speed. Neither is general haste or Energy haste.
	registerAttackSpeedEnchant(2543, 22841, "Arcanum of Rapidity", 1, true)
	registerAttackSpeedEnchant(34, 7217, "Iron Counterweight", 3, false)
	// Minor Haste (13928) also has casting-speed aura 65. A casting-speed
	// aura is not the SpellHaste stat: Berserking has the same aura type but
	// does not shorten Arcane Missiles. Its melee haste remains in the
	// enchant's stats and participates in the provisional Energy model.
	core.NewEnchantEffect(931, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.Env.IsForever() {
			character.PseudoStats.CastSpeedMultiplier *= 1.01
		}
	})
}

// Recovery is an equip aura (SpellItemEnchantment effect 3), not a per-hand
// weapon-hit proc. 70291 driver1248761 admits melee autos/specials; its sourced
// outcome filter is dodge/parry. A single owner ICD also prevents dual copies.
func applyRecoveryEnchant(agent core.Agent) {
	character := agent.GetCharacter()
	if !character.Env.IsForever() || character.GetAura("Recovery") != nil {
		return
	}
	icd := core.Cooldown{Timer: character.NewTimer(), Duration: 10 * time.Second}
	heal := character.RegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: 1248759},
		SpellSchool:      core.SpellSchoolPhysical,
		ProcMask:         core.ProcMaskEmpty,
		Flags:            core.SpellFlagPassiveSpell | core.SpellFlagHelpful | core.SpellFlagNoOnCastComplete,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			spell.CalcAndDealHealing(sim, &character.Unit, character.MaxHealth()*.05, spell.OutcomeHealing)
		},
	})
	core.MakePermanent(character.RegisterAura(core.Aura{
		Label: "Recovery", ActionID: core.ActionID{SpellID: 1248761}, Icd: &icd,
		OnSpellHitDealt: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !spell.ProcMask.Matches(core.ProcMaskMelee) || !result.Outcome.Matches(core.OutcomeDodge|core.OutcomeParry) || !icd.IsReady(sim) {
				return
			}
			if character.MainHand().Enchant.EffectID != 8721 && character.OffHand().Enchant.EffectID != 8721 {
				return
			}
			icd.Use(sim)
			heal.Cast(sim, &character.Unit)
		},
	}))
}

// ValidateForeverRevelationModel separates supported simulation payloads from
// legal item equipment. The client does not provide Warrior/Rogue payloads.
func ValidateForeverRevelationModel(class proto.Class, model *proto.ForeverRevelationModel) error {
	if revelationPayload(class) == 0 {
		return fmt.Errorf("Revelation8217: unsupported source class payload for %s (not an item class restriction)", class)
	}
	if model == nil || !model.Enabled {
		return fmt.Errorf("Revelation8217 requires an explicitly enabled provisional model")
	}
	if math.IsNaN(model.BaseChance) || math.IsInf(model.BaseChance, 0) || model.BaseChance < 0 || model.BaseChance > 1 {
		return fmt.Errorf("Revelation8217 provisional baseChance must be finite in [0,1]")
	}
	if math.IsNaN(model.CritExponent) || math.IsInf(model.CritExponent, 0) || model.CritExponent < 0 {
		return fmt.Errorf("Revelation8217 provisional critExponent must be finite and nonnegative")
	}
	return nil
}

// Conditional on an eligible NONCRITICAL direct event, not a flat chance per
// cast. This law and the outcome/event sampling convention are provisional.
func ForeverRevelationProcChance(model *proto.ForeverRevelationModel, effectiveCrit float64) float64 {
	return model.BaseChance * math.Pow(1-max(0, min(1, effectiveCrit)), model.CritExponent)
}

func revelationPayload(class proto.Class) int32 {
	switch class {
	case proto.Class_ClassMage:
		return 1248808
	case proto.Class_ClassPriest:
		return 1323377
	case proto.Class_ClassWarlock:
		return 1323392
	case proto.Class_ClassDruid:
		return 1323400
	case proto.Class_ClassHunter:
		return 1323410
	case proto.Class_ClassShaman:
		return 1323418
	case proto.Class_ClassPaladin:
		return 1323419
	}
	return 0
}

func revelationSpellEligible(class proto.Class, spell *core.Spell) bool {
	return revelationSpellIDs[class][spell.SpellID] &&
		spell.ProcMask.Matches(core.ProcMaskSpellDamage|core.ProcMaskSpellHealing|core.ProcMaskMeleeSpecial|core.ProcMaskRangedSpecial)
}

// Primary 70291 description tokens pin the damaging child for each rank.
// Every child is SCHOOL_DAMAGE, intersects its class payload, and lacks
// SPELL_ATTR2_CANT_CRIT (0x20000000). These legacy tick-ID aliases do not
// reclassify the channel's outcome-only container or any ordinary DoT/HoT.
var revelationPeriodicChildren = map[int32]int32{
	402174: 402284, 1240720: 1240727, 1240721: 1240730, 1316995: 1316993, // Penance
	10: 1279976, 6141: 1279977, 8427: 1279978, 10185: 1279979, 10186: 1279980, 10187: 1279949, // Blizzard
	16914: 1278965, 17401: 1278968, 17402: 1278759, // Hurricane
	1510: 1279721, 14294: 1279719, 14295: 1279715, // Volley
	5740: 1282380, 6219: 1282383, 11677: 1282384, 11678: 1282385, // Rain of Fire
}

func applyRevelationEnchant(agent core.Agent) {
	c := agent.GetCharacter()
	if !c.Env.IsForever() || c.GetAura("Revelation") != nil {
		return
	}
	if err := ValidateForeverRevelationModel(c.Class, c.ForeverRevelationModel); err != nil {
		panic(err)
	}
	model := c.ForeverRevelationModel
	equipped := func() bool {
		return c.MainHand().Enchant.EffectID == 8217 || c.OffHand().Enchant.EffectID == 8217
	}
	adjust := func(amount float64) {
		for _, spell := range c.Spellbook {
			if revelationSpellEligible(c.Class, spell) {
				spell.BonusDirectCritRating += amount
			}
		}
	}
	// Primary payload: one charge, +100 percentage points, 15 seconds, forms
	// retained. No global crit stat mutation: DoT/HoT/wand rolls are untouched.
	buff := c.RegisterAura(core.Aura{
		Label: "Revelation", ActionID: core.ActionID{SpellID: revelationPayload(c.Class)},
		Duration: 15 * time.Second, MaxStacks: 1,
		OnGain:   func(_ *core.Aura, _ *core.Simulation) { adjust(100 * core.SpellCritRatingPerCritChance) },
		OnExpire: func(_ *core.Aura, _ *core.Simulation) { adjust(-100 * core.SpellCritRatingPerCritChance) },
	})
	c.OnSpellRegistered(func(spell *core.Spell) {
		if !revelationSpellEligible(c.Class, spell) {
			return
		}
		if child := revelationPeriodicChildren[spell.SpellID]; child != 0 && revelationSpellIDs[c.Class][child] && spell.Flags.Matches(core.SpellFlagChanneled) {
			spell.RevelationPeriodicDirectChildID = child
		}
		if buff.IsActive() {
			spell.BonusDirectCritRating += 100 * core.SpellCritRatingPerCritChance
		}
		previous := spell.ConsumeDirectCritBonus
		spell.ConsumeDirectCritBonus = func(sim *core.Simulation) {
			if previous != nil {
				previous(sim)
			}
			// Reservation at the first eligible crit-capable outcome sample
			// prevents simultaneous/in-flight direct children sharing a charge.
			buff.Deactivate(sim)
		}
	})
	onDirectEvent := func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
		if !equipped() || spell.Unit != &c.Unit || !result.DirectDamageOrHealing ||
			!revelationSpellEligible(c.Class, spell) || result.DidCrit() ||
			(!result.Landed() && (!model.TriggerOnMiss || !result.Outcome.Matches(core.OutcomeMiss))) {
			return
		}
		if sim.Proc(ForeverRevelationProcChance(model, result.EffectiveCritChance), "Provisional Revelation") {
			buff.Activate(sim)
			buff.SetStacks(sim, 1)
		}
	}
	core.MakePermanent(c.RegisterAura(core.Aura{
		Label: "Revelation Driver", ActionID: core.ActionID{SpellID: 1248806},
		OnSpellHitDealt: onDirectEvent, OnHealDealt: onDirectEvent,
		OnPeriodicDamageDealt: onDirectEvent, // source-backed logical direct children only

	}))
	c.RegisterOnItemSwap(func(sim *core.Simulation) {
		if !equipped() {
			buff.Deactivate(sim)
		}
	})
}

// Exact SpellClassOptions family/mask intersection in client1.60.1.70291.
// Runtime qualification still requires a direct damage/heal result; source
// channel containers and retained action-ID aliases cannot qualify by ID alone.
var revelationSpellIDs = map[proto.Class]map[int32]bool{
	proto.Class_ClassMage: { // payload1248808 family3 mask[547494642, 0, 0, 8]
		10: true, 116: true, 120: true, 122: true, 205: true, 837: true, 865: true, 1449: true, 2120: true, 2121: true,
		2136: true, 2137: true, 2138: true, 2948: true, 5143: true, 5144: true, 5145: true, 6131: true, 6141: true, 7268: true,
		7269: true, 7270: true, 7322: true, 8406: true, 8407: true, 8408: true, 8412: true, 8413: true, 8416: true, 8417: true,
		8418: true, 8419: true, 8422: true, 8423: true, 8427: true, 8437: true, 8438: true, 8439: true, 8444: true, 8445: true,
		8446: true, 8492: true, 10159: true, 10160: true, 10161: true, 10179: true, 10180: true, 10181: true, 10185: true, 10186: true,
		10187: true, 10197: true, 10199: true, 10201: true, 10202: true, 10205: true, 10206: true, 10207: true, 10211: true, 10212: true,
		10215: true, 10216: true, 10230: true, 10273: true, 10274: true, 11113: true, 12557: true, 12611: true, 13018: true, 13019: true,
		13020: true, 13021: true, 13339: true, 13340: true, 13341: true, 13342: true, 13374: true, 13878: true, 14145: true, 15041: true,
		15091: true, 15241: true, 15244: true, 15574: true, 15735: true, 15736: true, 15744: true, 15790: true, 15791: true, 16046: true,
		16070: true, 16144: true, 16785: true, 17145: true, 17195: true, 17273: true, 17274: true, 17276: true, 17277: true, 17492: true,
		20228: true, 20229: true, 20623: true, 20679: true, 20795: true, 20828: true, 20832: true, 21229: true, 22272: true, 22273: true,
		22424: true, 22460: true, 22746: true, 23039: true, 23113: true, 23331: true, 24530: true, 25028: true, 25049: true, 25304: true,
		25345: true, 25346: true, 27618: true, 28323: true, 28863: true, 29163: true, 29515: true, 30092: true, 30095: true, 31378: true,
		31751: true, 42896: true, 350025: true, 400574: true, 400616: true, 400618: true, 400619: true, 400620: true, 400621: true, 400622: true,
		400623: true, 400640: true, 425225: true, 434443: true, 460858: true, 460860: true, 461896: true, 469055: true, 469056: true, 469057: true,
		469058: true, 469161: true, 1235318: true, 1236173: true, 1236291: true, 1239696: true, 1239697: true, 1239699: true, 1239700: true, 1240044: true,
		1240045: true, 1240046: true, 1240047: true, 1279949: true, 1279976: true, 1279977: true, 1279978: true, 1279979: true, 1279980: true, 1308651: true,
		1308937: true, 1312002: true, 1318432: true,
	},
	proto.Class_ClassPriest: { // payload1323377 family6 mask[161758864, 8486934, 128, 0]
		585: true, 591: true, 596: true, 598: true, 984: true, 996: true, 1004: true, 2050: true, 2052: true, 2053: true,
		2054: true, 2055: true, 2060: true, 2061: true, 6060: true, 6063: true, 6064: true, 8092: true, 8102: true, 8103: true,
		8104: true, 8105: true, 8106: true, 8129: true, 8131: true, 9472: true, 9473: true, 9474: true, 10797: true, 10874: true,
		10875: true, 10876: true, 10915: true, 10916: true, 10917: true, 10933: true, 10934: true, 10945: true, 10946: true, 10947: true,
		10960: true, 10961: true, 10963: true, 10964: true, 10965: true, 13908: true, 15007: true, 15068: true, 15407: true, 16568: true,
		17137: true, 17138: true, 17165: true, 17311: true, 17312: true, 17313: true, 17314: true, 17843: true, 18807: true, 19236: true,
		19238: true, 19240: true, 19241: true, 19242: true, 19243: true, 19296: true, 19299: true, 19302: true, 19303: true, 19304: true,
		19305: true, 22203: true, 22205: true, 22313: true, 22800: true, 22802: true, 22803: true, 22821: true, 22919: true, 23455: true,
		23458: true, 23459: true, 23642: true, 23953: true, 23979: true, 24152: true, 25314: true, 25316: true, 26044: true, 26143: true,
		26565: true, 27286: true, 27608: true, 27636: true, 27640: true, 27803: true, 27804: true, 27805: true, 28309: true, 28310: true,
		28883: true, 29407: true, 401937: true, 401955: true, 402174: true, 402261: true, 402277: true, 402284: true, 402289: true, 412526: true,
		474204: true, 474268: true, 1215740: true, 1232758: true, 1236153: true, 1238817: true, 1240720: true, 1240721: true, 1240723: true, 1240724: true,
		1240727: true, 1240730: true, 1240732: true, 1240733: true, 1240734: true, 1240736: true, 1240770: true, 1240771: true, 1240772: true, 1240773: true,
		1240774: true, 1277331: true, 1277332: true, 1277333: true, 1277334: true, 1277335: true, 1277370: true, 1277371: true, 1277372: true, 1277374: true,
		1277376: true, 1277377: true, 1277378: true, 1293973: true, 1309595: true, 1309633: true, 1309635: true, 1309636: true, 1310076: true, 1316991: true,
		1316992: true, 1316993: true, 1316994: true, 1316995: true,
	},
	proto.Class_ClassWarlock: { // payload1323392 family5 mask[541161, 8650816, 0, 0]
		686: true, 689: true, 695: true, 699: true, 705: true, 709: true, 1088: true, 1106: true, 1120: true, 1949: true,
		5676: true, 5740: true, 5857: true, 6219: true, 6353: true, 6789: true, 7641: true, 7651: true, 8288: true, 8289: true,
		11659: true, 11660: true, 11661: true, 11675: true, 11677: true, 11678: true, 11681: true, 11682: true, 11683: true, 11684: true,
		11699: true, 11700: true, 17877: true, 17919: true, 17920: true, 17921: true, 17922: true, 17923: true, 17924: true, 17925: true,
		17926: true, 17962: true, 18867: true, 18868: true, 18869: true, 18870: true, 18871: true, 18930: true, 18931: true, 18932: true,
		24826: true, 25307: true, 28412: true, 350026: true, 403501: true, 403677: true, 403685: true, 403686: true, 403687: true, 403688: true,
		403689: true, 403858: true, 412758: true, 460692: true, 460698: true, 460699: true, 460700: true, 1282380: true, 1282383: true, 1282384: true,
		1282385: true, 1293693: true, 1293694: true, 1293812: true, 1293813: true, 1293817: true, 1293818: true, 1316697: true,
	},
	proto.Class_ClassDruid: { // payload1323400 family7 mask[4194341, 2, 0, 0]
		2912: true, 5176: true, 5177: true, 5178: true, 5179: true, 5180: true, 5185: true, 5186: true, 5187: true, 5188: true,
		5189: true, 6778: true, 6780: true, 8903: true, 8905: true, 8949: true, 8950: true, 8951: true, 9758: true, 9875: true,
		9876: true, 9888: true, 9889: true, 9912: true, 16914: true, 17401: true, 17402: true, 18562: true, 20687: true, 21668: true,
		25297: true, 25298: true, 27530: true, 405953: true, 429820: true, 449431: true, 463285: true, 1278759: true, 1278965: true, 1278968: true,
	},
	proto.Class_ClassHunter: { // payload1323410 family9 mask[10240, 0, 0, 0]
		1510: true, 3044: true, 14281: true, 14282: true, 14283: true, 14284: true, 14285: true, 14286: true, 14287: true, 14294: true,
		14295: true, 19597: true, 19676: true, 19677: true, 19678: true, 19679: true, 19680: true, 19681: true, 19682: true, 19683: true,
		19684: true, 19685: true, 19686: true, 22908: true, 1271102: true, 1277851: true, 1278060: true, 1278061: true, 1279715: true, 1279719: true,
		1279721: true, 1280004: true, 1280044: true,
	},
	proto.Class_ClassShaman: { // payload1323418 family11 mask[2148532675, 266240, 0, 0]
		331: true, 332: true, 403: true, 421: true, 529: true, 547: true, 548: true, 913: true, 915: true, 930: true,
		939: true, 943: true, 959: true, 1064: true, 2860: true, 6041: true, 8004: true, 8005: true, 8008: true, 8010: true,
		8042: true, 8044: true, 8045: true, 8046: true, 8056: true, 8058: true, 10391: true, 10392: true, 10395: true, 10396: true,
		10412: true, 10413: true, 10414: true, 10466: true, 10467: true, 10468: true, 10472: true, 10473: true, 10605: true, 10622: true,
		10623: true, 15207: true, 15208: true, 25357: true, 27624: true, 408341: true, 408342: true, 408343: true, 408344: true, 408345: true,
		408423: true, 408424: true, 408426: true, 408427: true, 408428: true, 408439: true, 408440: true, 408441: true, 408442: true, 408443: true,
		408472: true, 408473: true, 408474: true, 408475: true, 408477: true, 408479: true, 408481: true, 408482: true, 408484: true, 408490: true,
		408491: true, 408681: true, 408683: true, 408685: true, 408687: true, 408688: true, 408689: true, 408690: true, 434357: true, 436376: true,
		1220744: true, 1220746: true, 1220747: true, 1220748: true, 1220749: true, 1220750: true, 1220751: true, 1238299: true, 1238300: true, 1238373: true,
		1238376: true,
	},
	proto.Class_ClassPaladin: { // payload1323419 family10 mask[3223323648, 2097792, 64, 67108864]
		635: true, 639: true, 647: true, 1026: true, 1042: true, 2812: true, 3472: true, 10318: true, 10328: true, 10329: true,
		19750: true, 19939: true, 19940: true, 19941: true, 19942: true, 19943: true, 20183: true, 20187: true, 20280: true, 20281: true,
		20282: true, 20283: true, 20284: true, 20285: true, 20286: true, 20411: true, 20412: true, 20413: true, 20414: true, 20467: true,
		20473: true, 20929: true, 20930: true, 20963: true, 20964: true, 20965: true, 20966: true, 24239: true, 24274: true, 24275: true,
		25292: true, 25902: true, 25903: true, 25911: true, 25912: true, 25913: true, 25914: true, 429145: true, 429146: true, 429151: true,
		444894: true, 458856: true, 1310909: true, 1310910: true, 1310911: true, 1310912: true, 1310914: true, 1311590: true, 1311591: true, 1311592: true,
		1311593: true, 1311594: true, 1311595: true, 1311596: true, 1311597: true, 1311598: true, 1311599: true, 1311604: true, 1311605: true, 1311606: true,
		1311650: true, 1311655: true, 1311806: true,
	},
}

func registerAttackSpeedEnchant(id, spellID int32, label string, percent float64, ranged bool) {
	core.NewEnchantEffect(id, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.GetAura(label) != nil {
			return
		}
		count := func() float64 {
			n := 0.0
			for _, item := range character.Equipment {
				if item.Enchant.EffectID == id {
					n++
				}
			}
			return n
		}
		var haste, multiplier float64
		aura := character.RegisterAura(core.Aura{
			Label: label, ActionID: core.ActionID{SpellID: spellID},
			Duration: core.NeverExpires, BuildPhase: core.CharacterBuildPhaseGear,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				if count() > 0 {
					aura.Activate(sim)
				}
			},
			OnGain: func(_ *core.Aura, sim *core.Simulation) {
				haste = percent * count()
				multiplier = 1 + haste/100
				if !character.Env.IsForever() {
					character.AddStatDynamic(sim, stats.MeleeHaste, haste)
					return
				}
				character.MultiplyMeleeSpeed(sim, multiplier)
				if ranged {
					character.MultiplyRangedSpeed(sim, multiplier)
				}
			},
			OnExpire: func(_ *core.Aura, sim *core.Simulation) {
				if !character.Env.IsForever() {
					character.AddStatDynamic(sim, stats.MeleeHaste, -haste)
					return
				}
				character.MultiplyMeleeSpeed(sim, 1/multiplier)
				if ranged {
					character.MultiplyRangedSpeed(sim, 1/multiplier)
				}
			},
		})
		character.RegisterOnItemSwap(func(sim *core.Simulation) {
			aura.Deactivate(sim)
			if count() > 0 {
				aura.Activate(sim)
			}
		})
	})
}
