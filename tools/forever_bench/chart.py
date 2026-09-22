#!/usr/bin/env python3
"""Render a race-by-spec DPS matrix from forever_bench results."""

import argparse
import json
from pathlib import Path

import matplotlib
matplotlib.use("Agg")
import matplotlib.pyplot as plt
from matplotlib.offsetbox import AnnotationBbox, OffsetImage
from matplotlib.patches import Rectangle
import numpy as np

from build_display import BUILDS, CLASS_COLORS, HYBRID_PARENTS, RACES
from sensitivity import DISPLAY_METRICS, format_metric, load_columns


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("results", type=Path)
    parser.add_argument("--output", type=Path)
    parser.add_argument("--faction", choices=["all", "alliance", "horde"], default="all")
    parser.add_argument("--icons", type=Path, default=Path("assets/img/spec_icons"))
    parser.add_argument("--sensitivity", type=Path)
    args = parser.parse_args()
    data = json.loads(args.results.read_text())
    columns = load_columns(args.sensitivity, args.results, args.faction)
    rows = [r for r in data["Results"] if args.faction == "all" or r["Faction"].lower() == args.faction]
    if not rows:
        raise ValueError("No matching results")
    available_races = {r["Race"] for r in rows}
    races = [r for r in RACES if r in available_races]
    keys = {r["Key"] for r in rows}
    builds = [b for b in BUILDS if b[0] in keys]
    if len(builds) != len(keys) or len(races) != len(available_races):
        raise ValueError("Missing display metadata")
    peaks = {key: max(r["DPS"] for r in rows if r["Key"] == key) for key in keys}
    builds.sort(key=lambda b: peaks[b[0]], reverse=True)
    lookup = {}
    for row in rows:
        pair = (row["Key"], row["Race"])
        if pair in lookup:
            raise ValueError(f"Duplicate result: {pair}")
        lookup[pair] = row
    roster = json.loads(Path("assets/db_inputs/forever_races.json").read_text())["races"]
    eligible = {r["name"]: set(r["classes"]) for r in roster}
    values = np.full((len(builds), len(races)), np.nan)
    for y, (key, cls, _, _) in enumerate(builds):
        for x, race in enumerate(races):
            row = lookup.get((key, race))
            if row:
                if cls not in eligible[race]:
                    raise ValueError(f"Unavailable combination in results: {key}/{race}")
                values[y, x] = row["DPS"]
            elif cls in eligible[race]:
                raise ValueError(f"Missing available combination: {key}/{race}")

    extra = len(DISPLAY_METRICS) if columns else 0
    width = len(races) + extra
    fig, ax = plt.subplots(figsize=(19 if columns else 15, max(6, len(builds) * .45 + 2.9)))
    fig.subplots_adjust(left=.27 if columns else .315, right=.925, top=.83, bottom=.135)
    cmap = plt.colormaps["YlGnBu"].copy()
    cmap.set_bad("#edf0f4")
    ceiling = float(np.nanmax(values))
    image = ax.imshow(values, cmap=cmap, vmin=0, vmax=ceiling, aspect="auto")
    labels = [r.replace(" ", "\n") for r in races]
    if columns:
        labels += ["Tier 1\ngain", "Gear\n+10%", "Scaling\namp."]
    ax.set_xlim(-.5, width-.5)
    ax.set_xticks(range(width), labels, fontsize=10)
    ax.xaxis.tick_top()
    ax.tick_params(axis="x", length=0, pad=9)
    ax.set_yticks([])
    ax.set_xticks(np.arange(-.5, width, 1), minor=True)
    ax.set_yticks(np.arange(-.5, len(builds), 1), minor=True)
    ax.grid(which="minor", color="white", linewidth=1.2)
    ax.tick_params(which="minor", bottom=False, left=False)
    for spine in ax.spines.values():
        spine.set_visible(False)

    groups = []
    for y, (key, cls, label, icon) in enumerate(builds):
        if groups and groups[-1][0] == cls:
            groups[-1][1].append(y)
        else:
            groups.append((cls, [y]))
        path = args.icons / f"{icon}.jpg"
        if not path.exists():
            raise FileNotFoundError(f"{path}; run tools/forever_bench/fetch_icons.py")
        box = AnnotationBbox(OffsetImage(plt.imread(path), zoom=.31), (-2.05, y),
                             frameon=False, annotation_clip=False)
        ax.add_artist(box)
        ax.text(-1.78, y, label + ("*" if key in HYBRID_PARENTS else ""),
                ha="left", va="center", fontsize=10, color="#253248", clip_on=False)
        best = float(np.nanmax(values[y]))
        for x, value in enumerate(values[y]):
            if np.isnan(value):
                ax.text(x, y, "—", ha="center", va="center", color="#9ba5b4", fontsize=11)
            else:
                ax.text(x, y, f"{value:.0f}", ha="center", va="center", fontsize=10,
                        color="white" if value > ceiling * .58 else "#15243b",
                        fontweight="bold" if value == best else "normal")
                if value == best:
                    ax.add_patch(Rectangle((x-.46, y-.43), .92, .86, fill=False,
                                           edgecolor="#f4a938", linewidth=1.4))
        if columns:
            for index, metric in enumerate(DISPLAY_METRICS):
                x = len(races) + index
                summary = columns[key][metric]
                if metric == "amplification":
                    value = summary["Amplification"]
                    color = "#475569" if value is None or abs(value - 1) <= summary["MonteCarlo95Bound"] else (
                        "#236343" if value > 1 else "#8b5939")
                else:
                    color = "#236343" if summary["GainPercent"] >= 0 else "#a44240"
                ax.add_patch(Rectangle((x-.5, y-.5), 1, 1, facecolor="#edf3ee", edgecolor="white"))
                ax.text(x, y, format_metric(metric, summary), ha="center", va="center", fontsize=10,
                        color=color, fontweight="bold")
    for cls, positions in groups:
        lo, hi = min(positions), max(positions)
        ax.text(-2.38, (lo+hi)/2, cls, ha="right", va="center", fontsize=10.5,
                fontweight="bold", color="#253248", clip_on=False)
        ax.add_patch(Rectangle((-2.30, lo-.46), .045, hi-lo+.92,
                               facecolor=CLASS_COLORS[cls], linewidth=0, clip_on=False))
        if lo:
            ax.axhline(lo-.5, color="#9aa6b6", linewidth=1.2)
    factions = {r["Race"]: r["Faction"] for r in rows}
    for faction, color in [("Alliance", "#244b80"), ("Horde", "#813b43")]:
        cols = [x for x, race in enumerate(races) if factions[race] == faction]
        if cols:
            mid = (min(cols)+max(cols))/2
            transform = ax.get_xaxis_transform()
            ax.add_patch(Rectangle((min(cols)-.5, 1.068), len(cols), .043, transform=transform,
                                   facecolor=color, linewidth=0, clip_on=False))
            ax.text(mid, 1.0895, faction, transform=transform, va="center",
                    ha="center", fontsize=11, fontweight="bold", color="white")
            icon = Path("assets/img/wowhead/icons") / f"{faction.lower()}.png"
            ax.add_artist(AnnotationBbox(OffsetImage(plt.imread(icon), zoom=1.6),
                                        (mid-.92, 1.0895), xycoords=transform,
                                        frameon=False, annotation_clip=False))
    if columns:
        ax.add_patch(Rectangle((len(races)-.5, 1.068), extra, .043, transform=ax.get_xaxis_transform(),
                               facecolor="#3e6152", linewidth=0, clip_on=False))
        ax.text(len(races)+(extra-1)/2, 1.0895, "Gains & scaling", transform=ax.get_xaxis_transform(),
                ha="center", va="center", fontsize=11, fontweight="bold", color="white")
        ax.axvline(len(races)-.5, color="#8d9f98", linewidth=2)
    cax = fig.add_axes([.94, .23, .01, .5])
    fig.colorbar(image, cax=cax, label="Damage per second")
    fig.suptitle("WoW Forever · DPS", y=.977, fontsize=23, fontweight="bold")
    tier = "Tier 1 bonuses enabled" if data["Tier1Bonuses"] else "Tier 1 override disabled"
    fig.text(.5, .942, f"Level 60 · {data['Duration']:g}s single target · {tier}", ha="center", fontsize=11)
    samples = "/".join(f"{n:,}" for n in sorted({r["Iterations"] for r in rows}))
    fig.text(.07, .085,
             f"{samples} iterations per result · rows ordered by peak mean DPS · outlined cell: row peak · — unavailable\n"
              "Crafted/dungeon/PvP equipment · paid shared-hit normalization · no world buffs",
             fontsize=9, color="#475569")
    if columns:
        fig.text(.07, .05,
                 "Equal race weights. Tier 1: on vs off. Gear: item stats + weapon damage; fixed enchants, procs and rotation; paid hit recalculated.\n"
                 "Scaling amp. = +50% gain / (5 × +10% gain). 1× linear; >1× accelerating; <1× flattening; — too small/noisy. Not proof of exponential growth.\n"
                 "Caps and resource thresholds affect curvature. Hypothetical upgrades, not future-item predictions.",
                 fontsize=8.5, color="#475569")
    fig.text(.07, .02,
             "* Hybrid rows depend on guardian, proc or rage assumptions; see build notes. All rows use a beta model. Game icons via Wowhead.",
             fontsize=8.5, color="#64748b")
    prefix = args.output or args.results.with_suffix("")
    prefix.parent.mkdir(parents=True, exist_ok=True)
    for extension in ("svg", "png"):
        path = prefix.with_suffix("." + extension)
        fig.savefig(path, dpi=150, facecolor="white", metadata={"Date": None} if extension == "svg" else None)
        if extension == "svg":
            path.write_text("\n".join(line.rstrip() for line in path.read_text().splitlines()) + "\n")
        print(path)
    plt.close(fig)


if __name__ == "__main__":
    main()
