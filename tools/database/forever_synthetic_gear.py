#!/usr/bin/env python3
"""Rebuild the level-65 *hypothetical* Forever equipment workbook.

Reads pinned client tables and the verified catalog to generate labeled synthetic proposals.
This module does not modify the real-item catalog or existing profiles.
"""
from __future__ import annotations

import collections
import csv
import hashlib
import json
import math
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
REF = ROOT / 'assets/db_inputs/forever_model_tables'
CATALOG = json.loads((ROOT / 'assets/db_inputs/forever_gear_catalog.json').read_text())
ITEMS = CATALOG['items']
BY_ID = {r['id']: r for r in ITEMS}
BUILD = CATALOG['clientBuild']
P = math.log(2) / math.log(1.5)  # Empirical Classic-style Lp budget hypothesis; not a verified Forever price curve.
MOD = {
    3:'Agility', 4:'Strength', 5:'Intellect', 6:'Spirit', 7:'Stamina',
    31:'HitRating', 32:'CritRating', 38:'AttackPower',39:'RangedAttackPower',
    41:'HealingPower',42:'SpellDamage',43:'MP5',45:'SpellPower',50:'BonusArmor',
    84:'HolyPower',85:'FirePower',86:'NaturePower',87:'FrostPower',88:'ShadowPower',89:'ArcanePower',
}
STAT_COLS = [
    'Stamina','Strength','Agility','Intellect','Spirit','HitRating','CritRating',
    'AttackPower','RangedAttackPower','SpellPower','SpellDamage','MP5',
    'FirePower','FrostPower','ArcanePower','ShadowPower','NaturePower','HolyPower',
    'HealingPower','BonusArmor',
]
GENERIC = {'Stamina','Strength','Agility','Intellect','Spirit','HitRating','CritRating'}
# Upper-frontier item-implied exchange prices, not server/Blizzard item prices.
# The explicitly hypothetical v2 caster/hybrid mirrors use these assumed
# prices; equality under this model is not evidence of equal game DPS budgets.
BUDGET_PRICE = {'AttackPower':.66,'RangedAttackPower':.48,'SpellPower':.9,
                'SpellDamage':1.,'MP5':3.,'BonusArmor':.07,'HealingPower':.52,
                **{s:.885 for s in ('HolyPower','FirePower','FrostPower','ArcanePower','ShadowPower','NaturePower')}}
SLOT = {1:'Head',2:'Neck',3:'Shoulders',5:'Chest',6:'Waist',7:'Legs',8:'Feet',9:'Wrists',
        10:'Hands',11:'Finger',12:'Trinket',13:'One-hand',14:'Shield',15:'Ranged bow',
        16:'Back',17:'Two-hand',20:'Chest',21:'Main hand',22:'Off hand weapon',23:'Held off-hand',
        25:'Thrown',26:'Ranged/wand',28:'Relic'}
GROUP = {**dict.fromkeys((1,5,7,17,20),0),**dict.fromkeys((3,6,8,10,12),1),
         **dict.fromkeys((2,9,11,14,16,23),2),**dict.fromkeys((13,21,22),3),
         **dict.fromkeys((15,25,26,28),4)}
ARMOR_TYPES = {1:'Cloth',2:'Leather',3:'Mail',4:'Plate'}
BODY_SLOTS = [1,3,5,9,10,6,7,8]
ARMOR_LOC = {1:'Clothmodifier',2:'Leathermodifier',3:'Chainmodifier',4:'Platemodifier'}
TABLES = {n:{int(row['ID']):row for row in csv.DictReader((REF / (n + '.csv')).open())} for n in
          ('RandPropPoints','ItemArmorTotal','ItemArmorQuality','ItemArmorShield','ArmorLocation',
           'ItemDamageOneHand','ItemDamageTwoHand','ItemDamageRanged')}
VARIANCE = {int(row['ID']):float(row['DmgVariance']) for row in csv.DictReader((REF/'WeaponSparseVariance.csv').open())}


def round_nearest(x):
    return math.floor(x+.5)


def scale(slot, level=65, quality=4):
    prefix={2:'Good',3:'Superior',4:'Epic'}[quality]
    return float(TABLES['RandPropPoints'][level][f'{prefix}F_{GROUP[slot]}'])


