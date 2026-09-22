import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { chromium } from 'playwright';

const base = process.env.SITE_URL || 'http://localhost:8080/classic/';
const bundle = JSON.parse(readFileSync('ui/core/forever_ranked_profiles.json', 'utf8'));
assert.equal(bundle.profiles.length, 147, 'web defaults must cover the complete roster');
const browser = await chromium.launch({ executablePath: process.env.CHROMIUM_PATH || undefined });
try {
	for (const [route, id] of [
		['mage', 'arcane__gnome'],
		['mage', 'fire__orc'],
		['shadow_priest', 'shadow__undead'],
		['balance_druid', 'balance__tauren'],
		['enhancement_shaman', 'enhancement__dwarf'],
		['hunter', 'marksmanship__orc'],
	]) {
		const profile = bundle.profiles.find(profile => profile.id === id);
		assert.ok(profile, `missing ${id}`);
		const context = await browser.newContext();
		const page = await context.newPage();
		await page.route(/^https?:\/\/([^/]+\.)?(zamimg|wowhead|googletagmanager)\.com\//, request => request.abort());
		await page.goto(`${base}${route}/`, { waitUntil: 'networkidle' });
		await page.getByRole('button', { name: 'Simulate', exact: true }).waitFor();
		const storageKey = `__classic_${route === 'enhancement_shaman' ? 'enhacement_shaman' : route}__currentSettings__`;
		await page.waitForFunction(key => localStorage.getItem(key), storageKey);
		const fresh = await page.evaluate(key => JSON.parse(localStorage.getItem(key)), storageKey);
		assert.ok(fresh.player.foreverTier1Bonuses, 'fresh default lacks Tier 1');
		assert.ok(fresh.player.equipment.items.filter(item => item.enchant).length >= 9, 'fresh default lacks enchants');
		assert.equal(fresh.encounter.duration, 300);

		await page.getByLabel('Ranked build', { exact: true }).selectOption(id);
		await page.getByRole('button', { name: 'Load build', exact: true }).click();
		const stored = await page.evaluate(key => JSON.parse(localStorage.getItem(key)), storageKey);
		const expected = profile.settings;
		for (const field of ['equipment', 'race', 'talentsString', 'rotation', 'consumes', 'buffs']) {
			assert.deepEqual(stored.player[field], expected.player[field], `${id}: ${field}`);
		}
		assert.deepEqual(stored.player.bonusStats.stats, expected.player.bonusStats.stats, `${id}: paid hit`);
		assert.deepEqual(stored.raidBuffs, expected.raidBuffs);
		assert.deepEqual(stored.partyBuffs, expected.partyBuffs);
		assert.deepEqual(stored.debuffs, expected.debuffs);
		await page.getByRole('button', { name: 'Simulate', exact: true }).click();
		await page.getByText('Save as Reference', { exact: true }).first().waitFor({ timeout: 180000 });
		const actual = Number(await page.locator('.results-sim-dps .topline-result-avg').first().innerText());
		assert.ok(Math.abs(actual - profile.dps) < .015, `${id}: WASM ${actual} vs native ${profile.dps}`);
		console.log(`${id}: fully loaded from the picker, ${actual} DPS, native match`);

		if (route === 'enhancement_shaman') {
			await page.getByRole('tab', { name: 'Settings', exact: true }).click();
			const critLabel = page.getByText('Crit Aura (3%)', { exact: true });
			assert.ok(await critLabel.isVisible());
			assert.equal(await page.getByText('Moonkin Aura', { exact: true }).isVisible(), false);
			assert.equal(await page.getByText('Leader of the Pack', { exact: true }).isVisible(), false);
			const button = critLabel.locator('..').locator(':scope > a.icon-picker-button');
			await button.click();
			const off = await page.evaluate(key => JSON.parse(localStorage.getItem(key)).raidBuffs, storageKey);
			assert.equal(Boolean(off.moonkinAura || off.leaderOfThePack), false);
			await button.click();
			const on = await page.evaluate(key => JSON.parse(localStorage.getItem(key)).raidBuffs, storageKey);
			assert.equal(on.moonkinAura, true);
			assert.equal(Boolean(on.leaderOfThePack), false);
			console.log('One crit-aura control; off clears both source flags, on selects one');
		}

		const other = bundle.profiles.find(candidate => candidate.key === profile.key && candidate.id !== id);
		await page.evaluate(({ key, settings }) => localStorage.setItem(key, JSON.stringify(settings)),
			{ key: storageKey, settings: other.settings });
		await page.goto(`${base}${route}/?profile=${id}`, { waitUntil: 'networkidle' });
		await page.getByRole('button', { name: 'Simulate', exact: true }).waitFor();
		const linked = await page.evaluate(key => JSON.parse(localStorage.getItem(key)), storageKey);
		assert.deepEqual(linked.player.equipment, expected.player.equipment, `${id}: direct ranking link`);
		assert.equal(linked.player.race, expected.player.race, `${id}: direct-link race`);
		assert.deepEqual(linked.player.bonusStats.stats, expected.player.bonusStats.stats, `${id}: direct-link paid hit`);
		await context.close();
	}
} finally {
	await browser.close();
}
