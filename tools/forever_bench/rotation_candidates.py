"""Rotation variants derived from the frozen input profiles."""

from copy import deepcopy
from functools import lru_cache
from itertools import permutations
import json
from pathlib import Path


def balance(player):
    choices = []
    moonfire, insect_swarm, starfire, wrath = 9835, 24977, 25298, 9912
    for filler in ("eclipse", "starfire", "wrath"):
        for dots in ("both", "moonfire", "insect-swarm", "none"):
            p = deepcopy(player)
            keep = []
            for entry in p["rotation"]["priorityList"]:
                spell = entry.get("action", {}).get("castSpell", {}).get("spellId", {}).get("spellId")
                if spell == moonfire and dots not in ("both", "moonfire"):
                    continue
                if spell == insect_swarm and dots not in ("both", "insect-swarm"):
                    continue
                if filler != "eclipse" and spell in (starfire, wrath):
                    continue
                keep.append(entry)
            if filler != "eclipse":
                keep.append({"action": {"castSpell": {"spellId": {
                    "spellId": starfire if filler == "starfire" else wrath,
                }}}})
            p["rotation"]["priorityList"] = keep
            choices.append((f"{filler}-{dots}", p))
    for label, order in (
        ("eclipse-first", [starfire, insect_swarm, moonfire, wrath]),
        ("moonfire-first", [moonfire, insect_swarm, starfire, wrath]),
        ("eclipse-moonfire-first", [starfire, moonfire, insect_swarm, wrath]),
    ):
        p = deepcopy(player)

        def priority(entry):
            spell = entry.get("action", {}).get("castSpell", {}).get("spellId", {}).get("spellId")
            return order.index(spell) if spell in order else -1

        p["rotation"]["priorityList"].sort(key=priority)
        choices.append((label, p))
    return choices


def spell_id(entry):
    action = entry.get("action", {})
    return (action.get("castSpell") or action.get("channelSpell") or {}).get("spellId", {}).get("spellId")


def cast(spell, condition=None):
    action = {"castSpell": {"spellId": {"spellId": spell}}}
    if condition is not None:
        action["condition"] = condition
    return {"action": action}


def compare(value, op, constant):
    return {"cmp": {"op": op, "lhs": value, "rhs": {"const": {"val": str(constant)}}}}


def replace_spell(rotation, old, new):
    if isinstance(rotation, dict):
        if rotation.get("spellId") == old:
            rotation["spellId"] = new
            rotation.pop("rank", None)
        for value in rotation.values():
            replace_spell(value, old, new)
    elif isinstance(rotation, list):
        for value in rotation:
            replace_spell(value, old, new)


def feral(player):
    choices = []
    for rip_points in (4, 5):
        for bite_energy in (40, 60, 100, None):
            for builder in (9830, 9850):
                for rake in (False, True):
                    p = deepcopy(player)
                    if not any(spell_id(a) == 31018 for a in p["rotation"]["priorityList"]):
                        index = next(i for i, a in enumerate(p["rotation"]["priorityList"]) if spell_id(a) in (9830, 9850))
                        p["rotation"]["priorityList"].insert(index, cast(31018, {"and": {"vals": [
                            compare({"currentComboPoints": {}}, "OpGe", 5),
                            compare({"currentEnergy": {}}, "OpLe", 40),
                        ]}}))
                    actions = []
                    for entry in p["rotation"]["priorityList"]:
                        spell = spell_id(entry)
                        if spell == 9904:
                            continue
                        if spell == 9896:
                            entry["action"]["condition"]["and"]["vals"][0]["cmp"]["rhs"]["const"]["val"] = str(rip_points)
                        if spell == 31018:
                            if bite_energy is None:
                                continue
                            entry["action"]["condition"] = {"and": {"vals": [
                                compare({"currentComboPoints": {}}, "OpGe", 5),
                                compare({"currentEnergy": {}}, "OpLe", bite_energy),
                            ]}}
                        if spell in (9830, 9850):
                            if rake:
                                actions.append(cast(9904, {"and": {"vals": [
                                    {"not": {"val": {"dotIsActive": {"spellId": {"spellId": 9904}}}}},
                                    compare({"remainingTime": {}}, "OpGe", "9s"),
                                ]}}))
                            entry = cast(builder)
                        actions.append(entry)
                    p["rotation"]["priorityList"] = actions
                    choices.append((f"rip{rip_points}-bite{bite_energy}-builder{builder}-rake{int(rake)}", p))
    for seconds in (4, 6, 8):
        for points in (3, 4, 5):
            p = deepcopy(player)
            actions = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 31018]
            index = next(i for i, a in enumerate(actions) if spell_id(a) == 9896)
            actions.insert(index, cast(31018, {"and": {"vals": [
                compare({"remainingTime": {}}, "OpLe", f"{seconds}s"),
                compare({"currentComboPoints": {}}, "OpGe", points),
            ]}}))
            p["rotation"]["priorityList"] = actions
            choices.append((f"finish-bite-{seconds}s-{points}cp", p))
    return choices