def body_armor(slot, material, level=65, quality=4):
    if slot==14:
        return int(TABLES['ItemArmorShield'][level][f'Quality_{quality}'])
    if slot==16:material=1
    base=float(TABLES['ItemArmorTotal'][level][ARMOR_TYPES[material]])
    multiplier=float(TABLES['ItemArmorQuality'][level][f'Qualitymod_{quality}'])
    pos=float(TABLES['ArmorLocation'][5 if slot==20 else slot][ARMOR_LOC[material]])
    return round_nearest(base*multiplier*pos)


def damage_table(item):
    # Verified against all catalog weapon damage ranges, not inferred from names.
    slot=item['inventoryType'];kind=item['subclass']
    return 'ItemDamageTwoHand' if slot in (17,15,25) or slot==26 and kind in (2,3,18) else 'ItemDamageOneHand'


def damage(item, speed=None):
    speed = item['weapon']['speed'] if speed is None else speed
    dps = float(TABLES[damage_table(item)][item['itemLevel']][f'Quality_{item["quality"]}'])
    spread=VARIANCE[item['id']]
    return math.floor(dps*speed*(1-spread/2)),round_nearest(dps*speed*(1+spread/2)),dps


def normalized_alloc(item):
    return {MOD[a['mod']]:a['allocation']/10000 for a in item.get('statAllocations',[]) if a['mod'] in MOD}


def norm(alloc,slot=None):
    def price(stat):
        if stat=='AttackPower':
            return .41 if slot in (15,25,26,28) else .51 if slot in (13,21,22) else .66
        if stat=='SpellPower' and slot==23:return .712
        if stat=='SpellPower' and slot==17:return .78
        if stat=='MP5' and slot in (13,21,22):return 2.3
        return 1 if stat in GENERIC else BUDGET_PRICE.get(stat,1)
    return sum((v*price(k))**P for k,v in alloc.items())**(1/P)


def item_stats(alloc,slot):
    points=scale(slot)
    return {k:round_nearest(points*v) for k,v in alloc.items() if round_nearest(points*v)}


def reduce_stamina(alloc,add,ratio=.20,slot=None):
    """Trade ONLY Stamina against a different generic stat; preserve hypothesized cost exactly.

    Special stat allocations stay unchanged. If the selected template has too little
    Stamina, do not invent additional Stamina. No merging with an existing stat.
    """
    alloc=alloc.copy()
    if 'Stamina' not in alloc or add in alloc:
        raise ValueError(f'need Stamina and a fresh stat slot: {alloc}, {add}')
    total=norm(alloc,slot)**P
    old=alloc['Stamina']**P
    new=min(old,ratio*total)
    alloc['Stamina']=new**(1/P)
    delta=max(0,old-new)
    if delta:alloc[add]=delta**(1/P)
    assert abs(norm(alloc,slot)**P-total)<1e-8
    return alloc


def conservatively_project(alloc,source_slot,target_slot):
    """Do not make a template larger when the target slot has a higher assumed cost."""
    old=norm(alloc,source_slot)
    new=norm(alloc,target_slot)
    factor=min(1,old/new) if new else 1
    return {k:v*factor for k,v in alloc.items()}


def source_kind(item):
    return ','.join(sorted({s['kind'] for s in item['sources']}))


