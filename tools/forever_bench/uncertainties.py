#!/usr/bin/env python3
"""Validate and render the canonical mechanics/coverage register."""

import argparse
from collections import Counter
from html import escape
import json
from pathlib import Path
from urllib.parse import quote

ROOT = Path(__file__).resolve().parents[2]
REGISTER = ROOT / "docs/uncertainties.json"
SOURCE_ROOT = "https://github.com/gunba/wow-forever-sim/blob/forever/"
STATUSES = {
    "needs-evidence": "Needs evidence",
    "needs-code": "Implementation gap",
    "model-choice": "Model assumption",
    "provisional": "Accepted working answer",
    "resolved": "Resolved",
    "out-of-scope": "Outside current scenarios",
}
KINDS = {"mechanic", "implementation", "scenario", "data"}
STAGES = {
    "current-beta": "Current beta · level 30",
    "level30": "Level-30 beta",
    "launch": "Beyond level 30 / launch",
    "offline": "Sources / code",
}
DISPOSITIONS = {
    "evidence-answer": "Source-backed answer",
    "credible-answer": "Credible working answer",
    "already-resolved": "Previously resolved",
    "implementation": "Known implementation work",
    "partial": "Partly answered",
    "experiment": "Game test needed",
    "source-gap": "Source gap",
    "model-choice": "Scenario decision",
    "not-material": "Outside current scenarios",
}
PRIORITIES = {"high", "medium", "low"}
ACTIVE = {"needs-evidence", "needs-code", "model-choice"}


def apply_review(data, review):
    """Reconcile a complete adjudication packet with stable register records."""
    existing = {item["id"]: item for item in data["items"]}
    additions = review["coverage_additions"]
    if set(existing) - set(additions) != set(review["entries"]):
        raise ValueError("Review must cover every original register ID exactly")
    fallback = {
        "partial": "needs-evidence", "experiment": "needs-evidence",
        "source-gap": "needs-evidence", "implementation": "needs-code",
        "evidence-answer": "resolved", "already-resolved": "resolved",
        "credible-answer": "provisional", "model-choice": "model-choice",
        "not-material": "out-of-scope",
    }
    items = []
    for ident, finding in (review["entries"] | additions).items():
        item = dict(existing.get(ident, additions.get(ident, {}).get("register", {})))
        previous = item.get("review", {}).get("previous_status", item.get("status"))
        item["id"] = ident
        item.update(finding.get("register_update", {}))
        item.pop("access", None)
        item.pop("resolution", None)
        status = finding.get("proposed_status", fallback[finding["disposition"]])
        if finding["disposition"] == "credible-answer" and status == "resolved":
            status = "provisional"
        item["status"] = status
        item["priority"] = finding.get("proposed_priority", item["priority"])
        item["sources"] = list(dict.fromkeys(item.get("sources", []) + finding["sources"]))
        item["related"] = list(dict.fromkeys(item.get("related", []) + finding.get("related", [])))
        item["review"] = {key: finding[key] for key in (
            "disposition", "confidence", "answer", "remaining", "implementation", "tests")}
        item["review"]["previous_status"] = previous
        if finding.get("code_ready"):
            item["review"]["code_ready"] = finding["code_ready"]
        items.append(item)
    data["schema_version"] = 2
    data["reviewed_revision"] = review["baseline_revision"]
    data["review"] = {
        "date": review["review_date"],
        "original_entries": review["baseline_entries"],
        "reviewed_original_entries": len(review["entries"]),
        "added_entries": len(additions),
        "before": review["baseline_counts"],
        "stage_definitions": review["stage_definitions"],
        "evidence": [
            "docs/question_review.md", "artifacts/question_review/review.json",
            "artifacts/question_review/sources.json",
            "artifacts/question_review/class_evidence.json",
            "artifacts/question_review/scenario_data_evidence.json",
            "artifacts/question_review/seed_ranges.json",
        ],
    }
    data["items"] = items
    return data


