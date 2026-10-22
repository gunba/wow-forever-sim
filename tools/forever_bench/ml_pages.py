"""Learned (reinforcement-learning) results for the review portal: the matrix's first column and one page per build.

The reports come from tools/forever_rl/run_all.py (artifacts/ml/<key>.json). Each has two results: the policy network,
which chose every action in the engine and cannot be loaded into the simulator (the ceiling), and, when converted
(tools/forever_rl/learned_bench.py adds "profile"), an ordinary profile: the published setup with the chosen talents and
the published APL plus overrides converted from the network, benched exactly like the other cells.
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
.ml.nogain{background:#26242e;color:#a7a2b8}
"""

CAVEATS = (
    "The network chose every action in the simulator, with perfect reactions and exact knowledge of every timer, at the "
    "same 50 ms resolution the APL is checked at. It is this simulator's ceiling for the build under its current mechanics, "
    "not a target a player can be expected to reach, and it uses whatever the model allows. Pets act on their own. It "
    "cannot be loaded into the simulator; the profile is the part of it that converts to APL rules."
)


def load(directory=ML_DIR):
    out = {}
    if not directory.exists():
        return out
    for p in sorted(directory.glob("*.json")):
        r = json.loads(p.read_text())
        if "key" in r:  # one report per build; index.json and learned_results.json are not
            out[r["key"]] = r
    return out


def header_cell():
    return ('<th rowspan="2" scope="col" class="ml" title="What a learned network reached at the build\'s best race. '
            '&quot;sim&quot; opens the simmable profile converted from it (the published APL with overrides), where one keeps '
            'a gain; &quot;no APL&quot; means no conversion did yet.">Learned<br><span>best race</span></th>')


def network(r):
    """The network's DPS, its paired gain and SE, the published setup's DPS and the number of fights. The published setup
    is played by the engine's own APL path when measured (r["native"]); older reports compare with training's emulation."""
    if r.get("native"):
        N = r["native"]
        return N["network"], N["network_vs_published"], N["network_vs_published_se"], N["published"], N["fights"]
    L = r["learned"]
    return L["dps"], L["vs_published"], L["vs_published_se"], r["published"]["dps"], r["fights"]


def network_gained(r):
    """The network beat the published setup by at least two standard errors."""
    _, gain, se, _, _ = network(r)
    return gain > 2 * se


def gained(r):
    """A simmable profile beat the published setup, on paired fights and on the bench."""
    return bool((r.get("profile") or {}).get("gain"))


