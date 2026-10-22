"""Learned (reinforcement-learning) results for the review portal: the matrix's first column and one page per build.

The reports come from tools/forever_rl/run_all.py (artifacts/ml/<key>.json). A learned result is a policy network that
chose every action in the engine; it is not an APL and cannot be loaded into the simulator.
"""
import json
import shutil
from html import escape
from pathlib import Path

ML_DIR = Path("artifacts/ml")

ML_CSS = """
.ml{background:#2c2440;color:#e4d7ff;white-space:nowrap}.ml span{font-size:.75rem;color:#b9a9e6;font-weight:400}
thead th.ml{text-align:center;min-width:0}
.mlpage table{width:auto;min-width:40%}.mlpage td,.mlpage th{text-align:left}.mlpage td.n{text-align:right}
.mlpage .cols{display:flex;gap:3rem;flex-wrap:wrap}.mlpage .cols>div{flex:1;min-width:320px}
.mlpage .big{font-size:1.4rem}.pos{color:#9be3b0}.neg{color:#f2a3a3}
"""

CAVEATS = (
    "The learned result is a neural network that chose every action in the simulator, with perfect reactions and exact "
    "knowledge of every timer, at the same 50 ms resolution the APL is checked at. It is this simulator's ceiling for the "
    "build under its current mechanics, not a target a player can be expected to reach, and it uses whatever the model "
    "allows. Pets act on their own. It cannot be loaded into the simulator: it is not an APL."
)


def load(directory=ML_DIR):
    out = {}
    if not directory.exists():
        return out
    for p in sorted(directory.glob("*.json")):
        if p.name != "index.json":
            r = json.loads(p.read_text())
            out[r["key"]] = r
    return out


def header_cell():
    return '<th rowspan="2" scope="col" class="ml" title="Rotation and talents learned in the simulator, at the build\'s best race">Learned<br><span>best race</span></th>'


def cell(key, ml):
    r = ml.get(key)
    if r is None:
        return '<td class="ml unavailable">—</td>'
    L = r["learned"]
    title = (f'Learned rotation and talents on {r["race"]}: {L["dps"]:.1f} DPS; '
             f'{L["vs_published"]:+.1f} ± {L["vs_published_se"]:.1f} against the published rotation and talents on the same '
             f'{r["fights"]:,} fights. Not an APL.')
    return (f'<td class="ml"><a href="ml/{escape(key)}.html" title="{escape(title)}">{L["dps"]:.0f}</a>'
            f'<br><span>{escape(r["race"])}</span></td>')


def _name(action):
    for prefix in ("cast:", "macro:"):
        if action.startswith(prefix):
            action = action[len(prefix):]
    if action.endswith("#1"):
        # Tag 1 is the on-next-swing queue of Heroic Strike and Cleave.
        action = action[:-2] + " (queue)"
    if action.startswith("unqueue:"):
        action = "Unqueue " + action[len("unqueue:"):]
    if action.startswith("wait:"):
        action = "Wait " + action[len("wait:"):]
    return action


def _num(v, fmt="{:.0f}"):
    return "—" if v is None or v != v else fmt.format(v)


def _signed(v, se=None, digits=1):
    cls = "pos" if v > 0 else "neg" if v < 0 else ""
    s = f"{v:+.{digits}f}" + (f" ± {se:.{digits}f}" if se is not None else "")
    return f'<span class="{cls}">{s}</span>'