def load_register(path=REGISTER):
    data = json.loads(Path(path).read_text())
    if data["schema_version"] != 2:
        raise ValueError("Expected the adjudicated register schema")
    ids = set()
    for item in data["items"]:
        ident = item["id"]
        if ident in ids:
            raise ValueError(f"Duplicate uncertainty ID: {ident}")
        ids.add(ident)
        for field in ("title", "question", "current_model", "impact"):
            if not isinstance(item.get(field), str) or not item[field].strip():
                raise ValueError(f"{ident}: missing {field}")
        if item["category"] not in data["categories"]:
            raise ValueError(f"{ident}: unknown category")
        if item["status"] not in STATUSES or item["kind"] not in KINDS:
            raise ValueError(f"{ident}: unknown status/kind")
        if item["priority"] not in PRIORITIES:
            raise ValueError(f"{ident}: unknown priority")
        review = item["review"]
        if review["disposition"] not in DISPOSITIONS or review["confidence"] not in {"high", "moderate", "low"}:
            raise ValueError(f"{ident}: unknown disposition/confidence")
        for field in ("answer", "remaining", "implementation"):
            if not isinstance(review.get(field), str) or not review[field].strip():
                raise ValueError(f"{ident}: missing reviewed {field}")
        for test in review["tests"]:
            if test["stage"] not in STAGES or not test["requires"].strip() or not test["measure"].strip():
                raise ValueError(f"{ident}: incomplete test prerequisites/observations")
        if item["status"] == "needs-evidence" and not review["tests"]:
            raise ValueError(f"{ident}: evidence question has no answer route")
        if not item.get("sources"):
            raise ValueError(f"{ident}: no evidence/code references")
        for source in item["sources"]:
            if source.startswith("https://"):
                continue
            if not (ROOT / source.split("#")[0]).is_file():
                raise ValueError(f"{ident}: missing source {source}")
        for other in item.get("related", []):
            if other == ident:
                raise ValueError(f"{ident}: self-reference")
    for item in data["items"]:
        if set(item.get("related", [])) - ids:
            raise ValueError(f'{item["id"]}: unknown related ID')
    summary = data["review"]
    if summary["reviewed_original_entries"] != summary["original_entries"]:
        raise ValueError("Review coverage is incomplete")
    if summary["original_entries"] + summary["added_entries"] != len(ids):
        raise ValueError("Review accounting does not match register")
    if sum(summary["before"].values()) != summary["original_entries"]:
        raise ValueError("Baseline status accounting does not match register")
    return data


def source_url(source):
    return source if source.startswith("https://") else SOURCE_ROOT + quote(source, safe="/#")


def ordered_items(data):
    priority = {"high": 0, "medium": 1, "low": 2}
    return sorted(data["items"], key=lambda x: (
        x["status"] not in ACTIVE,
        priority[x["priority"]], x["id"],
    ))


def work_queue(data, stage):
    return [(item, test) for item in ordered_items(data) if item["status"] in ACTIVE
            for test in item["review"]["tests"] if test["stage"] == stage]


def stages_for(item):
    return list(dict.fromkeys(test["stage"] for test in item["review"]["tests"]))


def accounting(data):
    current = Counter(item["status"] for item in data["items"])
    original = Counter(item["status"] for item in data["items"]
                       if item["review"]["previous_status"] is not None)
    return [(key, data["review"]["before"].get(key, 0), original[key], current[key])
            for key in STATUSES]


