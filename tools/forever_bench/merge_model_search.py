#!/usr/bin/env python3
"""Merge converged follow-up swaps into a complete initial modeled search.

The final profile for each continued pair is replayed with the original seed
before independent paired confirmation, keeping the first comparison matched.
"""
from __future__ import annotations

import argparse
import json
from pathlib import Path
import re
import shutil
import subprocess


def main():
    parser=argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--initial',type=Path,required=True)
    parser.add_argument('--followup',type=Path,required=True)
    parser.add_argument('--binary',type=Path,required=True)
    parser.add_argument('--output',type=Path,required=True)
    args=parser.parse_args()
    args.output.mkdir(parents=True,exist_ok=True)
    original=json.loads((args.initial/'results.json').read_text())
    baseline=json.loads((args.initial/'baseline.json').read_text())
    summary=json.loads((args.initial/'summary.json').read_text())
    manifest=json.loads((args.initial/'manifest.json').read_text())
    pairs={tuple(pair) for pair in summary['unconverged']}
    if not pairs or summary['completed']!=summary['requested'] or summary['failures']:
        raise ValueError('Expected a complete first pass with unconverged profiles')
    reports={}
    for path in sorted(args.followup.glob('*.gear-search.json')):
        row=json.loads(path.read_text())
        pair=(row['Build'],row['Race'])
        if pair not in pairs or pair in reports or not row['Converged']:
            raise ValueError(f'Follow-up result missing, duplicated or unconverged: {pair}')
        reports[pair]=(path,row)
    if reports.keys()!=pairs:
        raise ValueError(f'Missing follow-up searches: {pairs-reports.keys()}')
    source_by_pair={(r['Key'],r['Race']):i for i,r in enumerate(original['Results'])}
    changed=0
    for pair,(path,report) in reports.items():
        if pair not in source_by_pair:
            raise ValueError(f'Unknown pair: {pair}')
        if report['Final']['BaselinePlayer']==original['Results'][source_by_pair[pair]]['BaselinePlayer']:
            continue
        changed+=1
        slug=pair[0]+'__'+re.sub(r'[^a-z0-9]+','_',pair[1].lower()).strip('_')
        selected={**original,'Results':[report['Final']]}
        selected_path=args.output/(slug+'.selected.json')
        selected_path.write_text(json.dumps(selected,indent=2)+'\n')
        prefix=args.output/(slug+'.matched')
        subprocess.run([str(args.binary),'-build',pair[0],'-race',pair[1],
                        '-baseline-results',str(selected_path),'-iterations',
                        str(manifest['settings']['iterations']),'-seed',
                        str(manifest['settings']['seed']),'-output',str(prefix)],check=True)
        replay=json.loads(Path(str(prefix)+'.json').read_text())['Results'][0]
        if replay['BaselinePlayer']!=report['Final']['BaselinePlayer'] or replay['Warnings']:
            raise ValueError(f'Matched replay changed the selected profile: {pair}')
        original['Results'][source_by_pair[pair]]=replay
        slot=next(i for i,r in enumerate(baseline['Results']) if (r['Key'],r['Race'])==pair)
        summary['reports'][slot]=str(path.resolve())
    original['GearSearchFollowup']=[{'Key':pair[0],'Race':pair[1], 'Report':path.name}
                                    for pair,(path,_) in sorted(reports.items())]
    summary['initialUnconverged']=summary.pop('unconverged')
    summary['unconverged']=[]
    summary['followupReports']=[str(path) for path,_ in reports.values()]
    for filename,obj in [('results.json',original),('summary.json',summary)]:
        (args.output/filename).write_text(json.dumps(obj,indent=2)+'\n')
    for filename in ('baseline.json','manifest.json'):
        shutil.copyfile(args.initial/filename,args.output/filename)
    print(f'Confirmed {len(reports)} converged follow-ups; {changed} new profiles needed a matched first-seed replay')


if __name__=='__main__':main()
