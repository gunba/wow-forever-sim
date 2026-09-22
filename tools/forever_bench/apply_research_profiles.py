#!/usr/bin/env python3
"""Prepare benchmark players and APL presets from validated build recipes."""

import argparse
from copy import deepcopy
import gzip
import json
from pathlib import Path
import re

from build_display import expected_roster


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--archive", type=Path, default=Path("artifacts/research_builds/validation.json.gz"))
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument("--write-apls", action="store_true")
    args = parser.parse_args()
    with gzip.open(args.archive, "rt") as stream:
        data = json.load(stream)
    profiles = {(row["Key"], row["Race"]): deepcopy(row) for row in data["BaselineProfiles"]}
    definitions = {
        match[0]: match[1:] for match in re.findall(
            r'\{"([^"]+)", "[^"]+", "([^"]+)", "([^"]+)", "([^"]+)",',
            Path("tools/forever_bench/profiles.go").read_text(),
        )
    }
    for choice in data["Choices"]:
        if not choice["Selected"]:
            continue
        key, candidate = choice["Key"], choice["Candidate"]
        entries = [entry for entry in data["Comparisons"]
                   if entry["build"] == key and entry["candidate"] == candidate]
        players = {}
        for entry in entries:
            row = data["Runs"][entry["path"]]
            player = row["BaselinePlayer"]
            parent = profiles[choice["Parent"], row["Race"]]["BaselinePlayer"]
            left, right = deepcopy(player), deepcopy(parent)
            for item in (left, right):
                item.pop("rotation")
                item.pop("talentsString")
            if left != right:
                raise ValueError(f"{key}/{row['Race']}: changed non-talent/APL input")
            if row.get("Warnings") or abs(row["Hit"]["Balance"]) > 1e-6:
                raise ValueError(f"{key}/{row['Race']}: invalid validation result")
            players[row["Race"]] = player
        if set(players) != {race for build, race in expected_roster() if build == key}:
            raise ValueError(f"{key}: incomplete race coverage")
        variants = {(p["talentsString"], json.dumps(p["rotation"], sort_keys=True))
                    for p in players.values()}
        if len(variants) != 1:
            raise ValueError(f"{key}: expected one shared talent/APL preset")
        talents, _ = next(iter(variants))
        directory, preset, apl = definitions[key]
        source = Path(f"ui/{directory}/presets.ts").read_text()
        match = re.search(r"export const " + re.escape(preset) +
                          r"\s*=.*?talentsString:\s*'([0-9-]+)'", source, re.S)
        if not match or match[1] != talents:
            raise ValueError(f"{key}: talent preset does not match the selected recipe")
        if args.write_apls:
            Path(f"ui/{directory}/apls/{apl}.apl.json").write_text(
                json.dumps(next(iter(players.values()))["rotation"], indent=2) + "\n")
        for race, player in players.items():
            profiles[key, race] = {"Key": key, "Race": race, "BaselinePlayer": deepcopy(player)}
    if set(profiles) != expected_roster():
        raise ValueError("Selected profiles do not cover the current roster")
    args.output.parent.mkdir(parents=True, exist_ok=True)
    args.output.write_text(json.dumps({"Results": list(profiles.values())}, indent=2) + "\n")
    print(f"Prepared {len(profiles)} benchmark players")


if __name__ == "__main__":
    main()
