#!/usr/bin/python

# Reads a spell's numbers out of a client build's DB2 tables, the way the sim needs them: damage range,
# spell power / attack power coefficients, dot ticks, cast time, cost, cooldown, level. Tables come from
# wago.tools as CSV and are cached under tools/data_watch/.cache.
#
#   tools/data_watch/spell_client.py 25306                   # one spell id, Forever beta
#   tools/data_watch/spell_client.py Fireball               # every rank of a spell by name
#   tools/data_watch/spell_client.py Fireball --era         # the same from Classic Era, for comparison
#   tools/data_watch/spell_client.py Fireball --json
#   tools/data_watch/spell_client.py --learned mage          # every spell id the class trains, one line each
#
# Forever reuses some Era spell ids and introduces new ones. SkillLineAbility also retains legacy rune
# entries, so the learned() list alone is not proof of current availability. Id ranges do not distinguish
# NPC spells from player spells: both occur in the 1.3M range.
#
# From Python: `from spell_client import Client; Client().spell(25306)`.
#
# The two builds store a damage range differently. Era: EffectBasePoints + 1 .. EffectBasePoints + EffectDieSides.
# Forever: EffectBasePointsF with Variance, low = base * (1 - v/2), high = base * (1 + v/2). Checked on Fireball
# rank 1, which is 14-22 in both. EffectBonusCoefficient is the spell power coefficient. 1 is a real value on a 3.5 sec
# cast (Fireball rank 12) but is also what effects that never scale with spell power carry (weapon strikes), so
# read it together with the effect type. `Coefficient` is retail's scaling-class column and is
# zero everywhere, do not read it.

import argparse
import collections
import csv
import json
import os
import sys
import urllib.request

FOREVER = '1.60.1.69893'
ERA = '1.15.9.69722'
CACHE = os.path.join(os.path.dirname(os.path.abspath(__file__)), '.cache')
TABLES = ['SpellName', 'Spell', 'SpellEffect', 'SpellMisc', 'SpellCastTimes', 'SpellPower', 'SpellCooldowns',
          'SpellDuration', 'SpellLevels', 'SpellRadius']
CLASS_MASK = dict(warrior=1, paladin=2, hunter=4, rogue=8, priest=16, shaman=64, mage=128, warlock=256, druid=1024)
CLASS_SKILLS = dict(warrior={26, 256, 257}, paladin={594, 267, 184}, hunter={50, 163, 51},
                   rogue={38, 39, 253}, priest={56, 78, 613}, shaman={375, 373, 374},
                   mage={6, 8, 237}, warlock={593, 355, 354}, druid={573, 574, 134})
POWER = {0: 'mana', 1: 'rage', 2: 'focus', 3: 'energy'}


def table(build, name):
	os.makedirs(CACHE, exist_ok=True)
	path = os.path.join(CACHE, f'{name}_{build}.csv')
	if not os.path.exists(path):
		req = urllib.request.Request(f'https://wago.tools/db2/{name}/csv?build={build}', headers={'User-Agent': 'wowsims-forever spell_client'})
		with urllib.request.urlopen(req, timeout=300) as r:
			data = r.read()
		with open(path, 'wb') as f:
			f.write(data)
	with open(path, encoding='utf-8') as f:
		return list(csv.DictReader(f))


def num(v):
	v = float(v or 0)
	return int(v) if v.is_integer() else round(v, 4)


