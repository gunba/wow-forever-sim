#!/usr/bin/env python3
"""Level-60 stat weights and weapon-speed models for every benchmark build, exported per class.

    python3 tools/forever_weights/generate.py            # everything
    python3 tools/forever_weights/generate.py --classes warrior,paladin

Method (see README.md): central finite differences with common random numbers on the published benchmark
requests (DPS builds) and the benchmark tank workloads (tank builds); weapon speed from a speed x weapon-DPS
x attack-power grid fitted per speed. Runs are cached under --work, keyed by the engine binary, so a rerun
after an engine change recomputes everything and an interrupted run resumes.
"""
import argparse, collections, concurrent.futures, copy, hashlib, json, os, pathlib, re, subprocess, sys

import numpy as np

REPO = pathlib.Path(__file__).resolve().parents[2]
HERE = pathlib.Path(__file__).resolve().parent
RESULTS = REPO / "artifacts/modelled_gear/forever_dps_5min.json"
INPUTS = REPO / "artifacts/modelled_gear/forever_input_profiles.json"
DB_PATH = REPO / "assets/database/db.json"

RACE_ORDER = ["Tauren", "Orc", "Troll", "Undead", "Human", "Dwarf", "Night Elf", "Gnome", "Windshaper", "High Order"]
TREES = {"warrior": ["arms", "fury", "protection"], "paladin": ["holy", "protection", "retribution"],
         "hunter": ["beast_mastery", "marksmanship", "survival"], "rogue": ["assassination", "combat", "subtlety"],
         "priest": ["discipline", "holy", "shadow"], "shaman": ["elemental", "enhancement", "restoration"],
         "mage": ["arcane", "fire", "frost"], "warlock": ["affliction", "demonology", "destruction"],
         "druid": ["balance", "feral", "restoration"]}
TANKS = {"tank_warrior": "shield_tank", "protection_paladin": "protection_paladin_tank", "feral_tank_druid": "bear_tank",
         "fury_tank": "fury_tank"}
LAMBDAS = {"lambda_0_5": 0.5, "lambda_1": 1.0, "threat": 0.0}
WARRIOR_DPS_KEYS = {"fury": "fury_dps", "fury_2h": "fury_2h_dps", "arms": "arms_dps", "fury_sunder": "fury_sunder_dps"}
SHARED_WEAPON_MODEL = {"fury_sunder": "fury"}
NO_WINDFURY = {"fury", "fury_tank"}
SCHOOLS = {"mage": ["arcane", "fire", "frost"], "warlock": ["shadow", "fire"], "priest": ["shadow", "holy"],
           "druid": ["nature", "arcane"], "shaman": ["nature", "fire", "frost"], "paladin": ["holy"]}

# Stat probes: (bonusStats index list, pseudoStats index list, delta).
S = {"str": 0, "agi": 1, "sta": 2, "int": 3, "spi": 4, "spell_power": 5, "arcane_damage": 6, "fire_damage": 7, "frost_damage": 8,
     "holy_damage": 9, "nature_damage": 10, "shadow_damage": 11, "mp5": 12, "spell_hit_pct": 13, "spell_crit_pct": 14,
     "ap": 17, "hit_pct": 18, "crit_pct": 19, "expertise_pct": 22, "armor": 26, "rap": 27, "defense": 28, "block_pct": 29,
     "block_value": 30, "dodge_pct": 31, "parry_pct": 32, "healing": 41, "spell_damage": 42, "feral_ap": 43}
PROBE = {k: ([v], []) for k, v in S.items()}
PROBE["ap"] = ([17, 27], [])                  # item "+N Attack Power" raises melee and ranged attack power
PROBE["haste_pct"] = ([20, 15], [])           # "attack speed and casting speed"
PROBE["weapon_damage"] = ([], [30])           # BonusPhysicalDamage, as set by "+N Weapon Damage" items
DELTA = {"str": 50, "agi": 50, "sta": 75, "int": 50, "spi": 50, "mp5": 10, "ap": 100, "rap": 100, "feral_ap": 100,
         "hit_pct": 2, "crit_pct": 2, "haste_pct": 3, "expertise_pct": 1, "spell_power": 50, "spell_damage": 50, "healing": 50,
         "spell_hit_pct": 2, "spell_crit_pct": 2, "armor": 600, "defense": 15, "dodge_pct": 2, "parry_pct": 2, "block_pct": 2,
         "block_value": 30, "weapon_damage": 4, "mh_weapon_dps": 5, "oh_weapon_dps": 5, "two_hand_weapon_dps": 5,
         "ranged_weapon_dps": 5, **{f"{s}_damage": 50 for s in ("arcane", "fire", "frost", "holy", "nature", "shadow")}}
