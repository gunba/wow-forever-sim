#!/usr/bin/env python3
"""Reproducible fixed-equipment tank talent and rotation comparisons."""

import argparse
import concurrent.futures
import copy
import hashlib
import itertools
import json
import os
from pathlib import Path
import subprocess

ROLES = {
    "tank_warrior": ("warrior", "Human", (0, 0, 31), "shieldSlam"),
    "protection_paladin": ("paladin", "Human", (0, 31, 0), "holyShield"),
    "feral_tank_druid": ("druid", "Tauren", (0, 31, 0), "berserk"),
}


def canonical(value):
    return json.dumps(value, sort_keys=True, separators=(",", ":"))


def signature(player):
    return hashlib.sha256(canonical(player).encode()).hexdigest()[:16]


def talents(key):
    return json.loads(Path(f"ui/core/talents/trees/{ROLES[key][0]}.json").read_text())


def decode(key, text):
    parts = text.split("-")
    return {node["fieldName"]: int(parts[t][i]) if t < len(parts) and i < len(parts[t]) else 0
            for t, tree in enumerate(talents(key)) for i, node in enumerate(tree["talents"])}


def encode(key, allocation):
    return "-".join("".join(str(allocation.get(node["fieldName"], 0))
                           for node in tree["talents"]).rstrip("0")
                    for tree in talents(key)).rstrip("-")


def legal(key, allocation):
    if sum(allocation.values()) != 51 or not allocation.get(ROLES[key][3]):
        return False
    for t, tree in enumerate(talents(key)):
        nodes = tree["talents"]
        if sum(allocation[n["fieldName"]] for n in nodes) < ROLES[key][2][t]:
            return False
        for node in nodes:
            points = allocation[node["fieldName"]]
            if not 0 <= points <= node.get("maxPoints", 1):
                return False
            if not points:
                continue
            row = node["location"]["rowIdx"]
            if sum(allocation[n["fieldName"]] for n in nodes
                   if n["location"]["rowIdx"] < row) < row * 5:
                return False
            parent = node.get("prereqLocation")
            if parent:
                source = next(n for n in nodes if n["location"] == parent)
                if allocation[source["fieldName"]] != source.get("maxPoints", 1):
                    return False
    return True


def neighbors(key, allocation):
    nodes = [n for tree in talents(key) for n in tree["talents"]]
    for source in nodes:
        field = source["fieldName"]
        if not allocation[field]:
            continue
        for target in nodes:
            dest = target["fieldName"]
            if dest == field or target.get("notSimulated"):
                continue
            candidate = dict(allocation)
            candidate[field] -= 1
            candidate[dest] += 1
            if legal(key, candidate):
                yield candidate


def talent_variants(key, player, only_neighbors=False):
    base = decode(key, player["talentsString"])
    yield "same-input", player
    allocations = list(neighbors(key, base))
    if not only_neighbors:
        if key == "tank_warrior":
            for cruelty, deflection, sunder, bloodrage, anticipation in itertools.product(
                    range(6), range(6), (0, 3), (0, 2), (3, 5)):
                candidate = dict(base)
                candidate.update(improvedDisarm=0, vanguard=0, concussionBlow=1,
                                 improvedShieldBash=0, cruelty=cruelty, deflection=deflection,
                                 improvedSunderArmor=sunder, improvedBloodrage=bloodrage,
                                 anticipation=anticipation)
                candidate["improvedHeroicStrike"] += 51 - sum(candidate.values())
                allocations.append(candidate)
        elif key == "protection_paladin":
            for strength, seals, deflection, judgement, conviction, keep_fury in itertools.product(
                    (0, 5), (0, 3), (1, 3, 5), (0, 2), (0, 3, 5), (0, 1)):
                candidate = dict(base)
                candidate.update(guardiansFavor=0, improvedSealOfFury=keep_fury,
                                 divineStrength=strength, improvedSeals=seals,
                                 deflection=deflection, improvedJudgement=judgement,
                                 conviction=conviction)
                candidate["benediction"] += 51 - sum(candidate.values())
                allocations.append(candidate)
        else:
            for utility, heart, shred, crit, rend in itertools.product(
                    (False, True), (0, 5), (0, 3), (0, 2), range(6)):
                candidate = dict(base)
                candidate.update(naturalShapeshifter=0, reflection=0, giftOfNature=0,
                                 heartOfTheWild=heart, shreddingAttacks=shred,
                                 predatoryInstincts=crit, rendAndTear=rend)
                if not utility:
                    candidate.update(brutalImpact=0, feralCharge=0)
                candidate["furor"] += 51 - sum(candidate.values())
                allocations.append(candidate)
    seen = {player["talentsString"]}
    for allocation in allocations:
        if not legal(key, allocation):
            continue
        text = encode(key, allocation)
        if text in seen:
            continue
        seen.add(text)
        candidate = copy.deepcopy(player)
        candidate["talentsString"] = text
        yield "talents-" + text, candidate


def cmp(name, op, value):
    return {"cmp": {"op": op, "lhs": {name: {}}, "rhs": {"const": {"val": str(value)}}}}


