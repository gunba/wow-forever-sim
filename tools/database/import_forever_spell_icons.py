#!/usr/bin/env python3
"""Bundle Forever display metadata for benchmark spells and internal aliases."""

import concurrent.futures
import html
import json
from pathlib import Path
import re
import urllib.request


def spell_ids(value):
    if isinstance(value, dict):
        for key, child in value.items():
            if key == "spellId" and isinstance(child, int) and child > 0:
                yield child
            else:
                yield from spell_ids(child)
    elif isinstance(value, list):
        for child in value:
            yield from spell_ids(child)


def tooltip(spell_id):
    url = f"https://nether.wowhead.com/forever/tooltip/spell/{spell_id}?lvl=60"
    try:
        with urllib.request.urlopen(url, timeout=30) as response:
            data = json.load(response)
        if not data.get("name") or not data.get("icon"):
            raise ValueError("no display metadata")
        rank = re.search(r"\bRank\s+(\d+)", html.unescape(re.sub("<[^>]+>", " ", data.get("tooltip", ""))))
        return spell_id, {
            "id": spell_id, "name": data["name"], "icon": data["icon"].lower(),
            "rank": int(rank[1]) if rank else 0, "hasBuff": bool(data.get("buff")),
        }
    except Exception as error:
        print(f"Unresolved spell icon {spell_id}: {error}")
        return spell_id, None


def main():
    output = Path("assets/db_inputs/forever_spell_icons.json")
    cached = {r["id"]: r for r in json.loads(output.read_text())["spellIcons"]} if output.exists() else {}
    database = json.loads(Path("assets/database/db.json").read_text())
    known = {r["id"]: r for r in database["spellIcons"] if r.get("icon")}
    aliases = {}
    names = {}
    for path in Path("ui/core/spells").glob("*.json"):
        for key, entry in json.loads(path.read_text()).items():
            names[int(key)] = entry["ability"]
            if entry.get("foreverId"):
                aliases[int(key)] = entry["foreverId"]
    talent_icons = {}
    for path in Path("ui/core/talents/trees").glob("*.json"):
        for tree in json.loads(path.read_text()):
            for talent in tree["talents"]:
                if talent.get("name") and talent.get("icon"):
                    talent_icons.setdefault(talent["name"], set()).add(talent["icon"])
    needed = set(spell_ids(json.loads(Path("artifacts/forever_dps_5min.json").read_text())))
    needed = {aliases.get(id, id) for id in needed} | set(aliases.values())
    missing = sorted(needed - known.keys() - cached.keys())
    with concurrent.futures.ThreadPoolExecutor(max_workers=6) as pool:
        for id, row in pool.map(tooltip, missing):
            if row:
                cached[id] = row
            elif len(talent_icons.get(names.get(id), set())) == 1:
                # Display-only alias: do not point its tooltip at a different rank.
                cached[id] = {
                    "id": id, "name": names[id],
                    "icon": next(iter(talent_icons[names[id]])).lower(),
                }
                print(f"Using the talent icon for internal action {id} ({names[id]})")
    for alias, actual in aliases.items():
        row = cached.get(actual) or known.get(actual)
        if row:
            cached[actual] = row
            cached[alias] = {**row, "id": alias}
    output.write_text(json.dumps({"spellIcons": [cached[id] for id in sorted(cached)]}, indent=2) + "\n")
    print(f"Bundled {len(cached)} spell metadata records")


if __name__ == "__main__":
    main()