# Hit probed at fixed levels: below the 8% special cap, between the special and dual-wield white caps, below the spell cap.
LOW_SHIFT = {"hit_at4": ("hit_pct", 4.0), "hit_at17": ("hit_pct", 17.0), "spell_hit_at8": ("spell_hit_pct", 8.0)}
MELEE_SKILLS, RANGED_SKILLS = list(range(6, 16)), list(range(16, 20))
SLOT = {"mh": 14, "oh": 15, "2h": 14, "ranged": 16}
WEAPON_STAT = {"mh": "mh_weapon_dps", "oh": "oh_weapon_dps", "2h": "two_hand_weapon_dps", "ranged": "ranged_weapon_dps"}
SPEEDS = {"mh": [1.5, 1.8, 2.1, 2.4, 2.7, 3.0], "oh": [1.3, 1.5, 1.7, 2.0, 2.3, 2.6, 2.9], "2h": [2.6, 2.9, 3.2, 3.5, 3.8, 4.1],
          "ranged": [1.8, 2.1, 2.4, 2.7, 3.0, 3.3]}
MULTS, GRID_AP = [0.6, 1.0, 1.4], [0, 600]
WORKLOADS = {"single": [], "multi": ["-targets", "3", "-duration", "90", "-swing", "1500", "-healing", "2500"]}


class Ctx:
    pass


C = Ctx()


def sh(cmd, **kw):
    return subprocess.run(cmd, check=True, **kw)


def go_binary():
    import shutil
    for cand in (os.environ.get("GO"), shutil.which("go"), str(pathlib.Path.home() / ".local/share/go/bin/go")):
        if cand and pathlib.Path(cand).exists():
            return cand
    sys.exit("No Go toolchain: install one to ~/.local/share/go or set GO=/path/to/go")


def build_binaries(work):
    env = {**os.environ, "GOFLAGS": os.environ.get("GOFLAGS", "-mod=mod"), "GOPROXY": os.environ.get("GOPROXY", "off"),
           "GOTOOLCHAIN": os.environ.get("GOTOOLCHAIN", "local")}
    bins = work / "bin"; bins.mkdir(parents=True, exist_ok=True)
    for name, pkg in (("forever_tanks", "./tools/forever_tanks"), ("hpstat", "./tools/forever_weights/hpstat")):
        sh([go_binary(), "build", "-tags", "with_db", "-trimpath", "-buildvcs=false", "-o", str(bins / name), pkg], cwd=REPO, env=env)
    return bins / "forever_tanks", bins / "hpstat"


def talent_spec(cls, talents):
    pts = [sum(int(c) for c in part if c.isdigit()) for part in (talents.split("-") + ["", "", ""])[:3]]
    return TREES[cls][max(range(3), key=lambda i: pts[i])]


def classify(key, cls, player, db):
    eq = player["equipment"]["items"]
    item = lambda i: db.get((eq[i] if len(eq) > i else {}).get("id"), {})
    mh, oh, rg = item(14), item(15), item(16)
    if key in ("feral", "feral_tank_druid"):
        return "feral"
    if cls == "hunter":
        return "hunter"
    if cls in ("mage", "warlock", "priest") or key in ("balance", "elemental", "stormcaller"):
        return "caster"
    if mh.get("handType") == 4:
        return "2h"
    if oh.get("weaponType") == 7:
        return "shield"
    return "dw"


