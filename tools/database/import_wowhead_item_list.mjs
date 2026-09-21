// Extract a complete Wowhead Listview script captured from the browser.
import { readFileSync, writeFileSync } from 'node:fs';
import assert from 'node:assert/strict';
import JSON5 from 'json5';

const [input, output] = process.argv.slice(2);
if (!input || !output) throw new Error('Usage: node import_wowhead_item_list.mjs capture.js output.json');
const script = readFileSync(input, 'utf8');
const match = script.match(/(?:var\s+)?listviewitems\s*=\s*/);
assert.ok(match, 'Missing complete item Listview data');
const start = match.index + match[0].length;
const items = JSON5.parse(script.slice(start, script.indexOf(';\n', start)));
const metadataStart = script.indexOf('{', script.indexOf('WH.Gatherer.addData(3,'));
const metadataEnd = script.indexOf('});', metadataStart);
const metadata = JSON5.parse(script.slice(metadataStart, metadataEnd + 1));
assert.equal(items.length, 706);
assert.equal(new Set(items.map(item => item.id)).size, 706);
assert.ok(items.every(item => item.level === 65));
const source = 'https://www.wowhead.com/forever/items/min-level:65/max-level:65/quality:3:4:5:6/slot:5:8:11:10:1:23:7:21:2:22:13:15:26:28:14:3:25:12:17:6:9';
writeFileSync(output, JSON.stringify({
	source, count: items.length,
	items: items.map(item => ({ ...item, displayMetadata: metadata[item.id] })),
}, null, 2) + '\n');
console.log(`Exported ${items.length} distinct ilvl-65 records to ${output}`);