def synthetic(identifier,slot,material,archetype,template,alloc,policy,confidence,note='',speed=None):
    src=BY_ID[template]
    assert src['itemLevel']==65 and src['quality']==4 and src['clientStatRecord']
    v=item_stats(alloc,slot)
    weapon=src.get('weapon') if speed is not None else None
    dmin,dmax,dps=damage(src,speed) if weapon else (None,None,None)
    armor=body_armor(slot,material) if (material in ARMOR_TYPES and slot in BODY_SLOTS+[16]) or slot==14 else None
    item_type=ARMOR_TYPES.get(material,'--') if slot in BODY_SLOTS+[16] else \
        'Shield' if slot==14 else 'Weapon subclass '+str(material) if speed is not None else '--'
    class_use=({'Cloth':'Mage/Priest/Warlock','Leather':'Druid/Rogue','Mail':'Hunter/Shaman',
                'Plate':'Paladin/Warrior'}.get(item_type,'Equipment restrictions apply')
               if slot in BODY_SLOTS+[16,2,11,12] else
               'Weapon training and hand restrictions apply; Shaman cannot dual wield')
    return dict(ID=identifier,Slot=SLOT[slot],SlotID=slot,Material=item_type,
                ClassUse=class_use,
                Archetype=archetype,Policy=policy,SourceID=src['id'],SourceName=src['name'],
                SourceKind=source_kind(src),SourceSlot=SLOT[src['inventoryType']],
                Scale=scale(slot),Allocation=alloc,Stats=v,Armor=armor,
                BaseBlockValue=src.get('baseBlockValue', 0) if slot==14 else 0,
                Speed=speed,Min=dmin,Max=dmax,DPS=dps,
                SourceBudget=norm(normalized_alloc(src),src['inventoryType']),
                ProposedBudget=norm(alloc,slot),
                Confidence=confidence,Notes=note)


def evidence():
    armor=[];weapons=[];budget=[];prices=[]
    for r in ITEMS:
        if r.get('clientStatRecord') and r.get('armor') and r['class']==4:
            inv=r['inventoryType'];material=r['subclass']
            if inv==14 and material==6:actual=body_armor(14,6,r['itemLevel'],r['quality'])
            elif inv in BODY_SLOTS+[16,20] and material in ARMOR_TYPES:
                actual=body_armor(inv,material,r['itemLevel'],r['quality'])
            else:continue
            armor.append((r['id'],r['name'],r['itemLevel'],r['quality'],SLOT[inv],
                          ARMOR_TYPES.get(material,'Shield'),r['armor'],actual,actual-r['armor']))
        if r.get('clientStatRecord') and r.get('weapon') and r['id'] in VARIANCE:
            low,high,dps=damage(r)
            weapons.append((r['id'],r['name'],r['itemLevel'],r['quality'],SLOT[r['inventoryType']],
                            r['weapon']['speed'],VARIANCE[r['id']],damage_table(r),round(dps,5),
                            r['weapon']['min'],low,low-r['weapon']['min'],
                            r['weapon']['max'],high,high-r['weapon']['max']))
        if (r.get('clientStatRecord') and r.get('statAllocations') and r['itemLevel']==65
                and any(s['kind'] in ('crafted','pvp') for s in r['sources'])):
            a=normalized_alloc(r)
            pnorm=sum(v**P for k,v in a.items() if k in GENERIC)**(1/P)
            budget.append((r['id'],r['name'],r['quality'],SLOT.get(r['inventoryType'],'?'),
                           source_kind(r),', '.join(f'{k} {v}' for k,v in r['stats'].items()),
                           bool(r['effects']),bool(r['setId']),
                           all(k in GENERIC for k in a),pnorm,norm(a,r['inventoryType'])))
            special=[(k,v) for k,v in a.items() if k not in GENERIC]
            if r['quality']==4 and len(special)==1 and not r['effects'] and not r['setId']:
                stat,value=special[0]
                generic_spent=sum(v**P for k,v in a.items() if k in GENERIC)
                if generic_spent < 1:
                    implied=(1-generic_spent)**(1/P)/value
                    prices.append((r['id'],r['name'],SLOT.get(r['inventoryType'],'?'),
                        stat,round(value,5),round(generic_spent,5),round(implied,5),
                        ', '.join(f'{k}={v}' for k,v in r['stats'].items()),
                        'Equality requires a fully spent budget; otherwise this is only an upper price.',
                        f'https://www.wowhead.com/forever/item={r["id"]}'))
    return armor,weapons,budget,prices


