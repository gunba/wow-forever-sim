#!/usr/bin/env python3
"""Archive the complete modeled gear-search and confirmation evidence compactly."""
from __future__ import annotations

import argparse
import gzip
import json
from pathlib import Path
import shutil


def packed_json(path: Path, pieces: list[tuple[str, Path]]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open('wb') as raw:
        with gzip.GzipFile(filename='', mode='wb', fileobj=raw, mtime=0, compresslevel=9) as out:
            out.write(b'{')
            for index, (name, source) in enumerate(pieces):
                if index:
                    out.write(b',')
                out.write(json.dumps(name).encode()+b':')
                with source.open('rb') as stream:
                    shutil.copyfileobj(stream, out)
            out.write(b'}\n')


def main():
    parser=argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--search', required=True, type=Path)
    parser.add_argument('--confirmation', required=True, type=Path)
    parser.add_argument('--initial-search', type=Path,
                        help='first complete pass if some pairs required a follow-up')
    parser.add_argument('--output', required=True, type=Path)
    args=parser.parse_args()
    summary=json.loads((args.search/'summary.json').read_text())
    confirmation=json.loads((args.confirmation/'selections.json').read_text())
    if summary['completed']!=summary['requested'] or summary['failures']:
        raise ValueError('Cannot archive incomplete gear selection')
    if len(confirmation)!=summary['completed']:
        raise ValueError('Confirmation decisions do not cover every race/build profile')
    args.output.mkdir(parents=True, exist_ok=True)
    reports=[Path(path) for path in summary['reports']]
    if any(not path.exists() for path in reports):
        raise ValueError('A selected report is missing')
    summary['reports']=[path.name for path in reports]
    summary['followupReports']=[Path(path).name for path in summary.get('followupReports',[])]
    summary_path=args.output/'summary.json'
    summary_path.write_text(json.dumps({
        'scenario':'modeled-65-v1',
        'searchManifest':json.loads((args.search/'manifest.json').read_text()),
        'confirmationManifest':json.loads((args.confirmation/'manifest.json').read_text()),
        'searchSummary':summary,
        'decisions':confirmation,
    },indent=2)+'\n')
    # Stream the large report array rather than materializing all 184 trials.
    with (args.output/'search.json.gz').open('wb') as raw:
        with gzip.GzipFile(filename='',mode='wb',fileobj=raw,mtime=0,compresslevel=9) as out:
            out.write(b'{"manifest":')
            out.write((args.search/'manifest.json').read_bytes().strip())
            out.write(b',"reports":[')
            for index,path in enumerate(reports):
                if index:out.write(b',')
                with path.open('rb') as stream:shutil.copyfileobj(stream,out)
            out.write(b']}\n')
    if summary.get('followupReports'):
        followups=json.loads((args.search/'summary.json').read_text())['followupReports']
        with (args.output/'convergence_followups.json.gz').open('wb') as raw:
            with gzip.GzipFile(filename='',mode='wb',fileobj=raw,mtime=0,compresslevel=9) as out:
                out.write(b'{"reports":[')
                for index,source in enumerate(followups):
                    if index:out.write(b',')
                    with Path(source).open('rb') as stream:shutil.copyfileobj(stream,out)
                out.write(b']}\n')
    if args.initial_search:
        first=json.loads((args.initial_search/'summary.json').read_text())
        if first['completed']!=len(reports):
            raise ValueError('Initial search does not cover the final roster')
        with (args.output/'initial_search.json.gz').open('wb') as raw:
            with gzip.GzipFile(filename='',mode='wb',fileobj=raw,mtime=0,compresslevel=9) as out:
                out.write(b'{"manifest":')
                out.write((args.initial_search/'manifest.json').read_bytes().strip())
                out.write(b',"reports":[')
                for index,source in enumerate(first['reports']):
                    if index:out.write(b',')
                    with Path(source).open('rb') as stream:shutil.copyfileobj(stream,out)
                out.write(b']}\n')
    packed_json(args.output/'validation.json.gz',[
        ('manifest',args.confirmation/'manifest.json'),
        ('decisions',args.confirmation/'selections.json'),
        ('pairs',args.confirmation/'pairs.json'),
    ])
    packed_json(args.output/'before.json.gz',[
        ('searchBaseline',args.search/'baseline.json'),
        ('searchFinal',args.search/'results.json'),
        ('selected',args.confirmation/'selected.json'),
    ])
    print(f"Archived {len(reports)} complete searches and {len(confirmation)} decisions to {args.output}")


if __name__=='__main__':main()
