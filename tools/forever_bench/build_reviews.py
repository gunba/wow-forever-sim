#!/usr/bin/env python3
"""Generate class-review summaries from the published benchmark requests."""

import argparse
import json
from pathlib import Path

from build_display import BUILDS, BUILD_CAVEATS


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--results", type=Path, default=Path("artifacts/forever_dps_5min.json"))
    parser.add_argument("--output", type=Path, default=Path("docs/build_reviews.md"))
    args = parser.parse_args()
    names = {}
    for path in Path("ui/core/spells").glob("*.json"):
        for key, value in json.loads(path.read_text()).items():
            if key.isdigit() and isinstance(value, dict):
                names[int(key)] = value.get("ability", f"Spell {key}")
    database = json.loads(Path("assets/database/db.json").read_text())
    item_records = {i["id"]: i for i in database["items"]}
    synthetic = {
        item["id"]: item
        for path in ("forever_synthetic_gear.json", "forever_synthetic_gear_v2.json")
        for item in json.loads((Path("assets/db_inputs") / path).read_text())["items"]
    }
    items = {i: item["name"] for i, item in item_records.items()}
    enchants = {e["effectId"]: e["name"] for e in database["enchants"]}
    display = {key: (cls, spec) for key, cls, spec, _ in BUILDS}

    def action_name(action):
        if action.get("spellId"):
            spell = action["spellId"]
            label = names.get(spell, f"Spell {spell}")
            if action.get("rank"):
                label += f" (rank {action['rank']})"
            return f"[{label}](https://www.wowhead.com/forever/spell={spell})"
        if action.get("itemId"):
            item = action["itemId"]
            if item in synthetic:
                reference = synthetic[item]["modelReferenceItemID"]
                return f"[{items.get(item, f'Modeled item {item}')} (reference only)](https://www.wowhead.com/forever/item={reference})"
            return f"[{items.get(item, f'Item {item}')}](https://www.wowhead.com/forever/item={item})"
        label = {"OtherActionAttack": "Auto-attack", "OtherActionShoot": "Shoot"}.get(action.get("otherId"), action.get("otherId", "Unknown action"))
        if action.get("tag"):
            label += f" (tag {action['tag']})"
        return label

    def condition(value):
        if "const" in value:
            return value["const"]["val"]
        for key, label in {
            "currentMana": "Mana", "currentManaPercent": "Mana fraction",
            "currentEnergy": "Energy", "currentRage": "Rage",
            "currentComboPoints": "Combo points", "remainingTime": "Time remaining",
            "currentTime": "Time elapsed", "numberTargets": "Target count",
        }.items():
            if key in value:
                return label
        for key, word in (("and", " AND "), ("or", " OR ")):
            if key in value:
                return "(" + word.join(condition(v) for v in value[key]["vals"]) + ")"
        if "not" in value:
            return "NOT " + condition(value["not"]["val"])
        for kind in ("cmp", "math"):
            if kind in value:
                v = value[kind]
                op = {"OpLe": "≤", "OpLt": "<", "OpGe": "≥", "OpGt": ">", "OpEq": "=",
                      "OpNe": "≠", "OpMul": "×", "OpDiv": "/", "OpAdd": "+", "OpSub": "−"}[v["op"]]
                return f"{condition(v['lhs'])} {op} {condition(v['rhs'])}"
        for key, suffix, identifier in (
            ("dotIsActive", "DoT active", "spellId"),
            ("dotRemainingTime", "DoT time remaining", "spellId"),
            ("auraIsActive", "active", "auraId"),
            ("auraNumStacks", "stacks", "auraId"),
        ):
            if key in value:
                return f"{action_name(value[key][identifier])} {suffix}"
        if "totemRemainingTime" in value:
            return value["totemRemainingTime"]["totemType"] + " totem time remaining"
        return "`" + json.dumps(value, separators=(",", ":")) + "`"

    results = json.loads(args.results.read_text())
    modeled = results.get("GearScenario") in {"modeled-65-v1", "modeled-65-v2"}
    groups = {}
    for row in results["Results"]:
        groups.setdefault(row["Key"], []).append(row)
    ordered = sorted(groups, key=lambda key: max(r["DPS"] for r in groups[key]), reverse=True)
    matrix_path = ("../artifacts/modelled_gear/forever_dps_5min.png"
                   if modeled else "../artifacts/forever_dps_5min.png")
    raw_path = ("../artifacts/modelled_gear/forever_dps_5min.json"
                if modeled else "../artifacts/forever_dps_5min.json")
    search_path = ("../artifacts/modelled_gear_search/current/summary.json"
                   if modeled else "../artifacts/gear_search/summary.json")
    lines = [
        "# Build reviews", "",
        ("These are hypothetical modeled-only gear results, not obtainable item rankings. "
         "Sources linked for modeled items are allocation references only and have different "
         "stats. Gear is compared under the [model assumptions](modelled_gear.md)."
         if modeled else "These summaries describe the published loadouts. They are simulation "
         "results, not independent confirmation of server mechanics or proof of a global optimum."), "",
        f"The tables and [matrix]({matrix_path}) use the same "
        f"{len(results['Results'])} common-seed replays. "
        + ("The modeled search and original real-item benchmark are separate; neither proves "
           "available launch gear." if modeled else "Equipment selections came from an earlier "
           "mechanics revision; these results use the corrected engine. Historical search gains "
           "are not directly comparable to this release."), "",
        "The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific "
        "Tier 1 bonuses, and " +
        ("hit from selected gear, enchants, talents and racials only. " if
         results.get("GearScenario") == "modeled-65-v2" else
         "paid shared-hit normalization. ") +
        "[Scenario and exchange model](../tools/forever_bench/README.md) · "
        "[In-game checks](in_game_checks.md)", "",
    ]
    slots = ["Head", "Neck", "Shoulders", "Back", "Chest", "Wrists", "Hands",
             "Waist", "Legs", "Feet", "Ring 1", "Ring 2", "Trinket 1", "Trinket 2",
             "Main hand", "Off hand", "Ranged/relic"]
    for key in ordered:
        cls, spec = display[key]
        rows = groups[key]
        representative = max(rows, key=lambda r: r["DPS"])
        p = representative["BaselinePlayer"]
        points = [sum(map(int, part)) for part in p["talentsString"].split("-")]
        points += [0] * (3-len(points))
        lines += [f"## {cls} — {spec}", "",
                  f"**Talents:** {'/'.join(map(str, points))} · `{p['talentsString']}`", "",
                  f"[Requests and results]({raw_path}) · "
                  f"[Equipment search]({search_path})", ""]
        if key in BUILD_CAVEATS:
            lines += ["**Model limitations:** " + " ".join(BUILD_CAVEATS[key]), ""]
        if len({json.dumps(row["BaselinePlayer"]["rotation"], sort_keys=True) for row in rows}) > 1:
            lines += [
                f"**Rotation variants:** The priorities below are for {representative['Race']}. "
                "Other races may use a different saved APL; their exact rotations are in "
                "the Ranked builds selector and raw requests.", "",
            ]
        lines += ["", "### Results", "",
                  "| Race | DPS | Standard error | Mana-limited seconds |",
                  "|---|---:|---:|---:|"]
        for r in rows:
            lines.append(f"| {r['Race']} | {r['DPS']:.2f} | "
                         f"{r['StandardError']:.2f} | {r['OOMSeconds']:.2f} |")
        lines += ["", "Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.", ""]
        lines += [f"### Equipment — {representative['Race']}", "",
                  "| Slot | Item | Item level | Enchant |", "|---|---|---:|---|"]
        for slot, entry in zip(slots, p["equipment"]["items"]):
            item = item_records.get(entry.get("id"))
            if item:
                modeled_item = synthetic.get(item['id'])
                reference = modeled_item["modelReferenceItemID"] if modeled_item else item['id']
                suffix = ' (modeled; reference only)' if modeled_item else ''
                lines.append(f"| {slot} | [{item['name']}{suffix}](https://www.wowhead.com/forever/item={reference}) | "
                             f"{item['ilvl']} | {enchants.get(entry.get('enchant'), '—')} |")
        lines += ["", "Other races can use different equipment. Their complete setups are "
                  "available in the simulator's Ranked builds selector.", ""]
        unsupported = sorted({str(effect) for row in rows for effect in row.get("UnmodeledSetBonuses") or []})
        if unsupported:
            lines += ["**Unmodeled equipped-set effects:** " + "; ".join(unsupported), ""]
        if p["rotation"].get("prepullActions"):
            lines += ["### Before the pull", ""]
            for entry in p["rotation"]["prepullActions"]:
                a = entry["action"]
                text = action_name(a["castSpell"]["spellId"]) if "castSpell" in a else "`" + json.dumps(a) + "`"
                lines.append(f"- {condition(entry['doAtValue'])}: {text}.")
            lines.append("")
        lines += ["### Rotation priorities", ""]
        for index, entry in enumerate(p["rotation"].get("priorityList", []), 1):
            a = entry["action"]
            if "castSpell" in a:
                text = "Cast " + action_name(a["castSpell"]["spellId"])
            elif "autocastOtherCooldowns" in a:
                text = "Use ready automatic cooldowns"
            else:
                text = "`" + json.dumps(a, separators=(",", ":")) + "`"
            if "condition" in a:
                text += " when " + condition(a["condition"])
            lines.append(f"{index}. {text}.")
        lines += ["", f"### Damage breakdown — {representative['Race']}", "",
                  "| Action | DPS |", "|---|---:|"]
        damage = []

        def collect(metrics, prefix=""):
            for action in metrics.get("actions", []):
                # This benchmark has one enemy, in target slot zero. Later
                # slots are players/pets: SW:D backlash, Sappers and Demonic
                # Rune must not appear as damage dealt to the boss.
                targets = action.get("targets", [])
                total = targets[0].get("damage", 0) if targets else 0
                if total > 0:
                    damage.append((total / representative["Iterations"] / results["Duration"],
                                   prefix + action_name(action["id"])))
            for pet in metrics.get("pets", []):
                collect(pet, pet.get("name", "Pet") + ": ")

        collect(representative["Metrics"])
        damage.sort(reverse=True)
        lines += [f"| {name} | {dps:.2f} |" for dps, name in damage[:8]]
        lines += ["", "### Resource flow", "",
                  "| Resource | Action | Net amount per fight |",
                  "|---|---|---:|"]
        resources = sorted(representative["Metrics"].get("resources", []),
                           key=lambda resource: abs(resource.get("actualGain", 0)), reverse=True)
        for resource in resources[:12]:
            lines.append(f"| {resource['type'].removeprefix('ResourceType')} | "
                         f"{action_name(resource['id'])} | "
                         f"{resource.get('actualGain', 0)/representative['Iterations']:+.1f} |")
        lines.append("")
    args.output.write_text("\n".join(lines).rstrip() + "\n")
    print(args.output)


if __name__ == "__main__":
    main()