def proposals():
    result=[]
    families=[
        ('Strength / crit / hit',279253,{'Agility':'Strength'},'HitRating','high','verified generic price group; armor changes are projected'),
        ('Agility / crit / hit',279253,{},'HitRating','high','verified generic price group; armor changes are projected'),
        ('Strength / agility / crit / hit',22651,{},'Agility','high','verified generic price group; armor changes are projected'),
        ('Agility / AP / crit',20068,{},'CritRating','medium','AP allocation retained; source is a cloak'),
        ('Intellect / spell power / crit',272685,{},'CritRating','conditional','spell-power template is an off-hand; cross-slot cost unverified'),
        ('Intellect / MP5 / crit',272681,{},'CritRating','conditional','MP5 cost not independently determined'),
        ('Intellect / fire power / crit',279266,{},'CritRating','medium','school power held fixed; cross-slot projection'),
        ('Crit / shadow power / intellect',279269,{},'Intellect','medium','school power held fixed; cross-slot projection'),
        ('Strength / intellect / spell power / hit',272682,{},'HitRating','conditional','hybrid template is a two-hander; cross-slot cost unverified'),
    ]
    seq=1
    for slot in BODY_SLOTS+[16,2,11]:
        materials=[1,2,3,4] if slot in BODY_SLOTS else [1] if slot==16 else [0]
        for material in materials:
            for name,id,swap,add,confidence,note in families:
                src=BY_ID[id]
                a=normalized_alloc(src)
                for start,end in swap.items():a[end]=a.pop(start)
                a=conservatively_project(a,src['inventoryType'],slot)
                a=reduce_stamina(a,add,slot=slot)
                # Material is illustrative, not a source profession or a fabricated obtainable recipe.
                status='armor type and slot projected' if slot in BODY_SLOTS else 'slot projected, no actual crafted/PvP neck/ring'
                result.append(synthetic(f'SYN-{seq:04}',slot,material,name,id,a,'20% Stamina cost (maximum)',
                                        confidence,note+'; '+status))
                seq+=1
    # Same base cap as shoulders is client-observed for trinket inventoryType=12,
    # but there are no suitable real level-65 passive stat trinkets to validate availability.
    for fraction in (.80,1.00):
        for name,id,swap,add,_,note in families:
            if name=='Intellect / MP5 / crit':continue
            a=normalized_alloc(BY_ID[id])
            for start,end in swap.items():a[end]=a.pop(start)
            a=conservatively_project(a,BY_ID[id]['inventoryType'],12)
            a=reduce_stamina(a,add,slot=12)
            a={k:v*fraction for k,v in a.items()}
            result.append(synthetic(f'SYN-{seq:04}',12,0,'Passive '+name,id,a,
                                    f'{fraction:.0%} of template allocation, 20% Stamina share',
                                    'unverified', 'No confirmed ilvl65 passive trinket; sensitivity scenario, not a verified item. '+note))
            seq+=1
    # Weapons: base DPS and damage variance are exact table predictions. Distinct
    # speeds are alternate synthetic items, NEVER extra DPS for a fast weapon.
    weapon_families=[
        ('Sword one-hand',272452,[1.8,2.4,2.9]),
        ('Axe one-hand',272592,[1.8,2.4,2.9]),
        ('Mace one-hand',279261,[1.8,2.3,2.9]),
        ('Dagger main-hand',272683,[1.7,2.0,2.5]),
        ('Dagger off-hand (unverified)',272683,[1.7,2.0]),
        ('Fist main-hand',272597,[1.8,2.4,2.9]),
        ('Fist off-hand',272598,[1.7,2.0,2.4]),
        ('Axe two-hand',272593,[3.0,3.4,3.8]),
        ('Mace two-hand',272601,[3.0,3.4,3.8]),
        ('Sword two-hand',272604,[3.0,3.4,3.8]),
        ('Polearm two-hand',272602,[3.0,3.4,3.8]),
        ('Staff two-hand',272603,[2.8,3.4,3.8]),
        ('Spell two-hand',272682,[2.8,3.4,3.8]),
        ('Gun ranged',279273,[1.8,2.5,3.0]),
        ('Crossbow ranged',272595,[2.0,2.5,3.0]),
        ('Bow ranged',272594,[1.8,2.5,3.0]),
        ('Wand ranged',279246,[1.5,1.7,2.0]),
    ]
    for label,id,speeds in weapon_families:
        src=BY_ID[id]
        for speed in speeds:
            a=normalized_alloc(src)
            offhand=label.startswith('Dagger off-hand')
            slot=22 if offhand else src['inventoryType']
            result.append(synthetic(f'SYN-{seq:04}',slot,src['subclass'],label,id,a,
                                    'Source stat allocations unchanged','low' if offhand else 'medium',
                                    'Client DPS + variance are exact for known weapons; this speed/item is hypothetical. '
                                    'An ilvl65 off-hand dagger has NOT been verified. ' if offhand else
                                    'Client DPS + variance are exact for known weapons; this speed/item is hypothetical. '
                                    'Check hand/class constraints and do not dual-wield Shaman.',speed=speed))
            seq+=1
    # There is no verified level-65 crafted/PvP spell-power one-hand weapon:
    # project a real spell-power offhand's allocation to the one-hand budget,
    # while retaining an observed weapon type/variance for damage.
    caster_weapon=BY_ID[272683]
    a=conservatively_project(normalized_alloc(BY_ID[272685]),23,13)
    result.append(synthetic(f'SYN-{seq:04}',13,caster_weapon['subclass'],
                            'Caster dagger one-hand (projected)',272683,a,
                            'Spell-power offhand pattern projected onto weapon','unverified',
                            'Stat pattern comes from real offhand item 272685, not this dagger. '
                            'Weapon stat price and coexistence of this DPS with spell power are unverified.',
                            speed=2.0))
    result[-1]['ExtraSourceID']=272685
    seq+=1
    for label,id in [('Physical shield',272591),('Hybrid damage shield',278469),
                     ('Spell-power held off-hand',272685)]:
        r=BY_ID[id];a=normalized_alloc(r)
        # No conversion of healing-only stats to damage.
        result.append(synthetic(f'SYN-{seq:04}',r['inventoryType'],r['subclass'],label,id,a,
                                'Original stat allocation, not a real new item','medium',
                                'Source template copied. HealingPower remains healing only; shield/relic effects omitted.'))
        seq+=1
    return result