def shaman_caster(player):
    choices = []
    bolt_ids = {4: 915, 6: 6041, 8: 10392, 10: 15208}
    for reserve in (10, 35, 60):
        for rank, spell in bolt_ids.items():
            p = deepcopy(player)
            actions = []
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) in bolt_ids.values():
                    continue
                actions.append(entry)
            condition = {"cmp": {
                "op": "OpGe", "lhs": {"currentMana": {}},
                "rhs": {"math": {"op": "OpMul", "lhs": {"remainingTime": {}},
                                 "rhs": {"const": {"val": str(reserve)}}}},
            }}
            actions.extend([cast(15208, condition), cast(spell)])
            p["rotation"]["priorityList"] = actions
            choices.append((f"reserve{reserve}-bolt{rank}", p))
    for label, spell in (("no-chain-lightning", 10605), ("no-flame-shock", 29228), ("no-searing-totem", 10438)):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != spell]
        choices.append((label, p))
    for rank, spell in ((1, 8050), (3, 8053), (5, 10448)):
        p = deepcopy(player)
        replace_spell(p["rotation"], 29228, spell)
        choices.append((f"flame-shock-{rank}", p))
    p = deepcopy(player)
    for entry in p["rotation"]["priorityList"]:
        if spell_id(entry) == 10605:
            entry["action"].pop("condition", None)
    choices.append(("chain-on-cooldown", p))
    p = deepcopy(player)
    p["rotation"]["priorityList"].sort(key=lambda a: 0 if "autocastOtherCooldowns" in a["action"] else 1)
    choices.append(("cooldowns-first", p))
    if any(spell_id(a) == 1238300 for a in player["rotation"]["priorityList"]):
        p = deepcopy(player)
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) == 29228:
                entry["action"]["condition"] = compare({"dotRemainingTime": {"spellId": {"spellId": 29228}}}, "OpLe", "2s")
        choices.append(("refresh-flame-before-lava", p))
    return choices


def enhancement(player):
    choices = []
    for stacks in (3, 4, 5):
        for chain in (False, True):
            for stormstrike_first in (False, True):
                p = deepcopy(player)
                actions = []
                stormstrike = None
                for entry in p["rotation"]["priorityList"]:
                    if spell_id(entry) == 15208:
                        condition = entry["action"]["condition"]
                        condition["cmp"]["rhs"]["const"]["val"] = str(stacks)
                        if chain:
                            actions.append(cast(10605, deepcopy(condition)))
                    if spell_id(entry) == 17364 and stormstrike_first:
                        stormstrike = entry
                        continue
                    actions.append(entry)
                if stormstrike is not None:
                    index = next(i for i, a in enumerate(actions) if spell_id(a) in (15208, 10605))
                    actions.insert(index, stormstrike)
                p["rotation"]["priorityList"] = actions
                choices.append((f"maelstrom{stacks}-chain{int(chain)}-strikefirst{int(stormstrike_first)}", p))
    for label, ids in (
        ("no-earth-shock", {10414}),
        ("no-flame-shock", {29228}),
        ("no-shocks", {10414, 29228}),
        ("no-searing", {10438}),
        ("external-utility-totems", {10442, 10627}),
    ):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) not in ids]
        p["rotation"]["prepullActions"] = [a for a in p["rotation"].get("prepullActions", []) if spell_id(a) not in ids]
        choices.append((label, p))
    for label, old, new in (("earth-shock-1", 10414, 8042), ("flame-shock-1", 29228, 8050)):
        p = deepcopy(player)
        replace_spell(p["rotation"], old, new)
        choices.append((label, p))
    return choices


