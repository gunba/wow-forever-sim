// Self-check for the browser scrubber. No framework: run it with
//
//   npx tsx ui/scrub/scrub.test.ts
//
// Each case is one of the three mistakes the first version of this made. They are pinned
// here rather than described in a comment because a privacy tool that fails open is worse
// than no privacy tool, and "it looked right" is how it got written the first time.

import assert from 'node:assert';

import { plausible, readRecords, scrub } from './scrub';

/** A length-prefixed, NUL-terminated string, the way the client writes one. */
function str(text: string): number[] {
	const bytes = new TextEncoder().encode(text);
	const length = bytes.length + 1;
	return [length & 0xff, (length >> 8) & 0xff, ...bytes, 0];
}

function u32(value: number): number[] {
	return [value & 0xff, (value >> 8) & 0xff, (value >> 16) & 0xff, (value >>> 24) & 0xff];
}

/** One actor record: name, class, then hits / spell / spell / damage and three unread words. */
function record(name: string, className: string, hits: number, spellId: number, damage: number): number[] {
	return [...str(name), ...str(className), ...u32(hits), ...u32(spellId), ...u32(spellId), ...u32(damage), ...u32(0), ...u32(0), ...u32(0)];
}

function file(...parts: number[][]): Uint8Array {
	return new Uint8Array(parts.flat());
}

// Reads what is there.
{
	const bytes = file(record('Terry Oldman', 'WARRIOR', 3, 772, 120));
	const records = readRecords(bytes);
	assert.strictEqual(records.length, 1);
	assert.deepStrictEqual(records[0], { name: 'Terry Oldman', className: 'WARRIOR', hits: 3, spellId: 772, spellAgain: 772, damage: 120 });
}

// Mistake one: a name that also appears without a class after it must still go. The first
// version replaced only at record offsets and left this copy behind.
{
	const bytes = file(record('Terry Oldman', 'WARRIOR', 1, 6603, 10), str('Terry Oldman'), u32(0));
	const { scrubbed } = scrub(bytes);
	assert.ok(!Buffer.from(scrubbed).includes('Terry Oldman'), 'a copy of the name survived');
}

// Mistake two: a scrubbed name that is a prefix of another string must not corrupt it. The
// first version turned "Murloc Streamrunner" into "Player Streamrunner".
{
	const bytes = file(record('Murloc', 'WARRIOR', 1, 6603, 5), str('Murloc Streamrunner'), u32(0));
	const { scrubbed } = scrub(bytes);
	const text = Buffer.from(scrubbed).toString('latin1');
	assert.ok(!text.includes('Player Streamrunner'), 'a longer string was corrupted by a shorter name');
	assert.ok(text.includes('Murloc Streamrunner'), 'a string that was never a record was rewritten');
}

// The file must not change length, because it holds offsets nobody has worked out.
{
	const bytes = file(record('Bary Oldman', 'PALADIN', 4, 6603, 194));
	const { scrubbed } = scrub(bytes);
	assert.strictEqual(scrubbed.length, bytes.length);
}

// The damage survives the scrub, which is the entire point of sending the file.
{
	const bytes = file(record('Terry Oldman', 'WARRIOR', 87, 6603, 9121), record('Bary Oldman', 'PALADIN', 15, 78, 8258));
	const before = readRecords(bytes).map(r => [r.hits, r.spellId, r.damage]);
	const { scrubbed, namesRemoved } = scrub(bytes);
	const after = readRecords(scrubbed).map(r => [r.hits, r.spellId, r.damage]);
	assert.deepStrictEqual(after, before);
	assert.strictEqual(namesRemoved, 2);
}

// Mistake four, found only by running the browser port against a real file: a record whose
// numbers read as nonsense still has a real character name in front of it. Filtering those
// records out before collecting names made the browser scrub 5 names where the Python tool
// scrubs 24. The summary is filtered; the scrubbing never is.
{
	const junk = [...str('Geosculptor Yip'), ...str('PALADIN'), ...u32(248474), ...u32(1), ...u32(2), ...u32(3815178240), ...u32(0), ...u32(0), ...u32(0)];
	const bytes = file(record('Terry Oldman', 'WARRIOR', 1, 6603, 10), junk);
	const records = readRecords(bytes);
	assert.strictEqual(records.filter(plausible).length, 1, 'the nonsense record should not be summarised');
	const { scrubbed, namesRemoved } = scrub(bytes);
	assert.strictEqual(namesRemoved, 2, "the nonsense record's name must still be scrubbed");
	assert.ok(!Buffer.from(scrubbed).includes('Geosculptor Yip'), 'a name survived because its record looked odd');
}

// Valid short and UTF-8 names must not disappear from the discovery pass.
for (const name of ['Al', 'Björn', '李雷']) {
	const bytes = file(record(name, 'ROGUE', 2, 1752, 36), str(name));
	const { scrubbed, namesRemoved } = scrub(bytes);
	assert.strictEqual(namesRemoved, 1, `missed ${name}`);
	assert.ok(!Buffer.from(scrubbed).includes(Buffer.from(str(name))), `left ${name}`);
	assert.strictEqual(scrubbed.length, bytes.length);
	assert.strictEqual(readRecords(scrubbed)[0].damage, 36);
}

console.log('scrub: all checks passed');