def markdown(data):
    meta = data["review"]
    lines = [
        "# Simulator questions and coverage",
        "",
        "<!-- Generated by tools/forever_bench/uncertainties.py. Edit uncertainties.json. -->",
        "",
        f"Reviewed against `{data['reviewed_revision']}`. {data['scope']}",
        "",
        "This is the current register. Earlier audits preserve historical findings, not current task status. "
        "Research answers and implementation fixes are separate: **combat mechanics and benchmark results have not changed in this review**.",
        "",
        f"**Coverage:** {meta['reviewed_original_entries']}/{meta['original_entries']} original entries reviewed; "
        f"{meta['added_entries']} additional coverage gaps. Reviewed {meta['date']}.",
        "",
        "## Before and after",
        "",
        "| Status | Before | Original entries now | Including additions |",
        "|---|---:|---:|---:|",
        *[f"| {STATUSES[key]} | {before} | {original} | {total} |"
          for key, before, original, total in accounting(data)],
        "",
        "A working answer is credible enough for an explicitly qualified model, not verified server behavior. "
        "Confidence applies to the stated finding, not every effect of that class. "
        "An implementation gap is not a request for a game experiment.",
        "",
        "## Prioritized follow-up",
        "",
        "Each queue is ordered by priority, then stable ID. A question can have partial tests at several stages; "
        "these queue counts must not be added together as distinct questions. Accepted-answer corroboration and "
        "out-of-scenario work remain inside their records, not in these pending queues.",
        "",
    ]
    for stage, label in STAGES.items():
        queue = work_queue(data, stage)
        lines += [f"### {label}", "", meta["stage_definitions"][stage], ""]
        for item, test in queue:
            lines += [
                f'- **{item["priority"].title()} · [{item["id"]} — {item["title"]}](#{item["id"].lower()})**',
                f'  - Prerequisites: {test["requires"]}',
                f'  - Observe / verify: {test["measure"]}',
            ]
        lines += [""]
    lines += [
        "## Evidence and implementation follow-up",
        "",
        "The individual records identify source-ready corrections separately from assumptions. "
        "No item marked Implementation gap has been silently fixed by this research publication.",
        "",
        *[f'- **[{item["id"]} — {item["title"]}](#{item["id"].lower()})**: {item["review"]["code_ready"]}'
          for item in ordered_items(data) if item["review"].get("code_ready")],
        "",
        *[f"- [{source}]({source_url(source)})" for source in meta["evidence"]],
        "",
        "## Updating an answer",
        "",
        "Edit `docs/uncertainties.json`, retaining the ID and qualified legacy references. "
        "Update the reviewed answer, confidence, remaining scope and concrete tests; distinguish an answered "
        "source question from outstanding implementation. Regenerate with "
        "`python3 tools/forever_bench/uncertainties.py`.",
        "",
        "Old T50, T51 and T52 labels were reused for different topics. The qualified aliases below "
        "preserve those references without conflating their answers. Private chat and identifying logs stay local.",
        "",
        "## Inventory coverage",
        "",
        *["- " + note for note in data["coverage"]],
        "",
    ]
    for category, label in data["categories"].items():
        lines += [f"## {label}", ""]
        for item in ordered_items(data):
            if item["category"] != category:
                continue
            lines += [
                f'<a id="{item["id"].lower()}"></a>',
                "",
                f'### {item["id"]} — {item["title"]}',
                "",
                f'**{STATUSES[item["status"]]} · {item["priority"].title()} priority · '
                f'{item["review"]["confidence"].title()} confidence**',
                "",
                f'**Disposition:** {DISPOSITIONS[item["review"]["disposition"]]}',
                "",
                f'**Question:** {item["question"]}',
                "",
                f'**Finding:** {item["review"]["answer"]}',
                "",
                f'**Current implementation / scenario:** {item["current_model"]}',
                "",
                f'**Impact:** {item["impact"]}',
                "",
                f'**Remaining scope:** {item["review"]["remaining"]}',
                "",
                f'**Implementation / model follow-up:** {item["review"]["implementation"]}',
                "",
            ]
            for test in item["review"]["tests"]:
                qualifier = "" if item["status"] in ACTIVE else "Optional / conditional · "
                lines += [
                    f'**{qualifier}{STAGES[test["stage"]]}**',
                    f'- Prerequisites: {test["requires"]}',
                    f'- Observe / verify: {test["measure"]}', "",
                ]
            lines += ["**References:** " + " · ".join(
                f"[{source}]({source_url(source)})" for source in item["sources"]), ""]
            if item.get("legacy"):
                lines += ["**Earlier references:** " + "; ".join(item["legacy"]), ""]
            if item.get("related"):
                lines += ["**Related:** " + ", ".join(item["related"]), ""]
    return "\n".join(lines)