def ranged_hunter(player):
    choices = []
    is_bm = any(spell_id(a) == 1293527 for a in player["rotation"]["priorityList"])
    aimed_ids = (19434, 20900, 20901, 20902, 20903, 20904)
    for shot in (None, 2643, *aimed_ids):
        for gap in ("0s", "1s", "1.5s", "2s"):
            if shot is None and gap != "0s":
                continue
            p = deepcopy(player)
            actions = []
            inserted = False
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) in (*aimed_ids, 2643):
                    if not inserted and shot is not None:
                        actions.append(cast(shot, compare({"autoTimeToNext": {"autoType": "Ranged"}}, "OpGe", gap)))
                        inserted = True
                    continue
                actions.append(entry)
            p["rotation"]["priorityList"] = actions
            choices.append((f"primary{shot}-gap{gap}", p))
    for label, ids in (
        ("no-arcane", {14287}), ("no-serpent", {25295}),
        ("no-arcane-or-serpent", {14287, 25295}),
    ):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) not in ids]
        choices.append((label, p))
    for threshold in ("10%", "15%", "20%", "25%", "50%", "75%"):
        p = deepcopy(player)
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) == 14287:
                entry["action"]["condition"] = compare({"currentManaPercent": {}}, "OpGe", threshold)
        choices.append((f"arcane-mana{threshold}", p))
    for threshold in ("15%", "25%"):
        for seconds in ("10s", "20s", "30s"):
            p = deepcopy(player)
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) == 14287:
                    entry["action"]["condition"] = {"or": {"vals": [
                        compare({"currentManaPercent": {}}, "OpGe", threshold),
                        compare({"remainingTime": {}}, "OpLe", seconds),
                    ]}}
            choices.append((f"arcane-mana{threshold}-finish{seconds}", p))
    for gap in ("0s", ".5s", "1.5s"):
        p = deepcopy(player)
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) == 14287:
                entry["action"]["condition"] = compare({"autoTimeToNext": {"autoType": "Ranged"}}, "OpGe", gap)
        choices.append((f"arcane-gap{gap}", p))
    p = deepcopy(player)
    for entry in p["rotation"]["priorityList"]:
        if spell_id(entry) == 25295:
            entry["action"]["condition"] = {"and": {"vals": [
                entry["action"]["condition"],
                compare({"remainingTime": {}}, "OpGe", "12s"),
            ]}}
    choices.append(("serpent-duration", p))
    if is_bm:
        for threshold in ("0%", "50%", "95%"):
            p = deepcopy(player)
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) == 19577:
                    entry["action"]["condition"] = compare({"currentManaPercent": {}}, "OpGt", threshold)
            choices.append((f"intimidation-mana{threshold}", p))
    else:
        for gap in ("0s", "2s", "3s"):
            p = deepcopy(player)
            p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 1310786]
            p["rotation"]["priorityList"].insert(1, cast(
                1310786, compare({"autoTimeToNext": {"autoType": "Ranged"}}, "OpGe", gap),
            ))
            choices.append((f"sniper-gap{gap}", p))
        for mana in ("15%", "25%", "40%", "60%"):
            p = deepcopy(player)
            p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 1310786]
            p["rotation"]["priorityList"].insert(1, cast(
                1310786, compare({"currentManaPercent": {}}, "OpGe", mana),
            ))
            choices.append((f"sniper-mana{mana}", p))
    return choices


def survival(player):
    choices = []
    originals = {spell_id(a): a for a in player["rotation"]["priorityList"] if spell_id(a) in (1317257, 14271, 14305)}
    for order in permutations(originals):
        p = deepcopy(player)
        fixed = [a for a in p["rotation"]["priorityList"] if spell_id(a) not in originals]
        p["rotation"]["priorityList"] = fixed + [deepcopy(originals[id]) for id in order]
        choices.append(("priority-" + "-".join(map(str, order)), p))
    for label, spell in (("no-trap", 14305), ("no-mongoose", 14271)):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != spell]
        choices.append((label, p))
    p = deepcopy(player)
    replace_spell(p["rotation"], 14305, 14317)
    choices.append(("explosive-trap", p))
    for seconds in ("6s", "12s", "15s"):
        p = deepcopy(player)
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) == 14305:
                entry["action"]["condition"] = compare({"remainingTime": {}}, "OpGe", seconds)
        choices.append((f"trap-duration{seconds}", p))
    return choices


def arcane(player):
    choices = []
    for stacks in (1, 2, 3, 4):
        for reset_spell in (25304, 25306, 1237313, 25345):
            for mana in ("20%", "40%"):
                p = deepcopy(player)
                for entry in p["rotation"]["priorityList"]:
                    conditions = entry["action"].get("condition", {}).get("or", {}).get("vals", [])
                    if any(v.get("cmp", {}).get("lhs", {}).get("auraNumStacks", {}).get("auraId", {}).get("spellId") == 30451 for v in conditions):
                        entry["action"]["castSpell"]["spellId"] = {"spellId": reset_spell}
                        entry["action"]["condition"] = {"or": {"vals": [
                            compare({"auraNumStacks": {"auraId": {"spellId": 30451}}}, "OpGe", stacks),
                            compare({"currentManaPercent": {}}, "OpLt", mana),
                        ]}}
                choices.append((f"blast{stacks}-reset{reset_spell}-mana{mana}", p))
    for filler in (25304, 1237313, 25345):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [
            a for a in p["rotation"]["priorityList"] if spell_id(a) not in (30451, 25304, 25345, 25306, 1237313)
        ] + [cast(filler)]
        choices.append((f"filler-only-{filler}", p))
    return choices


