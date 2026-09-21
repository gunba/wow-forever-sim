#!/usr/bin/env python3
"""Stage the benchmark matrix and replay files alongside the built simulator."""

import argparse
from html import escape
import json
from pathlib import Path
import re
import shutil

from build_display import BUILDS, CLASS_COLORS, RACES
from sensitivity import SCENARIOS, load_columns


SIM_PATHS = {
    "balance": "balance_druid", "feral": "feral_druid",
    "elemental": "elemental_shaman", "stormcaller": "elemental_shaman",
    "enhancement": "enhancement_shaman", "retribution": "retribution_paladin",
    "shadow": "shadow_priest", "smite": "smite_priest",
}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--results", type=Path, default=Path("artifacts/forever_dps_5min.json"))
    parser.add_argument("--profiles", type=Path, default=Path("artifacts/ui_profiles"))
    parser.add_argument("--output", type=Path, default=Path("dist/classic/review"))
    parser.add_argument("--sensitivity", type=Path, default=Path("artifacts/forever_sensitivity.json"))
    args = parser.parse_args()
    data = json.loads(args.results.read_text())
    rows = data["Results"]
    columns = load_columns(args.sensitivity, args.results)
    lookup = {(r["Key"], r["Race"]): r for r in rows}
    builds = sorted(BUILDS, key=lambda b: max(r["DPS"] for r in rows if r["Key"] == b[0]), reverse=True)
    args.output.mkdir(parents=True, exist_ok=True)
    shutil.copytree(args.profiles, args.output / "profiles", dirs_exist_ok=True)
    shutil.copytree("assets/img/spec_icons", args.output / "icons", dirs_exist_ok=True)
    for faction in ("alliance", "horde"):
        shutil.copyfile(Path("assets/img/wowhead/icons") / f"{faction}.png", args.output / "icons" / f"{faction}.png")
    shutil.copyfile(args.sensitivity, args.output / "sensitivity.json")
    shutil.copytree(args.sensitivity.parent / "sensitivity", args.output / "sensitivity", dirs_exist_ok=True)
    for extension in ("json", "csv", "svg", "png"):
        shutil.copyfile(args.results.with_suffix("." + extension), args.output / ("results." + extension))
    for name in ("build_reviews.md", "in_game_checks.md", "energy_audit.md", "auto_attack_audit.md"):
        shutil.copyfile(Path("docs") / name, args.output / name)
    body = []
    for key, class_name, label, icon in builds:
        simulator = SIM_PATHS.get(key, class_name.lower())
        cells = []
        for race in RACES:
            row = lookup.get((key, race))
            if row is None:
                cells.append('<td class="unavailable" aria-label="Unavailable">—</td>')
                continue
            race_file = re.sub(r"[^a-z0-9]+", "_", race.lower()).strip("_")
            filename = f"{key}__{race_file}.json"
            if not (args.profiles / filename).exists():
                raise ValueError(f"Missing replay profile: {filename}")
            title = f"Mean {row['DPS']:.2f} DPS; SE {row['StandardError']:.2f}; mana-limited {row['OOMSeconds']:.2f}s"
            cells.append(f'<td><a download href="profiles/{filename}" title="{escape(title)}">{row["DPS"]:.0f}</a></td>')
        for metric, *_ in SCENARIOS:
            value = columns[key][metric]
            title = (f'Equal-weight mean over {value["Races"]} races; conservative 95% Monte Carlo bound '
                     f'±{value["MonteCarlo95Bound"]:.2f} percentage points. Fixed talents and rotation.')
            cells.append(f'<td class="gain" title="{escape(title)}">{value["GainPercent"]:+.1f}%</td>')
        body.append(
            f'<tr><th scope="row" style="color:{CLASS_COLORS[class_name]}">'
            f'<a href="../{simulator}/"><img src="icons/{icon}.jpg" alt="">'
            f'{escape(class_name)}<br><span>{escape(label)}</span></a></th>{"".join(cells)}</tr>'
        )
    factions = {r["Race"]: r["Faction"] for r in rows}
    heading = "".join(f'<th scope="col" class="{factions[r].lower()}">{escape(r)}</th>' for r in RACES)
    heading += '<th scope="col" class="gain">Tier 1 gain</th><th scope="col" class="gain">Gear +10%</th><th scope="col" class="gain">Gear +20%</th>'
    groups = '<tr><th rowspan="2" scope="col">Class / build</th>'
    for faction in ("Alliance", "Horde"):
        count = sum(factions[r] == faction for r in RACES)
        groups += (f'<th colspan="{count}" scope="colgroup" class="faction {faction.lower()}">'
                   f'<img src="icons/{faction.lower()}.png" alt="">{faction}</th>')
    groups += '<th colspan="3" scope="colgroup" class="faction gain">Mean DPS gain</th></tr>'
    document = """<!doctype html>
<html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>Forever DPS benchmark</title>
<style>
body{background:#15171d;color:#e8e9ed;font:16px system-ui;margin:2rem;line-height:1.5}
a{color:#a9d8ff}main{max-width:1800px;margin:auto}.matrix{overflow-x:auto}
table{border-collapse:collapse;width:100%;font-variant-numeric:tabular-nums}
th,td{padding:.6rem;border-bottom:1px solid #343945;text-align:right}
tbody th,thead th[rowspan]{text-align:left;min-width:190px}thead{background:#242834}
tbody tr:hover{background:#242834}th a{color:inherit;text-decoration:none}
th img{width:34px;height:34px;float:left;margin:.15rem .7rem 0 0;border-radius:4px}
th span{font-weight:400}.unavailable{color:#68707e}td a{color:inherit}
.note{color:#b9c0cf;font-size:.9rem}.links{display:flex;gap:1rem;flex-wrap:wrap}
.alliance{background:#244b80}.horde{background:#813b43}
.faction{text-align:center}.faction img{float:none;width:auto;height:23px;vertical-align:middle;margin:0 .5rem}
.gain{background:#243b30;color:#b6efc8;white-space:nowrap}
</style><main>
<h1>Forever DPS benchmark</h1>
<p>Level 60 · five-minute single target · 5,000 iterations per race/build · full role-specific Tier 1 bonuses.</p>
<p>Rows are ordered by each build’s highest mean DPS. Click a DPS cell to download its replay profile;
open the simulator through the build name, then use <strong>Import → JSON</strong>.</p>
<p class="note">Crafted/dungeon gear; no world buffs. Hit is normalized through a paid benchmark budget,
not an obtainable reforging system. Imported bonus stats contain that fixed adjustment:
changing gear, talents or race requires recalculation for a fair comparison.
Energy scaling with general haste is a model assumption. These are tested builds, not proven global optima.</p>
<nav class="links"><a href="results.png">Chart PNG</a><a href="results.svg">Chart SVG</a>
<a href="results.csv" download>CSV</a><a href="results.json" download>Raw requests/results</a>
<a href="sensitivity.json" download>Gain calculations</a>
<a href="sensitivity/tier1_off.json" download>Tier 1 off</a>
<a href="sensitivity/gear_110.json" download>Gear +10% run</a>
<a href="sensitivity/gear_120.json" download>Gear +20% run</a>
<a href="build_reviews.md">Build reviews</a><a href="in_game_checks.md">In-game checks</a>
<a href="energy_audit.md">Energy model</a><a href="auto_attack_audit.md">Auto-attack model</a></nav>
<p class="note">Hover a result for its standard error and mana-limited time. A dash means that race/class combination is unavailable.</p>
<p class="note">Gain columns average each available race’s percentage DPS change with equal weights.
Tier 1 gain compares bonuses on versus off, using the same build.
Gear columns increase item/suffix stats and weapon damage together; enchants, weapon speed/skill,
procs, consumables and buffs stay fixed. Tier 1 stays on and paid hit is recalculated.
These are hypothetical sensitivities with fixed talents/rotations, not stat weights or predictions of future items.
Cat weapon-DPS scaling remains an open mechanic. Hover gain cells for Monte Carlo uncertainty.</p>
<div class="matrix"><table><thead>""" + groups + "<tr>" + heading + \
        "</tr></thead><tbody>" + "".join(body) + """</tbody></table></div>
<footer class="note"><p>Unofficial beta simulator. Built on
<a href="https://github.com/wowsims/classic">WoWSims Classic</a> and
<a href="https://github.com/ElliotWood/Forever">ElliotWood/Forever</a>.
<a href="https://github.com/gunba/wow-forever-sim">Source and issues</a>.
Game icons via Wowhead.</p></footer></main></html>
"""
    (args.output / "index.html").write_text(document)
    print(f"Review site staged at {args.output}")


if __name__ == "__main__":
    main()