def cast(spell, condition=None, target=None):
    action = {"castSpell": {"spellId": {"spellId": spell}}}
    if target is not None:
        action["castSpell"]["target"] = {"type": "Target", "index": target}
    if condition:
        action["condition"] = condition
    return {"action": action}


def spell_id(row):
    return row.get("action", {}).get("castSpell", {}).get("spellId", {}).get("spellId", 0)


def apl_variants(key, player, guard_pass=False):
    yield "same-input", player
    base = player["rotation"]
    rows = base["priorityList"]
    if key == "tank_warrior":
        for shield_first, queue, sunder_dump, aoe_clap, shield_rage in itertools.product(
                (False, True), (20, 25, 30) if guard_pass else (20, 30, 45, 60),
                (True,) if guard_pass else (False, True), (False, True),
                (10, 15) if guard_pass else (20,)):
            new = copy.deepcopy(rows)
            revenge = next(i for i, row in enumerate(new) if spell_id(row) == 25288)
            slam = next(i for i, row in enumerate(new) if spell_id(row) == 23925)
            if shield_first:
                new[revenge], new[slam] = new[slam], new[revenge]
            for row in new:
                if spell_id(row) == 2565:
                    row["action"]["condition"]["and"]["vals"][0] = cmp("currentRage", "OpGe", shield_rage)
                if spell_id(row) in (25286, 20569):
                    row["action"]["condition"]["and"]["vals"][-1] = cmp("currentRage", "OpGe", queue)
            if not sunder_dump:
                new = [row for row in new if not (
                    spell_id(row) == 11597 and "cmp" in row["action"].get("condition", {})
                    and "currentRage" in row["action"]["condition"]["cmp"]["lhs"])]
            if aoe_clap:
                new.insert(min(revenge, slam), cast(11581, cmp("numberTargets", "OpGe", 3)))
            candidate = copy.deepcopy(player)
            candidate["rotation"]["priorityList"] = new
            yield f"shieldfirst{shield_first}-queue{queue}-sunder{sunder_dump}-clap{aoe_clap}-blockrage{shield_rage}", candidate
    elif key == "protection_paladin":
        defensive = copy.deepcopy(rows[:4])
        for strike_first, rank, reserve, judgement_first in itertools.product(
                (False, True), (26573, 20116, 20922, 20923, 20924), (0, 15, 30, 50), (False, True)):
            judgement, strike = cast(20271), cast(10333)
            attacks = [judgement, strike] if judgement_first else [strike, judgement]
            consecrate = cast(rank, cmp("currentManaPercent", "OpGe", f"{reserve}%"))
            if strike_first:
                new = defensive + attacks + [consecrate, cast(24239)]
            else:
                new = defensive + [consecrate] + attacks + [cast(24239)]
            candidate = copy.deepcopy(player)
            candidate["rotation"]["priorityList"] = new
            yield f"strikefirst{strike_first}-consecrate{rank}-reserve{reserve}-judge{judgement_first}", candidate
    else:
        for maul, swipe, bite_first, heal, spread in itertools.product(
                (15, 25, 35, 50), (False, True), (False, True), (30, 50), (False,)):
            new = copy.deepcopy(rows)
            # Compare explicit health-gated regeneration and proactive Barkskin
            # with the inherited cooldown-manager behavior.
            new[0:1] = [cast(22842, cmp("currentHealthPercent", "OpLt", f"{heal}%")),
                        cast(22812), rows[0]]
            for row in new:
                if spell_id(row) == 9881:
                    row["action"]["condition"] = cmp("currentRage", "OpGe", maul)
            if not swipe:
                new = [row for row in new if spell_id(row) != 9908]
            if not bite_first:
                index = next(i for i, row in enumerate(new) if spell_id(row) == 1238073)
                bite = new.pop(index)
                filler = max(i for i, row in enumerate(new) if spell_id(row) == 414644)
                new.insert(filler, bite)
            if spread:
                # Apply missing Lacerate stacks to extra attackers explicitly;
                # a fixed target cast does not reset the autoattack target.
                index = max(i for i, row in enumerate(new) if spell_id(row) == 414644)
                for target in (1, 2):
                    cond = {"and": {"vals": [
                        cmp("numberTargets", "OpGt", target),
                        {"cmp": {"op": "OpLe", "lhs": {"auraRemainingTime": {
                            "sourceUnit": {"type": "Target", "index": target},
                            "auraId": {"spellId": 414647}}},
                            "rhs": {"const": {"val": "4.5s"}}}},
                    ]}}
                    new.insert(index, cast(414644, cond, target))
            candidate = copy.deepcopy(player)
            candidate["rotation"]["priorityList"] = new
            yield f"maul{maul}-swipe{swipe}-bitefirst{bite_first}-heal{heal}-spread{spread}", candidate