def fire(player):
    choices = []
    fillers = (25306, 10149, 10150, 1237313, 1237312)
    for reserve in (0, 30, 60, 90, 110, 150):
        for scorch in (2948, 10207):
            p = deepcopy(player)
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) in fillers:
                    entry["action"]["castSpell"]["spellId"] = {"spellId": 25306}
                    entry["action"]["condition"]["cmp"]["rhs"]["math"]["rhs"]["const"]["val"] = str(reserve)
                if spell_id(entry) in (2948, 10207) and "condition" not in entry["action"]:
                    entry["action"]["castSpell"]["spellId"] = {"spellId": scorch}
            choices.append((f"fireball-reserve{reserve}-scorch{scorch}", p))
    for label, omit in (("no-fire-blast", 10199), ("no-hot-streak-pyro", 18809)):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != omit]
        choices.append((label, p))
    for filler in (10149, 10150, 1237313, 1237312):
        for reserve in (0, 60, 110):
            p = deepcopy(player)
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) in fillers:
                    entry["action"]["castSpell"]["spellId"] = {"spellId": filler}
                    entry["action"]["condition"]["cmp"]["rhs"]["math"]["rhs"]["const"]["val"] = str(reserve)
            choices.append((f"fire-filler{filler}-reserve{reserve}", p))
    return choices


def frost(player):
    choices = []
    for filler in (25304, 1237313, 1237312, 10181, 10179):
        for lance in ("proc", "two-stacks", "never"):
            p = deepcopy(player)
            if lance != "never" and not any(spell_id(a) == 30455 for a in p["rotation"]["priorityList"]):
                p["rotation"]["priorityList"].insert(-1, cast(
                    30455, {"auraIsActive": {"auraId": {"spellId": 44543}}},
                ))
            actions = []
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) == 30455:
                    if lance == "never":
                        continue
                    if lance == "two-stacks":
                        entry["action"]["condition"] = compare({"auraNumStacks": {"auraId": {"spellId": 44543}}}, "OpGe", 2)
                    else:
                        entry["action"]["condition"] = {"auraIsActive": {"auraId": {"spellId": 44543}}}
                if spell_id(entry) in (25304, 1237313, 1237312, 10181, 10179):
                    entry["action"]["castSpell"]["spellId"] = {"spellId": filler}
                actions.append(entry)
            p["rotation"]["priorityList"] = actions
            choices.append((f"frost-filler{filler}-lance{lance}", p))
    return choices


