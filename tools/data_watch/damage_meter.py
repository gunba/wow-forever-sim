#!/usr/bin/python

# Reads the client's own damage meter cache, Cache/DamageMeter.bin, and prints what it
# recorded: who did what damage with which ability, straight out of a running game.
#
#   tools/data_watch/damage_meter.py                       # default install path
#   tools/data_watch/damage_meter.py --file <DamageMeter.bin>
#   tools/data_watch/damage_meter.py --scrub out.bin       # copy with player names removed
#
# Why this exists: every other tool here reads a data table, which says what an ability is
# meant to do. This is the only file on the machine that says what the server actually paid
# out, and Forever blocks addons from reading damage, so it is the only one there will be.
#
# Format, worked out by hand against a levelling warrior's file. Records are a name and a
# class as length-prefixed strings followed by seven little-endian uint32s:
#
#   uint16  len          including the trailing NUL
#   char[]  name         actor: a player, a pet, or an NPC
#   uint16  len
#   char[]  class        WARRIOR, PALADIN, ... - players only, absent on NPCs
#   uint32  hits
#   uint32  spellID      6603 is Auto Attack; 772 Rend, 78 Heroic Strike all appear
#   uint32  spellID      repeated, and every record seen so far has them equal
#   uint32  damage
#   uint32  x3           not identified; small, and one tracks overkill closely enough to guess
#
# What this file is NOT safe to share as it stands. It carries the character names of the
# player and of everyone they grouped with, which is why --scrub exists. Scrubbing replaces
# each class-tagged name in place with a same-length placeholder, so every offset in the
# file stays where it was, and every occurrence is replaced rather than only the ones the
# record walk found - a name can appear without a class after it. The result is verified
# before it is written: if any scrubbed name survives, nothing is written at all.
#
# It over-scrubs on purpose. NPCs carry class tokens too - a Stonetusk Boar reads as a
# WARRIOR - so every class-tagged name goes, not only the players'. That loses the record of
# what was being fought, which is a real cost for validating against a target's armour. The
# trade is deliberate: a privacy tool should fail closed, and there is no field here that
# separates a player from an NPC.
#
# What it does NOT remove: names that never appear with a class token, which is how pets
# read. They are weakly identifying and the tool does not claim otherwise.

import argparse
import collections
import os
import struct
import sys

DEFAULT = r'C:\Program Files (x86)\World of Warcraft\_classic_beta_\Cache\DamageMeter.bin'

CLASSES = {
	'WARRIOR', 'PALADIN', 'HUNTER', 'ROGUE', 'PRIEST',
	'SHAMAN', 'MAGE', 'WARLOCK', 'DRUID',
}


def read_string(blob, i):
	"""The string at i as (text, next offset), or (None, i) if there is not one there."""
	if i + 2 > len(blob):
		return None, i
	length = struct.unpack_from('<H', blob, i)[0]
	if length < 2 or i + 2 + length > len(blob):
		return None, i
	raw = blob[i + 2:i + 2 + length]
	if not raw.endswith(b'\x00'):
		return None, i
	try:
		text = raw[:-1].decode('utf-8')
	except UnicodeDecodeError:
		return None, i
	if not text.isprintable():
		return None, i
	return text, i + 2 + length


def records(blob):
	"""Every actor record, as (name offset, name, class, hits, spellID, damage)."""
	out = []
	i = 0
	while i < len(blob) - 4:
		name, after_name = read_string(blob, i)
		if name:
			klass, after_class = read_string(blob, after_name)
			if klass in CLASSES and after_class + 28 <= len(blob):
				hits, spell, spell_again, damage = struct.unpack_from('<4I', blob, after_class)
				out.append((i, name, klass, hits, spell, spell_again, damage))
				i = after_class + 28
				continue
		i += 1
	return out


NUL = b'\x00'


def scrub(blob, rows, path):
	"""A copy with every occurrence of every named actor replaced by a placeholder.

	Two things this has to get right.

	Same length, because the file has offsets in it that were not worked out; changing a
	string's length risks moving something another part of the file points at.

	Every occurrence, not just the one the record walk found. The first version of this
	replaced names at the offsets records() returned and left one behind, because a name
	can appear in the file without a class after it - so the walk never saw that copy. A
	scrubber that removes 61 of 62 copies of a name has not removed it. The names are
	collected from the class-tagged records, then replaced everywhere they occur.
	"""
	names = {}
	for _, name, *_ in rows:
		if name not in names:
			names[name] = 'Player%d' % (len(names) + 1)

	# Matched with the length prefix and the NUL, so only a whole string can match. A plain
	# substring replace turned "Murloc Streamrunner" into "Player Streamrunner" because
	# "Murloc" was itself a scrubbed name.
	out = bytes(blob)
	for name, placeholder in names.items():
		encoded = name.encode('utf-8')
		padded = placeholder.ljust(len(encoded), '.')[:len(encoded)]
		envelope = struct.pack('<H', len(encoded) + 1) + encoded + NUL
		replacement = struct.pack('<H', len(encoded) + 1) + padded.encode() + NUL
		out = out.replace(envelope, replacement)

	leaked = [n for n in names if struct.pack('<H', len(n.encode()) + 1) + n.encode() + NUL in out]
	if leaked:
		sys.exit(f'refusing to write {path}: {len(leaked)} name(s) survived scrubbing')

	with open(path, 'wb') as f:
		f.write(out)
	return names


def main():
	parser = argparse.ArgumentParser(description="What the client's damage meter recorded.")
	parser.add_argument('--file', default=DEFAULT)
	parser.add_argument('--scrub', metavar='OUT', help='write a copy with player names removed')
	args = parser.parse_args()

	if not os.path.exists(args.file):
		sys.exit(f'no damage meter cache at {args.file}; pass --file')
	blob = open(args.file, 'rb').read()
	rows = records(blob)
	print(f'{args.file}\n  {len(blob):,} bytes, {len(rows)} actor records')

	if args.scrub:
		names = scrub(blob, rows, args.scrub)
		print(f'\n  {len(names)} player names replaced, {args.scrub} written')
		print('  pet and NPC names are not class-tagged and are left as they are')
		return

	by_actor = collections.defaultdict(lambda: [0, 0])
	by_spell = collections.defaultdict(lambda: [0, 0])
	for _, name, klass, hits, spell, spell_again, damage in rows:
		if spell != spell_again or not 0 < damage < 10_000_000 or hits > 10_000:
			continue
		by_actor[(name, klass)][0] += damage
		by_actor[(name, klass)][1] += hits
		by_spell[spell][0] += damage
		by_spell[spell][1] += hits

	print('\nby actor:')
	for (name, klass), (damage, hits) in sorted(by_actor.items(), key=lambda kv: -kv[1][0]):
		print(f'  {name:20} {klass:8} {damage:10,} damage  {hits:6,} hits')

	print('\nby ability:')
	for spell, (damage, hits) in sorted(by_spell.items(), key=lambda kv: -kv[1][0]):
		per = damage / hits if hits else 0
		print(f'  spell {spell:<8} {damage:10,} damage  {hits:6,} hits  {per:8.1f} average')


if __name__ == '__main__':
	main()