def full_budget(alloc, slot):
    """Scale a projected allocation to one hypothesized slot budget."""
    capacity = norm(alloc, slot)
    if capacity <= 0:
        raise ValueError(f'empty modeled allocation for slot {slot}')
    return {stat: amount / capacity for stat, amount in alloc.items()}


def proposals_v2():
    """Equal-capacity synthetic-only pool; retain v1 separately for replay."""
    result = proposals()
    for index, row in enumerate(result):
        if row['SlotID'] == 12:
            continue  # Passive-trinket fractions remain distinct sensitivity assumptions.
        if row['SlotID'] in BODY_SLOTS and 'Strength' in row['Allocation']:
            material = next(k for k, v in ARMOR_TYPES.items() if v == row['Material'])
            if material != 4:
                # A plate allocation is not evidence for Strength on every
                # leather/mail/cloth slot. Retain the model ID, but substitute
                # leather/mail armor. Cloth has no usable client-verified
                # offensive template here, so use a caster off-hand allocation.
                reference = {1: 272685, 2: 279253, 3: 22676}[material]
                src = BY_ID[reference]
                if material == 1:
                    alloc = conservatively_project(normalized_alloc(src), src['inventoryType'], row['SlotID'])
                    alloc = reduce_stamina(alloc, 'HitRating', slot=row['SlotID'])
                    archetype = 'Intellect / spell power / hit (cloth)'
                elif material == 2:
                    alloc = conservatively_project(normalized_alloc(src), src['inventoryType'], row['SlotID'])
                    alloc = reduce_stamina(alloc, 'HitRating', slot=row['SlotID'])
                    archetype = 'Agility / crit / hit (leather)'
                else:
                    alloc = conservatively_project(normalized_alloc(src), src['inventoryType'], row['SlotID'])
                    if row['Archetype'] == 'Strength / crit / hit':
                        # One existing mail variant spends the source's MP5
                        # allocation on Hit instead. This is an explicitly
                        # hypothetical swap at the estimated 3:1 MP5 price,
                        # not a claim that the source item carries Hit.
                        alloc['HitRating'] = alloc.pop('MP5') * 3
                        archetype = 'Strength / intellect / crit / hit (mail)'
                    else:
                        archetype = 'Strength / intellect / crit / MP5 (mail)'
                result[index] = synthetic(
                    row['ID'], row['SlotID'], material, archetype, reference, alloc,
                    '20% Stamina cost (maximum); non-plate projection',
                    'unverified', 'Leather and mail use current same-material '
                    'items. The mail Hit variant replaces the source MP5 at an '
                    'assumed cost; the source does not have Hit. Cloth uses a '
                    'caster off-hand allocation, not verified cloth armor. '
                    'All slots and budgets remain hypothetical.',
                )
                row = result[index]
        row['Allocation'] = full_budget(row['Allocation'], row['SlotID'])
        row['Stats'] = item_stats(row['Allocation'], row['SlotID'])
        row['ProposedBudget'] = norm(row['Allocation'], row['SlotID'])
    seq = len(result) + 1

    # The genuine four-generic-stat leg allocation is mirrored across offensive
    # archetypes. One Strength unit becomes one assumed-cost unit of SpellPower,
    # and Agility becomes Intellect. This is equal *item cost*, not equal DPS:
    # Intellect does not grant baseline spell power in Forever.
    for slot in BODY_SLOTS + [16, 2, 11, 12]:
        for material in ([1, 2, 3, 4] if slot in BODY_SLOTS else [1] if slot == 16 else [0]):
            physical = full_budget(normalized_alloc(BY_ID[22651]), slot)
            physical = reduce_stamina(physical, 'Agility', slot=slot)
            mirrored = physical.copy()
            mirrored['SpellPower'] = mirrored.pop('Strength') / (
                .712 if slot == 23 else .78 if slot == 17 else .9)
            mirrored['Intellect'] = mirrored.pop('Agility')
            mirrored = full_budget(mirrored, slot)
            if slot == 12:
                mirrored = {k: value * .8 for k, value in mirrored.items()}
            item = synthetic(f'SYN2-{seq:04}', slot, material,
                             'Intellect / spell power / crit / hit (matched cost)',
                             22651, mirrored,
                             '80% passive trinket capacity' if slot == 12 else
                             '20% Stamina; mirrored physical/caster modeled cost',
                             'unverified',
                             'Four-stat caster mirror of a real physical allocation. '
                             'Spell-power price 0.9 is a benchmark hypothesis; '
                             'Intellect is not converted to spell power.')
            result.append(item)
            seq += 1

    # The observed three-offensive-stat physical allocation carries much
    # more hit than the four-stat variant. Mirror that allocation as a caster
    # and as a hybrid so the gear-only benchmark can reach spell hit without
    # fabricated bonuses or a forced full set of one template.
    for slot in BODY_SLOTS + [16, 2, 11, 12]:
        for material in ([1, 2, 3, 4] if slot in BODY_SLOTS else [1] if slot == 16 else [0]):
            physical = normalized_alloc(BY_ID[279253])
            physical['Strength'] = physical.pop('Agility')
            physical = conservatively_project(physical, 2, slot)
            physical = reduce_stamina(physical, 'HitRating', slot=slot)
            physical = full_budget(physical, slot)
            for archetype, replaced in [
                ('Spell power / crit / hit (matched cost)', 'Strength'),
                ('Strength / spell power / hit (matched cost)', 'CritRating'),
            ]:
                alloc = physical.copy()
                alloc['SpellPower'] = alloc.pop(replaced) / .9
                label = archetype
                reference = 279253
                if slot in BODY_SLOTS and material in (1, 2, 3) and 'Strength' in alloc:
                    # Do not present an Agility->Strength swap on a leather
                    # source as plausible non-plate armor. Mail does have
                    # Strength examples; use its actual hybrid source pattern.
                    if material == 3:
                        reference = 20203
                        alloc = full_budget(normalized_alloc(BY_ID[reference]), slot)
                        label = 'Strength / agility / intellect / MP5 (mail)'
                    elif material == 2:
                        alloc = physical.copy()
                        alloc['SpellPower'] = alloc.pop('Strength') / .9
                        label = 'Spell power / crit / hit (leather)'
                    else:
                        alloc = full_budget(normalized_alloc(BY_ID[272685]), slot)
                        alloc = reduce_stamina(alloc, 'HitRating', slot=slot)
                        reference = 272685
                        label = 'Intellect / spell power / hit (cloth)'
                alloc = full_budget(alloc, slot)
                if slot == 12:
                    alloc = {k: value * .8 for k, value in alloc.items()}
                result.append(synthetic(
                    f'SYN2-{seq:04}', slot, material, label, reference, alloc,
                    '80% passive trinket capacity' if slot == 12 else
                    '20% Stamina; mirrored high-hit modeled cost',
                    'unverified',
                    'Cloth uses a caster off-hand pattern, not verified cloth '
                    'armor; leather/mail use same-material items. Hit on '
                    'leather replaces projected Stamina; source 279253 has '
                    'no Hit. Spell-power prices are unverified.',
                ))
                seq += 1

    # The available q4/65 wand is healing-focused. A lower-level damage wand
    # supplies the stat pattern; its cost is rescaled to the q4/65 ranged cap.
    allocation = full_budget(normalized_alloc(BY_ID[249385]), 26)
    for speed in (1.5, 1.7, 2.0):
        item = synthetic(f'SYN2-{seq:04}', 26, 19, 'Spell-power wand (projected)',
                         279246, allocation, 'Full projected ranged capacity',
                         'unverified', 'Damage-stat pattern from item 249385, '
                         'not the real healing wand; no level-65 damage wand is verified.',
                         speed=speed)
        item['ExtraSourceID'] = 249385
        result.append(item)
        seq += 1

    for speed in (1.6, 2.2):
        item = synthetic(f'SYN2-{seq:04}', 25, 16, 'Thrown weapon (projected)',
                         272599, normalized_alloc(BY_ID[272599]),
                         'Full ranged stat capacity', 'unverified',
                         'Stats and damage modeled from a level-65 gun; '
                         'a matching thrown weapon has not been observed.',
                         speed=speed)
        item['SubclassOverride'] = 16
        result.append(item)
        seq += 1

    # Relic source records have effects rather than stats. Replace those with
    # explicitly hypothetical fixed-stat items at the same group-4 slot cap.
    for reference, archetypes in [
        (279250, ('Physical idol', 'Spell-power idol')),
        (279247, ('Physical libram', 'Spell-power libram')),
        (279249, ('Physical totem', 'Spell-power totem')),
    ]:
        for archetype in archetypes:
            if archetype.startswith('Physical'):
                alloc = full_budget(normalized_alloc(BY_ID[20068]), 28)
            else:
                alloc = full_budget(normalized_alloc(BY_ID[272685]), 28)
            item = synthetic(f'SYN2-{seq:04}', 28, BY_ID[reference]['subclass'],
                             archetype, reference, alloc,
                             'Full projected ranged/relic capacity', 'unverified',
                             'The source relic has a real effect, not these fixed stats. '
                             'No level-65 fixed-stat relic source verifies this allocation.')
            item['ExtraSourceID'] = 20068 if archetype.startswith('Physical') else 272685
            result.append(item)
            seq += 1
    return result