def rogue(player):
    choices = []
    base = deepcopy(player)
    actions = base["rotation"]["priorityList"]
    # The two-second-tick Tea/Adrenaline Rush gates do not describe smooth
    # Energy. Tea is already the configured conjured consumable in these inputs.
    actions[:] = [a for a in actions if a["action"].get("castSpell", {}).get("spellId", {}).get("itemId") != 7676]
    tea = {"action": {
        "condition": compare({"currentEnergy": {}}, "OpLe", 10),
        "castSpell": {"spellId": {"itemId": 7676}},
    }}
    actions.insert(0, tea)
    for entry in actions:
        if spell_id(entry) == 13750:
            entry["action"]["condition"] = compare({"currentEnergy": {}}, "OpLe", 40)
    choices.append(("energy-cooldowns", deepcopy(base)))

    builder = next(spell_id(a) for a in reversed(actions) if spell_id(a) in (11294, 1241584, 16511))
    for cp in (3, 4, 5):
        for gate_builder in (False, True):
            p = deepcopy(base)
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) == 31016:
                    entry["action"]["condition"] = compare({"currentComboPoints": {}}, "OpGe", cp)
                if spell_id(entry) == builder and not gate_builder:
                    entry["action"].pop("condition", None)
            choices.append((f"eviscerate{cp}-buildergate{int(gate_builder)}", p))
    for cp in (1, 2, 3, 5):
        for refresh in ("0s", "2s", "3s"):
            p = deepcopy(base)
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) == 6774:
                    entry["action"]["condition"] = {"and": {"vals": [
                        compare({"currentComboPoints": {}}, "OpGe", cp),
                        {"or": {"vals": [
                            {"not": {"val": {"auraIsActive": {"auraId": {"spellId": 6774}}}}},
                            compare({"auraRemainingTime": {"auraId": {"spellId": 6774}}}, "OpLe", refresh),
                        ]}},
                        compare({"remainingTime": {}}, "OpGe", "6s"),
                    ]}}
            choices.append((f"slice{cp}-refresh{refresh}", p))
    for cp in (4, 5):
        p = deepcopy(base)
        entries = p["rotation"]["priorityList"]
        entries[:] = [a for a in entries if spell_id(a) != 11275]
        index = next(i for i, a in enumerate(entries) if spell_id(a) == 31016)
        entries.insert(index, cast(11275, {"and": {"vals": [
            compare({"currentComboPoints": {}}, "OpGe", cp),
            {"not": {"val": {"dotIsActive": {"spellId": {"spellId": 11275}}}}},
            compare({"remainingTime": {}}, "OpGt", "12s"),
        ]}}))
        choices.append((f"rupture{cp}", p))
    p = deepcopy(base)
    p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 11275]
    choices.append(("no-rupture", p))
    if builder == 1241584:
        for cp in (1, 3, 4, 5):
            for refresh in ("0s", "2s"):
                p = deepcopy(base)
                entries = p["rotation"]["priorityList"]
                entries[:] = [a for a in entries if spell_id(a) != 1310703]
                index = next(i for i, a in enumerate(entries) if spell_id(a) == 31016)
                entries.insert(index, cast(1310703, {"and": {"vals": [
                    compare({"currentComboPoints": {}}, "OpGe", cp),
                    compare({"auraRemainingTime": {"auraId": {"spellId": 1310703}}}, "OpLe", refresh),
                    compare({"remainingTime": {}}, "OpGt", "9s"),
                ]}}))
                choices.append((f"venom{cp}-refresh{refresh}", p))
    if builder in (11294, 1241584):
        for cp in (0, 3, 4, 5):
            p = deepcopy(base)
            entries = p["rotation"]["priorityList"]
            entries[:] = [a for a in entries if spell_id(a) != 14177]
            index = next(i for i, a in enumerate(entries) if spell_id(a) == 31016)
            entries.insert(index, cast(14177, compare({"currentComboPoints": {}}, "OpGe", cp)))
            choices.append((f"cold-blood{cp}", p))
    if builder == 16511:
        for before_finishers in (False, True):
            p = deepcopy(base)
            entries = p["rotation"]["priorityList"]
            entries[:] = [a for a in entries if spell_id(a) != 14278]
            before = (11275, 31016) if before_finishers else (16511,)
            index = next(i for i, a in enumerate(entries) if spell_id(a) in before)
            entries.insert(index, cast(14278, compare({"currentComboPoints": {}}, "OpLt", 4)))
            choices.append((f"ghostly-strike-early{int(before_finishers)}", p))
        p = deepcopy(base)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) not in (1856, 14183, 11269)]
        choices.append(("no-stealth-cycle", p))
    return choices


def retribution(player):
    choices = []
    for twist_floor in (0, 5, 15, 30, 50):
        for consecration_floor in (None, 0, 20, 40):
            p = deepcopy(player)
            if consecration_floor is None:
                p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 20924]
            elif not any(spell_id(a) == 20924 for a in p["rotation"]["priorityList"]):
                p["rotation"]["priorityList"].append(cast(20924))
            for entry in p["rotation"]["priorityList"]:
                spell = spell_id(entry)
                condition = entry["action"].get("condition", {})
                if spell in (20293, 20920) and "and" in condition:
                    condition["and"]["vals"][0]["cmp"]["rhs"]["const"]["val"] = f"{twist_floor}%"
                if spell == 20924:
                    entry["action"]["condition"] = compare({"currentManaPercent": {}}, "OpGe", f"{consecration_floor}%")
            choices.append((f"twist{twist_floor}-consecration{consecration_floor}", p))
    for label, omit in (("no-twist", {20293}), ("no-consecration", {20924}), ("no-exorcism", {10333})):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [
            a for a in p["rotation"]["priorityList"]
            if spell_id(a) not in omit and not (
                label == "no-twist" and spell_id(a) == 20920 and "and" in a["action"].get("condition", {})
            )
        ]
        choices.append((label, p))
    p = deepcopy(player)
    replace_spell(p["rotation"], 20293, 20154)
    choices.append(("twist-righteousness1", p))
    for first in (10333, 24239, 20924):
        p = deepcopy(player)
        entries = p["rotation"]["priorityList"]
        entry = next((a for a in entries if spell_id(a) == first), None)
        if entry is None:
            if first != 20924:
                continue
            entry = cast(20924, compare({"currentManaPercent": {}}, "OpGe", "20%"))
        else:
            entries.remove(entry)
        entries.insert(1, entry)
        choices.append((f"priority{first}", p))
    return choices