def stat_list(b):
    w, cls = b["weapons"], b["cls"]
    if w == "caster":
        s = ["int", "spi", "mp5", "spell_power", "spell_damage", "spell_hit_pct", "spell_hit_at8", "spell_crit_pct", "haste_pct"]
        s += [f"{x}_damage" for x in SCHOOLS.get(cls, [])]
        if b["key"] == "stormcaller":
            s += ["mh_weapon_dps", "str", "agi", "ap", "hit_pct", "crit_pct"]
        return s
    s = ["str", "agi", "ap", "hit_pct", "hit_at4", "crit_pct", "haste_pct", "expertise_pct", "weapon_damage"]
    if w == "dw" or b["key"] in ("survival", "pet_melee"):
        s += ["hit_at17"]
    if w != "feral":
        s += ["weapon_skill"]
    s += {"dw": ["mh_weapon_dps", "oh_weapon_dps"], "2h": ["two_hand_weapon_dps"], "shield": ["mh_weapon_dps"],
          "feral": ["two_hand_weapon_dps", "feral_ap"], "hunter": ["ranged_weapon_dps", "rap", "mh_weapon_dps", "oh_weapon_dps", "ranged_weapon_skill"]}[w]
    if cls in ("paladin", "shaman", "druid", "hunter"):
        s += ["int", "spi", "mp5", "spell_power"] + [f"{x}_damage" for x in SCHOOLS.get(cls, []) if cls != "hunter"]
        if cls != "hunter":
            s += ["spell_crit_pct", "spell_hit_pct"]
    if b["tank"]:
        s += ["sta", "armor", "defense", "dodge_pct", "parry_pct"] + (["block_pct", "block_value"] if w == "shield" else [])
    return list(dict.fromkeys(s))


def grid_slots(b):
    return {"dw": ["mh", "oh"], "2h": ["2h"], "shield": ["mh"], "hunter": ["ranged"] + (["mh", "oh"] if b["key"] in ("survival", "pet_melee") else []),
            "feral": [], "caster": []}[b["weapons"]]


def load_builds(classes):
    results = json.loads(RESULTS.read_text())["Results"]
    inputs = {(r["Key"], r["Race"]): r["BaselinePlayer"] for r in json.loads(INPUTS.read_text())["Results"]}
    db = {it["id"]: it for it in json.loads(DB_PATH.read_text())["items"]}
    races = collections.defaultdict(list); rows = {}
    for r in results:
        races[r["Key"]].append(r["Race"]); rows[(r["Key"], r["Race"])] = r
    builds = {}
    for key, rs in races.items():
        race = next((x for x in RACE_ORDER if x in rs), rs[0])
        row = rows[(key, race)]
        player = inputs[(key, race)] if key in TANKS else row["Request"]["raid"]["parties"][0]["players"][0]
        cls = player["class"].replace("Class", "").lower()
        builds[key] = dict(key=key, race=race, cls=cls, label=row["Build"], tank=key in TANKS, row=row, player=player,
                           spec=talent_spec(cls, player["talentsString"]), weapons=classify(key, cls, player, db))
    # Jordan's dual-wield Fury tank: benchmark Protection Warrior gear, Fury off-hand, his tank rotation.
    ft = json.loads((HERE / "fury_tank.json").read_text())
    p = copy.deepcopy(builds["tank_warrior"]["player"])
    fury = rows[("fury", builds["tank_warrior"]["race"])]["Request"]["raid"]["parties"][0]["players"][0]
    p["equipment"]["items"][15] = copy.deepcopy(fury["equipment"]["items"][15])
    p["talentsString"], p["rotation"] = ft["talents"], ft["rotation"]
    p["consumes"]["offHandImbue"] = p["consumes"].get("mainHandImbue", "ElementalSharpeningStone")
    p["tankWarrior"]["options"]["stance"] = ft["stance"]
    builds["fury_tank"] = dict(key="fury_tank", race=builds["tank_warrior"]["race"], cls="warrior", label="Fury tank (dual wield)",
                               tank=True, row=None, player=p, spec="fury", weapons="dw")
    for k, bb in builds.items():
        st = (bb["row"] or rows[("tank_warrior", bb["race"])])["Stats"]
        bb["ref_hit"], bb["ref_spell_hit"] = st[18], st[13]
    if classes:
        builds = {k: v for k, v in builds.items() if v["cls"] in classes}
    for b in builds.values():
        b["stats"] = stat_list(b)
        b["slots"] = grid_slots(b)
    return builds, db


# ---------------------------------------------------------------- requests

def tank_template(b, workload):
    out = C.work / "templates" / f"{b['key']}__{workload}.json"
    if not out.exists():
        out.parent.mkdir(parents=True, exist_ok=True)
        pf = out.with_suffix(".player.json"); pf.write_text(json.dumps(b["player"]))
        sh([str(C.bin), "-player", str(pf), "-iterations", "1", "-seed", "1", "-output", str(out), *WORKLOADS[workload]],
           stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL, env={**os.environ, "GOMAXPROCS": "1"})
        pf.unlink()
    return json.loads(out.read_text())["request"]