def render_html(data):
    """Self-contained, progressively enhanced register for the review page."""
    meta = data["review"]
    overview = "".join(
        f"<tr><th>{STATUSES[key]}</th><td>{before}</td><td>{original}</td><td>{total}</td></tr>"
        for key, before, original, total in accounting(data))
    plans = []
    buttons = []
    for stage, label in STAGES.items():
        queue = work_queue(data, stage)
        count = len({item["id"] for item, _ in queue})
        buttons.append(f'<button type="button" data-stage="{stage}">{label} · {count}</button>')
        tasks = "".join(
            f'<li><a href="#{item["id"]}">{item["id"]} — {escape(item["title"])}</a> '
            f'<strong>{item["priority"].title()}</strong>'
            f'<p><strong>Requires:</strong> {escape(test["requires"])}</p>'
            f'<p><strong>Observe / verify:</strong> {escape(test["measure"])}</p></li>'
            for item, test in queue)
        plans.append(f'<details><summary>{label} · {count} questions</summary>'
                     f'<p>{escape(meta["stage_definitions"][stage])}</p><ol>{tasks}</ol></details>')
    corrections = "".join(
        f'<li><a href="#{item["id"]}">{item["id"]} — {escape(item["title"])}</a>: '
        f'{escape(item["review"]["code_ready"])}</li>'
        for item in ordered_items(data) if item["review"].get("code_ready"))
    groups = []
    for category, label in data["categories"].items():
        cards = []
        items = [x for x in ordered_items(data) if x["category"] == category]
        for item in items:
            legacy = "; ".join(item.get("legacy", []))
            review = item["review"]
            routes = stages_for(item)
            search = json.dumps(item, ensure_ascii=False)
            links = " · ".join(
                f'<a href="{escape(source_url(src), quote=True)}">{escape(src)}</a>'
                for src in item["sources"])
            tests = "".join(
                f'<li><strong>{STAGES[test["stage"]]}</strong>'
                f'<p><strong>Requires:</strong> {escape(test["requires"])}</p>'
                f'<p><strong>Observe / verify:</strong> {escape(test["measure"])}</p></li>'
                for test in review["tests"])
            tests_label = "Answer routes" if item["status"] in ACTIVE else "Optional / conditional follow-up"
            related = " · ".join(f'<a href="#{other}">{other}</a>' for other in item.get("related", []))
            cards.append(
                f'<details class="question-card" id="{item["id"]}" '
                f'data-status="{item["status"]}" data-priority="{item["priority"]}" '
                f'data-access="{" ".join(routes)}" data-search="{escape(search.casefold(), quote=True)}">'
                f'<summary><span class="question-id">{item["id"]}</span> {escape(item["title"])} '
                f'<span class="question-state">{STATUSES[item["status"]]} · {item["priority"]}</span></summary>'
                f'<p><strong>Question:</strong> {escape(item["question"])}</p>'
                f'<p class="question-finding"><strong>{DISPOSITIONS[review["disposition"]]} '
                f'· {review["confidence"]} confidence:</strong> {escape(review["answer"])}</p>'
                f'<p><strong>Current implementation / scenario:</strong> {escape(item["current_model"])}</p>'
                f'<p><strong>Impact:</strong> {escape(item["impact"])}</p>'
                f'<p><strong>Remaining scope:</strong> {escape(review["remaining"])}</p>'
                f'<p><strong>Implementation / model follow-up:</strong> {escape(review["implementation"])}</p>'
                + (f'<details class="question-tests"><summary>{tests_label}</summary><ul>{tests}</ul></details>' if tests else "")
                +
                f'<p class="question-links">{links}</p>'
                + (f'<p class="note">Earlier references: {escape(legacy)}</p>' if legacy else "")
                + (f'<p class="note">Related: {related}</p>' if related else "")
                + f'<a href="#{item["id"]}" class="note">Link to {item["id"]}</a></details>'
            )
        active = sum(x["status"] in ACTIVE for x in items)
        groups.append(
            f'<details class="question-group"><summary>{escape(label)} '
            f'<span class="question-count">{active} active / {len(items)} total</span></summary>'
            + "".join(cards) + "</details>")
    status_options = "".join(f'<option value="{key}">{label}</option>' for key, label in STATUSES.items())
    access_options = "".join(f'<option value="{key}">{label}</option>' for key, label in STAGES.items())
    return (
        '<section id="questions" aria-labelledby="question-heading"><h2 id="question-heading">Questions &amp; coverage</h2>'
        f'<p><strong>{meta["reviewed_original_entries"]}/{meta["original_entries"]} questions reviewed; '
        f'{meta["added_entries"]} new coverage gaps.</strong> '
        'Sources, remaining tests and implementation work are separated below.</p>'
        '<p class="note">This review does not change combat mechanics or benchmark numbers. '
        'A source-backed correction listed here is not necessarily implemented. '
        '“Beyond level 30” means outside announced beta access; move it earlier if access expands.</p>'
        '<details class="question-overview"><summary>Review accounting and evidence</summary>'
        '<div class="question-accounting"><table><thead><tr><th>Status</th><th>Before</th>'
        '<th>Original entries now</th><th>Including additions</th></tr></thead>'
        '<tbody>' + overview + '</tbody></table></div>'
        '<p>Accepted working answers are qualified, not verified server facts. Confidence applies to the stated finding. '
        '<a href="question_review.md">Review summary</a> · '
        '<a href="question_review/review.json">Adjudications</a> · '
        '<a href="question_review/sources.json">Source ledger</a></p></details>'
        '<details class="question-plan"><summary>Prioritized test and implementation plan</summary>'
        '<p>High priority first. A question may have partial checks at several stages. '
        'Accepted-answer corroboration is optional and remains within its record.</p>'
        + "".join(plans) + '</details>'
        '<details class="question-plan"><summary>Source-supported implementation follow-up · not applied</summary>'
        '<p>These specific changes do not need an unknown server rule to be invented. '
        'Some records also contain separate unresolved behavior.</p><ol>'
        + corrections + '</ol></details>'
        '<div class="question-shortcuts">' + "".join(buttons) + '</div>'
        '<form class="question-filters" role="search" onsubmit="return false">'
        '<label>Find a question <input id="question-search" type="search" placeholder="Class, spell, ID…"></label>'
        '<label>Status <select id="question-status"><option value="active">Active</option>'
        '<option value="all">All statuses</option>' + status_options + '</select></label>'
        '<label>Priority <select id="question-priority"><option value="all">All priorities</option>'
        '<option value="high">High</option><option value="medium">Medium</option><option value="low">Low</option></select></label>'
        '<label>Answer route <select id="question-access"><option value="all">All routes</option>'
        + access_options + '</select></label></form>'
        '<p id="question-results" class="note" aria-live="polite"></p>'
        + "".join(groups)
        + '<p class="note"><a href="uncertainties.md">Complete readable register</a> · '
        '<a href="uncertainties.json">Structured register</a></p></section>'
    )