def shadow(player):
    choices = []
    for inner_focus in (19280, 10947):
        p = deepcopy(player)
        entries = p["rotation"]["priorityList"]
        sequence = next(a for a in entries if "strictSequence" in a["action"])
        entries.remove(sequence)
        index = next(i for i, a in enumerate(entries) if spell_id(a) == inner_focus)
        sequence["action"]["strictSequence"]["actions"][1] = cast(inner_focus)["action"]
        sequence["action"]["condition"] = deepcopy(entries[index]["action"].get("condition", {}))
        entries.insert(index, sequence)
        choices.append((f"inner-focus{inner_focus}", p))
    for spell in (19280, 10894, 10947):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"]
            if spell_id(a) != spell and not (spell == 19280 and "strictSequence" in a["action"])]
        choices.append((f"omit{spell}", p))
    for seconds in (6, 18, 24):
        p = deepcopy(player)
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) == 10894:
                entry["action"]["condition"]["and"]["vals"][1] = compare({"remainingTime": {}}, "OpGe", f"{seconds}s")
        choices.append((f"pain-end{seconds}s", p))
    for dots in (False, True):
        p = deepcopy(player)
        checks = [{"spellCanCast": {"spellId": {"spellId": 10947}}}]
        if dots:
            for spell in (10894, 19280):
                ready = [
                    {"not": {"val": {"dotIsActive": {"spellId": {"spellId": spell}}}}},
                    {"spellCanCast": {"spellId": {"spellId": spell}}},
                ]
                if spell == 10894:
                    ready.append(compare({"remainingTime": {}}, "OpGe", "10s"))
                checks.append({"and": {"vals": ready}})
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) == 18807:
                entry["action"] = {"channelSpell": {
                    "spellId": {"spellId": 18807},
                    "interruptIf": {"or": {"vals": checks}},
                    "instantInterrupt": False,
                }}
        choices.append((f"mind-flay-interrupt-dots{int(dots)}", p))
    return choices


def smite(player):
    choices = []
    for reserve in (0, 20, 40, 60, 80, 100):
        for rank, spell in ((3, 598), (4, 984), (6, 6060)):
            p = deepcopy(player)
            entries = [a for a in p["rotation"]["priorityList"] if spell_id(a) not in (598, 984, 6060, 10934)]
            entries.extend([
                cast(10934, {"cmp": {"op": "OpGe", "lhs": {"currentMana": {}},
                    "rhs": {"math": {"op": "OpMul", "lhs": {"remainingTime": {}},
                                    "rhs": {"const": {"val": str(reserve)}}}}}}),
                cast(spell),
            ])
            p["rotation"]["priorityList"] = entries
            choices.append((f"reserve{reserve}-smite{rank}", p))
    for spell in (15261, 10894, 1316995):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != spell]
        choices.append((f"omit{spell}", p))
    for refresh in (0, 1, 2, 3):
        p = deepcopy(player)
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) == 15261:
                entry["action"]["condition"] = compare(
                    {"dotRemainingTime": {"spellId": {"spellId": 15261}}}, "OpLe", f"{refresh}s")
        choices.append((f"holy-fire-refresh{refresh}s", p))
    p = deepcopy(player)
    entries = p["rotation"]["priorityList"]
    penance = next(a for a in entries if spell_id(a) == 1316995)
    entries.remove(penance)
    entries.insert(1, penance)
    choices.append(("penance-first", p))
    p = deepcopy(player)
    sequence = next(a for a in p["rotation"]["priorityList"] if "strictSequence" in a["action"])
    sequence["action"]["strictSequence"]["actions"][1] = cast(1316995)["action"]
    choices.append(("inner-focus-penance", p))
    return choices


@lru_cache
def talent_trees(class_name):
    return json.loads((Path(__file__).resolve().parents[2] /
                       f"ui/core/talents/trees/{class_name}.json").read_text())


def talent_rank(player, class_name, field):
    parts = player["talentsString"].split("-")
    parts += [""] * (3 - len(parts))
    for tree, text in zip(talent_trees(class_name), parts):
        for i, talent in enumerate(tree["talents"]):
            if talent["fieldName"] == field:
                return int(text[i]) if i < len(text) else 0
    raise KeyError(field)


def dot_maintenance(spell):
    return cast(spell, {"cmp": {
        "op": "OpLe",
        "lhs": {"dotRemainingTime": {"spellId": {"spellId": spell}}},
        "rhs": {"spellCastTime": {"spellId": {"spellId": spell}}},
    }})