def page(r, label, class_name, simulator, css):
    P, L, A = r["published"], r["learned"], r.get("actions") or {}
    race_file = r["race"].lower().replace(" ", "_")
    gain_pct = 100 * L["vs_published"] / P["dps"] if P["dps"] else 0
    rot_gain = L["published_talents_vs_published"]
    parts = [f"""<!doctype html><html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>{escape(class_name)} {escape(label)} · learned</title><style>{css}</style><main class="mlpage">
<p><a href="../index.html#matrix">← DPS matrix</a></p>
<h1>{escape(class_name)} · {escape(label)} · learned ({escape(r["race"])})</h1>
<p class="big">{L["dps"]:.1f} DPS, {_signed(L["vs_published"], L["vs_published_se"])} ({gain_pct:+.1f}%) against the published
rotation and talents ({P["dps"]:.1f}) on the same {r["fights"]:,} fights.</p>
<p>With the published talents the learned rotation scores {L["published_talents_dps"]:.1f} ({_signed(rot_gain, L["published_talents_vs_published_se"])});
the rest comes from the talents below. <a href="../{escape(simulator)}/?profile={escape(r["key"])}__{escape(race_file)}">Open the published
{escape(r["race"])} setup</a> · <a href="{escape(r["key"])}.json">Report (JSON)</a></p>
<p class="note">{escape(CAVEATS)} <a href="../learned-rotations.md">How these were made</a></p>
<h2>Talents</h2>
<p>Published <code>{escape(P["talents"])}</code><br>Learned <code>{escape(L["talents"])}</code></p>"""]
    if r["talent_changes"]:
        parts.append("<table><thead><tr><th>Talent</th><th>Tree</th><th>Published</th><th>Learned</th></tr></thead><tbody>")
        for c in r["talent_changes"]:
            parts.append(f'<tr><td>{escape(c["talent"])}</td><td>{escape(c["tree"])}</td><td class="n">{c["published"]}</td>'
                         f'<td class="n">{c["learned"]}</td></tr>')
        parts.append("</tbody></table>")
    else:
        parts.append("<p>The search kept the published talents.</p>")
    if A:
        rows = [x for x in A["actions"] if x["action"] != "wait" and (x["apl"] >= 0.05 or x["policy"] >= 0.05)]
        rows.sort(key=lambda x: -(x["policy"] + x["apl"]))
        execute = any(x["apl_execute"] > 0 or x["policy_execute"] > 0 for x in rows)
        rage = any((x.get("policy_rage") or 0) > 0 for x in rows)
        parts.append(f'<h2>Actions</h2><p>Per fight, over {A["fights"]:,} fights of each (the APL on the published talents, the '
                     f'learned rotation on its own). Where both had the same choice, the learned rotation took the APL\'s '
                     f'{100 * A["agreement"]:.1f}% of the time.</p><table><thead><tr><th>Action</th><th>APL</th><th>Learned</th><th>Change</th>'
                     + ("<th>APL in execute</th><th>Learned in execute</th>" if execute else "")
                     + ("<th>APL rage</th><th>Learned rage</th>" if rage else "") + "</tr></thead><tbody>")
        for x in rows:
            parts.append(f'<tr><td>{escape(_name(x["action"]))}</td><td class="n">{x["apl"]:.1f}</td><td class="n">{x["policy"]:.1f}</td>'
                         f'<td class="n">{_signed(x["policy"] - x["apl"])}</td>'
                         + (f'<td class="n">{x["apl_execute"]:.1f}</td><td class="n">{x["policy_execute"]:.1f}</td>' if execute else "")
                         + (f'<td class="n">{_num(x.get("apl_rage"))}</td><td class="n">{_num(x.get("policy_rage"))}</td>' if rage else "")
                         + "</tr>")
        parts.append("</tbody></table>")
        if A.get("disagreements"):
            parts.append("<h2>Where it differs</h2><p>The most common points where the APL would have done one thing and the learned "
                         "rotation did another, per fight (on the learned talents).</p><table><thead><tr><th>APL would</th>"
                         "<th>Learned did</th><th>Per fight</th></tr></thead><tbody>")
            for d in A["disagreements"][:12]:
                parts.append(f'<tr><td>{escape(_name(d["apl_would"]))}</td><td>{escape(_name(d["policy_did"]))}</td>'
                             f'<td class="n">{d["per_fight"]:.1f}</td></tr>')
            parts.append("</tbody></table>")
        if A.get("cooldowns"):
            parts.append("<h2>Cooldowns</h2><p>Mean fight time of each use, in seconds (share of fights with that use).</p>"
                         "<table><thead><tr><th>Action</th><th>APL</th><th>Learned</th></tr></thead><tbody>")
            for c in A["cooldowns"]:
                fmt = lambda uses: " · ".join(f'{u["mean_s"]:.0f} s ({100 * u["share"]:.0f}%)' for u in uses) or "—"
                parts.append(f'<tr><td>{escape(_name(c["action"]))}</td><td>{fmt(c["apl"])}</td><td>{fmt(c["policy"])}</td></tr>')
            parts.append("</tbody></table>")
        tl = A.get("timeline") or {}
        if tl.get("policy"):
            parts.append('<h2>The first minute</h2><p>One fight, the same seed for both.</p><div class="cols">')
            for title, key in (("APL", "apl"), ("Learned", "policy")):
                parts.append(f"<div><h3>{title}</h3><table><thead><tr><th>Time</th><th>Action</th>"
                             + ("<th>Rage</th>" if rage else "") + "</tr></thead><tbody>")
                for e in tl.get(key, [])[:120]:
                    parts.append(f'<tr><td class="n">{e["t"]:.2f}</td><td>{escape(_name(e["action"]))}</td>'
                                 + (f'<td class="n">{e["rage"]}</td>' if rage else "") + "</tr>")
                parts.append("</tbody></table></div>")
            parts.append("</div>")
    T = r["training"]
    parts.append(f'<h2>Training</h2><p>{T["decisions"] / 1e6:.0f} million decisions over {T["fights"]:,} fights and '
                 f'{T["builds_seen"]:,} talent builds in {T["minutes"]:.0f} minutes ({T["decisions_per_fight"]:.0f} decisions a fight); '
                 f'stopped: {escape(T["stop"])}. Method: <a href="https://github.com/gunba/wow-forever-sim/blob/forever/tools/forever_rl/README.md">'
                 f'tools/forever_rl</a>.</p></main></html>')
    return "".join(parts)


def write(ml, output, builds, sim_paths, css, directory=ML_DIR):
    """Pages and reports under output/ml for every build with a learned result."""
    if not ml:
        return
    d = output / "ml"
    d.mkdir(parents=True, exist_ok=True)
    meta = {key: (class_name, label) for key, class_name, label, icon in builds}
    for key, r in ml.items():
        class_name, label = meta.get(key, ("", key))
        simulator = sim_paths.get(key, class_name.lower())
        (d / f"{key}.html").write_text(page(r, label, class_name, simulator, css))
        shutil.copyfile(directory / f"{key}.json", d / f"{key}.json")
    if (directory / "index.json").exists():
        shutil.copyfile(directory / "index.json", d / "index.json")
