"""Value the proc and on-use part of items for every benchmark profile.

Used by generate.py --item-effects ITEMS.json, where ITEMS.json lists items with their client ItemEffect rows
({id, name, class, subclass, inventoryType, speed, item_effects: [{trigger, spell, cooldown_ms}], ...}).

- Items the simulator already implements keep its effect ("method": "simulated").
- Other items get a generic effect read from the client spell tables (damage, damage over time, stat, attack power,
  spell damage, haste and crit buffs, extra attacks, armor reduction), registered by the tool binary's -itemfx
  ("method": "estimated": the effect is simulated, but its proc rate is assumed when the client tables do not give one).
- Everything else is "not_modelled" with the reason.

Each value is the profile's score with the item's effect minus the score with an identical item without it (same random
seed), in the profile's usual units (main stat = 1). The item's own stats are excluded: both items carry none.
"""
import collections, copy, csv, hashlib, json, os, re, sys

HERE = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, os.path.join(HERE, "..", "data_watch"))
import spell_client  # noqa: E402

SCHOOL_BITS = [(1, "physical"), (2, "holy"), (4, "fire"), (8, "nature"), (16, "frost"), (32, "shadow"), (64, "arcane")]
STAT_FROM_MISC = {0: 0, 1: 1, 2: 2, 3: 3, 4: 4}              # MOD_STAT misc value -> sim stat index
SCHOOL_DAMAGE_STAT = {2: 9, 4: 7, 8: 10, 16: 8, 32: 11, 64: 6}  # MOD_DAMAGE_DONE school -> <school>_damage stat
ASSUMED_COH_PPM = 1.0
PHYSICAL = ("dw", "2h", "shield", "feral")
TRINKET_SLOT, MH, OH, RANGED = 12, 14, 15, 16
# Proc flags (SpellAuraOptions.ProcTypeMask_0).
DEAL_MELEE, TAKE_MELEE, DEAL_RANGED, DEAL_SPELL = 0x4 | 0x10, 0x8 | 0x20, 0x40 | 0x100, 0x10000 | 0x1000 | 0x40000


class Tables:
    def __init__(self, build):
        t = lambda n: spell_client.table(build, n)
        self.name = {int(r["ID"]): r["Name_lang"] for r in t("SpellName")}
        self.eff = collections.defaultdict(list)
        for r in t("SpellEffect"):
            self.eff[int(r["SpellID"])].append(r)
        self.misc = {int(r["SpellID"]): r for r in t("SpellMisc") if r["DifficultyID"] == "0"}
        self.dur = {int(r["ID"]): int(r["Duration"]) for r in t("SpellDuration")}
        self.opt = {int(r["SpellID"]): r for r in t("SpellAuraOptions") if r["DifficultyID"] == "0"}
        self.ppm = {int(r["ID"]): float(r["BaseProcRate"]) for r in t("SpellProcsPerMinute")}

    def duration(self, sid):
        d = self.dur.get(int(self.misc.get(sid, {}).get("DurationIndex") or 0), 0)
        return d if d > 0 else 0

    def school(self, sid):
        mask = int(self.misc.get(sid, {}).get("SchoolMask") or 1)
        return next((n for bit, n in SCHOOL_BITS if mask & bit), "physical")

    def effects(self, sid):
        return sorted(self.eff.get(sid, []), key=lambda e: int(e["EffectIndex"]))


def num(v):
    return float(v or 0)