CSS = """
#questions{margin:1.5rem 0 2rem}
.question-filters{display:flex;flex-wrap:wrap;gap:.8rem;align-items:end}
.question-filters label{display:flex;flex-direction:column;font-size:.9rem;gap:.25rem}
.question-filters input,.question-filters select{background:#242834;color:#eee;border:1px solid #68707e;padding:.5rem;border-radius:4px}
.question-shortcuts{display:flex;flex-wrap:wrap;gap:.5rem;margin:1rem 0}
.question-shortcuts button{background:#243b52;border:1px solid #527797;color:#eee;padding:.5rem .7rem;border-radius:5px;cursor:pointer}
.question-overview,.question-plan{margin:.75rem 0}
.question-overview>summary,.question-plan>summary,.question-plan details>summary,.question-tests>summary{cursor:pointer;padding:.4rem 0}
.question-accounting{overflow-x:auto}.question-accounting table{width:auto;min-width:440px}
.question-plan details{margin:.65rem 1rem}.question-plan li{margin:.8rem 0;max-width:1050px}
.question-plan li p{margin:.25rem 0}.question-finding{border-left:3px solid #779dc0;padding-left:.75rem}
.question-group{margin:.5rem 0;border:1px solid #343945;border-radius:5px}
.question-group>summary{padding:.65rem;cursor:pointer;font-weight:600}
.question-count,.question-state{color:#acb6c8;font-weight:400;font-size:.82rem;margin-left:.6rem}
.question-card{padding:.55rem .9rem;border-top:1px solid #343945}
.question-card>summary{cursor:pointer}
.question-card p{margin:.6rem 0;max-width:1150px}
.question-id{color:#a9d8ff;font-variant-numeric:tabular-nums}
.question-links{font-size:.8rem;overflow-wrap:anywhere}
.question-card:target{outline:2px solid #a9d8ff;outline-offset:-2px}
[hidden]{display:none!important}
.resource-group{margin:.7rem 0}.resource-group>summary{cursor:pointer}
"""

