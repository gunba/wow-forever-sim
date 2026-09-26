import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { chromium } from 'playwright';

const site = process.env.SITE_URL || 'http://localhost:8080/classic/';
const string = text => {
	const bytes = Buffer.from(`${text}\0`, 'utf8');
	const prefix = Buffer.alloc(2);
	prefix.writeUInt16LE(bytes.length);
	return Buffer.concat([prefix, bytes]);
};
const names = ['Al', 'Björn', '李雷'];
const fixture = Buffer.concat(names.flatMap(name => {
	const data = Buffer.alloc(28);
	[2, 1752, 1752, 36].forEach((value, index) => data.writeUInt32LE(value, index * 4));
	return [string(name), string('ROGUE'), data, string(name)];
}));
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
try {
	const page = await browser.newPage();
	const posts = [];
	page.on('request', request => {
		if (request.method() === 'POST') posts.push(request.url());
	});
	await page.goto(new URL('scrub/', site).toString(), { waitUntil: 'networkidle' });
	await page.getByText('Files stay in this browser', { exact: false }).waitFor();
	await page.locator('#meter-drop input').setInputFiles({
		name: 'fixture.bin', mimeType: 'application/octet-stream', buffer: fixture,
	});
	await page.getByText('3 names removed', { exact: true }).waitFor();
	assert.equal(await page.getByRole('button', { name: 'Send it', exact: true }).count(), 0);
	const [download] = await Promise.all([
		page.waitForEvent('download'),
		page.getByRole('link', { name: 'Download locally' }).click(),
	]);
	const cleaned = readFileSync(await download.path());
	assert.equal(cleaned.length, fixture.length);
	for (const name of names) assert.ok(!cleaned.includes(string(name)), `name survived: ${name}`);
	await page.locator('#hotfix-drop input').setInputFiles({
		name: 'DBCache.bin', mimeType: 'application/octet-stream', buffer: Buffer.from('XFTH'),
	});
	await page.getByText('Nothing was uploaded.', { exact: false }).waitFor();
	assert.deepEqual(posts, []);
	console.log('Cache tools: short/UTF-8 names scrubbed; download preserves length; no uploads');
} finally {
	await browser.close();
}
