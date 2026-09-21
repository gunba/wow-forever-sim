#!/usr/bin/env python3
"""Download the existing game icons used in benchmark charts."""

import argparse
from concurrent.futures import ThreadPoolExecutor
import json
from pathlib import Path
from urllib.request import urlopen

from build_display import BUILDS


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--output", type=Path, default=Path("assets/img/spec_icons"))
    args = parser.parse_args()
    args.output.mkdir(parents=True, exist_ok=True)
    sources = {icon: f"https://wow.zamimg.com/images/wow/icons/large/{icon}.jpg" for _, _, _, icon in BUILDS}

    def fetch(item):
        icon, url = item
        path = args.output / f"{icon}.jpg"
        if not path.exists():
            with urlopen(url, timeout=30) as response:
                path.write_bytes(response.read())
        return icon

    for icon in ThreadPoolExecutor(4).map(fetch, sources.items()):
        print(icon)
    (args.output / "sources.json").write_text(json.dumps(sources, indent=2) + "\n")


if __name__ == "__main__":
    main()