def base_request(b, workload):
    req = copy.deepcopy(tank_template(b, workload) if b["tank"] else b["row"]["Request"])
    return req


def pad(lst, n):
    lst = list(lst or [])
    return lst + [0.0] * (n - len(lst))


def bump(player, stats_idx=(), pseudo_idx=(), d=0.0):
    bs = player.setdefault("bonusStats", {})
    st, ps = pad(bs.get("stats"), 44), pad(bs.get("pseudoStats"), 40)
    for i in stats_idx:
        st[i] += d
    for i in pseudo_idx:
        ps[i] += d
    bs["stats"], bs["pseudoStats"] = st, ps


def weapon_variant(player, slot, db, *, dps_add=0.0, speed=None, mult=1.0, tag=""):
    it = db[player["equipment"]["items"][slot]["id"]]
    spd = it["weaponSpeed"]
    avg = (it["weaponDamageMin"] + it["weaponDamageMax"]) / 2
    spread = (it["weaponDamageMax"] - it["weaponDamageMin"]) / (it["weaponDamageMax"] + it["weaponDamageMin"])
    dps = avg / spd * mult + dps_add
    s = speed or spd
    new_avg = dps * s
    new_id = 1500000000 + int(hashlib.sha1(f"{it['id']}:{slot}:{tag}:{dps_add}:{s}:{mult}".encode()).hexdigest()[:8], 16) % 500000000
    entry = {"id": new_id, "name": it["name"] + " (probe)", "type": it["type"], "weaponDamageMin": new_avg * (1 - spread),
             "weaponDamageMax": new_avg * (1 + spread), "weaponSpeed": s, "stats": it["stats"], "ilvl": it.get("ilvl", 0)}
    for k in ("weaponType", "handType", "rangedWeaponType"):
        if it.get(k):
            entry[k] = it[k]
    player.setdefault("database", {}).setdefault("items", []).append(entry)
    player["equipment"]["items"][slot]["id"] = new_id


def apply_probe(b, req, stat, amount, db):
    p = req["raid"]["parties"][0]["players"][0]
    if stat in LOW_SHIFT:
        base, target = LOW_SHIFT[stat]
        ref = b["ref_spell_hit"] if base == "spell_hit_pct" else b["ref_hit"]
        bump(p, PROBE[base][0], PROBE[base][1], target - ref + amount)
    elif stat == "weapon_skill":
        bump(p, (), MELEE_SKILLS, amount)
    elif stat == "ranged_weapon_skill":
        bump(p, (), RANGED_SKILLS, amount)
    elif stat in ("mh_weapon_dps", "oh_weapon_dps", "two_hand_weapon_dps", "ranged_weapon_dps"):
        slot = {"mh_weapon_dps": 14, "oh_weapon_dps": 15, "two_hand_weapon_dps": 14, "ranged_weapon_dps": 16}[stat]
        weapon_variant(p, slot, db, dps_add=amount, tag=stat)
    else:
        idx, pidx = PROBE[stat]
        bump(p, idx, pidx, amount)


def run_request(tag, req):
    out = C.work / "runs" / (tag + ".json")
    if not out.exists():
        out.parent.mkdir(parents=True, exist_ok=True)
        rf = out.with_suffix(".request.json"); rf.write_text(json.dumps(req))
        sh([str(C.bin), "-request", str(rf), "-output", str(out)], stdout=subprocess.DEVNULL, stderr=subprocess.PIPE,
           env={**os.environ, "GOMAXPROCS": "1"})
        rf.unlink()
    return out


def metrics(out, want_stats=False):
    d = json.loads(out.read_text())
    m = d["result"]["raidMetrics"]["parties"][0]["players"][0]
    r = {"dps": m["dps"]["avg"], "tps": m.get("threat", {}).get("avg", 0), "dtps": m.get("dtps", {}).get("avg", 0),
         "death": m.get("chanceOfDeath", 0), "warnings": d.get("warnings")}
    if want_stats:
        r["stats"] = json.loads(subprocess.run([str(C.hp), str(out)], capture_output=True, text=True, check=True).stdout)
    return r


# ---------------------------------------------------------------- jobs

