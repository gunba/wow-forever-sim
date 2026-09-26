import assert from 'node:assert/strict';
import { chromium } from 'playwright';

const base = process.env.SITE_URL || 'http://localhost:8080/classic/';
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
try {
	const page = await browser.newPage();
	await page.route(/^https?:\/\/([^/]+\.)?(zamimg|wowhead|googletagmanager)\.com\//, route => route.abort());
	for (const [key, race] of [['tank_warrior', 'orc'], ['protection_paladin', 'human'], ['feral_tank_druid', 'night_elf']]) {
		await page.goto(new URL('review/', base).href, { waitUntil: 'domcontentloaded' });
		assert.equal(await page.locator('#tanks tbody tr').count(), 17);
		const id = `${key}__${race}`;
		const link = page.locator(`a[href$="?profile=${id}"]`);
		assert.equal(await link.count(), 1);
		assert.equal(await link.getAttribute('href'), `../${key}/?profile=${id}`);
		const expected = await (await page.request.get(new URL(`review/profiles/${id}.json`, base).href)).json();
		await link.click();
		await page.getByRole('button', { name: 'Simulate', exact: true }).waitFor();
		await page.waitForFunction(({ key, expected }) => {
			const player = JSON.parse(localStorage.getItem(`__classic_${key}__currentSettings__`))?.player;
			return player?.race === expected.player.race && player?.talentsString === expected.player.talentsString;
		}, { key, expected });
		const stored = await page.evaluate(key => JSON.parse(localStorage.getItem(`__classic_${key}__currentSettings__`)), key);
		for (const field of ['race', 'talentsString', 'equipment', 'rotation', 'healingModel']) {
			assert.deepEqual(stored.player[field], expected.player[field], `${id}: ${field}`);
		}
		assert.deepEqual(stored.encounter, expected.encounter);
		assert.deepEqual(stored.tanks, expected.tanks);
		console.log(`${id}: matrix link loads the exact tank profile`);
	}
} finally {
	await browser.close();
}
