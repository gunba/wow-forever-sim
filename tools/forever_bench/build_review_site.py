#!/usr/bin/env python3
"""Stage the benchmark matrix and replay files alongside the built simulator."""

import argparse
from html import escape
import json
from pathlib import Path
import re
import shutil

from build_display import BUILDS, BUILD_CAVEATS, CLASS_COLORS, HYBRID_PARENTS, RACES
from sensitivity import DISPLAY_METRICS, format_metric, load_columns
from uncertainties import CSS as QUESTIONS_CSS, SCRIPT as QUESTIONS_SCRIPT
from uncertainties import load_register, markdown as questions_markdown, render_html as questions_html


SIM_PATHS = {
    "balance": "balance_druid", "feral": "feral_druid",
    "elemental": "elemental_shaman", "stormcaller": "elemental_shaman",
    "enhancement": "enhancement_shaman", "retribution": "retribution_paladin",
    "retribution_physical": "retribution_paladin",
    "fury_sunder": "warrior",
    "shadow": "shadow_priest", "smite": "smite_priest",
    "tank_warrior": "tank_warrior",
    "protection_paladin": "protection_paladin",
    "feral_tank_druid": "feral_tank_druid",
}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--results", type=Path, default=Path("artifacts/modelled_gear/forever_dps_5min.json"))
    parser.add_argument("--profiles", type=Path, default=Path("artifacts/modelled_gear/ui_profiles"))
    parser.add_argument("--output", type=Path, default=Path("dist/classic/review"))
    parser.add_argument("--sensitivity", type=Path, default=Path("artifacts/modelled_gear/forever_sensitivity.json"))
    parser.add_argument("--gear-search", type=Path, default=Path("artifacts/modelled_gear_search/current"))
    args = parser.parse_args()
    data = json.loads(args.results.read_text())
    model_v2 = data.get("GearScenario") == "modeled-65-v2"
    modeled = data.get("GearScenario") in {"modeled-65-v1", "modeled-65-v2"}
    rows = data["Results"]
    unenchanted = all(
        not any(item.get("enchant") for item in row["BaselinePlayer"]["equipment"]["items"])
        for row in rows
    )
    columns = load_columns(args.sensitivity, args.results)
    lookup = {(r["Key"], r["Race"]): r for r in rows}
    items = {item["id"]: item for item in json.loads(Path("assets/database/db.json").read_text())["items"]}
    invalid = set()
    for row in rows:
        if row["Key"] not in {"enhancement", "elemental", "stormcaller"}:
            continue
        off_hand = row["BaselinePlayer"]["equipment"]["items"][15].get("id", 0)
        if items.get(off_hand, {}).get("weaponType") in {1, 2, 3, 4, 6, 8, 9}:
            invalid.add((row["Key"], row["Race"]))
    builds = sorted((b for b in BUILDS if any(r["Key"] == b[0] for r in rows)), key=lambda b: max(
        (r["DPS"] for r in rows if r["Key"] == b[0] and (r["Key"], r["Race"]) not in invalid),
        default=-1,
    ), reverse=True)
    args.output.mkdir(parents=True, exist_ok=True)
    register = load_register()
    (args.output / "uncertainties.json").write_text(json.dumps(register, indent=2) + "\n")
    (args.output / "uncertainties.md").write_text(questions_markdown(register) + "\n")
    shutil.copytree(args.profiles, args.output / "profiles", dirs_exist_ok=True)
    shutil.copytree("assets/img/spec_icons", args.output / "icons", dirs_exist_ok=True)
    for faction in ("alliance", "horde"):
        shutil.copyfile(Path("assets/img/wowhead/icons") / f"{faction}.png", args.output / "icons" / f"{faction}.png")
    shutil.copyfile(args.sensitivity, args.output / "sensitivity.json")
    shutil.copytree(args.sensitivity.parent / "sensitivity", args.output / "sensitivity", dirs_exist_ok=True)
    shutil.copytree("artifacts/mana_regen", args.output / "mana_regen", dirs_exist_ok=True)
    shutil.copytree("artifacts/research_builds", args.output / "research_builds", dirs_exist_ok=True)
    shutil.copytree(args.gear_search, args.output / "gear_search", dirs_exist_ok=True)
    if modeled:
        shutil.copyfile("assets/db_inputs/forever_synthetic_gear_v2.json"
                        if model_v2 else "assets/db_inputs/forever_synthetic_gear.json",
                        args.output / "synthetic_gear.json")
        shutil.copyfile("docs/modelled_gear.md", args.output / "modelled_gear.md")
        shutil.copyfile("docs/modelled_build_reviews.md", args.output / "modelled_build_reviews.md")
        shutil.copyfile("artifacts/modelled_gear/proposed_gear.xlsx", args.output / "proposed_gear.xlsx")
        if model_v2:
            for filename in ("gear_equity.json", "selected_gear_equity.csv",
                             "spec_gear_equity.csv"):
                shutil.copyfile(Path("artifacts/modelled_gear") / filename,
                                args.output / filename)
    for directory in ("profile_corrections", "windfury"):
        shutil.copytree(Path("artifacts") / directory, args.output / directory, dirs_exist_ok=True)
    shutil.copyfile("artifacts/spell_coverage.json", args.output / "spell_coverage.json")
    for extension in ("json", "csv", "svg", "png"):
        shutil.copyfile(args.results.with_suffix("." + extension), args.output / ("results." + extension))
    for name in ("build_reviews.md", "build_updates.md", "gear_updates.md", "in_game_checks.md", "check_dispositions.md", "spell_coverage.md", "windfury.md", "energy_audit.md", "auto_attack_audit.md", "crit_model.md", "forever_gear_data.md", "mechanics_review.md", "history_review.md", "upstream-forever-review-2026-09-23.md", "upstream-forever-followup-2026-09-23.md", "mana_regeneration.md", "mythicsim_review.md"):
        shutil.copyfile(Path("docs") / name, args.output / name)
    for name in ("weekly_review.md", "weekly_review_commits.csv", "flurry_review.md", "upstream_elliot_review.md", "tank_benchmark.md", "tank_selection.md", "forever-70009.md"):
        shutil.copyfile(Path("docs") / name, args.output / name)
    body = []
    for key, class_name, label, icon in builds:
        simulator = SIM_PATHS.get(key, class_name.lower())
        cells = []
        for race in RACES:
            row = lookup.get((key, race))
            if (key, race) in invalid:
                cells.append('<td title="Illegal dual-wield Shaman layout">Invalid</td>')
                continue
            if row is None:
                cells.append('<td class="unavailable" aria-label="Unavailable">—</td>')
                continue
            race_file = re.sub(r"[^a-z0-9]+", "_", race.lower()).strip("_")
            filename = f"{key}__{race_file}.json"
            if not (args.profiles / filename).exists():
                raise ValueError(f"Missing replay profile: {filename}")
            title = f"Mean {row['DPS']:.2f} DPS; SE {row['StandardError']:.2f}; mana-limited {row['OOMSeconds']:.2f}s"
            title += " " + " ".join(BUILD_CAVEATS.get(key, []))
            if row.get("Tank"):
                tank = row["Tank"]
                title += (f' Tank scenario: {tank["TPS"]:.1f} TPS; {tank["DTPS"]:.1f} DTPS; '
                          f'{100*tank["ChanceOfDeath"]:.2f}% modeled death probability.')
            if row.get("UnmodeledSetBonuses"):
                title += f'; {len(row["UnmodeledSetBonuses"])} equipped-set effects omitted'
            cells.append(f'<td><a href="../{simulator}/?profile={key}__{race_file}" title="{escape(title)}">{row["DPS"]:.0f}</a></td>')
        for metric in DISPLAY_METRICS:
            if any(k == key for k, _ in invalid):
                cells.append('<td class="unavailable">—</td>')
                continue
            value = columns[key][metric]
            if metric == "amplification":
                title = (f'Ratio of equal-race-weight gains over {value["Races"]} races: '
                         f'+50% gear gain {value["Gear50GainPercent"]:+.2f}% / '
                         f'(5 × +10% gear gain {value["Gear10GainPercent"]:+.2f}%). '
                         '1× linear; >1× accelerating; <1× flattening. ')
                title += (value["UnavailableReason"] if value["Amplification"] is None else
                          f'Conservative 95% Monte Carlo bound ±{value["MonteCarlo95Bound"]:.2f}×. '
                          'Caps and resource thresholds can affect this measure.')
            else:
                title = (f'Equal-weight mean over {value["Races"]} races; conservative 95% Monte Carlo bound '
                         f'±{value["MonteCarlo95Bound"]:.2f} percentage points. Fixed talents and rotation.')
            cells.append(f'<td class="gain" title="{escape(title)}">{format_metric(metric, value)}</td>')
        body.append(
            f'<tr><th scope="row" style="color:{CLASS_COLORS[class_name]}">'
            f'<a href="../{simulator}/?build={key}"><img src="icons/{icon}.jpg" alt="">'
            f'{escape(class_name)}<br><span>{escape(label)}{"*" if key in HYBRID_PARENTS else ""}</span></a></th>{"".join(cells)}</tr>'
        )
    factions = {r["Race"]: r["Faction"] for r in rows}
    heading = "".join(f'<th scope="col" class="{factions[r].lower()}">{escape(r)}</th>' for r in RACES)
    heading += '<th scope="col" class="gain">Tier 1 gain</th><th scope="col" class="gain">Gear +10%</th><th scope="col" class="gain">Scaling amp.</th>'
    groups = '<tr><th rowspan="2" scope="col">Class / build</th>'
    for faction in ("Alliance", "Horde"):
        count = sum(factions[r] == faction for r in RACES)
        groups += (f'<th colspan="{count}" scope="colgroup" class="faction {faction.lower()}">'
                   f'<img src="icons/{faction.lower()}.png" alt="">{faction}</th>')
    groups += '<th colspan="3" scope="colgroup" class="faction gain">Gains &amp; scaling</th></tr>'
    equipment_note = (
        "All equipment is hypothetical level-65 modeled gear; real enchants remain. "
        "Hit comes from items, enchants, talents and racials, with no paid conversion."
        if model_v2 else
        "Equipment and any historical hit-budget adjustments are recorded in each exact profile."
    )
    equipment_links = (
        '<a href="modelled_gear.md">Projection method and limits</a> · '
        '<a href="proposed_gear.xlsx" download>Gear spreadsheet</a> · '
        '<a href="synthetic_gear.json">Modeled item catalog</a> · '
        '<a href="gear_search/summary.json">Gear-selection evidence</a> · '
        '<a href="spec_gear_equity.csv">Stat/budget audit</a> · '
        '<a href="selected_gear_equity.csv">Per-race gear/hit audit</a>'
        if model_v2 else '<a href="forever_gear_data.md">Equipment sources and gaps</a>'
    )
    build_review = "modelled_build_reviews.md" if modeled else "build_reviews.md"
    tank_rows = [r for r in rows if r.get("Tank")]
    tank_html = ""
    if tank_rows:
        table = []
        validation = json.loads(Path("artifacts/tanks/current/validation.json").read_text())
        selected_tanks = {(r["Key"], r["Race"]): r for r in validation["selections"]}
        for row in sorted(tank_rows, key=lambda r: (r["Build"], -r["DPS"])):
            tank = row["Tank"]
            multi = selected_tanks[(row["Key"], row["Race"])]["Multi"]
            table.append(
                f'<tr><th>{escape(row["Build"])} · {escape(row["Race"])}</th>'
                f'<td>{row["DPS"]:.1f}</td><td>{tank["TPS"]:.1f}</td>'
                f'<td>{tank["DTPS"]:.1f}</td><td>{tank["TMI"]:.1f}</td>'
                f'<td>{100*tank["ChanceOfDeath"]:.2f}%</td>'
                f'<td title="Lowest mean enemy TPS in the separate 90-second, three-attacker validation">'
                f'{multi["LeastTargetTPS"]:.1f}</td>'
                f'<td>{", ".join(f"{t:.0f}" for t in multi["TargetTPS"])}</td></tr>')
        tank_html = (
            '<section id="tanks"><h2>Tank performance</h2>'
            '<p>Tank rows take frontal attacks: 3,000 base damage every two seconds, parry haste, '
            '1,500 HPS in three-second heals. Other DPS rows do not take these attacks. '
            'External support differs between tank classes; this is not an equal-support survival ranking.</p>'
            '<p class="note">Iterations continue after a modeled death. DPS is not discounted for death downtime; '
            'survival is evaluated separately.</p>'
            '<p><a href="tank_selection.md">Selection safeguards and three-attacker scenario</a> · '
            '<a href="tanks/validation.json">Matched validation and per-target threat</a> · '
            '<a href="tanks/metrics.csv">Tank metrics CSV</a> · '
            '<a href="tanks/archives.json">Research archives</a></p>'
            '<div class="matrix"><table><thead><tr><th>Build / race</th><th>DPS</th><th>TPS</th>'
            '<th>DTPS</th><th>TMI</th><th>Modeled death</th>'
            '<th>3 attackers: minimum TPS</th><th>TPS by attacker</th></tr></thead><tbody>'
            + "".join(table) + '</tbody></table></div></section>'
        )
        shutil.copytree("artifacts/tanks/current", args.output / "tanks", dirs_exist_ok=True)
    document = """<!doctype html>
<html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>Forever simulations</title>
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
""" + QUESTIONS_CSS + """</style><main>
<h1>Forever simulations</h1>
<p>Level 60 · five-minute single target · 5,000 iterations per race/build · full role-specific Tier 1 bonuses.</p>
<!-- equipment-status -->
<!-- invalid-results -->
<p class="note">""" + equipment_note + """ Ordinary MP5; no world or raid campfire buffs.
These are tested profiles, not proven global optima or measured class balance.</p>
<nav class="links" aria-label="Benchmark navigation"><a href="#matrix">DPS matrix</a>
<a href="#questions">Questions &amp; coverage</a><a href="results.png">Chart</a>
<a href="""" + build_review + """">Build details</a></nav>
""" + questions_html(register) + """
<details class="resource-group"><summary>Method, gear and gain columns</summary>
<p>""" + equipment_note + """</p><p>""" + equipment_links + """</p>
<p class="note">Gear is selected by legal slot-coordinate comparisons, not an exhaustive combination search.
Projected allocation budgets and special-stat prices are assumptions, not a recovered Blizzard formula.
All selected items, enchants and scenario settings are retained in the replay profiles.</p>
<p class="note">Gain columns average each available race’s percentage DPS change equally.
Tier 1 compares bonuses on versus off with the same build.
Gear scenarios increase item/suffix stats and weapon damage together; enchants, weapon speed/skill,
procs, consumables and buffs stay fixed. Tier 1 stays on.
Scaling amplification is the mean +50% gain divided by five times the mean +10% gain:
1× is linear, above 1× accelerates, below 1× flattens. A negative value means the +50% scenario loses DPS.
An amplification dash means the +10% gain is too small or noisy for a useful ratio.
This is finite-range curvature, not proof of exponential growth, isolated synergy or future item strength.</p></details>
<details class="resource-group"><summary>Replay files and downloads</summary>
<nav class="links"><a href="results.png">PNG</a><a href="results.svg">SVG</a>
<a href="results.csv" download>CSV</a><a href="results.json" download>Raw requests/results</a>
<a href="profiles/index.json">Profile index</a><a href="sensitivity.json">Gain accounting</a>
<a href="sensitivity/tier1_off.json">Tier off</a><a href="sensitivity/gear_110.json">Gear +10%</a>
<a href="sensitivity/gear_150.json">Gear +50%</a>
<a href="research_builds/validation.json.gz">Historical build validation</a></nav></details>
<details class="resource-group"><summary>Evidence and earlier reviews</summary>
<p class="note">These reports describe their pinned revisions. Current dispositions are in the register above.</p>
<nav class="links"><a href="upstream_elliot_review.md">Latest class-effects review</a>
<a href="weekly_review.md">September 26 review</a><a href="flurry_review.md">Flurry evidence</a>
<a href="tank_benchmark.md">Tank scenario evidence</a>
<a href="spell_coverage.md">Spell inventory</a><a href="forever-70009.md">70009 patch</a>
<a href="mechanics_review.md">Mechanics review</a><a href="history_review.md">Change-history review</a>
<a href="mythicsim_review.md">Independent engine comparison</a>
<a href="build_updates.md">Earlier build comparisons</a>
<a href="in_game_checks.md">Legacy test IDs</a><a href="check_dispositions.md">Legacy dispositions</a>
<a href="energy_audit.md">Energy</a><a href="auto_attack_audit.md">Auto-attacks</a>
<a href="crit_model.md">Crit</a><a href="mana_regeneration.md">Mana</a>
<a href="windfury.md">Historical Windfury comparison</a></nav></details>
<h2 id="matrix">DPS matrix</h2>
<p>Sorted by each build’s peak mean DPS. Click a cell to load that exact race/build setup.
The <strong>Ranked builds</strong> picker also loads complete profiles.</p>
<p class="note">Hover cells for uncertainty and resource information. A dash means unavailable.
* Separate build variants; their assumptions are in the register and build details.</p>
<div class="matrix"><table><thead>""" + groups + "<tr>" + heading + \
        "</tr></thead><tbody>" + "".join(body) + """</tbody></table></div>
<footer class="note"><p>Unofficial beta simulator. Built on
<a href="https://github.com/wowsims/classic">WoWSims Classic</a> and
<a href="https://github.com/ElliotWood/Forever">ElliotWood/Forever</a>.
<a href="https://github.com/gunba/wow-forever-sim">Source and issues</a>.
Game icons via Wowhead.</p></footer></main><script>""" + QUESTIONS_SCRIPT + """</script></html>
"""
    if unenchanted:
        document = document.replace(
            "<!-- equipment-status -->",
            '<p style="border-left:4px solid #e3b65f;padding:.5rem 1rem;background:#373022">'
            '<strong>Starting equipment:</strong> these runs use unenchanted seed loadouts. '
            'They are not an optimized-gear comparison.</p>',
        )
    if invalid:
        document = document.replace(
            "<!-- invalid-results -->",
            '<p style="border-left:4px solid #e36565;padding:.5rem 1rem;background:#402529">'
            '<strong>Invalid Shaman results:</strong> the historical dual-wield entries are hidden. '
            'Shamans cannot dual wield. Downloadable raw data and chart images still contain those '
            'superseded runs; they are not a current ranking.</p>',
        )
    document = document.replace('<footer class="note">', tank_html + '<footer class="note">')
    if tank_rows:
        document = document.replace(
            '<a href="#questions">Questions &amp; coverage</a>',
            '<a href="#questions">Questions &amp; coverage</a><a href="#tanks">Tank metrics</a>')
        document = document.replace(
            '<!-- equipment-status -->',
            '<p class="note">Rows marked Tank include incoming attacks and modeled healing; '
            'their rage, threat and survival are scenario-dependent. See <a href="#tanks">tank metrics</a>.</p>')
    (args.output / "index.html").write_text(document)
    print(f"Review site staged at {args.output}")


if __name__ == "__main__":
    main()