def spell_effects(T, sid, depth=0):
    """Generic effects of a spell (following triggered spells), plus reasons for the parts that are not modelled."""
    out, why = [], []
    if depth > 3 or sid not in T.misc:
        return out, [f"spell {sid} not in client tables"] if depth == 0 else []
    dur, school, stacks = T.duration(sid), T.school(sid), int(num(T.opt.get(sid, {}).get("CumulativeAura")))
    buff = collections.Counter()
    for e in T.effects(sid):
        eff, aura, bp, var = int(e["Effect"]), int(e["EffectAura"]), num(e["EffectBasePointsF"]), num(e["Variance"])
        misc, period, trig, target = int(num(e["EffectMiscValue_0"])), int(num(e["EffectAuraPeriod"])), int(num(e["EffectTriggerSpell"])), int(num(e["ImplicitTarget_0"]))
        if eff in (2, 9) and bp > 0:                     # school damage, health leech
            out.append({"kind": "damage", "spell": sid, "school": school, "min": bp * (1 - var / 2), "range": bp * var})
        elif eff == 19:
            out.append({"kind": "extra_attacks", "spell": sid, "n": max(1, int(bp))})
        elif eff in (64, 142) and trig:
            o, w = spell_effects(T, trig, depth + 1); out += o; why += w
        elif eff in (6, 35, 65, 119, 128, 129):           # aura, party/raid/area auras
            if aura == 3 and bp > 0 and period > 0:
                out.append({"kind": "dot", "spell": sid, "school": school, "tick": bp, "period_ms": period,
                            "ticks": max(1, (dur or period) // period)})
            elif aura == 29 and bp > 0:
                for m in ([0, 1, 2, 3, 4] if misc == -1 else [misc]):
                    if m in STAT_FROM_MISC:
                        buff[STAT_FROM_MISC[m]] += bp
            elif aura == 99 and bp > 0:
                buff[17] += bp
            elif aura == 124 and bp > 0:
                buff[27] += bp
            elif aura == 13 and bp > 0:
                if misc & 126 == 126:
                    buff[42] += bp
                else:
                    for bit, stat in SCHOOL_DAMAGE_STAT.items():
                        if misc & bit:
                            buff[stat] += bp
            elif aura in (52, 290) and bp > 0:
                buff[19] += bp
            elif aura == 71 and bp > 0:
                buff[14] += bp
            elif aura in (138, 319, 192) and bp > 0:
                out.append({"kind": "haste", "spell": sid, "pct": bp, "duration_ms": dur or 10000})
            elif aura == 22 and bp < 0 and misc & 1:
                out.append({"kind": "armor_debuff", "spell": sid, "amount": -bp, "duration_ms": dur or 30000, "stacks": max(1, stacks)})
            elif aura == 23 and trig:
                o, w = spell_effects(T, trig, depth + 1)
                ticks = max(1, (dur or period) // max(period, 1))
                for x in o:
                    if x["kind"] == "damage":
                        out.append({"kind": "dot", "spell": trig, "school": x["school"], "tick": x["min"] + x["range"] / 2,
                                    "period_ms": period, "ticks": ticks})
                    else:
                        why.append(f"periodic {x['kind']} not modelled")
                why += w
            elif aura not in (0, 42):
                why.append(f"aura {aura} on spell {sid} ({T.name.get(sid)}) not modelled")
        elif eff not in (0,):
            why.append(f"effect {eff} on spell {sid} ({T.name.get(sid)}) not modelled")
    if buff:
        out.append({"kind": "stats", "spell": sid, "stats": {str(k): v for k, v in buff.items()}, "duration_ms": dur or 10000})
    return out, why


def classify(item, T):
    """Generic spec(s) for an item, or reasons it is not modelled."""
    specs, why = [], []
    weapon = item["class"] == 2
    for ie in item["item_effects"]:
        sid, trig = ie["spell"], ie["trigger"]
        base = {"item": item["id"], "name": item["name"][:40]}
        if trig == 2:
            eff, w = spell_effects(T, sid)
            why += w
            if eff:
                specs.append({**base, "trigger": "coh" if weapon else "melee_hit", "ppm": ASSUMED_COH_PPM, "effects": eff,
                              "assumed": f"{ASSUMED_COH_PPM} proc per minute"})
        elif trig == 1:
            proc = [e for e in T.effects(sid) if int(e["EffectAura"]) == 42 and int(num(e["EffectTriggerSpell"]))]
            shield = [e for e in T.effects(sid) if int(e["EffectAura"]) in (15, 43)]     # damage shield, damage on hit taken
            if proc:
                o = T.opt.get(sid, {})
                mask = int(num(o.get("ProcTypeMask_0")))
                chance, icd = num(o.get("ProcChance")), int(num(o.get("ProcCategoryRecovery")))
                trigger = ("melee_hit" if mask & DEAL_MELEE else "struck" if mask & TAKE_MELEE else
                           "spell_hit" if mask & DEAL_SPELL else "ranged_hit" if mask & DEAL_RANGED else None)
                eff, w = spell_effects(T, int(num(proc[0]["EffectTriggerSpell"])))
                why += w
                if trigger is None:
                    why.append(f"proc flags {hex(mask)} not modelled")
                elif eff:
                    spec = {**base, "trigger": trigger, "effects": eff, "icd_ms": icd}
                    ppm = T.ppm.get(int(num(o.get("SpellProcsPerMinuteID"))))
                    if ppm:                                   # the client's procs-per-minute rate overrides ProcChance
                        spec["ppm"] = ppm
                    elif 0 < chance < 100:
                        spec["chance"] = chance / 100
                    else:                                     # 0, or 100% with no rate: the server's rate is unknown
                        spec["ppm"], spec["assumed"] = ASSUMED_COH_PPM, f"{ASSUMED_COH_PPM} proc per minute"
                    specs.append(spec)
            elif shield:
                bp, var = num(shield[0]["EffectBasePointsF"]), num(shield[0]["Variance"])
                specs.append({**base, "trigger": "struck", "chance": 1.0,
                              "effects": [{"kind": "damage", "spell": sid, "school": T.school(sid), "min": bp * (1 - var / 2), "range": bp * var}]})
            else:
                auras = sorted({int(e["EffectAura"]) for e in T.effects(sid)})
                why.append(f"equip effect (aura {auras}) is a spell modifier, scripted or utility")
        elif trig == 0:
            eff, w = spell_effects(T, sid)
            why += w
            if eff:
                specs.append({**base, "trigger": "use", "cooldown_ms": max(ie.get("cooldown_ms") or 0, 0) or 120000, "effects": eff})
    return specs, why


def relevant(b, spec):
    """Whether a spec can act for profile b, and which slot it is tested in."""
    w = b["weapons"]
    kinds = {e["kind"] for e in spec["effects"]}
    melee = w in PHYSICAL or (w == "hunter" and "mh" in b["slots"])
    trig = spec["trigger"]
    if trig in ("coh", "melee_hit"):
        return melee and w != "feral" if trig == "coh" else melee
    if trig == "ranged_hit":
        return w == "hunter"
    if trig == "spell_hit":
        return w == "caster" or b["cls"] in ("paladin", "shaman", "druid")
    if trig == "struck":
        return b["tank"]
    if trig == "use":
        if kinds & {"damage", "dot", "armor_debuff"}:
            return True
        if kinds & {"haste", "extra_attacks"}:
            return melee or w == "hunter"
        stats = {int(k) for e in spec["effects"] if e["kind"] == "stats" for k in e["stats"]}
        physical_stats, caster_stats = {0, 1, 17, 19, 27}, {3, 4, 5, 6, 7, 8, 9, 10, 11, 14, 42}
        return bool(stats & (physical_stats if w != "caster" else caster_stats)) or (b["tank"] and 2 in stats)
    return False


def weapon_slots(item, b):
    """Slots to test a chance-on-hit weapon in for profile b (inventory types: 13 one-hand, 21 main, 22 off, 17 two-hand,
    15/25/26 ranged)."""
    inv, w = item["inventoryType"], b["weapons"]
    if inv == 17:
        return [MH] if w == "2h" else []
    if inv in (13, 21, 22):
        hands = {13: [MH, OH], 21: [MH], 22: [OH]}[inv]
        if w == "dw" or (w == "hunter" and "mh" in b["slots"]):
            return hands
        return [MH] if w == "shield" and MH in hands else []
    if inv in (15, 25, 26) and item["subclass"] != 19:
        return [RANGED] if w == "hunter" else []
    return []


# ---------------------------------------------------------------- runs (called from generate.py with its module as g)

WEAPON_TYPE = {0: 1, 1: 1, 4: 4, 5: 4, 6: 6, 7: 9, 8: 9, 10: 8, 13: 3, 15: 2}      # item subclass -> proto WeaponType
RANGED_TYPE = {2: 1, 3: 3, 18: 2, 16: 6, 19: 8}                                  # item subclass -> proto RangedWeaponType
TANK_AUTOCAST = {"action": {"autocastOtherCooldowns": {}}}


SLOT_INDEX = {"HEAD": 0, "BACK": 3, "CHEST": 4, "WRIST": 5, "HANDS": 6, "LEGS": 8, "FEET": 9}
PROVISIONAL_ENCHANTS = {8217: "the simulator has only an opt-in provisional model of Revelation"}
ENCHANT_BASE_ID = 1100000000          # synthetic item ids that carry an unimplemented enchant's effect


DEFENSIVE_STATS = {2, 26, 28, 29, 30, 31, 32, 34, 35, 36, 37, 38, 39, 40}


def defensive(specs):
    """Whether any effect could change damage taken (otherwise tank values use the threat term only)."""
    for spec in specs:
        for e in spec["effects"]:
            if e["kind"] == "stats" and DEFENSIVE_STATS & {int(k) for k in e["stats"]}:
                return True
            if e["kind"] not in ("damage", "dot", "extra_attacks", "armor_debuff", "haste", "stats"):
                return True
    return False


def spec_hash(specs):
    return hashlib.sha1(json.dumps(specs, sort_keys=True).encode()).hexdigest()[:8] if specs else "sim"


def enchant_slots(e, b):
    """Equipment slots to test enchant e in for profile b."""
    s, w = set(e["slots"]), b["weapons"]
    dw = w == "dw" or (w == "hunter" and "mh" in b["slots"])
    if s == {"OFFHAND"}:                                   # shield spikes
        return [OH] if w == "shield" else []
    if s & {"MAINHAND", "TWOHAND"}:
        if w == "2h":
            return [MH] if "TWOHAND" in s else []
        if w in ("dw", "shield") or dw:
            return [MH] + ([OH] if dw and "OFFHAND" in s else []) if "MAINHAND" in s else []
        return [MH] if w == "caster" and not e["melee_proc"] else []
    if "RANGED" in s:
        return [RANGED] if w == "hunter" else []
    return [SLOT_INDEX[x] for x in e["slots"] if x in SLOT_INDEX][:1]


class Plan:
    def __init__(self, g, items_path, build, enchants_path=None):
        self.g = g
        self.items = {it["id"]: it for it in json.load(open(items_path))}
        T = Tables(build)
        enchants = json.load(open(enchants_path)) if enchants_path else []
        ids = g.C.work / "itemfx_ids.json"
        ids.write_text(json.dumps(sorted(self.items) + sorted({e["enchant_id"] for e in enchants})))
        registered = json.loads(g.subprocess.run([str(g.C.hp), "effects", str(ids)], capture_output=True, text=True, check=True).stdout)
        self.entry, self.hash = {}, {}
        specs = []
        for iid, it in self.items.items():
            s, why = classify(it, T)
            sim = any(registered[str(iid)][:2])
            method = "simulated" if sim else ("estimated" if s else "not_modelled")
            notes = []
            text = " ".join(it.get("chance_on_hit", []) + it.get("equip_procs", []) + it.get("use", []))
            if "Static Electricity" in text:
                notes.append("Static Electricity stacking is not modelled")
            if re.search(r"times as much damage to|damage to (Beasts|Demons|Dragonkin|Elementals|Giants|Humanoids|Undead|Plants|Aquatic)", text):
                notes.append("creature-type bonus not modelled (benchmark target is Dragonkin)")
            only = re.search(r"(?:attacks|casts|hits) against ([A-Z][A-Za-z' ]+?) have a chance", text)
            if only and not sim:
                notes.append(f"only procs against {only.group(1)}: no value against the benchmark target")
                s = []
            if not sim:
                notes += sorted({x["assumed"] for x in s if x.get("assumed")})
                specs += s
            if method != "simulated" and why:
                notes.append("not modelled: " + "; ".join(sorted(set(why)))[:300])
            self.entry[iid] = {"method": method, "specs": s, "notes": notes, "sim": sim}
            self.hash[iid] = "sim" if sim else spec_hash(s)
        self.enchants = {}
        rows = {int(r["ID"]): r for r in spell_client.table(build, "SpellItemEnchantment")}
        for e in enchants:
            eid = e["enchant_id"]
            reg = registered[str(eid)]
            row = rows.get(eid, {})
            effs = [(int(num(row.get(f"Effect_{i}"))), int(num(row.get(f"EffectArg_{i}"))), num(row.get(f"EffectPointsMin_{i}"))) for i in range(3)]
            sim = reg[1] or reg[2]
            synth = ENCHANT_BASE_ID + eid
            s, why = [], []
            for kind, arg, pts in effs:
                if kind == 1 and arg:
                    eff, w = spell_effects(T, arg)
                    why += w
                    if eff:
                        spec = {"item": synth, "name": e["name"][:40], "trigger": "coh", "effects": eff}
                        if pts > 0:
                            spec["chance"] = pts / 100
                        else:
                            spec["ppm"], spec["assumed"] = ASSUMED_COH_PPM, f"{ASSUMED_COH_PPM} proc per minute"
                        s.append(spec)
                elif kind == 3 and arg:
                    o, w = classify({"id": synth, "name": e["name"], "class": 4, "item_effects": [{"trigger": 1, "spell": arg}]}, T)
                    s += o; why += w
                elif kind == 5:
                    why.append("stat enchant (scored as stats)")
            method = "simulated" if sim else ("estimated" if s else "not_modelled")
            notes = ([] if sim else sorted({x["assumed"] for x in s if x.get("assumed")}))
            if eid in PROVISIONAL_ENCHANTS:
                method, notes = "not_modelled", notes + [PROVISIONAL_ENCHANTS[eid]]
            if e.get("note") in ("skill", "movement speed", "attack power against a creature type") or (e.get("note") or "").startswith("aura 59"):
                notes.append("no combat value against the benchmark target (" + e["note"] + ")")
                method, s = "not_modelled", []
            if method != "simulated" and why:
                notes.append("not modelled: " + "; ".join(sorted(set(why)))[:300])
            self.enchants[eid] = {"spell": e["spell"], "name": e["name"], "slots": e["slots"], "method": method,
                                  "specs": s, "notes": notes, "sim": sim, "melee_proc": any(k == 1 for k, _, _ in effs)}
            self.hash[("enchant", eid)] = "sim" if sim else spec_hash(s)
            if not sim:
                specs += s
        self.spec_path = g.C.work / "itemfx_specs.json"
        self.spec_path.write_text(json.dumps(sorted(specs, key=lambda x: (x["item"], x["trigger"])), sort_keys=True))

    def tests(self, b):
        """(item, slot) pairs to run for build b. slot: an equipment index; trinket slot for non-weapon effects."""
        out = []
        for iid, e in self.entry.items():
            if e["method"] == "not_modelled":
                continue
            it = self.items[iid]
            specs = e["specs"]
            if any(s["trigger"] == "coh" for s in specs) or (e["sim"] and it["class"] == 2):
                out += [(iid, slot) for slot in weapon_slots(it, b)]
            elif any(relevant(b, s) for s in specs) or (e["sim"] and not specs):
                out.append((iid, TRINKET_SLOT))
        return out

    def enchant_tests(self, b):
        out = []
        eq = b["player"]["equipment"]["items"]
        for eid, e in self.enchants.items():
            if e["method"] == "not_modelled" or (e["specs"] and not any(relevant(b, s) for s in e["specs"])):
                continue
            out += [(eid, slot) for slot in enchant_slots(e, b) if len(eq) > slot and (eq[slot] or {}).get("id")]
        return out

    def request(self, b, wl, iid, slot, with_effect):
        g = self.g
        req = g.base_request(b, wl)
        req["simOptions"]["iterations"], req["simOptions"]["randomSeed"] = g.C.item_iterations, str(g.C.seed)
        p = req["raid"]["parties"][0]["players"][0]
        rot = p.get("rotation", {})
        it = self.items[iid]
        on_use = any(x["trigger"] == 0 for x in it["item_effects"])
        if b["tank"] and on_use and "autocastOtherCooldowns" not in json.dumps(rot):
            # Tank rotations do not use item cooldowns; both runs of an on-use item get automatic cooldown use.
            rot.setdefault("priorityList", []).insert(0, TANK_AUTOCAST)
        eq = p["equipment"]["items"]
        src = g.C.db.get((eq[slot] or {}).get("id")) or {"type": 12, "name": "Trinket", "stats": [0] * 44}
        keep = ("type", "weaponType", "handType", "rangedWeaponType", "weaponDamageMin", "weaponDamageMax", "weaponSpeed", "ilvl")
        entry = {k: src[k] for k in keep if k in src}
        entry["stats"] = [0] * len(src.get("stats") or [0] * 44)
        if slot in (MH, OH, RANGED) and it.get("speed"):
            dps = (src["weaponDamageMin"] + src["weaponDamageMax"]) / 2 / src["weaponSpeed"]
            spread = (src["weaponDamageMax"] - src["weaponDamageMin"]) / (src["weaponDamageMax"] + src["weaponDamageMin"])
            avg = dps * it["speed"]
            entry.update(weaponSpeed=it["speed"], weaponDamageMin=avg * (1 - spread), weaponDamageMax=avg * (1 + spread))
            if slot == RANGED and it["subclass"] in RANGED_TYPE:
                entry["rangedWeaponType"] = RANGED_TYPE[it["subclass"]]
            elif it["subclass"] in WEAPON_TYPE:
                entry["weaponType"] = WEAPON_TYPE[it["subclass"]]
        entry["id"] = iid if with_effect else 1300000000 + slot * 1000000 + iid % 1000000
        entry["name"] = it["name"] + ("" if with_effect else " (no effect)")
        p.setdefault("database", {}).setdefault("items", []).append(entry)
        eq[slot] = {"id": entry["id"]}
        return req

    def enchant_request(self, b, wl, eid, slot, with_effect):
        """The profile's own item in the slot, with the enchant or without any enchant. Unimplemented enchants ride on a
        copy of the item (same stats) whose synthetic id carries the enchant's effect."""
        g = self.g
        req = g.base_request(b, wl)
        req["simOptions"]["iterations"], req["simOptions"]["randomSeed"] = g.C.item_iterations, str(g.C.seed)
        p = req["raid"]["parties"][0]["players"][0]
        eq = p["equipment"]["items"]
        item = {k: v for k, v in (eq[slot] or {}).items() if k != "enchant"}
        e = self.enchants[eid]
        if e["sim"]:
            if with_effect:
                item["enchant"] = eid
        else:
            src = g.C.db[item["id"]]
            keep = ("type", "weaponType", "handType", "rangedWeaponType", "weaponDamageMin", "weaponDamageMax", "weaponSpeed", "ilvl")
            entry = {k: src[k] for k in keep if k in src}
            entry["stats"] = list(src.get("stats") or [0] * 44)
            entry["id"] = ENCHANT_BASE_ID + eid + (0 if with_effect else 50000000)
            entry["name"] = e["name"] + ("" if with_effect else " (no effect)")
            p.setdefault("database", {}).setdefault("items", []).append(entry)
            item = {"id": entry["id"]}
        eq[slot] = item
        return req

    def job(self, key, wl, iid, slot, with_effect, kind="item"):
        g = self.g
        b = g.C.builds[key]
        if kind == "enchant":
            tag = f"en__{key}__{wl}__{iid}__{slot}__{int(with_effect)}__{self.hash[('enchant', iid)]}"
            make = lambda: self.enchant_request(b, wl, iid, slot, with_effect)
        else:
            tag = f"fx__{key}__{wl}__{iid}__{slot}__{int(with_effect)}__{self.hash[iid]}"
            make = lambda: self.request(b, wl, iid, slot, with_effect)
        failed = g.C.work / "runs" / (tag + ".failed")
        if failed.exists():
            return None
        try:
            return g.metrics(g.run_request(tag, make(), extra=["-itemfx", str(self.spec_path)]))
        except g.subprocess.CalledProcessError as err:     # e.g. a class-specific simulator item on another class
            (g.C.work / "runs" / (tag + ".request.json")).unlink(missing_ok=True)
            failed.write_text((err.stderr or b"").decode(errors="replace")[:2000])
            return None

    def jobs(self):
        out = []
        for key, b in self.g.C.builds.items():
            for wl in self.g.workload_sets(b)[""]:
                for iid, slot in self.tests(b):
                    out += [(key, wl, iid, slot, True), (key, wl, iid, slot, False)]
                for eid, slot in self.enchant_tests(b):
                    out += [(key, wl, eid, slot, True, "enchant"), (key, wl, eid, slot, False, "enchant")]
        return out

    def export(self, R, by_class, profile_keys):
        """Fill by_class[cls]["item_effects"] and ["enchant_effects"]. R: job tuple -> metrics; profile_keys: build key ->
        [(profile key, lam, norm)]."""
        g = self.g

        def fill(out, key, b, slots, jkey, specs, use_only):
            wls = g.workload_sets(b)[""]
            offensive = bool(specs) and not defensive(specs)
            for slot in slots:
                a = {wl: R.get(jkey(wl, slot, True)) for wl in wls}
                ref = {wl: R.get(jkey(wl, slot, False)) for wl in wls}
                if any(v is None for v in (*a.values(), *ref.values())):
                    continue
                for pkey, lam, norm in profile_keys[key]:
                    # An offensive effect cannot change damage taken; its survival term would only be noise.
                    v = g.value(b, a, ref, 0.0 if offensive else lam) / norm
                    if v < -1:
                        msg = ("using it on cooldown lowers the score in the simulator: valued at 0" if use_only else
                               "the simulator scores this effect negative (for example a channel that replaces normal attacks)")
                        if msg not in out.get("note", ""):
                            out["note"] = (out.get("note", "") + "; " if out.get("note") else "") + msg
                    if use_only and v < 0:
                        v = 0.0
                    out[pkey + (":oh" if slot == OH and b["weapons"] != "shield" else "")] = round(v, 2)

        for key, b in g.C.builds.items():
            tests, etests = self.tests(b), self.enchant_tests(b)
            for iid, e in self.entry.items():
                out = by_class[b["cls"]].setdefault("item_effects", {}).setdefault(str(iid), {})
                out["method"] = e["method"]
                if e["notes"] and "note" not in out:
                    out["note"] = "; ".join(e["notes"])
                use_only = all(x["trigger"] == 0 for x in self.items[iid]["item_effects"])
                fill(out, key, b, [s for i, s in tests if i == iid], lambda wl, slot, w: (key, wl, iid, slot, w), e["specs"], use_only)
            for eid, e in self.enchants.items():
                out = by_class[b["cls"]].setdefault("enchant_effects", {}).setdefault(str(e["spell"]), {})
                out.update(enchant_id=eid, name=e["name"], method=e["method"])
                if e["notes"] and "note" not in out:
                    out["note"] = "; ".join(e["notes"])
                fill(out, key, b, [s for i, s in etests if i == eid], lambda wl, slot, w: (key, wl, eid, slot, w, "enchant"), e["specs"], False)