def weight_jobs(b):
    jobs = []
    wls = list(WORKLOADS) if b["tank"] else ["single"]
    for wl in wls:
        jobs.append((b["key"], wl, None, 0))
        for s in b["stats"]:
            if s in ("weapon_skill", "ranged_weapon_skill"):
                jobs += [(b["key"], wl, s, 5), (b["key"], wl, s, 10)]
            elif s in LOW_SHIFT:
                jobs += [(b["key"], wl, s, 1), (b["key"], wl, s, -1)]
            else:
                jobs += [(b["key"], wl, s, DELTA[s]), (b["key"], wl, s, -DELTA[s])]
    return jobs


def do_weight_job(job):
    key, wl, stat, amount = job
    b = C.builds[key]
    req = base_request(b, wl)
    req["simOptions"]["iterations"], req["simOptions"]["randomSeed"] = C.iterations, str(C.seed)
    if stat:
        apply_probe(b, req, stat, amount, C.db)
    tag = f"w__{key}__{wl}__{stat or 'base'}__{amount:+g}"
    out = run_request(tag, req)
    need_stats = b["tank"] or stat is None
    return job, metrics(out, want_stats=need_stats)


def grid_jobs(b, nowf=False):
    jobs = []
    if b["key"] in SHARED_WEAPON_MODEL:
        return jobs
    wls = list(WORKLOADS) if b["tank"] else ["single"]
    for slot in b["slots"]:
        if nowf and slot not in ("mh", "oh"):
            continue
        ref = ref_speed(b, slot)
        for s in sorted(set(SPEEDS[slot] + [ref])):
            for m in MULTS:
                for ap in GRID_AP:
                    for wl in wls:
                        jobs.append((b["key"], wl, slot, s, m, ap, nowf))
    return jobs


def ref_speed(b, slot):
    return C.db[b["player"]["equipment"]["items"][SLOT[slot]]["id"]]["weaponSpeed"]


def do_grid_job(job):
    key, wl, slot, s, m, ap, nowf = job
    b = C.builds[key]
    req = base_request(b, wl)
    req["simOptions"]["iterations"], req["simOptions"]["randomSeed"] = C.grid_iterations, str(C.seed)
    p = req["raid"]["parties"][0]["players"][0]
    weapon_variant(p, SLOT[slot], C.db, speed=s, mult=m, tag="grid")
    bump(p, [27] if slot == "ranged" else [17], [], ap)
    if nowf:
        req["raid"]["parties"][0].setdefault("buffs", {})["windfuryTotem"] = False
    tag = f"g__{key}__{wl}__{slot}__{s}__{m}__{ap}" + ("__nowf" if nowf else "")
    out = run_request(tag, req)
    want = s == ref_speed(b, slot) and m == 1.0
    return job, metrics(out, want_stats=want)


# ---------------------------------------------------------------- fitting

def value(b, a, ref, lam):
    """Strength-like value of result a relative to ref (DPS: DPS points; tanks: % threat + lam * % time-to-die)."""
    if not b["tank"]:
        return a["single"]["dps"] - ref["single"]["dps"]
    T = Sv = 0.0
    for wl in WORKLOADS:
        x, y = a[wl], ref[wl]
        T += 50 * (x["tps"] / y["tps"] - 1)
        # Health only changes with Stamina-type probes, which always read stats; speed-grid runs share the weapon's stats.
        hx = x["stats"][34] if x.get("stats") else (y["stats"][34] if y.get("stats") else 1.0)
        hy = y["stats"][34] if y.get("stats") else hx
        Sv += 50 * ((hx / x["dtps"]) / (hy / y["dtps"]) - 1)
    return T + lam * Sv


def fit_weights(b, W, lam, main=None):
    wls = list(WORKLOADS) if b["tank"] else ["single"]
    get = lambda stat, amount: {wl: W[(b["key"], wl, stat, amount)] for wl in wls}
    base = get(None, 0)
    raw = {}
    for s in b["stats"]:
        if s in ("weapon_skill", "ranged_weapon_skill"):
            v5, v10 = value(b, get(s, 5), base, lam) / 5, value(b, get(s, 10), get(s, 5), lam) / 5
            raw[s] = v5; raw[s + "_after5"] = v10
        elif s in LOW_SHIFT:
            raw[s] = (value(b, get(s, 1), base, lam) - value(b, get(s, -1), base, lam)) / 2
        else:
            d = DELTA[s]
            raw[s] = (value(b, get(s, d), base, lam) - value(b, get(s, -d), base, lam)) / (2 * d)
    if main is None:
        if b["weapons"] == "caster":
            main = "spell_power"
        elif b["cls"] == "warrior":
            main = "str"
        elif b["cls"] in ("rogue", "hunter"):
            main = max(("str", "agi"), key=lambda k: raw.get(k, 0))
        else:                                       # hybrids: Paladin, Shaman, Druid
            main = max(("str", "agi", "spell_power"), key=lambda k: raw.get(k, 0))
    norm = raw[main]
    w = {k: round(v / norm, 3) for k, v in raw.items()}
    return w, main, norm, base