def cell(key, ml, simulator):
    r = ml.get(key)
    if r is None:
        return '<td class="ml unavailable">—</td>'
    dps, gain, se, pub, fights = network(r)
    race, Q = escape(r["race"]), r.get("profile")
    if not network_gained(r):
        title = (f'The learned network on {r["race"]} ({dps:.1f}; {gain:+.1f} ± {se:.1f} on {fights:,} paired fights) '
                 f'did not beat the published setup ({pub:.1f}).')
        return (f'<td class="ml nogain"><a href="ml/{escape(key)}.html" title="{escape(title)}">no gain</a>'
                f'<br><span>{race}</span></td>')
    title = f'Learned network on {r["race"]}: {dps:.1f} DPS; {gain:+.1f} ± {se:.1f} against the published setup ({pub:.1f}) on the same {fights:,} fights. '
    if gained(r):
        title += f'Simmable profile: {Q["dps"]:.1f} ({Q["dps"] - Q["published_dps"]:+.1f} against the published cell).'
        sim = (f'<a href="../{escape(simulator)}/?profile={escape(Q["id"])}" title="Open the simmable profile: the published APL '
               f'with overrides converted from the network">sim {Q["dps"]:.0f}</a>')
    else:
        title += "No APL reproduces this gain yet; its page lists the network's talents and actions."
        sim = "no APL"
    return (f'<td class="ml"><a href="ml/{escape(key)}.html" title="{escape(title)}">{dps:.0f}</a>'
            f'<br><span>{race} · {sim}</span></td>')


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
    P, L, A, Q = r["published"], r["learned"], r.get("actions") or {}, r.get("profile")
    race_file = r["race"].lower().replace(" ", "_")
    dps, gain, se, pub, fights = network(r)
    parts = [f"""<!doctype html><html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>{escape(class_name)} {escape(label)} · learned</title><style>{css}</style><main class="mlpage">
<p><a href="../index.html#matrix">← DPS matrix</a></p>
<h1>{escape(class_name)} · {escape(label)} · learned ({escape(r["race"])})</h1>
<p class="big">Network: {dps:.1f} DPS, {_signed(gain, se)} ({100 * gain / pub:+.1f}%) against the published setup ({pub:.1f})
on the same {fights:,} fights.</p>"""]
    if not network_gained(r):
        parts.append("<p><strong>The learner did not beat the published setup for this build:</strong> its best policy is "
                     "within noise of the APL or below it, so the published setup stands as the best known.</p>")
    elif gained(r):
        share = (Q["paired_vs_published"] / gain) if gain > 0 else 0
        parts.append(f"""<h2>Simmable profile</h2>
<p class="big">{Q["dps"]:.1f} DPS benched like the matrix, {_signed(Q["dps"] - Q["published_dps"])} against the published cell
({Q["published_dps"]:.1f}). <a href="../{escape(simulator)}/?profile={escape(Q["id"])}">Open it in the simulator</a></p>
<p>On {Q["paired_fights"]:,} paired fights it gains {_signed(Q["paired_vs_published"], Q["paired_vs_published_se"])}, {100 * share:.0f}% of the
network's gain on those fights. It uses the {escape(Q["talent_source"])} talents <code>{escape(Q["talents"])}</code> and the published APL
with {len(Q["overrides"])} entr{"ies" if len(Q["overrides"]) != 1 else "y"} added{f' and {len(Q["removed"])} removed' if Q.get("removed") else ""}.</p>""")
        if Q["overrides"]:
            parts.append("<h3>Added above the published APL</h3><p>Checked in order before the published entries. Each edit was "
                         "kept only if it raised DPS on paired fights.</p><ol>"
                         + "".join(f"<li><code>{escape(o)}</code></li>" for o in Q["overrides"]) + "</ol>")
        if Q.get("removed"):
            parts.append("<h3>Removed from the published APL</h3><ul>"
                         + "".join(f"<li><code>{escape(o)}</code></li>" for o in Q["removed"]) + "</ul>")
    else:
        best = (f" (its best attempt: {_signed(Q['paired_vs_published'], Q['paired_vs_published_se'])} on {Q['paired_fights']:,} paired fights)"
                if Q else "")
        parts.append(f"""<h2>No APL derived</h2>
<p>The conversion to APL rules did not reproduce this gain{best}, so there is no simmable profile for it. The network's talents and
actions below show what it does differently from the published setup; they are the place to start investigating.</p>""")
    parts.append(f"""<h2>The network</h2>
<p>With the published talents the network scores {L["published_talents_dps"]:.1f}; with its own talents, below, {L["dps"]:.1f}
(both on training's {r["fights"]:,} confirmation fights). <a href="../{escape(simulator)}/?profile={escape(r["key"])}__{escape(race_file)}">Open
the published {escape(r["race"])} setup</a> · <a href="{escape(r["key"])}.json">Report (JSON)</a></p>
<p class="note">{escape(CAVEATS)} <a href="../learned-rotations.md">How these were made</a></p>
<h2>Talents</h2>
<p>Published <code>{escape(P["talents"])}</code><br>Network <code>{escape(L["talents"])}</code></p>""")
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
                     f'network on its own). Where both had the same choice, the network took the APL\'s '
                     f'{100 * A["agreement"]:.1f}% of the time.</p><table><thead><tr><th>Action</th><th>APL</th><th>Network</th><th>Change</th>'
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
