// Reads the client's own damage meter cache and strips the character names out of it.
//
// Cache/DamageMeter.bin is the only file on the machine that records what the server
// actually paid out. It is also the only one that carries the names of the player and of
// everyone they grouped with, which is why it cannot simply be attached to a public issue.
//
// This runs in the browser. The file is read with FileReader and never uploaded, so the
// names are gone before anything leaves the machine. It is a port of
// tools/data_watch/damage_meter.py; the two must agree, and the three mistakes that port
// already made are pinned in scrub.test.ts.
//
// Format, worked out by hand against a levelling warrior's file:
//
//   uint16  len          including the trailing NUL
//   char[]  name         actor: a player, a pet, or an NPC
//   uint16  len
//   char[]  class        WARRIOR, PALADIN, ... - players only, absent on NPCs
//   uint32  hits
//   uint32  spellID      repeated, equal in every record seen so far
//   uint32  spellID
//   uint32  damage
//   uint32  x3           not identified

export const CLASS_TOKENS = new Set(['WARRIOR', 'PALADIN', 'HUNTER', 'ROGUE', 'PRIEST', 'SHAMAN', 'MAGE', 'WARLOCK', 'DRUID']);

export type ActorRecord = {
	name: string;
	className: string;
	hits: number;
	spellId: number;
	spellAgain: number;
	damage: number;
};

/**
 * Whether a record's numbers are worth showing. Not used for scrubbing - a record with a
 * nonsense damage field still carries a real name, and dropping it would leave that name in
 * the file. Only the summary is filtered, so what is shown is what can be believed.
 */
export function plausible(record: ActorRecord): boolean {
	return record.spellId === record.spellAgain && record.damage > 0 && record.damage < 10_000_000 && record.hits > 0 && record.hits < 10_000;
}

export type ScrubResult = {
	scrubbed: Uint8Array;
	records: ActorRecord[];
	namesRemoved: number;
};

/** The length-prefixed string at `offset`, or null if there is not one there. */
function readString(view: DataView, bytes: Uint8Array, offset: number): { text: string; next: number } | null {
	if (offset + 2 > bytes.length) return null;
	const length = view.getUint16(offset, true);
	if (length < 2 || offset + 2 + length > bytes.length) return null;
	if (bytes[offset + 2 + length - 1] !== 0) return null;
	let text: string;
	try {
		text = new TextDecoder('utf-8', { fatal: true, ignoreBOM: true }).decode(bytes.subarray(offset + 2, offset + 1 + length));
	} catch {
		return null;
	}
	if (/[\p{C}\p{Z}]/u.test(text.replace(/ /g, ''))) return null;
	return { text, next: offset + 2 + length };
}

/** Every actor record in the file. */
export function readRecords(bytes: Uint8Array): ActorRecord[] {
	const view = new DataView(bytes.buffer, bytes.byteOffset, bytes.byteLength);
	const out: ActorRecord[] = [];
	let i = 0;
	while (i < bytes.length - 4) {
		const name = readString(view, bytes, i);
		if (name) {
			const className = readString(view, bytes, name.next);
			if (className && CLASS_TOKENS.has(className.text) && className.next + 28 <= bytes.length) {
				const at = className.next;
				const hits = view.getUint32(at, true);
				const spellId = view.getUint32(at + 4, true);
				const spellAgain = view.getUint32(at + 8, true);
				const damage = view.getUint32(at + 12, true);
				// Pushed whatever the numbers look like. Filtering here is what made the browser
				// port scrub 5 names where the Python tool scrubs 24: a record whose damage
				// field reads as nonsense still has a real character name in front of it, and
				// the name is the part that matters. plausible() below filters the summary.
				out.push({ name: name.text, className: className.text, hits, spellId, damage, spellAgain });
				i = at + 28;
				continue;
			}
		}
		i++;
	}
	return out;
}

/**
 * A copy of the file with every named actor replaced by a placeholder.
 *
 * Three things this has to get right, each of which it got wrong first:
 *
 * Every occurrence, not only the ones the record walk found. A name can appear without a
 * class after it, so walking records and replacing at those offsets left one copy of a name
 * behind. Removing 61 of 62 copies of a name has not removed it.
 *
 * Whole strings only. A plain substring replace turned "Murloc Streamrunner" into "Player
 * Streamrunner", because "Murloc" was itself a scrubbed name. Matching includes the length
 * prefix and the trailing NUL, so only a complete string can match.
 *
 * Same length, because the file has offsets in it that were never worked out; changing a
 * string's length risks moving something another part of the file points at.
 *
 * Throws rather than returning a file with a name still in it.
 */
export function scrub(bytes: Uint8Array): ScrubResult {
	const records = readRecords(bytes);
	const placeholders = new Map<string, string>();
	for (const record of records) {
		if (!placeholders.has(record.name)) placeholders.set(record.name, `Player${placeholders.size + 1}`);
	}

	const out = new Uint8Array(bytes);
	for (const [name, placeholder] of placeholders) {
		const byteLength = new TextEncoder().encode(name).length;
		const padded = placeholder.padEnd(byteLength, '.').slice(0, byteLength);
		for (const offset of findStringOffsets(bytes, name)) {
			for (let i = 0; i < padded.length; i++) out[offset + 2 + i] = padded.charCodeAt(i);
		}
	}

	for (const name of placeholders.keys()) {
		if (findStringOffsets(out, name).length > 0) {
			throw new Error(`refusing to hand back a file with ${name.length} characters of a name still in it`);
		}
	}

	return { scrubbed: out, records, namesRemoved: placeholders.size };
}

/** Offsets of the length prefix of every complete occurrence of `text` as a string. */
function findStringOffsets(bytes: Uint8Array, text: string): number[] {
	const encoded = new TextEncoder().encode(text);
	const needle = new Uint8Array(encoded.length + 3);
	new DataView(needle.buffer).setUint16(0, encoded.length + 1, true);
	needle.set(encoded, 2);
	needle[needle.length - 1] = 0;

	const out: number[] = [];
	outer: for (let i = 0; i + needle.length <= bytes.length; i++) {
		for (let j = 0; j < needle.length; j++) {
			if (bytes[i + j] !== needle[j]) continue outer;
		}
		out.push(i);
	}
	return out;
}