def fit_grid(b, slot, G, norm, lam, nowf):
    wls = list(WORKLOADS) if b["tank"] else ["single"]
    s0 = ref_speed(b, slot)
    speeds = sorted(set(SPEEDS[slot] + [s0]))
    get = lambda s, m, ap: {wl: G[(b["key"], wl, slot, s, m, ap, nowf)] for wl in wls}
    D0 = None
    it = C.db[b["player"]["equipment"]["items"][SLOT[slot]]["id"]]
    D0 = (it["weaponDamageMin"] + it["weaponDamageMax"]) / 2 / it["weaponSpeed"]
    api = 27 if slot == "ranged" else 17
    AP0, AP1 = get(s0, 1.0, 0)["single"]["stats"][api], get(s0, 1.0, 600)["single"]["stats"][api]
    slopes = [value(b, get(s0, 1.4, ap), get(s0, 0.6, ap), lam) / norm / (0.8 * D0) for ap in GRID_AP]
    a = float(np.mean(slopes))
    rows, worst = {}, 0.0
    for s in speeds:
        X, y = [], []
        for m in MULTS:
            for ap, APt in zip(GRID_AP, (AP0, AP1)):
                X.append([m * D0, APt / 14, 1.0]); y.append(value(b, get(s, m, ap), get(s0, m, ap), lam) / norm)
        X, y = np.array(X), np.array(y)
        c, *_ = np.linalg.lstsq(X, y, rcond=None)
        worst = max(worst, float(np.abs(y - X @ c).max()))
        dv = c[0] * D0 + c[1] * AP0 / 14 + c[2]
        mult = round(float(1 + dv / (a * D0)), 3) if abs(a * D0) > 1e-6 else 1.0
        rows[f"{s:g}"] = {"h_weapon_dps": round(float(c[0]), 4), "h_ap14": round(float(c[1]), 4), "h_const": round(float(c[2]), 2),
                          "at_reference_str": round(float(dv), 1), "multiplier_at_reference": mult}
    return {"reference_speed": s0, "reference_weapon_dps": round(D0, 2), "reference_ap": round(AP0),
            "ap_kind": "ranged" if slot == "ranged" else "melee", "dps_weight_at_reference": round(a, 4),
            "dps_weight_by_ap": [round(x, 4) for x in slopes], "speed_terms": rows, "fit_max_residual": round(worst, 2)}


def talent_spell_hit(b):
    """Damage-weighted talent and school spell hit of the profile's spells, from its base run."""
    out = C.work / "runs" / f"w__{b['key']}__single__base__+0.json"
    bonus = json.loads(subprocess.run([str(C.hp), str(out), "spellhit"], capture_output=True, text=True, check=True).stdout)
    m = json.loads(out.read_text())["result"]["raidMetrics"]["parties"][0]["players"][0]
    dmg = collections.Counter()
    for a in m["actions"]:
        sid = a["id"].get("spellId")
        if sid:
            dmg[str(sid)] += sum(t.get("damage", 0) for t in a["targets"])
    spell_dmg = {k: v for k, v in dmg.items() if k in bonus and bonus[k]["school"] != 1 and v > 0}
    total = sum(spell_dmg.values())
    if not total:
        return 0.0, {}
    avg = sum(v * bonus[k]["bonus_hit_pct"] for k, v in spell_dmg.items()) / total
    top = sorted(spell_dmg.items(), key=lambda kv: -kv[1])[:5]
    return round(avg, 2), {k: {"share": round(v / total, 3), "talent_hit_pct": bonus[k]["bonus_hit_pct"]} for k, v in top}


