#!/usr/bin/env python3
"""Generate class-review summaries from retained build evidence."""

import argparse
import json
from pathlib import Path

from build_display import BUILDS


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--input", type=Path, default=Path("artifacts/optimization"))
    parser.add_argument("--output", type=Path, default=Path("docs/build_reviews.md"))
    args = parser.parse_args()
    names = {}
    for path in Path("ui/core/spells").glob("*.json"):
        for key, value in json.loads(path.read_text()).items():
            if key.isdigit() and isinstance(value, dict):
                names[int(key)] = value.get("ability", f"Spell {key}")
    items = {i["id"]: i["name"] for i in json.loads(Path("assets/database/db.json").read_text())["items"]}
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

    bundles = [json.loads(p.read_text()) for p in args.input.glob("*.json")]
    bundles.sort(key=lambda b: max(r["DPS"] for r in b["retained"]["Results"]), reverse=True)
    lines = [
        "# Build reviews", "",
        "These summaries cover completed build comparisons. They are simulation results, "
        "not independent confirmation of server mechanics or proof of a global optimum.", "",
        "The tables below use each build's independent validation run. "
        "The [final matrix](../artifacts/forever_dps_5min.png) uses a separate common-seed "
        "run of all 147 combinations, so small Monte Carlo differences are expected. "
        "The final matrix also includes the later Eureka nested-charge correction; "
        "the optimization validation tables remain historical records.", "",
        "The benchmark uses level 60, 300 seconds, one level-63 target, complete role-specific "
        "Tier 1 bonuses, and paid shared-hit normalization. "
        "[Scenario and exchange model](../tools/forever_bench/README.md) · "
        "[In-game checks](in_game_checks.md)", "",
    ]
    for bundle in bundles:
        key = bundle["build"]
        cls, spec = display[key]
        rows = bundle["retained"]["Results"]
        representative = max(rows, key=lambda r: r["DPS"])
        p = representative["BaselinePlayer"]
        points = [sum(map(int, part)) for part in p["talentsString"].split("-")]
        points += [0] * (3-len(points))
        lines += [f"## {cls} — {spec}", "",
                  f"**Talents:** {'/'.join(map(str, points))} · `{p['talentsString']}`", "",
                  f"[Requests, results and search evidence](../artifacts/optimization/{key}.json)", ""]
        if "Mechanics" not in bundle["retained"]:
            lines += [
                "> Historical comparison: this run predates continuous autos and "
                "haste-scaled Energy. Revalidation is pending where those rules affect the build.", "",
            ]
        lines += ["- " + note for note in bundle["notes"]]
        lines += ["", "### Results", "",
                  "| Race | Baseline DPS | Retained DPS | Change | Mana-limited seconds |",
                  "|---|---:|---:|---:|---:|"]
        by_race = {r["Race"]: r for r in rows}
        for result in bundle["comparison"]:
            r = by_race[result["race"]]
            lines.append(f"| {r['Race']} | {result['baselineDPS']:.2f} | {r['DPS']:.2f} | "
                         f"{result['gainPercent']:+.2f}% | {r['OOMSeconds']:.2f} |")
        lines += ["", "Mana-limited time counts failed mana-cost checks; it is not necessarily zero-damage time.", ""]
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
                total = sum(t.get("damage", 0) for t in action.get("targets", []))
                if total > 0:
                    damage.append((total / representative["Iterations"] / bundle["retained"]["Duration"],
                                   prefix + action_name(action["id"])))
            for pet in metrics.get("pets", []):
                collect(pet, pet.get("name", "Pet") + ": ")

        collect(representative["Metrics"])
        damage.sort(reverse=True)
        lines += [f"| {name} | {dps:.2f} |" for dps, name in damage[:8]]
        lines.append("")
    args.output.write_text("\n".join(lines).rstrip() + "\n")
    print(args.output)


if __name__ == "__main__":
    main()