def warlock(player):
    choices = []
    filler = 1293813 if talent_rank(player, "warlock", "incinerate") else 25307
    filler_ids = (25307, 1293813, 11700, 11704)
    dots = [25311, 25309]
    if talent_rank(player, "warlock", "siphonLife"):
        dots.append(18881)
    for mask in range(1 << len(dots)):
        p = deepcopy(player)
        entries = [a for a in p["rotation"]["priorityList"] if spell_id(a) not in dots]
        index = next(i for i, a in enumerate(entries) if spell_id(a) in filler_ids)
        entries[index:index] = [dot_maintenance(spell) for i, spell in enumerate(dots) if mask & (1 << i)]
        p["rotation"]["priorityList"] = entries
        choices.append((f"dots-{mask}", p))
    for label, omit in (("agony-only", {603}), ("doom-only", {11713}),
                        ("no-curse", {603, 11713})):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) not in omit]
        choices.append((label, p))
    for seconds in (6, 12, 18, 24):
        p = deepcopy(player)
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) in dots + [11713]:
                condition = entry["action"].get("condition", {})
                entry["action"]["condition"] = {"and": {"vals": [
                    condition, compare({"remainingTime": {}}, "OpGe", f"{seconds}s"),
                ]}}
        choices.append((f"dot-end{seconds}s", p))
    for mana in (5, 15, 25, 40):
        p = deepcopy(player)
        for entry in p["rotation"]["priorityList"]:
            if spell_id(entry) == 11689:
                entry["action"]["condition"] = compare({"currentManaPercent": {}}, "OpLt", f"{mana}%")
        choices.append((f"life-tap{mana}", p))
    for first in (25311, 25309, 603, 17924, 17923, 18871):
        if first == 18871 and not talent_rank(player, "warlock", "shadowburn"):
            continue
        p = deepcopy(player)
        entries = p["rotation"]["priorityList"]
        existing = next((a for a in entries if spell_id(a) == first), None)
        entries[:] = [a for a in entries if spell_id(a) != first]
        action = deepcopy(existing) if existing and first in (25311, 25309, 603) else cast(first)
        index = next(i for i, a in enumerate(entries) if spell_id(a) == 1311680) + 1
        entries.insert(index, action)
        choices.append((f"priority{first}", p))
    if talent_rank(player, "warlock", "conflagrate"):
        for seconds in (0, 1, 3, 6, 15):
            p = deepcopy(player)
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) == 18932:
                    entry["action"]["condition"] = compare(
                        {"dotRemainingTime": {"spellId": {"spellId": 25309}}}, "OpLe", f"{seconds}s")
            choices.append((f"conflagrate-{seconds}s", p))
    if talent_rank(player, "warlock", "decimation"):
        p = deepcopy(player)
        entries = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 17924]
        index = next(i for i, a in enumerate(entries) if spell_id(a) in filler_ids)
        entries.insert(index, cast(17924, {"auraIsActive": {"auraId": {"spellId": 440873}}}))
        p["rotation"]["priorityList"] = entries
        choices.append(("soul-fire-decimation", p))
    channels = [11700]
    if talent_rank(player, "warlock", "wrack"):
        channels.append(11704)
    for channel in channels:
        for interrupt in (False, True):
            p = deepcopy(player)
            checks = [{"and": {"vals": [
                {"spellCanCast": {"spellId": {"spellId": spell}}},
                {"not": {"val": {"dotIsActive": {"spellId": {"spellId": spell}}}}},
            ]}} for spell in dots if any(spell_id(a) == spell for a in p["rotation"]["priorityList"])]
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) in filler_ids:
                    entry["action"] = {"channelSpell": {
                        "spellId": {"spellId": channel},
                        "interruptIf": {"or": {"vals": checks}} if interrupt else {"const": {"val": "false"}},
                        "instantInterrupt": False,
                    }}
            choices.append((f"channel{channel}-interrupt{int(interrupt)}", p))
    p = deepcopy(player)
    for entry in p["rotation"]["priorityList"]:
        if spell_id(entry) in filler_ids:
            entry["action"] = cast(filler)["action"]
    choices.append(("direct-filler", p))
    return choices


