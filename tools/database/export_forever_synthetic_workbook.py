#!/usr/bin/env python3
"""Export the reviewed hypothetical equipment workbook from the pinned model catalog.

Optional dependency: XlsxWriter. Does not alter the simulator or real item records.
"""
from __future__ import annotations

import collections
import hashlib
import json
import xlsxwriter
from forever_synthetic_gear import *

def main():
    armor,weapons,budget,prices=evidence()
    assert all(z[-1]==0 for z in armor),[(z[0],z[-1]) for z in armor if z[-1]][:10]
    assert all(z[11]==z[14]==0 for z in weapons),[(z[0],z[11],z[14]) for z in weapons if z[11] or z[14]][:10]
    out=proposals_v2()
    assert len({r['ID'] for r in out})==len(out)
    assert all(r['Scale'] in (47,36,27,20,15) for r in out)
    path=ROOT/'artifacts/modelled_gear/proposed_gear.xlsx'
    w=xlsxwriter.Workbook(str(path))
    w.set_properties({'title':'WoW Forever proposed level-65 gear',
                      'subject':'Hypothetical equipment anchored to build '+BUILD,
                      'comments':'Synthetic scenarios are not obtainable items or finalized sim profiles.'})
    header=w.add_format({'bold':True,'bg_color':'#25364b','font_color':'#ffffff','text_wrap':True,'border':1,'valign':'vcenter'})
    text=w.add_format({'text_wrap':True,'valign':'top'})
    num=w.add_format({'num_format':'0.00','valign':'top'})
    mono=w.add_format({'font_name':'Consolas','font_size':10})
    caution=w.add_format({'bg_color':'#ffebc2','text_wrap':True})
    link=w.add_format({'font_color':'#1052a0','underline':1})
    percent=w.add_format({'num_format':'0.0%'})
    def sheet(name,columns,rows,widths=None,freeze=1):
        sh=w.add_worksheet(name);sh.freeze_panes(freeze,0)
        sh.set_row(0,36)
        for c,title in enumerate(columns):sh.write(0,c,title,header)
        sh.autofilter(0,0,len(rows),len(columns)-1)
        for i,row in enumerate(rows,1):
            for j,val in enumerate(row):
                if val is None:continue
                if isinstance(val,str) and val.startswith('https://'):
                    sh.write_url(i,j,val,link,val)
                elif isinstance(val,(int,float)):sh.write_number(i,j,val)
                else:sh.write(i,j,val)
        for j,width in enumerate(widths or []):sh.set_column(j,j,width)
        return sh
    readme=[
      ('Status','Current hypothetical gear-only benchmark. No SYN2-* item is confirmed obtainable. Earlier real and synthetic-model releases remain archived separately.'),
      ('Data','Forever beta build '+BUILD+'; pinned public catalog and client tables in assets/db_inputs/.'),
      ('Stat formula','round_half_up(RandPropPoints[level, quality, inventory slot] × StatPercentEditor / 10000). Proven for catalog client item records.'),
      ('Weapon formula','DPS = ItemDamageOneHand or TwoHand[level, quality] by verified weapon slot/type; min=floor(DPS × speed × (1 - DmgVariance/2)); max=round_half_up(DPS × speed × (1 + DmgVariance/2)).'),
      ('Armor formula','round_half_up(ItemArmorTotal[level, armor class] × ItemArmorQuality[level, quality] × ArmorLocation[slot, armor class]); robe uses chest location. Shield from ItemArmorShield.'),
      ('Budget evidence','Generic primary stats, Stamina, unified Hit and Crit approximately obey a 1.7095-power norm. This matches many existing item allocations; it is a supported working model, not proof of the server formula.'),
      ('Special costs','AP, spell power, school power and MP5 cannot be assigned a verified universal exchange price. Caster four-stat and caster/hybrid high-hit items mirror a projected physical allocation where Stamina was deliberately exchanged for Hit (the source item has no Hit) at the explicitly assumed armor spell-power price. Equal modeled cost does not imply equal DPS or Intellect-to-SP conversion.'),
      ('Price clues','The Price clues tab solves one special-stat multiplier per eligible no-effect/no-set item assuming its full budget is spent. If an item is underbudget, that result is an upper bound, not the actual game price.'),
      ('Stamina policy','For armor/neck/rings: reduce large source Stamina allocation to at most 20% of the modeled p-budget and transfer only that cost to a new primary/rating stat. No free offensive stats in the working model. Compare to actual source stats before adopting.'),
      ('Weapon policy','Stats remain exactly the real source allocation. Only speed changes; theoretical table DPS stays fixed within a weapon type. Cannot assert novel speeds will drop.'),
      ('Trinkets','No confirmed ilvl65 passive stat trinket. Only 80%-capacity projections enter rankings; the eight 100% variants are excluded sensitivity cases.'),
      ('Editing cells','Stat amounts recalculate when allocation basis points change. Source/proposed budget figures are generation-time snapshots: rerun generate.py to recompute them.'),
      ('Sources','All 65 crafted/PvP records are listed on Verified Sources; coverage gaps on Coverage. Source URLs are item-specific Wowhead Forever pages. Tables from wago.tools build '+BUILD+'.'),
      ('Constraints','Do not apply raid-tier set bonuses from synthetic pieces; retain role Tier1 override once. Respect class armor/weapon training, hand use, Shaman no dual wield, profession restrictions on REAL crafted items, unique rings/trinkets and racial hit.'),
      ('Hit comparison','The current model gets hit only from equipped items, enchants, talents and racials. It does not exchange offensive stats for hit. Excess hit is wasted; Tauren can meet the cap with fewer hit-bearing pieces.'),
      ('No claiming availability','Crafted/PvP references are real data, but projected slot/stat combinations, respeeded weapons and passive trinkets are what-if scenarios, not datamined item records. Full-catalog armor/weapon proof includes ineligible items only to test the client formula.'),
      ('Comparability','All ranked equipment is modeled and level 65, with slot budgets evaluated under the same explicit Lp hypothesis. Intellect does not grant baseline spell power. Cross-archetype DPS equivalence cannot be established by item budgets alone.'),
    ]
    sh=w.add_worksheet('Start here');sh.set_column(0,0,23);sh.set_column(1,1,121);sh.freeze_panes(1,0)
    sh.write_row(0,0,['Topic','Explanation'],header)
    for i,row in enumerate(readme,1):sh.write_row(i,0,row,text);sh.set_row(i,45)
    sh.autofilter(0,0,len(readme),1)
    # Input allocation rows are editable. The visible stat cells below contain Excel
    # formulas with cached numeric answers for viewers that do not recalculate files.
    alloc_sh=w.add_worksheet('Editable allocations')
    alloc_sh.freeze_panes(1,0);alloc_sh.set_row(0,32)
    acol=['Synthetic ID','Stat','Allocation (basis points)','Source ID','Source allocation (basis points)','Source policy']
    alloc_sh.write_row(0,0,acol,header)
    alloc_sh.set_column(0,0,18);alloc_sh.set_column(1,1,22);alloc_sh.set_column(2,4,23);alloc_sh.set_column(5,5,54)
    arows={};arow=1
    for row in out:
        source=normalized_alloc(BY_ID[row['SourceID']])
        for stat,raw in row['Allocation'].items():
            alloc_sh.write_row(arow,0,[row['ID'],stat,raw*10000,row['SourceID'],
                                         source.get(stat,0)*10000,row['Policy']])
            arows[(row['ID'],stat)]=arow+1;arow+=1
    alloc_sh.autofilter(0,0,arow-1,len(acol)-1)
    cols=['Synthetic ID','Slot','Inventory type','Armor / weapon type','Class archetype',
          'Modeled item name (with amounts)','Status / policy','Reference ID','Source item name','Source category',
          'Reference URL','Source slot','Stat capacity (q4/65)']+STAT_COLS+[
          'Base armor','Base Block','Speed (s)','Minimum damage','Maximum damage','Base weapon DPS',
          'Source modeled capacity used','Proposed modeled capacity used','Confidence','Limitations']
    sh=w.add_worksheet('Proposed gear');sh.freeze_panes(1,7);sh.set_row(0,47)
    for j,title in enumerate(cols):sh.write(0,j,title,header)
    sh.set_column(0,0,16);sh.set_column(1,1,16);sh.set_column(2,3,13);sh.set_column(4,4,22)
    sh.set_column(5,5,99);sh.set_column(6,6,34);sh.set_column(7,7,13);sh.set_column(8,8,41)
    sh.set_column(9,9,18);sh.set_column(10,10,37);sh.set_column(11,12,14)
    sh.set_column(13,13+len(STAT_COLS)-1,14)
    sh.set_column(13+len(STAT_COLS),len(cols)-2,16)
    sh.set_column(len(cols)-1,len(cols)-1,95)
    for i,r in enumerate(out,1):
        url=f'https://www.wowhead.com/forever/item={r["SourceID"]}'
        row=[r['ID'],r['Slot'],r['SlotID'],r['Material'],r['ClassUse'],modeled_name(r),
             r['Policy'],r['SourceID'],r['SourceName'],r['SourceKind'],url,r['SourceSlot'],r['Scale']]
        for j,v in enumerate(row):
            if j==10:sh.write_url(i,j,v,link,v)
            else:sh.write(i,j,v)
        for j,name in enumerate(STAT_COLS,13):
            index=arows.get((r['ID'],name))
            if index:
                sh.write_formula(i,j,f'=ROUND($M{i+1}*\'Editable allocations\'!$C${index}/10000,0)',None,
                                 r['Stats'].get(name,0))
            else:sh.write_number(i,j,0)
        tail=[r['Armor'],r.get('BaseBlockValue',0),r['Speed'],r['Min'],r['Max'],r['DPS'],
              r['SourceBudget'],r['ProposedBudget'],r['Confidence'],r['Notes']]
        for j,v in enumerate(tail,13+len(STAT_COLS)):
            if v is not None:sh.write(i,j,v)
    sh.autofilter(0,0,len(out),len(cols)-1)
    proof_cols=['Item ID','Verified item','Item level','Quality','Slot','Armor material',
                'Observed armor','Predicted armor','Error']
    sheet('Armor proof',proof_cols,armor,[12,46,12,10,18,16,15,16,10])
    wcols=['Item ID','Verified item','Item level','Quality','Slot','Speed',
           'DmgVariance','Client DPS table','Table DPS','Observed min','Calculated min',
           'Min error','Observed max','Calculated max','Max error']
    sheet('Weapon proof',wcols,weapons,[12,48,11,10,17,10,13,25,13,15,16,11,15,17,11])
    bcols=['Item ID','Verified item','Quality','Slot','Source type','Verified stats',
           'Has effect','Has set','Only generic stats?','Generic-cost norm',
           'Exploratory special-price norm','Source URL']
    brows=[list(row)+[f'https://www.wowhead.com/forever/item={row[0]}'] for row in budget]
    sheet('Budget evidence',bcols,brows,[12,48,10,18,20,78,13,13,21,22,32,49])
    sheet('Price clues',
          ['Item ID','Verified item','Slot','Only non-generic stat',
           'Raw special allocation / 10000','Generic p-budget spent','Implied price if full',
           'Verified quantities','Interpretation','Source URL'],
          prices,[12,49,18,24,29,24,22,66,91,49])
    vcols=['Item ID','Item','Item level','Quality','Slot','Material/subclass',
           'Source','Verified stats','Armor','Weapon min','Weapon max','Speed',
           'Client allocation?','Set','Effects','Requires profession','URL']
    verified=[]
    for r in ITEMS:
        if r['itemLevel']==65 and any(s['kind'] in ('crafted','pvp') for s in r['sources']):
            weapon=r.get('weapon',{})
            verified.append((r['id'],r['name'],65,r['quality'],SLOT.get(r['inventoryType'],'?'),
                ARMOR_TYPES.get(r['subclass'],'subclass '+str(r['subclass'])) if r['class']==4
                else 'Weapon subclass '+str(r['subclass']),source_kind(r),
                ', '.join(f'{k}={v}' for k,v in r['stats'].items()),r.get('armor'),
                weapon.get('min'),weapon.get('max'),weapon.get('speed'),
                bool(r.get('clientStatRecord')),r['setId'],bool(r['effects']),
                r['requiredSkill'],f'https://www.wowhead.com/forever/item={r["id"]}'))
    sheet('Verified Sources',vcols,verified,[12,48,12,10,18,20,17,89,12,14,14,11,17,12,11,19,49])
    cv=[]
    for slot in [1,2,3,16,5,9,10,6,7,8,11,12,13,21,22,17,14,23,15,26,28]:
        real=[r for r in ITEMS if r['itemLevel']==65 and
              (r['inventoryType']==slot or slot==5 and r['inventoryType']==20) and
              any(s['kind'] in ('crafted','pvp') for s in r['sources'])]
        sub=[r for r in out if r['SlotID']==slot]
        quality=collections.Counter(r['quality'] for r in real)
        cv.append((SLOT[slot],slot,quality[4],quality[3],sum(r.get('clientStatRecord',False) for r in real),
                   sum(s['kind']=='crafted' for r in real for s in r['sources']),
                   sum(s['kind']=='pvp' for r in real for s in r['sources']),
                   len(sub),scale(slot),
                   'No real ilvl65 crafted/PvP source; entirely projected' if not real else
                   'Real item(s), but new stat/slot combinations are projected'))
    sheet('Coverage', ['Slot','Inventory ID','Real q4 ilvl65','Real q3 ilvl65',
                       'Real client stat records','Crafted source flags','PvP source flags',
                       'What-if variants','Q4 stat capacity','Coverage caveat'],
          cv,[22,17,20,20,27,23,21,21,23,83])
    model_rows=[('Client build',BUILD,'Pinned reference/ files'),
                ('Quality 4 slot budget group 0',47,'Head/chest/legs/two-hand'),
                ('Quality 4 slot budget group 1',36,'Shoulders/waist/feet/hands/trinket'),
                ('Quality 4 slot budget group 2',27,'Neck/wrist/ring/shield/cloak/held off-hand'),
                ('Quality 4 slot budget group 3',20,'One-hand weapons'),
                ('Quality 4 slot budget group 4',15,'Bows/guns/crossbows/wands/relics'),
                ('Lp exponent',P,'Classic-style working hypothesis, not proven Forever price curve'),
                ('Generic stat unit cost',1,'Empirically supported by clean L65 items'),
                ('AP model price, armor',.66,'One clean L65 cloak pair; other slots differ'),
                ('AP model price, weapon',.51,'Different cost implied by existing PvP blades'),
                ('AP model price, ranged',.41,'Different cost implied by PvP ranged weapons'),
                ('SP model price, armor',.90,'Two clean L65 PvP leg examples imply an upper price ~0.927; not proof of a universal price'),
                ('SP model price, offhand',.712,'One L65 PvP tome, not generalizable to armor'),
                ('SP model price, two-hand',.78,'One L65 PvP two-hander; another crafted axe implies a higher upper price'),
                ('MP5 model price, armor',3.,'Conservative armor-side projection; clean L65 PvP mail legs upper ~3.066'),
                ('MP5 model price, one-hand',2.3,'Clean L65 PvP battle mace upper ~2.268 plus rounding'),
                ('School power model price',.885,'Three matching new crafted templates; limited validation'),
                ('MP5 model price',2.3,'One clean q4/65 weapon implies ~2.27; others differ'),
                ('Passive trinket lower scenario',.80,'No verified ilvl65 passive analog'),
                ('Passive trinket full scenario',1.00,'No verified ilvl65 passive analog'),
                ('Armor formula exact matches',sum(v[-1]==0 for v in armor),f'{len(armor)} tested catalog records'),
                ('Weapon formula exact matches',sum(v[11]==v[14]==0 for v in weapons),f'{len(weapons)} tested catalog records'),
                ('Source code release','https://github.com/gunba/wow-forever-sim','Reference catalog at assets/db_inputs/forever_gear_catalog.json'),
                ('Client table reference',f'https://wago.tools/db2/RandPropPoints/csv?build={BUILD}','Other table names in reference/'),
    ]
    sheet('Model', ['Parameter','Value','Provenance / limit'],model_rows,[45,72,115])
    selected_results = ROOT/'artifacts/modelled_gear/forever_dps_5min.json'
    if selected_results.exists():
        selected = json.loads(selected_results.read_text())
        if selected.get('GearScenario') == 'modeled-65-v2':
            import sys
            sys.path.insert(0, str(ROOT/'tools/forever_bench'))
            from gear_equity import inspect
            profiles, _ = inspect(selected, json.loads(
                (ROOT/'assets/db_inputs/forever_synthetic_gear_v2.json').read_text()))
            columns = ['Build', 'Race', 'DPS', 'ItemCount', 'MeanItemLevel',
                       'EstimatedBudgetUnits', 'AssumedCapacityPoints',
                       'MeleeHitPercent', 'SpellHitPercent',
                       'WorstHitShortfallPercent', 'Strength', 'Agility', 'Intellect',
                       'AttackPower', 'SpellPower', 'CritRating', 'HitRating', 'MP5']
            sheet('Selected profiles', columns,
                  [[profile[column] for column in columns] for profile in profiles],
                  [23, 18, 13, 15, 18, 20, 22, 16, 16, 25] + [15]*8)
    w.close()
    print(f'{path}\n{len(out)} proposed variants; {len(verified)} ilvl65 crafted/PvP source records; '
          f'{len(armor)} exact armor rows, {len(weapons)} exact weapon rows')
    for fn in sorted(REF.iterdir()):
        print(hashlib.sha256(fn.read_bytes()).hexdigest(),fn.relative_to(ROOT))

if __name__=='__main__':main()
