#!/usr/bin/env python3
"""Package complete tank-search JSON records without process logs or binaries."""

import argparse
import concurrent.futures
import gzip
import hashlib
import json
import os
from pathlib import Path


def archive(source, output):
    files = sorted(source.glob("*.json"))
    if not files or not (source / "manifest.json").is_file():
        raise ValueError(f"No complete research manifest in {source}")
    path = output / (source.name + ".json.gz")
    records = []
    with path.open("wb") as raw, gzip.GzipFile(fileobj=raw, mode="wb", mtime=0) as stream:
        stream.write(b'{"format":1,"files":{')
        for i, file in enumerate(files):
            data = file.read_bytes()
            json.loads(data)
            if i:
                stream.write(b",")
            stream.write(json.dumps(file.name).encode() + b":" + data.strip())
            records.append({"file": file.name, "sha256": hashlib.sha256(data).hexdigest(),
                            "bytes": len(data)})
        stream.write(b"}}\n")
    return {"archive": path.name, "sha256": hashlib.sha256(path.read_bytes()).hexdigest(),
            "bytes": path.stat().st_size, "records": records}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--source", type=Path, nargs="+", required=True)
    parser.add_argument("--output", type=Path, required=True)
    cpus = len(os.sched_getaffinity(0)) if hasattr(os, "sched_getaffinity") else (os.cpu_count() or 1)
    parser.add_argument("--workers", type=int, default=cpus)
    args = parser.parse_args()
    if len({path.name for path in args.source}) != len(args.source):
        raise SystemExit("Source directory names must be unique")
    args.output.mkdir(parents=True, exist_ok=True)
    with concurrent.futures.ThreadPoolExecutor(max_workers=args.workers) as executor:
        archives = list(executor.map(lambda path: archive(path, args.output), args.source))
    (args.output / "archives.json").write_text(json.dumps({
        "format": 1,
        "description": "Original JSON values keyed by source filename; text hashes describe the source records before JSON framing.",
        "archives": archives,
    }, indent=2) + "\n")
    print(f"{len(archives)} archives; {sum(a['bytes'] for a in archives)/1024**2:.1f} MiB")


if __name__ == "__main__":
    main()