def warrior(player):
    choices = []
    primary = 23894 if talent_rank(player, "warrior", "bloodthirst") else 21553
    for rage in (20, 30, 40, 50, 60, 70):
        for execute in (False, True):
            for early in (False, True):
                p = deepcopy(player)
                entries = p["rotation"]["priorityList"]
                heroic = next((a for a in entries if spell_id(a) == 25286), None)
                if heroic is None:
                    heroic = cast(25286)
                    heroic["action"]["castSpell"]["spellId"]["tag"] = 1
                    entries.append(heroic)
                conditions = [compare({"currentRage": {}}, "OpGe", rage)]
                if not execute:
                    conditions.append({"not": {"val": {"isExecutePhase": {"threshold": "E20"}}}})
                heroic["action"]["condition"] = {"and": {"vals": conditions}}
                if early:
                    entries.remove(heroic)
                    index = next(i for i, a in enumerate(entries) if "autocastOtherCooldowns" in a["action"]) + 1
                    entries.insert(index, heroic)
                choices.append((f"heroic{rage}-execute{int(execute)}-early{int(early)}", p))
    majors = tuple(spell for spell in (primary, 1680, 20662, 11605)
                   if any(spell_id(a) == spell for a in player["rotation"]["priorityList"]))
    for order in permutations(majors):
        p = deepcopy(player)
        entries = p["rotation"]["priorityList"]
        positions = [i for i, a in enumerate(entries) if spell_id(a) in majors]
        by_spell = {spell_id(entries[i]): entries[i] for i in positions}
        for position, spell in zip(positions, order):
            entries[position] = by_spell[spell]
        choices.append(("priority-" + "-".join(map(str, order)), p))
    for spell in (11574, 11605, 1680, 25286):
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) != spell]
        choices.append((f"omit{spell}", p))
    for rage in (15, 30, 50):
        p = deepcopy(player)
        entries = p["rotation"]["priorityList"]
        entry = next((a for a in entries if spell_id(a) == 11605), None)
        if entry is None:
            entry = cast(11605)
            entries.insert(len(entries) - 1, entry)
        entry["action"]["condition"] = compare({"currentRage": {}}, "OpGe", rage)
        choices.append((f"slam-rage{rage}", p))
    for seconds in (0, 9, 18):
        p = deepcopy(player)
        entries = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 11574]
        index = next(i for i, a in enumerate(entries) if spell_id(a) == primary)
        entries.insert(index, cast(11574, {"and": {"vals": [
            {"not": {"val": {"dotIsActive": {"spellId": {"spellId": 11574}}}}},
            compare({"remainingTime": {}}, "OpGe", f"{seconds}s"),
        ]}}))
        p["rotation"]["priorityList"] = entries
        choices.append((f"rend-end{seconds}", p))
    if talent_rank(player, "warrior", "spearingStrike"):
        for rage in (15, 50, 70):
            p = deepcopy(player)
            entries = [a for a in p["rotation"]["priorityList"] if spell_id(a) != 1310222]
            entries.append(cast(1310222, compare({"currentRage": {}}, "OpGe", rage)))
            p["rotation"]["priorityList"] = entries
            choices.append((f"spearing-rage{rage}", p))
    for seconds in (30, 60, 300):
        p = deepcopy(player)
        # The paired Arms stance conditions must follow Recklessness timing.
        def retime(value):
            if isinstance(value, dict):
                if value == {"const": {"val": "15s"}}:
                    value["const"]["val"] = f"{seconds}s"
                else:
                    for child in value.values():
                        retime(child)
            elif isinstance(value, list):
                for child in value:
                    retime(child)
        retime(p["rotation"])
        choices.append((f"recklessness-{seconds}s", p))
    if primary == 23894:
        for rage in (10, 20, 35, 50):
            p = deepcopy(player)
            for entry in p["rotation"]["priorityList"]:
                if spell_id(entry) == 2457:
                    entry["action"]["condition"]["and"]["vals"][1] = compare({"currentRage": {}}, "OpLe", rage)
            choices.append((f"overpower-dance{rage}", p))
        p = deepcopy(player)
        p["rotation"]["priorityList"] = [a for a in p["rotation"]["priorityList"] if spell_id(a) not in (2457, 2458, 11585)]
        choices.append(("no-overpower-dance", p))
    return choices


def variants(build, player):
    generators = {
        "balance": balance, "feral": feral, "elemental": shaman_caster,
        "stormcaller": shaman_caster, "enhancement": enhancement,
        "beast_mastery": ranged_hunter, "marksmanship": ranged_hunter,
        "survival": survival,
        "arcane": arcane, "fire": fire, "frost": frost,
        "combat": rogue, "mutilate": rogue, "subtlety": rogue,
        "retribution": retribution, "shadow": shadow, "smite": smite,
        "demonology": warlock, "affliction": warlock, "ds_ruin": warlock,
        "destruction": warlock, "fury": warrior, "arms": warrior,
    }
    if build not in generators:
        raise NotImplementedError(f"No rotation candidate family for {build}")
    choices = [("baseline", deepcopy(player))] + generators[build](player)
    seen = set()
    unique = []
    for label, p in choices:
        signature = json.dumps(p, sort_keys=True)
        if signature not in seen:
            seen.add(signature)
            unique.append((label, p))
    return unique