def failures(candidate, reference):
    reasons = []
    for encounter in ("Single", "Multi"):
        c, r = candidate[encounter]["Tank"], reference[encounter]["Tank"]
        for field, limit in (("DTPS", 1.01), ("TMI", 1.02)):
            if c[field] > r[field] * limit:
                reasons.append(encounter + ":" + field)
        if c["ChanceOfDeath"] > r["ChanceOfDeath"] + .0025:
            reasons.append(encounter + ":death")
        for field, limit in (("TPS", .99), ("LeastTargetTPS", .98)):
            if c[field] < r[field] * limit:
                reasons.append(encounter + ":" + field)
        if candidate[encounter]["Warnings"]:
            reasons.append(encounter + ":warnings")
    return reasons


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--binary", type=Path, required=True)
    parser.add_argument("--input", type=Path, required=True)
    parser.add_argument("--phase", choices=("talents", "neighbors", "apl", "apl-guard", "confirm"), required=True)
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument("--iterations", type=int, default=200)
    parser.add_argument("--iterations-by-build", type=json.loads, default={},
                        help='JSON overrides, e.g. {"protection_paladin":100000}')
    parser.add_argument("--seed", type=int, default=20261031)
    parser.add_argument("--workers", type=int, default=min(24, os.cpu_count() or 1))
    parser.add_argument("--all-races", action="store_true")
    args = parser.parse_args()
    if any(key not in ROLES or not isinstance(n, int) or n < 1
           for key, n in args.iterations_by_build.items()):
        raise SystemExit("Invalid per-build iteration counts")
    args.binary = args.binary.resolve()
    args.output.mkdir(parents=True, exist_ok=True)
    rows = json.loads(args.input.read_text())["Results"]
    jobs, labels = {}, {}
    for row in rows:
        key, race = row["Key"], row["Race"]
        if key not in ROLES or not args.all_races and race != ROLES[key][1]:
            continue
        player = row["BaselinePlayer"]
        variants = ([("same-input", player)] if args.phase == "confirm" else
                    apl_variants(key, player, args.phase == "apl-guard") if args.phase.startswith("apl") else
                    talent_variants(key, player, args.phase == "neighbors"))
        controls = json.loads(Path(f"artifacts/tanks/{key}_1t_20261993_selected.json").read_text())
        control = controls["request"]["raid"]["parties"][0]["players"][0]
        control["race"] = player["race"]
        variants = itertools.chain([("published-control", control)], variants)
        for label, variant in variants:
            uid = key + "__" + race.lower().replace(" ", "_") + "__" + signature(variant)
            jobs[uid] = key, race, variant
            labels.setdefault(uid, []).append(label)
    manifest = {
        "phase": args.phase, "iterations": args.iterations, "seed": args.seed,
        **({"iterationsByBuild": args.iterations_by_build} if args.iterations_by_build else {}),
        "binarySHA256": hashlib.sha256(args.binary.read_bytes()).hexdigest(),
        "inputSHA256": hashlib.sha256(args.input.read_bytes()).hexdigest(),
        "jobs": {uid: {"build": key, "race": race, "labels": labels[uid], "player": player}
                 for uid, (key, race, player) in jobs.items()},
    }
    manifest_path = args.output / "manifest.json"
    if manifest_path.exists() and json.loads(manifest_path.read_text()) != manifest:
        raise SystemExit("Inputs changed; choose another output directory")
    manifest_path.write_text(json.dumps(manifest, indent=2) + "\n")

    def run(uid):
        key, race, player = jobs[uid]
        count = args.iterations_by_build.get(key, args.iterations)
        prefix = args.output / uid
        path = prefix.with_suffix(".player.json")
        path.write_text(json.dumps(player))
        result_path = prefix.with_suffix(".pair.json")
        if not result_path.exists():
            with prefix.with_suffix(".log").open("w") as log:
                subprocess.run([str(args.binary), "-build", key, "-race", race,
                                "-player", str(path), "-tank-pair", "-iterations", str(count),
                                "-seed", str(args.seed), "-output", str(prefix)],
                               env={**os.environ, "GOMAXPROCS": "1"}, stdout=log,
                               stderr=subprocess.STDOUT, check=True)
        return uid, json.loads(result_path.read_text())

    results = {}
    with concurrent.futures.ThreadPoolExecutor(max_workers=args.workers) as executor:
        for uid, pair in executor.map(run, jobs):
            results[uid] = pair
    controls = {(key, race): results[uid] for uid, (key, race, _) in jobs.items()
                if "published-control" in labels[uid]}
    ranked = []
    for uid, pair in results.items():
        key, race, player = jobs[uid]
        ranked.append({"id": uid, "build": key, "race": race, "labels": labels[uid],
                       "DPS": pair["Single"]["DPS"], "MultiDPS": pair["Multi"]["DPS"],
                       "rejectedFor": failures(pair, controls[key, race]),
                       "talents": player["talentsString"]})
    ranked.sort(key=lambda row: row["DPS"], reverse=True)
    (args.output / "ranking.json").write_text(json.dumps(ranked, indent=2) + "\n")
    for key in ROLES:
        top = [row for row in ranked if row["build"] == key and not row["rejectedFor"]][:8]
        print(key, json.dumps(top, indent=2), flush=True)
    print(f"{len(results)} paired scenarios completed", flush=True)


if __name__ == "__main__":
    main()