ABBREVIATIONS = {
    'Stamina': 'Sta', 'Strength': 'Str', 'Agility': 'Agi', 'Intellect': 'Int',
    'Spirit': 'Spi', 'CritRating': 'Crit', 'HitRating': 'Hit',
    'AttackPower': 'AP', 'RangedAttackPower': 'RAP',
    'SpellPower': 'SP', 'SpellDamage': 'Dmg', 'HealingPower': 'Heal',
    'FirePower': 'Fire', 'ShadowPower': 'Shadow', 'MP5': 'MP5',
}


def modeled_name(proposal):
    label = f"Modeled: {proposal['Archetype']} — {proposal['Slot']}"
    if proposal['Material'] in ARMOR_TYPES.values():
        label += f" ({proposal['Material']})"
    numbers = [f"+{amount} {ABBREVIATIONS.get(stat, stat)}"
               for stat, amount in proposal['Stats'].items() if amount]
    if proposal['Armor']:
        numbers.append(f"{proposal['Armor']} armor")
    if proposal.get('BaseBlockValue'):
        numbers.append(f"{proposal['BaseBlockValue']} Block")
    if proposal['Speed'] is not None:
        numbers.append(f"{proposal['Min']}–{proposal['Max']} dmg @ {proposal['Speed']:g}s")
    return label + ' [' + ', '.join(numbers) + ']'