class Client:
	def __init__(self, build=FOREVER):
		self.build = build
		t = {n: table(build, n) for n in TABLES}
		self.names = {int(r['ID']): r['Name_lang'] for r in t['SpellName']}
		self.text = {int(r['ID']): r for r in t['Spell']}
		self.effects = collections.defaultdict(list)
		for r in t['SpellEffect']:
			self.effects[int(r['SpellID'])].append(r)
		self.misc = {int(r['ID']): r for r in t['SpellMisc']}
		misc_by_spell = {}
		for r in t['SpellMisc']:
			if 'SpellID' in r:
				misc_by_spell[int(r['SpellID'])] = r
		self.misc_by_spell = misc_by_spell or self.misc
		self.cast = {r['ID']: int(r['Base']) for r in t['SpellCastTimes']}
		self.duration = {r['ID']: int(r['Duration']) for r in t['SpellDuration']}
		self.radius = {r['ID']: num(r['Radius']) for r in t['SpellRadius']}
		self.power = collections.defaultdict(list)
		for r in t['SpellPower']:
			self.power[int(r['SpellID'])].append(r)
		self.cooldown = {int(r['SpellID']): r for r in t['SpellCooldowns']}
		self.levels = {int(r['SpellID']): r for r in t['SpellLevels']}

	def learned(self, class_name):
		name = class_name.lower()
		mask, skills = CLASS_MASK[name], CLASS_SKILLS[name]
		# New Death ranks have ClassMask=0 but are in Priest's Shadow skill line.
		# Both paths are discovery evidence; neither proves trainer availability.
		return sorted({int(r['Spell']) for r in table(self.build, 'SkillLineAbility')
		               if int(r['ClassMask'] or 0) & mask or int(r['SkillLine']) in skills})

	def ids(self, name):
		return sorted(i for i, n in self.names.items() if n.lower() == name.lower())

	def spell(self, spell_id):
		spell_id = int(spell_id)
		if spell_id not in self.names:
			return None
		misc = self.misc_by_spell.get(spell_id, {})
		duration = self.duration.get(misc.get('DurationIndex'), 0)
		lv = self.levels.get(spell_id, {})
		cd = self.cooldown.get(spell_id, {})
		out = dict(id=spell_id, name=self.names[spell_id], rank=self.text.get(spell_id, {}).get('NameSubtext_lang', ''),
		           level=num(lv.get('SpellLevel') or lv.get('BaseLevel')), castTimeMs=self.cast.get(misc.get('CastingTimeIndex'), 0),
		           durationMs=duration, cooldownMs=num(cd.get('RecoveryTime')), categoryCooldownMs=num(cd.get('CategoryRecoveryTime')),
		           gcdMs=num(cd.get('StartRecoveryTime')), cost=[], effects=[],
		           description=self.text.get(spell_id, {}).get('Description_lang', ''))
		for p in self.power.get(spell_id, []):
			out['cost'].append(dict(power=POWER.get(int(p['PowerType']), p['PowerType']), flat=num(p['ManaCost']), pctOfBase=num(p['PowerCostPct'])))
		for e in sorted(self.effects.get(spell_id, []), key=lambda e: int(e['EffectIndex'])):
			if 'EffectBasePointsF' in e and 'EffectDieSides' not in e:
				base, var = float(e['EffectBasePointsF']), float(e['Variance'] or 0)
				low, high = base * (1 - var / 2), base * (1 + var / 2)
			else:
				base, dice = float(e['EffectBasePoints']), int(e['EffectDieSides'] or 0)
				low, high = (base + 1, base + dice) if dice else (base, base)
			period = int(e['EffectAuraPeriod'] or 0)
			eff = dict(index=int(e['EffectIndex']), effect=int(e['Effect']), aura=int(e['EffectAura']), low=num(round(low)), high=num(round(high)),
			           spCoefficient=num(round(float(e['EffectBonusCoefficient']), 4)), apCoefficient=num(round(float(e['BonusCoefficientFromAP']), 4)),
			           perLevel=num(e['EffectRealPointsPerLevel']), perComboPoint=num(e['EffectPointsPerResource']),
			           periodMs=period, triggerSpell=int(e['EffectTriggerSpell'] or 0), chainTargets=int(e['EffectChainTargets'] or 0),
			           radius=self.radius.get(e['EffectRadiusIndex_0'], 0), miscValue=int(e['EffectMiscValue_0'] or 0))
			if period and duration:
				eff['ticks'] = duration // period
			out['effects'].append(eff)
		return out


def show(s):
	cost = ', '.join(f"{c['flat']} {c['power']}" + (f" ({c['pctOfBase']}% of base)" if c['pctOfBase'] else '') for c in s['cost']) or 'free'
	print(f"{s['id']} {s['name']} {s['rank']}  level {s['level']}  cast {s['castTimeMs'] / 1000:g}s  {cost}"
	      + (f"  cd {s['cooldownMs'] / 1000:g}s" if s['cooldownMs'] else '') + (f"  duration {s['durationMs'] / 1000:g}s" if s['durationMs'] else ''))
	for e in s['effects']:
		bits = [f"e{e['index']} effect {e['effect']}" + (f" aura {e['aura']}" if e['aura'] else ''),
		        f"{e['low']}" + (f"-{e['high']}" if e['high'] != e['low'] else '')]
		if e['spCoefficient']:
			bits.append(f"sp {e['spCoefficient']}")
		if e['apCoefficient']:
			bits.append(f"ap {e['apCoefficient']}")
		if e['periodMs']:
			bits.append(f"every {e['periodMs'] / 1000:g}s x{e.get('ticks', '?')}")
		for k in ('perLevel', 'perComboPoint', 'triggerSpell', 'chainTargets', 'radius'):
			if e[k]:
				bits.append(f'{k} {e[k]}')
		print('    ' + '  '.join(bits))


def main():
	parser = argparse.ArgumentParser(description='Spell numbers from a client build.')
	parser.add_argument('spell', nargs='?', help='spell id or exact spell name')
	parser.add_argument('--learned', metavar='CLASS', help='list class-associated candidate spells, including retained records')
	parser.add_argument('--era', action='store_true', help=f'read Classic Era {ERA} instead of Forever {FOREVER}')
	parser.add_argument('--build')
	parser.add_argument('--json', action='store_true')
	args = parser.parse_args()
	client = Client(args.build or (ERA if args.era else FOREVER))
	if args.learned:
		for i in client.learned(args.learned):
			s = client.spell(i)
			if s:
				print(f"{i}	{s['name']}	{s['rank']}	level {s['level']}")
		return
	ids = [int(args.spell)] if args.spell.isdigit() else client.ids(args.spell)
	spells = [s for s in (client.spell(i) for i in ids) if s]
	if not spells:
		sys.exit(f'no spell {args.spell!r} in {client.build}')
	if args.json:
		print(json.dumps(spells, indent=1))
	else:
		for s in spells:
			show(s)


if __name__ == '__main__':
	main()