SCRIPT = """
(() => {
 const root = document.querySelector('#questions');
 if (!root) return;
 const cards = [...root.querySelectorAll('.question-card')];
 const groups = [...root.querySelectorAll('.question-group')];
 const search = root.querySelector('#question-search'), status = root.querySelector('#question-status');
 const priority = root.querySelector('#question-priority'), access = root.querySelector('#question-access');
 function filter() {
   const terms = search.value.toLowerCase().trim().split(/\\s+/).filter(Boolean);
   let n = 0;
   for (const card of cards) {
     const d = card.dataset;
     const matches = (status.value === 'all' || (status.value === 'active'
       ? ['needs-evidence','needs-code','model-choice'].includes(d.status) : status.value === d.status))
       && (priority.value === 'all' || priority.value === d.priority)
       && (access.value === 'all' || d.access.split(' ').includes(access.value))
       && terms.every(term => d.search.includes(term));
     card.hidden = !matches;
     if (matches) n++;
   }
   for (const group of groups) {
     const visible = [...group.querySelectorAll('.question-card')].filter(c => !c.hidden).length;
     group.hidden = !visible;
     group.querySelector('.question-count').textContent = `${visible} shown`;
     if ((terms.length || access.value !== 'all') && visible) group.open = true;
   }
   root.querySelector('#question-results').textContent = `${n} of ${cards.length} entries shown`;
 }
 function revealLink() {
   const card = cards.find(c => '#' + c.id === location.hash);
   if (!card) return;
   status.value = 'all'; priority.value = 'all'; access.value = 'all'; search.value = '';
   filter(); card.open = true; card.closest('.question-group').open = true;
   card.scrollIntoView({block:'nearest'});
 }
 for (const control of [search,status,priority,access]) control.addEventListener('input', filter);
 for (const button of root.querySelectorAll('[data-stage]')) button.addEventListener('click', () => {
   status.value = 'active'; priority.value = 'all'; search.value = '';
   access.value = button.dataset.stage; filter();
 });
 window.addEventListener('hashchange', revealLink);
 filter(); revealLink();
})();
"""


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--check", action="store_true", help="check the generated Markdown without changing it")
    parser.add_argument("--apply-review", type=Path, help="reconcile a complete adjudication packet before rendering")
    args = parser.parse_args()
    if args.apply_review:
        if args.check:
            parser.error("--check cannot apply a review")
        data = apply_review(json.loads(REGISTER.read_text()), json.loads(args.apply_review.read_text()))
        REGISTER.write_text(json.dumps(data, indent=2, ensure_ascii=False) + "\n")
    data = load_register()
    output = ROOT / "docs/uncertainties.md"
    text = markdown(data) + "\n"
    if args.check:
        if not output.exists() or output.read_text() != text:
            raise SystemExit("docs/uncertainties.md is stale; regenerate it")
    else:
        output.write_text(text)
    print(f'{len(data["items"])} entries; ' + ", ".join(
        f"{key}={value}" for key, value in sorted(Counter(x["status"] for x in data["items"]).items())))


if __name__ == "__main__":
    main()