def export(builds, W, G, out_dir):
    by_class = collections.defaultdict(lambda: {"profiles": {}, "weapon_model": {}})
    for key, b in builds.items():
        lams = LAMBDAS if b["tank"] else {"dps": None}
        spec_main = fit_weights(b, W, 0.0)[1]          # one main stat per build (tanks: chosen on threat)
        for lname, lam in lams.items():
            w, main, norm, base = fit_weights(b, W, lam if lam is not None else 0.0, spec_main)
            if b["tank"]:
                pkey = f"{TANKS[key]}_{lname}"
            else:
                pkey = WARRIOR_DPS_KEYS.get(key, f"{key}_dps")
            mkey = (pkey if b["tank"] else SHARED_WEAPON_MODEL.get(key, key)) if b["slots"] else None
            st = base["single"]["stats"]
            dw = b["weapons"] == "dw" or (b["weapons"] == "hunter" and "mh" in b["slots"])
            prof = {k: v for k, v in w.items() if k not in LOW_SHIFT}
            if b["weapons"] == "caster":
                th, top = talent_spell_hit(b)
                prof["spell_hit_model"] = {"cap_pct": 16, "cap_pct_plus2": 5, "below_cap": w["spell_hit_at8"],
                                           "at_reference": w["spell_hit_pct"], "reference_spell_hit_pct": round(st[13], 2),
                                           "talent_spell_hit_pct": th, "talent_spell_hit_top_spells": top,
                                           "note": "talent_spell_hit_pct = damage-weighted talent and school hit the cap already counts: gear hit "
                                                   "is worth below_cap per 1% until gear hit reaches cap_pct - talent_spell_hit_pct. "
                                                   "below_cap was measured at 8% gear hit plus talents."}
            else:
                prof["hit_model"] = {"special_cap_pct": 8, "below_special_cap": w["hit_at4"], "between_caps": w.get("hit_at17", 0.0) if dw else 0.0,
                                     "white_cap_pct": 27 if dw else 8, "reference_hit_pct": round(st[18], 2),
                                     "note": "measured at 4% and 17% hit"}
            meta = {"class": b["cls"].upper(), "spec": b["spec"], "role": "tank" if b["tank"] else "dps", "label": b["label"],
                    "weapons": b["weapons"], "weapon_model_key": mkey, "main_stat": main, "race": b["race"],
                    "reference": {"ap": round(st[17]), "rap": round(st[27]), "spell_power": round(st[5]), "melee_hit_pct": round(st[18], 2),
                                  "spell_hit_pct": round(st[13], 2), "health": round(st[34])}}
            if b["tank"]:
                meta["lambda"] = lam
                meta["label"] = b["label"] + {"lambda_0_5": ", threat + half survival", "lambda_1": ", threat + survival", "threat": ", threat only"}[lname]
            prof["meta"] = meta
            by_class[b["cls"]]["profiles"][pkey] = prof
            if b["slots"] and mkey == (pkey if b["tank"] else key):
                by_class[b["cls"]]["weapon_model"][mkey] = {s: fit_grid(b, s, G, norm, lam if lam is not None else 0.0, False) for s in b["slots"]}
                if key in NO_WINDFURY:
                    by_class[b["cls"]]["weapon_model"][mkey + "_nowf"] = {
                        s: fit_grid(b, s, G, norm, lam if lam is not None else 0.0, True) for s in b["slots"] if s in ("mh", "oh")}
    out_dir.mkdir(parents=True, exist_ok=True)
    for cls, data in by_class.items():
        doc = {"schema": 2, "class": cls.upper(), "source": f"wow-forever-sim {C.rev}, level 60, client build {C.build}, benchmark profiles "
               f"(modelled gear), {C.iterations} iterations per stat probe and {C.grid_iterations} per speed-grid point, common seed {C.seed}.",
               "units": "Weights are per point as items show it, normalised so the profile's meta.main_stat = 1 (Strength or Agility for "
                        "physical profiles, spell_power for casters). *_pct = per 1%. defense per Defense point, armor per armor point, "
                        "block_value per point, weapon_damage per '+N Weapon Damage', *_weapon_dps per point of weapon DPS at the reference "
                        "speed, weapon_skill per point for the first 5 points of skill in the wielded weapon types (weapon_skill_after5 after "
                        "that), ap = '+N Attack Power' (melee and ranged), rap = ranged attack power only, spell_power = damage and healing, "
                        "spell_damage = damage only, <school>_damage = that school only, haste_pct = attack and casting speed.",
               "tank_merge": "<tank>_lambda_X = % threat + X * % longer time-to-die (max health / damage taken per second), averaged over "
                             "one boss (300 s) and a three-attacker pack (90 s).",
               "hit_note": "Melee: each 1% of hit is worth hit_model.below_special_cap until special_cap_pct (8% vs +3, about 6% vs +2), "
                           "then between_caps (dual wield white swings) until white_cap_pct (27% vs +3, 25% vs +2), then 0. "
                           "Casters: spell_hit_model.below_cap per 1% until cap_pct (16% vs +3, 5% vs +2, including talent school hit), then 0.",
               "profiles": data["profiles"],
               "weapon_model": {"formula": "weapon_value = a * D + h_weapon_dps(s) * D + h_ap14(s) * AP / 14 + h_const(s) * AP / reference_ap. "
                                           "D = weapon DPS, s = speed, AP = current melee attack power (ranged attack power when ap_kind is "
                                           "ranged), a = dps_weight_at_reference. Interpolate between listed speeds, clamp outside.",
                                "notes": ["Look up a profile's model with meta.weapon_model_key. '<key>_nowf' = no Windfury Totem in the group.",
                                          "Feral forms use the weapon's DPS but not its speed: no speed model for feral profiles.",
                                          "Casters: weapon DPS is irrelevant (stat sticks).",
                                          "dps_weight_at_reference = 0 means the build does not use that weapon (e.g. melee Hunters' ranged slot)."],
                                "profiles": data["weapon_model"]}}
        (out_dir / f"{cls}.json").write_text(json.dumps(doc, indent=1) + "\n")
        print(f"wrote {out_dir / (cls + '.json')}: {len(data['profiles'])} profiles, {len(data['weapon_model'])} weapon models")


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument("--out", type=pathlib.Path, default=pathlib.Path.home() / ".cache/forever-sim-audit/statw/exports")
    ap.add_argument("--work", type=pathlib.Path, default=pathlib.Path.home() / ".cache/forever-weights")
    ap.add_argument("--classes", default="")
    ap.add_argument("--iterations", type=int, default=20000)
    ap.add_argument("--grid-iterations", type=int, default=10000)
    ap.add_argument("--seed", type=int, default=1795000001)
    ap.add_argument("--workers", type=int, default=max(1, (os.cpu_count() or 2) - 2))
    ap.add_argument("--smoke", action="store_true", help="tiny iteration counts into a separate cache, to check every build runs")
    args = ap.parse_args()
    bins = build_binaries(args.work)
    engine = hashlib.sha256(bins[0].read_bytes()).hexdigest()[:12]
    C.bin, C.hp = bins
    C.iterations, C.grid_iterations, C.seed = (100, 100, args.seed) if args.smoke else (args.iterations, args.grid_iterations, args.seed)
    C.work = args.work / (engine + ("-smoke" if args.smoke else ""))
    C.rev = subprocess.run(["git", "rev-parse", "--short", "HEAD"], cwd=REPO, capture_output=True, text=True).stdout.strip()
    C.build = (REPO / "assets/db_inputs/forever_client_build.txt").read_text().strip()
    classes = {c.strip().lower() for c in args.classes.split(",") if c.strip()}
    C.builds, C.db = load_builds(classes)
    for b in C.builds.values():                     # templates first (cheap, serial)
        if b["tank"]:
            for wl in WORKLOADS:
                tank_template(b, wl)
    wj = [j for b in C.builds.values() for j in weight_jobs(b)]
    gj = [j for b in C.builds.values() for j in grid_jobs(b)] + [j for b in C.builds.values() if b["key"] in NO_WINDFURY for j in grid_jobs(b, True)]
    print(f"{len(C.builds)} builds, {len(wj)} weight runs, {len(gj)} grid runs, engine {engine}", flush=True)
    W, G = {}, {}
    with concurrent.futures.ThreadPoolExecutor(args.workers) as ex:
        done = 0
        for job, m in ex.map(do_weight_job, wj):
            W[job] = m; done += 1
            if done % 100 == 0:
                print(f"  weights {done}/{len(wj)}", flush=True)
        done = 0
        for job, m in ex.map(do_grid_job, gj):
            G[job] = m; done += 1
            if done % 200 == 0:
                print(f"  grid {done}/{len(gj)}", flush=True)
    warn = {k: v["warnings"] for k, v in W.items() if v.get("warnings") and k[2] is None}
    for k, v in warn.items():
        print("APL warnings", k[0], v[:3])
    export(C.builds, W, G, args.out if not args.smoke else args.work / "smoke-exports")


if __name__ == "__main__":
    main()
